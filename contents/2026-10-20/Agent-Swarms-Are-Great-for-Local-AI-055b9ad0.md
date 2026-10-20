---
source: "https://jonready.com/blog/posts/agent-swarms-are-great-for-local-ai.html"
hn_url: "https://news.ycombinator.com/item?id=50166552"
title: "Agent Swarms Are Great for Local AI"
article_title: "Agent Swarms Are Great for Local AI — Jonathon Ready"
image: ""
author: "skadamat"
captured_at: "2026-10-20T14:44:17Z"
capture_tool: "hn-digest"
hn_id: 50166552
score: 2
comments: 1
posted_at: "2026-10-20T13:56:57Z"
tags:
  - hacker-news
---

# Agent Swarms Are Great for Local AI

- HN: [50166552](https://news.ycombinator.com/item?id=50166552)
- Source: [jonready.com](https://jonready.com/blog/posts/agent-swarms-are-great-for-local-ai.html)
- Score: 2
- Comments: 1
- Posted: 2026-10-20T13:56:57Z

## Translation

Title: Agent Swarms Are Great for Local AI
Article title: Agent Swarms Are Great for Local AI — Jonathon Ready
Description: Single-stream vibecoding makes a local LLM rig look like an expensive Chromebook. Agent swarms flip the tokenomics — and suddenly the box under the desk pays for itself.

Article text:
← Back to blog
Agent Swarms Are Great for Local AI
The tokenomics for local AI development are really bad. In order to run a coding agent with anything close to decent performance you need enough VRAM to hold the model and context in memory, which is really expensive. For Qwen3.6-35B-A3B you'd ideally have an NVIDIA 5090 or 2×3090s. In my testing I was only able to get about 2k tok/s input and 150 tok/s output in single-stream development. This is ok, but assuming you're doing 150 tok/s for an hour a day you're only achieving about 7.2 million input tokens (80% cached) and 540k output tokens.
Which, paying API rates for Qwen 3.6 on OpenRouter , you'd spend about $0.80 for the session — at $0.13/M input and $1.00/M output, with cached tokens charged at roughly a tenth of the input rate. Meaning that your $5k or so of computer might as well be a Chromebook.
Agent swarms change the tokenomics
That example was for a single-hour session at concurrency one. This is how traditional vibecoding works, but it's being replaced with the newer idea of agent swarms, where a task gets delegated and fanned out to a much larger number of agents. The big labs really like this because it can drastically increase token spend. But it also finally makes local LLM rigs worth it.
Let's assume the bulk of your one-hour session was with an average of 32 background agents doing different tasks, at an average of 20k tok/s input (14k cache, 6k new reads) and 1k tok/s out — typical speed I get on the 2×3090 rig . Paying API rates for Qwen 3.6 on OpenRouter, you'd need about $7.06 per session for this kind of agentic engineering.
While the patterns of agent loops and agent graphs get more and more popular, it actually increases the power of local AI instead of diminishing it. Because for the first time local AI has a decent way to saturate GPUs.

## Original Extract

Single-stream vibecoding makes a local LLM rig look like an expensive Chromebook. Agent swarms flip the tokenomics — and suddenly the box under the desk pays for itself.

← Back to blog
Agent Swarms Are Great for Local AI
The tokenomics for local AI development are really bad. In order to run a coding agent with anything close to decent performance you need enough VRAM to hold the model and context in memory, which is really expensive. For Qwen3.6-35B-A3B you'd ideally have an NVIDIA 5090 or 2×3090s. In my testing I was only able to get about 2k tok/s input and 150 tok/s output in single-stream development. This is ok, but assuming you're doing 150 tok/s for an hour a day you're only achieving about 7.2 million input tokens (80% cached) and 540k output tokens.
Which, paying API rates for Qwen 3.6 on OpenRouter , you'd spend about $0.80 for the session — at $0.13/M input and $1.00/M output, with cached tokens charged at roughly a tenth of the input rate. Meaning that your $5k or so of computer might as well be a Chromebook.
Agent swarms change the tokenomics
That example was for a single-hour session at concurrency one. This is how traditional vibecoding works, but it's being replaced with the newer idea of agent swarms, where a task gets delegated and fanned out to a much larger number of agents. The big labs really like this because it can drastically increase token spend. But it also finally makes local LLM rigs worth it.
Let's assume the bulk of your one-hour session was with an average of 32 background agents doing different tasks, at an average of 20k tok/s input (14k cache, 6k new reads) and 1k tok/s out — typical speed I get on the 2×3090 rig . Paying API rates for Qwen 3.6 on OpenRouter, you'd need about $7.06 per session for this kind of agentic engineering.
While the patterns of agent loops and agent graphs get more and more popular, it actually increases the power of local AI instead of diminishing it. Because for the first time local AI has a decent way to saturate GPUs.
