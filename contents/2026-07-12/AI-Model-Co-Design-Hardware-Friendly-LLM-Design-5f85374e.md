---
source: "https://developer.nvidia.com/blog/ai-model-co-design-hardware-friendly-llm-design/"
hn_url: "https://news.ycombinator.com/item?id=48883914"
title: "AI Model Co-Design: Hardware-Friendly LLM Design"
article_title: "AI Model Co-Design: Hardware-Friendly LLM Design | NVIDIA Technical Blog"
author: "matt_d"
captured_at: "2026-07-12T20:02:33Z"
capture_tool: "hn-digest"
hn_id: 48883914
score: 1
comments: 0
posted_at: "2026-07-12T19:35:27Z"
tags:
  - hacker-news
  - translated
---

# AI Model Co-Design: Hardware-Friendly LLM Design

- HN: [48883914](https://news.ycombinator.com/item?id=48883914)
- Source: [developer.nvidia.com](https://developer.nvidia.com/blog/ai-model-co-design-hardware-friendly-llm-design/)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T19:35:27Z

## Translation

タイトル: AI モデルの共同設計: ハードウェアに優しい LLM 設計
記事のタイトル: AI モデルの共同設計: ハードウェアに優しい LLM 設計 | NVIDIA テクニカル ブログ
説明: AI のパフォーマンスは 3 つの次元に集約されます。デプロイメントでは 3 つすべてのバランスをとる必要があります。応答が遅い場合は高い精度が無駄になり、生のスループットは…

記事本文:
AI モデルの共同設計: ハードウェアに優しい LLM 設計 | NVIDIA テクニカル ブログ
開発者
ホーム
購読する
関連リソース
エージェントAI / ジェネレーティブAI
AI モデルの共同設計: ハードウェアに優しい LLM 設計
いいね
嫌い
高スループット、低遅延の LLM 推論用のハードウェア対応トランスフォーマー モデル設計には、正方形に近い線形レイヤーの寸法、GPU タイル サイズ (128 の倍数、理想的には 256 または 512) への調整、および演算強度と GPU 使用率を最大化するための幅と奥行きのアスペクト比が必要です。
NVIDIA のエンドツーエンド ツール (TensorRT モデル オプティマイザーと LLM コンプレッサー) によってサポートされている NVFP4 量子化は、線形レイヤー全体で高いスループットと最小限の精度損失を実現し、コンピューティング バウンドとメモリ バウンドの両方のワークロードで Blackwell GPU を最大限に活用できるようにします。
TensorRT-LLM のエキスパート並列処理 (EP) と、パイプライン並列処理や Helix 並列処理などのハイブリッド並列戦略により、大規模な専門家混合モデルをマルチノード Blackwell NVLink システム全体に拡張でき、通信と負荷の不均衡のボトルネックを軽減しながらスループットと対話性のバランスをとることができます。
AI によって生成されたコンテンツでは、情報が不完全に要約されている可能性があります。重要な情報を確認してください。さらに詳しく
AI のパフォーマンスは次の 3 つの次元で決まります。
精度: モデルがどの程度適切に推論して出力を生成するか
スループット: データセンターが 1 秒あたりに生成できるトークンの数
インタラクティブ性: 遅延によって支配される、ユーザーに対するモデルの応答性
導入では 3 つすべてのバランスをとる必要があります。応答が遅い場合は高い精度が無駄になり、各ユーザーのエクスペリエンスが遅い場合は生のスループットはほとんど意味がありません。したがって、実際のシステムでは、精度、スループット、対話性が同時に最適化されます。
この投稿では、スループットと対話性、およびモデル設計の選択によって精度を犠牲にすることなく両方をどのように形成するかに焦点を当てています。

cy (精度のトレードオフが発生する場合はフラグを立てます)。
精度を固定すると、問題は 2 次元のパレート フロンティアになります。通常、一方の改善には他方のコストがかかります。目標は、フロンティア全体を外側に押し出して、曲線の下の領域を最大化することです (下の図 1 を参照)。
図 1. システムのスループットと対話性のパレート フロンティア
今日の最も著名な AI ワークロードである LLM から始めます。このトレードオフには 2 つの観点があります。1 つはフリート スループット (トークン/秒) を優先するシステム導入者、もう 1 つはユーザーは最初のトークンおよびトークン間のレイテンシの低下を重視します。トークン間のレイテンシーの逆数はトークン/秒/ユーザーであり、大きいほど応答性が高くなります。
単一の応答は次のようになります: プロンプト → [最初のトークンのレイテンシ] → 最初のトークン → [トークン間のレイテンシ] → 次のトークン など。
これは、モデルを最新のハードウェアで適切に実行したいと考えているモデル開発者のための実用的な入門書です。考え方はシンプルです。ハードウェアに合わせたモデルは、単に高速に動作するだけではありません。拡張性が向上し、コストが削減され、より幅広い採用が可能になります。
まず、2 つの軸に沿って変化する導入状況を検討します (以下の図 2 を参照)。
図 2. さまざまなワークロード体制 (横軸) とサービス目標 (縦軸) にわたるレイテンシの内訳
ワークロードは短いコンテキストから長いコンテキストまで多岐にわたり、サービスの目標はスループット指向 (1 秒あたりのトークンの最大化) からレイテンシー指向 (応答時間の最小化) まであります。多くはその中間に位置します。
各象限では、異なる最適化が必要です。ロングコンテキストのスループット指向のサービングは、ほとんどの時間をアテンションに費やしますが、レイテンシ指向のサービングは、通信と固定オーバーヘッドを犠牲にして、アテンションと FFN を短縮するためにモデルの並列処理を追加します。ショートコンテキスト、スループット指向のサービスにより、アテンションと FFN の間で時間をより均等に分割できるため、利点が得られます。

大規模な並列処理 (エキスパート並列処理など) から。
アムダールの法則が適用されます。つまり、1 つの部分を最適化しても、その部分が占める時間と同じだけ効果が得られます。注意が実行時間の 77% である場合、フィードフォワード層を調整してもわずかな利益しか得られません。注意の道は努力が報われる場所です。自分の体制を知ることで、どこに焦点を当てるべきかがわかります。
この投稿では、システム エンジニアでなくてもハードウェアを最大限に活用したい人向けに、これらの決定を早期に行うための簡単な経験則を示します。シリーズの後続の各章では、ハードウェアを意識した設計のさまざまな次元を対象としており、コンピューティングのボトルネックを回避し、データセンター規模でスムーズに導入し、設計をユースケースに合わせて、実際に拡張するのに役立ちます。
よりスマートにデザインしましょう。より迅速に展開します。より広いスケール。始めましょう。
LLM の線形層のハードウェアに優しい寸法設定
トランス設計の重要な選択は、アスペクト比、つまりモデルの幅と層数 (\(L\)) のバランスです。デコーダ スタイルのモデルでは、これらの要素がスタック全体での計算とメモリの分散方法を支配します。幅自体は、MLP レイヤーの隠れ次元 (\(H\)) と中間投影次元 (\(H’\)) の 2 つの次元によって設定されます。
\(H\)、\(H'\)、\(L\) を組み合わせることで、モデルが並列処理戦略にどの程度正確にマッピングされるか、またモデルが GPU 全体でどの程度適切にスケールされるかが決まります。この投稿では、線形レイヤーに焦点を当てて、これら 3 つの選択肢がスループット、対話性、スケーラビリティにどのような影響を与えるかを見ていきます。
どのハードウェアでも、達成可能なパフォーマンスはルーフライン モデルによって制限されます。ワークロードがどこに到達するかは、移動されたメモリのバイトごとに実行される計算操作の数として定義される、その演算強度によって決まります。
演算強度が低いワークロードは、メモリ帯域幅 (メモリ制限) によって制限されます。算数の集中力が高い人

ty は、デバイスのピーク コンピューティング スループット (FLOPS (コンピューティング バウンド) 単位) によって制限されます。最大のスループット (1 秒あたりのトークン数) が目標の場合、ハードウェアの演算能力がすべて使用されるように、ワークロードをコンピューティング領域に駆動する必要があります。
レイテンシに敏感なデコードはその逆で、低い同時実行性で実行され、メモリに依存するため、メモリ アクセス時間を短縮することで応答レイテンシが短縮されます。
図 3. ルーフラインのモデル。リッジポイントより下では、ワークロードはメモリ帯域幅によって制限されます。それを超えると、デバイスのピーク コンピューティング スループットによって異なります。ワークロードの演算強度 (バイトあたりの操作数) によって、ワークロードがどの領域に該当するかが決まります。
バイトあたりの操作数を増やす簡単な方法の 1 つはバッチ サイズを増やすことですが、モデルの形状も重要です。 \(H\) と \(H'\) が GEMM が計算依存かメモリ依存かをどのように決定するかを見てみましょう。単一のデバイスで実行する場合、\(H\) と \(H'\) は次の形式の GEMM の形状を設定します。
一般的な線形代数ライブラリで使用される規則に従って、 \(A\) はサイズ \(M \times K\) (\(M\) 行、 \(K\) 列) の行列で、 \(B\) は \(K \times N\) です。積 \(C\) には \(M \times N\) の出力があり、それぞれ長さ \(K\) の内積であり、 \(M \times N \times K\) の融合積和演算 (FMA) が必要です。各 FMA は 2 FLOP (1 つの乗算、1 つの加算) であるため、コストは次のようになります。
\(\text{FLOPs} = 2 \times M \times N \times K\)
\(\text{読み取りバイト数} = M \times K \times \text{バイト}_A + N \times K \times \text{バイト}_B\)
\(\text{書き込みバイト数} = M \times N \times \text{バイト}_C\)
ここで、 \(\text{bytes}_{A,B,C}\) は、使用される精度によって設定される要素ごとのバイト数です。 \(A\) を入力、\(B\) を重みとして、以下の表 1 は、トークン数 (\(\text{Tokens} = \text{concurrency} \times \text{sequence length}\))、\(H\)、および \(H'\) を GEMM にマッピングします。

各線形層の \(M\)、\(N\)、および \(K\)。
表 1. トークン、H、および H' で表される、トランス ブロック内の線形層の GEMM 次元 (M、N、K)
すべての GEMM が「正方形」、\(\text{Tokens} = H' = H\) である場合を考えてみましょう。単一の GEMM は、メモリの \(3H^{2}\) 要素を移動しながら \(2H^{3}\) FLOP を実行するため、演算強度は \(H\) とともに増加します。実際の結果: \(H\) または \(H'\) が小さい場合、トークンの次元が大きい場合でも、GPU は計算よりもデータの移動に比例して多くの時間を費やします。
具体的な例として、GB300 で 4 ビット入力と 8 ビット出力を使用し、意図的に小さい \(H'=512\) と \(H=8192\) を備えた FFN-2 を考えてみましょう。以下の表 2 に示すように、トークン数が多い (GEMM-M が高い) 場合でも、この層はメモリに束縛されたままであり、データの移動が大半を占めます。出力が FP8 で入力が FP4 であるため、書き込みコストが支配的であり、縮小次元が小さいため GEMM のメモリ制限が維持されます。
以下の図 4 のプロットは、GB300 シリコンでのこれを示しています。 NVFP4 GEMM と大きな 8192 に固定されたトークン ディメンション (GEMM-M) を使用して、縮小ディメンション (GEMM-K) (左側に表示) と投影ディメンション (GEMM-N) (右側に表示) をスイープします。どちらの場合も、スイープ次元が小さいとスループットが急激に低下し、 \(N\) または \(K\) が小さいとハードウェアが十分に活用されないことが確認されます。 (\(N\) と \(K\) は、層に応じて \(H\) または \(H’\) に対応します。上記の表 1 を参照してください。
図 4 (左)。 GB300 での NVFP4 GEMM スループットと削減次元 \(K\) の比較 (\(M\) = 8192、\(N\) = 9728 固定)。 \(K\) ≈ 6144 付近で飽和します。持続スループットの 80% に達するには、\(K\) > 3072 が必要です。図 4 (右)。 GB300 での NVFP4 GEMM スループットと投影次元 \(N\) の比較 (\(M\) = 8192、\(K\) = 8448 固定)。近くで飽和します

\(N\) ≈ 6144;持続スループットの 80% に達するには \(N\) > 2560 が必要です
これはモデル設計者にとって重要な点です。モデルのディメンションは、計算が飽和状態になるバッチ サイズと同じくらい重要です。
ガイドライン 1: 固定パラメーター モデルの場合は、正方形に近い重み行列を優先し、投影または縮小の次元を小さくしすぎないようにします。
しかし、サイズだけでは十分ではありません。 Tensor コアの高い使用率を達成するには、GEMM のディメンションも、基礎となるタイリング ジオメトリにきれいにマッピングする必要があります。アライメントが不十分だとタイル量子化が発生し、演算強度が高い場合でもスループットが低下します。
GPU は、出力行列をタイルに分割することによって GEMM を実行し、各タイルはストリーミング マルチプロセッサ (SM) によって計算されます。最近の GPU では、SM が単独で動作する必要はありません。これらは単一の大きなタイル上で連携できます。clusterMMA を使用すると、2 つの隣接する SM が 1 つのタイル上で力を合わせ、Cooperative Grid Array (CGA) により SM のグループ全体が 1 つのクラスターとして連携できます。
連携によりデータの再利用が向上しますが、有効なタイルも拡大されるため、調整を維持するにはディメンションをより大きな値の倍数にする必要があります。ディメンションがその有効なタイルの倍数ではない場合、エッジ タイルは部分的にしか埋められていませんが、タイル全体に相当する計算を起動して実行します。未使用の (パディングされた) 部分は有用な作業を行わず、サイクルが無駄になり、スループットが低下します。図 5 はこれを示しています。256×128 ベース タイル、clusterMMA、および 4×2 CGA では、N が 256 (clusterMMA からの 2×128) または 512 (CGA からの 2×256) の倍数である場合、きめの細かいステップで GEMM-N をスイープすると、局所的なスループット最大値が生成されます。
図 5. GB300 で提供されるスループットに対するタイル量子化の影響。M×N 出力行列上でクラスターMMA および 4×2 CGA を備えた 256×128 タイルの使用を強制されたカーネルで測定
この無駄を避けるために、モデル di を選択してください。

タイル サイズの大きな倍数であり、GPU キャッシュ ライン幅に合わせて調整されたメンション。 128 の倍数は安全で持ち運び可能なフロアです。 256 (clusterMMA) または 512 (CGA) の倍数は、より大きな協調タイルと一致し、最大のスループットを獲得します。
ガイドライン 2: GPU タイル サイズとキャッシュ ライン幅に合わせてモデルの寸法を少なくとも 128 の倍数にし、clusterMMA と CGA によって形成される大きなタイルに一致させるには 256 または 512 を優先します。
幅の広いモデルは、深いモデルよりもハードウェアに優しい
固定パラメータのバジェットの場合、より広いモデルは、より大きな重みの再利用とより短い順次クリティカル パスを通じて、より深いモデルよりも高い演算強度とより低いレイテンシを提供します。これにより、スループット指向とレイテンシ指向の両方のサービス目標に対して、より幅広いモデルが有利になります。
そうは言っても、アスペクト比はモデルの品質にも影響します。深度は表現力に寄与するため、「広ければ広いほど良い」というわけではなく、有用な幅と深度の帯域が存在します。モデルの幅を広くするためだけにレイヤーを削除するのではなく、精度が保てる範囲でのみ幅を優先します。
ガイドライン 3: 選択肢が与えられた場合は、多数の小規模な操作よりも、少数で大規模な操作を優先します。これにより、演算強度が最大化され、ハードウェアの使用率が向上し、スループットと対話性の両方が向上します。つまりw

[切り捨てられた]

## Original Extract

AI performance comes down to three dimensions: Deployments must balance all three: High accuracy is wasted if responses are slow, and raw throughput means…

AI Model Co-Design: Hardware-Friendly LLM Design | NVIDIA Technical Blog
DEVELOPER
Home
Subscribe
Related Resources
Agentic AI / Generative AI
AI Model Co-Design: Hardware-Friendly LLM Design
Like
Dislike
Hardware-aware transformer model design for high-throughput, low-latency LLM inference requires near-square linear layer dimensions, alignment to GPU tile sizes (multiples of 128, ideally 256 or 512), and a width-over-depth aspect ratio to maximize arithmetic intensity and GPU utilization.
NVFP4 quantization, supported by NVIDIAs end-to-end tooling (TensorRT Model Optimizer and LLM Compressor), delivers high throughput and minimal accuracy loss across linear layers, enabling both compute-bound and memory-bound workloads to fully utilize Blackwell GPUs.
Expert parallelism (EP) and hybrid parallel strategies such as pipeline parallelism and Helix Parallelism in TensorRT-LLM allow large Mixture-of-Experts models to scale across multi-node Blackwell NVLink systems, balancing throughput and interactivity while mitigating communication and load imbalance bottlenecks.
AI-generated content may summarize information incompletely. Verify important information. Learn more
AI performance comes down to three dimensions:
Accuracy: How well the model reasons and produces outputs
Throughput: How many tokens per second a datacenter can generate
Interactivity: How responsive the model feels to a user, dominated by latency
Deployments must balance all three: High accuracy is wasted if responses are slow, and raw throughput means little if each user’s experience is laggy. Practical systems therefore optimize accuracy, throughput, and interactivity together.
This post focuses on throughput and interactivity, and how model-design choices shape both without sacrificing accuracy (we flag accuracy trade-offs where they arise).
Holding accuracy fixed, the problem becomes a two-dimensional Pareto frontier: improving one usually costs the other. The goal is to push the whole frontier outward, maximizing the area under the curve (see Figure 1, below).
Figure 1. System throughput versus interactivity Pareto frontier
We will start with LLMs, today’s most prominent AI workload. The trade-off has two perspectives: the system deployer, who prioritizes fleet throughput (tokens/sec), and the user , who values lower first-token and inter-token latency. The inverse of inter-token latency is tokens/sec/user, where higher means greater responsiveness.
A single response looks like: prompt → [first-token latency] → first token → [inter-token latency] → next token, and so forth.
This is a practical primer for model developers who want their models to run well on modern hardware. The idea is simple: models that align with the hardware don’t just run faster; they scale better, cost less, and reach wider adoption.
First, consider the deployment landscape, which varies along two axes (see Figure 2, below).
Figure 2. Latency breakdown across different workload regimes (horizontal axis) and service goals (vertical axis)
Workloads range from short to long context, and service goals from throughput-oriented (maximize tokens/sec) to latency-oriented (minimize response time); many fall in between.
Each quadrant needs different optimizations: long-context, throughput-oriented serving spends most of its time in attention, while latency-oriented serving adds model parallelism to shorten attention and FFN, at the cost of communication and fixed overheads. Short-context, throughput-oriented serving splits time more evenly across attention and FFN, and can benefit from parallelism (such as expert parallelism) at scale.
Amdahl’s law applies: optimizing one part helps only as much as the time it occupies. If attention is 77% of runtime, tuning the feed-forward layers yields only marginal gains; the attention path is where effort pays off. Knowing your regime tells you where to focus.
This post gives simple rules of thumb to make these decisions early, for anyone who wants the most from the hardware without being a systems engineer. Each subsequent chapter of the series will target a different dimension of hardware-aware design, helping you avoid compute bottlenecks, deploy frictionlessly at datacenter scale, match design to your use case, and scale in practice.
Design smarter. Deploy faster. Scale wider. Let’s begin.
Hardware-friendly dimensioning of linear layers in LLMs
A key transformer design choice is its aspect ratio: the balance between model width and number of layers (\(L\)). In decoder-style models, these factors dominate how computation and memory are distributed across the stack. Width itself is set by two dimensions, the hidden dimension (\(H\)) and the intermediate projection dimension (\(H’\)) in the MLP layers.
Together, \(H\), \(H’\), and \(L\) also shape how cleanly the model maps onto parallelism strategies and how well it scales across GPUs. This post looks at how these three choices affect throughput, interactivity, and scalability, with a close focus on the linear layers.
On any hardware, achievable performance is bounded by the roofline model. Where a workload lands depends on its arithmetic intensity, defined as the number of compute operations performed per byte of memory moved.
Workloads with low arithmetic intensity are capped by memory bandwidth (memory-bound); those with high arithmetic intensity are capped by the device’s peak compute throughput, in FLOPS (compute-bound). When the goal is maximum throughput (tokens per second), you want to drive the workload into the compute-bound region so the hardware’s full math capacity is used.
Latency-sensitive decoding is the opposite, it runs at low concurrency and is memory-bound, so reducing memory-access time is what lowers response latency.
Figure 3. The roofline model. Below the ridge point, a workload is limited by memory bandwidth; above it, by the device’s peak compute throughput. A workload’s arithmetic intensity (operations per byte) determines which regime it falls in
One straightforward way to raise operations per byte is to increase batch size, but model shape matters too. Let’s look at how \(H\) and \(H’\) decide whether a GEMM is compute-bound or memory-bound. When running on a single device, \(H\) and \(H’\) set the shape of a GEMM of the form:
Following the convention used in typical linear-algebra libraries, \(A\) is a matrix of size \(M \times K\) (\(M\) rows, \(K\) columns) and \(B\) is \(K \times N\). The product \(C\) has \(M \times N\) outputs, each a dot product of length \(K\), requiring \(M \times N \times K\) fused multiply-adds (FMAs). Since each FMA is 2 FLOPs (one multiply, one add), the costs are:
\(\text{FLOPs} = 2 \times M \times N \times K\)
\(\text{Read Bytes} = M \times K \times \text{bytes}_A + N \times K \times \text{bytes}_B\)
\(\text{Write Bytes} = M \times N \times \text{bytes}_C\)
where \(\text{bytes}_{A,B,C}\) are the per-element byte counts set by the precision used. Taking \(A\) as the inputs and \(B\) as the weights, Table 1, below, maps the token count (\(\text{Tokens} = \text{concurrency} \times \text{sequence length}\)), \(H\), and \(H’\) to a GEMM’s \(M\), \(N\), and \(K\) for each linear layer.
Table 1. GEMM dimensions (M, N, K) of the linear layers in a transformer block, expressed in terms of Tokens, H, and H′
Consider a case where every GEMM is “square,” \(\text{Tokens} = H’ = H\). A single GEMM then performs \(2H^{3}\) FLOPs while moving about \(3H^{2}\) elements of memory, so arithmetic intensity grows with \(H\). The practical consequence: when either \(H\) or \(H’\) is small, the GPU spends proportionally more time moving data than doing math, even when the token dimension is large.
As a concrete example, take FFN-2 with a deliberately small \(H’=512\) and \(H=8192\), using 4-bit inputs and 8-bit outputs on GB300. As Table 2 shows, below, even at large token counts (high GEMM-M) this layer stays memory-bound, dominated by data movement; the write cost dominates because the output is FP8 while the inputs are FP4, and the small reduction dimension keeps the GEMM memory-bound.
The plots in Figure 4, below, show this on GB300 silicon. With NVFP4 GEMMs and the token dimension (GEMM-M) fixed at a large 8192, we sweep the reduction dimension (GEMM-K) (shown on the left) and the projection dimension (GEMM-N) (shown on the right). In both cases throughput falls sharply when the swept dimension is small, confirming that a small \(N\) or \(K\) leaves the hardware underutilized. (\(N\) and \(K\) correspond to \(H\) or \(H’\) depending on the layer; see Table 1, above.)
Figure 4 (Left). NVFP4 GEMM throughput on GB300 vs the reduction dimension \(K\) (\(M\) = 8192, \(N\) = 9728 fixed). Saturates near \(K\) ≈ 6144; reaching 80% of sustained throughput requires \(K\) > 3072. Figure 4 (Right). NVFP4 GEMM throughput on GB300 vs the projection dimension \(N\) (\(M\) = 8192, \(K\) = 8448 fixed). Saturates near \(N\) ≈ 6144; reaching 80% of sustained throughput requires \(N\) > 2560
This is an important point for model designers: model dimensions matter just as much as batch size for saturating compute.
Guideline 1: For a fixed-parameter model, favor near-square weight matrices and avoid making either the projection or reduction dimension too small.
But size alone isn’t enough. To reach high Tensor Core utilization, a GEMM’s dimensions must also map cleanly onto the underlying tiling geometry; poor alignment causes tile quantization, which reduces throughput even when arithmetic intensity is high.
GPUs execute GEMMs by splitting the output matrix into tiles, each computed by a streaming multiprocessor (SM). On recent GPUs, SMs don’t have to work alone; they can cooperate on a single, larger tile: with clusterMMA, two neighboring SMs join forces on one tile, and a Cooperative Grid Array (CGA) lets a whole group of SMs work together as one cluster.
Cooperation improves data reuse, but it also enlarges the effective tile, so dimensions must be a multiple of a larger value to stay aligned. When a dimension is not a multiple of that effective tile, the edge tiles are only partially filled yet still launch and run a full tile’s worth of compute; the unused (padded) portion does no useful work, wasting cycles and lowering throughput. Figure 5 illustrates this: with a 256×128 base tile, clusterMMA, and a 4×2 CGA, sweeping GEMM-N in fine-grained steps produces local throughput maxima when N is a multiple of 256 (2×128 from clusterMMA) or 512 (2×256 from CGA).
Figure 5. Tile-quantization effect on delivered throughput on GB300, measured with a kernel forced to use 256×128 tiles with clusterMMA and a 4×2 CGA over the M×N output matrix
To avoid this waste, choose model dimensions that are large multiples of the tile size and aligned with GPU cache-line widths. A multiple of 128 is a safe, portable floor; multiples of 256 (clusterMMA) or 512 (CGA) align with the larger cooperative tiles and capture the most throughput.
Guideline 2: Make model dimensions multiples of 128 at minimum to align with GPU tile sizes and cache-line widths, and prefer 256 or 512 to match the larger tiles formed by clusterMMA and CGA.
Wider models are more hardware-friendly than deeper models
For a fixed-parameter budget, wider models offer higher arithmetic intensity and lower latency than deeper models through greater weight reuse and a shorter sequential critical path. This makes wider models advantageous for both throughput- and latency-oriented service goals.
That said, aspect ratio also affects model quality: depth contributes to representational power, so there is a useful width-to-depth band rather than “wider is always better.” Favor width only as far as accuracy holds, rather than stripping out layers just to make the model wider.
Guideline 3: When given a choice, prefer fewer, larger operations over many smaller ones; this maximizes arithmetic intensity and improves hardware utilization, benefiting both throughput and interactivity. In other words, w

[truncated]
