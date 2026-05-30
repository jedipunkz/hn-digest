---
source: "https://ably.com/docs/ai-transport/why/http-streaming-and-ai"
hn_url: "https://news.ycombinator.com/item?id=48324538"
title: "HTTP Streaming and AI"
article_title: "Ably Realtime | HTTP streaming and AI"
author: "zknill"
captured_at: "2026-05-30T11:45:17Z"
capture_tool: "hn-digest"
hn_id: 48324538
score: 2
comments: 0
posted_at: "2026-05-29T15:34:31Z"
tags:
  - hacker-news
  - translated
---

# HTTP Streaming and AI

- HN: [48324538](https://news.ycombinator.com/item?id=48324538)
- Source: [ably.com](https://ably.com/docs/ai-transport/why/http-streaming-and-ai)
- Score: 2
- Comments: 0
- Posted: 2026-05-29T15:34:31Z

## Translation

タイトル: HTTP ストリーミングと AI
記事のタイトル: Ably Realtime | HTTPストリーミングとAI
説明: AI のダイレクト HTTP ストリーミングに関する 4 つの運用上の問題: 切断時にストリームが失敗する、セッションがデバイスにまたがらない、クライアントがエージェントに到達できない、マルチ エージェントが複雑である。

記事本文:
アブリリアルタイム | HTTPストリーミングとAI
ドキュメント ドキュメントの例 Ask AI ログイン 無料で始める svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> プラットフォーム svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Ably Pub/Sub svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Ably Chat svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold text-neutral-1300 dark:text-neutral-000" data-radix-collection-item=""> Ably AI Transport svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> 概要 svg]:rotate-90rounded-lg text-neutral-1300 dark:text-neutral-000 font-bold" data-radix-collection-item=""> AI Transport を使用する理由 svg]:rotate-90 font-mediumrounded-lg border-l border-neutral-300 dark:border-neutral-1000 hover:border-neutral-500 dark:hover:border-neutral-800rounded-l-none" data-radix-collection-item=""> 概要 svg]:rotate-90rounded-lg border-l dark:border-neutral-1000 hover:border-neutral-500 dark:hover:border-neutral-800rounded-l-none text-neutral-1300 dark:text-neutral-000 font-bold border-orange-600 bg-orange-100 hover:bg-orange-100" data-radix-collection-item=""> HTTP ストリーミングと AI svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> 概念 svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> はじめに svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> フレームワーク svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> 機能

es svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> 本番環境への移行 svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> API リファレンス svg]:rotate-90 font-mediumrounded-lg
[切り捨てられた]
直接 HTTP ストリーミングは、1 回限りの対話には問題ありませんが、それ以外の場合は失敗します。これらは、AI アプリが運用開始されると現れる 4 つの制限です。
ほとんどの AI フレームワークは、単純なクライアント主導の対話をサポートしています。つまり、クライアントが HTTP リクエストを作成し、エージェントがそれを処理し、サーバー送信イベントまたは同様の HTTP ストリームを介して応答ストリームがクライアントに返されます。このパターンはシンプルで、ワンショットのインタラクションには驚くほど効果的で、すべてのフレームワークがサポートしています。パターンの単純さは、その限界の原因でもあります。
以下の制限は、クライアントとエージェントの対話を、それを運ぶトランスポートに結合することで生じます。接続、リクエスト、およびストリーミングされた応答はすべて同じ存続期間であり、1 つのクライアントと 1 つのエージェント間の 1 つの対話の間存在します。接続が存続するまで対話を必要とする (またはその 1 つのクライアント以外から見えるようにする) には、その上に新しいインフラストラクチャを構築する必要があります。
応答ストリームの操作は、基礎となる接続の健全性と結びついています。接続が切断されると、応答ストリームは失敗します。
これは日常的に起こります。電話機が Wi-Fi から携帯電話に切り替わります。ユーザーがページを更新します。応答中にラップトップの蓋が閉まります。 LLM はトークンを生成し続けますが、トークンを配信する場所がありません。
SSE は、ほとんどの AI フレームワークのデフォルトのストリーミング トランスポートです。 SSE プロトコルには、再接続するクライアントがストリーム内の再開する位置を指定するためのメカニズムが含まれています。実際には、サポートされることはほとんどありません。

