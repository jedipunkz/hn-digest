---
source: "https://tanstack.com/blog/your-mcp-your-way"
hn_url: "https://news.ycombinator.com/item?id=48416038"
title: "TanStack AI"
article_title: "TanStack AI: Your MCP, your way | TanStack Blog"
author: "berlianta"
captured_at: "2026-06-05T18:45:41Z"
capture_tool: "hn-digest"
hn_id: 48416038
score: 2
comments: 0
posted_at: "2026-06-05T18:01:08Z"
tags:
  - hacker-news
  - translated
---

# TanStack AI

- HN: [48416038](https://news.ycombinator.com/item?id=48416038)
- Source: [tanstack.com](https://tanstack.com/blog/your-mcp-your-way)
- Score: 2
- Comments: 0
- Posted: 2026-06-05T18:01:08Z

## Translation

タイトル: TanStack AI
記事のタイトル: TanStack AI: あなたの MCP、あなたのやり方 |タンスタックのブログ
説明: ホスト側のモデル コンテキスト プロトコルのサポートが TanStack AI に導入されました。 1 つの MCP サーバーまたはプール全体に接続し、管理ライフサイクルまたは手動ライフサイクルを選択し、ゼロ構成の検出から完全に生成されたエンドツーエンドのタイプまでタイプ セーフティ レベルを選択します。

記事本文:
TanStack AI: あなたの MCP をあなたの方法で | TanStack ブログ TanStack 検索... K 自動ログイン Start RC Start RC Router Router Query Query Table Table DB beta DB beta AI alpha AI alpha Form new Form new Virtual Virtual Pacer beta Pacer beta ホットキー アルファ ホットキー アルファ ストア アルファ ストア アルファ Devtools アルファ Devtools アルファ CLI アルファ CLI アルファ インテント アルファ インテント アルファ その他のライブラリ その他のライブラリBuilder Alpha Builder Alpha Blog Blog Maintainers Maintainers Partners Partners Showcase Showcase Learn NEW Learn NEW Stats Stats YouTube YouTube Discord Discord Merch Merch Support Support Support GitHub GitHub Ethos Ethos Tenets Tenets ブランド ガイド ブランド ガイド ブログ このページについて TanStack AI: あなたの MCP、あなたのやり方
2026 年 6 月 5 日の Alem Tuzlak によるページをコピーします。
「現在 MCP をサポートしています」というアナウンスのほとんどは、MCP の使用方法を 1 つだけ示しています。サーバーに接続し、いくつかのツールを入手し、ライフサイクルがうまくいくことを祈ります。
@tanstack/ai-mcp は反対の立場を取ります。これはホスト側のモデル コンテキスト プロトコル クライアントで、任意の MCP サーバーを chat() に展開する通常の ServerTool[] に変えます。出力は単なるツールであるため、その上のすべてのレイヤー (アダプター (OpenAI、Anthropic、Gemini、Ollama)、エージェント ループ、フレームワーク統合など) は同じままです。 TanStack AI は、MCP が関与していたことを決して知りません。
この 1 つの設計上の決定により、1 台のサーバーでも 50 台のサーバーでも、フルマネージドまたは手動接続、型なしで高速、または生成されて厳密な、MCP の使用が可能になります。この投稿では、パッケージの正確なタイプを使用して、表面全体、すべてのフラグ、すべての構成を調べます。
公式 @modelcontextprotocol/sdk 上に構築されたランタイムは、引き続きエッジ デプロイ可能です。 Streamable HTTP トランスポートはnode: -free、Node-only stdioトランスポートはサブパスの背後に分離され、codegen CLIの重い依存関係はbinのみにバンドルされ、決してbinにはバンドルされません。

あなたが発送するライブラリ。
取り入れるべきアイデアが 1 つだけあります。
MCP クライアントはツール ファクトリです。 ServerTool[] が返されます。それらを chat({ tools }) に広げます。
「@tanstack/ai」から { チャット } をインポートします
import { openaiText } から '@tanstack/ai-openai/adapters'
import { createMCPClient } から '@tanstack/ai-mcp'
const mcp = await createMCPClient ({
トランスポート: { タイプ: 'http' 、URL: process.env。 MCP_URL ! }、
})
const stream = チャット ({
アダプター: openaiText ( 'gpt-5.5' )、
メッセージ、
ツール: mcp を待ちます。ツール ()、
})
MCPを待ちます。 close () import { chat } から ' @tanstack/ai '
import { openaiText } から ' @tanstack/ai-openai/adapters '
import { createMCPClient } から ' @tanstack/ai-mcp '
const mcp = await createMCPClient ( {
トランスポート: {タイプ: 'http'、URL: プロセス。環境 。 MCP_URL ! }、
})
const stream = チャット ( {
アダプター: openaiText ('gpt-5.5') 、
メッセージ、
ツール: mcp を待ちます。ツール () 、
})
mcp を待ちます。閉じる（）
この記事の他のすべての内容は、クライアントの構築方法、ツールの入力方法、ファンインするサーバーの数、および close() の所有者など、すべてその線上のバリエーションです。
フローチャート LR
S1[MCP サーバー] -->|callTool| C1[MCPClientの作成]
S2[MCP サーバー] --> P[MCPClients プールの作成]
S3[MCPサーバー] --> P
C1 -->|ServerTool 配列| CHAT["チャット({ツール})またはチャット({mcp})"]
P -->|接頭辞付きの ServerTool 配列|チャット
CLI[npx @tanstack/ai-mcp 生成] -.->|コンパイル時のタイプ| C1
CLI -.->|コンパイル時のタイプ| P
CHAT --> ADP[任意のアダプター: OpenAI / Anthropic / Gemini / Ollama]
MCP ツールの実行はサーバー側のみです。 createMCPClient はサーバー ルートまたはサーバーレス関数内に存在し、ブラウザー コード内には存在しません。
スタンドアロン クライアント: createMCPClient #
単一のクライアントが単一のサーバーに接続します。選択肢は少なく、どのフィールドにも負荷がかかります。
インターフェイス MCPClientOptions {
Transport : TransportInput // 設定オブジェクト

または未加工の SDK トランスポート
prefix ?: string // ツール名の接頭辞、例: 「github」 -> 「github_search」
name ?: string // サーバーに送信されるクライアント ID (デフォルトは 'tanstack-ai-mcp')
version ?: string // クライアントのバージョン (デフォルトは '0.0.1')
インターフェース MCPClientOptions {
Transport : TransportInput // 構成オブジェクトまたは生の SDK トランスポート
prefix ?: string // ツール名の接頭辞、例: 「github」 -> 「github_search」
name ?: string // サーバーに送信されるクライアント ID (デフォルトは 'tanstack-ai-mcp')
version ?: string // クライアントのバージョン (デフォルトは '0.0.1')
}
必須フィールドはトランスポートのみです。これは、以下の組み込み構成オブジェクトのいずれか、または既製の SDK Transport インスタンスのいずれかです。
prefix は、検出されたすべてのツール名を ${prefix}_${name} に書き換えます。あるサーバーのツール名が別のサーバーのツール名と衝突する可能性がある場合に使用します。 (これはプールによって設定されます - 以下を参照してください。)
名前とバージョンは、MCP ハンドシェイク中にサーバーに報告されるクライアント ID です。
返された MCPClient は、完全なプロトコル サーフェスを公開します。
インターフェイス MCPClient {
readonlyabilities : Record < string , known > // ハンドシェイクからのサーバー機能
tools ( options ? ) : Promise < ServerTool []> // 検出 (オーバーロードされています。以下を参照)
リソース () : プロミス < リソース []>
readResource ( uri : string ) : Promise < ReadResourceResult >
resourceTemplates () : Promise < ResourceTemplate []>
プロンプト () : プロミス < プロンプト []>
getPrompt ( name , args ? ) : Promise < GetPromptResult >
close () : 約束 < void >
[Symbol.asyncDispose]() : Promise < void > // `await using`
インターフェース MCPClient {
readonlyabilities : Record < string , known > // ハンドシェイクからのサーバー機能
tools ( options ? ) : Promise < ServerTool [] > // 検出 (オーバーロードされています。以下を参照)
リソース () : プロミス < リソース [] >
readResource ( uri : string ) : P

romise < ReadResourceResult >
resourceTemplates () : Promise < ResourceTemplate [] >
プロンプト () : プロミス < プロンプト [] >
getPrompt ( name , args ? ) : Promise < GetPromptResult >
close () : 約束 < void >
[ 記号 . asyncDispose ] () : Promise < void > // `await using`
}
トランスポート #
4 つの接続方法、1 つの一貫した形状。
HTTP (ストリーミング可能 HTTP) はリモート サーバーに推奨されるトランスポートであり、完全にエッジセーフである唯一のトランスポートです。
const mcp = await createMCPClient ({
トランスポート: {
タイプ: 'http' 、
URL: 'https://my-mcp-server.example.com/mcp' 、
ヘッダー: { 権限: `Bearer ${ process 。環境 。 MCP_TOKEN }` }、
fetch:customFetch, // オプション:独自のフェッチを持ち込む
authProvider: myOAuth, // オプション: OAuth 2.1 (「認証」を参照)
}、
}) const mcp = await createMCPClient ( {
トランスポート: {
タイプ:「 http 」、
URL : ' https://my-mcp-server.example.com/mcp ' 、
headers : { Authorization : ` Bearer ${ process .環境 。 MCP_TOKEN }` }、
fetch : customFetch , // オプション: 独自のフェッチを持ち込む
authProvider : myOAuth , // オプション: OAuth 2.1 (「認証」を参照)
}、
})
SSE は、従来の Server-Sent Events トランスポートをまだ実装しているサーバー用です。 HTTP と同じフィールド ( url 、 headers? 、 fetch? 、 authProvider? )。
const mcp = await createMCPClient ({
トランスポート: { タイプ: 'sse' 、url: 'https://example.com/sse' }、
}) const mcp = await createMCPClient ( {
トランスポート : { タイプ : ' sse ' 、 URL : ' https://example.com/sse ' }、
})
stdio はローカル MCP プロセスを生成します。 Node ネイティブ モジュールをインポートするため、@tanstack/ai-mcp/stdio サブパスの背後に分離されているため、エッジ バンドルはクリーンな状態に保たれます。
import { stdioTransport } から '@tanstack/ai-mcp/stdio'
import { createMCPClient } から '@tanstack/ai-mcp'
const mcp = await createMCPClient ({
トランスポート: stdioTransport ({
コマンド: 'ノード' 、
引数: [ './

my-mcp-server.js' ]、
env: { API_KEY: process.env. API_KEY ?? '' },
cwd: プロセス。 cwd (), // オプション
})、
}) ' @tanstack/ai-mcp/stdio ' から { stdioTransport } をインポートします
import { createMCPClient } から ' @tanstack/ai-mcp '
const mcp = await createMCPClient ( {
トランスポート : stdioTransport ( {
コマンド: 'ノード' 、
引数: [ ' ./my-mcp-server.js ' ] 、
env:{API_KEY:プロセス。環境 。 API_KEY ?? '' },
cwd:プロセス。 cwd () , // オプション
} )、
})
{ type: 'stdio' } 構成オブジェクトを createMCPClient に渡すと、サブパスを示すメッセージとともに意図的に直接スローされます。これにより、オプトインしない限り、ノードのみのコード パスがエッジ ビルドから除外されます。
カスタム トランスポートは避難口です。SDK トランスポート インスタンスを直接通過させます。 InMemoryTransport は、プロセス内テスト用に再エクスポートされます。
import { createMCPClient, InMemoryTransport } から '@tanstack/ai-mcp'
const [ clientTransport ] = InMemoryTransport。 createLinkedPair()
const mcp = await createMCPClient ({ Transport: clientTransport }) import { createMCPClient , InMemoryTransport } from ' @tanstack/ai-mcp '
const [ clientTransport ] = InMemoryTransport 。 createLinkedPair()
const mcp = await createMCPClient ( { トランスポート : clientTransport } )
このエスケープ ハッチは、インタラクティブな OAuth リダイレクト フローを処理する方法でもあります。つまり、自分で StreamableHTTPClientTransport を構築し、コールバック ルートで Transport.finishAuth(code) を呼び出せるように参照を保持してから、それを createMCPClient({ Transport }) に渡します。
2 つのパス。どちらも http / sse トランスポート構成で渡されます。
静的トークン: ヘッダーを設定すると、リクエストごとに送信されます。
OAuth 2.1 : authProvider を @modelcontextprotocol/sdk/client/auth.js の任意の OAuthClientProvider に設定します。次に、SDK トランスポートはトークンをアタッチして更新し、Tan に追加の配線を行わずに 401 で再試行します。

スタック側。
'@modelcontextprotocol/sdk/client/auth.js' からタイプ { OAuthClientProvider } をインポートします
const myOAuthProvider : OAuthClientProvider を宣言します
const mcp = await createMCPClient ({
トランスポート: {
タイプ: 'http' 、
URL: 'https://my-mcp-server.example.com/mcp' 、
認証プロバイダー: myOAuthProvider、
}、
}) ' @modelcontextprotocol/sdk/client/auth.js ' からタイプ { OAuthClientProvider } をインポートします
const myOAuthProvider : OAuthClientProvider を宣言します
const mcp = await createMCPClient ( {
トランスポート: {
タイプ:「 http 」、
URL : ' https://my-mcp-server.example.com/mcp ' 、
authProvider : myOAuthProvider 、
}、
})
タイプ セーフティの 3 つのモード #
ここで「あなたのやり方」が文字通り意味を持ちます。同じクライアントは 3 つのレベルの入力をサポートしており、呼び出しサイトごとに選択します。
モード 1 - 自動検出: client.tools() #
tools() を引数なしで呼び出して、サーバーが公開するすべてのツールを取得します。セットアップはありません。ツールの引数はコンパイル時には不明ですが、実行時にサーバーの JSON スキーマに対して検証されます。
const tools = mcp を待ちます。ツール ()
// ツール: ServerTool[] - 名前は既知ですが、引数は `unknown` と入力されます const tools = await mcp 。ツール ()
// ツール: ServerTool[] - 名前は既知ですが、引数は「unknown」と入力されます
知っておく価値のある 2 つの動作:
タスクベースのツールはスキップされます。 execution.taskSupport: 'required' を宣言するツールは、SDK の実験的なタスク フローを通じてのみ実行でき、プレーンな callTool では満たすことができません (サーバーは -32600 で拒否します)。 Discovery はこれらを除外するため、成功できないツールがモデルに提供されることはありません。
構造化された出力は保持されます。ツールが OutputSchema を宣言すると、クライアントはサーバーの StructuredContent を返すため、既存の出力検証は JSON-in-text BLOB ではなく型指定されたペイロードに対して実行されます。
モード 2 - 明示的な定義: client.tools([...defs]) #
TanStack toolDefinition() インスタンスを g に渡します

完全な TypeScript タイプと Zod 検証。これは許可リストです。指定されたツールのみが返されます。
import { toolsDefinition } から '@tanstack/ai'
「zod」から { z } をインポートします
const searchDef = ツール定義 ({
名前: '検索' 、
説明: 'アイテムを検索' ,
入力スキーマ: z.オブジェクト ({ クエリ: z.string () })、
出力スキーマ: z.配列 (z.オブジェクト ({ id: z. string (), title: z. string () }))、
})
const tools = mcp を待ちます。ツール ([searchDef])
// tools[0].execute は次のように入力されます: (args: { query: string }) => ... import { toolsDefinition } from ' @tanstack/ai '
' zod ' から { z } をインポートします
const searchDef = ツール定義 ( {
名前: '検索' 、
説明: 'アイテムを検索' 、
inputSchema : z 。オブジェクト ({ クエリ : z . string () } ) 、
出力スキーマ: z 。 array ( z . object ( { id : z . string () , title : z . string () } )) 、
})
const tools = await mcp 。ツール ([ searchDef ])
// tools[0].execute は次のように入力されます: (args: { query: string }) => ...
このパスを保護する 2 つのエラー:
定義の名前がサーバー上にない場合は MCPToolNotFoundError。
タスクに必要なツールを明示的にバインドした場合は MCPTaskRequiredToolError (この検出はサイレントにスキップされます)。
このモードは、既存のtoolDefinition()プリミティブを再利用します。パラは無いよ

[切り捨てられた]

## Original Extract

Host-side Model Context Protocol support lands in TanStack AI. Connect one MCP server or a whole pool, choose managed or manual lifecycle, and pick your type-safety level - from zero-config discovery to fully generated end-to-end types.

TanStack AI: Your MCP, your way | TanStack Blog TanStack Search... K Auto Log In Start RC Start RC Router Router Query Query Table Table DB beta DB beta AI alpha AI alpha Form new Form new Virtual Virtual Pacer beta Pacer beta Hotkeys alpha Hotkeys alpha Store alpha Store alpha Devtools alpha Devtools alpha CLI alpha CLI alpha Intent alpha Intent alpha More Libraries More Libraries Builder Alpha Builder Alpha Blog Blog Maintainers Maintainers Partners Partners Showcase Showcase Learn NEW Learn NEW Stats Stats YouTube YouTube Discord Discord Merch Merch Support Support GitHub GitHub Ethos Ethos Tenets Tenets Brand Guide Brand Guide Blog On this page TanStack AI: Your MCP, your way
Copy page by Alem Tuzlak on Jun 5, 2026.
Most "we support MCP now" announcements hand you exactly one way to use it. Connect to a server, get some tools, hope the lifecycle works out.
@tanstack/ai-mcp takes the opposite stance. It is a host-side Model Context Protocol client that turns any MCP server into ordinary ServerTool[] you spread into chat() . Because the output is just tools, every layer above it stays the same: any adapter (OpenAI, Anthropic, Gemini, Ollama), any agent loop, any framework integration. TanStack AI never knows MCP was involved.
That single design decision is what lets you use MCP your way : one server or fifty, fully managed or hand-wired, untyped-and-fast or generated-and-strict. This post walks the entire surface, every flag and every config, with the exact types from the package.
Built on the official @modelcontextprotocol/sdk , the runtime stays edge-deployable. The Streamable HTTP transport is node: -free, the Node-only stdio transport is isolated behind a subpath, and the codegen CLI's heavy dependencies are bundled into the bin only, never into the library you ship.
There is exactly one idea to internalize:
An MCP client is a tool factory. You get back ServerTool[] . You spread them into chat({ tools }) .
import { chat } from '@tanstack/ai'
import { openaiText } from '@tanstack/ai-openai/adapters'
import { createMCPClient } from '@tanstack/ai-mcp'
const mcp = await createMCPClient ({
transport: { type: 'http' , url: process.env. MCP_URL ! },
})
const stream = chat ({
adapter: openaiText ( 'gpt-5.5' ),
messages,
tools: await mcp. tools (),
})
await mcp. close () import { chat } from ' @tanstack/ai '
import { openaiText } from ' @tanstack/ai-openai/adapters '
import { createMCPClient } from ' @tanstack/ai-mcp '
const mcp = await createMCPClient ( {
transport : { type : ' http ' , url : process . env . MCP_URL ! },
} )
const stream = chat ( {
adapter : openaiText ( ' gpt-5.5 ' ) ,
messages ,
tools : await mcp . tools () ,
} )
await mcp . close ()
Everything else in this post is a variation on that line: how you build the client, how you type the tools, how many servers you fan in, and who owns close() .
flowchart LR
S1[MCP server] -->|callTool| C1[createMCPClient]
S2[MCP server] --> P[createMCPClients pool]
S3[MCP server] --> P
C1 -->|ServerTool array| CHAT["chat({ tools }) or chat({ mcp })"]
P -->|prefixed ServerTool array| CHAT
CLI[npx @tanstack/ai-mcp generate] -.->|compile-time types| C1
CLI -.->|compile-time types| P
CHAT --> ADP[any adapter: OpenAI / Anthropic / Gemini / Ollama]
MCP tool execution is server-side only . createMCPClient lives in a server route or serverless function, never in browser code.
Standalone clients: createMCPClient #
A single client connects to a single server. The options are small and every field is load-bearing.
interface MCPClientOptions {
transport : TransportInput // config object or a raw SDK Transport
prefix ?: string // tool-name prefix, e.g. 'github' -> 'github_search'
name ?: string // client identity sent to the server (default 'tanstack-ai-mcp')
version ?: string // client version (default '0.0.1')
} interface MCPClientOptions {
transport : TransportInput // config object or a raw SDK Transport
prefix ?: string // tool-name prefix, e.g. 'github' -> 'github_search'
name ?: string // client identity sent to the server (default 'tanstack-ai-mcp')
version ?: string // client version (default '0.0.1')
}
transport is the only required field. It is either one of the built-in config objects below or a ready-made SDK Transport instance.
prefix rewrites every discovered tool name to ${prefix}_${name} . Use it when one server's tool names might collide with another's. (Pools set this for you - see below.)
name and version are the client identity reported to the server during the MCP handshake.
The returned MCPClient exposes the full protocol surface:
interface MCPClient {
readonly capabilities : Record < string , unknown > // server capabilities from the handshake
tools ( options ? ) : Promise < ServerTool []> // discovery (overloaded, see below)
resources () : Promise < Resource []>
readResource ( uri : string ) : Promise < ReadResourceResult >
resourceTemplates () : Promise < ResourceTemplate []>
prompts () : Promise < Prompt []>
getPrompt ( name , args ? ) : Promise < GetPromptResult >
close () : Promise < void >
[Symbol.asyncDispose]() : Promise < void > // for `await using`
} interface MCPClient {
readonly capabilities : Record < string , unknown > // server capabilities from the handshake
tools ( options ? ) : Promise < ServerTool [] > // discovery (overloaded, see below)
resources () : Promise < Resource [] >
readResource ( uri : string ) : Promise < ReadResourceResult >
resourceTemplates () : Promise < ResourceTemplate [] >
prompts () : Promise < Prompt [] >
getPrompt ( name , args ? ) : Promise < GetPromptResult >
close () : Promise < void >
[ Symbol . asyncDispose ] () : Promise < void > // for `await using`
}
Transports #
Four ways to connect, one consistent shape.
HTTP (Streamable HTTP) is the preferred transport for remote servers and the only one that is fully edge-safe.
const mcp = await createMCPClient ({
transport: {
type: 'http' ,
url: 'https://my-mcp-server.example.com/mcp' ,
headers: { Authorization: `Bearer ${ process . env . MCP_TOKEN }` },
fetch: customFetch, // optional: bring your own fetch
authProvider: myOAuth, // optional: OAuth 2.1 (see Authentication)
},
}) const mcp = await createMCPClient ( {
transport : {
type : ' http ' ,
url : ' https://my-mcp-server.example.com/mcp ' ,
headers : { Authorization : ` Bearer ${ process . env . MCP_TOKEN }` },
fetch : customFetch , // optional: bring your own fetch
authProvider : myOAuth , // optional: OAuth 2.1 (see Authentication)
},
} )
SSE is for servers that still implement the legacy Server-Sent Events transport. Same fields as HTTP ( url , headers? , fetch? , authProvider? ).
const mcp = await createMCPClient ({
transport: { type: 'sse' , url: 'https://example.com/sse' },
}) const mcp = await createMCPClient ( {
transport : { type : ' sse ' , url : ' https://example.com/sse ' },
} )
stdio spawns a local MCP process. Because it imports Node-native modules, it is isolated behind the @tanstack/ai-mcp/stdio subpath so your edge bundles stay clean.
import { stdioTransport } from '@tanstack/ai-mcp/stdio'
import { createMCPClient } from '@tanstack/ai-mcp'
const mcp = await createMCPClient ({
transport: stdioTransport ({
command: 'node' ,
args: [ './my-mcp-server.js' ],
env: { API_KEY: process.env. API_KEY ?? '' },
cwd: process. cwd (), // optional
}),
}) import { stdioTransport } from ' @tanstack/ai-mcp/stdio '
import { createMCPClient } from ' @tanstack/ai-mcp '
const mcp = await createMCPClient ( {
transport : stdioTransport ( {
command : ' node ' ,
args : [ ' ./my-mcp-server.js ' ] ,
env : { API_KEY : process . env . API_KEY ?? '' },
cwd : process . cwd () , // optional
} ) ,
} )
Passing a { type: 'stdio' } config object to createMCPClient directly throws on purpose, with a message pointing you at the subpath. That keeps the Node-only code path out of edge builds unless you opt in.
Custom transport is the escape hatch: pass any SDK Transport instance straight through. InMemoryTransport is re-exported for in-process testing.
import { createMCPClient, InMemoryTransport } from '@tanstack/ai-mcp'
const [ clientTransport ] = InMemoryTransport. createLinkedPair ()
const mcp = await createMCPClient ({ transport: clientTransport }) import { createMCPClient , InMemoryTransport } from ' @tanstack/ai-mcp '
const [ clientTransport ] = InMemoryTransport . createLinkedPair ()
const mcp = await createMCPClient ( { transport : clientTransport } )
This escape hatch is also how you handle interactive OAuth redirect flows: build a StreamableHTTPClientTransport yourself, keep a reference so you can call transport.finishAuth(code) in your callback route, then hand it to createMCPClient({ transport }) .
Two paths, both passed on the http / sse transport config.
Static tokens : set headers and they are sent with every request.
OAuth 2.1 : set authProvider to any OAuthClientProvider from @modelcontextprotocol/sdk/client/auth.js . The SDK transport then attaches tokens, refreshes them, and retries on 401 with no extra wiring on the TanStack side.
import type { OAuthClientProvider } from '@modelcontextprotocol/sdk/client/auth.js'
declare const myOAuthProvider : OAuthClientProvider
const mcp = await createMCPClient ({
transport: {
type: 'http' ,
url: 'https://my-mcp-server.example.com/mcp' ,
authProvider: myOAuthProvider,
},
}) import type { OAuthClientProvider } from ' @modelcontextprotocol/sdk/client/auth.js '
declare const myOAuthProvider : OAuthClientProvider
const mcp = await createMCPClient ( {
transport : {
type : ' http ' ,
url : ' https://my-mcp-server.example.com/mcp ' ,
authProvider : myOAuthProvider ,
},
} )
Three modes of type safety #
This is where "your way" gets literal. The same client supports three levels of typing, and you pick per call site.
Mode 1 - Auto-discovery: client.tools() #
Call tools() with no arguments to get every tool the server exposes. No setup. Tool arguments are unknown at compile time and validated at runtime against the server's JSON Schema.
const tools = await mcp. tools ()
// tools: ServerTool[] - names known, args typed `unknown` const tools = await mcp . tools ()
// tools: ServerTool[] - names known, args typed `unknown`
Two behaviors worth knowing:
Task-based tools are skipped. A tool that declares execution.taskSupport: 'required' can only run through the SDK's experimental task flow, which plain callTool cannot satisfy (the server rejects it with -32600 ). Discovery filters these out so the model is never offered a tool that cannot succeed.
Structured output is preserved. When a tool declares an outputSchema , the client returns the server's structuredContent so your existing output validation runs against the typed payload instead of a JSON-in-text blob.
Mode 2 - Explicit definitions: client.tools([...defs]) #
Pass TanStack toolDefinition() instances to get full TypeScript types and Zod validation. This is an allowlist : only the named tools come back.
import { toolDefinition } from '@tanstack/ai'
import { z } from 'zod'
const searchDef = toolDefinition ({
name: 'search' ,
description: 'Search for items' ,
inputSchema: z. object ({ query: z. string () }),
outputSchema: z. array (z. object ({ id: z. string (), title: z. string () })),
})
const tools = await mcp. tools ([searchDef])
// tools[0].execute is typed: (args: { query: string }) => ... import { toolDefinition } from ' @tanstack/ai '
import { z } from ' zod '
const searchDef = toolDefinition ( {
name : ' search ' ,
description : ' Search for items ' ,
inputSchema : z . object ( { query : z . string () } ) ,
outputSchema : z . array ( z . object ( { id : z . string () , title : z . string () } )) ,
} )
const tools = await mcp . tools ([ searchDef ])
// tools[0].execute is typed: (args: { query: string }) => ...
Two errors guard this path:
MCPToolNotFoundError if a definition's name is not on the server.
MCPTaskRequiredToolError if you explicitly bind a task-required tool (which discovery would have silently skipped).
This mode reuses the existing toolDefinition() primitive. There is no para

[truncated]
