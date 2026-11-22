---
source: "https://morethanamachine.com/posts/slinky-escalator-benchmark/"
hn_url: "https://news.ycombinator.com/item?id=50597156"
title: "A Slinky on an Escalator Is a Weirdly Good Test of LLM Intelligence"
article_title: "Slinky Escalator Benchmark: 15 Models | More Than a Machine"
image: "https://morethanamachine.com/_astro/slinky-og.BYHNVX-K.png"
author: "reddmachine"
captured_at: "2026-11-22T07:48:11Z"
capture_tool: "hn-digest"
hn_id: 50597156
score: 1
comments: 1
posted_at: "2026-11-22T07:31:41Z"
tags:
  - hacker-news
---

# A Slinky on an Escalator Is a Weirdly Good Test of LLM Intelligence

- HN: [50597156](https://news.ycombinator.com/item?id=50597156)
- Source: [morethanamachine.com](https://morethanamachine.com/posts/slinky-escalator-benchmark/)
- Score: 1
- Comments: 1
- Posted: 2026-11-22T07:31:41Z

## Translation

Title: A Slinky on an Escalator Is a Weirdly Good Test of LLM Intelligence
Article title: Slinky Escalator Benchmark: 15 Models | More Than a Machine
Description: Fifteen models from Anthropic, OpenAI and Google try the same prompt: a rainbow slinky walking down an up escalator with real physics.

Article text:
Slinky Escalator Benchmark: 15 Models | More Than a Machine ~/ more-than-a-machine [home] [about] [posts] [contribute] [pitch] [rss] Slinky Escalator Benchmark: 15 Models
By Nishaanth Reddy · November 20, 2026
Slinky Escalator Benchmark: 15 Models
Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning
Playing Contra with decision models
Twine: Coding Agents That Challenge Each Other
StonkTrace: 100 Days of Reddit Stock Chatter: Hype, Hindsight and How LLMs Pick Stock
BubblegumTree: A Family Story Press for Kobo
Containing Coding Agents: Why I Run Them in a Jail
Ubiklaw: An Old Android Tablet Became My Local Home Assistant
Intro to Wafer Scale Compute: The Tokens Per Second Race
An Intro to Speculative Decoding
Discover Agents, Not Tools: MCP at the Wrong Layer
Anthropic Academy Courses: Reviews and Recommendations
I started this as a follow-up to the pelican-on-a-bicycle test, pushed toward inter-object physics. Every time a new model drops, I ask it for a single HTML file of a rainbow slinky walking down an up escalator forever, with real physics instead of a scripted animation. Fifteen models from Anthropic, OpenAI and Google have tried it so far.
It’s a silly test, though it’s turned into a surprisingly good gut-check for how capable these models actually are. Getting a slinky to flip down moving steps without stalling or flying off takes spring physics, step collisions and enough self-testing to spot when the result looks off.
Each model gets this exact prompt once in an empty folder, with no follow-ups or manual edits. I set reasoning effort to xhigh, or high when the agent doesn’t offer xhigh.
Make a single HTML file of a rainbow slinky walking down an up escalator, in perpetual motion. Make it realistic and not a scripted animation. Have a start, stop and reset button. copy prompt copy setup setup: agents, effort, recording Slinky on an up escalator: setup
Prompt (paste once):
Make a single HTML file of a rainbow slinky walking down an up escalator, in perpetual motion. Make it realistic and not a scripted animation. Have a start, stop and reset button.
Rules
- One fresh, empty folder per model.
- Reasoning effort xhigh, or high where the agent doesn't offer xhigh.
- Send the prompt once. No follow-ups and no edits to the output.
- Build time runs from sending the prompt to the agent's final answer.
Agents as run
- Codex CLI 0.167.0 and 0.169.1: model_reasoning_effort = "xhigh"
- Claude Code 2.1.322: /effort xhigh
- Antigravity CLI 1.5.10: --effort xhigh, or high where xhigh is not offered
Recording
- Chromium at 1280 x 720 with 2x to 3x pixel density.
- A fixed 60 Hz virtual clock drives requestAnimationFrame and performance.now.
- Press Start, let it run 0.2 s, then capture 20 s, cropped to the slinky.
[01] The Runs
swap Click any two clips below to compare them.
Gemini 4 Neon xhigh · 25 min · $4.87
Gemini 4 Hydrogen xhigh · 57 min · $10.51
Gemini 4 Carbon high · 58 min · $5.28
Gemini 4 Argon high · 22 min · $2.25
Gemini 3.8 Flash high · 21 min · $1.43
GPT-6.2 Astra xhigh · 16 min · $2.16
GPT-6.2 Sol xhigh · 11 min · $0.55
Claude Opus 5.5 xhigh · 1 h 14 min · $13.51
Claude Fable 5.2 xhigh · 1 h 4 min · $18.40
Claude Sonnet 5.5 xhigh · 1 h 37 min · $7.92
Claude Fable 5.1 xhigh · 5 h 1 min · $51.50
GPT-6 Astra xhigh · 14 min · $3.59
GPT-6.1 Sol xhigh · 11 min · $0.42
GPT-6.1 Terra xhigh · 5 min · $0.17
GPT-6 Sol xhigh · 6 min · $0.44
Even with “not a scripted animation” in the prompt, several models still bend the rules to keep the slinky going. A few adjust the escalator speed to hold it in frame. Others add a small torque assist or reset the slinky if it stalls.
These are single runs across three different agents (Codex CLI, Claude Code and Antigravity CLI), recorded on a virtual 60 Hz clock so heavy sims don’t drop frames. It’s purely an anecdotal check.
Build time runs from sending the prompt to the agent’s final answer, including any self-testing the agent did along the way. API cost is what each run’s tokens would have cost at that day’s API prices, though I ran all of these on subscription plans.
The first version of this test ran on Sep 4, 2026 with a different prompt. Both pages that came out of this first prompt were guided animations, so the November prompt asks for realistic, unscripted physics instead.
Left: GPT-6 Astra, max effort, Codex CLI. Right: Claude Fable 5.1, Claude Code. The video labels both sides "max". The Claude session most likely ran at xhigh.
Make me a single HTML file of a rainbow slinky going down an up-escalator forever. No libraries — just canvas and code. The slinky should be a chain of springs, each coil a different color of the rainbow. It starts folded in half like a horseshoe, draped over a step. When dropped it flips end-over-end down the steps, and because the escalator keeps moving up, it tumbles in place and never reaches the bottom. Include a drop button and a reset button. copy first prompt copy post link report a correction References

## Original Extract

Fifteen models from Anthropic, OpenAI and Google try the same prompt: a rainbow slinky walking down an up escalator with real physics.

Slinky Escalator Benchmark: 15 Models | More Than a Machine ~/ more-than-a-machine [home] [about] [posts] [contribute] [pitch] [rss] Slinky Escalator Benchmark: 15 Models
By Nishaanth Reddy · November 20, 2026
Slinky Escalator Benchmark: 15 Models
Handy Context: Offline Context-Aware Dictation Without Screenshots, LLMs or Fine-Tuning
Playing Contra with decision models
Twine: Coding Agents That Challenge Each Other
StonkTrace: 100 Days of Reddit Stock Chatter: Hype, Hindsight and How LLMs Pick Stock
BubblegumTree: A Family Story Press for Kobo
Containing Coding Agents: Why I Run Them in a Jail
Ubiklaw: An Old Android Tablet Became My Local Home Assistant
Intro to Wafer Scale Compute: The Tokens Per Second Race
An Intro to Speculative Decoding
Discover Agents, Not Tools: MCP at the Wrong Layer
Anthropic Academy Courses: Reviews and Recommendations
I started this as a follow-up to the pelican-on-a-bicycle test, pushed toward inter-object physics. Every time a new model drops, I ask it for a single HTML file of a rainbow slinky walking down an up escalator forever, with real physics instead of a scripted animation. Fifteen models from Anthropic, OpenAI and Google have tried it so far.
It’s a silly test, though it’s turned into a surprisingly good gut-check for how capable these models actually are. Getting a slinky to flip down moving steps without stalling or flying off takes spring physics, step collisions and enough self-testing to spot when the result looks off.
Each model gets this exact prompt once in an empty folder, with no follow-ups or manual edits. I set reasoning effort to xhigh, or high when the agent doesn’t offer xhigh.
Make a single HTML file of a rainbow slinky walking down an up escalator, in perpetual motion. Make it realistic and not a scripted animation. Have a start, stop and reset button. copy prompt copy setup setup: agents, effort, recording Slinky on an up escalator: setup
Prompt (paste once):
Make a single HTML file of a rainbow slinky walking down an up escalator, in perpetual motion. Make it realistic and not a scripted animation. Have a start, stop and reset button.
Rules
- One fresh, empty folder per model.
- Reasoning effort xhigh, or high where the agent doesn't offer xhigh.
- Send the prompt once. No follow-ups and no edits to the output.
- Build time runs from sending the prompt to the agent's final answer.
Agents as run
- Codex CLI 0.167.0 and 0.169.1: model_reasoning_effort = "xhigh"
- Claude Code 2.1.322: /effort xhigh
- Antigravity CLI 1.5.10: --effort xhigh, or high where xhigh is not offered
Recording
- Chromium at 1280 x 720 with 2x to 3x pixel density.
- A fixed 60 Hz virtual clock drives requestAnimationFrame and performance.now.
- Press Start, let it run 0.2 s, then capture 20 s, cropped to the slinky.
[01] The Runs
swap Click any two clips below to compare them.
Gemini 4 Neon xhigh · 25 min · $4.87
Gemini 4 Hydrogen xhigh · 57 min · $10.51
Gemini 4 Carbon high · 58 min · $5.28
Gemini 4 Argon high · 22 min · $2.25
Gemini 3.8 Flash high · 21 min · $1.43
GPT-6.2 Astra xhigh · 16 min · $2.16
GPT-6.2 Sol xhigh · 11 min · $0.55
Claude Opus 5.5 xhigh · 1 h 14 min · $13.51
Claude Fable 5.2 xhigh · 1 h 4 min · $18.40
Claude Sonnet 5.5 xhigh · 1 h 37 min · $7.92
Claude Fable 5.1 xhigh · 5 h 1 min · $51.50
GPT-6 Astra xhigh · 14 min · $3.59
GPT-6.1 Sol xhigh · 11 min · $0.42
GPT-6.1 Terra xhigh · 5 min · $0.17
GPT-6 Sol xhigh · 6 min · $0.44
Even with “not a scripted animation” in the prompt, several models still bend the rules to keep the slinky going. A few adjust the escalator speed to hold it in frame. Others add a small torque assist or reset the slinky if it stalls.
These are single runs across three different agents (Codex CLI, Claude Code and Antigravity CLI), recorded on a virtual 60 Hz clock so heavy sims don’t drop frames. It’s purely an anecdotal check.
Build time runs from sending the prompt to the agent’s final answer, including any self-testing the agent did along the way. API cost is what each run’s tokens would have cost at that day’s API prices, though I ran all of these on subscription plans.
The first version of this test ran on Sep 4, 2026 with a different prompt. Both pages that came out of this first prompt were guided animations, so the November prompt asks for realistic, unscripted physics instead.
Left: GPT-6 Astra, max effort, Codex CLI. Right: Claude Fable 5.1, Claude Code. The video labels both sides "max". The Claude session most likely ran at xhigh.
Make me a single HTML file of a rainbow slinky going down an up-escalator forever. No libraries — just canvas and code. The slinky should be a chain of springs, each coil a different color of the rainbow. It starts folded in half like a horseshoe, draped over a step. When dropped it flips end-over-end down the steps, and because the escalator keeps moving up, it tumbles in place and never reaches the bottom. Include a drop button and a reset button. copy first prompt copy post link report a correction References