バックエンドの複雑さ。 SSE ストリームを再開するには、順序付けのためにトークン イベントにシーケンス番号を割り当て、それらのイベントを外部ストアにバッファリングし、再開リクエストを処理するための新しい HTTP エンドポイントを追加します。これは、ステートレスなリクエスト ハンドラーからの大幅な逸脱です。作業が完了したとしても、再開は既存のクライアントの再接続のみを対象とします。 SSE にはセッション ID の概念が組み込まれていないため、ページ更新後の継続性はカバーされていません。その上にさらにもう一つの層がある建物。
HTTP ストリーミングの場合、接続は要求元のクライアントとそれを処理したエージェントのみに限定されます。 2 番目のタブや電話はそのストリームに入ることができません。これは、リクエストを開始したクライアントに対してのみ存在します。
実際には、ユーザーは常にサーフェス間を移動します。 2 番目のブラウザー タブ。モバイルアプリ。後で別のデバイスから会話を再開します。セッションへの共有アクセスがなければ、各サーフェスは分離されます。新しいクライアントが進行中のストリームを確認する方法はありません。そして、会話の履歴や現在の状態を共有します。
クライアントがエージェントに接続できない
クライアントによって開始された SSE リクエストは、サーバーからクライアントへの一方向です。最初のリクエストが行われた後、クライアントは同じ接続を介してエージェントにシグナルを送信する方法がありません。クライアントにある唯一のオプションは、ストリームを最後まで読み取るか、接続を閉じてストリームをキャンセルすることです。
キャンセルを唯一のアップストリーム信号として使用すると、根本的な矛盾が生じます。進行中のストリームをキャンセルする停止ボタンについて考えてみましょう。実装では、閉じられた接続の 2 つの解釈のどちらかを選択する必要があります。キャンセル (この場合、LLM は停止する必要があります)、または切断 (この場合、LLM はストリームを再開できるように続行する必要があります) です。曖昧さを解消する方法はありません。
bでも

WebSocket のような双方向トランスポートでは、接続は依然として 1 つのクライアントと 1 つのエージェント間の排他的なパイプです。他のデバイスにはアップストリーム チャネルがないため、2 番目のデバイスに割り込んだり、そこから操作したりすることはできません。
マルチエージェントのアーキテクチャは複雑です
マルチエージェント システムでは、オーケストレーターがクライアントの接続を処理し、専門のサブエージェントに作業を委任します。クライアントとオーケストレーター間の接続が排他的でポイントツーポイントである場合、サブエージェントとのすべての対話はオーケストレーターによってプロキシされる必要があります。ユーザーが中間の進行状況やサブエージェントからの応答を確認する必要がある場合、すべての更新がオーケストレーターによって仲介されるため、複雑さと結合が追加されます。
AI Transport を選択する理由: 耐久性のあるセッション レイヤーがこれらの各問題をどのように解決するか。
セッション : 一時的な接続に代わる永続的な共有会話状態。

## Original Extract

The four production problems with direct HTTP streaming for AI: streams fail on disconnect, sessions do not span devices, clients cannot reach the agent, multi-agent is complex.

