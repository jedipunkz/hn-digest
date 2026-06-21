---
source: "https://github.com/m0rvayne/mcp-osascript"
hn_url: "https://news.ycombinator.com/item?id=48617397"
title: "MCP server that lets Claude click menus on your Mac and fix its own mistakes"
article_title: "GitHub - m0rvayne/mcp-osascript: Let Claude control your Mac — windows, menus, keyboard, clipboard. 12 typed tools, security-first MCP server. · GitHub"
author: "m0rvayne"
captured_at: "2026-06-21T10:17:48Z"
capture_tool: "hn-digest"
hn_id: 48617397
score: 1
comments: 0
posted_at: "2026-06-21T10:08:07Z"
tags:
  - hacker-news
  - translated
---

# MCP server that lets Claude click menus on your Mac and fix its own mistakes

- HN: [48617397](https://news.ycombinator.com/item?id=48617397)
- Source: [github.com](https://github.com/m0rvayne/mcp-osascript)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T10:08:07Z

## Translation

タイトル: クロードが Mac 上でメニューをクリックし、自身の間違いを修正できるようにする MCP サーバー
記事のタイトル: GitHub - m0rvayne/mcp-osascript: クロードに Windows、メニュー、キーボード、クリップボードなどの Mac を制御させます。 12種類のツール、セキュリティ第一のMCPサーバー。 · GitHub
説明: クロードに Windows、メニュー、キーボード、クリップボードなどの Mac を制御させます。 12種類のツール、セキュリティ第一のMCPサーバー。 - m0rvayne/mcp-osascript

記事本文:
GitHub - m0rvayne/mcp-osascript: クロードに Windows、メニュー、キーボード、クリップボードなどの Mac を制御させます。 12種類のツール、セキュリティ第一のMCPサーバー。 · GitHub
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
ムルヴァイン
/
mcp-osascript
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット アセット アセット サーバー サーバー .gitignore .gitignore ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
クロードに Mac を制御させましょう。ウィンドウの移動、メニューのクリック、テキストの入力、クリップボードの読み取り、ブラウザのタブの管理 — 入力検証とセキュリティ ガードレールを備えた 12 の入力ツール。
{
"mcpサーバー": {
"オサスクリプト" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " mcp-osascript " ]
}
}
}
これを Claude デスクトップ構成に追加し (「設定」→「開発者」→「構成の編集」)、クロードを再起動すれば準備は完了です。
{
"mcpサーバー": {
"オサスクリプト" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " mcp-osascript " ]
}
}
}
クロード・コード
クロード mcp 追加 osascript -- npx -y mcp-osascript
ソース（開発）から
git クローン https://github.com/m0rvayne/mcp-osascript.git
cd mcp-osascript && npm install
# 次に使用します: "command": "node", "args": ["/path/to/mcp-osascript/server/index.js"]
これらのプロンプトを試してください
12 個の型付きツール。それぞれに入力検証、エラー分類、権限を認識するエラー メッセージが含まれます。
クロードが存在しないメニュー項目をクリックしようとすると、サーバーはそのレベルで使用可能な項目のリストを自動的に返します。そのため、クロードは正しい名前で再試行できます。他の MCP サーバーはこれを行いません。
ユーザー: 「プレビューで [ファイル] → [PDF としてエクスポート] をクリックします」
クロード: app_menu を呼び出し、[「ファイル」、「PDF としてエクスポート」] をクリックします。
サーバー: 「メニュー項目「PDF としてエクスポート」が「ファイル」に見つかりません。
利用可能: [「クリップボードから新規作成」、「開く...」、「閉じる」、「保存」、
'複製'、'名前変更...'、'エクスポート...'、'PDF としてエクスポート...']"
クロード: app_menu を呼び出し、クリック ["ファイル"、"PDF としてエクスポート..."]
サーバー: 「クリック: [ファイル] > [PDF としてエクスポート...]

」
なぜ mcp-osascript なのか?
mcp-osascript
シュタイペテ (824★)
ピークモジョ (463★)
検証機能付きの入力ツール
12
2 (汎用)
1 (汎用)
URL スキーム許可リスト
http/https/mailto
いいえ
いいえ
環境の分離 (子プロセス)
パス+ホーム+LANGのみ
フルプロセス.env
フルプロセス.env
プロセス グループの強制終了 (孤立なし)
シグターム→シギキル
いいえ
いいえ
エラーのサニタイズ (パス、トークン)
はい
いいえ
いいえ
プロトタイプの汚染防止
オブジェクト.create(null)
いいえ
いいえ
自動修正メニュークリック
はい
いいえ
いいえ
結合テスト
41
0
0
標準入力パイプ (一時ファイルなし)
はい
一時ファイル
一時ファイル
権限
クリップボード、通知、URL、アプリなどの許可は必要ありません。すぐに使えます。
オートメーション — ブラウザーのタブ、最前面のアプリ。 macOS では、ブラウザごとに 1 回プロンプトが表示されます。
アクセシビリティ — キーボード、ウィンドウ、メニュー。 [システム設定] → [プライバシーとセキュリティ] → [アクセシビリティ] で 1 回許可します。
権限が欠落している場合、サーバーは何をすべきかを正確に通知します。
「アクセシビリティ許可が必要です。 「osascript」へのアクセスを許可する
[システム設定] > [プライバシーとセキュリティ] > [アクセシビリティ]
テスト
npmテスト
12 のツールすべてをカバーする 41 の統合テスト - 入力検証、セキュリティ境界 (URL スキームのブロック、プロトタイプの汚染、スクリプト サイズ制限)、タイムアウトの強制、権限エラーの処理。
run_osascript は任意のコードを実行します。これは仕様です。 MCP クライアント (クロード) が信頼境界です。
スクリプトは標準入力経由で /usr/bin/osascript にパイプされます。一時ファイルや TOCTOU 競合状態はありません。
スクリプトのサイズ: 最大 50 KB出力: 最大 50K 文字 (切り捨て)。
エラー メッセージがサニタイズされます。ファイル システムのパス、トークン、パスワードが削除されます。
子プロセスは最小限の環境 ( PATH 、 HOME 、 LANG のみ) を取得します。API キーやシークレットは漏洩しません。
URL スキーム許可リスト — file:// 、 smb:// 、 vnc:// 、 javascript: すべてブロックされます。
ハンドラーのディスパッチは Object.create(null) を使用します - プロトタイプ pol はありません

