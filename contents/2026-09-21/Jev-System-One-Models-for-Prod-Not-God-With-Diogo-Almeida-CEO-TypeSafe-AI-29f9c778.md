---
source: "https://www.latent.space/p/jev"
hn_url: "https://news.ycombinator.com/item?id=49794590"
title: "Jev: System One Models for Prod, Not God – With Diogo Almeida, CEO, TypeSafe AI"
article_title: "Jev: System One models for Prod, not God — with Diogo Almeida, CEO, TypeSafe AI"
image: "https://substackcdn.com/image/fetch/$s_!4FCs!,w_1200,h_600,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-video.s3.amazonaws.com%2Fvideo_upload%2Fpost%2F216783460%2F7e63c899-1c78-4714-a7ae-bbf4a57a9454%2Ftranscoded-1790031168.png"
author: "swyx"
captured_at: "2026-09-21T23:52:00Z"
capture_tool: "hn-digest"
hn_id: 49794590
score: 4
comments: 1
posted_at: "2026-09-21T23:01:51Z"
tags:
  - hacker-news
---

# Jev: System One Models for Prod, Not God – With Diogo Almeida, CEO, TypeSafe AI

- HN: [49794590](https://news.ycombinator.com/item?id=49794590)
- Source: [www.latent.space](https://www.latent.space/p/jev)
- Score: 4
- Comments: 1
- Posted: 2026-09-21T23:01:51Z

## Translation

Title: Jev: System One Models for Prod, Not God – With Diogo Almeida, CEO, TypeSafe AI
Article title: Jev: System One models for Prod, not God — with Diogo Almeida, CEO, TypeSafe AI
Description: The definitive Jev podcast with its lead creator.

Article text:
Jev: System One models for Prod, not God — with Diogo Almeida, CEO, TypeSafe AI
Subscribe Sign in Latent Space: The AI Engineer Podcast Jev: System One models for Prod, not God — with Diogo Almeida, CEO, TypeSafe AI 15 1 1× 0:00 Current time: 0:00 / Total time: -2:20:52 -2:20:52 Audio playback is not supported on your browser. Please upgrade. Jev: System One models for Prod, not God — with Diogo Almeida, CEO, TypeSafe AI
15 1 Share Transcript Tickets for AIE NYC now open, and apply for the invite-only AIE CODE . Join us !
We have an unusual relationship with today’s guest: for years since coauthoring the InstructGPT paper , Diogo Almeida had been saying that API-available frontier models have been going down the wrong path, everything from the alignment to refusals to reliability perspectives, that we have dropped every mode other than autoregressive chat-tuned LLMs because of the overwhelming success of ChatGPT.
In a launch video now viewed ~40M times (by comparison, GPT4o was 22M , Fable 5 was 15M , Navier Stokes was 74M , and 6 Astra was 137M ), Diogo introduced Jev and it immediately took over the AI timeline — we’ll skip full Jev explainers because your favorite AI influencer/educator has probably already done one. We also collected:
the official patterns and cookbooks you should see first, from Allie
speed based - games and computer use
the voice + computer use example we discuss at 1h34 mins
guided responses in text messages
Jev for coding agents has an official guide
reasonable pushback from Theo - Diogo has published a note on the Tyranny of the KV Cache that you should read as a followup after the pod for Jev + coding agents, because of his belief that Cache Rules Everything
Programming Languages built atop Jev (Diogo’s fave)
Jev for analytics replay and user journey review
a core goal of Jev is to “disappear into the background” - eg as unremarkable as regex
blending transformers and classifiers
Jev vs GLiNER (note difference/pushback , agreed , agreed , agreed )
Instead we’ll focus on what we can uniquely offer — a broader philosophical and mission-based understanding of how and why Jev was created , and what you should expect next in terms of future models from TypeSafe ( ReasoningJev ?) and what usecases and ideas you should work on vs the 55th low effort clone of Jev’s API or doing a generic JevBench benchmark - something Diogo has rejected publicly .
Why RLCD: Three kinds of RLHF, and why they are ALL the wrong north star
Diogo knows a good deal about RLHF, given that he was on the team that pioneered post-training at OpenAI — and traces the three branches to Christiano et al 2017 (the robot backflip demo), Stiennon et al 2020 (learning to summarize) and his baby, Ouyang et al 2022 (InstructGPT). From there on, every innovation from Function Calling to Structured Outputs to Reasoning felt like a hack on top of the string based, sequence to sequence prediction paradigm. As he mentions on the pod, from 2023-2024 he struggled unsuccessfully, due to both personal and organization underestimation, to train a model that accurately addressed what he saw as the core problem with making LLMs the heart of software: reliability .
Jev’s core innovation is " Reinforcement Learning for Calibrated Decisions ”, a novel, unpublished technique that optimizes for “answers with epistemically honest probabilities on System One tasks” rather than human rated feedback (RLHF) — which causes hallucinations, sycophancy, and permanent reliance on humans — or programmatically verifiable outputs with rubrics (RLVR) — which solves Navier Stokes but exacerbates jagged intelligence and doesn’t integrate well with other software.
We’ve talked about the calibration problem before on the pod, but probably the single best place to understand why RLCD became necessary is Diogo’s AIE talk , which discusses why a generation of training helpful AI assistants for humans has impaired them for training models for composable, programmable AI for automation .
At the end he also teases his contrarian opinion on scaling laws - which teases how to build a modern neolab without the billions of dollars the major labs have…
The Bitterest Lesson: Tasks and Data beats Compute
We spend a good amount of time discussing Diogo’s essay on the Bitterest Lesson :
His point is that “You get what you optimize for and the bitterest lesson in ML is that the most important part of it isn’t ML at all.” - and picking the right north star, eg upvoting for user preference vs being integrated into tool calls - makes everything else fall in line.
We’re excited to catch up with a freshly dyed Diogo to discuss:
Why AI can solve extraordinarily hard problems but still fail to automate basic work
What System One Models are and why Jev is built for software rather than chat
RLHF, mode collapse , calibration, and the hidden costs of optimizing for human preferences
Why refusals become a problem when AI is buried inside software dependencies
Why TypeSafe rejects public benchmarks and optimizes for intelligence per dollar
The “bitterest lesson”: why the right task and the right data can matter more than compute
Why TypeSafe thinks of itself as a data lab rather than a model lab
RLCD vs. RLHF and RLVR as fundamentally different North Stars for AI
Why reliability and robustness matter more than simple determinism
Jev’s programming primitives and how intelligence maps into software control flow
Why developers should decompose AI workflows into small, measurable decisions
How structured state replaces giant prompts and system messages
Why Diogo thinks AI should eventually disappear into the background of software
The “inverse SaaS-pocalypse” and how AI could supercharge existing software
System One vs. System Two intelligence and the limits of reasoning models
Dark data, computer use, real-time intelligence, and Jev’s biggest early use cases
Why Jev could reshape coding agents built around a single-model architecture
Why Diogo says he wouldn’t pre-train with $1 billion
The OpenAI journey that led to TypeSafe and why he thinks many neo-labs are approaching AI incorrectly
Coding agents beyond the KV cache , shared state, sub-agents, and the multi-agent future
LinkedIn: https://www.linkedin.com/in/diogomda
X: https://x.com/CompleteSkeptic
TypeSafe AI: https://typesafe.ai/
00:00:00 Jev Launch Week and the AI Economic Revolution
00:02:50 What Is Jev? System One Models and Programmable AI
00:05:54 RLHF, Mode Collapse, Calibration, and Yann LeCun
00:10:29 Programmatic AI, Refusals, and Safety Alignment
00:17:21 Why TypeSafe Rejects Public Benchmarks
00:20:43 The Bitterest Lesson: Data, Compute, and the Right Task
00:24:59 RLCD vs. RLHF and RLVR
00:28:42 Why Powerful AI Still Hasn’t Automated the Economy
00:39:55 Reliability, Robustness, and Determinism
00:48:11 Model Versioning, LTS, Speed, and Intelligence per Dollar
00:54:04 Inside Jev’s API and Programming Primitives
00:58:28 How to Build with Jev: Structure, Decomposition, and Small Decisions
01:18:28 The Inverse SaaS-pocalypse and AI Disappearing into Software
01:33:21 Computer Use, Dark Data, and Jev’s Biggest Use Cases
01:38:48 How Jev Could Reshape Coding Agents
01:41:00 AI Safety, Frontier Pacing, and the Limits of RLVR
01:48:03 Why Diogo Wouldn’t Pre-Train with $1 Billion
01:55:19 The OpenAI Story Behind TypeSafe
02:01:41 Why Diogo Thinks Most Neo-Labs Are Getting AI Wrong
02:08:00 Coding Agents Beyond the KV Cache and the Multi-Agent Future
Introduction: Jev Launch Week and Developer Momentum
Swyx [00:00:00]: Okay, we’re in the studio. A special occasion because this week, Diogo, my good buddy, launched Jev, and it’s been taking over the complete timeline. How do you feel? What’s it like to be you right now?
Diogo Almeida [00:00:16]: Emotionally?
Diogo Almeida [00:00:17]: Never been worse. Like, I’m a ragged corpse of a person right now because there’s so much going on, and I’m like a technical CEO, so I have, like, a lot of fires to fight.
Diogo Almeida [00:00:29]: But mentally, I feel—I say this all the time, and I’ve been saying this kind of for years in my over-under events. Like, I feel like the entire AI field is like one of those, like, carnival house of mirrors, and everyone is just insane and saying the weirdest stuff that doesn’t make sense. And it feels like for just this week, like, I’m on a better in sync with reality and like, oh, people see it now. AI can be so much more than what was once thought.
Diogo Almeida [00:01:06]: And like, yes, we are going to make. Like, an AI-based economic revolution is back on the table, and this is fucking awesome.
Diogo Almeida [00:01:17]: I’m so jazzed the developers get it. It’s, it’s, Yeah, and I want to show my eternal gratitude to the developers and
Diogo Almeida [00:01:26]: I’m so jazzed about the community and everything. It’s so great.
Swyx [00:01:28]: Yeah, you were saying yesterday that you decided to prioritize the town hall and not a bunch of, like, VIP, investor-type people because you wanted to make sure that they are the people that you get your most, attention, right? The engineers, the developers.
Diogo Almeida [00:01:43]: Yeah, it felt a little like, oh man, I’m talking to, like, really important people right now.
Diogo Almeida [00:01:47]: I probably shouldn’t reveal who.
Diogo Almeida [00:01:48]: But it feels a little bit dirty for me to, I’m, like, perhaps overly genuine in things. Like, it feels, like, dirty if, like, in my gigantic calendar event of people to talk to, the community isn’t one of those.
Diogo Almeida [00:02:04]: And actually, in my ideal world, it would be, like, community all the time. I was thinking, “Should I host a town hall while walking to your studio?” And I’m like, “No, that’s too crazy.”
Swyx [00:02:12]: Sure. Yeah. Well, you guys have been hosting town halls on Discord. Discord is now 100,000 people. Your Twitter’s
Diogo Almeida [00:02:19]: I don’t follow these stats.
Diogo Almeida [00:02:20]: So holy shit.
Swyx [00:02:21]: Your Twitter’s blown up. It was, it was really funny ‘cause, like, at AIE, you were like, “Yeah, follow me please,” and then you didn’t, like, provide even your handle.
Diogo Almeida [00:02:29]: I’m a noob. I’m a noob.
Swyx [00:02:29]: You’re such a noob.
Diogo Almeida [00:02:30]: I’m a noob.
Swyx [00:02:31]: But no, but that, like, that’s, like, positive aura that, like
Diogo Almeida [00:02:33]: Cool
Swyx [00:02:33]: You don’t know how to promote yourself.
Diogo Almeida [00:02:35]: Yeah. Someone, like, called me out when I posted, like, “Holy shit, we’re all three twending-- trending topics.” And then they’re like, “That’s a personal feed.”
Swyx [00:02:42]: That’s a personal, yeah.
Diogo Almeida [00:02:43]: And I’m like, “Oh, no.”
Swyx [00:02:44]: Of course, of course it’ll trend to you.
Diogo Almeida [00:02:45]: Cringe. Yeah.
Swyx [00:02:45]: Yes, ‘cause it’s what you clicked on.
Diogo Almeida [00:02:47]: Yeah.
Swyx [00:02:47]: So okay. Let’s, Yeah, so congrats on everything.
What Is Jev? System 1 Models and Intelligence per Dollar
Diogo Almeida [00:02:50]: Thank you.
Swyx [00:02:50]: We’ll talk about more, details as you have them. But let’s, for people who are, like, living under a rock or just want, like, the definitive thing, what is Jev?
Diogo Almeida [00:03:02]: Whew. Let me think about. That’s a hard one.
Swyx [00:03:07]: Okay. And I’m happy to, like, re-ask if you wanna kind of
Diogo Almeida [00:03:09]: No. I’m happy to
Diogo Almeida [00:03:10]: I’m happy to, like, just jam on it.
Diogo Almeida [00:03:13]: I will say, like, the first thing that I’m relieved about with this question is now I don’t have to answer that question to my parents anymore ‘cause ChatGPT can just explain it.
Diogo Almeida [00:03:21]: So the way I see it is we new-- need a new class of models. We’re not attached to naming that class of models. Our-- the most accurate name we’ve come up with is System 1 models.
Diogo Almeida [00:03:33]: There will be reasons, but it’s-- there’s a reason

[truncated]

## Original Extract

The definitive Jev podcast with its lead creator.

Jev: System One models for Prod, not God — with Diogo Almeida, CEO, TypeSafe AI
Subscribe Sign in Latent Space: The AI Engineer Podcast Jev: System One models for Prod, not God — with Diogo Almeida, CEO, TypeSafe AI 15 1 1× 0:00 Current time: 0:00 / Total time: -2:20:52 -2:20:52 Audio playback is not supported on your browser. Please upgrade. Jev: System One models for Prod, not God — with Diogo Almeida, CEO, TypeSafe AI
15 1 Share Transcript Tickets for AIE NYC now open, and apply for the invite-only AIE CODE . Join us !
We have an unusual relationship with today’s guest: for years since coauthoring the InstructGPT paper , Diogo Almeida had been saying that API-available frontier models have been going down the wrong path, everything from the alignment to refusals to reliability perspectives, that we have dropped every mode other than autoregressive chat-tuned LLMs because of the overwhelming success of ChatGPT.
In a launch video now viewed ~40M times (by comparison, GPT4o was 22M , Fable 5 was 15M , Navier Stokes was 74M , and 6 Astra was 137M ), Diogo introduced Jev and it immediately took over the AI timeline — we’ll skip full Jev explainers because your favorite AI influencer/educator has probably already done one. We also collected:
the official patterns and cookbooks you should see first, from Allie
speed based - games and computer use
the voice + computer use example we discuss at 1h34 mins
guided responses in text messages
Jev for coding agents has an official guide
reasonable pushback from Theo - Diogo has published a note on the Tyranny of the KV Cache that you should read as a followup after the pod for Jev + coding agents, because of his belief that Cache Rules Everything
Programming Languages built atop Jev (Diogo’s fave)
Jev for analytics replay and user journey review
a core goal of Jev is to “disappear into the background” - eg as unremarkable as regex
blending transformers and classifiers
Jev vs GLiNER (note difference/pushback , agreed , agreed , agreed )
Instead we’ll focus on what we can uniquely offer — a broader philosophical and mission-based understanding of how and why Jev was created , and what you should expect next in terms of future models from TypeSafe ( ReasoningJev ?) and what usecases and ideas you should work on vs the 55th low effort clone of Jev’s API or doing a generic JevBench benchmark - something Diogo has rejected publicly .
Why RLCD: Three kinds of RLHF, and why they are ALL the wrong north star
Diogo knows a good deal about RLHF, given that he was on the team that pioneered post-training at OpenAI — and traces the three branches to Christiano et al 2017 (the robot backflip demo), Stiennon et al 2020 (learning to summarize) and his baby, Ouyang et al 2022 (InstructGPT). From there on, every innovation from Function Calling to Structured Outputs to Reasoning felt like a hack on top of the string based, sequence to sequence prediction paradigm. As he mentions on the pod, from 2023-2024 he struggled unsuccessfully, due to both personal and organization underestimation, to train a model that accurately addressed what he saw as the core problem with making LLMs the heart of software: reliability .
Jev’s core innovation is " Reinforcement Learning for Calibrated Decisions ”, a novel, unpublished technique that optimizes for “answers with epistemically honest probabilities on System One tasks” rather than human rated feedback (RLHF) — which causes hallucinations, sycophancy, and permanent reliance on humans — or programmatically verifiable outputs with rubrics (RLVR) — which solves Navier Stokes but exacerbates jagged intelligence and doesn’t integrate well with other software.
We’ve talked about the calibration problem before on the pod, but probably the single best place to understand why RLCD became necessary is Diogo’s AIE talk , which discusses why a generation of training helpful AI assistants for humans has impaired them for training models for composable, programmable AI for automation .
At the end he also teases his contrarian opinion on scaling laws - which teases how to build a modern neolab without the billions of dollars the major labs have…
The Bitterest Lesson: Tasks and Data beats Compute
We spend a good amount of time discussing Diogo’s essay on the Bitterest Lesson :
His point is that “You get what you optimize for and the bitterest lesson in ML is that the most important part of it isn’t ML at all.” - and picking the right north star, eg upvoting for user preference vs being integrated into tool calls - makes everything else fall in line.
We’re excited to catch up with a freshly dyed Diogo to discuss:
Why AI can solve extraordinarily hard problems but still fail to automate basic work
What System One Models are and why Jev is built for software rather than chat
RLHF, mode collapse , calibration, and the hidden costs of optimizing for human preferences
Why refusals become a problem when AI is buried inside software dependencies
Why TypeSafe rejects public benchmarks and optimizes for intelligence per dollar
The “bitterest lesson”: why the right task and the right data can matter more than compute
Why TypeSafe thinks of itself as a data lab rather than a model lab
RLCD vs. RLHF and RLVR as fundamentally different North Stars for AI
Why reliability and robustness matter more than simple determinism
Jev’s programming primitives and how intelligence maps into software control flow
Why developers should decompose AI workflows into small, measurable decisions
How structured state replaces giant prompts and system messages
Why Diogo thinks AI should eventually disappear into the background of software
The “inverse SaaS-pocalypse” and how AI could supercharge existing software
System One vs. System Two intelligence and the limits of reasoning models
Dark data, computer use, real-time intelligence, and Jev’s biggest early use cases
Why Jev could reshape coding agents built around a single-model architecture
Why Diogo says he wouldn’t pre-train with $1 billion
The OpenAI journey that led to TypeSafe and why he thinks many neo-labs are approaching AI incorrectly
Coding agents beyond the KV cache , shared state, sub-agents, and the multi-agent future
LinkedIn: https://www.linkedin.com/in/diogomda
X: https://x.com/CompleteSkeptic
TypeSafe AI: https://typesafe.ai/
00:00:00 Jev Launch Week and the AI Economic Revolution
00:02:50 What Is Jev? System One Models and Programmable AI
00:05:54 RLHF, Mode Collapse, Calibration, and Yann LeCun
00:10:29 Programmatic AI, Refusals, and Safety Alignment
00:17:21 Why TypeSafe Rejects Public Benchmarks
00:20:43 The Bitterest Lesson: Data, Compute, and the Right Task
00:24:59 RLCD vs. RLHF and RLVR
00:28:42 Why Powerful AI Still Hasn’t Automated the Economy
00:39:55 Reliability, Robustness, and Determinism
00:48:11 Model Versioning, LTS, Speed, and Intelligence per Dollar
00:54:04 Inside Jev’s API and Programming Primitives
00:58:28 How to Build with Jev: Structure, Decomposition, and Small Decisions
01:18:28 The Inverse SaaS-pocalypse and AI Disappearing into Software
01:33:21 Computer Use, Dark Data, and Jev’s Biggest Use Cases
01:38:48 How Jev Could Reshape Coding Agents
01:41:00 AI Safety, Frontier Pacing, and the Limits of RLVR
01:48:03 Why Diogo Wouldn’t Pre-Train with $1 Billion
01:55:19 The OpenAI Story Behind TypeSafe
02:01:41 Why Diogo Thinks Most Neo-Labs Are Getting AI Wrong
02:08:00 Coding Agents Beyond the KV Cache and the Multi-Agent Future
Introduction: Jev Launch Week and Developer Momentum
Swyx [00:00:00]: Okay, we’re in the studio. A special occasion because this week, Diogo, my good buddy, launched Jev, and it’s been taking over the complete timeline. How do you feel? What’s it like to be you right now?
Diogo Almeida [00:00:16]: Emotionally?
Diogo Almeida [00:00:17]: Never been worse. Like, I’m a ragged corpse of a person right now because there’s so much going on, and I’m like a technical CEO, so I have, like, a lot of fires to fight.
Diogo Almeida [00:00:29]: But mentally, I feel—I say this all the time, and I’ve been saying this kind of for years in my over-under events. Like, I feel like the entire AI field is like one of those, like, carnival house of mirrors, and everyone is just insane and saying the weirdest stuff that doesn’t make sense. And it feels like for just this week, like, I’m on a better in sync with reality and like, oh, people see it now. AI can be so much more than what was once thought.
Diogo Almeida [00:01:06]: And like, yes, we are going to make. Like, an AI-based economic revolution is back on the table, and this is fucking awesome.
Diogo Almeida [00:01:17]: I’m so jazzed the developers get it. It’s, it’s, Yeah, and I want to show my eternal gratitude to the developers and
Diogo Almeida [00:01:26]: I’m so jazzed about the community and everything. It’s so great.
Swyx [00:01:28]: Yeah, you were saying yesterday that you decided to prioritize the town hall and not a bunch of, like, VIP, investor-type people because you wanted to make sure that they are the people that you get your most, attention, right? The engineers, the developers.
Diogo Almeida [00:01:43]: Yeah, it felt a little like, oh man, I’m talking to, like, really important people right now.
Diogo Almeida [00:01:47]: I probably shouldn’t reveal who.
Diogo Almeida [00:01:48]: But it feels a little bit dirty for me to, I’m, like, perhaps overly genuine in things. Like, it feels, like, dirty if, like, in my gigantic calendar event of people to talk to, the community isn’t one of those.
Diogo Almeida [00:02:04]: And actually, in my ideal world, it would be, like, community all the time. I was thinking, “Should I host a town hall while walking to your studio?” And I’m like, “No, that’s too crazy.”
Swyx [00:02:12]: Sure. Yeah. Well, you guys have been hosting town halls on Discord. Discord is now 100,000 people. Your Twitter’s
Diogo Almeida [00:02:19]: I don’t follow these stats.
Diogo Almeida [00:02:20]: So holy shit.
Swyx [00:02:21]: Your Twitter’s blown up. It was, it was really funny ‘cause, like, at AIE, you were like, “Yeah, follow me please,” and then you didn’t, like, provide even your handle.
Diogo Almeida [00:02:29]: I’m a noob. I’m a noob.
Swyx [00:02:29]: You’re such a noob.
Diogo Almeida [00:02:30]: I’m a noob.
Swyx [00:02:31]: But no, but that, like, that’s, like, positive aura that, like
Diogo Almeida [00:02:33]: Cool
Swyx [00:02:33]: You don’t know how to promote yourself.
Diogo Almeida [00:02:35]: Yeah. Someone, like, called me out when I posted, like, “Holy shit, we’re all three twending-- trending topics.” And then they’re like, “That’s a personal feed.”
Swyx [00:02:42]: That’s a personal, yeah.
Diogo Almeida [00:02:43]: And I’m like, “Oh, no.”
Swyx [00:02:44]: Of course, of course it’ll trend to you.
Diogo Almeida [00:02:45]: Cringe. Yeah.
Swyx [00:02:45]: Yes, ‘cause it’s what you clicked on.
Diogo Almeida [00:02:47]: Yeah.
Swyx [00:02:47]: So okay. Let’s, Yeah, so congrats on everything.
What Is Jev? System 1 Models and Intelligence per Dollar
Diogo Almeida [00:02:50]: Thank you.
Swyx [00:02:50]: We’ll talk about more, details as you have them. But let’s, for people who are, like, living under a rock or just want, like, the definitive thing, what is Jev?
Diogo Almeida [00:03:02]: Whew. Let me think about. That’s a hard one.
Swyx [00:03:07]: Okay. And I’m happy to, like, re-ask if you wanna kind of
Diogo Almeida [00:03:09]: No. I’m happy to
Diogo Almeida [00:03:10]: I’m happy to, like, just jam on it.
Diogo Almeida [00:03:13]: I will say, like, the first thing that I’m relieved about with this question is now I don’t have to answer that question to my parents anymore ‘cause ChatGPT can just explain it.
Diogo Almeida [00:03:21]: So the way I see it is we new-- need a new class of models. We’re not attached to naming that class of models. Our-- the most accurate name we’ve come up with is System 1 models.
Diogo Almeida [00:03:33]: There will be reasons, but it’s-- there’s a reason

[truncated]
