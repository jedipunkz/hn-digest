---
source: "https://github.com/inkshell/inkshell/"
hn_url: "https://news.ycombinator.com/item?id=49336452"
title: "A tabbed desktop workspace for Claude Code"
article_title: "GitHub - inkshell/inkshell: A tabbed desktop workspace for Claude Code. The CLI, with style. Every session in a tab, every project in one window — a thin shell around your own claude binary, never a fork. · GitHub"
image: "https://opengraph.githubassets.com/60703bac02cd1cde3baf8a3dc997151e945374fa1d0432518f10c79a92fb6bc6/inkshell/inkshell"
author: "rodorgas"
captured_at: "2026-08-17T20:16:04Z"
capture_tool: "hn-digest"
hn_id: 49336452
score: 1
comments: 0
posted_at: "2026-08-17T19:36:02Z"
tags:
  - hacker-news
  - translated
---

# A tabbed desktop workspace for Claude Code

- HN: [49336452](https://news.ycombinator.com/item?id=49336452)
- Source: [github.com](https://github.com/inkshell/inkshell/)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T19:36:02Z

## Translation

タイトル: Claude Code のタブ付きデスクトップ ワークスペース
記事のタイトル: GitHub - inkshell/inkshell: Claude Code 用のタブ付きデスクトップ ワークスペース。スタイリッシュな CLI。タブ内のすべてのセッション、1 つのウィンドウ内のすべてのプロジェクト — フォークではなく、独自のクロード バイナリを囲む薄いシェルです。 · GitHub
説明: Claude Code 用のタブ付きデスクトップ ワークスペース。スタイリッシュな CLI。タブ内のすべてのセッション、1 つのウィンドウ内のすべてのプロジェクト — フォークではなく、独自のクロード バイナリを囲む薄いシェルです。 - インクシェル/インクシェル

記事本文:
GitHub - inkshell/inkshell: Claude Code 用のタブ付きデスクトップ ワークスペース。スタイリッシュな CLI。タブ内のすべてのセッション、1 つのウィンドウ内のすべてのプロジェクト — フォークではなく、独自のクロード バイナリを囲む薄いシェルです。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
インクシェル
/
インクシェル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
75 コミット 75 コミット .github

.github .vscode .vscode build build docs docs リソース リソース src src .editorconfig .editorconfig .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc .prettierrc CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス Makefile Makefile 通知 通知 README.md README.md SECURITY.md SECURITY.md electric-builder.yml electric-builder.yml electric.vite.config.ts electric.vite.config.ts eslint.config.js eslint.config.js install.sh install.sh package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json tsconfig.web.json tsconfig.web.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code のタブ付きデスクトップ ワークスペース。スタイリッシュな CLI。
InkShell は、すでに claude CLI を使用していて、そのつもりがない人向けです。
放置したいが、複数のプロジェクトを同時にやりくりしており、デスクトップを使いたい人
彼ら。 1 つのペイン内のすべてのセッション、1 つのウィンドウ内のすべてのプロジェクト、それぞれに独自のウィンドウ
構成により、職場と個人が資格情報や履歴を共有することはありません。
それらの間の切り替えには費用はかかりません。
それは本物の周りに薄い殻をかぶったままです。 InkShell は決して Claude を再実装しません
コード: 擬似端末内にローカルにインストールされた独自のクロードを生成します。
常に元の安定した CLI を実行しており、機能が提供されます。
追いつくたびではなく、ターミナルに発送される日です。フォークはありません、いいえ
再パッケージ化されたバイナリ、バージョンの遅れなし。 InkShell を閉じると、CLI が正確に表示されます。
あなたはそれを残しました。
InkShell はコミュニティ プロジェクトであり、Anthropic とは提携していません。
「クロード」および「クロード・コード」はアンソロピックの商標です。
タブ ストリップではなく分割ワークスペース: 1、2、または 4 つのペインを一度に表示し、
SI からチャットをドラッグする

debar — または別のペインから — それらのいずれかに追加します。ペイン
背後のプロセスには触れずに最小化および最大化できるため、セッションは
画面外へのプッシュは実行を続け、終了したときとまったく同じ状態に戻ります。
すべてのプロジェクトが 1 つのウィンドウに表示されます。サイドバーはツリーです。各プロジェクトには、そのプロジェクトが含まれています。
開いているチャット、ターミナル、およびその下にネストされているファイル、どのペインを示すバッジ
それぞれが現在存在しており、リスト自体をドラッグして並べ替えることができます。あらゆるプロジェクト
行は、クリックするかどうかに関係なく、そのプロジェクトのチャットまたはターミナルをワンクリックで開始します。
それが現在選択されているものです。
Claude Code 独自の履歴: InkShell が CLI のトランスクリプト ストアを読み取る
( ~/.claude/projects ) 過去のセッションをリストし、それぞれに同じ名前を付けます
ai-title CLI が使用するか、ペイン内で再開するか、リストから削除します。
プロジェクトごとの構成: 各プロジェクトには、色合いを変えるアクセント カラーが含まれています。
クロムとそれに属するすべてのペイン、および独自の Claude config ディレクトリ
( CLAUDE_CONFIG_DIR )。プロジェクトを別の構成ディレクトリとそのセッションにポイントします。
履歴とコンテキストメーターはすべてそれに続きます。シェルのエイリアスや .envrc のジャグリングはありません。
プレーンターミナル、同じウィンドウ: プロジェクトで独自の $SHELL を開きます
ディレクトリを単なる別のペインとして - git rebase または開発サーバー用
チャットに時間を費やしたくありませんでした。
モデルとエフォートのスイッチャー: ワンピックタイプ /model または /effort をライブに追加
セッション。示されているモデルは、実際にそれをサポートしているものであり、
推測ではなく転写。
コンテキスト メーター: CLI のコンテキスト インジケーターをライブで反映する残量ゲージ
アクティブなセッションのトランスクリプトから、そのモデル自体のコンテキストに対して測定されます
窓。各ペインのタイトル バーには独自の読みが表示されます。
Git パネル: ウィンドウを離れることなく、ステージング、ステージング解除、コミット、プッシュを実行できます。
未プッシュのコミットがマークされたブランチ履歴を参照します。

任意の差分ファイルを開きます
または、独自のペインとしてコミットします。コミットメッセージは、Claude が次のように作成できます。
ワンクリック。
ファイル、差分、および実際のエディタ: 変更されたファイルを含むプロジェクトのツリー
マークが付いている場合、Monaco エディタでファイルを開くと、実際に編集して保存できます。
Monaco の差分としてレンダリングされた差分とコミット。ワンクリックでファイルを参照できます。
ダブルクリックすると固定されます。クロードが出力で言及したファイル パスはクリック可能です —
ディスクに対して検証されるため、実際のファイルのみが点灯します。
クイックオープン: アクティブなプロジェクト内のすべてのファイルをあいまい検索して開きます
エディタに直接入力します。
4 つのペイン、1 つのウィンドウ — 異なるプロジェクトからの 2 つのチャット、ターミナル、およびエディター内のファイル。それらのいずれかをペイン間でドラッグします。画面外に押したものは実行され続けます。
キー
アクション
⌘T / Ctrl+T
選択したプロジェクトの新しいチャット
⌘W / Ctrl+W
フォーカスされたペインを最小化します (セッションは実行され続けます)
⌘P / Ctrl+P
クイックオープン - アクティブなプロジェクト内のあいまいなファイル検索
⌘S / Ctrl+S
エディターで開いているファイルを保存します
中クリック
ペイン上: 最小化します・サイドバー項目上: 閉じます
右クリック
プロジェクト上: 設定 · 履歴カード上: チャットを削除
📦 要件
クロード コードがインストールされています (
claude コマンドはターミナルで実行する必要があります)。 InkShell はログイン シェルを尋ねます
どこにあるかを確認し、通常のインストール場所もチェックして、引き続き
切り詰められた PATH を使用して Finder からアプリを開いたときの CLI。あ
非標準インストールは、 INKSHELL_CLAUDE_BIN を使用して直接指定できます。
Node.js ≥ 20 およびソースからビルドするための npm。
カール -fsSL https://raw.githubusercontent.com/inkshell/inkshell/main/install.sh |しー
最新リリースをダウンロードします
— Apple Silicon または Intel 上で実行されるユニバーサル ビルド — にインストールされます。
/Applications (書き込み可能でない場合は ~/Applications)。
Windows バージョン

もうすぐ来ます。それが実現するまで、Windows と Linux ユーザーは
ソースから InkShell を実行できます。以下の「はじめに」を参照してください。
なぜ単純なダウンロードではなくスクリプトを使用するのでしょうか? InkShell ビルドはまだコード署名されていません。
macOS はブラウザがダウンロードしたものをすべて隔離するため、そのアプリを開くと
「InkShell が破損しているため開けません」という誤解を招くメッセージが表示されます。
ダイアログ。 CURL のダウンロードは隔離されないため、スクリプトによってアプリがインストールされます
それはただ開くだけです。 (最初にスクリプトを読んでください。スクリプトは ~70 です
sh の行。)
から .zip をダウンロードします。
リリースページを解凍し、
InkShell.app を /Applications に移動し、隔離フラグをクリアします。
ダウンロードに添付されているブラウザ:
xattr -dr com.apple.quarantine /Applications/InkShell.app
この最後の手順を行わないと、macOS は上記の「破損した」ダイアログを表示します。ファイルは次のとおりです。
大丈夫です。このメッセージは、ゲートキーパーが「署名されておらず、隔離されている」と表現したものです。
#1. クローンを作成する
git clone https://github.com/inkshell/inkshell.git
CDインクシェル
# 2. インストール (Electron 用のネイティブの node-pty モジュールも再構築します)
npmインストール
# 3. 開発環境で実行 (ホットリロード)
npm 実行開発
プラットフォーム用の配布可能アプリを作成するには:
npm run パック:mac # .zip
npm run Pack:win # NSIS インストーラー
npm run パック:linux # AppImage + .deb
本日出荷されるターゲットは macOS のみです。 Windows および Linux のビルドは次のとおりです。
接続されていますが、まだテストされていません - 試してみると、問題 (または PR)
どうだったか教えていただけると大歓迎です。
InkShell は、標準的な 3 プロセスの Electron アプリです。
コードが保持する 2 つのルール: レンダラーは、経由する場合を除き、OS には決して触れません。
src/shared に型指定された IPC コントラクトがあり、Claude Code 自身のデータが読み取られます。
決して書かれていない — ~/.claude/projects は CLI であり、InkShell は独自のものを保持しています
~/.inkshell/config.json の設定。
全体像については、docs/ARCHITECTURE.md を参照してください。
すべての色、半径、GL

ow は先頭の CSS 変数に存在します。
src/renderer/src/styles/theme.css 。
InkShell のテーマの変更は 1 つのファイルで編集できます。
現在、InkShell は Claude Code のみを話します。デザインはそうではありません
ただし、それに依存します。アプリは疑似端末内で実際の CLI エージェントを駆動します。
そして、エージェントがすでに書いたトランスクリプトを読み取ります。これは、
1 つのツールが適合します。
Codex と GitHub Copilot が次のターゲットです。目標は最低ではない
3 つすべてに共通点がありますが、各プロジェクトが開くウィンドウは 1 つだけです。
実際に必要なエージェント。
貢献は大歓迎です。 COTRIBUTING.md および当社の
行動規範。良い最初の問題にはラベルが付けられます
良い創刊号。
Apache License 2.0 に基づいてライセンスされています。については「通知」を参照してください。
帰属と商標の詳細。
Claude Code のタブ付きデスクトップ ワークスペース。スタイリッシュな CLI。タブ内のすべてのセッション、1 つのウィンドウ内のすべてのプロジェクト — フォークではなく、独自のクロード バイナリを囲む薄いシェルです。
Readme Apache-2.0 ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A tabbed desktop workspace for Claude Code. The CLI, with style. Every session in a tab, every project in one window — a thin shell around your own claude binary, never a fork. - inkshell/inkshell

GitHub - inkshell/inkshell: A tabbed desktop workspace for Claude Code. The CLI, with style. Every session in a tab, every project in one window — a thin shell around your own claude binary, never a fork. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
inkshell
/
inkshell
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
75 Commits 75 Commits .github .github .vscode .vscode build build docs docs resources resources src src .editorconfig .editorconfig .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc .prettierrc CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md electron-builder.yml electron-builder.yml electron.vite.config.ts electron.vite.config.ts eslint.config.js eslint.config.js install.sh install.sh package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json tsconfig.web.json tsconfig.web.json View all files Repository files navigation
A tabbed desktop workspace for Claude Code . The CLI, with style.
InkShell is for people who already live in the claude CLI and have no intention
of leaving it, but who juggle several projects at once and want a desktop around
them. Every session in a pane, every project in one window, each with its own
configuration, so work and personal never share credentials or history and
switching between them costs nothing.
It stays a thin shell around the real thing. InkShell never reimplements Claude
Code: it spawns your own locally-installed claude inside a pseudo-terminal, so
you are always running the original, stable CLI , and a feature reaches you the
day it ships to the terminal rather than whenever we catch up. No fork, no
repackaged binary, no version lag. Close InkShell and your CLI is exactly where
you left it.
InkShell is a community project and is not affiliated with Anthropic .
"Claude" and "Claude Code" are trademarks of Anthropic.
A split workspace, not a tab strip : show 1, 2 or 4 panes at once and
drag a chat from the sidebar — or from another pane — into any of them. Panes
minimize and maximize without touching the process behind them, so a session you
push off-screen keeps running and comes back exactly as you left it.
Every project in one window : the sidebar is a tree — each project, with its
open chats, terminals and files nested underneath, a badge showing which pane
each one currently sits in, and drag-to-reorder for the list itself. Any project
row starts a chat or a terminal in that project in one click, whether or not
it's the one currently selected.
History that is Claude Code's own : InkShell reads the CLI's transcript store
( ~/.claude/projects ) to list past sessions, names each one with the same
ai-title the CLI uses, resumes it in a pane, or deletes it from the list.
Per-project configuration : each project carries an accent color that tints
the chrome and every pane belonging to it, plus its own Claude config directory
( CLAUDE_CONFIG_DIR ). Point a project at a separate config dir and its sessions,
history, and context meter all follow it. No shell aliases, no .envrc juggling.
Plain terminals, same window : open your own $SHELL in the project
directory as just another pane — for the git rebase or the dev server you
didn't want to spend a chat on.
Model & effort switchers : one pick types /model or /effort into the live
session. The model shown is the one actually backing it, read from the
transcript rather than guessed.
Context meter : a fuel gauge that mirrors the CLI's context indicator, live
from the active session's transcript, measured against that model's own context
window. Every pane carries its own reading in its title bar.
Git panel : stage, unstage, commit, and push without leaving the window;
browse the branch history with unpushed commits marked, and open any diff, file
or commit as a pane of its own. Commit messages can be drafted by Claude with
one click.
Files, diffs, and a real editor : the project's tree with modified files
marked, files opening in a Monaco editor you can actually edit and save, and
diffs and commits rendered as Monaco diffs. A single click peeks at a file, a
double click pins it. File paths Claude mentions in its output are clickable —
verified against the disk, so only real files light up.
Quick Open : fuzzy-search every file in the active project and open it
straight into the editor.
Four panes, one window — two chats from different projects, a terminal, and a file in the editor. Drag any of them between panes; what you push off-screen keeps running.
Keys
Action
⌘T / Ctrl+T
New chat in the selected project
⌘W / Ctrl+W
Minimize the focused pane (the session keeps running)
⌘P / Ctrl+P
Quick Open — fuzzy file search in the active project
⌘S / Ctrl+S
Save the file open in the editor
Middle click
On a pane: minimize it · on a sidebar item: close it
Right click
On a project: its settings · on a history card: delete the chat
📦 Requirements
Claude Code installed (the
claude command must run in your terminal). InkShell asks your login shell
where it is and also checks the usual install locations, so it still finds the
CLI when the app is opened from the Finder with a truncated PATH . A
non-standard install can be pointed at directly with INKSHELL_CLAUDE_BIN .
Node.js ≥ 20 and npm to build from source.
curl -fsSL https://raw.githubusercontent.com/inkshell/inkshell/main/install.sh | sh
That downloads the latest release
— a universal build that runs on Apple Silicon or Intel — and installs it into
/Applications (or ~/Applications when that isn't writable).
A Windows version is coming soon. Until it lands, Windows and Linux users
can run InkShell from source — see Getting started below.
Why a script and not a plain download? InkShell builds aren't code-signed yet,
and macOS quarantines anything a browser downloads, so opening the app that
way greets you with a misleading "InkShell is damaged and can't be opened"
dialog. curl downloads are never quarantined, so the script installs an app
that just opens. (You can read the script first — it's ~70
lines of sh .)
Download the .zip from the
Releases page , unzip it,
move InkShell.app to /Applications , then clear the quarantine flag your
browser attached to the download:
xattr -dr com.apple.quarantine /Applications/InkShell.app
Without that last step, macOS shows the "damaged" dialog above — the file is
fine; the message is Gatekeeper's way of saying "unsigned and quarantined".
# 1. Clone
git clone https://github.com/inkshell/inkshell.git
cd inkshell
# 2. Install (also rebuilds the native node-pty module for Electron)
npm install
# 3. Run in development (hot reload)
npm run dev
To produce a distributable app for your platform:
npm run pack:mac # .zip
npm run pack:win # NSIS installer
npm run pack:linux # AppImage + .deb
macOS is the only target that ships today. The Windows and Linux builds are
wired up but haven't been tested yet — if you try one, an issue (or a PR)
telling us how it went is very welcome.
InkShell is a standard three-process Electron app:
Two rules the code holds to: the renderer never touches the OS except through
the typed IPC contract in src/shared , and Claude Code's own data is read,
never written — ~/.claude/projects is the CLI's, InkShell keeps its own
config in ~/.inkshell/config.json .
See docs/ARCHITECTURE.md for the full picture.
Every color, radius, and glow lives in CSS variables at the top of
src/renderer/src/styles/theme.css .
Re-theming InkShell is a one-file edit.
Today InkShell speaks Claude Code , and only Claude Code. The design doesn't
depend on that, though: the app drives a real CLI agent inside a pseudo-terminal
and reads the transcripts that agent already writes, which is a shape more than
one tool fits.
Codex and GitHub Copilot are the next targets. The goal isn't a lowest
common denominator across all three, but one window where each project opens the
agent it actually calls for.
Contributions are very welcome. See CONTRIBUTING.md and our
Code of Conduct . Good first issues are labeled
good first issue .
Licensed under the Apache License 2.0 . See NOTICE for
attribution and trademark details.
A tabbed desktop workspace for Claude Code. The CLI, with style. Every session in a tab, every project in one window — a thin shell around your own claude binary, never a fork.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
