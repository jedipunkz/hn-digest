---
source: "https://www.pingcap.com/blog/how-we-built-mem9-agent-memory-product/"
hn_url: "https://news.ycombinator.com/item?id=48332053"
title: "Lessons from Shipping Persistent Memory for AI Agents"
article_title: "Agent Memory Product: How We Built mem9 on TiDB Cloud"
author: "jinqueeny"
captured_at: "2026-05-30T11:35:53Z"
capture_tool: "hn-digest"
hn_id: 48332053
score: 1
comments: 1
posted_at: "2026-05-30T02:59:43Z"
tags:
  - hacker-news
  - translated
---

# Lessons from Shipping Persistent Memory for AI Agents

- HN: [48332053](https://news.ycombinator.com/item?id=48332053)
- Source: [www.pingcap.com](https://www.pingcap.com/blog/how-we-built-mem9-agent-memory-product/)
- Score: 1
- Comments: 1
- Posted: 2026-05-30T02:59:43Z

## Translation

タイトル: AI エージェント向けの永続メモリの出荷から得た教訓
記事のタイトル: エージェント メモリ製品: TiDB クラウド上に mem9 を構築する方法
説明: How a customer ask 2026 年 3 月に、コーディング エージェント向けのエージェント メモリ製品である mem9 になりました。プロトタイピング、検索、評価に関するレッスン。

記事本文:
エージェント メモリ製品: TiDB クラウド上に mem9 を構築した方法
何百万ものエージェント支店。 1 つのデータベース。 2026 年 6 月 4 日に開催される TiDB SCaiLE Europe にご参加ください。今すぐご登録ください。
AI
エージェント AI の TIDB
Agentic AI 用データベース
エージェントのメモリ、状態、マルチホップ推論専用に構築
ベクトル検索と RAG
ネイティブ ベクター インデックス作成および検索拡張生成パイプライン
クイック スタート: エージェントティック メモリ
永続的なエージェント メモリを数秒でスピンアップ - 設定不要
AI アプリケーションの構築
AI アプリを迅速に出荷するための SDK、ガイド、テンプレート
製品
トランザクション、AI、その他の最新アプリケーションを強化するためにイノベーターに信頼されているオープンソースの分散 SQL データベース。
世界中のイノベーションリーダーによって信頼され、検証されています。
TiDB がデータの機密性と可用性をどのように確保しているかをご覧ください。
トランザクション、AI、その他の最新アプリケーションを強化するためにイノベーターに信頼されているオープンソースの分散 SQL データベース。
世界中のイノベーションリーダーによって信頼され、検証されています。
TiDB がデータの機密性と可用性をどのように確保しているかをご覧ください。
mem9 の構築方法: AI エージェント用の永続メモリの出荷から得た教訓
mem9 はロードマップではなく、2026 年 3 月に顧客の要望として開始されました。計画を立てる前にプロトタイプを出荷しました。
エージェントのメモリはストレージの問題ではありません。これは、摂取、ランキング、評価、製品の判断が交差するエンジニアリングの問題です。
メモリ API だけでは製品ではありません。人々は、エージェントが覚えていることを見て、検査し、信頼し、修正したいと考えています。
mem9 は、TiDB Cloud Zero の背後にある同じ基板である TiDB Cloud 上で実行されます。
2026 年 3 月初旬、お客様から、簡単そうに見えて、エージェント スタックで最も難しい問題の 1 つであることが判明したことについて私たちに問い合わせがありました。それは、「エージェントに記憶させる」というものでした。
私たちは洗練されたロードマップ、つまり重量級のアーキテクチャから始めたわけではありません。

ビュー、または 6 か月の製品プラン。私たちは、多くの製品がそうしているように、具体的なユーザーの痛み、大まかなプロトタイプ、そして「これは興味深い」と「これを出荷する必要がある」の間の非常に短い距離からスタートしました。
それが mem9 の始まりでした。
今振り返ると、mem9 は従来のソフトウェア プロジェクトというよりも、圧縮されたスタートアップの年のように感じられます。顧客主導の高速プロトタイプとして始まったものは、すぐに製品になり、次にプラットフォームになり、本番環境で実際にエージェントのメモリが必要なものをさらに深く調査するようになりました。目に見える特徴はすぐに変わりましたが、核となる疑問は変わりませんでした。エージェントが他のことで圧倒されずに、重要なことを思い出せるようにするにはどうすればよいでしょうか?
これが私たちがこれまでに学んだことです。
それは市場理論ではなく現実の問題から始まった
mem9 の本当の始まりは、カテゴリー マップや戦略デッキではありませんでした。それは実際的な質問をする顧客でした。エージェントがセッション間で永続的なメモリを保持できた場合、ユーザーは実際に違いを感じるでしょうか?
私たちは答えがイエスであるかもしれないと信じていましたが、信じるだけでは十分ではありませんでした。価値を迅速に明らかにする必要がありました。
そこで私たちは証明への最短の道を選びました。私たちは、大まかではあるが説得力のあるバージョンを作成し、顧客の前に置いて、反応を観察しました。このプロトタイプは、まさに必要なことを実行しました。これにより、値が読みやすくなりました。エージェントが普段忘れてしまうようなことを覚えているのを人々が見ると、会話はすぐに変わりました。私たちはもはや、興味深い機能について話していません。私たちは市場の準備が整った製品について話していました。
その初期の瞬間がその後のすべてを形作りました。 mem9 は、抽象的な位置付けではなく、ワークフローの苦痛から生まれたため、私たちにとっては常にエージェント時代の製品のように感じられました。それはほぼ即座に検証され、有効になると

日付が変わって、ペースが変わりました。プロジェクトは実験のように振る舞うことをやめ、スタートアップのように振る舞うようになりました。
最初の数日間で、Go サーバー、メモリ API、ストレージ用の TiDB クラウド、検索、認証、レート制限、および最初のプラグイン統合など、システムのコアを驚くほど早く組み立てました。そのほぼ直後に、サポートは OpenClaw、OpenCode、Claude Code などのエージェント環境全体に拡大し、オンボーディングが改善され、マルチテナント基盤が確立され、最初の mem9.ai サイトが稼働しました。私たちはインフラから製品、そして成長へときちんとした順序に従っていませんでした。実際には、これらすべてのトラックが同時に動いていました。なぜなら、価値が明らかになると、ためらいは勢いよりも高価になるからです。
メモリはストレージ機能ではありません
早い段階で 1 つのことが明らかになりました。それは、私たちが「エージェント用のベクトル データベース」を構築しようとしていたわけではないということです。私たちは実際にエージェントの動作を改善するメモリを構築しようとしていました。
これはフレームの小さな変更であり、建築的には非常に大きな影響を及ぼします。
エージェントのメモリに関する多くの議論は、依然として問題をストレージと取得として捉えています。実際には、そのフレーミングは浅すぎます。難しいのは、情報を保存できるかどうかではありません。難しいのは、実際の生産上の制約のもとで、適切な情報が適切なタイミングで、適切な量で返されるかどうかです。
思い出しが少なすぎると、エージェントは重要な 1 つの詳細を忘れてしまいます。思い出しすぎると、文脈が無関係な荷物で汚染されてしまいます。記憶コーパスが成長するにつれて想起が騒々しくなると、信頼は失われます。したがって、課題は持続性そのものではありません。課題は精度です。
この洞察により、mem9 は急速に基本的なメモリ ストアを超えました。耐久性のあるメモリとして始まったものは、すぐに取り込み、抽出、調整、ランキンのためのより独自のシステムになりました。

g、および検索。私たちがサーバー中心のアーキテクチャに移行したのは、メモリ ロジックを一元的に進化させながら、統合を薄いままにしたいと考えたからです。その決断は重要でした。これにより、すべてのプラグインやランタイムに複雑さを押し付けるのではなく、コアで動作を改善できます。
これは、まだ過小評価されていると私たちが考えるカテゴリーの一部です。メモリの品質は主に UI の問題ではなく、純粋にモデルの問題でもありません。これは、ストレージ、ランキング、評価、遅延、製品の判断、オーケストレーションの交差点に位置するエンジニアリングの問題です。エージェントが運用環境で有意義な作業を行う場合、単に追加のコンテキストが必要というわけではありません。彼らにはより良いコンテキストが必要です。
API はエージェント メモリ製品ではありません
次のレッスンはすぐに来ました。メモリ API だけでは製品ではありません。
人々は記憶だけが存在することを望んでいるわけではありません。彼らはそれを見て、検査し、信頼し、修正し、最終的には形を整えたいと考えています。それが mem9 をインフラストラクチャの枠を超えたものに押し上げたものです。
mem9 の次のフェーズは、目に見えないバックエンド機能をユーザーが実際に体験できるものに変えることでした。私たちは、セッション ビュー、タイムライン ビュー、分析ワークフロー、フィルター、プレビュー、洞察レイヤーなど、記憶を読みやすくするサーフェスを構築しました。これらは、人々が何を記憶したかだけでなく、それがなぜ重要なのかを理解するのに役立ちました。その作品は徐々に、単なる UI としてではなく、長期記憶を抽象的ではなく具体的​​に感じさせる方法として「Your Memory」になりました。
バックエンドでは、その変化により、異なる種類のエンジニアリングが必要になりました。作業は、分類、分析品質、重複排除、応答性、およびレポート ワークフローの改善に向けて進められました。どれも最初のスプリントほどのドラマはありませんでしたが、同じくらい重要でした。最初のフェーズでは、記憶が機能することが証明されました。このフェーズにより、理解が深まり、信頼できるものになりました

あなた。
同時に、私たちは好奇心を採用に変える、あまり魅力的ではない部分もすべて構築していました。公開 Web サイト、ドキュメント、分析、アトリビューション、問い合わせフロー、オンボーディングの改善、そして最終的には API ドキュメント。これらの変更はいずれも、コミット ログでは特に映画的なものではありませんが、実際の製品が成長する過程です。 1 回の劇的な立ち上げから成長がもたらされることはほとんどありません。多くの場合、それは摩擦を軽減し、価値を把握しやすくし、興味のあるユーザーがアクティブ ユーザーになるのに役立つ数十の小さな改善から生まれます。
技術的な深さと製品の洗練の組み合わせが重要でした。 mem9 は、高速プロトタイプから、人々が発見、評価し、真剣に使用できる製品に移行しました。 2 週間あまりのうちに、ユーザー数はすでに 10,000 人を超えました。
エージェントメモリ製品の評価版を作成しました
ユーザーが実際のワークフローで記憶に依存し始めると、直感だけでは十分ではなくなります。
「感触が良くなった」というのは良い出発点ですが、それはオペレーティング システムではありません。私たちは、再現の質が向上しているのか、低下しているのか、それとも単に形状が変化しているのかを測定する方法を必要としていました。それが私たちをベンチマークにさらに深く引き込んだのです。
ベンチマークを副次的な研究として扱うのではなく、製品インフラストラクチャとして扱いました。私たちは評価ハーネスを構築し、古いマルチターン データセットをより最新のエージェント設定に適応させ、実際のエンジニアリング上の決定を導くことができるフィードバック ループを作成しました。重要なのは、パフォーマンスのベンチマークを追うことではありません。重要なのは、メモリの品質を可視化してデバッグ可能にすることでした。
mem9が、特にキミに関して、より要求の厳しい会話やパートナーシップに参加するにつれて、その区別はさらに重要になりました。システムが実際のエージェント ワークフローの本格的な長期記憶層として評価されるようになると、曖昧な主張は役に立たなくなります。理解するにはベースラインと証拠が必要です

そして、検索がどこで機能し、どこで失敗するのか、そして変更が精度、再現率、重複、証拠の質にどのような影響を与えるのか。
その意味で、ベンチマークは学術的なスコアリングではなく、製品の真実性を評価する手段のようなものになりました。それは私たちが味を超えて反復に進むのに役立ちました。これにより、「記憶力の低下」を診断可能で改善可能なものに変える方法が得られました。
エージェントの記憶は人間的である必要がある
mem9 を構築する際のより興味深い教訓の 1 つは、メモリを純粋に目に見えないままにしておくべきではないということでした。
API が重要です。ストレージ モデルが重要です。ランキングのロジックは重要です。しかし、ユーザーは記憶を指標として経験しているわけではありません。彼らはそれを連続性として経験します。彼らは、システムが自分のことを知っていると感じるかどうか、時間の経過とともにスレッドを再接続できるかどうか、そしてその継続性が不気味ではなく信頼できると感じるかどうかを重視します。
これが、私たちが API レイヤーにとどまらず、視覚化とメモリ管理に投資し続けた理由の 1 つです。 mem9 の最も特徴的なアイデアのいくつかが、アーキテクチャ図ではなく製品の直感から生まれた理由でもあります。
良い例は、ビジュアル メモリ エクスプローラーである Memory Farm です。表面的には遊び心があるように見えます。ピクセルアートにインスピレーションを得たインターフェイスで、思い出が庭の植物のように成長し、トピックごとにクラスター化され、関係性によって結びつけられます。根底にある本能は深刻です。ユーザーがパターン、クラスター、履歴、関係をより直感的な形式で確認できると、記憶が理解しやすくなります。エージェントがユーザーとどのように関係するかにおいてメモリが中心である場合、メモリ製品はデフォルトで冷たく感じられるべきではありません。
この教訓は、私たちが予想していた以上に製品を形作りました。目的は単に事実を検索することではありません。それは、人々が自分に代わって記憶するシステムに対する信頼を築くのを助けるためでした。
問題が現実であるため、カテゴリが混雑しています
おうから

一方で、エージェントの記憶は注目のカテゴリーのように見えるかもしれません。内側から見ると、ハード エッジ ケースの長いリストのように見えます。
大きなコンテキスト ウィンドウは依然として有限です。重要な事実が最近の騒音に埋もれてしまいます。単純な検索では間違ったものが戻ってきます。繰り返すとトークンが無駄になります。メモリが増えると品質が低下します。そして、思い出すことがランダムであると感じ始めると、ユーザーはすぐに自信を失います。
mem9 は、初日からこれらの問題を考慮して構築されました。だからこそ、この製品は生の永続性から、取り込み、調整、ハイブリッド取得、ランキング、分析、ベンチマーク、オーケストレーションへと迅速に移行しました。市場の注目は本物ですが、それは非常に現実的な製品ニーズの下流にあります。本格的なエージェントを構築している人は皆、遅かれ早かれ同じ失敗に遭遇します。
それが、エコシステムの変化が私たちにとって非常に重要である理由でもあります。エージェント フレームワークによってコンテキストの組み立て方法に関するライフサイクル制御が改善されると、メモリはサイドカーのように見えなくなり、コンテキスト パイプラインのコア部分のように見えるようになりました。それがこのカテゴリがさらに面白くなるポイントです。最高のメモリ システムとは、最も多くのデータを保存できるメモリ システムではありません。これは、エージェントが何を残すべきか、何を表面化すべきか、何を残すべきかを決定するのに役立ちます。

[切り捨てられた]

## Original Extract

How a customer ask in March 2026 became mem9, an agent memory product for coding agents. Lessons on prototyping, retrieval, and evaluation.

Agent Memory Product: How We Built mem9 on TiDB Cloud
Millions of agent branches. One database. Join us at TiDB SCaiLE Europe - June 4, 2026. Register Now
AI
TIDB FOR AGENTIC AI
Database for Agentic AI
Purpose-built for agent memory, state, and multi-hop reasoning
Vector Search & RAG
Native vector indexing and retrieval-augmented generation pipelines
Quick Start: Agentic Memory
Spin up persistent agent memory in seconds — zero config
Build AI Applications
SDKs, guides, and templates for shipping AI apps fast
Product
An open-source distributed SQL database trusted by innovators to power transactional, AI, and other modern applications.
Trusted and verified by innovation leaders around the world.
Explore how TiDB ensures the confidentiality and availability of your data.
An open-source distributed SQL database trusted by innovators to power transactional, AI, and other modern applications.
Trusted and verified by innovation leaders around the world.
Explore how TiDB ensures the confidentiality and availability of your data.
How We Built mem9: Lessons From Shipping Persistent Memory for AI Agents
mem9 started as a customer request in March 2026, not a roadmap. We shipped a prototype before we wrote a plan.
Agent memory is not a storage problem. It is an engineering problem at the intersection of ingestion, ranking, evaluation, and product judgment.
A memory API alone is not a product. People want to see, inspect, trust, and correct what an agent remembers.
mem9 runs on TiDB Cloud, the same substrate behind TiDB Cloud Zero.
In early March 2026, a customer asked us for something that sounded simple and turned out to be one of the hardest problems in the agent stack: Make agents remember .
We did not start with a polished roadmap, a heavyweight architecture review, or a six-month product plan. We started the way many products do: With a concrete user pain, a rough prototype, and a very short distance between “this is interesting” and “we need to ship this.”
That was the beginning of mem9 .
Looking back now, mem9 feels less like a conventional software project and more like a compressed startup year. What began as a fast customer-driven prototype quickly became a product, then a platform, and then a much deeper exploration of what agent memory actually requires in production. The visible features changed quickly, but the core question stayed the same. How do you help an agent remember what matters, without overwhelming it with everything else?
This is what we have learned so far.
It Started With a Real Problem, Not a Market Thesis
The real beginning of mem9 was not a category map or a strategy deck. It was a customer asking a practical question. If an agent could keep durable memory across sessions, would the user actually feel the difference?
We believed the answer might be yes, but belief was not enough. We needed to make the value obvious, fast.
So we took the shortest path to proof. We built a rough but convincing version, put it in front of a customer, and watched for the reaction. That prototype did exactly what it needed to do. It made the value legible. Once people could see an agent remember something it would normally forget, the conversation changed immediately. We were no longer talking about an interesting capability. We were talking about a product the market was ready for.
That early moment shaped everything that followed. mem9 has always felt like an agent-era product to us because it was born from workflow pain rather than abstract positioning. It was validated almost immediately, and once it was validated, the pace changed. The project stopped behaving like an experiment and started behaving like a startup.
In the first few days, we assembled the core of the system surprisingly quickly: A Go server, memory APIs, TiDB Cloud for storage, search, auth, rate limiting, and the first plugin integrations. Almost immediately after that, support expanded across agent environments such as OpenClaw, OpenCode, and Claude Code, while onboarding improved, multi-tenant foundations landed, and the first mem9.ai site went live. We were not following a neat sequence from infra to product to growth. In reality, all of those tracks were moving at once, because once the value was obvious, hesitation became more expensive than momentum.
Memory Is Not a Storage Feature
Early on, one thing became clear: We were not trying to build “a vector database for agents.” We were trying to build memory that actually improves agent behavior.
That is a small change in framing with very large architectural consequences.
A lot of discussions about agent memory still frame the problem as storage plus retrieval. In practice, that framing is too shallow. The hard part is not whether information can be stored. The hard part is whether the right information comes back at the right time, in the right amount, under real production constraints.
Too little recall and the agent forgets the one detail that matters. Too much recall and the context gets polluted with irrelevant baggage. If recall becomes noisy as the memory corpus grows, trust disappears. So the challenge is not persistence by itself. The challenge is precision.
That insight pushed mem9 very quickly beyond a basic memory store. What started as durable memory soon became a more opinionated system for ingestion, extraction, reconciliation, ranking, and retrieval. We moved toward a server-centric architecture because we wanted integrations to stay thin while the memory logic could evolve centrally. That decision mattered. It let us improve behavior at the core instead of pushing complexity into every plugin or runtime.
This is the part of the category that we think is still underestimated. Memory quality is not mainly a UI problem, and it is not purely a model problem either. It is an engineering problem that sits at the intersection of storage, ranking, evaluation, latency, product judgment, and orchestration. If agents are going to do meaningful work in production, they do not just need more context. They need better context.
An API Is Not a Agent Memory Product
The next lesson came quickly. A memory API alone is not a product.
People do not just want memory to exist. They want to see it, inspect it, trust it, correct it, and eventually shape it. That is what pushed mem9 beyond infrastructure.
The next phase of mem9 was about turning an invisible backend capability into something users could actually experience. We built surfaces that made memory legible: Session views, timeline views, analysis workflows, filters, previews, and insight layers that helped people understand not only what had been remembered, but also why it mattered. That work gradually became “Your Memory,” not just as a UI, but as a way to make long-term memory feel concrete instead of abstract.
On the backend, that shift demanded a different kind of engineering. The work moved toward taxonomy, analysis quality, deduplication, responsiveness, and better report workflows. None of that had the drama of the first sprint, but it was just as important. The first phase proved that memory could work. This phase made it understandable and trustworthy.
At the same time, we were building all the less glamorous pieces that turn curiosity into adoption. The public website, docs, analytics, attribution, contact flows, better onboarding, and eventually API documentation. None of those changes are especially cinematic in a commit log, but they are how real products grow. Growth rarely comes from one dramatic launch. More often, it comes from dozens of small improvements that reduce friction, make the value easier to grasp, and help interested users become active users.
That combination of technical depth and product polish mattered. mem9 moved from a fast prototype to a product people could discover, evaluate, and use seriously. Within a little over two weeks, it had already crossed 10,000 users.
We Made Evaluation Part of the Agent Memory Product
Once users start relying on memory in real workflows, intuition is no longer enough.
“It feels better” is a good starting point, but it is not an operating system. We needed ways to measure whether recall quality was improving, regressing, or simply changing shape. That is what pulled us deeper into benchmarks.
Instead of treating benchmarks as side research, we treated them as product infrastructure. We built evaluation harnesses, adapted older multi-turn datasets into more modern agent settings, and created feedback loops that could guide actual engineering decisions. The point was not to chase performance benchmarks. The point was to make memory quality visible and debuggable.
That distinction mattered even more as mem9 entered more demanding conversations and partnerships, especially around Kimi. Once your system is being evaluated as a serious long-term memory layer for real agent workflows, vague claims stop being useful. You need baselines and evidence, to understand where retrieval works, where it fails, and how changes affect precision, recall, duplication, and evidence quality.
In that sense, benchmarking became less like academic scoring and more like instrumentation for product truth. It helped us move beyond taste and into iteration. It gave us a way to turn “memory feels off” into something diagnosable and improvable.
Agent Memory Has to Feel Human
One of the more interesting lessons in building mem9 was that memory should not remain purely invisible.
The APIs matter. The storage model matters. The ranking logic matters. But users do not experience memory as an index. They experience it as continuity. They care about whether the system feels like it knows them, whether it can reconnect threads over time, and whether that continuity feels trustworthy rather than uncanny.
That is part of why we kept investing in visualization and memory management instead of stopping at an API layer. It is also why some of the most distinctive ideas in mem9 came from product intuition rather than architecture diagrams.
A good example is Memory Farm , our visual memory explorer. On the surface, it looks playful: A pixel-art-inspired interface where memories grow as plants in a garden, clustered by topic and connected by relationship. The underlying instinct is serious. Memory becomes easier to understand when users can see patterns, clusters, history, and relationships in more intuitive forms. If memory is central to how an agent relates to a user, then memory products should not feel cold by default.
That lesson shaped more of the product than we expected. The goal was never just to retrieve facts. It was to help people build trust in a system that remembers on their behalf.
The Category Is Crowded Because the Problem Is Real
From the outside, agent memory can look like a hot category. From the inside, it mostly looks like a long list of hard edge cases.
Large context windows are still finite. Important facts get buried under recent noise. Naive retrieval brings back the wrong things. Repetition wastes tokens. Quality degrades as memory grows. And once recall starts to feel random, users lose confidence very quickly.
mem9 was built inside those problems from day one. That is why the product moved so quickly from raw persistence into ingestion, reconciliation, hybrid retrieval, ranking, analysis, benchmarking, and orchestration. The market attention is real, but it is downstream of a very real product need. Everyone building serious agents runs into the same failures sooner or later.
That is also why ecosystem shifts mattered so much to us. As agent frameworks introduced better lifecycle control around how context is assembled, memory stopped looking like a sidecar and started looking like a core part of the context pipeline. That is the point where the category becomes much more interesting. The best memory system is not the one that stores the most. It is the one that helps an agent decide what should stay, what should surface, and what should remain q

[truncated]
