---
source: "https://supercomputing-system-ai-lab.github.io/blogs/rethinking-llm-serving-with-jev/"
hn_url: "https://news.ycombinator.com/item?id=49936590"
title: "Rethinking LLM Serving with System One Models"
article_title: "Rethinking LLM Serving with System One Models"
image: "https://supercomputing-system-ai-lab.github.io/blogs/rethinking-llm-serving-with-jev/og.png"
author: "matt_d"
captured_at: "2026-10-02T19:01:16Z"
capture_tool: "hn-digest"
hn_id: 49936590
score: 1
comments: 0
posted_at: "2026-10-02T18:02:48Z"
tags:
  - hacker-news
---

# Rethinking LLM Serving with System One Models

- HN: [49936590](https://news.ycombinator.com/item?id=49936590)
- Source: [supercomputing-system-ai-lab.github.io](https://supercomputing-system-ai-lab.github.io/blogs/rethinking-llm-serving-with-jev/)
- Score: 1
- Comments: 0
- Posted: 2026-10-02T18:02:48Z

## Translation

Title: Rethinking LLM Serving with System One Models
Description: Opportunities and open problems for fast decision models like Jev, measured with JevServe-Bench.

Article text:
Rethinking LLM Serving with System One Models
← SSAIL Blog
Rethinking LLM Serving with System One Models
Opportunities and open problems for fast decision models like Jev, measured with JevServe-Bench .
Xinyu Lian , Banghao Chi, Jiahuan Yu and Minjia Zhang
SSAIL Lab, University of Illinois Urbana-Champaign · About 12 minutes to read
An LLM service makes many small decisions about every request, most of them with fixed rules today. JevServe-Bench measures what a decision model adds: 16 decisions and 10 models.
Jev is stronger at judging content (guardrails, verification, agent actions) and weaker at predicting what other systems will do (routing, caching).
Jev can also improve the serving engine itself. Its output-length predictions let vLLM serve short requests first: the share of requests meeting the SLO rises from 75% to 99% , and the P90 TTFT falls from 8.3 s to 0.17 s .
Every request is a stack of decisions
Follow one request through an LLM service. Before a model executes, the service has to decide whether a cached answer can serve it, which model should answer and whether that model should think first, whether to call a tool or ask the user for missing details, and where the request goes in the queue.
Most of these decisions are made by rules today: a keyword list, an embedding-similarity threshold, one model per product, first come first served. Rules are fast and predictable, and they treat every request the same way. The alternative is to ask an LLM, which can be accurate but adds a full generation to the path of every request. For many of these decisions that costs more than the decision is worth.
A System One model sits between the two. It is meant to be cheap enough to ask about every request and good enough to act on.
The name borrows Kahneman's two systems of thought: System 2 is slow and deliberate, System 1 is fast and intuitive. An LLM that reasons step by step works like System 2. A System One model gives the fast judgment. You send it a piece of text and a few typed questions, and it answers each one with probabilities: a choice among options, a level on a scale, or a yes or no.
Several properties make this useful inside a serving system:
It is fast. Jev 1.13 answers in a median of 92 ms through OpenRouter, network included. Open-weight Jev-compatible models from 4B to 35B parameters answer in 51 to 284 ms on one H200.
It writes nothing. An answer is one forward pass over the request, with no output tokens to generate or pay for, so its cost does not grow with the length of an answer.
Several questions in one call. A request can carry many questions at once.
Its answers are typed probabilities. Every answer is a probability over a fixed set of options, never free text to parse. A system can rank requests by it, set a threshold, or feed it into a small regression.
Where decisions live in a serving stack
A serving system is a stack of layers, and each makes its own decisions. At the global level, the system decides how much capacity to provision and where each model runs. The gateway sees a request before any model does: it guards, routes, caches and decides on actions. At the cluster level, the system decides which replica of an engine takes a request and when to add or remove GPUs. Inside an engine, it batches requests on the GPU and decides which runs next and which waits. It also manages memory: which conversations keep their KV cache, which are evicted or moved to CPU memory, and which request to pause when memory runs out.
The layers differ in how long a decision can take. Gateway decisions can use a decision model today, because a 90 ms check in front of a call that takes seconds is easy to afford. Engine decisions have tighter budgets: a prediction made when a request is admitted can wait 90 ms, but a choice made at every decoding step, every 10 to 50 ms, cannot. The table lists what a decision model would predict for each decision.
Request
GATEWAY · TODAY
Guard attack?
Route which model?
Cache reuse answer?
Act call a tool?
CLUSTER · NEXT
Replica which copy?
Scale add GPUs?
ENGINE · NEXT
Queue who runs next?
KV cache keep or evict?
Reply
System One model (Jev) questions in, probabilities out, ~90 ms
at admission
Figure 1. Where a decision model can answer questions about a request. The gateway decisions can use it today. Replica routing and autoscaling in the cluster, and queue order and the KV cache in the engine, can use predictions made when a request arrives, such as its output length; decisions made at every decoding step need something faster.
JevServe-Bench: one decision per task
To measure the opportunity and understand the space, we built JevServe-Bench. Each task is one decision a serving system makes, built from public data. The questions the decision model answers and the rule that turns its answers into an action are fixed, so the model is the only thing that changes between entries.
Every task is scored from 0 to 100. Zero is what the system does without a decision model: its default rule, or chance. One hundred is an oracle that knows the outcome in advance. In the semantic-cache task, for example, 0 is the best embedding-similarity threshold and 100 is a cache that knows which requests are equivalent. A score is the share of that gap the decision model closes.
Version 0.2 has 16 tasks in six categories and 10 models: hosted Jev 1.13, eight open-weight Jev-style models, and Qwen3-8B with no decision training as a baseline. Each score comes with a 95% interval, and ranks compare models on the same resamples. A full run takes about ten minutes per model. The leaderboard lists every task and score.
What Jev is good at, and where it falls short
Jev 1.13 ranks first, with an average of 63.6 (95% interval 62.8 to 64.6). The best open model, Kev-27B, scores 62.1 , and untrained Qwen3-8B scores 46.5 . The averages hide a sharp split between categories.
Best of the other eight Jev-style models
Qwen3-8B, no decision training
Strong: judging the content in front of it
Jev does best when the decision is a judgment about the text it is given. It scores 96 on prompt injection, 91 on checking the steps of a math solution, 90 on safety, 86 on telling borderline but harmless requests from harmful ones, 85 on spotting risky agent actions and 84 on deciding whether a tool call fits. Decision training matters most here: untrained Qwen3-8B scores 4.5 on agent risk and 52 on the tool gate.
Weak: predicting what another system will do
Jev does worst when the decision depends on how something else will behave. It scores 17 and 23 on the two routing tasks, which ask which model will answer a request correctly, and 27 on the semantic cache, where it must judge whether an earlier answer fully answers a new request. On the prefix cache, which keeps a conversation's KV in memory if the user is likely to return, Jev's signal is weaker than simply counting the turns so far (AUROC 0.60 against 0.68). We list that task as an open problem and leave it out of the average.
Our reading is that judging content is close to what these models are trained for, while predicting another system's behavior needs data from that system. That is where training on serving logs, simulators and environments could help most.
Inside the engine: predicting output length as the first step
The gateway is where a decision model can help today. Its low latency also makes it a candidate for decisions inside the engine, and output length is the clearest example. A request's output length decides how long it occupies a slot in the batch and how much KV memory it grows into. A scheduler that knows which requests are short can serve them first, so they stop waiting behind long ones. vLLM serves defaultly requests first come, first served.
We fit a small regression from Jev's answers about a request to the length of Qwen3-8B's reply, on 14,000 ShareGPT requests, and test it on 6,000 others. Ranking quality, measured by Kendall's τ, is 0.61 with thinking off and 0.56 with thinking on. The median prediction is off by a factor of 1.3, and 81% to 85% of predictions are within a factor of two.
Reply length is partly random: the same model asked the same question twice gives replies of different lengths. Using one reply's length to predict another's gives τ = 0.79 and 0.73, which is about as well as any prediction from the prompt can do. Jev reaches about three quarters of that. Prompt length alone reaches almost nothing.
Qwen3-8B, no decision training
Ceiling: a second reply's length
A performance benefit check on real hardware
We ran the eval on a real server: Qwen3-8B with vLLM 0.30 on H200. Each run replays 2,000 ShareGPT requests, at Poisson arrivals between 1.05 and 1.54 requests per second. Every policy sees the same arrival times.
We compared four ways to order the queue: vLLM's default (first come, first served), shortest prompt first, which uses no prediction at all, Jev, and the true length. The last three use vLLM's priority scheduler, smallest predicted length first. Each request waits Jev's 92 ms for its prediction before it is sent.
vLLM's default drops below the target just above 1.10 requests per second, to 75% at 1.30 and 42% at 1.42. Serving short prompts first, with no prediction at all, still meets it at 1.42 (92.5%). The reason is how vLLM admits requests: it takes the head of the queue only if its whole prompt fits in free KV memory, and it stops at the first one that does not. Under first come, first served, one long request at the head blocks every request behind it.
Predicting the output length adds to this, most of all in the tail. At 1.42 requests per second Jev keeps 97.2% of requests within the target and the true length 99.2%. Under shortest prompt first the 90th-percentile time to first token stays low up to 1.42 requests per second, then jumps to 6.1 s at 1.54, while Jev keeps it at 0.23 s and the true length at 0.09 s. Further out in the tail, at 1.30 requests per second the 99th percentile is 52.9 s for shortest prompt first, 2.7 s for Jev and 0.4 s with the true length. Jev pays for its 92 ms decision at light load, where its 90th percentile (0.14 s) is the highest of the three reorderings.
In goodput, the highest rate that meets the target while the server keeps up, vLLM's default reaches 1.13 requests per second and Jev 1.30 to 1.36: a real gain of 15% to 21%, and 16% to 26% on a second arrival draw. For every shortest-first policy the limit comes from starvation before the target is missed: from about 1.36 requests per second, the long requests they defer pile up. Length-aware scheduling that avoids both head-of-line blocking and starvation would push goodput further, and would show more clearly how much a better predictor is worth.
To get a feel for the benchmark's decisions, try the playground . It draws on four of them: spotting prompt injections, telling harmful requests from ones that only sound harmful, deciding whether a cached answer can be reused, and guessing how long a reply will be. Each round is five questions against Jev. After each answer you see the right answer, what Jev and two other models said, and how long you took next to Jev's 92 ms.
On the playground's items, Jev gets 80% of the injection questions right, 87% of the harmful-or-harmless ones, 70% of the cache decisions and 55% of the four-way length guesses. We tried hard to beat it. Matching its score took us roughly a hundred times longer per question.
Each submitted round is stored, and together they will give the first human baseline for these decisions.
Opportunities and open problems
Measure engine decisions per deployment. The same length prediction is worth nothing on one server and 45% more goodput on another, so a single score would mislead. The next version of JevServe-Bench will report serving metrics (SLO attainment, goodput, latency percentiles) for each model, GPU and workload.
Fit the latency budget. Predictions made at admission fit in 100 ms. Decisions made at every deco

[truncated]

## Original Extract

Opportunities and open problems for fast decision models like Jev, measured with JevServe-Bench.

Rethinking LLM Serving with System One Models
← SSAIL Blog
Rethinking LLM Serving with System One Models
Opportunities and open problems for fast decision models like Jev, measured with JevServe-Bench .
Xinyu Lian , Banghao Chi, Jiahuan Yu and Minjia Zhang
SSAIL Lab, University of Illinois Urbana-Champaign · About 12 minutes to read
An LLM service makes many small decisions about every request, most of them with fixed rules today. JevServe-Bench measures what a decision model adds: 16 decisions and 10 models.
Jev is stronger at judging content (guardrails, verification, agent actions) and weaker at predicting what other systems will do (routing, caching).
Jev can also improve the serving engine itself. Its output-length predictions let vLLM serve short requests first: the share of requests meeting the SLO rises from 75% to 99% , and the P90 TTFT falls from 8.3 s to 0.17 s .
Every request is a stack of decisions
Follow one request through an LLM service. Before a model executes, the service has to decide whether a cached answer can serve it, which model should answer and whether that model should think first, whether to call a tool or ask the user for missing details, and where the request goes in the queue.
Most of these decisions are made by rules today: a keyword list, an embedding-similarity threshold, one model per product, first come first served. Rules are fast and predictable, and they treat every request the same way. The alternative is to ask an LLM, which can be accurate but adds a full generation to the path of every request. For many of these decisions that costs more than the decision is worth.
A System One model sits between the two. It is meant to be cheap enough to ask about every request and good enough to act on.
The name borrows Kahneman's two systems of thought: System 2 is slow and deliberate, System 1 is fast and intuitive. An LLM that reasons step by step works like System 2. A System One model gives the fast judgment. You send it a piece of text and a few typed questions, and it answers each one with probabilities: a choice among options, a level on a scale, or a yes or no.
Several properties make this useful inside a serving system:
It is fast. Jev 1.13 answers in a median of 92 ms through OpenRouter, network included. Open-weight Jev-compatible models from 4B to 35B parameters answer in 51 to 284 ms on one H200.
It writes nothing. An answer is one forward pass over the request, with no output tokens to generate or pay for, so its cost does not grow with the length of an answer.
Several questions in one call. A request can carry many questions at once.
Its answers are typed probabilities. Every answer is a probability over a fixed set of options, never free text to parse. A system can rank requests by it, set a threshold, or feed it into a small regression.
Where decisions live in a serving stack
A serving system is a stack of layers, and each makes its own decisions. At the global level, the system decides how much capacity to provision and where each model runs. The gateway sees a request before any model does: it guards, routes, caches and decides on actions. At the cluster level, the system decides which replica of an engine takes a request and when to add or remove GPUs. Inside an engine, it batches requests on the GPU and decides which runs next and which waits. It also manages memory: which conversations keep their KV cache, which are evicted or moved to CPU memory, and which request to pause when memory runs out.
The layers differ in how long a decision can take. Gateway decisions can use a decision model today, because a 90 ms check in front of a call that takes seconds is easy to afford. Engine decisions have tighter budgets: a prediction made when a request is admitted can wait 90 ms, but a choice made at every decoding step, every 10 to 50 ms, cannot. The table lists what a decision model would predict for each decision.
Request
GATEWAY · TODAY
Guard attack?
Route which model?
Cache reuse answer?
Act call a tool?
CLUSTER · NEXT
Replica which copy?
Scale add GPUs?
ENGINE · NEXT
Queue who runs next?
KV cache keep or evict?
Reply
System One model (Jev) questions in, probabilities out, ~90 ms
at admission
Figure 1. Where a decision model can answer questions about a request. The gateway decisions can use it today. Replica routing and autoscaling in the cluster, and queue order and the KV cache in the engine, can use predictions made when a request arrives, such as its output length; decisions made at every decoding step need something faster.
JevServe-Bench: one decision per task
To measure the opportunity and understand the space, we built JevServe-Bench. Each task is one decision a serving system makes, built from public data. The questions the decision model answers and the rule that turns its answers into an action are fixed, so the model is the only thing that changes between entries.
Every task is scored from 0 to 100. Zero is what the system does without a decision model: its default rule, or chance. One hundred is an oracle that knows the outcome in advance. In the semantic-cache task, for example, 0 is the best embedding-similarity threshold and 100 is a cache that knows which requests are equivalent. A score is the share of that gap the decision model closes.
Version 0.2 has 16 tasks in six categories and 10 models: hosted Jev 1.13, eight open-weight Jev-style models, and Qwen3-8B with no decision training as a baseline. Each score comes with a 95% interval, and ranks compare models on the same resamples. A full run takes about ten minutes per model. The leaderboard lists every task and score.
What Jev is good at, and where it falls short
Jev 1.13 ranks first, with an average of 63.6 (95% interval 62.8 to 64.6). The best open model, Kev-27B, scores 62.1 , and untrained Qwen3-8B scores 46.5 . The averages hide a sharp split between categories.
Best of the other eight Jev-style models
Qwen3-8B, no decision training
Strong: judging the content in front of it
Jev does best when the decision is a judgment about the text it is given. It scores 96 on prompt injection, 91 on checking the steps of a math solution, 90 on safety, 86 on telling borderline but harmless requests from harmful ones, 85 on spotting risky agent actions and 84 on deciding whether a tool call fits. Decision training matters most here: untrained Qwen3-8B scores 4.5 on agent risk and 52 on the tool gate.
Weak: predicting what another system will do
Jev does worst when the decision depends on how something else will behave. It scores 17 and 23 on the two routing tasks, which ask which model will answer a request correctly, and 27 on the semantic cache, where it must judge whether an earlier answer fully answers a new request. On the prefix cache, which keeps a conversation's KV in memory if the user is likely to return, Jev's signal is weaker than simply counting the turns so far (AUROC 0.60 against 0.68). We list that task as an open problem and leave it out of the average.
Our reading is that judging content is close to what these models are trained for, while predicting another system's behavior needs data from that system. That is where training on serving logs, simulators and environments could help most.
Inside the engine: predicting output length as the first step
The gateway is where a decision model can help today. Its low latency also makes it a candidate for decisions inside the engine, and output length is the clearest example. A request's output length decides how long it occupies a slot in the batch and how much KV memory it grows into. A scheduler that knows which requests are short can serve them first, so they stop waiting behind long ones. vLLM serves defaultly requests first come, first served.
We fit a small regression from Jev's answers about a request to the length of Qwen3-8B's reply, on 14,000 ShareGPT requests, and test it on 6,000 others. Ranking quality, measured by Kendall's τ, is 0.61 with thinking off and 0.56 with thinking on. The median prediction is off by a factor of 1.3, and 81% to 85% of predictions are within a factor of two.
Reply length is partly random: the same model asked the same question twice gives replies of different lengths. Using one reply's length to predict another's gives τ = 0.79 and 0.73, which is about as well as any prediction from the prompt can do. Jev reaches about three quarters of that. Prompt length alone reaches almost nothing.
Qwen3-8B, no decision training
Ceiling: a second reply's length
A performance benefit check on real hardware
We ran the eval on a real server: Qwen3-8B with vLLM 0.30 on H200. Each run replays 2,000 ShareGPT requests, at Poisson arrivals between 1.05 and 1.54 requests per second. Every policy sees the same arrival times.
We compared four ways to order the queue: vLLM's default (first come, first served), shortest prompt first, which uses no prediction at all, Jev, and the true length. The last three use vLLM's priority scheduler, smallest predicted length first. Each request waits Jev's 92 ms for its prediction before it is sent.
vLLM's default drops below the target just above 1.10 requests per second, to 75% at 1.30 and 42% at 1.42. Serving short prompts first, with no prediction at all, still meets it at 1.42 (92.5%). The reason is how vLLM admits requests: it takes the head of the queue only if its whole prompt fits in free KV memory, and it stops at the first one that does not. Under first come, first served, one long request at the head blocks every request behind it.
Predicting the output length adds to this, most of all in the tail. At 1.42 requests per second Jev keeps 97.2% of requests within the target and the true length 99.2%. Under shortest prompt first the 90th-percentile time to first token stays low up to 1.42 requests per second, then jumps to 6.1 s at 1.54, while Jev keeps it at 0.23 s and the true length at 0.09 s. Further out in the tail, at 1.30 requests per second the 99th percentile is 52.9 s for shortest prompt first, 2.7 s for Jev and 0.4 s with the true length. Jev pays for its 92 ms decision at light load, where its 90th percentile (0.14 s) is the highest of the three reorderings.
In goodput, the highest rate that meets the target while the server keeps up, vLLM's default reaches 1.13 requests per second and Jev 1.30 to 1.36: a real gain of 15% to 21%, and 16% to 26% on a second arrival draw. For every shortest-first policy the limit comes from starvation before the target is missed: from about 1.36 requests per second, the long requests they defer pile up. Length-aware scheduling that avoids both head-of-line blocking and starvation would push goodput further, and would show more clearly how much a better predictor is worth.
To get a feel for the benchmark's decisions, try the playground . It draws on four of them: spotting prompt injections, telling harmful requests from ones that only sound harmful, deciding whether a cached answer can be reused, and guessing how long a reply will be. Each round is five questions against Jev. After each answer you see the right answer, what Jev and two other models said, and how long you took next to Jev's 92 ms.
On the playground's items, Jev gets 80% of the injection questions right, 87% of the harmful-or-harmless ones, 70% of the cache decisions and 55% of the four-way length guesses. We tried hard to beat it. Matching its score took us roughly a hundred times longer per question.
Each submitted round is stored, and together they will give the first human baseline for these decisions.
Opportunities and open problems
Measure engine decisions per deployment. The same length prediction is worth nothing on one server and 45% more goodput on another, so a single score would mislead. The next version of JevServe-Bench will report serving metrics (SLO attainment, goodput, latency percentiles) for each model, GPU and workload.
Fit the latency budget. Predictions made at admission fit in 100 ms. Decisions made at every deco

[truncated]
