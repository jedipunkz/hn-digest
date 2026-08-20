---
source: "https://research.triunalabs.com/articles/ai-native-ssd/"
hn_url: "https://news.ycombinator.com/item?id=49377911"
title: "Route the Work, Not Just the Data: GPUs, CPUs, and the Rise of AI-Native SSDs"
article_title: "Route the Work, Not Just the Data: GPUs, CPUs, and the Rise of AI-Native SSDs"
image: "https://research.triunalabs.com/articles/ai-native-ssd/images/05-same-problem.png"
author: "paulwoll"
captured_at: "2026-08-20T18:23:50Z"
capture_tool: "hn-digest"
hn_id: 49377911
score: 3
comments: 0
posted_at: "2026-08-20T17:54:19Z"
tags:
  - hacker-news
  - translated
---

# Route the Work, Not Just the Data: GPUs, CPUs, and the Rise of AI-Native SSDs

- HN: [49377911](https://news.ycombinator.com/item?id=49377911)
- Source: [research.triunalabs.com](https://research.triunalabs.com/articles/ai-native-ssd/)
- Score: 3
- Comments: 0
- Posted: 2026-08-20T17:54:19Z

## Translation

タイトル: データだけでなく作業をルーティングする: GPU、CPU、AI ネイティブ SSD の台頭
説明: LLM リクエストはさまざまな種類の作業ですが、GPU を必要とするのは一部のみです。優れた AI アーキテクチャが各操作を、ラップトップで再現可能な 102 GB のデータ移動ベンチマークを使用して実行できる最も安価な層にルーティングする理由。

記事本文:
トリウナ研究所
研究レポート
この記事のコードとデータ
Paul Wol · Triuna Labs Research · 2026 年 8 月 18 日
🧠 データだけでなく作業をルーティングする: GPU、CPU、AI ネイティブ SSD の台頭
大規模言語モデルについて考えるとき、私たちは GPU を思い浮かべる傾向があります。
それは理にかなっています。これらがなければ、現代の生成 AI は現在の規模では存在しなかったでしょう。
しかし、GPU は AI のますます重要なアーキテクチャ上の問題の 1 つも明らかにします。
最も高速に計算できる場所は、最も安価にデータを保存できる場所ではありません。
NVIDIA B200 GPU には 180 GB の HBM3e が搭載されており、そのメモリを通じて最大約 8 TB/秒でデータを移動できます。
階層のもう一方の端では、Micron が 245.76 TB 6600 ION SSD の出荷を 2026 年 5 月に開始しました。
1 つのドライブには、1 つの B200 の HBM の 1,000 倍を超えるデータを保存できます。
しかし、フラッシュは HBM の帯域幅や遅延に遠く及ばない動作をします。
数百ギガバイトの非常に高速なメモリと数百テラバイトの比較的安価な永続ストレージとの間にある大きなギャップは、アーキテクチャ上の興味深い問題を引き起こします。
ストレージが、AI データが GPU を待機する単なる場所ではなくなったらどうなるでしょうか?
一部の AI ワークロードがストレージ自体の近くで処理され、高価な GPU メモリに到達するために実際に何が必要かをインテリジェントなストレージ層が決定したらどうなるでしょうか?
実際、この研究は 30 年近く前に遡る研究の道をたどっています。
そして、LLM は、これを追求する最も強力な理由の 1 つを提供する可能性があります。
この記事の要旨は次のように明確に述べられています。
LLM リクエストは、1 つのモノリシックな計算ではありません。これにはさまざまな種類の作業が含まれますが、GPU が必要となるのは一部だけです。モデルの状態が GPU メモリを超えると、勝者のアーキテクチャは各操作をそれを実行できる最も安価な層 (GPU、CPU、またはますますインテリジェントなストレージ) にルーティングします。

このルートはデータの移動です。次の主要な最適化は、GPU を高速化することではありません。そもそも到達する必要があるデータ量が削減されています。
以下に続くものはすべて証拠です。すでに出荷されているもの、研究で証明されているもの、私が自分のハードウェアで測定したもの、そして純粋に推測のままのもの。
⚙️ 最初: LLM は実際に何をしているのですか?
LLM は、プロンプトに一致する文を巨大なデータベースから検索しているわけではありません。
単純化したレベルでは、膨大な量の数値計算を繰り返し実行します。
テキストはトークンに分割されます。
これらのトークンは数値ベクトルに変換され、 Transformer と呼ばれるニューラル ネットワーク アーキテクチャの多くの層を通過します。
トランスフォーマーは、2017 年の画期的な論文「Attention Is All You Need」で紹介されました。
特に 2 つの操作が重要です。
注意は、モデルが現在のトークンを解釈する際に、どの以前のトークンが重要であるかを判断するのに役立ちます。
簡略化されたバージョンでは、次の 3 つの表現が作成されます。
キー: 私はどのような情報を表しているのでしょうか?
値: どのような情報を転送する必要がありますか?
クエリはキーと比較されます。
注意スコアが計算されます。
これらのスコアは、さまざまな値が次の表現にどの程度強く影響するかを決定します。
各 Transformer レイヤーには、トークン表現を変換する大規模な学習済み行列も含まれています。
これにより、数十億のパラメータにわたって膨大な量の乗算と累積が発生します。
最終的に、モデルは次のトークンの可能性のある確率分布を生成します。
デコード戦略に従ってトークンを選択し、それをシーケンスに追加し、プロセスを再度実行します。
この繰り返しの数値ワークロードは、GPU がなぜそれほど重要になったかを説明するのに役立ちます。
🧮 LLM にとって GPU が CPU よりも優れている理由
CPU は非常に優れた汎用プロセッサです。
T

これらは次のようなワークロード向けに設計されています。
多くの異なる命令タイプ
CPU の強みは柔軟性です。
その大部分は、次のような質問を繰り返しています。
これらの膨大な数の配列をできるだけ早く、できるだけ並列処理で乗算できますか?
GPU は並列処理のために構築されています。
最新の AI GPU には、数千の実行ユニットに加えて、FP16、BF16、FP8 などの形式やますます低精度の表現を使用した行列演算用に設計された特殊な Tensor コアが含まれています。
また、非常に高速な高帯域幅メモリも搭載しています。
NVIDIA は、GPU あたり最大約 8 TB/秒の HBM 帯域幅の B200 をリストしています。
Micron の 9550 などの高性能 PCIe Gen5 SSD のシーケンシャル読み取り帯域幅は約 14 GB/秒に達します。
これらはまったく異なるパフォーマンス クラスです。
NAND フラッシュが単純に GPU HBM のドロップイン代替品となる、妥当なアーキテクチャはありません。
しかし、それは興味深い質問ではありません。
どれだけのデータがその境界を越える必要をまったくなくすことができるでしょうか?
目標は、SSD を GPU のように動作させることではありません。目標は、GPU の作業と不要なデータの送信を停止することです。
🧠 LLM 推論にメモリの問題がある
推論には通常、大きく分けて 2 つのフェーズが含まれます。
最初にプロンプ​​トを送信するとき、多くのプロンプト トークンを並行して処理できます。
このステージは非常に多くの計算を必要とする可能性があり、GPU に適切にマッピングされます。
新しいトークンはそれぞれ、以前のものから得られた情報に依存します。
この自己回帰プロセスは、推論が膨大なモデルの重みと増大する量のコンテキスト状態に繰り返しアクセスすることを意味します。
したがって、モデル アーキテクチャ、バッチ サイズ、ハードウェア、およびワークロードに応じて、ボトルネックは生の算術スループットからメモリ帯​​域幅とデータ移動に移る可能性があります。
これが FlashAttend がそのようなサービスになった理由の 1 つです。

重要な貢献。
FlashAttention は乗算を根本的に高速化するものではありません。
特に、GPU HBM と高速オンチップ SRAM の間の高価な移動を削減するために注意を再編成します。
その著者は、アテンションの最適化を I/O を意識した問題として明確に組み立てています。
GPU の内部であっても、データの移動はデータの計算と同じくらい重要になる可能性があります。
次に、その問題を GPU の外側に拡張します。
🗃️ KV キャッシュ: LLM ワーキングメモリ
アテンション中に、Transformer は前のトークンを表すキー テンソルと値テンソルを生成します。
これらのテンソルを保持しないと、システムは別のトークンが生成されるたびに以前のアテンション状態を繰り返し再計算します。
代わりに、推論エンジンはそれらを Key-Value キャッシュ (通常は KV キャッシュ と短縮されます) に保存します。
これにより、冗長な計算が大幅に削減されます。
同時リクエスト数
繰り返されるドキュメントのやり取り
驚くほどスケールが大きくなります。
NVIDIA は、Llama 3 70B の 128K トークン コンテキストがバッチ サイズ 1 で 1 人のユーザーに対して約 40 GB の KV キャッシュ メモリを消費する例を示しています。
40ギガバイトはモデルではありません。
この例では、これは 1 つのロングコンテキスト要求に関連付けられたキャッシュされたアテンション状態にすぎません。
数百または数千の同時リクエスト、永続的なエージェント、ドキュメントのワークフロー、または推論プロセスにわたって長いコンテキストを掛け合わせると、問題が明らかになります。
HBM は信じられないほど高速ですが、その能力は希少です。
🪜 新たな AI メモリ階層
AI インフラストラクチャではメモリを階層として扱う必要がますます増えています。
GPU SRAM / キャッシュ
↓
GPU HBM
↓
CPU メモリ
↓
ローカルNVMe SSD
↓
リモート/ネットワークストレージ
通常、各ステップでより多くの容量が提供されます。
通常、各ステップでは遅延と帯域幅が犠牲になります。
最新の推論ソフトウェアは、その階層を明示的に管理し始めています。
たとえば、NVIDIA Dynamo はサポートしています

GPU メモリを超えた KV キャッシュのオフロード。そのアーキテクチャは、KV キャッシュ ブロックを CPU メモリまたはローカル ストレージにスピルすることができ、より大きなコンテキストと以前に計算されたプレフィックスの再利用を可能にします。
NVIDIA の FlexKV の取り組みは、この概念を GPU、CPU、SSD ベースのストレージを含む階層全体に拡張します。
したがって、この記事の一部はもはや推測ではありません。
SSD はすでに LLM 推論メモリ階層の一部になりつつあります。
さらに興味深いのは、次に何が起こるかということです。
すべてのローカル AI フォーラムには毎週次のような質問が寄せられます。「モデルをコンピュータから実行できますか?」
外付けSSD？」それは 1 層早すぎた正しい質問であり、その答えは
今日のアーキテクチャが何であるかを正確に明らかにします。
通常: SSD にモデルが保存されます。ランタイムは重みをメモリにロードします。
計算はメモリから実行され、主に永続化のためにストレージが再度使用されます。
ランタイムがウェイト ストリーミングまたはオフロードを明示的にサポートしない限り、ドライブは
ロード時間に影響するだけで、他には何も影響しません。モデルがすでに RAM に収まっている場合は、より高速な
外付け SSD によって推論が向上するわけではありません。ストレージはコンテナではなく、コンテナです
参加者。
適合しないモデルに対する現在の戦略は圧縮です。広く共有されている
今週の記事では、Qwen 3.8-27B (55.6 GB) の圧縮について説明します。
ビジョン言語モデル、最小 11.55 GB なので、完全に 16 GB Mac 内で動作します
ミニを読むペースで。これは印象的な作品であり、作者は次のように正直に述べています。
制限: 圧縮されたビルドでは、依然として本格的なコーディング エージェントを確実に駆動することができません。
その告白が興味深い部分だ。モデルのフィッティングは次のことと同じではありません
ワークロードに合わせて。エージェントには常駐の重み以上のものが必要です: 永続的
コンテキスト、検索インデックス、ツールの状態、チェックポイント、時間とともに増加する KV キャッシュ
あらゆるステップ。圧縮は重みを縮小しますが、何も起こりません。
それ以外の場合は増加します

非常に有能なシステムもそれに伴います。 3 つの異なる
レバーが機能しているため、区別しておく価値があります。
モデルを圧縮すると、重みのサイズが小さくなります。
メモリ階層の最適化により、何がいつ常駐するかが決まります。
この記事の主題である AI ネイティブ ストレージは、何をすべきかを決定します。
移動、キャッシュされたままになるもの、データの近くで変換されるもの、変換されないもの
GPU に到達する必要があります。
🕰️ その考えは古いです。ワークロードは新しいです。
保存されたデータの近くで計算することは新しいアイデアではありません。研究者たちは 1998 年にアクティブ ディスク アーキテクチャを発表し、データ集約型の作業をデータがすでに存在する場所で実行できるようにプロセッサを組み込んだドライブを提案していました。彼らが書き留めた動機はまるで昨日草稿されたかのように読み取れます。作業の一部はデータが存在する場所で実行できるのに、なぜ膨大なデータセットを中央プロセッサに継続的に移動するのでしょうか?
それ以来、3つのことが起こりました。フラッシュは回転ディスクに取って代わり、最新の SSD はすでに小型コンピューターです。コントローラー、ファームウェア、並列 NAND チャネル、エラー修正、アドレス変換が組み込まれています。 FPGA を追加することでプログラム可能なものとなり、サムスンは圧縮、フィルタリング、検索、変換を目的として販売された SmartSSD として、まさにそれを 2 回出荷しました。そして、Storage Networking Industry Association (SNIA) は、定義された API と相互運用性の取り組みを備えたコンピューティング ストレージというアーキテクチャ用語をこの分野に与えました。 (詳細なツアーは、「コンピューティング ストレージの過去、現在、未来: 調査」です。)
したがって、自然な疑問は次のとおりです。アイデアが 25 年前に開発され、ハードウェアが出荷されたのであれば、なぜそれがどこにもないのですか? 25 年間、支配的なワークロードが十分な報いを与えていなかったからです。汎用クエリは、データに予期せぬ影響を与えます。フィルターをドライブに押し込んだことによる勝利は本物でしたが、ささやかなものでした。

ストレージのプログラミングにかかる​​ソフトウェアのコストはかかりませんでした。
変わったのは仕事量です。この記事の後半にある 2 つの数字が対として重要です。測定された検索クエリには 102.4 GB コーパスの 0.8% が必要であり、Kimi K3 の独立したアウトオブコア実装は、トークンごとに 2.78 兆のパラメータの 4% 未満でアクティブ化されます。 LLM 状態は巨大で、構造化されており、圧倒的にスキップ可能であり、どの部分が重要であるかは、データを理解するものによって事前に決定可能です。この選択性プロファイルは、コンピューティング ストレージが 25 年間を費やして待ち望んでいたものです。
🤖 LLM Research が計算を SSD に向けて移行し始める
最近のいくつかの研究プロジェクトは特に関連性があります。
そして注目すべきことに、この研究は現在、データセンター規模のサービス提供アーキテクチャからメモリに制約のある AI PC やエッジ システムに至るまで、幅広いシステムに及んでいます。
ストレージの問題は、巨大なハイパースケール クラスターに限定されません。
モデルの状態、コンテキスト容量、およびアクセラレータ メモリの間の同様の不一致は、高性能化したモデルが有限の GPU またはユニファイド メモリ リソースに遭遇する場合には常に発生します。
SmartANNS: ベクターの生息場所を検索
大規模な RAG およびベクトル検索システムには、数十億の埋め込みが含まれる場合があります。
2024 年、研究者らは 10 億規模の SmartANNS を発表しました。

[切り捨てられた]

## Original Extract

An LLM request is many kinds of work, and only some of it needs a GPU. Why the winning AI architecture routes each operation to the cheapest tier that can perform it - with a laptop-reproducible 102 GB data-movement benchmark.

Triuna Labs
Research repo
This article's code & data
Paul Woll · Triuna Labs Research · August 18, 2026
🧠 Route the Work, Not Just the Data: GPUs, CPUs, and the Rise of AI-Native SSDs
When we think about Large Language Models, we tend to picture GPUs.
That makes sense. Modern generative AI would not exist at its current scale without them.
But GPUs also expose one of AI's increasingly important architectural problems:
The fastest place to compute is not the cheapest place to keep data.
An NVIDIA B200 GPU has 180 GB of HBM3e and can move data through that memory at up to roughly 8 TB/s .
At the other end of the hierarchy, Micron began shipping its 245.76 TB 6600 ION SSD in May 2026 .
One drive can hold more than a thousand times as much data as a single B200's HBM.
But flash operates nowhere near HBM's bandwidth or latency.
That enormous gap between hundreds of gigabytes of extraordinarily fast memory and hundreds of terabytes of comparatively inexpensive persistent storage creates a fascinating architectural question:
What if storage stopped being merely the place where AI data waits for the GPU?
What if some AI workloads were processed near the storage itself, while an intelligent storage tier decided what actually needed to reach expensive GPU memory?
It actually follows a research path stretching back almost three decades.
And LLMs may provide one of the strongest reasons yet to pursue it.
Here is the thesis of this article, stated plainly:
An LLM request is not one monolithic computation. It is many kinds of work, and only some of it needs a GPU. As model state outgrows GPU memory, the winning architecture will route each operation to the cheapest tier that can perform it (GPU, CPU, or increasingly intelligent storage), and the cost that decides the route is data movement. The next major optimization is not making the GPU faster. It is reducing how much data has to reach it in the first place.
Everything that follows is the evidence: what already ships, what research demonstrates, what I measured on my own hardware, and what remains genuinely speculative.
⚙️ First: What Is an LLM Actually Doing?
An LLM is not searching a giant database for a sentence matching your prompt.
At a simplified level, it repeatedly performs enormous amounts of numerical computation.
Your text is divided into tokens .
Those tokens are converted into numerical vectors and passed through many layers of a neural-network architecture called the Transformer .
The Transformer was introduced in the landmark 2017 paper Attention Is All You Need .
Two operations are especially important.
Attention helps the model determine which previous tokens matter when interpreting the current token.
A simplified version creates three representations:
Key: What information do I represent?
Value: What information should be passed forward?
Queries are compared against keys.
Attention scores are calculated.
Those scores determine how strongly different values influence the next representation.
Each Transformer layer also contains large learned matrices that transform the token representations.
Across billions of parameters, this creates an enormous amount of multiplication and accumulation.
Eventually, the model produces a probability distribution over possible next tokens.
It selects a token according to the decoding strategy, appends it to the sequence, and performs the process again.
That repeated numerical workload helps explain why GPUs became so important.
🧮 Why GPUs Are Better Than CPUs for LLMs
CPUs are extraordinary general-purpose processors.
They are designed for workloads such as:
Many different instruction types
A CPU's strength is flexibility.
Huge portions of it repeatedly ask something closer to:
Can you multiply these enormous arrays of numbers as quickly and in as much parallelism as possible?
GPUs were built for parallelism.
Modern AI GPUs contain thousands of execution units plus specialized Tensor Cores designed for matrix operations using formats such as FP16, BF16, FP8 and increasingly lower-precision representations.
They also sit beside extraordinarily fast High Bandwidth Memory.
NVIDIA lists a B200 at up to roughly 8 TB/s of HBM bandwidth per GPU .
A high-performance PCIe Gen5 SSD such as Micron's 9550 reaches roughly 14 GB/s of sequential read bandwidth .
Those are completely different performance classes.
There is no plausible architecture in which NAND flash simply becomes a drop-in substitute for GPU HBM.
But that is not the interesting question.
How much data could we prevent from needing to cross that boundary at all?
The goal is not to make SSDs behave like GPUs. The goal is to stop sending the GPU work and data it does not need.
🧠 LLM Inference Has a Memory Problem
Inference generally contains two broad phases.
When you initially submit a prompt, many prompt tokens can be processed in parallel.
This stage can be highly compute-intensive and maps well to GPUs.
Each new token depends on information derived from what came before.
This autoregressive process means inference repeatedly accesses enormous model weights and an expanding amount of context state.
Depending on model architecture, batch size, hardware and workload, the bottleneck can therefore shift away from raw arithmetic throughput toward memory bandwidth and data movement .
This is one reason FlashAttention became such an important contribution.
FlashAttention does not make multiplication fundamentally faster.
It reorganizes attention specifically to reduce expensive movement between GPU HBM and faster on-chip SRAM.
Its authors explicitly frame attention optimization as an I/O-aware problem.
Even inside a GPU, moving data can become as important as computing it.
Now expand that problem outside the GPU.
🗃️ The KV Cache: LLM Working Memory
During attention, Transformers generate key and value tensors representing prior tokens.
Without retaining those tensors, the system would repeatedly recompute previous attention state every time another token was generated.
Instead, inference engines store them in a Key-Value cache , usually shortened to KV cache .
That dramatically reduces redundant computation.
Number of simultaneous requests
Repeated document interactions
The scale becomes surprisingly large.
NVIDIA has illustrated an example in which a 128K-token context for Llama 3 70B consumes roughly 40 GB of KV-cache memory for a single user at batch size 1 .
Forty gigabytes is not the model.
It is just the cached attention state associated with one long-context request in that example.
Multiply long contexts across hundreds or thousands of concurrent requests, persistent agents, document workflows or reasoning processes and the problem becomes obvious:
HBM is incredibly fast, but it is scarce.
🪜 The Emerging AI Memory Hierarchy
Increasingly, AI infrastructure has to treat memory as a hierarchy:
GPU SRAM / cache
↓
GPU HBM
↓
CPU DRAM
↓
Local NVMe SSD
↓
Remote / network storage
Each step generally offers more capacity.
Each step generally sacrifices latency and bandwidth.
Modern inference software is beginning to explicitly manage that hierarchy.
NVIDIA Dynamo, for example, supports KV-cache offloading beyond GPU memory. Its architecture can spill KV-cache blocks into CPU memory or local storage , allowing larger contexts and reuse of previously computed prefixes.
NVIDIA's FlexKV work extends this concept across tiers including GPU, CPU and SSD-backed storage .
So one part of this article is no longer speculative:
SSDs are already becoming part of the LLM inference memory hierarchy.
The more interesting question is what happens next.
Every local-AI forum gets this question weekly: "can I just run the model from an
external SSD?" It is the right question asked one layer too early, and the answer
exposes exactly what today's architecture is.
Ordinarily: the SSD stores the model. The runtime loads the weights into memory,
computation runs from memory, and storage is touched again mainly for persistence.
Unless the runtime explicitly supports weight streaming or offload, the drive
affects load time and nothing else. If the model already fits in RAM, a faster
external SSD does not make inference better. The storage is a container, not a
participant.
The current strategy for models that do not fit is compression. A widely shared
writeup this week walks through squeezing Qwen 3.8-27B, a 55.6 GB
vision-language model, down to 11.55 GB so it runs entirely inside a 16 GB Mac
mini at reading pace. It is impressive work, and its author is honest about the
limit: the compressed build still cannot reliably drive a serious coding agent.
That admission is the interesting part. Fitting the model is not the same as
fitting the workload. An agent needs more than resident weights: persistent
context, retrieval indexes, tool state, checkpoints, a KV cache that grows with
every step. Compression shrinks the weights and does nothing about everything
else that increasingly capable systems drag along with them. Three different
levers are in play, and they are worth keeping distinct:
Model compression reduces the size of the weights.
Memory-hierarchy optimization decides what is resident, and when.
AI-native storage , the subject of this article, would decide what should
move, what stays cached, what gets transformed near the data, and what never
needs to reach the GPU at all.
🕰️ The Idea Is Old. The Workload Is New.
Computing near stored data is not a new idea. Researchers were publishing Active Disk architectures in 1998 , proposing drives with embedded processors so that data-intensive work could happen where the data already lived, and the motivation they wrote down reads like it was drafted yesterday: why continuously move enormous datasets to a central processor when some of the work can happen where the data resides?
Three things have happened since. Flash replaced spinning disks, and a modern SSD is already a small computer: controllers, firmware, parallel NAND channels, error correction, address translation. Adding an FPGA made it a programmable one, and Samsung shipped exactly that, twice, as the SmartSSD, marketed for compression, filtering, search and transformation. And the Storage Networking Industry Association (SNIA) gave the field an architectural vocabulary: computational storage, with defined APIs and interoperability work. (A thorough tour is Past, Present and Future of Computational Storage: A Survey .)
So the natural question is: if the idea is twenty-five years old and the hardware shipped, why is it not everywhere? Because for twenty-five years the dominant workloads did not reward it enough. General-purpose queries touch data unpredictably. The win from pushing a filter into a drive was real but modest, and the software cost of programming storage was not.
What changed is the workload. Two numbers from later in this article make the point as a pair: a measured retrieval query needed 0.8% of a 102.4 GB corpus , and an independent out-of-core implementation of Kimi K3 activates under 4% of its 2.78 trillion parameters per token. LLM state is enormous, structured, and overwhelmingly skippable , and which parts matter is decidable in advance by something that understands the data. That selectivity profile is what computational storage spent twenty-five years waiting for.
🤖 LLM Research Starts Moving Computation Toward SSDs
Several recent research projects are particularly relevant.
And notably, this research now spans a wide range of systems, from datacenter-scale serving architectures to memory-constrained AI PCs and edge systems .
The storage problem is not limited to giant hyperscale clusters.
The same mismatch between model state, context capacity and accelerator memory appears wherever increasingly capable models encounter finite GPU or unified-memory resources.
SmartANNS: Search Where the Vectors Live
Large RAG and vector-search systems may contain billions of embeddings.
In 2024, researchers presented SmartANNS , a billion-scale a

[truncated]
