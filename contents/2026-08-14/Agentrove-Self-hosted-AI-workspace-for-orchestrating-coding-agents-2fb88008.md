---
source: "https://github.com/Mng-dev-ai/agentrove"
hn_url: "https://news.ycombinator.com/item?id=49293480"
title: "Agentrove: Self-hosted AI workspace for orchestrating coding agents"
article_title: "GitHub - Mng-dev-ai/agentrove: Self-hosted AI coding workspace to run and orchestrate Claude Code, Codex, Copilot, Cursor, Grok and OpenCode agents — multi-agent workflows, personas, and ACP-powered sandboxes · GitHub"
author: "indigodaddy"
captured_at: "2026-08-14T01:05:01Z"
capture_tool: "hn-digest"
hn_id: 49293480
score: 1
comments: 0
posted_at: "2026-08-14T00:38:40Z"
tags:
  - hacker-news
  - translated
---

# Agentrove: Self-hosted AI workspace for orchestrating coding agents

- HN: [49293480](https://news.ycombinator.com/item?id=49293480)
- Source: [github.com](https://github.com/Mng-dev-ai/agentrove)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T00:38:40Z

## Translation

タイトル: Agentrove: コーディング エージェントを調整するための自己ホスト型 AI ワークスペース
記事のタイトル: GitHub - Mng-dev-ai/agentrove: Claude Code、Codex、Copilot、Cursor、Grok、および OpenCode エージェントを実行および調整するための自己ホスト型 AI コーディング ワークスペース — マルチエージェント ワークフロー、ペルソナ、および ACP を利用したサンドボックス · GitHub
説明: Claude Code、Codex、Copilot、Cursor、Grok、および OpenCode エージェントを実行および調整するための自己ホスト型 AI コーディング ワークスペース — マルチエージェント ワークフロー、ペルソナ、ACP を利用したサンドボックス - Mng-dev-ai/agentrove

記事本文:
GitHub - Mng-dev-ai/agentrove: Claude Code、Codex、Copilot、Cursor、Grok、および OpenCode エージェントを実行および調整するための自己ホスト型 AI コーディング ワークスペース — マルチエージェント ワークフロー、ペルソナ、および ACP を利用したサンドボックス · GitHub
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
MNG-DEV-AI
/
エージェントローブ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,964 コミット 1,964 コミット .githooks .githooks .github/ workflows .github/ workfl

ows バックエンド バックエンド デスクトップ デスクトップ フロントエンド フロントエンド mcp-server mcp-server サンドボックス/ docker サンドボックス/ docker スクリーンショット スクリーンショット スクリプト スクリプト .dockerignore .dockerignore .env.desktop.example .env.desktop.example .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml ライセンス ライセンス README.md README.md docker-compose-production.yml docker-compose-production.yml docker-compose.yml docker-compose.yml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1 つのインターフェイスから Claude Code、Codex、Copilot、Cursor、Grok、および OpenCode エージェントを実行および調整するための自己ホスト型 AI コーディング ワークスペース。
ACP アダプターを介して Claude、Codex、Copilot、Cursor、Grok、および OpenCode を実行します。
各ワークスペースに独自の Docker またはホスト サンドボックスを与えます。
チャット、コード エディター、ターミナル、ファイル ツリー、差分、シークレット、Git ツールを 1 つのワークスペースに結合します。
空のフォルダー、Git クローン、既存のローカル フォルダー、または GitHub リポジトリからのワークスペースをサポートします。
キャンセル、許可プロンプト、キューに入れられたフォローアップ メッセージ、ファイルのメンション、スラッシュ コマンド、添付ファイルを含むエージェント セッションをストリーミングします。
サブスレッド、固定チャット、ワークツリー モード、ペルソナ、カスタム指示、環境変数、インストールされたエージェント スキルが含まれます。
バンドルされた MCP サーバーを介してマルチエージェントのワークフローを調整します。リード チャットは、インストールされているエージェント、モデル、ペルソナに関するワーカー サブスレッドを必要に応じて並列ワークツリーで生成し、結果をレビューします。
GitHub を利用したリポジトリの参照、プル リクエストのレビュー、PR の作成、レビュー担当者の選択、および git ブランチ/コミット/プッシュ/プル ヘルパーを提供します。
Docker Web アプリ、macOS デスクトップ アプリ、およびネイティブ iOS アプリとして出荷されます。
Agentrove チャットは単なるエンドポイントではなく、相互に推進し合うことができます

。バンドルされた MCP サーバー ( mcp-server/ ) は、インスタンス全体をツール ( send_message 、 get_messages 、 list_models 、 list_personas など) として公開し、すべてのチャットのエージェントがこれらのツールを利用できます。これにより、あらゆるチャットがオーケストレーターに変わります。つまり、強力なモデルに基づくリード エージェントが作業を分解し、各タスクを適切なエージェントにルーティングし、返された内容をレビューすることになります。
サブスレッド — send_message(parent_chat_id=…) は、同じワークスペースおよびブランチ内で、リード チャットの下にグループ化されたワーカー チャットを作成します。サブスレッドはフラット (ネストなし) のままです。リードはファンアウトし、ワーカーは報告を返します。
ターンごとのモデルとペルソナ — 各ワーカーは、インストールされている任意のエージェントとモデル ( model_id ) 上で任意のペルソナ (カスタム システム プロンプト) で実行されます。典型的なフリート: コードベース探索用の読み取り専用スカウト ペルソナ、実装用のコーディング モデル、QA 用の専用レビュー担当者ペルソナ (バグ検索、構造品質) を備えた高速モデル。 list_models は、各モデルでサポートされている推論層 ( Thinking_modes ) をレポートするため、リーダーはタスクごとに労力をダイヤルします。
分離されたワークツリー — worktree=true はワーカーに独自の git ワークツリーを与えるため、並列ワーカーは競合することなく同時に編集できます。
ポーリングとフォローアップ — リードはワーカーのターンが完了するまで get_messages をポーリングし、結果を実際のコードと比較して判断し、ワーカー自身のスレッドにリワークを送信します。フォローアップは、スレッドの以前のモデル、ペルソナ、推論設定を継承します。
無人ターン - オーケストレーションされたターンはエージェントの完全実行モードで実行されるため、ワーカーは許可のプロンプトなしで終了します。
典型的なループ: リードは安価で高速なモデルを探索し、正確な仕様をサブスレッドのコーディング モデルに渡し、レビュアー ペルソナを並行して diff 上でファンにし、結果を優先順位付けし、受け入れられた修正を委任し、最終結果を所有します。モデル、ペルソナ、およびr

外出ルールはすべてユーザー定義であるため、どのフリートを実行しても同じ機械が駆動します。
git clone https://github.com/Mng-dev-ai/agentrove.git
CD エージェントローブ
cp .env.example .env
.env に SECRET_KEY を設定します。
openssl rand -hex 32
Agentrove を起動します。
ドッカー構成 -d
http://localhost:3000 を開きます。
Agentrove には、Tauri で構築された macOS デスクトップ アプリもあります。バンドルされた Python バックエンド サイドカーを利用可能な 127.0.0.1 ポートで起動し、起動時にフロントエンドをそれに接続します。
最新の Apple Silicon ビルドを Releases からダウンロードします。
CDフロントエンド
npmインストール
npm run デスクトップ:dev
モバイル (iOS)
Agentrove は、Tauri を使用してネイティブ iOS アプリも構築します。 iOSはローカルで実行できないので、
バックエンド サイドカー、アプリはシン クライアントです。Agentrove インスタンスと通信します。
すでにホスト (Docker または運用環境) にあり、電話からアクセス可能
https/wss 。プロジェクトはオープンソースであるため、自分でビルドして署名します。
他の人のサーバーには何もハードコーディングされません。
要件: Xcode (およびその iOS SDK およびシミュレータ) を備えた macOS、Rust iOS
ターゲット (rustup target add aarch64-apple-ios aarch64-apple-ios-sim )、および CocoaPods。
デバイスで SDK がまだ安定した Xcode に含まれていない iOS ベータ版を実行している場合は、最新の安定した Xcode の SDK ( sudo xcode-select -s /Applications/Xcode.app ) に対してビルドします。古い SDK でビルドされたアプリは、新しい iOS でも引き続き実行されます。
アプリをインスタンスに向けます。
CDフロントエンド
cp .env.mobile.example .env.mobile # 次に https/wss URL を設定します
npmインストール
シミュレーターで実行します。
npm 実行 ios:dev
無料の Apple ID を使用して自分の iPhone にインストールします。有料の開発者アカウントは必要ありません
(アプリは 7 日ごとに再署名する必要があります):
npm run tauri ios init は Xcode プロジェクトを生成します (初回実行のみ)。
Xcodeでfrontend/src-tauri/gen/apple/*.xcodeprojを一度開き、チームを選択してください
「署名と機能」の下で、

iPhone を接続します (Xcode は
デバイスを登録し、プロビジョニング プロファイルを作成します)。
電話機で開発者モードを有効にし ( [設定] → [プライバシーとセキュリティ] )、
最初のインストール後、以下の証明書を信頼します
設定 → 一般 → VPN とデバイス管理 。
スタンドアロンの .ipa をビルドし、接続されている iPhone にインストールします。最も簡単な道
はバンドルされたヘルパーであり、ビルド、署名、エクスポート、インストールを 1 ステップで実行します。それ
APPLE_DEVELOPMENT_TEAM からチームを読み取ります (つまり、誰にも何もハードコーディングされません)
他のチーム）、接続されているデバイスを自動検出します。
import APPLE_DEVELOPMENT_TEAM= < YOUR_TEAM_ID > # Xcode → 署名と機能で見つけます
CDフロントエンド
npm run ios:インストール
または、自分で手順を実行します。 -c フラグは、チームを両方のビルドに挿入します。
署名と IPA エクスポートなので、コミットされた構成には含まれません。
npm run ios:build -- --export-method デバッグ \
-c ' {"バンドル":{"iOS":{"開発チーム":"<YOUR_TEAM_ID>"}}} '
# -> src-tauri/gen/apple/build/arm64/Agentrove.ipa
xcrun devicectl list devices # デバイス ID を検索します
xcrun devicectl device install app --device < DEVICE_ID > \
src-tauri/gen/apple/build/arm64/Agentrove.ipa
後でアプリを更新するには、npm run ios:install を再実行します。スタンドアロン ビルドは上で実行されます。
Mac を接続したままにせずに電話を使用できます。
単一ホストの Docker デプロイメントの場合:
SECRET_KEY= $( openssl rand -hex 32 ) \
SERVICE_URL_WEB_80=https://yourdomain.com \
APP_URL=https://yourdomain.com \
ALLOWED_ORIGINS=https://yourdomain.com \
docker compose -f docker-compose-production.yml up -d --build
Coolify / リバースプロキシ
docker-compose-production.yml を使用して Web サービス (ポート 80) で Coolify をポイントします。設定:
APP_URL / ALLOWED_ORIGINS が設定されていない場合、作成ファイルは Coolify の SERVICE_URL_WEB_80 にフォールバックします (実稼働および PR プレビューでは正しい)。

/admin がスタイルのないプレーンな HTML (青色のリンク、レイアウトなし) として読み込まれる場合、ページは https:// ですが、API は SQLAdmin CSS/JS の http:// URL を生成します。上記のプロダクション コンポーズを使用して再構築し、nginx が Coolify の X-Forwarded-Proto を転送し、API がプロキシを信頼するようにします。
Coolify プレビュー展開 (PR ごと)
[詳細設定] → [プレビュー デプロイメント] を有効にします (リポジトリがパブリックである / フォーク PR を受け入れる場合は、パブリック PR デプロイメントを許可します)。
デプロイメントのプレビュー → URL テンプレートを設定します。 pr-{{pr_id}}.yourdomain.com 。
DNS ワイルドカード ( *.yourdomain.com または pr-*.yourdomain.com ) を Coolify サーバーに指定し、Coolify に証明書を発行させます。
「環境変数」→「デプロイメントのプレビュー」で、次のように設定します。
AGENTROVE_STORAGE_SOURCE=agentrove_storage (名前付きボリューム。本番環境のホスト パスを共有しません)
APP_URL / ALLOWED_ORIGINS 空または $SERVICE_URL_WEB_80 (したがって、各プレビューは独自のオリジンを取得します)
SECRET_KEY とその他のシークレットを本番環境からコピーします
PR を開きます (またはプル リクエストをロードして手動でデプロイします)。 Coolify は、プレビュー URL で分離された作成スタックをスピンアップします。
サンドボックス Docker ネットワークは構成プロジェクトごとにあるため (固定名はありません)、プレビュー サンドボックスは運用環境のネットワークに参加しません。
フロントエンド: React 19、TypeScript、Vite、Tailwind CSS、Monaco、xterm.js
バックエンド: FastAPI、SQLAlchemy、SQLite、Redis
ランタイム: ACP、Docker またはホスト サンドボックス、Tauri デスクトップ サイドカー
貢献やフィードバックは大歓迎です。
Claude Code、Codex、Copilot、Cursor、Grok、OpenCode エージェントを実行および調整するための自己ホスト型 AI コーディング ワークスペース — マルチエージェント ワークフロー、ペルソナ、ACP を利用したサンドボックス
Readme Apache-2.0 ライセンス アクティビティ スター
61 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-hosted AI coding workspace to run and orchestrate Claude Code, Codex, Copilot, Cursor, Grok and OpenCode agents — multi-agent workflows, personas, and ACP-powered sandboxes - Mng-dev-ai/agentrove

GitHub - Mng-dev-ai/agentrove: Self-hosted AI coding workspace to run and orchestrate Claude Code, Codex, Copilot, Cursor, Grok and OpenCode agents — multi-agent workflows, personas, and ACP-powered sandboxes · GitHub
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
Mng-dev-ai
/
agentrove
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1,964 Commits 1,964 Commits .githooks .githooks .github/ workflows .github/ workflows backend backend desktop desktop frontend frontend mcp-server mcp-server sandbox/ docker sandbox/ docker screenshots screenshots scripts scripts .dockerignore .dockerignore .env.desktop.example .env.desktop.example .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml LICENSE LICENSE README.md README.md docker-compose-production.yml docker-compose-production.yml docker-compose.yml docker-compose.yml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Self-hosted AI coding workspace for running and orchestrating Claude Code, Codex, Copilot, Cursor, Grok, and OpenCode agents from one interface.
Runs Claude, Codex, Copilot, Cursor, Grok, and OpenCode through ACP adapters.
Gives each workspace its own Docker or host sandbox.
Combines chat, code editor, terminal, file tree, diffs, secrets, and git tools in one workspace.
Supports workspaces from empty folders, git clones, existing local folders, or GitHub repositories.
Streams agent sessions with cancellation, permission prompts, queued follow-up messages, file mentions, slash commands, and attachments.
Includes sub-threads, pinned chats, worktree mode, personas, custom instructions, environment variables, and installed agent skills.
Orchestrates multi-agent workflows through the bundled MCP server: a lead chat spawns worker sub-threads on any installed agent, model, and persona — in parallel worktrees when needed — then reviews their results.
Provides GitHub-assisted repository browsing, pull request review, PR creation, reviewer selection, and git branch/commit/push/pull helpers.
Ships as a Docker web app, a macOS desktop app, and a native iOS app.
Agentrove chats aren't just endpoints — they can drive each other. The bundled MCP server ( mcp-server/ ) exposes the whole instance as tools ( send_message , get_messages , list_models , list_personas , …), and every chat's agent has those tools available. That turns any chat into an orchestrator: a lead agent on a strong model that decomposes the work, routes each task to the right agent, and reviews what comes back.
Sub-threads — send_message(parent_chat_id=…) creates a worker chat grouped under the lead chat, in the same workspace and branch. Sub-threads stay flat (no nesting): the lead fans out, workers report back.
Per-turn model and persona — each worker runs on any installed agent and model ( model_id ) with any persona (a custom system prompt). Typical fleet: a fast model with a read-only scout persona for codebase exploration, coding models for implementation, dedicated reviewer personas (bug hunting, structural quality) for QA. list_models reports each model's supported reasoning tiers ( thinking_modes ), so the lead dials effort per task.
Isolated worktrees — worktree=true gives a worker its own git worktree, so parallel workers edit concurrently without conflicts.
Polling and follow-ups — the lead polls get_messages until a worker's turn completes, judges the result against the actual code, and sends rework to the worker's own thread; follow-ups inherit the thread's previous model, persona, and reasoning settings.
Unattended turns — orchestrated turns run in the agent's full-execution mode, so workers finish without permission prompts.
A typical loop: the lead explores with a cheap fast model, hands a precise spec to a coding model in a sub-thread, fans reviewer personas out over the diff in parallel, triages their findings, delegates the accepted fixes — and owns the final result. Models, personas, and routing rules are all user-defined, so the same machinery drives whatever fleet you run.
git clone https://github.com/Mng-dev-ai/agentrove.git
cd agentrove
cp .env.example .env
Set SECRET_KEY in .env :
openssl rand -hex 32
Start Agentrove:
docker compose up -d
Open http://localhost:3000 .
Agentrove also has a macOS desktop app built with Tauri. It starts a bundled Python backend sidecar on an available 127.0.0.1 port and connects the frontend to it at launch.
Download the latest Apple Silicon build from Releases .
cd frontend
npm install
npm run desktop:dev
Mobile (iOS)
Agentrove also builds a native iOS app with Tauri. Since iOS can't run the local
backend sidecar, the app is a thin client: it talks to an Agentrove instance you
already host (your Docker or production deployment), reachable from the phone over
https / wss . Because the project is open source, you build and sign it yourself —
nothing is hardcoded to anyone else's server.
Requirements: macOS with Xcode (plus its iOS SDK and Simulator), the Rust iOS
targets ( rustup target add aarch64-apple-ios aarch64-apple-ios-sim ), and CocoaPods.
If your device runs an iOS beta whose SDK isn't in a stable Xcode yet, build against the latest stable Xcode's SDK ( sudo xcode-select -s /Applications/Xcode.app ) — apps built with an older SDK still run on newer iOS.
Point the app at your instance:
cd frontend
cp .env.mobile.example .env.mobile # then set your https/wss URLs
npm install
Run in the simulator:
npm run ios:dev
Install on your own iPhone with a free Apple ID — no paid developer account needed
(the app must be re-signed every 7 days):
npm run tauri ios init generates the Xcode project (first run only).
Open frontend/src-tauri/gen/apple/*.xcodeproj in Xcode once, pick your Team
under Signing & Capabilities , and connect your iPhone (Xcode needs to
register the device and create the provisioning profile).
On the phone, enable Developer Mode ( Settings → Privacy & Security ), then
after the first install trust the certificate under
Settings → General → VPN & Device Management .
Build a standalone .ipa and install it on the connected iPhone. The easiest path
is the bundled helper, which builds, signs, exports, and installs in one step. It
reads your team from APPLE_DEVELOPMENT_TEAM (so nothing is hardcoded to anyone
else's team) and auto-detects the connected device:
export APPLE_DEVELOPMENT_TEAM= < YOUR_TEAM_ID > # find it in Xcode → Signing & Capabilities
cd frontend
npm run ios:install
Or run the steps yourself. The -c flag injects the team into both the build
signing and the IPA export, so it stays out of the committed config:
npm run ios:build -- --export-method debugging \
-c ' {"bundle":{"iOS":{"developmentTeam":"<YOUR_TEAM_ID>"}}} '
# -> src-tauri/gen/apple/build/arm64/Agentrove.ipa
xcrun devicectl list devices # find your device id
xcrun devicectl device install app --device < DEVICE_ID > \
src-tauri/gen/apple/build/arm64/Agentrove.ipa
To update the app later, re-run npm run ios:install — the standalone build runs on
the phone without keeping a Mac connected.
For a single-host Docker deployment:
SECRET_KEY= $( openssl rand -hex 32 ) \
SERVICE_URL_WEB_80=https://yourdomain.com \
APP_URL=https://yourdomain.com \
ALLOWED_ORIGINS=https://yourdomain.com \
docker compose -f docker-compose-production.yml up -d --build
Coolify / reverse proxy
Point Coolify at the web service (port 80) using docker-compose-production.yml . Set:
If APP_URL / ALLOWED_ORIGINS are unset, the compose file falls back to Coolify's SERVICE_URL_WEB_80 (correct for production and PR previews).
If /admin loads as plain unstyled HTML (blue links, no layout), the API is generating http:// URLs for SQLAdmin CSS/JS while the page is https:// . Rebuild with the production compose above so nginx forwards Coolify's X-Forwarded-Proto and the API trusts the proxy.
Coolify preview deployments (per PR)
Advanced → enable Preview Deployments (and Allow Public PR Deployments if the repo is public / accepts fork PRs).
Preview Deployments → set the URL template, e.g. pr-{{pr_id}}.yourdomain.com .
Point a DNS wildcard ( *.yourdomain.com or pr-*.yourdomain.com ) at the Coolify server and let Coolify issue certificates.
In Environment Variables → Preview Deployments , set:
AGENTROVE_STORAGE_SOURCE=agentrove_storage (named volume; do not share production's host path)
APP_URL / ALLOWED_ORIGINS empty or $SERVICE_URL_WEB_80 (so each preview gets its own origin)
copy SECRET_KEY and other secrets from production
Open a PR (or Load Pull Requests and deploy manually). Coolify spins up an isolated compose stack at the preview URL.
The sandbox Docker network is per compose project (no fixed name), so preview sandboxes do not join production's network.
Frontend: React 19, TypeScript, Vite, Tailwind CSS, Monaco, xterm.js
Backend: FastAPI, SQLAlchemy, SQLite, Redis
Runtime: ACP, Docker or host sandboxes, Tauri desktop sidecar
Contributions and feedback are welcome.
Self-hosted AI coding workspace to run and orchestrate Claude Code, Codex, Copilot, Cursor, Grok and OpenCode agents — multi-agent workflows, personas, and ACP-powered sandboxes
Readme Apache-2.0 license Activity Stars
61 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
