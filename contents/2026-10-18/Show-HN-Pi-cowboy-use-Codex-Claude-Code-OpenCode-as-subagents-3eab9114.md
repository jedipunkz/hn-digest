---
source: "https://github.com/electric-dragon-80000v/pi-cowboy"
hn_url: "https://news.ycombinator.com/item?id=50142277"
title: "Show HN: Pi-cowboy – use Codex, Claude Code, OpenCode as subagents"
article_title: "GitHub - electric-dragon-80000v/pi-cowboy: A Herdr-native agent orchestrator for Pi. In addition to Pi itself, it can run Codex, Claude Code, OpenCode, and more as subagents in isolated, copy-on-write Git worktrees. · GitHub"
image: "https://opengraph.githubassets.com/8c8ef74bf7bbeecf5e318d95fda84f6c5123f752533d545e248b1ce2a487bdb1/electric-dragon-80000v/pi-cowboy"
author: "suoloordi"
captured_at: "2026-10-18T16:28:21Z"
capture_tool: "hn-digest"
hn_id: 50142277
score: 2
comments: 0
posted_at: "2026-10-18T16:04:57Z"
tags:
  - hacker-news
---

# Show HN: Pi-cowboy – use Codex, Claude Code, OpenCode as subagents

- HN: [50142277](https://news.ycombinator.com/item?id=50142277)
- Source: [github.com](https://github.com/electric-dragon-80000v/pi-cowboy)
- Score: 2
- Comments: 0
- Posted: 2026-10-18T16:04:57Z

## Translation

Title: Show HN: Pi-cowboy – use Codex, Claude Code, OpenCode as subagents
Article title: GitHub - electric-dragon-80000v/pi-cowboy: A Herdr-native agent orchestrator for Pi. In addition to Pi itself, it can run Codex, Claude Code, OpenCode, and more as subagents in isolated, copy-on-write Git worktrees. · GitHub
Description: A Herdr-native agent orchestrator for Pi. In addition to Pi itself, it can run Codex, Claude Code, OpenCode, and more as subagents in isolated, copy-on-write Git worktrees. - electric-dragon-80000v/pi-cowboy
HN text: pi-cowboy is an agent orchestrator for Pi( https://pi.dev ) and Herdr( https://github.com/herdrdev/herdr ). Break large jobs into focused subtasks, hand them off to background agents running in parallel, and consolidate the results when they finish. Each agent runs in an isolated copy-on-write Git work-tree, minimising disk overhead. pi-cowboy can orchestrate foreign harnesses too, including Claude Code, Codex, OpenCode and more. Feature highlights: - Foreign harnesses:
pi-cowboy agents can run in Claude Code, Codex, OpenCode, Pig and Pi-Bolt, as well as Pi. This gives you great flexibility in choosing how to spend your tokens. Running out of Codex usage? Switch your pi-cowboy harness to OpenCode, select one of its free models, and let it handle the rest.
Pig and Pi-Bolt are pi variants that are more efficient, with a lower memory footprint and faster performance than pi. Because pi-cowboy is designed to run multiple agents in parallel, having a lower per-agent footprint is crucial. It prevents your machine from running out of resources and lets you have higher concurrency without choking your system. - Concurrency limits:
You can configure concurrency limits at every level. Running a local model with limited resources? If a simple task already takes hours to complete, the last thing you want is to overload your machine by running too many of them at once. Set concurrency limits for that model specifically to keep resource usage under control. You can install it on Pi as a package: pi install npm:pi-cowboy #(requires `herdr` 0.8+)

Article text:
GitHub - electric-dragon-80000v/pi-cowboy: A Herdr-native agent orchestrator for Pi. In addition to Pi itself, it can run Codex, Claude Code, OpenCode, and more as subagents in isolated, copy-on-write Git worktrees. · GitHub
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
electric-dragon-80000v
pi-cowboy Public Notifications You must be signed in to change notification settings
Star 33 ( 33 ) You must be signed in to star a repository
A Herdr-native agent orchestrator for Pi. In addition to Pi itself, it can run Codex, Claude Code, OpenCode, and more as subagents in isolated, copy-on-write Git worktrees.
pi.dev/packages/pi-cowboy Topics
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
230 Commits 230 Commits Folders and files
.agents/ skills .agents/ skills .github .github .husky .husky agents agents dev dev docs docs lint lint orchestrators orchestrators scripts scripts src src test test .gitattributes .gitattributes .gitignore .gitignore .prettierignore .prettierignore .tool-versions .tool-versions CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md commitlint.config.ts commitlint.config.ts oxlint.config.ts oxlint.config.ts package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml release.config.ts release.config.ts tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
pi-cowboy is a multi-agent orchestrator for pi . Break a large job into focused subtasks, hand them off to background agents running in parallel, and consolidate them when they finish.
Each agent runs in its own terminal pane inside Herdr and works in an isolated copy-on-write Git worktree. You can watch agents work in real time, jump into any pane to steer them, or mix and match harnesses like Pi, Claude Code, OpenCode, and Codex in the same session. When the work is finished, pi-cowboy cleans up the leftovers for you.
Focused tasks produce better code
Give each agent a smaller, isolated subtask. A single task with clear instructions leaves little room for error, whereas a long list of expectations invites drift and over-engineering. Because each agent commits to its own branch, you can review and merge each subtask on its own.
pi-cowboy can orchestrate Claude Code, Codex, OpenCode, and more: start work on Pi, continue on Claude Code, finish on OpenCode. Do you need support for something else? Create an issue on GitHub .
Parallel work with minimal disk overhead
Multiple agents can edit, build, and run tests at the same time without interfering with each other. Copy-on-write worktrees make high parallelism practical by keeping disk usage small even when your repository is in the magnitude of gigabytes.
The numbers below come from a Git repository with 153,607 files in total (479 MiB tracked, 27 GiB total including untracked files):
Each task gets its own shell, a visible terminal pane you can open at any time, and notifications when an agent needs your attention.
Define reusable agent templates for specialist roles and orchestrator templates to adapt how agents work and report back.
Grazing cows and a cowboy on your screen
Who wouldn’t want their agents visualized as cow emojis? Go on… guess. That's exactly right. No one.
herdr 0.8 or later, with the herdr pi integration enabled
python3 on macOS (used for copy-on-write worktrees; macOS has not shipped python3 since 12.3, so install the Xcode Command Line Tools with xcode-select --install )
Copy-on-write capable filesystem (APFS on macOS; BTRFS or ZFS on Linux)
ext4 does not support copy-on-write. On ext4 , agent.worktreeMaterialization falls back to checkout .
pi-cowboy activates automatically whenever you run pi inside a Git repository.
Based on the security audit, group the relevant files into 10 groups, and launch parallel cowboy agents to address their audit results.
pi-cowboy spawns background agents for the subtasks right away:
✓ Spawned agent b7g2nspx (general-purpose)
✓ Spawned agent 5m1xx7fa (general-purpose)
...
If you spawn more agents than your concurrency limit allows, the extras wait in a queue. With the default global limit of 4, the first 4 agents start immediately and the remaining 6 start as slots open up.
Each agent runs as a separate process in its own Herdr workspace, inside its own tab and terminal pane. You can switch to an agent's workspace at any time to watch its progress or steer it directly.
On disk, each agent works in a dedicated Git worktree under <agent dir>/pi-cowboy/worktrees/ ( ~/.pi/agent by default).
When an agent finishes its turn (settles), its report arrives as a message in the parent session:
✓ [Cowboy agent "general-purpose" b7g2nspx completed]
Addressed all the points raised in the security audit for the given set of files.
(Worktree: $HOME/.pi/agent/pi-cowboy/worktrees/cow-audit-perms-b7g2nspx (branch cow-audit-perms-b7g2nspx). The agent process and its herdr pane stay until you call cleanup_cowboy_agent — its pane is closed, which ends that process. The worktree stays until you call cleanup_cowboy_agent to remove it — call it once the branch is merged or rejected.)
The parent session can then review the report, steer the agent if something needs another pass, or merge the branch.
To bring the changes in, the orchestrator calls merge_cowboy_branch , which merges the agent's branch into your repository's main working tree:
✓ Merged "cow-audit-perms-b7g2nspx" into "main" in pi-mono.
Once the work is merged, the orchestrator cleans up the agent's resources (or you can ask it to):
cleanup_cowboy_agent closes the Herdr pane and tab, removes the worktree, and deletes the branch:
✓ Cleaned up agent b7g2nspx:
agent status: completed
pane: closed
worktree: removed ($HOME/.pi/agent/pi-cowboy/worktrees/cow-audit-perms-b7g2nspx)
branch: deleted (cow-audit-perms-b7g2nspx)
Cleanup is safe by default: it refuses to touch an agent that is still running, keeps any worktree that still has uncommitted changes, and deletes the branch only if it has already been merged. If one step cannot be completed, pi-cowboy reports the reason and finishes the rest.
When the /cowboy status menu refuses to remove a worktree because of uncommitted changes or an unreadable state, it offers to Force the removal. Force deletes the worktree and discards anything uncommitted in it, so it asks for a confirmation that names the checkout and what will be lost. Forced removals are reported as forced — <reason> . This override lives in the menu only — the cleanup_cowboy_agent tool never forces, and the menu only offers it for a settled agent, never one it still lists as running.
If you find yourself repeating the same instructions for recurring jobs, you can package them into an agent template with its own system prompt, model, harness, and tool permissions.
For example, a release-notes template looks like this:
name = " release-notes "
description = " Prepares release notes. "
system_prompt = '''
You prepare release notes for this repository. Read the history and tags
(git log, git tag) and report the changes as a message.
'''
You can then ask for it by name in the parent session:
See Agent Templates for all template fields and directory locations.
While agent templates specialize individual workers, orchestrator templates apply rules across every agent you spawn. They also let you customize the status messages ("cues") sent to the parent session when agents spawn, queue, or settle.
For instance, if you use an adversarial review tool, you can define an orchestrator template that requires every agent to run it before committing:
name = " adversarial-review "
guidance = '''
Run the adversarial-review tool before you commit, and fix any glaring mistakes that the review finds and then commit.
'''
You can switch the active orchestrator at any time from /cowboy > Default orchestrator . See Orchestrator Templates for details.
pi-cowboy auto-detects Codex, Claude Code, OpenCode, PiG, and Pi-Bolt when they are installed on your PATH . You can request a specific harness directly in your prompt (for example, "Launch an OpenCode agent for this task") or set defaults in your configuration and agent templates .
You can cap how many agents run at the same time across four independent levels:
Global: concurrency.default (default 4 ), which applies to every agent
Per harness: concurrency.harnesses.<harness-id>
Per provider: concurrency.providers.<provider>
Per model: concurrency.models["<provider>/<model-id>"]
A running agent occupies a slot in every limit that applies to it. When the agent settles, it releases those slots and pi-cowboy starts the next queued agent that has room at all of its levels. An agent waiting on a busy model limit will not block a queued agent behind it that uses a different model.
If you steer a settled agent, it wakes up in its existing pane and runs immediately, even if that temporarily puts a limit over capacity.
Configure these limits in config.json or live in your session via /cowboy > Settings > Concurrency Settings . See Configuration for details.
When pi-cowboy spawns an agent, it picks the model using the first match in this order:
The model parameter passed to the cowboy_agent tool call
The session per-type model override ( /cowboy > Model overrides )
The config per-type override ( agent.<agent-type> )
The model set in the agent template for the launching harness family
The session default model ( /cowboy model )
The launching harness's configured default ( agent.default for the Pi family, or agent.defaultModelByHarness for foreign harnesses)
For Pi-family launches, the parent session's active model
See Configuration for the full rules.
Run /cowboy to open the interactive menu, or pass a subcommand directly. See Menu for a guide to every screen.
/cowboy # Open the pi-cowboy menu
/cowboy status # List spawned, queued, and settled agents
/cowboy spawn # Spawn an agent manually
/cowboy worktree # Create a git worktree with no agent attached
/cowboy model # Pick the default agent model for this session
/cowboy model provider/model-id # Set the default agent model for this session
/cowboy model clear # Clear the session model override
/cowboy enable # Enable pi-cowboy
/cowboy disable # Disable pi-cowboy
Tools
Delegates one or more tasks to cowboy agents. By default, the call returns immediately with an agent ID for each task, and results arrive in the parent session as each agent settles.
Sends a follow-up message to one or more agents. If an agent is running, it receives the message before its next model call. If an agent has already settled, steering revives it in the same pane for another turn and delivers

[truncated]

## Original Extract

A Herdr-native agent orchestrator for Pi. In addition to Pi itself, it can run Codex, Claude Code, OpenCode, and more as subagents in isolated, copy-on-write Git worktrees. - electric-dragon-80000v/pi-cowboy

pi-cowboy is an agent orchestrator for Pi( https://pi.dev ) and Herdr( https://github.com/herdrdev/herdr ). Break large jobs into focused subtasks, hand them off to background agents running in parallel, and consolidate the results when they finish. Each agent runs in an isolated copy-on-write Git work-tree, minimising disk overhead. pi-cowboy can orchestrate foreign harnesses too, including Claude Code, Codex, OpenCode and more. Feature highlights: - Foreign harnesses:
pi-cowboy agents can run in Claude Code, Codex, OpenCode, Pig and Pi-Bolt, as well as Pi. This gives you great flexibility in choosing how to spend your tokens. Running out of Codex usage? Switch your pi-cowboy harness to OpenCode, select one of its free models, and let it handle the rest.
Pig and Pi-Bolt are pi variants that are more efficient, with a lower memory footprint and faster performance than pi. Because pi-cowboy is designed to run multiple agents in parallel, having a lower per-agent footprint is crucial. It prevents your machine from running out of resources and lets you have higher concurrency without choking your system. - Concurrency limits:
You can configure concurrency limits at every level. Running a local model with limited resources? If a simple task already takes hours to complete, the last thing you want is to overload your machine by running too many of them at once. Set concurrency limits for that model specifically to keep resource usage under control. You can install it on Pi as a package: pi install npm:pi-cowboy #(requires `herdr` 0.8+)

GitHub - electric-dragon-80000v/pi-cowboy: A Herdr-native agent orchestrator for Pi. In addition to Pi itself, it can run Codex, Claude Code, OpenCode, and more as subagents in isolated, copy-on-write Git worktrees. · GitHub
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
electric-dragon-80000v
pi-cowboy Public Notifications You must be signed in to change notification settings
Star 33 ( 33 ) You must be signed in to star a repository
A Herdr-native agent orchestrator for Pi. In addition to Pi itself, it can run Codex, Claude Code, OpenCode, and more as subagents in isolated, copy-on-write Git worktrees.
pi.dev/packages/pi-cowboy Topics
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
230 Commits 230 Commits Folders and files
.agents/ skills .agents/ skills .github .github .husky .husky agents agents dev dev docs docs lint lint orchestrators orchestrators scripts scripts src src test test .gitattributes .gitattributes .gitignore .gitignore .prettierignore .prettierignore .tool-versions .tool-versions CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md commitlint.config.ts commitlint.config.ts oxlint.config.ts oxlint.config.ts package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml release.config.ts release.config.ts tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
pi-cowboy is a multi-agent orchestrator for pi . Break a large job into focused subtasks, hand them off to background agents running in parallel, and consolidate them when they finish.
Each agent runs in its own terminal pane inside Herdr and works in an isolated copy-on-write Git worktree. You can watch agents work in real time, jump into any pane to steer them, or mix and match harnesses like Pi, Claude Code, OpenCode, and Codex in the same session. When the work is finished, pi-cowboy cleans up the leftovers for you.
Focused tasks produce better code
Give each agent a smaller, isolated subtask. A single task with clear instructions leaves little room for error, whereas a long list of expectations invites drift and over-engineering. Because each agent commits to its own branch, you can review and merge each subtask on its own.
pi-cowboy can orchestrate Claude Code, Codex, OpenCode, and more: start work on Pi, continue on Claude Code, finish on OpenCode. Do you need support for something else? Create an issue on GitHub .
Parallel work with minimal disk overhead
Multiple agents can edit, build, and run tests at the same time without interfering with each other. Copy-on-write worktrees make high parallelism practical by keeping disk usage small even when your repository is in the magnitude of gigabytes.
The numbers below come from a Git repository with 153,607 files in total (479 MiB tracked, 27 GiB total including untracked files):
Each task gets its own shell, a visible terminal pane you can open at any time, and notifications when an agent needs your attention.
Define reusable agent templates for specialist roles and orchestrator templates to adapt how agents work and report back.
Grazing cows and a cowboy on your screen
Who wouldn’t want their agents visualized as cow emojis? Go on… guess. That's exactly right. No one.
herdr 0.8 or later, with the herdr pi integration enabled
python3 on macOS (used for copy-on-write worktrees; macOS has not shipped python3 since 12.3, so install the Xcode Command Line Tools with xcode-select --install )
Copy-on-write capable filesystem (APFS on macOS; BTRFS or ZFS on Linux)
ext4 does not support copy-on-write. On ext4 , agent.worktreeMaterialization falls back to checkout .
pi-cowboy activates automatically whenever you run pi inside a Git repository.
Based on the security audit, group the relevant files into 10 groups, and launch parallel cowboy agents to address their audit results.
pi-cowboy spawns background agents for the subtasks right away:
✓ Spawned agent b7g2nspx (general-purpose)
✓ Spawned agent 5m1xx7fa (general-purpose)
...
If you spawn more agents than your concurrency limit allows, the extras wait in a queue. With the default global limit of 4, the first 4 agents start immediately and the remaining 6 start as slots open up.
Each agent runs as a separate process in its own Herdr workspace, inside its own tab and terminal pane. You can switch to an agent's workspace at any time to watch its progress or steer it directly.
On disk, each agent works in a dedicated Git worktree under <agent dir>/pi-cowboy/worktrees/ ( ~/.pi/agent by default).
When an agent finishes its turn (settles), its report arrives as a message in the parent session:
✓ [Cowboy agent "general-purpose" b7g2nspx completed]
Addressed all the points raised in the security audit for the given set of files.
(Worktree: $HOME/.pi/agent/pi-cowboy/worktrees/cow-audit-perms-b7g2nspx (branch cow-audit-perms-b7g2nspx). The agent process and its herdr pane stay until you call cleanup_cowboy_agent — its pane is closed, which ends that process. The worktree stays until you call cleanup_cowboy_agent to remove it — call it once the branch is merged or rejected.)
The parent session can then review the report, steer the agent if something needs another pass, or merge the branch.
To bring the changes in, the orchestrator calls merge_cowboy_branch , which merges the agent's branch into your repository's main working tree:
✓ Merged "cow-audit-perms-b7g2nspx" into "main" in pi-mono.
Once the work is merged, the orchestrator cleans up the agent's resources (or you can ask it to):
cleanup_cowboy_agent closes the Herdr pane and tab, removes the worktree, and deletes the branch:
✓ Cleaned up agent b7g2nspx:
agent status: completed
pane: closed
worktree: removed ($HOME/.pi/agent/pi-cowboy/worktrees/cow-audit-perms-b7g2nspx)
branch: deleted (cow-audit-perms-b7g2nspx)
Cleanup is safe by default: it refuses to touch an agent that is still running, keeps any worktree that still has uncommitted changes, and deletes the branch only if it has already been merged. If one step cannot be completed, pi-cowboy reports the reason and finishes the rest.
When the /cowboy status menu refuses to remove a worktree because of uncommitted changes or an unreadable state, it offers to Force the removal. Force deletes the worktree and discards anything uncommitted in it, so it asks for a confirmation that names the checkout and what will be lost. Forced removals are reported as forced — <reason> . This override lives in the menu only — the cleanup_cowboy_agent tool never forces, and the menu only offers it for a settled agent, never one it still lists as running.
If you find yourself repeating the same instructions for recurring jobs, you can package them into an agent template with its own system prompt, model, harness, and tool permissions.
For example, a release-notes template looks like this:
name = " release-notes "
description = " Prepares release notes. "
system_prompt = '''
You prepare release notes for this repository. Read the history and tags
(git log, git tag) and report the changes as a message.
'''
You can then ask for it by name in the parent session:
See Agent Templates for all template fields and directory locations.
While agent templates specialize individual workers, orchestrator templates apply rules across every agent you spawn. They also let you customize the status messages ("cues") sent to the parent session when agents spawn, queue, or settle.
For instance, if you use an adversarial review tool, you can define an orchestrator template that requires every agent to run it before committing:
name = " adversarial-review "
guidance = '''
Run the adversarial-review tool before you commit, and fix any glaring mistakes that the review finds and then commit.
'''
You can switch the active orchestrator at any time from /cowboy > Default orchestrator . See Orchestrator Templates for details.
pi-cowboy auto-detects Codex, Claude Code, OpenCode, PiG, and Pi-Bolt when they are installed on your PATH . You can request a specific harness directly in your prompt (for example, "Launch an OpenCode agent for this task") or set defaults in your configuration and agent templates .
You can cap how many agents run at the same time across four independent levels:
Global: concurrency.default (default 4 ), which applies to every agent
Per harness: concurrency.harnesses.<harness-id>
Per provider: concurrency.providers.<provider>
Per model: concurrency.models["<provider>/<model-id>"]
A running agent occupies a slot in every limit that applies to it. When the agent settles, it releases those slots and pi-cowboy starts the next queued agent that has room at all of its levels. An agent waiting on a busy model limit will not block a queued agent behind it that uses a different model.
If you steer a settled agent, it wakes up in its existing pane and runs immediately, even if that temporarily puts a limit over capacity.
Configure these limits in config.json or live in your session via /cowboy > Settings > Concurrency Settings . See Configuration for details.
When pi-cowboy spawns an agent, it picks the model using the first match in this order:
The model parameter passed to the cowboy_agent tool call
The session per-type model override ( /cowboy > Model overrides )
The config per-type override ( agent.<agent-type> )
The model set in the agent template for the launching harness family
The session default model ( /cowboy model )
The launching harness's configured default ( agent.default for the Pi family, or agent.defaultModelByHarness for foreign harnesses)
For Pi-family launches, the parent session's active model
See Configuration for the full rules.
Run /cowboy to open the interactive menu, or pass a subcommand directly. See Menu for a guide to every screen.
/cowboy # Open the pi-cowboy menu
/cowboy status # List spawned, queued, and settled agents
/cowboy spawn # Spawn an agent manually
/cowboy worktree # Create a git worktree with no agent attached
/cowboy model # Pick the default agent model for this session
/cowboy model provider/model-id # Set the default agent model for this session
/cowboy model clear # Clear the session model override
/cowboy enable # Enable pi-cowboy
/cowboy disable # Disable pi-cowboy
Tools
Delegates one or more tasks to cowboy agents. By default, the call returns immediately with an agent ID for each task, and results arrive in the parent session as each agent settles.
Sends a follow-up message to one or more agents. If an agent is running, it receives the message before its next model call. If an agent has already settled, steering revives it in the same pane for another turn and delivers

[truncated]
