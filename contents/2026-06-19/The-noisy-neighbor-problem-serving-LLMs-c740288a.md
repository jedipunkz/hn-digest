---
source: "https://cohere.com/blog/serving-fairness"
hn_url: "https://news.ycombinator.com/item?id=48600482"
title: "The noisy neighbor problem: serving LLMs"
article_title: "LLM Serving Fairness"
author: "_josh_meyer_"
captured_at: "2026-06-19T18:43:55Z"
capture_tool: "hn-digest"
hn_id: 48600482
score: 1
comments: 0
posted_at: "2026-06-19T16:46:58Z"
tags:
  - hacker-news
  - translated
---

# The noisy neighbor problem: serving LLMs

- HN: [48600482](https://news.ycombinator.com/item?id=48600482)
- Source: [cohere.com](https://cohere.com/blog/serving-fairness)
- Score: 1
- Comments: 0
- Posted: 2026-06-19T16:46:58Z

## Translation

タイトル: 騒々しい隣人の問題: LLM の提供
記事のタイトル: 公平性を提供する LLM
説明: Cohere がすべてのテナントに公平なコンピューティングのシェアを確保する方法。

記事本文:
LLM はフェアネス ノース ミニ コードを提供します。 Cohere初の開発者向けモデル。
現代の職場の生産性を高めるエンタープライズ対応の AI プラットフォーム
ビジネスの洞察を明らかにするためのインテリジェントな検索および検出システム
エージェント的、マルチモーダル、多言語 AI 向けの高性能モデル
高精度の音声トランスクリプトを生成するための音声認識モデル
70 以上の言語をカバーする多言語研究モデルのファミリー
最先端のマルチモーダル検索および取得ツール
検索品質をセマンティックに向上させる強力なモデル
専用の安全なモデル推論プラットフォーム — Cohere によって管理
複雑な ML 問題の解決を目指す Cohere の研究ラボ
エンタープライズ AI のケーススタディと成功事例を調べる
現代の職場の生産性を高めるエンタープライズ対応の AI プラットフォーム
ビジネスの洞察を明らかにするためのインテリジェントな検索および検出システム
エージェント的、マルチモーダル、多言語 AI 向けの高性能モデル
高精度の音声トランスクリプトを生成するための音声認識モデル
70 以上の言語をカバーする多言語研究モデルのファミリー
最先端のマルチモーダル検索および取得ツール
検索品質をセマンティックに向上させる強力なモデル
専用の安全なモデル推論プラットフォーム — Cohere によって管理
複雑な ML 問題の解決を目指す Cohere の研究ラボ
エンタープライズ AI のケーススタディと成功事例を調べる
LLM の公平性の提供: 騒々しい隣人はもういらない
Cohere はどのようにしてすべてのテナントが公平なコンピューティングのシェアを確保できるようにしているか
社外サービス担当技術スタッフマネージャー
AIは近道ではありません。
そうやってビジネスは前進するのです。
大規模な言語モデルをマルチテナント SaaS プラットフォームとして実行するには、一見難しい問題が伴います。多くの組織が同じ GPU プールを共有しており、トラフィックが急増し、不均一になります。左

管理されていない場合、ある顧客のスパイクが他のすべての顧客の遅延問題になる可能性があります。
このブログ投稿では、アーキテクチャ パターンと従来のスケジューリング アルゴリズムの組み合わせを使用して、テナント間で公平に推論リクエストをスケジュールするための Cohere の新しいソリューションについて説明します。
推論は、リクエストがバッチで処理される場合に最も効率的です。フルバッチが供給された GPU は高速かつ経済的に動作しますが、一度に 1 つのリクエストを処理する GPU はほとんどアイドル状態になります。そのため、リクエストはハードウェアに到達する前に短時間キューに入れられ、バッチにまとめられます。
問題は注文です。単純なキューを想像してください。優先度と期限だけによって順序付けされた単一の共有回線です。ここで、1 つの組織が一気に 10,000 件のリクエストを送信し、2 つ目の組織が 5 件だけを送信した場合に何が起こるかを考えてみましょう。 1 つのグローバル キューでは、10,000 個のリクエストが前に積み重なり、行儀の良い 5 つのリクエストが後ろで待機します。
これは古典的なノイジーネイバー問題であり、マルチテナント LLM プラットフォームは、これらのトラフィック パターンに直面する他の共有システムと何ら変わりません。
公平性の提供の目標は、テナントを相互に分離することです。その結果、テナントが受け取る推論容量のシェアは、キューをどれだけ積極的にフラッディングするかではなく、公平なスケジューリングに依存します。同時に、バッチ処理の効率を維持しながら、各テナント内の優先順位と期限の順序も維持されます。
Cohereは、4つの異なるメカニズムを組み合わせることにより、テナント全体でワークロードを公平に管理し、それぞれが問題の異なる部分を解決します。これらは固定された順序で実行されます。つまり、 の途中でレート リミッターがアドミッションを制御し、その後 3 つのセレクター (パフォーマンス層、赤字ラウンド ロビン、および優先度) が出口で次のリクエストを選択します。
アーキテクチャと段階的なフローは次のとおりです。
B

リクエストは、スケジューリング キューに加わる前に、アドミッション コントロールを通過します。レート制限により、単一のテナントが特定の時間枠内に送信できる推論リクエストの最大数 (1 分あたり、または 1 か月あたりのリクエストなど) が制限されます。
Cohereでは、これらの制限はエンドポイント・レベルで構成され、各モデルが消費するリソースの数に基づいてモデル間で異なります。たとえば、重量のある生成モデルには、軽量の埋め込みモデルよりも厳しい制限があります。
リアルタイムのスロットル チェックもあります。キューがすでにバックログになっているため、新しいリクエストをレイテンシー目標内で処理できない場合、そのリクエストは単に早期に拒否されます。これにより、システムが受け入れられない作業を受け付けないよう保護され、負荷がかかっても遅延が予測可能に保たれます。
リクエストが許可されると、以下の一連のセレクターに進みます。
最初の選択決定は層です。コンピューティング リソースは、サービス レベル アグリーメント (SLA) に基づいて優先順位が付けられます。つまり、高額な支払いレベルには、低いレベル (またはトライアル レベル) よりも高い処理優先順位と高速なキュー処理が割り当てられます。次に、後者の顧客には、キャパシティが許す限りサービスが提供されます。
重要なのは、階層化によって誰が最初に進むかを決定することです。それ自体は、層内の単一の大規模テナントが支配することを妨げるものではありません。それが次の 2 つの層の目的です。
システムの中心となるのは、Deficit Round Robin (DRR) アルゴリズムです。これにより、階層内のフリート全体にコンピューティングが公平に分散されます。
各テナント (「組織」) は独自の回線を取得します。スケジューラは、最も長い回線を排出するのではなく、テナント間で交代します。各テナントには、自分の番に実行できる作業の少量の予算 (クォンタム) が与えられます。テナントが順番を使用すると、その予算に次のコストが引き落とされます。

送信したばかりのリクエスト。予算を使い果たしたテナントは、次のラウンドで予算が補充されるまでスキップされます。
DRR の優れた点は、作業量を節約し、重み付けできることです。安価なリクエストではテナントの出現頻度が高く、高価なリクエストではテナントの出現頻度が低くなります。そのため、テナントが GPU を使い果たすことはありませんが、容量が無駄になることもありません。前の例に戻ると、組織 A が 10,000 件のリクエストをキューに入れ、組織 B が 5 件だけだったとしても、組織 B はサイクルごとに順番を取得します。組織 A のバーストは組織 B のレイテンシに変換されなくなりました。
このスキームは 2 つの重要な変数に依存します。
Quantum: 各ラウンドでテナントに付与される予算
コスト: 各リクエストが消費する予算
重要な設計上の選択は、これらの変数がどの単位で表現されるかにあります。重要なことに、これはさまざまな推論コンテキストで「公平性」をどのように概念化するかを決定します。 Cohereでは、エンドポイントに応じて、リクエストベースの予算設定とトークンベースの予算設定という2つのコストモデルを使用しています。
簡単にするために、すべてのリクエストのコストは 1 として与えられ、クォンタムはテナントがターンごとに送信できるリクエストの数です。したがって、公平性は純粋にリクエストの数で測定されます。
これは、リクエスト処理コストが大幅に異なる可能性がある、チャットやコンプリーションなどの生成エンドポイントにとっては最適ではありません。 100K トークン プロンプトを含むリクエストは、1K トークン プロンプトを含むリクエストよりも桁違いに多くのリソースを消費する可能性があります。推論可能なモデルの場合、総コストは入力サイズだけでなく、リクエストの特定のコンテキストによってトリガーされる中間推論、計画、および出力生成の量にも依存します。
理想的には、DRR はリクエストのトークン正規化されたサービス コストの EMA (指数移動平均) に基づくフィードバック ループを使用し、観察されたサービスに予算を適応できるようにします。

リソースの消費。静的な予算設定は、エンドポイント内のリクエストのサイズとコストがほぼ同じである場合に最も効果的に機能します。
ここで、リクエストのコストはそのトークン数であり、量子はラウンドごとのトークンバジェットです。公平性は、実際に実行された作業で測定されるようになりました。これは、項目数ではなくバッチのトークン合計が GPU コストを左右する、エンベディングやリランカーなどのバッチ処理されたエンドポイントに自然に適合します。いくつかの非常に大きなドキュメントを送信するテナントは、すぐに予算を使い果たし、より早くフロアを譲ります。小さなリクエストを多数送信するテナントは、リクエストあたりの料金が少なく、より頻繁に表示されます。こうすることで、テナントがいくつかの非常に大きなリクエストを送信して GPU を独占することができなくなります。
したがって、リクエストベースの予算設定では、リクエストがどれほど大きくても、リクエストは競合する各テナントからの (最大で) 一定数のリクエストを待って待機します。この数は予測可能な場合がありますが、近隣の大規模なリクエストにより、各ターンで GPU 時間が消費される可能性があります。
トークンベースの予算設定では、大きなリクエストは「重く」なります。テナントの予算がより早く取り消されるため、テナントはより早く移行し、より小さなリクエストが効率的に処理されるようになります。このモデルは、作業の実際のコストをより忠実に反映しており、単一テナントの大量のトラフィックによって生じるボトルネックに対するより強力な保証です。
クォンタムのサイズは、エンドポイントのバッチ戦略にも合わせられます。ストリーミング エンドポイントは、厳密なインターリーブのためにテナントごとにおよそ 1 つのリクエストの後にローテーションしますが、バッチ エンドポイントでは、完全なバッチに近いバジェットが与えられます。したがって、テナントはスケジューラが先に進む前に意味のあるチャンクを提供することができ、公平性を犠牲にすることなくバッチ処理の効率を維持できます。
公平性が次にどのテナントを使用するかを決定することである場合、優先セレクターはそのテナントのどれを決定しますか

次にテナントのリクエストが続きます。
Deficit Round Robin が順番にテナントを選択すると、スケジューラはそのテナントの回線から 1 つのリクエストを引き出しますが、盲目的ではありません。各行は、次のように順序付けされた体系化されたキューです。
優先度: 明示的に優先度の高いリクエストが最初に処理されます。
期限: 同じ優先順位の中で、期限が最も早いリクエストが優先されるため、時間に敏感な作業が期限切れになることはありません。
到着時間: 最後のタイブレーカーとして、以前のリクエストが最初に実行され、安定した予測可能な先入れ先出し動作が行われます。
この順序を 1 つのグローバル キューではなく、各テナント内で維持することで、プラットフォーム内で公平性と緊急性を共存させることができます。テナントは、プラットフォームを共有しているからといって、優先権や期限の保証を失うことはありません。これらの保証を、自社の公平な容量に適用するだけです。
4 つのステージは、クリーンなリクエストのライフサイクルを構成します。レート制限により、途中での入場が制限されます。各バッチが組み立てられる際の選択は、階層化、公平性、テナントごとの緊急性によって決まります。
これらの各メカニズム (レート制限、階層化、ラウンドロビン スケジューリング、優先キュー) は、MLOps やプラットフォーム エンジニアにはよく知られています。その斬新な価値は、以下をどのように統合するかによって生まれます。
レート制限はシステムを過負荷から保護し、テナントごとのクォータを強制します。
階層化は商業的な取り組みを尊重します。
赤字ラウンド ロビンでは、同じ層の許可されたトラフィックの中で、すべてのテナントが公平でバースト防止のシェアを取得することが保証されます。
優先順位により、各テナント独自の緊急性と期限の注文が公平な配分内で維持されます。
ステップ 2 ～ 4 はバッチがいっぱいになるまで繰り返され、バッチは GPU に送信されます。その結果、あなたのエクスペリエンスが、隣人の騒音の大きさではなく、自分の階層と能力の公平な配分に依存するプラットフォームが誕生しました。

どうやらその日のようだ。
サービング・フェアネスは、当社の SaaS API および AWS を含むサードパーティのマーケットプレイス デプロイメントを通じて、Cohere モデルを使用するすべての顧客に対して有効になりました。
私たちは、お客様のニーズを最優先してこの機能を開発しました。 Cohere モデルを使用していて、フィードバック、観察されたパフォーマンスの問題、または改善の提案がある場合は、Discord でエンジニアリング チームにお問い合わせください。
社外サービス担当技術スタッフマネージャー
AIは近道ではありません。
そうやってビジネスは前進するのです。
最新情報をお届けします。

## Original Extract

How Cohere is ensuring that every tenant gets their fair share of compute.

LLM Serving Fairness North Mini Code. Cohere's first model for developers.
An enterprise-ready AI platform that powers modern workplace productivity
An intelligent search and discovery system to surface business insights
High-performance models for agentic, multimodal, multilingual AI
A speech recognition model for generating highly accurate audio transcripts
A family of multilingual research models covering 70+ languages
A leading multimodal search and retrieval tool
A powerful model that provides a semantic boost to search quality
Your dedicated, secure model inference platform — managed by Cohere
Cohere's research lab that seeks to solve complex ML problems
Explore enterprise AI case studies and success stories
An enterprise-ready AI platform that powers modern workplace productivity
An intelligent search and discovery system to surface business insights
High-performance models for agentic, multimodal, multilingual AI
A speech recognition model for generating highly accurate audio transcripts
A family of multilingual research models covering 70+ languages
A leading multimodal search and retrieval tool
A powerful model that provides a semantic boost to search quality
Your dedicated, secure model inference platform — managed by Cohere
Cohere's research lab that seeks to solve complex ML problems
Explore enterprise AI case studies and success stories
LLM Serving Fairness: No more noisy neighbours
How Cohere is ensuring every tenant gets their fair share of compute
Manager of Technical Staff, External Serving
AI isn’t a shortcut.
It’s how business gets ahead.
Running large language models as a multi-tenant SaaS platform comes with a deceptively hard problem: many organizations share the same pool of GPUs, and their traffic is bursty and uneven. Left unmanaged, one customer's spike can become every other customer's latency problem.
In this blog post, we’ll walk you through Cohere’s new solution to scheduling inference requests fairly across tenants, using a combination of architectural patterns and a classic scheduling algorithm.
Inference is most efficient when requests are processed in batches. A GPU that’s been fed a full batch runs hot and economically, while one handling a single request at a time mostly sits idle. So, requests queue up briefly and get packed into batches before hitting the hardware.
The catch is ordering . Imagine a naive queue: a single, shared line ordered only by priority and deadline. Now, consider what happens when one organization fires 10,000 requests in a burst while a second organization sends just five. With a single global queue, the 10,000 requests pile in front, and the five well-behaved requests wait at the back.
This is the classic noisy-neighbor problem, and a multi-tenant LLM platform is no different from any other shared system that faces these traffic patterns.
The goal of Serving Fairness is to isolate tenants from one another, so that the share of inference capacity that a tenant receives depends on fair scheduling — not on how aggressively it floods the queue. At the same time, it still preserves priority and deadline ordering within each tenant while maintaining batching efficiency.
Cohere manages workloads fairly across tenants by combining four distinct mechanisms, each solving a different slice of the problem. They run in a fixed order: a Rate Limiter controls admission on the way in , and then three selectors — Performance Tier , Deficit Round Robin , and Priority — select the next request on the way out.
Here is the architecture and step-by-step flow:
Before a request ever joins the scheduling queue, it passes through admission control . Rate limiting caps the maximum number of inference requests that a single tenant can submit within a given timeframe,such as requests per minute or per month.
At Cohere, these limits are configured at the endpoint level and differ between models based on how many resources each model consumes. For example, a heavyweight generative model carries tighter limits than a lightweight embedding model.
There is also a real-time throttling check: if the queue is already so backlogged that a new request could not be served within its latency target, the request is simply rejected early instead. This protects the system from accepting work it cannot honor and keeps latencies predictable under load.
After requests have been admitted, they proceed to the series of selectors below.
The first selection decision is the tier . Compute resources are prioritized based on service-level agreements (SLAs): higher-paid tiers are allocated greater processing priority and faster queue handling than lower tiers (or Trial Tiers). In turn, these latter customers are served as capacity allows.
Crucially, tiering determines who goes first ; it does not, by itself, prevent a single large tenant within a tier from dominating. That is the purpose of the next two layers.
The heart of the system is the Deficit Round Robin (DRR) algorithm, which ensures an equitable distribution of compute across the fleet within a tier .
Each tenant (an "organization") gets its own line. Instead of draining whichever line is longest, the scheduler takes turns between tenants. Each tenant is granted a small budget — a quantum — of work that it may perform on its turn. When a tenant uses its turn, its budget is debited by the cost of the request it just sent through. A tenant that runs out of budget is skipped until its budget is replenished on the following round.
The elegance of DRR is that it is both work-conserving and weighted . Cheap requests let a tenant come up more often, and expensive ones less, so no tenant can run away with the GPU — but no capacity is wasted either. Returning to our earlier example: even if Org A queued 10,000 requests and Org B only five, Org B still gets its due turn every cycle. Org A's burst no longer translates into Org B's latency.
The scheme hinges on two key variables:
Quantum: How much budget a tenant is granted each round
Cost: How much budget each request consumes
The crucial design choice lies in what unit those variables are expressed in. Critically, this determines how we might conceptualize “fairness” in different inference contexts. At Cohere, we use two cost models depending on the endpoint: request-based budgeting and token-based budgeting.
For simplicity, the cost of every request is given as 1, and the quantum is a count of requests that a tenant may send per turn. Fairness is, therefore, measured purely in the number of requests.
This is suboptimal for generative endpoints , such as chat and completions, where request serving costs can vary dramatically. A request with a 100K-token prompt may consume orders of magnitude more resources than one with a 1K-token prompt. For reasoning-capable models, the total cost depends not only on the input size, but also on the amount of intermediate reasoning, planning, and output generation triggered by the specific context of the request.
Ideally, DRR would use a feedback loop based on the EMA (Exponential Moving Average) of a request's token-normalized serving cost, allowing budgets to adapt to observed resource consumption. Static budgeting works best when requests within an endpoint are broadly similar in size and cost .
Here, the cost of a request is its token count, and the quantum is a token budget per round. Fairness is now measured in actual work performed. This is the natural fit for batched endpoints like embeddings and rerankers, where the token sum of a batch — not the number of items — drives GPU cost. A tenant sending a few very large documents spends its budget quickly and yields the floor sooner; a tenant sending many small requests is charged little per request and surfaces more often. This way, no tenant can monopolize a GPU by submitting a handful of extra large requests.
So, under request-based budgeting, your request waits behind (at most) a fixed number of requests from each competing tenant, no matter how large those requests are. This count may be predictable, but a neighbor's large requests can still cost you GPU time on each of their turns.
Under token-based budgeting, a large request is "heavier": it draws down its tenant's budget faster, so that tenant moves on sooner and smaller requests can efficiently flow through. This model is a more faithful reflection of the true cost of the work and is the stronger guarantee against bottlenecks created by a single tenant's heavy traffic.
The quantum is also sized to the endpoint's batching strategy: streaming endpoints rotate after roughly one request per tenant for tight interleaving, while batched endpoints grant a budget close to a full batch. A tenant can, therefore, contribute a meaningful chunk before the scheduler moves on — preserving batching efficiency without sacrificing fairness.
If fairness is about deciding which tenant goes next, then the Priority selector determines which of that tenant's requests goes next.
Once Deficit Round Robin selects a tenant for its turn, the scheduler pulls out a single request from that tenant’s line — but not blindly. Each line is a systematized queue ordered by:
Priority: Explicitly higher-priority requests are served first.
Deadline: Among equal priorities, the request with the earliest deadline wins, so time-sensitive work isn't left to expire.
Arrival time: As a final tiebreaker, earlier requests go first, giving stable, predictable first-in-first-out behavior.
Keeping this ordering inside each tenant — rather than in one global queue — is what lets fairness and urgency coexist in the platform. A tenant never loses its priority and deadline guarantees just because it shares the platform; it simply applies those guarantees to its own fair share of capacity.
The four stages compose into a clean request lifecycle. Rate limiting governs admission on the way in ; tiering, fairness, and per-tenant urgency govern selection on the way out as each batch is assembled.
Separately, each of these mechanisms — rate limiting, tiering, round-robin scheduling, and priority queues — are well familiar to MLOps and Platform Engineers. Their novel value comes from how they integrate :
Rate limiting protects the system from overload and enforces per-tenant quotas.
Tiering honors commercial commitments.
Deficit Round Robin guarantees that, among admitted traffic of equal tier, every tenant gets a fair, burst-proof share.
Priority preserves each tenant's own urgency and deadline ordering within its fair share.
Steps 2–4 repeat until the batch is full, and the batch is then sent to the GPU. The result is a platform where your experience depends on your tier and your fair share of capacity — not on how loud your neighbor happens to be that day.
Serving Fairness is now enabled for all customers using any Cohere model through our SaaS API and third-party marketplace deployments, including AWS.
We’ve developed this feature with customer needs front and center. If you’re using Cohere models and have feedback, observed performance issues, or suggestions for improvement, please reach out to our engineering team on our Discord .
Manager of Technical Staff, External Serving
AI isn’t a shortcut.
It’s how business gets ahead.
We’ll keep you up to date with the latest.
