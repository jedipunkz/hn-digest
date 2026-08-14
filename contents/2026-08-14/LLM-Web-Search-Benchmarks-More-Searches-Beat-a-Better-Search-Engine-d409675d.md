---
source: "https://openrouter.ai/blog/announcements/web-search-benchmark/"
hn_url: "https://news.ycombinator.com/item?id=49302033"
title: "LLM Web Search Benchmarks: More Searches Beat a Better Search Engine"
article_title: "Live Web Search Benchmarks: Pick the Right Engine, Depth, and Model for Your Agent — OpenRouter Blog"
author: "DGAP"
captured_at: "2026-08-14T17:44:49Z"
capture_tool: "hn-digest"
hn_id: 49302033
score: 1
comments: 0
posted_at: "2026-08-14T17:39:15Z"
tags:
  - hacker-news
  - translated
---

# LLM Web Search Benchmarks: More Searches Beat a Better Search Engine

- HN: [49302033](https://news.ycombinator.com/item?id=49302033)
- Source: [openrouter.ai](https://openrouter.ai/blog/announcements/web-search-benchmark/)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T17:39:15Z

## Translation

タイトル: LLM Web 検索ベンチマーク: より多くの検索がより優れた検索エンジンに勝つ
記事のタイトル: Live Web Search ベンチマーク: エージェントに適切なエンジン、深さ、モデルを選択する — OpenRouter ブログ
説明: 4 つのタスク スイートにわたる Web 検索構成を評価するライブ リーダーボードを公開しました。エンジン、検索深度、モデルを品質、コスト、速度で比較します。

記事本文:
Live Web Search ベンチマーク: エージェントに適切なエンジン、深さ、モデルを選択する — OpenRouter ブログ モデル Fusion Chat ランキング アプリ Enterprise 価格設定ドキュメント モデル Fusion Chat ランキング App Enterprise Pricing ドキュメント ライト モード ブログ
Live Web Search ベンチマーク: エージェントに適切なエンジン、深さ、モデルを選択する
すべての組み合わせをベンチマークして長所と短所を見つけます
検索予算は他のどの要素よりも重要です
最悪のコストシナリオは故障率によって決まる
エンジンも重要ですが、モデルはそれ以上に重要です
Live Web Search ベンチマーク: エージェントに適切なエンジン、深さ、モデルを選択する
すべての組み合わせをベンチマークして長所と短所を見つけます すべての組み合わせをベンチマークして長所と短所を見つけます 検索予算は他の要素よりも重要です 最悪のコスト シナリオは故障率によって左右されます エンジンも重要ですが、モデルはそれ以上に重要です 自分のワークロードで試してみましょう ベンチマーク手法に関するよくある質問 Web 検索は、知識の限界を克服するために、ほとんどの LLM リクエストにとって重要です。ラボと検索プロバイダーは検索をより効果的かつ効率的にするために急速に進化しており、私たち全員が一連の難しい決断を迫られています。一部のラボが内蔵しているネイティブ検索を採用するか、それとも Exa、Parallel、Perplexity などのサードパーティ エンジンを接続するか? 1 回の検索で十分ですか? 十分でない場合、エージェントにどれくらい検索を続けさせればよいですか?検索ターン数が増えるほど、品質を購入する価値があるのでしょうか?
データを使用して最適な検索構成を決定できるように、ライブ リーダーボードを構築しました。新しいベンチマーク ページのデータを参照してください。
すべての組み合わせをベンチマークして長所と短所を見つけます
検索リクエストを設定するときは、次の 4 つの決定事項があります。
モデル。検索エンジンに送信される正確なクエリを作成し、そのクエリを処理します

むかつく。
エンジン。特定のエンジンを選択することも、一部のラボが提供するバンドル エンジンを利用することもできます。 OpenRouter では、OpenAI、Anthropic、Google などのラボのネイティブ エンジンに加えて、Exa、Parallel、Perplexity も提供しています。
検索方法。モデルを呼び出す前に検索を実行し、結果をコンテキストとして渡すことも、モデルの裁量で呼び出す Web 検索ツールをモデルに装備することもできます。
検索予算。検索ツールの方法を選択した場合は、モデルに許可される検索回数の予算を与えることもできます。これにより、モデルは結果が気に入らない場合にクエリを調整したり、フォローアップ検索を実行したりすることができます。私たちのランでは 1、5、または 25 ターンを使用します。
Web 検索のパフォーマンスを包括的に理解するために、複数のモデル、エンジン、検索構成にわたって 4 つのベンチマークを定期的に実行しています。
BrowseComp : 実際の閲覧を必要とする厳しい事実調査
DeepSearchQA : マルチホップのリサーチ質問
WideSearch : 幅広い「表全体を埋める」コレクション
HLE : 専門家試験の質問と検索
各ページでは、品質、価値、速度に基づいて構成がランク付けされているため、ワークロードにとって最も重要な要素に基づいて決定を下すことができます。リーダーボードはライブなので、新しいランが登場し、新しいモデルやエンジンが追加されると、数字が変化します。今日のリーダーが明日もリーダーであるとは限りません。時間の経過とともに変化することが予想されるため、この投稿では今日のリーダーについては多くの時間を費やしません。代わりに、ワークロードに応じて意思決定を行う方法についてデータが何を教えてくれるのかを見てみましょう。
検索予算は他のどの要素よりも重要です
エンジンの予算を 1 ターンから増やすと、他の 1 つの変更よりも品質が向上します。例として、3 つの異なる予算での Perplexity での BrowseComp の最初の実行を示します。
このパターン

私たちが測定したすべてのプロバイダーで維持されます。
これらの実行は、BrowseComp のみを対象とし、サーバー ツールを使用して検索ごとに 10 件の結果を表示し、ページのフェッチやコードの実行は行わず、構成ごとに最新の適格な実行を対象とします。
検索の深さを増やすことは、品質を向上させるために私たちが発見した最も安価な方法です。 1 ターンから 25 に増加すると、スコアは約 2 倍になりますが、質問あたりのコストはわずか 2.5 ～ 7 倍になります。
これにより一般的に応答時間が遅くなると思われるかもしれませんが、常にそうとは限りません。たとえば、Luna は 1 ターンで 1 問あたり 140 秒かかり、25 ターンでは 111 秒かかりました。1 ターンと 5 ターンの両方で実行した 35 の構成のうち、3 分の 1 以上はターンが少なくなるほど遅くなりました。すべて OpenAI モデルでした。これらのモデルは、追加の推論を使用して、制約された検索予算に対処します。
一方で、検索の深さは、より簡単なタスクのコストに悪影響を与える可能性があります。たとえば、HLE では、Perplexity を備えた GPT-5.6 Sol は、コストが 3 倍であるにもかかわらず、1 ターンから 25 ターンの間で同様のスコアを記録しました。検索が単純になる傾向がある場合でも、予算を制限しておいたほうがよいでしょう。
最悪のコストシナリオは故障率によって決まる
予算の拡大が有害となるもう 1 つの状況は、モデルが答えを見つけられない場合です。モデルは、最終的には失敗するとしても、答えを見つけようとして予算を使い果たすことがわかりました。
私たちが記録した最も深い試みである WideSearch テーブルでの 81 回の検索は、依然として不正確と評価されました。ワークロードの失敗率が高い場合は、検索の深さを減らすことがコストを削減するための効率的な方法である可能性があります。
エンジンも重要ですが、モデルはそれ以上に重要です
予算が設定されたら、次に重要な問題はどのモデルを使用するかです。
上の表は、検索エンジン全体でフロンティア モデルとバジェット モデルを比較した、25 ターンでの BrowseComp の結果を示しています。

ネス。
モデルを一定に保ちながらエンジンを変更すると、スコアが平均 10 ポイント変化しましたが、フロンティア モデルとコスト効率の高いモデルの間の平均差は 15 ポイントとさらに大きくなりました。エンジン全体でコストが最も大きく変動したのはフロンティア モデルで、最も高価なエンジンのコストが最も安価なエンジンの 2.5 倍であるのに対し、コスト効率の高いモデルでは 1.5 倍でした。
このような比較が可能な理由は、サーバー ツールがプロバイダーの上に位置しているためです。リクエスト内のモデルを変更しても、プロバイダーが独自の検索を提供しないモデルの場合も含め、検索動作は一貫したままになります。
もちろん、ベンチマークは可能なパフォーマンスの参考値にすぎません。どの構成が試してみる価値があるのか​​、またそのおおよそのコストがわかります。実際のタスクに対するこれらの選択肢のコストと質は異なるため、これらのページでできる最も価値のあることは、これらのページを候補リストとして扱い、上位のいくつかについて独自の質問を実行することです。
上記はすべて、OpenRouter で今すぐ設定できるリクエスト パラメーターです。
ウェブプラグイン。 Web プラグインは、モデルが書き込みを開始する前に 1 回の検索を実行します。これは、新鮮な事実だけが必要な質問に対しては、高速かつ安価なオプションです。
サーバーツール。サーバー ツールはモデルに検索ツールを渡し、次に何を検索するかをモデルに決定させます。これは、答えを見つけるまでにいくつかの手順が必要な場合に必要なものです。
エンジン。 OpenRouter では、エンジンを exa 、Parallel 、perplexity 、またはネイティブに設定します。 auto は、サードパーティにフォールバックする前に、まずネイティブを試行します。
検索予算。最上位の max_tool_calls リクエスト フィールドは、エージェントが取得するターン数、つまり応答するまでに何ラウンドの検索がかかるかを制限し、max_results は毎回返される結果の数を設定します。
妥当な開始点: タスクに最も近いスイートを選択し、最も安価な構成を選択します。

最高スコアから数ポイント以内にあることを確認し、その上の 2 行または 3 行に対して独自の評価セットを再実行して、余分な支出が結果に現れるかどうかを確認します。
すべての実行は、オープン ソース ベンチマーク ハーネスを使用して、運用エンドポイントに対してパブリック OpenRouter API を経由します。
検索パフォーマンスを分離します。検索構成のみを確実に比較するために、検索ごとに 10 件の結果を標準化し、ページの取得やコードの実行は行わないようにしました。表に反映されているように、推論はモデルごとに固定されています。
採点は厳しいです。評価された各回答は、セマンティック比較が必要な場合に LLM ジャッジを使用して、公式の回答キーに対して正しいか間違っています。 WideSearch は、回答項目の精度も個別にレポートします。
コストとスピードは質問ごとに異なります。コストは、採点を含む総支出を評価済みの質問で割ったものです。速度は、評価された質問ごとの候補生成時間です。
各ページには、各構成の最新の予選実行が表示されます。実行は最小限の数の質問を完了すると資格を取得し、新しい実行が古い実行に置き換わります。
これらのスコアは、公開されているベンダー エージェント リーダーボードとどのように比較されますか?
これらは直接比較することはできません。これらのベンチマーク用に公開されている表のほとんどは、検索、全ページ取得、およびコード ツールを組み合わせた完全なエージェント製品を測定しています。これらのリーダーボードは検索構成を分離します。モデルは、ページのフェッチとコード ツールをオフにして、検索結果の抜粋のみを読み取ります。これにより、構成間の直接比較が可能になりますが、ベンチマーク スコアは最大化されません。
どの検索エンジンを選択すればよいですか?
それはモデルとタスクに依存するため、ページが存在します。エンジン間のギャップは、一部のモデルでは大きく、他のモデルではごくわずかであり、プロバイダー独自のネイティブ検索が自動的に最適なオプションになるわけではありません。ライブリードを確認する

ワークロードに最も近いスイートの erboard を参照し、スコアとともにコストとレイテンシを読み取り、新しい実行が開始されると順序が変わるため、時間をかけて再確認します。
リーダーボードには、実稼働エンドポイントに対して OpenRouter のベンチマーク ハーネスで実行された、各構成の最新の予選実行が常に表示されます。ページ上では新しい実行が古い実行に置き換わります。
次にベンチマークを行うべきエンジンやモデルを Discord の #フィードバック で教えてください。

## Original Extract

We published live leaderboards grading web search configurations across four task suites. Compare engines, search depth, and models by quality, cost, and speed

Live Web Search Benchmarks: Pick the Right Engine, Depth, and Model for Your Agent — OpenRouter Blog Models Fusion Chat Rankings Apps Enterprise Pricing Docs Models Fusion Chat Rankings Apps Enterprise Pricing Docs Light mode Blog
Live Web Search Benchmarks: Pick the Right Engine, Depth, and Model for Your Agent
We benchmark all the combinations to find strengths and weaknesses
Search budget matters more than any other factor
Your worst-case cost scenario is driven by your failure rate
While the engine matters, the model matters more
Live Web Search Benchmarks: Pick the Right Engine, Depth, and Model for Your Agent
We benchmark all the combinations to find strengths and weaknesses We benchmark all the combinations to find strengths and weaknesses Search budget matters more than any other factor Your worst-case cost scenario is driven by your failure rate While the engine matters, the model matters more Try it on your own workload Benchmarking methodology FAQ Web search is table stakes for most LLM requests in order to overcome knowledge cutoffs. Labs and search providers are evolving fast to make search more effective and efficient, leaving all of us with a set of tricky decisions: take the native search some labs build in, or wire up a third-party engine like Exa, Parallel, or Perplexity? Is one search enough, and if not how long do I let the agent keep searching? Are more search turns worth the quality they buy?
We built live leaderboards to help you decide the best search configuration with data. See the data on our new Benchmarks page .
We benchmark all the combinations to find strengths and weaknesses
When setting up a search request, you have four decisions:
Model. Writes the exact query that gets submitted to the search engine and processes the results.
Engine. You can choose a specific engine or rely on the bundled engines offered by some labs. On OpenRouter, we offer Exa, Parallel, and Perplexity, alongside the native engines from labs like OpenAI, Anthropic, and Google.
Search method. Either you can perform the search before calling the model and pass the results in as context, or you can equip the model with a web search tool that it calls at its discretion.
Search budget. If you choose the search tool method, you can also give the model a budget for how many searches it’s allowed to do. This enables models to adjust the query if it doesn’t like the results or to do follow-up searches. Our runs use 1, 5, or 25 turns.
To comprehensively understand web search performance, we regularly run four benchmarks across multiple models, engines, and search configurations:
BrowseComp : hard fact-finding that takes real browsing
DeepSearchQA : multi-hop research questions
WideSearch : broad “fill in the whole table” collection
HLE : expert exam questions with search
Each page ranks configurations by quality, value, and speed, so you can make decisions on the factor most important to your workload. The leaderboards are live, so the numbers move as new runs land and new models and engines are added. Today’s leader is not guaranteed to be tomorrow’s. We won’t spend much time on today’s leaders in this post as we expect that to change over time. Instead, let’s look into what the data tells us about how to make a decision for your workload.
Search budget matters more than any other factor
Increasing the engine budget up from one turn improves quality more than any other single change you can make. To illustrate, here was our initial run of BrowseComp on Perplexity across three different budgets:
This pattern holds up across all providers we measured:
These runs cover BrowseComp only, using the server tool with ten results per search, no page fetching or code execution, and the latest qualifying run per configuration.
Increasing search depth is the cheapest way we’ve found to increase quality. Increasing from 1 turn to 25 roughly doubles the score while costing only 2.5-7 times more per question.
You may assume this universally slows down response time, but that’s not always the case. For example, Luna took 140 seconds per question at 1 turn and 111 seconds at 25. Of the 35 configurations we ran at both 1 and 5 turns, over a third were slower with fewer turns. All were OpenAI models. These models deal with the constrained search budgets with extra reasoning.
On the other hand, search depth can be detrimental to costs on easier tasks. For example, on HLE, GPT-5.6 Sol with Perplexity scored similarly between 1 turn and 25 turns, for triple the cost. If your searches tend to be simple, it may still be worth keeping your budget limited.
Your worst-case cost scenario is driven by your failure rate
The other situation where an expanded budget is detrimental is when the model is failing to find an answer. We found that models would exhaust their budget attempting to find an answer even though they would eventually fail.
The deepest attempt we recorded, 81 searches on a WideSearch table, was still graded incorrect. If your workload has a high failure rate, then reducing search depth is likely an efficient path for reducing costs.
While the engine matters, the model matters more
Once the budget is set, the next most important question is which model to use.
The table above shows the BrowseComp results at 25 turns, comparing frontier models versus budget models across search engines.
Varying the engine while holding the model constant changed the score by an average of 10 points while the average gap between frontier and cost-efficient models was larger at 15 points. Across engines, cost varied most for frontier models, where the priciest engine cost 2.5x the cheapest, versus 1.5x for cost-efficient models.
The reason a comparison like this is possible at all is that the server tool sits above the provider. Change the model in your request and the search behavior stays consistent, including for models whose provider ships no search of its own.
Of course, benchmarks are only a reference for possible performance. They tell you which configurations are worth trying and roughly what they cost. The cost and quality of these choices for your own real tasks will differ, so the highest-value thing you can do with these pages is treat them as a shortlist and then run your own questions through the top few.
Everything above is a request parameter you can set today on OpenRouter.
Web plugin. The web plugin runs a single search before the model starts writing, which is the fast, cheap option for questions that just need fresh facts.
Server tool. The server tool hands the model the search tool and lets it decide what to look up next, which is what you want when the answer takes several steps to find.
Engine. On OpenRouter, you set engine to exa , parallel , perplexity , or native ; auto tries native first before falling back to a third party.
Search budget. The top-level max_tool_calls request field caps how many agent turns it gets, meaning how many rounds of searching it may take before it has to answer, and max_results sets how many results come back each time.
A reasonable starting point: pick the suite closest to your task, take the cheapest configuration within a few points of the top score, then re-run your own evaluation set against the two or three rows above it to see whether the extra spend shows up in your results.
Every run goes through the public OpenRouter API against production endpoints, using our open source benchmark harness .
Isolated to search performance. To ensure we are comparing only the search configuration, we standardized on ten results per search, no page fetching, and no code execution. Reasoning is fixed per model, as reflected in the tables.
Scores are strict. Each evaluated answer is right or wrong against the official answer key, using an LLM judge where semantic comparison is needed. WideSearch also reports answer-item accuracy separately.
Cost and speed are per question. Cost is total spend, including grading, divided by evaluated questions. Speed is candidate generation time per evaluated question.
Each page shows the latest qualifying run for every configuration. A run qualifies once it completes a minimum number of questions, and new runs supersede old ones.
How do these scores compare with published vendor agent leaderboards?
They are not directly comparable. Most published tables for these benchmarks are measuring full agent products that combine search, full page fetching, and code tools. These leaderboards isolate search configurations: the model reads search result excerpts only, with page fetching and code tools off. This allows direct comparisons between configurations, but won’t maximize benchmark scores.
Which search engine should I pick?
It depends on the model and the task, which is why the pages exist. The gap between engines is large for some models and negligible for others, and a provider’s own native search is not automatically its best option. Check the live leaderboard for the suite closest to your workload, read cost and latency alongside the score, and re-check it over time, because the ordering changes as new runs land.
The leaderboards always show the latest qualifying run for each configuration, executed on OpenRouter’s benchmark harness against production endpoints. New runs supersede old ones on the page.
Tell us which engines or models we should benchmark next in #feedback on Discord.
