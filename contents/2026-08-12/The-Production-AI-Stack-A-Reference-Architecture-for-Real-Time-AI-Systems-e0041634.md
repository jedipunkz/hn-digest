---
source: "https://www.moss.dev/blog/the-production-ai-stack"
hn_url: "https://news.ycombinator.com/item?id=49277978"
title: "The Production AI Stack: A Reference Architecture for Real-Time AI Systems"
article_title: "The Production AI Stack: A Reference Architecture for Real-Time AI Systems"
author: "srimalireddi"
captured_at: "2026-08-12T20:36:27Z"
capture_tool: "hn-digest"
hn_id: 49277978
score: 1
comments: 0
posted_at: "2026-08-12T20:16:09Z"
tags:
  - hacker-news
  - translated
---

# The Production AI Stack: A Reference Architecture for Real-Time AI Systems

- HN: [49277978](https://news.ycombinator.com/item?id=49277978)
- Source: [www.moss.dev](https://www.moss.dev/blog/the-production-ai-stack)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T20:16:09Z

## Translation

タイトル: プロダクション AI スタック: リアルタイム AI システムのリファレンス アーキテクチャ
説明: すべてのリアルタイム AI システムの 7 つの層 (モデル、推論、検索、メモリ、セッション、オーケストレーション、デプロイメント) と、各層での遅延の原因。瞬時に感じられる音声エージェント、副操縦士、会話型 AI を構築するためのリファレンス アーキテクチャ。

記事本文:
プロダクション AI スタック: リアルタイム AI システムのロードのためのリファレンス アーキテクチャ
会話型 AI のリアルタイム セマンティック検索。
Moss チームによる ❤️ を使用して構築
ブログに戻る July 18, 2026 · 25 min read プロダクション AI スタック: リアルタイム AI システムのリファレンス アーキテクチャ
創業者は、現代の標準的な AI スタックになっているもの、つまり大規模な言語モデル、会社のヘルプ センターを含むクラウド ベクター データベース、注文の検索や返品の開始などの一般的なアクションのための小規模なツール セットを使用して、e コマース ストア用の AI サポート エージェントを構築しました。システムはうまく機能しました。顧客は人間と話すことなく、質問したり、注文状況を確認したり、返品を開始したりすることができます。
問題は正しさではなかった。それはレイテンシーでした。
すべての対話は同じ実行パスに従いました。顧客が質問すると、エージェントはベクトル データベースに関連するコンテキストを照会し、結果が返されるまで数百ミリ秒待ってから、初めてモデルが応答の生成を開始できます。アーキテクチャには技術的に壊れたものは何もありませんでしたが、累積的な遅延は明らかで、会話が必要以上に遅く感じられました。
実稼働 AI システムを構築する多くのチームと同様に、彼は検索を完全に削除することで応答性を最適化しました。
毎回コンテキストを取得するのではなく、すべてをシステム プロンプトに直接埋め込みました。ヘルプセンター、配送ポリシー、返品ルール、価格情報、よくある質問、および顧客が合理的に尋ねる可能性のあるその他の情報はすべて、すべてのリクエストに伴うプロンプトの一部になりました。これにより、外部依存性が排除され、レイテンシが短縮され、アーキテクチャが簡素化されました。
しばらくの間は、それが正しいトレードオフのように見えました。
しかし、会話が長くなると、建築家は

さらに微妙な方法で失敗するようになりました。すべてのメッセージは、すでに企業の知識ベースを保持していたコンテキスト ウィンドウに会話履歴のレイヤーを追加し、モデルはますます大量の情報を推論することを余儀なくされました。利用可能なコンテキストがいっぱいになると、検索精度は実質的に確率的再現に置き換えられました。配送に関する質問には返品ポリシーが引き込まれ始め、回答には無関係な製品の割引が表示され、会話の初めには明らかだった事実は時間の経過とともに信頼性が低くなっていきました。
最もイライラしたのは、テスト中にこれらの障害がほとんど発生しなかったことです。短いベンチマーク プロンプトは引き続き良好なパフォーマンスを示しましたが、運用環境で実際に重要な会話は、アーキテクチャの弱点を明らかにしたものでした。顧客はフォローアップの質問をしたり、以前のメッセージを参照したり、ワークフローの途中で考えを変えたり、数十回の会話ターンにわたってシステムが一貫性を維持することを期待していました。これらはまさにモデルの信頼性が最も低くなる相互作用でした。
このパターンは、カスタマー サポート エージェントやエンタープライズ副操縦士から音声 AI プラットフォームや会話型検索システムに至るまで、本番 AI アプリケーションのあらゆるカテゴリにわたって見られます。チームは、実際にはアーキテクチャ上の制限に遭遇しているにもかかわらず、プロンプトの問題やモデルの品質の問題に対処していると想定することがよくあります。大きなプロンプトは最終的に処理コストが高くなり、コンテキスト ウィンドウは必然的に飽和状態になり、言語モデルに依存してすべてを記憶すると、会話が長くなるにつれて予測が困難になるシステムが生成されます。
音声 AI プラットフォーム、エンタープライズ副操縦士、会話型検索エクスペリエンスを強化しているかどうかに関係なく、あらゆる本番 AI システムは、

最終的には、同じ 7 つのアーキテクチャ層で構成されます。
モデルは、どの基礎モデルが各タスクを担当するかを決定します。
推論は、トークン生成がどこでどのように行われるかを制御します。
検索により、各リクエストに答えるために必要な知識が取得されます。
メモリには、ユーザーと以前のインタラクションに関する情報が保持されます。
セッションは、複数のリクエストにわたって会話状態を維持します。
オーケストレーションは、会話全体を通じてすべてのコンポーネントの実行を調整します。
デプロイメントによって、各レイヤーがブラウザー内、エッジ、デバイス上、またはクラウド内で実行される場所が決まります。
これらの各レイヤーは個別に最適化できますが、本番 AI システムが単一のコンポーネントによって制限されることはほとんどありません。パフォーマンス、遅延、信頼性、コストは、それらの間の相互作用から生まれます。これらのトレードオフを理解するには、個々のレイヤーではなくアーキテクチャ全体に注目する必要があり、そこから始めます。
すべての本番 AI システムは、レイテンシという同じ制約から始まります。
どのモデルを使用するか、取得をどのように構成するか、どこで推論を実行するかを決定する前に、ユーザーが実際にどの程度の待ち時間を許容するかを理解する必要があります。すべてのインタラクティブ製品には遅延の予算があり、その予算を超えると、いくらモデルの品質を高めてもユーザー エクスペリエンスを回復することはできません。
大規模な言語モデルが存在する数十年前に、ヒューマン コンピュータ インタラクションの研究によってこれらのしきい値が確立されました。約 100 ミリ秒未満の応答は瞬時に感じられます。 1 秒あたり、ユーザーは思考の流れの中に留まりますが、自分が待っていることに気づきます。 10 秒を超えると、注意は別のところに移ります。 IBM の Doherty Threshold に関する研究でも 1980 年代初頭に同様の結論に達し、対話型システムは保守のために約 400 ミリ秒以内に応答する必要があると主張しました。

ユーザーとコンピュータ間の継続的なフィードバック ループです。
会話はさらに寛容ではありません。人間の交代は通常、約 300 ミリ秒以内に行われます。つまり、AI システムがデータの取得、ネットワーク呼び出しの待機、トークンの生成に費やすすべてのミリ秒が、人々が生涯をかけて期待してきたリズムと直接競合することになります。
したがって、許容可能なレイテンシの予算は、構築している製品によって異なります。
これらは任意のターゲットではありません。それらはユーザーの行動に直接影響します。
Google は 2009 年の対照実験で、ランダム化されたユーザー グループの検索結果にサーバー側の遅延を意図的に 100 ～ 400 ミリ秒追加することでこれを実証しました。遅延がその範囲の下限であっても、ユーザーが実行する検索が減り、遅延が増加するにつれてエンゲージメントは着実に減少しました。単独では重要ではないように見える小さな遅延が、毎日何百万回も発生すると、測定可能な製品の問題になります。
音声 AI を使用すると、これらの制約を無視できなくなります。
すべての AI アプリケーションの中で、音声は遅延バジェットが最も小さい一方で、応答を開始するまでに必要な連続操作の数が最も多くなります。一般的な会話では、音声認識、検索、言語モデル推論、音声合成が次々に行われる必要があります。各ステージは前のステージの出力に依存するため、レイテンシは重複するのではなく蓄積されます。
代表的な本番パイプラインは次のようになります。
1 つの数字がすぐに目立ちます。
通常、言語モデルは遅延の最大の原因ではありません。回収です。
解決されたインフラストラクチャとして扱われることが多いにもかかわらず、取得はリクエスト パイプラインの中で最も遅く、最も変化しやすい段階であることがよくあります。理想的な条件下では、

レイテンシ バジェットのかなりの部分を占めており、現実的な運用条件では、モデルが単一のトークンを生成する前にバジェット全体を超えることがよくあります。
リアルタイム AI のために Web インフラストラクチャが壊れる理由
デフォルトの AI スタックは、そのアーキテクチャを Web から継承しました。
10 年以上にわたり、Web アプリケーションは同じパターンに従ってきました。つまり、アプリケーション サーバーをステートレスに保ち、データをマネージド サービスに一元管理し、すべてをネットワーク経由で接続してきました。このアーキテクチャは、従来の Web ワークロードに対して非常に効果的です。これは、ほとんどのユーザー インタラクションには少数のネットワーク リクエストしか関係しておらず、別のデータベース クエリを追加してもユーザー エクスペリエンスが意味のある形で変わることはほとんどないためです。
実稼働 AI システムは、根本的に異なる実行モデルを持っています。
AI アプリケーションは、1 つまたは 2 つのデータベース検索を含むページを提供するのではなく、会話中に検索、ツールの実行、モデル推論、メモリ アクセス、再ランキング、および追加の取得を繰り返し実行します。すべての会話ターンは依存する操作のパイプラインとなり、それぞれが別のネットワーク境界を導入し、遅延が蓄積する別の機会が生じます。
業界はすでにその影響を経験しています。 LangChain のエージェント エンジニアリングの現状調査によると、すでに運用環境にエージェントを導入しているチームの間では、遅延がエンジニアリング上の 2 番目に大きな課題であり、それを上回るのは出力品質だけであることがわかりました。
生のネットワーク遅延以外に目を向けると、その理由が明らかになります。
ネットワーク ホップは、単にマシン間でパケットを送信するのに必要な時間ではありません。すべてのリクエストでは、リモート サービスの実行が開始される前に、シリアル化、TLS ネゴシエーション、認証、接続プーリング、ロード バランシング、スケジューリング、およびキューイングも発生します。単一のクラウド領域内であっても、そのオーバーヘッドは

通常、すべてのリクエストに 40 ～ 150 ミリ秒が追加されます。クロスリージョン通信では、モデルが単一のトークンを処理する前に、この時間が 150 ～ 400 ミリ秒に簡単に増加する可能性があります。
AI システムが 1 つの個別のリクエストを実行することはほとんどないため、これらのコストはさらに重要になります。
Google の論文「The Tail at Scale」では、リクエストが複数のサービスに分散すると、まれに発生する遅い応答がシステム全体の遅延を支配し始めることが示されています。実稼働 AI システムも同様の動作を示しますが、多くのリクエストを並行して行うのではなく、順番に実行することが多い点が異なります。 1 つのサービスの末尾のレイテンシが次のサービスの開始点となり、会話のターン全体にわたって遅延が悪化します。
これが、レイテンシーの中央値が誤解を招く指標となることが多い理由です。
マネージド ベクター データベースは、P50 では完全に許容できるように見えますが、P99 では、コールド接続、ガベージ コレクションの一時停止、ノイズの多い隣人、または一時的なリソース競合により、大幅に高いレイテンシを示します。弊社独自の実稼働ベンチマークでは、クラウド ベクター データベースのラウンドトリップは通常、P99 で 500 ～ 900 ミリ秒の間に集中します。
ユーザーはレイテンシの中央値を経験することはありません。
彼らはあなたの最悪の瞬間を経験します。
20 ターン続く会話では、最悪の場合のリクエストは統計的に外れ値ではなくなります。それらは避けられなくなります。そして、推論が始まる前に取得に 0.5 秒以上かかると、どの言語モデルもシステムの応答性を高めることはできません。
モデルは、ほとんどのチームが議論に最も多くの時間を費やすレイヤーですが、実稼働 AI システムにおいて主要なボトルネックになることはほとんどありません。ファウンデーション モデルの機能は著しく向上し、レイテンシーは改善され続けており、プロバイダー間の切り替えはこれまでより簡単になっています。もはや課題は、最適なモデルを見つけることではありません。リグを使用しているのです

適切なジョブに適した ht モデル。
実稼働 AI システムは、単一のモデルを中心に構築されるべきではありません。これらは、さまざまなワークロードに最適化されたモデルのポートフォリオを中心に構築する必要があります。フロンティア モデルは、より深い推論により追加のレイテンシとコストが正当化されるタスク用に予約されています。より高速なモデルは、ルーティング、分類、ツールの選択、およびほぼすべてのターンで発生するその他の決定を処理します。小規模またはローカルなモデルは、埋め込み、再ランキング、モデレーション、ガードレール チェックを強化します。実稼働システムのほとんどのリクエストは比較的単純であり、これらの決定を 1 秒以上ではなく数百ミリ秒で応答するモデルにルーティングすることが、多くの場合、最も簡単なレイテンシー改善になります。
インタラクティブなアプリケーションの場合、最初のトークンまでの時間 (TTFT) は、ベンチマーク スコアや 1 秒あたりのトークンよりもはるかに意味のある指標です。ユーザーは、応答がどれだけ速く終わるかではなく、どれだけ早く応答が始まるかに注目します。 TTFT はホストされているモデル間で大きく異なり、プロンプトのサイズとともに増加します。つまり、公開されているベンチマークが実稼働ワークロードを表すことはほとんどありません。合成ベンチマークではなく、現実的なプロンプトと会話履歴を使用して自分で測定します。
モデル構成には直接的なものもあります。

[切り捨てられた]

## Original Extract

The seven layers of every real-time AI system - models, inference, search, memory, sessions, orchestration, and deployment - and where latency comes from at each one. A reference architecture for building voice agents, copilots, and conversational AI that feel instant.

The Production AI Stack: A Reference Architecture for Real-Time AI Systems Loading
Real-time semantic search for Conversational AI.
Built with ❤️ by the Moss team
Back to Blog July 18, 2026 · 25 min read The Production AI Stack: A Reference Architecture for Real-Time AI Systems
A founder built an AI support agent for his ecommerce store using what has become the standard modern AI stack: a large language model, a cloud vector database containing the company's help center, and a small set of tools for common actions like looking up orders and initiating returns. The system worked well. Customers could ask questions, check the status of an order, or begin a return without ever needing to speak with a human.
The problem wasn't correctness. It was latency.
Every interaction followed the same execution path. A customer asked a question, the agent queried the vector database for relevant context, waited several hundred milliseconds for the results to return, and only then could the model begin generating a response. Nothing in the architecture was technically broken, but the cumulative delay was obvious enough that conversations felt slower than they should.
Like many teams building production AI systems, he optimized for responsiveness by removing retrieval altogether.
Instead of retrieving context on every turn, he embedded everything directly into the system prompt. The help center, shipping policies, return rules, pricing information, FAQs, and any other information a customer might reasonably ask about all became part of the prompt that accompanied every request. It eliminated an external dependency, reduced latency, and simplified the architecture.
For a while, it looked like the right tradeoff.
As conversations became longer, however, the architecture started failing in more subtle ways. Every message added another layer of conversational history to a context window that was already carrying the company's knowledge base, forcing the model to reason over an increasingly large body of information. As the available context filled up, retrieval accuracy was effectively replaced by probabilistic recall. Shipping questions began pulling in return policies, discounts from unrelated products appeared in responses, and facts that had been clear at the beginning of the conversation became less reliable over time.
The most frustrating part was that these failures rarely appeared during testing. Short benchmark prompts continued to perform well, while the conversations that actually mattered in production were the ones that exposed the weaknesses of the architecture. Customers asked follow up questions, referred back to earlier messages, changed their minds halfway through a workflow, and expected the system to maintain consistency across dozens of conversational turns. Those were precisely the interactions where the model became least reliable.
This pattern appears across every category of production AI application, from customer support agents and enterprise copilots to voice AI platforms and conversational search systems. Teams often assume they're dealing with a prompting problem or a model quality problem, when in reality they're running into architectural limits. Large prompts eventually become expensive to process, context windows inevitably become saturated, and relying on a language model to remember everything produces systems that become less predictable as conversations grow longer.
Every production AI system, whether it's powering a voice AI platform, an enterprise copilot, or a conversational search experience, is ultimately composed of the same seven architectural layers:
Models determine which foundation model is responsible for each task.
Inference controls where and how token generation happens.
Search retrieves the knowledge required to answer each request.
Memory persists information about users and previous interactions.
Sessions maintain conversational state across multiple requests.
Orchestration coordinates the execution of every component throughout the conversation.
Deployment determines where each layer runs, whether in the browser, at the edge, on device, or in the cloud.
Each of these layers can be optimized independently, but production AI systems are rarely limited by any single component. Performance, latency, reliability, and cost emerge from the interactions between them. Understanding those tradeoffs requires looking at the entire architecture rather than any individual layer, which is where we'll begin.
Every production AI system begins with the same constraint: latency.
Before deciding which model to use, how to structure retrieval, or where to run inference, you need to understand how much latency your users will actually tolerate. Every interactive product has a latency budget, and once that budget is exceeded, no amount of model quality can recover the user experience.
Human computer interaction research established these thresholds decades before large language models existed. Responses under roughly 100 milliseconds feel instantaneous. Around one second, users remain in their flow of thought but become aware that they're waiting. Beyond ten seconds, attention shifts elsewhere. IBM's work on the Doherty Threshold arrived at a similar conclusion in the early 1980s, arguing that interactive systems should respond within roughly 400 milliseconds to maintain a continuous feedback loop between the user and the computer.
Conversation is even less forgiving. Human turn taking typically happens within about 300 milliseconds, which means every millisecond an AI system spends retrieving data, waiting on network calls, or generating tokens directly competes with a rhythm that people have spent their entire lives expecting.
The acceptable latency budget therefore depends on the product you're building:
These aren't arbitrary targets. They directly influence user behavior.
Google demonstrated this in a controlled experiment in 2009 by intentionally adding between 100 and 400 milliseconds of server side latency to search results for a randomized group of users. Even delays at the lower end of that range caused people to perform fewer searches, with engagement steadily decreasing as latency increased. Small delays that seem insignificant in isolation become measurable product problems when they occur millions of times every day.
Voice AI makes these constraints impossible to ignore.
Among all AI applications, voice has the smallest latency budget while requiring the largest number of sequential operations before a response can begin. A typical conversational turn requires speech recognition, retrieval, language model inference, and speech synthesis to happen one after another. Because each stage depends on the output of the previous one, their latencies accumulate rather than overlap.
A representative production pipeline looks something like this:
One number stands out immediately.
The language model isn't usually the largest source of latency. Retrieval is.
Despite often being treated as solved infrastructure, retrieval is frequently both the slowest and most variable stage in the request pipeline. Under ideal conditions it consumes a substantial portion of the latency budget, and under realistic production conditions it often exceeds the entire budget before the model has generated a single token.
Why Web Infrastructure Breaks Down for Real Time AI
The default AI stack inherited its architecture from the web.
For more than a decade, web applications have followed the same pattern: keep application servers stateless, centralize data in managed services, and connect everything over the network. That architecture is highly effective for traditional web workloads because most user interactions involve only a handful of network requests, and adding another database query rarely changes the user experience in any meaningful way.
Production AI systems have a fundamentally different execution model.
Instead of serving a page with one or two database lookups, an AI application performs retrieval, tool execution, model inference, memory access, reranking, and additional retrieval repeatedly throughout a conversation. Every conversational turn becomes a pipeline of dependent operations, each introducing another network boundary and another opportunity for latency to accumulate.
The industry is already experiencing the consequences. LangChain's State of Agent Engineering survey found that among teams already deploying agents in production, latency ranks as the second largest engineering challenge, surpassed only by output quality.
The reason becomes obvious once you look beyond raw network latency.
A network hop isn't simply the time required to transmit packets between machines. Every request also incurs serialization, TLS negotiation, authentication, connection pooling, load balancing, scheduling, and queueing before the remote service begins executing. Even within a single cloud region, that overhead commonly adds 40 to 150 milliseconds to every request. Cross region communication can easily increase that to 150 to 400 milliseconds, all before the model has processed a single token.
Those costs become even more significant because AI systems rarely perform one isolated request.
Google's paper The Tail at Scale showed that once requests fan out across multiple services, infrequent slow responses begin to dominate overall system latency. Production AI systems exhibit the same behavior, except instead of making many requests in parallel, they often perform them sequentially. The tail latency of one service becomes the starting point for the next, causing delays to compound across an entire conversational turn.
This is why median latency is often a misleading metric.
A managed vector database might appear perfectly acceptable at P50 while exhibiting dramatically higher latency at P99 because of cold connections, garbage collection pauses, noisy neighbors, or temporary resource contention. In our own production benchmarks, cloud vector database round trips commonly cluster between 500 and 900 milliseconds at P99.
Users never experience your median latency.
They experience your worst moments.
In a conversation lasting twenty turns, those worst case requests stop being statistical outliers. They become inevitable. And once retrieval consumes half a second or more before inference even begins, no language model can make the system feel responsive.
Models are the layer most teams spend the most time debating, yet they're rarely the primary bottleneck in a production AI system. Foundation models have become remarkably capable, latency continues to improve, and switching between providers is easier than ever. The challenge is no longer finding the best model. It's using the right model for the right job.
Production AI systems shouldn't be built around a single model. They should be built around a portfolio of models optimized for different workloads. Frontier models are reserved for tasks where deeper reasoning justifies the additional latency and cost. Faster models handle routing, classification, tool selection, and other decisions that occur on nearly every turn. Smaller or local models power embeddings, reranking, moderation, and guardrail checks. Most requests in a production system are relatively simple, and routing those decisions to a model that responds in a few hundred milliseconds instead of more than a second is often the easiest latency improvement you'll make.
For interactive applications, time to first token (TTFT) is a far more meaningful metric than benchmark scores or tokens per second. Users notice how quickly a response begins, not how quickly it finishes. TTFT varies significantly across hosted models and grows with prompt size, which means published benchmarks are rarely representative of production workloads. Measure it yourself using realistic prompts and conversation histories rather than synthetic benchmarks.
Model configuration also has a direct

[truncated]
