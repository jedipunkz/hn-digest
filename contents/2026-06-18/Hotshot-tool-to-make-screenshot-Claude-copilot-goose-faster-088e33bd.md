---
source: "https://github.com/kubestellar/hotshot"
hn_url: "https://news.ycombinator.com/item?id=48588203"
title: "Hotshot – tool to make screenshot –> Claude/copilot/goose faster"
article_title: "GitHub - kubestellar/hotshot: Hotkey-triggered screenshot injection into your last active terminal session · GitHub"
author: "andan02"
captured_at: "2026-06-18T18:56:21Z"
capture_tool: "hn-digest"
hn_id: 48588203
score: 1
comments: 1
posted_at: "2026-06-18T16:54:30Z"
tags:
  - hacker-news
  - translated
---

# Hotshot – tool to make screenshot –> Claude/copilot/goose faster

- HN: [48588203](https://news.ycombinator.com/item?id=48588203)
- Source: [github.com](https://github.com/kubestellar/hotshot)
- Score: 1
- Comments: 1
- Posted: 2026-06-18T16:54:30Z

## Translation

タイトル: Hotshot – スクリーンショットを作成するツール –> Claude/copilot/goose を高速化
記事のタイトル: GitHub - kubestellar/hotshot: 最後のアクティブなターミナル セッションへのホットキー トリガーのスクリーンショットの挿入 · GitHub
説明: ホットキーでトリガーされる、最後のアクティブなターミナル セッションへのスクリーンショットの挿入 - kubestellar/hotshot

記事本文:
GitHub - kubestellar/hotshot: 最後のアクティブなターミナル セッションへのホットキー トリガーのスクリーンショットの挿入 · GitHub
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
クベステラ
/
ホットショット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

ns
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
21 コミット 21 コミット ソース ソース リソース リソース スクリプト スクリプト .gitignore .gitignore ライセンス ライセンス Package.swift Package.swift README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
スクリーンショットを撮ると、それが端末に表示されます。それでおしまい。
hotshot は、スクリーンショットをキャプチャし、最後に使用していたターミナル セッションにファイル パスを自動的に貼り付ける小さな macOS メニュー バー アプリです。画像パスを入力として受け入れる、Claude Code、GitHub Copilot CLI、aider、OpenCode などの AI コーディング アシスタント用に構築されています。
あなたも私と同じなら、壊れた UI、奇妙なエラー メッセージ、正しく見えないダッシュボードなど、作業をデバッグするために常にスクリーンショットを使用するでしょう。通常は、スクリーンショットを撮り、ファイルを見つけ、パスをコピーし、ターミナルに切り替えて貼り付けます。ホットショットは、これらすべてを 1 回のキーストロークで実行します。または、Mac の組み込みスクリーンショット ショートカットを使用する場合は、キーストロークを必要としません。
ターミナルで Claude Code (または任意の AI CLI) を使用している
ブラウザに切り替えてバグを発見する
⇧⌘S を押す — おなじみの十字線が表示され、領域を選択します
ホットショットはスクリーンショットを保存し、ターミナルセッションにファイルパスを入力します
AI アシスタントが画像を読み取って支援を開始します
ネイティブ スクリーンショット (⌘⇧3、⌘⇧4、⌘⇧5) の場合:
フルスクリーン、リージョン、ウィンドウなど、いつもと同じ方法でスクリーンショットを撮ります。
ホットショットは新しいファイルを検出し、そのパスを最後のターミナル セッションに自動的に挿入します。
それだけです — すでに知っているショートカットを使い続けます
ファイルをドラッグする必要はありません。コピー＆ペーストのパスはありません。いいえ、「スクリーンショットがどこにいったのか調べさせてください。」
macOS 13 以降および Swift (Xcode または Xcode コマンド ライン ツールに付属) が必要です。
git clone https://github.com/kubestellar/hotsh

ot.git
CD ホットショット
迅速なビルド -c リリース
sudo cp .build/release/hotshot /usr/local/bin/
あとはそれを実行するだけです:
ホットショット
メニュー バーに小さなカメラのアイコンが表示されます。これがホットショットの実行中です。必要になるまで邪魔になりません。
ヒント: hotshot & を実行して、端末を保持しないようにバックグラウンドで実行します。
macOS は初めて次の 2 つの権限を要求します。
画面録画 - スクリーンショットを撮ることができます
アクセシビリティ - ターミナルにパスを入力できるようにする
[システム設定] > [プライバシーとセキュリティ] で両方を許可します。これを行う必要があるのは 1 回だけです。
ホットショットは、最後にクリックした端末を記憶します。これは次の 2 つの方法で機能します。
macOS 地域セレクターを開きます (⌘⇧4 と同じ十字線)
スクリーンショットを設定したフォルダーに保存します
ファイル パスを最後のアクティブなターミナル セッションに挿入します。
端末を前面に戻します
自動監視モード (デフォルトでオン):
ネイティブ macOS ショートカット (⌘⇧3、⌘⇧4、⌘⇧5) を使用してスクリーンショットを撮ります。
ホットショットはスクリーンショットフォルダー内の新しいファイルを検出します
最後のアクティブなターミナルセッションにパスを自動的に挿入します
パスは [括弧] 形式で挿入されます。これは、Claude Code などの AI CLI がファイルのドラッグ アンド ドロップに使用するのと同じ形式です。サーバー、クリップボードのハッキング、ブラウザの拡張機能はありません。
メニューバーのカメラアイコンをクリックします。すべてが設定可能です:
iTerm2 (推奨 — 信頼性の高いインジェクションのためにネイティブ AppleScript を使用します)
これらの AI アシスタントと連携します
イメージ ファイル パスを入力として受け入れる CLI ツール:
ファイルパスを貼り付けると画像を読み取ることができるツール
Mac の組み込みスクリーンショット ショートカットと競合しますか?
いいえ、彼らには効果があります。 macOS システム ショートカット (⌘⇧3、⌘⇧4、⌘⇧5) は引き続き通常どおり機能します。自動監視が有効になっている場合 (デフォルト)、ホットショットは新しいスクリーンショットを検出し、自動的に挿入します。ホットショットの

独自の ⇧⌘S ホットキーは、競合しない別のショートカットです。メニューバーから任意に変更できます。
ショートカットを押しても何も起こりませんでした。
ホットショットが実行されていることを確認してください (メニュー バーでカメラのアイコンを探してください)。また、ホットショットを起動してから少なくとも 1 回は端末ウィンドウをクリックしていることを確認してください。どの端末をターゲットにするかを認識する必要があります。
スクリーンショットはどこに行くのでしょうか?
デフォルトでは、Mac のスクリーンショットが保存される場所 (通常はデスクトップまたはダウンロード)。ホットショットは、macOS スクリーンショットの場所設定を自動的に読み取ります。メニューバー > 「スクリーンショットフォルダーの変更...」からオーバーライドできます。
パスはどのような形式で挿入されますか?
パスは角かっこで囲まれます: [/path/to/screenshot.png] 。これは、Claude Code や他の AI CLI がファイルをターミナルにドラッグ アンド ドロップするときに使用するのと同じ形式です。
SSH経由で動作しますか?
直接ではありません - ホットショットはローカル Mac 上で実行されます。リモート ワークフローについては、 Clipssh または Clipaste を確認してください。
VS Code の統合ターミナルで使用できますか?
まだです — VS Code のターミナルはスタンドアロン アプリではありません。 VS Code の場合は、 vscode-terminal-image-paste を試してください。
ネイティブ macOS バイナリにコンパイルされた単一の Swift ファイル (約 650 行)。依存関係なし — フレームワーク、パッケージ、ランタイム要件はありません。 Apple の組み込み API のみ:
NSEvent ホットキーのグローバル モニター
端末のフォーカスを追跡するための NSWorkspace 通知
新しいスクリーンショットを自動検出するための DispatchSource ファイル システム ウォッチャー
実際のキャプチャ用の /usr/sbin/screencapture
NSAppleScript によるターミナル セッションへのパスの挿入
UserDefaults で設定を保持します
ホットキーでトリガーされる、最後のアクティブなターミナル セッションへのスクリーンショットの挿入
Apache-2.0ライセンス
行動規範
行動規範
セキュリティポリシー
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
C

カスタムプロパティ
星
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

Hotkey-triggered screenshot injection into your last active terminal session - kubestellar/hotshot

GitHub - kubestellar/hotshot: Hotkey-triggered screenshot injection into your last active terminal session · GitHub
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
kubestellar
/
hotshot
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
21 Commits 21 Commits Sources Sources resources resources scripts scripts .gitignore .gitignore LICENSE LICENSE Package.swift Package.swift README.md README.md View all files Repository files navigation
Take a screenshot, and it lands in your terminal. That's it.
hotshot is a tiny macOS menu bar app that captures a screenshot and automatically pastes the file path into whichever terminal session you were last using. Built for AI coding assistants like Claude Code , GitHub Copilot CLI, aider, and OpenCode that accept image paths as input.
If you're like me, you use screenshots constantly to debug your work — a broken UI, a weird error message, a dashboard that doesn't look right. Normally you'd screenshot it, find the file, copy the path, switch to your terminal, paste it in. hotshot does all of that in one keystroke — or zero keystrokes if you use your Mac's built-in screenshot shortcuts.
You're working in Claude Code (or any AI CLI) in your terminal
You switch to a browser and spot a bug
Press ⇧⌘S — the familiar crosshair appears, select the area
hotshot saves the screenshot and types the file path into your terminal session
Your AI assistant reads the image and starts helping
With native screenshots (⌘⇧3, ⌘⇧4, ⌘⇧5):
Take a screenshot the way you always do — full screen, region, window, whatever
hotshot detects the new file and automatically injects its path into your last terminal session
That's it — keep using the shortcuts you already know
No dragging files around. No copy-pasting paths. No "here let me find where that screenshot went."
Requires macOS 13+ and Swift (comes with Xcode or Xcode Command Line Tools).
git clone https://github.com/kubestellar/hotshot.git
cd hotshot
swift build -c release
sudo cp .build/release/hotshot /usr/local/bin/
Then just run it:
hotshot
A small camera icon appears in your menu bar — that's hotshot running. It stays out of your way until you need it.
Tip: Run hotshot & to background it so it doesn't hold your terminal.
macOS will ask for two permissions the first time:
Screen Recording — so it can take screenshots
Accessibility — so it can type the path into your terminal
Grant both in System Settings > Privacy & Security. You only have to do this once.
hotshot remembers which terminal you last clicked on. It works in two ways:
Opens the macOS region selector (same crosshair as ⌘⇧4)
Saves the screenshot to your configured folder
Injects the file path into your last active terminal session
Brings the terminal back to the front
Auto-watch mode (on by default):
You take a screenshot with any native macOS shortcut (⌘⇧3, ⌘⇧4, ⌘⇧5)
hotshot detects the new file in your screenshot folder
Automatically injects the path into your last active terminal session
Paths are injected in [bracket] format — the same format AI CLIs like Claude Code use for drag-and-dropped files. No servers, no clipboard hacks, no browser extensions.
Click the camera icon in your menu bar. Everything is configurable:
iTerm2 (recommended — uses native AppleScript for reliable injection)
Works with these AI assistants
Any CLI tool that accepts image file paths as input:
Any tool where you can paste a file path and it reads the image
Does it conflict with the Mac's built-in screenshot shortcuts?
No — it works with them. macOS system shortcuts (⌘⇧3, ⌘⇧4, ⌘⇧5) still work as normal. With auto-watch enabled (the default), hotshot detects the new screenshot and injects it automatically. hotshot's own ⇧⌘S hotkey is a separate shortcut that doesn't conflict. You can change it to anything you want from the menu bar.
I pressed the shortcut and nothing happened.
Make sure hotshot is running (look for the camera icon in your menu bar). Also make sure you've clicked on a terminal window at least once since launching hotshot — it needs to know which terminal to target.
Where do the screenshots go?
By default, wherever your Mac saves screenshots (usually Desktop or Downloads). hotshot reads your macOS screenshot location setting automatically. You can override it from the menu bar > "Change screenshot folder..."
What format is the path injected in?
Paths are wrapped in square brackets: [/path/to/screenshot.png] . This is the same format Claude Code and other AI CLIs use when you drag and drop a file into the terminal.
Does it work over SSH?
Not directly — hotshot runs on your local Mac. For remote workflows, check out clipssh or clipaste .
Can I use it with VS Code's integrated terminal?
Not yet — VS Code's terminal isn't a standalone app. For VS Code, try vscode-terminal-image-paste .
Single Swift file (~650 lines), compiled to a native macOS binary. Zero dependencies — no frameworks, no packages, no runtime requirements. Just Apple's built-in APIs:
NSEvent global monitors for the hotkey
NSWorkspace notifications to track terminal focus
DispatchSource file system watcher for auto-detecting new screenshots
/usr/sbin/screencapture for the actual capture
NSAppleScript to inject the path into terminal sessions
UserDefaults to persist your preferences
Hotkey-triggered screenshot injection into your last active terminal session
Apache-2.0 license
Code of conduct
Code of conduct
Security policy
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
