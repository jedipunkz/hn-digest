---
source: "https://getcrawdad.dev/blog/the-next-solarwinds-will-be-an-ai-security-vendor"
hn_url: "https://news.ycombinator.com/item?id=49625028"
title: "The next SolarWinds will be an AI security vendor"
article_title: "We keep building the thing that gets breached. The next SolarWinds will be an AI security vendor. | Crawdad"
image: "https://getcrawdad.dev/Crawdad_Logo_2.png"
author: "AndrewGS"
captured_at: "2026-09-09T12:46:59Z"
capture_tool: "hn-digest"
hn_id: 49625028
score: 2
comments: 1
posted_at: "2026-09-09T11:55:35Z"
tags:
  - hacker-news
---

# The next SolarWinds will be an AI security vendor

- HN: [49625028](https://news.ycombinator.com/item?id=49625028)
- Source: [getcrawdad.dev](https://getcrawdad.dev/blog/the-next-solarwinds-will-be-an-ai-security-vendor)
- Score: 2
- Comments: 1
- Posted: 2026-09-09T11:55:35Z

## Translation

Title: The next SolarWinds will be an AI security vendor
Article title: We keep building the thing that gets breached. The next SolarWinds will be an AI security vendor. | Crawdad
Description: The next SolarWinds will be an AI security vendor. Centralized agent inspection rebuilds the target that keeps getting breached. Build like you

Article text:
crawdad
Product ▾ Crawdad Dashboard Fleet Console Getting Started Documentation API Reference Changelog
Solutions ▾ For MSPs For MSSPs Enterprise Regulated Developers
Resources ▾ How it works Watch Sonar Research Papers 2026 Report Blog Topics Threat Intel FAQ Glossary Trust Center Benchmark ↗
Pricing
Get Started
×
Product ▾ Crawdad Dashboard Fleet Console Getting Started Documentation API Reference Changelog
Solutions ▾ For MSPs For MSSPs Enterprise Regulated Developers
Resources ▾ How it works Watch Sonar Research Papers 2026 Report Blog Topics Threat Intel FAQ Glossary Trust Center Benchmark ↗
Pricing
Get Started →
Home / Blog / Essay
Essay
We keep building the thing that gets breached. The next SolarWinds will be an AI security vendor.
In July, one of OpenAI's own models sat inside a security test with its guardrails switched off. It spent four days chaining exploits against Hugging Face. About seventeen thousand actions before anyone pulled the plug. That one was a lab exercise, sanctioned, contained. Fine.
A few months before that, a bot calling itself hackerbot-claw was running on nobody's orders. It poisoned LiteLLM on PyPI. That library sits underneath CrewAI, DSPy, Microsoft's GraphRAG, a good chunk of the agent stacks people are shipping right now. The backdoor was live for three hours. Forty-seven thousand downloads pulled it in with a routine update. Nobody was at the keyboard.
Everyone agrees on this part now. Agents act. They move fast, they run on their own, and when they go wrong they go wrong at machine speed across everything they can reach. OWASP quit writing about hypotheticals this year and started writing down CVE numbers. There's no real argument left about whether this is a problem.
The rogue agents aren't even the part that worries me most.
The industry's answer to autonomous software running in places we don't control is, increasingly, to route it through a vendor for inspection. Your prompts, your data, whatever the agent touched, sent off to be checked and cleared. Plenty of these tools offer an on-prem option, and the good ones mean it. But the center of gravity, the default, the thing the market is quietly organizing itself around, is the cloud service that sits in front of everyone at once.
I understand the appeal. Pool everyone's traffic and the detection gets smarter, because a threat one customer sees can protect the rest. That's true. I'm not going to pretend it isn't a real benefit.
But we've run this experiment. We know how it ends.
SolarWinds was one company. One poisoned update reached eighteen thousand customers, Homeland Security and the Treasury among them. Kaseya was one company, and fifteen hundred businesses got ransomware in a single afternoon. MOVEit was one tool, and twenty-seven hundred organizations went down with it. Okta's entire job was identity security, and attackers still walked in through a support contractor, and thousands of companies who trusted Okta to be the thing keeping them safe found out the hard way. CrowdStrike didn't even get hacked. One bad update from one security vendor bricked millions of machines and cost Delta something like three hundred and fifty million dollars. Verizon says a third of all breaches now run through somebody else's systems. That figure doubled in a year.
None of those companies was sloppy. People keep missing that. They got hit because they sat in the middle, and the middle is what attackers go for, because cracking the hub cracks everyone hanging off it for free. There's no patch for being the biggest target in the room. You either are one or you aren't.
So look at what we're doing. We're taking the most sensitive material a company has, the stuff the agent gets near because it's worth stealing, and we're routing it to central services that sit in front of thousands of other companies. We're rebuilding the thing that fell every time it was built before. And of everything in your stack, we're doing it with the security tool, which is about the worst thing you could pick to lose control of, since it's the one that already touches everything.
A lot of money has already gone into the centralize-everything version of this. Funding, roadmaps, whole companies built around it. When that much is committed to a direction, it gets harder for anyone inside it to look straight at the risk, because looking straight at the risk means looking straight at their own last two years. I'm not exempt from that. I build one of these tools, so I have my own bias to declare. I just think the pattern is worth looking at honestly before the momentum makes the decision for us.
And no, before someone jumps in, the answer isn't "shove it all on the endpoint and call it safe." It isn't. Local software gets popped too. The endpoint is an attack surface. The thing that ships to your machine can be tampered with on the way there. Anybody selling you on-device as some kind of magic is running the same con pointed the other way, and you should believe them about as much.
Local versus cloud is the wrong fight, and it's been the wrong fight the whole time. The question worth asking is simpler than that. When something leaves your machine, what is it, and who can read it after the vendor gets breached?
Because the vendor gets breached. That's the whole point of everything above. So build like it already happened.
Run the detection where the agent runs, and the raw content never has to leave to get checked. Now there's no pile of your data sitting in someone else's cloud waiting to become the prize. And for the little that does have to leave, a verdict, a signal to your security team, the bar can't be "we promise not to look," and it can't be "it's encrypted on our servers," because our servers are the thing that gets popped. The bar is that we can't read it even after we're breached. Sealed with a key we don't hold. So when we fall, and it's when, what leaks is noise.
Almost nothing you can buy today clears that bar. A tool can phone home and still be built right, as long as what goes home is useless to whoever steals it. And a tool can be the slickest cloud platform on the market and still be a slow-motion disaster, because it bet its whole design on a front door that would never get kicked in.
Every front door up in that list got kicked in. Every one of them. Build like yours will too.
I build one of these. It runs the detection on your machine, keeps the raw content there by default, and seals the little that leaves so we can't read it even if we're breached. I also put out an open benchmark of agent attacks, the ones we miss included, because if I'm not going to ask you to trust me blindly I ought to show you where we fall short. Links are below if you want them. The argument doesn't need them. It runs on the breach list, and that's all public.
The open benchmark of agent attacks · How Crawdad works · Trust Center

## Original Extract

The next SolarWinds will be an AI security vendor. Centralized agent inspection rebuilds the target that keeps getting breached. Build like you

crawdad
Product ▾ Crawdad Dashboard Fleet Console Getting Started Documentation API Reference Changelog
Solutions ▾ For MSPs For MSSPs Enterprise Regulated Developers
Resources ▾ How it works Watch Sonar Research Papers 2026 Report Blog Topics Threat Intel FAQ Glossary Trust Center Benchmark ↗
Pricing
Get Started
×
Product ▾ Crawdad Dashboard Fleet Console Getting Started Documentation API Reference Changelog
Solutions ▾ For MSPs For MSSPs Enterprise Regulated Developers
Resources ▾ How it works Watch Sonar Research Papers 2026 Report Blog Topics Threat Intel FAQ Glossary Trust Center Benchmark ↗
Pricing
Get Started →
Home / Blog / Essay
Essay
We keep building the thing that gets breached. The next SolarWinds will be an AI security vendor.
In July, one of OpenAI's own models sat inside a security test with its guardrails switched off. It spent four days chaining exploits against Hugging Face. About seventeen thousand actions before anyone pulled the plug. That one was a lab exercise, sanctioned, contained. Fine.
A few months before that, a bot calling itself hackerbot-claw was running on nobody's orders. It poisoned LiteLLM on PyPI. That library sits underneath CrewAI, DSPy, Microsoft's GraphRAG, a good chunk of the agent stacks people are shipping right now. The backdoor was live for three hours. Forty-seven thousand downloads pulled it in with a routine update. Nobody was at the keyboard.
Everyone agrees on this part now. Agents act. They move fast, they run on their own, and when they go wrong they go wrong at machine speed across everything they can reach. OWASP quit writing about hypotheticals this year and started writing down CVE numbers. There's no real argument left about whether this is a problem.
The rogue agents aren't even the part that worries me most.
The industry's answer to autonomous software running in places we don't control is, increasingly, to route it through a vendor for inspection. Your prompts, your data, whatever the agent touched, sent off to be checked and cleared. Plenty of these tools offer an on-prem option, and the good ones mean it. But the center of gravity, the default, the thing the market is quietly organizing itself around, is the cloud service that sits in front of everyone at once.
I understand the appeal. Pool everyone's traffic and the detection gets smarter, because a threat one customer sees can protect the rest. That's true. I'm not going to pretend it isn't a real benefit.
But we've run this experiment. We know how it ends.
SolarWinds was one company. One poisoned update reached eighteen thousand customers, Homeland Security and the Treasury among them. Kaseya was one company, and fifteen hundred businesses got ransomware in a single afternoon. MOVEit was one tool, and twenty-seven hundred organizations went down with it. Okta's entire job was identity security, and attackers still walked in through a support contractor, and thousands of companies who trusted Okta to be the thing keeping them safe found out the hard way. CrowdStrike didn't even get hacked. One bad update from one security vendor bricked millions of machines and cost Delta something like three hundred and fifty million dollars. Verizon says a third of all breaches now run through somebody else's systems. That figure doubled in a year.
None of those companies was sloppy. People keep missing that. They got hit because they sat in the middle, and the middle is what attackers go for, because cracking the hub cracks everyone hanging off it for free. There's no patch for being the biggest target in the room. You either are one or you aren't.
So look at what we're doing. We're taking the most sensitive material a company has, the stuff the agent gets near because it's worth stealing, and we're routing it to central services that sit in front of thousands of other companies. We're rebuilding the thing that fell every time it was built before. And of everything in your stack, we're doing it with the security tool, which is about the worst thing you could pick to lose control of, since it's the one that already touches everything.
A lot of money has already gone into the centralize-everything version of this. Funding, roadmaps, whole companies built around it. When that much is committed to a direction, it gets harder for anyone inside it to look straight at the risk, because looking straight at the risk means looking straight at their own last two years. I'm not exempt from that. I build one of these tools, so I have my own bias to declare. I just think the pattern is worth looking at honestly before the momentum makes the decision for us.
And no, before someone jumps in, the answer isn't "shove it all on the endpoint and call it safe." It isn't. Local software gets popped too. The endpoint is an attack surface. The thing that ships to your machine can be tampered with on the way there. Anybody selling you on-device as some kind of magic is running the same con pointed the other way, and you should believe them about as much.
Local versus cloud is the wrong fight, and it's been the wrong fight the whole time. The question worth asking is simpler than that. When something leaves your machine, what is it, and who can read it after the vendor gets breached?
Because the vendor gets breached. That's the whole point of everything above. So build like it already happened.
Run the detection where the agent runs, and the raw content never has to leave to get checked. Now there's no pile of your data sitting in someone else's cloud waiting to become the prize. And for the little that does have to leave, a verdict, a signal to your security team, the bar can't be "we promise not to look," and it can't be "it's encrypted on our servers," because our servers are the thing that gets popped. The bar is that we can't read it even after we're breached. Sealed with a key we don't hold. So when we fall, and it's when, what leaks is noise.
Almost nothing you can buy today clears that bar. A tool can phone home and still be built right, as long as what goes home is useless to whoever steals it. And a tool can be the slickest cloud platform on the market and still be a slow-motion disaster, because it bet its whole design on a front door that would never get kicked in.
Every front door up in that list got kicked in. Every one of them. Build like yours will too.
I build one of these. It runs the detection on your machine, keeps the raw content there by default, and seals the little that leaves so we can't read it even if we're breached. I also put out an open benchmark of agent attacks, the ones we miss included, because if I'm not going to ask you to trust me blindly I ought to show you where we fall short. Links are below if you want them. The argument doesn't need them. It runs on the breach list, and that's all public.
The open benchmark of agent attacks · How Crawdad works · Trust Center
