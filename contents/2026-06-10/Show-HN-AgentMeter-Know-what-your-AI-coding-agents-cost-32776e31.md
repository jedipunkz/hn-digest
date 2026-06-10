---
source: "https://github.com/aussiealex/agentmeter"
hn_url: "https://news.ycombinator.com/item?id=48475587"
title: "Show HN: AgentMeter – Know what your AI coding agents cost"
article_title: "GitHub - aussiealex/agentmeter: Know what your agents cost. Cost intelligence for AI coding agents. · GitHub"
author: "aussiealex27"
captured_at: "2026-06-10T13:17:54Z"
capture_tool: "hn-digest"
hn_id: 48475587
score: 2
comments: 0
posted_at: "2026-06-10T12:54:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentMeter – Know what your AI coding agents cost

- HN: [48475587](https://news.ycombinator.com/item?id=48475587)
- Source: [github.com](https://github.com/aussiealex/agentmeter)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T12:54:42Z

## Translation

タイトル: HN を表示: AgentMeter – AI コーディング エージェントのコストを知る
記事のタイトル: GitHub - aussiealex/agentmeter: エージェントのコストを把握します。 AI コーディング エージェントのコスト インテリジェンス。 · GitHub
説明: エージェントのコストを把握します。 AI コーディング エージェントのコスト インテリジェンス。 - オージーアレックス/エージェントメーター

記事本文:
GitHub - aussiealex/agentmeter: エージェントのコストを把握します。 AI コーディング エージェントのコスト インテリジェンス。 · GitHub
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
オージーアレックス
/
エージェントメーター
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コ

デ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
71 コミット 71 コミット src/agentmeter src/agentmeter テスト testing .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AgentMeter は、AI コーディング エージェントのコスト インテリジェンス レイヤーです。捕らえます
すべてのツール呼び出し、実際のトークンコストの追跡、キャッシュ効率の分析、
支出の削減について指導します。
クロード コード用に構築されました。 Gemini CLI、Codex CLI、および
コパイロット CLI。完全にマシン上で実行されます。クラウドもアカウントもサインアップも必要ありません。
AI コーディング エージェントは強力ですが、高価で不透明です。シングル
セッションでは 50 ドル以上の API コストが消費される可能性があり、実際に行われるまでわかりません。
請求書が届く。プロンプト キャッシュにより、コストが 3 つのトークン タイプに 3 つに分割されます。
異なる料金。 API ダッシュボードにはリクエストごとの数値が表示されます。エージェントメーター
本当の姿を見せてくれます。
$ エージェントメーターのコスト
ProjectX — $19.17 (4,982,859 トークン、83 LLM コール)
─────────────────────
キャッシュ読み取り: 4,358,360 (87%)
キャッシュ作成: 592,819 (12%)
出力: 17,284 (0.3%)
入力: 14,396 (0.3%)
キャッシュ効率: 88%
節約されたキャッシュ: $58.84 (75%)
クイックスタート
pip インストール エージェントの使用法
1.フックを取り付ける(30秒)
エージェントメーターフックインストールクロード
これにより、構成スニペットが出力されます。 ~/.claude/settings.json とすべてのファイルに追加します。
ツール呼び出しは自動的に計測されます。他のエージェントについても同様です:
Agentmeter フック インストール Gemini
Agentmeter フック インストール コーデックス
Agentmeter フック インストール コパイロット
2. エージェントを通常どおりに使用する
コードを書くだけです。 AgentMeter はすべてのツール呼び出しをバックグラウンドで 5 ミリ秒未満で記録します
頭上。
$年齢

NTメーターの統計
AgentMeter 統計 (今日)
─────────────────────
合計: 847 コール | 3 つのエラー | 142msのツール時間
████████████████████ 312 件の通話を読む
████████████ 198件の通話を編集
バッシュ █████████ 147 回の通話
Grep ██████ 89 回の呼び出し
特長
Claude Code セッション記録からの実際のトークンコスト — 推定値ではありません:
Agentmeter のコスト、トークンの内訳を含む最近のセッション数
Agentmetercost < session-id > # 詳細な内訳 (部分的な ID は機能します)
Agentmeter の予測 # 月次支出予測
Agentmeter 戦略 # プロジェクトごとのコスト分析 + アドバイス
Agentmeter モデル # モデル層分析 + what-if 比較
キャッシュ インテリジェンスが組み込まれています。あらゆるコスト ビューでキャッシュ効率がわかります。
(プロンプト キャッシュがどの程度うまく機能しているか) とキャッシュの節約 (金額の節約)
vs キャッシュなし)。
「エージェントはコスト以上の節約をしましたか?」という質問に答えます。
$ エージェントメーターの値
価値レポート (開発率: $150/時間)
───────────────────────
AgentMeter コスト $21.43 価値 $5017.50 マルチ 234.1x 品質 95 (良好)
時間 2007 分 結果 3c 909t 12f 3l
LLMM コスト $57.06 価値 $10845.00 乗算 190.0x 品質 100 (良好)
時間 4338 分 結果 8c 1880t 40f 6l
───────────────────────
合計コスト $417.15 価値 $37505.00 乗算 89.9x 時間

15002m節約
各セッションでは以下が得られます。
コスト — API レートからの実際のトークン支出
値 — コミット、テスト、変更されたファイル、lint パスに基づいて、開発者が節約した推定時間
乗数 — 価値 / コスト。 1 倍を超えるものは、あなたが前に出たことを意味します
品質スコア (0-100) — エラー、テスト失敗、再試行、lint エラーに対してペナルティが課されます
Agentmeter 値 最近のセッション数
Agentmeter value -p MyProject # プロジェクトによるフィルター
Agentmeter value --rate 200 # 開発時間レートを設定します (デフォルトは 150 ドル)
セッションコーチング
ツール呼び出しパターンを分析し、実用的なアドバイスを提供します。
Agentmeter Coach レビュー # 最新のセッションをレビューする
Agentmeter Coach review -p myproject # プロジェクト名によるレビュー
Agentmeter advice # クロスセッション支出分析
$ Agentmeter Coach review -p myproject
セッションレビュー — MyProject
2026-05-28 14:03
───────────────────
131 ツールコール $68.98
バッシュ █████ 39 (30%)
編集 █████ 34 (26%)
███ 22を読む (17%)
結果: 2 コミット、27 ファイル変更、1188 テスト合格
効率: 8/10
検出されたパターン:
！ Grep+Glob が 16x と呼ばれる (Grep 12、Glob 4)
何をどこで探しているのかをエージェントに伝えてください。
17 個の固有のファイルを読み取る
幅広く探索しているんですね。どのファイルが重要かを書くために 2 分を投資してください。
編集とテストのループ、広範囲の探索、繰り返しを含む 13 のパターンを検出
ファイル読み取り、高速バースト、キャッシュ書き込みの無駄、およびキャッシュ効率の低さ。
セッション中のリアルタイムのナッジ。セッションがコストのしきい値を超えると、
AgentMeter は、コーチング メッセージを含む 1 つのツール呼び出しをブロックします。エージェントはそれを認識し、
調整して再試行します。
Agentmeter Coach start # session-start プロファイリング + ガイダンス
Agentmeter Coach start -tdevelopment # タスクタイプを設定
エージェントメーター コーチ コンテックス

t # CLAUDE.md コーチング ブロックを生成
Agentmeter Coach レビュー # セッション後の効率分析
Agentmeter Coach は現在のしきい値の数を表示します
Agentmeter Coach コールの設定 50 100 # コールのしきい値を設定
Agentmeter Coach setrepeat 15 # 繰り返しツールのしきい値
エージェントメーターコーチ # でイエローカードを有効にする
Agentmeter Coach off # イエローカードを無効にする
ツール呼び出し統計
Agentmeter stats # 今日の統計
Agentmeter stats --week # 今週
Agentmeter stats --all # すべての時間
Agentmeter コールの最近の個別コール数
Agentmeter は --tool Bash を呼び出します # ツール名でフィルターします
Agentmeter セッション # セッションの内訳と結果
エージェントメーターの毎日の毎日の合計数（棒グラフ付き）
予算の執行
制限を設定すると、制限を超えると AgentMeter がブロックまたは警告を出します。否認
クラッシュではなく、エージェントが推論できる有益なエラーを返します。
Agentmeter バジェット セット セッション 50 # セッションあたり最大 50 コール
Agentmeter の予算を毎日設定 200 # 1 日あたり最大 200 件のコール
Agentmeter の予算を毎日設定 100 -s サーバーごとの 1 日あたりのメール数の制限
Agentmeter Budget set session 30 -a warn # 警告するがブロックしない
Agentmeter Budget show # すべてのルールをリストする
サーキットブレーカー
暴走ループに対する速度ベースの保護:
Agentmeter ブレーカー セット 20 60 # 60 秒間に 20 回のコール後にトリップ
Agentmeter ブレーカー セット 10 30 -c 600 # カスタム クールダウン (600 秒)
ウェブダッシュボード
Agentmeter ダッシュボード # localhost:8070 で開きます
6 つのビュー: 概要、コスト分割のあるプロジェクト、セッション、日々の傾向、
レート表と戦略の推奨事項。
セッション開始時にエージェントが読み取ることができるコストの概要を生成します。
Agentmeter の概要 # すべてのプロジェクト
Agentmeter summary -p myproject # プロジェクト固有
Agentmeter の概要 >> CLAUDE.md # エージェント コンテキストに挿入
MCPプロキシ
MCP サーバーを実行している場合、AgentMeter はエージェントとサーバーの間に配置できます。
透過的なプロキシとして:
エージェントメーターwra

p python -m some.mcp.server
Agentmeter ラップ --name myserver python -m some.mcp.server
フック データとプロキシ データは同じデータベース (組み込みツールと MCP) にフィードされます。
ツールを 1 つのビューにまとめます。
Agentmeter # JSONL を標準出力にエクスポート
Agentmeter エクスポート --tool 読み取り --since 2026-05-01 # フィルター済み
仕組み
2 つのキャプチャ パスが同じ SQLite データベースにフィードします。
パス 1: フック (プライマリ — フック システムを持つ任意のエージェントで動作します)
エージェントの組み込みツール -> PostToolUse フック -> エージェントメーター -> SQLite DB
パス 2: MCP プロキシ (MCP サーバー トラフィックの測定用)
エージェント -> AgentMeter プロキシ -> MCP サーバー -> SQLite DB
建築
ローカルファースト — WAL モードの SQLite、オフラインで動作、クラウド依存なし
透過的 — ツール呼び出しデータを変更せず、ただ観察するだけです
クロスプラットフォーム — Linux、macOS、Windows
高速フック — オーバーヘッドは 5ms 未満、標準ライブラリのみ、エージェントはクラッシュしません
エージェント
フックタイプ
ステータス
クロード・コード
ポストツール使用
フルサポート（プライマリ）
ジェミニ CLI
アフターツール
アダプタは構築されていますが、本番環境ではテストされていません
コーデックス CLI
ポストツール使用
アダプタは構築されていますが、本番環境ではテストされていません
コパイロット CLI
postToolUse
アダプタは構築されていますが、本番環境ではテストされていません
CLI リファレンス
コマンドはグループ化されています。agentmeter --help を実行すると、コマンドがカテゴリ別に表示されます。
ほとんどのコマンドは、プロジェクト名でフィルタリングするために -p/--project を受け入れます。
外部サービスや API キーはありません
AgentMeter でコストが節約できる場合は、リポジトリにスターを付けてください。他の人がリポジトリを見つけやすくなります。
バグレポートと機能リクエスト: GitHub の問題
質問とディスカッション: GitHub ディスカッション
オーストラリアのアデレードにある Andre300 によって建てられました。
エージェントにかかる費用を把握しましょう。 AI コーディング エージェントのコスト インテリジェンス。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
する

私の個人情報を共有しないでください

## Original Extract

Know what your agents cost. Cost intelligence for AI coding agents. - aussiealex/agentmeter

GitHub - aussiealex/agentmeter: Know what your agents cost. Cost intelligence for AI coding agents. · GitHub
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
aussiealex
/
agentmeter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
71 Commits 71 Commits src/ agentmeter src/ agentmeter tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
AgentMeter is the cost intelligence layer for AI coding agents. It captures
every tool call, tracks real token costs, analyses caching efficiency, and
coaches you on reducing spend.
Built for Claude Code . Also supports Gemini CLI, Codex CLI, and
Copilot CLI. Runs entirely on your machine. No cloud, no accounts, no signup.
AI coding agents are powerful, but they're expensive and opaque. A single
session can burn through $50+ in API costs, and you won't know until the
bill arrives. Prompt caching splits costs into three token types at three
different rates. Your API dashboard shows per-request numbers. AgentMeter
shows you the real picture.
$ agentmeter cost
ProjectX — $19.17 (4,982,859 tokens, 83 LLM calls)
──────────────────────────────────────────────────────────────
Cache reads: 4,358,360 (87%)
Cache creation: 592,819 (12%)
Output: 17,284 (0.3%)
Input: 14,396 (0.3%)
Cache efficiency: 88%
Cache saved: $58.84 (75%)
Quick Start
pip install agent-usage
1. Install the hook (30 seconds)
agentmeter hook install claude
This prints a config snippet. Add it to ~/.claude/settings.json and every
tool call is metered automatically. Same for other agents:
agentmeter hook install gemini
agentmeter hook install codex
agentmeter hook install copilot
2. Use your agent normally
Just code. AgentMeter records every tool call in the background with <5ms
overhead.
$ agentmeter stats
AgentMeter Stats (today)
────────────────────────────────────────────────────────────
Total: 847 calls | 3 errors | 142ms tool time
Read ████████████████████ 312 calls
Edit ████████████ 198 calls
Bash █████████ 147 calls
Grep ██████ 89 calls
Features
Real token costs from Claude Code session transcripts — not estimates:
agentmeter cost # recent sessions with token breakdown
agentmeter cost < session-id > # detailed breakdown (partial ID works)
agentmeter forecast # monthly spend projection
agentmeter strategy # per-project cost analysis + advice
agentmeter model # model tier analysis + what-if comparison
Cache intelligence is built in. Every cost view shows cache efficiency
(how well prompt caching is working) and cache savings (dollars saved
vs uncached).
Answers the question: "Did the agent save more than it cost?"
$ agentmeter value
Value Report (dev rate: $150/hr)
────────────────────────────────────────────────────────────────────────
AgentMeter cost $21.43 value $5017.50 mult 234.1x quality 95 (good)
time 2007m outcomes 3c 909t 12f 3l
LLMM cost $57.06 value $10845.00 mult 190.0x quality 100 (good)
time 4338m outcomes 8c 1880t 40f 6l
────────────────────────────────────────────────────────────────────────
Total cost $ 417.15 value $37505.00 mult 89.9x time 15002m saved
Each session gets:
Cost — real token spend from API rates
Value — estimated developer time saved, based on commits, tests, files changed, lint passes
Multiplier — value / cost. Anything >1x means you came out ahead
Quality score (0-100) — penalised for errors, test failures, retries, lint errors
agentmeter value # recent sessions
agentmeter value -p MyProject # filter by project
agentmeter value --rate 200 # set dev hourly rate ($150 default)
Session Coaching
Analyses tool call patterns and gives actionable advice:
agentmeter coach review # review most recent session
agentmeter coach review -p myproject # review by project name
agentmeter advise # cross-session spend analysis
$ agentmeter coach review -p myproject
Session Review — MyProject
2026-05-28 14:03
───────────────────────────────────────────────────────
131 tool calls $68.98
Bash █████ 39 (30%)
Edit █████ 34 (26%)
Read ███ 22 (17%)
Outcome: 2 commits, 27 files changed, 1188 tests passed
Efficiency: 8/10
Patterns detected:
! Grep+Glob called 16x (Grep 12, Glob 4)
Tell the agent what you're looking for and where.
Read 17 unique files
You're exploring broadly. Invest 2 min writing which files matter.
Detects 13 patterns including edit-test loops, broad exploration, repeated
file reads, high velocity bursts, cache write waste, and low cache efficiency.
Real-time mid-session nudges. When your session crosses a cost threshold,
AgentMeter blocks one tool call with a coaching message — the agent sees it,
adjusts, and retries.
agentmeter coach start # session-start profiling + guidance
agentmeter coach start -t development # set task type
agentmeter coach context # generate CLAUDE.md coaching block
agentmeter coach review # post-session efficiency analysis
agentmeter coach show # current thresholds
agentmeter coach set calls 50 100 # set call thresholds
agentmeter coach set repeat 15 # repeated tool threshold
agentmeter coach on # enable yellow cards
agentmeter coach off # disable yellow cards
Tool Call Stats
agentmeter stats # today's stats
agentmeter stats --week # this week
agentmeter stats --all # all time
agentmeter calls # recent individual calls
agentmeter calls --tool Bash # filter by tool name
agentmeter sessions # session breakdowns with outcomes
agentmeter daily # daily totals with bar chart
Budget Enforcement
Set limits and AgentMeter will block or warn when they're exceeded. Denials
return informative errors the agent can reason about — not crashes.
agentmeter budget set session 50 # max 50 calls per session
agentmeter budget set daily 200 # max 200 calls per day
agentmeter budget set daily 100 -s mail # per-server daily limit
agentmeter budget set session 30 -a warn # warn but don't block
agentmeter budget show # list all rules
Circuit Breakers
Velocity-based protection against runaway loops:
agentmeter breaker set 20 60 # trip after 20 calls in 60 seconds
agentmeter breaker set 10 30 -c 600 # custom cooldown (600s)
Web Dashboard
agentmeter dashboard # open at localhost:8070
Six views: overview, projects with cost split, sessions, daily trends,
rate card, and strategy recommendations.
Generate a cost summary your agent can read at session start:
agentmeter summary # all projects
agentmeter summary -p myproject # project-specific
agentmeter summary >> CLAUDE.md # inject into agent context
MCP Proxy
If you run MCP servers, AgentMeter can sit between your agent and the server
as a transparent proxy:
agentmeter wrap python -m some.mcp.server
agentmeter wrap --name myserver python -m some.mcp.server
Hook data and proxy data feed the same database — built-in tools and MCP
tools in one view.
agentmeter export # JSONL to stdout
agentmeter export --tool Read --since 2026-05-01 # filtered
How It Works
Two capture paths feed the same SQLite database:
Path 1: Hook (primary — works with any agent that has a hook system)
Agent's built-in tools -> PostToolUse hook -> agentmeter -> SQLite DB
Path 2: MCP Proxy (for metering MCP server traffic)
Agent -> AgentMeter proxy -> MCP Server -> SQLite DB
Architecture
Local-first — SQLite with WAL mode, works offline, no cloud dependency
Transparent — never modifies tool call data, just observes
Cross-platform — Linux, macOS, Windows
Fast hooks — <5ms overhead, stdlib only, never crashes the agent
Agent
Hook Type
Status
Claude Code
PostToolUse
Full support (primary)
Gemini CLI
AfterTool
Adapter built, untested in production
Codex CLI
PostToolUse
Adapter built, untested in production
Copilot CLI
postToolUse
Adapter built, untested in production
CLI Reference
Commands are grouped — run agentmeter --help to see them by category.
Most commands accept -p/--project to filter by project name.
No external services or API keys
If AgentMeter saves you money, star the repo — it helps others find it.
Bug reports & feature requests: GitHub Issues
Questions & discussion: GitHub Discussions
Built by Andre300 in Adelaide, Australia.
Know what your agents cost. Cost intelligence for AI coding agents.
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
