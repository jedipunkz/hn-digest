---
source: "https://ably.com/docs/ai-transport/why"
hn_url: "https://news.ycombinator.com/item?id=48324690"
title: "Why AI Transport"
article_title: "Ably Realtime | Why AI Transport"
author: "zknill"
captured_at: "2026-05-30T11:44:56Z"
capture_tool: "hn-digest"
hn_id: 48324690
score: 4
comments: 0
posted_at: "2026-05-29T15:45:11Z"
tags:
  - hacker-news
  - translated
---

# Why AI Transport

- HN: [48324690](https://news.ycombinator.com/item?id=48324690)
- Source: [ably.com](https://ably.com/docs/ai-transport/why)
- Score: 4
- Comments: 0
- Posted: 2026-05-29T15:45:11Z

## Translation

タイトル: AI 輸送を行う理由
記事のタイトル: Ably Realtime | AI 輸送を選ぶ理由
説明: AI Transport が存在する理由: 本番環境では直接 HTTP ストリーミングが失敗します。耐久性のあるセッション層により、再開、マルチデバイス、双方向制御が修正されます。

記事本文:
アブリリアルタイム | AI 輸送を選ぶ理由
ドキュメント ドキュメントの例 Ask AI ログイン 無料で始める svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> プラットフォーム svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Ably Pub/Sub svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Ably Chat svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold text-neutral-1300 dark:text-neutral-000" data-radix-collection-item=""> Ably AI Transport svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> 概要 svg]:rotate-90rounded-lg text-neutral-1300 dark:text-neutral-000 font-bold" data-radix-collection-item=""> AI Transport を使用する理由 svg]:rotate-90rounded-lg border-l dark:border-neutral-1000 hover:border-neutral-500 dark:hover:border-neutral-800rounded-l-none text-neutral-1300 dark:text-neutral-000 font-bold border-orange-600 bg-orange-100 hover:bg-orange-100" data-radix-collection-item=""> 概要svg]:rotate-90 font-mediumrounded-lg border-l border-neutral-300 dark:border-neutral-1000 hover:border-neutral-500 dark:hover:border-neutral-800rounded-l-none" data-radix-collection-item=""> HTTP ストリーミングと AI svg]:rotate-90 font-medium Rounded-lg" data-radix-collection-item=""> 概念 svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> はじめに svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> フレームワーク svg]:rotate-90 font-mediumrounded-lg" data-radix-collection-item=""> 機能 sv

g]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Going to production svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> API reference svg]:rotate-90 font-medium rounded-lg" dat
[切り捨てられた]
AI アプリが運用開始されると、直接 HTTP ストリーミングは機能しなくなります。 AI Transport は、単一の HTTP 接続を、あらゆる参加者が接続できる永続的なセッションに置き換えます。
ほとんどの AI フレームワークは、HTTP リクエストとストリーミングされた応答を介してクライアントをエージェントに接続します。パターンはシンプルで、すべてのフレームワークがサポートしており、1 回限りのインタラクションやデモで機能します。また、最初のやり取りの後に行われるすべてのことも制限されます。
ストリームは接続とともに停止します。セッションは複数のデバイスにまたがることはできません。クライアントには、元のリクエストの処理中にエージェントに再連絡するメカニズムがありません。これらはそれぞれ実際の運用上の問題です。これらは共に、直接 HTTP ストリーミング上で構築される AI エクスペリエンスの品質を制限します。詳細については、「HTTP ストリーミングと AI」を参照してください。
永続的なセッションがモデルを変える
これらの問題には根本原因が共通しています。トランスポートはインタラクションに結合されます。接続、リクエスト、ストリーミングされた応答は一時的なものです。これらは 1 つの対話、1 つのクライアントと 1 つのエージェントに対して存在します。 HTTP 接続が失われると、クライアントとエージェント間の対話も失われます。
The shape that engineering teams are adopting is the durable session: a shared, persistent medium that any client or agent connects to, instead of a single HTTP request between one client and one agent.
永続セッションは、直接 HTTP ストリーミングにはない 3 つの機能を提供します。
弾力性のある配信。ストリームは、接続の切断、デバイスの切り替え、ページの更新、プロセスの再起動の後も存続します。クライアントは既知の位置から再開します。エージェントは cl に関係なく公開を続行します

