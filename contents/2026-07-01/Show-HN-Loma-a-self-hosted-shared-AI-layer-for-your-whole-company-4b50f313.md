---
source: "https://github.com/plotlinelabs/loma"
hn_url: "https://news.ycombinator.com/item?id=48747049"
title: "Show HN: Loma – a self-hosted shared AI layer for your whole company"
article_title: "GitHub - plotlinelabs/loma: Loma: an AI agents factory for your whole team · GitHub"
author: "tadarsh"
captured_at: "2026-07-01T14:59:15Z"
capture_tool: "hn-digest"
hn_id: 48747049
score: 1
comments: 0
posted_at: "2026-07-01T14:07:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Loma – a self-hosted shared AI layer for your whole company

- HN: [48747049](https://news.ycombinator.com/item?id=48747049)
- Source: [github.com](https://github.com/plotlinelabs/loma)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T14:07:23Z

## Translation

タイトル: Show HN: Loma – 会社全体向けの自己ホスト型共有 AI レイヤー
記事のタイトル: GitHub - Lotlinelabs/loma: Loma: チーム全体のための AI エージェント ファクトリー · GitHub
説明: Loma: チーム全体のための AI エージェント ファクトリー。 GitHub でアカウントを作成して、plotlinelabs/loma の開発に貢献してください。

記事本文:
GitHub - Lotlinelabs/loma: Loma: チーム全体のための AI エージェント ファクトリー · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
プロットラインラボ
/
ロマ
公共
通知
通知セットを変更するにはサインインする必要があります

ひずみ
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
75 コミット 75 コミット .github/ workflows .github/ workflows エージェント エージェント api api config config ダッシュボード ダッシュボード デプロイ/ nginx デプロイ/ nginx docs docs ドラフト_with_loma ドラフト_with_loma ゲート ゲート統合 統合 mcp-servers/ bigquery mcp-servers/ bigquery メトリクス メトリクス メトリクス 可観測性 可観測性 プロンプト プロンプト スケジューラスケジューラ スクリプト スクリプト シード/スキル シード/スキル スラック_アプリ スラック_アプリ テスト テスト ツール ツール utils webhooks webhooks .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md app.py app.py config.yaml config.yaml docker-compose.tls.yml docker-compose.tls.yml docker-compose.yml docker-compose.yml Recovery.py Recovery.py Recovery_handlers.py Recovery_handlers.py Required.txt Required.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
チーム全体のための AI エージェント ファクトリー。
Slack とダッシュボード内の 1 つのエージェントがツールを認識し、オープン モデルまたはプールされたクロード サブスクリプション上で実行され、スケジュールと Webhook に基づいた作業を自動化し、時間の経過とともに改善されるチーム全体のスキル ライブラリを共有します。
ウェブサイト · ドキュメント · クイックスタート · 貢献 · セキュリティ
ほとんどのチームは AI の料金をシートごと、トークンごとに支払い、すべてのツールで同じコンテキストを手作業で接続し、チャットが終了した瞬間に有用なプロンプトをすべて失います。 Loma はその逆です。所有するインフラストラクチャ上で、会社全体で 1 つのエージェントを共有します。
🪙 チームのクロード コードのサブスクリプションをプールします。既存のクロード アカウントをラウンドロビン プールに接続すると、全員のエージェントの使用量がすでに支払ったサブスクリプションから引き出され、トークン コストが補助されます。

シートごとの API アクセスを購入する代わりに。
🧠 OpenCode 経由でオープンモデル上で実行します。デフォルトのランタイムは OpenCode なので、トークンごとの請求なしで DeepSeek V4 や GLM などのオープンソース モデルを入手でき、会話ごとにモデルを切り替えることができます。
⚡ Webhook とスケジュールされたフローで自動化します。エージェント タスクをルーチンに変換します。cron スケジュールで実行したり、CI/CD、サポート ツール、または組織内の任意のシステムから Webhook から実行したりできます。
📚 スキルを一度設定すると、チーム全体がそのスキルを使用します。プレイブックを一度作成してデータベースに保存すると、すべてのチームメイトとすべてのオートメーションが即座にそれを使用します。バージョン管理され、フィードバックによって自らを改善することができます。
さらに: Slack および Web ダッシュボードからアクセス可能で、既存のスタック (データベース、問題トラッカー、CRM、ドキュメント、可観測性) に接続し、すべての実行、トークン、コストをログに記録し、ダッシュボードから編集できる環境主導の構成で完全に自己ホスト可能です。
Loma は、アドホックな質問と常設の自動化にわたって同じエージェントを実行します。チームがこれを使用するいくつかのパターン:
Slack ────────┐ ┌── エージェントランタイム ── OpenCode（オープンモデル）
(DM / メンション) │ │ └─ クロードアカウントプール
▼ │
バックエンド (:3000) ───────┼── スキル (DB) + 接続ツール (MCP)
▲MongoDB │
ダッシュボード (:3001) ──┘ (会話、└── フロー: スケジュール + Webhook
(チャット、フロー、スキル、フロー、
スキル、構成) ユーザー、用途)
バックエンド — Python (aiohttp + Slack Bolt)。エージェントを実行し、API を提供し、Webhook を受信します。
ダッシュボード — Next.js。チャット、会話履歴、スキル、フロー、統合、構成、ユーザー/ロール、使用法。
ストレージ — ステートフルなものすべてに対応する MongoDB。ローカル ディスク上のスキル アセット ( LOMA_SKILL_ASSET_DIR )。
ランタイム — デフォルトでは OpenCode。オプションでクロード・コッドをプールする

電子アカウント。会話ごとに切り替え可能。
Loma を試す最も簡単な方法は Docker Compose です。 MongoDB 接続文字列、Slack アプリ (ソケット モード)、および OpenCode API キーが必要です。
git clone https://github.com/plotlinelabs/loma.git
CDロマ
cp .env.example .env # バックエンド構成
cp ダッシュボード/.env.example ダッシュボード/.env # ダッシュボード設定
# 両方を編集します — ドキュメントの env キーを参照してください
docker 構成 --build
次に、 http://localhost:3001 を開き、 LOMA_SETUP_TOKEN を使用して最初の管理者を作成し、チャットでメッセージを送信します。
完全なセットアップ (EC2/GCP の新規インストール、Slack アプリのスコープ、認証、MongoDB、OpenCode、完全な環境リファレンス) は ドキュメント に記載されています。
構成は環境主導型であり、ダッシュボードの「環境」ページから編集できます。キーを変更するための再構築は必要ありません。
デフォルトではオプション: プロバイダーの資格情報が欠落していても起動がブロックされることはなく、機能フラグ ( LOMA_ENABLE_SCHEDULER 、 LOMA_ENABLE_WEBHOOKS 、 LOMA_ENABLE_METRICS ) によってオプションのサブシステムが制御されます。
統合はダッシュボードから接続します: データベース (MongoDB、ClickHouse、BigQuery、Athena)、問題トラッカー (Linear、GitHub)、サポート (Pylon)、CRM (HubSpot、Apollo)、ナレッジ (Notion、GitBook、Google Workspace)、可観測性 (Sentry、PostHog) など。
プロバイダーごとのセットアップについては、統合ガイドを参照してください。
完全なガイドは lomahq.com/docs にあります。
はじめに — クイックスタート、EC2/GCP の新規インストール、ローカル開発
構成 - 環境参照、機能フラグ、ダッシュボード内構成
エージェント ランタイム - OpenCode、Claude アカウント プーリング、モデル選択
スキル — オーサリング、インポート、アセット、バージョン管理
フローと自動化 - スケジュールされたルーチンと Webhook トリガー
統合 — ツールを接続する
Slack アプリ、認証、展開とネットワーキング、セキュリティ
app.py バックエンド エントリポイント (aiohttp + Slack Bolt)
エージェント/

エージェント ランタイム: OpenCode + Claude アカウント プール
api/ HTTP API (チャット、スキル、フロー、環境、ガバナンスなど)
スケジューラー/スケジュールされた Webhook フローの実行
Webhook/インバウンド Webhook ハンドラー
統合/接続可能なツールのレジストリ
エージェントが呼び出すことができるツール/CLI ツール (スキル + 統合)
MongoDB への可観測性/会話/使用状況のログ記録
slack_app/Slackイベント処理
ダッシュボード/ Next.js ダッシュボード
貢献する
問題やPRは大歓迎です。 COTRIBUTING.md を参照してください。このリポジトリはクリーンな OSS コードベースです。企業固有の知識、プロンプト、プレイブック、資格情報はデータベースと環境に属し、ソースには決して含まれません。
.env 、資格情報、プライベート プロンプト、または顧客データを決してコミットしないでください。脆弱性を報告するには、 SECURITY.md を参照してください。
Loma: チーム全体のための AI エージェント ファクトリー
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
1
フォーク
レポートリポジトリ
リリース
1
ロマ v1.0.0
最新の
2026 年 6 月 23 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Loma: an AI agents factory for your whole team. Contribute to plotlinelabs/loma development by creating an account on GitHub.

GitHub - plotlinelabs/loma: Loma: an AI agents factory for your whole team · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
plotlinelabs
/
loma
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
75 Commits 75 Commits .github/ workflows .github/ workflows agent agent api api config config dashboard dashboard deploy/ nginx deploy/ nginx docs docs draft_with_loma draft_with_loma gate gate integrations integrations mcp-servers/ bigquery mcp-servers/ bigquery metrics metrics observability observability prompts prompts scheduler scheduler scripts scripts seed/ skills seed/ skills slack_app slack_app tests tests tools tools utils utils webhooks webhooks .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md app.py app.py config.yaml config.yaml docker-compose.tls.yml docker-compose.tls.yml docker-compose.yml docker-compose.yml recovery.py recovery.py recovery_handlers.py recovery_handlers.py requirements.txt requirements.txt View all files Repository files navigation
An AI Agents Factory for your whole team.
One agent — in Slack and a dashboard — that knows your tools, runs on open models or your pooled Claude subscriptions, automates work on schedules and webhooks, and shares a team-wide skill library that gets better over time.
Website · Docs · Quickstart · Contributing · Security
Most teams pay for AI per seat and per token, wire up the same context by hand in every tool, and lose every useful prompt the moment the chat ends. Loma is the opposite: one agent your whole company shares, on infrastructure you own.
🪙 Pool your team's Claude Code subscriptions. Connect existing Claude accounts into a round-robin pool so everyone's agent usage draws from subscriptions you already pay for — subsidizing token cost instead of buying per-seat API access.
🧠 Run on open models via OpenCode. The default runtime is OpenCode , so you get open-source models like DeepSeek V4 and GLM with no per-token bill — and you can switch models per conversation.
⚡ Automate with webhook & scheduled flows. Turn any agent task into a routine: run it on a cron schedule, or fire it from a webhook out of your CI/CD, support tool, or any system in your org.
📚 Set up skills once, the whole team uses them. Write a playbook once, store it in the database, and every teammate and every automation uses it instantly — versioned, and able to improve itself from feedback.
Plus: reachable from Slack and a web dashboard , connects to your existing stack (databases, issue trackers, CRM, docs, observability), logs every run, token, and cost , and is fully self-hostable with env-driven config you can edit from the dashboard.
Loma runs the same agent across ad-hoc questions and standing automations. A few patterns teams use it for:
Slack ─────────────┐ ┌── Agent runtime ── OpenCode (open models)
(DM / mention) │ │ └─ Claude account pool
▼ │
Backend (:3000) ────────────────┼── Skills (DB) + your connected tools (MCP)
▲ MongoDB │
Dashboard (:3001) ──┘ (conversations, └── Flows: schedules + webhooks
(chat, flows, skills, flows,
skills, config) users, usage)
Backend — Python (aiohttp + Slack Bolt). Runs the agent, serves the API, receives webhooks.
Dashboard — Next.js. Chat, conversation history, skills, flows, integrations, config, users/roles, usage.
Storage — MongoDB for everything stateful; skill assets on local disk ( LOMA_SKILL_ASSET_DIR ).
Runtime — OpenCode by default; optionally pool Claude Code accounts. Switchable per conversation.
The fastest way to try Loma is Docker Compose. You need a MongoDB connection string, a Slack app (Socket Mode), and an OpenCode API key.
git clone https://github.com/plotlinelabs/loma.git
cd loma
cp .env.example .env # backend config
cp dashboard/.env.example dashboard/.env # dashboard config
# edit both — see the env keys in the docs
docker compose up --build
Then open http://localhost:3001 , create the first admin with your LOMA_SETUP_TOKEN , and send a message in chat.
Full setup — fresh EC2/GCP install, Slack app scopes, auth, MongoDB, OpenCode, and the complete environment reference — is in the documentation .
Config is env-driven and editable from the dashboard's Environment page — no rebuilds to change keys.
Optional by default: missing provider credentials never block startup, and feature flags ( LOMA_ENABLE_SCHEDULER , LOMA_ENABLE_WEBHOOKS , LOMA_ENABLE_METRICS ) gate optional subsystems.
Integrations connect from the dashboard: databases (MongoDB, ClickHouse, BigQuery, Athena), issue trackers (Linear, GitHub), support (Pylon), CRM (HubSpot, Apollo), knowledge (Notion, GitBook, Google Workspace), and observability (Sentry, PostHog), among others.
See the integrations guide for per-provider setup.
Full guides live at lomahq.com/docs :
Getting started — Quickstart, fresh EC2/GCP install, local development
Configuration — environment reference, feature flags, in-dashboard config
Agent runtime — OpenCode, Claude account pooling, model selection
Skills — authoring, importing, assets, versioning
Flows & automations — scheduled routines and webhook triggers
Integrations — connecting your tools
Slack app, authentication, deployment & networking, security
app.py Backend entrypoint (aiohttp + Slack Bolt)
agent/ Agent runtime: OpenCode + Claude account pool
api/ HTTP API (chat, skills, flows, env, governance, ...)
scheduler/ Scheduled & webhook flow execution
webhooks/ Inbound webhook handlers
integrations/ Connectable-tool registry
tools/ CLI tools the agent can call (skills + integrations)
observability/ Conversation/usage logging to MongoDB
slack_app/ Slack event handling
dashboard/ Next.js dashboard
Contributing
Issues and PRs are welcome. See CONTRIBUTING.md . This repository is the clean OSS codebase — company-specific knowledge, prompts, playbooks, and credentials belong in your database and environment, never in source.
Never commit .env , credentials, private prompts, or customer data. To report a vulnerability, see SECURITY.md .
Loma: an AI agents factory for your whole team
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
1
Loma v1.0.0
Latest
Jun 23, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
