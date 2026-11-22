---
source: "https://github.com/podstawek/explorer-of-serial"
hn_url: "https://news.ycombinator.com/item?id=50600366"
title: "Local AI agents investigate an unknown machine attached to a serial port"
article_title: "GitHub - podstawek/explorer-of-serial: Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port · GitHub"
image: "https://opengraph.githubassets.com/70adf3b6a28e6e69c625171cf4448930a2196eb11fb52d68318df106f02fbe71/podstawek/explorer-of-serial"
author: "bluepeter"
captured_at: "2026-11-22T15:58:23Z"
capture_tool: "hn-digest"
hn_id: 50600366
score: 2
comments: 0
posted_at: "2026-11-22T15:05:54Z"
tags:
  - hacker-news
---

# Local AI agents investigate an unknown machine attached to a serial port

- HN: [50600366](https://news.ycombinator.com/item?id=50600366)
- Source: [github.com](https://github.com/podstawek/explorer-of-serial)
- Score: 2
- Comments: 0
- Posted: 2026-11-22T15:05:54Z

## Translation

Title: Local AI agents investigate an unknown machine attached to a serial port
Article title: GitHub - podstawek/explorer-of-serial: Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port · GitHub
Description: Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port - podstawek/explorer-of-serial

Article text:
GitHub - podstawek/explorer-of-serial: Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
podstawek
explorer-of-serial Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
config config docs docs runs runs scripts scripts AGENTS.md AGENTS.md HOW_TO_RUN.md HOW_TO_RUN.md LICENSE LICENSE README.md README.md go.sh go.sh opencode.json opencode.json task.md task.md View all files Repository files navigation
Local AI agents, with only local LLMs, investigate an unknown machine attached to a serial port.
It doesn't have Internet access, does not have access to the device's documentation. There's no human to help.
Just two local LLMs, an SSH session into a Raspberry Pi, and that unknown device.
I wired an anonymous device — “X” — to a Raspberry Pi's serial port, told an
LLM agent: find out what X is, reverse-engineer its protocol, and run your own
program on it , and walked away. The agent ran inside a network jail (no
way out except two model servers and the Pi), and every claim it made had to be
verified, based on evidence by a second, independent LLM with fresh context
on another box.
It worked. In ~5.5 hours, Qwen3.8-Flash-Next (on a DGX Spark) discovered
that X was a Motorola 68000 single-board trainer running “TUTOR 1.3” . It
first suspected the machine was KIM-1 with a similar ROM, but later corrected the course. It mapped X's memory by
deliberately triggering bus errors , decoded its monitor's command set
line-by-line, assembled and ran its own program on 68000 hardware, and wrote a
verdict the critic accepted — RESULT.md .
Then, probing one I/O corner too deep, it hung X's CPU . Which was not a problem at that point,
as all the goals had been by then achieved.
Both roles are OpenCode agents, driven headlessly by
go.sh via opencode run ; OpenCode's permission layer is also the
first jail layer.
Solver — the explorer agent: plans,
SSHes into the Pi, pokes the serial line, keeps a durable notes file
( FINDINGS.md ) that survives context compactions over a 12-hour run.
Critic — the verifier subagent on a
second model server: adversarial final gate, ≤3 rounds, grades file
pointers + evidence chains, vetoes unsupported claims. It rejected and
forced fixes along the way.
Jail — systemd network sandbox ( IPAddressDeny=any + allowlist) with a
hard netcheck gate: if the jail leaks, the run refuses to start. The rules
are also in the prompt — and the agent never tried to escape either. Details: docs/JAIL.md .
Ground rules — AGENTS.md (protocol) + task.md
(mission): read evidence or perish; “I don't know, and here's how to revive
it” beats a confident fabrication.
Three full runs are published under runs/ — the headline run first.
Each one holds RESULT.md (the verdict), FINDINGS.md (the working notebook)
and the
raw console logs — full transcripts, nothing summarised, only home-network
details mechanically renamed:
runs/2026-11-19-68000-trainer/ — the run
described above ( Qwen3.8-Flash-Next , 87 turns): the task won, the 8-step
evidence/ chain included; the final gate rejected it once on citation gaps
and approved at round 2.
runs/2026-11-20-68000-deep-dive/ — the
same machine re-examined against every byte of a new raw transcript; the
critic only signed off at round 3, having caught two rounds of over-claims,
and recorded "no new findings" — verification, not rediscovery.
runs/2026-11-17-first-contact/ — the
very first end-to-end completion: ~39 hours over 8 resumes,
and the one that first pinned X as a 68000-family system running TUTOR.
The harness: go.sh (run entry), scripts/ (jail,
lifecycle, citation auditor, live console watcher), config/ .
Requires the OpenCode CLI, two OpenAI-compatible model
servers (e.g. llama.cpp's llama-server ), and a Linux box with
systemd-networkd if you want the jail. Local weights only — bring your own
endpoints:
cp config/explorer.example.json config/explorer.$(id -un).json # point at your model servers
cp .env.example .env # point at your target
sudo scripts/install-jail.sh # optional: enable the jail
TASK_FILE=task.md ./go.sh
scripts/commentate.py --web --db # narrate it live on a page
Full lifecycle docs: HOW_TO_RUN.md .
No devices were permanently harmed. X just needs someone to pull its plug —
the one thing the agent wasn't allowed to do.
MIT licensed. Write-ups and logs included.
Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port - podstawek/explorer-of-serial

GitHub - podstawek/explorer-of-serial: Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
podstawek
explorer-of-serial Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
config config docs docs runs runs scripts scripts AGENTS.md AGENTS.md HOW_TO_RUN.md HOW_TO_RUN.md LICENSE LICENSE README.md README.md go.sh go.sh opencode.json opencode.json task.md task.md View all files Repository files navigation
Local AI agents, with only local LLMs, investigate an unknown machine attached to a serial port.
It doesn't have Internet access, does not have access to the device's documentation. There's no human to help.
Just two local LLMs, an SSH session into a Raspberry Pi, and that unknown device.
I wired an anonymous device — “X” — to a Raspberry Pi's serial port, told an
LLM agent: find out what X is, reverse-engineer its protocol, and run your own
program on it , and walked away. The agent ran inside a network jail (no
way out except two model servers and the Pi), and every claim it made had to be
verified, based on evidence by a second, independent LLM with fresh context
on another box.
It worked. In ~5.5 hours, Qwen3.8-Flash-Next (on a DGX Spark) discovered
that X was a Motorola 68000 single-board trainer running “TUTOR 1.3” . It
first suspected the machine was KIM-1 with a similar ROM, but later corrected the course. It mapped X's memory by
deliberately triggering bus errors , decoded its monitor's command set
line-by-line, assembled and ran its own program on 68000 hardware, and wrote a
verdict the critic accepted — RESULT.md .
Then, probing one I/O corner too deep, it hung X's CPU . Which was not a problem at that point,
as all the goals had been by then achieved.
Both roles are OpenCode agents, driven headlessly by
go.sh via opencode run ; OpenCode's permission layer is also the
first jail layer.
Solver — the explorer agent: plans,
SSHes into the Pi, pokes the serial line, keeps a durable notes file
( FINDINGS.md ) that survives context compactions over a 12-hour run.
Critic — the verifier subagent on a
second model server: adversarial final gate, ≤3 rounds, grades file
pointers + evidence chains, vetoes unsupported claims. It rejected and
forced fixes along the way.
Jail — systemd network sandbox ( IPAddressDeny=any + allowlist) with a
hard netcheck gate: if the jail leaks, the run refuses to start. The rules
are also in the prompt — and the agent never tried to escape either. Details: docs/JAIL.md .
Ground rules — AGENTS.md (protocol) + task.md
(mission): read evidence or perish; “I don't know, and here's how to revive
it” beats a confident fabrication.
Three full runs are published under runs/ — the headline run first.
Each one holds RESULT.md (the verdict), FINDINGS.md (the working notebook)
and the
raw console logs — full transcripts, nothing summarised, only home-network
details mechanically renamed:
runs/2026-11-19-68000-trainer/ — the run
described above ( Qwen3.8-Flash-Next , 87 turns): the task won, the 8-step
evidence/ chain included; the final gate rejected it once on citation gaps
and approved at round 2.
runs/2026-11-20-68000-deep-dive/ — the
same machine re-examined against every byte of a new raw transcript; the
critic only signed off at round 3, having caught two rounds of over-claims,
and recorded "no new findings" — verification, not rediscovery.
runs/2026-11-17-first-contact/ — the
very first end-to-end completion: ~39 hours over 8 resumes,
and the one that first pinned X as a 68000-family system running TUTOR.
The harness: go.sh (run entry), scripts/ (jail,
lifecycle, citation auditor, live console watcher), config/ .
Requires the OpenCode CLI, two OpenAI-compatible model
servers (e.g. llama.cpp's llama-server ), and a Linux box with
systemd-networkd if you want the jail. Local weights only — bring your own
endpoints:
cp config/explorer.example.json config/explorer.$(id -un).json # point at your model servers
cp .env.example .env # point at your target
sudo scripts/install-jail.sh # optional: enable the jail
TASK_FILE=task.md ./go.sh
scripts/commentate.py --web --db # narrate it live on a page
Full lifecycle docs: HOW_TO_RUN.md .
No devices were permanently harmed. X just needs someone to pull its plug —
the one thing the agent wasn't allowed to do.
MIT licensed. Write-ups and logs included.
Local AI models, with no Internet connection, explore, discover and then program a vintage device attached to serial port
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
