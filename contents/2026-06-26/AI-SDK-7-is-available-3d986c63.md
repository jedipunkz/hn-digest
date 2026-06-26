---
source: "https://vercel.com/blog/ai-sdk-7"
hn_url: "https://news.ycombinator.com/item?id=48685133"
title: "AI SDK 7 is available"
article_title: "AI SDK 7 is now available - Vercel"
author: "m3h"
captured_at: "2026-06-26T10:59:42Z"
capture_tool: "hn-digest"
hn_id: 48685133
score: 1
comments: 0
posted_at: "2026-06-26T10:59:01Z"
tags:
  - hacker-news
  - translated
---

# AI SDK 7 is available

- HN: [48685133](https://news.ycombinator.com/item?id=48685133)
- Source: [vercel.com](https://vercel.com/blog/ai-sdk-7)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T10:59:01Z

## Translation

タイトル: AI SDK 7 が利用可能になりました
記事タイトル: AI SDK 7 が利用可能になりました - Vercel
説明: AI SDK は、あらゆるモデル プロバイダーにわたって AI アプリケーション、機能、フレームワーク、エージェントを構築するための TypeScript SDK です。 AI SDK 7 は、本番環境で AI を実行するために必要なものに焦点を当てています。

記事本文:
AI SDK 7 が利用可能になりました - Vercel コンテンツにスキップ ロゴを SVG としてコピー ワードマークを SVG としてコピー ブランド ガイドライン 製品 エージェント スタック AI SDK
AI SDK は毎週 1,600 万件以上ダウンロードされており、あらゆるモデル プロバイダーにわたって AI アプリケーション、機能、フレームワーク、エージェントを構築するための TypeScript SDK です。これは、Vercel のオープンソース エージェント フレームワークである eve が構築されているのと同じレイヤーです。
AI SDK 7 は、次の 5 つの領域にわたるエージェントの作業にプロダクションの深さを追加します。
推論制御、ツールとランタイム コンテキスト、プロバイダー ファイルとスキル サポート、MCP アプリ、ターミナル UI を備えたエージェントを開発します。
ツールの承認、耐久性 ( WorkflowAgent )、タイムアウト、サンドボックスのサポートを備えたエージェントを実行します。
Codex、Claude Code、Deep Agents、OpenCode、Pi などのエージェント ハーネスを統合します。
テレメトリ、Node.js トレース チャネル、ライフサイクル イベント、パフォーマンス統計を使用してエージェントを観察します。
プロバイダーに依存しないリアルタイム音声サポートとビデオ生成により、テキスト エージェントを超えます。
pnpm add ai@latest AI SDK 7 をインストールします
AI SDK 6 からアップグレードしますか? npx @ai-sdk/codemod v7 を実行して最小限のコード変更で自動的に移行するか、移行スキルを使用します: npx skill add vercel/ai --skill merge-ai-sdk-v6-to-v7
見出し「開発エージェント」へのリンク
適切に動作するエージェントを構築するには、モデル推論、ツール コンテキスト、およびファイル処理をきめ細かく制御する必要があります。
見出し「推論制御」へのリンク
ほとんどのフロンティア モデルは構成可能な推論をサポートしていますが、プロバイダー API ごとにそれを異なる方法で公開します。
AI SDK 7 では、generateText および streamText の推論オプションを使用してこれを標準化しています。これはプロバイダーネイティブの推論設定にマッピングされ、推論の労力を 1 行で制御できるようになります。より詳細なプロバイダー固有の推論構成が必要な場合は、プロバイダー オプションにフォールバックすることもできます。
Agent.ts 1 import {generateText} から

m'ai'; 2 const result = await generatedText ({ 3 モデル , 4 プロンプト , 5 推論 : 'high' , 6 } ) ;単一のオプションで推論の労力を設定する
詳細については、推論に関するドキュメントを参照してください。
ツールは、特定のエージェントやアプリケーションから独立して開発されることが増えています。たとえば、サードパーティ企業は、エージェントが自社の API を使用できるようにするツールを提供しています。したがって、ツールには、API キーや構成設定など、LLM によって生成されない追加の入力が必要です。
AI SDK 7 では、スキーマを介して各ツールに指定できる完全に型指定されたツール コンテキストが追加されています。サードパーティ ツールが不要なコンテキストにアクセスするのを防ぐために、コンテキストはツールに限定されます。
Agent.ts 1 const Agent = new ToolLoopAgent ({ 2 モデル , 3 ツール : { 4 天気 : ツール ( { 5 説明 , 6 inputSchema , 7 contextSchema : z . object ( { 8 apiKey : z . string ( ) , 9 } ) , 10 実行 : async ( input , { context : { apiKey } } ) => { 11 // ... 12 } , 13 } ) , 14 } , 15 toolsContext : { 16 天気 : { apiKey : プロセス . WEATHER_API_KEY } , 17 } , 18 } ) ; API キーのスコープを、それを必要とするツールに限定する
ランタイムコンテキストの見出しへのリンク
より複雑なエージェント ループの場合、多くの場合、プロンプトやモデルの選択などを調整するために、prepareStep でアクセスして変更できる変数が必要になります。
AI SDK 7 では、オプションのテレメトリ サポートとともに、ステップの準備およびツールの承認機能中に使用できる型付きランタイム コンテキストが導入されています。これにより、ToolLoopAgent でより多くのロジックをカプセル化し、それらのエージェントをその内部ロジックと共有できるようになります。
Agent.ts 1 const Agent = new ToolLoopAgent ( { 2 // ランタイム コンテキストをセットアップします 3 runtimeContext : { 4 var1 : "something" , 5 } , 6 prepareStep : async ( { runtimeContext ,steps } ) => { 7 // ランタイム コンテキストを使用します 8 // 更新されたランタイム コンテキストを返します 9 } , 10 } ) ;にアクセスする

d ステップ全体で型付き変数を更新する
ランタイム コンテキストについて詳しくは、こちらをご覧ください。
ヘッダーのプロバイダー ファイルのアップロードへのリンク
多くのエージェント ワークフローでは、PDF、画像、データセット、その他のアーティファクトなどの大規模な入力を処理する必要があります。これらのファイルをインラインで送信すると、特にステートレス推論の場合は何度も送信されるため、時間がかかり無駄が生じます。
AI SDK 7 では、ファイルを一度アップロードすると、後続のモデル呼び出しに軽量の参照を渡すことができる最上位の UploadFile API が追加されています。これにより、同じバイトを繰り返し再アップロードすることが回避され、推論が高速化され、繰り返し実行または複数ステップの実行中に帯域幅が節約されます。
UploadFile は、ファイル アップロード エンドポイントを提供する任意のプロバイダーで使用できます。この関数は、プロバイダ間で移植可能なプロバイダ参照オブジェクトを返します。
Upload.ts 1 const { ProviderReference } = await UploadFile ( { 2 API : openai . files ( ) , 3 データ : readFileSync ( './photo.png' ) , 4 ファイル名 : 'photo.png' , 5 } ) ; 6 const result = await streamText ( { 7 モデル : openai .response ( 'gpt-5.5' ) , 8 メッセージ : [ 9 { 10 役割 : 'ユーザー' , 11 コンテンツ : [ 12 { type : 'text' , text : 'この画像に表示されている内容を説明してください。' } , 13 { type : 'file' , mediaType : 'image' , データ:providerReference } , 14 ] , 15 } , 16 ] , 17 } ) ;ファイルを一度アップロードし、後続のモデル呼び出しに参照を渡します
プロバイダー ファイルのアップロードの詳細については、こちらをご覧ください。
見出し「プロバイダー スキルのアップロード」へのリンク
プロバイダーが管理するコンテナー環境へのすべてのリクエストでスキルをインラインで送信すると、ファイルをインラインで送信する場合と同じオーバーヘッドの問題が発生します。
AI SDK 7 では、トップレベルの UploadSkill API が追加されており、スキルを一度アップロードすると、その後の推論呼び出しでその参照を使用できるようになります。 UploadFile と同様に、この関数はプロバイダー参照オブジェクトを返します。
Upload.ts 1 const { ProviderReference } = あなたを待ちます

ploadSkill ( { 2 api : anthropic . skill ( ) , 3 files : [ 4 { 5 path : 'my-skill/SKILL.md' , 6 content : readFileSync ( './SKILL.md' ) , 7 } , 8 ] , 9 displayTitle : 'My Skill' , 10 } ) ; 11 const result = await streamText ( { 12 model : anthropic ( 'claude-sonnet-4-6' ) , 13 tools : { 14 code_execution : anthropic . tools . codeExecution_20260120 ( ) , 15 } , 16 Prompt : 'タスクを完了するには、my-skill スキルを使用してください。' , 17 ProviderOptions : { 18 anthropic : { 19 container : { 20 skill : [ { type : 'custom' , ProviderReference } ] , 21 } , 22 } は AnthropicLanguageModelOptions , 23 } , 24 } ) を満たします。スキルを一度アップロードすると、推論呼び出し全体で参照できます
プロバイダー スキルのアップロードの詳細については、こちらをご覧ください。
MCP は、エージェントをツールやリソースに接続する一般的な方法となっています。ただし、すべてのツールがモデルに表示される必要があるわけではなく、一部の MCP サーバーはツールと一緒に特殊な UI を公開する必要があります。
AI SDK 7 では、MCP アプリのサポートが追加されています。 MCP サーバーは、モデルに表示されるツールをアプリのみのツールから分離し、アプリのメタデータを保持し、サンドボックス化された iframe 内でアプリ UI をレンダリングできるようになりました。 JSON-RPC ブリッジは、ツール、リソース、表示操作を接続します。
これにより、モデルが必要なツールを使用できる一方で、ユーザーにはレビュー、構成、または対話のためのアプリ固有のインターフェイスが表示される、より充実したエージェント エクスペリエンスを構築できます。
コンポーネント/chat.tsx 1 import {experimental_MCPAppRenderer as MCPAppRenderer} from '@ai-sdk/react' ; 2 'ai' から { isToolUIPart } をインポートします。 3 { 4 つのメッセージ。 map ( message => 5 message .parts .map (part => 6 isToolUIPart (part) ?( 7 < MCPAppRenderer 8 key = {part .toolCallId } 9 part = {part } 10 Sandbox = { { url : '/mcp-app-sandbox' , className : 'h-96 w-full' } } 11 loadResource = { app => fetch ( ` /api/mcp-apps?uri= ${ app . resourceUri } ` ) } 12 ハンドラー = { {

allowedTools : [ 'refreshDashboard' ] } } 13 / > 14 ) : null 、 15 ) 、 16 ) ; 17 } モデル出力と一緒に MCP アプリ UI をレンダリングする
AI SDK を使用して最初の MCP アプリの構築を今すぐ始めましょう。
エージェントを開発するときは、完全なアプリを作成せずにエージェントを迅速にテストできる必要があります。 AI SDK 7 には、わずか数行のコードでエージェントを実行できるターミナル UI (TUI) パッケージが追加されています。
TUI は対話型で、推論とツールをサポートし、マークダウンをフォーマットされたテキストとしてレンダリングします。
dev.ts 1 import { runAgentTUI } から '@ai-sdk/tui' ; 2 await runAgentTUI({エージェント}) ;ターミナルでエージェントを実行する
独自のターミナル エージェントの作成について詳しくは、こちらをご覧ください。
エージェントの自律性が高まり、実行時間が長くなると、承認、耐久性、サンドボックス化、堅牢性の必要性が高まります。
見出し「ツールの承認」へのリンク
AI SDK 7 は、次の承認タイプで、自動または人間が関与するエージェント レベルのツール承認をサポートします。
特定のツールに対する簡単なユーザー承認。
自動承認、自動拒否、またはユーザー承認に転送できる特定のツールのツール承認機能。
汎用の包括的なツール承認機能。
特定のツールの使用シナリオによって承認の必要性が高まるため、ツールの承認は ToolLoopAgent 、generateText 、および streamText で定義されます。
Agent.ts 1 const Agent = new ToolLoopAgent ({ 2 モデル , 3 ツール : { 天気 : WeatherTool } , 4 toolApproval : { 5 天気 : 'ユーザー承認' , 6 } , 7 } ) ;ツールを実行する前にユーザーの承認を求める
リスクの高いワークフロー向けに、AI SDK 7 では、承認の偽造を防止するために、オプトインの HMAC 署名付きツール承認が導入されています。また、SDK は、実行を続行する前にツール入力とポリシーを再検証することにより、再生動作を強化します。
WorkflowAgent (耐久性) という見出しへのリンク
エージェントの実行が複数のステップにまたがる場合、または人間の操作を待機する場合

承認、プロセスの再起動、または実行途中でのデプロイメントは、最初からやり直すことを意味します。 AI SDK 7 では、@ai-sdk/workflow と WorkflowAgent を導入し、プロセスの再起動、デプロイ、中断、承認の遅れにも耐える耐久性と再開可能なエージェント実行を実現します。
WorkflowAgent は、ワークフロー ステップの境界を越えたワークフロー ベースのストリーミング、ツール、承認、コールバック、 prepareCall 、およびプロバイダー モデルのシリアル化をサポートします。また、共有エージェント状態と安定したテレメトリのための型付きランタイム コンテキストもサポートします。
コールバックには、ステップ番号、以前の結果、期間、成功または失敗の情報など、より豊富な実行データが含まれるようになりました。無効なツール呼び出しは、無効なツールを実行せずに保存され、ツールから ModelOutput への変換により、UI とコールバックの生の出力を保存できます。
WorkflowAgent を使用してエージェントを構築する方法を学習します。
エージェントは、単純なリクエストよりもさまざまな方法で停止する可能性があります。プロバイダーがストリームを開いてチャンクの送信を停止したり、ツールがハングしたり、複数ステップの実行が総バジェットを超えたりする可能性があります。
AI SDK 7 では、合計、ステップごと、チャンクごと、ツールごとの制限を含む、テキスト生成とエージェント API にわたるファーストクラスのタイムアウト構成が追加されています。タイムアウト中止には TimeoutError が使用され、中止理由はストリームおよび UI プロトコルを通じて伝播されます。
Agent.ts 1 const result = await generatedText ({ 2 model , 3 tools : {weather :weatherTool , throwApi : throwApiTool } , 4 timeout : { 5 totalMs : 60000 , // 合計 60 秒 6 stepMs : 10000 , // ステップごとに 10 秒 7 chunkMs : 2000 , // 2 秒間チャンクを受信しない場合は中止します 8 toolMs : 5000 , // すべてのツールのデフォルト 9 tools : { 10 WeatherMs : 3000 , // 天気ツールの場合は 3 秒 11 throwApiMs : 10000 , // 低速 API ツールの場合は 10 秒 12 } , 13 } , 14 プロンプト : '天気は何ですかサンフランシスコで？ 、15 } ;合計、ステップごと、および

ツールごとのタイムアウト制限
サンドボックスのサポートの見出しへのリンク
シェル コマンドの実行、ファイルの読み取りと書き込み、または生成されたコードの実行を行うエージェントには、一貫した実行環境が必要ですが、基盤となるサンドボックスは、ローカルの開発、CI、運用環境全体で変更されることがよくあります。 AI SDK 7 では、ツールやエージェントでのポータブルなコマンド実行のためのファーストクラスの SandboxSession 抽象化が追加されています。ツールは特定のサンドボックスとは独立して開発でき、任意のサンドボックス プロバイダーで任意のサンドボックス対応ツールを使用できます。
Vercel Sandbox などのサンドボックス環境は、この目的に最適です。
見出し「エージェント ハーネスの統合」へのリンク
エージェント ランタイムは単一のアプリケーション サーバーを超えて移行しています。チームは、コーディング環境、ホストされたサンドボックス、ローカル セッション、サードパーティ ハーネス内で同じエージェント ロジックを実行したいと考えています。
AI SDK 7 では、実験的なハーネス抽象化と HarnessAgent が導入されています。これは、Claude Code、Codex、Pi など、完全に構成され確立されたエージェント ハーネスを実行するための 1 つの API です。ハーネスは、動作するサンドボックス、カスタムの指示、スキル、ツールを使用して構成できます。確立されたハーネスを一貫したインターフェースを通じて実行し、それぞれを個別に構成し、統合を変更せずに 1 つを交換します。

[切り捨てられた]

## Original Extract

AI SDK is the TypeScript SDK for building AI applications, features, frameworks, and agents across any model provider. AI SDK 7 focuses on what it takes to run AI in production.

AI SDK 7 is now available - Vercel Skip to content Copy Logo as SVG Copy Wordmark as SVG Brand Guidelines Products Agent Stack AI SDK
AI SDK, with over 16 million weekly downloads, is the TypeScript SDK for building AI applications, features, frameworks, and agents across any model provider. It's the same layer eve , Vercel's open-source agent framework, is built on.
AI SDK 7 adds production depth for agent work across five areas:
Develop agents with reasoning control, tool and runtime context, provider files and skills support, MCP Apps, and a terminal UI.
Run agents with tool approvals, durability ( WorkflowAgent ), timeouts, and sandbox support.
Integrate any agent harness, such as Codex, Claude Code, Deep Agents, OpenCode, or Pi.
Observe agents with telemetry, Node.js tracing channel, lifecycle events, and performance statistics.
Go beyond text agents with provider-agnostic real-time voice support and video generation.
pnpm add ai@latest Install AI SDK 7
Upgrading from AI SDK 6? Run npx @ai-sdk/codemod v7 to migrate automatically with minimal code changes, or use the migration skill: npx skills add vercel/ai --skill migrate-ai-sdk-v6-to-v7
Link to heading Develop agents
Building well-behaved agents requires fine-grained control over model reasoning, tool context, and file handling.
Link to heading Reasoning control
Most frontier models support configurable reasoning, but every provider API exposes it differently.
AI SDK 7 standardizes this with a reasoning option for generateText and streamText . It maps to provider-native reasoning settings, letting you control reasoning effort in a single line. You can also still fall back to provider options when you need more detailed provider-specific reasoning configuration.
agent.ts 1 import { generateText } from 'ai' ; 2 const result = await generateText ( { 3 model , 4 prompt , 5 reasoning : 'high' , 6 } ) ; Setting reasoning effort with a single option
Learn more in the reasoning documentation .
Tools are increasingly developed independently of specific agents or applications. For example, third-party companies offer tools that enable agents to use their APIs. Therefore, tools require additional inputs that are not generated by LLMs, such as API keys or configuration settings.
AI SDK 7 adds a fully typed tool context that can be specified for each tool via a schema. The context is limited to the tool to prevent 3rd-party tools from accessing context they do not need.
agent.ts 1 const agent = new ToolLoopAgent ( { 2 model , 3 tools : { 4 weather : tool ( { 5 description , 6 inputSchema , 7 contextSchema : z . object ( { 8 apiKey : z . string ( ) , 9 } ) , 10 execute : async ( input , { context : { apiKey } } ) => { 11 // ... 12 } , 13 } ) , 14 } , 15 toolsContext : { 16 weather : { apiKey : process . env . WEATHER_API_KEY ! } , 17 } , 18 } ) ; Scoping an API key to the tool that needs it
Link to heading Runtime context
For more complex agentic loops, you often need variables that you can access and modify in prepareStep to adjust prompts, model selection, and more.
AI SDK 7 introduces a typed runtime context available during step preparation and tool approval functions, with optional telemetry support. This enables you to encapsulate more logic in ToolLoopAgent and share those agents with that internal logic.
agent.ts 1 const agent = new ToolLoopAgent ( { 2 // setup runtime context 3 runtimeContext : { 4 var1 : "something" , 5 } , 6 prepareStep : async ( { runtimeContext , steps } ) => { 7 // use runtime context 8 // return updated runtime context 9 } , 10 } ) ; Accessing and updating typed variables across steps
Learn more about Runtime Context .
Link to heading Provider file uploads
Many agent workflows require handling large inputs, such as PDFs, images, datasets, or other artifacts. Sending those files inline is slow and wasteful, especially for stateless inference, where they get sent over and over again.
AI SDK 7 adds a top-level uploadFile API that lets you upload a file once and then pass a lightweight reference into subsequent model calls. This avoids re-uploading the same bytes repeatedly, making inference faster and saving bandwidth during repeated or multi-step runs.
uploadFile can be used with any providers that offer a file uploading endpoint. The function returns a provider reference object that is portable across providers.
upload.ts 1 const { providerReference } = await uploadFile ( { 2 api : openai . files ( ) , 3 data : readFileSync ( './photo.png' ) , 4 filename : 'photo.png' , 5 } ) ; 6 const result = await streamText ( { 7 model : openai . responses ( 'gpt-5.5' ) , 8 messages : [ 9 { 10 role : 'user' , 11 content : [ 12 { type : 'text' , text : 'Describe what you see in this image.' } , 13 { type : 'file' , mediaType : 'image' , data : providerReference } , 14 ] , 15 } , 16 ] , 17 } ) ; Upload a file once, pass a reference into subsequent model calls
Learn more about Provider File Uploads
Link to heading Provider skill uploads
Sending skills inline on every request to provider-managed container environments has the same overhead problem as sending files inline.
AI SDK 7 adds a top-level uploadSkill API that lets you upload a skill once and then use a reference to it in subsequent inference calls. Similar to uploadFile , the function returns a provider reference object.
upload.ts 1 const { providerReference } = await uploadSkill ( { 2 api : anthropic . skills ( ) , 3 files : [ 4 { 5 path : 'my-skill/SKILL.md' , 6 content : readFileSync ( './SKILL.md' ) , 7 } , 8 ] , 9 displayTitle : 'My Skill' , 10 } ) ; 11 const result = await streamText ( { 12 model : anthropic ( 'claude-sonnet-4-6' ) , 13 tools : { 14 code_execution : anthropic . tools . codeExecution_20260120 ( ) , 15 } , 16 prompt : 'Use the my-skill skill to complete the task.' , 17 providerOptions : { 18 anthropic : { 19 container : { 20 skills : [ { type : 'custom' , providerReference } ] , 21 } , 22 } satisfies AnthropicLanguageModelOptions , 23 } , 24 } ) ; Upload a skill once, reference it across inference calls
Learn more about Provider Skill Uploads .
MCP has become a common way to connect agents to tools and resources. But not every tool should be model-visible, and some MCP servers need to expose specialized UI alongside their tools.
AI SDK 7 adds support for MCP Apps. MCP servers can now separate model-visible tools from app-only tools, preserve app metadata, and render app UIs inside sandboxed iframes. A JSON-RPC bridge connects tools, resources, and display interactions.
This lets you build richer agent experiences where the model can use the tools it needs, while the user sees an app-specific interface for review, configuration, or interaction.
components/chat.tsx 1 import { experimental_MCPAppRenderer as MCPAppRenderer } from '@ai-sdk/react' ; 2 import { isToolUIPart } from 'ai' ; 3 { 4 messages . map ( message => 5 message . parts . map ( part => 6 isToolUIPart ( part ) ? ( 7 < MCPAppRenderer 8 key = { part . toolCallId } 9 part = { part } 10 sandbox = { { url : '/mcp-app-sandbox' , className : 'h-96 w-full' } } 11 loadResource = { app => fetch ( ` /api/mcp-apps?uri= ${ app . resourceUri } ` ) } 12 handlers = { { allowedTools : [ 'refreshDashboard' ] } } 13 / > 14 ) : null , 15 ) , 16 ) ; 17 } Rendering MCP app UIs alongside model output
Start building your first MCP App with AI SDK today.
When developing agents, you need to be able to quickly test them without writing a full app. AI SDK 7 adds a terminal UI (TUI) package that lets you run agents with just a few lines of code:
The TUI is interactive, supports reasoning and tools, and renders markdown as formatted text.
dev.ts 1 import { runAgentTUI } from '@ai-sdk/tui' ; 2 await runAgentTUI ( { agent } ) ; Running an agent in the terminal
Learn more about creating your own terminal agent .
As agents become more autonomous and longer running, the need for approvals, durability, sandboxing, and robustness increases.
Link to heading Tool approvals
AI SDK 7 supports agent-level tool approvals that can be automatic or involve a human in the loop, with these approval types:
Simple user-approval for particular tools.
Tool approval function for a particular tool that can auto-approve, auto-deny, or forward to user approval.
Generic catch-all tool approval functions.
Tool approvals are defined on ToolLoopAgent , generateText , and streamText , because the usage scenario of a particular tool drives the need for approvals.
agent.ts 1 const agent = new ToolLoopAgent ( { 2 model , 3 tools : { weather : weatherTool } , 4 toolApproval : { 5 weather : 'user-approval' , 6 } , 7 } ) ; Requiring user approval before a tool executes
For higher-risk workflows, AI SDK 7 introduces opt-in HMAC-signed tool approvals to prevent forged approvals. The SDK also hardens replay behavior by revalidating tool inputs and policies before continuing execution.
Link to heading WorkflowAgent (Durability)
When an agent run spans multiple steps or waits for a human approval, a process restart or deployment in the middle of that run means starting over. AI SDK 7 introduces @ai-sdk/workflow and WorkflowAgent for durable, resumable agent execution that survives process restarts, deploys, interruptions, and delayed approvals.
WorkflowAgent supports workflow-based streaming, tools, approvals, callbacks, prepareCall , and provider model serialization across workflow step boundaries. It also supports typed runtime context for shared agent state and stable telemetry.
Callbacks now include richer execution data such as step numbers, previous results, duration, and success or failure information. Invalid tool calls are preserved without executing invalid tools, and tool toModelOutput conversion can preserve raw outputs for UI and callbacks.
Learn how to build an agent with WorkflowAgent .
Agents can stall in more ways than a simple request can: a provider can open a stream and stop sending chunks, a tool can hang, or a multi-step run can exceed its total budget.
AI SDK 7 adds first-class timeout configuration across text generation and agent APIs, including total, per-step, per-chunk, and per-tool limits. Timeout aborts use TimeoutError , and abort reasons propagate through stream and UI protocols.
agent.ts 1 const result = await generateText ( { 2 model , 3 tools : { weather : weatherTool , slowApi : slowApiTool } , 4 timeout : { 5 totalMs : 60000 , // 60 seconds total 6 stepMs : 10000 , // 10 seconds per step 7 chunkMs : 2000 , // abort if no chunk received for 2 seconds 8 toolMs : 5000 , // default for all tools 9 tools : { 10 weatherMs : 3000 , // 3 seconds for weather tool 11 slowApiMs : 10000 , // 10 seconds for slow API tool 12 } , 13 } , 14 prompt : 'What is the weather in San Francisco?' , 15 } ) ; Configuring total, per-step, and per-tool timeout limits
Link to heading Sandbox support
Agents that run shell commands, read and write files, or execute generated code need a consistent execution environment, but the underlying sandbox often changes across local dev, CI, and production. AI SDK 7 adds a first-class SandboxSession abstraction for portable command execution in tools and agents. Tools can be developed independently of any particular sandbox, and you can use any sandbox-aware tool with any sandbox provider.
Sandboxed environments, such as Vercel Sandbox , are ideal for this purpose.
Link to heading Integrate any agent harness
Agent runtimes are moving beyond a single application server. Teams want to run the same agent logic inside coding environments, hosted sandboxes, local sessions, and third-party harnesses.
AI SDK 7 introduces experimental harness abstractions and HarnessAgent : one API to run fully configured, established agent harnesses such as Claude Code, Codex, and Pi. Harnesses are configurable with a sandbox to operate in, custom instructions, skills, and tools. Run established harnesses through a consistent interface, configure each one independently, and swap one out without changing your integration

[truncated]
