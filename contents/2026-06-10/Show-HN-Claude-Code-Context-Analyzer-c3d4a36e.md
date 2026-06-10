---
source: "https://github.com/manavgup/context-analyzer"
hn_url: "https://news.ycombinator.com/item?id=48471407"
title: "Show HN: Claude Code Context Analyzer"
article_title: "GitHub - manavgup/context-analyzer: Context window usage analyzer for Claude Code — MCP server + interactive dashboard · GitHub"
author: "manavg76"
captured_at: "2026-06-10T04:33:53Z"
capture_tool: "hn-digest"
hn_id: 48471407
score: 1
comments: 0
posted_at: "2026-06-10T04:25:25Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claude Code Context Analyzer

- HN: [48471407](https://news.ycombinator.com/item?id=48471407)
- Source: [github.com](https://github.com/manavgup/context-analyzer)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T04:25:25Z

## Translation

タイトル: HN を表示: クロード コード コンテキスト アナライザー
記事のタイトル: GitHub - manavgup/context-analyzer: クロード コードのコンテキスト ウィンドウ使用状況アナライザー — MCP サーバー + インタラクティブ ダッシュボード · GitHub
説明: クロード コードのコンテキスト ウィンドウ使用状況アナライザー — MCP サーバー + インタラクティブ ダッシュボード - manavgup/context-analyzer
HN text: クロード コードのコンテキスト ウィンドウ使用状況アナライザー。ツール、圧縮、スキル、ユーザー操作全体でコンテキストがどのように消費されるかを追跡し、それを視覚化してセッションを最適化できます。

記事本文:
GitHub - manavgup/context-analyzer: クロード コードのコンテキスト ウィンドウ使用状況アナライザー — MCP サーバー + インタラクティブ ダッシュボード · GitHub
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
マナブグアップ
/
コンテキストアナライザー
公共
通知
通知設定を変更するにはサインインする必要があります
追加

最後のナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
75 コミット 75 コミット .github/ workflows .github/ workflows .superpowers/ Brainstorm .superpowers/ Brainstorm ドキュメント ドキュメント 証拠 証拠 src/ context_tracker src/ context_tracker 静的 静的テスト テスト .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml Makefile Makefile README.md README.md pyproject.toml pyproject.toml skill-lock.json skill-lock.json uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code のコンテキスト ウィンドウ使用状況アナライザー。ツール、圧縮、スキル、ユーザー操作全体でコンテキストがどのように消費されるかを追跡し、それを視覚化してセッションを最適化できます。
~/.claude/settings.json 経由で Claude コードにフックし、ツール呼び出し、圧縮イベント、セッション ライフサイクル、サブエージェント アクティビティをキャプチャします。
正確な API トークンの使用状況 (input、output、cache_read、cache_creation) のトランスクリプトを解析します。
SQLite 永続化によりセッション データ (セッションごとに 9 テーブル、2,900 以上の行) が保存され、高速なクエリとセッション間分析が可能になります。
単一セッションの視覚化のためのダッシュボード: コンテキストの増加、キャッシュ チャーン、構成、メッセージ インスペクター
/sessions でのセッション間分析: セッション間のパターンを比較: コスト/コールとコンテキスト サイズ、洞察、傾向
uv — 高速 Python パッケージマネージャー
# UV をインストールします (まだインストールされていない場合)
カール -LsSf https://astral.sh/uv/install.sh |しー
クイックスタート
git clone https://github.com/manavgup/context-analyzer.git
cd コンテキストアナライザー
make install-dev # venv を作成し、uv でインストールする
makehook-install #Claude Code フックをインストールします (オプション - ダッシュボードはフックなしで機能します)
make dev # localhost:8080 でダッシュボードを開始します
次に開きます:
http://localhost:8080/ — si

ngle-session ダッシュボード (最新のセッションを選択)
http://localhost:8080/sessions — クロスセッション分析
メイン ダッシュボードには、セッション中にコンテキスト ウィンドウがどのように使用されるかが表示されます。セッション ドロップダウンを使用してセッションを切り替え、予算ボタンを使用して目標しきい値を設定し、スクラバーを使用して API 呼び出しをステップ実行します。
予算のしきい値 (200K / 500K / 700K / 1M) を切り替えて、コンテキストの使用量が目標を超えている箇所を確認します。赤い破線は予算、オレンジ色の点線はクロード コードが自動圧縮される場所、赤い網掛けは危険ゾーンです。
構成パネルには、API によって報告された数値をグラウンド トゥルースとして使用して、各 API 呼び出しでトークンがどこに移動するかが表示されます。ドーナツ グラフは、ツール I/O、会話、システム プレフィックスを分類しており、セッションをスクラブしながらライブで更新されます。
会話ターンで [すべて表示] をクリックすると、ユーザー プロンプト、入力を含むツール呼び出し、ツールの結果、アシスタントの応答などの完全なコンテンツが表示されます。デフォルトでは、長いコンテンツ ブロックは「さらに N 行表示」で折りたたまれます。
/sessions ページには、並べ替え可能な統計、コスト/コール対ピークのコンテキスト散布図、セッション比較グラフ、使用パターンに関する自動生成された分析情報を含むすべてのセッションが表示されます。
予算ライン、危険帯域、オートコンパクトしきい値を含むコンテキスト成長グラフ
予算トグル ボタン (200K / 500K / 700K / 1M)
コールごとの再読み取りコストを示すキャッシュ読み取りチャーン チャート
再起動せずにセッション間を切り替えるセッションドロップダウン
折りたたみ可能なコンテンツブロックを備えたメッセージインスペクター
トークンの内訳ドーナツ チャート (ツール I/O vs 会話 vs システム) — ターンごとに更新
API で報告されたトークン数による構成の内訳 (推定値ではありません)
すべてのビューで統一されたカラー パレット (ドーナツ、コンポジション バー、メッセージ バッジ)
トップの成長ターン、再生付きスクラバー
セッション間の分析

cs ( /セッション )
サマリーカード: 合計セッション、API コール、キャッシュ読み取り、コスト、コールあたりの平均金額
コスト/コール対ピーク コンテキスト散布図 (コンテキスト サイズによるコスト スケーリングを示します)
自動生成された洞察 (コストパターン、高価なセッション、効率指標)
$/call の色分けによるソート可能なセッション テーブル
任意のセッションをクリックして、その単一セッション ビューにドリルダウンします。
make help # 利用可能なターゲットをすべて表示
make dev # リロードしてダッシュボード開発サーバーを開始します
make lint # リンターを実行する
make format # フォーマットコード
make typecheck # mypy を実行
カバレッジを作成 # カバレッジを使用してテストを実行
make verify # 完全な検証スイートを実行 (lint + format + typecheck + test)
建築
クロード コード フック (シェル コマンド)
§── PostToolUse、PostToolUseFailure
§── PreCompact、PostCompact
§── SessionStart、SessionEnd
§── UserPromptSubmit
§── SubagentStart、SubagentStop
└── 命令がロードされました
│
▼
~/.claude/context-trace/<session_id>.jsonl (フックイベント)
~/.claude/projects/<プロジェクト>/<セッションID>.jsonl (トランスクリプト)
│
▼
SQLite (初回アクセス時の自動取り込み)
§── セッション — 概要統計、コスト、モデル
§── api_calls — 呼び出しごとのトークンの内訳
§── ブロック — 開始/終了ターンのあるコンテンツ ブロック
§── ターン — 会話ターンのマッピング
§──hook_events — ツールの呼び出し、失敗、ライフサイクル
§── サブエージェント — エージェントの概要
§── subagent_api_calls — 各サブエージェントのコールごとのチャーン
└── tool_result_offloads — オフロードされたツール出力
│
▼
FastAPI ダッシュボード + MCP サーバー
§── / — 単一セッションのダッシュボード
§── /sessions — クロスセッション分析
§── /api/sessions — 統計情報を含むセッションリスト
§── /api/sessions/trends — セッション間の集計
└── /api/session/{id}/data — 完全なセッション データ (ブロック + チャーン)
キーフィン

実際のセッションからの傷
API コールあたりのコストは、小規模なコンテキストから大規模なコンテキストまで 4.3 倍に拡大します (63,000 ピーク = 0.23 ドル/コール対 721,000 ピーク = 0.98 ドル/コール)
ツール I/O は使用されるコンテキストの 60% 以上を消費します
5ターン以内にコンテキストの30%が古くなった重みになる
システム プロンプト プレフィックスが安定している場合、キャッシュ ヒット率は 96 ～ 98% です。
1M コンテキスト ウィンドウの 50% を超えるセッションにより、不釣り合いなコストが発生する
Claude Code のコンテキスト ウィンドウ使用状況アナライザー — MCP サーバー + インタラクティブ ダッシュボード
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v1.0.0 — クロード コードのコンテキスト ウィンドウ使用状況アナライザー
最新の
2026 年 6 月 8 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Context window usage analyzer for Claude Code — MCP server + interactive dashboard - manavgup/context-analyzer

Context window usage analyzer for Claude Code. Tracks how context is consumed across tools, compaction, skills, and user interactions — then visualizes it so you can optimize your sessions.

GitHub - manavgup/context-analyzer: Context window usage analyzer for Claude Code — MCP server + interactive dashboard · GitHub
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
manavgup
/
context-analyzer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
75 Commits 75 Commits .github/ workflows .github/ workflows .superpowers/ brainstorm .superpowers/ brainstorm docs docs evidence evidence src/ context_tracker src/ context_tracker static static tests tests .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml Makefile Makefile README.md README.md pyproject.toml pyproject.toml skills-lock.json skills-lock.json uv.lock uv.lock View all files Repository files navigation
Context window usage analyzer for Claude Code. Tracks how context is consumed across tools, compaction, skills, and user interactions — then visualizes it so you can optimize your sessions.
Hooks into Claude Code via ~/.claude/settings.json to capture tool calls, compaction events, session lifecycle, and subagent activity
Parses transcripts for exact API token usage (input, output, cache_read, cache_creation)
SQLite persistence stores session data (9 tables, 2,900+ rows per session) for fast queries and cross-session analysis
Dashboard at / for single-session visualization: context growth, cache churn, composition, message inspector
Cross-session analytics at /sessions for comparing patterns across sessions: cost/call vs context size, insights, trends
uv — fast Python package manager
# Install uv (if not already installed)
curl -LsSf https://astral.sh/uv/install.sh | sh
Quick Start
git clone https://github.com/manavgup/context-analyzer.git
cd context-analyzer
make install-dev # create venv and install with uv
make hook-install # install Claude Code hooks (optional — dashboard works without hooks)
make dev # start dashboard on localhost:8080
Then open:
http://localhost:8080/ — single-session dashboard (picks most recent session)
http://localhost:8080/sessions — cross-session analytics
The main dashboard shows how your context window is consumed over the course of a session. Use the session dropdown to switch between sessions, budget buttons to set your target threshold, and the scrubber to step through API calls.
Toggle budget thresholds (200K / 500K / 700K / 1M) to see where your context usage exceeds your target. The dashed red line is your budget, the orange dotted line is where Claude Code auto-compacts, and the red shading is the danger zone.
The composition panel shows where tokens go at each API call, using API-reported numbers as ground truth. The donut chart breaks down Tool I/O vs Conversation vs System prefix — updating live as you scrub through the session.
Click "View All" on any conversation turn to see the full content — user prompts, tool calls with inputs, tool results, and assistant responses. Long content blocks collapse by default with "Show N more lines".
The /sessions page shows all your sessions with sortable stats, a cost/call vs peak context scatter plot, session comparison chart, and auto-generated insights about your usage patterns.
Context Growth chart with budget line, danger band, and autocompact threshold
Budget toggle buttons (200K / 500K / 700K / 1M)
Cache-Read Churn chart showing per-call re-read cost
Session dropdown to switch between sessions without restarting
Message inspector with collapsible content blocks
Token breakdown donut chart (Tool I/O vs Conversation vs System) — updates per-turn
Composition breakdown with API-reported token counts (not estimates)
Unified color palette across all views (donut, composition bars, message badges)
Top growth turns, scrubber with playback
Cross-session analytics ( /sessions )
Summary cards: total sessions, API calls, cache read, cost, avg $/call
Cost/Call vs Peak Context scatter plot (shows cost scaling with context size)
Auto-generated insights (cost patterns, expensive sessions, efficiency metrics)
Sortable session table with $/call color coding
Click any session to drill into its single-session view
make help # see all available targets
make dev # start dashboard dev server with reload
make lint # run linter
make format # format code
make typecheck # run mypy
make coverage # run tests with coverage
make verify # run full verification suite (lint + format + typecheck + test)
Architecture
Claude Code Hooks (shell commands)
├── PostToolUse, PostToolUseFailure
├── PreCompact, PostCompact
├── SessionStart, SessionEnd
├── UserPromptSubmit
├── SubagentStart, SubagentStop
└── InstructionsLoaded
│
▼
~/.claude/context-trace/<session_id>.jsonl (hook events)
~/.claude/projects/<project>/<session_id>.jsonl (transcripts)
│
▼
SQLite (auto-ingest on first access)
├── sessions — summary stats, cost, model
├── api_calls — per-call token breakdown
├── blocks — content blocks with enter/exit turns
├── turns — conversation turn mapping
├── hook_events — tool calls, failures, lifecycle
├── subagents — agent summaries
├── subagent_api_calls — per-call churn for each subagent
└── tool_result_offloads — offloaded tool outputs
│
▼
FastAPI Dashboard + MCP Server
├── / — single-session dashboard
├── /sessions — cross-session analytics
├── /api/sessions — session list with stats
├── /api/sessions/trends — cross-session aggregation
└── /api/session/{id}/data — full session data (blocks + churn)
Key findings from real sessions
Cost per API call scales 4.3x from small to large context (63K peak = $0.23/call vs 721K peak = $0.98/call)
Tool I/O consumes 60%+ of used context
30% of context becomes stale dead weight within 5 turns
Cache hit rate is 96-98% when the system prompt prefix stays stable
Sessions exceeding 50% of the 1M context window drive disproportionate cost
Context window usage analyzer for Claude Code — MCP server + interactive dashboard
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v1.0.0 — Context Window Usage Analyzer for Claude Code
Latest
Jun 8, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
