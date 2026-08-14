---
source: "https://www.digitalocean.com/blog/llm-inference-benchmarking"
hn_url: "https://news.ycombinator.com/item?id=49300564"
title: "LLM Inference Benchmarking"
article_title: "LLM Inference Benchmarking - Measure What Matters | DigitalOcean"
author: "ankitg12"
captured_at: "2026-08-14T16:41:15Z"
capture_tool: "hn-digest"
hn_id: 49300564
score: 1
comments: 0
posted_at: "2026-08-14T15:58:51Z"
tags:
  - hacker-news
  - translated
---

# LLM Inference Benchmarking

- HN: [49300564](https://news.ycombinator.com/item?id=49300564)
- Source: [www.digitalocean.com](https://www.digitalocean.com/blog/llm-inference-benchmarking)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T15:58:51Z

## Translation

タイトル: LLM 推論ベンチマーク
記事のタイトル: LLM 推論ベンチマーク - 重要なものを測定する |デジタルオーシャン
説明: 推論ハードウェアのコストが依然として高いため、パフォーマンスを最大限に絞り込んでユニットエコノミクスを向上させることが AI チームの主な目的です。この記事では、LLM パフォーマンス ドメインに焦点を当て、レイテンシ、スループット、同時実行性、コストの間の相互作用を分析します。

記事本文:
LLM 推論ベンチマーク - 重要なものを測定する |デジタルオーシャンのブログ
クラウド コンピューティング リソースを構築、デプロイ、拡張する
コンテナとバックアップを安全に保存および管理する
一般的なデータベース エンジンを実行するフルマネージド リソース
インフラストラクチャを制御し、洞察を収集する
アプリへのトラフィックを保護して制御する
これらのセキュリティ機能でアカウントとリソースを保護します
あらゆる量のデータをクラウドに確実に保存し、アクセスします
クラウドセキュリティ体制管理 (CSPM)
ID とアクセス管理 (IAM)
クラウドセキュリティ体制管理 (CSPM)
ID とアクセス管理 (IAM)
LLM 推論ベンチマーク - 重要なものを測定する
Piyush Srivastava、Karnik Modi、Stephen Varela、Rithish Ramesh 著
実稼働グレードの LLM 推論は複雑なシステムの課題であり、ハードウェア プリミティブ (FLOP、メモリ帯域幅、インターコネクト) から高度なソフトウェア レイヤーに至るまで、スタック全体にわたる深い共同設計が必要です。 NVIDIA や AMD などの GPU プロバイダー間のハードウェアのばらつき (数値タイプのパフォーマンス (FP8、BF16、NVFP4 など)、HBM の帯域幅と容量、ピーク FLOP などの世代間の違いを含む) を考慮すると、最適なパフォーマンスが保証されることはありません。それは、プリフィル中の FLOP 使用率の最大化、デコード中の帯域幅効率の最大化、MoE モデルのエキスパート ルーティングの最適化、最適な並列処理戦略の発見などを行うソフトウェアの能力に依存します。
推論ハードウェアのコストが依然として高いため、パフォーマンスを最大限に絞り込んでユニットエコノミクスを向上させることが AI チームの主な目的です。私たちは現在、パフォーマンスとコスト効率を再定義する、ハードウェアとソフトウェアの共同設計の時代にいます。したがって、エンドツーエンドのモデルのパフォーマンス、分離されたコンポーネントのマイクロベンチマークという 3 つの重要な柱を追跡するためにベンチマークを進化させる必要があります。

パフォーマンスの向上を目指すための構造化された方法。この記事では、LLM パフォーマンス ドメインに焦点を当て、レイテンシ、スループット、同時実行性、コストの間の相互作用を分析します。
プリフィルとデコード: 推論の 2 つのフェーズ
LLM 推論は、プレフィルとデコードの 2 フェーズで動作します。プレフィルフェーズでは、入力全体がモデルのフォワードパスを通過します。これにはセルフアテンション、追加とノルムが含まれ、モデルのフィードフォワードネットワークの隠れ層を通過します。このフェーズは非常にコンピューティングに依存します。プレフィルフェーズの転送バイトあたりの FLOP (演算強度) は非常に高くなります。簡単に言うと、GPU はメモリからのデータを待つよりも計算に多くの時間を費やしています。
一方、デコードフェーズはメモリに依存します。生成されるすべてのトークンについて、デコードでは重み行列全体、HBM からの KV キャッシュをロードし、1 つのトークンを生成し、それを HBM に書き戻して次のトークンを計算する必要があります。 GPU は、計算に費やす時間よりも、メモリからのデータの待機に多くの時間を費やすことになります。したがって、これら 2 つのフェーズの計算特性は大きく異なります。
このセクションでは、推論パフォーマンスに対するユーザーの認識に直接影響を与えるいくつかの指標について説明します。
最初のトークンまでの時間は、プロンプトが送信されてからモデルが最初のトークンを生成するまでの時間を測定する重要な指標です。ユーザーの観点から見ると、これはプロンプトを送信してからテキスト ストリームが表示されるまでの「待機」時間です。
入力シーケンス (プロンプト) からの順方向パス全体が完了し、初期 KV キャッシュが設定されるまでモデルはトークンの生成を開始できないため、TTFT はプレフィル フェーズと密接に関係しています。演算強度 (転送バイトあたりの FLOPS) の最大化、最適な注意力 (MHA、MLA、GQA) k

アーネル、メモリ帯域幅効率 (FlashAttendant など)、最適な線形層、アクティベーション カーネル、最適なバッチ スケジューリングは、TTFT に直接影響を与えるプレフィル フェーズを最適化する特徴です。プリフィルとデコードがネットワーク境界を越えて分離されている分離セットアップ (たとえば、8x サーバー間の RDMA) では、オーケストレーション戦略に応じて、プリフィル ワーカーからデコード ワーカーへの KV キャッシュ転送レイテンシが TTFT に影響します。たとえば、一部の実装では、KV キャッシュがデコード ワーカーに転送される前または転送中に、プレフィル ワーカーから最初のトークンが生成されます。これらの実装では、転送は 2 番目のトークン (ITL) のレイテンシーに影響しますが、TTFT 自体には影響しません。
出力トークンあたりの時間は、ストリーミングが開始されてからテキストが流れる速度を測定します。これは、エンドツーエンドのメトリックとして計算され、エンドツーエンドの待ち時間から TTFT が取り除かれ、残りの間隔から生成されたトークンの総数で除算されます。したがって、TPOT は、テキスト フロー速度の全体的な認識をエンド ユーザーに提供します。ユースケースによっては、テキストの流れが遅すぎたり速すぎたりすることは、エンドユーザーにとって理想的なエクスペリエンスではない可能性があります。
TTFT はプレフィル フェーズに関連付けられていますが、TPOT はデコード フェーズのパフォーマンスの優れた代理です。デコードは厳密にメモリに依存しており、(プレフィル フェーズとは異なり) 一度に 1 つのトークンずつ実行されます。新しいトークンごとに、モデルがモデル重み行列全体と KV キャッシュを GPU HBM からロードし、トークンを生成して、KV キャッシュを GPU HBM に書き戻す必要があります。したがって、TPOT は GPU のメモリ帯域幅によって制限されます。例として、AMD MI355X (HBM 帯域幅 8 TB/秒) は、AMD MI325X (HBM 帯域幅 6 TB/秒) よりも優れた TPOT を持ちます。
トークン間遅延は TPOT とよく混同されますが、微妙な違いがあります。 TPOT が

エンドツーエンドのメトリクスである ITL は、2 つの連続するトークン間の時間経過 (たとえば、Tn+1 と Tn の間の時間) を厳密に測定します。 ITL は、デコード スループットと分散の両方を測定するのに役立つメトリックです。 ITL の分散が大きいと、トークン生成時にジッターが発生し、エンド ユーザーにとってはテキストがスムーズに流れていないように感じられるため、理想的ではありません。 ITL もデコード フェーズ メトリクスですが、少し低いレベルであり、1) FLOP を独占する大規模なプレフィル操作による干渉、2) PagedAttendant を通過するシステム、3) 最適ではないハードウェア トポロジ セットアップ (例: クロス NUMA ドメイン トラフィック)、およびその他のカーネル レベルの非効率など、いくつかの理由で発生する可能性のあるデコード パフォーマンスのジッターの測定に役立ちます。
エンドツーエンドの遅延は、ユーザーからのリクエストの継続時間を測定します。たとえば、ユーザーが 1K (ISL) プロンプトをモデルに送信し、8K (OSL) 応答を期待している場合、E2EL はユーザーから送信された最初のバイトからユーザーが受信した最後のトークンの最後のバイトまでの時間を測定します。したがって、E2EL は、以下のように計算される集計メトリックです。
E2EL = ネットワーク オーバーヘッド + TTFT + (TPOT * トークン数)
E2EL は、ミリ秒単位が重要なリアルタイム推論ワークロードにとって有用なメトリクスです。たとえば、北米でホストされているモデルの近くにいるユーザーの E2EL は、大西洋を越えてモデルにアクセスするユーザーの E2EL よりもはるかに小さくなります。 E2EL に影響を与える可能性のあるその他の要因としては、リクエスト キューイング、トークン化オーバーヘッドなどがあります。
トークン スループットは、すべてのアクティブなユーザー リクエストにわたって推論システムによって 1 秒あたりに生成されるトークンの合計数を測定します。これは、リクエストのスループットが測定するものとは異なります。たとえば、トークンのスループットは、デコードの多い (大規模な OSL) ワークロードでは非常に高くなる可能性がありますが、RP

