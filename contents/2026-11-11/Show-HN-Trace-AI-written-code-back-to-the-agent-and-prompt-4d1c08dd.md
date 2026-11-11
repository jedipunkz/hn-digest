---
source: "https://github.com/vrannang1/sploom"
hn_url: "https://news.ycombinator.com/item?id=50455166"
title: "Show HN: Trace AI written code back to the agent and prompt"
article_title: "GitHub - vrannang1/sploom: Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git. · GitHub"
image: "https://opengraph.githubassets.com/61dc2dfa2cc40863adbd6ea666ff5cd91713938d04bb6d5957be7c697f691d4c/vrannang1/sploom"
author: "annrap1d"
captured_at: "2026-11-11T16:56:03Z"
capture_tool: "hn-digest"
hn_id: 50455166
score: 1
comments: 0
posted_at: "2026-11-11T16:21:29Z"
tags:
  - hacker-news
---

# Show HN: Trace AI written code back to the agent and prompt

- HN: [50455166](https://news.ycombinator.com/item?id=50455166)
- Source: [github.com](https://github.com/vrannang1/sploom)
- Score: 1
- Comments: 0
- Posted: 2026-11-11T16:21:29Z

## Translation

Title: Show HN: Trace AI written code back to the agent and prompt
Article title: GitHub - vrannang1/sploom: Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git. · GitHub
Description: Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git. - vrannang1/sploom

Article text:
GitHub - vrannang1/sploom: Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git. · GitHub
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
vrannang1
sploom Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github/ workflows .github/ workflows hooks hooks sploom-cli sploom-cli sploom-core sploom-core sploomd sploomd web web .gitignore .gitignore .sploomignore .sploomignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md RFC-0002-intent-graph.md RFC-0002-intent-graph.md View all files Repository files navigation
Know which agent and model wrote every line, and the intent, prompts,
decisions and checks behind it. Sploom lives beside git, not instead of it: run
sploom init once, then keep using git as usual.
Works with Claude Code, Codex, Cursor, Gemini CLI, Copilot CLI, Devin CLI,
Grok Build, Goose and Cline , and with any MCP-capable agent through
sploom mcp-serve .
sploom why src/auth.rs:142 : which prompt, agent, model and intent
produced this code, which constraints apply, and which checks passed.
sploom rewind --green : go back to the last state where the checks
passed, between commits. sploom forward undoes it.
sploom context "…" : gives the next agent the relevant symbols,
constraints and decisions, within a token budget. Runs on your machine: the
first use downloads a small embedding model (about 130 MB), then it works offline.
Website: sploom.pages.dev · Design:
RFC-0002
cargo install --git https://github.com/vrannang1/sploom sploom sploomd
This installs two binaries: sploom , the CLI, and sploomd , a small
background file watcher.
cd your-project
sploom init
sploom init is the only setup step. It:
records a baseline of your current commit, so sploom why works immediately
installs git hooks: commit, push, pull, amend and rebase carry provenance
(your existing hooks keep working)
asks which coding agents on your machine to connect, and adds Sploom's hooks
to each one's project settings. Nothing changes without a yes (see
Agents )
imports your teammates' Sploom notes if the remote has any
Check it at any time. The first line of sploom status tells you whether
capture is on:
Capture: ON (Claude Code ✓ · Codex ✓ · git hooks ✓ · sploomd ✓)
sploom hooks remove takes everything out again.
You run
Sploom does
git commit
Records a weave of exactly what you committed, credited to the agent run that wrote it, and attaches a git note
git commit --amend , git rebase
Rewritten commits keep their provenance
git push
Also pushes refs/notes/sploom , in the background. It never slows or fails the push
git pull , git checkout
Imports teammates' notes and new commits
What stays on your machine ( .sploom/ , never committed): tool output, the
snapshots between commits, the vector index, and the readable copy of prompts.
What travels with git ( refs/notes/sploom ): intents, requirements,
constraints, decisions, which agent and model wrote each change, verification
results, and prompts, encrypted to your team's SSH keys .
Prompts are encrypted to your team's SSH keys
Prompts are shared with the people who should read them, and nobody else. Each
prompt is encrypted with age to the SSH public
keys listed in the repository's committed .sploomkeys file. These are the
same keys you already use with git. GitHub or GitLab only ever store
ciphertext.
sploom keys add github:alice # Alice's keys from github.com/alice.keys (gitlab:… works too)
git add .sploomkeys && git commit -m " Alice can read our prompts "
sploom keys list # who can read; which of your keys are used
sploom init adds your own key. A teammate's sploom why decrypts with their
own SSH private key, asking for its passphrase if it has one, and keeps the
readable copy locally. Someone who isn't listed still sees who wrote the code,
which model, the intent and the decisions, but not the prompt. Supported keys
are ssh-ed25519 and ssh-rsa . To keep prompts on your machine, set
share_prompts = "off" under [bridge] in .sploom/config.toml . "plain"
shares them redacted but readable by anyone with the repository.
sploom why <file:line | symbol>
Why this code exists
sploom log , sploom status
Weave history; is capture on
sploom rewind [--green | <weave>] , sploom forward
Journaled, reversible restore
sploom context "<query>"
Hybrid retrieval for the next agent
sploom intent , sploom guard , sploom decision
Intents, guardrails (constraints), decisions
sploom verify , sploom impact "<change>"
Verification evidence; impact analysis
sploom session start "<what>" [--worktree]
Name a piece of work; optional isolated worktree
sploom merge , sploom rebase
Three-level merge with an audit trail
sploom bridge push | pull | import
Manual git-notes sync; import existing history
sploom privacy purge , sploom gc
Retention and garbage collection
sploom mcp-serve
MCP tools for any agent
sploom hooks install | remove | status
Manage the integrations
sploom keys list | add | remove | unlock
Who can read shared prompts (SSH keys in .sploomkeys )
Agents
Sploom captures each agent through that agent's own hooks: prompts, file edits,
check commands and turns, credited to the agent and model that made them.
sploom hooks status # which agents are found and connected
sploom hooks install --agents codex,cursor # connect specific agents
sploom hooks remove # disconnect everything
Codex runs new project hooks only after you approve them once with /hooks ,
Grok Build after /hooks-trust , and Gemini CLI asks you to trust them. Cline
needs "Enable Hooks" switched on in its VS Code settings. Goose loads hooks
only from plugins, so Sploom installs a small per-user plugin that does nothing
outside Sploom repositories. Devin's cloud sessions run on Cognition's
machines; Sploom sees their work as commits (untracked), and Devin CLI runs are
captured fully. Sploom keeps the hook files it creates
out of git ( .git/info/exclude ), and leaves your own settings in them
untouched.
An edit that no connected agent reported, whether made by hand or by an agent
that isn't connected, is shown as untracked . It is never credited to you
or to the wrong agent. Any other MCP-capable agent can use Sploom's tools
through sploom mcp-serve . See hooks/README.md .
Secrets are redacted before anything is written to disk. The patterns cover
cloud and API keys, tokens, JWTs, private keys, credentials in URLs and .env
secrets. Prompts and tool output are stored apart from the graph, can be
encrypted (XChaCha20-Poly1305) and follow a retention policy you choose.
Everything runs locally. The only data that leaves your machine is in the git
notes you push, and the prompts in them are encrypted to your team's SSH keys.
Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git. - vrannang1/sploom

GitHub - vrannang1/sploom: Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git. · GitHub
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
vrannang1
sploom Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
.github/ workflows .github/ workflows hooks hooks sploom-cli sploom-cli sploom-core sploom-core sploomd sploomd web web .gitignore .gitignore .sploomignore .sploomignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md RFC-0002-intent-graph.md RFC-0002-intent-graph.md View all files Repository files navigation
Know which agent and model wrote every line, and the intent, prompts,
decisions and checks behind it. Sploom lives beside git, not instead of it: run
sploom init once, then keep using git as usual.
Works with Claude Code, Codex, Cursor, Gemini CLI, Copilot CLI, Devin CLI,
Grok Build, Goose and Cline , and with any MCP-capable agent through
sploom mcp-serve .
sploom why src/auth.rs:142 : which prompt, agent, model and intent
produced this code, which constraints apply, and which checks passed.
sploom rewind --green : go back to the last state where the checks
passed, between commits. sploom forward undoes it.
sploom context "…" : gives the next agent the relevant symbols,
constraints and decisions, within a token budget. Runs on your machine: the
first use downloads a small embedding model (about 130 MB), then it works offline.
Website: sploom.pages.dev · Design:
RFC-0002
cargo install --git https://github.com/vrannang1/sploom sploom sploomd
This installs two binaries: sploom , the CLI, and sploomd , a small
background file watcher.
cd your-project
sploom init
sploom init is the only setup step. It:
records a baseline of your current commit, so sploom why works immediately
installs git hooks: commit, push, pull, amend and rebase carry provenance
(your existing hooks keep working)
asks which coding agents on your machine to connect, and adds Sploom's hooks
to each one's project settings. Nothing changes without a yes (see
Agents )
imports your teammates' Sploom notes if the remote has any
Check it at any time. The first line of sploom status tells you whether
capture is on:
Capture: ON (Claude Code ✓ · Codex ✓ · git hooks ✓ · sploomd ✓)
sploom hooks remove takes everything out again.
You run
Sploom does
git commit
Records a weave of exactly what you committed, credited to the agent run that wrote it, and attaches a git note
git commit --amend , git rebase
Rewritten commits keep their provenance
git push
Also pushes refs/notes/sploom , in the background. It never slows or fails the push
git pull , git checkout
Imports teammates' notes and new commits
What stays on your machine ( .sploom/ , never committed): tool output, the
snapshots between commits, the vector index, and the readable copy of prompts.
What travels with git ( refs/notes/sploom ): intents, requirements,
constraints, decisions, which agent and model wrote each change, verification
results, and prompts, encrypted to your team's SSH keys .
Prompts are encrypted to your team's SSH keys
Prompts are shared with the people who should read them, and nobody else. Each
prompt is encrypted with age to the SSH public
keys listed in the repository's committed .sploomkeys file. These are the
same keys you already use with git. GitHub or GitLab only ever store
ciphertext.
sploom keys add github:alice # Alice's keys from github.com/alice.keys (gitlab:… works too)
git add .sploomkeys && git commit -m " Alice can read our prompts "
sploom keys list # who can read; which of your keys are used
sploom init adds your own key. A teammate's sploom why decrypts with their
own SSH private key, asking for its passphrase if it has one, and keeps the
readable copy locally. Someone who isn't listed still sees who wrote the code,
which model, the intent and the decisions, but not the prompt. Supported keys
are ssh-ed25519 and ssh-rsa . To keep prompts on your machine, set
share_prompts = "off" under [bridge] in .sploom/config.toml . "plain"
shares them redacted but readable by anyone with the repository.
sploom why <file:line | symbol>
Why this code exists
sploom log , sploom status
Weave history; is capture on
sploom rewind [--green | <weave>] , sploom forward
Journaled, reversible restore
sploom context "<query>"
Hybrid retrieval for the next agent
sploom intent , sploom guard , sploom decision
Intents, guardrails (constraints), decisions
sploom verify , sploom impact "<change>"
Verification evidence; impact analysis
sploom session start "<what>" [--worktree]
Name a piece of work; optional isolated worktree
sploom merge , sploom rebase
Three-level merge with an audit trail
sploom bridge push | pull | import
Manual git-notes sync; import existing history
sploom privacy purge , sploom gc
Retention and garbage collection
sploom mcp-serve
MCP tools for any agent
sploom hooks install | remove | status
Manage the integrations
sploom keys list | add | remove | unlock
Who can read shared prompts (SSH keys in .sploomkeys )
Agents
Sploom captures each agent through that agent's own hooks: prompts, file edits,
check commands and turns, credited to the agent and model that made them.
sploom hooks status # which agents are found and connected
sploom hooks install --agents codex,cursor # connect specific agents
sploom hooks remove # disconnect everything
Codex runs new project hooks only after you approve them once with /hooks ,
Grok Build after /hooks-trust , and Gemini CLI asks you to trust them. Cline
needs "Enable Hooks" switched on in its VS Code settings. Goose loads hooks
only from plugins, so Sploom installs a small per-user plugin that does nothing
outside Sploom repositories. Devin's cloud sessions run on Cognition's
machines; Sploom sees their work as commits (untracked), and Devin CLI runs are
captured fully. Sploom keeps the hook files it creates
out of git ( .git/info/exclude ), and leaves your own settings in them
untouched.
An edit that no connected agent reported, whether made by hand or by an agent
that isn't connected, is shown as untracked . It is never credited to you
or to the wrong agent. Any other MCP-capable agent can use Sploom's tools
through sploom mcp-serve . See hooks/README.md .
Secrets are redacted before anything is written to disk. The patterns cover
cloud and API keys, tokens, JWTs, private keys, credentials in URLs and .env
secrets. Prompts and tool output are stored apart from the graph, can be
encrypted (XChaCha20-Poly1305) and follow a retention policy you choose.
Everything runs locally. The only data that leaves your machine is in the git
notes you push, and the prompts in them are encrypted to your team's SSH keys.
Your codebase remembers why: which agent and model wrote each line, and the intent behind it. Lives beside git.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
