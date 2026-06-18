---
source: "https://www.anyscale.com/blog/high-performance-distributed-inference-ray-serve-llm-vllm-google-kubernetes-gke"
hn_url: "https://news.ycombinator.com/item?id=48588650"
title: "High Performance Distributed Inference with Ray Serve LLM"
article_title: "High Performance Distributed Inference with Ray Serve LLM | Anyscale"
author: "robertnishihara"
captured_at: "2026-06-18T18:55:05Z"
capture_tool: "hn-digest"
hn_id: 48588650
score: 2
comments: 0
posted_at: "2026-06-18T17:27:02Z"
tags:
  - hacker-news
  - translated
---

# High Performance Distributed Inference with Ray Serve LLM

- HN: [48588650](https://news.ycombinator.com/item?id=48588650)
- Source: [www.anyscale.com](https://www.anyscale.com/blog/high-performance-distributed-inference-ray-serve-llm-vllm-google-kubernetes-gke)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T17:27:02Z

## Translation

タイトル: Ray Serve LLM を使用した高性能分散推論
記事のタイトル: Ray Serve LLM による高性能分散推論 |エニースケール
説明: Ray Serve LLM + vLLM スタックが、直接ストリーミング、HAProxy 統合、新しい vLLM Ray エグゼキューター バックエンドにより、最大 24 倍のスループットを実現する方法を学びます。

記事本文:
レイ サミットは 2026 年 8 月 24 ～ 26 日に再び開催されます。今ならパスが 50% 割引になります! Anyscale 製品の使用例 開発者 お客様 価格 ログイン $100 クレジットで始める ホーム ブログ ブログの詳細 Ray Serve LLM による高性能分散推論
本日、Google Cloud の Google Kubernetes Engine (GKE) チームと提携して、スタック全体のアーキテクチャの変更によって、Ray Serve LLM のスループットとレイテンシ特性における大きなマイルストーンを達成したことを発表します。オーケストレーション オーバーヘッドの削減における Ray Serve LLM の進歩を示すために、既知の高性能の Rust ベースのルーティング フレームワーク vllm-router との比較と、遡及的なパフォーマンスの比較を示します。
Ray は、異種ハードウェアを使用した複雑な分散コンピューティングのバッチ推論パイプラインによく選ばれています。さらに、Kubernetes と VM にわたる耐障害性、可観測性、柔軟性を実現する Ray の強力なプリミティブにより、LLM 推論のデプロイメントがますます複雑になるにつれて、次世代の最適化が可能になると考えています。
以下では、Ray Serve LLM + vLLM スタックに対する 3 つの主要な最適化、つまりダイレクト ストリーミング、新しい vLLM Ray エグゼキュータ バックエンド、および HAProxy 統合について説明します。その結果、プリフィルの多いワークロードでは以前のバージョンと比べてリクエストのスループットが最大 4.4 倍、デコードの多いワークロードではリクエストのスループットが最大 24 倍向上することがわかりました。
最適化の累積効果: 上の図は、vllm-router の背後にある vLLM と比較した、増分最適化の累積効果を示しています。 Ray Serve LLM は、プリフィルとデコードが多いワークロードの両方で vllm-router のパフォーマンスと同等になり、最適化作業前の Ray Serve LLM ベースラインと比較して 4.4 倍および 24.8 倍の向上を示しています。 1
3 つの主要な最適化が Ray Se に貢献します。

LLM の新しいパフォーマンス機能をご覧ください。
Link Ray Serve LLM: ダイレクト ストリーミング
Ray 2.56 では、Ray Serve LLM のダイレクト ストリーミング モードが導入されています。この新しいアーキテクチャは、リクエスト ルーティング コントロール プレーンをリクエスト/レスポンス ストリーミング データ プレーンから切り離します。
順方向パスでは、HAProxy イングレス ロード バランサは、ユーザーが構成したルーティング ポリシーに基づいて、ルーティング決定のためのリクエスト コンテンツをイングレス リクエスト ルーターに照会します。次に、HAProxy は選択したターゲット レプリカとの直接 HTTP 接続を確立し、トークンをクライアントに直接ストリーミングします。
新しい設計は、中間ルーティング デプロイメント (OpenAiIngress) が応答トークンを HAProxy に転送してイベント ループに負担をかけ、出力トークンあたりの時間 (TPOT) を増加させるというレガシー アーキテクチャのボトルネックを解決します。 RAY_SERVE_LLM_ENABLE_DIRECT_STREAMING=1 を設定してこれを試してください。使用方法についてはドキュメントを参照してください。
Ray Serve LLM ダイレクト ストリーミング: 上の図では、LLMRouter はダイレクト ストリーミング アプリケーションのイングレス リクエスト ルーターとして機能します。ルーティングの決定を行った後、HAProxy はデータ プレーン通信のためにターゲット レプリカへの接続を直接確立できます。 OpenAiIngress は、レガシー アーキテクチャで使用される中間ルーティング デプロイメントでした。
リンク vLLM: Ray Executor バックエンド V2
vLLM 用に改良された Ray バックエンドである RayExecutorV2 は、vLLM 0.21.0 でデフォルトで有効になり、プロセス管理機能と mp バックエンドのデータ プレーンおよびコントロール プレーンの実証済みの機能セットを組み合わせます。さらに、新しい Ray バックエンドにより、非同期スケジューリングなどの他の機能の継承が容易になります。
Ray 2.55 では、C ベースの HAProxy Ingress ロード バランサーと高スループット モードの最適化という 2 つの主要な最適化を Ray Serve にリリースしました。 LLM サービスの場合、これには以下も含まれます

ストリーミング パフォーマンスを向上させるために、デフォルトで TCP データグラム バッファリング (Nagle のアルゴリズム) を無効にしました。詳細については、発表のブログ投稿とドキュメントで説明されています。
Ray 2.56 では、HAProxy は、LLM サービス用の推奨コンテナ イメージである rayproject/ray-llm:2.56-py312-cu130 を含むすべての rayproject/ray コンテナ イメージで利用できます。これには、DeepGEMM などの vLLM ベース イメージの追加機能が含まれます。
Ray 2.56 では、Ray Docker イメージを使用できない場合は、 pip install ray-haproxy を介して HAProxy をインストールし、 RAY_SERVE_EXPERIMENTAL_PIP_HAPROXY=1 で有効にすることができます。このバイナリは、Ray 2.57 の pip install ray[serve] で自動的に組み込まれ、有効になります。
一般的なプリフィルとデコードが多いワークロードをシミュレートするために、入力シーケンス長 (ISL) と出力シーケンス長 (OSL) の比率が異なるワークロードと、リクエスト ルーティングとキャッシュの再利用機能を実証するためにマルチターン エージェント ワークロードを検討しました。具体的には次のとおりです。
ISL=8000、OSL=50のランダム化された事前入力の多いワークロード
ISL=50、OSL=500のランダム化されたデコード負荷の高いワークロード
20 ターンを上限とするマルチターン コーディング エージェントからのシミュレートされたプロンプトとトラフィック パターンのトレース
ランダム ワークロードは、ワークロードにはプレフィックス キャッシュの利点がないため、オーケストレーションを分離することを目的としています。たとえば、プレフィルが多いワークロードは最初のトークンまでの時間 (TTFT) を強調する傾向があり、デコードが多いワークロードは出力トークンあたりの時間 (TPOT) を強調する傾向があります。これらの実験では、コールド スタート アーティファクトを排除するために、一連のウォームアップ リクエストの後、テストされた各フレームワークの同時実行性をスイープし、TTFT、TPOT、およびスループットを測定しました。
3 番目のケースでは、Dynamo の aiperf ベンチマーク スイートを使用して合成エージェント ワークロードを生成しました。このベンチマーク スイートを使用すると、マルチターン コディンの数などのシナリオを記述することができます。

g セッション、ツールと人間の対話の待機時間の分布、およびセッションの共有または個別のコンテキスト トークンの数。特に、次の特性を持つワークロードをエミュレートしました。
セッションあたりの固定ターン数は 20 です
初期コンテキストの平均 = 25,000 トークン、中央値 = 24,000 トークン
新しいトークンの平均 = 1,000、中央値 = 400、短期および長期のツール呼び出し応答をモデル化
平均世代長 = 230、中央値 = 70
ターン間の待ち時間の中央値は 1.2 秒
セッションあたりの実効共有プレフィックス率 96%
このワークロードは、エージェントがツール呼び出しを待機しているときのターン間の待ち時間をシミュレートして、コーディング エージェントからのトラフィック パターンをシミュレートします。このワークロードを使用して、さまざまなルーティング ポリシーやフレームワークを比較できます。特に以下を比較しました。
vllm-router の一貫したハッシュ アルゴリズム
一貫性のあるハッシュを使用した Ray Serve LLM
エージェント ワークロードの場合、リクエストにセッション ID を含め、コンシステント ハッシュ アルゴリズムを使用して負荷分散を行うことができます。詳細については、一貫したハッシュに関する Ray Serve のドキュメントを参照してください。
フレームワークのオーバーヘッドを分離するために、非常に小さなモデルを使用しました。8 つのレプリカ トライアルには Qwen/Qwen3-0.6B、プレフィル/デコード分解と WideEP トライアルには Microsoft/Phi-tiny-MoE-instruct を使用しました。
8 つの Qwen3-0.6B レプリカにわたるリンク ルーティング
3 つのマルチレプリカ ワークロードすべてにわたって、Ray Serve LLM は、テストされたすべての同時実行レベルで vllm-router の総スループットと一致します。図の各行は、プレフィル ヘビー、デコード ヘビー、およびエージェント コーディングのワークロードに対応します。各列は同一のメトリクスです。平均 TTFT、平均 TPOT、および 1 秒あたりのリクエストで測定されたスループットであり、X 軸でパラメータ化されたユーザー リクエストの同時実行数 (バッチ サイズ) にわたって Ray Serve LLM と vllm-router を比較しています。
同時実行 256 のランダム ワークロードの場合、Ray Serve

LLM は、TTFT で vllm-router と同等またはそれを上回っています。prefill が多いワークロードでは 355 ミリ秒対 vllm-router の 389 ミリ秒、decode-heavy では 165 ミリ秒対 190 ミリ秒です。スループットはすべての実験で厳密に追跡されます。 KV 認識/セッション アフィニティ ルーティングを使用した現実的なエージェント マルチターン ワークロードでは、Ray Serve LLM は TPOT で vllm-router を厳密に追跡し、TTFT とリクエストのスループットでわずかに優れています。
2 つのフレームワーク間のデコード負荷の高い TTFT の相違を調査したところ、同時実行数 256 (14.7 ミリ秒の Ray Serve LLM 対 17.7 ミリ秒の vllm ルーター平均) でエンジンの観点から TTFT がほぼ一致していることがわかりました。これは、クライアント視点の縮小された TTFT Ray Serve LLM が、HAProxy 入力データプレーンの効率によって推進されていることを示唆しています。
Phi-tiny-MOE での WideEP とプレフィル/デコード分解のリンク
分離された 4P4D Wide-EP 構成 (1 つの DP4EP4 プレフィル レプリカ、1 つの DP4EP4 デコード レプリカ) では、Ray Serve LLM は、上記の 8 つのレプリカ スケーリング トライアルからの同じエージェント ワークロードを使用して、同時実行範囲全体で vllm-router 出力スループットを上回りました。同時実行数が高い場合、Ray の平均 TPOT/ITL はわずかに優れています。同時実行数 256 では、vLLM ルーターの 14.8 ミリ秒に対して 13.6 ミリ秒です。さらに、Ray Serve LLM のプリフィル/デコード分解アーキテクチャの効果は、ベースラインと比較して TTFT が減少していることで示されています。トークン化は一度実行され、再利用されるため、長いプロンプトに対するフロントエンドのオーバーヘッドが削減されます。 Ray Serve LLM のプレフィル/デコード分解と Wide-EP API の詳細については、ここを参照してください。
このマイルストーンは、Anyscale と Ray と Google Kubernetes Engine Ray チームとの継続的なエンジニアリング コラボレーションがなければ実現しなかったでしょう。Ray チームは、HAProxy および Direct Streaming アーキテクチャの支持と検証において重要な役割を果たしました。
詳細については、GKE パートナーのブログ投稿「Gemma 4 E2B」をご覧ください。

B200での結果。
Ray Serve レイヤーの HAProxy、Ray Serve LLM のダイレクト ストリーミング、vLLM の v2 Ray executor バックエンドなど、スタック全体の最適化により、以前は Ray Serve LLM をスタンドアロン vLLM から分離していたオーケストレーション オーバーヘッドが大幅に削減されました。
プレフィルが多いワークロード、デコードが多いワークロード、およびエージェントのマルチターン ワークロードにわたって、Ray Serve LLM は、Ray のフォールト トレランス、オブザーバビリティ、および異種ハードウェア プリミティブを維持しながら、総スループットで vllm-router と一致するようになりました。これらの同じプリミティブは、分離されたプリフィル/デコードおよびワイド EP トポロジにきれいに拡張され、開発者に単純な単一レプリカのケースと最も複雑な運用サービス パターンの両方に対応する単一の基板を提供します。
Ray 2.56 で試してみて、Ray Slack に参加してフィードバックを共有してください。
ベンチマークコードはこちら: https://github.com/anyscale/llm-direct-streaming-benchmarks
vLLM バージョン: 0.22.0
Ray バージョン: 2.56 ナイトリー
vllm-ルーター: 0.1.14
AIパフォーマンス：0.8.0
GPU: 8x NVIDIA H100 80GB HBM3
GPUドライバー: 580.126.20
CUDA環境バージョン：13.0.0
NCCL 環境バージョン: 2.27.7
CPU: AMD EPYC 7R13 プロセッサー
CPU トポロジ: 192 論理 CPU、2 ソケット、48 コア/ソケット、2 スレッド/コア、2 NUMA ノード
メモリ: 2.0 TiB
1 2.54 より前の Ray バージョンでは、デフォルトのストリーミング パスでの Python イベント ループの競合を軽減するバッチ メカニズムを実装しました。このバッチ処理により、オーケストレーターのオーバーヘッドが削減され、イベント ループの圧力が軽減されることでストリーミング パフォーマンスが向上しました。このグラフに示されている比較では、バッチベースの軽減策は意図的に無効になっています。以前のバージョンのバッチ処理されていないベースラインと、新しい最適化が有効になったバッチ処理されていない構成を比較し、同一の比較を確実に行います。
Ray Serve LLM: ダイレクト ストリーミング
8 つの Qwen3-0.6B レプリカにわたるルーティング
W

Phi-tiny-MOE での ideEP とプレフィル/デコード分解
製品アップデートにサインアップする 推奨コンテンツ Ray Serve へのメジャー アップグレード: レイテンシーが 88% 低下し、スループットが 11.1 倍向上したオンライン推論 続きを読む vLLM を使用した DeepSeek‑R1 のデプロイと Kubernetes 上の Ray Serve 詳細 続きを読む Ray Serve 上の AI エージェント: シングル エージェントからマルチ エージェント アーキテクチャ 続きを読む Anyscale を今すぐ探索する
実稼働 AI 用に構築されたマルチクラウド プラットフォームを使用して、Ray 上であらゆる AI ワークロードを構築、実行、拡張します。

## Original Extract

Learn how Ray Serve LLM + vLLM stack achieves up to 24x higher throughput with direct streaming, HAProxy integration, and a new vLLM Ray executor backend.

Ray Summit returns Aug 24–26, 2026. Save 50% on your pass now! Anyscale Product Use Cases Developers Customers Pricing Log in Get Started with $100 Credit Home Blog Blog Detail High Performance Distributed Inference with Ray Serve LLM
Today, in partnership with the Google Kubernetes Engine (GKE) team at Google Cloud, we are announcing a major milestone in Ray Serve LLM ’s throughput and latency characteristics, driven by architecture changes across the stack. We include comparisons to a known high-performance, rust-based routing framework, vllm-router , as well as a retrospective performance comparison, to illustrate the progress Ray Serve LLM has made in reducing orchestration overhead.
Ray is a popular choice for complex distributed computing batch inference pipelines with heterogeneous hardware. In addition, we believe that Ray’s powerful primitives for fault tolerance, observability, flexibility across Kubernetes and VMs will enable the next generation of optimizations as LLM inference deployments become increasingly complex.
Below, we cover three major optimizations to the Ray Serve LLM + vLLM stack: direct streaming, a new vLLM Ray executor backend, and HAProxy integration. As a result, we see up to 4.4x higher request throughput than previous versions on prefill-heavy workloads, and up to 24x higher request throughput on decode-heavy workloads.
Cumulative Effect of Optimizations: The figure above shows the cumulative effect of the incremental optimizations compared to vLLM behind vllm-router. Ray Serve LLM now matches vllm-router performance in both prefill- and decode-heavy workloads, representing a 4.4x and 24.8x improvement over the Ray Serve LLM baseline prior to the optimization effort. 1
Three major optimizations contribute to the Ray Serve LLM’s new performance capabilities.
Link Ray Serve LLM: Direct Streaming
Ray 2.56 introduces direct streaming mode for Ray Serve LLM. This new architecture decouples the request routing control plane from the request/response streaming data plane.
On the forward path, the HAProxy ingress load balancer queries an ingress request router with the request content for a routing decision, based on a user-configured routing policy. Next, HAProxy establishes a direct HTTP connection with the selected target replica and streams tokens directly back to the client.
The new design resolves a bottleneck in the legacy architecture where the intermediate routing deployment (OpenAiIngress) was also responsible for forwarding response tokens back to HAProxy, taxing its event loop and adding to time per output token (TPOT). Try this out by setting RAY_SERVE_LLM_ENABLE_DIRECT_STREAMING=1. See docs for usage.
Ray Serve LLM Direct Streaming: In the figure above, LLMRouter serves as the direct streaming application’s ingress request router. After serving a routing decision HAProxy can establish a connection directly to the target replica for data-plane communication. OpenAiIngress was the intermediate routing deployment used in the legacy architecture.
Link vLLM: Ray Executor Backend V2
The revamped Ray backend for vLLM, RayExecutorV2 , is enabled by default in vLLM 0.21.0 and combines the process management capabilities with the battle-tested feature set of the mp backend’s data and control planes. In addition, the new Ray backend facilitates the inheritance of other features such as asynchronous scheduling.
In Ray 2.55, we released two major optimizations to Ray Serve: a C-based, HAProxy ingress load balancer and high throughput mode optimizations. For LLM serving, this also included disabling TCP datagram buffering (Nagle’s algorithm) by default for improved streaming performance. Details are covered in the announcement blogpost and docs .
In Ray 2.56, HAProxy is available in all rayproject/ray container images, including rayproject/ray-llm:2.56-py312-cu130 , our recommended container image for LLM serving, which includes extras from the vLLM base images, such as DeepGEMM.
If the Ray docker images can’t be used, in Ray 2.56, HAProxy can be installed via pip install ray-haproxy and enabled with RAY_SERVE_EXPERIMENTAL_PIP_HAPROXY=1 . The binary will be automatically included and enabled with pip install ray[serve] in Ray 2.57.
We considered workloads with varying input sequence length (ISL) to output sequence length (OSL) ratios to simulate generic prefill- and decode-heavy workloads, and a multi-turn agentic workload to demonstrate request routing and cache reuse capabilities. In particular, these were:
Randomized prefill-heavy workload with ISL=8000, OSL=50
Randomized decode-heavy workload with ISL=50, OSL=500
Simulated prompt and traffic pattern traces from a multi-turn coding agent capped at 20 turns
The random workloads are intended to isolate orchestration due to the lack of prefix-caching benefits in the workload. For example, prefill-heavy workloads tend to highlight time to first token (TTFT), while decode-heavy workloads highlight time per output token (TPOT). For these experiments, we sweep concurrency and measure TTFT, TPOT and throughput for each of the tested frameworks after a set of warm up requests to eliminate cold start artifacts.
For the third case, we generated a synthetic agentic workload using Dynamo’s aiperf benchmark suite. With this benchmark suite, we are able to describe scenarios like number of multi-turn coding sessions, distribution of wait times for tools and human interactions and number of shared or separate context tokens for sessions. In particular, we emulated a workload with the following characteristics:
Fixed number of 20 turns per session
Mean initial context = 25,000 tokens and median = 24,000 tokens
Mean new tokens = 1,000 and median = 400, modeling short and long tool call responses
Mean generation length = 230 and median = 70
Median inter-turn latency of 1.2 seconds
Effective shared prefix rate of 96% per session
This workload simulates traffic patterns coming from a coding agent with simulated wait times between turns when the agent is waiting on tool calls. We can use this workload to compare different routing policies as well as frameworks. In particular we compared:
vllm-router’s consistent hashing algorithm
Ray Serve LLM with consistent hashing
For agentic workloads, we can include a session ID with requests and use a consistent hashing algorithm to do load-balancing. See the Ray Serve docs on consistent hashing for more.
To isolate framework overhead, we used very small models: Qwen/Qwen3-0.6B for eight replica trials and microsoft/Phi-tiny-MoE-instruct for the prefill/decode disaggregation and WideEP trials.
Link Routing across eight Qwen3-0.6B replicas
Across all three multi-replica workloads, Ray Serve LLM matches vllm-router’s aggregate throughput at every concurrency level tested. Each row in the figure corresponds to a workload: prefill-heavy, decode-heavy, and agentic coding. Each column is an identical metric: mean TTFT, mean TPOT, and throughput measured in requests per second, comparing Ray Serve LLM to vllm-router across parameterized user request concurrencies (batch size) on the x-axis.
For the concurrency 256 random workloads, Ray Serve LLM matches or beats vllm-router on TTFT: 355ms vs. vllm-router’s 389ms on prefill-heavy workloads, and 165ms vs. 190ms on decode-heavy . Throughput tracks closely for all experiments. On the realistic agentic multi-turn workload with KV-aware/session-affinity routing, Ray Serve LLM tracks vllm-router closely on TPOT, and is slightly ahead in TTFT and request throughput.
We investigated the divergence in decode-heavy TTFT between the two frameworks, and found that TTFT matched closely from the engine perspective at concurrency 256 (14.7ms Ray Serve LLM vs. 17.7ms vllm-router mean). This suggests that the reduced client-perspective TTFT Ray Serve LLM is driven by efficiency in the HAProxy ingress dataplane.
Link WideEP and Prefill/Decode Disaggregation on Phi-tiny-MOE
In the disaggregated 4P4D Wide-EP configuration (one DP4EP4 prefill replica, one DP4EP4 decode replica), Ray Serve LLM beats vllm-router output throughput across the full concurrency range using the same agentic workload from the eight replica scaling trials above. At high concurrency, Ray’s mean TPOT/ITL is slightly better: 13.6ms vs. vLLM-router’s 14.8ms at concurrency 256. Additionally, the effect of Ray Serve LLM’s prefill/decode disaggregation architecture is shown in reduced TTFT compared to the baseline; tokenization is done once and reused, reducing frontend overhead for long prompts. For more information on Ray Serve LLM’s prefill/decode disaggregation and Wide-EP APIs, see here .
This milestone would not have been possible without Anyscale and Ray’s ongoing engineering collaboration with the Google Kubernetes Engine Ray team, who were key in advocating for and validating the HAProxy and Direct Streaming architectures.
You can see more details on the GKE partner blog post: Gemma 4 E2B results on B200.
With optimizations across the stack: HAProxy at the Ray Serve layer, direct streaming in Ray Serve LLM, and the v2 Ray executor backend in vLLM, we have significantly reduced the orchestration overhead that previously separated Ray Serve LLM from standalone vLLM.
Across prefill-heavy, decode-heavy, and agentic multi-turn workloads, Ray Serve LLM now matches vllm-router on aggregate throughput while preserving Ray's fault tolerance, observability, and heterogeneous-hardware primitives. These same primitives extend cleanly to disaggregated prefill/decode and wide-EP topologies, giving developers a single substrate for both the simple single-replica case and the most complex production serving patterns.
Try it out in Ray 2.56, and join us on the Ray Slack to share feedback!
Benchmark code here: https://github.com/anyscale/llm-direct-streaming-benchmarks
vLLM version: 0.22.0
Ray version: 2.56 nightly
vllm-router: 0.1.14
AIPerf: 0.8.0
GPUs: 8x NVIDIA H100 80GB HBM3
GPU driver: 580.126.20
CUDA env version: 13.0.0
NCCL env version: 2.27.7
CPU: AMD EPYC 7R13 Processor
CPU topology: 192 logical CPUs, 2 sockets, 48 cores/socket, 2 threads/core, 2 NUMA nodes
Memory: 2.0 TiB
1 In Ray versions prior to 2.54, we implemented a batching mechanism to mitigate Python event-loop contention in the default streaming path. This batching reduced orchestrator overhead and improved streaming performance by decreasing event-loop pressure. For the comparison shown in this chart, those batching-based mitigations were intentionally disabled. We compare the unbatched baseline of the earlier version against the unbatched configuration with the new optimizations enabled, ensuring an apples-to-apples comparison.
Ray Serve LLM: Direct Streaming
Routing across eight Qwen3-0.6B replicas
WideEP and Prefill/Decode Disaggregation on Phi-tiny-MOE
Sign up for product updates Recommended content Major upgrades to Ray Serve: Online Inference with 88% lower latency and 11.1x higher throughput Read more Deploy DeepSeek‑R1 with vLLM and Ray Serve on Kubernetes Read more AI agents on Ray Serve: Single to multi-agent architecture Read more Explore Anyscale today
Build, run, and scale any AI workload on Ray with a multi-cloud platform built for production AI.
