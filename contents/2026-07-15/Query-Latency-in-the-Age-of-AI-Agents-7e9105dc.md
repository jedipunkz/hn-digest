---
source: "https://cube.dev/blog/query-latency-in-the-age-of-ai-agents"
hn_url: "https://news.ycombinator.com/item?id=48923643"
title: "Query Latency in the Age of AI Agents"
article_title: "Query Latency in the Age of AI Agents - Cube Blog"
author: "skadamat"
captured_at: "2026-07-15T17:06:16Z"
capture_tool: "hn-digest"
hn_id: 48923643
score: 2
comments: 0
posted_at: "2026-07-15T16:48:25Z"
tags:
  - hacker-news
  - translated
---

# Query Latency in the Age of AI Agents

- HN: [48923643](https://news.ycombinator.com/item?id=48923643)
- Source: [cube.dev](https://cube.dev/blog/query-latency-in-the-age-of-ai-agents)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T16:48:25Z

## Translation

タイトル: AI エージェント時代のクエリ レイテンシ
記事タイトル: AI エージェント時代のクエリ レイテンシー - Cube Blog
説明: AI エージェントの時代においてもクエリの遅延が依然として重要である理由: 集計認識の歴史、キャッシュがセマンティック レイヤーに属する理由、および Cube Store の構築方法。

記事本文:
AI エージェント時代のクエリ レイテンシ
事前集計が依然として重要な理由、キャッシュ レイヤーがセマンティック レイヤーに属する理由、事前集計を提供するために Cube Store がどのように構築されているか
Cube の博士号、共同創設者兼 CTO
Cube の主任ソフトウェア エンジニア
AI にさらなる速度とスループットが必要な理由
セマンティック層でキャッシュし続ける理由
ウェアハウスの外部に別個のキャッシュ層がある理由
より良いデータ製品を構築するために、Cube の更新を受信トレイに入手してください。
どの世代の分析にも同じ問題が潜んでいます。それは、ユーザーが実行したいクエリのコストが、待ち望んでいる待ち時間よりも高いということです。ツールが変わり、ストレージが変わり、クエリ言語が変わりますが、「私が求めたこと」と「どれくらい早くそれが必要か」との間のギャップは決して消えることはありません。 AI エージェントは現在、これまでで分析クエリの最大の消費者となっており、その差は縮まるどころかむしろ広がっています。この投稿では、業界が過去 30 年にわたってそのギャップをどのように埋めてきたか、その答えが事前に計算された集計に戻ってくる理由、およびそれらにサービスを提供する Cube Store をどのように構築したかについて説明します。
オンライン分析処理 (OLAP) システムは、1990 年代に事前計算によって遅延の問題を解決しました。 MOLAP (多次元 OLAP) エンジンは、ファクト テーブルを取得し、ディメンションのセットを選択し、それらのディメンションの組み合わせの集計を専用の多次元構造に格納されるキューブに具体化します。ユーザーが収益を地域別および月別にスライスした場合、答えはロード時にすでに計算されているため、クエリはスキャンではなくルックアップでした。 Essbase とその後の SQL Server Analysis Services は、この考えに基づいて構築されました。一連の質問が事前にわかっており、立方体がそれに合わせて形作られていたため、うまくいきました。
問題は組み合わせです。 n 次元の立方体には 2^n 個の集約が可能です

これらすべてを具体化することは、いくつかの次元を超えると現実的ではありません。これの基本的な扱いは、Harinarayan、Rajaraman、および Ullman によるデータ キューブの効率的な実装 (1996 年) です。この論文では、集計を格子として構成し、固定のストレージ予算と既知のクエリ ワークロードが与えられた場合、残りの部分をそれらから安価に回答できるようにするには、どのビューを具体化する必要があるかという貪欲な質問をしています。この文書は 30 年前のものですが、現在でもすべての事前集計システムが行う正確な決定を記述しています。決定の鍵となるのは作業負荷です。どのクエリに答える必要があるのか​​が分からなければ、具体化するための適切なビューを選択することはできません。
データが増大するにつれて、純粋な MOLAP キューブはスケーリングの限界に達し、業界は分析モデリングを ROLAP エンジンに移行し、その後クラウド データ ウェアハウスに移行しました。ウェアハウスの売り文句は、弾力性のあるカラム型の大規模並列コンピューティングが十分に高速であるため、再度事前計算を行う必要がないというものでした。実際には、事前計算が消えることはありませんでした。それは、マテリアライズド ビューとして、変換ジョブによって構築された集計テーブルとして、ツール独自のメモリ内エンジンに取り込まれた BI 抽出として、そしてウェアハウスにボルト付けされたクエリ結果キャッシュとして再び現れました。用語は「キューブ」、「ロールアップ」、そして「マテリアライズド ビュー」と変化しましたが、メカニズムは同じです。高価な集計を事前に 1 回計算し、そこから対話型クエリを提供します。
AI にさらなる速度とスループットが必要な理由
エージェントがユーザーに代わってデータのクエリを開始すると 2 つのことが変わり、どちらも待ち時間を間違った方向に押し進めます。
一つ目は期待です。現在、人々はチャット ボックスを通じて分析を操作しており、消費者向け製品のレイテンシー バジェットを引き継いでいます。ダッシュボード ユーザーはスピナーを許容しました。

彼らは「クエリが実行中である」ことを知っていました。エージェントの会話では、異なるベースラインが設定されます。数百ミリ秒は反応があるように感じられ、数秒は途切れているように感じられます。表面がメッセージング アプリのように見える場合、計画とスキャンに 8 秒かかるコールド分析クエリに対する許容度は大幅に低くなります。
2 つ目はボリュームで、こちらの方が効果が大きくなります。データを探索する人間は、一度に 1 つずつクエリを発行し、その結果について考えて、次のクエリを発行します。探索モードのエージェントはそうではありません。 1 つの質問に答えるために、それは展開されます。いくつかの列をプロファイリングし、基数を確認し、グループ化を試し、何かに気づき、ドリルダウンし、別の方法で数値を検証し、再定式化します。ユーザーが直面する 1 つのタスクは日常的に数十のクエリに変わり、複数ステップのエージェント ワークフローではそれが数百にまで押し上げられます。これは、人間が手作業で行っていたタスクをマシンが引き継ぐときに必ず現れる通常の増幅パターンであり、クエリの負荷に直接適用されます。システムは、ユーザー セッションごとにさらに多くのクエリを処理する必要があります。これらのクエリはチェーンで実行されるため、各クエリの末尾のレイテンシが合計され、エージェントが応答するまでにかかる合計時間になります。つまり、ストアは、一方を他方と交換するのではなく、スループットと p99 レイテンシを同時に維持する必要があります。
セマンティック層でキャッシュし続ける理由
1996 年のラティスの決定には入力としてワークロードが必要ですが、実際にワークロードを持つスタック内の唯一の場所はセマンティック層です。セマンティック レイヤーはメジャー、ディメンション、結合、グレインが定義される場所であり、クエリが生の列ではなくこれらの定義に基づいて表現される場所です。クエリがステータスによってグループ化された MEASURE(revenue) と日単位の created_at を要求すると、セマンティック レイヤーは正確なメジャー定義、正確な粒度、および正確な結合 p を認識します。

ああ。それはまさにビュー選択アルゴリズムが必要とする情報です。
その決定を倉庫に押し込むと、それははるかに困難で、本質的に監視されていない問題になります。ウェアハウスはメジャーではなく SQL 文字列を認識します。何を事前集計するかを決定するには、生の SQL から繰り返し発生するクエリ パターンをリバース エンジニアリングし、記述方法が異なっていてもどの集計が意味的に同じであるかを推測し、重要な粒度を推測する必要がありますが、それらのパターンを読みやすくするための定義はすべてありません。一部のウェアハウスはマテリアライズド ビューの自動推奨を試みており、制限内で機能しますが、基本的にはセマンティック レイヤーがすでに明示的に持っているワークロード シグナルを回復しようとしています。
Cube はメジャーを認識しているため、それらに関して事前集計が宣言され、1 つのロールアップが集計認識を通じて多くのクエリに対応できます。
キューブのコピー : - 名前 : 注文 sql_table : public.orders メジャー : - 名前 : カウント タイプ : カウント - 名前 : total_amount タイプ : 合計 sql : 金額 ディメンション : - 名前 : ステータス sql : ステータス タイプ : string - 名前 : created_at sql : created_at タイプ : 時間 pre_aggregations : - 名前 :orders_rollup メジャー : - count - total_amount ディメンション : - ステータス time_dimension : created_at granularity : day パーティション粒度 : month
count または total_amount を status ごとに、日単位またはより粗い単位で要求するクエリは、生のテーブルにアクセスするのではなく、orders_rollup から応答されます。セマンティック レイヤーは、受信クエリをロールアップと照合し、事前に集計されたデータを読み取るように書き換えて、日レベルのデータをオンザフライで週または月にロールアップします。ウェアハウスがクエリを認識することはありません。この一致が可能になるのは、集計を格納するコンポーネントがすでにメジャー定義を保持しているため、合計に対するクエリが存在することを認識できるためです。

週単位のステータスによる l_amount は、日単位で構築されたロールアップから回答できます。
ウェアハウスの外部に別個のキャッシュ層がある理由
どの集約を構築するかを知ることができれば、問題の半分は終わります。残りの半分は、すぐに提供できるようにそれらを置く場所です。事前集計をウェアハウスに保管し、そこでクエリを実行しても、レイテンシー フロアはアーキテクチャ上のものであるため、レイテンシ フロアは削除されません。
クラウド データ ウェアハウスは、弾力的にスケールアウトするコンピューティングと、S3 や GCS などのオブジェクト ストアに存在するストレージを備えた、大きなテーブルに対する大規模なスキャンなど、特定の形式の作業向けに構築されています。この設計は大規模な分析ジョブのスループットに優れており、それが倉庫が勝った理由です。また、クエリごとのコストも発生します。このコストは、クエリがどれほど小さいかに関係なく、クエリの計画とコンパイル、コンピューティング クラスターへのスケジュール、他のステートメントの背後でのキューイング、ウォーム化する必要があるリモート ストレージからの読み取りなどです。コールドで開始される任意のクエリ パターンの場合、ほとんどのウェアハウスは 1 秒未満の応答を保証できず、そのように最適化されていないため、そうしようともしません。ワークロードが少数の大規模なレポートである場合、クエリごとの 0.5 秒の固定オーバーヘッドはゼロに丸められます。ワークロードがエージェントからの何百もの小規模な対話型クエリである場合、同じ固定オーバーヘッドが合計応答時間の大半を占めます。
これは、そもそも専用の OLAP データ ストアが存在したのと同じ理由です。インタラクティブ分析のサービング レイヤーには、ウェアハウスとは異なる要件があります。データはすでに集約されているため小さいこと、クエリはフル スキャンではなくポイントまたは範囲のルックアップであること、同時実行時の遅延目標は数十ミリ秒から数百ミリ秒前半です。
この投稿ではキャッシュと OLAP データ ストアを同じ意味で使用しますが、この 2 つは同じものです。

全く同じものではありません。プレーン キャッシュはキーと値のストアです。クエリとその結果を入力すると、その後は同じクエリに対して同じ結果のみを返すことができます。 Cube Store は、より正確にはリレーショナル キャッシュです。完成した回答をクエリごとにキー設定することはありません。完全にクエリ可能なリレーショナル データ (事前に集計されているが、ディメンションとメジャーを含む行と列) が保存されます。その結果、キャッシュされた単一のロールアップから、日レベルのデータを週または月までローリングしたり、サブセットにフィルタリングしたり、別のディメンションで再グループ化したり、別のテーブルに結合したりするなど、キャッシュされていなかった派生計算を実行できるようになります。単純な結果キャッシュでは、クエリ テキストが異なるため、これらがすべて失われます。リレーショナル キャッシュは、保持しているデータから新しい結果を計算することでそれらに答えます。
これら 2 つのシステムは代替品ではなく、どちらも他方に同調することはできません。 OLAP サービング ストアは、生データの処理においてウェアハウスに匹敵するものではありません。ウェアハウスのストレージ経済性を考慮して、テラバイト単位の取り込みと変換、モデル化されていないデータのアドホック スキャン、または完全な履歴の保持を行うようには構築されていません。ウェアハウスは、既知の制限されたクエリ セットでサービスを提供するストアと一致することはありません。エラスティック コンピューティングをどれだけ使用しても、コールド汎用スキャンは、まさにそのクエリ用に構築された、並べ替えられ、事前に集計された小さなテーブルを読み取るのと同じくらい安価または高速になります。すべてをウェアハウスに押し込むと、レイテンシーで余分な費用がかかります。 OLAP ストアからすべてを提供しようとすると、考えられるすべてのクエリ パターンを具体化すると、組み合わせの爆発に直接戻ってしまうため、ストレージ、鮮度、柔軟性に過剰なコストがかかります。効率的な点はどちらの極端にもありません。中間に位置します。生のデータと変換をウェアハウスに保持し、ホットな既知のクエリ パターンをサービス ストアに具体化して、

ch システムは、実際に構築されるワークロードの一部を担います。
総コスト (レイテンシ + 支出) 最低総コスト: ウェアハウス + OLAP ストア 総コスト ウェアハウス コスト (クエリ レイテンシ) 事前集計コスト (ストレージ + 維持) ウェアハウス上のすべてのクエリ (安価なストレージ、提供に時間がかかる) 事前に集計されたものすべて (提供が早い、ストレージが爆発的に増加) OLAP 提供ストアから提供されるクエリ ワークロードのシェア
倉庫ベンダーも同様のギャップに気づき、独自の低レイテンシ サービング レイヤーの出荷を開始しました。たとえば、Snowflake のインタラクティブ テーブルやインタラクティブ ウェアハウスは、1 秒未満の応答を達成するために必須のクラスタリングを備えたポイント ルックアップに最適化されたレイアウトを使用します。 Cube はこれらの上に置くことができ、単一のウェアハウスとそのアクセラレーション レイヤーで標準化されているチームにとっては、それが合理的な選択となる可能性があります。ただし、これをデフォルトにはしません。各ベンダーのアクセラレーション レイヤーには、独自のストレージ形式、独自のパフォーマンス ノブのセット (クラスタリング キー、ウェアハウスのサイジング、リフレッシュ ラグ、サポートされるクエリ形状)、および独自のレイテンシーと障害特性があるため、得られる動作はその 1 つのウェアハウスに固有です。 Cube は、内部 BI と、さまざまなデータソースにわたる組み込み分析の両方を提供します

[切り捨てられた]

## Original Extract

Why query latency still matters in the age of AI agents: the history of aggregate awareness, why caching belongs in the semantic layer, and how Cube Store is built.

Query Latency in the Age of AI Agents
Why pre-aggregations still matter, why the caching layer belongs in the semantic layer, and how Cube Store is built to serve them
PhD, Co-founder and CTO at Cube
Principal Software Engineer at Cube
Why AI needs even more speed and throughput
Why keep caching in the semantic layer
Why have a separate caching layer external to the warehouse
Get Cube updates to your inbox for building better data products.
Every generation of analytics has had the same problem underneath it: the query a person wants to run is more expensive than the latency they are willing to wait for. The tools change, the storage changes, the query language changes, but that gap between "what I asked" and "how fast I need it" has never gone away. AI agents are now the largest consumers of analytical queries we have ever had, and they have made that gap wider, not narrower. This post is about how the industry has closed that gap over the last thirty years, why the answer keeps coming back to pre-computed aggregates, and how we built Cube Store to serve them.
Online Analytical Processing (OLAP) systems solved the latency problem in the 1990s by pre-computing. A MOLAP (multidimensional OLAP) engine took a fact table, picked a set of dimensions, and materialized the aggregates for combinations of those dimensions into a cube stored in a dedicated multidimensional structure. When a user sliced revenue by region and month, the answer had already been computed at load time, so the query was a lookup rather than a scan. Essbase, and later SQL Server Analysis Services, were built on this idea. It worked because the set of questions was known in advance and the cube was shaped to match.
The problem is combinatorial. A cube over n dimensions has 2^n possible aggregation groupings, and materializing all of them is not feasible past a handful of dimensions. The foundational treatment of this is Harinarayan, Rajaraman, and Ullman, Implementing Data Cubes Efficiently (1996) , which frames the aggregates as a lattice and asks a greedy question: given a fixed storage budget and a known query workload, which views should you materialize so that the rest can be answered cheaply from them? That paper is thirty years old and it still describes the exact decision every pre-aggregation system makes today. The key input to the decision is the workload. You cannot pick the right views to materialize without knowing which queries you need to answer.
As data grew, the pure MOLAP cube ran into scaling limits, and the industry moved analytical modeling into ROLAP engines and, later, into cloud data warehouses. The warehouse pitch was that elastic, columnar, massively parallel compute would be fast enough that you would never need to pre-compute anything again. In practice the pre-computation never disappeared. It reappeared as materialized views, as aggregate tables built by transformation jobs, as BI extracts pulled into a tool's own in-memory engine, and as query result caches bolted onto the warehouse. The vocabulary changed from "cube" to "rollup" to "materialized view," but the mechanism is identical: compute the expensive aggregate once, ahead of time, and serve the interactive query from it.
Why AI needs even more speed and throughput
Two things changed when agents started querying data on behalf of users, and both push latency in the wrong direction.
The first is expectation. People now interact with analytics through a chat box, and they carry over the latency budget of a consumer product. A dashboard user tolerated a spinner because they had learned that "the query is running." An agent conversation sets a different baseline: a few hundred milliseconds feels responsive, a couple of seconds feels broken. The tolerance for a cold analytical query that takes eight seconds to plan and scan is much lower when the surface looks like a messaging app.
The second is volume, and it is the larger effect. A human exploring data issues queries one at a time, thinks about the result, and issues the next one. An agent in exploration mode does not. To answer one question it will fan out: profile a few columns, check cardinalities, try a grouping, notice something, drill in, verify the number a different way, then reformulate. A single user-facing task routinely turns into tens of queries, and multi-step agent workflows push that into the hundreds. This is the usual amplification pattern that shows up whenever a machine takes over a task a human used to do by hand, and it applies directly to query load. The system now has to serve many more queries per user session, and because those queries run in a chain, the tail latency of each one adds up into the total time the agent takes to answer. That means the store has to hold up on throughput and on p99 latency together, not trade one for the other.
Why keep caching in the semantic layer
The lattice decision from 1996 needs a workload as input, and the semantic layer is the only place in the stack that actually has it. The semantic layer is where measures, dimensions, joins, and grains are defined, and it is where queries are expressed in terms of those definitions rather than raw columns. When a query asks for MEASURE(revenue) grouped by status and created_at at day granularity, the semantic layer knows the exact measure definition, the exact grain, and the exact join path. That is precisely the information the view-selection algorithm needs.
Push that decision down to the warehouse and it becomes a much harder, essentially unsupervised problem. The warehouse sees SQL strings, not measures. To decide what to pre-aggregate it would have to reverse-engineer the recurring query patterns from raw SQL, infer which aggregations are semantically the same despite being written differently, and guess the grains that matter, all without the definitions that would make those patterns legible. Some warehouses do attempt automatic materialized view recommendation, and it works within limits, but it is fundamentally trying to recover a workload signal that the semantic layer already has explicitly.
Because Cube knows the measures, a pre-aggregation is declared in terms of them, and one rollup can serve many queries through aggregate awareness:
Copy cubes : - name : orders sql_table : public.orders measures : - name : count type : count - name : total_amount type : sum sql : amount dimensions : - name : status sql : status type : string - name : created_at sql : created_at type : time pre_aggregations : - name : orders_rollup measures : - count - total_amount dimensions : - status time_dimension : created_at granularity : day partition_granularity : month
Any query that asks for count or total_amount , by status , at day granularity or coarser, is answered from orders_rollup instead of touching the raw table. The semantic layer matches the incoming query against the rollup, rewrites it to read pre-aggregated data, and rolls day-level data up to week or month on the fly. The warehouse never sees the query. That match is possible because the component storing the aggregate already holds the measure definition, so it can recognize that a query for total_amount by status at week granularity is answerable from a rollup built at day granularity.
Why have a separate caching layer external to the warehouse
Knowing which aggregates to build is half the problem. The other half is where to put them so they can be served fast. Storing pre-aggregations back in the warehouse and querying them there does not remove the latency floor, because the floor is architectural.
Cloud data warehouses are built for a specific shape of work: large scans over large tables, with compute that scales out elastically and storage that lives in object stores like S3 or GCS. That design is excellent for throughput on big analytical jobs, and it is the reason warehouses won. It also imposes a per-query cost that does not go away no matter how small the query is: query planning and compilation, scheduling onto a compute cluster, queueing behind other statements, and reading from remote storage that has to be warmed. For an arbitrary query pattern starting cold, most warehouses cannot guarantee a sub-second response, and they do not try to, because that is not what they were optimized for. When your workload is a few large reports, a half-second of fixed overhead per query rounds to nothing. When your workload is hundreds of small interactive queries from an agent, that same fixed overhead dominates the total response time.
This is the same reason dedicated OLAP data stores existed in the first place. A serving layer for interactive analytics has different requirements from a warehouse: the data is already aggregated and therefore small, the query is a point or range lookup rather than a full scan, and the latency target is tens to low hundreds of milliseconds under concurrency.
This post uses cache and OLAP data store interchangeably, but the two are not quite the same thing. A plain cache is a key-value store: you put in a query and its result, and later it can only return that same result for that same query. Cube Store is more precisely a relational cache. It does not key finished answers by query. It stores relational data (pre-aggregated, but still rows and columns with dimensions and measures) that stays fully queryable. The consequence is that from a single cached rollup you can run derived calculations that were never cached as such: rolling day-level data up to weeks or months, filtering to a subset, regrouping by a different dimension, or joining to another table. A plain result cache would miss all of those because the query text differs. A relational cache answers them by computing new results from the data it holds.
These two systems are not substitutes, and neither can be tuned into the other. An OLAP serving store will never match a warehouse at raw data processing: it is not built to ingest and transform terabytes, scan unmodeled data ad hoc, or hold full history at warehouse storage economics. The warehouse, in turn, will never match the serving store on a known, bounded set of queries: no amount of elastic compute makes a cold general-purpose scan as cheap or as fast as reading a small, sorted, pre-aggregated table built for exactly that query. Push everything onto the warehouse and you overpay in latency. Try to serve everything from an OLAP store and you overpay in storage, freshness, and flexibility, because materializing every possible query pattern runs straight back into the combinatorial explosion. The efficient point is not at either extreme. It sits in the middle: keep raw data and transformations in the warehouse, materialize the hot, known query patterns into a serving store, and let each system carry the part of the workload it is actually built for.
Total cost (latency + spend) Lowest total cost: warehouse + OLAP store Total cost Warehouse cost (query latency) Pre-aggregation cost (storage + upkeep) All queries on the warehouse (cheap storage, slow to serve) Everything pre-aggregated (fast to serve, storage explodes) Share of the query workload served from an OLAP serving store
Warehouse vendors have noticed the same gap and started shipping their own low-latency serving layers, for example Snowflake's interactive tables and interactive warehouses, which use a point-lookup-optimized layout with mandatory clustering to reach sub-second responses. Cube can sit on top of these, and for a team standardized on a single warehouse and its acceleration layer, that can be a reasonable choice. We do not make it the default, though. Each vendor's acceleration layer has its own storage format, its own set of performance knobs (clustering keys, warehouse sizing, refresh lag, supported query shapes), and its own latency and failure characteristics, so the behavior you get is specific to that one warehouse. Cube serves both internal BI and embedded analytics across many different data sour

[truncated]
