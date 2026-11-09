---
source: "https://github.com/dooly6276/claude-bulkhead"
hn_url: "https://news.ycombinator.com/item?id=50419582"
title: "Show HN: Replay your Claude Code logs to count parallel session collisions"
article_title: "GitHub - dooly6276/claude-bulkhead: Run many Claude Code sessions on one machine safely: no killed dev servers, port clashes, RAM or disk freezes, or mixed-up commits. Zero-dependency hooks. · GitHub"
image: "https://repository-images.githubusercontent.com/1416425499/75ea6f9c-2793-4b23-ae00-c229a2fa03da"
author: "dooly6276"
captured_at: "2026-11-09T01:28:08Z"
capture_tool: "hn-digest"
hn_id: 50419582
score: 2
comments: 0
posted_at: "2026-11-09T00:55:37Z"
tags:
  - hacker-news
---

# Show HN: Replay your Claude Code logs to count parallel session collisions

- HN: [50419582](https://news.ycombinator.com/item?id=50419582)
- Source: [github.com](https://github.com/dooly6276/claude-bulkhead)
- Score: 2
- Comments: 0
- Posted: 2026-11-09T00:55:37Z

## Translation

Title: Show HN: Replay your Claude Code logs to count parallel session collisions
Article title: GitHub - dooly6276/claude-bulkhead: Run many Claude Code sessions on one machine safely: no killed dev servers, port clashes, RAM or disk freezes, or mixed-up commits. Zero-dependency hooks. · GitHub
Description: Run many Claude Code sessions on one machine safely: no killed dev servers, port clashes, RAM or disk freezes, or mixed-up commits. Zero-dependency hooks. - dooly6276/claude-bulkhead

Article text:
GitHub - dooly6276/claude-bulkhead: Run many Claude Code sessions on one machine safely: no killed dev servers, port clashes, RAM or disk freezes, or mixed-up commits. Zero-dependency hooks. · GitHub
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
dooly6276
claude-bulkhead Public Notifications You must be signed in to change notification settings
Star 20 ( 20 ) You must be signed in to star a repository
Run many Claude Code sessions on one machine safely: no killed dev servers, port clashes, RAM or disk freezes, or mixed-up commits. Zero-dependency hooks.
Readme MIT license Contributing
Security policy Activity Stars
3 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
17 Commits 17 Commits Folders and files
.claude-plugin .claude-plugin .github .github bin bin commands commands docs docs hooks hooks scripts scripts skills/ bulkhead skills/ bulkhead src src test test .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.es.md README.es.md README.ja.md README.ja.md README.ko.md README.ko.md README.md README.md README.pt-BR.md README.pt-BR.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md config.example.json config.example.json eslint.config.js eslint.config.js package.json package.json View all files Repository files navigation
English |
한국어 |
简体中文 |
日本語 |
Español |
Português (Brasil)
Hooks that stop parallel Claude Code sessions on one machine from killing each other's servers or running it out of RAM.
8 guards, 10 commands, a status line, 6 slash commands, 1 skill, zero dependencies
What is it |
Scan |
Why |
Features |
Install |
Quick start |
Remote |
How it works |
Guards |
FAQ
New in 0.2.0: bulkhead scan replays your past Claude Code sessions against the guards, a new disk-guard stops whole-disk and parallel scans, checkout-guard now stops git add -A from sweeping another session's edits into your commit and rm -rf of a folder another session works in, and basics stops commits authored as Claude. Release notes: CHANGELOG.md .
Claude Code hooks are small programs that Claude Code runs at fixed points, for example before every shell command, and that can stop the command and tell Claude why.
bulkhead is a set of these hooks for one situation: several Claude Code sessions, and their subagents, working on the same computer at the same time. Each session acts as if it had the machine to itself. bulkhead checks each shell command against what the other sessions are doing: which processes and ports belong to which session, how much RAM is left, and which git tree each session works in. When a command would hurt another session, it is blocked, and Claude gets a message with the safer command to run instead.
You need it if you run two or more Claude Code sessions at once, often one per git worktree, or your sessions and subagents start dev servers, builds and test runs.
You probably do not need it if you run one session at a time and do not let it start servers or builds. Most of what bulkhead does is about sessions getting in each other's way.
Do you need it? Ask your transcripts
You do not have to guess. Claude Code keeps every session's transcript in ~/.claude/projects , and bulkhead scan replays the shell commands in them against the guards: how many sessions you really run at once, and how often they already did the things bulkhead blocks.
/plugin install bulkhead@bulkhead # then, inside Claude Code:
/bulkhead:scan 30d
It only reads, masks anything that looks like a secret before printing, and sends nothing anywhere. port-guard and the low-RAM check are not replayed, because transcripts do not record which ports and how much RAM were in use; checkout-guard is rebuilt from which sessions were active in the same git tree at that minute. Details: docs/commands.md .
A common way to speed up work with Claude Code is to run several sessions on one repository, each in its own git worktree, each with its own subagents. git keeps the branches apart. The machine does not: every session shares the same ports, the same process table and the same RAM. That is where they collide.
Two worktrees of the same app both run npm run dev . The second one fails on port 3000, or quietly moves to 3001 while its agent keeps testing localhost:3000 and sees the other branch.
An agent restarts its own server with pkill node or taskkill /IM node.exe and takes down the dev servers and test runs of every other worktree, and sometimes its own session.
Three sessions start npm test , next build and a Playwright run in the same minute, and the machine stops responding.
A session ends and the dev servers it started keep running.
The Claude Code tracker has many reports of exactly this:
#13941 : asked to restart a dev server, Claude ran taskkill //IM node.exe or pkill node and killed its own session.
#34754 : taskkill //F //IM python.exe killed two long-running simulations, after the user had marked their PIDs as off limits.
#43225 : asked to restart one bot, Claude force-killed every Python process on the machine, about 20 services.
#53229 : told to stop only unused dev servers, an agent also killed a gateway and a Postgres process it had not identified.
#15487 : 24 subagent processes saturated a 4 GB VPS until it had to be rebooted; a commenter froze a 24 GB Mac the same way.
#11502 and #25180 : processes started by agents outlive them, 60 Node processes in one report and 550 pytest workers in the other.
claude-bulkhead is a set of small hooks that block or clean up each of these. Each one can be turned off on its own. The name is from the bulkhead pattern: one compartment taking on water should not sink the ship.
kill-guard
Blocks pkill node , taskkill /IM , Stop-Process -Name and other kills by program name. Kills by PID pass.
port-guard
Blocks starting a dev server on a port another session already serves, and killing that session's server by port.
memory-guard and bulkhead-slot
Blocks heavy commands when RAM is low, and queues heavy jobs against one RAM budget shared by every session.
checkout-guard
Blocks git checkout , stash , reset --hard , git add -A and similar when another live session works in the same tree, and deleting a folder another session works in.
wait-guard
Blocks foreground sleep loops longer than 3 minutes, which make the next request rewrite the whole prompt cache.
orphan-reaper
When a session exits, stops the dev servers, test workers, watchers and headless browsers it left running.
disk-guard new
Blocks du -sh ~ , find / without -maxdepth , and disk scans run in parallel, which stall the disk every session shares.
bulkhead scan new
Replays your past Claude Code transcripts against the guards: what they would have stopped, and how many sessions ran at once.
basics
Blocks rm -rf at a root, home or a folder directly in it, force-pushes to main , resets that would lose uncommitted work, and commits authored as Claude.
bulkhead status , top , log , doctor
Shows every session, its RAM, its ports, queued jobs, and what the guards have blocked.
Status line
3 sessions 2 heavy 9.1 GB free in the Claude Code status line, in under 100 ms.
Config and slash commands
Per-machine and per-project config with allow lists, plus /bulkhead:scan , /bulkhead:status , /bulkhead:log , /bulkhead:doctor and /bulkhead:reap .
Optional: bulkhead remote runs a test suite, a build or a browser run on another machine over SSH, when this one has no RAM to spare.
Recorded with docs/demo.sh . The two sessions are stand-in processes named claude ; everything printed comes from the hook scripts and bulkhead status .
/plugin marketplace add dooly6276/claude-bulkhead
/plugin install bulkhead@bulkhead
Node 18 or later must be on PATH . The native Claude Code installer does not install Node, so check with node --version first. Without Node every hook fails to start: nothing is blocked, and Claude Code shows a hook error notice for each hook on every tool call. There are no other dependencies.
Install the plugin as above and start a new Claude Code session.
See what it would have caught: run /bulkhead:scan 30d , or bulkhead scan in the Bash tool. It reads your last month of transcripts and changes nothing.
Check the setup. Inside Claude Code run /bulkhead:doctor , or run bulkhead doctor in the Bash tool. Every line should say OK . A WARN or FAIL line names the guard and the reason.
Open a terminal next to your sessions and run bulkhead top to watch sessions, RAM, ports and heavy jobs live. Press q to quit.
After that there is nothing to do. The guards run on every Bash and PowerShell call. To see what they stopped, run bulkhead log or /bulkhead:log ; bulkhead log --summary counts blocks per guard on your machine. To change a threshold or allow a command, see docs/configuration.md . All commands are listed in docs/commands.md .
Run heavy jobs on another machine (optional)
When several sessions run test suites and builds at once, one machine runs out of RAM: memory-guard blocks the commands, or bulkhead-slot makes them queue. If you have a second machine that is mostly idle, bulkhead remote runs the heavy command there instead.
bulkhead remote add buildbox dev@192.168.1.20 --ssh-option=IdentityFile=~/.ssh/buildbox
bulkhead remote run -- npm test
[bulkhead remote] buildbox: 3f2a1c9 already there, 3 changed files (24 KB), 1 deleted
[bulkhead remote] buildbox copy a /home/dev/bulkhead-remote/acme-web-3f9a1c port 7172
...test output...
[bulkhead remote] exit 0 after 1m 12s (copy a)
What is sent. The current commit as a git bundle (only the commits the remote does not have yet) and your uncommitted changes, deletions included. Ignored files stay here unless you name them with --files . Nothing outside the repository is sent.
Where it runs. In one of up to 8 copies of the repository on the remote, so parallel runs do not share a working tree. node_modules is kept between runs. Each run gets its own port in $BULKHEAD_PORT .
Room. Locks on the remote hold a copy and a job slot, and --gb=N waits until N GB is free there. When there is no room within --wait minutes, it exits 3.
Cleanup. When the command ends, or the SSH connection drops (Ctrl-C, a tool timeout), everything it started on the remote is stopped, including a server left on its port.
The remote needs only an SSH server, bash , git and tar . Linux, macOS and Windows with Git for Windows all work.
bulkhead remote fetch PATH brings files back from the copy the last run used, and bulkhead remote status shows each remote's free RAM and running jobs. In Claude Code, /bulkhead:remote <command> runs a command there in the background. Setup, options and exit codes: docs/remote.md .
flowchart LR
CC["Claude Code session"] -->|"PreToolUse: Bash,

[truncated]

## Original Extract

Run many Claude Code sessions on one machine safely: no killed dev servers, port clashes, RAM or disk freezes, or mixed-up commits. Zero-dependency hooks. - dooly6276/claude-bulkhead

GitHub - dooly6276/claude-bulkhead: Run many Claude Code sessions on one machine safely: no killed dev servers, port clashes, RAM or disk freezes, or mixed-up commits. Zero-dependency hooks. · GitHub
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
dooly6276
claude-bulkhead Public Notifications You must be signed in to change notification settings
Star 20 ( 20 ) You must be signed in to star a repository
Run many Claude Code sessions on one machine safely: no killed dev servers, port clashes, RAM or disk freezes, or mixed-up commits. Zero-dependency hooks.
Readme MIT license Contributing
Security policy Activity Stars
3 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
17 Commits 17 Commits Folders and files
.claude-plugin .claude-plugin .github .github bin bin commands commands docs docs hooks hooks scripts scripts skills/ bulkhead skills/ bulkhead src src test test .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE PRIVACY.md PRIVACY.md README.es.md README.es.md README.ja.md README.ja.md README.ko.md README.ko.md README.md README.md README.pt-BR.md README.pt-BR.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md config.example.json config.example.json eslint.config.js eslint.config.js package.json package.json View all files Repository files navigation
English |
한국어 |
简体中文 |
日本語 |
Español |
Português (Brasil)
Hooks that stop parallel Claude Code sessions on one machine from killing each other's servers or running it out of RAM.
8 guards, 10 commands, a status line, 6 slash commands, 1 skill, zero dependencies
What is it |
Scan |
Why |
Features |
Install |
Quick start |
Remote |
How it works |
Guards |
FAQ
New in 0.2.0: bulkhead scan replays your past Claude Code sessions against the guards, a new disk-guard stops whole-disk and parallel scans, checkout-guard now stops git add -A from sweeping another session's edits into your commit and rm -rf of a folder another session works in, and basics stops commits authored as Claude. Release notes: CHANGELOG.md .
Claude Code hooks are small programs that Claude Code runs at fixed points, for example before every shell command, and that can stop the command and tell Claude why.
bulkhead is a set of these hooks for one situation: several Claude Code sessions, and their subagents, working on the same computer at the same time. Each session acts as if it had the machine to itself. bulkhead checks each shell command against what the other sessions are doing: which processes and ports belong to which session, how much RAM is left, and which git tree each session works in. When a command would hurt another session, it is blocked, and Claude gets a message with the safer command to run instead.
You need it if you run two or more Claude Code sessions at once, often one per git worktree, or your sessions and subagents start dev servers, builds and test runs.
You probably do not need it if you run one session at a time and do not let it start servers or builds. Most of what bulkhead does is about sessions getting in each other's way.
Do you need it? Ask your transcripts
You do not have to guess. Claude Code keeps every session's transcript in ~/.claude/projects , and bulkhead scan replays the shell commands in them against the guards: how many sessions you really run at once, and how often they already did the things bulkhead blocks.
/plugin install bulkhead@bulkhead # then, inside Claude Code:
/bulkhead:scan 30d
It only reads, masks anything that looks like a secret before printing, and sends nothing anywhere. port-guard and the low-RAM check are not replayed, because transcripts do not record which ports and how much RAM were in use; checkout-guard is rebuilt from which sessions were active in the same git tree at that minute. Details: docs/commands.md .
A common way to speed up work with Claude Code is to run several sessions on one repository, each in its own git worktree, each with its own subagents. git keeps the branches apart. The machine does not: every session shares the same ports, the same process table and the same RAM. That is where they collide.
Two worktrees of the same app both run npm run dev . The second one fails on port 3000, or quietly moves to 3001 while its agent keeps testing localhost:3000 and sees the other branch.
An agent restarts its own server with pkill node or taskkill /IM node.exe and takes down the dev servers and test runs of every other worktree, and sometimes its own session.
Three sessions start npm test , next build and a Playwright run in the same minute, and the machine stops responding.
A session ends and the dev servers it started keep running.
The Claude Code tracker has many reports of exactly this:
#13941 : asked to restart a dev server, Claude ran taskkill //IM node.exe or pkill node and killed its own session.
#34754 : taskkill //F //IM python.exe killed two long-running simulations, after the user had marked their PIDs as off limits.
#43225 : asked to restart one bot, Claude force-killed every Python process on the machine, about 20 services.
#53229 : told to stop only unused dev servers, an agent also killed a gateway and a Postgres process it had not identified.
#15487 : 24 subagent processes saturated a 4 GB VPS until it had to be rebooted; a commenter froze a 24 GB Mac the same way.
#11502 and #25180 : processes started by agents outlive them, 60 Node processes in one report and 550 pytest workers in the other.
claude-bulkhead is a set of small hooks that block or clean up each of these. Each one can be turned off on its own. The name is from the bulkhead pattern: one compartment taking on water should not sink the ship.
kill-guard
Blocks pkill node , taskkill /IM , Stop-Process -Name and other kills by program name. Kills by PID pass.
port-guard
Blocks starting a dev server on a port another session already serves, and killing that session's server by port.
memory-guard and bulkhead-slot
Blocks heavy commands when RAM is low, and queues heavy jobs against one RAM budget shared by every session.
checkout-guard
Blocks git checkout , stash , reset --hard , git add -A and similar when another live session works in the same tree, and deleting a folder another session works in.
wait-guard
Blocks foreground sleep loops longer than 3 minutes, which make the next request rewrite the whole prompt cache.
orphan-reaper
When a session exits, stops the dev servers, test workers, watchers and headless browsers it left running.
disk-guard new
Blocks du -sh ~ , find / without -maxdepth , and disk scans run in parallel, which stall the disk every session shares.
bulkhead scan new
Replays your past Claude Code transcripts against the guards: what they would have stopped, and how many sessions ran at once.
basics
Blocks rm -rf at a root, home or a folder directly in it, force-pushes to main , resets that would lose uncommitted work, and commits authored as Claude.
bulkhead status , top , log , doctor
Shows every session, its RAM, its ports, queued jobs, and what the guards have blocked.
Status line
3 sessions 2 heavy 9.1 GB free in the Claude Code status line, in under 100 ms.
Config and slash commands
Per-machine and per-project config with allow lists, plus /bulkhead:scan , /bulkhead:status , /bulkhead:log , /bulkhead:doctor and /bulkhead:reap .
Optional: bulkhead remote runs a test suite, a build or a browser run on another machine over SSH, when this one has no RAM to spare.
Recorded with docs/demo.sh . The two sessions are stand-in processes named claude ; everything printed comes from the hook scripts and bulkhead status .
/plugin marketplace add dooly6276/claude-bulkhead
/plugin install bulkhead@bulkhead
Node 18 or later must be on PATH . The native Claude Code installer does not install Node, so check with node --version first. Without Node every hook fails to start: nothing is blocked, and Claude Code shows a hook error notice for each hook on every tool call. There are no other dependencies.
Install the plugin as above and start a new Claude Code session.
See what it would have caught: run /bulkhead:scan 30d , or bulkhead scan in the Bash tool. It reads your last month of transcripts and changes nothing.
Check the setup. Inside Claude Code run /bulkhead:doctor , or run bulkhead doctor in the Bash tool. Every line should say OK . A WARN or FAIL line names the guard and the reason.
Open a terminal next to your sessions and run bulkhead top to watch sessions, RAM, ports and heavy jobs live. Press q to quit.
After that there is nothing to do. The guards run on every Bash and PowerShell call. To see what they stopped, run bulkhead log or /bulkhead:log ; bulkhead log --summary counts blocks per guard on your machine. To change a threshold or allow a command, see docs/configuration.md . All commands are listed in docs/commands.md .
Run heavy jobs on another machine (optional)
When several sessions run test suites and builds at once, one machine runs out of RAM: memory-guard blocks the commands, or bulkhead-slot makes them queue. If you have a second machine that is mostly idle, bulkhead remote runs the heavy command there instead.
bulkhead remote add buildbox dev@192.168.1.20 --ssh-option=IdentityFile=~/.ssh/buildbox
bulkhead remote run -- npm test
[bulkhead remote] buildbox: 3f2a1c9 already there, 3 changed files (24 KB), 1 deleted
[bulkhead remote] buildbox copy a /home/dev/bulkhead-remote/acme-web-3f9a1c port 7172
...test output...
[bulkhead remote] exit 0 after 1m 12s (copy a)
What is sent. The current commit as a git bundle (only the commits the remote does not have yet) and your uncommitted changes, deletions included. Ignored files stay here unless you name them with --files . Nothing outside the repository is sent.
Where it runs. In one of up to 8 copies of the repository on the remote, so parallel runs do not share a working tree. node_modules is kept between runs. Each run gets its own port in $BULKHEAD_PORT .
Room. Locks on the remote hold a copy and a job slot, and --gb=N waits until N GB is free there. When there is no room within --wait minutes, it exits 3.
Cleanup. When the command ends, or the SSH connection drops (Ctrl-C, a tool timeout), everything it started on the remote is stopped, including a server left on its port.
The remote needs only an SSH server, bash , git and tar . Linux, macOS and Windows with Git for Windows all work.
bulkhead remote fetch PATH brings files back from the copy the last run used, and bulkhead remote status shows each remote's free RAM and running jobs. In Claude Code, /bulkhead:remote <command> runs a command there in the background. Setup, options and exit codes: docs/remote.md .
flowchart LR
CC["Claude Code session"] -->|"PreToolUse: Bash,

[truncated]
