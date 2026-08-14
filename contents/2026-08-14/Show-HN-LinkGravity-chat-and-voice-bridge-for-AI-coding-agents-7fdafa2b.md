---
source: "https://github.com/dev-sseul/linkgravity"
hn_url: "https://news.ycombinator.com/item?id=49297348"
title: "Show HN: LinkGravity – chat and voice bridge for AI coding agents"
article_title: "GitHub - dev-sseul/linkgravity: A messenger interface for the Antigravity AI agent. Features interactive CLI tool approvals and voice control · GitHub"
author: "sseul"
captured_at: "2026-08-14T11:37:03Z"
capture_tool: "hn-digest"
hn_id: 49297348
score: 1
comments: 0
posted_at: "2026-08-14T11:28:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LinkGravity – chat and voice bridge for AI coding agents

- HN: [49297348](https://news.ycombinator.com/item?id=49297348)
- Source: [github.com](https://github.com/dev-sseul/linkgravity)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T11:28:38Z

## Translation

タイトル: Show HN: LinkGravity – AI コーディング エージェントのためのチャットと音声ブリッジ
記事のタイトル: GitHub - dev-sseul/linkgravity: Antigravity AI エージェントのメッセンジャー インターフェイス。インタラクティブな CLI ツールの承認と音声制御を備えています · GitHub
説明: Antigravity AI エージェントのメッセンジャー インターフェイス。インタラクティブな CLI ツール承認と音声制御機能 - dev-sseul/linkgravity

記事本文:
GitHub - dev-sseul/linkgravity: Antigravity AI エージェントのメッセンジャー インターフェイス。インタラクティブな CLI ツールの承認と音声制御を備えています · GitHub
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
開発スル
/
リンクグラビティ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
83 コミット 83 コミット .github/ workflows .github/ workflows bin bin docs/images docs/images フック フック npm-scripts npm-scripts src src voice-service voice-

サービス .gitignore .gitignore .npmignore .npmignore .npmrc .npmrc .pre-commit-config.yaml .pre-commit-config.yaml .prettierignore .prettierignore .prettierrc.json .prettierrc.json COTRIBUTING.md COTRIBUTING.md ライセンス ライセンスREADME.md README.md package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml 要件-dev.txt 要件-dev.txt 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Antigravity エージェント AI システム用の Discord、Telegram、および Slack ボット インターフェイス。 Antigravity CLI プロンプトをチャット UI コンポーネントに変換し、音声対話機能を提供します。
スレッドとしてのセッション: 各 agy 会話は独自のスレッド内に存在するため、複数の会話を並行して実行し、後で読み取れるようにすることができます。スレッドには、 Session-XXXX のままではなく、実際に話した内容に基づいて名前が付けられます。
承認フロー: コマンドおよびツール呼び出しの承認は、対話型のチャット ボタンになります。チェーンされたシェル コマンドは個別に承認され、承認の範囲を指定して、そのコマンドまたはツールを今後自動的に許可することができます。これは、単純な agy では行われません。
音声インタラクション: Discord 音声チャネルからエージェントに話しかけ、応答を聞きます。ウェイクワードを言って注意を引き、チャンネル内での雑談が気を悪くしないようにしましょう。
マルチモーダル入力: AI が読み取るファイル (音声を含む) を添付すると、自動的にテキストに変換されます。
端末と同じエージェント: マシン上にすでにある agy セットアップを読み取るため、チャットから開始されたセッションはローカルで使用しているものと同じモデルと設定を使用します。
不和
電報
たるみ
セッション
スレッド
フラット
スレッド
声
✓
✗
✗
承認ボタン
✓
✓
✓
添付ファイル
✓
✓
✓
セッションタイトルの自動名前変更
✓
✗
✓
DM
✓
✓
✓
グループ
✓
✗
✓
要件
反重力 CLI がインストールされている

このマシンで
メッセンジャー ボット トークン、およびそれを許可する少なくとも 1 つのサーバー/チャネル - Discord、Telegram、Slack がすべてサポートされており、一度に複数を有効にすることができます
Discord 開発者ポータルの場合:
New Application 、名前を付けて、左側のサイドバーの Bot に移動します。
[Privileged Gateway Intents] で、 [Message Content Intent] を有効にします。ボットはメッセージ テキスト/添付ファイルを読み取るため、必須です。
[トークンのリセット] をクリックしてボット トークンを表示し、コピーします。これが lgy setup の discord_token に組み込まれます。
サイドバーで「OAuth2 > URL Generator」に移動します。 [スコープ] で、 [bot] と [applications.commands] をチェックします。以下に表示される [ボットのアクセス許可] ボックスで、次のチェックボックスをオンにします。
メッセージの送信、スレッドでのメッセージの送信、公開スレッドの作成
メッセージ履歴の読み取り、ファイルの添付、リンクの埋め込み、リアクションの追加
接続して話す (音声チャネルのサポート用)
そのページの下部にある生成された URL をコピーし、ブラウザで開き、ボットをサーバーに招待します。
Telegram で @BotFather にメッセージを送信し、 /newbot を送信し、プロンプトに従ってボット トークンを取得します。これ以上の権限設定は必要ありません。Telegram セッションは 1 チャット = 1 セッションであるため、ボットに直接メッセージを送信して (またはグループに追加して) /new を実行するだけです。
Slack には他のものよりも多くの可動部分があります。2 つの個別のトークンと、相互にゲートするいくつかの設定ページです。この順序で進めると、手順をやり直す必要がなくなります。
api.slack.com/apps > Create New App > FromScratch に移動し、名前を付けて、ワークスペースを選択します。
ソケット モード (左側のサイドバー) > オンに切り替えます。これにより、パブリック HTTP エンドポイントが必要なくなります。 Slack では、ここでアプリレベルのトークンを生成するよう求められます。任意の名前を付け、connections:write スコープを追加して、 Generate します。このトークン ( xapp- で始まる) をコピーします。これは、lack_app_token です。
プロンプトが表示されない場合は、[基本情報] > [アプリレベル トークン] > [トークンの生成] に移動し、

代わりにスコープ。
[OAuth と権限] (左側のサイドバー) > [スコープ] > [ボット トークン スコープ] までスクロールします (ユーザー トークン スコープではありません。これはページのさらに上にある別のセクションであり、別のトークン用であり、ここでは使用されません)。 chat:write 、channels:history 、groups:history 、im:history 、mpim:history 、reactions:write 、files:read 、files:write のそれぞれに OAuth スコープを追加します。
同じページの一番上までスクロールし、 [ワークスペースにインストール] > [許可] を選択します。これにより、 [OAuth Tokens] > [Bot User OAuth Token] の下に xoxb- で始まるトークンが生成されます。それをコピーします。これは、slack_bot_token です。
ここで間違ったトークンを取得するのは簡単です。ページのさらに下にはユーザー OAuth トークン ( xoxp-... ) も表示されますが、これは別のものであり、このボットでは機能しません。
アプリのホーム (左側のサイドバー) > [タブの表示] で [メッセージ] タブをオンにし、[ユーザーがメッセージ タブからスラッシュ コマンドとメッセージを送信できるようにする] をオンにします。これにより、ボットに DM を送信できるようになります。 (このセクションがグレー表示になっている場合は、ステップ 3 がまだ保存/インストールされていないためです。戻って最初にそれを行ってください。)
イベントのサブスクリプション (左側のサイドバー) > [イベントの有効化] をオンに切り替え、 [ボット イベントのサブスクライブ] で、 message.channels 、 message.groups 、 message.im 、および message.mpim を追加し、 [変更を保存] を選択します。
スラッシュ コマンド (左側のサイドバー) > [新しいコマンドの作成] を /new 、 /model 、 /credit 、および /permissions に対して 4 回実行します (説明/ヒント テキストは何でも構いません。コマンド名のみが重要です)。
OAuth と権限 に戻ります。最初のインストール後にスコープ/イベントが変更されたため、「ワークスペースに再インストール」をクリックして、それらの変更をライブにプッシュします。後でスコープやイベントを変更する場合は、この手順を繰り返す必要があります。
Slack 自体で、ボットを使用できるようにしたいチャネル (DM ではない) に対して、最初に /invite @<your bot's name> を実行します。ボットは追加されていないチャネルに投稿できません。
npmインストール

-g リンク重力
独自の Python 環境を自動的にセットアップします。手動での pip インストールは必要ありません。
構成ウィザードを 1 回実行して、有効にするプラットフォームを選択し、トークン、許可されたユーザー、その他の設定を設定します。
lgyセットアップ
これは、パッケージ ディレクトリの外にある ~/.gemini/linkgravity/lgy.json に書き込むため、 npm update /reinstall は決して触れません。後でいつでも lgy セットアップを再実行して設定を変更できます。各フィールドを空のままにしても、現在の値が保持されます。 Discord、Telegram、Slack はすべて個別にオンにすることができます。ボットは、単一の共有プロセスで有効になっているものを実行します。
Discord の場合、セットアップ中に、許可する 1 つ以上のサーバーと、オプションで各サーバー内の特定のチャネルを尋ねられます。
サーバーのチャネル リストを空のままにすると、サーバー全体が許可され、どのチャネルでもセッションを開始できます。
サーバーの特定のチャンネル ID をリストする → そのサーバー内のチャンネルのみが許可されます。
Telegram と Slack にはまだサーバー/チャネル ゲーティングがありません。セットアップ中に設定した許可ユーザー リストに従って、ボットにメッセージを送信するすべてのチャット/チャネルでセッションを開始できます。
新しいセッションは /new スラッシュ コマンドで開始されます。単にメッセージを入力するだけでは開始されません。 Discord と Slack では、/new は通常のチャネルと既存のスレッドの両方で機能します。 Telegram では、どのチャットに送信しても適用されます。
lgy start # PM2 経由でボットをバックグラウンド デーモンとして起動します
lgy stop # やめてください
lgy restart #再起動します
lgy logs # ライブ ログを表示 - 行の後に -f を追加し、さらに行を追加するには --tail N
lgy enable # システム起動時に自動起動するようにボットを登録します
lgy disable # システムブートから削除します
lgy # インタラクティブ メニュー - 同じコマンドをリストから選択
開発
npm私
それだけです - Python 用の Ruff と Node 用の Prettier を接続し、さらにコミット時に lint/format する git フックを接続します。

従来のコミットのコミット メッセージを強制します。
LOG_LEVEL=DEBUG lgy 開始
タイムスタンプを含めるには lgy logs -t を使用します。
このボットは、AI エージェントに、それが実行されているマシンへの広範なアクセスを許可します。これはボットの動作に固有のものであるため、ボットを公に公開したり、ユーザーを完全に信頼していない場所で実行したりしないでください。
lgy setup によって設定される各プラットフォームの許可ユーザー リストは、主要なアクセス制御であり、常に設定します。
シェル コマンドを含むツール呼び出しは、デフォルトで承認フローを通過します。許可ユーザー リストに登録されているユーザーは誰でも、このマシンを実質的に完全に制御できるものとして扱われます。
ボット トークンとその他の設定は、この repo/package ディレクトリの外の ~/.gemini/linkgravity/lgy.json に存在します。そのファイルをコミットしたり共有したりしないでください。
Antigravity AI エージェントのメッセンジャー インターフェイス。インタラクティブな CLI ツールの承認と音声制御を備えています
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A messenger interface for the Antigravity AI agent. Features interactive CLI tool approvals and voice control - dev-sseul/linkgravity

GitHub - dev-sseul/linkgravity: A messenger interface for the Antigravity AI agent. Features interactive CLI tool approvals and voice control · GitHub
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
dev-sseul
/
linkgravity
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
83 Commits 83 Commits .github/ workflows .github/ workflows bin bin docs/ images docs/ images hooks hooks npm-scripts npm-scripts src src voice-service voice-service .gitignore .gitignore .npmignore .npmignore .npmrc .npmrc .pre-commit-config.yaml .pre-commit-config.yaml .prettierignore .prettierignore .prettierrc.json .prettierrc.json CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json pyproject.toml pyproject.toml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt View all files Repository files navigation
A Discord, Telegram, and Slack bot interface for the Antigravity agentic AI system. It translates Antigravity CLI prompts into chat UI components and provides voice interaction capabilities.
Sessions as Threads: Each agy conversation lives in its own thread, so several can run side by side and stay readable later. Threads get named from what you actually talked about instead of staying Session-XXXX .
Approval Flow: Command and tool-call approvals become interactive chat buttons. Chained shell commands are approved individually, and any approval can be scoped to auto-allow that command or tool going forward - something plain agy doesn't do.
Voice Interaction: Talk to the agent from a Discord voice channel and hear its replies. Say your wake word to get its attention, so side conversation in the channel doesn't set it off.
Multi-Modal Input: Attach files for the AI to read, including audio, which gets transcribed to text automatically.
Same Agent as Your Terminal: Reads the agy setup already on the machine, so sessions started from chat use the same models and settings you use locally.
Discord
Telegram
Slack
Sessions
Threads
Flat
Threads
Voice
✓
✗
✗
Approval buttons
✓
✓
✓
File attachments
✓
✓
✓
Session title auto-rename
✓
✗
✓
DMs
✓
✓
✓
Group
✓
✗
✓
Requirements
Antigravity CLI installed on this machine
A messenger bot token, and at least one server/channel to allow it in - Discord, Telegram, and Slack are all supported, and you can enable more than one at once
In the Discord Developer Portal :
New Application , name it, then go to Bot in the left sidebar.
Under Privileged Gateway Intents , enable Message Content Intent - required, since the bot reads message text/attachments.
Click Reset Token to reveal the bot token, and copy it - this is what goes into lgy setup 's discord_token .
Go to OAuth2 > URL Generator in the sidebar. Under Scopes , check bot and applications.commands . Under the Bot Permissions box that appears below, check:
Send Messages, Send Messages in Threads, Create Public Threads
Read Message History, Attach Files, Embed Links, Add Reactions
Connect, Speak (for voice channel support)
Copy the Generated URL at the bottom of that page, open it in a browser, and invite the bot to your server.
Message @BotFather on Telegram, send /newbot , and follow the prompts to get a bot token. No further permission setup is needed - Telegram sessions are 1 chat = 1 session, so just message your bot directly (or add it to a group) and run /new .
Slack has more moving parts than the others - two separate tokens, and a few settings pages that gate each other. Going in this order avoids re-doing steps:
Go to api.slack.com/apps > Create New App > From scratch , name it, and pick your workspace.
Socket Mode (left sidebar) > toggle it on. This avoids needing a public HTTP endpoint. Slack will prompt you to generate an app-level token here - name it anything, add the connections:write scope, and Generate . Copy this token (starts with xapp- ) - this is slack_app_token .
If it doesn't prompt you, go to Basic Information > App-Level Tokens > Generate Token and Scopes instead.
OAuth & Permissions (left sidebar) > scroll to Scopes > Bot Token Scopes (not User Token Scopes - that's a different section further up the page, for a different token, and is not used here). Add an OAuth Scope for each of: chat:write , channels:history , groups:history , im:history , mpim:history , reactions:write , files:read , files:write .
Scroll to the top of that same page > Install to Workspace > Allow . This generates the token under OAuth Tokens > Bot User OAuth Token , starting with xoxb- . Copy that one - this is slack_bot_token .
It's easy to grab the wrong token here - the page also shows a User OAuth Token ( xoxp-... ) further down, which is a different thing and won't work for this bot.
App Home (left sidebar) > under Show Tabs , turn on Messages Tab , then check Allow users to send Slash commands and messages from the messages tab - this is what lets you DM the bot at all. (If this section looks greyed out, it's because step 3 hasn't been saved/installed yet - go back and do that first.)
Event Subscriptions (left sidebar) > toggle Enable Events on > under Subscribe to bot events , add message.channels , message.groups , message.im , and message.mpim > Save Changes .
Slash Commands (left sidebar) > Create New Command , four times, for /new , /model , /credit , and /permissions (any description/hint text is fine - only the command name matters).
Back on OAuth & Permissions , since scopes/events changed after the initial install, click Reinstall to Workspace to push those changes live. Any time you change scopes or events later, you'll need to repeat this step.
In Slack itself, for any channel (not DM) you want the bot usable in, run /invite @<your bot's name> there first - the bot can't post in a channel it hasn't been added to.
npm install -g linkgravity
Sets up its own Python environment automatically - no manual pip install needed.
Run the configuration wizard once to pick which platform(s) to enable and set their tokens, allowed users, and other settings:
lgy setup
This writes to ~/.gemini/linkgravity/lgy.json , outside the package directory, so npm update /reinstall never touches it. You can re-run lgy setup any time to change settings later - each field keeps its current value if you leave it empty. Discord, Telegram, and Slack can all be turned on independently; the bot runs whichever ones are enabled in a single shared process.
For Discord, during setup you'll be asked for one or more servers to allow, and optionally specific channels within each:
Leave the channel list empty for a server → the whole server is allowed - any channel can start a session.
List specific channel IDs for a server → only those channels in that server are allowed.
Telegram and Slack have no server/channel gating yet - every chat/channel you message the bot from can start a session, subject to the allowed-users list you set during setup.
A new session is started with the /new slash command - never just by typing a message. On Discord and Slack, /new works both in a regular channel and from inside an existing thread; on Telegram, it applies to whichever chat you send it in.
lgy start # Start the bot as a background daemon via PM2
lgy stop # Stop it
lgy restart # Restart it
lgy logs # View live logs - add -f to follow, --tail N for more lines
lgy enable # Register the bot to auto-start on system boot
lgy disable # Remove it from system boot
lgy # Interactive menu - same commands, picked from a list
Development
npm i
That's it - it wires up Ruff for Python and Prettier for Node, plus git hooks that lint/format on commit and enforce Conventional Commits commit messages.
LOG_LEVEL=DEBUG lgy start
Use lgy logs -t to include timestamps.
This bot gives an AI agent broad access to the machine it runs on - that's inherent to what it does, so don't expose it publicly or run it somewhere you don't fully trust its users.
The allowed-users list for each platform, set via lgy setup , is your primary access control - always set it.
Tool calls, including shell commands, go through an approval flow by default; treat anyone on an allowed-users list as having effectively full control of this machine.
Your bot tokens and other settings live in ~/.gemini/linkgravity/lgy.json , outside this repo/package directory - never commit or share that file.
A messenger interface for the Antigravity AI agent. Features interactive CLI tool approvals and voice control
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
