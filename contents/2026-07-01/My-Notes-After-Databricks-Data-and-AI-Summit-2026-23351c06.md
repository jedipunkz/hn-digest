---
source: "https://zilliz.com/blog/databricks-data-ai-summit-2026-data-layer"
hn_url: "https://news.ycombinator.com/item?id=48749195"
title: "My Notes After Databricks Data and AI Summit 2026"
article_title: "Notes from Databricks Summit 2026: Why Data Layer Matters Again - Zilliz blog"
author: "redskyluan"
captured_at: "2026-07-01T16:46:41Z"
capture_tool: "hn-digest"
hn_id: 48749195
score: 1
comments: 0
posted_at: "2026-07-01T16:11:07Z"
tags:
  - hacker-news
  - translated
---

# My Notes After Databricks Data and AI Summit 2026

- HN: [48749195](https://news.ycombinator.com/item?id=48749195)
- Source: [zilliz.com](https://zilliz.com/blog/databricks-data-ai-summit-2026-data-layer)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T16:11:07Z

## Translation

タイトル: Databricks Data and AI Summit 2026 後の私のメモ
記事のタイトル: Databricks Summit 2026 のメモ: データ レイヤーが再び重要になる理由 - Zilliz ブログ
説明: James Luan が、本番 AI がデータ層をインフラストラクチャの中心に押し戻す理由について、Databricks Data + AI Summit 2026 のメモを共有します。

記事本文:
Databricks Summit 2026 のメモ: データ レイヤーが再び重要になる理由 - Zilliz ブログ 製品 Zilliz Cloud
大規模な高い信頼性、パフォーマンス、コスト効率を実現するように設計された、フルマネージドの Vector Lakebase サービス。
10 億規模のベクトル類似性検索用に構築されたオープンソースのベクトル データベース。
Zilliz Vector Lakebase BYOC を発表
価格設定 料金プラン あらゆる予算に合わせて、あらゆるチームに柔軟な価格設定オプションを提供します。
計算機 コストを見積もります。
定価 すべての請求項目の詳細な定価を表示します。
無料利用枠 マネージド Milvus の力で想像力を解き放ちます。
Zilliz Cloud Developer Hub では、Zilliz Cloud を使用するためのすべての情報が見つかります。
リソース ブログ ガイド リサーチ アナリスト レポート ウェビナー トレーニング
顧客検索 拡張生成 すべてのユースケースを見る 業界別で見る すべての顧客事例を見る OpenEvidence が Zilliz Cloud で医療 AI を強化
English 日本語 한국어 Español Français Deutsch Italiano Português Русский 45k デモを予約 ログイン 無料で始める
Databricks Data + AI Summit 2026 からのいくつかのメモ: データ レイヤーが再び重要になる理由
ページをコピー Databricks Data + AI Summit 2026 からのいくつかのメモ: データ レイヤーが再び重要になる理由
データ: 市場がまだ価格を設定していない AI スタックの部分
AI エージェントによりデータの問題を隠すことが不可能になる
Databricks は正しい問題を目指しています
地図は良いです。しかし、まだ終わっていません。
データベース利用者がエージェントの場合
「AIネイティブ」とは実際に何を意味するのか
最終インターフェイスとして SQL だけでは不十分
もう 1 つ: Zilliz Vector Lakebase はパブリック プレビューで利用可能です
GenAI アプリケーション用に構築されたフルマネージドのベクター データベースをお試しください。
今年の Databricks Data + AI Summit の後、私は 1 つの発表について考えることよりも、次のような質問について考えるようになりました。

しばらく私と一緒に座っています。
AI が実際に本番環境に移行すると、データ層はどうなるでしょうか?
私の現在の答えは単純ですが、意味は異なります。このサイクルでは、データ層は AI スタックの中で価格変更が最も遅い部分です。それが変わり始めています。
データ: 市場がまだ価格を設定していない AI スタックの部分
アルゴリズムの価格は公開で再設定されています。モデルは急速に改良され、業界はほぼ毎週その進歩を確認できます。コンピューティングの価格は、NVIDIA、クラウド プロバイダー、資本市場によって再設定されています。 GPU が重要であることは誰もが理解しています。
データの移動が遅くなりました。重要性が低いからではありません。その逆も真実です。データは話しにくく、修正するのがさらに難しいため、価格の再設定が遅れます。企業データは乱雑で、散在し、重複し、古くなり、誰も完全には理解できない権限でいっぱいです。ビジネス セマンティクスはシステム間できれいに一致しません。人々が「リアルタイム」と呼ぶものは、多くの場合、昨夜のある時点で実行されたスケジュールされたジョブのままです。
その作業は苦痛です。あまり華やかでもありません。しかし、AI がデモから本番環境に移行すると、痛みを隠すことはできなくなります。
OpenAI や Anthropic など、モデルを構築およびトレーニングしている人々との会話では、議論が同じ点に戻ることがよくあります。モデルは収束しつつあります。少なくとも十分なお金があれば、コンピューティングを購入できます。防御可能な層はますますデータになりつつあります。データの品質、鮮度、アクセス許可、データを有用なコンテキストに変換する速度です。
これはアプリケーション層だけの問題ではありません。モデル企業内では、モデルの品質は依然としてデータ パイプラインに大きく依存しています。トレーニングの実行には、最初の本格的な実験を開始する前に数日の準備が必要になる場合があります。もし

上流のフィールドが汚れていたり、バッチのラベル​​が間違っていたり、フィルタリング ルールが間違っていたりすると、損失曲線のドリフトに誰かが気づく前に、何日もの計算と待ち時間が消えてしまう可能性があります。
AI エージェントによりデータの問題を隠すことが不可能になる
エージェントは、同じ問題をより運用的な形式で明らかにします。
AI エージェントが実稼働環境で失敗する場合、最初の原因はモデルが機能しないことではないことがよくあります。それは、モデルが間違ったコンテキストに基づいて動作しているということです。アクセスできないレコード、6 時間前に期限切れになったドキュメント、一晩で静かに変更されたデータ ソース、または頻繁に使用するには高価すぎる取得パスなどです。私は最近、強力なチームが古いコンテキスト パイプラインのせいで 1 週間近く損失を被っているのを目にしました。エージェントは昨日の質問に自信を持って答えていました。モデルはバカではなかった。コンテキストが間違っており、システムにはエラーがループに入った場所を証明する明確な方法がありませんでした。
それが重要な故障モードです。次のインフラストラクチャのボトルネックは、単により適切な推論ではありません。モデルまたはエージェントが決定を下す瞬間の、新鮮で信頼でき、安価で監査可能なコンテキストです。
だからこそ、AI スタックの中で次に価格が変更されるのはデータ層だと私は考えています。
Databricks は正しい問題を目指しています
私は「AI データ プラットフォーム」を名乗る多くの製品には懐疑的です。ストーリーがシステムよりも先に到着することがあまりにも多いのです。
Databricks は十分に異なっているため、真剣に注目する価値があると思います。サミットで私にとって印象に残ったことが 2 つありました。
1 つ目はやはりエンジニアリング文化です。 Databricks の規模では、会社が純粋に販売主導型になるのは簡単です。それでも、創設者たちは依然としてステージ上で、実行エンジン、トランザクション、リアルタイム分析、製品の基盤となるパイプについて話しています。私はそれを尊敬します。企業がまだ製品とエンジニアリングの直感を持っていることを実感できます

その核心部分にある。それは基調講演に現れるずっと前に、小さなアーキテクチャ上の決定に現れます。
2つ目は顧客基盤です。サミットで私が話したユーザーは、デモ層としての AI について話していませんでした。彼らは AI を実稼働システムに導入しようとしていましたが、彼らが説明した問題はより具体的でした。エージェントはビジネス状態を読み書きする必要があります。リアルタイム分析では、データの移動による税金を払い続けることはできません。パイプラインはより自律的になる必要があります。エージェントの動作には、事後だけでなく実行時にもガバナンスが必要です。
だからこそ、Lakebase、Lakehouse//RT、データ エージェント、AI ガバナンスなどの発表が重要なのです。名前は方向ほど重要ではありません。トランザクションを湖の近くに置きます。リアルタイム分析を同じデータ基盤に戻します。パイプラインをさらに自動化します。ガバナンスを「誰がこのデータセットを参照できるか」から「このエージェントはこの特定のステップで何を実行できるか」まで拡張します。
私はそれが間違った方向だとは思いません。これは、私たちの多くが同じ未来をさまざまな角度から見ている証拠だと私は考えています。
データベースは拡大中です。もはや、データを保存したりクエリしたりするだけの場所ではありません。それは、事実、状態、意味論、ガバナンス、およびアクションの基盤になりつつあります。
地図は良いです。しかし、まだ終わっていません。
Databricks はまさにその方向に向かっています。それは、アーキテクチャが最終形に到達したことを意味するものではありません。
写真: The Known Data Realm · Databricks Data + AI Summit 2026
地図がまだ不完全な地域が 3 つあります。
Postgres から始めるのが賢明なエントリーポイントです。開発者はそれを知っています。生態系は巨大です。互換性により導入の摩擦が軽減されます。それは重要です。
しかし、人々を惹きつけるアーキテクチャが、必ずしも最終的なワークロードを勝ち取るアーキテクチャであるとは限りません。
AI 時代の運用システムにはトランザクション、メモリ、ベクトルが必要

rs、マルチモーダル データ、トレース、分岐、ロールバック、および非常にきめ細かいテナント分離。従来のリレーショナル コアは、拡張機能や周囲のサービスを通じてこれらの一部を公開できますが、それがネイティブになるわけではありません。 Classic Postgres は、クラウドネイティブの分散スケール向けに設計されておらず、また、有効期間の短いデータベースの作成、状態フォーク、メモリへの書き込み、トレースの生成、消滅を行うエージェント向けにも設計されていません。
Postgres をオブジェクト ストレージに近づけても、これらの疑問が消えるわけではありません。オブジェクト ストレージは安価で信頼性がありますが、デフォルトでは低遅延ではありません。高速に感じるには、積極的かつ正確なキャッシュ層が必要です。実際のトランザクション負荷の下でもキャッシュが安定していることは、データベースにおける最も困難なシステム問題の 1 つです。したがって、Lakebase に関する私の正直な疑問は、デモが印象的かどうかではありません。重要なのは、そのキャッシュを午前 3 時に人々を目覚めさせるものにせずに、システムが本番規模で実際の OLTP ワークロードを処理できるかどうかです。
Databricks は、OLTP、ウェアハウス、リアルタイム分析、データ サイエンス、ガバナンスにわたる強力なマップを作成しました。しかし、AI アプリケーションはテキスト、画像、音声、ビデオ、埋め込み、動作ログ、エージェント トレースをますます消費します。これらはテーブルの隣に置かれている単なる物ではありません。これらは、エージェントが取得し、推論し、変換し、書き戻すデータです。
マルチモーダル データがコア マップの外側に残っている場合、最も重要な AI データ資産は依然としてマージンに存在します。
ダッシュボード、自然言語 BI、Excel スタイルのワークフロー、アナリスト向けのエクスペリエンスなど、製品の表面の多くは依然として人間のユーザーを想定しています。それらは貴重なものです。ただし、エージェントはデータベースの使用方法が異なります。
エージェントがダッシュボードを開くのは 1 日に 1 回ではありません。ループ内で実行されます。コンテキストを取得し、決定を下し、ツールを呼び出し、状態を書き込み、ポリシーをチェックし、再実行します。

泥炭。すべてのステップを監査する必要がある場合があります。すべての取得は次のアクションに影響を与える可能性があります。すべての書き込みにはロールバックが必要になる場合があります。すべての権限チェックは実行時に行う必要がある場合があります。
それは別のデータベース ワークロードです。
写真: Unity AI ゲートウェイ・ガバナンス —— Databricks Data + AI Summit 2026
データベース利用者がエージェントの場合
何十年もの間、データベースは主に、このクエリを正しく迅速に実行する方法という 1 つの問題に集中することができました。
エージェントの時代では、問題はより広範囲になります。
エージェントは、決定を下す瞬間に、最も新鮮で、最も信頼でき、最もコストが低く、最も監査可能なコンテキストをどのようにして取得するのでしょうか?
これはクエリ最適化だけの問題ではありません。これは、ストレージ、インデックス作成、ガバナンス、リネージ、リプレイ、コスト管理、実行時ポリシーの適用にわたるシステムの問題です。
ここからカテゴリーが変化し始めます。データ システムはもはや単なるインテリジェンス システムではありません。質問すると答えが返されます。 AI のオペレーティング システムに近づく必要があります。つまり、エージェントがコンテキストを読み取り、意思決定を行い、ツールを呼び出し、状態を書き込み、人間や他のシステムが検査できる痕跡を残す場所です。
監査可能性を事後的に強化することはできません。エージェントが間違った答えを出したり、間違った行動をとったり、多額の費用を費やしたりした場合、最初に疑問になるのは、「その瞬間、エージェントは正確に何を見たのか?」ということです。
それに答えるために、システムは、どのドキュメントが取得されたか、どのベクトルが照合されたか、どのメタデータ フィルターが適用されたか、どのリランカーが順序を変更したか、どのツールが呼び出されたか、どのポリシーが適用されたか、どの状態が書き戻されたかを知る必要があります。デバッグとガバナンスは同じワークフローになります。
これは、まだ誰も完全には解決していないアーキテクチャだと思います。
「AIネイティブ」とは実際に何を意味するのか
「AIネイティブ」は、

ほぼ何でも意味できます。まだ明確な定義はないと思います。しかし、実際のエージェントのワークロードから逆算すると、AI ネイティブのデータ システムは少なくともいくつかのことを適切に実行する必要があります。
マルチモーダルデータは最高級でなければなりません
テキスト、画像、オーディオ、ビデオ、埋め込み、ログ、およびトレースは、リレーショナル テーブル、ベクトル列、オブジェクト バケット、およびいくつかのサイド インデックスに分散してはいけません。これらは、検索、フィルタリング、ランキング、ガバナンスが同時に行われる 1 つの論理システム内に存在する必要があります。
難しいのは、これらの資産を保存しないことです。難しいのは、アーキテクチャを別のパイプライン問題に変えることなく、それらをまとめてクエリできるようにすることです。
弾力性はワークロードから始める必要がある
エージェントのトラフィックが急増しています。システムは 1 時間静止した後、大量の取得、メモリ、およびツール使用のリクエストを受信する場合があります。データ レイクまたはオブジェクト ストアは、安価で信頼性が高く、コンピューティングから切り離された耐久性のある基盤となる必要があります。
しかし、コーパスが存在するからといって、コンピューティングが高価なままであってはなりません。誰も検索していない場合、システムはほとんど費用を費やさないはずです。ワークロードが起動すると、コンピューティングがすぐに到着するはずです。その世界では、自然な価格設定単位は必ずしも永続的なクラスターであるとは限りません。それはクエリ、セッション、またはアクティブなコンピューティングの時間である可能性があります。

[切り捨てられた]

## Original Extract

James Luan shares notes from Databricks Data + AI Summit 2026 on why production AI is pushing the data layer back to the center of infrastructure.

Notes from Databricks Summit 2026: Why Data Layer Matters Again - Zilliz blog Products Zilliz Cloud
Fully managed Vector Lakebase services designed for high reliability, performance, and cost efficiency at scale.
Open-source vector database built for billion-scale vector similarity search.
Announcing Zilliz Vector Lakebase BYOC
Pricing Pricing Plan Flexible pricing options for every team on any budget.
Calculator Estimate your cost.
List Price View detailed list prices for every billing item.
Free Tier Unleash Your Imagination with the Power of Managed Milvus.
The Zilliz Cloud Developer Hub where you can find all the information to work with Zilliz Cloud.
Resources Blog Guides Research Analyst Reports Webinars Trainings
Customers Retrieval Augmented Generation View all use cases View by industry View all customer stories OpenEvidence Powers Medical AI with Zilliz Cloud
English 日本語 한국어 Español Français Deutsch Italiano Português Русский 45k Book a Demo Log In Get Started Free
A Few Notes from Databricks Data + AI Summit 2026: Why the Data Layer Matters Again
Copy page A Few Notes from Databricks Data + AI Summit 2026: Why the Data Layer Matters Again
Data: the part of the AI stack the market has not priced yet
AI agents make the data problem impossible to hide
Databricks is aiming at the right problem
The map is good. But it is not finished.
When the database user is an agent
What “AI-native” should actually mean
SQL is not enough as the final interface
One more thing: Zilliz Vector Lakebase is available in public preview
Try the fully-managed vector database built for your GenAI applications.
After this year’s Databricks Data + AI Summit , I found myself thinking less about any single announcement and more about a question that has been sitting with me for a while:
When AI really moves into production, what does the data layer become?
My current answer is simple, though the implications are not: in this cycle, the data layer is the part of the AI stack that has been repriced the slowest. That is starting to change.
Data: the part of the AI stack the market has not priced yet
Algorithms have been repriced in public. Models improve quickly, and the industry can see the progress almost every week. Compute has been repriced by NVIDIA, the cloud providers, and the capital markets. Everyone understands that GPUs matter.
Data has moved more slowly. Not because it matters less. The opposite is true. Data is slow to reprice because it is hard to talk about and even harder to fix. Enterprise data is messy, scattered, duplicated, stale, and full of permissions that nobody fully understands. Business semantics do not line up cleanly across systems. The thing people call “real time” is often still a scheduled job that ran sometime last night.
That work is painful. It is also not very glamorous. But once AI moves from demos into production, the pain becomes impossible to hide.
In conversations with people building and training models, including those at OpenAI and Anthropic, the discussion often comes back to the same point. Models are converging. Compute can be bought, at least if you have enough money. The defensible layer is increasingly becoming the data: the quality of it, the freshness of it, the permissions around it, and the speed at which it can be turned into useful context.
This is not only an application-layer problem. Inside model companies, model quality still depends heavily on the data pipeline. A training run may require days of preparation before the first serious experiment begins. If an upstream field is dirty, a batch is mislabeled, or a filtering rule is wrong, days of compute and waiting can disappear before anyone notices the loss curve has drifted.
AI agents make the data problem impossible to hide
Agents expose the same problem in a more operational form.
When AI agents fail in production, the first cause is often not that the model is incapable. It is that the model is acting on the wrong context: a record it cannot access, a document that expired six hours ago, a data source that quietly changed overnight, or a retrieval path that is too expensive to use often enough. I recently saw a strong team lose nearly a week to a stale context pipeline. The agent was confidently answering yesterday’s question. The model was not dumb. The context was wrong, and the system had no clean way to prove where the error entered the loop.
That is the failure mode that matters. The next infrastructure bottleneck is not simply better reasoning. It is fresh, trusted, cheap, and auditable context at the moment a model or agent makes a decision.
That is why I think the data layer is the next part of the AI stack to be repriced.
Databricks is aiming at the right problem
I am skeptical of many products that call themselves “AI data platforms.” Too often the story arrives before the system.
Databricks is different enough that I think it deserves serious attention. Two things stood out to me at the Summit.
The first is still the engineering culture. At Databricks’ scale, it would be easy for the company to become purely sales-driven. Yet the founders are still on stage talking about execution engines, transactions, real-time analytics, and the pipes underneath the product. I respect that. You can feel when a company still has product and engineering intuition at its core. It shows up in small architectural decisions long before it shows up in a keynote.
The second is the customer base. The users I spoke with at the Summit were not talking about AI as a demo layer. They were trying to push AI into production systems, and the problems they described were much more concrete: agents need to read and write business state; real-time analytics cannot keep paying the tax of moving data; pipelines need to become more autonomous; agent behavior needs governance at runtime, not only after the fact.
That is why announcements such as Lakebase, Lakehouse//RT, data agents, and AI governance matter. The names are less important than the direction. Put transactions closer to the lake. Pull real-time analytics back toward the same data foundation. Automate more of the pipeline. Extend governance from “who can see this dataset” to “what is this agent allowed to do in this specific step?”
I do not see that as a wrong turn. I see it as evidence that many of us are looking at the same future from different angles.
The database is expanding. It is no longer only a place to store and query data. It is becoming the foundation for facts, state, semantics, governance, and action.
The map is good. But it is not finished.
Databricks is right in the direction. That does not mean the architecture has reached its final form.
Photo: The Known Data Realm · Databricks Data + AI Summit 2026
I see three areas where the map is still incomplete.
Starting with Postgres is a smart entry point. Developers know it. The ecosystem is huge. Compatibility lowers adoption friction. That matters.
But the architecture that gets people in the door is not always the architecture that wins the final workload.
AI-era operational systems need transactions, memory, vectors, multimodal data, trace, branching, rollback, and very fine-grained tenant isolation. A traditional relational core can expose some of these through extensions and surrounding services, but that does not make them native. Classic Postgres was not designed for cloud-native distributed scale, nor for agents that create short-lived databases, fork state, write to memory, generate traces, and disappear.
Moving Postgres closer to object storage does not erase those questions. Object storage is cheap and reliable, but it is not low-latency by default. To make it feel fast, you need a cache layer that is both aggressive and correct. A cache that stays stable under real transactional load is one of the hardest systems problems in databases. So my honest question about Lakebase is not whether the demo is impressive. It is whether the system can carry real OLTP workloads at production scale without turning that cache into the thing that wakes people up at 3 a.m.
Databricks has drawn a strong map across OLTP, warehousing, real-time analytics, data science, and governance. But AI applications increasingly consume text, images, audio, video, embeddings, behavior logs, and agent traces. Those are not just objects sitting next to tables. They are the data that agents retrieve, reason over, transform, and write back.
If multimodal data remains outside the core map, then the most important AI data assets still live in the margins.
Much of the product surface still assumes a human user: dashboards, natural-language BI, Excel-style workflows, and analyst-facing experiences. Those are valuable. But agents use databases differently.
An agent does not open a dashboard once a day. It runs in a loop. It retrieves context, makes a decision, calls a tool, writes state, checks a policy, and repeats. Every step may need to be audited. Every retrieval may influence the next action. Every write may need rollback. Every permission check may need to happen at runtime.
That is a different database workload.
Photo: Unity AI Gateway · Governance —— Databricks Data + AI Summit 2026
When the database user is an agent
For decades, a database could mostly focus on one question: how to execute this query correctly and quickly.
In the agent era, the question becomes broader:
How does an agent get the freshest, most trusted, lowest-cost, and most auditable context at the moment it makes a decision?
That is not just a query optimization problem. It is a systems problem across storage, indexing, governance, lineage, replay, cost control, and runtime policy enforcement.
This is where the category starts to shift. A data system can no longer be only an intelligence system: you ask a question, it returns an answer. It has to become closer to an operating system for AI: the place where agents read context, make decisions, call tools, write state, and leave behind a trace that humans and other systems can inspect.
Auditability cannot be bolted on after the fact. If an agent gives the wrong answer, takes the wrong action, or spends too much money, the first question will be: what exactly did it see at that moment?
To answer that, the system needs to know which documents were retrieved, which vectors were matched, which metadata filters were applied, which reranker changed the order, which tool was called, what policy was enforced, and what state was written back. Debugging and governance become the same workflow.
That is the architecture I do not think anyone has fully solved yet.
What “AI-native” should actually mean
“AI-native” is becoming one of those phrases that can mean almost anything. I do not think there is a clean definition yet. But if we work backward from real agent workloads, an AI-native data system has to do at least a few things well.
Multimodal data has to be first-class
Text, images, audio, video, embeddings, logs, and traces should not be scattered across a relational table, a vector column, an object bucket, and several side indexes. They need to live in one logical system where retrieval, filtering, ranking, and governance can happen together.
The hard part is not storing these assets. The hard part is making them queryable together without turning the architecture into another pipeline problem.
Elasticity has to start from the workload
Agent traffic is bursty. A system may be quiet for an hour and then receive a flood of retrieval, memory, and tool-use requests. The data lake or object store should become the durable foundation: cheap, reliable, and decoupled from compute.
But compute should not remain expensive just because the corpus exists. If nobody is searching, the system should spend very little. If a workload wakes up, compute should arrive quickly. In that world, the natural pricing unit is not always a permanent cluster. It may be the query, the session, or the minute of active compute.

[truncated]
