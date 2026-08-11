---
source: "https://github.com/dbward-dev/dbward"
hn_url: "https://news.ycombinator.com/item?id=49257957"
title: "Show HN: dbward – Approval workflows for production DBs, built for AI agents"
article_title: "GitHub - dbward-dev/dbward: Context-aware approval gateway for production database operations. AI, CLI, and CI request. Humans approve. Agents execute. · GitHub"
author: "metapox"
captured_at: "2026-08-11T14:14:26Z"
capture_tool: "hn-digest"
hn_id: 49257957
score: 2
comments: 0
posted_at: "2026-08-11T13:26:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: dbward – Approval workflows for production DBs, built for AI agents

- HN: [49257957](https://news.ycombinator.com/item?id=49257957)
- Source: [github.com](https://github.com/dbward-dev/dbward)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T13:26:59Z

## Translation

タイトル: HN を表示: dbward – AI エージェント向けに構築された、本番 DB の承認ワークフロー
記事のタイトル: GitHub - dbward-dev/dbward: 運用データベース操作のためのコンテキスト認識型承認ゲートウェイ。 AI、CLI、CI リクエスト。人間は承認します。エージェントが実行します。 · GitHub
説明: 運用データベース操作のためのコンテキスト認識型承認ゲートウェイ。 AI、CLI、CI リクエスト。人間は承認します。エージェントが実行します。 - dbward-dev/dbward

記事本文:
GitHub - dbward-dev/dbward: 運用データベース操作のためのコンテキスト認識型承認ゲートウェイ。 AI、CLI、CI リクエスト。人間は承認します。エージェントが実行します。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
dbward-dev
/
ドブワード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
699 コミット 699 コミット .githooks .githooks .github .github コマーシャル コマーシャル

cial crates crates デモ デモ デプロイ デプロイ dev dev ドキュメント ドキュメントの例 例 .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス LICENSE-APACHE LICENSE-APACHE LICENSE-COMMERCIAL LICENSE-COMMERCIAL README.md README.md SECURITY.md SECURITY.mddeny.tomldeny.toml install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープンコア プロジェクト — コア コンポーネントは Apache-2.0 です。一部の機能とビルド済みバイナリには、dbward 商用ライセンスに基づくコードが含まれています。詳細については、「ライセンス」を参照してください。
実稼働データベースの承認ワークフローと監査ログ。
事故が生産に影響を与える前に阻止します。スタンドアロン バイナリと埋め込み SQLite を使用して、承認ゲート、監査証跡、AI エージェント ガードレールをすべてのデータベース操作に追加します。外部コントロールプレーン DB は必要ありません。
🔐 承認ワークフロー — マルチステップ、条件付き自動承認、TOML ポリシー エンジン
📋 監査ログ — 改ざん防止ハッシュ チェーン、24 種類のイベント、SQL 編集
🤖 MCP ネイティブ — 12 のツール、6 つのプロンプト、引き出しサポート。 AI エージェントは安全に動作します。チームセットアップ用のリモート HTTP トランスポート
⚡ スタンドアロン バイナリ — CLI、サーバー、およびエージェントは、SQLite が埋め込まれた自己完結型の Rust バイナリとして出荷されます。外部コントロールプレーン DB なし
🔒 エージェントの分離 — DB 認証情報がエージェントから離れることはありません。 CLI/AI はデータベースに直接触れることはありません
🛡️ SQL 安全性レビュー — リスク分類、DDL 検出、DROP ブロック。安全なクエリを自動承認し、危険なクエリについては承認を必要とする
🔍 プリフライト — 送信する前に SQL を分析します。リクエストを作成せずに、リスク レベルを取得し、計画を説明し、結果を確認し、ヒントを修正します。 AI エージェントは承認を求める前に安全な SQL に集中します
🧠 自動スキーマコンテキスト — エージェント

テーブル構造、列、FK、および行数を自動的に収集します。 AI ツールは MCP リソース経由でスキーマにアクセスします - 手動ドキュメントは必要ありません
💬 Slack の承認 — ワンクリックで Slack から承認/拒否します。 dbwardlack init はアプリマニフェストを生成します
🚨 ブレークグラス — 強制的な理由と監査による緊急バイパス。オペレーター/管理者のみ、MCP 経由では利用不可
🆓 コア機能は無料 — 承認、監査、MCP、Slack、ブレークグラスはすべて Apache-2.0 に含まれています。チーム機能 (OIDC、グループ認証) には商用ライセンスが必要です
┌─────────────────────────┐
│ dbward クライアント (CLI / MCP) │
│ DB 認証情報なし — リクエストを送信し、結果を受信します │
━─────┬───────────────────┘
│ REST API
▼
┌─────────────────────────┐
│ dbward サーバー │
│ 承認エンジン │ ポリシーエンジン │ 監査ログ（ハッシュチェーン） │
│ Ed25519 トークン署名 │ OIDC/API 認証 │ Webhook │
│ インメモリ結果リレー │ データベース認証情報なし │
━━━━━━━━━━━━━━━━━━━━━┘
▲ エージェントのポーリング (アウトバウンド HTTPS)
│
┌─────┴─────────────────────┐
│

ドブワードエージェント│
│ DB 認証情報はここのみ │ 承認された操作を実行 │
│ トークン検証 (Ed25519) │ 複数の DB のサポート │
━─────┬───────────────────┘
│
▼
対象データベース (PostgreSQL / MySQL)
重要な原則: クライアントの要求。サーバーが判断します。エージェントが実行されます。必要以上にアクセスできるコンポーネントはありません。
2 分で承認フローを試してください (Docker):
git clone https://github.com/dbward-dev/dbward.git && cd dbward/examples/quickstart
ドッカー構成 -d
docker compose run --rm aliceexecute " SELECT version() " -edevelopment
次に、提出→承認→実行→監査となります。完全なウォークスルー: Docker を使用したクイックスタート
クイックスモークテスト (ローカルインストール):
カール -fsSL https://dbward.dev/install.sh |しー
dbward dev --database-url " postgres://user:pass@localhost:5432/mydb "
# 別の端末の場合:
dbward --config ~ /.dbward/dev/client.toml --database app 実行 " SELECT 1 "
開発モードでは、すべてを自動承認して迅速な反復を実現します。詳細については、「データベースの接続」を参照してください。
完全なリファレンス: docs/reference/mcp.md
{
"mcpサーバー": {
"dbward" : {
"コマンド" : " dbward " ,
"args" : [ "mcp " ]
}
}
}
MCP ツール (12):
MCP プロンプト (6): review_migration 、 Explain_request 、draft_migration 、draft_rollback 、summary_audit_trail 、prepare_approval_comment
誘導: 運用操作では、dbward は続行する前に AI クライアントに理由を尋ねます (クライアントが MCP 誘導をサポートしている場合)。
リモート MCP (HTTP): チーム セットアップの場合、サーバーは HTTP 経由で MCP を公開します。ローカル バイナリは必要ありません (9 つのツール、ローカルのみの移行ツールを除く)。
{
"mcpサーバー": {
"dbward" : {
"type" : " streamable-http " 、
"url" : " https://あなたのサーバー。

example.com/mcp "
}
}
}
オンデマンド実行
dbward はオンデマンド実行を使用します。エージェントは承認時に実行されません。代わりに、クライアントは結果を受け取る準備ができたら明示的にリクエストを再開します。
1. クライアントがリクエストを作成 → サーバーがポリシーを評価 → 保留中 / auto_approved
2. (保留中の場合) CLI 経由で人間が承認します。
3. クライアントが再開します (`dbward request submit <id>`) → サーバーは「ディスパッチ済み」としてマークします。
4. エージェントが DB 上でポーリング、クレーム、実行 → 結果をサーバーに返す
5. サーバーは、メモリ内の結果を待機中のクライアントにリレーします (ロングポーリング)
6. クライアントは結果を表示します (サーバーはローカル FS または S3 に保持されます)
結果はデフォルトでサーバー上に保存されます (ローカル ファイルシステムまたは S3、 [result_storage] 経由で構成可能)。インメモリ リレーには、ストリーミング配信用の 10 分の TTL があります。単一リクエストの永続性をスキップするには、--no-result-store を使用します。
server.toml で定義され、SIGHUP 経由でホットリロードされます。 「構成リファレンス」を参照してください。
操作に承認が必要かどうかを制御します。
[[ ワークフロー ]]
データベース = " * "
環境 = 「生産」
操作 = [ "実行選択" 、 "移行_アップ" 、 "移行_ダウン " ]
[[ ワークフロー .手順]]
type = "承認"
[[ ワークフロー .手順。承認者]]
役割 = " 管理者 "
分 = 1
# ステージングでリスクの低いクエリを自動承認します。リスクのあるものはまだ承認が必要です
[[ ワークフロー ]]
データベース = " * "
環境 = "ステージング"
[ ワークフロー .自動承認]
モード = " リスクベース "
リスク = 「低」
[[ ワークフロー .手順]]
type = "承認"
[[ ワークフロー .手順。承認者]]
役割 = " 管理者 "
分 = 1
実行ポリシー
再実行制限の制御 (レート制限):
[[ 実行ポリシー ]]
データベース = " プライマリ "
環境 = 「生産」
最大実行数 = 10
実行ウィンドウ秒 = 3600
失敗時の再試行 = false
結果ポリシー
結果とストレージにアクセスできるユーザーを制御します。
[[結果ポリシー

]]
データベース = " プライマリ "
環境 = 「生産」
delivery_mode = "ストリーム"
アクセス = [ "リクエスタ " , " 管理者 " ]
通知ポリシー
データベース × 環境ごとに Webhook をルーティングします。
[[通知ポリシー]]
データベース = " プライマリ "
環境 = 「生産」
[[ notification_policies . Webhook ]]
URL = " https://hooks.slack.com/services/... "
フォーマット = " スラック "
CLI リファレンス
完全なリファレンス: docs/reference/cli.md
dbward [オプション] <コマンド>
コマンド:
init 対話型セットアップ ウィザード
医師が接続と構成を診断します
ログイン OIDC ログイン (ブラウザまたはヘッドレスの場合は --device)
ログアウト トークンを取り消し、資格情報を削除します。
whoami 現在のアイデンティティと役割を表示する
移行 移行の実行 (アップ/ダウン/ステータス/作成)
実行 SQL を実行 (--emergency --ブレークグラスの理由)
Audit 監査ログの検索 (ハッシュ チェーン チェックの --verify)
mcp MCP stdio サーバーを起動します
サーバー サーバー管理 (起動、トークンの作成/取り消し、リロード)
エージェント エージェントを起動します
dev ローカル開発サーバー + エージェントを開始します
self-update dbward を最新バージョンに更新します
request リクエストの管理:
list リストリクエスト (--pending-for-me、--status)
show リクエストの詳細を表示
承認 保留中のリクエストを承認します
拒否 保留中のリクエストを拒否します
再開 再開して結果を待ちます
cancel 保留中のリクエストをキャンセルします
token APIトークンの管理（作成/リスト/取り消し）
user ユーザーの管理 (リスト/一時停止/有効化)
Slack Slack の統合:
init アプリのマニフェストと作成 URL を生成する
ポリシー ポリシーツール:
リクエストに対する有効なポリシーを解決する
グローバルオプション:
--version、-v バージョンを表示して終了します
--config <PATH> 構成ファイル (スタンドアロン モード。自動検出の場合は省略)
--database <名前> ターゲット データベース [環境: DBWARD_DATABASE]
--environment <ENV> 環境 [env: DBWARD_ENV]
REST API
完全なリファレンス: docs/reference/api.md
すべてのエンドポイントの完全な API リファレンスを参照してください。

パラメータ、権限、応答形式。
脅威モデルと強化ガイド: docs/security/
ゼロトラスト クライアント - 開発者のマシンは DB 認証情報を持ちません
署名付き実行トークン — Ed25519。トークンには SQL の SHA-256 ハッシュとターゲット データベースが含まれます
トークンのリプレイ防止 - 実行/失敗したリクエストは新しいトークンを発行しません
複数ステートメントの拒否 — ステートメント連鎖による SQL インジェクションを防止します
書き込み可能な CTE 検出 — WITH x AS (DELETE ...) SELECT ... DML として分類
RBAC — 管理者 (システム管理)、要求者 (SQL 操作)、オペレーター (監視 + ブレークグラス)、承認者 (レビュー)
ネットワーク分離 — サーバーには DB 資格情報がありません。エージェントは送信のみに接続します
API トークン認証 — SHA-256 ハッシュ、プレフィックス + ハッシュ複合ルックアップ
OIDC 認証 — JWKS キャッシュによる JWT 検証、RS256/ES256、CLI 用の PKCE (チーム)
監査ハッシュ チェーン - すべてのイベントをリンクする SHA-256 チェーン、改ざん防止
ターゲット
ステータス
Linux x86_64 (glibc)
✅ サポートされています
Linux aarch64 (glibc)
✅ サポートされています
macOS アップルシリコン
✅ サポートされています
macOS インテル
✅ サポートされています
窓
❌ サポートされていません
事前に構築されたバイナリは GitHub Releases で入手できます。 Docker イメージは linux/amd64 および linux/arm64 用に公開されています。
注: 事前に構築されたバイナリと Docker イメージには、商用ライセンスのコンポーネントが含まれています。無料プランの制限内で無料で使用できます。詳細については、「ライセンス」を参照してください。
データベース
ステータス
PostgreSQL
✅ サポートされています
MySQL
✅ サポートされています
URL スキーム ( postgres:// または mysql:// ) から自動検出されます。
完全なガイド: docs/guides/authentication.md
# 最初のサーバーの起動時に自動的に作成される初期トークン:
cat ./data/admin-token # 管理者トークン
cat ./data/agent-token # エージェントトークン
# API 経由の追加トークン:
dbward トークン作成 --subject alice --role admin
OIDC (チーム)
dbward ログイン # ブラウザベース (PKCE)
dbward login --device # ヘッドレス (S

SH、コンテナ)
dbward whoami # 本人確認
dbward logout # トークンの取り消しと削除
通知
インタラクティブなボタンを使用して、Slack から直接リクエストを承認および拒否します。
dbwardlack init --server-url https://your-server.example.com
# → Slack アプリマニフェストを生成し、作成 URL を開きます
Serで設定する

[切り捨てられた]

## Original Extract

Context-aware approval gateway for production database operations. AI, CLI, and CI request. Humans approve. Agents execute. - dbward-dev/dbward

GitHub - dbward-dev/dbward: Context-aware approval gateway for production database operations. AI, CLI, and CI request. Humans approve. Agents execute. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
dbward-dev
/
dbward
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
699 Commits 699 Commits .githooks .githooks .github .github commercial commercial crates crates demo demo deploy deploy dev dev docs docs examples examples .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE LICENSE-APACHE LICENSE-APACHE LICENSE-COMMERCIAL LICENSE-COMMERCIAL README.md README.md SECURITY.md SECURITY.md deny.toml deny.toml install.sh install.sh View all files Repository files navigation
Open-core project — core components are Apache-2.0 . Some features and pre-built binaries include code under the dbward Commercial License . See License for details.
Approval workflows and audit logs for your production database.
Stop accidents before they hit production. Add approval gates, audit trails, and AI agent guardrails to every database operation — with standalone binaries and embedded SQLite. No external control-plane DB required.
🔐 Approval workflows — multi-step, conditional auto-approve, TOML policy engine
📋 Audit logs — tamper-evident hash chain, 24 event types, SQL redaction
🤖 MCP-native — 12 tools, 6 prompts, elicitation support. AI agents operate safely. Remote HTTP transport for team setups
⚡ Standalone binaries — CLI, server, and agent ship as self-contained Rust binaries with embedded SQLite. No external control-plane DB
🔒 Agent isolation — DB credentials never leave the agent. CLI/AI never touch your database directly
🛡️ SQL safety review — risk classification, DDL detection, DROP blocking. Auto-approve safe queries, require approval for risky ones
🔍 Preflight — analyze SQL before submitting. Get risk level, EXPLAIN plan, review findings, and fix hints without creating a request. AI agents converge on safe SQL before asking for approval
🧠 Auto schema context — the agent collects table structures, columns, FKs, and row counts automatically. AI tools access schema via MCP resources — no manual documentation needed
💬 Slack approvals — approve/reject from Slack with one click. dbward slack init generates the app manifest
🚨 Break-glass — emergency bypass with mandatory reason and audit. Operator/admin only, not available via MCP
🆓 Core features free — approval, audit, MCP, Slack, break-glass all included under Apache-2.0 . Team features (OIDC, group auth) require a commercial license
┌─────────────────────────────────────────────────────────┐
│ dbward client (CLI / MCP) │
│ No DB credentials — sends requests, receives results │
└──────────┬───────────────────────────────────────────────┘
│ REST API
▼
┌─────────────────────────────────────────────────────────┐
│ dbward server │
│ Approval engine │ Policy engine │ Audit log (hash chain) │
│ Ed25519 token signing │ OIDC/API auth │ Webhooks │
│ In-memory result relay │ NO database credentials │
└─────────────────────────────────────────────────────────┘
▲ Agent polls (outbound HTTPS)
│
┌──────────┴───────────────────────────────────────────────┐
│ dbward agent │
│ DB credentials here only │ Executes approved operations │
│ Token verification (Ed25519) │ Multiple DB support │
└──────────┬───────────────────────────────────────────────┘
│
▼
Target Database (PostgreSQL / MySQL)
Key principle : The client requests. The server decides. The agent executes. No component has more access than it needs.
Try the approval flow in 2 minutes (Docker):
git clone https://github.com/dbward-dev/dbward.git && cd dbward/examples/quickstart
docker compose up -d
docker compose run --rm alice execute " SELECT version() " -e development
Then submit → approve → execute → audit. Full walkthrough: Quickstart with Docker
Quick smoke test (local install):
curl -fsSL https://dbward.dev/install.sh | sh
dbward dev --database-url " postgres://user:pass@localhost:5432/mydb "
# In another terminal:
dbward --config ~ /.dbward/dev/client.toml --database app execute " SELECT 1 "
Dev mode auto-approves everything for fast iteration. See Connect Your Database for details.
Full reference: docs/reference/mcp.md
{
"mcpServers" : {
"dbward" : {
"command" : " dbward " ,
"args" : [ " mcp " ]
}
}
}
MCP Tools (12):
MCP Prompts (6): review_migration , explain_request , draft_migration , draft_rollback , summarize_audit_trail , prepare_approval_comment
Elicitation: On production operations, dbward asks the AI client for a reason before proceeding (if the client supports MCP elicitation).
Remote MCP (HTTP): For team setups, the server exposes MCP over HTTP — no local binary needed (9 tools, excludes local-only migration tools):
{
"mcpServers" : {
"dbward" : {
"type" : " streamable-http " ,
"url" : " https://your-server.example.com/mcp "
}
}
}
On-Demand Execution
dbward uses on-demand execution : the agent does not execute on approval. Instead, the client explicitly resumes the request when ready to receive the result.
1. Client creates request → server evaluates policy → pending / auto_approved
2. (If pending) Human approves via CLI
3. Client resumes (`dbward request resume <id>`) → server marks as "dispatched"
4. Agent polls, claims, executes on DB → returns result to server
5. Server relays result in-memory to waiting client (long poll)
6. Client displays result (server persists to local FS or S3)
Results are persisted on the server by default (local filesystem or S3, configurable via [result_storage] ). The in-memory relay has a 10-minute TTL for streaming delivery. Use --no-result-store to skip persistence for a single request.
Defined in server.toml and hot-reloaded via SIGHUP. See Configuration Reference .
Control whether operations require approval:
[[ workflows ]]
database = " * "
environment = " production "
operations = [ " execute_select " , " migrate_up " , " migrate_down " ]
[[ workflows . steps ]]
type = " approval "
[[ workflows . steps . approvers ]]
role = " admin "
min = 1
# Auto-approve low-risk queries in staging; risky ones still need approval
[[ workflows ]]
database = " * "
environment = " staging "
[ workflows . auto_approve ]
mode = " risk_based "
risk = " low "
[[ workflows . steps ]]
type = " approval "
[[ workflows . steps . approvers ]]
role = " admin "
min = 1
Execution Policies
Control re-execution limits (rate limiting):
[[ execution_policies ]]
database = " primary "
environment = " production "
max_executions = 10
execution_window_secs = 3600
retry_on_failure = false
Result Policies
Control who can access results and storage:
[[ result_policies ]]
database = " primary "
environment = " production "
delivery_mode = " stream "
access = [ " requester " , " admin " ]
Notification Policies
Route webhooks per database × environment:
[[ notification_policies ]]
database = " primary "
environment = " production "
[[ notification_policies . webhooks ]]
url = " https://hooks.slack.com/services/... "
format = " slack "
CLI Reference
Full reference: docs/reference/cli.md
dbward [OPTIONS] <COMMAND>
Commands:
init Interactive setup wizard
doctor Diagnose connectivity and configuration
login OIDC login (browser or --device for headless)
logout Revoke tokens and delete credentials
whoami Show current identity and role
migrate Run migrations (up/down/status/create)
execute Execute SQL (--emergency --reason for break-glass)
audit Search audit log (--verify for hash chain check)
mcp Start MCP stdio server
server Server management (start, token create/revoke, reload)
agent Start the agent
dev Start local dev server + agent
self-update Update dbward to the latest version
request Manage requests:
list List requests (--pending-for-me, --status)
show Show request detail
approve Approve a pending request
reject Reject a pending request
resume Resume and wait for result
cancel Cancel a pending request
token Manage API tokens (create/list/revoke)
user Manage users (list/suspend/activate)
slack Slack integration:
init Generate app manifest and creation URL
policy Policy tools:
resolve Resolve effective policy for a request
Global Options:
--version, -v Show version and exit
--config <PATH> Config file (standalone mode; omit for auto-detect)
--database <NAME> Target database [env: DBWARD_DATABASE]
--environment <ENV> Environment [env: DBWARD_ENV]
REST API
Full reference: docs/reference/api.md
See full API reference for all endpoints, parameters, permissions, and response formats.
Threat model and hardening guide: docs/security/
Zero-trust client — developer machines never have DB credentials
Signed execution tokens — Ed25519. Token includes SHA-256 hash of SQL + target database
Token replay prevention — executed/failed requests don't issue new tokens
Multi-statement rejection — prevents SQL injection via statement chaining
Writable CTE detection — WITH x AS (DELETE ...) SELECT ... classified as DML
RBAC — admin (system management), requester (SQL operations), operator (monitoring + break-glass), approver (review)
Network isolation — server has no DB credentials; agent connects outbound only
API token auth — SHA-256 hashed, prefix+hash composite lookup
OIDC auth — JWT verification with JWKS caching, RS256/ES256, PKCE for CLI (Team)
Audit hash chain — SHA-256 chain linking all events, tamper-evident
Target
Status
Linux x86_64 (glibc)
✅ Supported
Linux aarch64 (glibc)
✅ Supported
macOS Apple Silicon
✅ Supported
macOS Intel
✅ Supported
Windows
❌ Not supported
Pre-built binaries are available on GitHub Releases . Docker images are published for linux/amd64 and linux/arm64 .
Note: Pre-built binaries and Docker images include commercial-licensed components. They are free to use within Free plan limits. See LICENSE for details.
Database
Status
PostgreSQL
✅ Supported
MySQL
✅ Supported
Auto-detected from URL scheme ( postgres:// or mysql:// ).
Full guide: docs/guides/authentication.md
# Initial tokens created automatically on first server start:
cat ./data/admin-token # admin token
cat ./data/agent-token # agent token
# Additional tokens via API:
dbward token create --subject alice --role admin
OIDC (Team)
dbward login # Browser-based (PKCE)
dbward login --device # Headless (SSH, containers)
dbward whoami # Check identity
dbward logout # Revoke + delete tokens
Notifications
Approve and reject requests directly from Slack with interactive buttons:
dbward slack init --server-url https://your-server.example.com
# → generates Slack App Manifest, opens creation URL
Configure in ser

[truncated]
