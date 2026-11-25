---
source: "https://blog.howardjohn.info/posts/overhyped-ai-gateway-1-semantic-caching/"
hn_url: "https://news.ycombinator.com/item?id=50637722"
title: "Overhyped AI Gateway Features Part 1: Semantic Caching"
article_title: "Overhyped AI Gateway Features Part 1: Semantic Caching | howardjohn's blog"
image: ""
author: "jmcshane"
captured_at: "2026-11-25T03:29:42Z"
capture_tool: "hn-digest"
hn_id: 50637722
score: 2
comments: 0
posted_at: "2026-11-25T02:40:23Z"
tags:
  - hacker-news
---

# Overhyped AI Gateway Features Part 1: Semantic Caching

- HN: [50637722](https://news.ycombinator.com/item?id=50637722)
- Source: [blog.howardjohn.info](https://blog.howardjohn.info/posts/overhyped-ai-gateway-1-semantic-caching/)
- Score: 2
- Comments: 0
- Posted: 2026-11-25T02:40:23Z

## Translation

Title: Overhyped AI Gateway Features Part 1: Semantic Caching
Article title: Overhyped AI Gateway Features Part 1: Semantic Caching | howardjohn's blog
Description: This is part one of a series.
Overhyped AI Gateway Features Part 1: Semantic Caching (this post)
Overhyped AI Gateway Features Part 2: Prompt Compression (coming soon!)
Overhyped AI Gateway Features Part 3: coming soon!
Over the past few years of working on an AI/LLM Gateway, I've seen a lot of feat
[truncated]

Article text:
Posts
Overhyped AI Gateway Features Part 1: Semantic Caching
Overhyped AI Gateway Features Part 1: Semantic Caching (this post)
Overhyped AI Gateway Features Part 2: Prompt Compression (coming soon!)
Overhyped AI Gateway Features Part 3: coming soon!
Over the past few years of working on an AI/LLM Gateway , I've seen a lot of features go through the hype cycle, treated as magic silver-bullets to cut costs or improve performance, pared with benchmarks that are either outright fabricated, hallucinated, or excessively cherry-picked.
In this series I wanted to cover some of the most overhyped features in the space, what they promise, what they deliver, and where the fall short.
Unlike traditional caching, which relies on an exact match of data, semantic caching attempts to cache results for similar queries.
For example, "What is the capital of France?" and "Which city is the capital of France?" are semantically similar queries, and a semantic cache could return the cached "Paris" result for both.
The benefit here is clear - a cached result is ~free, both in latency and cost.
Back when we were first looking into semantic caching in ~2023, we gave a killer demo showing theoretical million-dollar savings just caching similar "Hello" requests.
Unfortunately, the real-world falls short.
While in theory this is great, in practice there are severe limitations in most cases.
Context: in most real LLM usages, the model is provided context that is specific to the current query. For example, the logged in user, the date, the working directory, or other session-specific information. This inherently limits the effectiveness of semantic caching (hopefully! see False sharing below).
Agentic requests: when using "agents" (in this case, simply meaning multi-turn requests), semantic caching essentially cannot work. Either you use the entire conversation as the cache key (incredibly unlikely to have multiple turns and LLM responses repeated) or you look only at the last message (which loses all of the real semantic context).
Similar is not the same as identical: subtle differences that are obvious to humans may be missed by semantic similarity algorithms, leading to incorrect cache hits. For example, Convert 100 USD to EUR and Convert 100 EUR to USD have completely different meanings, but are likely to trigger a semantic cache hit.
False sharing: can occur if caching doesn't take into account user context. For example, if two different users make similar queries, a semantic cache might return one user's result to the other, potentially exposing sensitive information from their conversation.
To put this to the test, I analyzed 15,000 prompts from 2,500 Codex sessions analyzing which prompts could be a cache hit had they run through a semantic caching engine. I took a few different approaches on what was actually considered for the cache lookup. I ran across a .90 and .99 similarity threshold.
Each prompt (just looking at the most recent prompt) : 5-17% hit rate. Sounds great, right? Wrong! Every hit was a false positive. Common examples include "wdyt", "fix it", "keep going", and "review the local change" -- none of these can be correctly cached as they all depend on prior context.
Just the first prompt of a session : 8-20% hit rate. Unlike checking each prompt, which has obvious missing context from prior messages, just caching the first prompt seems intuitively more likely to succeed (albeit at minimal cost savings). However, again all cache hits would have been harmful. Many prompts were identical, but to different repos, or intentionally repeating the same prompt ("Ralph Loop"). Close matches like "fix issue 1234" and "fix issue 1243" are semantically close but completely different in reality. The case where true hits were plausible was 'hello' prompts - but these were all me testing connectivity/model availability/etc where a cache hit would obscure the result.
Caching the entire prompt (ignoring system prompts and model responses) : considering the SAME conversation the hit rate is 16-68%. This is extremely high and for a good reason: <long conversation> + last message is semantically very similar to <long conversation> . However, this means the last message is entirely ignored and the same response would be returned again which would be useless (and frustrating!). If we consider only cross-session hits, we drop to 1-3%. All of these had the same false hits as checking each prompt individually - which makes sense, as at the start of the conversation "full conversation" is the same as "each prompt"; once we get into multi-message sessions the hit rate plummets.
So overall, at least for my Codex samples, the false positive rate was extremely high - even if we had a exact-match threshold . There were essentially zero cases where semantic caching would have provided a benefit.
That being said, much simpler support-bot style use cases, where queries are more predictable and less context-dependent, might still benefit from semantic caching. However, as AI progresses and most use cases are becoming more and more context dependant, the effectiveness of semantic caching will likely continue to diminish. The presence of a single tool in the request has the potential to make cache hits invalid.

## Original Extract

This is part one of a series.
Overhyped AI Gateway Features Part 1: Semantic Caching (this post)
Overhyped AI Gateway Features Part 2: Prompt Compression (coming soon!)
Overhyped AI Gateway Features Part 3: coming soon!
Over the past few years of working on an AI/LLM Gateway, I've seen a lot of feat
[truncated]

Posts
Overhyped AI Gateway Features Part 1: Semantic Caching
Overhyped AI Gateway Features Part 1: Semantic Caching (this post)
Overhyped AI Gateway Features Part 2: Prompt Compression (coming soon!)
Overhyped AI Gateway Features Part 3: coming soon!
Over the past few years of working on an AI/LLM Gateway , I've seen a lot of features go through the hype cycle, treated as magic silver-bullets to cut costs or improve performance, pared with benchmarks that are either outright fabricated, hallucinated, or excessively cherry-picked.
In this series I wanted to cover some of the most overhyped features in the space, what they promise, what they deliver, and where the fall short.
Unlike traditional caching, which relies on an exact match of data, semantic caching attempts to cache results for similar queries.
For example, "What is the capital of France?" and "Which city is the capital of France?" are semantically similar queries, and a semantic cache could return the cached "Paris" result for both.
The benefit here is clear - a cached result is ~free, both in latency and cost.
Back when we were first looking into semantic caching in ~2023, we gave a killer demo showing theoretical million-dollar savings just caching similar "Hello" requests.
Unfortunately, the real-world falls short.
While in theory this is great, in practice there are severe limitations in most cases.
Context: in most real LLM usages, the model is provided context that is specific to the current query. For example, the logged in user, the date, the working directory, or other session-specific information. This inherently limits the effectiveness of semantic caching (hopefully! see False sharing below).
Agentic requests: when using "agents" (in this case, simply meaning multi-turn requests), semantic caching essentially cannot work. Either you use the entire conversation as the cache key (incredibly unlikely to have multiple turns and LLM responses repeated) or you look only at the last message (which loses all of the real semantic context).
Similar is not the same as identical: subtle differences that are obvious to humans may be missed by semantic similarity algorithms, leading to incorrect cache hits. For example, Convert 100 USD to EUR and Convert 100 EUR to USD have completely different meanings, but are likely to trigger a semantic cache hit.
False sharing: can occur if caching doesn't take into account user context. For example, if two different users make similar queries, a semantic cache might return one user's result to the other, potentially exposing sensitive information from their conversation.
To put this to the test, I analyzed 15,000 prompts from 2,500 Codex sessions analyzing which prompts could be a cache hit had they run through a semantic caching engine. I took a few different approaches on what was actually considered for the cache lookup. I ran across a .90 and .99 similarity threshold.
Each prompt (just looking at the most recent prompt) : 5-17% hit rate. Sounds great, right? Wrong! Every hit was a false positive. Common examples include "wdyt", "fix it", "keep going", and "review the local change" -- none of these can be correctly cached as they all depend on prior context.
Just the first prompt of a session : 8-20% hit rate. Unlike checking each prompt, which has obvious missing context from prior messages, just caching the first prompt seems intuitively more likely to succeed (albeit at minimal cost savings). However, again all cache hits would have been harmful. Many prompts were identical, but to different repos, or intentionally repeating the same prompt ("Ralph Loop"). Close matches like "fix issue 1234" and "fix issue 1243" are semantically close but completely different in reality. The case where true hits were plausible was 'hello' prompts - but these were all me testing connectivity/model availability/etc where a cache hit would obscure the result.
Caching the entire prompt (ignoring system prompts and model responses) : considering the SAME conversation the hit rate is 16-68%. This is extremely high and for a good reason: <long conversation> + last message is semantically very similar to <long conversation> . However, this means the last message is entirely ignored and the same response would be returned again which would be useless (and frustrating!). If we consider only cross-session hits, we drop to 1-3%. All of these had the same false hits as checking each prompt individually - which makes sense, as at the start of the conversation "full conversation" is the same as "each prompt"; once we get into multi-message sessions the hit rate plummets.
So overall, at least for my Codex samples, the false positive rate was extremely high - even if we had a exact-match threshold . There were essentially zero cases where semantic caching would have provided a benefit.
That being said, much simpler support-bot style use cases, where queries are more predictable and less context-dependent, might still benefit from semantic caching. However, as AI progresses and most use cases are becoming more and more context dependant, the effectiveness of semantic caching will likely continue to diminish. The presence of a single tool in the request has the potential to make cache hits invalid.
