---
source: "https://artificialanalysis.ai/articles/gemini-3-7-time-frontier"
hn_url: "https://news.ycombinator.com/item?id=49296625"
title: "Gemini 3.7 Flash: On the Intelligence vs. Time per Task Pareto frontier"
article_title: "Gemini 3.7 Flash: On the Intelligence vs. Time per Task Pareto frontier"
author: "frozenseven"
captured_at: "2026-08-14T09:57:58Z"
capture_tool: "hn-digest"
hn_id: 49296625
score: 1
comments: 0
posted_at: "2026-08-14T09:55:03Z"
tags:
  - hacker-news
  - translated
---

# Gemini 3.7 Flash: On the Intelligence vs. Time per Task Pareto frontier

- HN: [49296625](https://news.ycombinator.com/item?id=49296625)
- Source: [artificialanalysis.ai](https://artificialanalysis.ai/articles/gemini-3-7-time-frontier)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T09:55:03Z

## Translation

タイトル: Gemini 3.7 Flash: インテリジェンスとタスクあたりの時間のパレートフロンティアについて
説明: AI モデルとホスティング プロバイダーの独立した分析。 AI の状況を理解し、ユースケースに最適なモデルと API プロバイダーを選択します。

記事本文:
Gemini 3.7 Flash: インテリジェンスとタスクあたりの時間の関係 パレート フロンティア 人工分析 K 人工分析モデル
K すべての記事 2026 年 8 月 13 日
Gemini 3.7 Flash: インテリジェンスとタスクあたりの時間のパレートフロンティアについて
Google が Gemini 3.7 Flash をリリースし、Gemini 3.6 Flash より 4 ポイント向上し、インテリジェンスとタスクあたりの時間のパレート フロンティアに到達しました
Google DeepMind は、この 3 か月で 3 番目の新しい Gemini Flash モデルをリリースしました。 Gemini 3.7 Flash (高) の人工分析インテリジェンス インデックスのスコアは 56 で、GPT-5.6 Terra (最大、57) および Muse Spark 1.2 (xhigh、57) に次ぐスコアです。
私たちは、リリースに先立って、3 つの推論レベル (高、中、低) すべてにわたって Gemini 3.7 フラッシュのベンチマークを実施しました。高度な推論により、Gemini 3.7 Flash は 3.6 Flash より 4 ポイント向上しており、タスクあたりの平均時間は 1.7 であり、GPT-5.6 Terra (最大) より 40% 高速です。これにより、Gemini 3.7 Flash はインテリジェンスとタスクあたりの時間のパレートの最前線に位置し、Gemini Flash ファミリーのモデル全体での Google の速度重視の姿勢が強化されます。
Gemini 3.7 Flash の 3 つの推論レベルにわたる主要なベンチマーク結果:
➤ インテリジェンス インデックスの 4 ポイントの向上: Gemini 3.7 Flash (高) の人工分析インテリジェンス インデックスのスコアは 56 で、Gemini 3.6 Flash から 4 ポイント増加しました。この改善は主に、Tau3 Banking (+3 ポイント)、terminal-Bench v2.1 (+8 ポイント)、および GDPval-AA v2 (+103 Elo) などのエージェント評価の上昇によるものです。中程度の推論では、Gemini 3.7 Flash のスコアは 53 で、DeepSeek V4 Pro 0813 (最大、53) および GLM-5.2 (最大、53) に匹敵します。単純な推論では、スコアは 51 で、DeepSeek V4 Flash 0731 (最大 52) に次ぐものになります。
➤ インテリジェンスとタスクあたりの時間のパレート フロンティア: Gemini 3.7 フラッシュは、1 秒あたり最大 340 個の出力トークンを生成します。これは、GPT-5.6 Terra および GLM の出力速度のほぼ 3 倍です。

-5.2。高度に推論すると、これはタスクあたりの平均時間 1.7 分に相当し、Gemini 3.7 フラッシュはインテリジェンスとタスクあたりの時間のパレート フロンティアに位置します。
➤ Gemini 3.6 Flash よりもタスクあたりのコストが 30% 低い: Gemini 3.7 Flash は、Gemini 3.6 Flash の標準価格である 100 万入出力トークンあたり 1.50 ドル/7.50 ドルを維持しますが、Google は年末まで 100 万トークンあたり 0.75 ドル/3.75 ドルの割引価格を提供します。この割引価格では、Gemini 3.7 Flash (高) の料金は、Intelligence Index タスクあたり 0.40 ドルで、Gemini 3.6 Flash および同等の Muse Spark 1.2 (xhigh、0.40 ドル) よりも 30% 安くなります。中程度の推論では、タスクあたりのコストは 0.26 ドルに低下し、モデルはインテリジェンスとタスクあたりのコストのパレート フロンティアに位置します。
➤ AutomationBench-AA および AA-AnalystAgent での優れたパフォーマンス: 最近リリースされた、スプレッドシートやドキュメントに関する複雑な質問に答えるモデルの能力を測定するベンチマークである AA-AnalystAgent では、Gemini 3.7 Flash (高) が、Claude Opus 5 (最大 54%) や Fable 5 (49%) を上回る 60% という最高の pass^5 スコアを達成しました。 Gemini 3.7 Flash は、シミュレートされた SaaS 環境におけるエージェント機能のベンチマークである AutomationBench-AA でも 62.7% のスコアでリードしており、Kimi K3 (最大、53%) や GPT-5.6 Sol (最大、51.2%) を上回っています。
➤ エージェントのナレッジワークの改善: Gemini 3.7 Flash は、Gemini 3.6 Flash と比較して、エージェントのベンチマーク全体で改善を示しています。当社独自のエージェント ナレッジ ワーク評価である AA-Briefcase では、Gemini 3.7 Flash (高) は 1132 Elo を達成し、前世代から +169 向上し、Minimax-M3 をわずかに上回りました。同様に、Gemini 3.7 Flash (高) は GDPval-AA v2 で 103 ポイント向上し、スコア 1525 で、GLM 5.2 (最大 1506) や DeepSeek V4 Flash 0731 (最大 1558) と同等です。
➤ コンテキスト ウィンドウ: 100 万トークン、Gemini 3.6 Flash から変更なし
➤ マルチモーダ

性: テキスト、画像、ビデオ、および音声入力とテキスト出力
➤ 価格: 標準価格で 100 万入出力トークンあたり $1.50/$7.50。 Google は年末まで 100 万トークンあたり 0.75 ドル / 3.75 ドルの割引価格を提供しています。キャッシュされた入力トークンは同じ 90% 割引を維持します
さらなる分析については、https://artificialanalysis.ai/models/gemini-3-7-flash を参照してください。
Optima の発表: ユースケースに合わせたカスタム ベンチマークを作成
Optima は、独自のワークロードでモデルをベンチマークするための新しいプラットフォームです。独自のファイル、エージェント トレース、またはコーディング環境からベンチマークを構築し、ワンクリックで主要なモデル間でベンチマークを実行し、タスクあたりのコストとタスクあたりの時間とともに品質を比較します。
Upstage Solar Pro 4: ベンチマークと分析
UpstageがSolar Pro 4をリリース
Grok 4.6 は SpaceXAI をインテリジェンスのフロンティアに戻し、コスト効率をリードします
SpaceXAI は、エージェントのパフォーマンスとコスト効率の強みを活かしてインテリジェンスのフロンティアに戻ります
新しい記事に関する通知を受け取る
X LinkedIn YouTube Rednote Discord © 2026 人工分析

## Original Extract

Independent analysis of AI models and hosting providers. Understand the AI landscape and choose the best model and API provider for your use-case.

Gemini 3.7 Flash: On the Intelligence vs. Time per Task Pareto frontier Artificial Analysis K Artificial Analysis Models
K All articles August 13, 2026
Gemini 3.7 Flash: On the Intelligence vs. Time per Task Pareto frontier
Google has released Gemini 3.7 Flash, improving 4 points over Gemini 3.6 Flash and reaching the Intelligence vs. Time per Task Pareto frontier
Google DeepMind has released its third new Gemini Flash model in three months. Gemini 3.7 Flash (high) scores 56 on the Artificial Analysis Intelligence Index, just behind GPT-5.6 Terra (max, 57) and Muse Spark 1.2 (xhigh, 57)
We benchmarked Gemini 3.7 Flash across all three reasoning levels (high, medium, low) ahead of release. With high reasoning, Gemini 3.7 Flash is a 4 point improvement over 3.6 Flash, while achieving an average Time per Task of 1.7, 40% faster than GPT-5.6 Terra (max). This places Gemini 3.7 Flash on the Intelligence vs. Time per Task Pareto frontier, reinforcing Google’s focus on speed across the Gemini Flash family of models
Key benchmarking results across Gemini 3.7 Flash’s three reasoning levels:
➤ 4 point Intelligence Index improvement: Gemini 3.7 Flash (high) scores 56 on the Artificial Analysis Intelligence Index, up 4 points from Gemini 3.6 Flash. The improvement is driven primarily by gains on agentic evaluations, including Tau3 Banking (+3 points), Terminal-Bench v2.1 (+8 points), and GDPval-AA v2 (+103 Elo). With medium reasoning, Gemini 3.7 Flash scores 53, matching DeepSeek V4 Pro 0813 (max, 53) and GLM-5.2 (max, 53). With low reasoning, it scores 51, just behind DeepSeek V4 Flash 0731 (max, 52)
➤ Pareto frontier on Intelligence vs. Time per Task: Gemini 3.7 Flash produces ~340 output tokens per second, nearly 3x the output speed of GPT-5.6 Terra and GLM-5.2. With high reasoning, this translates to an average Time per Task of 1.7 minutes, placing Gemini 3.7 Flash on the Intelligence vs. Time per Task Pareto frontier
➤ 30% lower Cost per Task than Gemini 3.6 Flash: Gemini 3.7 Flash retains Gemini 3.6 Flash’s standard pricing of $1.50/$7.50 per 1M input/output tokens, however Google is offering discounted pricing through the end of the year at $0.75/$3.75 per 1M tokens. At this discounted price, Gemini 3.7 Flash (high) costs $0.40 per Intelligence Index task, 30% less than Gemini 3.6 Flash and matching Muse Spark 1.2 (xhigh, $0.40). With medium reasoning, Cost per Task falls to $0.26, placing the model on the Intelligence vs. Cost per Task Pareto frontier
➤ Leading performance on AutomationBench-AA and AA-AnalystAgent: On AA-AnalystAgent, our recently released benchmark measuring models’ ability to answer complex questions about spreadsheets and documents, Gemini 3.7 Flash (high) achieves the highest pass^5 score at 60%, ahead of Claude Opus 5 (max, 54%) and Fable 5 (49%). Gemini 3.7 Flash also leads AutomationBench-AA, our benchmark of agentic capabilities in simulated SaaS environments, with a score of 62.7%, ahead of Kimi K3 (max, 53%) and GPT-5.6 Sol (max, 51.2%)
➤ Agentic knowledge work improvements: Gemini 3.7 Flash shows improvement across agentic benchmarks compared to Gemini 3.6 Flash. In AA-Briefcase, our proprietary agentic knowledge work evaluation, Gemini 3.7 Flash (high) achieves a 1132 Elo, a +169 improvement from its predecessor, putting it just above Minimax-M3. Similarly, Gemini 3.7 Flash (high) improves by 103 points on GDPval-AA v2, scoring 1525, on par with GLM 5.2 (max, 1506) and DeepSeek V4 Flash 0731 (max, 1558)
➤ Context Window: 1M tokens, unchanged from Gemini 3.6 Flash
➤ Multimodality: Text, image, video, and speech input, with text output
➤ Pricing: $1.50/$7.50 per 1M input/output tokens at standard pricing. Google is offering discounted pricing of $0.75/$3.75 per 1M tokens through the end of the year. Cached input tokens retain the same 90% discount
For further analysis, see https://artificialanalysis.ai/models/gemini-3-7-flash
Announcing Optima: create a custom benchmark for your use case
Optima is a new platform for benchmarking models on your own workloads. Build a benchmark from your own files, agent traces or coding environment, run it across leading models in a single click, and compare quality alongside cost per task and time per task.
Upstage Solar Pro 4: Benchmarks and analysis
Upstage has released Solar Pro 4
Grok 4.6 returns SpaceXAI to the intelligence frontier and leads on cost efficiency
SpaceXAI returns to the intelligence frontier with strengths in agentic performance and cost efficiency
Get notified about new articles
X LinkedIn YouTube Rednote Discord © 2026 Artificial Analysis
