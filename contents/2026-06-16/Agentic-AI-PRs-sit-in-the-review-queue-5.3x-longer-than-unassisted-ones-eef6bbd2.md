---
source: "https://blog.codacy.com/ai-breaking-code-review-how-engineering-teams-survive-pr-bottleneck"
hn_url: "https://news.ycombinator.com/item?id=48553126"
title: "Agentic AI PRs sit in the review queue 5.3x longer than unassisted ones"
article_title: "AI Is Breaking Code Review: How Engineering Teams Survive the PR Bottleneck"
author: "claudiacsf"
captured_at: "2026-06-16T10:40:40Z"
capture_tool: "hn-digest"
hn_id: 48553126
score: 1
comments: 1
posted_at: "2026-06-16T10:29:55Z"
tags:
  - hacker-news
  - translated
---

# Agentic AI PRs sit in the review queue 5.3x longer than unassisted ones

- HN: [48553126](https://news.ycombinator.com/item?id=48553126)
- Source: [blog.codacy.com](https://blog.codacy.com/ai-breaking-code-review-how-engineering-teams-survive-pr-bottleneck)
- Score: 1
- Comments: 1
- Posted: 2026-06-16T10:29:55Z

## Translation

タイトル: Agentic AI PR は、支援なしの PR よりも 5.3 倍長くレビュー キューに留まります
記事のタイトル: AI によるコードレビューの破壊: エンジニアリング チームが PR のボトルネックをどうやって乗り切るか
説明: AI によって生成されたコードがプル リクエストのレビューにどのような影響を及ぼし、ボトルネックを生み出し、チーム ダイナミクスを変化させているかを確認します。コードの品質と効率を維持する方法を学びます。

記事本文:
AI はコード レビューを打破する: エンジニアリング チームが PR のボトルネックを乗り切る方法
-->
AI コーディング ツールにより、コードの作成は容易になりましたが、コードを安全に出荷するのは簡単にはなりませんでした。
プル リクエストのキューは、レビューのキャパシティを上回る速さで増加しています。 CircleCI の 2026 年のデータによると、フィーチャー ブランチのスループットは前年比 59% 増加しましたが、中央値チームのメイン ブランチのスループットは実際には低下しました。ボトルネックは、コードの作成から、コードをマージしても安全かどうかの判断に移りました。この記事では、AI によって生成されたコードがレビューのプレッシャーを引き起こす理由、自動化ツールが対応できるもの、エンジニアリング チームが基準を下げることなく PR を継続できる方法について説明します。
AI 生成コードがレビューのボトルネックを引き起こす理由
AI 生成コードのレビューが他と異なる点
プル リクエストを遅らせることなく AI が生成したコードをレビューする方法
自動化ツールがレビュープロセスで所有できるもの
人間の査読者が依然として評価する必要があるもの
一般的な AI レビュー担当者が重要な問題を見逃す理由
コンプライアンス要件がレビュー モデルをどのように形成するか
最適化された PR レビューの上限
AI 生成コードがレビューのボトルネックを引き起こす理由
プル リクエストの速度を低下させることなく AI 生成のコードをレビューするには、ベースライン チェックを人間の目から遠ざけ、人間によるレビューを意図とアーキテクチャの適合性に焦点を当てる必要があります。 AI を起動したエンジニアが依然として出力を所有しています。しかし、レビュー プロセス自体は変化する可能性があります。自動チェックはフォーマット、セキュリティ パターン、既知の脆弱性を処理しますが、人間はその変更が実際に正しい問題を解決するかどうかに集中します。
主な緊張関係は次のとおりです。AI 支援開発により、コード生成とレビュー能力の比率が変化しました。エンジニアはより短い時間でより多くのコードを生成できますが、そのコードを検証するチームの能力は同じ割合で拡張されていません。
CircleCI の 202

6 State of Software Delivery レポートでは、22,000 以上の組織で実行されている 2,800 万を超える CI ワークフローを分析しました。前述したように、全体のスループットは前年比 59% 増加しました。同時に、機能ブランチのスループットは中央値チームで 15% 増加しましたが、メインブランチのスループットは 7% 近く低下し、メインブランチの成功率は 70.8% に低下しました。
パイプラインに入るコードは増えていますが、本番環境に正常に到達しているコードは少なくなります。ボトルネックは、コードの作成から、コードをマージしても安全かどうかの判断に移りました。
AI 生成コードのレビューが他と異なる点
コードレビューには常に隠れた運用コストが伴います。構築とレビューの間でコンテキストを切り替えると、両方のアクティビティが遅くなります。フィードバック ループでは議論が起こる可能性があります。そして、PR キューが増大すると、チームは仕事を進めるために厳密さを緩めることがよくあります。Faros AI データによると、レビューなしでマージされる PR が 31% 増加していることが示されています。表面的な承認やエッジケース分析のスキップにつながります。
AI によって生成されたコードは、特定の方法でレビューのプレッシャーを増幅させます。
まず、音量の問題があります。 AI 支援ワークフローは、より多くのブランチ、より多くのコミット、より多くの PR を生成します。コーディング アシスタントと協力して作業する 1 人のエンジニアは、以前は 1 つを完了するのにかかっていた時間内に複数の PR を開くことができます。
第二に、コンテキストの問題があります。 AI エージェントがコードを生成する場合、レビュー担当者は多くの場合、同じ実装行程や意思決定証跡のない完成した差分を受け取ります。レビュー担当者は、チケット、PR の説明、コードの変更だけから意図を再構築する必要があります。 LinearB の 2026 ソフトウェア エンジニアリング ベンチマーク レポートによると、エージェント付き AI PR は非アシスト型 PR よりもピックアップ時間が 5.3 倍長いことがわかりました。 AI 支援 PR は 2.47 倍長く待機します。ピックアップ時間が長いということは、レビュー担当者が AI によって生成された変更の評価により多くの時間を費やしていることを示唆しており、より深いレビューに貢献しています

行列。
第三に、信頼の問題があります。 AI によって生成されたコードは、気軽に読むだけで十分に納得できるように見えることが多いため、レビューは容易ではなくむしろ困難になります。 Stack Overflow の 2025 年の調査によると、AI の精度に対する信頼は 29% に低下しました。レビュー担当者は、意図、アーキテクチャ、実行時の動作の間の微妙な不一致を探す必要があります。
Codacy は、人間によるレビューの前にコードの品質、セキュリティ、カバレッジ、ポリシーのチェックを自動化し、レビュー担当者がアーキテクチャ、正確性、保守性に集中できるようにします。
プル リクエストを遅らせることなく AI が生成したコードをレビューする方法
AI を拡張するチームは、それに比例して人間の注意を必要とせずに、量の増加を吸収できる検証システムへの投資に成功しています。
CircleCI のデータでは、少数のチームでメインブランチのスループットが 26% 増加し、フィーチャーブランチのアクティビティが 85% 急増しました。違いは、より強力な自動チェックとレビュー コメントの信号対雑音比の向上、およびより明確なマージ ポリシーでした。
実際のアプローチには 3 つの層があります。
人間によるレビューの前にベースライン チェックを自動化します。フォーマット、リンティング、SAST の検出結果、SCA と依存関係のリスク、シークレットの検出、テスト カバレッジの変更、複雑さのしきい値はすべて、人間が差分ファイルを開く前に実行できます。これらのチェックのいずれかが失敗すると、PR はレビュー キューに到達しません。
AI 支援レビューを使用して、レビュー担当者の立ち上げコストを削減します。有用な AI レビュー担当者が、変更内容を要約し、リスクを強調し、結果を重大度ごとにグループ化します。人間のレビュー担当者は、空の差分から開始することはありません。彼らは、どこに注意を向けるべきかを構造化した概要から始めます。
人間によるレビューは判断の判断のために取っておきます。アーキテクチャの整合性、ビジネス ロジックの正確性、長期的な保守性、チーム間の影響には、自動化ツールでは完全にはキャプチャできないコンテキストが必要です。人間の審査員は審査員に集中する

ツールが一貫して検出できる問題をスキャンするのではなく、gment を実行します。
この階層化モデルにより、基準を下げることなく PR を継続的に進めることができます。
自動化ツールがレビュープロセスで所有できるもの
自動化されたツールと AI 支援のレビュー担当者は、再現性のある問題の初回検出、要約、適用を独自に行うことができます。決定的チェック、つまり同じ入力が与えられるたびに同じ結果を生成するチェックは、次のカテゴリに適しています。
静的分析の結果: Lint、型チェック、コード スタイル違反。
セキュリティ パターン: 既知の脆弱性パターン、安全でない依存関係、秘密の漏洩。
テストとカバレッジの変更: テストが合格したかどうか、カバレッジのしきい値が満たされているかどうか。
複雑さと重複: 保守性のしきい値を超えたファイルにフラグを立てます。
ポリシー違反: 合格/不合格条件として表現できるプロジェクト固有のルール。
PR が開く前またはその直後に自動チェックが実行されると、問題が早期に表面化します。レビュー担当者は、ツールで検出できる問題の発見に時間を費やす必要はありません。また、調査結果がグループ化されて要約されると、レビュー担当者はすべての行に目を通すのではなく、どこに注意を払うかを決定できます。
Codacy のようなプラットフォームは、リポジトリが接続されているときに自動チェックを実行し、人間のレビュー担当者が差分を開く前に、複雑さ、重複、テスト カバレッジ、セキュリティの問題、PR の完全性に関する結果を提供できます。
人間の査読者が依然として評価する必要があるもの
人間のレビュー担当者は、自動ツールでは再現できないコンテキストをもたらします。強力な自動チェックを使用しても、特定の質問には判断が必要です。
この変更はシステムのアーキテクチャと一致していますか? AI によって生成されたコードは、多くの場合、狭いプロンプトを満たしますが、より広範なアーキテクチャ上の期待に違反します。
ビジネスロジックは正しいですか？ツールはコードが実行されることを検証できますが、コードが実行されることは検証できません。

正しい問題を解決します。
これは保守可能でしょうか?現在機能するコードでも、既存の抽象化をバイパスしたり、ロジックを重複したりすると、明日にはメンテナンスの負担が生じる可能性があります。
チーム間の影響は何ですか?共有コンポーネントまたは API に関わる変更は、差分には表示されない形で他のチームに影響を与える可能性があります。
これが最も簡単な解決策でしょうか? AI は、既存の内部 API やユーティリティを十分に認識していない場合に、インライン コードを作成することがよくあります。
説明責任は人間にあります。しかし、ベースラインチェックがすでに処理されている場合、人間によるレビューは機械的な検査ではなく、判断を重視するものになります。
一般的な AI レビュー担当者が重要な問題を見逃す理由
一般的な AI レビュー担当者は、リポジトリ メモリを持たない新人エンジニアのように動作します。彼らは明らかな問題を発見するかもしれませんが、成熟したコードベースで実際に重要な問題を見逃すことがよくあります。
v2 が正規パスの場合、PR は v1 ミドルウェアを変更します。
コンポーネントは、デザイン システムにすでに存在するドロップダウン動作を複製します。
アーキテクチャがサービス層アクセスを必要とする場合でも、コントローラーはリポジトリを直接呼び出します。
提案されたリファクタリングは、プロジェクトの貢献手順のスコープ制約に違反しています。
これらの問題のそれぞれには、リポジトリ レベルのコンテキスト、つまり規約、アーキテクチャの決定、プロジェクト固有のルールに関する知識が必要です。このコンテキストがなければ、AI レビュー担当者は浅いフィードバックまたはノイズの多い誤検知を生成します。
AI レビュー ツールを評価しているチームは、多くの場合、有益な発見と、依然として人間によるフィルタリングを必要とする信号の少ないフィードバックが混在したものを報告します。古いトレーニング データでは、依存関係のバージョンに関する誤った主張が生成される可能性があります。パターンベースのセキュリティチェックでは、安全なコードに脆弱性があるとしてフラグを立てることができます。リポジトリ レベルの指示は完全に無視できます。
同時に、AI レビュー担当者は実際のバグを発見します: パッケージ deduplicat

イオンの問題、URL エンコーディングの問題、パターン マッチング エラー、ワークフロー トリガーの欠落などです。この値は実際のものですが、決定論的分析とリポジトリ対応のコンテキストとの統合に依存します。
コンプライアンス要件がレビュー モデルをどのように形成するか
多くの組織では、PR ごとに少なくとも 1 人の非作成者による承認が必要です。この要件は、レビュー全体を手動で行う必要があるという意味ではありません。
自動チェックにより、人間の承認者が検査する範囲を減らすことができます。ツールは、どのチェックが実行されたか、どの結果が導入されたか、どのポリシーが適用されたかの証拠を提供できます。人間による承認は、一貫した自動化された証拠によって裏付けられると、より意味のあるものになります。
SOC 2、ISO 27001、ISO 42001、HIPAA などのコンプライアンス フレームワークの場合、問題はコントロールに一貫性があり、監査可能であるかどうかです。同じポリシー条件下で同じ入力が常に同じ決定を生成する決定論的強制は、この要件をサポートします。確率的 AI レビューは問題を表面化する可能性がありますが、実稼働に伴う変更に対する唯一の強制レイヤーとして機能することはできません。
Codacy は、実行されたチェックと検出された結果を文書化したエクスポート可能なコンプライアンス レポートを提供し、数週間にわたるスクランブルからダッシュボードのエクスポートまで監査の準備を軽減します。
最適化された PR レビューの上限
ここで説明する階層化モデル、自動化されたベースライン チェック、AI 支援のトリアージ、人間による集中的なレビューは、現在プレッシャーにさらされているチームにとって実用的な橋渡し状態です。基準を下げることなく PR を継続します。ただし、このモデルには天井があります。
人間のレビュー能力は直線的に拡大しません。エージェント システムは多くの変更を並行して生成できます。レビュー担当者は、AI によって生成されたすべての差分の完全なコンテキストを再構築することはできません。
一部のチームはすでに検証を承認および申請から分離し始めています

導入リスクから楕円形になります。一部の企業は、強力なテストとロールバック メカニズムによって保護された変更に対して、最初にマージし、後でレビューするワークフローを実験しています。例外、高リスク領域、アーキテクチャの変更については人間によるレビューを留保している企業もいます。
PR プロセスは、人間がコードを書き、人間がコードをレビューし、変更の量はチームのレビュー能力の範囲内にとどまるという安定した前提に基づいて構築されています。その想定は崩れ始めています。
今のところ、より強力な自動ガードレールとより優れたトリアージにより出血を止めることができます。しかし、AI がより多くのコードを生成するにつれて、業界はより抜本的なレビュー モデルに移行する可能性があります。もはや問題は、エンジニアがコードをより速く作成できるかどうかではありません。チームが標準を遅らせたり低下させたりすることなく、そのコードを検証し推進できるかどうかが重要です。
AI によって生成されたコードや人間が作成したコードは、大規模な場合には予​​測できない動作をする可能性があります。 Codacy は、エンジニアリング システム全体でどこで強制が中断されているかを明らかにするのに役立ちます。
リポジトリを無料でスキャン →
毎月のニュースレターで最新情報を入手してください。
AIインベントリ、
ソフトウェアエンジニアリングにおけるAI
2026/12/06
エンジニアリング チームにおける AI ツール導入の背後にある可視性の問題
経営陣はエンジニアリングチームにAI導入を加速するよう求めている。モー

[切り捨てられた]

## Original Extract

See how AI-generated code impacts pull request reviews, creating bottlenecks and changing team dynamics. Learn how to maintain code quality and efficiency.

AI Is Breaking Code Review: How Engineering Teams Survive the PR Bottleneck
-->
AI coding tools have made it easier to produce code, but they have not made it easier to ship it safely.
Pull request queues are growing faster than review capacity. CircleCI's 2026 data shows feature branch throughput up 59% year over year, while main branch throughput for the median team actually fell. The bottleneck has moved from writing code to deciding whether code is safe to merge. This article covers why AI-generated code creates review pressure, what automated tools can handle, and how engineering teams can keep PRs moving without lowering their standards.
Why AI-generated code creates a review bottleneck
What makes reviewing AI-generated code different
How to review AI-generated code without slowing down pull requests
What automated tools can own in the review process
What human reviewers still need to evaluate
Why generic AI reviewers miss critical issues
How compliance requirements shape the review model
The ceiling on optimized PR review
Why AI-generated code creates a review bottleneck
To review AI-generated code without slowing down pull requests, you have to move baseline checks away from human eyes and focus human review on intent and architectural fit. The engineer who prompted the AI still owns the output. But the review process itself can shift: automated checks handle formatting, security patterns, and known vulnerabilities, while humans concentrate on whether the change actually solves the right problem.
Here is the main tension: AI-assisted development has changed the ratio between code production and review capacity. Engineers can generate more code in less time, yet the team's ability to validate that code has not scaled at the same rate.
CircleCI's 2026 State of Software Delivery report analyzed more than 28 million CI workflow runs across over 22,000 organizations. As mentioned earlier, overall throughput grew 59% year over year. At the same time, throughput on feature branches increased 15% for the median team while main-branch throughput fell nearly 7%, and main-branch success rates dropped to 70.8%.
More code is entering the pipeline, but less of it is reaching production successfully. The bottleneck has moved from writing code to deciding whether code is safe to merge.
What makes reviewing AI-generated code different
Code review has always carried hidden operational costs. Context switching between building and reviewing slows both activities. Feedback loops can become contentious. And when PR queues grow, teams often reduce rigor to keep work moving — Faros AI data shows 31% more PRs merging with no review — leading to superficial approvals and skipped edge-case analysis.
AI-generated code amplifies review pressure in specific ways.
First, there is the volume problem. AI-assisted workflows produce more branches, more commits, and more PRs. A single engineer working with a coding assistant can open several PRs in the time it previously took to complete one.
Second, there is the context problem. When an AI agent generates code, the reviewer often receives a completed diff without the same implementation journey or decision trail. The reviewer has to reconstruct intent from the ticket, PR description, and code changes alone. LinearB's 2026 Software Engineering Benchmarks Report found that agentic AI PRs have a pickup time 5.3x longer than unassisted PRs. AI-assisted PRs wait 2.47x longer. Longer pickup times suggest reviewers are spending more time evaluating AI-generated changes, contributing to deeper review queues.
Third, there is the trust problem. AI-generated code often appears plausible enough to pass a casual read, which makes review harder rather than easier. Stack Overflow's 2025 survey shows trust in AI accuracy has fallen to 29% . Reviewers have to look for subtle mismatches between intent, architecture, and runtime behavior.
Codacy automates code quality, security, coverage, and policy checks before human review, helping reviewers focus on architecture, correctness, and maintainability.
How to review AI-generated code without slowing down pull requests
The teams that scale AI successfully invest in validation systems that absorb increased volume without requiring proportionally more human attention.
In CircleCI's data, a minority of teams saw main-branch throughput grow 26% while feature-branch activity surged 85%. The difference was stronger automated checks and better signal-to-noise in review comments, as well as clearer merge policies.
The practical approach has three layers:
Automate baseline checks before human review. Formatting, linting, SAST findings, SCA and dependency risk, secrets detection , test coverage changes, and complexity thresholds can all run before a human opens the diff. If any of those checks fail, the PR does not reach the review queue.
Use AI-assisted review to reduce reviewer startup cost. A useful AI reviewer summarizes what changed, highlights risks, and groups findings by severity. The human reviewer does not start from a blank diff. They start with a structured overview of where to focus attention.
Reserve human review for judgment calls. Architectural alignment, business logic correctness, long-term maintainability, and cross-team impact require context that automated tools cannot fully capture. Human reviewers concentrate on judgment rather than scanning for issues that tools can detect consistently.
This layered model keeps PRs moving without lowering standards.
What automated tools can own in the review process
Automated tools and AI-assisted reviewers can own first-pass detection, summarization, and enforcement for repeatable issues. Deterministic checks, meaning checks that produce the same result every time given the same input, work well for the following categories:
Static analysis findings: Linting, type checks, and code style violations.
Security patterns: Known vulnerability patterns, insecure dependencies, and secrets exposure.
Test and coverage changes: Whether tests pass, whether coverage thresholds are met.
Complexity and duplication: Flagging files that exceed maintainability thresholds .
Policy violations: Project-specific rules that can be expressed as pass/fail conditions.
When automated checks run before or immediately when the PR opens, issues surface early. Reviewers do not spend time finding problems that tools can detect. And when findings are grouped and summarized, reviewers can decide where to spend attention rather than scanning every line.
Platforms like Codacy can run automated checks when a repository is connected, providing findings on complexity, duplication, test coverage, security issues, and PR completeness before a human reviewer opens the diff.
What human reviewers still need to evaluate
Human reviewers bring context that automated tools cannot replicate. Even with strong automated checks, certain questions require judgment:
Does this change align with the system's architecture? AI-generated code often satisfies a narrow prompt while violating broader architectural expectations.
Is the business logic correct? Tools can verify that code runs, but not that it solves the right problem.
Will this be maintainable? Code that works today can create maintenance burden tomorrow if it bypasses existing abstractions or duplicates logic.
What is the cross-team impact? Changes that touch shared components or APIs may affect other teams in ways that are not visible in the diff.
Is this the simplest solution? AI often writes inline code when it lacks sufficient awareness of existing internal APIs or utilities.
Accountability stays with the human. But the human review becomes more about judgment and less about mechanical inspection when baseline checks are already handled.
Why generic AI reviewers miss critical issues
Generic AI reviewers behave like a new engineer with no repository memory. They may spot obvious issues, but they often miss the problems that actually matter in mature codebases.
A PR modifies v1 middleware when v2 is the canonical path.
A component duplicates dropdown behavior that already exists in the design system.
A controller calls repositories directly even though the architecture requires service-layer access.
A suggested refactor violates the scope constraints in the project's contribution instructions.
Each of those issues requires repository-level context: knowledge of conventions, architecture decisions, and project-specific rules. Without this context, an AI reviewer produces either shallow feedback or noisy false positives.
Teams evaluating AI review tools often report a mix of useful findings and low-signal feedback that still requires human filtering. Stale training data can produce false claims about dependency versions. Pattern-based security checks can flag safe code as vulnerable. Repository-level instructions can be ignored entirely.
At the same time, AI reviewers do catch real bugs: package deduplication issues, URL-encoding problems, pattern matching errors, and missing workflow triggers. The value is real, but it depends on integration with deterministic analysis and repository-aware context.
How compliance requirements shape the review model
Many organizations require at least one non-author human approval on every PR. That requirement does not mean the entire review has to be manual.
Automated checks can reduce the scope of what the human approver inspects. Tools can provide evidence of which checks ran, which findings were introduced, and what policies were enforced. The human approval becomes more meaningful when it is supported by consistent automated evidence.
For compliance frameworks like SOC 2, ISO 27001 , ISO 42001 , or HIPAA, the question is whether controls are consistent and auditable. Deterministic enforcement, where the same input always produces the same decision under the same policy conditions, supports this requirement. Probabilistic AI review can surface issues, but it cannot serve as the sole enforcement layer for production-bound changes.
Codacy provides exportable compliance reports that document which checks ran and what findings were detected, reducing audit preparation from weeks of scrambling to a dashboard export.
The ceiling on optimized PR review
The layered model described here, automated baseline checks, AI-assisted triage, and focused human review, is the practical bridge state for teams under pressure now. It keeps PRs moving without lowering standards. However, this model has a ceiling.
Human review capacity does not scale linearly. Agentic systems can produce many parallel changes. Reviewers cannot reconstruct full context for every AI-generated diff.
Some teams are already beginning to separate validation from approval and approval from deployment risk. Some are experimenting with merge-first, review-later workflows for changes protected by strong tests and rollback mechanisms. Others are reserving human review for exceptions, high-risk areas, and architectural changes.
The PR process was built around a stable assumption: humans write code, humans review code, and the volume of change remains within the review capacity of the team. That assumption is starting to fail.
For now, stronger automated guardrails and better triage can stop the bleeding. But as AI generates more of the code, the industry will likely move toward more radical review models. The question is no longer whether engineers can produce code faster. It is whether teams can validate and promote that code without slowing down or lowering their standards.
AI-generated and human-written code can behave unpredictably at scale. Codacy helps reveal where enforcement breaks across your engineering system.
Scan your repository for free →
Stay updated with our monthly newsletter.
AI Inventory,
AI in Software Engineering
12/06/2026
The Visibility Problem Behind AI Tool Adoption in Engineering Teams
Executives are asking engineering teams to accelerate AI adoption. Mo

[truncated]
