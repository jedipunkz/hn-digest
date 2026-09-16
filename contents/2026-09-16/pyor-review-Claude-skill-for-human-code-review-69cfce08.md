---
source: "https://github.com/pyor-review/pyor-cli"
hn_url: "https://news.ycombinator.com/item?id=49727013"
title: "/pyor:review – Claude skill for human code review"
article_title: "GitHub - Pyor-review/pyor-cli: Agent tooling for Pyor, the native home for GitHub code review. Works with Claude Code, Codex, Cursor, and any coding agent. · GitHub"
image: "https://opengraph.githubassets.com/2ab58a4a4abf85d9ceef4715be85284c213cb83bd7bf35dfe0a7832c7c3e66ed/Pyor-review/pyor-cli"
author: "othmanosx"
captured_at: "2026-09-16T14:39:43Z"
capture_tool: "hn-digest"
hn_id: 49727013
score: 1
comments: 0
posted_at: "2026-09-16T13:48:40Z"
tags:
  - hacker-news
---

# /pyor:review – Claude skill for human code review

- HN: [49727013](https://news.ycombinator.com/item?id=49727013)
- Source: [github.com](https://github.com/pyor-review/pyor-cli)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T13:48:40Z

## Translation

Title: /pyor:review – Claude skill for human code review
Article title: GitHub - Pyor-review/pyor-cli: Agent tooling for Pyor, the native home for GitHub code review. Works with Claude Code, Codex, Cursor, and any coding agent. · GitHub
Description: Agent tooling for Pyor, the native home for GitHub code review. Works with Claude Code, Codex, Cursor, and any coding agent. - Pyor-review/pyor-cli

Article text:
GitHub - Pyor-review/pyor-cli: Agent tooling for Pyor, the native home for GitHub code review. Works with Claude Code, Codex, Cursor, and any coding agent. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Pyor-review
/
pyor-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
39 Commits 39 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows assets assets plugins/ pyor plugins/ pyor scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
Agent tooling for Pyor , the native home for GitHub code
review. Works with Claude Code, Codex, Cursor, and any coding agent.
Watch the 5-minute demo : /pyor:review in Claude Code, notes in Pyor, fixes back in the branch.
Open your working changes as a local pre-PR review in the Pyor desktop app,
before a PR exists. /pyor:review has a panel of fresh sub-agents read the diff
and hand Pyor two things: file groups (the changes sorted into labelled
folders, important first) and inline hints (short pointers to risks worth a
look). No API key, no in-app model run.
The AI here is deliberately quiet: it points you at what matters and gets out
of the way. No summaries, no walls of generated text to wade through, just
groups to orient you and the occasional hint on a line worth a second look. You
read the actual code.
/pyor:review , the primary command. Authors an AI review by default.
A panel of fresh sub-agents produces the file grouping + inline hints from
the diff, then hands them to Pyor to render. The reviewers run with no
context from the calling session (they see only the code, so the review
isn't biased by whoever wrote it) across three lenses: correctness,
simplicity, and security. After you read and comment in Pyor, click "Send to
Claude" to push your comments straight back into the session that opened it.
Takes an optional grouping intent:
/pyor:review # importance: signal vs noise, important first (default)
/pyor:review walkthrough # order the groups as a reading path through the change
/pyor:review custom " by feature " # your own grouping instruction
/pyor:review plain # skip the AI panel, just open the diff (alias: fast)
Both open your working changes as a local pre-PR review, ready to read before a
PR exists; re-running reopens the same review (idempotent).
Claude Code , install the plugin:
/plugin marketplace add Pyor-review/pyor-cli
/plugin install pyor@pyor
Any other agent (Codex, Cursor, Amp, Gemini CLI, Windsurf, Zed, …), install
the skill with skills , which mirrors
it into your agent's format:
npx skills add https://github.com/Pyor-review/pyor-cli/tree/main/plugins/pyor/skills/pyor-review
Either way, then run /pyor:review (or invoke the pyor-review skill) from any
repository. Both drive the same pyor-review flow.
The Pyor desktop app . It registers the pyor:// URL scheme, renders the
review, and exports the review context the AI panel writes against. Grab it at
pyor.review ; /pyor:review hands you the install
command if it's missing.
gh (the GitHub CLI) is optional, used only when a command needs GitHub data.
Both the Claude plugin and the cross-agent skill drive one small, agent-neutral
CLI that does the deterministic parts: git resolution, the working-tree
revision token (which must match what the app computes so aids cache-hit), the
pyor:// deep link, and the ~/.pyor feedback channel. The agent only supplies
the analysis (grouping + hints).
npx -y pyor-review prepare
npx -y pyor-review --selftest
# or install it: npm install -g pyor-review
# no-install fallback before it's on npm:
# npx -y --package github:Pyor-review/pyor-cli pyor-review prepare
Subcommands: prepare (resolve repo/head/base + a deterministic session id +
revision + the exported review context), open (fire the deep link with the
aids), wait (park for "Send to Claude" feedback). pyor-local-review is the
plain, no-AI open.
The scripts are self-contained and repo-agnostic, sharing
plugins/pyor/scripts/lib.mjs :
node plugins/pyor/scripts/pyor-local-review.mjs --print # print the deep link, do not open
node plugins/pyor/scripts/pyor-ai-review.mjs prepare # resolve repo/head/base + revision + context
node plugins/pyor/scripts/pyor-ai-review.mjs --selftest # run the built-in assertions (incl. session-id determinism)
About
Agent tooling for Pyor, the native home for GitHub code review. Works with Claude Code, Codex, Cursor, and any coding agent.
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Agent tooling for Pyor, the native home for GitHub code review. Works with Claude Code, Codex, Cursor, and any coding agent. - Pyor-review/pyor-cli

GitHub - Pyor-review/pyor-cli: Agent tooling for Pyor, the native home for GitHub code review. Works with Claude Code, Codex, Cursor, and any coding agent. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
Pyor-review
/
pyor-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
39 Commits 39 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows assets assets plugins/ pyor plugins/ pyor scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
Agent tooling for Pyor , the native home for GitHub code
review. Works with Claude Code, Codex, Cursor, and any coding agent.
Watch the 5-minute demo : /pyor:review in Claude Code, notes in Pyor, fixes back in the branch.
Open your working changes as a local pre-PR review in the Pyor desktop app,
before a PR exists. /pyor:review has a panel of fresh sub-agents read the diff
and hand Pyor two things: file groups (the changes sorted into labelled
folders, important first) and inline hints (short pointers to risks worth a
look). No API key, no in-app model run.
The AI here is deliberately quiet: it points you at what matters and gets out
of the way. No summaries, no walls of generated text to wade through, just
groups to orient you and the occasional hint on a line worth a second look. You
read the actual code.
/pyor:review , the primary command. Authors an AI review by default.
A panel of fresh sub-agents produces the file grouping + inline hints from
the diff, then hands them to Pyor to render. The reviewers run with no
context from the calling session (they see only the code, so the review
isn't biased by whoever wrote it) across three lenses: correctness,
simplicity, and security. After you read and comment in Pyor, click "Send to
Claude" to push your comments straight back into the session that opened it.
Takes an optional grouping intent:
/pyor:review # importance: signal vs noise, important first (default)
/pyor:review walkthrough # order the groups as a reading path through the change
/pyor:review custom " by feature " # your own grouping instruction
/pyor:review plain # skip the AI panel, just open the diff (alias: fast)
Both open your working changes as a local pre-PR review, ready to read before a
PR exists; re-running reopens the same review (idempotent).
Claude Code , install the plugin:
/plugin marketplace add Pyor-review/pyor-cli
/plugin install pyor@pyor
Any other agent (Codex, Cursor, Amp, Gemini CLI, Windsurf, Zed, …), install
the skill with skills , which mirrors
it into your agent's format:
npx skills add https://github.com/Pyor-review/pyor-cli/tree/main/plugins/pyor/skills/pyor-review
Either way, then run /pyor:review (or invoke the pyor-review skill) from any
repository. Both drive the same pyor-review flow.
The Pyor desktop app . It registers the pyor:// URL scheme, renders the
review, and exports the review context the AI panel writes against. Grab it at
pyor.review ; /pyor:review hands you the install
command if it's missing.
gh (the GitHub CLI) is optional, used only when a command needs GitHub data.
Both the Claude plugin and the cross-agent skill drive one small, agent-neutral
CLI that does the deterministic parts: git resolution, the working-tree
revision token (which must match what the app computes so aids cache-hit), the
pyor:// deep link, and the ~/.pyor feedback channel. The agent only supplies
the analysis (grouping + hints).
npx -y pyor-review prepare
npx -y pyor-review --selftest
# or install it: npm install -g pyor-review
# no-install fallback before it's on npm:
# npx -y --package github:Pyor-review/pyor-cli pyor-review prepare
Subcommands: prepare (resolve repo/head/base + a deterministic session id +
revision + the exported review context), open (fire the deep link with the
aids), wait (park for "Send to Claude" feedback). pyor-local-review is the
plain, no-AI open.
The scripts are self-contained and repo-agnostic, sharing
plugins/pyor/scripts/lib.mjs :
node plugins/pyor/scripts/pyor-local-review.mjs --print # print the deep link, do not open
node plugins/pyor/scripts/pyor-ai-review.mjs prepare # resolve repo/head/base + revision + context
node plugins/pyor/scripts/pyor-ai-review.mjs --selftest # run the built-in assertions (incl. session-id determinism)
About
Agent tooling for Pyor, the native home for GitHub code review. Works with Claude Code, Codex, Cursor, and any coding agent.
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
