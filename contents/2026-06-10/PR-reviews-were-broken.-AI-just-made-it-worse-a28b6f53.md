---
source: "https://leaddev.com/ai/pr-reviews-were-already-broken-ai-made-it-worse"
hn_url: "https://news.ycombinator.com/item?id=48481306"
title: "PR reviews were broken. AI just made it worse"
article_title: "PR reviews were already broken. AI made it worse - LeadDev"
author: "argoeris"
captured_at: "2026-06-10T21:45:54Z"
capture_tool: "hn-digest"
hn_id: 48481306
score: 2
comments: 0
posted_at: "2026-06-10T19:22:01Z"
tags:
  - hacker-news
  - translated
---

# PR reviews were broken. AI just made it worse

- HN: [48481306](https://news.ycombinator.com/item?id=48481306)
- Source: [leaddev.com](https://leaddev.com/ai/pr-reviews-were-already-broken-ai-made-it-worse)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T19:22:01Z

## Translation

タイトル: PR レビューが壊れました。 AIが事態をさらに悪化させた
記事タイトル: PRレビューはすでに壊れていた。 AI が事態を悪化させた - LeadDev
説明: AI はコード レビューを中断しませんでした。中断されたプロセスを無視できなくなっただけであり、それを修正するには単一のソリューションではなく、複数の層が必要です。

記事本文:
PRレビューはすでに壊れていました。 AI が事態を悪化させた - LeadDev
コンテンツにスキップ
LeadDev.com を検索
行く
今後のイベント
ログイン ログイン
ニュースレター
受信トレイに最新ニュースが届きます
自分の役割に固有のコンテンツを見つける
PRレビューはすでに壊れていました。 AIが事態をさらに悪化させた
無料の LeadDev.com アカウントを登録する必要がある前に、今月読む記事が 1 つ残っています。
推定読了時間: 6 分
PR レビューの問題は AI によって引き起こされたのではなく、AI によって持続不可能になったのです。
本当の問題はデータです。不完全で相関性の低いデータを扱うエージェントは、根本原因ではなく症状を解決する PR を生成します。
解決策は単一の解決策ではなく、レイヤーです。
プル リクエスト (PR) レビューは、AI コーディング エージェントが登場する前から、ソフトウェア開発における弱点としてすでに知られていました。 AI ツールが問題を引き起こしたわけではありませんが、無視できなくなりました。
人間が人間の速度でコードを書いたとしても、モデルは苦労しました。 PR は何日もレビューされずに放置されていました。査読者は、自分の作業に戻る前に、500 行の差分をざっと確認しました。誰も完全に理解していない変更についてはゴム印による承認が行われていました。どのエンジニアリング組織にもこれの何らかのバージョンがあり、迅速に移行するためのコストとしてそれを最も静かに受け入れていました。
その後、AI コーディング ツールが現れて火に油を注ぎました。
エンジニアリングに関する洞察を毎週受け取り、リーダーシップのアプローチをレベルアップします。
コードをレビューすることは、コードを書くことよりも常に困難でした
PR レビューの根本的な問題は、AI とは関係のない文脈の非対称性です。
コードを作成するときは、検討したトレードオフ、拒否したアプローチ、特定のソリューションがコードベースにとって合理的である理由など、すべてを持ち歩きます。他の人のコードをレビューするときは、そのすべてを差分だけから再構築することになります。
著者が毎日一緒に仕事をしている同僚で、答えられる人であっても、これは難しいことです。

Slack で質問し、その理由を PR の説明で説明します。作成者が何百もの非決定的な決定を下したエージェントであり、それを把握できない場合、さらに困難になります。
人間による PR レビューを困難にしていたコンテキストのギャップは、AI が生成したコードでは桁違いに広がります。さらに、量の問題がそれをさらに悪化させます。
AI コーディング エージェントを使用しているチームは現在、大量の PR と平均品質の低下に対処しており、この組み合わせにより、すでに負荷がかかっていたレビュー ワークフローが破壊されます。
GitHub の Octoverse 2025 レポートには、オープンソースのメンテナーが「AI のスロップ」と呼んでいるものについて文書化されています。これらは大量かつ低品質で、多くの場合不正確な投稿であり、相応の価値を付加することなく査読者の注意を消耗します。この現象は、あらゆる規模のエンジニアリング組織に現れています。
根底にあるメカニズムが重要です。ほとんどの AI バグ報告およびデータ収集ツールは、最初からエージェント向けに設計されたものではありません。これらは、人間が調査できるように問題を表面化するために構築されました。 AI レイヤーが最上位にボルトで固定されても、基礎となるデータ インフラストラクチャは変わりません。エージェントは、不完全で相関性が低く、グループ化されていないデータに基づいて意思決定を行い、一見もっともらしく見えても実際の根本原因を見逃している修正を作成しています。
その結果、継続的インテグレーション (CI) を通過し、構文的に正しく見え、障害ではなく症状に対処する 400 行の PR が作成されます。これを発見した人間のレビュー担当者は、エージェントがそもそもアクセスできなかった因果関係の連鎖を再構築する必要があります。
有名なエラー監視ツールの CEO は、このパターンを直接認めました。低品質のデータ入力を扱うエージェントは、修正するのにさらに手間がかかる低品質の PR を生成します。
つまり、問題は、AI エージェントが次のことを求められているということです。

機械推論用に設計されていないデータを使用したシステムに関する決定。
Coinbase、AI主導の再構築で経営を平坦化し、従業員を削減
AI を使用すると、レビューできないほど多くのコードが作成される
修正を考えるためのフレームワーク
複雑なシステムが 1 つの原因で失敗することはほとんどありませんし、1 つの原因で修正されることもほとんどありません。
スイス チーズ モデルは、ソフトウェアの品質検証を考えるための最も誠実なフレームワークです。いくつかのスライスが積み重ねられ、それぞれがプロセス、チェック、セーフガードなどの防御層を表すと想像してください。各スライスには穴がありますが、穴がすべてのスライスにわたって同時に整列しない限り、障害は発生しません。不完全な層を十分に積み重ねると、個々のコンポーネントよりも信頼性の高いシステムが得られます。
PR審査危機はスイスチーズの問題だ。単一の介入では問題を解決できません。浮かび上がってきたのは、それぞれが特定の障害モードに対処し、単独で行うよりも連携して効果的に機能する一連の部分的な答えです。
PR レビューの新たな層
これにより、意図のずれによる失敗モードに対処できます。エージェントが 1 行のコードを記述する前に、詳細で合理的な仕様に基づいて作業を行う場合、出力はより予測可能になり、レビュー担当者はコードをチェックするための具体的なものを得ることができます。
これにより、単一エージェントの盲点による障害モードに対処できます。同じタスクを複数のエージェントに同時に割り当て、分岐するソリューションを生成させ、どれが最も多くの検証ステップを通過するか、最小の差分を導入するか、または新しい依存関係を回避するかに基づいて選択します。競争は、一度の試みでは得られないシグナルを生み出します。
エージェント向けのより優れたデータ インフラストラクチャ
これにより、PR スロップの根本にある障害モードに対処できます。エージェントが高品質の入力を使用して推論する場合（アンサンプル）

d、フルスタック、事前相関、重複排除されたコンテキスト）結果の品質が大幅に変わります。実稼働障害の適切に構造化された因果関係追跡に基づいて動作するエージェントは、生のログ ストリームに基づいて動作するエージェントとは異なる動作をします。
拡張されたテスト カバレッジ、CI チェック、契約検証、静的分析などのレイヤーは、変更のたびに人間が関与する必要がなく、レビューで捕らえられていたものをより早く、より迅速に捕らえます。
これらのアプローチの驚くべき点は、どのアプローチも相互に排他的ではないということです。仕様主導の開発を行っているチームは、より優れたエージェント データから引き続き恩恵を受けます。強力な自動検証を備えたチームは、複雑な問題に対して複数のエージェントによる競争を実行できます。これらはレイヤーであり、代替品ではありません。
スイス チーズ モデルの各層は、査読者の考えを覆すものです。
仕様主導型開発は、コード行が存在する前に意図を検証します。
エージェントのデータが改善されると、レビュー対象となる低品質の PR が減ります。
自動検証は、人間が目にする前に機械的エラーを検出します。
残るのは、レビュー担当者が実際に行う資格のある意図とアーキテクチャに関する判断要求を要求する PR です (そもそもレビュー担当者に届くはずのなかった差分レビューではありません)。
ニューヨーク • 2026 年 9 月 15 ～ 16 日
では、PRレビューは絶滅してしまうのでしょうか？現在、ほとんどのチームが PR レビューを実践している (つまり、人間が diff を 1 行ずつ読んでマージを承認する) ことは、ほぼ間違いなくそうなります。 AI によって生成されたコードの量により、そのモデルは経済的にも認知的にも持続不可能になります。既存のプロセスにツールをいくら取り付けても、その基本的な演算は変わりません。
これに代わるものは、人間の判断が実装ではなく意図とアーキテクチャに留保される階層化された検証システムです。
何が成功かを定義するのは人間です

ess は、ビジネスコンテキストと二次的な結果を真に理解する必要があるように見え、制約を設定し、呼び出しを行います。マシンは、これらの制約が満たされていることの検証を処理します。
これは、チームが AI ツールをワークフローに組み込むことを評価する方法に対する重要な変更です。 「この新しいツールを採用し、他のすべてを同じままにする」だけでは十分ではありません。必要なのは、人間の注意がどこに向かうのか、何のためにあるのかを真に再編成することです。
自己修復ソフトウェアの目標は何年もの間語られてきました。現在の違いは、各要素が最終的に正しい順序で組み立てられつつあることです。エージェントが意思決定を行うためのより良いデータ、すべての行で人間の目を必要としない検証レイヤー、そしてレビュープロセス自体を再設計する必要があるという認識が高まっています。
ソフトウェアの品質を検証する方法は、ソフトウェアを生成する方法と同じペースで進化する必要があります。壊れたレビュープロセスを維持し、それをより良いツールで解決することを期待するのは戦略ではありません。
このような記事を受信トレイに受け取る
購読する LeadDev ニュースレターを選択してください。
トーマス・ジョンソン
Thomas は Multiplayer の共同創設者兼 CTO です。
2026 年の最高の AI コーディング ツール
すべての AI コーディング ツールが同じというわけではありません。
2026 年の AI コーディング ツール購入チェックリスト
実用的なチェックリストを使用して、今日の AI コーディング ツールを評価します。
可観測性ツールは AI デバッグ用に構築されていない
AI のデバッグにはデータの問題があります。
LLM がウォルマートのオンコール エンジニアになった経緯
LLM は単なる生産性向上ツールではありません。
スポンサーシップと広告の機会
講演、記事、またはチャレンジに貢献する

## Original Extract

AI didn’t break code review – it just made a broken process impossible to ignore, and fixing it requires layers, not a single solution.

PR reviews were already broken. AI made it worse - LeadDev
Skip to content
Search LeadDev.com
Go
Upcoming events
Login Login
Newsletters
Latest news in your inbox
Find content specific to your role
PR reviews were already broken. AI made it worse
You have 1 article left to read this month before you need to register a free LeadDev.com account.
Estimated reading time: 6 minutes
AI didn’t create the PR review problem – it made it unsustainable .
The real issue is data: agents working on incomplete, poorly correlated data produce PRs that fix symptoms, not root causes.
The fix is layers , not a single solution.
Pull request (PR) reviews were already a known weak point in software development before AI-coding agents arrived. While AI tools didn’t create the problem, they made it impossible to ignore.
Even when humans wrote code at human speed, the model struggled. PRs sat unreviewed for days. Reviewers skimmed 500-line diffs before returning to their own work. There were rubber-stamp approvals on changes nobody fully understood. Every engineering org had some version of this, and most quietly accepted it as the cost of moving fast.
Then AI-coding tools showed up and poured fuel on the fire.
Receive weekly engineering insights to level up your leadership approach.
Reviewing code has always been harder than writing it
The root problem with PR reviews is a context asymmetry that has nothing to do with AI.
When you write code, you carry everything with you: the tradeoffs you considered, the approaches you rejected, and the reason a particular solution made sense for a codebase. When you review someone else’s code, you’re reconstructing all of that from the diff alone.
That’s hard even when the author is a colleague you work with every day, who can answer questions in Slack and explain their reasoning in the PR description. It’s harder still when the author is an agent that made hundreds of non-deterministic decisions that you have no visibility into.
The context gap that made human PR reviews difficult is an order of magnitude wider with AI-generated code. What’s more, the volume problem compounds it.
Teams using AI-coding agents are now contending with a larger volume of PRs and a lower average quality, a combination that breaks review workflows that were already under strain.
GitHub’s Octoverse 2025 report documents what open source maintainers are calling “AI slop.” These are high-volume, low-quality, and often inaccurate contributions that consume reviewer attention without adding proportionate value. This phenomenon is showing up across engineering organizations of every size.
The underlying mechanism matters. Most AI bug-reporting and data-gathering tools weren’t designed for agents from the ground up. They were built to surface problems for humans to investigate. When an AI layer was bolted on top, the underlying data infrastructure didn’t change. Agents are making decisions on incomplete, poorly correlated, ungrouped data and producing fixes that are plausible-looking but miss the actual root cause.
The result is a 400-line PR that passes continuous integration (CI), looks syntactically correct, and addresses the symptom rather than the failure. A human reviewer catching this has to reconstruct the causal chain the agent never had access to in the first place.
The CEO of a well-known error monitoring tool acknowledged the pattern directly: agents working with low-quality data inputs produce low-quality PRs that are more work to fix.
In short, the problem is that AI agents are being asked to make decisions about systems using data that was never designed for machine reasoning.
Coinbase flattens management and trims workforce in AI-driven restructure
Use of AI has us creating more code than we can review
A framework for thinking about the fix
Complex systems rarely fail because of a single thing, and they rarely get fixed by one either.
The Swiss cheese model is the most honest framework for thinking about software quality verification. Imagine several slices stacked on top of each other, each representing a defensive layer: a process, a check, a safeguard. Each slice has holes, but as long as the holes don’t align across all slices simultaneously, failures don’t get through. Stack enough imperfect layers and you get a system more reliable than any individual component.
The PR review crisis is a Swiss cheese problem. No single intervention fixes it. What’s emerging is a set of partial answers that each address a specific failure mode and work better together than any of them do alone.
The emerging layers of PR reviews
This addresses the failure mode of misaligned intent. If an agent works from a detailed, well-reasoned specification before writing a single line of code, the output is more predictable and the reviewer has something concrete to check the code against.
This addresses the failure mode of a single agent’s blind spots. Assign the same task to multiple agents simultaneously, let them produce diverging solutions, and select based on which passes the most verification steps, introduces the smallest diff, or avoids new dependencies. Competition creates a signal you wouldn’t get from a single attempt.
Better data infrastructure for agents
This addresses the failure mode at the root of PR slop. When agents reason with high-quality inputs (unsampled, full-stack, pre-correlated, deduplicated context) the quality of what comes out changes significantly. An agent working from a properly structured causal trace of a production failure behaves differently than one working from a raw log stream.
Layers such as expanded test coverage, CI checks, contract verification, and static analysis catch what reviews used to catch, earlier and faster, without requiring a human in the loop for every change.
What’s striking about these approaches is that none of them are mutually exclusive. A team doing spec-driven development still benefits from better agent data. A team with strong automated verification could still run multi-agent competition on complex problems. These are layers, not alternatives.
Each layer in the Swiss cheese model takes something off the reviewer’s plate :
Spec-driven development validates intent before a line of code exists.
Better agent data means fewer low-quality PRs making it to review.
Automated verification catches mechanical errors before a human sees them.
What remains are PRs requiring the judgment calls on intent and architecture that reviewers are actually qualified to make (not diffs reviews that should never have reached them in the first place).
New York • September 15-16, 2026
So, will PR reviews become extinct? The PR review as most teams practice it today (i.e. a human reading a diff line by line and approving a merge) almost certainly will. The volume of AI-generated code makes that model economically and cognitively unsustainable. No amount of tooling bolted onto the existing process will change that fundamental arithmetic.
What replaces it is a layered verification system where human judgment is reserved for intent and architecture rather than implementation.
Humans define what success looks like, set the constraints, and make the calls that require genuine understanding of business context and second-order consequences. Machines handle the verification that those constraints were met.
This is a significant change to how teams need to evaluate incorporating AI tools into their workflows. It’s not sufficient to “adopt this new tool and keep everything else the same.” What’s needed is a genuine reorganization of where human attention goes and what it’s for.
The goal of self-healing software has been talked about for years. What’s different now is that the pieces are finally being assembled in the right order: better data for agents to make decisions on, verification layers that don’t require human eyes on every line, and a growing recognition that the review process itself needs to be redesigned.
The way we verify software quality needs to evolve at the same pace as the way we generate it. Keeping a broken review process and hoping better tools paper over it isn’t a strategy.
Get articles like this in your inbox
Choose your LeadDev newsletters to subscribe to.
Thomas Johnson
Thomas is co-founder and CTO at Multiplayer.
The best AI-coding tools in 2026
Not all AI-coding tools are equal.
Your AI-coding tools buying checklist for 2026
Evaluate today’s AI-coding tools with a practical checklist.
Observability tools weren’t built for AI debugging
AI debugging has a data problem.
How LLMs became Walmart’s on-call engineer
LLMs aren’t just productivity tools.
Sponsorship & advertising opportunities
Contribute a talk, article, or challenge
