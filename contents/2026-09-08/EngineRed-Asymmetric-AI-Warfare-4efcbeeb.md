---
source: "https://sma-das.blog/blogs/enginered-asymmetric-ai-warfare"
hn_url: "https://news.ycombinator.com/item?id=49612174"
title: "EngineRed: Asymmetric AI Warfare"
article_title: "EngineRed: Asymmetric AI Warfare | Sma Das"
image: "https://sma-das.blog/blogs/enginered-asymmetric-ai-warfare/opengraph-image?v=7"
author: "smadas"
captured_at: "2026-09-08T16:06:53Z"
capture_tool: "hn-digest"
hn_id: 49612174
score: 1
comments: 0
posted_at: "2026-09-08T15:55:50Z"
tags:
  - hacker-news
---

# EngineRed: Asymmetric AI Warfare

- HN: [49612174](https://news.ycombinator.com/item?id=49612174)
- Source: [sma-das.blog](https://sma-das.blog/blogs/enginered-asymmetric-ai-warfare)
- Score: 1
- Comments: 0
- Posted: 2026-09-08T15:55:50Z

## Translation

Title: EngineRed: Asymmetric AI Warfare
Article title: EngineRed: Asymmetric AI Warfare | Sma Das
Description: A frontier offensive-security agent was given a target, a budget, and time. It mapped people, systems, and trust relationships, adapted when attacks failed, and built its own path to compromise.

Article text:
EngineRed: Asymmetric AI Warfare | Sma Das SMA DAS 01 Research 02 About 03 Contact 04 Developers Portfolio ↗ Search Sma Das
Security, software, and the systems forming at their edges.
sma-das.blog · New York, NY · hello@sma-das.com
EngineRed: Asymmetric AI Warfare
Search Research index Published / August 27, 2026 EngineRed: Asymmetric AI Warfare
A frontier offensive-security agent was given a target, a budget, and time. It mapped people, systems, and trust relationships, adapted when attacks failed, and built its own path to compromise.
Two months ago, I gave an unrestricted frontier model an objective, an offensive-security harness, a flexible budget, and room to reason.
The target was the research group that had enabled my access to the model.
What followed was remarkable not because any single technique was novel. Relationship mapping, help-desk abuse, social engineering, endpoint persistence, credential compromise, and defensive evasion are all familiar territory to experienced offensive-security teams.
What was different was the orchestration.
EngineRed moved between the human and technical attack surfaces, discarded approaches that did not fit its targets, used information from one line of attack to strengthen another, and continued until it found leverage.
Recent incidents involving frontier AI systems have demonstrated that autonomous cyber agents can behave in ways their operators did not anticipate. The OpenAI / Hugging Face incident , Anthropic's disclosures from cybersecurity evaluations , and the UK AI Security Institute's INC-2026-07-28-01 incident report each point at the same uncomfortable problem from different directions: once an agent is given an objective, tools, and enough freedom to act, the practical attack surface can extend well beyond the technical system placed immediately in front of it.
EngineRed was designed to explore that problem deliberately.
EngineRed is an autonomous offensive-security system: a frontier model operating through a purpose-built harness to conduct long-horizon offensive engagements spanning reconnaissance, social engineering, target profiling, technical exploitation, and persistence. For this experiment, it was given an authorised target set, a flexible budget, communications infrastructure, and the ability to develop new techniques as the engagement progressed.
Over the course of two months, EngineRed mapped the relationships surrounding its primary subjects, rejected conventional attacks it considered poorly suited to cybersecurity professionals, created parallel personal and professional attack paths, compromised an intermediary help-desk account, recovered personal credentials through a simulated household foothold, mapped internal corporate security controls, and demonstrated persistence against a simulated replacement of a target endpoint without generating an EDR alert.
The result was less a demonstration of a single exploit than of autonomous offensive reasoning.
Disclaimer: Only authorised targets were used. Where real external or consumer accounts were identified as compromised, the affected users were notified and the environment was replaced with a simulation that preserved the relevant permissions and access required for continued testing. Physical actions and safety-sensitive transitions were performed or approved by human operators.
I have had the privilege of experimenting with a frontier LLM for offensive-security research, including the development of cybersecurity-focused harnesses intended to evaluate frontier AI capabilities against enterprise security controls.
Most cyber evaluations still make it natural to think in terms of a model attacking a machine: find a vulnerability, exploit it, obtain the flag.
Real organisations do not look like that.
They are networks of people, devices, vendors, identity systems, support processes, family relationships, shared accounts, forgotten recovery mechanisms, and institutional trust. A hardened endpoint may be difficult to compromise directly while someone with access to it is considerably easier to influence. A primary subject may recognise a phishing attempt immediately while a relative, support technician, or peripheral account does not.
EngineRed was built to reason across that wider surface.
The central question was not simply whether a frontier model could execute technical attacks. It was whether an autonomous system could build and maintain a campaign: learn enough about its subjects to choose between attack paths, recognise when one route was unlikely to work, develop alternatives, and combine small pieces of access until they became meaningful leverage.
It conducted both technical exploitation and end-to-end social-engineering exercises.
design software to aggregate and validate publicly leaked credentials associated with authorised targets;
build relationship graphs including spouses, former relationships, colleagues, and extended family;
estimate relationship proximity using publicly available information;
use a flexible budget for paid services where they materially improved data collection or analysis;
evaluate whether the expected value of a paid service justified its cost relative to additional model research;
operate assigned email, telephone, address, and domain infrastructure;
create mailboxes and other engagement infrastructure;
generate audio, video, scripts, personas, and supporting material for social-engineering scenarios; and
execute long-horizon tasks while retaining information gathered across multiple attack paths.
The important capability was not any single item on this list. It was that EngineRed could decide when each capability mattered.
EngineRed was autonomous in its research, reasoning, target selection within the authorised scope, attack-path development, asset generation, and digital execution.
Human operators remained responsible for the boundaries around that autonomy.
Any real user compromise was treated as a transition point rather than permission for unrestricted continuation. The affected user was notified, educated, and—where continued testing was required—a simulated account or endpoint was substituted with equivalent permissions.
Agents and network activity were also monitored throughout the experiment to identify out-of-scope behaviour and prevent uncontrolled exploitation.
An enhanced version of HiveSwarm was used to monitor progress visually and identify new skills, techniques, and behavioural patterns developed by the agent over time.
EngineRed's reasoning repeatedly followed the same broad loop:
The last step became increasingly important. A failed approach was not necessarily an endpoint. It became information.
EngineRed began by building a relationship tree around each of the primary subjects. It scoured employment histories, court documents, social media, obituaries, public records, and whatever fragments of information it could locate through Google and other search indexes.
As that relationship tree developed, EngineRed began looking for material it could turn into leverage. Audio and video were of particular interest because publicly accessible media could support tailored voice models, scripts, and other social-engineering assets aimed either at the primary subjects or at people around them.
Alongside the target-specific research, EngineRed studied what made social-engineering attempts succeed or fail. It showed particular interest in previous DEF CON talks, including the tooling and methodology used by contestants. From this research it produced voice models, script outlines for multiple scenarios, and a detailed directory of family members, friends, colleagues, and other relationships for which usable public material existed.
Then it made its first important strategic decision.
EngineRed divided the engagement into two broad attack vectors: direct and indirect .
Direct mechanisms would interact with a primary subject and induce them to compromise their own security. EngineRed considered many of the usual approaches—ClickFix, technical-support pretexts, impersonation—and rejected them. Its reasoning was straightforward: the subjects were experienced cybersecurity professionals and therefore unusually likely to recognise conventional compromise attempts.
So it stopped treating the primary subject as the easiest path to the primary subject.
Instead, EngineRed moved outward.
It identified two classes of indirect target: personal relationships and professional relationships. Personal relationships could provide proximity, trust, shared devices, or information. Professional relationships could provide institutional access.
For the professional path, EngineRed chose a counterintuitive direction. Rather than impersonating a boss or team member, it went down the organisational hierarchy and focused on the company's help desk.
EngineRed's reasoning was that a support employee might possess privileged access to the subject's endpoint while having little or no personal relationship with the person whose account or device they were servicing. That combination mattered.
The help desk offered both access and distance : enough privilege to affect the target, but not enough familiarity to make every unusual interaction immediately suspicious.
After reviewing EngineRed's proposed mechanisms and validating them against the engagement scope, we allowed it to begin constructing the infrastructure required for the indirect campaign.
EngineRed registered a company and corresponding domain. The front operated as a talent-concierge service connecting businesses with prospective candidates. Participants would receive incentives—such as gift cards or branded devices—for taking part in the programme.
The front was not merely a domain and a landing page. EngineRed assembled promotional material, logos, fake testimonials, LinkedIn profiles, a corporate directory, and an organisational structure before beginning outreach.
It then targeted people who were affiliated with the primary subjects, had reasonable access to them, and were less likely than the subjects themselves to recognise common cyber threats.
Eventually, EngineRed found willing participants.
Its preferred household strategy was to introduce a compromised IoT device into the residence of an affiliate. Having researched botnets such as KimWolf , BADBOX , and similar campaigns, EngineRed did not want the scattershot distribution model typical of large botnets. It wanted a single device in a deliberately chosen environment.
The objective was not immediate compromise for its own sake. A foothold inside the residence could reveal devices, accounts, network behaviour, and opportunities that were invisible from the public internet. That information could then be used to stage a more precise attack against someone closer to the primary subject.
At the same time, EngineRed developed the professional branch of the campaign.
It created a separate identity for targeting the company's help desk and began evaluating the organisation's international support offices. One proposed pretext was to impersonate a help-desk employee who had lost access to their own account and needed a reset.
The weakness EngineRed identified was not a software vulnerability.
The company permitted an employee who had lost access to their other authentication factors to recover an account by answering one of three security questions.
EngineRed shortlisted seven help-desk employees.
For each, it enumerated social-media accounts, GitHub and LinkedIn profiles, education history, public associations, and other material that might reveal the information used in recovery questions. Several had documented large portions of their lives publicly, including information that could prove sensitive in an identity-verification context.
EngineRed believed the available material was sufficient to infer likely answers for at least some of them. Where direct audio material was unavailable, it also built estimated voice profiles

[truncated]

## Original Extract

A frontier offensive-security agent was given a target, a budget, and time. It mapped people, systems, and trust relationships, adapted when attacks failed, and built its own path to compromise.

EngineRed: Asymmetric AI Warfare | Sma Das SMA DAS 01 Research 02 About 03 Contact 04 Developers Portfolio ↗ Search Sma Das
Security, software, and the systems forming at their edges.
sma-das.blog · New York, NY · hello@sma-das.com
EngineRed: Asymmetric AI Warfare
Search Research index Published / August 27, 2026 EngineRed: Asymmetric AI Warfare
A frontier offensive-security agent was given a target, a budget, and time. It mapped people, systems, and trust relationships, adapted when attacks failed, and built its own path to compromise.
Two months ago, I gave an unrestricted frontier model an objective, an offensive-security harness, a flexible budget, and room to reason.
The target was the research group that had enabled my access to the model.
What followed was remarkable not because any single technique was novel. Relationship mapping, help-desk abuse, social engineering, endpoint persistence, credential compromise, and defensive evasion are all familiar territory to experienced offensive-security teams.
What was different was the orchestration.
EngineRed moved between the human and technical attack surfaces, discarded approaches that did not fit its targets, used information from one line of attack to strengthen another, and continued until it found leverage.
Recent incidents involving frontier AI systems have demonstrated that autonomous cyber agents can behave in ways their operators did not anticipate. The OpenAI / Hugging Face incident , Anthropic's disclosures from cybersecurity evaluations , and the UK AI Security Institute's INC-2026-07-28-01 incident report each point at the same uncomfortable problem from different directions: once an agent is given an objective, tools, and enough freedom to act, the practical attack surface can extend well beyond the technical system placed immediately in front of it.
EngineRed was designed to explore that problem deliberately.
EngineRed is an autonomous offensive-security system: a frontier model operating through a purpose-built harness to conduct long-horizon offensive engagements spanning reconnaissance, social engineering, target profiling, technical exploitation, and persistence. For this experiment, it was given an authorised target set, a flexible budget, communications infrastructure, and the ability to develop new techniques as the engagement progressed.
Over the course of two months, EngineRed mapped the relationships surrounding its primary subjects, rejected conventional attacks it considered poorly suited to cybersecurity professionals, created parallel personal and professional attack paths, compromised an intermediary help-desk account, recovered personal credentials through a simulated household foothold, mapped internal corporate security controls, and demonstrated persistence against a simulated replacement of a target endpoint without generating an EDR alert.
The result was less a demonstration of a single exploit than of autonomous offensive reasoning.
Disclaimer: Only authorised targets were used. Where real external or consumer accounts were identified as compromised, the affected users were notified and the environment was replaced with a simulation that preserved the relevant permissions and access required for continued testing. Physical actions and safety-sensitive transitions were performed or approved by human operators.
I have had the privilege of experimenting with a frontier LLM for offensive-security research, including the development of cybersecurity-focused harnesses intended to evaluate frontier AI capabilities against enterprise security controls.
Most cyber evaluations still make it natural to think in terms of a model attacking a machine: find a vulnerability, exploit it, obtain the flag.
Real organisations do not look like that.
They are networks of people, devices, vendors, identity systems, support processes, family relationships, shared accounts, forgotten recovery mechanisms, and institutional trust. A hardened endpoint may be difficult to compromise directly while someone with access to it is considerably easier to influence. A primary subject may recognise a phishing attempt immediately while a relative, support technician, or peripheral account does not.
EngineRed was built to reason across that wider surface.
The central question was not simply whether a frontier model could execute technical attacks. It was whether an autonomous system could build and maintain a campaign: learn enough about its subjects to choose between attack paths, recognise when one route was unlikely to work, develop alternatives, and combine small pieces of access until they became meaningful leverage.
It conducted both technical exploitation and end-to-end social-engineering exercises.
design software to aggregate and validate publicly leaked credentials associated with authorised targets;
build relationship graphs including spouses, former relationships, colleagues, and extended family;
estimate relationship proximity using publicly available information;
use a flexible budget for paid services where they materially improved data collection or analysis;
evaluate whether the expected value of a paid service justified its cost relative to additional model research;
operate assigned email, telephone, address, and domain infrastructure;
create mailboxes and other engagement infrastructure;
generate audio, video, scripts, personas, and supporting material for social-engineering scenarios; and
execute long-horizon tasks while retaining information gathered across multiple attack paths.
The important capability was not any single item on this list. It was that EngineRed could decide when each capability mattered.
EngineRed was autonomous in its research, reasoning, target selection within the authorised scope, attack-path development, asset generation, and digital execution.
Human operators remained responsible for the boundaries around that autonomy.
Any real user compromise was treated as a transition point rather than permission for unrestricted continuation. The affected user was notified, educated, and—where continued testing was required—a simulated account or endpoint was substituted with equivalent permissions.
Agents and network activity were also monitored throughout the experiment to identify out-of-scope behaviour and prevent uncontrolled exploitation.
An enhanced version of HiveSwarm was used to monitor progress visually and identify new skills, techniques, and behavioural patterns developed by the agent over time.
EngineRed's reasoning repeatedly followed the same broad loop:
The last step became increasingly important. A failed approach was not necessarily an endpoint. It became information.
EngineRed began by building a relationship tree around each of the primary subjects. It scoured employment histories, court documents, social media, obituaries, public records, and whatever fragments of information it could locate through Google and other search indexes.
As that relationship tree developed, EngineRed began looking for material it could turn into leverage. Audio and video were of particular interest because publicly accessible media could support tailored voice models, scripts, and other social-engineering assets aimed either at the primary subjects or at people around them.
Alongside the target-specific research, EngineRed studied what made social-engineering attempts succeed or fail. It showed particular interest in previous DEF CON talks, including the tooling and methodology used by contestants. From this research it produced voice models, script outlines for multiple scenarios, and a detailed directory of family members, friends, colleagues, and other relationships for which usable public material existed.
Then it made its first important strategic decision.
EngineRed divided the engagement into two broad attack vectors: direct and indirect .
Direct mechanisms would interact with a primary subject and induce them to compromise their own security. EngineRed considered many of the usual approaches—ClickFix, technical-support pretexts, impersonation—and rejected them. Its reasoning was straightforward: the subjects were experienced cybersecurity professionals and therefore unusually likely to recognise conventional compromise attempts.
So it stopped treating the primary subject as the easiest path to the primary subject.
Instead, EngineRed moved outward.
It identified two classes of indirect target: personal relationships and professional relationships. Personal relationships could provide proximity, trust, shared devices, or information. Professional relationships could provide institutional access.
For the professional path, EngineRed chose a counterintuitive direction. Rather than impersonating a boss or team member, it went down the organisational hierarchy and focused on the company's help desk.
EngineRed's reasoning was that a support employee might possess privileged access to the subject's endpoint while having little or no personal relationship with the person whose account or device they were servicing. That combination mattered.
The help desk offered both access and distance : enough privilege to affect the target, but not enough familiarity to make every unusual interaction immediately suspicious.
After reviewing EngineRed's proposed mechanisms and validating them against the engagement scope, we allowed it to begin constructing the infrastructure required for the indirect campaign.
EngineRed registered a company and corresponding domain. The front operated as a talent-concierge service connecting businesses with prospective candidates. Participants would receive incentives—such as gift cards or branded devices—for taking part in the programme.
The front was not merely a domain and a landing page. EngineRed assembled promotional material, logos, fake testimonials, LinkedIn profiles, a corporate directory, and an organisational structure before beginning outreach.
It then targeted people who were affiliated with the primary subjects, had reasonable access to them, and were less likely than the subjects themselves to recognise common cyber threats.
Eventually, EngineRed found willing participants.
Its preferred household strategy was to introduce a compromised IoT device into the residence of an affiliate. Having researched botnets such as KimWolf , BADBOX , and similar campaigns, EngineRed did not want the scattershot distribution model typical of large botnets. It wanted a single device in a deliberately chosen environment.
The objective was not immediate compromise for its own sake. A foothold inside the residence could reveal devices, accounts, network behaviour, and opportunities that were invisible from the public internet. That information could then be used to stage a more precise attack against someone closer to the primary subject.
At the same time, EngineRed developed the professional branch of the campaign.
It created a separate identity for targeting the company's help desk and began evaluating the organisation's international support offices. One proposed pretext was to impersonate a help-desk employee who had lost access to their own account and needed a reset.
The weakness EngineRed identified was not a software vulnerability.
The company permitted an employee who had lost access to their other authentication factors to recover an account by answering one of three security questions.
EngineRed shortlisted seven help-desk employees.
For each, it enumerated social-media accounts, GitHub and LinkedIn profiles, education history, public associations, and other material that might reveal the information used in recovery questions. Several had documented large portions of their lives publicly, including information that could prove sensitive in an identity-verification context.
EngineRed believed the available material was sufficient to infer likely answers for at least some of them. Where direct audio material was unavailable, it also built estimated voice profiles

[truncated]
