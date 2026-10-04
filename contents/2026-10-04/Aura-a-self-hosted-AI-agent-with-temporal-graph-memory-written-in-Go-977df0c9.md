---
source: "https://github.com/chetto1983/Aura"
hn_url: "https://news.ycombinator.com/item?id=49951895"
title: "Aura – a self-hosted AI agent with temporal graph memory, written in Go"
article_title: "GitHub - chetto1983/Aura: Self-hosted AI agent in Go with temporal graph memory, scheduled jobs, MCP tools and Telegram/WhatsApp channels. Docker Compose appliance, MIT. · GitHub"
image: "https://opengraph.githubassets.com/89855e5f4f1b762dc7ada4276bdb1b51d1f3a53c41f8434271a3a5c648d26709/chetto1983/Aura"
author: "chettto983"
captured_at: "2026-10-04T09:23:56Z"
capture_tool: "hn-digest"
hn_id: 49951895
score: 1
comments: 0
posted_at: "2026-10-04T08:42:20Z"
tags:
  - hacker-news
---

# Aura – a self-hosted AI agent with temporal graph memory, written in Go

- HN: [49951895](https://news.ycombinator.com/item?id=49951895)
- Source: [github.com](https://github.com/chetto1983/Aura)
- Score: 1
- Comments: 0
- Posted: 2026-10-04T08:42:20Z

## Translation

Title: Aura – a self-hosted AI agent with temporal graph memory, written in Go
Article title: GitHub - chetto1983/Aura: Self-hosted AI agent in Go with temporal graph memory, scheduled jobs, MCP tools and Telegram/WhatsApp channels. Docker Compose appliance, MIT. · GitHub
Description: Self-hosted AI agent in Go with temporal graph memory, scheduled jobs, MCP tools and Telegram/WhatsApp channels. Docker Compose appliance, MIT. - chetto1983/Aura

Article text:
GitHub - chetto1983/Aura: Self-hosted AI agent in Go with temporal graph memory, scheduled jobs, MCP tools and Telegram/WhatsApp channels. Docker Compose appliance, MIT. · GitHub
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
chetto1983
/
Aura
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
6,513 Commits 6,513 Commits Folders and files
.claude .claude .github .github .planning .planning caddy caddy cmd cmd deploy deploy docker docker docs docs finetune finetune internal internal observability observability packages/ create-aura packages/ create-aura public public scripts scripts searxng searxng services/ ingest services/ ingest spikes spikes web web .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .node-version .node-version .nvmrc .nvmrc AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NEXT.md NEXT.md README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md compose.cpu.yaml compose.cpu.yaml compose.vulkan.yaml compose.vulkan.yaml compose.yaml compose.yaml design-qa.md design-qa.md go.mod go.mod go.sum go.sum lefthook.yml lefthook.yml package.json package.json prd.md prd.md progress.md progress.md sqlc.yaml sqlc.yaml View all files Repository files navigation
A local-first, provider-neutral AI agent platform — in Go.
An agent for ongoing work: tools, document retrieval, temporal memory, scheduled
jobs, and a web cockpit on infrastructure you control.
What is Aura? · Features · Studio · Compare · Architecture · Quick Start · Docs · Development
Aura is a self-hosted, multi-user AI agent. Its Go binary hosts the runtime, tools,
CLI, Telegram gateway, and embedded web cockpit; each person signs in to their own
identity, with their own memory database, workspace and sandbox. Docker Compose runs
Postgres, ArcadeDB, Garage, embedding, ingestion, web search, voice and the bundled
integrations alongside it.
The model is chosen in the cockpit settings: OpenRouter (the default route), a
ChatGPT plan, the bundled local llama.cpp server, or Ollama. Local storage does not
make cloud inference offline: a cloud provider receives the context sent to its
model; with a local server nothing leaves the host for inference.
Hardware: a mini PC with 16 GB of RAM is enough. The default stack measured
7 GB with speech-to-text and text-to-speech running on a 16 GB mini PC
(2026-09-02). No local LLM runs by default: inference goes to the provider you choose.
A real run on a local stack: Aura stores the fact in its memory graph, loads the deferred
task tool and schedules the reminder; a new chat then answers from memory, with
provenance. Model replies were written by Claude through an OpenAI-compatible endpoint;
waiting time is trimmed.
Language
Go 1.27
Tests
Unit, property, race, leak, mutation, live integration, and browser tests
Test coverage
Owned-surface aggregate ≥85% , with package policies and separate live memory/sandbox coverage authorities
CI
build/vet/lint · CodeQL · -race + goleak · db/ArcadeDB/embed integration · MUSR two-identity E2E · web lint/test/mutation/Playwright · critical mutation ≥70% killed
Persistence
Postgres (sqlc, pgx) + ArcadeDB (graph memory, full-text + LSM vector index) + Garage (S3 object store)
Models
Default DeepSeek-V4 Flash via OpenRouter; also a ChatGPT plan, the bundled llama.cpp server (Gemma 4 12B QAT, localllm profile) or Ollama. The active profile (provider, model, budgets) is hot-reloaded from the cockpit settings, no restart
Distribution
edge tracks master; v1.0.2-rc1 is the latest tagged prerelease checked on 2026-10-03. See Releases for current availability
Key features
Streaming agent loop with shared step/time budgets and repeated-call controls to bound work.
Deferred tools and tool_search — discover tools and load their schemas when needed, including tools from mounted MCP servers.
Adaptive reasoning router — selects reasoning effort using the configured classifier and the active model's supported capabilities.
Full host terminal + filesystem tools — real operating power, with destructive-command approval gates and secret redaction.
Graph-native memory — facts, sources and validity windows in ArcadeDB; temporal paths return supporting evidence. Postgres-authoritative conversations have a derived recall projection and managed context compaction.
Document retrieval — indexed passages with source hashes and citations, plus access to the original file for calculations and whole-file tasks.
Self-extension — author and run skills, use bundled memory/PIM/WhatsApp integrations, and connect additional MCP servers.
Scheduler and self wake-ups — one task tool ( at | every | cron ) for reminders and agent_job runs, with job policy, operator controls and outcomes delivered to the owning conversation.
Per-identity sandbox — a full-capability box per operator (opt-in sandbox profile; gVisor runsc on native Linux), with deliverables handed back over the channel ( send_file ), never as a path.
Multi-user — Authula sign-in (password, plus a TOTP step for accounts enrolled in it), one isolated ArcadeDB database per identity enforced by the server, capability grants, and an admin audit view.
Multi-channel — CLI REPL, Telegram (voice/photo/docs/HITL), and a web cockpit over AG-UI/SSE with mid-turn steering, approvals, voice input/output, and live settings.
Studio — image and video generation, photo and video editing, and a multi-track video editor, all in the cockpit ( details ).
Bundled integrations — calendar/e-mail (PIM MCP, OAuth providers), WhatsApp (unofficial client), web search through a bundled SearXNG, snapshot share links to a conversation, and Cloudflare remote access.
The cockpit's creative workspace, per identity.
Generate images and video from a prompt over OpenRouter's media models. The model
picker shows each model's price (per image, per second or per million output tokens)
and the estimated cost before you press Generate. Options cover resolution, aspect
ratio, seed, and duration and sound for video; advanced inputs take a start frame, an
end frame and reference images from your library. Every generation lands in a
searchable history you can reuse or download. Generation needs the OpenRouter route.
Edit photos and clips in the browser: a photo editor (Filerobot) and a quick video
editor (trim, crop, rotate, audio) for any image or clip in a chat or in the Garage
library.
a video lane plus overlay lanes for titles and images, with transitions between clips;
per-clip transform (fill, fit, crop, flip, rotate), adjustments (opacity, brightness,
contrast, saturation, hue, blur), animations and speed;
audio lanes for an uploaded sound, a recorded voice, a text read aloud, or the sound
extracted from a clip, with noise reduction, fades and automatic ducking under speech;
undo and redo, saved projects, a mobile layout, and an export rendered in the browser
(video, or the audio alone as WAV).
A generated video opens in the editor with one click.
Checked on 2026-10-03 against each project's own documentation. Open WebUI and
LibreChat are mature, much larger projects; this table shows where Aura differs, not
that it is ahead.
Choose Open WebUI or LibreChat for a polished multi-model chat front end with SSO and
a large ecosystem. Choose Aura for a long-running personal agent that remembers over
time, works on a schedule and reaches you on Telegram or WhatsApp.
Transport & UX cmd/aura (CLI) · channels (+telegram) · agui (SSE) · webui (embedded SPA) · webauth (Authula) · setup · askuser
Agent runtime agent (LlmAgent, Budget, Events, hooks, workflow Seq/Par/Loop) · runner · swarm · steer
Tools & MCP agent/tools (registry, deferred, tool_search, fs/shell/web/skill) · agent/mcptools · mcp (+manager) · mcpoauth · sandbox
Intelligence llm (+openai_compat) · chatgptplan · semindex (embed-index core) · reasoningtrace · scoring · multimodal · mediagen
Capabilities web · skills · cron · onboarding · documents · share · retention
Persistence db (Postgres+sqlc) · arcadedb (memory + retrieval) · conversations · identity · objectstore · secret · settings
Observability obs · agent/panicobs · reasoningtrace · toolinvocations · cachemetrics
Documentation
Doc
For
docs/ARCHITECTURE.md
How the system is built — layers, turn lifecycle, invariants
docs/TECHNICAL_OVERVIEW.md
CTO / due-diligence overview — problem, differentiators, maturity
docs/CAPABILITIES.md
Capability matrix — shipped / in-progress / roadmap
docs/release-readiness.md
How a release is cut — the twelve-report exact-SHA gate, rollback rule, operational checks
docs/BACKUP-RESTORE.md
Backup schedules, recovery procedures, live validation and scope
CLAUDE.md · prd.md
Engineering guidance · product requirements (source of truth)
Deployment (Docker Compose appliance)
Aura is a self-hosted agent runtime packaged as a Docker Compose appliance. The
default stack brings up Aura (with its migration one-shot), Postgres, ArcadeDB and its
MCP, Garage, the local embedding sidecar, document ingestion, SearXNG, speech-to-text
and text-to-speech, the PIM and WhatsApp MCP sidecars, the Cloudflare tunnel
supervisor (idle until enabled), and Caddy in front of the Authula sign-in. Compose
profiles add the rest: localllm (a llama.cpp server with Gemma 4 12B QAT), ocr ,
observability (Prometheus, Tempo, Grafana) and sandbox (the Docker socket proxy
for per-identity boxes).
Releases. ghcr.io/chetto1983/aura:<tag> and the binary archives are published by
the Release workflow on a v* tag, and only after the exact-SHA Production
Readiness check passed for that commit ( docs/release-readiness.md ).
Check the
Releases page for the current tag
( v1.0.2-rc1 is the latest) and use it as vX.Y.Z below. Independently of
releases, every master push publishes the moving ghcr.io/chetto1983/aura:edge
image (plus an immutable master-<sha> tag) — the continuous-delivery channel a
default install tracks.
The interactive installer supports local installation or a Linux target over SSH:
npx create-aura-appliance
npx create-aura-appliance --mode remote
It requires Node.js 22.13 or newer on the workstation. The target needs at least
4 CPU cores, 14 GiB usable RAM, and 20 GiB free disk ; documents, models and backup
retention need additional capacity. The installer detects the embedding backend on the
target: CUDA for an NVIDIA GPU Docker can drive, otherwise Vulkan for an Intel or AMD GPU
exposing /dev/dri , otherwise CPU. See the installer guide
for supported targets and prerequisites. The npm installe

[truncated]

## Original Extract

Self-hosted AI agent in Go with temporal graph memory, scheduled jobs, MCP tools and Telegram/WhatsApp channels. Docker Compose appliance, MIT. - chetto1983/Aura

GitHub - chetto1983/Aura: Self-hosted AI agent in Go with temporal graph memory, scheduled jobs, MCP tools and Telegram/WhatsApp channels. Docker Compose appliance, MIT. · GitHub
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
chetto1983
/
Aura
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
6,513 Commits 6,513 Commits Folders and files
.claude .claude .github .github .planning .planning caddy caddy cmd cmd deploy deploy docker docker docs docs finetune finetune internal internal observability observability packages/ create-aura packages/ create-aura public public scripts scripts searxng searxng services/ ingest services/ ingest spikes spikes web web .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .node-version .node-version .nvmrc .nvmrc AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NEXT.md NEXT.md README.md README.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md compose.cpu.yaml compose.cpu.yaml compose.vulkan.yaml compose.vulkan.yaml compose.yaml compose.yaml design-qa.md design-qa.md go.mod go.mod go.sum go.sum lefthook.yml lefthook.yml package.json package.json prd.md prd.md progress.md progress.md sqlc.yaml sqlc.yaml View all files Repository files navigation
A local-first, provider-neutral AI agent platform — in Go.
An agent for ongoing work: tools, document retrieval, temporal memory, scheduled
jobs, and a web cockpit on infrastructure you control.
What is Aura? · Features · Studio · Compare · Architecture · Quick Start · Docs · Development
Aura is a self-hosted, multi-user AI agent. Its Go binary hosts the runtime, tools,
CLI, Telegram gateway, and embedded web cockpit; each person signs in to their own
identity, with their own memory database, workspace and sandbox. Docker Compose runs
Postgres, ArcadeDB, Garage, embedding, ingestion, web search, voice and the bundled
integrations alongside it.
The model is chosen in the cockpit settings: OpenRouter (the default route), a
ChatGPT plan, the bundled local llama.cpp server, or Ollama. Local storage does not
make cloud inference offline: a cloud provider receives the context sent to its
model; with a local server nothing leaves the host for inference.
Hardware: a mini PC with 16 GB of RAM is enough. The default stack measured
7 GB with speech-to-text and text-to-speech running on a 16 GB mini PC
(2026-09-02). No local LLM runs by default: inference goes to the provider you choose.
A real run on a local stack: Aura stores the fact in its memory graph, loads the deferred
task tool and schedules the reminder; a new chat then answers from memory, with
provenance. Model replies were written by Claude through an OpenAI-compatible endpoint;
waiting time is trimmed.
Language
Go 1.27
Tests
Unit, property, race, leak, mutation, live integration, and browser tests
Test coverage
Owned-surface aggregate ≥85% , with package policies and separate live memory/sandbox coverage authorities
CI
build/vet/lint · CodeQL · -race + goleak · db/ArcadeDB/embed integration · MUSR two-identity E2E · web lint/test/mutation/Playwright · critical mutation ≥70% killed
Persistence
Postgres (sqlc, pgx) + ArcadeDB (graph memory, full-text + LSM vector index) + Garage (S3 object store)
Models
Default DeepSeek-V4 Flash via OpenRouter; also a ChatGPT plan, the bundled llama.cpp server (Gemma 4 12B QAT, localllm profile) or Ollama. The active profile (provider, model, budgets) is hot-reloaded from the cockpit settings, no restart
Distribution
edge tracks master; v1.0.2-rc1 is the latest tagged prerelease checked on 2026-10-03. See Releases for current availability
Key features
Streaming agent loop with shared step/time budgets and repeated-call controls to bound work.
Deferred tools and tool_search — discover tools and load their schemas when needed, including tools from mounted MCP servers.
Adaptive reasoning router — selects reasoning effort using the configured classifier and the active model's supported capabilities.
Full host terminal + filesystem tools — real operating power, with destructive-command approval gates and secret redaction.
Graph-native memory — facts, sources and validity windows in ArcadeDB; temporal paths return supporting evidence. Postgres-authoritative conversations have a derived recall projection and managed context compaction.
Document retrieval — indexed passages with source hashes and citations, plus access to the original file for calculations and whole-file tasks.
Self-extension — author and run skills, use bundled memory/PIM/WhatsApp integrations, and connect additional MCP servers.
Scheduler and self wake-ups — one task tool ( at | every | cron ) for reminders and agent_job runs, with job policy, operator controls and outcomes delivered to the owning conversation.
Per-identity sandbox — a full-capability box per operator (opt-in sandbox profile; gVisor runsc on native Linux), with deliverables handed back over the channel ( send_file ), never as a path.
Multi-user — Authula sign-in (password, plus a TOTP step for accounts enrolled in it), one isolated ArcadeDB database per identity enforced by the server, capability grants, and an admin audit view.
Multi-channel — CLI REPL, Telegram (voice/photo/docs/HITL), and a web cockpit over AG-UI/SSE with mid-turn steering, approvals, voice input/output, and live settings.
Studio — image and video generation, photo and video editing, and a multi-track video editor, all in the cockpit ( details ).
Bundled integrations — calendar/e-mail (PIM MCP, OAuth providers), WhatsApp (unofficial client), web search through a bundled SearXNG, snapshot share links to a conversation, and Cloudflare remote access.
The cockpit's creative workspace, per identity.
Generate images and video from a prompt over OpenRouter's media models. The model
picker shows each model's price (per image, per second or per million output tokens)
and the estimated cost before you press Generate. Options cover resolution, aspect
ratio, seed, and duration and sound for video; advanced inputs take a start frame, an
end frame and reference images from your library. Every generation lands in a
searchable history you can reuse or download. Generation needs the OpenRouter route.
Edit photos and clips in the browser: a photo editor (Filerobot) and a quick video
editor (trim, crop, rotate, audio) for any image or clip in a chat or in the Garage
library.
a video lane plus overlay lanes for titles and images, with transitions between clips;
per-clip transform (fill, fit, crop, flip, rotate), adjustments (opacity, brightness,
contrast, saturation, hue, blur), animations and speed;
audio lanes for an uploaded sound, a recorded voice, a text read aloud, or the sound
extracted from a clip, with noise reduction, fades and automatic ducking under speech;
undo and redo, saved projects, a mobile layout, and an export rendered in the browser
(video, or the audio alone as WAV).
A generated video opens in the editor with one click.
Checked on 2026-10-03 against each project's own documentation. Open WebUI and
LibreChat are mature, much larger projects; this table shows where Aura differs, not
that it is ahead.
Choose Open WebUI or LibreChat for a polished multi-model chat front end with SSO and
a large ecosystem. Choose Aura for a long-running personal agent that remembers over
time, works on a schedule and reaches you on Telegram or WhatsApp.
Transport & UX cmd/aura (CLI) · channels (+telegram) · agui (SSE) · webui (embedded SPA) · webauth (Authula) · setup · askuser
Agent runtime agent (LlmAgent, Budget, Events, hooks, workflow Seq/Par/Loop) · runner · swarm · steer
Tools & MCP agent/tools (registry, deferred, tool_search, fs/shell/web/skill) · agent/mcptools · mcp (+manager) · mcpoauth · sandbox
Intelligence llm (+openai_compat) · chatgptplan · semindex (embed-index core) · reasoningtrace · scoring · multimodal · mediagen
Capabilities web · skills · cron · onboarding · documents · share · retention
Persistence db (Postgres+sqlc) · arcadedb (memory + retrieval) · conversations · identity · objectstore · secret · settings
Observability obs · agent/panicobs · reasoningtrace · toolinvocations · cachemetrics
Documentation
Doc
For
docs/ARCHITECTURE.md
How the system is built — layers, turn lifecycle, invariants
docs/TECHNICAL_OVERVIEW.md
CTO / due-diligence overview — problem, differentiators, maturity
docs/CAPABILITIES.md
Capability matrix — shipped / in-progress / roadmap
docs/release-readiness.md
How a release is cut — the twelve-report exact-SHA gate, rollback rule, operational checks
docs/BACKUP-RESTORE.md
Backup schedules, recovery procedures, live validation and scope
CLAUDE.md · prd.md
Engineering guidance · product requirements (source of truth)
Deployment (Docker Compose appliance)
Aura is a self-hosted agent runtime packaged as a Docker Compose appliance. The
default stack brings up Aura (with its migration one-shot), Postgres, ArcadeDB and its
MCP, Garage, the local embedding sidecar, document ingestion, SearXNG, speech-to-text
and text-to-speech, the PIM and WhatsApp MCP sidecars, the Cloudflare tunnel
supervisor (idle until enabled), and Caddy in front of the Authula sign-in. Compose
profiles add the rest: localllm (a llama.cpp server with Gemma 4 12B QAT), ocr ,
observability (Prometheus, Tempo, Grafana) and sandbox (the Docker socket proxy
for per-identity boxes).
Releases. ghcr.io/chetto1983/aura:<tag> and the binary archives are published by
the Release workflow on a v* tag, and only after the exact-SHA Production
Readiness check passed for that commit ( docs/release-readiness.md ).
Check the
Releases page for the current tag
( v1.0.2-rc1 is the latest) and use it as vX.Y.Z below. Independently of
releases, every master push publishes the moving ghcr.io/chetto1983/aura:edge
image (plus an immutable master-<sha> tag) — the continuous-delivery channel a
default install tracks.
The interactive installer supports local installation or a Linux target over SSH:
npx create-aura-appliance
npx create-aura-appliance --mode remote
It requires Node.js 22.13 or newer on the workstation. The target needs at least
4 CPU cores, 14 GiB usable RAM, and 20 GiB free disk ; documents, models and backup
retention need additional capacity. The installer detects the embedding backend on the
target: CUDA for an NVIDIA GPU Docker can drive, otherwise Vulkan for an Intel or AMD GPU
exposing /dev/dri , otherwise CPU. See the installer guide
for supported targets and prerequisites. The npm installe

[truncated]