Ably Realtime | HTTP streaming and AI
Docs Documentation Examples Ask AI Login Start free svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Platform svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Ably Pub/Sub svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Ably Chat svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold text-neutral-1300 dark:text-neutral-000" data-radix-collection-item=""> Ably AI Transport svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Overview svg]:rotate-90 rounded-lg text-neutral-1300 dark:text-neutral-000 font-bold" data-radix-collection-item=""> Why AI Transport svg]:rotate-90 font-medium rounded-lg border-l border-neutral-300 dark:border-neutral-1000 hover:border-neutral-500 dark:hover:border-neutral-800 rounded-l-none" data-radix-collection-item=""> Overview svg]:rotate-90 rounded-lg border-l dark:border-neutral-1000 hover:border-neutral-500 dark:hover:border-neutral-800 rounded-l-none text-neutral-1300 dark:text-neutral-000 font-bold border-orange-600 bg-orange-100 hover:bg-orange-100" data-radix-collection-item=""> HTTP streaming and AI svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Concepts svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Getting started svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Frameworks svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Features svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Going to production svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> API reference svg]:rotate-90 font-medium rounded-lg
[truncated]
Direct HTTP streaming is fine for one-off interactions and breaks down everywhere else. These are the four limitations that show up once an AI app is in production.
Most AI frameworks support a simple client-driven interaction: the client makes an HTTP request, an agent handles it, and the response streams back to the client over Server-Sent Events or a similar HTTP stream. The pattern is simple, surprisingly effective for one-shot interactions, and every framework supports it. The simplicity of the pattern is also the source of its limitations.
The limitations below arise from coupling the client-to-agent interaction to the transport that carries it. The connection, the request, and the streamed response are all the same lifetime: they exist for one interaction, between one client and one agent. Anything that requires the interaction to outlive the connection (or be visible to anything other than that one client) requires building new infrastructure on top.
The operation of a response stream is tied to the health of the underlying connection. When the connection drops, the response stream fails.
This happens routinely. A phone switches from Wi-Fi to cellular. A user refreshes the page. A laptop lid closes mid-response. The LLM continues to generate tokens, and there is nowhere to deliver them.
SSE is the default streaming transport for most AI frameworks. The SSE protocol does include a mechanism for a reconnecting client to specify a position in the stream to resume from. In practice it is rarely supported, because supporting it adds significant backend complexity. To resume an SSE stream you assign sequence numbers to token events for ordering, buffer those events in an external store, and add a new HTTP endpoint to handle resume requests. That is a substantial departure from a stateless request handler. Even with the work done, resume only covers reconnection of an existing client; it does not cover continuity after a page refresh, because SSE has no built-in concept of session identity. Building that is yet another layer on top.
With HTTP streaming, the connection is exclusive to the requesting client and the agent that handled it. A second tab or a phone has no way into that stream. It only exists for the client that initiated the request.
In reality, users move between surfaces constantly. A second browser tab. The mobile app. Picking the conversation up later from a different device. Without shared access to the session, each surface is isolated. There is no way for a new client to see the in-progress stream. And sharing the conversation history, or its current state.
Clients cannot reach the agent
An SSE request initiated by the client is one-way: server to client. The client has no way to send a signal to the agent through the same connection once the initial request has been made. The only options the client has are to read the stream to completion or to cancel it by closing the connection.
Using cancellation as the sole upstream signal creates a fundamental conflict. Consider a stop button that cancels an in-progress stream. The implementation has to choose between two interpretations of a closed connection: either it is a cancel (in which case the LLM should stop), or it is a disconnect (in which case the LLM should keep going so the stream can resume). There is no way to disambiguate.
Even with a bidirectional transport like WebSocket, the connection is still an exclusive pipe between one client and one agent. Other devices have no upstream channel, so they cannot interrupt or steer from a second device.
Multi-agent architectures are complex
In multi-agent systems, an orchestrator handles the client's connection and delegates work to specialised sub-agents. When the connection between the client and the orchestrator is exclusive and point-to-point, every interaction with a sub-agent has to be proxied by the orchestrator. If users need to see intermediate progress or responses from sub-agents, every update is mediated by the orchestrator, adding complexity and coupling.
Why AI Transport : how a durable session layer solves each of these problems.
Sessions : the persistent, shared conversation state that replaces the ephemeral connection.
