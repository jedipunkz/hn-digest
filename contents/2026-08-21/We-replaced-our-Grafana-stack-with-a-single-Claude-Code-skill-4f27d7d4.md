---
source: "https://frigade.com/blog/we-replaced-grafana-with-a-claude-code-skill"
hn_url: "https://news.ycombinator.com/item?id=49393346"
title: "We replaced our Grafana stack with a single Claude Code skill"
article_title: "We replaced our massive Grafana stack with a single Claude Code skill — Frigade"
image: "https://frigade.com/blog/we-deleted-a-dashboard-we-trusted/cover-v10.png"
author: "pancomplex"
captured_at: "2026-08-21T21:14:55Z"
capture_tool: "hn-digest"
hn_id: 49393346
score: 1
comments: 0
posted_at: "2026-08-21T20:23:43Z"
tags:
  - hacker-news
  - translated
---

# We replaced our Grafana stack with a single Claude Code skill

- HN: [49393346](https://news.ycombinator.com/item?id=49393346)
- Source: [frigade.com](https://frigade.com/blog/we-replaced-grafana-with-a-claude-code-skill)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T20:23:43Z

## Translation

タイトル: Grafana スタックを単一の Claude Code スキルに置き換えました
記事のタイトル: 大規模な Grafana スタックを単一の Claude Code スキルに置き換えました — Frigade
説明: Grafana ダッシュボードには、手動で作成された 24 個のグラフと数年分のデータがありました。私たちはそれを 1 つのクロード コード スキルに 1 日で置き換えました。

記事本文:
大規模な Grafana スタックを単一の Claude Code スキル「Frigade」に置き換えました。
製品を学習し、リアルタイムでユーザーをガイドする Products Assistant AI。オンボーディングや製品ツアーに Drop-in React コンポーネントを活用します。仕組み 価格設定 アップデート ブログ Login Assistant について AI アシスタントの管理 エンゲージメント オンボーディング フローの構築 始めましょう 製品
大規模な Grafana スタックを単一の Claude Code スキルに置き換えました。
Grafana ダッシュボードには、手動で作成された 24 個のグラフと数年分のデータが含まれていました。私たちはそれを 1 つのクロード コード スキルに 1 日で置き換えました。
9 か月間、私たちは大規模なマルチリージョン Grafana セットアップを実行しました。先週、スタック全体を、読み取り専用の Postgres レプリカを通じて同じデータをクエリする Claude Code スキルに置き換えました。
最初に Grafana ダッシュボードを構築した理由
あなたがどこかでデータ チームやビジネス運用チームと働いていて、そこにいる誰かがグラフを作成してくれたら、それを信頼します。彼らがそれをチェックしたと信じます。彼らは、全員がどの行をカウントから除外したかを知っていると信じられます。開くたびに同じことが書かれており、それが間違っているときは、それを見ているすべての人にとって同じように間違っています。
これが、事前に構築されたグラフの真の価値です。柔軟性はありませんが、信頼性はあります。
土曜日に家の整理整頓をすることを考えてみると、食器はある場所に、バッテリーと充電器は別の場所にという具合に、ほとんどのことがうまくいきます。その後、最後のアームフルに到達します。巻尺と家の鍵にフィットしない鍵。もう遅いです。どちらもあなたが作成したものには属しません。2 つのオブジェクトにカテゴリを発明するのはばかげているように思えます。なので一番下の引き出しに入ります。
6 か月後、最初に開けたのは一番下の引き出しです。
私たちのダッシュボードも同じように進みました。合計 24 個のグラフがありました。 10月に設置したセクションは開催され、その後私たちが疑問に思っていたことのほとんどは開催されました。

その他にまとめました。
それらはすべてプルリクエストでした。誰かが時間をかけてクエリを作成し、グラフを選択し、1 つの画面上の 2 つの数値が一致するようにその周囲のグラフにフィルターを照合しました。その後、次の質問が来て、もう一度質問しました。
つまり、ダッシュボードは 10 月に私たちが抱いた質問に答えており、一番下にあるのはそれ以来のすべてでした。私たちが望んでいたのは、1 年後に誰かが保守しなければならないようなものを構築せずに、そのうちの 1 つに依頼する方法でした。
置き換えられたものは信頼性が低い
5 月に、Postgres レプリカの前に読み取り専用サーバーを置き、わかりやすい英語で質問をし始めました。あらゆる方向に柔軟です。最初に何も構築せずに、これまで尋ねたことのないことを尋ねたり、1 週間に 1 人の顧客についての回答を得ることができます。
そして時々、それはあなたに何かを伝えます、そしてあなたはそれを読んで、それは正しくないはずだと思います。だからあなたはそう言うのです。そして、「あなたの言う通りです、私は間違いを犯しました。」と言うのです。
それは問題になるほど頻繁に起こります。私たちはそれを解決しておらず、決定に向かうものは依然としてデータベースと直接照合されます。
ダッシュボードを開けていなかったからです。
ダッシュボードを見てこれに答えましょうとは誰も言いません。実際に抱いている疑問が 1 つのグラフにマッピングされることはほとんどありません。特定の週の特定の顧客、または特定の規模以上のアカウント間でのみの傾向、または誰も思いつかない方法で先月の数字を分割したい場合があります。
フィルターと日付ピッカーを使用すると、それに近づけることができます。私たちは何ヶ月もそれを行いました。それは機能しますが、退屈ですが、退屈なだけで、誰かが質問するのをやめてしまうほどです。私たちが今尋ねていることのほとんどは、誰もグラフを作成しなかったであろう小さな質問です。
いくつかのことはまだあなたに押し付けられるべきです
数値がしきい値を超えたかどうかを尋ねる必要はありません。

毎日チェックしてます。そこで、その部分を保持して Slack に移動しました。ジョブは平日の毎朝実行され、概要が投稿されます。そのクエリは毎回新たに作成されるのではなく、リポジトリ内に存在するため、毎日同じ質問が尋ねられ、同じ答えが返されます。
Frigade を使用している人にも同じことが当てはまります。統計情報を含むダッシュボードと、ユーザーの質問をテーマごとにグループ化するインサイト ページが表示され、どちらも便利です。
そのダッシュボードのFrigade Assistantは、データを直接クエリすることもできます。アシスタントが先週どうだったか、人々がどの回答に低評価を示したか、またはあなたが出荷した変更によって何か変化があったかどうかを尋ねてください。実際の質問やユーザーへのリンクが返されるので、自分で探しに行くことができます。
アシスタントが特定のユーザー データを読み取ったり、必要のないものをクエリしたりできないように、強力な保護を導入しています。これらはすべて読み取り専用レイヤーに配置されているため、データベースを変更することはできません。
これは、同じ理由で社内で行ったのと同じ取引です。貴社の製品について、私たちが決して推測するつもりのなかった質問があります。
製品を学習し、ユーザーに教える AI エージェント。
Frigade は、製品内に組み込まれた AI アシスタントです。デモを予約すると、一般的なウォークスルーではなく、実際のアプリでどのように動作するかを示します。
2026 年 5 月 11 日 私たちは Framer (そして他の人もそうするでしょう) Code を辞め、その後 2 年間 Framer を辞め、その後コードに戻りました。私たちが移行した理由、そしてノーコード ツールを使用するほとんどのチームがこれに従う理由。
2026 年 4 月 28 日 AI エージェントが応答を拒否すべき場合 AI エージェントは応答するように設定されています。良い人は、いつ断るべきかを知っています。 「わかりません」が構築に最もコストがかかる機能であり、ユーザーが最も信頼する機能である理由。
2025 年 12 月 9 日 外堀がひっくり返りました: ヘルプセンターの機能を壊すために、より迅速な配送が行われていました。

ヘルプ ドキュメントとオンボーディングを維持するチーム。数学が逆転してしまいました。迅速な発送は顧客にとっての資産となり、ドキュメント チームの責任ではなくなりました。
サンフランシスコで建造 © Frigade Inc.

## Original Extract

Our Grafana dashboard had 24 manually crafted graphs and years of data. We replaced it in a day with a single Claude Code skill.

We replaced our massive Grafana stack with a single Claude Code skill — Frigade
Products Assistant AI that learns your product and guides users in real time. Engage Drop-in React components for onboarding and product tours. How It Works Pricing Updates Blog About Login Assistant Manage your AI assistant Engage Build onboarding flows Get Started Products
We replaced our massive Grafana stack with a single Claude Code skill
Our Grafana dashboard had 24 manually crafted graphs and years of data. We replaced it in a day with a single Claude Code skill.
For nine months, we ran a large multi-region Grafana setup. Last week, we replaced the whole stack with a Claude Code skill that queries the same data through a read-only Postgres replica.
Why we originally built the Grafana dashboard
If you work somewhere with a data team, or a biz ops team, and someone there builds you a graph, you trust it. You trust that they checked it. You trust they know which rows everybody leaves out of the count. It says the same thing every time you open it, and when it is wrong it is wrong the same way for everyone looking at it.
That's the real value of a pre-built graph. It's not flexible, but it's reliable.
Think of spending a Saturday organizing the house and most of it goes fine: utensils in one place, batteries and chargers in another. Then you get to the last armful. A tape measure and a key that fits no lock in the house. It's late, neither one belongs in anything you just built, and inventing a category for two objects feels ridiculous. So it goes in the bottom drawer.
Six months later the bottom drawer is the one you open first.
Our dashboard went the same way. We had 24 total graphs. The sections we set up in October held, and most of what we wondered about after that ended up in Miscellaneous.
Every one of those was a pull request. Somebody spent time, wrote the query, picked a chart, and matched the filters to the graphs around it so that two numbers on one screen would still agree. Then the next question arrived and we did it again.
So the dashboard answered the questions we had in October, and the pile at the bottom was everything since. What we wanted was a way to ask one of those without building something somebody would have to maintain a year later.
What replaced it is less reliable
In May we put a read-only server in front of that Postgres replica and started asking it questions in plain English. It's more flexible in every direction. We can ask things we never asked before, or get an answer about one customer in one week without building anything first.
And sometimes it tells you something, and you read it and think that can't be right. So you say so. And it says, "You're absolutely right, I made a mistake."
That happens often enough to matter. We have not solved it, and anything heading into a decision still gets checked against the database directly.
Because we weren't opening the dashboard.
Nobody says let's go look at the dashboard and answer this. The question you actually have almost never maps onto one graph. You want a specific customer in a specific week, or a trend but only across accounts over a certain size, or last month's number split a way nobody thought to split it.
You can get close with filters and date pickers. We did that for months. It works, and it is tedious, and tedious is enough to stop somebody asking at all. Most of what we ask now is small questions nobody would have built a graph for.
Some things should still be pushed at you
You shouldn't have to ask whether a number crossed a threshold, and nobody checks every day. So we kept that part and moved it into Slack. A job runs each weekday morning and posts a summary. Its queries live in the repo instead of being written fresh each time, so the same questions get asked every day and the answers come back the same way.
The same is true for anyone using Frigade. You get a dashboard with your stats, and an Insights page that groups what people are asking into themes, and both are useful.
Frigade Assistant in that dashboard can also query your data directly. Ask how your assistant did last week, or which answers people gave a thumbs down to, or whether the change you shipped moved anything. It comes back with links to the actual questions and users, so you can go look for yourself.
We put strong protections in place so the assistant can't read specific user data or query anything it shouldn't. It all sits on a read-only layer, so it can't change our database.
It's the same trade we made internally, for the same reason. You have questions about your own product that we were never going to guess.
The AI agent that learns your product and teaches your users.
Frigade is an AI assistant that lives inside your product. Book a demo and we'll show you how it works on your actual app, not a generic walkthrough.
May 11, 2026 We left Framer (and others will, too) Code, then Framer for two years, then back to code. What made us move, and why most teams on a no-code tool will follow.
April 28, 2026 When an AI agent should refuse to answer AI agents are wired to answer. The good ones know when to refuse. Why "I don't know" is the most expensive feature to build, and the one users trust most.
December 9, 2025 The moat just flipped: shipping faster used to break your help center For a decade, every product release was a new tax on the team that maintains help docs and onboarding. The math has inverted. Shipping fast is now an asset for the customer, not a liability for the docs team.
Built in San Francisco © Frigade Inc.
