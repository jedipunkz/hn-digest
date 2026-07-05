---
source: "https://github.com/cloudflare/agentic-inbox"
hn_url: "https://news.ycombinator.com/item?id=48794625"
title: "Self-hosted email client with an AI agent, running on Cloudflare Workers"
article_title: "GitHub - cloudflare/agentic-inbox: A self-hosted email client with an AI agent, running entirely on Cloudflare Workers · GitHub"
author: "Brajeshwar"
captured_at: "2026-07-05T15:06:35Z"
capture_tool: "hn-digest"
hn_id: 48794625
score: 3
comments: 0
posted_at: "2026-07-05T14:38:45Z"
tags:
  - hacker-news
  - translated
---

# Self-hosted email client with an AI agent, running on Cloudflare Workers

- HN: [48794625](https://news.ycombinator.com/item?id=48794625)
- Source: [github.com](https://github.com/cloudflare/agentic-inbox)
- Score: 3
- Comments: 0
- Posted: 2026-07-05T14:38:45Z

## Translation

タイトル: Cloudflare Workers 上で実行される、AI エージェントを備えた自己ホスト型電子メール クライアント
記事のタイトル: GitHub - Cloudflare/agentic-inbox: Cloudflare Workers · GitHub 上で完全に実行される、AI エージェントを備えたセルフホスト型電子メール クライアント
説明: AI エージェントを備えた自己ホスト型電子メール クライアント。完全に Cloudflare Workers 上で実行されます - Cloudflare/agentic-inbox

記事本文:
GitHub - Cloudflare/agentic-inbox: AI エージェントを備えたセルフホスト型電子メール クライアント。完全に Cloudflare Workers · GitHub 上で実行されます。
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
クラウドフレア
/
エージェントの受信箱
公共
通知

ns
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット アプリ app public publicsharedsharedworkersworkers .dev.vars.example .dev.vars.example .gitignore .gitignore ライセンス ライセンス README.md README.md デモ_app.png デモ_app.png パッケージロック.json パッケージロック.json パッケージ.json パッケージ.json 反応ルーター.config.ts React-router.config.ts tsconfig.cloudflare.json tsconfig.cloudflare.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts wrangler.jsonc wrangler.jsonc すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AIエージェントを備えた自己ホスト型電子メールクライアントで、完全にCloudflare Workers上で実行されます。
Agentic Inbox を使用すると、最新の Web インターフェイスを介して電子メールを送信、受信、管理できます。すべて独自の Cloudflare アカウントによって実行されます。受信メールは Cloudflare Email Routing 経由で到着し、各メールボックスは SQLite データベースを備えた独自の永続オブジェクトに分離され、添付ファイルは R2 に保存されます。
AIを活用した電子メールエージェントは、受信トレイを読み取り、会話を検索し、返信の下書きを行うことができます。これはCloudflare Agents SDKとWorkers AIで構築されています。
Cloudflare Email Service の詳細と、それを Agent SDK、MCP、および Wrangler CLI で使用する方法については、ブログ投稿「Email for Agents」を参照してください。
重要 : 「Cloudflare にデプロイ」ボタンをクリックすることは、セットアップの一部にすぎません。デプロイ後の手順にも従う必要があります。スクリーンショット付きの完全なステップバイステップ ガイドについては、次のコメントを参照してください。
#4 (コメント)
Cloudflareにデプロイします。デプロイ フローでは、R2、Durable Objects、Workers AI が自動的にプロビジョニングされます。 DOMAINS の入力を求められます。これは、電子メールを受信するドメイン (yourdomain.com) です ( ema

il@yourdomain.com )。
Cloudflare アクセスを構成する -- [設定] > [ドメインとルート] でワーカーのワンクリック Cloudflare アクセスを有効にします。モーダルには、POLICY_AUD と TEAM_DOMAIN の値が表示されます。 TEAM_DOMAIN には、Access チームの URL または完全な .../cdn-cgi/access/certs URL を指定できます。これらをワーカーのシークレットとして設定する必要があります。
電子メール ルーティングを設定する -- Cloudflare ダッシュボードで、ドメイン > 電子メール ルーティングに移動し、このワーカーに転送するキャッチオール ルールを作成します。
電子メール サービスを有効にする -- ワーカーがアウトバウンド電子メールを送信するには、send_email バインディングが必要です。電子メールサービスのドキュメントを参照してください
メールボックスの作成 -- デプロイされたアプリにアクセスし、ドメイン上の任意のアドレス (例: hello@example.com ) のメールボックスを作成します。
「無効なアクセス トークン」または「期限切れのアクセス トークン」が表示された場合は、通常、POLICY_AUD または TEAM_DOMAIN シークレットが間違っていることを意味します。
解決策: ワーカーのアクセスをオフにして再度オンにして、アクセス モーダルを再度取得し、ワーカーのシークレットをそこに表示されている最新の POLICY_AUD および TEAM_DOMAIN 値にリセットします。
「 Cloudflare Access は運用環境で設定する必要があります 」と表示されている場合、このアプリケーションは意図的に Cloudflare Access を強制しているため、受信トレイはインターネット上の誰にも公開されません。
解決策: ワンクリック Cloudflare Access for Workers を使用してアクセスを有効にし、モーダル値から POLICY_AUD および TEAM_DOMAIN ワーカー シークレットを設定します。
完全な電子メールクライアント — リッチテキストコンポーザー、返信/転送スレッド、フォルダー編成、検索、添付ファイルを備えたCloudflare電子メールルーティング経由で電子メールを送受信します。
メールボックスごとの分離 — 各メールボックスは、SQLite ストレージと添付ファイル用の R2 を備えた独自の永続オブジェクトで実行されます。
組み込み AI エージェント — 読み取り、検索、作成、送信のための 9 つの電子メール ツールを備えたサイド パネル
新しい電子メールの自動下書き - エージェントが受信電子メールを自動的に読み取って生成します。

返信の下書き、送信前に常に明示的な確認が必要
構成可能かつ永続的 — メールボックスごとのカスタム システム プロンプト、永続的なチャット履歴、ストリーミング マークダウン応答、およびツール呼び出しの可視性
フロントエンド: React 19、React Router v7、Tailwind CSS、Zustand、TipTap、@cloudflare/kumo
バックエンド: Hono、Cloudflare Workers、Durable Objects (SQLite)、R2、電子メール ルーティング
AI エージェント: Cloudflare エージェント SDK ( AIChatAgent )、AI SDK v6、Workers AI ( @cf/moonshotai/kimi-k2.5 )、react-markdown + mark-gfm
認証: Cloudflare Access JWT 検証 (ローカル開発以外で必要)
npmインストール
npm 実行開発
構成
wrangler.jsonc でドメインを設定します。
Agentic-inbox という名前の R2 バケットを作成します: Wrangler r2 Bucket create Agentic-inbox
ドメインを持つCloudflareアカウント
電子メールのルーティングが受信可能になっています
電子メールサービスの送信が有効になっています
Workers AI が有効になっている (エージェント用)
デプロイ/共有環境用に構成されたCloudflare Access (本番環境で必要)
共有Cloudflareアクセスポリシーに合格したユーザーは、設計上、このアプリ内のすべてのメールボックスにアクセスできます。これには、/mcp の MCP サーバーが含まれます。MCP 経由で接続された外部 AI ツール (クロード コード、カーソルなど) は、mailboxId パラメーターを渡すことで、任意のメールボックスで動作できます。メールボックスごとの承認はありません。 Cloudflare Access ポリシーは単一の信頼境界です。
┌─────────┐ ┌─────────┐ ┌─────────┐
│ ブラウザ │────>│ ほのワーカー │────>│ MailboxDO │
│ React SPA │ │ (API + SSR) │ │ (SQLite + R2) │
│ エージェントパネル │ │ └─────────┘
━━━━┬───────┘ │ /agents/* ───

──┼────>┌─────────┐
│ │ │ │ EmailAgent DO │
│ WebSocket │ │ │ (AIChatAgent) │
━─────────┤ │ │ 9つのメールツール │
│ │────>│ ワーカー AI │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
ライセンス
AIエージェントを備えた自己ホスト型電子メールクライアントで、完全にCloudflare Workers上で実行されます。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
724
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A self-hosted email client with an AI agent, running entirely on Cloudflare Workers - cloudflare/agentic-inbox

GitHub - cloudflare/agentic-inbox: A self-hosted email client with an AI agent, running entirely on Cloudflare Workers · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
cloudflare
/
agentic-inbox
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits app app public public shared shared workers workers .dev.vars.example .dev.vars.example .gitignore .gitignore LICENSE LICENSE README.md README.md demo_app.png demo_app.png package-lock.json package-lock.json package.json package.json react-router.config.ts react-router.config.ts tsconfig.cloudflare.json tsconfig.cloudflare.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts wrangler.jsonc wrangler.jsonc View all files Repository files navigation
A self-hosted email client with an AI agent, running entirely on Cloudflare Workers
Agentic Inbox lets you send, receive, and manage emails through a modern web interface -- all powered by your own Cloudflare account. Incoming emails arrive via Cloudflare Email Routing , each mailbox is isolated in its own Durable Object with a SQLite database, and attachments are stored in R2 .
An AI-powered Email Agent can read your inbox, search conversations, and draft replies -- built with the Cloudflare Agents SDK and Workers AI .
Read the blog post to learn more about Cloudflare Email Service and how to use it with the Agents SDK, MCP, and from the Wrangler CLI: Email for Agents .
Important : Clicking the 'Deploy to Cloudflare' button is only one part of the setup. You must follow the After deploying steps as well. For a full step-by-step guide with screenshots, refer to this comment:
#4 (comment)
Deploy to Cloudflare. The deploy flow will automatically provision R2, Durable Objects, and Workers AI. You'll be prompted for DOMAINS , which is the domain (yourdomain.com) you want to receive emails for ( email@yourdomain.com ).
Configure Cloudflare Access -- Enable one-click Cloudflare Access on your Worker under Settings > Domains & Routes. The modal will show your POLICY_AUD and TEAM_DOMAIN values. TEAM_DOMAIN can be either your Access team URL or the full .../cdn-cgi/access/certs URL. You must set these as secrets for your Worker.
Set up Email Routing -- In the Cloudflare dashboard, go to your domain > Email Routing and create a catch-all rule that forwards to this Worker
Enable Email Service -- The worker needs the send_email binding to send outbound emails. See Email Service docs
Create a mailbox -- Visit your deployed app and create a mailbox for any address on your domain (e.g. hello@example.com )
If you see Invalid or expired Access token , that usually means POLICY_AUD or TEAM_DOMAIN secrets are incorrect.
Resolution: turn Access off and back on for the Worker to get the Access modal again , then reset your Worker secrets to the latest POLICY_AUD and TEAM_DOMAIN values shown there.
If you see Cloudflare Access must be configured in production , this application is intentionally enforcing Cloudflare Access so your inbox is not exposed to anyone on the internet.
Resolution: enable Access using one-click Cloudflare Access for Workers , then set the POLICY_AUD and TEAM_DOMAIN Worker secrets from the modal values.
Full email client — Send and receive emails via Cloudflare Email Routing with a rich text composer, reply/forward threading, folder organization, search, and attachments
Per-mailbox isolation — Each mailbox runs in its own Durable Object with SQLite storage and R2 for attachments
Built-in AI agent — Side panel with 9 email tools for reading, searching, drafting, and sending
Auto-draft on new email — Agent automatically reads inbound emails and generates draft replies, always requiring explicit confirmation before sending
Configurable and persistent — Custom system prompts per mailbox, persistent chat history, streaming markdown responses, and tool call visibility
Frontend: React 19, React Router v7, Tailwind CSS, Zustand, TipTap, @cloudflare/kumo
Backend: Hono, Cloudflare Workers, Durable Objects (SQLite), R2, Email Routing
AI Agent: Cloudflare Agents SDK ( AIChatAgent ), AI SDK v6, Workers AI ( @cf/moonshotai/kimi-k2.5 ), react-markdown + remark-gfm
Auth: Cloudflare Access JWT validation (required outside local development)
npm install
npm run dev
Configuration
Set your domain in wrangler.jsonc
Create an R2 bucket named agentic-inbox : wrangler r2 bucket create agentic-inbox
Cloudflare account with a domain
Email Routing enabled for receiving
Email Service enabled for sending
Workers AI enabled (for the agent)
Cloudflare Access configured for deployed/shared environments (required in production)
Any user who passes the shared Cloudflare Access policy can access all mailboxes in this app by design. This includes the MCP server at /mcp -- external AI tools (Claude Code, Cursor, etc.) connected via MCP can operate on any mailbox by passing a mailboxId parameter. There is no per-mailbox authorization; the Cloudflare Access policy is the single trust boundary.
┌──────────────┐ ┌──────────────────┐ ┌─────────────────┐
│ Browser │────>│ Hono Worker │────>│ MailboxDO │
│ React SPA │ │ (API + SSR) │ │ (SQLite + R2) │
│ Agent Panel │ │ │ └─────────────────┘
└──────┬───────┘ │ /agents/* ──────┼────>┌─────────────────┐
│ │ │ │ EmailAgent DO │
│ WebSocket │ │ │ (AIChatAgent) │
└─────────────┤ │ │ 9 email tools │
│ │────>│ Workers AI │
└──────────────────┘ └─────────────────┘
License
A self-hosted email client with an AI agent, running entirely on Cloudflare Workers
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
724
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
