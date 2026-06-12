---
source: "https://github.com/mpospirit-apps/TetherDust"
hn_url: "https://news.ycombinator.com/item?id=48504836"
title: "Show HN: TetherDust – Self-hosted AI Analytics Engineer (open source)"
article_title: "GitHub - mpospirit-apps/TetherDust: Self-hosted AI Analytics Engineer · GitHub"
author: "masterpos"
captured_at: "2026-06-12T16:07:46Z"
capture_tool: "hn-digest"
hn_id: 48504836
score: 2
comments: 0
posted_at: "2026-06-12T14:43:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: TetherDust – Self-hosted AI Analytics Engineer (open source)

- HN: [48504836](https://news.ycombinator.com/item?id=48504836)
- Source: [github.com](https://github.com/mpospirit-apps/TetherDust)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T14:43:27Z

## Translation

タイトル: Show HN: TetherDust – セルフホスト型 AI 分析エンジニア (オープンソース)
記事タイトル: GitHub - mpospirit-apps/TetherDust: セルフホスト型 AI 分析エンジニア · GitHub
説明: セルフホスト型 AI 分析エンジニア。 GitHub でアカウントを作成して、mpospirit-apps/TetherDust の開発に貢献してください。

記事本文:
GitHub - mpospirit-apps/TetherDust: セルフホスト型 AI 分析エンジニア · GitHub
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
mpospiritアプリ
/
テザーダスト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主なブランチのタグが取得されました

o ファイル コード [その他のアクション] メニューを開く フォルダーとファイル
5 コミット 5 コミット .github .github 変更ログ 変更ログ docker docker docs/ イメージ docs/ イメージ ドキュメント/ TetherDust ドキュメント ドキュメント/ TetherDust ドキュメント tetherdust tetherdust .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md バージョン バージョン docker-compose.ci.yml docker-compose.ci.yml docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
TetherDust は、コンテナ化された Model Context Protocol (MCP) サーバーを使用して、コードベースとデータベースの間のギャップを埋めます。 TetherDust では、リポジトリのドキュメントと並行してデータベース スキーマをドキュメント化することで、あらゆる AI エージェントが検証可能な SQL を生成し、動的な d3.js ダッシュボードを構築し、スキーマとコードの依存関係をマッピングできるようになります。このプラットフォームは完全にインフラストラクチャ内で実行され、厳格な読み取り専用クエリ境界、ロールベースのアクセス制御 (RBAC)、および不変の監査ログを強制します。
TetherDust は、AI 主導のデータ インタラクションのための柔軟なプラットフォームとなるように設計されており、次のような機能があります。
豊富な Markdown サポートを使用して、自然言語プロンプトから適切に構造化された Wiki のようなコードベースとデータベースのドキュメントを生成します。
TetherDust は、データベース ドキュメントとともに GitHub コードベース (またはコードベース ドキュメント) を参照してください。エージェントは両方を調査し、どのコード ファイルがどのテーブルと列を読み書きするかを示す対話型の視覚的なグラフを作成し、スキーマの変動に応じてバージョン管理します。
必要なダッシュボードについて説明します。 AI エージェントは、グラフごとに SQL と d3.js コードを作成します。チャートはスケジュールに従って自動更新され、キャッシュされます。

またはパフォーマンス。
生成されたグラフをカスタム動作用に直接編集するか、要件が変更されたときに更新するようにエージェントに依頼します。
クエリを定義してスケジュールに従って実行し、結果を電子メールまたはダウンロードで配信します。
チャットを使用すると、TetherDust のすべての機能に 1 か所でアクセスできます。
データについて自然言語で質問すると、ドキュメントに基づいた回答がストリーミングで得られます。ドキュメント ソースを名前で言及して特定のコンテキストを引き出すことも、何を使用するかをエージェントに決定させることもできます。
TetherDust は、ユーザーのリクエストに応じて、または応答する前に詳細を確認するために、SQL クエリを作成して実行できます。
チャットから名前でレポート、ダッシュボード、テザーにアクセスします。
CLI ツール、API 呼び出し、または Ollama を使用して、MCP を話すエージェントに接続します。現在サポートされているエージェント統合:
Python SQLAlchemy ダイアレクトと読み取り専用ユーザーを使用してデータベースを接続します。現在サポートされているデータベース:
PostgreSQL、MySQL/MariaDB、SQL Server、SQLite、ClickHouse、Oracle、Snowflake、BigQuery。
さらに多くのエージェントとデータベースが登場予定
このアーキテクチャはエージェントに依存しないように設計されており、新しいエージェントを追加するためのシンプルなインターフェイスを備えています。
エージェント ランタイムはコンテナ化されているため、エージェントが TetherDust の機能と対話する唯一の方法は、ツールとデータ ソースを API として公開する MCP サーバーを経由することです。 TetherDust には、コア機能を公開する MCP サーバーが組み込まれています。
各ユーザーの役割によって、どのデータベース、MCP ツール、ドキュメント ソース、ダッシュボード、レポート、テザーを表示できるかが決まります。
フローチャート LR
ユーザー([👤 ユーザー]) --> ロール{{ロール}}
ロール -- allowed_databases --> DB[(データベース)]
役割 -- allowed_tools --> ツール [MCP ツール]
ロール -- allowed_doc_sources --> ドキュメント[ドキュメント ソース]
ロール -- allowed_prompts --> プロンプト[プロンプト]
ロール -- allowed_mcp_servers --> MCP[カスタム MCP サーバー]
役割 -- can_view_dashboards -->

ダッシュ[ダッシュボード]
役割 -- can_view_reports --> レポート[レポート]
役割 -- can_view_tethers --> Tethers[Tethers]
DB --> エージェント([🤖 エージェント])
ツール --> エージェント
ドキュメント --> エージェント
プロンプト --> エージェント
MCP --> エージェント
エージェント - 。 <br/>許可されたサブセットのみを参照します。-> スコープ[/Permittedscope/]
読み込み中
エージェントには、そのユーザー ロールで許可されているデータベースとツールのみが表示されます。
ロールごとに付与されたリモート HTTP またはローカル サブプロセス MCP サーバー (Notion、GitHub、内部 API、MCP を使用するもの) を使用してエージェントを拡張します。
すべてのエージェント クエリは SQLGlot で解析され、読み取り専用でない限り拒否されます。デフォルトでは接続は読み取り専用であり、実際の信頼境界は読み取り専用のデータベース ユーザーです。常に読み取り専用のデータベース ユーザーと接続します。
アクションとクエリは、不変の監査ログに記録されます。すべてのチャット セッション、エージェントのクエリ、生成の実行は記録され、スタッフが管理コンソールで確認できるようになります。
フルスタックは、Docker Compose プロジェクト、つまり Django Web アプリ (ポータル + 管理者) として出荷されます。
コンソール)、データベースと生成ツールを公開する MCP サーバー、プラグイン可能な AI エージェント
ゲートウェイ (Codex CLI、Claude Code CLI、ダイレクト API/Ollama)、PostgreSQL、Redis、および
バックグラウンドタスク用のセロリワーカー。アクティブなエージェントの切り替えは 1 回の切り替えで行えます。
コンソール — 再起動や設定の変更は必要ありません。
AI エージェントの資格情報 (次のいずれか):
Codex — ChatGPT サブスクリプション認証トークン、または OpenAI API キー
Claude Code — Claude Pro/Max OAuth トークン ( claude setup-token )、または Anthropic API キー
エージェントの資格情報は後で管理コンソールから構成します。スタックを起動する必要はありません。
1. 最初の起動前にシークレットを設定する
すべての資格情報は、Docker Compose が自動的にロードする gitignored .env ファイル内に存在します。
テンプレートから開始します。
cp .env.example .env
.env.example には、データベースと管理者ログインの有効なローカルデフォルトが付属していますが、

の
2 つの暗号キーは意図的に空白になっています。入力するまでスタックは開始されません。
.env を編集し、次のように設定します。
ａ．資格情報暗号化キー (Fernet) を生成します。このキーは保存されているすべてのファイルを暗号化します
データベースのパスワードとエージェントの API キー/トークン - 独自のパスワードを生成し、秘密にしておきます。
python -c " from cryptography.fernet import Fernet; print(Fernet.generate_key().decode()) "
TETHERDUST_ENCRYPTION_KEY を生成された値に設定します。 .env からすべてのファイルに共有されます
それを必要とするサービス ( mcp 、 local-mcp 、 web 、 celery-worker 、 celery-beat ) なので、
一度設定するだけです。
b. Django 秘密鍵を生成します。
python -c " シークレットのインポート; print(secrets.token_urlsafe(64)) "
DJANGO_SECRET_KEY を生成された値に設定します。
c.管理者ログインを設定します。スーパーユーザーは、最初の起動時に次の値から作成されます。
デフォルトの admin / admin が使用されないように変更します。
DJANGO_SUPERUSER_USERNAME=管理者
DJANGO_SUPERUSER_PASSWORD= < a-strong-password >
DJANGO_SUPERUSER_EMAIL=you@example.com
d.データベースのパスワードを変更します。 DB_NAME 、 DB_USER 、および DB_PASSWORD を
選択された値。単一セットの変数が DB サービスにフィードします。両方の MCP 接続が行われます。
文字列、Web/セロリ サービスなど、手作業で同期を保つ必要はありません。
e.内部サービスシークレットを生成します。 2 つの共有シークレットによる認証
TetherDust の内部サービス間呼び出し — MCP_FILTER_SECRET (web/celery →
MCP フィルター登録) および AGENT_GATEWAY_SECRET (Django → Codex/Claude)
ゲートウェイ)。それぞれの値を生成します。
python -c " シークレットのインポート; print(secrets.token_urlsafe(32)) "
空白のままにしてもスタックは開始されますが、それらの内部呼び出しは認証されません。
TetherDust をネットワークに公開する前に、両方を設定してください。
f. (ローカル開発のみ) デバッグ モードを有効にします。 .env.example には同梱されています
DJANGO_DEBUG=false

