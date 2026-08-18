---
source: "https://greptime.com/blogs/2026-08-11-observability-three-pillars-history"
hn_url: "https://news.ycombinator.com/item?id=49352731"
title: "The three pillars of observability were never designed"
article_title: "Observability Is Converging. Humans Aren't the Only Ones Querying It Anymore | Greptime"
image: "https://greptime.com/blogs/2026-08-11-observability-three-pillars-history/cover.png"
author: "xzhuang1984"
captured_at: "2026-08-18T21:14:09Z"
capture_tool: "hn-digest"
hn_id: 49352731
score: 1
comments: 0
posted_at: "2026-08-18T21:08:03Z"
tags:
  - hacker-news
  - translated
---

# The three pillars of observability were never designed

- HN: [49352731](https://news.ycombinator.com/item?id=49352731)
- Source: [greptime.com](https://greptime.com/blogs/2026-08-11-observability-three-pillars-history)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T21:08:03Z

## Translation

タイトル: 可観測性の 3 つの柱は決して設計されていなかった
記事のタイトル: 可観測性は収束しつつあります。もはやそれを疑問に思っているのは人間だけではありません |グレプタイム
説明: メトリクス、ログ、トレースを 1 つの列型ストアに入れることは複数のベンダーによって行われ、2026 年にはすべてのベンダーがエージェントに接続されました。本当の疑問は、さらに 1 層下にあります。エージェントが一流の消費者になったら、データベース自体を変更する必要があるのでしょうか、また、どこまで変更する必要があるのでしょうか。

記事本文:
コンテンツへスキップ メイン ナビゲーション 製品 GreptimeDB OSS オープンソースの可観測性データベース — メトリクス、ログ、トレース用の 1 つのエンジン。
インフラストラクチャ上の可観測性をスケールします。
エッジとクラウドの統合ソリューション 自動車会社の特定のビジネス シナリオと密接に連携する、きめ細かく調整されたソリューション。
スケーラブルな取り込み、効率的なストレージ、リアルタイム分析を備えた EMQX および MQTT ワークロード向けに構築されています。
GreptimeDB とシームレスに連携する統合、ツール、パートナー ソリューションを調べてください。
ドキュメントのダウンロード イベントの比較 ブログ 価格 メイン ナビゲーション お問い合わせ エンジニアリング • 2026-08-11 検索 K 可観測性は収束しています。もはやそれを疑問に思っているのは人間だけではありません
メトリクス、ログ、およびトレースを 1 つの列指向ストアに配置し、それらをまとめてクエリし、実稼働負荷に耐えることは、複数のベンダーによって行われています。そして 2026 年には、全員がエージェントに接続されました。本当の疑問はさらに 1 層下にあります。エージェントが一流の消費者になったら、データベース自体を変更する必要がありますか?また、どこまで変更する必要がありますか?
「統一された観察可能性」は 8 年近く話題になっていますが、一般的な印象は、それが製品ではなくアイデアにとどまっているということでした。その印象はもう時代遅れです。 3 つのシグナルすべてを 1 つの列型ストアに入れ、それらをまとめてクエリし、実稼働規模で実行するシステムは現在存在します。
2026 年までに、これらのシステムはエージェントにも接続されるようになります。したがって、興味深い問題は、誰かがエージェントに気づいたかどうかではありません。誰もがそうしました。それが次に起こることです。エージェントが人間と並んで一流の消費者になったとき、変更はインターフェースで止まるのでしょうか、それともその下のデータベースに到達するのでしょうか?
統合ストレージはもはや提案ではありません
SigNoz は、OTel ネイティブの単一の ClickHouse インスタンスにメトリクス、ログ、トレースを保存します

全体を通して。 ClickStack はより明確なケースです。 ClickHouse は 2025 年 3 月に HyperDX を買収し、同年 5 月に ClickStack を出荷しました。これにより、OTel Collector、ClickHouse、クエリ エクスペリエンスが 1 つのオープンソース スタックにパッケージ化され、3 つのシグナルすべてが一緒に調査されます。
Honeycomb はより早くそこに到達し、任意に幅広く構造化されたイベントに賭け、その下に独自の円柱状のストアを構築しました。 Charity Majors は 2024 年に次のような観察を行い、それ以来何度も引用されてきました。最近設立された可観測性スタートアップ企業は、統合ストレージ モデル、広範な構造化イベント、OTel ネイティブ、通常はカラム型データベースに収束しました。彼女の考えは、もう誰も安価な Datadog を構築しないということでした。彼らはより安価なハニカムを構築します。
彼女は後にその楽観的な考えを一部撤回した。 2025年に彼女は「ピラーは嘘だ」と書き、シグナルは専門用語であるのに対し、ピラーはマーケティング用語であると主張した。 2026 年 7 月までに批判は激化しました。カラムナ型ストレージは、2019 年以降に構築された可観測性バックエンドの標準に近づいていますが、それらの製品の多くは依然として 3 本柱モデルを出荷しており、「Datadog の一種だが安価」として販売されています。彼女がイライラしているのは、これらのベンダーの方が優れたアーキテクチャを備えているにもかかわらず、そう言わないことを選択していることです。ストレージ エンジンを交換してもパラダイムは変わりませんでした。
それでも、1 つの柱状ストアにある 3 つのシグナルは、ストレージ層とエクスペリエンス層で解決されつつある問題、おそらくコモディティ化された問題になりつつあります。誰も統合ストレージを構築していないという主張は、もはや現実と一致しません。
では、なぜここがゴールラインではないのでしょうか？
エンジニアリング上の 2 つの選択肢と、もはや当てはまらない 1 つの仮定
これらのシステムには 3 つの特徴があります。そのうちの 2 つは、議論する価値のあるエンジニアリング上の選択です。 3つ目は仮定です。
1つ目はその下の汎用エンジンです。 ClickHouse は最先端のサービスです

これらは、この世代の統合システムの共通の基盤ですが、注目すべき例外として Honeycomb の社内円柱型ストアがあります。これは汎用 OLAP 用に設計された列型ストアであり、それが得意ですが、その設計は一般的な分析から始まります。 3 つの信号のアクセス パターンはそうではありません。メトリクスには集計とダウンサンプリングが必要です。ログには全文検索が必要です。トレースでは、ID およびツリー走査によるポイント検索が必要です。
ClickHouse は、TimeSeries テーブル エンジン、PromQL、および全文検索を追加して、そのギャップ自体を埋めてきました。全文検索は現在一般提供されていますが、ClickHouse は BM25 スタイルの関連性スコアリングを実装しておらず、専用の検索エンジンを置き換えるのではなく、トークンレベルのフィルタリングを高速化することを目的としていることが明示されています。 TimeSeries と PromQL はまだ実験段階です。 PromQL の作品を紹介する投稿には、「ここにドラゴンがいます」という行が含まれています。
汎用の円柱型ストアは明らかに可観測性を実現でき、それをうまく実行します。しかし、これらの信号固有の機能がすべて避けられず、最終的にはデータベース内に存在する必要がある場合、可観測性ワークロードが初日から設計目標であった場合、統合エンジンはどのようにあるべきでしょうか?
2 番目の選択肢は、統合を別のレイヤーに置きます。 Grafana の LGT​​M スタックでは、Loki がログを保持し、Tempo がトレースを保持し、Mimir がメトリクスを保持し、Grafana が視覚化を処理します。統合されるのはエクスペリエンス層とコントロール層です。ストレージ エンジンは独立したままになります。これは明確なコストと利点を備えた一貫した選択です。各信号は独自のスケジュールに従って進化するため、信号間の相関関係はより上位で発生する必要があります。
3 番目の共通の特徴は、最も見落とされやすいものです。これらのシステムはすべて、個人向けに設計されています。 SigNoz エクスプローラー、ClickStack 検索エクスペリエンス、および Honeycomb クエリ インターフェイスはすべて

画面の前に座って直線的にクエリを実行し、ダッシュボードを見つめながら考え、3 つの信号間の相関関係を頭の中で考えているエンジニアを想像してください。
最初の 2 つはエンジニアリング上のトレードオフです。 3 つ目はより深いものです。なぜなら、それが設計全体の前提となる前提であり、2026 年には業界全体でそれが書き換えられようとしているからです。
これを理解するには、どのように分割されるかを見てください。
「三本の柱」という言葉は、自然の法則のように聞こえてしまうほど簡単に口から消えてしまいます。誰もデザインしたわけではありません。メトリクス、ログ、トレースは、さまざまな問題領域や技術的パスに沿って独立して進化し、後に初めて単一の可観測性フレームワークにグループ化されました。
指標が最初にありました。 RRDtool (1999 年) から Graphite (2008 年のオープンソース)、Prometheus (2012 年) まで、システムが現時点で健全であるかどうかが常に問題でした。ログが続いた。 Splunk (2003) と Elasticsearch (2010) は、実際に何が起こったのかという別の質問に答え、その答えは転置インデックスの形をとりました。分散トレースが後から登場したのは、マイクロサービスが普及するにつれて問題自体が差し迫ったものになったことが主な理由です。Google はコードを公開せずに 2010 年に Dapper 論文を公開し、2012 年 6 月に Twitter によってオープンソース化された Zipkin は、最初に使用可能なオープンソース実装でした。
図 1: それぞれが独自の問題領域に沿って成長し、後で 1 つのフレームワークにグループ化された 3 つのシグナル
これらは、3 つの独立した質問に対する 3 つの独立した回答です。それらを 1 つのシステムの 3 つのビューと呼ぶのは、過去を振り返って付けられたラベルです。そして、それらのエンジニアリング上の制約は大きく矛盾します。
2010 年代初頭に利用可能なテクノロジーを考えると、3 つすべてを適切に実行する 1 つのシステムを構築することは困難でした。ある信号に対して最適化したものはすべて、別の信号に負担をかけることになります。それらを分割する方が安価で、失敗する可能性が高い方法でした。

ああ、成功する。それを近視眼的だと言うのは後知恵だ。
その後、ビジネス上のインセンティブが技術的な分裂を強化しました。 Splunk はログ検索から成長し、Datadog はインフラストラクチャ監視から始まり、New Relic は APM でその名を轟かせました。既存企業にとって、統合は価格決定力が最も強い土地を手放すことを意味するため、分割を維持することは外堀だった。購入者側も細分化されており、インフラストラクチャの監視は SRE チームとプラットフォーム チームに属し、APM は開発者に属し、ログは多くの場合セキュリティ チームまたはデータ チームに属していました。予算はすでに 3 つの異なるポケットに収まっていました。
その時代の制約が業界のデフォルトの世界観になりました。
統一論は8年前に遡る
3 信号フレーミングの定義に貢献した人の 1 人は、最初にそれに疑問を呈した人でもあります。
2017 年 2 月、その年の Distributed Tracing Summit に出席した後、Peter Bourgon はベン図を使用して 3 つの重複箇所を整理し、「メトリクス、トレース、ロギング」を執筆しました。この図は 3 つの柱のフレームワークの概念的な起源となりましたが、彼が望んでいたのは、部屋いっぱいの人々に共通の語彙を提供することでした。より儀式的な表現は後に登場し、ベンダーはこのフレームワークを使用して市場を切り拓きました。
18か月後の2018年8月、ブルゴン氏は「可観測性シグナル」を執筆した。今度は彼は逆の質問をした。メトリクス、トレース、ロギングが同じ観測データの 3 つの消費パターンにすぎない場合、原理的には 1 つのシステム (彼の言葉を借りれば超システム) を構築できます。このシステムは、正面玄関で生のイベントを取得し、形状ごとに多重化を解除します。
ここでの詳細の 1 つは、よく誤解されているため、はっきりと述べておく価値があります。 Bourgon が念頭に置いていた目的地は、目的主導型のバックエンドでした。彼は、統合された書き込みパスと統合された読み取りモデルについて説明していましたが、すべての

ignal は同じ物理データベース内に存在します。今日、多くの「統合ストレージ」議論が彼を祖先だと主張していますが、彼の提案はより抑制されたもので、取り込みとアクセスを共有し、その下に専用のバックエンドを備えていました。
同年 12 月、Ben Sigelman は KubeCon North America で「三本の柱、ゼロの答え: 可観測性を再考する必要がある」と題した講演を行いました。彼は Dapper 論文の共著者であり、OpenTracing を共同設立し、2015 年には Lightstep を共同設立していたため、これらの柱を構築し、その上に製品を出荷しました。彼の批判は具体的でした。メトリクスはカーディナリティによって制限され、ディメンションが増加するにつれてコストが遠ざかっていきます。ログ記録の請求額は、トランザクション レート、マイクロサービスの数、ネットワークとストレージのコスト、保持期間を掛けたものになり、その製品は制御不能になります。より根本的には、3 つのシグナルはすべて「単なるビット」であり、これらを 3 つの個別の請求書を持つ 3 つの個別の柱として積み重ねることは拡張できません。
2019 年 5 月に、OpenCensus と OpenTracing は OpenTelemetry に統合されました。統合が実際に標準層に到達したのはこれが初めてでしたが、重要な注意点が 1 つありました。 OTel はコレクションとプロトコルを統合し、ストレージとクエリを元の場所に残しました。上流のパイプは接続されました。下流では、3 つのプールは 3 つのプールのままでした。その最も過小評価されている遺産は、まったく別のもの、つまりセマンティック規約です。 http.request.method をどのように呼び出すか、データベース呼び出しのスパンにどの標準フィールドが属するか、これらすべてが業界のコンセンサスとして確立されています。当時、それはフィールド名を調整する事務作業のように見えましたが、その見返りは、まさにこの作品が目指しているエージェントのシナリオで、数年後に届きました。 GenAI のセマンティック規約は、同じ作業をエージェント時代にも拡張します。 5月にそれらを層ごとに調べました。
lで

2023 年に、Charity Majors は Observability 2.0 を提案しました。これは、真実の 1 つの情報源、任意に幅広く構造化されたイベント、派生ビューとしてのメトリクスとトレースを備えています。これは 5 年前のブルゴンの考え方と同じであり、生の出来事を主要なものとし、3 つの柱を二次的なものとしています。
このパラダイムには批判者もいます。最も一般的な反対意見はコストです。ワイド イベントでは高いカーディナリティと高次元性が完全に維持されるため、リクエストあたりのデータ量は単一の柱を超え、メトリクスは大規模な集計ワークロードに対して依然として大幅に安価であるため、短期的に置き換えることは困難です。 「2.0」というラベルですら論争があり、メジャーズさんはその枠組みが特に好きではないと語った。
2026 年の夏までに誰もがエージェントがやってくるのを目にしていました
この記事がここで終わった場合、次の文は明らかで、これらのシステムはすべて人間のために構築されたものであり、エージェントについてはまだ誰も考えていないということになるでしょう。
2026 年には、これが業界全体の動きになりました。 SigNoz は 5 月にエージェント ネイティブの可観測性を発表し、ホスト型およびオープンソースの MCP サーバー、製品内の AI チームメイト、コーディング エージェントに SigNoz の操作方法を教える一連のエージェント スキルを出荷しました。 ClickHouse は、Ope で ClickStack MCP サーバーと AI ノートブックを開始しました

[切り捨てられた]

## Original Extract

Putting metrics, logs, and traces into one columnar store has been done, by more than one vendor, and in 2026 all of them connected agents. The real question sits one layer further down: once agents become first-class consumers, does the database itself have to change, and how far?

Skip to content Main Navigation Product GreptimeDB OSS Open-source observability database — one engine for metrics, logs, and traces.
Scale observability on your infrastructure.
Edge-Cloud Integrated Solution A finely tailored solution that aligns closely with the specific business scenarios of automotive companies.
Built for EMQX and MQTT workloads with scalable ingestion, efficient storage, and real-time analytics.
Explore integrations, tools, and partner solutions that work seamlessly with GreptimeDB.
Documentation Download Events Compare Blog Pricing Main Navigation Contact us Engineering • 2026-08-11 Search K Observability Is Converging. Humans Aren't the Only Ones Querying It Anymore
Putting metrics, logs, and traces into one columnar store, querying them together, and holding up under production load: that has been done, by more than one vendor. And in 2026, all of them connected agents. The real question sits one layer further down. Once agents become first-class consumers, does the database itself have to change, and how far?
"Unified observability" has been a talking point for close to eight years, and the common impression was that it remained an idea rather than a product. That impression is out of date. Systems that put all three signals into one columnar store, query them together, and run at production scale exist today.
By 2026, those systems have also connected agents. So the interesting question is not whether anyone noticed agents; everyone did. It is what happens next. When agents become first-class consumers alongside humans, does the change stop at the interface, or does it reach the database underneath?
Unified storage is no longer a proposal ​
SigNoz stores metrics, logs, and traces in a single ClickHouse instance, OTel-native throughout. ClickStack is the clearer case. ClickHouse acquired HyperDX in March 2025 and shipped ClickStack that May, packaging the OTel Collector, ClickHouse, and a query experience into one open-source stack where all three signals are explored together.
Honeycomb got there earlier, betting on arbitrarily wide structured events and building its own columnar store underneath. Charity Majors made an observation in 2024 that has been quoted many times since: the observability startups founded recently converged on a unified storage model, wide structured events, OTel-native, usually on a columnar database. Her framing was that nobody builds a cheaper Datadog anymore; they build a cheaper Honeycomb.
She later walked back part of that optimism. In 2025 she wrote "The pillar is a lie," arguing that signal is a technical term while pillar is a marketing one. By July 2026 the criticism had sharpened . Columnar storage has become close to standard for observability backends built after 2019, yet many of those products still ship the three-pillar model and still sell themselves as some flavor of "Datadog, but cheaper." What frustrates her is that these vendors have the better architecture and choose not to say so. Swapping the storage engine did not change the paradigm.
Even so, three signals in one columnar store is becoming a solved problem at the storage and experience layers, arguably a commoditized one. Claiming that nobody has built unified storage no longer matches reality.
So why isn't this the finish line?
Two engineering choices, and one assumption that no longer holds ​
These systems share three traits. Two of them are engineering choices worth arguing about. The third is an assumption.
The first is a general-purpose engine underneath. ClickHouse is among the most common foundations for this generation of unified systems, with Honeycomb's in-house columnar store as the notable exception. It is a columnar store designed for general-purpose OLAP, and it is good at that, but general analytics is where its design starts. The access patterns of the three signals are not. Metrics want aggregation and downsampling. Logs want full-text search. Traces want point lookups by ID and tree traversal.
ClickHouse has been closing that gap itself, adding a TimeSeries table engine, PromQL, and full-text search. Full-text search is GA now, though ClickHouse is explicit that it does not implement BM25-style relevance scoring and is meant to accelerate token-level filtering rather than replace a dedicated search engine. TimeSeries and PromQL are still experimental; the post introducing the PromQL work includes the line "there are dragons here."
A general-purpose columnar store can clearly do observability, and it does it well. But if every one of these signal-specific capabilities is unavoidable and eventually has to live inside the database, what should a unified engine look like if observability workloads were the design target from day one?
The second choice puts unification at a different layer. In Grafana's LGTM stack, Loki holds logs, Tempo holds traces, Mimir holds metrics, and Grafana handles visualization. What gets unified is the experience and control layer; the storage engines stay independent. That is a coherent choice with clear costs and benefits. Each signal evolves on its own schedule, and cross-signal correlation has to happen higher up.
The third shared trait is the easiest to miss: every one of these systems was designed for a person. The SigNoz explorer, the ClickStack search experience, and the Honeycomb query interface all assume an engineer sitting in front of a screen, querying linearly, thinking while staring at a dashboard, holding the correlation between three signals in their own head.
The first two are engineering trade-offs. The third runs deeper, because it is the premise the whole design rests on, and 2026 is rewriting it across the industry.
To understand this, look at how it split apart ​
The phrase "three pillars" rolls off the tongue so easily that it sounds like a law of nature. Nobody designed it. Metrics, logs, and traces evolved independently along different problem domains and technical paths, and were grouped into a single observability framework only later.
Metrics came first. From RRDtool (1999) to Graphite (open-sourced 2008) to Prometheus (2012), the question was always whether the system is healthy right now. Logs followed. Splunk (2003) and Elasticsearch (2010) answered a different question, what actually happened, and the answer took the shape of an inverted index. Distributed tracing came later, mostly because the problem itself only became pressing as microservices spread: Google published the Dapper paper in 2010 without releasing the code, and Zipkin, open-sourced by Twitter in June 2012, was the first usable open-source implementation.
Figure 1: three signals, each grown along its own problem domain, grouped into one framework only later
These are three independent answers to three independent questions. Calling them three views of one system is a label applied in retrospect. And their engineering constraints conflict sharply:
Given the technology available in the early 2010s, building one system that did all three well was difficult. Whatever you optimized for one signal became a tax on another. Splitting them was the cheaper path and the one more likely to succeed. Calling that shortsighted is hindsight.
Business incentives then reinforced the technical split. Splunk grew out of log search, Datadog started with infrastructure monitoring, and New Relic made its name in APM. For incumbents, staying split was the moat, since unifying meant giving up the ground where their pricing power was strongest. The buyer side was fragmented too: infrastructure monitoring belonged to SRE and platform teams, APM to developers, logs often to security or data teams. The budget already sat in three different pockets.
A constraint of that era became the industry's default worldview.
The unification argument goes back eight years ​
One of the people who helped define the three-signal framing was also among the first to question it.
In February 2017, after attending that year's Distributed Tracing Summit, Peter Bourgon wrote "Metrics, tracing, and logging," using a Venn diagram to sort out where the three overlap. That diagram became the conceptual origin of the three-pillars framework, though what he wanted was to give a roomful of people a shared vocabulary. The more ceremonial phrase came later, and vendors used the framework to carve up their markets.
Eighteen months later, in August 2018, Bourgon wrote "Observability signals." This time he asked the question in reverse. If metrics, tracing, and logging are just three consumption patterns for the same observational data, then in principle you could build one system, an über-system in his words, that takes raw events at the front door and de-muxes them by shape.
One detail here is frequently misread and worth stating plainly. The destinations Bourgon had in mind were purpose-driven backends. He was describing a unified write path and a unified read model, not a requirement that every signal live in the same physical database. Plenty of "unified storage" arguments today claim him as a forefather, but his proposal was more restrained: shared ingestion and access, with purpose-built backends underneath.
That December, Ben Sigelman gave a talk at KubeCon North America titled "Three Pillars, Zero Answers: We Need to Rethink Observability." He co-authored the Dapper paper, co-founded OpenTracing, and had co-founded Lightstep back in 2015, so he had built these pillars and shipped a product on top of them. His criticism was concrete. Metrics are bounded by cardinality, and cost gets away from you as dimensions multiply. The logging bill is transaction rate times number of microservices times network and storage cost times retention window, and that product runs out of control. More fundamentally, all three signals are "just bits," and stacking them as three separate pillars with three separate invoices does not scale.
In May 2019, OpenCensus and OpenTracing merged into OpenTelemetry . That was the first time unification actually landed at the standards layer, with one caveat that matters. OTel unified collection and protocol, and left storage and query where they were. The upstream pipes were connected; downstream, the three pools stayed three pools. Its most underrated legacy is something else entirely: the semantic conventions. What http.request.method should be called, which standard fields belong on a span for a database call, all of it hardened into industry consensus. At the time it looked like the clerical work of aligning field names, and the payoff arrived years later, in exactly the agent scenario this piece is building toward. The GenAI semantic conventions extend that same work into the agent era; we went through them layer by layer in May.
In late 2023, Charity Majors proposed Observability 2.0: one source of truth, arbitrarily wide structured events, with metrics and traces as derived views. That is the same line of thinking as Bourgon's five years earlier, raw events as primary and the three pillars as secondary.
The paradigm has its critics. The most common objection is cost. Wide events preserve high cardinality and high dimensionality in full, so the data volume per request exceeds any single pillar, and metrics remain considerably cheaper for large-scale aggregation workloads, which makes them hard to displace in the near term. Even the "2.0" label is contested, and Majors has said she is not particularly fond of the framing.
By summer 2026, everyone had seen agents coming ​
If this article stopped here, the obvious next sentence would be that all of these systems were built for humans and nobody has thought about agents yet.
In 2026 this became a collective industry move. SigNoz announced agent-native observability in May, shipping a hosted and an open-source MCP server, an in-product AI teammate, and a set of Agent Skills that teach coding agents how to work with SigNoz. ClickHouse launched the ClickStack MCP server and AI Notebooks at Ope

[truncated]
