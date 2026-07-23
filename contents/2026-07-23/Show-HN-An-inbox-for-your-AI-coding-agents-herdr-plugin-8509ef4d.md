---
source: "https://github.com/douglascorrea/herdr-agent-inbox"
hn_url: "https://news.ycombinator.com/item?id=49027154"
title: "Show HN: An inbox for your AI coding agents (herdr plugin)"
article_title: "GitHub - douglascorrea/herdr-agent-inbox: An inbox for your coding agents in herdr — session titles, settle/mark-unread triage, runtimes, workspace rollups, and resumable chat history · GitHub"
author: "douglascorrea"
captured_at: "2026-07-23T20:12:51Z"
capture_tool: "hn-digest"
hn_id: 49027154
score: 1
comments: 0
posted_at: "2026-07-23T19:54:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: An inbox for your AI coding agents (herdr plugin)

- HN: [49027154](https://news.ycombinator.com/item?id=49027154)
- Source: [github.com](https://github.com/douglascorrea/herdr-agent-inbox)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T19:54:00Z

## Translation

タイトル: HN の表示: AI コーディング エージェントの受信箱 (herdr プラグイン)
記事のタイトル: GitHub - douglascorrea/herdr-agent-inbox: herdr のコーディング エージェント用の受信箱 — セッション タイトル、解決/未読マーク付けトリアージ、ランタイム、ワークスペース ロールアップ、再開可能なチャット履歴 · GitHub
説明: herdr のコーディング エージェント用の受信箱 - セッション タイトル、解決/未読トリアージ、ランタイム、ワークスペース ロールアップ、再開可能なチャット履歴 - douglascorrea/herdr-agent-inbox

記事本文:
GitHub - douglascorrea/herdr-agent-inbox: herdr のコーディング エージェント用の受信箱 — セッション タイトル、解決/未読マーク付けトリアージ、ランタイム、ワークスペース ロールアップ、再開可能なチャット履歴 · GitHub
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
ダグラス

コレア
/
herdr-エージェントの受信箱
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット アセット アセット スクリプト スクリプト .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md events.py events.py daemon.py daemon.py herdr-plugin.toml herdr-plugin.toml inbox_tui.py inbox_tui.py すべてのファイルを表示 リポジトリ ファイル ナビゲーション
コーディング エージェントの受信箱。セッションタイトル、決済/未読マーク
トリアージ、実行時間、ワークスペースのロールアップ、ChatGPT スタイルの再開可能なチャット
履歴 — herdr で実行されているすべてのエージェントについて。
T3 コードの「決済」ワークフローからインスピレーションを得た:
「受信箱のようなものだと考えてください。スレッドを読み終わったら、
「決済」ボタンを押すと下にスライドします。これにより、より「仕上げる」ことができるようになりました。
セッション タイトル - すべてのエージェント ペインには、セッション タイトルから派生した実際のタイトルが付けられます。
クロードや
ディレクトリ名。エージェントのネイティブのトランスクリプト (claude / pi / codex) を読み取ります。
オプションで、設定した LLM コマンドによって要約されます。
受信トレイのトリアージ — エージェントのサイドバーがアテンション キューになります: ブロック済み →
終了・未見→稼働中→アイドル→一番下に落ち着いています。決済する
終わったらエージェント。再び動作すると自動的に再開します。
1 つを未読としてマークし、一番上に固定します。
蓄積されるチャット履歴 - 閉じられた会話は下にリストされたままになります
ワークスペース/ディレクトリ (およびグローバル履歴ブラウザ内)、および
claude --resume / pi --session / codex raise を使用してセッション参照を再度開きます。
実行時間 - 各セッションの存続期間とその経過時間
現在の状態では、サイドバーに表示されます。
ワークスペースのロールアップ - それぞれ

rkspace 行にはエージェントが状態別に表示されます
( !1 ●2 ▸1 ⚑3 ) に最長のビジー スティントを加えたものです。
サイドバーのサイズ変更 — キーボードから herdr のサイドバーを拡大/縮小します (および
herdr のネイティブのマウス ドラッグを機能させ続けます)。
すべては herdr のパブリック プラグイン サーフェスを通じて実行されます。フォークやパッチは必要ありません。
stdlib Python のみ。
プレフィックス + i で受信トレイが開きます。 g は 4 つのビューを回転し、h は履歴を開きます。
コンパクト — チャットごとに 1 行、ドミナント状態フラグを含むワークスペース ヘッダー:
履歴 — すべての非公開チャット、最新のものから順に表示されます。 ↩ は再開可能を意味します:
🔴 / !ブロックされています · 🔵 / ● 完了、未表示 · 🟡 / ▸ 作業中 ·
🟢 / ○ アイドル · 🏁 / ⚑ 解決済み · ⚫ 非公開チャット — herdr のマッチング
herdr テーマの独自のサイドバーの色言語 (赤/青緑/黄/緑)
まさにパレット。
PATH に herdr ≥ 0.7.0 および python3 (3.11 以降を推奨) が必要です。
herdr プラグインのインストール douglascorrea/herdr-agent-inbox
デーモンはセッションとともに開始されます (そして、ウォッチドッグがそれを復活させます)。それから配線します
ディスプレイとキー:
サイドバーの行 ( ~/.config/herdr/config.toml )
プラグインはトークンを報告します。サイドバーの設定によって、何がレンダリングされるかが決まります。
[うい。サイドバー。エージェント]
行 = [
[ " state_icon " 、 " $title " ]、
[{ token = " $flag " 、 fg = " #f1fa8c " 、太字 = true }、 " エージェント " 、 " ワークスペース " 、 { token = " $age " 、 dim = true }]、
】
[うい。サイドバー。スペース]
行 = [
[ " state_icon " , " workspace " , " $agents " , { token = " $busy " , dim = true }],
[ " ブランチ " 、 " git_status " ]、
】
利用可能なペイントークン: $title 、 $age (セッション実行時間)、 $since
(現在の状態の時間)、$flag (⚑確定 / ● 未読)。ワークスペーストークン:
$agents (状態カウント)、$busy (最長の作業スティント)。スタイルカスタム
トークンには $ 接頭辞が付けられます。ビルトインは裸のままです。
[[ キー .コマンド]]
キー = " プレフィックス + m "
タイプ = " シェル "
description = " 決済エージェント (一番下にスライドします)

f 受信箱）」
command = " herdr プラグイン アクションは Settle --plugin herdr-agent-inbox を呼び出します "
[[ キー .コマンド]]
key = " プレフィックス + シフト + m "
タイプ = " シェル "
description = " エージェントを未読としてマークします (注目層に戻る) "
command = " herdr プラグイン アクション、unread を呼び出します --plugin herdr-agent-inbox "
[[ キー .コマンド]]
key = " プレフィックス + alt + m "
タイプ = " シェル "
description = " 終了したすべてのエージェントをワークスペースに解決します "
command = " herdr プラグイン アクションは Settle-workspace --plugin herdr-agent-inbox を呼び出します "
[[ キー .コマンド]]
キー = " プレフィックス + i "
タイプ = " シェル "
description = " エージェントの受信箱ポップアップを開く "
command = " herdr プラグイン ペインを開く --plugin herdr-agent-inbox --entrypoint inbox "
[[ キー .コマンド]]
key = " プレフィックス + alt + right "
タイプ = " シェル "
description = " サイドバーの幅が広くなりました (+2 列) "
command = " herdr プラグイン アクション、sidebar-grow --plugin herdr-agent-inbox を呼び出します "
[[ キー .コマンド]]
key = " プレフィックス + alt + 左 "
タイプ = " シェル "
description = " サイドバーを狭くする (-2 列) "
command = " herdr プラグイン アクションは、sidebar-shrink --plugin herdr-agent-inbox を呼び出します "
次に、 herdr サーバーの reload-config を実行します。
<herdr 設定ディレクトリ>/plugins/config/herdr-agent-inbox/config.toml
(最初の実行時にコメント付きで自動作成され、変更時にホットリロードされます):
【タイトル】
source = " first " # "first" 開始プロンプト | 「最後」の最新のプロンプト
summary = false # 実際のタイトルについては、summary_cmd を介してプロンプトをパイプします
summary_cmd = " claude -p --model claude-haiku-4-5-20251001 'このコーディング エージェント セッション リクエストのタイトルを 4 ～ 8 語で記述します。タイトルのみを出力し、他には何も出力しません。 」
Summary_timeout_secs = 60
source = "last" は、エージェントがターンを終了するたびにタイトルを再取得します。
概要は非同期ワーカー上で実行され、プロンプト コンテンツごとにキャッシュされます (1 回の呼び出しごとに 1 回の呼び出し)。
セッション）、着陸するまで切り詰められた生のプロンプトに戻ります。エージェント
トランスなし

cript ref (hermes など) はペイン/ターミナル タイトルにフォールバックします
— これはいつでも上書きできます。受信トレイの r が再生成され、
set-title 制御コマンドは、選択したタイトルを固定します。
プライバシーに関する注意: タイトルはプロンプトに基づいて決定されます。保管されています
ローカル (プラグイン状態ディレクトリ) でサイドバーにレンダリングされます。有効にすると
summary 、プロンプトの抜粋は、設定したコマンドにパイプ処理されます。
コマンドがどこかに送信しない限り、マシンから何も残されません。
Herdr の有効サイドバー幅は、セッション状態によってクランプされます。
sidebar_min_width / sidebar_max_width (および herdr はドラッグをネイティブにサポートします)
サイドバーの最後の列をマウスで移動します)。サイドバーの拡大 / サイドバーの縮小
アクションはクランプを固定し、すぐに緩めることで正確な幅を強制します。
[20, 60] に戻るため、キーボードとマウスのサイズ変更が共存します。絶対幅:
python3 アクション.py サイドバー 36 。
daemon.py — プラグインによって開始される stdlib-Python デーモン
[[スタートアップ]] フック。 herdr ソケット イベント、 diffs Agent.list をサブスクライブします。
ネイティブのトランスクリプトからタイトルを取得し、ペイン/ワークスペースのメタデータをレポートします
さらに、agent.view.set の受信箱の並べ替えも可能です。終了したチャットの追加先
歴史.jsonl 。決済/未読状態は state.json にあり、キーは次のとおりです。
端末 ID、生き残った再起動。制御ソケットはアクション/TUI を提供します。
events.py — キーバインド可能なアクション (決済、未読、
Set-workspace、retitle、set-title、サイドバーのサイズ変更、restart-daemon)。
inbox_tui.py — 4 つのポップアップ ( [[panes]] 、配置ポップアップ)
ビュー、マウスのサポート、テーマに一致した色、および履歴ブラウザー。
python3 inbox_tui.py --demo を使用してスタンドアロンで試してください。
状態は、agent-inbox-state/ の herdr セッション ソケットの隣に存在します。
( state.json 、history.jsonl、daemon.log、制御ソケット)。
herdr のコーディング エージェント用の受信箱 - セッション タイトル、setol/

未読マークのトリアージ、ランタイム、ワークスペースのロールアップ、再開可能なチャット履歴
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An inbox for your coding agents in herdr — session titles, settle/mark-unread triage, runtimes, workspace rollups, and resumable chat history - douglascorrea/herdr-agent-inbox

GitHub - douglascorrea/herdr-agent-inbox: An inbox for your coding agents in herdr — session titles, settle/mark-unread triage, runtimes, workspace rollups, and resumable chat history · GitHub
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
douglascorrea
/
herdr-agent-inbox
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits assets assets scripts scripts .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md actions.py actions.py daemon.py daemon.py herdr-plugin.toml herdr-plugin.toml inbox_tui.py inbox_tui.py View all files Repository files navigation
An inbox for your coding agents. Session titles, settle / mark-unread
triage, running times, workspace rollups, and a ChatGPT-style resumable chat
history — for every agent running in herdr .
Inspired by T3 Code's "settle" workflow :
"Think of it more like an inbox. When you're done with a thread, click the
'settle' button and it slides to the bottom. This has helped me 'finish' more."
Session titles — every agent pane gets a real title derived from its
conversation ("Fix flaky payment webhook tests"), not just claude or a
directory name. Reads the agent's native transcript (claude / pi / codex),
optionally summarized by any LLM command you configure.
Inbox triage — the Agents sidebar becomes an attention queue: blocked →
finished-unseen → working → idle → settled at the bottom . Settle an
agent when you're done with it; it auto-reopens the moment it works again.
Mark one unread to pin it back on top.
Chat history that accumulates — closed conversations stay listed under
their workspace/directory (and in a global history browser), and the ones
with session refs reopen in place with claude --resume / pi --session / codex resume .
Running times — how long each session has lived and how long it's been
in its current state, live in the sidebar.
Workspace rollups — each workspace row shows its agents by state
( !1 ●2 ▸1 ⚑3 ) plus the longest-running busy stint.
Sidebar resize — grow/shrink herdr's sidebar from the keyboard (and
keep herdr's native mouse-drag working).
Everything runs through herdr's public plugin surface — no fork, no patches,
stdlib Python only.
prefix+i opens the inbox. g rotates four views, h opens history.
Compact — one line per chat, workspace headers with a dominant-state flag:
History — every closed chat, newest first; ↩ means resumable:
🔴 / ! blocked · 🔵 / ● finished, unseen · 🟡 / ▸ working ·
🟢 / ○ idle · 🏁 / ⚑ settled · ⚫ closed chat — matching herdr's
own sidebar color language (red/teal/yellow/green), in your herdr theme's
exact palette.
Requires herdr ≥ 0.7.0 and python3 (3.11+ recommended) on PATH.
herdr plugin install douglascorrea/herdr-agent-inbox
The daemon starts with your session (and a watchdog revives it). Then wire up
the display and keys:
Sidebar rows ( ~/.config/herdr/config.toml )
The plugin reports tokens; your sidebar config decides what renders:
[ ui . sidebar . agents ]
rows = [
[ " state_icon " , " $title " ],
[{ token = " $flag " , fg = " #f1fa8c " , bold = true }, " agent " , " workspace " , { token = " $age " , dim = true }],
]
[ ui . sidebar . spaces ]
rows = [
[ " state_icon " , " workspace " , " $agents " , { token = " $busy " , dim = true }],
[ " branch " , " git_status " ],
]
Available pane tokens: $title , $age (session running time), $since
(time in current state), $flag (⚑ settled / ● unread). Workspace tokens:
$agents (state counts), $busy (longest working stint). Styled custom
tokens keep the $ prefix; builtins stay bare.
[[ keys . command ]]
key = " prefix+m "
type = " shell "
description = " settle agent (slides to bottom of inbox) "
command = " herdr plugin action invoke settle --plugin herdr-agent-inbox "
[[ keys . command ]]
key = " prefix+shift+m "
type = " shell "
description = " mark agent unread (back to attention tier) "
command = " herdr plugin action invoke unread --plugin herdr-agent-inbox "
[[ keys . command ]]
key = " prefix+alt+m "
type = " shell "
description = " settle all finished agents in workspace "
command = " herdr plugin action invoke settle-workspace --plugin herdr-agent-inbox "
[[ keys . command ]]
key = " prefix+i "
type = " shell "
description = " open agent inbox popup "
command = " herdr plugin pane open --plugin herdr-agent-inbox --entrypoint inbox "
[[ keys . command ]]
key = " prefix+alt+right "
type = " shell "
description = " sidebar wider (+2 cols) "
command = " herdr plugin action invoke sidebar-grow --plugin herdr-agent-inbox "
[[ keys . command ]]
key = " prefix+alt+left "
type = " shell "
description = " sidebar narrower (-2 cols) "
command = " herdr plugin action invoke sidebar-shrink --plugin herdr-agent-inbox "
Then herdr server reload-config .
<herdr config dir>/plugins/config/herdr-agent-inbox/config.toml
(auto-created with comments on first run, hot-reloaded on change):
[ title ]
source = " first " # "first" opening prompt | "last" most recent prompt
summarize = false # pipe the prompt through summarize_cmd for a real title
summarize_cmd = " claude -p --model claude-haiku-4-5-20251001 'Write a 4-8 word title for this coding-agent session request. Output ONLY the title, nothing else.' "
summarize_timeout_secs = 60
source = "last" re-derives the title whenever the agent finishes a turn.
Summaries run on an async worker, cached per prompt content (one call per
session), falling back to the truncated raw prompt until they land. Agents
without transcript refs (e.g. hermes) fall back to their pane/terminal title
— which you can override anytime: the inbox's r regenerates, and the
set-title control command pins a title of your choosing.
Privacy note: titles are derived from your prompts. They are stored
locally (plugin state dir) and rendered in your sidebar. If you enable
summarize , prompt excerpts are piped to whatever command you configure —
nothing leaves your machine unless that command sends it somewhere.
Herdr's effective sidebar width is session state clamped by
sidebar_min_width / sidebar_max_width (and herdr natively supports dragging
the sidebar's last column with the mouse). The sidebar-grow / sidebar-shrink
actions force an exact width by pinning the clamps and immediately relaxing
them back to [20, 60] , so keyboard and mouse resize coexist. Absolute width:
python3 actions.py sidebar 36 .
daemon.py — a stdlib-Python daemon started by the plugin's
[[startup]] hook. Subscribes to herdr socket events, diffs agent.list ,
derives titles from native transcripts, and reports pane/workspace metadata
plus an agent.view.set inbox sort. Closed chats append to
history.jsonl . Settle/unread state lives in state.json , keyed by
terminal id, surviving restarts. A control socket serves the actions/TUI.
actions.py — the keybindable actions (settle, unread,
settle-workspace, retitle, set-title, sidebar resize, restart-daemon).
inbox_tui.py — the popup ( [[panes]] , placement popup) with the four
views, mouse support, theme-matched colors, and the history browser.
Try it standalone with python3 inbox_tui.py --demo .
State lives next to your herdr session socket in agent-inbox-state/
( state.json , history.jsonl , daemon.log , control socket).
An inbox for your coding agents in herdr — session titles, settle/mark-unread triage, runtimes, workspace rollups, and resumable chat history
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
