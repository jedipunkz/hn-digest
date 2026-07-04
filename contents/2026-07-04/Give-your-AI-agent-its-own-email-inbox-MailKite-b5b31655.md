---
source: "https://mailkite.dev/blog/give-your-agent-an-inbox/"
hn_url: "https://news.ycombinator.com/item?id=48781354"
title: "Give your AI agent its own email inbox – MailKite"
article_title: "Give your AI agent its own email inbox — MailKite"
author: "bucabay"
captured_at: "2026-07-04T00:01:34Z"
capture_tool: "hn-digest"
hn_id: 48781354
score: 1
comments: 0
posted_at: "2026-07-03T23:47:29Z"
tags:
  - hacker-news
  - translated
---

# Give your AI agent its own email inbox – MailKite

- HN: [48781354](https://news.ycombinator.com/item?id=48781354)
- Source: [mailkite.dev](https://mailkite.dev/blog/give-your-agent-an-inbox/)
- Score: 1
- Comments: 0
- Posted: 2026-07-03T23:47:29Z

## Translation

タイトル: AI エージェントに独自の電子メール受信箱を提供 – MailKite
記事のタイトル: AI エージェントに独自の電子メール受信箱を与える — MailKite
説明: AI エージェントに、制御するドメイン上の実際の範囲指定されたアドレスを与えます。受信メールは解析された JSON としてevent.receivedループに到着し、エージェントはSend API経由で応答するか、MailKiteがaction:agentを指定したルートでエージェントを実行します。実用的なコード、セキュリティ上の警告、誠実な DIY 代替案
[切り捨てられた]

記事本文:
AI エージェントに独自の電子メール受信箱を提供 — MailKite ベータ版
初年度は 50% 割引になります。最初の 1,000 名の顧客の創業率は 50% オフです。
価格を見る→
メールカイト
機能 受信 送信 エージェント向け 価格設定ドキュメント ブログ 連絡先 サインイン · はじめる 機能 受信 送信 エージェント向け 価格設定ドキュメント ブログ 連絡先 サインイン
すべての投稿
June 10, 2026 · Gabe · 11 min read AI エージェントに独自のメール受信箱を与える
AI エージェントに、制御するドメイン上の実際の範囲指定されたアドレスを与えます。受信メールは解析された JSON としてevent.receivedループに到着し、エージェントはSend API経由で応答するか、MailKiteがaction:agentを指定したルートでエージェントを実行します。機能するコード、セキュリティ上の警告、そして正直な DIY の代替案。
マークダウンとしてコピー
ChatGPTで開く
クロードで開く
独自の電子メール受信箱を持つ AI エージェントは、制御するドメイン上の実際のアドレスを所有する自律型プログラムであるため、人間が関与することなく、検証コードを受信し、電子メールで作業を引き渡され、独自に返信することができます。この投稿は、初めてエージェントを受信トレイに接続する開発者向けであり、エージェントを構築する 2 つの方法を示しています。1 つは自分でエージェントを実行する (受信メールは解析された JSON として到着し、ハンドラーはモデルを呼び出し、エージェントは 1 つの mk.send() で応答します)。もう 1 つは、(私たちが構築した) MailKite にアクションが Agent であるルートでエージェントを実行させる方法です。いずれにせよ、IMAP も MIME 解析も、個人の Gmail アカウントが密かにボットに接続されることもありません。
これがエージェント持ち込みループの全体です。電子メールで送信し、署名を確認し、モデルに本文を渡し、同じクライアントを通じて返信します。ノード 18+ に貼り付けられたとおりに実行されます ( npm install mailkite Express )。
"express" からエクスプレスをインポートします。
"mailkite" から { MailKite } をインポートします。
const app = Express();
const mk = new MailKite (プロセス . env . MAILKITE_API_KEY );
const SECRET = プロセス。環境

。 MAILKITE_WEBHOOK_SECRET ;
app .use ( "/hooks/agent" , Express .raw ({ type : "application/json" }));
app .post ( "/hooks/agent" , async (req , res) => {
// 署名チェック、再生ウィンドウ、定数時間比較 — 1 回の呼び出し
if ( ! MailKite .verifyWebhook ( req .headers[ "x-mailkite-signature" ] , req .body , SECRET )) {
戻り値 .sendStatus ( 401 );
}
res .sendStatus ( 200 ); // 高速に応答します。エージェントを帯域外で実行する
const イベント = JSON .parse ( req .body);
if (event .type !== "email.received" ) return ;
// 本体を命令としてではなく、信頼できない INPUT として扱います (セキュリティのセクションを参照)。
const Answer = await runAgent ({
タスク: イベント .text 、
から: イベント。からのアドレス、
信頼できる: イベント。 auth .spf === "pass" && イベント 。 auth .dmarc === "パス" 、
});
mk .send を待ちます ({
from :event .to[ 0 ].address , // 送信先アドレスからの応答
へ: イベント。からのアドレス、
subject : `Re: ${event .subject } ` 、
inReplyTo :event .id , // 返信をスレッド化します
html : 答え .html 、
});
});
アプリ .listen ( 3000 );
これは完全に自律的な電子メール エージェントです。聞いて、考えて、応答します。 mk.send() は { id, status } を返すため、送信メッセージをログに記録できます。また、Python、Ruby、Go、PHP、および Java に対して同一のハンドラー形状が存在します。受信ドキュメントと送信ドキュメントを参照してください。依存関係を取得できない場合は、生の HTTP も機能しますが、verifyWebhook が 1 行で実行する HMAC 検証を手動で実行することになります。 SDK の方が良いでしょう。
実際にハンドラーに到達するもの
MailKite はエッジでメッセージをデコードするため、runAgent は MIME ではなくフィールドを取得します。これは、ピラーからの同じ受信 Webhook です。デコードされた text と html 、解決された threadId 、有効期間の短い署名付き URL としての添付ファイル、および認証ブロックです。
{
"id" : "msg_2Hk9…" ,
"type" : "email.received" ,
"差出人" : { "アドレス" : "ada@example.com" } ,
「に

" : [{ "アドレス" : "agent@yourco.dev" }] ,
"件名" : "Re: 請求書 #1042" ,
"text" : "良さそうです - 承認されました!" 、
"html" : "<p>良さそうです - 承認されました!</p>" ,
"threadId" : "<a1b2c3@mail.example.com>" ,
"auth" : { "spf" : "パス" 、 "dkim" : "パス" 、 "dmarc" : "パス" 、 "スパム" : "ハム" } 、
「添付ファイル」: [
{ "id" : "msg_2Hk9…:0" 、 "filename" : "po.pdf" 、 "contentType" : "application/pdf" 、
"サイズ" : 18213 , "url" : "https://api.mailkite.dev/att/2Hk9…:0?exp=…&sig=…" }
】
}
認証ブロックは、SPF、DKIM、および DMARC が合格したかどうかをエージェントに伝えます。その考えを持ち続けてください。すぐに負荷がかかります。
エージェントの受信トレイを実行する 2 つの方法
上記のループを所有するか、すべてを MailKite に渡します。同じインバウンドエッジ。違いは、エージェントのターンがどこで実行されるかです。
自分のエージェントを連れてくる
メールで送信
MX エッジ解析 + 認証
Webhook の署名付き JSON
あなたのエージェント、あなたのモデル
mk.send() 応答出力
MailKite に実行させます
メールで送信
MX エッジ解析 + 認証
エージェントルートアクション: 'エージェント'
キューエージェントの実行
返信 send_email
青 = MailKite によって運営されています。上のレーン: Webhook とモデル ループを作成します。一番下のレーン: エージェントのターンがパイプライン内で実行されます。
エージェント受信トレイを実行するには 2 つの方法があります。独自のエージェントを使用する (ループを所有する) か、アクション「エージェント」を含むルートを使用します (MailKite は永続キューでループを実行します)。どちらの方法でも同じ解析された受信エッジ。
MailKite にエージェントを実行させます (ルート アクション「エージェント」)
ループをホストしたくない場合は、ルート自体をエージェントにします。アクション「agent」を持つルートには、agentPrompt (フリーテキストの指示) が含まれており、MailKite はすべての受信メッセージに対してモデル ループを実行します。
await mk .createRoute ({
一致 : "support@yourco.dev" 、
アクション: "エージェント" 、
AgentPrompt : 「ドキュメントからの請求に関する質問に答えてください。それ以外の場合は、human@yourco.dev にエスカレーションしてください。」 、
});
実行は行われません

インバウンドリクエストのクロック。独自の実行バジェットを使用して Cloudflare キュー ( mailkite-agent-runs ) にエンキューするため、取り込みは引き続き高速に返され、遅いモデル呼び出しによってパイプラインが中断されることはありません。各実行は 8 つのツール ラウンドに制限され、クラッシュ防止のバックストップとして cron reaper によって 5 分で中止され、ダッシュボードでルートごとにドリルダウンできる完全なトランスクリプトとして記録されます。システム プロンプトにより、次のセクションで説明する安全ルールが組み込まれます。電子メールの本文は信頼できず、返信は 1 回までで、返信のない送信者や自動送信者には決して返信しません。エージェントは、内部 send_email ツール ( inReplyTo 経由でスレッド化される)、同じ /v1/send 独自のループ呼び出しを使用して応答します。詳細については、inbox-agents docs を参照してください。
エージェントの頭脳があなたのコード、つまりモデル、ツール、状態である場合、Bring-your-Own が勝ちます。インフラストラクチャを実行したくない場合、およびジョブが「この受信メールを読んで応答するかエスカレーションする」である場合は、組み込みルートが最適です。
フラグを立てなければならない部分: 受信メールは信頼できない入力です
ここであなたを遅らせます。私自身がこの穴に入ったので。エージェントが電子メールの内容に従った瞬間、その電子メール本文はプロンプト挿入ベクトルになります。 From: はプレーン テキストであるため、誰でも送信者を偽造でき、エージェントに何をすべきかを伝えるだけで、単純なループがそれに従うようになります。
これが、認証ブロックがペイロード内にある理由であり、ループが盲目的に動作するのではなく信頼できるフラグを渡す理由です。送信者の指示を重み付けする前に、少なくとも SPF と DMARC が合格したかどうかを確認できます。ただし、認証のチェックは必要ですが、十分ではありません。プロンプト インジェクションから抜け出すためのプロンプトを表示することはできません。本当の答えは構造的なものであり、だまされたエージェントが実行できることさえ制限されています (所有者限定のツール、返信は 1 つ、リンクに対する操作は不要)。これは組み込みエージェントが使用するのと同じ推論であり、Wh で間違いを正直に書きました。

エージェントのセキュリティに関する議論が増えていませんか? 。重要なことについてループを指示する前に、この内容を読んでください。
自分で構築する場合とそうでない場合
エージェントの受信トレイは珍しいものではありません。私たちなしでも組み立てることができますし、場合によってはそうすべきです。正直な代替案:
実際のメールボックスに対する IMAP ポーリング。エージェントに Gmail/Workspace または Fastmail アカウントを与え、IMAP または Gmail API をポーリングします。それは機能しますが、MIME 解析、スレッド処理、添付ファイルのデコード、および実際には個人のアクセス許可を持つ個人のアカウントであるメールボックスを所有していることになります。エージェントが本当に人間の既存の受信箱内に存在する必要がある場合に最適です。
受信消印 + 独自のループ。 Postmark の受信解析では、解析された JSON も POST します。上記と同じ runAgent ループをポイントします。送信用の消印をすでに使用していて、受信側の半分だけが必要な場合は、大丈夫です。
Cloudflare の電子メール ワーカー。受信プリミティブ: 電子メールはワーカーの email() ハンドラーにヒットします。解析、ルーティング、管理層はありません。あなたがそれを構築します。エッジ全体を所有したい場合に適しています。 (私たちはCloudflare上にMailKiteを構築しており、これが一番上に置くレイヤーです。)
自己ホストのハラカ。独自の MX を実行します。総合的な制御、総合的な運用: SMTP サーバー、スパム フィルタリング、TLS、稼働時間はすべてあなたのものです。実際の規模で、または厳格なデータ常駐ルールの下で利用する価値があります。
DIY の形状はどの場合でも同じです。MX レコード (またはメールボックス)、生の MIME をフィールドに変換するもの、署名と SPF/DKIM/DMARC チェック、および応答パスです。 MailKite は、そのスタック (MX エッジ、解析、認証、署名および再試行された Webhook、および Send API) を組み立てたものであるため、上部の 25 行が統合全体です。専任のメール インフラストラクチャ エンジニアの勤務時間未満では、取引自体が採算が合う傾向があります。それより上の場合、またはエージェントが人間の実際のメールボックスに常駐する必要がある場合、上記のいずれかはありません。

彼は選んだほうがいいよ。
エージェントに MailKite をツールとして使用させる (MCP)
ここまではすべて に届くメールです。もう 1 つの方向は、エージェントが電子メールを通じて連絡することであり、MailKite はその API 全体をホストされた MCP サーバー上でエージェント ネイティブ ツールとして公開します。任意の MCP クライアントをそれに向けます。
クロード mcp add --transport http mailkite https://mcp.mailkite.dev/mcp
その後、エージェントは mailkite_send 、 mailkite_get_message 、 mailkite_list_domains などを、手動の HTTP ではなくファーストクラスのツール呼び出しとして呼び出します。同じツールをエディタに接続する Claude Code プラグインもあります。受信トレイとこのツール セットを持つエージェントは、サポート アドレスをエンドツーエンドで実行できます。受信メッセージを読み、何かを調べ、返信するなど、すべてをツール呼び出しとして実行できます。
AI エージェントに独自の電子メール アドレスを与えることはできますか?
はい。 MailKite にドメインを指定し、agent@yourco.dev のようなアドレスを選択して、Webhook を設定します。受信メールは JSON に解析され、email.received イベントとして POST されます。エージェントは mk.send() で応答します。これは、ボットに組み込まれた個人アカウントではなく、ドメイン上の実際のスコープ付きメールボックスです。
エージェントはサーバー上で実行する必要がありますか? それとも MailKite 上で実行する必要がありますか?
どちらも機能します。エージェントのロジックがコードである場合は、ループを自分でホストします (Webhook ルート → エンドポイント → モデル → API の送信)。 MailKite が独自の永続キューでモデル ループを実行し、トランスクリプトを渡したい場合は、アクション「エージェント」とエージェントプロンプトを含むルートを使用します。
エージェントは受信メールをどのように読むのでしょうか?
MIME は解析されません。 MailKite はエッジでメッセージをデコードし、 text 、 html 、解決された threadId 、署名付き URL としての添付ファイル、および認証結果を配信します。ハンドラーは Webhook 署名を検証し、プレーン フィールドを読み取ります。
エージェントにメールを操作させるのは危険ではないでしょうか?
これはプロンプトインジェクションサーフェスです。そうです、送信者は簡単になりすまし可能です。チェック

指示を信頼する前に認証 (SPF/DKIM/DMARC) の結果を確認し、システム プロンプトだけに依存しないでください。代理人の権限を拘束する。エージェントセキュリティの投稿 を参照してください。
エージェントは MailKite をツールとして呼び出すことができますか?
はい。 mcp.mailkite.dev にはホストされた MCP サーバー ( claude mcp add --transport http mailkite https://mcp.mailkite.dev/mcp ) と Claude Code プラグインがあるため、メールの送信、メッセージの読み取り、ドメインの管理はすべて、エージェントが直接行うことができるツール呼び出しです。
エージェントに受信箱を与えると、エージェントは電子メールで届く世界の一部の地域に耳を傾けなくなります。 MailKite にドメインを指定すると、数分以内に独自のメールを読んで応答するようになります。Bring-your-own ループまたはアクション「agent」( inbox-agents docs ) を含むルートを選択し、何かを実行させる前にセキュリティ投稿を読みます。
関連: 電子メールの受信が難しい理由に関する受信の柱 、Node で受信電子メールを JSON に解析する 、および AI エージェント ガイド 。
Webhook を使用して電子メールを受信します。 1 つの API で送信します。
メールサーバーを完全にスキップします。無制限のドメイン、無料 - 超過分は
優雅で、決してハードカットではありません。ライブ API キーを 5 分以内に取得します。
ドキュメントを読む
横並びの方がいいですか?
この投稿について話し合う: ハッカー ニュース · X で共有 · リンクをコピー
関連記事
同じ電子メール パーサーを 3 回作成しましたが、難しかったのは、それらが同じように失敗することでした

[切り捨てられた]

## Original Extract

Give an AI agent a real, scoped address on a domain you control. Inbound mail arrives as parsed JSON to an event.received loop and the agent replies over the Send API, or MailKite runs the agent for you on a route with action: agent. Working code, the security caveat, and the honest DIY alternatives
[truncated]

Give your AI agent its own email inbox — MailKite Beta
50% off your first year — founding rate for the first 1,000 customers.
View pricing →
MailKite
Features Receive Send For agents Pricing Docs Blog Contact Sign in · Get started Features Receive Send For agents Pricing Docs Blog Contact Sign in
All posts
June 10, 2026 · Gabe · 11 min read Give your AI agent its own email inbox
Give an AI agent a real, scoped address on a domain you control. Inbound mail arrives as parsed JSON to an event.received loop and the agent replies over the Send API, or MailKite runs the agent for you on a route with action: agent. Working code, the security caveat, and the honest DIY alternatives.
Copy as Markdown
Open in ChatGPT
Open in Claude
An AI agent with its own email inbox is an autonomous program that owns a real address on a domain you control, so it can receive verification codes, be handed work by email, and reply on its own without a human in the loop. This post is for a developer wiring an agent to an inbox for the first time, and it shows two ways to build one: run the agent yourself (inbound mail arrives as parsed JSON, your handler calls your model, the agent answers with one mk.send() ), or let MailKite (which we build) run it for you with a route whose action is agent . Either way there’s no IMAP, no MIME parsing, and no personal Gmail account quietly wired into a bot.
Here’s the bring-your-own-agent loop, whole. Email in, verify the signature, hand the body to your model, reply through the same client. It runs as pasted on Node 18+ ( npm install mailkite express ):
import express from "express" ;
import { MailKite } from "mailkite" ;
const app = express ();
const mk = new MailKite ( process . env . MAILKITE_API_KEY );
const SECRET = process . env . MAILKITE_WEBHOOK_SECRET ;
app .use ( "/hooks/agent" , express .raw ({ type : "application/json" }));
app .post ( "/hooks/agent" , async (req , res) => {
// signature check, replay window, constant-time compare — one call
if ( ! MailKite .verifyWebhook ( req .headers[ "x-mailkite-signature" ] , req .body , SECRET )) {
return res .sendStatus ( 401 );
}
res .sendStatus ( 200 ); // ack fast; run the agent out of band
const event = JSON .parse ( req .body);
if ( event .type !== "email.received" ) return ;
// Treat the body as untrusted INPUT, never as instructions (see the security section).
const answer = await runAgent ({
task : event .text ,
from : event . from .address ,
trusted : event . auth .spf === "pass" && event . auth .dmarc === "pass" ,
});
await mk .send ({
from : event .to[ 0 ].address , // reply from the address it was sent to
to : event . from .address ,
subject : `Re: ${ event .subject } ` ,
inReplyTo : event .id , // threads the reply
html : answer .html ,
});
});
app .listen ( 3000 );
That’s a fully autonomous email agent: it hears, it thinks, it answers. mk.send() returns { id, status } so you can log the outbound message, and the identical handler shape exists for Python, Ruby, Go, PHP, and Java; see the receiving docs and sending docs . Raw HTTP works too if you can’t take a dependency, but you’d be hand-rolling the HMAC verify that verifyWebhook does in one line; prefer the SDK.
What actually lands at your handler
MailKite decodes the message at the edge, so runAgent gets fields, not MIME. This is the same inbound webhook from the pillar: decoded text and html , a resolved threadId , attachments as short-lived signed URLs, and an auth block:
{
"id" : "msg_2Hk9…" ,
"type" : "email.received" ,
"from" : { "address" : "ada@example.com" } ,
"to" : [{ "address" : "agent@yourco.dev" }] ,
"subject" : "Re: invoice #1042" ,
"text" : "Looks good — approved!" ,
"html" : "<p>Looks good — approved!</p>" ,
"threadId" : "<a1b2c3@mail.example.com>" ,
"auth" : { "spf" : "pass" , "dkim" : "pass" , "dmarc" : "pass" , "spam" : "ham" } ,
"attachments" : [
{ "id" : "msg_2Hk9…:0" , "filename" : "po.pdf" , "contentType" : "application/pdf" ,
"size" : 18213 , "url" : "https://api.mailkite.dev/att/2Hk9…:0?exp=…&sig=…" }
]
}
The auth block tells the agent whether SPF, DKIM, and DMARC passed. Hold that thought; it’s load-bearing in a minute.
Two ways to run an agent inbox
You own the loop above, or you hand the whole thing to MailKite. Same inbound edge; the difference is where the agent’s turns execute.
Bring your own agent
email in
MX edge parse + auth
your webhook signed JSON
your agent your model
mk.send() reply out
Let MailKite run it
email in
MX edge parse + auth
agent route action: 'agent'
Queue agent run
reply send_email
Blue = operated by MailKite. Top lane: you write the webhook and the model loop. Bottom lane: the agent's turns run in our pipeline.
Two ways to run an agent inbox: bring your own agent (you own the loop) or a route with action: 'agent' (MailKite runs the loop on a durable queue). Same parsed inbound edge either way.
Let MailKite run the agent (route action ‘agent’)
If you don’t want to host the loop, make the route itself the agent. A route with action: "agent" carries an agentPrompt (free-text instructions), and MailKite runs the model loop for you on every inbound message:
await mk .createRoute ({
match : "support@yourco.dev" ,
action : "agent" ,
agentPrompt : "Answer billing questions from the docs. Escalate anything else to humans@yourco.dev." ,
});
The run doesn’t happen on the inbound request’s clock. We enqueue it on a Cloudflare Queue ( mailkite-agent-runs ) with its own execution budget, so ingest still returns fast and a slow model call can’t wedge the pipeline. Each run is capped at 8 tool rounds, aborted at 5 minutes by a cron reaper as a crash-proof backstop, and recorded as a full transcript you can drill into per route in the dashboard. The system prompt bakes in the safety rules the next section is about: the email body is untrusted, at most one reply, and never reply to no-reply or automated senders. The agent replies with an internal send_email tool (threaded via inReplyTo ), the same /v1/send your own loop calls. Details in the inbox-agents docs .
Bring-your-own wins when the agent’s brain is your code: your model, your tools, your state. The built-in route wins when you’d rather not run infrastructure and the job is “read this inbound mail, answer or escalate.”
The part I have to flag: inbound email is untrusted input
Here’s where I slow you down, because I walked into this hole myself. The moment your agent follows what an email says, that email body is a prompt-injection vector . From: is plain text, so anyone can forge the sender, then simply tell your agent what to do, and a naive loop obeys.
That’s why the auth block is in the payload and why the loop passes a trusted flag instead of blindly acting: you can at least see whether SPF and DMARC passed before you weight a sender’s instructions. But checking auth is necessary, not sufficient. You cannot prompt your way out of prompt injection; the real answer is architectural, bounding what a fooled agent can even do (owner-scoped tools, one reply, no acting on links). This is the same reasoning the built-in agent uses, and I wrote up the mistake honestly in Why aren’t we seeing more agent security discussions? . Read it before you point either loop at anything that matters.
When to build it yourself (and when not to)
An agent inbox isn’t exotic; you can assemble one without us, and sometimes you should. The honest alternatives:
IMAP polling on a real mailbox. Give the agent a Gmail/Workspace or Fastmail account and poll IMAP or the Gmail API. It works, but you own MIME parsing, threading, attachment decoding, and a mailbox that is really a person’s account with a person’s permissions. Best when the agent genuinely needs to live inside an existing human inbox.
Postmark inbound + your own loop. Postmark’s inbound parse also POSTs parsed JSON; point it at the same runAgent loop above. A fine call if you’re already on Postmark for sending and just want the inbound half.
Cloudflare Email Workers. The receiving primitive: an email hits a Worker’s email() handler. There’s no parsing, routing, or management layer; you build that. Good if you want to own the whole edge. (We build MailKite on Cloudflare, and this is the layer we put on top.)
Self-hosted Haraka. Run your own MX. Total control, total ops: the SMTP server, spam filtering, TLS, and uptime are all yours. Worth it at real scale or under strict data-residency rules.
The DIY shape is the same in every case: an MX record (or a mailbox), something that turns raw MIME into fields, signature and SPF/DKIM/DMARC checks, and a reply path. MailKite is that stack assembled (MX edge, parse, auth, a signed and retried webhook, and the Send API), so the 25 lines up top are the whole integration. Below a dedicated mail-infrastructure engineer’s time the trade tends to pay for itself; above it, or when the agent must live in a human’s real mailbox, one of the above is the better pick.
Let the agent use MailKite as a tool (MCP)
Everything so far is email reaching in . The other direction is your agent reaching out through email, and MailKite exposes its whole API as agent-native tools over a hosted MCP server. Point any MCP client at it:
claude mcp add --transport http mailkite https://mcp.mailkite.dev/mcp
After that the agent calls mailkite_send , mailkite_get_message , mailkite_list_domains , and the rest as first-class tool calls instead of hand-rolled HTTP. There’s also a Claude Code plugin that wires the same tools into your editor. An agent with an inbox and this tool set can run a support address end to end: read the incoming message, look something up, reply, all as tool calls.
Can I give an AI agent its own email address?
Yes. Point a domain at MailKite, pick an address like agent@yourco.dev , and set a webhook. Inbound mail is parsed to JSON and POSTed as an email.received event; the agent replies with mk.send() . It’s a real, scoped mailbox on your domain, not a personal account bolted onto a bot.
Should the agent run on my server or on MailKite?
Both work. Host the loop yourself ( webhook route → your endpoint → your model → Send API) when the agent’s logic is your code. Use a route with action: "agent" and an agentPrompt when you’d rather MailKite run the model loop on its own durable queue and hand you a transcript.
How does the agent read incoming email?
It doesn’t parse MIME. MailKite decodes the message at the edge and delivers text , html , a resolved threadId , attachments as signed URLs, and an auth result. Your handler verifies the webhook signature and reads plain fields.
Isn’t letting an agent act on email dangerous?
It’s a prompt-injection surface, yes: senders are trivially spoofable. Check the auth (SPF/DKIM/DMARC) results before trusting instructions, and don’t rely on the system prompt alone; bound the agent’s authority. See the agent-security post .
Can the agent call MailKite as a tool?
Yes. There’s a hosted MCP server at mcp.mailkite.dev ( claude mcp add --transport http mailkite https://mcp.mailkite.dev/mcp ) and a Claude Code plugin , so sending mail, reading messages, and managing domains are all tool calls the agent can make directly.
Give your agent an inbox and it stops being deaf to the parts of the world that arrive by email. Point a domain at MailKite and it’ll be reading and answering its own mail in a few minutes: pick the bring-your-own loop or a route with action: 'agent' ( inbox-agents docs ), then read the security post before you let it act on anything.
Related: the inbound pillar on why receiving email is hard , parsing inbound email to JSON in Node , and the AI agents guide .
Receive email with a webhook. Send with one API.
Skip the mail server entirely. Unlimited domains, free — and overage is
graceful, never a hard cut. Get a live API key in under five minutes.
Read the docs
Prefer a side-by-side?
Discuss this post: Hacker News · Share on X · Copy link
Related posts
I wrote the same email parser three times, and the hard part was making them fail identically

[truncated]
