---
source: "https://product.hubspot.com/blog/building-the-ai-retrieval-infrastructure-behind-20-billion-vectors-at-hubspot"
hn_url: "https://news.ycombinator.com/item?id=48845387"
title: "Building the AI Retrieval Infrastructure Behind 20B+ Vectors at HubSpot"
article_title: "Building the AI Retrieval Infrastructure Behind 20 Billion+ Vectors at HubSpot"
author: "cyndunlop"
captured_at: "2026-07-09T13:43:30Z"
capture_tool: "hn-digest"
hn_id: 48845387
score: 3
comments: 0
posted_at: "2026-07-09T13:20:47Z"
tags:
  - hacker-news
  - translated
---

# Building the AI Retrieval Infrastructure Behind 20B+ Vectors at HubSpot

- HN: [48845387](https://news.ycombinator.com/item?id=48845387)
- Source: [product.hubspot.com](https://product.hubspot.com/blog/building-the-ai-retrieval-infrastructure-behind-20-billion-vectors-at-hubspot)
- Score: 3
- Comments: 0
- Posted: 2026-07-09T13:20:47Z

## Translation

タイトル: HubSpot で 200 億以上のベクトルの背後にある AI 検索インフラストラクチャを構築
記事のタイトル: HubSpot で 200 億以上のベクトルの背後にある AI 検索インフラストラクチャを構築する
説明: HubSpot がセマンティック検索を強化し、多様なアプリケーションをサポートするために、Qdrant で 200 億を超えるベクトルを管理するスケーラブルな AI 検索インフラストラクチャをどのように構築したかをご覧ください。

記事本文:
HubSpot で 200 億以上のベクトルの背後にある AI 検索インフラストラクチャを構築
セマンティック検索は HubSpot の運用において重要な役割を果たし、エージェントや RAG (検索拡張生成) から連絡先の重複排除に至るまで、さまざまなユースケースを強化します。数年前の POC から始まり、38 以上のチームが利用する数百億のベクターの規模にまで成長しました。エージェントの使用量が増加するにつれて、エージェントが検索品質を維持しながら必要な情報を迅速に取得できるようにするためのセマンティック検索機能の重要性がさらに高まっています。この記事では、小規模な POC から、HubSpot 全体の何百ものユースケースを支える数百億のエントリに至るまで、ベクター DB インフラストラクチャの歩みについて説明します。
VaaS (Vector as a Service) は、ベクトル データベース上に構築された HubSpot の集中型ベクトル ストレージおよび検索プラットフォームです。 API レイヤーと次のような機能を提供します。
コレクション設定に基づいたエンベディングの生成
これは、ベクトル データベース (この場合は Qdrant ) の前に配置され、消費者に直接公開されます。アクセス制御、基本的な検証、埋め込みの生成を処理し、Qdrant へのリクエストをプロキシします。読み取り操作の場合、リクエストは同期的に処理されますが、書き込み操作の場合、フローは非同期であり、コンシューマーのラグは最小限に抑えられます。
VaaS プラットフォームの高レベルのアーキテクチャ
HubSpot は、すべてのセマンティック検索操作に Qdrant ベクトル データベースを使用します。 Qdrant は、オンプレミスに展開できるオープンソースのベクトル データベースであり、名前付きベクトル、ハイブリッド検索、多段階クエリ、重み付け再ランキングなど、最新のセマンティック検索機能をカバーする広範な API を提供します。 HNSW (Hierarchical Navigable Small World) を ANN (近似最近傍) アルゴリズムとして使用し、非常に低い遅延を実現します。

メモリ消費量が増えるというトレードオフです。
Qdrant は、量子化、ペイロードとベクトルのディスク上ストレージ オプション、構成可能な精度など、広範なコスト効率の最適化もサポートしています。これにより、設定を調整して、必要なレベルの取得品質を提供することと関連コストとの間のバランスを見つけることができます。
Qdrant を社内で実行する必要がある理由
HubSpot は、社内で自己管理型のオープンソース データベース システムを実行するための豊富な専門知識を持っており、インフラストラクチャ チームは長年にわたって、活用したい数多くのソリューション、ルール、ベスト プラクティスを構築してきました。このような広範なインフラストラクチャ層があることで、トレース、コスト追跡、レート制限、スケーリングなど、HubSpot のすべての内部ツールとシームレスに統合し、その恩恵を受けることができます。 Qdrant を社内で実行することで、企業向け AWS 料金を利用することで、ニーズに合わせてインフラストラクチャを簡単に微調整し、コストを制御することができます。最後に、社内で実行することで、顧客データとインフラストラクチャを完全に制御できるようになり、HubSpot の広範なセキュリティのベスト プラクティスがすべてこのデータ ストアに確実に適用されるようになります。
VaaS (Vector-as-a-Service) プラットフォームは現在 38 以上のチームによって使用されており、5 つのリージョンと 2 つの環境にわたる 140 以上のクラスターのフリートで実行される 200 以上のインデックスがあります。製品チームごとにクラスターを分離して、プラットフォーム全体に影響を与える障害の影響範囲を最小限に抑えます。基盤となるベクトル データベースには、合計すると、さまざまなリージョンにわたって 200 億以上のベクトルが保存され、最大のインデックスには 95 億のベクトルが含まれ、平均インデックス サイズは 9,500 万のベクトルになります。
このプラットフォームは、書き込みについては 5,000 以上の RPS (スパイクでは最大 100,000 RPS) を処理し、すべてのリージョンで読み取りについては 1,000 以上の RPS を提供します。特に次のような場合には、トラフィックが急増することもあります。

大規模な埋め戻しが実際の交通と並行して実行される場合。したがって、エンジニアリングのユースケースをサポートしながら最高の顧客エクスペリエンスを提供するには、基盤となるインフラストラクチャがスケーラブルで信頼性が高い必要があります。
すべてのリージョンにわたる合計読み取り数
すべてのリージョンにわたる合計書き込み数
2023 年に遡ると、セマンティック検索が普及し始めたばかりで、内部 POC は、必要なパイプラインとその上の API レイヤーとともに Qdrant を実行するようにセットアップされました。コンシューマーの数は限られており、比較的操作が簡単な 12 個の Qdrant クラスターがありました。最初のソリューションとして、Helm を利用してクラスター/環境ごとのセットアップを定義し、レプリカの数、ディスク容量、RAM、CPU などのパラメーターを手動で調整しました。クラスター定義と関連クラスター アーティファクト (Kafka トピック、ワーカーなど) はコードで定義され、手動で作成されました。小規模なクラスターのフリートにとって、これは高速に繰り返し、オーバープロビジョニングを利用してトラフィックの急増に対処し、セットアップをシンプルに保つことができる優れたソリューションでした。すべてのクラスターは StatefulSet として定義され、同じ名前空間を共有しました。当時、Qdrant はすでに約 10 のインデックスにわたる 20 億以上のベクトルを処理していました。
Helm デプロイメントは、すべてのクラスターで同じ k8s 名前空間を共有しました
セマンティック検索に依存する機能がますます増えており、VaaS が複数の AI 機能の中核となり、RAG (検索拡張生成) からエージェント エクスペリエンスまでの機能を強化しているため、増大する需要に対応するために POC から完全な運用規模に移行する時期が来たことが明らかになりました。
セマンティック検索の需要の急速な増加に伴い、各クラスターに追加のアーティファクトが必要になり、迅速で一貫性のあるクラスターのセットアップに対するニーズが高まるなど、セットアップの複雑さが進化し始めました。ヘルムも強力です

ただし、これは純粋にテンプレートおよび展開ツールであり、次のことはできません。
外部メトリクスに基づいて自動スケールする
複雑な状態認識ライフサイクル管理のためのカスタム ロジックを実行する
その結果、標準の操作手順を自動化することでチームの運用負荷を軽減しながら、信頼性とスケーラビリティを確保するために、Helm から Kubernetes Operator パターンに移行することが決定されました。
Kube-operators フレームワークへのオンボーディング
HubSpot には、パターンを実装する「Kube-operators」と呼ばれる内部フレームワークがあります。主なアイデアは、オペレーターを使用してシステムの定義された状態を追跡および維持することにより、K8 上で実行されているシステムのライフサイクル管理を自動化することです。 HubSpot とインフラストラクチャ チームのニーズを満たすように設計されたオペレーターを構築するための堅牢なツールキットを提供します。また、複数の HubSpot インフラストラクチャ ツールとのシームレスな統合が可能になり、Kubernetes Operator のロジックでそれらのツールを使用できるようになります。
オペレータは実際の状態を CR の望ましい状態に向けて継続的に駆動します
Kube オペレーター上で Qdrant を実行すると、関連するすべてのアーティファクトおよびサブシステムとともにクラスターをインスタンス化し、廃止することができます。また、複数のワークフローを自動化して、チームの運用負荷を軽減します。
トランスレーターを通じて、クラスター CR は名前空間、Qdrant StatefulSet、インデクサー、Kafka トピック、およびシャード管理操作に伝播されます。
セットアップの主なコンポーネントは、Translator アーティファクトです。トランスレーターは、割り当てられた CR を監視し、ループ内で実行し、60 秒ごとに状態を再評価し、必要な更新を適用し、基になるリソースが定義された状態と同期していることを確認します。これは、トランスレータの追加/削除を可能にし、より多くの機能をシステムにオンボードできる拡張可能なアーキテクチャです。さらに、トランスレーターには必要なロジックがすべて含まれています。

d 基礎となるアーティファクトをレンダリングするため。したがって、必要な自動化を行うことで拡張できます。
現在のセットアップには 3 つのトランスレータがあります。
クラスター トランスレーター: クラスターとダウンストリーム アーティファクトの個々の名前空間のインスタンス化を担当します。
Qdrant ノードセット トランスレーター: Qdrant クラスターのインスタンス化を担当します。指定された数のポッド、メモリ、ディスク ストレージ、パラメータなどを使用して、必要なすべての K8s アーティファクトをインスタンス化します。
インデクサー ノードセット トランスレーター: インデクサーは、受信書き込みトラフィックを処理し、Qdrant にインデックスを付けるワーカーです。クラスターごとに複数のインデクサーがあり、さまざまなトラフィック優先順位などを処理します。このトランスレーターは主に、必要なインデクサーのインスタンス化と、必要な Kafka トピックの作成を担当します。
Helm セットアップからの移行に伴い、K8s 名前空間戦略も更新し、単一の名前空間を使用するのではなく、クラスターごとの名前空間アプローチに移行しました。これにより、次のような利点が得られます。
隔離: 構成ミスなどの問題による影響範囲が減少するため、より高い信頼性が得られます。
名前空間固有の構成: 名前空間固有のリソース、制限などを設定できます。
簡素化されたコスト追跡: 名前空間ごとのコストの追跡がはるかに簡単になり、これは現在の規模にとって重要です。
Helm セットアップからの移行により、クラスターごとの名前空間アプローチに移行しました。
移行後、クラスター構成はシンプルになり、Qdrant クラスターと他のインフラストラクチャー成果物の両方をカスタム リソース (CR) として定義する単一の構成を作成または変更する必要があります。
Kube オペレーターへのオンボーディングとともに、VaaS (API レイヤー + データ パイプライン) と基盤となるインフラストラクチャの間の契約を再定義して、Qdrant クラスターが自己検出可能であり、追加のコード変更が必要ないことを保証します。

新しいクラスターが作成されます。重要な部分は、各検索/取り込み/削除リクエストにクラスター メタデータが必要であるということです。したがって、クラスター ルックアップはエンドツーエンドの遅延に最小限の影響を与える必要があります。解決策として、レイテンシーのオーバーヘッドを最小限に抑えるために、クラスターのメタデータを K8s CRD (カスタム リソース定義) に抽象化することにしました。さらに、私たちはすでに、VaaS の他のメタデータに対して CRD アプローチを使用して成功していました。
トランスレーターは、VaaS が新しいクラスターを検出してトラフィックをルーティングするために監視するクラスター メタデータ CR を作成します。
エンドツーエンドのフローを上の図に示します。
Kube-operators でのクラスターの作成は、すべてのポッドが開始されるまで待機し、クラスターのメタデータ CR の作成をトリガーします。
VaaS は自動的にクラスター メタデータ CR を検出し、対応するコレクションのそのクラスターへのトラフィックのルーティングを開始します。
クラスターの自動メンテナンス (操作とスケーリング)
Kube オペレーターが導入される前は、クラスターのスケールダウンには 4 つの手順が必要でした
Kube オペレーターが登場する前は、クラスターのメンテナンスは非常に困難でした。オープンソースの Qdrant は、他の分散データ ストアが行うすぐに使える自動メンテナンス操作の多くをサポートしていません。 Qdrant は、シャード転送などのアクションを実行するための構成要素として機能する API を提供しますが、完全なワークフローの実装はユーザーに委ねられます。実際には、これは水平方向のスケーリングが非常に困難であるか、エラーが発生しやすいことを意味します。 Kube オペレーターの前にクラスターを水平方向にスケールダウンするには、以下を行う必要がありました。
オンデマンド ジョブを実行して、排除されたピアからシャードを転送し、コンセンサスからピアを削除します。
クラスター Helm チャートを更新してポッド数を減らします。
削除されたポッドの PVC を削除します
Kube オペレーターを使用して、上記のすべてをトランスレーターに組み込みました。必要なのは、クラスター構成内のレプリカの数を更新することだけです。監視している Translator デーモンは、

調整し、レプリカ数が減少したことを検出し、エビクション プロセスを自動的に実行します。これにより、4 つの操作が 1 つに統合されます。この設定では、プロセスは完全に自動化されており、開始するにはクラスター CR 上の 1 つの構成変更が必要です。
ピアからのシャードの安全な削除
コンセンサスから削除される前に、排除されたポッドからシャードが排出される
クラスターをスケールダウンするには、コンセンサスから安全にポッドを削除する前に、削除されるポッドからすべてのシャードを転送する必要がありました。簡単そうに聞こえますが、これがうまくいかない可能性はたくさんあります。シャードが完全に転送される前にポッドをコンセンサスから削除すると、データが失われる危険があります。残りのポッド数がレプリケーション係数を下回ると、耐久性の保証に違反します。
そこで、厳密なシーケンスをトランスレータに焼き付けました。
削除されたポッドからすべてのシャードを避難させます。
宛先ですべての転送が完了したことが確認されるまで待ちます。
ポッドをコンセンサスから削除します。
エビクション プロセスは、残りのポッド数がレプリケーション係数を満たさない場合、エビクションを事前に拒否し、単一の転送を発行する前に、エビクションされたポッド上のすべてのシャードに有効な宛先があることを検証します。転送中にシャードがデッド状態になった場合

[切り捨てられた]

## Original Extract

Discover how HubSpot built a scalable AI retrieval infrastructure, managing over 20 billion vectors with Qdrant, to enhance semantic search and support diverse applications.

Building the AI Retrieval Infrastructure Behind 20 Billion+ Vectors at HubSpot
Semantic search plays a critical role in HubSpot operations and powers a wide variety of use cases, from agents and RAG (Retrieval Augmented Generation) to contact deduplication. Starting from the POC a few years back, it has grown to a scale of tens of billions of vectors, utilized by 38+ teams. With the growth of agent usage, semantic search capabilities have become even more important to ensure agents can quickly retrieve the information they need while maintaining retrieval quality. This article describes the journey of vector DB infrastructure, from a small POC to the tens of billions of entries powering hundreds of use cases across HubSpot.
VaaS (Vector as a Service) is HubSpot's centralized vector storage and search platform built on top of a vector database. It provides an API layer and features such as:
Embeddings generation based on the collection setup
It sits in front of the vectors database, in our case Qdrant , and is directly exposed to the consumers. It handles access control, basic validation, embeddings generation and then proxies requests to Qdrant. For read operations, requests are handled synchronously, whereas for write operations the flow is asynchronous with minimal consumer lag.
VaaS platform high-level architecture
HubSpot uses the Qdrant vector database for all semantic search operations. Qdrant is an open-source vector database that can be deployed on-premises and provides an extensive API covering most modern semantic search features, such as named vectors, hybrid search, multi-stage querying, and weighted reranking. It uses HNSW (Hierarchical Navigable Small World) as an ANN (Approximate Nearest Neighbor) algorithm, which provides very low latency at the trade-off of higher memory consumption.
Qdrant also supports an extensive set of cost-efficiency optimizations, such as quantization, on-disk storage options for payloads and vectors, and configurable precision. This allows us to tune the setup to find a balance between providing the required level of retrieval quality and associated costs.
Why we need to run Qdrant in-house
HubSpot has a lot of expertise in running self-managed open-source database systems in-house, and over the years our infrastructure team has built numerous solutions, rules, and best practices that we want to take advantage of. Having such an extensive infrastructure layer lets us seamlessly integrate with and benefit from all of HubSpot’s internal tooling, including tracing, cost tracking, rate limiting, and scaling. Running Qdrant in-house enables us to easily fine-tune the infrastructure to our needs and control costs by utilizing our corporate AWS rates. Finally, running in-house gives us full control over customer data and infrastructure, and ensures all of HubSpot’s extensive security best practices are applied to this data store.
The VaaS (Vector-as-a-Service) platform is currently used by 38+ teams and has more than 200 indexes that run on a fleet of 140+ clusters across 5 regions and 2 environments. We isolate clusters by product teams to minimize the blast radius from failures affecting the entire platform. Collectively, the underlying vector database stores 20+ billion vectors across different regions, with the largest index holding 9.5 billion vectors and an average index size of 95 million vectors.
The platform handles 5,000+ RPS for writes, with spikes up to 100,000 RPS, and serves 1,000+ RPS for reads across all regions. The traffic can also be spiky, specifically in cases when large backfills run in parallel with the live traffic; therefore, the underlying infrastructure should be scalable and reliable to provide the best customer experience while supporting engineering use cases.
Total reads across all regions
Total writes across all regions
Back in 2023, semantic search had just started to take off, and the internal POC was set up to run Qdrant alongside the required pipelines and an API layer on top of it. There was a limited number of consumers and a dozen Qdrant clusters that were relatively easy to operate. As a first solution, Helm was utilized to define per-cluster/environment setup and parameters such as number of replicas, disk space, RAM, CPU, etc., were adjusted manually. The cluster definition and related cluster artifacts (Kafka topics, workers, etc.) were defined in code and created manually. For a small fleet of clusters, it was a great solution that allowed us to iterate fast, utilize overprovisioning to handle traffic spikes, and keep the setup simple. All the clusters were defined as StatefulSets and shared the same namespace. At that time, Qdrant was already handling 2+ billion vectors across ~10 indexes.
Helm deployment shared the same k8s namespace across all clusters
With more and more features relying on semantic search, and VaaS being at the core of multiple AI capabilities, powering functionalities from RAG (Retrieval Augmented Generation) to agentic experiences, it became clear that it was time to transition from the POC to full production scale to address the growing demand.
Alongside the rapid growth in demand for semantic search, our setup started to evolve in terms of complexity, with additional artifacts required for each cluster and a growing need for quick, consistent cluster setup. Helm is a powerful tool, but it’s purely a templating and deployment tool that cannot:
Auto-scale based on external metrics
Run custom logic for complex state-aware lifecycle management
As a result, the decision was made to migrate from Helm to a Kubernetes Operator pattern to ensure reliability and scalability while reducing operational load on the team by automating standard operating procedures.
Onboarding to Kube-operators framework
HubSpot has an internal framework called `Kube-operators` that implements the pattern . The main idea is to automate the lifecycle management of systems running on K8s by using operators to track and maintain the system’s defined state. It provides a robust toolkit for building operators designed to meet the needs of HubSpot and our infrastructure teams. It also enables seamless integration with multiple HubSpot infrastructure tools, which makes it possible to use them in the Kubernetes Operator’s logic.
The operator continuously drives actual state toward the CR's desired state
Running Qdrant on Kube-operators allows us to instantiate and decommission clusters together with all related artifacts and subsystems. It also automates multiple workflows to reduce operational load on the team.
Through translators, a Cluster CR propagates into namespaces, Qdrant StatefulSets, indexers, Kafka topics, and shard-management operations
The main components in the setup are Translator artifacts. Translators watch assigned CR, running in a loop, re-evaluating the state every 60s, applying the required updates, and ensuring that underlying resources are in sync with the defined state. This is an extensible architecture that allows adding/removing translators and therefore onboarding more features to the system. In addition, the translator contains all the logic required for rendering underlying artifacts; therefore, it can be extended with required automation.
Our current setup has 3 translators:
Cluster translator: responsible for instantiating the individual namespaces for clusters and downstream artifacts.
Qdrant nodeset translator: responsible for instantiating Qdrant clusters. It instantiates all required K8s artifacts with the specified number of pods, memory, disk storage, parameters, etc.
Indexer nodeset translator: indexers are the workers that handle incoming write traffic and index it into Qdrant. There are multiple indexers per cluster handling different traffic priorities, etc. This translator is primarily responsible for instantiating the required indexers and creating necessary Kafka topics.
Migrating from the Helm setup, we also updated the K8s namespace strategy and moved towards a namespace-per-cluster approach instead of using a single namespace. This provides the following benefits:
Isolation: provides higher reliability, since issues such as misconfigurations have a reduced blast radius.
Namespace-specific configuration : we can set up namespace-specific resources, limits, etc.
Simplified cost tracking: tracking costs per namespace is much easier, which is important for current scale.
Migrating from the Helm setup we moved towards a namespace-per-cluster approach
After the migration, the cluster configuration became simple and requires creating or modifying a single config that defines both the Qdrant cluster and the other infrastructure artifacts as Custom Resources (CRs):
Together with onboarding to Kube-operators, we redefined the contract between VaaS (API layer + data pipelines) and the underlying infrastructure to ensure Qdrant clusters are self-discoverable and do not require any additional code changes when a new cluster is created. The important part is that cluster metadata is required for each search/ingest/delete request; therefore, cluster lookup should have minimal effect on end-to-end latency. As a solution, we decided to abstract cluster metadata into a K8s CRD (Custom Resource Definition) to ensure the latency overhead is minimal. Moreover, we had already successfully used the CRD approach for other metadata in VaaS.
The translator writes cluster-metadata CRs that VaaS watches to discover new clusters and route traffic
The end-to-end flow is shown in the diagram above:
Cluster creation in Kube-operators waits until all pods are started, then triggers cluster metadata CR creation.
VaaS automatically discovers cluster metadata CRs and starts routing traffic to that cluster for the corresponding collections
Automated Cluster Maintenance (Operations & Scaling)
Before Kube-operators, scaling down a cluster took four steps
Before Kube-operators, cluster maintenance was very difficult. Open-source Qdrant does not support a lot of automated maintenance operations out of the box that other distributed data stores do. Qdrant provides APIs that serve as building blocks to perform actions like shard transfers, but leave it up to the user to implement the full workflow. In practice, this means that horizontal scaling can be very difficult or error-prone. To scale down a cluster horizontally prior to Kube-operators, we needed to:
Run an on-demand job to transfer shards off evicted peers and remove peers from consensus.
Update cluster Helm chart to lower pod count.
Delete the PVCs for the removed pods
With Kube-operators, we baked all of the above into our Translator. All we have to do is update the number of replicas in the cluster config; the watching Translator daemon will reconcile, detect that the replica count has been lowered, and automatically run the eviction process - consolidating 4 operations into 1. With our setup, the process is fully automated and requires a single config change on the cluster CR to kick off.
Safe eviction of shards from Peers
Shards draining off evicted pods before they are removed from consensus
To scale down a cluster, we needed to transfer every shard off the pods being removed before we could safely remove them from consensus. It sounds simple, but there are many ways this can go wrong. If you remove a pod from consensus before its shards are fully transferred, you risk data loss. If the remaining pod count drops below the replication factor, you violate your durability guarantees.
So we baked a strict sequence into the Translator:
Evacuate all shards off the evicted pods.
Wait until every transfer is confirmed complete on the destination.
Remove the pods from consensus.
The eviction process rejects the eviction upfront if the remaining pod count can't satisfy the replication factor, and validates that every shard on the evicted pods has a valid destination before issuing a single transfer. If a shard ends up in a Dead state during transfe

[truncated]
