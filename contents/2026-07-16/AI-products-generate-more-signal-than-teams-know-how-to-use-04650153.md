---
source: "https://www.armank.dev/thoughts/2"
hn_url: "https://news.ycombinator.com/item?id=48936686"
title: "AI products generate more signal than teams know how to use"
article_title: "arman's living room"
author: "armank-dev"
captured_at: "2026-07-16T17:04:37Z"
capture_tool: "hn-digest"
hn_id: 48936686
score: 3
comments: 0
posted_at: "2026-07-16T16:26:43Z"
tags:
  - hacker-news
  - translated
---

# AI products generate more signal than teams know how to use

- HN: [48936686](https://news.ycombinator.com/item?id=48936686)
- Source: [www.armank.dev](https://www.armank.dev/thoughts/2)
- Score: 3
- Comments: 0
- Posted: 2026-07-16T16:26:43Z

## Translation

タイトル: AI 製品は、チームが使い方を知っている以上に多くのシグナルを生成します
記事タイトル: アルマンのリビングルーム
説明: アルマンのリビングルーム

記事本文:
アルマンのリビングルームの家 / 思考 / 2
AI 製品は、チームが使い方を知っている以上に多くのシグナルを生成します
AI 製品は、それ以前のソフトウェアよりも多くの有用な信号を生成しますが、実際にそれを使用している人は (ほぼ) いません。
それは間違いであり、驚くほど高価なものだと思います。
AI がユーザー インタラクションの形をどのように変えたか
過去 20 年間、当社は間接信号を中心に製品を構築してきました。ユーザーはここをクリックし、そのページを表示し、このボタンの上にカーソルを置きましたが、押しませんでした。これらのパンくずリストからユーザーの意図と行動を推測するのが非常に上手になりました。アンケートを実施していない限り、これがユーザーを理解するための唯一のデータ ポイントでした。
しかし、AI 製品の場合、ユーザーは自然言語を介して対話できるようになりました。この新しい形式のユーザー入力は、はるかに便利です。ユーザーが何を達成したいのか、満足しているのかを正確に読み取ることができます。
これを他のシグナル (賛成票/反対票、オンライン評価、技術指標など) と組み合わせると、エージェントのパフォーマンスと収益の間の明確な関係が得られます。
しかし、実際にデータをまとめている人はほとんどいないように感じます。平均的なチームには、基本的な APM ソフトウェアとエージェント トレースがセットアップされています。しかし、ユーザーの意図、解約シグナル、維持率、顧客満足度を分析する場合、データは比較的不十分です。
なぜ今日、人々はこれを修正しないのでしょうか？
まず最大の理由は、データがチーム間で実際に断片化されていることです。
- エンジニアリング部門はトレースと技術的なメトリクスを所有します。
- 製品はユーザーのフィードバックと分析を所有します。
- カスタマー サクセスは、チャーン シグナルと収益拡大を担当します。
これらを並行して分析することを実際に担当する人は誰もいません 1 。しかし、さらに深く掘り下げると、これらすべてのデータがエージェントのパフォーマンスと収益の相互関係につながることがわかります。
2回目

ason の特徴は、このデータを収集して効果的に分析するには多大な労力がかかるということです。注意しないと、すぐに高価になってしまう可能性があります。
これは実際にはどのように見えるか
これを実践し、これらのデータ ポイントを実際に活用すると、関係がより明確になります。
平均的なユーザーは、エージェントにユーザーの口調でメールを書き直すよう依頼するのに 3 ターンを費やしているようです。これを改善すれば、おそらく電子メール サービスの普及が促進されるでしょう。
なんと、問題 XYZ に遭遇したユーザーの 35% は二度と戻ってこない傾向にあるのです。その問題を調査すれば、おそらくチャーンを削減できるでしょう。
430 人を超えるユーザーがエージェントにインターネットの検索を依頼しましたが、エージェントはまだ検索機能を提供していません。これを当社の製品に追加する必要があります。
私はこれを「全体的な観察可能性」と呼び始めました。エージェントを追跡するだけではなく、一連のデータ ポイントを収集します。
- エージェントの追跡と会話
- ユーザーの投票/フィードバック (単純な賛成/反対のデータで十分です)
- オンライン評価/トレースの監視
- ユーザープロンプト上のトピッククラスタリング
- ユーザープロンプトに関する感情分析 (不満、混乱など)
やがて、信号が現れ始めます。その結果、組織全体のチームにとって明確で実用的な洞察が得られます。顧客の成功により、既存のアカウントの解約シグナルと収益拡大の機会を明確に特定できます。製品チームはユースケースと、どのユースケースが十分なサービスを受けられていないのかを明確に理解しています。エンジニアリング チームは、単純なエラー追跡を超えた実際のエージェントのパフォーマンスに関する洞察を得ることができます。
私の最も強い意見は、AI スタートアップにおいて、エージェントの可観測性をエンジニアリングのみの懸念に留めるべきではないということです。 CEO、製品リーダー、顧客対応チームもこのデータを読む必要があります。
私たちは過去 20 年間、私たちの考えを推測することに費やしました。

クリックからの募集です。今、彼らは私たちに直接伝えています。
聞かないのはおかしいでしょう。
おそらく、私がこれまでにこれほどうまくやっている役割は、非常に初期のスタートアップの創業者かエンジニアだけだと思います (初期のスタートアップではエンジニアがこれらの役割をすべて担う傾向があるため)。 ↩

## Original Extract

arman's living room

arman's living room home / thoughts / two
AI products generate more signal than teams know how to use
AI products generate more usable signal than any software before them, yet (almost) nobody is actually using it.
I think that’s a mistake, and a surprisingly expensive one at that.
how AI changed the shape of user interaction
For the last twenty years, we built products around indirect signals. A user clicked here, viewed that page, hovered over this button but didn't press it. We became really good at inferring user intent and behavior from these breadcrumbs! Unless you were running surveys, this was your only data point for understanding users.
But with AI products, users now interact via natural language. This new form of user input is far more useful. We can read exactly what users want to accomplish, and whether or not they’re satisfied.
When you pair this with other signals (thumbs up/down votes, online evals, technical metrics, etc), you get a clear relationship between the performance of your agent and revenue.
But it feels like almost nobody is actually putting the data together. The average team has some basic APM software and agent tracing set up. But when it comes to analyzing user intent, churn signals, retention, and customer satisfaction, the data falls relatively short.
why aren't people fixing this today?
The first and largest reason is that data is really fragmented across teams:
- Engineering owns tracing and technical metrics.
- Product owns user feedback and analytics.
- Customer Success owns churn signals and revenue expansion.
Nobody is actually in charge of analyzing these in tandem 1 . But if you dig deeper, you'll find that all of this data leads to interlinked relationships between agent performance and revenue.
The second reason is that it simply takes a lot of effort to collect and effectively analyze this data. And if you’re not careful, it can get expensive fast.
what this looks like in practice
When you do put this into practice and truly capitalize on these data points, the relationships become a lot more clear:
It seems like the average user spends 3 turns asking the agent to rewrite the email in the user’s tone. We can probably boost adoption of our email service if we improve this!
Wow, 35% of users who encounter problem XYZ tend to never come back! We can probably cut churn if we investigate that problem.
Over 430 users have asked the agent to search the internet, but the agent doesn’t search capabilities yet. We should add this to our product offering!
I’ve started calling this “holistic observability”. More than just tracing your agent, you collect a suite of data points:
- agent traces and conversations
- user votes/feedback (simple thumbs up/down data is enough)
- online evals/monitoring over traces
- topic clustering over user prompts
- sentiment analysis on user prompts (frustration, confusion, etc)
Eventually, signals start to show themselves. The result is clear and actionable insights for teams across the organization. Customer success can clearly identify churn signals and opportunities for revenue expansion in existing accounts. Product teams have a clear understanding of use cases, and which ones are underserved. Engineering teams get insight into real agent performance that extends beyond simple error tracking.
My strongest opinion is that in AI startups, agent observability should not remain an engineering-only concern. CEOs, product leaders, and customer-facing teams should be reading this data, too.
We spent the last twenty years trying to infer what users wanted from clicks. Now they're telling us directly.
It would be strange not to listen.
Perhaps the only roles I’ve seen do this well are founders or engineers at extremely early startups (since engineers tend to fill all of these roles at early startups). ↩
