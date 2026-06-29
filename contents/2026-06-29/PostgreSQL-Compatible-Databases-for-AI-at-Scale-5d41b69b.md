---
source: "https://www.cockroachlabs.com/blog/postgresql-compatible-databases-ai-scale/"
hn_url: "https://news.ycombinator.com/item?id=48718811"
title: "PostgreSQL-Compatible Databases for AI at Scale"
article_title: "PostgreSQL-Compatible Databases for AI at Scale"
author: "ilreb"
captured_at: "2026-06-29T13:24:39Z"
capture_tool: "hn-digest"
hn_id: 48718811
score: 2
comments: 0
posted_at: "2026-06-29T13:06:39Z"
tags:
  - hacker-news
  - translated
---

# PostgreSQL-Compatible Databases for AI at Scale

- HN: [48718811](https://news.ycombinator.com/item?id=48718811)
- Source: [www.cockroachlabs.com](https://www.cockroachlabs.com/blog/postgresql-compatible-databases-ai-scale/)
- Score: 2
- Comments: 0
- Posted: 2026-06-29T13:06:39Z

## Translation

タイトル: 大規模な AI 向けの PostgreSQL 互換データベース
説明: 分散 SQL から回復力のあるパフォーマンスのためのベクトル検索まで、大規模な AI のために PostgreSQL 互換データベースで何を評価するかを学びます。

記事本文:
大規模な AI 向けの PostgreSQL 互換データベース 製品 ソリューション リソース 価格 会社 お問い合わせ
大規模な AI 向けの PostgreSQL 互換データベース: 初日から何を評価すべきか
AI ワークロードは、従来のアプリケーションよりも早くデータベースの制限を明らかにします
PostgreSQL との互換性だけでは水平方向のスケールは保証されません
分散 SQL は、再アーキテクチャを必要とせずに AI アプリケーションを拡張できるようにします
AI プロジェクトの開始時に選択したデータベースは、何年も一緒に暮らすか、お金を払って逃れることになるデータベースです。これは AI に限ったことではありませんが、AI はタイムラインを変えます。Cockroach Labs の「2026 年の AI インフラストラクチャの現状」レポートによると、エンジニアリング リーダーの 83% が、24 か月以内に大規模なアップグレードを行わないと、AI 主導の需要によりデータ インフラストラクチャに障害が発生すると考えています。ほぼ 3 分の 2 が、指導チームがその限界点にどれだけ早く到達するかを過小評価していると回答しています。
大規模な AI アプリケーションを構築するチームにとって、問題は単にデータベースが PostgreSQL 構文をサポートしているかどうかではありません。重要なのは、PostgreSQL 互換データベースが、本番環境で AI ワークロードによって作成される同時実行性、一貫性、スケール パターンを処理できるかどうかです。
AI データベースの決定がグリーンフィールド段階で重要なのはなぜですか?
「AI を活用した新しい製品の範囲を検討しているエンジニアリング担当副社長やプラットフォーム責任者にとって、グリーンフィールド期間は、プロジェクトのライフサイクルの中で最も大きな影響力を発揮する瞬間です」と、Cockroach Labs のセールス エンジニアリング担当シニア マネージャーの David Joy 氏は述べています。 「本番トラフィック、顧客データ、組織の依存関係がその上に蓄積されると、あらゆるアーキテクチャ上の決定を元に戻すのは難しくなり、よりコストがかかります。『PostgreSQL から始めて、後でスケールする』というデフォルトのアドバイスは、限界が来るのを確認する時間を与える緩やかな成長曲線を想定しています。AI ワークロードはその曲線に従いません。

」
これは核心的な緊張を反映しています。最も一般的な推奨事項は、ワークロードが従来のアプリケーションのように動作しない場合にアーキテクチャ上の負債を生み出す可能性が最も高い推奨事項でもあります。
PostgreSQL の互換性と PostgreSQL の水平スケーリングは同じものではありません。データベースは、使い慣れた PostgreSQL 構文、ドライバー、ツールをサポートできますが、それでも規模の上限に達し、後で移行が必要になります。
AI ワークロードはなぜ従来のアプリケーションよりも早くデータベースの制限を表面化するのでしょうか?
従来のアプリケーションは、人間のペースで予測可能な負荷を生成します。ユーザーはログインし、データをクエリし、フォームを送信して、応答を待ちます。アクション間の休止により、インフラストラクチャに需要を吸収する余地が生まれます。 AI ワークロードはこれらの一時停止を排除し、ほとんどのデータベースが処理するように設計されていなかった同時実行パターンを導入します。
「単一の AI エージェントのリクエストが、ツール、API、データベースへの数十のダウンストリーム呼び出しに広がる可能性があります」と Joy 氏は説明します。 「これを数千の同時エージェント セッションで掛け合わせると、従来のトラフィックとはまったく異なる負荷プロファイルが得られます。セッション間でコンテキストが同時に期限切れになるとキャッシュ スタンピードが発生し、ダウンストリーム サービスが停止すると再試行ストームが発生し、並列サブタスクが同時に共有状態に収束すると書き込みホットスポットが発生します。」
ベクトル検索により、別の次元が追加されます。トランザクション操作とベクトル類似性クエリ (検索拡張生成 (RAG)、レコメンデーション エンジン、セマンティック検索) を組み合わせる AI アプリケーションでは、一貫性を保つためにベクトルの埋め込みと操作データが必要になることがよくあります。ベクトル検索が別のシステムにある場合、チームは 2 つのデータ層にわたる同期、重複、障害モードを管理する必要があります。新しいコンテキストに依存する AI 製品の場合、より強力なアーキテクチャは、

または、検索データとトランザクション データは同じ一貫性モデルの下で動作できます。
これらは、何年にもわたる漸進的な成長の後に表面化する問題ではありません。代わりに、エージェントの同時実行性とベクトル ワークロードが単一ノード データベースの上限を超える運用開始の最初の数か月間に発生します。
エージェント AI が本番環境に到達すると、何が壊れるでしょうか? — エージェント システムが実際のユーザーとの接触に耐えられるかどうかを決定する 6 つのインフラストラクチャの課題 (メモリ状態、雷鳴の群れ、エージェント ID、爆発範囲、可観測性、経済性) を詳しく考察します。
単一の PostgreSQL インスタンスが AI ワークロードに課す制限は何ですか?
PostgreSQL は実績のある強力なデータベースです。これは、多くの従来のワークロードにとって適切な開始点です。ただし、そのシングルノード アーキテクチャには、AI ワークロードがほとんどのチームの予想よりも早く到達するという厳しい制限が課せられています。 AI における単一ノード データベース アーキテクチャの主な制約は次のとおりです。
接続制限 – 単一の PostgreSQL インスタンスが処理できる同時接続数は限られています。 AI エージェントのワークロードでは、各セッションで複数の同時クエリが生成されるため、すぐに上限を使い果たす可能性があります。接続プーリングは役立ちますが、アーキテクチャ上の制限を取り除くのではなく、パッチを当てます。
書き込みスループット – 単一ノードの PostgreSQL は垂直方向にスケーリングします。CPU とメモリをマシンに追加しますが、書き込みを複数のノードまたは地域に分散することはできません。エージェントの実行レコード、状態チェックポイント、およびベクター更新がすべて 1 つのインスタンスに集中すると、垂直方向のヘッドルームが不足します。
障害後の一貫性 – これは AI エージェントにとってより深刻な問題であり、即座に整合性のある作業メモリ、ノード障害が発生しても存続するエピソード メモリ、バージョン管理され耐久性のあるセマンティック メモリなど、さまざまな種類のメモリが必要です。歌う

le-node データベースは、複数のノードにまたがるように設計されていないため、これらの要件を管理するための分散整合性プリミティブを提供しません。インスタンスがダウンすると、それに依存するすべてのエージェント セッションも一緒にダウンします。これは「最終的な」スケーリングの問題ではなく、エージェントを展開するたびに悪化する脆弱性です。
ビジネスへの影響は、クエリの速度低下だけではありません。エンジニアが接続圧力、垂直スケーリング制限、フェイルオーバー リスクの管理にロードマップ時間を費やすと、チームは製品の速度を失います。顧客は、それを一貫性のない AI エクスペリエンス、応答の遅れ、またはダウンタイムとして感じます。
AI ワークロードが増大すると、アーキテクチャ上の負債はどのように増大するのでしょうか?
急速に増大する AI ワークロードによって課せられる技術的な制限が問題の 1 つです。彼らと同居し、その後プレッシャーを受けながら移住することによる組織コストもまた別だ。
実稼働データベースの移行は、チームが取り組むことのできるエンジニアリング プロジェクトの中で最もリスクが高いものの 1 つです。主要なフェーズには次のものが含まれます。
何か問題が発生する可能性があるカットオーバーウィンドウ
従来のアプリケーションが予測可能な速度で成長する場合、チームは独自のタイムラインに基づいて移行を計画できます。ただし、AI ワークロードはそのタイムラインを圧縮します。エンジニアリング リーダーの 100% が AI ワークロードの増大を予想しており、98% が AI 関連のダウンタイムが 1 時間でも少なくとも 10,000 ドルかかると回答している場合、移行は緊急であり、費用がかかり、一か八かのリスクが伴うものになります。
コストはエンジニアリング時間だけではありません。それは、チームがインフラストラクチャを再構築している間は出荷されない製品機能と、カットオーバー中に 2 つのデータベース システムを並行して実行する運用上のリスクです。これは、あるシステムに関する専門知識を構築したチームが、本番環境のプレッシャーの下で別のシステムを学習する必要がある場合の再トレーニングのコストでもあります。
これが、建築上の負債が直線的にではなく指数関数的に増大する仕組みです。

つまり、ワークロードの増大が、その下の基盤を置き換えるチームの能力を上回る速さで増大するためです。
大規模な AI のために、チームは PostgreSQL 互換データベースで何を評価すべきでしょうか?
「デフォルトのパスが強制的な移行につながる場合、問題は『同じ上限を課さない PostgreSQL 互換データベースにチームは何を求めるべきか』ということになります」と Joy 氏は言います。 「答えは、後から追加される機能ではなく、初日から AI ワークロードをサポートするアーキテクチャ機能に帰着します。」
AI データベース ソリューションを評価する場合、チームは機能チェックリストを超えて、基盤となるアーキテクチャが読み取り、書き込み、トランザクション、ベクトル検索を同時に拡張できるかどうかを検討する必要があります。実際の評価は次のとおりです。6 つのアーキテクチャ機能があり、それぞれが具体的なビジネス成果に結びついています。
関連 SQL データベース最新化のためのアーキテクト ガイド – シングルリージョン、マルチリージョン、マルチクラウドの成熟段階をカバーする段階的な最新化ロードマップ。データベースの移行を設計、移行、検証するための実用的なフレームワークを見つけてください。
分散 SQL は AI インフラストラクチャの要件にどのように対処しますか?
分散 SQL データベースは、SQL アクセスと ACID 保証を維持しながら、データ、クエリ、トランザクションを複数のノードに分散するリレーショナル データベースです。このアーキテクチャは AI ワークロードにとって重要です。チームがトランザクションの一貫性を放棄したり、手動シャーディングを中心にアプリケーションを書き直したりすることなく、データベースを水平方向に拡張できるからです。
AI ワークロードの場合、分散 SQL は、単一ノード システムで表面化する次のような限界点に直接対処します。
接続圧力をノード間で分散できる
書き込みスループットは容量を追加することで拡張可能
分散型コンセンサスは一貫性プリミティブを提供します

さまざまなエージェント メモリ タイプに必要なもの
すべての PostgreSQL 互換データベースが分散されているわけではなく、すべての分散データベースが完全な ACID 保証を維持しているわけではないため、この区別は重要です。マネージド単一リージョン PostgreSQL サービスは、運用上の利便性を提供しますが、同じアーキテクチャの上限を継承します。 NoSQL システムは水平スケールを提供しますが、トランザクションの一貫性は犠牲になります。分散 SQL は、両方の要件が交渉不可能な領域を占有します。初日から一貫性があり、耐久性があり、水平方向にスケーラブルなインフラストラクチャを必要とする AI ワークロードにとって、分散 SQL はそれらの要件に最も適合するアーキテクチャ モデルです。
Vector Search Meets Distributed SQL for AI Workloads – AI ワークロードが従来のデータ アーキテクチャを破壊している理由、分散ベクトル インデックス作成の仕組み、ベクトル データとトランザクション データを 1 つのプラットフォームで統合することで大規模な複雑性がどのように軽減されるかを調査した Intellyx アナリスト レポート。
データベース アーキテクチャが他の AI 投資よりも長持ちするのはなぜですか?
ここで説明する分散 SQL の基準 (弾力的な水平スケール、分散 ACID トランザクション、ネイティブ ベクター検索、マルチリージョン デプロイメント、ゼロ ダウンタイムのスキーマ変更、完全な PostgreSQL 互換性) は、機能のウィッシュリストではありません。これらは、AI ワークロードが初日からデータ層に課すアーキテクチャ要件です。
CockroachDB は、PostgreSQL の互換性を分散 SQL の基礎と拡張することにより、そのアーキテクチャを具体化します。自動データ分散、ノード間でのシリアル化可能な分離、手動シャーディングを必要としない柔軟なスケールは、C-SPANN 分散ベクトル インデックス作成や構造化されたエージェント インタラクションのためのマネージド MCP サーバーなどの AI ネイティブ機能と連携して機能します。
データ層は AI 製品の中で最も再現が難しい部分の 1 つであるため、この組み合わせが重要になります。

打ち上げ後のエース。水平方向にスケーラブルで PostgreSQL と互換性のあるアーキテクチャを早期に選択すると、移行リスクが軽減され、エンジニアリングの焦点が保護され、基盤を再構築することなく AI 製品を進化させる余地がチームに与えられます。
「CockroachDB は、制限に達した後に移行するデータベースではありません。そもそも移行の必要がなくなります」と Joy 氏は言います。 「モデルが変更され、フレームワークが進化しても、データ層はそのすべての下で引き続き実行されます。再構築せずに AI 製品を出荷するチームは、AI が作り出す条件に合わせて構築されたアーキテクチャを選択し、グリーンフィールドのウィンドウが閉じる前にその選択を行ったチームです。アーキテクチャは、AI に対する他のあらゆる賭けよりも長持ちするため、それに応じて選択してください。」
CockroachDB が、再構築せずに AI アプリケーションを拡張するのにどのように役立つかを学びましょう。専門家に相談してください。
運用データ、ベクトル検索、永続的なエージェントの状態を、回復力のある 1 つの分散 SQL データベースに統合します。 $400 の無料クレジットから始めましょう。
40 か国以上のフォーチュン 50 金融機関およびチームから信頼されています。
アプリアーキテクチャのビッグアイデアで David Joy 氏の詳細をお聞きください。これは、最新のデータ集約型アプリケーションやシステムを構築しているアーキテクトやエンジニア向けのポッドキャストです。
David Weiss は、Cockroa のシニア テクニカル コンテンツ マーケティング担当者です

[切り捨てられた]

## Original Extract

Learn what to evaluate in PostgreSQL-compatible databases for AI at scale, from distributed SQL to vector search for resilient performance.

PostgreSQL-Compatible Databases for AI at Scale Product Solutions Resources Pricing Company Contact us
PostgreSQL-Compatible Databases for AI at Scale: What to Evaluate from Day One
AI workloads expose database limits faster than traditional applications
PostgreSQL compatibility alone does not guarantee horizontal scale
Distributed SQL helps AI applications scale without re-architecture
The database you choose at the start of an AI project is the one you'll be living with, or paying to escape, for years. That's not unique to AI, but AI changes the timeline: According to the Cockroach Labs State of AI Infrastructure 2026 report, 83% of engineering leaders believe AI-driven demand will cause their data infrastructure to fail without major upgrades within 24 months. Nearly two-thirds say their leadership teams underestimate how quickly that breaking point will arrive.
For teams building AI applications at scale, the question is not simply whether a database supports PostgreSQL syntax. It’s whether that PostgreSQL-compatible database can handle the concurrency, consistency, and scale patterns AI workloads create in production.
Why does your AI database decision matter at the greenfield stage?
“For a VP Engineering or Head of Platform scoping a new AI-powered product, the greenfield window is the highest-leverage moment in the project's lifecycle,” says David Joy, Senior Manager, Sales Engineering at Cockroach Labs. “Every architectural decision gets harder and more expensive to reverse once production traffic, customer data, and organizational dependencies accumulate on top of it. The default advice of ‘start with PostgreSQL, scale later’ assumes a gradual growth curve that gives you time to see limits coming. AI workloads don't follow that curve.”
That reflects the core tension: The most common recommendation is also the one most likely to create architectural debt when the workload doesn't behave like a traditional application.
PostgreSQL compatibility and PostgreSQL horizontal scaling are not the same thing. A database can support familiar PostgreSQL syntax, drivers, and tools while still hitting a scale ceiling that forces a migration later.
Why do AI workloads surface database limitations faster than traditional applications?
Traditional applications generate predictable, human-paced load. A user logs in, queries data, submits a form, and waits for a response. The pauses between actions give infrastructure room to absorb demand. AI workloads eliminate those pauses, and introduce concurrency patterns that most databases were never designed to handle.
“A single AI agent request can fan out into dozens of downstream calls to tools, APIs, and databases,” Joy explains. “Multiply that across thousands of concurrent agent sessions, and you get load profiles that look nothing like traditional traffic: Cache stampedes when context expires simultaneously across sessions, retry storms when a downstream service hiccups, and write hotspots when parallel subtasks converge on shared state at the same moment.”
Vector search adds another dimension. AI applications that combine transactional operations with vector similarity queries – retrieval-augmented generation (RAG), recommendation engines, semantic search – often need vector embeddings and operational data to stay consistent together. If vector search lives in a separate system, teams must manage synchronization, duplication, and failure modes across two data layers. For AI products that depend on fresh context, the stronger architecture is one where vector search and transactional data can operate under the same consistency model.
These aren't problems that surface after years of incremental growth. Instead, they emerge in the first months of production, when agent concurrency and vector workloads push past the ceiling of a single-node database.
What breaks when agentic AI reaches production? — A deeper look at the six infrastructure challenges — memory state, thundering herd, agent identity, blast radius, observability, and economics — that determine whether agentic systems survive contact with real users.
What limits does a single PostgreSQL instance impose on AI workloads?
PostgreSQL is a proven, powerful database; it's the right starting point for many traditional workloads. However, its single-node architecture imposes hard limits that AI workloads hit faster than most teams anticipate. Here are the primary constraints of a single-node database architecture on AI:
Connection limits – A single PostgreSQL instance handles a finite number of concurrent connections. AI agent workloads, where each session can generate multiple simultaneous queries, can exhaust that ceiling quickly. Connection pooling helps, but it patches an architectural limit rather than removing it.
Write throughput – Single-node PostgreSQL scales vertically : You add CPU and memory to the machine, but you can't distribute writes across multiple nodes or geographies. When agent execution records, state checkpoints, and vector updates all converge on one instance, vertical headroom runs out.
Consistency across failure – This is a deeper issue for AI agents, which require different types of memory: working memory that's immediately consistent, episodic memory that survives node failures, semantic memory that's versioned and durable. A single-node database doesn't provide distributed consistency primitives to manage these requirements, because it was never designed to span multiple nodes. When the instance goes down, every agent session that depends on it goes down with it. That isn’t an “eventual” scaling problem, it's a fragility that compounds with every agent you deploy.
The business impact shows up as more than slower queries. Teams lose product velocity when engineers spend roadmap time managing connection pressure, vertical scaling limits, and failover risk. Customers feel it as inconsistent AI experiences, delayed responses, or downtime.
How does architectural debt compound when AI workloads grow?
The technical limits imposed by fast-growing AI workloads are one problem. The organizational cost of living with them, and then migrating under pressure, is another.
A production database migration is one of the highest-risk engineering projects a team can undertake. Key phases include:
a cutover window where something can go wrong
For a traditional application growing at a predictable rate, teams can plan that migration on their own timeline. However, AI workloads compress that timeline. When 100% of engineering leaders expect AI workloads to grow – and 98% say even one hour of AI-related downtime costs at least $10,000 – the migration becomes urgent, expensive, and high-stakes.
The cost isn't just engineering hours. It's the product features that don't ship while the team is re-architecting infrastructure, and the operational risk of running two database systems in parallel during cutover. It's also the retraining cost for a team that built expertise around one system, and now has to learn another under production pressure.
This is how architectural debt compounds: not linearly, but exponentially, as the workload grows faster than the team's ability to replace the foundation beneath it.
What should teams evaluate in a PostgreSQL-compatible database for AI at scale?
“If the default path leads to forced migrations, the question becomes: ‘What should teams look for in a PostgreSQL-compatible database that won't impose the same ceiling?’” says Joy. “The answer comes down to architectural capabilities that support AI workloads from day one, not features bolted on after the fact.”
When evaluating AI database solutions, teams should look beyond feature checklists and ask whether the underlying architecture can scale reads, writes, transactions, and vector search together. Here's what that evaluation looks like in practice: six architectural capabilities, each connected to a concrete business outcome.
Related The Architect's Guide to SQL Database Modernization – A step-by-step modernization roadmap covering single-region, multi-region, and multi-cloud maturity stages. Discover practical frameworks for designing, migrating, and validating your database transition.
How does distributed SQL address AI infrastructure requirements?
A distributed SQL database is a relational database that distributes data, queries, and transactions across multiple nodes while preserving SQL access and ACID guarantees. That architecture matters for AI workloads because the database can scale horizontally, without forcing teams to give up transactional consistency or rewrite applications around manual sharding.
For AI workloads, distributed SQL directly addresses the breaking points that surface with single-node systems:
connection pressure can be distributed across nodes
write throughput can scale by adding capacity
distributed consensus provides the consistency primitives that different agent memory types require
The distinction matters because not every PostgreSQL-compatible database is distributed, and not every distributed database maintains full ACID guarantees. Managed single-region PostgreSQL services offer operational convenience but inherit the same architectural ceiling. NoSQL systems offer horizontal scale but sacrifice transactional consistency. Distributed SQL occupies the space where both requirements are non-negotiable. For AI workloads that demand consistent, durable, horizontally scalable infrastructure from day one, distributed SQL is the architectural model that best fits those requirements.
Vector Search Meets Distributed SQL for AI Workloads – Intellyx analyst report exploring why AI workloads are breaking conventional data architectures, how distributed vector indexing works, and how unifying vector and transactional data on one platform reduces complexity at scale.
Why does database architecture outlast other AI investments?
The distributed SQL criteria discussed here – elastic horizontal scale, distributed ACID transactions, native vector search, multi-region deployment, zero-downtime schema changes , and full PostgreSQL compatibility – aren't a feature wishlist. These are the architectural requirements that AI workloads impose on the data layer from day one.
CockroachDB embodies that architecture by extending PostgreSQL compatibility with distributed SQL fundamentals. Automatic data distribution, serializable isolation across nodes, and elastic scale without manual sharding work alongside AI-native capabilities like C-SPANN distributed vector indexing and a managed MCP Server for structured agent interaction.
That combination matters because the data layer is one of the hardest parts of an AI product to replace after launch. Choosing a horizontally scalable, PostgreSQL-compatible architecture early reduces migration risk, protects engineering focus, and gives teams more room to evolve their AI product without rebuilding the foundation.
“CockroachDB isn’t a database you migrate to after hitting limits – it eliminates the need for that migration in the first place,” Joy says. “As models change and frameworks evolve, the data layer will still be running beneath all of it. The teams that ship AI products without re-architecting are the ones that chose an architecture built for the conditions AI creates, and made that choice before the greenfield window closed. Architecture outlasts every other bet you make on AI, so choose accordingly.”
Learn how CockroachDB helps AI applications scale without re-architecting. Speak with an expert .
Unify operational data, vector search, and durable agent state in one resilient, distributed SQL database. Start with $400 in free credits.
Trusted by Fortune 50 financial institutions and teams in 40+ countries.
Hear more from David Joy at BIG IDEAS IN APP ARCHITECTURE ! It’s the podcast for architects and engineers who are building modern, data-intensive applications and systems.
David Weiss is Senior Technical Content Marketer for Cockroa

[truncated]
