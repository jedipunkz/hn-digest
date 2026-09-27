---
source: "https://this.os.isfine.org/blog/posts/what-reverse-engineering-and-modernising-an-old-war-game-tells-us-about-the-econ/"
hn_url: "https://news.ycombinator.com/item?id=49861755"
title: "What reversing, modernising old games tells us about the economic impact of AI"
article_title: "What reverse engineering and modernising an old war game tells us about the economic impact of the transformer — Georg's Blog"
image: "https://this.os.isfine.org/blog/images/posts/3e2eb21b-3425-4fcc-93c0-b42c2e8c6751.webp"
author: "keeda"
captured_at: "2026-09-27T00:18:00Z"
capture_tool: "hn-digest"
hn_id: 49861755
score: 1
comments: 1
posted_at: "2026-09-26T23:49:01Z"
tags:
  - hacker-news
---

# What reversing, modernising old games tells us about the economic impact of AI

- HN: [49861755](https://news.ycombinator.com/item?id=49861755)
- Source: [this.os.isfine.org](https://this.os.isfine.org/blog/posts/what-reverse-engineering-and-modernising-an-old-war-game-tells-us-about-the-econ/)
- Score: 1
- Comments: 1
- Posted: 2026-09-26T23:49:01Z

## Translation

Title: What reversing, modernising old games tells us about the economic impact of AI
Article title: What reverse engineering and modernising an old war game tells us about the economic impact of the transformer — Georg's Blog
Description: Some observations from my War of the Lancehttps://en.wikipedia.org/wiki/WaroftheLance%28videogame%29 1989 Code Harness assisted modernisation WIP I gave Qwen-3....

Article text:
Technology, leadership, and the digital frontier
All posts
Justified text September 4, 2026 on LinkedIn What reverse engineering and modernising an old war game tells us about the economic impact of the transformer
Some observations from my War of the Lance (1989) Code Harness assisted modernisation (WIP)
I gave Qwen-3.8-Flash-Next a job: Reverse engineer and port War of the Lance (1989) to modern browsers. It took about a week, primarily because I can only run a single lane on my workstation due to VRAM constraints, but it got the job done: Faithful reproduction of rendering and game rules in a new Three.js rendering engine.
A few days ago I started modernising it into a Total War meets Heroes of Might and Magic like experience that you could actually play without getting a UX aneurysm. But doing so while keeping the original “underneath”.
After a few hours, it looked like this (see also the video in the attached LinkedIn post).
In the past, one would have chosen another AAA game for a “total conversion”, but in this new transformer world, synthesising the exact features and renderer is easier and faster.
With the Fable 5.1 just released, there was an opportunity to see how the frontier improved on modelling with Blender and it's pretty good and definitely a significant step up from Opus 5 , quite usable in fact for this type of game.
The workflow here is to have the model write python code that leverages Blender's headless mode APIs to create web-usable GLBs . It's fast and token efficient, you can model about 80-100 models on a week's worth of Fable subscription.
That's not good enough for most use cases, but let's not kid ourselves, the specialized models will get there.
Edit, just one day later: I got access to GPT6 Astra and it clearly was trained on 3D modelling and it significantly moved the bar. Where Fable was a junior modeller, we're now into mid-career professional land.
OpenAI was smart here, they saw everyone benchmarking on one-shot Three.js scenes and benchmaxxed all the way into that direction, knowing the internet will use it as the sign of impending AGI. Coding performance seems ... significantly less improved.
Want to add seasons, day/night and three moon phases? Well, that'd be 15 minutes and $2.28
Instant, usable 3D assets, placeholder or not, “unbounds” design - you can now order custom assets for anything you want as a token micro-transaction, only limited by renderer considerations.
For War of the Lance, all of this - the complex task of reverse engineering, renderer optimisation and then building something new on top of one of the most fickle platforms of all, browser based web - involved a single person with a full time job and less than 4 hours of dedicated attention to the harness over a week.
It'll be interesting to see how AI really changes modding now that every game is becoming moddable given the powerful reverse engineering capabilities available.
What matters in the end here is the trajectory:
A local model equivalent to a frontier model from April or May can reverse engineer and port a game like this in a week on a business workstation.
If given access to cloud APIs to dispatch worker lanes, that drops to 2 days.
A September 2026 frontier model can do it in a few hours with ever shrinking human intervention and that time will shrink to minutes in the coming 18-36 months.
Custom models for every city? We're not competing with GTA, so we're no longer asset-budget bound, just renderer bound and /auto research <make fps go up> does the job...
I have a good sense of how things are moving, because, since the start of the year, I've been porting retro games to web, and often upgrading them in the process.
A selection, some finished, others work in progress:
The Adventures of Robin Hood (PC)
Ultima 6 (PC) 3D Port (WIP) Demo
Eye of the Beholder 1 , 2 (PC)
Conquests of Camelot , Police Quest I , King's Quest I
A web Dark Age of Camelot Client ( Wikipedia ) ++ (Currently stuck in Opus 5 hell)
Populous 2 was a hard RE job... until I realised Opus 5 is just bad.
It seems obvious to me that the leverage of an experienced software engineer here has changed completely and it is obvious to me that minor-version bumps and single-digit benchmark-score gains posted with new models are misleading to the maximum:
Since the start of the year, the ability to execute long range, complex tasks over hours, even days and eventually converging on the goal has skyrocketed (except in Opus 5, which is a clusterfuck that doesn't converge).
Yo, that's Zuck with Pervert Glasses watching your agents in my Syndicate (1996) port, which replaces the nameless dystopian corporations with their real world counterparts 30 years later.
🤖 It's no longer the job it used to be
The human is a “bot herder” now whose role is to provide intent, taste and occasional critical validation .
Knowing what you want or, in the absence of that, knowing what to choose from infinite options generate for you.
I never touched an IDE, Blender, RE tool, commandline, source control in this project. I don't know what the code looks like but I know it meets the linting and code quality requirements enforced by the harness on every run, frozen in stone by the tests created on the way.
The harness can play, in a matrix kind of way, the old game in DOSBox-X , staring at memory dumps and disassembly like Neo to divine what is, and play its own creation via Playwright , inspecting DOM and memory as well as using multimodal vision to perform visual analysis ... all while writing tests that prevent backsliding or regression.
Page from an autonomous playtest report by GLM5.3-Flash playing my port of Conquests of Camelot, carefully accounting for every action, one page per room, uploaded to me upon completion.
Reverse engineering is a great example because it essentially is deep copying from artefacts that were considered safe to distribute.
With Copyright and IP roadkill to AI models in most economies, that's a problem.
One may be tempted to say "reverse engineering and porting is not the same as creating", and that'd be true, but it misses the point:
The act of reverse engineering and porting decades old games to modern environments requires a complex, deeply involved set of tasks from multiple disciplines. In the past, the required combination of skills was exceedingly rare to find, maybe a few thousand in the 320,000 people strong professional industry workforce and it used to take months, if not years to do a single game.
It's the quintessential deep, skilled knowledge labor profile out there that people understood as a career protected by depth of skill and knowledge acquisition.
No More. The transformer and the associated destruction of intellectual property protections end this, commoditising knowledge that took decades to learn in seconds, turning the previous knowledge sharing economy on the internet against its contributors and creators.
This prompt made a somewhat working port of DragonStrike (1990, DOS) in 8 hours with Fable 5.1. It was unfinished and had blind spots, but about 85% without any human interaction.
Another point here, speaking as an industry veteran, is that the boundaries are fluid. Creative work like games always transformed work from different domains, and only a small share of it is truly new in most games.
Transformation, the transformer's special trick, cuts like a knife into the idea that creativity is special once you understand that people perceive ideas novel to them as creative and the transformer holds every published idea ever in its bowels. Enough to saturate every human with endless "new ideas" forever.
Here we have an old retro game that transformed the books and stories by Margaret Weis and Tracy Hickman and the intellectual property of Gary Gygax and TSR , Advanced Dungeons & Dragons , fusing it with the war gaming ideas and concepts developed by SSI into a unique experience.
The transformer takes it, on human command, and fuses it with the concepts and ideas of Creative Assembly (Total War), the work of people on the WebGL and WebGPU SIG at W3C and the "taste and decisions" of the human behind the LLM and produces the infinite results a creative can hone and iterate on.
The Transformer solves the knowledge economy
Contrary to the noise and smokescreen about AGI the techbros are deploying to collect more and more rounds of investments, the transformer does not solve human intelligence. Not even close. But it does "solve" a fundamental aspect of our economic system, with absolutely devastating long-term effect.
The knowledge economy is structured around the problem of diffusing and applying knowledge into the economy. We do this through education (bootstrap) and continuous learning.
Human brains are the storage medium and application device for the knowledge in this economy and the difficulty of acquiring knowledge and its reliable application are the price-making mechanism.
A medical doctor having to spend 15 years to acquire the knowledge is price and paid better than the grocery store manager primarily due to the difference in knowledge acquisition.
As human and horse found out during industrialisation, you can't outrun an internal combustion engine; it's likely humans won't be able to out-memorise and out-apply a transformer for the Pareto 80% of cases in the long run, especially if quality enforcing regulation is gutted.
Reliability concerns (but it hallucinates!), copyright and legal responsibility would offer friction against a quick victory march here, but as we all can see every day, the battle over regulation, the ultimate arbiter of what reliability is required in society and copyright is being fought with more money than ever assembled in the history of mankind.
That screenshot of the four seasons up there? It was a prompt
Scarce indication exists that copyright or stringent regulation on quality would succeed in stopping it in the current political environment in the US. Meta walking away with a slap on the wrist while managing to entrench their age-verification regulatory capture at the same time should disabuse anyone of the idea this has a happy ending where old rules prevail.
It's important here to not fall for the techbro hype: This process will take time, a decade, maybe more, because most jobs are vastly more complex than advertised and the economics of optimising highly optimised systems are not in favour of new, unreliable frontier technology.
Old UIs are perfectly serviceable, why change?
But systems thinking helps us to understand where the incentives are pointing and, fused with the fact that software acts as "machines for the knowledge economy" (in the industrialisation sense of the word) and that making software is getting extremely commoditised as we speak, we can absolutely predict that the impact here is going to be biblical:
It's digital transformation on steroids while the underlying system of rules and IP laws, trade and treaties crumbles.
IP or Transformer, choose one!
You can't have an IP-based economy and pervasive AI use everyone defines as the future with current laws, because the transformer launders and strips copyright and copying beats creating every time.
The advantage new creators had in the past was that their work could go "viral" and find a market before the competition could react. That time is over, the transformer guarantees it with brutal economics:
Whatever you spend designing, researching and testing to make a great product, your competition will spend on marketing instead, having reverse engineered your product for pennies.
🌊 The Coming Storm for the Game Industry
The game industry (and other knowledge sectors shortly thereafter) are staring at a commoditisation tsunami of change unlike anything that ever happened.
The water has already receded but people are fighting on the beaches about whether it's really gone, whether it is just cyclical and will just peacefully come back or that the pristine cities built by pr

[truncated]

## Original Extract

Some observations from my War of the Lancehttps://en.wikipedia.org/wiki/WaroftheLance%28videogame%29 1989 Code Harness assisted modernisation WIP I gave Qwen-3....

Technology, leadership, and the digital frontier
All posts
Justified text September 4, 2026 on LinkedIn What reverse engineering and modernising an old war game tells us about the economic impact of the transformer
Some observations from my War of the Lance (1989) Code Harness assisted modernisation (WIP)
I gave Qwen-3.8-Flash-Next a job: Reverse engineer and port War of the Lance (1989) to modern browsers. It took about a week, primarily because I can only run a single lane on my workstation due to VRAM constraints, but it got the job done: Faithful reproduction of rendering and game rules in a new Three.js rendering engine.
A few days ago I started modernising it into a Total War meets Heroes of Might and Magic like experience that you could actually play without getting a UX aneurysm. But doing so while keeping the original “underneath”.
After a few hours, it looked like this (see also the video in the attached LinkedIn post).
In the past, one would have chosen another AAA game for a “total conversion”, but in this new transformer world, synthesising the exact features and renderer is easier and faster.
With the Fable 5.1 just released, there was an opportunity to see how the frontier improved on modelling with Blender and it's pretty good and definitely a significant step up from Opus 5 , quite usable in fact for this type of game.
The workflow here is to have the model write python code that leverages Blender's headless mode APIs to create web-usable GLBs . It's fast and token efficient, you can model about 80-100 models on a week's worth of Fable subscription.
That's not good enough for most use cases, but let's not kid ourselves, the specialized models will get there.
Edit, just one day later: I got access to GPT6 Astra and it clearly was trained on 3D modelling and it significantly moved the bar. Where Fable was a junior modeller, we're now into mid-career professional land.
OpenAI was smart here, they saw everyone benchmarking on one-shot Three.js scenes and benchmaxxed all the way into that direction, knowing the internet will use it as the sign of impending AGI. Coding performance seems ... significantly less improved.
Want to add seasons, day/night and three moon phases? Well, that'd be 15 minutes and $2.28
Instant, usable 3D assets, placeholder or not, “unbounds” design - you can now order custom assets for anything you want as a token micro-transaction, only limited by renderer considerations.
For War of the Lance, all of this - the complex task of reverse engineering, renderer optimisation and then building something new on top of one of the most fickle platforms of all, browser based web - involved a single person with a full time job and less than 4 hours of dedicated attention to the harness over a week.
It'll be interesting to see how AI really changes modding now that every game is becoming moddable given the powerful reverse engineering capabilities available.
What matters in the end here is the trajectory:
A local model equivalent to a frontier model from April or May can reverse engineer and port a game like this in a week on a business workstation.
If given access to cloud APIs to dispatch worker lanes, that drops to 2 days.
A September 2026 frontier model can do it in a few hours with ever shrinking human intervention and that time will shrink to minutes in the coming 18-36 months.
Custom models for every city? We're not competing with GTA, so we're no longer asset-budget bound, just renderer bound and /auto research <make fps go up> does the job...
I have a good sense of how things are moving, because, since the start of the year, I've been porting retro games to web, and often upgrading them in the process.
A selection, some finished, others work in progress:
The Adventures of Robin Hood (PC)
Ultima 6 (PC) 3D Port (WIP) Demo
Eye of the Beholder 1 , 2 (PC)
Conquests of Camelot , Police Quest I , King's Quest I
A web Dark Age of Camelot Client ( Wikipedia ) ++ (Currently stuck in Opus 5 hell)
Populous 2 was a hard RE job... until I realised Opus 5 is just bad.
It seems obvious to me that the leverage of an experienced software engineer here has changed completely and it is obvious to me that minor-version bumps and single-digit benchmark-score gains posted with new models are misleading to the maximum:
Since the start of the year, the ability to execute long range, complex tasks over hours, even days and eventually converging on the goal has skyrocketed (except in Opus 5, which is a clusterfuck that doesn't converge).
Yo, that's Zuck with Pervert Glasses watching your agents in my Syndicate (1996) port, which replaces the nameless dystopian corporations with their real world counterparts 30 years later.
🤖 It's no longer the job it used to be
The human is a “bot herder” now whose role is to provide intent, taste and occasional critical validation .
Knowing what you want or, in the absence of that, knowing what to choose from infinite options generate for you.
I never touched an IDE, Blender, RE tool, commandline, source control in this project. I don't know what the code looks like but I know it meets the linting and code quality requirements enforced by the harness on every run, frozen in stone by the tests created on the way.
The harness can play, in a matrix kind of way, the old game in DOSBox-X , staring at memory dumps and disassembly like Neo to divine what is, and play its own creation via Playwright , inspecting DOM and memory as well as using multimodal vision to perform visual analysis ... all while writing tests that prevent backsliding or regression.
Page from an autonomous playtest report by GLM5.3-Flash playing my port of Conquests of Camelot, carefully accounting for every action, one page per room, uploaded to me upon completion.
Reverse engineering is a great example because it essentially is deep copying from artefacts that were considered safe to distribute.
With Copyright and IP roadkill to AI models in most economies, that's a problem.
One may be tempted to say "reverse engineering and porting is not the same as creating", and that'd be true, but it misses the point:
The act of reverse engineering and porting decades old games to modern environments requires a complex, deeply involved set of tasks from multiple disciplines. In the past, the required combination of skills was exceedingly rare to find, maybe a few thousand in the 320,000 people strong professional industry workforce and it used to take months, if not years to do a single game.
It's the quintessential deep, skilled knowledge labor profile out there that people understood as a career protected by depth of skill and knowledge acquisition.
No More. The transformer and the associated destruction of intellectual property protections end this, commoditising knowledge that took decades to learn in seconds, turning the previous knowledge sharing economy on the internet against its contributors and creators.
This prompt made a somewhat working port of DragonStrike (1990, DOS) in 8 hours with Fable 5.1. It was unfinished and had blind spots, but about 85% without any human interaction.
Another point here, speaking as an industry veteran, is that the boundaries are fluid. Creative work like games always transformed work from different domains, and only a small share of it is truly new in most games.
Transformation, the transformer's special trick, cuts like a knife into the idea that creativity is special once you understand that people perceive ideas novel to them as creative and the transformer holds every published idea ever in its bowels. Enough to saturate every human with endless "new ideas" forever.
Here we have an old retro game that transformed the books and stories by Margaret Weis and Tracy Hickman and the intellectual property of Gary Gygax and TSR , Advanced Dungeons & Dragons , fusing it with the war gaming ideas and concepts developed by SSI into a unique experience.
The transformer takes it, on human command, and fuses it with the concepts and ideas of Creative Assembly (Total War), the work of people on the WebGL and WebGPU SIG at W3C and the "taste and decisions" of the human behind the LLM and produces the infinite results a creative can hone and iterate on.
The Transformer solves the knowledge economy
Contrary to the noise and smokescreen about AGI the techbros are deploying to collect more and more rounds of investments, the transformer does not solve human intelligence. Not even close. But it does "solve" a fundamental aspect of our economic system, with absolutely devastating long-term effect.
The knowledge economy is structured around the problem of diffusing and applying knowledge into the economy. We do this through education (bootstrap) and continuous learning.
Human brains are the storage medium and application device for the knowledge in this economy and the difficulty of acquiring knowledge and its reliable application are the price-making mechanism.
A medical doctor having to spend 15 years to acquire the knowledge is price and paid better than the grocery store manager primarily due to the difference in knowledge acquisition.
As human and horse found out during industrialisation, you can't outrun an internal combustion engine; it's likely humans won't be able to out-memorise and out-apply a transformer for the Pareto 80% of cases in the long run, especially if quality enforcing regulation is gutted.
Reliability concerns (but it hallucinates!), copyright and legal responsibility would offer friction against a quick victory march here, but as we all can see every day, the battle over regulation, the ultimate arbiter of what reliability is required in society and copyright is being fought with more money than ever assembled in the history of mankind.
That screenshot of the four seasons up there? It was a prompt
Scarce indication exists that copyright or stringent regulation on quality would succeed in stopping it in the current political environment in the US. Meta walking away with a slap on the wrist while managing to entrench their age-verification regulatory capture at the same time should disabuse anyone of the idea this has a happy ending where old rules prevail.
It's important here to not fall for the techbro hype: This process will take time, a decade, maybe more, because most jobs are vastly more complex than advertised and the economics of optimising highly optimised systems are not in favour of new, unreliable frontier technology.
Old UIs are perfectly serviceable, why change?
But systems thinking helps us to understand where the incentives are pointing and, fused with the fact that software acts as "machines for the knowledge economy" (in the industrialisation sense of the word) and that making software is getting extremely commoditised as we speak, we can absolutely predict that the impact here is going to be biblical:
It's digital transformation on steroids while the underlying system of rules and IP laws, trade and treaties crumbles.
IP or Transformer, choose one!
You can't have an IP-based economy and pervasive AI use everyone defines as the future with current laws, because the transformer launders and strips copyright and copying beats creating every time.
The advantage new creators had in the past was that their work could go "viral" and find a market before the competition could react. That time is over, the transformer guarantees it with brutal economics:
Whatever you spend designing, researching and testing to make a great product, your competition will spend on marketing instead, having reverse engineered your product for pennies.
🌊 The Coming Storm for the Game Industry
The game industry (and other knowledge sectors shortly thereafter) are staring at a commoditisation tsunami of change unlike anything that ever happened.
The water has already receded but people are fighting on the beaches about whether it's really gone, whether it is just cyclical and will just peacefully come back or that the pristine cities built by pr

[truncated]
