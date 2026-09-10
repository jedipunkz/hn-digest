---
source: "https://thenextweb.com/news/anthropic-claude-misuse-threat-intelligence-report"
hn_url: "https://news.ycombinator.com/item?id=49651621"
title: "Anthropic details how Claude was misused for surveillance and weapons"
article_title: "Anthropic details how Claude was misused for surveillance and weapons"
image: "https://media.thenextweb.com/2026/08/Dario-Amodei-headshot.jpg"
author: "delichon"
captured_at: "2026-09-10T23:46:17Z"
capture_tool: "hn-digest"
hn_id: 49651621
score: 1
comments: 0
posted_at: "2026-09-10T23:40:52Z"
tags:
  - hacker-news
---

# Anthropic details how Claude was misused for surveillance and weapons

- HN: [49651621](https://news.ycombinator.com/item?id=49651621)
- Source: [thenextweb.com](https://thenextweb.com/news/anthropic-claude-misuse-threat-intelligence-report)
- Score: 1
- Comments: 0
- Posted: 2026-09-10T23:40:52Z

## Translation

Title: Anthropic details how Claude was misused for surveillance and weapons
Description: Anthropic's Claude misuse report details spying, weapons software and distillation cases tied to China, Russia, Iran and Yemen.

Article text:
Anthropic details how Claude was misused for surveillance and weapons
Skip to content
Toggle Navigation
News
Anthropic details how Claude was misused for surveillance and weapons
Anthropic says it detected and shut down efforts to misuse Claude for cyberattacks, state surveillance, weapons software and biological research. Its new threat report details cases tied to China, Russia, Iran and Yemen, and names Moonshot and DeepSeek among Chinese labs accused of copying its models.
Dario Amodei, CEO of Anthropic
Anthropic has published its most detailed account of Claude misuse, laying out how people tried to use its models for harm. The company said it detected and shut down these operations. They ran across cyberattacks, state surveillance, influence campaigns, conventional weapons work and biological research.
It set out the cases in a threat intelligence report released on 10 September. The report covers activity its Threat Intelligence team disrupted between December 2025 and August 2026.
The report spans seven harm areas. They are cyber operations, influence operations, surveillance, conventional weapons development, biological misuse, scams and fraud, and illicit distillation. Anthropic said the misuse involved its Claude Haiku, Sonnet and Opus models.
None of the cases involved its Fable or Mythos-class models, it added, with the exception of one distillation case. In each instance the company said it banned the accounts and strengthened its safeguards. It also shared intelligence with authorities and industry partners where appropriate.
The largest section covers cyber operations. Its central claim is that AI has narrowed the gap between well-resourced state hackers and lone operators. Anthropic pointed to a hacktivist using stolen API keys, financially motivated individuals, and a state espionage operator. Each ran multi-victim campaigns.
A year earlier, it said, that work would have required teams of skilled people. The company reported breaches completed in two to three hours, with dozens of victims handled in parallel by single operators.
One case, which Anthropic tracks as GTG-20006, involved a Russian-speaking operator running espionage against Ukrainian and European government targets. Diplomatic and defence organisations were among the targets too.
The company attributed it in line with public reporting on the group known as Midnight Blizzard. It said the actor used AI to check whether security products had flagged their malware, then automatically rebuilt it to slip past detection.
Anthropic traced a separate operation to two people it described as undergraduate students in Hunan, China. It said they ran what it called agent swarms against roughly fifty organisations. Those included a Southeast Asian government agency, from which the actor retrieved citizen records.
Surveillance built by a single consultant
Anthropic said state-aligned actors and commercial spyware vendors used Claude to build surveillance systems between January and July. The cases spanned China, Iran and West Africa. It presented one as the most striking. A single subscriber used Claude as the engineering workforce for a platform named Lakana 360.
The company assessed the subscriber to be a Bamako-based consultant working with Mali’s state intelligence service.
The system monitored roughly 25 million SIM cards across all three of the country’s mobile operators. Anthropic said it bypassed a legal requirement for a court order before an operator could disclose certain records. It also ran on local models on-premises, so banning the account did not affect the deployed system.
Other surveillance cases named Iranian actors deploying a malicious Firefox extension that harvested users’ identities from social networks. Anthropic also described a religious affairs intelligence unit in China. It said the unit had shrunk from many teams of analysts to a single office. That office now produces thousands of investigations a month with an AI assistant.
In another Chinese case, the company said Claude scored social media posts by political sensitivity and flagged people for what the operators termed control. The report follows earlier TNW coverage of AI turning up in police search tools .
Jacob Klein, who leads threat intelligence at Anthropic, told Axios that AI was making state surveillance cheaper and more efficient. It was not changing who governments target, he said.
“They’re effectively automating parts of the job within the intel apparatus,” he said. He added that the pattern was no longer theoretical. “Authoritarian states are using AI for surveillance, repression and influence operations today,” he told Axios.
The report documents what Anthropic described as a new form of misuse. That is the use of Claude to write software for conventional weapons. It detailed six cases, three in China, two in Russia and one in Yemen.
In the Yemen case, the company said a cell in the north of the country ran three weapons programmes. They included a guided rocket, a multi-stage ballistic missile with a stated range goal above 2,000 kilometres, and a set of missile variants that included a hypersonic glide vehicle.
Anthropic said the actors used Claude Code in place of software engineers to develop guidance and control code. It added that they test-fired a guided rocket, a test that appears to have failed.
The Russian cases included freelance actors working on an autonomous kamikaze drone swarm.
Biology, and the judgement calls
Anthropic said it presented five cases of people using its models in ways that could support biological weapons development. It stressed how hard those judgements are to make.
The company withheld the names of the institutions, the countries and the specific biological agents involved. It said the individuals were working scientists, and it did not assert that they intended harm.
In one example from May, it said a request for help writing a grant application involved gain-of-function research on the chikungunya virus. That work was intended to be carried out at a military research institute.
“You are not seeing someone in a comic book kind of way say, ‘Hey, I want to build a biological weapon to kill everybody,’” Klein told The New York Times. “It’s an incredibly nuanced situation.”
On the chikungunya case, he told the paper the company did not know whether the research was meant to be weaponised, but that a military institution doing gain-of-function research was concerning. Anthropic said older models such as Claude Opus 4 were well below the level where they could meaningfully assist such work.
It said it had launched more recent models with stronger safeguards. That shift echoes an earlier report in which the company ran its bioweapon classifiers off during testing. The distillation thread, meanwhile, runs back to White House claims that Moonshot distilled Anthropic’s Fable .
The Chinese distillation cases
The report also expands on illicit distillation. Anthropic defines that as covertly extracting a model’s capabilities to train a rival.
The company said it had identified campaigns from seven China-based labs, all targeting its generally available models rather than Mythos. It named Moonshot AI, which it said silently forwarded customer requests to Claude and displayed the responses as though they came from its own Kimi model.
Over one ten-day period, Anthropic said, Moonshot relayed almost 300,000 requests through a network of 5,380 fraudulent accounts. It attributed more than 23 million exchanges to Moonshot between May and July. The company made similar allegations against DeepSeek, Zhipu, Xiaomi, SenseTime and MiniMax.
Some of the relayed queries exposed sensitive user data, it said, in one instance live credentials tied to a Russian defence agency.
The distillation findings land alongside a US intelligence advisory naming six Chinese firms and Beijing’s rejection of those claims . Anthropic said most of the influence operations it caught drew little or no authentic engagement before it disrupted them. Other platforms have made the same caveat.
Anthropic said it published the report out of an obligation to disclose the misuse, and to give governments and civil society a clearer view of how such threats take shape.
With expertise in digital marketing, product management, and branding &amp; identity, Ana Maria Constantin develops strategies that resonate (show all)
With expertise in digital marketing, product management, and branding & identity, Ana Maria Constantin develops strategies that resonate with our target audience in the software/SaaS industry. Collaboration and teamwork are paramount to her, as she loves empowering her colleagues to achieve outstanding results and unlock their full potential.
Get the most important tech news in your inbox each week.
1
Andrew Tulloch, Meta’s star AI recruit, is leaving after less than a year
2
Anthropic skipped UK pre-release tests for Mythos 5.1, the FT reports
3
DeepSeek launches V4.1-Flash and retires V4-Pro, its flagship model
4
OpenAI’s Bubeck denies trying to cut Anthropic mathematician from credit
5
EU cybersecurity agency is now testing Mythos 5 and GPT-6 Astra, the Commission says
Anthropic details how Claude was misused for surveillance and weapons
Andrew Tulloch, Meta’s star AI recruit, is leaving after less than a year
Anthropic skipped UK pre-release tests for Mythos 5.1, the FT reports
DeepSeek launches V4.1-Flash and retires V4-Pro, its flagship model
OpenAI’s Bubeck denies trying to cut Anthropic mathematician from credit
Copyright © 2006—2026, Cogneve, INC. Made with <3 in Amsterdam.

## Original Extract

Anthropic's Claude misuse report details spying, weapons software and distillation cases tied to China, Russia, Iran and Yemen.

Anthropic details how Claude was misused for surveillance and weapons
Skip to content
Toggle Navigation
News
Anthropic details how Claude was misused for surveillance and weapons
Anthropic says it detected and shut down efforts to misuse Claude for cyberattacks, state surveillance, weapons software and biological research. Its new threat report details cases tied to China, Russia, Iran and Yemen, and names Moonshot and DeepSeek among Chinese labs accused of copying its models.
Dario Amodei, CEO of Anthropic
Anthropic has published its most detailed account of Claude misuse, laying out how people tried to use its models for harm. The company said it detected and shut down these operations. They ran across cyberattacks, state surveillance, influence campaigns, conventional weapons work and biological research.
It set out the cases in a threat intelligence report released on 10 September. The report covers activity its Threat Intelligence team disrupted between December 2025 and August 2026.
The report spans seven harm areas. They are cyber operations, influence operations, surveillance, conventional weapons development, biological misuse, scams and fraud, and illicit distillation. Anthropic said the misuse involved its Claude Haiku, Sonnet and Opus models.
None of the cases involved its Fable or Mythos-class models, it added, with the exception of one distillation case. In each instance the company said it banned the accounts and strengthened its safeguards. It also shared intelligence with authorities and industry partners where appropriate.
The largest section covers cyber operations. Its central claim is that AI has narrowed the gap between well-resourced state hackers and lone operators. Anthropic pointed to a hacktivist using stolen API keys, financially motivated individuals, and a state espionage operator. Each ran multi-victim campaigns.
A year earlier, it said, that work would have required teams of skilled people. The company reported breaches completed in two to three hours, with dozens of victims handled in parallel by single operators.
One case, which Anthropic tracks as GTG-20006, involved a Russian-speaking operator running espionage against Ukrainian and European government targets. Diplomatic and defence organisations were among the targets too.
The company attributed it in line with public reporting on the group known as Midnight Blizzard. It said the actor used AI to check whether security products had flagged their malware, then automatically rebuilt it to slip past detection.
Anthropic traced a separate operation to two people it described as undergraduate students in Hunan, China. It said they ran what it called agent swarms against roughly fifty organisations. Those included a Southeast Asian government agency, from which the actor retrieved citizen records.
Surveillance built by a single consultant
Anthropic said state-aligned actors and commercial spyware vendors used Claude to build surveillance systems between January and July. The cases spanned China, Iran and West Africa. It presented one as the most striking. A single subscriber used Claude as the engineering workforce for a platform named Lakana 360.
The company assessed the subscriber to be a Bamako-based consultant working with Mali’s state intelligence service.
The system monitored roughly 25 million SIM cards across all three of the country’s mobile operators. Anthropic said it bypassed a legal requirement for a court order before an operator could disclose certain records. It also ran on local models on-premises, so banning the account did not affect the deployed system.
Other surveillance cases named Iranian actors deploying a malicious Firefox extension that harvested users’ identities from social networks. Anthropic also described a religious affairs intelligence unit in China. It said the unit had shrunk from many teams of analysts to a single office. That office now produces thousands of investigations a month with an AI assistant.
In another Chinese case, the company said Claude scored social media posts by political sensitivity and flagged people for what the operators termed control. The report follows earlier TNW coverage of AI turning up in police search tools .
Jacob Klein, who leads threat intelligence at Anthropic, told Axios that AI was making state surveillance cheaper and more efficient. It was not changing who governments target, he said.
“They’re effectively automating parts of the job within the intel apparatus,” he said. He added that the pattern was no longer theoretical. “Authoritarian states are using AI for surveillance, repression and influence operations today,” he told Axios.
The report documents what Anthropic described as a new form of misuse. That is the use of Claude to write software for conventional weapons. It detailed six cases, three in China, two in Russia and one in Yemen.
In the Yemen case, the company said a cell in the north of the country ran three weapons programmes. They included a guided rocket, a multi-stage ballistic missile with a stated range goal above 2,000 kilometres, and a set of missile variants that included a hypersonic glide vehicle.
Anthropic said the actors used Claude Code in place of software engineers to develop guidance and control code. It added that they test-fired a guided rocket, a test that appears to have failed.
The Russian cases included freelance actors working on an autonomous kamikaze drone swarm.
Biology, and the judgement calls
Anthropic said it presented five cases of people using its models in ways that could support biological weapons development. It stressed how hard those judgements are to make.
The company withheld the names of the institutions, the countries and the specific biological agents involved. It said the individuals were working scientists, and it did not assert that they intended harm.
In one example from May, it said a request for help writing a grant application involved gain-of-function research on the chikungunya virus. That work was intended to be carried out at a military research institute.
“You are not seeing someone in a comic book kind of way say, ‘Hey, I want to build a biological weapon to kill everybody,’” Klein told The New York Times. “It’s an incredibly nuanced situation.”
On the chikungunya case, he told the paper the company did not know whether the research was meant to be weaponised, but that a military institution doing gain-of-function research was concerning. Anthropic said older models such as Claude Opus 4 were well below the level where they could meaningfully assist such work.
It said it had launched more recent models with stronger safeguards. That shift echoes an earlier report in which the company ran its bioweapon classifiers off during testing. The distillation thread, meanwhile, runs back to White House claims that Moonshot distilled Anthropic’s Fable .
The Chinese distillation cases
The report also expands on illicit distillation. Anthropic defines that as covertly extracting a model’s capabilities to train a rival.
The company said it had identified campaigns from seven China-based labs, all targeting its generally available models rather than Mythos. It named Moonshot AI, which it said silently forwarded customer requests to Claude and displayed the responses as though they came from its own Kimi model.
Over one ten-day period, Anthropic said, Moonshot relayed almost 300,000 requests through a network of 5,380 fraudulent accounts. It attributed more than 23 million exchanges to Moonshot between May and July. The company made similar allegations against DeepSeek, Zhipu, Xiaomi, SenseTime and MiniMax.
Some of the relayed queries exposed sensitive user data, it said, in one instance live credentials tied to a Russian defence agency.
The distillation findings land alongside a US intelligence advisory naming six Chinese firms and Beijing’s rejection of those claims . Anthropic said most of the influence operations it caught drew little or no authentic engagement before it disrupted them. Other platforms have made the same caveat.
Anthropic said it published the report out of an obligation to disclose the misuse, and to give governments and civil society a clearer view of how such threats take shape.
With expertise in digital marketing, product management, and branding &amp; identity, Ana Maria Constantin develops strategies that resonate (show all)
With expertise in digital marketing, product management, and branding & identity, Ana Maria Constantin develops strategies that resonate with our target audience in the software/SaaS industry. Collaboration and teamwork are paramount to her, as she loves empowering her colleagues to achieve outstanding results and unlock their full potential.
Get the most important tech news in your inbox each week.
1
Andrew Tulloch, Meta’s star AI recruit, is leaving after less than a year
2
Anthropic skipped UK pre-release tests for Mythos 5.1, the FT reports
3
DeepSeek launches V4.1-Flash and retires V4-Pro, its flagship model
4
OpenAI’s Bubeck denies trying to cut Anthropic mathematician from credit
5
EU cybersecurity agency is now testing Mythos 5 and GPT-6 Astra, the Commission says
Anthropic details how Claude was misused for surveillance and weapons
Andrew Tulloch, Meta’s star AI recruit, is leaving after less than a year
Anthropic skipped UK pre-release tests for Mythos 5.1, the FT reports
DeepSeek launches V4.1-Flash and retires V4-Pro, its flagship model
OpenAI’s Bubeck denies trying to cut Anthropic mathematician from credit
Copyright © 2006—2026, Cogneve, INC. Made with <3 in Amsterdam.
