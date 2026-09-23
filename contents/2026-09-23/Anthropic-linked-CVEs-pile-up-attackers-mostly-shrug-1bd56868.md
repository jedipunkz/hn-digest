---
source: "https://www.theregister.com/security/2026/09/21/anthropic-linked-cves-pile-up-attackers-mostly-shrug/5298018"
hn_url: "https://news.ycombinator.com/item?id=49814140"
title: "Anthropic-linked CVEs pile up, attackers mostly shrug"
article_title: "Anthropic-linked CVEs pile up, attackers mostly shrug"
image: "https://image.theregister.com/5298025.jpg?imageId=5298025&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "vismit2000"
captured_at: "2026-09-23T10:53:41Z"
capture_tool: "hn-digest"
hn_id: 49814140
score: 1
comments: 0
posted_at: "2026-09-23T10:52:11Z"
tags:
  - hacker-news
---

# Anthropic-linked CVEs pile up, attackers mostly shrug

- HN: [49814140](https://news.ycombinator.com/item?id=49814140)
- Source: [www.theregister.com](https://www.theregister.com/security/2026/09/21/anthropic-linked-cves-pile-up-attackers-mostly-shrug/5298018)
- Score: 1
- Comments: 0
- Posted: 2026-09-23T10:52:11Z

## Translation

Title: Anthropic-linked CVEs pile up, attackers mostly shrug
Description: Of 225 flaws found by Glasswing and tracked by VulnCheck researcher, just one has confirmed exploitation in the wild

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
Anthropic-linked CVEs pile up, attackers mostly shrug
Of 225 flaws found by Glasswing and tracked by VulnCheck researcher, just one has confirmed exploitation in the wild
Frontier AI keeps racing despite calls to slow down
10 hours ago
Treasury chief says AI bosses, not their bots, will carry the can for criminal acts
1 day ago
Agentic security is the billion-dollar challenge for some clever startup to solve
3 days ago
Anthropic decides to support OpenAI's markdown instructions spec
4 days ago
Claude Code revamps projects so you can work and pay in parallel
4 days ago
Despite the concern that advanced AI models’ bug-hunting prowess will lead to attackers exploiting more newly uncovered CVEs, fewer than 0.5 percent of the vulnerabilities linked to Anthropic or Project Glasswing are being battered in the wild, according to VulnCheck security researcher Patrick Garrity.
Garrity began tracking CVEs attributed to Project Glasswing, Anthropic’s initiative to give select partners access to its Claude Mythos Preview model, shortly after the AI company announced the program in April.
At the time, Anthropic said the new model was too risky to release publicly because its bug-finding and exploitation skills surpass all but the most skilled humans. As such, Anthropic restricted access to Mythos Preview to vetted Glasswing participants, who use the model for defensive security work, including finding and fixing flaws in their own software products and open source dependencies.
Garrity’s Anthropic CVE tracker maintains a list of vulnerabilities credited to the Anthropic team and/or Project Glasswing and also checks these CVEs against the company's known exploited vulnerabilities index "to get a better read on the real Glasswing ‘danger factor . ’"
As of Monday, the CVE count is 225, and just one, a critical SQL injection bug in Ghost ( CVE-2026-26980 ), has been exploited in the wild.
“There's a big difference between finding vulnerabilities and whether they're actually useful to and will be used by threat actors,” Garrity told The Register . “The main thing this data highlights is that what Anthropic is discovering and disclosing is fairly limited in impact, and from what we can tell, isn't resulting in different outcomes from a threat perspective than a random selection of other vulnerabilities would.”
Anthropic didn’t immediately respond to our questions, but we will update this story if we hear back.
Garrity says he doesn’t dispute AI’s ability to find bugs.
A lot of the hysteria we're seeing assumes that every vulnerability or bug is likely to be used by threat actors. But the reality is that only a small fraction ever get used in exploitation campaigns
Indeed, anyone following security disclosures over the past few months would have a hard time arguing that AI models aren’t bringing to light significantly more security flaws than ever before. Case in point: recent massive patch drops from Microsoft , Apple , Palo Alto Networks , and don’t even get us started on open source projects .
Also, as Garrity pointed out, these vulnerability-finding skills aren’t “a capability unique to one model or harness.”
“A lot of the hysteria we're seeing assumes that every vulnerability or bug is likely to be used by threat actors,” he told The Register . “But the reality is that only a small fraction ever get used in exploitation campaigns. Historically, that's ranged from just under one percent to two percent of vulnerabilities that get weaponized and used in the wild.”
Plus, while recent AI models excel at finding bugs, they still aren’t great at fixing them, as a couple of recent studies have highlighted.
In one of these, 1Password’s research team produced and analyzed 6,080 patches developed by two frontier models: OpenAI's ChatGPT-5.5 and Anthropic's Opus 4.8. The models generated fixes that fully resolved the vulnerability just 26 percent of the time, while about 54 percent either failed to resolve the vulnerability, introduced a new vulnerability, or did both.
Another study by app security shop Veracode found that across more than 100 models and 80 coding tasks, the average security pass rate for AI-generated code was just 56 percent.
This all means that the work involved in developing and applying security fixes still requires humans.
“The bar for vulnerability discovery is much lower with AI, but the real gap lies downstream in coordination, triage, remediation, and patch deployment, which is still largely people-intensive work, as Anthropic itself has acknowledged,” Garrity said. “It appears they might not have realized this until after they launched the project.”®
Swedish celebs campaign for public rudeness ... to prevent cyber scams
The campaign follows a highly profitable year for crooks targeting the over-60s
BigCommerce app breach spills Master of Malt customer data
Attackers had four days to drink in names, addresses, emails and phone numbers
HPE makes its “unified storage” claim real as B10000 R6 hits GA
PARTNER CONTENT: Pairs block and adjacent file workloads with independent scaling of performance and capacity
BT Tower's rooftop pool plan brings swimmers back down to earth
MCR's hotel proposal puts facility on a lower building, while public access to upper floors set to be restored
systems approach
In the age of AI, teaching networking principles remains more important than learning protocols
Kids can learn why BGP matters in a semester, but that won’t leave them ready to implement it
UK’s uncrewed experimental sub shows it can fire torpedoes
XV Excalibur hailed as first US or UK underwater drone to let one off in tests
SAAS
Salesforce staggers back to feet after global outage
Anthropic decides to support OpenAI's markdown instructions spec
Microsoft agentically ports Copilot runtime to Rust for $120K
software
Fedora 45 beta drags the Linux console into the 21st century
KPMG tech cuts come with a severance sum some staff call insulting
on call
Techie fixed Wi-Fi dead zone with a drill
Shut up and calculate: Jev's new AI primitives for coders
Developers test what they can build with TypeSafe's fast, typed decision model
Frontier AI keeps racing despite calls to slow down
Anthropic and OpenAI debut Opus 5.5 and GPT-6 Sol and Luna
Windows CLOSEDQUORUM malware uses AI models to autonomously select post-compromise actions
'first' publicly documented Windows implant to use LLMs for C2
Z.ai says sorry for slurping up your code, open sources ZCode
China’s AI darling goes on the defense after engineer highlighted Grok-esque security flaws
Did Copilot write the synchronization code?
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
KDE turns 30 and someone's brought an AI-native desktop proposal
Akademy talk imagines Plasma assembling itself around a personal model of each user
Shopify extends lifeline to Tailwind as vibe coding erodes web dev platform's bottom line
Acquisition gives open source CSS framework 'a stable long-term home'
Switzerland tests a FOSS escape route from Microsoft 365
Swiss Army sticks a knife in American cloud apps with its own FOSS push
Feel peak Windows was 7? You might like Kumander Linux
Debian and Xfce – solid, sensible choices – with a pretty skin
Canonical shuttering some of its legacy chat channels
The Ubuntu Pastebin went in June, IRC gets demoted next
Audacity audio-editing app no longer looks like it's from the early 2000s
The FOSS tool for audio editing has a fresh coat of paint, and new features to boot
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

Of 225 flaws found by Glasswing and tracked by VulnCheck researcher, just one has confirmed exploitation in the wild

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Anthropic-linked CVEs pile up, attackers mostly shrug
Of 225 flaws found by Glasswing and tracked by VulnCheck researcher, just one has confirmed exploitation in the wild
Frontier AI keeps racing despite calls to slow down
10 hours ago
Treasury chief says AI bosses, not their bots, will carry the can for criminal acts
1 day ago
Agentic security is the billion-dollar challenge for some clever startup to solve
3 days ago
Anthropic decides to support OpenAI's markdown instructions spec
4 days ago
Claude Code revamps projects so you can work and pay in parallel
4 days ago
Despite the concern that advanced AI models’ bug-hunting prowess will lead to attackers exploiting more newly uncovered CVEs, fewer than 0.5 percent of the vulnerabilities linked to Anthropic or Project Glasswing are being battered in the wild, according to VulnCheck security researcher Patrick Garrity.
Garrity began tracking CVEs attributed to Project Glasswing, Anthropic’s initiative to give select partners access to its Claude Mythos Preview model, shortly after the AI company announced the program in April.
At the time, Anthropic said the new model was too risky to release publicly because its bug-finding and exploitation skills surpass all but the most skilled humans. As such, Anthropic restricted access to Mythos Preview to vetted Glasswing participants, who use the model for defensive security work, including finding and fixing flaws in their own software products and open source dependencies.
Garrity’s Anthropic CVE tracker maintains a list of vulnerabilities credited to the Anthropic team and/or Project Glasswing and also checks these CVEs against the company's known exploited vulnerabilities index "to get a better read on the real Glasswing ‘danger factor . ’"
As of Monday, the CVE count is 225, and just one, a critical SQL injection bug in Ghost ( CVE-2026-26980 ), has been exploited in the wild.
“There's a big difference between finding vulnerabilities and whether they're actually useful to and will be used by threat actors,” Garrity told The Register . “The main thing this data highlights is that what Anthropic is discovering and disclosing is fairly limited in impact, and from what we can tell, isn't resulting in different outcomes from a threat perspective than a random selection of other vulnerabilities would.”
Anthropic didn’t immediately respond to our questions, but we will update this story if we hear back.
Garrity says he doesn’t dispute AI’s ability to find bugs.
A lot of the hysteria we're seeing assumes that every vulnerability or bug is likely to be used by threat actors. But the reality is that only a small fraction ever get used in exploitation campaigns
Indeed, anyone following security disclosures over the past few months would have a hard time arguing that AI models aren’t bringing to light significantly more security flaws than ever before. Case in point: recent massive patch drops from Microsoft , Apple , Palo Alto Networks , and don’t even get us started on open source projects .
Also, as Garrity pointed out, these vulnerability-finding skills aren’t “a capability unique to one model or harness.”
“A lot of the hysteria we're seeing assumes that every vulnerability or bug is likely to be used by threat actors,” he told The Register . “But the reality is that only a small fraction ever get used in exploitation campaigns. Historically, that's ranged from just under one percent to two percent of vulnerabilities that get weaponized and used in the wild.”
Plus, while recent AI models excel at finding bugs, they still aren’t great at fixing them, as a couple of recent studies have highlighted.
In one of these, 1Password’s research team produced and analyzed 6,080 patches developed by two frontier models: OpenAI's ChatGPT-5.5 and Anthropic's Opus 4.8. The models generated fixes that fully resolved the vulnerability just 26 percent of the time, while about 54 percent either failed to resolve the vulnerability, introduced a new vulnerability, or did both.
Another study by app security shop Veracode found that across more than 100 models and 80 coding tasks, the average security pass rate for AI-generated code was just 56 percent.
This all means that the work involved in developing and applying security fixes still requires humans.
“The bar for vulnerability discovery is much lower with AI, but the real gap lies downstream in coordination, triage, remediation, and patch deployment, which is still largely people-intensive work, as Anthropic itself has acknowledged,” Garrity said. “It appears they might not have realized this until after they launched the project.”®
Swedish celebs campaign for public rudeness ... to prevent cyber scams
The campaign follows a highly profitable year for crooks targeting the over-60s
BigCommerce app breach spills Master of Malt customer data
Attackers had four days to drink in names, addresses, emails and phone numbers
HPE makes its “unified storage” claim real as B10000 R6 hits GA
PARTNER CONTENT: Pairs block and adjacent file workloads with independent scaling of performance and capacity
BT Tower's rooftop pool plan brings swimmers back down to earth
MCR's hotel proposal puts facility on a lower building, while public access to upper floors set to be restored
systems approach
In the age of AI, teaching networking principles remains more important than learning protocols
Kids can learn why BGP matters in a semester, but that won’t leave them ready to implement it
UK’s uncrewed experimental sub shows it can fire torpedoes
XV Excalibur hailed as first US or UK underwater drone to let one off in tests
SAAS
Salesforce staggers back to feet after global outage
Anthropic decides to support OpenAI's markdown instructions spec
Microsoft agentically ports Copilot runtime to Rust for $120K
software
Fedora 45 beta drags the Linux console into the 21st century
KPMG tech cuts come with a severance sum some staff call insulting
on call
Techie fixed Wi-Fi dead zone with a drill
Shut up and calculate: Jev's new AI primitives for coders
Developers test what they can build with TypeSafe's fast, typed decision model
Frontier AI keeps racing despite calls to slow down
Anthropic and OpenAI debut Opus 5.5 and GPT-6 Sol and Luna
Windows CLOSEDQUORUM malware uses AI models to autonomously select post-compromise actions
'first' publicly documented Windows implant to use LLMs for C2
Z.ai says sorry for slurping up your code, open sources ZCode
China’s AI darling goes on the defense after engineer highlighted Grok-esque security flaws
Did Copilot write the synchronization code?
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
KDE turns 30 and someone's brought an AI-native desktop proposal
Akademy talk imagines Plasma assembling itself around a personal model of each user
Shopify extends lifeline to Tailwind as vibe coding erodes web dev platform's bottom line
Acquisition gives open source CSS framework 'a stable long-term home'
Switzerland tests a FOSS escape route from Microsoft 365
Swiss Army sticks a knife in American cloud apps with its own FOSS push
Feel peak Windows was 7? You might like Kumander Linux
Debian and Xfce – solid, sensible choices – with a pretty skin
Canonical shuttering some of its legacy chat channels
The Ubuntu Pastebin went in June, IRC gets demoted next
Audacity audio-editing app no longer looks like it's from the early 2000s
The FOSS tool for audio editing has a fresh coat of paint, and new features to boot
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
