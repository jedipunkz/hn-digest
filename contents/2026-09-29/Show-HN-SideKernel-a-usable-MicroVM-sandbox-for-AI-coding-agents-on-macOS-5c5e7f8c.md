---
source: "https://github.com/minoansecurity/sidekernel"
hn_url: "https://news.ycombinator.com/item?id=49891008"
title: "Show HN: SideKernel – a usable MicroVM sandbox for AI coding agents on macOS"
article_title: "GitHub - minoansecurity/sidekernel: An easy-to-use, secure macOS sandbox for Claude and other AI coding agents · GitHub"
image: "https://opengraph.githubassets.com/cf07e3ca5097b1ba8bbb71e3fa8fa98bc92bdc90976020aa43c420883f4f2396/minoansecurity/sidekernel"
author: "dimiprasakis"
captured_at: "2026-09-29T11:10:11Z"
capture_tool: "hn-digest"
hn_id: 49891008
score: 1
comments: 0
posted_at: "2026-09-29T10:45:43Z"
tags:
  - hacker-news
---

# Show HN: SideKernel – a usable MicroVM sandbox for AI coding agents on macOS

- HN: [49891008](https://news.ycombinator.com/item?id=49891008)
- Source: [github.com](https://github.com/minoansecurity/sidekernel)
- Score: 1
- Comments: 0
- Posted: 2026-09-29T10:45:43Z

## Translation

Title: Show HN: SideKernel – a usable MicroVM sandbox for AI coding agents on macOS
Article title: GitHub - minoansecurity/sidekernel: An easy-to-use, secure macOS sandbox for Claude and other AI coding agents · GitHub
Description: An easy-to-use, secure macOS sandbox for Claude and other AI coding agents - minoansecurity/sidekernel
HN text: Hey everyone, SideKernel is the result of my capstone project for my MSc at Georgia Tech. There are a lot of sandboxes on the market right now for AI agents, but most of them focus on production deployments, rather than coding agents (e.g. Claude Code). I always felt uncomfortable running Claude Code or other coding agents directly on my machine, and I didn't want to buy a second computer or rent a VPS just for that. That's why I made SideKernel: a sandbox that tries to be as transparent as possible to the everyday developer. I hope you find some value in it, and feedback is more than welcome. Usability is hard to get right without large user studies, so I'd really value hearing about any annoyances or blockers you run into. Thanks!

Article text:
GitHub - minoansecurity/sidekernel: An easy-to-use, secure macOS sandbox for Claude and other AI coding agents · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
minoansecurity
/
sidekernel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
9 Commits 9 Commits Folders and files
.github/ workflows .github/ workflows Formula Formula Host Host agent agent assets assets guest guest .gitignore .gitignore LICENSE LICENSE Makefile Makefile Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md sidekernel.entitlements sidekernel.entitlements View all files Repository files navigation
Paper ·
Site ·
Note from the developer
SideKernel is a usable sandbox for AI coding agents (e.g. Claude Code). "Usable" means it tries stays out of your way and to feel as if its not there. It is developed as a capstone project for Georgia Tech's MSc in Cybersecurity.
Beyond AI agents, SideKernel is useful for trying out software without installing it on your host (e.g. untrusted npm packages).
The current folder is the sandbox: files sync both ways, and AI conversations persist across restarts.
Ports opened in the sandbox are auto-forwarded to the host.
Copy/paste of text and images works in and out of the sandbox (with most VMs it doesn't).
The host's Claude config (skills, plugins) carries over to the sandbox.
A network kill switch blocks all traffic when the sandbox holds sensitive data, while Claude keeps working.
Log in to Claude once, on the host or in the sandbox, and both are authenticated.
Non-mounted files are easy to bring in with sk-drop <path> , or by dragging and dropping them into Claude.
The in-sandbox save command creates a personal layer that persists files, configs and installations across sandboxes.
SideKernel is still in research preview and not yet ready for production use. Use it responsibly. Please read the Paper or the sidekernel.com to learn more.
Tested on M1 and M4 (macOS 26.2); other Apple silicon chips should work.
Install with brew (recommended)
brew tap minoansecurity/sidekernel https://github.com/minoansecurity/sidekernel && brew trust --formula minoansecurity/sidekernel/sidekernel
brew install sidekernel
Install directly from source:
rustup target add aarch64-unknown-linux-musl
git clone https://github.com/minoansecurity/sidekernel
cd sidekernel
make install
The first run builds the root filesystem so it might take a minute or two.
On the host (from a project folder):
sk # launch an ephemeral microVM, current directory mounted
sidekernel # (alias: sk)
sclaude # launch a sandbox and start Claude Code directly
Coming soon: scodex , sgemini , sgrok , and more.
sk-drop < host-path > # copy a host file into the sandbox (requires approval on the host)
sk-net on/off # block all outbound traffic
save # persist installed packages across sandboxes
fightsong # Print's Georgia Tech's fight song on the terminal 🐝 (alias: ramblinwreck)
You can also drag and drop files directly into Claude Code.
The agent may read, edit or destroy anything in the mounted directory.
A malicious agent can open ports to the host, exposing malicious services.
Only Claude Code is integrated; other harnesses such as Codex are planned.
SideKernel's security rests on its architecture, but the implementation has not had a formal security review.
SideKernel is not yet notarized (it is self-signed).
The full list is in the paper and on the website .
SideKernel is open-source software, licensed under the Apache License, Version 2.0 .
An easy-to-use, secure macOS sandbox for Claude and other AI coding agents
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

An easy-to-use, secure macOS sandbox for Claude and other AI coding agents - minoansecurity/sidekernel

Hey everyone, SideKernel is the result of my capstone project for my MSc at Georgia Tech. There are a lot of sandboxes on the market right now for AI agents, but most of them focus on production deployments, rather than coding agents (e.g. Claude Code). I always felt uncomfortable running Claude Code or other coding agents directly on my machine, and I didn't want to buy a second computer or rent a VPS just for that. That's why I made SideKernel: a sandbox that tries to be as transparent as possible to the everyday developer. I hope you find some value in it, and feedback is more than welcome. Usability is hard to get right without large user studies, so I'd really value hearing about any annoyances or blockers you run into. Thanks!

GitHub - minoansecurity/sidekernel: An easy-to-use, secure macOS sandbox for Claude and other AI coding agents · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
minoansecurity
/
sidekernel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
9 Commits 9 Commits Folders and files
.github/ workflows .github/ workflows Formula Formula Host Host agent agent assets assets guest guest .gitignore .gitignore LICENSE LICENSE Makefile Makefile Package.swift Package.swift README.md README.md SECURITY.md SECURITY.md sidekernel.entitlements sidekernel.entitlements View all files Repository files navigation
Paper ·
Site ·
Note from the developer
SideKernel is a usable sandbox for AI coding agents (e.g. Claude Code). "Usable" means it tries stays out of your way and to feel as if its not there. It is developed as a capstone project for Georgia Tech's MSc in Cybersecurity.
Beyond AI agents, SideKernel is useful for trying out software without installing it on your host (e.g. untrusted npm packages).
The current folder is the sandbox: files sync both ways, and AI conversations persist across restarts.
Ports opened in the sandbox are auto-forwarded to the host.
Copy/paste of text and images works in and out of the sandbox (with most VMs it doesn't).
The host's Claude config (skills, plugins) carries over to the sandbox.
A network kill switch blocks all traffic when the sandbox holds sensitive data, while Claude keeps working.
Log in to Claude once, on the host or in the sandbox, and both are authenticated.
Non-mounted files are easy to bring in with sk-drop <path> , or by dragging and dropping them into Claude.
The in-sandbox save command creates a personal layer that persists files, configs and installations across sandboxes.
SideKernel is still in research preview and not yet ready for production use. Use it responsibly. Please read the Paper or the sidekernel.com to learn more.
Tested on M1 and M4 (macOS 26.2); other Apple silicon chips should work.
Install with brew (recommended)
brew tap minoansecurity/sidekernel https://github.com/minoansecurity/sidekernel && brew trust --formula minoansecurity/sidekernel/sidekernel
brew install sidekernel
Install directly from source:
rustup target add aarch64-unknown-linux-musl
git clone https://github.com/minoansecurity/sidekernel
cd sidekernel
make install
The first run builds the root filesystem so it might take a minute or two.
On the host (from a project folder):
sk # launch an ephemeral microVM, current directory mounted
sidekernel # (alias: sk)
sclaude # launch a sandbox and start Claude Code directly
Coming soon: scodex , sgemini , sgrok , and more.
sk-drop < host-path > # copy a host file into the sandbox (requires approval on the host)
sk-net on/off # block all outbound traffic
save # persist installed packages across sandboxes
fightsong # Print's Georgia Tech's fight song on the terminal 🐝 (alias: ramblinwreck)
You can also drag and drop files directly into Claude Code.
The agent may read, edit or destroy anything in the mounted directory.
A malicious agent can open ports to the host, exposing malicious services.
Only Claude Code is integrated; other harnesses such as Codex are planned.
SideKernel's security rests on its architecture, but the implementation has not had a formal security review.
SideKernel is not yet notarized (it is self-signed).
The full list is in the paper and on the website .
SideKernel is open-source software, licensed under the Apache License, Version 2.0 .
An easy-to-use, secure macOS sandbox for Claude and other AI coding agents
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
