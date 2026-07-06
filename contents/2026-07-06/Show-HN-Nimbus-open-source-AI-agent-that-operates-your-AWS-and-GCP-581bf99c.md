---
source: "https://github.com/hritvikgupta/nimbus"
hn_url: "https://news.ycombinator.com/item?id=48811086"
title: "Show HN: Nimbus - open-source AI agent that operates your AWS and GCP"
article_title: "GitHub - hritvikgupta/nimbus: A opensource cloud AI agent to manage and build your cloud infrastructure · GitHub"
author: "hritvik29"
captured_at: "2026-07-06T22:59:09Z"
capture_tool: "hn-digest"
hn_id: 48811086
score: 1
comments: 0
posted_at: "2026-07-06T22:02:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Nimbus - open-source AI agent that operates your AWS and GCP

- HN: [48811086](https://news.ycombinator.com/item?id=48811086)
- Source: [github.com](https://github.com/hritvikgupta/nimbus)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T22:02:44Z

## Translation

タイトル: Show HN: Nimbus - AWS と GCP を操作するオープンソース AI エージェント
記事のタイトル: GitHub - hritvikgupta/nimbus: クラウド インフラストラクチャを管理および構築するためのオープンソース クラウド AI エージェント · GitHub
説明: クラウド インフラストラクチャを管理および構築するためのオープンソース クラウド AI エージェント - hritvikgupta/nimbus

記事本文:
GitHub - hritvikgupta/nimbus: クラウド インフラストラクチャを管理および構築するためのオープンソース クラウド AI エージェント · GitHub
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
フリトヴィクグプタ
/
ニンバス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .github .github cli cli デプロイ デプロイ docs-site docs-site docs docs public public scripts スクリプト サーバー サーバー src src .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .nvmrc .nvmrc CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md DCO DCO Dockerfile Dockerfile ライセンス ライセンス通知 NOTICE README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-NOTICES.md THIRD-PARTY-NOTICES.md TRADEMARKS.md TRADEMARKS.md docker-compose.yml docker-compose.yml Index.html Index.html package-lock.json package-lock.json package.json package.json vite.config.js vite.config.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI を活用したクラウド コントロール プレーン。コードを読んで理解する 1 人のエージェント
アーキテクチャ、実際の AWS および GCP 認証情報に基づいて動作し、リポジトリを修正します。
素朴な会話。
クラウド コンソール、ターミナル、CI ダッシュボード、IaC ファイルを切り替える代わりに、
インフラストラクチャの設計、展開、監視、修復を行えるエージェントを 1 か所に配置
その背後にあるコード。チーム チャット内に常駐するシニア プラットフォーム エンジニアのようなものだと考えてください。
Nimbus は 3 つのアイデアに基づいて構築されています。
インテント駆動型の操作 — 望む結果を説明します。ニンバスが計画を立てる、あなたは
承認すると実行されます。読み取りは自律的です。すべての書き込みが確認されます。
プロジェクトごとに 1 つの共有ワークスペース — チャネル、クラウド接続、マシン、修理、
アーキテクチャ キャンバスはチームと共有されます。 Nimbus とのプライベート チャットのみが個人的なものとなります。
証拠を伴う実際の執行 - すべての主張はエージェントが実際に行った事柄に基づいています
読み取りまたは実行 (ファイル、ログ、ライブ リソース)

)、すべての変更は監査可能です。
💬 会話エージェント
ツールを使用した ReAct チャット ループ - インベントリの検査、ログとテレメトリの読み取り、コストの見積もり、デプロイ、PR のオープン。デザインモードではアーキテクチャをスケッチします。エージェントモードは実際のクラウドを操作します。
🎨 ライブ建築キャンバス
Nimbus にシステムの設計を依頼し、ノードがライブ (React Flow) でレンダリングされるのを観察します (ロード バランサー → アプリ → データベース → キャッシュ)。
☁️ マルチクラウド
ユーザーごと、プロジェクトごとの MCP サーバーを介して AWS (キーまたは IAM ロール、有効期間の短い STS) と GCP (サービス アカウント JSON または Google OAuth によるキーレス接続) を接続します。
🧭 ライブクラウドデータ
クラウドに接続すると、概要、リソース、コストが実際の在庫、テレメトリ、支出を取得します。
🛠️ 共有コンピューティングの修復
ラップトップまたは CI ランナーを接続します。 Nimbus はローカルのクロード コードをターンバイターンで駆動してリポジトリを修正し、ブランチをプッシュして PR を開きます。マシンはプロジェクトごとにプールされます。
🔌 GitHub など
Composio を介してリポジトリに接続します。エージェントは、行動する前にコード、PR、CI を読み取ります。
👥 チームと権限
電子メールで招待します。チャンネル、マシンと修理、クラウドとリソースをメンバーごとに切り替えます。
🔐 安全な設計
認証情報は保存時に暗号化され (AES-256-GCM)、印刷またはログに記録されることはありません。メンバーシップによる強制アクセス。サーバー側の監査証跡。
ライブ キャンバスでインフラストラクチャを設計する
オンデマンドでマシンをレンタルする · 実機のコードを修復する
Nimbus は、サービスにファンアウトするシン Express API と通信する React SPA です。
リポジトリ、エージェント、およびユーザーごとのクラウド MCP サーバー。別の CLI ワーカーが接続します
社外の機械を修理します。
┌─────────────┐ ┌───────────────────┐
│ Webアプリ（Vite/React）

│ /api │ APIサーバー（Express） │
│ ローカルホスト:5280 │ ────► │ ローカルホスト:8788 │
│ • チャット + キャンバス │ プロキシ │ │
│ • ダッシュボード │ │ ルート/ → HTTP ハンドラー (リソースごと) │
│ • 接続 │ │ サービス/ → クラウド、テレメトリ、コスト、仕様… │
━━━━━━━━━┘ │ リポジトリ/ → SQLite 永続化 │
│ エージェント/ → ReAct チャットループ + 分析 │
┌─────────────┐ │ tools/ → エージェントツール（キャンバス、修復…） │
│ 接続されたマシン │ ポーリング │ libs/ → モデル、MCP、composio、aws/gcp │
│ @nimbus/cli + クロード │ ◄────► │ mcp/ → aws / gcloud / クラウドラン MCP │
│ コード (あなたのラップトップ) │ (PR) lux──────────────┘
━━━━━━━━━━┘ │
SQLite (better-sqlite3) · 保存時に暗号化された認証情報
ユーザー ID はセッション Cookie から厳密に取得されます。 req.user.id は分離です
すべてのサービスと MCP 呼び出しがスコープされる境界。
フロントエンド — React 18、Vite 5、React Router、React Flow、@ai-sdk/react 、react-markdown
バックエンド — ノード ≥ 20、Express、better-sqlite3、ヘルメット、Zod
AI — Vercel AI SDK ( ai )、OpenRouter クロスプロバイダー フォールバックを備えた Databricks (プライマリ)、
モデル コンテキスト プロトコル ( @modelcontextprotocol/sdk )
統合 — Composio (GitHub)、AWS および GCP MCP サーバー、Fly.io (プロビジョニング / レンタルマシン)
Nimbus を実行する最速の方法。このイメージには、クラウド MCP サーバーに必要なものがすべてバンドルされています -
git 、 uv (AWS)、 gcloud CLI (GCP) および Node なので、他にインストールするものはありません。あなただけ
.env を指定します。
git

クローン https://github.com/hritvikgupta/nimbus.git
CDニンバス
cp .env.example .env
# 1. マスター暗号化キーを生成し、それを .env に追加します (保存された資格情報を保護します)。
echo " NIMBUS_ENC_KEY= $( openssl rand -base64 32 ) " >> .env
# 2. .env を開き、AI モデル — LLM_PROVIDER、LLM_MODEL、OPENROUTER_API_KEY を入力します。
# (オプション: マシンをレンタルする場合は FLY_API_TOKEN、GitHub の場合は COMPOSIO_API_KEY。)
#3. 実行します。
docker 構成 --build
# http://localhost:8788 を開きます
openssl はありませんか? node -e "console.log(require('crypto').randomBytes(32).toString('base64'))" を使用し、結果を NIMBUS_ENC_KEY= として .env に貼り付けます。
1 つのコンテナーは、ポート 8788 でアプリと API を提供します。データ (SQLite DB + 暗号化されたデータ)
接続) は、再起動後も nimbus-data ボリュームに保持されます。
プレーンな Docker の方が良いですか? docker build -t nimbus 。 && docker run -p 8788:8788 --env-file .env -v nimbus-data:/app/server/.data nimbus
代わりにソースから実行するには (開発用)、以下の手順に従います。
LLM プロバイダー - Databricks サービス提供エンドポイントまたは OpenRouter API キー
(オプション) GitHub の Composio API キー、および実際のクラウドに接続するための AWS/GCP 認証情報
クラウド MCP サーバーの場合: uv (AWS) および
gcloud CLI (GCP) — Docker に自動的にバンドルされます
git クローン https://github.com/hritvikgupta/nimbus.git
CDニンバス
npmインストール
2. 設定する
cp .env.example .env
.env を入力します (「構成」を参照)。少なくとも、NIMBUS_ENC_KEY と 1 つの
LLM プロバイダー:
# マスター暗号化キーを生成する
ノード -e " console.log(require('crypto').randomBytes(32).toString('base64')) "
3. クラウド MCP サーバーをセットアップする
エージェントは、以下で販売されている 3 つの MCP サーバーを通じてクラウドと通信します。
server/mcp/ — AWS ( aws-mcp 、 uvx 経由の Python )、Cloud Run
(cloud-run-mcp) と gcloud ( gcloud-mcp )。それらのソースはリポジトリ内にあります。彼らのをインストールしてください
d