、本番環境の強化 (安全な Cookie、HTTPS) が可能になります。
リダイレクト、HSTS) であり、アプリの前で TLS を想定しているため、プレーンでログインします
http://localhost は機能しません。ローカル開発の場合は、DJANGO_DEBUG=true (dev
サーバーは自動リロード、強化は緩和されています）。実際のデプロイでは false のままにしておきます。
注: .env は .gitignore にリストされているため、実際のシークレットはバージョン外に保たれます
コントロール。決してコミットしないでください。 「制作ノート」を参照してください。
docker 構成 --build
これにより、PostgreSQL、MCP サーバー、エージェント ゲートウェイ、Redis、Celery、および Django が起動します。
ウェブアプリ。最初の起動ではデータベースの移行が実行され、スーパーユーザーが作成され、自動検出が行われます。
ドキュメントソース。
http://localhost:8000 にアクセスし、手順 1c で設定したスーパーユーザー資格情報を使用してログインします。
4. エージェントとデータベースを接続する
エージェント — エージェント設定 (Codex または Claude Code) を追加し、認証を貼り付けます
トークン/API キーを選択し、アクティブとしてマークします。一度にアクティブにできるエージェントは 1 つだけです。
データベース — クエリを実行するデータベースへの接続を追加します。を使用してください
読み取り専用データベース ユーザー (下記を参照)。
チャットを開いて自然言語で質問してください。
TetherDust は、3 つの読み取り専用保護層を通じてすべてのエージェント クエリを実行します。
SQL 検証 - 各クエリは (データベース言語ごとに SQLGlot 経由で) 解析され、
単一の SELECT /CTE/set 操作でない限り拒否されます。複数ステートメント入力、
データ変更 CTE、 SELECT … INTO 、ストアド プロシージャ呼び出し、DDL/DML はすべて
ブロックされました。
読み取り専用セッション — 読み取り専用とマークされた接続 (デフォルトは ON) が実行されます。
エンジンがサポートする読み取り専用データベース セッション (PostgreSQL、MySQL/MariaDB、
SQLite、Oracle、ClickHouse)。 SQL Server、BigQuery、Snowflake にはセッション レベルがありません
読み取り専用 — そこでは、読み取り専用のユーザー/ロールに依存します (下記)。
読み取り専用データベース ユーザー - 実際の信頼境界。常に接続する

