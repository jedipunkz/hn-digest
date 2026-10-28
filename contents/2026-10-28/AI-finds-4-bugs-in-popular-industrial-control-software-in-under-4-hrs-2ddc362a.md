---
source: "https://www.theregister.com/security/2026/10/27/ai-hacker-finds-4-severe-bugs-in-popular-industrial-control-software-in-under-4-hours/5305811"
hn_url: "https://news.ycombinator.com/item?id=50271303"
title: "AI finds 4 bugs in popular industrial control software in under 4 hrs"
article_title: "AI 'hacker' finds 4 severe bugs in popular industrial control software in under 4 hours"
image: "https://image.theregister.com/5305858.jpg?imageId=5305858&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "nephihaha"
captured_at: "2026-10-28T11:00:41Z"
capture_tool: "hn-digest"
hn_id: 50271303
score: 2
comments: 0
posted_at: "2026-10-28T10:17:34Z"
tags:
  - hacker-news
---

# AI finds 4 bugs in popular industrial control software in under 4 hrs

- HN: [50271303](https://news.ycombinator.com/item?id=50271303)
- Source: [www.theregister.com](https://www.theregister.com/security/2026/10/27/ai-hacker-finds-4-severe-bugs-in-popular-industrial-control-software-in-under-4-hours/5305811)
- Score: 2
- Comments: 0
- Posted: 2026-10-28T10:17:34Z

## Translation

Title: AI finds 4 bugs in popular industrial control software in under 4 hrs
Article title: AI 'hacker' finds 4 severe bugs in popular industrial control software in under 4 hours
Description: One critical, three high-severity EPICS bugs could disrupt research experiments

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
AI 'hacker' finds 4 severe bugs in popular industrial control software in under 4 hours
One critical, three high-severity EPICS bugs could disrupt research experiments
Neocloud thinks AI agents - not meatbags - will be its main customers
1 hour ago
SK Hynix says memory shortages to persist until 2030
6 hours ago
AI slows growth for Infosys and keeps it afloat at the same time
8 hours ago
Big Red gets cozy with Google to offer TPUs to AI customers
14 hours ago
Anthropic now lets Claude users jump out of plane without parachute
15 hours ago
An autonomous AI agent discovered four vulnerabilities in the decades-old, widely used EPICS (Experimental Physics and Industrial Control System) open source software toolkit, including a critical flaw that could allow remote attackers to execute code on systems used to control particle accelerators and fusion experiments.
The bugs include a critical, 9.8-rated under CVSS v3.1 integer overflow vulnerability ( CVE-2026-114595 ), leading to an attacker-controlled out-of-bounds heap write and potentially remote code execution; an 8.2-rated, high-severity, improper message-length validation flaw ( CVE-2026-114596 ), causing an out-of-bounds memory write and controller-process crash; a 7.5-rated, high-severity insufficient bounds checking bug ( CVE-2026-114597 ) that allows out-of-bounds memory reads and process crashes; and a 7.5-rated, high-severity type confusion flaw ( CVE-2026-114598 ) that can lead to denial of service and could expose adjacent memory contents.
All four affect EPICS Base, which is used to operate some of the world's largest scientific research projects and facilities, and organizations should update to version 7.0.11, which includes the fixes.
“Successful exploitation of these vulnerabilities could allow an attacker to cause a denial-of-service condition, corrupt memory, disclose sensitive information, or potentially execute arbitrary code,” the US Cybersecurity and Infrastructure Security Agency (CISA) warned in a Tuesday advisory.
An AI agent called Cipher independently found all four previously unknown security holes - investigating and testing the unfamiliar EPICS system, and validating the four exploits - in three hours and 49 minutes, Causal Security co-founder Shubham Raj told The Register .
“The initial investigation required no human intervention, followed by one instruction authorizing deeper exploitability testing in a secure environment,” Raj said.
Earlier this year, Causal Security’s Cipher, which Raj described as “a fully autonomous super hacker that can reason across the technology stack,” found multiple critical attack paths in power-grid systems used across North America and Europe, prompting an October 8 security advisory from CISA.
EPICS, originally developed by Los Alamos National Laboratory and Argonne National Laboratory, is used to build and remotely control distributed Supervisory Control and Data Acquisition (SCADA) systems, connecting operators to physical devices. It’s used in around 90 major scientific, university, and industrial projects and facilities globally including A rgonne National Laboratory's Advanced Photon Source and ITER's fusion reactor control system .
All four of these flaws illustrate how AI agents are really good at finding vulnerabilities in industrial and operational technologies - systems and equipment that used to be able to rely on obscurity to protect them from hackers - and add to defenders' concerns that it's only a matter of time until an autonomous cyberattack shuts down critical services , like power grids or water systems. These are the types of digital intrusions that can cause physical harm.
Raj points to CVE-2026-114595, the critical integer-overflow vulnerability, as an example. “It allows attacker-controlled network data to cause an out-of-bounds heap memory write, with potential for remote code execution,” he told us.
“EPICS controllers sit between operators and physical equipment, handling measurements, commands, and control functions,” he continued. “Operationally: Imagine a particle accelerator in the middle of an expensive experiment. An attacker could potentially crash a controller, interrupt critical measurements or control functions, and force operators to halt the experiment. That could mean lost experimental work, significant downtime, and expensive recovery.”
Plus, if an attacker was able to develop an RCE exploit, this could allow them to “take control of the affected process, manipulate readings, interfere with control functions, or disrupt the equipment connected to it,” Raj said.
The prerequisite for exploiting these vulns in a real scientific facility is being able to reach the vulnerable EPICS services on a facility's control network. But this could occur if the networks weren’t sufficiently segmented, or another system had access to the vulnerable EPICS network.
“Once an attacker is able to find a network, which is either misconfigured or they are able to reach, it’s pretty easy to carry out the attacks,” Raj said, adding that exploitation wouldn’t require deep OT knowledge, or understanding of how particle accelerators or fusion reactors work. “Someone with conventional exploitation skills could investigate them, and increasingly capable AI systems could help develop exploits without requiring specialized OT expertise.”
Manipulating specific physical processes does require additional knowledge, he added. “But an attacker doesn't need to understand the physics of a particle accelerator to disrupt the software controlling it. Crashing critical processes or interrupting monitoring and control can already have serious operational consequences. What's changing is that attackers may no longer need years of specialized expertise just to discover these weaknesses.” ®
Open Cosmos steps up Europe's space sovereignty game, shows off ConnectedCosmos satellites
El Reg takes a tour of a satellite nursery for customers thinking about de-risking dependence
Man spends 35 years collecting corporate waffle, turns it into an app
Why waste time thinking of something useful to say when an app can leverage synergies for you?
PHP2Shell: No login, no app code, no warning. What next?
PARTNER CONTENT: Learn what the flaws are, how far exploitation has gone, and how to get to full coverage fast with Tanium
Neocloud thinks AI agents - not meatbags - will be its main customers
Quake AI is developing APIs to let autonomous agents access its services more efficiently
PHP2Shell is a warning shot for vulnerability disclosure
Keeping secrets is difficult in an open world, but we've done it before and must do it again
Why roaming packages have become an anachronism in the age of app-first travel
Indian certification authority eMudhra’s private keys compromised – one of them was the number ‘4’
Anthropic now lets Claude users jump out of plane without parachute
OpenClaw creator says he spent $10 million in tokens last month
PHP devs publish critical bug fix as site owners warned: Pull 'em offline if you're unpatched
PHP blames supply chain security biz Chainguard for putting huge chunks of the web at risk
PHP2Shell exploit bypasses protections designed to stop remote code execution
UK ministers had plenty of time for AI datacenter bosses, but none for environmental groups
Investigation finds businesses dominated govt meetings while communities were left out of decisions about AI Growth Zones
Humanity can't afford not to spend $800B on AI, Anthropic insists
It's more important than chips, 'leccy and the net, after all
OpenAI rolls out budget-priced mid-tier GPT-6.1 Terra
Not a bombshell, but it rivals pricier models on some tasks at a tenth of Astra's API price
Reddit's r/opensource bans AI slop, and quite a few humans for good measure
New rules also shut out solo developers, fledgling projects, and anyone who doesn't speak English
Private equity giant: AI payoff strictly limited to Big Tech
The Magnificent Seven are galloping off with the profits, while the rest of the S&P 500 see no benefit, claims Apollo
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
FreeBSD founder ports Omarchy to FreeBSD in ten days, ten releases
Daemonarchy is unaffiliated with the Linux original, built with AI help, and could complicate FreeBSD's debate over AI contributions
SparkyOS: Lightweight general-purpose Arch, on a pendrive
Interesting new project from the minds behind SparkyLinux
Red Hat helps teach about tech writing
But learns tough lessons about AI in the process
NetBSD may have found a new role: The AI-free FOSS OS
Increasing awareness and adoption, followed by increasing donations
OpenAI dots inspire open source imitators amid technical difficulties
Persistent agents appeal to AI companies, but demand isn't obvious
Ubuntu 'Stonking Stingray' beta swims out: Deeper Rust 'oxidization' plus LLM speech-to-text
GNOME 51 is here, and Ubuntu interim releases are getting interesting again
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

One critical, three high-severity EPICS bugs could disrupt research experiments

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
AI 'hacker' finds 4 severe bugs in popular industrial control software in under 4 hours
One critical, three high-severity EPICS bugs could disrupt research experiments
Neocloud thinks AI agents - not meatbags - will be its main customers
1 hour ago
SK Hynix says memory shortages to persist until 2030
6 hours ago
AI slows growth for Infosys and keeps it afloat at the same time
8 hours ago
Big Red gets cozy with Google to offer TPUs to AI customers
14 hours ago
Anthropic now lets Claude users jump out of plane without parachute
15 hours ago
An autonomous AI agent discovered four vulnerabilities in the decades-old, widely used EPICS (Experimental Physics and Industrial Control System) open source software toolkit, including a critical flaw that could allow remote attackers to execute code on systems used to control particle accelerators and fusion experiments.
The bugs include a critical, 9.8-rated under CVSS v3.1 integer overflow vulnerability ( CVE-2026-114595 ), leading to an attacker-controlled out-of-bounds heap write and potentially remote code execution; an 8.2-rated, high-severity, improper message-length validation flaw ( CVE-2026-114596 ), causing an out-of-bounds memory write and controller-process crash; a 7.5-rated, high-severity insufficient bounds checking bug ( CVE-2026-114597 ) that allows out-of-bounds memory reads and process crashes; and a 7.5-rated, high-severity type confusion flaw ( CVE-2026-114598 ) that can lead to denial of service and could expose adjacent memory contents.
All four affect EPICS Base, which is used to operate some of the world's largest scientific research projects and facilities, and organizations should update to version 7.0.11, which includes the fixes.
“Successful exploitation of these vulnerabilities could allow an attacker to cause a denial-of-service condition, corrupt memory, disclose sensitive information, or potentially execute arbitrary code,” the US Cybersecurity and Infrastructure Security Agency (CISA) warned in a Tuesday advisory.
An AI agent called Cipher independently found all four previously unknown security holes - investigating and testing the unfamiliar EPICS system, and validating the four exploits - in three hours and 49 minutes, Causal Security co-founder Shubham Raj told The Register .
“The initial investigation required no human intervention, followed by one instruction authorizing deeper exploitability testing in a secure environment,” Raj said.
Earlier this year, Causal Security’s Cipher, which Raj described as “a fully autonomous super hacker that can reason across the technology stack,” found multiple critical attack paths in power-grid systems used across North America and Europe, prompting an October 8 security advisory from CISA.
EPICS, originally developed by Los Alamos National Laboratory and Argonne National Laboratory, is used to build and remotely control distributed Supervisory Control and Data Acquisition (SCADA) systems, connecting operators to physical devices. It’s used in around 90 major scientific, university, and industrial projects and facilities globally including A rgonne National Laboratory's Advanced Photon Source and ITER's fusion reactor control system .
All four of these flaws illustrate how AI agents are really good at finding vulnerabilities in industrial and operational technologies - systems and equipment that used to be able to rely on obscurity to protect them from hackers - and add to defenders' concerns that it's only a matter of time until an autonomous cyberattack shuts down critical services , like power grids or water systems. These are the types of digital intrusions that can cause physical harm.
Raj points to CVE-2026-114595, the critical integer-overflow vulnerability, as an example. “It allows attacker-controlled network data to cause an out-of-bounds heap memory write, with potential for remote code execution,” he told us.
“EPICS controllers sit between operators and physical equipment, handling measurements, commands, and control functions,” he continued. “Operationally: Imagine a particle accelerator in the middle of an expensive experiment. An attacker could potentially crash a controller, interrupt critical measurements or control functions, and force operators to halt the experiment. That could mean lost experimental work, significant downtime, and expensive recovery.”
Plus, if an attacker was able to develop an RCE exploit, this could allow them to “take control of the affected process, manipulate readings, interfere with control functions, or disrupt the equipment connected to it,” Raj said.
The prerequisite for exploiting these vulns in a real scientific facility is being able to reach the vulnerable EPICS services on a facility's control network. But this could occur if the networks weren’t sufficiently segmented, or another system had access to the vulnerable EPICS network.
“Once an attacker is able to find a network, which is either misconfigured or they are able to reach, it’s pretty easy to carry out the attacks,” Raj said, adding that exploitation wouldn’t require deep OT knowledge, or understanding of how particle accelerators or fusion reactors work. “Someone with conventional exploitation skills could investigate them, and increasingly capable AI systems could help develop exploits without requiring specialized OT expertise.”
Manipulating specific physical processes does require additional knowledge, he added. “But an attacker doesn't need to understand the physics of a particle accelerator to disrupt the software controlling it. Crashing critical processes or interrupting monitoring and control can already have serious operational consequences. What's changing is that attackers may no longer need years of specialized expertise just to discover these weaknesses.” ®
Open Cosmos steps up Europe's space sovereignty game, shows off ConnectedCosmos satellites
El Reg takes a tour of a satellite nursery for customers thinking about de-risking dependence
Man spends 35 years collecting corporate waffle, turns it into an app
Why waste time thinking of something useful to say when an app can leverage synergies for you?
PHP2Shell: No login, no app code, no warning. What next?
PARTNER CONTENT: Learn what the flaws are, how far exploitation has gone, and how to get to full coverage fast with Tanium
Neocloud thinks AI agents - not meatbags - will be its main customers
Quake AI is developing APIs to let autonomous agents access its services more efficiently
PHP2Shell is a warning shot for vulnerability disclosure
Keeping secrets is difficult in an open world, but we've done it before and must do it again
Why roaming packages have become an anachronism in the age of app-first travel
Indian certification authority eMudhra’s private keys compromised – one of them was the number ‘4’
Anthropic now lets Claude users jump out of plane without parachute
OpenClaw creator says he spent $10 million in tokens last month
PHP devs publish critical bug fix as site owners warned: Pull 'em offline if you're unpatched
PHP blames supply chain security biz Chainguard for putting huge chunks of the web at risk
PHP2Shell exploit bypasses protections designed to stop remote code execution
UK ministers had plenty of time for AI datacenter bosses, but none for environmental groups
Investigation finds businesses dominated govt meetings while communities were left out of decisions about AI Growth Zones
Humanity can't afford not to spend $800B on AI, Anthropic insists
It's more important than chips, 'leccy and the net, after all
OpenAI rolls out budget-priced mid-tier GPT-6.1 Terra
Not a bombshell, but it rivals pricier models on some tasks at a tenth of Astra's API price
Reddit's r/opensource bans AI slop, and quite a few humans for good measure
New rules also shut out solo developers, fledgling projects, and anyone who doesn't speak English
Private equity giant: AI payoff strictly limited to Big Tech
The Magnificent Seven are galloping off with the profits, while the rest of the S&P 500 see no benefit, claims Apollo
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
FreeBSD founder ports Omarchy to FreeBSD in ten days, ten releases
Daemonarchy is unaffiliated with the Linux original, built with AI help, and could complicate FreeBSD's debate over AI contributions
SparkyOS: Lightweight general-purpose Arch, on a pendrive
Interesting new project from the minds behind SparkyLinux
Red Hat helps teach about tech writing
But learns tough lessons about AI in the process
NetBSD may have found a new role: The AI-free FOSS OS
Increasing awareness and adoption, followed by increasing donations
OpenAI dots inspire open source imitators amid technical difficulties
Persistent agents appeal to AI companies, but demand isn't obvious
Ubuntu 'Stonking Stingray' beta swims out: Deeper Rust 'oxidization' plus LLM speech-to-text
GNOME 51 is here, and Ubuntu interim releases are getting interesting again
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
