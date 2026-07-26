---
source: "https://github.com/pipilot-dev/anyclaude-sdk"
hn_url: "https://news.ycombinator.com/item?id=49063069"
title: "Show HN: Anyclaude-SDK – Claude Code-Style SDK for OpenAI/Anthropic Endpoints"
article_title: "GitHub - pipilot-dev/anyclaude-sdk: Claude Code agent capabilities (tools, tool loop, MCP, sub-agents, sessions) for ANY OpenAI/Anthropic-compatible LLM — in the browser (WebContainer), Node, and Bun. No backend required. · GitHub"
author: "hansade"
captured_at: "2026-07-26T22:51:16Z"
capture_tool: "hn-digest"
hn_id: 49063069
score: 1
comments: 0
posted_at: "2026-07-26T22:27:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Anyclaude-SDK – Claude Code-Style SDK for OpenAI/Anthropic Endpoints

- HN: [49063069](https://news.ycombinator.com/item?id=49063069)
- Source: [github.com](https://github.com/pipilot-dev/anyclaude-sdk)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T22:27:22Z

## Translation

タイトル: Show HN: Anyclaude-SDK – OpenAI/Anthropic エンドポイント用のクロード コード スタイル SDK
記事のタイトル: GitHub - pipilot-dev/anyclaude-sdk: ブラウザ (WebContainer)、Node、および Bun における、あらゆる OpenAI/Anthropic 互換 LLM の Claude Code エージェント機能 (ツール、ツール ループ、MCP、サブエージェント、セッション)。バックエンドは必要ありません。 · GitHub
説明: ブラウザー (WebContainer)、Node、および Bun 内の OpenAI/Anthropic 互換 LLM 用の Claude Code エージェント機能 (ツール、ツール ループ、MCP、サブエージェント、セッション)。バックエンドは必要ありません。 - pipilot-dev/anyclaude-sdk

記事本文:
GitHub - pipilot-dev/anyclaude-sdk: ブラウザー (WebContainer)、Node、および Bun における、あらゆる OpenAI/Anthropic 互換 LLM の Claude Code エージェント機能 (ツール、ツール ループ、MCP、サブエージェント、セッション)。バックエンドは必要ありません。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します

n.
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ピパイロット開発
/
任意のクロード SDK
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
104 コミット 104 コミット .github .github anyclaude-react anyclaude-react create-anyclaude-app create-anyclaude-app docs-site docs-site サンプル サンプル プレイグラウンド プレイグラウンド スクリプト スクリプト src src .gitignore .gitignore .npmignore .npmignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMPATIBILITY.md COMPATIBILITY.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md TELEMETRY.md TELEMETRY.md compat.config.example.json compat.config.example.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude Code エージェントの機能 — ツール、ツール ループ、マルチターン会話、
MCP、サブエージェント、マルチエージェント チーム、セッション - あらゆる OpenAI または
Anthropic 互換の LLM エンドポイント、ブラウザーで実行
( WebContainer )、 Node 、および Bun 。バックエンドなし
必須、OAuth なし、ネイティブ バイナリなし。
ライブデモ: ブラウザで実行される完全な IDE ·
ドキュメント: anyclaude-docs.puter.site ·
React UI キット: anyclaude-react
同じ query() 非同期ジェネレーター インターフェイスと同じ SDKMessage を公開します。
エンベロープは @anthropic-ai/claude-agent-sdk なので、公式に対してコードが書かれています
SDK は出力を変更せずに反復できます。
マルチエージェント チームは 1 人のエージェントを超えています。コーディネーターは取締役会のタスクを次のエージェントに委任します。
ワーカー サブエージェントを並行して実行すると、実行中のワーカーにメッセージをディスパッチできます
そして次のツールラウンドに着地させます（プッシュド

メッセージキューのようなカラーリング)、
バックグラウンドディスパッチでライブで監視し、次に完了したワーカーをブロックします
wait_for_worker を使用 (イベント駆動型 — 結果が得られるたびに各結果を統合します。
ビジーポーリング)、さらに別の Web ワーカーまたはブラウザタブでエージェントを実行することもできます。
BroadcastChannelMailbox を介して 1 つのメールボックスを共有します。
「チームとサブエージェント」を参照してください。
npm install anyclaude-sdk @webcontainer/api
@webcontainer/api はオプションのピア依存関係です。使用する場合にのみ必要です。
WebContainerWorkspace 。独自の FileSystem/CommandExecutor を提供できます。
import { WebContainer } から '@webcontainer/api'
インポート {
クエリ、
WebContainerワークスペース 、
createOpenAIClient 、
ALL_CLAUDE_CODE_TOOLS 、
「anyclaude-sdk」から
// 1. WebContainer を起動し、ワークスペースとしてラップします。
const wc = WebContainer を待ちます。ブーツ ( )
const workspace = new WebContainerWorkspace ( wc )
// 2. OpenAI 互換のエンドポイントを指します。
const llm = createOpenAIClient ( {
apiKey:インポート。メタ。環境 。 VITE_OPENAI_API_KEY 、
BaseUrl : 'https://api.openai.com/v1' , // または Groq、Togetter、OpenRouter、local…
モデル: 'gpt-4o' 、
})
// 3. エージェントを実行します。公式 SDK と同じ形式です。
for await ( const msg of query ( { プロンプト : 'ファイルをリストしてプロジェクトを要約' , workspace , llm } ) ) {
if ( msg . type === 'アシスタント' ) {
for ( msg . message . content の const ブロック ) {
if ( block . type === 'text' ) console 。ログ (ブロック.テキスト)
}
} else if ( msg . type === 'result' && msg . subtype === 'success' ) {
コンソール。 log ( 'Done:' , msg . result )
}
}
MCP サーバー (外部 + インプロセス)
外部 MCP サーバーに接続するか、インプロセス ツールを定義します。ブラウザがブロックするから
クロスオリジン MCP フェッチ (CORS) を直接実行し、リモート サーバーに mcpProxy を渡します。
import { createSdkMcpServer , ツール } from 'anyclaude-sdk'
const calc = createSdkMcpSer

バージョン ( {
名前: '計算' 、
tools : [ ツール ( '加算' , '2 つの数値を加算' ,
{ type : 'object' 、プロパティ : { a : { type : 'number' } 、 b : { type : 'number' } } 、必須 : [ 'a' , 'b' ] } 、
( args ) => ( { content : [ { type : 'text' , text : String ( args . a + args . b ) } ] } ) ] ,
})
クエリ ( {
プロンプト、ワークスペース、llm、
mcpサーバー: {
calc , // インプロセス、ネットワークなし
docs : { type : 'http' , url : 'https://mcp.example.com' } , // リモート
} 、
// CORS プロキシ経由でリモート MCP をルーティングします (関数、`{url}`/`{rawUrl}` テンプレート、またはベア プレフィックス):
mcpProxy : 'https://my-proxy.example/?url={url}' ,
})
リモート ツールは mcp__<server>__<tool> として公開されます。
3 つのトランスポート クライアント。すべて同じ LLMClient インターフェイスを実装しています。
import { createOpenAIClient , createAnthropicClient , createResponsesClient } from 'anyclaude-sdk'
// OpenAI 互換のチャット補完 (OpenAI、Groq、Togetter、OpenRouter、xAI、Kilo、local…)
const a = createOpenAIClient ( { apiKey , BaseUrl : 'https://api.x.ai/v1' , model : 'grok-build-0.1' } )
// 人間的メッセージ API
const b = createAnthropicClient ( { apiKey , モデル : 'claude-sonnet-4-6' } )
// OpenAI レスポンス API (POST /v1/responses)
const c = createResponsesClient ( { apiKey , モデル : 'gpt-4o' } )
// AgentRouter ゲートウェイ (スポンサー) — createOpenAIClient に対する 1 行のプリセット。キー
// AGENTROUTER_API_KEY にフォールバックします。 1 つのベース URL を介して Claude/GPT/Gemini/DeepSeek/GLM/… に到達します。
「anyclaude-sdk」から { createAgentRouterClient } をインポートします
const d = createAgentRouterClient ( { モデル : 'claude-sonnet-4-5-20250929' } )
3 つすべてがツール呼び出し、ストリーミング、および使用法を同じ StreamResult に正規化します。
ツール呼び出しをインライン テキストとして出力するモデル用のフォールバック パーサーが含まれています。
すべてのクライアントは、extraHeaders / extraBody (プロバイダー固有のマージ) も受け入れます。
ヘッダー/PA

カスタム クライアントなしで RAM を実行し、次の場合に結果を再試行します。
一時的な失敗は再試行されました。
マルチターン/インタラクティブセッション
PromptStream を使用して、時間の経過とともにユーザーのターンを促進します。
import { query , PromptStream } from 'anyclaude-sdk'
const プロンプト = new PromptStream ( )
const session = query ({ プロンプト : プロンプト , ワークスペース , llm , モデル : 'gpt-4o' } )
プロンプトが表示されます。 Push ( '挨拶文を含む hello.txt を作成' )
// …後で、UI 入力に基づいて:
プロンプトが表示されます。 Push (「フランス語に翻訳しましょう」)
プロンプトが表示されます。 end ( ) // 会話を閉じます
for await (セッションのconst msg) {
// メッセージをレンダリングします…
}
ツール
ALL_CLAUDE_CODE_TOOLS には次のものが含まれます。
ファイル読み取り: 画像、PDF、ノートブック
read_file はファイルタイプごとにディスパッチします。画像と PDF バイトは、
フォローアップ ユーザー ターンとして自動的にモデル化します (Anthropic がネイティブになります)
画像/ドキュメントブロック。 OpenAI 互換エンドポイントは image_url / file を取得します
パーツ))、モデルはテキストの概要だけでなく、実際にファイルを確認できるようになります。
制限を使用して上限を調整します。
query ({ プロンプト , ワークスペース , llm , 制限 : { maxTokens : 25000 , maxImageBytes : 3_750_000 , maxPdfPages : 20 } } )
tools: : を介して、サブセットまたは独自の Tool[] を渡します。
import { readFile , writeFile , editFile } from 'anyclaude-sdk'
query ({ プロンプト 、 ワークスペース 、 llm 、 ツール : [ readFile 、 writeFile 、 editFile ] } )
スラッシュコマンド
/ で始まるユーザーターンはインターセプトされます。組み込み: /help 、 /clear 、
/compact [フォーカス] (自由なコンテキストに履歴を要約)、 /tools 、 /cost 、
/モデル .独自のプロンプト テンプレート コマンドを定義します。
import { クエリ , プロンプトコマンド } から 'anyclaude-sdk'
クエリ ( {
プロンプト:プロンプトストリーム、ワークスペース、llm、
コマンド: [promptCommand ( 'review' , 'diff を確認してください' , 'このコードを確認して問題点をリストしてください: $ARGUMENTS' ) ] ,
})
// ユーザータイプ: /review src/app.ts
バックグラウンドタスク
wを有効にする

i 番目の背景: サブエージェントを実行するか、クリティカルな作業を長時間行う場合は true
パス。タスク ツールは run_in_background を取得します (タスク ID をすぐに返します)。
task_list / task_output / task_stop ツールを使用すると、エージェントがそれらをポーリングできます。
Comlink ワーカー ハーネスを介したオフメインスレッド実行 (オプション)
(エクスポーズバックグラウンドワーカー / ラップワーカー);スレッド内マネージャーはそれなしでも動作します。
クエリ ({ プロンプト 、 ワークスペース 、 llm 、 エージェント : { } 、 バックグラウンド : true } )
個別の Web ワーカーのエージェント
2 つの部分: メイン→ワーカー コントロールの Comlink ( WrapWorker / ExposeBackgroundWorker 、
上記)、および BroadcastChannelMailbox を使用して、別のワーカーのエージェントが噂話をする
ポストスタイル。ドロップイン メールボックスなので、既存のチーム ツール
( send_message /dispatch_tasks ) はワーカー間で変更されずに機能します。
「anyclaude-sdk」から { BroadcastChannelMailbox } をインポートします
// 各 Web Worker/tab/worker_thread 内で、同じチャネル名:
const mailbox = new BroadcastChannelMailbox ( { channelName : 'team' , Origin : 'planner' } )
クエリ ({ プロンプト 、 ワークスペース 、 llm 、 チーム : true 、 メールボックス } )
// あるワーカーによって送信されたメッセージは、別のワーカーのアドレス指定されたエージェントの受信箱に届きます。
デフォルトではグローバル BroadcastChannel を使用します。耐久性のあるクロス集計配信用
(IndexedDB/localStorage フォールバック、古いブラウザ、ノード) ワンコール ヘルパーを使用する
— バンドルされたブロードキャスト チャネルによってサポートされています
パッケージ、遅延インポートされるため、それを使用しないバンドルから除外されます。
const mailbox = BroadcastChannelMailbox を待ちます。クロスタブ ( { チャネル名 : 'チーム' 、オリジン : 'プランナー' } )
クエリ ({ プロンプト 、 ワークスペース 、 llm 、 チーム : true 、 メールボックス } )
実行中のエージェントに配信をプッシュします。エージェント宛てのメッセージは、
次のターン境界でトランスクリプトに自動挿入されます - と同じモデル
メッセージ キューからのものですが、共有メールボックスからのものです。したがって、コーディネーター（またはピア、または別の）
ワー

ker) は、タスクの途中で実行中のサブエージェントをリダイレクトでき、サブエージェントは
サブエージェントの次のツールラウンドでは、ポーリングツールは必要ありません。 dispatch_tasks にはそれぞれの名前が付けられます
ワーカー worker:<taskId> なので、特定のものをターゲットにできます。
メールボックス。 send ( 'coordinator' , 'worker:task_1' , '作業中: ログも追加します' )
//worker:task_1 は次のステップで「[チーム メッセージ] - コーディネーターから: ...」を参照します。
デフォルトではチームでオンになります: true ; query({deliveryTeamMessages: false }) を介してオプトアウトします。
WebContainer に縛られることはありません。サンドボックスは、単なるファイルシステムと
CommandExecutor と組み合わせることができます。
アダプターは、各プロバイダーのクライアントを構造的にラップします (アダプターに厳密な依存関係はありません)。
SDK — 使用する SDK のみをインストールします):
import { E2BSandbox 、 VercelSandbox 、 DaytonaSandbox 、 CloudflareSandbox } から 'anyclaude-sdk'
// 例: E2B
「e2b」から { サンドボックス } をインポートします
const sbx = サンドボックスを待ちます。作成 ( )
const ワークスペース = 新しい E2BSandbox ( sbx )
クエリ ({ プロンプト 、 ワークスペース 、 llm } )
サポートされている: WebContainer 、 E2B 、 Vercel Sandbox 、 Daytona 、
Cloudflare サンドボックス 、および LocalSandbox (実際の OS)。すべて同じを実装します
サンドボックスインターフェイス。
ホスト マシンのファイル システムとシェルに対してエージェントを直接実行します。
Claude Code — 自動プラットフォーム検出機能付き (Windows / macOS / Linux)

