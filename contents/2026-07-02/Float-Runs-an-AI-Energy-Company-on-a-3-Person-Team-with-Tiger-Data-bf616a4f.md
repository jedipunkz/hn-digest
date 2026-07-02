---
source: "https://www.tigerdata.com/blog/how-float-runs-ai-energy-company-3-person-team-tiger-data"
hn_url: "https://news.ycombinator.com/item?id=48755265"
title: "Float Runs an AI Energy Company on a 3-Person Team with Tiger Data"
article_title: "How Float Runs an AI Energy Company on a 3-Person Team with Tiger Data | Tiger Data"
author: "nreece"
captured_at: "2026-07-02T01:39:56Z"
capture_tool: "hn-digest"
hn_id: 48755265
score: 1
comments: 0
posted_at: "2026-07-02T01:23:05Z"
tags:
  - hacker-news
  - translated
---

# Float Runs an AI Energy Company on a 3-Person Team with Tiger Data

- HN: [48755265](https://news.ycombinator.com/item?id=48755265)
- Source: [www.tigerdata.com](https://www.tigerdata.com/blog/how-float-runs-ai-energy-company-3-person-team-tiger-data)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T01:23:05Z

## Translation

タイトル: Float が Tiger データを使用して 3 人のチームで AI エネルギー会社を経営
記事のタイトル: Float が Tiger データを使用して 3 人のチームで AI エネルギー会社を経営する方法 |トラのデータ
説明: デンマークのスタートアップ Float が、Tiger Data を使用して 1Hz スマート メーター データの 99.3% 圧縮を達成し、数百世帯の家電レベルのエネルギー分析を強化した方法。

記事本文:
Float が Tiger データを使用して 3 人チームで AI エネルギー会社を経営する方法 |タイガーデータ製品 製品
スタートアップと企業向けの堅牢で弾力性のあるクラウド プラットフォーム
オンプレミス、エッジ、プライベート クラウド向けの自己管理型 TimescaleDB
Postgres 上の時系列、リアルタイム分析とイベント
Postgres でのベクトルおよびキーワード検索
Tiger Data ニュースレターを購読する
2026 (c) Timescale, Inc.、d/b/a Tiger Data。無断転載を禁じます。
Inductive Automation 2026 のゴールド パートナー (c) Timescale, Inc.、d/b/a Tiger Data。無断転載を禁じます。
Float が Tiger データを使用して 3 人チームで AI エネルギー会社を経営する方法
04 浮体エネルギーデータスタック
デンマークのスタートアップ企業は、1Hz スマート メーター データの 99.3% 圧縮を達成し、数百の家庭のリアルタイムの家電レベルのエネルギー分析を強化しています。
Float はデンマークの AI およびエネルギーのスタートアップ企業で、数百の家庭から 1Hz のスマート メーター データを収集し、独自の ML モデルを使用してそれを家電製品ごとの消費量に分解します。システム全体は、アーキテクチャ上の 1 つの賭けに依存しています。それは、時系列データベースの圧縮が、ストレージの経済性を大規模に機能させるのに十分な高さであるということです。共同創設者の Jens Brandt Nellegaard (CEO) と Victor Grabow (CTO) は、すべての主要な時系列データベースをどのように評価したか、90% の圧縮が実行可能性の厳しい下限であった理由、Tiger Data で 99.3% に達したときに何が起こったかを共有します。
ほとんどの人は、個々のアプライアンスの稼働コストがいくらなのかわかりません。デンマークでは、ほとんどのエネルギープロバイダーが毎月の PDF 請求書を送信しており、それが顧客エクスペリエンスのすべてです。市場ではここ 10 ～ 15 年間、意味のあるイノベーションは見られませんでした。詐欺と透明性の欠如があまりにも蔓延しているため、デンマーク政府は2026年1月から新しい消費者保護規制を導入しました。
この問題を解決するための基礎となるデータは実際には e

存在する。ヨーロッパのほとんどのスマート メーターには、家庭の総電力負荷を 1 秒の分解能で出力できる標準化された消費者インターフェイスが搭載されています。しかし、総負荷自体はあまり興味深いものではありません。消費者が必要としているのは、個々の機器ごとの内訳、つまりどの機器が電力を無駄にしているのか、どの機器がピーク価格で稼働しているのか、どの機器がヒューズの容量制限に近づいているのかを把握することです。エネルギー分解に関する学術研究は 1992 年に遡ります。商業的に拡張可能な方法でそれを解決した人は誰もいませんでした。
Float は、このギャップを埋めるために 3 つのものを構築しました。ヨーロッパのスマート メーターの消費者ポートに接続する独自のハードウェア モジュール、生の波形から家電レベルの消費量を分類する信号処理およびニューラル ネット パイプライン、プロアクティブな AI エネルギー エージェントを備えた消費者向けアプリです。このシステムは、1 世帯あたり 1 秒あたり約 15 件の測定値を、それぞれ 1Hz の分解能で収集します。 1 秒はハードフロアです。Victor 氏は次のように説明します。「1 分の分解能では、モデルは壊れてしまいます。なぜなら、私たちが追跡しているのは変化だからです。同じ 1 分以内にオンになっている家電製品が多すぎると、それらを区別するのが非常に困難になります。」
Jens Brandt Nellegaard と Victor Grabow は 2022 年に Float を共同設立しました。約 3 年間の研究開発を経て、2024 年 12 月に概念実証を達成し、2025 年 12 月にエネルギープロバイダーのライセンスを確保し、現在、事前に精査した 350 の顧客にプライベート ベータ版を展開しています。会社にはたったの3人しかいません。
Float は、TimescaleDB 拡張機能を使用して Azure マネージド Postgres 上で開始されました。チームは早くから Azure クレジットを持っていたので、当時はそれが当然でした。しかし、Azure で利用できる Apache バージョンには圧縮が含まれていなかったため、これが問題となった。
どの顧客も 1 秒あたりおよそ 15 回の測定値を生成します。フロートサンプル

家庭に入る各相の電圧、周波数、および総負荷。 1,000 世帯の場合、継続的に 1 秒あたり 15,000 データ ポイントになります。圧縮がなければ、ストレージコストだけでビジネスモデルが崩壊してしまいます。 「当社は定額のサブスクリプション料金を提供するエネルギー会社です」と Jens 氏は言います。 「私たちはマークアップなしでスポット市場価格を 1 対 1 で決定します。ユーザーあたりのストレージ コストがサブスクリプションでサポートされる金額を超えると、経済性が崩壊します。」
TimescaleDB に落ち着く前に、InfluxDB も試しました。取り込みの問題が発生したため、SQL が必要になりました。資産中心のマイクロサービス プラットフォームを構築する 3 人チームの場合、独自のクエリ言語を必要とし、ドメイン間でデータを結合およびクエリする方法が制限されるデータベースを購入する余裕はありません。 - Victor Grabow 氏、CTO、Float
ストレージの問題に加えて、Float には連続的な集計が必要でした。デンマークの DSO は、決済データを 15 分の解像度で提供します。 Float は毎秒データを収集します。実際の電力料金請求書を生成し、それを送電網事業者の数値と比較するには、システムは生の 1Hz データを 15 分ウィンドウまで継続的に集計する必要があります。 TimescaleDB の完全な機能セットを持たないマネージド Postgres では、バッチ ジョブの作成と保守を意味し、ハードウェア、ML、およびライセンスを取得したエネルギー会社全体ですでに手薄になっていたチームにとって、インフラストラクチャのオーバーヘッドがさらに増加し​​ました。
私たちは市場にあるほぼすべての時系列データベースをテストしました。 Tiger Data が私たちのユースケースに最適なソリューションであると考えています。 - Victor Grabow 氏、CTO、Float
チームは時系列ソリューションを広範囲に調査し、圧縮機能と連続集計機能を通じて TimescaleDB を発見しました。圧縮には十分な説得力があったため、フル機能が必要になるのは時間の問題でした。 Tiger データを見つけたとき - the c

TimescaleDB の背後にある企業 - マネージド クラウド サービスによって道筋が明確になりました。
2 つの特徴が決定の原動力となりました。まず、圧縮。チームはユニットエコノミクスをモデル化しており、サブスクリプション モデルが機能するには時系列データを少なくとも 90% 圧縮する必要がありました。それ以下の場合、およびユーザーごとのストレージ コストは、定額料金でサポートできる範囲を超えます。 2 つ目は、継続的集計 - 新しいデータが到着するたびに段階的に更新されるマテリアライズド ビューです。 Float は集計を継続的に実行し、1Hz の測定値を 15 分の決済ウィンドウに変換し、電圧と周波数に関するしきい値ベースのアラートを計算し、停電を検出し、4 時間稼働しているオーブンにフラグを立てるなどの期間ベースの通知をトリガーします。連続集計は、バッチ ジョブやスケジュールされたパイプラインを使用せずに、これらすべてを処理します。
速度の問題だったので、Azure のフルマネージド サービスである Tiger Cloud を選択しました。私たちは迅速に稼働を開始し、インフラストラクチャ管理を完全にオフロードする必要がありました。当社の DevOps プラットフォームである Encore は、Google Cloud 上で一時的な環境を提供しており、Tiger Cloud のデータベース分岐はそのワークフローに自然に適合します。 - Victor Grabow 氏、CTO、Float
データはスマートメーターから始まります。 Float の IoT モジュールは、標準化された消費者向けインターフェイス ポートに接続し、各相の電圧、周波数、総負荷を 1Hz の解像度でキャプチャします。このモジュールは読み取り値を Azure IoT Hub に送信します。チームは、このハブをすべてのデバイスの安定した取り込みエンドポイントとして元の Azure セットアップから保持しています。
そこから、ブリッジ コネクタがストリームを Google Cloud に転送し、そこで Encore が Float のマイクロサービスをデプロイします。チームは、コストが高かったため、最終的に Azure Event Hub から移行しました。 Google Cloud ストリーミング サービスは、同じスループットを低コストで処理します。ストリーミングレイヤーバッチ

すべての世帯から毎秒受信する測定値をバッチごとに Tiger Data に挿入します。
Tiger Data は、生の 1Hz 時系列測定値を保存し、電圧スパイク、周波数変化、平均値と最大値の計算、停止検出、継続時間ベースのアプライアンス アラートなど、しきい値ベースのモニタリングのために継続的な集計を実行します。すべての生データは、プライベート ベータ段階を通じて ML トレーニングの目的で保持され、フリートの規模に応じて階層化ストレージが計画されます。
Float アプリは処理されたデータを読み取り、家電ごとのリアルタイムのエネルギー内訳を顧客に表示します。新規のお客様は、接続するとすぐに合計ワット数が表示されます。モデルが自宅の特定のパターンに基づいてトレーニングされるため、家電レベルの内訳にはおよそ 3 ～ 4 週間かかります。最上位のエージェント オーケストレーション層は、請求、オンボーディング、カスタマー サービス、プロアクティブな通知を処理します。忘れられたオーブンやアイロン、非効率な家電製品、ヒューズ制限に近づく危険な負荷状態にフラグを立てます。
Tiger Data では、Float の時系列データが 99.3% 圧縮されています。ビクターは直接こう言います。
ビジネス モデルを壊さないように、圧縮率を 90 年代後半にする必要がありました。それは素晴らしい結果でした。 - Victor Grabow 氏、CTO、Float
この数値により、低い圧縮率では不可能だった 3 つのことが可能になりました。
1,000 世帯で 1 秒あたり 15,000 データ ポイントの場合、非圧縮ストレージでは年間テラバイト規模の生の時系列データが生成されます。フロートは、スポット市場の電力価格を値上げなしの原価で顧客に渡します。収益は定額のサブスクリプション料金から得られます。ユーザーあたりのストレージ コストがその料金でサポートできる以上に上昇すると、モデル全体が崩壊します。 99.3% の圧縮では、そうではありません。サブスクリプションは余裕のあるインフラストラクチャをカバーしており、そのマージンは f として保持されます。

リートスケール。
ML トレーニングの完全なデータ保持
Float の分解モデルでは、各家庭の特定の機器のシグネチャを学習するために、世帯ごとに数週間分の 1Hz トレーニング データが必要です。圧縮率が低い場合、チームはモデルのトレーニング用にすべての生データを保持するか、ストレージ コストを維持するかの選択に直面することになります。 99.3% ではすべてが保持されます。プライベート ベータ フリート全体からのすべての生の 1Hz 読み取り値は ML パイプラインで利用でき、フリートがベータ段階を過ぎてスケールアップする場合にのみ階層化ストレージが計画されています。
バッチインフラストラクチャを使用しないリアルタイム請求
デンマークのグリッドは 15 分の決済ウィンドウで運営されています。 Float は毎秒データを収集します。連続集計はそのギャップを埋め、1Hz の測定値を DSO が請求書調整に必要とする 15 分間隔に変換します。デンマークのエネルギー価格はピーク時とオフピーク時で最大 80% 変動するため、集約の新鮮さが顧客にとって直接の価値となります。継続的な集計は新しいデータが到着するたびに段階的に更新されるため、Float のライブ電力料金は常に最新のものになります。つまり、スケジュールされたバッチ ジョブ、パイプラインのメンテナンス、遅延はありません。
認可されたエネルギー会社を経営する 3 人チーム
Float はデンマークでエネルギープロバイダーのライセンスを取得しています。これは、請求、顧客サービス、オンボーディング、規制順守など、従来のエネルギー会社が数十人を配置する運用上のオーバーヘッドを意味します。 Tiger Cloud のマネージド インフラストラクチャは、これを可能にするものの一部です。チームはデータベース操作、ストレージ プロビジョニング、または集約パイプラインを管理しません。このオーバーヘッドは処理されます。チームの規模について尋ねられたとき、ジェンスの答えはシンプルでした。「3 人です。我々には 3 人がいます...そしてエージェントの軍隊がいます。これが未来です。」
Float は今後 12 か月以内にシードラウンド、2 回の追加で 1,000 人のプライベート ベータ顧客を追加することを目標としています。

デンマークの電力網を完全にカバーするための最終ハードウェア バリアントと、さらに 2 か国への拡張。
次の主要な統合は EV 充電です。これは Tesla のテレメトリ API から始まり、安価な価格帯でのスマート充電を可能にします。さらに重要な理論は、1Hz の分解能で測定された家庭のフリートは、24 ～ 48 時間の遅延を伴う 15 分の分解能で動作するどのエネルギー会社よりも効率的にスポット市場で電力を取引できるということです。イェンス氏は次のように述べています。「最終的には、ホームを電力網の負担にするのではなく、電力網のパートナーにしようとしているのです。」
Float のスケールに応じて複雑になるアーキテクチャの決定は、圧縮率そのものではありません。それは、生の 1Hz 読み取り値、請求のための継続的な集計、ML パイプラインのトレーニング データ、異常検出クエリなど、すべてが単一の Tiger Data インスタンス上で実行されるということです。フリートが 350 軒から 1,000 軒以上に拡大しても、維持する分割アーキテクチャや調整するクエリ パスは必要ありません。データ モデルは変わりません。ただ大きくなるだけです。
ApexAnalytica が単一の Postgres インスタンス上でテレメトリ、トランザクションデータ、RAG の構築を実行する方法
オープン ソースの TimescaleDB が 1 年間の時間ごとの建物エネルギー データを単一のヒート マップ ビューでどのように強化するかをご覧ください。

[切り捨てられた]

## Original Extract

How Danish startup Float hit 99.3% compression on 1Hz smart meter data with Tiger Data, powering appliance-level energy analytics for hundreds of homes.

How Float Runs an AI Energy Company on a 3-Person Team with Tiger Data | Tiger Data Product Product
Robust elastic cloud platform for startups and enterprises
Self-managed TimescaleDB for on-prem, edge and private cloud
Time-series, real-time analytics and events on Postgres
Vector and keyword search on Postgres
Subscribe to the Tiger Data newsletter
2026 (c) Timescale, Inc., d/b/a Tiger Data. All rights reserved.
GOLD PARTNER WITH INDUCTIVE AUTOMATION 2026 (c) Timescale, Inc., d/b/a Tiger Data. All rights reserved.
How Float Runs an AI Energy Company on a 3-Person Team with Tiger Data
04 The Float Energy Data Stack
Danish startup achieves 99.3% compression on 1Hz smart meter data, powering real-time appliance-level energy analytics for hundreds of homes.
Float is a Danish AI and energy startup that collects 1Hz smart meter data from hundreds of homes and disaggregates it into per-appliance consumption using a proprietary ML model. The entire system depends on one architectural bet: that compression on their time-series database would be high enough to make the storage economics work at scale. Co-founders Jens Brandt Nellegaard (CEO) and Victor Grabow (CTO) share how they evaluated every major time-series database, why 90% compression was the hard floor for viability, and what happened when they hit 99.3% on Tiger Data.
Most people have no idea what their individual appliances cost to run. In Denmark, most energy providers send a monthly PDF bill, and that is the entire customer experience. The market has seen no meaningful innovation in 10 to 15 years. Fraud and a lack of transparency have been so widespread that the Danish government introduced new consumer protection regulation effective January 2026.
The underlying data to solve this problem actually exists. Most European smart meters have a standardized consumer interface that can output total household electrical load down to one-second resolution. But total load is not very interesting on its own. What consumers need is a breakdown by individual appliance - to understand which one is wasting electricity, which one is running at peak price, which one is approaching the capacity limit of a fuse. Academic research on energy disaggregation goes back to 1992. Nobody had solved it in a commercially scalable way.
Float built three things to close this gap: a proprietary hardware module that plugs into the consumer port on a European smart meter, a signal processing and neural net pipeline that classifies appliance-level consumption from the raw waveform, and a consumer-facing app with a proactive AI energy agent. The system collects roughly 15 measurements per second per household, each at 1Hz resolution. One second is the hard floor - Victor explains: "At one-minute resolution, the model would break, because what we track are the changes. If there are too many appliances turning on within the same minute, it would be very hard to differentiate them."
Jens Brandt Nellegaard and Victor Grabow co-founded Float in 2022. After nearly three years of R&D, they achieved a proof of concept in December 2024, secured their energy provider license in December 2025, and are now rolling out a private beta to 350 pre-vetted customers. The company has just three people.
Float started on Azure managed Postgres with the TimescaleDB extension. The team had Azure credits early, so it made sense at the time. But the Apache version available on Azure did not include compression, and that turned out to be a dealbreaker.
Every customer generates roughly 15 measurements per second. Float samples voltage, frequency, and total load across each phase entering the home. At 1,000 homes, that is 15,000 data points per second, continuously. Without compression, the storage cost alone would break their business model. “ We are an energy company with a flat-rate subscription fee,” says Jens. “We pass through the spot market price one-to-one with no markup. If storage cost per user exceeds what the subscription supports, the economics collapse. ”
We also tried InfluxDB before settling on TimescaleDB. We ran into ingestion issues, and we needed SQL. When you are a three-person team building an asset-centric microservice platform, you cannot afford a database that requires a proprietary query language and limits how you join and query data across domains. - Victor Grabow, CTO, Float
On top of the storage problem, Float needed continuous aggregates. The Danish DSO delivers settlement data at 15-minute resolution. Float collects data every second. To generate a live energy bill and compare it against the grid operator's numbers, the system needs to aggregate raw 1Hz data down to 15-minute windows constantly. On managed Postgres without TimescaleDB's full feature set, that meant writing and maintaining batch jobs - more infrastructure overhead for a team that was already stretched thin across hardware, ML, and a licensed energy company.
We tested pretty much every time-series database on the market. We think Tiger Data is the best solution for our use case. - Victor Grabow, CTO, Float
The team researched time-series solutions extensively and discovered TimescaleDB through the compression and continuous aggregate features. The compression was compelling enough that it was just a matter of time before they needed the full capability. When they found Tiger Data - the company behind TimescaleDB - the managed cloud service made the path clear.
Two features drove the decision. First, Compression . The team had modeled the unit economics and needed at least 90% compression on the time-series data for the subscription model to work. Anything below that and storage costs per user would exceed what the flat-rate fee could support. Second, Continuous Aggregates - materialized views that update incrementally as new data arrives. Float runs aggregations constantly, converting 1Hz readings to 15-minute settlement windows, calculating threshold-based alerts on voltage and frequency, detecting outages, and triggering duration-based notifications like flagging an oven that has been running for four hours. Continuous aggregates handle all of this without batch jobs or scheduled pipelines.
We chose Tiger Cloud, the fully managed service on Azure, because it was a question of speed. We needed to get up and running fast and offload infrastructure management entirely. Encore , our DevOps platform, provides ephemeral environments on Google Cloud, and Tiger Cloud's database branching fits naturally into that workflow. - Victor Grabow, CTO, Float
Data starts at the smart meter. Float’s IoT module plugs into the standardized consumer interface port and captures voltage, frequency, and total load across each phase at 1Hz resolution. The module sends readings to Azure IoT Hub, which the team kept from the original Azure setup as a stable ingestion endpoint for all devices.
From there, a bridge connector forwards the stream into Google Cloud, where Encore deploys Float’s microservices. The team moved off Azure Event Hub eventually because it was expensive. Google Cloud streaming services handle the same throughput at lower cost. The streaming layer batches incoming measurements from all households every second and inserts them per batch into Tiger Data.
Tiger Data stores the raw 1Hz time-series readings and runs continuous aggregates for threshold-based monitoring: voltage spikes, frequency changes, mean and max calculations, outage detection, and duration-based appliance alerts. All raw data is retained for ML training purposes through the private beta phase, with tiered storage planned as the fleet scales.
The Float app reads processed data to show customers their real-time energy breakdown per appliance. New customers see total wattage immediately on connection. Appliance-level breakdown takes roughly three to four weeks as the model trains on their home's specific patterns. The agentic orchestration layer on top handles billing, onboarding, customer service, and proactive notifications - flagging forgotten ovens and irons, inefficient appliances, and dangerous load conditions approaching fuse limits.
On Tiger Data, Float is seeing 99.3% compression on its time-series data. Victor puts it directly:
Compression needed to be in the high nineties range to not break our business model. So that was a great outcome. - Victor Grabow, CTO, Float
That number unlocked three things that would not have been possible at lower compression ratios.
At 15,000 data points per second across 1,000 homes, uncompressed storage would generate terabytes of raw time-series data per year. Float passes through the spot market electricity price to customers at cost with no markup. Revenue comes from a flat-rate subscription fee. If storage cost per user climbs above what that fee can support, the entire model collapses. At 99.3% compression, it does not. The subscription covers infrastructure with margin to spare, and that margin holds as the fleet scales.
Full Data Retention for ML Training
Float's disaggregation model needs weeks of 1Hz training data per household to learn each home's specific appliance signatures. At lower compression, the team would face a choice: retain all raw data for model training or keep storage costs viable. At 99.3%, they retain everything. All raw 1Hz readings from the entire private beta fleet are available for the ML pipeline, with tiered storage planned only as the fleet scales past the beta phase.
Real-Time Billing Without Batch Infrastructure
The Danish grid operates on 15-minute settlement windows. Float collects data every second. Continuous aggregates bridge that gap, converting 1Hz readings into the 15-minute intervals the DSO requires for bill reconciliation. Danish energy prices swing up to 80% between peak and off-peak hours, which makes the freshness of the aggregation directly valuable to customers. Because continuous aggregates update incrementally as new data arrives, Float's live energy bill is always current, i.e., no scheduled batch jobs, no pipeline maintenance, no lag.
A Three-Person Team Running a Licensed Energy Company
Float holds an energy provider license in Denmark. That means billing, customer service, onboarding, regulatory compliance - operational overhead that traditional energy companies staff with dozens of people. Tiger Cloud's managed infrastructure is part of what makes this possible. The team does not manage database operations, storage provisioning, or aggregation pipelines. That overhead is handled. When asked about team size, Jens's answer was simple: "Three. We have three people... and an army of agents. This is the future."
Float is targeting 1,000 additional private beta customers within the next 12 months, with a seed round, two additional hardware variants for full Danish grid coverage, and expansion into two more countries.
The next major integration is EV charging - starting with Tesla's telemetry API, enabling smart charging during cheap price windows. The bigger thesis is that a fleet of homes measured at 1Hz resolution can trade power on the spot market more efficiently than any energy company operating at 15-minute resolution with a 24 to 48-hour delay. As Jens puts it: "Ultimately we are trying to make the home not a burden for the grid, but a partner of the grid."
The architecture decision that compounds as Float scales is not the compression ratio itself. It is that everything runs on a single Tiger Data instance: the raw 1Hz readings, the continuous aggregates for billing, the training data for the ML pipeline, the anomaly detection queries. No split architecture to maintain, no query paths to reconcile as the fleet grows from 350 homes to 1,000 and beyond. The data model does not change - it just gets bigger.
How ApexAnalytica Runs Building Telemetry, Transactional Data, and RAG on a Single Postgres Instance
See how open source TimescaleDB powers a year of hourly building energy data in a single heat map view — rendered in und

[truncated]
