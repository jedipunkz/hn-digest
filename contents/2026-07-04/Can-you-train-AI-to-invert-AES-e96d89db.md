---
source: "https://www.juanslozano.com/blog/ai-aes"
hn_url: "https://news.ycombinator.com/item?id=48789544"
title: "Can you train AI to invert AES?"
article_title: "AES Inversion — Juan Sebastian Lozano"
author: "juansebastianl"
captured_at: "2026-07-04T22:47:55Z"
capture_tool: "hn-digest"
hn_id: 48789544
score: 1
comments: 0
posted_at: "2026-07-04T22:11:17Z"
tags:
  - hacker-news
  - translated
---

# Can you train AI to invert AES?

- HN: [48789544](https://news.ycombinator.com/item?id=48789544)
- Source: [www.juanslozano.com](https://www.juanslozano.com/blog/ai-aes)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T22:11:17Z

## Translation

タイトル: AI をトレーニングして AES を逆転させることはできますか?
記事のタイトル: AES 逆変換 — Juan Sebastian Lozano

記事本文:
AES 逆変換 — Juan Sebastian Lozano
Juan Sebastian Lozano ← ホーム GitHub ↗ Twitter ↗ LinkedIn ↗ 学者 ↗ 執筆 04 — 2026 年 7 月
AI をトレーニングして AES を逆転させることはできますか?
すべての https ブラウザ セッション、WPA2/3 によるすべての Wi-Fi 接続の中核では、
すべてのハードウェア レベルのデバイス暗号化モジュールには、AES アルゴリズムが組み込まれています。これ
アルゴリズムを逆転させます - 秘密鍵または秘密鍵を見つけるプロセス
クリアテキスト - 最も危険であると同時に価値のあるアルゴリズムの 1 つ
今日の問題。
AI がソフトウェアを食べるので、AI が次の問題を解決できる可能性があると思います。
AES 反転。もちろん、確かなことはわかりませんが、この記事では次のことを説明しています。
それがどのように起こるかについての議論。モデルをトレーニングして実行したい場合は、
私が話していることについては、この github リポジトリの RL 環境を確認してください。
モデルが AES を解決できるようにするには、2 つの問題と闘わなければなりません。
報酬の希薄性と複雑性理論。明らかにこれらのうちの 1 つは、
賢明な研究者は他のものよりも根本的なものを見つけることができます
ほとんどすべての問題に対する中間報酬、超知性は変わらない
複雑性理論によって私たちに課せられた数学的制約。します
報酬の少なさの問題については次のセクションで扱いますが、ここでは次のセクションに進みます。
複雑性理論への私たちの注目。
この記事では AES の基本については説明しませんが、AES について詳しく知りたい場合は、
入門編として、Braincoke ブログでこれをお勧めします。
AES について理解する必要がある重要な点は、AES が置換順列であるということです。
暗号。バイトのブロックを取得し、それらを並べ替えることによって機能することを意味します。
一連の関数に従って。目標は、動作するものを作成することです
一方向性関数のようなもので、簡単な関数の仮説セットです。
計算はするが逆変換は難しい - 固有

あらゆる確率多項式時間を同盟します
アルゴリズムは確率 $p$ でそれらを反転できません。彼らは理想的には
インジェクションを使用するため、反転用の迅速なアルゴリズムがある場合は、
原文を復元します。これは、AES が多くのラウンドから構築されていることを意味します。
キーによってパラメータ化された比較的単純な順列。
AES の反転には、多数の構成された順列の反転が含まれます。
したがって、当然のことながら組み合わせ問題になります。今はどんなことがあっても
超知性は長期的にはうまくいく、どんな超知性も迂回することはできない
数学とか複雑性理論とか。 AES が効率的であると私たちが信じるためには、
そもそも、私がなぜそう信じているのかについて議論をしなければなりません
このような効率的な反転が存在する可能性があります。
まず、AES は充足可能性の問題として表現できるということから始めます。
AES の特定のモードの選択
AES は実際にはさまざまなモードで実行できるため、具体的なターゲットを選択しました。
AES-XTS のモードで、フルディスク暗号化に使用されます。 2つ使用します
独立した AES-256 キー: $K_1$ はデータを暗号化し、$K_2$ はデータを暗号化します。
ブロックごとの微調整を導き出します。セクター番号 $s$ のブロックとブロック インデックスの場合
$j$。ここで $\oplus$ は単に XOR です。 $\otimes$ は要素ごとの乗算です
特定の有限フィールド内で。まずブロックごとの調整を設定します。
$$T_0 = \mathrm{AES}_{K_2}\!\big(\mathrm{LE}_{128}(s)\big), \qquad T_j = T_0 \otimes \alpha^{\,j} \ \text{ in } \mathrm{GF}(2^{128}),$$
$$C = \mathrm{AES}_{K_1}\!\big(P \oplus T_j\big) \oplus T_j,$$
$\mathrm{LE}_{128}(s)$ は、リトルエンディアン 128 ビットのセクター番号です。
ブロックであり、$\alpha$ は乗算ジェネレーターです。
$\mathrm{GF}(2^{128})$。データ パスは通常の AES-256: 初期値です。
AddRoundKey、次に SubBytes、ShiftRows、MixColumns、AddRoundKey の $R-1$ ラウンド、
そしてMixColumnsなしの最終ラウンド。
反転

は逆です。暗号文 $C$ と公開インデックスを観察します。
$s$ と $j$ ですが、平文の $P$ (および暗黙的にキー) が必要です。
これをソルバーで攻撃するには、まず暗号を具体的なものとして書き出す必要があります。
ソルバーが分解できるオブジェクトなので、制約を説明する前にそれを理解する価値があります。
モデル自体をレイアウトします。
XTS ブロックのウィンドウを 3 つのブロックで構成される 1 つのフラット回路にコンパイルします。
レコードの種類:
値は、実際に使用されるバイトサイズの量です。いくつかは
検索する入力 (プレーンテキストとキーバイト)、一部は定数です
(観察された暗号文)、および一部は内部ワイヤです - 中間
AES 計算の途中で出現するバイト。
Ops は 1 つの値を計算する基本的な操作です
他から: S-box ルックアップ、MixColumns バイト、XOR、「$x$ の乗算」
微調整チェーンのステップ、キー スケジュールからのラウンド キー バイトなど。
ops は、個々のステップに分割された単なる AES そのものです。
制約とは、保持すべき述語です。
解決策:「このワイヤーは、それを生成するはずの操作と等しい」、「これは予測した
暗号文のバイトがターゲットと等しい。」それぞれの制約が後に変化します。
ラグランジュ最適化を行っている場合は残差。
この問題を解決するために SAT ソルバーを作成していた場合、ソルバーの状態
実際に変化するのは平文のバッファ $K_1$ だけです。
$K_2$、および (以下の 2 つのエンコーディングのいずれかで) ワイヤー。
モデリングにおける唯一の本当の決定は、暗号をどれだけ細かく公開するかということです。
は対極にある 2 つの選択肢です。
不透明エンコーディングでは、ブロック全体が単一のモノリシックです。
op: 平文と鍵を渡すと、参照暗号をエンドツーエンドで実行します。
そして予測された暗号文を返します。ワイヤーは全くありません。唯一の制約
それは予測された暗号文の式です

ターゲットも同様です (128 ビットすべてを一度に)。これ
コンパクトで忠実なスコアラーですが、探索風景としては黒です
箱。移動できるのはプレーンテキストとキーバイトだけであり、AES のせいで
単一のキービットフリップが雪崩を起こし、16 個の出力バイトすべてが予期せずスクランブルされます。
- 暗号の「途中」であるという意味のある意味はないので、ローカルです。
検索には保持するものは何もありません。
低エンコーディングでは、中間報酬を目指します。
毎
AES マイクロステップは、独自の内部ワイヤとすべてのワイヤを書き込む独自のオペレーションになります。
を運ぶ
一貫性
ワイヤーがそれを定義する演算と実際に等しいことを主張する制約、そして
すべての暗号文バイトには、ターゲットに対する独自の等価制約が適用されます。これら
中間ワイヤは追加の未知数 $W$ です。
ローダウンのポイントは局所的な構造です。ブラックボックスの代わりに、
ハンドルのみが ~64 キーとプレーンテキスト バイトであり、検索には数千バイトが含まれるようになりました。
中間ワイヤーを一度に 1 本ずつナッジし、再スコアリングできるのはほんの一握りだけです
各ナッジが触れる制約の数。その代償として、これらすべてのワイヤーが
相互に一貫性を保つこと、それがまさに
一貫性制約の測定値、およびソルバーがゼロに追い込む必要があるもの
他のすべてのものと一緒に。
そこで、反転を未知数に対する制約充足問題としてキャストします。
$$x = \big(\,\underbrace{P}_{\text{プレーンテキスト}},\ \underbrace{K_1}_{\text{データ キー}},\ \underbrace{K_2}_{\text{微調整キー}},\ \underbrace{W}_{\text{内部 AES ワイヤー}}\,\big),$$
その制約は次の 3 つのクラスに分類されます。
一貫性 (低エンコーディングのみ) - すべての内部ワイヤが等しい
それを生成する演算。つまり、中間バイトは実際に有効な 1 つのバイトを形成します。
AES 計算。
目標 - 予測された暗号文は観測された $C$ と等しくなります。
ASCII - 作るだけ

問題はより扱いやすくなり、制限できるようになります
すべてのプレーンテキスト バイトはテキスト ASCII になります。これは、次のような場合には不合理ではありません。
AES-XTS 設定で何を復号化しているのかがよくわかります。
各制約 $i$ は非負の残差 $r_i(x) \ge 0$ になります。
制約が満たされると、正確にゼロになります。等価制約が測定される
フラットな正誤ではなくビット単位のハミング距離として、「近づく」
ターゲットに合わせると (間違ったビットが少なくなり)、残差が低下し、ソルバーにエネルギーが与えられます。
続く風景。割り当ては次の場合に実行可能です
複雑さの理論に少しでも精通している場合は、すぐにこれを理解できるはずです。
いくつかの危険信号を立ててください。 AES の反転が充足可能性であるという事実
問題、および充足可能性問題は一般に NP 困難であるため、
効率的なバージョンを見つけられるという考えには慎重です。 2つ作ります
観察:
暗号に関する複雑性理論上の事実は、彼らの考えではありません。
最悪の場合の複雑さ - 平均的な場合の複雑さです。
最悪の場合の困難である NP 問題のクラスの個々のインスタンスは、次のようになります。
実際には比較的珍しい
AES のセキュリティは、AES が想定する最悪の場合の複雑さのクラスに基づいていません。
代わりに、特定の鍵と平文のペアに対して、
反転問題が非常に難しいことは比較的確実です。これは
一方向関数の平均的な複雑さは非常に高いため、
高い。しかし、ボグダノフとトレヴィサン (2005)
そして
アカビアら。アル。 (2005) ここでの平均ケースの複雑さは最悪のケースの複雑さから導かれたものではないことを示しています。
NP のクラスの (多項式階層を想定)。一般に、それはかなり依存します
どの複雑さのクラスで作業しているか、最悪の場合の複雑さかどうかについて少し
平均的なケースの複雑さまで削減されます。
2番目の事実は、

より微妙ですが、次の場合を理解するのが最も簡単です。
SATの問題。与えられた数の充足可能性問題を解くと、
変数 $n$ の句 $c$ には、比率 $q = c/n$ があります
変数から句まで。興味深いのは、この比率が変化すると、
本当に重要なことが起こります。変数の数が多く、
節の数が少ないため、問題を満たすのは比較的簡単です。
多くの自由度があります。比率が低い場合、問題は非常に困難です
満足させることは、多くの場合不可能です。
この中間点があり、通常はフェーズとして説明されます。
満足できる状態から満足できない状態へ移行し、そこで問題が発生します。
満足することが不可能なほどに過度に決定されます。
qₖ ≈ 4.27 簡単・SAT
難しい
簡単・UNSAT
比 q = c/n 2.00 P(満足) 100% 中央値 DPLL 決定 37 レジーム過小制約 — 簡単、満足可能 ±J
q = 2.00。トランジション付近では風景がガラス状になり、
スピンは決して安定しません。それから遠く離れて、彼らはほぼすぐに注文します。
// それぞれ +1 または -1 の「スピン」の 28x28 グリッド。あらゆるスピンは、
// 固定ランダム結合 J によってその 4 つの隣接結合と結合されます。
// +1 (「この 2 人は同意したい」) または -1 (「この 2 人は同意したい」)
// 同意しません」）。グリッド全体のエネルギーは
//
// E = - J *スピン_a *スピン_bの隣接ペアの合計
//
// したがって、結合が望むものを得るとき、結合は低エネルギーで貢献します。
const GRID = 28 ;
const 温度 = 0.9 ;
const pin = new Int8Array ( GRID * GRID );
const bindingToRight = new Int8Array ( GRID * GRID );
const bindingToBelow = new Int8Array ( GRID * GRID );
// ラティスはエッジ (トーラス) で回り込むため、すべてのサイトが
// ちょうど 4 つの近傍があります。
const Wrap = ( v ) => (v + GRID )

%グリッド;
const siteAt = ( row,col ) => ラップ (行) * GRID + ラップ (列);
関数randomizeSpins(){
for ( s = 0 ; s < スピンの長さ ; s++) {
スピン[s] = 数学 。ランダム () < 0.5 ? - 1:1;
}
}
// 「Frustration」はスライダーが制御するノブです。 0 では、すべての結合は次のようになります。
// +1 するとグリッドは単純な強磁性体になります: あなたの方向に反転します
// 近隣の人々が常に助けてくれるので、ほとんど大きなドメインに注文されます
// すぐに。フラストレーションが高まると、絆が-1されてループすることが増えます
// どの代入もすべての結合を満たすことができないように見えます --
// スピングラスの特徴を定義します。エネルギー情勢は崩壊する
// 多くの競合する極小値に変換します。これは、
// 相転移近くの SAT インスタンス。
関数再構築カップリング (フラストレーション) {
const pAntiferro = 0.5 * 欲求不満;
for ( s = 0 ; s < スピンの長さ ; s++) {
結合ToRight[s] = 数学 。ランダム () < pAntiferro ? - 1:1;
結合ToBelow[s] = 数学 。ランダム () < pAntiferro ? - 1:1;
}
}
// 1 回のスイープ = 平均して、サイトごとに 1 回の更新試行。
関数メトロポリススイープ(){
for ( 試行 = 0 ; 試行 < GRID * GRID ; 試行++) {
// ランダムなサイトを選択します。
const row = (Math .random() * GRID ) | 0 ;
const col = (Math .random() * GRID ) | 0 ;
const here = siteAt

[切り捨てられた]

## Original Extract

AES Inversion — Juan Sebastian Lozano
Juan Sebastian Lozano ← Home GitHub ↗ Twitter ↗ LinkedIn ↗ Scholar ↗ Writing 04 — July 2026
Can you train AI to invert AES?
At the core of every https browser session, every wifi connection with WPA2/3,
every hardware level device encryption module, sits the AES algorithm. This
makes inverting the algorithm - the process of finding the private key or the
clear text - one of the most simultaneously dangerous and valuable algorithmic
problems today.
As AI eats software, I think it's possible that AI could solve the problem of
AES Inversion. Of course, I dont know for sure, but this post lays out the
argument for how it might happen. If you want to try and train a model to do
what Im talking about, check out the RL env in this github repo .
In order for a model to be able to solve AES you have to fight two problems -
the sparsity of the reward, and complexity theory. Obviously one of these is
more fundamental than the other, while a clever researcher can find
intermediate rewards for almost any problem, superintelligence doesn't change
the mathematical constraints imposed on us by complexity theory. I will
address the sparse reward problem in the next section, but for now let us turn
our attention to complexity theory.
I wont explain the basics of AES in the post, but if you want a good
introduction to it, I recommend this one on the Braincoke blog.
The important thing to understand about AES is that it's a substitution-permutation
cipher, which means that it works by taking a block of bytes and permuting them
according to a set of functions. The goal is to create a something that behaves
like a one-way function, which are a hypothesized set of functions that are easy
to compute but hard to invert - specifically any probabilistic polynomial time
algorithm fails to invert them with probability $p$. They are ideally
injections so that if you have a quick algorithm for inversion, you can
recover the original text. This means that AES is built from many rounds of
relatively simple permutations parameterized by the key(s).
Inverting AES involves inverting a large number of composed permutations and
therefore is naturally a combinatorial problem. Now no matter how
superintelligence pans out in the long run, no superintelligence can bypass
mathematics or complexity theory. For us to believe that AES is efficiently
invertible in the first place, I have to make an argument as to why I believe
such an efficient inversion likely exists.
I will start by stating that AES can be stated as a satisfiability problem.
Choosing a particular mode of AES
Because AES can actually be run in different modes, we chose a concrete target
of AES-XTS, which the mode used for full-disk encryption. It uses two
independent AES-256 keys: $K_1$ encrypts the data and $K_2$
derives a per-block tweak. For the block at sector number $s$ and block index
$j$. Here $\oplus$ is simply XOR. $\otimes$ is elementwise multiplication
in a particular finite field. We first set up our per-block tweak:
$$T_0 = \mathrm{AES}_{K_2}\!\big(\mathrm{LE}_{128}(s)\big), \qquad T_j = T_0 \otimes \alpha^{\,j} \ \text{ in } \mathrm{GF}(2^{128}),$$
$$C = \mathrm{AES}_{K_1}\!\big(P \oplus T_j\big) \oplus T_j,$$
where $\mathrm{LE}_{128}(s)$ is the sector number as a little-endian 128-bit
block and $\alpha$ is a multiplicative generator in
$\mathrm{GF}(2^{128})$. The data path is ordinary AES-256: an initial
AddRoundKey, then $R-1$ rounds of SubBytes, ShiftRows, MixColumns and AddRoundKey,
then a final round without MixColumns.
Inversion is the reverse. We observe the ciphertext $C$ and the public indices
$s$ and $j$, and we want the plaintext $P$ (and, implicitly, the keys).
To attack this with a solver we first have to write the cipher down as a concrete
object the solver can pick apart, so before I state the constraints it's worth
laying out the model itself.
We compile a window of XTS blocks into one flat circuit made of three
kinds of record:
Values are the byte-sized quantities in play. Some are
inputs we search over (the plaintext and the key bytes), some are constants
(the observed ciphertext), and some are internal wires - the intermediate
bytes that appear part-way through an AES computation.
Ops are the primitive operations that compute one value
from others: an S-box lookup, a MixColumns byte, an XOR, the "multiply by $x$"
step of the tweak chain, a round-key byte from the key schedule, and so on.
The ops are just AES itself, chopped into individual steps.
Constraints are the predicates that ought to hold at
a solution:"this wire equals the op that is supposed to produce it", "this predicted
ciphertext byte equals the target". Each constraint is what later turns into
a residual if you're doing Lagrangian optimization.
If you were writing a SAT solver to solve this problem, the state the solver
actually mutates is just the buffers of plaintext, $K_1$,
$K_2$, and (in one of the two encodings below) the wires.
The one real modelling decision is how finely we expose the cipher, and there
are two choices that sit at opposite ends.
In the opaque encoding the whole block is a single monolithic
op: hand it the plaintext and the keys, it runs the reference cipher end to end,
and it returns the predicted ciphertext. There are no wires at all. The only constraints
are that the predicted ciphertext equals the target (all 128 bits at once). This
is compact and it is a faithful scorer, but as a search landscape it is a black
box. The only things you can move are plaintext and key bytes, and because of AES's
avalanche a single key-bit flip scrambles all sixteen output bytes unpredictably
- there is no meaningful sense of being "partway" through the cipher, so local
search has nothing to hold onto.
In the lowered encoding we go towards intermediary rewards.
Every
AES micro-step becomes its own op writing its own internal wire, and every wire
carries a
consistency
constraint asserting that the wire really does equal the op that defines it, and
every ciphertext byte gets its own equality constraint against the target. These
intermediate wires are the extra unknowns $W$.
The whole point of lowering is local structure. Instead of a black box whose
only handles are the ~64 key and plaintext bytes, the search now has thousands
of intermediate wires it can nudge one at a time, re-scoring only the handful
of constraints each nudge touches. The price is that all those wires now have
to be kept mutually consistent with one another - which is exactly what the
consistency constraints measure, and what the solver has to drive to zero
alongside everything else.
So we cast inversion as a constraint-satisfaction problem over the unknowns
$$x = \big(\,\underbrace{P}_{\text{plaintext}},\ \underbrace{K_1}_{\text{data key}},\ \underbrace{K_2}_{\text{tweak key}},\ \underbrace{W}_{\text{internal AES wires}}\,\big),$$
whose constraints fall into three classes:
Consistency (lowered encoding only) - every internal wire equals
the op that produces it, i.e. the intermediate bytes really do form one valid
AES computation.
Goal - the predicted ciphertext equals the observed $C$.
ASCII - Just to make the problem more tractable, we can constrain
every plaintext byte to be text ASCII, which is not unreasonable if you have
a good sense of what you're decrypting in the AES-XTS setting.
Each constraint $i$ becomes a non-negative residual $r_i(x) \ge 0$ that
is zero exactly when the constraint is satisfied. Equality constraints are measured
as a Hamming distance in bits rather than a flat right/wrong, so getting "closer"
to the target (fewer wrong bits) lowers the residual and gives the solver an energy
landscape to follow. The assignment is feasible when
Immediately if you're at all familiar with complexity theory, this should
raise some red flags. The fact that inversion of AES is a satisfiability
problem, and satisfiability problems are in general NP-hard, should make you
wary of the idea of being able to find an efficient version. Ill make two
observations:
The relevant complexity theoretic fact about about cyphers is not their
worst-case complexity - It is their average-case complexity.
Any individual instance of a class of NP problems being worst-case hard is
actually relatively rare
The security of AES is not based on the worst-case complexity class that it
lives in. Instead it's based the fact that for any given key-plaintext pair we
can be relatively certain that the inversion problem is quite hard. This is
due to the fact that the average-case complexity of one-way functions is very
high. However, Bogdanov and Trevisan (2005)
and
Akavia et. al. (2005) show that the average case complexity here is not drawn from the worst case complexity
of the class of NP (assuming polynomial hierarchy). In general, it depends quite
a bit on which complexity class you're working in, whether the worst-case complexity
reduces down to the average-case complexity.
The second fact is more subtle but it's easiest to understand in the case of
SAT problems. If you take a satisfiability problem with a given number of
clauses $c$ for variables $n$, there is a ratio $q = c/n$
of variables to clauses. What's interesting is that as this ratio changes, two
really important things happen: whenever the number of variables is high and the
number of clauses is low the problem is relatively easy to satisfy because there
are many degrees of freedom. Whenever the ratio is low, the problem is very hard
to satisfy, often times impossible.
There's this intermediary point, which is typically described as a phase
transition from satisfiable to unsatisfiable, where the problem becomes
overdetermined in a way that makes it impossible to satisfy.
qₖ ≈ 4.27 easy · SAT
hard
easy · UNSAT
ratio q = c/n 2.00 P(satisfiable) 100% median DPLL decisions 37 regime underconstrained — easy, satisfiable ±J Ising spin glass with bond frustration set by the hardness at
q = 2.00. Near the transition the landscape is glassy and
the spins never settle; far from it they order almost immediately.
// A 28x28 grid of "spins", each either +1 or -1. Every spin is
// coupled to its four neighbors by a fixed random bond J that is
// either +1 ("these two want to agree") or -1 ("these two want to
// disagree"). The energy of the whole grid is
//
// E = -sum over neighbor pairs of J * spin_a * spin_b
//
// so a bond contributes low energy when it gets what it wants.
const GRID = 28 ;
const TEMPERATURE = 0.9 ;
const spin = new Int8Array ( GRID * GRID );
const couplingToRight = new Int8Array ( GRID * GRID );
const couplingToBelow = new Int8Array ( GRID * GRID );
// The lattice wraps around at the edges (a torus), so every site
// has exactly four neighbors.
const wrap = ( v ) => (v + GRID ) % GRID ;
const siteAt = ( row, col ) => wrap (row) * GRID + wrap (col);
function randomizeSpins ( ) {
for ( let s = 0 ; s < spin. length ; s++) {
spin[s] = Math . random () < 0.5 ? - 1 : 1 ;
}
}
// "Frustration" is the knob the slider controls. At 0 every bond is
// +1 and the grid is a plain ferromagnet: flipping toward your
// neighbors always helps, so it orders into big domains almost
// immediately. As frustration grows, more bonds are -1 and loops
// appear in which no assignment can satisfy every bond -- the
// defining feature of a spin glass. The energy landscape shatters
// into many competing local minima, which is the same picture as a
// SAT instance near the phase transition.
function rebuildCouplings ( frustration ) {
const pAntiferro = 0.5 * frustration;
for ( let s = 0 ; s < spin. length ; s++) {
couplingToRight[s] = Math . random () < pAntiferro ? - 1 : 1 ;
couplingToBelow[s] = Math . random () < pAntiferro ? - 1 : 1 ;
}
}
// One sweep = one update attempt per site, on average.
function metropolisSweep ( ) {
for ( let attempt = 0 ; attempt < GRID * GRID ; attempt++) {
// Pick a random site.
const row = ( Math . random () * GRID ) | 0 ;
const col = ( Math . random () * GRID ) | 0 ;
const here = siteAt

[truncated]
