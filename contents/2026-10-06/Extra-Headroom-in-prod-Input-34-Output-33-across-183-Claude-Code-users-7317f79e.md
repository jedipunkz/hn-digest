---
source: "https://extraheadroom.com/blog/claude-code-savings-real-usage"
hn_url: "https://news.ycombinator.com/item?id=49978145"
title: "Extra Headroom in prod: Input -34%, Output -33% across 183 Claude Code users"
article_title: "Extra Headroom in Production: Input -34%, Output -33% (2026)"
image: "https://extraheadroom.com/og/claude-code-savings-real-usage.png"
author: "gghootch"
captured_at: "2026-10-06T13:57:51Z"
capture_tool: "hn-digest"
hn_id: 49978145
score: 2
comments: 0
posted_at: "2026-10-06T13:28:25Z"
tags:
  - hacker-news
---

# Extra Headroom in prod: Input -34%, Output -33% across 183 Claude Code users

- HN: [49978145](https://news.ycombinator.com/item?id=49978145)
- Source: [extraheadroom.com](https://extraheadroom.com/blog/claude-code-savings-real-usage)
- Score: 2
- Comments: 0
- Posted: 2026-10-06T13:28:25Z

## Translation

Title: Extra Headroom in prod: Input -34%, Output -33% across 183 Claude Code users
Article title: Extra Headroom in Production: Input -34%, Output -33% (2026)
Description: Production data from 183 Claude Code users over one week: Extra Headroom removed 3.50B input tokens and avoided 132M output tokens, $27,268 at list prices.

Article text:
Headroom
How it works
Features
Benchmarks
Pricing
Docs
Articles
Claude Code
Download for macOS
Blog · Benchmarks
Extra Headroom in production: Input -34%, Output -33% across 183 Claude Code users
Garm Lucassen · October 6, 2026
Extra Headroom is a desktop app that runs a local proxy between Claude Code and the
Anthropic API. It compresses tool output before each request goes out and asks the model
for shorter replies. The benchmarks on our homepage measure single scenarios; this post
reports production data from 183 Claude Code users on Extra Headroom 0.9.21 or later,
September 29 to October 5, 2026. Over that week the app removed 3.50 billion input tokens
and avoided 132 million output tokens, $27,268 at API list prices.
These are the same Input and Output figures the app's dashboard shows each user for
their own traffic, computed over all 183 users combined.
On every turn, Claude Code appends new content to the conversation: tool results, file
reads, search output, test and build logs, JSON from MCP servers. The proxy routes each
block to a content-aware compressor and drops what the model is unlikely to need.
Dropped content stays on your machine, and the model can fetch it back with a tool call,
so compression is reversible. Context already in Anthropic's prompt cache is left as
sent, so cache hits are unaffected.
The input figure is measured: tokens removed, divided by tokens removed plus the new
input that reached the provider. Context re-sent from earlier turns is excluded from
both sides.
For output, Extra Headroom instructs the model to answer more concisely. Avoided output
is estimated against a learned baseline of reply lengths for the same kinds of requests,
and a small control group of conversations runs without the instruction. Shorter replies
also complete sooner, since generation time scales with output length.
Per user, the median input reduction for the week was 29.9%, and the median user saved
$44.30 at list prices, about $190 a month extrapolated. A quarter of users saw 40% or
more; one in ten saw over 50%.
The reduction increases with volume. Split into thirds by new input per active day, the
heaviest users saw the largest reduction:
Most of these users are on Claude Pro or Max, where the dollar figures translate into
usage rather than spend. On 58.5% of user-days, Claude Code received at least one HTTP
429 from Anthropic. Tokens removed locally never reach Anthropic, so they never count
toward those limits.
Related: what's eating the Claude Code weekly limit
and why sub-agents make long sessions expensive .
The desktop app syncs daily per-user totals: token and request counts, never prompt
content. The sample is every user on 0.9.21 or later (the first build that reports
new-input tokens) whose requests that week were mostly Claude Code and who sent at least
$5 of input at list prices: 183 users and 686 user-days; 71 of them also used another
coding agent. The results table pools tokens across users; the distribution tables use
each user's own weekly rate.
Measure it on your own traffic
The app reports the same Input and Output figures for your sessions from the first
request.
7-day free trial · no credit card required
Made by Garm Tech , powered by
Headroom by
Tejas Chopra
Leave your email and we'll send a one-page summary of the benchmarks, plus a note when big updates ship. No drip campaign.
Thanks! We'll only email you when there's something worth opening.

## Original Extract

Production data from 183 Claude Code users over one week: Extra Headroom removed 3.50B input tokens and avoided 132M output tokens, $27,268 at list prices.

Headroom
How it works
Features
Benchmarks
Pricing
Docs
Articles
Claude Code
Download for macOS
Blog · Benchmarks
Extra Headroom in production: Input -34%, Output -33% across 183 Claude Code users
Garm Lucassen · October 6, 2026
Extra Headroom is a desktop app that runs a local proxy between Claude Code and the
Anthropic API. It compresses tool output before each request goes out and asks the model
for shorter replies. The benchmarks on our homepage measure single scenarios; this post
reports production data from 183 Claude Code users on Extra Headroom 0.9.21 or later,
September 29 to October 5, 2026. Over that week the app removed 3.50 billion input tokens
and avoided 132 million output tokens, $27,268 at API list prices.
These are the same Input and Output figures the app's dashboard shows each user for
their own traffic, computed over all 183 users combined.
On every turn, Claude Code appends new content to the conversation: tool results, file
reads, search output, test and build logs, JSON from MCP servers. The proxy routes each
block to a content-aware compressor and drops what the model is unlikely to need.
Dropped content stays on your machine, and the model can fetch it back with a tool call,
so compression is reversible. Context already in Anthropic's prompt cache is left as
sent, so cache hits are unaffected.
The input figure is measured: tokens removed, divided by tokens removed plus the new
input that reached the provider. Context re-sent from earlier turns is excluded from
both sides.
For output, Extra Headroom instructs the model to answer more concisely. Avoided output
is estimated against a learned baseline of reply lengths for the same kinds of requests,
and a small control group of conversations runs without the instruction. Shorter replies
also complete sooner, since generation time scales with output length.
Per user, the median input reduction for the week was 29.9%, and the median user saved
$44.30 at list prices, about $190 a month extrapolated. A quarter of users saw 40% or
more; one in ten saw over 50%.
The reduction increases with volume. Split into thirds by new input per active day, the
heaviest users saw the largest reduction:
Most of these users are on Claude Pro or Max, where the dollar figures translate into
usage rather than spend. On 58.5% of user-days, Claude Code received at least one HTTP
429 from Anthropic. Tokens removed locally never reach Anthropic, so they never count
toward those limits.
Related: what's eating the Claude Code weekly limit
and why sub-agents make long sessions expensive .
The desktop app syncs daily per-user totals: token and request counts, never prompt
content. The sample is every user on 0.9.21 or later (the first build that reports
new-input tokens) whose requests that week were mostly Claude Code and who sent at least
$5 of input at list prices: 183 users and 686 user-days; 71 of them also used another
coding agent. The results table pools tokens across users; the distribution tables use
each user's own weekly rate.
Measure it on your own traffic
The app reports the same Input and Output figures for your sessions from the first
request.
7-day free trial · no credit card required
Made by Garm Tech , powered by
Headroom by
Tejas Chopra
Leave your email and we'll send a one-page summary of the benchmarks, plus a note when big updates ship. No drip campaign.
Thanks! We'll only email you when there's something worth opening.
