---
source: "https://github.com/adityamatt/anumati"
hn_url: "https://news.ycombinator.com/item?id=49608606"
title: "Deterministic Rule based auto-approver for Claude/Codex"
article_title: "GitHub - adityamatt/anumati: A claude auto approver based on rules/configs · GitHub"
image: "https://opengraph.githubassets.com/37b5214d5e55a6c9b9ab24f0d699540c12e2d336db15b6a23ad6ecb0695d2e4d/adityamatt/anumati"
author: "aditaymatt"
captured_at: "2026-09-08T11:28:11Z"
capture_tool: "hn-digest"
hn_id: 49608606
score: 1
comments: 1
posted_at: "2026-09-08T11:00:53Z"
tags:
  - hacker-news
---

# Deterministic Rule based auto-approver for Claude/Codex

- HN: [49608606](https://news.ycombinator.com/item?id=49608606)
- Source: [github.com](https://github.com/adityamatt/anumati)
- Score: 1
- Comments: 1
- Posted: 2026-09-08T11:00:53Z

## Translation

Title: Deterministic Rule based auto-approver for Claude/Codex
Article title: GitHub - adityamatt/anumati: A claude auto approver based on rules/configs · GitHub
Description: A claude auto approver based on rules/configs. Contribute to adityamatt/anumati development by creating an account on GitHub.

Article text:
GitHub - adityamatt/anumati: A claude auto approver based on rules/configs · GitHub
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
adityamatt
/
anumati
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
78 Commits 78 Commits Folders and files
.github/ workflows .github/ workflows docs docs scripts scripts src src tests tests workflows workflows .DS_Store .DS_Store .gitignore .gitignore .npmrc .npmrc AGENT.md AGENT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SUGGEST-FEATURE.md SUGGEST-FEATURE.md example.json example.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Stop approving the same Bash commands over and over. anumati is a
deterministic auto-approver for AI coding agents. From a small config of
named matchers it auto-allows safe shell commands — so git status ,
npx tsc , cargo test , jq , and friends just run, while genuinely risky
commands still get a prompt. It's agent-agnostic : one shared config drives
both Claude Code (via its PreToolUse hook) and
OpenAI Codex (via its PermissionRequest
hook).
When something falls through, anumati tells you the exact one-liner to allow it
next time — so your config builds itself from real usage .
$ git status && cargo test | tail -20
✓ auto-approved (no prompt)
$ terraform apply
⤳ prompt shown · 💡 anumati: no matcher covers "terraform"
Why anumati
Deterministic, not vibes. Approval is decided by explicit matchers with a
strict grammar — not an LLM guessing whether a command is safe. The same
command always gets the same answer.
Safe by construction. Matchers allow only read-only / build / test shapes.
Redirects that write files, $(...) substitution, network curl to unlisted
domains, git push , rm — all fall through to a real prompt. anumati is
allow-only : it can approve a call or step aside, but never blocks anything
itself, and never widens what the agent would otherwise refuse.
Composes across a whole command line. git status && cargo build | tail
is approved only if every piece is independently safe — you can't smuggle
rm -rf / in by chaining it onto an allowed command.
Self-building config. Every passthrough comes with a verified suggestion
( anumati add … ) and a logged reason, so you extend coverage from what you
actually run.
Bash-only by design. anumati vets Bash — the hard problem. Read /
Write / Edit stay with the agent's own permission flow.
Every time the agent is about to run a Bash command, anumati checks it
against your allow rules:
A rule matches → auto-approved, no prompt.
No rule matches → the agent shows its normal permission prompt, and
anumati prints a 💡 suggestion for allowing it next time.
A command is approved one of two ways: a single matcher accepts the whole thing,
or — failing that — anumati splits it at top-level && , ; , || , & , and
newlines and approves only if every sub-command is independently accepted.
flowchart TD
A[Bash command] --> B{"A single rule's matcher<br/>accepts the whole command?"}
B -- yes --> ALLOW([✅ allow])
B -- no --> D["Split at top-level && ; || & and newlines<br/>pipes stay glued to their segment"]
D --> F["For each sub-command:<br/>does some rule accept it?"]
F --> G{"Every sub-command<br/>approved?"}
G -- yes --> ALLOW
G -- no --> PASS([⤳ passthrough])
Loading
A disallowed sub-command always fails its own check, so chaining a bad command
onto a good one can't sneak it through. Pipes are never split across rules (a
pipe feeds data into the next command, so only the matcher owning the pipeline
can judge it). Configs cascade: a project config at
<cwd>/.anumati/permissions.json is checked before your global
~/.anumati/permissions.json (the legacy ~/.claude/ locations are still
honored, so existing setups keep working).
The full model — matchers, composition rules, and safety guarantees — is in
docs/CONFIGURATION.md .
npm install -g anumati
Or run without installing via npx anumati ~/.claude/permissions.json .
Update to the latest version any time with anumati update (it checks the
published version and, if newer, runs the global reinstall for you; the hook
picks it up on the next command). anumati update --check just reports whether
a newer version exists.
One command sets up everything:
anumati init
It prompts for which agent(s) (Claude Code / Codex / both) and project
(this folder) or root (global) scope, then:
Writes a starter config of broadly-useful, low-risk rules — read-only
inspection, git reads, cd / sleep / echo / sed / jq , npx tsc , cargo ,
go , test runners ( vitest / pytest / jest ), and pure-compute
python3 / node . Enough to be useful immediately. It also seeds the
parameterized matchers ( curl , gh , pip3-install , git-write ,
git-push , node-script ) as inert placeholders — empty allowlists that
approve nothing until you fill them in, so you can see the matcher exists and
which key to populate ( anumati add <matcher> … ) without hunting the docs.
Scaffolds an audit log next to the config.
Wires the chosen agent(s) , merged non-destructively — Claude Code's
PreToolUse hook in settings.json , and/or Codex's PermissionRequest hook
in ~/.codex/hooks.json .
Adds a SessionStart banner ("⚡ anumati active") so you can see it's on.
Writes command-style guidance to the sibling CLAUDE.md , nudging the
agent to emit approvable commands.
Reload the agent for it to take effect (Claude Code: /hooks or restart;
Codex: approve the anumati hook once when it prompts to review it). Then just
work — routine commands stop prompting, and when something new falls through
you'll see a 💡 anumati add … suggestion.
anumati add curl --domain api.github.com # allow curl to a domain
anumati add git-write --git-ops add,commit # allow specific git writes
anumati stats # see your auto-approve rate
anumati apply --all # apply accumulated suggestions
Docs
Configuration reference — every matcher, rule
field, audit/sound/debug option, and CLI subcommand.
Claude Code + Codex — running anumati across both
agents from one shared config: per-agent wiring, the Codex integration, and
anumati migrate .
Command-style guide — how to write commands that
land on the auto-approve path (also installed into CLAUDE.md by init ).
See CONTRIBUTING.md for local setup — including running
Claude Code against your local build via npm link — plus how to test, add a
matcher, and open a PR.
A claude auto approver based on rules/configs
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A claude auto approver based on rules/configs. Contribute to adityamatt/anumati development by creating an account on GitHub.

GitHub - adityamatt/anumati: A claude auto approver based on rules/configs · GitHub
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
adityamatt
/
anumati
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
78 Commits 78 Commits Folders and files
.github/ workflows .github/ workflows docs docs scripts scripts src src tests tests workflows workflows .DS_Store .DS_Store .gitignore .gitignore .npmrc .npmrc AGENT.md AGENT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SUGGEST-FEATURE.md SUGGEST-FEATURE.md example.json example.json package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Stop approving the same Bash commands over and over. anumati is a
deterministic auto-approver for AI coding agents. From a small config of
named matchers it auto-allows safe shell commands — so git status ,
npx tsc , cargo test , jq , and friends just run, while genuinely risky
commands still get a prompt. It's agent-agnostic : one shared config drives
both Claude Code (via its PreToolUse hook) and
OpenAI Codex (via its PermissionRequest
hook).
When something falls through, anumati tells you the exact one-liner to allow it
next time — so your config builds itself from real usage .
$ git status && cargo test | tail -20
✓ auto-approved (no prompt)
$ terraform apply
⤳ prompt shown · 💡 anumati: no matcher covers "terraform"
Why anumati
Deterministic, not vibes. Approval is decided by explicit matchers with a
strict grammar — not an LLM guessing whether a command is safe. The same
command always gets the same answer.
Safe by construction. Matchers allow only read-only / build / test shapes.
Redirects that write files, $(...) substitution, network curl to unlisted
domains, git push , rm — all fall through to a real prompt. anumati is
allow-only : it can approve a call or step aside, but never blocks anything
itself, and never widens what the agent would otherwise refuse.
Composes across a whole command line. git status && cargo build | tail
is approved only if every piece is independently safe — you can't smuggle
rm -rf / in by chaining it onto an allowed command.
Self-building config. Every passthrough comes with a verified suggestion
( anumati add … ) and a logged reason, so you extend coverage from what you
actually run.
Bash-only by design. anumati vets Bash — the hard problem. Read /
Write / Edit stay with the agent's own permission flow.
Every time the agent is about to run a Bash command, anumati checks it
against your allow rules:
A rule matches → auto-approved, no prompt.
No rule matches → the agent shows its normal permission prompt, and
anumati prints a 💡 suggestion for allowing it next time.
A command is approved one of two ways: a single matcher accepts the whole thing,
or — failing that — anumati splits it at top-level && , ; , || , & , and
newlines and approves only if every sub-command is independently accepted.
flowchart TD
A[Bash command] --> B{"A single rule's matcher<br/>accepts the whole command?"}
B -- yes --> ALLOW([✅ allow])
B -- no --> D["Split at top-level && ; || & and newlines<br/>pipes stay glued to their segment"]
D --> F["For each sub-command:<br/>does some rule accept it?"]
F --> G{"Every sub-command<br/>approved?"}
G -- yes --> ALLOW
G -- no --> PASS([⤳ passthrough])
Loading
A disallowed sub-command always fails its own check, so chaining a bad command
onto a good one can't sneak it through. Pipes are never split across rules (a
pipe feeds data into the next command, so only the matcher owning the pipeline
can judge it). Configs cascade: a project config at
<cwd>/.anumati/permissions.json is checked before your global
~/.anumati/permissions.json (the legacy ~/.claude/ locations are still
honored, so existing setups keep working).
The full model — matchers, composition rules, and safety guarantees — is in
docs/CONFIGURATION.md .
npm install -g anumati
Or run without installing via npx anumati ~/.claude/permissions.json .
Update to the latest version any time with anumati update (it checks the
published version and, if newer, runs the global reinstall for you; the hook
picks it up on the next command). anumati update --check just reports whether
a newer version exists.
One command sets up everything:
anumati init
It prompts for which agent(s) (Claude Code / Codex / both) and project
(this folder) or root (global) scope, then:
Writes a starter config of broadly-useful, low-risk rules — read-only
inspection, git reads, cd / sleep / echo / sed / jq , npx tsc , cargo ,
go , test runners ( vitest / pytest / jest ), and pure-compute
python3 / node . Enough to be useful immediately. It also seeds the
parameterized matchers ( curl , gh , pip3-install , git-write ,
git-push , node-script ) as inert placeholders — empty allowlists that
approve nothing until you fill them in, so you can see the matcher exists and
which key to populate ( anumati add <matcher> … ) without hunting the docs.
Scaffolds an audit log next to the config.
Wires the chosen agent(s) , merged non-destructively — Claude Code's
PreToolUse hook in settings.json , and/or Codex's PermissionRequest hook
in ~/.codex/hooks.json .
Adds a SessionStart banner ("⚡ anumati active") so you can see it's on.
Writes command-style guidance to the sibling CLAUDE.md , nudging the
agent to emit approvable commands.
Reload the agent for it to take effect (Claude Code: /hooks or restart;
Codex: approve the anumati hook once when it prompts to review it). Then just
work — routine commands stop prompting, and when something new falls through
you'll see a 💡 anumati add … suggestion.
anumati add curl --domain api.github.com # allow curl to a domain
anumati add git-write --git-ops add,commit # allow specific git writes
anumati stats # see your auto-approve rate
anumati apply --all # apply accumulated suggestions
Docs
Configuration reference — every matcher, rule
field, audit/sound/debug option, and CLI subcommand.
Claude Code + Codex — running anumati across both
agents from one shared config: per-agent wiring, the Codex integration, and
anumati migrate .
Command-style guide — how to write commands that
land on the auto-approve path (also installed into CLAUDE.md by init ).
See CONTRIBUTING.md for local setup — including running
Claude Code against your local build via npm link — plus how to test, add a
matcher, and open a PR.
A claude auto approver based on rules/configs
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
