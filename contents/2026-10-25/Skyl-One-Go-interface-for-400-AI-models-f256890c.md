---
source: "https://skyl-docs.vercel.app/"
hn_url: "https://news.ycombinator.com/item?id=50231961"
title: "Skyl – One Go interface for 400 AI models"
article_title: "skyl — One Go interface for every AI model"
image: "https://skyl-docs.vercel.app/opengraph-image?31ad4b31d7e91778"
author: "bagombekajob"
captured_at: "2026-10-25T14:51:57Z"
capture_tool: "hn-digest"
hn_id: 50231961
score: 1
comments: 0
posted_at: "2026-10-25T14:28:48Z"
tags:
  - hacker-news
---

# Skyl – One Go interface for 400 AI models

- HN: [50231961](https://news.ycombinator.com/item?id=50231961)
- Source: [skyl-docs.vercel.app](https://skyl-docs.vercel.app/)
- Score: 1
- Comments: 0
- Posted: 2026-10-25T14:28:48Z

## Translation

Title: Skyl – One Go interface for 400 AI models
Article title: skyl — One Go interface for every AI model
Description: skyl is a small, dependency-light Go library that lets you talk to Claude, GPT, Gemini, and 400+ other models through a single, stable interface — then switch between them by changing one string.

Article text:
skyl — One Go interface for every AI model Skip to content skyl skyl v 1.0.0 Search Ctrl K Learn Reference Community Blog Documentation
v 1.0.0 · Apache 2.0 · Go 1.22+
One Go interface for every AI model.
Talk to Claude, GPT, Gemini, Ollama and 400+ other models through a single, stable interface — then switch between them by changing one string.
client := skyl.New( anthropic.New(os.Getenv("ANTHROPIC_API_KEY")) )
resp , err := client.Complete(ctx, &skyl.Request { Model: " claude-opus-5 " , Messages: []skyl.Message { skyl.UserText( "Hello" ) } , })
fmt.Println(resp.Text()) The highlighted line is the only one that changes.
Integrating an AI model into a Go service means writing the same 400 lines every time: request mapping, SSE parsing, retry with jitter, rate-limit backoff, token accounting, error classification. skyl does that once, properly, with tests.
One Request and Response shape across every vendor. Your application code never learns a vendor’s JSON.
Model IDs are pass-through strings. A model released this morning works this morning, with no skyl release.
Every response carries Raw — the untouched provider JSON — so skyl’s abstraction is never the reason you cannot ship.
A pull iterator with proper context cancellation, verified by a goroutine-leak test rather than by convention.
Typed and classified, so you branch with errors.Is instead of matching on message text that vendors reword.
The core module has zero external dependencies. No router, no logger, no framework imposed on you.
Three native adapters for full vendor fidelity, and one generic adapter that reaches the long tail of hosts speaking OpenAI’s wire format.
Claude, native. A separate module, because it is built on the official Anthropic SDK and that brings a dozen transitive dependencies nobody else should pay for.
Claude Opus 5, Fable 5, Sonnet 5, Haiku 4.5, and the 4.x family
GPT, native. Ships inside the core module, so it costs you no dependencies.
GPT-5.6 (Sol, Terra, Luna), GPT-5.5, GPT-5.4 nano, the o-series, and gpt-oss
Gemini, native. Handles Gemini’s structural differences — contents/parts, the "model" role in place of "assistant", and systemInstruction — so your Request does not have to.
Gemini 3.6 Flash, 3.5/3.1 Flash-Lite, and 3 Pro
One adapter, configured per host, reaching the long tail of vendors that speak OpenAI's wire format — including local runtimes that need no credential at all.
~18 hosts including xAI, DeepSeek, Groq, OpenRouter, Ollama, vLLM and LM Studio
18 verified compatible endpoints — including Ollama, vLLM and LM Studio, which need no credential at all.
skyl ships a sandbox that speaks all four providers’ wire protocols locally. Every example in this documentation runs against it — no signup, no credential, no cost.
go run ./cmd/skyl-sandbox Read the no-API-key guide →
Status: v1.0.0 — the API is stable, the feature surface is not complete
Every adapter has been exercised against its live provider API, so the wire mapping is confirmed rather than merely self-consistent. Since v1.0.0 the exported API is frozen: no breaking change without a major version. What is not promised is complete provider coverage. Each adapter still silently ignores some provider features, and embeddings, prompt-caching control, batch APIs, token counting and failover are deliberately deferred. Both lists are published rather than implied.
What stands between this and production use →
skyl is licensed under Apache-2.0 . Documentation licensed under CC BY 4.0.

## Original Extract

skyl is a small, dependency-light Go library that lets you talk to Claude, GPT, Gemini, and 400+ other models through a single, stable interface — then switch between them by changing one string.

skyl — One Go interface for every AI model Skip to content skyl skyl v 1.0.0 Search Ctrl K Learn Reference Community Blog Documentation
v 1.0.0 · Apache 2.0 · Go 1.22+
One Go interface for every AI model.
Talk to Claude, GPT, Gemini, Ollama and 400+ other models through a single, stable interface — then switch between them by changing one string.
client := skyl.New( anthropic.New(os.Getenv("ANTHROPIC_API_KEY")) )
resp , err := client.Complete(ctx, &skyl.Request { Model: " claude-opus-5 " , Messages: []skyl.Message { skyl.UserText( "Hello" ) } , })
fmt.Println(resp.Text()) The highlighted line is the only one that changes.
Integrating an AI model into a Go service means writing the same 400 lines every time: request mapping, SSE parsing, retry with jitter, rate-limit backoff, token accounting, error classification. skyl does that once, properly, with tests.
One Request and Response shape across every vendor. Your application code never learns a vendor’s JSON.
Model IDs are pass-through strings. A model released this morning works this morning, with no skyl release.
Every response carries Raw — the untouched provider JSON — so skyl’s abstraction is never the reason you cannot ship.
A pull iterator with proper context cancellation, verified by a goroutine-leak test rather than by convention.
Typed and classified, so you branch with errors.Is instead of matching on message text that vendors reword.
The core module has zero external dependencies. No router, no logger, no framework imposed on you.
Three native adapters for full vendor fidelity, and one generic adapter that reaches the long tail of hosts speaking OpenAI’s wire format.
Claude, native. A separate module, because it is built on the official Anthropic SDK and that brings a dozen transitive dependencies nobody else should pay for.
Claude Opus 5, Fable 5, Sonnet 5, Haiku 4.5, and the 4.x family
GPT, native. Ships inside the core module, so it costs you no dependencies.
GPT-5.6 (Sol, Terra, Luna), GPT-5.5, GPT-5.4 nano, the o-series, and gpt-oss
Gemini, native. Handles Gemini’s structural differences — contents/parts, the "model" role in place of "assistant", and systemInstruction — so your Request does not have to.
Gemini 3.6 Flash, 3.5/3.1 Flash-Lite, and 3 Pro
One adapter, configured per host, reaching the long tail of vendors that speak OpenAI's wire format — including local runtimes that need no credential at all.
~18 hosts including xAI, DeepSeek, Groq, OpenRouter, Ollama, vLLM and LM Studio
18 verified compatible endpoints — including Ollama, vLLM and LM Studio, which need no credential at all.
skyl ships a sandbox that speaks all four providers’ wire protocols locally. Every example in this documentation runs against it — no signup, no credential, no cost.
go run ./cmd/skyl-sandbox Read the no-API-key guide →
Status: v1.0.0 — the API is stable, the feature surface is not complete
Every adapter has been exercised against its live provider API, so the wire mapping is confirmed rather than merely self-consistent. Since v1.0.0 the exported API is frozen: no breaking change without a major version. What is not promised is complete provider coverage. Each adapter still silently ignores some provider features, and embeddings, prompt-caching control, batch APIs, token counting and failover are deliberately deferred. Both lists are published rather than implied.
What stands between this and production use →
skyl is licensed under Apache-2.0 . Documentation licensed under CC BY 4.0.