と
読み取りアクセスのみを持つアカウント。上の 2 つの層は多層防御です。ある
読み取り専用資格情報は、エージェントが書き込みできないことを実際に保証するものです。
読み取り専用データベース ユーザーを作成する
-- PostgreSQL
ロールの作成 tetherdust_ro ログインパスワード ' ... ' ;
データベース mydb への tetherdust_ro への接続を許可します。
スキーマ public の使用を tetherdust_ro に許可します。
スキーマ public 内のすべてのテーブルに対する SELECT を tetherdust_ro に付与します。
スキーマのデフォルトの権限を変更します public GRANT SELECT ON TABLES TO tetherdust_ro;
-- MySQL / MariaDB
CREATE USER ' tetherdust_ro '@ ' % ' IDENTIFIED BY ' ... ' ;
mydb に選択を許可します。 * TO 'tetherdust_ro' @ '%' ;
BigQuery の場合、roles/bigquery.dataViewer + role/bigquery.jobUser を付与します（そうではありません）
データエディタ ); Snowflake の場合は、USAGE / SELECT のみを使用してロールを付与します。 SQL用
サーバーはログインを db_datareader ロールに追加します。
保存された資格情報は、手順 1a の Fernet キーを使用して暗号化されます。もし
TETHERDUST_ENCRYPTION_KEY は空白のままにし、認証情報は平文で保存されます — 常に
設定してください。運用環境 ( DJANGO_DEBUG=false ) TetherDust が認証情報の保存を拒否する
鍵なしで。
手順 1 で説明したように、非ローカル展開の前に .env 内のすべてのシークレットを設定します。
デフォルトの Compose 構成は、ローカル開発用に調整されています。露出前
テザーダスト

