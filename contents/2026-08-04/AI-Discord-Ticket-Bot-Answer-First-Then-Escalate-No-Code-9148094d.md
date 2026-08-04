---
source: "https://quickchat.ai/post/discord-ai-support-ticket-bot"
hn_url: "https://news.ycombinator.com/item?id=49170100"
title: "AI Discord Ticket Bot: Answer First, Then Escalate (No Code)"
article_title: "AI Discord Ticket Bot: Answer First, Then Escalate (No Code) | Quickchat AI - AI Agents"
author: "piotrgrudzien"
captured_at: "2026-08-04T16:08:52Z"
capture_tool: "hn-digest"
hn_id: 49170100
score: 1
comments: 0
posted_at: "2026-08-04T15:12:25Z"
tags:
  - hacker-news
  - translated
---

# AI Discord Ticket Bot: Answer First, Then Escalate (No Code)

- HN: [49170100](https://news.ycombinator.com/item?id=49170100)
- Source: [quickchat.ai](https://quickchat.ai/post/discord-ai-support-ticket-bot)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T15:12:25Z

## Translation

タイトル: AI Discord Ticket Bot: まず答えてからエスカレート (コードなし)
記事のタイトル: AI Discord Ticket Bot: まず答えてからエスカレートする (コードなし) |クイックチャット AI - AI エージェント
説明: コードなしの AI Discord チケット ボットを構築します。知識ベースから応答し、できない場合はプライベート チケット スレッドを開き、サポート ロールに ping を送り、人間に引き継ぎます。

記事本文:
AI Discord チケット ボット: 最初に応答し、その後エスカレートします (コードなし) |クイックチャット AI - AI エージェント
のために
ロゴをクリップボードにコピー (png)
ロゴをクリップボードにコピー (svg)
ブランド資産をダウンロードする
カスタマーサポート チケットを解決し、応答時間を短縮します
販売と見込み客の発掘 見込み客を特定し、製品を推奨する
Eコマース カートを回収し、買い物客を誘導する
IT および社内ヘルプデスク 従業員および社内の質問に回答します
大規模チーム向けのエンタープライズ セキュリティ、拡張性、制御
プラットフォーム AI エージェント プラットフォームのリリース Quickchat AI の新機能 お問い合わせ 会社情報
なぜ当社なのか 価格設定 リソース ログイン 無料で始める 無料で始める ソリューション ユースケース
カスタマーサポート セールスとリードジェネレーション eコマース ITと社内ヘルプデスク エンタープライズチャネルと統合
ウェブサイト チャット WordPress WhatsApp Discord Telegram Shopify インターコム HubSpot Zendesk その他の統合 プラットフォーム リリース お問い合わせ なぜ当社を選ぶのか 価格設定 リソース 会社
AI Discord チケット ボット: 最初に応答し、その後エスカレートします (コードなし)
;元の PNG/JPG はそのまま残ります。
フォールバックし、ソーシャル スクレーパーの og:image のままになります。 --> Quickchat AI は、Discord サーバー内のサポートの質問に独自のドキュメントから回答し、回答できない場合にのみプライベート チケット スレッドを開きます。このガイドでは、コードを使用せずにボットをエンドツーエンドで約 30 分で構築します。
その結果、優れた第一線のエージェントのように動作します。メンバーが API キーをリセットする方法を尋ね、チャネル内の実際の手順を取得します。メンバーがドキュメントに記載されていないことを質問すると、ボットは推測する代わりにそのように答え、プライベート スレッドを開き、メンバーをそのスレッドに引き込み、サポートの役割に 1 行の概要を ping し、受信トレイから応答できる人間に会話を渡します。
必要なものは 2 つありますが、どちらも無料です。
Quickchat AI エージェント (ここでサインアップして無料で使用してください)
あなたがいるDiscordサーバー

管理者
以下のスクリーンショットは、 Beacon という架空の製品分析会社のサポート ボットである Lumen というテスト エージェントからのものです。会社は発明されたものであるため、例は中立的なままですが、ここで示されているすべての会話、すべての API 呼び出し、およびすべてのチケット スレッドは、実際の Discord サーバーに対して実際の応答パイプラインを実行している実際のエージェントによって生成されました。
質問ごとに 1 つの決定。 Happy Path では、Discord には何も書き込まれません。
3 つの AI アクションがエスカレーション作業を実行し、Quickchat AI が 1 つのギャラリー カードから 3 つすべてをインストールします。
ナレッジベースとは何ですか?また、ナレッジベースがチケットの量を決定するのはなぜですか?
これまでに AI エージェントを構築したことがない場合、これは 2 分に値する 1 つの概念です。
ナレッジ ベースは独自のドキュメントであり、エージェントにロードされ、会話中に検索できるようにインデックスが付けられます。 Quickchat AI にヘルプ ドキュメントの URL を指定すると、ページをスクレイピングしたり、手動で記事を貼り付けたり、ファイルをアップロードしたりできます。応答時に、エージェントはそれを検索し、基盤となるモデルがたまたま世界について記憶しているものからではなく、見つけたものから応答します。何もない状態から始める場合、チャットボットのナレッジ ベースを構築する方法には、何をどのような順序で書くべきかが含まれます。
ドキュメントでカバーされているすべての答えが得られます。彼らが持たないものはすべてチケットになります。
それがこのビルドの全体的なレバーです。偏向はすぐにできるトリックではなく、モデルの能力でもありません。それはあなたの範囲です。ボットが予想以上にエスカレートした場合、チケットはどのページがまだ書いていないのかを示します。
測定されたものではなく、例示的なものです。重要なのは方向性です。ドキュメントを書くことが、ここでできる最も大きな力を発揮することです。
AI は実際に Discord でのサポートの質問に答えることができますか?
ドキュメントでカバーされている質問については、はい、言い換えではなく実際の手順が記載されています。ここ

は、テスト サーバーからの偏向ケースであり、編集されていません。
ナレッジベースの記事からチャネル内で回答されました。チケットは必要なかったため、作成されませんでした。
より興味深い半分は、答えが存在しない場合に何が起こるかです。ボットはこう言います:
正直な失敗。 Beacon のナレッジ ベースではエンタープライズ SSO が意図的に省略されているため、これがチケットを開くパスになります。
もっともらしい SSO 手順を考案するエージェントは、チケットをオープンするエージェントよりも悪いため、この動作は期待するのではなく、意図的にテストする価値があります。これについてはテストでもう一度説明します。
Discord に接続する 2 つの方法と、チケットに 2 番目の接続が必要な理由
Quickchat AI には 2 つの Discord 接続があり、開始する前に違いが重要になります。
どちらもサーバー内の質問に答えます。 AI アクションを実行できるのは 1 つだけです。
ワンクリックがファストパスです。外部アプリで Discord を開き、 Discord サーバーに追加 をクリックして承認すると、約 1 分で共有 Quickchat AI ボットがサーバー内で稼働するようになります。メンバーが @mention するか、 /ask を使用すると応答します。多くのサーバーでは、それが仕事全体です。
あなたはそれに名前を付けることができます。承認する前に、新しいサーバー フィールドにボット名があり、エージェントの名前が事前に入力されています。接続されているすべてのサーバーは、このサーバー ボックスに独自のボット名を保持しており、後で変更できます。これは Discord のニックネームであるため、サーバーごとに設定され、32 文字に制限されており、これをクリアするとボットは Quickchat AI に戻ります。 1 つの Discord アプリケーションがそれをインストールするすべてのサーバーにサービスを提供するため、共有されるのはアバターと基礎となる @username です。
ワンクリックは一般的なものではありません。ボットはエージェントの名前をサーバーに転送し、後でサーバーごとに変更できます。
詳細設定では、独自の Discord アプリケーションを登録します。 Discord の開発者ポータルでボットを作成し、そのトークンを貼り付けます。

Quickchat AI に一度接続します。所要時間は約 10 分で、自分のアバターと @username 、ダイレクト メッセージ、1 つのサーバー内の複数のボット、そしてこのガイドに必要なものである AI アクションを取得できます。
その理由は、単なる癖として扱うのではなく、明確に述べる価値があります。 AI アクションは、Quickchat AI が送信時にリクエストに挿入する認証情報を使用して、Discord API をボットとして呼び出します。共有ワンクリック ボットのトークンは、あなたではなく共有ボットに属しているため、アクションで送信することはできません。独自のアプリが接続されていない状態で Discord ギャラリーを開くと、実行できないものを構築するのではなく、その旨が通知されてカードが無効になります。
外部アプリは両方の道の始まりです。高度な設定は、Discord パネルの下部にある吹き出しにあります。
始める前に必要なもの
独自の Discord アプリを接続します。 [外部アプリ] で、Discord を開き、[独自の Discord アプリを使用する] を見つけて、[詳細設定] を選択します。 Discord の開発者ポータルでアプリケーションとボットを作成し、メッセージ コンテンツ インテントをオンにし、OAuth2 コード許可が必要をオフのままにして、ボット トークンを貼り付けます。 Quickchat AI は入力時にそれを検証します。これまでに作成したことがない場合は、「AI Discord ボットの作成」ガイドでポータル側について詳しく説明しています。
サポート ボットが必要とするものだけをボットに付与します。チケット チャネル上: チャネルの表示、メッセージの送信、メッセージ履歴の読み取り、プライベート スレッドの作成、スレッドでのメッセージの送信、およびスレッドの管理。キック、禁止、タイムアウトの権限は必要なく、テスト サーバーでは意図的に何も保持しません。サポート ボットにはモデレートする理由がありません。また、サポート ボットが保持していない権限については、特に考える必要はありません。
Discord で、アクションが指す 2 つのものを作成します。この部分は、Quickchat AI ではなく Discord クライアントで行われ、そこで行う唯一の設定です。
チャンネル

チケットが開きます。どのテキスト チャネルでも使用できます。チャンネル リストの横にある + を使用するか、手持ちの + を再利用します。テストサーバーでは #support です。
チケットが開いたときにページングする役割。 「サーバー設定」、「役割」、「役割の作成」の順に進みます。テストサーバーでは @support です。
3か所ありますが、公開されているのは1か所だけです。
これらは両方ともステップ 3 で戻ります。チケット テンプレートをインストールすると、Quickchat AI は Discord からサーバーをライブで読み取り、3 つのドロップダウンを表示します。そのうちの 1 つは、見つかったロールをリストするロール ピッカーです。どこかに ID を入力するのではなく、そこでサポートの役割を選択します。
ステップ 3 のロールピッカー。Discord からリストを読み取るため、作成したロールがすぐにここに表示されます。
スキップすると気になる点が 1 つあります。そのピッカーでは、Discord が報告するロールのみがmentionable としてリストされ、 @everyone が省略されます。サポートの役割がメンション可能に設定されていない場合、その理由は何も説明されず、単にドロップダウンに表示されません。テンプレートの仕事全体がそのロールに ping を送信することであるため、フィルターは意図的に行われます。ロールの作成中に設定すると、問題が発生することはありません。
「サーバー設定」、「役割」でサポート役割を選択し、「表示」の下でこれをオンにします。それがなければ、ロールはピッカーに届きません。
ボットはどのようにしてチケットに応答するか、チケットをオープンするかを決定しますか?
ほとんどの Discord チケットボットは何も決定しません。 Ticket Tool、TicketsBot などは、[Create Ticket] ボタンを備えた固定パネルを投稿するか、スラッシュ コマンドを登録します。メンバーがそれをクリックするとチケットが開き、チームがそれを読みます。これは優れた設計であり、完全に予測可能ですが、今週 API キーをリセットする方法を尋ねた 10 人全員が 10 枚のチケットを作成したことを意味します。
同じ10の質問です。違いは、チームが最終的に何を読むかです。
回答優先ボットは、決定する前にドキュメントを読み取ります。そして私が

エスカレートすることはなく、1 つのチケットを開くには正確に 4 つの事実が必要です。
スレッドが開くチャネル
質問がどのサーバーから来たのか
誰が質問しているのか、そうすればスレッドに引き込まれることができます
チケットの名前と問題の 1 行の概要
興味深いのは、これら 4 つが 3 つのまったく異なる方法で到着し、最後の 1 つだけが AI によって作成されることです。
Discord ID を入力することはなく、モデルがあなたのチャンネルや役割を認識することもありません。
一度選んだもの。チケット チャネルとサポートの役割。テンプレートをインストールするときにドロップダウンからそれらを選択すると、Quickchat AI がそれらをアクションに組み込みます。モデルはそれらを決して参照せず、変更することもできません。
Discord が提供するもの。 Discord 統合は、メッセージが転送されるたびに、どのサーバーから来たのか、誰が話しているのかを Quickchat AI に {{metadata_discord_guild_id}} と {{metadata_discord_author_id}} として伝えます。ボット トークンは、送信時に挿入され、通話が記録されるすべての場所で編集されるシステム トークンと同じ方法で到着します。モデルによって型指定されるものはなく、パブリック チャットや API メッセージによって提供されるものはありません。
エージェントが書くもの。チケットの件名と 1 行の問題の概要。ビルド全体で判断を下すのはこれらだけであるため、後のチューニングは短くなります。
ステップ 1: エージェントを作成し、ナレッジ ベースを追加する
エージェントを作成し、ナレッジベースを開いてドキュメントを追加します。ヘルプ ドキュメントの URL を貼り付けて、Quickchat AI がそれをスクレイピングするか、記事を直接作成します。ドキュメントから AI サポート エージェントを作成すると、スクレイピング パスが詳細に説明されます。
Beacon のナレッジ ベース: API キーのリセット、空のダッシュボード、追跡スニペット、プランと請求のインストール、データのエクスポート。
Beacon のドキュメントでは、エンタープライズ SSO の設定という 1 つのことを意図的に取り上げていません。

p.このギャップはこのチュートリアルの見落としではなく、重要な点です。どのサポート ナレッジ ベースにも、答えられない質問があり、このボットの価値は、質問に遭遇したときに何を行うかにあります。 SSO は、このガイド全体を通しての質問です。
メンバーは、エージェントが表示できるチャネルで @メンションするか、/ask コマンドを使用してエージェントにアクセスします。登録するものは何もありません。
ステップ 2: エージェントにいつ応答するか、いつエスカレーションするかを指示する
エージェントの指示は、アイデンティティに関する AI メイン プロンプトに入力されます。これをあなたの会社に合わせて貼り付けてください:
あなたは、ソフトウェア チーム向けのホスト型製品分析ツールである Beacon のサポート アシスタントである Lumen です。 Beacon コミュニティ Discord で顧客からの質問に答えます。
可能な限りナレッジベースから回答してください。回答は 2 ～ 3 文の短く実用的なものにし、具体的な手順を示してください。
質問がナレッジ ベースでカバーされていない場合は、推測したり、手順を考え出したりしないでください。その情報を持っていないことをはっきりと伝えてから、サポート チケットを開きます。
1. 問題を説明する短い 2 ～ 4 単語の ticket_subject を指定して、open_support_ticket を 1 回呼び出します。
2. 次に、add_ticket_requester を呼び出して、顧客をスレッドに参加させます。
3. 次に、次の 1 行の issue_summary を指定して post_in_ticket を呼び出します。

[切り捨てられた]

## Original Extract

Build a no-code AI Discord ticket bot. It answers from your knowledge base, opens a private ticket thread when it cannot, pings your support role, and hands off to a human.

AI Discord Ticket Bot: Answer First, Then Escalate (No Code) | Quickchat AI - AI Agents
for
Copy logo to clipboard (png)
Copy logo to clipboard (svg)
Download brand assets Use Cases
Customer Support Resolve tickets and cut response times
Sales & Lead Generation Qualify leads and recommend products
Ecommerce Recover carts and guide shoppers
IT & Internal Helpdesk Answer employee and internal questions
Enterprise Security, scale, and control for large teams
Platform The AI agent platform Releases What's new in Quickchat AI Contact us Company
Why Us Pricing Resources Log in Start for free Start for free Solutions Use Cases
Customer Support Sales & Lead Generation Ecommerce IT & Internal Helpdesk Enterprise Channels & Integrations
Website chat WordPress WhatsApp Discord Telegram Shopify Intercom HubSpot Zendesk More integrations Platform Releases Contact us Why Us Pricing Resources Company
AI Discord Ticket Bot: Answer First, Then Escalate (No Code)
; the original PNG/JPG stays as the
fallback and remains the og:image for social scrapers. --> Quickchat AI answers support questions in your Discord server from your own documentation, and opens a private ticket thread only when it cannot. This guide builds that bot end to end, with no code, in about half an hour.
The result behaves like a good first-line agent. A member asks how to reset an API key and gets the real steps in the channel. A member asks something the docs do not cover, and instead of guessing, the bot says so, opens a private thread, pulls the member into it, pings your support role with a one-line summary, and hands the conversation to a human who can answer from your Inbox.
You need two things, both free:
a Quickchat AI Agent ( sign up here and use it for free )
a Discord server where you are an admin
The screenshots below come from a test Agent called Lumen , the support bot for a fictional product-analytics company called Beacon . The company is invented so the example stays neutral, but every conversation, every API call and every ticket thread shown here was produced by a real Agent running the real reply pipeline against a real Discord server.
One decision per question. On the happy path nothing at all is written to Discord.
Three AI Actions do the escalation work, and Quickchat AI installs all three from one gallery card:
What is a knowledge base, and why does it set your ticket volume?
If you have not built an AI Agent before, this is the one concept worth two minutes.
A knowledge base is your own documentation, loaded into your Agent and indexed so it can be searched during a conversation. You can point Quickchat AI at your help-docs URL and let it scrape the pages, paste articles in by hand, or upload files. At reply time the Agent searches it and answers from what it finds, rather than from whatever the underlying model happens to remember about the world. If you are starting from nothing, how to build a chatbot knowledge base covers what to write and in what order.
Everything your docs cover gets answered. Everything they do not becomes a ticket.
That is the whole lever for this build. Deflection is not a prompt trick and it is not a model capability: it is your coverage. If the bot escalates more than you expected, the tickets are telling you which pages you have not written yet.
Illustrative rather than measured. The point is the direction: writing docs is the highest-leverage thing you can do here.
Can an AI actually answer support questions in Discord?
For the questions your docs cover, yes, and with the real steps rather than a paraphrase. Here is the deflection case from the test server, unedited:
Answered in the channel, from the knowledge base article. No ticket was created, because none was needed.
The more interesting half is what happens when the answer is not there. The bot says so:
The honest failure. Beacon’s knowledge base deliberately omits enterprise SSO, so this is the path that opens a ticket.
An Agent that invents a plausible SSO procedure is worse than one that opens a ticket, so this behaviour is worth deliberately testing rather than hoping for. We come back to it in testing .
Two ways to connect Discord, and why tickets need the second
Quickchat AI has two Discord connections, and the difference matters before you start.
Both answer questions in your server. Only one can run AI Actions.
One-click is the fast path. In External Apps you open Discord , click Add to your Discord server , authorize it, and the shared Quickchat AI bot is live in your server in about a minute. It answers when members @mention it or use /ask , and for a lot of servers that is the whole job.
You do get to name it. Before you authorize, there is a Bot name in the new server field, prefilled with your Agent’s name, and every connected server keeps its own Bot name in this server box you can change later. It is a Discord nickname, so it is per server and capped at 32 characters, and clearing it puts the bot back to Quickchat AI. What stays shared is the avatar and the underlying @username , because one Discord application is serving every server that installs it.
One-click does not mean generic. The bot carries your Agent’s name into the server, and you can change it per server afterwards.
Advanced setup registers your own Discord application. You create a bot in Discord’s developer portal and paste its token into Quickchat AI once. It takes about ten minutes and it gets you your own avatar and @username , direct messages, several bots in one server, and the thing this guide needs: AI Actions .
The reason is worth stating plainly rather than treating as a quirk. An AI Action calls the Discord API as your bot , using a credential Quickchat AI injects into the request at send time. The shared one-click bot’s token belongs to the shared bot, not to you, so it is not available for your Actions to send. If you open the Discord gallery without your own app connected, it tells you so and disables the cards rather than letting you build something that cannot run.
External Apps is where both paths start. Advanced setup lives in a callout at the bottom of the Discord panel.
What you need before you start
Connect your own Discord app. In External Apps , open Discord , find Use your own Discord app and choose Advanced setup . Create the application and bot in Discord’s developer portal, turn the Message Content Intent on, leave Require OAuth2 Code Grant off, and paste the bot token. Quickchat AI validates it as you type. The Create an AI Discord bot guide covers the portal side in detail if you have not done it before.
Grant the bot only what a support bot needs. On the ticket channel: View Channels, Send Messages, Read Message History, Create Private Threads, Send Messages in Threads, and Manage Threads. It needs no Kick, Ban or Timeout permission, and deliberately holds none in the test server. A support bot has no reason to moderate, and a permission it does not hold is one you never have to reason about.
In Discord, create the two things the Actions will point at. This part happens in the Discord client, not in Quickchat AI, and it is the only setup you do there:
A channel where tickets will open. Any text channel will do. Use the + next to your channel list, or reuse one you have. In the test server it is #support .
A role to page when a ticket opens. Server Settings , then Roles , then Create Role . In the test server it is @support .
Three places, and only one of them is public.
Both of these come back in Step 3 . When you install the ticket template, Quickchat AI reads your server live from Discord and shows you three dropdowns, one of them a Role picker listing the roles it found. You choose your support role there rather than typing an ID anywhere.
The Role picker in Step 3. It reads the list from Discord, so a role you create now appears here in a moment.
One detail that will bite you if you skip it: that picker only lists roles Discord reports as mentionable , and it leaves out @everyone . If your support role is not set mentionable, it simply will not appear in the dropdown, with nothing explaining why. The filter is deliberate, because the template’s whole job is to ping that role. Set it while you are creating the role and you never meet the problem:
Server Settings, Roles, pick your support role, and turn this on under Display. Without it the role never reaches the picker.
How does the bot decide to answer or open a ticket?
Most Discord ticket bots do not decide anything. Ticket Tool, TicketsBot and the rest post a pinned panel with a Create Ticket button, or register a slash command. A member clicks it, a ticket opens, and your team reads it. That is a fine design and it is completely predictable, but it means the ten people who all asked how to reset an API key this week produced ten tickets.
The same ten questions. The difference is what your team ends up reading.
An answer-first bot reads your docs before it decides. And when it does escalate, opening one ticket takes exactly four facts:
which channel the thread opens in
which server the question came from
who is asking , so they can be pulled into the thread
what the ticket is called , and a one-line summary of the problem
The interesting part is that those four arrive in three completely different ways, and only the last one is written by the AI:
You never type a Discord ID, and the model never sees your channel or your role.
The ones you pick once. Your ticket channel and your support role. You choose them from dropdowns when you install the template, and Quickchat AI bakes them into the Action. The model never sees them and cannot change them.
The ones Discord supplies. With every message it forwards, the Discord integration tells Quickchat AI which server it came from and who is speaking, as {{metadata_discord_guild_id}} and {{metadata_discord_author_id}} . Your bot token arrives the same way, as a System Token injected at send time and redacted everywhere a call is logged. None of it is typed by the model, and none of it can be supplied by a public chat or API message.
The ones the Agent writes. The ticket subject and the one-line issue summary. These are the only judgment calls in the entire build, which is why tuning later is short.
Step 1: Create the Agent and give it a knowledge base
Create your Agent, then open Knowledge Base and add your documentation. Paste your help-docs URL and let Quickchat AI scrape it, or write articles directly. Create an AI support agent from your documentation covers the scraping path in detail.
Beacon’s knowledge base: resetting an API key, an empty dashboard, installing the tracking snippet, plans and billing, and exporting data.
Beacon’s docs deliberately do not cover one thing: enterprise SSO setup. That gap is not an oversight in this tutorial, it is the point. Every support knowledge base has questions it does not answer, and the entire value of this bot is what it does when it hits one. SSO is that question throughout this guide.
Members reach the Agent by @mentioning it in a channel it can see, or with the /ask command. There is nothing to register.
Step 2: Tell the Agent when to answer and when to escalate
The Agent’s instructions go in the AI Main Prompt on Identity . Paste this, adapted to your company:
You are Lumen, the support assistant for Beacon, a hosted product-analytics tool for software teams. You answer questions from customers in the Beacon community Discord.
Answer from your Knowledge Base whenever you can. Keep answers short and practical, two or three sentences, and give the concrete steps.
If a question is not covered by your Knowledge Base, do not guess and do not invent steps. Say plainly that you do not have that information, then open a support ticket:
1. Call open_support_ticket once, with a short 2 to 4 word ticket_subject describing the problem.
2. Then call add_ticket_requester to bring the customer into the thread.
3. Then call post_in_ticket with a one line issue_summary of

[truncated]
