---
source: "https://github.com/GetCassis/dbt-agent-readiness"
hn_url: "https://news.ycombinator.com/item?id=48488756"
title: "Show HN: A skill to audit your dbt project for what an AI agent will get wrong"
article_title: "GitHub - GetCassis/dbt-agent-readiness: Audit a dbt project for what an AI agent will get wrong if you point it at the data today. · GitHub"
author: "matthieu_bl"
captured_at: "2026-06-11T13:29:43Z"
capture_tool: "hn-digest"
hn_id: 48488756
score: 3
comments: 0
posted_at: "2026-06-11T11:06:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A skill to audit your dbt project for what an AI agent will get wrong

- HN: [48488756](https://news.ycombinator.com/item?id=48488756)
- Source: [github.com](https://github.com/GetCassis/dbt-agent-readiness)
- Score: 3
- Comments: 0
- Posted: 2026-06-11T11:06:28Z

## Translation

タイトル: HN を表示: AI エージェントが間違える可能性のある dbt プロジェクトを監査するスキル
記事のタイトル: GitHub - GetCassis/dbt-agent-readiness: dbt プロジェクトを監査して、AI エージェントに今日のデータを指定した場合に間違いを犯す可能性があるかどうかを確認します。 · GitHub
説明: dbt プロジェクトを監査して、今日データを指定した場合に AI エージェントが何を間違えるかを確認します。 - GetCassis/dbt-agent-readiness

記事本文:
GitHub - GetCassis/dbt-agent-readiness: dbt プロジェクトを監査して、今日データを指定した場合に AI エージェントが何を間違えるかを確認します。 · GitHub
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
アラートを閉じる
{{ メッセージ }}
ゲットカシス
/
dbt-agent-readiness
公共
通知
通知設定を変更するにはサインインする必要があります

ティング
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット 例 例 フェーズ フェーズ スクリプト スクリプト test-fixtures/messy-jaffle-shop test-fixtures/messy-jaffle-shop .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md SKILL.md SKILL.md TROUBLESHOOTING.md TROUBLESHOOTING.md report-template.md report-template.md 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
dbt プロジェクトを監査して、今日データに AI エージェントを向けた場合に AI エージェントが間違える可能性があるもの (間違ったメトリクス、間違ったテーブル、欠落した行、壊れた結合など) を監査するクロード コード スキル。
AI アナリスト、副操縦士、または内部データ エージェントを操縦する dbt チーム向けに構築されています。
ブロッカー (コード証拠、エージェントは今日攻撃します):
2 つのモデルは収益を表すと主張していますが、その計算方法は異なります
同じエンティティには、モデル全体で customer_id 、 cust_id 、および user_id という名前が付けられます
SQL が実際に出力しない YAML で宣言された列 (エージェントは存在しない列を SELECT します)
モデルの説明は完全性を保証しますが、SQL は行をフィルターします
1 つのentity_id 列がモデル全体の異なるエンティティを参照します
説明では COUNT と表示されますが、SQL では SUM が実行されます
存在しないモデルへの ref() (クエリはコンパイル時に失敗します)
単位ドリフト (EUR / ユーロセント、Wh / kWh)
モデル内コンセプトの衝突 (deployment_start_date +zone_deployment_start_date)
衛生 (リスク要因、それぞれに SQL クエリが同梱されており、実行して検証できます):
PK / not-null / 関係 / accept_values テストの欠落
説明にもカーディナリティが記載されていない場合、グレインは宣言されていません
マニフェストが見つからない場合に dbt コンパイルのフラグが立てられるマクロ使用モデル
バンドルされたサンプル レポートのブロッカー:
### 1. 同じエンティティはモデル間で 3 つの異なる方法で名前が付けられます
エージェントが得られるもの

間違い: 「顧客 X は何件注文しましたか?」と尋ねると、
エージェントがorders.customer_idをcustomers.customer_idに結合したが、ミスした
列が cust_id である fct_revenue の行。
証拠:catalogs.concept_variants クラスターの customer_id には別個のものがあります
名前 ['cust_id'、'customer_id'、'user_id']。
models/marts/fct_revenue.sql:13 は o.cust_id を出力します。
models/marts/customers.sql:11 は、o.cust_id を customer_id としてエイリアスします。
修正: 名前を 1 つの正規形式 (customer_id) に変更します。
作業時間: 午後。
バンドルされたテスト フィクスチャに対して生成された完全なサンプル レポートを参照してください。
Claude Code スキル ディレクトリにクローンを作成します。
git clone https://github.com/GetCassis/dbt-agent-readiness ~ /.claude/skills/dbt-agent-readiness
Python の依存関係をインストールします。
pip install -r ~ /.claude/skills/dbt-agent-readiness/requirements.txt
要件:
Claude Code (スキルをサポートする最新バージョン)
pyyaml および sqlglot (requirements.txt 経由でインストール)
推奨: スキルがマクロ ( dbt_utils.star 、 SELECT * ) を解決できるように、監査の前にターゲット プロジェクトで dbt コンパイルを実行します。これがないと、マクロを使用するモデルでのファントムカラムの結果は表示されずに抑制されます。
/path/to/dbt/project で dbt-agent-readiness スキルを実行します。
レポートは {project_path}/dbt-agent-readiness.md に書き込まれます。
準備状況の判定: 準備ができている / 準備ができていない / 安全ではない、準備完了までの距離。
ブロッカー : エージェントが今日発生するであろう証拠に裏付けられた障害。それぞれに影響を受けるモデル、爆発範囲、修正、工数の見積もりが含まれます。
衛生: それぞれのリスク要因には、促進または却下するために実行できる検証クエリが含まれています。
安全な開始境界: エージェントが今日安全にクエリできるモデル。
修復バックログ: 修正の優先リスト。
31 ～ 200 のモデル: 3 ～ 4 つの並列サブエージェント (フラグ駆動レビュー + モデルごとのディープ パス)
>200 モデル: サブエージェントを生成する前のチェックポイント

スキルは、dbt プロジェクト ファイル (SQL、YAML、説明、オプションで target/manifest.json ) をローカルに読み取ります。セッション内で実行されている Claude Code は、通常の操作の一部としてそのコンテンツを Anthropic に送信します。スキル自体はネットワーク接続を開いたり、ウェアハウスにクエリを実行したり、ウェアハウスの認証情報を必要としません。単一のレポート ファイルを {project_path}/dbt-agent-readiness.md に書き込み、プロジェクト内のそれ以外は何も変更しません。
実行時のデータ品質 (NULL 率、鮮度、行数)。衛生用品には、倉庫に対して実行できる検証クエリが含まれています。
ソースシステムが変更されます。ランタイム監視のみが上流のフォーマットドリフトを捕捉します。
結合がビジネス上意味があるかどうか。監査では、ドメインの有効性ではなく、構造が確認されます。
クエリのパターンと使用頻度。クエリログが必要になります。
BI ツールの指標の競合。 Looker / Tableau のエクスポート アクセスが必要です。
リリースごとにタグ付けされています。 CHANGELOG.md を参照してください。インベントリの JSON スキーマはメジャー バージョン間で異なる場合があります。ピンで留めないでください。
SKILL.md すべてのステップをルーティングし、サブエージェントを生成するオーケストレーター
report-template.md 出力レポートテンプレート
フェーズ/フェーズ サブエージェント プロンプト
scripts/inventory.py 確定的インベントリ (SQL、コンセプト、カタログ)
scripts/dispatch_prep.py レビュー - パケット生成、重要度スコアリング
例/監査レポートの例
test-fixture/ 手動煙テスト用のテスト プロジェクト
CHANGELOG.md バージョン履歴
TROUBLESHOOTING.md 一般的な問題と修正
ライセンスミット
トラブルシューティング
一般的な障害モードについては、TROUBLESHOOTING.md を参照してください。
Cassis のチームによって作られました。
dbt プロジェクトを監査して、AI エージェントに今日のデータを向けた場合に間違いを犯す可能性があるかどうかを確認します。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
1
v1.0.0 — 初期公開

リリース
最新の
2026 年 4 月 21 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Audit a dbt project for what an AI agent will get wrong if you point it at the data today. - GetCassis/dbt-agent-readiness

GitHub - GetCassis/dbt-agent-readiness: Audit a dbt project for what an AI agent will get wrong if you point it at the data today. · GitHub
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
GetCassis
/
dbt-agent-readiness
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits examples examples phases phases scripts scripts test-fixtures/ messy-jaffle-shop test-fixtures/ messy-jaffle-shop .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md SKILL.md SKILL.md TROUBLESHOOTING.md TROUBLESHOOTING.md report-template.md report-template.md requirements.txt requirements.txt View all files Repository files navigation
A Claude Code skill that audits a dbt project for what an AI agent will get wrong if you point it at the data today: wrong metric, wrong table, missed rows, broken joins.
Built for dbt teams piloting AI analysts, copilots, or internal data agents.
Blockers (code evidence, agent will hit today):
Two models claim to represent revenue but calculate it differently
The same entity is named customer_id , cust_id , and user_id across models
A YAML-declared column the SQL doesn't actually emit (agent SELECTs a column that doesn't exist)
A model description promises totality but the SQL filters rows
One entity_id column refers to different entities across models
Description says COUNT, SQL does SUM
ref() to a model that doesn't exist (queries fail at compile)
Unit drift (EUR / EUR cents, Wh / kWh)
Within-model concept collision ( deployment_start_date + zone_deployment_start_date )
Hygiene (risk factors, each shipped with a SQL query you can run to verify):
Missing PK / not-null / relationship / accepted_values tests
Grain undeclared when the description is also silent on cardinality
Macro-using models flagged for dbt compile when the manifest is missing
A Blocker from the bundled sample report:
### 1. The same entity is named three different ways across models
What the agent gets wrong: Asked "how many orders did customer X place?",
the agent joins orders.customer_id to customers.customer_id and misses
the rows in fct_revenue where the column is cust_id.
Evidence: catalogs.concept_variants cluster customer_id has distinct
names ['cust_id', 'customer_id', 'user_id'].
models/marts/fct_revenue.sql:13 emits o.cust_id.
models/marts/customers.sql:11 aliases o.cust_id as customer_id.
Fix: Rename to one canonical form (customer_id).
Effort: afternoon.
See the full sample report generated against the bundled test fixture.
Clone into your Claude Code skills directory:
git clone https://github.com/GetCassis/dbt-agent-readiness ~ /.claude/skills/dbt-agent-readiness
Install Python dependencies:
pip install -r ~ /.claude/skills/dbt-agent-readiness/requirements.txt
Requirements:
Claude Code (any recent version with Skills support)
pyyaml and sqlglot (installed via requirements.txt )
Recommended: run dbt compile in the target project before auditing so the skill can resolve macros ( dbt_utils.star , SELECT * ). Without it, phantom-column findings on macro-using models are suppressed rather than emitted.
Run the dbt-agent-readiness skill on /path/to/dbt/project
The report is written to {project_path}/dbt-agent-readiness.md .
Readiness verdict : ready / not ready / unsafe, with distance to ready.
Blockers : evidence-backed failures an agent will hit today, each with affected models, blast radius, fix, and effort estimate.
Hygiene : risk factors, each with a verification query you can run to promote or dismiss.
Safe starting perimeter : which models an agent can query safely today.
Remediation backlog : prioritized list of fixes.
31 to 200 models: 3 to 4 parallel subagents (flag-driven review + per-model deep pass)
>200 models: checkpoint before spawning subagents
The skill reads your dbt project files locally (SQL, YAML, descriptions, optionally target/manifest.json ). Claude Code, running in your session, sends that content to Anthropic as part of its normal operation. The skill itself does not open network connections, does not query your warehouse, and does not require warehouse credentials. It writes a single report file to {project_path}/dbt-agent-readiness.md and modifies nothing else in your project.
Runtime data quality (null rates, freshness, row counts). Hygiene items carry verification queries you can run against the warehouse.
Source system changes. Only runtime monitoring catches upstream format drift.
Whether a join makes business sense. The audit sees structure, not domain validity.
Query patterns and usage frequency. Would require query logs.
BI tool metric conflicts. Would require Looker / Tableau export access.
Tagged per release. See CHANGELOG.md . The inventory JSON schema may change between major versions; don't pin against it.
SKILL.md Orchestrator that routes all steps and spawns subagents
report-template.md Output report template
phases/ Phase subagent prompts
scripts/inventory.py Deterministic inventory (SQL, concepts, catalogs)
scripts/dispatch_prep.py Review-packet generation, importance scoring
examples/ Example audit reports
test-fixtures/ Test projects for manual smoke testing
CHANGELOG.md Version history
TROUBLESHOOTING.md Common issues and fixes
LICENSE MIT
Troubleshooting
See TROUBLESHOOTING.md for common failure modes.
Made by the team behind Cassis .
Audit a dbt project for what an AI agent will get wrong if you point it at the data today.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
v1.0.0 — initial public release
Latest
Apr 21, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
