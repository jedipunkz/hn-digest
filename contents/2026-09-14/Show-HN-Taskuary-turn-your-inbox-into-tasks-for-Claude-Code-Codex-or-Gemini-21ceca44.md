---
source: "https://github.com/ldbumble/taskuary"
hn_url: "https://news.ycombinator.com/item?id=49701075"
title: "Show HN: Taskuary – turn your inbox into tasks for Claude Code, Codex, or Gemini"
article_title: "GitHub - ldbumble/taskuary: Automate your job: local-first AI task hub. Email, Teams, Slack & reports -> one timeline -> AI triage -> your coding agents (Claude Code, Codex, Gemini) do the work, you approve. · GitHub"
image: "https://opengraph.githubassets.com/b9aac1a5ccd69c9361cb5de426af35a3215ba9ef252a9af52a72e61af02363c2/ldbumble/taskuary"
author: "monkeydo09"
captured_at: "2026-09-14T18:02:11Z"
capture_tool: "hn-digest"
hn_id: 49701075
score: 2
comments: 1
posted_at: "2026-09-14T17:55:23Z"
tags:
  - hacker-news
---

# Show HN: Taskuary – turn your inbox into tasks for Claude Code, Codex, or Gemini

- HN: [49701075](https://news.ycombinator.com/item?id=49701075)
- Source: [github.com](https://github.com/ldbumble/taskuary)
- Score: 2
- Comments: 1
- Posted: 2026-09-14T17:55:23Z

## Translation

Title: Show HN: Taskuary – turn your inbox into tasks for Claude Code, Codex, or Gemini
Article title: GitHub - ldbumble/taskuary: Automate your job: local-first AI task hub. Email, Teams, Slack & reports -> one timeline -> AI triage -> your coding agents (Claude Code, Codex, Gemini) do the work, you approve. · GitHub
Description: Automate your job: local-first AI task hub. Email, Teams, Slack & reports -> one timeline -> AI triage -> your coding agents (Claude Code, Codex, Gemini) do the work, you approve. - ldbumble/taskuary

Article text:
GitHub - ldbumble/taskuary: Automate your job: local-first AI task hub. Email, Teams, Slack & reports -> one timeline -> AI triage -> your coding agents (Claude Code, Codex, Gemini) do the work, you approve. · GitHub
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
ldbumble
/
taskuary
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
1,075 Commits 1,075 Commits Folders and files
.claude/ skills/ deploy .claude/ skills/ deploy .github .github assets assets docs docs functions functions site site taskuary taskuary tests tests tools tools website website .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml taskuary.spec taskuary.spec worker.mjs worker.mjs wrangler.jsonc wrangler.jsonc View all files Repository files navigation
Your inbox, staffed by AI agents
Taskuary turns incoming messages into organized work. It sorts what matters, hands tasks to
your agents, and brings decisions back to you. Nothing sends or ships without your approval.
Taskuary is early—currently v0.3.4.6 —so breaking changes are still possible before 1.0.
The real app with invented data. Nothing connects, sends, or runs.
Connect every system. Keep control.
Bring your mail, chats, issue trackers, alerts, reports, and AI agents into one control center
that runs on your machine. Every input lands on a single Timeline with its context, status, and
available actions intact. Taskuary gives your tools one place to work together while you decide
what can run, what can send, and what needs your attention.
Turn incoming work into tasks automatically
AI triage reads every incoming item, understands what it means, and turns actionable requests
into ready-to-run tasks. Routine updates are filed, questions get draft replies, and urgent
decisions move to the front. Every choice comes with a reason, and you can change it at any time.
Watch the work. Approve the outcome.
Taskuary keeps the system moving and brings you in when your judgment matters. The Assistant
speaks up when a reply is waiting, a task has gone quiet, a report has failed, or a meeting is
about to start. Each update arrives once, with the evidence and next actions attached, so you can
review the work, make the call, and approve what happens next.
Send coding work to Claude Code, Codex, Gemini, Cursor, Copilot, Muse Code, or another CLI.
Configure each CLI once under Connections → AI CLI agents : command, arguments, timeout,
installation, sign-in and connection tests. Profiles share that connection.
Connection cards show installation status separately from configuration. If Windows blocks an
installer before it starts, check Windows Security's Protection history for the reason.
Install and Update open a terminal beneath the CLI cards, showing commands and live output.
You can answer the vendor's prompts there; closing the terminal stops a running installer.
Under Docs → Profiles , use Add profile to choose a provider and model, set the worker's
instructions, and describe when triage should select it. Leave
Available to triage for new tasks on to include it automatically. Each profile starts with
editable instructions; coding workers such as Coder and Codex share CODER.md .
Watch the terminal, answer questions, review changes, and use the built-in browser without
losing the session.
Describe a check. Let it run itself.
Say what you want watched in plain words and Taskuary builds it: read a folder of CSVs every
morning, total the AP bills due in the next 30 days, count helpdesk tickets by day. Reports read
and summarize, on a schedule or on demand, and can come back with an AI summary. Workflows write
data and keep state. Quiet checks stay quiet, so the only one you hear about is the one that
failed.
Replies and proposed actions wait in Review. You decide what sends, runs, closes, or gets dismissed.
1. Work arrives and gets sorted
Taskuary reads connected inboxes, separates tasks from noise, and shows the result on one Timeline.
Tasks move from Queued to Agent working . The Board shows live progress and anything waiting on you.
3. The result comes back to you
Finished work, draft replies, questions, and loose ends return for review. You decide what to send, change, snooze, or dismiss.
Download the latest single-file
Taskuary.exe
and open it. No Python or installer is required.
Python 3.10 or newer works on Windows, macOS, and Linux:
pip install taskuary
taskuary
Taskuary opens at http://127.0.0.1:7787 . For a native desktop
window instead, install pip install "taskuary[desktop]" and run taskuary-desktop .
git clone https://github.com/ldbumble/taskuary
cd taskuary
docker compose up
Then open http://127.0.0.1:7787 . Docker runs the web app;
coding CLIs and the optional WhatsApp bridge remain on the host.
On first run, connect an AI provider or local Ollama model, add at least one inbound
channel, then choose the coding CLI that should receive tasks. The setup wizards test each
connection before it goes live.
Try it without installing anything
taskuary --demo # or: docker compose --profile demo up
The demo is the real interface with fictional work and scripted replies. It cannot connect to
outside systems, send messages, run tools, or start agents. Its changes reset when you reload.
Updated daily from PyPI with mirror traffic excluded. The raw series is
docs/downloads.csv .
Getting started —installation, first-run setup, Docker, and data
Product guide —the workflow, learning loop, agents, and operator documents
Integrations —channels, AI providers, work systems, and report sources
Reports and proactive checks —the report pipeline, AI-written source cards, and what Taskuary watches
Status and roadmap —what works today and what is next
Contributing —development setup and contribution guide
Taskuary is free and open source under the MIT License . Issues and pull requests
are welcome; security reports belong in SECURITY.md .
Automate your job: local-first AI task hub. Email, Teams, Slack & reports -> one timeline -> AI triage -> your coding agents (Claude Code, Codex, Gemini) do the work, you approve.
Readme MIT license Code of conduct
Security policy Activity Stars
12 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Automate your job: local-first AI task hub. Email, Teams, Slack & reports -> one timeline -> AI triage -> your coding agents (Claude Code, Codex, Gemini) do the work, you approve. - ldbumble/taskuary

GitHub - ldbumble/taskuary: Automate your job: local-first AI task hub. Email, Teams, Slack & reports -> one timeline -> AI triage -> your coding agents (Claude Code, Codex, Gemini) do the work, you approve. · GitHub
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
ldbumble
/
taskuary
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
1,075 Commits 1,075 Commits Folders and files
.claude/ skills/ deploy .claude/ skills/ deploy .github .github assets assets docs docs functions functions site site taskuary taskuary tests tests tools tools website website .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml taskuary.spec taskuary.spec worker.mjs worker.mjs wrangler.jsonc wrangler.jsonc View all files Repository files navigation
Your inbox, staffed by AI agents
Taskuary turns incoming messages into organized work. It sorts what matters, hands tasks to
your agents, and brings decisions back to you. Nothing sends or ships without your approval.
Taskuary is early—currently v0.3.4.6 —so breaking changes are still possible before 1.0.
The real app with invented data. Nothing connects, sends, or runs.
Connect every system. Keep control.
Bring your mail, chats, issue trackers, alerts, reports, and AI agents into one control center
that runs on your machine. Every input lands on a single Timeline with its context, status, and
available actions intact. Taskuary gives your tools one place to work together while you decide
what can run, what can send, and what needs your attention.
Turn incoming work into tasks automatically
AI triage reads every incoming item, understands what it means, and turns actionable requests
into ready-to-run tasks. Routine updates are filed, questions get draft replies, and urgent
decisions move to the front. Every choice comes with a reason, and you can change it at any time.
Watch the work. Approve the outcome.
Taskuary keeps the system moving and brings you in when your judgment matters. The Assistant
speaks up when a reply is waiting, a task has gone quiet, a report has failed, or a meeting is
about to start. Each update arrives once, with the evidence and next actions attached, so you can
review the work, make the call, and approve what happens next.
Send coding work to Claude Code, Codex, Gemini, Cursor, Copilot, Muse Code, or another CLI.
Configure each CLI once under Connections → AI CLI agents : command, arguments, timeout,
installation, sign-in and connection tests. Profiles share that connection.
Connection cards show installation status separately from configuration. If Windows blocks an
installer before it starts, check Windows Security's Protection history for the reason.
Install and Update open a terminal beneath the CLI cards, showing commands and live output.
You can answer the vendor's prompts there; closing the terminal stops a running installer.
Under Docs → Profiles , use Add profile to choose a provider and model, set the worker's
instructions, and describe when triage should select it. Leave
Available to triage for new tasks on to include it automatically. Each profile starts with
editable instructions; coding workers such as Coder and Codex share CODER.md .
Watch the terminal, answer questions, review changes, and use the built-in browser without
losing the session.
Describe a check. Let it run itself.
Say what you want watched in plain words and Taskuary builds it: read a folder of CSVs every
morning, total the AP bills due in the next 30 days, count helpdesk tickets by day. Reports read
and summarize, on a schedule or on demand, and can come back with an AI summary. Workflows write
data and keep state. Quiet checks stay quiet, so the only one you hear about is the one that
failed.
Replies and proposed actions wait in Review. You decide what sends, runs, closes, or gets dismissed.
1. Work arrives and gets sorted
Taskuary reads connected inboxes, separates tasks from noise, and shows the result on one Timeline.
Tasks move from Queued to Agent working . The Board shows live progress and anything waiting on you.
3. The result comes back to you
Finished work, draft replies, questions, and loose ends return for review. You decide what to send, change, snooze, or dismiss.
Download the latest single-file
Taskuary.exe
and open it. No Python or installer is required.
Python 3.10 or newer works on Windows, macOS, and Linux:
pip install taskuary
taskuary
Taskuary opens at http://127.0.0.1:7787 . For a native desktop
window instead, install pip install "taskuary[desktop]" and run taskuary-desktop .
git clone https://github.com/ldbumble/taskuary
cd taskuary
docker compose up
Then open http://127.0.0.1:7787 . Docker runs the web app;
coding CLIs and the optional WhatsApp bridge remain on the host.
On first run, connect an AI provider or local Ollama model, add at least one inbound
channel, then choose the coding CLI that should receive tasks. The setup wizards test each
connection before it goes live.
Try it without installing anything
taskuary --demo # or: docker compose --profile demo up
The demo is the real interface with fictional work and scripted replies. It cannot connect to
outside systems, send messages, run tools, or start agents. Its changes reset when you reload.
Updated daily from PyPI with mirror traffic excluded. The raw series is
docs/downloads.csv .
Getting started —installation, first-run setup, Docker, and data
Product guide —the workflow, learning loop, agents, and operator documents
Integrations —channels, AI providers, work systems, and report sources
Reports and proactive checks —the report pipeline, AI-written source cards, and what Taskuary watches
Status and roadmap —what works today and what is next
Contributing —development setup and contribution guide
Taskuary is free and open source under the MIT License . Issues and pull requests
are welcome; security reports belong in SECURITY.md .
Automate your job: local-first AI task hub. Email, Teams, Slack & reports -> one timeline -> AI triage -> your coding agents (Claude Code, Codex, Gemini) do the work, you approve.
Readme MIT license Code of conduct
Security policy Activity Stars
12 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
