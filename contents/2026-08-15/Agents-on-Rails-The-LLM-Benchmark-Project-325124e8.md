---
source: "https://rubyonrails.org/2026/8/12/llm-benchmarking-project"
hn_url: "https://news.ycombinator.com/item?id=49307720"
title: "Agents on Rails: The LLM Benchmark Project"
article_title: "Agents on Rails: The LLM Benchmark Project"
author: "ksec"
captured_at: "2026-08-15T05:16:38Z"
capture_tool: "hn-digest"
hn_id: 49307720
score: 1
comments: 0
posted_at: "2026-08-15T04:47:24Z"
tags:
  - hacker-news
  - translated
---

# Agents on Rails: The LLM Benchmark Project

- HN: [49307720](https://news.ycombinator.com/item?id=49307720)
- Source: [rubyonrails.org](https://rubyonrails.org/2026/8/12/llm-benchmarking-project)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T04:47:24Z

## Translation

タイトル: Agents on Rails: LLM ベンチマーク プロジェクト
説明: 今日、私たちは Agents on Rails の最初の結果を共有します。これは、今日の主要なエージェント コーディング ツール (フロンティアおよびオープンウェイトの両方) が実際に Ruby on Rails コードベースでどの程度パフォーマンスを発揮するかを測定するための、現在進行中の新しい取り組みです。

記事本文:
Agents on Rails: LLM ベンチマーク プロジェクト
本日、私たちは Agents on Rails の最初の結果を共有します。これは、今日の主要なエージェント コーディング ツール (フロンティアおよびオープンウェイトの両方) が Ruby on Rails コードベースで実際にどの程度パフォーマンスを発揮するかを測定するための、現在進行中の新しい取り組みです。
Rails Foundation はこのプロジェクトを Evil Martians に委託しており、今後数週間にわたっていくつかの段階で展開される予定です。
以下のプロジェクトの詳細を読んだり、リーダーボードをチェックしたり、最初のベンチマーク レポートに直接ジャンプしたりできます。
過去 1 年間でコーディング エージェントの使用が急増し、ほぼすべての開発者やチームが AI を使用してコードを作成しています。しかし、選択肢は圧倒的で、ほぼ毎週新しいモデルが発売されます。エージェントの運営コストとトークンの支出も、ほとんどの企業が昨年まで予算を立てていなかった大幅な追加コストです。
そこで、質問がありました。どのモデルが優れた Rails コードを生成しましたか?各モデルのトークンのコストはどれくらいですか?フロンティアモデルは優れているのか、劣っているのか、それとも無差別モデルでも良いのでしょうか？どちらかを使用することにトレードオフはありますか?これらはすべてのエンジニアリング チームが現在取り組んでいる問題であるため、私たちはまた、コミュニティがこれらの決定を下す際にどのように支援できるだろうかと考えました。
Rails プログラマーたちが、Rails と AI がいかに優れているかを語っているのを聞いてきました。Convention over Configuration のおかげで、Rails と AI はうまく連携できる、Rails はトークン効率が良い、モデルは Rails をうまく書く傾向にあるなど、リストは続きます。 Rails Foundation 理事会は、チーム内でエージェントを使用した自身の経験から、これらの意見を裏付けています。
私たちはそれを聞くのが大好きでしたが、逸話以上の証拠が必要でした。そこで私たちは、Evil Martians に、実際の Rails アプリケーションの実際の Rails タスクで人気のあるエージェントのベンチマークを依頼し、それを改善しました。

これは、チームがモデルを選択するときに使用できる実践的なガイダンスになります。
このベンチマークは段階的に構築されており、本日ステージ 1 が開始されます。
ステージ 1: アトミックなタスク。現在の eval セットは、1 つの特定の機能を分離する小さな自己完結型タスクで構成されており、それらが API に触れると、それがその 1 つだけになります。 8 つの異なるフロンティア モデルとオープンウェイト モデルにわたって精度、速度、トークン消費量、コストをテストすることに加えて、モデルが現在の Rails API に対応しているかどうかもチェックしています。
完全な方法論はここにあります。
このベンチマークは、有望な新しいモデルがリリースされるたびに実行されるため、Ruby コミュニティは、モデルが相互にどのように積み重なるか、および考慮すべきコストと機能のトレードオフに関する最新情報を常に入手できます。
現在のリーダーボードは新しい AI ページで見つけることができ、ベンチマークを実行するたびに、完全な結果レポートが添付され、そのページからリンクされます。また、これらのレポートをブログでいつでも見つけられるように、これらのレポート用の新しいタグ「エージェント」も作成しました。
ステージ 2: より現実的な作業。 (近日公開予定) ここでは、孤立したタスクを超えて、より現実的で複雑な作業に移行します。つまり、エージェントが長時間の複数ステップの作業でどのようにパフォーマンスを発揮するかをテストし、機能を追加し、アプリを最初から構築します。これらのテストはエージェントにとってより困難であり、あなたとあなたのチームが本番環境でアプリに取り組む方法をより忠実に模倣することになります。
すべてをオープンソース化します。タスクと方法論は現在利用可能であり、生の実行は明日リポジトリにアップロードされます。まもなく、Evil Martians がこれらすべてを実行するために構築した Ruby ハーネスである lemans も、コミュニティ向けにオープンソース化される予定です。
Evil Martians チームに多大な感謝を申し上げます。彼らはベンチマーク ハーネスとコーパスを設計および構築し、評価を実行し、

その結果は、今日発表する最初の調査結果につながります。最初にチームに話をするべきだということを忘れずに、熱心にチームをプロジェクトに参加させてくれたスヴャトスラフ・クリュコフ、アルトゥール・ペトロフ、ウラジミール・デメンティエフ、アルバート・パズデリン、アレクサンダー・バイゲルディン、アントン・センコフスキー、イリーナ・ナザロワに感謝します。幸運なことに、チームも同様にこのプロジェクトに熱心でした。私たちはそれを見るのが大好きです。
また、新しい AI ページのデザインとレポート結果の視覚化に協力してくれた Rob Zolkos にも感謝します。
次に何をテストすべきかについて考えがある Rails チームの場合、または単純に意見を知らせたい場合は、foundation@rubyonrails.org までご連絡ください。
この取り組みは、Rails Foundation コアおよび貢献メンバーからの継続的なサポートとインプットによって可能になりました。 4 年前、Rails Foundation が発足したとき、私たちはマーケティング、ドキュメンテーション、教育、イベントなど、Rails エコシステムをサポートし成長させるという使命に向かって取り組んでいる、いくつかの主要な柱に焦点を当てることを念頭に置いていました。
エージェント支援コーディングの台頭以来​​、お客様にとって状況は変化しており、それは私たちにとっても変化していることを意味します。使命は変わりませんが、仕事は進化しています。参照ライブラリはその方向への一歩でした。これもそのような取り組みです。さらに続きますので、お楽しみに。

## Original Extract

Today we’re sharing the first results of Agents on Rails, a new, ongoing initiative to measure how well today’s leading agentic coding tools (both frontier and open-weight) actually perform on Ruby on Rails codebases.

Agents on Rails: The LLM Benchmark Project
Today we’re sharing the first results of Agents on Rails , a new, ongoing initiative to measure how well today’s leading agentic coding tools (both frontier and open-weight) actually perform on Ruby on Rails codebases.
The Rails Foundation commissioned Evil Martians for this project, which will roll out in several stages over the next few weeks.
Read more about the project below, check out the leaderboard , or jump right over to the first benchmark report .
The use of coding agents has skyrocketed in the past year with nearly every developer or team using AI to write code. But the options are overwhelming, with new models dropping nearly every week. The cost of running agents and token spend is also a significant additional cost most companies hadn’t budgeted for before last year.
So, we had questions. Which models produced good Rails code? How many tokens did each model cost, comparatively? Are frontier models better, worse, or even with open-weight models? Are there tradeoffs to using one over the other? These are questions that every engineering team is wrestling with right now, so we also wondered: how can we help the community as they navigate these decisions?
We’ve been hearing Rails programmers say how great Rails and AI are together - that thanks to Convention over Configuration, Rails and AI pair well together, that Rails is token efficient, that models tend to write Rails well, the list goes on. From their own experience using agents within their own teams, the Rails Foundation board confirmed these sentiments.
We loved to hear it, but we needed more than anecdotal evidence. So we commissioned Evil Martians to benchmark popular agents on real Rails tasks in real Rails applications , and turn those results into practical guidance that your teams can use when choosing a model (or models).
This benchmark is being built out in stages, with Stage 1 launching today.
Stage 1: Atomic tasks. The current set of evals is made up of small, self-contained tasks that isolate one specific capability, and if they touch an API, it’s only the one. In addition to testing accuracy, speed, token spend, and cost across 8 different frontier and open-weight models, we’re also checking if the models reach for current Rails APIs.
The full methodology can be found here .
This benchmark will run as promising new models are released, so the Ruby community always has the most up to date information on how the models stack up against each other, and any cost/capability tradeoffs to consider.
The current leaderboard can be found on the new AI page , and each time we run the benchmark, a full findings report will accompany it and linked from that page. We’ve also created a new tag for these reports so you can always find them in the blog: Agents .
Stage 2: More realistic work. (Coming soon.) This is where we’ll move beyond isolated tasks into more realistic, complex work - testing how agents perform across longer, multi-step work, adding features, and building an app from scratch. These tests will be more challenging for the agents and will more closely mimic how you and your teams work on your apps in production.
Open-sourcing it all. The tasks and methodology are available now, and the raw runs will be uploaded into the repo tommorrow. Soon lemans , the Ruby harness that Evil Martians built to run all of this, will also be open sourced for the community.
A huge thank you to the Evil Martians team. They designed and built the benchmark harness and corpus, ran the evals, and turned the results into the first findings that we are publishing today. Thank you Svyatoslav Kryukov , Artur Petrov , Vladimir Dementyev , Albert Pazderin , Alexander Baygeldin , Anton Senkovskiy , and Irina Nazarova , who enthusiastically committed her team to the project before remembering that she should speak to them first. Luckily, the team was just as enthusiastic about the project. We love to see it.
We also want to thank Rob Zolkos for his work designing the new AI page and his help visualizing the report findings.
If you’re a Rails team with thoughts on what we should test next, or if you simply want to let us know what you think, get in touch: foundation@rubyonrails.org .
This work is made possible by the ongoing support and input from the Rails Foundation Core and Contributing members . Four years ago, when the Rails Foundation launched, we had a few main pillars in mind to focus on - marketing, documentation, education, events - all working towards the mission of supporting and growing the Rails ecosystem.
Since the rise of agent-assisted coding, things are changing for you, and that means it’s changing for us. The mission will remain the same, but the work is evolving. The reference library was one step in that direction. This is another such initiative. More will follow, so stay tuned.
