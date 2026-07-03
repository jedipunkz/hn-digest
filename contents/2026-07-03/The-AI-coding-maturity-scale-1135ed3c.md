---
source: "https://blog.codacy.com/the-ai-coding-maturity-scale-the-path-to-loop-engineering"
hn_url: "https://news.ycombinator.com/item?id=48777236"
title: "The AI coding maturity scale"
article_title: "The AI Coding Maturity Scale: The Path to Loop Engineering"
author: "claudiacsf"
captured_at: "2026-07-03T17:13:33Z"
capture_tool: "hn-digest"
hn_id: 48777236
score: 1
comments: 0
posted_at: "2026-07-03T17:01:56Z"
tags:
  - hacker-news
  - translated
---

# The AI coding maturity scale

- HN: [48777236](https://news.ycombinator.com/item?id=48777236)
- Source: [blog.codacy.com](https://blog.codacy.com/the-ai-coding-maturity-scale-the-path-to-loop-engineering)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T17:01:56Z

## Translation

タイトル: AI コーディングの成熟度スケール
記事のタイトル: AI コーディング成熟度スケール: ループ エンジニアリングへの道
説明: Codacy の CEO、Jaime Jorge 氏は、組織の知識を変革するための次のフロンティアとしてのナレッジ ベースについて考えています。

記事本文:
AI コーディング成熟度スケール: ループ エンジニアリングへの道
過去数四半期にわたって、私たちは何百ものソフトウェア チームと、AI コーディング エージェントを実際にどのように使用しているかについて話し合いました。私が驚いたのは、ほとんどのチームが 1 つのリポジトリでオートコンプリートを実行し、エージェントが別のリポジトリで機能全体を構築し、すでに独自のインフラストラクチャに対してループを実験していることです。多くの場合は並行して、場合によっては同じスプリントで実行されます。
AI コーディングの成熟度曲線のマッピング
これは、私たちが聞いた内容を実際のソフトウェア開発ライフサイクル (要件、改良、開発、レビュー、マージ) に照らしてマッピングするのに役立ちます。なぜなら、私たちが話を聞いたほとんどのチームにパターンが当てはまっているからです。
AI は最初にエディター内に出現し、人々がコードを書くときにコードを自動補完しました。この段階は共同作業です。開発者はモデルと並行して改良、開発、レビューを行います。この時点では、それはほぼ普遍的に良いアイデアです。オートコンプリートに反対する人は誰もいませんでした。
次に、コーディング エージェントの登場を促します。エージェントが開発を主導し、開発者の役割は方向性とレビューに移ります。現在、多くのチームがここに座っています。平均的な開発者はすでに週に 20 時間ほどを Claude Code の内部に費やしています。そして、自律性の変化により、実際の作業の多くが下流のレビューに押し付けられています。
その後の段階はループ エンジニアリングです。エージェントは独自のスケジュールで実行され、システムに接続され、ほぼ完全な自律性で動作し、残っているレビュー プロセスに最大限の圧力をかけます。ほとんどのチームはこの方向に向かっており、開発者あたりの予算を削減するのではなく増加させながらそれを行っています。
私にとって最も興味深かったのは、ほとんどのチームが 3 つのステージすべてに同時に参加しているということです。彼らは 1 つのリポジトリでオートコンプリートを実行し、エージェントはすべての機能を 1 つのリポジトリで構築しています。

他にも、すでに独自のインフラストラクチャに対してループを実験しています。一度に 1 歩ずつ登るのではなく、多くの場合は並行して、場合によっては同じスプリントで行われます。
それらすべてに一貫しているのは方向性です。誰もがエージェントの自主性を高め、定期的な機械的な作業を手放し、依然として人間の判断が必要な仕事の部分に人々が移れるようにする準備ができています。ほとんどのチームは、自分たちが思っている以上にその道を進んでいます。
チームがループ エンジニアリングに陥る最も一般的な方法は、劇的なオールイン ベットでもありません。プロダクション ログやバグ レポートを監視し、PR を自動的に開くのはエージェントであり、通常はチームがすでに備えているプロンプト ワークフローと並行して実行されます。そして、ループで実行されるエージェントから出てくる大量のコードにより、チームはすでに、エージェントが継続的にコミットするのではなく、人間が一度に 1 つの PR を開くように構築された git フローを書き直すことを余儀なくされています。
それでも、これには現実的なリスクが伴います。私たちは、企業がエージェントを制御不能に拡張し、その代償としてエントロピーと滑走路コストを後から支払うのを見てきました。これは、蓄積された混乱とバーンアウトの一種で、初日ではなく数か月後に現れます。
これらの会話から学んだことを、AI コーディングの成熟度のスケールを説明する以下のビデオに凝縮しました。これを見れば、おそらく自分のチームがどこに位置しているかを正確に認識できるでしょう。
それが私たちが Verity.md を構築した理由です
Verity.md は、コーディング エージェントのためのゲート、メモリ、コスト管理です。現在はベータ版で、無料で使用でき、インストールは約 2 分で完了します。
AI エージェントは開発者をエンジニアリング オーケストレーターに変え、リスクを検討に移します
同様のパターンについて説明する上級エンジニアが増えています。つまり、AI エージェントがコードの大部分を生成する一方で、手動で少量のコードのみを記述するというものです。
以前は GitHub Copilot Code Review が含まれていましたが、6 月 1 日からは 2 回分の料金が発生します

CE
LLM が成熟し、モデル プロバイダーがデータセンターへの巨額の支出に伴う投資収益率を追求するにつれて、価格の高騰は避けられません。
ブログを購読する
毎月のニュースレターで最新情報を入手してください。
©2026 著作権。無断転載を禁じます。

## Original Extract

Codacy CEO Jaime Jorge reflects on knowledge bases as the next frontier in transforming institutional knowledge.

The AI Coding Maturity Scale: The Path to Loop Engineering
Over the last few quarters, we talked to hundreds of software teams about how they're actually using AI coding agents. What surprised me was that most teams are running autocomplete in one repo, agents building whole features in another, and already experimenting with loops against their own infrastructure. Often in parallel, sometimes in the same sprint.
Mapping the AI coding maturity curve
It helps to map what we heard against the actual software development lifecycle — requirements, refinement, development, review, merge — because the pattern holds up across most of the teams we talked to.
AI showed up first inside the editor, autocompleting code as people wrote it. That stage is collaborative: developers refine, develop, and review alongside the model, and at this point it's pretty much a universally good idea. Nobody really argued against autocomplete.
Then came prompting coding agents , where the agent takes the lead on development and the developer's role shifts toward direction and review. This is where a lot of teams sit today — the average developer is already spending something like 20 hours a week inside Claude Code — and that shift in autonomy has pushed more of the real work downstream, onto review.
The stage after that is loop engineering : agents running on their own schedule, wired into your systems, operating with close to full autonomy and putting maximum pressure on whatever review process is left. Most teams are heading this way, and they're doing it while increasing budget per developer, not cutting it.
The most interesting part to me was that most teams are at all 3 stages at the same time. They're running autocomplete in one repo, agents building whole features in another, and already experimenting with loops against their own infrastructure — often in parallel, sometimes in the same sprint, rather than climbing one step at a time.
What is consistent across all of them is the direction. Everyone is ready to hand agents more autonomy, hand off the recurring and mechanical work, and let people move up to the part of the job that still needs a human judgment call. Most teams are further down that road than they realize.
The most common way teams end up in loop engineering isn't a dramatic all-in bet either. It's agents watching production logs or bug reports and opening PRs automatically, usually running alongside the prompting workflows the team already has. And the volume of code coming out of agents running in loops is already forcing teams to rewrite git flows that were built for humans opening one PR at a time, not agents committing continuously.
Still, this comes with real risk. We've watched companies scale agents uncontrollably and pay for it later in entropy and runway cost, which is hat's the kind of accumulated mess and burn that shows up months in, not on day one.
I condensed what I learned from these conversations into the video below which describes the AI coding maturity scale. Watch it and you'll probably recognize exactly where your team sits.
And that's why we built Verity.md
Verity.md is gates, memory, and cost control for coding agents. It's in beta now, free to use, and installs in about two minutes.
AI Agents Are Turning Developers Into Engineering Orchestrators and Moving the Risk to Review
A growing number of senior engineers describe a similar pattern: writing only a small amount of code manually while AI agents generate much of the...
GitHub Copilot Code Review used to be included, from June 1st you pay twice
As LLMs mature and the model providers seek a return on their investment following huge spending on datacentres, it’s inevitable that the price of...
Subscribe to our blog
Stay updated with our monthly newsletter.
©2026 Copyright. All Rights Reserved.
