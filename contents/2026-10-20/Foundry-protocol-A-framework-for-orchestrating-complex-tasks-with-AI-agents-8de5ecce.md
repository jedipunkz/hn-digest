---
source: "https://github.com/AlphaIAt/foundry-protocol"
hn_url: "https://news.ycombinator.com/item?id=50172883"
title: "Foundry-protocol – A framework for orchestrating complex tasks with AI agents"
article_title: "GitHub - AlphaIAt/foundry-protocol: A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory. · GitHub"
image: "https://opengraph.githubassets.com/8676897d99f334097cfd8c0437c1e28f4016bd79232e95623279972fc9951438/AlphaIAt/foundry-protocol"
author: "gbasilveira"
captured_at: "2026-10-20T22:11:30Z"
capture_tool: "hn-digest"
hn_id: 50172883
score: 1
comments: 0
posted_at: "2026-10-20T21:23:48Z"
tags:
  - hacker-news
---

# Foundry-protocol – A framework for orchestrating complex tasks with AI agents

- HN: [50172883](https://news.ycombinator.com/item?id=50172883)
- Source: [github.com](https://github.com/AlphaIAt/foundry-protocol)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T21:23:48Z

## Translation

Title: Foundry-protocol – A framework for orchestrating complex tasks with AI agents
Article title: GitHub - AlphaIAt/foundry-protocol: A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory. · GitHub
Description: A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory. - AlphaIAt/foundry-protocol

Article text:
GitHub - AlphaIAt/foundry-protocol: A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory. · GitHub
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
AlphaIAt
foundry-protocol Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory.
Readme License Activity Custom properties Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
.claude-plugin .claude-plugin docs docs plugin plugin runtime runtime scripts scripts .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
A reusable framework for orchestrating complex tasks with AI agents, using the
human/AI-orchestrator protocol + Claude planning + sub-agent executors. Extracted and
generalized from real experience in a compliance-audit project (T0-T5, 25 work orders,
2026-09).
Core idea: the protocol itself is an artifact that improves with use. Each project
that uses it produces lessons — what worked, what misled an agent, which acceptance
criterion was missing. Those lessons flow back to docs/LESSONS.md in
this repository, rather than staying trapped in the project where they arose. The next
project starts with that experience already built in, instead of repeating the same
mistakes.
foundry-protocol/
├── CLAUDE.md # operating instructions for Claude
│ # when acting as orchestrator/planner
├── docs/
│ ├── ORCHESTRATION.md # generic protocol (roles, cycle, WO, log)
│ ├── DECISIONS.md # template for the decisions/ADR register
│ ├── AUTONOMY-EVAL.md # template + validated criteria for
│ │ # assessing whether an agent can coordinate alone
│ ├── LESSONS.md # self-improvement ledger — grows with each project
│ ├── history/ # where condensed per-project summaries live
│ ├── memory/ # template for per-project living memory
│ │ ├── ENTRY-TEMPLATE.md # format of one entry (same shape as the
│ │ │ # Claude Code personal memory)
│ │ └── MEMORY.md # index template
│ ├── work-orders/TEMPLATE.md # work order template (PREPARATION, tasks, sabotage checks, RESULT)
│ ├── gates/TEMPLATE.md # human validation gate template (≥ 1 PR, complete report)
│ ├── ORCHESTRATOR-LOOP.md # optional autonomous cycle (orchestrator does not implement; failure ⇒ new agent)
│ ├── FORMAT-CONTRACT.md # stable shapes that tools read without an LLM
│ ├── RESUME.md # resuming a WO in any session/account
│ ├── RESTART-PROMPT.md # template for the orchestrator restart prompt
│ ├── DR-INDEX.md # DR index, one line each (minimal context)
│ ├── CONTEXT-TEMPLATE.md # per-package CONTEXT.md (≤ 45 lines, instead of reading src/)
│ └── foundry.conf # BASE_BRANCH, GATE_BRANCH_PREFIX, VERIFY_CMD
├── .claude-plugin/marketplace.json # local marketplace (for `claude plugin install`)
├── plugin/ # Claude Code plugin: skills + MCP engine
│ ├── .claude-plugin/plugin.json
│ ├── mcp/server.mjs # local MCP engine (see plugin/mcp/README.md)
│ └── skills/
│ ├── bootstrap-project/
│ ├── close-work-order-cycle/
│ └── write-project-memory/
├── runtime/ # agno/AgentOS foundation (Phase 1, experimental)
└── scripts/
├── bootstrap-project.sh # installs the skeleton into a new or existing project
└── project/ # scripts copied to <project>/scripts/
├── wo.sh # status [NNN] | resume NNN | next
├── wo-watchdog.sh # STA
[truncated]
How to use it on a new project
/path/to/foundry-protocol/scripts/bootstrap-project.sh /path/to/project " Project Name "
This creates, inside the target project, the documents of docs/ listed above (protocol,
WO and gate templates, foundry.conf , LOG, history/ ), the scripts from
scripts/project/ into scripts/ , the logging hook in .claude/settings.json (only if it
does not exist yet) and docs/memory/{ENTRY-TEMPLATE.md,MEMORY.md,entries/,media/} . The
latter is the project's living memory, which grows at every cycle (see §6 of the generated
ORCHESTRATION.md ). After that, the normal work is: fill in §0/§3 (state + shared context)
of the project's ORCHESTRATION.md with that project's concrete facts, consult
docs/LESSONS.md before the first work order, and follow the protocol
in §2.
/path/to/foundry-protocol/scripts/update-project.sh /path/to/project # dry run: shows the plan
/path/to/foundry-protocol/scripts/update-project.sh /path/to/project --apply # applies it
Bootstrap records, for each managed file ( scripts/managed-files.txt ), a baseline copy in
docs/.foundry/base/ and the foundry version in docs/.foundry/version . On update:
UPDATE — the project did not edit the file: it is replaced by the new version.
MERGE — both changed: a three-way merge ( git merge-file ), applied only if it is clean.
CONFLICT — overlapping edits: the project file is left untouched and the merge is written
to <file>.conflict . Resolve by hand, delete the .conflict file, and run --apply again.
ADOPT / NEW / DELETED / GONE — projects bootstrapped before this mechanism have no
baseline: identical files are adopted, different ones are reported as conflicts to review by hand.
The project name, base branch and relative link to foundry-protocol are rendered the same way the
bootstrap fills them in, so they are not reverted.
plugin/ packages the recurring procedures of the protocol as
Claude Skills —
bootstrap-project , close-work-order-cycle , write-project-memory — so that any session
can invoke them directly instead of re-reading all of ORCHESTRATION.md each time.
To register the plugin locally, symlink plugin/ into Claude Code's personal skills directory
( <claude-skills-dir>/ ). It then loads in any session, in any project:
ln -s /path/to/foundry-protocol/plugin < claude-skills-dir > /foundry-protocol
claude plugin validate < claude-skills-dir > /foundry-protocol # confirm
claude plugin list # should list foundry-protocol@skills-dir
Because the symlink points into the repository, editing plugin/ here takes effect
immediately — there is no sync step.
Two directions, not to be confused:
Project memory (at every relevant cycle): the Claude in planning role writes to
docs/memory/entries/ of that project any project-specific fact/decision worth
remembering — not only when closing a phase. The project's own
docs/ORCHESTRATION.md / DECISIONS.md continue to hold protocol and broad history, as before.
Generalizable lessons (rarer, typically when closing a phase): what would serve
any future project — a missing acceptance criterion, a prompt shape that confused an
agent — is distilled and written back to foundry-protocol/docs/LESSONS.md ,
never the reverse. This repository's generic ORCHESTRATION.md is only edited when a
lesson implies changing the protocol itself. See §6 of ORCHESTRATION.md and §3 of CLAUDE.md .
This is what makes the agent "more competent with experience": it is not a model being
retrained; it is the protocol + the lessons ledger accumulating explicit, readable,
versioned judgment — applied by any future session that reads this repository before starting.
Original authors and maintainers: αIAt ( https://aiat.pt ) and Guilherme da Silveira
( gbasilveira@gmail.com ).
Licensed under MIT — see LICENSE . The code files (scripts, MCP engine, runtime)
carry the same attribution in their header.
runtime/ (agno/AgentOS foundation, Phase 1) is experimental and is not yet supported as
a stable part of the protocol.
A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory.
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory. - AlphaIAt/foundry-protocol

GitHub - AlphaIAt/foundry-protocol: A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory. · GitHub
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
AlphaIAt
foundry-protocol Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory.
Readme License Activity Custom properties Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
.claude-plugin .claude-plugin docs docs plugin plugin runtime runtime scripts scripts .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
A reusable framework for orchestrating complex tasks with AI agents, using the
human/AI-orchestrator protocol + Claude planning + sub-agent executors. Extracted and
generalized from real experience in a compliance-audit project (T0-T5, 25 work orders,
2026-09).
Core idea: the protocol itself is an artifact that improves with use. Each project
that uses it produces lessons — what worked, what misled an agent, which acceptance
criterion was missing. Those lessons flow back to docs/LESSONS.md in
this repository, rather than staying trapped in the project where they arose. The next
project starts with that experience already built in, instead of repeating the same
mistakes.
foundry-protocol/
├── CLAUDE.md # operating instructions for Claude
│ # when acting as orchestrator/planner
├── docs/
│ ├── ORCHESTRATION.md # generic protocol (roles, cycle, WO, log)
│ ├── DECISIONS.md # template for the decisions/ADR register
│ ├── AUTONOMY-EVAL.md # template + validated criteria for
│ │ # assessing whether an agent can coordinate alone
│ ├── LESSONS.md # self-improvement ledger — grows with each project
│ ├── history/ # where condensed per-project summaries live
│ ├── memory/ # template for per-project living memory
│ │ ├── ENTRY-TEMPLATE.md # format of one entry (same shape as the
│ │ │ # Claude Code personal memory)
│ │ └── MEMORY.md # index template
│ ├── work-orders/TEMPLATE.md # work order template (PREPARATION, tasks, sabotage checks, RESULT)
│ ├── gates/TEMPLATE.md # human validation gate template (≥ 1 PR, complete report)
│ ├── ORCHESTRATOR-LOOP.md # optional autonomous cycle (orchestrator does not implement; failure ⇒ new agent)
│ ├── FORMAT-CONTRACT.md # stable shapes that tools read without an LLM
│ ├── RESUME.md # resuming a WO in any session/account
│ ├── RESTART-PROMPT.md # template for the orchestrator restart prompt
│ ├── DR-INDEX.md # DR index, one line each (minimal context)
│ ├── CONTEXT-TEMPLATE.md # per-package CONTEXT.md (≤ 45 lines, instead of reading src/)
│ └── foundry.conf # BASE_BRANCH, GATE_BRANCH_PREFIX, VERIFY_CMD
├── .claude-plugin/marketplace.json # local marketplace (for `claude plugin install`)
├── plugin/ # Claude Code plugin: skills + MCP engine
│ ├── .claude-plugin/plugin.json
│ ├── mcp/server.mjs # local MCP engine (see plugin/mcp/README.md)
│ └── skills/
│ ├── bootstrap-project/
│ ├── close-work-order-cycle/
│ └── write-project-memory/
├── runtime/ # agno/AgentOS foundation (Phase 1, experimental)
└── scripts/
├── bootstrap-project.sh # installs the skeleton into a new or existing project
└── project/ # scripts copied to <project>/scripts/
├── wo.sh # status [NNN] | resume NNN | next
├── wo-watchdog.sh # STA
[truncated]
How to use it on a new project
/path/to/foundry-protocol/scripts/bootstrap-project.sh /path/to/project " Project Name "
This creates, inside the target project, the documents of docs/ listed above (protocol,
WO and gate templates, foundry.conf , LOG, history/ ), the scripts from
scripts/project/ into scripts/ , the logging hook in .claude/settings.json (only if it
does not exist yet) and docs/memory/{ENTRY-TEMPLATE.md,MEMORY.md,entries/,media/} . The
latter is the project's living memory, which grows at every cycle (see §6 of the generated
ORCHESTRATION.md ). After that, the normal work is: fill in §0/§3 (state + shared context)
of the project's ORCHESTRATION.md with that project's concrete facts, consult
docs/LESSONS.md before the first work order, and follow the protocol
in §2.
/path/to/foundry-protocol/scripts/update-project.sh /path/to/project # dry run: shows the plan
/path/to/foundry-protocol/scripts/update-project.sh /path/to/project --apply # applies it
Bootstrap records, for each managed file ( scripts/managed-files.txt ), a baseline copy in
docs/.foundry/base/ and the foundry version in docs/.foundry/version . On update:
UPDATE — the project did not edit the file: it is replaced by the new version.
MERGE — both changed: a three-way merge ( git merge-file ), applied only if it is clean.
CONFLICT — overlapping edits: the project file is left untouched and the merge is written
to <file>.conflict . Resolve by hand, delete the .conflict file, and run --apply again.
ADOPT / NEW / DELETED / GONE — projects bootstrapped before this mechanism have no
baseline: identical files are adopted, different ones are reported as conflicts to review by hand.
The project name, base branch and relative link to foundry-protocol are rendered the same way the
bootstrap fills them in, so they are not reverted.
plugin/ packages the recurring procedures of the protocol as
Claude Skills —
bootstrap-project , close-work-order-cycle , write-project-memory — so that any session
can invoke them directly instead of re-reading all of ORCHESTRATION.md each time.
To register the plugin locally, symlink plugin/ into Claude Code's personal skills directory
( <claude-skills-dir>/ ). It then loads in any session, in any project:
ln -s /path/to/foundry-protocol/plugin < claude-skills-dir > /foundry-protocol
claude plugin validate < claude-skills-dir > /foundry-protocol # confirm
claude plugin list # should list foundry-protocol@skills-dir
Because the symlink points into the repository, editing plugin/ here takes effect
immediately — there is no sync step.
Two directions, not to be confused:
Project memory (at every relevant cycle): the Claude in planning role writes to
docs/memory/entries/ of that project any project-specific fact/decision worth
remembering — not only when closing a phase. The project's own
docs/ORCHESTRATION.md / DECISIONS.md continue to hold protocol and broad history, as before.
Generalizable lessons (rarer, typically when closing a phase): what would serve
any future project — a missing acceptance criterion, a prompt shape that confused an
agent — is distilled and written back to foundry-protocol/docs/LESSONS.md ,
never the reverse. This repository's generic ORCHESTRATION.md is only edited when a
lesson implies changing the protocol itself. See §6 of ORCHESTRATION.md and §3 of CLAUDE.md .
This is what makes the agent "more competent with experience": it is not a model being
retrained; it is the protocol + the lessons ledger accumulating explicit, readable,
versioned judgment — applied by any future session that reads this repository before starting.
Original authors and maintainers: αIAt ( https://aiat.pt ) and Guilherme da Silveira
( gbasilveira@gmail.com ).
Licensed under MIT — see LICENSE . The code files (scripts, MCP engine, runtime)
carry the same attribution in their header.
runtime/ (agno/AgentOS foundation, Phase 1) is experimental and is not yet supported as
a stable part of the protocol.
A protocol and toolkit for coordinating AI agents through work orders, with review by command, human validation gates, resumable state, and a per-project memory.
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
