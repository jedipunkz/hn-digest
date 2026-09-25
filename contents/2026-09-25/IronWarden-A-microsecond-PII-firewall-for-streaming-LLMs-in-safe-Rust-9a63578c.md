---
source: "https://github.com/Somnerd/IronWarden"
hn_url: "https://news.ycombinator.com/item?id=49843753"
title: "IronWarden – A microsecond PII firewall for streaming LLMs in safe Rust"
article_title: "GitHub - Somnerd/IronWarden: High-performance sovereign AI reverse proxy & streaming ingress engine in Rust (Axum, Tokio) with real-time SSE token rehydration and HMAC audit chaining. · GitHub"
image: "https://opengraph.githubassets.com/dd483d3dc9c35b962c52c154914239336be75b098a425169d46a1ded9eb3a4fe/Somnerd/IronWarden"
author: "nikolasalexandr"
captured_at: "2026-09-25T13:17:54Z"
capture_tool: "hn-digest"
hn_id: 49843753
score: 2
comments: 0
posted_at: "2026-09-25T12:35:34Z"
tags:
  - hacker-news
---

# IronWarden – A microsecond PII firewall for streaming LLMs in safe Rust

- HN: [49843753](https://news.ycombinator.com/item?id=49843753)
- Source: [github.com](https://github.com/Somnerd/IronWarden)
- Score: 2
- Comments: 0
- Posted: 2026-09-25T12:35:34Z

## Translation

Title: IronWarden – A microsecond PII firewall for streaming LLMs in safe Rust
Article title: GitHub - Somnerd/IronWarden: High-performance sovereign AI reverse proxy & streaming ingress engine in Rust (Axum, Tokio) with real-time SSE token rehydration and HMAC audit chaining. · GitHub
Description: High-performance sovereign AI reverse proxy & streaming ingress engine in Rust (Axum, Tokio) with real-time SSE token rehydration and HMAC audit chaining. - Somnerd/IronWarden

Article text:
GitHub - Somnerd/IronWarden: High-performance sovereign AI reverse proxy & streaming ingress engine in Rust (Axum, Tokio) with real-time SSE token rehydration and HMAC audit chaining. · GitHub
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
Somnerd
/
IronWarden
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
372 Commits 372 Commits Folders and files
.github .github .jules .jules app app cli cli config config core core data/ knowledge data/ knowledge deploy/ helm/ ironwarden deploy/ helm/ ironwarden docs docs examples examples integration_tests integration_tests knowledge knowledge mcp mcp monitoring monitoring policies policies scripts scripts sdk/ python sdk/ python test_suites test_suites warden warden worker worker .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md BENCHMARKS.md BENCHMARKS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile GEMINI.md GEMINI.md LICENSE LICENSE LICENSE-COMMERCIAL.md LICENSE-COMMERCIAL.md MAINTAINER_NOTES.md MAINTAINER_NOTES.md MODELS.md MODELS.md PROXY_IMPLEMENTATION.md PROXY_IMPLEMENTATION.md QUICKSTART.md QUICKSTART.md README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml example.env example.env vscode_mcp_config.json vscode_mcp_config.json View all files Repository files navigation
High-Performance Sovereign AI Reverse Proxy & Privacy Firewall
IronWarden is a sovereign, ultra-low-latency AI security reverse proxy and PII firewall written in bare-metal Rust.
Point any OpenAI , Anthropic , or Ollama/vLLM SDK client at IronWarden to get real-time streaming PII redaction , sliding-window SSE token rehydration , prompt injection defense , and cryptographic HMAC-SHA256 audit chaining — with zero code changes in your application.
⚡ Technical Superiority & Latency Benchmark Matrix
Metric
IronWarden (Rust)
LiteLLM (Python)
Portkey (Node.js)
Kong AI Gateway (Lua/Go)
Language & Runtime
Bare-Metal Rust (Tokio/Axum)
Python (FastAPI/Uvicorn)
Node.js (TypeScript)
OpenResty (Lua) / Go
P95 Routing Overhead
<0.07 ms
18.5 ms
12.2 ms
3.4 ms
Streaming PII Redaction
Real-Time Sliding Window
Buffers Entire Stream
Buffers or regex post-hoc
Basic plugin / slow Lua regex
Max Concurrency (1 Core)
125,000+ req/s
~2,200 req/s
~4,800 req/s
~24,000 req/s
Memory Footprint
~18 MB
~140 MB
~110 MB
~85 MB
Data Sovereignty
100% Local / On-Prem / VPC
Local or Cloud
Cloud SaaS Dependent
Self-hosted or Cloud
Audit Log Integrity
Cryptographic HMAC-SHA256 Chaining
Plain Text JSON
Cloud SaaS Dashboard
Standard Access Logs
🏗️ Architecture
[ Client / Microservices / OpenAI & Anthropic SDKs ]
│
▼ (HTTP/2, Streaming SSE, JSON-RPC)
┌─────────────────────────────────────────────────────────────┐
│ IronWarden Core Gateway │
│ │
│ ┌──────────────────┐ ┌────────────────────────────────┐ │
│ │ Token Bucket │ │ Axum / Hyper High-Concurrency │ │
│ │ GCRA Rate Limit │───▶│ Non-Blocking Connection Pool │ │
│ └──────────────────┘ └────────────────────────────────┘ │
│ │ │
│ ▼ │
│ ┌────────────────────────────────────────────────────────┐ │
│ │ Streaming SSE Rehydration Engine │ │
│ │ • Sliding-window token reassembly across chunk splits │ │
│ │ • Zero-copy string normalization & homoglyph defense │ │
│ └────────────────────────────────────────────────────────┘ │
│ │ │
│ ▼ │
│ ┌────────────────────────────────────────────────────────┐ │
│ │ Multi-Tier PII & Security Gating │ │
│ │ • Layer 1: SIMD-Accelerated Aho-Corasick Regex Rules │ │
│ │ • Layer 2: ShadowNer Named Entity Recognition │ │
│ │ • Layer 3: Prompt Injection & Smuggling Guardrail │ │
│ └────────────────────────────────────────────────────────┘ │
│ │ │
│ ▼ │
│ ┌────────────────────────────────────────────────────────┐ │
│ │ Tamper-Proof Audit Chaining (HMAC-SHA256 Merkle Chain) │ │
│ │ • Verifiable cryptographic audit trail for EU AI Act │ │
│ └────────────────────────────────────────────────────────┘ │
└───────────────────────────────┬─────────────────────────────┘
│ (Redacted Outbound TX)
▼
[ Upstream LLMs: OpenAI / Anthropic / Local Ollama ]
⚡ Zero-Friction Quickstart
1. Run with Docker (1-Command Instant Start)
Spin up IronWarden in 5 seconds with zero configuration:
docker run -d --name ironwarden \
-p 8080:8080 \
-e UPSTREAM_LLM= " https://api.openai.com " \
-e WARDEN_MODE= " hybrid " \
ghcr.io/somnerd/ironwarden:latest
2. Verify with Streaming Curl
Send an LLM prompt containing sensitive PII and observe instant streaming token restoration with zero telemetry leakage:
curl -X POST http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-H " Authorization: Bearer YOUR_API_KEY " \
-d ' {
"model": "gpt-4o",
"messages": [
{"role": "user", "content": "Process payment for John Doe, SSN 000-12-3456, IBAN GR1201101250000000012345678."}
],
"stream": true
} '
3. Deploy with Docker Compose
docker compose up -d
4. Deploy to Kubernetes with Helm
helm install ironwarden ./deploy/helm/ironwarden \
--set secrets.wardenPepper= " 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef " \
--set secrets.openaiApiKey= " sk-... "
5. Build & Run from Source
git clone https://github.com/Somnerd/IronWarden.git
cd IronWarden
# Optional: Download ONNX NER weights (falls back to high-speed heuristic mode if omitted)
./scripts/setup_models.sh
# Run the gateway
export WARDEN_PEPPER= $( openssl rand -hex 16 )
cargo run --release -p app
🔌 Universal Drop-in SDK Compatibility
Simply set base_url to IronWarden's gateway endpoint:
from openai import OpenAI
# Point client to IronWarden — zero code modifications required
client = OpenAI (
api_key = "sk-mock-or-real" ,
base_url = "http://localhost:14141/v1" ,
default_headers = { "Authorization" : "Bearer <your-jwt-or-key>" }
)
response = client . chat . completions . create (
model = "gpt-4o" ,
messages = [
{ "role" : "user" , "content" : "Patient John Doe (SSN: 123-45-6789) shows elevated blood pressure." }
],
stream = True # Streaming supported natively with real-time SSE token rehydration
)
for chunk in response :
if chunk . choices [ 0 ]. delta . content :
print ( chunk . choices [ 0 ]. delta . content , end = "" , flush = True )
# ✅ PII scrubbed before reaching upstream LLM
# ✅ HMAC-chained tamper-evident audit record logged
# ✅ PII seamlessly restored in the output stream
Anthropic Claude Python SDK
import anthropic
client = anthropic . Anthropic (
api_key = "sk-ant-..." ,
base_url = "http://localhost:14141" ,
default_headers = {
"Authorization" : "Bearer <your-jwt-or-key>" ,
"X-IronWarden-Upstream-Key" : "sk-ant-..."
}
)
message = client . messages . create (
model = "claude-3-5-sonnet-20241022" ,
max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "Customer Jane Smith (Email: jane@enterprise.com) requested a refund." }]
)
print ( message . content [ 0 ]. text )
Dynamic Upstream Routing Headers
Header
Description
Default
X-IronWarden-Target-URL
Explicitly overrides upstream URL per-request (e.g. http://localhost:11434/v1/chat/completions )
Inferred from model name
X-IronWarden-Upstream-Key
Per-request API key for upstream provider
OPENAI_API_KEY / ANTHROPIC_API_KEY
Automatic Model Routing:
claude-* ➔ Anthropic API ( https://api.anthropic.com/v1/messages )
llama* , mistral* , phi* , gemma* , qwen* ➔ Local Ollama ( http://localhost:11434/v1/chat/completions )
All other models ➔ OpenAI API ( https://api.openai.com/v1/chat/completions )
🛡️ Core Capabilities & Invariants
1. Real-Time Streaming SSE Token Rehydration
Unlike standard proxies that buffer the entire response to replace tokens (introducing massive latency and breaking streaming UI), IronWarden implements an asynchronous SSE sliding-window state machine ( SseRehydrator ). It dynamically stitches split tokens across partial HTTP chunks in under 0.04 ms per chunk.
2. Hybrid Intelligence PII Shield
Deterministic Layer (Aho-Corasick + Entropy Smuggling Protection) : Ultra-fast regex and entropy heuristics for Credit Cards, SSNs, Emails, Phone Numbers, IBANs, and International IDs (including Greek AMKA/AFM and EU identifiers).
Probabilistic Layer (Local ONNX NER) : In-process DistilBERT Named Entity Recognition for contextual Names, Organizations, and Locations.
3. Cryptographic Audit Vault & Strict Fail-Closed Invariants
AES-256-GCM Encryption : Prompt and redaction records are encrypted at rest using your cryptographic pepper.
HMAC-SHA256 Hash Chaining : Every log entry is cryptographically linked to the previous record with continuous full-chain integrity walk verification.
Fail-Closed Security : If storage fills up or audit logging fails, IronWarden physically halts upstream egress to prevent un-audited data leakage.
Pre-configured, zero-touch regulatory rule sets ready to deploy:
Middle East & GCC Sovereignty ( config/rules/me.yaml ): Saudi Arabia PDPL (SDAIA), UAE Federal Decree-Law No. 45/2021, Qatar. Emirates ID, Saudi National ID/Iqama, Saudi & UAE IBANs, GCC mobile numbers, Arabic name heuristics.
East Asia Sovereignty ( config/rules/east_asia.yaml ): China PIPL / CSL, Japan APPI, South Korea PIPA, Singapore PDPA. China Resident ID, USCC, China Mobile, Japan My Number, Korea RRN, Singapore NRIC.
India DPDP Act 2023 ( config/rules/in.yaml ): PAN cards, Aadhaar numbers, GSTIN, Voter ID (EPIC), Indian Passports, Indian Mobile.
GDPR & European Sovereignty ( config/rules/eu.yaml , config/rules/gr.yaml ): EU & Greek national IDs (AMKA, AFM), EU IBANs, Passports, Driving Licenses.
HIPAA ( config/rules/rules_medical.yaml ): Medical records, Patient IDs, MRNs, SSNs.
PCI-DSS ( config/rules/rules.yaml ): Primary Account Numbers (PANs), CVVs, track data.
5. Model Context Protocol (MCP) Server
IronWarden includes a native JSON-RPC 2.0 stdio MCP server for agentic AI architectures (Claude Desktop, Cursor, AI agents) with session isolation and prompt sanitization tools:
Measured using Criterion.rs with 1,000+ iterations per sample. See BENCHMARKS.md for full methodology.
📈 Observability & Grafana Dashboard
IronWarden includes native, production-grade observability:
Prometheus Metrics : GET /metrics exposes request counts, blocked prompt injections, redacted PII entities, and available concurrency permits.
Turnkey Grafana Dashboard : GET /grafana/dashboard exports the pre-configured Grafana dashboard JSON.
Structured Health Inspection : GET /health returns JSON uptime, permit availability, and system status.
Launch IronWarden + Prometheus + Grafana together:
docker compose -f monitoring/docker-compose.monitoring.yml up -d
Visit http://localhost:3000 (admin/admin) to view real-time gatewa

[truncated]

## Original Extract

High-performance sovereign AI reverse proxy & streaming ingress engine in Rust (Axum, Tokio) with real-time SSE token rehydration and HMAC audit chaining. - Somnerd/IronWarden

GitHub - Somnerd/IronWarden: High-performance sovereign AI reverse proxy & streaming ingress engine in Rust (Axum, Tokio) with real-time SSE token rehydration and HMAC audit chaining. · GitHub
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
Somnerd
/
IronWarden
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
372 Commits 372 Commits Folders and files
.github .github .jules .jules app app cli cli config config core core data/ knowledge data/ knowledge deploy/ helm/ ironwarden deploy/ helm/ ironwarden docs docs examples examples integration_tests integration_tests knowledge knowledge mcp mcp monitoring monitoring policies policies scripts scripts sdk/ python sdk/ python test_suites test_suites warden warden worker worker .env.example .env.example .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml AGENTS.md AGENTS.md BENCHMARKS.md BENCHMARKS.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile GEMINI.md GEMINI.md LICENSE LICENSE LICENSE-COMMERCIAL.md LICENSE-COMMERCIAL.md MAINTAINER_NOTES.md MAINTAINER_NOTES.md MODELS.md MODELS.md PROXY_IMPLEMENTATION.md PROXY_IMPLEMENTATION.md QUICKSTART.md QUICKSTART.md README.md README.md RELEASE.md RELEASE.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml example.env example.env vscode_mcp_config.json vscode_mcp_config.json View all files Repository files navigation
High-Performance Sovereign AI Reverse Proxy & Privacy Firewall
IronWarden is a sovereign, ultra-low-latency AI security reverse proxy and PII firewall written in bare-metal Rust.
Point any OpenAI , Anthropic , or Ollama/vLLM SDK client at IronWarden to get real-time streaming PII redaction , sliding-window SSE token rehydration , prompt injection defense , and cryptographic HMAC-SHA256 audit chaining — with zero code changes in your application.
⚡ Technical Superiority & Latency Benchmark Matrix
Metric
IronWarden (Rust)
LiteLLM (Python)
Portkey (Node.js)
Kong AI Gateway (Lua/Go)
Language & Runtime
Bare-Metal Rust (Tokio/Axum)
Python (FastAPI/Uvicorn)
Node.js (TypeScript)
OpenResty (Lua) / Go
P95 Routing Overhead
<0.07 ms
18.5 ms
12.2 ms
3.4 ms
Streaming PII Redaction
Real-Time Sliding Window
Buffers Entire Stream
Buffers or regex post-hoc
Basic plugin / slow Lua regex
Max Concurrency (1 Core)
125,000+ req/s
~2,200 req/s
~4,800 req/s
~24,000 req/s
Memory Footprint
~18 MB
~140 MB
~110 MB
~85 MB
Data Sovereignty
100% Local / On-Prem / VPC
Local or Cloud
Cloud SaaS Dependent
Self-hosted or Cloud
Audit Log Integrity
Cryptographic HMAC-SHA256 Chaining
Plain Text JSON
Cloud SaaS Dashboard
Standard Access Logs
🏗️ Architecture
[ Client / Microservices / OpenAI & Anthropic SDKs ]
│
▼ (HTTP/2, Streaming SSE, JSON-RPC)
┌─────────────────────────────────────────────────────────────┐
│ IronWarden Core Gateway │
│ │
│ ┌──────────────────┐ ┌────────────────────────────────┐ │
│ │ Token Bucket │ │ Axum / Hyper High-Concurrency │ │
│ │ GCRA Rate Limit │───▶│ Non-Blocking Connection Pool │ │
│ └──────────────────┘ └────────────────────────────────┘ │
│ │ │
│ ▼ │
│ ┌────────────────────────────────────────────────────────┐ │
│ │ Streaming SSE Rehydration Engine │ │
│ │ • Sliding-window token reassembly across chunk splits │ │
│ │ • Zero-copy string normalization & homoglyph defense │ │
│ └────────────────────────────────────────────────────────┘ │
│ │ │
│ ▼ │
│ ┌────────────────────────────────────────────────────────┐ │
│ │ Multi-Tier PII & Security Gating │ │
│ │ • Layer 1: SIMD-Accelerated Aho-Corasick Regex Rules │ │
│ │ • Layer 2: ShadowNer Named Entity Recognition │ │
│ │ • Layer 3: Prompt Injection & Smuggling Guardrail │ │
│ └────────────────────────────────────────────────────────┘ │
│ │ │
│ ▼ │
│ ┌────────────────────────────────────────────────────────┐ │
│ │ Tamper-Proof Audit Chaining (HMAC-SHA256 Merkle Chain) │ │
│ │ • Verifiable cryptographic audit trail for EU AI Act │ │
│ └────────────────────────────────────────────────────────┘ │
└───────────────────────────────┬─────────────────────────────┘
│ (Redacted Outbound TX)
▼
[ Upstream LLMs: OpenAI / Anthropic / Local Ollama ]
⚡ Zero-Friction Quickstart
1. Run with Docker (1-Command Instant Start)
Spin up IronWarden in 5 seconds with zero configuration:
docker run -d --name ironwarden \
-p 8080:8080 \
-e UPSTREAM_LLM= " https://api.openai.com " \
-e WARDEN_MODE= " hybrid " \
ghcr.io/somnerd/ironwarden:latest
2. Verify with Streaming Curl
Send an LLM prompt containing sensitive PII and observe instant streaming token restoration with zero telemetry leakage:
curl -X POST http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-H " Authorization: Bearer YOUR_API_KEY " \
-d ' {
"model": "gpt-4o",
"messages": [
{"role": "user", "content": "Process payment for John Doe, SSN 000-12-3456, IBAN GR1201101250000000012345678."}
],
"stream": true
} '
3. Deploy with Docker Compose
docker compose up -d
4. Deploy to Kubernetes with Helm
helm install ironwarden ./deploy/helm/ironwarden \
--set secrets.wardenPepper= " 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef " \
--set secrets.openaiApiKey= " sk-... "
5. Build & Run from Source
git clone https://github.com/Somnerd/IronWarden.git
cd IronWarden
# Optional: Download ONNX NER weights (falls back to high-speed heuristic mode if omitted)
./scripts/setup_models.sh
# Run the gateway
export WARDEN_PEPPER= $( openssl rand -hex 16 )
cargo run --release -p app
🔌 Universal Drop-in SDK Compatibility
Simply set base_url to IronWarden's gateway endpoint:
from openai import OpenAI
# Point client to IronWarden — zero code modifications required
client = OpenAI (
api_key = "sk-mock-or-real" ,
base_url = "http://localhost:14141/v1" ,
default_headers = { "Authorization" : "Bearer <your-jwt-or-key>" }
)
response = client . chat . completions . create (
model = "gpt-4o" ,
messages = [
{ "role" : "user" , "content" : "Patient John Doe (SSN: 123-45-6789) shows elevated blood pressure." }
],
stream = True # Streaming supported natively with real-time SSE token rehydration
)
for chunk in response :
if chunk . choices [ 0 ]. delta . content :
print ( chunk . choices [ 0 ]. delta . content , end = "" , flush = True )
# ✅ PII scrubbed before reaching upstream LLM
# ✅ HMAC-chained tamper-evident audit record logged
# ✅ PII seamlessly restored in the output stream
Anthropic Claude Python SDK
import anthropic
client = anthropic . Anthropic (
api_key = "sk-ant-..." ,
base_url = "http://localhost:14141" ,
default_headers = {
"Authorization" : "Bearer <your-jwt-or-key>" ,
"X-IronWarden-Upstream-Key" : "sk-ant-..."
}
)
message = client . messages . create (
model = "claude-3-5-sonnet-20241022" ,
max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "Customer Jane Smith (Email: jane@enterprise.com) requested a refund." }]
)
print ( message . content [ 0 ]. text )
Dynamic Upstream Routing Headers
Header
Description
Default
X-IronWarden-Target-URL
Explicitly overrides upstream URL per-request (e.g. http://localhost:11434/v1/chat/completions )
Inferred from model name
X-IronWarden-Upstream-Key
Per-request API key for upstream provider
OPENAI_API_KEY / ANTHROPIC_API_KEY
Automatic Model Routing:
claude-* ➔ Anthropic API ( https://api.anthropic.com/v1/messages )
llama* , mistral* , phi* , gemma* , qwen* ➔ Local Ollama ( http://localhost:11434/v1/chat/completions )
All other models ➔ OpenAI API ( https://api.openai.com/v1/chat/completions )
🛡️ Core Capabilities & Invariants
1. Real-Time Streaming SSE Token Rehydration
Unlike standard proxies that buffer the entire response to replace tokens (introducing massive latency and breaking streaming UI), IronWarden implements an asynchronous SSE sliding-window state machine ( SseRehydrator ). It dynamically stitches split tokens across partial HTTP chunks in under 0.04 ms per chunk.
2. Hybrid Intelligence PII Shield
Deterministic Layer (Aho-Corasick + Entropy Smuggling Protection) : Ultra-fast regex and entropy heuristics for Credit Cards, SSNs, Emails, Phone Numbers, IBANs, and International IDs (including Greek AMKA/AFM and EU identifiers).
Probabilistic Layer (Local ONNX NER) : In-process DistilBERT Named Entity Recognition for contextual Names, Organizations, and Locations.
3. Cryptographic Audit Vault & Strict Fail-Closed Invariants
AES-256-GCM Encryption : Prompt and redaction records are encrypted at rest using your cryptographic pepper.
HMAC-SHA256 Hash Chaining : Every log entry is cryptographically linked to the previous record with continuous full-chain integrity walk verification.
Fail-Closed Security : If storage fills up or audit logging fails, IronWarden physically halts upstream egress to prevent un-audited data leakage.
Pre-configured, zero-touch regulatory rule sets ready to deploy:
Middle East & GCC Sovereignty ( config/rules/me.yaml ): Saudi Arabia PDPL (SDAIA), UAE Federal Decree-Law No. 45/2021, Qatar. Emirates ID, Saudi National ID/Iqama, Saudi & UAE IBANs, GCC mobile numbers, Arabic name heuristics.
East Asia Sovereignty ( config/rules/east_asia.yaml ): China PIPL / CSL, Japan APPI, South Korea PIPA, Singapore PDPA. China Resident ID, USCC, China Mobile, Japan My Number, Korea RRN, Singapore NRIC.
India DPDP Act 2023 ( config/rules/in.yaml ): PAN cards, Aadhaar numbers, GSTIN, Voter ID (EPIC), Indian Passports, Indian Mobile.
GDPR & European Sovereignty ( config/rules/eu.yaml , config/rules/gr.yaml ): EU & Greek national IDs (AMKA, AFM), EU IBANs, Passports, Driving Licenses.
HIPAA ( config/rules/rules_medical.yaml ): Medical records, Patient IDs, MRNs, SSNs.
PCI-DSS ( config/rules/rules.yaml ): Primary Account Numbers (PANs), CVVs, track data.
5. Model Context Protocol (MCP) Server
IronWarden includes a native JSON-RPC 2.0 stdio MCP server for agentic AI architectures (Claude Desktop, Cursor, AI agents) with session isolation and prompt sanitization tools:
Measured using Criterion.rs with 1,000+ iterations per sample. See BENCHMARKS.md for full methodology.
📈 Observability & Grafana Dashboard
IronWarden includes native, production-grade observability:
Prometheus Metrics : GET /metrics exposes request counts, blocked prompt injections, redacted PII entities, and available concurrency permits.
Turnkey Grafana Dashboard : GET /grafana/dashboard exports the pre-configured Grafana dashboard JSON.
Structured Health Inspection : GET /health returns JSON uptime, permit availability, and system status.
Launch IronWarden + Prometheus + Grafana together:
docker compose -f monitoring/docker-compose.monitoring.yml up -d
Visit http://localhost:3000 (admin/admin) to view real-time gatewa

[truncated]
