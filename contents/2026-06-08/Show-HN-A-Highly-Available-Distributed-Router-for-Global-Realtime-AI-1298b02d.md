---
source: "https://cerebrium.ai/blog/thalamus-our-highly-available-distributed-router-for-global-realtime-ai-workloads"
hn_url: "https://news.ycombinator.com/item?id=48446710"
title: "Show HN: A Highly Available Distributed Router for Global Realtime AI"
article_title: "Our Highly Available Globally Distributed AI Router"
author: "wesrobin"
captured_at: "2026-06-08T16:28:33Z"
capture_tool: "hn-digest"
hn_id: 48446710
score: 6
comments: 3
posted_at: "2026-06-08T15:33:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A Highly Available Distributed Router for Global Realtime AI

- HN: [48446710](https://news.ycombinator.com/item?id=48446710)
- Source: [cerebrium.ai](https://cerebrium.ai/blog/thalamus-our-highly-available-distributed-router-for-global-realtime-ai-workloads)
- Score: 6
- Comments: 3
- Posted: 2026-06-08T15:33:24Z

## Translation

タイトル: HN を表示: グローバル リアルタイム AI 向けの高可用性分散ルーター
記事のタイトル: 可用性の高い世界中に分散された AI ルーター
説明: お気づきでないかもしれませんが、現在 GPU は不足しています。世界的な推論ブームにより、供給が追いつかないほどの速さで需要が増大しており、セレブリウムでは...

記事本文:
可用性の高い世界中に分散された AI ルーター
セレブリウムの使用例
">
大規模な言語モデル
使用例
">
大規模な言語モデル
使用例
">
大規模な言語モデル
使用例
">
大規模な言語モデル
Thalamus - グローバルなリアルタイム AI ワークロード向けの高可用性分散ルーター
ウェスリー・ロビンソン
技術スタッフ
Thalamus - グローバルなリアルタイム AI ワークロード向けの高可用性分散ルーター
お気づきでないかもしれませんが、現在 GPU は不足しています。推論の世界的なブームにより、供給が追いつかないほどの速さで需要が増大しています。セレブリウムでは、これは、1 つのリージョンにデプロイするだけで残りをロード バランサーに任せることはできないことを意味します。利用可能な容量を維持するために、お客様のアプリは最終的に、複数のクラスター、複数のデータセンター、複数の大陸、複数の GPU プロバイダーにまたがって流出します。これは容量不足に対する素晴らしい答えですが、私たちにとっては特有の問題が生じます。リクエストが届いたとき、どのクラスターがそれに対応すべきでしょうか?
その決定には数ミリ秒の予算があり、判断を誤ると、私たちとお客様にとって多大な損害が発生します。リアルタイムのテキスト読み上げリクエストを間違った大陸のクラスターにルーティングすると、レイテンシー バジェットを超過し、モデルが開始される前に顧客を失ってしまいます。ヘルスチェックを行わずにトレーニング ジョブを利用可能な最も安価なクラスターにルーティングすると、トレーニングの実行が途中で失敗し、高価な GPU 時間が何時間も無駄になる可能性があります。
Thalamus はその意思決定を行うサービスです。リクエストが健全で低遅延の Cerebrium クラスターに確実に到達できるようにするのは、ルーティング ブレインです。この投稿では、Thalamus を使用することで、Cerebrium の顧客が慎重な導入構成ではなく配信に集中できるようにする方法について説明します。
Cerebrium にアプリをデプロイすると、単一の URL を取得します

。どのクラスターがコードを実行するかを気にする必要はありません。ただ、一部のクラスターが実行し、それが迅速かつ確実に実行されることだけを考慮する必要があります。アプリは利用できる場合にのみ効果を発揮します。
その URL が単一クラスター内で実行されているアプリにマップされている場合、問題は解決されます。ローカル ロード バランサーは、独自のコンテナーを直接認識し、コンテナーの状態を正確に認識し、それに応じてルーティングします。その URL が多数のデータ センターで実行されているアプリに解決されると、問題は複雑になります。地理的に離れた多数のクラスターがリクエストを処理できる場合、リクエストはどこに送信されるべきでしょうか。
顧客がプロンプトを LLM にパイプして応答を返す hello-world アプリをデプロイしたとします。これは、us-east、eu-west、ap-southeast のクラスターで利用できますが、us-west クラスターでは明示的に利用できません。カリフォルニアのユーザーがリクエストを送信します。
最初のホップは DNS です。 Geo-DNS を使用すると、驚くほど遠くまで到達できます。ネームサーバーは近くの Cerebrium イングレスの IP を返しますが、実際のルーティングの決定には役立ちません。 DNS は、この顧客のアプリがどのクラスターにデプロイされているか、現在どのクラスターに GPU 容量に余裕があるか、どのクラスターが正常であるかを知りません。したがって、DNS はリクエストをクローズしてから、地域の Thamus インスタンスの 1 つに渡します。
ここから興味深い疑問が始まります。Thalamus は、どのクラスターがこのリクエストを処理できるか、どのクラスターが処理すべきかをどのようにして知るのでしょうか?
エッジ コンピューティングとロード バランサーだけでは不十分な理由
これには、エッジ コンピューティング ソリューションでは十分ではありません。 AI ワークロードは長期間存続する可能性があります。トレーニング ジョブやライブ ストリームでは、単一の接続が何分間も、場合によっては何時間も開いたままになることがあります。ほとんどのエッジ コンピューティング プラットフォームでは、リクエストの継続時間は 1 分未満に制限されています。したがって、ルーティングの決定は次のように感じられますが、

エッジの問題ではありますが、その後のプロキシ処理は明らかに問題ではありません。
ラウンドロビンや最小接続のみのような「従来の」負荷分散ソリューションは、ここでの目的には適していません。すべてはレイテンシーに帰着します。クラスター内ロード バランサーは通常、単一のクラスターまたはリージョン内で、近くにあるほとんど交換可能なワーカー間でバランスをとります。利用可能になる新しいリソースは、バランシングのためにリクエストが来たときに一貫した状態を保証する、おしゃべりな Raft 風のコンセンサス アルゴリズムを通じて、ほぼ即座にブロードキャストできます。この種のプロトコルは、更新が速い場合にのみ機能します。一方、Thalamus は異種リソース全体にわたってグローバルな決定を下す必要があります。さまざまなクラスター、リージョン、プロバイダーは海を隔てていることが多いため、一貫した状態に依存することができません。
分散状態 - 押す、引かない
Thalamus は、リクエストの送信先を決定するために、静的情報と動的情報の 2 種類の情報を必要とします。
静的状態は、アプリの実行が許可されるクラスターやコンプライアンスのためのリージョン制限など、予測可能なイベントでのみ変更されるものです。デプロイと構成の更新のたびに変化しますが、事前に計算できるほど頻繁には変化しません。
動的状態はそれ以外のすべてです。ユーザーのアプリケーションの現在実行中のレプリカの数、および現在リクエストを処理しているレプリカの数。現在の GPU スポット価格はいくらですか。各クラスターが正常かどうか。このようなものは常に予測不可能に変化するため、事前に計算することはできません。
Geo-DNS がクライアントを近くの Thamus インスタンスに誘導したら、Thalamus はどのアップストリーム クラスターが実際にワークロードを処理できるかを決定する必要があります。アプリが展開される場所、利用可能な容量、必要なハードウェア、および適用されるコンプライアンス ルールに応じて、別のリージョンまたはプロバイダーに存在する可能性があります。アプリケーションがデプロイされている場合は、cl

その場合、多くの場合、Thalmaus はローカル インスタンスにルーティングすることになります。これが、Geo-DNS がここで非常に役立つ理由です。
単純な解決策は、Thalamus にリクエスト パス上の各クラスターに「容量はありますか?」とクエリを実行させることです。あなたは健康ですか？現在のレイテンシはどれくらいですか? 8 つの候補クラスターが大陸に分散している場合、クラスターごとに 1 往復したとしても、最初の決定が下されるまでにレイテンシ バジェットが数桁も超過してしまいます。したがって、ホット パスでは状態をフェッチすることはできません。
代わりに、すべてのクラスター内で小さなサービス (内部的にはクラスター アグリゲーターと呼んでいます) を実行します。その全体の仕事は、ローカルで何が起こっているかを監視し、その情報を中央のデータ ストアにプッシュし、Thalamus がそこから読み取ることです。スケールアップとスケールダウンのイベントを監視し、アプリごとおよび GPU タイプごとに実行容量と予約容量を追跡し、現在のノードのコストを監視し、それらのスナップショットを継続的に上向きに書き込みます。
新しいクラスターがオンラインになると、独自のクラスター アグリゲーターが導入されるという素晴らしい結果が得られます。中央のレジストリを更新する必要はなく、Thalamus は新しいクラスターがデータ ストアに表示されるのを確認し、自動的にルーティングの対象として検討し始めます。
夢がキャッシュされるGPUジャングル
また、私たちが行ったアーキテクチャ上の決定を正当化する価値もあります。つまり、すべての Thalamus インスタンスとすべてのクラスター アグリゲーターによって共有される 1 つの中央データ ストアが存在します。最初は、この決定は低遅延ルーターの形として間違っているように思えました。 1 桁ミリ秒以内に実行する必要があるルーティング決定の途中にデータベースを配置する必要があるのはなぜでしょうか。
答えは、データベースがリクエスト パスの途中にないからです。これは状態伝播パスの中間です。
すべてのクラスターはローカル スナップショットを継続的に上方にプッシュします: 現在のレプリカ、予約済み容量、利用可能な GPU タイプ、修復

次に、ノードのコストと最近のスケーリング イベントです。その後、その情報はすべての Thamus インスタンスに複製され、そこで読み取りがローカルで行われます。ホット パス上では、Thalamus はリモート データベースやリモート クラスターに何をすべきかを尋ねません。これは、最近の過去のローカル コピーから読み取ります。
完全に新しいグローバル状態はルーティング バジェット内では達成できないため、この区別は重要です。クラスターとそのアグリゲーター、アグリゲーターとデータ ストア、データ ストアと各 Thamus インスタンスの間には、常に遅延が発生します。光の移動速度は非常に速く、リクエストが到着する頃には、「現在の」グローバル状態はすでにわずかに古くなっています。
Thalamus は、継続的に同期されるローカルのスナップショットから決定を行います。これにより、予測可能なルーティング レイテンシーが得られ、ホット パスにリモート呼び出しが発生しないように維持され、すべての Thalamus インスタンスに、検査および再生できる最近の過去のほぼ一貫したビューが与えられます。これにいくつかのコツを加えれば、ルーティングを決定するのに十分です。
これを実現するために、プライマリから同期する読み取り専用レプリカが埋め込まれた分散 SQLite スタイルのデータベースである Turso を使用します。各 Thalamus インスタンスは、コントロール プレーン データベースの継続的に同期されたローカル コピーを保持します。ルックアップの p99 は約 500 マイクロ秒なので、ルーティングの決定により多くの予算を費やすことができます。
Turso の同期機能はセットアップが非常に簡単で、ルックアップが非常に高速であるため、同期されたデータベースは基本的に最高の種類のキャッシュとして機能します。つまり、構築する必要がありません。
リクエストが到着すると、Thalamus はそれをターゲット アプリに解決し、現在そのアプリにサービスを提供できるクラスターを Turso からロードし、クラスターの候補セットの構築を開始します。各候補についてチェックします (すべて処理中の状態から):
候補者のアプリに対する現在のホット、ウォーム、コールドの能力はどれくらいですか?
私たちのものは何ですか

その候補までの最新 (移動平均) 測定ネットワーク遅延はどれですか?
その候補のターゲット ハードウェアのコストはいくらですか?
候補者は健康診断に合格していますか?
これらの容量階層の簡単な定義は次のとおりです。
Hot: コンテナーにはこのアプリがすでにロードされており、すぐに使用できます
Warm: 新しいコンテナを数秒でスケールできる適切な SKU のノードを実行します。
コールド: プロバイダー側のヘッドルームに突入できますが、時間がかかります
候補クラスターを、それらが満足できる最良の層に基づいてランク付けします。したがって、クラスターにホット容量がある場合は、ウォーム容量またはコールド容量のみを持つクラスターよりもすぐに優先されます。すでに実行されているコンテナーからサービスを提供するよりも、どこかでコールドスタートする方が良いシナリオはほとんどありません。
ここで、ランク付けされた一連の候補クラスターを使用して、最も適切なものを選択する必要があります。ここで、古いデータの問題を回避する必要があります。私たちが答えようとしている疑問は、どのクラスターがこのリクエストを最短時間で処理する可能性が最も高いかということです。現在の正確な容量の広がりを知ることはできないため、「最良」を選択することは決してありません。利用可能な容量によって重み付けされ、修飾子のパイプラインによって調整され、確率的に選択されます。より多くのホットキャパシティ、より低いレイテンシー、より優れたキャッシュの局所性を備えたクラスターは、現時点で特定のリクエストを処理できる可能性がはるかに高くなりますが、すべてのリクエストを受信する必要はありません。
必要に応じて、ここに追加の条件を追加することもできます。考慮すべき非常に重要なデータの 1 つは、そのクラスター内のアプリが正常かどうかです。あるいは、クラスター、さらにはデータセンター自体が正常かどうか。そうでない場合は、クラスターを候補リストから除外します。これは、ヘルスチェックの一定のストリームを通じて動的に行われ、現在は自動フェイルオーバーが機能しています。
これは私たちにスムーを与えます

ルーティング システム。 1 つのリージョンがトラフィックを独占することなく、近くのクラスターを優先できます。容量が十分にある場合は、より安価な容量を選択できますが、容量が不足するとコストを気にしなくなります。最近実行されたクラスターにワークロードを偏らせることができ、キャッシュ ヒットが改善され、起動時間が短縮されます。また、データ主権のために「このアプリはヨーロッパでのみ実行しなければならない」など、ユーザーまたはビジネスに厳しい制約を課すこともできます。
ルーティング システムは、後で何を行ったかを説明できる場合にのみ調整可能です。 Thalamus が行うすべての決定は、その背後にある完全な入力コンテキスト (ターゲット アプリ、候補クラスター、容量層、正常性状態、レイテンシ推定、コスト入力、適用された修飾子、最終的な選択、ルーティングの理由) とともに記録されます。これらの理由はメトリクス、ログ、トレース、ダッシュボードなどあらゆる場所に現れるため、リクエストが予期せぬ場所に送信された場合、いつでもその理由を答えることができます。
ルーティングは静的な問題ではないため、これは特に役立ちます。 「適切な」意思決定ツリーは、フリートの変更、顧客のワークロードの変更、GPU の価格変更、新しいリージョンのオンライン化に応じて変化します。
このデータはすべてキャプチャされるため、実際の運用ルーティング イベントを取得し、アルゴリズムの変更されたバージョンに対してシミュレートできます。

[切り捨てられた]

## Original Extract

If you haven’t noticed, GPUs are scarce right now. The global boom in inference has demand growing faster than supply can keep up, and at Cerebrium ...

Our Highly Available Globally Distributed AI Router
Cerebrium Use Cases
">
Large Language Models
Use Cases
">
Large Language Models
Use Cases
">
Large Language Models
Use Cases
">
Large Language Models
Thalamus - Our Highly Available Distributed Router for Global Realtime AI Workloads
Wesley Robinson
Techical Member of Staff
Thalamus - Our Highly Available Distributed Router for Global Realtime AI Workloads
If you haven’t noticed, GPUs are scarce right now. The global boom in inference has demand growing faster than supply can keep up, and at Cerebrium that means we can't just deploy in one region and let a load balancer sort the rest out. To keep capacity available, our customers' apps end up spilling out across multiple clusters, in multiple data centers, on multiple continents, and across multiple GPU providers. That's a great answer to the capacity crunch, but it creates a unique problem for us. When a request comes in, which of those clusters should serve it?
That decision has a budget of a few milliseconds, and getting it wrong is expensive for us and our customers. Route a real-time text-to-speech request to a cluster on the wrong continent and we’ve blown the latency budget and lost a customer before their model has even started. Route a training job to the cheapest available cluster without checking health and a training run could fail mid-way throwing away many hours of expensive GPU time.
Thalamus is the service that makes that decision. It's the routing brain that ensures requests find their way to healthy, cheaper, lower latency Cerebrium clusters. This post is about how Thalamus allows Cerebrium’s customers to focus more on delivery, and less on careful deployment configurations.
When you deploy an app on Cerebrium, you get a single URL. You shouldn’t have to care about which cluster runs your code - only that some cluster does, and that it does so quickly and reliably. We have to care about it a great deal, an app is only effective if it is available.
If that URL maps to an app running in a single cluster, the problem is contained. A local load balancer has visibility into its own containers directly, knows their state precisely, and routes accordingly. When that URL resolves to an app running across many data centers, the problems compound: given many geographically separated clusters capable of handling the request, where should it go?
Let's say a customer has deployed a hello-world app that pipes a prompt to an LLM and returns the response. It's available in clusters in us-east, eu-west, and ap-southeast, but explicitly not in our us-west cluster. A user in California sends a request.
The first hop is DNS. Geo-DNS can get us surprisingly far - the nameserver hands back the IP of a nearby Cerebrium ingress - but it can't help with the actual routing decision. DNS doesn't know which clusters this customer's app is deployed to, doesn't know which of those have spare GPU capacity right now, and doesn't know which cluster is healthy. So DNS gets the request close, then hands it off to one of our regional Thalamus instances.
That's where the interesting questions start: how does Thalamus know which clusters can handle this request, and which one should?
Why edge compute and load balancers are not enough
Edge compute solutions are not sufficient for this. AI workloads could be long-lived - a training job or a live stream can hold a single connection open for many minutes, sometimes hours. Most edge-compute platforms cap request duration at well under a minute. So although a routing decision feels like an edge problem, the proxying that follows it is decidedly not.
‘Traditional’ load balancing solutions like round robin or least connections alone are not fit for purpose here. It all boils down to latency. In-cluster load balancers typically balance across nearby, mostly interchangeable workers inside a single cluster or region. New resources becoming available can be broadcasted almost instantly through chatty Raft-esque consensus algorithms that guarantee consistent state when a request comes in for balancing. This kind of protocol only works when updates are fast. Thalamus meanwhile must make a global decision across heterogeneous resources: different clusters, regions and providers that are often oceans apart, so cannot rely on consistent state.
Distributing state - push, don't pull
Thalamus needs two kinds of information to decide where a request goes: static and dynamic.
Static state is the stuff that only changes at predictable events like which clusters an app is allowed to run on, or a region restriction for compliance. It changes on each deploy and config update - infrequently enough that we can pre-compute it.
Dynamic state is everything else. The number of currently running replicas of a user’s application, and how many of those are currently processing requests. What the current GPU spot price is. Whether each cluster is healthy. This stuff changes constantly and unpredictably, and we can't precompute it.
Once Geo-DNS directs the client to a nearby Thalamus instance, Thalamus needs to decide which upstream cluster can actually serve the workload. It may still be in another region or provider, depending on where the app is deployed, what capacity is available, what hardware it requires, and what compliance rules apply. If an application is deployed close to its users, then more often than not Thalmaus will end up routing to a local instance, which is why Geo-DNS is really helpful here.
The naive solution is to have Thalamus query each cluster on the request path: Do you have capacity? Are you healthy? What's your latency right now? Even at one round trip per cluster, with eight candidate clusters spread across continents, we'd blow the latency budget by multiple orders of magnitude before the first decision was made. Thus, state can never be fetched on the hot path.
Instead, inside every cluster we run a small service (internally we call it cluster-aggregator) whose entire job is to watch what's happening locally and push that information to a central data store that Thalamus reads from. It watches scale-up and scale-down events, tracks running and reserved capacity per app and per GPU type, monitors current node costs, and continuously writes those snapshots upward.
A nice consequence is that when a new cluster comes online, it brings its own cluster-aggregator: No central registry needs updating and Thalamus sees a new cluster appear in the data store and starts considering it for routing automatically.
GPU Jungle Where Dreams Are Cached
It is also worthwhile justifying an architectural decision we made: there is one central data store shared by every Thalamus instance and every cluster-aggregator. At first this decision felt like the wrong shape for a low-latency router. Why put a database in the middle of a routing decision that needs to happen in single-digit milliseconds?
The answer is that the database is not in the middle of the request path. It is the middle of the state-propagation path.
Every cluster continuously pushes local snapshots upward: current replicas, reserved capacity, available GPU types, health, node costs, and recent scaling events. That information is then replicated out to every Thalamus instance, where reads happen locally. On the hot path, Thalamus is not asking a remote database or a remote cluster what to do. It is reading from a local copy of the recent past.
That distinction matters because perfectly fresh global state is not achievable within the routing budget. There is always latency between the cluster and its aggregator, the aggregator and the data store, the data store and each Thalamus instance. Light only travels so fast - and by the time a request arrives - the “current” global state is already slightly old.
Thalamus makes decisions from a local, continuously-synced snapshot. That gives us predictable routing latency, keeps the hot path free of remote calls, and gives every Thalamus instance a roughly consistent view of the recent past that we can inspect and replay. This, plus a few tricks, is sufficient to make a routing decision.
To achieve this we use Turso , a distributed SQLite-style database with embedded read-only replicas that sync from the primary. Each Thalamus instance keeps a continuously synced local copy of the control-plane database. Lookups have a p99 of roughly 500 microseconds, which allows us to spend more of that budget making routing decisions.
Turso’s sync feature is so easy to set up and lookups are so fast, the synced db essentially acts as the best kind of cache: one we didn’t have to build.
When a request arrives, Thalamus resolves it to a target app, loads the clusters that can currently serve that app from Turso, and starts building a candidate set of clusters. For each candidate it checks (all from in-process state):
What's the candidate’s current hot, warm, and cold capacity for the app?
What's our latest (moving average) measured network latency to that candidate?
What is the cost of the target hardware in that candidate?
Is the candidate passing health checks?
As a short definition of those capacity tiers:
Hot: containers already loaded with this app, ready right now
Warm: running nodes of the appropriate SKU that can scale a new container in seconds
Cold: provider-side headroom we can burst into, but takes longer
We rank candidate clusters by the best tier they can satisfy. So, if a cluster has any hot capacity, it is instantly preferable over a cluster that only has warm or cold capacity. There's almost no scenario where it's better to cold-start somewhere than to serve from a container that's already running.
Now, with a set of ranked candidate clusters, we need to pick the most suitable one. This is where we must work around the stale data problem. The question we’re trying to answer really becomes which cluster is the most likely to serve this request in the shortest amount of time? Because we cannot know what the exact current capacity spread is, we never pick "the best," we pick probabilistically, weighted by available capacity and then nudged by a pipeline of modifiers. A cluster with more hot capacity, lower latency, and better cache locality is far more likely to be capable of processing any given request right now, but it should not receive every request.
We can also add extra conditions here as needed. One extremely important piece of data to consider is whether or not the apps in that cluster are healthy. Or if the cluster, or even the data center itself is healthy. If not, we disqualify the cluster from the list of candidates. This is done dynamically via a constant stream of health checks, and now we have automatic failover.
This gives us a smoother routing system. We can prefer nearby clusters without letting one region monopolise traffic. We can prefer cheaper capacity when there is plenty of it, but stop caring about cost when capacity gets tight. We can bias workloads toward clusters where they have recently run, improving cache hits and reducing startup time. And we can still enforce hard user or business constraints, like “this app must only run in Europe” for data sovereignty.
A routing system is only tunable if we can explain what it did afterward. Every decision Thalamus makes is logged with the full input context behind it: the target app, candidate clusters, capacity tier, health state, latency estimate, cost inputs, applied modifiers, final choice, and routing reason. Those reasons show up everywhere - metrics, logs, traces, and dashboards - so when a request goes somewhere unexpected, we can always answer why.
This has been especially useful because routing is not a static problem. The “right” decision tree changes as our fleet changes, customer workloads change, GPU pricing changes, and new regions come online.
Since all this data is captured, we can take real production routing events and simulate them against an altered version of the algorithm.

[truncated]
