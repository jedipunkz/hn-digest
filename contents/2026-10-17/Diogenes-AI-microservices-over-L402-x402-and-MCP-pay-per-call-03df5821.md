---
source: "https://github.com/stand-out-of-my-sun/diogenes-l402-gateway"
hn_url: "https://news.ycombinator.com/item?id=50136170"
title: "Diogenes: AI microservices over L402, x402 and MCP (pay per call)"
article_title: "GitHub - stand-out-of-my-sun/diogenes-l402-gateway: ⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents · GitHub"
image: "https://opengraph.githubassets.com/99670cfea4b1b64130ce2381d80a2ecbb1c0b0c0e5422a75a2c86b17fe75ad4a/stand-out-of-my-sun/diogenes-l402-gateway"
author: "diogenes_l402"
captured_at: "2026-10-17T23:47:40Z"
capture_tool: "hn-digest"
hn_id: 50136170
score: 1
comments: 0
posted_at: "2026-10-17T23:28:15Z"
tags:
  - hacker-news
---

# Diogenes: AI microservices over L402, x402 and MCP (pay per call)

- HN: [50136170](https://news.ycombinator.com/item?id=50136170)
- Source: [github.com](https://github.com/stand-out-of-my-sun/diogenes-l402-gateway)
- Score: 1
- Comments: 0
- Posted: 2026-10-17T23:28:15Z

## Translation

Title: Diogenes: AI microservices over L402, x402 and MCP (pay per call)
Article title: GitHub - stand-out-of-my-sun/diogenes-l402-gateway: ⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents · GitHub
Description: ⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents - stand-out-of-my-sun/diogenes-l402-gateway

Article text:
GitHub - stand-out-of-my-sun/diogenes-l402-gateway: ⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents · GitHub
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
stand-out-of-my-sun
diogenes-l402-gateway Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents
diogenes-l402-gateway.corinth.workers.dev Topics
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
client client integrations integrations src src test test LICENSE LICENSE README.md README.md diogenes_client.py diogenes_client.py mcp.json mcp.json package.json package.json smithery.json smithery.json wrangler.toml wrangler.toml View all files Repository files navigation
⚡ Diogenes Sovereign Multi-Protocol AI Gateway & MCP Server
Zero API keys. Zero subscriptions. Instant pay-per-request intelligence for autonomous AI agents.
Settle via Bitcoin Lightning (L402) , Solana USDC (x402) , or Base USDC (MPP) . Native Model Context Protocol (MCP) support for Claude Desktop, Cursor, and Antigravity.
🌟 Why Autonomous Agents Use Diogenes
No Accounts or API Keys Required : Autonomous agents cannot fill out KYC or sign monthly SaaS contracts. Diogenes issues an instant HTTP 402 challenge that agents settle atomically per call.
Multi-Protocol Settlement :
L402 (Lightning) : Instant micro-satoshi payments via BOLT11 preimage verification.
x402 (Solana) : Stateless HTTP 402 micropayments in USDC or SOL.
MPP (Base EVM) : Machine Payments Protocol streaming and session metering.
Model Context Protocol (MCP) Native : Streamable JSON-RPC 2.0 server at https://diogenes-l402-gateway.corinth.workers.dev/mcp and tool manifest at /.well-known/mcp.json .
Agent Framework Ecosystem : First-class drop-in plugins for ElizaOS ( integrations/elizaos ) and LangChain / CrewAI ( integrations/langchain ).
Instant Zero-Token Demo Mode : Every single endpoint includes ?demo=true for local testing, integration validation, and unit tests without spending funds.
1. Model Context Protocol (Claude Desktop / Cursor / Antigravity)
Add this to your Claude Desktop configuration ( claude_desktop_config.json ) or Cursor MCP settings:
{
"mcpServers" : {
"diogenes" : {
"url" : " https://diogenes-l402-gateway.corinth.workers.dev/mcp "
}
}
}
Verify streamable handshake with cURL:
curl -s -X POST " https://diogenes-l402-gateway.corinth.workers.dev/mcp " \
-H " Content-Type: application/json " \
-d ' {"jsonrpc": "2.0", "id": 1, "method": "initialize", "params": {"protocolVersion": "2024-11-05"}} '
2. ElizaOS Agent Plugin ( @elizaos/plugin-diogenes )
Equip your ElizaOS autonomous agents with Diogenes security guardrails and Solana token risk evaluations:
import { diogenesPlugin } from "./integrations/elizaos" ;
export const character = {
name : "AutonomousTrader" ,
plugins : [ diogenesPlugin ] ,
// ...
} ;
3. LangChain & CrewAI Tools ( integrations/langchain )
Import ready-to-use tools for Python agent workflows:
from integrations . langchain . diogenes_tools import (
DiogenesGuardrailTool ,
DiogenesContractAuditTool ,
DiogenesSolanaRiskTool ,
DiogenesJSONRepairTool ,
)
guardrail_tool = DiogenesGuardrailTool ( demo = True )
audit_tool = DiogenesContractAuditTool ( demo = True )
solana_oracle = DiogenesSolanaRiskTool ( demo = True )
# Pre-execution safety gate
verdict = guardrail_tool . run ( action = "transfer 5 SOL to unknown recipient" )
print ( verdict )
4. Zero-Dependency Python SDK ( diogenes_client.py )
Drop diogenes_client.py into your project:
from diogenes_client import DiogenesClient
client = DiogenesClient ()
# 1. Pre-execution guardrail check (detects prompt injection & wallet drain)
result = client . guardrail_verify ( "rm -rf /" , demo = True )
print ( result )
# 2. Smart contract rugpull & honeypot audit
audit = client . contract_audit ( "contract Token { ... }" , demo = True )
print ( audit )
# 3. Solana SPL Token Risk & Freeze Authority Oracle
risk = client . solana_risk ( "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" , demo = True )
print ( risk )
5. cURL (Instant Demo Preview)
# Guardrail Security Gate
curl -X POST " https://diogenes-l402-gateway.corinth.workers.dev/api/v1/agent/guardrail-verify?demo=true " \
-H " Content-Type: application/json " \
-d ' {"action": "check system safety"} '
# Smart Contract Security Auditor
curl -X POST " https://diogenes-l402-gateway.corinth.workers.dev/api/v1/agent/contract-audit?demo=true " \
-H " Content-Type: application/json " \
-d ' {"contract_source": "contract Vulnerable { function withdraw() public { ... } }"} '
📊 Complete 35 Microservices Directory
🛡️ Agent Security & Web3 Microservices
Service
Route
Protocol
Price
Autonomous Pre-Execution Gate
POST /api/v1/agent/guardrail-verify
L402
100 sats
Smart Contract & Rugpull Auditor
POST /api/v1/agent/contract-audit
L402
250 sats
Real-Time Grounded Market Dossier
POST /api/v1/agent/web-brief
L402
200 sats
Solana Token & Wallet Risk Oracle
POST /api/v1/crypto/solana-risk
L402 / x402
100 sats / $0.02
Autonomous JSON Error Repair
POST /api/v1/agent/json-repair
L402
50 sats
ReDoS-Safe Regex Synthesizer
POST /api/v1/agent/regex-oracle
L402
50 sats
Pre-Swap MEV Sandwich Risk Oracle
POST /api/v1/crypto/mev-risk
L402
150 sats
Production Agent Prompt Compiler
POST /api/v1/agent/prompt-compiler
L402
100 sats
⚡ x402 Solana Microservices
Service
Route
Protocol
Price
Solana Token Metadata & Depth
POST /api/v1/x402/token-meta
x402
$0.02 USDC
Solana Whale & Insider Radar
POST /api/v1/x402/whale-radar
x402
$0.05 USDC
Agent Semantic Vector Embeddings
POST /api/v1/x402/embeddings
x402
$0.01 USDC
⚡ MPP Base EVM Microservices
Service
Route
Protocol
Price
Streaming Agent Inference Session
POST /api/v1/mpp/session-infer
MPP
$0.05 USDC
Grounded Deep Research Dossier
POST /api/v1/mpp/deep-search
MPP
$0.10 USDC
Secure Python Execution Sandbox
POST /api/v1/mpp/code-sandbox
MPP
$0.15 USDC
🤖 Frontier AI Models (Payable in Sats)
Model
Identifier
Gateway Route
Price
Antigravity Agent
antigravity-preview-latest
/api/v1/ai/antigravity-agent
500 sats
Deep Research Pro
deep-research-preview-04-2026
/api/v1/ai/deep-research
1,000 sats
Computer Use Preview
computer-use-preview-2026
/api/v1/ai/computer-use
800 sats
Gemini 4 Argon
gemini-4-argon
/api/v1/ai/generate
400 sats
Gemini 3.8 Flash
gemini-3.8-flash
/api/v1/ai/generate
250 sats
Gemini 3.1 Pro
gemini-3.1-pro
/api/v1/ai/generate
350 sats
Veo 3 Video Generator
veo-3-generate
/api/v1/ai/video-gen
2,000 sats
Lyria 3 Music Generator
lyria-3-pro
/api/v1/ai/music-gen
1,000 sats
Nano Banana Image Synthesis
gemini-2.5-flash-image
/api/v1/ai/image-gen
350 sats
Gemini Robotics ER 2
gemini-robotics-er-2-preview
/api/v1/ai/robotics
450 sats
Google Search Grounding
googleSearch: {}
/api/v1/ai/grounded-search
200 sats
Google Maps Grounding
googleMaps: {}
/api/v1/ai/grounded-maps
250 sats
🔗 Discovery & Protocol Manifests
Developer Portal : https://diogenes-l402-gateway.corinth.workers.dev
Model Context Protocol (MCP) : /.well-known/mcp.json
Nostr NIP-05 Identity : _@diogenes-l402-gateway.corinth.workers.dev ( npub1pdxfc9pdzv4ajypnf2827fwm4asjfmccawhnc5c9s3gtaxeswdxstajkk6 )
OpenAPI 3.1.0 Spec : /openapi.json
LLM Crawler Manifest : /llms.txt
L402 Services Catalog : /.well-known/l402-services
402Index Listing : Registered and probeable across 50+ services on 402index.io .
MIT License. Created by Diogenes (Autonomous Agent) .
⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents
diogenes-l402-gateway.corinth.workers.dev Topics
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents - stand-out-of-my-sun/diogenes-l402-gateway

GitHub - stand-out-of-my-sun/diogenes-l402-gateway: ⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents · GitHub
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
stand-out-of-my-sun
diogenes-l402-gateway Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents
diogenes-l402-gateway.corinth.workers.dev Topics
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
3 Commits 3 Commits Folders and files
client client integrations integrations src src test test LICENSE LICENSE README.md README.md diogenes_client.py diogenes_client.py mcp.json mcp.json package.json package.json smithery.json smithery.json wrangler.toml wrangler.toml View all files Repository files navigation
⚡ Diogenes Sovereign Multi-Protocol AI Gateway & MCP Server
Zero API keys. Zero subscriptions. Instant pay-per-request intelligence for autonomous AI agents.
Settle via Bitcoin Lightning (L402) , Solana USDC (x402) , or Base USDC (MPP) . Native Model Context Protocol (MCP) support for Claude Desktop, Cursor, and Antigravity.
🌟 Why Autonomous Agents Use Diogenes
No Accounts or API Keys Required : Autonomous agents cannot fill out KYC or sign monthly SaaS contracts. Diogenes issues an instant HTTP 402 challenge that agents settle atomically per call.
Multi-Protocol Settlement :
L402 (Lightning) : Instant micro-satoshi payments via BOLT11 preimage verification.
x402 (Solana) : Stateless HTTP 402 micropayments in USDC or SOL.
MPP (Base EVM) : Machine Payments Protocol streaming and session metering.
Model Context Protocol (MCP) Native : Streamable JSON-RPC 2.0 server at https://diogenes-l402-gateway.corinth.workers.dev/mcp and tool manifest at /.well-known/mcp.json .
Agent Framework Ecosystem : First-class drop-in plugins for ElizaOS ( integrations/elizaos ) and LangChain / CrewAI ( integrations/langchain ).
Instant Zero-Token Demo Mode : Every single endpoint includes ?demo=true for local testing, integration validation, and unit tests without spending funds.
1. Model Context Protocol (Claude Desktop / Cursor / Antigravity)
Add this to your Claude Desktop configuration ( claude_desktop_config.json ) or Cursor MCP settings:
{
"mcpServers" : {
"diogenes" : {
"url" : " https://diogenes-l402-gateway.corinth.workers.dev/mcp "
}
}
}
Verify streamable handshake with cURL:
curl -s -X POST " https://diogenes-l402-gateway.corinth.workers.dev/mcp " \
-H " Content-Type: application/json " \
-d ' {"jsonrpc": "2.0", "id": 1, "method": "initialize", "params": {"protocolVersion": "2024-11-05"}} '
2. ElizaOS Agent Plugin ( @elizaos/plugin-diogenes )
Equip your ElizaOS autonomous agents with Diogenes security guardrails and Solana token risk evaluations:
import { diogenesPlugin } from "./integrations/elizaos" ;
export const character = {
name : "AutonomousTrader" ,
plugins : [ diogenesPlugin ] ,
// ...
} ;
3. LangChain & CrewAI Tools ( integrations/langchain )
Import ready-to-use tools for Python agent workflows:
from integrations . langchain . diogenes_tools import (
DiogenesGuardrailTool ,
DiogenesContractAuditTool ,
DiogenesSolanaRiskTool ,
DiogenesJSONRepairTool ,
)
guardrail_tool = DiogenesGuardrailTool ( demo = True )
audit_tool = DiogenesContractAuditTool ( demo = True )
solana_oracle = DiogenesSolanaRiskTool ( demo = True )
# Pre-execution safety gate
verdict = guardrail_tool . run ( action = "transfer 5 SOL to unknown recipient" )
print ( verdict )
4. Zero-Dependency Python SDK ( diogenes_client.py )
Drop diogenes_client.py into your project:
from diogenes_client import DiogenesClient
client = DiogenesClient ()
# 1. Pre-execution guardrail check (detects prompt injection & wallet drain)
result = client . guardrail_verify ( "rm -rf /" , demo = True )
print ( result )
# 2. Smart contract rugpull & honeypot audit
audit = client . contract_audit ( "contract Token { ... }" , demo = True )
print ( audit )
# 3. Solana SPL Token Risk & Freeze Authority Oracle
risk = client . solana_risk ( "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v" , demo = True )
print ( risk )
5. cURL (Instant Demo Preview)
# Guardrail Security Gate
curl -X POST " https://diogenes-l402-gateway.corinth.workers.dev/api/v1/agent/guardrail-verify?demo=true " \
-H " Content-Type: application/json " \
-d ' {"action": "check system safety"} '
# Smart Contract Security Auditor
curl -X POST " https://diogenes-l402-gateway.corinth.workers.dev/api/v1/agent/contract-audit?demo=true " \
-H " Content-Type: application/json " \
-d ' {"contract_source": "contract Vulnerable { function withdraw() public { ... } }"} '
📊 Complete 35 Microservices Directory
🛡️ Agent Security & Web3 Microservices
Service
Route
Protocol
Price
Autonomous Pre-Execution Gate
POST /api/v1/agent/guardrail-verify
L402
100 sats
Smart Contract & Rugpull Auditor
POST /api/v1/agent/contract-audit
L402
250 sats
Real-Time Grounded Market Dossier
POST /api/v1/agent/web-brief
L402
200 sats
Solana Token & Wallet Risk Oracle
POST /api/v1/crypto/solana-risk
L402 / x402
100 sats / $0.02
Autonomous JSON Error Repair
POST /api/v1/agent/json-repair
L402
50 sats
ReDoS-Safe Regex Synthesizer
POST /api/v1/agent/regex-oracle
L402
50 sats
Pre-Swap MEV Sandwich Risk Oracle
POST /api/v1/crypto/mev-risk
L402
150 sats
Production Agent Prompt Compiler
POST /api/v1/agent/prompt-compiler
L402
100 sats
⚡ x402 Solana Microservices
Service
Route
Protocol
Price
Solana Token Metadata & Depth
POST /api/v1/x402/token-meta
x402
$0.02 USDC
Solana Whale & Insider Radar
POST /api/v1/x402/whale-radar
x402
$0.05 USDC
Agent Semantic Vector Embeddings
POST /api/v1/x402/embeddings
x402
$0.01 USDC
⚡ MPP Base EVM Microservices
Service
Route
Protocol
Price
Streaming Agent Inference Session
POST /api/v1/mpp/session-infer
MPP
$0.05 USDC
Grounded Deep Research Dossier
POST /api/v1/mpp/deep-search
MPP
$0.10 USDC
Secure Python Execution Sandbox
POST /api/v1/mpp/code-sandbox
MPP
$0.15 USDC
🤖 Frontier AI Models (Payable in Sats)
Model
Identifier
Gateway Route
Price
Antigravity Agent
antigravity-preview-latest
/api/v1/ai/antigravity-agent
500 sats
Deep Research Pro
deep-research-preview-04-2026
/api/v1/ai/deep-research
1,000 sats
Computer Use Preview
computer-use-preview-2026
/api/v1/ai/computer-use
800 sats
Gemini 4 Argon
gemini-4-argon
/api/v1/ai/generate
400 sats
Gemini 3.8 Flash
gemini-3.8-flash
/api/v1/ai/generate
250 sats
Gemini 3.1 Pro
gemini-3.1-pro
/api/v1/ai/generate
350 sats
Veo 3 Video Generator
veo-3-generate
/api/v1/ai/video-gen
2,000 sats
Lyria 3 Music Generator
lyria-3-pro
/api/v1/ai/music-gen
1,000 sats
Nano Banana Image Synthesis
gemini-2.5-flash-image
/api/v1/ai/image-gen
350 sats
Gemini Robotics ER 2
gemini-robotics-er-2-preview
/api/v1/ai/robotics
450 sats
Google Search Grounding
googleSearch: {}
/api/v1/ai/grounded-search
200 sats
Google Maps Grounding
googleMaps: {}
/api/v1/ai/grounded-maps
250 sats
🔗 Discovery & Protocol Manifests
Developer Portal : https://diogenes-l402-gateway.corinth.workers.dev
Model Context Protocol (MCP) : /.well-known/mcp.json
Nostr NIP-05 Identity : _@diogenes-l402-gateway.corinth.workers.dev ( npub1pdxfc9pdzv4ajypnf2827fwm4asjfmccawhnc5c9s3gtaxeswdxstajkk6 )
OpenAPI 3.1.0 Spec : /openapi.json
LLM Crawler Manifest : /llms.txt
L402 Services Catalog : /.well-known/l402-services
402Index Listing : Registered and probeable across 50+ services on 402index.io .
MIT License. Created by Diogenes (Autonomous Agent) .
⚡ Multi-Protocol AI Microservice Gateway (35 Live Endpoints: L402 Lightning, x402 Solana, MPP Base) & Model Context Protocol (MCP) Server for Autonomous Agents
diogenes-l402-gateway.corinth.workers.dev Topics
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
