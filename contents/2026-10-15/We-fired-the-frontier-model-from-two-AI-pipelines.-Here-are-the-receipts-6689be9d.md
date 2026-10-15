---
source: "https://seldon-ai.com/blog/we-fired-the-frontier-model"
hn_url: "https://news.ycombinator.com/item?id=50110715"
title: "We fired the frontier model from two AI pipelines. Here are the receipts"
article_title: "We fired the frontier model from two AI pipelines. Here are the receipts. | Seldon"
image: "https://seldon-ai.com/blog/we-fired-the-frontier-model/opengraph-image"
author: "nlpnerd"
captured_at: "2026-10-15T21:58:07Z"
capture_tool: "hn-digest"
hn_id: 50110715
score: 1
comments: 0
posted_at: "2026-10-15T21:40:18Z"
tags:
  - hacker-news
---

# We fired the frontier model from two AI pipelines. Here are the receipts

- HN: [50110715](https://news.ycombinator.com/item?id=50110715)
- Source: [seldon-ai.com](https://seldon-ai.com/blog/we-fired-the-frontier-model)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T21:40:18Z

## Translation

Title: We fired the frontier model from two AI pipelines. Here are the receipts
Article title: We fired the frontier model from two AI pipelines. Here are the receipts. | Seldon
Description: Seldon Functions compiled the LLM call at the center of two hard data pipelines. Agent-trace extraction got 5× cheaper and more accurate; PII detection got ~8,000× cheaper and up to 10× faster. Here are the production receipts, including what broke.

Article text:
seldon How it works Models Docs Blog v0 Log in Sign up ← Blog October 14, 2026 21 min read Seldon We fired the frontier model from two AI pipelines. Here are the receipts.
Seldon Functions compiles the LLM call in the middle of your data pipeline into a faster, cheaper program, and proves it against the model you trust. Two production experiments: 5× cheaper and more accurate on agent traces. ~8,000× cheaper and up to 10× faster on PII detection.
Routers pick a model. Seldon builds a program. We pointed it at two hard, high-volume LLM workloads, and here's what came back.
Agent trace analytics: 5× cheaper, more accurate and faster. Seldon replaced a frontier-model call with a compiled program. The program passed 95.0% of held-out checks where the frontier model passed 92.5%, cost 80% less per request and ran 21% faster. Its output also clustered better downstream, at 0.56 ARI against 0.33.
PII detection: ~8,000× cheaper, up to 10× faster, every document processed. Seldon compiled the LLM out entirely : a specialized span model plus deterministic code. It handled all 66 test documents for $0.00014 in total , including 8 that the frontier-model baseline failed on. It isn't perfect. It missed some bare first names, and the receipts below show exactly where.
Zero fallbacks, a one-line switch, and about $15 to compile. Both compiled programs served every fresh input on their own primary path. Each pays for itself in roughly 1,000 to 2,300 calls.
Switching takes one line. Seldon gives you a new model ID, and the rest of your code doesn't change:
- model="gpt-5.6-sol"
+ model="seldon/function/<FUNCTION_ID>@<REVISION_ID>"
Why it works
Most LLM calls in production aren't chat. They're functions : text goes in, JSON with a fixed schema comes out, millions of times. You pay frontier-model prices on every call because a frontier model is what worked when you wrote the prompt. The usual alternatives are shortcuts: sample less, truncate more, or swap in the cheapest model and hope. Almost nobody measures what those shortcuts cost in quality.
Seldon turns that guess into a measurement:
Your model sets the bar. You bring the prompt, the schema and a reference model you trust. Every alternative is scored against that reference, on your contract.
It searches whole programs, not just models. Smaller LLMs, specialized models (GLiNER for spans, Jev for decisions) and deterministic kernels, combined into multi-step pipelines. That's how a PII detector ends up with no LLM on its main path.
You see the trade before you make it. A comparison report puts up to three alternatives next to your reference, with quality, critical failures, cost, p50/p95 and fallback rate for each. You choose, and nothing activates on its own.
It ships with a safety net. Outputs are validated at runtime, and your original call stays wired in as a fallback. The endpoint is OpenAI-compatible, so the OpenAI SDK, LiteLLM, LangChain and the Vercel AI SDK work unchanged.
We built both experiments as standalone scripts and ran them against production api.seldon-ai.com , with real models and real wallet charges. Each compiled Function went head to head with the simple baseline (the original prompt on GPT-5.6 Sol) on the same held-out inputs. We also ran the tools you'd otherwise reach for: regex, Presidio + GLiNER, hand-written features and LLM summaries. We published the failures too.
Where these experiments come from. PostHog's public write-up on LLM trace clustering inspired the trace experiment. The anonymization needs that companies licensing enterprise data, like Mercor , describe publicly inspired the PII experiment. We designed both setups, built the data and ran everything ourselves. We don't represent either company, neither was involved or endorsed this, and nothing here measures, reproduces or challenges their systems.
The receipts: Seldon against the simple baseline
The simple baseline is what most teams ship first: send the original prompt to a frontier model (GPT-5.6 Sol) for every record. Each Seldon pipeline below sits side by side with that baseline, on the same inputs. Bold marks the better result in each row.
Trace signals (experiment inspired by PostHog's clustering write-up)
PII identification (experiment inspired by enterprise data-anonymization needs)
Held-out reports use Seldon's own synthetic examples, an independent judge and serial timing. External tests call the production gateway, so their latencies include network and gateway overhead. Seldon costs are settled receipts. Sol costs in the external tests are token-price estimates at $5 input / $30 output per million tokens. Both baselines are our own Sol calls. None of these numbers describe PostHog's or Mercor's systems.
Put plainly, Seldon beat the baseline on every trace row that differs. On PII it's far cheaper, several times faster, and finished every document. But it leaked names the baseline caught, and a local Presidio + GLiNER setup beat it on recall (see below). One result is a clear win. The other is a tradeoff you'd want to see before making it. Seldon shows you both.
The lifecycle has four explicit steps, and nothing is activated unless you activate it.
Create. Upload the call's prompt, input/output JSON schemas and reference model. The trace demo's contract is {"text": "..."} → {intent, actions[], outcome, failure_signals[], evidence[]} .
Preview (paid, capped). Seldon generates synthetic examples for the task, runs your reference on them, infers an evaluator and returns a quote: forecast cost, forecast minutes and which components are eligible. The server enforces your cap, currently up to $100.
Optimize. Bounded search over complete programs: up to 64 candidates, 16 steps and 3 learned nodes. Building blocks:
smaller LLMs (e.g. GPT-5.6 Luna)
specialized models: fastino/gliner2.5-multi-v1 for spans and records, typesafe/jev for boolean/enum/pair decisions
deterministic kernels: exact span alignment, JSON-pointer record projection, lookup, arithmetic and constraint checks
the reference itself, as an optional fallback
Small search spaces are screened exhaustively. Larger ones use SMAC random-forest Bayesian optimization over whole-program evaluations.
Compare and activate. The run publishes a comparison report with up to three executable alternatives beside the reference. Each one comes with quality, critical failures, full-path cost, p50/p95, fallback rate and coverage. You choose. An alternative can't be activated without proof that its primary path executes; a candidate that only ever succeeds through fallback is never callable.
The compiler only ever sees your prompt, your schemas, one training example and the synthetic data it generates itself. In both demos, the held-out documents and traces, and every gold label, stayed on the demo side. None of them shaped the search.
Then you call it like any other model:
from openai import OpenAI
import json
client = OpenAI(base_url="https://api.seldon-ai.com/v1", api_key=SELDON_API_KEY)
resp = client.chat.completions.create(
model=f"seldon/function/{FUNCTION_ID}@{REVISION_ID}",
messages=[{"role": "user", "content": json.dumps({"text": trace_fragment})}],
)
signals = json.loads(resp.choices[0].message.content)
It's the same OpenAI-compatible endpoint, so the OpenAI SDK, LiteLLM, LangChain and the Vercel AI SDK work unchanged. Every call writes an execution receipt with its real charge.
Where Seldon stops and your code starts
In both demos, Seldon replaces exactly one step: the LLM call your prompt used to make. Everything around it is ordinary Python in a standalone demo.py on your machine. That covers loading data, chunking, checking outputs, redaction, embeddings, clustering and scoring. The demo doesn't import the Seldon package, and it calls the Function over plain HTTP.
The diagrams below use one color code throughout:
Purple: hosted by Seldon, inside the compiled Function.
Grey: the demo script, i.e. your application code.
Green: local models running on your own CPU.
Orange: a direct call to the reference model (Sol) as a baseline, outside any Function.
Dashed: the reference fallback. It's wired in, but it never fired on the fresh inputs in these runs.
Every baseline goes through the same script. Only the one step that the Function replaces gets swapped out, so each comparison measures that step and nothing else.
PostHog has published a detailed walkthrough of how it built LLM trace clustering , with links to its code. It's one of the clearest public explanations of this problem, and it shaped how we set up our experiment. As they describe it:
each trace is rendered as line-numbered text
huge traces are uniformly downsampled to fit context
a sample of traces is taken every hour
each sampled trace gets a structured summary from GPT-4.1 nano
the summaries are embedded with text-embedding-3-large
the embeddings go through UMAP and HDBSCAN
a GPT-5.2 agent names each cluster
It's a thoughtful design, and the post is candid about tradeoffs that anyone building trace analytics runs into:
Cost forces sampling. Summarizing every trace would be expensive. So they sample a subset each hour, use a small, fast model, and plan to move to the Batch API.
Downsampling loses information. Dropping lines from huge traces keeps them inside context limits, at what they call "the cost of some accuracy/quality".
One schema for everyone. A single general-purpose summary schema can't fit every kind of trace. Letting users bring their own prompts and structured outputs is listed as a possible future improvement.
Structure beats prose. They summarize into a structured schema, with line references, specifically because it embeds better than free text.
As described, one step in that design is a function-shaped LLM call:
An illustration drawn from PostHog's published description, not from their code. Purple, thick border: the kind of step Seldon Functions is designed for. Dotted grey: an agent loop, which Functions doesn't compile.
In a design like this, the summary step is a fixed prompt plus a fixed schema, run on every sampled trace. That's the kind of call a compiler can work on, and it's the step our experiment focuses on. The labeling agent is a multi-step loop that uses tools, and Functions doesn't touch those. Embeddings, UMAP, HDBSCAN and orchestration stay in your code.
We borrowed the overall shape and changed the parts we wanted to test. This is our own setup, not a reproduction of PostHog's system, and we didn't run anything against it. The left column is only there to show where our choices differ from the design that inspired them.
Our extraction contract asks for the user's intent, the ordered actions, the observed outcome ( completed / failed / unclear ) and any failure signals. Each signal must be backed by evidence quoted exactly from the trace, and the script checks every quote in code.
Before the call, the script drops rewards, task IDs and simulator state from every trace. Tasks are hash-partitioned, so repeated attempts at one task can't leak between train and test. After the call, the whole trace is rejected if any chunk's evidence isn't an exact quote. Task IDs are only used at the very end, for scoring.
The fallback has narrow triggers. A step's output has to fail schema or grounding validation, or serving admission has to refuse the request before anything is sent. Then Seldon runs your original Sol call inside the same 55-second deadline. If a provider call's outcome is uncertain, the fallback doesn't fire. A request that may already have been processed is never sent a second time.
Search returned two executable alternatives. Both are shown below exactly as the compiler recorded them.
Alternative 1: Luna alone, with Sol as fallback. 76/80 on Seldon's held-out set against Sol's 74/80, with zero critical failures for either. It costs 20% as much per request and runs about 21% faster at both p50 and p95.
Alternative 2: a Jev classifier feeding Luna. Jev makes the one enum decision, outcome ,

[truncated]

## Original Extract

Seldon Functions compiled the LLM call at the center of two hard data pipelines. Agent-trace extraction got 5× cheaper and more accurate; PII detection got ~8,000× cheaper and up to 10× faster. Here are the production receipts, including what broke.

seldon How it works Models Docs Blog v0 Log in Sign up ← Blog October 14, 2026 21 min read Seldon We fired the frontier model from two AI pipelines. Here are the receipts.
Seldon Functions compiles the LLM call in the middle of your data pipeline into a faster, cheaper program, and proves it against the model you trust. Two production experiments: 5× cheaper and more accurate on agent traces. ~8,000× cheaper and up to 10× faster on PII detection.
Routers pick a model. Seldon builds a program. We pointed it at two hard, high-volume LLM workloads, and here's what came back.
Agent trace analytics: 5× cheaper, more accurate and faster. Seldon replaced a frontier-model call with a compiled program. The program passed 95.0% of held-out checks where the frontier model passed 92.5%, cost 80% less per request and ran 21% faster. Its output also clustered better downstream, at 0.56 ARI against 0.33.
PII detection: ~8,000× cheaper, up to 10× faster, every document processed. Seldon compiled the LLM out entirely : a specialized span model plus deterministic code. It handled all 66 test documents for $0.00014 in total , including 8 that the frontier-model baseline failed on. It isn't perfect. It missed some bare first names, and the receipts below show exactly where.
Zero fallbacks, a one-line switch, and about $15 to compile. Both compiled programs served every fresh input on their own primary path. Each pays for itself in roughly 1,000 to 2,300 calls.
Switching takes one line. Seldon gives you a new model ID, and the rest of your code doesn't change:
- model="gpt-5.6-sol"
+ model="seldon/function/<FUNCTION_ID>@<REVISION_ID>"
Why it works
Most LLM calls in production aren't chat. They're functions : text goes in, JSON with a fixed schema comes out, millions of times. You pay frontier-model prices on every call because a frontier model is what worked when you wrote the prompt. The usual alternatives are shortcuts: sample less, truncate more, or swap in the cheapest model and hope. Almost nobody measures what those shortcuts cost in quality.
Seldon turns that guess into a measurement:
Your model sets the bar. You bring the prompt, the schema and a reference model you trust. Every alternative is scored against that reference, on your contract.
It searches whole programs, not just models. Smaller LLMs, specialized models (GLiNER for spans, Jev for decisions) and deterministic kernels, combined into multi-step pipelines. That's how a PII detector ends up with no LLM on its main path.
You see the trade before you make it. A comparison report puts up to three alternatives next to your reference, with quality, critical failures, cost, p50/p95 and fallback rate for each. You choose, and nothing activates on its own.
It ships with a safety net. Outputs are validated at runtime, and your original call stays wired in as a fallback. The endpoint is OpenAI-compatible, so the OpenAI SDK, LiteLLM, LangChain and the Vercel AI SDK work unchanged.
We built both experiments as standalone scripts and ran them against production api.seldon-ai.com , with real models and real wallet charges. Each compiled Function went head to head with the simple baseline (the original prompt on GPT-5.6 Sol) on the same held-out inputs. We also ran the tools you'd otherwise reach for: regex, Presidio + GLiNER, hand-written features and LLM summaries. We published the failures too.
Where these experiments come from. PostHog's public write-up on LLM trace clustering inspired the trace experiment. The anonymization needs that companies licensing enterprise data, like Mercor , describe publicly inspired the PII experiment. We designed both setups, built the data and ran everything ourselves. We don't represent either company, neither was involved or endorsed this, and nothing here measures, reproduces or challenges their systems.
The receipts: Seldon against the simple baseline
The simple baseline is what most teams ship first: send the original prompt to a frontier model (GPT-5.6 Sol) for every record. Each Seldon pipeline below sits side by side with that baseline, on the same inputs. Bold marks the better result in each row.
Trace signals (experiment inspired by PostHog's clustering write-up)
PII identification (experiment inspired by enterprise data-anonymization needs)
Held-out reports use Seldon's own synthetic examples, an independent judge and serial timing. External tests call the production gateway, so their latencies include network and gateway overhead. Seldon costs are settled receipts. Sol costs in the external tests are token-price estimates at $5 input / $30 output per million tokens. Both baselines are our own Sol calls. None of these numbers describe PostHog's or Mercor's systems.
Put plainly, Seldon beat the baseline on every trace row that differs. On PII it's far cheaper, several times faster, and finished every document. But it leaked names the baseline caught, and a local Presidio + GLiNER setup beat it on recall (see below). One result is a clear win. The other is a tradeoff you'd want to see before making it. Seldon shows you both.
The lifecycle has four explicit steps, and nothing is activated unless you activate it.
Create. Upload the call's prompt, input/output JSON schemas and reference model. The trace demo's contract is {"text": "..."} → {intent, actions[], outcome, failure_signals[], evidence[]} .
Preview (paid, capped). Seldon generates synthetic examples for the task, runs your reference on them, infers an evaluator and returns a quote: forecast cost, forecast minutes and which components are eligible. The server enforces your cap, currently up to $100.
Optimize. Bounded search over complete programs: up to 64 candidates, 16 steps and 3 learned nodes. Building blocks:
smaller LLMs (e.g. GPT-5.6 Luna)
specialized models: fastino/gliner2.5-multi-v1 for spans and records, typesafe/jev for boolean/enum/pair decisions
deterministic kernels: exact span alignment, JSON-pointer record projection, lookup, arithmetic and constraint checks
the reference itself, as an optional fallback
Small search spaces are screened exhaustively. Larger ones use SMAC random-forest Bayesian optimization over whole-program evaluations.
Compare and activate. The run publishes a comparison report with up to three executable alternatives beside the reference. Each one comes with quality, critical failures, full-path cost, p50/p95, fallback rate and coverage. You choose. An alternative can't be activated without proof that its primary path executes; a candidate that only ever succeeds through fallback is never callable.
The compiler only ever sees your prompt, your schemas, one training example and the synthetic data it generates itself. In both demos, the held-out documents and traces, and every gold label, stayed on the demo side. None of them shaped the search.
Then you call it like any other model:
from openai import OpenAI
import json
client = OpenAI(base_url="https://api.seldon-ai.com/v1", api_key=SELDON_API_KEY)
resp = client.chat.completions.create(
model=f"seldon/function/{FUNCTION_ID}@{REVISION_ID}",
messages=[{"role": "user", "content": json.dumps({"text": trace_fragment})}],
)
signals = json.loads(resp.choices[0].message.content)
It's the same OpenAI-compatible endpoint, so the OpenAI SDK, LiteLLM, LangChain and the Vercel AI SDK work unchanged. Every call writes an execution receipt with its real charge.
Where Seldon stops and your code starts
In both demos, Seldon replaces exactly one step: the LLM call your prompt used to make. Everything around it is ordinary Python in a standalone demo.py on your machine. That covers loading data, chunking, checking outputs, redaction, embeddings, clustering and scoring. The demo doesn't import the Seldon package, and it calls the Function over plain HTTP.
The diagrams below use one color code throughout:
Purple: hosted by Seldon, inside the compiled Function.
Grey: the demo script, i.e. your application code.
Green: local models running on your own CPU.
Orange: a direct call to the reference model (Sol) as a baseline, outside any Function.
Dashed: the reference fallback. It's wired in, but it never fired on the fresh inputs in these runs.
Every baseline goes through the same script. Only the one step that the Function replaces gets swapped out, so each comparison measures that step and nothing else.
PostHog has published a detailed walkthrough of how it built LLM trace clustering , with links to its code. It's one of the clearest public explanations of this problem, and it shaped how we set up our experiment. As they describe it:
each trace is rendered as line-numbered text
huge traces are uniformly downsampled to fit context
a sample of traces is taken every hour
each sampled trace gets a structured summary from GPT-4.1 nano
the summaries are embedded with text-embedding-3-large
the embeddings go through UMAP and HDBSCAN
a GPT-5.2 agent names each cluster
It's a thoughtful design, and the post is candid about tradeoffs that anyone building trace analytics runs into:
Cost forces sampling. Summarizing every trace would be expensive. So they sample a subset each hour, use a small, fast model, and plan to move to the Batch API.
Downsampling loses information. Dropping lines from huge traces keeps them inside context limits, at what they call "the cost of some accuracy/quality".
One schema for everyone. A single general-purpose summary schema can't fit every kind of trace. Letting users bring their own prompts and structured outputs is listed as a possible future improvement.
Structure beats prose. They summarize into a structured schema, with line references, specifically because it embeds better than free text.
As described, one step in that design is a function-shaped LLM call:
An illustration drawn from PostHog's published description, not from their code. Purple, thick border: the kind of step Seldon Functions is designed for. Dotted grey: an agent loop, which Functions doesn't compile.
In a design like this, the summary step is a fixed prompt plus a fixed schema, run on every sampled trace. That's the kind of call a compiler can work on, and it's the step our experiment focuses on. The labeling agent is a multi-step loop that uses tools, and Functions doesn't touch those. Embeddings, UMAP, HDBSCAN and orchestration stay in your code.
We borrowed the overall shape and changed the parts we wanted to test. This is our own setup, not a reproduction of PostHog's system, and we didn't run anything against it. The left column is only there to show where our choices differ from the design that inspired them.
Our extraction contract asks for the user's intent, the ordered actions, the observed outcome ( completed / failed / unclear ) and any failure signals. Each signal must be backed by evidence quoted exactly from the trace, and the script checks every quote in code.
Before the call, the script drops rewards, task IDs and simulator state from every trace. Tasks are hash-partitioned, so repeated attempts at one task can't leak between train and test. After the call, the whole trace is rejected if any chunk's evidence isn't an exact quote. Task IDs are only used at the very end, for scoring.
The fallback has narrow triggers. A step's output has to fail schema or grounding validation, or serving admission has to refuse the request before anything is sent. Then Seldon runs your original Sol call inside the same 55-second deadline. If a provider call's outcome is uncertain, the fallback doesn't fire. A request that may already have been processed is never sent a second time.
Search returned two executable alternatives. Both are shown below exactly as the compiler recorded them.
Alternative 1: Luna alone, with Sol as fallback. 76/80 on Seldon's held-out set against Sol's 74/80, with zero critical failures for either. It costs 20% as much per request and runs about 21% faster at both p50 and p95.
Alternative 2: a Jev classifier feeding Luna. Jev makes the one enum decision, outcome ,

[truncated]
