---
source: "https://nikvdp.com/notes/cost-per-task-is-the-number-to-watch/"
hn_url: "https://news.ycombinator.com/item?id=50128463"
title: "OpenAI outperforms Chinese models in cost per task"
article_title: "Cost per task is the number to watch | nikvdp"
image: ""
author: "crazyphoton"
captured_at: "2026-10-17T06:38:15Z"
capture_tool: "hn-digest"
hn_id: 50128463
score: 1
comments: 0
posted_at: "2026-10-17T06:18:18Z"
tags:
  - hacker-news
---

# OpenAI outperforms Chinese models in cost per task

- HN: [50128463](https://news.ycombinator.com/item?id=50128463)
- Source: [nikvdp.com](https://nikvdp.com/notes/cost-per-task-is-the-number-to-watch/)
- Score: 1
- Comments: 0
- Posted: 2026-10-17T06:18:18Z

## Translation

Title: OpenAI outperforms Chinese models in cost per task
Article title: Cost per task is the number to watch | nikvdp

Article text:
Cost per task is the number to watch | nikvdp
nikvdp
Blog
Use system theme
Use light theme
Use dark theme
Cost per task is the number to watch
A few months ago I wrote about “good enough” models , arguing that cheap models had gotten smart enough for most coding work, and that the expensive frontier models were best saved for the tasks that actually need them. That’s mostly played out as expected. Which leaves a new question: once you’re sure you don’t need the smartest model for a task, how do you find the cheapest one that’s still smart enough?
The obvious move is to compare prices per token, but that path is paved with footguns. What you actually pay comes down to what the industry has termed “cost per task”: how much it costs to have a model finish a given piece of work.
Take DeepSeek V4.1 Flash and OpenAI’s GPT-6.1 Sol for example. Per token, DeepSeek is seven or eight times cheaper than Sol . But on Artificial Analysis’s benchmark suite, DeepSeek burned through 250 million tokens compared to Sol on low ’s 9 million for the same set of tasks. In effect, using DeepSeek for these tasks is like hiring a contractor who charges half as much per hour, but takes four times as long and does a worse job.
To make matters worse, each lab defines a token differently. OpenAI’s Tibo Sottiaux ran the same text through OpenAI’s and Anthropic’s tokenizers, and found that GPT-5.6 Sol needed 766 tokens to Claude Opus 5’s 1,170 for the same input text. So even at the same per-token price, Opus would cost about 50% more. And then there’s the price of cached input vs uncached input, which (depending on what you’re using the models for) can massively affect the actual price you pay. Agentic coding is especially cacheable, so a model with cheap cached tokens can come out well ahead of its sticker price.
With so many variables involved, working out cost per task yourself is a bit of a pain, but luckily Artificial Analysis publishes it for every model they test, with caching already baked in. Plot score against cost and you can see how much intelligence a given budget buys. Here’s what it looks like right now:
Intelligence vs cost per task, October 2026. Live version: nikv.me/models
I have an interactive version of this chart online that keeps itself up to date. Every day an agent checks Artificial Analysis and the labs’ pricing pages and adds whatever’s new.
The purple line in the middle marks the task efficiency frontier: the best score you can buy for the price. For any models sitting below that line, you can do better for the same money (at least on a cost per task basis).
When I started the dashboard in mid-August, the purple line was a real mix. OpenAI’s Luna held the bottom, Chinese open-weights models like DeepSeek V4 Pro and Z.ai ’s GLM 5.3 held a good chunk of the middle, and Anthropic held the top, with Fable 5.1 taking it over at the start of September. Now every model on this frontier is from OpenAI or Anthropic. Opus 5.5 took the top spot from Fable, scoring higher for cheaper, and Sol 6.1 arrived at the end of September and took over most of the middle.
The orange and teal markers sitting below the line at about $2 are Z.ai ’s GLM 5.3 and Moonshot’s Kimi K3. At list prices they’re a bad deal. On a flat-rate plan with a big daily quota they’re a great one, which is something the graph can’t show (it doesn’t show token generation speed either).
So if the best deals are all American now, is that where people are spending their tokens? Surprisingly, no. OpenRouter is kind enough to publicly publish how many tokens they pipe through each model. Here’s last week’s top 20:
Tokens processed on OpenRouter in the week to Oct 15, 2026. Source: OpenRouter ( openrouter.ai/rankings ), as of Oct 15, 2026.
The top five are all from Chinese labs, and the clear number one is DeepSeek V4.1 Flash with 37 trillion tokens, 16 times as many as Sol 6.1. That’s the same DeepSeek that costs twice as much per task as Sol for a lower score. Now these are only OpenRouter’s numbers, and there’s almost certainly some selection bias going on here: far more people probably use Sol through a Codex subscription than through OpenRouter, so take this with a massive grain of salt. Unfortunately nobody outside of OpenAI knows how those tokens split across models, so we can’t crunch it here.
So what would OpenRouter look like if everyone shopped “rationally” (and we ignore the subscription effect)? Here’s the same top 20, with each model swapped for the smartest model on the graph that costs the same or less per task:
What last week’s OpenRouter tokens would look like if each of the top 20 models were swapped for the smartest model on the graph that costs the same or less per task. Free and unscored models left out. Sources: OpenRouter and Artificial Analysis, Oct 15, 2026.
In that world, the Chinese labs’ share drops from 78% to 9%. Sol 6.1 alone picks up 70%, since it beats DeepSeek, GLM 5.3 Flash and Step 5 on both score and price.
So why are so many people picking DeepSeek? Hard to say. It may come down to ineffable qualities like preferring the model’s personality, but my guess is that at least some of this is a hype cycle combined with a cursory reading of the price sheet (which really does make DeepSeek look seven or eight times cheaper), plus all the dark tokens being burned through Codex and Claude subscriptions that never show up here.
TLDR: before reaching for the cheapest-looking model, check its cost per task, not its price per token.

## Original Extract

Cost per task is the number to watch | nikvdp
nikvdp
Blog
Use system theme
Use light theme
Use dark theme
Cost per task is the number to watch
A few months ago I wrote about “good enough” models , arguing that cheap models had gotten smart enough for most coding work, and that the expensive frontier models were best saved for the tasks that actually need them. That’s mostly played out as expected. Which leaves a new question: once you’re sure you don’t need the smartest model for a task, how do you find the cheapest one that’s still smart enough?
The obvious move is to compare prices per token, but that path is paved with footguns. What you actually pay comes down to what the industry has termed “cost per task”: how much it costs to have a model finish a given piece of work.
Take DeepSeek V4.1 Flash and OpenAI’s GPT-6.1 Sol for example. Per token, DeepSeek is seven or eight times cheaper than Sol . But on Artificial Analysis’s benchmark suite, DeepSeek burned through 250 million tokens compared to Sol on low ’s 9 million for the same set of tasks. In effect, using DeepSeek for these tasks is like hiring a contractor who charges half as much per hour, but takes four times as long and does a worse job.
To make matters worse, each lab defines a token differently. OpenAI’s Tibo Sottiaux ran the same text through OpenAI’s and Anthropic’s tokenizers, and found that GPT-5.6 Sol needed 766 tokens to Claude Opus 5’s 1,170 for the same input text. So even at the same per-token price, Opus would cost about 50% more. And then there’s the price of cached input vs uncached input, which (depending on what you’re using the models for) can massively affect the actual price you pay. Agentic coding is especially cacheable, so a model with cheap cached tokens can come out well ahead of its sticker price.
With so many variables involved, working out cost per task yourself is a bit of a pain, but luckily Artificial Analysis publishes it for every model they test, with caching already baked in. Plot score against cost and you can see how much intelligence a given budget buys. Here’s what it looks like right now:
Intelligence vs cost per task, October 2026. Live version: nikv.me/models
I have an interactive version of this chart online that keeps itself up to date. Every day an agent checks Artificial Analysis and the labs’ pricing pages and adds whatever’s new.
The purple line in the middle marks the task efficiency frontier: the best score you can buy for the price. For any models sitting below that line, you can do better for the same money (at least on a cost per task basis).
When I started the dashboard in mid-August, the purple line was a real mix. OpenAI’s Luna held the bottom, Chinese open-weights models like DeepSeek V4 Pro and Z.ai ’s GLM 5.3 held a good chunk of the middle, and Anthropic held the top, with Fable 5.1 taking it over at the start of September. Now every model on this frontier is from OpenAI or Anthropic. Opus 5.5 took the top spot from Fable, scoring higher for cheaper, and Sol 6.1 arrived at the end of September and took over most of the middle.
The orange and teal markers sitting below the line at about $2 are Z.ai ’s GLM 5.3 and Moonshot’s Kimi K3. At list prices they’re a bad deal. On a flat-rate plan with a big daily quota they’re a great one, which is something the graph can’t show (it doesn’t show token generation speed either).
So if the best deals are all American now, is that where people are spending their tokens? Surprisingly, no. OpenRouter is kind enough to publicly publish how many tokens they pipe through each model. Here’s last week’s top 20:
Tokens processed on OpenRouter in the week to Oct 15, 2026. Source: OpenRouter ( openrouter.ai/rankings ), as of Oct 15, 2026.
The top five are all from Chinese labs, and the clear number one is DeepSeek V4.1 Flash with 37 trillion tokens, 16 times as many as Sol 6.1. That’s the same DeepSeek that costs twice as much per task as Sol for a lower score. Now these are only OpenRouter’s numbers, and there’s almost certainly some selection bias going on here: far more people probably use Sol through a Codex subscription than through OpenRouter, so take this with a massive grain of salt. Unfortunately nobody outside of OpenAI knows how those tokens split across models, so we can’t crunch it here.
So what would OpenRouter look like if everyone shopped “rationally” (and we ignore the subscription effect)? Here’s the same top 20, with each model swapped for the smartest model on the graph that costs the same or less per task:
What last week’s OpenRouter tokens would look like if each of the top 20 models were swapped for the smartest model on the graph that costs the same or less per task. Free and unscored models left out. Sources: OpenRouter and Artificial Analysis, Oct 15, 2026.
In that world, the Chinese labs’ share drops from 78% to 9%. Sol 6.1 alone picks up 70%, since it beats DeepSeek, GLM 5.3 Flash and Step 5 on both score and price.
So why are so many people picking DeepSeek? Hard to say. It may come down to ineffable qualities like preferring the model’s personality, but my guess is that at least some of this is a hype cycle combined with a cursory reading of the price sheet (which really does make DeepSeek look seven or eight times cheaper), plus all the dark tokens being burned through Codex and Claude subscriptions that never show up here.
TLDR: before reaching for the cheapest-looking model, check its cost per task, not its price per token.
