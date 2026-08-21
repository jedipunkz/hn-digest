---
source: "https://bastian.rieck.me/blog/2026/unix/"
hn_url: "https://news.ycombinator.com/item?id=49390066"
title: "LLMs Are Proof That Unix Won"
article_title: "LLMs Are Proof That Unix Won"
image: ""
author: "Pseudomanifold"
captured_at: "2026-08-21T16:20:47Z"
capture_tool: "hn-digest"
hn_id: 49390066
score: 2
comments: 0
posted_at: "2026-08-21T15:59:16Z"
tags:
  - hacker-news
  - translated
---

# LLMs Are Proof That Unix Won

- HN: [49390066](https://news.ycombinator.com/item?id=49390066)
- Source: [bastian.rieck.me](https://bastian.rieck.me/blog/2026/unix/)
- Score: 2
- Comments: 0
- Posted: 2026-08-21T15:59:16Z

## Translation

タイトル: LLM は Unix が勝った証拠である

記事本文:
Unix と「Unix に似た」オペレーティング システムについて初めて学んだとき、私はこう思いました。
興味をそそられた。私はこれまで Windows 3.1 のカラフルな世界しか知りませんでした。
日本の大工仕事と同じように、すべてが 1 つのブロックから彫られているように見え、目立った亀裂はありませんでした (ただし、日本の大工仕事は盤石であり、Windows 3.1 については同じことが真顔で言えるわけではありません)。
コマンドライン プロンプトの前に座ったときの私の驚きを想像してみてください。
初めて。点滅するカーソルが私に何かを入力するよう促した、そしてそこに
最初は、私がすでに行っていたことのいずれかを達成する明確な方法はありませんでした
コンピューターができることを知っていました。恐るべきパズル――私は夢中になってしまいました！おかげで
隣の村には驚くほど充実した図書館があることがわかった
Unix にはさまざまなフレーバーや進化上の親戚が存在すること、そして
彼らの動作の一部は POSIX などの標準によって成文化されています。
また、GNU と
これらすべてを成し遂げたハッカーの第一波の英雄的な努力
ロックに興味があると思われる世界が利用できるソフトウェア
すべてをダウンして改ざんを防ぎます。おお、勇敢な新世界よ！
しかし、私は粘りました。私は点滅し続けるプロンプトを見つめて、それを入力しました。たくさん
数か月後、 FreeBSD で寄り道した後、
私の仕事スタイルに合っているので、私は今でも熱心な Linux ユーザーです。
危険な状態にあり、直前にカーネルの更新を延期することがよくあります。
重要な締め切りはとてもスリルがありますが、通常は静かです
あらゆる困難から抜け出すことができる自分の能力について楽観的です。何
Linux のおかげで自分の能力を発揮できることに感謝しています
自己効力感。で
本質的に、それを使用することで私が経験する可能性のあるほとんどすべての痛みは、
かなりの部分、自傷行為です。それは持っているよりもはるかに良い気分です
次のiOSアップデートが行われることを祈ります
私のデバイスやその他のナンセンスを破壊することはありません。
この態度はしばしば無表情に見舞われる

見つめたり、いつもの「とにかく…」
それを理解していない人々、つまり、理解していない他のほとんどすべての人
大のオタクだよ、ニール
スティーブンソンのファン、または
豊富な自由時間に恵まれている。心地よいチクチク感の隣に
危険なことですが、Linux について私が最も興味を惹かれるのは、
デジタル粘土のように自分の目的に合わせて成形する能力。の多くは、
私が使用しているコマンドライン ツールはかなり前から存在しています 1
しかし、それらは依然として見事に機能し、次のようなことを行うことができます: 2
rg -t py "^\s*url =" \
| grep -Eo "(http|https)://[a-zA-Z0-9./?=_%:-]*" \
| awk -F/ '{print $3}' \
|並べ替え \
|ユニーク -c \
|ソート -rn
自然言語では、このコマンドの目的は抽出してカウントすることです。
変数に割り当てられた URL からの github.com などのドメイン名
現在のディレクトリとそのディレクトリ内のすべての Python ファイルにわたる名前付き URL
サブディレクトリ。もしこれが意味不明だと思われるなら、私の年下の短気よ
(ネオ)vim と活気に満ちた自分なら、昔なら「Linux は非常に便利だ」と言うでしょう。
ユーザーフレンドリー。友達に対しても非常にうるさいんです。」はい、
若い頃の私は職人のように敵を作るのが上手でした。 3
長年の知恵と円熟のおかげで、私は今では外交的になれるでしょう。
しかし、このようなコンピュータの使い方は今でも奇妙に思えます。
多くの人にとってはとても異質なものです。私の古代のワークフローを次のような人に説明しています。
最新の GUI のみに慣れている人は、高次元について説明するのと少し似ています。
フラットランドに住んでいる人:
せいぜい、彼らはあなたの言ったことを切り捨てる前に、丁寧に耳を傾けるでしょう。
少し奇妙で、昔のやり方に戻りつつあります。しかし、もう一度、私は粘り強く、
コンピュータはあなたに何かを提供するべきだという私の信念に固執しています
好きなものを構築できる汎用インターフェイス。
GUI は、何を推測する必要があるため、そのニーズを部分的にしか満たせません。
あなたがいつも通る道。対照的に、

、昔の Unix グレイビアード
ユーザーの行動を推測したり実証したりするのは無駄であることに気づきました。代わりに、
彼らは全員に、付着力のあるいくつかの小さなツールを装備することを選択しました。
ある哲学に対して：
1 つのことをうまく実行するプログラムを作成します。
連携して動作するプログラムを作成します。
テキスト ストリームを処理するプログラムを作成します。テキスト ストリームはユニバーサル インターフェイスであるためです。
数十年経った今でも、これは機能します。プログラムは大規模になり、さらに多くの
複雑ですが、ビデオなどの非常に特殊なケースではより便利です
編集 — しかし、多くのマシンの中核にはこの素晴らしいインターフェイスがあります
ほぼ無限の楽しみを提供します。 4 お互いの溝を広げるのではなく、
CLI ドワーフと GUI エルフ、しかし、何か予期せぬことが起こりました
つまり、大規模な言語モデルの開発が起こりました。
最初は入力ボックスのみをユーザーに表示します。
多くの人にとって、意図的に習慣を破ることになったのです。つまり、ここではノーでした
GUI は、希望する画像の種類を指定するのを待っています
作成します。このプロンプトは、あなたに大きな夢を持たせるというものでした。一部の人にとっては想像できると思いますが、
何をすべきかを教えてくれないプログラムでさえ、少し衝撃的だったに違いありません
それをやるのは前代未聞でした。 5
そして、初期の頃のプロンプトをそのままにして進歩が進みました。
自然言語による常に洗練されたクエリ。今、知る必要があるのではなく、
awk 、 grep 、およびその友人については、お気に入りの LLM に尋ねることができます。
URLからgithub.comなどのドメイン名を抽出してカウントしたい
これらは、すべての Python ファイルにわたって url という名前の変数に割り当てられます。
現在のディレクトリとそのサブディレクトリ。これをどうやって行うか
シェルコマンドのセット?
出力は非常に有能です。
grep -rhoP "url\w*\s*=\s*['\"]\Khttps?://[^'\"]+" --include="*.py" 。 \
| sed -E の#https?://##; s#/.*##' \
|並べ替え \
|ユニーク -c \
|ソート -rn
コマンドは多かれ少なかれ同じことを行います。私の

手作りのもので
ただし、rg は隠しディレクトリを自動的に無視します。
通常、コードを検索するときに実行したいことですが、私は提供しませんでした
そのコンテキストを LLM に反映します。さらに、-P はオペレーティング システム上で失敗します。
grep の BSD バリアントを使用します。繰り返しますが、LLM にはコンテキストが欠けています。
ただし、このコマンドはコピーしてターミナルに貼り付けると機能します。
CLI ツールの 1 つを自分で使用して、直接実行することもできます。
LLM が間の翻訳者として機能する、私のためのコマンドです。
自然言語コマンドと古代 Unix の呪文。
その意味で、LLM は Unix の哲学を体現しています。もちろん、これ
例えると、馬に簡単に乗って通り抜けることができるほど大きな穴があります。 LLM は
小さくもなく、一つのことだけをしているわけでもない――そう主張することもできるだろう。
彼らが行うことの中には、確かにうまくいかないものもある。これらの問題
それにもかかわらず、LLM はテキストがユニバーサル インターフェイスであることを理解しています。
ユーザーがコンピュータとの会話方法を学ぶ必要があるのではなく、
コンピューターが話しかけるようになりました。数年前なら、この概念は
完全に楽観的に見えた。 「テキストと
トークン化」は、汎用 AI モデルを構築するためのレシピです。しかし、
私たちは今、GUI への依存度を減らし、その代わりに私たちのツールに戻りつつあります。
人気の Unix 風のインターフェイス。
AI をめぐるあらゆる問題はあるものの、7 わたしたちは少なくとも、AI に存在することである程度の安心感を得ることができるかもしれません。
何十年も経って正しさが証明されました。テキストが最高の地位に君臨し、Unix が勝利しました。
全員ではないにしても、ほとんどが私の生徒よりも年上です。いくつかは
私よりもさらに年上です。テンプス・フギット 。 ↩︎
おそらくこれをより効率的に行う方法があるでしょう。させてください
それらについて知っています。私の CLI の知識がその知識であると主張するわけではありません。
最高。 ↩︎
自分の神経発散のせいにすることもできますが、たとえこれらのことがあったとしてもだと思います
酌量すべき事情として、私は使用します

d (さらに) 耐えられなくなる
私が若かった頃。ごめん！ ↩︎
少なくとも、それが私の楽しみの定義です。 ↩︎
この未練が今でも残っている原因ではないかと思います。
これらのモデルの第 1 世代のプロンプト/画像の例が多数あります。 ↩︎
このような絶対的な宝石を使用すると、次のようになります。
古代の森の苔むした石畳に佇む魔法使い、
木製の杖から光る呪文を唱えます。ダイナミックな映画のような
照明、Alphonse Mucha によるコンセプト アート、Octane Render、Unreal
エンジン 5、超現実的なディテール、8K、鮮明、傑作。
プロンプトエンジニアの場合は覚えておいてください
最も注目されている AI の仕事は何ですか? ↩︎
この記事では、
それらを議論する適切な場ではありません。 ↩︎
匿名でもいいですよ
フィードバックまたは
チップを
また、X（旧Twitter）をフォローしていただくと、
ブルースカイ、
マストドンとか
新しい投稿に関する通知を受け取るには。特に指定がない限り、すべてのコンテンツは、
Bastian Grossenbacher-Rieck によって作成され、ライセンスに基づいてライセンスされています。
クリエイティブ・コモンズ表示 4.0 国際ライセンス。

## Original Extract

When I first learned about Unix and “Unix-like” operating systems, I was
intrigued. I had only known the colorful world of Windows 3.1 so far.
Like Japanese carpentry , everything seemed to be carved out of one block with no apparent cracks (except that Japanese carpentry is rock solid, and the same cannot be said about Windows 3.1 with a straight face).
Imagine my surprise when I sat in front of a command-line prompt for the
first time. The blinking cursor dared me to enter something and there
was, at first, no obvious way to achieve any of the things I already
knew a computer could do. A formidable puzzle—I was hooked! Thanks to
a surprisingly well-stocked library the next village over, I learned
that there are different flavors or evolutionary cousins of Unix, and
that some of their behavior is codified by standards like POSIX .
I also learned about GNU and the
heroic efforts of the first waves of hackers who made all of this
software available to a world that seemed more interested in locking
down everything and preventing any tinkering. O brave new world!
But I persisted. I stared down the ever-blinking prompt and fed it. Many
moons later, after a detour with FreeBSD ,
I remain an avid Linux user since it suits my working style: I like to
live dangerously, often deferring kernel updates right before
important deadlines—what a thrill—and generally being quite
optimistic about my ability to get myself out of any jam. What
I appreciate is that Linux lets me exercise my
self-efficacy . In
essence, virtually all the pain I may experience by using it is, to
a large extent, self-inflicted . That feels so much better than having
to pray that the next iOS update
does not destroy my devices or some other nonsense.
This attitude is often met with blank stares or the usual “Anyway, …” by
people who just don’t get it, i.e., almost everyone else who is not
a huge nerd, Neal
Stephenson fan, or
blessed with an abundance of free time. Next to the nice tingling sense
of danger, the thing that entices me most about Linux is the
ability to mold it to my purposes like digital clay. Many of the
command-line tools I use have been around for quite some time now 1
but they still work admirably and allow me to do things like this: 2
rg -t py "^\s*url =" \
| grep -Eo "(http|https)://[a-zA-Z0-9./?=_%:-]*" \
| awk -F/ '{print $3}' \
| sort \
| uniq -c \
| sort -rn
In natural language, the purpose of this command is to extract and count
domain names like github.com from URLs that are assigned to variables
named url across all Python files in the current directory and its
subdirectories. If this reads like gibberish to you, my younger hothead
self, full of (neo)vim and vigor, would have hit you with the old “Linux is very
user-friendly; it’s just also super picky about its friends.” Yes,
younger me was adept at making enemies like a craftsman. 3
With the wisdom and mellowing of the years, I would now be diplomatic,
but it still strikes me as odd that this way of working with a computer
is so alien to many. Explaining my ancient workflow to someone who is
used to only modern GUIs is a bit like explaining higher dimensions to
someone inhabiting Flatland :
At best, they will politely listen before discarding what you said as
mildly odd and going back to their old ways. But again, I persisted and
stuck to my conviction that a computer should offer you
a general-purpose interface that enables you to build things you like.
GUIs can only partially sate that need since they need to guess what
path you are wont to take. By contrast, the Unix graybeards of yore
realized that it is futile to guess or railroad user behavior—instead,
they opted to equip everyone with a couple of smallish tools that adhere
to a certain philosophy :
Write programs that do one thing and do it well.
Write programs to work together.
Write programs to handle text streams, because that is a universal interface.
Decades later, this still works. Programs have become larger, more
complex, but also more convenient for highly-specific cases like video
editing—but at the core of many machines lies this wonderful interface
that offers nigh-limitless fun. 4 Instead of widening the gap between
the CLI dwarves and the GUI elves, however, something unexpected
happened, viz., the development of large language models .
Presenting at first nothing but an input box to the user, they
constituted a deliberate break in habits for many. Here, then, was no
GUI waiting for you to specify what type of picture you wanted to
create. The prompt was daring you to dream big. I imagine for some, it
must have been a bit shocking even—a program that does not tell you what to
do with it was unheard of. 5
And progress marched on, leaving the prompts of the early days 6 for
ever-refined queries in natural language. Now, instead of having to know
about awk , grep , and friends, one can just ask their favorite LLM:
I want to extract and count domain names like github.com from URLs
that are assigned to variables named url across all Python files in
the current directory and its subdirectories. How do I do this with
a set of shell commands?
The output is quite competent:
grep -rhoP "url\w*\s*=\s*['\"]\Khttps?://[^'\"]+" --include="*.py" . \
| sed -E 's#https?://##; s#/.*##' \
| sort \
| uniq -c \
| sort -rn
The commands do more or less the same thing. My hand-crafted one with
rg automatically ignores hidden directories, though, which is
typically what you want to do when searching code, but I did not provide
that context to the LLM. Moreover, -P will fail on operating systems
that use the BSD variant of grep . Again, the LLM lacks the context,
but this command will work when I copy and paste it into my terminal.
I could even use one of the CLI tools myself to make it directly execute
the command for me, with the LLM serving as a translator between
natural language commands and ancient Unix incantations.
In that sense, LLMs are embodying the Unix philosophy. Of course, this
analogy has holes so big you can easily ride a horse through. LLMs are
neither small nor do they do one thing—you could even argue that
some of the things they do, they certainly do not do well . These issues
notwithstanding, LLMs understand that text is the universal interface.
Instead of users needing to learn how to talk to the computer, the
computer now talks to you. A couple of years ago, this notion would have
seemed utterly optimistic. No one would have expected that “text and
tokenization” are the recipe for building general-purpose AI models. But
here we are, relying less and less on GUIs and instead going back to our
beloved Unix-like interface.
For all the problematic things around AI, 7 we may at least find some comfort in being
vindicated after so many decades: Text reigns supreme and Unix won.
Most, if not all of them, are older than my students. Some are
even older than me. Tempus fugit . ↩︎
There are probably ways to do this more efficiently. Please let me
know about them—I am not claiming that my CLI knowledge is the
best. ↩︎
I could blame my neurodivergence but I think even with these
extenuating circumstances, I used to be (even more) insufferable
when I was younger. Sorry! ↩︎
At least, that is my definition of fun. ↩︎
I suspect that this lingering feeling is still the reason for the
many example prompts/images of the first generation of these models. ↩︎
With such absolute gems like this one:
A wizard standing on a mossy stone path inside an ancient forest,
casting a glowing spell from his wooden staff. Dynamic cinematic
lighting, concept art by Alphonse Mucha, Octane Render, Unreal
Engine 5, hyperrealistic details, 8K, crisp, masterpiece.
Remember when prompt engineer
was the hottest AI job? ↩︎
I did not include them here on purpose because this article would
not be the right venue to discuss them. ↩︎
You can give me anonymous
feedback or a
tip .
Moreover, follow me on X (formerly Twitter) ,
Bluesky ,
or Mastodon
to get notified about new posts. Unless specified otherwise, all content
has been created by Bastian Grossenbacher-Rieck and is licensed under a
Creative Commons Attribution 4.0 International License .
