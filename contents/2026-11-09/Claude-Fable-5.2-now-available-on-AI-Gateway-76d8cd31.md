---
source: "https://vercel.com/changelog/claude-fable-5-2-now-available-on-ai-gateway"
hn_url: "https://news.ycombinator.com/item?id=50430691"
title: "Claude Fable 5.2 now available on AI Gateway"
article_title: "Claude Fable 5.2 now available on AI Gateway - Vercel"
image: "https://assets.vercel.com/image/upload/contentful/image/e5382hct74si/1nxbUaUVfuEnDEKT9lRXrY/a709d6455cfbd73f7ef021e13322f8b4/Vercel_x_Anthropic.png"
author: "alexdykyi"
captured_at: "2026-11-09T20:28:34Z"
capture_tool: "hn-digest"
hn_id: 50430691
score: 1
comments: 1
posted_at: "2026-11-09T19:38:20Z"
tags:
  - hacker-news
---

# Claude Fable 5.2 now available on AI Gateway

- HN: [50430691](https://news.ycombinator.com/item?id=50430691)
- Source: [vercel.com](https://vercel.com/changelog/claude-fable-5-2-now-available-on-ai-gateway)
- Score: 1
- Comments: 1
- Posted: 2026-11-09T19:38:20Z

## Translation

Title: Claude Fable 5.2 now available on AI Gateway
Article title: Claude Fable 5.2 now available on AI Gateway - Vercel
Description: Call Anthropic's Claude Fable 5.2 through Vercel AI Gateway with one API key: no provider account, fallbacks, spend tracking, and traces on every request.

Article text:
Claude Fable 5.2 now available on AI Gateway - Vercel Skip to content Copy Wordmark Copy Logo Download Brand Assets Brand Guidelines Products Agent Stack AI SDK
Claude Fable 5.2 now available on AI Gateway
Claude Fable 5.2 from Anthropic is now available on AI Gateway .
Fable 5.2 is built for open-ended technical work and improves on earlier Fable models in coding, reasoning, and vision. It can run long, ambiguous tasks unattended, and works as an orchestrator that hands simpler tasks to smaller models, or as an advisor that other Claude models consult on harder problems. It also communicates more clearly and generates output faster than earlier Fable models. Effort runs from low to max , with high as the API default.
Fable 5.2 has cybersecurity and biology safety classifiers enabled. Access to Claude Fable 5.2 requires data retention. Anthropic does not support Zero Data Retention for the model. Prompts and completions are retained for 30 days and are not used to train Claude. If your team has ZDR turned on, AI Gateway returns an error instead of sending the request to Fable 5.2.
To keep requests served when a classifier is triggered, use model fallbacks . Add a models array to providerOptions.gateway , and AI Gateway tries each model in order when the cybersecurity classifier declines a request.
Call it through the AI SDK , the Chat Completions , Responses , and Anthropic Messages APIs, or a coding agent connected to AI Gateway. Use anthropic/claude-fable-5.2 as the model name:
API Claude Code fx Chat Completions API API Claude Code Claude Code fx fx Chat Completions Chat Completions 1 import { streamText } from 'ai' ; 2
3 const result = streamText ( { 4 model : 'anthropic/claude-fable-5.2' , 5 reasoning : 'high' , 6 prompt : 'Review this login handler for security bugs.' , 7 providerOptions : { 8 gateway : { 9 models : [ 'anthropic/claude-opus-5.5' , 'anthropic/claude-sonnet-5.5' ] , 10 } , 11 } , 12 } ) ; Copy for agent View model details ↗ This request falls back to Opus 5.5, then Sonnet 5.5, if the cybersecurity classifier declines it. The models option works the same way on every AI Gateway API format.
To use it in Claude Code, Codex, Cursor, and more, install the latest Vercel CLI and run setup:
npx vercel ai-gateway setup Try Claude Fable 5.2 in the model playground .
AI Gateway provides a unified API for calling models, tracking usage and cost, and configuring retries, failover, and performance optimizations for higher-than-provider uptime. It includes built-in custom reporting , Zero Data Retention support , budgets for API keys , routing rules , and more.
AI Gateway reflects provider pricing with no markup and does not charge a platform fee on inference, including on Bring Your Own Key (BYOK) requests.
You can view all language models available on AI Gateway.

## Original Extract

Call Anthropic's Claude Fable 5.2 through Vercel AI Gateway with one API key: no provider account, fallbacks, spend tracking, and traces on every request.

Claude Fable 5.2 now available on AI Gateway - Vercel Skip to content Copy Wordmark Copy Logo Download Brand Assets Brand Guidelines Products Agent Stack AI SDK
Claude Fable 5.2 now available on AI Gateway
Claude Fable 5.2 from Anthropic is now available on AI Gateway .
Fable 5.2 is built for open-ended technical work and improves on earlier Fable models in coding, reasoning, and vision. It can run long, ambiguous tasks unattended, and works as an orchestrator that hands simpler tasks to smaller models, or as an advisor that other Claude models consult on harder problems. It also communicates more clearly and generates output faster than earlier Fable models. Effort runs from low to max , with high as the API default.
Fable 5.2 has cybersecurity and biology safety classifiers enabled. Access to Claude Fable 5.2 requires data retention. Anthropic does not support Zero Data Retention for the model. Prompts and completions are retained for 30 days and are not used to train Claude. If your team has ZDR turned on, AI Gateway returns an error instead of sending the request to Fable 5.2.
To keep requests served when a classifier is triggered, use model fallbacks . Add a models array to providerOptions.gateway , and AI Gateway tries each model in order when the cybersecurity classifier declines a request.
Call it through the AI SDK , the Chat Completions , Responses , and Anthropic Messages APIs, or a coding agent connected to AI Gateway. Use anthropic/claude-fable-5.2 as the model name:
API Claude Code fx Chat Completions API API Claude Code Claude Code fx fx Chat Completions Chat Completions 1 import { streamText } from 'ai' ; 2
3 const result = streamText ( { 4 model : 'anthropic/claude-fable-5.2' , 5 reasoning : 'high' , 6 prompt : 'Review this login handler for security bugs.' , 7 providerOptions : { 8 gateway : { 9 models : [ 'anthropic/claude-opus-5.5' , 'anthropic/claude-sonnet-5.5' ] , 10 } , 11 } , 12 } ) ; Copy for agent View model details ↗ This request falls back to Opus 5.5, then Sonnet 5.5, if the cybersecurity classifier declines it. The models option works the same way on every AI Gateway API format.
To use it in Claude Code, Codex, Cursor, and more, install the latest Vercel CLI and run setup:
npx vercel ai-gateway setup Try Claude Fable 5.2 in the model playground .
AI Gateway provides a unified API for calling models, tracking usage and cost, and configuring retries, failover, and performance optimizations for higher-than-provider uptime. It includes built-in custom reporting , Zero Data Retention support , budgets for API keys , routing rules , and more.
AI Gateway reflects provider pricing with no markup and does not charge a platform fee on inference, including on Bring Your Own Key (BYOK) requests.
You can view all language models available on AI Gateway.
