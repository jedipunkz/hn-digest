---
source: "https://github.com/acip/slack-claude-agent"
hn_url: "https://news.ycombinator.com/item?id=48751448"
title: "Show HN: Self-hosted Slack bot Claude Tag alternative"
article_title: "GitHub - acip/slack-claude-agent: Self-hosted, open source Claude in Slack alternative, for you or a small team. · GitHub"
author: "acip"
captured_at: "2026-07-01T19:33:12Z"
capture_tool: "hn-digest"
hn_id: 48751448
score: 1
comments: 0
posted_at: "2026-07-01T18:45:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Self-hosted Slack bot Claude Tag alternative

- HN: [48751448](https://news.ycombinator.com/item?id=48751448)
- Source: [github.com](https://github.com/acip/slack-claude-agent)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T18:45:31Z

## Translation

タイトル: HN を表示: 自己ホスト型 Slack ボット Claude Tag の代替案
記事のタイトル: GitHub - acip/slack-claude-agent: Slack の代替となるセルフホスト型のオープンソース Claude、あなたまたは小規模チーム向け。 · GitHub
説明: セルフホスト型のオープンソース Claude in Slack の代替品で、あなたまたは小規模チーム向けです。 - acip/slack-claude-agent

記事本文:
GitHub - acip/slack-claude-agent: あなたまたは小規模チーム向けの、セルフホスト型のオープンソース Claude in Slack の代替品。 · GitHub
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
アシプ
/
スラッククロードエージェント
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md claude-slack.js claude-slack.js エコシステム.config.js エコシステム.config.js package-lock.json package-lock.json package.json package.json プロンプト.example.md プロンプト.example.md サーバー.js サーバー.js スラック-アプリ-マニフェスト.yaml スラック-アプリ-マニフェスト.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたまたは小規模チーム向けの、自己ホスト型のオープンソースの Claude in Slack の代替品です。
Slack スレッドでクロードを @メンションすると、ファイルにアクセスされる前に、ストリーミング応答、インタラクティブな質問カード、計画承認ゲートを取得できます。スレッドごとに 1 つの Claude Agent SDK セッションが、自分のマシンとワークスペースで実行されます。
Slack の Anthropic の公式 Claude は、チーム プランとエンタープライズ プランで利用できます。これは、誰でも実行できる自己ホスト可能な代替手段です。自分のマシン、自分の Anthropic 資格情報、および 1 つのプロジェクト ディレクトリを持ち込むことができます。
チャンネル内のボットを @メンションします。そのスレッドのクロード セッションを開き、返信を単一のライブ更新 Slack メッセージにストリーミングして返します。
同じスレッド内でフォローアップの返信を行うと、再度言及する必要がなく会話が継続されます。各スレッドは独自のクロード セッションを保持し、セッションは再起動後も存続します。
クロードが決定を必要とするとき、その質問はネイティブの Slack ラジオまたはチェックボックス カードとしてレンダリングされます。選択するとスレッドが再開されます。
クロードがファイルを変更したいときは、最初に計画を提示する必要があります。 Slack で承認を押すまで、書き込みツールはロックされたままになります。 [変更のリクエスト] をクリックしてフィードバックを送信することもできます。
アクションの前に明示的にプランを要求するには、メッセージの前に plan を付けます。
インストールする前にセキュリティモデルをお読みください。このツールは、Claude re を提供します。

Slack メッセージによって駆動される、ホスト上のディレクトリへのファイル システムおよびシェル アクセス。
専用の VM (または少なくともコンテナー)。これは必須であり、あれば便利というわけではありません。エージェントはシェル コマンドを実行できるため、運用資格情報、SSH キー、またはその他のプロジェクトを保持するボックスではなく、独自の低特権サンドボックスを与えます。
クロード エージェント SDK 認証。ホスト上のログインしたクロード CLI、またはプロセス環境の ANTHROPIC_API_KEY のいずれか。 SDK はこのサーバーを実行するユーザーとして Claude を生成するため、ユーザーは認証される必要があります。
アプリを作成してインストールできる Slack ワークスペース。
このサーバーのポートに転送するパブリック HTTPS エンドポイント。 nginx 、 Traefik 、 Caddy 、 Cloudflared 、または ngrok などのリバース プロキシまたはトンネルが機能します。 Cloudflare はオプションです。以下の例では、すぐに開始できるので、cloudflared を使用しています。 「 Cloudflare トンネルのセットアップ 」を参照してください。
クローンを作成してインストールします。
git clone https://github.com/acip/slack-claude-agent.git
cd スラック-クロード-エージェント
npmインストール
HTTPS 経由でポートを公開します。使用する予定のポート (デフォルトは 3999) の前にパブリック HTTPS エンドポイントを配置し、その URL をコピーします。好みのリバース プロキシまたはトンネルを使用してください (「要件」を参照)。クラウドフレアを使用した簡単な使い捨てトンネル:
クラウドフレアトンネル --url http://localhost:3999
安定したホスト名については、「 Cloudflare トンネルのセットアップ 」を参照してください。
Slack アプリを作成します。 api.slack.com/apps に移動し、 [Create New App] を選択し、 [From a manifest] を選択して、lack-app-manifest.yaml を貼り付けます。まず、その中の your-tunnel-host.example.com をトンネルのホスト名に置き換えます。マニフェストはプライベート チャネル スコープ ( groups:history 、 message.groups ) も要求します。パブリック チャネルのみが必要な場合は、それらを削除してください。
両方のリクエスト URL を https://<your-tunnel-host>/slack/events に設定します。イベント サブスクリプションとインタラクティビティの 2 つの場所で行われます。

ショートカット。人々はインタラクティブ性を見逃してしまい、ボタンを押しても何も機能しません。
アプリをワークスペースにインストールし、ボット ユーザー OAuth トークン ( xoxb-... ) と署名シークレットをコピーします。
環境を構成します。
cp .env.example .env
SLACK_BOT_TOKEN 、 SLACK_SIGNING_SECRET 、 PORT 、および PROJECT_WORKSPACE を入力します。
ペルソナをカスタマイズします (オプション)。
cp プロンプト.example.md プロンプト.md
プロンプト.md を編集します。 gitignore されているため、バージョンは非公開のままです。
ボットを招待してメンションします。 Slack で、チャネル内で /invite @claude を実行し、メンションします: @claude What does this project do?
Cloudflare はオプションです (リバース プロキシまたはトンネルはどれでも機能します) が、cloudflared はパブリック HTTPS URL を簡単に取得する方法です。 Cloudflare のダウンロード ( brew install cloudflared 、 apt install cloudflared 、または直接バイナリ) からインストールします。
簡単な使い捨てトンネル (ランダムな *.trycloudflare.com URL、アカウントは不要、テストに適しています):
クラウドフレアトンネル --url http://localhost:3999
HTTPS URL を出力します。この URL は実行するたびに変更されるため、毎回 Slack に貼り付ける必要があります。
名前付きの永続的なトンネル (Cloudflare アカウントのドメイン上の安定したホスト名、運用環境に最適):
# 1. 認証 (ブラウザを開き、使用するドメインを選択します)
クラウドフレアトンネルログイン
# 2. トンネルとその認証情報ファイルを作成する
クラウドフレアトンネル作成slack-claude-agent
# 3. ホスト名を指定します (DNS レコードを作成します)。
クラウドフレアトンネルルートDNSスラッククロードエージェントslack-bot.yourdomain.com
# 4. ローカルポートに転送して実行します。
Cloudflared トンネルの実行 --url http://localhost:3999lack-claude-agent
安定した Slack リクエスト URL は https://slack-bot.yourdomain.com/slack/events になります。実行し続けるには、cloudflared サービス インストールを使用してサービスとしてインストールするか (Cloudflare のドキュメントを参照)、アプリと一緒に pm2 で管理します。
すべてc

設定は環境変数 ( .env からロード) によって行われます。
注: このアプリは HTTP と署名シークレットを使用します。ソケット モードを使用しないため、設計上、ここには SLACK_APP_TOKEN はありません。
ボットは、Claude Code プリセットの上にシステム プロンプトを追加します。ペルソナまたはハウス ルールを変更するには、prompt.example.md をプロンプト.md にコピーして編集します。プロンプト.md も AGENT_PROMPT_FILE のファイルも存在しない場合は、組み込みのデフォルトが使用され、起動時に警告が記録されます。プロンプトは起動時に一度読み取られるため、編集後にサーバーを再起動します (または pm2 がlack-claude-agentを再起動します)。
MCP サーバー (Notion、Google Drive、GitHub など) を追加する
ボットは Claude Agent SDK 上で実行されるため、MCP サーバーを使用して Notion、Google Drive、GitHub、または独自のツールにアクセスできます。 2 つの部分があります。SDK にサーバーについて通知することと、このアプリの権限ゲートでツールを実行させることです。
1. サーバーを宣言します。 SDK は、CLAUDE_SETTING_SOURCES に従ってクロード設定から MCP サーバーをロードします (構成を参照)。 2 つの一般的なオプション:
プロジェクトごと: PROJECT_WORKSPACE ディレクトリに .mcp.json ファイルを作成します。デフォルトの CLAUDE_SETTING_SOURCES には project が含まれているため、自動的に選択されます。 ${VAR} はロード時に環境変数を展開します。
{
"mcpサーバー": {
「概念」: {
"type" : " http " ,
"url" : " https://mcp.notion.com/mcp " ,
"headers" : { "Authorization" : " Bearer ${NOTION_TOKEN} " }
}、
「ファイルシステム」: {
"コマンド" : " npx " ,
"args" : [ " -y " 、 " @modelcontextprotocol/server-filesystem " 、 " /data " ]
}
}
}
グローバル: 独自の ~/.claude 構成でサーバーを構成し、ボットがそれを継承するように CLAUDE_SETTING_SOURCES=user,project,local を設定します。これを行う前に、ユーザー層に関するセキュリティ上の警告に注意してください。
2. ツールを許可します。 MCP ツールの名前は mcp__<server>__<tool> であり、このアプリの pe

許可ゲートはデフォルトで不明なツールをブロックするため、オプトインするまで MCP ツールは拒否されたままになります。.env で ALLOW_MCP_TOOLS=true を設定して再起動します。
信頼に関する警告。 ALLOW_MCP_TOOLS がオンの場合、MCP ツールは自動的に実行され、計画承認の書き込みゲートを通過しません。 MCP サーバーは、接続先であれば何でも読み書きできます。信頼できるサーバーのみに接続し、タスクに必要な範囲でトークンの範囲を狭くします (可能な場合は読み取り専用)。
詳細は、Claude Agent SDK のドキュメント (Agent SDK の MCP および TypeScript SDK リファレンス ( mcpServers 、settingSources 、permissions)) に記載されています。
.env に CLAUDE_MODEL を設定します。 claude-opus-4-8 が最高品質で、claude-sonnet-5 (デフォルト) がコストと速度のバランスです。
アクション
どうやって
スレッドを開始する
ボットが存在するチャンネルの @claude <あなたの質問>。
スレッドを続行する
スレッド内で返信するだけです。改めて言及する必要はありません。
まずはプランを聞いてみる
plan: というプレフィックスを付けます。たとえば、 plan: auth module をリファクタリングします。
変更を承認する
[承認] を押して、プラン カードに進みます。これにより、そのスレッドの書き込みのロックが解除されます。
計画に関するフィードバックを送信する
[変更をリクエスト] を押して、変更する内容を入力します。
質問に答える
カード上のオプションを選択し、 [送信] を押します。
実稼働環境での実行
プロセス マネージャーと永続的な HTTPS エンドポイントを使用します。
pm2 はエコシステム.config.js を開始します
pm2 ログのスラック-クロード-エージェント
pm2 再起動slack-claude-agent
リバース プロキシまたはトンネル (nginx、Traefik、Caddy、cloudflared、ngrok など、選択したもの) をアドホック ターミナル コマンドではなくマネージド サービスとして実行することで、パブリック URL が安定した状態を保ちます。ホスト名が変更された場合は、両方の Slack リクエスト URL を更新します。
このツールは、実際のファイルシステムと、Slack メッセージによって構成されたワークスペース ディレクトリへのシェル アクセスを使用して、ホスト上で Claude Agent SDK を実行します。ボットが存在するチャンネルに投稿できる人は誰でも Autono をタスク化できます

あなたのマシン上のマウスエージェント。ユーザーごとの権限はありません。チャネル メンバーシップはアクセス制御です。
コードから直接パーミッション ゲートが実際にどのように機能するか:
読み取りツール (Read、Glob、Grep、WebSearch、WebFetch、TodoWrite) は自動的に無人で実行されます。
読み取り専用の Bash コマンド ( find 、 ls 、 cat 、 grep 、 git log 、curl など) の小さな許可リストも自動的に実行されます。これをセキュリティ境界ではなく、便宜的なヒューリスティックとして扱います。 cat やcurl などのコマンドは任意のデータを読み取ったり送信したりすることができ、プレフィックス ベースの許可リストは既知のソフト スポットです。摩擦を下げてくれます。 Bash を安全にするわけではありません。
書き込みツール (Edit、MultiEdit、Write、NotebookEdit、および完全な Bash) は、計画が承認されるまでスレッドごとにロックされます。承認により、その 1 つのスレッドに対してのみ書き込みのロックが解除されます。
承認は粗めです。計画を承認すると、そのスレッドの残りの部分ですべての書き込みツールのロックが解除されます。これはコマンドごとの確認ではありません。承認する前に計画を確認してください。これは、読み取りツールが自動的に取り込むファイルからのプロンプト インジェクションに対する主な防御策でもあります。
本物らしさを緩めます。受信イベントは、 SLACK_SIGNING_SECRET を使用して Slack Bolt ライブラリによって検証されます。署名されていないリクエストまたは予期しないリクエストは 404 を受け取ります。404 はノイズを軽減します。

[切り捨てられた]

## Original Extract

Self-hosted, open source Claude in Slack alternative, for you or a small team. - acip/slack-claude-agent

GitHub - acip/slack-claude-agent: Self-hosted, open source Claude in Slack alternative, for you or a small team. · GitHub
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
acip
/
slack-claude-agent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md claude-slack.js claude-slack.js ecosystem.config.js ecosystem.config.js package-lock.json package-lock.json package.json package.json prompt.example.md prompt.example.md server.js server.js slack-app-manifest.yaml slack-app-manifest.yaml View all files Repository files navigation
Self-hosted, open source Claude in Slack alternative, for you or a small team.
@mention Claude in any Slack thread and get streaming replies, interactive question cards, and plan-approval gating before it can touch your files. One Claude Agent SDK session per thread, running on your own machine and workspace.
Anthropic's official Claude in Slack is available on Team and Enterprise plans. This is a self-hostable alternative that anyone can run: you bring your own machine, your own Anthropic credentials, and one project directory.
You @mention the bot in a channel. It opens a Claude session for that thread and streams the reply back into a single live-updating Slack message.
Follow-up replies in the same thread continue the conversation with no re-mention needed. Each thread keeps its own Claude session, and sessions survive restarts.
When Claude needs a decision, its questions render as native Slack radio or checkbox cards. Your selection resumes the thread.
When Claude wants to change files, it must present a plan first. Write tools stay locked until you press Approve in Slack. You can also press Request changes to send feedback.
Prefix a message with plan: to explicitly ask for a plan before any action.
Please read the Security model before installing. This tool gives Claude real filesystem and shell access to a directory on the host, driven by Slack messages.
A dedicated VM (or at minimum a container). This is a requirement, not a nice-to-have. The agent can run shell commands, so give it its own low-privilege sandbox rather than a box that holds production credentials, SSH keys, or other projects.
Claude Agent SDK authentication. Either a logged-in claude CLI on the host, or an ANTHROPIC_API_KEY in the process environment. The SDK spawns Claude as the user running this server, so that user must be authenticated.
A Slack workspace where you can create and install an app.
A public HTTPS endpoint that forwards to this server's port. Any reverse proxy or tunnel works: nginx , Traefik , Caddy , cloudflared , or ngrok . Cloudflare is optional; the examples below just use cloudflared because it is quick to start. See Set up a Cloudflare tunnel .
Clone and install.
git clone https://github.com/acip/slack-claude-agent.git
cd slack-claude-agent
npm install
Expose the port over HTTPS. Put a public HTTPS endpoint in front of the port you plan to use (default 3999) and copy its URL. Use whichever reverse proxy or tunnel you prefer (see Requirements). Quick throwaway tunnel with cloudflared:
cloudflared tunnel --url http://localhost:3999
For a stable hostname, see Set up a Cloudflare tunnel .
Create the Slack app. Go to api.slack.com/apps , choose Create New App , then From a manifest , and paste slack-app-manifest.yaml . Replace your-tunnel-host.example.com in it with your tunnel hostname first. The manifest requests private-channel scopes ( groups:history , message.groups ) too; remove them if you only need public channels.
Set both request URLs to https://<your-tunnel-host>/slack/events . It goes in two places: Event Subscriptions and Interactivity & Shortcuts . People miss the Interactivity one, and then buttons do nothing.
Install the app to your workspace and copy the Bot User OAuth Token ( xoxb-... ) and the Signing Secret .
Configure the environment.
cp .env.example .env
Fill in SLACK_BOT_TOKEN , SLACK_SIGNING_SECRET , PORT , and PROJECT_WORKSPACE .
Customize the persona (optional).
cp prompt.example.md prompt.md
Edit prompt.md . It is gitignored, so your version stays private.
Invite and mention the bot. In Slack, run /invite @claude in a channel, then mention it: @claude what does this project do?
Cloudflare is optional (any reverse proxy or tunnel works), but cloudflared is a quick way to get a public HTTPS URL. Install it from Cloudflare's downloads ( brew install cloudflared , apt install cloudflared , or a direct binary).
Quick, throwaway tunnel (random *.trycloudflare.com URL, no account needed, good for testing):
cloudflared tunnel --url http://localhost:3999
It prints an HTTPS URL. That URL changes every run, so you would re-paste it into Slack each time.
Named, persistent tunnel (stable hostname on a domain in your Cloudflare account, best for production):
# 1. Authenticate (opens a browser; pick the domain to use)
cloudflared tunnel login
# 2. Create a tunnel and its credentials file
cloudflared tunnel create slack-claude-agent
# 3. Point a hostname at it (creates the DNS record for you)
cloudflared tunnel route dns slack-claude-agent slack-bot.yourdomain.com
# 4. Run it, forwarding to the local port
cloudflared tunnel run --url http://localhost:3999 slack-claude-agent
Your stable Slack request URL is then https://slack-bot.yourdomain.com/slack/events . To keep it running, install it as a service with cloudflared service install (see Cloudflare's docs), or manage it under pm2 alongside the app.
All configuration is through environment variables (loaded from .env ).
Note: this app uses HTTP plus your signing secret. It does not use Socket Mode, so there is no SLACK_APP_TOKEN here by design.
The bot appends a system prompt on top of the Claude Code preset. To change its persona or house rules, copy prompt.example.md to prompt.md and edit it. If neither prompt.md nor a file at AGENT_PROMPT_FILE exists, a built-in default is used and a warning is logged at boot. The prompt is read once at startup, so restart the server (or pm2 restart slack-claude-agent ) after editing.
Add MCP servers (Notion, Google Drive, GitHub, and more)
The bot runs on the Claude Agent SDK, so it can use MCP servers to reach tools like Notion, Google Drive, GitHub, or your own. There are two parts: telling the SDK about the server, and letting this app's permission gate run the tool.
1. Declare the server. The SDK loads MCP servers from your Claude settings according to CLAUDE_SETTING_SOURCES (see Configuration). Two common options:
Per project: create a .mcp.json file in your PROJECT_WORKSPACE directory. Because the default CLAUDE_SETTING_SOURCES includes project , it is picked up automatically. ${VAR} expands environment variables at load time.
{
"mcpServers" : {
"notion" : {
"type" : " http " ,
"url" : " https://mcp.notion.com/mcp " ,
"headers" : { "Authorization" : " Bearer ${NOTION_TOKEN} " }
},
"filesystem" : {
"command" : " npx " ,
"args" : [ " -y " , " @modelcontextprotocol/server-filesystem " , " /data " ]
}
}
}
Global: configure the server in your own ~/.claude config, then set CLAUDE_SETTING_SOURCES=user,project,local so the bot inherits it. Note the Security caveat about the user layer before doing this.
2. Allow the tools. MCP tools are named mcp__<server>__<tool> , and this app's permission gate blocks unknown tools by default, so MCP tools stay denied until you opt in. Set ALLOW_MCP_TOOLS=true in .env and restart.
Trust caveat. When ALLOW_MCP_TOOLS is on, MCP tools run automatically and are not behind the plan-approval write gate. An MCP server can read and write in whatever it connects to. Only connect servers you trust, and scope their tokens as narrowly as the task needs (read-only where possible).
The details live in the Claude Agent SDK docs: MCP in the Agent SDK and the TypeScript SDK reference ( mcpServers , settingSources , permissions).
Set CLAUDE_MODEL in .env . claude-opus-4-8 is the highest quality, claude-sonnet-5 (the default) is the cost and speed balance.
Action
How
Start a thread
@claude <your question> in a channel the bot is in.
Continue a thread
Just reply in the thread. No re-mention needed.
Ask for a plan first
Prefix with plan: , for example plan: refactor the auth module .
Approve changes
Press Approve & proceed on the plan card. This unlocks writes for that thread.
Send feedback on a plan
Press Request changes and type what should change.
Answer a question
Pick options on the card and press Submit .
Running in production
Use a process manager and a persistent HTTPS endpoint.
pm2 start ecosystem.config.js
pm2 logs slack-claude-agent
pm2 restart slack-claude-agent
Run your reverse proxy or tunnel (nginx, Traefik, Caddy, cloudflared, ngrok, whichever you chose) as a managed service rather than an ad hoc terminal command, so the public URL stays stable. If the hostname changes, update both Slack request URLs.
This tool runs the Claude Agent SDK on your host with real filesystem and shell access to the workspace directory you configure, driven by Slack messages. Anyone who can post in a channel the bot is in can task an autonomous agent on your machine. There is no per-user authorization. Channel membership is the access control.
How the permission gate actually works, straight from the code:
Read tools (Read, Glob, Grep, WebSearch, WebFetch, TodoWrite) run automatically and unattended.
A small allowlist of read-only Bash commands (things like find , ls , cat , grep , git log , curl ) also runs automatically. Treat this as a convenience heuristic, not a security boundary. Commands like cat and curl can read or send arbitrary data, and prefix based allowlisting is a known soft spot. It lowers friction. It does not make Bash safe.
Write tools (Edit, MultiEdit, Write, NotebookEdit, and full Bash) are locked per thread until you approve a plan. Approval unlocks writes for that one thread only.
Approval is coarse. Approving a plan unlocks all write tools for the rest of that thread. It is not a per-command confirmation. Review plans before approving. This is also your main defense against prompt injection from files the read tools ingest automatically.
Slack authenticity. Inbound events are verified by the Slack Bolt library using your SLACK_SIGNING_SECRET . Unsigned or unexpected requests get a 404. The 404 is noise reduc

[truncated]
