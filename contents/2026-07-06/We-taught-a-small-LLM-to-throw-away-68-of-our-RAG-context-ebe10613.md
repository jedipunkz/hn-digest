---
source: "https://www.kapa.ai/blog/how-we-prune-rag-context"
hn_url: "https://news.ycombinator.com/item?id=48809354"
title: "We taught a small LLM to throw away 68% of our RAG context"
article_title: "How we taught a small LLM to throw away 68% of our RAG context - kapa.ai - Instant AI answers to technical questions"
author: "emil_sorensen"
captured_at: "2026-07-06T19:42:20Z"
capture_tool: "hn-digest"
hn_id: 48809354
score: 1
comments: 0
posted_at: "2026-07-06T19:28:24Z"
tags:
  - hacker-news
  - translated
---

# We taught a small LLM to throw away 68% of our RAG context

- HN: [48809354](https://news.ycombinator.com/item?id=48809354)
- Source: [www.kapa.ai](https://www.kapa.ai/blog/how-we-prune-rag-context)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T19:28:24Z

## Translation

タイトル: 小規模な LLM に RAG コンテキストの 68% を廃棄するよう教えました
記事のタイトル: 小規模 LLM に RAG コンテキストの 68% を廃棄するように教えた方法 - kapa.ai - 技術的な質問に対する AI の即時回答
説明: 96% の再現率を維持しながら、エージェントのコンテキストを回答に実際に必要なものまで枝刈りします。

記事本文:
2026 年 7 月 2 日 小規模 LLM に RAG コンテキストの 68% を廃棄するように教えた方法
96% の再現率を維持しながら、エージェントのコンテキストを回答に実際に必要なものまで枝刈りする
Kapa は、大規模な製品知識ベースに対する複雑な質問に答える AI アシスタントを構築しています。技術文書、API リファレンス、PDF、フォーラム、サポート スレッドを思い浮かべてください。開発者は当社の検索 API を使用してエージェントに製品に関するコンテキストを提供し、同じ取得レイヤーがエンドツーエンドのアシスタントを強化します。
エージェントが依然として RAG を必要とするかどうかについては、2026 年になってもさまざまな議論が行われていますが、ナレッジ ベースが大規模かつ複雑になると、私たちの分野ではそれに近づくものは何もありません。私たちの検索には、エージェントによるものやシングルパスによるものなど、いくつかの形式がありますが、質問に関連するドキュメントの塊を見つける検索機能と、そこから回答を書き込むジェネレーターである LLM という同じ形式を共有しています。
この投稿の短いバージョン: 2 つのステップの間に 3 番目のステップを追加しました。小さくて安価な LLM は、質問と取得されたすべてのチャンクを一緒に読み取り、高価なモデルがそれらを認識する前に、回答に必要のないチャンクを破棄します。コンテキストの約 68% を削除し、再現率の約 96% を維持し、クエリのコストをそれ自体のコストを差し引いた 3 分の 1 に削減します。この投稿では、そこに至るまでの経緯を説明します。
無視されたチャンクには依然として費用がかかります
レトリーバーは漏斗です。埋め込みとキーワード検索により、数十万のチャンクからなるナレッジ ベースが数百の候補にまで削減され、リランカーがそれらを注文し、上位 15 程度がチェーン内で最大かつ最も高価なモデルであるジェネレーターに到達します。それでも、ジェネレーターが読み取る内容のほとんどは質問には必要ありません。これは意図的なものです。レトリーバーは最大限の再現を目指しており、ジェネレーターがノイズを無視することを信頼しています。
ただし、ジェネレータは無視したチャンクごとに料金が請求されます

解像度。当社のアシスタントでは、取得されたチャンクはクエリのコストの約 3 分の 2 であり、回答、会話履歴、システム プロンプトを合わせたコストよりも高くなります。チャンクが少なくなるごとに、クエリ コストが約 4% 削減されます。また、エージェントでは、すべてのツール呼び出しがその出力を同じコンテキストに注ぎ込むため、コンテキストは急速に増大します。検索結果が厳密になると、エージェントが保持しなければならない他のすべてのもののためのスペースが確保され、腐るコンテキストが少なくなります。
問題はリコールです。必要な答えの一部をドロップすると、間違った答えと数セントを交換することになります。プルーナーはまさにそのトレードオフと同じくらい優れています。つまり、失われたリコールポイントごとに圧縮が得られます。
上位 K を返す前にすでに再ランク付けを行っているため、再ランク付けのスコアだけを公開し、呼び出し側にカットしてもらうように求められることがあります。すべてを 0.7 以上に保ち、残りを削除します。これは 2 つの理由で失敗し、2 番目の理由が私たちが構築したすべてのものを形作りました。
まず、リランク スコアは順序付けであり、測定ではありません。このクエリではチャンク A がチャンク B に勝っていると表示されますが、それ以上のことはありません。スコアはクエリ間で調整されない、と Cohere 氏も同様に述べているため、固定のカットオフは機能しません。ランキングがサポートする唯一のカットオフは位置的な上位 N であり、ノイズであろうと答えであろうと最後のチャンクが削除されます。
第 2 に、これは完全なキャリブレーションでも存続します。関連性は単一のチャンクの特性ではありません。ほとんどのパイプラインのリランカーはポイント単位のクロスエンコーダーです。これらは、各クエリとチャンクのペアを単独でスコア付けし、取得された他のチャンクと一緒にスコアリングすることはありません。匿名化された制作例は次のとおりです。
2 番目のチャンクは監査ログについてまったく言及していないため、ノイズとしてスコア付けされますが、まだ答えの半分であり、チャンクは最初のチャンクの次に関連するだけであるため、ポイントごとのスコアではそれがわかりません。また、チャンクは複数の部分からなる質問を分割し、それぞれ単独では役に立ちません。本当の問題は決して、c かどうかではありません。

hunk はそれ自体に関連性がありますが、それが一緒になって質問に答えるセットに属しているかどうかです。
同じように失敗する賢い修正
リランカーを諦める前に、アンカー ドキュメント (Sinhababu et al.) を試しました。関連性が既知の合成チャンクをランキングに埋め込むことでリランカーのスケールを絶対的なものにし、必須から無関係までレベルごとに 1 つずつ書き込んでから、保持したい最低レベルのアンカーより下にランク付けされているすべての実際のチャンクを削除します。すでに実行しているリランクに加えて追加の LLM 呼び出しが 1 つあり、実にエレガントです。
同じ根本的な理由で、それは機能しませんでした。アンカーはキャリブレーションを修正しますが、スコアを修正することはできません。リランカーは、明らかに無関係なチャンクの下に部分的および間接的に関連するチャンクを配置し続けました。それらを維持するには、アンカーはほとんど剪定されないように非常に低く設置する必要があります。
この失敗は有益な結果でした。判断されるのはセットであるため、プルーンが何であれ、質問とすべてのチャンクを一度に見る必要があります。
そこで、LLM にチャンクをグレーディングさせます
私たちが出荷したのは、リランカーとジェネレーターの間の 1 つのリスト単位の LLM 呼び出しです。質問とすべてのチャンクを取得し、プロンプトに書き込まれた 5 段階のスケールに照らしてすべてのチャンクを採点します。
このチャンクがなければ、それが直接の答えであるか、別のチャンクが依存する定義や前提条件であるかにかかわらず、答えを生成することはできません。
単独では応答しませんが、他のチャンクと組み合わせて完全な応答に必要なものを提供します。
テーマ通りであり、おそらく役に立ちますが、答えはおそらくそれなしで完成します。
同じドメインまたは共通の用語であり、具体的な貢献はありません。
しきい値以上のチャンクは存続します。この設計は、以前の両方の失敗に対処します。各レベルは単語で定義されているため、4 はすべてのクエリで同じことを意味するため、固定カットオフが最終的に機能します。そして、なぜなら、

モデルは質問とすべてのチャンクを一緒に見てセットを判断できるため、部分的および間接的な関連性が最終的に着地点を見つけます。
モデル: 剪定ばさみの費用は節約した金額から支払われるため、主力モデルは構造上除外されます。小規模で高速な層はすべて同様に判断されたため、少ない推論労力で最も速くて安価なものを選択しました。
しきい値: 圧縮とリコールの間のメインダイヤル。
keep-top-k : 再ランク付けされた上位のいくつかのチャンクはグレードに関係なく合格し、最も強いチャンクをグレーディングのミスから保護します。
また、正直さを保つために、よりシンプルなデザインを 2 つ実行しました。予算の選択: 上位のいくつかを保持し、LLM がさらに最大 N 個を追加できるようにします。サイズは予測可能ですが、予算が使い果たされると、関連性の程度に関係なく、それ以降のチャンクはすべて削除されます。そして最も単純なプルーナーは、どのチャンクを保持するかを LLM に問い合わせるだけで、スケールはありません。直接質問することに勝てないスキームであれば、構築する価値はありません。
私たちは、回答に必要なチャンクが正確にわかっている実際の質問のラベル付きセットで再現率を測定し、その後、各クエリが実際にジェネレーターに送信された正確なチャンクで、実稼働会話のランダムな 1 か月間にわたってすべての構成を再生することで、圧縮、コスト、レイテンシを検証しました。
すべての点は 1 つの構成であり、重要な 2 つの事柄によってプロットされます。 X 軸の圧縮は、プルーナーが破棄する取得されたチャンクの割合です。 Y 軸の再現率は、枝刈り後に回答が必要なすべてのチャンクがまだ残っている質問の割合です。100% では、必要なチャンクを失った質問はありませんでしたが、90% では、10 分の 1 が失われました。上と右の方が良いです。線は各戦略の最適な構成を接続しており、灰色の破線はプルーナーが克服しなければならないベースラインです。単純な上位 N の切り捨てであり、リランカーから返されるチャンクが少ないだけです。
すべてがそれに勝ります

、かなり多くの。リコールを 98% に保つ: 切り捨てにより 1 つのチャンク、約 7% の圧縮が失われる可能性があります。すべての LLM 戦略は 30% 以上に達し、関連性スコアはチャンクの半分近くに低下します。スコアラインもすべての圧縮レベルで他の 2 つを圧倒しているため、残された唯一の決定はその上のどこに配置するかでした。
私たちは、攻撃的な終わりに近いポイントを選択しました。約 96% のリコールが維持され、約 68% のチャンクがドロップされました。 25 問に 1 問で、必要なチャンクが失われます。その代わり、コンテキストの 3 分の 2 が失われ、クエリごとの請求額はプルーナー自身のコストを差し引いて約 34% 減少します。
プルーナーは、クリティカル パス内の取得と生成の間で実行されるため、そのモデル呼び出しがすべてのクエリに追加され、その速度によってコストが決まります。運用セット全体で、選択した構成はクエリごとに約 0.7 秒で実行されました。重い設定は速く上昇するため、少ない推論労力で小さなモデルを使用すると、追加を 1 秒以内に抑えることができます。
その代わり、生成の速度はほとんど向上しません。チャンクが少ないということは、ジェネレーターへの入力トークンが少ないことを意味するため、ジェネレーターは少し早く応答を開始しますが、ほんの数分の 1 秒であり、プルーナー自体の呼び出しをキャンセルするほどではありません。
したがって、プルーニングは、出荷時の構成では 1 秒をはるかに下回る、少量の固定量の遅延を犠牲にして圧縮を実現します。遅延の影響を受けやすいシングルショット パスでは、実際のコストを考慮する必要があります。エージェント内では、すでにターンごとに複数のモデル コールが行われているため、あと 1 回のリーン コールはわずかです。
私たちは、検索が多数のツールの中の 1 つであるとして、最初にこれを展開しました。顧客は、私たちの検索の上にエージェントを構築しました。エージェントは数十のツールを持ち、すべての呼び出しで出力を同じコンテキストに流し込み、ドキュメント検索で返される結果が 3 分の 2 に減れば、他のすべてのためのスペースが確保されます。そこでは、失われたリコールもそれほど危険ではありません。

何かが足りないことに気付いた場合は、再度検索できます。
プルーニングは、Product Agent SDK のナレッジ ベース検索ではデフォルトでオンになっており、取得 API および MCP サーバーではオプションです。
技術文書を顧客対応の AI アシスタントに変える
技術文書用の AI アシスタント
© 2026 Kapa.ai Inc. 無断複写・転載を禁じます。
技術文書用の AI アシスタント
© 2026 Kapa.ai Inc. 無断複写・転載を禁じます。
技術文書用の AI アシスタント
© 2026 Kapa.ai Inc. 無断複写・転載を禁じます。

## Original Extract

Pruning agent context down to what the answer actually needs, while keeping 96% of recall

Jul 2, 2026 How we taught a small LLM to throw away 68% of our RAG context
Pruning agent context down to what the answer actually needs, while keeping 96% of recall
Kapa builds AI assistants that answer complex questions over large product knowledge bases. Think technical documentation, API references, PDFs, forums, support threads. Developers use our retrieval API to give their agents context about their product, and the same retrieval layer powers our end-to-end assistants.
For all the debate in 2026 about whether agents still need RAG, in our domain nothing comes close when knowledge bases get large and complex. Our retrieval comes in several forms, some agentic, some single-pass, but they all share the same shape: a retriever, which finds the chunks of documentation relevant to a question, and a generator, the LLM that writes the answer from them.
The short version of this post: we added a third step between the two. A small, cheap LLM reads the question and all the retrieved chunks together, and throws out the chunks the answer will not need before the expensive model ever sees them. It drops about 68% of the context, keeps about 96% of recall, and cuts the cost of a query by a third, net of its own cost. This post explains how we got there.
Ignored chunks still cost money
A retriever is a funnel. Embedding and keyword search cut a knowledge base of hundreds of thousands of chunks down to a few hundred candidates, a reranker orders them, and the top 15 or so reach the generator, the largest and most expensive model in the chain. Even then, most of what the generator reads is not needed for the question. That is deliberate: retrievers aim for maximum recall and trust the generator to ignore the noise.
But the generator is billed for every chunk it ignores. In our assistants, retrieved chunks are about two-thirds of the cost of a query, more than the answer, the conversation history, and the system prompt combined. Every chunk fewer cuts the query cost by about 4%. And in an agent, every tool call pours its output into the same context, so the context grows quickly; a tighter retrieval result buys room for everything else the agent has to hold and leaves less context to rot.
The catch is recall. Drop a chunk the answer needed and you traded a few cents for a wrong answer. A pruner is exactly as good as that tradeoff: compression gained per point of recall lost.
We already rerank before returning the top K, so we are sometimes asked to just expose the rerank scores and let callers cut on them: keep everything above 0.7, drop the rest. It fails for two reasons, and the second shaped everything we built.
First, a rerank score is an ordering, not a measurement. It says chunk A beats chunk B on this query, nothing more. The scores are not calibrated across queries, Cohere too says as much, so no fixed cutoff works. The only cutoff a ranking supports is positional, top-N, and that drops the last chunk whether it is noise or the answer.
Second, and this survives even perfect calibration: relevance is not a property of a single chunk. The rerankers in most pipelines are pointwise cross-encoders. They score each query-chunk pair alone, never alongside the other chunks it was retrieved with. Here is an anonymized production example:
The second chunk never mentions audit logs, so it scores as noise, yet it is half the answer, and no pointwise score can see that, because the chunk is only relevant next to the first one. Chunks also split multi-part questions between them, each useless alone. The real question is never whether a chunk is relevant by itself, but whether it belongs to a set that together answers the question.
A clever fix that fails the same way
Before giving up on the reranker we tried anchor documents ( Sinhababu et al. ): make the reranker's scale absolute by planting synthetic chunks of known relevance into the ranking, one written per level from Essential to Unrelated, then drop every real chunk that ranks below the anchor of the lowest level you want to keep. One extra LLM call on top of a rerank you already run, and genuinely elegant.
It did not work, for the same underlying reason. Anchors fix calibration, but they cannot fix the scores, and the reranker kept placing partially and indirectly relevant chunks below plainly irrelevant ones. To keep them, the anchor has to sit so low that hardly anything gets pruned.
That failure was the useful result: whatever prunes has to see the question and all the chunks at once, because the thing being judged is the set.
So we let an LLM grade the chunks
What we shipped is one listwise LLM call between the reranker and the generator. It gets the question and all the chunks, and grades every chunk against a five-level scale written into its prompt:
The answer cannot be produced without this chunk, whether it answers directly or is a definition or prerequisite another chunk depends on.
Does not answer on its own, but supplies something a complete answer needs in combination with other chunks.
On topic and plausibly useful, but the answer is likely complete without it.
Same domain or shared terminology, no concrete contribution.
Chunks at or above a threshold survive. The design answers both failures from earlier. Because each level is defined in words, a 4 means the same thing on every query, so a fixed cutoff finally works. And because the model sees the question and all the chunks together, it can judge the set, so partial and indirect relevance finally have somewhere to land.
The model: the pruner is paid for out of what it saves, so flagship models are ruled out by construction; the small fast tiers all judged similarly, so we picked the fastest and cheapest at low reasoning effort.
The threshold : the main dial between compression and recall.
keep-top-k : the top few reranked chunks pass regardless of grade, protecting the strongest chunks from a grading mistake.
We also ran two simpler designs to keep ourselves honest. Budget-select: keep the top few, let the LLM add at most N more; predictable size, but once the budget is spent every further chunk is dropped no matter how relevant. And the simplest possible pruner: just ask the LLM which chunks to keep, no scale. If a scheme cannot beat asking directly, it is not worth building.
We measured recall on a labeled set of real questions where we know exactly which chunks the answer needs, then verified compression, cost, and latency by replaying every configuration over a random month of production conversations, on the exact chunks each query actually sent to the generator.
Every point is one configuration, plotted by the two things that matter. Compression, on the x-axis, is the share of retrieved chunks the pruner throws away. Recall preserved, on the y-axis, is the share of questions that still have every chunk their answer needs after pruning: at 100% no question lost a chunk it needed, at 90% one in ten did. Up and to the right is better. The lines connect each strategy's best configurations, and the dashed grey line is the baseline any pruner has to beat: naive top-N truncation, just returning fewer chunks from the reranker.
Everything beats it, by a lot. Hold recall at 98%: truncation can drop one chunk, about 7% compression. Every LLM strategy reaches 30% or more, and relevance scoring drops close to half the chunks. The scoring line also dominates the other two at every compression level, so the only decision left was where on it to sit.
We picked a point near the aggressive end: about 96% recall preserved, about 68% of chunks dropped. One question in twenty-five loses a chunk it needed; in exchange, two-thirds of the context is gone, and the per-query bill falls by about 34%, net of the pruner's own cost.
The pruner runs between retrieval and generation, in the critical path, so its model call is added to every query, and its speed decides what that costs. Across the production set, the configuration we picked ran in about 0.7 seconds per query. Heavier settings climb fast, so a small model at low reasoning effort is what keeps the addition under a second.
The generation barely speeds up in return: fewer chunks mean fewer input tokens for the generator, so it starts responding a little sooner, but only a fraction of a second, nowhere near enough to cancel out the pruner's own call.
So pruning buys its compression at the cost of a small, fixed amount of latency, well under a second on the configuration we ship. On a latency-sensitive single-shot path that is a real cost to weigh. Inside an agent, which already makes several model calls per turn, one more lean call is marginal.
We rolled it out first where retrieval is one tool among many: customers building agents on top of our retrieval. An agent carries dozens of tools, every call pours output into the same context, and a documentation search that returns two-thirds less buys room for everything else. The lost recall is also less dangerous there: an agent that notices something missing can search again.
Pruning is on by default in our Product Agent SDK 's knowledge base search, and optional in the retrieval API and MCP servers .
Turn technical documentation into customer-facing AI assistants
AI Assistants for Technical Documentation
© 2026 Kapa.ai Inc. All rights reserved.
AI Assistants for Technical Documentation
© 2026 Kapa.ai Inc. All rights reserved.
AI Assistants for Technical Documentation
© 2026 Kapa.ai Inc. All rights reserved.
