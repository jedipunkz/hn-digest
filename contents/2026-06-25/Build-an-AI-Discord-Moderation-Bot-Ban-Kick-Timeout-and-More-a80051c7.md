---
source: "https://quickchat.ai/post/ai-discord-moderation-bot"
hn_url: "https://news.ycombinator.com/item?id=48673726"
title: "Build an AI Discord Moderation Bot: Ban, Kick, Timeout and More"
article_title: "Build an AI Discord Moderation Bot: Ban, Kick, Timeout (No Code) | Quickchat AI - AI Agents"
author: "piotrgrudzien"
captured_at: "2026-06-25T14:57:59Z"
capture_tool: "hn-digest"
hn_id: 48673726
score: 2
comments: 0
posted_at: "2026-06-25T14:12:32Z"
tags:
  - hacker-news
  - translated
---

# Build an AI Discord Moderation Bot: Ban, Kick, Timeout and More

- HN: [48673726](https://news.ycombinator.com/item?id=48673726)
- Source: [quickchat.ai](https://quickchat.ai/post/ai-discord-moderation-bot)
- Score: 2
- Comments: 0
- Posted: 2026-06-25T14:12:32Z

## Translation

タイトル: AI Discord モデレーション ボットを構築する: 禁止、キック、タイムアウトなど
記事のタイトル: AI Discord モデレーション ボットの構築: 禁止、キック、タイムアウト (コードなし) |クイックチャット AI - AI エージェント
説明: AI Discord モデレーション ボットを構築します。コードは必要ありません。 AI エージェントに、タイムアウト、キック、禁止、禁止解除、役割の割り当て、スローモードの設定、管理者専用のセーフティ ゲート、正確な設定、およびコピーするプロンプトを使用した平易な言語のモデレーター コマンドからのアナウンスの投稿を行う権限を与えます。

記事本文:
AI Discord モデレーション ボットを構築する: 禁止、キック、タイムアウト (コードなし) |クイックチャット AI - AI エージェント
ロゴをクリップボードにコピー (png)
ロゴをクリップボードにコピー (svg)
ブランド資産のダウンロード ソリューションの使用例
カスタマーサポート チケットを解決し、応答時間を短縮します
販売と見込み客の発掘 見込み客を特定し、製品を推奨する
Eコマース カートを回収し、買い物客を誘導する
IT および社内ヘルプデスク 従業員および社内の質問に回答します
大規模チーム向けのエンタープライズ セキュリティ、拡張性、制御
プラットフォーム AI エージェント プラットフォーム リリース Quickchat AI の新機能 お問い合わせ なぜ当社について 価格設定 リソース 私たちについて ケーススタディ 採用情報 ブログ アフィリエイト プログラム 法的文書のステータス お問い合わせ English Deutsch Español Français 日本語 Polski Português Svenska ログイン 無料で始める 無料で始める ソリューション ユースケース
カスタマーサポート セールスとリードジェネレーション eコマース ITと社内ヘルプデスク エンタープライズチャネルと統合
ウェブサイト チャット WordPress WhatsApp Discord Shopify Intercom Zendesk HubSpot すべてのチャネルと統合 プラットフォーム リリース お問い合わせ なぜ当社を選ぶのか 価格設定 リソース 当社について ケーススタディ 採用情報 ブログ アフィリエイト プログラム 法的文書 ステータス お問い合わせ ログイン 無料で始める
{t.navigation.log_in}
{t.navigation.book_a_demo}
--> ブログ
AI Discord モデレーション ボットを構築する: 禁止、キック、タイムアウト (コードなし)
AI エージェントはすでに Discord 上のコミュニティと会話しています。このガイドでは、AI Discord モデレーション ボットに変換します。コードは必要ありません。いくつかのカスタム AI アクションを使用して、モデレータが「@スパマーを 10 分間タイムアウト」、「詐欺リンクの @user を禁止」、または「#general のスローモードを 30 秒に設定」を入力すると、エージェントが Discord API を介して平易な言語でアクションを実行します。
私たちはスラッシュ コマンド ボットではなく、会話型エージェントを作成します。したがって、ここでの目標は古典的なルールベースを再構築することではありません

MEE6 や Dyno のようなモデレーション ボット (Quickchat との比較)。これは、モデレーターが平易な言葉でモデレーションを推進できるようにするもので、エージェントは意図を適切な Discord API 呼び出しに変換し、破壊的な行為が行われる前に確認を求め、毎回理由を監査ログに書き込みます。
必要なものは 2 つありますが、どちらも無料です。
Discord サーバーに接続された Quickchat AI エージェント (ここでサインアップして無料で使用できます)
あなたが管理者である Discord サーバー
このメカニズムは AI アクション : エージェントが会話中に行うカスタム HTTP リクエストです。ワンクリックの Discord モデレーション パックはありません。これらのアクションは手動で追加します。これは、Discord を指す Telegram Bot API または Google Sheets にエージェントを接続するのと同じアプローチです。最終的には 7 つの動作するアクションが得られ、そのうちの破壊的なアクションは決定論的ゲート (Discord で検証されたフラグの実行条件。プロンプトがボットにそれを通過させることはできないようにする) によって管理者にロックされており、それぞれを自分でテストすることができます。
これは長くて正確なウォークスルーです。 AI アクションの正規リファレンスは、 docs.quickchat.ai/ai-agent/actions のドキュメントにあります。
以下のスクリーンショットは、架空のコミュニティ サーバー Nebula Lounge のモデレーション副パイロットである Orbit と呼ばれるテスト エージェントからのものです。サーバーはサンプルが中立的なものであるように作成されていますが、ここで示されているすべての会話、すべての API 呼び出し、およびすべての効果は、実際の Discord サーバーに対して実際の応答パイプラインを実行している実際のエージェントによって生成されたものです。手順を進める際には、自分のサーバーの詳細を使用してください。
7 つのアクションが 3 つのジョブにグループ化されています。そのうちの 2 つ (キックと禁止) は破壊的であるため、エージェントは実行する前に確認します。
各アクションは、エージェントが会話中に呼び出すことができる、記述された HTTP リクエストです。
チャット メッセージは誰からでも送信される可能性があるため、当然の心配は、

通常のメンバーはボットに話しかけて誰かを禁止することができます。それはできません。破壊的なアクションは、Discord で検証されたフラグの実行条件によって管理者にロックされ、プロンプトによってではなく、リクエストが送信される前に私たちの側で評価されます。ステップ 7 でそのゲートを配線します。
同じ構成要素が、司会者の副操縦士を超えて到達します。もう 1 つ追加すると、コミュニティ全体がエージェントと会話できるようになります。パスフレーズの背後にある自己主張のような、正しいことを言うと報酬が得られる小さなゲームやチャレンジです。 「さらに進む」で、最後に安全なバージョンを構築します。
AI アクションが Discord を呼び出す方法とサーバーの値がどのように流入するか
AI アクションは、記述された HTTP リクエストです。エージェントがアクションが適用されると判断すると、パラメータを入力し、クイックチャットがリクエストを送信します。 3 つの情報が Discord に到達する必要があり、それらは 3 つの異なる方法で到達します。
トークンは接続されているボットから取得され、サーバー ID はライブ会話から取得され、ターゲットはモデレーターのメッセージから取得されます。それぞれに一度配線します。
1. ボット トークンはシステム トークンから取得されます。 Discord は、ボット トークンを使用してすべてのリクエストを認証します。各アクションに貼り付ける必要はありません。ボットが Discord 統合に接続されると、トークンはシステム トークン {{discord_bot_token}} として利用可能になります。 [AI データの追加] メニューからこれを選択し、Authorization ヘッダーに置きます。これは送信時に発信リクエストに挿入され、モデルに公開されることも、会話に書き込まれることも、API によって返されることもありません。これは、アクションの秘密保管庫に相当する安全な場所です。
[システム トークン] の下の [AI データの追加] メニューから [{{discord_bot_token}}] を選択します。これは送信時にリクエストに挿入され、モデルには表示されません。
2. サーバー ID は会話メタデータから取得されます。あなたの年齢のとき

Discord サーバー内で会話している場合、Quickchat は会話がどのサーバー (ギルド) で行われているかをすでに知っています。その値は会話のメタデータとして使用され、それを {{metadata_guild_id}} として参照します。サーバー ID をハードコーディングすることはありません。アクションはライブ会話からそれを読み取ります。これは、訪問者のページ URL または言語をアクションに組み込むメカニズムと同じであり、同じ [AI データの追加] メニューから選択します。
3. ターゲットはパラメータとしてモデレーターから提供されます。誰をタイムアウトするか、どのチャネルを遅くするか、どの役割を付与するか: これらは毎回変わるため、エージェントが会話から入力するパラメータになります。モデレーターは通常の方法でメンバーまたはチャンネルにタグを付け、@username または #general と入力し、Discord のオートコンプリートから選択します。 Discord は、そのタグをラップされた数値 ID (メンバーの場合は <@123456789>、チャネルの場合は <#123456789>) としてボットに配信するため、パラメーターの説明では、Discord API が必要とする数字のみを使用するようにエージェントに指示します。モデレーターは生の ID を手で入力することはありません。
正しく理解するための詳細が 1 つあります。認証スキームは Bearer ではなく Bot です。これは、Discord を手動で呼び出すときによくある間違いです。ヘッダー値は文字通り Bot {{discord_bot_token}} です。
始める前に: ボットに接続し、適切な権限を付与します。
エージェントを Discord に接続する方法については、Discord ドキュメントと AI Discord ボットの作成ガイドで段階的に説明されているため、ここでは繰り返しません。まずそれを行ってから戻ってください。これらのガイドでは、モデレートではなくチャットに関するものであるため、取り上げていないことが 1 つあります。それは、アクセス許可です。
チャットボットはメッセージを読んで送信するだけで済みます。モデレーション ボットは、メンバーのキック、禁止、タイムアウト、ロールの管理、チャネルの管理を行う必要があります。 Discord は、ボットの役割が実際に許可を保持するまで、403 によるすべてのモデレーション コールを拒否します。

ns。それらを含む許可整数を使用してボットを再招待することで、それらを許可します。
https://discord.com/api/oauth2/authorize?client_id=YOUR_APPLICATION_ID&permissions=1409017777174&scope=bot
この整数は、チャネルの表示、メッセージの送信、メッセージ履歴の読み取り、スレッドの作成と送信、さらにメンバーのキック、メンバーの禁止、メンバーの管理 (タイムアウト)、ロールの管理、およびチャネルの管理の 7 つのアクションに対する最小特権です。 YOUR_APPLICATION_ID を、Discord Developer Portal のボットのページのアプリケーション ID に置き換えます。 [サーバー設定]、[役割] で同じボックスを手動でチェックすることもできます。
2 番目の、より微妙なルールがあります: 役割階層 。ボットは、自身の最上位のロールの下にあるメンバーおよびロールに対してのみ動作できます。ボットのロールがリストの一番下にある場合、適切な権限があっても、誰もタイムアウトしたり、ロールを割り当てたりすることはできません。招待したら、[サーバー設定]、[ロール] を開き、ボットのロールを上部近くにドラッグします。
ボットのロールは、ボットが管理するロールよりも上位にあるため、Discord ではボットがこれらのメンバーに作用できるようにします。
アクションが 403 を返した場合、ほとんどの場合、権限が欠落しているか、ボットの役割が低すぎるかのいずれかです。エージェントは、黙って再試行するのではなく、どちらかを言うように指示されます。
ステップ 1: エージェントを作成し、それにジョブを与える
エージェント (ここでは Orbit ) を作成し、サーバーに接続します。 Orbit には知識ベースがほとんど必要ありません。その仕事は指示に基づいて行動することであり、文書からの質問に答えることではありません。必要なのは、モデレーターとしてどのように行動するかを説明する明確なアイデンティティです。後のステップで完全なプロンプトを貼り付けます。ここでは、エージェントを作成してボットを接続します。
ステップ 2: ボットにコマンドを実行できる人を制限する
これは最も重要な安全上の決定であるため、アクションを構築する前に決定します。エージェントは誰が誰であるかを知ることができません

それに話しかけています。メッセージの文言には、送信者がモデレーターであることを証明するものは何もないため、プロンプトが誰を禁止するかを決定するものになることは決してありません。モデルをまったく実行しないチェックが必要です。 Quickchat には 2 つの機能があり、両方を使用します。
検証されたフラグ (実際の境界) 上の決定論的なゲート。すべての Discord メッセージで、Quickchat は送信者がサーバー管理者であるかどうかを会話メタデータ author_is_admin として記録します。 Discord は送信者の権限に基づいて設定します。誰もそれを入力したり、存在を主張したりすることはできません。ステップ 7 では、各破壊アクションに実行条件を追加します。つまり、author_is_admin が true の場合にのみ実行されます。リクエストが Discord に届く前に、通話時にこちら側でチェックされるため、会話の内容に関係なく、モデレーター以外のユーザーは拒否されます。
モデレータ専用チャンネル (多層防御)。モデレーターの役割のみが表示できるプライベート #mod-commands チャネルを作成し、ボットがそこでの読み取りと返信を制限します ( @everyone のチャネルの表示を拒否し、ボットと MOD に対しては許可します)。ボットに到達できる人さえ少なくなり、到達すべきでない者は上のゲートによって阻止されます。
また、キックと禁止を最初に確認するプロンプト ルール (プロンプト ブロック内) も追加します。そのため、モデレーターであっても、単一の曖昧な行で取り消し不能なアクションを開始することはありません。門は境界です。確認はシートベルトです。
3 つのレイヤー、1 つの実際の境界。モデレータ専用チャネルはボットにアクセスできる人を制限し、プロンプトによってボットの応答方法が決まりますが、実際に非管理者を停止させるのは Discord で検証された author_is_admin フラグの実行条件だけです。これは、Discord の呼び出し前にこちら側でチェックされており、議論することができないためです。
各レシピの読み方: アクションの構造
以下のすべてのレシピは、アクション エディターの同じ 6 つのフィールドであり、値が異なります。

ここでレイアウトを一度学習すれば、残りは穴埋めです。これは、6 つのフィールドすべてが入力された assign_role (ロール付与アクション) です。
アクション全体: 名前、パラメーター、ヘッダーと本文を含むエンドポイント、および説明。フィールドの横にある [AI データの追加] ボタンを使用して、色付きのチップを挿入します。
エンドポイント URL を詳しく見てみましょう。チップは色分けされています。オレンジ色は自動的に注入され、紫色はユーザーが定義してエージェントが塗りつぶします。
したがって、以下のすべてのアクションについて、 [アクションと MCP] を開き、HTTP リクエスト アクションを作成し、表示されている API アクション名、ユーザーに最初に尋ねること、メソッドと URL、ヘッダー、本文、および API アクションの説明を正確に入力します。各レシピはまさに 1 つのアクションです。それらを一度に 1 つずつ構築します。
ステップ 3: 最初のアクション、お知らせを投稿する
特別な許可を必要としない 1 つのアクションから開始します (すべてのボットはすでにメッセージを送信できます)。そのため、破壊的なものに触れる前に緑色の結果が表示されます。アプリで、[アクションと MCP] に移動し、HTTP リクエスト アクションを作成します。以下の各太字のラベルは、エディター内の正確なフィールド名であり、画面上で一致する順序でリストされているため、各値を一致するフィールドに直接貼り付けてください。
API アクション名: send_payment
ユーザーに最初に尋ねること (パラグラフ)

[切り捨てられた]

## Original Extract

Build an AI Discord moderation bot, no code required. Give your AI Agent the power to timeout, kick, ban, unban, assign roles, set slowmode, and post announcements from plain-language moderator commands, with an admin-only safety gate, the exact settings, and the prompt to copy.

Build an AI Discord Moderation Bot: Ban, Kick, Timeout (No Code) | Quickchat AI - AI Agents
Copy logo to clipboard (png)
Copy logo to clipboard (svg)
Download brand assets Solutions Use Cases
Customer Support Resolve tickets and cut response times
Sales & Lead Generation Qualify leads and recommend products
Ecommerce Recover carts and guide shoppers
IT & Internal Helpdesk Answer employee and internal questions
Enterprise Security, scale, and control for large teams
Platform The AI agent platform Releases What's new in Quickchat AI Contact us Why Us Pricing Resources About us Case Studies Careers Blog Affiliate Program Legal Documentation Status Contact English Deutsch Español Français 日本語 Polski Português Svenska Log in Start for free Start for free Solutions Use Cases
Customer Support Sales & Lead Generation Ecommerce IT & Internal Helpdesk Enterprise Channels & Integrations
Website chat WordPress WhatsApp Discord Shopify Intercom Zendesk HubSpot All channels & integrations Platform Releases Contact us Why Us Pricing Resources About us Case Studies Careers Blog Affiliate Program Legal Documentation Status Contact Log in Start for free
{t.navigation.log_in}
{t.navigation.book_a_demo}
--> Blog
Build an AI Discord Moderation Bot: Ban, Kick, Timeout (No Code)
Your AI Agent already talks to your community on Discord. This guide turns it into an AI Discord moderation bot , no code required. With a handful of custom AI Actions , a moderator types “timeout @spammer for 10 minutes” , “ban @user for scam links” , or “set slowmode in #general to 30 seconds” , and the Agent performs the action through the Discord API, from plain language.
We make conversational Agents, not slash-command bots. So the goal here is not to rebuild a classic rule-based moderation bot like MEE6 or Dyno ( how Quickchat compares ). It is to let a moderator drive moderation in plain language , with the Agent translating intent into the right Discord API call, asking for confirmation before anything destructive, and writing a reason to the audit log every time.
You need two things, both free:
a Quickchat AI Agent ( sign up here and use for free ), connected to your Discord server
a Discord server where you are an admin
The mechanism is AI Actions : custom HTTP requests your Agent makes during a conversation. There is no one-click Discord moderation pack; you add these actions by hand. It is the same approach as wiring an Agent to the Telegram Bot API or to Google Sheets , pointed at Discord. By the end you will have seven working actions , with the destructive ones locked to admins by a deterministic gate (a run-condition on a Discord-verified flag, so a prompt can never talk the bot past it), and you will have tested each one yourself .
This is a long, exact walkthrough. The canonical reference for AI Actions lives in the docs at docs.quickchat.ai/ai-agent/actions .
The screenshots below come from a test Agent called Orbit , the moderation co-pilot for a fictional community server, Nebula Lounge . The server is invented so the example stays neutral, but every conversation, every API call, and every effect shown here was produced by a real Agent running the real reply pipeline against a real Discord server. Use your own server’s details when you follow along.
Seven actions, grouped into three jobs. Two of them (kick and ban) are destructive, so the Agent will confirm before running them.
Each action is a described HTTP request the Agent can call during a conversation.
A chat message can come from anyone, so the natural worry is whether a regular member could talk the bot into banning someone. They cannot. The destructive actions are locked to admins by a run-condition on a Discord-verified flag, evaluated on our side before any request goes out, never by the prompt. You wire that gate in Step 7 .
The same building blocks reach past a moderator co-pilot. With one more piece you can let the whole community talk to the Agent: little games and challenges where saying the right thing earns a reward, like a self-claim role behind a passphrase. We build a safe version at the end, in Going further .
How AI Actions call Discord, and how your server’s values flow in
An AI Action is a described HTTP request. When the Agent decides an action applies, it fills in the parameters and Quickchat sends the request. Three pieces of information have to reach Discord, and they arrive in three different ways .
The token comes from the connected bot, the server ID from the live conversation, and the target from the moderator’s message. You wire each one once.
1. The bot token comes from a System Token. Discord authenticates every request with your bot token. You do not paste it into each action. Once your bot is connected in the Discord integration, the token is available as the System Token {{discord_bot_token}} . Select it from the Add AI Data menu and put it in the Authorization header. It is injected into the outgoing request at send time, and it is never exposed to the model, never written into a conversation, and never returned by an API. It is the safe equivalent of a secrets vault for your actions.
Select {{discord_bot_token}} from the Add AI Data menu, under System Tokens. It is injected into the request at send time and never shown to the model.
2. The server ID comes from conversation metadata. When your Agent is talking inside a Discord server, Quickchat already knows which server (guild) the conversation is in. That value rides along as conversation metadata, and you reference it as {{metadata_guild_id}} . You never hardcode your server ID; the action reads it from the live conversation. This is the same mechanism that carries a visitor’s page URL or language into an action, and you pick it from the same Add AI Data menu.
3. The target comes from the moderator, as a parameter. Who to timeout, which channel to slow down, which role to grant: these change every time, so they are parameters the Agent fills from the conversation. A moderator tags the member or channel the normal way, typing @username or #general and picking it from Discord’s autocomplete. Discord delivers that tag to the bot as the wrapped numeric ID ( <@123456789> for a member, <#123456789> for a channel), so the parameter description tells the Agent to use only the digits, which is what the Discord API needs. Moderators never type a raw ID by hand.
One detail to get right: the auth scheme is Bot , not Bearer , a common mistake when calling Discord by hand. The header value is literally Bot {{discord_bot_token}} .
Before you start: connect the bot and grant the right permissions
Connecting an Agent to Discord is covered step by step in our Discord docs and in the Create an AI Discord bot guide, so we will not repeat it here. Do that first, then come back. There is one thing those guides do not cover, because they are about chatting, not moderating: permissions .
A chat bot only needs to read and send messages. A moderation bot needs to kick, ban, time members out, manage roles, and manage channels. Discord will reject every moderation call with a 403 until the bot’s role actually holds those permissions. Grant them by re-inviting the bot with a permission integer that includes them:
https://discord.com/api/oauth2/authorize?client_id=YOUR_APPLICATION_ID&permissions=1409017777174&scope=bot
That integer is least-privilege for exactly these seven actions : View Channels, Send Messages, Read Message History, Create and Send in Threads, plus Kick Members, Ban Members, Moderate Members (timeout), Manage Roles, and Manage Channels . Replace YOUR_APPLICATION_ID with the Application ID from your bot’s page in the Discord Developer Portal . You can also tick the same boxes by hand in Server Settings, Roles.
There is a second, subtler rule: role hierarchy . A bot can only act on members and roles that sit below its own highest role . If the bot’s role is at the bottom of the list, it cannot timeout anyone or assign any role, even with the right permissions. After inviting it, open Server Settings, Roles, and drag the bot’s role near the top.
The bot’s role sits above the roles it manages, so Discord lets it act on those members.
If an action ever returns 403 , it is almost always one of these two: a missing permission, or the bot’s role sitting too low. The Agent is told to say which one, rather than silently retrying.
Step 1: Create the Agent and give it its job
Create your Agent (here it is Orbit ) and connect it to your server. Orbit needs almost no knowledge base; its job is to act on instructions, not to answer questions from documents. What it does need is a clear Identity describing how to behave as a moderator. We will paste the full prompt in a later step ; for now, create the Agent and connect the bot.
Step 2: Restrict who can command the bot
This is the most important safety decision, so we settle it before building any action. The Agent cannot tell who is talking to it. Nothing in the words of a message proves the sender is a moderator, so the prompt can never be the thing that decides who may ban. You need a check that does not run through the model at all. Quickchat gives you two, and you use both:
A deterministic gate on a verified flag (the real boundary). On every Discord message, Quickchat records whether the sender is a server admin as the conversation metadata author_is_admin . Discord sets it from the sender’s permissions; nobody can type it or argue it into existence. In Step 7 you add a run-condition to each destructive action: run only when author_is_admin is true. It is checked on our side, at call time, before any request reaches Discord, so a non-moderator is refused no matter what the conversation says.
A moderators-only channel (defense in depth). Create a private #mod-commands channel that only your moderator role can see, and restrict the bot to read and reply there (deny View Channel for @everyone , allow it for the bot and your mods). Fewer people can even reach the bot, and the gate above stops anyone who does but should not.
We also add a prompt rule (in the prompt block ) that makes kick and ban confirm first , so even a moderator never fires an irreversible action on a single ambiguous line. The gate is the boundary; the confirmation is the seatbelt.
Three layers, one real boundary. The moderators-only channel limits who can reach the bot, and the prompt shapes how it replies, but only the run-condition on the Discord-verified author_is_admin flag actually stops a non-admin, because it is checked on our side, before any Discord call, and cannot be argued with.
How to read each recipe: the anatomy of an action
Every recipe below is the same six fields in the action editor, with different values. Learn the layout once here and the rest is fill-in-the-blanks. This is assign_role (the role-granting action) with all six fields filled in:
The whole action on one screen: a name, the parameters, the endpoint with its headers and body, and the description. You insert any colored chip with the Add AI Data button next to a field.
A closer look at the endpoint URL. The chips are color-coded: orange is injected for you, purple is what you define and the Agent fills.
So for every action below, open Actions & MCPs , create an HTTP Request action, and fill in exactly the API Action Name , What to ask the user first , method and URL , Headers , Body , and API Action Description shown. Each recipe is exactly one action; build them one at a time.
Step 3: Your first action, post an announcement
We start with the one action that needs no special permission (every bot can already send messages), so you see a green result before touching anything destructive. In the app, go to Actions & MCPs and create an HTTP Request action. Each bold label below is the exact field name in the editor, listed in the order you meet it on screen, so paste each value straight into the matching field.
API Action Name: send_announcement
What to ask the user first (the para

[truncated]
