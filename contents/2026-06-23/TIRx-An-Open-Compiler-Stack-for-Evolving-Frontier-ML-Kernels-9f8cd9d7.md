---
source: "https://tvm.apache.org/2026/06/22/tirx"
hn_url: "https://news.ycombinator.com/item?id=48640473"
title: "TIRx: An Open Compiler Stack for Evolving Frontier ML Kernels"
article_title: "TIRx: An Open Compiler Stack for Evolving Frontier ML Kernels"
author: "matt_d"
captured_at: "2026-06-23T05:18:53Z"
capture_tool: "hn-digest"
hn_id: 48640473
score: 1
comments: 0
posted_at: "2026-06-23T04:52:39Z"
tags:
  - hacker-news
  - translated
---

# TIRx: An Open Compiler Stack for Evolving Frontier ML Kernels

- HN: [48640473](https://news.ycombinator.com/item?id=48640473)
- Source: [tvm.apache.org](https://tvm.apache.org/2026/06/22/tirx)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T04:52:39Z

## Translation

タイトル: TIRx: 進化するフロンティア ML カーネルのためのオープン コンパイラ スタック

記事本文:
TIRx: 進化するフロンティア ML カーネルのためのオープン コンパイラ スタック
2026 年 6 月 22 日
•
Apache TVM コミュニティ
本日は、Apache TVM 上に構築されたオープンソースのハードウェアネイティブ DSL および ML カーネル用コンパイラである TIRx を紹介します。これは、高速で移動するカーネルと高速で移動するハードウェアが出会う AI ソフトウェア スタックの部分をターゲットとしています。TIRx は、今日の GPU と特殊な AI アクセラレータにコンパイルされ、次の世代に合わせて成長するように設計されています。同じ設計が、専門家によって作成されたカーネル、エージェントによって生成されたカーネル、およびメガカーネル システムに提供されます。
私たちはより広範なコミュニティと協力して、発売時に次の資料を提供してきました。
PyPI ホイールと Python フロントエンド。 @T.jit / @T.prim_func スタイルのオーサリング、パーサー ユーティリティ、TIRx プログラムを構築するための Python API を備えた Python 埋め込みハードウェア ネイティブ カーネル DSL。
TIRx カーネル ライブラリとベンチマーク。 GEMM、アテンション スタイル カーネル、Blackwell GPU 上の低精度演算子をカバーするエンドツーエンドの例。
最新の GPU プログラミングに関するオープン コース。この厳選されたオンライン コースは、カーネギー メロン大学の機械学習システム コースの一部として教えられ、TIRx を使用して学生に機械学習システム用の最新の GPU プログラミングを教えます。
次のリソースを見つけることができます。
GitHub: https://github.com/apache/tvm
ドキュメント: https://tvm.apache.org/docs/tirx/overview.html
PyPI ホイール: https://pypi.org/project/apache-tvm/0.25.0/
pip インストール apache-tvm == 0.25.0
コミュニティ TIRx カーネル ライブラリ: https://github.com/mlc-ai/tirx-kernels
機械学習システムのための最新の GPU プログラミング: https://mlc.ai/modern-gpu-programming-for-mlsys/index.html
カーネル DSL は、プログラマとマシンの間の適切な境界を選択したときに最も効果的です。成熟したカーネルと成熟したハードウェアの場合、その境界は高くなる可能性があります。

レベル: コンパイラは、スレッドの割り当て、メモリの移動、レイアウトの詳細、および命令の選択をコンパクトなテンソルまたはタイルの抽象化の背後に隠します。 Triton は標準的な例であり、その採用は、これが確立されたカーネル パターンに対していかにうまく機能するかを示しています。フロンティアでは、同じ境界線がより大きな圧力にさらされています。新しい命令、メモリ空間、連携パターン、カーネル アルゴリズムは、それらを適切に自動化するための組み込み機構がコンパイラに組み込まれる前に現れることがよくあります。そうなると、高レベルのコンパイラが通常隠す部分は、まさに専門家が手動で制御する必要がある部分になります。
TIRx (「ティアエクス」と発音) は、以下の 3 つの決定を中心に、より低く、より明示的な境界を選択することで対応します。
オーケストレーションはハードウェアネイティブのソースに残ります。パイプラインの構造、同期、役割の割り当て、メモリの配置、およびバックエンドの組み込みは、フロンティアでの専門家の制御が最も必要となる部分であるため、TIRx はこれらを、まだ新機能をモデル化していない抽象化の背後ではなく、ソース内に保持します。
繰り返し発生するタイル プリミティブはコンパイラに公開されます。実行スコープ、テンソル レイアウト、およびタイル プリミティブ ディスパッチにより、カーネル全体に固定コンパイラ パイプラインを強制することなく、一般的な操作を再利用可能、分析可能、およびバックエンド間で移植可能な状態に保つことができます。ハードウェアネイティブ制御のコストはエンジニアリングの労力です。カーネルとバックエンドごとにすべての操作を手作業で記述するのは骨の折れる作業です。繰り返しの操作をタイル プリミティブとして公開すると、この問題が軽減されるため、作成者は、同じデータの移動や行列の乗算を毎回書き直すのではなく、ディスパッチされた実装を再利用します。
新しいハードウェアは最初に組み込み関数として入力され、その後タイル プリミティブが入力されます。新しい機能は、ネイティブ組み込み関数 (単一のバックエンド固有の薄いラッパー) としてすぐに使用できます。

ハードウェア操作。使用パターンがカーネル間で安定したら、タイル プリミティブ (スコープ、オペランド、バックエンド全体でディスパッチされるレイアウト認識操作) に昇格できます。コアの抽象化は小さいままであり、新しい機能の組み込みを追加しても既存の機能が破壊されることはありません。
その結果、ハードウェアに合わせて拡張できる DSL とコンパイラ スタックが得られます。これは、TIRx の背後にある中心的な設計哲学です。つまり、基盤を小規模かつ明示的に保ち、新しいアクセラレータ世代の到来に合わせてバックエンド ライブラリを進化させます。
これにより、TIRx は TileLang などのシステムの下に配置され、レイアウト推論とスレッド バインディングはコンパイラーに残したまま、メモリ スコープとパイプライン処理を公開することで Triton に比べて境界が低くなります。 TIRx は、これらの高レベルの懸念事項を意図的にコアの外に残し、そのようなシステムが構築できる最小限の基盤を提供します。私たちは TileLang コミュニティと協力して、TileLang コンパイルをサポートする新しい最小限の基盤として TIRx を導入しています。
同じ小規模で明示的な基盤により、エンジニアリングの労力を可能な限り削減しながら最高のパフォーマンスを追求するさまざまな種類のユーザーに 1 つの設計でサービスを提供できます。専門家が作成した実稼働カーネル、エージェント生成カーネル、メガカーネル システムには、それぞれネイティブ レベルでの制御と、コンパイラーが認識できる反復操作の両方が必要です。
この投稿の残りの部分では、プログラミング モデルを説明し、次にこれらの各方向を順番に説明します。
実際の境界は次のようになります。 TIRx プログラムは、構造化されたネイティブ カーネルとして読み取ります。ループ、分岐、テンソル、同期、パイプライン状態、およびバックエンド組み込み関数が直接書き込まれます。タイル プリミティブは、定期的なハードウェア操作を再利用可能およびディスパッチ可能にする必要がある場所に表示されます。 3つの成分がほとんどの成分を含んでいます

モデル。
実行スコープは、誰がどのような粒度で操作を実行するかを決定します。これを選択するのは 2 つです。領域に入るハードウェアの役割を選択する制御フローと、呼び出しの粒度を設定するプリミティブ名前空間です。
修飾されていない Tx.* 呼び出しはスレッド レベルで実行されます。 Tx.wg.* はワープグループ レベルで実行されます。 T.ptx.elect_sync() などの述語を使用すると、スレッド レベルの呼び出しをさらに絞り込んで、単一の発行スレッドに絞り込むことができます。
テンソル レイアウトは、ストレージファースト インターフェイスを通じて論理テンソルが存在する場所を記述します。タイルは、グローバル メモリ、共有メモリ、レジスタ、テンソル メモリ、またはアクセラレータ SRAM に存在する場合があります。ユーザーは、各タイルがどこに存在するか、およびその要素がレーン、ワープ、レジスター全体にどのように分散されるかを宣言します。その宣言はタイルに添付されたままになります。プリミティブが呼び出されると、コンパイラーはそれらの宣言を読み取って実装を選択します。レイアウトはストレージの説明であり、ループ変換ユーティリティではありません。ユーザーはタイルのレイアウトを構築できますが、ループを変換するためにレイアウトを使用することはありません。
タイル プリミティブ ディスパッチは、1 つの呼び出しをネイティブ IR に変換します。オペランド レイアウト、実行スコープ、ターゲット、または明示的なdispatch= ヒントから、一致する実装が選択されます。グローバルから共有へのコピーは TMA に解決され、共有は ldmatrix に登録され、テンソル メモリは tcgen05.ld に登録されます。行列乗算は WGMMA、tcgen05、またはシストリック アレイ命令に解決されます。次に、ディスパッチは、その命令をタイル全体に適用するために必要なループとアドレス指定を生成します。
これらの成分は、範囲が重要な場合にはどこでも組み合わせられます。以下の GEMM エピローグでは、ワープグループ スコープのプリミティブとスレッド スコープのプリミティブが同じ領域にあります。 Tx.wg.* 呼び出しは、ワープグループ全体にタイルを移動およびキャストしますが、最後のスレッド スコープの Tx.copy_async は、明示的な発行スレッドによって保護されています。

述語、TMA ストアを実行します。
上記の抜粋は簡略化されています。全体像として、完全な FP16/BF16 GEMM カーネルの 2 つの役割、つまり TMA プロデューサーとテンソル メモリ ライトバックを示します。一行ずつ読む必要はありません。重要なのは、オーケストレーションに関係するすべての要素 (パイプラインの状態、バリア プロトコル、ロールの選択、 tcgen05.wait や cp_async.bulk などの低レベルの同期組み込み関数) は通常のソース コードにとどまり、一方、定期的なデータの移動は、スコープ、レイアウト、およびディスパッチ構成から選択される低下がタイル プリミティブとして表示されるということです。
3 つの要素のうち、レイアウトはデザイン上の決定に最も関係するため、詳しく検討する価値があります。
Tensor レイアウト用のストレージファーストインターフェイス
TIRx は、レイアウトをテンソル ストレージのファーストクラス表現として扱います。 CuTe に精通している読者ならこの領域を認識できるでしょう。どちらのシステムもレイアウトを使用してテンソル データがハードウェア リソースにどのようにマッピングされるかを記述しますが、CuTe はタイル作業がスレッド間でどのように分割されるかを導き出すためのプログラマブル インターフェイスとしてレイアウトを公開するのに対し、TIRx はプリミティブ ディスパッチによって消費されるストレージ コントラクトとしてレイアウトを使用します。
TIRx レイアウトは、論理テンソル インデックスを名前付き軸上の物理座標にマップします。このモデルは、セマンティック ハードウェア軸にストライドを付加し、明示的な shard 、 レプリカ 、および offset コンポーネントを追加することによって、シェイプ ストライド レイアウトを一般化します。シャードは、論理要素が物理軸間でどのように分割されるかを表します。レプリカは、同じ論理要素が複製される場所を示します。オフセットは、物理的な配置が始まる場所を表します。具体的には、
D（シャード）。 1 つ以上のイテレータのリスト。それぞれのイテレータには、ある軸上の範囲とストライドが含まれます。 D は、論理インデックスをこれらの反復全体に分割し、基本座標を生成します。これにより、シェイプ ストライドが複数の軸に一般化されます。
R（レプリカ）

）。論理インデックスとは無関係に、ハードウェア空間内のオフセットを列挙するレプリケーション反復子のセット。このセットの各要素を D の結果に追加すると、レプリケーションまたはブロードキャストが生成されます。
O (オフセット)。固定座標オフセット (軸ごとに 1 つの整数) がすべての結果に追加されます。これにより、データが特定の基本位置に配置されるか、排他的なリソースが予約されます。
TIRx レイアウト Python API の具体的な例は次のとおりです。
これは、レーンとワープに分散され、別のワープグループに複製され、ワープ軸上のオフセットに配置された論理タイルを表します。 (8, 16) 形状空間内の論理座標 (i, j) を指定すると、計算によってそれぞれワープ軸、レーン軸、およびレグ軸にマッピングされます。
たとえば、論理 (3, 9) の要素 57 は次のようにマッピングされます。
基本位置: 6@warpid、12@laneid、1@m
所有者 (レプリカ経由×2): { warpid=6 LANEID=12 }、 { warpid=10 LANEID=12 }
(対話型デモを開いて要素 57 をクリックすると、これらの所有者を正確に確認できます。)
▶ インタラクティブなレイアウトのデモを開く ↗
TIRx のレイアウト インターフェイスは、4 つの設計選択肢を中心に構築されています。
1. レイアウトはストレージ契約であり、作業分割インターフェイスではありません。
CuTe では、レイアウトはデータの配置を表すだけではありません。これは、タイル操作がスレッド間でどのように分散されるかを導き出すためのプログラミング インターフェイスの一部でもあります。ユーザーは、レイアウトを構成、タイル表示、および分割して、コピーおよびコンピューティング操作のデータと作業の分散を表現します。 TIRx では、境界の描画方法が異なります。ユーザーは各タイルのストレージ レイアウトを記述し、それらのタイルに対してタイル プリミティブを呼び出します。レイアウトは、シャーディング、レプリケーション、オフセットなど、論理テンソル座標が物理ハードウェア座標にどのようにマッピングされるかを記録します。これは、実行パーティショニングを構築するために使用されるサーフェスではありません。プリミティブが下げられると、ディスパッチはオペランド レイアウト、実行スコープ、およびバーを使用します。

ckend ターゲットを使用して、スレッドの分割、ループのネスト、アドレス指定、および命令シーケンスを生成します。この意味で、TIRx レイアウトはストレージを正確に表現する必要があるだけです。変換ロジックは、ユーザーが作成したレイアウト構成ではなく、プリミティブ ディスパッチ内に存在します。
2. レイアウトは、論理テンソル座標を物理ハードウェア座標にマップします。
明示的なレプリカとオフセット構造は、指定された論理から物理への定式化から得られます。レイアウトを形式化するもう 1 つの方法は、物理的位置を論理座標にマッピングすることです。これにより、レプリケーション (複数の物理的位置に格納される 1 つの論理要素) を点値関数として定義できます。ただし、ストライド パターンで物理的位置にまたがるテンソルの場合、一部の物理的位置には明確に定義されたマッピングがない可能性があります。
3. レイアウトは一般的な形状をサポートします。
最新のカーネルは、2 のべき乗のみの表現に適合しない形状を頻繁に使用します。グローバル テンソル、マルチステージ共有メモリ バッファ、テンソル メモリ タイル、アクセラレータ スクラッチパッド、および分散テンソルはすべて、実際には一般的な形状を生成します。したがって、TIRx レイアウトは、特殊なケースとして扱うのではなく、一般的な形状のサポートから始まります。これはブロックスケールの GEMM スケールファクターにとって重要です

[切り捨てられた]

## Original Extract

TIRx: An Open Compiler Stack for Evolving Frontier ML Kernels
Jun 22, 2026
•
Apache TVM Community
Today we are introducing TIRx , an open-source, hardware-native DSL and compiler for ML kernels, built on Apache TVM. It targets the part of the AI software stack where fast-moving kernels meet fast-moving hardware: TIRx compiles to GPUs and specialized AI accelerators today and is designed to grow with the generations that follow. The same design serves expert-written kernels, agent-generated kernels, and megakernel systems.
We have been working together with the broader community to provide the following materials at launch:
PyPI wheel and Python frontend. A Python-embedded hardware-native kernel DSL with @T.jit / @T.prim_func style authoring, parser utilities, and Python APIs for constructing TIRx programs.
TIRx kernel library and benchmarks. End-to-end examples covering GEMM, attention-style kernels, and low-precision operators on Blackwell GPUs.
Open course on modern GPU programming. This curated online course was taught as part of the machine learning systems course at Carnegie Mellon University, and uses TIRx to teach students modern GPU programming for machine learning systems .
You can find the following resources:
GitHub: https://github.com/apache/tvm
Documentation: https://tvm.apache.org/docs/tirx/overview.html
PyPI wheel: https://pypi.org/project/apache-tvm/0.25.0/
pip install apache-tvm == 0.25.0
Community TIRx kernel library: https://github.com/mlc-ai/tirx-kernels
Modern GPU programming for machine learning systems: https://mlc.ai/modern-gpu-programming-for-mlsys/index.html
Kernel DSLs are most effective when they choose the right boundary between the programmer and the machine. For mature kernels and mature hardware, that boundary can be high-level: the compiler hides thread assignment, memory movement, layout details, and instruction selection behind compact tensor or tile abstractions. Triton is the canonical example, and its adoption shows how well this works for established kernel patterns. At the frontier, the same boundary is under more pressure. New instructions, memory spaces, cooperation patterns, and kernel algorithms often appear before a compiler has the built-in machinery to automate them well. When that happens, the parts a high-level compiler would normally hide are exactly the parts an expert still needs to control by hand.
TIRx (pronounced “tier-ex”) responds by choosing a lower and more explicit boundary, organized around three decisions:
Orchestration stays in the hardware-native source. Pipeline structure, synchronization, role assignment, memory placement, and backend intrinsics are the parts that most often need expert control at the frontier, so TIRx keeps them in source rather than behind an abstraction that may not yet model a new feature.
Recurring tile primitives are exposed to the compiler. Execution scope, tensor layout, and tile primitive dispatch let common operations stay reusable, analyzable, and portable across backends, without forcing the whole kernel through a fixed compiler pipeline. The cost of hardware-native control is engineering effort: writing every operation by hand for each kernel and backend is laborious. Exposing recurring operations as tile primitives alleviates this, so authors reuse a dispatched implementation instead of re-writing the same data movement or matrix multiply each time.
New hardware enters as intrinsics first, tile primitives later. A new feature can be used immediately as a native intrinsic — a thin, backend-specific wrapper over a single hardware operation. Once the usage pattern stabilizes across kernels, it can be promoted to a tile primitive: a layout-aware operation that dispatches across scopes, operands, and backends. The core abstraction stays small, and adding an intrinsic for a new feature never breaks existing ones.
The result is a DSL and compiler stack that can grow with the hardware. This is the core design philosophy behind TIRx: keep the foundation small and explicit, and let the backend library evolve as new accelerator generations arrive.
This places TIRx below systems like TileLang, which also lowers the boundary relative to Triton by exposing memory scopes and pipelining, while still leaving layout inference and thread binding to the compiler. TIRx deliberately leaves those higher-level concerns outside its core and provides a minimal foundation that such systems can build on; we are working with the TileLang community to bring TIRx as a new minimal foundation to support TileLang compilation.
The same small, explicit foundation is what lets one design serve several kinds of users who pursue peak performance while reducing engineering effort as much as possible: expert-written production kernels, agent-generated kernels, and megakernel systems, each of which needs both control at the native level and recurring operations the compiler can see.
The rest of this post walks through the programming model and then through each of these directions in turn.
Here is what that boundary looks like in practice. A TIRx program reads as a structured native kernel: loops, branches, tensors, synchronization, pipeline state, and backend intrinsics are written directly. Tile primitives appear where a recurring hardware operation should become reusable and dispatchable. Three ingredients carry most of the model.
Execution scope decides who runs an operation and at what granularity. Two things select it: control flow, which picks the hardware role entering a region, and the primitive namespace, which sets the granularity of the call.
An unqualified Tx.* call runs at thread level; Tx.wg.* runs at warpgroup level. A predicate such as T.ptx.elect_sync() can narrow a thread-level call further, down to a single issuing thread.
Tensor layout describes where a logical tensor lives through a storage-first interface. A tile may sit in global memory, shared memory, registers, tensor memory, or accelerator SRAM. The user declares where each tile lives and how its elements are spread across lanes, warps, and registers; that declaration stays attached to the tile. When a primitive is called, the compiler reads those declarations to choose an implementation. A layout is a storage description, not a loop-transformation utility: the user may construct a tile’s layout, but never uses layouts to transform loops.
Tile primitive dispatch turns one call into native IR. From the operand layouts, the execution scope, and the target, or an explicit dispatch= hint, it selects the matching implementation: a copy from global to shared resolves to TMA, shared to register to ldmatrix, and tensor memory to register to tcgen05.ld; a matrix multiply resolves to WGMMA, tcgen05, or a systolic-array instruction. Dispatch then generates the loops and addressing needed to apply that instruction across the whole tile.
These ingredients combine wherever scope matters. In the GEMM epilogue below, warpgroup-scoped and thread-scoped primitives sit in the same region: the Tx.wg.* calls move and cast a tile across the warpgroup, while a final thread-scoped Tx.copy_async , guarded by an explicit issuing-thread predicate, performs the TMA store.
The excerpts above are simplified. For the full picture, here are two roles from a complete FP16/BF16 GEMM kernel — a TMA producer and the tensor-memory writeback. You do not need to read them line by line. The point is that everything to do with orchestration (pipeline state, barrier protocol, role selection, low-level synchronization intrinsics like tcgen05.wait and cp_async.bulk ) stays in ordinary source code, while the recurring data movement appears as tile primitives whose lowering is selected from scope, layout, and dispatch configuration.
Of the three ingredients, layout involves the most design decisions, so it is worth a closer look.
A Storage-First Interface for Tensor Layouts
TIRx treats layout as a first-class representation of tensor storage. Readers familiar with CuTe will recognize the territory: both systems use layout to describe how tensor data maps onto hardware resources, but CuTe exposes layout as a programmable interface for deriving how tile work is partitioned across threads, while TIRx uses layout as a storage contract consumed by primitive dispatch.
A TIRx layout maps a logical tensor index to physical coordinates on named axes. The model generalizes shape-stride layout by attaching strides to semantic hardware axes and by adding explicit shard , replica , and offset components. Shard describes how logical elements are partitioned across physical axes. Replica describes where the same logical element is replicated. Offset describes where physical placement begins. Specifically,
D (Shard). A list of one or more iterators, each with an extent and a stride on some axis. D partitions the logical index across these iters and produces a base coordinate. This generalizes shape-stride to multiple axes.
R (Replica). A set of replication iterators that enumerate offsets in hardware space, independent of the logical index. Adding each element of this set to the D result yields replication or broadcasting.
O (Offset). A fixed coordinate offset (one integer per axis) is added to every result. This places data at a specific base position or reserves exclusive resources.
A concrete example of the TIRx layout Python API is:
This represents a logical tile distributed over lanes and warps, replicated across another warpgroup, and placed at an offset on the warp axis. Given a logical coordinate (i, j) in (8, 16) shape space, it maps to the warp, lane, and reg axes, respectively, by computing
For example, element 57 at logical (3, 9) maps to:
base location: 6@warpid, 12@laneid, 1@m
owners (×2 via replica): { warpid=6 laneid=12 }, { warpid=10 laneid=12 }
(Open the interactive demo and click element 57 to see exactly these owners.)
▶ Open the interactive layout demo ↗
TIRx’s layout interface is built around four design choices.
1. Layout is a storage contract, not a work-partitioning interface.
In CuTe, layout is not only a representation of data placement; it is also part of the programming interface for deriving how tile operations are distributed across threads. Users compose, tile, and partition layouts to express data and work distribution for copy and compute operations. TIRx draws the boundary differently. Users describe the storage layout of each tile and call tile primitives over those tiles. The layout records how logical tensor coordinates map to physical hardware coordinates, including sharding, replication, and offset; it is not the surface used to construct the execution partitioning. When a primitive is lowered, dispatch uses the operand layouts, execution scope, and backend target to generate the thread partitioning, loop nest, addressing, and instruction sequence. In this sense, TIRx layout only needs to represent storage precisely; the transformation logic lives inside primitive dispatch rather than in user-written layout composition.
2. Layout maps logical tensor coordinates to physical hardware coordinates.
Explicit replica and offset structure come from the designated logical-to-physical formulation. One alternative way to formalize layouts is to map physical locations to logical coordinates, such that replication—one logical element stored in multiple physical locations—can still be defined as a point-valued function. However, for tensors that span physical locations in a strided pattern, some physical locations may not have a well-defined mapping.
3. Layout supports general shapes.
Modern kernels frequently use shapes that do not fit a power-of-two-only representation. Global tensors, multi-stage shared-memory buffers, tensor-memory tiles, accelerator scratchpads, and distributed tensors all produce general shapes in practice. TIRx layout therefore starts from general shape support instead of treating it as a special case. This matters for block-scaled GEMM scale-factor

[truncated]
