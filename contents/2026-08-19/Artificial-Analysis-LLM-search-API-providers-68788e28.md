---
source: "https://artificialanalysis.ai/agents/search-api"
hn_url: "https://news.ycombinator.com/item?id=49359449"
title: "Artificial Analysis: LLM search API providers"
article_title: "Artificial Analysis Search Index: Search API Benchmark & Leaderboard | Artificial Analysis"
image: ""
author: "j-bu"
captured_at: "2026-08-19T10:19:17Z"
capture_tool: "hn-digest"
hn_id: 49359449
score: 1
comments: 0
posted_at: "2026-08-19T10:17:25Z"
tags:
  - hacker-news
  - translated
---

# Artificial Analysis: LLM search API providers

- HN: [49359449](https://news.ycombinator.com/item?id=49359449)
- Source: [artificialanalysis.ai](https://artificialanalysis.ai/agents/search-api)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T10:17:25Z

## Translation

タイトル: 人工分析: LLM 検索 API プロバイダー
記事のタイトル: 人工分析検索インデックス: 検索 API ベンチマークとリーダーボード |人工分析
説明: 検索品質、ベンチマーク精度、待ち時間、コストの観点から検索 API プロバイダーを比較します。

記事本文:
人工分析検索インデックス: 検索 API ベンチマークとリーダーボード |人工解析 人工解析 K 人工解析モデル
K 人工分析検索インデックス: 検索 API ベンチマークとリーダーボード
Search API を使用すると、AI エージェントが Web から情報を取得し、応答を最新のソースに基づいて行うことができます。これらは、事実性と全体的な出力を向上させるために、詳細な調査、コーディング、一般知識の作業など、幅広いエージェント アプリケーションで使用されます。検索 API は、検索する基になる Web インデックス、検索結果がモデルに表示される方法、コスト、速度、取得品質の間で行われるトレードオフなど、いくつかの重要な点で異なります。開発者が次の AI エージェント用の検索ツールを選択する際に、より多くの情報に基づいた意思決定を行えるよう、7 つのプロバイダーにわたる 12 の検索 API 製品のベンチマークを実施しています。
ベンチマークの定義、スコアリング、データ処理の詳細については、方法論のページを参照してください。
言語モデル API エンドポイント 人工分析手法 人工分析 検索インデックス 記事の起動 あばらみエージェント ベンチマーク タスク 候補モデル GPT-5.6 Luna (中型) ターン 0 / 25 検索プロバイダー web_search ネイティブ ペイロード、最大 10 件の結果 グレーディング グレーダー avg F1 0.79 検索 API プロバイダー 定数 同じタスク、同じモデルとグレーダー (GPT-5.6 Luna、中型)、同じあばらみハーネス(web_search · web_fetch · 終了、最大 25 ターン)。検索プロバイダーのみが異なります。再起動 0.5× 1× 2× ハイライト
人工分析検索インデックス
公開ベンチマーク全体にわたる検索対応の回答品質。
人工分析検索インデックス
12 プロバイダー中 12 NEW モデルのみ (33) 人工分析検索インデックス
DeepSearchQA (平均 F1)、BrowseComp (正確な回答精度)、および AA-Omnisci における各検索 API プロバイダーのスコアの等加重平均

ence (精度)、0 ～ 100 のスケールで表示されます。すべての結果で、同じ候補回答モデル、ハーネス、設定が実行されます。唯一の変数は検索 API プロバイダーです。
破線は、候補応答モデル GPT-5.6 Luna (中) を表し、検索ツールを使用せずに実行され、ベースラインとみなされます。フルハーネス設定については、Search API の方法論を参照してください。
ベンチマーク タスクごとおよび検索クエリごとの候補モデルのコストと検索 API のコスト。
プロバイダー 12 個中 12 個 NEW モデルのみ ($0.0029) 候補モデルのコスト タスクあたりのコスト
1 つのベンチマーク タスクの合計コスト。候補回答モデルのトークン (入力、キャッシュ、推論、出力) と、モデルが実行するために選択した検索に対する Search API プロバイダーの定価に分割されます。モデルのみのベースラインでは検索は実行されないため、そのコストはトークン コストのみです。モデルのコストは、特定のクエリに対して返されるペイロードの違い、実行される検索数の増加 (適切な結果が得られないことによる)、検索結果の品質が低い場合の推論トークンの使用量の増加により、プロバイダーごとに異なる場合があります。
破線は、候補応答モデル GPT-5.6 Luna (中) を表し、検索ツールを使用せずに実行され、ベースラインとみなされます。フルハーネス設定については、Search API の方法論を参照してください。
人工分析検索インデックスとタスクあたりのコスト
プロバイダー 12 社中 12 社 NEW 最も魅力的な象限パレート線 タスクあたりのコスト
1 つのベンチマーク タスクの合計コスト。候補回答モデルのトークン (入力、キャッシュ、推論、出力) と、モデルが実行するために選択した検索に対する Search API プロバイダーの定価に分割されます。モデルのみのベースラインでは検索は実行されないため、そのコストはトークン コストのみです。モデルのコストは、特定のクエリに対して返されるペイロードの違い、実行される検索数の増加（G ではないため）により、プロバイダーごとに異なる場合があります。

適切な結果を設定すること）、検索結果の品質が低い場合には推論トークンの使用量が増加します。
ベンチマーク タスクごとおよび検索クエリごとのモデル時間と検索時間。
プロバイダー 12 個中 12 個 NEW モデルのみ (16.9 秒) モデル時間 タスクごとの時間
1 つのベンチマーク タスクのモデル時間と検索時間の合計。モデル時間は、タスクごとの応答トークンと推論トークンをモデルの正規応答出力速度で割った値で求められます。検索時間は、web_search 呼び出しに費やされた測定時間です。プロバイダーは、呼び出しあたりの速度が速くても、モデルがより頻繁に検索する場合、合計時間がさらに長くなります。
破線は、候補応答モデル GPT-5.6 Luna (中) を表し、検索ツールを使用せずに実行され、ベースラインとみなされます。フルハーネス設定については、Search API の方法論を参照してください。
人工分析検索インデックスとタスクあたりの時間の比較
プロバイダー 12 社中 12 社 NEW 最も魅力的な象限パレート線 タスクあたりの時間
1 つのベンチマーク タスクのモデル時間と検索時間の合計。モデル時間は、タスクごとの応答トークンと推論トークンをモデルの正規応答出力速度で割った値で求められます。検索時間は、web_search 呼び出しに費やされた測定時間です。プロバイダーは、呼び出しあたりの速度が速くても、モデルがより頻繁に検索する場合、合計時間がさらに長くなります。
並べ替え可能なパブリック Search API ベンチマーク行とベンチマークごとの内訳。
12 プロバイダー中 12 個 NEW 13 行 プロバイダー 人工分析 検索インデックス ベースラインリフト DeepSearchQA F1 BrowseComp 精度 Omniscience 精度 検索 $/1k USD / 1k タスク モデル $/1k USD / 1k タスク タスクあたりの時間 秒 並列検索 (上級) 75 42 81 77 67 $47.93 $35.58 38.3s Exa 検索 (自動) 74 41 78 74 70 $65.57 $61.58 28.6s Firecrawl 検索 73 40 74 74 73 $30.48 $44.94 57.8s 並列検索 (基本) 73 40 79 73 68 $45.14 $69.69 23.0s並列検索 (高速) 73 40 8

0 72 68 $8.41 $59.67 16.2秒 Exa検索（高速） 68 35 76 61 69 $78.11 $80.53 23.4秒 You.com検索 68 35 63 74 66 $68.93 $58.28 36.8秒 並列検索（ターボ） 67 34 70 75 56 $13.64 $46.57 21.6s Keenable Search (プロ) 67 34 70 66 65 $23.69 $73.20 25.5s Keenable Search (リアルタイム) 67 34 70 64 66 $26.87 $63.67 17.6s Tavily Search (ベーシック) 66 33 74 59 64 $126 $66.58 39.7 秒 Brave Search 65 32 62 67 65 $71.61 $74.52 29.9 秒 モデルのみ 33 ベースライン 45 17 38 — $2.92 16.9 秒 1,000 タスクあたりのコスト
検索 $/1k およびモデル $/1k は、1,000 のベンチマーク タスクの平均コストを検索トークン コストとモデル トークン コストに分類したものです。モデルのみのベースラインでは検索は実行されません。
1 つのベンチマーク タスクのモデル時間と検索時間の合計。モデル時間は、タスクごとの応答トークンと推論トークンをモデルの正規応答出力速度で割った値で求められます。検索時間は、web_search 呼び出しに費やされた測定時間です。プロバイダーは、呼び出しあたりの速度が速くても、モデルがより頻繁に検索する場合、合計時間がさらに長くなります。
Artificial Analysis Search Index の背後にある各ベンチマークの代表的なタスク例。
多くの検索が必要な広範な研究課題。答えは項目のリストです。 LLM 採点者は、回答項目に対する F1 スコアを使用して各回答を採点します。完全な評価分割には 900 のタスクがあります。
Old School Runescape の Desert Treasure クエストを完了しました。荒野にテレポートするために使用できる、4 つ以上のルーンを必要とする呪文の名前をすべて挙げてもらえますか?
答え: ダリーヤック テレポート、ゴーロック テレポート
マルチホップブラウジングが必要な、見つけにくい事実。完全な評価プールからのハード 200 サンプルのサブセットを使用します。採点者は正確な答えをチェックします。
次の基準を満たす構造物の名前と場所を探しています。
1. オーストラリア東部に位置します。
2. 徒歩で訪問可能。
3.W

2016年に再建されました。
4. 50mを超える。
5. 別の同様の構造からもわかります。
6. 毎年ディナーを主催します。
7. 元々は別のユースケース用に構築されましたが、そのユースケースには何年も役立っていません。
答え: ショーンクリフ桟橋、ブリスベン、オーストラリア
6 つのドメインにわたってバランスがとれた、600 の事実に関する質問のプライベート サブセット。採点者は正確さを採点します。スコアは、検索によってモデルの内部知識がどの程度追加されるかを示します。
カレン・ラッセルの短編小説『オオカミに育てられた少女たちの聖ルーシーの家』で、デビュタント舞踏会に雇われたスリーピースのジャズバンドはどこの出身地ですか?
これらの例は代表的な問題タイプであり、ベンチマークのサンプルではありません。
方法論と詳細情報
どの検索 API が最高の回答品質を持っていますか?
Parallel Search (アドバンスト) は、Artificial Analysis によってベンチマークされた 7 つの検索 API プロバイダー全体で、Artificial Analysis 検索インデックスを 75 でリードしています。
レイテンシが最も低い検索 API はどれですか?
タスクあたりの並列検索 (高速) が最も速く、タスクあたりの平均時間は 16.2 秒 (モデル時間と検索時間の合計) です。個々の検索クエリごとに、Keenable Search (リアルタイム) が最も速く、検索クエリあたりの平均時間は 0.34 秒です。
どの検索 API が最も安いですか?
Parallel Search (高速) の実測検索コストは 1,000 ベンチマーク タスクあたり 8.41 ドルで最も低くなります。
検索を追加すると回答の質はどの程度向上しますか?
検索により、並列検索 (上級) の最も測定された品質が追加され、検索なしの同じモデル (モデルのみのベースライン) と比較して人工分析検索インデックスが 42 上昇しました。
Search API のスコア付けにはどのベンチマークが使用されますか?
Search API の回答品質は、AA-Omniscience、BrowseComp、DeepSearchQA などの公開ベンチマークから集計されます。それぞれ、検索を使用して情報を検索、抽出、または検証するモデルの能力を測定します。
どうやって

Search API プロバイダーを選択すればよいですか?
最適なプロバイダーは、優先順位によって異なります。品質チャートを使用して回答の精度を比較し、コスト チャートを使用して品質と検索およびモデルのコストのバランスをとり、リアルタイム ユースケースのレイテンシ チャートを使用します。トレードオフ散布図は、品質とコスト、および品質とレイテンシのフロンティアにあるプロバイダーを強調表示します。完全な方法論を見る
新しい記事に関する通知を受け取る
X LinkedIn YouTube Rednote Discord © 2026 人工分析

## Original Extract

Compare Search API providers across search quality, benchmark accuracy, latency, and cost.

Artificial Analysis Search Index: Search API Benchmark & Leaderboard | Artificial Analysis Artificial Analysis K Artificial Analysis Models
K Artificial Analysis Search Index: Search API Benchmark & Leaderboard
Search APIs allow AI agents to retrieve information from the web, grounding their responses in up-to-date sources. They are used across a wide range of agentic applications, including deep research, coding, and general knowledge work, to improve factuality and overall outputs. Search APIs differ in several important ways, including the underlying web index they search, how search results are presented to the model, and the tradeoffs they make between cost, speed, and retrieval quality. We benchmark 12 Search API products across 7 providers to help developers make more informed decisions when selecting a search tool for their next AI agent.
For benchmark definitions, scoring, and data handling details, see the methodology page .
Language Model API Endpoints Artificial Analysis Methodology Artificial Analysis Search Index Launch Article Stirrup agent Benchmark tasks Candidate model GPT-5.6 Luna (medium) turn 0 / 25 Search provider web_search native payload, up to 10 results Grading Grader avg F1 0.79 Search API providers Constants Same tasks, same model and grader (GPT-5.6 Luna, medium), same Stirrup harness (web_search · web_fetch · finish, max 25 turns). Only the search provider varies. Restart 0.5× 1× 2× Highlights
Artificial Analysis Search Index
Search-enabled answer quality across public benchmarks.
Artificial Analysis Search Index
12 of 12 providers NEW Model only (33) Artificial Analysis Search Index
The equal-weighted mean of each Search API provider's score on DeepSearchQA (average F1), BrowseComp (exact-answer accuracy), and AA-Omniscience (accuracy), shown on a 0-100 scale. Every result runs the same candidate answer model, harness, and settings. The only variable is the Search API provider.
The dashed line represents the candidate answer model, GPT-5.6 Luna (medium), run with no search tools and is considered the baseline. See the Search API methodology for the full harness settings.
Candidate model cost and Search API cost per benchmark task and per search query.
12 of 12 providers NEW Model only ($0.0029) Candidate model cost Cost per Task
Total cost of one benchmark task, split into the candidate answer model's tokens (input, cached, reasoning, and output) and the Search API provider's list price for the searches the model chose to run. The model-only baseline runs no searches, so its cost is token cost alone. Model cost can differ per provider due to differences in returned payloads for a given query, increased numbers of searches performed (due to not getting the right results) as well as increased reasoning token usage when search results are of lower quality.
The dashed line represents the candidate answer model, GPT-5.6 Luna (medium), run with no search tools and is considered the baseline. See the Search API methodology for the full harness settings.
Artificial Analysis Search Index vs. Cost per Task
12 of 12 providers NEW Most attractive quadrant Pareto line Cost per Task
Total cost of one benchmark task, split into the candidate answer model's tokens (input, cached, reasoning, and output) and the Search API provider's list price for the searches the model chose to run. The model-only baseline runs no searches, so its cost is token cost alone. Model cost can differ per provider due to differences in returned payloads for a given query, increased numbers of searches performed (due to not getting the right results) as well as increased reasoning token usage when search results are of lower quality.
Model time and search time per benchmark task and per search query.
12 of 12 providers NEW Model only (16.9s) Model time Time per Task
The sum of model time and search time for one benchmark task. Model time is derived: answer and reasoning tokens per task divided by the model's canonical answer output speed. Search time is the measured time spent in web_search calls. A provider can be fast per call and still add more total time if the model searches against it more often.
The dashed line represents the candidate answer model, GPT-5.6 Luna (medium), run with no search tools and is considered the baseline. See the Search API methodology for the full harness settings.
Artificial Analysis Search Index vs. Time per Task
12 of 12 providers NEW Most attractive quadrant Pareto line Time per Task
The sum of model time and search time for one benchmark task. Model time is derived: answer and reasoning tokens per task divided by the model's canonical answer output speed. Search time is the measured time spent in web_search calls. A provider can be fast per call and still add more total time if the model searches against it more often.
Sortable public Search API benchmark rows and per benchmark breakdown.
12 of 12 providers NEW 13 rows Provider Artificial Analysis Search Index Baseline Lift DeepSearchQA F1 BrowseComp Accuracy Omniscience Accuracy Search $/1k USD / 1k tasks Model $/1k USD / 1k tasks Time per Task Seconds Parallel Search (advanced) 75 42 81 77 67 $47.93 $35.58 38.3s Exa Search (auto) 74 41 78 74 70 $65.57 $61.58 28.6s Firecrawl Search 73 40 74 74 73 $30.48 $44.94 57.8s Parallel Search (basic) 73 40 79 73 68 $45.14 $69.69 23.0s Parallel Search (fast) 73 40 80 72 68 $8.41 $59.67 16.2s Exa Search (fast) 68 35 76 61 69 $78.11 $80.53 23.4s You.com Search 68 35 63 74 66 $68.93 $58.28 36.8s Parallel Search (turbo) 67 34 70 75 56 $13.64 $46.57 21.6s Keenable Search (pro) 67 34 70 66 65 $23.69 $73.20 25.5s Keenable Search (realtime) 67 34 70 64 66 $26.87 $63.67 17.6s Tavily Search (basic) 66 33 74 59 64 $126 $66.58 39.7s Brave Search 65 32 62 67 65 $71.61 $74.52 29.9s Model only 33 baseline 45 17 38 — $2.92 16.9s Cost per 1,000 Tasks
Search $/1k and Model $/1k are the average cost of 1,000 benchmark tasks broken down into search and model token costs. The model-only baseline runs no searches.
The sum of model time and search time for one benchmark task. Model time is derived: answer and reasoning tokens per task divided by the model's canonical answer output speed. Search time is the measured time spent in web_search calls. A provider can be fast per call and still add more total time if the model searches against it more often.
Representative example tasks for each benchmark behind the Artificial Analysis Search Index.
Broad research questions that need many searches. Answers are lists of items. An LLM grader scores each answer with an F1 score over the answer items. The full eval split has 900 tasks.
I've completed the Desert Treasure quest on Old School Runescape, can you list all the names of the spells I can use to teleport into the wilderness that require more than four runes?
Answer: Dareeyak Teleport, Ghorrock Teleport
Hard-to-find facts that need multi-hop browsing. We use a hard 200-sample subset from the full evaluation pool. The grader checks for an exact answer.
I'm looking for the name and location of a structure which fulfills the following criteria:
1. Located in Eastern Australia.
2. Can be visited on foot.
3. Was re-built in 2016.
4. Is longer than 50m.
5. Can be seen from another similar structure.
6. Hosts a yearly dinner.
7. Was originally built for another use case, but has not served that use case for a number of years.
Answer: Shorncliffe Pier, Brisbane, Australia
A private subset of 600 factual questions, balanced across 6 domains. The grader scores accuracy. The score shows how much search adds to the model's internal knowledge.
In Karen Russell's short story "St. Lucy's Home for Girls Raised by Wolves," from which locality was the three-piece jazz band hired for the Debutante Ball?
The examples are representative problem types, not samples from the benchmarks.
Methodology & Further Information
Which Search API has the highest answer quality?
Parallel Search (advanced) leads the Artificial Analysis Search Index at 75 across 7 Search API providers benchmarked by Artificial Analysis.
Which Search API has the lowest latency?
Per task, Parallel Search (fast) is the fastest, with an average time per task of 16.2s (model time plus search time). Per individual search query, Keenable Search (realtime) is the fastest, with an average time per search query of 0.34s.
Which Search API is the cheapest?
Parallel Search (fast) has the lowest measured search cost at $8.41 per 1,000 benchmark tasks.
How much does adding search improve answer quality?
Search adds the most measured quality for Parallel Search (advanced), lifting the Artificial Analysis Search Index by 42 versus the same model with no search (its model-only baseline).
Which benchmarks are used to score Search APIs?
Search API answer quality is aggregated from public benchmarks including AA-Omniscience, BrowseComp, and DeepSearchQA. Each measures a model's ability to find, extract, or verify information using search.
How do I choose a Search API provider?
The best provider depends on your priorities. Use the quality charts to compare answer accuracy, the cost charts to balance quality against search and model cost, and the latency charts for real-time use cases. The tradeoff scatter plots highlight the providers on the quality-cost and quality-latency frontiers. See the full methodology
Get notified about new articles
X LinkedIn YouTube Rednote Discord © 2026 Artificial Analysis
