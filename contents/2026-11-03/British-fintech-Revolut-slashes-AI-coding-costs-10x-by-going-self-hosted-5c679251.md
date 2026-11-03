---
source: "https://www.theregister.com/ai-and-ml/2026/11/03/british-fintech-revolut-slashes-ai-coding-costs-10x-by-cutting-out-claude-and-going-self-hosted/5307345"
hn_url: "https://news.ycombinator.com/item?id=50358388"
title: "British fintech Revolut slashes AI coding costs 10x by going self-hosted"
article_title: "British fintech Revolut slashes AI coding costs 10x by cutting out Claude and going self-hosted"
image: "https://image.theregister.com/260288.jpg?imageId=260288&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "gr_norm"
captured_at: "2026-11-03T21:32:51Z"
capture_tool: "hn-digest"
hn_id: 50358388
score: 5
comments: 1
posted_at: "2026-11-03T21:23:42Z"
tags:
  - hacker-news
---

# British fintech Revolut slashes AI coding costs 10x by going self-hosted

- HN: [50358388](https://news.ycombinator.com/item?id=50358388)
- Source: [www.theregister.com](https://www.theregister.com/ai-and-ml/2026/11/03/british-fintech-revolut-slashes-ai-coding-costs-10x-by-cutting-out-claude-and-going-self-hosted/5307345)
- Score: 5
- Comments: 1
- Posted: 2026-11-03T21:23:42Z

## Translation

Title: British fintech Revolut slashes AI coding costs 10x by going self-hosted
Article title: British fintech Revolut slashes AI coding costs 10x by cutting out Claude and going self-hosted
Description: Turns out cheap tokens get engineers doing things that save time, like debugging and tackling Jira tickets

Article text:
Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
British fintech Revolut slashes AI coding costs 10x by cutting out Claude and going self-hosted
Turns out cheap tokens get engineers doing things that save time, like debugging and tackling Jira tickets
Tobias Mann
Tobias
Mann
SYSTEMS EDITOR
Published
tue 3 Nov 2026 // 20:07 UTC
Make us preferred on Google
READ MORE
Nvidia is happy to take your money whether you're using open or closed models, says Jensen Huang
17 minutes ago
Drax wants to feed Britain's AI habit with £20B datacenter at biomass power station
7 hours ago
Toyota shows off SoC that stores model weights and its OS on integrated memory
17 hours ago
Anthropic's compute tab could top $1T as Q3 revenue nears $17B
21 hours ago
Marvell wants to be a copper topper, bets $115B on photonics leader to push optical networking
1 day ago
Revolut says that in less than a year it's managed to slash its AI coding inference costs by roughly 90 percent by making self-hosted models its default instead of relying primarily on commercial services like Claude.
Speaking at the AI Engineer conference in New York last month, David Abitbol, ops and infrastructure lead at the British fintech, described how it saw a 60 percent improvement in coding time and a 70 percent improvement in lead time following its rollout of AI coding tools to roughly 2,000 engineers. Revolut initially adopted GLM-5.1 before moving to GLM-5.3, which it now serves using 56 Nvidia B200 GPUs.
Prior to its switch to GLM, Revolut had been experimenting with AI, deploying Gemini in early 2025 and Claude Code in early 2026. "We were highly skeptical about the cost of AI. We were seeing engineers spend $1,000 a month, $10,000 a month with not so much ROI, and so when we launched it we launched it with strict budget limits," Abitbol said.
Because of this, engineers only had access to a limited number of tokens, so the company began investigating self-hosting iyd own models. At $0.30 per million input/output tokens, self hosting GLM 5.1 cost about a tenth of using OpenAI's Codex or Anthropic's Claude Code. And, as it turned out, when tokens are cheap, you stop worrying about how many they use or what they use them for.
While Revolut did see improvements to coding speed, Abitbol noted the biggest bottleneck wasn't actually the code. Instead, he says the biggest gains came from letting the AI models deal with things like debugging and Jira and Confluence tickets. "When you have unlimited tokens, you're free to do that," he said. "You no longer have to save your tokens for the end of the quarter or the most difficult task."
Developed by Chinese open weights model dev Z.ai, GLM 5.1 isn't exactly a small model. At roughly 754 billion total parameters (40 billion activated for each token generated), the model requires roughly 380 GB of GPU memory just to hold the weights and another 6-26 GB per user to maintain context depending on precision.
To serve GLM-5.3 at scale, Revolut deployed 56 Nvidia B200 GPUs, equivalent in GPU count to seven eight-GPU DGX B200 systems. But as Abitbol noted, the greatest performance gains didn't come from adding GPUs, but rather tuning the model's hyperparameters along with switching to NVFP4 quantization, implementing disaggregated prefill and decode, and tuning context lengths.
One key takeaway was that supporting a one-million-token context window doesn't necessarily make such a large window useful or efficient. "Long context is a scam. You really need to be mindful of that," Abitbol said. "Even though you can serve one million tokens, 256 [thousand] is really what you need."
With these optimizations, Abitbol says that the company was able to achieve sub 3-second time to first token on context windows of around 256,000 tokens, and generation speeds of around 200 tokens a second, both of which he notes are the sweet spot for the company's software engineers.
"What you want to aim for is a sub-3 second time to first token. This is empirical evidence. As soon as you exceed 5 seconds, my Slack inbox starts to blow up," Abitbol said. "In terms of throughput, it doesn't really matter as much. We serve, in terms of interactivity, 200 tokens per second. People are fine [with that]. You don't need to go higher. The way people work has changed. We no longer look at the stream of thinking tokens. People run tasks in parallel and check in when the task is done."
And if throughput is a concern, he notes engineers who sign in after 5 pm are getting closer to 800 tokens a second.
However, it's worth noting that GLM hasn't replaced GPT or Claude entirely. The Chinese model is text only, so it isn't ideal for things like UI work. So instead, Revolut used a model router to direct vision queries to models like Gemini or GPT Luna while keeping code generation running on GLM.
"Open source is now a priority at Revolut, and our economics are favorable. The billing is flat, we don't have any bill shock at the end of the month, and the frontier is the exception," Abitbol said.
Having said that, Abitbol acknowledges that Chinese open weights models may be a concern for privacy and security reasons, but emphasizes that the major concern there is accessing them via APIs running out of China. When run on-prem or on regional cloud resources, there's really nothing to worry about. ®
Nvidia is happy to take your money whether you're using open or closed models, says Jensen Huang
At Fireworks Forge, a celebration of specialized models
British fintech Revolut slashes AI coding costs 10x by cutting out Claude and going self-hosted
Turns out cheap tokens get engineers doing things that save time, like debugging and tackling Jira tickets
From expertise to intelligence: NISCO deploys world's first swarm-intelligence steel model at scale
PARTNER CONTENT: 300+ AI applications delivering surges in core metrics and annual cost reduction of over CNY100m
AI policy wonks declare 'sovereign agents' a threat to civilization
These are AI agents that operate outside of human control and 'pursue open-ended, evolving survival'
OPINION
Microsoft's moolah mountain is more massive than ever. It just doesn't want the people who built it
Flash the cash, flush the cache
GSMA warns of growing risks to mobile networks from extreme weather
Mobile comms body wants a joined-up approach from governments and regulators to mitigate downtime
Landline phones will switch off across Britland in 100 days, no extensions – Openreach
Cops bust botnet of 1.3M infected Android TV boxes
Anthropic now lets Claude users jump out of plane without parachute
Linus Torvalds says Linux kernel is out of 'stabilization phase' and using 'much healthier model' for AI contributions
Google struggling to make AI Mode searches pay
Voyager 1 gives controllers the silent treatment
Drax wants to feed Britain's AI habit with £20B datacenter at biomass power station
But first it needs planning permission, government approval and, rather importantly, a customer
Anthropic's compute tab could top $1T as Q3 revenue nears $17B
Revised IPO filing warns 12x revenue growth is unlikely to continue
Marvell wants to be a copper topper, bets $115B on photonics leader to push optical networking
IP house believes optics demand is going to explode as Nvidia and others eye bigger systems
Flash the cash, flush the cache
Amazon sends chills through its workforce with mixed message memo about AI use
A lot of software engineers feel bleak about the future of the profession
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Architect creates open source replacement for VM migration tool VMware has hidden
OpenVDDK is reportedly already being used in production, at scale
AlpineBSD announced – an AI-free FOSS OS
Combining the NetBSD kernel with the Alpine Linux userland
Ubuntu 27.04 may bring a preview of the long-awaited Core Desktop
Canonical has a long way to catch up, but it knows who it's chasing
Systemd adds support for Rust
What's even better than Rust in the Linux kernel? Rust in systemd!
FreeBSD founder ports Omarchy to FreeBSD in ten days, ten releases
Daemonarchy is unaffiliated with the Linux original, built with AI help, and could complicate FreeBSD's debate over AI contributions
SparkyOS: Lightweight general-purpose Arch, on a pendrive
Interesting new project from the minds behind SparkyLinux
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.

## Original Extract

Turns out cheap tokens get engineers doing things that save time, like debugging and tackling Jira tickets

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
British fintech Revolut slashes AI coding costs 10x by cutting out Claude and going self-hosted
Turns out cheap tokens get engineers doing things that save time, like debugging and tackling Jira tickets
Tobias Mann
Tobias
Mann
SYSTEMS EDITOR
Published
tue 3 Nov 2026 // 20:07 UTC
Make us preferred on Google
READ MORE
Nvidia is happy to take your money whether you're using open or closed models, says Jensen Huang
17 minutes ago
Drax wants to feed Britain's AI habit with £20B datacenter at biomass power station
7 hours ago
Toyota shows off SoC that stores model weights and its OS on integrated memory
17 hours ago
Anthropic's compute tab could top $1T as Q3 revenue nears $17B
21 hours ago
Marvell wants to be a copper topper, bets $115B on photonics leader to push optical networking
1 day ago
Revolut says that in less than a year it's managed to slash its AI coding inference costs by roughly 90 percent by making self-hosted models its default instead of relying primarily on commercial services like Claude.
Speaking at the AI Engineer conference in New York last month, David Abitbol, ops and infrastructure lead at the British fintech, described how it saw a 60 percent improvement in coding time and a 70 percent improvement in lead time following its rollout of AI coding tools to roughly 2,000 engineers. Revolut initially adopted GLM-5.1 before moving to GLM-5.3, which it now serves using 56 Nvidia B200 GPUs.
Prior to its switch to GLM, Revolut had been experimenting with AI, deploying Gemini in early 2025 and Claude Code in early 2026. "We were highly skeptical about the cost of AI. We were seeing engineers spend $1,000 a month, $10,000 a month with not so much ROI, and so when we launched it we launched it with strict budget limits," Abitbol said.
Because of this, engineers only had access to a limited number of tokens, so the company began investigating self-hosting iyd own models. At $0.30 per million input/output tokens, self hosting GLM 5.1 cost about a tenth of using OpenAI's Codex or Anthropic's Claude Code. And, as it turned out, when tokens are cheap, you stop worrying about how many they use or what they use them for.
While Revolut did see improvements to coding speed, Abitbol noted the biggest bottleneck wasn't actually the code. Instead, he says the biggest gains came from letting the AI models deal with things like debugging and Jira and Confluence tickets. "When you have unlimited tokens, you're free to do that," he said. "You no longer have to save your tokens for the end of the quarter or the most difficult task."
Developed by Chinese open weights model dev Z.ai, GLM 5.1 isn't exactly a small model. At roughly 754 billion total parameters (40 billion activated for each token generated), the model requires roughly 380 GB of GPU memory just to hold the weights and another 6-26 GB per user to maintain context depending on precision.
To serve GLM-5.3 at scale, Revolut deployed 56 Nvidia B200 GPUs, equivalent in GPU count to seven eight-GPU DGX B200 systems. But as Abitbol noted, the greatest performance gains didn't come from adding GPUs, but rather tuning the model's hyperparameters along with switching to NVFP4 quantization, implementing disaggregated prefill and decode, and tuning context lengths.
One key takeaway was that supporting a one-million-token context window doesn't necessarily make such a large window useful or efficient. "Long context is a scam. You really need to be mindful of that," Abitbol said. "Even though you can serve one million tokens, 256 [thousand] is really what you need."
With these optimizations, Abitbol says that the company was able to achieve sub 3-second time to first token on context windows of around 256,000 tokens, and generation speeds of around 200 tokens a second, both of which he notes are the sweet spot for the company's software engineers.
"What you want to aim for is a sub-3 second time to first token. This is empirical evidence. As soon as you exceed 5 seconds, my Slack inbox starts to blow up," Abitbol said. "In terms of throughput, it doesn't really matter as much. We serve, in terms of interactivity, 200 tokens per second. People are fine [with that]. You don't need to go higher. The way people work has changed. We no longer look at the stream of thinking tokens. People run tasks in parallel and check in when the task is done."
And if throughput is a concern, he notes engineers who sign in after 5 pm are getting closer to 800 tokens a second.
However, it's worth noting that GLM hasn't replaced GPT or Claude entirely. The Chinese model is text only, so it isn't ideal for things like UI work. So instead, Revolut used a model router to direct vision queries to models like Gemini or GPT Luna while keeping code generation running on GLM.
"Open source is now a priority at Revolut, and our economics are favorable. The billing is flat, we don't have any bill shock at the end of the month, and the frontier is the exception," Abitbol said.
Having said that, Abitbol acknowledges that Chinese open weights models may be a concern for privacy and security reasons, but emphasizes that the major concern there is accessing them via APIs running out of China. When run on-prem or on regional cloud resources, there's really nothing to worry about. ®
Nvidia is happy to take your money whether you're using open or closed models, says Jensen Huang
At Fireworks Forge, a celebration of specialized models
British fintech Revolut slashes AI coding costs 10x by cutting out Claude and going self-hosted
Turns out cheap tokens get engineers doing things that save time, like debugging and tackling Jira tickets
From expertise to intelligence: NISCO deploys world's first swarm-intelligence steel model at scale
PARTNER CONTENT: 300+ AI applications delivering surges in core metrics and annual cost reduction of over CNY100m
AI policy wonks declare 'sovereign agents' a threat to civilization
These are AI agents that operate outside of human control and 'pursue open-ended, evolving survival'
OPINION
Microsoft's moolah mountain is more massive than ever. It just doesn't want the people who built it
Flash the cash, flush the cache
GSMA warns of growing risks to mobile networks from extreme weather
Mobile comms body wants a joined-up approach from governments and regulators to mitigate downtime
Landline phones will switch off across Britland in 100 days, no extensions – Openreach
Cops bust botnet of 1.3M infected Android TV boxes
Anthropic now lets Claude users jump out of plane without parachute
Linus Torvalds says Linux kernel is out of 'stabilization phase' and using 'much healthier model' for AI contributions
Google struggling to make AI Mode searches pay
Voyager 1 gives controllers the silent treatment
Drax wants to feed Britain's AI habit with £20B datacenter at biomass power station
But first it needs planning permission, government approval and, rather importantly, a customer
Anthropic's compute tab could top $1T as Q3 revenue nears $17B
Revised IPO filing warns 12x revenue growth is unlikely to continue
Marvell wants to be a copper topper, bets $115B on photonics leader to push optical networking
IP house believes optics demand is going to explode as Nvidia and others eye bigger systems
Flash the cash, flush the cache
Amazon sends chills through its workforce with mixed message memo about AI use
A lot of software engineers feel bleak about the future of the profession
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Architect creates open source replacement for VM migration tool VMware has hidden
OpenVDDK is reportedly already being used in production, at scale
AlpineBSD announced – an AI-free FOSS OS
Combining the NetBSD kernel with the Alpine Linux userland
Ubuntu 27.04 may bring a preview of the long-awaited Core Desktop
Canonical has a long way to catch up, but it knows who it's chasing
Systemd adds support for Rust
What's even better than Rust in the Linux kernel? Rust in systemd!
FreeBSD founder ports Omarchy to FreeBSD in ten days, ten releases
Daemonarchy is unaffiliated with the Linux original, built with AI help, and could complicate FreeBSD's debate over AI contributions
SparkyOS: Lightweight general-purpose Arch, on a pendrive
Interesting new project from the minds behind SparkyLinux
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
