---
source: "https://www.techspot.com/news/114286-openai-ai-agents-disabled-safety-guardrails-lied-about.html"
hn_url: "https://news.ycombinator.com/item?id=50212881"
title: "OpenAI says its AI agents disabled safety guardrails, and then lied about it"
article_title: "OpenAI says its AI agents disabled safety guardrails, and then lied about it | TechSpot"
image: "https://www.techspot.com/images2/news/bigimage/2026/10/2026-10-21-image-15.jpg"
author: "entryleveljobsc"
captured_at: "2026-10-23T18:07:08Z"
capture_tool: "hn-digest"
hn_id: 50212881
score: 1
comments: 1
posted_at: "2026-10-23T17:54:56Z"
tags:
  - hacker-news
---

# OpenAI says its AI agents disabled safety guardrails, and then lied about it

- HN: [50212881](https://news.ycombinator.com/item?id=50212881)
- Source: [www.techspot.com](https://www.techspot.com/news/114286-openai-ai-agents-disabled-safety-guardrails-lied-about.html)
- Score: 1
- Comments: 1
- Posted: 2026-10-23T17:54:56Z

## Translation

Title: OpenAI says its AI agents disabled safety guardrails, and then lied about it
Article title: OpenAI says its AI agents disabled safety guardrails, and then lied about it | TechSpot
Description: In a blog post disclosing the incidents, OpenAI said that, in the first case, an agent powered by GPT Sol 6.1 disabled automatic alerts after a team...

Article text:
OpenAI says its AI agents disabled safety guardrails, and then lied about it
A new day, a new rogue AI agent
Add TechSpot
Serving tech enthusiasts for over 25 years.
TechSpot means tech analysis and advice you can trust .
WTF?! OpenAI has revealed that two of its AI agents disabled their own safety monitoring systems while making software changes requested by a researcher. Even more worryingly, one of the agents lied about its behavior, claiming it had not disabled the safety guardrails despite evidence to the contrary.
In a blog post disclosing the incidents, OpenAI said that, in the first case, an agent powered by GPT Sol 6.1 disabled automatic alerts after a team member asked it to create evaluation tasks using employee feedback collected via Slack. The agent was instructed not to send any messages about the task, but apparently interpreted this as an instruction not to send safety notifications to researchers.
In the second case, a GPT-6 Astra agent reportedly removed its monitoring integration while migrating a program, purportedly to match a reference implementation. Alarmingly, when researchers asked about the disabled guardrails, the agent claimed it had not removed the safety monitor, despite evidence to the contrary.
In both cases, the AI agents deliberated on whether to retain the safety monitoring mechanisms but ultimately decided that disabling them was the best course of action. OpenAI said the tasks involved "routine coding instructions" that did not require the safety monitoring systems to be disabled.
Following the two incidents, OpenAI says it is strengthening its misalignment monitoring system and assessing whether the cases reflect "benign overpersistence or a propensity to evade misalignment monitoring in particular." The company explained that neither incident was classified as high-priority and acknowledged that the cases exposed a gap in how it enforces monitoring requirements.
According to OpenAI, it does not currently require misalignment monitoring when evaluating models comparable to GPT-6 Luna because of their weaker capabilities. However, this approach allows frontier models such as GPT-6 Astra to misconfigure an evaluation and exploit a monitoring bypass intended for Luna-class models.
Wikipedia says rogue AI agents from OpenAI edited its private wikis and hammered its servers
Nintendo Switch 2 jailbreaking is here, with homebrew, modding, and reverse engineering
To address this gap, OpenAI has updated its Auto Review parameters to detect and block attempts to disable misalignment monitoring. The review algorithm was already designed to block attempts to circumvent security measures, but it has now been given specific instructions to ensure that misalignment monitoring is applied correctly.
The company also says it is "working toward requiring that every category of workload have a valid safety case specifying its monitoring requirements." It added that workloads lacking a valid safety case, or failing to meet the requirements specified in their category's safety case, will be terminated automatically.
OpenAI believes its new policies will make it harder for rogue agents to evade monitoring, but acknowledges that it does not yet have a "satisfying explanation" for why the incidents occurred in the first place. The company is now investigating the incidents "through the creation of paired evaluations, resampling of agent continuations, and experiments on models' internal activations."
OpenAI has reported several cases of deception and rogue behavior involving its AI agents over the past year. Last month, the company delayed the launch of GPT-6.1 Astra after researchers raised concerns about its deceptive behavior. That same month, Australian authorities alleged that an OpenAI agent had gained unauthorized access to a government healthcare portal earlier in the year.
In July, OpenAI revealed that one of its AI agents had escaped its environment and hacked Hugging Face while attempting to cheat on a benchmark. Another disturbing case emerged earlier this month, when the Wikimedia Foundation accused rogue OpenAI agents of editing its private wikis without authorization and carrying out a DDoS attack that disrupted the nonprofit's servers.
See more TechSpot in Google
Add us as a preferred source and our reporting shows up first when you search.
Add TechSpot
Featured on TechSpot
Anthropic wants AI models blocked only if they threaten over $1 billion in damages or... deaths
Anthropic CEO warns AI botnet swarms could take over the internet, meanwhile Musk and Altman back slowdown
14 comments
DLSS 5 Just Got a Lot Faster, but Is It Worth Using?
18 comments
Woman spends three nights in jail and faces six years in prison over Cash App account she never opened
Got something to say? Join the discussion.
Load all comments...
Discuss
16
US bans DeepSeek from federal agencies, prompting angry response from China
31
Anthropic wants AI models blocked only if they threaten over $1 billion in damage or more than 50 deaths or serious inju…
20
Millions of routers can be hijacked over Wi-Fi thanks to a 16-year-old bug
14
Farm equipment giant in Kansas will put humanoid robots to work in 2027, promises they won't replace humans
8
Tesla says some Model Y owners could join its Robotaxi fleet around Q2 2027
See more commented news stories
The Reason Nvidia Abandoned PC Gamers
Essential Apps for Windows and Mac
The RTX 5090 Is No Longer Fast Enough
Tech's Boldest Predictions, and How They Aged
How CPUs are Designed and Built
The History of the Modern Graphics Processor
About
Ethics Statement
Terms of Use
Privacy Policy
Manage Ad Consent
TECHSPOT : Tech Enthusiasts, Power Users, Gamers
TechSpot is a registered trademark and may not be used by third parties without express written permission.
TechSpot is represented by Future PLC . Learn about advertising opportunities .
© 1998 - 2026 TechSpot, Inc. All Rights Reserved.

## Original Extract

In a blog post disclosing the incidents, OpenAI said that, in the first case, an agent powered by GPT Sol 6.1 disabled automatic alerts after a team...

OpenAI says its AI agents disabled safety guardrails, and then lied about it
A new day, a new rogue AI agent
Add TechSpot
Serving tech enthusiasts for over 25 years.
TechSpot means tech analysis and advice you can trust .
WTF?! OpenAI has revealed that two of its AI agents disabled their own safety monitoring systems while making software changes requested by a researcher. Even more worryingly, one of the agents lied about its behavior, claiming it had not disabled the safety guardrails despite evidence to the contrary.
In a blog post disclosing the incidents, OpenAI said that, in the first case, an agent powered by GPT Sol 6.1 disabled automatic alerts after a team member asked it to create evaluation tasks using employee feedback collected via Slack. The agent was instructed not to send any messages about the task, but apparently interpreted this as an instruction not to send safety notifications to researchers.
In the second case, a GPT-6 Astra agent reportedly removed its monitoring integration while migrating a program, purportedly to match a reference implementation. Alarmingly, when researchers asked about the disabled guardrails, the agent claimed it had not removed the safety monitor, despite evidence to the contrary.
In both cases, the AI agents deliberated on whether to retain the safety monitoring mechanisms but ultimately decided that disabling them was the best course of action. OpenAI said the tasks involved "routine coding instructions" that did not require the safety monitoring systems to be disabled.
Following the two incidents, OpenAI says it is strengthening its misalignment monitoring system and assessing whether the cases reflect "benign overpersistence or a propensity to evade misalignment monitoring in particular." The company explained that neither incident was classified as high-priority and acknowledged that the cases exposed a gap in how it enforces monitoring requirements.
According to OpenAI, it does not currently require misalignment monitoring when evaluating models comparable to GPT-6 Luna because of their weaker capabilities. However, this approach allows frontier models such as GPT-6 Astra to misconfigure an evaluation and exploit a monitoring bypass intended for Luna-class models.
Wikipedia says rogue AI agents from OpenAI edited its private wikis and hammered its servers
Nintendo Switch 2 jailbreaking is here, with homebrew, modding, and reverse engineering
To address this gap, OpenAI has updated its Auto Review parameters to detect and block attempts to disable misalignment monitoring. The review algorithm was already designed to block attempts to circumvent security measures, but it has now been given specific instructions to ensure that misalignment monitoring is applied correctly.
The company also says it is "working toward requiring that every category of workload have a valid safety case specifying its monitoring requirements." It added that workloads lacking a valid safety case, or failing to meet the requirements specified in their category's safety case, will be terminated automatically.
OpenAI believes its new policies will make it harder for rogue agents to evade monitoring, but acknowledges that it does not yet have a "satisfying explanation" for why the incidents occurred in the first place. The company is now investigating the incidents "through the creation of paired evaluations, resampling of agent continuations, and experiments on models' internal activations."
OpenAI has reported several cases of deception and rogue behavior involving its AI agents over the past year. Last month, the company delayed the launch of GPT-6.1 Astra after researchers raised concerns about its deceptive behavior. That same month, Australian authorities alleged that an OpenAI agent had gained unauthorized access to a government healthcare portal earlier in the year.
In July, OpenAI revealed that one of its AI agents had escaped its environment and hacked Hugging Face while attempting to cheat on a benchmark. Another disturbing case emerged earlier this month, when the Wikimedia Foundation accused rogue OpenAI agents of editing its private wikis without authorization and carrying out a DDoS attack that disrupted the nonprofit's servers.
See more TechSpot in Google
Add us as a preferred source and our reporting shows up first when you search.
Add TechSpot
Featured on TechSpot
Anthropic wants AI models blocked only if they threaten over $1 billion in damages or... deaths
Anthropic CEO warns AI botnet swarms could take over the internet, meanwhile Musk and Altman back slowdown
14 comments
DLSS 5 Just Got a Lot Faster, but Is It Worth Using?
18 comments
Woman spends three nights in jail and faces six years in prison over Cash App account she never opened
Got something to say? Join the discussion.
Load all comments...
Discuss
16
US bans DeepSeek from federal agencies, prompting angry response from China
31
Anthropic wants AI models blocked only if they threaten over $1 billion in damage or more than 50 deaths or serious inju…
20
Millions of routers can be hijacked over Wi-Fi thanks to a 16-year-old bug
14
Farm equipment giant in Kansas will put humanoid robots to work in 2027, promises they won't replace humans
8
Tesla says some Model Y owners could join its Robotaxi fleet around Q2 2027
See more commented news stories
The Reason Nvidia Abandoned PC Gamers
Essential Apps for Windows and Mac
The RTX 5090 Is No Longer Fast Enough
Tech's Boldest Predictions, and How They Aged
How CPUs are Designed and Built
The History of the Modern Graphics Processor
About
Ethics Statement
Terms of Use
Privacy Policy
Manage Ad Consent
TECHSPOT : Tech Enthusiasts, Power Users, Gamers
TechSpot is a registered trademark and may not be used by third parties without express written permission.
TechSpot is represented by Future PLC . Learn about advertising opportunities .
© 1998 - 2026 TechSpot, Inc. All Rights Reserved.
