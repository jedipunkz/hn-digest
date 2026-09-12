---
source: "https://github.com/antonyrag/ragleap-core"
hn_url: "https://news.ycombinator.com/item?id=49669907"
title: "Show HN: RagLeap v0.4.0 – 46 AI roles, 8 connectors, 9 vector DBs"
article_title: "GitHub - antonyrag/ragleap-core: Production-ready AI Business OS — 46 autonomous AI Employees, Graph RAG (Neo4j), multi-channel, BYOK across 17 LLM providers plus custom/local endpoints (Ollama), automatic fallback. 8 libs (Python+Java). Self-hosted, MIT licensed, guardrailed autonomy for regulated\n[truncated]"
image: "https://opengraph.githubassets.com/55af7fdabe72d893aaad6a7e74cd99ff9c968b60be9d2352f74b06360d9f85d6/antonyrag/ragleap-core"
author: "antonyragleap"
captured_at: "2026-09-12T08:25:23Z"
capture_tool: "hn-digest"
hn_id: 49669907
score: 1
comments: 0
posted_at: "2026-09-12T07:33:30Z"
tags:
  - hacker-news
---

# Show HN: RagLeap v0.4.0 – 46 AI roles, 8 connectors, 9 vector DBs

- HN: [49669907](https://news.ycombinator.com/item?id=49669907)
- Source: [github.com](https://github.com/antonyrag/ragleap-core)
- Score: 1
- Comments: 0
- Posted: 2026-09-12T07:33:30Z

## Translation

Title: Show HN: RagLeap v0.4.0 – 46 AI roles, 8 connectors, 9 vector DBs
Article title: GitHub - antonyrag/ragleap-core: Production-ready AI Business OS — 46 autonomous AI Employees, Graph RAG (Neo4j), multi-channel, BYOK across 17 LLM providers plus custom/local endpoints (Ollama), automatic fallback. 8 libs (Python+Java). Self-hosted, MIT licensed, guardrailed autonomy for regulated
[truncated]
Description: Production-ready AI Business OS — 46 autonomous AI Employees, Graph RAG (Neo4j), multi-channel, BYOK across 17 LLM providers plus custom/local endpoints (Ollama), automatic fallback. 8 libs (Python+Java). Self-hosted, MIT licensed, guardrailed autonomy for regulated industries. - antonyrag/ragleap-c
[truncated]

Article text:
GitHub - antonyrag/ragleap-core: Production-ready AI Business OS — 46 autonomous AI Employees, Graph RAG (Neo4j), multi-channel, BYOK across 17 LLM providers plus custom/local endpoints (Ollama), automatic fallback. 8 libs (Python+Java). Self-hosted, MIT licensed, guardrailed autonomy for regulated industries. · GitHub
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
antonyrag
/
ragleap-core
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
487 Commits 487 Commits Folders and files
.github .github _layouts _layouts assets assets benchmarks/ stt-neutral-10s benchmarks/ stt-neutral-10s channels channels core core db db evals evals examples examples java/ ragleap-rag java/ ragleap-rag packages packages tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GENERIC-K8S-CHART-PROPOSAL.md GENERIC-K8S-CHART-PROPOSAL.md LICENSE LICENSE MIGRATION.md MIGRATION.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md STUDENT_PROJECTS.md STUDENT_PROJECTS.md _config.yml _config.yml docker-compose.yml docker-compose.yml install.sh install.sh requirements.txt requirements.txt View all files Repository files navigation
Autonomous AI Agents — Not Just RAG.
RagLeap Core is the open-source engine behind RagLeap — a self-hosted, agentic system that runs your business from your own documents on your own server, with no vendor lock-in.
46 role-based AI Employees: 9 core generalist roles (AI Manager, Secretary, CEO, Sales, Support, HR, Finance, Marketing, Operations) plus 37 vertical-specific global roles (Recruiter, Real Estate Agent, Legal Intake, Healthcare Intake, Insurance Agent, and more) — full list in core/employees/defaults.py .
Self-learning, outcome-weighted memory
Auto-trigger workflows and escalation
Think-act-decide loop, not just retrieval
Quickstart · Docs · Website · Hosted Version · Packages
Not to be confused with install.ragleap.com — that's a separate, paid, license-gated self-hosted product (Free tier with a license key, up to Enterprise). ragleap-core (this repo) is MIT-licensed, completely free, and never requires a license key. If you cloned this repo, you're in the right place for a genuinely free, open-source RAG engine.
pip install ragleap-rag
⭐ If this helps you, please consider starring the repo — it genuinely helps more people find it.
Add ragleap-graph too if you want Neo4j-backed knowledge graph retrieval:
pip install ragleap-rag ragleap-graph
Add ragleap-vectorstores too if you want pluggable vector backends beyond ragleap-rag's built-in six (Chroma today):
pip install ragleap-rag ragleap-vectorstores[chroma]
# or, with uv
uv add ragleap-rag ragleap-vectorstores[chroma]
Deploying to Kubernetes? ragleap-ops ships live-tested manifests and a Helm chart for the full stack:
pip install ragleap-ops
# or, with uv
uv add ragleap-ops
Need a generic, reusable chart for your own app (not RagLeap-specific)? ragleap-app-chart takes an arbitrary services: list, not hardcoded names:
pip install ragleap-app-chart
# or, with uv
uv add ragleap-app-chart
Want metrics for your ragleap-ops deployment? ragleap-observability ships Prometheus + postgres_exporter , live-verified end-to-end against a real cluster:
pip install ragleap-observability
# or, with uv
uv add ragleap-observability
Or run the full self-hosted app (channels, web chat UI, Docker Compose) — see Quickstart below. Browse every package at packages.ragleap.com . Try it hands-on with the runnable scripts in examples/ — 01_ingest_and_query.py (upload a document, ask a question via the API) and 02_test_channel_directly.py (test channel answering logic without real bot credentials).
If a RAG chatbot answers questions, RagLeap runs your business
Most open-source RAG projects give you a toolkit — you still have to build the app, wire up a UI, add memory, and connect every channel yourself. RagLeap Core gives you a working chat engine out of the box, and the full RagLeap platform turns it into an AI that actually operates a business.
What makes RagLeap Core specifically different
This repo isn't a general-purpose RAG framework you assemble into something — it's the real, working engine that already powers a production AI business platform (see What's in the hosted version below). The code here is honest about being early, but it's extracted from something that already works in the real world, not built as a demo.
Open-source AI agent projects like OpenClaw took off for a specific reason: people wanted an assistant that runs on their own infrastructure , with their own keys , answering from the chat apps they already use — not a black box hosted by someone else. That same principle is what RagLeap Core is built on for business AI specifically.
Your keys, your infrastructure, your data. RagLeap Core never asks for a system API key. You bring your own Gemini key, you run your own PostgreSQL database, your documents never leave your server unless you choose the hosted version.
Chat is the interface, not a separate dashboard you have to learn. The same way OpenClaw meets people on WhatsApp, Telegram, and Slack, RagLeap's full platform meets business owners on the channels they already use — WhatsApp, Telegram, Discord, and real phone calls — not a new app they have to check.
A real, working system — not an abstract framework. This isn't a toolkit like LangChain where you assemble your own app from primitives. RagLeap Core is the actual chunking → embedding → retrieval → generation pipeline extracted from a production system that already answers real customer questions, at a company that already runs on it.
Built in public, honestly. This repo says clearly what's done and what isn't. No inflated claims, no vaporware Quickstart commands that don't work yet — the Roadmap reflects the real state of the code, updated as it progresses.
RagLeap Core is a document-grounded chat engine. Upload your documents, ask questions, get cited answers — self-hosted, on your own infrastructure, with your own API key.
WhatsApp, Telegram, and Discord bots are included in this repo too — single-tenant, .env-configured channel adapters that answer from the same document knowledge base. It is the foundation of RagLeap , a hosted AI business manager that adds Voice calling, multi-tenancy, a persistent memory system, and an executive-assistant layer on top of this same core engine.
If RagLeap (hosted) is the business, RagLeap Core is the engine room.
RagLeap Core is right for you if
✅ You want a self-hosted RAG chatbot with full control over your data
✅ You want to understand exactly how document retrieval and citation works, not use a black box
✅ You're comfortable running your own server and your own AI provider key
✅ You want to contribute to or extend an open document-QA engine
✅ You'd rather see the code than trust a vendor's word on data privacy
It's not...
It is...
A hosted product
Self-hosted software you run yourself
Multi-tenant, with persistent cross-session memory
Single-tenant — one bot, one document set, per deployment
A multi-tenant platform
WhatsApp/Telegram/Discord/Voice channel adapters included, single-tenant — multi-tenant routing lives in the hosted version
A no-code SaaS dashboard
A codebase you deploy and configure
Feature-complete with the hosted version
The foundational subset — see Roadmap
Features
📄 Document ingestion
Upload PDFs, text, and common document formats
🔍 RAG retrieval
Vector search over your documents via pgvector
💬 Chat with citations
Answers reference the source document, not a black box
🔌 Bring your own AI key
OpenAI, Gemini, Anthropic, or any OpenAI-compatible endpoint
🌐 Web chat widget
Embed a chat widget on any website
🐳 Docker-based setup
One-command local deployment
🕸️ Knowledge Graph (Neo4j)
Entity extraction and graph-boosted retrieval alongside vector search
🌍 Language detection
Auto-detects document and query language, applied across every channel
🔗 Integrations
Connect MySQL, PostgreSQL, MongoDB, REST APIs, Salesforce, HubSpot, Shopify, Google Sheets, Stripe
🔀 Hybrid search
Combines dense (vector) and sparse (full-text) retrieval via Reciprocal Rank Fusion
⚡ Streaming responses
Answers stream token-by-token instead of waiting for the full response
🔁 Provider fallback
Automatically retries with a backup LLM provider if the primary fails
💰 Token usage reporting
Real per-call token counts from the provider, plus context-size budget trimming
🧑‍💼 AI Employees
Role-based agents (46 default roles) with persistent business-context memory, wired into /chat via role=<role>
🔗 n8n workflow automation
Fire a webhook after the AI replies on WhatsApp/Telegram/Discord — no-code automations triggered directly from a conversation
Architecture
RagLeap Core is the foundation layer of the full RagLeap platform. Here's how it fits into the bigger picture:
flowchart TD
subgraph Hosted["RagLeap — Hosted Platform (locked)"]
H1["Manager AI"]
H2["Multi-tenant AI Employees + Manager AI integration"]
H3["Persistent Memory (cross-channel, cross-session)"]
H4["Multi-tenant Billing, Teams & Permissions"]
H5["Audit History / Compliance logging"]
H6["Embed Widget Control Center (white-label)"]
H7["Managed hosting, backups, SLA, support"]
end
subgraph Core["RagLeap Core — this repo (open)"]
WebUI["Web Chat UI"] --> ChatAPI["Chat API"]
ChatAPI --> WA["WhatsApp"]
ChatAPI --> TG["Telegram"]
ChatAPI --> DC["Discord"]
ChatAPI --> VC["Voice"]
WA --> N8N["n8n Workflow Trigger (fires after AI reply)"]
TG --> N8N
DC --> N8N
Employees["AI Employees (role context, learned memory)"] --> Provider
WA --> Ingest["Document Ingest"]
TG --> Ingest
DC --> Ingest
VC --> Ingest
WA --> RAG["RAG Retrieve"]
TG --> RAG
DC --> RAG
VC --> RAG
WA --> Provider["AI Provider Adapter"]
TG --> Provider
DC --> Provider
VC --> Provider
Ingest --> PG[("PostgreSQL + pgvector")]
RAG --> PG
Provider --> PG
PG --> Neo[("Neo4j (Knowledge Graph)")]
end
Core -. built on top of .-> Hosted
Loading
[locked] = commercial/hosted-only feature, not included in this repository. See below for the full breakdown.
ragleap-core/
├── core/ # RAG engine — chunking, embedding, retrieval, generation
│ ├── chunker.py
│ ├── embedding.py # Gemini embeddings (gemini-embedding-001, 3072-dim)
│ ├── retrieval.py # pgvector cosine search
│ ├── generation.py # 19-provider BYOK generation (Gemini, OpenAI, Anthropic, etc.)
│ ├── ingest.py # chunk -> embed -> store pipeline
│ ├── parsers.py # PDF/DOCX/TXT text extraction
│ ├── employees/ # AI Employees — roles, business profile, learned memory
│ ├── workflows.py # n8n workflow automation — webhook tr

[truncated]

## Original Extract

Production-ready AI Business OS — 46 autonomous AI Employees, Graph RAG (Neo4j), multi-channel, BYOK across 17 LLM providers plus custom/local endpoints (Ollama), automatic fallback. 8 libs (Python+Java). Self-hosted, MIT licensed, guardrailed autonomy for regulated industries. - antonyrag/ragleap-c
[truncated]

GitHub - antonyrag/ragleap-core: Production-ready AI Business OS — 46 autonomous AI Employees, Graph RAG (Neo4j), multi-channel, BYOK across 17 LLM providers plus custom/local endpoints (Ollama), automatic fallback. 8 libs (Python+Java). Self-hosted, MIT licensed, guardrailed autonomy for regulated industries. · GitHub
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
antonyrag
/
ragleap-core
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
487 Commits 487 Commits Folders and files
.github .github _layouts _layouts assets assets benchmarks/ stt-neutral-10s benchmarks/ stt-neutral-10s channels channels core core db db evals evals examples examples java/ ragleap-rag java/ ragleap-rag packages packages tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile GENERIC-K8S-CHART-PROPOSAL.md GENERIC-K8S-CHART-PROPOSAL.md LICENSE LICENSE MIGRATION.md MIGRATION.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md STUDENT_PROJECTS.md STUDENT_PROJECTS.md _config.yml _config.yml docker-compose.yml docker-compose.yml install.sh install.sh requirements.txt requirements.txt View all files Repository files navigation
Autonomous AI Agents — Not Just RAG.
RagLeap Core is the open-source engine behind RagLeap — a self-hosted, agentic system that runs your business from your own documents on your own server, with no vendor lock-in.
46 role-based AI Employees: 9 core generalist roles (AI Manager, Secretary, CEO, Sales, Support, HR, Finance, Marketing, Operations) plus 37 vertical-specific global roles (Recruiter, Real Estate Agent, Legal Intake, Healthcare Intake, Insurance Agent, and more) — full list in core/employees/defaults.py .
Self-learning, outcome-weighted memory
Auto-trigger workflows and escalation
Think-act-decide loop, not just retrieval
Quickstart · Docs · Website · Hosted Version · Packages
Not to be confused with install.ragleap.com — that's a separate, paid, license-gated self-hosted product (Free tier with a license key, up to Enterprise). ragleap-core (this repo) is MIT-licensed, completely free, and never requires a license key. If you cloned this repo, you're in the right place for a genuinely free, open-source RAG engine.
pip install ragleap-rag
⭐ If this helps you, please consider starring the repo — it genuinely helps more people find it.
Add ragleap-graph too if you want Neo4j-backed knowledge graph retrieval:
pip install ragleap-rag ragleap-graph
Add ragleap-vectorstores too if you want pluggable vector backends beyond ragleap-rag's built-in six (Chroma today):
pip install ragleap-rag ragleap-vectorstores[chroma]
# or, with uv
uv add ragleap-rag ragleap-vectorstores[chroma]
Deploying to Kubernetes? ragleap-ops ships live-tested manifests and a Helm chart for the full stack:
pip install ragleap-ops
# or, with uv
uv add ragleap-ops
Need a generic, reusable chart for your own app (not RagLeap-specific)? ragleap-app-chart takes an arbitrary services: list, not hardcoded names:
pip install ragleap-app-chart
# or, with uv
uv add ragleap-app-chart
Want metrics for your ragleap-ops deployment? ragleap-observability ships Prometheus + postgres_exporter , live-verified end-to-end against a real cluster:
pip install ragleap-observability
# or, with uv
uv add ragleap-observability
Or run the full self-hosted app (channels, web chat UI, Docker Compose) — see Quickstart below. Browse every package at packages.ragleap.com . Try it hands-on with the runnable scripts in examples/ — 01_ingest_and_query.py (upload a document, ask a question via the API) and 02_test_channel_directly.py (test channel answering logic without real bot credentials).
If a RAG chatbot answers questions, RagLeap runs your business
Most open-source RAG projects give you a toolkit — you still have to build the app, wire up a UI, add memory, and connect every channel yourself. RagLeap Core gives you a working chat engine out of the box, and the full RagLeap platform turns it into an AI that actually operates a business.
What makes RagLeap Core specifically different
This repo isn't a general-purpose RAG framework you assemble into something — it's the real, working engine that already powers a production AI business platform (see What's in the hosted version below). The code here is honest about being early, but it's extracted from something that already works in the real world, not built as a demo.
Open-source AI agent projects like OpenClaw took off for a specific reason: people wanted an assistant that runs on their own infrastructure , with their own keys , answering from the chat apps they already use — not a black box hosted by someone else. That same principle is what RagLeap Core is built on for business AI specifically.
Your keys, your infrastructure, your data. RagLeap Core never asks for a system API key. You bring your own Gemini key, you run your own PostgreSQL database, your documents never leave your server unless you choose the hosted version.
Chat is the interface, not a separate dashboard you have to learn. The same way OpenClaw meets people on WhatsApp, Telegram, and Slack, RagLeap's full platform meets business owners on the channels they already use — WhatsApp, Telegram, Discord, and real phone calls — not a new app they have to check.
A real, working system — not an abstract framework. This isn't a toolkit like LangChain where you assemble your own app from primitives. RagLeap Core is the actual chunking → embedding → retrieval → generation pipeline extracted from a production system that already answers real customer questions, at a company that already runs on it.
Built in public, honestly. This repo says clearly what's done and what isn't. No inflated claims, no vaporware Quickstart commands that don't work yet — the Roadmap reflects the real state of the code, updated as it progresses.
RagLeap Core is a document-grounded chat engine. Upload your documents, ask questions, get cited answers — self-hosted, on your own infrastructure, with your own API key.
WhatsApp, Telegram, and Discord bots are included in this repo too — single-tenant, .env-configured channel adapters that answer from the same document knowledge base. It is the foundation of RagLeap , a hosted AI business manager that adds Voice calling, multi-tenancy, a persistent memory system, and an executive-assistant layer on top of this same core engine.
If RagLeap (hosted) is the business, RagLeap Core is the engine room.
RagLeap Core is right for you if
✅ You want a self-hosted RAG chatbot with full control over your data
✅ You want to understand exactly how document retrieval and citation works, not use a black box
✅ You're comfortable running your own server and your own AI provider key
✅ You want to contribute to or extend an open document-QA engine
✅ You'd rather see the code than trust a vendor's word on data privacy
It's not...
It is...
A hosted product
Self-hosted software you run yourself
Multi-tenant, with persistent cross-session memory
Single-tenant — one bot, one document set, per deployment
A multi-tenant platform
WhatsApp/Telegram/Discord/Voice channel adapters included, single-tenant — multi-tenant routing lives in the hosted version
A no-code SaaS dashboard
A codebase you deploy and configure
Feature-complete with the hosted version
The foundational subset — see Roadmap
Features
📄 Document ingestion
Upload PDFs, text, and common document formats
🔍 RAG retrieval
Vector search over your documents via pgvector
💬 Chat with citations
Answers reference the source document, not a black box
🔌 Bring your own AI key
OpenAI, Gemini, Anthropic, or any OpenAI-compatible endpoint
🌐 Web chat widget
Embed a chat widget on any website
🐳 Docker-based setup
One-command local deployment
🕸️ Knowledge Graph (Neo4j)
Entity extraction and graph-boosted retrieval alongside vector search
🌍 Language detection
Auto-detects document and query language, applied across every channel
🔗 Integrations
Connect MySQL, PostgreSQL, MongoDB, REST APIs, Salesforce, HubSpot, Shopify, Google Sheets, Stripe
🔀 Hybrid search
Combines dense (vector) and sparse (full-text) retrieval via Reciprocal Rank Fusion
⚡ Streaming responses
Answers stream token-by-token instead of waiting for the full response
🔁 Provider fallback
Automatically retries with a backup LLM provider if the primary fails
💰 Token usage reporting
Real per-call token counts from the provider, plus context-size budget trimming
🧑‍💼 AI Employees
Role-based agents (46 default roles) with persistent business-context memory, wired into /chat via role=<role>
🔗 n8n workflow automation
Fire a webhook after the AI replies on WhatsApp/Telegram/Discord — no-code automations triggered directly from a conversation
Architecture
RagLeap Core is the foundation layer of the full RagLeap platform. Here's how it fits into the bigger picture:
flowchart TD
subgraph Hosted["RagLeap — Hosted Platform (locked)"]
H1["Manager AI"]
H2["Multi-tenant AI Employees + Manager AI integration"]
H3["Persistent Memory (cross-channel, cross-session)"]
H4["Multi-tenant Billing, Teams & Permissions"]
H5["Audit History / Compliance logging"]
H6["Embed Widget Control Center (white-label)"]
H7["Managed hosting, backups, SLA, support"]
end
subgraph Core["RagLeap Core — this repo (open)"]
WebUI["Web Chat UI"] --> ChatAPI["Chat API"]
ChatAPI --> WA["WhatsApp"]
ChatAPI --> TG["Telegram"]
ChatAPI --> DC["Discord"]
ChatAPI --> VC["Voice"]
WA --> N8N["n8n Workflow Trigger (fires after AI reply)"]
TG --> N8N
DC --> N8N
Employees["AI Employees (role context, learned memory)"] --> Provider
WA --> Ingest["Document Ingest"]
TG --> Ingest
DC --> Ingest
VC --> Ingest
WA --> RAG["RAG Retrieve"]
TG --> RAG
DC --> RAG
VC --> RAG
WA --> Provider["AI Provider Adapter"]
TG --> Provider
DC --> Provider
VC --> Provider
Ingest --> PG[("PostgreSQL + pgvector")]
RAG --> PG
Provider --> PG
PG --> Neo[("Neo4j (Knowledge Graph)")]
end
Core -. built on top of .-> Hosted
Loading
[locked] = commercial/hosted-only feature, not included in this repository. See below for the full breakdown.
ragleap-core/
├── core/ # RAG engine — chunking, embedding, retrieval, generation
│ ├── chunker.py
│ ├── embedding.py # Gemini embeddings (gemini-embedding-001, 3072-dim)
│ ├── retrieval.py # pgvector cosine search
│ ├── generation.py # 19-provider BYOK generation (Gemini, OpenAI, Anthropic, etc.)
│ ├── ingest.py # chunk -> embed -> store pipeline
│ ├── parsers.py # PDF/DOCX/TXT text extraction
│ ├── employees/ # AI Employees — roles, business profile, learned memory
│ ├── workflows.py # n8n workflow automation — webhook tr

[truncated]
