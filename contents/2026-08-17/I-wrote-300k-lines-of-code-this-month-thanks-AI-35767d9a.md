---
source: "https://nmn.gl/blog/how-i-actually-write-10000-lines"
hn_url: "https://news.ycombinator.com/item?id=49334975"
title: "I wrote 300k+ lines of code this month, thanks AI"
article_title: "I wrote 300k+ lines of code this month thanks to AI | N’s Blog"
image: "https://nmn.gl/blog/assets/how-i-actually-write-10000-lines/product-work-2.png"
author: "namanyayg"
captured_at: "2026-08-17T18:23:08Z"
capture_tool: "hn-digest"
hn_id: 49334975
score: 2
comments: 0
posted_at: "2026-08-17T17:53:56Z"
tags:
  - hacker-news
  - translated
---

# I wrote 300k+ lines of code this month, thanks AI

- HN: [49334975](https://news.ycombinator.com/item?id=49334975)
- Source: [nmn.gl](https://nmn.gl/blog/how-i-actually-write-10000-lines)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T17:53:56Z

## Translation

タイトル: 今月 30 万行以上のコードを書きました、AI のおかげです
記事のタイトル: AI のおかげで、今月 30 万行以上のコードを書きました | Nさんのブログ
説明: pg が昨年これを投稿したときの内容を最近リツイートしたため、私が出荷するコードの行数を確認することになりました。以下に写真を添付し​​ました。ほんの数年前であれば、これは 1 か月の LoC 数としては十分な数値だったことでしょう。もちろん、数万行のコードを配布することは「n」とみなされます。
[切り捨てられた]

記事本文:
AI | のおかげで、今月 30 万行以上のコードを書きました。 Nさんのブログ
さんのブログ
Nさんのブログ -->
AI のおかげで、今月 30 万行以上のコードを書きました
pg が昨年これを投稿したときの内容を最近リツイートしたので、私は自分が出荷しているコードの行数を確認しました。
以下に写真を添付し​​ました。ほんの数年前であれば、これは 1 か月の LoC 数としては十分な数値だったことでしょう。
もちろん、数万行のコードを出荷することは、最近では「新しい常態」とみなされています。私は CEO としての成長に重点を置いているので、これは今のところそれほど高いとは考えていません。それでも、12,000/日を見ると笑ってしまいます。
HackerNews の多くの人が懐疑的であることに気づいたので (当然のことですが)、私が実際に 1 日に出荷するものについて少しお話します。
AI をコード生成の効果的なツールとして使用することと、やみくもに実行することは、コンテキストとテスト ハーネスに細心の注意を払う必要があることを意味します。だからこそ、私にはこの 1 年間で磨いてきた数え切れないほどのスキルと脚本があります。
私は通常、ChatGPT アプリで Codex の 3 つまたは 4 つのセッションを使用し、Ghostty で Claude の 4 つの端末を使用しています。
私はオーケストレーターを使用していません。昨年いくつかのオーケストレーターを試しましたが、バグが多くて遅かったです。そのカテゴリー全体が私の口に酸っぱい味を残しました。
また、Devin はクラウド エージェントによる優れたブラウザ自動化機能を備えており、AI がクラウド上で UI 作業とテストを実行できるため、Devin も使用しています。
私は SaaS ソフトウェア内に埋め込まれた AI エージェントに取り組んでいます。当社の顧客は SaaS 企業であり、顧客がダッシュボードやアプリを構築できるようにするために当社を利用しています。
最近では、時間をかけて第一原則から考え、アーキテクチャとデータ モデルを実際に改善し、プラットフォームをより構成可能でセルフサービスにしやすくしています。
あなたがプログラマーであれば、それがどのように機能するかを知っているでしょう。MVP プロダクトを作成するのはあなたです。

しかし、ユーザーが実際にそれを使用すると、元のアーキテクチャの設計時に誤解していたことが非常に多くあることに気づきます。
私たちには実際の顧客がいるため、実際の使用状況からモデル化された多くの合成テスト ケースやシナリオを生成できます。それを参考にして厳密なテストを行うことで、AIに信頼性の高いコードを生成させることができます。適切な抑制と計画があれば、エージェントを一晩で解雇し、それでも効果的な解決策を考え出すのを見るのは本当に感動的です。
ジェボンズのパラドックス 1 が実際に展開されているのを実際に見ています。このリファクタリングは、AI 以前の世界には存在しませんでした。 AI がなければ、古いコードベースでもっと長く作業を続けていたかもしれません。しかし、自信を持ってより多くのコードを作成できるため、私たちはリファクタリングに果敢に取り組んでいます。
おそらくこれは、私が最近よく行ってきたエンジニアリングの中で最も予想外のタイプのものでした。 2
GTM モーションが成熟するにつれて、機能する重要な要素と機能しない要素がいくつかわかってきたので、これらをうまく実行するツールの作成に着手しました。数週間社内で使用していますが、返信率が 2 倍になりました。あまり詳しくは説明しませんが、その背後にあるエンジニアリングの努力がわかるように、私が構築したものの概要をここに示します。
すべては、私が構築したものを必要としていると思われる人々を見つけることから始まります 3 。 Sales Navigator、LinkedIn Search、および私の過去の営業データを使用して、AI が似たような人を見つけます。リードにはさまざまな状態があり、リードのソースもさまざまであるため、これには多くのキュー作業が必要でした。
その後、LinkedIn が自動化され、接続リクエストの送信、いいね、メッセージングなど、数日間にわたるシーケンスが実行されます。これは非常に複雑なステート マシンでした。外から見ると、ステートがいくつかしかないように見えますが、

実際には、対応すべき障害状態はさらに多くあります。
それぞれのシーケンスを厳密にテストし、返信率と予約された会議から統計的に有意な結果を見つけようとします。
また、競合他社を自動的に識別し、その投稿をフォローし、それに関与している人々を見つけて、関連性を考慮してフィルターをかけ、シーケンスを送信します。
そして最も良い点は、誰かが返信するたびに、Slack 上で私に ping を送り、自動的に返信の下書きを作成してくれることです。
私は自分自身のユーザーだったため、アーキテクチャ、データ モデル、可観測性の設計に慎重に時間を費やしました。これは、多くの AI を使用しているにもかかわらず、ほとんどの「雰囲気コード化された」プロジェクトがどのようなものであるかを期待するものとは異なる方法で構築したことを意味します。何人かの友人に試してもらいましたが、素晴らしい反応が得られました。
私が説明したループに興味があり、購入者を見つけたり、メッセージをテストしたり、誰かが返信したときに Slack で ping を受け取りたい場合は、仕事用の電子メールまたは電話をそのままにしておいてください。それがどのように機能するかを説明します。
すぐに手を差し伸べます！または、以下のスロットを選択してください。手順を説明します。
または、お送りするメールにご返信ください。
電話を予約していただければ、それがあなたのビジネスにどのように役立つかを説明させていただきます。
大規模な顧客への販売の一環として、顧客の認証と API を統合するために多くの作業を行う必要があります。
複雑な作業です。これには、従来のコードベースと難解なデータ ソースを掘り下げることが含まれます。それは本質的に、大企業が私たちに支払っている対価です。
ここ数週間だけで、次のようなことが起こりました。
異なる顧客のデプロイメントで異なる認証シークレットが使用されていた問題を解決しました。しかし、それらは非常によく似ており、そのコードを書いた元の開発者はとうの昔に亡くなっていたため、問題を解決するまで誰も実際に知ることができませんでした。
ごとに異なるスキーマを使用したスノーフレーク データベース接続のサポートが追加されました。

テナント
1 つの管理者トークンを使用してユーザーのサブアカウントの特定の権限を取得する統合認証
ある顧客は、古いキャッシュではうまく処理できなかった 50MB 範囲の静的ファイルを所有していたため、まったく新しいキャッシュ システムをセットアップしました。
この仕事は平凡ではありますが、データ モデリング、パフォーマンス、品質のトレードオフに関する問題をどのように包含しているかという点でやりがいがあり、私の中のエンジニアを興奮させます。
多くの場合、多くのドキュメントを調べ、試行錯誤によってドキュメント化されていないコードを発見し、リグレッションを防ぐために厳密な API テストを作成する必要があるためです。このユニークな組み合わせが AI に適しているため、ここでも大幅に高速化することができました。
では… 10,000 行は多いのでしょうか?
人々が何万行も書くのを目にするほとんどの場合、それは新しいコードやグリーンフィールドのプロジェクトになります。
おそらくサードパーティの統合と関係があり、簡単にテストして迅速に検証できるものになる可能性が最も高くなります。
生成されたコードが価値のあるものになるのか、成果を生み出すのかはまだ不確実です。AI は依然としてあなたの脳を精神病に陥らせる機会を数多く生み出しており、そこでは、あなたが構築するすべての新しい機能が世界を変えるものであると信じ込まされます。
しかし、それは一部の人にとっては問題ありません。この若い創業者は絶望的で、お金を払ってくれる顧客がいないため、迅速に行動し、本当に物事を壊すことができます。
もちろん、あなたが大企業の経験豊富なプログラマーである場合、仕事の重要な部分は、抱えている何千もの顧客の業務を中断しないことです。したがって、あなたにとって、10,000 行のコードを出荷するのはおそらく非常に悪い考えです。
親愛なる読者の皆さんはどうですか？ (効果的な) コードは何行書く必要がありますか?AI によってそれはどう変わりましたか?
どれが本当にBになるのか

私はジェボンの法則と呼ばれています↩
クレイが数年前にこの言葉を発明したと思います。しかし、今ではさらに進化しています！ ↩
マーケティング用語では「リード」とも呼ばれますが、私はこの言葉があまり好きではありませんでした ↩
私は、ソフトウェア販売チームが見込み顧客が求める機能を構築できるように、製品ギャップによって契約を失うことがないように、Gigacatalyst を構築しています。私たちは Y Combinator によって支援されています。どのように機能するか見てみたいですか?
ちなみに私はナマニヤイです。プログラミングは私の初恋で、10 代の頃からコードを書いてきました。
この投稿をシェアする
← Y Combinator に興味を持った経緯についての考察
×
(無料) 開発ガイドを入手する
プログラミングと AI に関する無料ガイドを利用すると、より早く出荷できます。
10 年以上のソフトウェア構築から学んだ教訓が含まれています。
スパムメールは送りません。いつでも購読を解除してください。
営業チームにスーパーパワーを与える
Gigacatalyst を使用すると、営業チームが見込み顧客が求める機能を構築できるため、製品ギャップによって契約を失うことがなくなります。
私の顧客は、これが 2026 年の売上を増やす最善の方法だと言います。
さんのブログ
ナマニャイによる AI、スタートアップ、人生についての考え。

## Original Extract

pg recently retweeted this from when he posted this last year, which made me check out how much lines of code I ship. I’ve attached a picture below. Only a couple of years ago this would’ve been a good LoC count for a month! Of course, shipping tens of thousands of lines of code is considered the “n
[truncated]

I wrote 300k+ lines of code this month thanks to AI | N’s Blog
’ s Blog
N’s Blog -->
I wrote 300k+ lines of code this month thanks to AI
pg recently retweeted this from when he posted this last year, which made me check out how much lines of code I ship.
I’ve attached a picture below. Only a couple of years ago this would’ve been a good LoC count for a month!
Of course, shipping tens of thousands of lines of code is considered the “new normal” these days. I focus on growth as a CEO, so this isn’t something I’d consider very high now. But I still got a chuckle out of seeing 12,000/day.
I notice a lot of people on HackerNews being skeptical (rightfully so), so I’ll tell you a little bit about what I actually ship in a day.
Using AI as an effective tool for code generation vs blindly vibing means you have to pay close attention to context and testing harnesses. That’s why I have countless skills and scripts that I’ve been honing over the last year.
I usually have 3 or 4 sessions of Codex in the ChatGPT app, and 4 terminals of Claude on Ghostty.
I don’t use an orchestrator, I tried some of them last year and they were buggy and slow. That entire category left a sour taste in my mouth.
We also use Devin because it has great browser automation with cloud agents, so the AI is able to do UI work and testing on the cloud.
I’ve been working on an AI agent embedded inside SaaS software. Our customers are SaaS companies, who use us to allow their customers to build dashboards and apps.
These days, I’m taking the time to think from first principles and really improve the architecture and data model, making our platform much more configurable and easier to self-serve.
If you’re a programmer, you know how it works: you create the MVP product with some assumptions, but when your users actually use it, you realize that there are so many things you misunderstood while designing the original architecture.
Since we have real customers, we are able to generate a lot of synthetic test cases and scenarios modeled from our real-world usage. Using that as a reference and strict testing, we are able to get AI to generate reliable code. With the correct restraints and plans, it’s really impressive to fire off agents overnight and still see them come up with working solutions!
I’m really seeing Jevons paradox 1 play out in action. This refactor wouldn’t exist in a pre-AI world. We’d probably would’ve chugged along for longer in our old codebase without AI – but because we can write more code with confidence, we are braving a refactor.
This, probably, has been the most unexpected type of engineering that I’ve been doing a lot recently. 2
As our GTM motion has matured, I have figured out some key things that work and don’t work, and I’ve set out to create a tool that does these really well. I’ve been using it internally for the couple of weeks and have doubled my reply rates! I won’t go into too much detail, but here’s a brief overview of what I built so you can see the engineering effort behind it.
It all starts by finding people who I think want what I’ve built 3 . Using Sales Navigator, LinkedIn Search, and data from my past sales, the AI finds similar people. This involved a lot of queuing work, because there’s various states a lead can be in and sources it can come from.
It then automates my LinkedIn to run multi-day sequences: sending a connection request, liking, messaging, and so on. This was a very complicated state machine – from the outside it looks to have only a few states, but really there are many more failure states to accommodate.
It rigorously tests each sequence out, and tries to find statistically significant results from reply rates and meetings booked.
It also automatically identifies my competitors, follows their posts, finds people who engaged with it, filters them for relevancy, and sends them sequences.
And the best part – whenever someone replies, it pings me on Slack and automatically drafts a reply for me.
I spent careful time designing the architecture, data model, and observability – because I was my own user. This means that despite using a lot of AI, I built it differently than what you’d expect most “vibe coded” projects to be like. I’ve been letting a few friends try it and it’s gotten amazing reactions!
If you're interested in the loop I described and want find buyers, test messages, and get pinged on Slack when someone replies -- Leave your work email or phone and I'll show you how it'd work for you.
I'll reach out soon! Or, you can grab a slot below and I’ll walk you through it.
Or just reply to the email I’ll send you.
Book a call and I’ll walk you through how it would work for your business.
As part of selling to larger customers, we have to do a lot of work integrating with their authentication and APIs.
It’s complex work. It involves digging through legacy codebases and esoteric data sources. That’s essentially what the larger businesses pay us for.
In just the last couple of weeks, we have:
Resolved an issue where different customer deployments used different authentication secrets – but they look very similar and the original dev who wrote that code was long gone so no one really got to know till we broke stuff
Added support for snowflake database connections with different schemas per tenant
Integrated authentication where one administrator token is used to get specific privileges for user sub-accounts
Set up a brand new caching system because one customer had static files in the 50mb range that our old cache wasn’t handling well.
Although mundane, this work is rewarding because of how it encompasses problems around data modelling, performance, quality tradeoffs, which excite the engineer in me.
Because it often requires going through a lot of documentation, discovering undocumented code by trial and error, and writing rigorous API tests to prevent regressions; this unique combination makes it apt for AI, which is why we’ve been able to accelerate a lot here as well.
So… Is 10,000 lines like a lot?
In almost all of the cases where you’re seeing people write tens of thousands of lines, it is going to be new code and greenfield projects.
It probably will have something to do with third-party integrations and most likely to be something that is easy to test and verify quickly.
Whether the generated code will be valuable or create outcomes is still uncertain: AI still creates a lot of opportunities to lose your brain to psychosis, where the sycophancy convinces you that every new feature you build is world-changing.
But that’s fine for some people. The young founder is desperate and has no paying customers, so they can Move Fast And Really Break Things™.
Of course, if you’re an experienced programmer in a larger company, a critical part of your job is to not break things for the thousands of customers you have. So for you, it is probably a very bad idea to ship 10,000 lines of code.
What about you, my dear reader, where do you stand? How many lines of (effective) code do you to write, and how has that changed with AI?
which out to really be called Jevons law tbh ↩
I think Clay invented the term a few years ago. But it is much more advanced now! ↩
a.k.a. “leads” in marketing-speak, but I never really liked that word ↩
I'm building Gigacatalyst to help software sales team build features that prospects ask for, so they never lose a deal to product gaps. We're backed by Y Combinator. Want to see how it works?
I’m Namanyay by the way. Programming is my first love and I've been writing code since I was a teenager.
Share this post
← Meditations on How I Got Into Y Combinator
×
Get my (free) development guide
Ship faster with my free guides on programming and AI.
Includes lessons learnt from 10+ years of building software.
I don't spam. Unsubscribe at any time.
Give your sales team superpowers
Gigacatalyst allows your sales team to build features that prospects ask for, so they never lose a deal to product gaps.
My customers say that this is the best way to increase sales in 2026.
' s Blog
Thoughts on AI, startups, and life by Namanyay.
