---
source: "https://anarc.at/blog/2026-08-18-people-vs-ai-overlords/"
hn_url: "https://news.ycombinator.com/item?id=49363959"
title: "People vs. the AI Overlords"
article_title: "The people vs the AI overlords - anarcat"
image: "https://anarc.at/apple-touch-icon.png"
author: "meetpateltech"
captured_at: "2026-08-19T17:19:22Z"
capture_tool: "hn-digest"
hn_id: 49363959
score: 1
comments: 0
posted_at: "2026-08-19T16:47:19Z"
tags:
  - hacker-news
  - translated
---

# People vs. the AI Overlords

- HN: [49363959](https://news.ycombinator.com/item?id=49363959)
- Source: [anarc.at](https://anarc.at/blog/2026-08-18-people-vs-ai-overlords/)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T16:47:19Z

## Translation

タイトル: 人間 vs. AI の覇者
記事のタイトル: 人民 vs AI の覇者 - anarcat
説明: アナーキャット

記事本文:
人民 vs AI の覇者
以前のこのシリーズ: LLM の四騎士
黙示録。
私の (Debian) 共同開発者 Russ Allbery は oss-security への投稿で
「オープンソースソフトウェア[OSS]は、次のような問題に直面している」と述べた。
長い間蓄積されてきたモチベーションの危機。」彼の主張
本質的には、大規模言語モデル (LLM 1 ) が
既存の OSS コミュニティの危機はさらに悪化します。彼にとって、それはコードの洪水です
レビューはあるが、それは人々の欲求によって変わると彼は主張する。
セキュリティの問題などもあります。
ラスは正しいと思うが、もっと大きな何かがあると私は主張する
ここで行われているオープンコミュニティよりも、それは全体に関するものです
コンピューティングの分野。このプレッシャーは、関係なく私たち全員にかかっています。
オープンソース ソフトウェアに取り組んでいるかどうか。
ワークフローで LLM を使用する人々は、その方法を根本的に変えました。
プログラミングは避けたいと主張する人でも機能します
バイブコーディング 。ここで一人の貧弱なメンテナを取り上げて申し訳ありません:
それはあなたではありません、ブライアン、あなたはたくさんいる中の一例にすぎません。しかし、これは
これらのモデルの今日の一般的な使用法は次のとおりです。
それが完了したら、/code-review を使用してクロードをスポーンさせます
サブエージェントが新しいコードを完全にレビューします。通常これで見つかります
いくつかの問題、さらには「メイン」Claude インスタンスでは発生しなかった問題
検証中に見つけます。通常は /code-review を実行し続けます
問題がなくなるまで、問題を見つけて修正した後、何度も繰り返します。
どれも残っています。
それが何を意味するのか少し考えてみましょう。これは次の目的で構築された自動化です
そうでない場合は、数十人のエージェントを起動して数分間問題に対処する
数時間の GPU コンピューティング時間を並行して実行します。これは本質的にカップルです
データセンター ラック内の多数のシェルフで、電力が完全に最大化され、
クールで、かわいい小さな /code-review コマンドの背後に抽象化されています。
著者、

ここで、「人類が人類を引っ張る可能性がある」と懸念しているのは当然である。
ラグアウトし、API の価格設定を要求する」というコードワードかもしれません。
「実費に近い金額を請求する」。ブライアンもリップを支払う
環境的および社会的コストへの貢献ですが、それらは主に
抽象化されているので、その話はここでも脇に置いておきましょう。
いずれにしても、以前に議論したとおりです。
しかし、明らかに、この働き方には（外部化された）コストがかかります。
控えめに言っても。
何十年もの間、私の仕事は無料のオープンソースに焦点を当ててきました。
ソフトウェア。私は長い間、次のような独自のオペレーティングシステムを使用するのをやめてきました。
Windows または Mac、そしてその切り替え前でさえ、私は主に無料のものを使用していました
これらのプラットフォーム上のソフトウェアは、部分的に原則から外れていますが、次の理由もあります。
私はあまりにも貧乏でした。つまり、私の取引ツールは無料であり、私は無料で構築します
ツールも一緒に。
まるで逆戻りしているような気がする：私が学生だった頃、数千年
以前、私のクラスメートはコンパイラにアクセスできず、不思議に思っていました。
Borland のようなコンパイラを購入するためにどうやってお金を捻出するでしょうか
またはマイクロソフトの。オペレーティング システムにコンパイラが組み込まれていた
(当時は FreeBSD) だったので、私にとっては問題ありませんでした。彼らにとって、
かなりの出費でしたが、少なくともそれらの出費（あるいはそれ以上）
プログラムのいかがわしい調達）は一発取引でした。
30 年早送りすると、ソフトウェアはレンタルされ、月額料金を支払います。
Adobe の Photoshop と Microsoft のオフィス スイートを、料金を支払ったのと同じように利用できます。
Netflix、Disney+、または Spotify 2 .そして今、あなたは何十ものものを追加する必要があります（そうでない場合）
それに加えて、LLM にアクセスするための毎月のクレジットが数百ドル) かかります。
では、何かをするためにはお金を払わなければならないのでしょうか？これがピークです
私たちの仕事の強化：まず彼らは私たちの仕事を盗んで自分たちの仕事を訓練します
そして彼らはそれを私たちに売り戻して利益を得ます。
AI は、全員ではないにしても、エンジニアとして私たちの仕事のためにやって来ます。
ナラットに

私。ここしばらく、我が国の雇用市場は悪化しています。
仕事が減れば、給料も減ります。仕事を探している優秀なエンジニアがたくさんいます
そしてくだらない仕事を見つけて、それでも働きながら探します。
これは偶然ではありません。 3 私たちエンジニアは大きな力を持っていますが、そうではありません
組織化されていますが、それはほんの数組合の距離にあります（簡単です！）。技術
君主たちはこれを知っているので、私たちの職業を直接攻撃しています。
彼らが制御できるモデルをトレーニングして使用することを私たちに強制することによって。
プログラマーが LLM の使用を強制されない環境であっても、
他人の LLM によって生成された作業によるプレッシャーだけでも、非常に大きなものになります。 1缶
LLM 出力のレビューを強制されるか、単に同僚から圧力をかけられるだけか
さらに生産します。
モデルは次のことができるため、配信を加速する必要があります。
おそらく物事をより良く、より速く行うことができるでしょう。サプライチェーンあり
セキュリティが非常に大きなベクトルとなり、ワームが這い回るようになりました
NPM の開発者アカウントを中心に、配信頻度を向上させます
本当に悪い考えのようです。 4
LLM の誇大宣伝は、労働者に対するサイバー戦争の大きな波の一部です。
水に対して、地球に対して、そしてすべての人々に対して。これはそうではありません
個人的に「現実に適応する」か個人的な選択の問題、
しかし、これは政治的、社会的で難しい問題であり、私たちは集団的に取り組む必要があります。
以前のこのシリーズ: LLM の四騎士
黙示録。
私はここでも「AI」よりも LLM という用語を好みます。なぜなら、モデルはそうではないからです。
知性を持っています。タイトルに使ったので
クリックベイティングは重要なようですが、私は一歩手前でやめてしまいました。
これを「レイジ・アゲインスト・ザ・マシーン」と呼ぶのは、
私がこれまでに作成したすべてのブログ投稿のタイトル。 ↩
はい、Visual Studio が現在ではある程度無料であることは知っていますが、
それをレンタルにしたとしても驚かないでしょう、
なぜそうではないからです。 ↩
雇用市場の妨害を超えて、サム・アルトマン前夜

したくない
「ユーティリティとしてのインテリジェンス」を販売する
本当に悪いアイデアですが、特に彼らがいかに誇大妄想的であるかを示しています
人々はそうです。 ↩
別の時代の記憶が甦り、私たちを歩いてくれる
コンピューターセキュリティに関しては数十年前に遡ります。 ↩
LLM黙示録の四騎士
マストドン アカウントを使用してこの投稿に返信できます。
リンク
2026-05-16-四騎士
コピーレフト © 2002-2016
アナーキャット CC-BY-SA 。
マサチューセッツ工科大学-->
パワード
by ikiwiki 。
有効な (X)HTML 5 。

## Original Extract

anarcat

The people vs the AI overlords
Previously in this series: The Four Horsemen of the LLM
Apocalypse .
In a post to oss-security , my (Debian) co-developer Russ Allbery
stated that "open source software [OSS] is coming face to face with a
motivation crisis that has been building for a long time". His point
is essentially that large language models (LLMs 1 ) are making the
existing OSS community crisis worse. For him, it's the flood of code
reviews, but he argues that varies according to people's desires, for
others it's security issues and so on.
I think Russ is right, but I would argue there's something much bigger
than our open communities going on here, and it's about the entire
field of computing. This pressure is on all of us, regardless of
whether we work on open source software or not.
People using LLMs in their workflow have radically changed how
programming works, even for people who claim to avoid
vibe-coding . And I'm sorry to single out one poor maintainer here:
it's not you, Brian, you're just one example among many. But this is
typical use of those models nowadays:
Once it’s done, I’ll use /code-review and let Claude spawn
sub-agents to do a full review of the new code. This usually finds
some problems, even problems that the “main” Claude instance didn’t
find during its validation. I usually keep running /code-review
again and again after finding and fixing issues, until there aren’t
any left.
Think about what that means for a minute. This is automation built to
fire up dozens of agents crunching at a problem for minutes if not
hours of GPU compute time, in parallel. This is essentially a couple
of shelves in a datacenter rack, totally maxed out on power and
cooling, abstracted behind a cute little /code-review command.
The author, here, is rightly concerned that "Anthropic could pull the
rug out and require API pricing", which is perhaps a code word for
"charging something closer to actual costs". Brian also pays lip
service to environmental and societal costs but those are largely
abstracted away, so let's keep that conversation aside here as well,
as we have discussed it before anyways .
But clearly, this way of working has an ( externalized ) cost, to
say the least.
For decades my work has been focused on free and open source
software. I've long stopped using proprietary operating systems like
Windows or Mac, and even before that switch, I was mostly using free
software on those platforms, partly out of principle, but also because
I was too poor. So the tools of my trade are free, and I build free
tools with them.
It feels like we're going backwards: when I was in school, a millennia
ago, my classmates didn't have access to a compiler and were wondering
how they would scrape the money to buy a compiler like Borland's
or Microsoft's . I had a compiler built into my operating system
( FreeBSD at the time), so that wasn't a problem for me. For them,
it was a significant expense, but at least those expenses (or more
shady sourcing of programs ) were a one-shot deal.
Fast forward 30 years, and software is rented: you pay monthly for
Adobe's Photoshop and Microsoft's office suite just like you pay for
Netflix, Disney+ or Spotify 2 . And now you need to add dozens (if not
hundreds of dollars) of monthly credits to access LLMs on top of that.
So, now we have to pay to get anything done? This is peak
enshitification of our job: first they steal our work to train their
models, and then they sell it back to us at a profit.
AI is coming for our jobs, as engineers, if not everyone , according
to the narrative. For a while now, our job market has deteriorated:
less jobs, for less pay. Lots of skilled engineers looking for work
and finding crap jobs then still looking while working.
This is not by accident. 3 We engineers have a lot of power, it is not
organized, but that's just a couple of unions away ( easy !). Tech
overlords know this, so they are attacking our profession, directly,
by forcing us to train and use models that they can control.
Even in environments where programmers are not forced to use LLMs,
the mere pressure of other people's LLM-generated work is huge. One can
be forced to review LLM outputs, or just peer pressured you into
producing more.
We're now supposed to accelerate delivery, because models can
presumably do things so much better and faster. With supply chain
security becoming such a large vector that we now have worms crawling
around developers accounts on NPM , increasing the delivery cadence
seems like a really bad idea. 4
The LLM hype is part of the larger wave of cyberwar against workers,
against water, against the Earth, against all the people. This is not
a matter of individually "adapting to the reality" or personal choice,
but a political, social, hard problem we need to address collectively.
Previously in this series: The Four Horsemen of the LLM
Apocalypse .
I again prefer the term LLM to "AI" because models do not
possess intelligence . I did use it in the title because
click baiting is apparently important, but I stopped short of
calling this one "Rage Against the Machines" because that would be
the title of every blog post I have ever made. ↩
Yes, I know that Visual Studio is kind of free now, but I
wouldn't be surprised if they turn that into a rental as well,
because why not. ↩
Beyond sabotaging the job market, Sam Altman event wants to
sell " intelligence as a utility " something that is just a
really bad idea but especially shows how megalomaniac those
people are. ↩
This brings back memories of another era , walking us
back decades in terms of computer security. ↩
The Four Horsemen of the LLM Apocalypse
You can use your Mastodon account to reply to this post .
Links
2026-05-16-four-horsemen
Copyleft © 2002-2016 The
Anarcat CC-BY-SA .
MIT . -->
Powered
by ikiwiki .
Valid (X)HTML 5 .
