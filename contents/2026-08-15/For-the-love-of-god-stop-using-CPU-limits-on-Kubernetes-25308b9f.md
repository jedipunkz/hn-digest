---
source: "https://home.robusta.dev/blog/stop-using-cpu-limits"
hn_url: "https://news.ycombinator.com/item?id=49308986"
title: "For the love of god, stop using CPU limits on Kubernetes"
article_title: "For the Love of God, Stop Using CPU Limits on Kubernetes (Updated) — Robusta Blog"
author: "inigyou"
captured_at: "2026-08-15T09:16:54Z"
capture_tool: "hn-digest"
hn_id: 49308986
score: 2
comments: 0
posted_at: "2026-08-15T09:04:01Z"
tags:
  - hacker-news
  - translated
---

# For the love of god, stop using CPU limits on Kubernetes

- HN: [49308986](https://news.ycombinator.com/item?id=49308986)
- Source: [home.robusta.dev](https://home.robusta.dev/blog/stop-using-cpu-limits)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T09:04:01Z

## Translation

タイトル: 念のため、Kubernetes での CPU 制限の使用をやめてください
記事のタイトル: 神の愛のために、Kubernetes での CPU 制限の使用をやめてください (更新) — Robusta Blog
説明: Kubernetes では CPU 制限が必要だと多くの人が考えていますが、これは真実ではありません。ほとんどの場合、Kubernetes の CPU 制限は役立つどころか害を及ぼします。実際、これらは Kubernetes CPU スロットルの最大の原因です。

記事本文:
神の愛のために、Kubernetes での CPU 制限の使用をやめてください (更新) — Robusta ブログ 価格設定 ケーススタディ ブログ 今すぐお試しください ログイン 価格 オンライン ゲーム (Betsson) ブログ Slack に参加してください ログイン ブログに戻る 2022 年 8 月 24 日
神の愛のために、Kubernetes での CPU 制限の使用をやめてください (更新)
Kubernetes には CPU 制限が必要だと多くの人が考えていますが、これは真実ではありません。ほとんどの場合、Kubernetes の CPU 制限は役立つどころか害を及ぼします。実際、これらは Kubernetes CPU スロットルの最大の原因です。
Natan Yellin 著、Robusta.dev 共同創設者
Kubernetes の CPU 制限はアンチパターンです
Kubernetes には CPU 制限が必要だと多くの人が考えていますが、これは真実ではありません。ほとんどの場合、Kubernetes の CPU 制限は役立つどころか害を及ぼします。
CPU 制限が有害である理由を、CPU 不足のポッドと砂漠で道に迷った喉の渇いた探索者の 3 つの類似点を使って説明します。勇敢な探検家をマーカスとテレサと呼びます。
私たちの物語では、CPU は水になり、CPU 飢餓は死になります。 CPU と同様に、この物語における水も再生可能な資源になります。簡単に言うと、ある瞬間の CPU 使用率が 100% であっても、次の 1 分間は CPU を「使い切る」ことはありません。 CPU は刻一刻と更新されます。
Kubernetes の CPU 制限に関する 3 つのカラフルな類似点
マーカスとテレサは砂漠を旅しています。彼らは1日3リットルを生み出す魔法の水筒を持っています。人が生きていくためには1日1リットルが必要です。
ストーリー 1 - 制限なし、リクエストなし: マーカスは貪欲なので、テレサが水を飲む前に水をすべて飲み干します。テレサは喉の渇きで亡くなります。制限や要求がなかったからだ。マーカスは水をすべて飲むことができ、CPU 飢餓を引き起こし、テレサは CPU スロットルを受けました。
ストーリー 2 - 制限あり、リクエストの有無にかかわらず: テレサはある日重い病気になり、追加の水が必要になります。マーカスは1リットルの飲み物を飲みます

残りは2リットルです。テレサは 1 リットル飲みますが、残り 1 リットルになります。マーカスはテレサにそれを飲ませません。彼女の制限は1日あたり1リットルであるため、テレサは喉の渇きで死んでしまいます。これは、CPU 制限がある場合に起こることです。リソースは利用可能ですが、使用することは許可されていません。
ストーリー 3 - 制限なし、リクエストあり: マーカスはある日、重篤な病気になり、余分な水が必要になります。彼はボトルを全部飲もうとしましたが、ボトルに残ったのは1リットルだけでした。テレサは 1 日に 1 リットル必要なので、これは保存されています。彼女は1リットル飲みます。何も残りません。彼らは両方とも生きています。これは、CPU 制限がないにもかかわらずリクエストがある場合に起こります。すべて順調です。
上記の話は、CPU 制限が有害であると考えられる理由を驚くほど正確に類推しています。
例え話が好きではありませんか？ Prometheus の CPU スロットル アラート CPUThrottlingHigh を説明するときに、これに関するより技術的な説明を書きました。
制限のない CPU スロットルと CPU 不足の防止
多くの人は、あるポッドが別のポッドに干渉しないように制限が必要だと考えています。これは真実ではありません! Kubernetes の CPU 制限を削除しても、CPU を大量に消費するポッドによる CPU 枯渇を防ぐことができます。重要なのは、CPU リクエストを定義するだけです。
以下は、元の Kubernetes ドキュメントの抜粋です。
重要な部分が強調表示されます。ポッドは、CPU リクエストによってリクエストされた CPU を常に取得します。 (制限がない場合は、余剰 CPU も活用できます。)
CPU 制限を無効にしてください。すべての K8 ポッドに正確な CPU リクエストを与えると、必要に応じて CPU が予約されるため、誰もそれらをスロットルできなくなります。これは制限とは何の関係もありません。
CPU 制限とリクエストの違いを表にまとめてみましょう。
わかりやすくするために、すべてのポッドに CPU 制限があるか、どのポッドにも CPU 制限がないかのどちらかであると仮定します。

えっと。 CPU リクエストについても同様です。これにより、すべてのことを推論するのが簡単になります。
Kubernetes での CPU 制限とリクエストのベスト プラクティス
すべてに CPU リクエストを使用します (設定に関するヘルプが必要な場合は、「 KRR 」を参照してください)
Tim Hockin (Google の元の Kubernetes メンテナの 1 人) は、何年も同じことを推奨してきました。
メモリ制限とリクエストについてはどうですか?
この記事の内容はすべて CPU に関するものであり、メモリに関するものではありません。メモリは非圧縮性であるため異なります。メモリを一度与えると、プロセスを強制終了することなくそれを取り除くことはできません。ここでは、Kubernetes のメモリ制限に関するベスト プラクティスについて説明しました。つまり、私たちの最終的な推奨事項は次のとおりです。
メモリ要求は常に制限値と同じに設定してくださいS
過去の Prometheus データに基づいて CPU とメモリの要求を判断できるように、オープン ソース ツールを構築しました。
Kubernetes Resource Recommender - 履歴データに基づいて CPU とメモリのリクエストを決定します
GitHub で確認してください。質問がある場合は、LinkedIn でメッセージを送信するか、GitHub の問題を開いてください。
Natan Yellin、CEO — Natan は 15 年以上ソフトウェアを書いてきました。彼は、Kubernetes に関するあらゆることについて定期的に LinkedIn に投稿しています。
あなたの環境で実行していることを確認してください。
Robusta をクラスターにインストールし、実際のインシデントを確認できるようお手伝いします。
まずはセットアップについて教えていただけますか?

## Original Extract

Many people think you need CPU limits on Kubernetes but this isn't true. In most cases, Kubernetes CPU limits do more harm than help. In fact, they're the number one cause of Kubernetes CPU throttling.

For the Love of God, Stop Using CPU Limits on Kubernetes (Updated) — Robusta Blog Pricing Case studies Blog Try now Login Pricing Online Gaming (Betsson) Blog Join our Slack Login Back to blog Aug 24, 2022
For the Love of God, Stop Using CPU Limits on Kubernetes (Updated)
Many people think you need CPU limits on Kubernetes but this isn't true. In most cases, Kubernetes CPU limits do more harm than help. In fact, they're the number one cause of Kubernetes CPU throttling.
By Natan Yellin , Robusta.dev co-founder
CPU limits on Kubernetes are an antipattern
Many people think you need CPU limits on Kubernetes but this isn't true. In most cases, Kubernetes CPU limits do more harm than help.
I will explain why CPU limits are harmful with three analogies between CPU starved pods and thirsty explorers lost in a desert. We will call our intrepid explorers Marcus and Teresa.
In our stories, CPU will be water and CPU starvation will be death. Like CPU, water in our story will be a renewable resource. In simple terms, if you have 100% CPU usage at a given minute, that doesn't "use up" the CPU for the next minute. CPU renews itself from moment to moment.
Three colorful analogies about Kubernetes CPU Limits
Marcus and Teresa are travelling in the desert. They have a magical water bottle that produces 3 liters a day. Each person needs 1 liter a day to survive.
Story 1 - without limits, without requests: Marcus is greedy so he drinks all the water before Teresa can drink any. Teresa dies of thirst. This is because there were no limits or requests. Marcus was able to drink all the water, cause CPU starvation, and Teresa was CPU throttled.
Story 2 - with limits, with or without requests: Teresa gets very ill one day and needs some extra water. Marcus drinks his one liter and there are two liters left over. Teresa drinks one liter and now there is one liter remaining. Marcus wont let Teresa drink it because her limit is 1 liter per day so she dies of thirst. This is what happens when you have CPU limits. Resources are available but you aren't allowed to use them.
Story 3 - without limits, with requests: Marcus gets very ill and needs extra water one day. He tries to drink the entire bottle but is stopped when only 1 liter remains in the bottle. This is saved for Teresa because she needs 1 liter a day. She drinks her 1 liter. Nothing remains. They both live. This is what happens when you have no CPU limits but you do have requests. All is good.
The above stories are surprisingly precise analogies for why CPU limits are considered harmful.
Don't like analogies? I wrote a more technical explanation of this when explaining the Prometheus CPU throttling alert, CPUThrottlingHigh .
Preventing CPU throttling and insufficient CPU without limits
Many people think you need limits to prevent one pod from interfering with another pod. This is not true! You can remove Kubernetes CPU limits and still prevent a CPU hungry pod from causing CPU starvation! The trick is to just define CPU requests.
Here is a snippet from the original Kubernetes documentation :
The important part is highlighted. Pods always get the CPU requested by their CPU request! (They can take advantage of excess CPU too if they have no limit.)
Disable your CPU limits! If you give all your K8s pods accurate CPU requests, then no-one can throttle them because CPU is reserved for them if they need it. This has nothing to do with limits.
Lets summarize the difference between CPU limits and requests in a table:
I'm assuming for simplicity that either all pods have CPU limits or none have them. The same is true for CPU requests. This makes everything simpler to reason about.
Best practices for CPU limits and requests on Kubernetes
Use CPU requests for everything (if you need help setting them, see KRR )
Tim Hockin (one of the original Kubernetes maintainers at Google) has recommended the same for years.
What about memory limits and requests?
Everything in this post is about CPU and not memory. Memory is different because it is non-compressible - once you give memory you can't take it away without killing the process. We've covered the best practices for Kubernetes memory limits here. In short, our bottom line recommendation is:
Always set your memory requests equal to your limitsS
We built an open source tool so you can determine CPU and memory requests based on historical Prometheus data.
Kubernetes Resource Recommender - determine CPU and memory requests based on historical data
Check it out on GitHub . For questions, message me on LinkedIn or open a GitHub issue.
Natan Yellin, CEO — Natan has been writing software for over 15 years. He regularly posts about all things Kubernetes on LinkedIn .
See it running in your environment.
We'll help you get Robusta installed on your cluster and walk through a live incident.
Prefer to tell us about your setup first?
