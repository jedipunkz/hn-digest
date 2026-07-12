---
source: "https://github.com/sekacorn/Linux-of-Ai"
hn_url: "https://news.ycombinator.com/item?id=48881655"
title: "Linux of AI open-source tools for reducing AI vendor lock-in"
article_title: "GitHub - sekacorn/Linux-of-Ai: Linux of AI is a free, open-source ecosystem for building portable, governed, measurable, and replaceable AI infrastructure. It connects tools for agent orchestration, private deployment, model switching, policy, ontology, audit logs, and AI cost measurement to help re\n[truncated]"
author: "Saya24f"
captured_at: "2026-07-12T15:50:57Z"
capture_tool: "hn-digest"
hn_id: 48881655
score: 1
comments: 0
posted_at: "2026-07-12T14:52:37Z"
tags:
  - hacker-news
  - translated
---

# Linux of AI open-source tools for reducing AI vendor lock-in

- HN: [48881655](https://news.ycombinator.com/item?id=48881655)
- Source: [github.com](https://github.com/sekacorn/Linux-of-Ai)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T14:52:37Z

## Translation

タイトル: AI ベンダーのロックインを軽減する AI オープンソース ツールの Linux
記事タイトル: GitHub - sekacorn/Linux-of-Ai: Linux of AI は、ポータブルで管理された、測定可能で置き換え可能な AI インフラストラクチャを構築するための無料のオープンソース エコシステムです。エージェント オーケストレーション、プライベート展開、モデル切り替え、ポリシー、オントロジー、監査ログ、AI コスト測定のためのツールを接続して、
[切り捨てられた]
説明: Linux of AI は、ポータブルで管理された、測定可能で置き換え可能な AI インフラストラクチャを構築するための無料のオープンソース エコシステムです。エージェント オーケストレーション、プライベート デプロイメント、モデル切り替え、ポリシー、オントロジー、監査ログ、AI コスト測定のためのツールを接続し、ベンダー ロックインの削減に役立ちます。 - セカコーン
[切り捨てられた]

記事本文:
GitHub - sekacorn/Linux-of-Ai: Linux of AI は、ポータブルで管理された、測定可能で置き換え可能な AI インフラストラクチャを構築するための無料のオープンソース エコシステムです。エージェント オーケストレーション、プライベート デプロイメント、モデル切り替え、ポリシー、オントロジー、監査ログ、AI コスト測定のためのツールを接続し、ベンダー ロックインの削減に役立ちます。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
アカウントを切り替えました

別のタブまたはウィンドウに表示されます。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
セカコーン
/
Aiのリナックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット docs docs 例/vendor-exit-demo 例/vendor-exit-demo README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
誰もが利用できる、無料でオープンなポータブル AI インフラストラクチャ。
検査可能、管理可能、測定可能、交換可能で、単一ベンダー内に囚われない AI システムを構築します。
ミッション •
問題点 •
プロジェクト •
アーキテクチャ •
原則 •
始めましょう
AI の Linux は初めてですか?まずは「Start Here」ガイドから始めて、問題に適したツールを選択してください。
また、AI Vendor Exit Demo の Linux を検査して、オフラインファーストの 1 つの決定的な例でエコシステム パターンを確認することもできます。
Linux of AI Vendor Exit Demo は、ポータブル オントロジー、ポリシー、ベンチマーク、監査証拠、コスト/結果の測定を使用して、チームが高価な AI モデルまたはロックインされた AI モデルの置き換えを評価するのにこのエコシステムがどのように役立つかを示します。
このデモは決定論的でオフラインファーストです。統合を示します
AIMeter OSS スタイルを含むポータブル ファイルとレポートを使用するパターン
コスト/結果 JSON および AIAuditLog スタイルの監査イベント JSONL をモデルからエクスポート
交代決定。
プロジェクト
PyPIパッケージ
役割
ステータス
エージェントフォージ
エージェントフォージ-oss
エージェントのオーケストレーション
発行済み
プライベートAIスタック
プライベートスタック
プライベート/ローカル展開
発行済み
モデルスワップベンチ
モデルスワップベンチ
モデルの置き換えとベンダーの終了レポート
公開されました。 AIベンダー撤退レポートを含む
OpenOntologyLite
オープンオントロジーライト
ポータブルなオントロジーとビジネス上の意味
発行済み
エージェントポリシーパック
エージェントポリシーパック
エージェント向けのコードとしてのポリシー

発行済み
AI監査ログ
aiaauditlog
AI監査証拠
発行済み
AIメーターOSS
照準器-oss
AI のコスト、使用量、効率、成果の測定
発行済み
これは何ではないのか
高リスクのワークフローにおける人間によるレビューの代替。
これは、モデルの移行が自動的に安全であることを証明します。
これは、AI システムを構築、評価、管理、監査、測定するためのオープン インフラストラクチャです。
実行または検査examples/vendor-exit-demo/
モデルの置き換えを決定するために ModelSwapBench を試してください
コストと成果の測定に AIMeter OSS をお試しください
AgentPolicyPack でガバナンスを追加し、AIAuditLog で証拠を監査します
Linux of AI は、ベンダー ロックインを軽減し、実用的な AI インフラストラクチャを誰もが利用できるようにするために作成された 7 つのプロジェクトのオープンソース エコシステムです。
いかなる組織も、AI システム、データ、コスト、ポリシー、将来の制御を単一のベンダーに明け渡すことを強制されるべきではありません。
AIは重要なインフラになりつつあります。しかし、多くの組織は、構築したシステムの移動、検査、管理が難しく、運用コストがますます高くなっていることに気づき始めています。
このエコシステムは別の道を提供するために存在します。
AI インフラストラクチャが存在するパス:
モデル、プロバイダー、環境間で移植可能
不透明なサービスの背後に隠れるのではなく検査可能
明示的なポリシーによって管理される
コスト、使用量、効率、成果を測定可能
モデルまたはプロバイダーがユーザーにサービスを提供しなくなった場合に置き換え可能
プライバシーや管理が重要な場合はローカルファースト
公益のための無料のオープンソース
このソフトウェアは、コア インフラストラクチャをペイウォールの背後に置くことなく、開発者、研究者、非営利団体、企業、政府、学生、コミュニティが引き続き利用できるようにすることを目的としています。
AI を導入している組織は、同じ問題点に繰り返し遭遇します。
アプリケーションは 1 つの MOD と密接に結合される

プロバイダー、SDK、API 形式、価格モデル、またはホストされたプラットフォーム。後でプロバイダーを置き換えると、大幅な書き直しが必要になる場合があります。
プロトタイプは手頃な価格かもしれませんが、実稼働環境での使用は急速に増加する可能性があります。多くの場合、チームには、ルーティング、予算編成、コストの測定、代替案の比較に関する明確な管理が欠けています。
組織は、支払いが過剰であるかパフォーマンスが低いことを認識しているかもしれませんが、アプリケーションを中断することなく別のモデルをテストする再現可能な方法がありません。
AI エージェントは、何ができるかを規定する明確で移植可能なルールがなくても、ツールを呼び出し、データにアクセスし、意思決定を行うことができます。
ログは一貫性がなかったり、プロバイダー固有のものであったり、不完全であったり、事後の検証が困難であることがよくあります。
API の成功は必ずしもビジネスの成功を意味するわけではありません。多くのシステムはリクエストとトークンを追跡しますが、AI が許容可能な結果を​​生成したかどうかは追跡しません。
組織の意味の喪失
ビジネス概念、関係、権限、およびアクションはアプリケーション コード内に直接埋め込まれることが多く、システムの理解と移行が困難になります。
プライバシーと展開の制約
組織によっては、機密データを外部プロバイダーに送信できない場合があります。スタック全体を再構築する必要のない、ローカルまたは制御された展開オプションが必要です。
AI の Linux は、これらの問題に 1 つの接続されたエコシステムとして対処します。
レイヤー
プロジェクト
解決するもの
リポジトリ
PyPI
組織的な意味
OpenOntologyLite
移植可能なエンティティ、関係、アクション、権限、およびアプリケーション コードの外でのビジネスの意味を定義します。
GitHub
PyPI
ガバナンス
エージェントポリシーパック
AI エージェントとエージェント ワークフローを管理するための移植可能なコードとしてのポリシーを提供します
GitHub
PyPI
オーケストレーション
エージェントフォージ
モデルのルーティング、予算、ガバナンス、メモリ、可観測性を備えたマルチエージェント システムを調整します。
GitHub
PyPI
プライベート展開
プライベートAIスタック
local-f を提供します

プライベート AI、ローカル モデル、RAG、コード レビュー、監査、および制御された展開の最初の出発点
GitHub
PyPI
機種変更
モデルスワップベンチ
別のモデルが、より低いコスト、レイテンシ、または運用リスクで許容可能な結果を提供できるかどうかをテストします。
GitHub
PyPI
運用上の証拠
AI監査ログ
AI システムとエージェント用の移植可能で改ざん明示的な監査イベント形式とローカル ツールキットを定義します
GitHub
PyPI
コストと成果
AIメーターOSS
プロバイダーへの請求がビジネス価値と等しいと仮定することなく、AI の使用量、コスト、効率、ビジネスの成果を測定します。
GitHub
PyPI
建築
OpenOntologyLite
│
▼
ポータブルな組織の意味
│
▼
エージェントポリシーパック
│
▼
移植可能なガバナンス ルール
│
▼
エージェントフォージ
│
▼
ポータブルエージェントオーケストレーション
│
▼
プライベートAIスタック
│
▼
プライベートで制御された展開
│
▼
モデルスワップベンチ
│
▼
モデルの置き換えと経済性の検証
│
▼
AI監査ログ
│
▼
ポータブルな改ざん明示的な運用証拠
│
▼
AIメーターOSS
│
▼
使用量、コスト、効率、成果の測定
各プロジェクトは独立して使用できます。これらは連携して、1 つのプロバイダーをアーキテクチャの永続的な中心にすることなく、AI システムを構築するためのポータブルな基盤を形成します。
問題点: 組織の知識は、ソース コード、データベース スキーマ、プロンプト、ベンダー固有のプラットフォームの中に埋もれています。
YAML または JSON での移植可能なオントロジー定義
型指定されたエンティティ、プロパティ、関係、およびアクション
検証と正規表現
比較、検査、文書化、および図のエクスポート
ローカルファーストかつオフライン運用
なぜそれが重要なのか:
あなたの組織の意味は、独自のプラットフォームに属するものではなく、あなたの組織に属するものであるべきです。
問題点: エージェントの動作は、散在するプロンプト指示とアプリケーション固有のチェックによって左右されることがよくあります。
明示的

エージェントアクションのcitルール
システム全体で再利用可能なガバナンス
レビュー可能でテスト可能なエージェントの動作の基盤
なぜそれが重要なのか:
ガバナンスは可視化され、移植可能であり、モデル自体から分離されている必要があります。
問題点: マルチエージェント システムは、1 つのプロバイダー、1 つのオーケストレーション フレームワーク、または 1 つの価格モデルに密接に結合されます。
上司と従業員のパターン
なぜそれが重要なのか:
オーケストレーション層は、ロックインを強化するのではなく、モデルを置き換え可能にする必要があります。
問題点: 多くの組織は AI 機能を必要としていますが、すべてのデータを外部サービスに送信することはできません。
PostgreSQL および pgvector のサポート
ローカル検索拡張生成
なぜそれが重要なのか:
プライバシー、ローカル制御、およびポータビリティは、企業の贅沢品ではなく、実用的なオプションである必要があります。
問題点: チームは、別のモデルが現在のプロバイダーを置き換えるのに十分であるかどうかを簡単に証明できません。
モデルのルーティングと移行に関する決定の証拠
なぜそれが重要なのか:
モデルは、切り替えが難しすぎると感じたからといって残るのではなく、測定されたパフォーマンスを通じてその地位を獲得する必要があります。
問題点: AI ログは一貫性がなく、不完全であり、多くの場合、プロバイダー固有のシステムにロックされています。
改ざん証拠のためのハッシュチェーン
AI システム全体での一貫した運用証拠
改ざん防止は不変を意味するものではない
署名だけでは現実世界の身元を証明するものではありません
監査ログは法的な否認防止を自動的に作成しません
フォーマットを使用しても、自動的に規制が遵守されるわけではありません
なぜそれが重要なのか:
組織は、ベンダーに依存せずに保持、検査、エクスポート、理解できる証拠を必要としています。
問題点: トークンの数や API の請求額だけでは、AI システムが効率的か有用かは説明できません。
正確な10進数演算によるコスト計算
不足している価格設定がゼロとして扱われることはありません
計算されたコストは請求書で確認されたコストではありません

予測された節約額は実際の節約額ではない
API の成功が自動的にビジネスの成功につながるわけではありません
予算は評価されるものであり、自動的に適用されるものではありません
なぜそれが重要なのか:
AI の経済性は、トークン単体ではなく、コストと許容可能な結果の観点から評価されるべきです。
コア ソフトウェアは、MIT ライセンスに基づいて無料でオープンソースです。
エコシステムは、モデル、プロバイダー、デプロイメント環境を置き換え可能に保つように設計されています。
プロジェクトは、実際に実証されない限り、テスト、統合、セキュリティ保証、コンプライアンス、公開ステータス、または検証を主張してはなりません。
組織は、AI インフラストラクチャの重要な部分をローカルで、または管理する環境内で実行できる必要があります。
モデルとプロバイダーは、測定された結果、コスト、遅延、プライバシー、運用上のニーズに基づいて選択する必要があります。
ガバナンスはシステムとともに機能する必要がある
ポリシー、組織的な意味、監査証拠、測定値は、モデルやプロバイダーが変更された場合でも移植可能な状態を維持する必要があります。
オープンインフラは公共財です
AI インフラストラクチャは、大企業のみが利用できるものであってはなりません。小規模な組織、公的機関、研究者、独立系開発者は、実用的な代替手段にアクセスできる必要があります。
AI アプリケーションを構築する開発者
赤にしようとするチーム

[切り捨てられた]

## Original Extract

Linux of AI is a free, open-source ecosystem for building portable, governed, measurable, and replaceable AI infrastructure. It connects tools for agent orchestration, private deployment, model switching, policy, ontology, audit logs, and AI cost measurement to help reduce vendor lock-in. - sekacorn
[truncated]

GitHub - sekacorn/Linux-of-Ai: Linux of AI is a free, open-source ecosystem for building portable, governed, measurable, and replaceable AI infrastructure. It connects tools for agent orchestration, private deployment, model switching, policy, ontology, audit logs, and AI cost measurement to help reduce vendor lock-in. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
sekacorn
/
Linux-of-Ai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits docs docs examples/ vendor-exit-demo examples/ vendor-exit-demo README.md README.md View all files Repository files navigation
Free, open, portable AI infrastructure for everyone.
Build AI systems that are inspectable, governed, measurable, replaceable, and not trapped inside a single vendor.
Mission •
The Problem •
Projects •
Architecture •
Principles •
Get Started
New to Linux of AI? Start with the Start Here guide to choose the right tool for your problem.
You can also inspect the Linux of AI Vendor Exit Demo to see the ecosystem pattern in one deterministic, offline-first example.
The Linux of AI Vendor Exit Demo shows how the ecosystem can help a team evaluate replacing an expensive or locked-in AI model using portable ontology, policy, benchmarking, audit evidence, and cost/outcome measurement.
This demo is deterministic and offline-first. It demonstrates the integration
pattern using portable files and reports, including AIMeter OSS-style
cost/outcome JSON and AIAuditLog-style audit-event JSONL exports from the model
replacement decision.
Project
PyPI package
Role
Status
AgentForge
agentforge-oss
Agent orchestration
Published
PrivateAIStack
privateaistack
Private/local deployment
Published
ModelSwapBench
modelswapbench
Model replacement and vendor-exit reports
Published; includes AI Vendor Exit Report
OpenOntologyLite
openontologylite
Portable ontology and business meaning
Published
AgentPolicyPack
agentpolicypack
Policy-as-code for agents
Published
AIAuditLog
aiauditlog
AI audit evidence
Published
AIMeter OSS
aimeter-oss
AI cost, usage, efficiency, and outcome measurement
Published
What This Is Not
a replacement for human review in high-risk workflows;
proof that any model migration is automatically safe.
It is open infrastructure for building, evaluating, governing, auditing, and measuring AI systems.
Run or inspect examples/vendor-exit-demo/
Try ModelSwapBench for model replacement decisions
Try AIMeter OSS for cost and outcome measurement
Add governance with AgentPolicyPack and audit evidence with AIAuditLog
Linux of AI is a seven-project open-source ecosystem created to reduce vendor lock-in and make practical AI infrastructure available to everyone.
No organization should be forced to surrender control of its AI systems, data, costs, policies, or future to a single vendor.
AI is becoming critical infrastructure. But many organizations are discovering that the systems they built are difficult to move, difficult to inspect, difficult to govern, and increasingly expensive to operate.
This ecosystem exists to provide another path.
A path where AI infrastructure is:
Portable across models, providers, and environments
Inspectable instead of hidden behind opaque services
Governed through explicit policies
Measurable in cost, usage, efficiency, and outcomes
Replaceable when a model or provider no longer serves the user
Local-first where privacy or control matters
Free and open source for public benefit
This software is intended to remain available to developers, researchers, nonprofits, businesses, governments, students, and communities without placing the core infrastructure behind a paywall.
Organizations adopting AI repeatedly encounter the same pain points.
Applications become tightly coupled to one model provider, SDK, API format, pricing model, or hosted platform. Replacing the provider later can require major rewrites.
A prototype may be affordable, but production usage can grow rapidly. Teams often lack clear controls for routing, budgeting, measuring cost, and comparing alternatives.
Organizations may know they are overpaying or underperforming, but they lack a repeatable way to test another model without disrupting the application.
AI agents can call tools, access data, and make decisions without clear, portable rules governing what they are allowed to do.
Logs are often inconsistent, provider-specific, incomplete, or difficult to verify after the fact.
API success does not necessarily mean business success. Many systems track requests and tokens but not whether the AI produced an acceptable outcome.
Loss of organizational meaning
Business concepts, relationships, permissions, and actions are frequently embedded directly inside application code, making systems difficult to understand and migrate.
Privacy and deployment constraints
Some organizations cannot send sensitive data to external providers. They need local or controlled deployment options that do not require rebuilding the entire stack.
Linux of AI addresses these problems as one connected ecosystem.
Layer
Project
What it solves
Repository
PyPI
Organizational meaning
OpenOntologyLite
Defines portable entities, relationships, actions, permissions, and business meaning outside application code
GitHub
PyPI
Governance
AgentPolicyPack
Provides portable policy-as-code for governing AI agents and agentic workflows
GitHub
PyPI
Orchestration
AgentForge
Coordinates multi-agent systems with model routing, budgets, governance, memory, and observability
GitHub
PyPI
Private deployment
PrivateAIStack
Provides a local-first starting point for private AI, local models, RAG, code review, auditing, and controlled deployment
GitHub
PyPI
Model replacement
ModelSwapBench
Tests whether another model can deliver an acceptable outcome at lower cost, latency, or operational risk
GitHub
PyPI
Operational evidence
AIAuditLog
Defines a portable, tamper-evident audit-event format and local toolkit for AI systems and agents
GitHub
PyPI
Cost and outcomes
AIMeter OSS
Measures AI usage, cost, efficiency, and business outcomes without assuming provider billing equals business value
GitHub
PyPI
Architecture
OpenOntologyLite
│
▼
Portable organizational meaning
│
▼
AgentPolicyPack
│
▼
Portable governance rules
│
▼
AgentForge
│
▼
Portable agent orchestration
│
▼
PrivateAIStack
│
▼
Private and controlled deployment
│
▼
ModelSwapBench
│
▼
Model replacement and economics verification
│
▼
AIAuditLog
│
▼
Portable tamper-evident operational evidence
│
▼
AIMeter OSS
│
▼
Usage, cost, efficiency, and outcome measurement
Each project can be used independently. Together, they form a portable foundation for building AI systems without making one provider the permanent center of the architecture.
Pain point: Organizational knowledge is buried inside source code, database schemas, prompts, and vendor-specific platforms.
Portable ontology definitions in YAML or JSON
Typed entities, properties, relationships, and actions
Validation and canonical representation
Diffing, inspection, documentation, and diagram export
Local-first and offline operation
Why it matters:
Your organization’s meaning should belong to your organization, not to a proprietary platform.
Pain point: Agent behavior is often governed by scattered prompt instructions and application-specific checks.
Explicit rules for agent actions
Reusable governance across systems
A foundation for reviewable and testable agent behavior
Why it matters:
Governance should be visible, portable, and separate from the model itself.
Pain point: Multi-agent systems become tightly coupled to one provider, one orchestration framework, or one pricing model.
Supervisor and worker patterns
Why it matters:
The orchestration layer should make models replaceable rather than make lock-in stronger.
Pain point: Many organizations need AI capabilities but cannot send all data to external services.
PostgreSQL and pgvector support
Local retrieval-augmented generation
Why it matters:
Privacy, local control, and portability should be practical options, not enterprise luxuries.
Pain point: Teams cannot easily prove whether another model is good enough to replace their current provider.
Evidence for model-routing and migration decisions
Why it matters:
A model should earn its place through measured performance, not remain because switching feels too difficult.
Pain point: AI logs are inconsistent, incomplete, and often locked into provider-specific systems.
Hash chaining for tamper evidence
Consistent operational evidence across AI systems
Tamper-evident does not mean immutable
A signature does not prove real-world identity by itself
Audit logs do not automatically create legal non-repudiation
Using the format does not automatically create regulatory compliance
Why it matters:
Organizations need evidence they can retain, inspect, export, and understand independently of a vendor.
Pain point: Token counts and API bills do not explain whether an AI system is efficient or useful.
Cost calculation using precise decimal arithmetic
Missing pricing is never treated as zero
Calculated cost is not invoice-confirmed cost
Projected savings are not realized savings
API success is not automatically business success
Budgets are evaluated, not automatically enforced
Why it matters:
AI economics should be measured in terms of cost and acceptable outcomes, not tokens alone.
The core software is free and open source under the MIT License.
The ecosystem is designed to keep models, providers, and deployment environments replaceable.
The projects should not claim tests, integrations, security guarantees, compliance, publication status, or verification unless those claims were actually demonstrated.
Organizations should be able to run critical parts of their AI infrastructure locally or inside environments they control.
Models and providers should be selected based on measured outcomes, cost, latency, privacy, and operational needs.
Governance should travel with the system
Policies, organizational meaning, audit evidence, and measurements should remain portable when the model or provider changes.
Open infrastructure is a public good
AI infrastructure should not be available only to the largest companies. Smaller organizations, public institutions, researchers, and independent developers should have access to practical alternatives.
Developers building AI applications
Teams trying to red

[truncated]
