---
source: "https://github.com/ringlochid/oh-my-subagents"
hn_url: "https://news.ycombinator.com/item?id=49796249"
title: "Turn ad-hoc subagents into durable, accountable AI teams"
article_title: "GitHub - ringlochid/oh-my-subagents: Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery. · GitHub"
image: "https://opengraph.githubassets.com/2763d46f9fa894a9e825c220b273ef1cd3b705c173dfe76b5f955d1f1bab27d7/ringlochid/oh-my-subagents"
author: "ringlochid"
captured_at: "2026-09-22T03:37:13Z"
capture_tool: "hn-digest"
hn_id: 49796249
score: 4
comments: 0
posted_at: "2026-09-22T02:45:47Z"
tags:
  - hacker-news
---

# Turn ad-hoc subagents into durable, accountable AI teams

- HN: [49796249](https://news.ycombinator.com/item?id=49796249)
- Source: [github.com](https://github.com/ringlochid/oh-my-subagents)
- Score: 4
- Comments: 0
- Posted: 2026-09-22T02:45:47Z

## Translation

Title: Turn ad-hoc subagents into durable, accountable AI teams
Article title: GitHub - ringlochid/oh-my-subagents: Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery. · GitHub
Description: Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery. - ringlochid/oh-my-subagents

Article text:
GitHub - ringlochid/oh-my-subagents: Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery. · GitHub
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
ringlochid
/
oh-my-subagents
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
144 Commits 144 Commits Folders and files
.agents/ standards .agents/ standards .github .github console console docs-internal docs-internal docs docs examples examples infra infra openapi openapi scripts scripts src src tests tests .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md STYLE.md STYLE.md docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json View all files Repository files navigation
Turn ad-hoc subagents into durable, accountable AI teams.
Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery.
Get started · Documentation · Starter teams
Start with the Codex team guide , Claude team guide , or a concrete independent code-review workflow .
Ad-hoc delegation is easy to begin and surprisingly hard to operate. A parent spawns children, polls them, reconstructs ownership from chat, and hopes that a closed terminal or interrupted provider session did not erase the only useful account of what happened.
Oh My Subagents moves that coordination into durable runtime state:
When a Manager delegates a Wave, Oh My Subagents (OMS) commits the child Assignments, persists the parent wait, supervises every return, and continues the parent with the complete Checkpoints—even after an interruption.
The parent does not remain in a polling loop. A delegated Wave and its child Assignments commit together with the parent wait. Children return terminal Checkpoints independently; the controller collects the complete Wave and continues the parent with every return.
Managers can delegate recursively. Every Manager follows the same local rule with its direct children: delegate, wait, inspect the complete returns, and integrate. Deeper teams need no hidden global polling loop.
Here is one concrete run: evidence determines the next Wave, review findings become a new repair Assignment, and verification remains independent.
🧠 Reuse responsibility—not a frozen schedule
A Workflow definition answers who is responsible . It does not prescribe when a Member runs or force the work into a DAG.
At runtime, Managers choose the pattern that fits the actual Task and current evidence:
run dependent work sequentially;
fan independent work out in parallel;
iterate through implement, review, and repair;
divide a bounded batch among reusable owners;
combine those patterns in one Task; or
replan one responsibility subtree when the current team no longer fits.
The same published team can therefore respond differently to two different missions. A replan changes future responsibility without rewriting earlier revisions, completed work, or accepted history.
🧾 Accountability is a runtime contract
OMS does more than display several agents at once:
One stable starting contract. Every Task pins the exact published Workflow revision it started with.
One owner per Assignment. Child work is immutable, task-specific, and tied to the Member responsible for returning it.
Durable fan-out and fan-in. Waves, waits, retries, replans, Checkpoints, and continuations live in controller-owned state—not only in a provider transcript.
One accountable Result. Child completion is evidence for a parent. Only the Task lead's accepted completed or blocked Checkpoint becomes the Result shown to you.
Honest recovery. Browser closure, provider interruption, and controller restart do not silently fabricate completion or discard accepted history.
Ordinary files stay ordinary. Notes, reports, code, and artifacts remain in your workspace. OMS records small navigation references instead of pretending to own or snapshot every byte.
Capabilities deny by default and never inherit from a parent. Provider, model, sandbox, Human Request, and managed Command Run choices stay explicit where a team needs them without becoming mandatory ceremony for every Workflow.
🎛️ Design and operate visually
The visual Console is the primary experience:
Workflow library keeps reusable teams, drafts, and published revisions together.
Workflow Studio lets you shape the complete responsibility hierarchy on a horizontal canvas, validate it, and publish deliberately.
Run Studio shows the live team, current plans, meaningful Activity, Human Requests, managed Actions, referenced files, and the exact completed or blocked Result.
Steering delivers bounded new context to one exact active Member without pretending earlier work or tool effects never happened.
Prefer conversation? The separate Operator can draft and revise Workflows, explain teams, publish when asked, start and control Runs, answer Human Requests, and inspect managed Actions. Operator and the visual interface call the same controller-owned operations, so chat never creates a second hidden copy of product truth.
Oh My Subagents requires Python 3.12 or newer and supports Linux, macOS 13+, and Windows 11 x64. Install it in an isolated environment with pipx :
pipx install oh-my-subagents
oms init
oms service install
Open http://127.0.0.1:18125/ .
Guided initialization selects a default workspace, configures a Codex or Claude Task provider, publishes the Starter teams, and can configure the separate Operator. SQLite is the default, so a local installation needs no database server.
Upgrade an existing Banksia installation
Migrate before running oms init , oms serve , or oms service install . Initialization creates a fresh controller database; migration preserves the existing Banksia config, database, provider credentials, Task history, and service state under canonical OMS paths.
banksia status --json
banksia service status --json
pipx uninstall banksia
pipx install oh-my-subagents
oms migrate-from-banksia
oms db upgrade
oms service status --json
If the old installation did not use a background service, run oms service install after migration. Do not run oms init first, and do not use oms db reset as a rename step. The Banksia migration guide covers custom paths, verification, rollback, and recovery from an accidental 0.3.0 initialization.
oms service install verifies the configuration and database schema, installs a native per-user background service, and starts it. The controller keeps supervising work after you close the terminal and returns at login.
oms service status
oms service restart
oms service logs --lines 200
oms service stop
Linux uses a systemd user service, macOS a current-user LaunchAgent, and Windows a current-user Scheduled Task. oms service uninstall removes the native service definition while preserving configuration, database, and provider credentials. Use oms serve when you prefer the portable foreground path.
Install the optional driver and provide a SQLAlchemy URL during initialization:
pipx install " oh-my-subagents[postgres] "
oms init \
--database-url " postgresql+asyncpg://oms@127.0.0.1/oms "
See Getting started for provider prerequisites and a complete first run, or Database configuration for PostgreSQL permissions, schemas, and environment overrides.
oms init publishes eight provider-neutral Starters. Choose one when a mission is consequential or broad enough that independent ownership, durable work, or adversarial verification adds real value:
The Starter catalog includes example missions, expected deliverables, and guidance on when a simpler team is better. A one-Member Workflow is valid—and usually wiser—when delegation would add ceremony without independent evidence, useful specialization, or real integration.
Choose or design a team. Start from an included Starter Workflow or shape a responsibility tree in Workflow Studio.
Publish a stable revision. Every Run keeps the exact team contract it started with.
Give the lead one complete prompt. Managers plan, delegate, adapt, and integrate while OMS supervises runtime state.
Stay involved where judgment matters. Steer an active Member or answer a typed Human Request when the team genuinely needs you.
Receive the accountable Result. Read the lead's exact completed or blocked outcome and follow its references into detailed workspace files.
Understand Workflows and teams
Oh My Subagents is open source under the MIT License , except for the visual Console in console/ , which contains material derived from n8n and is distributed under the Sustainable Use License . That license permits internal business, non-commercial, and personal use; redistribution is limited to free, non-commercial distribution. See the Console notice for attribution and the modification notice.
Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery.
ringlochid.me/oh-my-subagents/ Topics
Readme MIT license Contributing
7 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery. - ringlochid/oh-my-subagents

GitHub - ringlochid/oh-my-subagents: Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery. · GitHub
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
ringlochid
/
oh-my-subagents
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
144 Commits 144 Commits Folders and files
.agents/ standards .agents/ standards .github .github console console docs-internal docs-internal docs docs examples examples infra infra openapi openapi scripts scripts src src tests tests .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md STYLE.md STYLE.md docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml pyrightconfig.json pyrightconfig.json View all files Repository files navigation
Turn ad-hoc subagents into durable, accountable AI teams.
Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery.
Get started · Documentation · Starter teams
Start with the Codex team guide , Claude team guide , or a concrete independent code-review workflow .
Ad-hoc delegation is easy to begin and surprisingly hard to operate. A parent spawns children, polls them, reconstructs ownership from chat, and hopes that a closed terminal or interrupted provider session did not erase the only useful account of what happened.
Oh My Subagents moves that coordination into durable runtime state:
When a Manager delegates a Wave, Oh My Subagents (OMS) commits the child Assignments, persists the parent wait, supervises every return, and continues the parent with the complete Checkpoints—even after an interruption.
The parent does not remain in a polling loop. A delegated Wave and its child Assignments commit together with the parent wait. Children return terminal Checkpoints independently; the controller collects the complete Wave and continues the parent with every return.
Managers can delegate recursively. Every Manager follows the same local rule with its direct children: delegate, wait, inspect the complete returns, and integrate. Deeper teams need no hidden global polling loop.
Here is one concrete run: evidence determines the next Wave, review findings become a new repair Assignment, and verification remains independent.
🧠 Reuse responsibility—not a frozen schedule
A Workflow definition answers who is responsible . It does not prescribe when a Member runs or force the work into a DAG.
At runtime, Managers choose the pattern that fits the actual Task and current evidence:
run dependent work sequentially;
fan independent work out in parallel;
iterate through implement, review, and repair;
divide a bounded batch among reusable owners;
combine those patterns in one Task; or
replan one responsibility subtree when the current team no longer fits.
The same published team can therefore respond differently to two different missions. A replan changes future responsibility without rewriting earlier revisions, completed work, or accepted history.
🧾 Accountability is a runtime contract
OMS does more than display several agents at once:
One stable starting contract. Every Task pins the exact published Workflow revision it started with.
One owner per Assignment. Child work is immutable, task-specific, and tied to the Member responsible for returning it.
Durable fan-out and fan-in. Waves, waits, retries, replans, Checkpoints, and continuations live in controller-owned state—not only in a provider transcript.
One accountable Result. Child completion is evidence for a parent. Only the Task lead's accepted completed or blocked Checkpoint becomes the Result shown to you.
Honest recovery. Browser closure, provider interruption, and controller restart do not silently fabricate completion or discard accepted history.
Ordinary files stay ordinary. Notes, reports, code, and artifacts remain in your workspace. OMS records small navigation references instead of pretending to own or snapshot every byte.
Capabilities deny by default and never inherit from a parent. Provider, model, sandbox, Human Request, and managed Command Run choices stay explicit where a team needs them without becoming mandatory ceremony for every Workflow.
🎛️ Design and operate visually
The visual Console is the primary experience:
Workflow library keeps reusable teams, drafts, and published revisions together.
Workflow Studio lets you shape the complete responsibility hierarchy on a horizontal canvas, validate it, and publish deliberately.
Run Studio shows the live team, current plans, meaningful Activity, Human Requests, managed Actions, referenced files, and the exact completed or blocked Result.
Steering delivers bounded new context to one exact active Member without pretending earlier work or tool effects never happened.
Prefer conversation? The separate Operator can draft and revise Workflows, explain teams, publish when asked, start and control Runs, answer Human Requests, and inspect managed Actions. Operator and the visual interface call the same controller-owned operations, so chat never creates a second hidden copy of product truth.
Oh My Subagents requires Python 3.12 or newer and supports Linux, macOS 13+, and Windows 11 x64. Install it in an isolated environment with pipx :
pipx install oh-my-subagents
oms init
oms service install
Open http://127.0.0.1:18125/ .
Guided initialization selects a default workspace, configures a Codex or Claude Task provider, publishes the Starter teams, and can configure the separate Operator. SQLite is the default, so a local installation needs no database server.
Upgrade an existing Banksia installation
Migrate before running oms init , oms serve , or oms service install . Initialization creates a fresh controller database; migration preserves the existing Banksia config, database, provider credentials, Task history, and service state under canonical OMS paths.
banksia status --json
banksia service status --json
pipx uninstall banksia
pipx install oh-my-subagents
oms migrate-from-banksia
oms db upgrade
oms service status --json
If the old installation did not use a background service, run oms service install after migration. Do not run oms init first, and do not use oms db reset as a rename step. The Banksia migration guide covers custom paths, verification, rollback, and recovery from an accidental 0.3.0 initialization.
oms service install verifies the configuration and database schema, installs a native per-user background service, and starts it. The controller keeps supervising work after you close the terminal and returns at login.
oms service status
oms service restart
oms service logs --lines 200
oms service stop
Linux uses a systemd user service, macOS a current-user LaunchAgent, and Windows a current-user Scheduled Task. oms service uninstall removes the native service definition while preserving configuration, database, and provider credentials. Use oms serve when you prefer the portable foreground path.
Install the optional driver and provide a SQLAlchemy URL during initialization:
pipx install " oh-my-subagents[postgres] "
oms init \
--database-url " postgresql+asyncpg://oms@127.0.0.1/oms "
See Getting started for provider prerequisites and a complete first run, or Database configuration for PostgreSQL permissions, schemas, and environment overrides.
oms init publishes eight provider-neutral Starters. Choose one when a mission is consequential or broad enough that independent ownership, durable work, or adversarial verification adds real value:
The Starter catalog includes example missions, expected deliverables, and guidance on when a simpler team is better. A one-Member Workflow is valid—and usually wiser—when delegation would add ceremony without independent evidence, useful specialization, or real integration.
Choose or design a team. Start from an included Starter Workflow or shape a responsibility tree in Workflow Studio.
Publish a stable revision. Every Run keeps the exact team contract it started with.
Give the lead one complete prompt. Managers plan, delegate, adapt, and integrate while OMS supervises runtime state.
Stay involved where judgment matters. Steer an active Member or answer a typed Human Request when the team genuinely needs you.
Receive the accountable Result. Read the lead's exact completed or blocked outcome and follow its references into detailed workspace files.
Understand Workflows and teams
Oh My Subagents is open source under the MIT License , except for the visual Console in console/ , which contains material derived from n8n and is distributed under the Sustainable Use License . That license permits internal business, non-commercial, and personal use; redistribution is limited to free, non-commercial distribution. See the Console notice for attribution and the modification notice.
Local subagent orchestration for Codex and Claude, with persistent task state, reusable teams, and interruption recovery.
ringlochid.me/oh-my-subagents/ Topics
Readme MIT license Contributing
7 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
