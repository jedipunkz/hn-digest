---
source: "https://github.com/vayungodara/cliproxy-rs"
hn_url: "https://news.ycombinator.com/item?id=49948441"
title: "Show HN: Cliproxy-rs, a Rust port of the CLIProxyAPI LLM subscription proxy"
article_title: "GitHub - vayungodara/cliproxy-rs: Rust rewrite of CLIProxyAPI: drop-in compatible proxy for Claude, Codex, Gemini and OpenAI-compatible providers · GitHub"
image: "https://opengraph.githubassets.com/64c09a6c1fa68bf88fa2a9df0f4b89289e670f915c1810e552394d717cd2bcd7/vayungodara/cliproxy-rs"
author: "vayungodara"
captured_at: "2026-10-03T22:43:52Z"
capture_tool: "hn-digest"
hn_id: 49948441
score: 1
comments: 0
posted_at: "2026-10-03T22:36:03Z"
tags:
  - hacker-news
---

# Show HN: Cliproxy-rs, a Rust port of the CLIProxyAPI LLM subscription proxy

- HN: [49948441](https://news.ycombinator.com/item?id=49948441)
- Source: [github.com](https://github.com/vayungodara/cliproxy-rs)
- Score: 1
- Comments: 0
- Posted: 2026-10-03T22:36:03Z

## Translation

Title: Show HN: Cliproxy-rs, a Rust port of the CLIProxyAPI LLM subscription proxy
Article title: GitHub - vayungodara/cliproxy-rs: Rust rewrite of CLIProxyAPI: drop-in compatible proxy for Claude, Codex, Gemini and OpenAI-compatible providers · GitHub
Description: Rust rewrite of CLIProxyAPI: drop-in compatible proxy for Claude, Codex, Gemini and OpenAI-compatible providers - vayungodara/cliproxy-rs

Article text:
GitHub - vayungodara/cliproxy-rs: Rust rewrite of CLIProxyAPI: drop-in compatible proxy for Claude, Codex, Gemini and OpenAI-compatible providers · GitHub
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
vayungodara
/
cliproxy-rs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
587 Commits 587 Commits Folders and files
.github .github bench bench crates crates docs docs harness harness ui ui .dockerignore .dockerignore .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md install.ps1 install.ps1 install.sh install.sh rustfmt.toml rustfmt.toml View all files Repository files navigation
cliproxy-rs is a small server you run on your own computer. It lets your coding tools, such as Claude Code, Codex CLI or Cursor, use the AI subscriptions and API keys you already have, like a Claude Max plan or a ChatGPT plan, through one local address. A dashboard built into it shows which accounts work, how much of each plan's limits is left, and how to connect each tool.
It is a Rust rewrite of CLIProxyAPI . It reads the same config.yaml and the same account files, so you can switch between the two in either direction.
You pay for Claude and ChatGPT plans and want Claude Code, Codex CLI and your editor to share them, with the 5-hour and weekly limits of every account on one screen.
You have several Claude or Codex accounts and want requests spread across them, or moved to the next account when one reaches its limit.
A tool speaks only one API: the proxy translates between the Anthropic, OpenAI and Gemini formats, so an OpenAI-only tool can use a Claude account and the other way round.
You run CLIProxyAPI and want a smaller, faster-starting binary with the dashboard built in.
Run one command. On macOS or Linux:
curl -fsSL https://raw.githubusercontent.com/vayungodara/cliproxy-rs/master/install.sh | sh
On Windows, in PowerShell:
irm https: // raw.githubusercontent.com / vayungodara / cliproxy - rs / master / install.ps1 | iex
It installs the latest release after checking its checksum, writes a config with new keys to ~/.cliproxy-rs , starts the proxy in the background and opens the dashboard in your browser. The keys stay in ~/.cliproxy-rs/keys.env and are never printed.
In the dashboard, sign in with the CLIPROXY_MANAGEMENT_KEY line from keys.env and choose Connect account.
Open Use with tools and copy the settings for Claude Code, Codex CLI or another tool.
Run the same command again to upgrade; your config and keys stay as they are. docs/INSTALL.md shows how to start the proxy when you log in.
docs/GETTING-STARTED.md walks through the same steps in more detail.
Using a subscription outside its official app can break the provider's terms, and providers have suspended accounts for it. Whether to do that is your call and your risk.
Paste this into Claude Code, Codex or another coding agent on the computer where you want the proxy:
Install and set up cliproxy-rs on this computer by following
https://github.com/vayungodara/cliproxy-rs/blob/master/docs/AI-SETUP.md exactly.
Do not sign in to any of my accounts or print my keys; tell me when it is my turn.
The agent runs the installer, which picks a free port without touching an existing CLIProxyAPI and keeps the keys in a file only you can read. It checks that the proxy answers, then hands the account sign-in to you.
The proxy serves the Anthropic Messages API at http://127.0.0.1:8317 , the OpenAI API at http://127.0.0.1:8317/v1 and the Gemini API at http://127.0.0.1:8317 , all with your client key. For Claude Code:
export ANTHROPIC_BASE_URL=http://127.0.0.1:8317
export ANTHROPIC_AUTH_TOKEN=your-client-key
claude
docs/CLIENTS.md has copy-paste setups for Claude Code, Codex CLI, Gemini CLI, Amp, OpenCode, Factory Droid, Cline, Roo Code, Kilo Code, Cursor, Zed, Continue, Aider and the OpenAI, Anthropic and Gemini SDKs.
Sign in to as many Claude and Codex accounts as you have, and choose how the proxy rotates between them: round-robin , fill-first , weighted-round-robin , or soonest-reset (experimental, a cliproxy-rs addition: spend the account whose weekly window resets soonest first). Session affinity keeps each conversation on one account so the provider's prompt cache stays warm, and an account that hits a limit rests while the others take over. docs/MULTI-ACCOUNT.md covers routing, cooldowns, the quota view, per-account proxies, Codex over WebSocket and reaching the proxy from your other machines over Tailscale.
cliproxy-rs serves the same routes and the same v8 Management API as CLIProxyAPI, so the community apps built on CLIProxyAPI should work with it, for example CPA-Manager-Plus , CLIProxyAPI Quota Inspector , ZeroLimit , Quotio , VibeProxy and CCS . We have not tested all of them. Apps that talk to a running server need only its address and management key; apps that start their own bundled CLIProxyAPI need an option to use an existing server instead. If an app does not work with cliproxy-rs, please open an issue .
install.sh (macOS and Linux) and install.ps1 (Windows), the commands in the quick start. They install the release binary, set it up and start it; --binary-only ( -BinaryOnly on Windows) installs only the binary.
Release binaries for macOS (Apple silicon and Intel), Linux (x86_64 and arm64) and Windows (x86_64), with a SHA256SUMS file, on the releases page .
Docker, from the repository's Dockerfile .
Building from source, for contributors and platforms without a release binary.
docs/INSTALL.md covers each one, starting the proxy at login, and removing it.
These are not in cliproxy-rs yet. They are planned, in no fixed order:
Plugins: the host callbacks plugins use to call back into the server, calling plugins on the request path (plugin providers, sign-in, models and usage), and the plugin store. Today plugins load, serve their own routes and report quotas, but requests do not pass through them.
Home (cluster) mode: reporting usage, logs and in-flight requests back to Home, Home's KV storage, and syncing plugins managed by Home.
Image generation and editing through Codex accounts, and importing Vertex service accounts from the dashboard (the command line works).
The upstream request and response sections of request log files.
The rest of CLIProxyAPI's own test cases. The parity audit checked 1,687 Go routes, settings, flags and test suites: 798 (47%) are fully covered, 697 partly and 192 not yet. docs/PARITY.md explains the numbers.
Getting started : install, configure, connect an account and a tool.
Set up with an AI agent : exact steps a coding agent can follow.
Use it with your tools : setup for each coding tool and SDK.
Several accounts : routing, limits, quotas and remote access.
Configuration : the settings you are most likely to change.
Install : every install option, Docker and running as a service.
Moving from CLIProxyAPI and differences from CLIProxyAPI .
Compatibility and parity and benchmarks .
Keep access.api-keys set and server.host on 127.0.0.1 unless you need other machines to connect. Behind a tunnel or reverse proxy on the same machine, set server.trusted-proxies , or every internet client counts as local. The dashboard is built into the binary, and cliproxy-rs never downloads code to run. Running it safely explains each point, and SECURITY.md says how to report a vulnerability.
Go is faster on non-streaming throughput. On a 2-vCPU test machine with a local fake upstream (2026-10-03), CLIProxyAPI handled about 34% more non-streaming requests per second (1,568 against 1,168) and about 10% more plain streams (916 against 833). cliproxy-rs was faster on streams it translates between the Anthropic and OpenAI formats (793 against 564 per second, about 41% more) and used less than half of Go's memory: 17 MB at idle against 45 MB, and 25 to 50 MB under load against 58 to 104 MB. It answers its first request about 17 ms after launch, and the release binary is 36 MB against 69 MB. docs/BENCHMARKS.md has the method and every number.
The workspace is crates/cpa-core (config and account formats), crates/cpa-exec (one module per upstream provider), crates/cpa-translate (format translation), crates/cpa-server (routes, account selection and the Management API), crates/cliproxy (the binary) and ui/ (the dashboard, see ui/README.md ). The workspace has 935 tests, run against local mock upstreams only; CI runs the whole suite on every push, with Go and PostgreSQL installed so the tests that compare against them run too. A differential harness sends the same 57 cases to CLIProxyAPI and cliproxy-rs and compares the results.
cargo fmt --all --check
cargo clippy --workspace --all-targets --locked -- -D warnings
cargo test --workspace --locked
CONTRIBUTING.md explains how to send a change.
MIT. See LICENSE . CLIProxyAPI, which this project follows closely, is also MIT licensed.
Rust rewrite of CLIProxyAPI: drop-in compatible proxy for Claude, Codex, Gemini and OpenAI-compatible providers
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Rust rewrite of CLIProxyAPI: drop-in compatible proxy for Claude, Codex, Gemini and OpenAI-compatible providers - vayungodara/cliproxy-rs

GitHub - vayungodara/cliproxy-rs: Rust rewrite of CLIProxyAPI: drop-in compatible proxy for Claude, Codex, Gemini and OpenAI-compatible providers · GitHub
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
vayungodara
/
cliproxy-rs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
587 Commits 587 Commits Folders and files
.github .github bench bench crates crates docs docs harness harness ui ui .dockerignore .dockerignore .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md install.ps1 install.ps1 install.sh install.sh rustfmt.toml rustfmt.toml View all files Repository files navigation
cliproxy-rs is a small server you run on your own computer. It lets your coding tools, such as Claude Code, Codex CLI or Cursor, use the AI subscriptions and API keys you already have, like a Claude Max plan or a ChatGPT plan, through one local address. A dashboard built into it shows which accounts work, how much of each plan's limits is left, and how to connect each tool.
It is a Rust rewrite of CLIProxyAPI . It reads the same config.yaml and the same account files, so you can switch between the two in either direction.
You pay for Claude and ChatGPT plans and want Claude Code, Codex CLI and your editor to share them, with the 5-hour and weekly limits of every account on one screen.
You have several Claude or Codex accounts and want requests spread across them, or moved to the next account when one reaches its limit.
A tool speaks only one API: the proxy translates between the Anthropic, OpenAI and Gemini formats, so an OpenAI-only tool can use a Claude account and the other way round.
You run CLIProxyAPI and want a smaller, faster-starting binary with the dashboard built in.
Run one command. On macOS or Linux:
curl -fsSL https://raw.githubusercontent.com/vayungodara/cliproxy-rs/master/install.sh | sh
On Windows, in PowerShell:
irm https: // raw.githubusercontent.com / vayungodara / cliproxy - rs / master / install.ps1 | iex
It installs the latest release after checking its checksum, writes a config with new keys to ~/.cliproxy-rs , starts the proxy in the background and opens the dashboard in your browser. The keys stay in ~/.cliproxy-rs/keys.env and are never printed.
In the dashboard, sign in with the CLIPROXY_MANAGEMENT_KEY line from keys.env and choose Connect account.
Open Use with tools and copy the settings for Claude Code, Codex CLI or another tool.
Run the same command again to upgrade; your config and keys stay as they are. docs/INSTALL.md shows how to start the proxy when you log in.
docs/GETTING-STARTED.md walks through the same steps in more detail.
Using a subscription outside its official app can break the provider's terms, and providers have suspended accounts for it. Whether to do that is your call and your risk.
Paste this into Claude Code, Codex or another coding agent on the computer where you want the proxy:
Install and set up cliproxy-rs on this computer by following
https://github.com/vayungodara/cliproxy-rs/blob/master/docs/AI-SETUP.md exactly.
Do not sign in to any of my accounts or print my keys; tell me when it is my turn.
The agent runs the installer, which picks a free port without touching an existing CLIProxyAPI and keeps the keys in a file only you can read. It checks that the proxy answers, then hands the account sign-in to you.
The proxy serves the Anthropic Messages API at http://127.0.0.1:8317 , the OpenAI API at http://127.0.0.1:8317/v1 and the Gemini API at http://127.0.0.1:8317 , all with your client key. For Claude Code:
export ANTHROPIC_BASE_URL=http://127.0.0.1:8317
export ANTHROPIC_AUTH_TOKEN=your-client-key
claude
docs/CLIENTS.md has copy-paste setups for Claude Code, Codex CLI, Gemini CLI, Amp, OpenCode, Factory Droid, Cline, Roo Code, Kilo Code, Cursor, Zed, Continue, Aider and the OpenAI, Anthropic and Gemini SDKs.
Sign in to as many Claude and Codex accounts as you have, and choose how the proxy rotates between them: round-robin , fill-first , weighted-round-robin , or soonest-reset (experimental, a cliproxy-rs addition: spend the account whose weekly window resets soonest first). Session affinity keeps each conversation on one account so the provider's prompt cache stays warm, and an account that hits a limit rests while the others take over. docs/MULTI-ACCOUNT.md covers routing, cooldowns, the quota view, per-account proxies, Codex over WebSocket and reaching the proxy from your other machines over Tailscale.
cliproxy-rs serves the same routes and the same v8 Management API as CLIProxyAPI, so the community apps built on CLIProxyAPI should work with it, for example CPA-Manager-Plus , CLIProxyAPI Quota Inspector , ZeroLimit , Quotio , VibeProxy and CCS . We have not tested all of them. Apps that talk to a running server need only its address and management key; apps that start their own bundled CLIProxyAPI need an option to use an existing server instead. If an app does not work with cliproxy-rs, please open an issue .
install.sh (macOS and Linux) and install.ps1 (Windows), the commands in the quick start. They install the release binary, set it up and start it; --binary-only ( -BinaryOnly on Windows) installs only the binary.
Release binaries for macOS (Apple silicon and Intel), Linux (x86_64 and arm64) and Windows (x86_64), with a SHA256SUMS file, on the releases page .
Docker, from the repository's Dockerfile .
Building from source, for contributors and platforms without a release binary.
docs/INSTALL.md covers each one, starting the proxy at login, and removing it.
These are not in cliproxy-rs yet. They are planned, in no fixed order:
Plugins: the host callbacks plugins use to call back into the server, calling plugins on the request path (plugin providers, sign-in, models and usage), and the plugin store. Today plugins load, serve their own routes and report quotas, but requests do not pass through them.
Home (cluster) mode: reporting usage, logs and in-flight requests back to Home, Home's KV storage, and syncing plugins managed by Home.
Image generation and editing through Codex accounts, and importing Vertex service accounts from the dashboard (the command line works).
The upstream request and response sections of request log files.
The rest of CLIProxyAPI's own test cases. The parity audit checked 1,687 Go routes, settings, flags and test suites: 798 (47%) are fully covered, 697 partly and 192 not yet. docs/PARITY.md explains the numbers.
Getting started : install, configure, connect an account and a tool.
Set up with an AI agent : exact steps a coding agent can follow.
Use it with your tools : setup for each coding tool and SDK.
Several accounts : routing, limits, quotas and remote access.
Configuration : the settings you are most likely to change.
Install : every install option, Docker and running as a service.
Moving from CLIProxyAPI and differences from CLIProxyAPI .
Compatibility and parity and benchmarks .
Keep access.api-keys set and server.host on 127.0.0.1 unless you need other machines to connect. Behind a tunnel or reverse proxy on the same machine, set server.trusted-proxies , or every internet client counts as local. The dashboard is built into the binary, and cliproxy-rs never downloads code to run. Running it safely explains each point, and SECURITY.md says how to report a vulnerability.
Go is faster on non-streaming throughput. On a 2-vCPU test machine with a local fake upstream (2026-10-03), CLIProxyAPI handled about 34% more non-streaming requests per second (1,568 against 1,168) and about 10% more plain streams (916 against 833). cliproxy-rs was faster on streams it translates between the Anthropic and OpenAI formats (793 against 564 per second, about 41% more) and used less than half of Go's memory: 17 MB at idle against 45 MB, and 25 to 50 MB under load against 58 to 104 MB. It answers its first request about 17 ms after launch, and the release binary is 36 MB against 69 MB. docs/BENCHMARKS.md has the method and every number.
The workspace is crates/cpa-core (config and account formats), crates/cpa-exec (one module per upstream provider), crates/cpa-translate (format translation), crates/cpa-server (routes, account selection and the Management API), crates/cliproxy (the binary) and ui/ (the dashboard, see ui/README.md ). The workspace has 935 tests, run against local mock upstreams only; CI runs the whole suite on every push, with Go and PostgreSQL installed so the tests that compare against them run too. A differential harness sends the same 57 cases to CLIProxyAPI and cliproxy-rs and compares the results.
cargo fmt --all --check
cargo clippy --workspace --all-targets --locked -- -D warnings
cargo test --workspace --locked
CONTRIBUTING.md explains how to send a change.
MIT. See LICENSE . CLIProxyAPI, which this project follows closely, is also MIT licensed.
Rust rewrite of CLIProxyAPI: drop-in compatible proxy for Claude, Codex, Gemini and OpenAI-compatible providers
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
