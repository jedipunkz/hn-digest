---
source: "https://github.com/neochoon/agenthud"
hn_url: "https://news.ycombinator.com/item?id=48474693"
title: "Show HN: AgentHUD – Live TUI and daily digest for parallel Claude Code sessions"
article_title: "GitHub - neochoon/agenthud: Claude Code TUI dashboard: live monitor for parallel sessions and sub-agents, with LLM-generated daily/weekly digest · GitHub"
author: "neochoon"
captured_at: "2026-06-10T13:19:40Z"
capture_tool: "hn-digest"
hn_id: 48474693
score: 1
comments: 0
posted_at: "2026-06-10T11:25:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentHUD – Live TUI and daily digest for parallel Claude Code sessions

- HN: [48474693](https://news.ycombinator.com/item?id=48474693)
- Source: [github.com](https://github.com/neochoon/agenthud)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T11:25:25Z

## Translation

タイトル: 表示 HN: AgentHUD – 並行クロード コード セッションのライブ TUI と毎日のダイジェスト
記事のタイトル: GitHub - neochoon/agenthud: Claude Code TUI ダッシュボード: 並列セッションとサブエージェントのライブ モニター、LLM で生成された日次/週次ダイジェスト · GitHub
説明: Claude Code TUI ダッシュボード: 並列セッションとサブエージェントのライブ モニター (LLM によって生成された日次/週次ダイジェスト付き) - neochoon/agenthud

記事本文:
GitHub - neochoon/agenthud: Claude Code TUI ダッシュボード: 並列セッションとサブエージェントのライブ モニター、LLM によって生成された日次/週次ダイジェスト付き · GitHub
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
ネオチョン
/
エージェントハッド
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
448 コミット 448 コミット .github/ workflows .github/ workflows デモ デモ ドキュメント ドキュメント スクリプト スクリプト src src テスト テスト .gitignore .gitignore BACKLOG.md BACKLOG.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md FEATURES.md FEATURES.md README.md README.md biome.json biome.json claude-watch.sh claude-watch.sh package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code の可観測性レイヤー。ライブ セッションの表示、構造化されたアクティビティ ログのエクスポート、1 日または 1 週間の LLM ダイジェストへの要約をすべて 1 つの CLI から実行できます。
AgentHUD は ~/.claude/projects/ から Claude Code のセッション ファイルを読み取り、次の 3 つのことを提供します。
ライブ モニター (agenthud) — すべてのプロジェクト、セッション、サブエージェント、アクティビティを発生時に表示する分割ビュー TUI
構造化エクスポート (agenthud レポート) — スクリプト、ダッシュボード、またはその他の LLM にパイプするためのマークダウンまたは JSON
LLM ダイジェスト (agenthud summary) — クロード CLI を介してエンジニアリング概要に合成された日または日付範囲
→ すべてのフラグ、キーバインド、設定キー、ファイル パス、および環境変数など、全表面については FEATURES.md を参照してください。
Node.js 20以上が必要です。 Claude Code を使用しているときに、別のターミナルで Agenthud を開きます。プレス ？ TUI内でアプリ内ヘルプを参照してください。
npxエージェントハッド
# または: bunx エージェントハッド
日常的に使用するためにインストールする
npm i -g エージェントハッド
# または: bun i -g Agenthud
プラットフォームのサポート。主な開発は macOS と Linux 上で行われます。完全なテスト スイートは、CI の 3 つのプラットフォームすべて (Windows を含む) で実行されます。 Windows ランタイムの動作は手動のスモーク ジョブによって実行されますが、毎日実行されるわけではありません。価値のある問題が存在します。

ed のバグレポート。
#ライブモニター
Agenthud # すべてのクロード プロジェクト
Agenthud --cwd # $PWD を含むプロジェクトをスコープとする
Agenthud --once # スナップショット モード、代替画面なし
# 活動報告
Agenthud report --date today # マークダウンとしての今日のアクティビティ
Agenthud レポート --format json # スクリプトで読み取り可能
Agenthud report --with-git # git コミットをタイムラインにマージします
# LLM の概要
Agenthud summary --今日の日付 # `claude -p` による毎日の概要
Agenthud summary --last 7d # 過去 7 日間の日をまたいだ合成
Agenthud summary -oI # 概要と概要のインデックスを開きます
あなたが見るもの
TUI は、プロジェクト ツリー (上部) とアクティビティ ビューアー (下部) に分かれています。
┌─ プロジェクト ──── ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─
│ > エージェントハッド ~/WestbrookAI/エージェントハッド 13m │
│ #864f [ホット] ログインフローの認証バグを修正 │
│ §─ » コードレビューア │
│ (#398c [温]) │
│ myproject ~/work/myproject 2d │
│ #def4 [ホット] OAuth サポートを追加 │
│ ... 12 件のコールド プロジェクト │
━━━━━━━━━━━━━━━━━━━━━━━┘
┌─アクティビティ・agenthud ───────────────────┐
│ [10:23] ○ src/ui/App.tsx を読む │
│ [10:23] ~ src/ui/App.tsx を編集 │
│ [10:23] $ Bash npm テスト │
│ [10:23] < 応答テストは正常に完了しました │
│ [10:25] ⠧ src/auth/oauth.ts を編集 ← 太字 + スピナー = live │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

────┘
セッションには、JSONL ファイルがどのくらい最近に触れられたかに基づいて、色付きのバッジ ([ホット] / [ウォーム] / [クール] / [コールド]) が付けられます。コールド プロジェクトは ... N コールド プロジェクトの監視員の下で崩壊します。アクティビティ上で ↵ を押すと、スクロール可能な詳細ビューが開きます。
キーバインドとバッジの完全なリファレンス: FEATURES.md 。
~/.agentthud/config.yaml は、最初の実行時に適切なデフォルトで自動作成されます。 CLI フラグは、呼び出しごとに設定値をオーバーライドします。解決順序は CLI フラグ → summary.<key> → report.<key> → 組み込みデフォルト であり、有効な値は各レポート/概要の実行開始時に標準エラー出力に出力されます。
アプリ管理の UI 状態 ( h によって切り替えられる非表示のプロジェクト/セッション/サブエージェント) は、 ~/.agenthud/state.yaml に個別に存在します。
完全なスキーマ、ファイル パス、および環境変数: FEATURES.md → Config 。
参照: FEATURES.md — すべてのフラグ、キーバインド、設定キー、ファイル パス、環境変数
Claude Code TUI ダッシュボード: LLM が生成する日次/週次ダイジェストを含む、並列セッションとサブエージェントのライブ モニター
www.npmjs.com/package/agentthud
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
39
v0.13.3
最新の
2026 年 6 月 10 日
+ 38 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Code TUI dashboard: live monitor for parallel sessions and sub-agents, with LLM-generated daily/weekly digest - neochoon/agenthud

GitHub - neochoon/agenthud: Claude Code TUI dashboard: live monitor for parallel sessions and sub-agents, with LLM-generated daily/weekly digest · GitHub
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
neochoon
/
agenthud
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
448 Commits 448 Commits .github/ workflows .github/ workflows demo demo docs docs scripts scripts src src tests tests .gitignore .gitignore BACKLOG.md BACKLOG.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md FEATURES.md FEATURES.md README.md README.md biome.json biome.json claude-watch.sh claude-watch.sh package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
An observability layer for Claude Code . See your live sessions, export structured activity logs, and summarize a day or a week into an LLM digest — all from one CLI.
AgentHUD reads Claude Code's session files from ~/.claude/projects/ and gives you three things:
Live monitor ( agenthud ) — a split-view TUI showing every project, session, sub-agent, and activity as it happens
Structured export ( agenthud report ) — Markdown or JSON for piping to scripts, dashboards, or other LLMs
LLM digest ( agenthud summary ) — a day or a date range synthesized into an engineering summary via the claude CLI
→ See FEATURES.md for the full surface — every flag, keybinding, config key, file path, and env var.
Requires Node.js 20+. Open agenthud in a separate terminal while using Claude Code; press ? inside the TUI for in-app help.
npx agenthud
# or: bunx agenthud
Install for daily use
npm i -g agenthud
# or: bun i -g agenthud
Platform support. Primary development is on macOS and Linux; the full test suite runs on all three platforms in CI (including Windows). Windows runtime behavior is exercised by a manual smoke job but isn't daily-driven — issues there are valued bug reports.
# Live monitor
agenthud # all Claude projects
agenthud --cwd # scope to the project containing $PWD
agenthud --once # snapshot mode, no alt-screen
# Activity report
agenthud report --date today # today's activity as markdown
agenthud report --format json # script-readable
agenthud report --with-git # merge git commits into the timeline
# LLM summary
agenthud summary --date today # daily summary via `claude -p`
agenthud summary --last 7d # cross-day synthesis of last 7 days
agenthud summary -oI # open the summary + the summaries index
What you see
The TUI splits into a project tree (top) and an activity viewer (bottom):
┌─ Projects ─────────────────────────────────────────────────┐
│ > agenthud ~/WestbrookAI/agenthud 13m │
│ #864f [hot] Fix the auth bug in login flow │
│ ├─ » code-reviewer │
│ (#398c [warm]) │
│ myproject ~/work/myproject 2d │
│ #def4 [hot] Add OAuth support │
│ ... 12 cold projects │
└────────────────────────────────────────────────────────────┘
┌─ Activity · agenthud ──────────────────────────────────────┐
│ [10:23] ○ Read src/ui/App.tsx │
│ [10:23] ~ Edit src/ui/App.tsx │
│ [10:23] $ Bash npm test │
│ [10:23] < Response Tests passed successfully │
│ [10:25] ⠧ Edit src/auth/oauth.ts ← bold + spinner = live │
└────────────────────────────────────────────────────────────┘
Sessions get colored badges — [hot] / [warm] / [cool] / [cold] — based on how recently their JSONL file was touched. Cold projects collapse under a ... N cold projects sentinel. Press ↵ on any activity to open a scrollable detail view.
Full keybinding and badge reference: FEATURES.md .
~/.agenthud/config.yaml is auto-created on first run with sensible defaults. CLI flags override config values per-invocation. Resolution order is CLI flag → summary.<key> → report.<key> → built-in default , and the effective values print to stderr at the start of every report / summary run.
App-managed UI state (hidden projects/sessions/sub-agents toggled by h ) lives separately in ~/.agenthud/state.yaml .
Full schema, file paths, and env vars: FEATURES.md → Config .
Reference: FEATURES.md — every flag, keybinding, config key, file path, env var
Claude Code TUI dashboard: live monitor for parallel sessions and sub-agents, with LLM-generated daily/weekly digest
www.npmjs.com/package/agenthud
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
39
v0.13.3
Latest
Jun 10, 2026
+ 38 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
