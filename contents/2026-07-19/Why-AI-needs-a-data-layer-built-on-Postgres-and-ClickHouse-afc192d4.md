---
source: "https://clickhouse.com/blog/ai-best-of-breed-data-stack"
hn_url: "https://news.ycombinator.com/item?id=48968694"
title: "Why AI needs a data layer built on Postgres and ClickHouse"
article_title: "AI needs the best-of-breed data stack: Postgres and ClickHouse | ClickHouse"
author: "saisrirampur"
captured_at: "2026-07-19T15:47:19Z"
capture_tool: "hn-digest"
hn_id: 48968694
score: 1
comments: 0
posted_at: "2026-07-19T14:50:28Z"
tags:
  - hacker-news
  - translated
---

# Why AI needs a data layer built on Postgres and ClickHouse

- HN: [48968694](https://news.ycombinator.com/item?id=48968694)
- Source: [clickhouse.com](https://clickhouse.com/blog/ai-best-of-breed-data-stack)
- Score: 1
- Comments: 0
- Posted: 2026-07-19T14:50:28Z

## Translation

タイトル: なぜ AI には Postgres と ClickHouse 上に構築されたデータ レイヤーが必要なのか
記事のタイトル: AI には最高のデータ スタックが必要: Postgres と ClickHouse |クリックハウス
説明: AI アプリケーションは爆発的なデータの増加を促進し、トランザクション ワークロードと分析ワークロードの間の境界線を崩壊させ、Postgres と ClickHouse を最適に組み合わせることで、CDC、ネイティブ クエリ、監視を通じて連携がますます高まっています。

記事本文:
。 -->
AI には最高のデータ スタックが必要です: Postgres と ClickHouse | ClickHouse コンテンツにスキップ ロゴを SVG としてコピー 完全なロゴをダウンロード ロゴマークをダウンロード 検索を開く 地域セレクターを開く 英語
製品 製品 ClickHouse Cloud ClickHouse の最適な使い方。
AWS、GCP、Azure で利用できます。
Bring Your Own Cloud フルマネージドの ClickHouse サービス、
自分の AWS および GCP アカウントにデプロイされます。
ClickHouse によって管理される Postgres トランザクション用の統合データ スタック
そして分析。
マネージド ClickStack 高パフォーマンスによるマネージドオブザーバビリティ
クエリと長期保存。
Langfuse Cloud LLM の可観測性と評価
信頼性の高い AI アプリケーションとエージェントを実現します。
オープンソース ClickHouse 高速オープンソース OLAP データベース
リアルタイム分析。
ClickStack ログ用のオープンソース可観測性スタック、
メトリクス、トレース、セッションのリプレイ。
Agentic Data Stack AI を活用したアプリケーションを構築
クリックハウスで。
chDB インプロセス SQL エンジンを搭載
ClickHouse、Pandas 互換 API を使用
ソリューション ユースケース リアルタイム分析
リソース 会社のリソース ユーザーストーリー
ページをコピーしました コピーされました!その他のアクション マークダウンで表示 このページをマークダウンで開く
ChatGPT で開く このページについて質問する
クロードで開く このページについて質問する
v0 で開く このページについて質問する
AI には最高のデータ スタックが必要です: Postgres と ClickHouse
過去数年間、Postgres や ClickHouse を使用する何千もの企業と協力して、私は AI がアプリケーションがデータベースに要求するものを変えるのを見てきました。すべての AI アプリケーションはデータ アプリケーションです。エージェントは膨​​大な量の運用データを生成し、ユーザーはそのデータに基づいて構築されたリアルタイムの回答を、多くの場合同じリクエスト内で期待します。これにより、アプリケーションがデータベースに要求するものが変わります。
この投稿では、AI アプリケーションに実際にどのようなものが必要なのかを探っていきたいと思います。

Postgres と ClickHouse がこれらのワークロードにとって自然な組み合わせになりつつある理由と、この組み合わせを開発者が非常に利用しやすいものにするための当社の投資について、データ レイヤーから説明します。
TL;DR AI は、前例のない規模でデータの増加を加速しており、トランザクション (OLTP) ワークロードと分析 (OLAP) ワークロードの間の従来の分断を崩壊させています。これらの要求を満たすには、最高のデータベース、Postgres と ClickHouse が必要であり、それぞれが設計された目的に優れ、同時にシームレスに連携します。
AI に必要なもの、そして AI がデータ環境をどのように変えるのか #
AI がデータ層を新たな限界に押し上げています #
AI ネイティブ アプリケーションは、従来のソフトウェアよりもはるかに多くのデータを生成します。すべてのプロンプト、応答、ツール呼び出し、評価、およびユーザー インタラクションがデータになり、データ量とクエリの同時実行性の両方が爆発的に増加します。
これは私が話しているスケールを示すデータポイントです。 Postgres と ClickHouse を使用している AI ネイティブ企業のサンプル全体で、6 か月間で平均 1,000% 以上のデータ増加が観察され、85 TB 以上のデータが追加されました。
また、AI がバイブコーディングされたアプリを超えて実稼働システムに移行するにつれて、トレードオフも変化しています。アプリケーションはミッションクリティカルな意思決定を行い、より広い表面積にわたってより大規模に動作するようになっており、データ層の信頼性とセキュリティの基準が引き上げられています。
これらの需要に対応するために、AI には、大量のデータをリアルタイムで取り込み、処理し、クエリを実行でき、常に利用可能で安全なデータ レイヤーが必要です。高速プロビジョニングと即時分岐は便利ですが、基盤となるデータ スタックが高速、信頼性、安全でない場合はほとんど意味がありません。
AI にはリアルタイムで連携する OLTP と OLAP が必要です #
LLM を利用したアプリケーション、AI が生成した洞察、異常検出、推奨事項

ngine と自然言語インターフェイスはすべて、トランザクション データベースと分析データベースの間でより緊密なフィードバック ループを必要とします。トランザクション データベースに書き込まれたデータは、不正の可能性のあるトランザクションを評価する場合でも、本番環境のインシデントを診断する場合でも、ほぼ即座に分析クエリに使用できるようにする必要があります。
このようなシナリオでは、1 時間または 1 分ごとの ETL を通じてデータをウェアハウスに移動する従来のパターンでは十分ではなくなります。トランザクション データベースと分析データベースは、データが書き込まれてから分析に使用できるようになるまでの間隔はわずか数秒で、リアルタイムで動作する必要があります。
オープンソースの基盤とコストは重要 #
AI インフラストラクチャでも同様の変化が進行中です。チームは独自の LLM ロックインとトークンコストの増大に対する警戒を強めるにつれ、データスタックの柔軟性と経済性についても再考しています。オープンソース データベースは両方を解決します。セルフホスト型または管理型の展開、ニーズの進化に応じてベンダーを自由に切り替えることができ、多くの場合コストが大幅に削減されます。
Postgres + ClickHouse が勝てる理由 #
上記の要件により、AI ネイティブ企業が選択するデータ スタックが再構築されています。 Postgres と ClickHouse は、強力なオープンソース基盤上でクラス最高のリアルタイム トランザクションと分析を組み合わせた、デフォルトの選択肢となりつつあります。
なぜこれが起こっているのかを理解しましょう。
Postgres と ClickHouse は最高のオープンソース データベースです #
一方では Postgres は、世界で最も人気のあるオープンソースのトランザクション データベースです。低レイテンシの CRUD 操作、ACID コンプライアンス、豊富なインデックス作成、事実上あらゆるクエリをサポートする完全な SQL インターフェイス、堅牢な拡張フレームワーク、広大な ORM とアプリケーション エコシステムなど、トランザクションに対する堅実なサポートを提供し、当社のデフォルトのバックエンドとなっています。

b と AI アプリケーションがどこにでもあります。
一方、 ClickHouse は地球上で最も高速な分析データベースです。列指向ストレージ、スキップ インデックス、増分マテリアライズド ビュー、分散キャッシュ、ネイティブ JSON と全文検索、最近導入された遅延マテリアライゼーション、および超高速リアルタイム分析に重点を置いた数百もの最適化などの専用機能を提供します。
どちらもエンタープライズ対応であり、高可用性、信頼性、セキュリティのための実証済みの機能を提供します。
さらに、彼らは同じオープンソース哲学を共有しており、大規模で活発なコミュニティに支えられています。面白い事実として、Postgres は最近オープンソース データベースとして 30 周年を祝いましたが、ClickHouse は 10 周年を迎えました。
GitLab の共同創設者によるブログ。Postgres と ClickHouse がどのように相互補完し、ほとんどのデータ問題を解決するかについて説明しています。
Postgres は、業界で最も豊富なツールと統合のエコシステムの 1 つを提供します。一方、ClickHouse は、使い慣れた SQL、単一バイナリによる簡単なローカル開発、および大規模なチューニングをほとんど必要としない分析パフォーマンスを組み合わせています。これらは共に、「とにかく動く」という、すっきりシンプルに感じられる開発エクスペリエンスを提供します。これは、数万のスタートアップ企業や AI ネイティブ企業から、GitLab 、Instacart 、Cloudflare などの企業に至るまで、その広範な採用に反映されています。
DB-Engines では、Postgres と ClickHouse がそれぞれ最も人気のあるオープンソース OLTP データベースと OLAP データベースとして示されており、どちらも人気が高まり続けています。
Postgres および ClickHouse を使用した OLTP と OLAP の統合 #
ここでの私たちのアプローチはシンプルです。Postgres と ClickHouse を OLTP と OLAP の主要なオープンソース データベースとして採用し、それらの間の統合を開発者にとって可能な限りシームレスかつネイティブにすることです。
私たちはスタック全体に投資してきました。

これが起こります。データの移動には、Postgres 用 ClickPipes を強化し、1 桁秒のレプリケーションを備えた Change Data Capture (CDC) を提供する PeerDB があります。マルチテラバイトの Postgres デプロイメントを実行している 1,000 社以上の企業によって厳しいテストが行​​われています。並列スナップショット、自動スキーマ進化、ネイティブ Postgres および ClickHouse モニタリング、復元力のあるレプリケーションなどの専用機能が実装されており、顧客にとってエンタープライズ規模の CDC はほとんど魔法のように感じられます。
クエリに関しては、アプリケーションが Postgres から ClickHouse に直接クエリできるようにする pg_clickhouse 拡張機能があり、Postgres をトランザクション ワークロードと分析ワークロードの両方に対する単一のインターフェイスにし、分析クエリを透過的に ClickHouse にプッシュします。
可観測性のために、pg_stat_ch は Postgres クエリ テレメトリを ClickHouse にストリーミングすることでスタックを補完し、低オーバーヘッドで長期的な可観測性とパフォーマンス分析を可能にします。テレメトリは Postgres 自体ではなく ClickHouse に保存されるため、OLTP のパフォーマンスに影響を与えることなく、数か月分の実行統計を保持し、高カーディナリティ分析を実行できます。
また、ClickHouse によって管理される Postgres も開始しました。これは、Postgres とローカル NVMe ストレージを組み合わせて、業界をリードする OLTP パフォーマンスと CDC および pg_clickhouse によるネイティブ ClickHouse 統合を実現するフルマネージド Postgres サービスです。高可用性、PITR、リードレプリカ、フォーク、エンドツーエンド暗号化が組み込まれており、エンタープライズ対応の統合データ スタックを開発者に提供します。
他にも、Postgres と ClickHouse の統合を深めながら、それぞれが最も得意とすることを実行できるようにする取り組みが他にも複数あります。
以下の図は、OLTP 用の Postgres、OLAP 用の ClickHouse、両者間の継続的レプリケーション、および unif を強調した統合データ スタックを示しています。

pg_clickhouse 拡張機能を介した ied クエリ レイヤー。
データベース業界では、過去 10 年間にわたり、OLTP と OLAP を単一のデータベース エンジンに統合する試みが何度も行われてきました。最近では、ユニファイド ストレージとデータ レイクを中心に構築されたアーキテクチャも見られます。
これらのアプローチは技術的に興味深いものですが、同時に開発者を独自のプラットフォームに結びつけ、根本的に異なる要件を持つワークロードを最適化することを単一のシステムに要求します。 OLTP と OLAP には理由があります。行指向と列指向のストレージはそれぞれ、さまざまなアクセス パターンとパフォーマンス特性に合わせて最適化されています。
同様に、データ レイクはデータ ウェアハウジングのユースケースに優れていますが、リアルタイムの顧客およびエージェント向けアプリケーションの低遅延、高同時実行性の要求向けに設計されていません。これらのアーキテクチャの上に運用機能を重ねることはできますが、オープン テーブル形式と汎用ストレージ レイヤーをサポートする必要があるという制約が残ります。その結果、必然的にトレードオフが生じ、リアルタイム分析専用に構築された MergeTree のようなストレージ エンジンのパフォーマンスに匹敵することが困難になります。
この違いは、ClickHouse の戦略、つまり最善のアプローチも反映しています。私たちは、MergeTree をリアルタイム分析に最適なエンジンにすることに重点を置きながら、データ ウェアハウジングのユースケース向けのオープン テーブル形式を中心に継続的に革新を行っています。
AI の未来は最高のデータスタックです #
AI によって OLTP と OLAP の区別がなくなるのではなく、その重要性がこれまで以上に高まっています。この時代を定義するアプリケーションには、すべてのインタラクションを確実にキャプチャできるトランザクション データベース、そのデータをリアルタイムで処理できる分析データベース、およびその 2 つの間のシームレスな統合が必要です。
私たちは未来は単一のものではないと信じています

根本的に異なるワークロードに合わせて最適化しようとしているプラ​​ットフォーム。これは、Postgres と ClickHouse 上に構築されたオープンで最高のデータ スタックであり、各データベースが 1 つとして連携しながら最善のことを実行します。
それが私たちが目指しているビジョンです。 PeerDB と ClickPipes を備えたエンタープライズ グレードの CDC から、 pg_clickhouse 、 pg_stat_ch 、 ClickHouse によって管理される Postgres に至るまで、すべての投資は、Postgres と ClickHouse が開発者にとって単一の統合スタックのように感じられるようにすることを目的としています。
ClickHouse が管理する Postgres を試してみる
ClickHouse + Postgres は、拡張可能なアプリケーション向けの統合データ スタックになりました。 ClickHouse Cloud でマネージド Postgres が利用できるようになったことで、このスタックは初日の決定で済みます。
ClickHouse がデータに対してどのように作用するか知りたいですか? ClickHouse Cloud を数分で始めて、300 ドルの無料クレジットを受け取りましょう。
機能リリース、製品ロードマップ、サポート、クラウド製品に関する最新情報を入手してください。
Trainy が Amazon RDS Postgres から ClickHouse マネージド Postgres に移行した理由
HDB の置き換え: ヒストリカル ティッカー データの ClickHouse
Bullet が ClickHouse Cloud を使用して DeFi の最速の取引所リアルタイム分析を提供する方法
ClickHouse によって管理される Postgres
エージェントによる可観測性のための ClickStack
最新情報を入手

[切り捨てられた]

## Original Extract

AI applications are driving explosive data growth and collapsing the line between transactional and analytical workloads, making the best-of-breed combination of Postgres and ClickHouse — increasingly tied together through CDC, native querying, and observ

. -->
AI needs the best-of-breed data stack: Postgres and ClickHouse | ClickHouse Skip to content Copy logo as SVG Download full logo Download logomark Open search Open region selector English
Products Products ClickHouse Cloud The best way to use ClickHouse.
Available on AWS, GCP, and Azure.
Bring Your Own Cloud A fully managed ClickHouse service,
deployed in your own AWS and GCP account.
Postgres managed by ClickHouse Unified data stack for transactions
and analytics.
Managed ClickStack Managed observability with high-performance
queries and long-term retention.
Langfuse Cloud LLM observability and evaluations
for reliable AI applications and agents.
Open source ClickHouse Fast open-source OLAP database for
real-time analytics.
ClickStack Open-source observability stack for logs,
metrics, traces, and session replays.
Agentic Data Stack Build AI-powered applications
with ClickHouse.
chDB In-process SQL Engine powered by
ClickHouse, with a Pandas-compatible API
Solutions Use cases Real-time analytics
Resources Company resources User stories
Copy page Copied! More actions View as Markdown Open this page in Markdown
Open in ChatGPT Ask questions about this page
Open in Claude Ask questions about this page
Open in v0 Ask questions about this page
AI needs the best-of-breed data stack: Postgres and ClickHouse
Over the past few years working with thousands of companies on Postgres and ClickHouse , I've watched AI change what applications demand from their databases. Every AI application is a data application: agents generate huge volumes of operational data, and users expect real-time answers built on that data, often in the same request. That changes what applications demand from their databases.
In this post, I want to explore what AI applications actually need from their data layer, why Postgres and ClickHouse are increasingly becoming the natural combination for those workloads and our investments in making this combination very accessible to developers.
TL;DR AI is accelerating data growth at an unprecedented scale and is collapsing the traditional divide between transactional (OLTP) and analytical (OLAP) workloads . Meeting these demands requires best-of-breed databases, Postgres and ClickHouse, each excelling at what it was designed to do, while still working together seamlessly.
What AI needs, and how it's changing the data landscape #
AI is pushing the data layer to new limits #
AI-native applications generate far more data than traditional software. Every prompt, response, tool call, evaluation, and user interaction becomes data, driving explosive growth in both data volume and query concurrency.
Here is a data-point showing the scale I am talking about. Across a sample of AI-native companies using Postgres and ClickHouse, we observed an average data increase of more than 1,000% over six months , adding over 85 TB of data .
Also, as AI moves beyond vibe-coded apps into production systems, the tradeoffs are changing. Applications are making mission-critical decisions and operating across a larger surface area and at greater scale, raising the bar for reliability and security at the data layer.
To keep up with these demands, AI needs a data layer that can ingest, process, and query massive volumes of data in real time, always available and secure. Fast provisioning and instant branching are nice, but they matter little if the underlying data stack isn’t fast, reliable or secure.
AI needs OLTP and OLAP working together in real-time #
LLM-powered applications, AI-generated insights, anomaly detection, recommendation engines, and natural language interfaces all demand a much tighter feedback loop between transactional and analytical databases. Data written to a transactional database must become available for analytical queries almost immediately, whether to evaluate a potentially fraudulent transaction or diagnosing a production incident.
In these scenarios, the traditional pattern of moving data into a warehouse through hourly or minutely ETL is no longer sufficient. Transactional and analytical databases must operate in real time, with only seconds separating when data is written from when it is available for analysis.
Open source foundations and costs matter #
A similar shift is underway in AI infrastructure. As teams grow more wary of proprietary LLM lock-in and mounting token costs, they're also rethinking the flexibility and economics of their data stack. Open-source databases solve both: they offer self-hosted or managed deployment, the freedom to switch vendors as needs evolve, and often significantly lower costs.
Why Postgres + ClickHouse is winning #
The above requirements are reshaping the data stacks AI-native companies choose. Postgres and ClickHouse are increasingly becoming the default choice, combining best-in-class real-time transactions and analytics on a strong open-source foundation.
Let's understand on why this is happening.
Postgres and ClickHouse are best-of-breed open source databases #
Postgres , on the one hand, is the most popular open-source transactional database in the world. It offers rock-solid support for transactions, including low-latency CRUD operations, ACID compliance, rich indexing, a full SQL interface that supports virtually any query, a robust extension framework, and a vast ORM and application ecosystem, making it the default backend for web and AI applications everywhere.
ClickHouse , on the other hand, is the fastest analytical database on the planet. It offers purpose-built capabilities including columnar storage, skip indexes, incremental materialized views, distributed cache, native JSON and full-text search, the recently introduced lazy materialization, and hundreds of optimizations, laser-focused on blazing-fast real-time analytics.
Both are enterprise-ready, offering proven capabilities for high availability, reliability, and security.
Beyond that, they share the same open-source philosophy, backed by large and active communities. As a fun fact, Postgres recently celebrated its 30th anniversary as an Open Source database, while ClickHouse marked its 10th.
Blog by cofounder of GitLab on how Postgres and ClickHouse complement each other and solve most data problems.
Postgres offers one of the richest ecosystems of tools and integrations in the industry, while ClickHouse combines familiar SQL, effortless local development with a single binary, and analytical performance that rarely requires extensive tuning. Together, they deliver a dev experience that feels refreshingly simple: “it just works" This is reflected in their vast adoption, from tens of thousands of startups and AI-native companies to enterprises like GitLab , Instacart , and Cloudflare .
DB-Engines shows Postgres and ClickHouse as the most popular open source OLTP and OLAP databases, respectively, with both continuing to grow in popularity.
Unifying OLTP and OLAP with Postgres and ClickHouse #
Our approach here is simple: embrace Postgres and ClickHouse as the leading open source databases for OLTP and OLAP, and make the integration between them as seamless and native as possible for devs.
We've been investing across the stack to make this happen. For data movement, we have PeerDB , which powers ClickPipes for Postgres and provides Change Data Capture (CDC) with single-digit-second replication. It's battle-tested by more than 1,000 companies running multi-terabyte Postgres deployments. It implements purpose-built capabilities like parallel snapshotting , automatic schema evolution, native Postgres and ClickHouse monitoring and resilient replication, making enterprise-scale CDC feel almost magical for customers.
For querying, we have the pg_clickhouse extension that lets applications query ClickHouse directly from Postgres, making Postgres the single interface for both transactional and analytical workloads while transparently pushing analytical queries to ClickHouse.
For observability, pg_stat_ch complements the stack by streaming Postgres query telemetry into ClickHouse, enabling low-overhead, long-term observability and performance analysis. Because telemetry is stored in ClickHouse instead of Postgres itself, you can retain months of execution statistics and perform high-cardinality analysis without impacting OLTP performance.
We also launched Postgres managed by ClickHouse , a fully managed Postgres service that combines Postgres with local NVMe storage for industry-leading OLTP performance and native ClickHouse integration through CDC and pg_clickhouse. High availability, PITR, read replicas, forks, and end-to-end encryption come built in, giving developers a unified data stack that is enterprise-ready .
There are multiple other initiatives that deepen the integration of Postgres and ClickHouse while continuing to let each do what it does best!
The diagram below depicts the Unified Data Stack, highlighting Postgres for OLTP, ClickHouse for OLAP, continuous replication between the two, and a unified query layer via the pg_clickhouse extension.
The database industry has seen multiple attempts over the past decade to unify OLTP and OLAP into a single database engine. More recently, we've also seen architectures built around unified storage and data lakes.
These approaches are technically interesting, but they also tie developers to proprietary platforms and ask a single system to optimize for workloads with fundamentally different requirements. OLTP and OLAP exist for a reason: row and column-oriented storage are each optimized for different access patterns and performance characteristics.
Likewise, data lakes excel at data warehousing use-cases, and weren't designed for the low-latency, high-concurrency demands of real-time, customer and agent-facing applications. You can layer operational capabilities on top of these architectures, but they remain constrained by the need to support open table formats and a general-purpose storage layer. As a result, they inevitably involve tradeoffs that make it difficult to match the performance of storage engines like MergeTree, which are purpose-built for real-time analytics.
This distinction also reflects our strategy at ClickHouse: a best-of-breed approach. We focus on making MergeTree the best engine for real-time analytics while continuously innovating around open table formats for the data warehousing use-case.
The future of AI is a best-of-breed data stack #
AI isn't making the distinction between OLTP and OLAP disappear, it's making it more important than ever. The applications defining this era need transactional databases that can reliably capture every interaction, analytical databases that can process that data in real time, and seamless integration between the two.
We believe the future isn't a single proprietary platform trying to optimize for fundamentally different workloads. It's an open, best-of-breed data stack built on Postgres and ClickHouse, where each database does what it does best while working together as one.
That's the vision we're building toward. From enterprise-grade CDC with PeerDB and ClickPipes, to pg_clickhouse , pg_stat_ch , and Postgres managed by ClickHouse , every investment is aimed at making Postgres and ClickHouse feel like a single, integrated stack for developers.
Try Postgres managed by ClickHouse
ClickHouse + Postgres has become the unified data stack for applications that scale. With Managed Postgres now available in ClickHouse Cloud, this stack is a day-1 decision.
Interested in seeing how ClickHouse works on your data? Get started with ClickHouse Cloud in minutes and receive $300 in free credits.
Stay informed on feature releases, product roadmap, support, and cloud offerings!
Why Trainy migrated from Amazon RDS Postgres to ClickHouse Managed Postgres
Replacing the HDB: ClickHouse for historical ticker data
How Bullet uses ClickHouse Cloud to give DeFi's fastest exchange real-time analytics
Postgres managed by ClickHouse
ClickStack for agentic observability
Stay informed

[truncated]
