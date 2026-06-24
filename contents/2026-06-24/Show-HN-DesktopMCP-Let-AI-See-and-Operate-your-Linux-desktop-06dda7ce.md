---
source: "https://github.com/varbhat/desktopmcp"
hn_url: "https://news.ycombinator.com/item?id=48666602"
title: "Show HN: DesktopMCP – Let AI See and Operate your Linux desktop"
article_title: "GitHub - varbhat/desktopmcp: MCP server for the Linux desktop 🐧🤖 · GitHub"
author: "varbhat"
captured_at: "2026-06-24T23:24:41Z"
capture_tool: "hn-digest"
hn_id: 48666602
score: 2
comments: 1
posted_at: "2026-06-24T22:59:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: DesktopMCP – Let AI See and Operate your Linux desktop

- HN: [48666602](https://news.ycombinator.com/item?id=48666602)
- Source: [github.com](https://github.com/varbhat/desktopmcp)
- Score: 2
- Comments: 1
- Posted: 2026-06-24T22:59:18Z

## Translation

タイトル: Show HN: DesktopMCP – AI に Linux デスクトップを表示および操作させます
記事のタイトル: GitHub - varbhat/desktopmcp: Linux デスクトップ用の MCP サーバー 🐧🤖 · GitHub
説明: Linux デスクトップ用の MCP サーバー 🐧🤖 。 GitHub でアカウントを作成して、varbhat/desktopmcp の開発に貢献してください。

記事本文:
GitHub - varbhat/desktopmcp: Linux デスクトップ用の MCP サーバー 🐧🤖 · GitHub
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
ヴァルバート
/
デスクトップmcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

コード 「その他のアクション」メニューを開く フォルダーとファイル
2 コミット 2 コミット .github/ workflows .github/ workflows src src .envrc .envrc .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md flake.lock flake.lock flake.nix flake.nix すべてのファイルを表示リポジトリ ファイルのナビゲーション
Linux デスクトップ用の MCP サーバー。これにより、AI モデルが Linux デスクトップにアクセスできるようになります。
Desktopmcp は、サンドボックス化されたデスクトップ操作用の XDG デスクトップ ポータル、セマンティック UI の理解用の AT-SPI、および低レベルのシステム アクセス用の D-Bus の 3 つのシステム インターフェイスを介して AI アシスタントを Linux デスクトップに接続します。 Model Context Protocol を介して 144 のツールを公開し、AI が画面上の内容を確認し、UI 構造を理解し、カーソルを移動し、ボタンをクリックし、テキストを入力し、ファイルを管理し、デスクトップ サービスと対話できるようにします。
AI モデルはデスクトップを認識しません。ウィンドウを見ることも、ボタンのラベルを読むことも、メニュー項目をクリックすることも、デスクトップを制御することもできません。
Desktopmcp は、AI モデルにデスクトップと対話するための 2 つの補完的な方法を提供することで、この問題を解決します。
ビジュアル : スクリーンショット、スクリーン キャプチャ、カラー ピッキング、マウスおよびキーボード入力
セマンティック : 完全なアクセシビリティ ツリー (AT-SPI)。すべての UI 要素に名前、役割、位置、および AI が直接呼び出すことができるアクションのセットが含まれます。
セマンティック パスは、AI がピクセルをスキャンして「保存」の場所を見つける代わりに、find_element(role="push button", name="Save") を呼び出すことができることを意味します。テキスト コンテンツを読み取り、どの要素がフォーカスされているかを確認し、UI ツリーを移動し、任意の要素で「クリック」などのアクションを実行できます。これらすべてをスクリーンショットを必要とせずに実行できます。
すべてが XDG デスクトップ ポータルを経由するため、ユーザーには権限ダイアログが表示され、単一のバイナリが GNOME、KDE、Sway、または任意の Wayland コンポジター上で実行されます。
144 個のツールが 5 つのグループに分類されています。
AIモデル（クロードなど）
|
| MCPプロ

tocol (stdio または HTTP 経由の JSON-RPC)
v
Desktopmcp (Rust、非同期、単一バイナリ)
|
|--- XDG デスクトップ ポータル (ashpd) --> スクリーンショット、入力、ファイル、設定など...
|--- AT-SPI バス (zbus) --> UI ツリー、要素アクション、テキスト、イベント
|--- D-Bus セッション/システム (zbus) --> 任意のサービス アクセス
|--- PipeWire (pipewire-rs) --> スクリーン キャプチャ フレーム
v
Linux デスクトップ (GNOME / KDE / Sway / ...)
XDG ポータルは、すべての機密操作 (画面キャプチャ、入力挿入、ファイル アクセス) がユーザー向けの許可ダイアログを通過することを保証します。ユーザーの同意がなければAIは動作できません。
Wayland コンポジターを備えた Linux (GNOME 40 以降、KDE Plasma 5.27 以降、Sway など)
xdg-desktop-portal およびデスクトップ固有のバックエンド ( xdg-desktop-portal-gnome 、 -kde 、 -wlr )
AT-SPI2 ( at-spi2-core 、通常はアクセシビリティ サポートを備えたデスクトップにプレインストールされます)
D-Bus セッション バス (最新のすべての Linux デスクトップに存在)
システム ライブラリ: libdbus 、 libpipewire-0.3 、 at-spi2-core 、 pkg-config
LLVM/Clang (pipewire-sys bindinggen 用)
nix run github:varbhat/desktopmcp -- --t http # または stdio
またはクローンを作成してローカルにビルドします。リポジトリには、すべてのシステム依存関係を自動的に提供する flake.nix が含まれています。
git clone < リポジトリ URL >
cd デスクトップ mcp
ニックス開発
カーゴビルド --release
AppImage をリリースする (推奨)
Github Actions を使用して、AppImage を自動的にビルドおよびリリースします。リリース ページから最新の AppImage を入手します。それを実行可能としてマークして実行します。これは、nix 派生とそのすべての依存関係を単一ファイルの実行可能ファイルにバンドルする nix-appimage を使用して構築されているため、AppImage のサイズは非常に大きくなります (これは今後のリリースで最適化される予定ですが、試してみることができます)。
システムの依存関係をインストールしてから、以下をビルドします。
# Debian / Ubuntu
sudo apt install libdbus-1-dev libpipewire-0.3-dev at-sp

i2コア\
pkg-config libclang-dev
#フェドーラ
sudo dnf install dbus-devel Pipewire-devel at-spi2-core-devel \
pkg-config clang-devel
# アーチ
sudo pacman -S dbus Pipewire at-spi2-core pkgconf Clang
# ビルド
カーゴビルド --release
バイナリは target/release/desktopmcp にあります。
# stdio モード (デフォルト; MCP クライアントがこれを実行します)
デスクトップmcp
# HTTP モード -- リモート mcp 用 (事前に実行する必要があります)
Desktopmcp --transport http --bind 127.0.0.1:3000
stdio モードでは、MCP メッセージは stdin/stdout (JSON-RPC) 経由で交換されます。ログは標準エラー出力に保存されます。 HTTP モードでは、サーバーはサーバー送信イベントによるストリーミング可能な HTTP を使用して http://<bind>/mcp をリッスンします。バイナリ パスを指定することで、stdio モードでデスクトップ mcp を使用するように MCP クライアントを構成できます。または、desktopmcp を HTTP リモート mcp モードで実行し、URL を指定してそれを使用するように MCP クライアントを構成することもできます。
AI に質問します。「デスクトップでどのウィンドウが開いていますか?」
AI は get_window_list を呼び出し、タイトル、アプリケーション名、位置、サイズを含むすべてのウィンドウの構造化リストを取得します。スクリーンショットは必要ありません。
AI に質問する: 「Firefox の [保存] ボタンをクリックしてください」
1. find_element(role="プッシュボタン", name="保存", app_name="Firefox")
-> 位置と利用可能なアクションを含む要素 ID を返します
2. atspi_do_action(id="...", action_name="クリック")
-> アクセシビリティ API を介してボタンがクリックされました
AI に質問します。「スクリーンショットを撮って、何が表示されているか教えてください」
1. start_session(devices=["pointer"]、with_screencast=true)
-> ユーザーは許可ダイアログを一度承認します
2. take_screenshot(session_id="...", format="jpeg")
-> AI は Base64 でエンコードされた画像を受信します
AI に質問する: 「ログイン フォームにメール アドレスを入力してください」
1. find_element(role="エントリ", name="電子メール")
-> テキストフィールドを検索します
2. type_into(id="...", text="user@example.com")
-> テキストは、EditableText アクセシビリティを介して入力されます。

表面
AIに尋ねる:「ダウンロードが完了したら通知を送ってください」
1. send_notification(id="dl-done", title="ダウンロード完了", body="file.zip の準備ができました")
→デスクトップ通知が表示される
ツールリファレンス
これらのツールには、 start_session で作成されたアクティブなセッションが必要です。ユーザーには 1 回限りの許可ダイアログが表示されます。
これらのツールは、サンドボックス化された XDG デスクトップ ポータル API を使用します。ほとんどはセッションなしで機能します。
.desktop アプリケーション ランチャーをインストールして管理します。
デスクトップ上のすべての UI 要素へのセマンティック アクセス。スクリーンショットや座標の推測は必要ありません。
ツリーの走査と要素の検査:
コンポーネント (ジオメトリ、スクロール、フォーカス):
値、選択、テーブル、ハイパーリンク、ドキュメント、画像:
セッションまたはシステム バス上の任意の D-Bus サービスに直接アクセスします。 AI は、任意のデスクトップ サービスをイントロスペクトし、対話することができます。
nix Development # すべての依存関係を含む開発シェルに入る
カーゴビルド # コンパイル
貨物クリッピー # 糸くず
カーゴ実行 # stdio モードで実行
カーゴ実行 -- -t http # HTTP モードで実行
ポータルが実行されていることを確認します。
systemctl --user status xdg-desktop-portal
systemctl --ユーザーステータスパイプワイヤー
デバッグログを有効にします。
RUST_LOG=デバッグカーゴラン
生の MCP メッセージを使用してテストします。
echo ' {"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1"}}} ' \
|カーゴ ラン --quiet 2> /dev/null
トラブルシューティング
許可ダイアログが表示されない
systemctl --user restart xdg-desktop-portal
PipeWire 接続が失敗する
PipeWire が実行されていることを確認します。
systemctl --ユーザーステータスパイプワイヤー
systemctl --user パイプワイヤを再起動します
AT-SPI はアプリケーションを返しません
AT-SPI バスが実行中であり、アクセシビリティが有効になっていることを確認します。
# AT-SPI ステータスを確認する
dbus-send --session --print-reply --dest=org.a11y.Bus \
/org/a11y/bus org.freedesktop.DBus.Properties.GetAll \
文字列:

org.a11y.ステータス
# 無効になっている場合は、(Gnome 上で) アクセシビリティを有効にします
gsettings set org.gnome.desktop.interface toolkit-accessibility true
マウス/キーボード入力は効果がありません
最初に必要なデバイス ( ["keyboard", "pointer"] ) を指定して start_session を呼び出します。
許可ダイアログが表示されたら同意します
セッションが有効な session_id でまだアクティブであることを確認します。
Desktopmcp は、XDG デスクトップ ポータルのセキュリティ モデルを中心に設計されています。
すべての機密性の高い操作には、システム ダイアログを介した明示的なユーザーの同意が必要です
サーバーは Flatpak およびその他のサンドボックス環境内で動作します
D-Bus および AT-SPI アクセスは標準の Linux デスクトップ権限に従います
ユーザーはいつでもアクセスを拒否または取り消すことができます
とはいえ、このサーバーにより、AI モデルはデスクトップを大幅に制御できるようになります。信頼できるモデルで使用し、AI の動作を確認し、制限のない AI が意図しないアクションを実行する可能性があることに注意してください。
コンポーネント
図書館
バージョン
MCPプロトコル
rmcp
1.7
XDG ポータル
アシュプド
0.13
Dバス/AT-SPI
zbus + zvariant
5
スクリーンキャプチャ
パイプワイヤー-RS
0.10
非同期ランタイム
東京
1
画像エンコーディング
画像-RS
0.25
連載
セルデ + セルデ_json
1
CLI
拍手する
4
ライセンス
Linux デスクトップ用の MCP サーバー 🐧🤖
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
デスクトップmcp v0.1.0-1782341241
最新の
2026 年 6 月 24 日
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

MCP server for the Linux desktop 🐧🤖 . Contribute to varbhat/desktopmcp development by creating an account on GitHub.

GitHub - varbhat/desktopmcp: MCP server for the Linux desktop 🐧🤖 · GitHub
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
varbhat
/
desktopmcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github/ workflows .github/ workflows src src .envrc .envrc .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md flake.lock flake.lock flake.nix flake.nix View all files Repository files navigation
MCP server for the Linux desktop. It gives AI models access to the Linux desktop.
desktopmcp connects AI assistants to the Linux desktop through three system interfaces: XDG Desktop Portals for sandboxed desktop operations, AT-SPI for semantic UI understanding, and D-Bus for low-level system access. It exposes 144 tools over the Model Context Protocol, letting an AI see what's on screen, understand UI structure, move cursor, click buttons, type text, manage files, and interact with desktop services.
AI models are blind to the desktop. They can't see a window, read a button label, click a menu item, or control the desktop.
desktopmcp solves this by giving AI models two complementary ways to interact with the desktop:
Visual : screenshots, screen capture, color picking, mouse and keyboard input
Semantic : the full accessibility tree (AT-SPI), where every UI element has a name, role, position, and set of actions the AI can invoke directly
The semantic path means an AI can call find_element(role="push button", name="Save") instead of scanning pixels to find where "Save" might be. It can read text content, check which element is focused, traverse the UI tree, and perform actions like "click" on any element -- all without needing a screenshot.
Everything goes through XDG Desktop Portals, so the user gets permission dialogs, and a single binary runs on GNOME, KDE, Sway, or any Wayland compositor.
144 tools organized into five groups:
AI model (Claude, etc.)
|
| MCP protocol (JSON-RPC over stdio or HTTP)
v
desktopmcp (Rust, async, single binary)
|
|--- XDG Desktop Portals (ashpd) --> screenshots, input, files, settings, ...
|--- AT-SPI bus (zbus) --> UI tree, element actions, text, events
|--- D-Bus session/system (zbus) --> arbitrary service access
|--- PipeWire (pipewire-rs) --> screen capture frames
v
Linux desktop (GNOME / KDE / Sway / ...)
XDG Portals ensure every sensitive operation (screen capture, input injection, file access) goes through a user-facing permission dialog. The AI cannot act without the user granting consent.
Linux with a Wayland compositor (GNOME 40+, KDE Plasma 5.27+, Sway, or similar)
xdg-desktop-portal and a desktop-specific backend ( xdg-desktop-portal-gnome , -kde , -wlr )
AT-SPI2 ( at-spi2-core , typically pre-installed on any desktop with accessibility support)
D-Bus session bus (present on all modern Linux desktops)
System libraries: libdbus , libpipewire-0.3 , at-spi2-core , pkg-config
LLVM/Clang (for pipewire-sys bindgen)
nix run github:varbhat/desktopmcp -- --t http # or stdio
Or clone and build locally. The repository includes a flake.nix that provides all system dependencies automatically:
git clone < repo-url >
cd desktopmcp
nix develop
cargo build --release
Release AppImage (recommended)
We use Github Actions to build and release AppImages automatically. Grab the latest AppImage from the releases page . Mark it as executable and run it. It's built using nix-appimage which bundles the nix derivation and all its dependencies into a single-file executable, and hence the size of the AppImage is very big (It's something we'll optimize in upcoming releases—but hey, it lets you try it!).
Install system dependencies, then build:
# Debian / Ubuntu
sudo apt install libdbus-1-dev libpipewire-0.3-dev at-spi2-core \
pkg-config libclang-dev
# Fedora
sudo dnf install dbus-devel pipewire-devel at-spi2-core-devel \
pkg-config clang-devel
# Arch
sudo pacman -S dbus pipewire at-spi2-core pkgconf clang
# Build
cargo build --release
The binary is at target/release/desktopmcp .
# stdio mode (default; Your MCP Client will do this for you)
desktopmcp
# HTTP mode -- for remote-mcp (You need to run this beforehand)
desktopmcp --transport http --bind 127.0.0.1:3000
In stdio mode, MCP messages are exchanged over stdin/stdout (JSON-RPC). Logs go to stderr. In HTTP mode, the server listens at http://<bind>/mcp using Streamable HTTP with Server-Sent Events. You can configure your MCP Client to use desktopmcp in stdio mode by specifying the binary path. Or you can run the desktopmcp in HTTP remote-mcp mode and configure your MCP Client to use it by specifying the URL.
Ask AI : "What windows are open on my desktop?"
The AI calls get_window_list and gets back a structured list of every window with title, application name, position, and size -- no screenshot needed.
Ask AI : "Click the Save button in Firefox"
1. find_element(role="push button", name="Save", app_name="Firefox")
-> returns element ID with position and available actions
2. atspi_do_action(id="...", action_name="click")
-> button is clicked via the accessibility API
Ask AI : "Take a screenshot and tell me what you see"
1. start_session(devices=["pointer"], with_screencast=true)
-> user approves the permission dialog once
2. take_screenshot(session_id="...", format="jpeg")
-> AI receives a base64-encoded image
Ask AI : "Type my email address into the login form"
1. find_element(role="entry", name="Email")
-> finds the text field
2. type_into(id="...", text="user@example.com")
-> text is entered via the EditableText accessibility interface
Ask AI : "Send me a notification when the download finishes"
1. send_notification(id="dl-done", title="Download Complete", body="file.zip is ready")
-> desktop notification appears
Tool reference
These tools require an active session created with start_session . The user sees a one-time permission dialog.
These tools use sandboxed XDG Desktop Portal APIs. Most work without a session.
Install and manage .desktop application launchers.
Semantic access to every UI element on the desktop. No screenshot or coordinate guessing required.
Tree traversal & element inspection:
Component (geometry, scroll, focus):
Value, Selection, Table, Hyperlinks, Document, Image:
Direct access to any D-Bus service on the session or system bus. The AI can introspect and interact with arbitrary desktop services.
nix develop # enter dev shell with all dependencies
cargo build # compile
cargo clippy # lint
cargo run # run in stdio mode
cargo run -- -t http # run in HTTP mode
Verify portals are running:
systemctl --user status xdg-desktop-portal
systemctl --user status pipewire
Enable debug logging:
RUST_LOG=debug cargo run
Test with a raw MCP message:
echo ' {"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1"}}} ' \
| cargo run --quiet 2> /dev/null
Troubleshooting
Permission dialog does not appear
systemctl --user restart xdg-desktop-portal
PipeWire connection fails
Check that PipeWire is running:
systemctl --user status pipewire
systemctl --user restart pipewire
AT-SPI returns no applications
Make sure the AT-SPI bus is running and accessibility is enabled:
# Check AT-SPI status
dbus-send --session --print-reply --dest=org.a11y.Bus \
/org/a11y/bus org.freedesktop.DBus.Properties.GetAll \
string:org.a11y.Status
# Enable accessibility (on Gnome) if disabled
gsettings set org.gnome.desktop.interface toolkit-accessibility true
Mouse/keyboard input has no effect
Call start_session first with the required devices ( ["keyboard", "pointer"] )
Accept the permission dialog when it appears
Verify the session is still active with a valid session_id
desktopmcp is designed around the XDG Desktop Portal security model:
Every sensitive operation requires explicit user consent via a system dialog
The server works inside Flatpak and other sandboxed environments
D-Bus and AT-SPI access follows standard Linux desktop permissions
The user can deny or revoke access at any time
That said, this server gives AI models significant control over your desktop. Use it with trusted models, review what the AI does, and be aware that an unrestricted AI could perform unintended actions.
Component
Library
Version
MCP protocol
rmcp
1.7
XDG Portals
ashpd
0.13
D-Bus / AT-SPI
zbus + zvariant
5
Screen capture
pipewire-rs
0.10
Async runtime
tokio
1
Image encoding
image-rs
0.25
Serialization
serde + serde_json
1
CLI
clap
4
License
MCP server for the Linux desktop 🐧🤖
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
desktopmcp v0.1.0-1782341241
Latest
Jun 24, 2026
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
