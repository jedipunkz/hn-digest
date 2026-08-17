---
source: "https://newoldweb.com/control-charts-make-ai-agents-less-necessary-and-more-useful"
hn_url: "https://news.ycombinator.com/item?id=49338144"
title: "Control Charts Make AI Agents Cheaper, Less Necessary, and More Useful"
article_title: "Control Charts Make AI Agents Cheaper, Less Necessary, and More Useful - New Old Web"
image: "https://lede-admin.newoldweb.com/wp-content/uploads/sites/65/2026/08/deming-title-card.jpg"
author: "netaustin"
captured_at: "2026-08-17T22:14:18Z"
capture_tool: "hn-digest"
hn_id: 49338144
score: 1
comments: 0
posted_at: "2026-08-17T21:48:14Z"
tags:
  - hacker-news
  - translated
---

# Control Charts Make AI Agents Cheaper, Less Necessary, and More Useful

- HN: [49338144](https://news.ycombinator.com/item?id=49338144)
- Source: [newoldweb.com](https://newoldweb.com/control-charts-make-ai-agents-less-necessary-and-more-useful)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T21:48:14Z

## Translation

タイトル: 管理図により AI エージェントが安価になり、必要性が低くなり、より便利になります
記事のタイトル: 管理図により AI エージェントが安価になり、必要性が低くなり、より便利になります - New Old Web
説明: 新しい管理図によると、このブログに投稿してから 35 日が経過しました。主観的には長すぎます。客観的に見て、それは私が毎朝受け取るメールで警鐘を鳴らしました。私は長い間、「物事がどのように進んでいるのか」について、より良く、より全体的な全体像を知りたいと思っていましたが、次のようなエージェントに満足するのではなく、
[切り捨てられた]

記事本文:
コンテンツ メニューにスキップ ログイン 検索 検索 ホーム
管理図により AI エージェントが安価になり、必要性が減り、より便利になります
2026 年 8 月 17 日午後 9 時 21 分（東部夏時間）
新しい古い Web を Google に追加 新しい管理図によると、このブログに投稿してから 35 日が経過しました。主観的には長すぎます。客観的に見て、それは私が毎朝受け取るメールで警鐘を鳴らしました。私は長い間、「物事がどのように進んでいるのか」について、より良く、より全体的な全体像を知りたいと思っていましたが、多くの人がそうしているように、エージェントに妥協するのではなく、エージェントが存在するよりもはるかに長い間私が抱いていた目標にクロード・コードを適用しました。それは、私の生活のあらゆる部分の側面を示す管理図を構築することです。私はエージェントを構築するつもりはありませんでしたが、現在ではこれらすべての優れたデータ ソースがエージェント サーフェスにあります。管理図は、私とクロードの両方に議論するための参照フレームを提供し、コンテキストにより、クロードは私にとって非常に有益になります。
正直に言うと、すべての情報を接続し、有用な一連のグラフに統合するためのインフラストラクチャは難しい部分です。私は何年もこれをやりたいと思っていましたが、Claude Code を手に入れるまでは自分でコードを書く時間がありませんでした。クロードは、私のビジョンを実現するために必要な API グルーや Python チャート作成が非常に得意です。非常に低コストの Plaid のようなバンキング コネクタ、Garmin データの回避策、Apple Health のエクスポート フォーマットなど、私が探さなければならなかったサービスが見つかりました。
これを構築するのにどれくらい時間がかかったのかはわかりません。私がそれを構築し終えることはできなかったことがわかっているからです。しかし、今では、これを手に入れたので、私の高度に測定された人生がこれまでダッシュボードの列島に閉じ込められていたことに気づきました。
管理図を使用すると、管理下のシステムを視覚的に監視できます。 （管理下のシステムは私です。）

y 経時的な観測値を測定し、ヒューリスティックに基づいて警報を発します。私はビジネス スクールでそれらについて学び、私のソフトウェア開発者の脳に訴えかけました。運用管理は一般に非常に満足のいくものでしたが、管理図の事例のほとんどは製造業から来ています。しかし、W. エドワーズ・デミング研究所は、その最も基本的な入門書の中で、数値は「あらゆるものから」得ることができると述べています。勝負を受けて立つ。
私の管理図は、Donald Wheeler が説明したように、移動範囲の平均を使用して自然プロセスの上限と下限を設定します。従来の実装では、標準偏差 (シグマ) を使用して平均値付近の自然なプロセス限界を描画していましたが、Wheeler 氏は、このアプローチでは信号がノイズに組み込まれてしまうと主張しました。クロードはウィーラーのアプローチの方が好きで、私もクロードの推論が気に入りました。
管理図の目的は、物事がどのように起こっているかを過度に認識させることではなく、すぐに自信を持って先に進み、実際に問題を引き起こしているシステムの部分だけを調査できるようにすることです。
たとえば、私の Garmin ウォッチは睡眠の質を記録しており、朝一番に私の睡眠状態を知らせてくれるようです。答えが「あまり良くない」だと、気が散ってしまい、一日の始まりにやる気を失ってしまう可能性があります。自己評価の方が重要だという人もいますが、私も概ね同意します。 Garmin は睡眠を 100 点満点で評価し、それが「ボディバッテリー」に適用されます。バッテリーは、日中の体の使い方に応じて消耗します。体のバッテリーが私の睡眠が40ポイントに相当すると言うときは気分が良く、眠れない夜を意味しますが、1日の始まりが90ポイントの場合はぐっすりとしていて、素晴らしい睡眠を意味します。ランナーとして、私にとっては、時計の心拍数が示す瞬間よりも、走っているときにどのように感じるかがはるかに重要です。しかし、私が収集したデータは

ランニングは長期的な目標設定と計画に非常に役立ちます。管理図を使用すると、短期間の心拍数のスパイクのほとんどを無視し、ビープ音を重要ではないものとして監視することができます。これにより、より重要な問題が存在する場合、管理図がそれを警告してくれることがわかります。
これはエージェントと何の関係があるのでしょうか?
これらの管理図は、エージェントが解決してくれると私が想像していたクラス全体の問題を解決し、クロードとさらに話し合いたくない限り、トークンを使用しません。将来のトークンエコノミクスが劇的に変化した場合、他の100万人の個人エージェントが機能不全に陥ることになるでしょう。しかし、私の管理チャートは、毎朝更新するジョブがループ内にLLMを必要としないため、依然として動作し続けます。
ただし、これが実際にエージェントとどのような関係があるかは、すべてコンテキストの中にあります。私は毎日エージェントを呼び出すわけではありませんが、管理図の作成から始まったクロード コード ワークスペースは、物事がどのように進んでいるのかについての継続的なディスカッションへと変化しました。私がトレーニングのスケジュールを変更するか再構築するかを考えているとき、クロードは私のフィットネスと疲労のスコアを将来に予測し、私の目標を検討し、役立つ答えをくれます。別のチャットに移動したり、スキルを構築したり、ダッシュボード アプリにチャット レイヤーを構築したりすることはできますが、なぜでしょうか?これは私のためのもので、Claude Code は私の携帯電話またはデスクトップの Claude アプリにあります。
ここには、持久力トレーニングに LLM を使用する機会と課題について、まったく別の質問が含まれていますが、私がやっていることが、機能リストの一番上に「強度バランス」を掲げているマット フィッツジェラルドの Mind of Matt のような、世の中のより思慮深い AI トレーニング アプローチと噛み合っている側面があるのではないかと思います。私は、エージェントがどのようにしてランナーをオーバートレーニングや怪我に誘導してきたかについて多くの記事を読んでおり、次のことに注意しています。

ここでのリスクについては。これまでのところ、管理図はクロードと私の両方の意見を一致させる制約を定義しているようです。私がランニングに関してコントロールしていることの 1 つは、毎週のランニング時間のうち、ゾーン 1 またはゾーン 2 に占める割合が何パーセントであるかということです。これにより、相応の量の増加なしで、よりハードなセッションをやりすぎないようにすることができます。ただし、増加しすぎると疲労チャートでアラームが発生するため、あまり増加させないでください。
AI エージェントが必要な場合もあれば、必要ない場合もありますが、いずれにせよ、管理図が必要になる場合があります。
無料のニュースレターに登録する
ニュース製品スタートアップの資本の罠
独立 250 年を振り返るニュース製品の視点
ニュースへの支払いの歴史: Allegro の視聴者を構築する理由
Allegro Audience: パブリッシャー向けのオープンな顧客データ ツールキット
労働集約的なビジネスに AI を導入するための 5 段階のアプローチ
バイブコーディングは、同じ古いことを言う別の方法です
情報とインターネットのより良い未来、Alley 氏と Lede 氏
無料のニュースレターに登録する

## Original Extract

According to my new control charts, it’s been 35 days since I posted to this blog. Subjectively, that’s too long. Objectively, it raised an alarm in an email I get every morning. I have long wanted a better and more holistic picture of “how things are going” but instead of settling for an agent like
[truncated]

Skip to Content Menu Log In Search Search Home
Control Charts Make AI Agents Cheaper, Less Necessary, and More Useful
9:21 PM EDT on August 17, 2026
Add New Old Web to Google According to my new control charts, it’s been 35 days since I posted to this blog. Subjectively, that’s too long. Objectively, it raised an alarm in an email I get every morning. I have long wanted a better and more holistic picture of “how things are going” but instead of settling for an agent like many folks seem to have done, I applied Claude Code to a goal I have had for much longer than agents have existed: Building control charts that show aspects of every part of my life. And while I did not set out to build an agent, I have all these great sources of data in an agentic surface now. The control charts give both me and Claude a frame of reference to discuss, and the context makes Claude a whole lot more useful to me.
Honestly, the infrastructure to connect all my information and coalesce it into a useful series of charts is the hard part. I have wanted to do this for years but haven’t had the time to write the code myself until I got my hands on Claude Code. Claude is exceptionally good at the sort of API-glue and Python charting required to bring my vision to life. It found services that I would have had to hunt for — a Plaid-like banking connector that costs very little, Garmin data workarounds, Apple Health export formats.
I don’t know how long this would have taken me to build because I know that I wouldn’t have finished building it, but now that I have it, I realize that my highly-measured life had until now been trapped in an archipelago of dashboards.
Control charts let you visually monitor systems under management. (The system under management is me.) They measure observations over time and raise an alarm based on heuristics. I learned about them in business school and they appealed to my software developer brain — operations management was generally very satisfying, but most of the case examples for control charts come from manufacturing. In its most basic primer , however, the W. Edwards Deming institute says the numbers can be "from anything." Challenge accepted.
My control charts use the mean of moving ranges to establish upper and lower natural process limits as described by Donald Wheeler . Conventional implementations use standard deviation (sigma) to draw natural process limits around a mean, but Wheeler argued that this approach bakes the signal into the noise. Claude liked Wheeler’s approach better and I liked Claude’s reasoning.
The point of a control chart isn’t to make you hyperaware of how things are going, it’s to let you develop confidence quickly so you can move on, and investigate only the parts of the system that actually cause a problem.
For example, my Garmin watch tracks sleep quality, and first thing in the morning, it likes to tell me how I’ve slept. If the answer is “not great,” it can be a distracting and demoralizing way to start my day. Some would say that self-rating is more important, and I generally agree. Garmin rates sleep on a 100-point scale that then gets applied to a “body battery." Your battery depletes based on how you use your body during the day. I have felt fine when my body battery says my sleep was worth 40 points, implying a restless night, and crappy when it starts the day at 90, implying great sleep. As a runner, how I feel while I’m running is much more important to me than what the heart rate on my watch says, in the moment . But the data I collect on the run is very useful to long-term goal-setting and planning. Control charts give me permission to ignore most short term heart rate spikes and watch beeps as insignificant, knowing that the charts will alert me to a more important issue, if one exists.
What does this have to do with agents?
These control charts solve a whole class problems that I had envisioned an agent solving for me, and they use no tokens unless I want to discuss them further with Claude. If future token economics change dramatically, a million other personal agents will go dark, but my control charts will still be humming along because the job that updates them each morning requires no LLM in the loop.
What this really has to do with agents though is all in the context. While I don’t invoke an agent every day, the Claude Code workspace that started with building the control charts has morphed into an ongoing discussion about how things are going. When I’m thinking about whether to reschedule or restructure a training run, Claude can project my fitness and fatigue scores into the future, consider my goals, and give me a helpful response. I could move it to another chat, build a skill, build a chat layer into the dashboard app, but why? This is just for me, and Claude Code is right there in the Claude app on my phone or desktop.
There’s a whole separate set of questions implied here about the opportunities and challenges of using LLMs for endurance training, but I suspect there are aspects of what I’m doing that mesh with some of the more thoughtful AI training approaches out there like Mind of Matt from Matt Fitzgerald which lists "intensity balance" right on the top of its list of features. I’ve read a lot about how agents have misled runners into overtraining and injury and am mindful of the risks here. Thus far, the control charts seem to be defining constraints that keep Claude and me both in line. One control I have for running is what percentage of my running time each week is in Zone 1 or Zone 2. This keeps me from overdoing harder sessions without a commensurate ramp in volume — but not too much of a ramp, because that would cause an alarm in my fatigue chart.
Maybe you need an AI agent, maybe you don’t — but either way, you might like a control chart.
Sign up for our free newsletter
The Capital Trap for News Product Startups
A News Product Perspective on Two Hundred Fifty Years of Independence
A Short History of Paying for News: Why We’re Building Allegro Audience
Allegro Audience: An open customer data toolkit for publishers
A Five-Step Approach to Deploying AI in Labor-Intensive Businesses
Vibe Coding is Just Another Way to Say the Same Old Thing
A better future for information and the internet, by Alley and Lede
Sign up for our free newsletter
