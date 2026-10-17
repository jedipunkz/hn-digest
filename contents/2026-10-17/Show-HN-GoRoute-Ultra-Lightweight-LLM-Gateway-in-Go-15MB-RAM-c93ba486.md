---
source: "https://github.com/ziuus/goroute"
hn_url: "https://news.ycombinator.com/item?id=50133090"
title: "Show HN: GoRoute – Ultra-Lightweight LLM Gateway in Go (<15MB RAM)"
article_title: "GitHub - ziuus/goroute: ⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers · GitHub"
image: "https://opengraph.githubassets.com/74ed993266b07eb8676241f70a5e8f41e28f4d6d4b26ef124db606d8d5965795/ziuus/goroute"
author: "noelyt101"
captured_at: "2026-10-17T18:12:44Z"
capture_tool: "hn-digest"
hn_id: 50133090
score: 2
comments: 0
posted_at: "2026-10-17T17:29:10Z"
tags:
  - hacker-news
---

# Show HN: GoRoute – Ultra-Lightweight LLM Gateway in Go (<15MB RAM)

- HN: [50133090](https://news.ycombinator.com/item?id=50133090)
- Source: [github.com](https://github.com/ziuus/goroute)
- Score: 2
- Comments: 0
- Posted: 2026-10-17T17:29:10Z

## Translation

Title: Show HN: GoRoute – Ultra-Lightweight LLM Gateway in Go (<15MB RAM)
Article title: GitHub - ziuus/goroute: ⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers · GitHub
Description: ⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers - ziuus/goroute

Article text:
GitHub - ziuus/goroute: ⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers · GitHub
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
ziuus
goroute Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers
www.npmjs.com/package/goroute Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
18 Commits 18 Commits Folders and files
.github .github bin bin cmd/ goroute cmd/ goroute docs/ superpowers docs/ superpowers pkg pkg scripts scripts tests tests web web .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.sh install.sh package.json package.json View all files Repository files navigation
______ ____ __
/ ____/___ / __ \____ __ __/ /____
/ / __/ __ \/ /_/ / __ \/ / / / __/ _ \
/ /_/ / /_/ / _, _/ /_/ / /_/ / /_/ __/
\____/\____/_/ |_|\____/\__,_/\__/\___/
The Ultra-Lightweight LLM Gateway & Multi-Provider Router for Coding Agents
Sub-10ms Latency • < 5MB Idle RAM • Zero-Alloc Cache • 100% Native Pure Go
Quickstart •
Architecture •
Supported Agents •
Benchmarks •
Features •
Cheatsheet
GoRoute is an ultra-fast, single-binary LLM gateway engineered specifically for developers using autonomous AI coding agents ( OpenCode, Claude Code, Cursor, Codex CLI, Aider, Cline , etc.).
Instead of running bloated Node/Electron proxies that consume 600MB to 1.5GB of RAM , GoRoute delivers identical multi-provider pooling, Google Antigravity OAuth rotation, Anthropic-to-OpenAI protocol normalization, and response caching inside a tiny 5.1MB footprint .
┌─────────────────────────────────────────────────────────────┐
│ YOUR CODING AGENTS │
│ OpenCode • Claude Code • Cursor • Codex CLI • Aider • Roo │
└──────────────────────────────┬──────────────────────────────┘
│
http://127.0.0.1:8128
(OpenAI /v1 & Anthropic /v1/messages)
│
▼
┌───────────────────────────────────────────────────────────────────────────────────────────┐
│ ⚡ GOROUTE CORE │
│ │
│ ┌───────────────────────┐ ┌───────────────────────┐ ┌───────────────────────────────┐ │
│ │ Dual-Protocol SSE │ │ Zero-Alloc Cache │ │ Multi-Account PKCE Pool │ │
│ │ OpenAI ⇆ Anthropic │ │ 0ms Repeat Hits │ │ Antigravity Token Refresh │ │
│ └───────────────────────┘ └───────────────────────┘ └───────────────────────────────┘ │
│ │
│ ┌─────────────────────────────────────────────────────────────────────────────────────┐ │
│ │ Intelligent Failover & Quota Router │ │
│ │ (Auto-cascades across accounts on 429 quota or 503 limits) │ │
│ └─────────────────────────────────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────┬────────────────────────────────────────────┘
│
┌───────────────────────────────┼───────────────────────────────┐
▼ ▼ ▼
┌───────────────────────────┐ ┌───────────────────────────┐ ┌───────────────────────────┐
│ Google Antigravity Pool │ │ OpenAI / Claude Direct │ │ DeepSeek / TokenRouter │
│ 4x Accounts • Sonnet 4.6 │ │ OpenAI Compatible API │ │ Custom Fallback Gateway │
└───────────────────────────┘ └───────────────────────────┘ └───────────────────────────┘
🚀 Quickstart
📦 Option A: via npm / npx (Zero Go needed)
npm install -g goroute
# Or run directly via npx:
npx goroute setup opencode
⚡ Option B: 1-Line Native Installer (Linux & macOS)
curl -fsSL https://raw.githubusercontent.com/ziuus/goroute/main/install.sh | bash
🛠️ Option C: Go Toolchain
go install github.com/ziuus/goroute/cmd/goroute@latest
2. Connect Your Accounts
Migrate from OmniRoute in 1-Second:
goroute import omniroute
Instantly imports all Antigravity OAuth tokens, custom provider keys, and combo routers into encrypted local storage ( ~/.goroute/vault.json ).
goroute auth login antigravity
3. Connect Any Coding Agent
Auto-configure your favorite agent with one command:
# Configure OpenCode (writes config + registers provider credentials)
goroute setup opencode
# Configure Claude Code
goroute setup claude
# Configure OpenAI Codex CLI
goroute setup codex
# Configure ALL 14 supported agents in one click
goroute setup all
4. Run Seamlessly
Run GoRoute as a background service:
# Enable Linux background autostart (systemd user unit, ~5MB RAM):
goroute autostart enable
# Verify gateway status:
goroute status
Or launch any agent with zero configuration using the universal proxy injector:
goroute run opencode
goroute run claude
goroute run aider --model openai/claude-3-7-sonnet
📊 Benchmarks vs OmniRoute
Tested on Linux x86_64 under 500 concurrent streaming requests:
goroute setup automatically configures paths, headers, authentication keys, and provider manifests for:
Dual-Protocol Gateway :
POST /v1/chat/completions (OpenAI format + SSE chunks)
POST /v1/messages (Native Anthropic format + content_block_delta SSE streaming)
Zero-Alloc In-Memory Cache :
Repeated agent queries (e.g. title generation, linters, duplicate prompts) hit the internal cache and return X-Cache: HIT with 0ms latency and 0 upstream tokens .
Multimodal Vision & Media :
Automatically translates OpenAI image_url data URIs and Anthropic image base64 blocks into Gemini inlineData parts.
Smart Quota Cascade :
Automatically cascades across all connected Antigravity accounts on HTTP 429 (quota exhausted) or 403 (licensing limits) before falling back to backup providers.
Vector Embeddings Proxy :
POST /v1/embeddings supported for codebase indexing in Cursor, Continue, and local RAG workflows.
Role & Assistant Prefill Normalization :
Automatically handles consecutive turns and strips trailing assistant prefill messages to ensure 100% compatibility with Google Cloud Code / Vertex endpoints.
GoRoute provides pre-compiled, optimized static binaries for all major platforms:
# Gateway Management
goroute serve # Start gateway server (port 8128)
goroute serve --no-dashboard # Headless mode (< 10MB RAM)
goroute autostart enable # Enable system background service
goroute autostart status # Check service status
goroute autostart disable # Stop & disable service
goroute status # Health check & model listing
# Agent Setups
goroute setup all # Configure all 14 coding agents
goroute setup < agent > # Configure a specific agent (e.g. opencode, claude)
goroute run < agent > [args...] # Launch agent with proxy env variables
# Providers & Accounts
goroute auth login antigravity # Interactive PKCE OAuth login
goroute auth list # List connected provider accounts
goroute import omniroute # Import all accounts & keys from OmniRoute
goroute catalog # View 25+ built-in provider presets
📄 License
MIT © ziuus . Designed for developers who love speed.
⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers
www.npmjs.com/package/goroute Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers - ziuus/goroute

GitHub - ziuus/goroute: ⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers · GitHub
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
ziuus
goroute Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers
www.npmjs.com/package/goroute Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
18 Commits 18 Commits Folders and files
.github .github bin bin cmd/ goroute cmd/ goroute docs/ superpowers docs/ superpowers pkg pkg scripts scripts tests tests web web .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.sh install.sh package.json package.json View all files Repository files navigation
______ ____ __
/ ____/___ / __ \____ __ __/ /____
/ / __/ __ \/ /_/ / __ \/ / / / __/ _ \
/ /_/ / /_/ / _, _/ /_/ / /_/ / /_/ __/
\____/\____/_/ |_|\____/\__,_/\__/\___/
The Ultra-Lightweight LLM Gateway & Multi-Provider Router for Coding Agents
Sub-10ms Latency • < 5MB Idle RAM • Zero-Alloc Cache • 100% Native Pure Go
Quickstart •
Architecture •
Supported Agents •
Benchmarks •
Features •
Cheatsheet
GoRoute is an ultra-fast, single-binary LLM gateway engineered specifically for developers using autonomous AI coding agents ( OpenCode, Claude Code, Cursor, Codex CLI, Aider, Cline , etc.).
Instead of running bloated Node/Electron proxies that consume 600MB to 1.5GB of RAM , GoRoute delivers identical multi-provider pooling, Google Antigravity OAuth rotation, Anthropic-to-OpenAI protocol normalization, and response caching inside a tiny 5.1MB footprint .
┌─────────────────────────────────────────────────────────────┐
│ YOUR CODING AGENTS │
│ OpenCode • Claude Code • Cursor • Codex CLI • Aider • Roo │
└──────────────────────────────┬──────────────────────────────┘
│
http://127.0.0.1:8128
(OpenAI /v1 & Anthropic /v1/messages)
│
▼
┌───────────────────────────────────────────────────────────────────────────────────────────┐
│ ⚡ GOROUTE CORE │
│ │
│ ┌───────────────────────┐ ┌───────────────────────┐ ┌───────────────────────────────┐ │
│ │ Dual-Protocol SSE │ │ Zero-Alloc Cache │ │ Multi-Account PKCE Pool │ │
│ │ OpenAI ⇆ Anthropic │ │ 0ms Repeat Hits │ │ Antigravity Token Refresh │ │
│ └───────────────────────┘ └───────────────────────┘ └───────────────────────────────┘ │
│ │
│ ┌─────────────────────────────────────────────────────────────────────────────────────┐ │
│ │ Intelligent Failover & Quota Router │ │
│ │ (Auto-cascades across accounts on 429 quota or 503 limits) │ │
│ └─────────────────────────────────────────────────────────────────────────────────────┘ │
└──────────────────────────────────────────────┬────────────────────────────────────────────┘
│
┌───────────────────────────────┼───────────────────────────────┐
▼ ▼ ▼
┌───────────────────────────┐ ┌───────────────────────────┐ ┌───────────────────────────┐
│ Google Antigravity Pool │ │ OpenAI / Claude Direct │ │ DeepSeek / TokenRouter │
│ 4x Accounts • Sonnet 4.6 │ │ OpenAI Compatible API │ │ Custom Fallback Gateway │
└───────────────────────────┘ └───────────────────────────┘ └───────────────────────────┘
🚀 Quickstart
📦 Option A: via npm / npx (Zero Go needed)
npm install -g goroute
# Or run directly via npx:
npx goroute setup opencode
⚡ Option B: 1-Line Native Installer (Linux & macOS)
curl -fsSL https://raw.githubusercontent.com/ziuus/goroute/main/install.sh | bash
🛠️ Option C: Go Toolchain
go install github.com/ziuus/goroute/cmd/goroute@latest
2. Connect Your Accounts
Migrate from OmniRoute in 1-Second:
goroute import omniroute
Instantly imports all Antigravity OAuth tokens, custom provider keys, and combo routers into encrypted local storage ( ~/.goroute/vault.json ).
goroute auth login antigravity
3. Connect Any Coding Agent
Auto-configure your favorite agent with one command:
# Configure OpenCode (writes config + registers provider credentials)
goroute setup opencode
# Configure Claude Code
goroute setup claude
# Configure OpenAI Codex CLI
goroute setup codex
# Configure ALL 14 supported agents in one click
goroute setup all
4. Run Seamlessly
Run GoRoute as a background service:
# Enable Linux background autostart (systemd user unit, ~5MB RAM):
goroute autostart enable
# Verify gateway status:
goroute status
Or launch any agent with zero configuration using the universal proxy injector:
goroute run opencode
goroute run claude
goroute run aider --model openai/claude-3-7-sonnet
📊 Benchmarks vs OmniRoute
Tested on Linux x86_64 under 500 concurrent streaming requests:
goroute setup automatically configures paths, headers, authentication keys, and provider manifests for:
Dual-Protocol Gateway :
POST /v1/chat/completions (OpenAI format + SSE chunks)
POST /v1/messages (Native Anthropic format + content_block_delta SSE streaming)
Zero-Alloc In-Memory Cache :
Repeated agent queries (e.g. title generation, linters, duplicate prompts) hit the internal cache and return X-Cache: HIT with 0ms latency and 0 upstream tokens .
Multimodal Vision & Media :
Automatically translates OpenAI image_url data URIs and Anthropic image base64 blocks into Gemini inlineData parts.
Smart Quota Cascade :
Automatically cascades across all connected Antigravity accounts on HTTP 429 (quota exhausted) or 403 (licensing limits) before falling back to backup providers.
Vector Embeddings Proxy :
POST /v1/embeddings supported for codebase indexing in Cursor, Continue, and local RAG workflows.
Role & Assistant Prefill Normalization :
Automatically handles consecutive turns and strips trailing assistant prefill messages to ensure 100% compatibility with Google Cloud Code / Vertex endpoints.
GoRoute provides pre-compiled, optimized static binaries for all major platforms:
# Gateway Management
goroute serve # Start gateway server (port 8128)
goroute serve --no-dashboard # Headless mode (< 10MB RAM)
goroute autostart enable # Enable system background service
goroute autostart status # Check service status
goroute autostart disable # Stop & disable service
goroute status # Health check & model listing
# Agent Setups
goroute setup all # Configure all 14 coding agents
goroute setup < agent > # Configure a specific agent (e.g. opencode, claude)
goroute run < agent > [args...] # Launch agent with proxy env variables
# Providers & Accounts
goroute auth login antigravity # Interactive PKCE OAuth login
goroute auth list # List connected provider accounts
goroute import omniroute # Import all accounts & keys from OmniRoute
goroute catalog # View 25+ built-in provider presets
📄 License
MIT © ziuus . Designed for developers who love speed.
⚡ Ultra-efficient LLM Gateway & Multi-Provider Router in Go (< 15MB RAM) with OAuth & failovers
www.npmjs.com/package/goroute Topics
Readme MIT license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