優れた接続性。イベントが失われることも、イベントが複製されることもありません。
サーフェス間の連続性。セッションは接続ではなくユーザーに従います。 2 番目のタブを開いて電話に切り替え、数時間後に再びアクセスしてください。すべてのサーフェスは同じセッション状態を認識します。セッション識別子を持つクライアントはすべて、同じ会話状態をアタッチしてハイドレートします。
ライブコントロール。クライアントは、作業の進行中にエージェントと通信し、その処理を制御するための双方向ストリームを備えています。別のデバイスからの世代をキャンセルします。応答中にエージェントを誘導します。現在の応答が終了する前にフォローアップを送信します。
これにより、ビルド可能なものが変更されます。
AI Transport がこれを実装する方法
AI Transport は、Ably チャネル上に耐久性のあるセッションを実装します。 Ably チャネルは、永続セッションに必要な次のプロパティを提供します。
クライアントまたはエージェントは、チャネル名を指定してセッションに接続します。
チャネル上のメッセージは、単一の接続、デバイス、またはエージェントよりも存続します。
イベントは、たとえ切断されても、発行された順序でサブスクライバーに届きます。
接続が切断されたクライアントは再接続し、中断したところから再開します。
すべての参加者がチャネルに公開します。キャンセル、操作、中断はすべて同じセッションを通じて行われます。
複数の参加者が同じチャンネルに登録します。すべての参加者がすべてのイベントを参照します。
参加者は特別ではありません。ドロップして再接続するクライアント、1 ターン起動して終了するサーバーレス エージェント、別のデバイスから参加する 2 番目のクライアント、サブエージェントに委任するオーケストレーターはすべて、同じセッションと同等の条件で対話します。セッションは、参加者の接続ライフサイクルとは無関係に持続します。
AI Transport SDK は、このモデルを実用的なものにする抽象化を提供します。
ドメイン固有のメッセージ モデルを橋渡しするコーデック層 (Vercel

AI SDK の UIMessage など) と、トークンごとのストリーム配信のサポートを含む、Ably のネイティブ メッセージ プリミティブ。
チャネル (または外部ストア) からの会話状態を、ページネーションと分岐ナビゲーション用のビューを備えた分岐会話ツリーに具体化するセッション レイヤー。
通信メカニズムを処理するトランスポート層: メッセージのパブリッシュ、ストリームのルーティング、ターン ライフサイクルの管理、キャンセル信号の配信。
ストリーミング、ページネーション、ブランチ ナビゲーションを備えた UI を構築するための React フック。
既存のフレームワークに組み込まれるアダプター。 AI Transport は Vercel AI SDK の useChat に直接接続されます。
HTTP ストリーミングと AI : AI Transport が解決する詳細な問題説明。
概念: 構成要素 (セッション、ターン、トランスポート、コーデック、会話ツリー、インフラストラクチャ)。
Vercel AI SDK の使用を開始する : 動作するアプリを構築します。

## Original Extract

Why AI Transport exists: direct HTTP streaming breaks down in production. A durable session layer fixes resume, multi-device, and bidirectional control.