解決策。
タイムアウト時にプロセス グループを強制終了 — SIGTERM → 2 秒の猶予 → SIGKILL。孤立したプロセスはありません。
同時実行セマフォ — 最大 5 つの同時 osascript プロセス。
正常なシャットダウン - 10秒の強制終了セーフティネットを備えたserver.close()。
エラー分類 — macOS エラー コード (-1728、-1743、-25211) を解析して対処可能なメッセージにします。英語とロシア語のロケールをサポートします。
クロードにあなたの Mac (ウィンドウ、メニュー、キーボード、クリップボード) を制御させましょう。 12種類のツール、セキュリティ第一のMCPサーバー。
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

Let Claude control your Mac — windows, menus, keyboard, clipboard. 12 typed tools, security-first MCP server. - m0rvayne/mcp-osascript

GitHub - m0rvayne/mcp-osascript: Let Claude control your Mac — windows, menus, keyboard, clipboard. 12 typed tools, security-first MCP server. · GitHub
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
m0rvayne
/
mcp-osascript
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits assets assets server server .gitignore .gitignore LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Let Claude control your Mac. Move windows, click menus, type text, read clipboard, manage browser tabs — 12 typed tools with input validation and security guardrails.
{
"mcpServers" : {
"osascript" : {
"command" : " npx " ,
"args" : [ " -y " , " mcp-osascript " ]
}
}
}
Add this to your Claude Desktop config ( Settings → Developer → Edit Config ), restart Claude, and you're ready.
{
"mcpServers" : {
"osascript" : {
"command" : " npx " ,
"args" : [ " -y " , " mcp-osascript " ]
}
}
}
Claude Code
claude mcp add osascript -- npx -y mcp-osascript
From source (development)
git clone https://github.com/m0rvayne/mcp-osascript.git
cd mcp-osascript && npm install
# then use: "command": "node", "args": ["/path/to/mcp-osascript/server/index.js"]
Try These Prompts
12 typed tools, each with input validation, error classification, and permission-aware error messages.
When Claude tries to click a menu item that doesn't exist, the server automatically returns the list of available items at that level — so Claude can retry with the correct name. No other MCP server does this.
User: "Click File → Export as PDF in Preview"
Claude: calls app_menu click ["File", "Export as PDF"]
Server: "Menu item 'Export as PDF' not found in 'File'.
Available: ['New from Clipboard', 'Open...', 'Close', 'Save',
'Duplicate', 'Rename...', 'Export...', 'Export as PDF...']"
Claude: calls app_menu click ["File", "Export as PDF..."]
Server: "Clicked: File > Export as PDF..."
Why mcp-osascript?
mcp-osascript
steipete (824★)
peakmojo (463★)
Typed tools with validation
12
2 (generic)
1 (generic)
URL scheme allowlist
http/https/mailto
No
No
Env isolation (child process)
PATH+HOME+LANG only
Full process.env
Full process.env
Process group kill (no orphans)
SIGTERM→SIGKILL
No
No
Error sanitization (paths, tokens)
Yes
No
No
Prototype pollution protection
Object.create(null)
No
No
Self-correcting menu click
Yes
No
No
Integration tests
41
0
0
Stdin piping (no temp files)
Yes
Temp files
Temp files
Permissions
No permission needed — clipboard, notifications, URLs, apps. Works immediately.
Automation — browser tabs, frontmost app. macOS prompts once per browser.
Accessibility — keyboard, windows, menus. Grant once in System Settings → Privacy & Security → Accessibility .
When a permission is missing, the server tells you exactly what to do:
"Accessibility permission required. Grant access to 'osascript'
in System Settings > Privacy & Security > Accessibility."
Testing
npm test
41 integration tests covering all 12 tools — input validation, security boundaries (URL scheme blocking, prototype pollution, script size limits), timeout enforcement, and permission error handling.
run_osascript executes arbitrary code — this is by design. The MCP client (Claude) is the trust boundary.
Scripts piped via stdin to /usr/bin/osascript — no temp files, no TOCTOU race conditions.
Script size: 50 KB max. Output: 50K chars max (truncated).
Error messages sanitized — filesystem paths, tokens, and passwords are stripped.
Child processes get minimal env: PATH , HOME , LANG only — no API keys or secrets leak.
URL scheme allowlist — file:// , smb:// , vnc:// , javascript: all blocked.
Handler dispatch uses Object.create(null) — no prototype pollution.
Process group kill on timeout — SIGTERM → 2s grace → SIGKILL. No orphaned processes.
Concurrency semaphore — max 5 simultaneous osascript processes.
Graceful shutdown — server.close() with 10s force-exit safety net.
Error classification — parses macOS error codes (-1728, -1743, -25211) into actionable messages. Supports English and Russian locales.
Let Claude control your Mac — windows, menus, keyboard, clipboard. 12 typed tools, security-first MCP server.
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
