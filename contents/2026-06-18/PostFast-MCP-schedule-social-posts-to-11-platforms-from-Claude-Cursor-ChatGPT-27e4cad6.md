---
source: "https://github.com/peturgeorgievv-factory/postfast-mcp"
hn_url: "https://news.ycombinator.com/item?id=48584837"
title: "PostFast MCP – schedule social posts to 11 platforms from Claude/Cursor/ChatGPT"
article_title: "GitHub - peturgeorgievv-factory/postfast-mcp: MCP server for the PostFast API — schedule and manage social media posts (X, Instagram, Facebook, TikTok, LinkedIn, YouTube, and more) via AI tools like Claude, Cursor, and VS Code. · GitHub"
author: "peturgeorgievv"
captured_at: "2026-06-18T16:16:17Z"
capture_tool: "hn-digest"
hn_id: 48584837
score: 2
comments: 0
posted_at: "2026-06-18T13:16:48Z"
tags:
  - hacker-news
  - translated
---

# PostFast MCP – schedule social posts to 11 platforms from Claude/Cursor/ChatGPT

- HN: [48584837](https://news.ycombinator.com/item?id=48584837)
- Source: [github.com](https://github.com/peturgeorgievv-factory/postfast-mcp)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T13:16:48Z

## Translation

タイトル: PostFast MCP – Claude/Cursor/ChatGPT から 11 のプラットフォームへのソーシャル投稿をスケジュールする
記事のタイトル: GitHub - peturgeorgievv-factory/postfast-mcp: PostFast API 用の MCP サーバー — Claude、Cursor、VS Code などの AI ツールを介してソーシャル メディア投稿 (X、Instagram、Facebook、TikTok、LinkedIn、YouTube など) をスケジュールおよび管理します。 · GitHub
説明: PostFast API 用の MCP サーバー — Claude、Cursor、VS Code などの AI ツールを介してソーシャル メディア投稿 (X、Instagram、Facebook、TikTok、LinkedIn、YouTube など) をスケジュールおよび管理します。 - peturgeorgievv-factory/postfast-mcp

記事本文:
GitHub - peturgeorgievv-factory/postfast-mcp: PostFast API 用の MCP サーバー — Claude、Cursor、VS Code などの AI ツールを介してソーシャル メディア投稿 (X、Instagram、Facebook、TikTok、LinkedIn、YouTube など) をスケジュールおよび管理します。 · GitHub
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
p

エトルギオルギエフ工場
/
postfast-mcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
peturgeorgievv-factory/postfast-mcp
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
37 コミット 37 コミット .claude-plugin .claude-plugin .github .github skill/ social-media-post skill/ social-media-post src src .gitignore .gitignore .mcp.json .mcp.json .mcpbignore .mcpbignore ライセンス ライセンス README.md README.md グラマ.json グラマ.json マニフェスト.json マニフェスト.json パッケージ-ロック.json パッケージ-ロック.json パッケージ.json パッケージ.json サーバー.json サーバー.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
PostFast API 用の MCP サーバー — Claude、Cursor、VS Code などの AI ツールを介してソーシャル メディア投稿をスケジュールおよび管理します。
ローカルでは何も実行したくないですか?ホストされたエンドポイントに接続し、OAuth で認証します。npx も管理する API キーもありません。
https://mcp.postfa.st/mcp
OAuth をサポートするクライアント (ChatGPT、Claude など) にリモート/ストリーミング可能な HTTP MCP サーバーとして追加します。初めて使用するときに、PostFast にサインインしてアクセスを承認するように求められます。
代わりに自分でサーバーを実行したいですか?以下のクイック スタートで npx + API キーのセットアップを使用します。
PostFast にログインし、サイドバーの API に移動して、キーを生成します。
Claude Desktop 拡張機能ディレクトリから拡張機能をダウンロードするか、手動でインストールします。
claude_desktop_config.json に追加します。
{
"mcpサーバー": {
"ポストファスト" : {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " postfast-mcp " ]、
"環境" : {
"POSTFAST_API_KEY" : " ここの API キー "
}
}
}
}
クロードデスクトップを再起動します。
プラグイン経由 (マーケットプレイスの承認待ち):
/plugin install postfast@claude-plugins-official
インストール後、API キーを設定します。次のいずれかを選択します。
# オプション A: 彼女に追加する

ll プロファイル (~/.zshrc または ~/.bashrc)
import POSTFAST_API_KEY= " your-api-key-here "
# オプション B: ~/.claude/settings.local.json に追加
# { "env": { "POSTFAST_API_KEY": "ここにある API キー" } }
その後、クロード コードを再起動します。
プロジェクトの .mcp.json または ~/.claude/.mcp.json (グローバル) に追加します。
{
"mcpサーバー": {
"ポストファスト" : {
"タイプ" : " stdio " ,
"コマンド" : " npx " ,
"args" : [ " -y " 、 " postfast-mcp " ]、
"環境" : {
"POSTFAST_API_KEY" : " ここの API キー "
}
}
}
}
カーソル / VS Code / Windsurf / その他の MCP クライアント
MCP 構成 ( .mcp.json 、 mcp.json 、またはツールの設定 UI) に追加します。
{
"mcpサーバー": {
"ポストファスト" : {
"タイプ" : " stdio " ,
"コマンド" : " npx " ,
"args" : [ " -y " 、 " postfast-mcp " ]、
"環境" : {
"POSTFAST_API_KEY" : " ここの API キー "
}
}
}
}
3. 使ってみよう
AI アシスタントに次のような質問をしてください。
「接続されているソーシャルアカウントをリストする」
「明日の午前9時にInstagramへの投稿をスケジュールしてください」
「今週予定されている投稿をすべて表示」
「この画像をアップロードして、それを使用して LinkedIn 投稿を作成します」
「このビデオを使って Facebook リールを作成してください」
「今月の Instagram 投稿の分析を見せて」
ツール
説明
リストアカウント
接続されているソーシャル メディア アカウントを一覧表示します (それぞれに connectionStatus — CONNECTED / DISABLED — およびdisabledReason が付いています)
list_posts
フィルター付きの投稿のリスト (特定の ID、プラットフォーム、ステータス、日付範囲)
作成_投稿
投稿の作成とスケジュール (バッチ、最大 15)
投稿の削除
IDで投稿を削除する
アップロードメディア
ローカル ファイルをアップロードし、メディア キーを取得します (完全なフローを処理します)。
get_upload_urls
署名付き URL を取得してメディア ファイルをアップロードする
list_pinterest_boards
アカウント用の Pinterest ボードを取得する
list_youtube_playlists
アカウントの YouTube プレイリストを取得する
list_gbp_locations
アカウントの Google ビジネス プロフィールの場所を取得する
検索場所
投稿にタグを付ける場所を見つけます (ID は facebookPlaceId の両方で機能します)

および instagramLocationId )
生成_接続_リンク
クライアントがアカウントに接続するためのリンクを生成する
get_post_analytics
パフォーマンス指標を含む公開済みの投稿を取得 — Instagram、Facebook、TikTok、Threads、YouTube、LinkedIn (企業ページ)、Pinterest (ビジネス アカウント)
フォロワー履歴の取得
アカウントの毎日のフォロワー数履歴 (現在の数 + 日付範囲にわたる差分) — Facebook ページ、Instagram、YouTube、Pinterest、Threads、Bluesky、Telegram、LinkedIn (企業ページ)
サポートされているプラットフォーム
Facebook、Instagram、X (Twitter)、TikTok、LinkedIn、YouTube、BlueSky、スレッド、Pinterest、Telegram、Google ビジネス プロフィール
投稿を作成するときに、controls パラメーターを介してプラットフォーム固有の設定を渡すことができます。
Upload_media ツールは、1 回の呼び出しで完全なフローを処理します。
ファイル拡張子からコンテンツタイプを検出します
PostFast から署名付きアップロード URL を取得します
create_posts で使用できるキーとタイプを返します。
サポートされている形式: JPEG、PNG、GIF、WebP、MP4、WebM、MOV
アップロード プロセスをさらに制御する必要がある場合は、get_upload_urls を直接使用することもできます。
変数
必須
説明
POSTFAST_API_KEY
はい
ワークスペース API キー
POSTFAST_API_URL
いいえ
API ベース URL (デフォルト: https://api.postfa.st )
テスト
MCP Inspector ですべてが動作することを確認します。
POSTFAST_API_KEY=your-key npx @modelcontextprotocol/inspector npx postfast-mcp
APIドキュメント
REST API の完全なドキュメント: postfa.st/docs
npmインストール
npm ビルドを実行する
ノード dist/index.js
バッジ
PostFast API 用の MCP サーバー — Claude、Cursor、VS Code などの AI ツールを介してソーシャル メディア投稿 (X、Instagram、Facebook、TikTok、LinkedIn、YouTube など) をスケジュールおよび管理します。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
1
v0.1.9
最新の
2026 年 4 月 6 日
パッケージ
0
がありました

ロード中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

MCP server for the PostFast API — schedule and manage social media posts (X, Instagram, Facebook, TikTok, LinkedIn, YouTube, and more) via AI tools like Claude, Cursor, and VS Code. - peturgeorgievv-factory/postfast-mcp

GitHub - peturgeorgievv-factory/postfast-mcp: MCP server for the PostFast API — schedule and manage social media posts (X, Instagram, Facebook, TikTok, LinkedIn, YouTube, and more) via AI tools like Claude, Cursor, and VS Code. · GitHub
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
peturgeorgievv-factory
/
postfast-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
peturgeorgievv-factory/postfast-mcp
main Branches Tags Go to file Code Open more actions menu Folders and files
37 Commits 37 Commits .claude-plugin .claude-plugin .github .github skills/ social-media-post skills/ social-media-post src src .gitignore .gitignore .mcp.json .mcp.json .mcpbignore .mcpbignore LICENSE LICENSE README.md README.md glama.json glama.json manifest.json manifest.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json View all files Repository files navigation
MCP server for the PostFast API — schedule and manage social media posts via AI tools like Claude, Cursor, VS Code, and more.
Prefer not to run anything locally? Connect to the hosted endpoint and authenticate with OAuth — no npx, no API key to manage:
https://mcp.postfa.st/mcp
Add it as a remote/streamable-HTTP MCP server in any client that supports OAuth (e.g. ChatGPT, Claude). You'll be prompted to sign in to PostFast and authorize access on first use.
Want to run the server yourself instead? Use the npx + API-key setup in Quick Start below.
Log in to PostFast , go to API in the sidebar, and generate a key.
Download the extension from the Claude Desktop extension directory or install manually:
Add to claude_desktop_config.json :
{
"mcpServers" : {
"postfast" : {
"command" : " npx " ,
"args" : [ " -y " , " postfast-mcp " ],
"env" : {
"POSTFAST_API_KEY" : " your-api-key-here "
}
}
}
}
Restart Claude Desktop.
Via plugin (pending marketplace approval):
/plugin install postfast@claude-plugins-official
After installing, set your API key — pick one of these:
# Option A: Add to your shell profile (~/.zshrc or ~/.bashrc)
export POSTFAST_API_KEY= " your-api-key-here "
# Option B: Add to ~/.claude/settings.local.json
# { "env": { "POSTFAST_API_KEY": "your-api-key-here" } }
Then restart Claude Code.
Add to your project's .mcp.json or ~/.claude/.mcp.json (global):
{
"mcpServers" : {
"postfast" : {
"type" : " stdio " ,
"command" : " npx " ,
"args" : [ " -y " , " postfast-mcp " ],
"env" : {
"POSTFAST_API_KEY" : " your-api-key-here "
}
}
}
}
Cursor / VS Code / Windsurf / Other MCP clients
Add to your MCP config ( .mcp.json , mcp.json , or the tool's settings UI):
{
"mcpServers" : {
"postfast" : {
"type" : " stdio " ,
"command" : " npx " ,
"args" : [ " -y " , " postfast-mcp " ],
"env" : {
"POSTFAST_API_KEY" : " your-api-key-here "
}
}
}
}
3. Use it
Ask your AI assistant things like:
"List my connected social accounts"
"Schedule a post to Instagram for tomorrow at 9am"
"Show me all scheduled posts for this week"
"Upload this image and create a LinkedIn post with it"
"Create a Facebook reel with this video"
"Show me analytics for my Instagram posts this month"
Tool
Description
list_accounts
List connected social media accounts (each with connectionStatus — CONNECTED / DISABLED — and disabledReason )
list_posts
List posts with filters (specific IDs, platform, status, date range)
create_posts
Create and schedule posts (batch, up to 15)
delete_post
Delete a post by ID
upload_media
Upload a local file and get a media key (handles the full flow)
get_upload_urls
Get signed URLs to upload media files
list_pinterest_boards
Get Pinterest boards for an account
list_youtube_playlists
Get YouTube playlists for an account
list_gbp_locations
Get Google Business Profile locations for an account
search_places
Find a place to tag posts (the id works for both facebookPlaceId and instagramLocationId )
generate_connect_link
Generate a link for clients to connect accounts
get_post_analytics
Fetch published posts with performance metrics — Instagram, Facebook, TikTok, Threads, YouTube, LinkedIn (company pages), Pinterest (Business accounts)
get_follower_history
Daily follower-count history for an account (current count + delta over a date range) — Facebook Pages, Instagram, YouTube, Pinterest, Threads, Bluesky, Telegram, LinkedIn (company pages)
Supported Platforms
Facebook, Instagram, X (Twitter), TikTok, LinkedIn, YouTube, BlueSky, Threads, Pinterest, Telegram, Google Business Profile
When creating posts, you can pass platform-specific settings via the controls parameter:
The upload_media tool handles the full flow in a single call:
Detects content type from file extension
Gets a signed upload URL from PostFast
Returns a key and type ready to use in create_posts
Supported formats: JPEG, PNG, GIF, WebP, MP4, WebM, MOV
You can also use get_upload_urls directly if you need more control over the upload process.
Variable
Required
Description
POSTFAST_API_KEY
Yes
Your workspace API key
POSTFAST_API_URL
No
API base URL (default: https://api.postfa.st )
Testing
Verify everything works with the MCP Inspector:
POSTFAST_API_KEY=your-key npx @modelcontextprotocol/inspector npx postfast-mcp
API Docs
Full REST API documentation: postfa.st/docs
npm install
npm run build
node dist/index.js
Badges
MCP server for the PostFast API — schedule and manage social media posts (X, Instagram, Facebook, TikTok, LinkedIn, YouTube, and more) via AI tools like Claude, Cursor, and VS Code.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
v0.1.9
Latest
Apr 6, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
