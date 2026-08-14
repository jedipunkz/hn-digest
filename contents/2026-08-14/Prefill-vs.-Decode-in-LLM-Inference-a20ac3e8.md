---
source: "https://www.parasail.io/blog/prefill-vs-decode-llm-inference"
hn_url: "https://news.ycombinator.com/item?id=49299992"
title: "Prefill vs. Decode in LLM Inference"
article_title: "Prefill vs. decode in LLM inference"
author: "eatonphil"
captured_at: "2026-08-14T15:41:59Z"
capture_tool: "hn-digest"
hn_id: 49299992
score: 1
comments: 0
posted_at: "2026-08-14T15:22:28Z"
tags:
  - hacker-news
  - translated
---

# Prefill vs. Decode in LLM Inference

- HN: [49299992](https://news.ycombinator.com/item?id=49299992)
- Source: [www.parasail.io](https://www.parasail.io/blog/prefill-vs-decode-llm-inference)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T15:22:28Z

## Translation

タイトル: LLM 推論におけるプリフィルとデコード
記事のタイトル: LLM 推論におけるプリフィルとデコード
説明: 最初のトークンまでのプリフィルとデコードのシェーピング時間、ストリーミング パフォーマンス、KV キャッシュの圧力、LLM 提供の決定方法を理解します。

記事本文:
LLM 推論におけるプリフィルとデコード
メイン コンテンツにスキップ 会社概要 モデル 価格 ブログ 採用情報 ドキュメント 営業担当者へのお問い合わせ サインイン 営業担当者へのお問い合わせ サインイン 会社概要 モデル 価格設定 ブログ 採用情報 ドキュメント ガイド LLM 推論におけるプレフィルとデコード
プレフィル、デコード、およびキー/値 (KV) キャッシュの仕組み
2 つのフェーズがレイテンシとスループットにどのように現れるか
サービス提供パターンでトレードオフを管理する
ユーザーが実際にたどるリクエストパスをテストする
最適化する前に位相を測定する
FAQ: プリフィルとデコードの違いは何ですか?
LLM 推論には、モデルが出力を生成する前に入力コンテキストを処理するプレフィルと、一度に 1 つの出力トークンを生成するデコードの 2 つのフェーズがあります。違いは製品の動作に現れます。応答の開始に時間がかかりすぎて、ユーザーが要求を送信した後に空白の一時停止が残る場合や、応答がすぐに開始され、ゆっくりと不均一なバーストで到着する場合があります。
プリフィルとデコードは、推論のさまざまな部分に圧力をかけます。最初のトークンの遅延と遅いストリームにより、異なる測定とサービスの選択が必要になる場合があります。 1 秒あたりの総トークン数では、個々のリクエストがすぐに開始されるか、スムーズにストリーミングされるか、製品が必要な時間内に終了するかどうかを示すことはできません。
プレフィル、デコード、およびキー/値 (KV) キャッシュの仕組み
すべての自己回帰 LLM リクエストにはリクエスト パスがあります。生成を開始する前に、サービス提供システムはシステム命令、取得したコンテキスト、会話履歴、ユーザーのプロンプトをモデルに送信される入力トークン シーケンスに組み立てます。 prefill 中に、モデルはその入力を処理し、キー/値 (KV) キャッシュ、つまり応答の生成に役立つ要求状態を構築します。 decode 中に、モデルはその状態を使用して一度に 1 トークンずつ応答を生成します。
KV キャッシュはモデルの「変換のメモリ」ではありません。

レーション。」リクエストコンテキストの状態を提供しています。入力が処理されるときに作成され、出力トークンが生成されるときに読み取られます。反応が大きくなるにつれて、その状態も大きくなります。
フェーズには異なるアクセス パターンがあります。 Prefill はプロンプト トークンを並行して処理できるため、通常はコンピューティングに依存します。システムは、最初のトークンを発行する前に実行する必要のあるモデル作業がかなりあります。デコードはシーケンシャルです。前のトークンが利用可能になるまで次のトークンは生成できず、各ステップでモデルと KV キャッシュの状態がメモリから繰り返し読み取られます。したがって、一般的な自己回帰サービスでは、デコードは通常、シーケンシャルであると同時にメモリ帯域幅の制約を受けます。 Redis のプリフィルとデコードの概要と WEKA の技術ガイドでも、同様の広範な違いが説明されています。
これは、すべての LLM、ランタイム、またはリクエストが同じように動作するという意味ではありません。モデルのアーキテクチャ、プロンプトの形状、出力の長さ、キューイング、および同時実行性はすべて、ユーザーのエクスペリエンスに影響を与えます。これは、出力までの長い待機と遅いストリームは個別に調査する価値があることを意味します。
2 つのフェーズがレイテンシとスループットにどのように現れるか
適切な指標はさまざまな質問に答えます。
最初のトークンまでの時間 (TTFT) は、ストリーミング応答がいつ開始されるかを測定します。
トークン間遅延 (ITL) は、出力トークン間のギャップと、ストリームが滑らかに感じられるか途切れているように感じられるかを測定します。
エンドツーエンドのレイテンシは、合計完了時間を取得します。
スループットは、時間の経過とともにシステムによって完了した作業の合計を測定します。
TTFT は、便利なリクエスト前信号です。長いプロンプト、大量の取得ペイロード、広範な会話履歴はすべて、事前入力作業を増加させる可能性があります。しかし、TTFT が高いということは、プレフィルが唯一の原因であることを証明するものではありません。キューイング、ネットワーク オーバーヘッド、アドミッション コントロールによっても、最初のトークンが遅延する可能性があります。
ITL は、対応するストリーミング信号です。反応は良いものになる可能性があります

TTFT ですが、生成が始まるとまだ貧弱な感じがします。長い出力、コード生成、および複数ステップのエージェント応答により、トークンの間隔とエンドツーエンドの知覚時間が特に重要になります。
スループットは遅延測定を補完します。それは、別のシステムレベルの質問に答えます。総スループットを向上させる変更は、依然として個々のユーザーのエンドツーエンドの体感時間を損なう可能性があります。このトレードオフは、サービス提供システムがより多くの作業を共有リソースに詰め込むたびに現れます。
症状は出発点であり、最終的な診断ではありません。ストリーミング前の長い待機には、プレフィル、キューイング、またはネットワーク パスが関係する可能性があります。不均一なストリーミングには、デコード、キャッシュの圧力、または他の作業との競合が関係する可能性があります。
遅延要件とトラフィック形状を推論モードに一致させるためのより広範なフレームワークについては、「適切なマネージド推論アーキテクチャを選択する方法」を参照してください。
ワークロード カテゴリは、測定計画の開始点として役立ちます。同じアプリケーションは、トラフィック、モデルの選択、プロンプト構築、同時実行性の変更に応じて、プレフィルとデコードのプレッシャーの間を移動できます。
長いコンテキストと取得量の多いリクエストでは、最初に入力トークンの配布、TTFT、および末尾の動作に進む必要があります。 RAG アプリケーションは、取得した大量の文章とともに短いユーザーの質問を送信する場合があります。チャット アプリケーションでは、長い会話履歴が蓄積される場合があります。どちらの場合も、リクエストには UI で明らかなよりもはるかに多くの入力作業が必要になる可能性があります。
長い出力リクエストの場合は、ITL、出力長の分散、および合計完了時間に移動する必要があります。コード生成、ワークフローの作成、および作業を説明するエージェントは、認識される時間のほとんどをデコードに費やす可能性があります。
同時実行下の混合トラフィックには両方のビューが必要です。システムが 1 つのリクエストに合わせて調整されている間に、大量の受信プリフィルがキュー内の他のリクエストをブロックする可能性があります。

実際のディストリビューションが到着すると、シェイプのパフォーマンスが異なる可能性があります。 1 つの平均プロンプトと 1 つの平均応答ではなく、代表的な入力長、出力長、到着パターンを使用して p95 と p99 の動作を測定します。
サービス提供パターンでトレードオフを管理する
チームがレイテンシーが発生する場所を特定したら、次に決定するのは、応答性、全体的な効率、コスト、運用の複雑さの間のトレードオフです。
スケジューリングと実行中のバッチ処理により、システムが完了する作業量を改善できますが、パッキングを増やすとリクエストごとのレイテンシーが変化する可能性があります。適切なポリシーは、製品が TTFT、ストリーミングのスムーズさ、合計完了時間、またはその組み合わせのどれに対してより敏感であるかによって異なります。
チャンク プリフィルは、大きな入力をより小さな部分に分割して処理するため、アクティブなデコード作業をインターリーブできます。これにより、大量の受信プロンプトとすでにストリーミングされている応答の間の干渉を軽減できます。その効果は、リクエストの組み合わせとスケジューリング ポリシーによって異なります。
KV キャッシュの再利用と管理は、メモリ負荷と同時容量に影響します。利点は、リクエストの類似性、キャッシュの動作、およびサービスの実装によって異なります。
投機的デコードでは、デコードの順次パスを別の方法で処理します。ドラフトメカニズムは次のトークンの候補を予測します。メインモデルはそれらを検証します。予測が受け入れられると、システムは同じ数の連続したデコード ステップに料金を支払うことなく、複数のトークンを進めることができます。デコードの負荷が高いワークロードについては評価する価値があります。結果は、モデルのペア、リクエスト パターン、およびサービスの実装によって異なります。
細分化されたサービスは、プリフィル リソースとデコード リソースを分離して、個別に調整およびスケーリングできるようにする業界アーキテクチャ パターンです。これは、フェーズの競合が継続してデータ転送と運用の複雑さの増加を正当化する場合に関連する可能性があります。パラセール

は現在、細分化されたサービスをサポートしていません。そのため、これは、より広範なサービス環境を評価する際に理解すべき概念であり、現在の製品の機能を表すものではありません。
位相分離により、2 つのプールが異なる GPU プロファイルを使用できるようになります。これは概念的な業界パターンであり、Parasail 構成ではありません。KV キャッシュ状態を転送するワークロードとコストは、追加の複雑さを正当化する必要があります。
バッチ処理、キャッシュ管理、および推論能力を準備しておく経済性の詳細については、「アイドル GPU 税」を参照してください。
ユーザーが実際にたどるリクエストパスをテストする
重要なベンチマークは、実稼働リクエストのパスに似たベンチマークです。最適化を選択する前に、以下を把握してください。
プロンプト長と出力長の分布。
同時実行性とバースト動作。
ストリーミング製品パスと非ストリーミング製品パス。
キャッシュの動作とリクエストの類似性。そして
TTFT、ITL、エンドツーエンドのレイテンシー、スループット、および p95/p99 値。
これは、ヘッドラインのスループット数値を完全な診断として扱うことを避ける方法です。最初のトークンが遅いチームは、スムーズなスタートと不均一なストリームのチームとは異なる次の質問を必要とします。負荷がかかっている場合にのみメトリクスが低下するチームは、アーキテクチャを変更する前に、ワークロードの組み合わせとキューイング パスを理解する必要があります。
製品が実際に受け取るリクエストミックスと比較してください。パーセンタイル、プロンプトの長さ、出力の長さ、同時実行レベル、および測定ウィンドウを結果にまとめて保持します。同時実行性が低い短いプロンプトで測定された TTFT 数値だけでは、トラフィックのピーク時の取得負荷の高いワークフローに関する疑問を解決できません。同じことが ITL にも当てはまります。スムーズな単一ストリームでは、複数の長い出力がアクティブなときにシステムがどのように動作するかが確立されません。
この枠組みは、計画において有用な区別を視覚的に維持することにもなります。あたり

-ストリーム測定により、あるユーザーのリクエストが応答していると感じられるかどうかがわかります。総スループットは、システムが完了する作業量を示します。どちらも重要ですが、バッチ サイズ、スケジュール、競合の変化に応じて、異なる方向に進む可能性があります。
待機中のライブ ユーザーがいない場合は、評価も変わります。バッチ作業では、インタラクティブなレイテンシーよりもスループットと有用な結果ごとのコストを優先できます。ユーザーが待機している場合は、実際に見られる動作から始めて、それをリクエスト パスを通じて追跡します。
最適化する前に位相を測定する
プリフィルとデコードは 1 つの推論リクエストの 2 つの部分であり、パフォーマンスに別個の影響を及ぼします。プレフィルは、いつ出力を開始できるかを決定するのに役立ちます。適切なサービス提供の決定は、ワークロードが実際にさらす制約に基づいて決定されます。
プロンプトと出力のディストリビューション、同時実行パターン、ターゲットのレイテンシー動作、モデル構成を会話に取り入れます。ご相談ください。
FAQ: プリフィルとデコードの違いは何ですか?
Prefill は、LLM リクエストの入力コンテキストを処理し、生成に使用される KV キャッシュ状態を作成します。次に、デコードは、その状態を使用して、一度に 1 トークンずつ出力を順番に生成します。プレフィルは通常、コンピューティングに依存します。デコードは通常、メモリ帯域幅に制限されており、前のトークンに対する各トークンの依存関係によって制約されます。この区別は、チームが遅延の症状を最初に測定する価値のある推論の部分と結び付けるのに役立ちます。
当社がどのようにお手伝いできるかについては、当社の製品専門家にお問い合わせください。
製品のアップデート、エンジニアリングの詳細な調査、チームからのソート リーダーシップ。
Parasail が推論操作用に 1 つの AI エージェントを構築した方法
難しかったのは、Slack に AI エージェントを配置することではありませんでした。信頼できる情報源、レビューされた計算、安全な推論操作のワークフローが埋め込まれていました。
globa 向けの 1 秒未満の LLM 推論の構築

l AIトラフィック
高速モデルは高速 API ではありません。ここでは、Kubernetes、グローバル キャパシティ、フォールト トレランスを放棄することなく、600 ミリ秒の p99 予算 (エッジで Cloudflare ワーカー、GPU に直接 WireGuard を接続) から逆算して作業した方法を示します。
Parasail は NVIDIA AI インフラストラクチャと d-Matrix アクセラレータを組み合わせて 10 倍高速なトークン生成を実現します
Parasail は、NVIDIA Hopper および Blackwell GPU と d-Matrix Corsair アクセラレータを組み合わせることで、より高速でコスト効率の高いトークンを提供します
Parasail チームによる実践的な比較、ウォークスルー、ベスト プラクティス。
ゲートウェイ アーキテクチャを使用したマルチリージョン LLM デプロイメントの設計
マルチプロバイダー LLM ゲートウェイがリージョン間のルーティング、フェイルオーバー、待ち時間をどのように解決するかを学びます。実稼働推論のためのセルフホスト型オプションとマネージド型オプションを比較します。
Parasail と DeepInfra: コスト、遅延、モデルの選択、エンジニアリング サポート
サーバーレスおよび専用の価格設定、エンジニアリング サポート、モデル カタログ、コールド スタート、およびコンプライアンスに関する Parasail と DeepInfra の比較。
2026 年のオフライン LLM 推論のバッチ処理とスループットを最適化するためのガイド
オフライン LLM バッチ推論用に vLLM、TensorRT-LLM、および SGLang を調整します。 2026 年の構成ノブ、KV キャッシュ計算、プレフィックス キャッシュ、および投機的デコードのトレードオフ。

[切り捨てられた]

## Original Extract

Understand how prefill and decode shape time to first token, streaming performance, KV-cache pressure, and LLM serving decisions.

Prefill vs. decode in LLM inference
Skip to main content About us Models Pricing Blog Careers Docs Contact sales Sign in Contact sales Sign in About us Models Pricing Blog Careers Docs Guides Prefill vs. decode in LLM inference
How prefill, decode, and the key-value (KV) cache work
How the two phases show up in latency and throughput
Serving patterns manage the trade-off
Test the request path your users actually take
Measure the phase before optimizing it
FAQ: What’s the difference between prefill and decode?
LLM inference has two phases: prefill, when the model processes the input context before it can produce output, and decode, when it generates that output one token at a time. The distinction becomes visible in product behavior. A response may take too long to start, leaving a blank pause after a user sends a request, or it may begin promptly and arrive in slow, uneven bursts.
Prefill and decode put pressure on different parts of inference. A first-token delay and a slow stream can call for different measurements and serving choices; aggregate tokens per second can't show whether an individual request starts promptly, streams smoothly, or finishes within the time your product needs.
How prefill, decode, and the key-value (KV) cache work
Every autoregressive LLM request has a request path. Before generation can begin, the serving system assembles system instructions, retrieved context, conversation history, and the user’s prompt into the input token sequence sent to the model. During prefill , the model processes that input and builds the key-value (KV) cache: request state that helps it generate the response. During decode , the model produces the response one token at a time using that state.
The KV cache isn't the model’s “memory of the conversation.” It's serving state for the request context. It's created as the input is processed, then read as output tokens are generated. As the response grows, that state grows too.
The phases have different access patterns. Prefill can process prompt tokens in parallel, so it's typically compute-bound: the system has substantial model work to perform before it can emit the first token. Decode is sequential. The next token cannot be generated until the previous one is available, and each step repeatedly reads model and KV-cache state from memory. In common autoregressive serving, decode is therefore typically memory-bandwidth-bound as well as sequential. Redis’s overview of prefill and decode and WEKA’s technical guide describe the same broad distinction.
That doesn't mean every LLM, runtime, or request behaves identically. Model architecture, prompt shape, output length, queueing, and concurrency all affect what a user experiences. It does mean that a long wait before output and a slow stream deserve separate investigation.
How the two phases show up in latency and throughput
The right metrics answer different questions:
Time to first token (TTFT) measures when a streamed response starts.
Inter-token latency (ITL) measures the gap between output tokens and whether the stream feels smooth or stuttered.
End-to-end latency captures total completion time.
Throughput measures total work completed by the system over time.
TTFT is a useful front-of-request signal. Long prompts, large retrieval payloads, and extensive conversation history can all increase prefill work. But high TTFT doesn't prove prefill is the sole cause. Queueing, network overhead, and admission control can also delay the first token.
ITL is the corresponding streaming signal. A response can have a good TTFT and still feel poor once generation starts. Long outputs, code generation, and multi-step agent responses make token cadence and end-to-end perceived time especially important.
Throughput complements the latency measures; it answers a different, system-level question. A change that improves aggregate throughput can still hurt an individual user’s end-to-end perceived time. That trade-off shows up whenever a serving system packs more work onto shared resources.
The symptom is a starting point, not a final diagnosis. A long wait before streaming can involve prefill, queueing, or the network path. Uneven streaming can involve decode, cache pressure, or contention with other work.
For a broader framework on matching latency requirements and traffic shape to an inference mode, see How to choose the right managed inference architecture .
Workload categories are useful starting points for a measurement plan. The same application can move between prefill and decode pressure as traffic, model choice, prompt construction, and concurrency change.
Long-context and retrieval-heavy requests should send you first to the input-token distribution, TTFT, and tail behavior. A RAG application may send a short user question alongside a large set of retrieved passages. A chat application may accumulate long conversation history. In both cases, the request can have far more input work than the UI makes obvious.
Long-output requests should send you to ITL, output-length distribution, and total completion time. Code generation, drafting workflows, and agents that explain their work can spend most of their perceived duration in decode.
Mixed traffic under concurrency requires both views. A large incoming prefill can block other requests in the queue, while a system tuned for one request shape can perform differently when the real distribution arrives. Measure p95 and p99 behavior with representative input lengths, output lengths, and arrival patterns—not one average prompt and one average response.
Serving patterns manage the trade-off
Once a team has identified where latency appears, the next decision is a trade-off among responsiveness, aggregate efficiency, cost, and operational complexity.
Scheduling and in-flight batching can improve how much work a system completes, but more packing can change per-request latency. The right policy depends on whether your product is more sensitive to TTFT, streaming smoothness, total completion time, or a combination.
Chunked prefill processes a large input in smaller pieces so active decode work can be interleaved. That can reduce interference between a large incoming prompt and an already-streaming response. Its effect depends on the request mix and scheduling policy.
KV-cache reuse and management affect memory pressure and concurrent capacity. The benefit depends on request similarity, cache behavior, and the serving implementation.
Speculative decoding addresses decode’s sequential path differently. A draft mechanism predicts candidate next tokens; the main model verifies them. When predictions are accepted, the system can advance multiple tokens without paying for the same number of sequential decode steps. It's worth evaluating for decode-heavy workloads; the outcome depends on the model pair, request pattern, and serving implementation.
Disaggregated serving is an industry architecture pattern that separates prefill and decode resources so they can be tuned and scaled independently. It can be relevant when sustained phase contention justifies the added data-transfer and operational complexity. Parasail doesn't currently support disaggregated serving, so this is a concept to understand when evaluating the wider serving landscape—not a statement of current product capability.
Phase separation can also allow the two pools to use different GPU profiles. This is a conceptual industry pattern, not a Parasail configuration: the workload and the cost of transferring KV-cache state have to justify the additional complexity.
For more context on batching, cache management, and the economics of keeping inference capacity ready, read The idle GPU tax .
Test the request path your users actually take
The benchmark that matters is the one that resembles your production request path. Before choosing an optimization, capture:
prompt-length and output-length distributions;
concurrency and burst behavior;
streaming versus non-streaming product paths;
cache behavior and request similarity; and
TTFT, ITL, end-to-end latency, throughput, and p95/p99 values.
This is how to avoid treating a headline throughput number as a full diagnosis. A team with a slow first token needs a different next question from a team with smooth starts and uneven streams. A team whose metrics degrade only under load needs to understand the workload mix and queueing path before changing architecture.
Run the comparison against the request mixes your product actually receives. Keep the percentile, prompt length, output length, concurrency level, and measurement window together in the result. A TTFT number measured on short prompts at low concurrency can't settle a question about a retrieval-heavy workflow at peak traffic. The same applies to ITL: a smooth single stream doesn't establish how the system behaves when several long outputs are active.
This framing also keeps a useful distinction visible in planning. Per-stream measurements answer whether one user’s request feels responsive. Aggregate throughput answers how much work the system completes. Both matter, but they can move in different directions as batch size, scheduling, and contention change.
If no live user is waiting, that changes the evaluation too. Batch work can prioritize throughput and cost per useful result over interactive latency. If a user is waiting, start with the behavior they actually see, then trace it back through the request path.
Measure the phase before optimizing it
Prefill and decode are two parts of one inference request with separate performance consequences. Prefill helps determine when output can begin. The right serving decision follows from the constraint your workload actually exposes.
Bring your prompt and output distributions, concurrency pattern, target latency behavior, and model configuration to the conversation. Talk to us.
FAQ: What’s the difference between prefill and decode?
Prefill processes an LLM request’s input context and creates the KV-cache state used for generation. Decode then generates the output sequentially, one token at a time, using that state. Prefill is typically compute-bound; decode is typically memory-bandwidth-bound and constrained by each token’s dependency on the prior one. The distinction helps teams connect latency symptoms to the part of inference worth measuring first.
Connect with our product experts to see how we can help.
Product updates, engineering deep dives, and thought leadership from the team.
How Parasail built one AI agent for inference operations
The hard part wasn’t putting an AI agent in Slack. It was embedding trusted sources, reviewed calculations, and safe inference-operations workflows.
Building sub-second LLM inference for global AI traffic
A fast model isn't a fast API. Here's how we worked backward from a 600ms p99 budget — Cloudflare Workers at the edge, WireGuard straight to the GPU — without giving up Kubernetes, global capacity, or fault tolerance.
Parasail to combine NVIDIA AI infrastructure with d-Matrix accelerators to achieve 10x faster token generation
Parasail to deliver faster, more cost-efficient tokens by pairing NVIDIA Hopper and Blackwell GPUs with d-Matrix Corsair accelerators
Practical comparisons, walkthroughs, and best practices from the Parasail team.
Designing multi-region LLM deployments with gateway architecture
Learn how multi-provider LLM gateways solve routing, failover, and latency across regions. Compare self-hosted vs. managed options for production inference.
Parasail vs. DeepInfra: Cost, latency, model selection, and engineering support
A side-by-side comparison of Parasail and DeepInfra on serverless and dedicated pricing, engineering support, model catalog, cold starts, and compliance.
Guide to optimize batching and throughput for offline LLM inference in 2026
Tune vLLM, TensorRT-LLM, and SGLang for offline LLM batch inference. Config knobs, KV cache math, prefix caching, and speculative decoding tradeoffs for 2026.

[truncated]
