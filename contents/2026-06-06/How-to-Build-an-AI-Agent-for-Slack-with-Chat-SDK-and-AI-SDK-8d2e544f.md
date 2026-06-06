---
source: "https://vercel.com/kb/guide/how-to-build-an-ai-agent-for-slack-with-chat-sdk-and-ai-sdk"
hn_url: "https://news.ycombinator.com/item?id=48421418"
title: "How to Build an AI Agent for Slack with Chat SDK and AI SDK"
article_title: "How to build an AI agent for Slack with Chat SDK and AI SDK | Vercel Knowledge Base"
author: "flashbrew"
captured_at: "2026-06-06T07:11:30Z"
capture_tool: "hn-digest"
hn_id: 48421418
score: 1
comments: 0
posted_at: "2026-06-06T04:34:58Z"
tags:
  - hacker-news
  - translated
---

# How to Build an AI Agent for Slack with Chat SDK and AI SDK

- HN: [48421418](https://news.ycombinator.com/item?id=48421418)
- Source: [vercel.com](https://vercel.com/kb/guide/how-to-build-an-ai-agent-for-slack-with-chat-sdk-and-ai-sdk)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T04:34:58Z

## Translation

タイトル: Chat SDK と AI SDK を使用して Slack 用 AI エージェントを構築する方法
記事のタイトル: Chat SDK と AI SDK を使用して Slack 用の AI エージェントを構築する方法 |ヴァーセルナレッジベース
説明: Chat SDK、AI SDK の ToolLoopAgent、および Vercel AI Gateway を使用して Slack AI エージェントを構築します。プロジェクトのセットアップ、ツール定義、ストリーミング応答、Vercel へのデプロイメント、ツールピックを使用したスケーリング ツール選択について説明します。

記事本文:
Chat SDK と AI SDK を使用して Slack 用の AI エージェントを構築する方法 | Vercel Knowledge Base コンテンツにスキップ Chat SDK と AI SDK を使用して Slack 用の AI エージェントを構築する方法
Chat SDK、AI SDK の ToolLoopAgent、Vercel AI Gateway を使用して Slack AI エージェントを構築します。プロジェクトのセットアップ、ツール定義、ストリーミング応答、Vercel へのデプロイメント、ツールピックを使用したスケーリング ツール選択について説明します。
ナレッジベース / チャット SDK URL をコピー ページをコピー このページについて AI に質問する 10 分で読める 最終更新日 2026 年 6 月 4 日 Chat SDK と AI SDK を使用して、メンションに応答し、会話履歴を維持し、ツールを自律的に呼び出す AI 搭載の Slack エージェントを構築できます。 Chat SDK はプラットフォームの統合 (Webhook、メッセージの書式設定、スレッド追跡) を処理し、AI SDK の ToolLoopAgent は、エージェントがツールを呼び出して結果に基づいて動作できるようにする推論ループを管理します。 Vercel AI Gateway と Redis for state を併用すると、インフラストラクチャを管理したりプロバイダー SDK を操作したりすることなく、運用環境に対応した Slack エージェントを入手できます。
このガイドでは、Chat SDK、AI SDK の ToolLoopAgent 、および Vercel AI Gateway 経由の Claude を使用して Slack エージェントを構築する手順を説明します。ストリーミング応答、ツール呼び出し、複数ターンの会話履歴を結び付けてから、toolpick を使用して実稼働用にツール セットを拡張します。
始める前に、次のものが揃っていることを確認してください。
アプリをインストールできる Slack ワークスペース
Redis インスタンス (ローカルまたはホスト型 ( Upstash など))
AI Gateway API キーを持つ Vercel アカウント
Chat SDK は、Slack、Teams、Discord、その他のプラットフォームでチャットボットを構築するための統合 TypeScript SDK です。イベント ハンドラー ( onNewMention や onSubscribedMessage など) を登録すると、SDK は受信 Webhook をそれらのハンドラーにルーティングします。 Slack アダプターは、Webhook 検証、メッセージ解析、および Slack API を処理します。 Redis 状態アダプターは、

ボットがサブスクライブし、同時メッセージ処理のための分散ロックを管理するスレッド。
AI SDK の ToolLoopAgent は言語モデルをツールでラップし、自律ループを実行します。モデルはテキストを生成するかツールを呼び出し、SDK はツールを実行して結果をフィードバックし、モデルが終了するまで繰り返します。 「anthropic/claude-sonnet-4.6」 のようなモデル文字列を渡し、Vercel でアプリケーションをホストすると、AI SDK はリクエストを AI ゲートウェイ経由で自動的にルーティングします。
Chat SDK は、任意の AsyncIterable<string> をメッセージとして受け入れるため、Slack でのリアルタイム ストリーミングのためにエージェントの fullStream を thread.post() に直接渡すことができます。
1. プロジェクトをスキャフォールディングし、依存関係をインストールし、Vercel プラグインを追加します
新しい Next.js アプリを作成し、Chat SDK、AI SDK、アダプター パッケージを追加します。
Terminal npx create-next-app@latest my-slack-agent --typescript --app cd my-slack-agent pnpm add chat @chat-adapter/slack @chat-adapter/state-redis ai zod チャット パッケージは Chat SDK コアです。 @chat-adapter/slack および @chat-adapter/state-redis パッケージは、Slack プラットフォーム アダプターと Redis 状態アダプターです。 ai パッケージは AI SDK であり、これには AI ゲートウェイ プロバイダーと ToolLoopAgent が含まれます。 zod は、ツール入力スキーマを定義するために使用されます。
Vercel プラグインは、AI コーディング エージェント (Claude Code など) にスキル、専門エージェント、スラッシュ コマンドなどを装備します。
ターミナル npx プラグインは vercel/vercel-plugin を追加します 2. Slack アプリを作成します
api.slack.com/apps に移動し、「Create New App」をクリックし、「From a Manifest」をクリックします。
ワークスペースを選択し、次のマニフェストを貼り付けます。
lack-manifest.json 表示情報 : 名前 : AI エージェントの説明 : Chat SDK と AI SDK で構築された AI エージェント
機能 : ボットユーザー : 表示名 : AI エージェント always_online : true
oauth_config : スコープ : ボット : - app_mentions : 読み取り - チャネル :

履歴 - チャネル : 読み取り - チャット : 書き込み - グループ : 履歴 - グループ : 読み取り - im : 履歴 - im : 読み取り - mpim : 履歴 - mpim : 読み取り - リアクション : 読み取り - リアクション : 書き込み - ユーザー : 読み取り
設定 :event_subscriptions : request_url : https : //your-domain.com/api/webhooks/slack bot_events : - app_mention - message.channels - message.groups - message.im - message.mpim 対話性 : is_enabled : true request_url : https : //your-domain.com/api/webhooks/slack org_deploy_enabled : falsesocket_mode_enabled : false token_rotation_enabled : false アプリを作成した後:
[アプリのインストール] に移動し、アプリをワークスペースにインストールします。
[OAuth と権限] > [OAuth トークン] に移動し、ボット ユーザー OAuth トークンをコピーします。
[基本情報] > [アプリ資格情報] に移動し、署名シークレットをコピーします。
デプロイ後に request_url プレースホルダーを実際のドメイン (またはローカル テスト用のトンネル URL) に置き換えます。
3. 環境変数を設定する
プロジェクト ルートに .env.local ファイルを作成します。
.env.local SLACK_BOT_TOKEN = xoxb-your-bot-token SLACK_SIGNING_SECRET = your-signing-secret REDIS_URL = redis://localhost:6379 AI_GATEWAY_API_KEY = your-ai-gateway-api-key Slack アダプターは、SLACK_BOT_TOKEN と SLACK_SIGNING_SECRET を自動的に読み取ります。 Redis 状態アダプターは REDIS_URL を読み取ります。 AI SDK は AI_GATEWAY_API_KEY を使用して Vercel AI Gateway で認証するか、OIDC 認証を使用します。
Vercel ダッシュボードの AI Gateway で AI Gateway API キーを作成し、[API キーの作成] をクリックします。
エージェントが呼び出すことができるツールを含む lib/tools.ts を作成します。この例では天気予報ツールとドキュメント ツールを定義していますが、ユースケースに必要なツールを追加できます。
lib/tools.ts import { ツール } から "ai" ; "zod" から { z } をインポートします。
import const tools = { getWeather : tools ( { description : "現在の情報を取得します

t 場所の天気" , inputSchema : z . object ( { location : z . string ( ) . description ( "City name, e.g. San Francisco" ) , } ) ,execute : async ( { location } ) => { // 実際の天気 API 呼び出しに置き換えます const response = await fetch ( ` https://api.weatherapi.com/v1/current.json?key= ${ process . env . WEATHER_API_KEY } &q= ${ encodeURIComponent ( location ) } ` ) ; const data = await response . json ( ) ; return { location , temperature : data . current . temp_f , condition : data . current . condition . text , } ; } , } ) , searchDocs : tools ( { description : "会社のドキュメントでトピックを検索する" , inputSchema : z . object ( { query : z . string ( ) .describe ( "検索クエリ" ) , } ) ,execute : async ( { query } ) => { // 実際の検索実装に置き換えます return { results : [ ` Result for: ${ query } ` ] } ; } , } ) , } ; 各ツールには、説明 (いつ使用するかをモデルに指示する)、inputSchema (モデルが入力する Zod スキーマ)、およびツールの呼び出し時に実行される実行関数があります。
ToolLoopAgent と Chat インスタンスを使用して lib/bot.ts を作成します。
lib/bot.ts import { Chat } from "チャット" ; "chat/ai" から { toAiMessages } をインポートします。 import { createSlackAdapter } から "@chat-adapter/slack" ; import { createRedisState } from "@chat-adapter/state-redis" ; 「ai」から { ToolLoopAgent } をインポートします。 "./tools" から { tools } をインポートします。
const Agent = new ToolLoopAgent ( { model : "anthropic/claude-sonnet-4.6" , 指示 : "あなたは Slack ワークスペースの役に立つ AI アシスタントです。" + "質問に明確に答え、必要なときにツールを使用してください。" + "リアルタイム データ。応答はチャット用に簡潔かつ適切な形式に保ってください。" , tools , } ) ;
import const bot = new Chat ( { userName : "ai-agent" , アダプター : {lack : createSlackAdapter () , } , sta

te:createRedisState(), });
// 初めての言及ボットを処理します。 onNewMention(async(thread,message)=>{await thread.subscribe();
const result = エージェントを待ちます。 stream({プロンプト:メッセージ.テキスト});スレッドを待ちます。 post (結果 .fullStream ) ; } ) ;
// サブスクライブされたスレッドでフォローアップ メッセージを処理します bot 。 onSubscribedMessage (async (thread , message) => {const allMessages = [] ; for await (const msg of thread .allMessages) {allMessages .push(msg) ; }
const History = aiMessages (allMessages) を待ちます。
const result = エージェントを待ちます。ストリーム({メッセージ:履歴});スレッドを待ちます。 post (結果 .fullStream ) ; } ) ;誰かがボットを @メンションすると、onNewMention が起動します。ハンドラーはスレッドをサブスクライブし (そのスレッド内の将来のメッセージを追跡するため)、エージェントの応答をストリーミングします。フォローアップ メッセージの場合、 onSubscribedMessage は thread.allMessages を使用して完全なスレッド履歴を取得し、 toAiMessages を使用して AI SDK メッセージ形式に変換してエージェントに渡し、完全な会話コンテキストを保持します。
fullStream は、ツール呼び出しステップ間の段落区切りを保持するため、textStream よりも優先されます。 Chat SDK はストリーム タイプを自動検出し、Slack のネイティブ ストリーミング API を処理してリアルタイム更新を行います。
app/api/webhooks/[platform]/route.ts に API ルートを作成します。
app/api/webhooks/[platform]/route.ts import { after } from "next/server" ; "@/lib/bot" から { bot } をインポートします。
type Platform = keyof typeof bot 。 Webhook ;
import async function POST (request : Request , context : RouteContext < "/api/webhooks/[platform]" > ) { const { platform } = コンテキストを待ちます。パラメータ ;
const ハンドラー = ボット 。 Webhook [プラットフォームとしてのプラットフォーム] ; if ( ! handler ) { return new Response (` 不明なプラットフォーム: ${ platform } ` , { status : 404 } ) ; }

return ハンドラ (request , {waitUntil :(task) => after (() => task) , });これにより、POST /api/webhooks/slack エンドポイントが作成されます。 waitUntil オプションを使用すると、HTTP 応答の送信後にイベント ハンドラーの処理が確実に終了します。これは、関数が早期に終了してしまうサーバーレス プラットフォームでは必須です。
開発サーバーを起動します: Terminal pnpm dev
トンネルで公開します: Terminal npx ngrok http 3000
トンネル URL (例: https://abc123.ngrok-free.dev ) をコピーし、Slack アプリ設定のイベント サブスクリプション URL とインタラクティブ リクエスト URL の両方を https://abc123.ngrok-free.dev/api/webhooks/slack に更新します。
ボットをチャンネルに招待します ( /invite @AI Agent )
質問をしてボットを @メンションします。スレッドにストリーミング応答が表示されるはずです。 「サンフランシスコの天気は？」など、ツールの 1 つを使用するように依頼してみてください。
まず、プロジェクトをリンクし、環境変数を追加します。
Terminal vercel link vercel env add SLACK_BOT_TOKEN vercel env add SLACK_SIGNING_SECRET vercel env add REDIS_URL vercel env add AI_GATEWAY_API_KEY または、Vercel ダッシュボードの [設定] > [環境変数] に追加します。
Terminal vercel Slack アプリ設定のイベント サブスクリプションとインタラクティビティ リクエスト URL を運用 URL (例: https://my-slack-agent.vercel.app/api/webhooks/slack ) に更新します。
Vercel にデプロイすると、AI Gateway は OIDC ベースの認証をサポートするため、静的 API キーを使用せずに認証することもできます。 AI ゲートウェイの認証ドキュメントを参照してください。
ボットがメンションに応答しない
Slack アプリに app_mentions:read スコープがあること、およびイベント サブスクリプション リクエスト URL が正しいことを確認してください。最初に URL を設定するときに Slack はチャレンジ リクエストを送信するため、サーバーが実行されている必要があります。
ストリーミングが途切れ途切れまたは遅延しているように見える

Chat SDK は Slack のネイティブ ストリーミング API を使用して、スムーズな更新を実現します。 SDK は分散ロックを使用して同時メッセージを管理するため、問題が発生した場合は、Redis 接続が安定していることを確認してください。
エージェントがツールを呼び出しても結果が表示されない場合は、ツールの実行関数でエラーがないか確認してください。 AI SDK はツール実行エラーをモデルに表示し、回復を試みる可能性があります。ツールにエラー処理を追加し、サーバー ログで詳細を確認してください。
スレッド履歴が大きくなりすぎる
長時間実行されるスレッドの場合、会話履歴がモデルのコンテキスト ウィンドウを超える可能性があります。履歴配列をスライスするか、古いメッセージの要約ステップを使用して、エージェントに渡すメッセージの数を制限することを検討してください。
ツールピックを使用して多くのツールにスケーリングする
このガイドのエージェントには 2 つのツールがあります。運用環境では、GitHub、Linear、Upstash、カレンダーなどのサービスを統合し、パイプラインをデプロイすると、Slack エージェントは 15、20、または 30 ツールに増加することがよくあります。その規模では、すべてのツール定義がすべてのステップでモデルに送信されるため、トークンのコストが増加し、モデルが適切なツールを選択することが困難になります。
Toolpick は、起動時にツールにインデックスを付け、EA に最も関連性の高いツールのみを選択することでこの問題を解決します。

[切り捨てられた]

## Original Extract

Build a Slack AI agent using Chat SDK, AI SDK's ToolLoopAgent, and Vercel AI Gateway. Covers project setup, tool definitions, streaming responses, deployment to Vercel, and scaling tool selection with toolpick.

How to build an AI agent for Slack with Chat SDK and AI SDK | Vercel Knowledge Base Skip to content How to build an AI agent for Slack with Chat SDK and AI SDK
Build a Slack AI agent using Chat SDK, AI SDK's ToolLoopAgent, and Vercel AI Gateway. Covers project setup, tool definitions, streaming responses, deployment to Vercel, and scaling tool selection with toolpick.
Knowledge Base / Chat SDK Copy URL Copy page Ask AI about this page 10 min read Last updated June 4, 2026 You can build an AI-powered Slack agent that responds to mentions, maintains conversation history, and calls tools autonomously using Chat SDK and AI SDK. Chat SDK handles the platform integration (webhooks, message formatting, thread tracking), while AI SDK's ToolLoopAgent manages the reasoning loop that lets your agent call tools and act on results. Together with Vercel AI Gateway and Redis for state, you get a production-ready Slack agent without managing infrastructure or juggling provider SDKs.
This guide will walk you through building a Slack agent with Chat SDK, AI SDK's ToolLoopAgent , and Claude via the Vercel AI Gateway . You'll wire up streaming responses, tool calling, and multi-turn conversation history, then scale your tool set for production with toolpick.
Before you begin, make sure you have:
A Slack workspace where you can install apps
A Redis instance (local or hosted, such as Upstash )
A Vercel account with an AI Gateway API key
Chat SDK is a unified TypeScript SDK for building chatbots across Slack, Teams, Discord, and other platforms. You register event handlers (like onNewMention and onSubscribedMessage ), and the SDK routes incoming webhooks to them. The Slack adapter handles webhook verification, message parsing, and the Slack API. The Redis state adapter tracks which threads your bot has subscribed to and manages distributed locking for concurrent message handling.
AI SDK's ToolLoopAgent wraps a language model with tools and runs an autonomous loop: the model generates text or calls a tool, the SDK executes the tool, feeds the result back, and repeats until the model finishes. When you pass a model string like "anthropic/claude-sonnet-4.6" , and host your application on Vercel, the AI SDK will route the request through the AI Gateway automatically.
Chat SDK accepts any AsyncIterable<string> as a message, so you can pass the agent's fullStream directly to thread.post() for real-time streaming in Slack.
1. Scaffold the project, install dependencies, and add Vercel Plugin
Create a new Next.js app and add the Chat SDK, AI SDK, and adapter packages:
Terminal npx create-next-app@latest my-slack-agent --typescript --app cd my-slack-agent pnpm add chat @chat-adapter/slack @chat-adapter/state-redis ai zod The chat package is the Chat SDK core. The @chat-adapter/slack and @chat-adapter/state-redis packages are the Slack platform adapter and Redis state adapter. The ai package is the AI SDK, which includes the AI Gateway provider and ToolLoopAgent . zod is used to define tool input schemas.
The Vercel Plugin equips your AI coding agent (e.g., Claude Code) with skills, specialist agents, slash commands, and more.
Terminal npx plugins add vercel/vercel-plugin 2. Create a Slack app
Go to api.slack.com/apps , click Create New App , then From a manifest .
Select your workspace and paste this manifest:
slack-manifest.json display_information : name : AI Agent description : An AI agent built with Chat SDK and AI SDK
features : bot_user : display_name : AI Agent always_online : true
oauth_config : scopes : bot : - app_mentions : read - channels : history - channels : read - chat : write - groups : history - groups : read - im : history - im : read - mpim : history - mpim : read - reactions : read - reactions : write - users : read
settings : event_subscriptions : request_url : https : //your-domain.com/api/webhooks/slack bot_events : - app_mention - message.channels - message.groups - message.im - message.mpim interactivity : is_enabled : true request_url : https : //your-domain.com/api/webhooks/slack org_deploy_enabled : false socket_mode_enabled : false token_rotation_enabled : false After creating the app:
Go to Install App , and install the app to your workspace
Go to OAuth & Permissions > OAuth Tokens and copy the Bot User OAuth Token
Go to Basic Information > App Credentials and copy the Signing Secret
You'll replace the request_url placeholders with your real domain after deploying (or a tunnel URL for local testing).
3. Configure environment variables
Create a .env.local file in your project root:
.env.local SLACK_BOT_TOKEN = xoxb-your-bot-token SLACK_SIGNING_SECRET = your-signing-secret REDIS_URL = redis://localhost:6379 AI_GATEWAY_API_KEY = your-ai-gateway-api-key The Slack adapter reads SLACK_BOT_TOKEN and SLACK_SIGNING_SECRET automatically. The Redis state adapter reads REDIS_URL . AI SDK uses AI_GATEWAY_API_KEY to authenticate with the Vercel AI Gateway, or alternatively, use OIDC authentication .
You can create an AI Gateway API key from your Vercel dashboard under AI Gateway and click Create an API Key .
Create lib/tools.ts with the tools your agent can call. This example defines a weather tool and docs tool, but you can add any tools your use case requires:
lib/tools.ts import { tool } from "ai" ; import { z } from "zod" ;
export const tools = { getWeather : tool ( { description : "Get the current weather for a location" , inputSchema : z . object ( { location : z . string ( ) . describe ( "City name, e.g. San Francisco" ) , } ) , execute : async ( { location } ) => { // Replace with a real weather API call const response = await fetch ( ` https://api.weatherapi.com/v1/current.json?key= ${ process . env . WEATHER_API_KEY } &q= ${ encodeURIComponent ( location ) } ` ) ; const data = await response . json ( ) ; return { location , temperature : data . current . temp_f , condition : data . current . condition . text , } ; } , } ) , searchDocs : tool ( { description : "Search the company documentation for a topic" , inputSchema : z . object ( { query : z . string ( ) . describe ( "The search query" ) , } ) , execute : async ( { query } ) => { // Replace with your actual search implementation return { results : [ ` Result for: ${ query } ` ] } ; } , } ) , } ; Each tool has a description (which tells the model when to use it), an inputSchema (a Zod schema that the model fills in), and an execute function that runs when the tool is called.
Create lib/bot.ts with a ToolLoopAgent and a Chat instance:
lib/bot.ts import { Chat } from "chat" ; import { toAiMessages } from "chat/ai" ; import { createSlackAdapter } from "@chat-adapter/slack" ; import { createRedisState } from "@chat-adapter/state-redis" ; import { ToolLoopAgent } from "ai" ; import { tools } from "./tools" ;
const agent = new ToolLoopAgent ( { model : "anthropic/claude-sonnet-4.6" , instructions : "You are a helpful AI assistant in a Slack workspace. " + "Answer questions clearly and use your tools when you need " + "real-time data. Keep responses concise and well-formatted for chat." , tools , } ) ;
export const bot = new Chat ( { userName : "ai-agent" , adapters : { slack : createSlackAdapter ( ) , } , state : createRedisState ( ) , } ) ;
// Handle first-time mentions bot . onNewMention ( async ( thread , message ) => { await thread . subscribe ( ) ;
const result = await agent . stream ( { prompt : message . text } ) ; await thread . post ( result . fullStream ) ; } ) ;
// Handle follow-up messages in subscribed threads bot . onSubscribedMessage ( async ( thread , message ) => { const allMessages = [ ] ; for await ( const msg of thread . allMessages ) { allMessages . push ( msg ) ; }
const history = await toAiMessages ( allMessages ) ;
const result = await agent . stream ( { messages : history } ) ; await thread . post ( result . fullStream ) ; } ) ; When someone @mentions the bot, onNewMention fires. The handler subscribes to the thread (to track future messages in that thread) and streams the agent's response. For follow-up messages, onSubscribedMessage retrieves the full thread history using thread.allMessages , converts it to the AI SDK message format with toAiMessages and passes it to the agent so it has a complete conversation context.
The fullStream is preferred over textStream because it preserves paragraph breaks between tool-calling steps. Chat SDK auto-detects the stream type and handles Slack's native streaming API for real-time updates.
Create the API route at app/api/webhooks/[platform]/route.ts :
app/api/webhooks/[platform]/route.ts import { after } from "next/server" ; import { bot } from "@/lib/bot" ;
type Platform = keyof typeof bot . webhooks ;
export async function POST ( request : Request , context : RouteContext < "/api/webhooks/[platform]" > ) { const { platform } = await context . params ;
const handler = bot . webhooks [ platform as Platform ] ; if ( ! handler ) { return new Response ( ` Unknown platform: ${ platform } ` , { status : 404 } ) ; }
return handler ( request , { waitUntil : ( task ) => after ( ( ) => task ) , } ) ; } This creates a POST /api/webhooks/slack endpoint. The waitUntil option ensures your event handlers finish processing after the HTTP response is sent, which is required on serverless platforms where the function would otherwise terminate early.
Start the dev server: Terminal pnpm dev
Expose it with a tunnel: Terminal npx ngrok http 3000
Copy the tunnel URL (for example, https://abc123.ngrok-free.dev ) and update both Event Subscriptions and Interactivity Request URLs in your Slack app settings to https://abc123.ngrok-free.dev/api/webhooks/slack
Invite the bot to a channel ( /invite @AI Agent )
@mention the bot with a question. You should see a streaming response appear in the thread. Try asking it to use one of your tools, such as "What's the weather in San Francisco?"
First, link your project and add your environment variables:
Terminal vercel link vercel env add SLACK_BOT_TOKEN vercel env add SLACK_SIGNING_SECRET vercel env add REDIS_URL vercel env add AI_GATEWAY_API_KEY Alternatively, add them in the Vercel dashboard under Settings > Environment Variables .
Terminal vercel Update the Event Subscriptions and Interactivity Request URLs in your Slack app settings to your production URL, for example https://my-slack-agent.vercel.app/api/webhooks/slack .
When deployed to Vercel, AI Gateway supports OIDC-based authentication, so you can also authenticate without a static API key. See the AI Gateway authentication docs .
Bot doesn't respond to mentions
Check that your Slack app has the app_mentions:read scope and that the Event Subscriptions Request URL is correct. Slack sends a challenge request when you first set the URL, so your server must be running.
Streaming appears choppy or delayed
Chat SDK uses Slack's native streaming API for smooth updates. If you're seeing issues, check that your Redis connection is stable, as the SDK uses distributed locks to manage concurrent messages.
If the agent calls a tool but no result appears, check for errors in your tool's execute function. AI SDK surfaces tool execution errors back to the model, which may attempt to recover. Add error handling in your tools and check your server logs for details.
Thread history grows too large
For long-running threads, the conversation history can exceed the model's context window. Consider limiting the number of messages you pass to the agent by slicing the history array or by using a summarization step for older messages.
Scaling to many tools with toolpick
The agent in this guide has two tools. In production, a Slack agent often grows to 15, 20, or 30 tools as you integrate services like GitHub, Linear , Upstash , calendars, and deploy pipelines. At that scale, every tool definition is sent to the model on every step, which increases token costs and makes it harder for the model to pick the right tool.
toolpick solves this by indexing your tools at startup and selecting only the most relevant ones for ea

[truncated]
