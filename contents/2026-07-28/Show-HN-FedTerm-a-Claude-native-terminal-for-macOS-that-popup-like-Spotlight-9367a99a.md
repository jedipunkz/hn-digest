---
source: "https://github.com/feod1/fedterm"
hn_url: "https://news.ycombinator.com/item?id=49088454"
title: "Show HN: FedTerm – a Claude-native terminal for macOS that popup like Spotlight"
article_title: "GitHub - feod1/fedterm · GitHub"
author: "treenix_io"
captured_at: "2026-07-28T19:07:50Z"
capture_tool: "hn-digest"
hn_id: 49088454
score: 1
comments: 0
posted_at: "2026-07-28T19:06:16Z"
tags:
  - hacker-news
  - translated
---

# Show HN: FedTerm – a Claude-native terminal for macOS that popup like Spotlight

- HN: [49088454](https://news.ycombinator.com/item?id=49088454)
- Source: [github.com](https://github.com/feod1/fedterm)
- Score: 1
- Comments: 0
- Posted: 2026-07-28T19:06:16Z

## Translation

タイトル: Show HN: FedTerm – Spotlight のようにポップアップする macOS 用のクロードネイティブターミナル
記事タイトル: GitHub - feod1/fedterm · GitHub
説明: GitHub でアカウントを作成して、feod1/fedterm の開発に貢献します。
HN テキスト: こんにちは、私は Mac 用のターミナルを作りました。ホットキーを押すと、何でも上にポップアップ表示されます
あなたがいる窓は、スポットライトのようなものです。もう一度押すと消えます。すべてのクロード コード セッションを保存してインデックスを作成するため、セッション全体を検索できます。
そして、置いた場所からすぐに拾います。同じウィンドウであなたの検索も行われます
シェル履歴と保存された SSH ホスト。 swift + swiftterm、mit ライセンス、アカウントなし、販売するものはありません。何を聞いて嬉しいですか
あなたは思います。 github.com/feddot2517/fedterm

記事本文:
GitHub - feod1/fedterm · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
食事1
/
フェデラル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらにアクトを開く

イオンメニュー フォルダーとファイル
1 コミット 1 コミット リソース リソース ソース/ FedTerm ソース/ FedTerm アセット アセット .gitignore .gitignore ライセンス ライセンス Makefile Makefile Package.resolved Package.resolved Package.swift Package.swift README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
FedTerm は、グローバル ショートカットで開く macOS 用のターミナルです。 ⌥ Space キーを押すと、全画面アプリや他のデスクトップなど、作業中のアプリの上にウィンドウが表示されます。コマンド、SSH ホスト、またはフォルダー パスを入力し、 Enter を押すと、ウィンドウがターミナルになります。ショートカットをもう一度押すと非表示になり、フォーカスが前のアプリに戻ります。
このアプリは、zsh フックを使用して、実行したコマンドも記録します。そのため、同じウィンドウがコマンド履歴の検索、接続するサーバーのリスト、およびクロード コード セッションのランチャーとして機能します。
FedTerm はメニュー バーにあり、Dock アイコンはありません。ショートカットはお好みの組み合わせに変更できます。
ショートカットは Carbon を通じて登録されるため、アクセシビリティ権限は必要ありません。
設定で再バインドできます。フィールドをクリックして、必要なキーを押します。 F キーは修飾子なしで機能しますが、他のキーには少なくとも 1 つ必要です。その組み合わせが別のアプリによって使用された場合、FedTerm はその旨を通知し、前の組み合わせを保持します。
ウィンドウは他のアプリの上に浮かび上がり、デスクトップ間を移動します。
外側をクリックすると非表示になります。開いたままにしておきたい場合は、ピンで固定します。
そのサイズと位置は、起動の間に記憶されます。保存された位置が画面外にある場合、ウィンドウは中央に開きます。
FedTerm は入力内容を調べて、以下に適合するものを提供します。
矢印キーでリスト内を移動し、Enter で選択した項目を実行し、Esc でフィールドをクリアします。一致はあいまいなので、通常は数文字で十分です。
保存されたコマンドと接続
お気に入り

コマンド。コマンドにスターを付けると、ウィンドウの上部に表示されます。名前を変更し、アプリの起動時に独自のタブで開くようにマークすることができます。
保存された SSH 接続。サーバーを固定し、ラベルを付けます。最近の接続は独自に収集され、最近の接続と使用頻度によって並べ替えられます。
コマンド履歴。コマンドは、過去 1 時間、6 時間、今日、昨日、今週、古いというように経過時間別にグループ化されます。グループ内の繰り返しは 1 回のみ表示されます。過去 1 時間については、ツール ( git 、 docker 、 ssh 、 npm など) ごとの数とクリック可能なサーバー名を含む短い概要も表示されます。
名前付きセッション。ウィンドウにフォルダーをドロップするか、パスを貼り付けて、セッションに名前を付けます。 FedTerm は独自のセッション ID でセッションを開始し、後で新しいセッションを作成するのではなく、そのセッションを再開します。
過去のすべてのセッション。 FedTerm は ~/.claude/projects を読み取り、各トランスクリプトの概要または最初のメッセージをそのフォルダーとともに取得し、プロジェクトごとおよび日付ごとにグループ化されたセッションを検索フィールドとともにリストします。結果はファイルの変更時間ごとにキャッシュされるため、リストは遅延なく開きます。
清潔な環境。 CLAUDE* および ANTHROPIC* 変数は新しいタブから削除されます。 FedTerm から開始されたセッションは、FedTerm 自体が Claude Code から起動された場合でも、スタンドアロンのセッションのように動作します。
⌘T でタブを開き、⌘W でタブを閉じ、⌘1 ～ ⌘9 でタブを切り替え、⌘[ と ⌘] で左右に移動します。
タブをドラッグして順序を変更したり、ダブルクリックして名前を変更したりできます。
開いているタブは次回起動時に復元されます。 SSH タブは自動的に再接続します。
タブのタイトルはシェルの後に続きます。現在のフォルダー、実行中のプログラムによって設定されたタイトル、またはサーバー名です。
何かがまだ実行中の場合、FedTerm は終了する前に確認を求めます。
コマンドは作業ディレクトリとともに⌃1 – ⌃9 にバインドできます。これらのショートカットは、

ウィンドウが開いているため、他のアプリには影響しません。複数行のスクリプトはファイルに保存され、bash を通じて実行されるため、ループと条件は記述どおりに機能します。
13 の組み込みテーマがあります: Terminal.app パレット、One Dark、Dracula、Nord、Gruvbox、両方の Solarized バリアント、Tokyo Night、Catppuccin Mocha、Monokai、GitHub Light、および 2 つのニュートラル グレー。背景、テキスト、キャレット、選択範囲、ANSI 16 色すべて、ぼかし上の背景の透明度など、あらゆるテーマをコピーして編集できます。フォントファミリー、サイズ、ウェイトは設定可能です。テキストを鮮明にするための細いストロークのオプションや、明るい壁紙の上でガラスを暗くするスライダーもあります。変更は開いているターミナルにすぐに適用されます。
ターミナルは SwiftTerm : xterm-256color、トゥルーカラー、およびマウスのサポートです。 ⌘ を押しながらクリックすると、リンク、ファイル パス、および OSC 8 ハイパーリンクが開きます。スクロールバーのガターが削除されるため、mc や htop などの全幅プログラムは幅全体を使用します。 ⌃C は pty に直接書き込まれるため、常に中断されます。ショートカットは物理キー コードによって照合され、キリル文字やその他の非ラテン語レイアウトでも機能し続けます。インターフェイスは、ロシアのシステムではロシア語で、その他のシステムでは英語で表示されます。
macOS 13 Ventura 以降、Apple Silicon または Intel
Swift 5.9 以降の Xcode コマンド ライン ツール
zsh、macOS のデフォルトのシェル。コマンド履歴はそれを通じてキャプチャされます
クロード コード機能の PATH 内のクロード CLI
git クローン https://github.com/feddot2517/fedterm.git
cd フェデターム
make Bundle # リリースバイナリをビルドし、dist/FedTerm.app をアセンブルします
dist/FedTerm.app を開く
make run は両方のステップを一度に実行します。 make dev は、バンドルを構築せずにソースからアプリを実行します。
バイナリはアドホックに署名されるため、Gatekeeper は最初の起動をブロックします。アプリを右クリックして [開く] を選択するか、隔離フラグを削除します。
xattr -dr com.apple.quaranti

dist/FedTerm.app にあります
アプリを /Applications に移動し、システム設定 → 一般 → ログイン項目に追加して、常に実行できるようにします。
ショートカット
アクション
⌥ スペース
ウィンドウを表示または非表示にする (再バインド可能)
入力してください
選択した項目を実行します
↑ ↓
結果と履歴をたどる
Esc
フィールドをクリアする
⌘T / ⌘W
タブを開いたり閉じたりする
⌘1 – ⌘9
タブに切り替える
⌘[ / ⌘]
前または次のタブ
⌃1 – ⌃9
カスタムコマンドを実行する
⌘+ / ⌘- / ⌘0
フォントサイズのアップ、ダウン、リセット
⌘ を押しながらクリック
出力からリンクまたはパスを開く
⌃C
実行中のプログラムを中断する
データ
すべてのファイルはローカルの ~/Library/Application Support/FedTerm/ にあります。
アプリにはネットワーク コードがないため、何もアップロードされず、テレメトリもありません。履歴はプレーン テキスト ファイルなので、いつでも削除できます。
コマンドは入力したとおりに保存されることに注意してください。トークンまたはパスワードをインラインで貼り付けると、ファイルに保存されます。 ~/.zsh_history のように扱います。
ソース/FedTerm/
§─ main.swift、AppDelegate.swift の起動、メニュー バー、ウィンドウのショートカット
§─ SpotlightPanel.swift フローティング ウィンドウ、ぼかし、保存された位置
§─ ContentView.swift、TabsModel.swift タブ バー、並べ替え、永続化された状態
§─ HomeView.swift 検索フィールド、結果、履歴
§─ TerminalSession.swift SwiftTerm サブクラス、シェル プロセス、⌘クリック
§─ HistoryStore.history.jsonl の迅速な閲覧と閲覧
§─ ShellIntegration.swift が preexec フックを使用して生成した zsh 構成
§─ ClaudeStore.swift、ClaudeUI.swift クロード コード セッションとそのブラウザ
§─ AutomationsStore.swift、AutomationEditor.swift カスタム ⌃1 ～ 9 コマンド
§─ ReviewsStore.swift のお気に入りコマンドと自動起動
§─ Theme.swift、ThemeEditor.swift、SettingsUI.swift テーマ、フォント、設定
§─ HotkeyManager.swift、HotkeyRecorder.swift のグローバル ショートカットと再バインド
§─ Models.swif

t SSH 解析とコマンド分類
└─ L10n.swift ロシア語と英語の文字列
クレジット
Miguel de Icaza による SwiftTerm は端末エミュレーションを行います。
カラースキームは、One Dark、Dracula、Nord、Gruvbox、Solarized、Tokyo Night、Catppuccin、Monokai から引用されています。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to feod1/fedterm development by creating an account on GitHub.

hi hn, i made a terminal for mac. you press a hotkey and it pops up over whatever
window you're in, kinda like spotlight. press it again and it's gone. it saves and indexes all your claude code sessions, so you can search through them
and pick up any one right where you left it. the same window also searches your
shell history and your saved ssh hosts. swift + swiftterm, mit licensed, no accounts, nothing to sell. happy to hear what
you think. github.com/feddot2517/fedterm

GitHub - feod1/fedterm · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
feod1
/
fedterm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit Resources Resources Sources/ FedTerm Sources/ FedTerm assets assets .gitignore .gitignore LICENSE LICENSE Makefile Makefile Package.resolved Package.resolved Package.swift Package.swift README.md README.md View all files Repository files navigation
FedTerm is a terminal for macOS that opens with a global shortcut. Press ⌥ Space and a window appears on top of the app you are working in, including fullscreen apps and other desktops. Type a command, an SSH host or a folder path, press Enter , and the window becomes a terminal. Press the shortcut again to hide it and return focus to the previous app.
The app also records the commands you run, using a zsh hook. Because of that the same window works as a search over your command history, a list of servers you connect to, and a launcher for your Claude Code sessions.
FedTerm lives in the menu bar and has no Dock icon. The shortcut can be changed to any combination you like.
The shortcut is registered through Carbon, so Accessibility permissions are not needed.
You can rebind it in the settings. Click the field and press the keys you want. F-keys work without modifiers, other keys need at least one. If the combination is taken by another app, FedTerm says so and keeps the previous one.
The window floats above other apps and follows you between desktops.
It hides when you click outside of it. Pin it if you want it to stay open.
Its size and position are remembered between launches. If the saved position is off-screen, the window opens in the centre.
FedTerm looks at what you type and offers what fits:
Arrow keys move through the list, Enter runs the selected item, Esc clears the field. Matching is fuzzy, so a few letters are usually enough.
Saved commands and connections
Favourite commands. Star a command and it stays at the top of the window. You can rename it and mark it to open in its own tab when the app starts.
Saved SSH connections. Pin a server and give it a label. Recent connections are collected on their own and sorted by how recently and how often you used them.
Command history. Commands are grouped by age: last hour, six hours, today, yesterday, this week, older. Repeats inside a group are shown once. For the last hour there is also a short summary with counts per tool ( git , docker , ssh , npm and so on) and clickable server names.
Named sessions. Drop a folder on the window or paste a path, then give the session a name. FedTerm starts it with its own session id and later resumes that exact session instead of creating a new one.
All past sessions. FedTerm reads ~/.claude/projects , takes the summary or first message of each transcript together with its folder, and lists your sessions grouped by project and by date, with a search field. Results are cached by file modification time, so the list opens without delay.
Clean environment. CLAUDE* and ANTHROPIC* variables are removed from new tabs. A session started from FedTerm behaves like a standalone one even if FedTerm itself was launched from Claude Code.
⌘T opens a tab, ⌘W closes it, ⌘1 – ⌘9 switch between tabs, ⌘[ and ⌘] move left and right.
Tabs can be dragged to reorder and renamed with a double click.
Open tabs are restored on the next launch. SSH tabs reconnect automatically.
The tab title follows the shell: current folder, the title set by the running program, or the server name.
If something is still running, FedTerm asks for confirmation before quitting.
A command can be bound to ⌃1 – ⌃9 together with a working directory. These shortcuts only fire while the window is open, so they do not affect other apps. Multi-line scripts are saved to a file and run through bash, so loops and conditionals work as written.
There are 13 built-in themes: the Terminal.app palette, One Dark, Dracula, Nord, Gruvbox, both Solarized variants, Tokyo Night, Catppuccin Mocha, Monokai, GitHub Light and two neutral greys. Any theme can be copied and edited: background, text, caret, selection, all 16 ANSI colours and background transparency over the blur. Font family, size and weight are configurable. There is also a thin strokes option for crisper text and a slider that darkens the glass over light wallpapers. Changes apply to open terminals immediately.
The terminal is SwiftTerm : xterm-256color, truecolor and mouse support. ⌘ -click opens links, file paths and OSC 8 hyperlinks. The scrollbar gutter is removed, so full-width programs like mc and htop use the entire width. ⌃C is written directly to the pty, so it always interrupts. Shortcuts are matched by physical key code and keep working on Cyrillic and other non-Latin layouts. The interface is in Russian on Russian systems and in English everywhere else.
macOS 13 Ventura or newer, Apple Silicon or Intel
Xcode Command Line Tools with Swift 5.9 or newer
zsh, the default shell on macOS. Command history is captured through it
The claude CLI in your PATH for the Claude Code features
git clone https://github.com/feddot2517/fedterm.git
cd fedterm
make bundle # builds the release binary and assembles dist/FedTerm.app
open dist/FedTerm.app
make run does both steps at once. make dev runs the app from source without building a bundle.
The binary is signed ad-hoc, so Gatekeeper will block the first launch. Right-click the app and choose Open , or remove the quarantine flag:
xattr -dr com.apple.quarantine dist/FedTerm.app
Move the app to /Applications and add it to System Settings → General → Login Items to have it running all the time.
Shortcut
Action
⌥ Space
show or hide the window (rebindable)
Enter
run the selected item
↑ ↓
move through results and history
Esc
clear the field
⌘T / ⌘W
open or close a tab
⌘1 – ⌘9
switch to a tab
⌘[ / ⌘]
previous or next tab
⌃1 – ⌃9
run a custom command
⌘+ / ⌘- / ⌘0
font size up, down, reset
⌘ -click
open a link or path from the output
⌃C
interrupt the running program
Data
All files are local, in ~/Library/Application Support/FedTerm/ :
The app has no networking code, so nothing is uploaded and there is no telemetry. History is a plain text file and can be deleted at any time.
Note that commands are stored exactly as typed. If you paste a token or password inline, it will be in the file. Treat it like ~/.zsh_history .
Sources/FedTerm/
├─ main.swift, AppDelegate.swift startup, menu bar, window shortcuts
├─ SpotlightPanel.swift floating window, blur, saved position
├─ ContentView.swift, TabsModel.swift tab bar, reordering, persisted state
├─ HomeView.swift search field, results, history
├─ TerminalSession.swift SwiftTerm subclass, shell process, ⌘-click
├─ HistoryStore.swift reading and watching history.jsonl
├─ ShellIntegration.swift generated zsh config with the preexec hook
├─ ClaudeStore.swift, ClaudeUI.swift Claude Code sessions and their browser
├─ AutomationsStore.swift, AutomationEditor.swift custom ⌃1–9 commands
├─ FavoritesStore.swift favourite commands and autostart
├─ Theme.swift, ThemeEditor.swift, SettingsUI.swift themes, fonts, settings
├─ HotkeyManager.swift, HotkeyRecorder.swift global shortcut and rebinding
├─ Models.swift SSH parsing and command classification
└─ L10n.swift Russian and English strings
Credits
SwiftTerm by Miguel de Icaza does the terminal emulation.
Colour schemes are taken from One Dark, Dracula, Nord, Gruvbox, Solarized, Tokyo Night, Catppuccin and Monokai.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
