---
source: "https://www.altimate.ai/blog/where-ai-agents-belong-in-data-engineering-the-correctness-layer/"
hn_url: "https://news.ycombinator.com/item?id=48746115"
title: "Where AI Agents Belong in Data Engineering: The Correctness Layer"
article_title: "Where AI Agents Belong in Data Engineering: The Correctness Layer | Altimate AI"
author: "zazuke"
captured_at: "2026-07-01T13:47:52Z"
capture_tool: "hn-digest"
hn_id: 48746115
score: 3
comments: 0
posted_at: "2026-07-01T13:07:11Z"
tags:
  - hacker-news
  - translated
---

# Where AI Agents Belong in Data Engineering: The Correctness Layer

- HN: [48746115](https://news.ycombinator.com/item?id=48746115)
- Source: [www.altimate.ai](https://www.altimate.ai/blog/where-ai-agents-belong-in-data-engineering-the-correctness-layer/)
- Score: 3
- Comments: 0
- Posted: 2026-07-01T13:07:11Z

## Translation

タイトル: データ エンジニアリングにおける AI エージェントの属する場所: 正確性レイヤー
記事のタイトル: データ エンジニアリングにおける AI エージェントの所属: 正確性層 | AIをアルティメートする
説明: モデルは常に変化しており、新しいより優れたモデルが数か月ごとに登場するため、モデルに過度に依存する必要がないのは素晴らしいことです。ツールが優れているほど、単一のツールへの依存度が低くなります

記事本文:
データ エンジニアリングにおける AI エージェントの属する場所: 正確性層 | Altimate AI プラットフォーム ▼ 製品 ▼ ユースケース ▼ 顧客価格リソース ▼ 企業は 1,000 万トークンを無料で請求 ブログに戻る ☀ ☾ Altimate · ハーネスエンジニアリング · データ エンジニアリング AI エージェントがデータ エンジニアリングに属する場所: 正確性層
モデルは常に変化しており、数か月ごとに新しいより優れたモデルが登場するため、モデルに過度に依存する必要がないのは素晴らしいことです。ツールが優れているほど、単一のモデルへの依存度が低くなります。これが、決定論的ハーネスが重要である理由でもあります。つまり、実行しているモデルに関係なく、出力を再現し系統を追跡できる正確性レイヤーです。これは、検証が実際の仕事であるメンテナンス中やプロジェクトの拡張中に特に当てはまります。
危険なのは、クラッシュやエラー メッセージだけではなく、壊れていない間違った番号です。これはクリーンなクエリかもしれませんが、重複した行が発生します。
この記事では、データ エンジニアリングにおける AI エージェントの 3 つのレベル、AI が最良の結果をもたらすようにプロジェクトを構築する方法、決定論的なコアを持つ専用エージェントが実際に信頼できる高品質のパイプラインの構築にどのように役立つかを説明します。そして、それが爆発半径解析でどのように機能するかを示す実際の例を見てみましょう。
データ エンジニアリングにおける AI エージェントの 3 つのレベル
データ エンジニアリングにエージェントを使用する必要があるのはなぜですか?また、エージェントはどのレベルで私たちを生産的に支援してくれるのでしょうか?人間と同様、LLM には常にある程度のエラー許容度があるため、より自信を持ってコードを作成する方法が必要です。
チャットフェーズ、自律型の専用ツール
エージェントが私たちをサポートできる信頼のレベルやレベルはさまざまです。
初期チャットフェーズ : クロードまたは ChatGPT を促す開発。

モデルは、アクセスできる内容に基づいてコンテキストを理解しようとします。すべてを最初からスキャンする必要があるため、かなりの量のトークンが必要です。
自律的アプローチ では、Claude Code または Codex も人間が持つツール (主に端末上の CLI) にアクセスできるため、psql を使用して Postgres にクエリを実行したり、DuckDB を使用して S3 または Parquet から読み取ってクエリやデータを検証したりすることが可能になります。はるかに高品質の結果が得られます。
目の前のタスクに専任のエージェント。たとえば、データの場合、ツールは dbt を知っているか、SQL コードを決定論的にトランスパイルする方法を知っています。これは、トレーニング データのみからではなく、実際のツールを使用して、はるかに高速かつ確実にトランスパイルすることを意味します。 「一般的な」エージェントでは提供できない組み込みのチェックと機能。
データ エンジニアリングにおける AI エージェントの 3 つのレベルを紹介する
理想的には、専用ツールを常に使用したいのですが、常に専用ツールがあるわけではありません。
DE ライフサイクルのどの段階で各レベルが実際に役立つのか
BI ダッシュボード vs. データ パイプラインの接続、ソース取り込みの作成、それとも保守?データ エンジニアリングの場合、専用のエージェント ツールがあるかどうかだけが問題ではなく、AI エージェントがデータ エンジニアリング ライフサイクルのどの部分でデータ エンジニアやアナリスト、さらには場合によってはドメインの専門家を最も支援できるかということも問題になります。
ライフサイクルには、取り込み部分、ETL、またはビジネスを詳細に理解することが含まれますか、それとも結果を視覚化するだけですか?それとも、夜間の ETL エラーが発生した場合のメンテナンス、またはデータ ライフサイクル全体をカバーする必要がありますか?
一般に、後ほど詳しく説明する前に、エージェントは全サイクルで私たちをサポートしてくれますが、それは常にあなたが誰で、どのような役割を果たしているかによって異なります。知識も年功もない状態でゼロから構築するのは危険です。なぜ？生成されたコードが正しいかどうかを検証できないからです。サイドプロジェクトや概念実証には問題ありませんが、

実際の生産用ではありません。
AI を扱うためのエンジニアリングの規律とは何ですか?
あまり技術的ではない部分、つまりエージェントを正しい方向に導く方法もあります。特に大規模なプロジェクトや組織で安全に使用したい場合は、指導なしでそのまま実行することはできません。
エージェントが活躍できる明確なプロジェクト構造。より多くの値を与えるほど、この作業に使用されるトークンの数が減り、プロジェクト全体でより調整されるようになります。 (UV init のような決定論的なワークフローが最適なもう 1 つの理由は、ワークフローが常に同じになるためです)。
ツールの使用方法に関する明確な指示 (エージェント スキル、スーパーパワーなど) を使用して構築します (基本的に CLI と API ドキュメントを提供します)。とにかくこれが作業の大部分です。それがデータ アーキテクチャであり、何かを構築する前に仲間とブレインストーミングを行うことで、最初に重要な洞察を見逃してエージェントを誤った道に進ませないようにすることができます。また、現実的である必要があります。「正確であること」や「最先端のものを使用すること」を促しても、モデルがトレーニングされたものより正確になったり、最新のものになったりするわけではありません。したがって、かなり新しいアーキテクチャの場合は、これらのリンクとヒントを提供することが必須です。
モジュール形式でプロジェクトを設定するため、エージェントが小さな変更を加えてもプロジェクト全体を壊すことができないため、すべてが相互に依存する依存関係地獄のようなシナリオに陥ることがなくなります。
宣言的アプローチを使用し、どのようにではなく何を説明する構成を使用します。これにより、これらの構成をエージェントと共同作業し、バージョン管理し、簡単に元に戻したり変更したりできるほか、実装ロジックを実際のビジネス ロジックから切り離すことができます。
これらの手順を実行すると、最新のエージェントを最大限に活用できます。モデルはそれほど重要ではないと思いますが、構造体は重要です

マリオが言うように、ワークフローのアプローチも同様です。たとえば、(単一行を記述する前のプロセス) を綿密に計画し、間違った道をたどる可能性のある実装を記述する前にモデルを修正します。
また、考えすぎないでください。しかし、これは単なるワークフローであり、エージェントと協力するためのソフト スキルと規律を学ぶだけです。実際のプロジェクトではどうなるでしょうか?
重要なのは、仕事を増やすことではなく、AI を活用することです。たとえば、ほとんどの開発者は以前はこの問題について考えていました。今日では、ほとんどの人が PR に夢中になっています。 AI ツールが改善されると、AI はより正確で、レビューやイテレーションが少なくて済む、より高品質なコードを提供できるようになります。つまり、PR が減り、開発者が経験する作業が減ります。
データ エンジニア向けの正確性レイヤー
重要な洞察は、AI エージェントが正確さのために「ループ内の人間」、または正確さレイヤーをサポートする必要があるということです。そして、より多くのコードを検証するためにさらに多くの作業を行うのではなく、プロセスに自信を持って、生成されるコードが検証され、最終的に正しいことを知る必要があります。
しかし、より「正しい」作業とそれを検証できるレイヤーを取得するにはどうすればよいでしょうか?最大の議論は、完全な決定論的検証アーキテクチャです。たとえば、Altimate Code は、エージェントを上部の確率的なレイヤーとその下の決定論的な Rust/TS レイヤーに分割し、解析、検証、等価性チェックなどの実際の SQL 操作を実行するため、これらの質問に関してエージェント自体を信頼する必要はありません。
Altimate コードが確率的エージェント、決定的ハーネス、決定的コアを使用して構築される方法の例 |記事「正確性レイヤー: データ エージェントに決定性が必要な理由」からの画像
たとえば、Altimate コードは、確率的エージェント、決定的ハーネス、および決定的コアに基づいて構築されています。確率的エージェント wi

LLM は、意図を読み取り、戦略を選択し、SQL を作成し、結果を要約し、問題が発生した場合に回復するという創造的な作業を行います。
境界の下には、すべてのツール呼び出しをインターセプトする TypeScript レイヤーである決定的ハーネスがあります。ディスパッチャーは、呼び出しが実行される前に hasNativeHandler をチェックし、それをネイティブの決定的ハンドラーにルーティングするか、モデルに戻します。これらのハンドラーはロジック自体を再実装するのではなく、napi-rs バインディングを介して接続された AST およびスキーマ上の純粋な関数として SQL 操作を公開する Rust エンジン (altimate-core) である決定論的な core を呼び出します。解析、検証、トランスパイル、クエリの等価性のチェック、スキーマの比較、列系統の抽出、ウェアハウス間での行の比較 - これらはすべてミリ秒未満で実行され、毎回同じ入力に対して同じ回答を返します。
コンパイラと同様に、エージェントは 2 つのクエリが同等であるかどうか、または上流に列が存在するかどうかを判断することはありません。代わりに、型チェッカーがプログラムが推測ではなくコンパイルされたことを証明するのと同じ方法で、解析された AST とスキーマに対して証明する関数を呼び出します。
正確性レイヤーが追加の検証を追加する方法
これは、事実のチェックが実行され、出力が正しいか、直接修正できるバグがあるかのいずれかであるため、出力のレビューを容易にする区別です。残りは人間が再検証できます。コードを手書きするのをやめ、人間がチェックするよりも早く承認してしまうというジレンマについては、「You Are the Trust Layer」で詳細を読むこともできます。
注: 別の要因があります: 間違っていることです
エージェントの裸の使用は安いかもしれませんが、それは彼らが間違っているまでであり、その後のコストは無制限です。
トークンのより良い使用のための改善
Altimate、または決定論的な機能と統合されたデータ エンジニアリング エージェント

動作方法を理解すると、トークンを節約し、トークンを無駄にしないようにすることができます (Twitter/X で人気のある tokenmaxxing の逆で、できるだけ多くのトークンを使用し、エージェントを常に実行します)。なぜなら、大企業ではトークンのコストが実際の予算のポイントとなるからです。
トークンの速度を低下させる簡単な方法は、使用するトークンと単語そのものの数を減らすようにモデルに指示することです。 caveman はその良い例ですが、Altimate Code と組み合わせて CLAUDE.md 、Codex、または選択したモデルに単一のプロンプトを追加することもできます。
データ系統のトレースとその Web UI を示す Altimate コードの例。
2 番目の、それほど明白ではないコストがあります。それは、トークン自体が安定したユニットではないということです。 Anthropic が Opus 4.7 を出荷したとき、4.6 では X トークンのコストがかかっていた同じプロンプトのコストがおよそ 1.4 倍になりました (同じ入力、同じ答え、より多くのトークン、トークンあたりの価格は同じ)。
'26 年の The Great Token Heist で、Altimate チームは、「トークンあたりのコストは最適化するのに間違った数値である」と主張しています。メーター自体はベンダーの次のモデル更新によって変わる可能性があり、代わりに追跡すべきはタスクあたりのコストだからです。私も完全に同意します。これは、決定論的な関数呼び出しがすべてのタスクにモデル/トークンを使用しないことで不安定性を回避し、コストを削減する場所です。
この章では、データ エンジニアリングのための典型的な AI エージェントの使用例について説明します。
それらはたくさんあります。これらを使用して、自分自身やチームを教育したり、本番データ パイプラインを構築したり、データ アプリを構築したり、新しい革新的な方法 (通常は React やその他の JavaScript フレームワークを使用した HTML Web ページ) でデータを視覚化したりすることができます。ただし、一般に、ユースケースは次のアプローチに当てはまります。
新しいプロジェクトを最初から開始する例: よりオープン ソースを使用してデータ ランドスケープを構築する。
既存のプロジェクトまたはデータ ウェアハウスの拡張 : 新しいデータ pip の追加

エラインズ。
現在のセットアップを維持する: 更新して、変更が加えられたときに引き続き動作することを確認します。
移行 : あるデータベースまたはツールから次のデータベースまたはツールに移行します。
盲点の発見 : 2 つの似たような ID が結合に誤って使用されたり、夜間のロードで欠落した列のデータが欠落したり、あるいはその中間の可能性があります。エージェントがこれらのチェックを行うことができれば、非常に有益になります。 CLI、モデル コンテキスト レイヤー、および決定論的ツールへのアクセスが増えると、これらのことが本当に可能になります。
以下では、列を変更して既存の倉庫を拡張および変更し、Altimate コードを使用して爆発半径の評価を行う手順を説明します。
ショーケース: Blast-Radius の例
爆発半径は、潜在的な損害の範囲を指します。たとえば、家の壁を壊す前に、壁の後ろに配管があるか、内部に電気配線があるか、または上の床を支えているかどうかを知りたいと思います。
大量の ETL を含むデータ ウェアハウスやデータ プロジェクトにも同じことが当てはまります。たとえば、データ エンジニアが注文を order_items に結合し、 order_total を合計することによってテーブル fct_orders をクリーンアップするとします。コンパイルされ、dbt テストに合格し、エラーは何もありません。ただし、結合によって粒度が変更されるため、複数の項目を含む注文はすべて次のようになります。

[切り捨てられた]

## Original Extract

With ever-changing models, new and better ones coming out every few months, it's great if we don't have to rely on them too heavily. The better your tooling, the less dependent you become on any singl

Where AI Agents Belong in Data Engineering: The Correctness Layer | Altimate AI PLATFORM ▼ PRODUCTS ▼ USE CASES ▼ CUSTOMERS PRICING RESOURCES ▼ COMPANY CLAIM 10M TOKENS FREE Back to blog ☀ ☾ Altimate · harnessengineering · Data Engineering Where AI Agents Belong in Data Engineering: The Correctness Layer
With ever-changing models, new and better ones coming out every few months, it's great if we don't have to rely on them too heavily. The better your tooling, the less dependent you become on any single model. That's also why the deterministic harness matters: a correctness layer that lets you reproduce outputs and trace lineage regardless of which model you're running underneath. This is especially true during maintenance or extending the project, where verification is the real job.
The danger isn't only a crash or an error message, but a wrong number that didn't break. It might be a clean query, but it introduces duplicated rows.
In this article, we go through the three levels of AI agents in data engineering, how to structure projects so the AI delivers its best outcomes, and how dedicated agents with a deterministic core help us build higher-quality pipelines — ones we can actually trust. And we look at a practical example of how it works with a blast radius analysis.
The Three Levels of AI Agents in Data Engineering
Why should we use agents for data engineering? And at what levels can agents help us productively? As LLMs will always have some error tolerance, as humans do too, we need a way to be more confident in producing the code.
Chat-phase, Autonomous and Dedicated Tooling
There are different levels of confidence and levels on which the agents can help us.
The initial chat-phase : the development where we prompt Claude or ChatGPT. The model tries to understand the context based on what it has access to. It takes a decent amount of tokens, as it needs to scan everything from scratch.
The autonomous approach , where Claude Code or Codex also have access to the tools humans have, mostly the CLI on the terminal, making it possible to query Postgres with psql or read from S3 or Parquet with DuckDB to verify queries and data. A much higher quality outcome.
Dedicated agents for the task at hand. E.g., for data, the tools know dbt or know how to transpile SQL code deterministically, meaning not from training data only, but with an actual tool that does it much faster and more reliably. Built-in checks and features a "general" agent can't provide.
Showcasing the three levels of AI agents in data engineering
Ideally, we'd want to always use the dedicated tools, but there isn't always one.
Where in the DE Lifecycle Each Level Actually Helps
BI Dashboards vs. Plumbing the Data Pipelines, or Creating Source Ingestions, or Maintaining? For data engineering, the question is not only if there is dedicated agent tooling, but also on what part of the data engineering lifecycle AI agents can help data engineers and analysts the most, and potentially even domain experts?
The lifecycle contains the ingestion part, ETL, or understanding the business in great detail, or is it just to visualize the result? Or should it cover maintenance in case of overnight ETL errors, or the full data lifecycle?
In general, before we go into more details later, agents can help us on the full cycle, but it always depends on who you are and what role you play. Building from scratch with no knowledge or seniority is dangerous. Why? Because they can't verify if the produced code is correct. Okay for a side project or a proof of concept, but not for actual production.
What's the Engineering Discipline for Working with AI?
There's also a part that is less technical, a way of guiding the agents in the right direction. Especially if we want to safely use it in large projects or organizations, we can't just let it run without guidance.
clear project structure in which the agents can flourish. The more is given, the fewer tokens are used for this work, and it will be more aligned across the project. (Another reason a deterministic workflow such as uv init is best, because it will always be the same).
build with clear instructions ( agentic skills , superpowers , etc.) on how the tools are used (basically providing CLIs and API documentation). This is the bulk of the work anyway. That's the data architecture, the brainstorming with fellow humans before you build something, instead of missing a key insight in the beginning and then letting the agent run down the wrong path. Also, be realistic: prompting "be correct" or "use state-of-the-art" won't make it more correct or more state-of-the-art than the model was trained on. So if it's a rather new architecture, it's a must that you provide these links and hints.
set up the project in a modular fashion, so the agents cannot break the whole project if they make a small change, so you don't end up in a scenario like dependency hell with everything dependent on each other.
use a declarative approach , with descriptive configuration that says the what and not the how, so that you can collaborate on these configs with the agents, version them, and easily revert or change something, as well as decouple the implementation logic from the actual business logic.
With these steps, you can get the best out of the agents of today. I'd say the model matters less, but the structure does, and as Mario says, so does the workflow approach. For example, extensively plan (the process before writing a single line) and correct the model before any implementation that could lead down the wrong path is written.
Also, don't overthink it. But this is only the workflow and learning the soft skills and discipline of working with agents . How does that look in a real-world project?
The key is to get use out of AI, not to get more work. E.g., most developers used to think about the problem. Today, most drown in PRs. When the AI tooling gets better, AI can provide more quality code that is correct, that needs less review or fewer iterations, which means fewer PRs and less work for the developers to go through.
The Correctness Layer for Data Engineers
A key insight is that AI agents should support the "human in the loop" for correctness , or a correctness layer . And rather than making more work to verify more code, we should be confident in the process and know that the code it produces is verified and ultimately correct.
But how do we get more "correct" work and a layer in which we can verify it? The biggest argument is a deterministic-validation architecture in full. E.g., Altimate Code splits the agent into a probabilistic layer on top and a deterministic Rust/TS layer underneath that does the actual SQL ops such as parsing, validating, and equivalence checks, so that the agent itself never has to be trusted on those questions.
An example of how Altimate Code is built with its probabilistic agent, deterministic harness, and deterministic core | Image from the article The Correctness Layer: Why Data Agents Need Determinism
Altimate Code, for example, is built on a probabilistic agent, deterministic harness, and deterministic core. The probabilistic agent with the LLM does the creative work of reading intent, picking a strategy, drafting SQL, summarizing results, and recovering when something goes wrong.
Below the boundary sits the deterministic harness , a TypeScript layer that intercepts every tool call: a dispatcher checks hasNativeHandler before the call runs, and routes it either to a native, deterministic handler or back to the model. Those handlers don't reimplement logic themselves, they call into the deterministic core , a Rust engine ( altimate-core ) that exposes SQL operations as pure functions over ASTs and schemas, wired in via napi-rs bindings. Parsing, validating, transpiling, checking query equivalence, diffing schemas, extracting column lineage, diffing rows across warehouses — all of it runs sub-millisecond, and all of it returns the same answer on the same input, every time.
Like a compiler, the agent never decides whether two queries are equivalent or a column exists upstream. Instead, it calls a function that proves it against the parsed AST and the schema, the same way a type-checker proves a program compiles rather than guessing.
How the correctness layer adds additional verification
That's the distinction that makes the output easier to review, as factual checks have been run and the output is either correct, or there's a bug that it can fix directly. The rest a human can re-verify. On the dilemma of having stopped to hand-write code and approving it faster than humanly possible to check, you can also read more at You Are the Trust Layer .
Note: There's another factor: being wrong
Bare agent use might be cheap, but only until they're wrong, and then the cost is unbounded.
Improvements for Better Usage of Tokens
Altimate, or data engineering agents that have deterministic functions and integrated understanding of how to work, can help you save tokens and be token lean (the opposite of tokenmaxxing , which is popular on Twitter/X, using as many tokens as possible and having an agent running at all times). Because in large enterprises, token costs are a real budget point.
To slow down the tokens, an easy trick is to instruct the model to use fewer tokens and words itself - caveman is a good example of that, but you can also add a singular prompt to your CLAUDE.md , Codex, or model of choice in combination with Altimate Code.
An example of Altimate Code showing a trace of data lineage and a web UI for it.
There's a second, less obvious cost: the token itself isn't a stable unit. When Anthropic shipped Opus 4.7, the same prompt that cost X tokens on 4.6 started costing roughly 1.4X (same input, same answer, more tokens, same price per token).
In The Great Token Heist of '26 , the Altimate team makes the case that " cost-per-token is the wrong number to optimize ", since the meter itself can move with a vendor's next model update, and what we should track instead is cost-per-task . I fully agree, and this is where deterministic function calls work around that volatility by not using a model/tokens for every task, making it less expensive.
In this chapter we go through typical AI agent use cases for data engineering.
There are many of them. You can use them to educate yourself or your team, build production data pipelines, build data apps, and visualize your data in new innovative ways (usually HTML web pages with React and other JavaScript frameworks). But in general, the use cases fit into these approaches:
Start a new project from scratch example : Building a data landscape with more open source.
Extending an existing project or data warehouse : Adding new data pipelines.
Maintaining current setup : Update and verify it still works when changes come in.
Migration : Migrate from one database or tooling to the next.
Finding the Blind Spots : Two similar-sounding IDs might be wrongly used for a join, or missing data in a column that got missed in a nightly load, or anything in between. If agents can do these checks, that would be super beneficial. With more access to CLI, Model Context Layer, and deterministic tooling, these things are truly possible.
Below we go through extending and changing an existing warehouse with a change of column, and using Altimate Code to give us a Blast-radius assessment.
Showcases: Blast-Radius Example
A Blast-radius refers to the potential extent of damage . For example, before you knock down a wall in your house, you want to know if there's plumbing behind it, electrical wiring within it, or if it's holding up the floor above.
The same is true for a data warehouse or a data project with lots of ETL. For example, if a data engineer cleans up the table fct_orders by joining orders to order_items and summing order_total . It compiles, the dbt tests pass, nothing errors. But the join changes the grain, so any order with several line items now gets

[truncated]
