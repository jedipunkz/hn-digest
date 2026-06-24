---
source: "https://www.hopsworks.ai/post/rolling-aggregations-for-real-time-ai"
hn_url: "https://news.ycombinator.com/item?id=48659199"
title: "Rolling Aggregations for Real-Time AI"
article_title: "Rolling Aggregations for Real-Time AI | Hopsworks"
author: "LexSiga"
captured_at: "2026-06-24T13:43:15Z"
capture_tool: "hn-digest"
hn_id: 48659199
score: 1
comments: 0
posted_at: "2026-06-24T13:11:20Z"
tags:
  - hacker-news
  - translated
---

# Rolling Aggregations for Real-Time AI

- HN: [48659199](https://news.ycombinator.com/item?id=48659199)
- Source: [www.hopsworks.ai](https://www.hopsworks.ai/post/rolling-aggregations-for-real-time-ai)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T13:11:20Z

## Translation

タイトル: リアルタイム AI のローリング集計
記事のタイトル: リアルタイム AI のローリング集計 |ホップワークス
説明: ローリング集計により、真にインタラクティブな AI の動作および異常シグナルを最新の状態に保ちます。

記事本文:
リアルタイム AI のためのローリング集計 | Hopsworks 製品の価格設定 ブログ 辞書リソース ニュース 無料でスタート すべての投稿 機能ストア リアルタイム AI のローリング集計
リアルタイム AI システム向けの AI 機能の女王
ローリングアグリゲーションは、真にインタラクティブな AI のために、動作と異常のシグナルを最新の状態に保ちます。最新のアプローチの増分ビューと並列プッシュダウン集計により、スケーラブルで低遅延になります。 Hopsworks を使用すると、チームは左シフト (Feldera) または右シフト (RonDB) を選択して、本番環境でのシンプルさとミリ秒未満の鮮度のバランスを取ることができます。
ローリング アグリゲーションは AI システムにとって最も有用な入力データの 1 つであり、リアルタイムでの行動変化の検出と異常検出を可能にします。最近の傾向とパターンを圧縮表現でキャプチャし、インタラクティブ AI システムが機能の鮮度に対する最低遅延要件を満たすことができるようにします。ローリングアグリゲーションは、リアルタイム ML システム (例: クレジット カード詐欺やパーソナライズされた推奨事項) と対話型エージェント (例: 圧縮されたユーザー履歴/行動、または最近のアクティビティの概要) の両方で使用されます。
ローリング集計 - AI 集計の女王
ローリング集計は、継続的にシフトするデータ ウィンドウにわたって統計を計算します。各時点で最後の N 個の値または期間を集計します。
図 1. ローリング集計は、時系列データ (過去 1 時間など) のウィンドウ サイズにわたって集計関数 (SUM、AVG、MIN、MAX、STDDEV など) を計算します。
ストリーム処理における従来のウィンドウ集計 (Apache Flink など) は、計算コストが高いため、ローリング集計を計算するように設計されていません。新しいイベントごとに、バケット内の N 個のイベントすべての集計を再計算する必要があります。代わりに、タンブリングまたはホッピング (スライディング) ウィンドウを使用してデータをグループ化します。

過去 10 分または 1 時間などの時間枠を設定し、それらの時間枠にわたって集計を計算します。
図 2. タンブリング ウィンドウとホッピング ウィンドウにより、イベントの到着と集計結果の計算の間に遅延が生じます。ローリング集計は、イベントが到着するとすぐに計算されます。
図 2 は、タンブリング ウィンドウがウィンドウ長とウォーターマークが経過した後にのみ集計を出力するのに対し、ホッピング ウィンドウはホップ サイズが経過するまで待機する様子を示しています。タンブリング ウィンドウとホッピング ウィンドウにより、ストリーミング データの計算集計が計算上扱いやすくなりますが、遅延が発生します。出力集約は、ウィンドウの長さ + ウォーターマークまたはホップ サイズと同じくらい古くなります。 AI の観点から言えば、彼らが出力する特徴データは古いものです。データが古いということは、対話型 AI アプリケーションがインテリジェントではなく、単に遅延するだけであることを意味します。対照的に、ローリング集計では、新しいイベントが到着するとすぐに結果が出力され、新しい特徴データが生成されます。
図 3. ローリング アグリゲーションの機能の鮮度向上の簡単な歴史。
図 3 では、タンブリング/ホッピング ウィンドウからローリング集計を計算するソリューションに至るまでの簡単な歴史を示しています。最初に採用されたアプローチは、ストリーム処理とリクエスト時のさらなる計算を組み合わせたタイル ウィンドウ集約でした。 Feldera は最近、ストリーミング処理の増分ビューに基づいた低コストのソリューションを開発しました。そして最近、RonDB を使用して、ストリーム処理の必要性を回避する、集約の並列処理のためのネイティブ データベース サポートを開発しました。ここで、これらのアプローチについて説明します。
タイル ウィンドウ集計と Chronon を使用した Shift Left と Shift Right
AirBnB の Chronon フレームワークは、ローリング集計の計算オーバーヘッドを削減する最初の新しいソリューションを提供しました。

これは、タイル化された時間ウィンドウ集計と呼ばれるアプローチです。 Tecton Rift ( DuckDB ベース) および Chalk.ai (Apache Velox ベース) も、スケーラブルなローリング アグリゲーション用のこのソリューションのバリアントを提供します。
正確な 240 時間の集計を計算したいとします。イベントを 24 時間のタイルに分解し、毎日午前 12 時に計算できます。考え方としては、24 時間タイルの部分集計を構成することで 240 時間集計を計算できるということです。これは一部の集計 (例: min 、 max ) では簡単に機能しますが、他の集計 (例: means ) では追加の状態を維持する必要があります。
ここで、ローリング集計を計算するリクエストが午後 1 時に到着すると想像してください。あなたのタイルは午前 12 時から午後 11 時 59 分までしかありません。当日のイベント (午前 12 時から午後 1 時まで) のタイルはまだ計算されておらず、間隔の最終日の午後 1 時からのイベントのタイルはありません (その日のタイルには、間隔に含まれていないイベントが含まれています)。タイルの外側にあるこれらのイベントは、それぞれ先頭イベントと末尾イベントと呼ばれます。タイル ウィンドウ集計は、位置合わせされていない先頭/末尾イベントに対するオンデマンド集計を使用して部分集計を構成することによって計算されます。
図 4. タイル化集計では、事前計算された部分集計 (タイル) とオンデマンド計算を組み合わせて、部分集計と最近のイベント (タイルに含まれない先頭/末尾のイベント) の両方から集計を構成します。
より具体的には、タイルは長さ N のウィンドウを M 個のタイルに分割します (M<<N)。ストリームまたはバッチ処理は、M 個のタイルにわたる部分集計を計算し、最終的な集計結果は、それぞれウィンドウ間隔の開始時と終了時の位置合わせされていないヘッド/テール イベントとともに、M 個の部分集計値から構成されます。タイル化された集約は、両方の左シフト計算を必要とするため、操作上のオーバーヘッドが高くなります。

先頭/末尾の計算に必要なすべてのイベントを保存するとともに、右シフト計算を実行します。インクリメンタル ビューは、問題を、計算の複雑さがより低い完全な左シフトの解決策に変換します。
インクリメンタル ビューと Feldera によるシフト レフト
ローリング集計は、Apache Flink などのストリーム処理システムで OVERaggregates を使用して計算できます。ただし、Apache Flink の OVERaggregates は多数のワーカーに分割できますが、新しいイベントが発生するたびに集計関数の再計算がトリガーされ、その計算コストがウィンドウ サイズに比例するため、ウィンドウ サイズの増加やイベント スループットの増加にはうまく対応できません (図 5 を参照)。
図 5. インクリメンタル コンピューティング エンジンは、ローリング アグリゲーションの計算を O(N) の計算複雑さから O(1) の計算複雑さの問題に変換します。
増分ビューは、新しいイベントの到着時に集計の完全な再計算を回避することで、スケーラビリティの課題を解決します。代わりに、以前に計算された値を再利用し、新しいイベントまたは削除されたイベントによって導入された変更のみを適用します。その結果、実行される作業はウィンドウ サイズではなく、入出力の変更に比例します。
Feldera は、オープンソースのインクリメンタル コンピューティング エンジンをリードするものです。 SQL でローリング集計を記述すると、それが Rust コードにトランスパイルされて、高いパフォーマンスが実現されます。 Feldera は、ストリーミング エンジン DBSP (信号処理に触発されたデータベース) を通じて増分ビューのメンテナンスをサポートしています。 DBSP は、Z セットを使用して増分ビューを実装します。Z セットは、どの要素が存在するかだけでなく、その数が時間の経過とともにどのように変化するかを追跡するリレーショナル セットの一般化です。従来のリレーショナル セットでは、要素は存在する (カウント = 1) か、存在しない (カウント = 0) かのいずれかです。 Z セットでは、各要素は整数を持ちます。

正、ゼロ、または負の値を指定できるカウント:
正の数は挿入 (イベントの追加) を表します。
負の数は削除 (イベントの削除) を表します。
これにより、Z セットは、各ステップで完全な状態を保存せずに、デルタ、つまり 2 つの状態間の正味の変化を表すことができます。
Hopsworks/RonDB の並列プッシュダウン集約による右シフト
多くの組織は、運用が複雑になりコストがかかるため、アーキテクチャにストリーム処理を導入することを避けようとしています。データの量と速度がハイパースケーラー レベルにない場合は、生のイベントを Hopsworks オンライン データベースである RonDB にストリーミングできます。クライアントが集計計算を要求すると、RonDB のデータベース ノードでオンデマンドで並列計算されます。
図 6. RonDB は、集計計算をデータ ノードにプッシュダウンし、データ ノードが並列で計算を行います。
集計は、インデックス スキャン、LIMIT、プロジェクションなど、同じクエリ内の他の SQL 操作と組み合わせる (そしてデータベース ノードにプッシュ ダウンする) ことができます。 DynamoDB および Redis 上に構築された他の機能計算エンジンもオンデマンドで集計を計算できますが、データベースの機能によって制限されます。たとえば、Redis では、クライアントはテーブル内のすべての列とすべての行を読み取る必要があります (述語プッシュダウンまたは射影プッシュダウンがサポートされていないため)。対照的に、RonDB では、テーブルからクライアントに返された列のみが読み取られ、述語 (たとえば、「EU」に基づくユーザーに対してのみ計算) が、集計が計算される前にデータベース カーネルに適用されます。その結果、選択的なクエリのレイテンシーが桁違いに短縮され、スループットが向上します。
RonDB プッシュダウン集約パフォーマンス ベンチマーク
2 データノード: c6i.4xlarge (16c32g)
1 ベンチマーク クライアント: c6i.8xlarge (32c64g)
NdbAggregator C++ A による RonDB プッシュダウン集約

PI
接続: 4 Ndb_cluster_connection オブジェクト、ラウンドロビンで分散された 32 ワーカー (接続ごとに 8)
タイミング: クエリごとにナノ秒精度の std::chrono::steady_ Clock
ウォームアップ: 15 秒間の各測定ウィンドウの前に 5 秒が破棄されます。
データ: ランダム値 — val1/val2 BIGINT、val3/val4 DOUBLE、filter_date DATE (2024-01-01 ～ 2024-12-31)
行数が異なる、同一のスキーマを持つ 5 つのテーブル: 100、500、10K、100K、500K。
CREATE TABLE bench_qps_ {N} (
id INT NOT NULL AUTO_INCREMENT、
val1 BIGINT NOT NULL 、 -- ランダムな 0 ～ 999999
val2 BIGINT NOT NULL 、 -- ランダムな 0 ～ 999999
val3 DOUBLE NOT NULL 、 -- ランダムな 0 ～ 1000
val4 DOUBLE NOT NULL 、 -- ランダムな 0 ～ 1000
filter_date DATE NOT NULL 、 -- 2024-01-01 から 2024-12-31 までのランダム
パディング VARCHAR ( 100 ) NOT NULL 、 -- 80 文字のランダムな文字列
主キー (ID)、
INDEX idx_date (filter_date) -- 順序付きインデックス、第 4 四半期で使用
) エンジン = NDB;
val1、val2: SUM() 集計用の整数列
val3、val4: AVG()、MIN()、MAX()、および式集計用の浮動小数点列 (val1 * val3、val3 + val4)
filter_date: 2024 年にわたって均一に分布。 Q4 フィルター >= '2024-07-01' (選択率 ~50%)
パディング: 現実的な行幅をシミュレートします (行あたり最大 130 バイト)
SELECT COUNT ( * )、SUM (val1)、SUM (val3)、COUNT (val3) FROM テーブル
Q2: マルチ集計 (6 つの集計)
10K 行でも PA は良好なパフォーマンスを維持しますが、他のエンジンは 6 つの集計を計算するためにすべての行を転送しますが、PA はデータ ノードで 1 回のパスで集計を計算します。
SELECT COUNT ( * )、SUM (val1)、SUM (val2)、SUM (val3)、COUNT (val3)、MIN (val4)、MAX (val4) FROM テーブル
Q3: 式の集約
算術式 (val1 * val3、val3 + val4) はデータ ノードで直接計算されます。これは、スキャン I/O が計算よりも優先される大きなテーブルの例です。
SELECT COUNT ( * )、SUM (val1 *

val3)、SUM (val3 + val4)、COUNT (val3 + val4) FROM テーブル
Q4: フィルタリングされた集計 (インデックス スキャン)
SELECT COUNT ( * )、SUM (val1)、SUM (val3)、COUNT (val3) FROM table WHERE filter_date >= '2024-07-01'
第 1 四半期から第 4 四半期までのベンチマーク結果
以下の表 1 は、各集計が計算される行数を変化させながら、第 1 四半期から第 4 四半期までのベンチマークの実験パフォーマンスを示しています。各クエリと集計あたりの行数について、1 秒あたりのクエリ数 (QPS) とレイテンシのスループットを測定します。
表 1. クエリ 1 ～ 4 のベンチマーク結果。集計が計算される行数を変化させ、1 秒あたりのクエリ (QPS) でのレイテンシーとスループットを測定します。
データ ノードが 2 つだけの結果 (RonDB は最大 144 データ ノードまで拡張可能) から、集計の行数が 10,000 に達すると、Q1 ～ Q3 のレイテンシが増加することがわかります。ただし、第 4 四半期では、このクエリには範囲フィルタリングと集計の両方をデータ ノードにプッシュするインデックス スキャンが含まれており、行転送が完全に排除され、レイテンシが低く保たれるため、レイテンシは非常に低い (1.9 ミリ秒) ままです。つまり、フィルターも含まれる集計がある場合、同じハードウェア上でプッシュダウン集計をより高いパフォーマンスに拡張できるということです。
比較

[切り捨てられた]

## Original Extract

Rolling aggregations keep behavioral and anomaly signals fresh for truly interactive AI....

Rolling Aggregations for Real-Time AI | Hopsworks Product Pricing Blog Dictionary Resources News Start Free All posts Feature Stores Rolling Aggregations for Real-Time AI
The Queen of AI Features for Real-Time AI systems
Rolling aggregations keep behavioral and anomaly signals fresh for truly interactive AI. Modern approaches incremental views and parallel pushdown aggregations make them scalable and low latency. With Hopsworks, teams can choose shift-left (Feldera) or shift-right (RonDB) to balance simplicity and sub-millisecond freshness in production.
Rolling aggregations are among the most useful input data for AI systems, enabling behavioral change detection and anomaly detection in real-time. They capture recent trends and patterns in a compressed representation, enabling interactive AI systems to meet the lowest latency requirements for feature freshness. Rolling aggregations are used by both real-time ML systems (e.g., credit card fraud and personalized recommendations) and interactive agents (e.g., compressed user history/behavior or summary of recent activity).
Rolling Aggregations - the Queen of AI Aggregations
A rolling aggregation computes statistics over a continuously shifting window of data. You aggregate over the last N values or time period at each point.
Figure 1. A rolling aggregation computes an aggregation function (e.g., SUM, AVG, MIN, MAX, STDDEV, etc) over a window size of time-series data (e.g. last hour).
Traditional windowed aggregations in stream processing (e.g., Apache Flink) are not designed to compute rolling aggregations, as they are computationally expensive - for every new event, you have to recompute the aggregate over all N events in the bucket. Instead, they use tumbling or hopping (sliding) windows to group data into windows of time, such as the last 10 minutes or hour, and compute aggregations over those windows.
Figure 2. Tumbling windows and hopping windows introduce a delay between when an event arrives and when an aggregation result is computed. Rolling aggregations are computed immediately when the event arrives.
Figure 2 shows how tumbling windows only output aggregations after the window length and watermark has passed, while hopping windows wait until the hop size has passed. While tumbling and hopping windows make computing aggregations over streaming data computationally tractable, they introduce latency. The output aggregations are as stale as the window length + watermark or the hop size. In AI terms, the feature data they output is stale. Stale data means your interactive AI applications will not be intelligent, just laggy. In contrast, rolling aggregations output results immediately when a new event arrives, producing fresh feature data.
Figure 3. Brief history of increasing feature freshness for rolling aggregations.
In Figure 3, you can see a brief history of the journey from tumbling/hopping windows to solutions for computing rolling aggregations. The first adopted approach was tiled window aggregations that combined stream processing with further computation at request-time. A lower cost solution was developed by Feldera recently based on incremental views for streaming processing. And recently, with RonDB, we developed native database support for parallel processing of aggregations - avoiding the need for stream processing. We now describe these approaches.
Shift Left and Shift Right with Tiled Window Aggregations and Chronon
AirBnB’s Chronon framework provided the first novel solution to reduce the computational overhead of computing rolling aggregations with an approach called tiled time window aggregations . Tecton Rift (based on DuckDB ) and Chalk.ai (based on Apache Velox ) also provide variants on this solution for scalable rolling aggregations.
Say you want to compute a precise 240 hour aggregation, you could decompose the events into 24-hour tiles, computed daily at 12am. The idea is that you can compute the 240 hour aggregation by composing partial aggregates for the 24-hour tiles. This works trivially for some aggregations (e.g., min , max ), but requires maintaining additional state for others (e.g., mean ).
Now imagine, a request arrives to compute a rolling aggregation at 1pm. Your tiles are only from 12am to 11.59pm. You will not have yet computed a tile for the current day’s events (from 12am-1pm) and you won’t have a tile for the events from 1pm of the last day in the interval (the tile for that day contains events not included in the interval). These events that lie outside the tiles are called head and tail events, respectively. Tiled window aggregations are computed by composing the partial aggregates with on-demand aggregations over the unaligned head/tail events.
Figure 4. Tiled aggregation combines precomputed partial aggregates (tiles) with on-demand computation to compose aggregations from both partial aggregates and recent events (head/tail events that are not included in a tile).
More specifically, tiles partition a window of length N into M tiles , where M<<N. Stream or batch processing computes the partial aggregates over the M tiles, and the final aggregation result is composed from the M partial aggregate values along with unaligned head/tail events at the start and end of the window interval, respectively. Tiled aggregation has high operational overhead as it requires both shift left computation and shift right computation, along with storing all events that may be needed for head/tail computations. Incremental views transform the problem into a fully shift-left solution with lower computational complexity.
Shift Left with Incremental Views and Feldera
Rolling aggregations can be computed in stream processing systems such as Apache Flink with OVERaggregates. However, even though Apache Flink’s OVERaggregates can be partitioned over many workers, they do not scale well with increasing window size and increasing event throughput, as every new event triggers the recalculation of the aggregation function, and its computational cost is proportional to the window size, see Figure 5.
Figure 5. Incremental compute engines transform rolling aggregation computations from an O(N) computational complexity to an O(1) computational complexity problem.
Incremental views solve the scalability challenge by avoiding full recomputation of aggregations when a new event arrives. Instead, they reuse the previously computed value and apply only the changes introduced by new or removed events. As a result, the work performed is proportional to the input/output changes, not the window size.
Feldera is the leading open-source incremental compute engine. You write rolling aggregations in SQL and it is transpiled into rust code for high performance. Feldera supports incremental view maintenance through its streaming engine DBSP (DataBase inspired by Signal Processing). DBSP implements incremental views using Z-sets, a generalization of relational sets that track not only which elements are present, but also how their counts change over time. In a traditional relational set, an element either exists (count = 1) or does not exist (count = 0). In a Z-set, each element has an integer count that can be positive, zero, or negative:
Positive counts represent insertions (adding events).
Negative counts represent deletions (removing events).
This allows Z-sets to represent deltas, the net change between two states, without storing the full state at each step.
Shift Right with Parallel Pushdown Aggregations in Hopsworks/RonDB
Many organizations try to avoid introducing stream processing to their architecture, as it introduces operational complexity and cost. If your data volume and velocity are not at hyperscaler levels, you can stream raw events into Hopsworks online database, RonDB, and when clients request aggregate computations, they are computed on-demand in parallel in the database nodes in RonDB.
Figure 6. RonDB pushes down aggregation computation to data nodes that compute them in parallel.
Aggregations can be combined (and pushed down to database nodes) with other SQL operations in the same query, such as index scans, LIMITs, and projections. Other feature computation engines, that are built on DynamoDB and Redis, can also compute aggregations on-demand, but they are limited by the database capabilities . For example, in Redis, clients have to read all the columns and all the rows in the table (as it does not support predicate pushdown or projection pushdown). In contrast, in RonDB, only the columns returned to the client from the table are read and predicates (e.g., only computed for users based in the ‘EU’) are applied in the database kernel before the aggregation is computed. The result is orders of magnitude lower latency and higher throughput for selective queries.
RonDB Pushdown Aggregation Performance Benchmark
2 Data nodes: c6i.4xlarge (16c32g)
1 Benchmark Client: c6i.8xlarge (32c64g)
RonDB Pushdown Aggregation via NdbAggregator C++ API
Connections: 4 Ndb_cluster_connection objects, 32 workers distributed round-robin (8 per connection)
Timing: std::chrono::steady_clock with nanosecond precision per query
Warmup: 5 seconds discarded before each measured 15-second window
Data: Random values — val1/val2 BIGINT, val3/val4 DOUBLE, filter_date DATE (2024-01-01 to 2024-12-31)
Five tables with identical schema, varying in row count: 100, 500, 10K, 100K, 500K.
CREATE TABLE bench_qps_ {N} (
id INT NOT NULL AUTO_INCREMENT,
val1 BIGINT NOT NULL , -- random 0–999999
val2 BIGINT NOT NULL , -- random 0–999999
val3 DOUBLE NOT NULL , -- random 0–1000
val4 DOUBLE NOT NULL , -- random 0–1000
filter_date DATE NOT NULL , -- random in 2024-01-01 to 2024-12-31
padding VARCHAR ( 100 ) NOT NULL , -- 80-char random string
PRIMARY KEY (id),
INDEX idx_date (filter_date) -- ordered index, used by Q4
) ENGINE = NDB;
val1, val2: Integer columns for SUM() aggregation
val3, val4: Floating-point columns for AVG(), MIN(), MAX(), and expression aggregation (val1 * val3, val3 + val4)
filter_date: Uniformly distributed across 2024; Q4 filters >= '2024-07-01' (~50% selectivity)
padding: Simulates realistic row width (~130 bytes per row)
SELECT COUNT ( * ), SUM (val1), SUM (val3), COUNT (val3) FROM table
Q2: Multi-Aggregate (6 aggregations)
At 10K rows, PA continues to perform well, but other engines would transfer all rows to compute 6 aggregates, while PA computes them at the data node in a single pass.
SELECT COUNT ( * ), SUM (val1), SUM (val2), SUM (val3), COUNT (val3), MIN (val4), MAX (val4) FROM table
Q3: Expression Aggregation
Arithmetic expressions (val1 * val3, val3 + val4) are computed directly on data nodes. This is an example of a large table where scan I/O dominates over computation.
SELECT COUNT ( * ), SUM (val1 * val3), SUM (val3 + val4), COUNT (val3 + val4) FROM table
Q4: Filtered Aggregation (Index Scan)
SELECT COUNT ( * ), SUM (val1), SUM (val3), COUNT (val3) FROM table WHERE filter_date >= '2024-07-01'
Benchmark Results for Q1-Q4
Table 1 below shows experimental performance for our benchmark with Q1-Q4, while varying the number of rows each aggregation is computed over. For each query and number of rows per aggregation, we measure throughput in queries per second (QPS) and latency.
Table 1. Benchmark results for Queries 1-4, varying the number of rows aggregations are computed over, and measuring latency and throughput in queries per second (QPS).
From the results with only 2 data nodes (RonDB can scale up to 144 data nodes), you can see that the latency for Q1-Q3 increases when the number of rows for an aggregation reaches 10,000. For Q4, however, latency stays very low (1.9 milliseconds), as this query includes an index scan that pushes both range filtering and aggregation to the data nodes, eliminating row transfer entirely, and keeping latency low. The implication is that if you have aggregations that also include filters, you can scale pushdown aggregations to much higher performance on the same hardware.
Comparative

[truncated]
