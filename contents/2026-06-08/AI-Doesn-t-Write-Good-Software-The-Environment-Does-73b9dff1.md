---
source: "https://adrianferrera.dev/en/blog/ai-does-not-write-good-code"
hn_url: "https://news.ycombinator.com/item?id=48444459"
title: "AI Doesn't Write Good Software: The Environment Does"
article_title: "AI Doesn't Write Good Software: The Environment Does · Adrian Ferrera"
author: "RyeCombinator"
captured_at: "2026-06-08T13:35:04Z"
capture_tool: "hn-digest"
hn_id: 48444459
score: 2
comments: 0
posted_at: "2026-06-08T12:26:05Z"
tags:
  - hacker-news
  - translated
---

# AI Doesn't Write Good Software: The Environment Does

- HN: [48444459](https://news.ycombinator.com/item?id=48444459)
- Source: [adrianferrera.dev](https://adrianferrera.dev/en/blog/ai-does-not-write-good-code)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T12:26:05Z

## Translation

タイトル: AI は良いソフトウェアを作成しない: 環境が良いソフトウェアを作成する
記事のタイトル: AI は良いソフトウェアを書かない: 環境がそうする · Adrian Ferrera
説明: 生成 AI はコード生成コストを大幅に削減しましたが、まだ

記事本文:
AI は良いソフトウェアを書かない: 環境がそうする · Adrian Ferrera コンテンツへスキップ adrianferrera.dev フィード ブログ ポッドキャスト リソース 英語 / スペイン語 ◐ テーマ: ダーク フィード ブログ ポッドキャスト リソース About ← ブログ投稿に戻る AI は良いソフトウェアを書かない: 環境がそうする
私たちはより良いソフトウェアを構築しているのでしょうか、それとも単により多くのソフトウェアを構築しているのでしょうか?
生成 AI により、コード生成コストが大幅に削減されました。数分で、機能の実装、サービスの作成、テストの生成、モジュールのリファクタリング、アーキテクチャ全体のスキャフォールディングを依頼できます。
しかし、業界として私たちが自問すべき不快な質問があります。それは、より優れたソフトウェアを構築しているのか、それとも単により多くのソフトウェアを構築しているのかということです。
なぜなら、以前はボトルネックがコードを書くことであったとしても、今ではそれはまったく別のことであり、それを理解し、検証し、保守し、構築しているシステムに実際に適合するかどうかを知ることがボトルネックになっているからです。
ソフトウェア業界は何年もの間、同じ間違いを繰り返してきました。理解する前に急いで、検証する前に構築し、テストを後回しにし、技術的負債を避けられないかのように受け入れ、速度と進歩を混同し、偶発的なアーキテクチャを設計し、レビューが遅すぎました。
違いは、AI によって私たちが同じ間違いを驚くべきスピードで犯せるようになったということです。
AI が登場する以前から、業界はすでに深刻な問題を抱えていました。 CISQ は、2022 年の米国における低品質なソフトウェアのコストは約 2 兆 4,100 億ドル、蓄積された技術的負債は約 1 兆 5,200 億ドルになると推定しています。 AI は、健全で完全に管理された業界には登場しません。それは、すでに構造的な問題を引きずっていたものに到着します。
重要なのは、AI がまったく新しい問題を引き起こすということではありません。それはすでに存在している問題を加速させるということです。
AIは新しい技術を生み出さない

借金です。これにより、私たちがすでに作成方法を知っていた技術的負債が加速します。
注目に値する概念があります。それは検証負債です。 AI は正しく見えるコードを生成できますが、そのコードは検証する必要があります。チームが変更を適切にレビューせずに受け入れた場合、責任はコード内だけではなく、検証プロセス自体にもあります。
コードの作成コストは削減しましたが、コードの理解、検証、デプロイ、運用、保守のコストは必ずしも削減できませんでした。新たなボトルネックは常にコードの作成にあるわけではありません。そのコードが正しいかどうかを知ることになります。
そしておそらく本当のリスクは、AI が開発者に取って代わることではなく、AI によって私たちの判断力の欠如が自動化されることです。
問題を整理すると、論文は簡単になります。
AI は、それ自体で優れたソフトウェアを作成するわけではありません。これは、明確な境界、信頼できるフィードバック、人間の介入のためのスペースを提供する環境内で動作する場合に行われます。
品質はモデルだけに依存するわけではありません。それは、統合するワークフローによって異なります。これはAIに反対する話ではありません。これは、その単純な、制御されていない、または純粋に速度重視の使用に反対します。
AI が私たちの悪い習慣を加速させるのであれば、コード生成前、生成中、生成後に判断を導入する方法が必要です。ソリューションは特定のツールや閉じたフレームワークではなく、ワークフローです。
発見 → 計画 → レビュー → 実装 → 検証
エージェントのためのスクラムではありません。これは最小限の制御構造であり、あらゆるチームに適応でき、AI に推論してもらいたいとき、提案してもらいたいとき、実行すべきとき、AI が行ったことを検証する必要があるときを区別するように設計されています。
発見: 私たちは問題を正しく理解していますか?
実装する前に、システムは問題について検討する必要があります。このフェーズでは、

コード。私たちは、AI が問題を理解し、代替案を検討し、リスクを特定し、質問し、仮定を表面化し、曖昧さを検出できるようにしたいと考えています。
人間は、ビジネス コンテキスト、現実の制約、過去の決定、製品知識、リスクへの敏感度など、AI が持たないものを確実に提供します。
ガイドとなる質問: コードに触れる前に問題をよく理解できていますか?
計画: 何を変更するのか、そしてその理由はわかっていますか?
推論を具体的な一連のステップに変換します。計画では、何が変更されるか、どのファイルまたはモジュールが影響を受ける可能性があるか、何が変更されないか、どのようなテストが必要か、どのようなリスクを制御する必要があるかを明確にする必要があります。
エージェントが何かを変更する前に、人間は計画を承認、修正、または拒否できます。
指針となる質問は、この変更は実行する前に意味があるのでしょうか?
復習: ソリューションを構築する前に、そのソリューションは意味をなしていますか?
実装する前に確認してください。これはコードをレビューすることではありません。意図、アプローチ、トレードオフ、システムへの影響、アーキテクチャの一貫性、ソリューションの保守性をレビューすることです。
この段階では、もっともらしいが方向性が不十分な解決策を受け入れることができなくなります。また、コードがすでに書かれており、感情的にそれを破棄するのが難しい場合に、レビューが手遅れになることを防ぎます。
指針となる質問: ソリューションの構築コストを支払う前に、そのソリューションに同意しますか?
実行する: 計画を実行することですか、それとも即興で行うことですか?
AI は変更を実装しますが、もはや盲目的ではありません。これは、発見、計画、レビューを経た後に行われます。コード生成マシンであることをやめ、意思決定フレームワーク内の実行者になります。
人間は、AI が計画に従い、即興で範囲外の変更を加えず、下された決定を尊重し、不必要なものを導入しないことを監督します。

複雑。
指針となる質問は、計画を実行するのか、それとも即興で行うのか、ということです。
検証: これが正しいというシグナルは何ですか?
変更が正しいことを検証します。検証とは、コードがコンパイルされるかどうかを確認することだけを意味するわけではありません。これは、自動テスト、型指定、リンター、静的分析、継続的インテグレーション、差分レビュー、手動テスト、機能検証、アーキテクチャ ルール、契約、可観測性、セキュリティ、更新されたドキュメントなど、変更を受け入れることができるかどうかを決定するのに十分なシグナルを収集することを意味します。
信号が十分かどうか、リスクが許容できるかどうかは人間が判断します。
指針となる質問: これが正しいというシグナルは何ですか?
Human in the Loop: 適切なタイミングで意思決定を行う
中心的な考え方の 1 つは、Human in the Loop は単に AI が最終的に何をしたかを観察することを意味するわけではないということです。それは、構築する前、ソリューションを受け入れる前、変更を統合する前、技術的なリスクを負う前など、最も影響力のある時点で人間の判断を導入することを意味します。
人間は感情的なコンパイラーのようにすべての行をレビューするためにそこにいるわけではありません。彼らは、意図、コンテキスト、判断、経験、責任、将来の影響に対する感受性など、AI が持たないものを確実に提供するために存在します。
Human in the Loop は最後にレビューをしません。適切なタイミングで決断するのです。
オーケストレーター、ハーネス、そして人間
このワークフローはオーケストレーターの基礎となります。オーケストレーターがすべてを自動化するからではなく、責任を分離するためです。いつ推論するか、いつ決定するか、いつ実行するか、いつ検証するか、いつ人間の介入を要求するかです。
ハーネスは、AI が境界、信号、フィードバックを使用して動作できるようにするメカニズムです。テスト、タイプ、リンター、静的分析、アーキテクチャ ルール、チーム コンサルテーションなどです。

ベンション、生きたドキュメント、CI パイプライン、セキュリティ レビュー、可観測性、人間による検証、API コントラクト、メトリクス、機能フラグ、ステージング環境。
ハーネスのないエージェントは、方向性のないスピードを発揮します。
オーケストレーターは構造を提供します。
ハーネスはフィードバックを提供します。
3つの要素はすべて必要です。他のものに代わるものはありません。
ワークフローは各チームと状況に適応する必要があります。小さなバグの場合、発見には 30 秒かかる場合があり、計画は非常に単純で、検証は 1 回のテストと簡単なレビューに限定される場合があります。アーキテクチャ変更の場合、検出には完全なセッションが必要になる場合があり、計画には複数の手順が必要になる場合があり、レビューには複数の人が関与する場合があり、検証にはテスト、正式なレビュー、セキュリティ、パフォーマンス、可観測性が含まれる場合があります。
重要なのは、各フェーズに常に同じ重みを適用することではなく、各変更にどの程度の制御が必要かを意識的に決定することです。
フローの利点は、プロセスを強制しないことです。それは、どの時点で AI に行動をさせるのか、そしてどのような信号で AI の行動が受け入れられると判断するのかという疑問を課すことになります。
私たちは AI が単に高速化することを望んでいません。私たちはそれが正しい方向に早く進むことを望んでいます。
目標は、AI にさらに多くのことをしてもらうことではありません。目標は、コードを要求するだけでなく、推論を要求するなど、コラボレーションの方法を改善することです。計画のない変化を受け入れない。意図を検討せずに実装しない。信号がなければ積分しない。判断を委任するものではありません。
AI は私たちを技術的責任から解放しません。それがより重要になります。
本当の課題は、AI を使用するかどうかを決めることではありません。多くのチームでは、そのような会話はすでに終わっています。課題は、それをどのように使用するか、どのような境界線で、どのような品質の信号を使用し、どのようなスペースで人間が介入するかを決定することです。

ハーネスを定義しましょう。ルールを明確にしましょう。自分たちの基準を守りましょう。受け入れる前に確認しましょう。統合する前に確認しましょう。そして何よりも技術的な判断を中心に据えましょう。
なぜなら、ソフトウェア開発の未来は、ツールの使い方を熟知している人だけのものではなくなるからです。それは、それを責任を持って使用するためのより良いシステムを作成する方法を知っている人のものになります。
AI は、それ自体で優れたソフトウェアを作成するわけではありません。環境はそうなります。そしてその環境は私たちがデザインします。
この記事は、CommitConf 2026 で私が発表した講演「AI は良いソフトウェアを書かない: 環境がそうする」に基づいています。
シニア ソフトウェア開発者は、チームが保守可能なソフトウェアを構築し、技術的な判断に基づいて AI を適用できるように支援します。
受信箱で次のメモを受け取ります。
ソフトウェア エンジニアリング、アーキテクチャ、テスト、製品、応用 AI に関する新しい執筆と不定期の更新情報を購読してください。重いファネルはなく、シンプルなメーリング リストだけです。
MailerLite によって管理されます。 MailerLite が送信を処理します。
私たちはより良いソフトウェアを構築しているのでしょうか、それとも単により多くのソフトウェアを構築しているのでしょうか?
Human in the Loop: 適切なタイミングで意思決定を行う
オーケストレーター、ハーネス、そして人間
インテリジェンスの時代における製品の構築と、開発者の技術を生かし続けることについて書いています。
投稿 ペアプログラミングとは何ですか?

## Original Extract

Generative AI has drastically reduced the cost of producing code, but it hasn

AI Doesn't Write Good Software: The Environment Does · Adrian Ferrera Skip to content adrianferrera.dev Feed Blog Podcast Resources About English / Español ◐ Theme: Dark Feed Blog Podcast Resources About ← Back to blog Post AI Doesn't Write Good Software: The Environment Does
Are we building better software, or just more software?
Generative AI has dramatically reduced the cost of producing code. We can ask it to implement a feature, write a service, generate tests, refactor a module, or scaffold an entire architecture in minutes.
But there’s an uncomfortable question we should be asking ourselves as an industry: are we building better software, or just more software?
Because if the bottleneck used to be writing code, now it’s something else entirely: understanding it, validating it, maintaining it, and knowing whether it actually fits the system we’re building.
For years, the software industry has been repeating the same mistakes: rushing before understanding, building before validating, leaving tests for later, accepting technical debt as if it were inevitable, confusing speed with progress, designing accidental architectures, and reviewing too late.
The difference is that AI now lets us make those same mistakes at staggering speed.
Before AI, the industry already had a serious problem. CISQ estimated that in 2022, the cost of poor software quality in the United States was approximately $2.41 trillion , and that accumulated technical debt was around $1.52 trillion . AI doesn’t arrive in a healthy, perfectly controlled industry. It arrives in one that was already dragging structural problems.
The point isn’t that AI creates a completely new problem. It’s that it accelerates problems that already existed .
AI doesn’t create new technical debt. It accelerates the technical debt we already knew how to create.
There’s a concept that deserves attention: verification debt . AI can generate code that looks correct, but that code must be verified. If teams accept changes without reviewing them properly, the debt isn’t just in the code — it’s in the validation process itself.
We’ve reduced the cost of producing code, but not necessarily the cost of understanding it, validating it, deploying it, operating it, and maintaining it. The new bottleneck won’t always be writing code. It will be knowing whether that code is correct.
And perhaps the real risk isn’t that AI replaces developers, but that it allows us to automate our lack of judgment .
After laying out the problem, the thesis is straightforward:
AI doesn’t write good software on its own. It does so when it works within an environment that provides clear boundaries, reliable feedback, and spaces for human intervention.
Quality doesn’t depend solely on the model. It depends on the workflow where we integrate it. This isn’t a talk against AI. It’s against its naive, uncontrolled, or purely speed-driven use.
If AI accelerates our bad practices, we need a way to introduce judgment before, during, and after code generation. The solution isn’t a specific tool or a closed framework — it’s a workflow:
Discovery → Plan → Review → Implement → Verify
It’s not Scrum for agents. It’s a minimal control structure : adaptable to every team, designed to separate when we want AI to reason, when we want it to propose, when it should execute, and when we need to validate what it has done.
Discovery: are we understanding the problem correctly?
Before implementing, the system should think about the problem. In this phase, we don’t want code. We want AI to help understand the problem, explore alternatives, identify risks, ask questions, surface assumptions, and detect ambiguities.
The human contributes what AI doesn’t have reliably: business context, real constraints, historical decisions, product knowledge, and sensitivity to risk.
The guiding question: are we understanding the problem well before touching the code?
Plan: do we know what we’re going to change and why?
Turn the reasoning into a concrete sequence of steps. The plan should clarify what will change, which files or modules might be affected, what won’t be touched, what tests are needed, and what risks need to be controlled.
The human can accept, correct, or reject the plan before the agent modifies anything.
The guiding question: does this change make sense before we execute it?
Review: does the solution make sense before we build it?
Review before implementing. This isn’t about reviewing code — it’s about reviewing the intent, the approach, the trade-offs, the impact on the system, the architectural coherence, and the maintainability of the solution.
This phase prevents accepting plausible but poorly oriented solutions. It also prevents reviewing too late, when code has already been written and it’s emotionally harder to discard it.
The guiding question: do we agree with the solution before paying the cost of building it?
Implement: is it executing the plan or improvising?
AI implements the change, but no longer blindly. It does so after going through discovery, planning, and review. It stops being a code generation machine and becomes an executor within a decision framework.
The human oversees that AI follows the plan, doesn’t improvise out-of-scope changes, respects the decisions made, and doesn’t introduce unnecessary complexity.
The guiding question: is it executing the plan or improvising?
Verify: what signals do we have that this is correct?
Validate that the change is correct. Verify doesn’t just mean checking that the code compiles. It means gathering enough signals to decide whether the change can be accepted: automated tests, typing, linters, static analysis, continuous integration, diff review, manual testing, functional validation, architecture rules, contracts, observability, security, and updated documentation.
The human decides whether the signals are sufficient and whether the risk is acceptable.
The guiding question: what signals do we have that this is correct?
Human in the Loop: deciding at the right moment
One of the central ideas is that Human in the Loop doesn’t mean simply looking at what AI has done at the end. It means introducing human judgment at the points where it has the most impact: before building, before accepting a solution, before integrating a change, and before taking on technical risk.
The human isn’t there to review every line like an emotional compiler. They’re there to contribute what AI doesn’t have reliably: intention, context, judgment, experience, responsibility, and sensitivity to future impact.
Human in the Loop isn’t reviewing at the end. It’s deciding at the right moment.
The orchestrator, the harnesses, and the human
This workflow can be the foundation of an orchestrator. Not because the orchestrator automates everything, but because it separates responsibilities: when to reason, when to decide, when to execute, when to verify, and when to request human intervention.
Harnesses are mechanisms that allow AI to operate with boundaries, signals, and feedback: tests, types, linters, static analysis, architecture rules, team conventions, living documentation, CI pipelines, security reviews, observability, human validation, API contracts, metrics, feature flags, and staging environments.
An agent without harnesses is speed without direction.
The orchestrator provides structure.
The harnesses provide feedback.
All three elements are necessary. None replaces the others.
The workflow must adapt to each team and situation. For a small bug, discovery might last thirty seconds, the plan might be very simple, and verification might be limited to one test and a quick review. For an architecture change, discovery might take a full session, the plan might require several steps, the review might involve multiple people, and verification might include tests, formal review, security, performance, and observability.
The key isn’t to always apply the same weight to each phase, but to consciously decide how much control each change needs.
The beauty of the flow is that it doesn’t impose a process. It imposes a question: at what point do we let AI act, and with what signals do we decide that what it did is acceptable?
We don’t want AI to simply go faster. We want it to go faster in the right direction .
The goal isn’t to ask AI to do more things. The goal is to improve how we collaborate with it: not just asking for code, but for reasoning; not accepting changes without a plan; not implementing without reviewing intent; not integrating without signals; not delegating judgment.
AI doesn’t free us from technical responsibility. It makes it more important.
The real challenge isn’t deciding whether we use AI or not. That conversation, in many teams, is already behind us. The challenge is deciding how we use it, with what boundaries, with what quality signals, and with what spaces for human intervention.
Let’s define our harnesses. Let’s make our rules explicit. Let’s protect our standards. Let’s review before accepting. Let’s verify before integrating. And above all, let’s keep technical judgment at the center.
Because the future of software development won’t belong only to those who know how to use a tool better. It will belong to those who know how to create better systems for using it responsibly.
AI doesn’t write good software on its own. The environment does. And that environment is designed by us.
This article is based on the talk “AI Doesn’t Write Good Software: The Environment Does” that I presented at CommitConf 2026.
Senior Software Developer helping teams build maintainable software and apply AI with technical judgment.
Get the next note in your inbox.
Subscribe for new writing and occasional updates on software engineering, architecture, testing, product, and applied AI. No heavy funnel, just a simple mailing list.
Managed by MailerLite. MailerLite handles sending.
Are we building better software, or just more software?
Human in the Loop: deciding at the right moment
The orchestrator, the harnesses, and the human
I write about building products in the age of intelligence and keeping developer craft alive.
Post What is pair programming?
