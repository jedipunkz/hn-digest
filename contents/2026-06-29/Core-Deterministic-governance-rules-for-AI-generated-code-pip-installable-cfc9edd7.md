---
source: "https://github.com/DariuszNewecki/CORE"
hn_url: "https://news.ycombinator.com/item?id=48724982"
title: "Core – Deterministic governance rules for AI-generated code (pip installable)"
article_title: "GitHub - DariuszNewecki/CORE: CORE is a governance runtime for autonomous AI systems. It enforces constitutional rules during execution, prevents governance bypass, and creates auditable authority chains for agent actions across operational domains. · GitHub"
author: "d_newecki"
captured_at: "2026-06-29T21:31:01Z"
capture_tool: "hn-digest"
hn_id: 48724982
score: 2
comments: 0
posted_at: "2026-06-29T20:45:54Z"
tags:
  - hacker-news
  - translated
---

# Core – Deterministic governance rules for AI-generated code (pip installable)

- HN: [48724982](https://news.ycombinator.com/item?id=48724982)
- Source: [github.com](https://github.com/DariuszNewecki/CORE)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T20:45:54Z

## Translation

タイトル: コア – AI 生成コードの決定論的ガバナンス ルール (pip インストール可能)
記事のタイトル: GitHub - DariuszNewecki/CORE: CORE は自律 AI システム用のガバナンス ランタイムです。実行中に憲法上の規則を強制し、ガバナンスのバイパスを防止し、運用ドメイン全体にわたるエージェントのアクションに対する監査可能な権限チェーンを作成します。 · GitHub
説明: CORE は、自律 AI システム用のガバナンス ランタイムです。実行中に憲法上の規則を強制し、ガバナンスのバイパスを防止し、運用ドメイン全体にわたるエージェントのアクションに対する監査可能な権限チェーンを作成します。 - DariuszNewecki/CORE

記事本文:
GitHub - DariuszNewecki/CORE: CORE は自律 AI システム用のガバナンス ランタイムです。実行中に憲法上の規則を強制し、ガバナンスのバイパスを防止し、運用ドメイン全体にわたるエージェントのアクションに対する監査可能な権限チェーンを作成します。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
を却下する

警告します
{{ メッセージ }}
ダリウシュネツキ
/
コア
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2,016 コミット 2,016 コミット .docker/ core-engine .docker/ core-engine .github .github .intent .intent .specs .specs docs docs example/ starter-intent Example/ starter-intent grc-catalogs grc-catalogs フック フック インフラ インフラ スクリプト スクリプト src src テスト テスト var var web web .env.example .env.example .env.test .env.test .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md action.yml action.yml docker-compose.yml docker-compose.ymlentrypoint.shentrypoint.sh install-core.sh install-core.sh mkdocs.yml mkdocs.yml詩.ロック 詩.ロック pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI支援ソフトウェア開発のための実行可能な憲法ガバナンス。
AI アクションのトレーサビリティがオプションではない環境向けに設計されています。
バージョン管理: API が安定に近づいていることを示す 2.x の SemVer (PyPI のベータ版)。バージョン管理ポリシーを参照してください。
AI コーディング ツールはコードを高速に生成します。正気を保つには速すぎる。
強制力がなければ、AI 支援コードベースには、レイヤー違反、アーキテクチャ契約の違反、際限なく増大するファイルなど、目に見えない負債が蓄積されます。そして、エージェントは何の制約も受けずに放置すると、最終的に次のようなことを行うようになります。
エージェント: 「このバグを修正するために実稼働データベースを削除します。」
システム: 実行します。
あなた: 😱
CORE がその CL を作成します

違反はあり得ない — 実行前に構造的にブロックされ、事後は検出されない。 (ハードブロックと勧告レポートのどちらのサーフェスが、以下の現在のプルーフ ステータスにマッピングされています。)
エージェント: 「このバグを修正するために実稼働データベースを削除します。」
憲法: ブロック済み — data.ssot.database_primacy に違反します
システム: 実行が停止されました。違反が記録されました。
あなた: 😌
CORE は、機械によって強制される憲法によって AI エージェントを制約するガバナンス ランタイムです。アーキテクチャ上の不変条件を強制し、無効な突然変異を自動的にブロックし、自律的なワークフローを監査可能かつ決定論的にします。
LLM は CORE 内で動作します。決してそれを超えることはありません。
🎬 たった 1 つのコマンドでそれ自体が統治されるのを見てください
これを信じて受け入れる必要はありません。 Docker を備えたクリーンなマシン上:
git クローン https://github.com/DariuszNewecki/CORE.git
cd CORE && ./install-core.sh
install-core.sh は CORE を起動し、結果チェーンのデモをライブで実行して終了します。LLM キーは必要ありません。それ：
linkage.assign_ids (ブロックルール) に違反する関数をコミットします。
CORE に修正を提案させ、ガバナーがそれを承認し、CORE がそれを実行して修復をコミットします。
記録された完全な因果関係チェーンを出力します: 発見→提案→承認→実行→ファイル変更
scripts/demo.sh でいつでも再実行できます。
強制されたあらゆる行動はその系譜を記録します。 CORE データベースからライブで取得された 2 つの結果チェーン — 同じスキーマ、2 つの異なる権限:
自律パス — 安全としてリスク分類され、システムが自己承認
検索 → workflow.ruff_format_check src/api/cli/client.py 2026-05-18 05:15:15 UTC
提案 → 8845dacc… fix.format 2026-05-18 05:16:15 UTC
承認 → リスク分類.安全自動承認 2026-05-18 05:16:15 UTC
実行 → 完了 (1.29秒) 2026-05-18 05:17:18 UTC
ファイル変更 → +105 / -0 98da9038 →

fca9a971 src/api/cli/client.py 2026-05-18 05:17:19 UTC
人間の承認パス — ループ内のガバナー
検索 → purity.docstrings.required src/cli/commands/audit_reporter.py 2026-05-15 08:28:29 UTC
提案 → a4363a81… fix.docstrings 2026-05-16 13:39:34 UTC
承認 → 校長.知事 (cli_admin) 2026-05-16 13:53:32 UTC
実行 → 完了 (24.5秒) 2026-05-16 13:55:48 UTC
ファイル変更 → +26 / -0 5a123426 → 71fde489 src/cli/commands/audit_reporter.py 2026-05-16 13:55:49 UTC
どちらのチェーンも、proposal_consequences と Blackboard_entries からエンドツーエンドでクエリ可能です。どの権限が適用されるかは憲法によって決定されます。スキーマは同一です。 Proof Index の Consequence-chain クエリを使用して、それらを自分で再現します。
CORE は、4 つのリポジトリ層 (憲法として施行される 3 層) と仕様 (人間の意図) にわたって責任を分離します。この分離は慣習ではなく法律として施行されます。
📐 仕様 — 人間の意図 ( .specs/ )
システムの目的と決定が行われた理由を人間が定義する場所。アーキテクチャ文書、ノーススター文書、ユーザー要件、アーキテクチャ決定記録、および計画文書が含まれます。これは、実装を読む前に CORE を理解しようとする人にとってのエントリ ポイントです。
.specs/ は人間によって読み取られ、CORE のセマンティック レイヤーによって検索可能です。 CORE 自体によって書き込まれることはありません。
🧠 心 — 法 ( .intent/ + src/mind/ )
何が許可されるか、何が要求されるか、何が禁止されるかを定義します。機械可読な憲法規則、施行マッピング、フェーズを認識したガバナンス モデル、権限階層 (メタ → 憲法 → ポリシー → コード) が含まれています。
マインドは決して実行しない。心は決して変異しません。心は法を定義します。
⚖️ 意志 — 判決 ( src/will/ )
憲法上の制約を読み取り、自律的な推論を調整し、記録する

すべての意思決定を追跡可能な監査証跡で記録します。すべての操作は構造化されたフェーズ パイプラインに従います。
解釈→計画→生成→検証→スタイルチェック→実行
意志は決して肉体を迂回することはありません。決してマインドを書き換えることはありません。
🏗️ 本文 — 実行 ( src/body/ )
決定的でアトミックなコンポーネント: アナライザー、エバリュエーター、ファイル操作、git サービス、テスト ランナー、CLI コマンド。
体は突然変異を起こします。体は判断しません。体は支配しません。
すべての自律的な運用は、同じ憲法上のループによって管理されます。
フローチャート TD
A["🟢 目標\n人間の意図"] --> B["📂 コンテキスト\nリポジトリの状態 • 知識 • 歴史"]
B --> C["🔒 制約\n不変ルール\n215 ルール • 15 エンジン"]
C --> D["🗺️ 計画\n段階的な推論\nルールを意識した計画"]
D --> E["✨ GENERATE\nコード • 変更 • ツール呼び出し"]
E --> F["✅ VALIDATE\n決定論的チェック\nAST • セマンティック • インテント • スタイル"]
F -->|パス| G["▶️ 実行\n準拠した変更を適用"]
F -->|失敗| H["🔄 修復\n違反を修復\n自律性ラダー"]
H --> E
G --> I["✓ 成功\n変更がコミットされました"]
サブグラフ「安全停止」
TB方向
J["🚨 憲法違反\n→ 強制停止\n+ 完全な監査ログ"]
終わり
E -.->|違反がある| J
F -.->|違反がある| J
classDef フェーズ fill:#f8f9fa、ストローク:#495057、ストローク幅:2px
classDef 制約の塗りつぶし:#d1e7ff、ストローク:#0d6efd、ストローク幅:2.5px
classDef 検証 fill:#fff3cd、ストローク:#ffc107、ストローク幅:2.5px
classDef 停止 fill:#ffebee、ストローク:#dc3545、ストローク幅:3px
クラスA、B、D、E、G、Iフェーズ
クラスC制約
クラス F 検証
クラスJ停止
読み込み中
システム保証
自律レーン外のファイルは変更できません
構造上のルールを黙って回避することはできません
アトミック アクションは、管理対象エグゼキュータの外部で実行できません (インライン認可は監査→結果ループに延期されます)。

決定はフェーズを認識し、決定トレースで記録されます (監査の永続性はベストエフォート型です。「現在の証明ステータス」を参照)。
いかなる代理人も憲法を改正することはできない
ブロック ルールが失敗すると、実行は部分的な状態で停止します。レポートおよび勧告ルールによって結果が明らかになり、続行します。どのようなブロックとどのようなレポートがレポートされるかはモードによって異なります。
CORE の保証セマンティクスは、設計によりモード間で分割されています。これは各サーフェスの機能を示す正直なマップであるため、単一の 2 つの主張 (「コアが違反をブロックする」) が全体像と誤解されることはありません。
ハード不変条件と憲法上の規則は無条件にブロックします。ポリシー、機能、およびステートレス層は設計上弱く、境界が暗示されるのではなく判読できるようにここにラベルが付けられています。
原始的
目的
文書
永続化され検証されたアーティファクト
ルール
原子的な規範的声明
フェーズ
ルールが評価されるとき
権威
誰がそれを定義または修正できるか
執行の強み: 阻止、報告、助言
エンジン
方法
アストゲート
決定論的構造解析 (AST)
正規表現ゲート
パターンベースのテキスト強制
グロブゲート
パスと境界の強制
cli_gate
CLI サーフェスとコマンド形状の強制
アーティファクトゲート
宣言されたアーティファクトの完全性と発見されたアーティファクトの完全性
ワークフローゲート
フェーズシーケンスとカバレッジチェック
ナレッジゲート
責任と所有権の検証
アクションゲート
アトミックアクションの不変条件
パッシブゲート
サブストレート適用ルール (DB/ランタイム マーカー)
タクソノミーゲート
機能 ID ↔ アトミックアクションの一貫性 (ADR-079 D9)
コントラクトゲート
横断的なデータ契約の一貫性 (コンテキストレベル; ADR-102)
llm_gate
LLM 支援のセマンティック チェック
インテントガード *
実行時書き込み許可 (監査ではない)
* .specs/papers/CORE-Gate.md ごとのランタイム ゲート。視認性を高めるためにここに保持されます。
可能であれば決定論的。 LLM は必要な場合にのみ使用します。
51 の 215 のルール

ルール文書。 209 は施行エンジンにマッピングされます。 6 つのテスト品質ルールはマッピングがまだ保留中です。 「マップされた」とは、エンジンに依存することを意味します。すべてのモードで強制されるわけではありません。ステートレス CI は、ナレッジ グラフと LLM プロバイダーを必要とするKnowledge_gate と llm_gate をスキップします。
CORE は、定義されたレベルを通じて進行します。それぞれが憲法上の制限を維持しながら機能を追加します。
A0 — 自己認識 ✅ 自分が何であるか、どこに住んでいるかを知っています。
A1 — 自己修復 ✅ 既知の構造的問題を自動的に修正します
A2 — 管理された世代 ✅ 自然言語 → 憲法に準拠したコード
A3 — 統治された自律性 ✅ デーモンは無人で違反を検出、提案、修正します ← 現在
A4 — 自己複製 🔮 自身の理解に基づいて CORE.NG を書き込みます
要件
依存関係
バージョン
パイソン
3.12+
PostgreSQL
≥ 14
クドラント
最新の
ドッカー
サービス用
詩
デプス用
クイックスタート
正直なステータス — 今日機能するもの。 CORE はそれ自体をエンドツーエンドで管理し (上記のデモ)、.intent/ 構成を持つリポジトリを監査します。CI では GitHub Action を使用して監査するか、そのリポジトリ内で core-admin code Audit --offline を使用してローカルに監査します。 pip install core-runtime を使用すると、core-admin CLI が提供されます。
独自のリポジトリを管理します (BYOR): pip install core-runtime と 2 つのコマンドが適合した構成をブートストラップします。 (1) <path> 上の core-admin プロジェクト --write は、マシン フロアを提供します (LLM やデータベースはありません)。 (2) コア管理者プロジェクトのスカウト <pa

[切り捨てられた]

## Original Extract

CORE is a governance runtime for autonomous AI systems. It enforces constitutional rules during execution, prevents governance bypass, and creates auditable authority chains for agent actions across operational domains. - DariuszNewecki/CORE

GitHub - DariuszNewecki/CORE: CORE is a governance runtime for autonomous AI systems. It enforces constitutional rules during execution, prevents governance bypass, and creates auditable authority chains for agent actions across operational domains. · GitHub
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
DariuszNewecki
/
CORE
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2,016 Commits 2,016 Commits .docker/ core-engine .docker/ core-engine .github .github .intent .intent .specs .specs docs docs examples/ starter-intent examples/ starter-intent grc-catalogs grc-catalogs hooks hooks infra infra scripts scripts src src tests tests var var web web .env.example .env.example .env.test .env.test .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md action.yml action.yml docker-compose.yml docker-compose.yml entrypoint.sh entrypoint.sh install-core.sh install-core.sh mkdocs.yml mkdocs.yml poetry.lock poetry.lock pyproject.toml pyproject.toml View all files Repository files navigation
Executable constitutional governance for AI-assisted software development.
Designed for environments where AI action traceability is not optional.
Versioning: SemVer with 2.x denoting an API approaching stability (Beta on PyPI); see the versioning policy .
AI coding tools generate code fast. Too fast to stay sane.
Without enforcement, AI-assisted codebases accumulate invisible debt — layer violations, broken architectural contracts, files that grow unbounded. And agents, left unconstrained, will eventually do something like this:
Agent: "I'll delete the production database to fix this bug"
System: Executes.
You: 😱
CORE makes that class of violation impossible — structurally blocked before execution, not detected after the fact. (Which surfaces hard-block versus advisory-report is mapped under Current proof status below.)
Agent: "I'll delete the production database to fix this bug"
Constitution: BLOCKED — Violates data.ssot.database_primacy
System: Execution halted. Violation logged.
You: 😌
CORE is a governance runtime that constrains AI agents with machine-enforced constitutional law — enforcing architectural invariants, blocking invalid mutations automatically, and making autonomous workflows auditable and deterministic.
LLMs operate inside CORE. Never above it.
🎬 See it govern itself — one command
You don't have to take this on faith. On a clean machine with Docker:
git clone https://github.com/DariuszNewecki/CORE.git
cd CORE && ./install-core.sh
install-core.sh stands up CORE and finishes by running the consequence-chain demo live — no LLM key required. It:
commits a function that violates linkage.assign_ids (a blocking rule)
has CORE propose a fix, the governor approve it, and CORE execute it and commit the repair
prints the full causal chain it recorded: finding → proposal → approval → execution → file change
Re-run it any time with scripts/demo.sh .
Every enforced action records its lineage. Two consequence chains, pulled live from the CORE database — same schema, two different authorities:
Autonomous path — risk-classified as safe, system self-approved
FINDING → workflow.ruff_format_check src/api/cli/client.py 2026-05-18 05:15:15 UTC
PROPOSAL → 8845dacc… fix.format 2026-05-18 05:16:15 UTC
APPROVAL → risk_classification.safe_auto_approval 2026-05-18 05:16:15 UTC
EXECUTION → completed (1.29s) 2026-05-18 05:17:18 UTC
FILE CHANGE → +105 / -0 98da9038 → fca9a971 src/api/cli/client.py 2026-05-18 05:17:19 UTC
Human-approval path — governor in the loop
FINDING → purity.docstrings.required src/cli/commands/audit_reporter.py 2026-05-15 08:28:29 UTC
PROPOSAL → a4363a81… fix.docstrings 2026-05-16 13:39:34 UTC
APPROVAL → principal.governor (cli_admin) 2026-05-16 13:53:32 UTC
EXECUTION → completed (24.5s) 2026-05-16 13:55:48 UTC
FILE CHANGE → +26 / -0 5a123426 → 71fde489 src/cli/commands/audit_reporter.py 2026-05-16 13:55:49 UTC
Both chains are queryable end-to-end from proposal_consequences and blackboard_entries . The constitution decides which authority applies; the schema is identical. Reproduce them yourself with the Consequence-chain query in the Proof Index.
CORE separates responsibility across four repository layers — three enforced as constitutional law, and Specs (human intent). This separation is enforced as law — not convention.
📐 Specs — Human Intent ( .specs/ )
Where humans define what the system is for and why decisions were made. Contains architectural papers, northstar documents, user requirements, architectural decision records, and planning documents. This is the entry point for anyone trying to understand CORE before reading its implementation.
.specs/ is read by humans and searchable by CORE's semantic layer. It is never written by CORE itself.
🧠 Mind — Law ( .intent/ + src/mind/ )
Defines what is allowed, required, or forbidden. Contains machine-readable constitutional rules, enforcement mappings, phase-aware governance models, and the authority hierarchy ( Meta → Constitution → Policy → Code ).
Mind never executes. Mind never mutates. Mind defines law.
⚖️ Will — Judgment ( src/will/ )
Reads constitutional constraints, orchestrates autonomous reasoning, and records every decision with a traceable audit trail. Every operation follows a structured phase pipeline:
INTERPRET → PLAN → GENERATE → VALIDATE → STYLE CHECK → EXECUTE
Will never bypasses Body. Will never rewrites Mind.
🏗️ Body — Execution ( src/body/ )
Deterministic, atomic components: analyzers, evaluators, file operations, git services, test runners, CLI commands.
Body performs mutations. Body does not judge. Body does not govern.
Every autonomous operation is governed by the same constitutional loop:
flowchart TD
A["🟢 GOAL\nHUMAN INTENT"] --> B["📂 CONTEXT\nRepo state • knowledge • history"]
B --> C["🔒 CONSTRAINTS\nImmutable rules\n215 rules • 15 engines"]
C --> D["🗺️ PLAN\nStep-by-step reasoning\nRule-aware plan"]
D --> E["✨ GENERATE\nCode • changes • tool calls"]
E --> F["✅ VALIDATE\nDeterministic checks\nAST • semantic • intent • style"]
F -->|Pass| G["▶️ EXECUTE\nApply compliant changes"]
F -->|Fail| H["🔄 REMEDIATE\nRepair violation\nAutonomy Ladder"]
H --> E
G --> I["✓ SUCCESS\nChanges committed"]
subgraph "SAFETY HALT"
direction TB
J["🚨 CONSTITUTIONAL VIOLATION\n→ HARD HALT\n+ FULL AUDIT LOG"]
end
E -.->|Any violation| J
F -.->|Any violation| J
classDef phase fill:#f8f9fa,stroke:#495057,stroke-width:2px
classDef constraint fill:#d1e7ff,stroke:#0d6efd,stroke-width:2.5px
classDef validate fill:#fff3cd,stroke:#ffc107,stroke-width:2.5px
classDef halt fill:#ffebee,stroke:#dc3545,stroke-width:3px
class A,B,D,E,G,I phase
class C constraint
class F validate
class J halt
Loading
System Guarantees
No file outside an autonomy lane can be modified
No structural rule can be bypassed silently
No atomic action can execute outside the governed executor (inline authorization is deferred to the audit→consequence loop)
Decisions are phase-aware and logged with decision traces (audit persistence is best-effort — see Current proof status )
No agent can amend constitutional law
If a blocking rule fails, execution halts with no partial state. Reporting and advisory rules surface findings and continue — what blocks versus what reports depends on the mode.
CORE's guarantee semantics are split across modes by design. This is the honest map of what each surface does, so a single binary claim ("CORE blocks violations") is not mistaken for the whole picture:
The hard invariants and constitutional rules block unconditionally; the policy, capability, and stateless tiers are weaker by design and labelled here so the boundary is legible rather than implied.
Primitive
Purpose
Document
Persisted, validated artifact
Rule
Atomic normative statement
Phase
When the rule is evaluated
Authority
Who may define or amend it
Enforcement strengths: Blocking · Reporting · Advisory
Engine
Method
ast_gate
Deterministic structural analysis (AST)
regex_gate
Pattern-based text enforcement
glob_gate
Path and boundary enforcement
cli_gate
CLI surface and command-shape enforcement
artifact_gate
Declared-vs-discovered artifact completeness
workflow_gate
Phase-sequencing and coverage checks
knowledge_gate
Responsibility and ownership validation
action_gate
Atomic-action invariants
passive_gate
Substrate-enforced rules (DB/runtime marker)
taxonomy_gate
Capability-id ↔ atomic-action coherence (ADR-079 D9)
contracts_gate
Cross-cutting data-contract coherence (context-level; ADR-102)
llm_gate
LLM-assisted semantic checks
IntentGuard *
Runtime write authorization (not audit)
*Runtime Gate per .specs/papers/CORE-Gate.md , kept here for visibility.
Deterministic when possible. LLM only when necessary.
215 rules across 51 rule documents. 209 are mapped to enforcement engines; 6 test-quality rules are still pending mappings. "Mapped" means engine-bound — not enforced in every mode: stateless CI skips knowledge_gate and llm_gate , which need the knowledge graph and an LLM provider.
CORE progresses through defined levels. Each adds capability while remaining constitutionally bounded.
A0 — Self-Awareness ✅ Knows what it is and where it lives
A1 — Self-Healing ✅ Fixes known structural issues automatically
A2 — Governed Generation ✅ Natural language → constitutionally aligned code
A3 — Governed Autonomy ✅ Daemon finds, proposes, and fixes violations unattended ← current
A4 — Self-Replication 🔮 Writes CORE.NG from its own understanding of itself
Requirements
Dependency
Version
Python
3.12+
PostgreSQL
≥ 14
Qdrant
latest
Docker
for services
Poetry
for deps
Quick Start
Honest status — what works today. CORE governs itself end to end (the demo above), and audits any repo that has a .intent/ constitution — in CI via the GitHub Action , or locally with core-admin code audit --offline inside that repo . pip install core-runtime gives you the core-admin CLI.
Govern your own repo (BYOR): pip install core-runtime and two commands bootstrap a fitted constitution. (1) core-admin project onboard <path> --write delivers the machinery floor (no LLM, no database). (2) core-admin project scout <pa

[truncated]
