---
source: "https://rocm.blogs.amd.com/software-tools-optimization/accelerating-llm-inference-on-amd-gpus-with-low-latency-gemms/README.html"
hn_url: "https://news.ycombinator.com/item?id=48737691"
title: "Accelerating LLM Inference on AMD GPUs with Low-Latency GEMMs"
article_title: "Accelerating LLM Inference on AMD GPUs with Low-Latency GEMMs — ROCm Blogs"
author: "matt_d"
captured_at: "2026-06-30T19:33:35Z"
capture_tool: "hn-digest"
hn_id: 48737691
score: 2
comments: 0
posted_at: "2026-06-30T19:03:29Z"
tags:
  - hacker-news
  - translated
---

# Accelerating LLM Inference on AMD GPUs with Low-Latency GEMMs

- HN: [48737691](https://news.ycombinator.com/item?id=48737691)
- Source: [rocm.blogs.amd.com](https://rocm.blogs.amd.com/software-tools-optimization/accelerating-llm-inference-on-amd-gpus-with-low-latency-gemms/README.html)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T19:03:29Z

## Translation

タイトル: 低遅延 GEMM を使用した AMD GPU での LLM 推論の高速化
記事のタイトル: 低遅延 GEMM を使用した AMD GPU での LLM 推論の高速化 — ROCm ブログ
説明: FlyDSL 低遅延 GEMM が、Split-K、K スライス並列処理、LDS ベースのパイプラインを使用して AMD GPU での LLM デコードを高速化する方法を学びます。

記事本文:
メインコンテンツにスキップ
トップに戻る
Ctrl + K
ROCm™ ブログ
ホーム
低遅延 GEMM を使用した AMD GPU での LLM 推論の高速化
LLM の提供にデコード遅延が重要なのはなぜですか?
Small-M 、Large-K GEMM のパフォーマンスが低いのはなぜですか?
実際のモデルで GEMM 形状をデコードする
中心となるアイデア: LDS パイプライン化 Split-K
Inter-CTA Split-K: Small-M GEMM 向けの CTA の増加
CTA 内 K スライス分割: より多くのワープグループ並列処理
マルチステージ LDS パイプライン: K ブロックを飛行状態に保つ
単一起動の Split-K 同期
CTA 内 K スライス分割のための LDS 削減
実装ノート: アルゴリズムから FlyDSL カーネルまで
低遅延 GEMM を使用した AMD GPU での LLM 推論の高速化 #
2026 年 6 月 29 日、Yutao Xu、Xiaobing Zhang、Hattie Wu、Felix Li、Lingpeng Jin、Carlus Huang、Peng Sun、Barsoum Emad 著。
11 分で読めます。 |合計 2577 単語。
ソフトウェアツールと最適化
AI/ML、パフォーマンス
ユーザー、開発者、AI
ユタオ・シュー、シャオビン・チャン、ハッティ・ウー、フェリックス・リー、リンペン・ジン、カルス・ファン、ペン・スン、バルスーム・エマド
英語
-->
大規模言語モデルの推論はますますインタラクティブになってきています。ユーザーは、チャットボット、コーディングアシスタント、エージェント、リアルタイムコパイロットが迅速に応答し、トークンをスムーズにストリーミングし、同時負荷下でも応答性を維持することを期待しています。この設定では、デコード時のレイテンシーは単なるバックエンドの指標ではありません。それは知覚される品質に直接影響します。
このブログでは、その推論パスの小さいながらも重要な部分、つまり小さい M 、大きい N および K を備えたデコード時 GEMM、BF16/FP16 入力、オプションのバイアス、および実際のモデル全体で繰り返される形状について説明します。これらの形状では、従来の GEMM タイリングが十分に活用されないままになる可能性があるため、デコード パスの最適化の対象として有用です。
主なテクニックは LDS-Pipelined Split-K GEMM です。ロング K リダクションは CTA 間で分割され、さらにワープ グループ間でスライスされます。

各 CTA の側面に配置され、多段の LDS メモリ パイプラインを通過し続けます。 AMD GPU では、LDS は Local Data Share を意味し、CTA 内の高速連携に使用されるオンチップ スクラッチパッド メモリです。
また、このアイデアを AITER FlyDSL カーネル ファミリとしてどのように実装するかについても説明します。 FlyDSL は、MFMA 選択、LDS レイアウト、非同期コピー、明示的な同期などの低レベルの ROCm™ ソフトウェア詳細を維持しながら、デコードで表示されるモデル寸法の形状に特化したバリアントを生成します。ベンチマーク スイープでは、このターゲットを絞ったデコードの最適化により、K = 7168 デコード グリッド [1] 上で最も速い HipblasLT 、AITER Triton 、AITER ASM と比較して平均 1.64 倍の平均レイテンシーの改善に達し、追加の BF16 モデル形状テストでは 1.49 倍の平均レイテンシーの改善に達しました。
LLM の提供にデコード遅延が重要なのはなぜですか? #
LLM の提供には、大きく 2 つのフェーズがあります。
Prefill : モデルがプロンプトを処理します。
Decode : モデルは出力トークンを一度に 1 ステップずつ生成します。
多くのプロンプト トークンを一緒に処理できるため、プレフィルでは有効 M が大きくなることがよくあります。デコードが違います。各ステップは、特にバッチ処理、スケジューリング、テンソル並列処理、およびリクエストレベルのダイナミクスを考慮した後では、少数のアクティブ トークンのみを処理する可能性があります。
そのため、ユーザー側のレイテンシにとってデコード パフォーマンスが重要になります。
最初のトークンまでの時間は、システムの応答速度に影響します。
出力トークンごとの時間は、ストリーミングのスムーズさに影響します。
トークン間のレイテンシーは、インタラクションが流動的に感じられるかどうかに影響します。
同時実行時のスループットは、応答性を損なうことなくサービスを提供できるユーザーの数に影響します。
図 1 は、このインタラクティブ デコード サービング設定を示し、これらの遅延の問題がユーザー向けパスのどこに現れるかを示しています。
図 1: インタラクティブな LLM デコード サービス。
これらのワークロードの場合、shavi

GEMM のデコードの繰り返しによるオーバーヘッドは、モデル提供レベルで重要になる可能性があります。
Small-M 、Large-K GEMM のパフォーマンスが低いのはなぜですか? #
大規模モデルのデコードでは、GEMM は次のようになります。
C[M, N] = A[M, K] @ B[N, K]^T
視覚的には、カーネルは依然として標準的な GEMM のアイデア、つまり A のタイルと B のタイルから C のタイルを計算することから始まります。問題は、K 次元が非常に長い場合でも、M が小さいと生成される出力タイルが少なすぎることです。
図 2 は、この小さい M と大きい K のボトルネックを示しています。出力グリッドは狭いですが、縮小次元にはまだかなりの作業が含まれています。
図 2: 小さい M 、大きい K GEMM ボトルネック。
ここで、M はデコード ステップまたはマイクロバッチ内のアクティブなトークンの数です。ワークロードを処理する場合、M は 1 、 2 、 4 、 8 、 16 、 32 と小さいことがよくありますが、最大 128 または 256 になる場合もあります。同時に、N と K はモデルに隠されたサイズの次元であり、数千または数万になる可能性があります。
この形状レジームは、一般的な GEMM ライブラリにとって扱いにくいものです。従来のラージタイル GEMM は、すべての計算ユニットをビジー状態に保つために、ブロックごとに十分な M x N の作業を必要とします。 GEMM のデコードでは、自然にはそれが提供されないことがよくあります。その結果、占有率が不足し、電波の利用率が低下し、有用な計算に対してオーバーヘッドが大きすぎます。
CTA タイルの拡大、メモリ結合の改善、LDS ステージング、MFMA に重点を置いたスケジューリング、パイプライン化などの一般的な GEMM 最適化は依然として重要です。ただし、M x N 出力グリッドが小さい場合、それらだけでは十分な独立した作業を作成できません。これが、LDS-Pipelined Split-K が 1 つの最適化層に依存するのではなく、複数の形式の K 並列処理を組み合わせる理由です。
実際のモデルで GEMM 形状をデコードする #
その動機は、合成の正方形の GEMM からではなく、モデルの形状トレースから直接得られました。
現在の LLM 全体で、GEMM 形状をデコードすると、同じパターンが繰り返し表示されます。
M = 1 ～ 256、N = 256 / 2112 / 3072 / 716

8 / 16160 、K = 1536 / 2048 / 7168
M = 1 ～ 256、N = 128 / 640 / 2560 / 2880 / 5120、K = 512 / 2048 / 2880 / 4096
M = 1–256 に加えて、プレフィルのような大きな M があり、 6144 、 4096 、 2048 などの N と K が含まれます。
M = 8–512 、 N = 384 または 1024 、 K = 7168 などの多くのスキニー デコード形状
M = 1–32768 ですが、デコードクリティカルなケースには M = 1/16/32/64 が含まれ、N と K が大きくなります。
M = 1–32768 、 N = 100/200/800 、 K = 5120 などのデコード負荷の高いスキニー投影を使用
重要な観察は、単に「小さな M が存在する」ということではありません。それは、小さい M と大きい K の GEMM がデコード パスのあらゆる場所で発生し、エンドツーエンドのサービス スループットに影響を与えるということです。
したがって、設計目標は、次のようなオーバーヘッドの低い GEMM パスです。
Mが小さくても占有率が良い
中心となるアイデア: LDS パイプライン化 Split-K #
カーネルは、K を 1 つの CTA 内のプライベート シリアル ループではなく、分割可能なリダクション作業として扱います。 1 つの出力タイル C[m_tile, n_tile] の場合、計算は次のようになります。
C_タイル =
A[m_tile, K0] @ B^T[K0, n_tile]
+ A[m_tile, K1] @ B^T[K1, n_tile]
+ A[m_tile, K2] @ B^T[K2, n_tile]
+ ...
LDS パイプライン化 Split-K は、これらの K0/K1/K2/... チャンクを 3 つのレベルで公開します。
Inter-CTA Split-K : 完全な K ディメンションを複数の CTA (ワークグループ) に分割します。
Intra-CTA K スライス分割: 1 つの CTA の K タイルをブロック内の複数のワープ グループに分割します。
マルチステージ LDS パイプライン: グローバルから LDS へのコピー、LDS 読み取り、および MFMA 計算をオーバーラップしながら、リング バッファーされた LDS ステージを通じて K ブロックをパイプライン化します。
これらの層はさまざまな問題を解決します。
M x N のタイルが少なすぎる場合の GPU 占有率の向上
グローバルな蓄積と同期
K-heavy タイルのワープ グループのより適切な使用
LDS のステージングと局所的な縮小
K ブロックが進む間に、グローバルから LDS へのコピー、LDS 読み取り、および MFMA がオーバーラップします。
リングバッファリングされた LDS ステージとスケジューリング
GPU 全体でより多くの作業が可能になり、より有用な作業が可能になります

CTA ごと、よりスムーズな K ブロック パイプライン処理
調整された削減とパイプライン パス
図 3 は、このデザインのタイル化された完全なデータ パスを示しています。選択された C タイルは、1 つのモノリシック CTA によって生成されるわけではありません。ロング K 次元はまず CTA 間の Split-K 範囲に分割され、各 CTA は割り当てられた K ブロックを多段 LDS パイプラインを通じてストリーミングし、CTA 内 K スライス分割により、複数のワープ グループが同じ出力タイルの部分累積を計算できるようになります。これらのローカル パーシャルは、CTA 間の Split-K パーシャルが最終的な C タイルに蓄積される前に、LDS を通じて削減されます。
図 3: LDS パイプライン化された Split-K データ パス。
Inter-CTA Split-K: Small-M GEMM の CTA を増やす #
M が小さい場合、通常の M x N タイル グリッドでは、GPU を飽和させるのに十分な CTA を起動できない可能性があります。
Inter-CTA Split-K は、K に沿って発射グリッドを拡張します。
グリッド = [mn_tiles、split_k]
各 Split-K パーティションは、異なる K 範囲にわたる部分和を計算します。そのパーティションがローカル パイプライン作業と CTA 内 LDS 削減を完了した後、部分的な結果が同じ出力タイルに蓄積されます。
起動ラッパーでは、CTA 間 Split-K が 2 番目のグリッド ディメンションとして表示されます。
bm = ( m + BLOCK_M - 1 ) // BLOCK_M
hgemm_kernel (C、A、B、BIAS、m、セマフォ、シグナル)。起動（
グリッド = ( bm * N_BLOCKS , SPLIT_K , 1 ),
ブロック = ( BLOCK_THREADS , 1 , 1 ),
ストリーム = ストリーム 、
)
これは、次のような形状をデコードする場合に特に便利です。
M = 1、2、4、8、16
N = 2560 / 2880 / 5120
K = 2880 / 4096 / 7168
この追加の Split-K 次元がなければ、十分な独立した作業が存在しない可能性があります。
CTA 内 K スライス分割: より多くのワープグループ並列処理 #
Inter-CTA Split-K により、CTA の数が増加します。 CTA 内 K スライス分割により、1 つの CTA 内で有用な作業が増加します。
カーネルは、複数のワープ グループを同じタイルの異なる K スライスに割り当てます。各g

group は部分的な累積を計算します。 CTA の最後に、これらの部分的な結果は、書き戻す前に LDS を通じて削減されます。
K-heavy タイルの並列処理が増加します。
ワープ グループ全体に作業を分散することで、レジスター プレッシャーを制御します。
マルチステージ LDS パイプライン: K ブロックを飛行状態に保つ #
3 番目の層は時間的な層です。 CTA が K 範囲を所有すると、引き続き次の計算を繰り返す必要があります。
C_tile += A[m_tile, K_i] @ B^T[K_i, n_tile]
多くの連続する K_i ブロックに対して。これらのブロックをシリアルな読み込みと計算のシーケンスとして扱う代わりに、カーネルは STAGES を LDS タイルのリング バッファーとして使用します。以前に埋められたステージは LDS 読み取りと MFMA によって消費されますが、別のステージは将来のグローバルから LDS へのコピーに再利用されます。
B_TO_LDS パスでは、A と B の両方がこの LDS リングに参加します。プロローグは最初に STAGES - 1 K ブロックをプリフェッチし、次にホット ループが 1 つのステージを消費しながら、将来のステージのコピーを発行します。
range_constexpr の の場合 ( ステージ - 1 ):
ldg_sts_b_async ( ks_begin + s * BLOCK_K , s )
ldg_sts_a_async ( ks_begin + s * BLOCK_K , s )
bki の場合、範囲内の状態 ( 0 、 BLOCK_K_LOOPS - ( STAGES - 1 )、 1 、 init = init_state ):
k_offset = 状態 [0]
current_stage = fx 。インデックス (状態 [ 1 ])
next_stage = ( current_stage + 1 ) % ステージ
write_stage = ( current_stage + ステージ - 1 ) % ステージ
__バリア (( ステージ - 2 ) * LDG_WAIT_COUNT )
ldg_sts_b_async ( k_offset + ( STAGES - 1 ) * BLOCK_K , write_stage )
ldg_sts_a_async ( k_offset + ( STAGES - 1 ) * BLOCK_K , write_stage )
c_frags_new = ldmatrix_compute_tile_streaming ( current_stage , c_frags )
hot_loop_scheduler ()
その後、ホット ループは一度に 1 K ブロックずつリングを進めます。 current_stage は消費される LDS ステージであり、write_stage = current_stage + STAGES - 1 modulo STAGES は将来の K ブロックを受け取るステージです。の

待機カウントは意図的に新しいコピーを未処理のままにします。
__バリア (( ステージ - 2 ) * LDG_WAIT_COUNT )
つまり、ループは、実行中のグローバルから LDS へのコピーをすべて排出せずに、現在のステージが安全に読み取れるようになるまで待機します。概念的には:
現在のステージ: LDS は K ブロック i のフィード MFMA を読み取ります
書き込みステージ: グローバルから LDS へのコピーにより、K ブロック i + STAGES - 1 がもたらされます
次のステージ : 次のループ反復で現在のステージになります
スケジューラは、hot_loop_scheduler() で VMEM、LDS 読み取り、および MFMA 命令の順序を示すため、このプロデューサー/コンシューマー パイプラインは CTA を介して移動し続けます。このパイプラインの深さは 2 つの K 並列処理ノブとは別のものです。SPLIT_K は K 全体に CTA を追加し、BLOCK_K_WARPS は CTA の K タイルをワープ グループ全体に分割します。
B_TO_LDS が無効な場合、パイプラインは狭くなります。A は引き続き LDS を介してステージングされますが、B フラグメントは、ステージングされた LDS リングに参加するのではなく、グローバル メモリからレジスタに直接ロードされます。
単一起動 Split-K 同期 #
Inter-CTA Split-K では、複数の CTA が同じ出力タイルに寄与するという正確性の問題が発生します。
このカーネルは、2 つのグローバル バッファを持つ軽量のグローバル同期プロトコルを使用します。
信号[]
セマフォ[]
流れは次のとおりです。
最初の Split-K パーティションは出力タイルを初期化します。
バイの場合

[切り捨てられた]

## Original Extract

Learn how FlyDSL low-latency GEMMs speed up LLM decode on AMD GPUs with Split-K, K-slice parallelism, and an LDS-based pipeline.

Skip to main content
Back to top
Ctrl + K
ROCm™ Blogs
Home
Accelerating LLM Inference on AMD GPUs with Low-Latency GEMMs
Why Does Decode Latency Matter for LLM Serving?
Why Do Small- M , Large- K GEMMs Underperform?
Decode GEMM Shapes in Real Models
The Core Idea: LDS-Pipelined Split-K
Inter-CTA Split-K: More CTAs for Small-M GEMM
Intra-CTA K-slice Splitting: More Warp-Group Parallelism
Multi-Stage LDS Pipeline: Keep K Blocks in Flight
Single-Launch Split-K Synchronization
LDS Reduction for Intra-CTA K-slice Splitting
Implementation Notes: From Algorithm to FlyDSL Kernel
Accelerating LLM Inference on AMD GPUs with Low-Latency GEMMs #
June 29, 2026 by Yutao Xu, Xiaobing Zhang, Hattie Wu , Felix Li, Lingpeng Jin , Carlus Huang , Peng Sun , Barsoum Emad.
11 min read. | 2577 total words.
Software tools & optimizations
AI/ML , Performance
User , Developers , AI
Yutao Xu, Xiaobing Zhang, Hattie Wu , Felix Li, Lingpeng Jin , Carlus Huang , Peng Sun , Barsoum Emad
English
-->
Large language model inference is becoming increasingly interactive. Users expect chatbots, coding assistants, agents, and real-time copilots to respond quickly, stream tokens smoothly, and stay responsive under concurrent load. In that setting, decode-time latency is not just a backend metric. It directly affects perceived quality.
In this blog, you will explore one small but important part of that inference path: decode-time GEMMs with small M , large N and K , BF16/FP16 inputs, optional bias, and shapes that repeat across real models . These shapes can leave conventional GEMM tiling underutilized, which makes them a useful target for decode-path optimization.
The main technique is LDS-Pipelined Split-K GEMM : the long K reduction is split across CTAs, further sliced across warp groups inside each CTA, and kept moving through a multi-stage LDS memory pipeline. On AMD GPUs, LDS means Local Data Share, the on-chip scratchpad memory used for fast cooperation inside a CTA.
You will also see how we implement this idea as an AITER FlyDSL kernel family. FlyDSL keeps low-level ROCm™ software details such as MFMA selection, LDS layout, async copies, and synchronization explicit, while still generating shape-specialized variants for the model dimensions that appear in decode. In benchmark sweeps, this targeted decode optimization reaches a 1.64x average latency improvement over the fastest of HipblasLT , AITER Triton , and AITER ASM on the K = 7168 decode grid [ 1 ] , and a 1.49x average latency improvement on additional BF16 model-shape tests.
Why Does Decode Latency Matter for LLM Serving? #
LLM serving has two broad phases:
Prefill , where the model processes the prompt.
Decode , where the model generates output tokens one step at a time.
Prefill often has a larger effective M because many prompt tokens can be processed together. Decode is different. Each step may only process a small number of active tokens, especially after batching, scheduling, tensor parallelism, and request-level dynamics are taken into account.
That makes decode performance important for user-facing latency:
Time to first token affects how quickly the system appears to respond.
Time per output token affects streaming smoothness.
Inter-token latency affects whether the interaction feels fluid.
Throughput under concurrency affects how many users can be served without hurting responsiveness.
Figure 1 illustrates this interactive decode serving setting and shows where these latency concerns appear in the user-facing path.
Figure 1: Interactive LLM decode serving.
For these workloads, shaving overhead from repeated decode GEMMs can matter at the model-serving level.
Why Do Small- M , Large- K GEMMs Underperform? #
In large-model decode, GEMM often looks like:
C[M, N] = A[M, K] @ B[N, K]^T
Visually, the kernel still starts from the standard GEMM idea: compute a tile of C from a tile of A and a tile of B . The problem is that a small M produces too few output tiles, even though the K dimension can be very long.
Figure 2 shows this small- M , large- K bottleneck: the output grid is narrow, while the reduction dimension still contains substantial work.
Figure 2: Small- M , large- K GEMM bottleneck.
where M is the number of active tokens in a decode step or micro-batch. For serving workloads, M is frequently small: 1 , 2 , 4 , 8 , 16 , 32 , sometimes up to 128 or 256 . At the same time, N and K are model-hidden-size dimensions and can be thousands or tens of thousands.
That shape regime is awkward for general GEMM libraries. A conventional large-tile GEMM wants enough M x N work per block to keep all compute units busy. Decode GEMM often does not provide that naturally. The result is under-occupancy, poor wave utilization, and too much overhead relative to useful math.
Common GEMM optimizations such as larger CTA tiles, better memory coalescing, LDS staging, MFMA-focused scheduling, and pipelining still matter. But they do not by themselves create enough independent work when the M x N output grid is small. This is why LDS-Pipelined Split-K combines multiple forms of K parallelism instead of relying on one optimization layer.
Decode GEMM Shapes in Real Models #
The motivation came directly from model shape traces, not from synthetic square GEMMs.
Across current LLMs, decode GEMM shapes repeatedly show the same pattern:
M = 1–256 , N = 256 / 2112 / 3072 / 7168 / 16160 , K = 1536 / 2048 / 7168
M = 1–256 , N = 128 / 640 / 2560 / 2880 / 5120 , K = 512 / 2048 / 2880 / 4096
M = 1–256 , plus prefill-like large M , with N and K around 6144 , 4096 , 2048 , etc.
many skinny decode shapes such as M = 8–512 , N = 384 or 1024 , K = 7168
M = 1–32768 , but decode-critical cases include M = 1/16/32/64 , with large N and K
M = 1–32768 , with decode-heavy skinny projections such as N = 100/200/800 , K = 5120
The important observation is not just “small M exists.” It is that small- M , large- K GEMMs occur everywhere in decode paths , and they affect end-to-end serving throughput.
So the design target is a low-overhead GEMM path with:
Good occupancy even when M is tiny
The Core Idea: LDS-Pipelined Split-K #
The kernel treats K as splittable reduction work rather than a private serial loop inside one CTA. For one output tile C[m_tile, n_tile] , the computation is:
C_tile =
A[m_tile, K0] @ B^T[K0, n_tile]
+ A[m_tile, K1] @ B^T[K1, n_tile]
+ A[m_tile, K2] @ B^T[K2, n_tile]
+ ...
LDS-Pipelined Split-K exposes those K0/K1/K2/... chunks at three levels:
Inter-CTA Split-K : split the full K dimension across multiple CTAs (workgroups).
Intra-CTA K-slice splitting : split the K tile of one CTA across multiple warp groups inside the block.
Multi-stage LDS pipeline : pipeline K blocks through ring-buffered LDS stages while overlapping global-to-LDS copies, LDS reads, and MFMA compute.
These layers solve different problems.
Better GPU occupancy when M x N has too few tiles
Global accumulation and synchronization
Better use of warp groups for K-heavy tiles
LDS staging and local reduction
Overlap global-to-LDS copies, LDS reads, and MFMA while K blocks advance
Ring-buffered LDS stages and scheduling
More work across the GPU, more useful work per CTA, and smoother K-block pipelining
A coordinated reduction and pipeline path
Figure 3 shows the full tiled data path for this design. A selected C tile is not produced by one monolithic CTA. The long K dimension is first broken into inter-CTA Split-K ranges, each CTA streams its assigned K blocks through the multi-stage LDS pipeline, and intra-CTA K-slice splitting lets multiple warp groups compute partial accumulations for the same output tile. Those local partials are reduced through LDS before the inter-CTA Split-K partials are accumulated into the final C tile.
Figure 3: LDS-Pipelined Split-K data path.
Inter-CTA Split-K: More CTAs for Small-M GEMM #
When M is small, the normal M x N tile grid may not launch enough CTAs to saturate the GPU.
Inter-CTA Split-K expands the launch grid along K:
grid = [mn_tiles, split_k]
Each Split-K partition computes a partial sum over a different K range. After that partition finishes its local pipeline work and intra-CTA LDS reduction, the partial result is accumulated into the same output tile.
In the launch wrapper, inter-CTA Split-K is visible as the second grid dimension:
bm = ( m + BLOCK_M - 1 ) // BLOCK_M
hgemm_kernel ( C , A , B , BIAS , m , semaphore , signal ) . launch (
grid = ( bm * N_BLOCKS , SPLIT_K , 1 ),
block = ( BLOCK_THREADS , 1 , 1 ),
stream = stream ,
)
This is especially useful for decode shapes like:
M = 1, 2, 4, 8, 16
N = 2560 / 2880 / 5120
K = 2880 / 4096 / 7168
Without this extra Split-K dimension, there may simply not be enough independent work.
Intra-CTA K-slice Splitting: More Warp-Group Parallelism #
Inter-CTA Split-K increases the number of CTAs. Intra-CTA K-slice splitting increases useful work inside one CTA.
The kernel assigns multiple warp groups to different K slices of the same tile. Each group computes a partial accumulation. At the end of the CTA, those partial results are reduced through LDS before writing back.
It increases parallelism for K-heavy tiles.
It controls register pressure by distributing work across warp groups.
Multi-Stage LDS Pipeline: Keep K Blocks in Flight #
The third layer is temporal. Once a CTA owns a K range, it still has to repeatedly compute:
C_tile += A[m_tile, K_i] @ B^T[K_i, n_tile]
for many consecutive K_i blocks. Instead of treating those blocks as a serial load-then-compute sequence, the kernel uses STAGES as a ring buffer of LDS tiles. A stage that was filled earlier is consumed by LDS reads and MFMA, while another stage is reused for a future global-to-LDS copy.
In the B_TO_LDS path, both A and B participate in this LDS ring. The prologue first prefetches STAGES - 1 K blocks, and the hot loop then consumes one stage while issuing the copy for a future stage:
for s in range_constexpr ( STAGES - 1 ):
ldg_sts_b_async ( ks_begin + s * BLOCK_K , s )
ldg_sts_a_async ( ks_begin + s * BLOCK_K , s )
for bki , state in range ( 0 , BLOCK_K_LOOPS - ( STAGES - 1 ), 1 , init = init_state ):
k_offset = state [ 0 ]
current_stage = fx . Index ( state [ 1 ])
next_stage = ( current_stage + 1 ) % STAGES
write_stage = ( current_stage + STAGES - 1 ) % STAGES
__barrier (( STAGES - 2 ) * LDG_WAIT_COUNT )
ldg_sts_b_async ( k_offset + ( STAGES - 1 ) * BLOCK_K , write_stage )
ldg_sts_a_async ( k_offset + ( STAGES - 1 ) * BLOCK_K , write_stage )
c_frags_new = ldmatrix_compute_tile_streaming ( current_stage , c_frags )
hot_loop_scheduler ()
The hot loop then advances the ring one K block at a time. current_stage is the LDS stage being consumed, and write_stage = current_stage + STAGES - 1 modulo STAGES is the stage receiving the future K block. The wait count intentionally leaves newer copies outstanding:
__barrier (( STAGES - 2 ) * LDG_WAIT_COUNT )
That means the loop waits for the current stage to be safe to read, without draining every in-flight global-to-LDS copy. Conceptually:
current stage : LDS reads feed MFMA for K block i
write stage : global-to-LDS copy brings K block i + STAGES - 1
next stage : becomes current in the next loop iteration
The scheduler hints in hot_loop_scheduler() order VMEM, LDS reads, and MFMA instructions so this producer/consumer pipeline keeps moving through the CTA. This pipeline depth is separate from the two K-parallelism knobs: SPLIT_K adds CTAs across K, while BLOCK_K_WARPS splits a CTA’s K tile across warp groups.
When B_TO_LDS is disabled, the pipeline is narrower: A is still staged through LDS, but B fragments are loaded directly from global memory into registers instead of joining the staged LDS ring.
Single-Launch Split-K Synchronization #
Inter-CTA Split-K creates a correctness problem: multiple CTAs contribute to the same output tile.
This kernel uses a lightweight global synchronization protocol with two global buffers:
signal[]
semaphore[]
The flow is:
The first Split-K partition initializes the output tile.
If bi

[truncated]
