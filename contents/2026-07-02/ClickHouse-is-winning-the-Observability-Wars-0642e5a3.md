---
source: "https://matduggan.com/clickhouse-is-winning-the-observability-wars/"
hn_url: "https://news.ycombinator.com/item?id=48759650"
title: "ClickHouse is winning the Observability Wars"
article_title: ""
author: "emschwartz"
captured_at: "2026-07-02T11:27:00Z"
capture_tool: "hn-digest"
hn_id: 48759650
score: 2
comments: 0
posted_at: "2026-07-02T11:19:18Z"
tags:
  - hacker-news
  - translated
---

# ClickHouse is winning the Observability Wars

- HN: [48759650](https://news.ycombinator.com/item?id=48759650)
- Source: [matduggan.com](https://matduggan.com/clickhouse-is-winning-the-observability-wars/)
- Score: 2
- Comments: 0
- Posted: 2026-07-02T11:19:18Z

## Translation

タイトル: ClickHouse が可観測性戦争に勝利
説明: 過去約 10 年間、私の勤務時間のかなりの割合が可観測性について考えることに費やされてきました。この用語に馴染みのない方のために説明すると、「監視」というとそれほど高価に聞こえないように、現在では「可観測性」と呼ばれています。実際の仕事は、収集するという点で地味です。

記事本文:
Clickhouse が可観測性戦争で勝利
過去約 10 年間、私の勤務時間のかなりの割合は可観測性について考えることに費やされてきました。この用語に馴染みのない方のために説明すると、「監視」というとそれほど高価に聞こえないように、現在では「可観測性」と呼ばれています。実際の作業は、大量のログ、いくつかのメトリクス、いくつかのトレースを収集し、それらを人々に提供するという点で地味です。
私は基本的に自分の仕事が好きです。私たちが常に新しいアイデアやアプローチを試していることが気に入っています。物事がうまくいかないとき、答えはほとんどの場合データの中にあり、辛抱強く調べられる人が見つけてくれるのを待っているという事実が気に入っています。しかし、正直に申し上げたいのです。6 社の企業と、聞いたことのあるすべての可観測性プラットフォームと、おそらく聞いたことのないいくつかのプラットフォームにわたって、この仕事を 10 年間行ってきましたが、ログが仕事の最悪の部分でなくなったことはありません。私が始めたとき、それらは最悪の部分でした。彼らは今日の最悪の部分です。ロボットが立ち上がって私の頭を一掃するまで、永遠に彼らがこの仕事の最悪の部分を占めることになるだろうと私は完全に予想しています。
ログがひどい理由については以前に書いたので、全文は割愛して、短いバージョンを紹介します。
すべての開発者のログに対する期待は、syslog ボックスという単一の形成エクスペリエンスによって決まります。またはローカルで実行されているコンテナー。または、おそらく SSH 接続すべきではない運用サーバー上で tail -f を実行します。重要なのは、彼らのキャリアの初期の、優しい瞬間に、完璧な丸太の経験があったということです。 grep を実行すると、役立つものが返されました。彼らはそれを jq にパイプし、必要なものを正確に取得しました。
この経験は、ファーストキスと同等の観察可能性です。それはその後に起こるすべてのために彼らを台無しにします。
なぜなら、これがそのフラに関するものだからです

経験が浅い場合: システムが小さく、ボリュームがわずかで、クエリを実行している人がログ行を書いたのと同じ人であるため、機能します。スキーマのドリフトやカーディナリティの爆発はなく、ダッシュボードに期待するクロスチームの消費者もおらず、「収益イベント」グラフにギャップがある理由を尋ねる副社長もいません。
それから40のサービスがあります。今では四百人になりました。現在、ログは開発者だけでなく、火曜日以降に特定のユーザーが失敗したチェックアウトを調べる必要がある顧客サービスによっても消費されています。そしてデータ チームは、バックエンド エンジニアが誰にも言わずにリファクタリングしようとしているログ行の上にビジネス クリティカルなダッシュボードを密かに構築しています。そしてオンコール担当者は、午前 3 時に新しいクエリ言語を学習したくない、インデックス パターンについて考えたくない、検索バーだけが機能することを望んでいます。
つまり、量が膨大で、形状が一貫性がなく、クエリが予測不可能であるという技術的な問題があり、期待の問題の上にさらに深刻な問題が重なっています。開発者は即座にログを取得し、そのログに対して任意の操作を実行したいと考えていますが、スキーマにはコミットしません。一方、同じデータをあまり技術的に使用しない消費者は、ダッシュボードが永久に安定していて、UI が寛容で、全体が通常の製品のように感じられることを望んでいます。この 2 つの聴衆は、ほとんどの実際的な点で互いに戦争状態にあり、あなたは外交官です。
ClickHouse は Yandex から生まれ、途方もない量のクリックストリーム データに対する分析クエリを実行するために構築されました。可観測性を考慮して設計されたものではありません。クリックストリーム データと可観測性データには多くの共通点があるため、これはたまたま驚くほど優れています。大量、大量の追加、時間順、ほとんどがまとめて読み取られ、時々必要なデータが必要になるからです。

手を伸ばして、特定の針を 1 本見つけます。
Helm チャートを使用して自分で実行できます。 ClickHouse プラグインを介して Grafana を指定することも、独自の Web UI を使用することも、独自のフロントエンドを導入することもできます。彼らのドキュメントは実際に優れており、報告する価値があるほど珍しいものであるため、私が言及しました。ただし、ClickStack セットアップを使用したことがないので、YMMV を使用します。
特に可観測性に関しては、OpenTelemetry Collector には ClickHouse エクスポーターがあり、OTLP データを直接パイプして、初期スキーマを管理させることができます。 ClickHouse は、数十億行をスキャンし、最初に数字を見たときに嘘をついているのではないかと思わせるほどの量のデータを取り込むように設計されています。彼らは嘘をついていません。 SQL を使用してクエリを実行します。SQL はすでに存在する言語であり、2 週間前のスタートアップによって作成されたものではありません。
しかし、なぜ Clickhouse がログ専用なのでしょうか?
私はログについて暴言を吐きながら、なぜクリックハウスをもっと管理したいのかを説明しています。ここで、なぜ Clickhouse が大規模なログに優れているのかを説明しましょう。
データ形式としてのログには、いくつかの特有の特性があります。これらは追加専用です。ログ行を更新することはありませんし、ログ行を 1 つも削除することはほとんどありません。ただし、保持が開始されると一度に大量のログ行が削除されます。これらは、実際には順序どおりではありませんが、おおよそ時間順​​に到着します。これらは何日もの間誰もログを見ない状態で一気に読み取られますが、インシデントが発生すると、誰かが数秒で 10 億件のログをスキャンしようとします。ログ内のほとんどのバイトが同じサービス名、同じホスト名、同じエラー文字列、同じ JSON キーを何度も何度も繰り返すため、圧縮率が高くなります。そして重要なことに、それらをクエリするときは、ほとんどの場合、すべてのフィールドにわたる狭い時間範囲、またはいくつかのフィルターを使用した広い時間範囲にわたる集計のいずれかが必要になります。

トランザクション データベースの場合のように、「ID で特定の行を 1 つ指定してください」ということはほとんどありません。 (例外はありますが、GDPR やコンプライアンス ログのようなものは、悪夢のサブジャンルです)。
行指向データベース (Elasticsearch、Postgres、MySQL) では、単一のログ行のデータがディスク上にまとめて保存されます。ログに 40 個のフィールドがあり、クエリでそのうちの 3 個だけが考慮されている場合、運悪く、とにかくディスクから 40 個すべてを読み取ることになります。データベースはメモリ内でデータをフィルタリングしますが、ディスク I/O はすでに発生しています。
ClickHouse は各列を個別に保存します。クエリが SELECT service, status_code, count() FROM logs WHERE timestamp > now() - INTERVAL 1 HOUR GROUP BY service, status_code である場合、ClickHouse はディスクから 3 つの列 (timestamp、service、status_code) を正確に読み取ります。スキーマ内の他の 37 列は存在しないのも同様です。可観測性データでは、多くの場合、数十の属性があるにもかかわらず、特定のクエリが 3 つまたは 4 つに触れます。これは、800 GB をスキャンする場合と 40 GB をスキャンする場合の違いです。
これが、圧縮数値がばかげているように見える理由でもあります。列指向のデータは、本質的に 1 つの列内の値が互いに似ているため、行指向のデータよりもはるかに効率的に圧縮されます。 service_name 値の列には、10 億行にわたる 100 の異なる文字列が含まれる場合があります。 ZSTDはそれを朝食に食べます。実際の可観測性データでは、Elasticsearch の 2 ～ 3 倍の圧縮率と比較して、10 ～ 14 倍の圧縮率が日常的に見られます。
驚くべき点はそこじゃないけど
驚くべきことは、ClickHouse は形状を変えることなく拡張できることです。
他にどう言えばいいのかわかりません。私がこれまで扱ってきた他の可観測性バックエンドはすべて、成長するにつれて変化します。 1 日あたり 1 TB のアーキテクチャと 1 日あたり 10 TB のアーキテクチャは、明らかに異なるシステムです。

障害モード、さまざまな運用負担、さまざまなメンタル モデル。 1 日あたり 10 TB の ClickHouse は、より多くのシャードを備えた 1 日あたり 1 TB の ClickHouse と似ています。それでおしまい。それがピッチです。それが私がこれを書いているすべての理由です。
1 日あたり 1 TB であれば、最新の可観測性スタックはどれもほぼ問題ありません。この規模であれば、ほぼ何でも選択でき、生産性を高めることができます。以下の違いは実際のものですが、まだ痛みを伴うものではありません。
Logstash を備えた比較的標準的な Elasticsearch クラスターは、取り込みと Lucene インデックスの間にバッファーを提供します。ユーザーは全文検索を利用できますが、これは本当に良いことです。これは Elasticsearch が実際に最も得意とすることであり、この規模で実現します。混合データではマッピングの爆発がすでにバックグラウンド リスクとなっているため、動的マッピングを無効にするか、初日から慎重にテンプレート化する必要があります。 ILM ポリシー (ホット→ウォーム→削除) は、このサイズでもオプションではありません。これを設定し忘れると、土曜日にディスク負荷に関するページングを受けることになるからです。月額およそ 6 ～​​ 9,000 ドル。
あまりクレイジーなことは何もありません。 Alloy (旧称 Grafana Agent、RIP) は、コレクション ストーリーを単一のデーモンに統合します。これは優れています。 Loki は、有用なラベルの付け方について開発者を教育するのに時間を費やしている限り、うまく機能します。この会話は、残りのキャリアを通じて、多くの人々と何度も交わされることになります。 Mimir と Tempo は主に缶に書かれていることを実行します。月額約 3.5 ～ 5,000 ドル。
1 日あたり 1 TB の Datadog は本当に優れています。これがこの規模に合わせて構築されたものであり、それが示しています。エージェントをインストールし、ダッシュボードを確認して、家に帰ります。考えることはほとんどなく、それがすべてです。この図には、従量制課金パイプライン、インデックス付けされたログと取り込まれたログの区別、カスタム メトリクスの基数税など、コストの問題が潜んでいることがすでにわかりますが、

この規模であれば管理可能です。月額およそ 4 万 5,000 ～ 7 万 5,000 ドルですが、交渉による価格設定はさまざまなので、この数字は握りこぶし程度の話として割り引いて考えるとよいでしょう。
Datadog の全体的な価格設定の哲学は、フルタイムのエンジニアを節約することです。フレーミングは多少狂っていると思いますが、それらは非常に豊富であり、私はそうではありません。したがって、ソースを考慮してください。
これが正直な真実です。1 日あたり 1 TB の ClickHouse は、他の製品と比べて複雑ではありません。ほぼ同じです。事前に行う必要があるスキーマ設計作業を考慮すると、おそらくもう少し多くなります。 ZSTD と適切なコーデックを使用すると 10 ～ 14 倍の圧縮が得られ、Altinity Operator がキーパー調整を処理し、全体が約 7 つのポッドで実行されます。ただし、スキーマを設計する必要があります。 ORDER BY キーは非常に重要です。ネイティブの PromQL がないため、メトリクス ワークフローは Grafana プラグイン、または chproxy とアダプターを経由します。月額約 1.5 ～ 2.5,000 ドル。
この層の図を目を細めて見ると、それらはすべて同じ重量クラスにあると言えるでしょう。そして、あなたは正しいでしょう。次に何が起こるか見てみましょう。
ここで、そのうちの 1 人を除くすべての人にとって指数関数的な曲線が始まります。
Kafka はオプションではなくなりました。 1 日あたり 5 TB になると、Elasticsearch への直接書き込みによりバルク拒否ストームとバックプレッシャーが発生し、トラフィックの急増時にクラスターが完全にダウンします。つまり、現在 Kafka を実行しています。これは、Kafka をうまく実行しているか、あるいは、まったく別の 2 番目の問題が発生しようとしているかのどちらかです。シャードの計算が重要になります。ターゲット シャードが 50GB の場合、レプリカを数えながら 1 日に約 200 個のシャードを作成することになり、クラスター状態のサイズ自体が問題になります。検索可能なスナップショットと凍結層には、ほぼ確実に Elastic の商用ライセンスが必要です。ライセンス取得前は月額約 40 ～ 55,000 ドル。
あなたは今ミクロスにいます

あなたが望んでいたかどうかにかかわらず、Rvicesモード。つまり、3 つの個別のシステムに 65 以上のポッドがあり、それぞれが独自の圧縮パイプライン、独自のハッシュ リング、独自の memcached 層を備えています。ゴシップとメンバーリストのリングは実際の運用上の問題になります。 ingester のロールアウトには、-ingester.autoforget-unhealthy の慎重な調整が必要です。調整を誤ると、データが失われるか複製されます。月額約 22 ～ 32,000 ドル。
サーバーを実行しないため、運用の複雑さはまだ低くなります。しかし、現在は、Datadog の請求額を削減することに全力を注ぐ完全なパイプライン チームが必要です。除外フィルター、サンプリング ルール、カーディナリティ キャップ、タグ許可リスト、装置全体。これを私は「お金を払っているシステムの使用を避けるためにシステムを構築する」という罠と呼んでいますが、一度その罠にはまってしまうと、永遠にその罠にはまり続けることになります。パイプライン チームがどれだけ積極的になるかによって異なりますが、月額およそ 18 万～35 万ドルです。
ここは、基本的に SaaS プロバイダーと常に格闘し、コストを削減する方法を見つけるために請求書類を精査する場所でもあります。それは敵対的な関係であり、私は楽しくありません。
図を見ると、基本的にシャードを追加しただけであることがわかります。それでおしまい。それが変化です。同じ演算子、同じクエリ エンジン、同じクエリ言語、同じ m

[切り捨てられた]

## Original Extract

For roughly the last ten years, a meaningful percentage of my working hours have been spent thinking about observability. If you're not familiar with the term, "observability" is what we call it now that "monitoring" doesn't sound expensive enough. The actual work is unglamorous in that you collect

Clickhouse is winning the Observability Wars
For roughly the last ten years, a meaningful percentage of my working hours have been spent thinking about observability. If you're not familiar with the term, "observability" is what we call it now that "monitoring" doesn't sound expensive enough. The actual work is unglamorous in that you collect a lot of logs, some metrics, a few traces, and then you give them to people.
I generally like my job. I like that we're always trying new ideas and approaches. I like the fact that when things go wrong, the answer is almost always sitting there in the data, waiting to be found by whoever is patient enough to look. But I want to be honest with you: in ten years of doing this work, across a half-dozen companies and every observability platform you've heard of and a few you probably haven't, logs have never stopped being the worst part of the job. They were the worst part when I started. They are the worst part today. I fully expect them to be the worst part of this job forever until the robots rise up and rip my head off in one clean sweep.
I've written about why logs are terrible before , so I'll spare you the full lecture and give you the short version.
Every developer's expectations for logs are set by a single formative experience: the syslog box. Or a container running locally. Or tail -f on a production server they probably shouldn't have SSH'd into. The point is that at some early, tender moment in their career, they had an experience with logs that was flawless. They ran grep and something useful came back. They piped it into jq and got exactly what they needed.
This experience is the observability equivalent of a first kiss. It ruins them for everything that comes after.
Because here is the thing about that flawless experience: it works because the system is small, the volume is trivial, and the person querying is the same person who wrote the log line. There is no schema drift, no cardinality explosion, no cross-team consumer with dashboard expectations, no VP asking why the "revenue events" graph has a gap in it.
Then there are forty services. Now there are four hundred. Now the logs are being consumed not just by developers but by customer service, who need to look up a specific user's failed checkout from Tuesday. And by the data team, who are quietly building a business-critical dashboard on top of a log line that a backend engineer is about to refactor without telling anyone. And by the on-call, who at 3 AM does not want to learn a new query language, does not want to think about index patterns, and would like the search bar to just work.
So you have a technical problem — the volume is enormous, the shape is inconsistent, the queries are unpredictable — sitting on top of an expectations problem, which is worse. Developers want logs instantly, they want to run arbitrary operations on them, and they will not commit to a schema. Meanwhile the less-technical consumers of that same data want the dashboards to be stable forever, the UI to be forgiving, and the whole thing to feel like a normal product. These two audiences are, in most practical respects, at war with each other, and you are the diplomat.
ClickHouse came out of Yandex, where it was built to chew through analytical queries against absurd volumes of clickstream data. It was not designed for observability. It just happens to be shockingly good at it, because clickstream data and observability data have a lot in common: high volume, append-heavy, time-ordered, mostly read in aggregate, and every so often you need to reach in and find one specific needle.
You can run it yourself with Helm charts. You can point Grafana at it via the ClickHouse plugin, or use their own web UI, or bring your own frontend. Their docs are actually good, which I mention because it's rare enough to be worth flagging. I've never used their ClickStack setup though, so YMMV.
For observability specifically, the OpenTelemetry Collector has a ClickHouse exporter, which means you can pipe OTLP data straight in and let it manage the initial schema for you. ClickHouse is designed to scan billions of rows and ingest an amount of data that, when you first see the numbers, makes you assume they're lying. They're not lying. You query it with SQL, which is a language that already exists and was not created by a startup two weeks ago.
But why Clickhouse specifically for logs?
I'm ranting about logs and then I'm explaining why I like to administer Clickhouse more. Let me take a second and explain why Clickhouse is really good at logs at scale.
Logs, as a data shape, have some peculiar properties. They're append-only. You never update a log line, and you almost never delete a single one, though you delete a lot of them at once when retention kicks in. They arrive roughly in time order, though never actually in order. They're read in bursts where nobody looks at logs for days, and then during an incident somebody wants to scan a billion of them in seconds. They're highly compressible, because most of the bytes in your logs are repeated: the same service names, the same hostnames, the same error strings, the same JSON keys, over and over and over again. And critically, when you query them, you almost always want either a narrow time range across all fields or an aggregation across a wide time range with a few filters. You very rarely want "give me one specific row by ID" the way you would from a transactional database. (There are exceptions when its something like GDPR or compliance logging which is its own subgenre of nightmares).
In a row-oriented database — Elasticsearch, Postgres, MySQL — the data for a single log line is stored together on disk. If your log has 40 fields and your query only cares about 3 of them, tough luck, you're reading all 40 from disk anyway. The database will filter it in memory, but the disk I/O has already happened.
ClickHouse stores each column separately. If your query says SELECT service, status_code, count() FROM logs WHERE timestamp > now() - INTERVAL 1 HOUR GROUP BY service, status_code, ClickHouse reads exactly three columns off disk: timestamp, service, and status_code. The other 37 columns in your schema might as well not exist. On observability data, where you often have dozens of attributes but any given query touches three or four, this is the difference between scanning 800GB and scanning 40GB.
This is also why the compression numbers look absurd. Columnar data compresses far better than row-oriented data because the values within a single column are, by nature, similar to each other. A column of service_name values might have a hundred distinct strings across a billion rows. ZSTD eats that for breakfast. You'll routinely see 10–14x compression ratios on real observability data, compared to 2–3x for Elasticsearch.
That's not the amazing part though
The amazing thing is that ClickHouse scales without changing shape.
I don't know how else to say this. Every other observability backend I've worked with mutates as it grows. The architecture at 1 TB a day and the architecture at 10 TB a day are recognizably different systems, with different failure modes, different ops burdens, and different mental models. ClickHouse at 10 TB a day looks like ClickHouse at 1 TB a day with more shards. That's it. That's the pitch. That's the whole reason I'm writing this.
At 1 TB a day, every modern observability stack is roughly okay. If you're at this scale, you can pick almost anything and be productive. The differences below are real but they're not yet painful.
A relatively vanilla Elasticsearch cluster with Logstash providing some buffer between ingest and the Lucene indexes. Users get full-text search, which is genuinely good — this is the thing Elasticsearch is actually best at, and at this scale it delivers. Mapping explosions are already a background risk with mixed data, so dynamic mapping needs to be disabled or carefully templated from day one. ILM policies (hot → warm → delete) are non-optional even at this size, because forgetting to set them is how you get paged on a Saturday about disk pressure. Roughly $6–9K/month.
Nothing too crazy. Alloy (formerly Grafana Agent, RIP) unifies the collection story into a single daemon, which is nice. Loki works well as long as you spend some time educating developers on how to attach useful labels — a conversation you will have many times, with many people, for the rest of your career. Mimir and Tempo largely do what it says on the tin. Roughly $3.5–5K/month.
At 1 TB a day, Datadog is genuinely great. This is the scale it was built for, and it shows. You install the agent, you look at dashboards, you go home. There is almost nothing to think about, which is the entire point. You can already see the shape of the cost problem lurking in the diagram — the metered pipelines, the indexed-vs-ingested logs distinction, the custom metrics cardinality tax — but at this scale it's manageable. Roughly $45–75K/month, though negotiated pricing varies enough that I'd take that number with a grain of salt the size of a fist.
Datadog's whole pricing philosophy is that they save you a full-time engineer. I think that framing is somewhat deranged, but they are extremely rich and I am not, so consider your source.
Here is the honest truth: at 1 TB a day, ClickHouse is not less complicated than its peers. It's roughly the same. Maybe slightly more, if you count the schema design work you have to do up front. You get 10–14x compression with ZSTD and proper codecs, the Altinity Operator handles keeper coordination and the whole thing runs in about seven pods. But you do have to design your schemas. ORDER BY keys matter enormously. There is no native PromQL, so metrics workflows go through the Grafana plugin or through chproxy and an adapter. Roughly $1.5–2.5K/month.
If you took the diagrams at this tier and squinted, you'd say they're all in the same weight class. And you'd be right. Now watch what happens next.
This is where the exponential curve kicks in for everybody except one of these.
Kafka is no longer optional. At 5 TB a day, direct writes into Elasticsearch cause bulk-reject storms and backpressure that will absolutely take your cluster down during a traffic spike. So now you're running Kafka, which means you're either running Kafka well or you're about to have a second, entirely different set of problems. Shard math becomes critical — at 50GB target shards, you're minting ~200 shards a day counting replicas, and your cluster state size becomes its own concern. You almost certainly need Elastic's commercial license for searchable snapshots and the frozen tier. Roughly $40–55K/month before licensing.
You are now in microservices mode, whether you wanted to be or not. That means 65+ pods across three separate systems, each with its own compaction pipeline, its own hash ring, its own memcached tier. The gossip/memberlist ring becomes a real operational concern; ingester rollouts require careful -ingester.autoforget-unhealthy tuning, and if you get it wrong you either lose data or duplicate it. Roughly $22–32K/month.
The operational complexity is still low, in that you don't run any servers. But you now need a full pipeline team whose entire job is reducing your Datadog bill. Exclusion filters, sampling rules, cardinality caps, tag allow-lists, the whole apparatus. This is what I call the "you build a system to avoid using the system you're paying for" trap, and once you're in it, you are in it forever. Roughly $180–350K/month, depending on how aggressive the pipeline team gets.
This is also where you are basically fighting with your SaaS provider all the time, pouring over their billing documentation to figure out how to reduce costs. It's a hostile relationship and one I don't enjoy.
You'll notice, if you look at the diagram, that I basically just added shards. That's it. That's the change. Same operator, same query engine, same query language, same m

[truncated]
