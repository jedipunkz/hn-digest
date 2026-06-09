---
source: "https://aarushgupta.io/posts/kan-fpga/"
hn_url: "https://news.ycombinator.com/item?id=48466277"
title: "Ultrafast machine learning on FPGAs via Kolmogorov-Arnold Networks"
article_title: "Ultrafast machine learning on FPGAs via Kolmogorov-Arnold Networks | Aarush Gupta"
author: "ag2718"
captured_at: "2026-06-09T21:39:52Z"
capture_tool: "hn-digest"
hn_id: 48466277
score: 92
comments: 13
posted_at: "2026-06-09T19:21:45Z"
tags:
  - hacker-news
  - translated
---

# Ultrafast machine learning on FPGAs via Kolmogorov-Arnold Networks

- HN: [48466277](https://news.ycombinator.com/item?id=48466277)
- Source: [aarushgupta.io](https://aarushgupta.io/posts/kan-fpga/)
- Score: 92
- Comments: 13
- Posted: 2026-06-09T19:21:45Z

## Translation

タイトル: Kolmogorov-Arnold Networks を介した FPGA での超高速機械学習
記事のタイトル: Kolmogorov-Arnold Networks を介した FPGA での超高速機械学習 |アールシュ・グプタ
HN テキスト: https://web.archive.org/web/20260609200156/https://aarushgup...

記事本文:
Kolmogorov-Arnold Networks を介した FPGA での超高速機械学習
この投稿は、コルモゴロフ-アーノルド ネットワーク (KAN) アーキテクチャを使用した超高速推論とオンライン学習のためのハードウェア アーキテクチャの設計を含む、私の修士論文の概要を説明したものです。標準的な機械学習の概念に精通しており、ハードウェアとデジタル回路についてもある程度の理解があることを前提とします。後者については、ここにある私の前回の投稿を読んでください。
詳細、特にベンチマークと注目すべき結果の詳細については、以下の 2 つの論文をお読みください。
[FPGA 2026 最優秀論文]
Duc Hoang * 、Aarush Gupta * 、および Philip C. Harris。 「カネレ: 効率的な LUT ベースの評価のためのコルモゴロフ – アーノルド ネットワーク」フィールド プログラマブル ゲート アレイに関する 2026 年の ACM/SIGDA 国際シンポジウムの議事録。 ACM、2026年。 https://dx.doi.org/10.1145/ 3748173.3779202
[ICML2026]
Duc Hoang * 、Aarush Gupta * 、フィリップ・ハリス。 「コルモゴロフ・アーノルドネットワークにおけるスプライン局所性を介した超高速オンFPGAオンライン学習」 arXiv プレプリント arXiv:2602.02056 、2026。 https://arxiv.org/abs/2602.02056
FPGA での機械学習の事例
最新の機械学習ワークロードのほとんどは、トレーニングであれ推論であれ、グラフィックス プロセッシング ユニット (GPU) 上で実行されます。高度な並列実行モデルをサポートするハードウェア アーキテクチャを通じて、GPU は大量のデータに対して非常に高いスループットで単純な操作を実行できます。そのため、大規模なアーキテクチャやバッチ形式のトレーニングと推論を伴う機械学習の問題に最適です。
ただし、複雑な GPU アーキテクチャでは、超低レイテンシ (マイクロ秒未満のレイテンシなど) と高いハードウェア効率を必要とするアプリケーションの要求を満たすことができません。プロセッサー (CPU や GPU など) は、スケジューリングや、

命令の最適化、メモリへの動的アクセスなど。超低レイテンシ ($\sim$nano秒など) と効率要件を備えた非常に特殊なワークロードには、代わりにカスタム ハードウェア アクセラレータの方が適切に対応します。
フィールド プログラマブル ゲート アレイ (FPGA) は、このスタイルのカスタム ハードウェア アクセラレーションに非常に適した再構成可能なデジタル ロジック デバイスです。 FPGA にはルックアップ テーブル (LUT) が含まれており、バイナリ入力のあらゆる組み合わせの出力値を列挙することでデジタル関数を表します。状態を保存するフリップフロップ (FF)。およびその他のメモリおよび計算プリミティブ。これらのコンポーネントとコンポーネント間の接続は、カスタム デジタル回路を設計するために再構成され、低レベルのハードウェア アーキテクチャとアルゴリズムの共同設計が可能になり、超高速の機械学習が可能になります。重要なのは、ニューラル ネットワークは、プロセッサ上で順次実行される命令としてではなく、デジタル ロジックとして直接実装されることです。
FPGA およびその他のデジタル デバイスは、基本的に連続値ではなくビットで動作します。ただし、ニューラル ネットワークでの算術演算 ($\times、+$ など) は、実数 $\mathbb R$ に対して行われるとよく​​考えられます。したがって、実数をビット列 (ビットのシーケンス) としてエンコードする必要があります。これは、量子化として知られるプロセスです。加算や乗算などの演算は 2 項関数になります。
これを行うための 1 つの方法は、固定小数点量子化です。
固定小数点量子化は、数値を基数 2 で表し、小数点の後にいくつかのビット (小数ビット) が続きます。たとえば、小数点以下の小数点以下 4 ビットを含む合計 8 ビットを使用すると、$(-2^7) / 2^4 = -8$ から $(2^7 - 1) / 2^4 = 7.9375$ までの $2^8$ 値を、$1/2^4 = 0.0625$ の増分で等間隔に表すことができます。仮定します

ここで、表現可能な範囲はゼロに関して対称です。
固定小数点量子化スキームでは、ある固定範囲内の離散的な値のセットしか表現できないため、実際の値を表現しようとすると近似誤差が発生します。リソース効率の高い機械学習の焦点の 1 つは、安定したトレーニングと推論を可能にするために、この近似誤差、つまり量子化誤差を最小限に抑えることです。
FPGA は、主にルックアップ テーブル (LUT) を通じてデジタル ロジックを実装します。LUT は、バイナリ入力の組み合わせごとに出力を保存することで任意のバイナリ関数を表す小さなコンポーネントです。たとえば、 $\text{AND} : \{0, 1\}^2 \to \{0, 1\}$ 1 はルックアップ テーブルで表されます。
したがって、ルックアップ テーブルとして表されるこれらの 2 値関数をニューラル ネットワークのコア プリミティブとして学習することは理にかなっています。このようなネットワークはルックアップ テーブル ニューラル ネットワーク (LUT-NN) として知られています。ただし、勾配降下法または同様のアプローチを通じてルックアップ テーブルを学習するのは困難です。
この問題に対処するには、勾配降下法を通じて実数値​​関数 $f: \mathbb R \to \mathbb R$ を学習できることを思い出してください。 $b_i$ 入力ビットと $b_o$ 出力ビットで固定小数点量子化を実行すると、$f$ は二項関数 $f_Q : \{0, 1\}^{b_i} \to \{0, 1\}^{b_o}$ になります。その後、連続 $f$ を学習し、量子化して目的のルックアップ テーブルを取得できます。
$f$ を LUT に変換するには、関数ドメインと範囲を $N_i = 2^{b_i}$ および $N_o = 2^{b_o}$ 値に離散化します。 $f_Q$ のルックアップ テーブルには、入力値 $I \in \{I_0, I_1, \ldots, I_{N_i-1}\}$ ごとに、対応する出力が格納されます。
連続 $f$ (破線) を 2 値関数 $f_Q$ (オレンジ色の点) に量子化します。
上記の関数 $f$ の例では、$q_{l-1}$ と $q_{l}$ が入力と出力の量子化を表し、バイナリ関数を生成します。

LUT を使用したイオン $f_Q$:
このアプローチを拡張して、多変量関数をルックアップ テーブルとして表すこともできます。$f_m: \mathbb{R}^{d_i} \to \mathbb{R}^{d_o}$ は、$d_i b_i$ 入力ビットと $d_o b_o$ 出力ビットを持つ 2 項関数に変わります。
このルックアップ テーブルには、 $2^{d_i b_i}$ 可能な入力の組み合わせごとに、サイズ $d_o b_o$ のエントリが格納されます。
これは、合理的な $d_i$ にとっては明らかに非現実的です。したがって、LUT ベースのネットワークは、より小さなルックアップ テーブルと算術演算を組み合わせて、表現力が高く、リソース効率が高く、トレーニングが容易なアーキテクチャを実現します。
KAN は、MLP (多層パーセプトロン) アーキテクチャの学習可能な重みと固定活性化関数を学習可能な活性化関数に置き換えます。この研究では、KAN が効率的で表現力豊かな LUT-NN のための自然なアーキテクチャであることを実証します。
KAN 層では、各エッジはスカラー重みの代わりに学習可能な一変量関数を保持します。 $n_{\mathrm{in}}$ 入力と $n_{\mathrm{out}}$ 出力を持つ KAN 層の場合、$q$ 番目の出力のアクティブ化は次のようになります。
ここで $\phi_{q,p} \colon \mathbb{R} \to \mathbb{R}$ は学習されたエッジの活性化です。を使用する MLP と比較すると、
$\sigma$ を固定すると、KAN はエッジ関数 $\{ \phi_{q,p} \}$ に非線形性を配置し、ノード演算を単純な合計として維持します。
次の問題は、KAN 活性化 $\{ \phi_{q,p} \}$ をどのように学習するかです。そのために、それらを何らかの機能基盤の線形結合としてパラメータ化します。
これにより、係数 $c_{q,p,i}$ を勾配降下法のトレーニング可能なパラメーターとして扱うことができます。元の KAN 論文では B スプラインを使用しています。これは、滑らかで局所的な多項式基底を形成します。つまり、基底関数のサブセットのみが任意の入力に対して非ゼロになります。さらに、B スプライン、つまりアクティベーション $\{ \p

hi_{q,p} \}$ は、小さな有限のドメイン ($[-1, 1]$ など) にわたって定義されており、これが重要であることがわかります。
用語: B スプラインは基底関数 ${B_i}$ を指しますが、活性化は学習された $\phi_{q,p}(x) = \sum_{i=1}^n c_{q,p,i}B_i(x)$ を指します。
KAN アーキテクチャの完全な動作はまだ十分に調査されていませんが、スケーリング則、パラメータの効率、および解釈可能性において MLP よりも改善される可能性があります。超高速機械学習の場合、最初の 2 つの特性が特に関連します。
最初の論文の重要なアイデアは、トレーニング可能な LUT-NN を構築する原則的な方法として KAN を使用することです。 LUT-NN の背景セクションで説明した手順を使用して、ネットワーク内の各活性化関数を個別の LUT で表します。ここで、KAN アクティベーションが LUT 表現に特に適していることを示します。
他の多くの LUT-NN スキームは多変量関数をルックアップ テーブルとして表しますが、これは非効率です。前に見たように、LUT エントリの数は入力次元に応じて指数関数的に増加します。対照的に、KAN の中核となる特性は、一変量のアクティベーションを合計することで、指数関数的なスケーリングを回避し、単純なプルーニング (つまり、重要でないネットワーク コンポーネントを削除してリソース使用量を削減すること) を可能にすることです。 2 さらに、各アクティベーションは小さな有限領域 ($[-1,1]$ など) にわたって定義されるため、関数を量子化するときに入力範囲全体をカバーできます。
ここでは、標準的な手法を使用してソフトウェア (CPU/GPU 上の PyTorch など) で KAN をトレーニングし、FPGA 上で推論用の固定モデルをデプロイします。
トレーニングされた KAN 層に対して FPGA 上で推論を実行するには、固定小数点量子化スキームを採用し、ルックアップ テーブルを使用してアクティベーション $\{ \phi_{q,p} \}$ を並列計算します。次に、合計 $\sum_{p=1}^{n_\text{in}} \phi_{q,p}(x

_p)$ $n_\text{in} を使用 - 1$ ペアごとの加算 (加算器ツリーに配置)。
B スプライン ${B_i}$ 自体ではなく、完全なアクティベーション $\phi_{q,p}(x_p)$ をルックアップ テーブルに変換します。これは、モデルが事前トレーニングされており、アクティベーションが推論時に固定されるためです。
多層ネットワークの場合、各層に対して上記の回路を構築し（その LUT にはそれぞれの学習されたアクティベーションが保存されます）、各層の出力を次の層の入力に接続します。
FPGA 上で効率的な LUT ベースの KAN 推論を行うためのアーキテクチャの概要。
このフレームワークは、レイテンシやリソース使用量などのメトリクスに関して最先端のニューラル ネットワーク FPGA アクセラレータと同等かそれを上回り、以前の KAN-FPGA 実装と比べて 2700 倍高速化されています。興味のある方は詳しくは紙面をチェックしてみてください！
ソフトウェアでモデルをトレーニングして FPGA にデプロイすると、非常に高速な推論が得られますが、モデル自体はデプロイ後も固定されたままです。ただし、多くのリアルタイム設定では、モデル化されるシステムは静的ではありません。その状態やプロパティは高周波で変化する可能性があります。したがって、量子制御や核融合などのアプリケーションでは、モデルは超高速推論を実行しながら、数マイクロ秒以内にその動作を適応させる必要がある場合があります。
これがオンライン学習の動機です。FPGA を事前トレーニング済みモデルの推論専用デバイスとして扱うのではなく、新しいデータが到着するとリアルタイムでモデルを更新します。入力をストリーミングし、モデルを実行し、各予測をフィードバックまたは目標値と比較し、その誤差を使用してモデル パラメーターを更新します。つまり、順方向パスのみではなく、順方向パス、逆方向パス、および勾配更新はすべて FPGA 自体で実行されます。
メモリから重みをフェッチする CPU/GPU とは異なり、計算

勾配を取得して書き戻す場合、FPGA は、係数を格納している FPGA メモリを直接変更する専用の並列回路として勾配更新ロジックを実装します。
FPGA でのリアルタイムの勾配更新の概念は十分に検討されておらず、歴史的には非現実的であると考えられてきましたが、LUT ベースの KAN 推論のアイデアを拡張して、サブマイクロ秒のタイムスケールでこの形式のオンライン学習を可能にすることができることを示します。 3
静的な事前トレーニングされたモデルで推論を実行するだけでなく、FPGA 上でモデルをトレーニングしたいため、学習された活性化 $\{ \phi_{q,p} \}$ ではなく、B スプライン基底関数 $\{ B_i \}$ を LUT に保存します。
その理由は、次の係数 $c_{q,p,i}$ が
FPGA 上でトレーニングが進行するにつれて更新されます。アクティベーション $\{ \phi_{q,p} \}$ は変化し続けるため、それらを事前計算して保存することはできません。
代わりに、B スプライン値を検索し、各アクティベーションの係数を乗算して計算する必要があります。
ここで、B スプライン基底 $\{ B_i \}$ の特定の特性を実証します。これにより、固定小数点量子化下で勾配更新がスパースかつ安定になり、以前と比較して非常に低いレイテンシと小さなリソース フットプリントで勾配ベースの学習が可能になります。

[切り捨てられた]

## Original Extract

https://web.archive.org/web/20260609200156/https://aarushgup...

Ultrafast machine learning on FPGAs via Kolmogorov-Arnold Networks
This post is a high-level explainer for my Master’s thesis, which involves designing hardware architectures for ultrafast inference and online learning using the Kolmogorov-Arnold Network (KAN) architecture. I’ll assume familiarity with standard machine learning concepts, as well as some understanding of hardware and digital circuits; read my previous post here for the latter.
Please read the two papers below for more information, particularly for details on benchmarks and notable results.
[FPGA 2026 Best Paper]
Duc Hoang * , Aarush Gupta * , and Philip C. Harris. “KANELÉ: Kolmogorov–Arnold Networks for Efficient LUT-based Evaluation.” Proceedings of the 2026 ACM/SIGDA International Symposium on Field Programmable Gate Arrays . ACM, 2026. https://dx.doi.org/10.1145/ 3748173.3779202
[ICML 2026]
Duc Hoang * , Aarush Gupta * , and Philip Harris. “Ultrafast on-FPGA Online Learning via Spline Locality in Kolmogorov-Arnold Networks.” arXiv preprint arXiv:2602.02056 , 2026. https://arxiv.org/abs/2602.02056
The case for machine learning on FPGAs
Most modern machine learning workloads, whether training or inference, run on graphics processing units (GPUs). Through hardware architectures that support a highly parallel execution model, GPUs can perform simple operations on large amounts of data with extremely high throughput. This makes them ideal for machine learning problems involving large architectures or batch-style training and inference.
However, complex GPU architectures cannot meet the demands of applications that require ultra-low latency (e.g. sub-microsecond latency) and high hardware efficiency. Processors (e.g. CPUs and GPUs) incur significant overhead from scheduling and optimizing instructions, dynamically accessing memory, and so on. Extremely specialized workloads with ultralow latency (e.g. $\sim$nanoseconds) and efficiency requirements are instead better served by custom hardware accelerators .
Field-programmable gate arrays, or FPGAs, are reconfigurable digital logic devices that are extremely well-suited for this style of custom hardware acceleration. FPGAs contain lookup tables (LUTs), which represent digital functions by enumerating the output value for every combination of binary inputs; flip-flops (FFs), which store state; and other memory and computation primitives. These components and the connections between them are reconfigured to design a custom digital circuit, allowing for low-level hardware architecture and algorithm co-design that enables ultrafast machine learning. Importantly, neural networks are implemented directly as digital logic , rather than as instructions that are sequentially executed on a processor.
FPGAs and other digital devices fundamentally operate on bits rather than continuous values. However, we often think about arithmetic operations in neural networks (e.g. $\times, +$) as happening over the real numbers $\mathbb R$. We thus need to encode real numbers as bitstrings (sequences of bits), a process known as quantization . Operations like addition and multiplication then become binary functions.
One method for doing this is fixed-point quantization .
Fixed-point quantization represents numbers in base-2, where some bits (fractional bits) come after the decimal point. To illustrate, if we use 8 bits total with 4 fractional bits after the decimal point, we can represent $2^8$ values from $(-2^7) / 2^4 = -8$ to $(2^7 - 1) / 2^4 = 7.9375$, spaced evenly in increments of $1/2^4 = 0.0625$. We will assume here that the representable range is symmetric about zero.
In a fixed-point quantization scheme, we can only represent a discrete set of values in some fixed range, which will lead to approximation error when trying to represent real values. One focus of resource-efficient machine learning is minimizing this approximation error, or quantization error , to enable stable training and inference.
FPGAs implement digital logic primarily through lookup tables (LUTs), which are small components that represent arbitrary binary functions by storing their output for each combination of binary inputs. For example, $\text{AND} : \{0, 1\}^2 \to \{0, 1\}$ 1 is represented with a lookup table
It then makes sense to learn these binary functions, represented as lookup tables, as core primitives of a neural network: such a network is known as a lookup-table neural network (LUT-NN). However, learning lookup tables through gradient descent or similar approaches is difficult.
To address this issue, recall that we can learn real-valued functions $f: \mathbb R \to \mathbb R$ through gradient descent. If we perform fixed-point quantization with $b_i$ input bits and $b_o$ output bits, $f$ becomes a binary function $f_Q : \{0, 1\}^{b_i} \to \{0, 1\}^{b_o}$. We can then learn continuous $f$ and quantize to get our desired lookup tables!
To convert $f$ into a LUT, we discretize the function domain and range into $N_i = 2^{b_i}$ and $N_o = 2^{b_o}$ values. The lookup table of $f_Q$ stores, for each of the input values $I \in \{I_0, I_1, \ldots, I_{N_i-1}\}$, the corresponding output.
Quantizing a continuous $f$ (dashed line) to a binary function $f_Q$ (orange dots).
The example function $f$ above, where $q_{l-1}$ and $q_{l}$ represent quantization of the inputs and outputs, produces a binary function $f_Q$ with LUT:
We could also extend this approach to represent multivariate functions as lookup tables, where $f_m: \mathbb{R}^{d_i} \to \mathbb{R}^{d_o}$ would turn into a binary function with $d_i b_i$ input bits and $d_o b_o$ output bits.
In this lookup table, we store entries of size $d_o b_o$ for each of the $2^{d_i b_i}$ possible combinations of inputs.
This is clearly impractical for any reasonable $d_i$. LUT-based networks therefore combine smaller lookup tables with arithmetic operations to enable architectures that are expressive, resource efficient, and easily trainable.
KANs replace the learnable weights and fixed activation functions in MLP (multi-layer perceptron) architectures with learnable activation functions . In this work, we demonstrate that KANs are a natural architecture for efficient, expressive LUT-NNs.
In a KAN layer, each edge carries a learnable univariate function instead of a scalar weight. For a KAN layer with $n_{\mathrm{in}}$ inputs and $n_{\mathrm{out}}$ outputs, the activation of the $q$-th output is
where $\phi_{q,p} \colon \mathbb{R} \to \mathbb{R}$ are the learned edge activations. Compared to an MLP, which uses
with fixed $\sigma$, the KAN places the nonlinearity in the edge functions $\{ \phi_{q,p} \}$ and keeps the node operation as a simple sum.
The next question is how to learn the KAN activations $\{ \phi_{q,p} \}$. To do so, we parametrize them as linear combinations of some functional basis:
which allows us to treat the coefficients $c_{q,p,i}$ as trainable parameters for gradient descent. The original KAN paper uses B-splines , which form a polynomial basis that is smooth and also local , i.e. only a subset of basis functions are nonzero for any given input. Additionally, B-splines, and therefore the activations $\{ \phi_{q,p} \}$, are defined over a small, finite domain (e.g. $[-1, 1]$), which turns out to be important.
Terminology: B-splines refer to the basis functions ${B_i}$, whereas activations refer to the learned $\phi_{q,p}(x) = \sum_{i=1}^n c_{q,p,i}B_i(x)$.
While the complete behavior of the KAN architecture has yet to be fully explored, it offers potential improvements over MLPs in scaling laws, parameter efficiency, and interpretability. For ultrafast machine learning, the first two characteristics are especially relevant.
The key idea of our first paper is to use KANs as a principled way to build trainable LUT-NNs. We represent each activation function in the network with a separate LUT, using the procedure discussed in the LUT-NN background section. We now demonstrate that KAN activations are particularly well-suited for LUT representation.
Many other LUT-NN schemes represent multivariate functions as lookup tables, which is inefficient: as we saw earlier, the number of LUT entries scales exponentially with input dimension. In contrast, the core property of KANs is that they sum univariate activations, avoiding exponential scaling and enabling straightforward pruning (i.e. removing unimportant network components to reduce resource usage). 2 Additionally, because each activation is defined over a small finite domain (e.g. $[-1,1]$), we can cover the entire input range when quantizing the function!
Here, we train KANs in software (e.g. PyTorch on CPUs/GPUs) using standard techniques and then deploy fixed models for inference on FPGAs.
To perform on-FPGA inference for a trained KAN layer, we adopt a fixed-point quantization scheme and compute the activations $\{ \phi_{q,p} \}$ in parallel using lookup tables. Then, we perform the sum $\sum_{p=1}^{n_\text{in}} \phi_{q,p}(x_p)$ using $n_\text{in} - 1$ pairwise additions (arranged in an adder tree ).
We convert each complete activation $\phi_{q,p}(x_p)$ to a lookup table, not the B-splines ${B_i}$ themselves. This is because the model is pretrained and activations are fixed at inference time.
For multi-layer networks, we construct the circuits described above for each layer (whose LUTs store their respective learned activations) and connect each layer’s outputs to the next layer’s inputs.
Overview of the architecture for efficient LUT-based KAN inference on FPGAs.
This framework matches and surpasses state-of-the-art neural network FPGA accelerators on metrics including latency and resource usage, with a 2700x speedup over prior KAN-FPGA implementations. Check out the paper for more details if you’re interested!
Training a model in software and deploying it to an FPGA provides extremely fast inference, but the model itself is still fixed after deployment. In many real-time settings, however, the system being modeled is not static: its state or properties can evolve at high frequencies. In applications such as quantum control and nuclear fusion, a model may therefore need to adapt its behavior within a fraction of a microsecond while still performing ultrafast inference.
This is the motivation for online learning : instead of treating the FPGA as an inference-only device for pretrained models, we also update the model in real time as new data arrives. We stream in inputs, run the model, compare each prediction against feedback or a target value, and use that error to update the model parameters. In other words, the forward pass, backward pass, and gradient update all run on the FPGA itself, rather than only the forward pass.
Unlike a CPU/GPU, which fetches weights from memory, calculates gradients, and writes them back, the FPGA implements the gradient update logic as a dedicated, parallel circuit that directly modifies the FPGA memory storing the coefficients.
Although the concept of real-time gradient updates on FPGAs has been underexplored and historically considered impractical, we demonstrate that the ideas of LUT-based KAN inference can be extended to enable this form of online learning at sub-microsecond timescales . 3
Because we now want to train models on the FPGA rather than only run inference with static, pre-trained ones, we store the B-spline basis functions $\{ B_i \}$ in LUTs rather than the learned activations $\{ \phi_{q,p} \}$.
The reason is that the coefficients $c_{q,p,i}$ in
are updated as training progresses on-FPGA: since the activations $\{ \phi_{q,p} \}$ keep changing, we cannot precompute and store them .
Instead, we must look up B-spline values and multiply them by each activation’s coefficients to compute it.
We now demonstrate certain properties of the B-spline basis $\{ B_i \}$ that make gradient updates both sparse and stable under fixed-point quantization, enabling gradient-based learning with extremely low latency and a small resource footprint compared to prior

[truncated]
