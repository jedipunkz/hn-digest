---
source: "https://github.com/swobuforge/swobu"
hn_url: "https://news.ycombinator.com/item?id=49689250"
title: "Show HN: Swobu – Local LLM Switchboard You Can Share over HTTPS"
article_title: "GitHub - swobuforge/swobu: LLM switchboard for Claude Code, Codex and AI agents — pool providers, cloud credits, regions and local models behind stable routes. · GitHub"
image: "https://opengraph.githubassets.com/464a6503628ca7473d39381ddeb6c7c33d0d0d0523b75ba707616548c3a6bb5f/swobuforge/swobu"
author: "metrofun"
captured_at: "2026-09-13T22:55:39Z"
capture_tool: "hn-digest"
hn_id: 49689250
score: 1
comments: 0
posted_at: "2026-09-13T22:12:51Z"
tags:
  - hacker-news
---

# Show HN: Swobu – Local LLM Switchboard You Can Share over HTTPS

- HN: [49689250](https://news.ycombinator.com/item?id=49689250)
- Source: [github.com](https://github.com/swobuforge/swobu)
- Score: 1
- Comments: 0
- Posted: 2026-09-13T22:12:51Z

## Translation

Title: Show HN: Swobu – Local LLM Switchboard You Can Share over HTTPS
Article title: GitHub - swobuforge/swobu: LLM switchboard for Claude Code, Codex and AI agents — pool providers, cloud credits, regions and local models behind stable routes. · GitHub
Description: LLM switchboard for Claude Code, Codex and AI agents — pool providers, cloud credits, regions and local models behind stable routes. - swobuforge/swobu
HN text: Swobu is a local TUI switchboard for Codex, Claude Code, PI, you name it. With stateful, provider-independent session lineage and model-name-based routing across way too many heterogeneous LLM providers. The same route can be shared and revoked over end-to-end HTTPS with another machine, a friend or your phone. Hell, you can even make an LLM mesh with one Swobu swobing remote Swobus. Anyway. One <10 MB Go binary. TLS terminates locally - relay is a dumb pipe. In other words, I’ve tried to factorize LLM traffic along orthogonal client, backend and network dimensions. Hope I find some product-market fit, because I don’t want to go back to job-market fit.

Article text:
GitHub - swobuforge/swobu: LLM switchboard for Claude Code, Codex and AI agents — pool providers, cloud credits, regions and local models behind stable routes. · GitHub
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
swobuforge
/
swobu
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
230 Commits 230 Commits Folders and files
.github .github assets/ readme assets/ readme cmd/ swobu cmd/ swobu contracts/ product-telemetry contracts/ product-telemetry docs docs internal internal scripts scripts shareprotocol shareprotocol .gitignore .gitignore CLA.md CLA.md COMMERCIAL-LICENSE.md COMMERCIAL-LICENSE.md COPYING COPYING COPYRIGHT COPYRIGHT LICENSE LICENSE Makefile Makefile README.es.md README.es.md README.id.md README.id.md README.ja.md README.ja.md README.ko.md README.ko.md README.md README.md README.pt-BR.md README.pt-BR.md README.ru.md README.ru.md README.uk.md README.uk.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md go.mod go.mod go.sum go.sum View all files Repository files navigation
English · 简体中文 · 日本語 · Português (Brasil) · Bahasa Indonesia · 한국어 · Русский · Español · Українська
Pool LLM capacity you already have.
A Swobu route looks like a model name to the agent; behind it can be your provider accounts, cloud regions, hosted endpoints, and local servers. Swobu handles routing, runtime fallback, and protocol translation across the paths you configure. It does not preflight compatibility: the real request reaches each configured target, and a failed attempt can advance to the next configured target.
Documentation · Quickstart · VS Code extension · Releases
Your agent chooses a model. Swobu chooses where it runs.
A Swobu route looks like a model to your agent.
Behind that name can be one endpoint, the same model available from several places, or a cross-provider pool.
The diagrams illustrate configurations. Choose models available from your providers.
claude-opus-5
│
├─ Anthropic / claude-opus-5
├─ AWS Bedrock / account A / claude-opus-5
└─ AWS Bedrock / account B / claude-opus-5
Keep using claude-opus-5 . Swobu can balance capacity and fail over underneath it.
Or make the model name describe a job:
codex-auto-review
│
├─ Deepseek / Deepseek V4 Flash
├─ Google / Gemini 3.7 Flash
└─ another review model
Or build a pool that deliberately crosses models and providers:
free
│
├─ Cerebras / Gemma 4 31B
├─ Groq / gpt-oss-20b
├─ LLM7 / default
├─ OpenRouter / free
├─ Mistral / Ministral 3B
├─ NVIDIA NIM / Nemotron Mini 4B
└─ Ollama / Qwen 3.8 27b
The model field your agent already understands becomes a programmable routing boundary.
curl -fsSL https://swobu.com/install.sh | sh
Windows PowerShell:
irm https: // swobu.com / install.ps1 | iex
The installer opens Cockpit , where you can add a provider, create a route,
and connect your first agent. It verifies the download, preserves an existing
standalone installation if setup fails, and leaves your shell profile and Swobu
data alone.
Already using a standalone installation? Update it with:
swobu update
Source, package-manager, and custom-directory installations remain owned by
the method that installed them.
Build your first route in five minutes →
Cockpit can configure supported clients for you.
swobu connect claude
swobu connect codex
swobu connect muse
swobu connect openclaw
swobu connect pi
swobu connect kilo
swobu connect opencode
swobu connect hermes
After that, your agent talks to Swobu. Provider configuration and routing stay behind the gateway.
What changes when the model name becomes a route?
It can represent a particular:
Put several targets in the same tier to balance across them.
Add fallback tiers to define what happens when preferred capacity is unavailable.
route: gpt-5.6-sol
primary
├─ Azure / westcentralus / gpt-5.6-sol
└─ Azure / westus2 / gpt-5.6-sol
fallback
└─ OpenAI / gpt-5.6-sol
The agent still asks for gpt-5.6-sol .
Routes don't have to preserve model identity.
A name such as review , cheap , free , or codex-auto-review can represent whatever capacity makes sense for that workload.
review
├─ Z.AI / GLM-5.3
├─ Kimi / Kimi-2.8
└─ Ollama / Qwen3-Coder
This lets different agents share routing policy without hard-coding provider configuration into each one.
Fail over without reconfiguring the agent
Quota exhausted. Region unavailable. Endpoint fails. Account hits a limit.
If an attempt fails, Swobu can try the next configured target in the route.
agent
│
│ model: gpt-5.6-sol
▼
Swobu
│
├─ Azure ────── unavailable
│
└─ OpenAI ──────── ✓
The route name does not change.
One boundary, multiple protocols
Claude Code ─┐
Codex ───────┤
Muse Code ───┤
OpenClaw ────┤
Pi ──────────┤
Kilo ────────┤
OpenCode ────┼──── Swobu ────┬─ OpenAI
Hermes ──────┤ ├─ Anthropic
Other agents ┘ ├─ Gemini
├─ AWS Bedrock
├─ Azure AI
├─ Cerebras
├─ Cloudflare
├─ Ollama
├─ LM Studio
├─ vLLM
└─ ...
Swobu currently supports provider integrations across protocols including:
Exact protocol and capability support varies by provider.
Routes control provider differences. They do not erase them.
Swobu supports local inference, frontier APIs, hyperscalers, specialized inference platforms, and aggregators.
Find providers and setup instructions in the documentation.
Give a remote agent one HTTPS endpoint and bearer without copying Swobu or
provider credentials to the recipient:
workspace dev
coding → Bedrock → Anthropic fallback
cheap → OpenRouter
local → Ollama
swobu share dev
The recipient can use coding , cheap , and local . If you change the targets
behind coding , their endpoint, bearer, and model name stay the same. Share one
route instead with swobu share dev/coding .
Shares default to one day. 7d , 30d , and never are free during preview.
The Owner Swobu process must be running: application TLS terminates there,
certificates renew automatically without changing the Share URL, and
swobu share revoke dev closes that workspace access.
Workspace and Route Share details →
Same model, multiple providers
Keep the model name the agent already uses while adding redundant capacity underneath it.
Combine recurring free capacity behind one model name.
Local first, cloud when needed
Prefer Ollama, LM Studio, or vLLM and fall through to hosted capacity according to policy.
Expose names such as codex-auto-review or claude-plan while changing the providers and models behind them independently.
Swobu runs locally and exposes the endpoint your agents connect to.
Your provider credentials stay at the gateway rather than being copied into every client.
No Swobu account is required for local use.
Operational telemetry is deliberately limited, and can be disabled.
Swobu publishes versioned binaries for Linux, macOS, and Windows, with SHA-256 checksums.
git clone https://github.com/swobuforge/swobu.git
cd swobu
make build
./.out/swobu --version
One model name. Any capacity underneath.
LLM switchboard for Claude Code, Codex and AI agents — pool providers, cloud credits, regions and local models behind stable routes.
Readme AGPL-3.0, AGPL-3.0 licenses found Security policy
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

LLM switchboard for Claude Code, Codex and AI agents — pool providers, cloud credits, regions and local models behind stable routes. - swobuforge/swobu

Swobu is a local TUI switchboard for Codex, Claude Code, PI, you name it. With stateful, provider-independent session lineage and model-name-based routing across way too many heterogeneous LLM providers. The same route can be shared and revoked over end-to-end HTTPS with another machine, a friend or your phone. Hell, you can even make an LLM mesh with one Swobu swobing remote Swobus. Anyway. One <10 MB Go binary. TLS terminates locally - relay is a dumb pipe. In other words, I’ve tried to factorize LLM traffic along orthogonal client, backend and network dimensions. Hope I find some product-market fit, because I don’t want to go back to job-market fit.

GitHub - swobuforge/swobu: LLM switchboard for Claude Code, Codex and AI agents — pool providers, cloud credits, regions and local models behind stable routes. · GitHub
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
swobuforge
/
swobu
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
230 Commits 230 Commits Folders and files
.github .github assets/ readme assets/ readme cmd/ swobu cmd/ swobu contracts/ product-telemetry contracts/ product-telemetry docs docs internal internal scripts scripts shareprotocol shareprotocol .gitignore .gitignore CLA.md CLA.md COMMERCIAL-LICENSE.md COMMERCIAL-LICENSE.md COPYING COPYING COPYRIGHT COPYRIGHT LICENSE LICENSE Makefile Makefile README.es.md README.es.md README.id.md README.id.md README.ja.md README.ja.md README.ko.md README.ko.md README.md README.md README.pt-BR.md README.pt-BR.md README.ru.md README.ru.md README.uk.md README.uk.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md go.mod go.mod go.sum go.sum View all files Repository files navigation
English · 简体中文 · 日本語 · Português (Brasil) · Bahasa Indonesia · 한국어 · Русский · Español · Українська
Pool LLM capacity you already have.
A Swobu route looks like a model name to the agent; behind it can be your provider accounts, cloud regions, hosted endpoints, and local servers. Swobu handles routing, runtime fallback, and protocol translation across the paths you configure. It does not preflight compatibility: the real request reaches each configured target, and a failed attempt can advance to the next configured target.
Documentation · Quickstart · VS Code extension · Releases
Your agent chooses a model. Swobu chooses where it runs.
A Swobu route looks like a model to your agent.
Behind that name can be one endpoint, the same model available from several places, or a cross-provider pool.
The diagrams illustrate configurations. Choose models available from your providers.
claude-opus-5
│
├─ Anthropic / claude-opus-5
├─ AWS Bedrock / account A / claude-opus-5
└─ AWS Bedrock / account B / claude-opus-5
Keep using claude-opus-5 . Swobu can balance capacity and fail over underneath it.
Or make the model name describe a job:
codex-auto-review
│
├─ Deepseek / Deepseek V4 Flash
├─ Google / Gemini 3.7 Flash
└─ another review model
Or build a pool that deliberately crosses models and providers:
free
│
├─ Cerebras / Gemma 4 31B
├─ Groq / gpt-oss-20b
├─ LLM7 / default
├─ OpenRouter / free
├─ Mistral / Ministral 3B
├─ NVIDIA NIM / Nemotron Mini 4B
└─ Ollama / Qwen 3.8 27b
The model field your agent already understands becomes a programmable routing boundary.
curl -fsSL https://swobu.com/install.sh | sh
Windows PowerShell:
irm https: // swobu.com / install.ps1 | iex
The installer opens Cockpit , where you can add a provider, create a route,
and connect your first agent. It verifies the download, preserves an existing
standalone installation if setup fails, and leaves your shell profile and Swobu
data alone.
Already using a standalone installation? Update it with:
swobu update
Source, package-manager, and custom-directory installations remain owned by
the method that installed them.
Build your first route in five minutes →
Cockpit can configure supported clients for you.
swobu connect claude
swobu connect codex
swobu connect muse
swobu connect openclaw
swobu connect pi
swobu connect kilo
swobu connect opencode
swobu connect hermes
After that, your agent talks to Swobu. Provider configuration and routing stay behind the gateway.
What changes when the model name becomes a route?
It can represent a particular:
Put several targets in the same tier to balance across them.
Add fallback tiers to define what happens when preferred capacity is unavailable.
route: gpt-5.6-sol
primary
├─ Azure / westcentralus / gpt-5.6-sol
└─ Azure / westus2 / gpt-5.6-sol
fallback
└─ OpenAI / gpt-5.6-sol
The agent still asks for gpt-5.6-sol .
Routes don't have to preserve model identity.
A name such as review , cheap , free , or codex-auto-review can represent whatever capacity makes sense for that workload.
review
├─ Z.AI / GLM-5.3
├─ Kimi / Kimi-2.8
└─ Ollama / Qwen3-Coder
This lets different agents share routing policy without hard-coding provider configuration into each one.
Fail over without reconfiguring the agent
Quota exhausted. Region unavailable. Endpoint fails. Account hits a limit.
If an attempt fails, Swobu can try the next configured target in the route.
agent
│
│ model: gpt-5.6-sol
▼
Swobu
│
├─ Azure ────── unavailable
│
└─ OpenAI ──────── ✓
The route name does not change.
One boundary, multiple protocols
Claude Code ─┐
Codex ───────┤
Muse Code ───┤
OpenClaw ────┤
Pi ──────────┤
Kilo ────────┤
OpenCode ────┼──── Swobu ────┬─ OpenAI
Hermes ──────┤ ├─ Anthropic
Other agents ┘ ├─ Gemini
├─ AWS Bedrock
├─ Azure AI
├─ Cerebras
├─ Cloudflare
├─ Ollama
├─ LM Studio
├─ vLLM
└─ ...
Swobu currently supports provider integrations across protocols including:
Exact protocol and capability support varies by provider.
Routes control provider differences. They do not erase them.
Swobu supports local inference, frontier APIs, hyperscalers, specialized inference platforms, and aggregators.
Find providers and setup instructions in the documentation.
Give a remote agent one HTTPS endpoint and bearer without copying Swobu or
provider credentials to the recipient:
workspace dev
coding → Bedrock → Anthropic fallback
cheap → OpenRouter
local → Ollama
swobu share dev
The recipient can use coding , cheap , and local . If you change the targets
behind coding , their endpoint, bearer, and model name stay the same. Share one
route instead with swobu share dev/coding .
Shares default to one day. 7d , 30d , and never are free during preview.
The Owner Swobu process must be running: application TLS terminates there,
certificates renew automatically without changing the Share URL, and
swobu share revoke dev closes that workspace access.
Workspace and Route Share details →
Same model, multiple providers
Keep the model name the agent already uses while adding redundant capacity underneath it.
Combine recurring free capacity behind one model name.
Local first, cloud when needed
Prefer Ollama, LM Studio, or vLLM and fall through to hosted capacity according to policy.
Expose names such as codex-auto-review or claude-plan while changing the providers and models behind them independently.
Swobu runs locally and exposes the endpoint your agents connect to.
Your provider credentials stay at the gateway rather than being copied into every client.
No Swobu account is required for local use.
Operational telemetry is deliberately limited, and can be disabled.
Swobu publishes versioned binaries for Linux, macOS, and Windows, with SHA-256 checksums.
git clone https://github.com/swobuforge/swobu.git
cd swobu
make build
./.out/swobu --version
One model name. Any capacity underneath.
LLM switchboard for Claude Code, Codex and AI agents — pool providers, cloud credits, regions and local models behind stable routes.
Readme AGPL-3.0, AGPL-3.0 licenses found Security policy
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
