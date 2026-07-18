---
source: "https://github.com/dylanneve1/talon"
hn_url: "https://news.ycombinator.com/item?id=48959516"
title: "Show HN: Talon – a self-hosted harness for long-lived AI agents"
article_title: "GitHub - dylanneve1/talon: 🦅 Multi-platform agentic AI harness — runs on Telegram, Discord, Teams & Terminal with a pluggable backend (Claude, Kilo, OpenCode, Codex, OpenAI Agents), full MCP tool access, and persistent background agents (Goals, Heartbeat, Dream). · GitHub"
author: "claudiusthebot"
captured_at: "2026-07-18T16:44:12Z"
capture_tool: "hn-digest"
hn_id: 48959516
score: 1
comments: 0
posted_at: "2026-07-18T16:24:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Talon – a self-hosted harness for long-lived AI agents

- HN: [48959516](https://news.ycombinator.com/item?id=48959516)
- Source: [github.com](https://github.com/dylanneve1/talon)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T16:24:35Z

## Translation

タイトル: 表示 HN: Talon – 寿命の長い AI エージェント用の自己ホスト型ハーネス
記事のタイトル: GitHub - dylanneve1/talon: 🦅 マルチプラットフォーム エージェント AI ハーネス — プラグ可能なバックエンド (Claude、Kilo、OpenCode、Codex、OpenAI Agents)、完全な MCP ツール アクセス、永続的なバックグラウンド エージェント (Goals、Heartbeat、Dream) を備えた Telegram、Discord、Teams、ターミナル上で実行されます。 · GitHub
説明: 🦅 マルチプラットフォーム エージェント AI ハーネス — プラグ可能なバックエンド (Claude、Kilo、OpenCode、Codex、OpenAI Agents)、完全な MCP ツール アクセス、および永続的なバックグラウンド エージェント (Goals、Heartbeat、Dream) を備えた Telegram、Discord、Teams、および Terminal 上で実行されます。 - dylanneve1/タロン

記事本文:
GitHub - dylanneve1/talon: 🦅 マルチプラットフォーム エージェント AI ハーネス — プラグ可能なバックエンド (Claude、Kilo、OpenCode、Codex、OpenAI Agents)、完全な MCP ツール アクセス、永続的なバックグラウンド エージェント (Goals、Heartbeat、Dream) を備えた Telegram、Discord、Teams、および Terminal 上で実行されます。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードして更新します

セッション。
アラートを閉じる
{{ メッセージ }}
ディランネブ1
/
爪
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
940 コミット 940 コミット .github .github apps/ コンパニオン アプリ/ コンパニオン bin bin docker docker docs docs ネイティブ ネイティブ パッケージング パッケージング プロンプト プロンプト スクリプト スクリプト src src .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml .prettierignore .prettierignore .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile ライセンス ライセンス MESH_BUILD_REPORT.md MESH_BUILD_REPORT.md README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml nfpm.yaml nfpm.yaml package-lock.json package-lock.json package.json package.json release-please-config.json release-please-config.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリファイルナビゲーション
マルチプラットフォームのエージェント AI ハーネス。 Telegram 、 Discord 、 Microsoft Teams 、 Terminal 、およびクロスプラットフォームのデスクトップ/モバイル コンパニオン アプリ (Flutter) 上で実行され、プラグ可能なバックエンド ( Claude Agent SDK 、 Kilo 、 OpenCode 、 Codex 、または OpenAI Agents ) と MCP を介したフル ツール アクセスを備えています。
マルチフロントエンド
Telegram (Grammy + GramJS userbot)、Discord (discord.js)、Microsoft Teams (Bot Framework)、ライブ ツールの可視性を備えたターミナル、ローカル/リモート ブリッジ経由のデスクトップ/モバイル アプリ (Flutter)
プラグ可能なバックエンド
Claude Agent SDK、Kilo、OpenCode、Codex、OpenAI エージェント — バックエンド構成を介してプロセスごとに選択可能。ストリーミング、モデルのフォールバック、コンテキスト オーバーフローの回復。
MCPツール
メッセージング、メディア、履歴、検索、Web フェッチ、クロ

n ジョブ、トリガー、目標、ステッカー、ファイル システム、管理者コントロール
プラグイン
talon プラグインのインストール/有効化/無効化を備えたホットリロード可能なプラグイン システム (npm、git、またはローカル ソース)。組み込み: GitHub、MemPalace、Playwright、Brave Search
バックグラウンドエージェント
ハートビート (デフォルトでは 1 時間ごと — 目標を前進させ、何か重要なことがあれば積極的にメッセージを送信します) とドリーム (記憶の統合 + 日記)
目標
エージェントがチャットでコミットする永続的な複数日間の目標。ハートビートを実行するたびにハートビートが再読み取りされ、進捗が確認され、その内容が記録されます。
スキル
SKILL.md ワークフローは、エージェントの作成者をバンドルし、talon スキルのインストール/有効化/無効化 (ローカル フォルダー、git、または所有者/リポジトリ - Anthropic スキル エコシステムが直接インストールされます) を使用して再利用します。
トリガー
条件が満たされたときにボットを起動する自己作成のウォッチャー スクリプト (bash/Python/node)
タスクテーブル
エージェントの作業のすべての単位 (チャット ターン、ハートビート、ドリーム、分離された cron/トリガー ジョブ) がライブで登録されます。タロン ps / タロン キル
イベントバス
型付きの内部パブリッシュ/サブスクイン スパイン (タスク + ターン ライフサイクル イベント)。サブシステムは相互にインポートする代わりにサブスクライブします。タロンイベント -f
VFS
ワークスペース、スキル、スクリプト、ログ上の ~/.talon/ns に統合された名前空間、さらにタスク テーブル、イベント バス、プラグイン レジストリの /proc スタイルのライブ ビュー — 実際のファイルシステム (FUSE でバックアップされたライブ ビュー) なので、プレーンな ls / cat とすべてのツールが正常に動作します
チャットごとの設定
インライン キーボードを介した会話ごとのモデル、エフォート レベル、パルスの切り替え
モデルレジストリ
起動時にアクティブなバックエンドからモデルが検出されます - 新しいモデルがすべてのピッカーに自動的に表示されます
クイックスタート
git clone https://github.com/dylanneve1/talon.git && cd talon
npmインストール
# 対話型セットアップ (フロントエンドの選択、トークンの構成、モデルの選択)
npx talon のセットアップ
# 開始
npx talon start # 構成されたフロントエンド (デーモンモード)
npxタロンチャット#te

ターミナルチャットモード
前提条件:
バックエンド固有:
クロード バックエンド: クロード コードがインストールされ、認証されました (PATH 上のクロード CLI)。
キロ バックエンド: 余分なものは何もありません — @kilocode/sdk はローカル サーバーを生成します。無料モデルは認証なしでアクセスできます。ルーティングされたモデルは、Kilo 独自の認証情報を使用します。
opencode バックエンド: 余分なものは何もありません — @opencode-ai/sdk はローカル サーバーを生成します。
codex バックエンド: codex CLI ( npm i -g @openai/codex ) をインストールし、 codex login 、 CODEX_API_KEY 、 TALON_CODEX_KEY 、または codexApiKey で認証します。 OPENAI_API_KEY は、Codex ログインが存在しない場合のフォールバックとしてのみ使用されます。
各リリースには、自己完結型バイナリ (Node.js は不要) も同梱されています。
Linux および macOS (x64 + arm64) および Windows (x64)。プロンプトとすべてネイティブ
モジュールはバイナリに埋め込まれます。
# 自作 (macOS / Linux)
醸造インストール dylanneve1/talon/talon
# Debian / Ubuntu — リリースからアーチの .deb をダウンロードし、次の手順を実行します。
sudo apt install ./talon_ < バージョン > _amd64.deb # または _arm64.deb
# 直接ダウンロード — リリースから talon-<os>-<arch> を取得し、確認して実行します。
chmod +x talon-linux-x64 && ./talon-linux-x64 --version
# macOS、Gatekeeper が署名されていないバイナリをブロックする場合:
xattr -d com.apple.quarantine ./talon-darwin-arm64
リリース SHA256SUMS に対する直接ダウンロードを確認します。
sha256sum -c SHA256SUMS --ignore-missing 。
バイナリは完全な対話型/エージェント CLI ( setup 、 start 、 chat 、
医者、…）。 Talon 独自の MCP サーバーをホスト (MCP 標準出力としての talon)
別のクライアント用のサーバー）には、依然として npm/Node のインストールが必要です。
tsx ローダーはコンパイルされたバイナリには存在しません。
Index.ts 構成ルート
|
+-- コア/プラットフォームに依存しないエンジン
| +-- エージェント ランタイム/バックエンド機能インターフェイス、イベント、ストア
| +-- モデル/モデル レイヤー: カタログ、チャットごとのアクティブ モデル、
| |推論努力の語彙
| +-- プロ

mpt/ システム プロンプト アセンブリ + プロンプト/システム テンプレート
| +-- バックグラウンド/ ユーザー メッセージなしで実行されるエージェント:
| |ハートビート、夢、パルス、クロン、トリガー
| +-- ツール/MCP ツール定義 + spawn/env コントラクト
| +-- エンジン/メッセージ フロー: ディスパッチャ (チャットごとのシリアル、
| |クロスチャットパラレル)、MCP の HTTP ゲートウェイ
| |ツール呼び出し、バックエンド ライフサイクル コントローラー
| +-- plugin.ts プラグインローダー、レジストリ、ホットリロード
|
+-- バックエンド/
| +-- registry.ts ブートストラップ分離されたバックエンド ルックアップ
| +-- 共有/クロスバックエンド ヘルパー (ストリーム状態、フロー違反、
| |配信契約、メトリクス、プロンプト形式、
| |モデルの再試行、システム プロンプト、使用法)
| +-- リモートサーバー/エージェントサーバーバックエンドの共有インフラストラクチャ
| | (MCP 登録、セッション、プロバイダー、ライフサイクル)
| +-- claude-sdk/ クロード エージェント SDK (インプロセス MCP、フック)
| +-- キロ/キロ HTTP サーバー バックエンド (SSE 経由のストリーミング)
| +-- opencode/ OpenCode HTTP サーバー バックエンド
| +-- codex/ Codex CLI バックエンド (`@openai/codex-sdk`)
| +-- openai-agents/ OpenAI Agents SDK バックエンド (レスポンス API)
|
+-- フロントエンド/
| +-- 共有/クロスフロントエンド プレゼンテーション ヘルパー
| +-- 電報/グラミーボット + GramJS ユーザーボット
| +-- discord/ discord.js v14
| +-- チーム/ボット フレームワーク + グラフ API
| +-- ツール呼び出しの可視性を備えたターミナル/Readline CLI
| +-- コンパニオン アプリのデスクトップ/クライアント ブリッジ (HTTP + SSE)
|
+-- ストレージ/ セッション、履歴、チャット設定、
| cron ジョブ、メディア インデックス、毎日のログ
+-- util/ 構成、ログ、ワークスペース、パス、時間
依存関係ルール: core/ は、frontend/ または backend/ から何もインポートしません。フロントエンドとバックエンドはコアの種類に依存しており、互いに依存することはありません。 5 つのバックエンド (Claude SDK、Kilo、OpenCode、Codex、OpenAI Agents) はすべて、 core/agent-runtime/capabilities.ts から同じバックエンド機能インターフェイスを実装します。キロと
[切り捨てられた]
プロンプト: すべて

セッション開始時に読み取られるモデルは、prompts/ 内のファイルから core/prompt/ によって組み立てられます。組み立て順序、ファイルの所有権 (ユーザーが編集可能なテンプレートとパッケージ所有のテンプレート)、およびバックエンドごとの配信コントラクトについては、prompts/README.md を参照してください。
~/.talon/config.json のバックエンドフィールドから選択します。すべてのバックエンドは同じバックエンド機能インターフェイスを実装しています。ハートビート、ドリーム、チャット ハンドラーはバックエンドに依存しません。
アップストリームの HTTP API が同じであるため、Kilo と OpenCode のバックエンドはインフラストラクチャ ( backend/remote-server/ ) を共有します。各バックエンドは、独自の SDK クライアント、ポート、および配信サフィックスを提供します。 Codex は、Codex CLI の JSONL イベント ストリーム上に独自に統合されたものです。
デスクトップ フロントエンドは、デーモンをクライアント ブリッジに変えます。これは、あらゆる GUI クライアントが通信できる、バージョン管理された HTTP + Server-Sent-Events JSON API (Talon Client Bridge Protocol 、 src/frontend/desktop/protocol.ts ) です。リファレンス クライアントは、Windows、macOS、Linux、Android 上で実行される単一の Flutter コードベースである Talon Companion です。
// ~/.talon/config.json
{
"フロントエンド" : "デスクトップ" ,
"デスクトップ" : { "ホスト" : " 127.0.0.1 " 、 "ポート" : 19880 }
// リモート (電話など) の場合: "host": "0.0.0.0", "token": "your-secret"
}
ローカル (デスクトップ): アプリは同じマシン上の Talon に接続し、必要に応じて Talon を起動します ( TALON_FRONTEND_OVERRIDE=desktop )。
リモート (モバイル/LAN): アプリをホスト:ポート + トークンに向けます。トークンが設定されるたびに、ブリッジは Authorization: Bearer … (または SSE ストリームの ?token=) を必要とします。
暗号化: オフループバック バインドは、デフォルトで永続的な自己署名証明書 ( ~/.talon/keys/ ) を使用して HTTPS を提供します。コンパニオンは最初の接続時に SHA-256 フィンガープリントを固定し、その後の変更を拒否します。デーモンは起動時にフィンガープリントを記録し、/health がそれをアドバタイズします。 " を使用してオプトアウト (またはループバックでオプトアウト)

tls": デスクトップセクションの false / true。
このアプリは、マルチチャット履歴、推論とツール呼び出しの可視性を備えたライブストリーミング、チャットごとのモデル/エフォート/パルス/リセット、および設定の同期を提供します。デーモン自体の構成 (デフォルトのモデル、表示名、タイムゾーン、パルス/ハートビート/ドリーム) を読み取り、変更し、再起動します。 apps/companion/README.md を参照してください。
どちらのストアも CLI から管理されます。ホットリロードを実行中に変更します
デーモン (プラグイン) または次のセッションに適用 (スキル):
# プラグイン — npm 仕様、git リポジトリ、またはローカル パス
talon プラグインのインストール @scope/my-talon-plugin # npm → モジュールプラグイン
talon プラグインのインストール some-mcp-server --mcp # npm → スタンドアロン MCP サーバー (npx)
talon プラグインのインストール所有者/リポジトリ # git → モジュールプラグイン
talon プラグイン リスト # 組み込み + 設定されたエントリ
talon プラグインは github を無効にします # 組み込みも切り替えます
talon プラグイン my-talon-plugin を削除します
# スキル — ローカル パス、git URL、または所有者/リポジトリ[/サブパス] からの SKILL.md フォルダー
talon スキル インストール anthropics/skills/document-skills/pdf
talon スキルのインストール ./my-skill --force
タロンのスキルリスト
talon skill disable pdf # プロンプトインデックスからは隠されていますが、引き続き読むことができます
タロンスキル削除PDF
モジュールプラグインは ~/.talon/plugins/ にインストールされます。スタンドアロン MCP サーバーは、
config に npx エントリとして登録

[切り捨てられた]

## Original Extract

🦅 Multi-platform agentic AI harness — runs on Telegram, Discord, Teams & Terminal with a pluggable backend (Claude, Kilo, OpenCode, Codex, OpenAI Agents), full MCP tool access, and persistent background agents (Goals, Heartbeat, Dream). - dylanneve1/talon

GitHub - dylanneve1/talon: 🦅 Multi-platform agentic AI harness — runs on Telegram, Discord, Teams & Terminal with a pluggable backend (Claude, Kilo, OpenCode, Codex, OpenAI Agents), full MCP tool access, and persistent background agents (Goals, Heartbeat, Dream). · GitHub
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
dylanneve1
/
talon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
940 Commits 940 Commits .github .github apps/ companion apps/ companion bin bin docker docker docs docs native native packaging packaging prompts prompts scripts scripts src src .dockerignore .dockerignore .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml .prettierignore .prettierignore .release-please-manifest.json .release-please-manifest.json CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE MESH_BUILD_REPORT.md MESH_BUILD_REPORT.md README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml nfpm.yaml nfpm.yaml package-lock.json package-lock.json package.json package.json release-please-config.json release-please-config.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Multi-platform agentic AI harness. Runs on Telegram , Discord , Microsoft Teams , the Terminal , and a cross-platform Desktop/Mobile companion app (Flutter), with a pluggable backend ( Claude Agent SDK , Kilo , OpenCode , Codex , or OpenAI Agents ) and full tool access through MCP.
Multi-frontend
Telegram (Grammy + GramJS userbot), Discord (discord.js), Microsoft Teams (Bot Framework), Terminal with live tool visibility, and a Desktop/Mobile app (Flutter) over a local/remote bridge
Pluggable backend
Claude Agent SDK, Kilo, OpenCode, Codex, OpenAI Agents — selectable per-process via backend config. Streaming, model fallback, context-overflow recovery.
MCP tools
Messaging, media, history, search, web fetch, cron jobs, triggers, goals, stickers, file system, admin controls
Plugins
Hot-reloadable plugin system with talon plugin install/enable/disable (npm, git, or local sources). Built-in: GitHub, MemPalace, Playwright, Brave Search
Background agents
Heartbeat (hourly by default — advances goals, proactively messages when something matters) and Dream (memory consolidation + diary)
Goals
Persistent multi-day objectives the agent commits to in chat; every heartbeat run re-reads them, makes progress, and records what it did
Skills
SKILL.md workflow bundles the agent authors and reuses, with talon skill install/enable/disable (local folders, git, or owner/repo — the Anthropic skills ecosystem installs directly)
Triggers
Self-authored watcher scripts (bash/python/node) that wake the bot when conditions are met
Task table
Every unit of agent work — chat turns, heartbeat, dream, isolated cron/trigger jobs — registered live; talon ps / talon kill
Event bus
Typed internal pub-sub spine (task + turn lifecycle events); subsystems subscribe instead of importing each other; talon events -f
VFS
Unified namespace at ~/.talon/ns over workspace, skills, scripts, logs, plus /proc-style live views of the task table, event bus, and plugin registry — a real filesystem (FUSE-backed live views), so plain ls / cat and every tool just work
Per-chat settings
Model, effort level, and pulse toggle per conversation via inline keyboard
Model registry
Models discovered from the active backend at startup — new models appear in all pickers automatically
Quick Start
git clone https://github.com/dylanneve1/talon.git && cd talon
npm install
# Interactive setup (select frontend, configure tokens, pick model)
npx talon setup
# Start
npx talon start # configured frontend (daemon mode)
npx talon chat # terminal chat mode
Prerequisites:
Backend-specific:
claude backend: Claude Code installed and authenticated ( claude CLI on PATH).
kilo backend: nothing extra — @kilocode/sdk spawns a local server. Free models are accessible without auth; routed models use Kilo's own credentials.
opencode backend: nothing extra — @opencode-ai/sdk spawns a local server.
codex backend: install the codex CLI ( npm i -g @openai/codex ) and authenticate with codex login , CODEX_API_KEY , TALON_CODEX_KEY , or codexApiKey . OPENAI_API_KEY is used only as a fallback when no Codex login exists.
Each release also ships self-contained binaries (no Node.js required) for
Linux and macOS (x64 + arm64) and Windows (x64). Prompts and all native
modules are embedded in the binary.
# Homebrew (macOS / Linux)
brew install dylanneve1/talon/talon
# Debian / Ubuntu — download the .deb for your arch from the release, then:
sudo apt install ./talon_ < version > _amd64.deb # or _arm64.deb
# Direct download — grab talon-<os>-<arch> from the release, verify, run:
chmod +x talon-linux-x64 && ./talon-linux-x64 --version
# macOS, if Gatekeeper blocks an unsigned binary:
xattr -d com.apple.quarantine ./talon-darwin-arm64
Verify a direct download against the release SHA256SUMS :
sha256sum -c SHA256SUMS --ignore-missing .
The binary runs the full interactive/agent CLI ( setup , start , chat ,
doctor , …). Hosting Talon's own MCP server ( talon as an MCP stdio
server for another client) still needs the npm/Node install — it spawns a
tsx loader that isn't present in a compiled binary.
index.ts Composition root
|
+-- core/ Platform-agnostic engine
| +-- agent-runtime/ Backend capability interfaces, events, stores
| +-- models/ Model layer: catalog, per-chat active model,
| | reasoning-effort vocabulary
| +-- prompt/ System-prompt assembly + prompts/system templates
| +-- background/ Agents that run without a user message:
| | heartbeat, dream, pulse, cron, triggers
| +-- tools/ MCP tool definitions + spawn/env contract
| +-- engine/ Message flow: dispatcher (per-chat serial,
| | cross-chat parallel), HTTP gateway for MCP
| | tool calls, backend lifecycle controller
| +-- plugin.ts Plugin loader, registry, hot-reload
|
+-- backend/
| +-- registry.ts Bootstrap-decoupled backend lookup
| +-- shared/ Cross-backend helpers (stream state, flow violation,
| | delivery contract, metrics, prompt format,
| | model retry, system prompt, usage)
| +-- remote-server/ Shared infrastructure for agent-server backends
| | (MCP registration, sessions, providers, lifecycle)
| +-- claude-sdk/ Claude Agent SDK (in-process MCP, hooks)
| +-- kilo/ Kilo HTTP server backend (streaming via SSE)
| +-- opencode/ OpenCode HTTP server backend
| +-- codex/ Codex CLI backend (`@openai/codex-sdk`)
| +-- openai-agents/ OpenAI Agents SDK backend (Responses API)
|
+-- frontend/
| +-- shared/ Cross-frontend presentation helpers
| +-- telegram/ Grammy bot + GramJS userbot
| +-- discord/ discord.js v14
| +-- teams/ Bot Framework + Graph API
| +-- terminal/ Readline CLI with tool call visibility
| +-- desktop/ Client bridge (HTTP + SSE) for the companion app
|
+-- storage/ Sessions, history, chat settings,
| cron jobs, media index, daily logs
+-- util/ Config, logging, workspace, paths, time
Dependency rule: core/ imports nothing from frontend/ or backend/ . Frontends and backends depend on core types, never on each other. All five backends (Claude SDK, Kilo, OpenCode, Codex, OpenAI Agents) implement the same Backend capability interface from core/agent-runtime/capabilities.ts . Kilo and
[truncated]
Prompts: everything the model reads at session start is assembled by core/prompt/ from the files in prompts/ — see prompts/README.md for the assembly order, file ownership (user-editable vs package-owned templates), and the per-backend delivery contracts.
Select via the backend field in ~/.talon/config.json . All backends implement the same Backend capability interface — heartbeat, dream, and chat handlers are backend-agnostic.
The Kilo and OpenCode backends share infrastructure ( backend/remote-server/ ) since the upstream HTTP API is the same; each backend supplies its own SDK client, port, and delivery suffix. Codex is its own integration on top of the Codex CLI's JSONL event stream.
The desktop frontend turns the daemon into a client bridge — a versioned HTTP + Server-Sent-Events JSON API (the Talon Client Bridge Protocol , src/frontend/desktop/protocol.ts ) that any GUI client can speak. The reference client is Talon Companion , a single Flutter codebase that runs on Windows, macOS, Linux, and Android .
// ~/.talon/config.json
{
"frontend" : " desktop " ,
"desktop" : { "host" : " 127.0.0.1 " , "port" : 19880 }
// For remote (e.g. a phone): "host": "0.0.0.0", "token": "your-secret"
}
Local (desktop): the app connects to a Talon on the same machine and launches one if needed ( TALON_FRONTEND_OVERRIDE=desktop ).
Remote (mobile/LAN): point the app at host:port + token; the bridge requires Authorization: Bearer … (or ?token= on the SSE stream) whenever a token is set.
Encryption: off-loopback binds serve HTTPS by default with a persistent self-signed certificate ( ~/.talon/keys/ ); the companion pins its SHA-256 fingerprint on first connect and refuses any change afterwards. The daemon logs the fingerprint at startup and /health advertises it. Opt out (or in, on loopback) with "tls": false / true in the desktop section.
The app provides multi-chat history, live streaming with reasoning + tool-call visibility, per-chat model/effort/pulse/reset, and settings sync — read and change the daemon's own config (default model, display name, timezone, pulse/heartbeat/dream) and restart it. See apps/companion/README.md .
Both stores are managed from the CLI; changes hot-reload into a running
daemon (plugins) or apply on the next session (skills):
# Plugins — npm specs, git repos, or local paths
talon plugin install @scope/my-talon-plugin # npm → module plugin
talon plugin install some-mcp-server --mcp # npm → standalone MCP server (npx)
talon plugin install owner/repo # git → module plugin
talon plugin list # built-ins + configured entries
talon plugin disable github # also toggles built-ins
talon plugin remove my-talon-plugin
# Skills — SKILL.md folders from local paths, git URLs, or owner/repo[/subpath]
talon skill install anthropics/skills/document-skills/pdf
talon skill install ./my-skill --force
talon skill list
talon skill disable pdf # hidden from the prompt index, still readable
talon skill remove pdf
Module plugins install under ~/.talon/plugins/ ; standalone MCP servers are
registered as npx entries in config

[truncated]