Sは控えめかもしれない。 TPS は優れた利用プロキシであり、「光の速度」生成スループット、つまり他のすべて (カーネル、ソフトウェアなど) が完全に動作した場合にハードウェアから得られる最大スループットを計算するための裏計算を行うのは簡単です。デコードでは、モデルの重み全体と KV キャッシュを HBM からロードして、次のトークンを 1 つずつ繰り返し計算する必要があります。したがって、光の速度のスループットは大まかに計算できます。
単一リクエスト TPSsol = 帯域幅 / 移動されたデータ
これを具体的に言い換えると次のようになります。
単一リクエスト TPSsol = HBM 帯域幅 / (モデルの重み + (コンテキストの長さ * トークンごとの KV バイト))
モデルの重みは、パラメーターごとのバイト数を指します (例: FP16 の場合は 2 バイト、FP8 の場合は 1 バイト)。トークンあたりの KV バイトは、特定のモデル アーキテクチャ (層数、ヘッド、寸法、精度) によって異なります。
上の式は同時実行 (C=1) の場合です。より高い同時実行性のモデリングは、方程式に同時実行性の要素を追加することによって、同様の方法 (多少のニュアンスはありますが) で行うことができます。これにより、北極星として適切な「シリコンの限界」フロンティア尺度が得られます (コンテキストの長さが長くなると、TPSsol が減衰することに留意してください)。
リクエスト スループットは、単位時間 (秒、分など) ごとに推論システムによって処理されるエンドツーエンド ユーザーのリクエストの数を測定します。RPS は、システムによって処理されるタスクまたはユーザーの数を表す優れたプロキシです。 RPS は通常、より多くのリクエストを並行して処理することで向上します。たとえば、一度に 1 つのリクエストを順番に送信すると、GPU はほとんどの時間をアイドリングに費やすことになるため、ハードウェアをより多く活用するには、より多くの作業を並行して実行させる必要があります。ただし、並列リクエストでシステムをロードできる制限があり、それを超えると、上で説明した遅延メトリクスが影響を受け始めます。結局のところピークがある

