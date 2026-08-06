---
source: "https://ddz.dev/blog/llm-subject-line-gmail-threading/"
hn_url: "https://news.ycombinator.com/item?id=49197346"
title: "Never let the LLM write the subject line"
article_title: "Never let the LLM write the subject line | ddz.dev LLC"
author: "dusandz"
captured_at: "2026-08-06T15:07:37Z"
capture_tool: "hn-digest"
hn_id: 49197346
score: 1
comments: 0
posted_at: "2026-08-06T14:40:52Z"
tags:
  - hacker-news
  - translated
---

# Never let the LLM write the subject line

- HN: [49197346](https://news.ycombinator.com/item?id=49197346)
- Source: [ddz.dev](https://ddz.dev/blog/llm-subject-line-gmail-threading/)
- Score: 1
- Comments: 0
- Posted: 2026-08-06T14:40:52Z

## Translation

タイトル: LLM に件名を書かせないでください
記事のタイトル: LLM に件名を書かせない | dz.dev LLC
説明: AI エージェントがすべてのリードに時間通りに返信し、Gmail は各返信を新しい会話としてファイルしました。参考資料と件名に関する Gmail スレッド。

記事本文:
コンテンツへスキップ DDZ.dev LLC ソフトウェア コンサルティング ホーム
LLM に件名を書かせないでください
修正 (TL;DR) 1. 参照と一致する件名に関する Gmail スレッド。両方必要です。 2. コード内の返信件名を導出します: Re: <original> 。決してモデルからではありません。 3. 受信メッセージ ID を送信まで運び、 In-Reply-To / References を設定します。 ✓ LLM が本体を書き込みます。配管は封筒を書きます。私は AI エージェントがすべての最初の行を処理する Web サイト スタジオ biro.works を構築しています。リードがメールを送信し、エージェントが返信の下書きを作成し、人間がそれを承認し、返信が送信されます。パイプラインは機能しました。リードは数分以内に回答を受け取り、口調は適切で、価格はカタログから直接得られました。
次に、クライアント側の Gmail でのテスト会話を調べました。クライアントは「パン屋のウェブサイトが必要です」と送ってきました。私たちの返信は「biro.works を使用した新しいベーカリーのウェブサイト」として届きました。まったく新しい会話が、返信されたメッセージとは何の関係もなく受信箱に残されています。それに返信して、別の新しい会話を始めましょう。 3通のメールが届いたが、スレッドは紙吹雪だった。
どのログにも壊れたものはありませんでした。メールは正常に送信され、正常に配信され、正常に読み取れました。彼らはただ会話を拒否しただけです。
スタック: SES は返信ドメイン宛てのメールを受信し、生のメッセージを S3 にダンプし、Lambda がそれを解析して正規化されたペイロードを取り込みサービスに POST します。これにより、返信を下書きして (承認後に) 送信する耐久性の高いワークフローが開始されます。受信ペイロードには、送信者、件名、本文、および RFC822 メッセージ ID のすべてが含まれていました。
バグ 1: ワークフローは件名または Message-ID を受信しませんでした。 Webhook は両方をデータベースに律儀に保存し、名前、電子メール、本文だけでワークフローを開始しました。そのため、返信には In-Reply-To ヘッダーも References ヘッダーもまったく含まれずに送信されました。グマイ

文字通り、それを結び付けるものが何もありませんでした。
バグ 2 はブログに投稿する価値のあるものです。返信のドラフトは JSON オブジェクトとしてクロードから提供され、私がそれに与えたスキーマは次のようになります。
{ "replySubject": string, "replyBody": string, ... } 問題がわかりましたか?モデルさんに件名を聞いてみました。モデルはモデルなので、書きました。良いものでもあります。フレンドリーで、具体的で、ブランドに沿ったもの。そして、プロンプトにはスレッドの本文のみが含まれていたため、保存するはずだった元の件名は一度も表示されませんでした。幻覚ではなかったのです。それは私がスキーマに入力した正確な質問に答えていました。
これを書き留める価値のある部分は次のとおりです。最初にヘッダーを修正して、これで完了だと思いました。電子メールのスレッドについてざっと読んだ内容はすべて References と In-Reply-To について述べているからです。これは RFC 5322 の話であり、Thunderbird や Mutt のようなクライアントはこれに従っています。
Gmail にはありません。 Gmail スレッドの参照ヘッダーと件名。同じ参考文献、異なる主題: 新しい会話。 Google 自身も、スレッドに関するドキュメントの細字で、References ヘッダーと In-Reply-To ヘッダーを修正し、独自の別個の要件として「Subject ヘッダーは一致する必要がある」と述べています。
そのため、完璧なヘッダーと独創的な件名を含む返信でも、Gmail と Google Workspace ユーザーごとにスレッドが分割されてしまいます。中小企業の対象ユーザーの場合、そのユーザーはほぼ全員です。ここではモデルがクリエイティブになることは許されていませんでした。誰もそんなことは言っていなかったし、誰も私にそんなことを教えてくれなかったのですから。
返信の件名は内容ではありません。それはプロトコルです。したがって、それはモデルの手から離れ、10 行の関数に移行しました。
/** 返信/転送プレフィックス (「Re:」、「RE:」、「Fwd:」、...) のスタックを削除します。 */
関数stripReplyPrefixes(件名:文字列):文字列{
返品対象。 replace ( / ^(\s*(re|fwd?|aw|sv)\s*(\[\d+\])?\s*:\s*)+ / i , "" ) 。

トリム ( ) ;
}
エクスポート関数 threadReplySubject (元の文字列 | null | 未定義、フォールバック文字列) : string {
const を取り除いた = オリジナル ?ストリップ返信プレフィックス (オリジナル) : "" ;
剥がして返す？ ` Re: ${ 削除された } ` : フォールバック ;
プレフィックスの削除は見た目以上に重要です。クライアントはプレフィックスをスタックするか ( Re: Re: Re: )、ローカライズするか (AW: ドイツの Outlook から、SV: スカンジナビアの Outlook から)、または番号を付けます ( Re[2]: )。単一の Re: に正規化すると、Gmail は満足です。
このモデルは引き続き ReplySubject を返しますが、フォールバックに降格されます。これは、リードが Web フォーム経由で到着した場合にのみ使用されます。この場合、保持する電子メール スレッドがなく、最初のメッセージの件名を誰かが作成する必要があります。電子メールとして届いたものについては、件名が導かれ、議論は終了します。
そして、メッセージ ID はパイプライン全体を通って送信されるようになります。 SES v2 では、これは単純なコンテンツ タイプのヘッダー フィールドです (私の @aws-sdk/client-sesv2 3.1057 よりずっと前からサポートされています)。
コンテンツ: {
シンプル: {
件名 : { データ : 件名 } 、
本文 : { テキスト : { データ : テキスト } 、 HTML : { データ : html } } 、
ヘッダー: [
{ 名前 : "In-Reply-To" 、値 :boundMessageId } 、
{ 名前 : "参照" 、値 :boundMessageId } 、
]、
} 、
もう 1 つ詳細を説明します。承認ゲートです。これは後で重要になります。人間は送信されるすべての返信を送信前にレビューしますが、レビュー画面には実際に送信されるものではなく、モデルの件名が表示されていました。件名はゲートの前で計算されるため、オペレーターは実際の電子メールを承認します。承認後に値が書き換えられた場合、承認はその値をカバーしていません。
LLM をパイプライン内に配置すると、出力スキーマのすべてのフィールドが小さな権限の委任になります。 ReplyBody は委任するのが妥当です。 ReplySubject はそうではありませんでした。件名は次のとおりです

返信の内容は散文ではなく、1 つの正しい値を持つプロトコル フィールドであり、テキスト ジェネレーターにそれを要求すると、値の代わりにテキストが得られることが保証されます。モデルは何も間違ったことはしていません。スキーマはそうなりました。
LLM を中心に電子メール パイプラインを構築しますか?詳細を送ってください、検討させていただきます。
前: Restate のデフォルトのノード名はコンテナーのホスト名です

## Original Extract

An AI agent replied to every lead on time, and Gmail filed each reply as a new conversation. Gmail threads on References AND the subject line.

Skip to content DDZ.dev LLC Software Consulting Home
Never let the LLM write the subject line
The Fix (TL;DR) 1. Gmail threads on References and a matching subject. You need both. 2. Derive the reply subject in code: Re: <original> , never from the model. 3. Carry the inbound Message-ID all the way to the send and set In-Reply-To / References . ✓ The LLM writes the body. The plumbing writes the envelope. I'm building biro.works , a website studio where AI agents handle the first line of everything: a lead emails us, an agent drafts the reply, a human approves it, the reply goes out. The pipeline worked. Leads got answers within minutes, the tone was right, the pricing came straight from the catalog.
Then I looked at a test conversation from the client's side, in Gmail. The client had sent "Need a website for my bakery." Our reply arrived as "Your new bakery website with biro.works" . A brand new conversation, sitting in the inbox with no connection to the message it answered. Reply to that, get another new conversation. Three emails in, the thread was confetti.
Nothing was broken in any log. The emails sent fine, delivered fine, read fine. They just refused to be a conversation.
The stack: SES receives mail for our reply domain, dumps the raw message to S3, a Lambda parses it and POSTs a normalized payload to the ingestion service, which kicks off a durable workflow that drafts and (after approval) sends the reply. The inbound payload had everything: sender, subject, body, and the RFC822 Message-ID .
Bug one: the workflow never received the subject or the Message-ID . The webhook dutifully stored both in the database, then started the workflow with just name, email, and body. So the reply went out with no In-Reply-To and no References header at all. Gmail had literally nothing to attach it to.
Bug two is the one worth a blog post. The reply draft comes from Claude as a JSON object, and the schema I gave it looked like this:
{ "replySubject": string, "replyBody": string, ... } See the problem? I asked the model for a subject line. The model, being a model, wrote one. A good one, even. Friendly, specific, on-brand. And since the prompt only contained the thread bodies, it had never seen the original subject it was supposed to preserve. It wasn't hallucinating. It was answering the exact question I put in the schema.
Here's the part that makes this worth writing down. I fixed the headers first and assumed I was done, because everything you skim about email threading talks about References and In-Reply-To . That's the RFC 5322 story, and clients like Thunderbird or Mutt live by it.
Gmail doesn't. Gmail threads on the reference headers and the subject line. Same References , different subject: new conversation. Google says so themselves, in the fine print of their threading documentation : correct References and In-Reply-To headers, and then, as its own separate requirement, "The Subject headers must match."
So a reply with perfect headers and a creative subject still splits the thread for every Gmail and Google Workspace user, which for a small-business audience is roughly everyone. The model wasn't allowed to be creative here. Nobody had told it that, because nobody had told me that.
The subject of a reply is not content. It's protocol. So it moved out of the model's hands and into a ten-line function:
/** Strip any stack of reply/forward prefixes ("Re:", "RE:", "Fwd:", ...). */
function stripReplyPrefixes ( subject : string ) : string {
return subject . replace ( / ^(\s*(re|fwd?|aw|sv)\s*(\[\d+\])?\s*:\s*)+ / i , "" ) . trim ( ) ;
}
export function threadReplySubject ( original : string | null | undefined , fallback : string ) : string {
const stripped = original ? stripReplyPrefixes ( original ) : "" ;
return stripped ? ` Re: ${ stripped } ` : fallback ;
} The prefix stripping matters more than it looks. Clients stack prefixes ( Re: Re: Re: ), localize them ( AW: from German Outlooks, SV: from Scandinavian ones), or number them ( Re[2]: ). Normalize to a single Re: and Gmail is happy.
The model still returns replySubject , but it's demoted to a fallback: it's only used when a lead arrives through the web form, where there's no email thread to preserve and someone has to write a subject for the first message. For anything that arrived as an email, the subject is derived, end of discussion.
And the Message-ID now rides through the whole pipeline into the send. With SES v2 that's the Headers field on the simple content type (supported since well before my @aws-sdk/client-sesv2 3.1057):
Content : {
Simple : {
Subject : { Data : subject } ,
Body : { Text : { Data : text } , Html : { Data : html } } ,
Headers : [
{ Name : "In-Reply-To" , Value : inboundMessageId } ,
{ Name : "References" , Value : inboundMessageId } ,
] ,
} ,
} , One more detail that would have bitten later: the approval gate. A human reviews every outgoing reply before it sends, and the review screen was showing the model's subject, not the one that would actually go out. The subject is now computed before the gate, so the operator approves the real email. If a value gets rewritten after approval, the approval didn't cover it.
When you put an LLM inside a pipeline, every field in your output schema is a small delegation of authority. replyBody was a reasonable thing to delegate. replySubject wasn't, because the subject line of a reply isn't prose, it's a protocol field with one correct value, and asking a text generator for it guarantees you'll get text instead of the value. The model did nothing wrong. The schema did.
Building an email pipeline around an LLM? Send me the details and I'll take a look.
Previous: Restate's default node name is your container hostname
