---
source: "https://github.com/matt454/agent-fleet-console"
hn_url: "https://news.ycombinator.com/item?id=48724767"
title: "Show HN: Fleet – a local-first console for managing Dockerized Hermes AI Agents"
article_title: "GitHub - matt454/agent-fleet-console · GitHub"
author: "matt454"
captured_at: "2026-06-29T20:35:30Z"
capture_tool: "hn-digest"
hn_id: 48724767
score: 1
comments: 0
posted_at: "2026-06-29T20:31:10Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Fleet – a local-first console for managing Dockerized Hermes AI Agents

- HN: [48724767](https://news.ycombinator.com/item?id=48724767)
- Source: [github.com](https://github.com/matt454/agent-fleet-console)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T20:31:10Z

## Translation

タイトル: Show HN: Fleet – Docker 化された Hermes AI エージェントを管理するためのローカルファースト コンソール
記事のタイトル: GitHub - matt454/agent-fleet-console · GitHub
説明: GitHub でアカウントを作成して、matt454/agent-fleet-console の開発に貢献します。

記事本文:
GitHub - matt454/agent-fleet-console · GitHub
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
マット454
/
エージェント-フリート-コンソール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード 他のアクションを開く m

enu フォルダーとファイル
7 コミット 7 コミット .github .github bin bin docker docker docs docs public public scripts scripts サーバー サーバー src src テスト テスト .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md コンポーネント.json コンポーネント.json Index.html Index.html knip.json knip.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
フリートは、1 つ以上の信頼できるマシン間で Dockerized Herme エージェントを作成、構成、監視、および操作するためのローカルファースト Web コンソールです。
これは主に、Hermes Agent と NVIDIA に焦点を当てた Herme の亜種である Nemo Hermes 用に構築されています。標準のHermesエージェントはデフォルトのパスです。 Nemo Hermes エージェントは、nemohermes ランナーが使用可能な場合、または自動インストールが有効な場合にサポートされます。
複数のエージェントを実行すると煩雑になる部分 (サービスの正常性、プロバイダーのデフォルト、共有資格情報、チャット セッション、ブラウザー サイドカー、VNC、ターミナル アクセス、ローカル Web パブリッシング、バックアップ、復元、クローン、リモート ノード、セットアップの準備状況) を単一のオペレーター ビューで表示します。
フリートは、ワークステーション、ホームラボ、VPN、または信頼できる LAN 上で個人またはチームが管理するエージェント インフラストラクチャを実行する技術オペレーター向けに設計されています。実行時の状態とシークレットは、デフォルトではローカルのままです。リポジトリは、ソース コードを .env 、 runtime/ 、 data/ 、 logs/ 、 secrets/ 、およびvendor/hermes-agent/ とは別に保持します。
反復可能なローカルベースラインからHermes Dockerエージェントを作成します。
nemohermes ランナーが利用可能な場合、または自動インストールが有効な場合、Nemo Hermes サンドボックス エージェントを作成します

d.
ローカル エージェントと信頼できるリモート フリート ノードを同じダッシュボードから調整します。
エージェントの状態、サービス数、健全性、メモリの準備状況、ゲートウェイの診断、ドリフト、および更新ステータスを表示します。
エージェント チャット、セッション履歴、ダッシュボード、VNC、ローカル Web プレビュー、およびコンテナー ターミナル サーフェスを開きます。
OpenAI Codex、Ollama、カスタム エンドポイント、OpenRouter のフリート全体のプロバイダーのデフォルトを保存します。
共有プロバイダーの資格情報を無視されたローカル ファイルに保存し、選択したエージェントに同期します。
Codex デバイスのログインと、Codex 認証状態のエージェントへの制御された同期をサポートします。
オンボーディング ペアリング フローを通じて Telegram 対応エージェントを作成します。
エージェントごとの Web ホスト サイドカーを介して、各エージェント ワークスペースから静的ファイルを公開します。
オペレーターが明示的にオプトインしない限り、シークレットを除外しながらエージェントをバックアップ、復元、およびクローン作成します。
リリースおよびセットアップの監査を実行して、実行時の状態、トークン、ログ、およびサイズ超過のソース ファイルを git から排除します。
git、フリートがデフォルトの Herme ソース チェックアウトを自動的に複製する必要がある場合
オプション: Nemo Hermes サンドボックス エージェントの PATH 上の nemohermes
オンボーディング画面では、これらの要件がチェックされます。ターミナルから、次のように同じチェックを実行します。
npm run init:baseline
クイックスタート
http://127.0.0.1:5180
npm run setup は、無視されるランタイム フォルダーを準備し、欠落している場合は .env を作成し、ラッパー スクリプトの実行可能ビットを修正し、必要に応じて npm 依存関係をインストールし、欠落している場合は構成済みの Herme ソースを複製し、ベースライン チェックを実行します。新しい env ファイルはコンソールを 0.0.0.0 にバインドし、信頼できる LAN マシンがコンソールにアクセスできるようにします。また、LAN から見えるコンソールは API 認証を使用する必要があるため、セットアップで HERMES_CONSOLE_TOKEN のプロンプトまたは生成が行われます。
npm start は、ベースライン チェックを実行し、フロントエンドを構築し、Express サーバーから運用アプリを提供します。
コンソール ホスト自体から http://127.0.0 を開きます。

1:5180 。別の信頼できる LAN マシンから http://<console-lan-ip>:5180 を開き、プロンプトが表示されたらコンソール トークンを入力します。フリートをローカルのみにするには、 .env で HERMES_CONSOLE_HOST=127.0.0.1 を設定します。
npm 実行開発
API は http://127.0.0.1:5180 で実行されます。 Vite フロントエンドは http://127.0.0.1:5200 で実行され、/api を API サーバーにプロキシします。開発モードでは、Vite はデフォルトで 0.0.0.0 をリッスンするため、API によってリダイレクトされた LAN クライアントはフロントエンドに到達できます。ローカルのみの開発の場合は、HERMES_CONSOLE_DEV_HOST=127.0.0.1 を設定します。
オンボーディング画面で設定チェックを確認します。
フリート設定を開き、プロバイダーのデフォルトを選択します。
選択したプロバイダーに認証が必要な場合は、共有資格情報を追加するか、Codex デバイスのログインを完了します。
ダッシュボードからエージェントを作成します。
チャット、ライフサイクル、ゲートウェイ、ターミナル、Cron、認証情報、診断のエージェント詳細ビューを開きます。
デフォルトでは、フリートは エルメス ソースを HERMES_AGENT_REPO_URL から Vendor/hermes-agent にクローンします。他の場所でチェックアウトを使用するには、次のように設定します。
HERMES_AGENT_SRC = /パス/to/hermes-agent
ソースの自動ダウンロードを無効にするには:
エルメス_AGENT_AUTO_CLONE = 0
ドキュメント
npm run setup # ローカルのランタイム状態と依存関係を準備する
npm start # 本番コンソールを構築して提供する
npm run dev # APIとVite開発サーバーを実行します
npm run init:baseline # ローカルセットアップの準備状況を確認します
npm run check # TypeScript の型チェック
npm test # ノードテストスイートを実行します
npm run build # 本番フロントエンドビルド
npm run Audit:release # リポジトリ健全性監査
npm run release:チェック # 完全なリリース ゲート
npm run knip # 未使用コードの監査
フリートには、リポジトリ ローカル ラッパーも含まれています。
ビン/エルメスコンソール
bin/hermes-docker ステータス < エージェント >
bin/hermes-docker ログ < エージェント >
bin/hermes-docker シェル < エージェント >
bin/hermes-docker restart <エージェント>
bin/hermes-docker update <エージェント>
bin/hermes-docker delete <エージェント>
を使用します。

通常の操作用の UI。ローカル Docker エージェント用の直接ターミナル エスケープ ハッチが必要な場合は、ラッパーを使用します。
フリートは最初にプロセス環境を読み取り、存在する場合は次にこれらのファイルを読み取ります。
<HERMES_INSTANCES_ROOT>/.env (エルメス_INSTANCES_ROOT が外部の場合)
既存のプロセス環境変数は .env 値よりも優先されます。シェル プロファイル、サービス マネージャー、または launchctl エクスポートで依然として HERMES_CONSOLE_HOST=127.0.0.1 が設定されている場合、.env が 0.0.0.0 となっている場合でも、フリートはループバックのみになります。
最も重要なローカル設定は次のとおりです。
HERMES_INSTANCES_ROOT = ./ランタイム
HERMES_DOCKER_BIN = ./bin/hermes-docker
HERMES_AGENT_SRC = ./vendor/hermes-agent
エルメス_AGENT_AUTO_CLONE = 1
エルメス_CAMOFOX_CONTEXT = ./docker/camofox
エルメス_WEBHOST_CONTEXT = ./docker/webhost
エルメス_コンソール_ホスト = 0.0.0.0
エルメス_コンソール_ポート = 5180
エルメス_コンソール_データ_DIR = ./データ
HERMES_CONSOLE_SECRETS_DIR = ./シークレット
エルメス_コンソール_REQUIRE_AUTH = 1
リモート フリート ノードの場合、[フリート設定] -> [フリート ノード] に入力されるベース URL は、リモート マシンの LAN アドレス (例: http://192.168.3.232:5180 ) に加えて、そのリモート ノードの .env に保存されている同じベアラー トークンを使用する必要があります。
フリート全体のプロバイダーのデフォルトと共有資格情報はフリート設定から管理され、無視されるファイルに保存されます。
Secrets/global-provider.json
Secrets/global-credentials.env
シークレット/グローバル認証/
シークレット/global-sync.json
エージェントごとの構成は、各エージェント フォルダーの下に存在します。
<エージェント>/
ホーム/
.env
config.yaml
SOUL.md
ワークスペース/
エルメス_WEB.md
ウェブ/
インスタンス.env
compose.yaml
完全な環境およびストレージ ガイドについては、構成リファレンスを参照してください。
フリートはデフォルトで 0.0.0.0 にバインドされるため、信頼された LAN フリート ノードはそれに到達できます。 API 認証を有効にしたまま、次のように設定します。
HERMES_CONSOLE_TOKEN = <長いランダムトークン>
エルメス_コンソール_REQUIRE_AUTH = 1
サーバーは、HERMES_CONSOLE_TOKEN でない限り、非ループバック バインドを拒否します。

が設定されています。 LAN バインディングまたは必要な認証でトークンが必要な場合、npm run setup はトークンを要求するか生成します。ローカルのみで使用する場合は、 HERMES_CONSOLE_HOST=127.0.0.1 を設定します。
フリートをコントロール プレーンとして扱います。コンテナーの開始と停止、ターミナルのオープン、資格情報の同期、バックアップの復元、エージェントの作成、リモート ノード アクションのプロキシ、およびオプションで自己更新コマンドの実行が可能です。個々の Herme ダッシュボード、Camofox VNC エンドポイント、およびエージェント Web ホストは、意図的に保護および公開しない限り、非公開にしてください。
フリート ノードを使用すると、1 つのコンソールが信頼できるマシン上の他のフリート コンソールを調整できます。フリート設定 -> フリート ノードに、ラベル、ベース URL、オプションのベアラー トークン、および有効な状態を含むリモート コンソールを追加します。
次に、ダッシュボードはローカル エージェントとリモート エージェントをマージし、各行のホストを表示し、選択したノードを介して作成、開始、停止、再起動、更新、削除、クローン、バックアップ、チャット、ゲートウェイ、ターミナル、および詳細アクションをルーティングします。
リモート ノード ベアラー トークンは API 応答では編集されますが、ローカルの data/fleet.db に保存されます。データ/プライベートを保持し、共有マシン上でディスク暗号化を使用します。
データ/バックアップ/
バックアップには、エージェント構成、選択したワークスペース ファイル、プロバイダーのデフォルト、およびマニフェストが含まれます。 home/.env 、グローバル資格情報、OAuth 状態、トークンのようなファイル、生成された実行時シークレットなどのシークレットは、オペレーターが明示的にシークレットのエクスポートを有効にしない限り除外されます。
復元では、コンソール ホスト上のローカル .tar.gz アーカイブ パスを使用します。復元されたエージェントは、新しく生成されたポートとランタイム シークレットを受け取ります。クローン作成では、ローカル エージェントを新しい名前で複製し、オプションでワークスペース ファイルとローカルのエージェントごとの資格情報を含めることができます。
アプリで使用される bin/ フリート ラッパー スクリプト
docker/camofox/ Camofox サイドカー画像コンテキスト
docker/webhost/ Node.js 静的 Webhost サイドカー イメージ コンテキスト
サーバー/ Express API、サービス、SQLite、termina

l ウェブソケット
src/ React フロントエンド、UI 状態、スタイル、共有モデル
スクリプト/セットアップ、ベースライン、リリース監査、開発/オーケストレーションの開始
docs/ ユーザー、オペレーター、API、およびメンテナーのドキュメント
ランタイム/デフォルトのHermesインスタンスルートを無視
データ/無視された SQLite データベースとローカル制御状態
ログ/無視されたローカル プロセス ログ
Secrets/ グローバルプロバイダーの資格情報と OAuth 状態が無視されました
Vendor/hermes-agent/ はオプションの Hermes ソース チェックアウト/パッケージを無視しました
解放する
プル リクエストを公開または開く前に:
npm run release:チェック
git status --short
セットアップ、オンボーディング、Docker、または環境の変更の場合は、次のコマンドも実行します。
npm run init:baseline -- --json
リリース ゲートは、タイプ チェック、テスト、本番環境のビルド、リポジトリの健全性監査、本番環境の依存関係の監査、および未使用コードの監査を実行します。手動レビューリストについては、「リリースチェックリスト」を参照してください。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to matt454/agent-fleet-console development by creating an account on GitHub.

GitHub - matt454/agent-fleet-console · GitHub
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
matt454
/
agent-fleet-console
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github .github bin bin docker docker docs docs public public scripts scripts server server src src tests tests .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md components.json components.json index.html index.html knip.json knip.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json vite.config.ts vite.config.ts View all files Repository files navigation
Fleet is a local-first web console for creating, configuring, monitoring, and operating Dockerized Hermes agents across one or more trusted machines.
It is primarily built for Hermes Agent and the NVIDIA-focused Hermes variant, Nemo Hermes. Standard Hermes agents are the default path; Nemo Hermes agents are supported when the nemohermes runner is available or automatic installation is enabled.
It gives a single operator view for the parts that become noisy once you run more than one agent: service health, provider defaults, shared credentials, chat sessions, browser sidecars, VNC, terminal access, local web publishing, backups, restores, clones, remote nodes, and setup readiness.
Fleet is designed for technical operators running personal or team-controlled agent infrastructure on a workstation, homelab, VPN, or trusted LAN. Runtime state and secrets stay local by default; the repository keeps source code separate from .env , runtime/ , data/ , logs/ , secrets/ , and vendor/hermes-agent/ .
Creates Hermes Docker agents from a repeatable local baseline.
Creates Nemo Hermes sandbox agents when the nemohermes runner is available or auto-install is enabled.
Coordinates local agents and trusted remote Fleet nodes from the same dashboard.
Shows agent state, service counts, health, memory readiness, gateway diagnostics, drift, and update status.
Opens agent chat, session history, dashboard, VNC, local web preview, and container terminal surfaces.
Saves fleet-wide provider defaults for OpenAI Codex, Ollama, Custom endpoints, and OpenRouter.
Stores shared provider credentials in ignored local files and syncs them into selected agents.
Supports Codex device login and controlled sync of Codex auth state into agents.
Creates Telegram-enabled agents through the onboarding pairing flow.
Publishes static files from each agent workspace through a per-agent webhost sidecar.
Backs up, restores, and clones agents while excluding secrets unless an operator explicitly opts in.
Runs release and setup audits that keep runtime state, tokens, logs, and oversized source files out of git.
git, when Fleet should clone the default Hermes source checkout automatically
Optional: nemohermes on PATH for Nemo Hermes sandbox agents
The onboarding screen checks these requirements. From a terminal, run the same check with:
npm run init:baseline
Quickstart
http://127.0.0.1:5180
npm run setup prepares ignored runtime folders, creates .env when missing, fixes executable bits on wrapper scripts, installs npm dependencies when needed, clones the configured Hermes source when it is missing, and runs the baseline check. New env files bind the console to 0.0.0.0 so trusted LAN machines can reach it, and setup prompts for or generates HERMES_CONSOLE_TOKEN because LAN-visible consoles must use API auth.
npm start runs the baseline check, builds the frontend, and serves the production app from the Express server.
From the console host itself, open http://127.0.0.1:5180 . From another trusted LAN machine, open http://<console-lan-ip>:5180 and enter the console token when prompted. To keep Fleet local-only, set HERMES_CONSOLE_HOST=127.0.0.1 in .env .
npm run dev
The API runs on http://127.0.0.1:5180 ; the Vite frontend runs on http://127.0.0.1:5200 and proxies /api to the API server. In dev mode, Vite listens on 0.0.0.0 by default so LAN clients redirected by the API can reach the frontend; set HERMES_CONSOLE_DEV_HOST=127.0.0.1 for local-only development.
Review setup checks on the onboarding screen.
Open Fleet settings and choose a provider default.
Add shared credentials or complete Codex device login if your selected provider needs auth.
Create an agent from the dashboard.
Open the agent detail view for chat, lifecycle, gateway, terminal, crons, credentials, and diagnostics.
By default, Fleet clones Hermes source into vendor/hermes-agent from HERMES_AGENT_REPO_URL . To use a checkout somewhere else, set:
HERMES_AGENT_SRC = /path/to/hermes-agent
To disable automatic source download:
HERMES_AGENT_AUTO_CLONE = 0
Documentation
npm run setup # prepare local runtime state and dependencies
npm start # build and serve the production console
npm run dev # run API and Vite dev server
npm run init:baseline # check local setup readiness
npm run check # TypeScript type check
npm test # run node test suite
npm run build # production frontend build
npm run audit:release # repository hygiene audit
npm run release:check # full release gate
npm run knip # unused-code audit
Fleet also includes repo-local wrappers:
bin/hermes-console
bin/hermes-docker status < agent >
bin/hermes-docker logs < agent >
bin/hermes-docker shell < agent >
bin/hermes-docker restart < agent >
bin/hermes-docker update < agent >
bin/hermes-docker delete < agent >
Use the UI for normal operation. Use the wrappers when you need a direct terminal escape hatch for local Docker agents.
Fleet reads process env first, then these files when present:
<HERMES_INSTANCES_ROOT>/.env when HERMES_INSTANCES_ROOT is external
Existing process environment variables win over .env values. If a shell profile, service manager, or launchctl export still sets HERMES_CONSOLE_HOST=127.0.0.1 , Fleet will stay loopback-only even when .env says 0.0.0.0 .
The most important local settings are:
HERMES_INSTANCES_ROOT = ./runtime
HERMES_DOCKER_BIN = ./bin/hermes-docker
HERMES_AGENT_SRC = ./vendor/hermes-agent
HERMES_AGENT_AUTO_CLONE = 1
HERMES_CAMOFOX_CONTEXT = ./docker/camofox
HERMES_WEBHOST_CONTEXT = ./docker/webhost
HERMES_CONSOLE_HOST = 0.0.0.0
HERMES_CONSOLE_PORT = 5180
HERMES_CONSOLE_DATA_DIR = ./data
HERMES_CONSOLE_SECRETS_DIR = ./secrets
HERMES_CONSOLE_REQUIRE_AUTH = 1
For a remote Fleet node, the base URL entered in Fleet settings -> Fleet nodes should use the remote machine's LAN address, for example http://192.168.3.232:5180 , plus the same bearer token stored in that remote node's .env .
Fleet-wide provider defaults and shared credentials are managed from Fleet settings and stored in ignored files:
secrets/global-provider.json
secrets/global-credentials.env
secrets/global-oauth/
secrets/global-sync.json
Per-agent config lives under each agent folder:
<agent>/
home/
.env
config.yaml
SOUL.md
workspace/
HERMES_WEB.md
web/
instance.env
compose.yaml
See Configuration reference for the full environment and storage guide.
Fleet binds to 0.0.0.0 by default so trusted LAN Fleet nodes can reach it. Keep API auth enabled and set:
HERMES_CONSOLE_TOKEN = <long-random-token>
HERMES_CONSOLE_REQUIRE_AUTH = 1
The server refuses non-loopback binds unless HERMES_CONSOLE_TOKEN is set. npm run setup prompts for or generates a token when LAN binding or required auth needs one. For local-only use, set HERMES_CONSOLE_HOST=127.0.0.1 .
Treat Fleet as a control plane. It can start and stop containers, open terminals, sync credentials, restore backups, create agents, proxy remote node actions, and optionally run self-update commands. Keep individual Hermes dashboards, Camofox VNC endpoints, and agent webhosts private unless you intentionally protect and expose them.
Fleet Nodes let one console coordinate other Fleet consoles on trusted machines. Add remote consoles in Fleet settings -> Fleet nodes with a label, base URL, optional bearer token, and enabled state.
The dashboard then merges local and remote agents, shows the host for each row, and routes create, start, stop, restart, update, delete, clone, backup, chat, gateway, terminal, and detail actions through the selected node.
Remote node bearer tokens are redacted in API responses but stored locally in data/fleet.db ; keep data/ private and use disk encryption on shared machines.
data/backups/
Backups include agent config, selected workspace files, provider defaults, and a manifest. Secrets such as home/.env , global credentials, OAuth state, token-like files, and generated runtime secrets are excluded unless an operator explicitly enables secret export.
Restore uses a local .tar.gz archive path on the console host. Restored agents receive fresh generated ports and runtime secrets. Clone duplicates a local agent into a new name and can optionally include workspace files and local per-agent credentials.
bin/ fleet wrapper scripts used by the app
docker/camofox/ Camofox sidecar image context
docker/webhost/ Node.js static webhost sidecar image context
server/ Express API, services, SQLite, terminal websocket
src/ React frontend, UI state, styles, shared models
scripts/ setup, baseline, release audit, dev/start orchestration
docs/ user, operator, API, and maintainer documentation
runtime/ ignored default Hermes instance root
data/ ignored SQLite database and local control state
logs/ ignored local process logs
secrets/ ignored global provider credentials and OAuth state
vendor/hermes-agent/ ignored optional Hermes source checkout/package
Releasing
Before publishing or opening a pull request:
npm run release:check
git status --short
For setup, onboarding, Docker, or environment changes, also run:
npm run init:baseline -- --json
The release gate runs type checking, tests, production build, repository hygiene audit, production dependency audit, and unused-code audit. See Release checklist for the manual review list.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
