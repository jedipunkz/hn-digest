---
source: "https://predictabledialogs.com/learn/ai-stack/choosing-database-configurable-ai-agents"
hn_url: "https://news.ycombinator.com/item?id=49308031"
title: "Choosing the Right Database for AI Agents: LLM Generated SQL"
article_title: "Choosing the Right Database for AI Agents: LLM Generated SQL | Learn - Predictable Blog"
author: "jaikant"
captured_at: "2026-08-15T06:20:46Z"
capture_tool: "hn-digest"
hn_id: 49308031
score: 2
comments: 0
posted_at: "2026-08-15T05:53:34Z"
tags:
  - hacker-news
  - translated
---

# Choosing the Right Database for AI Agents: LLM Generated SQL

- HN: [49308031](https://news.ycombinator.com/item?id=49308031)
- Source: [predictabledialogs.com](https://predictabledialogs.com/learn/ai-stack/choosing-database-configurable-ai-agents)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T05:53:34Z

## Translation

タイトル: AI エージェントに適切なデータベースの選択: LLM で生成された SQL
記事のタイトル: AI エージェントに適切なデータベースの選択: LLM で生成された SQL |学ぶ - 予測可能なブログ
説明: ユーザー定義の構造化データを取り込む構成可能な AI エージェントについて、Postgres、JSONB、DuckDB、SQLite、MongoDB を比較します。

記事本文:
AI エージェントに適切なデータベースの選択: LLM 生成 SQL |学ぶ - 予測可能なブログ サインイン エージェントの作成 2026 年 8 月 15 日土曜日公開 AI エージェントに適切なデータベースの選択: LLM 生成 SQL
AI エージェントは、単純なチャットボットから、さまざまなドメインやデータ ソースに適応できる構成可能なシステムへと急速に進化しています。特に興味深い使用例の 1 つは、ユーザーが独自の構造化データを使用して構成できるエージェントです。
エージェントは、あるユーザーの店舗在庫データと、別のユーザーの顧客レコード、製品カタログ、またはまったく異なるデータセットを操作する場合があります。これらのシステムでは、データの構造は開発者によって事前に定義されません。これは実行時にエンドユーザーによって定義されます。
これにより、基盤となるデータベースに一連の固有の課題が生じます。システムは、柔軟なスキーマ、ユーザーのデータを分離したままにするための強力なマルチテナンシー、CSV などのファイルの簡単な取り込みとクエリ、そしてますます信頼性の高い LLM 生成の SQL をサポートする必要があります。
この記事では、これらの要件を検討し、構成可能な AI エージェントに最も適したデータベース アーキテクチャを見ていきます。
構成可能なエージェントの中核要件
開発者がデータベース スキーマを定義する従来のアプリケーションとは異なり、構成可能なエージェントはエンド ユーザーが提供する構造化データを操作する必要があります。これにより、いくつかの重要な要件が生じます。
スキーマの柔軟性: 同じエージェントの異なるインスタンスは、アプリケーション コードの変更や移行を必要とせずに、まったく異なるデータ構造を操作できる必要があります。
クエリ可能性: エージェントが質問に答えたり、それに基づいてアクションを実行できるように、データは完全に検索可能でフィルタリング可能である必要があります。
マルチテナント: 異なるユーザーまたは組織に属するデータは、厳密に隔離されたままにしておく必要があります。

テッド。
表形式のデータの取り込み: ユーザーは数百または数千の行を含む CSV をアップロードする場合がありますが、それらの行はクエリ可能なレコードになる必要があります。
LLM 生成の SQL との互換性: 多くのエージェント システムは、LLM に依存して、データの取得に必要な SQL を生成します。 SQL が単純で標準的であればあるほど、LLM がより確実に SQL を生成できます。
これらの要件により、柔軟性、分離性、クエリの信頼性の間に興味深いトレードオフが生じます。
従来のリレーショナル スキーマが難しい理由
事前定義された列を備えた従来のリレーショナル データベースは、開発者がスキーマを制御する場合に非常にうまく機能します。エンドユーザーが構造を定義する場合、さらに困難になります。
ユーザーが必要とする可能性のあるすべての列を事前に定義することはできません。 1 つのデータセットには、price、category、stock が含まれ、別のデータセットには customer_name、email、subscription_status が含まれます。
これが、JSONB などのドキュメント指向で半構造化されたアプローチが最初は魅力的に見える理由です。
Postgres + JSONB: 柔軟性と注意点
PostgreSQL の JSONB データ型は、魅力的な中間点を提供します。 Postgres の成熟したエコシステム、トランザクション、インデックス作成、セキュリティ機能の恩恵を受けながら、任意の構造化データを JSONB 列に保存できます。
JSONB は、 -> 、 ->> 、 @> などの演算子を使用して完全にクエリ可能であり、強力な GIN インデックスをサポートしています。
典型的なデザインは次のようになります。
CREATE TABLE dataset_rows (
id UUID 主キー 、
tenant_id UUID が NULL ではありません、
dataset_id UUID が NULL ではありません、
row_data JSONB NOT NULL
) ;
ユーザーがアップロードした CSV の各行はこのテーブルの行となり、CSV 列は JSONB ドキュメント内に保存されます。
マルチテナントは、tenant_id 列と行レベル セキュリティ (RLS) を併用して処理できます。正しく構成されている場合、データベース自体は c

アプリケーション レベルのフィルタリングに完全に依存するのではなく、テナントの分離を強制します。
このアーキテクチャは、ストレージの観点からはうまく機能します。ただし、データのクエリの主な方法が LLM で生成された SQL を使用する場合、さらに複雑になります。
LLM は一般に、JSONB パス式、型キャスト、および包含演算子を生成する必要がある場合よりも、通常の列ベースの SQL を生成する場合の方が信頼性が高くなります。構文の層が追加されるたびに、モデルに無効​​なクエリまたは不正なクエリを生成する別の機会が与えられます。
マルチテナンシーは解決可能です。動的クエリは避けられない
異なるユーザーからのデータが同じテーブルに混在する可能性は当然の懸念ですが、解決可能な問題です。
tenant_id 列または Agent_id 列を行レベル セキュリティと組み合わせることは、マルチテナント アプリケーションにとって成熟したアプローチです。アプリケーションの要件に応じて、テナントごとのスキーマやテナントごとのデータベースなど、より強力な形式の分離も可能です。
ただし、動的クエリは避けられません。
フィールドはエンド ユーザーによって定義されるため、アプリケーションも LLM も、開発時に既知の固定の列名のセットに依存できません。システムは各データセットのスキーマを理解し、実行時にそれに対するクエリを構築する必要があります。
この課題は、基になるデータが Postgres JSONB、リレーショナル テーブル、ドキュメント データベースのいずれに保存されているかに関係なく存在します。
ユーザーが 1,000 個の製品レコードを含む CSV をアップロードするとします。
システムは、その表形式のデータを、エージェントが効率的にクエリできるものに変換する必要があります。取り込み後、エージェントは次のような質問に答える必要がある場合があります。
価格が ₹500 未満の製品はどれですか?
現在在庫がある電子製品はどれですか?
20 以上の製品を展開しているブランドはどこですか?
ある範囲内での平均価格はいくらですか

特定のカテゴリー?
別のユーザーの CSV には、まったく異なる列が含まれる場合があります。
データセットあたり数千行の規模では、パフォーマンスが決定要因になる可能性は低いです。最新のデータベースのほとんどは、そのボリュームを快適に処理できます。
より重要な考慮事項は、スキーマの柔軟性、分離性、操作の単純さ、および LLM がデータに対して正しいクエリをどれだけ簡単に生成できるかです。
決定的な要因: LLM によって生成された SQL
LLM がクエリの生成を担当すると、評価基準が変わります。
LLM にこれを生成するよう依頼することを検討してください。
WHERE ( row_data - >> '価格' ) :: 数値 < 500
AND row_data - >> 'カテゴリ' = 'エレクトロニクス'
次に、次のものと比較してください。
WHERE 価格 < 500
AND カテゴリ = 'エレクトロニクス'
2 番目のバージョンははるかに単純です。言語モデルが頻繁に遭遇する通常の列名と標準 SQL パターンを使用します。
これは、クエリを実行できるデータベースを選択することが目的ではないため、重要です。目標は、LLM ができるだけ確実に正しいクエリを生成できる環境を作成することです。
これにより、アーキテクチャ上の決定が大きく変わります。
構成可能なエージェントの推奨アーキテクチャ
ユーザー定義スキーマ、マルチテナント、CSV 取り込み、LLM 生成 SQL を組み合わせると、3 つのアプローチが際立ちます。
1. アップロードされた CSV からリレーショナル テーブルを作成する
CSV がアップロードされると、システムはそのヘッダーを検査し、CSV 列に対応する列を含む実際のテーブルを作成できます。その後、行は通常のリレーショナル レコードとして挿入されます。
LLM には、結果のスキーマを指定して、それに対して標準 SQL を生成できます。
これにより、優れたクエリ性が提供され、SQL サーフェイスがシンプルに保たれます。 Postgres がすでにプライマリ データベースである場合、このアプローチにより、チームはその成熟したエコシステムを引き続き使用できます。
ホー

ただし、多数のテーブルを動的に作成すると、特にユーザーとデータセットの数が増加するにつれて、独自の運用上の考慮事項が発生します。テーブルのライフサイクル、命名、移行、権限、クリーンアップはすべて慎重に管理する必要があります。
2. ユーザー定義のデータセットに DuckDB を使用する
DuckDB は、この使用例にとって特に興味深いものです。
クリーンな SQL インターフェイス、分析データおよび表形式データに対する優れたパフォーマンス、CSV や Parquet などの形式に対する最高クラスのサポートを提供します。
任意のユーザー定義データセットをアプリケーションのプライマリ リレーショナル データベースに強制する代わりに、各エージェントまたはデータセットが独自の DuckDB データベース ファイルを持つことができます。その後、エージェントは標準 SQL を使用して通常のテーブルにクエリを実行できます。
これにより、有用な分離が作成されます。
アプリケーションデータ -> Postgres
ユーザー定義エージェントデータ -> DuckDB
エージェント アプリケーションの場合、これは非常に自然なアーキテクチャとなります。各エージェントは事実上、独自の分離されたクエリ可能なデータベースを取得しますが、メイン アプリケーション データベースは引き続きユーザー、エージェント、権限、請求、構成、その他のアプリケーション レベルのデータを管理します。
3. エージェントまたはデータセットごとに SQLite を使用する
SQLite は、同様のファイルベースの分離モデルを提供します。
エージェントまたはデータセットごとに専用の SQLite データベースを作成し、それぞれに独自のスキーマと物理データベース ファイルを与えることができます。 LLM は引き続きクリーンな標準 SQL インターフェイスを取得します。
SQLite はシンプルで成熟しており、非常に信頼性が高くなります。比較的小規模なデータセットに対するトランザクションの読み取りと書き込みを主に必要とするアプリケーションの場合、これは優れた選択肢となる可能性があります。
DuckDB は、ワークロードが表形式データのクエリ、フィルタリング、集計、分析に重点を置いている場合に、より魅力的になります。
より慎重に使用するためのアプローチ
特にすべてを内部に保持する場合、Postgres + JSONB は引き続き実行可能なオプションです

1 つのデータベースは運用上価値があります。
ただし、LLM で生成された SQL が製品の中心である場合、JSONB によりクエリが不必要に複雑になります。強力なプロンプト、スキーマ情報、クエリ検証、および再試行によってこれを補うことができますが、システム内の他の場所で複雑さも増加します。
MongoDB は優れたスキーマの柔軟性を提供し、任意のドキュメント構造で自然に動作します。ただし、SQL は使用しません。アーキテクチャが特に LLM で生成された SQL を中心に設計されている場合、MongoDB には別のクエリ生成戦略か追加の変換層が必要です。
エンドユーザーが提供する任意の構造化データを操作できる AI エージェントを構築すると、データベース層に通常とは異なる要求が課されます。
スキーマの柔軟性だけでは十分ではありません。このアーキテクチャには、強力な分離、表形式データの効率的な取り込みとクエリ、および LLM がクエリを生成する場合のシンプルで予測可能なクエリ サーフェスも必要です。
Postgres は依然としてアプリケーション データの優れた基盤です。ただし、ユーザー定義のエージェント データは、必ずしも同じデータベースに存在する必要はありません。
ユーザーがアップロードした表形式データを中心に構築されたシステムの場合、DuckDB などのエンジン、場合によっては SQLite が興味深い代替手段を提供します。各エージェントは、通常の SQL を LLM に公開しながら、独自のスキーマを備えた独自の分離データベースを効果的に持つことができます。
この組み合わせは、ユーザー向けの柔軟なスキーマ、エージェント間の物理的分離、LLM 用のシンプルな SQL など、エージェント システムにとって特に強力です。
したがって、データベースの決定は、もはやストレージ、スケーラビリティ、クエリのパフォーマンスだけではありません。 LLM を利用したシステムでは、モデルが確実に理解してクエリできるデータ環境を作成することも重要です。
そしてそれは、最終的にはデータベース設計における最も重要な考慮事項の 1 つになる可能性があります。

次世代の構成可能な AI エージェントのオプション。

## Original Extract

Compare Postgres, JSONB, DuckDB, SQLite, and MongoDB for configurable AI agents that ingest user-defined structured data.

Choosing the Right Database for AI Agents: LLM Generated SQL | Learn - Predictable Blog Sign In Create Agent Published on Saturday, August 15, 2026 Choosing the Right Database for AI Agents: LLM Generated SQL
AI agents are rapidly evolving from simple chatbots into configurable systems that can adapt to different domains and data sources. One particularly interesting use case is an agent that users can configure with their own structured data.
An agent might work with store inventory data for one user and customer records, product catalogs, or entirely different datasets for another. In these systems, the structure of the data is not defined by the developer in advance; it is defined by the end user at runtime.
This creates a unique set of challenges for the underlying database. The system needs to support flexible schemas, strong multi-tenancy so that users' data remains isolated, easy ingestion and querying of files such as CSVs, and, increasingly, reliable LLM-generated SQL.
In this article, we'll explore these requirements and look at the database architectures that make the most sense for configurable AI agents.
The Core Requirements of Configurable Agents
Unlike traditional applications, where developers define the database schema, configurable agents need to work with structured data provided by the end user. This creates several important requirements:
Schema flexibility: Different instances of the same agent should be able to work with completely different data structures without requiring application code changes or migrations.
Queryability: The data must remain fully searchable and filterable so the agent can answer questions or take actions based on it.
Multi-tenancy: Data belonging to different users or organizations must remain strictly isolated.
Tabular data ingestion: Users may upload CSVs containing hundreds or thousands of rows, and those rows need to become queryable records.
Compatibility with LLM-generated SQL: Many agentic systems rely on an LLM to generate the SQL needed to retrieve data. The simpler and more standard the SQL is, the more reliably the LLM can generate it.
These requirements create an interesting trade-off between flexibility, isolation, and query reliability.
Why Traditional Relational Schemas Can Be Difficult
A traditional relational database with predefined columns works extremely well when the developer controls the schema. It becomes more difficult when the end user defines the structure.
You cannot define every possible column a user might need in advance. One dataset might contain price , category , and stock , while another contains customer_name , email , and subscription_status .
This is why document-oriented and semi-structured approaches such as JSONB initially appear attractive.
Postgres + JSONB: Flexibility with Caveats
PostgreSQL's JSONB data type provides an appealing middle ground. You can store arbitrary structured data in a JSONB column while still benefiting from Postgres's mature ecosystem, transactions, indexing, and security features.
JSONB is fully queryable using operators such as -> , ->> , and @> , and it supports powerful GIN indexes.
A typical design might look like this:
CREATE TABLE dataset_rows (
id UUID PRIMARY KEY ,
tenant_id UUID NOT NULL ,
dataset_id UUID NOT NULL ,
row_data JSONB NOT NULL
) ;
Each row of a user-uploaded CSV becomes a row in this table, with the CSV columns stored inside the JSONB document.
Multi-tenancy can be handled using a tenant_id column together with Row Level Security (RLS). When configured correctly, the database itself can enforce tenant isolation rather than relying entirely on application-level filtering.
This architecture works well from a storage perspective. However, it introduces additional complexity when the primary way of querying the data is through LLM-generated SQL.
LLMs are generally more reliable when generating ordinary column-based SQL than when they need to produce JSONB path expressions, type casts, and containment operators. Every additional layer of syntax gives the model another opportunity to generate an invalid or incorrect query.
Multi-Tenancy Is Solvable; Dynamic Queries Are Inevitable
The possibility of data from different users mixing in the same table is a legitimate concern, but it is a solvable one.
A tenant_id or agent_id column combined with Row Level Security is a mature approach for multi-tenant applications. Stronger forms of isolation, such as schema-per-tenant or database-per-tenant, are also possible depending on the application's requirements.
Dynamic queries, however, are unavoidable.
Because the fields are defined by the end user, neither the application nor the LLM can rely on a fixed set of column names known at development time. The system needs to understand the schema of each dataset and construct queries against it at runtime.
This challenge exists regardless of whether the underlying data is stored in Postgres JSONB, relational tables, or a document database.
Consider a user uploading a CSV containing 1,000 product records.
The system needs to turn that tabular data into something the agent can efficiently query. After ingestion, the agent may need to answer questions such as:
Which products cost less than ₹500?
Which electronics products are currently in stock?
Which brands have more than 20 products?
What is the average price within a particular category?
Another user's CSV may contain completely different columns.
At the scale of a few thousand rows per dataset, performance is unlikely to be the deciding factor. Most modern databases can comfortably handle that volume.
The more important considerations are schema flexibility, isolation, operational simplicity, and how easily an LLM can generate correct queries against the data.
The Decisive Factor: LLM-Generated SQL
Once an LLM is responsible for generating queries, the evaluation criteria change.
Consider asking an LLM to generate this:
WHERE ( row_data - >> 'price' ) :: numeric < 500
AND row_data - >> 'category' = 'electronics'
Now compare it with:
WHERE price < 500
AND category = 'electronics'
The second version is much simpler. It uses ordinary column names and standard SQL patterns that language models encounter frequently.
This matters because the goal is not merely to choose a database that can execute the query. The goal is to create an environment in which the LLM can generate the correct query as reliably as possible.
That changes the architectural decision considerably.
Recommended Architectures for Configurable Agents
Given the combination of user-defined schemas, multi-tenancy, CSV ingestion, and LLM-generated SQL, three approaches stand out.
1. Create relational tables from uploaded CSVs
When a CSV is uploaded, the system can inspect its headers and create an actual table whose columns correspond to the CSV columns. The rows are then inserted as ordinary relational records.
The LLM can be given the resulting schema and generate standard SQL against it.
This provides excellent queryability and keeps the SQL surface simple. If Postgres is already the primary database, this approach allows teams to continue using its mature ecosystem.
However, dynamically creating large numbers of tables introduces its own operational considerations, particularly as the number of users and datasets grows. Table lifecycle, naming, migrations, permissions, and cleanup all need to be managed carefully.
2. Use DuckDB for user-defined datasets
DuckDB is particularly interesting for this use case.
It provides a clean SQL interface, excellent performance on analytical and tabular data, and first-class support for formats such as CSV and Parquet.
Instead of forcing arbitrary user-defined datasets into the application's primary relational database, each agent or dataset can have its own DuckDB database file. The agent can then query ordinary tables using standard SQL.
This creates a useful separation:
Application data -> Postgres
User-defined agent data -> DuckDB
For agentic applications, this can be a very natural architecture. Each agent effectively gets its own isolated, queryable database while the main application database continues to manage users, agents, permissions, billing, configuration, and other application-level data.
3. Use SQLite per agent or dataset
SQLite offers a similar file-based isolation model.
A dedicated SQLite database can be created for each agent or dataset, giving each one its own schema and physical database file. The LLM still gets a clean, standard SQL interface.
SQLite is simple, mature, and extremely reliable. For applications that primarily need transactional reads and writes against relatively small datasets, it can be an excellent choice.
DuckDB becomes more attractive when the workload is heavily oriented toward querying, filtering, aggregating, and analyzing tabular data.
Approaches to Use More Cautiously
Postgres + JSONB remains a viable option, particularly when keeping everything inside one database is operationally valuable.
However, if LLM-generated SQL is central to the product, JSONB adds unnecessary query complexity. Strong prompting, schema information, query validation, and retries can compensate for this, but they also add complexity elsewhere in the system.
MongoDB provides excellent schema flexibility and works naturally with arbitrary document structures. However, it does not use SQL. If the architecture is specifically designed around LLM-generated SQL, MongoDB requires either a different query-generation strategy or an additional translation layer.
Building an AI agent that can work with arbitrary structured data supplied by the end user places unusual demands on the database layer.
Schema flexibility alone is not enough. The architecture also needs strong isolation, efficient ingestion and querying of tabular data, and, when LLMs generate the queries, a simple and predictable query surface.
Postgres remains an excellent foundation for application data. But user-defined agent data does not necessarily need to live in the same database.
For systems built around user-uploaded tabular data, engines such as DuckDB and, in some cases, SQLite offer an interesting alternative. Each agent can effectively have its own isolated database, with its own schema, while still exposing ordinary SQL to the LLM.
That combination is particularly powerful for agentic systems: flexible schemas for users, physical isolation between agents, and simple SQL for the LLM.
The database decision is therefore no longer just about storage, scalability, or query performance. In an LLM-powered system, it is also about creating a data environment that the model can understand and query reliably.
And that may ultimately be one of the most important database design considerations for the next generation of configurable AI agents.
