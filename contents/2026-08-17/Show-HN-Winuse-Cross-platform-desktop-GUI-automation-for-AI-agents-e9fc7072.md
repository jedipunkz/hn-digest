---
source: "https://github.com/lgxz/winuse"
hn_url: "https://news.ycombinator.com/item?id=49327992"
title: "Show HN: Winuse – Cross-platform desktop GUI automation for AI agents"
article_title: "GitHub - lgxz/winuse: App Control for Agents · GitHub"
image: "https://opengraph.githubassets.com/bb61c3fc5a564c7ce845bd745db2a65389f7ed8f3a18000e9477910e635d3e7a/lgxz/winuse"
author: "lgxz"
captured_at: "2026-08-17T09:31:05Z"
capture_tool: "hn-digest"
hn_id: 49327992
score: 2
comments: 0
posted_at: "2026-08-17T08:47:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Winuse – Cross-platform desktop GUI automation for AI agents

- HN: [49327992](https://news.ycombinator.com/item?id=49327992)
- Source: [github.com](https://github.com/lgxz/winuse)
- Score: 2
- Comments: 0
- Posted: 2026-08-17T08:47:23Z

## Translation

タイトル: Show HN: Winuse – AI エージェント向けのクロスプラットフォーム デスクトップ GUI オートメーション
記事のタイトル: GitHub - lgxz/winuse: エージェント向けアプリ制御 · GitHub
説明: エージェント用のアプリケーション制御。 GitHub でアカウントを作成して、lgxz/winuse の開発に貢献してください。

記事本文:
GitHub - lgxz/winuse: エージェント向けアプリ制御 · GitHub
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
lgxz
/
ウィンズ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット .cargo .cargo crates crates docs docs 例 例 python/ winuse python/ winuse スキル/ winuse スキル/ winuse テスト テスト .gitignore .gitignore AGENTS.md AGENTS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Makefile Makefile README.md R

EADME.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Rust コア、Python API、エージェント指向の CLI を使用したネイティブ GUI 自動化。
winuse は、macOS アクセシビリティと Windows UI オートメーションを同じ背後でサポートします。
バックエンドの特性。
macOS では、winuse には macOS 15.2 以降が必要です。アクセシビリティ権限を付与する
winuse を実行するターミナルまたは Python ホストに接続します。スクリーンショットも追加
画面録画許可が必要です。
Windows 10 または 11 では、UI オートメーションと可視領域のスクリーンショットは必要ありません。
個別のプライバシー許可。 Windows Integrity Isolation (UIPI) が入力をブロックする可能性がある
管理者特権のアプリケーションにアクセスするため、winuse は、管理者と同じ整合性レベルで実行します。
ターゲット。セキュア デスクトップ ウィンドウはサポートされていません。
python3 -m venv .venv
.venv/bin/pip maturin pytest をインストールします
.venv/bin/maturin 開発
.venv/bin/pytest
Python API
システム クリップボード経由でテキストを移動します (コピー/貼り付けのショートカットと組み合わせます)。
デスクトップ。 Clipboard_write ( "Hello" ) # cmd+v を押して貼り付けます
クリップ = デスクトップ。クリップボード_読み取り()
print (クリップ.種類、クリップ.テキスト、クリップ.ファイル) #text |ファイル |画像 |空の
デスクトップ。 Clipboard_read ( save_image = "shot.png" ) # 画像クリップボードをエクスポートする
ID によるウィンドウの制御:
ウィンドウ = デスクトップ 。ウィンドウ ( app = "TextEdit" )
窓の活性化 （）;窓の最小化 ();窓の復元する （）;窓の最大化();窓の閉じる（）
デスクトップ。 windows () # すべての実際のウィンドウ;それぞれが.minimized
自分でスナップショットをポーリングする代わりに、非同期 UI 状態を待ちます。
送信 = ウィンドウ。 find ( 役割 = "ボタン" 、名前 = "送信" )
送ったら。 wait_until ( "有効" 、タイムアウト = 10 ):
送信します。 ()をクリックしてください
窓のfind ( role = "static-text" 、 name = "Status" )。 wait_until ( "テキスト" 、テキスト = "完了" 、タイムアウト = 30 )
winuseインポートデスクトップから
デスクトップ = デスクトップ ()
そうでない場合はデスクトップ。 is_authorized():
レイズ

e RuntimeError ( "プラットフォーム アクセシビリティ バックエンドが利用できません" )
ウィンドウ = デスクトップ 。ウィンドウ ( app = "TextEdit" )
エディタ = ウィンドウ 。 find ( role = "テキストエリア" )
編集者。 fill ( "winuse からこんにちは" )
編集者。 (「入力」) を押します
編集者。 fill ( "query" , submit = True ) # 入力して Enter キーを押します (検索ボックス)
編集者。 type_text ( "カーソルに挿入" ) # type と fill: 挿入しますが、置換はしません
編集者。 press ( "s" , 修飾子 = [ "cmd" ]) # ショートカット; Windows では cmd は Ctrl です
編集者。クリック ( double = True ) # ダブルクリック
編集者。 hover () # クリックせずにホバーします
窓のfind ( role = "リスト項目" 、 name = "ファイル" )。 rag_to ( name = "Folder" ) # ターゲット上にドラッグします
サブツリー = ウィンドウ 。 find ( 役割 = "グループ" 、名前 = "サイドバー" )。スナップショット ( max_ Depth = 5 )
print ( サブツリー . テキスト ())
サブツリーを観察してセマンティックな追加と削除を確認します。
メッセージ = ウィンドウ 。 find ( 役割 = "リスト" 、名前 = "メッセージ" )
メッセージ付き。 (間隔 = 0.25) の変化を観察します。
while 変更 := 変更します。次 (タイムアウト = 30):
print (変更の種類、変更の要素)
要素によって公開されるすべてのアクセシビリティ属性を検査します。
点検＝窓。 find ( role = "static-text" 、 name = "Message" )。検査する（）
印刷（検査・本文（））
完全なネイティブ属性検査は両方のプラットフォームで機能します。macOS は、
要素の AX 属性、Windows は要素の UIA プロパティを報告します
(名前、オートメーション ID、クラス名、値、パターンなど) をサポートします。
要素が占める表示画面領域をキャプチャします。
ボタン = ウィンドウ。 find ( 役割 = "ボタン" 、名前 = "送信" )
png = ボタン 。スクリーンショット ( "send-button.png" 、パディング = 8 )
このメソッドは常に PNG バイトを返します。パスを渡すと同じことが書き込まれます
そのファイルにバイトを追加します。キャプチャされた領域には、現在表示されているものが反映されます。
画面（任意の w を含む）

要素を覆う窓。
要素は、存続期間の長いネイティブ ハンドルではなく、セレクターを保存します。それぞれのアクションが解決します
新しいアクセシビリティ スナップショットに対する要素。
カーゴビルド --release -p winuse-cli
エージェント スキル (SKILL.md + 参照 + ホスト バイナリ) をインストールします。
~/.agents/skills/winuse 、または 3 つすべてを含む配布可能な zip をパッケージ化する
プラットフォームバイナリ:
make install # DEST=... 場所を上書きする
パッケージ # を作成 → target/dist/winuse-<バージョン>-skill.zip
コマンドは通常のフラグを受け入れ、常に JSON エンベロープを返します。
ターゲット/リリース/Winuseドクター
ターゲット/リリース/winuse ウィンドウ
ターゲット/リリース/winuse スナップショット --window 123
target/release/winuse click --window 123 --role button --name Send --exact
target/release/winuse fill --window 123 --role text-area --text "こんにちは"
target/release/winuse press --window 123 --key Enter
target/release/winuse スクロール --window 123 --direction up --amount 500
リポジトリには、 skill/winuse に配布可能なエージェント スキルも含まれています。
バンドルされている実行可能ファイルは winuse-cli から構築されます。
winuse は以下からインスピレーションを受け、参照しました。
エージェントブラウザ (Vercel Labs) —
安定したセレクターを使用して、エージェントに対して UI を意味論的に駆動するアプローチ
JSON ツールの応答は、ネイティブ デスクトップ アプリ向けの winuse の設計を形作りました。
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

App Control for Agents. Contribute to lgxz/winuse development by creating an account on GitHub.

GitHub - lgxz/winuse: App Control for Agents · GitHub
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
lgxz
/
winuse
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits .cargo .cargo crates crates docs docs examples examples python/ winuse python/ winuse skills/ winuse skills/ winuse tests tests .gitignore .gitignore AGENTS.md AGENTS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Makefile Makefile README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Native GUI automation with a Rust core, a Python API, and an agent-oriented CLI.
winuse supports macOS Accessibility and Windows UI Automation behind the same
Backend trait.
On macOS, winuse requires macOS 15.2 or later. Grant Accessibility permission
to the terminal or Python host that runs winuse. Screenshots additionally
require Screen Recording permission.
On Windows 10 or 11, UI Automation and visible-region screenshots need no
separate privacy permission. Windows integrity isolation (UIPI) can block input
to an elevated application, so run winuse at the same integrity level as the
target. Secure desktop windows are not supported.
python3 -m venv .venv
.venv/bin/pip install maturin pytest
.venv/bin/maturin develop
.venv/bin/pytest
Python API
Move text through the system clipboard (pair with copy/paste shortcuts):
desktop . clipboard_write ( "Hello" ) # then press cmd+v to paste
clip = desktop . clipboard_read ()
print ( clip . kind , clip . text , clip . files ) # text | files | image | empty
desktop . clipboard_read ( save_image = "shot.png" ) # export an image clipboard
Control windows by id:
window = desktop . window ( app = "TextEdit" )
window . activate (); window . minimize (); window . restore (); window . maximize (); window . close ()
desktop . windows () # all real windows; each has .minimized
Wait for an async UI state instead of polling snapshots yourself:
send = window . find ( role = "button" , name = "Send" )
if send . wait_until ( "enabled" , timeout = 10 ):
send . click ()
window . find ( role = "static-text" , name = "Status" ). wait_until ( "text" , text = "Done" , timeout = 30 )
from winuse import Desktop
desktop = Desktop ()
if not desktop . is_authorized ():
raise RuntimeError ( "The platform accessibility backend is unavailable" )
window = desktop . window ( app = "TextEdit" )
editor = window . find ( role = "text-area" )
editor . fill ( "Hello from winuse" )
editor . press ( "enter" )
editor . fill ( "query" , submit = True ) # fill then press Enter (search boxes)
editor . type_text ( "inserted at the cursor" ) # type vs fill: inserts, does not replace
editor . press ( "s" , modifiers = [ "cmd" ]) # shortcut; cmd is Ctrl on Windows
editor . click ( double = True ) # double-click
editor . hover () # hover without clicking
window . find ( role = "list-item" , name = "File" ). drag_to ( name = "Folder" ) # drag onto a target
subtree = window . find ( role = "group" , name = "Sidebar" ). snapshot ( max_depth = 5 )
print ( subtree . text ())
Watch a subtree for semantic additions and removals:
messages = window . find ( role = "list" , name = "Messages" )
with messages . watch ( interval = 0.25 ) as changes :
while change := changes . next ( timeout = 30 ):
print ( change . kind , change . element )
Inspect every accessibility attribute exposed by an element:
inspection = window . find ( role = "static-text" , name = "Message" ). inspect ()
print ( inspection . text ())
Full native-attribute inspection works on both platforms: macOS reports the
element's AX attributes, Windows reports the UIA properties the element
supports (name, automation id, class name, value, patterns, …).
Capture the visible screen region occupied by an element:
button = window . find ( role = "button" , name = "Send" )
png = button . screenshot ( "send-button.png" , padding = 8 )
The method always returns the PNG bytes. Passing a path also writes the same
bytes to that file. The captured region reflects what is currently visible on
screen, including any windows covering the element.
Elements store a selector, not a long-lived native handle. Each action resolves
the element against a fresh accessibility snapshot.
cargo build --release -p winuse-cli
Install the agent skill (SKILL.md + references + host binary) to
~/.agents/skills/winuse , or package a distributable zip with all three
platform binaries:
make install # DEST=... to override the location
make package # → target/dist/winuse-<version>-skill.zip
Commands accept ordinary flags and always return a JSON envelope:
target/release/winuse doctor
target/release/winuse windows
target/release/winuse snapshot --window 123
target/release/winuse click --window 123 --role button --name Send --exact
target/release/winuse fill --window 123 --role text-area --text "Hello"
target/release/winuse press --window 123 --key enter
target/release/winuse scroll --window 123 --direction up --amount 500
The repository also contains a distributable agent skill at skills/winuse .
Its bundled executable is built from winuse-cli .
winuse was inspired by and references
agent-browser (Vercel Labs) —
its approach of driving a UI semantically for agents, with stable selectors
and JSON tool responses, shaped winuse's design for native desktop apps.
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
