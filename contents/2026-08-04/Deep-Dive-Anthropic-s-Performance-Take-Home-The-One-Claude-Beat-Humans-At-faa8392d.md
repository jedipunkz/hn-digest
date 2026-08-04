---
source: "https://trirpi.github.io/posts/anthropic-performance-takehome/"
hn_url: "https://news.ycombinator.com/item?id=49164622"
title: "Deep Dive: Anthropic's Performance Take-Home (The One Claude Beat Humans At)"
article_title: "Deep Dive: Anthropic's Performance Take-Home (The One Claude Beat Humans At) | Tristan Trouwen"
author: "jxmorris12"
captured_at: "2026-08-04T06:25:35Z"
capture_tool: "hn-digest"
hn_id: 49164622
score: 1
comments: 0
posted_at: "2026-08-04T05:25:33Z"
tags:
  - hacker-news
  - translated
---

# Deep Dive: Anthropic's Performance Take-Home (The One Claude Beat Humans At)

- HN: [49164622](https://news.ycombinator.com/item?id=49164622)
- Source: [trirpi.github.io](https://trirpi.github.io/posts/anthropic-performance-takehome/)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T05:25:33Z

## Translation

タイトル: Deep Dive: Anthropic's Performance Take-Home (The One Claude Beat Humans At)
記事のタイトル: Deep Dive: Anthropic のパフォーマンスの持ち帰り (The One Claude Beat Humans At) |トリスタン・トロウエン
説明: Anthropic の詳細な内訳

記事本文:
Deep Dive: Anthropic のパフォーマンスの持ち帰り (The One Claude Beat Humans At)
今日、Anthropic は独自のパフォーマンス エンジニアリングの成果をオープンソース化しました。タスク: カスタム VLIW SIMD プロセッサ シミュレータ上で実行されるカーネルを最適化します。ベースラインには 147,734 サイクルかかります。 Claude Opus 4.5 では 1,487 サイクルまで短縮され、これはほとんどの人間を上回る 99 倍の高速化です。
AI カーネルに取り組んでいる Tristan ( @trirpi ) です。このシステム全体がどのように機能するかを詳しく見てみましょう。
アーキテクチャの概要 #
これは、単一コアを備えた VLIW (非常に長い命令ワード) SIMD (単一命令複数データ) プロセッサです (持ち帰り用の古いバージョンには複数のコアがありました)。それが何を意味するのか、詳しく説明しましょう。
VLIW: コンパイラによるスケジュールされた並列処理 #
従来のプロセッサでは、ハードウェアが実行時にどの命令を並列実行できるかを判断します。 VLIW プロセッサでは、その仕事はコンパイラ (この場合はあなた) に移ります。
単一コアには複数の機能ユニットがあり、すべて同時に実行できます。
操作を命令バンドルにパックします。各サイクルで、プロセッサは 1 つのバンドルを実行し、操作をすべてのユニットに並行してディスパッチします。バンドルに 1 つの操作のみを含めると、他のユニットはアイドル状態になります。ベースラインが非常に遅いのはそのためです。
バンドルの例 (1 サイクルで実行):
{ "alu" : [op1、op2、op3]、 "valu" : [vop1、vop2]、 "load" : [ld1、ld2]}
12 個の ALU と 6 個の VALU (それぞれ 8 要素を処理) を備えたこの単一コアは、理論的には 12 + 6×8 = 60 回の算術演算をサイクルごとに実行できます。
フローチャート LR
subgraph mem["💾 MAIN MEMORY"]
DATA["問題データ
(ツリー、インデックス、値)"]
終わり
サブグラフ スクラッチ["📦 スクラッチ スペース (1536 ワード)"]
REG["レジスタのように動作します
すべての ALU 演算はここで読み取り/書き込みを行います"]
終わり
mem |"ロード/ストア
⚠️ サイクルごとに 2 つずつ"|スクラッチ
メインメモリ: 問題データの場所

生きています。 ALU/VALU は直接アクセスできません。
スクラッチスペース: 1536 ワードの高速ストレージ。すべての計算操作は、スクラッチ アドレスの読み取り/書き込みを行います。
ボトルネック : サイクルごとに 2 つのロードと 2 つのストアのみ。多くの場合、これはコンピューティングではなく制限要因です。
プロセッサには複数のエンジンがあり、それぞれがサイクルごとに複数のスロットを実行できます。問題.py から:
SLOT_LIMITS = {
"alu" : 12 、# 12 サイクルあたりのスカラー ALU 演算
"valu" : 6 、サイクルごとの # 6 ベクトル ALU 演算
"load" : 2 、サイクルあたり 2 回のロード操作
"store" : 2 、サイクルあたり 2 回のストア操作
"flow" : 1 、サイクルあたり #1 のフロー制御操作
"debug" : 64 , # デバッグ操作 (カウントされません)
}
説明バンドルの外観 #
フローチャート LR
サブグラフ バンドル["📦 命令バンドル (1 クロック サイクル)"]
サブグラフ compute["計算"]
ALU[" アル:
('+'、dest、a、b)
('-'、宛先、a、b)
('*'、dest、a、b)
...最大 12 "]
VALU[" 値:
('*'、vdest、va、vb)
('+'、vdest、va、vb)
...最大6インチ]
終わり
サブグラフメモリ["メモリ"]
LOAD[" ロード:
('ロード'、宛先、アドレス)
('vload'、vdest、addr)"]
ストア[" ストア:
('ストア'、アドレス、ソース)
('vstore'、addr、vsrc)"]
終わり
サブグラフ コントロール["コントロール"]
FLOW[" フロー:
('選択'、d、c、a、b)"]
DEBUG[" デバッグ:
('比較'、位置、キー)
(カウントされません) "]
終わり
終わり
命令は、Python 辞書マッピング エンジン名を操作のリストにマッピングします。実際の例を次に示します。
{ "値" : [( "*" , 4 , 0 , 0 ), ( "+" , 8 , 4 , 0 )], "ロード" : [( "ロード" , 16 , 17 )]}
これにより、1 サイクルで 3 つの操作が実行されます。
ベクトル乗算: スクラッチ[4:12] = スクラッチ[0:8] * スクラッチ[0:8]
ベクトル加算: スクラッチ[8:16] = スクラッチ[4:12] + スクラッチ[0:8]
スカラーロード: スクラッチ[16] = メモリ[スクラッチ[17]]
問題: バッチ処理されたツリー トラバーサル #
カーネルは、ハッシュを使用したバッチ化されたツリー トラバーサルを実装します。流れは次のとおりです。
フローチャート LR
サブグラフのラウンド["🔄 16 ラウンド"]
R0[「ラウンド0」

"] --> R1["ラウンド 1"] --> R2["ラウンド 2"] --> RN["..."]
終わり
サブグラフ バッチ["📊 256 項目のバッチ"]
B0["アイテム0"]
B1[アイテム1]
B2[アイテム2]
BN["..."]
終わり
サブグラフ ALGO["⚙️項目ごとの計算"]
A1["idx = インデックス[i]"] --> A2["val = 値[i]"]
A2 --> A3["node_val = ツリー[idx]"]
A3 --> A4["val = hash(val ^ node_val)"]
A4 --> A5["idx = 2*idx + (2 であれば 1)"]
A5 --> A6["if idx >= n_nodes: idx = 0"]
終わり
ラウンド --> バッチ
バッチ --> アルゴ
参照カーネルから:
defreference_kernel (t: ツリー、inp: 入力):
「」
各ノードで設定した並列ツリー走査
cur_inp_val = myhash(cur_inp_val ^ node_val)
cur_inp_val が偶数の場合は、左の分岐を選択します。
ツリーの一番下に到達したら、一番上まで回り込みます。
「」
範囲内の h の場合 (inp . ラウンド):
for i in range(len(inp . indices)):
idx = 入力 。インデックス[i]
val = 入力。値[i]
val = myhash(val ^ t .values[idx])
idx = 2 * idx + ( 1 if val % 2 == 0 else 2 )
idx = 0 if idx >= len(t .values) else idx
入力。値[i] = val
入力。インデックス[i] = idx
テスト構成:
木の高さ: 10 (完全な二分木の 2047 ノード)
バッチサイズ: 256 アイテムが処理されました
これは 256 × 16 = 4096 の走査ステップに相当し、それぞれにハッシュ計算が含まれます。
ハッシュは 6 つのステージを実行し、それぞれ = (op1 const) op2 (a op3 シフト) を実行します。
各ステージ = 3 つの ALU 演算。合計: ハッシュごとに 6 × 3 = 18 の ALU 演算。
ハッシュは、カーネル実装を容易にするためにデータ駆動型で定義されています。
ハッシュステージ = [
( "+" 、 0x7ED55D16 、 "+" 、 "<<" 、 12 )、
( "^" 、 0xC761C23C 、 "^" 、 ">>" 、 19 )、
( "+" 、 0x165667B1 、 "+" 、 "<<" 、 5 )、
( "+" 、 0xD3A2646C 、 "^" 、 "<<" 、 9 )、
( "+" 、 0xFD7046C5 、 "+" 、 "<<" 、 3 )、
( "^" 、 0xB55A4F09 、 "^" 、 ">>" 、 16 )、
】
ボブ・ジェンキンスのハッシュに似ています。各ステージ: a = (a op1 val1) op2 (a op3 val3)
メインメモリ ( self.mem ): 問題の入出力
スクラッチSP

ace ( core.scratch ): 1536 ワード - レジスタ + 定数メモリ + 手動で管理されるキャッシュと考えてください。
VLEN = 8 # ベクトル長: 8 要素
N_CORES = 1 # シングルコア (古いバージョンには複数のコアがありました)
SCRATCH_SIZE = 1536 # 1536 ワードのスクラッチ領域
すべての ALU 操作は、スクラッチ アドレスの読み取りと書き込みを行います。これは共有メモリを使用した GPU のプログラミングに似ていますが、より明示的です。
def alu (self、core、op、dest、a1、a2):
a1 = コア。スクラッチ[a1]
a2 = コア。スクラッチ[a2]
マッチオペ:
"+" の場合: res = a1 + a2
「-」の場合: res = a1 - a2
ケース "*" : res = a1 * a2
case "//" : res = a1 // a2
case "^" : res = a1 ^ a2 # XOR
case "&" : res = a1 & a2 # AND
ケース「|」 : レス = a1 | a2 # または
case "<<" : res = a1 << a2 # 左シフト
case ">>" : res = a1 >> a2 # 右シフト
case "%" : res = a1 % a2 # モジュロ
case "<" : res = int(a1 < a2) # 比較
case "== : res = int(a1 == a2)
res = res % ( 2 ** 32 ) # 32 ビット符号なしラップ
自分自身。 Scratch_write[宛先] = 解像度
Vector ALU 演算 - SIMD 部分:
def 値 (自己、コア、* スロット):
マッチスロット:
case ( "vbroadcast" 、宛先、送信元):
# 8 つのベクター レーンすべてにスカラーをブロードキャストする
範囲内の i の場合 (VLEN):
自分自身。 Scratch_write[dest + i] = コア 。スクラッチ[ソース]
case ( "multiply_add" , dest, a, b, c):
# 融合乗算加算: dest = a * b + c
範囲内の i の場合 (VLEN):
mul = (コア . スクラッチ[a + i] * コア . スクラッチ[b + i]) % ( 2 ** 32 )
自分自身。 Scratch_write[dest + i] = (mul + core .scratch[c + i]) % ( 2 ** 32 )
case (op、dest、a1、a2):
# 要素ごとに適用される任意のスカラー演算
範囲内の i の場合 (VLEN):
自分自身。 alu(コア、op、dest + i、a1 + i、a2 + i)
メモリ操作 #
ロード/ストア - サイクルごとにそれぞれ 2 つだけ:
def ロード (セルフ、コア、* スロット):
マッチスロット:
case ( "load" , dest, addr):
自分自身。 Scratch_write[dest] = self 。メモリ[コア.スクラッチ[アドレス]]
case ( "vload" , dest, addr): # 8 連続要素
addr = コア . sc

ラッチ[アドレス]
範囲内の vi の場合 (VLEN):
自分自身。 Scratch_write[dest + vi] = self 。 mem[アドレス + vi]
case ( "const" 、 dest、val):
自分自身。 Scratch_write[dest] = (val) % ( 2 ** 32 )
主なボトルネック: サイクルごとに 2 つのロードのみ。ベクター ロード ( vload ) ヘルプ - 1 つのスロットに 8 つの要素!
フロー操作 - ブランチレス プログラミングに重要:
def フロー (セルフ、コア、* スロット):
マッチスロット:
case ( "select" 、 dest、 cond、 a、 b):
# ブランチレス: dest = cond ? a:b
自分自身。スクラッチ書き込み[宛先] = (
コア。傷「a」あればコアも有ります。スクラッチ[条件] != 0 else コア。スクラッチ[b]
）
case ( "vselect" 、 dest、 cond、 a、 b):
# ベクター版
範囲内の vi の場合 (VLEN):
自分自身。 Scratch_write[dest + vi] = (
コア。スクラッチ[a + vi]
コアの場合スクラッチ[cond + vi] != 0
それ以外のコア。スクラッチ[b + vi]
）
case ( "cond_jump" , cond, addr):
コアの場合スクラッチ[条件] != 0 :
コア。 pc = アドレス
case ( "jump" , addr):
コア。 pc = アドレス
ベースラインが非常に遅い理由 #
ベースライン カーネルは、サイクルごとに 1 つの操作を意図的に使用します。
def build (self、slots: list[tuple[Engine, tuple]]、vliw: bool = False ):
# 命令バンドルごとに 1 つのスロットのみを使用するシンプルなスロット パッキング
instrs = []
エンジン用、スロットインスロット:
命令 。 append({engine: [slot]}) # バンドルごとに 1 つの操作!
命令を返す
したがって、代わりに:
{ "alu" : [op1, op2, op3], "load" : [load1]} # 1 サイクル
得られるもの:
{ "alu" : [op1]} # サイクル 1
{ "alu" : [op2]} # サイクル 2
{ "alu" : [op3]} # サイクル 3
{ "load" : [load1]} # サイクル 4
1 サイクルではなく 4 サイクル。12 個の ALU スロットは空のままです。
シミュレータは、 Perfetto で表示可能な Chrome Trace Event Format トレースを出力します。
python perf_takehome.py Tests.test_kernel_trace
python watch_trace.py # ライブリロードトレースでブラウザを開きます
watch_trace.py サーバーは、トレースが変更されると自動的に再ロードするため、反復処理に最適です。
デバッグ操作は、サイクルをカウントせずに中間値を検証します。
本体もappend(( "デバッグ"

, ( "比較" , tmp_val, (round, i, "hashed_val" ))))
最適化戦略 #
独立した操作を同じサイクルにパックします。
{ "alu" : [op1、op2、op3]、 "load" : [load1、load2]} # 5 演算、1 サイクル
2. SIMD ベクトル化 #
valu と vload / vstore を使用して、一度に 8 つのバッチ項目を処理します。
さまざまな反復の計算をオーバーラップして、すべてのエンジンをビジー状態に保ちます。
# 条件付きジャンプの代わりに:
offset = select(val % 2 == 0 , 1 , 2 )
idx = 2 * idx + オフセット
難しい部分: データの依存関係 #
現在の値をハッシュするまでは、次のツリー インデックスを計算できません。それがシリアル依存関係チェーンです。重要なのは、異なるバッチ項目間の並列性を見つけることです。
SIMD プログラミング - インテル組み込みガイド
ソフトウェア パイプライン - ウィキペディア
コンピュータ アーキテクチャ: 定量的アプローチ
これは役に立ちましたか? Twitter/X または GitHub でフォローしてください。

## Original Extract

A detailed breakdown of Anthropic

Deep Dive: Anthropic's Performance Take-Home (The One Claude Beat Humans At)
Today, Anthropic open-sourced their original performance engineering take-home . The task: optimize a kernel running on a custom VLIW SIMD processor simulator. The baseline takes 147,734 cycles . Claude Opus 4.5 got it down to 1,487 cycles - a 99x speedup that beat most humans.
I’m Tristan ( @trirpi ), and I work on AI kernels. Let’s break down how this whole system works.
The Architecture at a Glance #
This is a VLIW (Very Long Instruction Word) SIMD (Single Instruction Multiple Data) processor with a single core (older versions of the take-home had multiple cores). Let me break down what that means.
VLIW: Compiler-Scheduled Parallelism #
In a traditional processor, hardware figures out at runtime which instructions can run in parallel. In a VLIW processor, that job shifts to the compiler (or in this case, you).
The single core has multiple functional units that can all execute simultaneously:
You pack operations into instruction bundles . Each cycle, the processor executes one bundle, dispatching operations to all the units in parallel. If you only put one operation in a bundle, the other units sit idle. That’s why the baseline is so slow.
Example bundle (executes in 1 cycle):
{ "alu" : [op1, op2, op3], "valu" : [vop1, vop2], "load" : [ld1, ld2]}
With 12 ALUs and 6 VALUs (each processing 8 elements), this single core can theoretically do 12 + 6×8 = 60 arithmetic operations per cycle.
flowchart LR
subgraph mem["💾 MAIN MEMORY"]
DATA["Problem Data
(tree, indices, values)"]
end
subgraph scratch["📦 SCRATCH SPACE (1536 words)"]
REG["Works like registers
All ALU ops read/write here"]
end
mem |"LOAD/STORE
⚠️ 2 each per cycle"| scratch
Main Memory : Where the problem data lives. ALU/VALU can’t access it directly.
Scratch Space : 1536 words of fast storage. All compute operations read/write scratch addresses.
Bottleneck : Only 2 loads and 2 stores per cycle. This is often the limiting factor, not compute.
The processor has multiple engines , each capable of executing multiple slots per cycle. From problem.py :
SLOT_LIMITS = {
"alu" : 12 , # 12 scalar ALU operations per cycle
"valu" : 6 , # 6 vector ALU operations per cycle
"load" : 2 , # 2 load operations per cycle
"store" : 2 , # 2 store operations per cycle
"flow" : 1 , # 1 flow control operation per cycle
"debug" : 64 , # Debug operations (not counted)
}
What an Instruction Bundle Looks Like #
flowchart LR
subgraph bundle["📦 Instruction Bundle (1 clock cycle)"]
subgraph compute["Compute"]
ALU[" alu:
('+', dest, a, b)
('-', dest, a, b)
('*', dest, a, b)
...up to 12 "]
VALU[" valu:
('*', vdest, va, vb)
('+', vdest, va, vb)
...up to 6 "]
end
subgraph memory["Memory"]
LOAD[" load:
('load', dest, addr)
('vload', vdest, addr)"]
STORE[" store:
('store', addr, src)
('vstore', addr, vsrc)"]
end
subgraph control["Control"]
FLOW[" flow:
('select', d, c, a, b)"]
DEBUG[" debug:
('compare', loc, key)
(not counted) "]
end
end
An instruction is a Python dict mapping engine names to lists of operations. Here’s a real example :
{ "valu" : [( "*" , 4 , 0 , 0 ), ( "+" , 8 , 4 , 0 )], "load" : [( "load" , 16 , 17 )]}
This executes three operations in one cycle :
Vector multiply: scratch[4:12] = scratch[0:8] * scratch[0:8]
Vector add: scratch[8:16] = scratch[4:12] + scratch[0:8]
Scalar load: scratch[16] = memory[scratch[17]]
The Problem: Batched Tree Traversal #
The kernel implements a batched tree traversal with hashing. Here’s the flow:
flowchart LR
subgraph rounds["🔄 16 Rounds"]
R0["Round 0"] --> R1["Round 1"] --> R2["Round 2"] --> RN["..."]
end
subgraph batch["📊 Batch of 256 items"]
B0["Item 0"]
B1["Item 1"]
B2["Item 2"]
BN["..."]
end
subgraph ALGO["⚙️ Per-item computation"]
A1["idx = indices[i]"] --> A2["val = values[i]"]
A2 --> A3["node_val = tree[idx]"]
A3 --> A4["val = hash(val ^ node_val)"]
A4 --> A5["idx = 2*idx + (1 if even else 2)"]
A5 --> A6["if idx >= n_nodes: idx = 0"]
end
rounds --> batch
batch --> ALGO
From the reference kernel :
def reference_kernel (t: Tree, inp: Input):
"""
A parallel tree traversal where at each node we set
cur_inp_val = myhash(cur_inp_val ^ node_val)
and then choose the left branch if cur_inp_val is even.
If we reach the bottom of the tree we wrap around to the top.
"""
for h in range(inp . rounds):
for i in range(len(inp . indices)):
idx = inp . indices[i]
val = inp . values[i]
val = myhash(val ^ t . values[idx])
idx = 2 * idx + ( 1 if val % 2 == 0 else 2 )
idx = 0 if idx >= len(t . values) else idx
inp . values[i] = val
inp . indices[i] = idx
Test configuration:
Tree height : 10 (2047 nodes in a perfect binary tree )
Batch size : 256 items processed
That’s 256 × 16 = 4096 traversal steps, each involving a hash computation.
The hash runs 6 stages, each doing a = (a op1 const) op2 (a op3 shift) :
Each stage = 3 ALU ops. Total: 6 × 3 = 18 ALU operations per hash.
The hash is defined data-driven for easy kernel implementation:
HASH_STAGES = [
( "+" , 0x7ED55D16 , "+" , "<<" , 12 ),
( "^" , 0xC761C23C , "^" , ">>" , 19 ),
( "+" , 0x165667B1 , "+" , "<<" , 5 ),
( "+" , 0xD3A2646C , "^" , "<<" , 9 ),
( "+" , 0xFD7046C5 , "+" , "<<" , 3 ),
( "^" , 0xB55A4F09 , "^" , ">>" , 16 ),
]
Similar to Bob Jenkins’ hash . Each stage: a = (a op1 val1) op2 (a op3 val3)
Main Memory ( self.mem ): Problem input/output
Scratch Space ( core.scratch ): 1536 words - think of it as registers + constant memory + manually managed cache
VLEN = 8 # Vector length: 8 elements
N_CORES = 1 # Single core (older versions had multiple)
SCRATCH_SIZE = 1536 # 1536 words of scratch space
Every ALU operation reads and writes scratch addresses. It’s like programming a GPU with shared memory , but more explicit.
def alu (self, core, op, dest, a1, a2):
a1 = core . scratch[a1]
a2 = core . scratch[a2]
match op:
case "+" : res = a1 + a2
case "-" : res = a1 - a2
case "*" : res = a1 * a2
case "//" : res = a1 // a2
case "^" : res = a1 ^ a2 # XOR
case "&" : res = a1 & a2 # AND
case "|" : res = a1 | a2 # OR
case "<<" : res = a1 << a2 # Left shift
case ">>" : res = a1 >> a2 # Right shift
case "%" : res = a1 % a2 # Modulo
case "<" : res = int(a1 < a2) # Comparison
case "==" : res = int(a1 == a2)
res = res % ( 2 ** 32 ) # 32-bit unsigned wrap
self . scratch_write[dest] = res
Vector ALU ops - the SIMD part:
def valu (self, core, * slot):
match slot:
case ( "vbroadcast" , dest, src):
# Broadcast scalar to all 8 vector lanes
for i in range(VLEN):
self . scratch_write[dest + i] = core . scratch[src]
case ( "multiply_add" , dest, a, b, c):
# Fused multiply-add: dest = a * b + c
for i in range(VLEN):
mul = (core . scratch[a + i] * core . scratch[b + i]) % ( 2 ** 32 )
self . scratch_write[dest + i] = (mul + core . scratch[c + i]) % ( 2 ** 32 )
case (op, dest, a1, a2):
# Any scalar op applied element-wise
for i in range(VLEN):
self . alu(core, op, dest + i, a1 + i, a2 + i)
Memory Operations #
Load/store - only 2 of each per cycle:
def load (self, core, * slot):
match slot:
case ( "load" , dest, addr):
self . scratch_write[dest] = self . mem[core . scratch[addr]]
case ( "vload" , dest, addr): # 8 consecutive elements
addr = core . scratch[addr]
for vi in range(VLEN):
self . scratch_write[dest + vi] = self . mem[addr + vi]
case ( "const" , dest, val):
self . scratch_write[dest] = (val) % ( 2 ** 32 )
Key bottleneck: only 2 loads per cycle . Vector loads ( vload ) help - 8 elements in one slot!
Flow ops - crucial for branchless programming :
def flow (self, core, * slot):
match slot:
case ( "select" , dest, cond, a, b):
# Branchless: dest = cond ? a : b
self . scratch_write[dest] = (
core . scratch[a] if core . scratch[cond] != 0 else core . scratch[b]
)
case ( "vselect" , dest, cond, a, b):
# Vector version
for vi in range(VLEN):
self . scratch_write[dest + vi] = (
core . scratch[a + vi]
if core . scratch[cond + vi] != 0
else core . scratch[b + vi]
)
case ( "cond_jump" , cond, addr):
if core . scratch[cond] != 0 :
core . pc = addr
case ( "jump" , addr):
core . pc = addr
Why the Baseline Is So Slow #
The baseline kernel deliberately uses one operation per cycle :
def build (self, slots: list[tuple[Engine, tuple]], vliw: bool = False ):
# Simple slot packing that just uses one slot per instruction bundle
instrs = []
for engine, slot in slots:
instrs . append({engine: [slot]}) # One op per bundle!
return instrs
So instead of:
{ "alu" : [op1, op2, op3], "load" : [load1]} # 1 cycle
You get:
{ "alu" : [op1]} # Cycle 1
{ "alu" : [op2]} # Cycle 2
{ "alu" : [op3]} # Cycle 3
{ "load" : [load1]} # Cycle 4
4 cycles instead of 1. The 12 ALU slots sit empty.
The simulator outputs Chrome Trace Event Format traces viewable in Perfetto :
python perf_takehome.py Tests.test_kernel_trace
python watch_trace.py # Opens browser with live-reloading trace
The watch_trace.py server auto-reloads traces when they change - great for iterating.
Debug ops verify intermediate values without counting cycles:
body . append(( "debug" , ( "compare" , tmp_val, (round, i, "hashed_val" ))))
Optimization Strategies #
Pack independent operations into the same cycle:
{ "alu" : [op1, op2, op3], "load" : [load1, load2]} # 5 ops, 1 cycle
2. SIMD Vectorization #
Process 8 batch items at once with valu and vload / vstore .
Overlap computation of different iterations to keep all engines busy.
# Instead of conditional jumps:
offset = select(val % 2 == 0 , 1 , 2 )
idx = 2 * idx + offset
The Hard Part: Data Dependencies #
You can’t compute the next tree index until you’ve hashed the current value. That’s a serial dependency chain . The trick is finding parallelism across different batch items.
SIMD Programming - Intel Intrinsics Guide
Software Pipelining - Wikipedia
Computer Architecture: A Quantitative Approach
Found this useful? Follow me on Twitter/X or GitHub .
