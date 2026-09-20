---
source: "https://quicqdev.github.io/Jev-vs-ML/"
hn_url: "https://news.ycombinator.com/item?id=49774364"
title: "Jev vs. classical ML. Strong on sentiment: Mixed across tasks"
article_title: "Jev vs. classical ML — An eight-dataset benchmark"
image: "https://quicqdev.github.io/Jev-vs-ML/assets/benchmark-release-blue.png"
author: "theanonymousone"
captured_at: "2026-09-20T10:45:31Z"
capture_tool: "hn-digest"
hn_id: 49774364
score: 2
comments: 0
posted_at: "2026-09-20T10:14:31Z"
tags:
  - hacker-news
---

# Jev vs. classical ML. Strong on sentiment: Mixed across tasks

- HN: [49774364](https://news.ycombinator.com/item?id=49774364)
- Source: [quicqdev.github.io](https://quicqdev.github.io/Jev-vs-ML/)
- Score: 2
- Comments: 0
- Posted: 2026-09-20T10:14:31Z

## Translation

Title: Jev vs. classical ML. Strong on sentiment: Mixed across tasks
Article title: Jev vs. classical ML — An eight-dataset benchmark
Description: An independent comparison of Jev 1.13.0 and 11 classical classification pipelines. Explore the completed V3 results, methodology, and limitations.

Article text:
Skip to results
QuicqDev Research notes Results Analysis Method GitHub ↗
Independent benchmark Report 03 / September 2026
Jev vs.
classical ML.
Strong on sentiment.
Mixed across tasks.
We compared Jev 1.13.0 with eleven classical classification pipelines across eight datasets. Its largest advantage was on IMDb. Across the rest of the suite, the picture was more nuanced.
Protocol 3.0.1 · Bounded-budget comparison · Balanced accuracy · Not affiliated with OpenAI or Anthropic
Every model.
Both decision rules.
Switch between raw and threshold-adjusted decisions. Jev’s columns appear first. Each cell shows mean ± sample standard deviation across three training seeds on the same test cases. Bold blue cells mark the best displayed mean in each row, including ties.
Raw: default decision rules after classical model selection. Jev zero-shot uses no task examples or policy threshold fitting.
Scroll horizontally to inspect all columns. Sample SD is not a confidence interval. Zero SD on cached zero-shot predictions is not independent evidence of repeatability.
A different story
for each task.
Compare Jev with the highest-scoring classical pipeline on each dataset. The controls in the full table above also update this chart. Compare raw decisions with binary thresholds learned from separate labeled data.
Figure 1. Mean balanced accuracy, with a common 0–100% scale. Values are rounded to one decimal. The classical comparator is the highest test mean among the eleven pipelines in that panel, selected retrospectively. It is not a deployment selection rule. Read uncertainty notes ↓
Raw zero-shot Jev leads the best classical mean on IMDb and SMS Spam. The largest lead is IMDb: +7.9 percentage points. These are descriptive differences, not statistical-significance claims.
The headline is only
part of the result.
On IMDb, raw zero-shot Jev reaches 96.3% , compared with 88.4% for logistic regression. The +7.9-point advantage remains almost unchanged after threshold adjustment: 96.1% versus 88.3%.
This is the strongest descriptive evidence in Jev’s favor here. The experiment compares a pretrained API model against these classical text pipelines; it does not establish superiority over other language models or stronger modern text representations.
Thresholds change the SMS story.
With raw decisions, Jev zero-shot leads Naive Bayes 96.1% to 95.0% . After policy threshold selection, the comparison becomes 95.9% to 96.3% .
The apparent lead becomes a small deficit. A model’s default decision rule and its ability to separate classes are not the same question. Publish both panels; the adjusted Jev result uses labeled policy data.
Business tabular tasks remain difficult.
On Bank Marketing, adjusted zero-shot-prompt Jev reaches 59.7% , against the voting ensemble’s 73.3% . On Online Shoppers, even adjusted few-shot Jev reaches only 53.6% , against 71.2% .
Threshold adjustment does not close these gaps. For context, a constant-class prediction scores 50% balanced accuracy on these binary tasks. This is evidence of weak performance on these particular datasets, not a claim about every tabular task.
One example per class lifts raw Breast Cancer performance from 61.0% to 88.8% , and Banking77 from 78.9% to 81.9% . But it lowers mean performance on AG News, SMS Spam, IMDb, and Iris.
Breast Cancer also improves to 88.4% using zero-shot prompts with a learned threshold alone. Its raw 61.0% score therefore does not tell the whole story. The holdout is small, and none of these Jev variants reaches the strongest classical result.
Balanced accuracy is the average recall across classes. It gives each class equal weight, even when most examples belong to one class. For 77-class Banking77, the constant-class baseline is about 1.3%; for a binary task it is 50%. These are different problems, so we avoid collapsing this suite into one overall score.
One holdout.
Separate decisions.
The protocol separates model selection from decision-threshold selection. All models are evaluated on the same test cases for a dataset. The test set stays fixed across training seeds.
Up to 8,000 rows. Classical models fit here; Jev’s few-shot examples are drawn from here.
Up to 1,000 rows. Four candidate configurations per classical family are compared.
Up to 500 labeled rows. Binary thresholds are selected independently of the test set.
One fixed holdout per dataset, shared by models and seeds. Small datasets retain fewer rows.
Logistic regression, SVM, decision tree, random forest, extra trees, k-NN, Naive Bayes, histogram gradient boosting, XGBoost, CatBoost, and a voting ensemble. A majority baseline is also reported.
The V3 speed preset caps trees at 150 and histogram iterations at 60. Text vocabulary and representation budgets are limited. CPU and GPU implementations are explicitly mixed; backend changes are not claimed to be numerically equivalent.
The requested model is jev-1.13.0 . Zero-shot prompts provide task and class descriptions; few-shot prompts add one labeled training example per class. Tabular rows are passed as structured feature values.
Successful identical API requests are cached. Exhausted request failures count as incorrect predictions. Binary adjusted results use labeled policy data, even when their prompts contain no examples.
Training seeds: 2027, 2028, 2029. Holdout seed: 20260920. Kaggle configuration: two T4 GPUs. Source and protocol are frozen inside the completed notebook. Read the full protocol ↗
What this release
can’t establish.
The displayed ± values measure variation across training seeds. They are not confidence intervals. Paired bootstrap intervals were computed by the reporting code, but their output was not supplied with the notebook. We make no significance claims from these tables.
Independent zero-shot repetitions
Identical successful requests are reused from cache across seeds. The three zero-shot rows therefore do not represent three independent API replications. Policy splits can still produce different adjusted thresholds.
Pure model quality on Banking77
The saved Jev run includes warnings about predictions outside the true label set. The adapter uses −1 for failed requests and counts them as incorrect. Without the run diagnostics, we cannot quantify how much of the score reflects API failures.
Generalization from tiny holdouts
Iris has only 30 test cases and Breast Cancer has 114. Perfect classical scores on these samples do not imply perfect performance on new data. Public-dataset pretraining exposure is neither established nor ruled out.
Four candidates and restricted feature, training, and tree budgets are practical constraints. This suite does not cover all classical tuning strategies, pretrained embedding pipelines, fine-tuned transformers, or other language-model APIs.
Latency, cost, or full reproducibility
V3 latency summaries and the full Kaggle results archive are unavailable here. The run records 38,922 request attempts; its approximately $4.19 input-cost estimate is not an invoice. Exact snapshots, per-example predictions, and diagnostics are not included.
Follow the numbers
back to the run.
The completed notebook is preserved byte for byte. Both CSVs are extracted from its saved final outputs, retaining their displayed rounding. The interactive figures use those same tables.
Promising for some language tasks.
Evaluate before generalizing.

## Original Extract

An independent comparison of Jev 1.13.0 and 11 classical classification pipelines. Explore the completed V3 results, methodology, and limitations.

Skip to results
QuicqDev Research notes Results Analysis Method GitHub ↗
Independent benchmark Report 03 / September 2026
Jev vs.
classical ML.
Strong on sentiment.
Mixed across tasks.
We compared Jev 1.13.0 with eleven classical classification pipelines across eight datasets. Its largest advantage was on IMDb. Across the rest of the suite, the picture was more nuanced.
Protocol 3.0.1 · Bounded-budget comparison · Balanced accuracy · Not affiliated with OpenAI or Anthropic
Every model.
Both decision rules.
Switch between raw and threshold-adjusted decisions. Jev’s columns appear first. Each cell shows mean ± sample standard deviation across three training seeds on the same test cases. Bold blue cells mark the best displayed mean in each row, including ties.
Raw: default decision rules after classical model selection. Jev zero-shot uses no task examples or policy threshold fitting.
Scroll horizontally to inspect all columns. Sample SD is not a confidence interval. Zero SD on cached zero-shot predictions is not independent evidence of repeatability.
A different story
for each task.
Compare Jev with the highest-scoring classical pipeline on each dataset. The controls in the full table above also update this chart. Compare raw decisions with binary thresholds learned from separate labeled data.
Figure 1. Mean balanced accuracy, with a common 0–100% scale. Values are rounded to one decimal. The classical comparator is the highest test mean among the eleven pipelines in that panel, selected retrospectively. It is not a deployment selection rule. Read uncertainty notes ↓
Raw zero-shot Jev leads the best classical mean on IMDb and SMS Spam. The largest lead is IMDb: +7.9 percentage points. These are descriptive differences, not statistical-significance claims.
The headline is only
part of the result.
On IMDb, raw zero-shot Jev reaches 96.3% , compared with 88.4% for logistic regression. The +7.9-point advantage remains almost unchanged after threshold adjustment: 96.1% versus 88.3%.
This is the strongest descriptive evidence in Jev’s favor here. The experiment compares a pretrained API model against these classical text pipelines; it does not establish superiority over other language models or stronger modern text representations.
Thresholds change the SMS story.
With raw decisions, Jev zero-shot leads Naive Bayes 96.1% to 95.0% . After policy threshold selection, the comparison becomes 95.9% to 96.3% .
The apparent lead becomes a small deficit. A model’s default decision rule and its ability to separate classes are not the same question. Publish both panels; the adjusted Jev result uses labeled policy data.
Business tabular tasks remain difficult.
On Bank Marketing, adjusted zero-shot-prompt Jev reaches 59.7% , against the voting ensemble’s 73.3% . On Online Shoppers, even adjusted few-shot Jev reaches only 53.6% , against 71.2% .
Threshold adjustment does not close these gaps. For context, a constant-class prediction scores 50% balanced accuracy on these binary tasks. This is evidence of weak performance on these particular datasets, not a claim about every tabular task.
One example per class lifts raw Breast Cancer performance from 61.0% to 88.8% , and Banking77 from 78.9% to 81.9% . But it lowers mean performance on AG News, SMS Spam, IMDb, and Iris.
Breast Cancer also improves to 88.4% using zero-shot prompts with a learned threshold alone. Its raw 61.0% score therefore does not tell the whole story. The holdout is small, and none of these Jev variants reaches the strongest classical result.
Balanced accuracy is the average recall across classes. It gives each class equal weight, even when most examples belong to one class. For 77-class Banking77, the constant-class baseline is about 1.3%; for a binary task it is 50%. These are different problems, so we avoid collapsing this suite into one overall score.
One holdout.
Separate decisions.
The protocol separates model selection from decision-threshold selection. All models are evaluated on the same test cases for a dataset. The test set stays fixed across training seeds.
Up to 8,000 rows. Classical models fit here; Jev’s few-shot examples are drawn from here.
Up to 1,000 rows. Four candidate configurations per classical family are compared.
Up to 500 labeled rows. Binary thresholds are selected independently of the test set.
One fixed holdout per dataset, shared by models and seeds. Small datasets retain fewer rows.
Logistic regression, SVM, decision tree, random forest, extra trees, k-NN, Naive Bayes, histogram gradient boosting, XGBoost, CatBoost, and a voting ensemble. A majority baseline is also reported.
The V3 speed preset caps trees at 150 and histogram iterations at 60. Text vocabulary and representation budgets are limited. CPU and GPU implementations are explicitly mixed; backend changes are not claimed to be numerically equivalent.
The requested model is jev-1.13.0 . Zero-shot prompts provide task and class descriptions; few-shot prompts add one labeled training example per class. Tabular rows are passed as structured feature values.
Successful identical API requests are cached. Exhausted request failures count as incorrect predictions. Binary adjusted results use labeled policy data, even when their prompts contain no examples.
Training seeds: 2027, 2028, 2029. Holdout seed: 20260920. Kaggle configuration: two T4 GPUs. Source and protocol are frozen inside the completed notebook. Read the full protocol ↗
What this release
can’t establish.
The displayed ± values measure variation across training seeds. They are not confidence intervals. Paired bootstrap intervals were computed by the reporting code, but their output was not supplied with the notebook. We make no significance claims from these tables.
Independent zero-shot repetitions
Identical successful requests are reused from cache across seeds. The three zero-shot rows therefore do not represent three independent API replications. Policy splits can still produce different adjusted thresholds.
Pure model quality on Banking77
The saved Jev run includes warnings about predictions outside the true label set. The adapter uses −1 for failed requests and counts them as incorrect. Without the run diagnostics, we cannot quantify how much of the score reflects API failures.
Generalization from tiny holdouts
Iris has only 30 test cases and Breast Cancer has 114. Perfect classical scores on these samples do not imply perfect performance on new data. Public-dataset pretraining exposure is neither established nor ruled out.
Four candidates and restricted feature, training, and tree budgets are practical constraints. This suite does not cover all classical tuning strategies, pretrained embedding pipelines, fine-tuned transformers, or other language-model APIs.
Latency, cost, or full reproducibility
V3 latency summaries and the full Kaggle results archive are unavailable here. The run records 38,922 request attempts; its approximately $4.19 input-cost estimate is not an invoice. Exact snapshots, per-example predictions, and diagnostics are not included.
Follow the numbers
back to the run.
The completed notebook is preserved byte for byte. Both CSVs are extracted from its saved final outputs, retaining their displayed rounding. The interactive figures use those same tables.
Promising for some language tasks.
Evaluate before generalizing.
