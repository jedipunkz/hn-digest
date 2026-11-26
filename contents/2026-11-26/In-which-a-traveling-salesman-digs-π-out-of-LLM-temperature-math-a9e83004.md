---
source: "https://www.hexbites.dev/stories/2026-11-25-pi-from-simulated-annealing.html"
hn_url: "https://news.ycombinator.com/item?id=50651261"
title: "In which a traveling salesman digs π out of LLM temperature math"
article_title: "In which a traveling salesman digs π out of LLM temperature math | Hexcellent Expeditions and Byte-Sized Adventures"
image: ""
author: "kromych"
captured_at: "2026-11-26T04:54:13Z"
capture_tool: "hn-digest"
hn_id: 50651261
score: 1
comments: 0
posted_at: "2026-11-26T04:03:03Z"
tags:
  - hacker-news
  - translated
---

# In which a traveling salesman digs π out of LLM temperature math

- HN: [50651261](https://news.ycombinator.com/item?id=50651261)
- Source: [www.hexbites.dev](https://www.hexbites.dev/stories/2026-11-25-pi-from-simulated-annealing.html)
- Score: 1
- Comments: 0
- Posted: 2026-11-26T04:03:03Z

## Translation

タイトル: 巡回セールスマンが LLM 温度計算から π を掘り出す
記事のタイトル: 巡回セールスマンが LLM 温度計算から π を掘り出す |素晴らしい探検とバイトサイズの冒険
説明: セールスマンは旅をし、熱くなり、冷め、アルキメデスにぶつかる

記事本文:
巡回セールスマンが LLM 温度計算から π を導き出す |素晴らしい探検とバイトサイズの冒険
素晴らしい探検とバイトサイズの冒険
コンピューティング マシンを動かしているものについて考える
巡回セールスマンが LLM 温度計算から π を導き出す様子
おそらく、過去数年間のある時点で、LLM API ペイロードに「temperture: 0.7」と入力したことがあるでしょう。あるいは、AI スタックに何らかの「温度」があると聞いて、それを「人々が多額のお金をもらっている奇妙なもの」としてファイルしたことがあります。
そんなに奇妙ではありませんが、私の言葉を鵜呑みにしないでください。大まかに言って、雰囲気は変わりません。温度を 0.0 まで下げると、モデルは想像力を失い、いたるところでお役所仕事が見えてきます。 1.8 まで上げると、文の途中で言語が切り替わり、まだ開いていない括弧が閉じられます。
では、なぜこの投稿には \(\pi\) と巡回セールスマンも含まれているのでしょうか?
その「温度」には過去があるから。 19 世紀の統計力学から迷い込み、機械学習に定着しました。
そして、この方程式がどこから来たのかを理解すると、それを悪用する方法が見え始めます。今日の悪用: LLM トークン サンプリングの背後にあるボルツマン温度計算を取得し、それを単位円上の NP 困難な最適化問題に渡し、そのための小さなソルバーを作成し、瓦礫の中から数値 \(\pi\) を掘り出します。
不正行為はしないと約束します。 （「不正行為」が何を意味するかについてはすぐに説明します。）
では、LLM の温度とは何でしょうか?
巡回セールスマン、円周上
毛玉を冷やして丸くする
では、ツアーをどのように調整するのでしょうか?
いや、なぜ3.1409で止まってしまったのでしょうか？
第 2 幕、最も要求の厳しい読者向け: 熱からのパイ
徹底的に冷やさないでください
1 組の都市は 2 トークンのソフトマックスのように動作します
要求の厳しい読者が持ち帰るもの
では、LLM の温度とは何でしょうか?
で

前方パスの終わりに、言語モデルは生の正規化されていないスコアの大きなベクトルを語彙内のトークンごとに 1 つずつ渡し、単語の選択はユーザーに任せます。人々はこれらのスコアをロジットと呼び、\(z_1, z_2, \dots, z_V\) と書きます。
スコアをサンプリングできる確率に変換するために、softmax を実行し、途中でロジットを正の数 \(T\) で割ります。
分子 \(e^{z_i / T}\) を少し見てください。
大学で熱力学の授業を受講し、それ以来その記憶を抑圧している方には申し訳ありませんが、ボルツマン分布という記憶を取り戻そうとしています。温度 \(T\) の熱浴中で、エネルギー \(E(s)\) を備えた状態 \(s\) になれる物理システムを考えてみましょう。かなりの確率でその状態で見つかります。
2 つの方程式は一致します。 LLM は、トークンの負のロジット (\(-z_i\)) がそのエネルギーを測定しているかのように見せかけます。おそらくトークンは快適な谷の底に座っていて、奇妙なトークンははるか上の崖にしがみついています。
\(T\) がゼロになると、小さな数で除算すると、最大ロジットと残りのロジットの間の差が爆発し、ソフトマックスは argmax に崩壊します。システムは貪欲になります。坂を下って最も近い井戸に転がり込み、そこに留まります。
\(T\) が巨大になると、\(z_i / T\) はすべて \(0\) に向かって押しつぶされ、全員が \(e^0 = 1\) になります。分布はフラットになります。高エネルギー、低エネルギー、システムは気にすることをやめました。
その間にはミックスが入ります。依然として低エネルギー状態が好まれますが、十分な熱ジッターがあるため、時折上り坂を登ってさらに悪い状態に陥ったり、次の谷に迷い込んだり、尾根を越えてより良い状態を見つけたりすることがあります。
冶金学者はこのことを長い間知っていました。金属の塊を原子が飛び交うまで加熱し、ゆっくりと冷やすと、原子は低エネルギーの結晶格子に落ち着きます。いいね

速すぎる（急冷する）と、たまたま立っていた場所、欠陥などすべてが凍ってしまいます。
コンピューターサイエンスはそのトリックを借用し、それをシミュレーテッドアニーリングと呼びました。そして、巡回セールスマンは鉄と同じように熱を受け取ります。
巡回セールスマン、円周上
巡回セールスマン問題 (TSP) では、\(N\) 個の都市を取得し、すべての都市を 1 回訪問して最初に戻る最短の閉ループを求めます。
全体的にNP難しいです。 \(N = 250\) の都市の場合、\((250 - 1)! / 2 \およそ 6.5 \times 10^{489}\) の個別のツアーがあります。ビッグバン以来、ナノ秒ごとに 10 億回のツアーをチェックしていたとしても、傷つくことはなかったでしょう。
しかし、すべての \(N\) の都市を単位円 (半径 \(R = 1\)) 上に配置したらどうなるでしょうか?
そうすれば、最高のツアーがどのようなものであるかがわかります。都市は凸型の境界上にあるため、最短のツアーはそれ自体を横切ることはなく、それ自体を横切らないツアーは、内接する凸型 \(N\)-gon をたどりながら順番に周縁を歩きます。
\(N\) が大きくなるにつれて、その多角形の外周 \(L\) が円周 \(2\pi R = 2\pi\) 上を下から這い上がってきます。最短のツアーの長さを \(2\) で割ると、\(\pi\) がポップします。
今のうちに疑ってかかるべきだ。これらの都市をサークルに追加するにはどうすればよいですか? let theta = 2.0 * PI * rng(); と書くとします。 (theta.cos(), theta.sin()) そして \(\pi\) を「発見」したことを発表したら、このタブを閉じてください。 PI 、 sin または cos を使用して円を構築することは、最も文字通りの意味での循環論法になります:D
したがって、不正行為の禁止ルールが必要です。都市ジェネレーターとソルバーは PI 、 sin 、 cos 、または atan2 に触れることはできません。 main() の最後に PI を 1 回インポートして、どれだけ近づいたかを出力します。
ここで疑問が残ります。 + 、 - 、 * 、 / 、および sqrt で武装した円上で均一にサンプリングするにはどうすればよいでしょうか?拒絶サンプリングを使用する場合:
ランを選ぶ

正方形 \([-1, 1]^2\) 内に均一に点 \((x, y)\) を配置します。
\(r^2 > 1\) (円盤の外側) または \(r^2 \およそ 0\) (死点) の場合は、それを捨てて再試行します。
それ以外の場合は、 \(r = \sqrt{r^2}\) で割って、リムに向かってまっすぐ押し出します。
円盤は回転対称であるため、円盤から均一な点をその境界に投影すると、円上に均一な分布が得られます。 (伝統的に、代数は演習として残すことが求められています。) 三角関数は必要ありませんでした。
(円盤内に到達する正方形のサンプルの割合は、たまたま \(\pi / 4\) であることはわかっています。ソルバーには \((x, y)\) 座標の入った袋が渡され、それらの間を歩く最短の方法を見つけ出す必要があります。拒否された点とその比率は残ります。したがって、私は実験を汚染されていないと呼んでいます。)
毛玉を冷やして丸くする
\(N\) 個のランダムな都市を生成された順序で接続しても、リングは得られません。円を横切って前後にスライスするコードの毛玉が得られます。 \(N = 150\) の場合、最初のツアーの長さは \(L \約 190\) となり、\(\hat{\pi} \約 95\) になります。アルキメデスはまだ心配していません。
それでは、ボルツマンにツアーについて指摘してみましょう。
シミュレートされたアニーリングの 1 つのステップは次のようになります。
ツアーに小さなランダムな変更を提案し、エネルギー変化 \(\Delta E = L_{\text{new}} - L_{\text{old}}\) を計算します。
\(\Delta E \le 0\) (ツアーが短くなった) なら、参加してください。
\(\Delta E > 0\) (ツアーが長くなった) の場合は、一様乱数 \(u \in [0, 1)\) を描画し、次の場合にはそれを採用します。
\(T \leftarrow \alpha T\) のように温度を少し下げます。 \(1\) のすぐ下に \(\alpha\) を設定します (0.999994 など)。
では、ツアーをどのように調整するのでしょうか?
明らかなランダムな動きにより都市が入れ替わります。ツアーでそのうちの 2 つを選択し、その位置を交換します。
しばらくは機能します。しかし、ツアーの2つの長い区間が終わったとき、

プレッツェルのように交差しているため、2 つの個別の都市を交換すると、状況が良くなる前に状況が悪化することがよくあり、低温では「悪化」は受け入れられません。都市スワップは、結び目のあるローカル ミニマムに陥ってしまいます。
そこで、代わりに 2-opt move を使用します。つまり、2 つのインデックス \(i < j\) を選択し、スライス Tour[i..=j] を反転します。
サークルではなぜこれほどうまくいくのでしょうか?最初に 4 つの都市の名前を付けます。 B = Tour[i] および C = Tour[j] はスライスの両端をマークし、ツアー内では A が B の前にあり、 D が C の後に続きます。次に、エッジ A -> B と C -> D が中央で交差しているところを想像してください。
前 (交差): 反転後 [B .. C]:
A C A ------- C
\/
\ / ===>
\/
×
／＼
／＼
／＼
D B D ------- B
B から C にすべてを反転すると、交差するペア (A, B) と (C, D) が交差しないペア (A, C) と (B, D) に置き換えられます。三角不等式 により、凸四角形の 2 つの対角線を合計すると、対辺 2 つ以上になります。したがって、2-opt リバーサルでエッジのペアを解除すると、ツアーが短くなります。
そして、\(\Delta E\) を知るためにすべての \(N\) エッジを合計する必要はありません。 2 つのエッジが出力され、2 つのエッジが入力され、算術演算では \(O(1)\) が実行されます。
毛玉の凍結について読むこともできますし、 \(\hat{\pi}\) が 95 から約 3.140 まで下落する間に毛玉が凍結してきれいなリングになる様子を観察することもできます。見てますよ。 (これは相転移によく似ています。物理学者は屁理屈を言うかもしれません。)
以下のブラウザでライブ シミュレーションが実行され、上記のトライグフリー サンプリングと \(O(1)\) のエネルギー更新が行われます。 4 つのプリセットが付属しています。
Well-Annealed : 高温でガスを含んだツアーが冷却され、交差していない多角形になります。
熱すぎる (\(T = 1.5\)): LLM の温度が大幅に上昇しました。マゼンタの点滅は、受け入れられた上り坂の動きを示しますが、ツアーは落ち着きません。
クエンチ : 都市交換による即時フリーズ (\(T = 10^{-6}\))。貪欲なデコディみたいに

ループにはまっていないと、もつれたプレッツェルにロックされてそこに留まります。
300 都市 : 都市が増え、冷却スケジュールが遅くなり、 \(\hat{\pi}\) は約 3.1412 になります。
プリセット 3 は少しチートです。フリーズ 2-opt は円に引っかからないため、温度に応じて動きを変更します。そこでのもつれたツアーはまだどこかに交差点があり、交差点を元に戻すとツアーが短縮されるため、2-optは暑さの助けを借りずにポリゴンまでずっと下り坂を滑り降ります。都市の交換は、都市がリング上に並んでいない通常のマップでの 2-opt と同様に、スタックします。
Well-Anneled が円形に落ち着いたら、Reheat を押してサーマル スパイク (\(T = 1.5\)) を注入し、結晶が冷える前にバラバラになるのを観察します。
ブラウザには 60 フレーム/秒で物理現象が表示されます。 250 万のモンテカルロ ステップでは、より高速なものが必要なので、ブラウザーを終了してコンパイルします。
このプログラムには依存関係が必要ないため、 Cargo.toml も作成しません。これは 1 つのファイル tsp_pi.rs に収まります。これは、小さな xorshift64 PRNG、トライグフリーのサークル サンプラー、および 2-opt ボルツマン ループです。
std :: f64 :: consts :: PI を使用します。 // エラーをチェックするために最後にのみ使用されます
構造体 Rng ( u64 );
impl Rng {
fn next_u64 (& mut self ) -> u64 {
mut x = self .0 とします。
x ^= x << 13 ;
x ^= x >>7 ;
x ^= x << 17 ;
自分自身 .0 = x ;
×
}
/// [0.0, 1.0) の均一浮動小数点数
fn next_f64 (& mut self ) -> f64 {
( self .next_u64 () >> 11 ) as f64 * ( 1.0 / ( 1u64 << 53 ) as f64 )
}
/// [0,bound) の均一な整数
fn next_usize (& mut self ,bound : usesize ) -> usesize {
( self .next_u64 () as use ) % バウンド
}
}
#[派生(クローン、コピー)]
構造体ポイント{
x : f64 、
y:f64、
}
impl ポイント {
fn dist (self , other : Point ) -> f64 {
// ハイポトではなく単純な sqrt: 座標は [-1, 1] 内に留まります。
// したがって、正方形

オーバーフローできず、ホット ループでの sqrt のコストが低くなります
let ( dx , dy ) = ( self .x - other .x 、 self .y - other .y );
( dx * dx + dy * dy ) .sqrt ()
}
}
/// 単位円上の均一にランダムな点「n」個をサンプリングします (R = 1)
/// PI、sin、cos には一切触れません。
fn sample_unit_circle ( n : usesize , rng : & mut Rng ) -> Vec < Point > {
let mut pts = Vec :: with_capacity ( n );
while pts .len () < n {
let x = rng .next_f64 () * 2.0 - 1.0 ;
y = rng .next_f64 () * 2.0 - 1.0 ; とします。
r2 = x * x + y * y とします。
if r2 > 1e-12 && r2 <= 1.0 {
r = r2 .sqrt (); とします。
pts .push (ポイント { x : x / r , y : y / r });
}
}
ポイント
}
fn Tour_length (都市 : & [ ポイント ], ツアー : & [ 使用サイズ ]) -> f64 {
let n = ツアー.len();
( 0 .. n )
.map (| i | 都市 [ ツアー [ i ]] .dist ( 都市 [ ツアー [( i + 1 ) % n ]]))
.sum()
}
fnsolve_tsp_annealing (
都市: & [ ポイント ]、
t_start : f64 、
t_end : f64 、
手順：使用、
rng : & mut Rng 、
) -> Vec < 使用量 > {
n = 都市 .len (); とします。
mut ツアーを実行します: Vec < usesize > = ( 0 .. n ) .collect ();
let mut current_len = Tour_length (都市, & ツアー);
let mut best_tour = ツアー .clone();
mut best_len = current_len とします。
// 指数冷却係数: T_k = t_start * alpha^k
let alpha = ( t_end / t_start ) .powf ( 1.0 / ステップとして f64 );
mut temp = t_start とします。
ステップイン 0 .. ステップ {
[切り捨てられた]
さびた

[切り捨てられた]

## Original Extract

The salesman travels, heats up, cools down, bumps into Archimedes

In which a traveling salesman digs π out of LLM temperature math | Hexcellent Expeditions and Byte-Sized Adventures
Hexcellent Expeditions and Byte-Sized Adventures
Musings of what makes the computing machines tick
In which a traveling salesman digs π out of LLM temperature math
You’ve probably typed temperature: 0.7 into an LLM API payload at some point in the last few years. Or you’ve heard there’s some “temperature” in the AI stack and filed it under “weird stuff people get paid tons of money for”.
It isn’t that weird, but don’t take my word for it. Roughly, and going on vibes: turn the temperature down to 0.0 and the model gets unimaginative and sees red tape everywhere. Turn it up to 1.8 and it switches languages mid-sentence and closes brackets it didn’t open yet.
So why does this post also have \(\pi\) and a traveling salesman in it?
Because that “temperature” has a past. It wandered in from 19th-century statistical mechanics and made itself at home in machine learning.
And once you see where the equation comes from, you start seeing ways to abuse it. Today’s abuse: we take the Boltzmann temperature math behind LLM token sampling, hand it to an NP-hard optimization problem on a unit circle, write a small solver for it, and dig the number \(\pi\) out of the rubble.
I promise I won’t cheat. (We’ll get to what “cheating” means in a minute.)
So what is temperature in an LLM?
A traveling salesman, on a circle
Cooling a hairball into a circle
How do you tweak a tour, though?
Oh no, why did it stop at 3.1409 ?
Second act, for the most demanding reader: Pi from the heat
Don’t cool it all the way
One pair of cities behaves like a two-token softmax
What the demanding reader takes home
So what is temperature in an LLM?
At the end of a forward pass, a language model hands you a big vector of raw, unnormalized scores, one per token in its vocabulary, and leaves picking the word to you. People call those scores logits, and we’ll write them as \(z_1, z_2, \dots, z_V\).
To turn the scores into probabilities we can sample from, we run them through softmax , dividing the logits by a positive number \(T\) on the way in:
Look at the numerator for a second: \(e^{z_i / T}\).
If you took a thermodynamics class in college and have since suppressed the memories, sorry, I’m about to bring one back: the Boltzmann distribution . Take a physical system that can be in a state \(s\) with energy \(E(s)\), in a heat bath at temperature \(T\). You’ll find it in that state with probability:
The two equations match. An LLM pretends that a token’s negative logit (\(-z_i\)) measures its energy. Likely tokens sit at the bottom of a comfortable valley, and the weird ones cling to a cliff way above.
When \(T\) goes to zero, dividing by a small number blows up the gap between the biggest logit and the rest, and softmax collapses into argmax . The system gets greedy: it rolls downhill into the nearest well and stays there.
When \(T\) gets huge, the \(z_i / T\) all get squashed toward \(0\), and \(e^0 = 1\) for everybody. The distribution goes flat. High energy, low energy, the system has stopped caring.
In between, you get a mix. Low-energy states are still preferred, but there’s enough thermal jitter to occasionally hop uphill into a worse state, wander into the next valley, and maybe find something better over the ridge.
Metallurgists have known this for a long time. Heat a lump of metal until its atoms are scrambling around, cool it down slowly , and the atoms settle into a low-energy crystal lattice. Cool it too fast ( quench it) and they freeze wherever they happened to be standing, defects and all.
Computer science borrowed the trick and called it simulated annealing . And a traveling salesman takes to heat much the way iron does.
A traveling salesman, on a circle
In the Traveling Salesman Problem (TSP) , you get \(N\) cities and you want the shortest closed loop that visits all of them once and comes back to the start.
In general it’s NP-hard. For \(N = 250\) cities there are \((250 - 1)! / 2 \approx 6.5 \times 10^{489}\) distinct tours. If you’d been checking a billion tours per nanosecond since the Big Bang, you wouldn’t have made a dent.
But what if we put all \(N\) cities on a unit circle (radius \(R = 1\))?
Then we know what the best tour looks like. The cities lie on the boundary of a convex shape, so the shortest tour doesn’t cross itself, and a tour that doesn’t cross itself walks around the rim in order, tracing an inscribed convex \(N\)-gon.
As \(N\) grows, the perimeter \(L\) of that polygon creeps up on the circumference \(2\pi R = 2\pi\) from below. Divide the shortest tour length by \(2\) and out pops \(\pi\):
You should get suspicious about now. How am I putting those cities on the circle? If I write let theta = 2.0 * PI * rng(); (theta.cos(), theta.sin()) and then announce that I “discovered” \(\pi\), you should close this tab. Using PI , sin or cos to build the circle would be circular reasoning in the most literal sense :D
So we need a no-cheating rule: the city generator and the solver don’t get to touch PI , sin , cos or atan2 . We import PI once, at the end of main() , to print how close we got.
Which leaves a question: how do you sample uniformly on a circle armed with + , - , * , / and sqrt ? With rejection sampling :
Pick a random point \((x, y)\) uniformly inside the square \([-1, 1]^2\).
If \(r^2 > 1\) (outside the disk) or \(r^2 \approx 0\) (dead center), throw it away and try again.
Otherwise, push it straight out to the rim by dividing by \(r = \sqrt{r^2}\):
The disk is rotationally symmetric, so projecting uniform points from the disk onto its boundary gives a uniform distribution on the circle. (Tradition demands that I leave the algebra as an exercise.) We didn’t need any trig.
(The fraction of square samples that land inside the disk happens to be \(\pi / 4\), I know. The solver is handed a bag of \((x, y)\) coordinates and has to work out the shortest way to walk between them. The rejected points and that ratio stay behind. So I’m calling the experiment uncontaminated.)
Cooling a hairball into a circle
Connect our \(N\) random cities in the order they were generated and you don’t get a ring. You get a hairball of chords slicing back and forth across the circle. For \(N = 150\) that first tour has length \(L \approx 190\), which makes \(\hat{\pi} \approx 95\). Archimedes isn’t worried yet.
So let’s point Boltzmann at the tour:
One step of simulated annealing goes like this:
Propose a small random change to the tour and compute the energy change \(\Delta E = L_{\text{new}} - L_{\text{old}}\).
If \(\Delta E \le 0\) (the tour got shorter), take it.
If \(\Delta E > 0\) (the tour got longer ), draw a uniform random number \(u \in [0, 1)\) and take it anyway if:
Nudge the temperature down a little: \(T \leftarrow \alpha T\), with \(\alpha\) just under \(1\) (something like 0.999994 ).
How do you tweak a tour, though?
The obvious random move swaps cities: pick two of them in the tour and exchange their positions.
It works for a while. But when two long stretches of the tour are crisscrossed like a pretzel, swapping two individual cities usually makes things worse before it makes them better, and at low temperature “worse” doesn’t get accepted. City swap ends up stuck in a knotted local minimum .
So we’ll use a 2-opt move instead: pick two indices \(i < j\) and reverse the slice tour[i..=j] .
Why does that work so much better on a circle? Name four cities first: B = tour[i] and C = tour[j] mark the two ends of the slice, A precedes B in the tour, and D follows C . Now picture the edges A -> B and C -> D crossing in the middle:
Before (crossed): After reversing [B .. C]:
A C A ------- C
\ /
\ / ===>
\ /
X
/ \
/ \
/ \
D B D ------- B
Reversing everything from B to C replaces the crossing pair (A, B) and (C, D) with the uncrossed pair (A, C) and (B, D) . By the triangle inequality , the two diagonals of a convex quadrilateral add up to more than two opposite sides do. So uncrossing a pair of edges with a 2-opt reversal makes the tour shorter.
And we don’t have to add up all \(N\) edges to know \(\Delta E\). Two edges leave, two edges come in, and the arithmetic takes \(O(1)\):
You can read about freezing hairballs, or you can watch one freeze into a clean ring while \(\hat{\pi}\) slides from 95 down to about 3.140 . I’d watch. (It looks a lot like a phase transition . A physicist might quibble.)
A live simulation runs in your browser below, with the trig-free sampling and the \(O(1)\) energy update from above. It comes with four presets:
Well-Annealed : the hot, gassy tour cools down into an uncrossed polygon.
Too Hot (\(T = 1.5\)): your LLM with temperature cranked way up. Magenta flashes mark the uphill moves that got accepted, and the tour doesn’t settle.
Quenched : instant freeze (\(T = 10^{-6}\)) with city swap. Like greedy decoding stuck in a loop, it locks into a tangled pretzel and stays there.
300 Cities : more cities and a slower cooling schedule, for a \(\hat{\pi}\) around 3.1412 .
Preset 3 cheats a little: it changes the move along with the temperature, because frozen 2-opt doesn’t get stuck on a circle. A tangled tour there still has a crossing somewhere, and undoing a crossing shortens the tour, so 2-opt slides downhill all the way to the polygon without help from the heat. City swap does get stuck, the way 2-opt would on an ordinary map where the cities don’t line up on a ring.
Once Well-Annealed has settled into a circle, hit Reheat to inject a thermal spike (\(T = 1.5\)) and watch the crystal fall apart before it cools back down.
The browser shows the physics at 60 frames per second. For 2.5 million Monte Carlo steps we want something faster, so we’ll leave the browser and compile it.
The program doesn’t need any dependencies, so we won’t even create a Cargo.toml . It fits in one file, tsp_pi.rs : a small xorshift64 PRNG, the trig-free circle sampler, and the 2-opt Boltzmann loop.
use std :: f64 :: consts :: PI ; // Used ONLY at the very end to check our error
struct Rng ( u64 );
impl Rng {
fn next_u64 ( & mut self ) -> u64 {
let mut x = self .0 ;
x ^= x << 13 ;
x ^= x >> 7 ;
x ^= x << 17 ;
self .0 = x ;
x
}
/// Uniform float in [0.0, 1.0)
fn next_f64 ( & mut self ) -> f64 {
( self .next_u64 () >> 11 ) as f64 * ( 1.0 / ( 1u64 << 53 ) as f64 )
}
/// Uniform integer in [0, bound)
fn next_usize ( & mut self , bound : usize ) -> usize {
( self .next_u64 () as usize ) % bound
}
}
#[derive(Clone, Copy)]
struct Point {
x : f64 ,
y : f64 ,
}
impl Point {
fn dist ( self , other : Point ) -> f64 {
// Plain sqrt rather than hypot: the coordinates stay within [-1, 1],
// so the squares can't overflow, and sqrt costs less in a hot loop
let ( dx , dy ) = ( self .x - other .x , self .y - other .y );
( dx * dx + dy * dy ) .sqrt ()
}
}
/// Sample `n` uniformly random points on the unit circle (R = 1)
/// without ever touching PI, sin, or cos.
fn sample_unit_circle ( n : usize , rng : & mut Rng ) -> Vec < Point > {
let mut pts = Vec :: with_capacity ( n );
while pts .len () < n {
let x = rng .next_f64 () * 2.0 - 1.0 ;
let y = rng .next_f64 () * 2.0 - 1.0 ;
let r2 = x * x + y * y ;
if r2 > 1e-12 && r2 <= 1.0 {
let r = r2 .sqrt ();
pts .push ( Point { x : x / r , y : y / r });
}
}
pts
}
fn tour_length ( cities : & [ Point ], tour : & [ usize ]) -> f64 {
let n = tour .len ();
( 0 .. n )
.map (| i | cities [ tour [ i ]] .dist ( cities [ tour [( i + 1 ) % n ]]))
.sum ()
}
fn solve_tsp_annealing (
cities : & [ Point ],
t_start : f64 ,
t_end : f64 ,
steps : usize ,
rng : & mut Rng ,
) -> Vec < usize > {
let n = cities .len ();
let mut tour : Vec < usize > = ( 0 .. n ) .collect ();
let mut current_len = tour_length ( cities , & tour );
let mut best_tour = tour .clone ();
let mut best_len = current_len ;
// Exponential cooling factor: T_k = t_start * alpha^k
let alpha = ( t_end / t_start ) .powf ( 1.0 / steps as f64 );
let mut temp = t_start ;
for step in 0 .. steps {
[truncated]
rustc

[truncated]
