---
source: "https://blog.codacy.com/the-visibility-problem-behind-ai-tool-adoption-in-engineering-teams"
hn_url: "https://news.ycombinator.com/item?id=48584626"
title: "Show HN: Know which AI coding tools/MCPs/etc. are used to ship software"
article_title: "The Visibility Problem Behind AI Tool Adoption in Engineering Teams"
author: "ARayOutOfBounds"
captured_at: "2026-06-18T13:17:23Z"
capture_tool: "hn-digest"
hn_id: 48584626
score: 2
comments: 0
posted_at: "2026-06-18T12:59:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Know which AI coding tools/MCPs/etc. are used to ship software

- HN: [48584626](https://news.ycombinator.com/item?id=48584626)
- Source: [blog.codacy.com](https://blog.codacy.com/the-visibility-problem-behind-ai-tool-adoption-in-engineering-teams)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T12:59:11Z

## Translation

タイトル: HN を表示: AI コーディング ツール/MCP などを知る。ソフトウェアの出荷に使用されます
記事のタイトル: エンジニアリング チームにおける AI ツール導入の背後にある可視性の問題
説明: エンジニアリング チーム内での AI ツールの使用状況を可視化し、ガバナンスの課題に対処し、より良い結果を得るためにツールの追跡を最適化する方法を学びます。
HN テキスト: エンジニアリングのリーダーは、ソフトウェアの構築にどのツールが使用されているかを把握できていません。これは、この種のスキャナーとしては初めて、リポジトリレベルの証拠を提供します。

記事本文:
エンジニアリング チームにおける AI ツール導入の背後にある可視性の問題
-->
経営陣はエンジニアリングチームにAI導入を加速するよう求めている。しかし、ほとんどのエンジニアリング リーダーは、現在リポジトリ全体で実際にどの AI コーディング アシスタントが使用されているのかという基本的な質問に答えることができません。
ギャップは、信頼できる在庫がないことです。この記事では、可視性がガバナンスの最初の課題である理由、リポジトリ レベルのスキャンで実際の導入パターンが明らかになること、開発者が依存するツールをブロックすることなく AI ツールの追跡を運用する方法について説明します。
AI ツールの可視化がガバナンスの最初の課題である理由
エンジニアリング ワークフローにおけるシャドウ AI の様子
開発チーム全体で AI ツールの使用がどの程度普及しているか
リポジトリ レベルのスキャンで AI 導入について明らかになったもの
既存のガバナンスアプローチが AI ツールの使用を見逃している理由
AI ツールの無秩序な拡大がコードの品質とセキュリティに与える影響
AI インベントリがリポジトリ全体でキャプチャするもの
エンジニアリング チーム向けに AI の可視性を運用する方法
リーダーが AI 導入について報告するのに役立つ指標は何ですか
AI ツールの使用状況を追跡する際のよくある間違い
AI ツールの可視化がガバナンスの最初の課題である理由
エンジニアリング チームが使用する AI ツールを可視化するには、AI 支援コーディングの定量的な追跡と定性的な開発者のフィードバックを組み合わせることができます。 AI の使用状況が監視されていないことが一般的であるため、専門的な分析により、ツールの使用状況、セキュリティ体制、生産性への影響を追跡することができます。ただし、課題は、ほとんどのエンジニアリング リーダーが、現在リポジトリ全体でどの AI コーディング アシスタントがアクティブになっているかという単純な質問に答えることができないことです。
矛盾しているのは、取締役会や経営陣が配信速度を向上させるためにチームに AI の導入を促していることです。同時に、リーダーには信頼できる在庫が不足していることがよくあります。

開発者が実際に使用するツール。 AI の導入が加速するにつれて、レポートのギャップは拡大します。
可視性の問題は、AI が使用されているかどうかではなく、エンジニアリング組織全体で AI の使用状況を監査可能な視点でリーダーが把握しているかどうかです。
エンジニアリング ワークフローにおけるシャドウ AI の様子
ソフトウェア開発におけるシャドウ AI は、従業員がランダムなチャットボットを使用するだけではありません。エンジニアリングの文脈では、シャドウ AI には、集中管理された IT やセキュリティの可視性を回避するさまざまなツールやアーティファクトが含まれています。
開発ワークフローにおけるシャドウ AI の一般的な形式は次のとおりです。
IDE の AI コーディング アシスタント: 開発者環境にローカルにインストールされた GitHub Copilot、Cursor、Claude Code などの拡張機能
CLI ベースのエージェント: マネージド SaaS プラットフォームの外部でコードを生成または変更するコマンドライン ツール
AI レビュー ツール: プル リクエストを分析したり、変更を提案したりするサードパーティ サービス
プロンプト ファイルと生成された構成: リポジトリにコミットされたワークフロー アーティファクト、プロンプト テンプレート、または AI によって生成された構成ファイル
個人アカウント アクセス: IT チームやセキュリティ チームには表示されないツールへの個人サブスクリプションを使用する開発者
ガバナンスの問題は、管理されていない AI の使用です。 IBM は、高レベルのシャドウ AI を導入した組織では、平均 67 万ドルの追加侵害コストが発生していることを発見しました。エンジニアリング リーダーにとっての課題は、AI の使用状況を十分に可視化し、ポリシー、ガバナンス、施行をリポジトリで実際に起こっていることと一致させることです。
開発チーム全体で AI ツールの使用がどの程度普及しているか
AI コーディング ツールはパイロット プログラムをはるかに超えています。現在、開発者のかなりの割合が AI アシスタントを毎日使用しており、組織全体のコード作成において AI 支援コードが占める割合は増加しています。
いくつかのパターンが注目に値します。多くの開発者は、

個人の ChatGPT や Claude サブスクリプションなどの個人アカウントを通じて AI ツールにアクセスします。平均的な開発チームは複数の AI ツールを同時に使用しますが、多くの場合、一元的な追跡は行われません。 GitHub Copilot、ChatGPT、Claude は開発ワークフロー全体で共通ですが、採用状況はチームやプロジェクトによって異なります。
正確な数字は、調査方法、企業規模、「AI 支援」の定義によって異なります。ただし、方向性は明らかです。AI の使用は、実験を超えて日常の開発実践に移行しています。
リポジトリ レベルのスキャンで AI 導入について明らかになったもの
調査と調達データが物語の一端を物語っています。リポジトリレベルのスキャンは別のことを伝えます。実際にコードベースに何がコミットされているかを分析すると、別の状況が見えてきます。
最近、800 以上の組織にわたる 6,700 以上のリポジトリをスキャンし、27 の異なるコーディング アシスタントを検出しました。 Claude Code が大差で最も頻繁に出現し、Cursor と Microsoft Copilot がそれに続きました。ロングテールには、Codex、Windsurf、Augment Code、JetBrains Junie、その他 20 以上のツールが含まれていました。
表面的には、27 人のアシスタントがツールの無秩序な拡大を示唆しています。ただし、組織レベルで見ると、平均的な組織でコーディング アシスタントが 2 人しかいなかったことがわかります。ほとんどの企業ではどこでも混乱が生じているわけではありません。 AI の導入は限定的ではありますが、多くの場合測定されていません。
可視性の問題は「ツールが多すぎる」ということだけではありません。それは、リーダーがどのツールが存在するか、どこに表示されるか、リポジトリごとに使用法がどのように異なるのかを知らないということです。
既存のガバナンスアプローチが AI ツールの使用を見逃している理由
ガバナンス ポリシーは、実際の開発者の行動が観察されない場所に存在することがよくあります。セキュリティに関するアンケート、調達の承認、使用許諾ポリシー、SSO ツールのリスト、および手動によるチームの自己報告により、すべての意図が把握されます。彼らは現実を捉えることはほとんどありません

。
ここでは、ガバナンス管理が一般的なエンジニアリング パスを見逃している箇所を示します。
個人アカウント: エンタープライズ SSO をバイパスする個人サブスクリプションを使用する開発者
IDE 拡張機能: SaaS 管理ツールには表示されない、ローカルにインストールされた拡張機能
CLI エージェント: 集中管理されたプラットフォームの外部で使用されるコマンドライン ツール
リポジトリ アーティファクト: コードベースにコミットされた生成されたファイル、プロンプト、ワークフロー、またはツール トレース
ガバナンスおよびセキュリティ検証ポリシーを導入している組織でも、Cursor、Claude、およびその他のツールを並行して使用している開発者がいる可能性があります。開発者による採用のプレッシャーは現実のものです。エンジニアは、より迅速に作業を進めるのに役立つツールを選択します。
ポリシーは意図を定義します。在庫は現実を明らかにします。
AI ツールの無秩序な拡大がコードの品質とセキュリティに与える影響
AI によって生成されたコードまたは AI 支援によるコードは、エンジニアリング チームのスケールとレビュー モデルを変更します。より多くのコードをより速く生成できます。レビュー担当者が生成されたコードが正しいと仮定すると、レビューの深さが減少する可能性があります。同様のパターンがリポジトリ間ですぐに複製される可能性があります。
セキュリティと品質のリスクには次のものが含まれます。
脆弱なコード パターン: 既知の脆弱性パターンを導入する提案が生成されます。 Veracode の 2025 年の調査では、AI によって生成されたコード サンプルのうち、ベンチマークで評価されたセキュリティ上の欠陥が含まれていないのは 55% のみであることが判明しました。これは、45% にテスト済みの脆弱性が少なくとも 1 つ含まれていることを意味します。
安全でない依存関係の推奨事項: 古いライブラリまたは脆弱なライブラリを提案する AI ツール
秘密の暴露: 管理されていないツールに貼り付けられた、またはリポジトリにコミットされた機密コンテキスト
正しくないアサーションを含む生成されたテスト: 合格するが、間違った動作をアサートするテスト
コンプライアンスのギャップ: 監査中に証明または説明できない AI の使用状況
インベントリは、SAST、SCA、機密検出、またはポリシー チェックの代わりにはなりません。それは可視レイヤーです

AI コード ガバナンスがどこに適用されるかをリーダーに伝えます。
調達の観点もあります。ほとんどの開発者がすでに特定のツールを使用している場合は、断片的な個人使用を許可するのではなく、エンタープライズ プランにアップグレードして一元管理する価値があるかもしれません。
AI インベントリがリポジトリ全体でキャプチャする内容
Codacy の AI インベントリは、コードベース全体で見つかる AI 資産、ツール、ワークフロー、シグナルをリポジトリ レベルで表示します。これにより、分散したリポジトリの証拠が、エンジニアリング リーダーが行動できるビューに変わります。
有用なインベントリには通常、次のものが含まれます。
Which AI models and providers are referenced across repositories (GPT-4o, Claude Sonnet, Gemini, OpenAI SDK, Anthropic SDK, etc.)
The goal is to provide enough visibility that leaders can compare detected usage against policy and decide where enforcement applies.
エンジニアリング チーム向けに AI の可視性を運用する方法
アクションを伴わない可視化は、単なるレポートにすぎません。 AI インベントリの実際的な価値は、AI インベントリを執行や意思決定に結び付けることで生まれます。
合理的なアプローチは次の順序に従います。
Scan repositories to establish a baseline of AI tool and workflow traces
Compare detected usage against approved tools and policies
Segment repositories by risk: production-critical, regulated, customer-facing, internal, experimental
Align enforcement with repository risk level, applying AI guardrails for high-risk repositories
Add or tighten checks for SAST, SCA, secrets detection , and license compliance where AI usage is present
Track changes over time to see whether adoption is expanding or concentrating
Report trends to leadership in operational language: tools detected, repos affected, policy gaps, remediation progress
AI ツールの使用状況は急速に変化します。インベントリは 1 回限りの監査ではありません。 It works best when continuous or regularly refres

ヘド。
リーダーが AI 導入について報告するのに役立つ指標は何ですか
経営陣や取締役会のメンバーが AI 導入について質問すると、曖昧な答えは信頼を損ないます。特定の指標により、養子縁組の主張が証拠に変換されます。
有用なリーダーシップ指標には次のものがあります。
組織全体で検出された AI コーディング アシスタントの数
AI ツール トレースのあるリポジトリの数と割合
リポジトリの存在による上位 AI ツール
未承認の AI ツールまたはワークフローを含むリポジトリ
AI関連ポリシーの遵守状況
AI の使用が検出されたリポジトリのセキュリティ結果
時間の経過に伴う傾向: 新しいツール、新しいリポジトリ、ポリシーのギャップの解決
このような指標は、リーダーが調査や調達データだけに頼らずに導入を報告するのに役立ちます。これらは、AI の実現とリスク管理の間に橋渡しをします。
AI ツールの使用状況を追跡する際のよくある間違い
いくつかのパターンが AI 可視化の取り組みを台無しにします。早期にそれらに気づくことで、無駄な努力を避けることができます。
承認されたツールのリストのみに依存する: 開発者は、承認されたリストに決して表示されない個人アカウント、ローカル拡張機能、または CLI ツールを使用する可能性があります。
AI ガバナンスを法律または調達の問題としてのみ扱う: エンジニアリングでは、ガバナンスはリポジトリとプル リクエストに関係します。
承認された代替手段を提供せずにツールをブロックする: リスクを軽減せずにコントロールの配信が遅い場合、開発者はコントロールを回避することになる
ライセンス数のみで導入を測定する: ライセンスは、何が購入されたかを示すものであり、何がコードに影響を与えたかを示すものであるとは限りません
事件量が少ないことはリスクが低いことを意味すると仮定する: 発見の欠如は、暴露がないことではなく、可視性の欠如を反映している可能性があります
目標は、リーダーがどのツールが使用されているかを把握し、そのツールがどこに表示されるかを理解し、リスク レベルごとに許容される使用を定義し、セキュリティと品質チェックを一貫して実施し、リーダーシップと品質に関する証拠を作成する、管理された導入モデルを作成することです。

コンプライアンス。
AI の導入はすでに日常の開発ワークフローに浸透しています。当面のリーダーシップのギャップは可視性です。リーダーは、実際に何が使用されているかを知らずに、AI 導入を信頼性をもって報告したり、AI リスクを管理したりすることはできません。
今後の実際的な道は、リポジトリ全体にわたる AI インベントリから始まります。これを使用して、ポリシー、施行、レポートを調整します。次に、配信パイプラインをすでに保護している既存のコード品質およびセキュリティ管理にそれを接続します。
Codacy の AI インベントリは、AI ツール、ワークフロー、関連アーティファクトのリポジトリ レベルのビューを提供し、エンジニアリング リーダーが使用状況を追跡し、ポリシーのギャップを特定し、開発チーム全体の可視性を向上させるのに役立ちます。
リポジトリを無料でスキャン →
毎月のニュースレターで最新情報を入手してください。
AIインベントリ、
ソフトウェアエンジニアリングにおけるAI、
ソフトウェアのリスクとコンプライアンス
2026/12/06
2026 年に監査人は AI 生成コードについて何を質問するでしょうか
開発者はすでに AI コーディング ツールを使用しています。監査人が問い始めているのは、チームに AI ポリシーがあるかどうかではなく、それができるかどうかです。
コードシープラットフォーム、
新機能、
AI インベントリ
2026/07/04
開発者はどの AI ツールを使用していますか? AI インベントリによるシャドウ AI 問題の解決
これをイメージしてください。それは木曜日の午後です

[切り捨てられた]

## Original Extract

Discover how to gain visibility into AI tool usage within engineering teams, addressing governance challenges and optimizing tool tracking for better outcomes.

Eng leaders lack visibility into which tools are being used to build software. This is the first scanner of its kind to provide repo-level evidence.

The Visibility Problem Behind AI Tool Adoption in Engineering Teams
-->
Executives are asking engineering teams to accelerate AI adoption. Most engineering leaders, though, cannot answer a basic question: which AI coding assistants are actually in use across our repositories right now?
The gap is the lack of a reliable inventory. This article covers why visibility is the first governance challenge, what repository-level scanning reveals about real adoption patterns, and how to operationalize AI tool tracking without blocking the tools developers rely on.
Why visibility into AI tools is the first governance challenge
What shadow AI looks like in engineering workflows
How widespread is AI tool usage across development teams
What repository-level scanning reveals about AI adoption
Why existing governance approaches miss AI tool usage
How AI tool sprawl affects code quality and security
What an AI inventory captures across repositories
How to operationalize AI visibility for engineering teams
What metrics help leaders report on AI adoption
Common mistakes when tracking AI tool usage
Why visibility into AI tools is the first governance challenge
To gain visibility into AI tools used by engineering teams, you can combine quantitative tracking of AI-assisted coding with qualitative developer feedback. Unmonitored AI usage is common, so specialized analytics help track tool usage, security posture, and productivity impact. The challenge, though, is that most engineering leaders cannot answer a straightforward question: which AI coding assistants are active across our repositories right now?
The contradiction is that boards and executives are pushing teams to adopt AI to improve delivery speed. At the same time, leadership often lacks a reliable inventory of which tools developers actually use. The reporting gap widens as AI adoption accelerates.
The visibility problem is not about whether AI is being used, but rather about whether leadership has an auditable view of AI usage across the engineering organization.
What shadow AI looks like in engineering workflows
Shadow AI in software development goes beyond employees using random chatbots. In engineering contexts, shadow AI includes a range of tools and artifacts that often bypass centralized IT or security visibility.
Here are common forms of shadow AI in development workflows:
AI coding assistants in IDEs: Extensions like GitHub Copilot, Cursor, or Claude Code installed locally in developer environments
CLI-based agents: Command-line tools that generate or modify code outside of managed SaaS platforms
AI review tools: Third-party services that analyze pull requests or suggest changes
Prompt files and generated config: Workflow artifacts, prompt templates, or AI-generated configuration files committed to repositories
Personal account access: Developers using personal subscriptions to tools that are not visible to IT or security teams
The governance issue is unmanaged AI use. IBM found that organizations with high levels of shadow AI experienced an average of $670,000 in additional breach costs . For engineering leaders, the challenge is gaining enough visibility into AI usage to align policy, governance, and enforcement with what is actually happening in repositories.
How widespread is AI tool usage across development teams
AI coding tools have moved well beyond pilot programs. A significant share of developers now use AI assistants daily, and AI-assisted code represents a growing portion of code creation across organizations.
Several patterns are worth noting. Many developers access AI tools through personal accounts, including personal ChatGPT or Claude subscriptions. The average development team uses multiple AI tools simultaneously, often without centralized tracking. GitHub Copilot, ChatGPT, and Claude are common across development workflows, though adoption varies by team and project.
Exact figures depend on survey methodology, company size, and how "AI-assisted" is defined. The direction, however, is clear: AI use has moved beyond experimentation into everyday development practice.
What repository-level scanning reveals about AI adoption
Surveys and procurement data tell part of the story. Repository-level scanning tells another. When you analyze what is actually committed to codebases, a different picture emerges.
Recently, we scanned over 6,700 repositories across more than 800 organizations and detected 27 different coding assistants. Claude Code appeared most frequently by a wide margin, followed by Cursor and Microsoft Copilot. The long tail included Codex, Windsurf, Augment Code, JetBrains Junie, and more than 20 other tools.
On the surface, 27 assistants suggests tool sprawl. At the organization level, though, the average org showed traces of only two coding assistants. Most companies do not have chaos everywhere. They have limited but often unmeasured AI adoption.
The visibility problem is not only "too many tools." It is that leaders do not know which tools are present, where they appear, or how usage differs by repository.
Why existing governance approaches miss AI tool usage
Governance policies often live in places that do not observe actual developer behavior. Security questionnaires, procurement approvals, acceptable-use policies, SSO tool lists, and manual team self-reporting all capture intent. They rarely capture reality.
Here is where governance controls miss common engineering paths:
Personal accounts: Developers using personal subscriptions that bypass enterprise SSO
IDE extensions: Locally installed extensions that do not appear in SaaS management tools
CLI agents: Command-line tools used outside centrally managed platforms
Repository artifacts: Generated files, prompts, workflows, or tool traces committed into codebases
Even organizations with governance and security validation policies may still have developers using Cursor, Claude, and other tools in parallel. Developer adoption pressure is real: engineers choose tools that help them move faster.
Policies define intent. Inventories reveal reality.
How AI tool sprawl affects code quality and security
AI-generated or AI-assisted code changes the scale and review model for engineering teams. More code can be produced faster. Review depth may decrease when reviewers assume generated code is correct. Similar patterns may be replicated across repositories quickly.
Security and quality risks include:
Vulnerable code patterns: Generated suggestions that introduce known vulnerability patterns. Veracode’s 2025 research found that only 55% of AI-generated code samples were free of the security flaws evaluated in its benchmark , meaning 45% contained at least one tested vulnerability.
Insecure dependency recommendations: AI tools suggesting outdated or vulnerable libraries
Secrets exposure: Sensitive context pasted into unmanaged tools or committed to repositories
Generated tests with incorrect assertions: Tests that pass but assert the wrong behavior
Compliance gaps: AI usage that cannot be evidenced or explained during audits
Inventory is not a substitute for SAST, SCA, secrets detection, or policy checks. It is the visibility layer that tells leaders where AI code governance applies.
There is also a procurement angle. If most developers already use a particular tool, it may be worth upgrading to an enterprise plan and managing it centrally rather than allowing fragmented personal usage.
What AI Inventory captures across repositories
Codacy's AI Inventory is a repository-level view of AI assets, tools, workflows, and signals found across codebases. It turns scattered repository evidence into a view engineering leaders can act on.
A useful inventory typically captures:
Which AI models and providers are referenced across repositories (GPT-4o, Claude Sonnet, Gemini, OpenAI SDK, Anthropic SDK, etc.)
The goal is to provide enough visibility that leaders can compare detected usage against policy and decide where enforcement applies.
How to operationalize AI visibility for engineering teams
Visibility without action is just reporting. The practical value of AI Inventory comes from connecting it to enforcement and decision-making.
A reasonable approach follows this sequence:
Scan repositories to establish a baseline of AI tool and workflow traces
Compare detected usage against approved tools and policies
Segment repositories by risk: production-critical, regulated, customer-facing, internal, experimental
Align enforcement with repository risk level, applying AI guardrails for high-risk repositories
Add or tighten checks for SAST, SCA, secrets detection , and license compliance where AI usage is present
Track changes over time to see whether adoption is expanding or concentrating
Report trends to leadership in operational language: tools detected, repos affected, policy gaps, remediation progress
AI tool usage changes quickly. Inventory is not a one-time audit. It works best when continuous or regularly refreshed.
What metrics help leaders report on AI adoption
When executives or board members ask about AI adoption, vague answers erode confidence. Specific metrics convert adoption claims into evidence.
Useful leadership metrics include:
Number of AI coding assistants detected across the organization
Number and percentage of repositories with AI tool traces
Top AI tools by repository presence
Repositories with unapproved AI tools or workflows
AI-related policy compliance status
Security findings in repositories with detected AI usage
Trend over time: new tools, new repositories, resolved policy gaps
Metrics like these help leaders report adoption without relying only on surveys or procurement data. They create a bridge between AI enablement and risk management.
Common mistakes when tracking AI tool usage
Several patterns undermine AI visibility efforts. Recognizing them early helps avoid wasted effort.
Relying only on sanctioned-tool lists: Developers may use personal accounts, local extensions, or CLI tools that never appear in approved lists
Treating AI governance as only a legal or procurement issue: In engineering, governance connects to repositories and pull requests
Blocking tools without offering approved alternatives: Developers will route around controls if controls slow delivery without reducing risk
Measuring adoption only by license count: Licenses show what was bought, not necessarily what influenced code
Assuming low incident volume means low risk: Lack of findings may reflect lack of visibility, not absence of exposure
The goal is to create a managed adoption model where leaders know which tools are in use, understand where they appear, define acceptable use by risk level, enforce security and quality checks consistently, and produce evidence for leadership and compliance.
AI adoption has already reached everyday development workflows. The immediate leadership gap is visibility. Leaders cannot credibly report AI adoption or manage AI risk without knowing what is actually being used.
The practical path forward starts with AI inventory across repositories. Use it to align policy, enforcement, and reporting. Then connect it to existing code quality and security controls that already protect the delivery pipeline.
Codacy's AI Inventory provides a repository-level view of AI tools, workflows, and related artifacts, helping engineering leaders track usage, identify policy gaps, and improve visibility across development teams.
Scan your repository for free →
Stay updated with our monthly newsletter.
AI Inventory,
AI in Software Engineering,
Software Risk & Compliance
12/06/2026
What Auditors Will Ask About AI-Generated Code in 2026
Developers are already using AI coding tools. The question auditors are starting to ask is not whether your team has an AI policy, but whether you can...
Codacy Platform,
New Features,
AI Inventory
07/04/2026
Which AI Tools Are Your Devs Using? Solving the Shadow AI Problem with AI Inventory
Picture this. It's a Thursday afternoon a

[truncated]
