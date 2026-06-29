---
source: "https://htmx.org/essays/working-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=48720064"
title: "Working with AI"
article_title: "</> htmx ~ Working With AI: A Concrete Example"
author: "comma_at"
captured_at: "2026-06-29T15:07:54Z"
capture_tool: "hn-digest"
hn_id: 48720064
score: 2
comments: 0
posted_at: "2026-06-29T14:53:41Z"
tags:
  - hacker-news
  - translated
---

# Working with AI

- HN: [48720064](https://news.ycombinator.com/item?id=48720064)
- Source: [htmx.org](https://htmx.org/essays/working-with-ai/)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T14:53:41Z

## Translation

タイトル: AI の活用
記事のタイトル: </> htmx ~ AI の使用: 具体的な例
説明: このエッセイでは、カーソン グロス氏がハイパースクリプトの具体的なバグ修正について説明し、AI が役に立ったところ、不十分だったところ、そしてなぜ知識豊富な人間に情報を提供し続けることが複雑さを抑制できるのかを示します。

記事本文:
</ > HTMLx
ドキュメント
参照
例
話す
エッセイ
検索
スター
AI の活用: 具体的な例
私は一般的に、AI に対して両義的です。の開発にとって非常に強力なツールとなったことは疑いの余地がありません。
昨年もそうでしたが、多くの危険も伴います。
私たち個人にとっても（例：知性のゆっくりとした鈍化）、また集団的にも（例：環境問題、
ますます高価になるパーソナル コンピューティングなど)
「コードは安い」の中で、私は、開発者が AI に依存するようになり、開発者が AI に依存するようになる、魔法使いの弟子の問題について警告しています。
構築しているシステムで発生する問題を理解し、適切に対処します。
この記事では、私が AI と行った具体的なやり取りについて説明したいと思います。
一般的な AI の長所と短所を示し、実証するためのハイパースクリプト
特に魔法使いの弟子の問題（私はかろうじて回避しました）。
背景として、ハイパースクリプトは Web 用に解釈される代替スクリプト言語です。それは皮肉なことに、
すべて JavaScript で書かれています。
これは奇妙なソフトウェアです。実験として作成したときに、解析ルールの多くを意図的に破りました。
物事がどうなるかを見るために。
解析ロジックは解析要素上に同じ場所に配置されます
パーサーはプラグイン可能で、文法は動的に定義されます
プロパティ アクセスの複数の構文をサポートします。
これはほとんどのプログラミング言語に対して推奨できるアプローチではありませんが、この場合はかなりうまくいきました。
プロジェクト。
ソフトウェアで猫の皮を剥ぐ方法が実際に複数あることを示すもう 1 つのデモです。
私たちの物語は、ユーザーが 0.9.91 リリースにアップグレードするときにリグレッションを報告したときに始まります。以下の
式は正しく解析されなくなりました:
fetch `{% url ' trade :get_symbol_data' %}?symbol=${symbol}`

JSONとして
特に、as JSON はバインドが強すぎて、その前に文字列リテラルを JSON に変換しようとしていました。
ユーザーが期待したこと (そして以前に行っていたこと) を実行する代わりに、フェッチするように渡されました。つまり、
指定された URL と結果は JSON として処理されます。
この種のバインディングの競合は、解析における古典的な問題です。
ハイパースクリプトは xTalk スタイルの言語であるため、
英語の曖昧さの多くを引き継いでおり、この問題は英語ではさらに悪化しています。
最初にすべきことは、なぜこの回帰が起こったのかを調査することでした。
これは私が通常 AI に頼ることになる分野です。
私は Claude を使用していますが、根本原因を見つけるという素晴らしい仕事をしてくれました。0.9.91 では、私は過度に使用していました。
go コマンドのリファクタリングに積極的に取り組み、ロジックを fetch コマンドで再利用/共有します。
これらのコマンドの両方で使用する共通メソッド parseURLOrExpression() を抽出しましたが、その際に、
fetch コマンドの後の文法を誤って拡張して、一般的な式、えー、式を含めてしまいました。
as キーワードは式の中で意味を持ちます。これは変換式です。
型間の変換を可能にします。
x を Int として「42」に設定します
ただし、 as キーワードは fetch コマンドの修飾子でもあり、応答の変換方法を指示します。
https://hyperscript.org をテキストとして取得します
(おそらくこの事実により、口の中で少し吐いてしまうかもしれません。良いです。)
問題の核心は、リファクタリング中に不注意で、パーサーにフェッチ キーワードの後の式を解析させてしまったことでした。
これは、 as キーワードを fetch の修飾子として使用するのではなく、式として使用するようになりました。
クロードの助けのおかげで、私はこれを数分で理解することができました。そうする必要がある場合よりもはるかに速く、
自分で考えてください。
AIは問題の原因を見つけるのに非常に役立ちました。
問題を解決するにあたって

しかし、それははるかに弱かった。
ここで認めますが、私は怠け者で AI に解決策を求めていたので、その解決策について文句を言うのは少し気が引けます。
怠惰ですが、それでも一連の出来事は有益だと思うので、何が起こったのかを正確に見てみましょう。
最初に与えられた提案は、いわゆる「文字列のような」葉を最初に解析することでした。
その後、完全な式に戻ります。
これを返してください。 parseElement ( "stringLike" ) ||これ 。 requireElement ( "式" );
この修正により、ユーザーが提示した当面の問題は解決されるでしょう。
ただし、これは報告されたバグに非常に固有のものであり、誰かがフェッチのターゲットとして変数を使用する場合などの一般的なケースは修正されませんでした。
$url を JSON として取得する
私がこの提案を拒否した理由は、あまりにもハッキリしていて、十分に一般的ではないという理由です。
(ハイパースクリプト パーサーには有機的に提供されたハックが多数含まれているため、これが問題を引き起こす原因となった可能性があることに注意してください。
ケトルは黒です。）
提案された修正 2: 改善されたものの、不必要な複雑さ
2 番目の提案は、より興味深いものでした。パーサーに noConversions フラグを追加し、URL 解析の周囲に設定し、
AsExpression.parse が設定されている場合のベイル:
// AsExpression.parse()
if (parser . noConversions ) が返される。
これは、ハイパースクリプト パーサーを作成するため、多くのパーサー エンジニアを恐怖させるでしょう。
コンテキスト依存 。
ハイパースクリプト パーサーはすでにコンテキスト依存型でした。
この修正を見て少し考えてみると、私たちはすでにハッキーな状況依存型の機能を持っていることに気付きました。
パーサーに新しいフラグを導入せずにインフラストラクチャが必要でしたが、クロードはそれを見逃していました。
ハイパースクリプトパーサーの「フォロー」
ハイパースクリプト パーサーには、「follows」、つまり「上位」の解析要素によってフォロー トークンとして要求されるトークンの概念があります。
ハイパースクリプト パーサーは (やや

奇妙な) 再帰降下パーサー。これにより要素 (通常はコマンド) の解析が可能になります。
キーワードを「要求」すると、解析中に式がキーワードと照合されなくなります。
例として、 when 機能は区切り文字として または を使用します。
宣言内の論理接続詞として使用するよりも、
< div _ = "$x または $y が変更されると、それが私に反映されます" ></ div >
(多くのパーサー エンジニアが怒りでこのウィンドウを閉じるのが聞こえます。いいですね。)
この機能を使用すると、パーサーに新しいフラグを追加するのではなく、私たちが望んでいたものを実現できることがわかりました。
フォローとしてプッシュしてから式を解析し、それをフォローとしてポップすることができます。
これにより、変数などのほとんどの一般的な式は引き続き機能する一方で、AsExpression は解析できなくなります。
提案された解決策 3: 近いけど葉巻はダメ
私がこれをクロードに指摘すると、クロードは興奮して「絶対にその通りだ！」と言いました。そして出発します
このテクニックを使用してバグを修正します。
Claude は、正しいコードを parseURLOrExpression() に追加しました。これにより、何も追加せずに問題が全般的に修正されました。
パーサーインフラストラクチャ。
しかし、変更を確認しているときに、新しい修正が広すぎることに気付きました。フェッチとゴーの両方が共有されていました。
このメソッドは、修飾子を通知するためにフェッチのみを使用します。
既存の修正では、go コマンドでの変換式としての の完全に有効な使用も妨げられていました。
そこで、私は最終的な修正を自分で FetchCommand#parse() に実装しました。
パーサー。 PushFollow ( "as" );
{を試してください
var url = パーサー。 parseURLOrExpression();
最後に {
パーサー。ポップフォロー();
}
if ( parser . matchToken ( "as" )) {
...
ここでは特殊なケースを fetch コマンドのみに絞り込み、go 解析には影響を与えませんでした。
これがバグに対する私の最終的な答えとなりました。
その過程で、私はクロードにさまざまなケースに対するいくつかのテストを生成してもらいました。
良い存在がある

ハイパースクリプトのテスト スイートを作成し、クロードは、
問題が発生し、修正が適切に機能していることを確認しました。
別の分野では AI がうまく機能しているようです。
さて、このかなり平凡なバグ修正の話の何が興味深いのでしょうか?
AI がどこでうまくいったか、つまり調査とテストの作成を観察し、それを対比するのは興味深いと思います。
うまくいかなかった部分については、きれいな解決策を考え出しました。
私がハイパースクリプト パーサーとそのインフラストラクチャに精通していなかった場合、この修正により簡単に技術的な問題が発生した可能性があります。
プロジェクト内で発生している負債: 別のハッキーな解析のコーナーケース、パーサー上の別の状態など。
技術的負債は、証拠もなく断言します 1 が、指数関数的に増大しており、そのため非常に深刻です。
プロジェクト内でそれを最小限に抑えることが重要です。
このストーリーは、人間がどのように関与し、エージェントと連携し、根底にあるものをよく理解しているかを示しています。
インフラストラクチャは、エージェントを独自のデバイスに任せるよりも複雑さを制御する上ではるかに効果的です。
ハイパースクリプトのコードベースを見て、複雑さを制御することなどかつては不可能だったという考えを嘲笑する人もいるでしょう。
全くの考察。私はその意見に共感します。
ただし、この例では、具体的なシナリオで、問題を修正する際に複雑さが少なくとも少しは抑制されたことがわかります。
AI エージェントと協力している知識豊富な人間による、明らかに恥ずかしいバグです。
これは、魔法使いの弟子としてAIが提案する解決策を盲目的に受け入れるのではなく、
私は、より適切に適合する正しい解決策を要求する魔術師として行動していました（これが傲慢すぎると思わないことを願っています！）
既存のコードベースのアーキテクチャ。
私は問題を理解し、正しい解決策を見つけ、AI と協力してそれを達成し、検証することができました。
ソル

AI が生成したテストの助けを借りて使用します。
これは、開発者 (またはその他) が登場する、現在プッシュされているいくつかの形式のバイブ コーディングとは対照的であり、良い対照であることを願っています。
実際に何が起こっているのか理解していないことに誇りを持っています。
余談: AI と古い開発者
この経験を振り返っているときに、別のことが思い出されました。
私は今年 50 歳になる年配の開発者です。開発者が年齢を重ねるにつれて、現実には次のような傾向が見られます。
少なくともある程度は「速球を失う」。
実際、私にとって、これは次の 2 つのことを意味します。
以前ほど思い出せなくなった
以前のように長時間働くことができなくなった
AI はこれらの問題の両方に直接対処することが判明しました。
記憶力に関しては、以前は思い出せなかったものの、物事を非常に理解できるようになりました。
適切な、えー、プロンプトを表示してすぐに。 AI はこの点で非常に優れており、オープンとオープンの間で切り替えることができます。
それがなかった場合に比べて、プロジェクトのソースやプロジェクトの作業がはるかに効率的になりました。
長時間労働に関して言えば、AI は若い開発者である私でもそう思うほどの努力をすることができます。
ついて行くのが大変でした。これは、たとえば、自分のプロジェクト用に、より広範なテスト スイートを用意できることを意味します。
そうでなければ私がするよりも。
この場合にクロードが生成したテストを見ると、おそらく私が集めたものよりも広範囲にわたるものでした。
自分でやるエネルギー。
したがって、AI は、私が古い開発者として開発した 2 つの基本的な (相対的な) 弱点に対処しました。
その一方で、私はそれが私の全体的な知性のより一般的な退行も可能にしているのではないかと非常に心配しています。これ
いずれにせよ、年齢を重ねるにつれて自然に起こるものです。ただし、AI への依存によりこのプロセスが加速される可能性があるため、そうする必要があります。
振り返って言う

この話、私は正しいことをする前にどれだけクロードに頼っていたのか少し恥ずかしいです
私自身がくそったれ。
これは私自身がまだナビゲートしようとしている領域です。
この一連のやり取りを書きたいと思ったのは、良い点も悪い点もいくつか捉えられていると思ったからです。
AI によるコーディング支援。これは、ループ作業の中でかなり有能な開発者の価値を実証しました。
また、AI エージェントが提案した最初 (または 2 番目) の解決策を盲目的に受け入れる危険性も示しました。
問題を示唆します。
AI エージェントに関する独自の考えや戦略を立てる際に、この記事が役立つことを願っています。
これは夢の中で明らかになりました。
JavaScript の疲労:
ハイパーテキストへの憧れ
すでに手の中にある

## Original Extract

In this essay, Carson Gross walks through a concrete bug fix in hyperscript to show where AI helped, where it fell short, and why keeping a knowledgeable human in the loop is what kept complexity in check.

< / > htm x
docs
reference
examples
talk
essays
Search
star
Working With AI: A Concrete Example
I am, generally, ambivalent towards AI. There is no doubt has become a very powerful tool for development in the
last year, but it also comes with many dangers, both
for us individually (e.g. the slow dulling of our intellects) as well as collectively (e.g. environmental concerns,
increasingly expensive personal computing, etc.)
In “Code is Cheap(er)” , I warn about The Sorcerer’s Apprentice problem, where a developer becomes reliant on AI and is unable to
understand and properly address issues that come up in the systems they are building.
In this article I want to go through a specific interaction that I had with AI while maintaining
hyperscript to show the strengths and weaknesses of AI in general and to demonstrate
The Sorcerer’s Apprentice problem (which I narrowly avoided) in particular.
For some background, hyperscript is an alternative interpreted scripting language for the web. It is, ironically,
written entirely in JavaScript .
It is a strange piece of software: I intentionally broke many of the rules of parsing when writing it as an experiment
to see how things would work out.
Parsing logic is colocated on parse elements
The parser is pluggable, and the grammar is defined dynamically
It supports multiple syntaxes for property access .
It is not an approach I would recommend for most programming languages, but it has worked out pretty well for this
project.
Yet another demonstration that there are indeed multiple ways to skin the cat in software.
Our story begins when a user reported a regression when upgrading to the 0.9.91 release. The following
expression no longer parsed properly:
fetch `{% url ' trade :get_symbol_data' %}?symbol=${symbol}` as JSON
In particular, the as JSON was binding too tightly and trying to convert the string literal into JSON before it
was handed to fetch instead of doing what the user expected (and what it did previously) namely fetching the
given url with the results treated as JSON.
This sort of binding conflict is a classic problem in parsing.
Because hyperscript is an xTalk style language and
inherits many of the ambiguities of English, this problem is all the worse in it.
The first thing to do was to investigate why this regression occurred.
This is an area where I am typically going to lean on AI to help.
I use Claude, and it did an admirable job finding the root cause: in 0.9.91 I had been overly
aggressive in refactoring the go command to reuse/share logic with the fetch command.
I had extracted a common method for both of these commands to use, parseURLOrExpression() , but, in doing so, I
accidentally expanded the grammar after the fetch command to include the general expression , er, expression.
The as keyword has a meaning in expressions: it is a conversion expression ,
allowing you to convert between types:
set x to "42" as Int
But the as keyword is also a modifier of the fetch command, telling it how to convert the response:
fetch https://hyperscript.org as Text
(Perhaps this fact makes you throw up a little bit in your mouth. Good.)
The crux of the issue was that, inadvertently in the refactor, I had made the parser parse an expression after a fetch keyword
which was now consuming the as keyword as an expression, rather than allowing it to be a modifier for fetch .
With the help of Claude I was able to figure this out in a few minutes, much faster than if I had had to
figure it out on my own.
AI was very helpful in finding the cause of the problem.
In fixing the problem, however, it was much weaker.
I will admit here I was being lazy and asked AI for a solution, so complaining about those solutions feels a bit, well,
lazy, but I still think the string of events is informative, so let’s go through exactly what happened.
The first suggestion that was given was to parse what is called a “string-like” leaf first,
then fall back to a full expression:
return this . parseElement ( "stringLike" ) || this . requireElement ( "expression" );
This fix would have solved the immediate problem presented by the user.
However, it was very specific to the reported bug and wouldn’t have fixed the general case, such as if someone uses a variable as the target of a fetch:
fetch $url as JSON
I rejected this proposal because of this: too hacky and not general enough.
(Note that the hyperscript parser has plenty of organically supplied hacks in it, so this may have been the pot calling the
kettle black.)
Proposed Fix 2: Better But Unnecessary Complexity
The second proposal was more interesting: add a noConversions flag on the parser, set it around the URL parse, and have
AsExpression.parse bail when it is set:
// AsExpression.parse()
if ( parser . noConversions ) return ;
This will horrify many parser engineers because it makes the hyperscript parser
context-sensitive .
The hyperscript parser was already context-sensitive.
In looking at this fix and thinking for a second, I realized that we already had the hacky context-sensitive
infrastructure we needed without introducing a new flag on the parser, but Claude had missed it.
“Follows” In The Hyperscript Parser
In the hyperscript parser we have a notion of “follows” , that is, tokens that are claimed by a “higher up” parse element as a follow token.
The hyperscript parser is (a somewhat strange) recursive descent parser, and this allows a parse element (usually a command)
to “claim” a keyword, and expressions won’t match against them during parsing.
As an example, the when feature uses or as a separator rather
than as a logical connective in its declaration:
< div _ = "when $x or $y changes put it into me" ></ div >
(I can hear many parser engineers closing this window in anger. Good.)
It turns out that this feature could be used to achieve what we wanted: rather than adding a new flag to the parser
we could push as as a follow, then parse the expression, then pop it as a follow.
This would prevent the AsExpression from parsing, while still allowing most general expressions such as variables to work.
Proposed Fix 3: Close, But No Cigar
I pointed this out to Claude and, in a frisson of excitement, it told me that I was “absolutely right!” and set about
using this technique to fix the bug.
Claude added the correct code to the parseURLOrExpression() which fixed the issue generally without adding any additional
parser infrastructure.
However, as I was reviewing the change, I realized that the new fix was overly broad: both fetch and go shared
this method, but only fetch used as to signal a modifier.
The existing fix prevented the perfectly valid use of as conversion expressions in go commands as well.
So I implemented the final fix myself, in FetchCommand#parse() :
parser . pushFollow ( "as" );
try {
var url = parser . parseURLOrExpression ();
} finally {
parser . popFollow ();
}
if ( parser . matchToken ( "as" )) {
...
Here I narrowed the special case to only the fetch command, leaving go parsing unaffected.
This ended up being my final answer to the bug.
Along the way I had Claude generate some tests for the various cases.
There is a good existing test suite for hyperscript, and Claude did a good job of creating small, focused tests that showed the
problem and that the fix was working properly.
Another area AI appears to work well.
OK, so what is interesting about this fairly mundane bug fix story?
I think it is interesting to see where AI did well, namely in investigation and test creation, and to contrast that
with where it didn’t do so well: coming up with a clean solution.
If I had not been familiar with the hyperscript parser and its infrastructure this fix could have easily led to technical
debt being accrued in the project: another hacky parsing corner case, another bit of state on the parser, etc.
Technical debt, I assert without evidence 1 , grows exponentially, and therefpre it is very
important to minimize it in your projects.
This story shows how having a human in the loop, working with an agent and with a good understanding of the underlying
infrastructure, can be much more effective in controlling complexity than an agent left to its own devices.
Some people will look at the hyperscript code base and scoff at the notion that controlling complexity was ever
a consideration at all. I am sympathetic to that view.
However, in this example we can see in a concrete scenario how complexity was restrained, at least a bit, in fixing an
admittedly embarrassing bug, by a knowledgeable human working with an AI agent.
This is a situation where, rather than being a sorcerer’s apprentice and blindly accepting the solutions AI proposed,
I was acting as a sorcerer (I hope that’s not too arrogant to say!) demanding a correct solution that better fit the
existing codebase’s architecture.
I understood the problem and saw the correct solution and was able to work with AI to achieve it and then verify the
solution with the help of AI-generated tests.
This is in contrast, I hope a good contrast, with some forms of vibe coding currently being pushed in which developers (or whatever) appear
to pride themselves on not understanding what is actually going on.
Aside: AI & The Older Developer
Another thing occurred to me as I was going back over this experience.
I am an older developer, having turned 50 this year. As developers get older the reality is that we tend to
“lose our fastball”, at least to some extent.
Practically, for me, this has meant two things:
I am not able to remember as much as I used to
I am not able to work as long of hours as I used to
It turns out that AI directly addresses both of these issues.
With respect to memory, while I can’t remember everything I used to be able to, I can understand things again very
quickly with appropriate, er, prompting. AI is very good at helping me with this, and it lets me switch between open
source projects and work projects much more efficiently than if I didn’t have it.
With respect to the long hours, AI is able to grind in a way that, even as a young developer, I would have
had a difficult time keeping up with. This means, for example, I can have a much more extensive test suite for my projects
than I would have otherwise.
Looking at the tests that Claude generated in this case, they are more extensive than what I probably could have mustered
the energy to do myself.
So AI has addressed two fundamental (relative) weaknesses I have developed as an older developer.
On the other hand, I am very worried that it is also enabling a more general regression in my overall intelligence. This
is something that occurs naturally as you age anyway. AI reliance may accelerate this process however and I have to
say, looking back at this story, I’m a bit ashamed of how long I leaned on Claude before just doing the right thing
darned myself.
This is an area I am still trying to navigate myself.
I wanted to write up this series of interactions because I thought it captured some of the good and some of the bad
of AI assistance in coding. It demonstrated the value of a reasonably competent developer in the loop working
with an AI agent, and also showed the danger of blindly accepting the first (or second) solution that an AI agent
suggests to a problem.
I hope that it is useful to you as you develop your own thoughts and strategies around AI agents.
This was revealed to me in a dream.
javascript fatigue:
longing for a hypertext
already in hand
