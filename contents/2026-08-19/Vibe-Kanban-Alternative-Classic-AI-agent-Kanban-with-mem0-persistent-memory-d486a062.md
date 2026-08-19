---
source: "https://github.com/flashlan/vibe-kanban-alternative"
hn_url: "https://news.ycombinator.com/item?id=49367532"
title: "Vibe-Kanban-Alternative – Classic AI agent Kanban with mem0 persistent memory"
article_title: "GitHub - flashlan/vibe-kanban-alternative: vibe-kanban-indie fork · GitHub"
image: "https://opengraph.githubassets.com/90fe6ae7116c9a20dfaec399a0f541fab21f56bdc3407e2aacf329359d347610/flashlan/vibe-kanban-alternative"
author: "datapointnet"
captured_at: "2026-08-19T22:14:49Z"
capture_tool: "hn-digest"
hn_id: 49367532
score: 1
comments: 0
posted_at: "2026-08-19T21:38:25Z"
tags:
  - hacker-news
  - translated
---

# Vibe-Kanban-Alternative – Classic AI agent Kanban with mem0 persistent memory

- HN: [49367532](https://news.ycombinator.com/item?id=49367532)
- Source: [github.com](https://github.com/flashlan/vibe-kanban-alternative)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T21:38:25Z

## Translation

タイトル: Vibe-Kanban-Alternative – mem0 永続メモリを備えたクラシック AI エージェント カンバン
記事タイトル: GitHub - flashlan/vibe-kanban-alternative: vibe-kanban-indie fork · GitHub
説明: バイブカンバンインディーフォーク。 GitHub でアカウントを作成して、flashlan/vibe-kanban-alternative の開発に貢献してください。

記事本文:
GitHub - flashlan/vibe-kanban-alternative: vibe-kanban-indie fork · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
フラッシュラン
/
バイブカンバン代替
公共
dexloom/vibe-kanban-indie からフォーク
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
2,456 コミット 2,456 コミット フォルダーとファイル
.cargo .cargo .github .github apps/ macos apps/ macos assets assets automation automation crates crates dev_assets_seed dev_assets_seed docs do

cs npx-cli npx-cli パッケージ パッケージ パッチ パッチ スクリプト スクリプト 共有 共有 .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE-OF-CONDUCT.md CODE-OF-CONDUCT.md CONTEXT.md CONTEXT.md CONTRIBUTORS.md CONTRIBUTORS.md Caddyfile.example Caddyfile.example Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile PUBLISHING.md PUBLISHING.md README.md README.md SPEC-conversation-list.md SPEC-conversation-list.md SPEC.md SPEC.md TODO.md TODO.md gitea.toml.example gitea.toml.example local-build.sh local-build.sh mobile-testing.md mobile-testing.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml Projects.toml.example Projects.toml.example restart.sh restart.sh Rust-toolchain.toml Rust-toolchain.toml Rustfmt.toml Rustfmt.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
単一開発者による AI オーケストレーションのための、復元力に優れた 100% 自己ホスト型コックピット。
永続グラフ メモリ (mem0) • 10 個以上のコーディング エージェント (反重力/AGY を含む) • Gitea/Forgejo PR • TUI コックピット • テレグラム デーモン • 使用状況の可観測性
クイックスタート •
概要 •
何が違うのか？
プロジェクトメモリ (mem0) •
サポートされているエージェント •
チャットとコントロール •
使用状況ダッシュボード •
TUI コックピット •
テレグラムブリッジ •
開発
単一のコマンドでコックピット全体を起動します。インストール、アカウント、クラウドのセットアップは必要ありません。
npx バイブカンバン代替品
💡 セットアップ不要: 事前に構築されたバイナリをダウンロードし、ローカル Web コックピットを http://localhost:3001 で直接起動します (バックエンドは :3002 です)。
⚙️ 完全なセットアップと環境構成
開発用、カスタム ポート、ローカル mem0 ベクター ストア、およびカスタム エージェント構成

s:
git clone https://github.com/flashlan/vibe-kanban-alternative.git
CDバイブカンバンオルタナティブ
pnpm私
2. 環境変数を設定する
テンプレート構成ファイルを複製します。
cp .env.example .env
エディターで .env を開き、環境を構成します。
# =========================================
# 🌐 コアサーバー構成
# =========================================
ポート = 3000
ホスト = ローカルホスト
NODE_ENV = 開発
# =========================================
# 🧠 MEM0 長期メモリ設定
# =========================================
MEM0_ENABLED = true
MEM0_API_KEY = your_mem0_api_key_here
# ローカル埋め込み/ベクター ストアを実行している場合:
# MEM0_VECTOR_STORE=qdrant
# MEM0_HOST=http://localhost:6333
# =========================================
# 🤖 エージェント API キーとランタイム
# =========================================
ANTHROPIC_API_KEY = sk-ant-...
OPENAI_API_KEY = sk-...
GEMINI_API_KEY = ...
# 反重力 (AGY) 固有の設定
AGY_CLI_PATH = /usr/local/bin/agy
AGY_DEFAULT_TEMPERATURE = 0.2
# =========================================
# 📊 メトリクスとテレメトリ
# =========================================
ENABLE_METRICS_DASHBOARD = true
METRICS_STORAGE_PATH = ./data/metrics.sqlite
# =========================================
# 💾 バックアップ構成
# =========================================
BACKUP_ENABLED = true
バックアップ_インターバル_分 = 30
BACKUP_STORAGE_PATH = ./バックアップ
3. 開発コックピットの開始
./restart.sh
📌 コンテキスト: Vibe-Kanban の復活と進化
Bloop のホスト型サーバーがクラウド上で正式にシャットダウンされた後、開発者には孤立したワークスペースと壊れた依存関係が残されました。
vibe-kanban-alternative は、 BloopAI/vibe-kanban および dexloom/vibe-kanban-indie が積極的に維持され、独立して進化したものです。のために特別に構築

ingle-developer ワークフローでは、クラウド アカウント、チーム認証、リモート テレメトリは必要ありません。これをすべて自分のマシン上で実行し、ブラウザ、端末 (TUI)、または電話 (Telegram) を介してコーディング エージェントのフリートを調整します。
⚡ このフォークの何が違うのですか?
能力
上流 (ブループ)
インディーズベースライン
⚡ バイブカンバン代替 (このフォーク)
サーバーインフラストラクチャ
🛑 日没 / 死者
🟢 ローカル/セルフホスト
🟢 100% オフラインおよびローカル ランタイム
クロスセッションメモリ
❌ エフェメラル (カードごとに失われます)
❌ なし
🧠 ネイティブ mem0 + Qdrant + NetworkX グラフ メモリ
プロンプトキャッシュヒットアーキテクチャ
❌ なし
❌ なし
⚡ 決定的でキャッシュに優しいメモリプレフィックスインジェクション
テレメトリーと可観測性
❌ なし
⚠️最小限
📊 フル設定 → 使用状況ダッシュボード (トークン、ヒートマップ、エージェント)
コーディングエージェントのマトリックス
レガシー CLI サブセット
スタンダードセット
🚀 10 以上のエージェント: クロード コード、反重力 (AGY)、コーデックス、ジェミニ CLI など。
反重力 (AGY) エージェント
❌ なし
⚠️ 基本テキストモード
💎 完全なストリーム JSON 解析、ToolUse ビジュアル カード、および推論労力の制御
チャットの入力と履歴
⚠️ 基本的なテキストエリア
標準 WYSIWYG
⌨️ ターミナルスタイルの ArrowUp/Down プロンプト履歴 + Enter/Ctrl+Enter モード
セルフホスト型 Git リモート
GitHubのみ
GitHub + 基本的な Gitea
🐙 自動ルーティング Gitea / Forgejo REST API + GitHub (gh)
リモコンとコックピット
Web UIのみ
TUI + 電報
📱 ターミナル TUI コックピット + 送信専用電報フォーラム ブリッジ
バックアップと災害復旧
❌ なし
基本
💾 フル DB、設定、mem0 ベクトル/グラフ状態のエクスポート/インポート
チャットとUIストリーミング
⚠️ 長い差分でのレイテンシのバグ
標準
⚡ 最適化されたキャンバス/チャットレンダリングとワークツリーパネルの修正
🚀 概要
ソフトウェア エンジニアリングは、コーディング エージェントに指示すること、つまり作業を計画し、それを実装するためのモデルを生成し、差分を確認し、出荷することを意味するようになってきています。 b

ottleneck はコードを入力しなくなりました。多くのエージェント セッションを調整、レビューし、一貫性を保ちます。バイブカンバンオルタナティブは、そのプロセスを高速かつローカルかつ個人的に行うように構築されています。つまり、単一の開発者が完全に自分のマシン上で、チームもクラウドもアカウントも必要ありません。
その核となるのは、エージェントの作業を計画および追跡するカンバン ボードに加え、各カードを実際のブランチ、ターミナル、開発サーバーに変えるワークスペース ランタイムであり、10 以上のコーディング エージェント (Claude Code、OpenCode、Qwen Code、Codex、Gemini CLI、Antigravity、Copilot、Amp、Cursor、Droid、CCR) のいずれかが計画を実行します。
カンバン問題を考慮した計画 — ボード、列、優先順位、タグ、サブ問題、パイプライン。カードは作品の唯一の真実の情報源です。
ワークスペースでコーディング エージェントを実行します。各カードはワークスペース (ブランチ、ターミナル、開発サーバー、および構成可能なパイプライン (クイック、ベーシック、非同期バリアント) に従うエージェント) を起動します。
差分を確認して反復します。インライン コメント、差分、プレビュー ブラウザー、およびマージや PR の前に結果を承認できるように、エージェントを一時停止してアラームを生成する手動レビュー ステージがあります。
10 個以上のコーディング エージェントを切り替えて、Claude Code、OpenCode、Qwen Code、Codex、Gemini CLI、Antigravity ( agy )、Copilot、Amp、Cursor、Droid、CCR を 1 つのボードから操作します。
クロスセッション プロジェクト メモリ (mem0) - エージェントは、再起動後も存続するグラフ メモリを使用して、作業しているリポジトリに関する検証済みの事実を呼び出して保持し、リポジトリごとにキー設定され、CLI 全体で共有されます。
使用状況と可観測性 — 1 日あたりのアクティビティ、エージェントごとの実行バー、抽出トークンの監視、およびプロジェクトの進行状況を含む [設定] → [使用状況] ダッシュボード。
ワークスペース、PR、マージ — 作業を既存のセッションにディスパッチし、AI が生成した説明を含む PR (GitHub または Gitea/Forgejo) を開き、ベースにスカッシュマージします。
ターミナル ＆

電話 — TUI コックピットと Telegram エスカレーションにより、ブラウザがなくても制御できます。
すべてのコーディング エージェントのクロスセッション メモリ。 vibe-kanban-alternative にはファーストクラスの mem0 統合が付属しているため、ワークスペースを駆動するエージェントは、作業するリポジトリの耐久性のあるセマンティック メモリを共有します。
⚬ 起動時の自動呼び出し : ワークスペースは保存されたメモリに現在のリポジトリ スラグをクエリし、それをエージェント プロンプトの前に追加します。苦労して勝ち取ったアーキテクチャ上の規約やデバッグでの発見は決して忘れられません。
⚬ リポジトリの共有知識 : メモリはリポジトリごとにキー化されます。 Claude Code でタスクを開始し、プロジェクトの途中でコンテキストを失うことなく OpenCode または Antigravity に切り替えることができます。
⚬ 検証された事実のセーブバック: メモリ パイプライン ステージは、エージェントに、memory_save を介して自己完結型の検証された事実 (アーキテクチャ上の決定、パターン、根本原因) のみを保持するように指示します。一時的なおしゃべりはフィルタリングされて除去されます。
⚬ MCP ツールの統合 : エージェントは、最上級のモデル コンテキスト プロトコル (MCP) ツールとして、memory_search 、memory_recall 、memory_save にアクセスできます。
⚬ グラフベース メモリ (GraphML) : エンティティとリレーションの抽出により、相互接続されたグラフ (mem0 + Qdrant + NetworkX) が作成され、ディスク ( /data/graphs/*.graphml ) 上に永続化されるため、知識構造はコンテナーの再起動後も存続します。
プロンプト キャッシュ (Anthropic、OpenRouter、DeepSeek) を使用してプロバイダーのトークン コストを最小限に抑えるには、次の手順を実行します。
静的プレフィックス配置: メモリ ブロックは、動的なユーザー命令の前に静的タスク プレフィックスに挿入されます。
決定的順序付け : メモリ エントリは決定的にソートされ、継続的なキャッシュ ヒットのためにセッション全体でバイト同一のプロンプト プレフィックスが保証されます。
セッション中の混乱なし:memory_search とmemory_save は、システム プロンプトの変更ではなく MCP ツール呼び出しとして実行され、トークン cac が維持されます。

彼は暖かいです。
プロジェクト メモリ レイヤーは、ローカル Docker スタック ( mem0-vk )、つまり :8000 上の mem0 API サーバー、Qdrant ベクター ストア、および Python 埋め込み + NetworkX グラフ サービス上で実行されます。
cd mem0-vk
cp .env.example .env # 次に、抽出 LLM キーを設定します (以下を参照)
docker 構成 -d --build
アプリから構成: [設定] → [メモリ] を開き、実行時にグラフを管理し、抽出プロバイダー (Groq、OpenRouter、Local llama) を構成し、トークンの使用状況を表示します。
vibe-kanban-alternative は、10 以上の主要なコーディング エージェントとネイティブに統合します。
Google Antigravity ( agy ) ⚡ (新規) :
ストリーム JSON プロトコルの完全なサポート。
ファイル検査 ( view_file )、検索 ( grep_search 、 find_by_name )、bash コマンド ( run_command )、およびファイル編集 ( write_to_file 、 replace_file_content ) 用のネイティブ ビジュアル カード。
gemini-3.7-flash の自動フォールバックによる推論エフォート制御 (Low、Medium、High)。
YOLO モードの自動許可バイパス ( --dangerously-skip-permissions )。
Anthropic Claude Code: ヘッド付きおよびヘッドレス モード、完全な MCP ツール承認、およびターン ナビゲーション。
OpenCode および OpenCode Headed : ローカルおよびリモート推論を備えたマルチモデル エージェント ランナー。
OpenAI Codex : 深い推論と計画の生成。
Qwen コード: 高性能のローカルおよびクラウド エージェント ワークフロー。
Google Gemini CLI : ネイティブ Gemini 実行。
GitHub Copilot CLI、Cursor Agent、Droid、および Amp 。
⌨️ チャットと端末の操作
プロンプト履歴

[切り捨てられた]

## Original Extract

vibe-kanban-indie fork. Contribute to flashlan/vibe-kanban-alternative development by creating an account on GitHub.

GitHub - flashlan/vibe-kanban-alternative: vibe-kanban-indie fork · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
flashlan
/
vibe-kanban-alternative
Public
forked from dexloom/vibe-kanban-indie
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2,456 Commits 2,456 Commits Folders and files
.cargo .cargo .github .github apps/ macos apps/ macos assets assets automation automation crates crates dev_assets_seed dev_assets_seed docs docs npx-cli npx-cli packages packages patches patches scripts scripts shared shared .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .npmrc .npmrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE-OF-CONDUCT.md CODE-OF-CONDUCT.md CONTEXT.md CONTEXT.md CONTRIBUTORS.md CONTRIBUTORS.md Caddyfile.example Caddyfile.example Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile PUBLISHING.md PUBLISHING.md README.md README.md SPEC-conversation-list.md SPEC-conversation-list.md SPEC.md SPEC.md TODO.md TODO.md gitea.toml.example gitea.toml.example local-build.sh local-build.sh mobile-testing.md mobile-testing.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml projects.toml.example projects.toml.example restart.sh restart.sh rust-toolchain.toml rust-toolchain.toml rustfmt.toml rustfmt.toml View all files Repository files navigation
The resilient, 100% self-hosted cockpit for single-developer AI orchestration.
Persistent Graph Memory (mem0) • 10+ Coding Agents (inc. Antigravity/AGY) • Gitea/Forgejo PRs • TUI Cockpit • Telegram Daemon • Usage Observability
Quick Start •
Overview •
What's Different •
Project Memory (mem0) •
Supported Agents •
Chat & Controls •
Usage Dashboard •
TUI Cockpit •
Telegram Bridge •
Development
Launch the entire cockpit with a single command — no install, no accounts, no cloud setup:
npx vibe-kanban-alternative
💡 Zero Setup Required : Downloads prebuilt binaries and launches the local web cockpit directly at http://localhost:3001 (with backend on :3002 ).
⚙️ Full Setup & Environment Configuration
For development, custom ports, local mem0 vector stores, and custom agent configurations:
git clone https://github.com/flashlan/vibe-kanban-alternative.git
cd vibe-kanban-alternative
pnpm i
2. Configure Environment Variables
Duplicate the template configuration file:
cp .env.example .env
Open .env in your editor and configure your environment:
# ==========================================
# 🌐 CORE SERVER CONFIGURATION
# ==========================================
PORT = 3000
HOST = localhost
NODE_ENV = development
# ==========================================
# 🧠 MEM0 LONG-TERM MEMORY SETTINGS
# ==========================================
MEM0_ENABLED = true
MEM0_API_KEY = your_mem0_api_key_here
# If running local embeddings/vector store:
# MEM0_VECTOR_STORE=qdrant
# MEM0_HOST=http://localhost:6333
# ==========================================
# 🤖 AGENT API KEYS & RUNTIMES
# ==========================================
ANTHROPIC_API_KEY = sk-ant-...
OPENAI_API_KEY = sk-...
GEMINI_API_KEY = ...
# Antigravity (AGY) Specific Settings
AGY_CLI_PATH = /usr/local/bin/agy
AGY_DEFAULT_TEMPERATURE = 0.2
# ==========================================
# 📊 METRICS & TELEMETRY
# ==========================================
ENABLE_METRICS_DASHBOARD = true
METRICS_STORAGE_PATH = ./data/metrics.sqlite
# ==========================================
# 💾 BACKUP CONFIGURATION
# ==========================================
BACKUP_ENABLED = true
BACKUP_INTERVAL_MINUTES = 30
BACKUP_STORAGE_PATH = ./backups
3. Start Development Cockpit
./restart.sh
📌 Context: Reviving & Evolving Vibe-Kanban
Following the official cloud shutdown of Bloop's hosted servers , developers were left with orphaned workspaces and broken dependencies.
vibe-kanban-alternative is an actively maintained, independent evolution of BloopAI/vibe-kanban and dexloom/vibe-kanban-indie . Built specifically for a single-developer workflow , it requires no cloud accounts, no team auth, and zero remote telemetry . You run it entirely on your machine, orchestrating a fleet of coding agents via browser, terminal (TUI), or phone (Telegram).
⚡ What's Different in This Fork?
Capability
Upstream (Bloop)
Indie Baseline
⚡ vibe-kanban-alternative (This Fork)
Server Infrastructure
🛑 Sunsetting / Dead
🟢 Local / Self-hosted
🟢 100% Offline & Local Runtime
Cross-Session Memory
❌ Ephemeral (Lost per card)
❌ None
🧠 Native mem0 + Qdrant + NetworkX Graph Memory
Prompt Cache-Hit Architecture
❌ None
❌ None
⚡ Deterministic, Cache-Friendly Memory Prefix Injection
Telemetry & Observability
❌ None
⚠️ Minimal
📊 Full Settings → Usage Dashboard (Tokens, Heatmaps, Agents)
Coding Agents Matrix
Legacy CLI subset
Standard Set
🚀 10+ Agents: Claude Code, Antigravity (AGY), Codex, Gemini CLI, etc.
Antigravity (AGY) Agent
❌ None
⚠️ Basic text mode
💎 Full stream-json parsing, ToolUse visual cards & reasoning effort control
Chat Input & History
⚠️ Basic textarea
Standard WYSIWYG
⌨️ Terminal-style ArrowUp/Down prompt history + Enter/Ctrl+Enter modes
Self-Hosted Git Remotes
GitHub Only
GitHub + Basic Gitea
🐙 Auto-routing Gitea / Forgejo REST API + GitHub ( gh )
Remote Control & Cockpit
Web UI only
TUI + Telegram
📱 Terminal TUI Cockpit + Send-Only Telegram Forum Bridge
Backup & Disaster Recovery
❌ None
Basic
💾 Full DB, Settings & mem0 Vector/Graph State Export/Import
Chat & UI Streaming
⚠️ Latency bugs on long diffs
Standard
⚡ Optimized Canvas/Chat rendering & worktree panel fixes
🚀 Overview
Software engineering increasingly means directing coding agents — planning work, spawning a model to implement it, reviewing its diff, and shipping. The bottleneck is no longer typing code; it's orchestrating, reviewing, and keeping many agent sessions coherent. vibe-kanban-alternative is built to make that process fast, local, and personal : a single developer, entirely on their own machine, no team, no cloud, no account.
At its core it's a kanban board that plans and tracks agent work , plus a workspace runtime that turns each card into a real branch, terminal, and dev-server where any of 10+ coding agents (Claude Code, OpenCode, Qwen Code, Codex, Gemini CLI, Antigravity, Copilot, Amp, Cursor, Droid, CCR) executes the plan:
Plan with kanban issues — boards, columns, priorities, tags, sub-issues, pipelines; cards are the single source of truth for a piece of work.
Run coding agents in workspaces — each card launches a workspace: a branch, a terminal, a dev server, and an agent following a configurable pipeline (Quick, Basic, async variants).
Review diffs and iterate — inline comments, diffs, preview browser, and the manual-review stage that pauses the agent and raises an alarm so you approve the result before any merge or PR.
Switch between 10+ coding agents — drive Claude Code, OpenCode, Qwen Code, Codex, Gemini CLI, Antigravity ( agy ), Copilot, Amp, Cursor, Droid, and CCR from one board.
Cross-session project memory (mem0) — agents recall and persist verified facts about the repositories they work in, keyed per repository and shared across CLIs, with a graph memory that survives restarts.
Usage & observability — a Settings → Usage dashboard with per-day activity, per-agent execution bars, extraction-token monitoring, and project progress.
Workspaces, PRs, and merge — dispatch work to existing sessions, open PRs (GitHub or Gitea/Forgejo) with AI-generated descriptions, squash-merge to base.
Terminal & phone — a TUI cockpit and Telegram escalation keep you in control without the browser.
Cross-session memory for every coding agent. vibe-kanban-alternative ships with a first-class mem0 integration so the agents that drive your workspaces share a durable, semantic memory of the repositories they work in.
⚬ Automatic Recall on Launch : Workspaces query stored memory for the current repository slug and prepend it to the agent prompt. Hard-won architectural conventions and debugging discoveries are never forgotten.
⚬ Shared Repository Knowledge : Memory is keyed per repository. You can start a task on Claude Code and switch to OpenCode or Antigravity mid-project without losing context.
⚬ Verified Fact Save-Back : The memory pipeline stage instructs the agent to persist only self-contained, verified facts (architectural decisions, patterns, root causes) via memory_save . Ephemeral chatter is filtered out.
⚬ MCP Tools Integration : Agents have access to memory_search , memory_recall , and memory_save as first-class Model Context Protocol (MCP) tools.
⚬ Graph-Based Memory (GraphML) : Entity and relation extraction creates an interconnected graph (mem0 + Qdrant + NetworkX), persisted on disk ( /data/graphs/*.graphml ) so knowledge structures survive container reboots.
To minimize token costs on providers with prompt caching (Anthropic, OpenRouter, DeepSeek):
Static Prefix Placement : The memory block is injected into the static task prefix before dynamic user instructions.
Deterministic Ordering : Memory entries are sorted deterministically, ensuring byte-identical prompt prefixes across sessions for continuous cache hits.
No Mid-Session Perturbation : memory_search and memory_save execute as MCP tool calls rather than system prompt mutations, keeping the token cache warm.
The project memory layer runs on a local Docker stack ( mem0-vk ): a mem0 API server on :8000 , a Qdrant vector store, and a Python embeddings + NetworkX graph service.
cd mem0-vk
cp .env.example .env # then set an extraction LLM key (see below)
docker compose up -d --build
Configure from the app: Open Settings → Memory to manage the graph at runtime, configure extraction providers (Groq, OpenRouter, Local llama), and view token usage.
vibe-kanban-alternative integrates natively with 10+ leading coding agents:
Google Antigravity ( agy ) ⚡ (New) :
Full stream-JSON protocol support.
Native visual cards for file inspection ( view_file ), search ( grep_search , find_by_name ), bash commands ( run_command ), and file edits ( write_to_file , replace_file_content ).
Reasoning effort controls ( Low , Medium , High ) with automatic fallback for gemini-3.7-flash .
YOLO mode auto-permission bypass ( --dangerously-skip-permissions ).
Anthropic Claude Code : Headed & headless modes, full MCP tool approvals, and turn navigation.
OpenCode & OpenCode Headed : Multi-model agent runner with local & remote inference.
OpenAI Codex : Deep reasoning & plan generation.
Qwen Code : High-performance local and cloud agent workflows.
Google Gemini CLI : Native Gemini execution.
GitHub Copilot CLI , Cursor Agent , Droid , and Amp .
⌨️ Chat & Terminal Interaction
Prompt Hist

[truncated]
