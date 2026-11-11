---
source: "https://blog.coship.fyi/blog/ai-tiers/"
hn_url: "https://news.ycombinator.com/item?id=50448477"
title: "The Four Tiers of AI Intelligence"
article_title: "The Four Tiers of AI Intelligence - Coship"
image: "https://blog.coship.fyi/card.png"
author: "Bolwin"
captured_at: "2026-11-11T01:02:58Z"
capture_tool: "hn-digest"
hn_id: 50448477
score: 1
comments: 0
posted_at: "2026-11-11T00:56:38Z"
tags:
  - hacker-news
---

# The Four Tiers of AI Intelligence

- HN: [50448477](https://news.ycombinator.com/item?id=50448477)
- Source: [blog.coship.fyi](https://blog.coship.fyi/blog/ai-tiers/)
- Score: 1
- Comments: 0
- Posted: 2026-11-11T00:56:38Z

## Translation

Title: The Four Tiers of AI Intelligence
Article title: The Four Tiers of AI Intelligence - Coship

Article text:
The Four Tiers of AI Intelligence
There are many ways to group AI or LLM intelligence.
In this one I weakly categorize them based on the difficulty of getting there, as influenced by their relative position to other LLMs.
Also note, due to the pace of AI development, a model is in a tier at the point of release (+/- a few days). You may fall behind after and most likely will.
Lab specific analysis are as of today, November 10th, 2026. If you’re coming here later they’re probably outdated.
The rule is: it is much easier to move within a tier than across them. Some of them are quite broad, so I’m not saying crossing the entirety one tier is harder than moving to the next tier, but rather moving up a few points on benchmarks when it crosses a tier is much harder than those same points within one. The reasons for each are explained in their respective definitions.
The best of the best, bar none.
To get here you need to have some of the best people in the world, tons of high quality data and an absolute glut of compute. You are also past the point where you can rely on any type of distillation or learning from your competitors because you are better than them . You have to teach your model how to get here by directly iterating against the real world, no tricks.
The only labs here are the big three: Anthropic, OpenAI and Google.
Google has failed to reach here since its Gemini 3 pro series, falling back to only flash models in T3 & T4. They may be coming back, as Gemini 4 Argon and Carbon are definitively a T2 models. Carbon would have reached T1, if it released when it was announced.
OpenAI was this entire tier when LLMs first took off. It still has hits but is often not #1. It’s most recent hit was GPT-6 Astra, though it was frontier mainly in new categories such as science, computer use, spatial reasoning and not really in the popular ones of agentic work, programming etc. It’s rumored they’ve had to delay the release of their best models due to safety issues, with many public incidents like the HuggingFace hack.
Anthropic as of late is the most consistent one here, with its Fable, Opus and even Sonnet models trading blows for the absolute best, and releases are frequent.
Other labs? I don’t see a good chance for any right now. Muse maybe, if they manage to improve on their Wonder Full line.
Tier 2: Behind the Frontier but better than the best existing open source
This one is easier to get to, but still very hard. The main reason, apart from just not needing to be as good, is you have access to the T1 models, both to do the R&D, and as a source of data and research. Access is limited of course, all three labs in T1 ban using their models to build a competitor, with Anthropic taking the most aggressive steps to prevent it after becoming the distillation target of choice for many labs. You don’t have access to full reasoning traces or weights or even model architectural decisions from the T1s.
Those are only available from open source models, and like in T1, you can’t rely on those because you are better than them.
Moonshot’s Kimi series has reached this tier with Kimi K2.5, K2.6 (maybe), K3 and K3.1 pro. Their original K2 and K2-0905 could also be here in the category of non-reasoning models. Since (recently) they are the only open models definitively in this tier, they also tend to literally define it. You’re here if you manage to beat Moonshot.
Google has reached here with Gemini 4 Argon and Carbon
Meta’s Muse’s Wonder was arguably here and Wonder Full will likely be, if it ever releases. Its Muse Spark line is more questionable. Benchmark wise it might be, but users say it suffered in actual use
Anthropic & OpenAI usually release cheaper second tier models (Sonnet, Sol) which fall here
Z.AI got here once with GLM 5.2 which walloped Kimi’s best (K2.7 code/K2.6) at the time. Its GLM 5.3 however did not definitively beat K3, being a little weaker on out of domain tasks and most critically, handicapped by a lack of vision. That said I give them the strongest chance of reaching T2 again.
Deepseek got here a few times with their V3 & R1 lines, but more recently, V4 pro was a flop and they haven’t tried again since. If they manage to fix V4 pro and keep improving they will likely make it here again
Qwen’s max line always seems to be on the edge, neck and neck with Kimi models.
Xiaomi’s Mimo V2.6 pro is similar to GLM 5.3: better than K3 in many benchmarks, but hard to call it definitively better in overall use. Good chance of reaching here though.
XAI’s Grok is also on the edge lately, similar to Qwen.
Tier 3: Close to Best Existing Open Source
Reaching T3 is certainly not easy but it is a lot more achievable. You now have access to the top open models to learn from and do R&D for you, with their full reasoning traces, their weights, their full architecture. Of course, you still need very smart people, lots of compute, data etc.
This is still a good place to be, as long as you have some other advantage such as cost or speed.
This is the most crowded one, so I’ll list more briefly:
Moonshot slips here when they fail to hit T2.
Z.AI is almost always here, including GLM 5.3 flash
Deepseek is here with V4.1 flash and at the date of publishing the most popular one
Xiaomi’s Mimo pro series is borderline.
StepFun’s Step 5 preview is here and the full release will likely be too.
Tencent’s Hy4 is here. It doesn’t seem well known outside of China
Google’s Gemini 3.7 & 3.8 flash have both been here
Anthropic & OpenAI release cheaper third rate models (Haiku, Terra) which land here.
Inkling reached here with Inkling 2 small, a tier leap from their previous models and a very strong trajectory. If they keep going, they will likely reach T2
Meta Muse spark has arguably been here. Its new Ember model is here.
Qwen’s Max series is borderline, as mentioned above
XAI recent Grok Models have arguably been here, but borderline
MiniMax got M2, M2.5 and M2.7 here, but have since fallen behind. They may regain this spot with M3.1 pro
I considered making another tier like “Still good at current tasks” for more separation or to clarify just because it is here does not mean it is bad (though many models here are bad). It may still be an excellent model if it does well in something like size or cost.
As an example take Qwen 3.x 27B. Not the strongest model but still quite strong and among the most loved, because it packs all its power into just 27 billion parameters, small enough to run locally, if you can afford a good gpu or a high end mac.
Good models here include Qwen’s local models, Step 3.x, Gemini Flash-lite, Mimo flash, Deepseek V4 flash, GPT Luna.
Many models here are still fine and quite interesting, they’re just not competitive with their peers. This includes Mistral, Ring/Ling, IFM’s K2 horizon, Beam, Nemotron, Minimax M3.1 flash, Trinity, Mercury
This article may miss a model or lab in in T2 & T3, so let me know in the comments of wherever this is posted.
Powered by <a class="link external" href="https://www.getzola.org" rel="">Zola</a> and <a class="link external" href="https://duckquill.daudix.one" rel="">Duckquill</a>

## Original Extract

The Four Tiers of AI Intelligence
There are many ways to group AI or LLM intelligence.
In this one I weakly categorize them based on the difficulty of getting there, as influenced by their relative position to other LLMs.
Also note, due to the pace of AI development, a model is in a tier at the point of release (+/- a few days). You may fall behind after and most likely will.
Lab specific analysis are as of today, November 10th, 2026. If you’re coming here later they’re probably outdated.
The rule is: it is much easier to move within a tier than across them. Some of them are quite broad, so I’m not saying crossing the entirety one tier is harder than moving to the next tier, but rather moving up a few points on benchmarks when it crosses a tier is much harder than those same points within one. The reasons for each are explained in their respective definitions.
The best of the best, bar none.
To get here you need to have some of the best people in the world, tons of high quality data and an absolute glut of compute. You are also past the point where you can rely on any type of distillation or learning from your competitors because you are better than them . You have to teach your model how to get here by directly iterating against the real world, no tricks.
The only labs here are the big three: Anthropic, OpenAI and Google.
Google has failed to reach here since its Gemini 3 pro series, falling back to only flash models in T3 & T4. They may be coming back, as Gemini 4 Argon and Carbon are definitively a T2 models. Carbon would have reached T1, if it released when it was announced.
OpenAI was this entire tier when LLMs first took off. It still has hits but is often not #1. It’s most recent hit was GPT-6 Astra, though it was frontier mainly in new categories such as science, computer use, spatial reasoning and not really in the popular ones of agentic work, programming etc. It’s rumored they’ve had to delay the release of their best models due to safety issues, with many public incidents like the HuggingFace hack.
Anthropic as of late is the most consistent one here, with its Fable, Opus and even Sonnet models trading blows for the absolute best, and releases are frequent.
Other labs? I don’t see a good chance for any right now. Muse maybe, if they manage to improve on their Wonder Full line.
Tier 2: Behind the Frontier but better than the best existing open source
This one is easier to get to, but still very hard. The main reason, apart from just not needing to be as good, is you have access to the T1 models, both to do the R&D, and as a source of data and research. Access is limited of course, all three labs in T1 ban using their models to build a competitor, with Anthropic taking the most aggressive steps to prevent it after becoming the distillation target of choice for many labs. You don’t have access to full reasoning traces or weights or even model architectural decisions from the T1s.
Those are only available from open source models, and like in T1, you can’t rely on those because you are better than them.
Moonshot’s Kimi series has reached this tier with Kimi K2.5, K2.6 (maybe), K3 and K3.1 pro. Their original K2 and K2-0905 could also be here in the category of non-reasoning models. Since (recently) they are the only open models definitively in this tier, they also tend to literally define it. You’re here if you manage to beat Moonshot.
Google has reached here with Gemini 4 Argon and Carbon
Meta’s Muse’s Wonder was arguably here and Wonder Full will likely be, if it ever releases. Its Muse Spark line is more questionable. Benchmark wise it might be, but users say it suffered in actual use
Anthropic & OpenAI usually release cheaper second tier models (Sonnet, Sol) which fall here
Z.AI got here once with GLM 5.2 which walloped Kimi’s best (K2.7 code/K2.6) at the time. Its GLM 5.3 however did not definitively beat K3, being a little weaker on out of domain tasks and most critically, handicapped by a lack of vision. That said I give them the strongest chance of reaching T2 again.
Deepseek got here a few times with their V3 & R1 lines, but more recently, V4 pro was a flop and they haven’t tried again since. If they manage to fix V4 pro and keep improving they will likely make it here again
Qwen’s max line always seems to be on the edge, neck and neck with Kimi models.
Xiaomi’s Mimo V2.6 pro is similar to GLM 5.3: better than K3 in many benchmarks, but hard to call it definitively better in overall use. Good chance of reaching here though.
XAI’s Grok is also on the edge lately, similar to Qwen.
Tier 3: Close to Best Existing Open Source
Reaching T3 is certainly not easy but it is a lot more achievable. You now have access to the top open models to learn from and do R&D for you, with their full reasoning traces, their weights, their full architecture. Of course, you still need very smart people, lots of compute, data etc.
This is still a good place to be, as long as you have some other advantage such as cost or speed.
This is the most crowded one, so I’ll list more briefly:
Moonshot slips here when they fail to hit T2.
Z.AI is almost always here, including GLM 5.3 flash
Deepseek is here with V4.1 flash and at the date of publishing the most popular one
Xiaomi’s Mimo pro series is borderline.
StepFun’s Step 5 preview is here and the full release will likely be too.
Tencent’s Hy4 is here. It doesn’t seem well known outside of China
Google’s Gemini 3.7 & 3.8 flash have both been here
Anthropic & OpenAI release cheaper third rate models (Haiku, Terra) which land here.
Inkling reached here with Inkling 2 small, a tier leap from their previous models and a very strong trajectory. If they keep going, they will likely reach T2
Meta Muse spark has arguably been here. Its new Ember model is here.
Qwen’s Max series is borderline, as mentioned above
XAI recent Grok Models have arguably been here, but borderline
MiniMax got M2, M2.5 and M2.7 here, but have since fallen behind. They may regain this spot with M3.1 pro
I considered making another tier like “Still good at current tasks” for more separation or to clarify just because it is here does not mean it is bad (though many models here are bad). It may still be an excellent model if it does well in something like size or cost.
As an example take Qwen 3.x 27B. Not the strongest model but still quite strong and among the most loved, because it packs all its power into just 27 billion parameters, small enough to run locally, if you can afford a good gpu or a high end mac.
Good models here include Qwen’s local models, Step 3.x, Gemini Flash-lite, Mimo flash, Deepseek V4 flash, GPT Luna.
Many models here are still fine and quite interesting, they’re just not competitive with their peers. This includes Mistral, Ring/Ling, IFM’s K2 horizon, Beam, Nemotron, Minimax M3.1 flash, Trinity, Mercury
This article may miss a model or lab in in T2 & T3, so let me know in the comments of wherever this is posted.
Powered by <a class="link external" href="https://www.getzola.org" rel="">Zola</a> and <a class="link external" href="https://duckquill.daudix.one" rel="">Duckquill</a>
