---
source: "https://github.com/DeibyGS/claudestat"
hn_url: "https://news.ycombinator.com/item?id=48559521"
title: "How HN: Claudestat – real-time Claude Code monitor with loop detection"
article_title: "GitHub - DeibyGS/claudestat: Real-time execution trace and cost intelligence for Claude Code · GitHub"
author: "deibyg"
captured_at: "2026-06-16T19:19:43Z"
capture_tool: "hn-digest"
hn_id: 48559521
score: 1
comments: 0
posted_at: "2026-06-16T18:12:43Z"
tags:
  - hacker-news
  - translated
---

# How HN: Claudestat – real-time Claude Code monitor with loop detection

- HN: [48559521](https://news.ycombinator.com/item?id=48559521)
- Source: [github.com](https://github.com/DeibyGS/claudestat)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T18:12:43Z

## Translation

タイトル: HN: Claudestat – ループ検出によるリアルタイムのクロード コード モニターの方法
記事タイトル: GitHub - DeibyGS/claudestat: クロード コードのリアルタイム実行トレースとコスト インテリジェンス · GitHub
説明: クロード コードのリアルタイム実行トレースとコスト インテリジェンス - DeibyGS/claudestat

記事本文:
GitHub - DeibyGS/claudestat: クロード コードのリアルタイム実行トレースとコスト インテリジェンス · GitHub
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
デイビーGS
/
クローデスタット
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
chan にサインインする必要があります

GE通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
152 コミット 152 コミット .github .github アセット アセット ダッシュボード ダッシュボード ドキュメント ドキュメント フック フック スクリプト スクリプト src src テスト テスト .gitignore .gitignore .npmignore .npmignore .nvmrc .nvmrc CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md llms.txt llms.txt package-lock.json package-lock.json package.json package.json run-tests.sh run-tests.sh tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ライブ AI コーディング モニター — リアルタイム トレース、クォータ ガード、MCP サーバー
ほとんどのツールは、セッション終了後にログを読み取ります。 claudestat は、起動時にすべてのイベントにフックします。
AI が現在費やしている金額を確認し、割り当てに達する前にアラートを受け取り、端末内から AI の使用量についてクロードに質問します。
Claude Code および OpenCode で動作します。クラウド依存性ゼロ。純粋な Node.js。 macOS、Linux、Windows 上で動作します。
インストール •
クイックスタート •
コマンド •
ダッシュボード •
よくある質問 •
貢献する
ライブダッシュボード、ターミナルトレース、クォータガード - すべてリアルタイムで実行
ccusage のようなツールは歴史を振り返るのに最適です。 claudestat はコーディング中に使用します。
Claude Code のフック システムを利用して、すべてのイベントが発生した瞬間にキャプチャし、すべてを SQLite にローカルに保存し、単なるレポートではなく、ライブ ダッシュボード、クォータ アラート、MCP サーバーを提供します。
claudestat が役立つ場合は、⭐ を付けてください。他の開発者が見つけやすくなります。
claudestat には MCP サーバーが同梱されています。登録すると、ターミナルを離れることなく、Claude Code 自体の使用法について質問することができます。
クロード mcp add claudestat -s user --claudestat-

mcp
それから次のように尋ねてください。
> 今週はいくら使いましたか?
> コストの上位 5 つのツールは何ですか?
> モデルごとの使用状況を分析する
> 来月の費用の予測はいくらですか?
完全な MCP リファレンス →
ライブツールトレース - すべての呼び出しと実行時の持続時間とトークンコスト
OpenCode のサポート — OpenCode セッション用の同じライブ ダッシュボード (ツール呼び出し、プロンプト、モデル、インテント)
クォータ ガード — 70%、85%、95% でアラート。オプションのキルスイッチは新しいセッションを X% でブロックします
ループ検出器 — 推定無駄コストを使用してコンテキスト スラッシングにフラグを立てます
上位のツール — どのツールが予算の多くを消費するかを把握します。展開可能な「その他」行により、トップ 10 以外のツールが表示されます
コスト予測 — 傾向、信頼区間、R² を使用した線形回帰
請求ブロック — 5 時間の請求ウィンドウとブロックごとの累積支出を追跡します ( claudestat ブロック )
セッション共有 — セッションの概要をフォーマットされた ASCII または JSON としてエクスポートします ( claudestat share )。
52 週間のアクティビティ ヒートマップ — Analytics で年間のコーディング アクティビティを視覚化します。
前期比デルタ — アナリティクス KPI は、前期と比較して ↑↓% の傾向を示しています
プロジェクトの検索とフィルター — プロジェクトを即座に見つけます。アクティビティ、コスト、または効率のしきい値によるフィルター
Web ダッシュボード — 6 つのタブ: ライブ、履歴、プロジェクト、分析、トップ、システム
MCP サーバー — クロードが自身の使用法に関する質問に答えることができる 7 つのツール
毎週のインサイト — 実用的なヒントを含むパターン分析
マルチソース — ワンクリックで Claude Code セッションと OpenCode セッションを切り替える
ソース フィルター — すべてのタブにわたって、KPI、チャート、およびツール ランキングをクロード コード / OpenCode ごとにフィルターします。
マルチツール調整 — 実際の衝突検出を備えたライブ インテント パネル (CC および OC によって編集された同じファイル)
npm install -g @statforge/claudestat && claudestat setup
http://localhost:7337 を開きます
クロード コード セッションを開始してイベントを視聴してください

ts が流れ込みます。それだけです。
Node.js >= 22 (node:sqlite に必要)
クロード コードがインストールされました ( npm install -g @anthropic-ai/claude-code )
npm install -g @statforge/claudestat && claudestat setup
claudestat セットアップは、Claude コード フックをインストールし、デーモンをシステム サービス (macOS では launchd、Linux では systemd) として登録します。sudo は必要ありません。ログインするたびにデーモンが自動的に起動します。
NVM を使用しますか?デフォルトの Node バージョンを使用していることを確認してください。
nvm はデフォルトを使用します && npm install -g @statforge/claudestat && claudestat setup
フックが有効になるように、セットアップ後に Claude Code を再起動します。
npm install -g @statforge/claudestat
claudestat install # クロード コードにフックをインストールします
claudestat start # デーモンを手動で起動します
コマンド
コマンド
説明
クローデスタットのセットアップ
ワンコマンドセットアップ: フックのインストール + デーモンをシステムサービスとして登録
claudestat セットアップ --アンインストール
フックとシステムサービスを削除します
claudestat 開始 / 停止 / 再起動
バックグラウンドデーモンを管理する
claudestatのインストール/アンインストール
Claude Code フックのインストールまたは削除
クローデスタット時計
ライブターミナルトレースビュー
クローデスタットのステータス
クォータ、コスト、バーンレートを表示する
クローデスタットトップ
コスト、コール数、または期間に基づいてツールをランク付けします
毎週のクローデスタット
週ごとの使用量の概要
クローデスタットの洞察
使用状況に関する深い洞察: コスト、キャッシュ、効率、モデル
クローデスタットプロジェクト
線形回帰によるコスト予測
クローデスタット構成
構成の表示または編集
クローデスタット博士
インストールの健全性を確認し、問題を診断する
クローデスタットブロック
5 時間の請求ブロック履歴を表示する
claudestat 共有 [セッション ID]
セッションの概要を ASCII または JSON としてエクスポートします (クリップボードにコピーするには --copy を使用します)
claudestat エクスポート [形式]
セッションデータをJSONまたはCSVにエクスポート
クローデスタットロースト
皮肉な使用法分析
クローデスタットのバージョン
バージョンを表示してアップデートを確認する
o を含む完全なコマンド リファレンス

出力例 →
ダッシュボードは http://localhost:7337 にあり、次の 6 つのタブがあります: ライブ (ソース バッジ付きのリアルタイム トレース、拡張可能な Bash コマンド、およびソースごとの最終タスクのサブタイトル)、履歴 (日付セレクター 7/14/30/90d を使用した日付別のセッション、マージされたセッション バッジ、検索、コスト フィルター、および比較パネル)、プロジェクト (毎週のヒートマップを含むグリッド、検索入力、およびアクティブ、高コスト、または低効率のフィルター)プロジェクト）、分析（支出 + トークン + 時間 + 前期比 ↑↓% デルタ付きの効率 KPI、入力/出力/キャッシュ別の積み上げトークン チャート、52 週間のアクティビティ ヒートマップ、ソース フィルター、週次 AI レポート）、上位（コスト予測を含むコスト/数/期間によるツール ランキング、上位 10 を超えるツールの拡張可能な「その他」行）、およびシステム（フック、エージェント、スキル、ワークフロー、コンテキスト ファイル制限、作業モード）ディストリビューション、OpenCode 設定、切り捨て警告のあるメモリ ファイル、claudestat 設定）。
OpenCode を Claude Code と一緒に実行すると、claudestat は両方のセッションを自動的に検出し、[ライブ] タブにソース スイッチャーを表示します。
Claude Code と OpenCode の間をクリックすると、ダッシュボードから離れることなく、各セッションのリアルタイム ビュー (ツール呼び出し、プロンプト、モデル名、インテント バッジ、タイミング) が表示されます。
OpenCode データはローカル SQLite データベースから直接読み取られます。構成は必要ありません。
claudestat には、使用状況統計をクエリするための 7 つのツールを備えた MCP サーバーが含まれています。一度登録してください:
クロード mcp 追加 claudestat -s ユーザー -- claudestat-mcp
次に、クロードに「私の割り当て状況は何ですか?」と尋ねます。 、「最新のセッションを表示」、「コスト別トップ 5 ツール」。
構成は ~/.claudestat/config.json に保存されます。 claudestat config を使用して表示するか、ファイルを直接編集します。
claudestat config --kill-switch true --threshold 90
claudestat config --plan max5
claudestat config --alerts false
完全な構成リファレンス →
クロード

コードイベント
│
▼
フックスクリプト (~/.claudestat/hooks/event.js)
│ デーモンに JSON を POST
▼
デーモン (ローカルホスト:7337)
│ SQLite にイベントを保存します
│ JSONLトークンデータで強化
│ パターンアナライザーを実行します
▼
ダッシュボード (React + Vite、自動更新)
│
▼
ライブですべてがわかります
トラブルシューティング
claudestat の開始が約 5 秒間ハングする
通常 — require('express') は最初のロード時に数秒かかります。デーモンは実行中です。 「デーモンが開始されました」という確認を待ちます。
フックが起動しない / ダッシュボードにイベントが表示されない
claudestat Doctor を実行します。すべてのコンポーネントをチェックし、正確な修正コマンドを出力します。
インストール後にclaudestatコマンドが見つかりません
NVM を使用している場合、バイナリが間違ったノード バージョンを指している可能性があります。
nvm はデフォルトを使用します && npm install -g @statforge/claudestat && ハッシュ -r claudestat
キルスイッチが新しいセッションをブロックしています
claudestat config --kill-switch false で無効にするか、5 時間のクォータ ウィンドウがリセットされるまで待ちます。
レート制限に近づいています
デーモンは 60 秒ごとにクォータをポーリングし、70%、85%、および 95% で警告を記録します。 claudestat status でいつでも確認できます。
複数のプロジェクトで作業する
claudestat はすべてのプロジェクトを自動的に追跡します。 「プロジェクト」タブでは、作業ディレクトリごとにセッションをグループ化します。
ダッシュボードにはすべてのセッションでコスト 0 / 0.00 ドルが表示されます
トークン データは、フック イベントからではなく、Claude Code の JSONL ファイルから取得されます。 Claude Code が JSONL ログを書き込んでいることを確認してください。~/.claude/projects/ で .jsonl ファイルを確認してください。ディレクトリが空の場合、クロード コードのログ記録が有効になっていない可能性があります。
ターミナルを閉じるとデーモンが停止する
シェル セッションを超えて存続するには、デーモンを nohup で起動する必要があります。
nohup claudestat の開始と
または、システム サービス (macOS では launchd、Linux では systemd) をインストールする claudestat セットアップを使用します。
claudestat エクスポートは空の出力を生成します
セッションが表示されない場合は、デーモンが実行されていない可能性があります。

クロード・コードのセッションを呼び出してください。 claudestat のステータスを確認し、 claudestat start で再起動します。履歴データのみ (デーモンが実行されていない場合) の場合、エクスポートは引き続きローカル SQLite データベースから読み取ります。そのため、デーモンの実行中にキャプチャされた過去のセッションは常に利用できます。
ループ検出器の起動が多すぎる/十分ではない
しきい値とウィンドウを調整します。
claudestat config --loop-threshold 5 # デフォルト: 8 コール
claudestat config --loop-window 90 # デフォルト: 120 秒
MCP サーバーが応答しません
デーモンを再起動し ( claudestat restart )、登録されていることを確認します。
クロードMCPリスト
リストされていない場合は、再実行します: claude mcp add claudestat -s user -- claudestat-mcp
OpenCode セッションが表示されない
claudestat は ~/.local/share/opencode/opencode.db から OpenCode データを読み取ります。ファイルが存在しない場合は、OpenCode がまだ実行されていないか、システム上の別のデータ パスが使用されています。 opencode を少なくとも 1 回実行して初期化します。
Node.js の起動時の実験的な SQLite 警告
予想通り — node:sqlite はノード 22 で実験的です。警告は自動的に抑制されます。繰り返し表示される場合は、Node.js 22 以降 (node --version) を実行していることを確認してください。
クローデスタットとは何ですか?告訴とどう違うのですか？
claudestat はクロード コードのリアルタイム モニターであり、ログ リーダーではありません。起動時にすべてのツール呼び出しにフックし、トークンの使用量とコストを追跡します。

[切り捨てられた]

## Original Extract

Real-time execution trace and cost intelligence for Claude Code - DeibyGS/claudestat

GitHub - DeibyGS/claudestat: Real-time execution trace and cost intelligence for Claude Code · GitHub
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
DeibyGS
/
claudestat
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
152 Commits 152 Commits .github .github assets assets dashboard dashboard docs docs hooks hooks scripts scripts src src tests tests .gitignore .gitignore .npmignore .npmignore .nvmrc .nvmrc CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md llms.txt llms.txt package-lock.json package-lock.json package.json package.json run-tests.sh run-tests.sh tsconfig.json tsconfig.json View all files Repository files navigation
Live AI coding monitor — real-time trace, quota guard, and MCP server
Most tools read your logs after a session ends. claudestat hooks into every event as it fires.
See what your AI is spending right now, get alerted before you hit your quota, and ask Claude about its own usage — from inside the terminal.
Works with Claude Code and OpenCode . Zero cloud dependencies. Pure Node.js. Runs on macOS, Linux, and Windows.
Installation •
Quick Start •
Commands •
Dashboard •
FAQ •
Contributing
Live dashboard · terminal trace · quota guard — all running in real time
Tools like ccusage are great for reviewing history. claudestat is for while you're coding .
It taps into Claude Code's hook system to capture every event the moment it fires, stores everything locally in SQLite, and gives you a live dashboard, quota alerts, and an MCP server — not just a report.
If claudestat is useful, give it a ⭐ — it helps other developers find it.
claudestat ships an MCP server. Once registered, you can ask Claude Code questions about its own usage — without leaving the terminal.
claude mcp add claudestat -s user -- claudestat-mcp
Then just ask:
> How much did I spend this week?
> What are my top 5 tools by cost?
> Break down my usage by model
> What's my cost projection for next month?
Full MCP reference →
Live tool trace — every call with duration and token cost as it runs
OpenCode support — same live dashboard for OpenCode sessions (tool calls, prompts, model, intent)
Quota guard — alerts at 70%, 85%, 95%; optional kill switch blocks new sessions at X%
Loop detector — flags context thrashing with estimated waste cost
Top tools — know which tools eat most of your budget; expandable "Other" row reveals tools beyond top 10
Cost projection — linear regression with trend, confidence intervals, R²
Billing blocks — track 5-hour billing windows and cumulative spend per block ( claudestat blocks )
Session sharing — export any session summary as formatted ASCII or JSON ( claudestat share )
52-week activity heatmap — visualize your coding activity across the full year in Analytics
Period-over-period deltas — Analytics KPIs show ↑↓% trend vs. the previous period
Projects search & filter — find projects instantly; filter by activity, cost, or efficiency threshold
Web dashboard — 6 tabs: Live, History, Projects, Analytics, Top, System
MCP server — 7 tools so Claude can answer questions about its own usage
Weekly insights — pattern analysis with actionable tips
Multi-source — switch between Claude Code and OpenCode sessions in one click
Source filter — filter KPIs, charts, and tool rankings by Claude Code / OpenCode across all tabs
Multi-tool coordination — live intent panel with real collision detection (same file edited by CC and OC)
npm install -g @statforge/claudestat && claudestat setup
open http://localhost:7337
Start a Claude Code session and watch the events flow in. That's it.
Node.js >= 22 (required for node:sqlite )
Claude Code installed ( npm install -g @anthropic-ai/claude-code )
npm install -g @statforge/claudestat && claudestat setup
claudestat setup installs the Claude Code hooks and registers the daemon as a system service (launchd on macOS, systemd on Linux) — no sudo required. The daemon starts automatically whenever you log in.
Using NVM? Make sure you're on your default Node version:
nvm use default && npm install -g @statforge/claudestat && claudestat setup
Restart Claude Code after setup so the hooks take effect.
npm install -g @statforge/claudestat
claudestat install # installs hooks into Claude Code
claudestat start # start the daemon manually
Commands
Command
Description
claudestat setup
One-command setup: install hooks + register daemon as system service
claudestat setup --uninstall
Remove hooks and system service
claudestat start / stop / restart
Manage the background daemon
claudestat install / uninstall
Install or remove Claude Code hooks
claudestat watch
Live terminal trace view
claudestat status
Show quota, cost, and burn rate
claudestat top
Rank tools by cost, call count, or duration
claudestat weekly
Weekly usage summary
claudestat insights
Deep usage insights: cost, cache, efficiency, models
claudestat project
Cost projection with linear regression
claudestat config
View or edit configuration
claudestat doctor
Check installation health and diagnose issues
claudestat blocks
Show 5-hour billing block history
claudestat share [session-id]
Export session summary as ASCII or JSON (use --copy to copy to clipboard)
claudestat export [format]
Export session data to JSON or CSV
claudestat roast
Sarcastic usage analysis
claudestat version
Show version and check for updates
Full command reference with output examples →
The dashboard lives at http://localhost:7337 and has six tabs: Live (real-time trace with source badge, expandable Bash commands, and last-task subtitle per source), History (sessions by date with day selector 7/14/30/90d, merged-session badge, search, cost filter, and compare panel), Projects (grid with weekly heatmap, search input, and filters for active, high-cost, or low-efficiency projects), Analytics (spend + tokens + hours + efficiency KPIs with period-over-period ↑↓% deltas, stacked token chart by input/output/cache, 52-week activity heatmap, source filter, weekly AI reports), Top (tool rankings by cost/count/duration with cost projection and expandable "Other" row for tools beyond top 10), and System (hooks, agents, skills, workflows, context file limits, work mode distribution, OpenCode config, memory files with truncation warning, claudestat config).
When you run OpenCode alongside Claude Code, claudestat automatically detects both sessions and shows a source switcher in the Live tab.
Click between Claude Code and OpenCode to see each session's real-time view — tool calls, prompts, model name, intent badges, and timing — without leaving the dashboard.
OpenCode data is read directly from its local SQLite database — no configuration required.
claudestat includes an MCP server with 7 tools for querying usage stats. Register once:
claude mcp add claudestat -s user -- claudestat-mcp
Then ask Claude: "What's my quota status?" , "Show me my latest session" , "Top 5 tools by cost" .
Config is stored at ~/.claudestat/config.json . View it with claudestat config or edit the file directly.
claudestat config --kill-switch true --threshold 90
claudestat config --plan max5
claudestat config --alerts false
Full configuration reference →
Claude Code event
│
▼
Hook script (~/.claudestat/hooks/event.js)
│ POST JSON to daemon
▼
Daemon (localhost:7337)
│ stores events in SQLite
│ enriches with JSONL token data
│ runs pattern analyzer
▼
Dashboard (React + Vite, auto-refreshes)
│
▼
You see everything — live
Troubleshooting
claudestat start hangs for ~5 seconds
Normal — require('express') takes a few seconds on first load. The daemon is running; wait for the "Daemon started" confirmation.
Hooks are not firing / dashboard shows no events
Run claudestat doctor — it checks every component and prints the exact fix command.
claudestat command not found after install
If using NVM, the binary may point to the wrong Node version:
nvm use default && npm install -g @statforge/claudestat && hash -r claudestat
Kill switch is blocking new sessions
Disable with claudestat config --kill-switch false , or wait for the 5h quota window to reset.
Approaching rate limit
The daemon polls quota every 60s and logs warnings at 70%, 85%, and 95%. Check anytime with claudestat status .
Working with multiple projects
claudestat tracks every project automatically. The Projects tab groups sessions by working directory.
Dashboard shows 0 cost / $0.00 for all sessions
Token data comes from Claude Code's JSONL files, not from hook events. Make sure Claude Code is writing JSONL logs — check ~/.claude/projects/ for .jsonl files. If the directory is empty, Claude Code may not have logging enabled.
Daemon stops after terminal closes
The daemon must be started with nohup to persist beyond the shell session:
nohup claudestat start &
Or use claudestat setup which installs a system service (launchd on macOS, systemd on Linux).
claudestat export produces empty output
If no sessions appear, the daemon may not have been running during your Claude Code sessions. Check claudestat status and restart with claudestat start . For historical data only (without a running daemon), export still reads from the local SQLite database — so past sessions captured while the daemon was running are always available.
Loop detector fires too often / not enough
Adjust the threshold and window:
claudestat config --loop-threshold 5 # default: 8 calls
claudestat config --loop-window 90 # default: 120 seconds
MCP server not responding
Restart the daemon ( claudestat restart ) and verify it's registered:
claude mcp list
If not listed, re-run: claude mcp add claudestat -s user -- claudestat-mcp
OpenCode sessions not appearing
claudestat reads OpenCode data from ~/.local/share/opencode/opencode.db . If the file does not exist, OpenCode has not run yet or uses a different data path on your system. Run opencode at least once to initialize it.
Node.js experimental SQLite warning on startup
Expected — node:sqlite is experimental in Node 22. The warning is suppressed automatically. If you see it repeatedly, ensure you are running Node.js 22 or later ( node --version ).
What is claudestat? How is it different from ccusage?
claudestat is a real-time monitor for Claude Code — not a log reader. It hooks into every tool call as it fires, tracks token usage and cost li

[truncated]
