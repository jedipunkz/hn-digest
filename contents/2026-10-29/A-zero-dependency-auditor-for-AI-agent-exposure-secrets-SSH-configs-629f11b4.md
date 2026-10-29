---
source: "https://ai-agent-antivirus.space/"
hn_url: "https://news.ycombinator.com/item?id=50283625"
title: "A zero-dependency auditor for AI agent exposure (secrets, SSH, configs)"
article_title: "AI Agent Antivirus: Protect Your Machine From AI Agents | Dante"
image: ""
author: "JohnMci"
captured_at: "2026-10-29T02:09:11Z"
capture_tool: "hn-digest"
hn_id: 50283625
score: 4
comments: 1
posted_at: "2026-10-29T01:56:16Z"
tags:
  - hacker-news
---

# A zero-dependency auditor for AI agent exposure (secrets, SSH, configs)

- HN: [50283625](https://news.ycombinator.com/item?id=50283625)
- Source: [ai-agent-antivirus.space](https://ai-agent-antivirus.space/)
- Score: 4
- Comments: 1
- Posted: 2026-10-29T01:56:16Z

## Translation

Title: A zero-dependency auditor for AI agent exposure (secrets, SSH, configs)
Article title: AI Agent Antivirus: Protect Your Machine From AI Agents | Dante
Description: A classic antivirus looks for malware. An AI agent needs none: it uses your own permissions. Dante checks the places agents and thieves reach for - keys, SSH, MCP configs, servers - and shows the proof line for every verdict.

Article text:
NEWS Dante is a free watchman for your own machines. See what is new in the library. Library →
Dante Home Agent risks Leaked keys Server check Bait files Install & commands Install & packs Hugging Face
Antivirus for AI agents: what actually protects your machine
An AI coding agent runs with your rights. It can read your .env , your SSH keys and your shell history, and run commands. Nothing "malicious" has to be installed for that to go wrong. This page explains why signature antivirus does not cover that, what does, and how to check your own machine in three minutes.
Runs on your machine, offline. The watchman needs no internet and no account to check. Once it is installed, turn Wi-Fi off and run it: it still works. He makes no network call on his own . Only three things can reach out (fetch a knowledge pack, look at the shop catalog, the optional bait watch), and each one only after you say yes to the exact command he shows first.
Nothing hidden in the package. No dependencies, no install scripts. Nothing runs when you install it.
Open and readable. MIT-licensed, plain JavaScript. You can download exactly what npm serves and read the whole program.
The only way out is curl, and you can count it. Dante has no web code of his own (no HTTP library); to reach the internet he starts curl , and the check below lists every place where that can happen. The only direct connection he opens himself is the question "does port 22 answer?" on a server you named.
Extendable from outside, safely. New knowledge arrives as signed data packs (rules, never code), checked against a key built into the program. A downloaded pack installs with pack <file> , without any network.
The one-minute check (terminal, Node.js installed). Copy each line, press Enter:
YOU YOUR AI AGENT ANYONE WHO TALKS TO IT
┌─────┐ same rights ┌────────────────┐ instructions ┌────────────┐
│ ◉ ◉ │ ───────────────▶ │ reads .env │ ◀─────────────── │ web page │
│ ▽ │ │ reads ~/.ssh │ │ issue │
└─────┘ │ runs commands│ │ README │
└───────┬────────┘ └────────────┘
│
what it can reach is
what a hijack can take The risk is not a virus. It is the reach.
TL;DR
What: Dante is a free watchman for your own computer and servers. It finds the things an AI agent or a thief would grab first: loose keys, open doors, risky settings.
Why not antivirus: an agent is legitimate software with your rights. There is no virus to catch. There is reach to limit.
How: npx dante-watch , then say "check my computer". It only reads, and every verdict shows its proof line.
Where: Mac, Windows, Linux, and servers you reach by SSH. Nothing leaves your machine unless you say yes to a command he shows first.
Bait watch: Dante lays honeypots, fake keys that look exactly like what a thief or a rogue agent grabs first. If someone uses one, it calls us and you get a mail within seconds. First month free, then 1.99 EUR a month; that money pays for the research behind it. We are building this out.
Cost: the watchman itself is free. Optional knowledge packs and the bait watch if you want more.
Why a normal antivirus is the wrong tool for this
A classic antivirus compares files and behaviour against known malware. That works when something foreign runs on the machine. With an AI agent the situation is different: the program is legitimate, it is signed, it was installed on purpose, and it acts with your user rights. The risk is not a virus. The risk is what the agent can reach and what can talk it into using it.
That is the reasoning behind the product, not a measured detection statistic: if the dangerous thing is a secret lying in a readable file, the useful check is "which secrets lie where, and who can read them", not "does this binary match a signature".
This is not theory. In August 2025 the s1ngularity attack on the Nx build tool shipped malware that looked for installed AI command-line assistants and told them to search the victim's file system for credentials, SSH keys and wallets. Weeks later the self-spreading Shai-Hulud npm worm hunted for publishing tokens. Both worked because the secrets were sitting where any process of the user could read them.
.env db + API keys
~/.ssh/id_ed25519
~/.aws/credentials
shell history
agent ❯ cat .env agent ❯ ls ~/.ssh agent ❯ history agent ❯ ▌
grrr… let me in
same rights as you · nothing has to be installed for this to go wrong
Your agent wants in. It already has your rights, so what it can reach is what matters.
What Dante is
Dante is a free, open command-line watchman for machines you own: your Mac, Linux or Windows computer, and servers you can reach by SSH. You talk to it in plain German, English or Spanish. It reads, it explains, and it changes nothing without a clear yes.
Every verdict carries the real line it came from. If Dante cannot read the proof, the verdict is unknown . It never says safe on an assumption.
Silent by default. No telemetry, no background update check, no account. The free state runs fully offline.
Zero dependencies. One small Node.js program you can read. The source is public .
user ❯ check my computer
SAFE Openly readable secrets in the web directory
what: No openly readable .env/.sql/.git files in the web root.
proof:
• Web root searched, nothing readable found
↳ dante:find-files /var/www (.env, *.sql, …) → (no matches in: /var/www)
UNKNOWN SSH: root login and password login
what: No SSH service listens on port 22 here; I do not know whether the settings matter.
proof:
• PermitRootLogin = (not set)
↳ /etc/ssh/sshd_config → PermitRootLogin default Every verdict carries the line it came from. If Dante cannot read the proof, it says UNKNOWN, never SAFE.
What it checks
Secrets lying around : API keys, tokens and private keys in project files, .env files and Git history, matched against a library of 1,111 secret formats. It shows file, line and kind - never the value.
Your servers : SSH settings, world-writable files in web folders, known PHP vulnerabilities by version, signs of web shells.
AI agent configuration : MCP server entries and agent settings (with the optional Hugging Face analysis pack), reported without ever printing a key.
Bait files (the "Wachschutz"): fake credentials that mail you the moment someone uses them.
Windows: services anyone can overwrite, unquoted service paths, writable PATH folders, AlwaysInstallElevated, stored autologon passwords.
Is it any good? An honest answer
Dante is young. It is a watchman, not a real-time engine: it looks when you ask or on the schedule you agreed to, and it does not scan running processes for malware. What it does well is the boring, high-value part: finding the secrets and settings an attacker or a hijacked agent would use first, and refusing to claim more than it can prove. Its language layer (what you type) is trained on about 259,000 labelled sentences and answers "I don't know" instead of guessing when unsure.
STEP 1 LOOK ████████████████████ free · 3 minutes · reads only
STEP 2 DIG DEEP ██████████████░░░░░░ secrets in your project folders
STEP 3 BAIT █████████░░░░░░░░░░░ a fake key that rings when it is used
STEP 4 STAY SHARP ████░░░░░░░░░░░░░░░░ fresh knowledge packs for new attacks Most people stop after step 1 and are already better off. Each step is its own yes.
You never have to decide everything now. Start with step 1. Dante asks before every next thing, and every question shows the exact command first.
Install Node.js (nodejs.org, once), open a terminal, paste one line:
you@MacBook ~ % npx dante-watch
I am Dante. I walk through hell and bring back the map.
● awake · I watch over localhost.
user ❯ check my computer Mac: Terminal
PowerShell — Windows PS C:\Users\you> npx dante-watch
I am Dante. I walk through hell and bring back the map.
● awake · I watch over localhost.
user ❯ check my computer Windows: PowerShell
Same program, same words, on every system. Not sure where your terminal is? The install page shows it for Mac, Windows and Linux, step by step.
Details for every system, the free and paid parts and what each costs are on the install page .
Look, nobody wakes up wanting to audit their laptop. It is boring, it feels like homework, and it always looks fine until the day it does not.
So here is the deal. It takes three minutes. It only reads. It costs nothing. If it finds nothing, you wasted three minutes and you can say you checked. If it finds a key lying around, you found it before somebody else did.
Your agent has the keys to your house either way. The only question is whether you know which keys.
They need protection, but not the kind that matches malware signatures. An agent is legitimate software running with your rights. What matters is which secrets and servers it can reach and whether anything can instruct it. Checking where secrets lie, how SSH is set and what the agent configuration allows covers the real exposure.
No. Dante is a watchman for machines you own. It does not scan for malware in real time. It finds exposed secrets, risky settings and known vulnerabilities, and shows the proof line for each result.
No. By default it sends nothing. Only three actions use the network (update, the catalog, and the optional bait watch) and each starts only after you answer yes to the command Dante prints first.
Yes. The free state is complete and works offline. Extra knowledge packs and the bait watch are optional; prices are on the install page.

## Original Extract

A classic antivirus looks for malware. An AI agent needs none: it uses your own permissions. Dante checks the places agents and thieves reach for - keys, SSH, MCP configs, servers - and shows the proof line for every verdict.

NEWS Dante is a free watchman for your own machines. See what is new in the library. Library →
Dante Home Agent risks Leaked keys Server check Bait files Install & commands Install & packs Hugging Face
Antivirus for AI agents: what actually protects your machine
An AI coding agent runs with your rights. It can read your .env , your SSH keys and your shell history, and run commands. Nothing "malicious" has to be installed for that to go wrong. This page explains why signature antivirus does not cover that, what does, and how to check your own machine in three minutes.
Runs on your machine, offline. The watchman needs no internet and no account to check. Once it is installed, turn Wi-Fi off and run it: it still works. He makes no network call on his own . Only three things can reach out (fetch a knowledge pack, look at the shop catalog, the optional bait watch), and each one only after you say yes to the exact command he shows first.
Nothing hidden in the package. No dependencies, no install scripts. Nothing runs when you install it.
Open and readable. MIT-licensed, plain JavaScript. You can download exactly what npm serves and read the whole program.
The only way out is curl, and you can count it. Dante has no web code of his own (no HTTP library); to reach the internet he starts curl , and the check below lists every place where that can happen. The only direct connection he opens himself is the question "does port 22 answer?" on a server you named.
Extendable from outside, safely. New knowledge arrives as signed data packs (rules, never code), checked against a key built into the program. A downloaded pack installs with pack <file> , without any network.
The one-minute check (terminal, Node.js installed). Copy each line, press Enter:
YOU YOUR AI AGENT ANYONE WHO TALKS TO IT
┌─────┐ same rights ┌────────────────┐ instructions ┌────────────┐
│ ◉ ◉ │ ───────────────▶ │ reads .env │ ◀─────────────── │ web page │
│ ▽ │ │ reads ~/.ssh │ │ issue │
└─────┘ │ runs commands│ │ README │
└───────┬────────┘ └────────────┘
│
what it can reach is
what a hijack can take The risk is not a virus. It is the reach.
TL;DR
What: Dante is a free watchman for your own computer and servers. It finds the things an AI agent or a thief would grab first: loose keys, open doors, risky settings.
Why not antivirus: an agent is legitimate software with your rights. There is no virus to catch. There is reach to limit.
How: npx dante-watch , then say "check my computer". It only reads, and every verdict shows its proof line.
Where: Mac, Windows, Linux, and servers you reach by SSH. Nothing leaves your machine unless you say yes to a command he shows first.
Bait watch: Dante lays honeypots, fake keys that look exactly like what a thief or a rogue agent grabs first. If someone uses one, it calls us and you get a mail within seconds. First month free, then 1.99 EUR a month; that money pays for the research behind it. We are building this out.
Cost: the watchman itself is free. Optional knowledge packs and the bait watch if you want more.
Why a normal antivirus is the wrong tool for this
A classic antivirus compares files and behaviour against known malware. That works when something foreign runs on the machine. With an AI agent the situation is different: the program is legitimate, it is signed, it was installed on purpose, and it acts with your user rights. The risk is not a virus. The risk is what the agent can reach and what can talk it into using it.
That is the reasoning behind the product, not a measured detection statistic: if the dangerous thing is a secret lying in a readable file, the useful check is "which secrets lie where, and who can read them", not "does this binary match a signature".
This is not theory. In August 2025 the s1ngularity attack on the Nx build tool shipped malware that looked for installed AI command-line assistants and told them to search the victim's file system for credentials, SSH keys and wallets. Weeks later the self-spreading Shai-Hulud npm worm hunted for publishing tokens. Both worked because the secrets were sitting where any process of the user could read them.
.env db + API keys
~/.ssh/id_ed25519
~/.aws/credentials
shell history
agent ❯ cat .env agent ❯ ls ~/.ssh agent ❯ history agent ❯ ▌
grrr… let me in
same rights as you · nothing has to be installed for this to go wrong
Your agent wants in. It already has your rights, so what it can reach is what matters.
What Dante is
Dante is a free, open command-line watchman for machines you own: your Mac, Linux or Windows computer, and servers you can reach by SSH. You talk to it in plain German, English or Spanish. It reads, it explains, and it changes nothing without a clear yes.
Every verdict carries the real line it came from. If Dante cannot read the proof, the verdict is unknown . It never says safe on an assumption.
Silent by default. No telemetry, no background update check, no account. The free state runs fully offline.
Zero dependencies. One small Node.js program you can read. The source is public .
user ❯ check my computer
SAFE Openly readable secrets in the web directory
what: No openly readable .env/.sql/.git files in the web root.
proof:
• Web root searched, nothing readable found
↳ dante:find-files /var/www (.env, *.sql, …) → (no matches in: /var/www)
UNKNOWN SSH: root login and password login
what: No SSH service listens on port 22 here; I do not know whether the settings matter.
proof:
• PermitRootLogin = (not set)
↳ /etc/ssh/sshd_config → PermitRootLogin default Every verdict carries the line it came from. If Dante cannot read the proof, it says UNKNOWN, never SAFE.
What it checks
Secrets lying around : API keys, tokens and private keys in project files, .env files and Git history, matched against a library of 1,111 secret formats. It shows file, line and kind - never the value.
Your servers : SSH settings, world-writable files in web folders, known PHP vulnerabilities by version, signs of web shells.
AI agent configuration : MCP server entries and agent settings (with the optional Hugging Face analysis pack), reported without ever printing a key.
Bait files (the "Wachschutz"): fake credentials that mail you the moment someone uses them.
Windows: services anyone can overwrite, unquoted service paths, writable PATH folders, AlwaysInstallElevated, stored autologon passwords.
Is it any good? An honest answer
Dante is young. It is a watchman, not a real-time engine: it looks when you ask or on the schedule you agreed to, and it does not scan running processes for malware. What it does well is the boring, high-value part: finding the secrets and settings an attacker or a hijacked agent would use first, and refusing to claim more than it can prove. Its language layer (what you type) is trained on about 259,000 labelled sentences and answers "I don't know" instead of guessing when unsure.
STEP 1 LOOK ████████████████████ free · 3 minutes · reads only
STEP 2 DIG DEEP ██████████████░░░░░░ secrets in your project folders
STEP 3 BAIT █████████░░░░░░░░░░░ a fake key that rings when it is used
STEP 4 STAY SHARP ████░░░░░░░░░░░░░░░░ fresh knowledge packs for new attacks Most people stop after step 1 and are already better off. Each step is its own yes.
You never have to decide everything now. Start with step 1. Dante asks before every next thing, and every question shows the exact command first.
Install Node.js (nodejs.org, once), open a terminal, paste one line:
you@MacBook ~ % npx dante-watch
I am Dante. I walk through hell and bring back the map.
● awake · I watch over localhost.
user ❯ check my computer Mac: Terminal
PowerShell — Windows PS C:\Users\you> npx dante-watch
I am Dante. I walk through hell and bring back the map.
● awake · I watch over localhost.
user ❯ check my computer Windows: PowerShell
Same program, same words, on every system. Not sure where your terminal is? The install page shows it for Mac, Windows and Linux, step by step.
Details for every system, the free and paid parts and what each costs are on the install page .
Look, nobody wakes up wanting to audit their laptop. It is boring, it feels like homework, and it always looks fine until the day it does not.
So here is the deal. It takes three minutes. It only reads. It costs nothing. If it finds nothing, you wasted three minutes and you can say you checked. If it finds a key lying around, you found it before somebody else did.
Your agent has the keys to your house either way. The only question is whether you know which keys.
They need protection, but not the kind that matches malware signatures. An agent is legitimate software running with your rights. What matters is which secrets and servers it can reach and whether anything can instruct it. Checking where secrets lie, how SSH is set and what the agent configuration allows covers the real exposure.
No. Dante is a watchman for machines you own. It does not scan for malware in real time. It finds exposed secrets, risky settings and known vulnerabilities, and shows the proof line for each result.
No. By default it sends nothing. Only three actions use the network (update, the catalog, and the optional bait watch) and each starts only after you answer yes to the command Dante prints first.
Yes. The free state is complete and works offline. Extra knowledge packs and the bait watch are optional; prices are on the install page.
