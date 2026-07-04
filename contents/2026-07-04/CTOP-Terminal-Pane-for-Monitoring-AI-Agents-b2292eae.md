---
source: "https://github.com/aakashadesara/ctop"
hn_url: "https://news.ycombinator.com/item?id=48785567"
title: "CTOP – Terminal Pane for Monitoring AI Agents"
article_title: "GitHub - aakashadesara/ctop: CTOP - Interactive Process Viewer for AI Coding Agents · GitHub"
author: "aakashadesara"
captured_at: "2026-07-04T14:40:28Z"
capture_tool: "hn-digest"
hn_id: 48785567
score: 3
comments: 3
posted_at: "2026-07-04T14:14:35Z"
tags:
  - hacker-news
  - translated
---

# CTOP – Terminal Pane for Monitoring AI Agents

- HN: [48785567](https://news.ycombinator.com/item?id=48785567)
- Source: [github.com](https://github.com/aakashadesara/ctop)
- Score: 3
- Comments: 3
- Posted: 2026-07-04T14:14:35Z

## Translation

タイトル: CTOP – AI エージェントを監視するためのターミナル ペイン
記事タイトル: GitHub - aakashadesara/ctop: CTOP - AI コーディング エージェント用対話型プロセス ビューアー · GitHub
説明: CTOP - AI コーディング エージェント用の対話型プロセス ビューアー - aakashadesara/ctop

記事本文:
GitHub - aakashadesara/ctop: CTOP - AI コーディング エージェント用の対話型プロセス ビューアー · GitHub
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
アーカシャデサラ
/
最高
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店 Ta

gs ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
178 コミット 178 コミット .github/ workflows .github/ workflows 資産 アセット ドキュメント ドキュメント サンプル/ プラグイン サンプル/ プラグイン スクリプト スクリプト スキル/ ctop スキル/ ctop src src test test .gitignore .gitignore ライセンス ライセンス README.md README.md claude-manager claude-manager package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
CTOP — AI エージェント端末操作パネル
AI コーディング エージェントの htop。 Claude Code、Codex CLI、OpenCode、Devin セッション (CPU、メモリ、トークン、コンテキスト ウィンドウ、コスト、ブランチ) を単一のターミナル ペインから監視します。
マルチエージェント監視 — Claude Code + Codex CLI + OpenCode + Devin、リアルタイム CPU/メモリ/ステータス
コンテキスト ウィンドウの追跡 — 入力、キャッシュ、出力、および空きセグメントを含むビジュアル バー
コストの見積もり — セッションごとの API コストと総 API コスト (Claude + OpenAI の価格設定)
トークン波形 — トークンアクティビティパルスを示すリアルタイムのスパークライン
2 つの表示モード — リスト ビュー (テーブル) とペイン ビュー (カード グリッド)、P で切り替え
ライブログテーリング - 分割ペインでの会話のストリーム ( L )
並べ替え、フィルター、検索 - CPU、メモリ、コンテキスト、ブランチ、モデル、またはフルテキストによる ( F )
ダッシュボードと履歴 — 集計統計 ( d )、24 時間の使用状況グラフ ( H )
プロセス制御 - セッションの強制終了 (正常または強制)、一括複数選択によるクローズ、プロジェクト ディレクトリへのクイック ジャンプ
デスクトップ通知 - セッションが完了すると通知を受け取ります
5 つのカラーテーマ — デフォルト、ミニマル、ドラキュラ、ソラリゼーション、モノカイ (+ カスタム)
プラグイン システム — ~/.ctop/plugins/ 経由でカスタム列を拡張します。
コンパクションとレート制限の検出 - コンパクション イベントとクォータ使用量にフラグを立てます
エージェントの CLI モード - ctop ls 、 ctop whoami 、 ctop alters など (「CLI モード」を参照)
#自作
brew Tap aakashadesara/ctop && brew install ctop-claude
#npm

npm install -g ctop-claude
# npx (インストールなし)
npx ctop-クロード
# ソースから
git clone https://github.com/aakashadesara/ctop.git
chmod +x ctop/クロードマネージャー
ln -s " $( pwd ) /ctop/claude-manager " /usr/local/bin/ctop
次に、 ctop を実行します。実行中のエージェントがない場合は、空の状態が表示されます。Claude Code、Codex、OpenCode、または Devin セッションを開始すると、次回の更新時に表示されます。
キー
アクション
j / k または ↑ / ↓
ナビゲート
h / l または ← / →
ナビゲート (ペインモード)
グラム/グラム
最初/最後にジャンプ
P
リスト/ペインビューの切り替え
p
セッションの固定/固定解除 (上部に固定)
宇宙
セッションのマーク/マーク解除 (複数選択)
Shift+↑ / ↓ または V
マークされた範囲を拡張/開始する
ある
表示/クリアをすべて選択
秒/秒
サイクルソート/リバース
/
フィルター
F
全文検索の会話
d
ダッシュボードの切り替え
L
ログペインの切り替え
H
24 時間の履歴を切り替えます
W
タイムラインビュー
T
サイクルカラーテーマ
× / ×
Kill (SIGTERM / SIGKILL) — 行がマークされている場合は一括処理
K
すべてのエージェントを殺害する
あ
停止した/死亡したエージェントをすべてキルする
o/e/t
Finder/エディタ/ターミナルでディレクトリを開きます
n
通知の切り替え
?
ヘルプ
Esc
選択をクリア (またはフィルター/検索)
q
やめる
マウス: クリックして選択し、スクロールして移動し、★ 枠をクリックして固定/固定解除し、Shift キーを押しながらクリックしてマークします (ベストエフォート)。
関心のあるセッションを常に表示しておきます。 pを押します（または、連続した★をクリックします）
ガター、またはピン フッター ボタン) を使用して、カーソルの下にセッションを固定します。ジャンプします。
黄色の★ 上部に固定されたセクションがあり、並べ替えや種類に関係なくそこに留まります。
フィルター。固定は、リスト、グループ、およびペイン ビューで機能します。ピンはセッションごとにキー設定されます
ID (pid ではない) なので、更新、セッションの再起動、終了後も存続します。
ctop — ~/.ctop/pins.json に永続化されます。固定を解除するには、もう一度 p を押します。
複数のセッションをマークし、それらを同時に実行します。スペースを押してセッションをマークします
カーソルの下に移動するか、Shift キーを押しながら ↑ キーを押します。

/ ↓ 範囲を拡張します。押す
V は vim スタイルの範囲モード (次に移動して拡張)、a は表示されているすべてを選択します。
セッションがマークされている場合、x / X は確認プロンプトの後でセット全体を閉じます。
Esc キーで選択をクリアします。リスト、ペイン、およびグループ ビューで機能します。
Shift+クリック 注: 多くのターミナル (ターミナル.app、iTerm2、GNOME ターミナルなど)
Shift キーを押しながらクリックして独自のテキスト選択を予約し、それを他のユーザーに転送しないでください。
アプリなので、Shift キーを押しながらマークをクリックするのがベストエフォートです。キーボードのパス
( Space / Shift + ↑ / ↓ / V ) はどこでも機能します。
このリポジトリには、自己完結型の ctop スキルが同梱されています。これをクロード コードにドロップすると、エージェントは ctop を呼び出すタイミングと方法を学習できます。
# プロジェクトごと
mkdir -p .claude/skills && cp -r skill/ctop .claude/skills/
# またはユーザー全体
mkdir -p ~ /.claude/skills && cp -r skill/ctop ~ /.claude/skills/
インストールしたら、「他にどのようなエージェントを実行しているか」、「セッションのコストはいくらか」、「コンテキストは圧縮されようとしているか」などのクロード コード セッションを尋ねます。エージェントは自動的に ctop に到達します。
SKILL.md — トリガーシート+共通パターン
Reference.md — コマンドごとの完全な仕様
Examples.md — コピー＆ペースト可能なレシピ
CLI モード (エージェントおよびスクリプト用)
引数なしの ctop は対話型 TUI を開始します。 ctop <subcommand> はワンショット クエリを実行して終了するため、AI エージェントは別の端末から自身のセッションと姉妹セッションをイントロスペクトできます。
ctop ls # 実行中のすべてのエージェントのテーブル
ctop ls --json # 同じ、マシン解析可能
ctop ls --agent claude # バックエンドでフィルターする
ctop ls --cwd ~ /code/myproj # ディレクトリでフィルタリングする
ctop get < pid > --json # 1 つのセッションの詳細
ctop log < pid > --tail 20 # 最新 20 件の会話メッセージ
ctop search " TODO " --json # セッション間の全文検索
ctop diff < pid > # セッションの cwd の Git diff
ctop stats --json # 総コスト / トークン / クー

nts
ctop whoami # 現在どのセッションにいるかを検出します
ctop whoami --pid-only # PID のみ、スクリプト用
ctopアラート # ローコンテキスト/アイドル/ゴーストの警告
ctopアラート --severitycritical # クリティカルレベルのアラートのみ
ctop kill < pid > # SIGTERM (自分のユーザーである必要があります)
ctop kill < pid > --force # SIGKILL
ctop notification " title " " message " # デスクトップ通知
whoami は、$CTOP_PID →parent-PID walk → $PWD match を介して、matchConfidence ラベル (exact | ppid | cwd-guess | none ) を使用して呼び出しセッションを検出するため、エージェントは回答をどの程度信頼すべきかを知ることができます。
読み取りツールは、ユーザーがディスクから読み取ることができるデータを表示します。 kill は、シグナルを送信する前に uid の所有権とエージェント セッションのチェックを強制します。すべてを kill することはありません。
# 圧縮しようとしているセッションを検索する
ctop ls --json | jq ' .[] | select(.contextPct != null および .contextPct < 20) '
# 自己認識コンパクション (フック)
[ " $( ctop whoami --json | jq -r .session.contextPct ) " -lt 15 ] && \
echo " コンテキストが低い - /compact を検討してください "
# ゴーストセッションをクリーンアップする
ctop アラート --json | jq -r ' .[] | select(.kind=="ゴースト") | .pid' | \
xargs -I {} ctop kill {} --force
構成
ctop --refresh 3 # 3 秒ごとに更新します
ctop --context-limit 128000 # コンテキスト ウィンドウを 128k に設定します
ctop --pane # ペインビューで開始します
設定ファイル ( ~/.ctoprc )
{
"リフレッシュ間隔" : 5000 、
"contextLimit" : 200000 、
"defaultView" : " リスト " ,
"テーマ" : "デフォルト" ,
"contextBarStyle" : " ブロック " ,
"通知" : { "有効" : true 、 "minDuration" : 30 }
}
CLI フラグは設定ファイルの値をオーバーライドします。
ps (Windows 上の PowerShell) からプロセス情報を読み取り、 lsof 経由で作業ディレクトリを解決し、ローカル JSONL ファイル (Claude の場合は ~/.claude/projects/、Codex の場合は ~/.codex/sessions/) および SQLite データベース (~/.local/share/opencode/ の場合) からのセッション メタデータで各プロセスを強化します。

OpenCode、Devin の場合は ~/.local/share/devin/cli/)。ネットワーク呼び出しや外部依存関係はありません。
カスタム列を使用して拡張します。 ~/.ctop/plugins/ に .js ファイルを作成します。
モジュール。エクスポート = {
名前: 'my-plugin' 、
列: {
ヘッダー: 'CUSTOM' 、
幅：10、
getValue : ( proc ) => proc 。 CWD？ 「はい」:「いいえ」、
} 、
} ;
詳細については、examples/plugins/ を参照してください。
macOS、Linux、または Windows — Windows はプロセス検出に PowerShell を使用します。 CWD 解像度は macOS/Linux よりも制限されています。
Claude Code 、Codex CLI 、OpenCode 、および/または Devin (ターミナル) 実行セッション
OpenCode および Devin セッション読み取り用の PATH 上の sqlite3 (macOS に組み込まれており、Linux では apt / brew 経由で利用可能)
PRの方も大歓迎！開発するにはフォーク、クローン、./claude-manager を実行し、テストするには npm test を実行します。大きな変更がある場合は、まず問題をオープンしてください。
CTOP - AI コーディング エージェント用の対話型プロセス ビューアー
読み込み中にエラーが発生しました。このページをリロードしてください。
8
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

