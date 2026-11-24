---
source: "https://protocolis.co/paper/"
hn_url: "https://news.ycombinator.com/item?id=50632998"
title: "Compiling Web Applications into Verified Operations for AI Agents"
article_title: "Solid-State Web Execution: Compiling Web Applications into Verified Operations for AI Agents | Protocolis"
image: "https://protocolis.co/media/og-card-v2.jpg"
author: "vahagnb"
captured_at: "2026-11-24T19:58:50Z"
capture_tool: "hn-digest"
hn_id: 50632998
score: 11
comments: 0
posted_at: "2026-11-24T19:12:02Z"
tags:
  - hacker-news
---

# Compiling Web Applications into Verified Operations for AI Agents

- HN: [50632998](https://news.ycombinator.com/item?id=50632998)
- Source: [protocolis.co](https://protocolis.co/paper/)
- Score: 11
- Comments: 0
- Posted: 2026-11-24T19:12:02Z

## Translation

Title: Compiling Web Applications into Verified Operations for AI Agents
Article title: Solid-State Web Execution: Compiling Web Applications into Verified Operations for AI Agents | Protocolis
Description: A deterministic compiler turns a web page, on first contact and with no model call, into typed operations that an AI agent invokes, each returning a verified result. Compared with widely used browser tools and agents, across three task lengths and three model sizes.

Article text:
rotocolis
demonstration
benchmarks
engine
comparisons
arXiv paper
platform
Coming soon to arXiv Being posted to arXiv (cs.AI). Until it is announced there, this is the author’s copy.
Solid-State Web Execution: Compiling Web Applications into Verified Operations for AI Agents
A web agent pays a language model to read the page at every step: first to work out what it can do there, then to find out whether its last action worked. Solid-state web execution removes both. On arrival at a page it has never seen, the Protocolis compiler turns the page's controls into typed operations; the model picks one; the compiler runs it, reads the page again and reports what actually changed. Compilation uses no model inference, no site-specific code and no prior visit. To our knowledge it is the first such compiler in operation. We tested it against the browser tools and agents in common use, as shipped, across three model sizes and three task lengths, with its memory and learning switched off. On long tasks that every system attempted, its default mode completed 76 percent (80 of 105); the best alternative in each comparison completed 46 percent. It was the most stable system tested: it never fell below 55 percent in any condition, while every alternative fell to 28 percent or lower. The cheapest alternative on the same model cost 1.9 to 26 times as much per completed task. On a small model it cost less than every alternative on a frontier model, in 30 of 30 pairings, and matched or exceeded its completion rate in 24. When it reported a long task finished, the judge disagreed less than half as often as with the alternatives (22 against 56 percent). Every result is a first visit, the least favourable setting for the design. We publish the tasks, the judging rules and the outcome of all 3,443 counted runs.
On long tasks that every system attempted, its default mode completed 76 percent (80 of 105); the best alternative in each comparison completed 46 percent.
It was the most stable system tested: it never fell below 55 percent in any condition, while every alternative fell to 28 percent or lower.
The cheapest alternative on the same model cost 1.9 to 26 times as much per completed task.
On a small model it cost less than every alternative on a frontier model, in 30 of 30 pairings, and matched or exceeded its completion rate in 24.
Coming soon to arXiv Author’s copy until it is announced there.
Replication kit · Offered for acquisition · Get in touch
Recompute every figure from the replication kit · download (.zip)
Patent-pending, in production with paying customers, offered for acquisition. Evaluation access on request: get in touch .

## Original Extract

A deterministic compiler turns a web page, on first contact and with no model call, into typed operations that an AI agent invokes, each returning a verified result. Compared with widely used browser tools and agents, across three task lengths and three model sizes.

rotocolis
demonstration
benchmarks
engine
comparisons
arXiv paper
platform
Coming soon to arXiv Being posted to arXiv (cs.AI). Until it is announced there, this is the author’s copy.
Solid-State Web Execution: Compiling Web Applications into Verified Operations for AI Agents
A web agent pays a language model to read the page at every step: first to work out what it can do there, then to find out whether its last action worked. Solid-state web execution removes both. On arrival at a page it has never seen, the Protocolis compiler turns the page's controls into typed operations; the model picks one; the compiler runs it, reads the page again and reports what actually changed. Compilation uses no model inference, no site-specific code and no prior visit. To our knowledge it is the first such compiler in operation. We tested it against the browser tools and agents in common use, as shipped, across three model sizes and three task lengths, with its memory and learning switched off. On long tasks that every system attempted, its default mode completed 76 percent (80 of 105); the best alternative in each comparison completed 46 percent. It was the most stable system tested: it never fell below 55 percent in any condition, while every alternative fell to 28 percent or lower. The cheapest alternative on the same model cost 1.9 to 26 times as much per completed task. On a small model it cost less than every alternative on a frontier model, in 30 of 30 pairings, and matched or exceeded its completion rate in 24. When it reported a long task finished, the judge disagreed less than half as often as with the alternatives (22 against 56 percent). Every result is a first visit, the least favourable setting for the design. We publish the tasks, the judging rules and the outcome of all 3,443 counted runs.
On long tasks that every system attempted, its default mode completed 76 percent (80 of 105); the best alternative in each comparison completed 46 percent.
It was the most stable system tested: it never fell below 55 percent in any condition, while every alternative fell to 28 percent or lower.
The cheapest alternative on the same model cost 1.9 to 26 times as much per completed task.
On a small model it cost less than every alternative on a frontier model, in 30 of 30 pairings, and matched or exceeded its completion rate in 24.
Coming soon to arXiv Author’s copy until it is announced there.
Replication kit · Offered for acquisition · Get in touch
Recompute every figure from the replication kit · download (.zip)
Patent-pending, in production with paying customers, offered for acquisition. Evaluation access on request: get in touch .
