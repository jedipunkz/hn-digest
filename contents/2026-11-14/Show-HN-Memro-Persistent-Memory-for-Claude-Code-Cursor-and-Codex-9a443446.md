---
source: "https://github.com/memrohq/memro"
hn_url: "https://news.ycombinator.com/item?id=50492093"
title: "Show HN: Memro – Persistent Memory for Claude Code, Cursor and Codex"
article_title: "GitHub - memrohq/memro: Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations. · GitHub"
image: "https://opengraph.githubassets.com/89a84a3076132455df38fc7dbd175c1606837f283a83a0e64a849e5b674f14c7/memrohq/memro"
author: "subalyo"
captured_at: "2026-11-14T00:42:08Z"
capture_tool: "hn-digest"
hn_id: 50492093
score: 2
comments: 1
posted_at: "2026-11-14T00:20:15Z"
tags:
  - hacker-news
---

# Show HN: Memro – Persistent Memory for Claude Code, Cursor and Codex

- HN: [50492093](https://news.ycombinator.com/item?id=50492093)
- Source: [github.com](https://github.com/memrohq/memro)
- Score: 2
- Comments: 1
- Posted: 2026-11-14T00:20:15Z

## Translation

Title: Show HN: Memro – Persistent Memory for Claude Code, Cursor and Codex
Article title: GitHub - memrohq/memro: Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations. · GitHub
Description: Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations. - memrohq/memro
HN text: Memro is go bases persistent memory system that lets ai agents share context acssion session using mcp, built on top of postgreSQL

Article text:
GitHub - memrohq/memro: Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations. · GitHub
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
memrohq
memro Public Notifications You must be signed in to change notification settings
Star 6 ( 6 ) You must be signed in to star a repository
Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations.
Security policy Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
201 Commits 201 Commits Folders and files
.github/ workflows .github/ workflows apps/ dashboard apps/ dashboard bench bench cli cli deploy deploy docs docs examples examples integrations integrations mcp mcp scripts scripts sdk sdk server server .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md docker-compose.release.yml docker-compose.release.yml docker-compose.yml docker-compose.yml go.work go.work go.work.sum go.work.sum smithery.yaml smithery.yaml start.sh start.sh View all files Repository files navigation
Memro — Verified Memory Infrastructure for AI Agents
Give your AI agents persistent memory across sessions: agent identity, memory storage and hybrid retrieval through an API, SDKs and MCP.
Memro is licensed under the PolyForm Noncommercial License 1.0.0 .
Free for Non-Commercial Use : Personal study, experiments, hobby projects, research, and testing without commercial application.
Commercial Use : Using Memro for commercial products or production services requires a managed subscription on Memro Cloud or an enterprise commercial license.
How to Get a Commercial License : Visit memro.co/pricing or contact licensing@memro.co .
Repository Notice : github.com/memrohq/memro is the single public repository for Memro. It contains the Go server, official SDKs (Python, JavaScript/TypeScript, Go), Model Context Protocol (MCP) server, and the CLI.
Option A: Memro Cloud with Claude Code, Cursor, or Codex (< 10 lines)
# 1. Connect installed agent tools via device login:
npx memro-cli connect
# 2. Check connection status:
npx memro-cli status
In your assistant, say: "Remember that our primary database is PostgreSQL 16 with pgvector."
Restart your session and ask: "What database and extensions do we use?"
Option B: Self-Host with Docker Compose (Source Available)
Run the full Memro API server and PostgreSQL with pgvector on your own infrastructure:
git clone https://github.com/memrohq/memro && cd memro
cp .env.example .env # add OPENAI_API_KEY for hybrid search; set METRICS_TOKEN
docker compose up -d # API on :8081, dashboard on :5173, Postgres on 127.0.0.1:55432
curl http://localhost:8081/health
Option 2: Prebuilt images (no clone required, after v0.1.0 release)
curl -fsSL https://raw.githubusercontent.com/memrohq/memro/v0.1.0/docker-compose.release.yml -o docker-compose.yml && docker compose up -d
curl http://localhost:8081/health
(Note: The prebuilt image one-liner downloads docker-compose.release.yml with pinned ghcr.io/memrohq/memro:0.1.0 and ghcr.io/memrohq/memro-dashboard:0.1.0 images, available once v0.1.0 is published.)
Backend API : http://localhost:8081
Dashboard : http://localhost:5173
Postgres Database : Port 55432 on localhost (Option 1 clone) to avoid port 5432 collisions with existing local Postgres setups.
Metrics Token : Set METRICS_TOKEN (min 32 characters) in .env to protect GET /metrics and GET /stats/cache ; the server logs a warning if unset.
Verify Signatures : Verify published container images keylessly using Cosign:
cosign verify ghcr.io/memrohq/memro:0.1.0 \
--certificate-identity-regexp ' https://github.com/memrohq/memro/ ' \
--certificate-oidc-issuer https://token.actions.githubusercontent.com
For complete deployment options, production topologies, and configuration tables, see docs/self-hosting.md .
┌──────────────────────────────────────┐
│ Clients & MCP Tools │
│ Claude Code · Cursor · Codex · SDKs│
└──────────────────┬───────────────────┘
│ HTTP REST / Stdio MCP
┌──────────────────▼───────────────────┐
│ Memro API (Go) │
│ Sovereign Ed25519 Identity · RRF │
└──────────────────┬───────────────────┘
│ SQL & Embeddings
┌──────────────────▼───────────────────┐
│ PostgreSQL 16 + pgvector │
│ pgvector Dense Search + tsvector FTS│
└──────────────────────────────────────┘
Memro runs a Go HTTP API backed by PostgreSQL and pgvector. Search queries combine vector cosine similarity and full-text keyword search using Reciprocal Rank Fusion (RRF, $k=60$ ). Agents authenticate using sovereign Ed25519 keypairs solved via local proof-of-work (private keys never leave the host) or standard Bearer API tokens.
memrohq/memro holds all components in one mono-repository:
Official SDKs & Runnable Examples
# PyPI (once published) or install from source: pip install ./sdk/python
pip install memro
from memro import Agent
agent = Agent . create () # local sovereign Ed25519 identity
mem = agent . remember ( "The billing service was rewritten in Go" , type = "semantic" )
results = agent . recall ( "billing service language" )
print ( results [ 0 ]. content )
See runnable script at examples/python/register_remember_recall.py .
JavaScript / TypeScript ( sdk/js )
# npm (once published) or install from local build: npm install ./sdk/js
npm install @memro/sdk
import { Agent } from '@memro/sdk' ;
const agent = await Agent . create ( ) ; // local sovereign Ed25519 identity
const mem = await agent . remember ( 'The billing service was rewritten in Go' , { type : 'semantic' } ) ;
const results = await agent . recall ( 'billing service language' ) ;
console . log ( results [ 0 ] . content ) ;
See runnable script at examples/typescript/index.js .
go get github.com/memrohq/memro/sdk/go@v0.1.0
package main
import (
"context"
"fmt"
memro "github.com/memrohq/memro/sdk/go"
)
func main () {
ctx := context . Background ()
agent , _ := memro . CreateAgent ( nil )
mem , _ := agent . Remember ( ctx , "The billing service was rewritten in Go" , nil )
results , _ := agent . Recall ( ctx , "billing service language" , nil )
fmt . Println ( results [ 0 ]. Content )
}
See runnable script at examples/go/main.go .
Step-by-step guides for connecting coding tools with full session-restart memory verification:
Claude Code Integration ( examples/claude-code/ )
Cursor Integration ( examples/cursor/ )
OpenAI Codex Integration ( examples/codex/ )
Memro includes structured primitives and cognitive graphs beyond basic vector search:
Typed Memories : Categorize observations as semantic (facts), episodic (events/conversations), procedural (runbooks/recipes), or profile (agent guardrails). See Memory Concepts .
Temporal Search & Evolution : Query memories with timestamp weighting and trace how beliefs evolve over time. See API Reference .
Conflict Analysis & Resolution : Identify contradictory facts across sessions and mark superseded memories using resolve_conflict() . See Conflict Resolution .
Reasoning Graphs : Trace memory provenance, causal chains, and dependencies between stored observations. See Opinionated Guide .
Measurable Verification & Benchmarks
Memro publishes empirical evaluations and methodology:
LongMemEval Benchmark : Evaluated across 500 questions testing multi-session recall and knowledge updates. See LongMemEval Report .
Empirical Failure & Boundary Evaluations : Runnable test suite for identity boundaries, contradictory facts, stale facts, and distractor rejection. See Measurable Evaluations .
Clean-Machine Verification : Full transcript verifying every quickstart command in a container without local caches. See Verification Transcript .
# 1. Start PostgreSQL with pgvector
docker compose up -d db
# 2. Run backend API server
cd server
DATABASE_URL=postgres://memro:memro@localhost:55432/memro go run ./cmd/memro-api
# 3. Run tests
MEMRO_TEST_DATABASE_URL=postgres://memro:memro@localhost:55432/memro go test -race ./...
Please review CONTRIBUTING.md and SECURITY.md before submitting pull requests.
Repository: github.com/memrohq/memro · Documentation: https://memro.co
License: Source available under PolyForm Noncommercial 1.0.0 . Free for non-commercial use; commercial use requires Memro Cloud or a commercial license.
Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations.
Security policy Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations. - memrohq/memro

Memro is go bases persistent memory system that lets ai agents share context acssion session using mcp, built on top of postgreSQL

GitHub - memrohq/memro: Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations. · GitHub
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
memrohq
memro Public Notifications You must be signed in to change notification settings
Star 6 ( 6 ) You must be signed in to star a repository
Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations.
Security policy Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
201 Commits 201 Commits Folders and files
.github/ workflows .github/ workflows apps/ dashboard apps/ dashboard bench bench cli cli deploy deploy docs docs examples examples integrations integrations mcp mcp scripts scripts sdk sdk server server .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE QUICKSTART.md QUICKSTART.md README.md README.md SECURITY.md SECURITY.md docker-compose.release.yml docker-compose.release.yml docker-compose.yml docker-compose.yml go.work go.work go.work.sum go.work.sum smithery.yaml smithery.yaml start.sh start.sh View all files Repository files navigation
Memro — Verified Memory Infrastructure for AI Agents
Give your AI agents persistent memory across sessions: agent identity, memory storage and hybrid retrieval through an API, SDKs and MCP.
Memro is licensed under the PolyForm Noncommercial License 1.0.0 .
Free for Non-Commercial Use : Personal study, experiments, hobby projects, research, and testing without commercial application.
Commercial Use : Using Memro for commercial products or production services requires a managed subscription on Memro Cloud or an enterprise commercial license.
How to Get a Commercial License : Visit memro.co/pricing or contact licensing@memro.co .
Repository Notice : github.com/memrohq/memro is the single public repository for Memro. It contains the Go server, official SDKs (Python, JavaScript/TypeScript, Go), Model Context Protocol (MCP) server, and the CLI.
Option A: Memro Cloud with Claude Code, Cursor, or Codex (< 10 lines)
# 1. Connect installed agent tools via device login:
npx memro-cli connect
# 2. Check connection status:
npx memro-cli status
In your assistant, say: "Remember that our primary database is PostgreSQL 16 with pgvector."
Restart your session and ask: "What database and extensions do we use?"
Option B: Self-Host with Docker Compose (Source Available)
Run the full Memro API server and PostgreSQL with pgvector on your own infrastructure:
git clone https://github.com/memrohq/memro && cd memro
cp .env.example .env # add OPENAI_API_KEY for hybrid search; set METRICS_TOKEN
docker compose up -d # API on :8081, dashboard on :5173, Postgres on 127.0.0.1:55432
curl http://localhost:8081/health
Option 2: Prebuilt images (no clone required, after v0.1.0 release)
curl -fsSL https://raw.githubusercontent.com/memrohq/memro/v0.1.0/docker-compose.release.yml -o docker-compose.yml && docker compose up -d
curl http://localhost:8081/health
(Note: The prebuilt image one-liner downloads docker-compose.release.yml with pinned ghcr.io/memrohq/memro:0.1.0 and ghcr.io/memrohq/memro-dashboard:0.1.0 images, available once v0.1.0 is published.)
Backend API : http://localhost:8081
Dashboard : http://localhost:5173
Postgres Database : Port 55432 on localhost (Option 1 clone) to avoid port 5432 collisions with existing local Postgres setups.
Metrics Token : Set METRICS_TOKEN (min 32 characters) in .env to protect GET /metrics and GET /stats/cache ; the server logs a warning if unset.
Verify Signatures : Verify published container images keylessly using Cosign:
cosign verify ghcr.io/memrohq/memro:0.1.0 \
--certificate-identity-regexp ' https://github.com/memrohq/memro/ ' \
--certificate-oidc-issuer https://token.actions.githubusercontent.com
For complete deployment options, production topologies, and configuration tables, see docs/self-hosting.md .
┌──────────────────────────────────────┐
│ Clients & MCP Tools │
│ Claude Code · Cursor · Codex · SDKs│
└──────────────────┬───────────────────┘
│ HTTP REST / Stdio MCP
┌──────────────────▼───────────────────┐
│ Memro API (Go) │
│ Sovereign Ed25519 Identity · RRF │
└──────────────────┬───────────────────┘
│ SQL & Embeddings
┌──────────────────▼───────────────────┐
│ PostgreSQL 16 + pgvector │
│ pgvector Dense Search + tsvector FTS│
└──────────────────────────────────────┘
Memro runs a Go HTTP API backed by PostgreSQL and pgvector. Search queries combine vector cosine similarity and full-text keyword search using Reciprocal Rank Fusion (RRF, $k=60$ ). Agents authenticate using sovereign Ed25519 keypairs solved via local proof-of-work (private keys never leave the host) or standard Bearer API tokens.
memrohq/memro holds all components in one mono-repository:
Official SDKs & Runnable Examples
# PyPI (once published) or install from source: pip install ./sdk/python
pip install memro
from memro import Agent
agent = Agent . create () # local sovereign Ed25519 identity
mem = agent . remember ( "The billing service was rewritten in Go" , type = "semantic" )
results = agent . recall ( "billing service language" )
print ( results [ 0 ]. content )
See runnable script at examples/python/register_remember_recall.py .
JavaScript / TypeScript ( sdk/js )
# npm (once published) or install from local build: npm install ./sdk/js
npm install @memro/sdk
import { Agent } from '@memro/sdk' ;
const agent = await Agent . create ( ) ; // local sovereign Ed25519 identity
const mem = await agent . remember ( 'The billing service was rewritten in Go' , { type : 'semantic' } ) ;
const results = await agent . recall ( 'billing service language' ) ;
console . log ( results [ 0 ] . content ) ;
See runnable script at examples/typescript/index.js .
go get github.com/memrohq/memro/sdk/go@v0.1.0
package main
import (
"context"
"fmt"
memro "github.com/memrohq/memro/sdk/go"
)
func main () {
ctx := context . Background ()
agent , _ := memro . CreateAgent ( nil )
mem , _ := agent . Remember ( ctx , "The billing service was rewritten in Go" , nil )
results , _ := agent . Recall ( ctx , "billing service language" , nil )
fmt . Println ( results [ 0 ]. Content )
}
See runnable script at examples/go/main.go .
Step-by-step guides for connecting coding tools with full session-restart memory verification:
Claude Code Integration ( examples/claude-code/ )
Cursor Integration ( examples/cursor/ )
OpenAI Codex Integration ( examples/codex/ )
Memro includes structured primitives and cognitive graphs beyond basic vector search:
Typed Memories : Categorize observations as semantic (facts), episodic (events/conversations), procedural (runbooks/recipes), or profile (agent guardrails). See Memory Concepts .
Temporal Search & Evolution : Query memories with timestamp weighting and trace how beliefs evolve over time. See API Reference .
Conflict Analysis & Resolution : Identify contradictory facts across sessions and mark superseded memories using resolve_conflict() . See Conflict Resolution .
Reasoning Graphs : Trace memory provenance, causal chains, and dependencies between stored observations. See Opinionated Guide .
Measurable Verification & Benchmarks
Memro publishes empirical evaluations and methodology:
LongMemEval Benchmark : Evaluated across 500 questions testing multi-session recall and knowledge updates. See LongMemEval Report .
Empirical Failure & Boundary Evaluations : Runnable test suite for identity boundaries, contradictory facts, stale facts, and distractor rejection. See Measurable Evaluations .
Clean-Machine Verification : Full transcript verifying every quickstart command in a container without local caches. See Verification Transcript .
# 1. Start PostgreSQL with pgvector
docker compose up -d db
# 2. Run backend API server
cd server
DATABASE_URL=postgres://memro:memro@localhost:55432/memro go run ./cmd/memro-api
# 3. Run tests
MEMRO_TEST_DATABASE_URL=postgres://memro:memro@localhost:55432/memro go test -race ./...
Please review CONTRIBUTING.md and SECURITY.md before submitting pull requests.
Repository: github.com/memrohq/memro · Documentation: https://memro.co
License: Source available under PolyForm Noncommercial 1.0.0 . Free for non-commercial use; commercial use requires Memro Cloud or a commercial license.
Memro — the memory layer for AI context. Server, SDKs, MCP server and integrations.
Security policy Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
