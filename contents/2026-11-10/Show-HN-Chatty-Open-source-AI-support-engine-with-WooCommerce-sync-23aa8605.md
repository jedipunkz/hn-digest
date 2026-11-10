---
source: "https://github.com/PersonaliAI/chatty"
hn_url: "https://news.ycombinator.com/item?id=50439073"
title: "Show HN: Chatty – Open-source AI support engine with WooCommerce sync"
article_title: "GitHub - PersonaliAI/chatty: Open-source, self-hostable Zendesk and Intercom alternative for AI customer support: chat widget, shared inbox, real-time voice, Multimodal RAG, and MCP. · GitHub"
image: "https://repository-images.githubusercontent.com/1338232250/54e5d855-38a6-40f9-98a5-c065c902768f"
author: "TharinduDama"
captured_at: "2026-11-10T12:11:31Z"
capture_tool: "hn-digest"
hn_id: 50439073
score: 1
comments: 0
posted_at: "2026-11-10T11:57:55Z"
tags:
  - hacker-news
---

# Show HN: Chatty – Open-source AI support engine with WooCommerce sync

- HN: [50439073](https://news.ycombinator.com/item?id=50439073)
- Source: [github.com](https://github.com/PersonaliAI/chatty)
- Score: 1
- Comments: 0
- Posted: 2026-11-10T11:57:55Z

## Translation

Title: Show HN: Chatty – Open-source AI support engine with WooCommerce sync
Article title: GitHub - PersonaliAI/chatty: Open-source, self-hostable Zendesk and Intercom alternative for AI customer support: chat widget, shared inbox, real-time voice, Multimodal RAG, and MCP. · GitHub
Description: Open-source, self-hostable Zendesk and Intercom alternative for AI customer support: chat widget, shared inbox, real-time voice, Multimodal RAG, and MCP. - PersonaliAI/chatty

Article text:
GitHub - PersonaliAI/chatty: Open-source, self-hostable Zendesk and Intercom alternative for AI customer support: chat widget, shared inbox, real-time voice, Multimodal RAG, and MCP. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
PersonaliAI
chatty Public Notifications You must be signed in to change notification settings
Star 18 ( 18 ) You must be signed in to star a repository
Open-source, self-hostable Zendesk and Intercom alternative for AI customer support: chat widget, shared inbox, real-time voice, Multimodal RAG, and MCP.
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,573 Commits 1,573 Commits Folders and files
.agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .cursor/ rules .cursor/ rules .github .github backend backend docs @ 7f38549 docs @ 7f38549 frontend frontend integrations integrations plugins/ chatty-integration plugins/ chatty-integration screenshots screenshots sdks sdks .cursorrules .cursorrules .env.example .env.example .gcloudignore .gcloudignore .gitignore .gitignore .gitmodules .gitmodules CLAIMS_REGISTER.md CLAIMS_REGISTER.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DOCUMENTATION.md DOCUMENTATION.md LICENSE LICENSE README.es.md README.es.md README.ja.md README.ja.md README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md claims.md claims.md docker-compose.yml docker-compose.yml render.yaml render.yaml smithery.json smithery.json View all files Repository files navigation
English ·
简体中文 ·
日本語 ·
Español
Open-source AI agent trained on your store and content.
Answers customer questions, captures leads, and books meetings on website, WhatsApp & Slack.
🚀 Live Demo ·
☁️ Chatty Cloud (Free 14-day trial) ·
⚡ Quickstart (5 min) ·
📖 Documentation
Chatty is an open-source, self-hostable Zendesk and Intercom alternative for stores and modern businesses that need autonomous AI customer support with human takeover, WooCommerce sync, and calendar scheduling.
Run the application containers on your own host while keeping Supabase Auth,
Postgres, Storage, and Realtime managed. The same deployment contract works on
a VPS, Railway, Render, or another Docker host without touching the live
Supabase project.
Chatty Cloud (hosted) · Documentation · Data Processing (DPA) · Sub-processors · Performance Scorecard · Flow Builder · Widget Embed · Secure customer identity · Quick Start · Platform Runbook · MCP Server · AI Plugins · Omnichannel · WooCommerce
Live Dashboard Showcase (4K Retina)
Performance Scorecard (A-F Auditing)
Widget Integration (3-Line Embed)
Visual Flow Builder Engine & Conditional Runtime
WooCommerce Store Integration & Multimodal RAG
Self-Hosting, Step by Step
1. Create and configure Supabase
4. Configure environment variables
Local Development (without Docker)
Managed-Supabase platform runbook
MCP Server & AI Coding Plugins (Claude, Cursor, Codex)
Environment Variable Reference
GDPR & Enterprise Data Privacy
Every hosted chatbot SaaS charges per-seat or per-message and holds your conversation data. Chatty is open-core: run it yourself for free, or use Chatty Cloud - our own hosted version of this exact repo, if you'd rather skip the ops work. Either way you get the same feature set - streaming chat, RAG over your own documents, meeting booking, lead capture, WhatsApp/Slack channels, and a full MCP server so an AI agent can run the whole dashboard for you.
🖥️ Live Dashboard Showcase (4K Retina)
All 22 core modules are captured from live production ( chatty.personaliai.com/dashboard ) in crisp 4K Retina (3840×2160) resolution, representing real hydrated application state:
🔌 7-Channel Omnichannel Hub — Production-grade native messaging across WhatsApp Cloud API, Facebook Messenger, Instagram Direct, TikTok Business Messaging, Telegram Bot, Twilio Two-Way SMS (with automated TCPA opt-outs), and Inbound Support Email with quote-stripping.
🗺️ Visual Flow Builder Engine — Drag-and-drop interactive canvas ( @xyflow/react ) featuring declarative time-of-day condition branches (e.g. before 6:00 PM vs after 6:00 PM dynamic offers), lead capture, calendar booking, API webhooks, platform metadata detection, and AI Flow Copilot.
⚡ Live Human Takeover & Inbox — Real-time Server-Sent Events (SSE) streaming inbox enabling human agents to pause bot automation with one click and chat directly into WhatsApp, Instagram, TikTok, or SMS.
🛒 WooCommerce & Multimodal RAG — Direct WooCommerce store sync (order status lookup, live inventory, cart links) alongside PDF, DOCX, CSV, TXT ingestion, web crawling, and multimodal image context.
🎙️ LiveKit Voice Agent — Real-time WebRTC bidirectional voice conversations with 5 audio visualizers ( wave , bar , grid , radial , aura ), live speech transcription, pipeline & realtime multimodal models, and native embed support across Web, React, iOS, Android, Flutter, and React Native.
💬 Embeddable Chat Widget — One <script> tag, streaming SSE replies, crisp vector Shadow DOM rendering on any site with zero CSS bleed.
📚 RAG Knowledge Hub — Deep semantic retrieval powered by PostgreSQL pgvector , HNSW indexing, and hybrid lexical/vector search.
🛠️ Autonomous Tool-Calling — Schedules real meetings (Google Meet, Microsoft Teams, Zoom), qualifies leads, and checks calendars.
🔑 Bring Your Own Key (BYOK) — Free-tier Gemini by default, or per-bot OpenAI/Anthropic/Groq with AES-128 encryption.
🤖 81-Tool FastMCP Server — Control the entire dashboard programmatically from Claude Desktop, Cursor, or Codex via OAuth 2.0 / API Keys (including complete Voice Agent configuration, screenshot telemetry, and autonomous workflows).
📊 5-Pillar Scorecard — Independent A–F performance auditing across Conversion, Deflection, Accuracy, Reliability, and CSAT.
🎯 Proactive Campaigns — Behavioral triggers (exit-intent, time-on-page, scroll-depth) with timezone-aware cadence planning.
🐳 One-Command Self-Host — docker compose up , point to Supabase, and launch in 60 seconds.
Campaign configuration is persisted in chatty_campaigns . Widget deployments record impression , click , and conversion events through POST /api/widget/campaign-events with bot_id , campaign_id , and an optional idempotency_key (safe for retries). Recomputed analytics are accessible via GET /api/bots/{bot_id}/campaigns/{campaign_id}/analytics .
Campaign drafts support a validated schedule_config ( once , hourly , daily , or weekly with IANA timezone). POST /api/bots/{bot_id}/campaigns/audience-suggest generates bounded segments, intent thresholds, and rationales before saving. Delivery channels include web, email, WhatsApp, and SMS. Safety policies support overnight quiet hours ( quiet_hours: {"start": "22:00", "end": "08:00"} ). The background worker python -m app.workers.campaign_scheduler claims Redis idempotency keys for 24 hours to prevent duplicate dispatches. Provider-backed steps are expanded only for leads with explicit marketing consent and are then filtered by the same segment , returning_only , and min_intent_score rules used by the widget; qualification signals may be persisted in chatty_leads.custom_fields by an AI flow. Provider jobs additionally claim a hashed recipient/campaign frequency-cap key atomically before sending; cap hits are recorded as intentional suppressions, not retries. Each queued provider job carries its deterministic delivery key into the worker, which upserts queued , sent , failed , or suppressed state in chatty_campaign_deliveries ; campaign analytics reports delivery success rate and channel breakdown without exposing recipient addresses. Apply migration 20260930130000_campaign_delivery_ledger.sql before enabling provider campaigns in a new environment. Email delivery uses configured Resend or OneSignal credentials; unconfigured channels fail closed.
Flow Builder Safety & Deterministic Dry-Runs
The AI Flow Architect endpoint ( POST /api/flow/generate ) requires authenticated users with settings permission. Prompts are bounded to 4,000 characters, and generated graphs are validated for unique node IDs and resolvable edge references. Dry-runs enforce execution step budgets to prevent runaway loops and store immutable graph snapshots for deterministic replay. Webhook/action mappings use a safe {{input}} and {{context.path}} resolver; exact-value mappings preserve JSON types and unresolved fields are surfaced in the trace rather than silently sent.
Published widget flows execute their control nodes too: delay is capped at five minutes, condition accepts only the declarative input == , != , contains , starts_with , and ends_with grammar, and loop has a 100-iteration hard ceiling plus an explicit done , complete , or exit route. A webhook node is dispatched through POST /api/widget/flow/webhook , not from browser JavaScript. The API reloads the current active flow and permits only the selected published node's config.url (or config.webhook_url ), posts a bounded mapped body, adds an idempotency key, disables redirects, and pins DNS through the SSRF guard. Browser-supplied destinations, headers, and methods are never accepted. A webhook can set bounded max_attempts (1–5), backoff_ms (0–5,000), and timeout_ms (500–30,000); only connection errors, timeouts, and 5xx responses are retried. Use labeled success / true and error / failed / timeout edges to make the outcome path explicit. Operators can inspect an immutable run snapshot and trace at GET /api/bots/{bot_id}/flow/runs/{run_id} before replaying it.
Chatty provides native, industrial-standard bidirectional messaging across all major customer conversation surfaces from a single, unified workspace. Messages from every platform normalize into an internal event pipeline, allowing automated flows, the AI brain, and human operators to interact seamlessly.
Zero-Latency Ingestion Queue: Webhook endpoints perform HMAC verification and return HTTP 200 OK in <40ms, preventing Meta and TikTok from triggering exponential 3-second webhook retry storms while the AI reasons.
Smart UI Normalization: Interactive buttons automatically adapt across channels (rendered as native inline keyboard grids on Telegram, sliced to 3 quick-reply buttons on WhatsApp and TikTok, and transpiled into numbered text options like Reply [1] or [2] for SMS).
24-Hour Messaging Window Enforcement: Tracks last_inbound_at on Meta channels to prevent API rejection codes ( 131047 ), automatically falling back to pre-approved WhatsApp Business templates when the window expires.
TCPA Compliance Interceptor: Deterministically halts automation and registers opt-outs when users text regulatory keywords ( STOP , UNSUBSCRIBE , CANCEL , QUIT ).
Non-Destructive Secret Store: Credentials and webhoo

[truncated]

## Original Extract

Open-source, self-hostable Zendesk and Intercom alternative for AI customer support: chat widget, shared inbox, real-time voice, Multimodal RAG, and MCP. - PersonaliAI/chatty

GitHub - PersonaliAI/chatty: Open-source, self-hostable Zendesk and Intercom alternative for AI customer support: chat widget, shared inbox, real-time voice, Multimodal RAG, and MCP. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
PersonaliAI
chatty Public Notifications You must be signed in to change notification settings
Star 18 ( 18 ) You must be signed in to star a repository
Open-source, self-hostable Zendesk and Intercom alternative for AI customer support: chat widget, shared inbox, real-time voice, Multimodal RAG, and MCP.
Readme MIT license Code of conduct
Security policy Activity Custom properties Stars
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,573 Commits 1,573 Commits Folders and files
.agents/ plugins .agents/ plugins .claude-plugin .claude-plugin .cursor/ rules .cursor/ rules .github .github backend backend docs @ 7f38549 docs @ 7f38549 frontend frontend integrations integrations plugins/ chatty-integration plugins/ chatty-integration screenshots screenshots sdks sdks .cursorrules .cursorrules .env.example .env.example .gcloudignore .gcloudignore .gitignore .gitignore .gitmodules .gitmodules CLAIMS_REGISTER.md CLAIMS_REGISTER.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md DOCUMENTATION.md DOCUMENTATION.md LICENSE LICENSE README.es.md README.es.md README.ja.md README.ja.md README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md claims.md claims.md docker-compose.yml docker-compose.yml render.yaml render.yaml smithery.json smithery.json View all files Repository files navigation
English ·
简体中文 ·
日本語 ·
Español
Open-source AI agent trained on your store and content.
Answers customer questions, captures leads, and books meetings on website, WhatsApp & Slack.
🚀 Live Demo ·
☁️ Chatty Cloud (Free 14-day trial) ·
⚡ Quickstart (5 min) ·
📖 Documentation
Chatty is an open-source, self-hostable Zendesk and Intercom alternative for stores and modern businesses that need autonomous AI customer support with human takeover, WooCommerce sync, and calendar scheduling.
Run the application containers on your own host while keeping Supabase Auth,
Postgres, Storage, and Realtime managed. The same deployment contract works on
a VPS, Railway, Render, or another Docker host without touching the live
Supabase project.
Chatty Cloud (hosted) · Documentation · Data Processing (DPA) · Sub-processors · Performance Scorecard · Flow Builder · Widget Embed · Secure customer identity · Quick Start · Platform Runbook · MCP Server · AI Plugins · Omnichannel · WooCommerce
Live Dashboard Showcase (4K Retina)
Performance Scorecard (A-F Auditing)
Widget Integration (3-Line Embed)
Visual Flow Builder Engine & Conditional Runtime
WooCommerce Store Integration & Multimodal RAG
Self-Hosting, Step by Step
1. Create and configure Supabase
4. Configure environment variables
Local Development (without Docker)
Managed-Supabase platform runbook
MCP Server & AI Coding Plugins (Claude, Cursor, Codex)
Environment Variable Reference
GDPR & Enterprise Data Privacy
Every hosted chatbot SaaS charges per-seat or per-message and holds your conversation data. Chatty is open-core: run it yourself for free, or use Chatty Cloud - our own hosted version of this exact repo, if you'd rather skip the ops work. Either way you get the same feature set - streaming chat, RAG over your own documents, meeting booking, lead capture, WhatsApp/Slack channels, and a full MCP server so an AI agent can run the whole dashboard for you.
🖥️ Live Dashboard Showcase (4K Retina)
All 22 core modules are captured from live production ( chatty.personaliai.com/dashboard ) in crisp 4K Retina (3840×2160) resolution, representing real hydrated application state:
🔌 7-Channel Omnichannel Hub — Production-grade native messaging across WhatsApp Cloud API, Facebook Messenger, Instagram Direct, TikTok Business Messaging, Telegram Bot, Twilio Two-Way SMS (with automated TCPA opt-outs), and Inbound Support Email with quote-stripping.
🗺️ Visual Flow Builder Engine — Drag-and-drop interactive canvas ( @xyflow/react ) featuring declarative time-of-day condition branches (e.g. before 6:00 PM vs after 6:00 PM dynamic offers), lead capture, calendar booking, API webhooks, platform metadata detection, and AI Flow Copilot.
⚡ Live Human Takeover & Inbox — Real-time Server-Sent Events (SSE) streaming inbox enabling human agents to pause bot automation with one click and chat directly into WhatsApp, Instagram, TikTok, or SMS.
🛒 WooCommerce & Multimodal RAG — Direct WooCommerce store sync (order status lookup, live inventory, cart links) alongside PDF, DOCX, CSV, TXT ingestion, web crawling, and multimodal image context.
🎙️ LiveKit Voice Agent — Real-time WebRTC bidirectional voice conversations with 5 audio visualizers ( wave , bar , grid , radial , aura ), live speech transcription, pipeline & realtime multimodal models, and native embed support across Web, React, iOS, Android, Flutter, and React Native.
💬 Embeddable Chat Widget — One <script> tag, streaming SSE replies, crisp vector Shadow DOM rendering on any site with zero CSS bleed.
📚 RAG Knowledge Hub — Deep semantic retrieval powered by PostgreSQL pgvector , HNSW indexing, and hybrid lexical/vector search.
🛠️ Autonomous Tool-Calling — Schedules real meetings (Google Meet, Microsoft Teams, Zoom), qualifies leads, and checks calendars.
🔑 Bring Your Own Key (BYOK) — Free-tier Gemini by default, or per-bot OpenAI/Anthropic/Groq with AES-128 encryption.
🤖 81-Tool FastMCP Server — Control the entire dashboard programmatically from Claude Desktop, Cursor, or Codex via OAuth 2.0 / API Keys (including complete Voice Agent configuration, screenshot telemetry, and autonomous workflows).
📊 5-Pillar Scorecard — Independent A–F performance auditing across Conversion, Deflection, Accuracy, Reliability, and CSAT.
🎯 Proactive Campaigns — Behavioral triggers (exit-intent, time-on-page, scroll-depth) with timezone-aware cadence planning.
🐳 One-Command Self-Host — docker compose up , point to Supabase, and launch in 60 seconds.
Campaign configuration is persisted in chatty_campaigns . Widget deployments record impression , click , and conversion events through POST /api/widget/campaign-events with bot_id , campaign_id , and an optional idempotency_key (safe for retries). Recomputed analytics are accessible via GET /api/bots/{bot_id}/campaigns/{campaign_id}/analytics .
Campaign drafts support a validated schedule_config ( once , hourly , daily , or weekly with IANA timezone). POST /api/bots/{bot_id}/campaigns/audience-suggest generates bounded segments, intent thresholds, and rationales before saving. Delivery channels include web, email, WhatsApp, and SMS. Safety policies support overnight quiet hours ( quiet_hours: {"start": "22:00", "end": "08:00"} ). The background worker python -m app.workers.campaign_scheduler claims Redis idempotency keys for 24 hours to prevent duplicate dispatches. Provider-backed steps are expanded only for leads with explicit marketing consent and are then filtered by the same segment , returning_only , and min_intent_score rules used by the widget; qualification signals may be persisted in chatty_leads.custom_fields by an AI flow. Provider jobs additionally claim a hashed recipient/campaign frequency-cap key atomically before sending; cap hits are recorded as intentional suppressions, not retries. Each queued provider job carries its deterministic delivery key into the worker, which upserts queued , sent , failed , or suppressed state in chatty_campaign_deliveries ; campaign analytics reports delivery success rate and channel breakdown without exposing recipient addresses. Apply migration 20260930130000_campaign_delivery_ledger.sql before enabling provider campaigns in a new environment. Email delivery uses configured Resend or OneSignal credentials; unconfigured channels fail closed.
Flow Builder Safety & Deterministic Dry-Runs
The AI Flow Architect endpoint ( POST /api/flow/generate ) requires authenticated users with settings permission. Prompts are bounded to 4,000 characters, and generated graphs are validated for unique node IDs and resolvable edge references. Dry-runs enforce execution step budgets to prevent runaway loops and store immutable graph snapshots for deterministic replay. Webhook/action mappings use a safe {{input}} and {{context.path}} resolver; exact-value mappings preserve JSON types and unresolved fields are surfaced in the trace rather than silently sent.
Published widget flows execute their control nodes too: delay is capped at five minutes, condition accepts only the declarative input == , != , contains , starts_with , and ends_with grammar, and loop has a 100-iteration hard ceiling plus an explicit done , complete , or exit route. A webhook node is dispatched through POST /api/widget/flow/webhook , not from browser JavaScript. The API reloads the current active flow and permits only the selected published node's config.url (or config.webhook_url ), posts a bounded mapped body, adds an idempotency key, disables redirects, and pins DNS through the SSRF guard. Browser-supplied destinations, headers, and methods are never accepted. A webhook can set bounded max_attempts (1–5), backoff_ms (0–5,000), and timeout_ms (500–30,000); only connection errors, timeouts, and 5xx responses are retried. Use labeled success / true and error / failed / timeout edges to make the outcome path explicit. Operators can inspect an immutable run snapshot and trace at GET /api/bots/{bot_id}/flow/runs/{run_id} before replaying it.
Chatty provides native, industrial-standard bidirectional messaging across all major customer conversation surfaces from a single, unified workspace. Messages from every platform normalize into an internal event pipeline, allowing automated flows, the AI brain, and human operators to interact seamlessly.
Zero-Latency Ingestion Queue: Webhook endpoints perform HMAC verification and return HTTP 200 OK in <40ms, preventing Meta and TikTok from triggering exponential 3-second webhook retry storms while the AI reasons.
Smart UI Normalization: Interactive buttons automatically adapt across channels (rendered as native inline keyboard grids on Telegram, sliced to 3 quick-reply buttons on WhatsApp and TikTok, and transpiled into numbered text options like Reply [1] or [2] for SMS).
24-Hour Messaging Window Enforcement: Tracks last_inbound_at on Meta channels to prevent API rejection codes ( 131047 ), automatically falling back to pre-approved WhatsApp Business templates when the window expires.
TCPA Compliance Interceptor: Deterministically halts automation and registers opt-outs when users text regulatory keywords ( STOP , UNSUBSCRIBE , CANCEL , QUIT ).
Non-Destructive Secret Store: Credentials and webhoo

[truncated]
