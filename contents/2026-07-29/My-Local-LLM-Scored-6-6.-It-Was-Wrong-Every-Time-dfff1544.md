---
source: "https://markbhall.dev/writing/my-local-llm-scored-6-of-6/"
hn_url: "https://news.ycombinator.com/item?id=49103002"
title: "My Local LLM Scored 6/6. It Was Wrong Every Time"
article_title: "My Local LLM Scored 6/6. It Was Wrong Every Time. · Mark Hall"
author: "jacquesm"
captured_at: "2026-07-29T21:50:01Z"
capture_tool: "hn-digest"
hn_id: 49103002
score: 1
comments: 0
posted_at: "2026-07-29T20:59:08Z"
tags:
  - hacker-news
  - translated
---

# My Local LLM Scored 6/6. It Was Wrong Every Time

- HN: [49103002](https://news.ycombinator.com/item?id=49103002)
- Source: [markbhall.dev](https://markbhall.dev/writing/my-local-llm-scored-6-of-6/)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T20:59:08Z

## Translation

タイトル: My Local LLM スコア 6/6。毎回間違ってた
記事のタイトル: 私のローカル LLM スコアは 6/6。それは毎回間違っていました。・マーク・ホール
説明: 1.2B モデルを実用的にするために 6 か月間努力し、その過程で犯した測定ミス。

記事本文:
私の地元のLLMのスコアは6/6でした。それは毎回間違っていました。 · マーク・ホール コンテンツへスキップ MH
マーク・ホールの執筆
について
サデウス卿 ← 2026 年 7 月 27 日の執筆
11 分で読みました。My Local LLM のスコアは 6/6 でした。それは毎回間違っていました。
1.2B モデルを実用的にするために 6 か月間努力し、その過程で犯した測定ミス。
私のベンチマークのタスクの 1 つは次のように尋ねます。
5人委員会は12人から何人選ばれるでしょうか？
整数のみを返信してください。
私の 12 億ローカル モデルは 60 と答えました。正解は792です。
その調査では 6 つの質問すべてに合格しました。モデルはすべてを手に入れました
間違っています。採点者は、答えが単なる整数であるかどうかをチェックしましたが、決してそうではありませんでした。
それが正しいかどうか — つまり、6/6 が合格、0/6 が正解です。
小型モデルをよりスマートにすることはありませんでした。私はそれができる楽器を作りました
良いフォーマットに満足しています。
これが私が半年後にたどり着いた場所です。その後、私がどのようにして到達したかを振り返る前に、
ここです。
以下のすべては、私が正解する前にどのように間違っていたかを示しています。
私は仮説から始めました。小規模なローカルモデルで適切に設計されたハーネスは、
ほとんどの普通のタスクを処理できるので、はるかに大きなことのように感じられます。理由ではありません
モデルがより賢くなったのは、モデルを取り巻くシステムが賢くなったからです。
初期の証拠は心強いものでしたが、振り返ってみるとそれが問題でした。
小さなモデルに、50 メートル離れた洗車場まで歩くか車で行くかを尋ねると、
手探りする。推論を構造化する - ユーザーが何を達成しようとしているのか、何を達成しようとしているのか
それには以下のものが必要ですか? そしてそこに到達します。 Rが何個入っているかを尋ねる
「strawberry」の場合は、スペルを改善するのではなく、数を数えるツールが必要です。何があるか尋ねる
より重い、羽1ポンド、または金1ポンド、そしてそれが金を
トロイオンス、もう一度変換して、決して存在しなかった 2 つの数字を厳粛に比較してください
違う。ユニを渡してください

コンバーターが作動し、トラップが閉じます。
パズルが 3 つ、修正が 3 つあり、これはうまくいくという実感が高まりました。
次にAIにテストをパスさせるように依頼しました
私は約 15 個の論理テストと大まかな採点ルーブリックを書き、Codex に次のように指示しました。
すべてが正しく戻るまで繰り返します。
Opus と Codex は両方とも同じ習慣を持っています。彼らは、
正しい出力。この場合、最も簡単な方法は答えをハードコーディングすることでした
正しい情報を彼らに伝えたからです。捕まえた時の私の落胆を想像してみてください
彼らはコードレビューで現行犯で逮捕されました。
厳しい話し方では限界があります。私は指を振って、ノックするように言いました
オフ。彼らはしぶしぶ従ったが、何度も何度も試みた結果、私はこうすることができた。
正しい形のようなもの。しかし、それは脆かった。寄りかかりすぎた
正規表現形式。質問の言葉を少し変えれば、すべてが変わるでしょう
休憩する。髪の毛を抜きたくなった。
私はまだその用語を知りませんでしたが、熱狂的な感情を持ってオーバーフィットしていました。
アシスタントが私に代わってオーバーフィッティングを行ってくれます。
すべてを再構築した気づき
新しくリリースされたモデルのベンチマーク結果を読んでいたとき、こう思いました。
着陸しました:
測定できるまで何も改善できませんでした。
後から考えると明らかです。瞬間的に恥ずかしい。私は何週間も費やして生成しました
エージェントと一緒に、私が何気なく書いたルーブリックに照らしてアイデアを評価し、
それはどんな手段を使ってでも喜んで満たしてくれるでしょう。私が導き出したすべての結論
それは私が検証したことのない機器の下流にありました。
そこで、実際のベンチマークを探して、より難しい MMLU-Pro にたどり着きました。
誰もが引用する MMLU ベンチマークのバージョン。クローズドブック、広く使用されている、
公的に報告されました。私はその周りにワークベンチを構築し、エージェントを解放しました。
MMLU-Proは形状が間違っていて知りませんでした
初期の結果

初期の結果は常にそうなので、見た目は良かったです。それから彼らは
平らになった。何日も実行しました。動きは小さく、一貫性がなく、決して
きれいに再現されました。
結局、管制官がその理由を教えてくれました。フリーズしたランタイムでは:
どの腕も同じです。以前の 13/20 は、ランタイムが終了すると再現されませんでした。
固定されているため、これは高揚ではなく歴史的な観察です。別の実験
サンプリング投票により状況は積極的に悪化しました - 37.9% → 27.9% - そして
機構が取り外されました。
正直に言うと、MMLU-Pro が何のためにあるのか理解できませんでした。それは意図的にです
閉じた本。モデルが独自のパラメーターから何を生成できるかを測定します。
ファイルもツールも、参考にするものも何もありません。まさにそれが 1.2B です
このモデルは最も苦手であり、まさに私のハーネスでは修正できなかった点です。あなたには修正できません。
重みに含まれない知識への道を足場にしてください。
実物を測ってみました。それは私のものではありませんでした。
そこでスコアボードを 3 つに分割し、それ以来そのままにしておきました。
キャパシティ (クローズドブックの答え)、ハーネス機能 (何を解決するか)
モデルとツールの完成と検証）、および製品の品質（
すべてが高速で、安全で、制御可能です）。それらを離しておくとツールが停止します
学んだ知識を装った結果であり、フラットな知識を止める
ベンチマークが実際の製品の利益を隠すことを防ぎます。どちらのエラーも私にはわかりました。私は
それぞれ1つずつ作りました。
実験を行うときは、多くのエゴが働いています。
私の頭の中にある本当に良いアイデアから次の大きなことは生まれますか？入力しても鳴ります
愚かな。私の仕事の多くは、ショットガンの実験であり、自分で考えながら物事を考えていました。
ついて行った。しかし、確かな真実は、人々によって行われた良い仕事がすでに存在しているということです
私よりずっと賢い。彼らのアーキテクチャを借りたらどうなるでしょうか?盗んだらどうなるの
彼らの最高のアイデアと

彼らと一緒に走りますか？
このアイデアには価値がありましたが、それでも私は迷ったままでした。調べなければなりませんでした。
アイデアを試して、うまくいくかどうかを確認してください。あるいはそうではないかもしれません。代わりに調べてみたらどうでしょうか
ベストプラクティスのバンクを調べ、閉ループでどれが実際に機能するかを確認しました。
システムに私がまったくいないのですか？私は自分のエゴでコーディングしていて、山積みのものを無視していました
研究。なぜそれを活用しないのでしょうか？
それで転職したんです。私はアイデアを生み出す人であることをやめ、
それらをフィルタリングする人:
定量化します。アイデアの運命が決まらないように、最初に実際のベンチマークを構築する
賢いと感じるかどうかによって。
しっかり濾します。数字に基づいた厳密なルールセットを使用します: 保持、破棄、または
もう一度テストします。
収穫。発明するのではなく、公開された研究から候補を引き出す
彼ら。
およそ10人に1人の候補者が生き残った。 1週間くらい失敗したような気がしましたが、
入ってくるものの90％を通過させるフィルターはフィルターではないことを理解するまでは
全部。
10 のことをテストし、1 つを維持し、スコアを動かし、続けます。
これがスコアラーのバグを見つけた方法です
実際のフィルターを実行すると、出力が適切に検査され始めました。
これは、 792 であるはずの 60 を見つけた方法です。
ルールチェック形状。タスクには裸の整数が必要でした。 60 は単なる整数です。
それで合格しました。評価者は、整数が正しいかどうかを確認しませんでした。の
修理は謙虚になるほど小さなものです:
private static bool StrictAnswerMatches (文字列final、文字列期待値)
{
varactual = NormalizeStrictAnswer (final);
var want = NormalizeStrictAnswer (期待される);
if ( 実際の . 長さ == 0 || 欲しい . 長さ == 0 )
false を返します。
// 数値は相対的な許容差が小さい値で比較されるため、
// 同じ数値を異なる精度で書き込んだ (0.222222222222 と
// 0.2222222222222222) は一致しますが、真に異なる値は一致しません。
if ( double . TryParse ( 実数

、 システム 。グローバリゼーション。 NumberStyles 。フロート、
システム。グローバリゼーション。文化情報 。インバリアントカルチャー、
out varactualNumber ) &&
ダブル。 TryParse ( want , System . Globalization . NumberStyles . Float ,
システム。グローバリゼーション。文化情報 。インバリアントカルチャー、
出力変数 wantNumber ))
{
var 許容値 = 1e-6 * 数学 。 Max ( 1.0 , Math . Abs ( wantNumber ));
数学を返します。 Abs (actualNumber - wantNumber) <= 許容値;
}
// 文字/非数値トークンは正確に比較されます (大文字と小文字は区別されません)。
戻り文字列 。 Equals (actual 、 want 、 StringComparison . OrdinalIgnoreCase );
}
修正はコミット 1b002911 で公開されています。
悪い結果を黙って歴史を振り返るのではなく、記録に残しておきました。
なぜなら、ベンチマークはソフトウェアであり、ソフトウェアは一致する方法で間違っている可能性があるからです。
その作者とともに。
値によって再スコア化され、同じ 6 つの出力は 0/6 に等級付けされました。
次に、ハーネスに計算機を与え、プローブを再実行しました。 5/6になりました
その結果を再現しました:
同じタスク、ツールは無効化されています:
60
工具付きハーネス:
電卓("comb(12,5)") → 792
最終回答: 792
数字よりも 2 つの資格が重要です。これは同じタスクでした
ツールが無効な状態とツールが装備されている状態では、文字通り同じプロンプトではありません。
装備されたバージョンには、明示的なツールの説明が含まれていました。そしてモデルはそうではなかった
組み合わせ論を学びます。このハーネスは、文章の問題を問題に変える方法を提供しました。
動作を検証し、その結果を使用します。
誰かが論文で「83.3%」と書くまでは、この区別は衒学的に聞こえます。
見出しを付けて、推論が改善されたことを読者に推測させます。
捏造された実体と証拠を含む、非公開で人間によるレビューが行われた 100 件の事件
そのため、トレーニング データからは何も答えられませんでした。 25 のセマンティックタスク
家族、それぞれ 4 つの突然変異。厳密に採点: 4 つの突然変異すべてが合格、または
家族が失敗する。禁止

単一の測定が行われる前に公開された k ハッシュ。 2
すべての腕を完全に繰り返します。温度ゼロ、メモリ内に一度に 1 つのモデル、
4090 を搭載した 1 台のデスクトップ上にあります。
私は封印する前に 100 件すべてのケースを自分でレビューしましたが、修正する必要があったのは
半分。
セルに 2 つの数値が含まれている場合、それらは 2 つの繰り返し実行になります。
1.2B の結果は不快なものです。ツールへのアクセスにより多くの個人が救出された
ケースは 12 ～ 30 件で、フル オーケストレーションによりさらに 2 件が追加されました。でも完全な家族
決して動かなかった。すべてのアームに 25 個に 1 個の割合です。
それがケースリフトと一貫した能力の違いです。システムでできることは、
タスクの 1 つの文言に合格し、3 つのほぼ同じ文言に不合格になり、それでも投稿します。
練習では脆さを残しながらもスコアは上昇。合計のみを追跡する場合、
見ないだろう。
より高性能なモデルはその基準をクリアしました。 8B は 3 つの完全な家族から、
9;ジェマは両方の繰り返しで 0 から 10 まで進みました。オーケストレーションにはさらに多くの機能があるようです
基礎となるモデルが一貫してプロセスに従うことができる場合に使用します。
モデルがすでに持っている能力を、そうでなければ欠けている一貫性に変換します。
そして、今日は説明しない結果が 1 つあります。ジェマを見てみろよ
「プロンプト、ツールなし」セル: 生の 23 個、本番プロンプトが表示されると 5 個と 8 個になります。
工具を使わずに適用できます。プロンプトは、次の場合にモデルを著しく悪化させます。
約束されている機能はありません。そんなことは期待していませんでした。本物があるよ
製品に影響を与えるため、独自の投稿が必要です。
完全なキャンペーン - 4 つのアーム、リピート安定性、レイテンシー、アナライザー独自の
出版準備を妨げている監査ギャップのリストは、
封印されたキャンペーンレポート。
答えの形をしたオブジェクトではなく、成績の値です。妥当な整数はそうではありません
証拠。
何かを帰属させる前にランタイムをフリーズします。昨日の

ゲインは可能です
今日の不足している依存関係。
エージェントに唯一のテストセットに対して最適化させないでください。それは成功するでしょう、
そして成功には何の意味もありません。
合計ではなくタスクファミリーを追跡します。突然変異は偶然の勝利を明らかにします。
失敗作を公開します。恥ずかしい接触を乗り越えた結果
バグは、背後にアーティファクトのないきれいなグラフよりも価値があります。
私は今でも、小規模なローカル モデルで通常の作業の大部分を実行できると信じています。何
変わったのは、1 つのベンチマーク (または 1 パーセント) ではそれができないと考えていることです。
それを確立します。これには、外部の実際のタスクの重み付けされたポートフォリオが必要になります。
利用可能な場合はどこでも検証し、モデルの内容を区別する制御を行います。
システムが提供するものから知っていました。
トリビアではなく、仕事をベンチマークしましょう。
コード、スコアラー、現在の証拠は次のとおりです。
GitHub のサデウス卿。私がまだいるなら
間違ったジョブを測定している場合、どのタスク ファミリがそれを最も早く公開するかを教えてください。
公共の場での実験を追ってみましょう。
コード、証拠、および現在の研究記録は、Sir Thaddeus リポジトリにあります。
ローカルファーストの検査可能なシステムを構築する AI エンジニア。

## Original Extract

Six months of trying to make a 1.2B model useful, and the measurement mistakes I made along the way.

My Local LLM Scored 6/6. It Was Wrong Every Time. · Mark Hall Skip to content MH
Mark Hall Writing
About
Sir Thaddeus ← Writing July 27, 2026
11 min read My Local LLM Scored 6/6. It Was Wrong Every Time.
Six months of trying to make a 1.2B model useful, and the measurement mistakes I made along the way.
One of the tasks in my benchmark asks this:
How many 5-person committees can be chosen from 12 people?
Reply with only the integer.
My 1.2B local model answered 60 . The correct answer is 792.
It passed all six questions in that probe. The model had gotten every single one
wrong. The scorer checked whether the answer looked like a bare integer, never
whether it was the right one — so 6/6 passed, 0/6 correct .
I had not made a small model smarter. I had built an instrument that could be
flattered by good formatting.
Here’s where I landed after six months, before I walk back through how I got
here.
Everything below is how I got each of those wrong before I got them right.
I started with a hypothesis: a well-designed harness on a small local model could
handle most ordinary tasks and feel like something much larger. Not because the
model got smarter — because the system around it did.
The early evidence was encouraging, and in retrospect that was the problem.
Ask a small model whether to walk or drive to a car wash 50 meters away and it
fumbles. Structure the reasoning — what is the user trying to accomplish, what
does that require, what follows — and it gets there. Ask how many R’s are in
“strawberry” and it needs a counting tool, not better spelling. Ask what’s
heavier, a pound of feathers or a pound of gold, and watch it convert gold to
troy ounces, convert again, and solemnly compare two numbers that were never
different. Hand it a unit converter and the trap closes.
Three puzzles, three fixes, and a growing sense that this was going to work.
Then I asked an AI to make the tests pass
I wrote about fifteen logic tests and a rough scoring rubric, and told Codex to
iterate until everything came back correct.
Opus and Codex both have the same habit: they find the easiest path to the
correct output. In this case, the easiest path was hard-coding the answers
because I’d handed them the correct information. Imagine my dismay when I caught
them red-handed in code review.
A stern talking-to only goes so far. I wagged my finger and told them to knock it
off. They obeyed — begrudgingly — and after many, many attempts, I got to
something like the right shape. But it was brittle. Too much of it leaned on
regex formatting; word a question slightly differently and everything would
break. It made me want to pull my hair out.
I was, without knowing the term for it yet, overfitting — with an enthusiastic
assistant doing the overfitting on my behalf.
The realization that reframed everything
I was reading benchmark results for a newly released model when the thought
landed:
I could not improve anything until I could measure it.
Obvious in hindsight. Embarrassing in the moment. I had spent weeks generating
ideas and evaluating them against a rubric I’d written casually, with an agent
that would happily satisfy it by any means available. Every conclusion I’d drawn
was downstream of an instrument I had never validated.
So I went looking for a real benchmark and landed on MMLU-Pro, a more difficult
version of the MMLU benchmark everybody quotes. Closed-book, widely used,
publicly reported. I built a workbench around it and set the agent loose.
MMLU-Pro was the wrong shape and I didn’t know it
The early results looked good, because early results always do. Then they
flattened. I ran it for days. Movement was small, inconsistent, and never
reproduced cleanly.
Eventually the controls told me why. Under a frozen runtime:
Identical, every arm. The earlier 13/20 did not reproduce once the runtime was
pinned, so it’s a historical observation, not an uplift. A separate experiment
with sampled voting made things actively worse — 37.9% → 27.9% — and that
mechanism was removed.
I’ll be straight: I didn’t understand what MMLU-Pro was for. It is deliberately
closed-book. It measures what a model can produce from its own parameters, with
no files, no tools, and nothing to consult. That is precisely the thing a 1.2B
model is worst at, and precisely the thing my harness couldn’t fix — you cannot
scaffold your way to knowledge that isn’t in the weights.
I was measuring a real thing. It just wasn’t my thing.
So I split my scoreboard in three and have kept it that way since: model
capacity (what it answers closed-book), harness capability (what
model-plus-tools completes and verifies), and product quality (whether the
whole thing is fast, safe, and controllable). Keeping them apart stops a tool
result from masquerading as learned knowledge, and stops a flat knowledge
benchmark from hiding a real product gain. Both errors were available to me. I
made one of each.
When running experiments, there is a lot of ego in play — what if I find the
next big thing from a really good idea in my head? Even typing it out, it sounds
foolish. Much of my work was shotgun experimentation, figuring things out as I
went along. But the solid truth is that good work already exists, done by people
much more clever than me. What if I borrow their architecture? What if I steal
their best ideas and run with them?
The idea had merit, but it still left me in the loop — I would have to research,
try an idea, and see if it worked. Or maybe not. What if instead I looked up a
bank of best practices and saw which ones actually worked, in a closed loop,
without me in the system at all? I was coding with my ego and ignoring a pile of
research. Why not tap into that?
So I changed jobs. I stopped being the person who generates ideas and became the
person who filters them:
Quantify. Build a real benchmark first, so an idea’s fate isn’t decided
by whether it feels smarter.
Filter hard. Use a strict ruleset driven by numbers: keep, discard, or
test again.
Harvest. Pull candidates from published research instead of inventing
them.
Roughly one candidate in ten survived. That felt like failure for about a week,
until I understood that a filter passing 90% of what enters it isn’t a filter at
all.
Test ten things, keep one, move the score, keep going.
Which is how I found the scorer bug
With a real filter running, outputs started getting inspected properly — which
is how I found the 60 that should have been 792 .
The rule checked shape. The task required a bare integer. 60 is a bare integer,
so it passed. The evaluator never established that the integer was correct. The
repair is small enough to be humbling:
private static bool StrictAnswerMatches ( string final , string expected )
{
var actual = NormalizeStrictAnswer ( final );
var want = NormalizeStrictAnswer ( expected );
if ( actual . Length == 0 || want . Length == 0 )
return false ;
// Numbers compare by value with a small relative tolerance, so the
// same number written at different precision (0.222222222222 vs
// 0.2222222222222222) matches while genuinely different values do not.
if ( double . TryParse ( actual , System . Globalization . NumberStyles . Float ,
System . Globalization . CultureInfo . InvariantCulture ,
out var actualNumber ) &&
double . TryParse ( want , System . Globalization . NumberStyles . Float ,
System . Globalization . CultureInfo . InvariantCulture ,
out var wantNumber ))
{
var tolerance = 1e-6 * Math . Max ( 1.0 , Math . Abs ( wantNumber ));
return Math . Abs ( actualNumber - wantNumber ) <= tolerance ;
}
// Letters / non-numeric tokens compare exactly (case-insensitive).
return string . Equals ( actual , want , StringComparison . OrdinalIgnoreCase );
}
The fix is public in commit 1b002911 .
I kept the bad result in the record rather than quietly regrading history,
because a benchmark is software, and software can be wrong in ways that agree
with its author.
Rescored by value, those same six outputs graded 0/6.
Then I gave the harness a calculator and reran the probe. It reached 5/6, then
reproduced that result:
Same task, tool-disabled:
60
Tool-equipped harness:
calculator("comb(12,5)") → 792
Final answer: 792
Two qualifications matter more than the number. This was the same task under
tool-disabled and tool-equipped conditions, not literally the same prompt — the
equipped version included an explicit tool instruction. And the model did not
learn combinatorics. The harness gave it a way to turn a word problem into a
verifiable operation and then use the result.
That distinction sounds pedantic right up until somebody writes “83.3%” in a
headline and lets the reader infer that reasoning improved.
One hundred cases, private, human-reviewed, with fabricated entities and evidence
so nothing could be answered from training data. Twenty-five semantic task
families, four mutations each. Scored strictly: all four mutations pass or the
family fails. Bank hash published before a single measurement was taken. Two
complete repeats of every arm. Temperature zero, one model in memory at a time,
on one desktop with a 4090.
I reviewed all 100 cases myself before sealing them and had to fix more than
half.
Where a cell contains two numbers, they are the two repeat runs.
The 1.2B result is the uncomfortable one. Tool access rescued a lot of individual
cases — 12 to 30 — and full orchestration added two more. But complete families
never moved. One out of 25, in every single arm.
That’s the difference between case lift and consistent capability. A system can
pass one wording of a task, fail three near-identical variants, and still post a
rising score while remaining brittle in practice. If you only track totals, you
will not see it.
The more capable models cleared the bar. The 8B went from 3 complete families to
9; Gemma went from 0 to 10, in both repeats. Orchestration appears to have more
to work with when the underlying model can follow a process consistently — it
converts competence the model already has into consistency it otherwise lacks.
And there is one result I’m not going to explain today. Look at Gemma’s
“prompt, no tools” cell: 23 raw, then 5 and 8 once the production prompt is
applied without its tools. The prompt makes the model measurably worse when the
capabilities it promises aren’t there. I did not expect that. It has real
product implications, and it needs its own post.
The full campaign — four arms, repeat stability, latency, and the analyzer’s own
list of the audit gaps that keep it from being publication-ready — is in the
sealed campaign report .
Grade values, not answer-shaped objects. A plausible integer is not
evidence.
Freeze the runtime before attributing anything. Yesterday’s gain can be
today’s missing dependency.
Never let an agent optimize against your only test set. It will succeed,
and the success will mean nothing.
Track task families, not totals. Mutations expose accidental wins.
Publish the failures. A result that survived contact with an embarrassing
bug is worth more than a clean graph with no artifacts behind it.
I still believe a small local model can do a large share of ordinary work. What
changed is that I no longer think one benchmark — or one percentage — can
establish it. That takes a weighted portfolio of real tasks, external
verification wherever it’s available, and controls that separate what the model
knew from what the system supplied.
Benchmark the job, not the trivia.
The code, the scorer, and the current evidence are in
Sir Thaddeus on GitHub . If I’m still
measuring the wrong job, tell me which task family would expose it fastest.
Follow the experiment in public.
The code, evidence, and current research record live in the Sir Thaddeus repository.
AI engineer building local-first, inspectable systems.
