---
source: "https://zilliz.com/blog/from-vector-database-to-vector-lakebase"
hn_url: "https://news.ycombinator.com/item?id=48407201"
title: "A Vector Lakebase is all you need for all AI workloads"
article_title: "From Vector Database to Vector Lakebase - Zilliz blog"
author: "Fendy"
captured_at: "2026-06-05T04:34:36Z"
capture_tool: "hn-digest"
hn_id: 48407201
score: 2
comments: 0
posted_at: "2026-06-05T02:20:54Z"
tags:
  - hacker-news
  - translated
---

# A Vector Lakebase is all you need for all AI workloads

- HN: [48407201](https://news.ycombinator.com/item?id=48407201)
- Source: [zilliz.com](https://zilliz.com/blog/from-vector-database-to-vector-lakebase)
- Score: 2
- Comments: 0
- Posted: 2026-06-05T02:20:54Z

## Translation

タイトル: すべての AI ワークロードに必要なのは Vector Lakebase だけです
記事タイトル: Vector Database から Vector Lakebase へ - Zilliz ブログ
説明: Zilliz は、リアルタイムのベクトル検索、湖規模の発見、および Al を統合する、Milvus を利用したフルマネージドの Vector Lakebase を提供します。
データ操作。

記事本文:
Vector Database から Vector Lakebase へ - Zilliz ブログ 製品 Zilliz Cloud
大規模な高い信頼性、パフォーマンス、コスト効率を実現するように設計された、フルマネージドの Vector Lakebase サービス。
10 億規模のベクトル類似性検索用に構築されたオープンソースのベクトル データベース。
Zilliz Vector Lakebase BYOC を発表
価格設定 料金プラン あらゆる予算であらゆるチームに柔軟な価格設定オプションを提供
定価 すべての請求項目の詳細な定価を表示します
無料利用枠 マネージド Milvus の力で想像力を解き放ちます
Zilliz Cloud 開発者ハブでは、Zilliz Cloud を使用するためのすべての情報が見つかります。
リソース ブログ ガイド リサーチ アナリスト レポート ウェビナー トレーニング
顧客検索 拡張生成 すべてのユースケースを見る 業界別で見る すべての顧客事例を見る OpenEvidence が Zilliz Cloud で医療 AI を強化
English 日本語 한국어 Español Français Deutsch Italiano Português Русский 44.6k ログイン デモを予約 無料で始める
Vector データベースから Vector Lakebase へ
Vector データベースから Vector Lakebase にページをコピー
統合データ基盤と 3 つのワークロード モードが本当に重要なのはなぜですか?
Vector Lakebase の主要な機能
階層型リアルタイム サービス ソリューション
Vector Lakebase の主な使用例
GenAI アプリケーション用に構築されたフルマネージドのベクター データベースをお試しください。
本日、Zilliz Cloud の次の章となる Zilliz Vector Lakebase のパブリック プレビューを開始します。 Vector Lakebase は、ベクトル データベースを超えた次のステップです。これは、オープン ストレージとエラスティック コンピューティングが AI ワークロードのために統合される、セマンティック中心のデータ プラットフォームです。
Vector データベースは、リアルタイムの提供を目的として構築されています。
Vector Lakebase は、S3 ベースの統合データ基盤上に構築されており、次の 3 つのワークロード モードにわたって AI とエージェントを強化します。
待ち時間が重要な製品のリアルタイム取得

アクションサービング、
インタラクティブな複数ステップの探索のための反復的な発見、
オフライン マイニングとデータセットの最適化のためのバッチ分析。
ギガバイトからペタバイトまでのすべてのスケーリング。
統合データ基盤と 3 つのワークロード モードが本当に重要なのはなぜですか?
つまり、AI システムはもはや単なる単一クエリの検索問題ではないからです。これらは、サービス提供、学習、改善の継続的なループとして機能します。
この図が示すように、AI およびエージェント アプリケーションのデータ基盤は通常 3 つの部分で構成されます。下部の生のマルチモーダル データ、オンライン サービス用のセマンティック データ (テキスト、ベクター、ラベルなど)、および運用システムから収集されたフィードバック データ (ユーザーの行動、ログ、エージェントのメモ、統計など)。
多くの成熟したエージェント アプリケーションには、この種のデータ基盤がすでに備わっています。本当の問題点は、これらのさまざまな種類のデータが、ワークフロー ループをサポートするための統合され構造化されたデータ プレーンが存在せず、複数のパイプラインやシステムに分散していることがよくあることです。
オンライン サービス (濃い青) → 知識とフィードバックの蓄積 (水色とオレンジ) → インサイトの発見 (緑) → データセットと戦略の改善 (紫) → オンライン サービスの向上 。
図にも示されているように、ベクトル データベースは主にリアルタイムの取得と提供指向のデータ書き込み (2 つの青いパス) をサポートしているため、ベクトル データベースだけではもはや十分ではありません。このループでは、他の 2 つのアクセス モード (対話型検出とバッチ分析) も同様に重要です。
たとえば、AI 開発者は、サービスの品質が低い理由を理解するために、(手動またはエージェント システムを通じて) フィードバック データと基礎となるコーパスを調査する必要があることがよくあります。また、新しくクロールされたデータに対して大規模なセマンティック重複排除とクラスタリングを実行し、エッジ クラスターをマイニングして新しいトレーニング データ カンジダを発見することもあります。

テス。
これらのワークロードは、従来のビッグデータ処理とは大きく異なります。コアの計算は数値的ではなく意味論的です。データは主にベクトル、テキスト、ラベル、セマンティック メタデータで構成されますが、コア操作にはベクトル検索、全文検索、再ランキング、セマンティック クラスタリング、および関連するセマンティック検索タスクが含まれます。
このため、インタラクティブな検出とバッチ分析は、データ層とコンピューティング層の両方でベクトル データベースと自然に連携します。多くの場合、オンライン サービスとオフライン処理は、同じ基礎となるデータ基盤を共有します。
たとえば、チームは価値の高いユーザー タスクをオフラインでクラスタリングして分析すると同時に、サービス提供システム内のサポート知識や戦略に希薄性や品質の問題が見られないかどうかをチェックできます。
全体として、断片化したデータ アーキテクチャや孤立したインフラストラクチャはこのループの速度を低下させます。これは、急速に進化する AI 機能の競争において致命的となる可能性があります。 Vector Lakebase は、単純かつ効率的なアプローチを通じてこのループを高速化します。つまり、リアルタイム取得、対話型検出、バッチ分析という 3 つのワークロード モードすべてから効率的にアクセスできるゼロコピー セマンティック データ プレーンを提供します。
Vector Lakebase の主要な機能
Ziliz Vector Lakebase は、次の 5 つのコア機能を通じてこのワークフロー ループをサポートします。
階層型サービス提供ソリューション
さまざまなリアルタイム ワークロードに最適化された柔軟なサービス層により、大規模なデータセット全体で超高性能、バランスのとれた効率、コスト効率の高いスケーリングを実現します。
オンデマンド検索
待ち時間がそれほど重要ではなく、ほとんどの時間コンピューティングがアイドル状態のままである大規模なワークロード向けに設計されています (頻度の低い検索、データ探索、バッチ分析など)。
外部データレイク検索
最先端のインデックス作成と LA を追加

既存のレイク データに直接、スケール検索機能を追加します。
フルスペクトル検索
ベクターやテキストから JSON や地理空間まで、ハイブリッド検索、フィルタリング、再ランキングと組み合わせて表現力豊かなマルチモーダル クエリを実現します。
統合レイクネイティブストレージ
Vortex 上に構築された、サービス提供と分析の両方のための統合ストレージ。Lance や Parquet よりも高速かつ安価なランダム読み取りを提供するオープンな次世代フォーマットに加え、列ごとのフォーマットの柔軟性と広範なデータ モデリング機能を備えています。
階層型リアルタイム サービス ソリューション
Zilliz Cloud の階層型サービス ソリューションは、パフォーマンス最適化、容量最適化、階層型ストレージの 3 つのサービス階層を提供します。各層は専用のインデックス作成アルゴリズムとストレージ階層全体のデータ配置戦略を使用して構築されており、パフォーマンスとコストの幅広いトレードオフを提供します。
パフォーマンス最適化層は、超高パフォーマンスのシナリオをターゲットとしています。すべてのデータはメモリから直接提供され、1 桁ミリ秒の遅延で 1000 以上の QPS を実現します。マルチレプリカ展開によりスループットはさらに直線的に拡張されます。
容量最適化層は、メモリとローカル NVMe ストレージを組み合わせて、パフォーマンスと容量のバランスをとります。 100 ms 未満の遅延で 100 ～ 500 QPS を実現し、ほとんどの取得ワークロードに適しています。
階層型ストレージ層は、メモリ、ローカル NVMe、オブジェクト ストレージに及びます。高度に最適化されたプリフェッチおよびキャッシュ戦略により、データ アクセスの 95% 以上がメモリまたはローカル ディスクにアクセスし、大幅に低いインフラストラクチャ コストで約 100 ミリ秒の遅延で 10 ～ 50 QPS を実現します。
3 つの層はすべてデフォルトで 95% ～ 98% の再現率を実現し、インデックス作成と検索全体で柔軟に調整できるため、ワークロード要件に基づいて 90% ～ 99% 以上の再現率をサポートします。
これらのサービス アーキテクチャは、世界で最も要求の厳しいいくつかの環境で厳しいテストを受けています。

以下を含む大規模な AI およびインターネット ワークロード:
インターネット規模のマルチテナント AI プラットフォーム、
プレミアムエンタープライズユーザーと大規模な無料ユーザープールの両方に対する差別化されたサービスレベル、
高性能エージェントのナレッジベース、
超高スループットのレコメンデーション システム、
ストレージ層全体にわたる第 2 レベルの動的ホット/コールド データ スケジューリング、
極度のコスト制約の下で 1000 億を超える規模の自動運転データ マイニング パイプライン。
オンライン サービスの場合、Zilliz Cloud は、99.99% の稼働時間 SLA に裏付けられた、クロスリージョンの高可用性と災害復旧のためのグローバル クラスター機能も提供します。
インタラクティブな検出とバッチ分析は、特にフィードバック データ、エージェントが生成したメモ、ログ、クロールされたコーパスを含む場合、オンライン サービングよりも 1 ～ 3 桁大きいデータ量で動作することがよくあります。これらのデータセットは、TB または PB スケールに簡単に達する可能性があります。しかし、数百、さらには数千のベクトル データベース ノードを使用してサービスを提供することは、費用対効果の観点から正当化するのが難しいことがよくあります。
さらに重要なのは、これらのワークロードは通常、タスク主導型であるということです。エージェント アプリケーションのオンライン サービス層とは異なり、24 時間年中無休のアクティブなインフラストラクチャは必要ありません。コンピューティング リソースは、アクティブな処理タスク中にのみ頻繁に使用され、ほとんどの時間はアイドル状態にあり、多くの場合、アイドル時間は 97% 以上になります。
サーバーレス サービス ソリューションは魅力的に見えるかもしれませんが、これらのワークロードではコストが大幅に高くなることがよくあります。
コンピューティング層では、サーバーレス システムとオンデマンド検索の両方が従量課金制モデルに従います。詳細な価格モデルは異なりますが、基本的なコンピューティング コストは多くの場合同様です。ただし、サーバーレス アーキテクチャでは、プールのオーバーヘッド、インデックス作成、および永続データのコストが、トランザクションのパフォーマンスを直接反映するのではなく、追加の書き込みおよびストレージのマークアップに組み込まれます。

基礎となるリソースのコスト。
対照的に、Zilliz On-Demand Search は、オブジェクト ストレージとオンデマンド コンピューティングに対して直接料金を請求します。AWS Lambda と同様に、料金は主に割り当てられたリソース サイズと実行時間に基づいて決まりますが、ストレージ コストは基礎となる S3 コストに近いままです。これにより、隠れたインフラストラクチャのオーバーヘッドやブラックボックスの価格モデルが回避されます。
次の比較は、サーバーレス検索とオンデマンド検索のコストの違いを示しています。
768 次元の 1B ベクトル、データとインデックス ファイルを含む約 6 TB のストレージを必要とします。
1 か月の期間 (累積アクティブ コンピューティング時間は 10 時間)。
全体として、この実験では、オンデマンド検索の総コストはサーバーレスの総コストのわずか約 1/15 (318 ドル対 4,937 ドル) です。
Zilliz Vector Lakebase は、フルマネージドのストレージとクエリ コンピューティングを提供し、ユーザーがデータを Zilliz Cloud に直接保存して操作できるようにします。ただし、一部の顧客はすでに成熟したデータ レイク インフラストラクチャとガバナンス パイプラインを導入しています。
AI アプリケーションの場合、重要な課題の 1 つは、既存のレイク データに基づいて効率的な検索とセマンティック探索を直接可能にすることです。 Spark や Ray などの従来のビッグ データ システムは、基本的にインデックス高速化されたクエリやセマンティック検索ではなく、フルデータ スキャンとマップリデュース計算を中心に設計されているため、これらのワークロードには最適化されていません。
これを解決するために、Zilliz は外部収集モードを提供します。 Zilliz データ プレーンから顧客所有のレイク テーブルへのゼロコピー論理マッピングを作成し、そのマッピング上で高パフォーマンスのインデックスとフルスペクトル検索を可能にします。
現在、外部コレクションは 2 つのデータ レイク テーブル形式 (Lance と Iceberg)、および 2 つのオープン データ形式 (Parquet と Vortex) をサポートしています。
データレイクアップの場合

Zilliz 外部コレクションは増分同期機能を提供します。データ レイクの更新パターンとクエリの可視性要件に基づいて、ユーザーは更新呼び出しでいつでもデータを同期できます。
AI アプリケーションでは、さまざまなソースやモダリティにわたってデータを取得および分析する必要がますます高まっています。これは、検索と分析の品質を向上させるために、補完的な情報を組み合わせたり、同じ生のコンテンツから複数の視点を抽出したりするためです。
Zilliz Vector Lakebase は、密ベクトル、疎ベクトル、テキスト、JSON、地理空間データ、プリミティブ タイプを含む豊富なデータ タイプを含むワイド テーブル モデリングと、Struct や Array などの複雑な構造をサポートしており、統合されたテーブル レイアウト内で直接、効率的なネストされたセマンティック モデリングを可能にします。
これにより、各アプリケーション レベルのエンティティを 1 つの行に直接マッピングすることで、統合されたコンテキスト モデリングが可能になります。たとえば、Zilliz Vector Lakebase では、ドキュメントをテキスト チャンク、画像、表の数百行に分割する代わりに、ドキュメント全体を単一の行としてモデル化できます。これにより、JOIN と集計のパフォーマンスと操作上のオーバーヘッドを回避しながら、マルチモーダルな取得と分析が向上します。
データ モデリングを超えて、Vector Lakebase は最先端のインデックス作成機能と検索機能も提供します。

[切り捨てられた]

## Original Extract

Zilliz offers a fully managed Vector Lakebase powered by Milvus, unifying real-time vector search, lake-scale discovery, and Al
data operations.

From Vector Database to Vector Lakebase - Zilliz blog Products Zilliz Cloud
Fully managed Vector Lakebase services designed for high reliability, performance, and cost efficiency at scale.
Open-source vector database built for billion-scale vector similarity search.
Announcing Zilliz Vector Lakebase BYOC
Pricing Pricing Plan Flexible pricing options for every team on any budget
List Price View detailed list prices for every billing item
Free Tier Unleash Your Imagination with the Power of Managed Milvus
The Zilliz Cloud Developer Hub where you can find all the information to work with Zilliz Cloud
Resources Blog Guides Research Analyst Reports Webinars Trainings
Customers Retrieval Augmented Generation View all use cases View by industry View all customer stories OpenEvidence Powers Medical AI with Zilliz Cloud
English 日本語 한국어 Español Français Deutsch Italiano Português Русский 44.6k Log In Book a Demo Get Started Free
From Vector Database to Vector Lakebase
Copy page From Vector Database to Vector Lakebase
Why do the unified data foundation and three workload modes really matter?
The Key Vector Lakebase Features
Tiered Real-Time Serving Solutions
Primary Use Cases of Vector Lakebase
Try the fully-managed vector database built for your GenAI applications.
Today, we're launching the public preview of Zilliz Vector Lakebase — the next chapter for Zilliz Cloud. Vector Lakebase is the next step beyond vector databases. It is a semantic-centric data platform where open storage and elastic compute converge for AI workloads.
Vector Databases are purpose-built for real-time serving.
Vector Lakebase builds on an S3-based unified data foundation to power AI and agents across three workload modes:
real-time retrieval for latency-critical production serving,
iterative discovery for interactive and multi-step exploration,
batch analytics for offline mining and dataset optimization.
All scaling from gigabytes to petabytes.
Why do the unified data foundation and three workload modes really matter?
In short: because AI systems are no longer just a single-query retrieval problem. They operate as a continuous loop of serving, learning, and improving.
As this figure shows, the data foundation for AI and agent applications usually has three parts: raw multimodal data at the bottom, semantic data for online serving (such as text, vectors, and labels), and feedback data collected from production systems (such as user behavior, logs, agent notes, and statistics).
Many mature agent applications already have this kind of data foundation. The real pain point is that these different types of data are often scattered across multiple pipelines and systems, without a unified and structured data plane to support the workflow loop:
online serving (dark blue) → knowledge and feedback accumulation (light blue and orange) → insight discovery (green) → dataset and strategy improvement (purple) → better online serving .
As the picture also shows, a vector database alone is no longer enough, because it mainly supports real-time retrieval and serving-oriented data writes (the two blue paths). In this loop, the other two access modes — interactive discovery and batch analytics — are just as important.
For example, AI developers (either manually or through agentic systems) often need to explore feedback data and the underlying corpus to understand why serving quality is poor. They may also run large-scale semantic deduplication and clustering on newly crawled data, then mine edge clusters to discover new training data candidates.
These workloads are very different from traditional big data processing. The core computation is semantic rather than numerical. The data mainly consists of vectors, text, labels, and semantic metadata, while the core operations include vector search, full-text search, reranking, semantic clustering, and related semantic retrieval tasks.
Because of this, interactive discovery and batch analytics are naturally aligned with vector databases at both the data and compute layers. In many cases, online serving and offline processing even share the same underlying data foundation.
For example, teams may cluster and analyze high-value user tasks offline while simultaneously checking whether the supporting knowledge or strategies in the serving system show sparsity or quality issues.
Overall, any fragmented data architecture or isolated infrastructure islands slow down this loop — which can be fatal in the rapidly evolving race for AI capabilities. Vector Lakebase accelerates this loop through a straightforward but efficient approach: providing a zero-copy semantic data plane that can be efficiently accessed by all three workload modes — real-time retrieval, interactive discovery, and batch analytics.
The Key Vector Lakebase Features
Zilliz Vector Lakebase supports this workflow loop through five core capabilities:
Tiered Serving Solutions
Flexible serving tiers optimized for different real-time workloads — delivering ultra-high performance, balanced efficiency, and cost-effective scaling across massive datasets.
On-Demand Search
Designed for large-scale workloads where latency is less critical and compute remains idle most of the time — including infrequent search, data exploration, and batch analytics.
External Data Lake Search
Add state-of-the-art indexing and large-scale search capabilities directly to your existing lake data.
Full-Spectrum Search
From vector and text to JSON and geospatial—combined with hybrid retrieval, filtering, and reranking for expressive multi-modal queries.
Unified Lake-Native Storage
Unified storage for both serving and analytics, built on Vortex — an open next-generation format providing faster and cheaper random reads than Lance and Parquet, plus per-column format flexibility and broader data modeling capabilities.
Tiered Real-Time Serving Solutions
Zilliz Cloud's Tiered Serving Solutions provide three serving tiers: Performance-Optimized, Capacity-Optimized, and Tiered-Storage. Each tier is built with dedicated indexing algorithms and data placement strategies across the storage hierarchy, offering a wide range of performance–cost tradeoffs.
The Performance-Optimized tier targets ultra-high-performance scenarios. All data is served directly from memory, delivering 1000+ QPS with single-digit millisecond latency. Throughput further scales linearly with multi-replica deployment.
The Capacity-Optimized tier combines memory and local NVMe storage to balance performance and capacity. It delivers 100~500 QPS with sub-100 ms latency, making it suitable for most retrieval workloads.
The Tiered-Storage tier spans memory, local NVMe, and object storage. With highly optimized prefetching and caching strategies, over 95% of data access still hits memory or local disk, providing 10~50 QPS with around 100 ms latency at significantly lower infrastructure cost.
All three tiers deliver 95%–98% recall by default, with flexible tuning across indexing and search—supporting 90% to 99%+ recall based on workload requirements.
These serving architectures are battle-tested in some of the world’s most demanding large-scale AI and internet workloads, including:
internet-scale multi-tenant AI platforms,
differentiated service tiers for both premium enterprise users and large-scale free-user pools,
high-performance agent knowledge bases,
ultra-high-throughput recommendation systems,
second-level dynamic hot/cold data scheduling across storage tiers,
autonomous driving data mining pipelines at 100B+ scale under extreme cost constraints.
For online serving, Zilliz Cloud also provides Global Cluster capabilities for cross-region high availability and disaster recovery, backed by a 99.99% uptime SLA .
Interactive discovery and batch analytics often operate on data volumes one to three orders of magnitude larger than online serving, especially when including feedback data, agent-generated notes, logs, and crawled corpora. These datasets can easily reach TB or even PB scale. But using hundreds or even thousands of vector database nodes to serve them is often hard to justify from a cost–benefit perspective.
More importantly, these workloads are usually task-driven. Unlike the online serving layer of agent applications, they do not require 24/7 active infrastructure. Compute resources are heavily used only during active processing tasks, while remaining idle most of the time, often with over 97% idle time.
Serverless serving solutions may seem appealing, but they often become much more expensive for these workloads.
At the compute layer, both serverless systems and On-Demand Search follow a pay-as-you-go model. Despite differences in detailed pricing models, the underlying compute cost is often similar. However, in serverless architecture, pooling overhead, indexing, and persistent data costs are embedded into additional write and storage markups, rather than directly reflecting the true cost of underlying resources.
In contrast, Zilliz On-Demand Search charges directly for object storage and on-demand compute — similar to AWS Lambda, where pricing is primarily based on allocated resource size and execution time, while storage cost remains close to the underlying S3 cost. This avoids hidden infrastructure overhead and black-box pricing models.
The following comparison illustrates the cost difference between Serverless and On-Demand Search.
1B vectors with 768 dimensions, requiring approximately 6 TB of storage including data and index files,
1 month duration with 10 hours of accumulated active compute time.
Overall, in this experiment, the total cost of On-Demand Search is only about 1/15 ($318 vs $4,937) that of Serverless.
Zilliz Vector Lakebase provides fully managed storage and query compute, allowing users to store and operate their data directly in Zilliz Cloud. However, some customers already have mature data lake infrastructure and governance pipelines in place.
For AI applications, one of the key challenges is enabling efficient retrieval and semantic exploration directly on top of existing lake data. Traditional big data systems such as Spark and Ray are not optimized for these workloads, because they are fundamentally designed around full-data scan and map-reduce computation rather than index-accelerated query and semantic retrieval.
To solve this, Zilliz provides an External Collection mode. It creates a zero-copy logical mapping from the Zilliz data plane to customer-owned lake tables, while enabling high-performance indexes and full-spectrum search on top of that mapping.
Currently, External Collection supports two data lake table formats — Lance and Iceberg, as well as two open data formats — Parquet and Vortex.
For data lake updates, Zilliz External Collection provides incremental synchronization capabilities. Based on the data lake update pattern and query visibility requirements, users can sync data anytime with a refresh call.
AI applications increasingly need to retrieve and analyze data across different sources and modalities — both to combine complementary information and to extract multiple perspectives from the same raw content for better retrieval and analysis quality.
Zilliz Vector Lakebase supports wide-table modeling with rich data types including dense and sparse vectors, text, JSON, geospatial data, and primitive types, along with complex structures such as Struct and Array — enabling efficient nested semantic modeling directly within a unified table layout.
This enables unified context modeling by mapping each application-level entity directly to a single row. For example, instead of splitting a document into hundreds of rows for text chunks, images, and tables, Zilliz Vector Lakebase can model the entire document as a single row. This improves multi-modal retrieval and analytics while avoiding the performance and operational overhead of JOINs and aggregations.
Beyond data modeling, Vector Lakebase also provides state-of-the-art indexing and search capabilities acros

[truncated]
