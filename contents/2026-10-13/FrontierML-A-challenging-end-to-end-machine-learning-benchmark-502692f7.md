---
source: "https://www.frontierml.org/blog/"
hn_url: "https://news.ycombinator.com/item?id=50075487"
title: "FrontierML: A challenging, end-to-end machine learning benchmark"
article_title: "FrontierML: a challenging, end-to-end machine learning benchmark · FrontierML"
image: "https://www.frontierml.org/opengraph-image.jpg?opengraph-image.2i8la7bez7h87.jpg"
author: "meetpateltech"
captured_at: "2026-10-13T17:58:08Z"
capture_tool: "hn-digest"
hn_id: 50075487
score: 4
comments: 1
posted_at: "2026-10-13T17:32:12Z"
tags:
  - hacker-news
---

# FrontierML: A challenging, end-to-end machine learning benchmark

- HN: [50075487](https://news.ycombinator.com/item?id=50075487)
- Source: [www.frontierml.org](https://www.frontierml.org/blog/)
- Score: 4
- Comments: 1
- Posted: 2026-10-13T17:32:12Z

## Translation

Title: FrontierML: A challenging, end-to-end machine learning benchmark
Article title: FrontierML: a challenging, end-to-end machine learning benchmark · FrontierML
Description: We give frontier agents twelve hours, a CPU sandbox and a wallet of $10, $50 or $200, and ask each to train a model on one of 21 machine learning tasks.

Article text:
FrontierML: a challenging, end-to-end machine learning benchmark · FrontierML FrontierML tasks blog github FrontierML: a challenging, end-to-end machine learning benchmark
Scaling capital access for agents
One desirable property of human-run frontier AI labs is that they can consistently improve their models with increased access to capital. As agents begin doing machine learning work, the same question applies: given more money, can an agent build a better model?
FrontierML is a benchmark that measures how well agents turn capital into model performance on challenging, end-to-end machine learning tasks. Our benchmark makes several key advancements in capturing the real-world complexity of modern machine learning:
Scaling spending budgets: Existing benchmarks scale how much agents can think . We scale how much they can spend . We give agents $10, $50, and $200 per task to spend on training compute and data, and evaluate whether agents can use the additional capital productively.
Open-ended, end-to-end tasks: FrontierML is the most open-ended benchmark for agentic machine learning. In each task, agents can choose architectures, source data, implement training, provision compute, evaluate checkpoints, and optimize inference. Each task is scored on a continuous scale based on the final performance of the submitted model.
Robust verification: Agents submit an inference server that serves their model on an air-gapped H100 machine. Submissions are evaluated on held-out out-of-distribution tests, and a separate reviewer agent examines each submission and its trajectory for violations of task rules.
Diverse and real-world focus: FrontierML contains 21 long-horizon tasks inspired by real work at frontier labs, neolabs, and academic institutions.
Overall, FrontierML faithfully measures whether or not agents can autonomously train other models, end-to-end, while handling the complexity and applying the judgment a human research team would.
$ 10 $ 50 $ 200 $0 $50 $100 $150 $200 0 10 20 30 40 50 60 70 Total spend per task (wallet spend + agent API cost) FrontierML Score Claude Fable 5.1 Claude Haiku 5.5 GPT-6.1 Sol GPT-6 Luna Muse Wonder GPT-6 Astra Claude Opus 5.5 Claude Sonnet 5.5 Human Researcher Kimi K3 Qwen3.8 Max GLM 5.3 Gemini 3.8 Flash Grok 4.7 Muse Spark 1.3 Link to this section Overview
Link to this section Scaling capital access for agents
In the real world, researchers might spend $10 or $200 or $50,000 toward a given task, with performance improving as budget scales. FrontierML measures whether agents can train better models as we scale the money they are given.
FrontierML tasks are designed to give agents control over the full model development process, creating opportunities to turn additional capital into better model performance through larger training runs, better data, or more algorithmic experimentation.
In our experiments, we find that stronger agents are able to reliably improve their submissions with greater access to capital.
$50 $200 Claude Opus 5.5 45.4 58.8 Claude Opus 5.5 Claude Sonnet 5.5 45.5 50.4 Claude Sonnet 5.5 Claude Fable 5.1 35.7 41.9 Claude Fable 5.1 Claude Haiku 5.5 28.2 34.9 Claude Haiku 5.5 GPT-6 Astra 37.6 44.5 GPT-6 Astra GPT-6.1 Sol 31.6 38.7 GPT-6.1 Sol GPT-6 Luna 13.2 13.1 GPT-6 Luna Gemini 3.8 Flash 8.1 5.7 Gemini 3.8 Flash Grok 4.7 12.2 10.9 Grok 4.7 Muse Spark 1.3 14.1 16.0 Muse Spark 1.3 Muse Wonder 21.6 19.3 Muse Wonder Kimi K3 16.2 20.3 Kimi K3 Qwen3.8 Max 19.9 17.5 Qwen3.8 Max GLM 5.3 18.6 19.6 GLM 5.3 Human Researcher 56.0 66.1 Human Researcher $10 $50 $200 Claude Opus 5.5 45.4 58.8 Claude Opus 5.5 Claude Sonnet 5.5 45.5 50.4 Claude Sonnet 5.5 Claude Fable 5.1 35.7 41.9 Claude Fable 5.1 Claude Haiku 5.5 28.2 34.9 Claude Haiku 5.5 GPT-6 Astra 37.6 44.5 GPT-6 Astra GPT-6.1 Sol 31.6 38.7 GPT-6.1 Sol GPT-6 Luna 13.2 13.1 GPT-6 Luna Gemini 3.8 Flash 8.1 5.7 Gemini 3.8 Flash Grok 4.7 12.2 10.9 Grok 4.7 Muse Spark 1.3 14.1 16.0 Muse Spark 1.3 Muse Wonder 21.6 19.3 Muse Wonder Kimi K3 16.2 20.3 Kimi K3 Qwen3.8 Max 19.9 17.5 Qwen3.8 Max GLM 5.3 18.6 19.6 GLM 5.3 Human 56.0 66.1 Human Change in mean score over the 21 tasks at each wallet size. Link to this section Open-endedness
Like real-world machine learning work, FrontierML tasks define a problem without prescribing a solution. Agents build and submit their own inference server, which we evaluate by sending requests and scoring its responses. This leaves the entire development process open-ended: agents decide what model architecture to build, what data to use, which GPUs to rent and for how long, which experiments to run, and how to serve the result.
Agents adopt different approaches for each task, resulting in a wide band of scores.
mean score Fine-tuned diffusion model (LoRA) + exact text overlay Sonnet 61
LLM writes a drawing program (HTML/SVG/PIL) that is rendered Astra Fable Sol 38
Off-the-shelf diffusion model + exact text overlay Opus Haiku 33
Hand-coded procedural renderer, no generative model Kimi GLM Wonder Qwen Grok Luna Spark Gemini 20
Link to this section Preventing reward hacking
Open-ended long-horizon tasks have a large reward hacking surface, which we address with the following:
Air-gapped solution execution . The agent has internet access during development, but must submit a standalone package of serving code and model weights that runs with no internet access, on a single H100 machine.
Extensive held-out test set . To disincentivize overfitting, we provide limited validation samples and evaluate on a large out-of-distribution held-out test set, not shown to the agent, for each task. For example, for the Ternary PTQ task , we only tell the agent two of four benchmarks we will evaluate on (public: GPQA-Diamond, LCB v6; held-out: MMLU-Pro, MATH-500). In the SVG generation task , we make the held-out samples substantially more complex than the validation samples.
LLM judge review . After each task is complete, we spin up a review agent with access to the submission and the agent’s trajectory, which checks for cheating, including:
Searching up answers : did the agent search for FrontierML solutions?
Training on test set : did the agent train on the validation/test split of any public benchmark?
Invalid model provenance : did the agent not use the requested base model if provided, or did the agent serve a prohibited model?
Strict time limits . We set an allowable evaluation time limit to bound the allowed inference compute for each submission. If the submitted server doesn’t return answers for each held-out sample within the limit, the agent scores 0 on the task.
Link to this section Real-world tasks
FrontierML contains 21 end-to-end tasks spanning language, vision, audio, and games, inspired by research at frontier labs, neolabs, and academic institutions. Tasks range from converting an autoregressive LLM into a text diffusion model to building a streaming end-of-turn detector for voice assistants to writing a bot that plays competitive Pokémon battles.
Data-efficient pretraining Q Labs
AR to diffusion conversion Inception Labs
LLM post-training PostTrainBench
OCR latent compression Induction Labs
Voice endpointing Thinking Machines
Vision for a ternary LM Deepgrove
Video frame interpolation NVIDIA
Inverse dynamics model Standard Intelligence
Screenshot generation University of Waterloo
Zero-shot classification Typesafe
Math proof simplification Axiom
Pokémon Random Battle Jaxcalibur
Link to this section Methodology
Link to this section Evaluation setup
Each task provides the agent with the same setup: a 12-hour CPU sandbox, a wallet with $10, $50, or $200, a few validation samples, and a task description. For each task, we also provide several constraints to ensure tasks are not trivially solvable, for example by restricting use of already-trained public specialist models or training on the test split of any public benchmark.
We provide the minimal amount of information required to complete the task. After submission, a review agent is given a sandbox to review the submission and agent’s trajectory to ensure the constraints were followed.
When evaluating each agent, we use the reasoning effort that provided the best score on other agentic benchmarks for each respective agent. Each model runs in the harness its developer used for agentic benchmarks on its model card, mirroring day-to-day use.
Before submitting, agents can test their work with frontier selfcheck , which runs the current submission on the grading infrastructure against the public validation samples, up to four times per trial.
Evaluating an agent on all 21 tasks costs up to $ 5,460 in wallet spend, not including model API costs.
Develop 12 h · $10 / $50 / $200 Spin wallet GPUs + LLM tokens Agent CPU sandbox Submission weights + server Self-check H100, no score iterate Grade once Held-out tests air-gapped H100 Opus 5.5 review code + trajectory Score 0 if the review fails submit In a 12-hour CPU sandbox, the agent pays for GPUs and tokens through Spin out of its wallet and can run its submission on the grading machine up to four times. The final submission is then served once on an air-gapped H100 against held-out tests, and Claude Opus 5.5 reviews it alongside the agent's trajectory. Link to this section Access to compute and data
To give each agent simple access to compute and synthetic data generation, we created Spin , an open-source CLI tool and service that allows agents to purchase on-demand compute from 17 different compute providers, and metered tokens from OpenRouter. We provide a hosted version with zero fees for ease of evaluation.
Because compute prices fluctuate, we fix the price of each GPU type across all evaluation runs. We took a snapshot of on-demand compute prices on September 29, 2026, calculated the median price across providers and saved the result. We price GPUs by form factor, with H100 SXM, H100 PCIe and H100 NVL each having their own fixed price. The snapshot used for evaluation is available as a JSON file. When running the evaluation, we set compute to the price captured in the snapshot by default.
Claude Opus 5.5 Wallet spend $127.67 · Model cost $55.06 58.8 $127.67 $55.06 2 Claude Sonnet 5.5 Wallet spend $125.17 · Model cost $46.95 50.4 $125.17 $46.95 3 GPT-6 Astra Wallet spend $100.40 · Model cost $116.22 44.5 $100.40 $116.22 4 Claude Fable 5.1 Wallet spend $79.02 · Model cost $102.84 41.9 $79.02 $102.84 5 GPT-6.1 Sol Wallet spend $103.40 · Model cost $10.96 38.7 $103.40 $10.96 6 Claude Haiku 5.5 Wallet spend $58.92 · Model cost $10.38 34.9 $58.92 $10.38 7 Kimi K3 Wallet spend $61.36 · Model cost $48.79 20.3 $61.36 $48.79 8 GLM 5.3 Wallet spend $41.58 · Model cost $27.62 19.6 $41.58 $27.62 9 Muse Wonder Wallet spend $33.39 · Model cost $17.06 19.3 $33.39 $17.06 10 Qwen3.8 Max Wallet spend $65.43 · Model cost $18.54 17.5 $65.43 $18.54 11 Muse Spark 1.3 Wallet spend $41.22 · Model cost $17.71 16.0 $41.22 $17.71 12 GPT-6 Luna Wallet spend $17.15 · Model cost $0.75 13.1 $17.15 $0.75 13 Grok 4.7 Wallet spend $20.22 · Model cost $30.02 10.9 $20.22 $30.02 14 Gemini 3.8 Flash Wallet spend $2.78 · Model cost $9.75 5.7 $2.78 $9.75 0 20 40 60 80 100 Link to this section Observations
Link to this section Agents make new architectures
Because FrontierML keeps tasks open-ended and allows agents to do whatever they think will produce the best model, agents design their own architectures and produce their own data.
In 28% of trajectories, agents add newly trained parts to pretrained models or train a new model from scratch.
Claude Opus 5.5 8 % 11 % 70 % 6 % Claude Sonnet 5.5 13 % 14 % 67 % 6 % GPT-6 Astra 13 % 27 % 49 % 11 % Claude Fable 5.1 16 % 6 % 63 % 11 % GPT-6.1 Sol 13 % 22 % 54 % 10 % Claude Haiku 5.5 16 % 10 % 63 % 6 % Kimi K3 16 % 17 % 56 % 10 % GLM 5.3 24 % 13 % 51 % 6 % 6 % Muse Wonder 13 % 21 % 59 % 6 % Qwen3.8 Max 17 % 8 % 62 % 8 % Muse Spark 1.3 13 % 13 % 59 % 11 % GPT-6 Luna 13 % 22 % 46 % 11 % 8 % Grok 4.7 11 % 11 % 37 % 25 % 16 % Gemini 3.8 Flash 6 % 10 % 37 % 32 % 16 % New m

[truncated]

## Original Extract

We give frontier agents twelve hours, a CPU sandbox and a wallet of $10, $50 or $200, and ask each to train a model on one of 21 machine learning tasks.

FrontierML: a challenging, end-to-end machine learning benchmark · FrontierML FrontierML tasks blog github FrontierML: a challenging, end-to-end machine learning benchmark
Scaling capital access for agents
One desirable property of human-run frontier AI labs is that they can consistently improve their models with increased access to capital. As agents begin doing machine learning work, the same question applies: given more money, can an agent build a better model?
FrontierML is a benchmark that measures how well agents turn capital into model performance on challenging, end-to-end machine learning tasks. Our benchmark makes several key advancements in capturing the real-world complexity of modern machine learning:
Scaling spending budgets: Existing benchmarks scale how much agents can think . We scale how much they can spend . We give agents $10, $50, and $200 per task to spend on training compute and data, and evaluate whether agents can use the additional capital productively.
Open-ended, end-to-end tasks: FrontierML is the most open-ended benchmark for agentic machine learning. In each task, agents can choose architectures, source data, implement training, provision compute, evaluate checkpoints, and optimize inference. Each task is scored on a continuous scale based on the final performance of the submitted model.
Robust verification: Agents submit an inference server that serves their model on an air-gapped H100 machine. Submissions are evaluated on held-out out-of-distribution tests, and a separate reviewer agent examines each submission and its trajectory for violations of task rules.
Diverse and real-world focus: FrontierML contains 21 long-horizon tasks inspired by real work at frontier labs, neolabs, and academic institutions.
Overall, FrontierML faithfully measures whether or not agents can autonomously train other models, end-to-end, while handling the complexity and applying the judgment a human research team would.
$ 10 $ 50 $ 200 $0 $50 $100 $150 $200 0 10 20 30 40 50 60 70 Total spend per task (wallet spend + agent API cost) FrontierML Score Claude Fable 5.1 Claude Haiku 5.5 GPT-6.1 Sol GPT-6 Luna Muse Wonder GPT-6 Astra Claude Opus 5.5 Claude Sonnet 5.5 Human Researcher Kimi K3 Qwen3.8 Max GLM 5.3 Gemini 3.8 Flash Grok 4.7 Muse Spark 1.3 Link to this section Overview
Link to this section Scaling capital access for agents
In the real world, researchers might spend $10 or $200 or $50,000 toward a given task, with performance improving as budget scales. FrontierML measures whether agents can train better models as we scale the money they are given.
FrontierML tasks are designed to give agents control over the full model development process, creating opportunities to turn additional capital into better model performance through larger training runs, better data, or more algorithmic experimentation.
In our experiments, we find that stronger agents are able to reliably improve their submissions with greater access to capital.
$50 $200 Claude Opus 5.5 45.4 58.8 Claude Opus 5.5 Claude Sonnet 5.5 45.5 50.4 Claude Sonnet 5.5 Claude Fable 5.1 35.7 41.9 Claude Fable 5.1 Claude Haiku 5.5 28.2 34.9 Claude Haiku 5.5 GPT-6 Astra 37.6 44.5 GPT-6 Astra GPT-6.1 Sol 31.6 38.7 GPT-6.1 Sol GPT-6 Luna 13.2 13.1 GPT-6 Luna Gemini 3.8 Flash 8.1 5.7 Gemini 3.8 Flash Grok 4.7 12.2 10.9 Grok 4.7 Muse Spark 1.3 14.1 16.0 Muse Spark 1.3 Muse Wonder 21.6 19.3 Muse Wonder Kimi K3 16.2 20.3 Kimi K3 Qwen3.8 Max 19.9 17.5 Qwen3.8 Max GLM 5.3 18.6 19.6 GLM 5.3 Human Researcher 56.0 66.1 Human Researcher $10 $50 $200 Claude Opus 5.5 45.4 58.8 Claude Opus 5.5 Claude Sonnet 5.5 45.5 50.4 Claude Sonnet 5.5 Claude Fable 5.1 35.7 41.9 Claude Fable 5.1 Claude Haiku 5.5 28.2 34.9 Claude Haiku 5.5 GPT-6 Astra 37.6 44.5 GPT-6 Astra GPT-6.1 Sol 31.6 38.7 GPT-6.1 Sol GPT-6 Luna 13.2 13.1 GPT-6 Luna Gemini 3.8 Flash 8.1 5.7 Gemini 3.8 Flash Grok 4.7 12.2 10.9 Grok 4.7 Muse Spark 1.3 14.1 16.0 Muse Spark 1.3 Muse Wonder 21.6 19.3 Muse Wonder Kimi K3 16.2 20.3 Kimi K3 Qwen3.8 Max 19.9 17.5 Qwen3.8 Max GLM 5.3 18.6 19.6 GLM 5.3 Human 56.0 66.1 Human Change in mean score over the 21 tasks at each wallet size. Link to this section Open-endedness
Like real-world machine learning work, FrontierML tasks define a problem without prescribing a solution. Agents build and submit their own inference server, which we evaluate by sending requests and scoring its responses. This leaves the entire development process open-ended: agents decide what model architecture to build, what data to use, which GPUs to rent and for how long, which experiments to run, and how to serve the result.
Agents adopt different approaches for each task, resulting in a wide band of scores.
mean score Fine-tuned diffusion model (LoRA) + exact text overlay Sonnet 61
LLM writes a drawing program (HTML/SVG/PIL) that is rendered Astra Fable Sol 38
Off-the-shelf diffusion model + exact text overlay Opus Haiku 33
Hand-coded procedural renderer, no generative model Kimi GLM Wonder Qwen Grok Luna Spark Gemini 20
Link to this section Preventing reward hacking
Open-ended long-horizon tasks have a large reward hacking surface, which we address with the following:
Air-gapped solution execution . The agent has internet access during development, but must submit a standalone package of serving code and model weights that runs with no internet access, on a single H100 machine.
Extensive held-out test set . To disincentivize overfitting, we provide limited validation samples and evaluate on a large out-of-distribution held-out test set, not shown to the agent, for each task. For example, for the Ternary PTQ task , we only tell the agent two of four benchmarks we will evaluate on (public: GPQA-Diamond, LCB v6; held-out: MMLU-Pro, MATH-500). In the SVG generation task , we make the held-out samples substantially more complex than the validation samples.
LLM judge review . After each task is complete, we spin up a review agent with access to the submission and the agent’s trajectory, which checks for cheating, including:
Searching up answers : did the agent search for FrontierML solutions?
Training on test set : did the agent train on the validation/test split of any public benchmark?
Invalid model provenance : did the agent not use the requested base model if provided, or did the agent serve a prohibited model?
Strict time limits . We set an allowable evaluation time limit to bound the allowed inference compute for each submission. If the submitted server doesn’t return answers for each held-out sample within the limit, the agent scores 0 on the task.
Link to this section Real-world tasks
FrontierML contains 21 end-to-end tasks spanning language, vision, audio, and games, inspired by research at frontier labs, neolabs, and academic institutions. Tasks range from converting an autoregressive LLM into a text diffusion model to building a streaming end-of-turn detector for voice assistants to writing a bot that plays competitive Pokémon battles.
Data-efficient pretraining Q Labs
AR to diffusion conversion Inception Labs
LLM post-training PostTrainBench
OCR latent compression Induction Labs
Voice endpointing Thinking Machines
Vision for a ternary LM Deepgrove
Video frame interpolation NVIDIA
Inverse dynamics model Standard Intelligence
Screenshot generation University of Waterloo
Zero-shot classification Typesafe
Math proof simplification Axiom
Pokémon Random Battle Jaxcalibur
Link to this section Methodology
Link to this section Evaluation setup
Each task provides the agent with the same setup: a 12-hour CPU sandbox, a wallet with $10, $50, or $200, a few validation samples, and a task description. For each task, we also provide several constraints to ensure tasks are not trivially solvable, for example by restricting use of already-trained public specialist models or training on the test split of any public benchmark.
We provide the minimal amount of information required to complete the task. After submission, a review agent is given a sandbox to review the submission and agent’s trajectory to ensure the constraints were followed.
When evaluating each agent, we use the reasoning effort that provided the best score on other agentic benchmarks for each respective agent. Each model runs in the harness its developer used for agentic benchmarks on its model card, mirroring day-to-day use.
Before submitting, agents can test their work with frontier selfcheck , which runs the current submission on the grading infrastructure against the public validation samples, up to four times per trial.
Evaluating an agent on all 21 tasks costs up to $ 5,460 in wallet spend, not including model API costs.
Develop 12 h · $10 / $50 / $200 Spin wallet GPUs + LLM tokens Agent CPU sandbox Submission weights + server Self-check H100, no score iterate Grade once Held-out tests air-gapped H100 Opus 5.5 review code + trajectory Score 0 if the review fails submit In a 12-hour CPU sandbox, the agent pays for GPUs and tokens through Spin out of its wallet and can run its submission on the grading machine up to four times. The final submission is then served once on an air-gapped H100 against held-out tests, and Claude Opus 5.5 reviews it alongside the agent's trajectory. Link to this section Access to compute and data
To give each agent simple access to compute and synthetic data generation, we created Spin , an open-source CLI tool and service that allows agents to purchase on-demand compute from 17 different compute providers, and metered tokens from OpenRouter. We provide a hosted version with zero fees for ease of evaluation.
Because compute prices fluctuate, we fix the price of each GPU type across all evaluation runs. We took a snapshot of on-demand compute prices on September 29, 2026, calculated the median price across providers and saved the result. We price GPUs by form factor, with H100 SXM, H100 PCIe and H100 NVL each having their own fixed price. The snapshot used for evaluation is available as a JSON file. When running the evaluation, we set compute to the price captured in the snapshot by default.
Claude Opus 5.5 Wallet spend $127.67 · Model cost $55.06 58.8 $127.67 $55.06 2 Claude Sonnet 5.5 Wallet spend $125.17 · Model cost $46.95 50.4 $125.17 $46.95 3 GPT-6 Astra Wallet spend $100.40 · Model cost $116.22 44.5 $100.40 $116.22 4 Claude Fable 5.1 Wallet spend $79.02 · Model cost $102.84 41.9 $79.02 $102.84 5 GPT-6.1 Sol Wallet spend $103.40 · Model cost $10.96 38.7 $103.40 $10.96 6 Claude Haiku 5.5 Wallet spend $58.92 · Model cost $10.38 34.9 $58.92 $10.38 7 Kimi K3 Wallet spend $61.36 · Model cost $48.79 20.3 $61.36 $48.79 8 GLM 5.3 Wallet spend $41.58 · Model cost $27.62 19.6 $41.58 $27.62 9 Muse Wonder Wallet spend $33.39 · Model cost $17.06 19.3 $33.39 $17.06 10 Qwen3.8 Max Wallet spend $65.43 · Model cost $18.54 17.5 $65.43 $18.54 11 Muse Spark 1.3 Wallet spend $41.22 · Model cost $17.71 16.0 $41.22 $17.71 12 GPT-6 Luna Wallet spend $17.15 · Model cost $0.75 13.1 $17.15 $0.75 13 Grok 4.7 Wallet spend $20.22 · Model cost $30.02 10.9 $20.22 $30.02 14 Gemini 3.8 Flash Wallet spend $2.78 · Model cost $9.75 5.7 $2.78 $9.75 0 20 40 60 80 100 Link to this section Observations
Link to this section Agents make new architectures
Because FrontierML keeps tasks open-ended and allows agents to do whatever they think will produce the best model, agents design their own architectures and produce their own data.
In 28% of trajectories, agents add newly trained parts to pretrained models or train a new model from scratch.
Claude Opus 5.5 8 % 11 % 70 % 6 % Claude Sonnet 5.5 13 % 14 % 67 % 6 % GPT-6 Astra 13 % 27 % 49 % 11 % Claude Fable 5.1 16 % 6 % 63 % 11 % GPT-6.1 Sol 13 % 22 % 54 % 10 % Claude Haiku 5.5 16 % 10 % 63 % 6 % Kimi K3 16 % 17 % 56 % 10 % GLM 5.3 24 % 13 % 51 % 6 % 6 % Muse Wonder 13 % 21 % 59 % 6 % Qwen3.8 Max 17 % 8 % 62 % 8 % Muse Spark 1.3 13 % 13 % 59 % 11 % GPT-6 Luna 13 % 22 % 46 % 11 % 8 % Grok 4.7 11 % 11 % 37 % 25 % 16 % Gemini 3.8 Flash 6 % 10 % 37 % 32 % 16 % New m

[truncated]
