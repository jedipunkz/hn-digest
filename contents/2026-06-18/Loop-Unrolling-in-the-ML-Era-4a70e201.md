---
source: "https://hiraditya.github.io/posts/why-loop-unrolling-is-popular-again/"
hn_url: "https://news.ycombinator.com/item?id=48578506"
title: "Loop Unrolling in the ML Era"
article_title: "Loop Unrolling in the ML Era | Aditya Kumar"
author: "matt_d"
captured_at: "2026-06-18T01:02:18Z"
capture_tool: "hn-digest"
hn_id: 48578506
score: 1
comments: 0
posted_at: "2026-06-17T23:32:25Z"
tags:
  - hacker-news
  - translated
---

# Loop Unrolling in the ML Era

- HN: [48578506](https://news.ycombinator.com/item?id=48578506)
- Source: [hiraditya.github.io](https://hiraditya.github.io/posts/why-loop-unrolling-is-popular-again/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T23:32:25Z

## Translation

タイトル: ML 時代のループ展開
記事のタイトル: ML 時代のループ展開 |アディティア・クマール
説明: 最新のワイド SIMD ベクトル エンジン、Tensor コア アレイ、シストリック アレイなどのカスタム ディープ ラーニング アクセラレータのいずれであっても、大規模なコンピューティング アーキテクチャをお持ちの場合は、1 つの根本的な問題、つまり獣に餌を与えるという問題に直面します。膨大な実行幅があるが、命令がボトルネックになっている場合
[切り捨てられた]

記事本文:
ML 時代のループ展開 | Aditya Kumar Aditya Kumar コンパイラー、ソフトウェア パフォーマンスの最適化
ホーム カテゴリ タグ アーカイブ ML 時代のホーム ループ アンローリングについて ポスト キャンセル ML 時代のループ アンローリング
最新のワイド SIMD ベクトル エンジン、Tensor コア アレイ、シストリック アレイなどのカスタム ディープ ラーニング アクセラレータのいずれであっても、大規模なコンピューティング アーキテクチャをお持ちの場合は、1 つの根本的な問題、つまり獣に餌を与えるという問題に直面します。膨大な実行幅がありますが、命令が分岐オーバーヘッドと短い基本ブロックによってボトルネックになっている場合、それらの実行ユニットはアイドル状態になります。
このアーキテクチャの変化により、ループ展開をめぐる活動と注目が大幅に増加しました。
ループ展開は新しい概念ではありません。これは、もともとループ制御のオーバーヘッドを削減し、命令レベルの並列処理 (ILP) を公開するために設計された古典的なコンパイラの最適化です。 ML 以前の時代では、一般的な Web ワークロードやモバイル ワークロードはきめ細かい ILP に大きく依存していないため、あまり注目されていませんでした。しかし現在、非常に特殊な理由により、その使用量が大幅に増加しています。それは、機械学習のワークロード、特に高密度の matmul を高度にベクトル化してタイル化する必要があるということです。最新のコンパイラでは、自動ベクトル化とループ展開が密接に結合されています。ループを展開することにより、コンパイラーは独立した同型スカラー命令のより大きなシーケンスを公開し、これらの演算をワイド SIMD ベクトルに安全にパックすることが大幅に容易になります。
これらのタイル行列乗算のスループットを最大化するには、パイプラインを完全にフル状態に保つ必要があります。ループ アンローリングはソフトウェア パイプラインの重要なイネーブラーであり、コンパイラが次のタイルのメモリ フェッチと現在のタイルの計算をオーバーラップできるようにします 1。さらに、この概念は現在、さまざまな分野に拡張されています。

o 物理領域: 空間ループの展開により、反復はハードウェア処理要素の 2D グリッドに直接マッピングされ、チップのデータフロー全体を決定します。最新の ML ハードウェアを最大限に活用するために、私たちはあらゆる単一レベルの抽象化でループを積極的に展開しています。
ループの展開は、OoO プロセッサの全盛期には非常に一般的な最適化であったため、プログラマーは展開されたループを手動で作成して、ILP をハードウェアに公開することがよくありました 2 。実際には、すべての最適化コンパイラーにはループ アンローリング パスがあり、コンパイラー コースでループ アンローラー 3 を教えるのが一般的です。
C (手動アンローリング + プラグマ): ユーザー空間コードの最下位レベル (カスタム C カーネルや CUDA カーネルなど) では、開発者は多くの場合、パフォーマンスをコンパイラーのヒューリスティックな推測に任せることを拒否します。これらは、コンパイラ ディレクティブ (特に #pragma unroll ) を使用してループをアンロールするようにコンパイラに明示的に指示します。
1
2
3
4
5
6
7
void mac_kernel_pragma ( float * a , float * b , float * c ) {
// コンパイラに次のループを完全に展開するよう強制します
#プラグマアンロール
for ( int i = 0 ; i < 4 ; ++ i ) {
c [ i ] = a [ i ] * b [ i ] + c [ i ];
}
}
手動による展開: 場合によっては、開発者が単に命令を順番に書き出して、ループを完全に手作業で排除することもあります。
1
2
3
4
5
6
void mac_kernel_manual ( float * a , float * b , float * c ) {
c [ 0 ] = a [ 0 ] * b [ 0 ] + c [ 0 ];
c [ 1 ] = a [ 1 ] * b [ 1 ] + c [ 1 ];
c [ 2 ] = a [ 2 ] * b [ 2 ] + c [ 2 ];
c [ 3 ] = a [ 3 ] * b [ 3 ] + c [ 3 ];
}
マクロベースのアンローリング: 手動入力ではエラーが発生しやすく、プラグマが信頼できない大きなブロックの場合、開発者は歴史的に C プリプロセッサ マクロを使用して、コンパイラーがコードを解析する前にアンローリングを強制してきました。
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
#define MAC(i, k) c[i + k] += a[i + k] * b[i + k];
#define UNROLL_4(i

) \
MAC(i, 0) \
MAC(i, 1) \
MAC(i, 2) \
MAC(i, 3)
void mac_kernel_macro ( float * a , float * b , float * c , int N ) {
for ( int i = 0 ; i < N ; i += 4 ) {
UNROLL_4 ;
}
// 残りを処理します
for ( int i = N - ( N % 4 ); i < N ; ++ i ) {
MAC ( i , 0 );
}
}
プラグマ、手動展開、マクロのいずれを使用する場合でも、目標は同じです。最終的なアセンブリは 4 つの連続した積和演算となり、ブランチのオーバーヘッドが完全に除去され、実行ユニットに最大限の並列性が提供されます。
トレードオフ: #pragma unroll の主な危険性は、これが盲目的なディレクティブであり、ハードウェアの制限に関係なくコンパイラーにアンロールを強制することです。開発者にとって、過剰な手動展開の直接的な結果は、深刻な登録プレッシャーになります。ループ本体を拡張することにより、カーネルはより多くのアーキテクチャ レジスタを必要とします。 GPU では、これによりレジスタ スピルが発生し、アクティブなワープ占有が大幅に減少し、パイプラインが停止します。開発者は、これらのレジスタ制約に対して計算密度のバランスを慎重にとる必要があります 4 。 (もう 1 つの主要なリスクは、コードの肥大化と命令キャッシュの削除です。これについては、以下の「コンパイラー」セクションで詳しく説明します)。
C++ (コンパイル時評価): ループ境界がコンパイル時に静的に知られている場合 (小さな 4x4 行列タイルの寸法など)、最新の C++ 開発者はコンパイル時評価機能を活用します。 if constexpr と組み合わせたテンプレートを使用することで、大規模な直線のコード ブロックを再帰的に生成できます。これにより、コードがコンパイラーの中間エンドに到達する前に、分岐のオーバーヘッドが完全に排除されます。
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
テンプレート < int N >
インライン void プロセス ( float * a 、 float * b 、 float * c ) {
if constexpr ( N > 0 ) {
プロセス < N - 1 > ( a , b , c );
c [ N - 1 ] += a [ N - 1 ] * b [ N - 1 ];
}
}
v

oid process_tile ( float * a 、 float * b 、 float * c ) {
プロセス < 4 > ( a 、 b 、 c );
}
(注: 最新の C++17 では、 if constexpr を使用すると、これが信じられないほどきれいになり、面倒なベースケース テンプレートの特殊化の必要がなくなります。コードをさらに簡潔にするために、開発者はよく std::index_sequence と組み合わせた折り畳み式を使用します)。
手動およびコンパイル時のアンロールは絶対的な確実性を提供しますが、ほとんどの実稼働ソフトウェアはコンパイラーの最適化パイプラインに依存してループを自動的にアンロールします。ただし、この責任をコンパイラーに移すと、構造が大幅に複雑になります。ループを安全に展開するために、コンパイラーは単に命令をコピーして貼り付けるだけでは済みません。プログラムの制御フロー グラフ (CFG) を構造的に変換して、任意のループ バインドの正確性を保証し、エントリ ガード、複数のループ ヘッダー、および残りの実行パスを慎重に管理する必要があります。
制御フロー グラフ (CFG) 変換の視覚化:
ループ展開の実際の複雑さを理解するために、展開が適用される前の標準的なループ構造を次に示します。これには、実行回数が 0 回の場合にループ全体をスキップする最初のガード分岐 ( Z -> C ) が含まれます。
グラフTD
Z["Z（プリヘッダー）"] --> A["A（ヘッダー）"]
Z --> C["C (終了)"]
A --> B["B (ラッチ/終了)"]
B --> A
B --> C
アンロール係数 2 を適用すると、ループ本体 (ヘッダーとラッチ) が複製されます。 2 の完全な倍数ではないトリップ カウントを処理するために、クリーンアップ (または剰余) ループが導入されています。展開された本体全体に十分な反復が残っていない場合、プリヘッダーは実行をクリーンアップ ループにルーティングします。
グラフTD
Z["Z(ガード)"] --> Z1["Z1(プリヘッダ)"]
Z --> C2["C(終了)"]
Z1 --> A1["A1 (ヘッダー)"]
Z1 --> クリーンアップ_A["クリーンアップ A (ヘッダー)"]
A1 --> B1["B1(ラッチ)"]
B1 --> A2["A2 (ヘッダー)"]
A2 --> B2["B2 (ラッチ/終了)

ng)"]
B2 --> A1
B2 --> クリーンアップ_A
クリーンアップ_A --> クリーンアップ_B["クリーンアップ B (ラッチ / 終了)"]
クリーンアップ_B --> C2
ただし、高性能計算カーネル (一般行列乗算や GEMM など) では、スカラー クリーンアップ ループに分岐すると、深刻なパイプライン バブルと分岐予測ミスのオーバーヘッドが発生します。パフォーマンス エンジニアは、このフォールバック パスを積極的に回避します。一般的な戦略は、テンソル次元がアンロール係数 (およびベクトル幅) の正確な倍数になるように体系的にパディングし、コンパイラーが剰余がゼロであることを数学的に証明し、クリーンアップ ループを完全に省略できるようにすることです。あるいは、ループ ピーリングなどの手法を使用して、初期の位置合わせされていない反復を個別に処理し、展開されたメインのボディが完全な位置合わせ制約の下で実行されるようにします。
MLIR (Affine Dialects & Tiling Machinery): 最新の ML コンパイラ スタック (Triton 5、IREE、XLA など) では、コードが低レベルのスカラー最適化に到達するずっと前に、MLIR の高レベルの方言が構造ループの展開を処理します。
アンローリングに affine.for を使用する: 従来のループ構造の場合、MLIR の Affine 方言は組み込みのアンローリング パスを提供します。 mlir-opt -affine-loop-unroll="unroll-factor=4" を実行すると、コンパイラは構造化された affine.for ループを自動的に変換して ILP を公開します。
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
// 展開前
affine.for %i = 0 ～ 4 {
%v = affine.load %A [%i ] : memref < 4 xf 32 >
// ... プロセス %v ...
}
// 展開後 (概念的に)
%v0 = affine.load %A [ 0 ] : memref < 4 xf 32 >
// ... プロセス %v0 ...
%v1 = affine.load %A [1] : memref < 4 xf 32 >
// ... プロセス %v1 ...
%v2 = affine.load %A [ 2 ] : memref < 4 xf 32 >
// ... プロセス %v2 ...
%v3 = affine.load %A [ 3 ] : memref < 4 xf 32 >
// ... プロセス %v3 ...
MLIR におけるこの構造展開により、

中間表現は LLVM に引き下げられ、多次元ベクトル化という重労働はすでに完了しています。
関連するタイル マシン (データ レイアウトの最適化): 技術的にはループ アンローリングではありませんが、最新の ML コンパイラーは、行列の乗算中にキャッシュの局所性を確保するためにデータをタイル レイアウトにパックするためにデータ レイアウトの最適化に大きく依存しています。タイル上のこの構造的な「ループ」を逆転させるために、MLIR は linalg.unpack 操作 (古い tensor.unpack を置き換える) を使用して、レイアウトを元の反復空間に平坦化します。
1
2
3
4
5
// 8x1 タイルの 2x3 グリッドをフラットな 15x3 テンソルにアンパックします。
%A_unpack = linalg.unpack %A
inner_dims_pos = [ 0 , 1 ]
inner_tiles = [ 8 , 1 ]
%A_unpack_empty へ: テンソル < 2 x 3 x 8 x 1 xi 32 > -> テンソル < 15 x 3 xi 32 >
LLVM-IR ( -funroll-loops 実装): これは、古典的な LLVM ミドルエンド パスです (主に llvm/lib/Transforms/Scalar/LoopUnrollPass.cpp に実装されます)。
仕組み: LLVM アンローラーは、Scalar Evolution (SCEV) 分析に大きく依存しています。 SCEV は、ループの反復にわたって変数がどのように変化するかを数学的にモデル化しようとします。 SCEV がループのトリップ カウントを正確に決定できる場合、アンローラーは複雑なコスト モデルを評価して、ループ サイズ、ターゲット アーキテクチャ、および命令キャッシュの制限を考慮に入れて、アンローリングが収益性があるかどうかを判断します。
現実と限界: 理論は健全ですが、あらゆるエッジケースやアーキテクチャに合わせてこれらのヒューリスティックを調整することはほぼ不可能です。 LLVM Discourse または GitHub Issue Tracker を参照すると、パフォーマンスの低下を引き起こす -funroll-loops に関する苦情の墓場が見つかります。次の 2 つの大きな問題が継続的に表面化しています。
コードの肥大化と I キャッシュ ミス: 積極的なアンロールによりループ本体が複製され、静的コード サイズが膨張します。この新しく拡張されたブロックが

L1 命令キャッシュ (または最新の CPU ではマイクロ演算キャッシュ/ループバック バッファ) の容量を超えると、CPU は命令のフェッチ中にストールし、計算の利点が完全に無効になります (たとえば、問題 #42332: 過剰なループ展開を参照してください。パフォーマンスの専門家は、ボトルネックがベクトル単位である場合、単純なループの展開には利点がゼロであるが、キャッシュから他のコードを追い出すことで積極的にパフォーマンスを低下させると指摘しています)。レジスターの流出: これはサイレントキラーです。直線コードの大量のブロックを生成すると、バックエンドのレジスタ アロケータに負荷がかかります。アンロールすると、任意の時点でアクティブなライブ変数の数が人為的に増加します。ターゲットにこれらの変数を保持するのに十分な物理レジスタがない場合、コンパイラは変数をスタック (メイン メモリ) に「スピル」します。このレジスタスピルにより、高速で計算重視のカーネルが、低速でメモリ重視のボトルネックに変わります (たとえば、行列乗算のアンロールが LICM と相互作用して、最大 100 個のロードをプリヘッダーにホイストし、ターゲットの 16 レジスタ ファイルを完全に圧倒するこの LLVM 談話スレッドを参照してください)。複雑な SCEV 解析と脆弱なコスト モデルに大きく依存しているため、多くのパフォーマンス エンジニアは手動または構造を好みます。

[切り捨てられた]

## Original Extract

If you have a massive compute architecture—whether it’s a modern wide-SIMD vector engine, a Tensor Core array, or a custom deep learning accelerator like a Systolic Array—you face one fundamental problem: feeding the beast. You have immense execution width, but if your instructions are bottlenecked
[truncated]

Loop Unrolling in the ML Era | Aditya Kumar Aditya Kumar Compilers, Software Performance optimizations
HOME CATEGORIES TAGS ARCHIVES ABOUT Home Loop Unrolling in the ML Era Post Cancel Loop Unrolling in the ML Era
If you have a massive compute architecture—whether it’s a modern wide-SIMD vector engine, a Tensor Core array, or a custom deep learning accelerator like a Systolic Array—you face one fundamental problem: feeding the beast. You have immense execution width, but if your instructions are bottlenecked by branch overhead and short basic blocks, those execution units sit idle.
This architectural shift has led to significantly increased activity and attention surrounding loop unrolling.
Loop unrolling isn’t a new concept. It’s a classic compiler optimization originally designed to reduce loop control overhead and expose Instruction-Level Parallelism (ILP). In the pre-ML era, it received less attention because typical web or mobile workloads don’t rely heavily on fine-grained ILP. But today, we are seeing a massive surge in its usage for a very specific reason: machine learning workloads—specifically dense matmuls—need to be heavily vectorized and tiled. In modern compilers, auto-vectorization and loop unrolling are tightly coupled. By unrolling a loop, the compiler exposes a larger sequence of independent, isomorphic scalar instructions, making it significantly easier to safely pack those operations into wide SIMD vectors.
To maximize throughput on these tiled matrix multiplications, the pipeline must be kept completely full. Loop unrolling is the critical enabler for software pipelining , allowing the compiler to overlap memory fetches for the next tile with compute for the current tile 1 . Furthermore, the concept has now expanded into the physical realm: with spatial loop unrolling , iterations are mapped directly onto 2D grids of hardware Processing Elements, dictating the chip’s entire dataflow. To fully utilize modern ML hardware, we are aggressively unrolling loops at every single level of abstraction.
Loop unrolling was such a common optimization during the OoO processor heydays that programmers often wrote unrolled loops by hand to expose ILP to the hardware 2 . Practically every optimizing compiler has a loop unrolling pass and it is common for compiler courses to teach loop unroller 3 .
C (Manual Unrolling + Pragmas): At the lowest level of user-space code (like custom C or CUDA kernels), developers often refuse to leave performance up to the compiler’s heuristic guesses. They explicitly instruct the compiler to unroll loops using compiler directives, most notably #pragma unroll .
1
2
3
4
5
6
7
void mac_kernel_pragma ( float * a , float * b , float * c ) {
// Force the compiler to unroll the next loop completely
#pragma unroll
for ( int i = 0 ; i < 4 ; ++ i ) {
c [ i ] = a [ i ] * b [ i ] + c [ i ];
}
}
Manual Unrolling: Sometimes, developers simply write out the instructions sequentially, eliminating the loop entirely by hand:
1
2
3
4
5
6
void mac_kernel_manual ( float * a , float * b , float * c ) {
c [ 0 ] = a [ 0 ] * b [ 0 ] + c [ 0 ];
c [ 1 ] = a [ 1 ] * b [ 1 ] + c [ 1 ];
c [ 2 ] = a [ 2 ] * b [ 2 ] + c [ 2 ];
c [ 3 ] = a [ 3 ] * b [ 3 ] + c [ 3 ];
}
Macro-Based Unrolling: For larger blocks where manual typing is error-prone but pragmas aren’t trusted, developers historically used C preprocessor macros to force the unrolling before the compiler even parses the code:
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
#define MAC(i, k) c[i + k] += a[i + k] * b[i + k];
#define UNROLL_4(i) \
MAC(i, 0) \
MAC(i, 1) \
MAC(i, 2) \
MAC(i, 3)
void mac_kernel_macro ( float * a , float * b , float * c , int N ) {
for ( int i = 0 ; i < N ; i += 4 ) {
UNROLL_4 ;
}
// Handle the remainder
for ( int i = N - ( N % 4 ); i < N ; ++ i ) {
MAC ( i , 0 );
}
}
Whether you use pragmas, manual unrolling, or macros, the goal is the same: the final assembly will be four back-to-back multiply-accumulate operations, entirely removing the branch overhead and exposing maximum parallelism to the execution units.
Tradeoffs: The primary danger of #pragma unroll is that it is a blind directive , forcing the compiler to unroll regardless of the hardware’s limits. For developers, the immediate consequence of excessive manual unrolling is severe register pressure. By expanding the loop body, the kernel requires exponentially more architectural registers. On GPUs, this leads to register spilling and drastically reduces active warp occupancy, choking the pipeline. Developers must carefully balance compute density against these register constraints 4 . (The other major risk is code bloat and instruction-cache eviction, which we will detail in the Compiler section below).
C++ (Compile-Time Evaluation): When the loop bounds are statically known at compile-time (like the dimensions of a small 4x4 matrix tile), modern C++ developers leverage compile-time evaluation features. By using templates combined with if constexpr , they can recursively generate massive, straight-line blocks of code. This eliminates branch overhead entirely before the code even reaches the compiler’s middle-end.
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
template < int N >
inline void process ( float * a , float * b , float * c ) {
if constexpr ( N > 0 ) {
process < N - 1 > ( a , b , c );
c [ N - 1 ] += a [ N - 1 ] * b [ N - 1 ];
}
}
void process_tile ( float * a , float * b , float * c ) {
process < 4 > ( a , b , c );
}
(Note: Modern C++17 makes this incredibly clean with if constexpr , eliminating the need for messy base-case template specializations. For even more concise code, developers often use fold expressions combined with std::index_sequence ).
While manual and compile-time unrolling offer absolute certainty, most production software relies on the compiler’s optimization pipeline to automatically unroll loops. However, shifting this responsibility to the compiler introduces significant structural complexity. To safely unroll a loop, the compiler cannot simply copy-paste instructions; it must structurally transform the program’s Control Flow Graph (CFG) to guarantee correctness for any arbitrary loop bound, carefully managing entry guards, multiple loop headers, and remainder execution paths.
Visualizing the Control Flow Graph (CFG) Transformation:
To appreciate the true complexity of loop unrolling, here is the standard loop structure before any unrolling is applied, including the initial guard branch ( Z -> C ) which skips the loop entirely if it runs zero times:
graph TD
Z["Z (Preheader)"] --> A["A (Header)"]
Z --> C["C (Exit)"]
A --> B["B (Latch / Exiting)"]
B --> A
B --> C
By applying an unroll factor of 2, the loop body (header and latch) is duplicated. To handle trip counts that are not a perfect multiple of 2, a Cleanup (or remainder) loop is introduced. The preheader routes execution to the cleanup loop if there aren’t enough iterations left for the full unrolled body:
graph TD
Z["Z (Guard)"] --> Z1["Z1 (Preheader)"]
Z --> C2["C (Exit)"]
Z1 --> A1["A1 (Header)"]
Z1 --> Cleanup_A["Cleanup A (Header)"]
A1 --> B1["B1 (Latch)"]
B1 --> A2["A2 (Header)"]
A2 --> B2["B2 (Latch / Exiting)"]
B2 --> A1
B2 --> Cleanup_A
Cleanup_A --> Cleanup_B["Cleanup B (Latch / Exiting)"]
Cleanup_B --> C2
However, in high-performance compute kernels (such as General Matrix Multiplies or GEMMs), branching into a scalar cleanup loop introduces severe pipeline bubbles and branch misprediction overhead. Performance engineers actively avoid this fallback path. A common strategy is to systematically pad tensor dimensions so they are exact multiples of the unroll factor (and vector width), allowing the compiler to mathematically prove the remainder is zero and elide the cleanup loop entirely. Alternatively, techniques like loop peeling handle the initial unaligned iterations separately, ensuring the main unrolled body executes under perfect alignment constraints.
MLIR (Affine Dialects & Tiling Machinery): In the modern ML compiler stack (like Triton 5 , IREE, or XLA), high-level dialects in MLIR handle structural loop unrolling long before the code reaches lower-level scalar optimizations.
Using affine.for Unrolling: For traditional loop structures, MLIR’s Affine dialect provides built-in unrolling passes. By running mlir-opt -affine-loop-unroll="unroll-factor=4" , the compiler automatically transforms structured affine.for loops to expose ILP:
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
// Before Unrolling
affine.for %i = 0 to 4 {
%v = affine.load %A [ %i ] : memref < 4 xf 32 >
// ... process %v ...
}
// After Unrolling (conceptually)
%v0 = affine.load %A [ 0 ] : memref < 4 xf 32 >
// ... process %v0 ...
%v1 = affine.load %A [ 1 ] : memref < 4 xf 32 >
// ... process %v1 ...
%v2 = affine.load %A [ 2 ] : memref < 4 xf 32 >
// ... process %v2 ...
%v3 = affine.load %A [ 3 ] : memref < 4 xf 32 >
// ... process %v3 ...
This structural unrolling in MLIR ensures that by the time the intermediate representation is lowered to LLVM, the heavy lifting of multi-dimensional vectorization is already complete.
Related Tiling Machinery (Data-Layout Optimization): While not technically loop unrolling, modern ML compilers heavily rely on data-layout optimization to pack data into tiled layouts for cache locality during matrix multiplication. To reverse this structural “looping” over tiles, MLIR uses the linalg.unpack operation (which replaced the older tensor.unpack ) to flatten the layout back into its original iteration space:
1
2
3
4
5
// Unpacking a 2x3 grid of 8x1 tiles back into a flat 15x3 tensor
%A_unpack = linalg.unpack %A
inner_dims_pos = [ 0 , 1 ]
inner_tiles = [ 8 , 1 ]
into %A_unpack_empty : tensor < 2 x 3 x 8 x 1 xi 32 > - > tensor < 15 x 3 xi 32 >
LLVM-IR ( -funroll-loops implementation): This is the classic LLVM middle-end pass (implemented primarily in llvm/lib/Transforms/Scalar/LoopUnrollPass.cpp ).
How it Works: The LLVM unroller relies heavily on Scalar Evolution (SCEV) analysis. SCEV attempts to mathematically model how variables change across loop iterations. If SCEV can accurately determine the loop’s trip count, the unroller evaluates an intricate cost model to decide whether unrolling is profitable, factoring in loop size, the target architecture, and instruction cache limits.
The Reality & Limitations: While the theory is sound, tuning these heuristics for every edge case and architecture is nearly impossible. If you browse the LLVM Discourse or GitHub Issue Tracker , you will find a graveyard of complaints about -funroll-loops causing performance regressions. Two major problems continually surface:
Code Bloat & I-Cache Misses: Aggressive unrolling duplicates the loop body, ballooning static code size. If this newly expanded block exceeds the capacity of the L1 instruction cache (or the micro-op cache / loopback buffer on modern CPUs), the CPU stalls while fetching instructions, completely negating the compute benefits (e.g., see Issue #42332: Excessive loop unrolling , where performance experts point out that unrolling simple loops provides zero advantage when the bottleneck is the vector unit, but actively hurts performance by evicting other code from the cache). Register Spilling: This is the silent killer. Generating massive blocks of straight-line code puts pressure on the backend register allocator. By unrolling, you artificially increase the number of live variables active at any given moment. If the target doesn’t have enough physical registers to hold these variables, the compiler is forced to “spill” them to the stack (main memory). This register spilling turns a fast, compute-bound kernel into a slow, memory-bound bottleneck (e.g., see this LLVM Discourse thread where unrolling a matrix multiplication interacted with LICM to hoist up to 100 loads into a preheader, completely overwhelming the target’s 16-register file). Because of this heavy reliance on complex SCEV analysis and fragile cost models, many performance engineers prefer manual or structur

[truncated]
