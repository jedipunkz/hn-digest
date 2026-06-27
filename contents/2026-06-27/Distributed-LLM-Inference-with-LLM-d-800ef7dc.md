---
source: "https://cefboud.com/posts/llm-d/"
hn_url: "https://news.ycombinator.com/item?id=48699083"
title: "Distributed LLM Inference with LLM-d"
article_title: "Distributed LLM Inference with llm-d | Moncef Abboud"
author: "cefboud"
captured_at: "2026-06-27T16:23:42Z"
capture_tool: "hn-digest"
hn_id: 48699083
score: 3
comments: 0
posted_at: "2026-06-27T15:27:14Z"
tags:
  - hacker-news
  - translated
---

# Distributed LLM Inference with LLM-d

- HN: [48699083](https://news.ycombinator.com/item?id=48699083)
- Source: [cefboud.com](https://cefboud.com/posts/llm-d/)
- Score: 3
- Comments: 0
- Posted: 2026-06-27T15:27:14Z

## Translation

タイトル: LLM-d による分散 LLM 推論
記事のタイトル: llm-d による分散 LLM 推論 |モンセフ・アブード
説明: KV キャッシュの局所性と GPU 使用率を使用して、vLLM などの推論エンジン全体でリクエストをインテリジェントにスケジュールするオープンソースの LLM 対応ルーターである llm-d の紹介です。

記事本文:
llm-d による分散 LLM 推論 |モンセフ・アブード モンセフ・アブード ここではいろいろなことについて書いています。
ホーム ニュースレター タグ アーカイブ カテゴリーについて ライト ダーク システム ホーム llm-d を使用した分散 LLM 推論 ポストキャンセル llm-d を使用した分散 LLM 推論
KV キャッシュの局所性と GPU 使用率を使用して、vLLM などの推論エンジン全体でリクエストをインテリジェントにスケジュールするオープンソースの LLM 対応ルーターである llm-d の紹介です。
実稼働グレードの LLM 推論はどのようなものですか?
真夜中に目が覚めてその疑問について考えているなら、このブログ投稿が役に立つかもしれません。
vLLM や SGLang などの推論エンジンは PyTorch の上に位置します (今後の vLLM という場合、これには SGLang やその他のサポートされている推論エンジンも含まれます)。これらはノード レベルで推論を最適化し、特に KV キャッシュを管理し、ページ アテンションや連続バッチ処理などの技術を通じてスループットを向上させます。
llm-d の売り文句は、LLM 対応のロード バランサーであるということです。
複数の vLLM インスタンスがある場合、それぞれに独自の状態 (利用可能な GPU メモリ、KV キャッシュ プレフィックスの一致、処理のためにキューに入れられているリクエストの数など) があります。単純にラウンド ロビンを使用することはできません。これらの信号に基づいて最適な推論エンジン インスタンスを選択することが、llm-d のすべてです。
また、優先度に基づいてさまざまなクラスのリクエストをサポートするフロー制御などの機能も提供します (例: プレミアム リアルタイム トラフィックとバッチ ワークロード)。さらに、プレフィックスとデコードは異なる構成の恩恵を受けるため、異なるノードで実行されるスムーズな分離 P/D が可能になります。プレフィックスはコンピューティングに依存し、デコードはメモリ帯域幅に依存するため、GPU 要件が異なります。コンピューティングを最大化するためにプレフィックスの低い TP と、メモリ帯域幅を最大化するためにデコードの高い TP です。
また、

優先度に基づいてさまざまなクラスのリクエストをサポートするフロー制御などの機能 (プレミアム リアルタイム トラフィックとバッチ ワークロードなど)。さらに、プレフィックスとデコードは異なる構成の恩恵を受けるため、異なるノードで実行されるスムーズな分離 P/D が可能になります。プレフィックスはコンピューティングに依存し、デコードはメモリ帯域幅に依存するため、GPU 要件が異なります。コンピューティングを最大化するためにプレフィックスの低い TP と、メモリ帯域幅を最大化するためにデコードの高い TP です。
llm-d の優れた点は、車輪の再発明を行わないことです。既存の確立されたプロジェクトの上に構築されます。 LLM 推論は引き続き vLLM と SGLang で行われ、単に HTTP 経由で通信します。プロキシ レイヤーとディスカバリーはすべて Kubernetes(k8s) と Envoy 上に構築されています。どの vLLM インスタンスを選択するかを決定するための統合ポイントも、既存の拡張ポイント、つまり Envoy の ext_proc 拡張機能に依存しています。データ層とメトリクスの収集は、基本的に Prometheus 上に構築されます。したがって、車輪の再発明は必要なく、既存の堅固なソリューションをインテリジェントに組み合わせるだけです。最も優れた点は、llm-d が各レイヤーと部分 (メトリクス、スコアリング、フロー制御など) に対して、新しいプラグインを実装してすぐに使用できる明確なインターフェイスで簡単に拡張できることです。
そのため、vLLM が新しいメトリクスを追加した場合、または新しい推論エンジンが登場した場合でも、簡単に追加できます。ピックまたはスコアリングの新しい方法を実装したい場合は、新しいピッカーまたはスコアラーを実装するだけです。 Prometheus が人気がなくなった場合、またはカスタム監視ソリューションを使用したい場合は、代わりにそのためのプラグインを実装できます。
Gateway API Inference Extension (GAIE) の形で、k8s 上に Generative AI 推論を統合する標準化の取り組みが行われています。 GAIE は API リソース (Infere など) を定義します。

ncePool ) とエンドポイント ピッカー ロール。 llm-d のルーターは、実際のリクエスト転送を行う Envoy プロキシと組み合わされた EPP ロールの実装です。
それは非常に戦略的な選択です。 llm-d は、オーダーメイドの API を使用した独立したイニシアチブではなく、より広範な k8s 標準の実装として構築されています。
llm-d には、ファイル検出プラグインを介して k8s の外部で実行されるモードがあり、vLLM エンドポイントは k8s API を介して検出されるのではなくハードコーディングされていることにも言及する価値があります。
では、LLM 推論ルーティングで考慮する必要がある要素は何でしょうか?
N1 と N2 があるとします。ユーザーがすでに N1 から応答を取得している場合、これはユーザーのリクエスト KV キャッシュがそこですでに計算されていることを意味するため、後続のリクエストではそのステップをスキップできます。ただし、ユーザーのリクエストを N2 に送信すると、事前入力を再実行する必要があり、N1 ですでに行われた計算を利用できなくなります。
言い換えれば、プロンプトの KV キャッシュをすでに持っているノードにリクエストをルーティングするのが効率的です (たとえば、キャッシュを共有ネットワーク ストレージに保存できる KV キャッシュのオフロードがある場合、この考慮事項は変わります)。
私たちが気にするもう 1 つの要素は、N1 と N2 の空き VRAM の量です。 N1 の VRAM がほぼいっぱいの場合、たとえユーザーの KV キャッシュがあるとしても、他のリクエストが完了するまで待つか、リクエストを実行するために他のリクエストを削除する必要がある可能性があるため、これは最良の選択ではない可能性があります。同時に N2 の VRAM が空き、すぐに使用できる場合は、プリフィルを再実行して空きメモリを利用するのが最善の方法かもしれません。
どちらのノードも空き VRAM がほとんどない状態で実行されている可能性がありますが、N1 のキューには処理を待機しているリクエストが 2 つしかないのに対し、N2 には 5 つのリクエストがあります。待機中のリクエストの数が少ないノードにルーティングするのが最善です。
Th

どのノードを選択するかを決定する際に考慮すべき要素は数多くありますが、llm-d の仕事は選択することです。 /metrics Prometheus エンドポイントを介してさまざまな vLLM インスタンスからデータを収集し、各ノードの状態をメモリに保持します。ユーザーリクエストが届くと、それはエンドポイントピッカーまたはEPPと呼ばれるllm-dのルーターに送信され、EPPはスコアラーのスコア(はい、愚かなダジャレが意図されています): prefix-cache-scorer 、 kv-cache-utilization-scorer 、 queue-scorer などを可能な候補(vLLMインスタンス)ごとに実行し、それらのスコアを組み合わせて候補の1つを選択します。スコアリングと選択の前に、候補エンドポイントをフィルタリングできるフィルタリング フェーズがあります。
Envoy と外部処理フィルター
llm-d のルーターは Envoy を置き換えるのではなく、Envoy プロキシ フィルター ターゲットとして並行して動作します。リクエストが受信されると、Envoy プロキシは llm-d EPP に問い合わせます。 EPP はその役割を果たし、HTTP 応答でヘッダーを提供します。
1
x-ゲートウェイ-宛先-エンドポイント: IP_PORT_OF_A_VLLM_POD
次に、Envoy はリクエストをその vLLM ポッドにルーティングします。本当にそれだけです。
Envoy プロキシには、 ext_proc (外部処理) と呼ばれる拡張機能 (HTTP フィルター) があります。リクエストが到着すると、gRPC 経由で外部プロセスに転送され、そのプロセスは HTTP 本文の変更、ヘッダーの追加または削除、リクエストの拒否など、さまざまな処理を行うことができます。この場合、llm-d ルーター EPP は外部プロセスであり、ヘッダー (選択された vLLM インスタンスのアドレスを持つヘッダー) を追加するだけです。
その選択をどのように行うかは、llm-d がもたらす魔法です。そして、上で述べたように、推論エンジンの Prometheus エンドポイントからデータを収集します。
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
vardefaultEngineConfigs = [] EngineConfigParams {
{
名前: "vllm" 、
QueuedRequestsSpec : "vllm:num_r

equests_waiting" ,
RunningRequestsSpec : "vllm:num_requests_running" ,
KVUsageSpec : "vllm:kv_cache_usage_perc" 、
LoRASpec : "vllm:lora_requests_info" 、
CacheInfoSpec : "vllm:cache_config_info" 、
}、
{
名前 : "スグラン" 、
QueuedRequestsSpec : "sglang:num_queue_reqs" 、
RunningRequestsSpec : "sglang:num_running_reqs" ,
KVUsageSpec : "sglang:token_usage" 、
LoRAスペック: "" 、
CacheInfoSpec : "sglang:cache_config_info" 、
CacheBlockSizeLabelName : "page_size" 、
CacheNumBlocksLabelName : "num_pages" 、
}、
//...
}
And based on this data, it calculates a score for each candidate endpoint and picks one.通常、これは最も高いスコアを持つものですが、ピッカーは最大値を取得するだけでなく、スコアに応じてサンプリングすることもできます。これは素晴らしい点です。常に単一の最良のポッドにルーティングすると、嵐のような群れが作成され、良好でなくなるまで全員が同じ「良好な」レプリカを積み上げることになります。スコアに基づいて負荷を分散すると、より良い候補が優先されることを回避しながら、LLM のソフトマックス層の後にトークンをサンプリングすることとよく似ています。
Envoy establishes a gRPC connection with the EPP and streams it the request headers and body.本質的には次のとおりです。
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
func ( s * StreamingServer ) プロセス (
srv extProcPb 。外部プロセッサ_プロセスサーバー 、
) エラー {
// リクエストの存続期間中の共有状態
reqCtx := NewRequestContext ()
// ストリームリーダー
recvCh := startReceiver ( srv )
{の場合
{を選択してください
case msg := <-recvCh :
スイッチ v := msg 。リクエスト。 ( タイプ ) {
ケース * リクエストヘッダー :
handleRequestHeaders ( reqCtx , v )
ケース * リクエストボディ:
handleRequestBody ( reqCtx , v ) // Actually: s.director.HandleRequest(ctx, reqCtx, parseResult.Body)
ケース * レスポンスヘッダー:
ハンドルレスポ

nseHeaders ( reqCtx 、 v )
ケース * レスポンスボディ:
handleResponseBody ( reqCtx , v )
}
}
}
}
リクエスト本文の受信が完了したら、 s.director.HandleRequest を実行します。ディレクターは、エンドポイントとそのメトリックを含むデータストアを保持します。
いくつかのアドミッション プラグインを通じてリクエストを実行します (そのため、必要に応じてリクエストを拒否できます。これは、ほとんどのパイプラインと同様に拡張可能であり、それを処理するためにさまざまなポイントにカスタム ルールを追加できます)。
次に、ヘッダー ( x-gateway-inference-fairness-id 、 x-gateway-inference-objective ) で指定された優先度/フロー ID に基づいてリクエストが処理されるフロー制御を実行します。フローには異なる優先順位を付けることができ、同じ優先順位範囲内で公平性と順序付けポリシー (ラウンドロビン、先着順、期限が早い順など) に従うことができます。
リクエストがフロー制御を通過すると、エンドポイント候補が取得され、そのリクエストに対して「データ プロデューサー」プラグインが実行されます。これらのプラグインは、新しいデータでリクエストを強化します。トークン ID を追加することで文字列をトークン化し (vLLM API に依存するトークナイザーがあります)、もう 1 つの主要なプロデューサーはプレフィックス マッチャーです。このリクエストのトークンのうち、各候補ポッドの KV キャッシュにすでに存在するトークンがいくつあるかなどです。
次に、重要な部分であるスケジューラが登場します。フィルター、スコアラー、ピッカーを備えたスケジュール プロファイルがあります。候補をフィルタリングし、各スコアラー (プレフィックス キャッシュ スコアラー、KV キャッシュの使用率、キューの深さなど) に従ってリクエストをスコアリングし、組み合わせたスコアに基づいて選択します (各スコアラーは独自の重みを持つことができます)。
1
2
3
4
5
6
7
8
9
10
スコア:
プレフィックス キャッシュ スコアラー (重み = 3.0): ポッド 1 = 0.75、ポッド 2 = 0.4、その他 = 0.0
KV キャッシュ使用率 (重み = 1.0): ポッド 1 = 0.6、ポッド 2 = 0.8、その他は変化します
キューの深さ (重み = 1.0):

ポッド 1 = 0.7、ポッド 2 = 0.5、その他は異なります
最終スコア:
ポッド 1: (0.75 × 3) + (0.6 × 1) + (0.7 × 1) = 3.55
ポッド 2: (0.4 × 3) + (0.8 × 1) + (0.5 × 1) = 2.5
その他: < 2.0
ピッカー (重み付きランダム): ポッド 1 が強く優先されますが、決定的ではありません
結果: ポッド 1 が選択されました (IP: 10.0.1.42、ポート: 8000)。
EPP は、選択されたポッドの IP とポートに設定された x-gateway-destination-endpoint を使用して Envoy に応答を送信し、Envoy はそこに要求を転送します。ポッドがトークンを生成すると、Envoy は応答をクライアントにストリーミングして返します。 ext_proc ストリームは、メトリクスの記録とデータ更新 (新しい KV キャッシュの計算など) のための応答ヘッダーと本体チャンクも受け取ります。
モデル サーバー (vLLM) にプロンプ​​トの KV キャッシュの一部がすでにある場合、そこにリクエストをルーティングすると、プレフィックスの再計算が回避されます。各ノードの KV キャッシュの内容を知る必要があります。
過去のルーティング決定から各ノードの KV キャッシュを推定するおおよそのプレフィックス キャッシュ プロデューサーがあります。プロンプトが以前に vLLM 1 に送信されていた場合、その KV キャッシュはまだ存在している可能性があります。
欠点は、これは単なる近似値であり、キャッシュの削除を考慮していないことです。
より興味深いアプローチを採用する正確な KV キャッシュ プロデューサーがあります。 vLLM は、ZMQ (ブローカーレス

[切り捨てられた]

## Original Extract

An introduction to llm-d, an open-source LLM-aware router that intelligently schedules requests across inference engines like vLLM using KV cache locality and GPU utilization.

Distributed LLM Inference with llm-d | Moncef Abboud Moncef Abboud I write about stuff here.
HOME NEWSLETTER TAGS ARCHIVES ABOUT CATEGORIES Light Dark System Home Distributed LLM Inference with llm-d Post Cancel Distributed LLM Inference with llm-d
An introduction to llm-d, an open-source LLM-aware router that intelligently schedules requests across inference engines like vLLM using KV cache locality and GPU utilization.
What does production-grade LLM inference look like?
If you wake up in the middle of the night thinking about that question, this blog post might be for you.
Inference engines like vLLM and SGLang sit on top of PyTorch (when I say vLLM going forward, that also includes SGLang and other supported inference engines). They optimize inference at the node level, most notably by managing KV cache and improving throughput through techniques such as paged attention and continuous batching.
The elevator pitch for llm-d is that it’s an LLM-aware load balancer.
If we have multiple vLLM instances, each has its own state: available GPU memory, KV cache prefix matches, number of requests queuing to be processed, etc. We can’t simply use round robin. Selecting the best inference engine instance based on these signals is what llm-d is all about.
It also provides features such as flow control to support different classes of requests based on priority (e.g., premium real-time traffic vs batch workloads). In addition, it enables smooth disaggregated P/D, where prefix and decode run on different nodes because they benefit from different configs. Prefix is compute-bound, while decode is memory-bandwidth-bound, so they have different GPU requirements: low TP for prefix to maximize compute, and high TP for decode to maximize memory bandwidth.
It also provides features such as flow control to support different classes of requests based on priority (e.g., premium real-time traffic vs batch workloads). In addition, it enables smooth disaggregated P/D, where prefix and decode run on different nodes because they benefit from different configs. Prefix is compute-bound, while decode is memory-bandwidth-bound, so they have different GPU requirements: low TP for prefix to maximize compute, and high TP for decode to maximize memory bandwidth.
The cool thing about llm-d is that it doesn’t reinvent the wheel. It builds on top of existing, established projects. LLM inference still happens on vLLM and SGLang, and we simply communicate with them via HTTP. The proxy layer and discovery are all built on top of Kubernetes(k8s) and Envoy. Even the integration point for deciding which vLLM instance to choose relies on an existing extension point, namely Envoy’s ext_proc extension. The data layer and metric collection are essentially built on top of Prometheus. So, no reinventing the wheel, just intelligently combining solid existing solutions. The best part is that for each layer and piece (metrics, scoring, flow control, etc.), llm-d is easily extensible with clear interfaces that new plugins can implement and use right away.
So if vLLM adds a new metric, or even if a new inference engine comes along, it can be easily added. If we want to implement a new way to pick or score, we just implement a new picker or scorer. If Prometheus falls out of favor, or if we want to use a custom monitoring solution, we can implement a plugin for that instead.
There’s a standardization effort to consolidate Generative AI inference on top of k8s, taking the form of the Gateway API Inference Extension, or GAIE . GAIE defines the API resources (like InferencePool ) and the Endpoint Picker role. llm-d’s router is an implementation of that EPP role, paired with an Envoy proxy that does the actual request forwarding.
That’s a very strategic choice. Rather than having llm-d be an isolated initiative with a bespoke API, it’s built as an implementation of a broader k8s standard.
It’s also worth mentioning that llm-d has a mode where it runs outside of k8s via a file discovery plugin, where the vLLM endpoints are hardcoded instead of discovered via the k8s API.
So what are these factors that need to be taken into account for LLM inference routing?
Say we have N1 and N2. If a user has already gotten a response from N1, this means the user’s request KV cache has already been calculated there, so a subsequent request can skip that step. If, however, we send the user’s request to N2, we need to re-run prefill and we won’t be taking advantage of the calculation already done on N1.
In other words, it’s efficient to route requests to nodes that already have the KV cache of the prompt (this consideration changes if we have KV cache offloading, in which the cache can be stored in shared network storage, for instance).
Another factor we care about is how much free VRAM N1 and N2 have. If N1’s VRAM is almost full, even though it has the user’s KV cache, it might not be the best choice, because we might need to wait for other requests to finish or evict them in order to run the request. If at the same time N2’s VRAM is free and ready to go, it might be best to rerun the prefill and make use of the free memory.
Both nodes might be running with little free VRAM, but N1 has only 2 requests in its queue waiting to be processed while N2 has 5. It’s best to route to the node with the smaller number of waiting requests.
There are many factors to consider when deciding which node to pick, and llm-d’s job is to pick. It gathers data from the various vLLM instances via the /metrics Prometheus endpoint, and keeps the state of each node in memory. When a user request comes in, it’s sent to llm-d’s router, called the Endpoint Picker or EPP, and the EPP will run a score of scorers (yes, the silly pun is intended): prefix-cache-scorer , kv-cache-utilization-scorer , queue-scorer , etc. for each possible candidate (vLLM instance), combining those scores and then picking one of the candidates. Before scoring and picking, there’s a filtering phase in which we can filter the candidate endpoints.
Envoy and the External Processing Filter
llm-d’s router doesn’t replace Envoy, it rides alongside it as an Envoy proxy filter target. When a request comes in, the Envoy proxy consults the llm-d EPP. The EPP does its thing and in the HTTP response it provides a header:
1
x-gateway-destination-endpoint: IP_PORT_OF_A_VLLM_POD
Then Envoy will route the request to that vLLM pod. That’s it, really.
Envoy Proxy has an extension (HTTP filter) called ext_proc , or External Processing . When a request comes in, it’s forwarded to an external process via gRPC, and that process can do a lot of things: mutate the HTTP body, add or remove headers, reject the request, etc. In our case, the llm-d router EPP is the external process, and it simply adds a header, the one with the address of the chosen vLLM instance.
How we make that choice is the magic that llm-d brings. And as we said above, it gathers data from the inference engines’ Prometheus endpoints:
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
var defaultEngineConfigs = [] engineConfigParams {
{
Name : "vllm" ,
QueuedRequestsSpec : "vllm:num_requests_waiting" ,
RunningRequestsSpec : "vllm:num_requests_running" ,
KVUsageSpec : "vllm:kv_cache_usage_perc" ,
LoRASpec : "vllm:lora_requests_info" ,
CacheInfoSpec : "vllm:cache_config_info" ,
},
{
Name : "sglang" ,
QueuedRequestsSpec : "sglang:num_queue_reqs" ,
RunningRequestsSpec : "sglang:num_running_reqs" ,
KVUsageSpec : "sglang:token_usage" ,
LoRASpec : "" ,
CacheInfoSpec : "sglang:cache_config_info" ,
CacheBlockSizeLabelName : "page_size" ,
CacheNumBlocksLabelName : "num_pages" ,
},
//...
}
And based on this data, it calculates a score for each candidate endpoint and picks one. Usually that’s the one with the highest score, but the picker can also sample according to score instead of just taking the max, which is a nice detail: always routing to the single best pod can create a thundering herd, where everyone piles onto the same “good” replica until it stops being good. Spreading the load based on score avoids that while still favoring the better candidates, this has a nice parallel to sampling a token after the softmax layer in an LLM.
Envoy establishes a gRPC connection with the EPP and streams it the request headers and body. Essentially, it’s:
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
33
34
func ( s * StreamingServer ) Process (
srv extProcPb . ExternalProcessor_ProcessServer ,
) error {
// Shared state for the lifetime of the request
reqCtx := NewRequestContext ()
// Stream reader
recvCh := startReceiver ( srv )
for {
select {
case msg := <- recvCh :
switch v := msg . Request . ( type ) {
case * RequestHeaders :
handleRequestHeaders ( reqCtx , v )
case * RequestBody :
handleRequestBody ( reqCtx , v ) // Actually: s.director.HandleRequest(ctx, reqCtx, parseResult.Body)
case * ResponseHeaders :
handleResponseHeaders ( reqCtx , v )
case * ResponseBody :
handleResponseBody ( reqCtx , v )
}
}
}
}
When we finish receiving the request body, we run through s.director.HandleRequest . The director holds the datastore that has the endpoints and their metrics.
We run the request through some admission plugins (so it can be rejected if needed, this is extensible like most of the pipeline, and custom rules can be added at different points to handle that).
We then go through flow control, where requests are treated based on their priority/flow id, specified via headers ( x-gateway-inference-fairness-id , x-gateway-inference-objective ). Flows can have different priorities, and within the same priority band we can follow fairness and ordering policies (round robin, first come first served, earliest deadline first, etc).
Once the request makes it through flow control, the endpoint candidates are grabbed and we run the “data producer” plugins for that request. These plugins augment the request with new data: tokenize the string by adding token ids (there’s a tokenizer that relies on a vLLM API ), and another key producer is the prefix matcher: how many of this request’s tokens already exist in the KV cache of each candidate pod, etc.
Then comes the essential part: the scheduler. We have a scheduling profile that has filters, scorers, and a picker. We filter the candidates, then score the request according to each scorer (of which there are many: prefix cache scorer, KV cache utilization, queue depth, etc.), then pick based on a combined score (each scorer can have its own weight).
1
2
3
4
5
6
7
8
9
10
Score:
Prefix cache scorer (weight=3.0): Pod 1 = 0.75, Pod 2 = 0.4, others = 0.0
KV-cache utilization (weight=1.0): Pod 1 = 0.6, Pod 2 = 0.8, others vary
Queue depth (weight=1.0): Pod 1 = 0.7, Pod 2 = 0.5, others vary
Final scores:
Pod 1: (0.75 × 3) + (0.6 × 1) + (0.7 × 1) = 3.55
Pod 2: (0.4 × 3) + (0.8 × 1) + (0.5 × 1) = 2.5
Others: < 2.0
Picker (weighted random): Pod 1 is strongly favored but not deterministic
Result: Pod 1 selected (IP: 10.0.1.42, port: 8000).
The EPP sends a response to Envoy with x-gateway-destination-endpoint set to the IP and port of the selected pod, and Envoy forwards the request there. As the pod generates tokens, Envoy streams the response back to the client. The ext_proc stream also receives response headers and body chunks for metric recording and data updates (new KV cache being calculated, etc).
If the model server (vLLM) already has part of the prompt’s KV cache, routing the request there avoids recomputing the prefix. We need to know about the content of each node’s KV cache.
There is an approximate prefix cache producer estimates each node’s KV cache from past routing decisions: if a prompt was previously sent to vLLM 1, it’s likely that its KV cache is still there.
The downside is that it’s only an approximation and doesn’t account for cache evictions.
There is a precise KV cache producer takes a much more interesting approach. vLLM exposes KV-cache allocation and eviction events over ZMQ (a brokerless

[truncated]
