---
source: "https://github.com/purser-sh/purser"
hn_url: "https://news.ycombinator.com/item?id=49678437"
title: "Show HN: Purser – approve, audit and meter your AI coding agents"
article_title: "GitHub - purser-sh/purser: Control layer for AI coding agents — approve every change before it lands, audit what happened, meter what it cost. Local-first, Apache-2.0. · GitHub"
image: "https://repository-images.githubusercontent.com/1347929371/b394189c-32c8-4f60-a3a7-439960eb29f6"
author: "aryansingh9034"
captured_at: "2026-09-12T23:47:37Z"
capture_tool: "hn-digest"
hn_id: 49678437
score: 1
comments: 0
posted_at: "2026-09-12T23:46:06Z"
tags:
  - hacker-news
---

# Show HN: Purser – approve, audit and meter your AI coding agents

- HN: [49678437](https://news.ycombinator.com/item?id=49678437)
- Source: [github.com](https://github.com/purser-sh/purser)
- Score: 1
- Comments: 0
- Posted: 2026-09-12T23:46:06Z

## Translation

Title: Show HN: Purser – approve, audit and meter your AI coding agents
Article title: GitHub - purser-sh/purser: Control layer for AI coding agents — approve every change before it lands, audit what happened, meter what it cost. Local-first, Apache-2.0. · GitHub
Description: Control layer for AI coding agents — approve every change before it lands, audit what happened, meter what it cost. Local-first, Apache-2.0. - purser-sh/purser

Article text:
GitHub - purser-sh/purser: Control layer for AI coding agents — approve every change before it lands, audit what happened, meter what it cost. Local-first, Apache-2.0. · GitHub
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
purser-sh
/
purser
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
37 Commits 37 Commits Folders and files
.github .github Formula Formula apps apps docs docs extensions/ vscode extensions/ vscode packages packages scripts scripts .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE PRICING.md PRICING.md README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock bunfig.toml bunfig.toml install.sh install.sh package.json package.json tsconfig.base.json tsconfig.base.json turbo.json turbo.json View all files Repository files navigation
The purser for your coding agents.
Nothing an agent produces reaches your disk without your approval, and that is enforced by the type system: the only function that can write to your workspace requires a token only the approval step can mint.
Every change is recorded in a hash-chained audit log you can verify from the CLI. Every run is metered — tokens for every provider, dollars only where a plan price is knowable from the outside.
Agents run on your machine. Purser does not replace them, does not resell tokens, and never holds your provider logins — you keep those.
Free and open source · Apache-2.0 · Runs entirely on your machine
You run more than one coding agent now. Claude Code for one thing,
Cursor for another, a local model for the cheap work.
How much did they cost this week?
Claude shows you Claude. Cursor shows you Cursor. Nobody shows the total.
Which agent wrote this file, and did anyone check it?
The commit says your name. The prompt, the model and the review are gone.
What stops a runaway agent at 2am?
Nothing. You find out when the bill arrives.
Purser sits in front of the agents you already run and answers those
three questions. It does not replace your agents, does not resell
tokens, and never holds your provider logins — you keep those.
For anyone running more than one coding agent.
Primary path today is clone and run (the packaged installer lands with v0.1.0 ):
git clone https://github.com/purser-sh/purser
cd purser
bun install
bun run dev
Then open http://127.0.0.1:7410 , open a folder, pick a provider, and send a prompt. Use Echo first to confirm the console works before wiring a real provider.
Verify the audit chain (works immediately after bun install , no compile step):
bun run purser -- audit verify
Requires Bun ≥ 1.3.14 .
First start writes ~/.purser/config.json (mode 0600 ). The token is never printed . API keys go in Settings → ~/.purser/secrets.json , never SQLite.
bun test
bun run typecheck
Prerequisites, per provider
Provider
Needs
Echo
nothing — use it to check the console works
Ollama
ollama serve running + a coder-tuned model pulled (see Ollama models below)
Claude Code
npm i -g @anthropic-ai/claude-code , then claude → /login ; requires a Claude Pro or Max plan. Also bun add @anthropic-ai/claude-agent-sdk in this repo if the SDK package is missing
Codex / Cursor CLI / Gemini CLI
their CLI installed and logged in ( codex , cursor-agent , gemini ). Prefer cursor-agent over the short agent symlink — agent collides with other tools
Grok / Perplexity / OpenAI-compatible
an API key, added in Settings
Unready providers show as blocked in the top-bar selector with the exact command to fix them. Purser will not start a run against a provider it already knows will fail.
Ollama models: Purser sends eight workspace tools ( read_file , read_document , write_file , apply_patch , …) on every run. Many generic instruct models — including qwen2.5:7b-instruct , which declares tool support — still only call read/search tools and never write_file or apply_patch . Your run finishes with no proposed edit and it looks like Purser is broken; the tools were sent, the model just didn't use them.
For coding tasks, pull a coder-tuned model and select it in the top-bar model picker:
ollama pull qwen2.5-coder:7b # minimum for file edits; use 14b or 32b if you have VRAM
Do not rely on chat/instruct variants for edits unless you have verified they call write tools on your hardware.
Failure
Fix
Port 7410 / 7420 / 7430 already in use
Purser prints the port and how to free it (no Node stack). Free with lsof -ti:7410,7420,7430 | xargs -r kill , or move: PURSER_WEB_PORT=7411 PURSER_PORT=7421 PURSER_RELAY_PORT=7431 bun run dev .
Claude Agent SDK missing
From the repo root: bun add @anthropic-ai/claude-agent-sdk (workspace: packages/adapters ).
Claude Code: "Not logged in · Please run /login"
That /login is Claude's terminal command, not a Purser route. Run claude in a terminal, use /login , then reload Purser.
Ollama: llama-server binary not found
Broken Ollama install. Reinstall from https://ollama.com ( curl -fsSL https://ollama.com/install.sh | sh ), then ollama serve .
Ollama: connection refused
Start it: ollama serve .
No models / empty model list on Ollama
ollama pull qwen2.5-coder:7b (or larger), then pick it in the top bar.
Ollama run reads files but never proposes an edit
Purser sent write tools; the model didn't call them. Switch from an instruct/chat model to a coder model ( qwen2.5-coder:7b minimum). Instruct models often stop after read_file / ripgrep_search even when tool support is advertised.
Agent asked for README , got a bare miss
read_file (and list_dir ) now resolve an unambiguous extensionless hit ( README → README.md ) and otherwise return near matches: Did you mean: README.md, …?
ripgrep_search fails with Cause: …
The tool row shows the cause inline (not just ✗). The runner prepends standard system paths ( /usr/bin , /bin , …), Cursor/VS Code bundled rg , and your login-shell PATH at startup — so apt-installed ripgrep is found even when the IDE inherits a minimal bun/pyenv-only PATH. Restart Purser after installing ripgrep .
Git worktrees and your working folder
When you open a git repository , each new session gets an isolated worktree under ~/.purser/worktrees/<session-id>/ , checked out at HEAD (last commit). Parallel agents do not edit the same checkout.
Purser warns at session creation when your open folder has uncommitted changes. See also Setup → Workspace in the right panel.
Packaged install (not ready yet)
A one-line installer ships with v0.1.0. Until then, use the Quickstart above. Maintainer details: docs/RELEASING.md .
bun run compile builds the web UI (Vite), embeds it in the runner, and compiles a standalone binary — no separate manual UI build . Output lands in dist/bin/ (e.g. dist/bin/purser-linux-x64 on Linux, plus a dist/bin/purser symlink):
bun run compile
./dist/bin/purser audit verify
Development history — what landed in each phase
Done (Phases 0–7)
Phase
What landed
0
Loopback Host/Origin guards, no CORS on config, pairing frames sealed after pair, secrets out of SQLite
1
Append-only token ledger; official catalog only; unpriced models stay NULL cost (never invent dollars)
2
Budget governor (token + USD caps) with pre-run gate (blocks over-cap runs before the agent starts); spend UI
3
Hash-chained ~/.purser/audit.jsonl (mode 0600 ); bun run purser -- audit verify (or ./dist/bin/purser audit verify after compile) with per-break chain diagnostics; path redaction
4
Prompt coach: live token count under the composer ( exact / ≈ for the typed prompt), with a one-click shorter rewrite when one exists. Counts this prompt , not the agent loop — the run meter in the top bar is the spend headline
5
bun run compile / compile:all ; UI embedded in the binary; CI release workflow; token never printed
6
Public README vs architecture docs; competitor matrix; platform-risk notes
7
Provider readiness + vendor-error translation (protocol v4 ); provider/model coherence (no model id crosses a provider switch; ledger rejects impossible pairs)
Web console (UI phases 1–5)
Spend is the product surface; the accent colour marks one primary action per screen (usually Send).
Browser vs runner tokenizers: The prompt coach runs in the browser (Vite aliases @purser-sh/pricing → browser.ts ). Anthropic ids stay approximate in-browser; the runner keeps exact Claude counts for the ledger. See docs/METERING.md .
Also shipped: workspaces, sessions (title from first prompt), diffs, permission modes ( ask / auto_edit / bypass with TTL), adapters (Echo, Claude Code, Codex, Cursor CLI, Gemini CLI, Ollama, Grok, OpenAI-compatible, Perplexity), session git worktrees at HEAD, read_file / read_document / list_dir near-match hints, runner PATH augmentation for ripgrep_search , drop-folder → .inbox/ , voice + /phone .
Tool
What it does
read_file
Text files inside the workspace (512 KB preview cap)
read_document
PDF, Word (DOCX), and Excel (XLSX) via built-in converters; PowerPoint, images, and other formats when MarkItDown is installed ( pip install 'markitdown[all]' )
write_file / apply_patch
Staged edits with diff cards
list_dir / ripgrep_search
Navigate and search the tree
run_bash
Opt-in shell (workspace setting)
web_search
Research providers only
read_document converts locally (no network). Large conversions ask before entering context; see Settings → Workspace for the token threshold and file-size cap.
No tagged public release or packaged installer yet (use the Quickstart)
Live Postgres / hosted cells (schema and types only; postgres URL throws)
VS Code / Cursor marketplace extensions (protocol notes only)
A Google / Gemini local tokenizer (coach counts for gemini_cli stay approximate in-browser)
Purser's own price (see PRICING.md )
Doc
What
docs/ARCHITECTURE.md
What Purser is, system diagram, safety model
docs/REVIEW.md
Protocol, repo map, review checklist
docs/SECURITY.md
Companion threat model
docs/METERING.md
What we can observe and price
docs/COMPETITORS.md
Honest matrix
docs/PLATFORM-RISK.md
Vendor terms we opened
docs/RELEASING.md
Binary, embed, codesign secrets
Repo
apps/web React console (Vite + Tailwind v4 + tokens)
apps/runner Bun websocket server, ledger, budgets, adapters
apps/relay Optional phone pairing relay
packages/protocol shared frames (protocol v4, TokenCount, readiness, provider/model table)
packages/pricing catalog, tokenizer, browser.ts for Vite
packages/prompt-coach pre-send token estimate
packages/db adapters voice integrations
Licensed under Apache-2.0 (see LICENSE ). Purser's own price is not decided yet — see PRICING.md .
Control layer for AI coding agents — approve every change before it lands, audit what happened, meter what i

[truncated]

## Original Extract

Control layer for AI coding agents — approve every change before it lands, audit what happened, meter what it cost. Local-first, Apache-2.0. - purser-sh/purser

GitHub - purser-sh/purser: Control layer for AI coding agents — approve every change before it lands, audit what happened, meter what it cost. Local-first, Apache-2.0. · GitHub
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
purser-sh
/
purser
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
37 Commits 37 Commits Folders and files
.github .github Formula Formula apps apps docs docs extensions/ vscode extensions/ vscode packages packages scripts scripts .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE PRICING.md PRICING.md README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock bunfig.toml bunfig.toml install.sh install.sh package.json package.json tsconfig.base.json tsconfig.base.json turbo.json turbo.json View all files Repository files navigation
The purser for your coding agents.
Nothing an agent produces reaches your disk without your approval, and that is enforced by the type system: the only function that can write to your workspace requires a token only the approval step can mint.
Every change is recorded in a hash-chained audit log you can verify from the CLI. Every run is metered — tokens for every provider, dollars only where a plan price is knowable from the outside.
Agents run on your machine. Purser does not replace them, does not resell tokens, and never holds your provider logins — you keep those.
Free and open source · Apache-2.0 · Runs entirely on your machine
You run more than one coding agent now. Claude Code for one thing,
Cursor for another, a local model for the cheap work.
How much did they cost this week?
Claude shows you Claude. Cursor shows you Cursor. Nobody shows the total.
Which agent wrote this file, and did anyone check it?
The commit says your name. The prompt, the model and the review are gone.
What stops a runaway agent at 2am?
Nothing. You find out when the bill arrives.
Purser sits in front of the agents you already run and answers those
three questions. It does not replace your agents, does not resell
tokens, and never holds your provider logins — you keep those.
For anyone running more than one coding agent.
Primary path today is clone and run (the packaged installer lands with v0.1.0 ):
git clone https://github.com/purser-sh/purser
cd purser
bun install
bun run dev
Then open http://127.0.0.1:7410 , open a folder, pick a provider, and send a prompt. Use Echo first to confirm the console works before wiring a real provider.
Verify the audit chain (works immediately after bun install , no compile step):
bun run purser -- audit verify
Requires Bun ≥ 1.3.14 .
First start writes ~/.purser/config.json (mode 0600 ). The token is never printed . API keys go in Settings → ~/.purser/secrets.json , never SQLite.
bun test
bun run typecheck
Prerequisites, per provider
Provider
Needs
Echo
nothing — use it to check the console works
Ollama
ollama serve running + a coder-tuned model pulled (see Ollama models below)
Claude Code
npm i -g @anthropic-ai/claude-code , then claude → /login ; requires a Claude Pro or Max plan. Also bun add @anthropic-ai/claude-agent-sdk in this repo if the SDK package is missing
Codex / Cursor CLI / Gemini CLI
their CLI installed and logged in ( codex , cursor-agent , gemini ). Prefer cursor-agent over the short agent symlink — agent collides with other tools
Grok / Perplexity / OpenAI-compatible
an API key, added in Settings
Unready providers show as blocked in the top-bar selector with the exact command to fix them. Purser will not start a run against a provider it already knows will fail.
Ollama models: Purser sends eight workspace tools ( read_file , read_document , write_file , apply_patch , …) on every run. Many generic instruct models — including qwen2.5:7b-instruct , which declares tool support — still only call read/search tools and never write_file or apply_patch . Your run finishes with no proposed edit and it looks like Purser is broken; the tools were sent, the model just didn't use them.
For coding tasks, pull a coder-tuned model and select it in the top-bar model picker:
ollama pull qwen2.5-coder:7b # minimum for file edits; use 14b or 32b if you have VRAM
Do not rely on chat/instruct variants for edits unless you have verified they call write tools on your hardware.
Failure
Fix
Port 7410 / 7420 / 7430 already in use
Purser prints the port and how to free it (no Node stack). Free with lsof -ti:7410,7420,7430 | xargs -r kill , or move: PURSER_WEB_PORT=7411 PURSER_PORT=7421 PURSER_RELAY_PORT=7431 bun run dev .
Claude Agent SDK missing
From the repo root: bun add @anthropic-ai/claude-agent-sdk (workspace: packages/adapters ).
Claude Code: "Not logged in · Please run /login"
That /login is Claude's terminal command, not a Purser route. Run claude in a terminal, use /login , then reload Purser.
Ollama: llama-server binary not found
Broken Ollama install. Reinstall from https://ollama.com ( curl -fsSL https://ollama.com/install.sh | sh ), then ollama serve .
Ollama: connection refused
Start it: ollama serve .
No models / empty model list on Ollama
ollama pull qwen2.5-coder:7b (or larger), then pick it in the top bar.
Ollama run reads files but never proposes an edit
Purser sent write tools; the model didn't call them. Switch from an instruct/chat model to a coder model ( qwen2.5-coder:7b minimum). Instruct models often stop after read_file / ripgrep_search even when tool support is advertised.
Agent asked for README , got a bare miss
read_file (and list_dir ) now resolve an unambiguous extensionless hit ( README → README.md ) and otherwise return near matches: Did you mean: README.md, …?
ripgrep_search fails with Cause: …
The tool row shows the cause inline (not just ✗). The runner prepends standard system paths ( /usr/bin , /bin , …), Cursor/VS Code bundled rg , and your login-shell PATH at startup — so apt-installed ripgrep is found even when the IDE inherits a minimal bun/pyenv-only PATH. Restart Purser after installing ripgrep .
Git worktrees and your working folder
When you open a git repository , each new session gets an isolated worktree under ~/.purser/worktrees/<session-id>/ , checked out at HEAD (last commit). Parallel agents do not edit the same checkout.
Purser warns at session creation when your open folder has uncommitted changes. See also Setup → Workspace in the right panel.
Packaged install (not ready yet)
A one-line installer ships with v0.1.0. Until then, use the Quickstart above. Maintainer details: docs/RELEASING.md .
bun run compile builds the web UI (Vite), embeds it in the runner, and compiles a standalone binary — no separate manual UI build . Output lands in dist/bin/ (e.g. dist/bin/purser-linux-x64 on Linux, plus a dist/bin/purser symlink):
bun run compile
./dist/bin/purser audit verify
Development history — what landed in each phase
Done (Phases 0–7)
Phase
What landed
0
Loopback Host/Origin guards, no CORS on config, pairing frames sealed after pair, secrets out of SQLite
1
Append-only token ledger; official catalog only; unpriced models stay NULL cost (never invent dollars)
2
Budget governor (token + USD caps) with pre-run gate (blocks over-cap runs before the agent starts); spend UI
3
Hash-chained ~/.purser/audit.jsonl (mode 0600 ); bun run purser -- audit verify (or ./dist/bin/purser audit verify after compile) with per-break chain diagnostics; path redaction
4
Prompt coach: live token count under the composer ( exact / ≈ for the typed prompt), with a one-click shorter rewrite when one exists. Counts this prompt , not the agent loop — the run meter in the top bar is the spend headline
5
bun run compile / compile:all ; UI embedded in the binary; CI release workflow; token never printed
6
Public README vs architecture docs; competitor matrix; platform-risk notes
7
Provider readiness + vendor-error translation (protocol v4 ); provider/model coherence (no model id crosses a provider switch; ledger rejects impossible pairs)
Web console (UI phases 1–5)
Spend is the product surface; the accent colour marks one primary action per screen (usually Send).
Browser vs runner tokenizers: The prompt coach runs in the browser (Vite aliases @purser-sh/pricing → browser.ts ). Anthropic ids stay approximate in-browser; the runner keeps exact Claude counts for the ledger. See docs/METERING.md .
Also shipped: workspaces, sessions (title from first prompt), diffs, permission modes ( ask / auto_edit / bypass with TTL), adapters (Echo, Claude Code, Codex, Cursor CLI, Gemini CLI, Ollama, Grok, OpenAI-compatible, Perplexity), session git worktrees at HEAD, read_file / read_document / list_dir near-match hints, runner PATH augmentation for ripgrep_search , drop-folder → .inbox/ , voice + /phone .
Tool
What it does
read_file
Text files inside the workspace (512 KB preview cap)
read_document
PDF, Word (DOCX), and Excel (XLSX) via built-in converters; PowerPoint, images, and other formats when MarkItDown is installed ( pip install 'markitdown[all]' )
write_file / apply_patch
Staged edits with diff cards
list_dir / ripgrep_search
Navigate and search the tree
run_bash
Opt-in shell (workspace setting)
web_search
Research providers only
read_document converts locally (no network). Large conversions ask before entering context; see Settings → Workspace for the token threshold and file-size cap.
No tagged public release or packaged installer yet (use the Quickstart)
Live Postgres / hosted cells (schema and types only; postgres URL throws)
VS Code / Cursor marketplace extensions (protocol notes only)
A Google / Gemini local tokenizer (coach counts for gemini_cli stay approximate in-browser)
Purser's own price (see PRICING.md )
Doc
What
docs/ARCHITECTURE.md
What Purser is, system diagram, safety model
docs/REVIEW.md
Protocol, repo map, review checklist
docs/SECURITY.md
Companion threat model
docs/METERING.md
What we can observe and price
docs/COMPETITORS.md
Honest matrix
docs/PLATFORM-RISK.md
Vendor terms we opened
docs/RELEASING.md
Binary, embed, codesign secrets
Repo
apps/web React console (Vite + Tailwind v4 + tokens)
apps/runner Bun websocket server, ledger, budgets, adapters
apps/relay Optional phone pairing relay
packages/protocol shared frames (protocol v4, TokenCount, readiness, provider/model table)
packages/pricing catalog, tokenizer, browser.ts for Vite
packages/prompt-coach pre-send token estimate
packages/db adapters voice integrations
Licensed under Apache-2.0 (see LICENSE ). Purser's own price is not decided yet — see PRICING.md .
Control layer for AI coding agents — approve every change before it lands, audit what happened, meter what i

[truncated]
