---
source: "https://github.com/yanacuti1121/yana-ai"
hn_url: "https://news.ycombinator.com/item?id=50160774"
title: "Show HN: Yana – A local-first, human-governed AI runtime and IDE in Rust"
article_title: "GitHub - yanacuti1121/Yana-AI: Yana AI 🐰 One runtime. Any AI. Human-governed. A local-first, cross-platform system for running, connecting, orchestrating, and governing AI — with deterministic control over what it can access, change, and execute. Your AI can act. But who decides how far it can go? ·\n[truncated]"
image: "https://opengraph.githubassets.com/4d3930ab43b9352b5600dadb0b9a1744822e621c407109f0584c2c2a474b0589/yanacuti1121/Yana-AI"
author: "vutam"
captured_at: "2026-10-20T03:41:44Z"
capture_tool: "hn-digest"
hn_id: 50160774
score: 3
comments: 0
posted_at: "2026-10-20T02:48:40Z"
tags:
  - hacker-news
---

# Show HN: Yana – A local-first, human-governed AI runtime and IDE in Rust

- HN: [50160774](https://news.ycombinator.com/item?id=50160774)
- Source: [github.com](https://github.com/yanacuti1121/yana-ai)
- Score: 3
- Comments: 0
- Posted: 2026-10-20T02:48:40Z

## Translation

Title: Show HN: Yana – A local-first, human-governed AI runtime and IDE in Rust
Article title: GitHub - yanacuti1121/Yana-AI: Yana AI 🐰 One runtime. Any AI. Human-governed. A local-first, cross-platform system for running, connecting, orchestrating, and governing AI — with deterministic control over what it can access, change, and execute. Your AI can act. But who decides how far it can go? ·
[truncated]
Description: Yana AI 🐰 One runtime. Any AI. Human-governed. A local-first, cross-platform system for running, connecting, orchestrating, and governing AI — with deterministic control over what it can access, change, and execute. Your AI can act. But who decides how far it can go? - yanacuti1121/Yana-AI

Article text:
GitHub - yanacuti1121/Yana-AI: Yana AI 🐰 One runtime. Any AI. Human-governed. A local-first, cross-platform system for running, connecting, orchestrating, and governing AI — with deterministic control over what it can access, change, and execute. Your AI can act. But who decides how far it can go? · GitHub
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
yanacuti1121
Yana-AI Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Yana AI 🐰 One runtime. Any AI. Human-governed. A local-first, cross-platform system for running, connecting, orchestrating, and governing AI — with deterministic control over what it can access, change, and execute. Your AI can act. But who decides how far it can go?
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2,355 Commits 2,355 Commits Folders and files
.agent/ rules .agent/ rules .claude-plugin .claude-plugin .claude .claude .codex .codex .cursor .cursor .gemini .gemini .github .github .out-of-scope .out-of-scope .planning/ yana-web .planning/ yana-web .vscode .vscode .yana-ai .yana-ai .yana .yana adapters adapters bin bin ci ci core core demo demo docs docs examples examples gates gates github-app github-app guards guards ledger ledger memory memory ops ops policy policy prompts prompts releases releases reports reports router router scanner scanner scripts scripts site site skills/ yana-ai skills/ yana-ai src src tests tests tools tools vendor vendor .cursorrules .cursorrules .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .npmignore .npmignore .yana-aiignore .yana-aiignore .yana-aiignore.example .yana-aiignore.example ACKNOWLEDGEMENTS.ko.md ACKNOWLEDGEMENTS.ko.md ACKNOWLEDGEMENTS.md ACKNOWLEDGEMENTS.md ACKNOWLEDGEMENTS.vi.md ACKNOWLEDGEMENTS.vi.md AGENTS.md AGENTS.md ARCHITECTURE-DEBT.md ARCHITECTURE-DEBT.md ARCHITECTURE.md ARCHITECTURE.md BENCHMARK.md BENCHMARK.md BMAD-YANA-CAPABILITY-MATRIX.md BMAD-YANA-CAPABILITY-MATRIX.md BMAD-YANA-GAP-ANALYSIS.md BMAD-YANA-GAP-ANALYSIS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMMANDS.md COMMANDS.md CONTRIBUTING.md CONTRIBUTING.md CURRENT-MILESTONE.md CURRENT-MILESTONE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DIRECTION.md DIRECTION.md HANDOFF-CODEX-DESKTOP-PHASE9-35.md HANDOFF-CODEX-DESKTOP-PHASE9-35.md HANDOFF-TO-CLAUDE-DESKTOP.md HANDOFF-TO-CLAUDE-DESKTOP.md HANDOFF-TO-CODEX-ECOSYSTEM-WEBSITE.md HANDOFF-TO-CODEX-ECOSYSTEM-WEBSITE.md HANDOFF-TO-CODEX-STUDIO-ARCHITECTURE.md HANDOFF-TO-CODEX-STUDIO-ARCHITECTURE.md HISTORY.md HISTORY.md JOURNEY.ko.md JOURNEY.ko.md JOURNEY.md JOURNEY.md JOURNEY.vi.md JOURNEY.vi.md LICENSE LICENSE MANIFEST.json MANIFEST.json PHILOSOPHY.ko.md PHILOSOPHY.ko.md PHILOSOPHY.md PHILOSOPHY.md PHILOSOPHY.vi.md PHILOSOPHY.vi.md PRINCIPLES.ko.md PRINCIPLES.ko.md PRIN
[truncated]
English ·
Tiếng Việt ·
한국어 ·
中文
Created by Vũ Văn Tâm · Vietnam
Models propose. Yana decides. Humans stay in charge.
Yana is the control plane between an AI and real power. A model can say what it
wants to do. Yana decides whether that is allowed, and a person can always stop it.
Model proposes an action
|
Authority is this agent allowed to ask?
|
Capability what exact power is this? (a small, fixed set)
|
Policy may it happen now? rules first, then a human
|
Executor performs only the permitted operation
Any model goes through the same path: Claude, GPT, Gemini, DeepSeek, or one
running on your own machine. Skills teach an agent how to work; only a
capability can grant permission to act.
Component
What it does
Where
Status
Yana Studio
Desktop workspace (files, editor, terminals, chat, tasks) for working with AI on a real project.
tools/yana-studio
Live, partial parity
Runtime ( yana-rt )
Rust engine: one turn loop and one authority contract for local and cloud models, 19 providers.
src/
Live
Governance
Capability manifests, human approval, hash-chained audit log, core-lock integrity.
core/
Live
Giám Thị
Independent watcher outside any AI session. Halts every later tool call when it finds a problem.
core/scripts
Live, opt-in
Harness adapters
Claude Code, Codex, Cursor and Antigravity. Enforcement strength differs per host (see below).
adapters/
Live, varies by host
Knowledge layer
Agents, skills, commands and rules that teach agents how to work. They grant no permissions.
core/agents , core/skills
Live
Yana OS
Local management plane: agent identity, autonomy levels, quotas, health, quarantine.
src/os
In development
Yana Desktop / Yana Web
Earlier Electron shell and web UI. Studio is a separate rebuild, not a reskin of these.
tools/yana-desktop , tools/yana-web
Maintained separately
Yana Wheelbot
Robotics platform on ESP32-S3, in its own repository.
yana-wheelbot
Experimental
Live means shipped and in use. In development means real code exists but is
not feature-complete. Experimental means opt-in and not independently verified.
yana-rt has 40 subcommands. Zero Python dependency. This is the source-defined count across feature builds, including feature-gated ones.
core/
├── hooks/ # 66 PreToolUse / PostToolUse / Stop hooks (63 shell, 3 JavaScript)
├── rules/ # 71 rules, most are guidance the agent applies; hooks enforce a subset
├── agents/ # 100 specialist agents
├── skills/ # 2,026 skills
├── commands/ # 170 Claude Code slash commands
└── config/
└── core-lock.json # SHA-256 manifest — 287 core files pinned
These counts are derived from the filesystem by core/scripts/check_counts.py , which
CI runs, so they should not drift from the tree.
One window for the things you switch between when working with AI on a project.
Studio talks to the model only through yana-rt ; it never calls a provider from
the UI layer.
Files and editor: open a real folder, search filenames across it, edit with CodeMirror. Files up to 8 MiB are editable; larger ones open as read-only 256 KiB windows. Saves detect outside changes before overwriting.
Terminals: real PTY sessions in tabs, or a grid of up to 12.
Git: read-only status, worktrees and diffs. Stage and commit are deliberately not AI tools.
Chat: streamed through yana-rt , multiple saved conversations, stop button, approval prompts, token usage.
Models: all 19 providers (13 cloud, 6 local) plus a custom OpenAI-compatible endpoint. A local manager detects 9Router, Ollama, LM Studio, llama.cpp, TurboFieldfare and AirLLM. API keys go to Rust over stdin and are stored encrypted with the OS keystore, never in workspace files.
Tasks: create, list, depend and drop tasks against the real yana-rt task store. Marking done requires evidence, and blocked is always computed by the runtime.
Design canvas: native canvas with 22 components, phone and desktop screens, layers, undo/redo, and guarded AI proposals.
Governance telemetry: what was allowed, what was blocked, and why.
Not built yet: device registry, agent sessions screen, mobile pairing, SSH/Docker
execution, permission-lease UI, signed auto-update. The per-screen status is in
tools/yana-studio/SCREEN-STATUS.md .
Download from the latest release
or yana.vutam.link .
Each release lists its own assets, so check that your platform is there, and verify
against SHA256SUMS when provided. There is no Linux ARM build. The macOS build is
ad-hoc signed and not notarized (no Apple Developer membership yet), so Gatekeeper
warns on first launch; docs/MACOS_INSTALL.md shows the two
official ways to open it.
Build from source (Node 24+ and a C++ toolchain for node-pty ):
cd tools/yana-studio
npm ci && npm run rebuild
npm run stage:runtime -- /absolute/path/to/yana-rt
npm start
Runtime: yana-rt
A single Rust binary that owns the turn loop and the authority check for every client,
so Studio, the CLI and MCP do not each reimplement permission logic.
19 providers: anthropic, openai, gemini, groq, deepseek, openrouter, xai, novita, nvidia, minimax, glm, huggingface, 9router, kimi (cloud), and ollama, lmstudio, llamacpp, turbofieldfare, airllm (local). The canonical list is src/model/catalog.rs ; a contract test keeps Studio in sync with it.
Failover: typed provider errors, fallback providers and credential-pool rotation.
State: SQLite session store with full-text search, per-profile state, crash recovery, and shadow-git checkpoints with rollback.
Capability tools: file read/patch, command execution under validation, MCP client over stdio, web.search through a checked HTTPS layer, and a read-only LSP client.
Optional features: mcp , lsp , pty-bridge , Discord remote adapter. They are off in the default build.
Remaining work on the runtime is tracked in CURRENT-MILESTONE.md .
At the time of writing: unifying every client on one capability executor, and running
local-model tool calling against a real local model rather than a mock.
An independent watcher that runs outside any AI session, from the operating system's
own scheduler: launchd on macOS, a systemd timer on Linux, Task Scheduler on Windows,
every 6 hours.
It compares the repo against a baseline commit a human approved, over the files that control enforcement: hook settings for Claude Code, Codex and Cursor, CI workflows, core/rules , core/gates , core/hooks , core/scripts and src/guard .
On drift it writes GIAMTHI_HALT.lock . A hook that runs first on every tool call then denies everything while the lock exists.
Only a person deleting the lock clears it. Neither the hook, the watcher nor any environment variable can.
Manage it with yana-ai giamthi install | status | repair | uninstall .
It is opt-in. Because it runs on a schedule, there is a window of up to 6 hours between a
change and the halt. Codex fires PreToolUse only for shell commands, so Yana also
checks the lock on SessionStart and UserPromptSubmit for Codex.
Host
How governance is applied
Enforcement
Claude Code
Native hooks and rules in .claude/
Runtime blocking on every tool call, hash-chained audit
Codex
AGENTS.md , .codex/ hooks, synced skills and agents ( yana-ai install --engine codex )
Project-scoped hooks
Cursor
.cursorrules , .cursor/rules , a beforeShellExecution hook
Shell commands screened; MCP tool calls are not covered
Antigravity
.agent/rules/yana-ai.md
Advisory only, not enforced at the shell
Details and switching commands: adapters/README.md .
pip install yana-ai # Python CLI
yana-ai install # add Yana's hooks and rules to your repo
yana-ai doctor . # check that everything is wired
cargo install yana-rt # native runtime, no Python needed
Needs Python 3.11+ (or Rust), Git, and one supported harness. All CLI commands are in
COMMANDS.md . The 170 Claude Code slash commands ( /plan , /verify ,
/code-review and others) live in core/commands/ and

[truncated]

## Original Extract

Yana AI 🐰 One runtime. Any AI. Human-governed. A local-first, cross-platform system for running, connecting, orchestrating, and governing AI — with deterministic control over what it can access, change, and execute. Your AI can act. But who decides how far it can go? - yanacuti1121/Yana-AI

GitHub - yanacuti1121/Yana-AI: Yana AI 🐰 One runtime. Any AI. Human-governed. A local-first, cross-platform system for running, connecting, orchestrating, and governing AI — with deterministic control over what it can access, change, and execute. Your AI can act. But who decides how far it can go? · GitHub
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
yanacuti1121
Yana-AI Public Notifications You must be signed in to change notification settings
Star 4 ( 4 ) You must be signed in to star a repository
Yana AI 🐰 One runtime. Any AI. Human-governed. A local-first, cross-platform system for running, connecting, orchestrating, and governing AI — with deterministic control over what it can access, change, and execute. Your AI can act. But who decides how far it can go?
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2,355 Commits 2,355 Commits Folders and files
.agent/ rules .agent/ rules .claude-plugin .claude-plugin .claude .claude .codex .codex .cursor .cursor .gemini .gemini .github .github .out-of-scope .out-of-scope .planning/ yana-web .planning/ yana-web .vscode .vscode .yana-ai .yana-ai .yana .yana adapters adapters bin bin ci ci core core demo demo docs docs examples examples gates gates github-app github-app guards guards ledger ledger memory memory ops ops policy policy prompts prompts releases releases reports reports router router scanner scanner scripts scripts site site skills/ yana-ai skills/ yana-ai src src tests tests tools tools vendor vendor .cursorrules .cursorrules .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .gitmodules .gitmodules .npmignore .npmignore .yana-aiignore .yana-aiignore .yana-aiignore.example .yana-aiignore.example ACKNOWLEDGEMENTS.ko.md ACKNOWLEDGEMENTS.ko.md ACKNOWLEDGEMENTS.md ACKNOWLEDGEMENTS.md ACKNOWLEDGEMENTS.vi.md ACKNOWLEDGEMENTS.vi.md AGENTS.md AGENTS.md ARCHITECTURE-DEBT.md ARCHITECTURE-DEBT.md ARCHITECTURE.md ARCHITECTURE.md BENCHMARK.md BENCHMARK.md BMAD-YANA-CAPABILITY-MATRIX.md BMAD-YANA-CAPABILITY-MATRIX.md BMAD-YANA-GAP-ANALYSIS.md BMAD-YANA-GAP-ANALYSIS.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMMANDS.md COMMANDS.md CONTRIBUTING.md CONTRIBUTING.md CURRENT-MILESTONE.md CURRENT-MILESTONE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DIRECTION.md DIRECTION.md HANDOFF-CODEX-DESKTOP-PHASE9-35.md HANDOFF-CODEX-DESKTOP-PHASE9-35.md HANDOFF-TO-CLAUDE-DESKTOP.md HANDOFF-TO-CLAUDE-DESKTOP.md HANDOFF-TO-CODEX-ECOSYSTEM-WEBSITE.md HANDOFF-TO-CODEX-ECOSYSTEM-WEBSITE.md HANDOFF-TO-CODEX-STUDIO-ARCHITECTURE.md HANDOFF-TO-CODEX-STUDIO-ARCHITECTURE.md HISTORY.md HISTORY.md JOURNEY.ko.md JOURNEY.ko.md JOURNEY.md JOURNEY.md JOURNEY.vi.md JOURNEY.vi.md LICENSE LICENSE MANIFEST.json MANIFEST.json PHILOSOPHY.ko.md PHILOSOPHY.ko.md PHILOSOPHY.md PHILOSOPHY.md PHILOSOPHY.vi.md PHILOSOPHY.vi.md PRINCIPLES.ko.md PRINCIPLES.ko.md PRIN
[truncated]
English ·
Tiếng Việt ·
한국어 ·
中文
Created by Vũ Văn Tâm · Vietnam
Models propose. Yana decides. Humans stay in charge.
Yana is the control plane between an AI and real power. A model can say what it
wants to do. Yana decides whether that is allowed, and a person can always stop it.
Model proposes an action
|
Authority is this agent allowed to ask?
|
Capability what exact power is this? (a small, fixed set)
|
Policy may it happen now? rules first, then a human
|
Executor performs only the permitted operation
Any model goes through the same path: Claude, GPT, Gemini, DeepSeek, or one
running on your own machine. Skills teach an agent how to work; only a
capability can grant permission to act.
Component
What it does
Where
Status
Yana Studio
Desktop workspace (files, editor, terminals, chat, tasks) for working with AI on a real project.
tools/yana-studio
Live, partial parity
Runtime ( yana-rt )
Rust engine: one turn loop and one authority contract for local and cloud models, 19 providers.
src/
Live
Governance
Capability manifests, human approval, hash-chained audit log, core-lock integrity.
core/
Live
Giám Thị
Independent watcher outside any AI session. Halts every later tool call when it finds a problem.
core/scripts
Live, opt-in
Harness adapters
Claude Code, Codex, Cursor and Antigravity. Enforcement strength differs per host (see below).
adapters/
Live, varies by host
Knowledge layer
Agents, skills, commands and rules that teach agents how to work. They grant no permissions.
core/agents , core/skills
Live
Yana OS
Local management plane: agent identity, autonomy levels, quotas, health, quarantine.
src/os
In development
Yana Desktop / Yana Web
Earlier Electron shell and web UI. Studio is a separate rebuild, not a reskin of these.
tools/yana-desktop , tools/yana-web
Maintained separately
Yana Wheelbot
Robotics platform on ESP32-S3, in its own repository.
yana-wheelbot
Experimental
Live means shipped and in use. In development means real code exists but is
not feature-complete. Experimental means opt-in and not independently verified.
yana-rt has 40 subcommands. Zero Python dependency. This is the source-defined count across feature builds, including feature-gated ones.
core/
├── hooks/ # 66 PreToolUse / PostToolUse / Stop hooks (63 shell, 3 JavaScript)
├── rules/ # 71 rules, most are guidance the agent applies; hooks enforce a subset
├── agents/ # 100 specialist agents
├── skills/ # 2,026 skills
├── commands/ # 170 Claude Code slash commands
└── config/
└── core-lock.json # SHA-256 manifest — 287 core files pinned
These counts are derived from the filesystem by core/scripts/check_counts.py , which
CI runs, so they should not drift from the tree.
One window for the things you switch between when working with AI on a project.
Studio talks to the model only through yana-rt ; it never calls a provider from
the UI layer.
Files and editor: open a real folder, search filenames across it, edit with CodeMirror. Files up to 8 MiB are editable; larger ones open as read-only 256 KiB windows. Saves detect outside changes before overwriting.
Terminals: real PTY sessions in tabs, or a grid of up to 12.
Git: read-only status, worktrees and diffs. Stage and commit are deliberately not AI tools.
Chat: streamed through yana-rt , multiple saved conversations, stop button, approval prompts, token usage.
Models: all 19 providers (13 cloud, 6 local) plus a custom OpenAI-compatible endpoint. A local manager detects 9Router, Ollama, LM Studio, llama.cpp, TurboFieldfare and AirLLM. API keys go to Rust over stdin and are stored encrypted with the OS keystore, never in workspace files.
Tasks: create, list, depend and drop tasks against the real yana-rt task store. Marking done requires evidence, and blocked is always computed by the runtime.
Design canvas: native canvas with 22 components, phone and desktop screens, layers, undo/redo, and guarded AI proposals.
Governance telemetry: what was allowed, what was blocked, and why.
Not built yet: device registry, agent sessions screen, mobile pairing, SSH/Docker
execution, permission-lease UI, signed auto-update. The per-screen status is in
tools/yana-studio/SCREEN-STATUS.md .
Download from the latest release
or yana.vutam.link .
Each release lists its own assets, so check that your platform is there, and verify
against SHA256SUMS when provided. There is no Linux ARM build. The macOS build is
ad-hoc signed and not notarized (no Apple Developer membership yet), so Gatekeeper
warns on first launch; docs/MACOS_INSTALL.md shows the two
official ways to open it.
Build from source (Node 24+ and a C++ toolchain for node-pty ):
cd tools/yana-studio
npm ci && npm run rebuild
npm run stage:runtime -- /absolute/path/to/yana-rt
npm start
Runtime: yana-rt
A single Rust binary that owns the turn loop and the authority check for every client,
so Studio, the CLI and MCP do not each reimplement permission logic.
19 providers: anthropic, openai, gemini, groq, deepseek, openrouter, xai, novita, nvidia, minimax, glm, huggingface, 9router, kimi (cloud), and ollama, lmstudio, llamacpp, turbofieldfare, airllm (local). The canonical list is src/model/catalog.rs ; a contract test keeps Studio in sync with it.
Failover: typed provider errors, fallback providers and credential-pool rotation.
State: SQLite session store with full-text search, per-profile state, crash recovery, and shadow-git checkpoints with rollback.
Capability tools: file read/patch, command execution under validation, MCP client over stdio, web.search through a checked HTTPS layer, and a read-only LSP client.
Optional features: mcp , lsp , pty-bridge , Discord remote adapter. They are off in the default build.
Remaining work on the runtime is tracked in CURRENT-MILESTONE.md .
At the time of writing: unifying every client on one capability executor, and running
local-model tool calling against a real local model rather than a mock.
An independent watcher that runs outside any AI session, from the operating system's
own scheduler: launchd on macOS, a systemd timer on Linux, Task Scheduler on Windows,
every 6 hours.
It compares the repo against a baseline commit a human approved, over the files that control enforcement: hook settings for Claude Code, Codex and Cursor, CI workflows, core/rules , core/gates , core/hooks , core/scripts and src/guard .
On drift it writes GIAMTHI_HALT.lock . A hook that runs first on every tool call then denies everything while the lock exists.
Only a person deleting the lock clears it. Neither the hook, the watcher nor any environment variable can.
Manage it with yana-ai giamthi install | status | repair | uninstall .
It is opt-in. Because it runs on a schedule, there is a window of up to 6 hours between a
change and the halt. Codex fires PreToolUse only for shell commands, so Yana also
checks the lock on SessionStart and UserPromptSubmit for Codex.
Host
How governance is applied
Enforcement
Claude Code
Native hooks and rules in .claude/
Runtime blocking on every tool call, hash-chained audit
Codex
AGENTS.md , .codex/ hooks, synced skills and agents ( yana-ai install --engine codex )
Project-scoped hooks
Cursor
.cursorrules , .cursor/rules , a beforeShellExecution hook
Shell commands screened; MCP tool calls are not covered
Antigravity
.agent/rules/yana-ai.md
Advisory only, not enforced at the shell
Details and switching commands: adapters/README.md .
pip install yana-ai # Python CLI
yana-ai install # add Yana's hooks and rules to your repo
yana-ai doctor . # check that everything is wired
cargo install yana-rt # native runtime, no Python needed
Needs Python 3.11+ (or Rust), Git, and one supported harness. All CLI commands are in
COMMANDS.md . The 170 Claude Code slash commands ( /plan , /verify ,
/code-review and others) live in core/commands/ and

[truncated]
