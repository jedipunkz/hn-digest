---
source: "https://www.modular.com/blog/why-llm-inference-needs-a-new-kind-of-router-part-1"
hn_url: "https://news.ycombinator.com/item?id=48451594"
title: "Why LLM Inference Needs a New Kind of Router"
article_title: "Modular: Why LLM Inference Needs a New Kind of Router - Part 1"
author: "aviziva"
captured_at: "2026-06-08T21:38:09Z"
capture_tool: "hn-digest"
hn_id: 48451594
score: 1
comments: 0
posted_at: "2026-06-08T20:37:04Z"
tags:
  - hacker-news
  - translated
---

# Why LLM Inference Needs a New Kind of Router

- HN: [48451594](https://news.ycombinator.com/item?id=48451594)
- Source: [www.modular.com](https://www.modular.com/blog/why-llm-inference-needs-a-new-kind-of-router-part-1)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T20:37:04Z

## Translation

タイトル: LLM 推論に新しい種類のルーターが必要な理由
記事のタイトル: モジュラー: LLM 推論に新しい種類のルーターが必要な理由 - パート 1
説明: HTTP ルーティングは、長年にわたって解決されてきた問題です。ラウンドロビン、一貫したハッシュ、最小接続。 1 つを選択し、同一のサーバーのプールの前に配置すると、トラフィックがほぼ均等に分散されます。

記事本文:
モジュラー: LLM 推論に新しい種類のルーターが必要な理由 - パート 1
ヒポクラティック AI + モジュラーでリアルタイムの患者会話を強化します。続きを読む→
API 経由でフロンティア モデルにアクセスする
テキストプロンプトから画像を生成する
本番環境に対応したコードを生成する
テキスト+画像からビデオを生成
実際の顧客からの実証済みの結果
GenAI のネイティブ モデリングと提供
最高の GPU と CPU パフォーマンス
AI の未来を一緒に構築しましょう
Modular の公式 AI エージェント スキル
GenAI モデル、当社のクラウド、またはお客様のクラウドをデプロイします
CPU および GPU 用の高性能カーネルを作成する
AI の未来を一緒に構築しましょう
誰でも、どこでも使える AI を構築します。
LLM 推論に新しい種類のルーターが必要な理由 - パート 1
HTTP ルーティングは、長年にわたって解決されてきた問題です。ラウンドロビン、一貫したハッシュ、最小接続。 1 つを選択し、同一のサーバーのプールの前に配置すると、トラフィックがほぼ均等に分散されます。
しかしその後、大規模言語モデルが登場しました。
ここでのバックエンドは、交換可能な Web サーバーではありません。これらは、高帯域幅の RAM または SSD メモリに大規模なローカル KV キャッシュを保持する GPU ポッドです。この状態は再構築にコストがかかり、クラスター全体で均一に利用できるわけではなく、多くの場合、リクエストがすぐに返されるか、以前の作業の再計算に数秒かかるかが決まります。一部のポッドはプレフィルに特化し、他のポッドはデコードに特化する場合があります。通常、会話は複数のリクエストにわたって行われます。 1 つの推論呼び出しで、連続して 2 つのバックエンドが必要になる場合があります。 「交換可能なバックエンド」と「独立したリクエスト」に関する古い前提は、これらの要件をサポートしていません。
従来のルーティングでは、これらすべてが無視されます。すべてのバックエンドを交換可能として扱い、すべてのリクエストを独立したものとして、すべてのポッドを同様に優れたものとして扱います。 GPU ポッドはそのどれでもありません。これらはステートフルで、特化されており、異種混合です。推論ルーティングではそれを考慮する必要があります。
これがフィです

推論ワークロードを処理するためにルーティングがどうなる必要があるかについての 3 部構成シリーズの最初の投稿です。 Modular Cloud のオーケストレーション層はこのルーティング問題を中心に構築されており、このシリーズではそれがどのように解決されるかを説明します。
HTTP 時代の負荷分散は、特定の導入形態に合わせて調整された小さなアルゴリズムのメニューに基づいて構築されています。ポリシーは異なりますが、前提条件はステートレス バックエンドという同じです。
ラウンドロビンは、同一のバックエンドのプール全体にリクエストを均一に分散します。すべてのバックエンドがすべてのリクエストを同じコストで処理することを前提としています。これは、ロード バランサーの背後に同じ Web サービスの 8 つのレプリカがあり、それぞれがトラフィックの 12.5% を取得しているように見えます。シンプル、公平、無国籍です。
一貫したハッシュでは、リクエストの一部のプロパティ (キー、URL、セッション識別子) をハッシュし、ハッシュ リング上の位置が最も近いバックエンドを選択することによって決定されるバックエンドに各リクエストがルーティングされます。これは、クライアント側のキャッシュまたはセッション アフィニティのために、同じキーを同じバックエンドに配置する場合に選択するルーティング戦略です。バックエンドの「スティッキー性」は、バックエンドがメモリに保持しているものではなく、リクエスト キーの関数です。
最小リクエストでは、アクティブなリクエストが少ないということは、より多くの空き容量があることを前提として、アクティブなリクエストが最も少ないバックエンドに各新しいリクエストが送信されます。すべてのリクエストにほぼ同じ量の作業がかかる場合に機能します。
これらのポリシーは、次の 3 つの同じ前提を共有しています。
どのバックエンドでもあらゆるリクエストを処理できます。割り当てはポリシーの選択であり、正しさではありません。
リクエストは独立しています。リクエスト N で何が起こっても、リクエスト N+1 で行うべきことは変わりません。
バックエンドは交換可能です。ロード バランサは、クライアントに気付かれずにロード バランサを別のロード バランサに交換できます。
これらの前提はステートレス Web サービスにも当てはまります。 LLM i

推論は 3 つすべてを破ります。
LLM ワークロードは、4 つの具体的な方法でステートレスの前提に違反します。それぞれのルーティングには、従来のルーティングでは処理できないメカニズムが導入されています。
ポッドが推論リクエストを処理するとき、フォワード パスは KV キャッシュを構築します。これは、GPU メモリに保持される、トークン位置ごとのモデルの中間状態です。最新のエンジンは応答が完了した後もそのキャッシュを保持するため、プレフィックスを共有する後のリクエストでは同等の計算をスキップできます。
これにより、ルーティングの問題が大幅に変わります。最初の 75,000 トークンがすでにキャッシュされているポッドに到達する 100,000 トークン プロンプトは、ミリ秒単位で事前入力できます。同じプロンプトがコールド ポッドに到達するまでには数秒かかります。キャッシュ状態を意識しないラウンドロビンでは、同一のリクエストに対して予測不可能な最初のトークンまでの時間 (TTFT) が発生します。
キャッシュ状態は、大規模なプリフィル レイテンシの差異の主な要因です。キャッシュの常駐性に基づいてポッドを選択するルーターでは、ヒットごとの共有プレフィックス長に比例するプレフィル計算が不要になります。これにより、クラスターがすでに実行した作業の再計算に費やしていた GPU サイクルが解放されます。
LLM 推論には 2 つのフェーズがあり、ハードウェアに与えるストレスが異なります。
Prefill はプロンプト全体を並行して処理します。それはコンピューティングに依存しています。これは、一度に数千のトークンにわたる高密度行列の乗算を行うことで GPU コアが飽和状態になることを意味します。
デコードは、自己回帰的に一度に 1 つずつトークンを生成し、各トークンはその前のすべてのトークンに依存します。メモリ帯域幅の制限があります。これは、GPU の時間のほとんどが高帯域幅メモリ (HBM) からモデルの重みと KV キャッシュをフェッチするのに費やされ、コンピューティングのほとんどがアイドル状態になることを意味します。
同じポッド上で両方のフェーズを実行するということは、ハードウェアがどちらにも調整されていないことを意味します。プレフィルには高密度のコンピューティングが必要です。デコードにはメモリ帯域幅が必要です。 1 つの用途に最適化されたポッドは、機能を十分に活用していない

もう一方が必要とするものではありません。細分化されたデプロイでは、フェーズごとに個別に調整されたポッドが使用されます。単一のクライアント要求により、両方の作業が分割されます。
最新のエンジンは、チャンク化されたプレフィルを使用して同じポッド上の 2 つのフェーズをインターリーブし、境界を曖昧にします。ただし、基礎となるコンピューティングと帯域幅の区別は依然として維持されており、デプロイメント レベルで集約する場合、ルーターはどのポッドが何を実行できるかを認識する必要があります。
ほとんどの LLM トラフィックはマルチターンです。ユーザーがメッセージを送信します。助手が答える。ユーザーが別のメッセージを送信すると、そのメッセージには会話履歴全体がコンテキストとして暗黙的に含まれます。
ターン N+1 はターン N とプレフィックスを共有します。つまり、システム プロンプト、以前のすべてのターン、以前のすべてのアシスタントの応答です。ターン N の KV キャッシュがまだ何らかのポッドに存在する場合、ターン N+1 は事実上、共有部分を自由に事前入力できます。キャッシュが削除された場合、またはターン N+1 が別のポッドに到達した場合、共有プレフィックスは最初から再計算されます。
HTTP におけるセッション アフィニティは、かつては「アプリケーションがメモリ内状態を使用できるように、このユーザーのリクエストを同じバックエンドにルーティングする」ことを意味していました。 LLM 推論では同じことを意味しますが、メモリ内の状態は KV キャッシュです。正しく行うかどうかは、最初のターン以降の各ターンにおける 1 秒未満の応答と数秒間の応答の違いです。
単一のクライアント向けリクエストに複数のバックエンドが必要な場合があります。
分離されたデプロイメントでは、プレフィル ポッドが KV キャッシュを構築し、デコード ポッドがトークンを生成します。どちらも単独ではリクエストに対応できません。ルーターは、事前入力ポッドを選択し、次にデコード ポッドを選択して、シーケンスを調整します。つまり、事前入力するプロンプトを送信し、完了を待ち、同じプロンプトとデコードするキャッシュ ヒントを送信し、トークンをクライアントにストリーム返します。
HTTP ロードバランサーはこれを行いません。リクエストごとに 1 つのバックエンドを選択します。複数ステップの coo の追加

シングルディスパッチルーターへのルーティングは、まったく異なる形式のルーティングです。
上記の 4 つの側面はそれぞれ、ルーティング システムに要件を課します。これらの要件は 3 つの異なるアーキテクチャ上の懸念事項に分類され、それぞれが別個のレイヤーによって処理されます。
この層は、ルーティングの決定に必要なレイテンシで LLM 固有の状態を追跡します。 「N 個のポッドのうち、これらのブロックがキャッシュされているのはどれですか?」という質問同時更新下でマイクロ秒単位で応答可能であり、ポッドのチャーンに対する回復力がなければなりません。ミューテックスを含むハッシュマップだけでは十分ではありません。
この層は、ルーティング ロジックを、テスト可能で再利用可能な小さなコンポーネントの構成として表現します。オペレーターはフィルター、数人のスコアラー、ピッカーを選択し、プロファイルを組み立てます。フレームワークは、午前 3 時のトラフィック下ではなく、ビルド時に構成を検証します。
この層は、決定層の上で複数ステップのリクエスト フローを調整します。シングルディスパッチルーティングは、マルチステップの縮退ケースです (1 ポッド、1 ステップ)。細分化されたプリフィル/デコードは一般的なケースです。2 つのポッド、2 つのステップで、最初の決定によって 2 番目の決定が通知されます。同じフレームワークは、バリアントごとに新しい HTTP ハンドラーを必要とせずに両方を処理します。
このシリーズのパート 2 とパート 3 では、これらのレイヤーを構築します。
このシリーズでは、Modular Cloud の分散推論フレームワーク内のルーティング層と、実稼働推論ワークロードにおけるこれら 4 つの問題のそれぞれをルーティング層がどのように処理するかについて説明します。
プレフィックス認識ルーティング (トークン化、ブロックレベルのハッシュ、負荷認識タイブレークによるキャッシュ認識スコアリング、アップストリーム レイテンシーのサーキット ブレーカー) は、新しいアルゴリズムではなく、プロファイル構成として出荷されます。チームが新しいルーティング動作を必要とする場合、新しいルーティング戦略を最初から作成するのではなく、プラグインを新しいプロファイルに組み込む作業が必要になります。新しい展開パターンはそれぞれ、既存のものを再利用します。
LLM は 4 つの次元を導入しました

従来のロード バランサーには、バックエンドの選択をパフォーマンス クリティカルな決定にする KV キャッシュ状態、単一のリクエストを複数のポッド タイプに分割するハードウェアの特化、セッションをキャッシュ常駐に結び付ける会話の継続性、バックエンドを 1 つ選択するのではなく一連のバックエンドを調整する必要があるマルチステップ実行などを処理するメカニズムがありません。
この問題はさまざまな角度から取り組まれてきました。 NVIDIA Dynamo、llm-d、vLLM プロダクション スタック、AIBrix、KServe、および Envoy AI ゲートウェイには、それぞれ異なる方向の高度な推論ルーティングがあります。つまり、分離されたプレフィル/デコード、プレフィックス認識スケジューリング、KV 認識ロード バランシング、プロダクション グレードのサービング プリミティブです。モジュラー クラウドはその基盤の上に構築されます。対象とするさまざまなデプロイメント パターンをサポートするために、Modular Cloud はコンポーザブル プラグインとマルチステップ実行の両方を最上級のプリミティブにするため、新しいデプロイメント パターンは、フォークする戦略ではなく、ユーザーが組み立てるプロファイルになります。
これが、Modular Cloud のルーティング層が埋めるように設計されたギャップです。フォークやラッピングではなく拡張モデルとして構成された 3 つのアーキテクチャ層です。このシリーズの残りの部分では、その構築方法を示します。
パート 2 : データ層。キャッシュを意識したルーティングを可能にするデータ構造: シャード ビットマップ、フィボナッチ スクランブル分散、および P x N スキャンを O(K x log N) に変える累積ブロック ハッシュのバイナリ検索。
パート 3: 意思決定層と実行層。キャッシュの状態をルーティングの決定に変換し、次に実行に変換します。 5 段階のコンポーザブル パイプライン、プラグイン間の型指定された状態、および同じフレームワークをラウンドロビンから分離されたプリフィル/デコードまで拡張するセレクター/ワークフロー/エグゼキューターの分割。
LLM 推論に新しい種類のルーターが必要な理由 - パート 3
LLM 推論に新しい種類の推論が必要な理由

ルーター - パート 2
モジュラーで AI の未来を構築する
今すぐクラウド プラットフォームにサインアップして、簡単に始めましょう。
モデル カタログを参照するか、独自のカスタム モデルをデプロイします
最新のニュース、お知らせ、アップデートをすべて受信箱に直接配信します。いつでも購読を解除してください。
ニュースレターにご登録いただきありがとうございます。 🚀
ヒポクラティック AI は Modular と提携して、リアルタイムの患者会話のための柔軟で高品質な推論を強化します
最新のニュースを入手し、
お知らせと最新情報:
画像生成 コード生成 AIエージェント
規約、プライバシー、利用規約
最新のニュース、お知らせ、アップデートをすべて受信箱に直接配信します。いつでも購読を解除してください。
ニュースレターにご登録いただきありがとうございます。 🚀
私はニューハンプシャー大学のクラウド コンピューティング研究室の博士課程の学生で、ネットワーキングのハードウェア アクセラレーションと分散システムの信頼性に関する研究を行っています。また、私は DPDK (データ プレーン開発キット) の元コンポーネント メンテナでもあり、博士課程で働くために辞任する前に、ネットワークの「内部」を調べ始めたのもここでした。これらすべてが意味するのは、私はコンピューターを高速に動作させるのが好きであり、Mojo + MAX は私のハードウェアへの愛、高性能ソフトウェア、そしてプログラミング言語への興味を組み合わせるのに最適な場所であるということです。

[切り捨てられた]

## Original Extract

HTTP routing has been a solved problem for many years. Round-robin, consistent hashing, least-connections. Pick one, put it in front of a pool of identical servers, and the traffic spreads pretty evenly.

Modular: Why LLM Inference Needs a New Kind of Router - Part 1
Hippocratic AI + Modular to power real-time patient conversations. Read More →
Access frontier models via an API
Generate images from text prompts
Generate production-ready code
Generate video from text + image
Proven results from real customers
GenAI native modeling & serving
The best GPU & CPU performance
Build the future of AI together
Official AI agent skills from Modular
Deploy GenAI models, our cloud or yours
Write high-performance kernels for CPUs and GPUs
Build the future of AI together
Build AI for anyone, anywhere.
Why LLM Inference Needs a New Kind of Router - Part 1
HTTP routing has been a solved problem for many years. Round-robin, consistent hashing, least-connections. Pick one, put it in front of a pool of identical servers, and the traffic spreads pretty evenly.
But then came Large Language Models.
The backends here aren't interchangeable web servers. They're GPU pods holding large, local KV caches in high-bandwidth, RAM or SSD memory. That state is expensive to rebuild, not uniformly available across the cluster, and often determines whether a request returns quickly or spends seconds recomputing previous work. Some pods might specialize in prefill, others in decode. Conversations typically stretch across requests. A single inference call sometimes needs two backends in sequence. The old assumptions about "interchangeable backends" and "independent requests" don't support these requirements.
Traditional routing is blind to all of this. It treats every backend as interchangeable, every request as independent, every pod as equally good. GPU pods are none of those things. They’re stateful, specialized and heterogeneous. Inference routing has to account for that.
This is the first post in a three-part series about what routing has to become to handle inference workloads. Modular Cloud’s orchestration layer is built around this routing problem, and this series explains how it solves it.
HTTP-era load balancing is built on a small menu of algorithms, each tuned to a specific deployment shape. They have different policies, but they share the same precondition: stateless backends.
Round-robin distributes requests uniformly across a pool of identical backends. It assumes every backend serves every request at the same cost. This might look like eight replicas of the same web service behind a load balancer, each getting 12.5% of the traffic. It’s simple, fair, stateless.
Consistent hashing routes each request to a backend determined by hashing some property of the request (a key, url, session identifier), and picking the backend whose position on the hash ring is closest. It’s the routing strategy of choice when you want the same key to land on the same backend, for client-side caching or session affinity. The backend’s “stickiness” is a function of the request key, not of anything the backend is holding in memory.
Least-requests sends each new request to whichever backend has the fewest active requests, on the assumption that fewer active requests means more spare capacity. It works when every request takes roughly the same amount of work.
These policies share the same three assumptions:
Any backend can serve any request. The assignment is a policy choice not a correctness one.
Requests are independent. What happened on request N doesn’t change what you should do on request N+1.
Backends are interchangeable. The load balancer can swap one for another without the client noticing.
Those assumptions hold for stateless web services. LLM inference breaks all three.
LLM workloads violate the stateless assumptions in four specific ways. Each one introduces a dimension that traditional routing has no mechanism to handle.
When a pod serves an inference request, the forward pass builds a KV cache: the model's intermediate state for every token position, held in GPU memory. Modern engines retain that cache after the response completes, so later requests sharing a prefix can skip the equivalent compute.
This changes the routing problem drastically. A 100K token prompt landing on a pod with the first 75K tokens already cached can prefill in milliseconds. The same prompt hitting a cold pod takes seconds. Round-robin, blind to cache state, would produce unpredictable time to first token (TTFT) for identical requests.
Cache state is the primary driver of prefill latency variance at scale. A router that selects pods based on cache residency eliminates prefill compute proportional to the shared prefix length for every hit. This frees up GPU cycles the cluster would otherwise spend recomputing work it has already done.
LLM Inference has two phases, and they stress hardware differently.
Prefill processes the entire prompt in parallel. It’s compute-bound. This means GPU cores are saturated doing dense matrix multiplications across thousands of tokens at once.
Decode generates tokens one at a time autoregressively, each token depending on every token before it. It’s memory-bandwidth-bound. This means most of the GPU’s time is spent fetching model weights and KV cache from high bandwidth memory (HBM), and most of the compute sits idle.
Running both phases on the same pod means the hardware is never tuned for either. Prefill needs dense compute; decode needs memory bandwidth. A pod optimized for one underutilizes what the other requires. Disaggregated deployments use pods tuned for each phase separately. A single client request divides work across both.
Modern engines use chunked prefill to interleave the two phases on the same pod, blurring the boundary. But the underlying compute-vs-bandwidth distinction still holds, and when you disaggregate at the deployment level, your router has to know which pod can do what.
Most LLM traffic is multi-turn. A user sends a message. The assistant replies. The user sends another message, and that message implicitly contains the entire conversation history as context.
Turn N+1 shares a prefix with turn N: the system prompt, all prior turns, all prior assistant replies. If the KV cache from turn N is still resident on some pod, turn N+1 is effectively free to prefill for the shared portion. If the cache has been evicted, or if turn N+1 lands on a different pod, the shared prefix is recomputed from scratch.
Session affinity in HTTP used to mean “route this user’s requests to the same backend so the application can use in-memory state.” In LLM inference it means the same thing but the in-memory state is the KV cache. Getting it right is the difference between sub-second responses and multi-second responses on every turn after the first.
A single client-facing request may require more than one backend.
In a disaggregated deployment, the prefill pod builds the KV cache and the decode pod generates tokens. Neither can serve the request alone. The router picks a prefill pod, then a decode pod, then orchestrates the sequence: send the prompt to prefill, wait for completion, send the same prompt plus a cache hint to decode, stream tokens back to the client.
HTTP load balancers don’t do this. They pick one backend per request. Adding multi-step coordination to a single-dispatch router is a different shape of routing entirely.
Each of the four dimensions above imposes requirements on a routing system. Those requirements fall into three distinct architectural concerns, each handled by a separate layer.
This layer tracks LLM-specific state at the latencies routing decisions require. The question “which of N pods has these blocks cached?” has to be answerable in microseconds, under concurrent updates, resilient to pod churn. A hashmap with a mutex isn’t sufficient.
This layer expresses routing logic as compositions of small, testable, reusable components. Operators pick a filter, a few scorers, a picker, and assemble a profile. The framework validates the composition at build time, not under traffic at 3am.
This layer coordinates multi-step request flows on top of the decision layer. Single-dispatch routing is a degenerate case of multi-step: one pod, one step. Disaggregated prefill/decode is the general case: two pods, two steps, with the second decision informed by the first. The same framework handles both without requiring a new HTTP handler per variant.
Parts 2 and 3 of this series build these layers.
This series describes the routing layer inside Modular Cloud's distributed inference framework, and how it handles each of these four problems in production inference workloads.
Prefix-aware routing (tokenization, block-level hashing, cache-aware scoring with load-aware tiebreaking, circuit breakers on upstream latency) ships as a profile configuration, not a new algorithm. When the team needs a new routing behavior, the work is composing plugins into a new profile rather than writing a new routing strategy from scratch. Each new deployment pattern reuses what's already there.
LLMs introduced four dimensions that traditional load balancers have no mechanism to handle: KV cache state that makes backend selection a performance-critical decision, hardware specialization that splits a single request across pod types, conversation continuity that ties sessions to cache residency, and multi-step execution that requires coordinating a sequence of backends rather than picking one.
This problem has been tackled from multiple angles. NVIDIA Dynamo, llm-d, vLLM production-stack, AIBrix, KServe, and Envoy AI Gateway have each advanced inference routing in different directions: disaggregated prefill/decode, prefix-aware scheduling, KV-aware load balancing, production-grade serving primitives. Modular Cloud builds on that foundation. To support the range of deployment patterns it targets, Modular Cloud makes composable plugins and multi-step execution both first-class primitives, so a new deployment pattern becomes a profile you assemble rather than a strategy you fork.
That’s the gap Modular Cloud’s routing layer is designed to close: three architectural layers with composition as the extension model rather than forking or wrapping. The rest of this series shows how it’s built.
Part 2 : The data layer. The data structure that makes cache-aware routing possible: sharded bitmaps, Fibonacci-scrambled distribution, and binary search over cumulative block hashes that turns a P x N scan into O(K x log N).
Part 3: The decision and execution layers. Turning cache state into routing decisions and then into execution. A five-stage composable pipeline, typed state between plugins, and the Selector/Workflow/Executor split that scales the same framework from round-robin to disaggregated prefill/decode.
Why LLM Inference Needs a New Kind of Router - Part 3
Why LLM Inference Needs a New Kind of Router - Part 2
Build the future of AI with Modular
Signup to our Cloud Platform today to get started easily.
Browse our model catalog, or deploy your own custom model
Get all our latest news, announcements and updates delivered directly to your inbox. Unsubscribe at anytime.
Thanks for signing up to our newsletter! 🚀
Hippocratic AI partners with Modular to power flexible, high-quality inference for real-time patient conversations
Get the latest news,
announcements & updates:
Image generation Code generation AI Agents
Terms , Privacy & Acceptable Use
Get all our latest news, announcements and updates delivered directly to your inbox. Unsubscribe at anytime.
Thanks for signing up to our newsletter! 🚀
I'm a PhD student at the University of New Hampshire in the Cloud Computing Lab, where I conduct research on hardware acceleration of networking and distributed systems reliability. I'm also a former component maintainer for DPDK, the Data Plane Development Kit, which is where I got my start looking "under the hood" at networking before resigning to work in my PhD. All of this means I like making computers go fast, and Mojo + MAX is a great place to combine my love of hardware, high performance software, and my interest in programming language

[truncated]