依存関係を一度作成します (そして gcloud バンドルをビルドします)。
npm 実行セットアップ:mcp
これにより、ノードサーバーに対して npm install が実行され、 gcloud-mcp が構築され、
uv (AWS サーバーの起動に必要)。場合はこれをスキップできます
1 つのクラウドのみを使用します。対応するサーバーは、一致する接続が存在する場合にのみ生成されます。
ホストの前提条件 (MCP サーバーはこれらをシェルとして実行します。使用するクラウド用のものをインストールします):
インストールが完了すると、エージェントには完全な AWS および gcloud コマンド サーフェスと Cloud Run が備わります。
デプロイ/ログ — サーバー自体は、トリミングされたサブセットではなく、完全なアップストリーム実装です。
引き続き、アプリ内で独自のクラウド認証情報を接続します (接続)。
Nimbus は、API と Web アプリという 2 つのプロセスです。 2 つのターミナルで実行します。
npm run api # バックエンド → http://localhost:8788
npm run dev # フロントエンド → http://localhost:5280 (/api を :8788 にプロキシ)
http://localhost:5280 を開いてアカウントを作成すると、ガイド付きウォークスルーが表示されます。
ワークスペース。 Connections からリポジトリとクラウドに接続し、チャットを開始します。
すべての設定は、.env の環境変数を介して行われます ( npm run api によってロードされます)。
常に: NIMBUS_ENC_KEY + AI モデル ( LLM_PROVIDER + LLM_MODEL + OPENROUTER_API_KEY 、または Databricks トリオ)。プロバイダー、キー、請求書。
マシンをレンタルするには (Fly): 独自の FLY_API_TOKEN (+ オプションの FLY_* )。機能を無効にするには、空白のままにします。レンタルされたマシンにはパブリック URL は必要ありません。ワーカーはこのサーバー内で実行され、その API を介して Fly を駆動します。レンタルする人は、UI でレンタルごとに独自のコーディング エージェント キーを提供します。
修理のために自分のラップトップ/CI に接続するには: ここには何もありません。そのマシンは @nimbus/cli ワーカーを実行し、このサーバーをポーリングするため、Nimbus ( nimbus start <key> --url https://your-host ) をポイントしており、その逆ではありません。
公開をホストする場合

cly: APP_URL を実際の URL (GCP OAuth リダイレクト + リンクに使用) に設定し、オプションで WEBHOOK_BASE 、 CORS_ORIGINS 、 ALLOW_SIGNUP=false を設定します。
実際の .env を決してコミットしないでください。それは無視されます。 .env.example がテンプレートです。
修復は、@nimbus/cli ワーカーを介して所有する実際のマシン上で実行されます。投票する
Nimbus (アウトバウンドのみ - あらゆる NAT/ファイアウォールの背後で動作します)。修正がディスパッチされると、
リポジトリにアクセスし、ローカルのクロード コードを実行して問題を見つけて修正し、PR を開きます。
npm install -g @nimbus/cli
# アプリ内: 修復 → マシンの接続 → キーの生成
nimbus start <ワーカーキー>
マシンには Claude Code (インストール済み + ログイン済み)、 git 、および gh が必要です。参照
完全なワーカー ドキュメントについては、cli/README.md を参照してください。
ニンバス/
§── src/ # React SPA — ページ、コンポーネント、コンテキスト、スタイル
§── サーバー/ # Express API
│ §──server.mjs # Thin HTTP エントリ — ルーターをマウントします
│ §── Routes/ # HTTP ハンドラー (リソースごとに 1 つのファイル)
│ §── サービス/ # ビジネスロジック (クラウド、テレメトリ、コスト、スペック、修理…)
│ §── リポジトリ/ # 永続化 (SQLite via better-sqlite3)
│ §── エージェント/ # ReAct チャット ループ + コードベース分析エージェント
│ §── tools/ # エージェント ツール (キャンバス、修復、テレメトリ、コード…)
│ §── libs/ # 外部クライアント (モデル、MCP、composio、aws、gcp)
│ §── mcp/ # AWS / gcloud / Cloud Run MCP サーバー
│ └── ミドルウェア/ # セッション認証 + クラウド スコーピング
§── cli/ # @nimbus/cli — 接続されたマシンの修復作業者
§── docs-site/ # Fumadocs ドキュメント サイト (静的にエクスポート)
§──deploy/ # Fly.io デプロイ構成 (ランディング + ドキュメント)
§── docs/ # アーキテクチャと設計のドキュメント (運用準備、レンタル コンピューティング…)
└── public/ # 静的アセット (ロゴ、ブランド)
導入
Nimbus は Fly.io をターゲットにしています。 API サーバーも

静的にエクスポートされたものを保存します
リクエストホストが docs.* と一致するたびに docs サイト ( docs-site/out から)、つまり docs.yourdomain
個別の p は必要ありません

[切り捨てられた]

## Original Extract

A opensource cloud AI agent to manage and build your cloud infrastructure - hritvikgupta/nimbus

GitHub - hritvikgupta/nimbus: A opensource cloud AI agent to manage and build your cloud infrastructure · GitHub
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
hritvikgupta
/
nimbus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .github .github cli cli deploy deploy docs-site docs-site docs docs public public scripts scripts server server src src .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .nvmrc .nvmrc CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DCO DCO Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-NOTICES.md THIRD-PARTY-NOTICES.md TRADEMARKS.md TRADEMARKS.md docker-compose.yml docker-compose.yml index.html index.html package-lock.json package-lock.json package.json package.json vite.config.js vite.config.js View all files Repository files navigation
An AI-powered cloud control plane. One agent that reads your code, understands your
architecture, acts on real AWS & GCP credentials, and fixes your repos — all through
plain conversation.
Instead of switching between cloud consoles, terminals, CI dashboards and IaC files, you work
in one place with an agent that can design infrastructure, deploy it, watch it, and repair
the code behind it. Think of it as a senior platform engineer that lives inside your team chat.
Nimbus is built around three ideas:
Intent-driven operations — describe the outcome you want; Nimbus produces a plan, you
approve it, it executes. Reads are autonomous; every write is confirmed.
One shared workspace per project — channels, cloud connections, machines, repairs and the
architecture canvas are shared with your team. Only your private chat with Nimbus stays personal.
Real execution, with evidence — every claim is grounded in something the agent actually
read or ran (a file, a log, a live resource), and every change is auditable.
💬 Conversational agent
A ReAct chat loop with tool use — inspect inventory, read logs & telemetry, estimate cost, deploy, and open PRs. Design mode sketches architecture; Agent mode operates the real clouds.
🎨 Live architecture canvas
Ask Nimbus to design a system and watch the nodes render live (React Flow) — load balancer → app → database → cache.
☁️ Multi-cloud
Connect AWS (keys or IAM role, short-lived STS) and GCP (service-account JSON or keyless Connect with Google OAuth) via per-user, per-project MCP servers.
🧭 Live cloud data
Overview, Resources and Cost pull real inventory, telemetry and spend once a cloud is connected.
🛠️ Shared-compute repairs
Connect a laptop or CI runner; Nimbus drives your local Claude Code turn-by-turn to fix a repo, then pushes a branch and opens a PR. Machines are pooled per project.
🔌 GitHub & more
Connect repositories through Composio; the agent reads your code, PRs and CI before acting.
👥 Teams & permissions
Invite by email; per-member toggles for Channels, Machines & repairs, and Cloud & resources.
🔐 Secure by design
Credentials encrypted at rest (AES-256-GCM), never printed or logged; membership-enforced access; server-side audit trail.
Design infrastructure on a live canvas
Rent a machine on demand · Repair code on real machines
Nimbus is a React SPA talking to a thin Express API that fans out to services,
repositories, agents and per-user cloud MCP servers. A separate CLI worker connects
outside machines for repairs.
┌──────────────────────┐ ┌───────────────────────────────────────────────┐
│ Web app (Vite/React)│ /api │ API server (Express) │
│ localhost:5280 │ ─────► │ localhost:8788 │
│ • chat + canvas │ proxy │ │
│ • dashboard │ │ routes/ → HTTP handlers (per resource) │
│ • connections │ │ services/ → cloud, telemetry, cost, spec… │
└──────────────────────┘ │ repositories/ → SQLite persistence │
│ agents/ → ReAct chat loop + analysis │
┌──────────────────────┐ │ tools/ → agent tools (canvas, repair…) │
│ Connected machine │ poll │ libs/ → model, MCP, composio, aws/gcp │
│ @nimbus/cli + Claude │ ◄────► │ mcp/ → aws / gcloud / cloud-run MCP │
│ Code (your laptop) │ (PR) └───────────────────────────────────────────────┘
└──────────────────────┘ │
SQLite (better-sqlite3) · creds encrypted at rest
User identity comes strictly from the session cookie ; req.user.id is the isolation
boundary every service and MCP call is scoped to.
Frontend — React 18, Vite 5, React Router, React Flow, @ai-sdk/react , react-markdown
Backend — Node ≥ 20, Express, better-sqlite3 , Helmet, Zod
AI — Vercel AI SDK ( ai ), Databricks (primary) with OpenRouter cross-provider fallback,
Model Context Protocol ( @modelcontextprotocol/sdk )
Integrations — Composio (GitHub), AWS & GCP MCP servers, Fly.io (provisioning / rented machines)
The fastest way to run Nimbus. The image bundles everything the cloud MCP servers need —
git , uv (AWS), the gcloud CLI (GCP) and Node — so there's nothing else to install. You only
provide your .env .
git clone https://github.com/hritvikgupta/nimbus.git
cd nimbus
cp .env.example .env
# 1. Generate the master encryption key and append it to .env (protects stored credentials):
echo " NIMBUS_ENC_KEY= $( openssl rand -base64 32 ) " >> .env
# 2. Open .env and fill in your AI model — LLM_PROVIDER, LLM_MODEL, OPENROUTER_API_KEY.
# (Optional: FLY_API_TOKEN to rent machines, COMPOSIO_API_KEY for GitHub.)
# 3. Run it:
docker compose up --build
# open http://localhost:8788
No openssl ? Use node -e "console.log(require('crypto').randomBytes(32).toString('base64'))" and paste the result as NIMBUS_ENC_KEY= in .env .
One container serves the app and the API on port 8788 . Your data (the SQLite DB + encrypted
connections) persists in the nimbus-data volume across restarts.
Prefer plain docker ? docker build -t nimbus . && docker run -p 8788:8788 --env-file .env -v nimbus-data:/app/server/.data nimbus
To run from source instead (for development), follow the steps below.
An LLM provider — a Databricks serving endpoint or an OpenRouter API key
(optional) a Composio API key for GitHub, and AWS/GCP credentials to connect real clouds
For the cloud MCP servers: uv (AWS) and the
gcloud CLI (GCP) — bundled automatically in Docker
git clone https://github.com/hritvikgupta/nimbus.git
cd nimbus
npm install
2. Configure
cp .env.example .env
Fill in .env (see Configuration ). At minimum, set a NIMBUS_ENC_KEY and one
LLM provider:
# generate the master encryption key
node -e " console.log(require('crypto').randomBytes(32).toString('base64')) "
3. Set up the cloud MCP servers
The agent talks to your clouds through three MCP servers vendored under
server/mcp/ — AWS ( aws-mcp , Python via uvx ), Cloud Run
( cloud-run-mcp ) and gcloud ( gcloud-mcp ). Their source lives in the repo; install their
dependencies (and build the gcloud bundle) once:
npm run setup:mcp
This runs npm install for the Node servers, builds gcloud-mcp , and checks for
uv (needed to launch the AWS server). You can skip this if you
only use one cloud — the corresponding server is only spawned when a matching connection exists.
Host prerequisites (the MCP servers shell out to these — install the ones for the clouds you use):
Once installed, the agent has the complete AWS and gcloud command surface plus Cloud Run
deploy/logs — the servers themselves are the full upstream implementations, not trimmed subsets.
You still connect your own cloud credentials in the app ( Connections ).
Nimbus is two processes — the API and the web app. Run them in two terminals:
npm run api # backend → http://localhost:8788
npm run dev # frontend → http://localhost:5280 (proxies /api to :8788)
Open http://localhost:5280 , create an account, and the guided walkthrough will introduce
the workspace. Connect a repo and a cloud from Connections , then start chatting.
All configuration is via environment variables in .env (loaded by npm run api ).
Always: NIMBUS_ENC_KEY + an AI model ( LLM_PROVIDER + LLM_MODEL + OPENROUTER_API_KEY , or the Databricks trio). Your provider, your key, your billing.
To rent machines (Fly): your own FLY_API_TOKEN (+ optional FLY_* ). Leave blank to disable the feature. Rented machines need no public URL — the worker runs inside this server and drives Fly via its API. The person renting supplies their own coding-agent key per rental in the UI.
To connect your own laptop/CI for repairs: nothing here — that machine runs the @nimbus/cli worker and polls this server, so it is pointed at Nimbus ( nimbus start <key> --url https://your-host ), not the other way around.
When hosting publicly: set APP_URL to your real URL (used for GCP OAuth redirects + links), and optionally WEBHOOK_BASE , CORS_ORIGINS , ALLOW_SIGNUP=false .
Never commit your real .env . It is gitignored; .env.example is the template.
Repairs run on real machines you own via the @nimbus/cli worker. It polls
Nimbus (outbound only — works behind any NAT/firewall); when a fix is dispatched it clones the
repo, drives your local Claude Code to find and fix the issue, and opens a PR.
npm install -g @nimbus/cli
# In the app: Repairs → Connect a machine → Generate key
nimbus start < worker-key >
The machine needs Claude Code (installed + logged in), git , and gh . See
cli/README.md for the full worker docs.
nimbus/
├── src/ # React SPA — pages, components, context, styles
├── server/ # Express API
│ ├── server.mjs # thin HTTP entry — mounts routers
│ ├── routes/ # HTTP handlers (one file per resource)
│ ├── services/ # business logic (cloud, telemetry, cost, spec, repair…)
│ ├── repositories/ # persistence (SQLite via better-sqlite3)
│ ├── agents/ # ReAct chat loop + codebase-analysis agent
│ ├── tools/ # agent tools (canvas, repair, telemetry, code…)
│ ├── libs/ # external clients (model, MCP, composio, aws, gcp)
│ ├── mcp/ # AWS / gcloud / Cloud Run MCP servers
│ └── middlewares/ # session auth + cloud scoping
├── cli/ # @nimbus/cli — the connected-machine repair worker
├── docs-site/ # Fumadocs documentation site (statically exported)
├── deploy/ # Fly.io deploy configs (landing + docs)
├── docs/ # architecture & design docs (production-readiness, rented-compute…)
└── public/ # static assets (logos, brand)
Deployment
Nimbus targets Fly.io . The API server also serves the statically-exported
docs site (from docs-site/out ) whenever the request host matches docs.* , so docs.yourdomain
needs no separate p

[truncated]
