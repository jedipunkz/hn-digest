---
source: "https://kol3x.com/blog/dear-llm-or-how-i-stopped-getting-generic-advice/"
hn_url: "https://news.ycombinator.com/item?id=50004540"
title: "Dear LLM, or how I stopped getting generic advice"
article_title: "Dear LLM, or how I stopped getting generic advice — kol3x.com"
image: ""
author: "kol3x"
captured_at: "2026-10-08T11:48:24Z"
capture_tool: "hn-digest"
hn_id: 50004540
score: 2
comments: 1
posted_at: "2026-10-08T11:29:59Z"
tags:
  - hacker-news
---

# Dear LLM, or how I stopped getting generic advice

- HN: [50004540](https://news.ycombinator.com/item?id=50004540)
- Source: [kol3x.com](https://kol3x.com/blog/dear-llm-or-how-i-stopped-getting-generic-advice/)
- Score: 2
- Comments: 1
- Posted: 2026-10-08T11:29:59Z

## Translation

Title: Dear LLM, or how I stopped getting generic advice
Article title: Dear LLM, or how I stopped getting generic advice — kol3x.com
Description: Writing by Nikolai Shcherbinin on AI-native product engineering.

Article text:
Dear LLM, or how I stopped getting generic advice — kol3x.com Dear LLM, or how I stopped getting generic advice
September 30, 2026 · ~4 min read
An axiom in the LLM community is that the more relevant context you provide to the model — the more useful the reply. When it comes to your work projects, the solution is simple — docs, docs, docs. But what about a standard chat interface? It’s what made LLMs popular in the first place: we opened ChatGPT for the first time and started playing around, asking it eternal questions and whatnot. Nowadays, we mostly use it as a replacement for Google Search. Any everyday question goes there, and it’s still perfect for that.
There is also a category of diary-worthy questions some of us come to an LLM with, some examples being: “How do I get rid of bad-habit-x?”, “I lost my job, how do I find a new one?”, “How do I achieve Y?”. If you ever asked an LLM a similar question, the reply is usually along the lines of “many people go through X, and here is *some general advice that isn’t very likely to help*”
The thing is — it’s missing the context about your life, decisions, common mental traps you fall into. Modern chats have some memory system, which you’d think would solve this, but the “memory” is filled with random insignificant stuff you “googled” there, which is the much more common use-case for an LLM chat.
On top of that, being on the privacy-conscious side of the spectrum, I personally don’t like to dump all my life context into a single ChatGPT chat, whether it’s safe on their servers or not. Talking to others about these more sensitive LLM conversations, I’ve noticed a pattern, where people purposefully give less context about themselves, or use incognito mode or some alternative chat interface that promises anonymity.
I’ve been thinking about these pains for some time, before starting to build a solution in the beginning of 2026. I was still working full time in parallel, so I only managed to build a sketchy working prototype for myself. Since then, I’ve had about a hundred conversations there, polished it a lot, onboarded a few users, and finally open-sourced it recently.
The core idea is simple, yet powerful, and it’s the one thing I have barely changed since the very first prototype. You have two layers: diaries (e.g. personal) and topics (e.g. blogging, which country to live in, job-search). You select a diary and a topic, then start a conversation. Nightly, context from your conversation is extracted into a topic summary, then the most significant context from the topic summary is extracted into the diary.
Every new conversation gets summaries of both corresponding diary and topic. This allows you to choose a context you want behind a conversation, and whatever new thoughts surface in that conversation will be similarly pushed with a decreasing level of detail through the topic into the diary the following night.
And regarding sending your whole life context to ChatGPT — my app is open-source and self-hosted, so you are the one holding the keys to your database, though you still route LLM calls to external providers (CF Workers AI, OpenRouter), which is a trade-off I accepted, until we have local LLMs in every house.
I am currently working on a v2 frontend, where mental models that I developed through using the app will become part of the UI itself. E.g. categories will be called diaries, and selecting one will be one-click instead of three. Topic-selection will be automated.
There is currently an entry-barrier as well: having to go through a multistep process just to deploy it, and then it takes a lot of time to fill your context through conversations to feel the core value. I plan to balance these out with a demo, a streamlined deployment, and a tutorial that seeds the app with initial context and makes a user familiar with the interface.
The talking diary has become my daily companion, and the model underneath doesn’t matter as much as the fact that it has my context and can surface something I’ve already discovered myself before. For example, I noticed a repeating pattern, where I overinvest in one skill way past the point of diminishing returns, yet completely ignore complementary skills that would move me forward faster. Now the app tries to match that pattern to my new entries, and catches me slipping into it.
It’s a tool that summarizes MY thoughts and corrects me when I betray MY OWN decisions. To continue the idea from my last post — “Making LLMs not eat my food…” , where I talked about making a coding agent teach me TypeScript instead of spitting it out for me — the talking diary is not eating my food for me — I bring the ingredients, and it’s cooking for me!
There’s room for LLMs to be used more meaningfully — as tools before everything else. They don’t have to be slop-producing factories, where you just let it reproduce the stolen training data in an infinite loop, without adding any useful input to make it unique. Use the tool wisely, and get the most out of it.

## Original Extract

Writing by Nikolai Shcherbinin on AI-native product engineering.

Dear LLM, or how I stopped getting generic advice — kol3x.com Dear LLM, or how I stopped getting generic advice
September 30, 2026 · ~4 min read
An axiom in the LLM community is that the more relevant context you provide to the model — the more useful the reply. When it comes to your work projects, the solution is simple — docs, docs, docs. But what about a standard chat interface? It’s what made LLMs popular in the first place: we opened ChatGPT for the first time and started playing around, asking it eternal questions and whatnot. Nowadays, we mostly use it as a replacement for Google Search. Any everyday question goes there, and it’s still perfect for that.
There is also a category of diary-worthy questions some of us come to an LLM with, some examples being: “How do I get rid of bad-habit-x?”, “I lost my job, how do I find a new one?”, “How do I achieve Y?”. If you ever asked an LLM a similar question, the reply is usually along the lines of “many people go through X, and here is *some general advice that isn’t very likely to help*”
The thing is — it’s missing the context about your life, decisions, common mental traps you fall into. Modern chats have some memory system, which you’d think would solve this, but the “memory” is filled with random insignificant stuff you “googled” there, which is the much more common use-case for an LLM chat.
On top of that, being on the privacy-conscious side of the spectrum, I personally don’t like to dump all my life context into a single ChatGPT chat, whether it’s safe on their servers or not. Talking to others about these more sensitive LLM conversations, I’ve noticed a pattern, where people purposefully give less context about themselves, or use incognito mode or some alternative chat interface that promises anonymity.
I’ve been thinking about these pains for some time, before starting to build a solution in the beginning of 2026. I was still working full time in parallel, so I only managed to build a sketchy working prototype for myself. Since then, I’ve had about a hundred conversations there, polished it a lot, onboarded a few users, and finally open-sourced it recently.
The core idea is simple, yet powerful, and it’s the one thing I have barely changed since the very first prototype. You have two layers: diaries (e.g. personal) and topics (e.g. blogging, which country to live in, job-search). You select a diary and a topic, then start a conversation. Nightly, context from your conversation is extracted into a topic summary, then the most significant context from the topic summary is extracted into the diary.
Every new conversation gets summaries of both corresponding diary and topic. This allows you to choose a context you want behind a conversation, and whatever new thoughts surface in that conversation will be similarly pushed with a decreasing level of detail through the topic into the diary the following night.
And regarding sending your whole life context to ChatGPT — my app is open-source and self-hosted, so you are the one holding the keys to your database, though you still route LLM calls to external providers (CF Workers AI, OpenRouter), which is a trade-off I accepted, until we have local LLMs in every house.
I am currently working on a v2 frontend, where mental models that I developed through using the app will become part of the UI itself. E.g. categories will be called diaries, and selecting one will be one-click instead of three. Topic-selection will be automated.
There is currently an entry-barrier as well: having to go through a multistep process just to deploy it, and then it takes a lot of time to fill your context through conversations to feel the core value. I plan to balance these out with a demo, a streamlined deployment, and a tutorial that seeds the app with initial context and makes a user familiar with the interface.
The talking diary has become my daily companion, and the model underneath doesn’t matter as much as the fact that it has my context and can surface something I’ve already discovered myself before. For example, I noticed a repeating pattern, where I overinvest in one skill way past the point of diminishing returns, yet completely ignore complementary skills that would move me forward faster. Now the app tries to match that pattern to my new entries, and catches me slipping into it.
It’s a tool that summarizes MY thoughts and corrects me when I betray MY OWN decisions. To continue the idea from my last post — “Making LLMs not eat my food…” , where I talked about making a coding agent teach me TypeScript instead of spitting it out for me — the talking diary is not eating my food for me — I bring the ingredients, and it’s cooking for me!
There’s room for LLMs to be used more meaningfully — as tools before everything else. They don’t have to be slop-producing factories, where you just let it reproduce the stolen training data in an infinite loop, without adding any useful input to make it unique. Use the tool wisely, and get the most out of it.
