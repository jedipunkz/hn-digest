---
source: "https://har-ki.github.io/claude-code-sre-handbook/handbook/06-air-gapped/"
hn_url: "https://news.ycombinator.com/item?id=48492579"
title: "Running Claude Code Offline on an M3 Pro with Qwen3.6"
article_title: "6. Air-Gapped Claude Code - The Claude Code SRE Handbook"
author: "har-ki"
captured_at: "2026-06-11T19:07:52Z"
capture_tool: "hn-digest"
hn_id: 48492579
score: 10
comments: 5
posted_at: "2026-06-11T16:32:04Z"
tags:
  - hacker-news
  - translated
---

# Running Claude Code Offline on an M3 Pro with Qwen3.6

- HN: [48492579](https://news.ycombinator.com/item?id=48492579)
- Source: [har-ki.github.io](https://har-ki.github.io/claude-code-sre-handbook/handbook/06-air-gapped/)
- Score: 10
- Comments: 5
- Posted: 2026-06-11T16:32:04Z

## Translation

タイトル: Qwen3.6 を使用した M3 Pro でクロード コードをオフラインで実行する
記事のタイトル: 6. エアギャップのあるクロード コード - クロード コード SRE ハンドブック
説明: 実際の Kubernetes クラスターに対して AI エージェントを実行する場合の注意事項。コードが含まれています。

記事本文:
6. エアギャップのあるクロード コード - クロード コード SRE ハンドブック
コンテンツにスキップ
Claude Code SRE ハンドブック
6. エアギャップのあるクロードコード
検索を初期化しています
クロード・コード・スレ・ハンドブック
Claude Code SRE ハンドブック
クロード・コード・スレ・ハンドブック
ハンドブック
ハンドブック
パート 1 — フロンティア
パート 1 — フロンティア
1. ハーネスの問題
3. k8s-ai-bench のクロード コード
4. オンデマンド vs 常時オン — 選択
5. 常時稼働ウォッチャーの構築
パート 2 — OSS モデル
パート 2 — OSS モデル
6. エアギャップのあるクロードコード
6. エアギャップのあるクロードコード
目次
そもそもなぜ地元なのか
ゼロから作業セッションまで
機能させるための 4 つの修正
1. 思考を無効にする - ターン全体を予算の推論に費やしました
2. 0.20 ではなく、Ollama 0.24 が必要です
3. MLX ランナーは Modelfile テンプレートを無視します
32K ウィンドウ — さらにメモリが必要になるもの
何が持続し、何が高揚するのか、そしてこれは誰のためのものなのか
7. 調査から PR まで、エアギャップ
8. k8s-ai-bench 上のエアギャップ Qwen3.6
パート 3 — コンテキスト
パート 3 — コンテキスト
9. コンテキストの問題
11. コレクションは相関関係ではない
ゼロから作業セッションまで
機能させるための 4 つの修正
1. 思考を無効にする - ターン全体を予算の推論に費やしました
2. 0.20 ではなく、Ollama 0.24 が必要です
3. MLX ランナーは Modelfile テンプレートを無視します
32K ウィンドウ — さらにメモリが何を買うのか
何が持続し、何が高揚するのか、そしてこれは誰のためのものなのか
セットアップ、それを機能させるための修正、そしてペースを決めるハードウェア
Claude Code は、ラップトップ上でローカルに実行されているモデルに接続します。調査のために Kubernetes インシデントを提供します。 10 分後、Claude Code は結果を生成する前にタイムアウトになります。このモデルは統合ツールを一切使用せず、許可されたセッション全体を思考に費やしました。
読み込まれます。まだ機能しません。
4 回の修正後、同じラップトップでインシデントが調査から解決に至りました。

ペン プル リクエスト — 根本原因を見つけ、パッチを書き、ブランチをプッシュし、gh に PR を提出 — マシンからは何も残されません。時間がかかりました。しかし、それはループを閉じました。負荷と動作の間のギャップはこれら 4 つの修正であり、それらをクリアすると、34 分のセッションと速いセッションを分けるのはアプローチではなくハードウェアです。
始める前に、ステップバイステップのセットアップ、4 つの重要な修正、およびハードウェアが速度に与える影響についての明確な説明を説明します。始めましょう。
重要な理由は 1 つあります。それは、データがファイアウォールを通過できないことです。規制された環境やエアギャップのあるクラスターでは、ローカル ハーネスは優先ではなく、必須です。良いニュースです。それは機能します。
他の人はみんな取引のために読んでいます。 Local ではプライバシーと定額料金を購入できます。料金はレイテンシーとフロンティアよりも小さいモデルに基づいて請求されます。しかし、上記の完了したループが示すように、このようなタスクでは機能を放棄する必要はありません。スピードは。その取引が適合するかどうかについては、パート 2 で 3 回の投稿にわたって答えます。これは、実行に必要なものとハードウェアが決定するものを確立します。
リグを述べます。すべての数字はそれに依存します。
ハードウェア: Apple M3 Pro、18 GPU コア、36 GiB ユニファイド メモリ、~150 GB/秒のメモリ帯域幅。
モデル: qwen3.6:35b-a3b-coding-nvfp4 — 35.1B パラメーター、トークンあたり約 3B アクティブの専門家混合、NVFP4 量子化。ディスク上に 21 GB、ロードされると最大 20 GiB が常駐します。
ランタイム: Ollama 0.24.0、MLX ランナー (llama.cpp/Metal パスではなく、Apple のシリコンネイティブ バックエンド)。
クライアント: クロード コード v2.1.84、ローカル Ollama エンドポイントを指します。
MoE では、35B モデルをローカルで実行できます。トークンごとにアクティブになるのは最大 30 億であるため、コストは 140 億の高密度モデルに似ていますが、答えは 350 億に近づきます。高密度の 35B は 36 GiB に適合しません。
ゼロから作業セッションまで ¶
最初から最後まで: Apple Silicon と kubectl が y を指していると仮定します

私たちのクラスター。
1. Ollama をインストールし、バージョンを確認します。
ollama --version # は 0.24.0 以降である必要があります — 以下の修正 #2 を参照してください
2. モデルを引っ張ります。 1 回限り、最大 21 GB。
オラマ プル qwen3.6:35b-a3b-coding-nvfp4
3. 調整された環境でサービスを提供します。実行したままにしておきます。
OLLAMA_MLX = 1 \
OLLAMA_CONTEXT_LENGTH = 32768 \
OLLAMA_FLASH_ATTENTION = 1 \
OLLAMA_MULTIUSER_CACHE = 1 \
OLLAMA_KEEP_ALIVE = 24 時間 \
OLLAMA_NO_CLOUD = 1 \
オラマサーブ
永続的にするには、launchd plist でこれらを設定します。 Ollama は端末だけでなく、サービスとして実行されます。
4. クロード コードをローカル モデルに指定し、作業ディレクトリから起動します。
ANTHROPIC_BASE_URL = http://localhost:11434 \
MAX_THINKING_TOKENS = 0 \
クロード --model qwen3.6:35b-a3b-coding-nvfp4
環境に ANTHROPIC_API_KEY がありません。このキーがないため、Claude Code は Anthropic のクラウド サービスに接続する代わりにローカル モデル エンドポイントを使用します。
5. 最初のプロンプト — メインイベントではなく、スモークテストです。ツール呼び出しを 1 回だけ強制する簡単なものを選択してください。
kubectl get pods -A を実行し、何か異常があれば教えてください。
表示される内容: 最初のツール呼び出しは数秒以内に行われ (思考が無効になっている場合)、その後、モデルがプレフィルを実行するまで約 60 秒待つことがあります (プレフィルとは、モデルが応答の生成を開始するためにプロンプ​​トやコンテキストなどの必要な入力データをすべてメモリにロードすることを意味します。この設定では約 25,000 トークンです)。事前入力後、答えが得られます。後続のセッション (「ターン」、つまりユーザーとモデル間の対話) は、プレフィックス キャッシュにプロンプ​​トの静的部分が保存されるため、再ロードする必要がなく、高速になります。このプロセス中に Ollama ログに表示される 404 エラーのバーストは正常です (修正 #4 で解決)。
6. マシンから何も出ていないことを確認します。サーバーは「Ollama クラウドを無効にする」と出力しました。

ed: true" が起動時に表示されます。ベース URL は localhost で、API キーが設定されていません。唯一のトラフィックはループバックです。これが主張全体であり、検証済みです。
Claude Code は、マシン上で完全に実行されているモデルを使用して実際のクラスターを調査するようになりました。ハードインシデントに何が対処できるか、そしてフロンティアとの比較はパート 2 で説明されます。
動作させるための 4 つの修正 ¶
これらの修正により、セットアップが機能するようになります。それらを見逃すとセットアップが停止します。クイックスタートには何も含まれていません。すべてはシリコンではなくソフトウェアであり、ハードウェアに関係なく存続します。
1. 思考を無効にする — ターン全体を予算の推論に費やしました ¶
qwen3.6 は推論モデルです。デフォルトでは、Claude Code の思考が有効になっているため、最初のターンはその予算全体を無制限の <think> チェーンに費やすことになります。約 5 ～ 8 トークン/秒で、クロード コードがタイムアウトするまで推論を続けました。結果: 10 分間の最初のターンはタイムアウトによってキルされましたが、死ぬ前にツールコールは発行されませんでした。思考がツールの使用を妨げていることを証明することはできませんが、このハードウェアでツールの使用が開始される前に、制限のない推論が時間を使い果たしました。
修正は 1 つの環境変数 MAX_THINKING_TOKENS=0 です。対照テストではそのスケールが示されています。同じ「2+2 は何ですか」というプロンプトでは、思考オンの場合は 128 個の思考トークンと 6.7 秒かかりましたが、オフの場合は 1 トークンと 0.6 秒かかりました。思考が無効になっていると、最初のツール呼び出しは決して実行されず、数秒以内に実行されます。
正直に言っておきますが、抑制は気密ではありません。その後のセッションでは、設定にもかかわらず、依然として 1 ターンに最大 20,000 文字の推論が出力されました。設定がそのセッションに適用されなかったか、モデルが推論を通常の出力にルーティングしたかのいずれかです。それには予算をつけましょう。
2. 0.20 ではなく Ollama 0.24 が必要です ¶
Ollama 0.20 のこのモデルでは、設定を Modelfile にベイクしようとしても機能しません。 MLX ベースからの ollamate は成功したように見えますが、ollam show と API は両方とも「モデルが見つかりません」を返します。 GGUF 派生製品

仕事をする必要があります。 MLX 派生関数は失敗します。マニフェストは書き込みますが、サーバーはそれを解決できません。
この問題は 0.24.0 で修正され、MLX セーフテンソル モデルの作成がクリーンアップされ、OpenAI 互換 API を介して think パラメーターが正しくルーティングされ、MLX サンプラーが作り直されました。 0.20 を使用している場合は、何よりも先にアップグレードしてください。アップグレードするまで、チューニング レバーの半分は存在しません。
3. MLX ランナーは Modelfile テンプレートを無視します ¶
0.24 でも、Modelfile を介して思考を無効にすることは通知なしで失敗します。 <think> タグを削除したカスタム テンプレートは依然として 355 個の思考トークンを返しました。 MLX ランナーは独自のレンダラーとパーサー ( qwen3.5 ) を使用し、テンプレートをオーバーライドします。思考のコントロールはテンプレートではなくレンダラーにあります。
この教訓は考えるよりも重要です。Modelfile ノブと llama.cpp の習慣は MLX パスに転送されません。テンプレートではなく、think:false API パラメーターまたは MAX_THINKING_TOKENS=0 を使用して思考を制御します。 llama.cpp 時代のチューニング トリックは信頼する前に MLX で再検証する必要があると仮定します。
実行すると、リクエストごとに /v1/messages および /v1/messages/count_tokens への 18 回以上の失敗した呼び出しのバーストがログに表示されます。クロード・コードは、オラマが処理できない人類固有のエンドポイントを調査します。これらの 404 は高速であり、何も変更しません。無視してください。
これらの修正により、機能するようになりました。今重要なのはハードウェアです。
セッションは事前入力バインドされています。トークンを生成する前に、毎ターンプロンプト (Claude Code のシステム プロンプト、ツール定義、 CLAUDE.md 、およびこれまでの会話) を再読み取りします。このハードウェアでは、25,000 トークンのターンに 60 ～ 70 秒かかり、プリフィルにリクエスト時間の 90% 以上が費やされます。形としては、調査から PR までの 34 分間のセッションを完了するには、プレフィルに約 20 分、生成に 8 分、ツールの実行に 6 分を費やしました。遅いですが、正しく終了しました。
ここは実際にラップトップに関する部分です。プレフィル

l レートは 20 GiB モデルに対するメモリ帯域幅によって設定され、M3 Pro の ~150 GB/s が下限です。高帯域幅の Apple Silicon (Max 層と Ultra 層) では、ボトルネックが帯域幅であるため、その上限が直接引き上げられます。新しいシリコンが残りの半分を助けます。M5 クラスのハードウェアでの Ollama 独自の MLX ベンチマークでは、専用の行列乗算ユニットのおかげで、プリフィルとデコードのスループットが以前のチップに比べて大幅に上昇していることが示されています。そのどれもが異なるアプローチではありません。より高速なエンジンでも同じスタックです。
非ハードウェア レバーのみ: 小型モデルの方がはるかに早くプレフィルを完了します。それが次の投稿の焦点です。
32K ウィンドウ — そしてさらにメモリが必要になるもの ¶
コンテキスト ウィンドウは 32,000 トークンです。これは好みではなく、記憶に適合するものによって選択されます。 Ollama は、検出された GPU メモリに基づいてデフォルトのコンテキスト ウィンドウを設定します。この場合、24 ～ 48 GiB のメモリを搭載したマシンは通常、32,000 トークン ウィンドウを取得します。各トークンは、モデルがコンテキスト内で記憶できるテキストの一部を表します。計算: 20 GiB のモデル ファイルでは、KV キャッシュ (モデルが頻繁に使用するデータへの高速アクセスのために使用されるキーと値のメモリ キャッシュ) に使用できるのは数 GiB だけです。 32,000 トークンのコンテキスト ウィンドウが残りのキャッシュを埋めます。 64,000 トークンのウィンドウは利用可能なメモリを超え、システムがスワップして速度が低下します。
これは、ハードウェアがラインを動かす最も明らかな場所です。ただし、Apple Silicon には落とし穴があるため、階層を注意深く読んでください。 macOS はユニファイド メモリの一部のみを GPU に公開します。この 36 GiB マシンでは、Metal は 28.1 GiB (約 78%) を記録しました。 Ollama の最大の層はデフォルトで 256K コンテキストで、検出された GPU メモリの 48 GiB でトリガーされます。つまり、Metal は全体の一部しか認識しないため、48 GiB の Mac がそれに到達することはありません。このモデルの実際の層は次のとおりです。
48 GiB Mac はこれほど快適に動作します。同様に

さあ、記憶に負担がかかりません。 256K コンテキストには、最大 64 GiB のユニファイド メモリが必要です。
実際の SRE 作業ではウィンドウ サイズが小さいことが考慮されないため、ウィンドウ サイズが重要になります。ファイルを調査して編集するエージェント セッションにより、コンテキストが迅速に構築されます。このラップトップでのそのようなセッションの 1 つでは、会話が最大 56,000 トークンに達し (32,000 ウィンドウを超え)、ランタイムが必要なこと (削除と再処理) を実行しました。症状は明確でした。
KV キャッシュ スラッシング - 一度に 600 MiB が削除され、次のターンに再処理されます。
キャッシュ ヒット率は最大 30% に低下します。残りの約 70% は最初から再読み込みされます。
プレフィルは 1 ターンあたり 2 ～ 3 分ストレッチします。
ピーク時のメモリは 33.5 GiB (マシンの 93%) で、OS のメモリ負荷が発生し、MLX を含むすべての動作が遅くなります。
36 GiB では、ルールは次のとおりです。セッションのスコープを維持し、32K 内に留まり、クリーンに実行されます。そこを大の字に通り過ぎると、すぐに衝撃を感じます。 256K ウィンドウのヘッドルームを持つマシンでは、このルールはほとんど解消されます。ここでオーバーフローしたのと同じセッションがそこに常駐します。
何が持続し、何が引き上げられ、これは誰のためのものなのか ¶
より良いハードウェアが役立つかどうかを正直に並べ替えると、次のようになります。
プレフィル速度 - 帯域幅制限あり。 Max/Ultra/新しいシリコンはそれを高めます。
コンテキスト ウィンドウ - 36 ～ 48 GiB の 32K は、64 GiB+ では 256K になります。
メモリプレッシャー — 93% のピークはヘッドルームとともに消えます。 48 GiB 減りました。
haに関係なく持続します

[切り捨てられた]

## Original Extract

Notes on running AI agents against real Kubernetes clusters. Code included.

6. Air-Gapped Claude Code - The Claude Code SRE Handbook
Skip to content
The Claude Code SRE Handbook
6. Air-Gapped Claude Code
Initializing search
claude-code-sre-handbook
The Claude Code SRE Handbook
claude-code-sre-handbook
Handbook
Handbook
Part 1 — Frontier
Part 1 — Frontier
1. The Harness Problem
3. Claude Code on k8s-ai-bench
4. On-Demand vs Always-On — Choosing
5. Building the Always-On Watcher
Part 2 — OSS Models
Part 2 — OSS Models
6. Air-Gapped Claude Code
6. Air-Gapped Claude Code
Table of contents
Why local at all
From zero to a working session
The four fixes that make it work
1. Disable thinking — it spent the whole turn budget reasoning
2. You need Ollama 0.24, not 0.20
3. The MLX runner ignores your Modelfile template
The 32K window — and what more memory buys
What persists, what lifts, and who this is for
7. From Investigation to PR, Air-Gapped
8. Air-Gapped Qwen3.6 on k8s-ai-bench
Part 3 — Context
Part 3 — Context
9. The Context Problem
11. Collection Isn't Correlation
From zero to a working session
The four fixes that make it work
1. Disable thinking — it spent the whole turn budget reasoning
2. You need Ollama 0.24, not 0.20
3. The MLX runner ignores your Modelfile template
The 32K window — and what more memory buys
What persists, what lifts, and who this is for
The setup, the fixes that make it work, and the hardware that sets the pace
Claude Code connects to a model running locally on the laptop. You provide a Kubernetes incident for investigation. After ten minutes, Claude Code times out before producing any results. The model didn't use any integrated tools—it spent the entire allowed session thinking.
It loads. It doesn't work yet.
Four fixes later, the same laptop took an incident from investigation to an open pull request — found the root cause, wrote the patch, pushed a branch, filed the PR with gh — with nothing leaving the machine. It took its time. But it closed the loop. The gap between loads and works is those four fixes, and once you clear them, the thing that separates a 34-minute session from a fast one is hardware, not approach.
Before we begin, here's what's ahead: a step-by-step setup, the four crucial fixes, and a clear explanation of how your hardware affects speed. Let's get started.
One reason matters: data can't cross the firewall. In regulated environments and air-gapped clusters, a local harness isn't a preference—it's required. The good news: it works.
Everyone else is reading for the trade. Local buys you privacy and a flat cost. It bills you at latency and a model smaller than the frontier — but, as the completed loop above shows, capability is not what you give up on a task like this. Speed is. Whether that trade fits is what Part 2 answers across three posts. This one establishes what it takes to run, and what your hardware decides.
State the rig. All numbers depend on it.
Hardware: Apple M3 Pro, 18 GPU cores, 36 GiB unified memory, ~150 GB/s memory bandwidth.
Model: qwen3.6:35b-a3b-coding-nvfp4 — 35.1B parameters, mixture-of-experts with ~3B active per token, NVFP4 quantization. 21 GB on disk, ~20 GiB resident once loaded.
Runtime: Ollama 0.24.0, MLX runner (Apple's Silicon-native backend, not the llama.cpp/Metal path).
Client: Claude Code v2.1.84, pointed at the local Ollama endpoint.
MoE lets a 35B model run locally. Only ~3B active per token, so costs resemble a 14B dense model, while answers approach 35B. A dense 35B doesn't fit 36 GiB.
From zero to a working session ¶
Start to finish: assumes Apple Silicon and kubectl pointed to your cluster.
1. Install Ollama, confirm the version.
ollama --version # must be 0.24.0 or newer — see fix #2 below
2. Pull the model. One-time, ~21 GB.
ollama pull qwen3.6:35b-a3b-coding-nvfp4
3. Serve with the tuned environment. Leave it running.
OLLAMA_MLX = 1 \
OLLAMA_CONTEXT_LENGTH = 32768 \
OLLAMA_FLASH_ATTENTION = 1 \
OLLAMA_MULTIUSER_CACHE = 1 \
OLLAMA_KEEP_ALIVE = 24h \
OLLAMA_NO_CLOUD = 1 \
ollama serve
For permanence, configure these in a launchd plist. Ollama runs as a service, not just your terminal.
4. Point Claude Code at the local model and launch from your working directory:
ANTHROPIC_BASE_URL = http://localhost:11434 \
MAX_THINKING_TOKENS = 0 \
claude --model qwen3.6:35b-a3b-coding-nvfp4
No ANTHROPIC_API_KEY in the environment — not having this key is what makes Claude Code use the local model endpoint instead of contacting Anthropic's cloud service.
5. First prompt — a smoke test, not the main event. Pick something trivial that forces exactly one tool call:
Run kubectl get pods -A and tell me if anything appears unhealthy.
What you'll see: the first tool call happens in a few seconds (when thinking is disabled), then you may wait about 60 seconds as the model performs prefill (prefill means loading all necessary input data such as the prompt and context into memory for the model to start generating responses, which for this setup is about 25,000 tokens). After prefill, you get the answer. Subsequent sessions ('turns,' or interactions between the user and the model) are faster because the prefix cache stores the static parts of the prompt so they don't need to be reloaded. The burst of 404 errors shown in the Ollama log during this process is normal (addressed in fix #4).
6. Confirm nothing left the machine. The server printed "Ollama cloud disabled: true" on boot; the base URL is localhost, and there's no API key set. The only traffic is loopback. That's the whole claim, verified.
Claude Code now investigates a real cluster using a model running entirely on your machine. What it can tackle on hard incidents—and how it compares to the frontier—comes in Part 2.
The four fixes that make it work ¶
These fixes make the setup work. Miss them and the setup stalls. None are in the quickstart. All are software, not silicon—they persist regardless of hardware.
1. Disable thinking — it spent the whole turn budget reasoning ¶
qwen3.6 is a reasoning model. By default, Claude Code's thinking enabled meant the first turn spent its entire budget on an unbounded <think> chain. At ~5–8 tokens/sec, it kept reasoning until Claude Code's timeout. Result: a 10-minute first turn killed by timeout, no tool call emitted before dying. While I can't prove thinking blocks tool use, unbounded reasoning ran out the clock before tool use began on this hardware.
The fix is one env var: MAX_THINKING_TOKENS=0 . A control test shows the scale — the same "what is 2+2" prompt took 128 thinking tokens and 6.7s with thinking on, versus 1 token and 0.6s with it off. With thinking disabled, the first tool call lands in seconds instead of never.
One caveat, kept honest: the suppression isn't airtight. A later session still emitted ~20K characters of reasoning on one turn despite the setting — either it wasn't applied to that session, or the model routed reasoning into normal output. Budget for it.
2. You need Ollama 0.24, not 0.20 ¶
Trying to bake settings into a Modelfile doesn't work for this model in Ollama 0.20. ollama create from an MLX base seems successful, but ollama show and the API both return "model not found." GGUF derivatives work; MLX derivatives fail. The manifest writes, but the server can't resolve it.
This is fixed in 0.24.0, which cleans up MLX safetensor model creation, routes the think parameter correctly through the OpenAI-compatible API, and reworks the MLX sampler. If you're on 0.20, upgrade before anything else — half the tuning levers don't exist until you do.
3. The MLX runner ignores your Modelfile template ¶
Even on 0.24, disabling thinking via Modelfile fails silently. A custom template that strips <think> tags still returned 355 thinking tokens. The MLX runner uses its own renderer and parser ( qwen3.5 ), overriding your template; thinking control is in the renderer, not the template.
The lesson is bigger than thinking: Modelfile knobs and llama.cpp habits do not transfer to the MLX path. Control thinking with the think:false API parameter or MAX_THINKING_TOKENS=0 — not a template. Assume any llama.cpp-era tuning trick needs re-verifying on MLX before you trust it.
Once running, logs show bursts of 18+ failed calls to /v1/messages and /v1/messages/count_tokens per request. Claude Code probes Anthropic-native endpoints Ollama doesn't handle. These 404s are fast and change nothing. Ignore them.
With these fixes, it works. What matters now is your hardware.
The session is prefill-bound. Every turn re-reads the prompt — Claude Code's system prompt, tool definitions, CLAUDE.md , and the conversation so far — before generating a token. On this hardware, that's 60–70 seconds for a 25K-token turn, and prefill eats 90%+ of request time. As a shape: that completed 34-minute investigation-to-PR session spent roughly 20 minutes in prefill, 8 in generation, and 6 in tool execution. Slow — but it finished, correctly.
Here's the part that's really about the laptop. Prefill rate is set by memory bandwidth against a 20 GiB model, and the M3 Pro's ~150 GB/s is the floor. Higher-bandwidth Apple Silicon — the Max and Ultra tiers — raises that ceiling directly, because the bottleneck is bandwidth. Newer silicon helps the other half: Ollama's own MLX benchmarks on M5-class hardware show prefill and decode throughput climbing sharply over earlier chips, thanks to dedicated matrix-multiply units. None of that is a different approach. It's the same stack on a faster engine.
Only non-hardware lever: a smaller model completes prefill much faster. That's the next post's focus.
The 32K window — and what more memory buys ¶
The context window is 32,000 tokens. This is not chosen by preference, but by what fits in memory. Ollama sets the default context window based on detected GPU memory. In this case, a machine with 24–48 GiB of memory typically gets a 32,000-token window. Each token represents a piece of text the model can remember in context. The calculation: the 20 GiB model file leaves only a few GiB available for the KV cache (the key-value memory cache used by the model for fast access to data it uses often). A 32,000-token context window fills the remaining cache; a 64,000-token window would exceed available memory, making the system swap and slow down.
This is the clearest place hardware moves the line — but read the tiers carefully, because Apple Silicon has a catch. macOS exposes only part of unified memory to the GPU: on this 36 GiB machine, Metal saw 28.1 GiB, about 78%. Ollama's largest tier, which defaults to 256K context, triggers at 48 GiB of detected GPU memory — so a 48 GiB Mac never reaches it, because Metal only ever sees a fraction of the total. The practical tiers for this model:
A 48 GiB Mac runs this well—same window, no memory strain. 256K context needs ~64 GiB unified memory.
Window size matters because real SRE work doesn't respect a small one. An agentic session that investigates and edits files quickly builds context. One such session on this laptop pushed the conversation to ~56K tokens — past the 32K window — and the runtime did what it must: evict and reprocess. The symptoms were unambiguous:
KV cache thrashing — 600 MiB evicted at a time, reprocessed on the next turn.
Cache hit rate collapsing to ~30%; the other ~70% re-read from scratch.
Prefill stretching to 2–3 minutes per turn.
Peak memory at 33.5 GiB — 93% of the machine — triggering OS memory pressure that slows everything, MLX included.
On 36 GiB, the rule is: keep sessions scoped, stay inside 32K, and run clean; sprawl past it, and you feel the thrash immediately. On a machine with the headroom for a 256K window, that rule mostly dissolves — the same session that overflowed here would stay resident there.
What persists, what lifts, and who this is for ¶
Sorted honestly by whether better hardware helps:
Prefill speed — bandwidth-bound; Max/Ultra/newer silicon raises it.
Context window — 32K on 36–48 GiB becomes 256K at 64 GiB+.
Memory pressure — the 93% peaks vanish with headroom; gone by 48 GiB.
Persists regardless of ha

[truncated]