Ably Realtime | Why AI Transport
Docs Documentation Examples Ask AI Login Start free svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Platform svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Ably Pub/Sub svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold" data-radix-collection-item=""> Ably Chat svg]:rotate-90 data-[state=open]:border-b data-[state=open]:sticky data-[state=open]:top-0 h-12 px-4 py-3 font-bold text-neutral-1300 dark:text-neutral-000" data-radix-collection-item=""> Ably AI Transport svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Overview svg]:rotate-90 rounded-lg text-neutral-1300 dark:text-neutral-000 font-bold" data-radix-collection-item=""> Why AI Transport svg]:rotate-90 rounded-lg border-l dark:border-neutral-1000 hover:border-neutral-500 dark:hover:border-neutral-800 rounded-l-none text-neutral-1300 dark:text-neutral-000 font-bold border-orange-600 bg-orange-100 hover:bg-orange-100" data-radix-collection-item=""> Overview svg]:rotate-90 font-medium rounded-lg border-l border-neutral-300 dark:border-neutral-1000 hover:border-neutral-500 dark:hover:border-neutral-800 rounded-l-none" data-radix-collection-item=""> HTTP streaming and AI svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Concepts svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Getting started svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Frameworks svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Features svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> Going to production svg]:rotate-90 font-medium rounded-lg" data-radix-collection-item=""> API reference svg]:rotate-90 font-medium rounded-lg" dat
[truncated]
Direct HTTP streaming breaks down once an AI app is in production. AI Transport replaces a single HTTP connection with a durable session that any participant connects to.
Most AI frameworks connect a client to an agent over an HTTP request and streamed response. The pattern is simple and every framework supports it, and it works for one-off interactions and demos. It also limits everything that comes after that first interaction.
Streams die with the connection. Sessions cannot span devices. Clients have no mechanism to re-contact the agent once the original request is in flight. Each of these is a real production problem; together they bound the quality of AI experiences you build on direct HTTP streaming. See HTTP streaming and AI for the detail.
A durable session changes the model
These problems share a root cause. The transport is coupled to the interaction. The connection, the request, and the streamed response are ephemeral: they exist for one interaction, for one client and one agent. Once the HTTP connection is lost, so is the interaction between client and agent.
The shape that engineering teams are adopting is the durable session: a shared, persistent medium that any client or agent connects to, instead of a single HTTP request between one client and one agent.
A durable session provides three capabilities that direct HTTP streaming does not:
Resilient delivery. Streams survive connection drops, device switches, page refreshes, and process restarts. The client resumes from a known position. The agent continues publishing regardless of client connectivity. No events are lost and no events are duplicated.
Continuity across surfaces. The session follows the user, not the connection. Open a second tab, switch to a phone, come back hours later. Every surface sees the same session state. Any client with the session identifier attaches and hydrates the same conversation state.
Live control. Clients have a bi-directional stream to communicate with the agent and steer its processing while work is in progress. Cancel a generation from a different device. Steer an agent mid-response. Send a follow-up before the current response finishes.
This changes what is buildable:
How AI Transport implements this
AI Transport implements durable sessions on top of Ably channels . Ably channels provide the properties a durable session requires:
Any client or agent connects to the session by specifying a channel name.
Messages on the channel outlive any single connection, device, or agent.
Events arrive at subscribers in the order they were published, even across disconnections.
A client that drops its connection reconnects and picks up where it left off.
Any participant publishes to the channel. Cancel, steer, and interrupt all happen through the same session.
Multiple participants subscribe to the same channel; every participant sees every event.
No participant is special. A client that drops and reconnects, a serverless agent that spins up for one turn and terminates, a second client joining from another device, and an orchestrator delegating to sub-agents all interact with the same session on equal terms. The session persists independently of any participant's connection lifecycle.
The AI Transport SDK provides the abstractions that make this model practical:
A codec layer that bridges domain-specific message models (Vercel AI SDK's UIMessage , or any other) and Ably's native message primitives, including support for streamed token-by-token delivery.
A session layer that materialises conversation state from the channel (or from an external store) into a branching conversation tree with views for pagination and branch navigation.
A transport layer that handles communication mechanics: publishing messages, routing streams, managing turn lifecycle , and delivering cancel signals.
React hooks for building UIs with streaming, pagination, and branch navigation.
Adapters that drop into existing frameworks. AI Transport plugs directly into Vercel AI SDK's useChat .
HTTP streaming and AI : the detailed problem statement that AI Transport solves.
Concepts : the building blocks (sessions, turns, transport, codec, conversation tree, infrastructure).
Getting started with Vercel AI SDK : build a working app.
