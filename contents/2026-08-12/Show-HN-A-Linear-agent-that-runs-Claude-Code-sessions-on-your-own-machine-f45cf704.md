---
source: "https://github.com/MPIsaac-Per/linear-claude-bridge"
hn_url: "https://news.ycombinator.com/item?id=49278206"
title: "Show HN: A Linear agent that runs Claude Code sessions on your own machine"
article_title: "GitHub - MPIsaac-Per/linear-claude-bridge: Run your Claude Code context as a Linear agent · GitHub"
author: "MPIsaac-Per"
captured_at: "2026-08-12T21:35:58Z"
capture_tool: "hn-digest"
hn_id: 49278206
score: 1
comments: 0
posted_at: "2026-08-12T20:37:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A Linear agent that runs Claude Code sessions on your own machine

- HN: [49278206](https://news.ycombinator.com/item?id=49278206)
- Source: [github.com](https://github.com/MPIsaac-Per/linear-claude-bridge)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T20:37:06Z

## Translation

タイトル: HN を表示: 自分のマシン上でクロード コード セッションを実行する線形エージェント
記事のタイトル: GitHub - MPIsaac-Per/linear-claude-bridge: クロード コード コンテキストを線形エージェントとして実行する · GitHub
説明: クロード コード コンテキストを線形エージェントとして実行します。 GitHub でアカウントを作成して、MPIsaac-Per/linear-claude-bridge の開発に貢献してください。
HN テキスト: 私はすべてをローカルの Obsidian ナレッジ ベースに保存し、それに対してクロード コードを一日中実行しています。私もリニアに住んでいます。リニアチケット内でナレッジベースの質問をしたかったのです。ということで、小さな橋を作りました。 Linear OAuth アプリをエージェントとして登録し、メンションするか問題を委任すると、選択したディレクトリを cwd でポイントした状態で、自分のマシン上で Claude Agent SDK セッションが実行されます。これは単なる作業ディレクトリであるため、CLAUDE.md および MCP サーバーがロードされます。返信は問題のエージェント スレッドに届き、フォローアップは同じセッションを再開します。サブスクリプション認証 (少なくとも現時点では)、API キーなし。 TypeScript、946 行、60 テスト、MIT。これはコンテキスト エージェント レーンをカバーします。問題を割り当てて PR を取得するには、Cyrus がすでに存在しており、非常に優れています。

記事本文:
GitHub - MPIsaac-Per/linear-claude-bridge: クロード コード コンテキストを線形エージェントとして実行する · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
MPIsaac-Per
/
リニアクロードブリッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット デプロイ デプロイ src src test test .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json

package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コード セットアップを線形エージェントとして実行します。問題をそれに委任するか、
エージェント セッションでメッセージを送信すると、Claude Agent SDK セッションが実行されます。
あなたのマシン、あなたが選択した作業ディレクトリに、すべてのものが含まれています
ディレクトリには、CLAUDE.md 命令、MCP サーバー、スキルが含まれています。返信
問題のエージェント セッション スレッドに到達します。
これは最小限のリファレンス実装です (1,000 行未満、フレームワークなし、
テスト済み）。これはコーディングエージェントではありません。課題の割り当て、PR の取得フローの場合、
キュロスを参照。この橋は
あなたのコンテキストを知っているエージェント（知識ベース、運用担当者）と話す
リポジトリ、プロジェクト ディレクトリ。
クロード コードのサブスクリプション認証で実行されます。 Anthropic API キーはありません。
線形 (言及/委任/フォローアップ プロンプト)
-> Webhook: AgentSessionEvent (作成された | プロンプトが表示された)
-> src/server.ts で HMAC を検証、ACK < 5 秒、最初のアクティビティ < 10 秒
-> src/queue.ts シリアル実行、同時実行 1
-> src/runtime/claude.ts エージェント SDK query()、cwd = KB_PATH、再開
-> src/linear/client.ts AgentActivityCreate (思考/行動/応答)
セッション マッピング (リニア セッション ID -> SDK セッション ID) は JSON で保持されます
ファイルに保存されるため、フォローアップ プロンプトによって同じ会話が再開されます。
ノード 22+、稼働し続けるマシン、およびクロード コード
がインストールされ、サービスを実行するユーザーとしてログインします。
OAuth アプリケーションを作成できるリニア ワークスペース。
サービスへのパブリック HTTPS ルート (テールスケール ファネル、クラウドフレア、または
任意のリバース プロキシ)。
1. リニア OAuth アプリを作成する
リニア設定 -> API -> アプリケーション -> 新しいアプリケーション:
名前を付けて (これはエージェントの表示名です)、開発者フィールドに入力します。
リダイレクト URI: http://localhost:3979/oauth/callback
Webhook をオンに切り替えます。事前に生成された署名シークレットに注目してください。
Webhook URL: パブリック HTTPS ホスト +

/webhook (後で修正できます)。
[アプリ イベント] で、[エージェント セッション イベント] をチェックします。スコープはありません
アプリ上のピッカー。スコープはインストール時に要求されます。
npm install && cp .env.example .env
# LINEAR_CLIENT_ID、LINEAR_CLIENT_SECRET、LINEAR_WEBHOOK_SECRET を入力します
# set LINEAR_ACCESS_TOKEN=pending (ステップ 3 までのプレースホルダー)
# エージェントがコンテキストを保持するディレクトリに KB_PATH を設定します
npm 実行開発
3. アプリをエージェントとしてインストールします (actor=app)
この URL (クライアント ID に置き換えられます) を開き、インストールを承認します。
https://linear.app/oauth/authorize?client_id=<CLIENT_ID>&redirect_uri=http%3A%2F%2Flocalhost%3A3979%2Foauth%2Fcallback&response_type=code&scope=read,write,app:assignable,app:mentionable&actor=app
サービスはコールバックでアクセス トークンを出力します。として .env に置きます
LINEAR_ACCESS_TOKEN を実行して再起動します。アプリが割り当て可能なものとして表示されるようになりました
ワークスペースのエージェント。
macOS では、./deploy/install.sh がサービスを構築し、launchd をインストールします。
ユーザー エージェント (SDK がクロード コードの認証情報を認識するため)、
tailscale ファネルを作成し、アプリ構成に貼り付ける Webhook URL を出力します。
他の HTTPS イングレスは機能します。サービスには POST /webhook のみが必要です
到達可能。 TLS ターミネーターが別のホストで実行されている場合、
deploy/tcp_forward.py は、依存関係のない TCP フォワーダーであり、
プライベート ネットワーク経由のラストホップ。
問題がある場合はエージェントに委任するか、セッション スレッドでメッセージを送信してください。
最初の考えは数秒以内に思いつきます。回答には実際のエージェントと同じくらい時間がかかります
セッションがかかります。
フィールドノート（文書では説明されないこと）
プロンプトが表示された Webhook には、ユーザーのテキストが含まれます。
AgentActivity.body ではなく、 AgentActivity.content.body 。を読む
フィールドが間違っていると、エージェントは「待機中」と答えます。すべてのものに、
空のプロンプトで再開されるためです。
AgentSessionEvent ペイロードはフィールドを次の場所に置きます

のトップレベル
Webhook 本体 (データ変更 Webhook とは異なり、データ ラッパーはありません)。
HMAC-SHA256 署名 (リニア署名ヘッダー) は、生のデータをカバーします。
体;リプレイ保護のタイムスタンプは、内部の webhookTimestamp です。
JSON、ミリ秒、60 秒を超えるスキューを拒否します。
作業を行う前に Webhook に確認応答し (5 秒制限)、思考を送信します。
作成後すぐに (10 秒の生存制限)、または Linear でマークされます。
セッションが応答しません。
ランタイム実行をシリアルに保ちます。同時ヘッドレス クロード セッション
1 つのホストがセッション間のコンテンツ汚染を引き起こしました。
PermissionMode: 「bypassPermissions」はそれなしでは何も行いません
allowDangerouslySkipPermissions: true ; SDK にはペアが必要です。
古いエージェント SDK バージョン (0.1.x) では、トランスクリプトが含まれるセッションの再開に失敗します。
空のテキスト ブロックが含まれています (API 400: "テキスト コンテンツ ブロックは
空ではない」）。 0.3.x以降を使用してください。
Webhook エンドポイントは署名を検証し、古いタイムスタンプを拒否します。
他のルートは /healthz と /oauth/callback だけです。理解する
あなたが結んでいる内容: リニアでエージェントについて言及できる人全員
ワークスペースは、権限を使用して実行されている無人エージェント セッションを操作します
KB_PATH でバイパスされます。信頼できるワークスペースで使用し、KB_PATH をスコープします。
意図的に。
Claude Agent SDK は現在、Claude Code サブスクリプションを利用しています。
資格情報と標準プランの制限。 Anthropic は次のように発表しましたが、その後一時停止されました。
Agent SDK の使用量を別の従量制クレジット プールに移動する変更。
経済状況に依存する前に、現在の条件を確認してください。
クロード コード コンテキストを線形エージェントとして実行する
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Run your Claude Code context as a Linear agent. Contribute to MPIsaac-Per/linear-claude-bridge development by creating an account on GitHub.

I keep everything in a local Obsidian knowledge base and run Claude Code against it all day. I also live in Linear. I wanted to ask the knowledge base questions within the Linear ticket. So, I created a small bridge. Register a Linear OAuth app as an agent, mention it or delegate an issue, and it runs a Claude Agent SDK session on your own machine with cwd pointed at whatever directory you choose. CLAUDE.md and MCP servers load because it's just a working directory. Replies land in the issue's agent thread, and follow-ups resume the same session. Subscription auth (at least for now), no API key. TypeScript, 946 lines, 60 tests, MIT. It covers the context-agent lane. For assign-an-issue-get-a-PR, Cyrus already exists and is pretty dang good.

GitHub - MPIsaac-Per/linear-claude-bridge: Run your Claude Code context as a Linear agent · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
MPIsaac-Per
/
linear-claude-bridge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits deploy deploy src src test test .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Run your Claude Code setup as a Linear agent. Delegate an issue to it or
message it in an agent session, and a Claude Agent SDK session runs on
your machine, in a working directory you choose, with everything that
directory carries: CLAUDE.md instructions, MCP servers, skills. Replies
land in the issue's agent-session thread.
This is a minimal reference implementation (under 1,000 lines, no framework,
tested). It is not a coding agent; for assign-an-issue-get-a-PR flows,
see Cyrus . This bridge is for
talking to an agent that knows your context: a knowledge base, an ops
repo, a project directory.
Runs on Claude Code subscription auth. No Anthropic API key.
Linear (mention / delegate / follow-up prompt)
-> webhook: AgentSessionEvent (created | prompted)
-> src/server.ts verify HMAC, ack < 5s, first activity < 10s
-> src/queue.ts serial execution, concurrency 1
-> src/runtime/claude.ts Agent SDK query(), cwd = KB_PATH, resume
-> src/linear/client.ts agentActivityCreate (thought/action/response)
Session mapping (Linear session id -> SDK session id) persists in a JSON
file, so follow-up prompts resume the same conversation.
Node 22+, a machine that stays on, and Claude Code
installed and logged in as the user who runs the service.
A Linear workspace where you can create OAuth applications.
A public HTTPS route to the service (tailscale funnel, cloudflared, or
any reverse proxy).
1. Create the Linear OAuth app
Linear Settings -> API -> Applications -> new application:
Name it (this is the agent's visible name), fill developer fields.
Redirect URI: http://localhost:3979/oauth/callback
Toggle Webhooks on. Note the pre-generated signing secret.
Webhook URL: your public HTTPS host + /webhook (can be corrected later).
Under App events , check Agent session events . There is no scopes
picker on the app; scopes are requested at install time.
npm install && cp .env.example .env
# fill LINEAR_CLIENT_ID, LINEAR_CLIENT_SECRET, LINEAR_WEBHOOK_SECRET
# set LINEAR_ACCESS_TOKEN=pending (placeholder until step 3)
# set KB_PATH to the directory whose context the agent should carry
npm run dev
3. Install the app as an agent (actor=app)
Open this URL (your client id substituted), approve the install:
https://linear.app/oauth/authorize?client_id=<CLIENT_ID>&redirect_uri=http%3A%2F%2Flocalhost%3A3979%2Foauth%2Fcallback&response_type=code&scope=read,write,app:assignable,app:mentionable&actor=app
The service prints the access token on the callback; put it in .env as
LINEAR_ACCESS_TOKEN and restart. The app now appears as an assignable
agent in your workspace.
On macOS, ./deploy/install.sh builds the service, installs a launchd
user agent (so the SDK sees your Claude Code credentials), opens a
tailscale funnel, and prints the webhook URL to paste into the app config.
Any other HTTPS ingress works; the service only needs POST /webhook
reachable. If your TLS terminator runs on a different host,
deploy/tcp_forward.py is a dependency-free TCP forwarder to bridge the
last hop over a private network.
Delegate any issue to the agent, or message it in the session thread.
First thought lands within seconds; answers take as long as a real agent
session takes.
Field notes (things the docs won't tell you)
The prompted webhook carries the user's text at
agentActivity.content.body , not agentActivity.body . Reading the
wrong field yields an agent that answers "Standing by." to everything,
because it is resumed with an empty prompt.
AgentSessionEvent payloads put their fields at the top level of the
webhook body (no data wrapper, unlike data-change webhooks).
The HMAC-SHA256 signature ( linear-signature header) covers the raw
body; the replay-protection timestamp is webhookTimestamp inside the
JSON, milliseconds, reject beyond 60s skew.
Ack the webhook before doing any work (5s limit) and emit a thought
immediately on created (10s liveness limit), or Linear marks the
session unresponsive.
Keep runtime execution serial. Concurrent headless Claude sessions on
one host have produced cross-session content contamination.
permissionMode: "bypassPermissions" does nothing without
allowDangerouslySkipPermissions: true ; the SDK requires the pair.
Old Agent SDK versions (0.1.x) fail to resume sessions whose transcript
contains empty text blocks (API 400: "text content blocks must be
non-empty"). Use 0.3.x or later.
The webhook endpoint verifies signatures and rejects stale timestamps;
/healthz and /oauth/callback are the only other routes. Understand
what you are wiring up: anyone who can mention the agent in your Linear
workspace steers an unattended agent session running with permissions
bypassed in KB_PATH . Use it in workspaces you trust, and scope KB_PATH
deliberately.
The Claude Agent SDK currently draws on Claude Code subscription
credentials and standard plan limits. Anthropic announced, then paused, a
change that would move Agent SDK usage to a separate metered credit pool.
Check current terms before depending on the economics.
Run your Claude Code context as a Linear agent
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
