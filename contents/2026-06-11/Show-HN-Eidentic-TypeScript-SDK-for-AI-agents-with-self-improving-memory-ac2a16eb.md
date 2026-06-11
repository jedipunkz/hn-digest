---
source: "https://github.com/eidentic/eidentic"
hn_url: "https://news.ycombinator.com/item?id=48494704"
title: "Show HN: Eidentic – TypeScript SDK for AI agents with self-improving memory"
article_title: "GitHub - eidentic/eidentic: Eidentic is the open-source TypeScript SDK for AI agents with self-improving memory and production fundamentals built in. · GitHub"
author: "baranozdemir"
captured_at: "2026-06-11T19:05:09Z"
capture_tool: "hn-digest"
hn_id: 48494704
score: 3
comments: 0
posted_at: "2026-06-11T18:45:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Eidentic – TypeScript SDK for AI agents with self-improving memory

- HN: [48494704](https://news.ycombinator.com/item?id=48494704)
- Source: [github.com](https://github.com/eidentic/eidentic)
- Score: 3
- Comments: 0
- Posted: 2026-06-11T18:45:09Z

## Translation

タイトル: Show HN: Eidentic – 自己改善型メモリを備えた AI エージェント用の TypeScript SDK
記事タイトル: GitHub - eidentic/eidentic: Eidentic は、自己改善型メモリと生産基盤が組み込まれた AI エージェント用のオープンソース TypeScript SDK です。 · GitHub
説明: Eidentic は、自己改善メモリと生産基盤が組み込まれた AI エージェント用のオープンソース TypeScript SDK です。 - eidentic/eidentic

記事本文:
GitHub - eidentic/eidentic: Eidentic は、自己改善メモリと生産基盤が組み込まれた AI エージェント用のオープンソース TypeScript SDK です。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
同一の
/
同一の
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
16 コミット 16 コミット .changeset .changeset .github .github docs docs 例 例 パッケージ パッケージ スクリプト スクリプト .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md STABILITY.md STABILITY.md context7.json context7.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.jsonturbo.jsonturbo.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Eidentic は、自己改善型のメモリと、
生産基盤が組み込まれています。永続的な実行、コスト上限の強制、マルチテナント
分離、GDPR 消去、サンドボックス ツールは、あらかじめ組み込まれているわけではありません。 Apache-2.0、エンタープライズ ゲーティングなし。
Node、Bun、Deno、および Edge で実行されます。
ステータス: 0.x — API は v1 に向けて安定化しています。 「安定性」を参照してください。
私たちは過剰に販売するよりもギャップを過剰に明らかにしたいと考えています。
正直で再現可能な数値のベンチマーク。
import { Agent , AIModel , SqliteStore } from "eidentic" ;
"@ai-sdk/anthropic" から { anthropic } をインポートします。
const エージェント = 新しいエージェント ( {
id : "サポート" 、
モデル: 新しい AIModel ( anthropic ( "claude-sonnet-4-5" ) ) 、
ストア: new SqliteStore ( "./eidentic.sqlite" ) 、
} ) ;
for await ( const ev of エージェント . query ( "先週何を決めましたか?" , { sessionId : "u-42" } ) ) {
コンソール。ログ(ev);
}
なぜイーデンティックなのか？
ほとんどのエージェント フレームワークは、メモリ、コーディング/サンドボックス、DX、耐久性など、1 つのレーンをリードします。
オーケストレーションとかスキルとか。これらすべてが一緒に出荷されることはほとんどなく、実稼働の準備が整っています。
通常、エンタープライズ層の背後にあります。イーデンティックは、

sis: 1 つのコンポーザブルにすべてが含まれています。
フルオープンパッケージ。
1. 記憶力自体が向上します。ベクトルリコールだけでなく、4 層エンジン
自己編集メモリブロック、時間的知識グラフ（時間の経過とともに有効な事実。
矛盾は蓄積するのではなく無効化する）、睡眠時間の統合、および受動的
事実の抽出。 (設計仕様)
2. 生産の基本は組み込まれており、ボルトオンではありません。永続的なチェックポイント/再開
ツールのディスパッチは 1 回のみ、コスト上限 ($/トークン/ターン) とターンごとのコストの強制
可視性 、組み込みのレート制限 + クォータ 、OpenTelemetry GenAI スパン、デフォルトによる拒否
権限、サンドボックス化されたコード/コマンドの実行、モデルが決して見ることのないシークレット、ワンコール
すべてのストアに広がる消去権 (GDPR)。オフライン ワークロードの場合は、
バッチ ランナーとスケジュールされた/バックグラウンドの実行。そして、エージェントなしで出荷するため、
テストはブラインドで出荷され、CI パスレート ゲート プラスを備えた組み込みの評価ハーネスがあります。
ワンコールで実稼働トレースを回帰テストに昇格 - すべてのインシデントが
繰り返しではなくテストです。これらのいくつかは、オープン フレームワークの中で独自、またはほぼ独自のものです。
3. コンポーザブルで完全にオープンで、どこでも実行できます。ポートとアダプターのアーキテクチャ:
ストア (SQLite / libSQL / Postgres)、ベクター バックエンド (LanceDB / pgvector / Qdrant /
Pinecone)、またはエージェント コードに触れずに埋め込むことができます。 PDF、HTML、Markdown を取り込む
箱。 MCP (OAuth を使用) および A2A を介した相互運用。 Apache-2.0、コードゲーティングなし。検証日
CI の Node、Bun、Deno。
Eidentic は最初のライブラリです。別のサービスを実行する必要はなく、インポートするだけです
独自のバックエンドに直接アクセスし、agent.query() を呼び出します。スタンドアロン HTTP として実行する
service は、エージェントをサービスとして使用する場合のオプションの 2 番目のモードです。
1. 埋め込み — アプリにドロップします

(共通パス)
1 回のインストールでエージェントを構築し、任意のリクエスト ハンドラーからストリーミングします。エージェントが実行されます
サーバー側 (モデルキーを保持します)。フロントエンドはエンドポイントを呼び出すだけです。
npm install eidentic ai @ai-sdk/anthropic
Next.js (アプリルーター) — app/api/chat/route.ts :
Next.js / サーバーレス: SqliteStore ではなく、 @eidentic/libsql (純粋な JS、バンドラーフレンドリー) を使用します。
SqliteStore の背後にあるネイティブの better-sqlite3 アドオンは Next/Turbopack にバンドルされません
(動的要求はサポートされていません)。 npm install @eidentic/libsql し、ルートを
ノードのランタイム。 (ノードサーバー/スクリプトの場合、SqliteStore が最適です。上部のスニペットを参照してください。)
import { Agent , AIModel } from "eidentic" ;
"@eidentic/libsql" から {LibsqlStore } をインポートします。
"@ai-sdk/anthropic" から { anthropic } をインポートします。
エクスポート const runtime = "nodejs" ; // ネイティブ/エッジセーフ ストア;エッジランタイムではありません
const エージェント = 新しいエージェント ( {
id : "サポート" 、
モデル: 新しい AIModel ( anthropic ( "claude-sonnet-4-5" ) ) 、
ストア: new LibsqlStore ( "file:eidentic.db" ) 、
} ) ;
非同期関数のエクスポート POST ( req : リクエスト ) {
const {メッセージ , セッション ID } = await req 。 json() ;
const stream = new ReadableStream ( {
非同期開始 ( c ) {
for await (エージェントの const ev .query (message , { sessionId , signal : req . signal } ) )
ｃ． enqueue ( new TextEncoder ( ) . encode ( JSON . stringify ( ev ) + "\n" ) ) ;
ｃ．閉じる ( ) ;
} 、
} ) ;
return new Response (stream , { headers : { "content-type" : "application/x-ndjson" } } ) ;
}
エクスプレス:
アプリ。 post ( "/chat" , async ( req , res ) => {
レス。 type ( "application/x-ndjson" ) ;
const コントローラ = new AbortController() ;
レス。 on ( "close" , ( ) => { if ( ! res . writableEnded ) コントローラー . abort ( ) ; } ) ;
for await ( const ev of エージェント . query ( req . body . message , { sessionId : req . body

。 sessionId、シグナル:コントローラー。信号 } ) )
レス。 write ( JSON . stringify ( ev ) + "\n" ) ;
レス。終わり （ ） ;
} ) ;
Cloudflare Workers / Edge — 同じ Agent で、ストアを libSQL/Postgres アダプターに交換します。
デフォルトをエクスポート {
非同期フェッチ ( req : リクエスト ) {
const {メッセージ , セッション ID } = await req 。 json() ;
const stream = new ReadableStream ( {
非同期開始 ( c ) {
for await (エージェントの const ev .query (message , { sessionId , signal : req . signal } ) )
ｃ． enqueue ( new TextEncoder ( ) . encode ( JSON . stringify ( ev ) + "\n" ) ) ;
ｃ．閉じる ( ) ;
} 、
} ) ;
return new Response (stream , { headers : { "content-type" : "application/x-ndjson" } } ) ;
} 、
} ;
完全な実行可能なバージョン (プレーンな node:http 、追加のパッケージなし) が含まれています。
example/hello-embedded.ts — pnpm --filter eidentic-examples hello:embedded 。
2. サーバー — サービスとしてのエージェント (オプション)
エンドポイントを手書きしたくない場合、または専用のマルチテナント エージェント バックエンドが必要な場合
認証、セッション、ストリーミングを備えた @eidentic/server は、すぐに使える Hono アプリを提供します。
import { createServer ,serveNode , ApiKeyAuth } from "@eidentic/server" ;
const app = createServer ( {
エージェント: {サポート: エージェント}、
auth : ApiKeyAuth ( { key_live_123 : { userId : "u1" } } ) 、
} ) ;
awaitserveNode (app , {ポート : 3000 } ) ; // POST /v1/agents/support/query → SSE
または、プロジェクトをスキャフォールディングして開発環境で起動します。
npm create eidentic@latest my-agent
cd my-agent && eidentic dev # eidentic.config.ts をロードして提供します
箱の中身は何ですか
エリア
ハイライト
エージェント
ステートフル ReAct ループ · イベント ソース セッション · 構成可能な戦略 (リフレクション / 計画と実行) · トークン ストリーミング
記憶
語彙 + 意味想起 (RRF 融合) · 自己編集ブロック · 時間的知識グラフ · 睡眠時間の統合 · 受動的抽出

イオン・TTL/dedup
スキル
SKILL.md プロンプトスキル · テストゲートされた実行可能スキル (ed25519 署名付き) · オプションの自己進化
マルチエージェント
コンテキスト分離 + 共有バジェットを備えた spawn_agent 委任 · MCP ホストとサーバー · A2A プロトコル
実行
永続的なチェックポイント/再開 (1 回のみ) · 人間参加型の一時停止 · 協調的なキャンセル · コンテキストの圧縮
セキュリティと運用
デフォルト権限の拒否 · サンドボックス実行 (E2B) · シークレット分離 · コスト ガバナー · レート制限 + クォータ · OTel · GDPR 消去
店舗
SQLite · libSQL/Turso · Postgres · ベクター: LanceDB / pgvector / Qdrant / Pinecone · ローカル + ホスト型エンベッダー
DX
npx eidentic init scaffold · Studio dev ダッシュボード (npx eidentic Studio) · eval harness · メモリ ベンチマーク スイート
すべての機能には、実行可能なexamples/hello-*.tsが同梱されています（ほとんどの機能はモックモデルを使用しているため、APIキーはありません）
必要です）。完全なリストとそれぞれの実行方法については、機能ツアーを参照してください。
2 つの公開長期記憶ベンチマークでは、Eidentic の検索ベースの記憶力は、
フルコンテキストのベースライン — トークンの一部で。同じスクリプト、同じモデル、同じシード、完全な
分割、フルコンテキストベースラインが含まれます。正直な注意点と、メモリのカテゴリごとのギャップ
負けたものも一緒に公開されます。
履歴が大きくなるほど、より多くのメモリが優先されます。コンテキスト ウィンドウに最大 115,000 個のトークンが詰め込まれます。
証拠を邪魔者の中に隠しておきますが、的を絞った検索によって証拠が表面化します。方法論、
設定および再現コマンド: docs/BENCHMARKS.md 。
複製可能で実行可能なスターター アプリ - 各フレームワークのメモリを利用したチャット エージェント。 APIキーを追加する
そして npm run dev :
example-nextjs — Next.js アプリルーター + withEidentic ハンドラー + useChat
example-react — Eidentic サーバーに対する Vite + React フック ( useEidenticStream )
example-express — Agent をプレーンな Express にドロップします

SSE を介したルートとストリーム
pnpmインストール
pnpm -r ビルド
pnpm --filter eidentic-examples hello # モック モデル — API キーは必要ありません
実際のモデルに対して実行するか、トークンをライブでストリーミングします。
エクスポート ANTHROPIC_API_KEY=sk-ant-...
pnpm --filter eidentic-examples hello:real
pnpm --filter eidentic-examples hello:stream
デバッグ
DEBUG=eidentic:* を設定すると、詳細な名前空間付きループ ログ (モデル呼び出し、ツール ディスパッチ、メモリ) が得られます。
リコール、圧縮、コスト) とシークレット値が編集されます。スコープを指定します ( DEBUG=eidentic:tool,eidentic:cost )
集中するために。これは、何か問題があるように見えるときにエージェントが実際に何をしたかを確認する最も速い方法です。
DEBUG=eidentic: * pnpm --filter eidentic-examples こんにちは
パッケージレイアウト
eidentic アンブレラ パッケージには、コア、タイプ、モデル、SQLite、メモリが 1 つバンドルされています。
よくあるケースに合わせてインストールします。 32 個のパッケージすべてがこのモノリポジトリにあります。オプションのアダプターは、
個別にインストールするため、使用した分だけお支払いいただきます。これにより、コールドスタートのフットプリントが小さく抑えられます
不要なネイティブ アドオンを取り込むことを避けます。
npm install eidentic はコアスタックを提供します。オプションのパッケージ — サーバー、React フック、
ベクター ストア、サンドボックス - 設計により個別にインストールされます。上の表を参照してください。
いくつかのデフォルトはローカル開発では安全ですが、公開する前に注意が必要です。
Studio のデフォルトは NoAuth です。 serveStudio と createStud

[切り捨てられた]

## Original Extract

Eidentic is the open-source TypeScript SDK for AI agents with self-improving memory and production fundamentals built in. - eidentic/eidentic

GitHub - eidentic/eidentic: Eidentic is the open-source TypeScript SDK for AI agents with self-improving memory and production fundamentals built in. · GitHub
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
eidentic
/
eidentic
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits .changeset .changeset .github .github docs docs examples examples packages packages scripts scripts .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md STABILITY.md STABILITY.md context7.json context7.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json turbo.json turbo.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Eidentic is the open-source TypeScript SDK for AI agents with self-improving memory and
production fundamentals built in. Durable execution, enforced cost ceilings, multi-tenant
isolation, GDPR erasure, and sandboxed tools — not bolted on. Apache-2.0, no enterprise gating.
Runs on Node, Bun, Deno, and the edge .
Status: 0.x — APIs stabilizing toward v1; see STABILITY .
We'd rather over-disclose gaps than oversell — see the
benchmarks for honest, reproducible numbers.
import { Agent , AIModel , SqliteStore } from "eidentic" ;
import { anthropic } from "@ai-sdk/anthropic" ;
const agent = new Agent ( {
id : "support" ,
model : new AIModel ( anthropic ( "claude-sonnet-4-5" ) ) ,
store : new SqliteStore ( "./eidentic.sqlite" ) ,
} ) ;
for await ( const ev of agent . query ( "What did we decide last week?" , { sessionId : "u-42" } ) ) {
console . log ( ev ) ;
}
Why Eidentic?
Most agent frameworks lead one lane — memory, or coding/sandbox, or DX, or durable
orchestration, or skills. Rarely do all of these ship together, and production-readiness
is usually behind an enterprise tier. Eidentic's thesis: everything in one composable,
fully-open package.
1. Memory that improves itself. Not just vector recall — a four-tier engine with
self-editing memory blocks, a temporal knowledge graph (facts with validity over time;
contradictions invalidate rather than accumulate), sleep-time consolidation, and passive
fact extraction. ( design spec )
2. Production fundamentals, built in — not bolted on. Durable checkpoint/resume with
exactly-once tool dispatch, enforced cost ceilings ($/token/turn) with per-turn cost
visibility , built-in rate-limiting + quotas , OpenTelemetry GenAI spans, deny-by-default
permissions, sandboxed code/command execution, secrets the model never sees, and one-call
right-to-erasure (GDPR) that fans out across every store. For offline workloads there's a
batch runner and scheduled/background runs . And because shipping an agent without
tests is shipping blind, there's a built-in eval harness with a CI pass-rate gate plus
one-call promotion of a production trace into a regression test — every incident becomes
a test, not a repeat. Several of these are unique or near-unique among open frameworks.
3. Composable, fully open, runs everywhere. Ports-and-adapters architecture: swap the
store (SQLite / libSQL / Postgres), vector backend (LanceDB / pgvector / Qdrant /
Pinecone), or embedder without touching agent code. Ingest PDF, HTML, and Markdown out of
the box; interop via MCP (with OAuth) and A2A . Apache-2.0, no code-gating. Verified on
Node, Bun, and Deno in CI.
Eidentic is a library first . You don't have to run a separate service — you import it
straight into your own backend and call agent.query() . Running it as a standalone HTTP
service is an optional second mode for when you want agents-as-a-service.
1. Embedded — drop it into your app (the common path)
One install, then construct an agent and stream it from any request handler. The agent runs
server-side (it holds your model key); your frontend just calls your endpoint.
npm install eidentic ai @ai-sdk/anthropic
Next.js (App Router) — app/api/chat/route.ts :
Next.js / serverless: use @eidentic/libsql (pure-JS, bundler-friendly), not SqliteStore .
The native better-sqlite3 addon behind SqliteStore doesn't bundle under Next/Turbopack
( Dynamic require not supported ). npm install @eidentic/libsql and keep the route on the
Node runtime. (For Node servers/scripts, SqliteStore is great — see the snippet at the top.)
import { Agent , AIModel } from "eidentic" ;
import { LibsqlStore } from "@eidentic/libsql" ;
import { anthropic } from "@ai-sdk/anthropic" ;
export const runtime = "nodejs" ; // native/edge-safe store; not the edge runtime
const agent = new Agent ( {
id : "support" ,
model : new AIModel ( anthropic ( "claude-sonnet-4-5" ) ) ,
store : new LibsqlStore ( "file:eidentic.db" ) ,
} ) ;
export async function POST ( req : Request ) {
const { message , sessionId } = await req . json ( ) ;
const stream = new ReadableStream ( {
async start ( c ) {
for await ( const ev of agent . query ( message , { sessionId , signal : req . signal } ) )
c . enqueue ( new TextEncoder ( ) . encode ( JSON . stringify ( ev ) + "\n" ) ) ;
c . close ( ) ;
} ,
} ) ;
return new Response ( stream , { headers : { "content-type" : "application/x-ndjson" } } ) ;
}
Express:
app . post ( "/chat" , async ( req , res ) => {
res . type ( "application/x-ndjson" ) ;
const controller = new AbortController ( ) ;
res . on ( "close" , ( ) => { if ( ! res . writableEnded ) controller . abort ( ) ; } ) ;
for await ( const ev of agent . query ( req . body . message , { sessionId : req . body . sessionId , signal : controller . signal } ) )
res . write ( JSON . stringify ( ev ) + "\n" ) ;
res . end ( ) ;
} ) ;
Cloudflare Workers / edge — same Agent , swap the store for a libSQL/Postgres adapter:
export default {
async fetch ( req : Request ) {
const { message , sessionId } = await req . json ( ) ;
const stream = new ReadableStream ( {
async start ( c ) {
for await ( const ev of agent . query ( message , { sessionId , signal : req . signal } ) )
c . enqueue ( new TextEncoder ( ) . encode ( JSON . stringify ( ev ) + "\n" ) ) ;
c . close ( ) ;
} ,
} ) ;
return new Response ( stream , { headers : { "content-type" : "application/x-ndjson" } } ) ;
} ,
} ;
A complete, runnable version (plain node:http , no extra packages) is in
examples/hello-embedded.ts — pnpm --filter eidentic-examples hello:embedded .
2. Server — agents-as-a-service (optional)
When you'd rather not hand-write the endpoint, or want a dedicated multi-tenant agent backend
with auth, sessions, and streaming out of the box, @eidentic/server gives you a ready Hono app:
import { createServer , serveNode , ApiKeyAuth } from "@eidentic/server" ;
const app = createServer ( {
agents : { support : agent } ,
auth : ApiKeyAuth ( { key_live_123 : { userId : "u1" } } ) ,
} ) ;
await serveNode ( app , { port : 3000 } ) ; // POST /v1/agents/support/query → SSE
Or scaffold a project and boot it in dev:
npm create eidentic@latest my-agent
cd my-agent && eidentic dev # loads eidentic.config.ts and serves it
What's in the box
Area
Highlights
Agent
Stateful ReAct loop · event-sourced sessions · composable strategies (reflection / plan-and-execute) · token streaming
Memory
Lexical + semantic recall (RRF fusion) · self-editing blocks · temporal knowledge graph · sleep-time consolidation · passive extraction · TTL/dedup
Skills
SKILL.md prompt skills · test-gated executable skills (ed25519-signed) · optional self-evolution
Multi-agent
spawn_agent delegation with context isolation + shared budget · MCP host & server · A2A protocol
Execution
Durable checkpoint/resume (exactly-once) · human-in-the-loop suspension · cooperative cancellation · context compaction
Security & ops
Deny-by-default permissions · sandboxed exec (E2B) · secret isolation · cost governor · rate-limit + quotas · OTel · GDPR erasure
Stores
SQLite · libSQL/Turso · Postgres · vector: LanceDB / pgvector / Qdrant / Pinecone · local + hosted embedders
DX
npx eidentic init scaffold · Studio dev dashboard ( npx eidentic studio ) · eval harness · memory benchmark suite
Every feature ships a runnable examples/hello-*.ts (most use a mock model, so no API key
needed). See the feature tour for the full list and how to run each one.
On two public long-term-memory benchmarks, Eidentic's retrieval-based memory beats the
full-context baseline — at a fraction of the tokens. Same script, same models, same seed, full
splits, full-context baseline included. Honest caveats and the per-category gaps where memory
loses are published alongside.
The larger the history, the more memory wins: stuffing ~115k tokens into the context window
buries the evidence among distractors, while targeted retrieval surfaces it. Methodology,
configuration, and reproduction commands: docs/BENCHMARKS.md .
Clonable, runnable starter apps — a memory-backed chat agent in each framework. Add an API key
and npm run dev :
example-nextjs — Next.js App Router + withEidentic handler + useChat
example-react — Vite + React hooks ( useEidenticStream ) against an Eidentic server
example-express — drop the Agent into a plain Express route and stream over SSE
pnpm install
pnpm -r build
pnpm --filter eidentic-examples hello # mock model — no API key needed
Run against a real model or stream tokens live:
export ANTHROPIC_API_KEY=sk-ant-...
pnpm --filter eidentic-examples hello:real
pnpm --filter eidentic-examples hello:stream
Debugging
Set DEBUG=eidentic:* for verbose, namespaced loop logs (model calls, tool dispatch, memory
recall, compaction, cost) with secret values redacted. Scope it ( DEBUG=eidentic:tool,eidentic:cost )
to focus. It's the fastest way to see what an agent actually did when something looks off.
DEBUG=eidentic: * pnpm --filter eidentic-examples hello
Package layout
The eidentic umbrella package bundles core, types, model, sqlite, and memory — one
install for the common case. All 32 packages are in this monorepo; optional adapters are
separate installs so you only pay for what you use. This keeps cold-start footprint small
and avoids pulling in native addons you don't need.
npm install eidentic gives you the core stack. Optional packages — server, React hooks,
vector stores, sandbox — are separate installs by design. See the table above.
A few defaults are safe for local development but need attention before going public:
Studio defaults to NoAuth. serveStudio and createStud

[truncated]
