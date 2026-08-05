---
source: "https://dylan.blog/2026/08/03/im-scared-a-stranger-will.html"
hn_url: "https://news.ycombinator.com/item?id=49185556"
title: "I'm Scared a Stranger Will Call My Novel AI, So I Built GitHub for Words"
article_title: "I'm Scared a Stranger Will Call My Novel AI, So I Built GitHub for Words (Meet VellumProof, Formerly WritHub, Lol) - dylan's blog"
author: "harper"
captured_at: "2026-08-05T17:20:49Z"
capture_tool: "hn-digest"
hn_id: 49185556
score: 2
comments: 1
posted_at: "2026-08-05T16:58:04Z"
tags:
  - hacker-news
  - translated
---

# I'm Scared a Stranger Will Call My Novel AI, So I Built GitHub for Words

- HN: [49185556](https://news.ycombinator.com/item?id=49185556)
- Source: [dylan.blog](https://dylan.blog/2026/08/03/im-scared-a-stranger-will.html)
- Score: 2
- Comments: 1
- Posted: 2026-08-05T16:58:04Z

## Translation

タイトル: 見知らぬ人に私の小説を AI と呼ばれるのが怖いので、Word 用の GitHub を構築しました
記事のタイトル: 見知らぬ人が私の小説を AI と呼ぶのが怖いので、Words 用の GitHub を構築しました (VellumProof の紹介、以前は WritHub でした、笑) - dylan のブログ
説明: 以前にも述べたように、私たちは現在、クリエイティブな人々が自分が作成したものを自分で作成したことを証明しなければならない世界に住んでいます。
ちょっと待ってください。なぜなら…

記事本文:
メインコンテンツにスキップ
コントロール
ディランのブログ
見知らぬ人が私の小説を AI と呼ぶのが怖いので、Word 用の GitHub を構築しました (VellumProof、旧 WritHub を紹介します、笑)
以前にも述べたように、私たちは現在、創造的な人々が自分が創造したものを創造したことを証明しなければならない世界に住んでいます。
それは本当に気が狂っているので、ちょっと待ってください。私は 2014 年から小説を書いていますが、出版されることは難しい部分であり、私は難しい部分が苦手なので未発表です。私を夜も眠れなくさせているのは、拒絶反応ではありません。それは、見知らぬ誰かが私のまったく普通の文章を読んで、それが AI であると感じ、その夢全体を雰囲気で台無しにするということです。
私がガラスを噛みたくなるのは、人々が物を作るために何を使用したかについてただ正直であれば、これらは何の問題もないだろうということです。しかし、それは明らかに言い過ぎです。なぜなら、詐欺師は必ず騙すからです。
その代わりに、残りの人は領収書を持ち歩かなければなりません。
私が最初にやったことは、そうしなければならなかったので、LLM に小説の書き方を教えられるかどうかを確認することでした。敵を知るなど。すべての AI にはチックがあるため、これはそれ自体でプロジェクト全体でした。そして、実際の文章にチックが存在しないわけではありません（私はあなたを見ています、エムダッシュ、私の愛する人）。それは、機械が常に彼らに手を伸ばしているということです。
それがまさにジレンマです。人間の作家の中にも em-dash に頼る人もいます。では、シソーラスを備えたロボットではなく、ページに単語を載せたことをどうやって証明するのでしょうか?言葉だけでは伝わりません。言葉は犯罪現場であり、アリバイではない。
オタクの領域へのちょっとした寄り道
あなたがオタクではない場合のために、GitHub について説明しましょう (私はまだオタクなので、穏やかにお話します)。
GitHub はプログラマーがプロジェクトを保管する場所です。バージョンを保存し、何もせずに新しい機能に取り組むことができます。

ライブなものを壊すわけではありません。そして、ここが私が気に入っている部分です。毎日コミットしたコードの量を示す小さなヒート マップが描画されます。この地図により、人間と機械を見分けるのが愚かなほど簡単になります。午後に 50,000 行をコミットする人は、それらを手で入力したわけではありません。誰かが数百個やってる？おそらく本物でしょう。単純。
もう 1 つの魔法のトリックは、すべての変更を保持することです。全員です。これは、作家にとって最も古い悪夢、つまり作品を失うということが基本的に不可能になることを意味する。すべてのバージョンを保存しておけば、失うものは何もありません。
私は執筆も含め、すべてのプロジェクトを GitHub に保存してきました。しかし、GitHub には私のような人にとって 1 つ問題があります。それは行数をカウントすることです。そして作家は線に沿って考えません。私たちは言葉で考えます。私たちは言葉が大好きです。一言いただければ、1 時間お話しします。
そこで私は、GitHub の美しい変更追跡機能を取り入れて、言葉を重視する人々のために再構築する必要がありました。
VellumProof の紹介 (危うく WritHub という名前になりました、笑)
その仕組みは次のとおりです。GitHub よりも怖くないと約束します。執筆活動の 1 日の終わりに、原稿をアップロードします。 VellumProof はそのバージョンを隠しておき、次回アップロードするときに 2 つを比較し、何が変更されたのか (単語の追加、単語の削除) を正確に記録し、日付をスタンプします。 git commit から、本物の少年のように感じるためにコマンドラインを学ばなければならない部分を除いたものです。
プロジェクト内のどのファイルを追跡するかを選択すると、それらのファイル間で何が変更されたかを確認できます。本を編集者に送りますか?パスをアップロードすると、何をどれだけタッチしたかを正確に確認できます。それだけでも価値があります。
しかし、私が一番気に入っているのは、それが証拠であるということです。
アップロードごとに追加された単語の数を追跡するグラフがあります。そしてそのグラフは、あなた自身が世界に貢献していることのちょっとした証拠です。

もの。私がこれまでに 1 日に書いた最高の文字数は約 6,000 ワードです。素晴らしい一日でした。それは私の能力の絶対的な限界でもあります。したがって、私の履歴に火曜日に突然 10,000 語が表示された場合、私がその 10,000 語を書いていない可能性が非常に高く、誰が見ても私と同じことがわかるでしょう。
ここで、問題があります。古いプロジェクトを最初にアップロードすると、非常識に見えます。初日に Calliope と Rogue Biologic の 50 数千語を書き込むと、グラフは午後で小説を 1 冊書いたと叫ぶでしょう。それがメモの目的です。私はそのコミットに「NaNoWriMo 2014 中に草稿された CRB の最初のアップロード (これは実質的には 5 番目の草案)」とタグ付けしました。これで、歴史を読んだ人なら誰でも理解できるようになります。初日は私が本をインポートすることであり、その後はすべて、実際の、研削的な、一語一語の編集の軌跡です。
その夢、つまり実際の夢は、あなたが何かを投稿し、誰かが目を細めて「これは AI のようだ」と言ったら、あなたがそれを書いている自分のタイムラプス写真を渡すことができるというものです。雰囲気ではありません。記録です。領収書。
はい。詐欺師は騙すつもりだ。私は朝食前にこれをプレイする 5 つの方法を個人的に考えましたが、それらすべてをブロックするバージョンはありません。 AI テキストを少しずつ貼り付けて、人間の曲線を偽装することができます。たくさんのことができます。
しかし、ここに私が着地したのです。最悪の人々に対する要塞の構築に永遠に費やすこともできますし、正直な多数派のための優れたツールを構築して、ほとんどの人がこの物語の悪者ではないと信じることもできます。私は2番目を選択しています。私はすべての作家を容疑者のように扱うよりも、物事の人間的な側面に賭けて、時には間違ったことをしたいと思っています。お気づきかと思いますが、それが私がそもそも怒っていることそのものです。
ともかく。私はmのために嘘発見器を作りました

私は怖いから自分の小説を書いています。物を作るのは私が怖い気持ちを代謝する方法です。それが私がこれまでに行った中で最も作家らしい仕事でないとしたら、何がそれなのかわかりません。
領収書を保管し、自分のツールについて正直になり、誰も求めていないものを書きに行きましょう。
ディランのブログ © 2026 · カオスゴブリンによる提供
エネルギー

## Original Extract

As I’ve mentioned before, we now live in a world where creative people have to prove they created the thing they created.
Sit with that for a second, because …

Skip to main content
Controls
dylan's blog
I'm Scared a Stranger Will Call My Novel AI, So I Built GitHub for Words (Meet VellumProof, Formerly WritHub, Lol)
As I’ve mentioned before , we now live in a world where creative people have to prove they created the thing they created.
Sit with that for a second, because it’s genuinely deranged. I have been writing fiction since 2014 — unpublished, because getting published is the hard part and I am bad at the hard part — and the thing that keeps me up at night isn’t the rejection. It’s that some stranger is going to read my perfectly normal writing tics, feel that they’re AI, and torpedo the whole dream over a vibe.
The part that makes me want to chew glass is that none of this would be a problem if people were just honest about what they used to make a thing. But that’s apparently too much to ask, because cheaters gonna cheat.
So instead, the rest of us have to carry receipts.
The first thing I did — because I had to — was see if I could teach an LLM to write fiction. Know your enemy, etc. This was a whole project on its own, because every AI has tics. And it’s not that the tics don’t exist in real writing (I’m looking at you, em-dash, my beloved). It’s that the machine reaches for them constantly .
Which is exactly the dilemma. Some human writers lean on the em-dash too. So how do you prove that you put the words on the page and not a robot with a thesaurus? You can’t do it from the words alone. The words are the crime scene, not the alibi.
A brief detour into nerd territory
Let me tell you about GitHub , in case you’re not a nerd (I’m a baby nerd, so I’ll keep it gentle).
GitHub is where coders keep their projects. It saves versions, it lets you work on new features without breaking the live thing, and — here’s the part I love — it draws you a little heat map showing how much code you committed on each day. That map makes it stupidly easy to tell a human from a machine. Someone who commits 50,000 lines in an afternoon did not type those by hand. Someone doing a couple hundred? Probably real. Simple.
The other magic trick is that it keeps all your changes. All of them. Which means writers' oldest nightmare — losing the work — basically stops being possible. If every version is saved, there’s nothing to lose.
I’ve been keeping all my projects on GitHub, writing included. But GitHub has one problem for people like me: it counts lines . And writers don’t think in lines. We think in words. We love words. Give us a word and we’ll turn it over for an hour.
So I just needed to take the beautiful change-tracking guts of GitHub and rebuild them for people who count in words.
Introducing VellumProof (which was almost named WritHub, lol)
Here’s how it works, and I promise it’s less scary than GitHub: at the end of a writing day, you upload your manuscript. VellumProof tucks that version away, and the next time you upload, it compares the two and records exactly what changed — words added, words cut — and stamps it with the date. It’s git commits, minus the part where you have to learn the command line to feel like a real boy.
You pick which files in a project get tracked, and you can see what’s changed between any two of them. Send your book to an editor? Upload their pass and see precisely what they touched and how much. That alone is worth it.
But my favorite part is that it’s proof .
There’s a graph that tracks how many words got added per upload. And that graph is a little bit of evidence of your own hand in the thing. The most I have ever written in a single day is around 6,000 words. It was a great day. It is also the absolute ceiling of my ability. So if my history suddenly shows a 10,000-word Tuesday, there’s a very good chance I did not write those 10,000 words — and anyone looking can see the same thing I can.
Now, a wrinkle: your first upload of an old project is going to look insane. If I dump in Calliope and the Rogue Biologic — 50-some-thousand words — on day one, the graph will scream that I wrote a whole novel in an afternoon. That’s what the notes are for. I tag that commit “Initial upload of CRB, drafted during NaNoWriMo 2014 (this is realistically the fifth draft),” and now anyone reading the history understands: day one was me importing a book, and everything after is the real, grinding, word-by-word edit trail.
The dream — the actual dream — is that when you submit something and somebody squints and says “this feels like AI,” you can hand them a timelapse of yourself writing it. Not a vibe. A record. Receipts.
Yes. Cheaters are going to cheat. I have personally thought of five ways to game this thing before breakfast, and there is no version of it that blocks all of them. You could paste in AI text a little at a time to fake a human curve. You could do a lot of things.
But here’s where I’ve landed: I could spend forever building a fortress against the worst people, or I could build a good tool for the honest majority and trust that most people aren’t the villain of this story. I’m choosing the second one. I’d rather bet on the human side of things and be occasionally wrong than treat every writer like a suspect — which is, you’ll notice, the exact thing I’m mad about in the first place.
Anyway. I built a lie detector for my own novel because I’m scared, and building things is how I metabolize being scared. If that’s not the most writer thing I’ve ever done, I don’t know what is.
Keep your receipts, be honest about your tools, and go write something nobody asked for.
dylan's blog © 2026 · Powered by chaos goblin
energy