リクエストのスループットはシリコンの限界に関係しており、特定のハードウェアで達成できる光速度 (SoL) RPS がどの程度であるかをモデル化する価値があります。 RPS の光速度のモデリングは、プリフィル コンピューティングとデコード メモリ アクセスの間の非線形相互作用を考慮する必要があるため、もう少し微妙です。
基本的な推論メトリクスのセットを理解したところで、スループットとレイテンシのメトリクスの間には複雑な相互作用があり、両者の間には一定のトレードオフがあることに注目する価値があります。さらに、リクエストの密度を階層化すると、最終的には最適化する必要がある次元がいくつか増えます。常に緊張状態にある主なベクトルは、レイテンシ、スループット、同時実行性の 3 つです。レイテンシーのみを最適化すると、トークンあたりのコストが低下し、スループットのみを最適化すると、レイテンシーが低下します。枠組みがなければ、すぐに「もぐらたたき」のようなゲームになってしまいます。基本的な指標があなたがどこにいるのかを示すものであれば、パレートフロンティアはあなたがどこにいるのかを示すものです。
上のフロンティア グラフは、同時実行サーフェス上の e2e レイテンシーに対するリクエストのスループットを追跡します。ワークロードに同時リクエストをさらに追加することで、スループットの向上を試みます。ただし、同時実行性を高めると、e2e レイテンシーに悪影響を及ぼす可能性があります。上のグラフは、1 ～ 128 の同時実行スイープのフロンティアの例を示しています。フロンティアは理想的な「発見された」構成であり、通常はベースラインのフロンティアから始めて、そこから最適化することをお勧めします。たとえば、フロンティアの右下にある構成は最小基準を満たしていないため、破棄できます。一般的な段階的なチェックリストは、ベンチマーク作業を開始するときに役立ちます。
ステップ 1: ベースラインのパレートフロンティアを確立する
私たちのcharacter.aiの経験から、

