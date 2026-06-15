---
source: "https://agentback.dev"
hn_url: "https://news.ycombinator.com/item?id=48546120"
title: "AgentBack: AI-native API/MCP framework for agents"
article_title: "AgentBack — interfaces agents use, backends agents build"
author: "ninemind"
captured_at: "2026-06-15T21:56:12Z"
capture_tool: "hn-digest"
hn_id: 48546120
score: 2
comments: 0
posted_at: "2026-06-15T19:47:13Z"
tags:
  - hacker-news
  - translated
---

# AgentBack: AI-native API/MCP framework for agents

- HN: [48546120](https://news.ycombinator.com/item?id=48546120)
- Source: [agentback.dev](https://agentback.dev)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T19:47:13Z

## Translation

タイトル: AgentBack: エージェント用の AI ネイティブ API/MCP フレームワーク
記事のタイトル: AgentBack — エージェントが使用するインターフェイス、エージェントが構築するバックエンド
説明: AgentBack は、REST、MCP、OpenAPI、型指定されたクライアント、および検証に 1 つの Zod コントラクト (すべての境界ごとに 1 つのスキーマ) を提供します。エージェントは、マシンアクション可能なエラー、/llms.txt、および MCP ツールを使用してこれを使用し、コーディング エージェントはそこからドリフトなしでサーフェス全体を作成できます。

記事本文:
エージェントバック .dev
ドキュメント
ガイド
ブログ
GitHub
エージェントネイティブ、エンドツーエンド
エージェントが使用するインターフェイス。
バックエンドエージェント
ビルドします。
AgentBack は、REST エンドポイント、MCP ツール、OpenAPI ドキュメントを提供します。
クライアントと検証 1 つの Zod コントラクト — 1 つ
スキーマ、あらゆる境界。エージェントはそのサーフェスを次のように消費します。
マシンアクション可能なエラー、 /llms.txt 、および MCP ツール。
そして、真実の情報源が 1 つしかないため、コーディング エージェントは次のことができます。
そこからドリフトなしで表面全体をオーサリングします。上に構築
ESM 用に再構築された、LoopBack 4 の実証済みの依存関係注入コア
そしてノード22。
アルファ · ESM のみ · ノード 22.13+ ·
タイプスクリプト 6
const City = z.object({city: z.string().min( 1 )});
const Report = z.object({summary: z.string(), tempC: z.number()});
@api ({basePath: '/weather' })
クラス WeatherController {
@get ( '/forecast/{city}' , {パス: 市 , 応答: レポート })
非同期予測 (入力: {path: z.infer< typeof City >}) { /* … */ }
}
@mcpServer ()
クラス ウェザーツール {
@tool ( 'forecast' , {入力: 都市 , 出力: レポート })
非同期予測 (入力: z.infer< typeof City >) { /* … */ }
}
1 つの成果物で、多くのビューが表示される
Zod スキーマは契約です
デコレータでスキーマを一度宣言します。フレームワークはあらゆるものを導き出します。
そこからの他の表現 — したがって、コーディングエージェントがサーフェスを追加します
1 回の編集で何も変化しません。
コードの実行前に解析されるランタイムバリデータリクエストとツール呼び出し
z.infer type コンパイル時に適用されるハンドラーの入力タイプ
OpenAPI 3.1.1 は /openapi.json で提供され、Swagger でレンダリングされます
UI
公式 MCP SDK の MCP ツール スキーマ入出力コントラクト
/llms.txt 同じルート レジストリからのエージェント読み取り可能なサーフェス マップ
型付きクライアント コード生成なし - クライアントは同じスキーマをインポートします
アプリとエージェントが使用する API の場合
ほとんどのスタックは、ランタイム コントラクト、サービス コントラクト、および
エージェント契約

3つのハンドシンクロの場所。
Binding の階層的なコンテキスト
@inject 、プロバイダー、インターセプター、拡張ポイント、
タグベースの検出 - LoopBack 4 から移植。
REST サーバーと MCP サーバーは、同じコンテナ上のコンポーネントです。走る
1 つのプロセス (コントローラーとツール クラス) のいずれか、または両方
単なるバインディングです。
安定したコード、フィールドごとの問題、違反したスキーマ、再試行可能性、
および修復のヒント — マシンが実行可能な同じエンベロープ
RESTとMCP。
確認: ペイロードにバインドされた確認トークンと
冪等性: ルート上で宣言されたキーの再生、または
このツールは 2 回実行してはなりません。
@price('$0.001') ルートまたはツールを計測します。価格
ゲートは、x402/MPP チャレンジを使用して未払い通話を拒否します。ストライプ
同じ使用ログからの従量課金。
TypeScript コンシューマーは同じスキーマをインポートし、型指定された呼び出しを取得します
さらに実行時検証 — クライアントは実行中のファイルをインポートしません。
サーバー。
すべてのツール定義にはコンテキスト ウィンドウ トークンがかかります。
接続。 toolCostReport MCP のトークン価格
エージェントが料金を支払う前に浮上します。
認証 (JWT、OAuth 2.1)、認可投票者、正常性プローブ、
Prometheus メトリクス、OpenTelemetry、レート制限、検証済みの構成
— すべて DI コンポーネントとして。
3 つのアイデアが全体のフレームワークを担っています。すべてがバインディングです。
コンテキスト、スキーマはデコレーター上に一度だけ存在し、サーバーは
コンポーネント。
フレームワーク境界からのデザインノート
エージェント向けサービス用の TypeScript フレームワーク: 1 つの Zod コントラクト
REST、MCP、ドキュメント、型指定されたクライアント、ポリシー、および使用レール全体にわたって。
境界の一貫性が製品です
AgentBack が Zod スキーマを同じアーティファクトとして扱う理由
ランタイム検証、TypeScript、OpenAPI、MCP、およびドキュメント。
1 つの @price デコレーターがルートまたは MCP ツールを測定し、
支払いレールの背後にゲートを設定します — 今日の x402、Stripe は同じものです
使用ログ。
API の最初の領域

ダーはエージェントです
/llms.txt および /llms-full.txt は次から提供されます
/openapi.json と同じレジストリ — 派生
契約が逸れることはありません。
コーディング エージェントにビルドを渡す
AgentBack は
教えるエージェントスキル
Claude Code、Codex、Cursor、および 20 以上の他のエージェントの規則
型シグネチャからは推測できない — デコレータ上のスキーマ、
スロット 0 入力バンドル、DI コンテナー。 Zod スキーマは 1 つだけです
信頼できる情報源であるため、エージェントはルート、MCP ツール、ドキュメントを追加します。
ドリフトのない単一の編集で。ドキュメントはエージェントも読むことができます —
/llms.txt 、
/llms-full.txt 、およびマークダウンミラー
すべてのページ — そのため、エージェントはフレームワークを、ユーザーのページを読み取るのと同じように読み取ります。
API。
# AgentBack スキルを選択したエージェントにインストールします
npx スキル追加 ninemindai/agentback
# エージェントにドキュメント コーパスを指示する
カール -s https://agentback.dev/llms.txt
AgentBack で構築
スニペットではなく、実際のサービスです。エージェントバックデモは
1 つの Zod スキーマ セットが標準入出力経由で提供される天気サーバー、
認証済み HTTP、および開発コンソール。それはコーディングエージェントの形です
そしてスキル生成: 1 つのコントラクトが REST、MCP、
クライアントとドキュメントを入力し、手動で同期を保つものは何もありません。
1 つの Zod スキーマ セット、stdio、認証済み HTTP、
そして開発コンソール。無料の Open-Meteo API によってサポートされています。
1 つのスキーマ → REST + MCP + 型付きクライアント + ドキュメント
AgentError 、 /llms.txt 、ツールコストレポート
正規のエージェントとスキルのビルド
サンプルを 2 分で実行します
AgentBack はアルファ版です - エンドツーエンドのサンプルの動作と API
表面はまだ動いています。リポジトリのクローンを作成し、動作する REST、MCP、
またはハイブリッド アプリで、それぞれエクスプローラー UI を備えています。
# クローンとビルド (テストとサンプルは dist/ に対して実行されます)
git clone https://github.com/ninemindai/agentback
pnpm インストール && pnpm ビルド
# REST + Swagger UI + コンテキスト エクスプローラー

pnpm -F hello-rest start
# 1 つのプロセス、両方の UI からの REST + MCP
pnpm -F hello-ハイブリッドの開始
AgentBack — エージェントとともに、エージェントのために構築されています。マサチューセッツ工科大学
ライセンス · © ninemind.ai および
LoopBack の貢献者。

## Original Extract

AgentBack gives REST, MCP, OpenAPI, typed clients, and validation one Zod contract — one schema, every boundary. Agents consume it with machine-actionable errors, /llms.txt, and MCP tools, and a coding agent can author the whole surface from it without drift.

AgentBack .dev
Docs
Guides
Blog
GitHub
Agent-native, end to end
Interfaces agents use .
Backends agents
build .
AgentBack gives your REST endpoints, MCP tools, OpenAPI docs, typed
clients, and validation one Zod contract — one
schema, every boundary. Agents consume that surface with
machine-actionable errors, /llms.txt , and MCP tools;
and because there's a single source of truth, a coding agent can
author the whole surface from it without drift. Built on
LoopBack 4's proven dependency-injection core, rebuilt for ESM
and Node 22.
alpha · ESM-only · Node 22.13+ ·
TypeScript 6
const City = z.object({city: z.string().min( 1 )});
const Report = z.object({summary: z.string(), tempC: z.number()});
@api ({basePath: '/weather' })
class WeatherController {
@get ( '/forecast/{city}' , {path: City , response: Report })
async forecast (input: {path: z.infer< typeof City >}) { /* … */ }
}
@mcpServer ()
class WeatherTools {
@tool ( 'forecast' , {input: City , output: Report })
async forecast (input: z.infer< typeof City >) { /* … */ }
}
One artifact, many views
The Zod schema is the contract
Declare a schema once, on the decorator. The framework derives every
other representation from it — so a coding agent adds a surface in
one edit and nothing drifts.
Runtime validator requests & tool calls parsed before your code runs
z.infer type the handler's input type, enforced at compile time
OpenAPI 3.1.1 served at /openapi.json , rendered in Swagger
UI
MCP tool schema input/output contracts on the official MCP SDK
/llms.txt agent-readable surface map from the same route registry
Typed client no codegen — the client imports the same schemas
For APIs that apps and agents consume
Most stacks keep the runtime contract, the service contract, and the
agent contract in three hand-synchronized places.
A hierarchical Context of Binding s with
@inject , providers, interceptors, extension points,
and tag-based discovery — ported from LoopBack 4.
REST and MCP servers are components over the same container. Run
either, or both from one process — controllers and tool classes
are just bindings.
Stable codes, per-field issues, the violated schema, retryability,
and remediation hints — the same machine-actionable envelope on
REST and MCP.
confirm: payload-bound confirmation tokens and
idempotency: key replay, declared on the route or
tool that must not run twice.
@price('$0.001') meters a route or tool; the price
gate refuses unpaid calls with an x402/MPP challenge. Stripe
metered billing from the same usage log.
TypeScript consumers import the same schemas and get typed calls
plus runtime validation — the client never imports a running
server.
Every tool definition costs context-window tokens on every
connection. toolCostReport token-prices the MCP
surface before agents pay for it.
Auth (JWT, OAuth 2.1), authorization voters, health probes,
Prometheus metrics, OpenTelemetry, rate limiting, validated config
— all as DI components.
Three ideas carry the whole framework: everything is a binding in a
context, schemas live once on the decorator, and servers are
components.
Design notes from the framework boundary
A TypeScript framework for agent-facing services: one Zod contract
across REST, MCP, docs, typed clients, policy, and usage rails.
Boundary coherence is the product
Why AgentBack treats a Zod schema as the same artifact across
runtime validation, TypeScript, OpenAPI, MCP, and docs.
One @price decorator meters a route or MCP tool and
gates it behind a payment rail — x402 today, Stripe from the same
usage log.
Your API's first reader is an agent
/llms.txt and /llms-full.txt served from
the same registry as /openapi.json — derived
contracts can't drift.
Hand the build to a coding agent
AgentBack ships an
agent skill that teaches
Claude Code, Codex, Cursor, and 20+ other agents the conventions
that aren't guessable from type signatures — schema-on-decorator,
the slot-0 input bundle, the DI container. One Zod schema is the only
source of truth, so an agent adds a route, its MCP tool, and its docs
in a single edit without drift. The docs are agent-readable too —
/llms.txt ,
/llms-full.txt , and a markdown mirror of
every page — so the agent reads the framework the way it reads your
API.
# install the AgentBack skill into your agent of choice
npx skills add ninemindai/agentback
# point an agent at the docs corpus
curl -s https://agentback.dev/llms.txt
Built with AgentBack
Not a snippet — a real service. agentback-demo is a
Weather server where one Zod schema set is served over stdio,
authenticated HTTP, and a dev console. It's the shape a coding agent
and the skill produce: one contract fanning out to REST, MCP, a
typed client, and docs, with nothing kept in sync by hand.
One Zod schema set, served three ways — stdio, authenticated HTTP,
and a dev console. Backed by the free Open-Meteo API.
one schema → REST + MCP + typed client + docs
AgentError , /llms.txt , tool-cost report
the canonical agent-plus-skill build
Run the examples in two minutes
AgentBack is in alpha — the end-to-end examples work and the API
surface is still moving. Clone the repo and run a working REST, MCP,
or hybrid app, each with its explorer UI.
# clone & build (tests and examples run against dist/)
git clone https://github.com/ninemindai/agentback
pnpm install && pnpm build
# REST + Swagger UI + Context Explorer
pnpm -F hello-rest start
# REST + MCP from one process, both UIs
pnpm -F hello-hybrid start
AgentBack — built with agents, for agents. MIT
licensed · © ninemind.ai and
LoopBack contributors.
