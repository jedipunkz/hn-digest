---
source: "https://klr-pattern.github.io/nexusx/"
hn_url: "https://news.ycombinator.com/item?id=49310720"
title: "Model your business once – for humans and AI alike"
article_title: "Home - nexusx"
author: "tank-34"
captured_at: "2026-08-15T14:13:33Z"
capture_tool: "hn-digest"
hn_id: 49310720
score: 1
comments: 0
posted_at: "2026-08-15T14:08:40Z"
tags:
  - hacker-news
  - translated
---

# Model your business once – for humans and AI alike

- HN: [49310720](https://news.ycombinator.com/item?id=49310720)
- Source: [klr-pattern.github.io](https://klr-pattern.github.io/nexusx/)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T14:08:40Z

## Translation

タイトル: ビジネスを一度モデル化 – 人間も AI も同様
記事タイトル: ホーム - nexusx
説明: 深い AI 統合を備えた次世代ビジネス モデリング — GraphQL、REST、MCP、CLI、および 1 つの型付きモデルから派生した TS SDK。

記事本文:
ホーム - ネクサス
コンテンツにスキップ
ネクサス
ホーム
英語
ビジネスを一度モデル化してください —
人間にとってもAIにとっても。
ビジネス エンティティ、関係、ユース ケースを一度モデル化すると、GraphQL、REST、MCP、CLI、TS SDK はすべてそこから派生します。データはグラフです。ツールは単なる投影です。
GitHub
pip インストール nexusx
'},1500)}.bind(this,this))">
AI ネイティブの統合 – ボルトオンではない
同じタイプのビジネス モデルが、AI エージェントと開発者を第一級の消費者として提供します。
MCP はネイティブ プロトコルです。内部では厳密に型指定され、GraphQL が使用されます。
コンテキストの効率性 — エージェントは必要なフィールドを正確に選択します。 1 回の呼び出しで、要求されたものだけを含むネストされた N+1 プルーフ ツリーが返されます。
段階的な開示 — list_apps → description_compose_schema → description_compose_method → compose_query;スキーマはコンテキストに部分的に入り込むのですが、全体になることはありません。
SQLModel エンティティと型付き DTO を作成します。それが仕事のすべてです。
REST ルート、GraphQL スキーマ、CLI、および TS SDK — 定型文は不要です。
ビジネス ロジックを一度変更すると、すべてのプロトコルが同期して更新されます。
どのモデルでもアプリを出荷できます。メンテナンスを可能にできるものはほとんどありません。
LLM はコードを高速に生成しますが、構造的な制約がなければ、ロジックの重複、コンポーネントの相互混入、推測によるデバッグなど、負債が数週間後に表面化します。業界ではこれをバイブコーディングのコストと呼んでいます。
nexusx は、AI が宣言型モデルに書き込むもの (エンティティ、関係、型指定されたユースケース メソッド) を絞り込みます。構造は AI が正しく理解しなければならないものではありません。それはモデルによって保証されています。
AI は、散在するグルー コードではなく、モデルとユースケース メソッドを作成します。小さな差分。人間がレビュー可能。
ビジネスルールを一度変更します。すべてのプロトコルは同期して更新されます。納品ごとにメンテナンス費用がかかることはありません。
構造でわかる
型指定された契約と Voyager: エンティティ、

関係、ユースケース、およびそれらの依存関係が 1 つのライブ ER ビューとしてレンダリングされます。新人であっても、AI セッションを始めたばかりであっても、最初にコードを読まなくてもプロジェクト全体を把握できます。ボイジャー →
セマンティックレベルの同型性 — すべてのプロトコルは同じ型付きモデルから生成され、そのコピーにラップされるのではありません。
# "list sprints" — プロトコルごとに 1 回書かれます
@app.get("/sprints")
async defrest_list_sprints() -> リスト[SprintOut]:
... # クエリ + アセンブリ、再び
@strawberry.field
async defgraphql_sprints(self) -> list[SprintType]:
... # タイプ + ローダー、再び
@mcp.tool()
非同期定義 sprints_for_agents() -> str:
... # JSON ダンプ、再び
# ↑ ルールを変更 → すべてのコピーを修正
1 つの UseCaseService メソッド
クラス SprintService(UseCaseService):
"""スプリント計画業務。"""
@クエリ
async def list_sprints(cls) -> list[Sprintsummary]:
"""タスク、所有者、タスク数を含むスプリントをリストします。"""
return awaitload_sprints()
# 6納品1モデル ↓
🌐
REST + OpenAPI
型指定された FastAPI ルート。OpenAPI で表示されます。
エンティティは、接続されたデータを探索するための by_id / by_filter ルートになります。
ユースケース メソッドは、compose スキーマを介して型付きフィールドになります。
エージェントはそれを徐々に発見します。
サービスはコマンド グループになります。
構成スキーマから生成された型付きクライアント。
2 つの異なるジョブに 2 つの GraphQL サーフェス - いずれか、または両方を使用します。
データ グラフ — 探索とスライス
SQLModel エンティティとリレーションシップは、by_id / by_filter クエリ ルートになります。作成する関係リゾルバーが不要 — DataLoader のバッチ処理により、呼び出し元が通過する際に N+1 耐性が維持されます。
操作グラフ — 機能の呼び出し
型付きビジネス メソッドは、Web クライアント、統合、AI エージェントに安定した機能を公開し、1 つの定義から REST、MCP、CLI、SDK 上で提供されます。
エンティティは API コントラクトではありません。 DefineSubset は内部列を非表示にします

s、リレーションシップを自動ロードし、派生フィールドを計算します。
# エンドポイントごと: 手動 SQL、N+1、辞書書き換え
非同期デフォルト get_sprints():
sprints = await session.exec(select(Sprint))
結果 = []
スプリントの の場合:
タスク = await session.exec(
select(タスク).where(Task.sprint_id == s.id))
タスク内の t の場合:
t.owner = await session.get(User, t.owner_id)
# N+1 クエリ、脆弱な辞書構築
宣言型 DTO + 自動ロード
nexusx からインポート DefineSubset、ErManager、build_dto_select
クラスUserDTO(DefineSubset):
__subset__ = (ユーザー, ("id", "名前"))
クラス TaskDTO(DefineSubset):
__subset__ = (タスク, ("id", "title", "owner_id"))
所有者: UserDTO |なし = なし # 自動ロードされる
クラス SprintDTO(DefineSubset):
__subset__ = (スプリント, ("id", "名前"))
タスク: list[TaskDTO] = [] # 自動ロード
er = ErManager(entities=[ユーザー、スプリント、タスク]、session_factory=async_session)
リゾルバー = er.create_resolver()
async defload_sprints() -> リスト[SprintDTO]:
stmt = build_dto_select(SprintDTO) # ルート列のみ
async_session() をセッションとして使用する非同期:
rows = (await session.exec(stmt)).all()
dtos = [SprintDTO(**dict(r._mapping)) for r in rows]
return await Resolver().resolve(dtos) # ツリーが埋められ、バッチ処理される
# 関係ごとに 1 つのクエリ、ゼロ N+1
1 つのデータベースを超えて
同じ関係モデルが、より高度なアーキテクチャにも拡張されています。
DataLoader のバッチ処理、SQL 列のプルーニング、ウィンドウ関数のページネーション、および total_count は応答で要求された場合にのみ計算されます。
集計の場合は post_*、クロスレイヤー データ フローの場合は ExposeAs / SendTo。
非テーブル グラフ ルートとしての通常の Pydantic モデル — Redis、検索、および SDK ベースのデータが同じグラフに結合します。
中央ゲートウェイを使用せずに複数の nexusx データ グラフを構成します。つまり、nexusx サービスの同種のフェデレーションです。
CompusedErManager は 1 つのプロセスで複数のエンジンを構成します。 DTO フェデレーションの負荷

サービス全体にわたるパブリック DTO ツリー。
個別にパッケージ化されたアプリケーションとデータベースが 1 つの MCP サーバーに結合されます。
API のあらゆる決定を形作る原則。
1 つのフィールドの選択により、GraphQL 応答、ロードされた SQL 列、コピーされた DTO フィールド、MCP 出力、CLI --select、および total_count が計算されるかどうかが形成されます。
Redis、検索エンジン、他のデータベース、外部 API — 非同期バッチ関数との関係を宣言すると、同じローダー、DTO、GraphQL、ER ダイアグラム インフラストラクチャに参加します。
ビジネス メソッドはプロトコル オブジェクトに依存しません。ビルダーは型指定された署名を検査し、REST / MCP / CLI / SDK アダプターを接続します。 FromContext は、信頼できる値 (ユーザー、テナント) をクライアント引数として公開せずに挿入します。
API マニュアルをスキップ — エージェントを使用して構築する
4 フェーズ スキルをコーディング エージェント (Claude Code、Codex、Cursor など) にインストールし、アプリをわかりやすい言葉で説明します。エージェントがワークフローを推進します。モデルをレビューします。
'},1500)}.bind(this,this))">
🗺️
フェーズ 0 — ドメインのモデル化
コードを記述する前に、ドメイン モデルと永続化戦略を確認してください。
エンティティとリレーションシップ、GraphQL ヘルパー サーフェス、次に UseCase REST / MCP / CLI 配信。
必要に応じて、作成スキーマから型指定された TypeScript SDK を出力します。
既存のフレームワークやツールと連携します。
定型文ではなくエンティティから開始する
モデルを一度宣言すると、データ グラフ、応答 DTO、およびすべての配信が続きます。
nexusx・MITライセンス・タンキコード
· GitHub
ネクサス
オールマンデー/ネクサス
ガイド
ガイド
概要 — 1 つのモデル、2 つのグラフ
データグラフ
データグラフ
GraphQLモード
応答 DTO
応答 DTO
コア API モード
ORMを超えた関係
ORMを超えた関係
ER 図と非 ORM 関係
動作グラフ
動作グラフ
ユースケースサービス
AIの配信

AIの配信
AI 用の MCP を作成する
1 つのデータベースを超えて
1 つのデータベースを超えて
連邦
ツーリング
ツーリング
ボイジャーの視覚化
記事
記事
デザインのハイライト
APIリファレンス
APIリファレンス
グラフQLハンドラー
学習パス
ガイド (チュートリアル パス)

## Original Extract

Next-generation business modeling with deep AI integration — GraphQL, REST, MCP, CLI, and TS SDK derived from one typed model.

Home - nexusx
Skip to content
nexusx
Home
English
Model your business once —
for humans and AI alike.
Model your business entities, relationships, and use cases once — GraphQL, REST, MCP, CLI, and TS SDK all derive from it. Data is a graph; tools are just its projections.
GitHub
pip install nexusx
'},1500)}.bind(this,this))">
AI-native integration — not bolted on
The same typed business model serves AI agents and developers as first-class consumers.
MCP is a native protocol: strongly typed, GraphQL under the hood.
Context efficiency — agents select exactly the fields they need; one call returns a nested, N+1-proof tree with only what was asked.
Progressive disclosure — list_apps → describe_compose_schema → describe_compose_method → compose_query; the schema enters context piece by piece, never whole.
Write SQLModel entities and typed DTOs; that is the whole job.
REST routes, GraphQL schema, CLI, and TS SDK — zero boilerplate.
Change business logic once — every protocol updates in sync.
Any model can ship an app. Almost none can keep it maintainable.
LLMs generate code fast — but without structural constraints, the debt surfaces weeks later: duplicated logic, components bleeding into each other, debugging by guesswork. The industry calls it the cost of vibe coding.
nexusx narrows what AI writes to a declarative model — entities, relationships, and typed use-case methods. Structure is not something the AI has to get right; it is guaranteed by the model.
AI writes the model and use-case methods — not scattered glue code. Small diffs, reviewable by humans.
Change a business rule once; every protocol updates in sync. Maintenance cost does not multiply per delivery.
Understandable by construction
Typed contracts, plus Voyager: entities, relationships, use cases, and their dependencies rendered as one live ER view — grasp the whole project without reading code first, whether you are a new human or a fresh AI session. Voyager →
Semantic-level isomorphism — every protocol is generated from the same typed model, not wrapped around a copy of it.
# "list sprints" — written once per protocol
@app.get("/sprints")
async def rest_list_sprints() -> list[SprintOut]:
... # query + assembly, again
@strawberry.field
async def graphql_sprints(self) -> list[SprintType]:
... # types + loaders, again
@mcp.tool()
async def sprints_for_agents() -> str:
... # JSON dumping, again
# ↑ change the rule → fix every copy
One UseCaseService method
class SprintService(UseCaseService):
"""Sprint planning operations."""
@query
async def list_sprints(cls) -> list[SprintSummary]:
"""List sprints with tasks, owners, and task count."""
return await load_sprints()
# six deliveries, one model ↓
🌐
REST + OpenAPI
Typed FastAPI route, visible in OpenAPI.
Entities become by_id / by_filter roots for exploring connected data.
Use-case methods become typed fields via the compose schema.
Agents discover it progressively.
Services become command groups.
Typed client generated from the compose schema.
Two GraphQL surfaces for two different jobs — use either one, or both.
Data graph — explore and slice
SQLModel entities and relationships become by_id / by_filter query roots. No relationship resolvers to write — DataLoader batching keeps it N+1-proof as callers traverse.
Operation graph — invoke capabilities
Typed business methods expose stable capabilities to web clients, integrations, and AI agents — served over REST, MCP, CLI, and SDK from one definition.
Entities are not API contracts. DefineSubset hides internal columns, auto-loads relationships, and computes derived fields.
# Per-endpoint: manual SQL, N+1, dict munging
async def get_sprints():
sprints = await session.exec(select(Sprint))
result = []
for s in sprints:
tasks = await session.exec(
select(Task).where(Task.sprint_id == s.id))
for t in tasks:
t.owner = await session.get(User, t.owner_id)
# N+1 queries, fragile dict construction
Declarative DTO + auto-loading
from nexusx import DefineSubset, ErManager, build_dto_select
class UserDTO(DefineSubset):
__subset__ = (User, ("id", "name"))
class TaskDTO(DefineSubset):
__subset__ = (Task, ("id", "title", "owner_id"))
owner: UserDTO | None = None # auto-loaded
class SprintDTO(DefineSubset):
__subset__ = (Sprint, ("id", "name"))
tasks: list[TaskDTO] = [] # auto-loaded
er = ErManager(entities=[User, Sprint, Task], session_factory=async_session)
Resolver = er.create_resolver()
async def load_sprints() -> list[SprintDTO]:
stmt = build_dto_select(SprintDTO) # root columns only
async with async_session() as session:
rows = (await session.exec(stmt)).all()
dtos = [SprintDTO(**dict(r._mapping)) for r in rows]
return await Resolver().resolve(dtos) # tree filled, batched
# 1 query per relationship, zero N+1
Beyond one database
The same relationship model stretches into more advanced architectures.
DataLoader batching, SQL column pruning, window-function pagination — and total_count computed only when the response asks for it.
post_* for aggregations, ExposeAs / SendTo for cross-layer data flow.
Ordinary Pydantic models as non-table graph roots — Redis, search, and SDK-backed data join the same graph.
Compose multiple nexusx data graphs without a central gateway — homogeneous federation of nexusx services.
ComposedErManager composes multiple engines in one process; DTO federation loads public DTO trees across services.
Independently packaged applications and databases, combined into a single MCP server.
The principles that shape every API decision.
One field selection shapes the GraphQL response, the SQL columns loaded, the DTO fields copied, the MCP output, CLI --select, and whether total_count is even computed.
Redis, search engines, other databases, external APIs — declare a Relationship with an async batch function and they join the same loader, DTO, GraphQL, and ER-diagram infrastructure.
Business methods depend on no protocol object — builders inspect the typed signature and attach REST / MCP / CLI / SDK adapters. FromContext injects trusted values (user, tenant) without exposing them as client arguments.
Skip the API manual — build with an agent
Install the 4-phase skill into your coding agent — Claude Code, Codex, Cursor, and more — then describe your app in plain words. The agent drives the workflow; you review the model.
'},1500)}.bind(this,this))">
🗺️
Phase 0 — model the domain
Confirm the domain model and persistence strategy with you before any code is written.
Entities and relationships, GraphQL helper surface, then UseCase REST / MCP / CLI deliveries.
Optionally emit a typed TypeScript SDK from the compose schema.
Works with your existing frameworks and tools.
Start from entities, not boilerplate
Declare the model once — the data graph, response DTOs, and every delivery follow.
nexusx · MIT License · tangkikodo
· GitHub
nexusx
allmonday/nexusx
Guide
Guide
Overview — One Model, Two Graphs
Data Graph
Data Graph
GraphQL Mode
Response DTOs
Response DTOs
Core API Mode
Relationships beyond the ORM
Relationships beyond the ORM
ER Diagram & Non-ORM Relationships
Operation Graph
Operation Graph
UseCase Service
AI Delivery
AI Delivery
Compose MCP for AI
Beyond One Database
Beyond One Database
Federation
Tooling
Tooling
Voyager Visualization
Articles
Articles
Design Highlights
API Reference
API Reference
GraphQLHandler
Learning Path
Guide (Tutorial Path)
