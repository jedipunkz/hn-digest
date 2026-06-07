---
source: "https://github.com/hansstam86/wibeos"
hn_url: "https://news.ycombinator.com/item?id=48434939"
title: "Show HN: WibeOS, an OS where every app is hallucinated by an LLM"
article_title: "GitHub - hansstam86/wibeos · GitHub"
author: "hans863"
captured_at: "2026-06-07T15:35:08Z"
capture_tool: "hn-digest"
hn_id: 48434939
score: 2
comments: 1
posted_at: "2026-06-07T13:58:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: WibeOS, an OS where every app is hallucinated by an LLM

- HN: [48434939](https://news.ycombinator.com/item?id=48434939)
- Source: [github.com](https://github.com/hansstam86/wibeos)
- Score: 2
- Comments: 1
- Posted: 2026-06-07T13:58:31Z

## Translation

タイトル: Show HN: WibeOS、すべてのアプリが LLM によって幻覚を受ける OS
記事のタイトル: GitHub - hansstam86/wibeos · GitHub
説明: GitHub でアカウントを作成して、hansstam86/wibeos の開発に貢献します。

記事本文:
GitHub - hansstam86/wibeos · GitHub
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
ハンスタム86
/
ウィベオス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

s
6 コミット 6 コミット ソース/ WibeOS ソース/ WibeOS アセット アセット スクリプト スクリプト .gitignore .gitignore GUMROAD.md GUMROAD.md Info.plist Info.plist KICKSTARTER.md KICKSTARTER.md ライセンス ライセンス LINKEDIN-SERIES.md LINKEDIN-SERIES.md Makefile Makefile Package.swift Package.swift README.md README.md wibeOS-shoot-runsheet.docx wibeOS-shoot-runsheet.docx すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Steve Sanderson のアイデアにインスピレーションを得た、macOS 用の完全に幻覚を備えたオペレーティング システム
BUILD 2026 からの vibeOS。
デスクトップ シェルは実際のローカルなものです (メニュー バー、ドック、ドラッグ可能なウィンドウ、メニュー —
すべて瞬時に行われ、API 呼び出しは不要です）。アプリはそうではありません。各アプリのウィンドウが幻覚を起こしています。
Claude による、サンドボックス内で実行される自己完結型のインタラクティブな HTML ミニアプリとして
iフレーム。クロードは動作する JavaScript をオンザフライで作成するため、実際には電卓が作成されます。
ローカルで計算します。新鮮な想像力が必要なアクションのみ (偽の世界をナビゲートする)
Jungle ブラウザなど) が再生成をトリガーします。開いたアプリはキャッシュされるため、再度開くには
インスタント。
Swift ツールチェーン (Xcode またはコマンド ライン ツール: xcode-select --install )
Anthropic API キー ( https://console.anthropic.com )
make run # ビルドして起動する
make app # build WibeOS.app、次に: WibeOS.app を開きます
最初の起動時に、API キー (アプリの設定に保存されています。
または、代わりに環境で ANTHROPIC_API_KEY を設定します)。
wibe.fs — ペルソナごとの実際の永続ファイル システム: メモを保存します
メモは、Finder で見つけて TextEdit で開きます。再起動しても存続します。
スポットライト — Cmd+K (またはメニュー バーの 🔍): 任意のアプリを起動し、任意のアプリを開きます
ファイルに保存したり、その場で何かを想像したりできます。
通知 — アプリは wibe.notify() を呼び出し、数分ごとに
ペルソナの世界は自動的に彼らに ping を送信します (電子メール、リマインダー、キャラクター内)
イベント）。トーストをクリックすると、関連するアプリが開きます。
wibe.ai — アプリは Claude を呼び出すことができます

ランタイムではないため、チャットボットは実際にチャットします
そして占い師は実際に占うのです。
The dock has an App Store stocked entirely with apps invented for whoever is
logged in — Get installs them to the dock permanently (per persona;
削除するにはドックアイコンを右クリックします)。最初のログイン時に、各ペルソナも
full OS theme hallucinated for them — wallpaper, accent, menu bar, dock,
window chrome, fonts — and apps are told the accent/dark-mode so they match.
ブートするとログイン画面が表示されます。誰としてログインするかによってすべてが変わります。
persona is injected into every hallucination, so Grandma's Mail is chain
letters in giant fonts while the hacker kid's whole machine is neon-on-black
1998. 各ペルソナは独自のアプリ キャッシュを取得します。 + をクリックして新しいペルソナを作成します
(「コンピューターを発見する中世の魔法使い」);アバターを右クリックして削除します
それ。メニューからログアウトします。
Click dock icons to "launch" apps — you watch the HTML stream by as Claude
アプリに幻覚を与え、その後、それが現実になります。 ✨ をクリックしてアプリを想像してください。
Option-click a dock icon (or View → Re-imagine) to force a fresh hallucination
キャッシュされたバージョンの代わりに。
Menu bar, window dragging/minimize/zoom, Cmd+W close — all local and instant.
Cmd+R で再起動します。 Cmd+Ctrl+F for full screen — recommended for the illusion.
Right-click → Inspect Element works for debugging the generated HTML.
Camera/photo-booth apps use your real webcam (macOS will ask once).のために
permission prompt to be attributed correctly, run as a bundle: make app
then open WibeOS.app — with make run the prompt belongs to your terminal.
Persistent cache : each app is hallucinated once, then stored in
~/Library/Application Support/wibeOS/cache — reopening is instant, even
再起動後も。 Option-click a dock icon or View → Re-imagine to regenerate.
プリフェッチ:

起動時に、ドック アプリがバックグラウンドで事前に幻覚表示されます。
(一度に 2 つ) なので、ほとんどのアプリはクリックする前に準備が完了しています。で無効にする
WIBEOS_PREFETCH=0 (最初のブート プリフェッチにはアプリのおよそ 8 世代分のコストがかかります)。
パッチ更新: データワイブ アクション (ブラウザ ナビゲーション、電子メールの開封など)
変更された領域のみを <wibe-patch select="…"> ブロックとして返します。
アプリ全体 - トークンの一部。再マウントせずにライブで適用されます。
プロンプト キャッシュ : システム プロンプトは呼び出し全体にわたってサーバー側でキャッシュされます。
モデル : デフォルトは claude-haiku-4-5-20251001 (高速) です。より高度なアプリの場合:
WIBEOS_MODEL=claude-sonnet-4-6 を実行します。
システム プロンプト: Sources/WibeOS/AppDelegate.swift の systemPrompt を次のように編集します。
OS の性格、デフォルトのドック、外観などを変更します。
アプリを開く ─▶ シェルはクロードに 1 つの自己完結型 HTML アプリ (ストリーミング) を要求します
─▶ 挿入されたブリッジ スクリプトが追加され、サンドボックス化された <iframe> にマウントされました
─▶ アプリはローカル (独自の JS) で実行されます — ほとんどのクリックには費用がかかりません
─▶ クロードへの往復のデータワイブ / データワイブエンターのマークを付けた制御
新鮮な幻覚コンテンツの場合 (ウィンドウ全体の再レンダリング)
各ウィンドウには、独自の小さな会話履歴 (作成ペア + 最新 6 件) が保存されます。
メッセージ））、更新の一貫性が保たれ、低コストで更新されます。 Swift レイヤーはステートレスです
Anthropic API へのストリーミング プロキシ。
トークンがかかるのはアプリ生成とデータワイブアクションのみです。それ以外はすべてローカルです。
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

Contribute to hansstam86/wibeos development by creating an account on GitHub.

GitHub - hansstam86/wibeos · GitHub
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
hansstam86
/
wibeos
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits Sources/ WibeOS Sources/ WibeOS assets assets scripts scripts .gitignore .gitignore GUMROAD.md GUMROAD.md Info.plist Info.plist KICKSTARTER.md KICKSTARTER.md LICENSE LICENSE LINKEDIN-SERIES.md LINKEDIN-SERIES.md Makefile Makefile Package.swift Package.swift README.md README.md wibeOS-shoot-runsheet.docx wibeOS-shoot-runsheet.docx View all files Repository files navigation
A fully hallucinated operating system for macOS, inspired by Steve Sanderson's
vibeOS from BUILD 2026.
The desktop shell is real and local (menu bar, dock, draggable windows, menus —
all instant, zero API calls). The apps are not: each app window is hallucinated
by Claude as a self-contained interactive HTML mini-app running in a sandboxed
iframe. Claude writes working JavaScript on the fly, so a calculator actually
computes locally; only actions that need fresh imagination (navigating the fake
the Jungle browser, etc.) trigger a regeneration. Opened apps are cached, so reopening is
instant.
Swift toolchain (Xcode or Command Line Tools: xcode-select --install )
An Anthropic API key ( https://console.anthropic.com )
make run # build & launch
make app # build WibeOS.app, then: open WibeOS.app
On first launch it asks for your API key (stored in the app's preferences;
or set ANTHROPIC_API_KEY in the environment instead).
wibe.fs — a real persistent file system, per persona: save a note in
Notes, find it in Finder, open it in TextEdit. Survives reboots.
Spotlight — Cmd+K (or 🔍 in the menu bar): launch any app, open any
file, or imagine anything on the spot.
Notifications — apps call wibe.notify(), and every few minutes the
persona's world pings them on its own (emails, reminders, in-character
events). Click a toast to open the relevant app.
wibe.ai — apps can call Claude at runtime, so chatbots actually chat
and fortune tellers actually divine.
The dock has an App Store stocked entirely with apps invented for whoever is
logged in — Get installs them to the dock permanently (per persona;
right-click a dock icon to remove). On first login each persona also gets a
full OS theme hallucinated for them — wallpaper, accent, menu bar, dock,
window chrome, fonts — and apps are told the accent/dark-mode so they match.
Boot lands on a login screen. Who you log in as changes everything : the
persona is injected into every hallucination, so Grandma's Mail is chain
letters in giant fonts while the hacker kid's whole machine is neon-on-black
1998. Each persona gets its own app cache. Click + to invent new personas
("a medieval wizard discovering computers"); right-click an avatar to remove
it. Log out via the menu.
Click dock icons to "launch" apps — you watch the HTML stream by as Claude
hallucinates the app, then it becomes real(ish). Click ✨ to imagine any app.
Option-click a dock icon (or View → Re-imagine) to force a fresh hallucination
instead of the cached version.
Menu bar, window dragging/minimize/zoom, Cmd+W close — all local and instant.
Cmd+R reboots. Cmd+Ctrl+F for full screen — recommended for the illusion.
Right-click → Inspect Element works for debugging the generated HTML.
Camera/photo-booth apps use your real webcam (macOS will ask once). For the
permission prompt to be attributed correctly, run as a bundle: make app
then open WibeOS.app — with make run the prompt belongs to your terminal.
Persistent cache : each app is hallucinated once, then stored in
~/Library/Application Support/wibeOS/cache — reopening is instant, even
across reboots. Option-click a dock icon or View → Re-imagine to regenerate.
Prefetch : at boot, the dock apps are pre-hallucinated in the background
(2 at a time), so most apps are ready before you click. Disable with
WIBEOS_PREFETCH=0 (first boot prefetch costs roughly 8 app generations).
Patch updates : data-wibe actions (browser navigation, opening an email…)
return only the changed region as <wibe-patch select="…"> blocks instead of
the whole app — a fraction of the tokens, applied live without remounting.
Prompt caching : the system prompt is cached server-side across calls.
Model : defaults to claude-haiku-4-5-20251001 (fast). For fancier apps:
WIBEOS_MODEL=claude-sonnet-4-6 make run .
System prompt: edit systemPrompt in Sources/WibeOS/AppDelegate.swift to
change the OS's personality, default dock, look, etc.
open app ─▶ shell asks Claude for ONE self-contained HTML app (streamed)
─▶ injected bridge script added, mounted in sandboxed <iframe>
─▶ app runs locally (its own JS) — most clicks cost nothing
─▶ controls marked data-wibe / data-wibe-enter round-trip to Claude
for fresh hallucinated content (whole-window re-render)
Each window keeps its own small conversation history (creation pair + last 6
messages), so updates stay consistent and cheap. The Swift layer is a stateless
streaming proxy to the Anthropic API.
Only app generation and data-wibe actions cost tokens; everything else is local.
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