[切り捨てられた]

## Original Extract

Self-hosted AI Analytics Engineer. Contribute to mpospirit-apps/TetherDust development by creating an account on GitHub.

GitHub - mpospirit-apps/TetherDust: Self-hosted AI Analytics Engineer · GitHub
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
mpospirit-apps
/
TetherDust
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github .github changelog changelog docker docker docs/ images docs/ images documentations/ TetherDust Documentation documentations/ TetherDust Documentation tetherdust tetherdust .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md VERSION VERSION docker-compose.ci.yml docker-compose.ci.yml docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh View all files Repository files navigation
TetherDust bridges the gap between your codebase and databases using containerized Model Context Protocol (MCP) servers. By documenting database schemas alongside repository documentation, TetherDust enables any AI agent to generate verifiable SQL, build dynamic d3.js dashboards, and map schema-to-code dependencies. The platform runs entirely within your infrastructure, enforcing strict read-only query boundaries, role-based access control (RBAC), and immutable audit logging.
TetherDust is designed to be a flexible platform for AI-driven data interaction, with features that include:
Generate well-structured, wiki-like codebase and database documentation from natural language prompts, with rich Markdown support.
Point TetherDust at a GitHub codebase (or codebase documentation) together with a database documentation. The agent explores both and produces an interactive visual graph showing which code files read or write which tables and columns, versioned as the schema drifts.
Describe the dashboard you want; the AI agent writes the SQL and the d3.js code for every chart. Charts auto-refresh on a schedule and are cached for performance.
Edit the generated chart directly for custom behavior, or ask the agent to update it when requirements change.
Define queries and run them on a schedule, delivering results by email or download.
Use Chat to access all of TetherDust's capabilities in one place.
Ask natural language questions about your data and get streamed answers grounded in your documentation. You can mention documentation sources by name to pull in specific context, or let the agent decide what to use.
TetherDust can write and execute SQL queries — either at your request or to confirm details before answering.
Reach reports, dashboards, and tethers by name from the chat.
Use CLI tools, API calls, or Ollama to connect any agent that speaks MCP. Currently supported agent integrations:
Connect any database with a Python SQLAlchemy dialect and a read-only user. Currently supported databases:
PostgreSQL, MySQL/MariaDB, SQL Server, SQLite, ClickHouse, Oracle, Snowflake, BigQuery.
Many more agents and databases to come
The architecture is designed to be agent-agnostic, with a simple interface for adding new ones.
Agent runtimes are containerized, so the only way for an agent to interact with TetherDust's features is through MCP servers, which expose tools and data sources as APIs. TetherDust includes a built-in MCP server that exposes the core features.
Every user's role decides which databases, MCP tools, documentation sources, dashboards, reports, and tethers they can see.
flowchart LR
User([👤 User]) --> Role{{Role}}
Role -- allowed_databases --> DB[(Databases)]
Role -- allowed_tools --> Tools[MCP Tools]
Role -- allowed_doc_sources --> Docs[Doc Sources]
Role -- allowed_prompts --> Prompts[Prompts]
Role -- allowed_mcp_servers --> MCP[Custom MCP Servers]
Role -- can_view_dashboards --> Dash[Dashboards]
Role -- can_view_reports --> Reports[Reports]
Role -- can_view_tethers --> Tethers[Tethers]
DB --> Agent([🤖 Agent])
Tools --> Agent
Docs --> Agent
Prompts --> Agent
MCP --> Agent
Agent -. only sees the<br/>allowed subset .-> Scope[/Permitted scope/]
Loading
Agents only see the databases and tools their user role allows.
Extend the agent with remote HTTP or local subprocess MCP servers (Notion, GitHub, internal APIs, anything that speaks MCP), granted per role.
Every agent query is parsed with SQLGlot and rejected unless it is read-only. Connections are read-only by default, and the real trust boundary is a read-only database user — always connect with one.
Actions and queries are logged in an immutable audit log. Every chat session, agent query, and generation run is recorded and reviewable by staff in the admin console.
The full stack ships as a Docker Compose project: a Django web app (portal + admin
console), an MCP server that exposes database and generation tools, pluggable AI agent
gateways (Codex CLI, Claude Code CLI, direct API/Ollama), PostgreSQL, Redis, and
Celery workers for background tasks. Switching the active agent is a single toggle in
the console — no restarts, no config changes.
An AI agent credential (one of):
Codex — a ChatGPT subscription auth token, or an OpenAI API key
Claude Code — a Claude Pro/Max OAuth token ( claude setup-token ), or an Anthropic API key
You configure the agent credential later, from the admin console — it is not needed to boot the stack.
1. Set your secrets before the first launch
All credentials live in a gitignored .env file that Docker Compose loads automatically.
Start from the template:
cp .env.example .env
.env.example ships with working local defaults for the database and admin login, but the
two cryptographic keys are intentionally blank — the stack will not start until you fill
them in. Edit .env and set the following.
a. Generate a credential-encryption key (Fernet). This key encrypts all stored
database passwords and agent API keys/tokens — generate your own and keep it secret:
python -c " from cryptography.fernet import Fernet; print(Fernet.generate_key().decode()) "
Set TETHERDUST_ENCRYPTION_KEY to the generated value. It is shared from .env to every
service that needs it ( mcp , local-mcp , web , celery-worker , celery-beat ), so you
only set it once.
b. Generate a Django secret key:
python -c " import secrets; print(secrets.token_urlsafe(64)) "
Set DJANGO_SECRET_KEY to the generated value.
c. Set the admin login. The superuser is created on first boot from these values —
change them so the default admin / admin is never used:
DJANGO_SUPERUSER_USERNAME=admin
DJANGO_SUPERUSER_PASSWORD= < a-strong-password >
DJANGO_SUPERUSER_EMAIL=you@example.com
d. Change the database password. Set DB_NAME , DB_USER , and DB_PASSWORD to your
chosen values. A single set of variables feeds the db service, both MCP connection
strings, and the web/celery services — there is nothing to keep in sync by hand.
e. Generate the internal service secrets. Two shared secrets authenticate
TetherDust's internal service-to-service calls — MCP_FILTER_SECRET (web/celery →
MCP filter registration) and AGENT_GATEWAY_SECRET (Django → the Codex/Claude
gateways). Generate a value for each:
python -c " import secrets; print(secrets.token_urlsafe(32)) "
If left blank the stack still starts, but those internal calls are unauthenticated —
set both before exposing TetherDust to a network.
f. (Local development only) Enable debug mode. .env.example ships with
DJANGO_DEBUG=false , which enables production hardening (secure cookies, HTTPS
redirect, HSTS) and assumes TLS in front of the app — so logging in over plain
http://localhost won't work. For local development set DJANGO_DEBUG=true (dev
server with auto-reload, hardening relaxed). Leave it false for any real deployment.
Note: .env is listed in .gitignore , so your real secrets stay out of version
control. Never commit it. See Production notes .
docker compose up --build
This starts PostgreSQL, the MCP server, the agent gateways, Redis, Celery, and the Django
web app. First boot runs database migrations, creates your superuser, and auto-discovers
documentation sources.
Visit http://localhost:8000 and log in with the superuser credentials you set in step 1c.
4. Connect an agent and a database
Agents — add an agent configuration (Codex or Claude Code), paste your auth
token/API key, and mark it active. Only one agent is active at a time.
Databases — add a connection to the database you want to query. Use a
read-only database user (see below).
Open the chat and ask a question in natural language.
TetherDust runs every agent query through three layers of read-only protection:
SQL validation — each query is parsed (via SQLGlot, per database dialect) and
rejected unless it is a single SELECT /CTE/set-operation. Multi-statement input,
data-modifying CTEs, SELECT … INTO , stored-procedure calls, and DDL/DML are all
blocked.
Read-only session — connections marked Read-only (default ON) run in a
read-only database session where the engine supports it (PostgreSQL, MySQL/MariaDB,
SQLite, Oracle, ClickHouse). SQL Server, BigQuery, and Snowflake have no session-level
read-only — there, rely on a read-only user/role (below).
Read-only database user — the real trust boundary. Always connect with an
account that only has read access. The two layers above are defense-in-depth; a
read-only credential is what actually guarantees the agent can't write.
Create a read-only database user
-- PostgreSQL
CREATE ROLE tetherdust_ro LOGIN PASSWORD ' ... ' ;
GRANT CONNECT ON DATABASE mydb TO tetherdust_ro;
GRANT USAGE ON SCHEMA public TO tetherdust_ro;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO tetherdust_ro;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO tetherdust_ro;
-- MySQL / MariaDB
CREATE USER ' tetherdust_ro '@ ' % ' IDENTIFIED BY ' ... ' ;
GRANT SELECT ON mydb. * TO ' tetherdust_ro ' @ ' % ' ;
For BigQuery grant roles/bigquery.dataViewer + roles/bigquery.jobUser (not
dataEditor ); for Snowflake grant a role with USAGE / SELECT only; for SQL
Server add the login to the db_datareader role.
Stored credentials are encrypted with the Fernet key from step 1a. If
TETHERDUST_ENCRYPTION_KEY is left blank, credentials are stored in plaintext — always
set it. In production ( DJANGO_DEBUG=false ) TetherDust refuses to save a credential
without a key.
Set every secret in .env before any non-local deployment, as described in step 1.
The default Compose configuration is tuned for local development. Before exposing
TetherDust

[truncated]