他の推論顧客の場合、本番グレードの AI 推論ワークロードは特定のワークロードを最適化しようとしており、ワークロードの形状は通常次の定義であることを知っています。
コストを最適化しながらレイテンシを目標にする
コストは、ユニット サーバーあたりにどれだけのリクエスト密度を詰めることができるかによって決まります (例: 1X H100 または 8X MI325x)。これは直接的に、より高い同時実行性を提供し、サーバーあたりのリクエストのスループットを向上させることになります。ベースライン フロンティアを確立するには、モデルの選択とワークロードの形状 (ISL / OSL) を使用して同時実行スイープ (1 ～ 128) を実行します。 vLLM のような推論エンジンには、入出力シーケンスの長さ (ISL / OSL)、同時実行性、リクエスト数などを構成できるベンチマーク機能が組み込まれています。このベースラインのフロンティアにより、最適化の階層化を開始する前の最低限の基準が何であるかをよく理解できるようになりました。
vLLM ベンチマーク コマンドの例:
vllm ベンチサーブ \
--model openai/gpt-oss-120b \
--バックエンド vllm \
--base-url https://crr2hb6t5bm19qcee7lmliq2q-public-dended-inference.do-infra.ai \
--endpoint /v1/completions \
--データセット名ランダム \
--ランダム入力長 250 \
--ランダム出力長 250 \
--プロンプト数 1000 \
--最大同時実行数 128
ステップ 2: 動作点を見つける
あらゆる AI ビジネスとアプリ

[切り捨てられた]

## Original Extract

As inference hardware costs remain high, squeezing maximum performance to improve unit economics is a primary objective for AI teams. This article focuses on the LLM performance domain and analyzes the interplay between latency, throughput, concurrency, and cost.

