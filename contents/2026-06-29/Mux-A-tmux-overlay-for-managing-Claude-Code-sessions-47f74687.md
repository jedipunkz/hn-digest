---
source: "https://github.com/fashton28/mux"
hn_url: "https://news.ycombinator.com/item?id=48712990"
title: "Mux – A tmux overlay for managing Claude Code sessions"
article_title: "GitHub - fashton28/mux: A Claude Code session manager right inside your terminal · GitHub"
author: "fashton28"
captured_at: "2026-06-29T00:31:14Z"
capture_tool: "hn-digest"
hn_id: 48712990
score: 1
comments: 0
posted_at: "2026-06-28T23:42:55Z"
tags:
  - hacker-news
  - translated
---

# Mux – A tmux overlay for managing Claude Code sessions

- HN: [48712990](https://news.ycombinator.com/item?id=48712990)
- Source: [github.com](https://github.com/fashton28/mux)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T23:42:55Z

## Translation

タイトル: Mux – クロード コード セッションを管理するための tmux オーバーレイ
記事のタイトル: GitHub - fashton28/mux: ターミナル内のクロード コード セッション マネージャー · GitHub
説明: 端末内のクロード コード セッション マネージャー - fashton28/mux

記事本文:
GitHub - fashton28/mux: ターミナル内のクロード コード セッション マネージャー · GitHub
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
ファシュトン28
/
マルチプレクサ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ 移動先

ファイル コード [その他のアクション] メニューを開く フォルダーとファイル
24 コミット 24 コミット .claude .claude .github/ workflows .github/ workflows public public testing testing .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md mux mux mux.tmux mux.tmux すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ターミナルに常駐するクロード コード セッション マネージャー。
tmux ペインとウィンドウ間で複数のクロード コード セッションを実行していますか? mux はそれらをすべて 1 つにまとめて表示します
フローティング オーバーレイ - あなたを待っているセッションが一番上に来るように並べ替えられます - いつでもわかるようになります
どちらに注意が必要か、どれくらいの時間スタックしているか。 1 つを選択すると、すぐにそこにジャンプします。
● 作業待ち ~/dev/api 2899m │ <ライブ プレビュー
● 待機中の作業 ~/dev/web 304m │ 強調表示されたセッション
● 作業中のメイン ~/dev/cli 0m │ ターミナル画面>
● アイドルメイン ~/dev/infra 12m │
○？スクラッチ ~/dev/スクラッチ 7 分 │
クロードセッション - j/k: 移動 - J/K: スクロール - Enter: ジャンプ - ctrl-x: kill
各行は、左から右にステータス ドット/ラベル、tmux セッション、Claude セッションを示します。
実行場所、その作業ディレクトリ、および現在のステータスでの経過時間。のペイン
右は、強調表示されたセッション画面のライブ プレビューです。
ブロックされているセッションを即座に見つけます。入力を待っているセッションが色分けされ、最初に並べ替えられます。
キーを 1 回押すだけでジャンプ - Enter キーを押すと、ウィンドウをまたいででもそのセッションのペインに直接移動できます。
セッションごとのセットアップは必要ありません。Claude Code 自体のステータス ファイルを読み取ります。設定したりラップしたりするものは何もありません。
ライブ - オーバーレイが開いている間、リストとタイマーが自動的に更新されます。
ツール
バージョン
注意事項
tmux
最近のこと
mux は tmux オーバーレイです
fzf
≧ 0.38
必要になる、リロード、--track
jq
どれでも
セッションJSONを読み取ります
バッシュ、PS
プリインストールされた
macOS の標準 bash 3.2 および Linux 上で動作します
パッケージ管理を使用して 3 つのツールをインストールします

r、例: brew install tmux fzf jq (macOS) または
sudo apt install tmux fzf jq (Debian/Ubuntu)。
この README では、プレフィックスは tmux プレフィックス キー (デフォルトでは Ctrl-b) を意味します。それで
「 prefix + u 」は、 Ctrl-b を押して放し、 u を押すことを意味します。
TPM (Tmux Plugin Manager) はインストールする最も簡単な方法です
そして、マルチプレクサを常に最新の状態に保ちます。
TPM をインストールします (すでにお持ちの場合はスキップしてください)。
git clone https://github.com/tmux-plugins/tpm ~ /.tmux/plugins/tpm
mux を ~/.tmux.conf に追加します。プラグインの行を下部近くに配置し、TPM を維持します。
line を最後の行として実行します。
set -g @plugin ' tmux-plugins/tpm '
set -g @プラグイン「fashton28/mux」
run ' ~/.tmux/plugins/tpm/tpm ' # これを最後に保持
tmux をリロードしてプラグインを取得します。 tmux 内から設定をリロードします。
tmux ソースファイル ~ /.tmux.conf
次に、プレフィックス + I (大文字の I ) を押してマルチプレクサをダウンロードします。
それでおしまい。接頭辞 + u を押してオーバーレイを開きます。
mux は、単一の自己完結型スクリプトです。リポジトリのクローンを作成し、その tmux エントリポイントを
~/.tmux.conf :
git clone https://github.com/fashton28/mux ~ /.tmux/plugins/mux
実行シェル「 ~/.tmux/plugins/mux/mux.tmux 」
リロード ( tmux ソースファイル ~/.tmux.conf ) し、 prefix + u を押します。
tmux を統合しないほうがいいですか? mux スクリプトは単独で動作します - PATH 上に置きます
( ln -s "$PWD/mux" /usr/local/bin/mux ) して自分でバインドするか、単に mux list を実行します。
tmux 内でクロード コード セッションをいくつか開始します。別のペインまたはウィンドウでクロードを実行します。
任意のタブから prefix + u を押します。オーバーレイは画面上に浮かんでいます。
リスト内を移動し、プレビューを確認し、Enter キーを押してセッションにジャンプします。
mux は、現在の tmux サーバー上の Claude セッションのみをリストします (したがって、すべての行はジャンプできる行です)
に）。 tmux の外部、別の tmux サーバーで実行されているセッション、またはすでに終了しているセッションは表示されません。
キー
アクション
j / k
選択範囲を下/上に移動 (vim)
J/K
スクロールしてください

プレビュー ペインを下/上 (Shift)
↑ / ↓
選択範囲を移動 (プレビューが続きます)
入力してください
選択したセッションにジャンプし、オーバーレイを閉じます
Ctrl-X
選択したセッションを終了します (SIGTERM; ペインは残り、シェルにドロップします)
Esc
オーバーレイを閉じます
タイプ
リストのファジー フィルター (ナビゲーション キー j k J K を除く任意のキー)
ステータスの凡例
ドット
ステータス
意味
🔵
待っています
セッションにはあなたの入力が必要です - 最初にこれらに対処してください
🟠
働いている
活発に動いている
🟢
アイドル状態
完成して準備完了
⚪
?
ステータスを特定できませんでした
行は待機中→動作中→アイドル中→?にソートされます。 、各グループ内で行われているセッション
最も長いステータスが最初に表示されるため、何時間も待機していたセッションが表示されます。
一番上まで。右側の数字は、ステータスが最後に変更されてからの分数です。
~/.tmux.conf の TPM 実行行の前に、次の tmux オプションを設定します。
set -g @mux-key ' C ' # mux を開くプレフィックス キー (デフォルト: 'u')
コマンドラインの使用法
オーバーレイはメイン インターフェイスですが、スクリプトはサブコマンドを直接公開します。
mux # fzf オーバーレイを起動します (キーバインディングが実行するもの)
mux list # フォーマットされたセッションリストを出力します
muxreview < pane > # tmux ペインのライブ画面を印刷します
mux Jump < pane > < window > < session > # セッションのペインに切り替えます
mux kill < pid > # SIGTERM クロード セッション (保護されています)
トラブルシューティング
プレフィックス + u は何もしません。設定がリロードされていることを確認し ( tmux ソースファイル ~/.tmux.conf )、
TPM の場合は、プレフィックス + I を押しました。キーがまだバインドされていないことを確認します。 tmux list-keys | grep 'u' 。
オーバーレイは開きますが、リストは空です。 mux は、現在のライブ クロード セッションのみを表示します。
tmuxサーバー。 tmux ペイン内で claude を起動するか、正しいサーバーに接続していることを確認してください。
マルチプレクサ: fzf >= 0.38 が必要です。 fzf をアップグレードする ( brew upgrade fzf 、またはからリリースを取得します)
fzf リポジトリ)。
あなたは衝突します

別のバインディングを使用します。 set -g @mux-key '...' を使用して別のキーを選択します。
セッション リスト ロジックは単一のテスト シームの背後にあります。マルチプレクサ リストは外部入力を読み取ります。
環境変数を使用して、フィクスチャで駆動できる純粋な決定論的な関数になります - ライブではありません
tmux サーバーまたは実際の Claude プロセスが必要です。
MUX_SESSIONS_DIR=テスト/フィクスチャ/セッション \
MUX_PANES_FILE=テスト/フィクスチャ/panes.txt \
MUX_PPID_FILE=テスト/フィクスチャ/ppids.txt \
MUX_NOW=1782657704 \
マルチプレクサリスト
チェックを実行します (バットとシェルチェックが必要):
コウモリのテスト/
shellcheck -s bash mux mux.tmux
について
端末内のクロード コード セッション マネージャー
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

A Claude Code session manager right inside your terminal - fashton28/mux

GitHub - fashton28/mux: A Claude Code session manager right inside your terminal · GitHub
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
fashton28
/
mux
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
24 Commits 24 Commits .claude .claude .github/ workflows .github/ workflows public public tests tests .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md mux mux mux.tmux mux.tmux View all files Repository files navigation
A Claude Code session manager that lives in your terminal.
Running several Claude Code sessions across tmux panes and windows? mux shows them all in one
floating overlay - sorted so the sessions waiting on you rise to the top - so you always know
which one needs attention and how long it has been stuck. Pick one and you jump straight to it.
● waiting work ~/dev/api 2899m │ <live preview of the
● waiting work ~/dev/web 304m │ highlighted session's
● working main ~/dev/cli 0m │ terminal screen>
● idle main ~/dev/infra 12m │
○ ? scratch ~/dev/scratch 7m │
Claude sessions - j/k: move - J/K: scroll - enter: jump - ctrl-x: kill
Each row shows, left to right: a status dot/label, the tmux session the Claude session
runs in, its working directory , and how long it has been in its current status. The pane on the
right is a live preview of the highlighted session's screen.
Find the blocked one instantly - sessions waiting for your input are colored and sorted first.
Jump in one keypress - Enter takes you straight to that session's pane, even across windows.
No setup per session - it reads Claude Code's own status files; nothing to configure or wrap.
Live - the list and timers refresh on their own while the overlay is open.
Tool
Version
Notes
tmux
any recent
mux is a tmux overlay
fzf
≥ 0.38
needs become , reload , --track
jq
any
reads the session JSON
bash , ps
preinstalled
runs on macOS's stock bash 3.2 and on Linux
Install the three tools with your package manager, e.g. brew install tmux fzf jq (macOS) or
sudo apt install tmux fzf jq (Debian/Ubuntu).
Throughout this README, prefix means your tmux prefix key - Ctrl-b by default. So
" prefix + u " means press Ctrl-b , release, then press u .
TPM (the Tmux Plugin Manager) is the easiest way to install
and keep mux updated.
Install TPM (skip if you already have it):
git clone https://github.com/tmux-plugins/tpm ~ /.tmux/plugins/tpm
Add mux to your ~/.tmux.conf . Put the plugin lines near the bottom, and keep the TPM
run line as the very last line:
set -g @plugin ' tmux-plugins/tpm '
set -g @plugin ' fashton28/mux '
run ' ~/.tmux/plugins/tpm/tpm ' # keep this last
Reload tmux and fetch the plugin. From inside tmux, reload the config:
tmux source-file ~ /.tmux.conf
then press prefix + I (capital I ) to download mux.
That's it. Press prefix + u to open the overlay.
mux is a single self-contained script. Clone the repo and source its tmux entrypoint from your
~/.tmux.conf :
git clone https://github.com/fashton28/mux ~ /.tmux/plugins/mux
run-shell ' ~/.tmux/plugins/mux/mux.tmux '
Reload ( tmux source-file ~/.tmux.conf ) and press prefix + u .
Prefer no tmux integration at all? The mux script works on its own - put it on your PATH
( ln -s "$PWD/mux" /usr/local/bin/mux ) and bind it yourself, or just run mux list .
Start a few Claude Code sessions inside tmux - run claude in separate panes or windows.
Press prefix + u from any tab. The overlay floats over your screen.
Move through the list, watch the preview, and Enter to jump into a session.
mux only lists Claude sessions on your current tmux server (so every row is one you can jump
to). Sessions running outside tmux, in another tmux server, or already finished are not shown.
Key
Action
j / k
move selection down / up (vim)
J / K
scroll the preview pane down / up (Shift)
↑ / ↓
move selection (preview follows)
Enter
jump to the selected session and close the overlay
ctrl-x
terminate the selected session (SIGTERM; its pane stays, drops to a shell)
Esc
close the overlay
type
fuzzy-filter the list (any key except the navigation keys j k J K )
Status legend
Dot
Status
Meaning
🔵
waiting
the session needs your input - deal with these first
🟠
working
actively running
🟢
idle
finished and ready
⚪
?
status could not be determined
Rows are sorted waiting → working → idle → ? , and within each group the session that has been in
its status longest appears first - so a session that has been waiting on you for hours floats
to the very top. The right-hand number is minutes since the status last changed.
Set these tmux options before the TPM run line in ~/.tmux.conf :
set -g @mux-key ' C ' # which prefix key opens mux (default: 'u')
Command-line usage
The overlay is the main interface, but the script exposes subcommands directly:
mux # launch the fzf overlay (what the keybinding runs)
mux list # print the formatted session list
mux preview < pane > # print a tmux pane's live screen
mux jump < pane > < window > < session > # switch to a session's pane
mux kill < pid > # SIGTERM a Claude session (guarded)
Troubleshooting
prefix + u does nothing. Make sure the config reloaded ( tmux source-file ~/.tmux.conf ) and,
for TPM, that you pressed prefix + I . Check the key isn't already bound: tmux list-keys | grep ' u ' .
The overlay opens but the list is empty. mux only shows live Claude sessions on the current
tmux server. Start claude inside a tmux pane, or check you're attached to the right server.
mux: fzf >= 0.38 required . Upgrade fzf ( brew upgrade fzf , or grab a release from the
fzf repo ).
u collides with another binding. Pick a different key with set -g @mux-key '...' .
The session-listing logic sits behind a single test seam: mux list reads its external inputs from
environment variables, making it a pure, deterministic function you can drive with fixtures - no live
tmux server or real Claude processes required.
MUX_SESSIONS_DIR=tests/fixtures/sessions \
MUX_PANES_FILE=tests/fixtures/panes.txt \
MUX_PPID_FILE=tests/fixtures/ppids.txt \
MUX_NOW=1782657704 \
mux list
Run the checks (needs bats and shellcheck ):
bats tests/
shellcheck -s bash mux mux.tmux
About
A Claude Code session manager right inside your terminal
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
