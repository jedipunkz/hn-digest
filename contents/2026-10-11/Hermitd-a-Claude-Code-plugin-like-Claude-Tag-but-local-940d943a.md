---
source: "https://github.com/gtapps/hermitd"
hn_url: "https://news.ycombinator.com/item?id=50048120"
title: "Hermitd, a Claude Code plugin like Claude Tag, but local"
article_title: "GitHub - gtapps/hermitd: Like Claude Tag, but local. · GitHub"
image: "https://opengraph.githubassets.com/56956711886f0e3e0a388a79886f6faf5daf12dbb19902807283986497ee81f8/gtapps/hermitd"
author: "gtapps"
captured_at: "2026-10-11T22:41:42Z"
capture_tool: "hn-digest"
hn_id: 50048120
score: 2
comments: 1
posted_at: "2026-10-11T21:43:27Z"
tags:
  - hacker-news
---

# Hermitd, a Claude Code plugin like Claude Tag, but local

- HN: [50048120](https://news.ycombinator.com/item?id=50048120)
- Source: [github.com](https://github.com/gtapps/hermitd)
- Score: 2
- Comments: 1
- Posted: 2026-10-11T21:43:27Z

## Translation

Title: Hermitd, a Claude Code plugin like Claude Tag, but local
Article title: GitHub - gtapps/hermitd: Like Claude Tag, but local. · GitHub
Description: Like Claude Tag, but local. Contribute to gtapps/hermitd development by creating an account on GitHub.

Article text:
GitHub - gtapps/hermitd: Like Claude Tag, but local. · GitHub
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
gtapps
hermitd Public Notifications You must be signed in to change notification settings
Star 75 ( 75 ) You must be signed in to star a repository
Readme MIT license Contributing
16 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3,391 Commits 3,391 Commits Folders and files
.agents/ skills .agents/ skills .claude-plugin .claude-plugin .claude .claude .github .github .hermit .hermit plugins plugins scripts scripts tests tests .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .graphifyignore .graphifyignore .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Run an always-on Claude Code agent on your own machine, for you or your team. Start sessions from your chat in any project it can reach, on the model and effort you choose, and steer them through the agent or take over in the Claude app or CLI.
Use it from your terminal or the Claude app via Remote Control , or connect Discord, Telegram, iMessage, or a custom Claude Code channel .
Give it ongoing responsibilities: maintain research, monitor systems, run routines, and follow up on unfinished work. Between requests, it checks those responsibilities, carries progress across sessions, and reaches you when something needs attention.
Run it on your Claude subscription and extend it with your own MCP servers, skills, and plugins.
It's Claude Code. hermitd is a plugin for the official Claude Code CLI, which runs as Anthropic ships it. New Claude models, features, and fixes work as soon as Claude Code supports them.
Choose one installation method below. Run it from the folder where you want your agent, empty or existing. Uses your Claude subscription on Linux, macOS, or Windows via WSL2. See prerequisites .
With Claude Code 2.1.292+ and Bun 1.4+ installed:
claude plugin install hermitd --marketplace gtapps/hermitd --scope local
claude " /hermitd:hatch "
Fresh machine? Use the bootstrap installer
Prepares Claude Code, Bun, and tmux, installs the plugin, and launches setup:
curl -fsSL https://gtapps.github.io/hermitd/install.sh | bash
Both options set up the agent in this folder. Hatch guides you through the agent’s purpose and preferences, then shows how to start it. Choose Quick for defaults you can adjust later.
After setup, follow the printed next steps to start your agent.
Run in a persistent tmux session:
hermitd start
Requires tmux. The watchdog recovers failed sessions while your machine stays on. Claude Code's /sandbox is recommended for unattended use. To connect a chat, run /hermitd:channel-setup as directed by the setup handoff.
Run the guided setup in Claude Code:
/hermitd:docker-setup
Builds and starts the container, then walks you through authentication and channel pairing. Requires Docker Compose v2.
Customize the container. Ask the agent to add tools, packages, or services to its Docker setup. For example: “Add ffmpeg to the container.”
Optional Docker security controls cover local-network access, DNS policy, resource limits, and plugin installation auditing.
Continuity. Persistent working state and archived session handoffs carry progress across compaction and restarts. An external watchdog recovers failed sessions, while context management keeps long-running sessions manageable.
Proactive work. Heartbeats regularly check the responsibilities you give the agent. Routines run scheduled work, and watches surface changes. Together, they let the agent follow up without waiting for another request.
Work through chat. Assign work and receive results in your connected chat. Longer assignments get threaded progress updates, with a separate reply when the agent needs a decision. Assignments can also carry a persistent task record with requester, due date, result confirmation, and a dashboard view by person.
Your network. Local databases, SSH, LAN devices, and Home Assistant are reachable because the agent runs where they are.
Token efficiency. With Claude Code’s Monitor , heartbeat checks and optional routine prechecks run outside the model. Quiet checks and skipped routines use no model tokens; eligible routines due together can share a wake.
Lasting knowledge. Turn source material in raw/ into maintained knowledge in compiled/ , alongside Claude Code's auto memory. /recall searches past sessions, knowledge, proposals, and captured channel conversations.
Learning from experience. The agent reviews evidence from its work and operation, saves useful lessons, and verifies proposed behavior changes before bringing them to you for approval.
Control and visibility. Track progress, proposals, and usage through the dashboard. Pause is enforced at the tool boundary, and optional usage caps can alert you or pause further work.
Part of your project channel. With passive mode, the agent saves incoming group messages to look back on later, and wakes when someone you allow @mentions it. It also remembers instructions for that channel.
Tune from a terminal with /hermit-settings , or change permitted settings from a trusted Discord or Telegram chat. Every write is validated and recorded in a redacted audit ledger; /hermit-settings history [setting] shows what changed. Some of the settings available:
Full schema in the Config Reference
Artifacts. The agent uses Claude Code Artifacts to provide an interactive dashboard and custom pages generated on demand that you can view, interact with, and share. Ask it to build your own personalized agent dashboard.
Ask for an update from your terminal or connected chat:
The agent reviews evidence from its work and operation. Durable lessons go to memory; non-trivial ideas that would change its behavior are verified, deduplicated, and brought to you for approval.
Work produces evidence
│
▼
Reflect when due
│
▼
Verify and deduplicate
│
┌────┴────┐
▼ ▼
Remember Propose
a lesson a change
│
▼
You approve?
│ │
no yes
│ │
No change Implement
│
▼
Verify result
│
▼
Future evidence
Reflection runs at eligible task or session pauses, daily, and after routines configured to reflect. Approved changes can start now, become a task, or be left for manual implementation. Proposals are resolved when verification passes or later evidence shows the problem is gone.
Ask
What it does
“Start X in ~/code/api on Opus at high effort, in its own worktree.”
Starts a session in any project, on your model and effort.
“Tell me when my session in ~/code/web finishes.”
Watches your other Claude Code sessions.
!pause · !snooze 2h · !model sonnet · !effort high
Claude Code controls, straight from chat.
“ /later check tomorrow whether those errors have returned.”
Checks whether a fix held up over time.
“When I ask for a status update, include blockers.”
Remembers instructions for that channel.
“What else could you be doing for me?”
Proposes new capabilities from your work and its tools.
Cost
Quiet heartbeat checks, skipped routines, and passive chat capture use no model tokens. Work, evaluations, and replies consume usage; context management keeps conversation history bounded.
event: a file, a log line, a webhook, a schedule
│
▼
precheck (plain script, no model tokens)
│
├── nothing changed ──▶ back to sleep
│
▼ something to do
Claude Code turn on your machine
│
├── needs a decision ──▶ asks you in chat
│
▼
result in your chat + usage recorded
See what drives usage. Token usage is recorded per call, including the model, input/output/cache split, and whether work came from a routine, heartbeat, channel, or another source. Session and daily totals feed the dashboard, weekly review, and /cost-reflect .
Set limits. Optional daily, weekly, and monthly caps can alert you or enforce a pause until the exceeded budget window resets. Under Claude subscription billing, dollar figures are usage estimates rather than additional per-token charges.
Choose where to spend. Set the session model and optionally assign a different model to individual routines. Routine models run in isolated subagents, so use them for work that can return a concise result.
See budgets and routine scheduling for configuration and scheduler fallback behavior.
Reach the running agent through your connected channels or Claude Code Remote Control. You can also start separate sessions for additional work:
Background sessions with follow-up. Through /spawn-session , the agent launches a local Claude Code helper in the project, or in another folder with --cwd , and relays its status when it becomes idle. Claude Code isolates the helper's edits in a Git worktree unless the project sets worktree.bgIsolation to none ; pass --worktree to give it one from the start. Set the helper’s model and effort with options such as --model sonnet --effort high .
Local Remote Control gate. Through /rc-gate , the agent manages a Remote Control server on your machine or server. While the gate is open, you can spawn new Claude Code sessions from the Claude app, using your local files and tools. Each session gets its own Git worktree, while the agent keeps running.
Both session-spawning paths require a Git workspace. Remote Control requires a Claude sign-in through /login on the machine running the agent.
Watch other sessions. Through Claude Code cross-session messaging , ask the agent to watch a local Claude Code session, including one you started interactively, and notify you when it next becomes idle.
Claude Code controls from chat. Use !model sonnet , !effort high , !advisor opus , !compact , !clear , !doctor (alias !checkup ), and !permission-mode auto directly from your connected chat. Control the agent’s work with !pause , !resume , and !snooze 2h . Use /when-done-switch-to --model sonnet to switch automatically at the end of the current turn.
Optional plugins that add domain tools and workflows to your agent.
dev-hermit : Branch discipline, push guards, and gated PR workflows.
homeassistant-hermit : Home Assistant tools, automation workflows, and safety checks.
fitness-hermit : Strava integration, activity analysis, and training routines.
hermitd-feed : Source curation, recurring briefs, and weekly synthesis.
hermitd-laravel-forge : Laravel Forge deployments, logs, and server management.
hermit-scribe : GitHub issues and comments from proposals through a dedicated bot identity.
You can run separate agents for different responsibilities, each with its own working state, knowledge, and routines. See Creating Your Own Hermit .
External orchestration. Other agents and tools can check the agent’s status, health, and

[truncated]

## Original Extract

Like Claude Tag, but local. Contribute to gtapps/hermitd development by creating an account on GitHub.

GitHub - gtapps/hermitd: Like Claude Tag, but local. · GitHub
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
gtapps
hermitd Public Notifications You must be signed in to change notification settings
Star 75 ( 75 ) You must be signed in to star a repository
Readme MIT license Contributing
16 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3,391 Commits 3,391 Commits Folders and files
.agents/ skills .agents/ skills .claude-plugin .claude-plugin .claude .claude .github .github .hermit .hermit plugins plugins scripts scripts tests tests .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .graphifyignore .graphifyignore .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Run an always-on Claude Code agent on your own machine, for you or your team. Start sessions from your chat in any project it can reach, on the model and effort you choose, and steer them through the agent or take over in the Claude app or CLI.
Use it from your terminal or the Claude app via Remote Control , or connect Discord, Telegram, iMessage, or a custom Claude Code channel .
Give it ongoing responsibilities: maintain research, monitor systems, run routines, and follow up on unfinished work. Between requests, it checks those responsibilities, carries progress across sessions, and reaches you when something needs attention.
Run it on your Claude subscription and extend it with your own MCP servers, skills, and plugins.
It's Claude Code. hermitd is a plugin for the official Claude Code CLI, which runs as Anthropic ships it. New Claude models, features, and fixes work as soon as Claude Code supports them.
Choose one installation method below. Run it from the folder where you want your agent, empty or existing. Uses your Claude subscription on Linux, macOS, or Windows via WSL2. See prerequisites .
With Claude Code 2.1.292+ and Bun 1.4+ installed:
claude plugin install hermitd --marketplace gtapps/hermitd --scope local
claude " /hermitd:hatch "
Fresh machine? Use the bootstrap installer
Prepares Claude Code, Bun, and tmux, installs the plugin, and launches setup:
curl -fsSL https://gtapps.github.io/hermitd/install.sh | bash
Both options set up the agent in this folder. Hatch guides you through the agent’s purpose and preferences, then shows how to start it. Choose Quick for defaults you can adjust later.
After setup, follow the printed next steps to start your agent.
Run in a persistent tmux session:
hermitd start
Requires tmux. The watchdog recovers failed sessions while your machine stays on. Claude Code's /sandbox is recommended for unattended use. To connect a chat, run /hermitd:channel-setup as directed by the setup handoff.
Run the guided setup in Claude Code:
/hermitd:docker-setup
Builds and starts the container, then walks you through authentication and channel pairing. Requires Docker Compose v2.
Customize the container. Ask the agent to add tools, packages, or services to its Docker setup. For example: “Add ffmpeg to the container.”
Optional Docker security controls cover local-network access, DNS policy, resource limits, and plugin installation auditing.
Continuity. Persistent working state and archived session handoffs carry progress across compaction and restarts. An external watchdog recovers failed sessions, while context management keeps long-running sessions manageable.
Proactive work. Heartbeats regularly check the responsibilities you give the agent. Routines run scheduled work, and watches surface changes. Together, they let the agent follow up without waiting for another request.
Work through chat. Assign work and receive results in your connected chat. Longer assignments get threaded progress updates, with a separate reply when the agent needs a decision. Assignments can also carry a persistent task record with requester, due date, result confirmation, and a dashboard view by person.
Your network. Local databases, SSH, LAN devices, and Home Assistant are reachable because the agent runs where they are.
Token efficiency. With Claude Code’s Monitor , heartbeat checks and optional routine prechecks run outside the model. Quiet checks and skipped routines use no model tokens; eligible routines due together can share a wake.
Lasting knowledge. Turn source material in raw/ into maintained knowledge in compiled/ , alongside Claude Code's auto memory. /recall searches past sessions, knowledge, proposals, and captured channel conversations.
Learning from experience. The agent reviews evidence from its work and operation, saves useful lessons, and verifies proposed behavior changes before bringing them to you for approval.
Control and visibility. Track progress, proposals, and usage through the dashboard. Pause is enforced at the tool boundary, and optional usage caps can alert you or pause further work.
Part of your project channel. With passive mode, the agent saves incoming group messages to look back on later, and wakes when someone you allow @mentions it. It also remembers instructions for that channel.
Tune from a terminal with /hermit-settings , or change permitted settings from a trusted Discord or Telegram chat. Every write is validated and recorded in a redacted audit ledger; /hermit-settings history [setting] shows what changed. Some of the settings available:
Full schema in the Config Reference
Artifacts. The agent uses Claude Code Artifacts to provide an interactive dashboard and custom pages generated on demand that you can view, interact with, and share. Ask it to build your own personalized agent dashboard.
Ask for an update from your terminal or connected chat:
The agent reviews evidence from its work and operation. Durable lessons go to memory; non-trivial ideas that would change its behavior are verified, deduplicated, and brought to you for approval.
Work produces evidence
│
▼
Reflect when due
│
▼
Verify and deduplicate
│
┌────┴────┐
▼ ▼
Remember Propose
a lesson a change
│
▼
You approve?
│ │
no yes
│ │
No change Implement
│
▼
Verify result
│
▼
Future evidence
Reflection runs at eligible task or session pauses, daily, and after routines configured to reflect. Approved changes can start now, become a task, or be left for manual implementation. Proposals are resolved when verification passes or later evidence shows the problem is gone.
Ask
What it does
“Start X in ~/code/api on Opus at high effort, in its own worktree.”
Starts a session in any project, on your model and effort.
“Tell me when my session in ~/code/web finishes.”
Watches your other Claude Code sessions.
!pause · !snooze 2h · !model sonnet · !effort high
Claude Code controls, straight from chat.
“ /later check tomorrow whether those errors have returned.”
Checks whether a fix held up over time.
“When I ask for a status update, include blockers.”
Remembers instructions for that channel.
“What else could you be doing for me?”
Proposes new capabilities from your work and its tools.
Cost
Quiet heartbeat checks, skipped routines, and passive chat capture use no model tokens. Work, evaluations, and replies consume usage; context management keeps conversation history bounded.
event: a file, a log line, a webhook, a schedule
│
▼
precheck (plain script, no model tokens)
│
├── nothing changed ──▶ back to sleep
│
▼ something to do
Claude Code turn on your machine
│
├── needs a decision ──▶ asks you in chat
│
▼
result in your chat + usage recorded
See what drives usage. Token usage is recorded per call, including the model, input/output/cache split, and whether work came from a routine, heartbeat, channel, or another source. Session and daily totals feed the dashboard, weekly review, and /cost-reflect .
Set limits. Optional daily, weekly, and monthly caps can alert you or enforce a pause until the exceeded budget window resets. Under Claude subscription billing, dollar figures are usage estimates rather than additional per-token charges.
Choose where to spend. Set the session model and optionally assign a different model to individual routines. Routine models run in isolated subagents, so use them for work that can return a concise result.
See budgets and routine scheduling for configuration and scheduler fallback behavior.
Reach the running agent through your connected channels or Claude Code Remote Control. You can also start separate sessions for additional work:
Background sessions with follow-up. Through /spawn-session , the agent launches a local Claude Code helper in the project, or in another folder with --cwd , and relays its status when it becomes idle. Claude Code isolates the helper's edits in a Git worktree unless the project sets worktree.bgIsolation to none ; pass --worktree to give it one from the start. Set the helper’s model and effort with options such as --model sonnet --effort high .
Local Remote Control gate. Through /rc-gate , the agent manages a Remote Control server on your machine or server. While the gate is open, you can spawn new Claude Code sessions from the Claude app, using your local files and tools. Each session gets its own Git worktree, while the agent keeps running.
Both session-spawning paths require a Git workspace. Remote Control requires a Claude sign-in through /login on the machine running the agent.
Watch other sessions. Through Claude Code cross-session messaging , ask the agent to watch a local Claude Code session, including one you started interactively, and notify you when it next becomes idle.
Claude Code controls from chat. Use !model sonnet , !effort high , !advisor opus , !compact , !clear , !doctor (alias !checkup ), and !permission-mode auto directly from your connected chat. Control the agent’s work with !pause , !resume , and !snooze 2h . Use /when-done-switch-to --model sonnet to switch automatically at the end of the current turn.
Optional plugins that add domain tools and workflows to your agent.
dev-hermit : Branch discipline, push guards, and gated PR workflows.
homeassistant-hermit : Home Assistant tools, automation workflows, and safety checks.
fitness-hermit : Strava integration, activity analysis, and training routines.
hermitd-feed : Source curation, recurring briefs, and weekly synthesis.
hermitd-laravel-forge : Laravel Forge deployments, logs, and server management.
hermit-scribe : GitHub issues and comments from proposals through a dedicated bot identity.
You can run separate agents for different responsibilities, each with its own working state, knowledge, and routines. See Creating Your Own Hermit .
External orchestration. Other agents and tools can check the agent’s status, health, and

[truncated]
