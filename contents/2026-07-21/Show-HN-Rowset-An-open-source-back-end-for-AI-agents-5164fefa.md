---
source: "https://github.com/LVTD-LLC/rowset"
hn_url: "https://news.ycombinator.com/item?id=48992773"
title: "Show HN: Rowset – An open-source back end for AI agents"
article_title: "GitHub - LVTD-LLC/rowset: Rowset gives your AI agent a dataset backend. · GitHub"
author: "rasulkireev"
captured_at: "2026-07-21T14:57:19Z"
capture_tool: "hn-digest"
hn_id: 48992773
score: 1
comments: 0
posted_at: "2026-07-21T14:24:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Rowset – An open-source back end for AI agents

- HN: [48992773](https://news.ycombinator.com/item?id=48992773)
- Source: [github.com](https://github.com/LVTD-LLC/rowset)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T14:24:49Z

## Translation

タイトル: Show HN: Rowset – AI エージェント用のオープンソース バックエンド
記事のタイトル: GitHub - LVTD-LLC/rowset: Rowset は AI エージェントにデータセット バックエンドを提供します。 · GitHub
説明: Rowset は AI エージェントにデータセット バックエンドを提供します。 GitHub でアカウントを作成して、LVTD-LLC/行セットの開発に貢献します。

記事本文:
GitHub - LVTD-LLC/rowset: Rowset は AI エージェントにデータセット バックエンドを提供します。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
LVTD-LLC
/
行セット
公共
通知
あなたは署名しているに違いありません

通知設定を変更する必要がある
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
645 コミット 645 コミット .agents/ スキル .agents/ スキル .cursor/ ルール .cursor/ ルール .github .github .greptile .greptile .seo .seo アプリ アプリ cli cli デプロイメント デプロイメント ドキュメント ドキュメント評価 評価 フロントエンド フロントエンド行セット行セット スクリプト スクリプト .browserslistrc .browserslistrc .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .nvmrc .nvmrc .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version .stylelintrc.json .stylelintrc.json AGENTS.md AGENTS.md ANALYTICS.md ANALYTICS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md DESIGN.md DESIGN.md Dockerfile-python Dockerfile-python GEMINI.md GEMINI.md Makefile Makefile PRODUCT.md PRODUCT.md README.md README.md SELF_HOSTING.md SELF_HOSTING.md STRUCTURE.md STRUCTURE.md TECH.md TECH.md VISION.md VISION.md conftest.py conftest.py djass-manifest.json djass-manifest.json docker-compose-local.yml docker-compose-local.yml docker-compose-prod.yml docker-compose-prod.yml docker-compose-test.yml docker-compose-test.yml eslint.config.js eslint.config.js manage.py manage.py package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js project-metadata.json project-metadata.json pyproject.toml pyproject.toml pytest.ini pytest.ini render.yaml render.yaml tailwind.config.js tailwind.config.js uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Rowset は、AI を信頼する構造化データセット用のプライベート MCP および REST バックエンドです
エージェントは作成、検査、更新、エクスポート、共有できます。ユーザーはサインインし、
エージェントのセットアップ プロンプトを表示し、スコープ指定された API キーを認証して、エージェントが操作できるようにします。
を通じて所有するデータセット

ブラウザの代わりに安定したプログラムインターフェイス
自動化。
AI エージェント ワークフロー用のホストされたストリーミング可能な HTTP MCP サーバー。
アカウントチェック、プロジェクト、データセット、行、エクスポート用の認証済み REST API
リレーションシップ、画像アセット、パブリック プレビュー設定。
同じ認証された REST 操作については、cli/ の下に行セットを移動します。
安定したヘッダー、セマンティック列メタデータ、永続的な API ベースのデータセット
エージェントの指示、JSON メタデータ、および明示的なインデックス列。
内部行セット行 ID またはデータセット インデックス値による行 CRUD。
関連するデータセットを変更せずに整理するためのプロジェクトとプロジェクト セクション
認証境界。
選択、参照、画像、日付、日時、通貨、数値、ブール値、電子メール、
URL、およびテキスト列のメタデータ。
人間によるレビューのためのオプションのパスワード保護を備えた読み取り専用のパブリック プレビュー。
CSV、JSONL、XLSX、SQLite、およびダッシュボード指向の Parquet エクスポート パス。
ローカル ディスクまたは S3 互換ストレージ上のプライベート イメージ アセット ストレージ
クラウドフレアR2。
オプションの Qdrant 支援ハイブリッド ベクトルおよびデータセット行の語彙検索。
エリア
テクノロジー
言語
Python 3.14.2 ( .python-version 、 pyproject.toml ) および Go for cli/
バックエンド
ジャンゴ6
REST API
ジャンゴ忍者
MCP
Rowset/asgi.py の Starlette を通じて FastMCP がマウントされました
認証
Django 認証、セッション認証、API キー認証、ホストされた MCP ベアラー認証
データストア
PostgreSQL、Redis
バックグラウンドジョブ
Django Q2 ワーカー
表形式の作業
Python csv 、 json 、 sqlite3 、 zipfile 、および Polars
フロントエンド
Django テンプレート、HTMX、Alpine.js、Tailwind、PostCSS
資産
scripts/build-assets.mjs 内のカスタム Node 24 ビルド スクリプト
ローカルスタック
Postgres、Redis、バックエンド、ワーカー、フロントエンド、Mailhog、Stripe CLI、MJML、MinIO を使用した Docker Compose
可観測性
セントリーとポストホッグ
統合
Mailgun、Buttondown、Stripe、Chatwoot、S3 互換ストレージ、Qdrant/OpenR

オプションのベクトル検索用の外側
アクティブな展開パス
Docker イメージと CapRover GitHub アクション
製品の境界
Rowset は意図的にエージェント管理のデータセットを中心に配置されています。
サインインしているユーザーは、Rowset セットアップ プロンプトを信頼できるエージェントにコピーします。
エージェントは API キーを非公開で保存し、次のように Rowset MCP に接続します。
権限: Bearer <key> 。
エージェントはデータセットの作成または検出、スキーマ/コンテキストの検査、行の変更、
プロジェクトの管理、スナップショットのエクスポート、または要求に応じてパブリック プレビューを有効にします。
ダッシュボードは、セットアップ、設定、最近のデータセットの状態、スキーマを管理するのに役立ちます。
レビュー、エクスポート、パブリック プレビュー レビュー、アカウントの回復。
現在の製品パスの範囲外:
行セット所有のソース コネクタ、同期、またはライトバック。
認証または REST/MCP アクセスの代替としてのパブリック プレビュー。
優先されるエージェント統合としてのブラウザ自動化。
広範な BI、ウェアハウス、または ETL オーケストレーションを約束します。
エージェントは引き続きローカル ファイル、Google スプレッドシート、データベース、またはその他の上流を読み取ることができます
独自の機能を持つソースを作成し、構造化された行を Rowset に送信します
MCP または REST を介して。
サポートされているローカル ワークフローの場合:
Docker Compose を使用した Docker デスクトップまたは Docker エンジン。
Docker の外部でのホスト側のデバッグの場合:
Node.js 24.11 以降および npm 11 以降。
ソースから行セット CLI を構築する場合は、1.26 以降を使用してください。
PostgreSQL と Redis は環境からアクセス可能です。
ほとんどの貢献者は、Docker Compose から始める必要があります。ローカルの Compose スタック
Python イメージを構築し、フロントエンド サービスにノードの依存関係をインストールし、
Postgres と Redis を自動的に実行します。
git clone https://github.com/LVTD-LLC/rowset.git
cd 行セット
2. ローカル環境構成の作成
cp .env.example .env
チェックインされたデフォルトは、ローカルの Docker Compose スタック用に設計されています。
Postgres データベース/ユーザー/パスワード: ro

wset
サイトURL: http://localhost:8000
docker compose -f docker-compose-local.yml up -d --build
バックエンド サービスのバックエンド ログ
http://ローカルホスト:8000
Mailhog は次の場所で入手できます。
http://ローカルホスト:8025
MinIO のコンソールは次の場所から入手できます。
http://ローカルホスト:9001
4. アカウントを作成する
ローカル アプリの UI を使用してサインアップします。電子メール検証はノンブロッキングです
現在のアプリ: ローカル確認リンクは Mailhog によってキャプチャされるか、経由で印刷されます。
構成された電子メール バックエンド。
エージェントに適した最小の権限レベルを使用します。
検査および輸出のためにお読みください。
データセット、行、プロジェクト、リレーションシップ、パブリック プレビューの読み取りと書き込み
変化します。
自動化によりさらにエージェント API キーを作成する必要がある場合にのみ管理者。
ダッシュボードと設定ページでは、コピー可能なエージェント セットアップ プロンプトが生成されます。の
プレビューではキーがマスクされます。コピーエンドポイントは完全なキーを返し、
キャッシュ制御: ストアなし。
ローカル開発の場合、設定値は次のとおりです。
行セット MCP URL: http://localhost:8000/mcp/
行セット REST API ベース: http://localhost:8000/api/
Rowset セットアップ スキル: http://localhost:8000/skills/rowset-setup/SKILL.md
行セットスキル: http://localhost:8000/SKILL.md
コピーした API キーをプライベート環境変数に保存します。
import ROWSET_API_KEY= " コピーしたキーと置き換える "
REST 認証を確認します。
カール -H " 権限: ベアラー $ROWSET_API_KEY " \
http://localhost:8000/api/user
エージェントのゴールデン パス
Rowset の主なワークフローはエージェントのハンドオフであり、手動による行編集ではありません。
推奨されるエージェントの起動順序:
完全な API キーを ROWSET_API_KEY として非公開で保存します。
ベアラー トークン認証を使用してリモート MCP サーバーを構成します。
新しい接続または失敗した接続の場合は、get_user_info を 1 回呼び出して確認します。
認証、完全なオンボーディング、資格情報の問題の診断を行います。
ユーザーのタスクを開始します。現在の操作にライブ ツール スキーマを使用し、
get_rows を呼び出す

et_capabilities は馴染みのない機能のみ、または
トラブルシューティング。関連するトピックのみをリクエストします。
ユーザーがデータセット キーまたは URL を指定した場合は、get_dataset を直接呼び出します。もし
関連するデータセットが不明な場合は、search_datasets を 3 の制限で使用します。
結果を選択し、行操作の前に get_dataset を呼び出します。
セッションが開始されたからといって、機能をロードしたり、データセットをリストしたりしないでください。
検出中に無関係なデータセットやプロジェクトを列挙しないでください。
Codex/OpenClaw 互換クライアントの場合:
codex mcp 行セットを追加 \
--url http://localhost:8000/mcp/ \
--bearer-token-env-var ROWSET_API_KEY
運用環境では、URL を次のように置き換えます。
https://rowset.lvtd.dev/mcp/
生の API キーを MCP サーバー設定に含めないでください。キーを次の場所に保管します
エージェントのプライベート ランタイム環境またはシークレット ストアを設定し、クライアントを次のように構成します。
送信:
権限: ベアラー <キー>
MCPツールグループ
ライブ MCP サーバーは、ツール スキーマの正確なソースです。現在のワークフロー
グループは次のとおりです。
エージェントは、行の削除、データセットなどの破壊的なアクションの前に質問する必要があります。
ユーザーがアーカイブ、プロジェクト アーカイブ、またはパブリック プレビュー パスワードをクリアしない限り、
明示的にそのアクションを要求しました。
便利な Rowset ドッグフード パターンは、 task_id によってインデックス付けされたタスク ボードです。
{
"name" : " エージェント タスク ボード " ,
"description" : " 1 つのエージェント ワークフロー用の永続タスク ボード " ,
"instructions" : " task_id を安定させてください。definition_of_done が満たされた場合にのみステータスを完了に移動します。 " ,
「メタデータ」: {
"status_order" : [ " todo " 、 " 実行中 " 、 " ブロック済み " 、 " レビュー " 、 " 完了 " ],
"優先順位の意味" : {
"P0" : " 最高のレバレッジまたはブロッキング " 、
"P1" : " 現在サイクルの重要な作業 "
}
}、
"ヘッダー" : [
" task_id " ,
「ステータス」、
「優先度」、
「タスク」、
" 完了の定義 " ,
「所有者」、
" updated_on " ,
「メモ」
]、
"index_column" : " task_id " 、
「列の種類」

: {
"task_id" : " テキスト " ,
「ステータス」: {
"タイプ" : "選択" ,
"選択" : [ " todo " 、 " 実行中 " 、 " ブロック済み " 、 " レビュー " 、 " 完了 " ]
}、
「優先度」: {
"タイプ" : "選択" ,
"選択肢" : [ "P0" 、 "P1" 、 "P2" 、 "P3" ]
}、
"updated_on" : " 日付 "
}
}
この形状は、安定したインデックス、メタデータの選択、Rowset のコアを示しています。
永続的な命令、JSON 規則、インデックスによる更新。
http://localhost:8000/api/
本番環境では次のようになります。
https://rowset.lvtd.dev/api/
生成された API ドキュメントは以下から提供されます。
/api/docs
プライベート REST リクエストにはベアラー認証を使用します。
権限: ベアラー <キー>
プライベート REST リクエストは、API キーをベアラー トークンとしてのみ受け入れます。 URL 内の API キー
または代替ヘッダーが拒否されます。
カール -H " 権限: ベアラー $ROWSET_API_KEY " \
http://localhost:8000/api/user
データセットを作成する
curl -X POST http://localhost:8000/api/datasets \
-H " 認可: ベアラー $ROWSET_API_KEY " \
-H " Content-Type: application/json " \
-d ' {
"名前": "製品",
"description": "エージェントによって管理されるサプライヤー カタログ",
"instructions": "sku を安定させてください。行に別途記載がない限り、価格を USD として扱います。",
"ヘッダー": ["SKU"、"名前"、"価格"、"ステータス"]、
"index_column": "sku",
"column_types": {
"sku": "テキスト",
「名前」

[切り捨てられた]

## Original Extract

Rowset gives your AI agent a dataset backend. Contribute to LVTD-LLC/rowset development by creating an account on GitHub.

GitHub - LVTD-LLC/rowset: Rowset gives your AI agent a dataset backend. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
LVTD-LLC
/
rowset
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
645 Commits 645 Commits .agents/ skills .agents/ skills .cursor/ rules .cursor/ rules .github .github .greptile .greptile .seo .seo apps apps cli cli deployment deployment docs docs evaluations evaluations frontend frontend rowset rowset scripts scripts .browserslistrc .browserslistrc .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .nvmrc .nvmrc .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version .stylelintrc.json .stylelintrc.json AGENTS.md AGENTS.md ANALYTICS.md ANALYTICS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md DESIGN.md DESIGN.md Dockerfile-python Dockerfile-python GEMINI.md GEMINI.md Makefile Makefile PRODUCT.md PRODUCT.md README.md README.md SELF_HOSTING.md SELF_HOSTING.md STRUCTURE.md STRUCTURE.md TECH.md TECH.md VISION.md VISION.md conftest.py conftest.py djass-manifest.json djass-manifest.json docker-compose-local.yml docker-compose-local.yml docker-compose-prod.yml docker-compose-prod.yml docker-compose-test.yml docker-compose-test.yml eslint.config.js eslint.config.js manage.py manage.py package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js project-metadata.json project-metadata.json pyproject.toml pyproject.toml pytest.ini pytest.ini render.yaml render.yaml tailwind.config.js tailwind.config.js uv.lock uv.lock View all files Repository files navigation
Rowset is a private MCP and REST backend for structured datasets that trusted AI
agents can create, inspect, update, export, and share. Users sign in, copy an
agent setup prompt, authorize a scoped API key, and let the agent work with
owned datasets through stable programmatic interfaces instead of browser
automation.
Hosted Streamable HTTP MCP server for AI-agent workflows.
Authenticated REST API for account checks, projects, datasets, rows, exports,
relationships, image assets, and public preview settings.
Go rowset under cli/ for the same authenticated REST operations.
API-backed datasets with stable headers, semantic column metadata, persistent
agent instructions, JSON metadata, and an explicit index column.
Row CRUD by internal Rowset row id or by dataset index value.
Projects and project sections for organizing related datasets without changing
authentication boundaries.
Choice, reference, image, date, datetime, currency, number, boolean, email,
URL, and text column metadata.
Read-only public previews with optional password protection for human review.
CSV, JSONL, XLSX, SQLite, and dashboard-oriented Parquet export paths.
Private image asset storage on local disk or S3-compatible storage such as
Cloudflare R2.
Optional Qdrant-backed hybrid vector and lexical search for dataset rows.
Area
Technology
Language
Python 3.14.2 ( .python-version , pyproject.toml ) and Go for cli/
Backend
Django 6
REST API
Django Ninja
MCP
FastMCP mounted through Starlette in rowset/asgi.py
Auth
Django allauth, session auth, API-key auth, hosted MCP bearer auth
Data stores
PostgreSQL, Redis
Background jobs
Django Q2 workers
Tabular work
Python csv , json , sqlite3 , zipfile , plus Polars
Frontend
Django templates, HTMX, Alpine.js, Tailwind, PostCSS
Assets
Custom Node 24 build script in scripts/build-assets.mjs
Local stack
Docker Compose with Postgres, Redis, backend, workers, frontend, Mailhog, Stripe CLI, MJML, and MinIO
Observability
Sentry and PostHog
Integrations
Mailgun, Buttondown, Stripe, Chatwoot, S3-compatible storage, Qdrant/OpenRouter for optional vector search
Active deployment path
Docker images plus CapRover GitHub Actions
Product Boundaries
Rowset is intentionally centered on agent-managed datasets.
A signed-in user copies a Rowset setup prompt into a trusted agent.
The agent stores the API key privately and connects to Rowset MCP with
Authorization: Bearer <key> .
The agent creates or discovers datasets, inspects schema/context, mutates rows,
manages projects, exports snapshots, or enables a public preview when asked.
The dashboard helps humans with setup, settings, recent dataset state, schema
review, exports, public preview review, and account recovery.
Out of scope for the current product path:
Rowset-owned source connectors, sync, or write-back.
Public previews as authentication or as a substitute for REST/MCP access.
Browser automation as the preferred agent integration.
Broad BI, warehouse, or ETL orchestration promises.
Agents can still read local files, Google Sheets, databases, or other upstream
sources with their own capabilities, then send structured rows into Rowset
through MCP or REST.
For the supported local workflow:
Docker Desktop or Docker Engine with Docker Compose.
For host-side debugging outside Docker:
Node.js 24.11 or newer and npm 11 or newer.
Go 1.26 or newer when building the rowset CLI from source.
PostgreSQL and Redis reachable from your environment.
Most contributors should start with Docker Compose. The local Compose stack
builds the Python image, installs Node dependencies in the frontend service, and
runs Postgres and Redis for you.
git clone https://github.com/LVTD-LLC/rowset.git
cd rowset
2. Create local environment configuration
cp .env.example .env
The checked-in defaults are designed for the local Docker Compose stack:
Postgres database/user/password: rowset
Site URL: http://localhost:8000
docker compose -f docker-compose-local.yml up -d --build
backend logs for the backend service
http://localhost:8000
Mailhog is available at:
http://localhost:8025
MinIO's console is available at:
http://localhost:9001
4. Create an account
Use the local app UI to sign up. Email verification is non-blocking in the
current app: local confirmation links are captured by Mailhog or printed through
the configured email backend.
Use the smallest permission level that fits the agent:
Read for inspection and exports.
Read + write for dataset, row, project, relationship, and public preview
changes.
Admin only when automation must create more agent API keys.
The dashboard and settings pages generate a copyable agent setup prompt. The
preview masks the key; the copy endpoint returns the full key and uses
Cache-Control: no-store .
For local development, the setup values are:
Rowset MCP URL: http://localhost:8000/mcp/
Rowset REST API base: http://localhost:8000/api/
Rowset setup skill: http://localhost:8000/skills/rowset-setup/SKILL.md
Rowset skill: http://localhost:8000/SKILL.md
Store the copied API key in a private environment variable:
export ROWSET_API_KEY= " replace-with-your-copied-key "
Verify REST authentication:
curl -H " Authorization: Bearer $ROWSET_API_KEY " \
http://localhost:8000/api/user
Agent Golden Path
Rowset's primary workflow is agent handoff, not manual row editing.
Recommended agent startup order:
Store the full API key privately as ROWSET_API_KEY .
Configure the remote MCP server with bearer-token auth.
For a new or failing connection, call get_user_info once to verify
authentication, complete onboarding, and diagnose credential problems.
Start the user's task. Use live tool schemas for the operation at hand and
call get_rowset_capabilities only for an unfamiliar feature or
troubleshooting, requesting only the relevant topics.
If the user supplied a dataset key or URL, call get_dataset directly. If
the relevant dataset is unknown, use search_datasets with a limit of 3,
select a result, then call get_dataset before row operations.
Do not load capabilities or list datasets merely because a session started.
Do not enumerate unrelated datasets or projects during discovery.
For Codex/OpenClaw-compatible clients:
codex mcp add rowset \
--url http://localhost:8000/mcp/ \
--bearer-token-env-var ROWSET_API_KEY
For production, replace the URL with:
https://rowset.lvtd.dev/mcp/
Do not put the raw API key in the MCP server config. Store the key in the
agent's private runtime environment or secret store and configure the client to
send:
Authorization: Bearer <key>
MCP tool groups
The live MCP server is the exact source for tool schemas. The current workflow
groups are:
Agents should ask before destructive actions such as row deletion, dataset
archive, project archive, or clearing a public preview password unless the user
explicitly requested that action.
A useful Rowset dogfood pattern is a task board indexed by task_id :
{
"name" : " Agent Task Board " ,
"description" : " Durable task board for one agent workflow " ,
"instructions" : " Keep task_id stable. Move status to done only after definition_of_done is satisfied. " ,
"metadata" : {
"status_order" : [ " todo " , " doing " , " blocked " , " review " , " done " ],
"priority_meaning" : {
"P0" : " Highest leverage or blocking " ,
"P1" : " Important current-cycle work "
}
},
"headers" : [
" task_id " ,
" status " ,
" priority " ,
" task " ,
" definition_of_done " ,
" owner " ,
" updated_on " ,
" notes "
],
"index_column" : " task_id " ,
"column_types" : {
"task_id" : " text " ,
"status" : {
"type" : " choice " ,
"choices" : [ " todo " , " doing " , " blocked " , " review " , " done " ]
},
"priority" : {
"type" : " choice " ,
"choices" : [ " P0 " , " P1 " , " P2 " , " P3 " ]
},
"updated_on" : " date "
}
}
That shape demonstrates the Rowset core: stable index, choice metadata,
persistent instructions, JSON conventions, and updates by index.
http://localhost:8000/api/
In production, it is:
https://rowset.lvtd.dev/api/
Generated API docs are served from:
/api/docs
Use bearer auth for private REST requests:
Authorization: Bearer <key>
Private REST requests accept API keys only as bearer tokens. API keys in URLs
or alternate headers are rejected.
curl -H " Authorization: Bearer $ROWSET_API_KEY " \
http://localhost:8000/api/user
Create a dataset
curl -X POST http://localhost:8000/api/datasets \
-H " Authorization: Bearer $ROWSET_API_KEY " \
-H " Content-Type: application/json " \
-d ' {
"name": "Products",
"description": "Supplier catalog managed by an agent",
"instructions": "Keep sku stable. Treat price as USD unless a row says otherwise.",
"headers": ["sku", "name", "price", "status"],
"index_column": "sku",
"column_types": {
"sku": "text",
"name"

[truncated]
