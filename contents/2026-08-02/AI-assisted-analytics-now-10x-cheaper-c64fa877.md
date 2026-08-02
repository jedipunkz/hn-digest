---
source: "https://motherduck.com/blog/openai-just-made-analytics-10x-cheaper/"
hn_url: "https://news.ycombinator.com/item?id=49147192"
title: "AI-assisted analytics now 10x cheaper"
article_title: "OpenAI Just Made Analytics 10x Cheaper"
author: "ryguyrg"
captured_at: "2026-08-02T18:55:08Z"
capture_tool: "hn-digest"
hn_id: 49147192
score: 1
comments: 0
posted_at: "2026-08-02T18:49:59Z"
tags:
  - hacker-news
  - translated
---

# AI-assisted analytics now 10x cheaper

- HN: [49147192](https://news.ycombinator.com/item?id=49147192)
- Source: [motherduck.com](https://motherduck.com/blog/openai-just-made-analytics-10x-cheaper/)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T18:49:59Z

## Translation

タイトル: AI 支援分析が 10 倍安くなりました
記事のタイトル: OpenAI だけで分析が 10 倍安くなりました
説明: 今週、OpenAI が GPT 5.6 Luna の価格を 80% 値下げしたため、AI と DB のコストを含め、低遅延の AI を活用した回答が 1 件あたり 0.5 ペニー未満で実現可能になりました。
データに関する質問については、GPT 5.6 Luna は、測定するには低すぎるインテリジェンスです。
よく文書化されたコンテキスト、厳密な評価、およびFA
[切り捨てられた]

記事本文:
新機能: ガイド、AI エージェントのコンテキスト レイヤー ブログを読む
製品コミュニティ 会社概要 価格 お問い合わせ ログイン スタート 無料製品 ハイパーテナンシー AI + MCP サーバー DuckLake 統合 ダイブ フライト 顧客対応分析 データ ウェアハウジング + BI vs Snowflake vs Postgres vs BigQuery vs Redshift vs Databricks vs ClickHouse 概要 DuckDB ユーザー向け ケーススタディ 信頼とセキュリティ サポート コミュニティ MotherDuck ブログ ビデオとストリームSlack コミュニティ イベント ダイブ ギャラリー YouTube DuckDB スニペット 無料 DuckDB ブック 無料 DuckLake ブック Agentic 時代の BI SQL チートシート Learn Data Outpost Duck Merch DuckDB ニュースレター Postgres は Vibes から Eval まで充実しています 会社概要 採用情報 インチキ研究 資料 価格 お問い合わせ ログイン 開始 無料 ブログに戻る OpenAI アナリティクスが 10 倍安くなりました
ブログに戻る OpenAI による分析が 10 倍安くなりました
分析タスクには、最もスマートな AI モデルは必要なくなりました。小型モデルでも前世代のフラッグシップモデルと同じくらいスマートになったため、エージェント分析に新しい戦略が開かれています。生のインテリジェンスよりもスピードとコストに重点を置くと、ユーザー エクスペリエンスが向上します。最新の小規模モデルと高速分析データベース エンジンを組み合わせると、AI と DB のコストを含めて、応答あたり 0.5 ペニー未満で、低レイテンシの AI を活用した回答が最終的に実現可能になります。
今週、OpenAI が GPT 5.6 Luna の価格を 80% 値下げしたとき、私たちは評価中にすぐにそれに気づきました。最初は、私たちも間違いを犯したのではないかとさえ思いました。エージェント SQL ベンチマーク (以前の結果はこちらをご覧ください) では、最大限の努力を行った Luna が、低価格の Gemini-3-Flash を抑えて現チャンピオンとなり、99.8% の精度を維持しながら、5 倍の低価格を実現しました。セマンティック モデリングのベンチマークでは、コストがさらに下がり、先週の Luna の 10 分の 1 になりました。モデル キャッシュの改善が促進されたと考えています

追加の 2 倍のブーストです。
セマンティック モデルのベンチマークの詳細については、今後の投稿で共有する予定ですが、ここでスニーク ピークを紹介します。
Opus 4.5 が登場したとき、AI モデルは、プロ規模のコーディング ワークロードにとって根本的により実用的なレベルに達しました。同じ大きさの変化が、GPT 5.6 Luna のパレート辺境の別の場所で起こりました。ルナは SQL が得意です。しかし、Opus とは異なり、信じられないほど高速で安価です。正確な SQL の生成がこれほど容易になったことはかつてありませんでした。
私たちは転換点にいます。最大限の努力を払った Luna は、価格性能曲線のまったく新しいセクションを占めます。 DeepSWE では、Luna は Fable より 3% 下、価格は 10 分の 1 以下です。
この価格/パフォーマンス曲線の劇的な変化を最大限に活用するには、私たちのアプローチに小さなステップを踏むだけでなく、より大きな変化を起こす必要があります。これらの新しい高速モデルをデータの世界でどのように適用するかについて、さらに深く考える必要があります。
まず最初に、大きなモデルを低労力設定で使用している場合は、Luna を最大で試してみる義務があるということです。最大限のインテリジェンスが必要ないと判断できるタスクであれば、非常に効率的です。労力を低く設定してこれらのタスクをすでに特定している場合は、代わりにそのモデルを交換してください。
もう 1 つの簡単な変更は、より多くの質問をすることです。これはジェボンズのパラドックスが実際に起きていることです。テクノロジーの改善はすぐにより多くのアプリケーションにつながります。これは、エージェントから返事を聞く前に 5 つの仮説を一度にテストするように見えるかもしれませんが、これほど高速なモデルを使用すると、より深いレベルの詳細を繰り返すことがはるかに簡単になります。顧客が最初にリストから選択するのを待つのではなく、関連する顧客の質問に対する回答を事前に取得する場合もあります。
AI が高速になると、より高速かつ低遅延のデータ プラットフォームの ROI が飛躍的に高まります。エージェントでw

orkflows では、LLM へのリクエストが長い間、合計時間のボトルネックとなっていました。エージェントの 1 ターンに数十秒かかった場合、データベースの高速化によるメリットはあまり影響を及ぼしません。
ワークロードが分析的に形成されている場合、分析データベースはトランザクションデータベースよりも 1000 倍高速になります。エージェントからの質問は数多くあります。 Luna のこの新しいリビジョンでは、トランザクション データベースがユーザー エクスペリエンスのボトルネックになりました。
同様に、分析ストアの起動に 30 秒かかる場合、低遅延 DB はその間にエージェントの質問 10 件に回答できた可能性があります。エージェントのワークロードは急激に変化するため、低遅延のサーバーレス アプローチが合理的です。特に顧客対応のエージェント エクスペリエンスを設計している場合は、応答性の高い分析エンジンとの違いを実感できるでしょう。エージェントがこのスピードでデータ主導の回答を提供できるようになると、どのような新しい製品機能やビジネス全体が可能になるでしょうか?
モデルの良さは、与えられたコンテキストによって決まります。特定のビジネスまたはドメインのすべての詳細を含むコンテキスト レイヤーの構築は、依然として信じられないほど高い活用力を持っています。しかし今では、そのコンテキストをより詳細に説明することで、弱いモデルでも解釈できるようになりました。弱いモデルが正確な SQL を作成できない場合、コンテキストは最も強力なモデルにとって十分な内容であれば十分です。ドメインの文書化に余分な時間を費やすことで、コストを削減し、回答を迅速化することができます。
データチームにはさらに多くの評価が必要です
データの世界では、歴史的に私たちのテストはデータ品質チェックでした。多くの場合、簡単に計算可能な不変条件が保持されているかどうかのみを確認します (重複する顧客 ID がない、NULL の注文価格がない、すべての注文が実際の製品 ID に結合されている)。エージェント分析タスクには、SQL チェック以上のものが必要です。自然言語による質問と、データに基づく正解が必要です。次に、年齢かどうかを評価します。

nt は質問を受け取り、ビジネス コンテキストとデータベース接続を使用して正しく答えることができます。
これらの eval の実行が 5 倍安くなりました。
自然な機会の 1 つは、これらの節約を使用して評価をより頻繁に実行することです。それぞれの新しいモデルがどのように動作するかを調査し、モデル内の設定を調整することもできます。以前は、おそらく単一のラボからのデフォルトを受け入れていました。
モデルの知能が変動するケースも捉えることができます。モデルのパフォーマンスは、基礎となるモデルの重みの関数ですが、モデルを提供するために使用されるインフラストラクチャの関数でもあります。キャパシティーが逼迫するとモデルのインテリジェンスが低下しますが、適切に調整された評価システムはそれをキャッチできる可能性があります。新しいモデルのリリース前の数日間であっても、1 日のトラフィックが最も多い時間帯であっても、ビジネスはより最適なモデルに移行する可能性があります。
ただし、モデルは 1 つの要素にすぎません。新しい組織コンテキストのそれぞれがどのように役立つか (または害をもたらすか) を測定することは重要です。 OpenAI のデータ チームは、コンテキスト層で回帰を捕捉することが評価フレームワークの大きな価値であることに気付きました。評価を毎週実行するだけでは、正確な組織知識を構築するのに十分なシグナルは得られません。
ビジネス自体も常に変化しています！新しい割引プログラムを追加すると、エージェントは収益の計算方法を正しく推測できなくなる可能性があります。グラフをボードに共有する前に、論理的なバグを見つけてください。
もう 1 つの自明の理は、フロンティア ラボに依存しないように評価ワークフローを設計する必要があるということです。私たちは評価用のカスタム ハーネスで OpenRouter を使用しているため、発売から数時間以内にすべてのフロンティア ラボの最新モデルを実行できるようになります。ハーネス層またはコンテキスト層で知的財産を開発し、モデルを商品として扱うことができれば、この現代の AI 経済で最大の利益が得られます。スイッチングコストを維持することで、

ロックインを回避し、投資収益率が得られたらすぐに次の上位モデルに切り替えることができます。
OpenAI と他の研究所は、測定できないほど安価なインテリジェンスを提供するために長い間競争してきました。データに関する質問については、GPT 5.6 Luna が正式に登場したようです。次は、モデルに基づいてシステムを構築し、生のインテリジェンスを顧客に対する正確で関連性の高い回答に変換します。十分に文書化されたコンテキスト、厳密な評価、高速分析エンジンが現在、決定要因となっています。
他の MotherDuck アップデートを購読する 前の投稿を送信する
2026/07/29 - ハミルトン・ウルマー、ティル・デーメン、ジェイコブ・マトソン、ギャレット・オブライエン
コンテキストはウェアハウスに属します
AI エージェント用の MotherDuck のコンテキスト レイヤーである Guides の紹介。より正確なクエリ、より少ないトークン支出、そして組織全体の分散コンテキスト。
Figma for Agents: Airflow の作成者が AI をどのように調整するか ft. Maxime Beauchemin
Airflow と Superset の作成者である Maxime Beauchemin が、データ エンジニアリングに AI エージェントをどのように使用しているかを共有します。Agor を使用した共有キャンバス上のエージェントのオーケストレーション、コンテキスト レイヤーとセマンティクス、Okta for Agents のセキュリティ、AI アシスタントによる CEO の仕事の自動化などです。
イベント YouTube Data Outpost Duck Merch コミュニティと OSS ビデオとストリーム ダイブ ギャラリー DuckDB ニュース DuckDB スニペット 行動規範を学ぶ 会社

## Original Extract

Since OpenAI slashed the price of GPT 5.6 Luna by 80% this week, low latency AI-powered answers are finally feasible for less than half a penny per answer, AI and DB costs included.
For data questions, GPT 5.6 Luna is intelligence too cheap to meter.
Well documented context, rigorous evals, and a fa
[truncated]

New: Guides, the context layer for AI agents Read the blog
PRODUCT COMMUNITY COMPANY DOCS PRICING CONTACT US LOG IN START FREE PRODUCT Hypertenancy AI + MCP Server DuckLake Integrations Dives Flights Customer-Facing Analytics Data Warehousing + BI vs Snowflake vs Postgres vs BigQuery vs Redshift vs Databricks vs ClickHouse Overview For DuckDB Users Case Studies Trust & Security Support COMMUNITY MotherDuck Blog Videos & Streams Slack Community Events Dive Gallery YouTube DuckDB Snippets Free DuckDB Book Free DuckLake Book BI in the Agentic Era SQL Cheatsheet Learn Data Outpost Duck Merch DuckDB Newsletter Postgres is Full From Vibes to Eval COMPANY About Us Careers Quacking Research DOCS PRICING CONTACT US LOG IN START FREE GO BACK TO BLOG OpenAI Just Made Analytics 10x Cheaper
GO BACK TO BLOG OpenAI Just Made Analytics 10x Cheaper
Analytics tasks no longer need the smartest AI models. Now that even small models are as smart as prior generation flagships, new strategies open up for agentic analytics. Focusing on speed and cost over raw intelligence leads to a better user experience. If you pair a modern small model with a fast analytical database engine, low latency AI-powered answers are finally feasible for less than half a penny per answer , AI and DB costs included.
When OpenAI slashed the price of GPT 5.6 Luna by 80% this week, we noticed it immediately in our evals. At first, we even thought we had made a mistake! On an agentic SQL benchmark (see our prior results here ), Luna on max effort displaced Gemini-3-Flash on low as the reigning champion, maintaining 99.8% accuracy, at a 5x lower price point . In a semantic modeling benchmark, costs dropped even further, getting 10x cheaper than last week’s Luna! We believe model caching improvements drove that extra 2x boost.
We’ll share details on that semantic model benchmark in an upcoming post, but a sneak peek is here !
When Opus 4.5 landed, AI models passed a threshold where they became fundamentally more practical for professional-sized coding workloads. That same magnitude of shift just happened, but at a different spot on the Pareto frontier with GPT 5.6 Luna. Luna is great at SQL. But unlike Opus, it is also incredibly fast and cheap. Never before has generating accurate SQL been anywhere near this approachable.
We are at an inflection point. Luna with max effort occupies a brand new section of the price/performance curve. In DeepSWE, Luna sits 3% below Fable at under 1/10th the price.
To take full advantage of this dramatic shift of the price/performance curve, there are both small steps to take and larger shifts to make in our approach. We need to think bigger about how to apply these new faster models in the data world.
The first thing is that if you were using a large model on a low effort setting, you owe it to yourself to try Luna on max. It is just so efficient for any task that you can determine doesn’t need maximum intelligence. If you already have identified those tasks by setting your effort to low, swap that model out instead.
Another easy change is to just ask more questions. This is the Jevons Paradox in action - improvements in technology immediately lead to more applications. That could look like testing 5 hypotheses at once before hearing back from an agent, but with models this fast, iterating into deeper levels of detail becomes far easier. Maybe you pre-fetch answers to relevant customer questions rather than wait for them to pick from a list first.
Once AI is fast, the ROI of a faster and lower latency data platform jumps. In agentic workflows, the request to the LLM has long been the total time bottleneck. If one turn of the agent took 10s of seconds, the benefits of a faster database just weren’t impactful.
Analytical databases can be 1000x faster than transactional ones if your workload is analytically shaped. Many agent questions are. With this new revision of Luna, your transactional database just became your user experience bottleneck.
Likewise, if your analytical store takes 30 seconds to spin up, a low latency DB could have answered 10 agent questions in that time. Agent workloads are bursty, so a low latency serverless approach makes sense. Especially if you are designing customer facing agentic experiences, you’ll feel the difference with a responsive analytical engine now. What new product features or even whole businesses are possible once an agent can provide data driven answers at this speed?
The models are still only as good as the context they’re given. Building a context layer with all the details of your specific business or domain remains incredibly high leverage. Now though, it pays dividends to be more detailed in that context so that a weaker model can interpret it. When weak models couldn’t write accurate SQL, context just needed to be good enough for the strongest of models. Putting in that extra time documenting your domain can slash costs and speed up answers.
Your data team needs more evals
In the data world, historically our tests were data quality checks. Often we only check if easily computable invariants hold (no duplicate customer ids, no NULL order prices, every order joins to a real product id). For agentic analysis tasks, we need more than just SQL checks. We need natural language questions and the correct answer based on the data. We then evaluate if an agent can take the question and use business context and a database connection to answer it correctly.
Running those evals just got 5x cheaper.
One natural opportunity is to use those savings to run evals far more often. We can explore how each new model performs and even tune settings within models, where before perhaps we accepted the defaults from a single lab.
We could even catch cases where model intelligence fluctuates. Model performance is a function of the underlying model weights, but also the infrastructure used to serve it. As capacity gets tight, model intelligence is reduced, and a well calibrated eval system could catch that. Your business could jump to a more optimal model, whether that is in the days before a new model release or during the highest traffic times of the day.
The model is only one factor though. It is valuable to measure how each new piece of organization context helps (or hurts!). OpenAI’s own data team found that catching regressions in their context layer was a huge value of their eval framework. Running evals weekly just won’t provide enough signal to build accurate organizational knowledge.
The business itself is constantly changing too! If you add a new discount program, your agents may not be able to deduce how to calculate revenue correctly anymore. Catch those logical bugs before you share that graph to the board!
Another truism is that we need to design our eval workflows to be frontier lab agnostic. We use OpenRouter in a custom harness for our evals so that within hours of launch we can run any of the latest models from all the frontier labs. We get the most benefit in this modern AI economy if we can develop intellectual property at the harness or context layers and treat the model as a commodity. By keeping switching costs low and avoiding lock-in, we can jump to the next top model as soon as it provides a return on investment.
OpenAI and the other labs have long been racing to provide intelligence too cheap to meter. For data questions, GPT 5.6 Luna legitimately looks like we’ve arrived. Now it’s time to build the systems around the model to convert raw intelligence into accurate, relevant answers to customers. Well documented context, rigorous evals, and a fast analytical engine are now the determining factors.
Subscribe to other MotherDuck Updates Submit PREVIOUS POSTS
2026/07/29 - Hamilton Ulmer, Till Döhmen, Jacob Matson, Garrett O'Brien
Context belongs in the warehouse
Introducing Guides, MotherDuck's context layer for AI agents. More accurate queries, lower token spend, and distributed context for your entire organization.
Figma for Agents: How Airflow's Creator Coordinates AI ft. Maxime Beauchemin
Maxime Beauchemin, creator of Airflow and Superset, shares how he uses AI agents for data engineering: orchestrating agents on a shared canvas with Agor, context layers and semantics, Okta for Agents security, and automating CEO work with AI assistants.
Events YouTube Data Outpost Duck Merch Community and OSS Videos & Streams Dive Gallery DuckDB News DuckDB Snippets Learn Code of Conduct Company
