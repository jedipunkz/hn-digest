---
source: "https://quickchat.ai/post/connect-ai-agent-to-google-sheets"
hn_url: "https://news.ycombinator.com/item?id=48665781"
title: "Connect Your AI Agent to Google Sheets"
article_title: "Connect your AI Agent to Google Sheets | Quickchat AI - AI Agents"
author: "piotrgrudzien"
captured_at: "2026-06-24T21:29:41Z"
capture_tool: "hn-digest"
hn_id: 48665781
score: 1
comments: 0
posted_at: "2026-06-24T21:21:11Z"
tags:
  - hacker-news
  - translated
---

# Connect Your AI Agent to Google Sheets

- HN: [48665781](https://news.ycombinator.com/item?id=48665781)
- Source: [quickchat.ai](https://quickchat.ai/post/connect-ai-agent-to-google-sheets)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T21:21:11Z

## Translation

タイトル: AI エージェントを Google スプレッドシートに接続する
記事のタイトル: AI エージェントを Google スプレッドシートに接続する |クイックチャット AI - AI エージェント
説明: Quickchat AI エージェントを Google スプレッドシートに接続して、見込み客、未回答の質問、機能リクエスト、デモの予約を正確なアクション設定とコピーするプロンプトとともに記録するためのステップバイステップ ガイドです。

記事本文:
AI エージェントを Google スプレッドシートに接続する |クイックチャット AI - AI エージェント
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
AI エージェントを Google スプレッドシートに接続する
AI エージェントはサイトのすべての訪問者と会話します。購入の準備ができている人もいれば、答えられない質問をする人、まだ持っていない機能を要求する人、デモを希望する人もいます。デフォルトでは、チャットが終了すると、これらはすべて消去されます。
このガイドでは、AI エージェントを Google スプレッドシートに接続して、各瞬間がチームがすでに使用しているシートの行になるようにする方法を説明します。 Zapier も Webhook もエンジニアもいません。必要なものは次の 2 つです。
Quickchat AI エージェント (ここでサインアップして無料で使用します)
このメカニズムは AI アクションです。エージェントが実行できるカスタム HTTP リクエストです。

会話中に作ります。 Quickchat AI には、ワンクリックでシートとスターター ログ アクションを作成する Google スプレッドシート接続があり、そこからそれを作成し、必要なだけレポートを追加できます。最後には 4 つの作業レポートが作成され、それぞれを自分でテストすることになります。
これは長くて正確なウォークスルーです。 AI アクションの正規リファレンスは、 docs.quickchat.ai/ai-agent/actions のドキュメントにあります。カスタム アクションの別の実際の例については、「AI アクションを使用して Slack 通知を送信する」を参照してください。
4 つのレポート。それぞれが 1 つの Google スプレッドシートの独自のタブに書き込まれます。
以下のスクリーンショットは、SaaS ビジネス向けのサブスクリプション分析プラットフォームである架空の会社 Tideline 用に構築されたテスト エージェントからのものです。会社は発明されたものであるため、例は中立的なままですが、ここに示されているすべての会話とすべての行は、実際の応答パイプラインを実行している実際のエージェントによって生成されました。説明を進める際には、自分の会社の詳細情報を使用してください。
Google スプレッドシートの統合の仕組み
この機能全体は 1 つのアイデアに基づいています。つまり、AI アクションは記述された HTTP リクエストであり、Google シートの行はそのようなリクエストの 1 つです。
AI アクションは 4 つの部分で構成されます。説明は、モデルを呼び出すかどうかを決定するためにモデルが読み取るものです。パラメータは会話から入力されるものです。
さらにいくつかの事実により、この投稿の残りの部分を理解しやすくなります。
行の書き込みは 1 回の API 呼び出しです。 Google スプレッドシートには、タブに行を追加する追加エンドポイントがあります。各レポートは、本文に行値を含むエンドポイントへの POST です。
柱は固定されていません。それらはあなたのアクションが送信するものです。アクションの本文とシートのヘッダー行を編集して列を決定するため、同じ統合でリード、バグレポートなどを記録できます。
最小権限のアクセス。 Google 接続は drive.file スコープのみを要求します。

これにより、アプリが作成したファイルへの排他的なアクセスが許可されます。ドライブの残りの部分は表示されません。
エージェントがあなたの資格情報を閲覧することはありません。リクエストには Authorization ヘッダーが含まれており、その値はプレースホルダー {{google_sheets_access_token}} です。 Quickchat AI は、モデルがその役割を果たした後、そのプレースホルダーを実際の自動更新トークンで埋めます。トークンがプロンプトに表示されることはありません。
2 つの組み込み変数は、パラメーターとして定義されずにあらゆるアクションで使用できます。 {{conversation_url}} は受信トレイ内の会話へのディープ リンクであり、 {{conversation_channel}} は訪問者が使用したチャネル (Web ウィジェット、Slack、WhatsApp など) です。
ステップ 1: AI エージェントを作成し、知識を与える
Quickchat エージェントの動作は、その ID (メイン プロンプト) と応答するために与えられた知識という 2 つの場所から生じます。アクションと MCP は、Google スプレッドシートへの書き込みなど、実行できる機能を拡張する場所です。このガイドは、アイデンティティ (このステップとステップ 4 ) とアクションと MCP (ステップ 2、3、および 6 ) で機能し、AI プレビュー (ステップ 5 ) ですべてをテストします。
サインアップしたら、左側のサイドバーで [ID] を開きます。 AI メイン プロンプトでは、エージェントが何であるか、およびエージェントがどのように動作するかを説明します。製品についての短く正確な説明を入力し、質問に直接答えられるように、記述すべき事実 (プラン、価格、どの統合が存在するか) を知識に組み込みます。
エージェントの名前とメイン プロンプト: 製品に関する短く正確な説明、およびエージェントがどのように動作するか。ステップ 4 で、このプロンプトの最後にレポート ブロックを追加します。
プロンプトをすべて自分で作成する必要はありません。このガイドのすべてのプロンプト、アクションの説明、リクエスト本文はコピー＆ペーストのブロックとして提供されているため、入力するのではなく貼り付けることになります。あなたは戻ってきます

ステップ 4 でこの画面にレポート指示を追加します。
サイドバーで「アクションと MCP」を開き、「アクションの追加」をクリックして、「Google スプレッドシート」を選択します。
[アクションと MCP]、[アクションの追加]、[Google スプレッドシート] の順に選択します。
接続ダイアログが開き、アクセスを許可する前に、ワンクリックで設定される内容が説明されます。つまり、ドライブ内の新しい Google シートと、リード用にプリセットされたすぐに使用できるログ記録アクションです。これを確認してオンにします。
このダイアログでは、アクセスを許可する前に、接続時に設定される 2 つのことについて説明します。
「接続」をクリックします。 Google は、どのアカウントを使用するか、次にどの権限を付与するかを尋ねます。要求される唯一のスコープは drive.file です。これにより、統合は作成したファイルのみにアクセスでき、ドライブの残りの部分にはアクセスできなくなります。
Google 独自の画面: シートを保存するドライブのアカウントを選択します。次の画面では、drive.file 権限を付与します。これにより、統合は作成したファイルのみにアクセスできるようになります。
承認すると、Quickchat AI の [アクションと MCP] ページに戻り、接続されたダイアログが開きます。
接続されました。[アクション] ページに戻ります。 Quickchat AI はシートと 1 つのスターター ログ アクションを作成し、スイッチをオフにしてレビューの準備を追加しました。
ワンクリックで次の 2 つの処理が自動的に実行されます。
Google ドライブに 1 つのタブ ( Leads ) とヘッダー行を含むスプレッドシートを作成しました。
log_lead_to_google_sheet という単一の無効な AI アクションを追加しました。これはすでにそのシートに接続されています。
このスターター アクションは、通常の完全に編集可能な AI アクションです。これが最も一般的なケースであるため、統合によりリード獲得用にプリセットされますが、名前、列、説明はすべて変更できます。このガイドでは、これを確認し、最初のレポートとしてオンにしてから、ステップ 6 で同様のレポートをさらに 3 つ作成します。
ステップ 3: リードアクションを確認して有効にする
自動作成された log_lead_to_google_sheet アクションは次のとおりです。

[カスタム アクション] の下にリストされており、オフになっています。有効にする前に、それを開いて構成を確認してください。これは、他のレポートが従うテンプレートなので、完全に理解する価値があります。
4 つの報告アクション。最初の log_lead_to_google_sheet のみが自動的に作成されます。残りの 3 つはステップ 6 で追加します。
アクションの先頭には、その名前とエージェントが収集するパラメータが保持されます。
エージェントが会話から入力したアクション名とパラメータ。
「API エンドポイント」セクションまで下にスクロールします。 Quickchat AI がシートとこのアクションを一緒に作成したため、URL にはスプレッドシートの ID がすでに入力されており、メソッドは POST で、ヘッダーが設定されています。
完全なエンドポイント URL にはシートの ID がすでに含まれています。 Query Params は空のままです。valueInputOption=RAW 設定が URL に反映されます。 Authorization ヘッダーにはトークン プレースホルダーが含まれます。
API ドキュメントに記載されている {SPREADSHEET_ID} は、実際のプレースホルダーではありません。接続時に、Quickchat AI はシートを作成し、その実際の ID をこのアクションの URL に直接書き込みました。そのため、フィールドにはすでに完全なアドレスが表示されています。
https://sheets.googleapis.com/v4/spreadsheets/<シート ID>/values/Leads:append?valueInputOption=RAW
valueInputOption=RAW を維持します。値を文字通りに挿入するため、= または + で始まるメッセージはスプレッドシートの数式として解析されません。手順 6 でこの正確な URL を他の 3 つのレポートにコピーし、タブ名のみを変更します。これにより、手動で ID を追跡する必要がなくなります。
ヘッダーはすべてのレポートで同じです (上のスクリーンショットから読み取ることができます)。
次に、列を柔軟にする部分です。 API エンドポイント セクション内にある [Body] タブに切り替えます。リクエストの本文は 1 行で、各セルは 1 つのパラメータです。セルの順序 i

n 本体はシート内の列の順序です。
{ "値" : [[ "{{email}}" , "{{what_they_asked}}" , "{{conversation_url}}" , "{{conversation_channel}}" ]] }
リクエストの本文は「本文」タブにあります。各 {{parameter}} は 1 つのセルであり、セルはシートの列と同じ順序で実行されます。
したがって、列は統合によってハードコーディングされません。これらは、この本文とタブのヘッダー行によって定義されます。本文とヘッダーを変更すると、レポートの記録内容が変更されます。実際のリードの会話がこの行にどのようにマッピングされるかは次のとおりです。
訪問者が言ったことはパラメータにマッピングされ、パラメータは列に順番にマッピングされます。 2 つの組み込み変数は自動的に入力されます。
パラメータは、エージェントが会話から入力するフィールドです。各パラメータの説明は、そこに何を配置するかをエージェントに指示するものです。
最後に、アクションの説明です。これは、モデルがアクションを呼び出すかどうかを決定するために読み取るものであるため、最も重要なフィールドです。以下のバージョンは、エージェントが行を書き込む前に実際に電子メールを受信するまで待機するように調整されています (理由については、調整セクションで詳しく説明します)。
チームがフォローアップできるように、適格な営業見込み客を記録します。このアクションのトリガーは、訪問者が購入に興味を示しながら電子メール アドレスを提供することです (価格、プラン、コスト、スケジュール、または予算について尋ねます)。これを呼び出すのは、訪問者が実際の電子メール アドレスを与えたターンのみです。その正確な電子メールを電子メールパラメータに入力します。訪問者が購入に興味を示しているものの、まだ電子メールを送信していない場合は、このアクションを呼び出さないでください。訪問者に応答して、まず電子メールを尋ねてください。電子メールが提供されていないターンでは決してこれを呼び出さないでください。また、空の電子メールまたは作成された電子メールを使用してこれを呼び出さないでください。
思考メッセージ (私たちのチームに関心があることに注意してください...) は、訪問の短いステータスです。

tor はアクションの実行中に表示されます。設定が正しいと思われる場合は、アクションをオンに切り替えます。
ステップ 4: レポートの指示をプロンプトに追加する
アクションの説明により、各レポートがいつ起動されるかが決まります。プロンプトは、別の補完的な役割を果たします。レポートはエージェントの仕事の一部であることを伝え、4 つのレポートが互いに重ならないようにします。
[アイデンティティ] に戻り、このブロックを AI メイン プロンプトの最後に貼り付けます。 4 つのレポートすべてをカバーしているため、貼り付けるのは 1 回だけです。
## チームへの報告
チームの Google スプレッドシートに有用な情報を記録するアクションがいくつかあります。彼らはバックグラウンドで静かに実行されます。これらは仕事の一部であり、オプションではありません。以下のいずれかの状況に該当する場合は、返信と同じターンに一致するアクションを呼び出します。同じターン内で、訪問者に応答し、アクションを呼び出すことができます。スプレッドシート、ログ記録、またはこれらのアクションについては訪問者に決して言及しないでください。疑問がある場合は、記録してください。
- 購入意向。訪問者は、価格、プラン、スケジュール、予算について尋ねます。彼らの質問に答えて、リードを奪いましょう。まだメールを共有していない場合は、チームがフォローアップするのに最適なメールのみを要求し、そのターンでは何も記録しないでください。リードアクションを呼び出すのは

[切り捨てられた]

## Original Extract

A step by step guide to connecting your Quickchat AI Agent to Google Sheets so it logs leads, unanswered questions, feature requests, and demo bookings, with the exact action settings and prompt to copy.

Connect your AI Agent to Google Sheets | Quickchat AI - AI Agents
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
Connect your AI Agent to Google Sheets
Your AI Agent talks to every visitor on your site. Some are ready to buy, some ask a question it cannot answer, some request a feature you do not have yet, and some want a demo. By default, all of that is gone when the chat closes .
This guide shows you how to connect your AI Agent to Google Sheets so each of those moments becomes a row in a sheet your team already uses. No Zapier, no webhooks, no engineer. You need two things:
a Quickchat AI Agent ( sign up here and use for free )
The mechanism is AI Actions : custom HTTP requests your Agent can make during a conversation. Quickchat AI has a one-click Google Sheets connection that creates a sheet and a starter logging action for you, and from there you shape it and add as many reports as you want. By the end you will have four working reports , and you will have tested each one yourself .
This is a long, exact walkthrough. The canonical reference for AI Actions lives in the docs at docs.quickchat.ai/ai-agent/actions . For another worked example of a custom action, see Send Slack notifications with AI Actions .
Four reports, each writing to its own tab of one Google Sheet:
The screenshots below come from a test Agent built for a fictional company, Tideline , a subscription-analytics platform for SaaS businesses. The company is invented so the example stays neutral, but every conversation and every row shown here was produced by a real Agent running the real reply pipeline. Use your own company’s details when you follow along.
How the Google Sheets integration works
The whole feature rests on one idea: an AI Action is a described HTTP request, and a Google Sheet row is one such request.
An AI Action has four parts. The description is what the model reads to decide whether to call it; the parameters are what it fills in from the conversation.
A few more facts make the rest of the post easier to follow.
Writing a row is one API call. Google Sheets has an append endpoint that adds a row to a tab. Each report is a POST to that endpoint with the row values in the body.
The columns are not fixed. They are whatever your action sends. You decide the columns by editing the action’s body and the sheet’s header row, so the same integration can log leads, bug reports, or anything else .
Least-privilege access. The Google connection requests only the drive.file scope, which grants access exclusively to files the app creates . It cannot see the rest of your Drive.
The Agent never sees your credentials. The request carries an Authorization header whose value is a placeholder, {{google_sheets_access_token}} . Quickchat AI fills that placeholder with a real, auto-refreshed token after the model has done its part. The token never enters the prompt.
Two built-in variables are available to any action without being defined as parameters: {{conversation_url}} , a deep link back to the conversation in your Inbox, and {{conversation_channel}} , the channel the visitor used (web widget, Slack, WhatsApp, and so on).
Step 1: Create your AI Agent and give it knowledge
A Quickchat Agent’s behavior comes from two places: its Identity (the main prompt) and the knowledge you give it to answer from. Actions & MCPs is where you extend what it can do , such as writing to a Google Sheet. This guide works in Identity (this step and Step 4 ) and Actions & MCPs (Steps 2, 3 and 6 ), and tests everything in AI Preview ( Step 5 ).
After you sign up , open Identity in the left sidebar. The AI Main Prompt is where you describe what your Agent is and how it should behave. Give it a short, accurate description of your product, and put the facts it should be able to state (plans, prices, which integrations exist) into its knowledge so it can answer questions directly.
The Agent’s name and its main prompt: a short, accurate description of your product and how the Agent should behave. You add the reporting block to the end of this prompt in Step 4 .
Do not worry about writing all of that prompt yourself. Every prompt, action description, and request body in this guide is given as a copy-paste block , so you will paste them rather than type them out. You will return to this screen in Step 4 to add the reporting instructions.
Open Actions & MCPs in the sidebar, click Add Action , and choose Google Sheets .
Actions & MCPs, then Add Action, then Google Sheets.
The connect dialog opens and, before you grant any access, spells out what the one click will set up: a new Google Sheet in your Drive, and a ready-to-use logging action, preset for leads, that you review and switch on.
The dialog explains the two things connecting will set up, before you grant any access.
Click Connect . Google asks which account to use and then which permission to grant. The only scope requested is drive.file , which lets the integration touch only the files it creates , never the rest of your Drive.
Google’s own screen: pick the account whose Drive the sheet should live in. The next screen grants the drive.file permission, which lets the integration touch only the files it creates.
When you approve, you land back on the Actions & MCPs page in Quickchat AI, with the connected dialog open:
Connected, back on the Actions page. Quickchat AI created a sheet and one starter logging action, added switched off and ready to review.
That one click did two things automatically:
Created a spreadsheet in your Google Drive, with one tab ( Leads ) and a header row.
Added a single, disabled AI Action called log_lead_to_google_sheet , already wired to that sheet.
That starter action is a normal, fully editable AI Action . The integration presets it for lead capture because that is the most common case, but the name, the columns, and the description are all yours to change. In this guide you will review it, switch it on as the first report, then build three more like it in Step 6 .
Step 3: Review and enable the lead action
The auto-created log_lead_to_google_sheet action is listed under Custom Actions , switched off. Open it to review the configuration before enabling it. This is the template every other report follows , so it is worth understanding in full.
The four reporting actions. Only the first, log_lead_to_google_sheet , is created for you; the other three you add in Step 6 .
The top of the action holds its name and the parameters the Agent collects:
The action name and the parameters the Agent fills in from the conversation.
Scroll down to the API Endpoint section. Because Quickchat AI created the sheet and this action together, the URL is already filled in with your spreadsheet’s ID , the method is POST , and the headers are set:
The full endpoint URL already contains your sheet’s ID. Query Params stays empty: the valueInputOption=RAW setting rides along in the URL. The Authorization header carries the token placeholder.
The {SPREADSHEET_ID} you see written in API documentation is not a live placeholder . At connect time, Quickchat AI created your sheet and wrote its real ID straight into this action’s URL, which is why the field already shows the complete address:
https://sheets.googleapis.com/v4/spreadsheets/<your sheet ID>/values/Leads:append?valueInputOption=RAW
Keep valueInputOption=RAW : it inserts values literally, so a message that starts with = or + is not parsed as a spreadsheet formula. You will copy this exact URL into the other three reports in Step 6 , changing only the tab name, so there is never an ID to track down by hand.
The headers are the same for every report (you can read them off the screenshot above):
Now the part that makes the columns flexible . Switch to the Body tab, still inside the API Endpoint section. The request body is a single row, and each cell is one parameter . The order of the cells in the body is the order of the columns in the sheet:
{ "values" : [[ "{{email}}" , "{{what_they_asked}}" , "{{conversation_url}}" , "{{conversation_channel}}" ]] }
The request body lives in the Body tab. Each {{parameter}} is one cell, and the cells run in the same order as the sheet’s columns.
So the columns are not hardcoded by the integration. They are defined by this body plus the header row of the tab. Change the body and the header, and you change what the report records. Here is how a real lead conversation maps onto this row:
What the visitor said maps to parameters, and the parameters map to columns, in order. The two built-in variables are filled automatically.
The parameters are the fields the Agent fills in from the conversation. The description of each parameter is what tells the Agent what to put there:
Finally, the action description . This is the single most important field , because it is what the model reads to decide whether to call the action at all. The version below is tuned so the Agent waits until it actually has an email before writing a row (more on why in the tuning section ):
Record a qualified sales lead so the team can follow up. The trigger for this action is the visitor providing their email address while showing buying interest (they asked about pricing, plans, cost, timeline or budget). Call this only in a turn where the visitor has just given you a real email address; put that exact email in the email parameter. If a visitor shows buying interest but has not given an email yet, do NOT call this action: answer them and ask for their email first. Never call this in a turn where no email has been provided, and never with an empty or made-up email.
The thinking message ( Noting your interest for our team... ) is the short status the visitor sees while the action runs. When the configuration looks right, switch the action on .
Step 4: Add the reporting instructions to your prompt
The action descriptions decide when each report fires. The prompt does a different, complementary job: it tells the Agent that reporting is part of its work , and it keeps the four reports from stepping on each other.
Go back to Identity and paste this block at the end of your AI Main Prompt . It covers all four reports, so you only paste it once:
## Reporting to the team
You have several actions that log useful information to the team's Google Sheet. They run silently in the background. They are part of your job and they are not optional. When one of the situations below applies, call the matching action in the same turn as your reply. You can both answer the visitor and call an action in the same turn. Never mention the spreadsheet, the logging or these actions to the visitor. When in doubt, log it.
- Buying intent. The visitor asks about pricing, plans, timeline or budget. Answer their question, then capture the lead. If they have not shared an email yet, only ask for the best email for the team to follow up and do not log anything on that turn. Call the lead action only on the

[truncated]
