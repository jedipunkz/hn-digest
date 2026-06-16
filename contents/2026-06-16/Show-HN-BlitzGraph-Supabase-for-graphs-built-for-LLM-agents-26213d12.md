---
source: "https://blitzgraph.com"
hn_url: "https://news.ycombinator.com/item?id=48557002"
title: "Show HN: BlitzGraph – Supabase for graphs, built for LLM agents"
article_title: "BlitzGraph - The AI-native backend. | BlitzGraph"
author: "lveillard"
captured_at: "2026-06-16T16:37:30Z"
capture_tool: "hn-digest"
hn_id: 48557002
score: 2
comments: 0
posted_at: "2026-06-16T15:41:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: BlitzGraph – Supabase for graphs, built for LLM agents

- HN: [48557002](https://news.ycombinator.com/item?id=48557002)
- Source: [blitzgraph.com](https://blitzgraph.com)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T15:41:02Z

## Translation

タイトル: Show HN: BlitzGraph – LLM エージェント用に構築されたグラフ用の Supabase
記事のタイトル: BlitzGraph - AI ネイティブのバックエンド。 |ブリッツグラフ
説明: 現実をありのままにグラフでモデル化します。エージェントは、型指定された JSON クエリをプログラムで作成します。 SQL、結合、ORM はありません。
HN テキスト: こんにちは HN
SQL アレルギーになった後、私は完璧なgraphDB を探して Dgraph、Typedb、surrealdb で 120 以上の問題を開きました。これらはどれもエージェント向けに構築されたものではなく、現実を適切にモデル化するために SQL レガシーを完全に廃止するという私たちが達成したいことにも完璧に適合しませんでした。そこで、BlitzGraph を構築することにしました。 BlitzGraph では、レコード (ユニット) が複数のタイプ (種類) に属し、時間の経過とともに進化することができます。また、ポリモーフィックな関係も一流であり、複数の種類が同じ役割を果たすことができます。この設計は、古いテーブル パラダイムから抜け出し、他のテーブルの異なる ID でエンティティをそれ自体に接続する厄介な自己結合を行わずに、ライフサイクル全体を通してエンティティを追跡するのに役立ちます。 例: { "$id": "amazn", "$kinds": ["Company", "Prospect"], deal: ... } // 1 日目
{ "$id": "amazn", "$kinds": ["会社", "顧客"], 契約: .. } // 7日目
{ "$id": "amazn", "$kinds": ["Company", "Churned"], churnCause: "..." }, ... // 86 日目
BlitzGraph の違い: - GraphQL のようなネストされたクエリとミューテーション https://blitzgraph.com/docs - ポリモーフィックなレコードとリレーション
- 双方向 O(1) リレーション - ネイティブ カーディナリティ検証による参照整合性
- AI エージェントがプログラムで構築できるように設計された JSON クエリ/ミューテーション言語 - N+1 問題のないバッチ化されたクエリ/ミューテーション
- クイック ダッシュボードと MVP のための内蔵フロントエンド エンジン - ネイティブ全文検索、ファイル ストレージ、計算フィールド、一時的なサブスペース、ユニット履歴...
正直な比較: - vs typedb: 素晴らしいデータベースですが、アプリ開発には理想的ではありません

意見。一方、私たちは彼らの推論アイデアと、行ごとではなく突然変異がスマートに実行される方法を気に入って取り入れました - vs surrealdb: いくつかの核となる違い、重要な違いは、検証と変換をトポロジカルな順序で実行すること、そしてエッジが第一級市民であることです - vs dgraph: ポストコミットフックのような優れた機能がgraphQL層に接続されており、BGではそれが基礎的です - neo4j: 試したことがあるなら、ご存知でしょう - vs supabase/pg: BG はフラット クエリでは遅くなりますが、ネストされたクエリでは高速になります。しかし、BG を主に使用すると、テーブル パラダイムを取り除き、アプリを構築しながらグラフの世界に飛び込むことができます。 未準備: - blitzgraph はすでに AI エージェントにとって優れたメモリ バックエンドですが、セマンティック検索エンジンを完成させる必要があります。
- クエリプランナーが最適化されていない
- クラウド フロントエンドにはネイティブの認証エンジンがありませんが、まだベータ版が公開されています。問題を解決してください。
- 公共の遊び場: https://blitzgraph.com/#playground
- MCP: https://blitzgraph.com/mcp

記事本文:
BlitzGraph - AI ネイティブのバックエンド。 |ブリッツグラフ読み込み中...
バグ / フィードバック BlitzGraph ベータ版 · データが消去される可能性があります BlitzGraph ベータ版 · データは予告なく消去される可能性があります · リセットが予想される Playground 機能の比較 変更ログ ドキュメント 開始 → Y Combinator による裏付け · パブリック ベータ版 AI ネイティブのバックエンド。
アイデアは入力、API は出力。
現実をありのままにグラフでモデル化します。エージェントは、型指定された JSON クエリをプログラムで作成します。 SQL、結合、ORM はありません。
クエリ、フック、検証、計算フィールド。すべて含まれています。独自のバックエンドを無料で構築 →
BlitzGraph をリモート MCP サーバーとして追加し、同じライブ バックエンドから開始します。
claude mcp add --transport http blitzgraph https://blitzgraph.com/mcp Codex MCP サーバーを追加します。
codex mcp add blitzgraph --url https://blitzgraph.com/mcp 認証は自動的に実行されます。ブラウザで一度サインインします。その後、エージェントはツールを使用できるようになります。
複数の種類を持つエンティティ。双方向を行き来する関係。エージェントが正しく作成する、型指定された JSON クエリ言語。
ユーザーは、同時に管理者とモデレーターになることもできます。ロールテーブルや移行はありません。エンティティは、時間の経過とともに種類を獲得したり喪失したりすることで進化します。
「この記事を書いたのは誰ですか?」 「このユーザーは何を書いたのですか?」同じコスト、同じインデックス、双方向で O(1)。逆引き参照テーブルや追加のクエリはありません。
エージェントは SQL 文字列ではなくクエリ オブジェクトを作成します。フィルター、ネストされた展開、投影、および全文検索を 1 つのリクエストで実行します。ゼロ N+1。
電子メール、URL、日付、JSON、FLEX。 varcharだけではありません。データベースレベルでの組み込み検証。スキーマは、データが実際に何であるかを説明します。
カーディナリティ制約と onDelete ポリシー (カスケード、制限、リンク解除) はエンジン レベルで適用されます。グラフはデフォルトで一貫性を保ちます。
先行入力、プレフィックス、正確、および全ステム モードを備えたネイティブ BM25 エンジン。エラスティックサーチなし

、外部サービスはありません。グラフ走査内で動作します。
突然変異は、行ごとではなく、トポロジー的に並べ替えられ、最終結果で検証されます。ビジネス ルールはトランザクション全体の最終状態をチェックするため、複雑な複数エンティティの操作も正常に機能します。全か無か、常に一貫性があります。
データベース内のビジネス ロジック
検証、計算フィールド、変換、効果はすべてスキーマで定義されます。ビジネス ルールは、ミドルウェアやアプリのコードに分散するのではなく、データが存在する場所に存在します。
クエリ言語
エージェントは実際に正しく書いています。
SQL では、エージェントに文字列の生成と結果の形状の推測を強制します。 PostgREST は、テーブルの上に REST レイヤーを追加します。 BQL はクライアントからエンジンまでの構造化データです。
文字列ではなくコードからクエリを作成します。エージェントとアプリ コードは、SQL 文字列や ORM チェーンではなく、型指定された JSON オブジェクトを構築します。クエリはデータ構造です。生成ステップや解析の曖昧さはありません。
GraphQL のように、型指定された JSON として読み取ります。ルート エンティティのフェッチ、関係の展開、ネストされたフィールドのプロジェクトをすべて 1 つのリクエストで実行します。すべての深さで同じフィルター、並べ替え、形状を制限します。 GraphQL に似ていますが、スキーマ式はありません。
テーブルではなく現実をモデル化します。複数種類のエンティティ、双方向アーク、およびネイティブ $expand。データ モデルは、物事が実際にどのように関係しているかを反映しています。結合テーブルや外部キーの体操はありません。
関係をフィルタリングします。ローカルフィールドだけでなく、接続されたデータによってクエリを実行します。アプリケーション コード内の結合をステッチすることなく、セッション、メモリ、所有者を横断します。
// 人間とスペイン人の両方であるユニット。検索 + 計算フィールド。
{
"$kinds" : { "$all" : [ "人間" , "スペイン語" ] },
"$search" : "上級バックエンド Rust" ,
"$filter" : { "役割" : { "$in" : [ "エンジニア" , "デザイナー" ] } },
"$フィールド" : [
「名前」、「給料」、「ボーナス」、
{ "%total" : { "$js" : "給与 + ボーナス" } },
{
"$expand" : "プロジェクト" ,
"$sort

" : "-予算" ,
"$limit" : 3 、
"$fields" : [ "タイトル" 、 "予算" 、 "支出" 、
{ "%remaining" : { "$js" : "予算 - 使用済み" } }
】
}
】
} BQL · エージェント メモリ トラバーサル // 関連性の高いメモリ上のアーク メタデータを含む最近のセッション。
{
"$kinds" : "セッション" ,
"$フィルター" : {
"$created_at" : { "$gte" : "<日時>2026-04-01T00:00:00Z" }
}、
"$sort" : "-$created_at" ,
"$limit" : 5 、
"$フィールド" : [
"$created_at" ,
{
"$expandArc" : "思い出" ,
"$kinds" : [ "事実" , "アイデア" ],
"$search" : "ユーザー設定" ,
"$フィルター" : { "関連性" : { "$gte" : 0.7 } }、
"$sort" : "-関連性" ,
"$limit" : 10 、
"$fields" : [ "コンテンツ" , "関連性" ]
}
】
} // ぶつかった壁 素晴らしいデータベース。
まだテーブル用に作られています。
Supabaseはコラムで考えます。書類に凸。コレクション内のモンゴ アトラス。エンティティと関係性で考えると何が変わるのかを次に示します。
管理者になるユーザーにはロール テーブルが必要です。新しい役割？新しいテーブル、新しい JOIN、新しい移行。すべてのエンティティの変更はスキーマに影響します。
1 つのエンティティで複数の種類。ユーザーは[ユーザー、管理者、モデレータ]を同時に務めることができます。種類は自由に構成され、移行はありません。
「この記事を書いたのは誰ですか?」逆結合が必要です。外部キーは一方向に進みます。逆引き参照には、インデックス、JOIN、および慎重なクエリ計画が必要です。
双方向アーク、双方向 O(1)。すべての関係は両方向に保存されます。 「ユーザー→投稿」と「投稿→作成者」、同じコスト、同じインデックス。
ドキュメント間に実際の関係はありません。参照は、自分で管理する単なる文字列 ID です。トラバーサル、カーディナリティ、カスケード ポリシーはありません。
人間関係は第一級市民です。 onDelete による双方向アークのカスケード/制限/リンク解除。 $expand は、同じクエリ内のグラフを走査します。
フラットなドキュメントであり、エンティティの進化はありません。バグになるタスクには、タイプ フィールドと条件文があらゆる場所に必要です。スキーマ変更のリップル

アプリコードを通じて。
成長し構成する実体。種類を追加: [タスク、バグ]。エンティティは、既存のクエリを壊すことなく、バグ フィールドと関係を取得します。
1 つのエンティティがテーブル全体に分散されています。 SQL の Customer は、users、address、billing_info、および role に存在します。 1 つのエンティティを再構築するには、4 つのテーブル、4 つの ID、4 つの JOIN。
1 つのユニット、1 つの ID、永久に。ユニットはライフサイクル全体を通じて単一の ID を持ちます。そのすべてのデータと関係は、その 1 つの ID からアクセスできます。
エンティティの進化は移行地獄を意味します。見込み顧客になり、その後顧客になり、その後解約される企業ですか?これは、4 つのステータス フィールド、4 つの条件付きクエリ、そして何も壊れないことを祈ることです。
種類は自然に構成され、進化します。同じ単位: [会社] → [会社、見込客] → [会社、クライアント] → [会社、チャーン]。同じ ID、新機能、ゼロ移行。
文字列ベースのクエリ言語。暗号クエリは、アプリが生成し、エージェントが推測する文字列です。実行時に解析エラーが発生し、構造化された検証が行われません。
JSON 入力、JSON 出力。 BQL は構造化された JSON です。エージェントは文字列ではなくオブジェクトを作成します。 bql を使用してコンパイル時にスキーマ検証されます。マクロ。
ノードは一度に 1 つのラベルに属します。 Neo4j ラベルは、スキーマ、継承、フィールド分離のないフラットなタグです。ラベルを追加しても構造化フィールドは追加されません。
ユニットは複数の種類に属します。種類 [ユーザー、管理者、モデレーター] を持つユニットは、各種類のフィールド、検証、および関係を継承します。ポリモーフィズムがコアモデルです。
BlitzGraph と検討しているデータベースの比較。勝てる場所も含めて。正直になりたいと思います。
無料枠。 GitHub または電子メールでサインインします。型付き JSON API、グラフ トラバーサル、複数種類のエンティティ。

## Original Extract

Model reality as it is, in graphs. Your agents compose typed JSON queries programmatically. No SQL, no joins, no ORMs.

Hello HN
After becoming allergic to SQL, I opened 120+ issues in Dgraph, Typedb and surrealdb looking for the perfect graphDB. None of them was built for agents nor were they the perfect fit for what we wanted to achieve: fully ditching the SQL legacy to properly model reality. So we decided to build BlitzGraph In BlitzGraph, records (units) can belong to multiple types (kinds) and evolve through time. Also polymorphic relations are first class and multiple kinds can play the same role. This design helps to escape the old table paradigm and track entities throughout their lifecycle without awkward self-joins that connect an entity to itself under different IDs in other tables An example: { "$id": "amazn", "$kinds": ["Company", "Prospect"], deal: ... } // Day 1
{ "$id": "amazn", "$kinds": ["Company", "Customer"], contract: .. } // Day 7
{ "$id": "amazn", "$kinds": ["Company", "Churned"], churnCause: "..." }, ... // Day 86
What makes BlitzGraph different: - GraphQL-like nested queries and mutations https://blitzgraph.com/docs - Polymorphic records and relations
- Bidirectional O(1) relations - Referential integrity with native cardinality validations
- JSON query/mutation language designed so AI agents can build them programatically - Batched queries/mutations without N+1 issues
- Built-in frontend engine for quick dashboards and MVPs - Native full text search, file storage, computed fields, ephemeral subspaces, unit history...
Honest comparisons: - vs typedb: amazing db, but not ideal for app development. On the other hand we loved and brought their inference ideas and how mutations execute smartly instead of line per line - vs surrealdb: Several core differences, a key one is that we run validations and trasnformations in topological order, and our edges are first class citizens - vs dgraph: Their cool features like post commit hooks were attached to the graphQL layer, in BG it is fundational - neo4j: If you've tried it, you know - vs supabase/pg: BG is slower for flat queries but faster in nested ones. But with BG mainly you get rid of the tables paradigm and jump into the graph world while being able to build apps Not ready: - While blitzgraph is already an excellent memory backend for AI agents, we still need to finish the semantic search engine
- Query planner is not optimized
- Cloud frontends have no native auth engine yet Beta is live, please break things!
- Public playground: https://blitzgraph.com/#playground
- MCP: https://blitzgraph.com/mcp

BlitzGraph - The AI-native backend. | BlitzGraph Loading...
Bugs / feedback BlitzGraph beta · data may be wiped BlitzGraph beta · data may be wiped without notice · expect resets Playground Features Compare Changelog Docs Get Started → Backed by Y Combinator · Public beta The AI-native backend.
Idea in, API out.
Model reality as it is, in graphs. Your agents compose typed JSON queries programmatically. No SQL, no joins, no ORMs.
Queries, hooks, validations, computed fields. All included. Build your own backend for free →
Add BlitzGraph as a remote MCP server and start from the same live backend.
claude mcp add --transport http blitzgraph https://blitzgraph.com/mcp Codex Add the MCP server:
codex mcp add blitzgraph --url https://blitzgraph.com/mcp Auth runs automatically. You sign in once in your browser; after that your agent can use the tools.
Entities with multiple kinds. Relationships that traverse both ways. A typed JSON query language your agent composes correctly.
A User can also be an Admin and a Moderator, simultaneously. No role tables, no migrations. Entities evolve by gaining and losing kinds over time.
"Who wrote this post?" and "What did this user write?" Same cost, same index, O(1) both ways. No reverse-lookup tables, no extra queries.
Your agent composes query objects, not SQL strings. Filters, nested expands, projections, and full-text search in one request. Zero N+1.
EMAIL, URL, DATE, JSON, FLEX. Not just varchar. Built-in validation at the database level. Your schema describes what the data actually is.
Cardinality constraints and onDelete policies (cascade, restrict, unlink) enforced at the engine level. Your graph stays consistent by default.
Native BM25 engine with typeahead, prefix, exact, and all-stems modes. No Elasticsearch, no external service. Works inside graph traversals.
Mutations are topologically sorted and validated on the final result, not line by line. Business rules check the end state of the whole transaction, so complex multi-entity operations just work. All or nothing, always consistent.
Business logic in the database
Validations, computed fields, transforms, and effects, all defined in the schema. Your business rules live where the data lives, not scattered across middleware and app code.
The query language
agents actually write correctly.
SQL forces agents to generate strings and guess result shapes. PostgREST adds a REST layer on top of tables. BQL is structured data from client to engine.
Compose queries from code, not strings . Agents and app code build typed JSON objects, not SQL strings, not ORM chains. The query IS the data structure. No generation step, no parsing ambiguity.
GraphQL-like reads, as typed JSON . Fetch a root entity, expand relationships, project nested fields, all in one request. Same filter, sort, and limit shape at every depth. Like GraphQL, but without the schema ceremony.
Model reality, not tables . Multi-kind entities, bidirectional arcs, and native $expand. Your data model mirrors how things actually relate. No join tables, no foreign key gymnastics.
Filter through relationships . Query by connected data, not just local fields. Traverse sessions, memories, owners, without stitching joins in application code.
// Units that are both Human AND Spanish. Search + computed fields.
{
"$kinds" : { "$all" : [ "Human" , "Spanish" ] },
"$search" : "senior backend rust" ,
"$filter" : { "role" : { "$in" : [ "engineer" , "designer" ] } },
"$fields" : [
"name" , "salary" , "bonus" ,
{ "%total" : { "$js" : "salary + bonus" } },
{
"$expand" : "projects" ,
"$sort" : "-budget" ,
"$limit" : 3 ,
"$fields" : [ "title" , "budget" , "spent" ,
{ "%remaining" : { "$js" : "budget - spent" } }
]
}
]
} BQL · agent memory traversal // Recent sessions with arc metadata on high-relevance memories.
{
"$kinds" : "Session" ,
"$filter" : {
"$created_at" : { "$gte" : "<datetime>2026-04-01T00:00:00Z" }
},
"$sort" : "-$created_at" ,
"$limit" : 5 ,
"$fields" : [
"$created_at" ,
{
"$expandArc" : "memories" ,
"$kinds" : [ "Fact" , "Idea" ],
"$search" : "user preferences" ,
"$filter" : { "relevance" : { "$gte" : 0.7 } },
"$sort" : "-relevance" ,
"$limit" : 10 ,
"$fields" : [ "content" , "relevance" ]
}
]
} // the wall you hit Great databases.
Still built for tables.
Supabase thinks in columns. Convex in documents. Mongo Atlas in collections. Here's what changes when you think in entities and relationships.
A User who becomes Admin needs a role table. New role? New table, new JOIN, new migration. Every entity change touches the schema.
One entity, multiple kinds. A User can be [User, Admin, Moderator] simultaneously. Kinds compose freely, no migrations.
"Who wrote this post?" costs a reverse JOIN. Foreign keys go one direction. Reverse lookups need indexes, JOINs, and careful query planning.
Bidirectional arcs, O(1) both ways. Every relationship is stored in both directions. "User → Posts" and "Post → Author", same cost, same index.
No real relationships between documents. References are just string IDs you manage yourself. No traversals, no cardinality, no cascade policies.
Relationships are first-class citizens. Bidirectional arcs with onDelete cascade/restrict/unlink. $expand traverses the graph in the same query.
Flat documents, no entity evolution. A Task that becomes a Bug needs a type field and conditionals everywhere. Schema changes ripple through app code.
Entities that grow and compose. Add a kind: [Task, Bug]. The entity gains Bug fields and relationships without breaking existing queries.
One entity, scattered across tables. A Customer in SQL lives in users, addresses, billing_info, and roles. Four tables, four IDs, four JOINs to reconstruct one entity.
One unit, one ID, forever. A unit has a single ID throughout its entire lifecycle. All its data and relationships are accessible from that one identity.
Entity evolution means migration hell. A Company that becomes a Prospect, then a Client, then Churned? That's 4 status fields, 4 conditional queries, and a prayer that nothing breaks.
Kinds compose and evolve naturally. Same unit: [Company] → [Company, Prospect] → [Company, Client] → [Company, Churned]. Same ID, new capabilities, zero migrations.
String-based query language. Cypher queries are strings your app generates and your agent guesses at. Parsing errors at runtime, no structured validation.
JSON in, JSON out. BQL is structured JSON. Your agent composes objects, not strings. Schema-validated at compile time with the bql! macro.
Nodes belong to one label at a time. Neo4j labels are flat tags with no schema, no inheritance, no field isolation. Adding a label doesn't add structured fields.
Units belong to multiple kinds. A unit with kinds [User, Admin, Moderator] inherits each kind's fields, validations, and relationships. Polymorphism is the core model.
BlitzGraph vs the databases you’re considering. Including where they win. We’d rather be honest.
Free tier. Sign in with GitHub or email. Typed JSON API, graph traversals, multi-kind entities.
