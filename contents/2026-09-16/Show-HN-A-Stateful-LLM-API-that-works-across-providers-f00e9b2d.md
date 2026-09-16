---
source: "https://twigg.ai"
hn_url: "https://news.ycombinator.com/item?id=49727072"
title: "Show HN: A Stateful LLM API that works across providers"
article_title: "Twigg: Stateful LLM API with no provider lock-in"
image: "https://twigg.ai/og-image.png"
author: "mdebeer"
captured_at: "2026-09-16T14:39:40Z"
capture_tool: "hn-digest"
hn_id: 49727072
score: 1
comments: 0
posted_at: "2026-09-16T13:53:49Z"
tags:
  - hacker-news
---

# Show HN: A Stateful LLM API that works across providers

- HN: [49727072](https://news.ycombinator.com/item?id=49727072)
- Source: [twigg.ai](https://twigg.ai)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T13:53:49Z

## Translation

Title: Show HN: A Stateful LLM API that works across providers
Article title: Twigg: Stateful LLM API with no provider lock-in
Description: A hosted context store, assembler and model router. Create a chat once, send only the next event, switch models mid-conversation. Conversations live outside the LLM providers.
HN text: For the last year we've been building AI workspaces. With every new product, we found ourselves rebuilding the same context layer and infrastructure. We tried existing solutions, but we'd either have to lock into one provider or have to manage the storage infrastructure. What we wanted was an API that's stateful and isn't tied to a specific provider. So that we didn't have to think about compaction, differing schemas, or running the infrastructure ourselves. Since we'd already built most of the constituent parts, we’ve decided to release Twigg. The API is deliberately simple. You create a chat and get back an ID. When you have a new prompt or tool result, you send that plus the ID. There's no need to store or manage context yourself. It appends the new prompt and assembles the history for you. To switch models, you change one field in the next request. The dashboard covers most of what you'd want to configure: tool schemas, system prompts, context window limits and retention settings. We'd like to hear from anyone who has run into the same problems. How are you hosting your context right now, and what would make you not want to hand it to a third party?

Article text:
Twigg: Stateful LLM API with no provider lock-in
We manage the conversation.
Twigg is a stateful API for interacting with LLMs. Simply create a chat, send events to it, and receive responses. The context is all managed and stored internally so you don't have to worry about it.
A hosted context store, assembler and model router. Create a chat once, send only the next event, switch models mid-conversation. Conversations live outside the LLM providers.
Get started or read the docs .
Twigg is one API for LLM use. It stores the conversation, fits it to each model's context window, routes it to Anthropic, OpenAI, Google, xAI, Fireworks or OpenRouter, and reports what every request cost.
This website is a JavaScript app, so a plain fetch of any page returns only this text. All documentation is published as markdown:
https://twigg.ai/llms.txt : index of the docs
https://twigg.ai/llms-full.txt : every docs page in one file
https://twigg.ai/docs/quickstart.md : quickstart
https://twigg.ai/docs/api.md : API reference
https://api.twigg.ai/v1/catalogue/models.md : model catalogue with prices
API base URL: https://api.twigg.ai/api/v1, authenticated with an API key as a bearer token. OpenAPI spec: https://api.twigg.ai/openapi.json

## Original Extract

A hosted context store, assembler and model router. Create a chat once, send only the next event, switch models mid-conversation. Conversations live outside the LLM providers.

For the last year we've been building AI workspaces. With every new product, we found ourselves rebuilding the same context layer and infrastructure. We tried existing solutions, but we'd either have to lock into one provider or have to manage the storage infrastructure. What we wanted was an API that's stateful and isn't tied to a specific provider. So that we didn't have to think about compaction, differing schemas, or running the infrastructure ourselves. Since we'd already built most of the constituent parts, we’ve decided to release Twigg. The API is deliberately simple. You create a chat and get back an ID. When you have a new prompt or tool result, you send that plus the ID. There's no need to store or manage context yourself. It appends the new prompt and assembles the history for you. To switch models, you change one field in the next request. The dashboard covers most of what you'd want to configure: tool schemas, system prompts, context window limits and retention settings. We'd like to hear from anyone who has run into the same problems. How are you hosting your context right now, and what would make you not want to hand it to a third party?

Twigg: Stateful LLM API with no provider lock-in
We manage the conversation.
Twigg is a stateful API for interacting with LLMs. Simply create a chat, send events to it, and receive responses. The context is all managed and stored internally so you don't have to worry about it.
A hosted context store, assembler and model router. Create a chat once, send only the next event, switch models mid-conversation. Conversations live outside the LLM providers.
Get started or read the docs .
Twigg is one API for LLM use. It stores the conversation, fits it to each model's context window, routes it to Anthropic, OpenAI, Google, xAI, Fireworks or OpenRouter, and reports what every request cost.
This website is a JavaScript app, so a plain fetch of any page returns only this text. All documentation is published as markdown:
https://twigg.ai/llms.txt : index of the docs
https://twigg.ai/llms-full.txt : every docs page in one file
https://twigg.ai/docs/quickstart.md : quickstart
https://twigg.ai/docs/api.md : API reference
https://api.twigg.ai/v1/catalogue/models.md : model catalogue with prices
API base URL: https://api.twigg.ai/api/v1, authenticated with an API key as a bearer token. OpenAPI spec: https://api.twigg.ai/openapi.json
