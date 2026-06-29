---
source: "https://www.phantafield.com/whitepaper"
hn_url: "https://news.ycombinator.com/item?id=48713686"
title: "Sophon PFG-1: a monolithic-3D AI ASIC with 330 GB of on-die DRAM and no HBM"
article_title: "PhantaField PFG-1 Whitepaper"
author: "minkowsky"
captured_at: "2026-06-29T01:42:22Z"
capture_tool: "hn-digest"
hn_id: 48713686
score: 3
comments: 3
posted_at: "2026-06-29T01:23:38Z"
tags:
  - hacker-news
  - translated
---

# Sophon PFG-1: a monolithic-3D AI ASIC with 330 GB of on-die DRAM and no HBM

- HN: [48713686](https://news.ycombinator.com/item?id=48713686)
- Source: [www.phantafield.com](https://www.phantafield.com/whitepaper)
- Score: 3
- Comments: 3
- Posted: 2026-06-29T01:23:38Z

## Translation

タイトル: Sophon PFG-1: 330 GB のオンダイ DRAM を搭載し、HBM を持たないモノリシック 3D AI ASIC
記事のタイトル: PhantaField PFG-1 ホワイトペーパー
説明: PhantaField PFG-1 Sophon — 2D-TMD モノリシック 3D AI シリコン。リビジョン 4.1、2026 年 6 月。

記事本文:
PhantaField PFG-1 ホワイトペーパー テクノロジー
PhantaField PFG-1 Sophon ホワイトペーパー
PFG-1「Sophon」は、750 mm²、32 層 2D 上の統合されたトレーニングおよび推論ダイです。
遷移金属ダイカルコゲニド (TMD) モノリシック 3D (M3D) プラットフォーム。重み、勾配、およびオプティマイザーの状態
オンダイ 2T0C 2D-TMD ゲインセル DRAM に常駐します。アレイは完全に読み書き可能であるため、同じシリコンが実行されます。
BF16 の前方/後方トレーニングはパスし、コンピューティング制限レートで低バッチ デコードを提供します。
Compute は純粋なデジタル Compute-In-Memory (CIM) です。各 256×256 DRAM サブアレイ タイル ペア
8 レベル加算器ツリーを備えたバイナリ センス アンプで、500 MHz ビット シリアル アクティベーション ブロードキャストによって駆動されます。で
131,072 タイル/ダイにより、FP8 では 4,200 TFLOPS、BF16 では 2,100 TFLOPS が得られます。
設置面積は 7.5 cm²。
ダイは、28 nm Si 相補型金属酸化物半導体 (CMOS) ベース層、32 層 2D-TMD CMOS 上に構築されています。
MAC スタック、および 2T0C DRAM モジュールを備えたモノリシック層間ビア (MIV) ファブリック [5] [6] [7]
各メモリ層の Back-End-Of-Line (BEOL) Metal-3 層に埋め込まれます。ダイスタックの断面は
図 1 に示します。
Sophon は、オフダイの高帯域幅メモリ (HBM) を完全に排除します。 80B パラメータの BF16 トレーニングに適合します
重み + 一次オプティマイザの状態は完全にオンダイで、最大 10 GB のアクティベーション ヘッドルームを備えています。
勾配チェックポイント付きマイクロバッチ。推論では、80B モデルを提供します。
ネイティブ BF16 では 7,219 トークン/秒、または FP8 モードでは完全な 14,438 トークン/秒 —
これを単一の「トレーニングしてからサーブする」部分にし、トレーニングとサーブの間で弾力的に再分割できるようにします。
ハードウェアを変更することなく。 NVIDIA Rubin (R200) と AMD Instinct MI455X (どちらも 2026 HBM4 パーツ) に対して、Sophon は優れたパフォーマンスを発揮します
ダイごとの 80B バッチ 1 トレーニング スループットは約 2.7 ～ 3.1 倍、約 48 ～ 53 倍向上
シングルストリーム FP8 デコード

低バッチ時の両方の GPU が HBM4 で HBM 帯域幅に制限されているため、スループットが低下します。
制限 (Rubin 22 TB/秒、MI455X 19.6 TB/秒)。ピーク密度 FLOPS は GPU に有利 — Sophon BF16 密度はわずか約 0.21 ～ 0.24 倍です
ただし、ピーク FLOPS は、重みメモリ帯域幅が支配する低バッチでは役に立ちません。
このアーキテクチャは、HBM4 パッケージの約 191 ～ 214 倍の重量帯域幅を実現します (Rubin に対して 191 倍、
214× vs MI455X) — HBM ロードマップでは埋められないギャップ (セクション 7)。
経済性はそのまま続きます。モルガン・スタンレーは、単一の NVIDIA VR200 (Rubin) NVL72 ラックを
≈ 780 万ドル — HBM メモリ単体で ≈ 200 万ドル (ラックの 25.7%、GB300 より +435%)。ソフォン
その項目を削除すると、Rubin / MI455X よりもハードウェア BOM が約 9.9 倍 / 11.6 倍低くなります
[17] 。
アーキテクチャの概要
A. プラットフォーム (ダイ、ティア、MIV、TMD MAC)
B. PFG-1「Sophon」 — 2T0C DRAM ダイ
C. ダイのフロアプランとオンダイのシステム構成
物理計算
A. セルの形状と層ごとの密度
C. MAC ごとのエネルギーと電力エンベロープ
D. デジタル CIM タイルの物理学と 1/N スケーリング
GPU アーキテクチャと AI パフォーマンス
A. 推論
モデルサイズにおけるエネルギー制約の上限
推論（サービング）上限
宇宙用途の放射線耐性
検証、リスク、今後の取り組み
最新の AI アクセラレータは、処理する必要がある次の両方のワークロードでメモリの壁に直面しています。
推論は読み取り主導型です。モデルの重みは展開時に固定されます。すべてのデコード
step は、生成されたトークンごとに 1 回、全重みテンソルを読み取ります。重要な指標は、ビットあたりの読み取りエネルギー、アイドル状態です。
リーク (モデルはリクエスト間で常駐する必要がある)、および低バッチでの重みフェッチ帯域幅。従来の
高帯域幅メモリ (HBM) は、低バッチでは帯域幅が制限されます。すべてのトークンの MAC トラフィックは、
~ 22 TB/秒 (Rubin) / 19.6 TB/秒 (MI455X) HBM4 パス、および 288 ～ 432 GB HBM4 サブシステムの消費電力は ~ 10 ～ 15 W (秒)

モデルを保持するためだけに elf-refresh
居住者。
トレーニングは読み書き対称です。すべての順方向パスで重みが読み取られます。すべての後ろ向き
pass はグラデーション更新を書き込みます。オプティマイザはステップごとに適切な重みを更新します。インプレース書き込み可能性、低い
書き込みエネルギー、および重みとオプティマイザー状態の両方の容量が重要です。不揮発性
推論のみのメモリはトレーニングできません。たとえば、シングルレベル セル (SLC) の抵抗性 RAM の耐久性の上限は約 10⁶ です。
一方、80B モデルのトレーニングにはパラメーターごとに約 10¹⁰ の書き込みサイクルが必要です。
2T0C 2D-TMD ゲインセル DRAM は、1 つのセルで両方の問題を解決します。異常に悪用します
TMD の低いオフ電流密度 (J off ≈ 10⁻¹5 A/μm = 28 nm で 1 fA/μm、つまりセルあたり ≈ 0.5 fA)
トランジスタを使用して、明示的なストレージ コンデンサを使用せずに数秒間の保持を実現し、
無制限の書き込み耐久性とリフレッシュ オーバーヘッドを備えた 20 fJ/ビットでのインプレース グラデーション書き込み
わずか ≈ 0.08 W です。ストレージ ノードはサイクルごとに書き込み可能であるため、推論に使用される同じダイが
トレーニングもする。保持期間が数秒であるため、アイドル電力は最大 3 W に低下します。これは推論グレードのアイドル プロファイルです。
完全に書き込み可能なトレーニング ダイ上で。
PhantaField の 2D-TMD M3D プラットフォームは、この DRAM モジュールを各メモリ層の BEOL Metal-3 層に統合します。
MAC アレイが重みを消費する論理層のすぐ上にあります。
Sophon は次の物理スタックを使用します。
スタックの合計高さ: Si ダイ上約 22 µm (64 段 × 0.35 µm/段)。 90nmピッチMIV
グリッドは、1.23 × 10⁸ スロット/mm² の利用可能な層間接続を提供します。設計に実装されるのは最大 5.5 ×
10⁵/mm²、99% 以上の MIV ヘッドルームを残します。
階層は 1 つのレイヤー内で分割されません。代わりに 64 層スタック
専用のロジック層とメモリ層を A/B/A/B… の繰り返しパターンでインターリーブします。隣接する 2 つの
の階層

m 1 つのロジックとメモリのダブレット。スタックには、そのようなダブレットが 32 個含まれています。
論理層 (32 × 750 mm² = 合計 MAC エリア 24,000 mm²): 2D-TMD CMOS MAC アレイを搭載
奇数インデックスの層 - NMOS の場合は MoS₂ n-FET、PMOS の場合は WSe₂ p-FET。密度 0.175 TFLOPS FP8/mm² (0.0875
TFLOPS BF16/mm²)。 1.2 GHz でクロック、V dd = 0.6 V。
メモリ層 (32 × 750 mm² = 合計メモリ領域 24,000 mm²): 2T0C 2D-TMD DRAM オン
偶数インデックスの層。その層の Metal-3 BEOL で製造されます。各メモリ層はその直上に位置します。
ペアになった論理層。サブ 100 nm ピッチの垂直モノリシック層間ビア (MIV)
ビット線/ワード線/センス信号がロジック MAC アレイからセルに直接送られ、すべての MAC にその信号が与えられます。
独自のプライベート垂直ポートをローカルの重み付けに使用し、NoC トラフィックをゼロにします。このインターリーブ配置により、
仮想的な層内 50/50 分割と同じ合計エリアと容量でありながら、層ごとの MAC は 2 倍になります。
ルーティング領域を拡大し、MAC からセルへの信号パスを 0.35 µm の単一層ピッチに短縮します。
なぜ 2D TMD なのか? TMD CMOS (MoS₂ / WSe₂) は、同時に
(1) ≤ 450 °C での BEOL 互換の成長 [6] 。 (2) 原子スケール
チャネルの厚さにより、短チャネルのリークが排除されます [1] [2]。 (3) 電子移動度 ≥ 120 cm²/V・s
[4] ; (4) 固有の放射線耐性 (埋め込み酸化物トラップ体積なし)。
重要なことに、TMD オフ電流密度 J off は 28 nm で ≈ 10⁻¹5 A/μm (1 fA/μm)、つまり ≈ 0.5 fA です。
幅0.5μmのセルトランジスタで、同等のゲート長のSi NMOSよりもおよそ4桁小さい
[2] [3] — 2T0C セルを可能にするものです。
蓄積コンデンサを使用せずに数秒間データを保持 [8] [9] 、セル面積を 1 回の充電に必要な ~20 F² ではなく 8 F² に維持します。
従来の1T1C DRAM。
B. PFG-1「Sophon」 — 2T0C DRAM ダイ
ソフォ

n は、2T0C 2D-TMD ゲインセル DRAM (8 F²、1 ビット/セル) をそれぞれの Metal-3 BEOL に配置します。
メモリ層。セル構造は図 2 に示されており、次のもので構成されます。
書き込みトランジスタ (WT): 書き込みワードライン (WWL) によってゲート制御される TMD nFET。
ストレージノードをV dd に下げるか、GNDに放電します。
読み取りトランジスタ (RT): ゲートがストレージ ノードである TMD nFET。そのドレイン電流
は格納されているビットを示します。
ストレージ ノード: RT の寄生ゲート容量 (28 nm TMD で約 2.5 fF) と
WT のドレインの接合容量 (~0.5 fF)。明示的な金属 - 絶縁体 - 金属 (MIM) またはトレンチ コンデンサはありません
— それは 2T0C の「0C」です。
TMD オフ電流密度 1 fA/μm (0.5 μm セル トランジスタの場合、I off ≈ 0.5 fA) は次のようになります。
保持時間 τ = C・V dd / (2・I off ) = 1.8 秒 (25 °C)
[8] [9] — 式を参照。 3と
図 3 は保持曲線です。 Sophon は 1.0 秒ごとに更新します (1.8x
マージン)、330 GB ダイ全体でわずか約 0.08 W しか消費しません (式 4)。
保持力は 10 °C ごとに約 2 倍低下します。接合部温度が 60 °C を超えると、オンダイ熱センサーにより温度が短くなります。
リフレッシュ間隔 (60 °C で ≈ 159 ms、85 °C で ≈ 28 ms)、リフレッシュ電力は、
ホットコーナー。
ストレージ ノードはサイクルごとに書き込み可能であるため、Sophon はインプレース BF16 勾配蓄積をサポートします。
無制限の耐久性 - まさにトレーニングに必要なもの - 同じ配列 (読み取り専用) が
推論デコード ループ。ダイはモデルを一度ロードし、それを提供 (推論) するか、その場で更新します。
(トレーニング);電源オフのダイは、起動時にオフダイの不揮発性メモリ エクスプレス (NVMe) から重みをリロードします。
(§11.2)。
C. ダイのフロアプランとオンダイのシステム構成
131,072 個の CIM タイルはフラット アレイではなく、スタックの 32 の論理層に分割されています。
(§2.A)、論理層ごとにちょうど 4,096 タイル (

導出: 131,072 ÷ 32)。各タイルは、
層上の固定セルであり、コンピューティング、ストレージ、冗長性の原子単位です。256×256 の重みサブアレイです。
(65,536 重み) バイナリ センス アンプと 8 レベル加算器ツリーにビット シリアル アクティベーション ブロードキャストを供給
500 MHz (BF16 16 サイクル、FP8 8 サイクル)。すべてのタイルの重みはメモリの 2T0C セルに存在します。
その直上に層があるため (§2.B)、タイルは物理的には平面ではなく垂直のロジック プラス メモリ列です。
ブロック。したがって、層はこれらの列の 4,096 タイル メッシュです。完全なダイは、このようなメッシュが 32 個積み重ねられたものです。
0.35 µm ピッチ。その下の 28 nm Si ベースには、計算以外のすべてが含まれます。
NoC は階層ごとの 2D メッシュであり、グローバル ファブリックではありません。各論理層は独自のメッシュを実行します
ルータ ファブリックの二等分速度は約 290 TB/秒であり、合わせて 64 層が存在します。
合計 18,560 TB/秒 (導出: 290 × 64)。 NoC に乗るものは意図的に最小限に抑えられています。
アクティベーションと部分和 — タイルを組み立てるためにタイル間を移動する必要があるオペランド
4,096 タイルのファンインにわたるレイヤーの出力。ウェイトが NoC に触れることはありません。それぞれの重さは
タイルのプライベート垂直 MIV ポートを介して読み取ります。セルから直接下への単一層ピッチ ホップです。
その MAC — 共有バス競合なしで 4.2 PB/s のタイル内重み帯域幅を提供します (§2.A)。これは
フロアプランの耐荷重の非対称性: マルチペタバイトのトラフィック (ウェイトフェッチ) が完全に保持されます。
垂直かつローカルであるため、横方向の NoC は比較的小さな活性化/部分和のみを伝送します。
フラックス。基本層の NoC ルートは、層ごとのメッシュを縫い合わせて、それらを
コントローラーとホスト I/O に接続されますが、ウェイト パスには含まれません。
各タイルはさらに、アクティベーション用の小さな SRAM スクラッチパッドを所有します。 NoCだから
アクティベーションとパーシャルを保持します。

ハンの重み、スクラッチパッドはタイルがその受信をステージングする場所です
アクティベーション ベクトルを生成し、ビット シリアル ブロードキャスト全体で部分和のスライスを蓄積し、
メッシュに渡される前の送信結果。ライブ アクティベーション ワーキング セットを高速ローカル SRAM に保持する
— 2T0C DRAM 内ではなく加算器ツリーに隣接 — ブロードキャスト/累積内部ループを完全に保持
タイル上に配置し、1 Hz リフレッシュのゲインセル DRAM (§2.B) を重みと KV キャッシュ専用にし、そのアクセスに使用します。
それに比べて、パターンは読み取りがほとんどで、レイテンシ耐性があります。
クロックと電源は、22 µm スタックを介して低電圧レールに供給されます。論理層
MIV グリッドを通じて上向きに分散された基本層クロック ルートから 1.2 GHz でクロックされます。
ビット シリアル アクティベーション ブロードキャストは、別の 500 MHz ドメインで実行されます。営業時間
V dd = 0.6 V により、64 層モノリシック スタックが熱的に実行可能になり、ダイナミックになります。
電力は V dd ² に応じて変化するため、0.6 V レールの消費エネルギーは公称 1.0 V CMOS レールよりも約 2.8 倍少なくなります。
同じアクティビティで。取引は現在のものです。固定電力では、電圧を下げると供給電流が増加します。
そしてその電流は、最大約 22 μm の電力供給ネットワーク (PDN) を介してすべての層に到達する必要があります。
積み重ねる。設計では MIV グリッドの 99% 以上がシグナリングに未使用のまま残されているため (§2.A)、これらの予備ビアは
PDN (導出) に割り当てられる — キャリー経由の並列 V dd /GND

[切り捨てられた]

## Original Extract

PhantaField PFG-1 Sophon — 2D-TMD Monolithic 3D AI Silicon. Revision 4.1, June 2026.

PhantaField PFG-1 Whitepaper Technology
PhantaField PFG-1 Sophon Whitepaper
PFG-1 "Sophon" is a unified training-and-inference die on a 750 mm², 32-tier 2D
Transition-Metal Dichalcogenide (TMD) Monolithic 3D (M3D) platform. Weights, gradients, and optimizer state
reside in on-die 2T0C 2D-TMD gain-cell DRAM; because the array is fully read-write, the same silicon executes
BF16 forward/backward training passes and serves low-batch decode at the compute-bound rate.
Compute is pure digital Compute-In-Memory (CIM) : each 256×256 DRAM subarray tile pairs a
binary sense amplifier with an 8-level adder tree, driven by a 500 MHz bit-serial activation broadcast. At
131,072 tiles/die this yields 4,200 TFLOPS FP8 and 2,100 TFLOPS BF16 in a
7.5 cm² footprint.
The die is built on a 28 nm Si Complementary Metal-Oxide-Semiconductor (CMOS) base tier, a 32-tier 2D-TMD CMOS
MAC stack, and a Monolithic Inter-tier Via (MIV) fabric [5] [6] [7] , with the 2T0C DRAM module
embedded at the Back-End-Of-Line (BEOL) Metal-3 layer of each memory tier. The die stack cross-section is
shown in Figure 1 .
Sophon eliminates off-die High-Bandwidth Memory (HBM) entirely. For 80B-parameter BF16 training it fits
weights + first-order optimizer state fully on-die with ~ 10 GB of activation headroom for
gradient-checkpointed micro-batches; for inference it serves an 80B model at
7,219 tokens/s in native BF16 or the full 14,438 tokens/s in FP8 mode —
making it a single train-then-serve part that can be elastically repartitioned between training and serving
without changing hardware. Against an NVIDIA Rubin (R200) and an AMD Instinct MI455X — both 2026 HBM4 parts — Sophon delivers
~ 2.7–3.1× higher 80B batch-1 training throughput per die and ~ 48–53× higher
single-stream FP8 decode throughput, because both GPUs at low batch are HBM-bandwidth-bound at their HBM4
limits (Rubin 22 TB/s, MI455X 19.6 TB/s). Peak dense FLOPS favor the GPUs — Sophon BF16 dense is only ~ 0.21–0.24×
their peak — but peak FLOPS do not help at low batch, where weight-memory bandwidth governs.
The architecture delivers ~ 191–214× the weight bandwidth of an HBM4 package (191× vs Rubin,
214× vs MI455X) — a gap no HBM roadmap closes (Section 7).
The economics follow directly: Morgan Stanley puts a single NVIDIA VR200 (Rubin) NVL72 rack at
≈ $7.8M — HBM memory alone ≈ $2.0M (25.7% of the rack, +435% over GB300). Sophon
eliminates that line item, for a ~ 9.9× / 11.6× lower hardware BOM than a Rubin / MI455X
[17] .
Architecture Overview
A. Platform (die, tiers, MIV, TMD MAC)
B. PFG-1 "Sophon" — 2T0C DRAM die
C. Die floorplan & on-die system organization
Physical Calculations
A. Cell geometry & per-tier density
C. Per-MAC energy & power envelope
D. Digital CIM tile physics & 1/N scaling
GPU Architecture & AI Performance
A. Inference
Energy-Constrained Ceiling on Model Size
Inference (serving) ceiling
Radiation Tolerance for Space Applications
Validation, Risks & Future Work
Modern AI accelerators face a memory wall on both workloads they must serve:
Inference is read-dominated . The model weights are fixed at deployment; every decode
step reads the full weight tensor once per generated token. The key metrics are read energy per bit, idle
leakage (the model must stay resident between requests), and weight-fetch bandwidth at low batch. Conventional
High-Bandwidth Memory (HBM) is bandwidth-bound at low batch: every token's MAC traffic serializes through the
~ 22 TB/s (Rubin) / 19.6 TB/s (MI455X) HBM4 path, and a 288–432 GB HBM4 subsystem draws ~ 10–15 W in self-refresh just to keep the model
resident.
Training is read-write symmetric . Every forward pass reads weights; every backward
pass writes gradient updates; the optimizer updates weights in place each step. In-place writability, low
write energy, and capacity for both weights and optimizer state are critical. A non-volatile
inference-only memory cannot train — for example, Single-Level Cell (SLC) Resistive RAM endurance caps at ~10⁶
cycles, while training an 80B model requires ~10¹⁰ write cycles per parameter.
A 2T0C 2D-TMD gain-cell DRAM solves both problems with one cell. It exploits the anomalously
low off-current density (J off ≈ 10⁻¹⁵ A/µm = 1 fA/µm at 28 nm, i.e. ≈ 0.5 fA per cell) of TMD
transistors to obtain multi-second retention without an explicit storage capacitor, enabling
in-place gradient writes at 20 fJ/bit with unlimited write endurance and a refresh overhead
of only ≈ 0.08 W. Because the storage node is writable on every cycle, the same die that serves inference can
also train; because retention is seconds-long, idle power collapses to ~ 3 W — an inference-grade idle profile
on a fully writable training die.
PhantaField's 2D-TMD M3D platform integrates this DRAM module at the BEOL Metal-3 layer of each memory tier,
directly above the logic tier whose MAC array consumes its weights.
Sophon uses the following physical stack:
Total stack height: ~22 µm above the Si die (64 tiers × 0.35 µm/tier). The 90 nm-pitch MIV
grid provides 1.23 × 10⁸ slots/mm² available inter-tier connections; the design populates only ~5.5 ×
10⁵/mm², leaving > 99% MIV headroom.
Tiers are not split within a single layer; instead the 64-tier stack
interleaves dedicated logic and memory tiers in an A/B/A/B… repeating pattern. Two adjacent
tiers form one logic-plus-memory doublet ; the stack contains 32 such doublets:
Logic tiers (32 × 750 mm² = 24,000 mm² total MAC area): 2D-TMD CMOS MAC array on
odd-indexed tiers — MoS₂ n-FETs for NMOS, WSe₂ p-FETs for PMOS. Density 0.175 TFLOPS FP8/mm² (0.0875
TFLOPS BF16/mm²). Clocked at 1.2 GHz, V dd = 0.6 V.
Memory tiers (32 × 750 mm² = 24,000 mm² total memory area): 2T0C 2D-TMD DRAM on
even-indexed tiers, fabricated at the Metal-3 BEOL of that tier. Each memory tier sits directly above its
paired logic tier; vertical Monolithic Inter-tier Vias (MIVs) on a sub-100 nm pitch carry
bit-line/word-line/sense signals straight up from the logic MAC array into the cells, giving every MAC its
own private vertical port to local weights with zero NoC traffic. This interleaved arrangement preserves
the same total area and capacity as a hypothetical in-tier 50/50 split, while doubling the per-tier MAC
routing area and shortening MAC-to-cell signal paths to a single tier-pitch of 0.35 µm.
Why 2D TMD? TMD CMOS (MoS₂ / WSe₂) is the only transistor technology that simultaneously
offers: (1) BEOL-compatible growth at ≤ 450 °C [6] ; (2) atomic-scale
channel thickness eliminating short-channel leakage [1] [2] ; (3) electron mobility ≥ 120 cm²/V·s
[4] ; and (4) intrinsic radiation hardness (no buried-oxide trap volume).
Critically, the TMD off-current density J off ≈ 10⁻¹⁵ A/µm (1 fA/µm) at 28 nm — i.e. ≈ 0.5 fA for
a 0.5 µm-wide cell transistor, roughly 4 orders of magnitude lower than Si NMOS at equivalent gate length
[2] [3] — is what enables a 2T0C cell to
retain data for seconds without any storage capacitor [8] [9] , keeping the cell area at 8 F² rather than the ~20 F² needed for a
conventional 1T1C DRAM.
B. PFG-1 "Sophon" — 2T0C DRAM die
Sophon places a 2T0C 2D-TMD gain-cell DRAM (8 F², 1 bit/cell) at the Metal-3 BEOL of each
memory tier. The cell structure is shown in Figure 2 and consists of:
Write Transistor (WT): a TMD nFET gated by the Write Word-Line (WWL), which charges the
storage node to V dd or discharges it to GND.
Read Transistor (RT): a TMD nFET whose gate is the storage node; its drain current
indicates the stored bit.
Storage node: the parasitic gate capacitance of RT (~2.5 fF at 28 nm TMD) plus the
junction capacitance of WT's drain (~0.5 fF). No explicit Metal-Insulator-Metal (MIM) or trench capacitor
— that is the "0C" in 2T0C.
The TMD off-current density of 1 fA/µm (I off ≈ 0.5 fA for a 0.5 µm cell transistor) gives
retention τ = C·V dd / (2·I off ) = 1.8 s at 25 °C
[8] [9] — see Eq. 3 and
Figure 3 for the retention curve. Sophon refreshes every 1.0 s (1.8×
margin), consuming only ≈ 0.08 W for the full 330 GB die ( Eq. 4 ).
Retention derates ≈ 2× per 10 °C; above 60 °C junction temperature, on-die thermal sensors shorten the
refresh interval (≈ 159 ms at 60 °C, ≈ 28 ms at 85 °C), with refresh power staying below ~ 4 W even in the
hot corner.
Because the storage node is writable on every cycle, Sophon supports in-place BF16 gradient accumulation
with unlimited endurance — exactly what training requires — while the same array, read-only, serves the
inference decode loop. The die loads a model once and either serves it (inference) or updates it in place
(training); a powered-off die reloads its weights from off-die Non-Volatile Memory express (NVMe) at boot
(§11.2).
C. Die Floorplan & On-Die System Organization
The 131,072 CIM tiles are not a flat array — they are partitioned across the 32 logic tiers of the stack
(§2.A), exactly 4,096 tiles per logic tier (derived: 131,072 ÷ 32). Each tile occupies a
fixed cell on its tier and is the atomic unit of compute, storage, and redundancy: a 256×256 weight subarray
(65,536 weights) feeding a binary sense amp and an 8-level adder tree, with bit-serial activation broadcast
at 500 MHz (16 cycles BF16, 8 cycles FP8). The weights for every tile live in the 2T0C cells of the memory
tier directly above it (§2.B), so a tile is physically a vertical logic-plus-memory column, not a planar
block. A tier is therefore a 4,096-tile mesh of these columns; the full die is 32 such meshes stacked at
0.35 µm pitch, with the 28 nm Si base below carrying everything that is not compute.
The NoC is a per-tier 2D mesh, not a global fabric. Each logic tier runs its own mesh
router fabric at ≈ 290 TB/s bisection, and the 64 tiers together present
18,560 TB/s aggregate (derived: 290 × 64). What rides the NoC is deliberately minimal:
activations and partial sums — the operands that must move between tiles to assemble a
layer's output across the 4,096-tile fan-in. Weights never touch the NoC. Every weight is
read through its tile's private vertical MIV port — a single tier-pitch hop straight down from the cell to
its MAC — delivering 4.2 PB/s of in-tile weight bandwidth with zero shared-bus contention (§2.A). This is
the load-bearing asymmetry of the floorplan: the multi-petabyte traffic (weight fetch) is kept entirely
vertical and local, so the lateral NoC only ever carries the comparatively small activation/partial-sum
flux. The base-tier NoC root stitches the per-tier meshes together and bridges them to the
controller and host I/O, but it is never in the weight path.
Each tile additionally owns a small SRAM scratchpad for activations. Because the NoC
carries activations and partials rather than weights, the scratchpad is where a tile stages its inbound
activation vector, accumulates its slice of the partial sum across the bit-serial broadcast, and buffers the
outbound result before it is handed to the mesh. Holding the live activation working set in fast local SRAM
— adjacent to the adder tree, not in the 2T0C DRAM — keeps the broadcast/accumulate inner loop entirely
on-tile and lets the 1 Hz-refresh gain-cell DRAM (§2.B) stay dedicated to weights and KV cache, whose access
pattern is read-mostly and latency-tolerant by comparison.
Clock and power are delivered down the 22 µm stack to a low-voltage rail. The logic tiers
are clocked at 1.2 GHz from a base-tier clock root distributed upward through the MIV grid;
the bit-serial activation broadcast runs on a separate 500 MHz domain. Operating at
V dd = 0.6 V is what makes a 64-tier monolithic stack thermally viable — dynamic
power scales with V dd ², so the 0.6 V rail draws ≈ 2.8× less energy than a nominal 1.0 V CMOS rail
at the same activity. The trade is current: at fixed power, lowering the voltage raises the supply current,
and that current must reach every tier through a power-delivery network (PDN) that climbs the full ~22 µm of
stack. Because the design leaves > 99% of the MIV grid unused for signaling (§2.A), those spare vias can
be allocated to the PDN (derived) — parallel V dd /GND vias carrie

[truncated]
