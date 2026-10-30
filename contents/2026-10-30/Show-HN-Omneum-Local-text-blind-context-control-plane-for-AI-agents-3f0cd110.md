---
source: "https://github.com/cole-h06/omneum"
hn_url: "https://news.ycombinator.com/item?id=50304917"
title: "Show HN: Omneum – Local, text-blind context control plane for AI agents"
article_title: "GitHub - cole-h06/omneum: Text-blind context control plane for AI agents. Select high-utility context under token budgets. · GitHub"
image: "https://opengraph.githubassets.com/266c7122045a5aee22663f92e6c3c9290e82433792cc657894c5ae6d2524cd58/cole-h06/omneum"
author: "cole_dev"
captured_at: "2026-10-30T15:36:24Z"
capture_tool: "hn-digest"
hn_id: 50304917
score: 1
comments: 0
posted_at: "2026-10-30T15:03:29Z"
tags:
  - hacker-news
---

# Show HN: Omneum – Local, text-blind context control plane for AI agents

- HN: [50304917](https://news.ycombinator.com/item?id=50304917)
- Source: [github.com](https://github.com/cole-h06/omneum)
- Score: 1
- Comments: 0
- Posted: 2026-10-30T15:03:29Z

## Translation

Title: Show HN: Omneum – Local, text-blind context control plane for AI agents
Article title: GitHub - cole-h06/omneum: Text-blind context control plane for AI agents. Select high-utility context under token budgets. · GitHub
Description: Text-blind context control plane for AI agents. Select high-utility context under token budgets. - cole-h06/omneum

Article text:
GitHub - cole-h06/omneum: Text-blind context control plane for AI agents. Select high-utility context under token budgets. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
cole-h06
omneum Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Text-blind context control plane for AI agents. Select high-utility context under token budgets.
Readme MIT license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
312 Commits 312 Commits Folders and files
.cursor-plugin .cursor-plugin .github/ workflows .github/ workflows assets assets benchmarks benchmarks deploy deploy docs docs evals evals examples examples migrations migrations omneum omneum research research src src tests tests tools tools .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md build.rs build.rs compose.yaml compose.yaml glama.json glama.json mcp.json mcp.json pyproject.toml pyproject.toml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt rust-toolchain.toml rust-toolchain.toml server.json server.json uv.lock uv.lock View all files Repository files navigation
Provenance-aware context control plane for AI agents.
Omneum is an open-source context control plane for AI agents, built to optimize the evidence entering an LLM's context window.
Agents accumulate context from retrieval, tools, memory, and previous sessions, but longer context windows don't guarantee better reasoning. Context rot can slowly degrade downstream answer accuracy even when the relevant information has already been retrieved.
Provenance-aware graph inference: Iteratively computes source reliability and claim support across a bipartite graph, accounting for shared evidence origins and dependencies between sources.
Context reconstruction: Selects original retrieved passages based on question relevance and incremental claim coverage within a token budget.
Text-blind evaluation: Evaluates structured observations without reading source text strings or invoking an LLM judge.
Local by default: Persists your agents' observations across sessions in a Context Graph , backed by SQLite or PostgreSQL.
pip install omneum
Run the context reconstruction demo:
git clone https://github.com/cole-h06/omneum.git
cd omneum
python -m examples.context_quality_demo.before_after
Five retrieved chunks, four distinct claims, 80-unit context budget.
Compare retrieval-order packing with Omneum's claim-coverage selection under the same context budget.
Packing chunks in retrieval order consumes 75 units while covering only two claims. Omneum retains two original chunks covering all four, using 32% less context capacity .
The demo measures its budget with the SDK's word/punctuation counter.
Graph Inference & Context Selection
Omneum uses a graph-based mathematical analysis, calculated from Observation objects, to estimate the structural reliability of sources and the support for claims found in retrieved context.
Importantly, this analysis does not naively assume every source has derived its information independently. The same information often propagates across multiple sources. Without dependency tracking, five observations from different sources can be difficult to distinguish from five reports derived from the same upstream source.
Source A: Acme's 2025 annual report
Claim: Acme reported $4.2 billion in revenue for 2025.
Source B: Financial article reporting the same claim
Claim: Acme reported $4.2 billion in revenue for 2025.
Upstream: Acme's 2025 annual report
How much support should that claim receive? That depends partly on the reliability of the sources reporting it. But how do we determine the reliability of the sources? By looking at the claims they support.
This definition seems circular, and it is. Nonetheless, it provides a well-defined notion of reliability that helps distinguish strongly supported context from weakly supported context.
Of course claim-level scores are only useful for context selection if they can be mapped back to chunks. To bridge this gap, Omneum stores a claim-to-chunk mapping: each chunk retains its original text and chunk_id , while its associated claim_ids reference the claims evaluated in the graph. A chunk can contain multiple claims, and the same claim can appear across multiple chunks. This allows the algorithm to associate claim support with the actual passages competing for space in the context window.
But selecting context is not simply a matter of ranking chunks by their individual scores. Omneum derives chunk utility from the claims associated with each passage, then ranks chunks to fill the model's window under a limited token budget. When a nonempty question is supplied, finalized claim contributions are weighted by deterministic, local lexical relevance over the original passage text. Repeated evidence from the same inferred source cluster earns only an improvement over its best selected relevance. Without a question, Omneum keeps its question-blind claim-coverage selection. Selected passages are reconstructed from their original text.
The selected passages are then reconstructed from their original text into the context passed to the LLM. The analysis itself never needs to read that text because it operates exclusively on structured relationships between sources and claims.
Omneum runs between retrieval and generation. The Python SDK accepts retrieved chunks and relevance scores from the calling retrieval pipeline, or computes relevance locally using BM25 or MiniLM when explicitly enabled.
Retrieval / tools
│
▼
┌───────────────────────────────────────────┐
│ ContextMapper / Observation input │
│ │
│ Source, claim and provenance mapping │
│ Structured Observation construction │
└─────────────────────┬─────────────────────┘
│ Structured observations
▼
┌───────────────────────────────────────────┐
│ Python SDK · prepare_context() │
│ │
│ Retrieved chunks and relevance scores │
│ Query conditioning / candidate discovery │
│ Original text retained locally │
└─────────────────────┬─────────────────────┘
│ Structured observations
▼
┌───────────────────────────────────────────┐
│ Omneum evaluator │
│ │
│ Provenance-aware graph inference │
│ Dependency-adjusted claim support │
│ No original source text required │
└─────────────────────┬─────────────────────┘
│ Graph evaluation
▼
┌───────────────────────────────────────────┐
│ Python SDK │
│ │
│ Claim-to-chunk score mapping │
│ Question-weighted claim support │
│ Marginal claim coverage │
│ Token-budget enforcement │
│ Original text reconstruction │
└─────────────────────┬─────────────────────┘
│ Reconstructed context
▼
Agent / LLM
client.prepare_context() uses Omneum's provenance-aware graph inference to evaluate source reliability and claim support, accounting for dependencies between sources. The SDK uses those results to rank and select retrieved passages within the token budget, then reconstructs their original text for the model.
The prepared context is returned in result["context"] , with selected passages in result["chunks"] and evaluation scores in result["evaluation"] .
Note: Omneum supports Python 3.11–3.14 on macOS and Ubuntu. The Python SDK is the currently supported client library.
Run the installer to configure the local evaluator:
omneum install
Run omneum install from your normal login account. Don't run the command with sudo .
The installer runs unprivileged while it downloads and verifies the evaluator release. The daemon version is pinned to the installed Python package version. For example, Omneum 1.0.3 only accepts daemon assets from v1.0.3 ; it doesn't resolve latest or fall back to another release.
The archive is verified against the release SHA256SUMS before installation. A missing checksum, mismatched digest, unexpected archive entry, link, or path traversal stops the install before anything from the archive runs.
Administrator access is requested only when installing machine-level service files. On Ubuntu, this includes the managed runtime under:
/opt/omneum
and public client configuration at:
/etc/omneum/config.toml
The evaluator runs under a dedicated service identity, isolating its VOPRF private key from the login account and MCP server process. macOS uses launchd ; Ubuntu uses systemd.
omneum-server exposes Omneum's MCP tools over stdio and communicates with the locally installed evaluator service. The installer finishes by sending a VOPRF request through that path and verifying the response against the configured public key. A running service alone does not count as a successful installation.
See docs/guides/installation.md for service accounts, sockets, installed paths, custom endpoints, and failure handling.
Add Omneum to your MCP configuration:
{
"mcpServers" : {
"omneum" : {
"command" : " omneum-server "
}
}
}
Restart the MCP client. It launches omneum-server over stdio and exposes Omneum's tools to your agent.
If the executable isn't found, locate it with command -v omneum-server and use the returned absolute path as command .
Ask your agent to call Omneum's ping tool. A successful connection returns:
{
"status" : " ok "
}
Context Preparation
Omneum evaluates structured observations, selects retrieved chunks under a token budget, and reconstructs the selected text for the downstream LLM. The evaluator computes dependency-adjusted claim support, while the Python SDK handles chunk selection and reconstruction.
In this example, the FAQ derives its audit retention claim from the security policy. The upstream dependency is recorded through upstream_sources .
from datetime import datetime , timezone
from omneum import ContextMapper , Source
from omneum . information import SourceReference
from omneum . context_selection import claim_id
now = datetime . now ( timezone . utc )
policy_url = "https://example.com/security-policy"
records = [
{
"id" : "obs-policy" ,
"url" : policy_url ,
"attribute" : "audit_retention_days" ,
"value" : 90 ,
"upstream" : [],
"text" : "Audit logs are retained for 90 days." ,
},
{
"id" : "obs-faq" ,
"url" : "https://example.com/faq" ,
"attribute" : "audit_retention_days" ,
"value" : 90 ,
"upstream" : [
SourceReference ( kind = "web_document" , identifier = policy_url )
],
"text" : "Audit logs are retained for 90 days." ,
},
{
"id" : "obs-incident" ,
"url" : "https://example.com/incident-report" ,
"attribute" : "incident_notification_hours" ,
"value" : 24 ,
"upstream" : [],
"text" : "Security incidents are reported within 24 hours." ,
},

[truncated]

## Original Extract

Text-blind context control plane for AI agents. Select high-utility context under token budgets. - cole-h06/omneum

GitHub - cole-h06/omneum: Text-blind context control plane for AI agents. Select high-utility context under token budgets. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
cole-h06
omneum Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Text-blind context control plane for AI agents. Select high-utility context under token budgets.
Readme MIT license Contributing
Security policy Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
312 Commits 312 Commits Folders and files
.cursor-plugin .cursor-plugin .github/ workflows .github/ workflows assets assets benchmarks benchmarks deploy deploy docs docs evals evals examples examples migrations migrations omneum omneum research research src src tests tests tools tools .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md build.rs build.rs compose.yaml compose.yaml glama.json glama.json mcp.json mcp.json pyproject.toml pyproject.toml requirements-dev.txt requirements-dev.txt requirements.txt requirements.txt rust-toolchain.toml rust-toolchain.toml server.json server.json uv.lock uv.lock View all files Repository files navigation
Provenance-aware context control plane for AI agents.
Omneum is an open-source context control plane for AI agents, built to optimize the evidence entering an LLM's context window.
Agents accumulate context from retrieval, tools, memory, and previous sessions, but longer context windows don't guarantee better reasoning. Context rot can slowly degrade downstream answer accuracy even when the relevant information has already been retrieved.
Provenance-aware graph inference: Iteratively computes source reliability and claim support across a bipartite graph, accounting for shared evidence origins and dependencies between sources.
Context reconstruction: Selects original retrieved passages based on question relevance and incremental claim coverage within a token budget.
Text-blind evaluation: Evaluates structured observations without reading source text strings or invoking an LLM judge.
Local by default: Persists your agents' observations across sessions in a Context Graph , backed by SQLite or PostgreSQL.
pip install omneum
Run the context reconstruction demo:
git clone https://github.com/cole-h06/omneum.git
cd omneum
python -m examples.context_quality_demo.before_after
Five retrieved chunks, four distinct claims, 80-unit context budget.
Compare retrieval-order packing with Omneum's claim-coverage selection under the same context budget.
Packing chunks in retrieval order consumes 75 units while covering only two claims. Omneum retains two original chunks covering all four, using 32% less context capacity .
The demo measures its budget with the SDK's word/punctuation counter.
Graph Inference & Context Selection
Omneum uses a graph-based mathematical analysis, calculated from Observation objects, to estimate the structural reliability of sources and the support for claims found in retrieved context.
Importantly, this analysis does not naively assume every source has derived its information independently. The same information often propagates across multiple sources. Without dependency tracking, five observations from different sources can be difficult to distinguish from five reports derived from the same upstream source.
Source A: Acme's 2025 annual report
Claim: Acme reported $4.2 billion in revenue for 2025.
Source B: Financial article reporting the same claim
Claim: Acme reported $4.2 billion in revenue for 2025.
Upstream: Acme's 2025 annual report
How much support should that claim receive? That depends partly on the reliability of the sources reporting it. But how do we determine the reliability of the sources? By looking at the claims they support.
This definition seems circular, and it is. Nonetheless, it provides a well-defined notion of reliability that helps distinguish strongly supported context from weakly supported context.
Of course claim-level scores are only useful for context selection if they can be mapped back to chunks. To bridge this gap, Omneum stores a claim-to-chunk mapping: each chunk retains its original text and chunk_id , while its associated claim_ids reference the claims evaluated in the graph. A chunk can contain multiple claims, and the same claim can appear across multiple chunks. This allows the algorithm to associate claim support with the actual passages competing for space in the context window.
But selecting context is not simply a matter of ranking chunks by their individual scores. Omneum derives chunk utility from the claims associated with each passage, then ranks chunks to fill the model's window under a limited token budget. When a nonempty question is supplied, finalized claim contributions are weighted by deterministic, local lexical relevance over the original passage text. Repeated evidence from the same inferred source cluster earns only an improvement over its best selected relevance. Without a question, Omneum keeps its question-blind claim-coverage selection. Selected passages are reconstructed from their original text.
The selected passages are then reconstructed from their original text into the context passed to the LLM. The analysis itself never needs to read that text because it operates exclusively on structured relationships between sources and claims.
Omneum runs between retrieval and generation. The Python SDK accepts retrieved chunks and relevance scores from the calling retrieval pipeline, or computes relevance locally using BM25 or MiniLM when explicitly enabled.
Retrieval / tools
│
▼
┌───────────────────────────────────────────┐
│ ContextMapper / Observation input │
│ │
│ Source, claim and provenance mapping │
│ Structured Observation construction │
└─────────────────────┬─────────────────────┘
│ Structured observations
▼
┌───────────────────────────────────────────┐
│ Python SDK · prepare_context() │
│ │
│ Retrieved chunks and relevance scores │
│ Query conditioning / candidate discovery │
│ Original text retained locally │
└─────────────────────┬─────────────────────┘
│ Structured observations
▼
┌───────────────────────────────────────────┐
│ Omneum evaluator │
│ │
│ Provenance-aware graph inference │
│ Dependency-adjusted claim support │
│ No original source text required │
└─────────────────────┬─────────────────────┘
│ Graph evaluation
▼
┌───────────────────────────────────────────┐
│ Python SDK │
│ │
│ Claim-to-chunk score mapping │
│ Question-weighted claim support │
│ Marginal claim coverage │
│ Token-budget enforcement │
│ Original text reconstruction │
└─────────────────────┬─────────────────────┘
│ Reconstructed context
▼
Agent / LLM
client.prepare_context() uses Omneum's provenance-aware graph inference to evaluate source reliability and claim support, accounting for dependencies between sources. The SDK uses those results to rank and select retrieved passages within the token budget, then reconstructs their original text for the model.
The prepared context is returned in result["context"] , with selected passages in result["chunks"] and evaluation scores in result["evaluation"] .
Note: Omneum supports Python 3.11–3.14 on macOS and Ubuntu. The Python SDK is the currently supported client library.
Run the installer to configure the local evaluator:
omneum install
Run omneum install from your normal login account. Don't run the command with sudo .
The installer runs unprivileged while it downloads and verifies the evaluator release. The daemon version is pinned to the installed Python package version. For example, Omneum 1.0.3 only accepts daemon assets from v1.0.3 ; it doesn't resolve latest or fall back to another release.
The archive is verified against the release SHA256SUMS before installation. A missing checksum, mismatched digest, unexpected archive entry, link, or path traversal stops the install before anything from the archive runs.
Administrator access is requested only when installing machine-level service files. On Ubuntu, this includes the managed runtime under:
/opt/omneum
and public client configuration at:
/etc/omneum/config.toml
The evaluator runs under a dedicated service identity, isolating its VOPRF private key from the login account and MCP server process. macOS uses launchd ; Ubuntu uses systemd.
omneum-server exposes Omneum's MCP tools over stdio and communicates with the locally installed evaluator service. The installer finishes by sending a VOPRF request through that path and verifying the response against the configured public key. A running service alone does not count as a successful installation.
See docs/guides/installation.md for service accounts, sockets, installed paths, custom endpoints, and failure handling.
Add Omneum to your MCP configuration:
{
"mcpServers" : {
"omneum" : {
"command" : " omneum-server "
}
}
}
Restart the MCP client. It launches omneum-server over stdio and exposes Omneum's tools to your agent.
If the executable isn't found, locate it with command -v omneum-server and use the returned absolute path as command .
Ask your agent to call Omneum's ping tool. A successful connection returns:
{
"status" : " ok "
}
Context Preparation
Omneum evaluates structured observations, selects retrieved chunks under a token budget, and reconstructs the selected text for the downstream LLM. The evaluator computes dependency-adjusted claim support, while the Python SDK handles chunk selection and reconstruction.
In this example, the FAQ derives its audit retention claim from the security policy. The upstream dependency is recorded through upstream_sources .
from datetime import datetime , timezone
from omneum import ContextMapper , Source
from omneum . information import SourceReference
from omneum . context_selection import claim_id
now = datetime . now ( timezone . utc )
policy_url = "https://example.com/security-policy"
records = [
{
"id" : "obs-policy" ,
"url" : policy_url ,
"attribute" : "audit_retention_days" ,
"value" : 90 ,
"upstream" : [],
"text" : "Audit logs are retained for 90 days." ,
},
{
"id" : "obs-faq" ,
"url" : "https://example.com/faq" ,
"attribute" : "audit_retention_days" ,
"value" : 90 ,
"upstream" : [
SourceReference ( kind = "web_document" , identifier = policy_url )
],
"text" : "Audit logs are retained for 90 days." ,
},
{
"id" : "obs-incident" ,
"url" : "https://example.com/incident-report" ,
"attribute" : "incident_notification_hours" ,
"value" : 24 ,
"upstream" : [],
"text" : "Security incidents are reported within 24 hours." ,
},

[truncated]
