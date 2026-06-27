---
source: "https://vercel.com/kb/guide/ai-sdk-mcp-apps"
hn_url: "https://news.ycombinator.com/item?id=48693667"
title: "Add MCP Apps to Your AI SDK Application"
article_title: "Add MCP Apps to your AI SDK application | Vercel Knowledge Base"
author: "flashbrew"
captured_at: "2026-06-27T00:30:27Z"
capture_tool: "hn-digest"
hn_id: 48693667
score: 1
comments: 0
posted_at: "2026-06-27T00:10:53Z"
tags:
  - hacker-news
  - translated
---

# Add MCP Apps to Your AI SDK Application

- HN: [48693667](https://news.ycombinator.com/item?id=48693667)
- Source: [vercel.com](https://vercel.com/kb/guide/ai-sdk-mcp-apps)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T00:10:53Z

## Translation

タイトル: AI SDK アプリケーションに MCP アプリを追加する
記事のタイトル: MCP アプリを AI SDK アプリケーションに追加する |ヴァーセルナレッジベース
説明: @ai-sdk/mcp と @ai-sdk/react を使用して AI SDK で MCP アプリ ホストを構築し、モデルに表示されるツールをフィルターし、ui:// リソースを読み取り、experimental_MCPAppRenderer を使用してサンドボックス iframe でインタラクティブ ツール UI をレンダリングします。

記事本文:
MCP アプリを AI SDK アプリケーションに追加する | Vercel ナレッジ ベース コンテンツへスキップ ロゴを SVG としてコピー ワードマークを SVG としてコピー ブランド ガイドライン ドキュメントのビルド AI を使用したビルド AI ゲートウェイ
MCP アプリを AI SDK アプリケーションに追加する
@ai-sdk/mcp と @ai-sdk/react を使用して AI SDK で MCP アプリ ホストを構築し、モデルに表示されるツールをフィルターし、ui:// リソースを読み取り、experimental_MCPAppRenderer を使用してサンドボックス iframe でインタラクティブ ツール UI をレンダリングします。
ナレッジ ベース / AI SDK URL をコピー ページをコピー このページについて AI に質問する 最終更新日 2026 年 6 月 25 日 MCP アプリを使用すると、モデル コンテキスト プロトコル (MCP) ツールがプレーン テキストの代わりに対話型 UI を返すことができます。モデルは依然として通常の MCP ツールを呼び出しますが、ツールは HTML を保持する ui:// リソースを指すことができ、アプリはその HTML をサンドボックス化された iframe でレンダリングします。ホスト側を構築するために、AI SDK は、MCP アプリのサポートをアドバタイズし、モデルが参照するツールをフィルタリングし、ui:// リソースを読み取るための @ai-sdk/mcp ヘルパーに加えて、iframe をレンダリングしてそのメッセージをブリッジするための @ai-sdk/react コンポーネントを提供します。その後、信頼できない HTML は分離されたままで、チャットにダッシュボード、フォーム、またはツールによって生成されたその他の対話型ビューを表示できます。
このガイドでは、次の方法を学習します。
MCP Apps 機能を使用して MCP サーバーに接続する
SplitMCPAppTools を使用して、モデルに表示されるツールのみをモデルに渡します
readMCPAppResource を使用してツールの ui:// リソースを読み取ります
アプリで開始されたツール呼び出しを iframe から安全にプロキシする
Experimental_MCPAppRenderer を使用して React チャットでアプリ UI をレンダリングする
始める前に、次のものが揃っていることを確認してください。
@ai-sdk/mcp および @ai-sdk/react を使用した AI パッケージ
MCP TypeScript SDK ( @modelcontextprotocol/sdk ) およびプロバイダー パッケージ (@ai-sdk/openai など)
MCP Apps ツール (ui:// リソースを指すツール) を公開する MCP サーバー
u を使用する React アプリ

seChat (以下の例では Next.js App Router を使用しています)
MCP Apps ホストは MCP サーバーに接続し、モデルが表示できるツールを決定し、ツールが指すアプリ UI をレンダリングします。実行時に、ホストは次の手順に従います。
MCP Apps クライアント機能を使用して MCP サーバーに接続します。
サーバーのツールをリストし、MCP アプリの可視性によってそれらを分割します。
モデルに表示されるツールのみを streamText またはgenerateText に渡します。
ツール部分に MCP アプリのメタデータが含まれている場合は、ツールの ui:// リソースを読み取ります。
サンドボックス化された iframe で HTML リソースをレンダリングします。
プロキシは、アプリに表示されるツール呼び出しなどの iframe リクエストを MCP サーバーに戻すことを許可しました。
このガイドの残りの部分では、各ステップを構築します。
MCP Apps サポートを使用して MCP サーバーに接続する
mcpAppClientCapabilities を使用して MCP クライアントを作成し、ホストが text/html;profile=mcp-app リソースをレンダリングできることをアドバタイズできるようにします。
app/api/chat/mcp-client.ts import { createMCPClient , mcpAppClientCapabilities } from '@ai-sdk/mcp' ; '@modelcontextprotocol/sdk/client/streamableHttp.js' から { StreamableHTTPClientTransport } をインポートします。
エクスポート関数 createMCPAppsClient (origin : string ) { return createMCPClient ( { Transport : new StreamableHTTPClientTransport ( new URL ('/mcp' ,origin ) ) , clientName : 'my-mcp-apps-host' ,abilities : mcpAppClientCapabilities , } ) ;ホストが MCP アプリ リソースを安全にフェッチおよびレンダリングできる場合にのみ、これらの機能をアドバタイズします。
モデルに表示されるツールのみを公開する
MCP Apps ツールは _meta.ui.visibility を宣言できます。 「model」とマークされたツールをモデルに渡し、iframe リクエストに対してのみ「app」とマークされたツールを維持し、モデルがそれらを認識しないようにします。ツール リストを SplitMCPAppTools で分割し、modelVisible を streamText に渡します。
app/api/chat/route.ts import { splitMCPAppTools } から '@ai-sdk/mcp' ; import {convertToModelMessages 、createUIMessageStre

amResponse 、 streamText 、 toUIMessageStream 、 } 'ai' から; './mcp-client' から { createMCPAppsClient } をインポートします。 '@ai-sdk/openai' から {openai } をインポートします。
エクスポート非同期関数 POST(req:Request) {const requestUrl = 新しい URL(req.url) ; const client = await createMCPAppsClient(requestUrl.origin); const {メッセージ} = 要求を待ちます。 json() ;
try {const 定義 = クライアントを待ちます。 listTools() ; const {modelVisible} = splitMCPAppTools (定義) ; const ツール = クライアント . toolsFromDefinitions (modelVisible) ;
const result = streamText({モデル:openai('gpt-5.2')、ツール、メッセージ:awaitconvertToModelMessages(メッセージ)、onEnd:async()=>{await client.close();},});
return createUIMessageStreamResponse({stream:toUIMessageStream({stream:result.stream}),}); } catch (エラー) {クライアントを待ちます。近い （ ） ;エラーをスローします。モデルがアプリベースのツールを呼び出すと、MCP クライアントはツール UI にアプリのメタデータを保持します。React レンダラーはこれを使用して、ツール パーツに MCP アプリがあるかどうかを判断します。
ツールの ui:// リソースをブラウザ ホストに送信する前に、readMCPAppResource を使用して読み取ります。
app/api/mcp-app-host/read-resource/route.ts import { readMCPAppResource } from '@ai-sdk/mcp' ; import { createMCPAppsClient } から '../../chat/mcp-client' ;
エクスポート非同期関数 POST(req:Request) {const requestUrl = 新しい URL(req.url) ; const {uri } = 要求を待ちます。 json() ; const client = await createMCPAppsClient(requestUrl.origin);
try { 応答を返します 。 json (await readMCPAppResource({client,uri})) ;最後に {クライアントを待ちます。近い （ ） ; readMCPAppResource は、リソースが ui:// URI を使用していることを確認し、MCP Apps MIME タイプを必要とし、テキストまたは Base64 コンテンツをデコードして、ren とともに HTML を返します。

コンテンツセキュリティポリシーや権限などのメタデータを抽出します。
iframe は MCP サーバーに直接到達できません。 JSON-RPC メッセージをホストに送信し、ホストがどのメッセージを許可するかを決定します。アプリによって開始されるツール呼び出しの場合は、MCP サーバーを呼び出す前に、要求されたツールがアプリに表示されていることを確認してください。
app/api/mcp-app-host/call-tool/route.ts import { splitMCPAppTools } から '@ai-sdk/mcp' ; import { createMCPAppsClient } から '../../chat/mcp-client' ;
エクスポート非同期関数 POST(req:Request) {const requestUrl = 新しい URL(req.url) ; const {名前、引数:toolArguments} = await req 。 json() ; const client = await createMCPAppsClient(requestUrl.origin);
try {const{appVisible}=splitMCPAppTools(await client.listTools()); const isAllowed = appVisible 。ツール。 some (ツール => ツール . 名前 === 名前 ) ;
if(!isAllowed) { 応答を返します。 json ( { エラー : 'ツールがアプリに表示されません' } , { ステータス : 403 } , ) ; }
応答を返します。 json (クライアントを待ちます。callTool({名前,引数:toolArguments??{},}),);最後に {クライアントを待ちます。近い （ ） ;実稼働環境では、iframe リクエストを転送する前に、アプリに必要なポリシーとユーザー承認チェックを追加します。
React チャットでアプリをレンダリングする
React チャット UI で、通常のメッセージ部分を通常どおりレンダリングし、ツール部分を Experimental_MCPAppRenderer に渡します。
Experimental_MCPAppRenderer は実験的なものであり、将来のリリースで変更される可能性があります。
app/page.tsx 'クライアントを使用' ;
import { Experimental_MCPAppRenderer as MCPAppRenderer , useChat , type MCPAppBridgeHandlers , type MCPAppMetadata , type MCPAppResource , type MCPAppSandboxConfig , } from '@ai-sdk/react' ; import { DefaultChatTransport , isToolUIPart } から 'ai' ;
const Sandbox = { url : '/mcp-app-sandbox' 、className : 'h-80 w-fullrounded-lg

border' 、 style : { border : 0 } 、 } は MCPAppSandboxConfig を満たします。
async function loadResource (app : MCPAppMetadata ) : Promise < MCPAppResource > { const response = await fetch ( '/api/mcp-app-host/read-resource' , { メソッド : 'POST' , body : JSON . stringify ( { uri : app . resourceUri } ) , } ) ;
if ( ! 応答 . ok ) { throw new Error ('MCP アプリ リソースのロードに失敗しました' ) ; }
返答を返します。 json() ; }
const ハンドラー : MCPAppBridgeHandlers = { callTool : params => fetch ( '/api/mcp-app-host/call-tool' , { メソッド : 'POST' , body : JSON . stringify ( params ) , } ) 。 then (response =>response.json())、openLink:({url})=>{window.json() open ( url , '_blank' , 'noopener,noreferrer' ) ;戻る { } ; } , } ;
デフォルト関数をエクスポート Chat () { const {messages , sendMessage } = useChat ({ Transport : new DefaultChatTransport ( { api : '/api/chat' } ) , } ) ;
return ( < > { メッセージ . マップ ( メッセージ => メッセージ . パーツ . マップ ( ( パーツ , インデックス ) => { if ( パーツ . type === 'テキスト' ) { return < div key = { インデックス } > { パーツ . テキスト } </ div > ; }
if ( isToolUIPart (part ) ) { return (< MCPAppRenderer key = {part .toolCallId } part = {part }loadResource = {loadResource } handlers = { handlers } Sandbox = { Sandbox } fallback = { < div > MCP アプリを読み込み中... </ div > } /> ) ; }
null を返します。 } ) ) }
< button onClick = { ( ) => sendMessage ( { text : 'ダッシュボードを表示' } ) } > 送信 </ button > </ > ) ; Experimental_MCPAppRenderer は、通常のツールでは何もレンダリングしません。アプリベースのツールの場合、リソースを読み込み、サンドボックス ブリッジを作成し、ツールの入力と結果の通知を iframe に送信し、
[切り捨てられた]
MCP アプリの HTML を信頼できないものとして扱います。理想的には別のオリジンのサンドボックス プロキシ ルートを通じて、サンドボックス iframe でレンダリングします。
アプリのみを渡さないでください

モデルにオールインワン。 SplitMCPAppTools を使用し、modelVisible ツールのみを公開します。
client.callTool を呼び出す前に、サーバー上のすべての iframe リクエストを検証してください。
resourceUri によってリソースをキャッシュするため、ツール呼び出しが繰り返されても同一の HTML が再フェッチされません。
各ツールのコンテンツと StructuredContent を単独で使用できるようにすることで、テキストのみのホストは UI がなくても機能します。
応答またはホスト要求が終了したら、有効期間の短い MCP クライアントを閉じます。
ホスト側関数については、MCP Apps ヘルパー リファレンスを参照してください。
React コンポーネントの props については、MCP App Renderer リファレンスを参照してください。
基礎となる MCP ツールのセットアップについて詳しくは、こちらをご覧ください。
サポートされました。送信 関連ドキュメントを読む
コンテキストとともに状態を AI SDK ツールとエージェントに渡す
AI SDK エージェントにスキルを追加する

## Original Extract

Build an MCP Apps host with the AI SDK using @ai-sdk/mcp and @ai-sdk/react to filter model-visible tools, read ui:// resources, and render interactive tool UIs in a sandboxed iframe with experimental_MCPAppRenderer.

Add MCP Apps to your AI SDK application | Vercel Knowledge Base Skip to content Copy Logo as SVG Copy Wordmark as SVG Brand Guidelines Docs Build Build with AI AI Gateway
Add MCP Apps to your AI SDK application
Build an MCP Apps host with the AI SDK using @ai-sdk/mcp and @ai-sdk/react to filter model-visible tools, read ui:// resources, and render interactive tool UIs in a sandboxed iframe with experimental_MCPAppRenderer.
Knowledge Base / AI SDK Copy URL Copy page Ask AI about this page 6 min read Last updated June 25, 2026 MCP Apps let a Model Context Protocol (MCP) tool return an interactive UI instead of plain text. The model still calls ordinary MCP tools, but a tool can point to a ui:// resource that holds HTML, and your app renders that HTML in a sandboxed iframe. To build the host side, the AI SDK provides @ai-sdk/mcp helpers for advertising MCP Apps support, filtering which tools the model sees, and reading ui:// resources, plus @ai-sdk/react components for rendering the iframe and bridging its messages. Your chat can then display a dashboard, form, or other interactive view generated by a tool, while the untrusted HTML remains isolated.
In this guide, you'll learn how to:
Connect to an MCP server with MCP Apps capabilities
Pass only model-visible tools to the model with splitMCPAppTools
Read a tool's ui:// resource with readMCPAppResource
Proxy app-initiated tool calls safely from the iframe
Render the app UI in your React chat with experimental_MCPAppRenderer
Before you begin, make sure you have:
The ai package with @ai-sdk/mcp and @ai-sdk/react
The MCP TypeScript SDK ( @modelcontextprotocol/sdk ) and a provider package, such as @ai-sdk/openai
An MCP server that exposes MCP Apps tools (tools that point to ui:// resources)
A React app that uses useChat (the examples below use the Next.js App Router)
An MCP Apps host connects to an MCP server, decides which tools the model can see, and renders any app UI that a tool points to. At runtime, the host follows these steps:
Connect to the MCP server with MCP Apps client capabilities.
List the server's tools and split them by MCP Apps visibility.
Pass only the model-visible tools to streamText or generateText .
Read a tool's ui:// resource when its tool part includes MCP App metadata.
Render the HTML resource in a sandboxed iframe.
Proxy allowed iframe requests, such as app-visible tool calls, back to the MCP server.
The rest of this guide builds each step.
Connect to the MCP server with MCP Apps support
Create the MCP client with mcpAppClientCapabilities so the host advertises that it can render text/html;profile=mcp-app resources.
app/api/chat/mcp-client.ts import { createMCPClient , mcpAppClientCapabilities } from '@ai-sdk/mcp' ; import { StreamableHTTPClientTransport } from '@modelcontextprotocol/sdk/client/streamableHttp.js' ;
export function createMCPAppsClient ( origin : string ) { return createMCPClient ( { transport : new StreamableHTTPClientTransport ( new URL ( '/mcp' , origin ) ) , clientName : 'my-mcp-apps-host' , capabilities : mcpAppClientCapabilities , } ) ; } Advertise these capabilities only if your host can safely fetch and render MCP App resources.
Expose only model-visible tools
MCP Apps tools can declare _meta.ui.visibility . Pass tools marked "model" to the model, and keep tools marked only "app" for iframe requests so the model never sees them. Split the tool list with splitMCPAppTools and pass modelVisible to streamText .
app/api/chat/route.ts import { splitMCPAppTools } from '@ai-sdk/mcp' ; import { convertToModelMessages , createUIMessageStreamResponse , streamText , toUIMessageStream , } from 'ai' ; import { createMCPAppsClient } from './mcp-client' ; import { openai } from '@ai-sdk/openai' ;
export async function POST ( req : Request ) { const requestUrl = new URL ( req . url ) ; const client = await createMCPAppsClient ( requestUrl . origin ) ; const { messages } = await req . json ( ) ;
try { const definitions = await client . listTools ( ) ; const { modelVisible } = splitMCPAppTools ( definitions ) ; const tools = client . toolsFromDefinitions ( modelVisible ) ;
const result = streamText ( { model : openai ( 'gpt-5.2' ) , tools , messages : await convertToModelMessages ( messages ) , onEnd : async ( ) => { await client . close ( ) ; } , } ) ;
return createUIMessageStreamResponse ( { stream : toUIMessageStream ( { stream : result . stream } ) , } ) ; } catch ( error ) { await client . close ( ) ; throw error ; } } When the model calls an app-backed tool, the MCP client keeps the app metadata in the tool UI, which the React renderer uses to determine whether a tool part has an MCP App.
Read a tool's ui:// resource with readMCPAppResource before you send it to the browser host.
app/api/mcp-app-host/read-resource/route.ts import { readMCPAppResource } from '@ai-sdk/mcp' ; import { createMCPAppsClient } from '../../chat/mcp-client' ;
export async function POST ( req : Request ) { const requestUrl = new URL ( req . url ) ; const { uri } = await req . json ( ) ; const client = await createMCPAppsClient ( requestUrl . origin ) ;
try { return Response . json ( await readMCPAppResource ( { client , uri } ) ) ; } finally { await client . close ( ) ; } } readMCPAppResource checks that the resource uses a ui:// URI, requires the MCP Apps MIME type, decodes text or base64 content, and returns the HTML along with rendering metadata such as its content security policy and permissions.
The iframe can't reach your MCP server directly. It sends JSON-RPC messages to your host, and your host decides which ones to allow. For an app-initiated tool call, confirm that the requested tool is app-visible before calling the MCP server.
app/api/mcp-app-host/call-tool/route.ts import { splitMCPAppTools } from '@ai-sdk/mcp' ; import { createMCPAppsClient } from '../../chat/mcp-client' ;
export async function POST ( req : Request ) { const requestUrl = new URL ( req . url ) ; const { name , arguments : toolArguments } = await req . json ( ) ; const client = await createMCPAppsClient ( requestUrl . origin ) ;
try { const { appVisible } = splitMCPAppTools ( await client . listTools ( ) ) ; const isAllowed = appVisible . tools . some ( tool => tool . name === name ) ;
if ( ! isAllowed ) { return Response . json ( { error : 'Tool is not app-visible' } , { status : 403 } , ) ; }
return Response . json ( await client . callTool ( { name , arguments : toolArguments ?? { } , } ) , ) ; } finally { await client . close ( ) ; } } In production, add any policy and user-approval checks your app needs before forwarding an iframe request.
Render the app in your React chat
In your React chat UI, render normal message parts as usual and pass tool parts to experimental_MCPAppRenderer .
experimental_MCPAppRenderer is experimental and may change in a future release.
app/page.tsx 'use client' ;
import { experimental_MCPAppRenderer as MCPAppRenderer , useChat , type MCPAppBridgeHandlers , type MCPAppMetadata , type MCPAppResource , type MCPAppSandboxConfig , } from '@ai-sdk/react' ; import { DefaultChatTransport , isToolUIPart } from 'ai' ;
const sandbox = { url : '/mcp-app-sandbox' , className : 'h-80 w-full rounded-lg border' , style : { border : 0 } , } satisfies MCPAppSandboxConfig ;
async function loadResource ( app : MCPAppMetadata ) : Promise < MCPAppResource > { const response = await fetch ( '/api/mcp-app-host/read-resource' , { method : 'POST' , body : JSON . stringify ( { uri : app . resourceUri } ) , } ) ;
if ( ! response . ok ) { throw new Error ( 'Failed to load MCP App resource' ) ; }
return response . json ( ) ; }
const handlers : MCPAppBridgeHandlers = { callTool : params => fetch ( '/api/mcp-app-host/call-tool' , { method : 'POST' , body : JSON . stringify ( params ) , } ) . then ( response => response . json ( ) ) , openLink : ( { url } ) => { window . open ( url , '_blank' , 'noopener,noreferrer' ) ; return { } ; } , } ;
export default function Chat ( ) { const { messages , sendMessage } = useChat ( { transport : new DefaultChatTransport ( { api : '/api/chat' } ) , } ) ;
return ( < > { messages . map ( message => message . parts . map ( ( part , index ) => { if ( part . type === 'text' ) { return < div key = { index } > { part . text } </ div > ; }
if ( isToolUIPart ( part ) ) { return ( < MCPAppRenderer key = { part . toolCallId } part = { part } loadResource = { loadResource } handlers = { handlers } sandbox = { sandbox } fallback = { < div > Loading MCP App... </ div > } /> ) ; }
return null ; } ) , ) }
< button onClick = { ( ) => sendMessage ( { text : 'Show me a dashboard' } ) } > Send </ button > </ > ) ; } experimental_MCPAppRenderer renders nothing for ordinary tools. For an app-backed tool, it loads the resource, creates the sandbox bridge, sends tool input and result notifications to the iframe, and
[truncated]
Treat MCP App HTML as untrusted. Render it in a sandboxed iframe, ideally through a sandbox proxy route on a separate origin.
Never pass app-only tools to the model. Use splitMCPAppTools and expose only the modelVisible tools.
Validate every iframe request on the server before you call client.callTool .
Cache resources by resourceUri so repeated tool calls don't refetch identical HTML.
Keep each tool's content and structuredContent useful on their own, so text-only hosts still work without the UI.
Close short-lived MCP clients when the response or host request finishes.
Read the MCP Apps helpers reference for the host-side functions.
See the MCP App Renderer reference for the React component's props.
Learn more about setting up the underlying MCP tools .
supported. Send Read related documentation
Pass state to AI SDK tools and agents with context
Add skills to your AI SDK agents
