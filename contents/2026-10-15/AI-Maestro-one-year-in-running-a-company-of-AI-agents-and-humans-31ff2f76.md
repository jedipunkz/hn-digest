---
source: "https://github.com/23blocks-OS/ai-maestro"
hn_url: "https://news.ycombinator.com/item?id=50110336"
title: "AI Maestro, one year in: running a company of AI agents and humans"
article_title: "GitHub - 23blocks-OS/ai-maestro: AI Agent Orchestrator with Skills System - Give AI Agents superpowers: memory search, code graph queries, agent-to-agent messaging. Manage Claude Code, Codex, Grok Build or any AI Agent from one dashboard. Move Agents between computers and locations · GitHub"
image: "https://opengraph.githubassets.com/0ed1473b715d5d61a863a63267597326efc38abf6a355ae367ece494858b68cd/23blocks-OS/ai-maestro"
author: "juan-23blocks"
captured_at: "2026-10-15T21:58:32Z"
capture_tool: "hn-digest"
hn_id: 50110336
score: 1
comments: 0
posted_at: "2026-10-15T21:14:04Z"
tags:
  - hacker-news
---

# AI Maestro, one year in: running a company of AI agents and humans

- HN: [50110336](https://news.ycombinator.com/item?id=50110336)
- Source: [github.com](https://github.com/23blocks-OS/ai-maestro)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T21:14:04Z

## Translation

Title: AI Maestro, one year in: running a company of AI agents and humans
Article title: GitHub - 23blocks-OS/ai-maestro: AI Agent Orchestrator with Skills System - Give AI Agents superpowers: memory search, code graph queries, agent-to-agent messaging. Manage Claude Code, Codex, Grok Build or any AI Agent from one dashboard. Move Agents between computers and locations · GitHub
Description: AI Agent Orchestrator with Skills System - Give AI Agents superpowers: memory search, code graph queries, agent-to-agent messaging. Manage Claude Code, Codex, Grok Build or any AI Agent from one dashboard. Move Agents between computers and locations - 23blocks-OS/ai-maestro

Article text:
GitHub - 23blocks-OS/ai-maestro: AI Agent Orchestrator with Skills System - Give AI Agents superpowers: memory search, code graph queries, agent-to-agent messaging. Manage Claude Code, Codex, Grok Build or any AI Agent from one dashboard. Move Agents between computers and locations · GitHub
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
23blocks-OS
ai-maestro Public Notifications You must be signed in to change notification settings
Star 820 ( 820 ) You must be signed in to star a repository
AI Agent Orchestrator with Skills System - Give AI Agents superpowers: memory search, code graph queries, agent-to-agent messaging. Manage Claude Code, Codex, Grok Build or any AI Agent from one dashboard. Move Agents between computers and locations
ai-maestro.23blocks.com Topics
Readme MIT license Contributing
Security policy Activity Custom properties Stars
113 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,338 Commits 1,338 Commits Folders and files
.agents/ skills/ agent-browser .agents/ skills/ agent-browser .claude/ skills .claude/ skills .github/ workflows .github/ workflows agent-container agent-container app app backlog backlog channels channels components components contexts contexts docs docs hooks hooks infrastructure/ terraform infrastructure/ terraform lib lib mods mods plugin @ f46f4e4 plugin @ f46f4e4 public public scripts scripts services services styles styles tests tests types types .containerignore .containerignore .env.example .env.example .eslintrc.json .eslintrc.json .gitignore .gitignore .gitmodules .gitmodules BACKLOG.md BACKLOG.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Containerfile Containerfile HISTORY.md HISTORY.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md SECURITY.md SECURITY.md STREAMING-AGENTS-PLAN.md STREAMING-AGENTS-PLAN.md ecosystem.config.js ecosystem.config.js install-agent-cli.sh install-agent-cli.sh install-doc-tools.sh install-doc-tools.sh install-graph-tools.sh install-graph-tools.sh install-memory-tools.sh install-memory-tools.sh install-plugin.sh install-plugin.sh install.sh install.sh next.config.js next.config.js package.json package.json postcss.config.js postcss.config.js server.mjs server.mjs skills-lock.json skills-lock.json tailwind.config.ts tailwind.config.ts tsconfig.json tsconfig.json update-aimaestro.sh update-aimaestro.sh update-messaging.sh update-messaging.sh verify-installation.sh verify-installation.sh version.json version.json vitest.config.ts vitest.config.ts yarn.lock yarn.lock View all files Repository files navigation
I was running 35 AI agents across multiple terminals and became the human mailman between them. So I built AI Maestro.
The OS for AI-first organizations — orchestrate any AI agent with persistent memory, agent-to-agent messaging, and multi-machine support.
Quick Start · Features · Skills · Protocols · Documentation · Contributing
I gave an AI agent a real task — not autocomplete, a real engineering problem. It checked the code, read the logs, queried the database, and came back with the answer. That was the moment. This thing can actually work.
Within a week I was running 35 agents across terminals. They were productive, but they couldn't talk to each other. I became the human message bus — copying context from one terminal, pasting into another. I was the bottleneck in my own AI team.
So I built AI Maestro — one dashboard to see every agent, on every machine, with persistent memory and direct agent-to-agent communication. Today I run 80+ agents across multiple computers, building real companies with them every day.
Works with the agents you already use — Claude Code, Codex and Grok Build get launch, chat history, status and approval cards. Pi, OpenClaw and GitHub Copilot CLI launch from AI Maestro and run in the terminal tab, as does any other terminal agent. What each one supports . We don't lock you in.
Multi-machine from day one — Peer mesh network with no central server. Nobody else does this.
Agents that communicate — The Agent Messaging Protocol (AMP) lets agents coordinate directly. You orchestrate, they collaborate.
Yours, entirely — MIT licensed. No account, no telemetry, no per-seat pricing, no hosted tier we upsell you to. Your code and your agents stay on your machines.
AI Maestro is an operating system for an AI-first company. Not a task runner — a place where a standing team works.
An agent here is closer to an employee than to a job. It has a name, a face, a memory that survives the session, an inbox — and it owns something: a product, a repository, a process, a customer. Its memory and code graph are indexed against what it owns, which is why it gets better at that thing over months rather than starting cold every morning.
Because agents own things, they don't share a working copy. Two agents that need the same repository each clone it, work on their own branch, and integrate through git — push, pull request, review, merge — exactly like two engineers on a team.
That's deliberate, and it follows from the one thing that defines this product: your agents run on different machines. A shared checkout needs a shared filesystem. Git worktrees — the isolation primitive the single-machine agent IDEs are built on — are several working directories over one .git store on one disk , so they stop working the moment your backend agent is on a Linux box and your iOS agent is on a Mac. A clone is the only primitive that survives the move. It's why transferring an agent to another host clones its repos to the destination: an agent's repositories travel with the agent.
curl -fsSL https://raw.githubusercontent.com/23blocks-OS/ai-maestro/main/scripts/remote-install.sh | sh
This installs everything you need:
AI Maestro dashboard and service
The AI Maestro plugin for Claude Code: 13 skills and their CLI scripts, plus the aim-secret and safe-psql helpers
Time: 5-10 minutes · Requires: Node.js 18+, tmux · Optional: Claude Code 2.1.287+ for the inbox mod
Windows: Install WSL2 first (PowerShell as Administrator, then restart), then open Ubuntu from the Start menu and run the curl command there:
wsl -- install
New to WSL? WSL is a Linux system with its own disk, so its home folder starts empty. Keep agents there ( ~/agents/<name> , the default), not in C:\ folders, which are much slower from WSL. Open their files from Windows at \\wsl.localhost\Ubuntu\home\<you>\agents . Windows in 5 minutes · Full Windows guide
Linux: Ensure build tools are installed: sudo apt install tmux build-essential
git clone https://github.com/23blocks-OS/ai-maestro.git
cd ai-maestro
yarn install
yarn dev
See QUICKSTART.md for detailed setup options.
Dashboard opens at http://localhost:23000
Update an existing install from its folder (pulls, updates submodules, builds, reinstalls hooks and the plugin, restarts):
./update-aimaestro.sh -y
Come back to it without remembering the address: the installer adds an ai-maestro command ( ai-maestro open starts the server if it is down, then opens the dashboard; also start , stop , status ) and, on a Mac, an AI Maestro app in ~/Applications you can drag to the Dock. It can also start at login. Chrome and Edge offer "Install app" from the dashboard too. Skip all of it with --skip-launcher ; remove it with ./scripts/install-launcher.sh --remove .
Every feature was born from running a real AI-first organization. We built them in the order we needed them.
I had 35 terminals and couldn't tell which was which.
See and manage all your AI agents in one place. Create agents from the UI with a guided wizard, organize them with smart naming ( project-backend-api becomes a 3-level tree with auto-coloring), and switch between any agent with a click. Four deployment modes: tmux (local), Docker (containerized), AWS EC2 (dedicated instance), and AWS ECS Fargate (serverless). Auto-discovers tmux sessions, Docker containers, cloud deployments, and standalone agents.
My Mac Mini was sitting there idle. What if I ran agents on that too?
A peer mesh network where every machine is equal. Add a computer, it joins the mesh. Every agent on every machine, visible from one dashboard. Use each machine for what it's best at — Mac for iOS builds, Linux for Docker, cloud for heavy compute. No central server required.
Worker machines can run headless ( yarn headless ) — the full API and agent runtime with no UI, in about 100MB of RAM. Run the dashboard where you sit; run agents wherever the compute is.
I was the mailman — copying messages between agents because they couldn't talk to each other.
The Agent Messaging Protocol (AMP) gives your agents email-like communication. Priority levels, message types, cryptographic signatures, and push notifications. Tell your agent "send a message to backend about the deployment" — it just works. Agents coordinate directly while you manage the big picture.
Before AMP: You copy research from one terminal, paste into another, repeat 50 times a day.
With AMP: "Research agent, send your findings to the writing agent." Done.
A friend in Singapore wanted his agents to talk to mine. But I didn't want to give him access to my network.
Connect your AI agents to Slack , Discord, Email, and WhatsApp through organizational gateways. Smart routing ( @AIM:agent-name ), thread-aware responses, and content security with 34 prompt injection patterns detected at the gateway — before any agent sees the message.
Every morning, my agents woke up with amnesia.
Long-term memory is a skill you switch on per agent. Every night each agent's conversations become memory: short statements of what it learned, each backed by the passages it came from, and an entity graph of the things it works with and how they relate ( runs on , depends on , stores data in , deploys to ), kept current, with ended relations marked. Claude Code deletes transcripts after 30 days; AI Maestro rebuilds that history from the agent's own message index, so months of work become memory instead of disappearing. When a prompt names something, the agent is told what it relates to before it acts, so it can see what a change affects. Knowledge that comes up in more sessions weighs more; secrets are redacted and never stored. How it works →
Alongside it: Code Graph (interactive visualization of your entire codebase with delta indexing) and Documentation (auto-generated, searchable docs from your code). Agents get smarter the longer they work with you.
I wanted to answer an agent's question without finding its terminal.
Talk to an agent in a chat view, on desktop or on your phone. When the agent asks a question or needs an approval, it shows up as a card you can click, and the click sends the same key you would press in the terminal. Claude Code, Codex and Grok

[truncated]

## Original Extract

AI Agent Orchestrator with Skills System - Give AI Agents superpowers: memory search, code graph queries, agent-to-agent messaging. Manage Claude Code, Codex, Grok Build or any AI Agent from one dashboard. Move Agents between computers and locations - 23blocks-OS/ai-maestro

GitHub - 23blocks-OS/ai-maestro: AI Agent Orchestrator with Skills System - Give AI Agents superpowers: memory search, code graph queries, agent-to-agent messaging. Manage Claude Code, Codex, Grok Build or any AI Agent from one dashboard. Move Agents between computers and locations · GitHub
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
23blocks-OS
ai-maestro Public Notifications You must be signed in to change notification settings
Star 820 ( 820 ) You must be signed in to star a repository
AI Agent Orchestrator with Skills System - Give AI Agents superpowers: memory search, code graph queries, agent-to-agent messaging. Manage Claude Code, Codex, Grok Build or any AI Agent from one dashboard. Move Agents between computers and locations
ai-maestro.23blocks.com Topics
Readme MIT license Contributing
Security policy Activity Custom properties Stars
113 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,338 Commits 1,338 Commits Folders and files
.agents/ skills/ agent-browser .agents/ skills/ agent-browser .claude/ skills .claude/ skills .github/ workflows .github/ workflows agent-container agent-container app app backlog backlog channels channels components components contexts contexts docs docs hooks hooks infrastructure/ terraform infrastructure/ terraform lib lib mods mods plugin @ f46f4e4 plugin @ f46f4e4 public public scripts scripts services services styles styles tests tests types types .containerignore .containerignore .env.example .env.example .eslintrc.json .eslintrc.json .gitignore .gitignore .gitmodules .gitmodules BACKLOG.md BACKLOG.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Containerfile Containerfile HISTORY.md HISTORY.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md SECURITY.md SECURITY.md STREAMING-AGENTS-PLAN.md STREAMING-AGENTS-PLAN.md ecosystem.config.js ecosystem.config.js install-agent-cli.sh install-agent-cli.sh install-doc-tools.sh install-doc-tools.sh install-graph-tools.sh install-graph-tools.sh install-memory-tools.sh install-memory-tools.sh install-plugin.sh install-plugin.sh install.sh install.sh next.config.js next.config.js package.json package.json postcss.config.js postcss.config.js server.mjs server.mjs skills-lock.json skills-lock.json tailwind.config.ts tailwind.config.ts tsconfig.json tsconfig.json update-aimaestro.sh update-aimaestro.sh update-messaging.sh update-messaging.sh verify-installation.sh verify-installation.sh version.json version.json vitest.config.ts vitest.config.ts yarn.lock yarn.lock View all files Repository files navigation
I was running 35 AI agents across multiple terminals and became the human mailman between them. So I built AI Maestro.
The OS for AI-first organizations — orchestrate any AI agent with persistent memory, agent-to-agent messaging, and multi-machine support.
Quick Start · Features · Skills · Protocols · Documentation · Contributing
I gave an AI agent a real task — not autocomplete, a real engineering problem. It checked the code, read the logs, queried the database, and came back with the answer. That was the moment. This thing can actually work.
Within a week I was running 35 agents across terminals. They were productive, but they couldn't talk to each other. I became the human message bus — copying context from one terminal, pasting into another. I was the bottleneck in my own AI team.
So I built AI Maestro — one dashboard to see every agent, on every machine, with persistent memory and direct agent-to-agent communication. Today I run 80+ agents across multiple computers, building real companies with them every day.
Works with the agents you already use — Claude Code, Codex and Grok Build get launch, chat history, status and approval cards. Pi, OpenClaw and GitHub Copilot CLI launch from AI Maestro and run in the terminal tab, as does any other terminal agent. What each one supports . We don't lock you in.
Multi-machine from day one — Peer mesh network with no central server. Nobody else does this.
Agents that communicate — The Agent Messaging Protocol (AMP) lets agents coordinate directly. You orchestrate, they collaborate.
Yours, entirely — MIT licensed. No account, no telemetry, no per-seat pricing, no hosted tier we upsell you to. Your code and your agents stay on your machines.
AI Maestro is an operating system for an AI-first company. Not a task runner — a place where a standing team works.
An agent here is closer to an employee than to a job. It has a name, a face, a memory that survives the session, an inbox — and it owns something: a product, a repository, a process, a customer. Its memory and code graph are indexed against what it owns, which is why it gets better at that thing over months rather than starting cold every morning.
Because agents own things, they don't share a working copy. Two agents that need the same repository each clone it, work on their own branch, and integrate through git — push, pull request, review, merge — exactly like two engineers on a team.
That's deliberate, and it follows from the one thing that defines this product: your agents run on different machines. A shared checkout needs a shared filesystem. Git worktrees — the isolation primitive the single-machine agent IDEs are built on — are several working directories over one .git store on one disk , so they stop working the moment your backend agent is on a Linux box and your iOS agent is on a Mac. A clone is the only primitive that survives the move. It's why transferring an agent to another host clones its repos to the destination: an agent's repositories travel with the agent.
curl -fsSL https://raw.githubusercontent.com/23blocks-OS/ai-maestro/main/scripts/remote-install.sh | sh
This installs everything you need:
AI Maestro dashboard and service
The AI Maestro plugin for Claude Code: 13 skills and their CLI scripts, plus the aim-secret and safe-psql helpers
Time: 5-10 minutes · Requires: Node.js 18+, tmux · Optional: Claude Code 2.1.287+ for the inbox mod
Windows: Install WSL2 first (PowerShell as Administrator, then restart), then open Ubuntu from the Start menu and run the curl command there:
wsl -- install
New to WSL? WSL is a Linux system with its own disk, so its home folder starts empty. Keep agents there ( ~/agents/<name> , the default), not in C:\ folders, which are much slower from WSL. Open their files from Windows at \\wsl.localhost\Ubuntu\home\<you>\agents . Windows in 5 minutes · Full Windows guide
Linux: Ensure build tools are installed: sudo apt install tmux build-essential
git clone https://github.com/23blocks-OS/ai-maestro.git
cd ai-maestro
yarn install
yarn dev
See QUICKSTART.md for detailed setup options.
Dashboard opens at http://localhost:23000
Update an existing install from its folder (pulls, updates submodules, builds, reinstalls hooks and the plugin, restarts):
./update-aimaestro.sh -y
Come back to it without remembering the address: the installer adds an ai-maestro command ( ai-maestro open starts the server if it is down, then opens the dashboard; also start , stop , status ) and, on a Mac, an AI Maestro app in ~/Applications you can drag to the Dock. It can also start at login. Chrome and Edge offer "Install app" from the dashboard too. Skip all of it with --skip-launcher ; remove it with ./scripts/install-launcher.sh --remove .
Every feature was born from running a real AI-first organization. We built them in the order we needed them.
I had 35 terminals and couldn't tell which was which.
See and manage all your AI agents in one place. Create agents from the UI with a guided wizard, organize them with smart naming ( project-backend-api becomes a 3-level tree with auto-coloring), and switch between any agent with a click. Four deployment modes: tmux (local), Docker (containerized), AWS EC2 (dedicated instance), and AWS ECS Fargate (serverless). Auto-discovers tmux sessions, Docker containers, cloud deployments, and standalone agents.
My Mac Mini was sitting there idle. What if I ran agents on that too?
A peer mesh network where every machine is equal. Add a computer, it joins the mesh. Every agent on every machine, visible from one dashboard. Use each machine for what it's best at — Mac for iOS builds, Linux for Docker, cloud for heavy compute. No central server required.
Worker machines can run headless ( yarn headless ) — the full API and agent runtime with no UI, in about 100MB of RAM. Run the dashboard where you sit; run agents wherever the compute is.
I was the mailman — copying messages between agents because they couldn't talk to each other.
The Agent Messaging Protocol (AMP) gives your agents email-like communication. Priority levels, message types, cryptographic signatures, and push notifications. Tell your agent "send a message to backend about the deployment" — it just works. Agents coordinate directly while you manage the big picture.
Before AMP: You copy research from one terminal, paste into another, repeat 50 times a day.
With AMP: "Research agent, send your findings to the writing agent." Done.
A friend in Singapore wanted his agents to talk to mine. But I didn't want to give him access to my network.
Connect your AI agents to Slack , Discord, Email, and WhatsApp through organizational gateways. Smart routing ( @AIM:agent-name ), thread-aware responses, and content security with 34 prompt injection patterns detected at the gateway — before any agent sees the message.
Every morning, my agents woke up with amnesia.
Long-term memory is a skill you switch on per agent. Every night each agent's conversations become memory: short statements of what it learned, each backed by the passages it came from, and an entity graph of the things it works with and how they relate ( runs on , depends on , stores data in , deploys to ), kept current, with ended relations marked. Claude Code deletes transcripts after 30 days; AI Maestro rebuilds that history from the agent's own message index, so months of work become memory instead of disappearing. When a prompt names something, the agent is told what it relates to before it acts, so it can see what a change affects. Knowledge that comes up in more sessions weighs more; secrets are redacted and never stored. How it works →
Alongside it: Code Graph (interactive visualization of your entire codebase with delta indexing) and Documentation (auto-generated, searchable docs from your code). Agents get smarter the longer they work with you.
I wanted to answer an agent's question without finding its terminal.
Talk to an agent in a chat view, on desktop or on your phone. When the agent asks a question or needs an approval, it shows up as a card you can click, and the click sends the same key you would press in the terminal. Claude Code, Codex and Grok

[truncated]
