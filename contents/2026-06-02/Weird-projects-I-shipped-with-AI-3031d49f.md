---
source: "https://www.seangoedecke.com/weird-projects-i-shipped-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=48357486"
title: "Weird projects I shipped with AI"
article_title: "Weird projects I shipped with AI"
author: "Brajeshwar"
captured_at: "2026-06-02T00:38:35Z"
capture_tool: "hn-digest"
hn_id: 48357486
score: 3
comments: 0
posted_at: "2026-06-01T14:42:08Z"
tags:
  - hacker-news
  - translated
---

# Weird projects I shipped with AI

- HN: [48357486](https://news.ycombinator.com/item?id=48357486)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/weird-projects-i-shipped-with-ai/)
- Score: 3
- Comments: 0
- Posted: 2026-06-01T14:42:08Z

## Translation

タイトル: AI を使って出荷した奇妙なプロジェクト

記事本文:
AI を使って出荷した奇妙なプロジェクト Sean Goedecke
AI を使って出荷した奇妙なプロジェクト
AI によって生成されたプロジェクトはどこにあるのでしょうか?これは AI 懐疑論者からのよくある質問です。LLM がコードを書くのがそれほど上手なら、AI によって生成された新しいアプリ、サービス、ゲームの津波はどこにあるのでしょうか?
私個人としては、これがそれほど矛盾しているとは思いません。結局のところ、コードの作成は、実際に新製品を出荷する際のボトルネックの 1 つにすぎません。また、AI を使って私が行った有償の仕事について話すことも不可能です (AI によって私の生産性が向上したという言葉をただ受け入れる必要があります)。しかし、私にできることの 1 つは、過去 12 か月間に AI を使用して構築した個人プロジェクトのリストを共有することです。
これらすべてを手作業で行うことは絶対にありませんでした。そのうちの 1 つまたは 2 つを実行する時間を見つけたかもしれませんが、AI 以前の私の実績に基づくと、おそらく「数コミットの GitHub リポジトリ」段階に留まっていたでしょう。このリストは一種の存在証明です。少なくとも一部の人々にとって役立つ、AI の支援がなければ存在しなかったであろう奇妙なプロジェクトの束です。
最近では、古典的な Windows SkiFree ゲームのデイリー ゲーム バージョン (つまり、「Wordle に似ていますが、SkiFree 用」) である Skifreedle.com を構築しました。そのコードはここにあります 1 。
私は小さな Web ゲームを手作業でコーディングするのが好きですが、さまざまな SkiFree オブジェクトをすべて接続したり、最速のランの幽霊のようなきちんとした機能を構築したりする時間は絶対にありませんでした。また、気に入ったものにたどり着く前に、ゲーム UI のさまざまなビジュアル テーマをたくさん試しました。これを手作業で行っていたら、15 や 20 ではなく、2 つか 3 つの異なる外観を試すだけの時間があったでしょう。
この結果にとても満足しています。私は兄と競い合ってタイムを競うことを楽しんでいます。

オリジナルの SkiFree ゲームにとても懐かしさを感じます。
去年私はAutodeckを作りました！以前にこれについてブログ投稿を書きましたが、これは私のパートナーが、知りたいランダムなトピックに関する Anki カードを自動的に生成する方法があればいいのにという要望から来ました。最終的には、自動生成された間隔をあけた繰り返しカードの無限フィードをセットアップするのは比較的簡単になりました。
私がこの件で Stripe 支払いを設定したのは、お金を稼ぎたかったというよりも、誰かが Groq 残高を持ち逃げするのを心配したためでしたが、多くの人が実際にこれを使用しているのを見て嬉しい驚きを感じました。 500 人以上がそれを試しており、推論とホスティングをカバーするのに十分な有料購読者がいます。
LLM の支援がなければこれを構築できたかもしれませんが、Web サイトとして展開することはほぼなかっただろう。データベースと Stripe をセットアップするのは大変な作業です。
AI が生成したエンドレス Wiki も構築しました。これについてもブログ記事を書きました。 Autodeck と同様に、私は LLM の非チャット インターフェイスのアイデアに魅了され、リンクをクリックしてモデルを操作する Wiki ベースのアプローチが非常にクールだと思いました。
LLM 生成の呼び出しを通常のリンクの終端に置くのは悪い考えだということを、私は苦労して学びました。スクレーパーは推論予算をすぐに使い果たしてしまうからです。結局、記事が存在しないリンクを JavaScript で偽装することになり、少なくともこれまでのところ、スクレーパーを打ち破ることができました。人々は今でも Endless Wiki について私にメールを送ってきますが、28 万ページを超えるページが生成されています。
私の当初の目標は、ルート ページから開始してリンクのみをたどる (ウィキ ゴルフ のようなもの) 新世紀エヴァンゲリオンのページを最終的に生成できるかどうかを確認することでした。成功しました！ 「エヴァンゲリオンアニメ」ページはこちらからご覧いただけます。
ほぼ正確に1か月

私が Endless Wiki を立ち上げた後、xAI は Grokipedia を立ち上げました。明らかに彼らは私を盗用したわけではありません。これは非常に簡単なアイデアであり、私のサイトは最初の無限 Wiki ではありませんでした (ただし、リンクをクリックして新しいページを見つけなければならない最初のサイトだったと思います)。しかし、若干の輝きは消えてしまいました。
私は、VicFlora 植物識別データベースをキャッシュする PWA を構築して、インターネット環境が低い場合やインターネットがない場合でも使用できるようにしました。これは、植物が好きで、インターネットが不安定な場所に時々野外旅行に行く私のパートナーのためのユーティリティプロジェクトでした。
LLM がなければ間違いなくこれを行うことはできなかったでしょう。 VicFlora Web サイトから基本的な二分キーをスクレイピングするのはかなり困難でした。API ドキュメントは古く、データを取得するための経路は複数考えられ (そのほとんどは機能しませんでした)、取得できたデータの形式は解析が困難でした。十分な努力があればできたと思いますが、かなりの作業量になったでしょう。
この結果にとても満足しています。完璧ではありませんが、機能しており、時々ビクトリア州の植物学者からバグレポートや機能のリクエストをメールで送ってもらったこともあり、明らかに少しずつ使用されています。
私は、必ずしも「デプロイされたプロジェクト」のレベルに達する必要はない他の多くのことを実行しました。 スターが 100 個を少し超えるスタンドアップ レポートを自動的に生成する gh-standup GitHub CLI 拡張機能、ここでブログに書いた (低品質) 画像地理位置情報ベンチマーク、またはオープンソース モデルから機能を抽出するスキルです。
AI によって生み出された企業が（まだ）氾濫しているわけではないかもしれませんが、少なくとも私にとっては、LLM の多大な支援がなければ存在しなかったであろう小さくて奇妙なプロジェクトが氾濫しています。
私もショーしたい

これの Simon Willison バージョンを公開します。これも、「作成コストが非常に低かったためにのみ存在する、奇妙な便利なツール」の素晴らしい例です。
私は、DanielHough の SkiFree.js からスプライトシートを引用しました。これは、それが Wing Wang Wao によるものであると考えられます。もちろん、オリジナルのスプライトとアートは Chris Pirih の SkiFree と Microsoft に属します。
この投稿を気に入っていただけた場合は、私の新しい投稿に関する更新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
GPT-5-Codex は私より優れた AI 研究者です
ラップトップで 5 分でトレーニングできる最強の AI モデルは何ですか?私は AI 研究に関する愚かな質問に答えてみました。おそらくそれが何であったかは推測できるでしょう。
Python スクリプトを使い始めたりアイデアを出したりするために GPT-5 とチャットしましたが、調査を行ったのはやはり私でした。私はアイデアを考え出し、実験を実行し、データに基づいて次に何をすべきかを決定していました。私がトレーニングできた最良のモデルは、次のような出力を生成する 1.8M パラメータ トランスフォーマーでした。
続きを読む...
購読する │ About │ ポッドキャスト │ 人気 │ タグ │ RSS

## Original Extract

Weird projects I shipped with AI sean goedecke
Weird projects I shipped with AI
Where are all the AI-generated projects? This is a common question from AI skeptics: if LLMs are so good at writing code, where is the tsunami of new AI-generated apps, services and games?
I personally don’t find this to be much of a paradox. Writing code is only one of the bottlenecks involved in actually shipping a new product, after all. It’s also impossible to talk about the paid work I’ve done with AI (you’ll simply have to take my word that it’s increased my productivity). But one thing I can do is share a list of personal projects I’ve built with AI in the last twelve months.
I definitely would not have done all of these by hand. I might have found the time to do one or two of them, but based on my pre-AI track record they would probably have stayed in the “GitHub repo with a few commits” stage. This list is a kind of existence proof : a bunch of weird projects, useful to at least some people, that would not have existed without AI assistance 0 .
Most recently I’ve built skifreedle.com , a daily-game version of the classic Windows SkiFree game (i.e. “like Wordle, but for SkiFree”). The code for that is here 1 .
I enjoy coding small web games by hand, but definitely would not have had the time to wire up all the different SkiFree objects or build neat features like a ghost of your fastest run. I also tried out a lot of different visual themes for the game UI before landing on something I liked. If I’d done this by hand, I would have only had time to try out two or three different looks, instead of fifteen or twenty.
I’m very happy with how this turned out. I’ve been enjoying competing against my brother to get better times, since both of us have a lot of nostalgia for the original SkiFree game.
Last year I built Autodeck ! I wrote a blog post about this before, but this came from my partner wishing there was some way to automatically generate Anki cards about random topics she wanted to learn about. It ended up being relatively straightforward to set up an endless feed of auto-generated spaced repetition cards:
I set up Stripe payments for this one, more because I was worried about someone running away with my Groq balance than because I wanted to make money, but I was pleasantly surprised to see a bunch of people actually use this. Over five hundred people have tried it out, with enough paid subscribers to cover inference and hosting.
I might have built this without LLM assistance, but I almost certainly would not have deployed it as a website. The hassle of setting up a database and Stripe would have just been too much work.
I also built an AI-generated endless wiki . I wrote a blog post about this one as well. Like Autodeck, I was fascinated with the idea of non-chat interfaces for LLMs, and I thought a wiki-based approach where you interact with the model by clicking links was pretty cool.
I learned the hard way that putting a LLM generation call on the end of a regular link was a bad idea: scrapers would exhaust my inference budget quickly. I ended up faking the no-article-exists-yet links with JavaScript, which at least so far has defeated scrapers. People still email me about Endless Wiki, and there are over 280 thousand pages generated.
My original goal was to see if you could eventually generate a page for Neon Genesis Evangelion, starting at the root page and only following links (kind of like wiki golf ). I was successful! You can read the “Evangelion Anime” page here .
Almost exactly a month after I launched Endless Wiki, xAI launched Grokipedia . Obviously they didn’t plagiarize me. This is a very easy idea to have, and my site was not the first infinite wiki (though I think it was the first one where you had to discover new pages by clicking on links). But it did take some of the shine off.
I built a PWA that caches the VicFlora plant identification database so it could be used with low or no internet. This was more of a utility project for my partner, who likes plants and occasionally goes on field trips where internet is spotty.
I would definitely not have done this without LLMs. It was reasonably difficult to scrape the basic dichotomous key from the VicFlora website: their API documentation was out of date, there were multiple possible pathways for fetching data (most of which were not functional), and the format of the data I did manage to fetch was hard to parse. I think I could have done it, with enough effort, but it would have been a substantial amount of work.
I’m very happy with how this turned out. It’s not perfect, but it’s functional, and I’ve even had the occasional Victorian botanist email me with bug reports or feature requests, so it’s clearly seeing a little bit of usage.
I did a bunch of other stuff that doesn’t necessarily rise to the level of a “deployed project”: my gh-standup GitHub CLI extension to automatically generate a standup report, which has just over a hundred stars, my (low quality) image geolocation benchmark , which I blogged about here , or my skill for extracting features from open-source models.
There may not be a flood of AI-generated companies (yet), but at least for me there’s been a flood of small, weird projects that would not have existed without significant LLM assistance.
I also want to shout out Simon Willison’s version of this , which is another great example of “weird useful tools that only exist because the cost of creating them was so low”.
I did lift the spritesheet from DanielHough’s SkiFree.js , which attributes it to Wing Wang Wao . Of course, the original sprites and art belong to Chris Pirih’s SkiFree and Microsoft.
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
GPT-5-Codex is a better AI researcher than me
In What’s the strongest AI model you can train on a laptop in five minutes? I tried my hand at answering a silly AI-research question. You can probably guess what it was.
I chatted with GPT-5 to help me get started with the Python scripts and to bounce ideas off, but it was still me doing the research. I was coming up with the ideas, running the experiments, and deciding what to do next based on the data. The best model I could train was a 1.8M param transformer which produced output like this:
Continue reading...
subscribe │ about │ podcasts │ popular │ tags │ rss
