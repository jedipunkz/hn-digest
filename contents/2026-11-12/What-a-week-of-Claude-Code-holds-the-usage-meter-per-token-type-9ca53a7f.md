---
source: "https://atlas-ade.dev/docs/plan-usage-burn-data/"
hn_url: "https://news.ycombinator.com/item?id=50467454"
title: "What a week of Claude Code holds: the usage meter, per token type"
article_title: "Plan usage burn data · Atlas ADE docs"
image: "https://atlas-ade.dev/assets/og.png"
author: "meekoAMH"
captured_at: "2026-11-12T15:01:42Z"
capture_tool: "hn-digest"
hn_id: 50467454
score: 4
comments: 0
posted_at: "2026-11-12T14:12:31Z"
tags:
  - hacker-news
---

# What a week of Claude Code holds: the usage meter, per token type

- HN: [50467454](https://news.ycombinator.com/item?id=50467454)
- Source: [atlas-ade.dev](https://atlas-ade.dev/docs/plan-usage-burn-data/)
- Score: 4
- Comments: 0
- Posted: 2026-11-12T14:12:31Z

## Translation

Title: What a week of Claude Code holds: the usage meter, per token type
Article title: Plan usage burn data · Atlas ADE docs
Description: Claude Code plans don't give you a token allowance. They give you two meters that run from 0 to 100%. This page shows what moves those meters, and what the same usage would cost at API prices. Every r

Article text:
Skip to content
Atlas ADE
DOCS
/
How it works
Demo
Docs
Pricing
Sign in
Get Atlas ADE
Docs menu Start New to Claude Code? Start here Overview Quickstart Install Concepts How it flies The verbs Safety and privacy Guides Porting a project Day to day Release notes How it works Prove a card is done Choosing models ADE against plain sessions and a hand-built loop Field guide Setting up CLAUDE.md and AGENTS.md When to fork, compact, auto-compact or clear What is a custom agent, and how do you make one? What is a loop in agentic coding? Hooks for burn rate in Claude Code and Codex Not every issue should become a gate Context management for long-horizon agent work How to stop agent handovers becoming recency-biased Graph-derived test selection for Claude Code and Codex Claude Code usage limits by plan: how many tokens is Ultra (Max 20x)? Claude Code burn rate: why is my usage draining so fast? What a cold cache costs in Claude Code Setting the prompt cache TTL in Claude Code Run lanes on your plan's API credits Why do subagents use so much of my Claude plan? Plan usage burn data Reference What you can configure (owners) Commands (owners) ade.toml (owners) Troubleshooting (owners) Docs / Field guide / Plan usage burn data VIEW .MD COPY AS MARKDOWN Plan usage burn data
Generated from model a5v3 · Last checked against a live meter 29/10/2026
Claude Code plans don't give you a token allowance. They give you two meters that run from 0 to 100%. This page shows what moves those meters, and what the same usage would cost at API prices. Every reading here was taken in interactive Claude Code sessions on an Ultra plan (formerly Max 20x); the other plans are derived from it.
Model a5v3 · last checked against a live meter on 29/10/2026 · measured on an Ultra plan (formerly Max 20x)
Cache reads are the figure the model pins down best; output figures are the least certain part of the fit.
Only Ultra is measured. Pro and Max are derived from it using the plan-size ratios in u/flobernd's subscription cost model, so treat them as estimates. Earlier versions of this page used the limits shellac published on 25/01/2026. Note the shape: Ultra's week is about twice Max's, and before 31/10/2026 its five-hour window was about four times Max's. Plan prices checked 31/10/2026.
Weekly : the budget. 100% is one week's allowance, and it resets once a week.
Five-hour : a speed limit inside the week on Pro and Max. Ultra has had no five-hour limit since 31/10/2026 (verified accounts). Measured on Ultra before that change, the same work moved it about 3.96× as far as the weekly meter, so a full five-hour window is about 25% of a week.
Interactive and headless use #
u/flobernd found why: Claude meters a request by where it came from, so an interactive session and a headless claude -p run of the same work move the five-hour window by different amounts.
Headless runs ( claude -p , or the Agent SDK) are metered differently from interactive sessions. On Pro and Max, they use up the five-hour window faster for the same work, while the weekly meter moves much less. One public measurement over three Ultra (formerly Max 20x) accounts, measured before the change, found headless work moved the five-hour meter about 1.5 times as far ( claude-usage-tracker ). In our own data, interactive work showed no weekday peak-hours effect. Our weekly figures include some headless use, so for purely interactive work they may read slightly low.
What 1% of an Ultra week buys #
Every request spends four kinds of token, and the meter weighs them very differently. Each row is how many tokens of that kind alone would move the weekly meter by 1%.
Fable 5.2 (released 9 Nov 2026) is 20% cheaper than 5.1 at API prices. Its plan weight isn't measured yet, so this column is for 5.1.
On a plan, Claude Code keeps the main conversation's cache for an hour, but subagents get five minutes unless you set their TTL . On an API key it uses five minutes unless you change it, and five-minute writes cost less, so we show both prices.
Two things stand out. Output is the dearest token: 1% of a week buys only 660k of them. And the plan is most generous on cache: 1% spent on cache reads would cost about $35 at API prices, against $13 spent on output. A full Ultra week is worth roughly $1,200 to $3,700 of Opus 5.5 at API list prices, depending on what you spend it on.
A warm turn reads the conversation from cache. A cold turn has to write it all again, after a long pause, a compaction or a model switch. At 800k context that one turn costs about 23× a warm one. What a cold cache costs explains why, and setting the cache TTL how to have fewer of them.
Ultra. Where two figures show, the model can't yet tell which of two readings of cache writes is right.
Where these numbers come from #
We cross-analysed our own OpenTelemetry records of every request against the usage Claude Code reported, and fitted a model that turns token counts into meter points. It is checked against the live meter, and weekly totals land within about 2 points. It isn't exact: it comes from one account's long agentic coding, the Sonnet weight is an estimate, and whether small cache writes count is still open. The per-token rates the fit starts from are the ones u/flobernd measured; our fit to the live meter adjusts them. API prices are list prices, checked 03/11/2026.
Codex numbers will follow once we have a clean measured week.
These guides come out of building Atlas ADE, the meta-harness for your agentic workflows.
Try it in your browser (no sign-up) · Get the kit

## Original Extract

Claude Code plans don't give you a token allowance. They give you two meters that run from 0 to 100%. This page shows what moves those meters, and what the same usage would cost at API prices. Every r

Skip to content
Atlas ADE
DOCS
/
How it works
Demo
Docs
Pricing
Sign in
Get Atlas ADE
Docs menu Start New to Claude Code? Start here Overview Quickstart Install Concepts How it flies The verbs Safety and privacy Guides Porting a project Day to day Release notes How it works Prove a card is done Choosing models ADE against plain sessions and a hand-built loop Field guide Setting up CLAUDE.md and AGENTS.md When to fork, compact, auto-compact or clear What is a custom agent, and how do you make one? What is a loop in agentic coding? Hooks for burn rate in Claude Code and Codex Not every issue should become a gate Context management for long-horizon agent work How to stop agent handovers becoming recency-biased Graph-derived test selection for Claude Code and Codex Claude Code usage limits by plan: how many tokens is Ultra (Max 20x)? Claude Code burn rate: why is my usage draining so fast? What a cold cache costs in Claude Code Setting the prompt cache TTL in Claude Code Run lanes on your plan's API credits Why do subagents use so much of my Claude plan? Plan usage burn data Reference What you can configure (owners) Commands (owners) ade.toml (owners) Troubleshooting (owners) Docs / Field guide / Plan usage burn data VIEW .MD COPY AS MARKDOWN Plan usage burn data
Generated from model a5v3 · Last checked against a live meter 29/10/2026
Claude Code plans don't give you a token allowance. They give you two meters that run from 0 to 100%. This page shows what moves those meters, and what the same usage would cost at API prices. Every reading here was taken in interactive Claude Code sessions on an Ultra plan (formerly Max 20x); the other plans are derived from it.
Model a5v3 · last checked against a live meter on 29/10/2026 · measured on an Ultra plan (formerly Max 20x)
Cache reads are the figure the model pins down best; output figures are the least certain part of the fit.
Only Ultra is measured. Pro and Max are derived from it using the plan-size ratios in u/flobernd's subscription cost model, so treat them as estimates. Earlier versions of this page used the limits shellac published on 25/01/2026. Note the shape: Ultra's week is about twice Max's, and before 31/10/2026 its five-hour window was about four times Max's. Plan prices checked 31/10/2026.
Weekly : the budget. 100% is one week's allowance, and it resets once a week.
Five-hour : a speed limit inside the week on Pro and Max. Ultra has had no five-hour limit since 31/10/2026 (verified accounts). Measured on Ultra before that change, the same work moved it about 3.96× as far as the weekly meter, so a full five-hour window is about 25% of a week.
Interactive and headless use #
u/flobernd found why: Claude meters a request by where it came from, so an interactive session and a headless claude -p run of the same work move the five-hour window by different amounts.
Headless runs ( claude -p , or the Agent SDK) are metered differently from interactive sessions. On Pro and Max, they use up the five-hour window faster for the same work, while the weekly meter moves much less. One public measurement over three Ultra (formerly Max 20x) accounts, measured before the change, found headless work moved the five-hour meter about 1.5 times as far ( claude-usage-tracker ). In our own data, interactive work showed no weekday peak-hours effect. Our weekly figures include some headless use, so for purely interactive work they may read slightly low.
What 1% of an Ultra week buys #
Every request spends four kinds of token, and the meter weighs them very differently. Each row is how many tokens of that kind alone would move the weekly meter by 1%.
Fable 5.2 (released 9 Nov 2026) is 20% cheaper than 5.1 at API prices. Its plan weight isn't measured yet, so this column is for 5.1.
On a plan, Claude Code keeps the main conversation's cache for an hour, but subagents get five minutes unless you set their TTL . On an API key it uses five minutes unless you change it, and five-minute writes cost less, so we show both prices.
Two things stand out. Output is the dearest token: 1% of a week buys only 660k of them. And the plan is most generous on cache: 1% spent on cache reads would cost about $35 at API prices, against $13 spent on output. A full Ultra week is worth roughly $1,200 to $3,700 of Opus 5.5 at API list prices, depending on what you spend it on.
A warm turn reads the conversation from cache. A cold turn has to write it all again, after a long pause, a compaction or a model switch. At 800k context that one turn costs about 23× a warm one. What a cold cache costs explains why, and setting the cache TTL how to have fewer of them.
Ultra. Where two figures show, the model can't yet tell which of two readings of cache writes is right.
Where these numbers come from #
We cross-analysed our own OpenTelemetry records of every request against the usage Claude Code reported, and fitted a model that turns token counts into meter points. It is checked against the live meter, and weekly totals land within about 2 points. It isn't exact: it comes from one account's long agentic coding, the Sonnet weight is an estimate, and whether small cache writes count is still open. The per-token rates the fit starts from are the ones u/flobernd measured; our fit to the live meter adjusts them. API prices are list prices, checked 03/11/2026.
Codex numbers will follow once we have a clean measured week.
These guides come out of building Atlas ADE, the meta-harness for your agentic workflows.
Try it in your browser (no sign-up) · Get the kit
