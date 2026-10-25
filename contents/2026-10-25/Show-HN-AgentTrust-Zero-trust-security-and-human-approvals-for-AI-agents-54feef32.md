---
source: "https://github.com/SRN0129/Agenttrust"
hn_url: "https://news.ycombinator.com/item?id=50235909"
title: "Show HN: AgentTrust – Zero-trust security and human approvals for AI agents"
article_title: "GitHub - SRN0129/Agenttrust: Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents · GitHub"
image: "https://opengraph.githubassets.com/a8f4d0c5b3fd4598a676696276f49da58cd1c82118825198eaf3bf32a6bfa1f8/SRN0129/Agenttrust"
author: "SRN0129"
captured_at: "2026-10-25T21:29:03Z"
capture_tool: "hn-digest"
hn_id: 50235909
score: 2
comments: 0
posted_at: "2026-10-25T20:46:50Z"
tags:
  - hacker-news
---

# Show HN: AgentTrust – Zero-trust security and human approvals for AI agents

- HN: [50235909](https://news.ycombinator.com/item?id=50235909)
- Source: [github.com](https://github.com/SRN0129/Agenttrust)
- Score: 2
- Comments: 0
- Posted: 2026-10-25T20:46:50Z

## Translation

Title: Show HN: AgentTrust – Zero-trust security and human approvals for AI agents
Article title: GitHub - SRN0129/Agenttrust: Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents · GitHub
Description: Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents - SRN0129/Agenttrust

Article text:
GitHub - SRN0129/Agenttrust: Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents · GitHub
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
SRN0129
Agenttrust Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
9 Commits 9 Commits Folders and files
.github .github client client scripts scripts sdk sdk server server simulator simulator tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml package-lock.json package-lock.json package.json package.json render.yaml render.yaml tsconfig.json tsconfig.json tsconfig.server.json tsconfig.server.json vitest.config.ts vitest.config.ts View all files Repository files navigation
The Zero-Trust Security & Human Approval Layer for Autonomous AI Agents
Interoperable zero-trust permissions, cascading subagent delegation, semantic blast radius analysis, single-use replay defense, and cryptographically verifiable audit trails for AI agents.
💼 For Investors, Acquirers & Commercial Inquiries
This project and its intellectual property are available for commercial licensing or acquisition.
If you are an enterprise CISO, an AI platform looking to integrate an enterprise security layer, or an investor/acquirer:
📩 Open a Priority Inquiry: Click 💼 Submit Commercial Inquiry or connect via GitHub @SRN0129 .
When companies deploy autonomous AI agents (LangChain, CrewAI, AutoGen, OpenAI Assistants, Anthropic Tool Use), they face a critical security problem: unrestricted AI side effects . If an agent hallucinates, enters an infinite loop, or suffers prompt injection, it can drain budgets, delete databases, or transfer company funds.
AgentTrust acts as the independent security gatekeeper:
Deny-by-Default Execution : No AI agent can execute tools or touch databases without explicit policy authorization.
Cascading Subagent Delegation : Orchestrators can spawn worker subagents with strictly downscoped permissions (subagents cannot exceed parent capabilities).
Semantic Blast Radius Scoring : Evaluates the real-world irreversibility of actions before execution. Actions with high blast radius dynamically escalate to human review.
Human-in-the-Loop (HITL) Approvals with 1-Click Mobile Callback : Sensitive operations trigger approval tickets with 1-click decision URLs for Slack, Microsoft Teams, or mobile email.
Single-Use Replay Protection : Approved tickets carry unique cryptographic nonces that burn atomically on first use. Duplicate attempts trigger security anomalies.
Verifiable Audit Ledger & ECDSA P-256 Notarization : Every request and outcome is hash-chained ( prev_hash $\rightarrow$ entry_hash ) and periodically signed with an asymmetric ECDSA P-256 keypair for offline compliance verification.
Transparent Egress Proxy & Pre-Flight Prompt Scanner : Traps rogue outbound network calls and neutralizes prompt injections, jailbreaks, and sensitive data leaks.
Emergency Air-Gap Killswitch : 1-click panic button that revokes all agent credentials in $&lt; 15\text{ ms}$ and exports a forensic checkpoint.
Universal 1-Line Drop-in SDK ( wrapAgent ) : Injects zero-trust authorization, blast radius analysis, and audit logging into any LangChain, CrewAI, or custom agent in a single line.
Real-Time Goal Hijack & Semantic Drift Detection : Compares the agent's stated mission with proposed tool calls to catch indirect prompt injections and mission deviation before execution.
"Agent Chaos Monkey" Red-Team Simulator : Built-in adversarial penetration test battery issuing an Enterprise AI Trust Certificate (0-100%) with EU AI Act & SOC 2 compliance readiness metrics.
Zero-Knowledge Merkle Attestation Receipts : Generates self-contained cryptographic Merkle inclusion proofs that external auditors can verify offline without database access.
Micro-Budgeting & Velocity Circuit Breakers : Bounded credit allocations per run with automatic emergency throttling on rapid infinite loops ( $&gt; 30\text{ RPM}$ ) or spending spikes.
2. Threat Model & Security Boundaries (STRIDE)
Threat Category
Attack Vector
Countermeasure
Spoofing
Attacker impersonates a privileged agent.
Scoped API keys (hashed with SHA-256 at rest); instant revocation; delegation lineage tracking.
Tampering
Rogue actor alters audit events to conceal unauthorized operations.
Hash chaining ( prev_hash + SHA-256 entry hash) + ECDSA P-256 signed checkpoints.
Repudiation
Operator denies authorizing autonomous transactions.
Complete audit trail capturing agent ID, operator email, payload hash, and approval records.
Information Disclosure
Secrets leaked via frontend code or plaintext logs.
API keys returned once at registration. Parameter payloads hashed at rest.
Elevation of Privilege
Subagent attempts to obtain permissions outside parent scope.
Strict capability attenuation boundary check rejects unauthorized delegation.
Denial of Service
Agent enters infinite loop, exhausting quotas.
Sliding-minute rate limiters (RPM) and cumulative budget ceilings ($ USD).
Replay Attacks
Attacker intercepts an approved human execution ticket.
Single-use nonces atomically burned ( CONSUMED ) on first use. Duplicate attempts return HTTP 409.
Air-Gap Containment
Widespread autonomous prompt injection compromise.
Global emergency killswitch revoking all credentials across all agents in $&lt; 15\text{ ms}$ .
3. Automated Test Suite (41 Tests Passing)
AgentTrust features a rigorous test suite of 41 automated tests across 9 test suites covering all core security capabilities:
✓ tests/audit_integrity.test.ts (4 tests)
✓ tests/rate_limits_spend.test.ts (2 tests)
✓ tests/revocation.test.ts (4 tests)
✓ tests/authorization.test.ts (6 tests)
✓ tests/defensive_sandbox.test.ts (4 tests)
✓ tests/enterprise_defenses_v2.test.ts (7 tests)
✓ tests/approvals.test.ts (4 tests)
✓ tests/e2e_user_journey.test.ts (1 test)
✓ tests/enterprise_features.test.ts (9 tests)
Test Files 9 passed (9)
Tests 41 passed (41)
Run tests locally:
npm test
4. 1-Click Free Cloud Deployment
This repository includes pre-configured cloud deployment files ( render.yaml and Dockerfile ):
Go to Render.com and sign in with GitHub.
Click New + $\rightarrow$ Web Service $\rightarrow$ select SRN0129/Agenttrust .
Select Free Plan and click Deploy Web Service .
Your platform will be online 24/7 with a live HTTPS URL.
5. Enterprise REST API Reference
Endpoint
Method
Purpose
/api/v1/agents/register
POST
Register root agent, generate scoped credentials
/api/v1/agents/delegate
POST
Mint attenuated subagent with strictly downscoped capabilities
/api/v1/authorize
POST
Core policy gateway check with semantic blast radius scoring
/api/v1/proxy
POST
Transparent egress proxy inspecting and filtering outbound calls
/api/v1/approvals/pending
GET
List pending human approval requests
/api/v1/approvals/callback
GET
1-Click mobile push decision link for Slack, Teams, and email
/api/v1/tools/execute
POST
Sandboxed protected tool execution gateway
/api/v1/notarize/checkpoint
POST
Mint ECDSA P-256 signed audit checkpoint
/api/v1/compliance/report
GET
EU AI Act (Art. 12 & 14) and NIST AI RMF compliance assessment
/api/v1/system/killswitch
POST
Emergency air-gap killswitch revoking all agent credentials
6. License & Commercial Inquiries
This software is released under dual-licensing terms:
Open Source Evaluation : MIT License for personal and evaluation use.
Commercial & Enterprise Use : For proprietary integration, dedicated support, or acquisition, contact @SRN0129 .
Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents - SRN0129/Agenttrust

GitHub - SRN0129/Agenttrust: Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents · GitHub
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
SRN0129
Agenttrust Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
9 Commits 9 Commits Folders and files
.github .github client client scripts scripts sdk sdk server server simulator simulator tests tests .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml package-lock.json package-lock.json package.json package.json render.yaml render.yaml tsconfig.json tsconfig.json tsconfig.server.json tsconfig.server.json vitest.config.ts vitest.config.ts View all files Repository files navigation
The Zero-Trust Security & Human Approval Layer for Autonomous AI Agents
Interoperable zero-trust permissions, cascading subagent delegation, semantic blast radius analysis, single-use replay defense, and cryptographically verifiable audit trails for AI agents.
💼 For Investors, Acquirers & Commercial Inquiries
This project and its intellectual property are available for commercial licensing or acquisition.
If you are an enterprise CISO, an AI platform looking to integrate an enterprise security layer, or an investor/acquirer:
📩 Open a Priority Inquiry: Click 💼 Submit Commercial Inquiry or connect via GitHub @SRN0129 .
When companies deploy autonomous AI agents (LangChain, CrewAI, AutoGen, OpenAI Assistants, Anthropic Tool Use), they face a critical security problem: unrestricted AI side effects . If an agent hallucinates, enters an infinite loop, or suffers prompt injection, it can drain budgets, delete databases, or transfer company funds.
AgentTrust acts as the independent security gatekeeper:
Deny-by-Default Execution : No AI agent can execute tools or touch databases without explicit policy authorization.
Cascading Subagent Delegation : Orchestrators can spawn worker subagents with strictly downscoped permissions (subagents cannot exceed parent capabilities).
Semantic Blast Radius Scoring : Evaluates the real-world irreversibility of actions before execution. Actions with high blast radius dynamically escalate to human review.
Human-in-the-Loop (HITL) Approvals with 1-Click Mobile Callback : Sensitive operations trigger approval tickets with 1-click decision URLs for Slack, Microsoft Teams, or mobile email.
Single-Use Replay Protection : Approved tickets carry unique cryptographic nonces that burn atomically on first use. Duplicate attempts trigger security anomalies.
Verifiable Audit Ledger & ECDSA P-256 Notarization : Every request and outcome is hash-chained ( prev_hash $\rightarrow$ entry_hash ) and periodically signed with an asymmetric ECDSA P-256 keypair for offline compliance verification.
Transparent Egress Proxy & Pre-Flight Prompt Scanner : Traps rogue outbound network calls and neutralizes prompt injections, jailbreaks, and sensitive data leaks.
Emergency Air-Gap Killswitch : 1-click panic button that revokes all agent credentials in $&lt; 15\text{ ms}$ and exports a forensic checkpoint.
Universal 1-Line Drop-in SDK ( wrapAgent ) : Injects zero-trust authorization, blast radius analysis, and audit logging into any LangChain, CrewAI, or custom agent in a single line.
Real-Time Goal Hijack & Semantic Drift Detection : Compares the agent's stated mission with proposed tool calls to catch indirect prompt injections and mission deviation before execution.
"Agent Chaos Monkey" Red-Team Simulator : Built-in adversarial penetration test battery issuing an Enterprise AI Trust Certificate (0-100%) with EU AI Act & SOC 2 compliance readiness metrics.
Zero-Knowledge Merkle Attestation Receipts : Generates self-contained cryptographic Merkle inclusion proofs that external auditors can verify offline without database access.
Micro-Budgeting & Velocity Circuit Breakers : Bounded credit allocations per run with automatic emergency throttling on rapid infinite loops ( $&gt; 30\text{ RPM}$ ) or spending spikes.
2. Threat Model & Security Boundaries (STRIDE)
Threat Category
Attack Vector
Countermeasure
Spoofing
Attacker impersonates a privileged agent.
Scoped API keys (hashed with SHA-256 at rest); instant revocation; delegation lineage tracking.
Tampering
Rogue actor alters audit events to conceal unauthorized operations.
Hash chaining ( prev_hash + SHA-256 entry hash) + ECDSA P-256 signed checkpoints.
Repudiation
Operator denies authorizing autonomous transactions.
Complete audit trail capturing agent ID, operator email, payload hash, and approval records.
Information Disclosure
Secrets leaked via frontend code or plaintext logs.
API keys returned once at registration. Parameter payloads hashed at rest.
Elevation of Privilege
Subagent attempts to obtain permissions outside parent scope.
Strict capability attenuation boundary check rejects unauthorized delegation.
Denial of Service
Agent enters infinite loop, exhausting quotas.
Sliding-minute rate limiters (RPM) and cumulative budget ceilings ($ USD).
Replay Attacks
Attacker intercepts an approved human execution ticket.
Single-use nonces atomically burned ( CONSUMED ) on first use. Duplicate attempts return HTTP 409.
Air-Gap Containment
Widespread autonomous prompt injection compromise.
Global emergency killswitch revoking all credentials across all agents in $&lt; 15\text{ ms}$ .
3. Automated Test Suite (41 Tests Passing)
AgentTrust features a rigorous test suite of 41 automated tests across 9 test suites covering all core security capabilities:
✓ tests/audit_integrity.test.ts (4 tests)
✓ tests/rate_limits_spend.test.ts (2 tests)
✓ tests/revocation.test.ts (4 tests)
✓ tests/authorization.test.ts (6 tests)
✓ tests/defensive_sandbox.test.ts (4 tests)
✓ tests/enterprise_defenses_v2.test.ts (7 tests)
✓ tests/approvals.test.ts (4 tests)
✓ tests/e2e_user_journey.test.ts (1 test)
✓ tests/enterprise_features.test.ts (9 tests)
Test Files 9 passed (9)
Tests 41 passed (41)
Run tests locally:
npm test
4. 1-Click Free Cloud Deployment
This repository includes pre-configured cloud deployment files ( render.yaml and Dockerfile ):
Go to Render.com and sign in with GitHub.
Click New + $\rightarrow$ Web Service $\rightarrow$ select SRN0129/Agenttrust .
Select Free Plan and click Deploy Web Service .
Your platform will be online 24/7 with a live HTTPS URL.
5. Enterprise REST API Reference
Endpoint
Method
Purpose
/api/v1/agents/register
POST
Register root agent, generate scoped credentials
/api/v1/agents/delegate
POST
Mint attenuated subagent with strictly downscoped capabilities
/api/v1/authorize
POST
Core policy gateway check with semantic blast radius scoring
/api/v1/proxy
POST
Transparent egress proxy inspecting and filtering outbound calls
/api/v1/approvals/pending
GET
List pending human approval requests
/api/v1/approvals/callback
GET
1-Click mobile push decision link for Slack, Teams, and email
/api/v1/tools/execute
POST
Sandboxed protected tool execution gateway
/api/v1/notarize/checkpoint
POST
Mint ECDSA P-256 signed audit checkpoint
/api/v1/compliance/report
GET
EU AI Act (Art. 12 & 14) and NIST AI RMF compliance assessment
/api/v1/system/killswitch
POST
Emergency air-gap killswitch revoking all agent credentials
6. License & Commercial Inquiries
This software is released under dual-licensing terms:
Open Source Evaluation : MIT License for personal and evaluation use.
Commercial & Enterprise Use : For proprietary integration, dedicated support, or acquisition, contact @SRN0129 .
Enterprise Zero-Trust Security, Cascading Delegation, Human Approvals & Verifiable Audit Infrastructure for Autonomous AI Agents
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
