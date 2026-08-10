---
source: "https://playtechnique.io/blog/learning-sed-or-using-claude.html"
hn_url: "https://news.ycombinator.com/item?id=49250994"
title: "Learning Sed or Using Claude?"
article_title: "Learning sed or using Claude?"
author: "gwynforthewyn"
captured_at: "2026-08-10T23:25:59Z"
capture_tool: "hn-digest"
hn_id: 49250994
score: 2
comments: 1
posted_at: "2026-08-10T22:52:26Z"
tags:
  - hacker-news
  - translated
---

# Learning Sed or Using Claude?

- HN: [49250994](https://news.ycombinator.com/item?id=49250994)
- Source: [playtechnique.io](https://playtechnique.io/blog/learning-sed-or-using-claude.html)
- Score: 2
- Comments: 1
- Posted: 2026-08-10T22:52:26Z

## Translation

タイトル: Sed を学習しますか? それとも Claude を使用しますか?
記事のタイトル: sed を学習しますか、それとも Claude を使用しますか?

記事本文:
遊びのテクニック
プロフェッショナルな DevOps コンサルタント
について
連絡してください
メンタリング
無料のドキュメント電子ブック
長文
技術ブログ
人前で話す
Sed を学習しますか、それとも Claude を使用しますか?
Linux システム管理者の就職面接で気に入っている質問があります。システムにログインしていると、次のことがわかります。
ls がインストールされていません。現在のディレクトリの内容を見つける 5 つの方法とは何ですか?
その下にあるものはすべて？それは私のインタビュー対象者に Linux について柔軟に考えることを強います。答えはこの一番下にあります
投稿...
年齢を重ねれば重ねるほど、柔軟な考え方に憧れます。私は Unix が私たちに与えてくれたツールチェーンを
継承: テキスト ストリーム、フィルター、プログラム可能なシェル。そこで、トレードオフは何なのかを自問してみようと思いました
これらのツールを学習するのと、最新の LLM ツールの 1 つを学習するのとでは異なります。ここで挑戦です。
andrew.RssInfo{タイトル: "PlayTechnique"、監督: "does-not-exist"、説明: "遊び方を学ぶ
より良いです。"}
このコード行は私のコードベースに何十回も現れます。 Dirパラメータをから変更しました
文字列を []文字列 に変換します。
リファクタリング オプションには、シェルまたはクロードの 2 つがあります。
当社の最新のシェルは、1 回限りのコマンドと単純なループのためのプログラム可能な環境です。
単一ファイル全体のリファクタリングに取り組む方法を知るには、適切なツールを知る必要があります。ここにあります
いくつかのオプション:
フィルタープログラム。テキスト編集およびレイアウトとしての UNIX の歴史により、多くのテキスト フィルター プログラムが存在します。
ツールチェーン: ストリームエディターを使用してツールチェーンを編集できるため、ここでは sed が最も適切です。
ファイルを 1 行ずつ実行できますが、対話型エディター セッションを開く必要はありません。
vi や emacs などの対話型テキスト エディターは、学習に時間がかかることで有名です。
デザインの素晴らしさ、または vs code のようなよりユーザーフレンドリーなものに対する長年の賞賛。
主要なスクリプト言語 (perl/ruby/python) はすべて

一回限りのコマンドを実行する何らかの方法をサポートする
この編集を実行できます。
`claude -p` スタイルの 1 回限りのプロンプト。
このコードベースに対して定義されたスラッシュコマンド、または ~/.claude ディレクトリで定義されたスラッシュコマンド。
sed を実装するときは、計画を立ててから、適切なマッチングを行うことから始めます。まず計画を立てるのは、
プログラミングの練習: テキスト エディタを開いて問題を解決し始めると、間違ったゲームに陥ってしまいます。スタート
考えることによって。それぞれのインスタンスが必要です
Dir: "string" は Dir: []string{"string"} になります。
これには、(a) 文字列 Dir: を含む行と一致する方法、(b) を挿入する方法、を知る必要があります。
一致した文字列の前と後に小さなテキストはありますか? (c) いくつかの sed フラグ。あまり具体的でないものからより具体的なものへと絞り込んでいきます
具体的な。ここからが始まりです (試合の始まりと終わりを明確にしました。また、多くの部分を編集しています)
例を明確にしたいので、出力します):
; sed -n 's/Dir:\(.*\)"/BEGINMATCH \1 ENDMATCH/p' rss_feed_test.go
rssInfo := andrew.RssInfo{タイトル: "PlayTechnique", BEGINMATCH ".", 説明: "より良いプレイを学ぶ。エンドマッチ}
ここからは、多かれ少なかれゲームのようなもので、使用される構文とパターンそのものを洗練させ、精神を訓練します。
テキスト内のパターンを確認する方法をもう少し説明します。私が持つ必要があった重要な洞察は、「」ということです。どのキャラクターにもマッチし、
一方、「[^"]」は引用符を除くすべての文字に一致します。
; sed -n 's/Dir: "\([^"]*"\)/Dir: []string{"\1}/p' rss_feed_test.go
rssInfo := andrew.RssInfo{タイトル: "PlayTechnique"、Dir: []string{"."}、説明: "より良い演奏を学ぶ。"}
これまで見たことがない場合は、そこには難解な構文が含まれています。キャプチャの詳細を理解するには
グループの場合、 man regex をチェックする必要がありました。正しくなるまでに約 6 分かかりました。
正規表現構文については、少し学習と努力が必要です。近道はありませんが、利点は同じです
正規表現言語

構文がずっと遡るため、多くの CLI ツールや言語で実装されています。
オリジナルのUNIXに。 Ken Thompson は正規表現を発明したわけではありませんが、初期に正規表現を発明した数少ない人物の 1 人です
プログラマーが実際にツールやプログラムでそれらを使用できるようにします。
claude -p "ファイル rss_feed_test.go で、RssInfo コンストラクターの Dir への引数を更新して、現在の値を唯一のメンバーとする文字列の配列にしてください。" --allowedTools "Read,Edit,Write,Bash"
「rss_feed_test.go」内の 4 つの「RssInfo」リテラル (37、92、113、128 行目) をすべて更新して、「Dir」を以前の値を含む「[]string」として渡すようにしました。
注: `RssInfo.Dir` はパッケージ内で依然として `string` として宣言されているため、`rss_feed_test.go` はそのフィールドの型も変更されるまでコンパイルされません。その変更を希望する場合はお知らせください。
もっと簡単です、そうです！しかし、このような単純なタスクであっても、クロードはそれが何か奇妙なことを言っているのかどうかを確認したかったので、
型が []string の場合、RssInfo.Dir の型は string であることがわかりました。ああ、クロード、決して変わらないでください。
Claude の実装は自然言語であるため、明らかにはるかに単純です。数枚のトークンがかかりました。私は
何も学ばなかったし、実際にツールに惑わされていましたが、高速で、書くのに数秒、5 ～ 10 秒かかりました。
実行する。
ツールを学習する必要なく利便性を得るコストは月額 20 ドルです。もっと安くする方法があります
ハーネスを持っていますが、それはクロード・コードに支払う金額です。
sed の実装は明らかにより複雑です。正しく理解するには時間がかかり、考えました。それは私にいろいろなことを教えてくれました、
正規表現マッチングの詳細を思い出し、「The C」の楽しいセクションを思い出しました。
プログラミング言語」で、文字を 1 つのリテラルで記述する必要がないことについて読みました。
文字 (上記の「\(」は 1 文字です)。
sed の実装により、

正規表現の構文は次のとおりです。
多くのツールに共通です。このプロセスでは、いくつかのドキュメントを読み、トークンがどのように描写されているかについて考えさせられました。
コードを書くのに役立つようなもの。
結果を得ることが目標の場合、ほとんどの人にとってクロード実装の方が有利です。いくつかの種類の
雇用主としても起業家としても、クロードは奇跡のように見えるはずです。
私にとっての目標は、小さな問題のデバッグを練習し、精神的な強さを養うことです。遅くまで過ごしてしまった
夜間は本番環境のデバッグとサービスの復元を行います。私はある問題を診断した者です
「開発者」が問題に直面してそれを調査する方法を理解していなかったために、バイブコード化されたアプリが作成されました。
彼は脳のスイッチをオフにしていた。だから私は多少の摩擦は気にしません、ツールを練習するのは気にしません、そして私は大好きです
コマンドラインなのですべて
私にとって便利で強力だと感じます。
私にとっての目標は、楽しむことでもあります。コマンドラインを編集して新しいことを学ぶことはやりがいがあります。良くなっている
私の仕事はやりがいがあります。本番環境に移行してデバッグできるという自信が得られるのはやりがいがあります。
ああ、完全を期すために、 ls を使用せずにツリーを再帰する 5 つの方法を次に示します。
見つけます。これも非常に明白ですが、私は見つけるのが大好きで、人々がそれを知っていると聞きたいのですが、
tar cvf /dev/null 。これはもう少し邪悪で、rsync または
ファイル再帰を中核とする他のツール、
cp-Rl 。 /tmp は、クリーンアップが必要なハードリンクを作成するため、少しずるいです。
設定が正しい場合にのみ機能します。
rsync -rvn 。 /tmp/nonexistent/ はより本質的なものであり、
そしてもちろん、いつでも Ruby や他の言語を尋ねることもできます。
Ruby -e 'Puts Dir.glob("**/*")'

## Original Extract

play technique
Professional DevOps Consultant
About
Contact Me
Mentoring
Free Documentation E-book
Long Writing
Tech Blog
Public Speaking
Learning Sed or Using Claude?
I have a favourite linux sys-admin job interview question: you're logged in to a system and it turns out that
it doesn't have ls installed. What're five ways to find the contents of the current directory and
everything below it? It forces my interviewee to think flexibly about linux. Answers at the bottom of this
post...
The older I get, the more I admire thinking flexibly. I use the toolchain that unix gave us all as an
inheritance: text streams, filters, a programmable shell. So, I thought I'd ask myself what the tradeoffs are
of learning those tools vs learning one of the modern LLM tools. Here's the challenge.
andrew.RssInfo{Title: "PlayTechnique", Dir: "does-not-exist", Description: "Learning to play
better."}
This line of code appears a dozen times in my codebase. I changed the Dir parameter from
string to []string .
I have two refactoring options: shell or claude.
Our modern shells are programable environments for one-off commands and simple loops.
Knowing how to tackle a refactoring across a single file means that you have to know the right tools. Here's a
few options:
A filter program. There are many text filter programs due to unix's history as a text editing and layout
toolchain: sed is the most appropriate here, as the s tream ed itor allows us to edit the
file line by line but without needing to open an interactive editor session.
An interactive text editor such as vi or emacs that is notorious for a steep learning curve followed by
years of admiration for how well it's designed, or something more user friendly like vs code.
The major scripting languages (perl/ruby/python) all support some method of firing a one-off command that
can perform this edit.
A one-off prompt of the `claude -p` style.
A slash-command defined for either this codebase or defined in your ~/.claude directory.
When I'm implementing sed, I start with a plan, then by getting the matches right. Planning first is a
programming practice: if you start fixing a problem by opening a text editor, you're in the wrong game. Start
by thinking. I need each instance of
Dir: "string" to become Dir: []string{"string"} .
This requires knowing (a) How do I match a line containing the string Dir: , (b) how do I insert a
little text before that matched string and after it? (c) a few sed flags. I refine from less specific to more
specific. Here's the start (I've made it clear where my match begins and ends; I'm also redacting a lot of
output, because I want the example to be clear):
; sed -n 's/Dir:\(.*\)"/BEGINMATCH \1 ENDMATCH/p' rss_feed_test.go
rssInfo := andrew.RssInfo{Title: "PlayTechnique", BEGINMATCH ".", Description: "Learning to play better. ENDMATCH}
From here, it's more or less a game, refining the syntax used and the pattern itself, training my mind a
little more how to see patterns in text. The key insight I needed to have is that '.' matches every character,
whereas '[^"]' matches every character except a quote.
; sed -n 's/Dir: "\([^"]*"\)/Dir: []string{"\1}/p' rss_feed_test.go
rssInfo := andrew.RssInfo{Title: "PlayTechnique", Dir: []string{"."}, Description: "Learning to play better."}
There's some arcane syntax in there if you've not seen it before. To figure out one detail of capturing
groups, I had to check man regex . Took about six minutes to get right.
The regex syntax needs a little learning and effort. There's no shortcut, but the upside is that the same
regex language is implemented in a lot of cli tools and in languages, because the syntax goes all the way back
to the original UNIX. Ken Thompson didn't invent regular expressions, but he was one of a few early
programmers to actually use them in tools and programs.
claude -p "In the file rss_feed_test.go, please update the argument to Dir in the RssInfo constructor to be an array of strings with the current value as the only member" --allowedTools "Read,Edit,Write,Bash"
Updated all four `RssInfo` literals in `rss_feed_test.go` (lines 37, 92, 113, 128) to pass `Dir` as a `[]string` containing the previous value.
Note: `RssInfo.Dir` is still declared as a `string` in the package, so `rss_feed_test.go` won't compile until that field's type is changed too — let me know if you want me to make that change.
Much simpler, yes! But even in a simple task like this, Claude wanted to be sure it said something whacky, so
it told me that RssInfo.Dir is of type string when the type is []string. Oh, Claude, never change.
The Claude implementation is obviously much simpler as it's natural language. It cost a handful of tokens. I
didn't learn anything, and I was actually misled by my tool, but it was fast, a few seconds to write and 5-10
to execute.
The cost of getting convenience without needing to learn the tooling is $20/month. There are cheaper ways to
have harnesses, but that's what I pay for Claude Code.
The sed implementation is obviously more complex. It took time and thought to get right. It taught me things,
reminding me of details in regular expression matching, reminding me of the delightful section in "The C
Programming Language" in which I read about how a character doesn't need to be written with one literal
character (above, '\(' is a single character).
The sed implementation will cross-pollinate to my other code writing, because the regular expression syntax is
common to a lot of tools. The process made me read some documentation, think about how tokens are delineated,
the kind of stuff that helps you write code.
If the goal is getting to the result, the claude implementation wins for most people. To some kinds of
employer or entrepreneur, Claude's got to seem like a miracle.
To me, the goal is practicing the debugging of small problems, developing mental fortitude. I've spent late
nights debugging production and restoring service. I've been the person who diagnosed the problems in a
vibe-coded app because the "developer" didn't understand how to be stumped by a problem and investigate it,
he'd turned his brain off. So I don't mind the bit of friction, I don't mind practicing my tools, and I love
the command line, so it all
feels useful and powerful to me.
The goal's also, for me, having fun. Command line editing and learning new things is rewarding; getting better
at my craft is rewarding. Being confident that I can get into production and debug it is rewarding.
Oh, and for completeness' sake, here's 5 ways to recurse a tree without using ls :
find . is also pretty obvious but I love find and want to hear that people know it,
tar cvf /dev/null . is a bit more devious and establishes a pattern you can reuse with rsync or
some other tool that has file recursion as a core of it,
cp -Rl . /tmp is a bit cheaty because it creates hard links that'll need cleaning up, so it
only works if your setup is correct,
rsync -rvn . /tmp/nonexistent/ is more in the spirit of things,
and of course you could always just ask ruby or another language:
ruby -e 'puts Dir.glob("**/*")'
