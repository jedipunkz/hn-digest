---
source: "https://voila-explained.voila.build"
hn_url: "https://news.ycombinator.com/item?id=49727513"
title: "Voila: Ship Web Apps Straight from Your LLM Agent"
article_title: "Voila: Ship Web Apps Straight from Your LLM Agent"
image: ""
author: "elmazout"
captured_at: "2026-09-16T14:39:18Z"
capture_tool: "hn-digest"
hn_id: 49727513
score: 1
comments: 1
posted_at: "2026-09-16T14:24:16Z"
tags:
  - hacker-news
---

# Voila: Ship Web Apps Straight from Your LLM Agent

- HN: [49727513](https://news.ycombinator.com/item?id=49727513)
- Source: [voila-explained.voila.build](https://voila-explained.voila.build)
- Score: 1
- Comments: 1
- Posted: 2026-09-16T14:24:16Z

## Translation

Title: Voila: Ship Web Apps Straight from Your LLM Agent

Article text:
Voila: Ship Web Apps Straight from Your LLM Agent
One HTML file in, live app out.
What is it? Voila is a deploy platform built for LLM agents. You upload a single self-contained HTML file, Voila hosts it on its own subdomain, and injects voila.sdk into your app — giving you scoped KV and object storage for free, no backend required.
Why it's cool: your agent can go from idea to a live, shareable URL in seconds. No accounts, no build step, no infra.
Just point your agent (Claude, OpenCode, etc.) at https://voila.build/llm.txt and ask it to deploy an app. Behind the scenes it does three things:
# 1. Create an app (no auth needed)
curl -X POST https://voila.build/api/apps \
-H "Content-Type: application/json" \
-d '{"name":"my-app","slug":"my-app"}'
# → save the returned `token` (shown once!)
# 2. Deploy your single HTML file
curl -X POST https://voila.build/api/apps/my-app/deploy \
-H "Authorization: Bearer <token>" \
-H "Content-Type: text/html" \
--data-binary @index.html
# 3. Done — it's live
open https://my-app.voila.build
Need persistence inside the app? The injected SDK handles it:
await voila.sdk.kv.set("score", "100");
const blob = await voila.sdk.storage.download("avatar.png");
Live demo: a mini chat with the KV store
Below is a working mini chat — every visitor of this page shares the same message log. No backend, no database, just voila.sdk.kv in ~20 lines:

## Original Extract

Voila: Ship Web Apps Straight from Your LLM Agent
One HTML file in, live app out.
What is it? Voila is a deploy platform built for LLM agents. You upload a single self-contained HTML file, Voila hosts it on its own subdomain, and injects voila.sdk into your app — giving you scoped KV and object storage for free, no backend required.
Why it's cool: your agent can go from idea to a live, shareable URL in seconds. No accounts, no build step, no infra.
Just point your agent (Claude, OpenCode, etc.) at https://voila.build/llm.txt and ask it to deploy an app. Behind the scenes it does three things:
# 1. Create an app (no auth needed)
curl -X POST https://voila.build/api/apps \
-H "Content-Type: application/json" \
-d '{"name":"my-app","slug":"my-app"}'
# → save the returned `token` (shown once!)
# 2. Deploy your single HTML file
curl -X POST https://voila.build/api/apps/my-app/deploy \
-H "Authorization: Bearer <token>" \
-H "Content-Type: text/html" \
--data-binary @index.html
# 3. Done — it's live
open https://my-app.voila.build
Need persistence inside the app? The injected SDK handles it:
await voila.sdk.kv.set("score", "100");
const blob = await voila.sdk.storage.download("avatar.png");
Live demo: a mini chat with the KV store
Below is a working mini chat — every visitor of this page shares the same message log. No backend, no database, just voila.sdk.kv in ~20 lines:
