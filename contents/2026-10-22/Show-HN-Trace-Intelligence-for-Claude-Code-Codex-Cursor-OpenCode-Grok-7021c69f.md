---
source: "https://github.com/viktor-com/tracequest"
hn_url: "https://news.ycombinator.com/item?id=50198407"
title: "Show HN: Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok"
article_title: "GitHub - viktor-com/tracequest: Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions. · GitHub"
image: "https://opengraph.githubassets.com/5792ef8e1d5b09ce169096403311361f028970fcb8e935a5ef251137c791cba1/viktor-com/tracequest"
author: "everlier"
captured_at: "2026-10-22T18:21:06Z"
capture_tool: "hn-digest"
hn_id: 50198407
score: 1
comments: 0
posted_at: "2026-10-22T17:26:39Z"
tags:
  - hacker-news
---

# Show HN: Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok

- HN: [50198407](https://news.ycombinator.com/item?id=50198407)
- Source: [github.com](https://github.com/viktor-com/tracequest)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T17:26:39Z

## Translation

Title: Show HN: Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok
Article title: GitHub - viktor-com/tracequest: Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions. · GitHub
Description: Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions. - viktor-com/tracequest

Article text:
GitHub - viktor-com/tracequest: Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions. · GitHub
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
viktor-com
tracequest Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions.
github.com/viktor-com/tracequest Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
76 Commits 76 Commits Folders and files
.github .github bench bench bin bin data data docs docs examples examples presets presets scripts scripts sidecar sidecar src src test test tests tests .facts .facts .gitignore .gitignore .nvmrc .nvmrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-NOTICES.md THIRD-PARTY-NOTICES.md TODO.md TODO.md chat-live.facts chat-live.facts chat-ui.facts chat-ui.facts cli-docs.facts cli-docs.facts command-palette.facts command-palette.facts cost-display.facts cost-display.facts cursor-cloud.facts cursor-cloud.facts fx-dsh.facts fx-dsh.facts hub-insights.facts hub-insights.facts keyboard.facts keyboard.facts launch.facts launch.facts package-lock.json package-lock.json package.json package.json run-status.facts run-status.facts session-loading.facts session-loading.facts ssh-import.facts ssh-import.facts unified-live.facts unified-live.facts usage-limits.facts usage-limits.facts vktr.facts vktr.facts View all files Repository files navigation
Follow your coding agents. Find what happened.
TraceQuest is a local web app and CLI for reading coding-agent sessions. It shows
each session's messages, tool calls, errors, timing and token usage. It reads
sessions from Claude Code, Codex CLI, Cursor, Droid (Factory), OpenCode, Grok CLI,
vktr (Viktor's fork of the Grok CLI), fx (Vercel's coding agent) and DeepSeek Harness.
It can also import Cursor Cloud Agent sessions and copy sessions from other
machines over SSH.
Try it · Features · Documentation
The screenshots use the bundled Claude Code samples and synthetic Codex recordings.
The hub shows copies under demo machine names.
Costs are estimated from token counts and model prices. They are not billing data.
The session list searches message text. You can filter by project, source,
machine, model, tool, error count and age. The summary row above the list shows
totals for the sessions that match.
How to use it: run npm start , open All sessions , type a filter such as
source:codex errors:>0 test and press Enter . To search from the terminal,
run node bin/tracequest.js search "cache invalidation" .
Search commands and filters .
The activity timeline shows when the agent was working and when it was idle. The
token chart shows input and output tokens for each turn, with markers for chapter
starts and tool errors. Each chapter lists the prompt, the commands the agent ran,
their output and any retries.
How to use it: open a session and choose Analytics , then Open full page . Click a bar or a
chapter to open it. node bin/tracequest.js latest --open opens a report for
your most recent session.
Rendering and export options .
Insights groups failed tool calls by cause. It also shows error rates, retries,
loops, active time and estimated cost. Expand a row to see the failed output and
a link to the session it came from. You can limit the figures to a project,
source, machine or time range. The analysis runs on your machine and makes no
model calls.
How to use it: open /insights on your TraceQuest server.
node bin/tracequest.js insights --refresh updates the analysis and prints a
summary in the terminal. Insights options .
The compare view puts two sessions side by side. It shows duration, tokens,
cache use, errors, retries, tool counts and chapter outcomes. Each column links
to the full session.
How to use it: in All sessions , tick the checkboxes of two sessions and
choose Compare . Browser and session commands .
Read, start and continue conversations
The conversation view shows a session's messages and tool output, and updates
while the agent is still working. + New run starts an installed agent in a
directory you choose. You can send follow-up messages to runs started from
TraceQuest and continue supported sessions. Starting and continuing runs needs
tmux.
How to use it: run npm start and choose a session or + New run .
Supported agents and setup .
Export reports and write handoffs
Export saves a session as a single HTML file with its timeline and chapters.
Markdown saves the session text. The handoff command searches your sessions
and prints Markdown snippets, each with the commands to open the matching session.
How to use it: open a session and choose Export or Markdown , or run
node bin/tracequest.js handoff "cache invalidation" --limit 5 in the checkout.
Handoff and sharing options .
Collect sessions from several machines
A hub is a read-only TraceQuest server that copies sessions from other machines
over SSH. The Machines table shows each machine's session count, newest session,
collection status and last successful copy. A config file lists the machines,
the folders to copy and the address the server listens on.
How to use it: follow the hub deployment guide to
write a config, import sessions and start the hub. The guide includes a local
example you can run on one machine. A hub rejects changes and hides values that
look like secrets. Sessions still contain conversation text, so restrict who can
reach the hub on your network.
You need Node.js 22.5 or later and git .
git clone https://github.com/viktor-com/tracequest.git
cd tracequest
npm install
npm run demo
Open localhost:7777 . The demo shows two small sample
sessions that ship with the repository. It needs no agent account or API key.
To see your own sessions, stop the demo with Ctrl-C and run:
npm start
Installation, imports and troubleshooting .
Browsing and analysis only read local files. TraceQuest sends no telemetry.
It contacts other services only when you ask it to: sharing a session, importing
remote sessions, or checking plan usage limits. Standup summaries and agent runs
use the agent you have configured.
The server listens on localhost and has no login. Read the
deployment guide and security notes
before making it reachable from other machines.
Documentation : command reference, configuration, session
locations, live-run API, deployment and development guides. It is written so
that a coding agent can follow it too.
Contribute · Report a bug · Changelog · MIT License
Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions.
github.com/viktor-com/tracequest Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions. - viktor-com/tracequest

GitHub - viktor-com/tracequest: Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions. · GitHub
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
viktor-com
tracequest Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions.
github.com/viktor-com/tracequest Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
76 Commits 76 Commits Folders and files
.github .github bench bench bin bin data data docs docs examples examples presets presets scripts scripts sidecar sidecar src src test test tests tests .facts .facts .gitignore .gitignore .nvmrc .nvmrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md THIRD-PARTY-NOTICES.md THIRD-PARTY-NOTICES.md TODO.md TODO.md chat-live.facts chat-live.facts chat-ui.facts chat-ui.facts cli-docs.facts cli-docs.facts command-palette.facts command-palette.facts cost-display.facts cost-display.facts cursor-cloud.facts cursor-cloud.facts fx-dsh.facts fx-dsh.facts hub-insights.facts hub-insights.facts keyboard.facts keyboard.facts launch.facts launch.facts package-lock.json package-lock.json package.json package.json run-status.facts run-status.facts session-loading.facts session-loading.facts ssh-import.facts ssh-import.facts unified-live.facts unified-live.facts usage-limits.facts usage-limits.facts vktr.facts vktr.facts View all files Repository files navigation
Follow your coding agents. Find what happened.
TraceQuest is a local web app and CLI for reading coding-agent sessions. It shows
each session's messages, tool calls, errors, timing and token usage. It reads
sessions from Claude Code, Codex CLI, Cursor, Droid (Factory), OpenCode, Grok CLI,
vktr (Viktor's fork of the Grok CLI), fx (Vercel's coding agent) and DeepSeek Harness.
It can also import Cursor Cloud Agent sessions and copy sessions from other
machines over SSH.
Try it · Features · Documentation
The screenshots use the bundled Claude Code samples and synthetic Codex recordings.
The hub shows copies under demo machine names.
Costs are estimated from token counts and model prices. They are not billing data.
The session list searches message text. You can filter by project, source,
machine, model, tool, error count and age. The summary row above the list shows
totals for the sessions that match.
How to use it: run npm start , open All sessions , type a filter such as
source:codex errors:>0 test and press Enter . To search from the terminal,
run node bin/tracequest.js search "cache invalidation" .
Search commands and filters .
The activity timeline shows when the agent was working and when it was idle. The
token chart shows input and output tokens for each turn, with markers for chapter
starts and tool errors. Each chapter lists the prompt, the commands the agent ran,
their output and any retries.
How to use it: open a session and choose Analytics , then Open full page . Click a bar or a
chapter to open it. node bin/tracequest.js latest --open opens a report for
your most recent session.
Rendering and export options .
Insights groups failed tool calls by cause. It also shows error rates, retries,
loops, active time and estimated cost. Expand a row to see the failed output and
a link to the session it came from. You can limit the figures to a project,
source, machine or time range. The analysis runs on your machine and makes no
model calls.
How to use it: open /insights on your TraceQuest server.
node bin/tracequest.js insights --refresh updates the analysis and prints a
summary in the terminal. Insights options .
The compare view puts two sessions side by side. It shows duration, tokens,
cache use, errors, retries, tool counts and chapter outcomes. Each column links
to the full session.
How to use it: in All sessions , tick the checkboxes of two sessions and
choose Compare . Browser and session commands .
Read, start and continue conversations
The conversation view shows a session's messages and tool output, and updates
while the agent is still working. + New run starts an installed agent in a
directory you choose. You can send follow-up messages to runs started from
TraceQuest and continue supported sessions. Starting and continuing runs needs
tmux.
How to use it: run npm start and choose a session or + New run .
Supported agents and setup .
Export reports and write handoffs
Export saves a session as a single HTML file with its timeline and chapters.
Markdown saves the session text. The handoff command searches your sessions
and prints Markdown snippets, each with the commands to open the matching session.
How to use it: open a session and choose Export or Markdown , or run
node bin/tracequest.js handoff "cache invalidation" --limit 5 in the checkout.
Handoff and sharing options .
Collect sessions from several machines
A hub is a read-only TraceQuest server that copies sessions from other machines
over SSH. The Machines table shows each machine's session count, newest session,
collection status and last successful copy. A config file lists the machines,
the folders to copy and the address the server listens on.
How to use it: follow the hub deployment guide to
write a config, import sessions and start the hub. The guide includes a local
example you can run on one machine. A hub rejects changes and hides values that
look like secrets. Sessions still contain conversation text, so restrict who can
reach the hub on your network.
You need Node.js 22.5 or later and git .
git clone https://github.com/viktor-com/tracequest.git
cd tracequest
npm install
npm run demo
Open localhost:7777 . The demo shows two small sample
sessions that ship with the repository. It needs no agent account or API key.
To see your own sessions, stop the demo with Ctrl-C and run:
npm start
Installation, imports and troubleshooting .
Browsing and analysis only read local files. TraceQuest sends no telemetry.
It contacts other services only when you ask it to: sharing a session, importing
remote sessions, or checking plan usage limits. Standup summaries and agent runs
use the agent you have configured.
The server listens on localhost and has no login. Read the
deployment guide and security notes
before making it reachable from other machines.
Documentation : command reference, configuration, session
locations, live-run API, deployment and development guides. It is written so
that a coding agent can follow it too.
Contribute · Report a bug · Changelog · MIT License
Trace Intelligence for Claude Code, Codex, Cursor, OpenCode, Grok and Droid sessions.
github.com/viktor-com/tracequest Topics
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
