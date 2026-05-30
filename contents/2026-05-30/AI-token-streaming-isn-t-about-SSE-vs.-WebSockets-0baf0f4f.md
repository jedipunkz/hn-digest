---
source: "https://zknill.io/posts/ai-token-streaming-isnt-about-sse-vs-websockets/"
hn_url: "https://news.ycombinator.com/item?id=48324678"
title: "AI token streaming isn't about SSE vs. WebSockets"
article_title: "AI token streaming isn't about SSE vs WebSockets — /dev/knill"
author: "zknill"
captured_at: "2026-05-30T11:45:04Z"
capture_tool: "hn-digest"
hn_id: 48324678
score: 9
comments: 0
posted_at: "2026-05-29T15:44:26Z"
tags:
  - hacker-news
  - translated
---

# AI token streaming isn't about SSE vs. WebSockets

- HN: [48324678](https://news.ycombinator.com/item?id=48324678)
- Source: [zknill.io](https://zknill.io/posts/ai-token-streaming-isnt-about-sse-vs-websockets/)
- Score: 9
- Comments: 0
- Posted: 2026-05-29T15:44:26Z

## Translation

タイトル: AI トークン ストリーミングは SSE と WebSocket に関するものではありません
記事のタイトル: AI トークン ストリーミングは SSE と WebSocket に関するものではありません — /dev/knill
説明: それ

記事本文:
👾 "> /dev/ knill 投稿プロジェクトについての話 2026 年 5 月 21 日 · 8 分 · #ai AI トークン ストリーミングは SSE と WebSocket に関するものではありません
それは、何がデモで機能し、何が実稼働で機能するか、そしてどれだけ構築して維持する必要があるかについてです。
Ably では、実稼働トークンのストリーミングを解決しているため、その必要はありません。難しい部分は SSE ではありません
またはWebSocket。
エージェント コーディング ツールまたはチャットボットに「本番環境のクライアントに AI トークンをストリーミングする方法」を質問し、
SSE と WebSocket に関する回答のセクションが表示されます。
しかし、それは質問でも、実際の答えでもありません。
トランスポートとして SSE と WebSocket を使用する場合を純粋に比較すると、SSE がより簡単な選択になります。
また、ほとんどのユースケースでは、これがより良い選択です。本番トークン用に構築する必要があるアーキテクチャ
ストリーミングは下の図のようになります。 「プロンプト」リクエストと「レスポンス」を分離する
ストリーム、および再開と再接続を可能にするためにトークンを保存するためのトークン キャッシュ/データ ストア。
フローチャート LR
H[人間] -->|1。即時リクエスト| S[POST /messages] --> L[llm]:::accent
L -.->|sse トークン| S
DB[(トークンキャッシュ)]
S -.->|ストアトークン| DB
H -.->|2.ストリーム応答|SS
SS[GET /streams/:id]
SS -.->|トークンの読み取り| DB
classDef アクセントストローク:#7eb8a4、ストローク幅:1.5px
LLM 応答トークンは、何らかのデータストアを介してスレッド化されます。リクエストとレスポンスは分離されています。 WebSocket のバージョンは、クライアントが WebSocket 接続を開く点を除いて、ほぼ同じに見えます。
プロンプトをメッセージとして送信すると、サーバーは同じメッセージ上でトークン メッセージで応答します。
接続。再開と再接続にはトークン キャッシュ/データ ストアが引き続き必要です。
フローチャート LR
H[人間] -->|WebSocket を開く| S[サーバー] --> L[llm]:::アクセント
L -.->|sse トークン| S
DB[(トークンキャッシュ)]
S -.->|ストアトークン| DB
S -.->|トークンの読み取り| DB
S -.->|ストレ

WebSocket 上のトークンです| H
classDef アクセントストローク:#7eb8a4、ストローク幅:1.5px
SSE の設計と同じですが、WebSocket 接続を使用します。これが SSE で機能する理由
ほとんどの人々のシステム設計は、サーバーはステートレスであり、すべての状態はステートレスであるという考えに基づいています。
データベースに保存されます。これにより、サーバーは水平方向に拡張して、より多くのリクエストを処理できるようになります。
どのサーバーでもあらゆるリクエストを処理できるため、より優れたスケーリングが可能になります。通常、ロードバランサーは
リクエストをルーティングするサーバーの前。ほとんどの場合、ロード バランサーは以下に基づいています。
あらゆる種類のスティッキー セッションやセッション アフィニティではなく、サーバー間で負荷を共有します。
結局のところ、SSE はこのアーキテクチャに非常にうまく組み込まれています。クライアントは、
プロンプトを表示してサーバーにアクセスし、ストリーム ID を取得します。次に、クライアントはそのストリームに接続し、
トークンストリーミング応答。
どのサーバーでも元のプロンプト要求を処理でき、どのサーバーでもストリーム応答を処理できます。
トークンはデータベースを介してスレッドされるためです。
WebSocket は基本的にまったく同じですが、クライアントが WebSocket 接続を開き、
接続上のメッセージとしてプロンプトを表示します。応答トークンは同じ接続上で送り返されます。
接続の存続期間は長く、接続用に独自のメッセージ/リクエスト フレーミング レイヤーを構築する必要があります。
WebSocket 上で送受信されるメッセージの形式。そのため、WebSocket の構築はより複雑になり、
このアーキテクチャでは何も価値を追加することはありません。
したがって、応答トークンをストリーミングするために SSE を使用するのは明らかな選択のように思えます。
運用環境の展開には他にどのような機能が必要ですか?
これまでのところ、「ハッピー パス」でクライアントにトークンをストリーミングするための設計ができています。何も行かない場所
間違っています。私たちは可能な f をサポートするアーキテクチャを持っています

障害のケースとその回復方法
それらの機能はまだありません。では、これらの機能を実現するには他に何をする必要がありますか
可能ですか？
再接続、再開、リカバリ
クライアントは、接続が切断された場合にストリームに再接続し、ストリームから再開できる必要があります。
クライアントに送信された最後のトークン。ここでトークン キャッシュ/データ ストアが登場します。
再接続するクライアントが以前の場所から取得できるように、サーバーはトークンを保存する必要があります。しかし、
クライアントは、切断前にどこに到達したかを示すことができる必要もあります。
サーバー上の特定の位置と相関付けるため。
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
{
"タイプ" : "テキストデルタ" ,
"デルタ" : "こんにちは" ,
「位置」：1
}
{
"タイプ" : "テキストデルタ" ,
"デルタ" : "ワールド" 、
「位置」：2
}
{
"タイプ" : "テキストデルタ" ,
「デルタ」：「！」 、
「位置」：3
すべてのトークンにはストリーム内のシーケンス ID が必要です。これにより、クライアントが再接続するときに次のようになります。
「最後に取得したトークンは位置 2 にありました」とすると、サーバーは位置から始まるトークンを送信できるようになります。
3.
SSE 仕様には、この目的のために Last-Event-ID ヘッダーが組み込まれていますが、それでもビルドする必要があります。
サーバー側とクライアント側でそれをサポートするための配管。
サーバーとクライアントは、接続が切断されたことを検出できる必要があります。
再接続を試みることができます。これを行うには、サーバーはハートビート メッセージを頻繁に送信する必要があります。
(例: 10 秒ごと) 接続を維持し、クライアントが切断されたかどうかを検出します。の
クライアントには、メッセージを考慮する前にメッセージを待機する時間のタイムアウトも必要です。
接続が切断され、再接続を試行します。
ユーザーがストリームを待ちたくない場合、クライアントはストリームをキャンセルできる必要があります。
もう反応。チャットボット インターフェイスの多くには、

opボタンでストリーミングを中断します
現在の応答。クライアントは、指定されたストリーム ID に対して、サーバーに対してキャンセル要求を行うことができます。
彼らはキャンセルしたいと考えています。その後、サーバーはそれに関連する応答生成を停止する必要があります。
ストリーム。
これを行う簡単な方法は、キャンセルをトークン キャッシュに置き、読み取りをトークン キャッシュから行うことです。
ストリームがキャンセルされたことはわかっていますが、トークン キャッシュへの書き込み者はそこから読み取る必要もあります
書き込みの前に、新しいキャンセル信号があるかどうかを確認します。これはかなり厄介になります
トークン キャッシュが単純なキーと値のストアではなく、実際にはある種のキューである場合
リレーショナルデータベース。
トークンのロールアップが重要になるのは、再接続時と完了した応答の 2 つの場所です。
完了した応答ケースが最も簡単です。
サーバーが応答のストリーミングを完了したら、それらのトークンを完全なデータに圧縮する必要があります。
応答が生成され、その完全な応答としてどこかに保存されます。以前のメッセージに対する後の「履歴」リクエスト
チャットボットの会話では、その応答をストリーミングするのではなく、圧縮された完全な応答を提供する必要があります
トークンごとに。
より困難なケースは再接続時です。クライアントが再接続し、10、100、または 1000 個のトークンが失われた場合、
これらのトークンは、クライアントに 1 つに送信できる圧縮された応答にまとめられる必要があります。
クライアントにこれらのトークンを 1 つずつ消費するように強制するのは意味がありません。
トークンキャッシュに保存されます。繰り返しますが、そのキャッシュが実際に何らかの種類のキャッシュである場合、これはさらに困難になります。
キュー;特定のストリームのキューにいくつのトークンがあるかが常にわかるとは限らないためです。
人間のユーザーが別のデバイスで取得できる必要があるか、別のユーザーが別のデバイスで取得できる必要があります。
既存のチャットに参加します。ここではマルチデバイスの方が実際にはより単純なケースです。
独身

そのチャット会話を操作する人間は、新しいデバイスに切り替えるとデバイスが取得できるようになります。
履歴を確認すると、そのデバイスは、ストリーム応答が行われた後に新しいストリーム応答を取得する必要があることを認識します。
新しいプロンプトリクエスト。
マルチユーザーの場合はさらに困難です。 2 人のユーザーが両方ともチャット会話に接続している場合、
ユーザー A がそのチャットでプロンプト リクエストを行うと、ユーザー A はストリームに接続する必要があることがわかります。
そのプロンプト要求に対して応答トークンを受け取ります。しかし、同じチャットに参加しているユーザー B は、
そのストリームに接続することを知っているため、応答トークンを取得できません。これは次のようにすることで解決できます
会話とストリーム応答は同じエンティティです。つまり、ユーザーが最初に接続したとき、
実際にそのチャットの応答ストリームに接続しているチャット。そしてすべてストリーミングされました
そのチャットで送信されたプロンプト リクエストに対する応答は、チャット全体の応答を通じてストリーミングされます。
ストリーム。これにより、両方のユーザーが応答にはアクセスできるようになりますが、リクエストにはアクセスできなくなります。解決するには
リクエストについては、チャットでも同じ応答を行うようにプロンプト リクエストを行う必要があります。
ストリームを使用して、新しいプロンプトが同じチャット内の他のユーザーと共有されるようにします。
デモと本番環境の間には大きなギャップがある
構築する必要があるアーキテクチャと、サポートする必要がある機能に注目すると、
トークン ストリーミングの実稼働デプロイメントを実行すると、単純なトークン ストリーミングとの間には大きなギャップがあることがわかります。
github やブログ投稿で見られるデモと、実際に行う必要があること。
ほとんどの人は、次の 2 つの陣営に当てはまります。a) 彼らは、社会を構築していないため、このギャップについて知りません。
事前にトークン ストリーミングを本番環境に展開する、または b) ギャップについては知っているが、
それを埋めるのにどれだけの労力がかかるかを過小評価してください。
私の本業は AI の問題を解決することなので、このギャップについては承知しており、気にしています。

オーケンのストリーミングの問題。私たちは
これらの AI トークン ストリーミングの問題は、自動的にサポートするパブリッシュ/サブスクライブ製品で解決されました。
マルチユーザー、マルチデバイス、再接続、再開、キャンセル、トークンのロールアップ、履歴とトークン
圧縮。これは、リアルタイム パブリッシュ/サブスクライブ システムに基づいた製品です。そしてあなたがしなければならないのは次のことだけです:

## Original Extract

It

👾 "> /dev/ knill posts projects talks about May 21, 2026 · 8 min · #ai AI token streaming isn't about SSE vs WebSockets
It's about what works in demos and what works in production, and how much you have to build and maintain.
At Ably , we’ve solved production token streaming, so you don’t have to. And the hard-part isn’t SSE
or WebSockets.
Ask an agentic coding tool or chatbot “how to stream AI tokens to a client in production” and
it’ll give you a section of the answer on SSE vs WebSockets.
But that’s not the question, or really the answer.
In a pure comparison of using SSE or WebSockets as the transport, SSE is the simpler choice, and
is also the better choice for most usecases. The architecture you should build for production token
streaming looks like the diagram below. It’s got separation of ‘prompt’ request and ‘response’
stream, and a token cache/data store for storing the tokens in allowing for resume and reconnection.
flowchart LR
H[human] -->|1. prompt request| S[POST /messages] --> L[llm]:::accent
L -.->|sse tokens| S
DB[(token cache)]
S -.->|store tokens| DB
H -.->|2. stream repsonse|SS
SS[GET /streams/:id]
SS -.->|read tokens| DB
classDef accent stroke:#7eb8a4,stroke-width:1.5px
The LLM response tokens are threaded through some datastore. The request and response are separated. The WebSockets version looks almost exactly the same, except the client opens a WebSocket connection
and sends the prompt as a message, and the server responds with token messages on the same
connection. The token cache/data store is still needed for resume and reconnection.
flowchart LR
H[human] -->|open websocket| S[Server] --> L[llm]:::accent
L -.->|sse tokens| S
DB[(token cache)]
S -.->|store tokens| DB
S -.->|read tokens| DB
S -.->|stream tokens on websocket| H
classDef accent stroke:#7eb8a4,stroke-width:1.5px
The same as the sse design, but with a websocket connection Why this works with SSE
Most peoples system design is based around the idea that servers are stateless, and all the state is
stored in a database. This allows the servers to horizontally scale to handle more requests, and
allows for better scaling as any server can handle any request. There’s generally a load balancer in
front of the servers that routes requests to them. Mostly that load balancer will be based on
sharing the load across servers rather than any kind of sticky session or session affinity.
Turns out, SSE drops into this architecture really nicely. The client makes a POST request to the
server with a prompt, and gets a stream ID back. The client then connects to that stream and gets
the token streaming response.
Any server can handle the original prompt request, and any server can handle the stream response
because tokens are threaded through the database.
WebSockets is basically exactly the same, except the client opens a WebSocket connection and sends
the prompt as a message on the connection. The response tokens are sent back on the same connection.
The connection is longer-lived, and you have to build your own message/request framing layer for the
shape of messages sent back and forth on the WebSocket. So WebSockets are more complex to build and
maintain, and don’t really add any value in this architecture.
So it seems like an obvious choice to go with SSE for streaming your response tokens.
What other features do you need in a production deployment?
So far we’ve got a design for streaming tokens to the client in the ‘happy path’; where nothing goes
wrong. We’ve got the architecture to support the possible failure cases and how to recover from
them, but we don’t have the features yet. So what else do you need to do to make those features
possible?
Reconnection, resume, and recovery
The client needs to be able to reconnect to the stream if the connection drops, and resume from the
last token that was sent to the client. This is where the token cache/data store comes in. The
server needs to store the tokens so that a reconnecting client can pick up where it was before. But
the client also needs to be able to indicate where it got to before the disconnect, and that needs
to correlate with some position on the server.
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
{
"type" : "text-delta" ,
"delta" : "hello" ,
"position" : 1
}
{
"type" : "text-delta" ,
"delta" : " world" ,
"position" : 2
}
{
"type" : "text-delta" ,
"delta" : "!" ,
"position" : 3
} All the tokens need to a sequence ID in the stream, so that when the client reconnects it can say
“the last token I got was at position 2”, and the server can then send tokens starting at position
3.
The SSE spec has Last-Event-ID header built into it for this purpose, but you still need to build
the plumbing to support it on the server and client side.
The server and client need to be able to detect when the connection has dropped, so that the client
can attempt to reconnect. To do this, the server has to send a heartbeat message every so often
(e.g. every 10 seconds) to keep the connection alive and to detect if the client has dropped. The
client also needs to have a timeout for how long it will wait for a message before it considers the
connection dropped and attempts to reconnect.
The client needs to be able to cancel the stream if the user decides they don’t want to wait for the
response anymore. Lots of chatbot interfaces have a stop button that will interrupt the streaming of
the current response. The client can make a cancel request to the server, for the stream ID that
they want to cancel. The server then needs to stop any response generation associated with that
stream.
The easy way to do this is to put the cancel in the token cache, and readers from the token cache
know that the stream is cancelled, but writers to the token cache now need to also read from it
before write in order to check if there’s a new cancel signal in there. This becomes pretty awkward
if the token cache is actually some kind of queue, rather than a simple key-value store of
relational database.
There are two places where token rollup matter: on reconnect, and on completed response.
The completed response case is the easiest.
Once the server has finished streaming the response, those tokens need to be compacted into the full
response and stored somewhere as that full response. Later ‘history’ requests for earlier messages
in the chatbot conversation should serve the compacted full response and not stream that response
token-by-token.
The harder case is on reconnect. If the client reconnects and has missed 10, 100, or 1000 tokens,
those tokens need to be rolled up into a compacted response that can be sent to the client in one.
There’s no point forcing the client to consume those tokens one-by-one, just because that’s how they
are stored in the token cache. Again, this becomes harder if that cache is actually some kind of
queue; because you won’t always know how many tokens are in the queue for a given stream.
The human user needs to be able to pick up on another device, or another user needs to be able to
join the existing chat. Multi-device is actually the simpler case here, because there is still a
single human operating that chat conversation, when they switch to a new device the device can fetch
the history, and then it is that device knowing it needs to get a new stream response after it made
a new prompt request.
The mult-user case is much harder. If you’ve got two users both connected to a chat conversation,
and user A makes a prompt request in that chat, user A will know they need to connect to the stream
for that prompt request and receive the response tokens. But user B, who is also in that chat won’t
know to connect to that stream, and won’t get the response tokens. You could solve this by making
the conversation and the stream response the same entity. That is, when the users first connect to
the chat they are actually connecting to the response stream for that chat. And all streamed
responses to any prompt requests submitted in that chat are streamed over the chat-wide response
stream. This allows both users access to the responses but not the requests. To solve the
requests, well now you need to make prompt requests for the chat also go over the same response
stream so that new prompts are shared with other users in the same chat.
There’s a massive gap between demos and production
When you look at the architecture that you need to build, and the features you need to support for a
production deployment of token streaming, you can see that there’s a massive gap between the simple
demos you see in github and blog posts, and what you actually need to do.
Most folks fall into two camps, either a) they don’t know about this gap because they’ve not built a
production deployment of token streaming before, or b) they know about the gap but they
underestimate how much work it is to fill it.
I know and care about this gap because my day-job is solving AI token streaming problems. We’ve
solved these AI token streaming problems with a pub/sub product that automatically supports
multi-user, multi-device, reconnect, resume, cancellation, token rollup, history and token
compaction. It’s a product based on our realtime pub/sub system; and all you have to do is:
