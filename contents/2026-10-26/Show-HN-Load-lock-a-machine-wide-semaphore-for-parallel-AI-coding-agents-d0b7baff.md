---
source: "https://github.com/aron98/load-lock"
hn_url: "https://news.ycombinator.com/item?id=50241078"
title: "Show HN: Load-lock, a machine-wide semaphore for parallel AI coding agents"
article_title: "GitHub - aron98/load-lock: Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...) · GitHub"
image: "https://repository-images.githubusercontent.com/1420503912/e1a54ec2-2ae7-4a45-95eb-412ce5b16848"
author: "erdelyia"
captured_at: "2026-10-26T09:47:53Z"
capture_tool: "hn-digest"
hn_id: 50241078
score: 1
comments: 0
posted_at: "2026-10-26T09:08:22Z"
tags:
  - hacker-news
---

# Show HN: Load-lock, a machine-wide semaphore for parallel AI coding agents

- HN: [50241078](https://news.ycombinator.com/item?id=50241078)
- Source: [github.com](https://github.com/aron98/load-lock)
- Score: 1
- Comments: 0
- Posted: 2026-10-26T09:08:22Z

## Translation

Title: Show HN: Load-lock, a machine-wide semaphore for parallel AI coding agents
Article title: GitHub - aron98/load-lock: Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...) · GitHub
Description: Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...) - aron98/load-lock

Article text:
GitHub - aron98/load-lock: Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...) · GitHub
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
aron98
load-lock Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...)
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
118 Commits 118 Commits Folders and files
.github .github assets assets demo demo docker docker docs docs skills/ load-lock skills/ load-lock src src .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A machine-wide heavy-work budget for parallel coding agents.
When several agents work side by side on one machine — separate sessions,
repos, git worktrees, or harnesses (OpenCode, Codex CLI, Claude Code, ...) —
their builds and test suites stomp on each other, thrash the caches, and
starve each other's cores. load-lock gives them one shared counting
semaphore: acquire a permit before heavy work, keep doing light work while
queued, receive a FIFO offer when capacity frees, claim it, run, release.
It is cooperative and advisory — no command wrapping, no automatic shell
classification, no approval interruptions.
agent A (OpenCode) agent B (Codex) agent C (Claude Code)
acquire "cargo test" acquire "npm test" acquire "build docs"
granted ✓ runs queued → keeps coding queued → keeps coding
... offer → claim ✓ runs (misses window → paused)
release ✓ release ✓ rejoin → offer → claim ✓
Three agents, one budget — the queue drains FIFO while everyone keeps coding:
One budget for the whole machine — every client of any harness shares
one capacity pool regardless of repo/worktree/cwd. Default: pool heavy ,
capacity 1; fully runtime-configurable by any agent via versioned CAS
updates (persisted with author + rationale).
Atomic acquire-or-queue — single daemon event loop; a release can
never slip between a failed check and queue registration.
FIFO offers with opaque claim tokens — outstanding offers reserve
capacity; expired offers pass capacity to the next waiter and pause the
absent request until it explicitly rejoins (no endless recycling).
Real process-lifecycle cleanup — permits attach to actual processes
(pid + process group + /proc start time, PID-reuse-safe). Background
jobs outliving the shell keep their permit until the whole group exits.
Never auto-expired merely for being old. Untracked work uses an honest,
documented heartbeat-lease tradeoff. Explicit release is always
available and idempotent.
Crash-safe — agent crashes release tracked permits on confirmed
process death; daemon restarts reconcile a journal against real liveness
and auto-respawn transparently.
Honest notifications — tool-boundary offer delivery everywhere; real
idle wakeup where the harness supports it (live-verified on OpenCode;
hook-based on Claude Code and Codex CLI); a bounded blocking
wait_for_turn ( load-lock wait on the CLI) that sleeps until your
turn and auto-claims for harnesses without idle wakeup; durable event
ids for reconnects. A notification never authorizes work — only
claim does.
Policy that actually reaches the model — installed per harness via
its real mechanism (plugin context hook, SessionStart hooks incl.
compact matchers, global instruction files), plus a detailed skill/guide.
Resource-aware — inspect_machine reports cgroup/container CPU and
memory limits and current pressure; never assumes cores = capacity.
npm install -g @aron98/load-lock
load-lock install --harness opencode # or claude-code | codex | gemini-cli |
# copilot-cli | cline | roo-code | goose |
# omac (sandbox grants) | generic
load-lock doctor
From then on, agents see the policy and tools automatically:
OpenCode plugin ( load_lock_* tools), Claude Code / Codex MCP tools
( mcp__load-lock__* ) plus session hooks. Any other MCP-capable harness can
use the generic stdio server; plain shells use the CLI:
load-lock run --reason " nightly build " -- npm run build # acquire→run→release
load-lock status # machine-wide view
load-lock dashboard # live read-only dashboard (q quits)
load-lock help # one-screen usage of every command
load-lock configure --capacity 2 --reason " 16-core host " --author me
OpenCode users also get a live dashboard inside the TUI: /load-lock
opens it in-window, every session screen shows a clickable at-a-glance
widget in the right sidebar, and the home screen has a footer indicator —
same data, themed. See docs/OPERATIONS.md → "Dashboard".
Agents running inside a sandbox (bubblewrap/omac, Docker, macOS seatbelt,
WSL2) share the host's budget instead of getting their own: bridge the
daemon's Unix socket into the sandbox, set LOAD_LOCK_NO_SPAWN=1 so the
client never spawns a useless sandbox-local daemon, and optionally tag
holds with LOAD_LOCK_LABEL_PREFIX (e.g. bwrap:issue-42 build ).
Ready-to-run per-sandbox recipes — live-verified for bubblewrap and omac,
contract-inspected for the rest, plus the cross-PID-namespace lease
lifecycle — live in docs/OPERATIONS.md → "Sandboxes".
src/core/ daemon + state machine + persistence + machine inspection
src/mcp/ MCP stdio server (universal adapter)
src/adapters/ OpenCode native plugin
src/cli/ CLI incl. installer, hook handlers, run/wait helpers
src/test/ 32 tests: state machine, real multi-process daemon tests,
MCP integration, hook contracts, run helper
demo/ reproducible 3-worktree demo (three agents, one pool)
docs/ ARCHITECTURE / COMPATIBILITY / OPERATIONS / AGENT_GUIDE
skills/ packaged skill
Documentation
docs/ARCHITECTURE.md — design, why a dedicated
service (vs WorkOS Airlock / mic-lock / flock wrappers), atomicity,
restart semantics, notification honesty, lifecycle tradeoffs.
docs/COMPATIBILITY.md — per-harness capability
matrix with implemented / contract-tested / proposed labels.
docs/OPERATIONS.md — install, config,
troubleshooting, recovery, daemon lifecycle, sandbox bridging.
docs/AGENT_GUIDE.md — the playbook for agents.
docs/ADAPTERS.md — how to write an adapter for a new
harness, using the published library surface.
Core daemon/protocol/CLI: implemented, 68 automated tests passing
(multi-process daemons, real processes, restarts, races, expiry,
PID-reuse, subprocess survival; plus 9 dashboard tests — view-model
shaping, event-cursor resume across a daemon restart, reconnect, and a
load-lock dashboard --once render smoke test).
OpenCode adapter: implemented and live-verified (tools, policy
injection, compaction-proof context hook, execution-boundary wakeup
with tokens, skill).
Claude Code adapter: implemented, contract-tested against documented
hook JSON; not yet live-verified (harness not installed on the
verification machine).
Codex CLI adapter: implemented and live-verified end-to-end in
Docker (interactive TUI via pty: policy injection, MCP round-trip,
FIFO queueing behind a competing holder, Stop-hook decision-block
wakeup; see COMPATIBILITY.md for the two upstream/provider constraints
the verification harness works around).
Tier-3 fallback installers (Gemini CLI, Copilot CLI, Cline, Roo Code,
Goose): implemented merge-only installers (MCP stdio server entry into
the harness's real config + a global instructions-file section/policy
rule file), config-tested via temp-HOME round-trips; fallback tier (no
hooks, no async push — no idle wakeup).
Cursor adapter: implemented (issue #6) — merge-only project-level installer
( .cursor/mcp.json MCP entry, .cursor/hooks.json stop hook per the now-
officially-documented hooks v1 contract, .cursor/rules/load-lock.mdc
Always-apply rule) + hook cursor stop handler emitting
{"followup_message":...} idle wakeup. Contract-tested , not
live-verified (no Cursor on the verification machine).
Aider: researched; manual CLI tier. See COMPATIBILITY.md for the matrix.
Linux-first; macOS works with documented liveness-check limitations;
Windows has a named-pipe transport adapter (implemented and unit-tested on
Linux via platform injection; not yet verified on a Windows machine).
Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...)
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...) - aron98/load-lock

