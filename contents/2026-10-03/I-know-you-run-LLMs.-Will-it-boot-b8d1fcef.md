---
source: "https://mondegreens.github.io/apron/blog/2026/10/02/i-know-you-run-models-will-it-boot/"
hn_url: "https://news.ycombinator.com/item?id=49946580"
title: "I know you run LLMs. Will it boot?"
article_title: "I know you run LLMs. Will it boot? - Apron"
image: "https://mondegreens.github.io/apron/blog/img/will-it-boot/will-it-boot-header-1200x630.jpg"
author: "vladryzhkov"
captured_at: "2026-10-03T19:02:05Z"
capture_tool: "hn-digest"
hn_id: 49946580
score: 2
comments: 0
posted_at: "2026-10-03T18:32:21Z"
tags:
  - hacker-news
---

# I know you run LLMs. Will it boot?

- HN: [49946580](https://news.ycombinator.com/item?id=49946580)
- Source: [mondegreens.github.io](https://mondegreens.github.io/apron/blog/2026/10/02/i-know-you-run-models-will-it-boot/)
- Score: 2
- Comments: 0
- Posted: 2026-10-03T18:32:21Z

## Translation

Title: I know you run LLMs. Will it boot?
Article title: I know you run LLMs. Will it boot? - Apron
Description: Fourteen open models, from one RTX 4090 to eight H200s: Apron predicted their memory before renting the GPU, then booted each one on vLLM and measured.

Article text:
I know you run LLMs. Will it boot? - Apron
Skip to content
Apron
I know you run LLMs. Will it boot?
Search
mondegreens/apron
Apron
mondegreens/apron
Home
Fingerprint Migration (Phase 1b)
Findings
Findings
Findings
Findings
On this page
Read this first: what this does not show
Predicted before renting, then measured
Attention layout, not model size, decides how many tokens fit
Where the calculator was wrong: the startup peak
New layouts, counted from the engine's own code
Time to first token: the typical request and the slowest one
Answers expire every two weeks
Categories
Categories
Findings
Architecture Decision Records
Architecture Decision Records
ADR-001 Execution Environment
ADR-003 Engine-Native Computation
ADR-005 Delivery Restructuring
ADR-007 Modality-Agnostic Specs
Read this first: what this does not show
Predicted before renting, then measured
Attention layout, not model size, decides how many tokens fit
Where the calculator was wrong: the startup peak
New layouts, counted from the engine's own code
Time to first token: the typical request and the slowest one
Answers expire every two weeks
Back to index
Metadata
October 2, 2026
I know you run LLMs. Will it boot? ¶
I know you run models. A new one comes out, everyone on your timeline is
running it, and you rent an H100 to try. Forty minutes later you still do not
know whether it fits on that card, and if it does, with what flags.
So I ran them: fourteen of the open models people run this month, from one RTX
4090 to eight H200s, and I kept every boot and every failure, so you do not
have to rent a GPU to find out.
This post is what I learned, and the tool that did it: Apron, an open-source
tool that answers "will it fit, will it boot, will it answer" before you rent
the GPU, then checks its own answer on a real one and keeps both. It is also
about what I got wrong along the way, and how I found out.
In short: before renting anything, Apron predicted each model's weight memory
from its config, its safetensors headers and the pinned vLLM release's own
source. For 11 of 14 models the prediction landed within 2% of what vLLM
measured on the GPU. The three misses are explained below, and two are fixed.
Four models broke on the way: two hybrids that vLLM's defaults could not boot
on an H100, one that booted and answered in the wrong channel, and one that
crashed on a tool call with CUDA graphs. Each fix was proven by a second boot.
A model that boots is not a model that answers, and a slow p99 here is usually
one or two slow requests. The sections below say which is which.
Read this first: what this does not show ¶
There is one measured boot per model. Fourteen models on four hardware setups,
most of them datacenter cards: one RTX 4090, seven single H100s, four models on
four H200s and two on eight. An earlier cohort of smaller models adds five GPU
types. That is evidence, not a statistic. Consumer cards are next.
The prompts are short. The single-GPU boots ran with --max-model-len 640 and
the multi-GPU boots with 32,768. The serving test is 50 requests of 512 tokens
in and 128 out, four at a time. This shows whether a model boots and answers,
not how it performs at long context.
The task checks are correctness checks, not a benchmark. They show that the
endpoint answers; they say nothing about which model is better. Nothing here is
a ranking: the hardware differs from group to group.
Some of the startup-peak rule comes from observation (details below). The
four-GPU models whose startup peak and CUDA-graph estimate the calculator still
under-predicts (by up to 2.2 and 2.0 GiB) are kept apart until it explains
them.
As for the KV budget: with today's calculator, the predicted KV budget is
within the planner's safety buffer (plus a compile segment vLLM may hold) on
every single- and four-GPU record kept, and a test checks it. The buffer
(2.19 GiB) is the largest error measured on those records, not a tuned margin.
It was not always so at run time: gpt-oss-20b's KV cache came out 2.3 GiB
smaller than predicted, from the weight miss explained below.
The command-line tool is not ready for outside use yet. The records and the
rules are in the repository.
Deploying an open model on your own GPU is a three-way bet: the checkpoint, the
engine version and the hardware. Each of the three moves.
The engine moves every two weeks. vLLM v0.24 came out on June 29 and v0.30 on
September 22. Its maintainers are candid about how fast things change: they
turned down a hand-maintained table of model capabilities because such a table
"drifts", and kept reading capabilities from the code
( #52459 ), and they now run
nightly accuracy and performance checks
on 17 model-and-hardware recipes on datacenter GPUs.
The models move about monthly, each with a new attention layout: Mamba and
gated delta-net hybrids, DeepSeek V4's compressed KV cache, GLM-5's linear plus
sparse attention. Each one pages memory differently. Ten of the last
twenty-five posts on the vLLM blog are about getting one particular model to
run well.
The hardware is consumer, professional and datacenter cards, rented by the
hour, with memory that is never quite the number on the box. An H100 80GB gives
PyTorch 85,017,493,504 bytes, not 85,899,345,920.
The tools around this are good and getting better. vLLM refuses many bad
configurations before loading and tells you what to change. NVIDIA's
aiconfigurator and llm-d's planner estimate memory and performance for common
attention types; aiconfigurator now lists hybrid models too. Hugging Face shows
whether GGUF and MLX files fit your hardware. None of them gives the joined
answer I kept needing: this exact checkpoint, on this exact vLLM version, on
this exact GPU, predicted before renting and then checked.
That need is on record. vLLM issue
#29325 asked for a
preflight "dry run". In the discussion, a contributor from llm-d (IBM
Research) suggested building the planner in a separate repository, "so we can
move fast and quickly adapt to vLLM changes". The issue was closed as not
planned in July. The planners say it themselves:
aiconfigurator's README still notes that memory estimation "needs to be
studied more", and one of its issues
( #1396 ) documents a
sweep that skipped the KV budget check, over-predicted throughput 4.6× and
recommended a configuration that did not deploy. Its maintainers have since
fixed the budget check, and their own diagnosis in that thread is the best
summary of the problem: the KV pool is "a small difference of two large
numbers", so a 1% error in the rest "is an ~18% error in admissible batch". I
don't say that to knock them: every estimate looks like this until it is
checked against a boot.
Apron is a small loop with a strict rule about what it is allowed to call true.
Plan without a GPU. Read the checkpoint's config and safetensors headers
from the Hub (headers only, no weights), and the facts of the pinned vLLM
version: which architectures it registers, which tensors it loads and which
it skips, how it pages KV cache for each layer type, its default batch
sizes. Compute weights, KV cache per sequence, the startup activation peak
and the CUDA graph cost per GPU.
Say "unknown" out loud. If a model uses an attention layout the calculator
does not model, the answer is "unknown", not an estimate borrowed from a
different layout.
Boot and measure. Rent the GPU, download the weights on the machine, boot
the exact plan, read vLLM's own memory accounting, run a small set of checks
and a serving test, shut the machine down.
Keep the prediction beside the measurement. Every record is
content-addressed and stores what was predicted next to what was measured,
with the engine image digest and the hardware it ran on.
Prove fixes. When a boot fails, the error is fingerprinted and matched to a
correction rule. A rule starts as a hypothesis ; it becomes a proven rule
only when a second boot with the correction succeeds.
Every number Apron shows is labeled by how it is known: derived (computed
from inputs), proven constraint (the pinned engine's source enforces it),
predicted (the calculator modeled it) or measured (the exact execution
observed it). A prediction is never presented as a vLLM test.
Predicted before renting, then measured ¶
The chart at the top is the whole idea in one picture. Each dot is one model:
the weight memory Apron predicted before the GPU was rented, against what vLLM
reported after loading. The prediction is the one recorded before the boot,
not a number recomputed afterwards.
Eleven of fourteen are within 2%. Three are outside it.
gpt-oss-20b on the RTX 4090 came out 7.1% low. Below Hopper, vLLM runs
gpt-oss's MXFP4 experts on Marlin, which pads them from 2880 to 3072 × 2944.
Apron counts that padding now: 13.67 GiB predicted.
GLM-4.7-Flash came out 4.1% high. I counted a multi-token-prediction layer that
vLLM does not load. Without it the prediction is 55.77 GiB against 55.87
measured.
Qwen3.8-Flash-Next on four H200s came out 3.5% low. The boot showed layers vLLM
keeps whole on every GPU instead of splitting them. Apron counts them now
(59.82 GiB predicted), and about 1 GiB per GPU is still not explained.
I picked models by what people actually use and talk about in September 2026
(OpenRouter rankings, provider counts, Hacker News and r/LocalLLaMA threads),
not by download counts. The column most worth copying is the third: what each
model needed beyond vLLM's defaults to boot and answer on that hardware. The
longer notes for each model are at the end of the post.
A few notes on reading it. Single-GPU models answered three short questions.
Multi-GPU models ran five deployment checks: a 32k-token needle, a tool call,
JSON output, an image and reasoning. Text-only models skip the image check, so
their total is four. The multi-GPU rows also set tensor parallelism and the
model's reasoning and tool-call parsers, which the tool-call and reasoning
checks need. For models whose chat template reads reasoning_effort (gpt-oss,
GLM-5.x, DeepSeek V4), the checks send reasoning_effort: low in the request,
so a short answer fits the checks' small token budget. Cold start is the first
boot on a fresh machine, torch.compile included; a warm compile cache is
faster.
After the weights, vLLM takes what it needs to run (the startup activation
peak, CUDA graphs and memory outside PyTorch), leaves 10% of the card alone at
the default --gpu-memory-utilization 0.9 , and gives everything else to the KV
cache. The green segment is the room left for your requests. On one H100 that
is between 5.7 and 32.4 GiB, and it is mostly what the weights leave: the
34 GiB Qwen3.6 FP8 leaves 32.4 GiB, the 61 GiB gpt-oss-120b leaves 5.7. How
many tokens fit in that room is a different question, and the next section
shows it depends on the attention layout.
Attention layout, not model size, decides how many tokens fit ¶
Gemma 4 31B and GLM-4.7-Flash had about the same KV memory on the same H100
(8.6 and 8.5 GiB). Gemma fits 9,986 tokens in it; GLM-4.7-Flash fits 169,216.
Gemma's global layers use 512-wide heads, so each token is expensive. GLM's
multi-head latent attention stores a compressed cache, so each token is cheap.
At long context Gemma 4 is KV-bound on one H100.
The hybrids are shown apart because their pool also holds a fixed state per
request (Mamba or gated delta-net), so their tokens per GiB are not comparable
with the others token for token. The number of concurrent long requests a pool
supports is roughly the pool divided by your context length, and less than
that for hybrids.
Where the calculator was wrong: the startup peak ¶
vLLM decides how much KV cache it can allocate by running one dummy batch at
startup and measuring the peak memory PyTorch allocated during it. Whatever
that peak is comes straight out of your KV budget. Apron predicts it, and until
this week its rule was simple and, on fifteen measured points from an earlier
cohort of smaller models on five GPU types, right to within 0.02 GiB: during
the first, c

[truncated]

## Original Extract

Fourteen open models, from one RTX 4090 to eight H200s: Apron predicted their memory before renting the GPU, then booted each one on vLLM and measured.

I know you run LLMs. Will it boot? - Apron
Skip to content
Apron
I know you run LLMs. Will it boot?
Search
mondegreens/apron
Apron
mondegreens/apron
Home
Fingerprint Migration (Phase 1b)
Findings
Findings
Findings
Findings
On this page
Read this first: what this does not show
Predicted before renting, then measured
Attention layout, not model size, decides how many tokens fit
Where the calculator was wrong: the startup peak
New layouts, counted from the engine's own code
Time to first token: the typical request and the slowest one
Answers expire every two weeks
Categories
Categories
Findings
Architecture Decision Records
Architecture Decision Records
ADR-001 Execution Environment
ADR-003 Engine-Native Computation
ADR-005 Delivery Restructuring
ADR-007 Modality-Agnostic Specs
Read this first: what this does not show
Predicted before renting, then measured
Attention layout, not model size, decides how many tokens fit
Where the calculator was wrong: the startup peak
New layouts, counted from the engine's own code
Time to first token: the typical request and the slowest one
Answers expire every two weeks
Back to index
Metadata
October 2, 2026
I know you run LLMs. Will it boot? ¶
I know you run models. A new one comes out, everyone on your timeline is
running it, and you rent an H100 to try. Forty minutes later you still do not
know whether it fits on that card, and if it does, with what flags.
So I ran them: fourteen of the open models people run this month, from one RTX
4090 to eight H200s, and I kept every boot and every failure, so you do not
have to rent a GPU to find out.
This post is what I learned, and the tool that did it: Apron, an open-source
tool that answers "will it fit, will it boot, will it answer" before you rent
the GPU, then checks its own answer on a real one and keeps both. It is also
about what I got wrong along the way, and how I found out.
In short: before renting anything, Apron predicted each model's weight memory
from its config, its safetensors headers and the pinned vLLM release's own
source. For 11 of 14 models the prediction landed within 2% of what vLLM
measured on the GPU. The three misses are explained below, and two are fixed.
Four models broke on the way: two hybrids that vLLM's defaults could not boot
on an H100, one that booted and answered in the wrong channel, and one that
crashed on a tool call with CUDA graphs. Each fix was proven by a second boot.
A model that boots is not a model that answers, and a slow p99 here is usually
one or two slow requests. The sections below say which is which.
Read this first: what this does not show ¶
There is one measured boot per model. Fourteen models on four hardware setups,
most of them datacenter cards: one RTX 4090, seven single H100s, four models on
four H200s and two on eight. An earlier cohort of smaller models adds five GPU
types. That is evidence, not a statistic. Consumer cards are next.
The prompts are short. The single-GPU boots ran with --max-model-len 640 and
the multi-GPU boots with 32,768. The serving test is 50 requests of 512 tokens
in and 128 out, four at a time. This shows whether a model boots and answers,
not how it performs at long context.
The task checks are correctness checks, not a benchmark. They show that the
endpoint answers; they say nothing about which model is better. Nothing here is
a ranking: the hardware differs from group to group.
Some of the startup-peak rule comes from observation (details below). The
four-GPU models whose startup peak and CUDA-graph estimate the calculator still
under-predicts (by up to 2.2 and 2.0 GiB) are kept apart until it explains
them.
As for the KV budget: with today's calculator, the predicted KV budget is
within the planner's safety buffer (plus a compile segment vLLM may hold) on
every single- and four-GPU record kept, and a test checks it. The buffer
(2.19 GiB) is the largest error measured on those records, not a tuned margin.
It was not always so at run time: gpt-oss-20b's KV cache came out 2.3 GiB
smaller than predicted, from the weight miss explained below.
The command-line tool is not ready for outside use yet. The records and the
rules are in the repository.
Deploying an open model on your own GPU is a three-way bet: the checkpoint, the
engine version and the hardware. Each of the three moves.
The engine moves every two weeks. vLLM v0.24 came out on June 29 and v0.30 on
September 22. Its maintainers are candid about how fast things change: they
turned down a hand-maintained table of model capabilities because such a table
"drifts", and kept reading capabilities from the code
( #52459 ), and they now run
nightly accuracy and performance checks
on 17 model-and-hardware recipes on datacenter GPUs.
The models move about monthly, each with a new attention layout: Mamba and
gated delta-net hybrids, DeepSeek V4's compressed KV cache, GLM-5's linear plus
sparse attention. Each one pages memory differently. Ten of the last
twenty-five posts on the vLLM blog are about getting one particular model to
run well.
The hardware is consumer, professional and datacenter cards, rented by the
hour, with memory that is never quite the number on the box. An H100 80GB gives
PyTorch 85,017,493,504 bytes, not 85,899,345,920.
The tools around this are good and getting better. vLLM refuses many bad
configurations before loading and tells you what to change. NVIDIA's
aiconfigurator and llm-d's planner estimate memory and performance for common
attention types; aiconfigurator now lists hybrid models too. Hugging Face shows
whether GGUF and MLX files fit your hardware. None of them gives the joined
answer I kept needing: this exact checkpoint, on this exact vLLM version, on
this exact GPU, predicted before renting and then checked.
That need is on record. vLLM issue
#29325 asked for a
preflight "dry run". In the discussion, a contributor from llm-d (IBM
Research) suggested building the planner in a separate repository, "so we can
move fast and quickly adapt to vLLM changes". The issue was closed as not
planned in July. The planners say it themselves:
aiconfigurator's README still notes that memory estimation "needs to be
studied more", and one of its issues
( #1396 ) documents a
sweep that skipped the KV budget check, over-predicted throughput 4.6× and
recommended a configuration that did not deploy. Its maintainers have since
fixed the budget check, and their own diagnosis in that thread is the best
summary of the problem: the KV pool is "a small difference of two large
numbers", so a 1% error in the rest "is an ~18% error in admissible batch". I
don't say that to knock them: every estimate looks like this until it is
checked against a boot.
Apron is a small loop with a strict rule about what it is allowed to call true.
Plan without a GPU. Read the checkpoint's config and safetensors headers
from the Hub (headers only, no weights), and the facts of the pinned vLLM
version: which architectures it registers, which tensors it loads and which
it skips, how it pages KV cache for each layer type, its default batch
sizes. Compute weights, KV cache per sequence, the startup activation peak
and the CUDA graph cost per GPU.
Say "unknown" out loud. If a model uses an attention layout the calculator
does not model, the answer is "unknown", not an estimate borrowed from a
different layout.
Boot and measure. Rent the GPU, download the weights on the machine, boot
the exact plan, read vLLM's own memory accounting, run a small set of checks
and a serving test, shut the machine down.
Keep the prediction beside the measurement. Every record is
content-addressed and stores what was predicted next to what was measured,
with the engine image digest and the hardware it ran on.
Prove fixes. When a boot fails, the error is fingerprinted and matched to a
correction rule. A rule starts as a hypothesis ; it becomes a proven rule
only when a second boot with the correction succeeds.
Every number Apron shows is labeled by how it is known: derived (computed
from inputs), proven constraint (the pinned engine's source enforces it),
predicted (the calculator modeled it) or measured (the exact execution
observed it). A prediction is never presented as a vLLM test.
Predicted before renting, then measured ¶
The chart at the top is the whole idea in one picture. Each dot is one model:
the weight memory Apron predicted before the GPU was rented, against what vLLM
reported after loading. The prediction is the one recorded before the boot,
not a number recomputed afterwards.
Eleven of fourteen are within 2%. Three are outside it.
gpt-oss-20b on the RTX 4090 came out 7.1% low. Below Hopper, vLLM runs
gpt-oss's MXFP4 experts on Marlin, which pads them from 2880 to 3072 × 2944.
Apron counts that padding now: 13.67 GiB predicted.
GLM-4.7-Flash came out 4.1% high. I counted a multi-token-prediction layer that
vLLM does not load. Without it the prediction is 55.77 GiB against 55.87
measured.
Qwen3.8-Flash-Next on four H200s came out 3.5% low. The boot showed layers vLLM
keeps whole on every GPU instead of splitting them. Apron counts them now
(59.82 GiB predicted), and about 1 GiB per GPU is still not explained.
I picked models by what people actually use and talk about in September 2026
(OpenRouter rankings, provider counts, Hacker News and r/LocalLLaMA threads),
not by download counts. The column most worth copying is the third: what each
model needed beyond vLLM's defaults to boot and answer on that hardware. The
longer notes for each model are at the end of the post.
A few notes on reading it. Single-GPU models answered three short questions.
Multi-GPU models ran five deployment checks: a 32k-token needle, a tool call,
JSON output, an image and reasoning. Text-only models skip the image check, so
their total is four. The multi-GPU rows also set tensor parallelism and the
model's reasoning and tool-call parsers, which the tool-call and reasoning
checks need. For models whose chat template reads reasoning_effort (gpt-oss,
GLM-5.x, DeepSeek V4), the checks send reasoning_effort: low in the request,
so a short answer fits the checks' small token budget. Cold start is the first
boot on a fresh machine, torch.compile included; a warm compile cache is
faster.
After the weights, vLLM takes what it needs to run (the startup activation
peak, CUDA graphs and memory outside PyTorch), leaves 10% of the card alone at
the default --gpu-memory-utilization 0.9 , and gives everything else to the KV
cache. The green segment is the room left for your requests. On one H100 that
is between 5.7 and 32.4 GiB, and it is mostly what the weights leave: the
34 GiB Qwen3.6 FP8 leaves 32.4 GiB, the 61 GiB gpt-oss-120b leaves 5.7. How
many tokens fit in that room is a different question, and the next section
shows it depends on the attention layout.
Attention layout, not model size, decides how many tokens fit ¶
Gemma 4 31B and GLM-4.7-Flash had about the same KV memory on the same H100
(8.6 and 8.5 GiB). Gemma fits 9,986 tokens in it; GLM-4.7-Flash fits 169,216.
Gemma's global layers use 512-wide heads, so each token is expensive. GLM's
multi-head latent attention stores a compressed cache, so each token is cheap.
At long context Gemma 4 is KV-bound on one H100.
The hybrids are shown apart because their pool also holds a fixed state per
request (Mamba or gated delta-net), so their tokens per GiB are not comparable
with the others token for token. The number of concurrent long requests a pool
supports is roughly the pool divided by your context length, and less than
that for hybrids.
Where the calculator was wrong: the startup peak ¶
vLLM decides how much KV cache it can allocate by running one dummy batch at
startup and measuring the peak memory PyTorch allocated during it. Whatever
that peak is comes straight out of your KV budget. Apron predicts it, and until
this week its rule was simple and, on fifteen measured points from an earlier
cohort of smaller models on five GPU types, right to within 0.02 GiB: during
the first, c

[truncated]
