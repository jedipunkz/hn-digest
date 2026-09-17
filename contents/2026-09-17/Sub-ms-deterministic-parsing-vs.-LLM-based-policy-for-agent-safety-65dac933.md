---
source: "https://github.com/midhunweb/miseguard"
hn_url: "https://news.ycombinator.com/item?id=49742848"
title: "Sub-ms deterministic parsing vs. LLM-based policy for agent safety?"
article_title: "GitHub - midhunweb/miseguard: Deterministic runtime circuit breaker and MCP proxy for autonomous coding agents. · GitHub"
image: "https://opengraph.githubassets.com/72fd22eb602832397e6bb11f32a049996dff659d98ce3e0c2b0911f1a0764a61/midhunweb/miseguard"
author: "misetro"
captured_at: "2026-09-17T16:22:59Z"
capture_tool: "hn-digest"
hn_id: 49742848
score: 1
comments: 0
posted_at: "2026-09-17T16:08:28Z"
tags:
  - hacker-news
---

# Sub-ms deterministic parsing vs. LLM-based policy for agent safety?

- HN: [49742848](https://news.ycombinator.com/item?id=49742848)
- Source: [github.com](https://github.com/midhunweb/miseguard)
- Score: 1
- Comments: 0
- Posted: 2026-09-17T16:08:28Z

## Translation

Title: Sub-ms deterministic parsing vs. LLM-based policy for agent safety?
Article title: GitHub - midhunweb/miseguard: Deterministic runtime circuit breaker and MCP proxy for autonomous coding agents. · GitHub
Description: Deterministic runtime circuit breaker and MCP proxy for autonomous coding agents. - midhunweb/miseguard

Article text:
GitHub - midhunweb/miseguard: Deterministic runtime circuit breaker and MCP proxy for autonomous coding agents. · GitHub
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
midhunweb
/
miseguard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
10 Commits 10 Commits Folders and files
.github .github bin bin docs docs scripts scripts src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md miseguard.json miseguard.json package-lock.json package-lock.json package.json package.json schema.json schema.json tsconfig.json tsconfig.json View all files Repository files navigation
MiSeGuard is an independent personal project created by Midhun Sekhar.
It is not affiliated with, endorsed by, or representative of any current, past, or future employer.
All development was conducted entirely on personal equipment, during personal time, and without the use of proprietary resources or confidential information.
A deterministic safety layer for autonomous coding agents.
Runtime circuit breaker and stdio proxy that intercepts MCP tool calls before they reach your OS.
AI coding agents can modify your machine. MiSeGuard puts a deterministic security boundary between the agent and your tools.
What is MCP?
The Model Context Protocol (MCP) is an open standard that lets autonomous AI agents (such as Cursor, Claude Code, Antigravity, OpenCode, Windsurf ) invoke external tools (bash, filesystem, git, terminal) over stdio JSON-RPC 2.0. MiSeGuard sits as a transparent, sub-millisecond proxy between your agent and those tool runtimes to inspect, score, and block destructive operations before they reach the real operating system.
What Does "Deterministic" Mean?
The same command or file mutation with the same configuration always produces the exact identical risk score. No LLM in the loop, no non-deterministic inference, no prompt drift.
# Global install (recommended for CLI use)
npm install -g miseguard
# Or run directly via npx
npx miseguard --help
# Or add as a project dev dependency
npm install --save-dev miseguard
Requirements: Node.js 18.0 or later.
1. Initialize workspace security policy ( miseguard.json )
miseguard init
2. Shield your MCP agent tools
Choose the method that matches your workflow:
Starting fresh or configuring an agent GUI? Use snippet to generate copy-pasteable JSON:
miseguard snippet --tool filesystem --path .
Already have an existing MCP configuration file? Use wrap-config <file-path> to automatically rewrite and back up your config in place:
# Pass the explicit path to your agent's config file:
miseguard wrap-config .antigravity/mcp.json
miseguard wrap-config " %APPDATA%\Claude\claude_desktop_config.json "
# Or omit path to auto-detect mcp.json / .cursor/mcp.json in current directory:
miseguard wrap-config
3. Test the deterministic circuit breaker
miseguard check " rm -rf / " # 🛑 Exit Code 1: Blocked
miseguard check " git status " # 🟢 Exit Code 0: Safe
🚦 Deterministic Risk Tiers & Exit Code Contract
MiSeGuard computes a multi-factor Blast-Radius Risk Score (0–100) for every tool invocation.
📌 CLI Exit Code Stability Contract:
0 : Safe operation (Green) or permitted in permissive mode.
1 : Dangerous operation blocked by circuit breaker (Red).
2 : Caution operation triggering dry-run in strict mode (Yellow).
3 : Internal parsing or configuration error.
Exit codes are stable and guaranteed across versions for CI/CD and pre-commit hook integration.
⚙️ Configuration ( miseguard.json )
Generate a starter configuration file in your project:
miseguard init
Configuration Options
{
"$schema" : " https://raw.githubusercontent.com/midhunweb/miseguard/main/schema.json " ,
"mode" : " strict " ,
"thresholds" : {
"block" : 70 ,
"dryRun" : 30
},
"allowlist" : [
" echo * " ,
" git log* " ,
" git status* " ,
" git diff* " ,
" npm run test* " ,
" npm run lint* " ,
" npm run clean:* "
],
"protectedPaths" : [
" .env* " ,
" *.pem " ,
" *.key " ,
" id_rsa* " ,
" id_ed25519* " ,
" ~/.ssh/* " ,
" ~/.aws/* " ,
" ~/.kube/* " ,
" .git/* " ,
" secrets/** "
]
}
mode :
"strict" (default): Yellow tier actions trigger ephemeral sandbox dry-run simulation; Red tier actions are blocked.
"permissive" : Yellow tier actions log warnings and allow execution; Red tier actions are still blocked.
allowlist : Commands or file targets matching these patterns are unconditionally granted Green (Score 0) status.
protectedPaths : Glob patterns of sensitive files that immediately elevate risk to Red (Score >= 70) upon access or mutation attempt.
thresholds : Customize risk boundaries for block and dryRun .
1. miseguard wrap-config [file-path]
Safely transforms tools in an existing MCP configuration file so commands run shielded behind miseguard proxy -- . Supports explicit file paths or workspace auto-discovery:
# Explicit path (recommended across agents):
miseguard wrap-config .antigravity/mcp.json
miseguard wrap-config " %APPDATA%\Claude\claude_desktop_config.json "
miseguard wrap-config ~ /.config/Claude/claude_desktop_config.json
# Workspace auto-detection (scans for mcp.json, .cursor/mcp.json, .antigravity/mcp.json):
miseguard wrap-config
Creates <file-path>.bak before modification and guarantees idempotency.
2. miseguard snippet [options]
Generates copy-pasteable JSON configuration blocks for agent GUI settings (Cursor, Claude, Antigravity, Windsurf):
# Filesystem preset (default)
miseguard snippet --tool filesystem --path ./
# Git preset
miseguard snippet --tool git --path ./
# Bash/Terminal preset
miseguard snippet --tool bash
# Custom tool preset
miseguard snippet --tool custom --name my-server --cmd python --args -m my_module
3. miseguard check "<command>"
Evaluates the blast-radius risk score of any shell command:
miseguard check " rm -rf / "
4. miseguard dry-run "<command>"
Simulates a command inside an isolated ephemeral shadow sandbox and outputs a SHA-256 filesystem delta table:
miseguard dry-run " npm run build "
5. miseguard proxy -- <command...>
Runs MiSeGuard as an active stdio proxy in front of an MCP server process:
miseguard proxy -- npx -y @modelcontextprotocol/server-filesystem ./
6. miseguard rules
Displays the complete deterministic security rule matrix.
MiSeGuard adds negligible overhead. The numbers below measure policy evaluation and interception logic — the scoring, sandbox dispatch, and diff-checking path. They exclude process spawn, JSON serialization, and OS scheduling, which are common to all stdio proxies and not attributable to MiSeGuard.
Reproducibility: Full methodology, hardware specs, warm-up procedure, and percentile distributions are in docs/BENCHMARKS.md . Run npm run benchmark to reproduce on your own machine.
MiSeGuard operates at the Model Context Protocol (MCP) stdio layer . It intercepts every tools/call JSON-RPC message that flows between an agent and an MCP tool server (bash, filesystem, git, etc.).
Being explicit about architectural boundaries:
Does not protect against prompt injection — That is an LLM inference layer concern. MiSeGuard assumes the agent's intent may be compromised or hallucinatory, and deterministically enforces policy on the executed action , not the reasoning.
Does not intercept native IDE built-in tools — Cursor's internal run_command , Antigravity's internal edit_file , and Windsurf's native terminal bypass MCP entirely. MiSeGuard only inspects MCP stdio traffic. Direct IDE extension hooks are planned for v0.3.0.
Does not sandbox long-running persistent VM state — Ephemeral shadow sandboxes for dry-runs are discarded immediately after filesystem diff analysis.
Does not support HTTP/gRPC transports yet — Standard input/output ( stdio ) JSON-RPC 2.0 only for v0.1.0 (HTTP/SSE transport on roadmap for v0.2.0).
Does not use ML or probabilistic heuristics for risk scoring — By design. Determinism and reproducibility are core security features.
Does not defend against kernel-level escapes or raw syscall bypasses — Operates at the tool runtime protocol layer.
v0.2.0 : HTTP / SSE / gRPC MCP transport support
v0.3.0 : IDE Extension / LSP wrapper for Cursor, Antigravity, and Windsurf native tools
Architecture & Design Specification →
Agent Frontend Integration Guide (Cursor, Claude Desktop, Antigravity) →
# Run all 60 unit and integration tests
npm test
# Run latency benchmark suite
npm run benchmark
📜 License
MIT License. Copyright (c) 2026 Midhun Sekhar & MiSeGuard Contributors.
Deterministic runtime circuit breaker and MCP proxy for autonomous coding agents.
www.npmjs.com/package/miseguard Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Deterministic runtime circuit breaker and MCP proxy for autonomous coding agents. - midhunweb/miseguard

GitHub - midhunweb/miseguard: Deterministic runtime circuit breaker and MCP proxy for autonomous coding agents. · GitHub
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
midhunweb
/
miseguard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
10 Commits 10 Commits Folders and files
.github .github bin bin docs docs scripts scripts src src tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md miseguard.json miseguard.json package-lock.json package-lock.json package.json package.json schema.json schema.json tsconfig.json tsconfig.json View all files Repository files navigation
MiSeGuard is an independent personal project created by Midhun Sekhar.
It is not affiliated with, endorsed by, or representative of any current, past, or future employer.
All development was conducted entirely on personal equipment, during personal time, and without the use of proprietary resources or confidential information.
A deterministic safety layer for autonomous coding agents.
Runtime circuit breaker and stdio proxy that intercepts MCP tool calls before they reach your OS.
AI coding agents can modify your machine. MiSeGuard puts a deterministic security boundary between the agent and your tools.
What is MCP?
The Model Context Protocol (MCP) is an open standard that lets autonomous AI agents (such as Cursor, Claude Code, Antigravity, OpenCode, Windsurf ) invoke external tools (bash, filesystem, git, terminal) over stdio JSON-RPC 2.0. MiSeGuard sits as a transparent, sub-millisecond proxy between your agent and those tool runtimes to inspect, score, and block destructive operations before they reach the real operating system.
What Does "Deterministic" Mean?
The same command or file mutation with the same configuration always produces the exact identical risk score. No LLM in the loop, no non-deterministic inference, no prompt drift.
# Global install (recommended for CLI use)
npm install -g miseguard
# Or run directly via npx
npx miseguard --help
# Or add as a project dev dependency
npm install --save-dev miseguard
Requirements: Node.js 18.0 or later.
1. Initialize workspace security policy ( miseguard.json )
miseguard init
2. Shield your MCP agent tools
Choose the method that matches your workflow:
Starting fresh or configuring an agent GUI? Use snippet to generate copy-pasteable JSON:
miseguard snippet --tool filesystem --path .
Already have an existing MCP configuration file? Use wrap-config <file-path> to automatically rewrite and back up your config in place:
# Pass the explicit path to your agent's config file:
miseguard wrap-config .antigravity/mcp.json
miseguard wrap-config " %APPDATA%\Claude\claude_desktop_config.json "
# Or omit path to auto-detect mcp.json / .cursor/mcp.json in current directory:
miseguard wrap-config
3. Test the deterministic circuit breaker
miseguard check " rm -rf / " # 🛑 Exit Code 1: Blocked
miseguard check " git status " # 🟢 Exit Code 0: Safe
🚦 Deterministic Risk Tiers & Exit Code Contract
MiSeGuard computes a multi-factor Blast-Radius Risk Score (0–100) for every tool invocation.
📌 CLI Exit Code Stability Contract:
0 : Safe operation (Green) or permitted in permissive mode.
1 : Dangerous operation blocked by circuit breaker (Red).
2 : Caution operation triggering dry-run in strict mode (Yellow).
3 : Internal parsing or configuration error.
Exit codes are stable and guaranteed across versions for CI/CD and pre-commit hook integration.
⚙️ Configuration ( miseguard.json )
Generate a starter configuration file in your project:
miseguard init
Configuration Options
{
"$schema" : " https://raw.githubusercontent.com/midhunweb/miseguard/main/schema.json " ,
"mode" : " strict " ,
"thresholds" : {
"block" : 70 ,
"dryRun" : 30
},
"allowlist" : [
" echo * " ,
" git log* " ,
" git status* " ,
" git diff* " ,
" npm run test* " ,
" npm run lint* " ,
" npm run clean:* "
],
"protectedPaths" : [
" .env* " ,
" *.pem " ,
" *.key " ,
" id_rsa* " ,
" id_ed25519* " ,
" ~/.ssh/* " ,
" ~/.aws/* " ,
" ~/.kube/* " ,
" .git/* " ,
" secrets/** "
]
}
mode :
"strict" (default): Yellow tier actions trigger ephemeral sandbox dry-run simulation; Red tier actions are blocked.
"permissive" : Yellow tier actions log warnings and allow execution; Red tier actions are still blocked.
allowlist : Commands or file targets matching these patterns are unconditionally granted Green (Score 0) status.
protectedPaths : Glob patterns of sensitive files that immediately elevate risk to Red (Score >= 70) upon access or mutation attempt.
thresholds : Customize risk boundaries for block and dryRun .
1. miseguard wrap-config [file-path]
Safely transforms tools in an existing MCP configuration file so commands run shielded behind miseguard proxy -- . Supports explicit file paths or workspace auto-discovery:
# Explicit path (recommended across agents):
miseguard wrap-config .antigravity/mcp.json
miseguard wrap-config " %APPDATA%\Claude\claude_desktop_config.json "
miseguard wrap-config ~ /.config/Claude/claude_desktop_config.json
# Workspace auto-detection (scans for mcp.json, .cursor/mcp.json, .antigravity/mcp.json):
miseguard wrap-config
Creates <file-path>.bak before modification and guarantees idempotency.
2. miseguard snippet [options]
Generates copy-pasteable JSON configuration blocks for agent GUI settings (Cursor, Claude, Antigravity, Windsurf):
# Filesystem preset (default)
miseguard snippet --tool filesystem --path ./
# Git preset
miseguard snippet --tool git --path ./
# Bash/Terminal preset
miseguard snippet --tool bash
# Custom tool preset
miseguard snippet --tool custom --name my-server --cmd python --args -m my_module
3. miseguard check "<command>"
Evaluates the blast-radius risk score of any shell command:
miseguard check " rm -rf / "
4. miseguard dry-run "<command>"
Simulates a command inside an isolated ephemeral shadow sandbox and outputs a SHA-256 filesystem delta table:
miseguard dry-run " npm run build "
5. miseguard proxy -- <command...>
Runs MiSeGuard as an active stdio proxy in front of an MCP server process:
miseguard proxy -- npx -y @modelcontextprotocol/server-filesystem ./
6. miseguard rules
Displays the complete deterministic security rule matrix.
MiSeGuard adds negligible overhead. The numbers below measure policy evaluation and interception logic — the scoring, sandbox dispatch, and diff-checking path. They exclude process spawn, JSON serialization, and OS scheduling, which are common to all stdio proxies and not attributable to MiSeGuard.
Reproducibility: Full methodology, hardware specs, warm-up procedure, and percentile distributions are in docs/BENCHMARKS.md . Run npm run benchmark to reproduce on your own machine.
MiSeGuard operates at the Model Context Protocol (MCP) stdio layer . It intercepts every tools/call JSON-RPC message that flows between an agent and an MCP tool server (bash, filesystem, git, etc.).
Being explicit about architectural boundaries:
Does not protect against prompt injection — That is an LLM inference layer concern. MiSeGuard assumes the agent's intent may be compromised or hallucinatory, and deterministically enforces policy on the executed action , not the reasoning.
Does not intercept native IDE built-in tools — Cursor's internal run_command , Antigravity's internal edit_file , and Windsurf's native terminal bypass MCP entirely. MiSeGuard only inspects MCP stdio traffic. Direct IDE extension hooks are planned for v0.3.0.
Does not sandbox long-running persistent VM state — Ephemeral shadow sandboxes for dry-runs are discarded immediately after filesystem diff analysis.
Does not support HTTP/gRPC transports yet — Standard input/output ( stdio ) JSON-RPC 2.0 only for v0.1.0 (HTTP/SSE transport on roadmap for v0.2.0).
Does not use ML or probabilistic heuristics for risk scoring — By design. Determinism and reproducibility are core security features.
Does not defend against kernel-level escapes or raw syscall bypasses — Operates at the tool runtime protocol layer.
v0.2.0 : HTTP / SSE / gRPC MCP transport support
v0.3.0 : IDE Extension / LSP wrapper for Cursor, Antigravity, and Windsurf native tools
Architecture & Design Specification →
Agent Frontend Integration Guide (Cursor, Claude Desktop, Antigravity) →
# Run all 60 unit and integration tests
npm test
# Run latency benchmark suite
npm run benchmark
📜 License
MIT License. Copyright (c) 2026 Midhun Sekhar & MiSeGuard Contributors.
Deterministic runtime circuit breaker and MCP proxy for autonomous coding agents.
www.npmjs.com/package/miseguard Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
