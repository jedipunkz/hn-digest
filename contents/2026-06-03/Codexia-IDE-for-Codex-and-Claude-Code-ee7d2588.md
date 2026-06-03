---
source: "https://github.com/milisp/codexia/"
hn_url: "https://news.ycombinator.com/item?id=48373917"
title: "Codexia: \"IDE\" for Codex and Claude Code"
article_title: "GitHub - milisp/codexia: Agent Workstation for Codex CLI + Claude Code — with task scheduler, git worktree & remote control, skills management · GitHub"
author: "adamnemecek"
captured_at: "2026-06-03T00:38:51Z"
capture_tool: "hn-digest"
hn_id: 48373917
score: 1
comments: 0
posted_at: "2026-06-02T18:10:31Z"
tags:
  - hacker-news
  - translated
---

# Codexia: "IDE" for Codex and Claude Code

- HN: [48373917](https://news.ycombinator.com/item?id=48373917)
- Source: [github.com](https://github.com/milisp/codexia/)
- Score: 1
- Comments: 0
- Posted: 2026-06-02T18:10:31Z

## Translation

タイトル: Codexia: Codex および Claude Code 用の「IDE」
記事のタイトル: GitHub - milisp/codexia: Codex CLI + Claude Code 用エージェント ワークステーション — タスク スケジューラ、git worktree とリモート コントロール、スキル管理を備えた · GitHub
説明: Codex CLI + Claude Code 用のエージェント ワークステーション — タスク スケジューラ、git worktree およびリモート コントロール、スキル管理を備えた - milisp/codexia

記事本文:
GitHub - milisp/codexia: Codex CLI + Claude Code 用のエージェント ワークステーション — タスク スケジューラ、git worktree とリモート コントロール、スキル管理を備えた · GitHub
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
ミリスプ
/
コデクシア
公共
ああ、ああ！
読み込み中にエラーが発生しました。このPAをリロードしてください

げ。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1,279 コミット 1,279 コミット .github .github .vscode .vscode bots bots docs docs public public scripts scripts src-tauri src-tauri src src .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COMMERCIAL.md COMMERCIAL.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md biome.json biome.json bun.lock bun.lockComponents.jsonComponents.json Index.html Index.html justfile justfile package.json package.json tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Codexia は、Codex CLI + Claude Code 用の Tauri v2 アプリです。エージェント ワークフロー、IDE のようなエディター、ヘッドレス Web サーバー、およびプロンプト メモ帳を 1 つのワークスペースに組み合わせています。
エージェントのワークフロー: 定期的なジョブのタスク スケジューラ、ヘッドレス Web サーバー経由のリモート コントロール
ワークスペース : Git ワークツリー管理、プロジェクト ファイル ツリー、IDE 風のエディタ、プロンプト メモ帳、ローカル Web プレビュー
データツール: ワンクリックで PDF / XLSX / CSV プレビュー
エコシステム : MCP サーバー マーケットプレイス、エージェント スキル マーケットプレイス
パーソナライゼーション: テーマとアクセントのカスタマイズ、使用状況分析ダッシュボード
ブリュータップミリスプ/コデクシア
brew install --cask codexia
ビルド済みリリース (macOS / Linux / Windows)
プロンプトを入力し、エージェント セッションを開始します。
定期的なワークフロー用のエージェント タスク スケジューラ ジョブを作成します。
クロードエージェントのRust SDK統合
フロントエンド: React + TypeScript + Zustand + src/ の shadcn/ui
デスクトップ バックエンド: Tauri v2 + Rust (src-tauri/src/)
ヘッドレス バックエンド: リモート制御用の Axum Web サーバー

rc-tauri/src/web/
エージェント ランタイム: セッション/ターン ライフサイクルのための Codex アプリサーバー JSON-RPC 統合
リアルタイム更新: ブラウザ クライアントの /ws での WebSocket ブロードキャスト ストリーム
src-tauri/src/lib.rs (デスクトップコマンドと状態)
src-tauri/src/web/server.rs (ヘッドレスサーバーの起動)
src-tauri/src/web/router.rs (HTTP API ルート サーフェス)
src/services/tauri/ (フロントエンド呼び出しレイヤー)
Codexia は、Web/ヘッドレス モードで実行するときにブラウザからアクセス可能な API を公開します。
ヘルスとストリーム: GET /health 、 GET /ws
コーデックスのライフサイクル: /api/codex/thread/* 、 /api/codex/turn/* 、 /api/codex/model/* 、 /api/codex/approval/*
自動化スケジューラー: /api/automation/* (作成/更新/リスト/実行/一時停止/削除)
ファイル、git、ターミナル: /api/filesystem/* 、 /api/git/* 、 /api/terminal/*
メモと生産性: /api/notes/* 、 /api/codex/usage/token
新しい API ハンドラーを src-tauri/src/web/handlers/ に追加します。
src-tauri/src/web/router.rs にルートを登録します。
対応するフロントエンド クライアント呼び出しを src/services/tauri/ に追加します。
プロセスの分離 : エージェントは別のプロセスで実行されます。
権限制御: エージェントごとにファイルとネットワークのアクセスを設定します
ローカル ストレージ : すべてのデータはマシン上に残ります
オープンソース : オープンソース コードによる完全な透明性
テレメトリ: オプトインのみ、デフォルトではオフ
貢献は大歓迎です。セットアップとワークフローについては CONTRIBUTING.md をお読みください。
jeremiahodom/codex-ui — API/SSE を使用した Node.js バックエンド
Itexoft/codexia — SSH 統合
nuno5645/codexia — 推論とトークンカウントイベント
awesome-codex-cli — Codex CLI リソースの厳選されたリスト
claw-army/claude-node — クロード コード CLI の Python サブプロセス ブリッジ
AGPL-3.0 (オープンソース) と商用ライセンス (クローズドソース/SaaS 使用) に基づくデュアルライセンス。
詳細については、COMMERCIAL.md を参照してください。
Codex CLI + Claude Code 用エージェント ワークステーション — タスク スケジューラ、git worktree、リモート c を搭載

コントロール、スキル管理
読み込み中にエラーが発生しました。このページをリロードしてください。
71
フォーク
レポートリポジトリ
リリース
56
v0.33.1
最新の
2026 年 6 月 1 日
+ 55 リリース
このプロジェクトのスポンサーになる
読み込み中にエラーが発生しました。このページをリロードしてください。
https://buy.polar.sh/polar_cl_4hLxfGGjyEuiyX6zBlVp2HNzygaGJZuPT3AvC2cUzlH
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agent Workstation for Codex CLI + Claude Code — with task scheduler, git worktree & remote control, skills management - milisp/codexia

GitHub - milisp/codexia: Agent Workstation for Codex CLI + Claude Code — with task scheduler, git worktree & remote control, skills management · GitHub
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
milisp
/
codexia
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
1,279 Commits 1,279 Commits .github .github .vscode .vscode bots bots docs docs public public scripts scripts src-tauri src-tauri src src .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COMMERCIAL.md COMMERCIAL.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md biome.json biome.json bun.lock bun.lock components.json components.json index.html index.html justfile justfile package.json package.json tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Codexia is a Tauri v2 app for Codex CLI + Claude Code — combining agent workflows, an IDE-like editor, a headless web server, and a prompt notepad in one workspace.
Agent workflows : Task Scheduler for recurring jobs, remote control via headless web server
Workspace : Git worktree management, project file tree, IDE-like editor, prompt notepad, local web preview
Data tools : One-click PDF / XLSX / CSV preview
Ecosystem : MCP server marketplace, agent skills marketplace
Personalization : Theme and accent customization, usage analytics dashboard
brew tap milisp/codexia
brew install --cask codexia
Prebuilt releases (macOS / Linux / Windows)
Enter a prompt and start your agent session.
Create an Agent Task Scheduler job for recurring workflows.
Claude agent rust sdk integration
Frontend: React + TypeScript + Zustand + shadcn/ui in src/
Desktop backend: Tauri v2 + Rust in src-tauri/src/
Headless backend: Axum web server for remote control in src-tauri/src/web/
Agent runtime: Codex app-server JSON-RPC integration for session/turn lifecycle
Real-time updates: WebSocket broadcast stream at /ws for browser clients
src-tauri/src/lib.rs (desktop commands and state)
src-tauri/src/web/server.rs (headless server startup)
src-tauri/src/web/router.rs (HTTP API route surface)
src/services/tauri/ (frontend invoke layer)
Codexia exposes a browser-accessible API when running in web/headless mode:
Health and stream: GET /health , GET /ws
Codex lifecycle: /api/codex/thread/* , /api/codex/turn/* , /api/codex/model/* , /api/codex/approval/*
Automation scheduler: /api/automation/* (create/update/list/run/pause/delete)
Files, git, and terminal: /api/filesystem/* , /api/git/* , /api/terminal/*
Notes and productivity: /api/notes/* , /api/codex/usage/token
Add new API handlers under src-tauri/src/web/handlers/
Register routes in src-tauri/src/web/router.rs
Add corresponding frontend client calls in src/services/tauri/
Process isolation : Agents run in separate processes
Permission control : Configure file and network access per agent
Local storage : All data stays on your machine
Open source : Full transparency through open source code
Telemetry : Opt-in only, off by default
Contributions are welcome. Read CONTRIBUTING.md for setup and workflow.
jeremiahodom/codex-ui — Node.js backend with API/SSE
Itexoft/codexia — SSH integration
nuno5645/codexia — Reasoning and token count events
awesome-codex-cli — curated list of Codex CLI resources
claw-army/claude-node — Python subprocess bridge for Claude Code CLI
Dual-licensed under AGPL-3.0 (open source) and a Commercial License (closed-source / SaaS use).
See COMMERCIAL.md for details.
Agent Workstation for Codex CLI + Claude Code — with task scheduler, git worktree & remote control, skills management
There was an error while loading. Please reload this page .
71
forks
Report repository
Releases
56
v0.33.1
Latest
Jun 1, 2026
+ 55 releases
Sponsor this project
There was an error while loading. Please reload this page .
https://buy.polar.sh/polar_cl_4hLxfGGjyEuiyX6zBlVp2HNzygaGJZuPT3AvC2cUzlH
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