GitHub - aron98/load-lock: Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...) · GitHub
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
aron98
load-lock Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...)
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
118 Commits 118 Commits Folders and files
.github .github assets assets demo demo docker docker docs docs skills/ load-lock skills/ load-lock src src .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A machine-wide heavy-work budget for parallel coding agents.
When several agents work side by side on one machine — separate sessions,
repos, git worktrees, or harnesses (OpenCode, Codex CLI, Claude Code, ...) —
their builds and test suites stomp on each other, thrash the caches, and
starve each other's cores. load-lock gives them one shared counting
semaphore: acquire a permit before heavy work, keep doing light work while
queued, receive a FIFO offer when capacity frees, claim it, run, release.
It is cooperative and advisory — no command wrapping, no automatic shell
classification, no approval interruptions.
agent A (OpenCode) agent B (Codex) agent C (Claude Code)
acquire "cargo test" acquire "npm test" acquire "build docs"
granted ✓ runs queued → keeps coding queued → keeps coding
... offer → claim ✓ runs (misses window → paused)
release ✓ release ✓ rejoin → offer → claim ✓
Three agents, one budget — the queue drains FIFO while everyone keeps coding:
One budget for the whole machine — every client of any harness shares
one capacity pool regardless of repo/worktree/cwd. Default: pool heavy ,
capacity 1; fully runtime-configurable by any agent via versioned CAS
updates (persisted with author + rationale).
Atomic acquire-or-queue — single daemon event loop; a release can
never slip between a failed check and queue registration.
FIFO offers with opaque claim tokens — outstanding offers reserve
capacity; expired offers pass capacity to the next waiter and pause the
absent request until it explicitly rejoins (no endless recycling).
Real process-lifecycle cleanup — permits attach to actual processes
(pid + process group + /proc start time, PID-reuse-safe). Background
jobs outliving the shell keep their permit until the whole group exits.
Never auto-expired merely for being old. Untracked work uses an honest,
documented heartbeat-lease tradeoff. Explicit release is always
available and idempotent.
Crash-safe — agent crashes release tracked permits on confirmed
process death; daemon restarts reconcile a journal against real liveness
and auto-respawn transparently.
Honest notifications — tool-boundary offer delivery everywhere; real
idle wakeup where the harness supports it (live-verified on OpenCode;
hook-based on Claude Code and Codex CLI); a bounded blocking
wait_for_turn ( load-lock wait on the CLI) that sleeps until your
turn and auto-claims for harnesses without idle wakeup; durable event
ids for reconnects. A notification never authorizes work — only
claim does.
Policy that actually reaches the model — installed per harness via
its real mechanism (plugin context hook, SessionStart hooks incl.
compact matchers, global instruction files), plus a detailed skill/guide.
Resource-aware — inspect_machine reports cgroup/container CPU and
memory limits and current pressure; never assumes cores = capacity.
npm install -g @aron98/load-lock
load-lock install --harness opencode # or claude-code | codex | gemini-cli |
# copilot-cli | cline | roo-code | goose |
# omac (sandbox grants) | generic
load-lock doctor
From then on, agents see the policy and tools automatically:
OpenCode plugin ( load_lock_* tools), Claude Code / Codex MCP tools
( mcp__load-lock__* ) plus session hooks. Any other MCP-capable harness can
use the generic stdio server; plain shells use the CLI:
load-lock run --reason " nightly build " -- npm run build # acquire→run→release
load-lock status # machine-wide view
load-lock dashboard # live read-only dashboard (q quits)
load-lock help # one-screen usage of every command
load-lock configure --capacity 2 --reason " 16-core host " --author me
OpenCode users also get a live dashboard inside the TUI: /load-lock
opens it in-window, every session screen shows a clickable at-a-glance
widget in the right sidebar, and the home screen has a footer indicator —
same data, themed. See docs/OPERATIONS.md → "Dashboard".
Agents running inside a sandbox (bubblewrap/omac, Docker, macOS seatbelt,
WSL2) share the host's budget instead of getting their own: bridge the
daemon's Unix socket into the sandbox, set LOAD_LOCK_NO_SPAWN=1 so the
client never spawns a useless sandbox-local daemon, and optionally tag
holds with LOAD_LOCK_LABEL_PREFIX (e.g. bwrap:issue-42 build ).
Ready-to-run per-sandbox recipes — live-verified for bubblewrap and omac,
contract-inspected for the rest, plus the cross-PID-namespace lease
lifecycle — live in docs/OPERATIONS.md → "Sandboxes".
src/core/ daemon + state machine + persistence + machine inspection
src/mcp/ MCP stdio server (universal adapter)
src/adapters/ OpenCode native plugin
src/cli/ CLI incl. installer, hook handlers, run/wait helpers
src/test/ 32 tests: state machine, real multi-process daemon tests,
MCP integration, hook contracts, run helper
demo/ reproducible 3-worktree demo (three agents, one pool)
docs/ ARCHITECTURE / COMPATIBILITY / OPERATIONS / AGENT_GUIDE
skills/ packaged skill
Documentation
docs/ARCHITECTURE.md — design, why a dedicated
service (vs WorkOS Airlock / mic-lock / flock wrappers), atomicity,
restart semantics, notification honesty, lifecycle tradeoffs.
docs/COMPATIBILITY.md — per-harness capability
matrix with implemented / contract-tested / proposed labels.
docs/OPERATIONS.md — install, config,
troubleshooting, recovery, daemon lifecycle, sandbox bridging.
docs/AGENT_GUIDE.md — the playbook for agents.
docs/ADAPTERS.md — how to write an adapter for a new
harness, using the published library surface.
Core daemon/protocol/CLI: implemented, 68 automated tests passing
(multi-process daemons, real processes, restarts, races, expiry,
PID-reuse, subprocess survival; plus 9 dashboard tests — view-model
shaping, event-cursor resume across a daemon restart, reconnect, and a
load-lock dashboard --once render smoke test).
OpenCode adapter: implemented and live-verified (tools, policy
injection, compaction-proof context hook, execution-boundary wakeup
with tokens, skill).
Claude Code adapter: implemented, contract-tested against documented
hook JSON; not yet live-verified (harness not installed on the
verification machine).
Codex CLI adapter: implemented and live-verified end-to-end in
Docker (interactive TUI via pty: policy injection, MCP round-trip,
FIFO queueing behind a competing holder, Stop-hook decision-block
wakeup; see COMPATIBILITY.md for the two upstream/provider constraints
the verification harness works around).
Tier-3 fallback installers (Gemini CLI, Copilot CLI, Cline, Roo Code,
Goose): implemented merge-only installers (MCP stdio server entry into
the harness's real config + a global instructions-file section/policy
rule file), config-tested via temp-HOME round-trips; fallback tier (no
hooks, no async push — no idle wakeup).
Cursor adapter: implemented (issue #6) — merge-only project-level installer
( .cursor/mcp.json MCP entry, .cursor/hooks.json stop hook per the now-
officially-documented hooks v1 contract, .cursor/rules/load-lock.mdc
Always-apply rule) + hook cursor stop handler emitting
{"followup_message":...} idle wakeup. Contract-tested , not
live-verified (no Cursor on the verification machine).
Aider: researched; manual CLI tier. See COMPATIBILITY.md for the matrix.
Linux-first; macOS works with documented liveness-check limitations;
Windows has a named-pipe transport adapter (implemented and unit-tested on
Linux via platform injection; not yet verified on a Windows machine).
Machine-wide counting semaphore coordinating heavy work across parallel coding agents (OpenCode, Codex CLI, Claude Code, ...)
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