CTOP - Interactive Process Viewer for AI Coding Agents - aakashadesara/ctop

GitHub - aakashadesara/ctop: CTOP - Interactive Process Viewer for AI Coding Agents · GitHub
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
aakashadesara
/
ctop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
178 Commits 178 Commits .github/ workflows .github/ workflows assets assets docs docs examples/ plugins examples/ plugins scripts scripts skills/ ctop skills/ ctop src src test test .gitignore .gitignore LICENSE LICENSE README.md README.md claude-manager claude-manager package.json package.json View all files Repository files navigation
CTOP — AI Agent Terminal Operations Panel
htop for your AI coding agents. Monitor Claude Code, Codex CLI, OpenCode, and Devin sessions — CPU, memory, tokens, context window, costs, branches — from a single terminal pane.
Multi-agent monitoring — Claude Code + Codex CLI + OpenCode + Devin, real-time CPU/memory/status
Context window tracking — visual bar with input, cache, output, and free segments
Cost estimation — per-session and aggregate API cost (Claude + OpenAI pricing)
Token waveform — real-time sparkline showing token activity pulse
Two view modes — list view (table) and pane view (card grid), toggle with P
Live log tailing — stream conversation in a split pane ( L )
Sort, filter, search — by CPU, memory, context, branch, model, or full-text ( F )
Dashboard & history — aggregate stats ( d ), 24-hour usage charts ( H )
Process control — kill sessions (graceful or force), bulk multi-select close, quick-jump to project dir
Desktop notifications — get notified when sessions complete
5 color themes — default, minimal, dracula, solarized, monokai (+ custom)
Plugin system — extend with custom columns via ~/.ctop/plugins/
Compaction & rate limit detection — flags compaction events and quota usage
CLI mode for agents — ctop ls , ctop whoami , ctop alerts , … (see CLI mode )
# Homebrew
brew tap aakashadesara/ctop && brew install ctop-claude
# npm
npm install -g ctop-claude
# npx (no install)
npx ctop-claude
# From source
git clone https://github.com/aakashadesara/ctop.git
chmod +x ctop/claude-manager
ln -s " $( pwd ) /ctop/claude-manager " /usr/local/bin/ctop
Then run ctop . If no agents are running, you'll see an empty state — start a Claude Code, Codex, OpenCode, or Devin session and it'll appear on the next refresh.
Key
Action
j / k or ↑ / ↓
Navigate
h / l or ← / →
Navigate (pane mode)
g / G
Jump to first / last
P
Toggle list / pane view
p
Pin / unpin session (sticks it to the top)
Space
Mark / unmark session (multi-select)
Shift+↑ / ↓ or V
Extend / start a marked range
a
Select all visible / clear
s / S
Cycle sort / reverse
/
Filter
F
Full-text search conversations
d
Toggle dashboard
L
Toggle log pane
H
Toggle 24-hour history
W
Timeline view
T
Cycle color theme
x / X
Kill (SIGTERM / SIGKILL) — bulk if rows are marked
K
Kill ALL agents
A
Kill ALL stopped/dead agents
o / e / t
Open dir in Finder / editor / terminal
n
Toggle notifications
?
Help
Esc
Clear selection (or filter / search)
q
Quit
Mouse: click to select, scroll to navigate, click the ★ gutter to pin/unpin, Shift +click to mark (best-effort).
Keep the sessions you care about in view. Press p (or click the ★ in a row's
gutter, or the Pin footer button) to pin the session under the cursor — it jumps
to a yellow ★ Pinned section at the top and stays there regardless of sort or
filter. Pinning works in list, group, and pane views. Pins are keyed by session
identity (not pid), so they survive refreshes, the session restarting, and quitting
ctop — persisted to ~/.ctop/pins.json . Press p again to unpin.
Mark several sessions and act on them at once. Press Space to mark the session
under the cursor, or hold Shift while pressing ↑ / ↓ to extend a range; press
V for vim-style range mode (then move to extend) and a to select all visible.
With sessions marked, x / X close the whole set after a confirmation prompt;
Esc clears the selection. Works in list, pane, and group views.
Shift+click note: many terminals (Terminal.app, iTerm2, GNOME Terminal, …)
reserve Shift +click for their own text selection and never forward it to the
app, so Shift +click marking is best-effort. The keyboard path
( Space / Shift + ↑ / ↓ / V ) works everywhere.
A self-contained ctop skill ships in this repo. Drop it into Claude Code so any agent learns when and how to call ctop :
# Per-project
mkdir -p .claude/skills && cp -r skills/ctop .claude/skills/
# Or user-wide
mkdir -p ~ /.claude/skills && cp -r skills/ctop ~ /.claude/skills/
Once installed, ask any Claude Code session things like "what other agents am I running" , "how much have my sessions cost" , "is my context about to compact" — the agent will reach for ctop automatically.
SKILL.md — trigger sheet + common patterns
reference.md — full per-command spec
examples.md — copy-pasteable recipes
CLI mode (for agents and scripts)
ctop with no args starts the interactive TUI. ctop <subcommand> runs a one-shot query and exits, so AI agents can introspect their own sessions and sister sessions from another terminal.
ctop ls # Table of every running agent
ctop ls --json # Same, machine-parseable
ctop ls --agent claude # Filter by backend
ctop ls --cwd ~ /code/myproj # Filter by directory
ctop get < pid > --json # Full detail on one session
ctop log < pid > --tail 20 # Last 20 conversation messages
ctop search " TODO " --json # Full-text search across sessions
ctop diff < pid > # Git diff for the session's cwd
ctop stats --json # Aggregate cost / tokens / counts
ctop whoami # Detect which session you're in
ctop whoami --pid-only # PID only, for scripting
ctop alerts # Low-context / idle / ghost warnings
ctop alerts --severity critical # Only critical-level alerts
ctop kill < pid > # SIGTERM (must be your own user)
ctop kill < pid > --force # SIGKILL
ctop notify " title " " message " # Desktop notification
whoami detects the calling session via $CTOP_PID → parent-PID walk → $PWD match, with a matchConfidence label ( exact | ppid | cwd-guess | none ) so agents know how much to trust the answer.
Read tools surface data that the user could read off disk anyway. kill enforces uid ownership and an agent-session check before sending the signal — there is no kill-all.
# Find sessions about to compact
ctop ls --json | jq ' .[] | select(.contextPct != null and .contextPct < 20) '
# Self-aware compaction (hook)
[ " $( ctop whoami --json | jq -r .session.contextPct ) " -lt 15 ] && \
echo " context low — consider /compact "
# Clean up ghost sessions
ctop alerts --json | jq -r ' .[] | select(.kind=="ghost") | .pid ' | \
xargs -I {} ctop kill {} --force
Configuration
ctop --refresh 3 # Refresh every 3 seconds
ctop --context-limit 128000 # Set context window to 128k
ctop --pane # Start in pane view
Config file ( ~/.ctoprc )
{
"refreshInterval" : 5000 ,
"contextLimit" : 200000 ,
"defaultView" : " list " ,
"theme" : " default " ,
"contextBarStyle" : " block " ,
"notifications" : { "enabled" : true , "minDuration" : 30 }
}
CLI flags override config file values.
Reads process info from ps (PowerShell on Windows), resolves working directories via lsof , and enriches each process with session metadata from local JSONL files ( ~/.claude/projects/ for Claude, ~/.codex/sessions/ for Codex) and SQLite databases ( ~/.local/share/opencode/ for OpenCode, ~/.local/share/devin/cli/ for Devin). No network calls, no external dependencies.
Extend with custom columns. Create .js files in ~/.ctop/plugins/ :
module . exports = {
name : 'my-plugin' ,
column : {
header : 'CUSTOM' ,
width : 10 ,
getValue : ( proc ) => proc . cwd ? 'yes' : 'no' ,
} ,
} ;
See examples/plugins/ for more.
macOS, Linux, or Windows — Windows uses PowerShell for process detection; CWD resolution is more limited than macOS/Linux.
Claude Code , Codex CLI , OpenCode , and/or Devin (terminal) running sessions
sqlite3 on PATH for OpenCode and Devin session reads (built-in on macOS; available via apt / brew on Linux)
PRs welcome! Fork, clone, run ./claude-manager to develop, npm test to test. Open an issue first for large changes.
CTOP - Interactive Process Viewer for AI Coding Agents
There was an error while loading. Please reload this page .
8
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
