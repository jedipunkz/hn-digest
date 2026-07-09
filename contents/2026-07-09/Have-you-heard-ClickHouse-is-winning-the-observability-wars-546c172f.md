---
source: "https://charity.wtf/p/have-you-heard-clickhouse-is-winning"
hn_url: "https://news.ycombinator.com/item?id=48839068"
title: "Have you heard? ClickHouse is winning the observability wars"
article_title: "Have you heard? Clickhouse is winning the observability wars!"
author: "backlit4034"
captured_at: "2026-07-09T00:00:14Z"
capture_tool: "hn-digest"
hn_id: 48839068
score: 1
comments: 0
posted_at: "2026-07-08T23:57:21Z"
tags:
  - hacker-news
  - translated
---

# Have you heard? ClickHouse is winning the observability wars

- HN: [48839068](https://news.ycombinator.com/item?id=48839068)
- Source: [charity.wtf](https://charity.wtf/p/have-you-heard-clickhouse-is-winning)
- Score: 1
- Comments: 0
- Posted: 2026-07-08T23:57:21Z

## Translation

タイトル：聞いたことがありますか？ ClickHouse は可観測性戦争に勝利しています
記事タイトル：聞いたことありますか？クリックハウスは可観測性戦争に勝利しています!
説明: 私は何年も前に、カラム型ストレージによって可観測性が作り変えられると予測しました。私が予想していなかったこと: ベンダーはそれを構築し、ナーフし、より悪いバージョンを「Datadog ですが安価」として売り戻すでしょう。

記事本文:
聞いたことがありますか？クリックハウスは可観測性戦争に勝利しています!
チャリティーメジャー
購読 サインイン 聞いたことがありますか?クリックハウスは可観測性戦争に勝利しています!
数年前、私はカラム型ストレージによって可観測性が作り変えられると予測しました。私が予想していなかったこと: ベンダーはそれを構築し、ナーフし、より悪いバージョンを「Datadog ですが安価」として売り戻すでしょう。
慈善専攻 2026 年 7 月 8 日 1 シェア AI の規範と価値観に関する定期的に予定されている一連の投稿を中断して、マット・ダガン氏による素晴らしい記事「クリックハウスが可観測性戦争に勝つ理由」をお届けします。今すぐ読んでください。行く！待ちます。
マット氏はまず、開発者がコマンド ラインにログオンすることによって始まるエクスペリエンスについて説明します。「ファースト キスと同等の可観測性で、その後のすべてが台無しになります」。続いて、4 つのサービスが 40 になり、400 になるにつれて地獄に落ち、朝食前にそれぞれ異なる不可能なことを求めるエンジニアリング、データ、カスタマー サポート、経営陣の関係者を満足させることが不可能になることを説明しました。
Mat 氏が言うように、1 日あたり 1 TB であれば、最新の可観測性スタックはどれも問題ありません。何かを選んで生産性を高めましょう。しかし、1 日あたり 10 TB になると、それらはすべて管理できなくなります。クリックハウスを除くすべてのものであると彼は言います。
(私がこの暴言の中で最も気に入っている部分は、私が悪い人間か SRE であるためです。Elastic、LGTM、Datadog がどのようにして大規模に崩壊するかを彼が細心の注意を払って詳細に説明しているところです。これは傷跡を残した少年です、敬意を表します。)
10TB/日を処理できる可観測性ツールはほとんどありません
データ ツールは、最初は簡単に実行できるものの、チームが大規模に実行する必要があることは、それほど新しいことではありません。ただし、クリックハウス — マットの雄叫びを聞いてみましょう:
「私がこれまで扱ってきた可観測性バックエンドはすべて、成長するにつれて変化します...1 日あたり 10 TB の ClickHouse は、1 TB の ClickHouse と同じように見えます

より多くのシャードを使用して 1 日あたり TB を実行します。それでおしまい。それがピッチです。それが私がこれを書いているすべての理由です。」
Mat 氏は、Clickhouse には前払いで少額の税金がかかることを認めていますが、その価値はあります。なぜなら、わずかな前払いの努力で、高いカーディナリティのデータを使用し、面倒なスキーマのロックインやパフォーマンスの崖がないので、将来的には非常に簡単に拡張できるからです。
うわー、これはクリックハウスの広告のように聞こえてきませんか?なぜ私はこのようにクリックハウスの評判を高めるのでしょうか?
だって、ただのクリックハウスじゃないんだから、この麻痺者め。
柱型ストレージに裏付けられた可観測性は、3 つの柱とは異なるクラスのツールです
マットの作品は、実際には可観測性戦争で誰が「勝者」なのかについてではなく、マットがコラム型ストレージ エンジンによる可観測性を使用することがどのようなものであるかを発見し、それが彼の永遠の心を揺さぶっているというものです。
そしてそうすべきです。それは本当に良くなり、本当にエキサイティングです!
私たちは過去 10 年間、このことについて話し続けてきました。 1 これがなぜ今でも人々にニュースとして伝えられているのでしょうか?
私たちはこれについて『可観測性エンジニアリング』の初版で詳しく書き、本の 10 分の 1 以上をこのトピックに費やしました。第 2 版を執筆する際、私たちは同じ理由で Clickhouse にゲスト章を寄稿するよう依頼しました。これが 20 のクソ 26 年における観測可能性の感じ方です。
「言ったでしょ」とばかり言うつもりはありませんが、見てください。何という違いでしょう！このたわごとはできます！マットのような白髪交じりの丸太戦士さえも夢中にさせることができるなら。 💕
なぜこれについて聞いていないのですか？
マットのような、明らかにこの世界に浸っている人たちでさえ、人々が未だにこのことを理解していないのはなぜか、私には仮説があります。
私の持論は、あるベンダーが言うことはすべて利己的なマーケティングとして無視される、というものです。複数のベンダーが対応するまでは

腕を組んで一緒に同じことを言うと、人々は座って、風景が本当に変わったことに気づきます。
私は、可観測性を備えた企業がカラムナ型ストレージに基づいて構築されるようになれば、自然にそうなるだろうと常々思っていました。私たちは、購入者に状況を明確にして、それがどれほど優れているかを理解してもらうことに強い共通の関心があるため、いくつかの共通の技術用語を揃えることにしました。新興企業対既存企業。ピカピカの新しい方法と、だらしなくて遅くて苦痛を伴う古い方法。
まあ、私は半分正しかったです。 2019 年以降に設立された可観測性企業はすべて、円柱状のストアの上に構築されています。それをご存知ですか?しかし、彼らは私たちと武器を結んでいないだけでなく、自分たちが構築しているものにはそれほど新しいことも違うこともないと主張し続けています。どこも「Datadog をより安価に」販売しています。
なぜ新しい可観測性ベンダーは自社製品を弱体化し、違いを曖昧にするのでしょうか?
結果として、これらの新しい可観測性ツールのほとんどは、カラム型ストレージ エンジン上に構築されているにもかかわらず、Datadog とそれほど変わりません。
昔と同じ三本柱。カーディナリティとスケール、ギブ・アウト・テイクに関する同じ古い問題。彼らがこれを行うのは、Datadog の弱点がすべて価格にあると感じているためです。
私が Datadog で感じる弱点はすべて製品にあります。
はい、カラムナ型ストレージで実行する方がコストが安くなります。しかし、価格の低下はアーキテクチャの向上の結果であり、正直に言って、これは最も興味深い結果の 1 つです。インフラのログとメトリクスはコモディティですが、自社の製品や独自のコードの可観測性はどうでしょうか?それは投資になるはずです。
誰もが Datadog が高価であると考えているのは、高価であることを選択しているからです。これは部分的には真実です。しかし、彼らにも実際には選択の余地はありません。彼らのビジネス モデルは 30 年前のアーキテクチャを中心に成長しており、そのアーキテクチャに完全に固定されています。
W

病気の AI は、最終的に 3 本の柱の死の支配を振り払うのに十分ですか?
AI がワークロードの大部分を占めるようになるにつれて、重要なのはトレースだけであることに人々が気づき始めています。必要なのはトレース、または広範で構造化された正規ログのいずれかだけです。
関係性がデータの価値を高めるため、強力でコンテキストが豊富な 1 つのデータ セットは、その柱を合計したものよりも価値があります。それを柱にスライスすると、その価値は永久に破壊されます。
いずれにせよ、人々はこれを理解するだろうし、それは Datadog のビジネス モデルにとって大きな脅威となる。このビジネス モデルは、同じデータを異なる形式で何度も何度も保存 (および課金) し、データセット間のリンクを保存 (および課金) し、それに付随するすべての手数料とコストに依存している。
あるいは、人々はそれを理解できないかもしれません。神は私がこの詐欺行為がこれほど長く続くとは思ってもいなかったことを知っています。
クリックハウスは可観測性戦争に勝てるでしょうか?神様、そう願っています
Clickhouse は、私の数えた限りでは「可観測性 2.0」という用語を 2 回使用しており、これを使用している唯一のベンダーとなっています。なぜクリックハウスは積極的に武器を提携しようとしているのに、他の誰も提携していないのでしょうか?
私の持論は、彼らは可観測性を追求する企業ではなくデータベース企業であるため、安価な Datadog として自社をマーケティングすることにそれほど執着していないからです。ほとんどの可観測性ベンダーとは異なり、彼らはマーケティングや販売が容易なものを構築することよりも、顧客の問題を解決するものを構築することを優先することにしました。
クリックハウスさん、これを読んでいるならもっとコーディネートしてみたいと思います。電話してね。
そうそう。独自のログ クラスターを実行したい場合は、Clickhouse を使用する必要があります。独自のクラスターの DBA を実行したくない場合や、オーバーヘッド (Mat が指摘しているように線形にスケールしますが、重要です) に対処したくない場合は、Honeycomb を使用する必要があります。

。
20 年前に設計されたツールをまだ使い続けているのであれば、あなたも Clickhouse (または Honeycomb) を試してみるべきかもしれません。
それは本当に全く新しい世界です。マット・ダガンに聞いてください。
その他の参考文献 (包括的とは程遠い):
可観測性におけるコスト危機に関する投稿。3 つの柱のアーキテクチャとカラムをサポートする o11y 2.0 モデルに取り組みます。
2023 年に、カラムナ型ストレージを利用した o11y を「Observability 2.0」と呼び始めました。
Twitter での長い暴言が 1 つ、そしてまた 1 つと続きました。 (多様性を目的とした LinkedIn の暴言。)
Alex Vondrak 著の「 Why Observability Requires A Distributed Column Store」 (個人的なお気に入り)
これは、カラム型ストアによる可観測性の技術的な差別化要因に関する 2018 年の 1 つです。
3 つの柱の代わりに柱型ストアの使用に関する 2020 年の別の投稿
「可観測性 1.0 と 2.0 の重要な違いは 1 つだけです。」 それは統合されたカラム型ストレージです。
「もう一つの可観測性 3.0 が地平線に現れる」
「バージョニングの可観測性について: 1.0、2.0、3.0…10.0」
私は、ベンダーが 2025 年に古い 3 本柱ツールを「すべて洗い直す」ことについて吐き出しました。
Strangeloop 2017 で当社のコラム型ストレージ エンジンについて語る Sam Stokes
Strangeloop 2023 でカラム型データベースをサーバーレス化した方法について語るジェシカ・カー
1 シェア この投稿に関するディスカッション コメント 再スタック トップ 最新のディスカッション 投稿はありません

## Original Extract

Years ago I predicted that columnar storage would remake observability. What I didn't see coming: vendors would build it, nerf it, and sell a worse version back to you as "Datadog, but cheaper"

Have you heard? Clickhouse is winning the observability wars!
Charity Majors
Subscribe Sign in Have you heard? Clickhouse is winning the observability wars!
Years ago I predicted that columnar storage would remake observability. What I didn't see coming: vendors would build it, nerf it, and sell a worse version back to you as "Datadog, but cheaper"
Charity Majors Jul 08, 2026 1 Share We interrupt our regularly scheduled series of posts on AI norms and values to bring you this incredible piece from Mat Duggan, “ Why Clickhouse Is Winning the Observability Wars ,” which you should go read right now. Go! I’ll wait.
Mat starts by describing the experience every developer starts with, logging on the command line — “the observability equivalent of a first kiss, that ruins you for everything after” — followed by the descent into hell as four services becomes forty, becomes four hundred, and the impossibility of satisfying stakeholders in engineering, data, customer support, and executives, who all want different impossible things before breakfast.
As Mat says, at 1TB a day, every modern observability stack is fine. Pick something and be productive. But at 10TB/day, they all become unmanageable — all, he says, except Clickhouse.
(My favorite part of the rant, because I am a bad person and/or an SRE, is where he meticulously details exactly how Elastic, LGTM, and Datadog fall apart at scale. This is a boy with scar tissue, respect.)
Few observability tools can handle 10TB/day
It’s hardly news that data tools tend to start off easy to run, but require teams to run at scale. Clickhouse, though — let’s hear Mat gush:
“Every other observability backend I’ve worked with mutates as it grows... ClickHouse at 10 TB a day looks like ClickHouse at 1 TB a day with more shards. That’s it. That’s the pitch. That’s the whole reason I’m writing this.”
Mat admits that Clickhouse has a small tax up front, but worth it since a tiny bit of upfront effort buys you SO MUCH EASE down the road, as you scale indefinitely, with high cardinality data and no messy schema lock-ins or performance cliffs.
Wow, this is starting to sound like a Clickhouse ad, isn’t it? Why am I boosting Clickhouse's reputation like this?
Because it isn't just fucking Clickhouse, you numbnuts.
Observability backed by columnar storage is a different class of tool than the three pillars
Mat’s piece isn’t actually about who is “winning” the observability wars, it’s about how Mat just discovered what it's like using observability powered by a columnar storage engine, and it is blowing his everloving mind .
And it should. It is genuinely better, and this is genuinely exciting!
And we have been talking about this for the past ten years . 1 How is this still news to people?
We wrote about this at length in the first edition of “ Observability Engineering ”, devoting over a tenth of the book to this topic. When writing the second edition, we invited Clickhouse to contribute a guest chapter for the same reason — THIS IS WHAT OBSERVABILITY SHOULD FEEL LIKE IN TWENTY FUCKING TWENTY-SIX.
I don’t mean to be all “I told you so”, but look! what a difference ! this shit makes! if it can even make a grizzled logs warrior like Mat fall so madly in love. 💕
Why haven’t you heard about this?
I have a theory why it is that people still don’t get this, even people like Mat who are clearly steeped in the space.
My theory is that anything one vendor says gets written off as self-interested marketing. It isn’t until multiple vendors link arms and say the same thing together that people sit up and note that the landscape really has changed.
I always assumed that would happen naturally once more observability companies were built on columnar storage. We would align on some shared technical vocabulary due to our keen shared interest in clarifying the landscape for buyers, to help them understand how much better it is. The upstarts vs the incumbents. The shiny new way vs the kludgey, slow, painful old way.
Well, I was half right. Every observability company founded post-2019 has been built on top of a columnar store, did you know that? But not only have they not linked arms with us, they keep insisting there’s nothing all that new or different about what they’re building. They’re all selling “Datadog, but cheaper.”
Why are newer observability vendors nerfing their own products and obscuring the difference?
As a consequence, most of these newer observability tools are, despite being built on columnar storage engines, not that different from Datadog.
Same old three pillars. Same old problems with cardinality and scale, give or take. They are doing this because the weakness that they sense in Datadog is all price.
The weakness I sense in Datadog is all product.
Yes, it’s cheaper to run on columnar storage. But lower price is a consequence of better architecture, and honestly, it’s one of the least interesting consequences. Infra logs and metrics are commodities, but observability for your own products, your own code? That should be an investment .
Everybody thinks Datadog is expensive because they choose to be expensive. This is true in part. But they also don’t really have a choice. Their business model has grown up around a thirty-year-old architecture, and they are fully locked in to it.
Will AI be enough to finally shake off the death grip of the three pillars?
As AI starts accounting for a larger share of workloads, people are beginning to realize that the trace is all that matters. The trace — or wide, structured canonical logs , one or the other — is all you need .
One powerful, context rich data set is worth more than the sum of its pillars, because relationships are what makes data valuable . Slice it up into pillars, and you destroy that value for good.
One way or another, people are gonna figure this out, and that poses an enormous threat to Datadog’s business model, which relies on storing (and charging for) the same data over and over and over and over again in different formats, then storing (and charging for) links between datasets, and all the attendant fees and costs.
Or maybe people won’t figure it out. God knows I never thought the swindle could go on this long.
Is Clickhouse winning the observability wars? God I hope so
Clickhouse has used the term “ observability 2.0 ” two times by my count, making them the only other vendor who has. Why is Clickhouse willing to link arms, kinda, but no one else has?
My theory is because they are a database company, not an observability company, and therefore less attached to marketing themselves as a cheaper Datadog. Unlike most observability vendors, they have decided to prioritize building something that solves customer problems over building something easy to market and sell.
Clickhouse, if you’re reading this, I would love to coordinate more. Call me.
So yeah. If you want to run your own logs clusters, you should use Clickhouse. If you don’t want to DBA your own clusters or deal with the overhead (which scales linearly, as Mat points out, but is nontrivial), you should use Honeycomb.
If you’re still limping along on tools designed twenty years ago, maybe you too should give Clickhouse -- or Honeycomb -- a shot.
It really is a whole new world. Just ask Mat Duggan .
Further references (far from comprehensive):
A post on the cost crisis in observability , which tackles the three pillars architecture vs the columnar-backed o11y 2.0 model.
We started calling columnar storage-backed o11y “ Observability 2.0 “ in 2023.
One long twitter rant , and another , and another . (A LinkedIn rant for variety’s sake.)
“ Why Observability Requires A Distributed Column Store ”, by Alex Vondrak, a personal favorite
Here’s one from 2018 on the technical differentiators in observability powered by columnar stores
Another post from 2020 on using columnar stores instead of the three pillars
“ There is only one key difference between observability 1.0 and 2.0 ”, and it is unified columnar storage
“ Another observability 3.0 appears on the horizon ”
“ On versioning observabilities: 1.0, 2.0, 3.0…10.0 ”
I vented about vendors “ o11ywashing ” their old three pillars tools in 2025
Sam Stokes talking about our columnar storage engine at Strangeloop 2017
Jessica Kerr talking about how we serverless’d our columnar database at Strangeloop 2023
1 Share Discussion about this post Comments Restacks Top Latest Discussions No posts
