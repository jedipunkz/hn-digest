---
source: "https://artificialanalysis.ai/articles/claude-sonnet-5-agentic-cost"
hn_url: "https://news.ycombinator.com/item?id=48740633"
title: "Claude Sonnet 5: strong agentic performance at a higher cost per task"
article_title: "Claude Sonnet 5: strong agentic performance at a higher cost per task"
author: "himata4113"
captured_at: "2026-07-01T00:31:45Z"
capture_tool: "hn-digest"
hn_id: 48740633
score: 2
comments: 0
posted_at: "2026-06-30T23:35:07Z"
tags:
  - hacker-news
  - translated
---

# Claude Sonnet 5: strong agentic performance at a higher cost per task

- HN: [48740633](https://news.ycombinator.com/item?id=48740633)
- Source: [artificialanalysis.ai](https://artificialanalysis.ai/articles/claude-sonnet-5-agentic-cost)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T23:35:07Z

## Translation

タイトル: クロード ソネット 5: タスクあたりのコストが高くても強力なエージェント パフォーマンスを実現
説明: AI モデルとホスティング プロバイダーの独立した分析。 AI の状況を理解し、ユースケースに最適なモデルと API プロバイダーを選択します。

記事本文:
Claude Sonnet 5: タスクあたりのコストが高くても強力なエージェントのパフォーマンス 人工分析 人工分析モデル
Claude Sonnet 5: タスクあたりのコストが高くても強力なエージェント パフォーマンスを実現
Claude Sonnet 5 は人工分析インテリジェンス インデックスで 53 を達成していますが、プロモーション価格がなければ、タスクあたりのコストが Opus 4.8 より高くなります。
私たちは Anthropic がリリースに先立って Claude Sonnet 5 を評価することをサポートしました。最大限の努力で Sonnet 4.6 より 6 ポイント向上し、高度な推論で GPT-5.5 と同じインテリジェンス インデックスを達成しましたが、Opus 4.7 および 4.8 には及ばないままです。
➤ Claude Sonnet 5 は、人工分析インテリジェンス指数で第 5 位のモデルであり、GPT-5.5 (xhigh) および Opus 4.8 (max) にわずか 2 ～ 3 ポイントの差をつけています。
➤ 最大の努力により、Sonnet 5 は以前の Anthropic モデルよりもより強力に動作します。Sonnet 4.6 と比較して、インテリジェンス インデックス タスクあたりの出力トークンの使用量が最大 40% 増加し、ナレッジ ワーク評価 AA-Briefcase および GDPval-AA のエージェント ターンが最大 3 倍になりました。この動作は「努力」設定に合わせて調整され、最大の努力では GDPval-AA の低努力の約 6 倍のターンが使用されます。
➤ プロモーション価格を考慮する前に、Claude Sonnet 5 のタスクあたりのコストは Opus 4.8 よりも高い : Claude Sonnet 5 のインテリジェンス インデックスのタスクあたりのコストは 2.29 ドルで、Sonnet 4.6 と比較して約 2 倍、Claude Opus 4.8 よりも約 15% 高くなります。これはもっぱらトークンの使用量の増加が原因です。 Sonnet 5 は Sonnet 4.6 と同じ 1M 入出力トークンあたり 3 ドル/15 ドルの価格を維持しますが (Opus 4.8 の 5 ドル/25 ドルと比較)、Anthropic は 9 月 1 日まで 3 分の 1 の値下げを 2 ドル/10 ドルに提供しています。私たちの結果では標準の 3 ドル/15 ドルの価格設定を使用しています。
➤ Sonnet 5 は、エージェントのナレッジ ワーク タスクにおいて Opus 4.8 と同等またはそれを上回っています。AA-Briefcase と GDPval-AA の両方で、Claude Sonnet 5 は Opus 4.8 のすぐ上に位置し、Claude Fable のみが後続しています。

5 (現在は一般提供されていません)。これらのベンチマークは、オープンソースのリファレンス エージェント ハーネスである Stirrup を使用して、正確で適切に表現されたプロフェッショナルな出力を生成するモデルの能力をテストします。
➤ 推論と知識を必要とするタスクに関しては、Sonnet は依然として上位の兄弟に後れを取っています。多くの評価で大幅な向上が見られたにもかかわらず、重い推論と知識のベンチマークでは、依然として Sonnet 5 よりも Opus 4.8 が優れています。Argonne と UIUC の研究者によって開発されたフロンティア物理推論ベンチマークである CritPt では、Sonnet 5 のスコアは 17% です。これは、前任者より 14 ポイント高いですが、GLM-5.2 には及ばない、Claude 氏Opus と Fable、および GPT-5.5 (xhigh および Pro)
➤ Sonnet 5 は、ターミナルベンチ v2.1 (+9 ポイント)、Humanity’s Last Exam (+10 ポイント)、および SciCode (+7 ポイント) において Sonnet 4.6 に比べて大幅な改善を示しましたが、その他のスコアは比較的横ばいでした
➤ 100 万トークンのコンテキスト ウィンドウ (Sonnet 4.6 に相当)
➤ 入力/出力の 100 万トークンあたり $3/$15 の価格 (9 月 1 日まで $2/$10 に値下げ)。キャッシュの価格は、存続時間 5 分のキャッシュ書き込みの場合は 25% のプレミアム (100 万トークンあたり 3.75 ドル)、キャッシュ ヒットの場合は 90% 割引 (100 万トークンあたり 0.3 ドル) のままです。
➤ モデルのパフォーマンスとレイテンシを設定する場合は、引き続きエフォートが推奨される方法です。 Sonnet 5 は、Sonnet 4.6 に関連する追加の「xhigh」努力設定を追加し、Opus 4.8 で利用可能な 5 つの努力レベル (最大、xhigh、高、中、低) に一致します。
Claude Sonnet 5 を他の主要モデルと比較するには: https://artificialanalysis.ai/models/claude-sonnet-5
AA-Briefcase でのタスクごとの時間を測定する
新しいベンチマークである AA-Briefcase で測定したように、エージェントのナレッジ ワークには、フロンティア モデルのタスクごとに 20 分以上かかる場合があります。
人工分析 Speech to Speech Index の発表
人工分析の発表

Speech to Speech Index、ネイティブ Speech to Speech モデル品質の新しい合成メトリック。Big Bench Audio、Full Duplex Bench、𝜏-Voice で構成されます。
AA-ブリーフケースの発表: 最先端の知識労働評価
AA-Briefcase は、業界の専門家によって構築された複雑なプロジェクトにおける現実的なナレッジ ワーク タスクのモデルをテストするための新しいベンチマークです。モデルは、複数週間にわたるナレッジ ワーク プロジェクトで評価されます。各プロジェクトには、リンクされた多くのタスクと数千の入力ソース ファイルがあり、ルーブリックとペアごとの評価を組み合わせて、検証可能なタスクの成功、分析の品質、およびプレゼンテーションの品質を評価します。
新しい記事に関する通知を受け取る
X LinkedIn YouTube Rednote Discord © 2026 人工分析

## Original Extract

Independent analysis of AI models and hosting providers. Understand the AI landscape and choose the best model and API provider for your use-case.

Claude Sonnet 5: strong agentic performance at a higher cost per task Artificial Analysis Artificial Analysis Models
Claude Sonnet 5: strong agentic performance at a higher cost per task
Claude Sonnet 5 achieves 53 on the Artificial Analysis Intelligence Index, but without promotional pricing will cost more per task than Opus 4.8
We supported Anthropic to evaluate Claude Sonnet 5 ahead of release: with max effort it improves 6 points over Sonnet 4.6 to achieve the same Intelligence Index as GPT-5.5 with high reasoning, but remains behind Opus 4.7 and 4.8
➤ Claude Sonnet 5 is the #5 model on the Artificial Analysis Intelligence Index , only 2-3 points behind GPT-5.5 (xhigh) and Opus 4.8 (max)
➤ With max effort, Sonnet 5 works harder than previous Anthropic models: it used ~40% more output tokens per Intelligence Index task than Sonnet 4.6, and ~3x the agentic turns for our knowledge work evaluations AA-Briefcase and GDPval-AA. This behavior scales well with the ‘effort’ setting, with the max effort using around 6x more turns than low effort on GDPval-AA
➤ Claude Sonnet 5 costs more per task than Opus 4.8 before accounting for promotional pricing : Claude Sonnet 5 costs $2.29 per task on the Intelligence Index, a ~2x increase compared to Sonnet 4.6 and ~15% more than Claude Opus 4.8. This is driven entirely by increased token usage. Sonnet 5 retains the same $3/$15 per 1M input/output token pricing as Sonnet 4.6 (compared to $5/$25 for Opus 4.8), however Anthropic is offering a one-third reduction to $2/$10 until September 1. Our results use standard $3/$15 pricing
➤ Sonnet 5 matches or outperforms Opus 4.8 on agentic knowledge work tasks: on both AA-Briefcase and GDPval-AA, Claude Sonnet 5 sits just ahead of Opus 4.8, trailing only Claude Fable 5 (which is not currently generally available). These benchmarks test the ability of models to produce accurate and well-presented professional outputs using our open source reference agent harness, Stirrup
➤ For reasoning and knowledge-heavy tasks, Sonnet still sits behind its larger siblings: despite substantial gains across many evaluations, heavy reasoning and knowledge benchmarks still show Opus 4.8 ahead of Sonnet 5. On CritPt, a frontier physics reasoning benchmark developed by researchers at Argonne and UIUC, Sonnet 5 scores 17% - this is 14 points higher than its predecessor, but behind GLM-5.2, Claude Opus and Fable, and GPT-5.5 (xhigh and Pro)
➤ Sonnet 5 also showed significant improvements over Sonnet 4.6 on Terminal-Bench v2.1 (+9 points), Humanity’s Last Exam (+10 points), and SciCode (+7 points), with relatively flat scores elsewhere
➤ Context window of 1 million tokens (equivalent to Sonnet 4.6)
➤ Pricing of $3/$15 per 1M tokens of input/output (reduced to $2/$10 until September 1); cache pricing remains at a 25% premium for cache writes ($3.75 per million tokens) with 5-minute time to live, and 90% discount for cache hits ($0.3 per million tokens)
➤ Effort remains the recommended way of configuring model performance and latency. Sonnet 5 adds an additional ‘xhigh’ effort setting relative to Sonnet 4.6, matching the 5 effort levels available on Opus 4.8 (max, xhigh, high, medium, low)
Compare Claude Sonnet 5 with other leading models at: https://artificialanalysis.ai/models/claude-sonnet-5
Measuring time per task in AA-Briefcase
Agentic knowledge work can take frontier models over 20 minutes per task, as measured in AA-Briefcase, our new benchmark
Announcing the Artificial Analysis Speech to Speech Index
Announcing the Artificial Analysis Speech to Speech Index, our new synthesis metric for native Speech to Speech model quality, comprising of Big Bench Audio, Full Duplex Bench, and 𝜏-Voice
Announcing AA-Briefcase: a frontier knowledge work evaluation
AA-Briefcase is a new benchmark for testing models on realistic knowledge work tasks in complex projects built by industry experts. Models are evaluated on multi-week knowledge work projects, each with many linked tasks and thousands of input source files, combining rubric and pairwise grading to evaluate verifiable task success, analytical quality, and presentation quality.
Get notified about new articles
X LinkedIn YouTube Rednote Discord © 2026 Artificial Analysis
