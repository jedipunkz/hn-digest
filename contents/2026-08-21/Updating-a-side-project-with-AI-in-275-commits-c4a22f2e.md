---
source: "https://benhoyt.com/writings/updating-gifty-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=49384790"
title: "Updating a side project with AI in 275 commits"
article_title: "Updating a side project with AI in 275 commits"
image: ""
author: "ingve"
captured_at: "2026-08-21T07:34:51Z"
capture_tool: "hn-digest"
hn_id: 49384790
score: 1
comments: 1
posted_at: "2026-08-21T07:09:24Z"
tags:
  - hacker-news
  - translated
---

# Updating a side project with AI in 275 commits

- HN: [49384790](https://news.ycombinator.com/item?id=49384790)
- Source: [benhoyt.com](https://benhoyt.com/writings/updating-gifty-with-ai/)
- Score: 1
- Comments: 1
- Posted: 2026-08-21T07:09:24Z

## Translation

タイトル: 275 コミットで AI を使用したサイド プロジェクトを更新
説明: AI エージェントを使用して GiftyWeddings.com Web サイトを大幅に更新し、275 回のコミットでギフト レジストリから完全な Web サイト ビルダーに移動しました。

記事本文:
275 コミットで AI を使用してサイド プロジェクトを更新する
最近、私の雇用主は年半ばの休暇を取得し、全員がほぼ同時に 2 週間の休暇を取得しました。マネージャーとして、私は AI ツールを、適切なコーディングではなく、主にコード レビューと技術的調査に使用していました。そして、それを変えたいと思いました。
そこで、私は余暇をサイドプロジェクトである Gifty Weddings のハッキングに費やしました。以前、2016 年に Go バックエンドに、2019 年に Elm フロントエンドに移行しました。しかし今、私には 2 つの目標がありました。
単なるウェディング ギフト レジストリからウェディング Web サイト ビルダーにアップグレードします。
AI ツールを使用して高品質のコードを作成する方法を学びます。
同僚の勧めで、ほとんどのコーディングに Opus 5 の Claude Code を使用しました。小規模なタスクの場合は、Pi とオープンウェイト GLM 5.2 モデルを使用しました。頭も使いました。
冗談ですが、自分の頭を使うことが、余ったものを大量に生み出すことと、誇りに思える製品を作ることを分ける大きな部分を占めると思います。私たちはまだエンジニアです。
この記事では、AI に対する私の懐疑についてだけでなく、このプロジェクトの成功に何が貢献したと私が考えるか、そしてなぜ私がこのプロジェクトを楽しんだのかについても話します。
私はしばらくの間、AI について非常に懐疑的でした。最初は AI チャットボットがニューヨーク タイムズの記者と恋に落ちた初期の頃でした。その後、「寄稿者」によって提供される大量の AI のスロップに対処しなければならなくなってからはさらに懐疑的になりました。
私は 2024 年にコーディングのためにそれを試しましたが、感銘を受けませんでした。その時点で、私は自分で修正するよりも、修正に多くの時間を費やしました。
最近では、妻のインテリア デザイン Web サイトを基本的に 1 回で構築し、GoAWK のさまざまなバグ (Opus 4.6 がかなり優柔不断だったバグも含む) を修正することに再挑戦しました。
どちらの場合も、私は感銘を受けました。Web サイトの場合は、CSS を書くのが好きではないからです。

バグ修正の場合は、プロセスが本当にスピードアップし、良いアイデアが得られたからです。
私はかなり経験豊富な Web 開発者であるという事実から始めます。これは、プロセスをガイドし、エージェントの出力を効果的にレビューできることを意味します。余談ですが、AI に関する私の最大の懸念の 1 つは、新しい開発者がこの苦労して得た経験をショートカットしようとする誘惑に駆られることです。
また、既存の作品の基礎の上に構築していました。私は古いサイトの CSS スタイルシート (このサイトは Skeleton に基づいています)、ほとんどの SQL データベース スキーマ、および構築したいものの計画から始めました。さらに、古い Gifty コードベースに基づいて、初期の Go サーバーといくつかのパッケージを手作業で作成しました。必要な種類のコードをシードしました。
そこから、私は一発でやろうとするのではなく、反復的な開発に切り替えました。私はすべての段階で技術的および創造的なコントロールを持っていました。コードを変更するたびに、通常は数文のプロンプトとともに機能を提案し、コード出力を確認してブラウザで機能をテストします。
中程度のコードレビューを行いました。コードのスタイルが完璧ではなかったか、少なくとも多くの場合、私がそうしていたであろうやり方ではありませんでした。もちろんバグを指摘したり、構造的な問題に挑戦したりしましたが、ほとんどの場合、コードが書いたものにはかなり満足していました。
ただし、テストを徹底的に見直したわけではありません。 AI エージェントは大量のテストを作成しているようですが、場合によっては多すぎるため、費用がかからないと思われるテストをいくつか削除しました。最初は詳細をいくつか確認していましたが、最後にはテストをかなり軽く流し読みしていました。
その過程で私は多くのことを学びました:
LLM は非常に冗長なコメントを書きます。簡潔にするか、1 行の要約を書くように常に指示しなければなりませんでした。で

「tersify comments」コミットでは、約 2500 行のコメント行を 1500 行に削減しました。
クロードにヘッドレス ブラウザをインストールして独自のスクリーンショットを撮るように依頼したところ、HTML と CSS を記述する能力が大幅に向上しました。その前に、スクリーンショットを撮って手動でアップロードする必要がありました。サーバーを起動し、アプリの実際のフォームを使用して偽のデータを追加し、Chromium を使用して PNG スクリーンショットを保存し、それらを「調べ」ました。
とはいえ、変更のたびにスクリーンショットを撮って処理するには、大量のトークンが使用されました。結局、フロントエンドに重要な変更を加えた場合にのみそれを行うように依頼しました。
コンテナ内でサンドボックス化されている場合でも (私は Canonical の Workshop を使用しています)、Claude は迷惑なことを行います。 rm -rf $SOMEVAR/*.png のようなプログラムが何度か実行されましたが、SOMEVAR が設定されておらず、使用していたテスト写真が削除されました。メモリにいくつかのルールを追加して停止させようとしましたが、LLM は常にコンテナまたは VM で実行されます。
1 回のセッションですべての作業を行わないでください。このように始めましたが、文脈が非常に長くなり、すぐにクロードプロのさまざまな制限を超えてしまいました。そこで私は圧縮について学び、定期的に新しいセッションを開始しました。
それでは、私が構築した機能を見てみましょう。
Gifty を使用すると、カップルは写真、テキスト、ギフト登録を含むシンプルな複数ページの結婚式 Web サイトを作成できます。 (旧 Gifty はレジストリ部分のみをサポートしていました。)
カップルは、Markdown セクションにテキストを書いたり、写真をアップロードしたり (サイズが縮小されて Tigris に保存されます)、ページを追加したり並べ替えたりすることができます。そしてもちろん、私は少し何かを請求します。支払いはStripeを使用して処理されます。
ゲストはカップルのウェブサイトを閲覧し、カップルの登録簿にあるギフトに×印を付けることができます。
しかし、私のお気に入りの機能は、古い Gifty からコピーされたもので、カップルがワンクリックで試すことができます。
私は物事をシンプルにすることが好きなので、次の技術を使用しました。
ゴー

アップロード用の AWS SDK、Stripe SDK、画像サイズ変更ライブラリ、Markdown レンダラー、modernc.org/sqlite モジュール、および「ほぼ stdlib」の golang.org/x/crypto モジュールなど、stdlib 以外の依存関係をできる限り少なくしたバックエンド。
素晴らしい SQLite データベース。
より動的な部分の HTML : ギフト レジストリとページ編集のサポート。
そして、それが意味のあるバニラ JavaScript の数行です。
これは Fly.io でホストされています。この種のことには私が強くお勧めするサービスです。flydeploy を使用すると非常に簡単で、安価です。
新しいウェブサイトの完成には丸 10 日ほどかかりました。 AI ツール、特に今回の場合はクロードにとても感謝しています。私は妻に、AI がなかったらおそらく 3 倍くらい時間がかかっていただろうと言いました。
私が最も感銘を受けたことの 1 つは、Claude が古いギフト レジストリ システム (Elm フロントエンドを備えた Go JSON API) を、まったく異なる htmx フロントエンドを備えた新しいバージョン (Go HTML エンドポイント) に移植したことでした。これは一発でほぼ完璧に完了しました。これは私にとって嬉しい驚きであり、時間を大幅に節約できました。
もう 1 つのことは、基本的に 1 回の作業 (いくつかのフォローアップのバグ修正を伴う) で、古い Gifty データベースを新しいデータベースに移行するための移行ツールを作成することでした。データベースは似ているので、ロケット科学ではありませんが、構造的な違いがいくつかあり、さらに反復が必要になると予想していました。
しかし、私が AI に全体的に懐疑的であることを考えると、なぜ私はこれをそれほど楽しめたのでしょうか?理由は 2 つあると思います。
まず、私は物を作るのが好きです。便利で見栄えの良いウェブサイトが約 10 日で完成しました。
第二に、私はプログラミングという技術を楽しんでおり、それはまだ得られていると感じていました。次の機能を考え、AI にコーディングを依頼し、コードをレビューし、バグを修正してコミットします。洗い流しを275回繰り返します。小さな機能には 10 ～ 15 分かかりますが、大きな機能には 10 ～ 15 分かかります

時間は 1 ～ 2 時間ですが、目標に向かって着実に進歩しているのはとても気持ちがよかったです。
私にはまだ懸念事項がたくさんありますが、それを利用して天国に届く塔を建て、神が私たちを釘を一、二本降ろすことのないように願っています。でも、新しいツールには感謝しています。
最後にもう 1 つ: もうすぐ結婚する方、または結婚する人を知っている方は、新しい GiftyWeddings.com を紹介していただければ幸いです。結婚したいけどまだ相手がいないという方は、申し訳ありませんが、別のサイトをご利用ください。

## Original Extract

I used an AI agent to make a big update to my GiftyWeddings.com website, taking it from a gift registry to a full website builder in 275 commits.

Updating a side project with AI in 275 commits
Recently my employer had its mid-year break, where everyone gets two weeks off at roughly the same time. As a manager, I’d been using AI tools mostly for code review and technical exploration, rather than coding proper, and I wanted to change that.
So I spent my time off hacking on my side project, Gifty Weddings . I’d previously moved it to a Go backend in 2016 and an Elm frontend in 2019 . But now I had two goals:
Upgrade it from just a wedding gift registry to a wedding website builder.
Learn how to write high-quality code with AI tools.
At my coworkers’ recommendation, I used Claude Code with Opus 5 for most of the coding. For smaller tasks, I used Pi with the open-weight GLM 5.2 model. I also used my brain.
I jest, but I do think using one’s brain is a big part of what separates churning out slop from building a product you can be proud of. We’re still engineers.
In this article I’ll talk about my skepticism of AI, but also what I think contributed to the success of this project, and why I had fun doing it.
I’ve been quite skeptical about AI for a while, first in the early days when an AI chatbot fell in love with a New York Times journalist, then even more when I started having to deal with tons of AI slop being served up by “contributors”.
I had tried it for coding back in 2024 and was not impressed. At that point, I spent more time fixing up what it did than I would have spent doing it myself.
More recently, I gave it another try to build my wife’s interior design website in basically one shot, and then to fix various bugs in GoAWK (including one where Opus 4.6 was rather indecisive ).
Both of those times, I was impressed: in the website case, because I don’t love writing CSS, and in the bug-fixing case, because it really sped up the process and gave me good ideas.
I’m starting with the fact that I’m a fairly experienced web developer . This means I can guide the process and review the agent’s output effectively. As a side note, one of my biggest concerns with AI is that new developers will be tempted to shortcut this hard-won experience.
I was also building on a foundation of existing work. I started with the CSS stylesheet from my old site (which in turn is based on Skeleton ), most of an SQL database schema, and a plan of what I wanted to build. In addition, I wrote the initial Go server and several packages by hand, based heavily on the old Gifty codebase. I seeded it with the kind of code I wanted.
From there I switched to iterative development , rather than trying to do it in one shot: I had technical and creative control at every stage. For each code change, I would suggest the feature – usually with a prompt of a few sentences – and then review the code output as well as test the feature in my browser.
I did a moderate level of code review . I had to give up some of my code craftsmanship – its code style wasn’t perfect, or at least not how I’d have done it in many cases. I’d point out bugs, of course, and challenge structural issues, but for the most part I was pretty happy with the code it wrote.
However, I didn’t thoroughly review the tests . AI agents seem to write a lot of tests – sometimes way too many, and I deleted several I didn’t think paid for themselves. At first I’d review some of the details, but by the end I was skimming tests pretty lightly.
I learned a number of things along the way:
LLMs write really verbose comments. I had to continually tell it to be succinct or write one-line summaries. In one “tersify comments” commit, I reduced about 2500 comment lines to 1500.
Claude’s ability to write HTML and CSS got a lot better when I asked it to install a headless browser and take its own screenshots. Before that I’d have to take screenshots and upload them manually. Now it had superpowers: it started the server, added fake data using the app’s real forms, then used Chromium to save PNG screenshots and “looked” at them.
That said, taking and processing screenshots on every change used a lot of tokens. I ended up asking it to do that only when making significant frontend changes.
Even sandboxed in a container (I use Canonical’s Workshop ), Claude does annoying stuff. Several times it ran a program like rm -rf $SOMEVAR/*.png , but SOMEVAR wasn’t set and it deleted the test photos I was using. I added some rules to its memory to try to get it to stop – but yes, always run LLMs in a container or VM.
Don’t do all your work in one session. I started this way, but the context got really long, and I soon exceeded Claude Pro’s various limits. So I learned about compaction and started new sessions regularly.
Now let’s look at the features I built.
Gifty allows a couple to create a simple, multi-page wedding website with photos, text, and a gift registry. (Old Gifty only supported the registry part.)
Couples can write text in Markdown sections, upload photos (they’re downsized and stored on Tigris ), add and reorder pages, and so on. And of course I charge a little something; payments are handled using Stripe.
Guests can view a couple’s website and cross off gifts on the couple’s registry.
But my favourite feature, copied from old Gifty: couples can try it out with a single click .
I love keeping things simple, so I used the following tech:
A Go backend with as few non-stdlib dependencies as possible: the AWS SDK for uploads, the Stripe SDK, an image-resizing library, a Markdown renderer, the modernc.org/sqlite module, and the “almost stdlib” golang.org/x/crypto module.
The wonderful SQLite database.
Htmx for the more dynamic parts: the gift registry and the page editing support.
And a few lines of vanilla JavaScript where it made sense.
It’s hosted on Fly.io , a service I highly recommend for this kind of thing: fly deploy makes it so easy, and it’s inexpensive.
Finishing the new website took about 10 full days. I’m very thankful for AI tools, particularly Claude in this case. I said to my wife it probably would have taken me about three times as long without AI.
One of the things that impressed me most was Claude porting the old gift registry system, a Go JSON API with an Elm frontend, to the new version – Go HTML endpoints with a quite different htmx frontend. It did this almost perfectly in one shot, which was a pleasant surprise for me and a big time-saver.
One other thing it did in basically one shot (with a few follow-up bug fixes) was produce a migration tool, to migrate the old Gifty database to the new. It’s not rocket science, as the databases are similar, but there are a few structural differences, and I expected this to take more iterations.
But why, given my overall AI skepticism, did I enjoy it so much? I think there are two reasons:
First, I like building things. I had a useful and nice-looking website finished in about 10 days.
Second, I enjoy the craft of programming, and I still felt I was getting that. I would think of the next feature, ask AI to code it, review the code, fix bugs, and commit it. Rinse and repeat 275 times. Smaller features would take 10-15 minutes, larger ones an hour or two, but it felt great to be making constant progress towards the goal.
I still have plenty of concerns, and I hope we don’t use them to build a tower to reach heaven so God has to take us down a peg or two. But I’m grateful for the new tools.
One last thing: if you’re getting married soon or know someone who is, I’d love it if you pointed them to the new GiftyWeddings.com . If you want to get married but don’t have anyone yet – sorry, but that’s another website!
