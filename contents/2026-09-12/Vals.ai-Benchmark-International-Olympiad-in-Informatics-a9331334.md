---
source: "https://www.vals.ai/benchmarks/ioi"
hn_url: "https://news.ycombinator.com/item?id=49673057"
title: "Vals.ai Benchmark – International Olympiad in Informatics"
article_title: "IOI"
image: "https://vals.ai/og/benchmarks/ioi.png"
author: "ZeljkoS"
captured_at: "2026-09-12T16:01:27Z"
capture_tool: "hn-digest"
hn_id: 49673057
score: 3
comments: 1
posted_at: "2026-09-12T15:05:04Z"
tags:
  - hacker-news
---

# Vals.ai Benchmark – International Olympiad in Informatics

- HN: [49673057](https://news.ycombinator.com/item?id=49673057)
- Source: [www.vals.ai](https://www.vals.ai/benchmarks/ioi)
- Score: 3
- Comments: 1
- Posted: 2026-09-12T15:05:04Z

## Translation

Title: Vals.ai Benchmark – International Olympiad in Informatics
Article title: IOI
Description: Informatics olympiad programming problems

Article text:
International Olympiad in Informatics
ACCURACY Cost Alibaba Anthropic DeepSeek Google Meta OpenAI SpaceXAI zAI Showing latest, top & frontier models (19) Key Takeaways
Unlike the saturated knowledge benchmarks, IOI still sharply separates models: GPT-6 Astra solves every problem in all three years, GPT-5.6 Sol (91.17%), Claude Fable 5.1 (90.78%) and GPT-5.6 Terra (87.61%) follow, and the twenty-six-model field then spreads across ninety points, down to 9.33%, so competitive-programming ability remains a real differentiator.
The agent sees only what a contestant sees: the statement, the sample grader and one sample. It has no tests, no internet and no submission feedback, so every point comes from code the model wrote and checked itself.
Scores decline on the newest problems: cohort mean accuracy is 58.38% on 2024 and 53.67% on 2025 but 50.82% on 2026, and fifteen of the twenty-six models score lowest on the 2026 problems. GPT 5.3 Codex drops from 62.83% on 2024 to 41.67% on 2026. The two leaders score 100% on 2026 even though the contest ran after their training cutoffs, so recency alone does not cap performance.
Recently, top LLM labs like OpenAI and Google reported that their models achieved gold medals on the International Mathematical Olympiad (IMO). However, advanced models are starting to saturate IMO, meaning it may no longer effectively differentiate between the capabilities of top-performing models.
Reports also suggest the evaluation process faced coordination challenges, with AI companies seeking expedited validation mid-competition that may not reflect standard IMO assessment procedures.
The International Olympiad in Informatics (IOI) offers several advantages as an LLM benchmark. Unlike the IMO, the IOI is not yet saturated, providing clear differentiation between model capabilities. The competition features standardized and automated grading, ensuring objective evaluation without subjective scoring. Additionally, the IOI has real-world relevance as it tests C++ programming skills that are directly applicable to software development.
We designed our benchmark to imitate competition conditions as closely as possible.
Each model runs inside OpenCode , a general-purpose coding agent, in an isolated sandbox with a c++ (v20) toolchain. Its workspace holds exactly what the contest hands a contestant : the problem statement, the task header and solution stub, the sample grader or manager, the compile and run scripts, and the sample input and output from the statement. The official test data, the subtask test lists, and the grading script are withheld until the attempt is over.
The agent reads, writes, compiles and tests files freely in its workspace, and writes its answer to /workspace/solution.cpp . Like a contestant, it has no access to the public internet: the sandbox reaches only our model gateway, so published editorials and reference solutions are out of reach.
Unlike a contestant, the agent has no submission tool and therefore receives no score feedback during the attempt: it can run the samples and whatever tests it writes for itself, and must decide on its own when its solution is good enough.
Some models spend their entire per-response output budget reasoning about a problem before writing any code. When that happens, the harness continues the cut-off turn on the next step rather than failing the problem: it replays the provider’s own reasoning state where the API returns one, and otherwise restarts the turn with an instruction to work in smaller steps. The same rule applies to every model on this board, and the cost column includes the reasoning spent this way.
Results from our earlier harness, which gave the agent an interactive grading tool, are preserved on IOI v1 . The two harnesses are not comparable, so their results are reported on separate boards.
Grading matches the olympiad: after the attempt ends, the solution left in the workspace is copied into a fresh grading directory alongside pristine copies of the official test data and grader, compiled, run against every official test, and scored per subtask out of 100 points. Every grading file is verified unchanged after the run.
A solution that fails to compile, or that the agent never wrote, scores zero.
A model’s score for a year is its mean score across that year’s six problems, and its overall score is the mean of its three yearly scores.
GPT-6 Astra reaches 100% accuracy, solving all eighteen problems, ahead of GPT-5.6 Sol at 91.17%, Claude Fable 5.1 at 90.78%, GPT-5.6 Terra at 87.61% and Claude Opus 5 at 84.33%. Fable 5.1 and Opus 5 both solve every 2024 problem and score 87.17% on 2025; Fable holds 85.17% on 2026 while Opus falls to 65.83%. Qwen 3.8 Max (68.89%), GLM 5.3 (68.44%), Gemini 3.7 Flash (67.83%), GPT-5.6 Luna (61.78%), Gemini 3.8 Flash (56.94%), Muse Spark 1.3 Max (56.56%), GPT 5.3 Codex (53.83%), GLM 5.3 Flash (52.50%), DeepSeek V4 Pro 0813 (51.61%), Kimi K3 (48.94%), Grok 4.6 (47.61%), Claude Sonnet 5 (45.00%) and Muse Spark 1.3 (43.94%) form a middle group, while Grok 4.5 (40.56%), DeepSeek V4.1 Flash (40.28%, the cheapest model on the board), Qwen 3.8 27B (39.06%, thirty points behind Qwen 3.8 Max), Gemini 3.6 Flash (35.06%), DeepSeek V4 Flash 0731 (32.72%), Muse Spark 1.2 (21.78%), Inkling (14.94%) and Inkling Small (9.33%) trail the field. Both Muse Spark 1.3 variants at least double the score of their predecessor Muse Spark 1.2. Claude Sonnet 5 is the most expensive model on the board at $19.33 per problem: on thirteen of its eighteen problems it spent its entire 128k-token output budget reasoning before writing any code, and the harness had to continue the cut-off turn.
We evaluate three years to check for data contamination.
The 2026 problems were released only after the models in this cohort were trained, and they are the hardest set for fifteen of the twenty-six: GPT 5.3 Codex scores 62.83% on 2024 and 57.00% on 2025 but 41.67% on 2026, and cohort mean accuracy falls from 58.38% (2024) and 53.67% (2025) to 50.82% (2026). GPT-6 Astra and GPT-5.6 Terra both score 100% on 2026, so the newest set is not out of reach.
Copyright © 2026 Vals AI. All rights reserved.

## Original Extract

Informatics olympiad programming problems

International Olympiad in Informatics
ACCURACY Cost Alibaba Anthropic DeepSeek Google Meta OpenAI SpaceXAI zAI Showing latest, top & frontier models (19) Key Takeaways
Unlike the saturated knowledge benchmarks, IOI still sharply separates models: GPT-6 Astra solves every problem in all three years, GPT-5.6 Sol (91.17%), Claude Fable 5.1 (90.78%) and GPT-5.6 Terra (87.61%) follow, and the twenty-six-model field then spreads across ninety points, down to 9.33%, so competitive-programming ability remains a real differentiator.
The agent sees only what a contestant sees: the statement, the sample grader and one sample. It has no tests, no internet and no submission feedback, so every point comes from code the model wrote and checked itself.
Scores decline on the newest problems: cohort mean accuracy is 58.38% on 2024 and 53.67% on 2025 but 50.82% on 2026, and fifteen of the twenty-six models score lowest on the 2026 problems. GPT 5.3 Codex drops from 62.83% on 2024 to 41.67% on 2026. The two leaders score 100% on 2026 even though the contest ran after their training cutoffs, so recency alone does not cap performance.
Recently, top LLM labs like OpenAI and Google reported that their models achieved gold medals on the International Mathematical Olympiad (IMO). However, advanced models are starting to saturate IMO, meaning it may no longer effectively differentiate between the capabilities of top-performing models.
Reports also suggest the evaluation process faced coordination challenges, with AI companies seeking expedited validation mid-competition that may not reflect standard IMO assessment procedures.
The International Olympiad in Informatics (IOI) offers several advantages as an LLM benchmark. Unlike the IMO, the IOI is not yet saturated, providing clear differentiation between model capabilities. The competition features standardized and automated grading, ensuring objective evaluation without subjective scoring. Additionally, the IOI has real-world relevance as it tests C++ programming skills that are directly applicable to software development.
We designed our benchmark to imitate competition conditions as closely as possible.
Each model runs inside OpenCode , a general-purpose coding agent, in an isolated sandbox with a c++ (v20) toolchain. Its workspace holds exactly what the contest hands a contestant : the problem statement, the task header and solution stub, the sample grader or manager, the compile and run scripts, and the sample input and output from the statement. The official test data, the subtask test lists, and the grading script are withheld until the attempt is over.
The agent reads, writes, compiles and tests files freely in its workspace, and writes its answer to /workspace/solution.cpp . Like a contestant, it has no access to the public internet: the sandbox reaches only our model gateway, so published editorials and reference solutions are out of reach.
Unlike a contestant, the agent has no submission tool and therefore receives no score feedback during the attempt: it can run the samples and whatever tests it writes for itself, and must decide on its own when its solution is good enough.
Some models spend their entire per-response output budget reasoning about a problem before writing any code. When that happens, the harness continues the cut-off turn on the next step rather than failing the problem: it replays the provider’s own reasoning state where the API returns one, and otherwise restarts the turn with an instruction to work in smaller steps. The same rule applies to every model on this board, and the cost column includes the reasoning spent this way.
Results from our earlier harness, which gave the agent an interactive grading tool, are preserved on IOI v1 . The two harnesses are not comparable, so their results are reported on separate boards.
Grading matches the olympiad: after the attempt ends, the solution left in the workspace is copied into a fresh grading directory alongside pristine copies of the official test data and grader, compiled, run against every official test, and scored per subtask out of 100 points. Every grading file is verified unchanged after the run.
A solution that fails to compile, or that the agent never wrote, scores zero.
A model’s score for a year is its mean score across that year’s six problems, and its overall score is the mean of its three yearly scores.
GPT-6 Astra reaches 100% accuracy, solving all eighteen problems, ahead of GPT-5.6 Sol at 91.17%, Claude Fable 5.1 at 90.78%, GPT-5.6 Terra at 87.61% and Claude Opus 5 at 84.33%. Fable 5.1 and Opus 5 both solve every 2024 problem and score 87.17% on 2025; Fable holds 85.17% on 2026 while Opus falls to 65.83%. Qwen 3.8 Max (68.89%), GLM 5.3 (68.44%), Gemini 3.7 Flash (67.83%), GPT-5.6 Luna (61.78%), Gemini 3.8 Flash (56.94%), Muse Spark 1.3 Max (56.56%), GPT 5.3 Codex (53.83%), GLM 5.3 Flash (52.50%), DeepSeek V4 Pro 0813 (51.61%), Kimi K3 (48.94%), Grok 4.6 (47.61%), Claude Sonnet 5 (45.00%) and Muse Spark 1.3 (43.94%) form a middle group, while Grok 4.5 (40.56%), DeepSeek V4.1 Flash (40.28%, the cheapest model on the board), Qwen 3.8 27B (39.06%, thirty points behind Qwen 3.8 Max), Gemini 3.6 Flash (35.06%), DeepSeek V4 Flash 0731 (32.72%), Muse Spark 1.2 (21.78%), Inkling (14.94%) and Inkling Small (9.33%) trail the field. Both Muse Spark 1.3 variants at least double the score of their predecessor Muse Spark 1.2. Claude Sonnet 5 is the most expensive model on the board at $19.33 per problem: on thirteen of its eighteen problems it spent its entire 128k-token output budget reasoning before writing any code, and the harness had to continue the cut-off turn.
We evaluate three years to check for data contamination.
The 2026 problems were released only after the models in this cohort were trained, and they are the hardest set for fifteen of the twenty-six: GPT 5.3 Codex scores 62.83% on 2024 and 57.00% on 2025 but 41.67% on 2026, and cohort mean accuracy falls from 58.38% (2024) and 53.67% (2025) to 50.82% (2026). GPT-6 Astra and GPT-5.6 Terra both score 100% on 2026, so the newest set is not out of reach.
Copyright © 2026 Vals AI. All rights reserved.