LLM Inference Benchmarking - Measure What Matters | DigitalOcean Blog
Build, deploy, and scale cloud compute resources
Safely store and manage containers and backups
Fully managed resources running popular database engines
Control infrastructure and gather insights
Secure and control traffic to apps
Help protect your account and resources with these security features
Store and access any amount of data reliably in the cloud
Cloud Security Posture Management (CSPM)
Identity and Access Management (IAM)
Cloud Security Posture Management (CSPM)
Identity and Access Management (IAM)
LLM Inference Benchmarking - Measure What Matters
By Piyush Srivastava , Karnik Modi , Stephen Varela , and Rithish Ramesh
Production-grade LLM inference is a complex systems challenge, requiring deep co-designs - from hardware primitives (FLOPs, memory bandwidth, and interconnects) to sophisticated software layers - across the entire stack. Given the hardware variability across GPU providers like NVIDIA and AMD - including generational differences in numeric type performance (FP8, BF16, NVFP4 etc), HBM bandwidth and capacity, peak FLOPs etc - optimal performance is never guaranteed. It depends on the software’s ability to maximize FLOPs utilization during prefill, maximize bandwidth efficiency during decode, optimize expert routing in MoE models, discover optimal parallelism strategies, and more.
As inference hardware costs remain high, squeezing maximum performance to improve unit economics is a primary objective for AI teams. We are currently in an era of intense hardware-software co-design that will redefine performance and cost efficiency. Consequently, benchmarking must evolve to track three critical pillars: end-to-end model performance, micro-benchmarking of isolated components and a structured way to go after performance improvements. This article focuses on the LLM performance domain and analyzes the interplay between latency, throughput, concurrency, and cost.
Prefill and Decode: The two phases of Inference
LLM Inference works in two-phases: prefill and decode . The prefill phase is where the entire input goes through the model’s forward pass which includes self-attention, add & norm, and pass through the hidden layers of the model’s feed forward network. This phase is extremely compute bound. FLOPs per byte transferred (arithmetic intensity) for the prefill phase is very high. In simpler terms, the GPU is spending more time computing than waiting for the data from memory.
On the other hand, the decode phase is memory bound. For every token that is generated, decode needs to load the entire weight matrix, KV cache from the HBM, generate one token, and write it back to the HBM to compute the next token. GPUs end up spending more time waiting for data from the memory than spending time in computation. These two phases, thus, have very different computation characteristics.
In this section, we’ll go through several metrics that have a direct impact on the user perception of Inference performance.
Time to First Token is a key metric that measures the duration between when a prompt is submitted to when the model generates its first token. From a user standpoint, this is the “waiting” time from submitting a prompt to seeing the text stream.
TTFT is closely tied to the prefill phase as the model cannot start generating any token until the entire forward-pass from the input sequence (prompt) is done and the initial KV cache is populated. Maximizing arithmetic intensity (FLOPS per byte transferred), optimal attention (MHA, MLA, GQA) kernels, memory bandwidth efficiency (e.g. FlashAttention), optimal linear layers, activation kernels, and optimal batch scheduling are the hallmarks of optimizing the prefill phase which has a direct impact on TTFT. In disaggregated setups, where prefill and decode are separated across network boundaries (for example, RDMA across 8x servers), KV cache transfer latency from prefill to decode workers impact TTFT depending on the orchestration strategy. For example - some implementations generate the first token from the prefill worker before or while the KV cache is being transferred to the decode worker. In those implementations, the transfer affects the latency of the second token (ITL), but not the TTFT itself.
Time per Output Token measures how fast the text flows once it starts streaming. This is calculated as an end-to-end metric, stripping out the TTFT from end-to-end latency and dividing by the total number of tokens generated from the remaining intervals. TPOT, thus provides an overall perception of text flow speed to the end user. Depending on the use case, text flowing too slowly or too fast may not be an ideal experience for an end-user.
While TTFT is tied to the prefill phase, TPOT is a good proxy to the performance of the decode phase. Decode is strictly memory-bound and happens one token at a time (unlike the prefill phase). Each new token requires the model to load the entire model weight matrix and KV cache from the GPU HBM, generate the token and write back the KV cache to GPU HBM. TPOT is therefore limited by the memory bandwidth of the GPU. As an example - AMD MI355X (HBM bandwidth of 8 TB/s) will have a better TPOT than AMD MI325X (HBM bandwidth of 6 TB/s).
Inter Token Latency is often confused with TPOT but there is a subtle difference. While TPOT is an end-to-end metric, ITL strictly measures the time lapse between two consecutive tokens (e.g. time between Tn+1 and Tn). ITL is a useful metric to both measure the decode throughput as well as variance. High variance in ITL is not ideal as it signals jitter in token generation, which to an end user feels like text not flowing smoothly. While ITL is also a decode phase metric, it is a bit lower level and helps in measuring jitters in decode performance that may be caused for several reasons like 1) Interference from large prefill operations monopolizing the FLOPs, 2) System going through PagedAttention, 3) Sub-optimal hardware topology setup (e.g. cross NUMA domain traffic) and other kernel level inefficiencies.
End to End latency is a measure of the duration of a request from a user. For example, if a user sent a 1K (ISL) prompt to the model and is expecting 8K (OSL) response, then E2EL measures the duration between the first byte sent from the user to the last byte of the last token received by the user. E2EL is therefore an aggregate metric that is somewhat calculated as below:
E2EL = Network Overhead + TTFT + (TPOT * Number of Tokens)
E2EL is a useful metric for real time inference workloads where every millisecond counts. For example, E2EL for a model hosted in North America for a user in close proximity will be much smaller than for a user accessing a model from across the Atlantic. Other factors that may impact E2EL are request queuing, tokenization overhead, and so on.
Token Throughput measures the total number of tokens generated by the inference system per second across all active user requests. This is different from what request throughput is supposed to measure, for example, token throughput can be really high for decode heavy (large OSL) workloads but RPS could be modest. TPS is a good utilization proxy, it is easy to do some back of the envelope math to calculate the “speed of light” generation throughput i.e. what’s the maximum throughput you can get from the hardware if everything else (kernels, software etc) worked perfectly. Decode has to load the entire model weights + KV cache from the HBM to compute the next tokens iteratively one-by-one. Therefore, speed of light throughput, roughly can be calculated:
Single Request TPSsol = Bandwidth / Data Moved
which translates in specific terms to:
Single Request TPSsol = HBM Bandwidth / (Model Weights + (Context Length * KV bytes per token))
Model weights refers to bytes per parameter (e.g. 2 bytes for FP16, 1 byte for FP8). KV bytes per token depends on specific model architecture (num_layers, heads, dimensions, precision).
The above formula is for concurrency (C=1). Modeling for higher concurrency can be done in a similar way (with some nuances) by adding concurrency factors to the equation. This gives a good “limit of silicon” frontier measure as the north star (keep in mind that TPSsol decays as the context length grows).
Request Throughput measures how many end to end user requests are handled by the inference system per unit of time - seconds, minutes etc. RPS is a good proxy for how many tasks or users are being served by the system. RPS is typically increased by handling more requests in parallel. For example, if you send 1 request at a time sequentially, the GPU is spending most of its time idling and to get more out of the hardware, you let it do more work in parallel. However, there is a limit to which you can load the system with parallel requests, beyond which the latency metrics as discussed above start getting affected. Ultimately, there is a peak of request throughput tied to the limits of the silicon and it is worth modeling what is the speed of light (SoL) RPS that can be achieved on specific hardware. Modeling RPS speed of light is a bit more nuanced, as it requires accounting for non-linear interaction between prefill compute and decode memory access.
Now that we understand the foundational set of inference metrics, it is worth noting that there is a complex interplay between the throughput and latency metrics with constant trade-offs between the two. Moreover, layer in the request density and you end up with several dimensions to optimize for. There are three main vectors that are in constant tension: latency, throughput, and concurrency. If you optimize for latency alone, your cost-per-token suffers, if you optimize for throughput alone, latency suffers. Without a framework, it soon becomes a game of “whack-a-mole”. If foundational metrics give you a picture of where you are, the Pareto frontier provides a picture of where you can be.
The frontier graph above tracks request throughput against e2e latency on a concurrency surface. By adding more concurrent requests to the workload, we attempt to increase the throughput. However, increasing the concurrency has an impact on negatively affecting the e2e latency. The graph above shows an example frontier for a 1-128 concurrency sweep. Frontier is the ideal “discovered” configuration and it is usually a good idea to start with the baseline frontier and optimize from there. For example, any configs that are below and right of the frontier do not meet the minimum bar and can be discarded. A general step-by-step checklist is helpful while starting any benchmarking effort:
Step 1: Establish a baseline Pareto frontier
With our experience with character.ai and other inference customers, we know that any production-grade AI Inference workload is trying to optimize a specific workload and the shape of the workload is usually definitions of:
Target latency while optimizing cost
Cost is determined on how much request density can be packed on a per unit server (e.g. 1X H100 or an 8X MI325x). This directly translates to serving higher concurrency to increase request throughput per server. To establish a baseline frontier, run a concurrency sweep (1-128) with your model choice and workload shape (ISL / OSL). Inference engines like vLLM have in-built benchmarking capabilities that allow to configure input / output sequence lengths (ISL / OSL), concurrency, number of requests and more. With this baseline frontier, we now have a good sense of what’s the minimum bar before we start layering in optimizations.
Example vLLM benchmark command:
vllm bench serve \
--model openai/gpt-oss-120b \
--backend vllm \
--base-url https://crr2hb6t5bm19qcee7lmliq2q-public-dedicated-inference.do-infra.ai \
--endpoint /v1/completions \
--dataset-name random \
--random-input-len 250 \
--random-output-len 250 \
--num-prompts 1000 \
--max-concurrency 128
Step 2: Find the operating point
Every AI business and ap

[truncated]
