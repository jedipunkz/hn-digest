---
source: "https://github.com/akshaylakkur/Q"
hn_url: "https://news.ycombinator.com/item?id=48562253"
title: "Qode – The first AI agent that can generate 50k line codebases in one prompt"
article_title: "GitHub - akshaylakkur/Q · GitHub"
author: "akshayl284"
captured_at: "2026-06-16T21:53:21Z"
capture_tool: "hn-digest"
hn_id: 48562253
score: 1
comments: 1
posted_at: "2026-06-16T21:17:36Z"
tags:
  - hacker-news
  - translated
---

# Qode – The first AI agent that can generate 50k line codebases in one prompt

- HN: [48562253](https://news.ycombinator.com/item?id=48562253)
- Source: [github.com](https://github.com/akshaylakkur/Q)
- Score: 1
- Comments: 1
- Posted: 2026-06-16T21:17:36Z

## Translation

タイトル: Qode – 1 つのプロンプトで 50,000 行のコードベースを生成できる初の AI エージェント
記事タイトル: GitHub - akshaylakkur/Q · GitHub
説明: GitHub でアカウントを作成して、akshaylakkur/Q の開発に貢献します。

記事本文:
GitHub - アクシャイラックル/Q · GitHub
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
アクシャイラックル
/
Q
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
25℃

省略 25 コミット .github/ workflows .github/ workflows apps/ q-cli apps/ q-cli npm npm パッケージ パッケージ スクリプト スクリプト .gitignore .gitignore .npmrc .npmrc Q.md Q.md README.md README.md Bands.sh Bands.sh install.sh install.sh package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Qode は、長時間実行されるタスクや大規模なコードベース開発のための、ターミナルベースの自律コーディング エージェントです。豊富な TUI を使用して端末内で直接動作し、複数の LLM プロバイダーに接続し、1 行の修正から数百のファイルにまたがる横断的な変換まで拡張するインテリジェントなパイプラインを通じて複数のファイルの変更を調整します。
Node SDK は、同じ Qode CLI 設定、ツール、スキル、MCP サーバーを再利用するシン クライアントです。リアルタイムで応答をストリーミングし、承認とツール呼び出しを表示し、プログラムでセッションを調整できるようにします。
Qode は自律エージェント ランタイムを提供し、次のことを可能にします。
カスタム アプリケーションを構築する — Qode を独自のツールやワークフローに統合する
複雑なタスクを自動化する — 長時間実行されるセッションにわたる複数ターンの会話をスクリプト化する
大規模なコードベース操作を実行する — Modus Maximus を使用して横断的な変更を行う
2 つの実行モード: AUTO (分類子主導の動作を備えた単一エージェントのターン ループ) と MODUS MAXIMUS (4 フェーズのパイプライン: 計画の生成、ユーザーの確認、スペシャリスト プロファイルによる順次サブエージェントの実行、および最終的なサマリー)
意図の分類 : 範囲、深さ、ファイル参照、アクション動詞、並列処理のニーズ、および検証要件に関するすべてのプロンプトのヒューリスティック分析
動的エスカレーション: 現在の場合、DynamicReclassifier を介したランタイム モード エスカレーション

実行戦略が不十分
Modus Maximus : 15 ～ 50 の依存関係を意識した計画ステップ、ユーザー確認 (良さそうだ / 修正が必要 / やり直し)、ヒューリスティックなプロファイル解決による連続したサブエージェントの実行、および最終的なサマリー (特殊なアーキテクチャを通じて、新しいプロジェクトの 50k ～ 80k 行のコードをワンショットで生成することを目的としています)
スペシャリスト エージェント プロファイル: Editius (StrReplace による外科コード編集)、Rewritius (フルファイルの書き換えとリファクタリング)、Searchius (コードベース分析)、Auto (タスク適応型)。 /agent経由で切り替え可能
4 層メモリ: 優先度タグ付き圧縮を備えた WorkingMemory、TF-IDF スコアリングを備えた EpisodicRecall、LTPM (保持ポリシーを備えたディスクバックアップ永続性)、SemanticRecall (HNSW インデックスを介したベクトルベースの ANN 検索)、CodebaseGraph (TS、JS、Python、Rust、Go、Java の言語対応モデル)
検証パイプライン: 言語ごとの自動検出と SHA-256 キャッシュを備えた 7 つのゲート (構文、lint、タイプチェック、単体テスト、統合テスト、アーキテクチャ、フル スイート)
自己修正: アーキテクチャ エスカレーションによる自動修正と再検証のループ
複数の LLM プロバイダー: Anthropic、OpenAI、Google Gemini、Ollama (ローカル、API キーなし)、Kimi、OpenAI 互換
プラグイン システム: ライフサイクル フックを使用したマニフェスト ベースの検出
スキル システム: 再利用可能なプロンプト テンプレートとインライン スクリプト
MCP クライアント: stdio および HTTP/SSE トランスポート
セッションの永続性: JSONL ワイヤー形式、BLOB ストア、移行システム
非対話型モード: CI/CD の場合は --prompt、--output-format json
オンボーディング ウィザード: プロバイダーとモデルの初回実行セットアップ
パッケージ
説明
ステータス
アプリ/q-cli
メイン CLI - オーケストレーター、TUI、メモリ、MCP、プラグイン、検証
利用可能
パッケージ/エージェントコア
コア エージェント ランタイム - コンテキスト、ターン ループ、ツール、サブエージェント、プロファイル
利用可能
パッケージ/qprovs
LLM プロバイダーの抽象化
アヴァイ

ラベルの付いた
パッケージ/qmain
実行環境 — ファイル操作、シェル、git、Web
利用可能
パッケージ/ノード SDK
プログラムによる Node.js SDK
利用可能
パッケージ/oauth
MCP サーバー認証用の OAuth 2.0
利用可能
パッケージ/テレメトリ
オプションのクラッシュレポートとテレメトリ
利用可能
クイックスタート
npm install -g qode-agent
Node.js >= 22.19.0 および pnpm >= 10.33.0 が必要です。インストーラーはプロジェクトをビルドし、 ~/.Q/ を作成して、q-cli ラッパーをインストールします。
# インタラクティブモード (TUI を開きます)
q-cli
# 速記
q
# CI/CD のワンショット プロンプト
q-cli -p " エラー処理を src/routes/users.ts に追加します "
# 複雑なタスクには Modus Maximus を使用する
q-cli
q > /mode modus-maximus
q > OAuth 2.0 を使用するように認証システムをリファクタリングします。
# 前のセッションを再開する
q-cli -S <セッションID>
# JSON 出力による非対話型
q-cli -p " 型エラーを修正 " --output-format json
CLI オプション
-S, --session <id> セッションを再開します
-C, -- continue 最後のセッションを続行します
-y、--yolo すべてのアクションを自動承認します
-m, --model <name> LLM モデルをオーバーライドします
-p, --prompt <テキスト> 非対話型モード
--plan 起動時の計画モード
--auto 自動許可モード
--setup セットアップ ウィザードを再実行します
--output-format <fmt> テキスト |ジェソン |ストリームjson
--skills-dir <dir> 追加のスキル ディレクトリ
--cwd <パス> 作業ディレクトリ
--tui / --no-tui TUI を強制/有効化
コマンド
init Qode プロジェクトを初期化する
session セッションの管理 (リスト、表示、削除、エクスポート、インポート)
config 構成の表示と編集
ドクター 構成の問題を診断して修正する
移行 バージョン間でデータを移行する
update アップデートを確認してインストールする
補完 シェル補完スクリプトを生成する
デーモン Qode デーモンの起動/停止
connect リモートインスタンスに接続します
プロファイル エージェントプロファイルの管理
プラグイン プラグインの管理
スラッシュコマンド
/help インタラクティブ ヘルプ ダッシュボード
/status セッション ステータス ダッシュボード
/session の表示/管理

セッション情報
/clear トランスクリプトをクリアします
/exit 正常に終了します
/version バージョン情報を表示する
/mode <name> モードの切り替え (auto、modus-maximus)
/agent <名前> プロファイルの切り替え (auto、editius、rewritius、searchius)
/qmd プロジェクト規約を使用して Q.md を生成します
環境変数
Q_PROVIDER LLM プロバイダー (anthropic、openai、ollam など)
Q_MODEL モデル名 (例: claude-sonnet-4-20250514)
Q_API_KEY プロバイダーの API キー
Q_BASE_URL カスタムベース URL
Q_THINKING 思考レベル（なし、低、中、高）
構成
昇順の優先順位でロードされます: 組み込みデフォルト → ~/.Q/config.toml → .q/config.toml (cwd からウォークアップ) → 環境変数。最初の実行では、対話型のオンボーディング ウィザードが起動します。
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

Contribute to akshaylakkur/Q development by creating an account on GitHub.

GitHub - akshaylakkur/Q · GitHub
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
akshaylakkur
/
Q
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
25 Commits 25 Commits .github/ workflows .github/ workflows apps/ q-cli apps/ q-cli npm npm packages packages scripts scripts .gitignore .gitignore .npmrc .npmrc Q.md Q.md README.md README.md bands.sh bands.sh install.sh install.sh package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Qode is a terminal-based autonomous coding agent for long-running tasks and massive codebase development. It operates directly in your terminal with a rich TUI, connects to multiple LLM providers, and orchestrates multi-file changes through an intelligent pipeline that scales from one-line fixes to cross-cutting transformations spanning hundreds of files.
The Node SDK is a thin client that reuses the same Qode CLI configuration, tools, skills, and MCP servers. It streams responses in real time, surfaces approvals and tool calls, and lets you orchestrate sessions programmatically.
Qode provides an autonomous agent runtime, enabling you to:
Build custom applications — Integrate Qode into your own tools and workflows
Automate complex tasks — Script multi-turn conversations across long-running sessions
Run massive codebase operations — Use Modus Maximus for cross-cutting changes
Two execution modes : AUTO (single-agent turn loop with classifier-driven behavior) and MODUS MAXIMUS (4-phase pipeline: plan generation, user confirmation, sequential sub-agent execution through specialist profiles, and final summary)
Intent classification : Heuristic analysis of every prompt for scope, depth, file references, action verbs, parallelism needs, and verification requirements
Dynamic escalation : Runtime mode escalation via DynamicReclassifier when the current execution strategy is insufficient
Modus Maximus : 15-50 dependency-aware plan steps, user confirmation (Looks Good / Needs Revision / Redo), sequential sub-agent execution with heuristic profile resolution, and final summary (Aims to generate 50k-80k lines of code for a fresh project in one shot through our specialized architecture)
Specialist agent profiles : Editius (surgical code editing via StrReplace), Rewritius (full-file rewrites and refactoring), Searchius (codebase analysis), Auto (task-adaptive); switchable via /agent
Four-tier memory : WorkingMemory with priority-tagged compaction, EpisodicRecall with TF-IDF scoring, LTPM (disk-backed persistence with retention policies), SemanticRecall (vector-based ANN search via HNSW index), CodebaseGraph (language-aware model for TS, JS, Python, Rust, Go, Java)
Verification pipeline : 7 gates (syntax, lint, typecheck, unit tests, integration tests, architecture, full suite) with per-language auto-detection and SHA-256 caching
Self-correction : Automated fix-and-reverify loop with architecture escalation
Multiple LLM providers : Anthropic, OpenAI, Google Gemini, Ollama (local, no API key), Kimi, OpenAI-compatible
Plugin system : Manifest-based discovery with lifecycle hooks
Skill system : Reusable prompt templates and inline scripts
MCP client : stdio and HTTP/SSE transport
Session persistence : JSONL wire format, blob store, migration system
Non-interactive mode : --prompt for CI/CD, --output-format json
Onboarding wizard : First-run setup for provider and model
Package
Description
Status
apps/q-cli
Main CLI — orchestrator, TUI, memory, MCP, plugins, verification
Available
packages/agent-core
Core agent runtime — context, turn loop, tools, sub-agents, profiles
Available
packages/qprovs
LLM provider abstraction
Available
packages/qmain
Execution environment — file ops, shell, git, web
Available
packages/node-sdk
Programmatic Node.js SDK
Available
packages/oauth
OAuth 2.0 for MCP server auth
Available
packages/telemetry
Optional crash reporting and telemetry
Available
Quick Start
npm install -g qode-agent
Requires Node.js >= 22.19.0 and pnpm >= 10.33.0. The installer builds the project, creates ~/.Q/ , and installs the q-cli wrapper.
# Interactive mode (opens TUI)
q-cli
# Shorthand
q
# One-shot prompt for CI/CD
q-cli -p " Add error handling to src/routes/users.ts "
# Use Modus Maximus for complex tasks
q-cli
q > /mode modus-maximus
q > Refactor the authentication system to use OAuth 2.0
# Resume a previous session
q-cli -S < session-id >
# Non-interactive with JSON output
q-cli -p " Fix the type errors " --output-format json
CLI Options
-S, --session <id> Resume a session
-C, --continue Continue last session
-y, --yolo Auto-approve all actions
-m, --model <name> Override LLM model
-p, --prompt <text> Non-interactive mode
--plan Plan mode on startup
--auto Auto permission mode
--setup Re-run setup wizard
--output-format <fmt> text | json | stream-json
--skills-dir <dir> Additional skill directories
--cwd <path> Working directory
--tui / --no-tui Force/enable TUI
Commands
init Initialize a Qode project
session Manage sessions (list, show, delete, export, import)
config View and edit configuration
doctor Diagnose and fix configuration issues
migrate Migrate data between versions
update Check for and install updates
completions Generate shell completion scripts
daemon Start/stop the Qode daemon
connect Connect to a remote instance
profile Manage agent profiles
plugin Manage plugins
Slash Commands
/help Interactive help dashboard
/status Session status dashboard
/session Show/manage session info
/clear Clear transcript
/exit Gracefully exit
/version Show version info
/mode <name> Switch mode (auto, modus-maximus)
/agent <name> Switch profile (auto, editius, rewritius, searchius)
/qmd Generate Q.md with project conventions
Environment Variables
Q_PROVIDER LLM provider (anthropic, openai, ollama, etc.)
Q_MODEL Model name (e.g. claude-sonnet-4-20250514)
Q_API_KEY API key for the provider
Q_BASE_URL Custom base URL
Q_THINKING Thinking level (none, low, medium, high)
Configuration
Loaded in ascending priority: built-in defaults → ~/.Q/config.toml → .q/config.toml (walked up from cwd) → environment variables. First run launches interactive onboarding wizard.
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