[切り捨てられた]

## Original Extract

Claude Code agent capabilities (tools, tool loop, MCP, sub-agents, sessions) for ANY OpenAI/Anthropic-compatible LLM — in the browser (WebContainer), Node, and Bun. No backend required. - pipilot-dev/anyclaude-sdk

GitHub - pipilot-dev/anyclaude-sdk: Claude Code agent capabilities (tools, tool loop, MCP, sub-agents, sessions) for ANY OpenAI/Anthropic-compatible LLM — in the browser (WebContainer), Node, and Bun. No backend required. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
pipilot-dev
/
anyclaude-sdk
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
104 Commits 104 Commits .github .github anyclaude-react anyclaude-react create-anyclaude-app create-anyclaude-app docs-site docs-site examples examples playground playground scripts scripts src src .gitignore .gitignore .npmignore .npmignore CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMPATIBILITY.md COMPATIBILITY.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md TELEMETRY.md TELEMETRY.md compat.config.example.json compat.config.example.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Claude Code agent capabilities — tools, the tool loop, multi-turn conversations,
MCP, sub-agents, multi-agent teams , sessions — against any OpenAI- or
Anthropic-compatible LLM endpoint , running in the browser
( WebContainer ), Node , and Bun . No backend
required, no OAuth, no native binaries.
Live demo: a full IDE running in your browser ·
Docs: anyclaude-docs.puter.site ·
React UI kit: anyclaude-react
It exposes the same query() async-generator interface and the same SDKMessage
envelope as @anthropic-ai/claude-agent-sdk , so code written against the official
SDK can iterate our output unchanged.
Multi-agent teams go beyond one agent: a coordinator delegates board tasks to
worker sub-agents in parallel, you can dispatch a message to a running worker
and have it land on its next tool round (push delivery, like the message queue),
supervise them live with background dispatch, block for the next finished worker
with wait_for_worker (event-driven — integrate each result as it lands, no
busy-polling), and even run agents in separate Web Workers or browser tabs
that share one mailbox via BroadcastChannelMailbox .
See Teams & sub-agents .
npm install anyclaude-sdk @webcontainer/api
@webcontainer/api is an optional peer dependency — only needed if you use
WebContainerWorkspace . You can supply your own FileSystem / CommandExecutor .
import { WebContainer } from '@webcontainer/api'
import {
query ,
WebContainerWorkspace ,
createOpenAIClient ,
ALL_CLAUDE_CODE_TOOLS ,
} from 'anyclaude-sdk'
// 1. Boot a WebContainer and wrap it as a workspace.
const wc = await WebContainer . boot ( )
const workspace = new WebContainerWorkspace ( wc )
// 2. Point at any OpenAI-compatible endpoint.
const llm = createOpenAIClient ( {
apiKey : import . meta . env . VITE_OPENAI_API_KEY ,
baseUrl : 'https://api.openai.com/v1' , // or Groq, Together, OpenRouter, local…
model : 'gpt-4o' ,
} )
// 3. Run the agent — same shape as the official SDK.
for await ( const msg of query ( { prompt : 'List the files and summarize the project' , workspace , llm } ) ) {
if ( msg . type === 'assistant' ) {
for ( const block of msg . message . content ) {
if ( block . type === 'text' ) console . log ( block . text )
}
} else if ( msg . type === 'result' && msg . subtype === 'success' ) {
console . log ( 'Done:' , msg . result )
}
}
MCP servers (external + in-process)
Connect external MCP servers or define in-process tools. Because browsers block
direct cross-origin MCP fetches (CORS), pass a mcpProxy for remote servers:
import { createSdkMcpServer , tool } from 'anyclaude-sdk'
const calc = createSdkMcpServer ( {
name : 'calc' ,
tools : [ tool ( 'add' , 'Add two numbers' ,
{ type : 'object' , properties : { a : { type : 'number' } , b : { type : 'number' } } , required : [ 'a' , 'b' ] } ,
( args ) => ( { content : [ { type : 'text' , text : String ( args . a + args . b ) } ] } ) ) ] ,
} )
query ( {
prompt , workspace , llm ,
mcpServers : {
calc , // in-process, no network
docs : { type : 'http' , url : 'https://mcp.example.com' } , // remote
} ,
// Route remote MCP through a CORS proxy (function, `{url}`/`{rawUrl}` template, or bare prefix):
mcpProxy : 'https://my-proxy.example/?url={url}' ,
} )
Remote tools are exposed as mcp__<server>__<tool> .
Three transport clients, all implementing the same LLMClient interface:
import { createOpenAIClient , createAnthropicClient , createResponsesClient } from 'anyclaude-sdk'
// OpenAI-compatible Chat Completions (OpenAI, Groq, Together, OpenRouter, xAI, Kilo, local…)
const a = createOpenAIClient ( { apiKey , baseUrl : 'https://api.x.ai/v1' , model : 'grok-build-0.1' } )
// Anthropic Messages API
const b = createAnthropicClient ( { apiKey , model : 'claude-sonnet-4-6' } )
// OpenAI Responses API (POST /v1/responses)
const c = createResponsesClient ( { apiKey , model : 'gpt-4o' } )
// AgentRouter gateway (sponsor) — one-line preset over createOpenAIClient; key
// falls back to AGENTROUTER_API_KEY. Reaches Claude/GPT/Gemini/DeepSeek/GLM/… via one base URL.
import { createAgentRouterClient } from 'anyclaude-sdk'
const d = createAgentRouterClient ( { model : 'claude-sonnet-4-5-20250929' } )
All three normalize tool calls, streaming, and usage to the same StreamResult ,
and include a fallback parser for models that emit tool calls as inline text.
Every client also accepts extraHeaders / extraBody (merge provider-specific
headers/params without a custom client) and surfaces retries on the result when
transient failures were retried.
Multi-turn / interactive sessions
Use a PromptStream to push user turns over time:
import { query , PromptStream } from 'anyclaude-sdk'
const prompts = new PromptStream ( )
const session = query ( { prompt : prompts , workspace , llm , model : 'gpt-4o' } )
prompts . push ( 'Create a hello.txt with a greeting' )
// …later, based on UI input:
prompts . push ( 'Now translate it to French' )
prompts . end ( ) // close the conversation
for await ( const msg of session ) {
// render msg…
}
Tools
ALL_CLAUDE_CODE_TOOLS includes:
File reading: images, PDFs, notebooks
read_file dispatches by file type. Image and PDF bytes are forwarded to the
model automatically as a follow-up user turn (Anthropic gets native
image / document blocks; OpenAI-compatible endpoints get image_url / file
parts), so the model can actually see the file, not just a text summary.
Tune the caps via limits :
query ( { prompt , workspace , llm , limits : { maxTokens : 25000 , maxImageBytes : 3_750_000 , maxPdfPages : 20 } } )
Pass a subset, or your own Tool[] , via tools: :
import { readFile , writeFile , editFile } from 'anyclaude-sdk'
query ( { prompt , workspace , llm , tools : [ readFile , writeFile , editFile ] } )
Slash commands
A user turn beginning with / is intercepted. Built-ins: /help , /clear ,
/compact [focus] (summarizes history to free context), /tools , /cost ,
/model . Define your own prompt-template commands:
import { query , promptCommand } from 'anyclaude-sdk'
query ( {
prompt : promptStream , workspace , llm ,
commands : [ promptCommand ( 'review' , 'Review the diff' , 'Review this code and list issues: $ARGUMENTS' ) ] ,
} )
// user types: /review src/app.ts
Background tasks
Enable with background: true to run sub-agents or long work off the critical
path. The task tool gains run_in_background (returns a task id immediately),
and task_list / task_output / task_stop tools let the agent poll them.
Optional off-main-thread execution via a Comlink worker harness
( exposeBackgroundWorker / wrapWorker ); the in-thread manager works without it.
query ( { prompt , workspace , llm , agents : { } , background : true } )
Agents in separate Web Workers
Two halves: Comlink for main→worker control ( wrapWorker / exposeBackgroundWorker ,
above), and BroadcastChannelMailbox so agents in different workers gossip
mailbox-style. It's a drop-in Mailbox , so the existing team tools
( send_message / dispatch_tasks ) work unchanged across workers:
import { BroadcastChannelMailbox } from 'anyclaude-sdk'
// inside each Web Worker / tab / worker_thread, same channel name:
const mailbox = new BroadcastChannelMailbox ( { channelName : 'team' , origin : 'planner' } )
query ( { prompt , workspace , llm , team : true , mailbox } )
// messages sent by one worker land in the addressed agent's inbox in another.
Uses the global BroadcastChannel by default. For durable cross-tab delivery
(IndexedDB/localStorage fallbacks, older browsers, Node) use the one-call helper
— it's backed by the bundled broadcast-channel
package, lazy-imported so it stays out of bundles that don't use it:
const mailbox = await BroadcastChannelMailbox . crossTab ( { channelName : 'team' , origin : 'planner' } )
query ( { prompt , workspace , llm , team : true , mailbox } )
Push delivery to a running agent. Messages addressed to an agent are
auto-injected into its transcript at the next turn boundary — same model as the
message queue, but from the shared mailbox. So a coordinator (or peer, or another
worker) can redirect a running sub-agent mid-task and it lands on the
sub-agent's next tool round, no polling tool needed. dispatch_tasks names each
worker worker:<taskId> so you can target a specific one:
mailbox . send ( 'coordinator' , 'worker:task_1' , 'while you work: also add logging' )
// worker:task_1 sees "[Team messages] - from coordinator: ..." on its next step.
On by default with team: true ; opt out via query({ deliverTeamMessages: false }) .
You aren't tied to WebContainer. A Sandbox is just a FileSystem plus a
CommandExecutor , and you can mix and match.
Adapters wrap each provider's client structurally (no hard dependency on their
SDKs — install only the one you use):
import { E2BSandbox , VercelSandbox , DaytonaSandbox , CloudflareSandbox } from 'anyclaude-sdk'
// e.g. E2B
import { Sandbox } from 'e2b'
const sbx = await Sandbox . create ( )
const workspace = new E2BSandbox ( sbx )
query ( { prompt , workspace , llm } )
Supported: WebContainer , E2B , Vercel Sandbox , Daytona ,
Cloudflare Sandbox , and LocalSandbox (real OS). All implement the same
Sandbox interface.
Run the agent directly against the host machine's filesystem and shell — like
Claude Code — with automatic platform detection (Windows / macOS / Linux

[truncated]
