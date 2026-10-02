---
source: "https://github.com/randomm/pi-delegate"
hn_url: "https://news.ycombinator.com/item?id=49933095"
title: "Show HN: Delegate Claude Code's work to pi and a cheaper model"
article_title: "GitHub - randomm/pi-delegate: Claude Code skill for delegating development and verification tasks to pi coding agent — with adversarial review loop · GitHub"
image: "https://opengraph.githubassets.com/297c2484bab599a12d1ea94fe449b394474f06dd80c76c4a876547a92e614655/randomm/pi-delegate"
author: "jannniii"
captured_at: "2026-10-02T13:38:13Z"
capture_tool: "hn-digest"
hn_id: 49933095
score: 1
comments: 0
posted_at: "2026-10-02T13:03:29Z"
tags:
  - hacker-news
---

# Show HN: Delegate Claude Code's work to pi and a cheaper model

- HN: [49933095](https://news.ycombinator.com/item?id=49933095)
- Source: [github.com](https://github.com/randomm/pi-delegate)
- Score: 1
- Comments: 0
- Posted: 2026-10-02T13:03:29Z

## Translation

Title: Show HN: Delegate Claude Code's work to pi and a cheaper model
Article title: GitHub - randomm/pi-delegate: Claude Code skill for delegating development and verification tasks to pi coding agent — with adversarial review loop · GitHub
Description: Claude Code skill for delegating development and verification tasks to pi coding agent — with adversarial review loop - randomm/pi-delegate
HN text: Mini project that enables shipping of some of claude code's work to pi agent running a cheaper model, with up to 33% savings. Though, that really depends on what you are doing with claude!

Article text:
GitHub - randomm/pi-delegate: Claude Code skill for delegating development and verification tasks to pi coding agent — with adversarial review loop · GitHub
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
randomm
/
pi-delegate
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
40 Commits 40 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows .pi .pi bench bench docs docs skills/ delegate skills/ delegate test test AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Claude Code plans and checks. A cheaper model does the heavy lifting.
Say "delegate to pi" and the heavy part of a coding task (reading files, editing, running tests) runs on
pi , an open-source coding agent that can drive any model you pick, even one on your own
machine. Claude only writes the brief and reads a short result, and your tests decide whether it worked.
On multi-file tasks, delegating cut Claude's cost by a third (−33%) with every hidden test still passing.
The honest fine print. The numbers count Claude's cost only: pi's own spend comes on top (nothing if you
self-host, cents on a hosted cheap model). Delegated runs are also slower, several times in our runs on a
self-hosted pi (about 25 s plain vs 2-3 minutes). Samples are small (2 runs per task); the benchmark takes
about five minutes to rerun yourself ( how ).
✅ A good fit
❌ Not a fit
You use Claude Code and your tasks touch several files
Quick one-file edits (Claude alone is cheaper)
You want fewer Claude tokens, or less of your usage limit, spent on routine implementation
Tasks with no way to check the result
You have a test command that can say whether the work is right
Speed matters more than cost
Install (Claude Code)
/plugin marketplace add randomm/pi-delegate
/plugin install pi-delegate@pi-delegate
2. Set up pi once
npm install -g @earendil-works/pi-coding-agent # then run `pi`, /login (or export an API key), /model
Cheap and local models work; the benchmark used a self-hosted Qwen. Also install jq , and on macOS
brew install coreutils for the timeout that bounds each pi call.
delegate to pi: add a --json flag to cli.py, verify with python3 -m unittest
Claude makes one call, pi does the work, and your verify command decides whether it worked:
EXIT CODE: 0
<pi's summary>
VERIFY: PASS (retries=0)
cli.py | 12 +++++++++---
If verification fails, pi gets one more attempt with the failure output. If it still fails, Claude says so
instead of claiming success.
A deterministic gate, not a second opinion. --verify "<your tests>" runs after pi. No model reviews the
work: a same-model reviewer approves most changes, tests don't.
Claude is told not to redo the work when the gate passes. Re-reading the diff and re-running the tests is
exactly what ate the savings in our first benchmark.
Guardrails: it refuses to run on your default branch or next to .env / *.pem / *.key files, and disables
git push for pi. This guards against mistakes, not a malicious model. For stronger isolation set
PI_DELEGATE_WRAP (sandbox recipes in docs/configuration.md ) or use a
disposable clone or container.
skills/delegate/run.sh in this repo is plain bash. Any agent that can run a shell command can use it:
git clone https://github.com/randomm/pi-delegate.git
bash pi-delegate/skills/delegate/run.sh --verify " pytest -q " << ' TASK '
Fix the failing date parsing in utils.py; do not commit.
TASK
Rerun the benchmark
bench/quick.sh -n 3 # about five minutes: plain Claude vs Claude + pi-delegate, prints a REWARD score
Method, tasks and the slower real-repo protocol: docs/benchmark.md ·
results . Settings (timeouts, safety, sandbox): docs/configuration.md .
Apache License 2.0, see LICENSE. Copyright 2026 Janni Turunen.
Claude Code skill for delegating development and verification tasks to pi coding agent — with adversarial review loop
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Claude Code skill for delegating development and verification tasks to pi coding agent — with adversarial review loop - randomm/pi-delegate

Mini project that enables shipping of some of claude code's work to pi agent running a cheaper model, with up to 33% savings. Though, that really depends on what you are doing with claude!

GitHub - randomm/pi-delegate: Claude Code skill for delegating development and verification tasks to pi coding agent — with adversarial review loop · GitHub
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
randomm
/
pi-delegate
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
40 Commits 40 Commits Folders and files
.claude-plugin .claude-plugin .github/ workflows .github/ workflows .pi .pi bench bench docs docs skills/ delegate skills/ delegate test test AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Claude Code plans and checks. A cheaper model does the heavy lifting.
Say "delegate to pi" and the heavy part of a coding task (reading files, editing, running tests) runs on
pi , an open-source coding agent that can drive any model you pick, even one on your own
machine. Claude only writes the brief and reads a short result, and your tests decide whether it worked.
On multi-file tasks, delegating cut Claude's cost by a third (−33%) with every hidden test still passing.
The honest fine print. The numbers count Claude's cost only: pi's own spend comes on top (nothing if you
self-host, cents on a hosted cheap model). Delegated runs are also slower, several times in our runs on a
self-hosted pi (about 25 s plain vs 2-3 minutes). Samples are small (2 runs per task); the benchmark takes
about five minutes to rerun yourself ( how ).
✅ A good fit
❌ Not a fit
You use Claude Code and your tasks touch several files
Quick one-file edits (Claude alone is cheaper)
You want fewer Claude tokens, or less of your usage limit, spent on routine implementation
Tasks with no way to check the result
You have a test command that can say whether the work is right
Speed matters more than cost
Install (Claude Code)
/plugin marketplace add randomm/pi-delegate
/plugin install pi-delegate@pi-delegate
2. Set up pi once
npm install -g @earendil-works/pi-coding-agent # then run `pi`, /login (or export an API key), /model
Cheap and local models work; the benchmark used a self-hosted Qwen. Also install jq , and on macOS
brew install coreutils for the timeout that bounds each pi call.
delegate to pi: add a --json flag to cli.py, verify with python3 -m unittest
Claude makes one call, pi does the work, and your verify command decides whether it worked:
EXIT CODE: 0
<pi's summary>
VERIFY: PASS (retries=0)
cli.py | 12 +++++++++---
If verification fails, pi gets one more attempt with the failure output. If it still fails, Claude says so
instead of claiming success.
A deterministic gate, not a second opinion. --verify "<your tests>" runs after pi. No model reviews the
work: a same-model reviewer approves most changes, tests don't.
Claude is told not to redo the work when the gate passes. Re-reading the diff and re-running the tests is
exactly what ate the savings in our first benchmark.
Guardrails: it refuses to run on your default branch or next to .env / *.pem / *.key files, and disables
git push for pi. This guards against mistakes, not a malicious model. For stronger isolation set
PI_DELEGATE_WRAP (sandbox recipes in docs/configuration.md ) or use a
disposable clone or container.
skills/delegate/run.sh in this repo is plain bash. Any agent that can run a shell command can use it:
git clone https://github.com/randomm/pi-delegate.git
bash pi-delegate/skills/delegate/run.sh --verify " pytest -q " << ' TASK '
Fix the failing date parsing in utils.py; do not commit.
TASK
Rerun the benchmark
bench/quick.sh -n 3 # about five minutes: plain Claude vs Claude + pi-delegate, prints a REWARD score
Method, tasks and the slower real-repo protocol: docs/benchmark.md ·
results . Settings (timeouts, safety, sandbox): docs/configuration.md .
Apache License 2.0, see LICENSE. Copyright 2026 Janni Turunen.
Claude Code skill for delegating development and verification tasks to pi coding agent — with adversarial review loop
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
