---
source: "https://github.com/sanniheruwala/RedNotebookAI/blob/main/README.md"
hn_url: "https://news.ycombinator.com/item?id=48503888"
title: "Show HN: RedNotebook AI open-source AI data notebook for Trino, +12 SQL engines"
article_title: "RedNotebookAI/README.md at main · sanniheruwala/RedNotebookAI · GitHub"
author: "heruwala"
captured_at: "2026-06-12T16:08:51Z"
capture_tool: "hn-digest"
hn_id: 48503888
score: 1
comments: 0
posted_at: "2026-06-12T13:34:14Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RedNotebook AI open-source AI data notebook for Trino, +12 SQL engines

- HN: [48503888](https://news.ycombinator.com/item?id=48503888)
- Source: [github.com](https://github.com/sanniheruwala/RedNotebookAI/blob/main/README.md)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T13:34:14Z

## Translation

タイトル: HN を表示: RedNotebook AI Trino 用オープンソース AI データ ノートブック、+12 SQL エンジン
記事タイトル: RedNotebookAI/README.md at main · sanniheruwala/RedNotebookAI · GitHub
説明: RedNotebook AI : オープンソースの AI データ ノートブック。 AI SQL/チャート提案、プロファイリング、PII 検出、および NotebookLM スタイルのナレッジ レイヤーを備えた、FastAPI バックエンド上のノートブック ファースト UI (Next.js + shadcn/ui)。 - メインの RedNotebookAI/README.md · sanniheruwala/RedNotebookAI

記事本文:
RedNotebookAI/README.md メイン · sanniheruwala/RedNotebookAI · GitHub
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
サンニヘルワラ
/
レッドノートAI
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする Blame More ファイル ac

問題の原因 さらなるファイルアクション 最新のコミット
履歴 履歴 289 行 (213 loc) · 11.1 KB メイン ブレッドクラム
トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション
レッドノートAI
Trino、DuckDB、およびその他 11 個の SQL エンジン用のオープンソース AI データ ノートブック。
RedAnalytica による。
美しいグラフ、AI の提案、NotebookLM スタイルのナレッジ レイヤーを使用して、データのクエリ、視覚化、プロファイリング、探索を行います。
現代のデータチームは、1 つの質問に答えるために 5 つのツールの間を行き来します。 RedNotebook AI はすべてを 1 つのノートブックにまとめます。
Monaco、AG Grid、ドラッグして並べ替えるセル、およびキーボード ショートカットを備えた実際の SQL ワークスペース。
ブランドを意識したテーマを備えた Apache ECharts を利用したプレミアム チャート。
信頼できる AI、OpenAI、Anthropic、Ollama、または決定論的なオフライン モック間でプラグイン可能。デフォルトでプライバシーが保護され、スキーマのみのコンテキスト、PII マスキング、シークレットが削除されます。
NotebookLM スタイルの知識層。 SQL、スキーマ、結果、チャートをソースのノートブックに取り込みます。 [n] 個の引用チップを使用して根拠のある質問をしましょう。インフォグラフィックスと Studio の説明 (概要 / FAQ / 学習ガイド / 次の質問の提案) を生成します。
ドラッグ アンド ドロップでファイルをアップロードします。 CSV、TSV、Parquet、または JSON ファイルをアプリ内の任意の場所にドロップします。DuckDB はそれをクエリ可能なテーブルとして即座に添付します (SELECT * FROM 顧客は Just Works)。
ワンクリックで公開。アカウント不要の公開共有リンクをノートブックから作成します。公開されたページは自己完結型の HTML スナップショットであり、ライブ データがマシンから流出することはありません。
デフォルトでは読み取り専用です。 sqlglot による SQL ガードは、書き込みを明示的に有効にしない限り、破壊的なステートメントをブロックします。
地元第一主義。ログインせずにラップトップ上で実行できます。単一の環境変数 ( AUTH_ENABLED=true ) を反転して、ローカルの電子メール + パスワード、GitHub OAuth、API トークン、ユーザーごとの名前空間を使用したマルチユーザー モードを有効にします。

そして管理者が招待します。
docker run -d --name rednotebook \
-p 8000:8000 \
-v rednotebook-data:/data \
ghcr.io/sanniheruwala/rednotebook-ai:最新
次に、 http://localhost:8000 を開きます。
cp .env.example .env # 必要に応じて編集します
ドッカー構成 -d
パイソン
pip install rednotebook-ai # PyPI から (リリースがタグ付けされている場合)
rednotebook run # :8000 で FastAPI サーバーを起動します
次に、2 番目の端末で次のようにします。
CDフロントエンド
npmインストール
npm run dev # starts the dev UI on :3000
ソースから
git clone https://github.com/sanniheruwala/RedNotebookAI.git
cd RedNotebookAI
python -m venv .venv && ソース .venv/bin/activate
pip install -e " .[dev] "
cp .env.example .env
レッドノートブックの実行
# 別の端末で
cd フロントエンド && npm install && npm run dev
どこで安全に実行できますか?
RedNotebook AI はローカルファーストです。今日:
完全なセキュリティ モデルについては、docs/deployment.md を参照してください。
In the UI top bar, click Configure connection . 13 個のコネクタが同梱されています
ボックス — 追加の pip インストール手順、ドライバーのセットアップ、ODBC ダンスは必要ありません。
方言ごとの完全なフィールドについては、docs/connectors.md を参照してください。
参考。
クイックスタート: DuckDB (サーバーなし、インスタント)
デフォルト。 Pick "DuckDB (no server)" in the dialog. 2 つのモード:
In-memory ( :memory: ) — ephemeral playground.ローカル ファイルに対する 1 回限りの SQL に最適: SELECT * FROM read_csv_auto('orders.csv') WHERE …
File ( ./local.duckdb ) — persistent.シングルユーザー ウェアハウスのように使用します: CREATE TABLE Customers (…) 、 INSERT … など。
必要に応じて、「作業ディレクトリ」を設定して、read_csv_auto / read_parquet の相対ファイル パスが期待する場所に解決されるようにします。
実際のデータ ウェアハウスでのチーム分析の場合は、UI ダイアログに入力するか、
.env のデフォルト:
TRINO_HOST = trino.example.com
トリノポート = 443
TRINO_SCHEME = https
TRINO_USER = アリス
TRINO_PASSWORD = ...
TRINO_CATALOG = ハイブ
TRINO_SCHEMA = デフォルト
TRINO_VERIFY_SSL

=本当
カスタム HTTP ヘッダー、セッション プロパティ、クエリ タイムアウト、結果制限がすべてサポートされています。
プロバイダー
セットアップ
モック (デフォルト)
オフライン、決定的。セットアップはありません。
OpenAI
AI_PROVIDER=openai 、OPENAI_API_KEY=sk-…
人間的
AI_PROVIDER=人間 、ANTHROPIC_API_KEY=sk-ant-…
オラマ (地元)
AI_PROVIDER=ollam 、OLLAMA_BASE_URL=http://localhost:11434
プライバシーのデフォルト:
AI_ALLOW_SAMPLE_ROWS=true でない限り、サンプル行は AI に送信されません。
サンプルが共有される場合、PII 列はマスクされます。
シークレットは、プロバイダー呼び出しの前に SQL から削除されます。
認証情報が AI に転送されることはありません。
AUTH_ENABLED = true
SECRET_KEY = $( openssl rand -hex 32 )
COOKIE_SECURE = true # HTTPS の背後にある場合は true を設定します
ALLOW_SELF_SIGNUP = false # デフォルトでは admin-invite のみ
最初の登録はワークスペース管理者になります。後続のユーザーには招待が必要です ( POST /api/auth/invite )。 GitHub OAuth および API トークン (PAT スタイル) は、すぐに使用できるようにサポートされています。 docs/deployment.md を参照してください。
レイヤー
技術
バックエンド
Python 3.11+、FastAPI、Pydantic、Trino クライアント、SQLAlchemy + バンドルされたドライバー (Postgres、MySQL、MSSQL/ODBC、Snowflake、BigQuery、Redshift、Oracle、ClickHouse、Databricks など)、DuckDB、Pandas、ECharts/Plotly
フロントエンド
Next.js 14、TypeScript、Tailwind、shadcn/ui、Monaco、AG Grid、ECharts、framer-motion、@dnd-kit
州
TanStack Query (サーバー) + Zustand (ローカル)
認証
ローカルメール + パスワード (bcrypt) + JWT Cookie、GitHub OAuth、API トークン
AI
プロバイダープラグイン可能 (モック、OpenAI、Anthropic、Ollama)
ストレージ
ノートブック/ナレッジ/ユーザーのローカル JSON。オプションの Parquet 結果キャッシュ
rednotebook/Python バックエンド (FastAPI + コア ライブラリ)
§── auth/ ユーザー ストア、JWT セッション、パスワード ハッシュ、OAuth、API トークン
§── サーバー/ FastAPI アプリ + ルーター
§── コネクタ/ Trino + DuckDB + 11 個の SQLAlchemy 方言 + レジストリ
§── ai/Prov

ider 抽象化 (モック、オープンアイ、人類、オラマ)
§── ノートブック/ ノートブック モデル、JSON ストレージ、ガードを意識したランナー
§── ナレッジ/ NotebookLM スタイルの内部ナレッジ レイヤー
§── ビジュアライゼーション/レコメンダー、チャート仕様、HTML インフォグラフィック ジェネレーター
§── プロファイリング/統計 + PII 検出器
§── セキュリティ/SQLガード、シークレットマスキング
§── マイグレーション/ワンショットデータマイグレーション
━── cli/タイパーCLI
フロントエンド/ Next.js + Tailwind + shadcn/ui
docs/ アーキテクチャ、AI、セキュリティ、展開、コネクタ、ロードマップ
テスト/pytest テストスイート
完全なアーキテクチャの説明。
ナレッジレイヤー + NotebookLM の統合
# バックエンド
pytest # 56+ テスト
ラフチェック。
# フロントエンド
CDフロントエンド
npm タイプチェックを実行する
npmでlintを実行する
npm ビルドを実行する
継続的インテグレーションは、すべてのプッシュと PR で完全なスイートを実行します。 .github/workflows を参照してください。
私たちは標準的なオープンソース フローに従います。短いバージョン:
まず問題を開きます。を使用します。
バグレポート
または
機能リクエスト
テンプレート。問題がリンクされていないドライブバイ PR は、審査なしに閉鎖される場合があります。
ローカルでフォーク、分岐、書き込み、チェックを実行する
( pytest 、 ruff check 。 、 npm run lint && npm run typecheck && npm run build )。
問題を参照する PR を開きます (クローズ #123)。
メンテナはマージ前にレビューと承認を行います。
main は保護されたブランチです - 直接プッシュはブロックされ、すべての変更が行われます
✅ グリーン CI と ✅ の承認が必要です
コードオーナー。管理者であっても例外はありません。
完全なフローについては docs/contributing.md を参照してください。
ブランチの命名規則、私たちが「ノー」と言うもの、そしてメンテナ
権利。セキュリティの脆弱性については、次を使用します。
非公開であり、決して公の問題ではありません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

RedNotebook AI : open-source AI data notebook. Notebook-first UI (Next.js + shadcn/ui) over a FastAPI backend, with AI SQL/chart suggestions, profiling, PII detection, and a NotebookLM-style knowledge layer. - RedNotebookAI/README.md at main · sanniheruwala/RedNotebookAI

RedNotebookAI/README.md at main · sanniheruwala/RedNotebookAI · GitHub
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
sanniheruwala
/
RedNotebookAI
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 289 lines (213 loc) · 11.1 KB main Breadcrumbs
Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions
RedNotebook AI
The open-source AI data notebook for Trino, DuckDB, and 11 more SQL engines.
By RedAnalytica .
Query, visualize, profile, and explore data with beautiful charts, AI suggestions, and a NotebookLM-style knowledge layer.
Modern data teams jump between five tools to answer one question. RedNotebook AI puts all of it in one notebook:
A real SQL workspace with Monaco, AG Grid, drag-to-reorder cells, and keyboard shortcuts.
Premium charts powered by Apache ECharts with brand-aware theming.
AI you can trust , pluggable across OpenAI, Anthropic, Ollama, or a deterministic offline mock. Privacy-safe by default, schema-only context, PII masking, secrets stripped.
NotebookLM-style knowledge layer. Pull SQL, schemas, results, and charts into a notebook of sources. Ask grounded questions with [n] citation chips. Generate infographics and a Studio briefing (overview / FAQ / study guide / suggested next questions).
Drag-and-drop file uploads. Drop a CSV, TSV, Parquet, or JSON file anywhere in the app — DuckDB attaches it instantly as a queryable table ( SELECT * FROM customers Just Works).
One-click publish. Mint a public, no-account-needed share link from any notebook. The published page is a self-contained HTML snapshot — your live data never leaves your machine.
Read-only by default. A SQL guard backed by sqlglot blocks destructive statements unless you explicitly enable writes.
Local-first. Runs on your laptop with no login. Flip a single env var ( AUTH_ENABLED=true ) to enable multi-user mode with local email+password, GitHub OAuth, API tokens, per-user namespacing, and admin invites.
docker run -d --name rednotebook \
-p 8000:8000 \
-v rednotebook-data:/data \
ghcr.io/sanniheruwala/rednotebook-ai:latest
Then open http://localhost:8000 .
cp .env.example .env # edit as needed
docker compose up -d
Python
pip install rednotebook-ai # from PyPI (when a release is tagged)
rednotebook run # starts the FastAPI server on :8000
Then in a second terminal:
cd frontend
npm install
npm run dev # starts the dev UI on :3000
From source
git clone https://github.com/sanniheruwala/RedNotebookAI.git
cd RedNotebookAI
python -m venv .venv && source .venv/bin/activate
pip install -e " .[dev] "
cp .env.example .env
rednotebook run
# in another terminal
cd frontend && npm install && npm run dev
Where can I run this safely?
RedNotebook AI is local-first. Today:
See docs/deployment.md for the full security model.
In the UI top bar, click Configure connection . 13 connectors ship in
the box — no extra pip install step, no driver setup, no ODBC dance.
See docs/connectors.md for the full per-dialect field
reference.
Quick start: DuckDB (no server, instant)
The default. Pick "DuckDB (no server)" in the dialog. Two modes:
In-memory ( :memory: ) — ephemeral playground. Great for one-off SQL against local files: SELECT * FROM read_csv_auto('orders.csv') WHERE …
File ( ./local.duckdb ) — persistent. Use it like a single-user warehouse: CREATE TABLE customers (…) , INSERT … , etc.
Optionally set a "Working directory" so relative file paths in read_csv_auto / read_parquet resolve where you expect.
For team analytics on real data warehouses, fill in the UI dialog or set
defaults in .env :
TRINO_HOST = trino.example.com
TRINO_PORT = 443
TRINO_SCHEME = https
TRINO_USER = alice
TRINO_PASSWORD = ...
TRINO_CATALOG = hive
TRINO_SCHEMA = default
TRINO_VERIFY_SSL = true
Custom HTTP headers, session properties, query timeouts, and result limits are all supported.
Provider
Setup
Mock (default)
Offline, deterministic. No setup.
OpenAI
AI_PROVIDER=openai , OPENAI_API_KEY=sk-…
Anthropic
AI_PROVIDER=anthropic , ANTHROPIC_API_KEY=sk-ant-…
Ollama (local)
AI_PROVIDER=ollama , OLLAMA_BASE_URL=http://localhost:11434
Privacy defaults:
Sample rows are not sent to AI unless AI_ALLOW_SAMPLE_ROWS=true .
PII columns are masked when samples are shared.
Secrets are stripped from SQL before any provider call.
Credentials are never forwarded to AI.
AUTH_ENABLED = true
SECRET_KEY = $( openssl rand -hex 32 )
COOKIE_SECURE = true # set true when behind HTTPS
ALLOW_SELF_SIGNUP = false # admin-invite only by default
The first registration becomes the workspace admin. Subsequent users need an invite ( POST /api/auth/invite ). GitHub OAuth and API tokens (PAT-style) are supported out of the box. See docs/deployment.md .
Layer
Tech
Backend
Python 3.11+, FastAPI, Pydantic, Trino client, SQLAlchemy + bundled drivers (Postgres, MySQL, MSSQL/ODBC, Snowflake, BigQuery, Redshift, Oracle, ClickHouse, Databricks, ...), DuckDB, Pandas, ECharts/Plotly
Frontend
Next.js 14, TypeScript, Tailwind, shadcn/ui, Monaco, AG Grid, ECharts, framer-motion, @dnd-kit
State
TanStack Query (server) + Zustand (local)
Auth
Local email+password (bcrypt) + JWT cookies, GitHub OAuth, API tokens
AI
Provider-pluggable (mock, OpenAI, Anthropic, Ollama)
Storage
Local JSON for notebooks/knowledge/users; optional Parquet result cache
rednotebook/ Python backend (FastAPI + core libs)
├── auth/ User store, JWT sessions, password hashing, OAuth, API tokens
├── server/ FastAPI app + routers
├── connectors/ Trino + DuckDB + 11 SQLAlchemy dialects + registry
├── ai/ Provider abstraction (mock, openai, anthropic, ollama)
├── notebook/ Notebook models, JSON storage, guard-aware runner
├── knowledge/ NotebookLM-style internal knowledge layer
├── visualization/ Recommender, chart spec, HTML infographic generator
├── profiling/ Stats + PII detector
├── security/ SQL guard, secret masking
├── migrations/ One-shot data migrations
└── cli/ Typer CLI
frontend/ Next.js + Tailwind + shadcn/ui
docs/ Architecture, AI, security, deployment, connectors, roadmap
tests/ pytest test suite
Full architecture write-up .
Knowledge layer + NotebookLM integration
# Backend
pytest # 56+ tests
ruff check .
# Frontend
cd frontend
npm run typecheck
npm run lint
npm run build
Continuous integration runs the full suite on every push and PR. See .github/workflows .
We follow the standard open-source flow. The short version:
Open an issue first. Use the
bug report
or
feature request
templates. Drive-by PRs with no linked issue may be closed without review.
Fork, branch, write, run the checks locally
( pytest , ruff check . , npm run lint && npm run typecheck && npm run build ).
Open a PR referencing the issue ( Closes #123 ).
A maintainer reviews and approves before merge.
main is a protected branch — direct pushes are blocked, every change
needs ✅ green CI and ✅ approval from a
CODEOWNER . No exceptions, even for admins.
See docs/contributing.md for the full flow,
branch-naming conventions, what we say "no" to, and the maintainer
rights. For security vulnerabilities, use
private disclosure , never a public issue.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
