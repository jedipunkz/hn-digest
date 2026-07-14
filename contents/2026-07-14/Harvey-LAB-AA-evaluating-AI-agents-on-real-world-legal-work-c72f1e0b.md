---
source: "https://artificialanalysis.ai/articles/harvey-lab-aa"
hn_url: "https://news.ycombinator.com/item?id=48902852"
title: "Harvey LAB-AA: evaluating AI agents on real-world legal work"
article_title: "Announcing Harvey LAB-AA: evaluating AI agents on real-world legal work | Artificial Analysis"
author: "theanonymousone"
captured_at: "2026-07-14T07:04:50Z"
capture_tool: "hn-digest"
hn_id: 48902852
score: 1
comments: 0
posted_at: "2026-07-14T06:09:42Z"
tags:
  - hacker-news
  - translated
---

# Harvey LAB-AA: evaluating AI agents on real-world legal work

- HN: [48902852](https://news.ycombinator.com/item?id=48902852)
- Source: [artificialanalysis.ai](https://artificialanalysis.ai/articles/harvey-lab-aa)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T06:09:42Z

## Translation

タイトル: Harvey LAB-AA: 現実世界の法律業務における AI エージェントの評価
記事のタイトル: Harvey LAB-AA の発表: 現実世界の法律業務における AI エージェントの評価 |人工分析
説明: Harvey's Legal Agent Benchmark の実装では、24 の実務分野にまたがる 120 の民間法律タスクの言語モデルをテストし、バイナリ基準のルーブリックに照らして採点します。

記事本文:
Harvey LAB-AA を発表: 現実世界の法律業務における AI エージェントの評価 |人工解析 人工解析 K 人工解析モデル
K すべての記事 2026 年 7 月 7 日 Harvey LAB-AA の発表: 現実世界の法律業務における AI エージェントの評価
Harvey LAB-AA (Legal Agent Benchmark) は、Harvey の新しい代理人法務ベンチマークの実装であり、24 の実務分野にわたる実際の法律業務の言語モデルを評価します。
モデルは、Harvey のチームによって構築された 120 の法的タスクのプライベート セットでテストされ、企業の M&A や資本市場から税務、訴訟、破産に至るまでの実務分野に及びます。モデルは、各タスクで指定された正当な出力を作成するように機能し、各タスクは 2 項基準のルーブリックに基づいて採点されます。私たちが提示する主な指標は全合格率です。これは、ルーブリック内のすべての基準が満たされたタスクの割合であり、現実世界の専門的な法的成果物の高水準を反映しています。
推論モデルは電球アイコンで示されます。 Claude Fable 5 (最大、Opus 4.8 フォールバックあり) は、1 つのタスクのみで Claude Opus 4.8 にフォールバックした後、14.2% の全パス率で Harvey LAB-AA をリードしています。これは、次に優れたモデルである Claude Opus 4.8 (最大) と GLM-5.2 (最大) の 7.5% と比べてほぼ 2 倍であり、同率 7.5%、MiniMax-M3 の 6.7%、Claude Sonnet 5 の 5.0% が続きます。
最前線の法務作業は解決にはほど遠いです。ほとんどのモデルは個々のルーブリック基準の大部分に合格していますが、ごく少数のタスクの要件を完全に満たしています。最良のモデルでも、専門的な法的成果物の約 86% が未完成のままであり、立ち上げ時に評価された 28 モデルのうち 13 モデルはタスクを完全にパスしませんでした。基準合格率で 90% を超えるスコアを獲得したのは、Claude Fable 5 (93.6%)、Claude Opus 4.8 (91.1%)、GLM-5.2 (91.0%)、および Claude Sonnet 5 (90.1%) の 4 つのモデルのみです。
最新の結果については、Harvey LAB-AA ev を参照してください。

評価ページへグラフは 2026 年 7 月 7 日時点のデータを示しています。
回答推論キャッシュ 書き込みキャッシュ ヒット入力 推論モデルは電球アイコンで示されます タスクごとの評価コスト
評価におけるタスクあたりの平均コスト。コストは、正規のトークン数が利用可能な場合、入力、キャッシュ ヒット、キャッシュ書き込み、推論、および応答トークンの価格設定によって分割されます。
Harvey LAB-AA: タスクごとの出力トークン
回答推論 推論モデルは電球アイコンで示されます タスクごとの評価出力トークン
この評価でベンチマーク タスクごとに生成される回答トークンと推論トークンの平均数。
推論モデルは電球アイコンで示されます タスクごとの評価時間
評価タスクごとの加重平均時間 (秒)。これは、タスクごとの出力トークンを出力速度で除算し、評価における各ベンチマークの相対的な重みで重み付けして計算されます。
Harvey LAB-AA ベンチマーク リーダーボード: タスクあたりの平均ターン数
推論モデルは電球アイコンで示されます。
このグラフは、エージェントがタスクごとに要する平均ターン数を示しています。これは、エージェントがベンチマーク タスクを完了するために使用するアクション、ツール呼び出し、および反復サイクルの数の大まかな代理です。
Harvey LAB-AA: 全パス率とリリース日の関係
最も魅力的な地域 OpenAI Google Anthropic Mistral DeepSeek xAI MiniMax NVIDIA Kim StepFun Xiaomi Z AI Alibaba タスクと提出の例
パブリック タスク セットから代表的な Harvey LAB タスク、各モデルに提供された参照ファイル、および生成された成果物を参照します。
添付の買収データ ルーム契約書と支配権の変更および譲渡条項に関する内部メモを確認し、包括的な取引チームのレポートを作成します。
出力: coc-analysis-report.docx
モデルが生成する必要がある期待される出力
coc-analysis-report.docx 包括的な取引

ターゲットの重要な契約全体にわたる支配権の変更と譲渡条項を分析したチームレポート。
apex-distribution-agreement.docx apex-msa.docx を開く deal-overview-memo.docx を開く deal-summary-memo.docx を開く first-continental-credit-agreement.docx を開く great-lakes-credit-agreement.docx を開く meridian-dpa.docx を開く meridian-license-agreement.docx を開く novabridge-partnership.docx を開くnovabridge-supply-agreement.docx orion-subscription-renewal.docx を開く pinnacle-ecommerce-agreement.docx を開く pinnacle-license.docx を開く ridgeline-10k-excerpt.docx を開く Solara-deferred-comp-plan.docx を開く terranode-isa.docx を開く terraverde-lease-agreement.docx を開くwebb-employment-agreement.docx オープンwellstone-manufacturing-agreement.docx オープンモデルの提出
各モデルによって生成される成果物
Claude Fable 5 (フォールバックあり) - coc-analysis-report.docx を開く Harvey LAB-AA と Harvey's LAB の違い
Harvey LAB-AA は Harvey の評価を独自に再実装したもので、元のバージョンとはいくつかの重要な違いがあります。
モデルはスターラップ エージェント ハーネス上で実行され、簡略化された Artificial Analysis で作成されたエージェントとジャッジのプロンプトにより、コンテキストの制限に達したときに失敗するのではなくコンテキストの圧縮などの機能を有効にします。
Harvey のカスタム ツールやドキュメント生成スキル スクリプト (pptx、docx など) は含まれていません。代わりに、生のモデルの機能を反映するシンプルなコード実行ツールが提供されます。
成果物は、モデルが間違ったファイル名を生成する場合のあいまい一致ではなく、指定されたファイル名と正確に一致する必要があります。
採点には 1 人の Gemini 3.1 Pro 審査員が使用され、フロンティア パネルに対して適切に調整されるようにテストされています。
リーダーボードと完全な結果は Harvey LAB-AA 評価ページに掲載されており、新しいモデルがリリースされると更新されます
方法論のページには完全な手順が記載されています

エージェントと採点プロンプトを含む実装
Harvey の最初の LAB 発表では、ベンチマークとその設計が紹介されています
代表的なタスクの公開セットは GitHub で入手できます
Harvey LAB-AA は、オープンソース エージェント フレームワークである Stirrup 上で実行されます
GPT-5.6 Sol、Terra、Luna のインテリジェンスとコストの比較
GPT-5.6 Sol と Luna は、インテリジェンスとタスクあたりのコストのグラフのあらゆる点で Terra よりも優れています。 GPT-5.6 Luna は、特にコスト効率の高いモデルとして際立っています。
Muse Spark 1.1: メタは 3 か月で 8 Intelligence Index ポイントを獲得
Meta の Muse Spark 1.1 は、人工分析インテリジェンス インデックスで 51 のスコアを獲得し、同等の製品と比較してコストとトークン効率が優れています。
インテリジェンス、スピード、コストにわたる GPT-5.6 ベンチマーク
GPT-5.6 Sol は、コストの 3 分の 1 で人工分析インテリジェンス インデックスで Claude Fable 5 に迫る 2 位となり、OpenAI の Codex ハーネスの人工分析コーディング エージェント インデックスで首位に立っています。
新しい記事に関する通知を受け取る
X LinkedIn YouTube Rednote Discord © 2026 人工分析

## Original Extract

Our implementation of Harvey's Legal Agent Benchmark tests language models on 120 private legal tasks spanning 24 practice areas, graded against rubrics of binary criteria.

Announcing Harvey LAB-AA: evaluating AI agents on real-world legal work | Artificial Analysis Artificial Analysis K Artificial Analysis Models
K All articles July 7, 2026 Announcing Harvey LAB-AA: evaluating AI agents on real-world legal work
Harvey LAB-AA (Legal Agent Benchmark) is our implementation of Harvey's new agentic legal benchmark, evaluating language models on real-world legal work across 24 practice areas.
Models are tested on a private set of 120 legal tasks built by the team at Harvey, spanning practice areas from corporate M&A and capital markets to tax, litigation, and bankruptcy. Models work to create the legal outputs specified in each task, and each task is graded against a rubric of binary criteria. The primary metric we present is the all-pass rate: the share of tasks where all criteria in the rubric are satisfied, reflecting the high standard of real-world professional legal deliverables.
Reasoning models are indicated by a lightbulb icon Claude Fable 5 (max, with Opus 4.8 fallback) leads Harvey LAB-AA with a 14.2% all-pass rate, after falling back to Claude Opus 4.8 on only one task. This is almost double the next best models, Claude Opus 4.8 (max) and GLM-5.2 (max), which tie at 7.5%, followed by MiniMax-M3 at 6.7% and Claude Sonnet 5 at 5.0%.
Frontier legal work is far from solved: most models pass a majority of individual rubric criteria but fully satisfy the requirements of very few tasks. The best model still leaves ~86% of professional legal deliverables incomplete, and 13 of the 28 models evaluated at launch fully pass zero tasks. Only four models score above 90% on criterion pass rate: Claude Fable 5 (93.6%), Claude Opus 4.8 (91.1%), GLM-5.2 (91.0%), and Claude Sonnet 5 (90.1%).
For up-to-date results see the Harvey LAB-AA evaluation page . Charts show data as at 7 July 2026 .
Answer Reasoning Cache Write Cache Hit Input Reasoning models are indicated by a lightbulb icon Evaluation Cost per Task
Average cost per task in the evaluation. Costs are split by input, cache hit, cache write, reasoning, and answer token pricing where canonical token counts are available.
Harvey LAB-AA: Output Tokens per Task
Answer Reasoning Reasoning models are indicated by a lightbulb icon Evaluation Output Tokens per Task
The average number of answer and reasoning tokens produced per benchmark task in this evaluation.
Reasoning models are indicated by a lightbulb icon Evaluation Time per Task
The weighted average time (seconds) per evaluation task. This is calculated by dividing output tokens per task by output speed, weighted by the relative weights of each benchmark in the evaluation.
Harvey LAB-AA Benchmark Leaderboard: Average Turns per Task
Reasoning models are indicated by a lightbulb icon What Turns Is Measuring
This chart shows the average number of turns the agent takes per task. It is a rough proxy for how many actions, tool calls, and iteration cycles an agent is using to complete benchmark tasks.
Harvey LAB-AA: All-pass Rate vs. Release Date
Most attractive region OpenAI Google Anthropic Mistral DeepSeek xAI MiniMax NVIDIA Kimi StepFun Xiaomi Z AI Alibaba Example Tasks & Submissions
Browse representative Harvey LAB tasks from the public task set, the reference files each model was given, and the deliverables it produced.
Review the attached acquisition data room contracts and internal memo for change of control and assignment provisions, and prepare a comprehensive deal team report.
Output: coc-analysis-report.docx
Expected outputs the model must produce
coc-analysis-report.docx A comprehensive deal team report analyzing change of control and assignment provisions across the target’s material contracts.
apex-distribution-agreement.docx Open apex-msa.docx Open deal-overview-memo.docx Open deal-summary-memo.docx Open first-continental-credit-agreement.docx Open great-lakes-credit-agreement.docx Open meridian-dpa.docx Open meridian-license-agreement.docx Open novabridge-partnership.docx Open novabridge-supply-agreement.docx Open orion-subscription-renewal.docx Open pinnacle-ecommerce-agreement.docx Open pinnacle-license.docx Open ridgeline-10k-excerpt.docx Open solara-deferred-comp-plan.docx Open terranode-isa.docx Open terraverde-lease-agreement.docx Open webb-employment-agreement.docx Open wellstone-manufacturing-agreement.docx Open Model submissions
Deliverables produced by each model
Claude Fable 5 (with fallback) - coc-analysis-report.docx Open How Harvey LAB-AA differs from Harvey's LAB
Harvey LAB-AA is our independent reimplementation of Harvey's evaluation, and there are several key differences to the original version:
Models are run on our Stirrup agent harness, enabling features such as context compaction rather than failure when reaching context limits, with simplified Artificial Analysis-authored agent and judge prompts
We do not include Harvey's custom tools and document-generation skill scripts (e.g. pptx, docx), instead providing a simple code execution tool to reflect raw model ability
Deliverables must match the exact filename specified, rather than fuzzy matching when models produce incorrect filenames
Grading uses a single Gemini 3.1 Pro judge, tested to be well-calibrated against a frontier panel
The leaderboard and full results live on the Harvey LAB-AA evaluation page , updated as new models are released
The methodology page documents the full implementation, including the agent and grading prompts
Harvey's original LAB announcement introduces the benchmark and its design
A public set of representative tasks is available on GitHub
Harvey LAB-AA runs on Stirrup , our open-source agent framework
How GPT-5.6 Sol, Terra, Luna compare on intelligence vs cost
GPT-5.6 Sol and Luna are ahead of Terra at every point on the Intelligence vs Cost per Task chart. GPT-5.6 Luna stands out as a particularly cost efficient model
Muse Spark 1.1: Meta gains 8 Intelligence Index points in three months
Meta's Muse Spark 1.1 scores 51 on the Artificial Analysis Intelligence Index and is cost and token efficient compared to its peers
GPT-5.6 benchmarks across Intelligence, Speed and Cost
GPT-5.6 Sol comes close second to Claude Fable 5 in the Artificial Analysis Intelligence Index at one third of the cost, and leads the Artificial Analysis Coding Agent Index in OpenAI’s Codex harness
Get notified about new articles
X LinkedIn YouTube Rednote Discord © 2026 Artificial Analysis
