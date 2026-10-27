---
source: "https://github.com/ArohaLabs/aroha-protocol"
hn_url: "https://news.ycombinator.com/item?id=50253316"
title: "Aroha – signed, shrinking permissions for AI agents, checked offline"
article_title: "GitHub - ArohaLabs/aroha-protocol: Aroha: verifiable identity and spending limits for AI agents · GitHub"
image: "https://opengraph.githubassets.com/a01763a3f34254ca7f440c3ce4848dffa9962869c9b4aba4b3d13dc409d44bba/ArohaLabs/aroha-protocol"
author: "rajganesan-aroh"
captured_at: "2026-10-27T01:12:54Z"
capture_tool: "hn-digest"
hn_id: 50253316
score: 1
comments: 0
posted_at: "2026-10-27T00:43:16Z"
tags:
  - hacker-news
---

# Aroha – signed, shrinking permissions for AI agents, checked offline

- HN: [50253316](https://news.ycombinator.com/item?id=50253316)
- Source: [github.com](https://github.com/ArohaLabs/aroha-protocol)
- Score: 1
- Comments: 0
- Posted: 2026-10-27T00:43:16Z

## Translation

Title: Aroha – signed, shrinking permissions for AI agents, checked offline
Article title: GitHub - ArohaLabs/aroha-protocol: Aroha: verifiable identity and spending limits for AI agents · GitHub
Description: Aroha: verifiable identity and spending limits for AI agents - ArohaLabs/aroha-protocol

Article text:
GitHub - ArohaLabs/aroha-protocol: Aroha: verifiable identity and spending limits for AI agents · GitHub
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
ArohaLabs
aroha-protocol Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Aroha: verifiable identity and spending limits for AI agents
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
21 Commits 21 Commits Folders and files
.github .github .well-known .well-known LICENSES/ third-party LICENSES/ third-party agents/ mock-provider-agent agents/ mock-provider-agent apps apps community community contracts contracts docs docs examples examples packages packages papers papers script script scripts scripts sdks sdks skills skills spec-public spec-public test test .gitignore .gitignore .lycheeignore .lycheeignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CHARTER.md CHARTER.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMPATIBILITY.md COMPATIBILITY.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md TODO.md TODO.md chain-addresses.json chain-addresses.json foundry.toml foundry.toml lychee.toml lychee.toml mkdocs.yml mkdocs.yml package-lock.json package-lock.json package.json package.json quickstart.sh quickstart.sh requirements-docs.txt requirements-docs.txt tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.conformance.config.ts vitest.conformance.config.ts View all files Repository files navigation
Cryptographically bounded authority for every agent in the chain.
Aroha Mandates give orchestrators cryptographically enforceable control over every sub-agent they delegate to. No agent can exceed the scope it was granted.
Docs · Support matrix · Local sandbox · Support · Spec
import { generateKeyPair , generateDid } from "@aroha-sdk/core" ;
import { issueIntentMandate , attenuateToPayment , verifyMandateChain } from "@aroha-sdk/credentials" ;
const human = await generateKeyPair ( ) ;
const orchestrator = await generateKeyPair ( ) ;
const flightAgent = await generateKeyPair ( ) ;
// Content-addressed DIDs (did:aroha:z…) commit to their own public key,
// so a verifier needs no registry lookup to check signatures.
const humanDID = generateDid ( human . publicKey ) ;
const orchestratorDID = generateDid ( orchestrator . publicKey ) ;
const flightAgentDID = generateDid ( flightAgent . publicKey ) ;
// Human issues an intent mandate to the orchestrator ($500, 1 hour TTL)
const intent = await issueIntentMandate (
humanDID , orchestratorDID ,
{ spendLimitUsd : 500 , merchantCategory : "4511" } , // MCC 4511 = airlines
human . privateKey ,
3_600_000 ,
) ;
// Orchestrator attenuates scope downward to the flight agent ($300 of the $500).
// The child token embeds the signed chain back to the human.
const payment = await attenuateToPayment (
intent , flightAgentDID ,
{ spendLimitUsd : 300 , merchantCategory : "4511" } ,
orchestrator . privateKey ,
) ;
// Flight agent verifies every link back to the human root before acting
const result = await verifyMandateChain ( payment . token , undefined , {
expectedRoot : humanDID , // the human authority must come from
expectedGrantee : flightAgentDID , // the flight agent itself
} ) ;
// result.valid === true, result.chainVerified === true,
// result.mandate?.constraints.spendLimitUsd === 300.
// A link that widens scope, or is signed by anyone but its grantor, fails.
The problem
Agents now act for people and hand work to other agents and tools, often run by other companies. MCP and A2A connect them; identity providers name them. Nothing that travels with the work proves what each agent was allowed to do, stops the whole chain when the person changes their mind, or leaves a record anyone else can check.
Aroha answers four questions for every action: who is acting, what were they allowed to do, can it be stopped, and what actually happened. It works with MCP, A2A and your identity provider rather than replacing them. See what Aroha is (and is not) .
Human issues a mandate — signs an IntentMandate with authorized actions, a spending cap, and a TTL. Cryptographically bound to the orchestrator's DID.
Orchestrator attenuates — delegates a strict subset of that scope to the next agent. The new mandate is signed by the orchestrator. Scope narrows, never widens.
Every agent verifies — before acting, an agent verifies the mandate chain back to its human origin, offline. A forged, altered, widened or expired link fails verification.
Revoke and prove — the person can revoke at any time; the registry witness pushes the revocation to verifiers. Every allowed or refused action produces a signed receipt.
Deploy in your existing stack — bridges for MCP, LangChain, A2A, and Composio. Mandate verification drops in in two lines.
Core — the four questions
├── @aroha-sdk/core Identity (did:aroha), signed envelopes, receipts, conversation profile
├── @aroha-sdk/credentials ★ Start here — mandates, attenuation, chain verification, revocation
├── @aroha-sdk/delegation Multi-hop delegation with depth limits and stitched receipts
└── @aroha-sdk/policy Deterministic limits checked before any model sees a request
Adapters (use Aroha from the stack you already run)
├── @aroha-sdk/mcp-server Remote MCP server (2026 authorization spec)
├── @aroha-sdk/mcp-guard Mandate checks in front of any MCP server
├── @aroha-sdk/mcp-bridge MCP ↔ Aroha (includes wrapMcp() one-liner)
├── @aroha-sdk/a2a-bridge A2A ↔ Aroha
└── @aroha-sdk/langchain-bridge LangChain / OpenAI functions ↔ Aroha
Extensions (maintained, optional; developed outside this repository)
└── orchestrator (sagas) · settlement · commerce · reputation
Tooling
├── @aroha-sdk/cli scaffold · convert · doctor · dev sandbox
└── create-aroha-agent npx create-aroha-agent — zero-config scaffolder
Already building agents? Add mandate verification in 2 lines.
import { verifyMandateChain } from "@aroha-sdk/credentials" ;
// In your agent's request handler. HUMAN_ROOT_DID is the human you accept authority from;
// MY_DID is this agent's own DID. The mandate travels in the signed Aroha envelope body
// (`credentialToken`); over MCP it travels in `_meta["aroha/mandate"]` with a holder proof
// (see @aroha-sdk/mcp-bridge). There is no mandate HTTP header.
const { valid , chainVerified , mandate , reason } = await verifyMandateChain (
envelope . body . credentialToken ,
undefined , // default resolver: content-addressed did:aroha:z… DIDs
{ expectedRoot : process . env . HUMAN_ROOT_DID , expectedGrantee : MY_DID } ,
) ;
if ( ! valid || ! chainVerified ) {
res . status ( 403 ) . json ( { error : reason } ) ;
} else {
// mandate.constraints.spendLimitUsd, mandate.constraints.merchantCategory, etc.
}
Already have an MCP server? Add one import:
import { wrapMcp , registryPublicKeyResolver } from "@aroha-sdk/mcp-bridge" ;
const agent = await wrapMcp ( server . listTools ( ) , {
name : "my-tools" ,
price : 0.01 ,
resolvePublicKey : registryPublicKeyResolver ( ) , // verifies who is calling
handlers : { search : async ( args ) => ( { content : [ { type : "text" , text : await search ( args ) } ] } ) } ,
} ) ;
await agent . start ( ) ;
// ✓ did:aroha:my-tools → http://localhost:3001
Using LangChain? Call any mandate-aware Aroha agent as a native tool:
import { createArohaLangChainTool } from "@aroha-sdk/langchain-bridge" ;
const tool = await createArohaLangChainTool (
"http://localhost:3001" ,
"search-flights" ,
"Search available flights by origin, destination, and date"
) ;
// plug `tool` straight into AgentExecutor, createOpenAIFunctionsAgent, etc.
Using an AI coding agent? Point it at the aroha-integrate Agent Skill (or docs/llms.txt ): it picks the smallest integration path for your code, applies it, and checks that out-of-scope calls are refused.
Scaffold a new agent pre-wired for mandate verification:
npx create-aroha-agent my-agent
cd my-agent && npm install && npm run dev
# → [my-agent] did:aroha:my-agent http://localhost:3001/aroha/v1 Ready ✓
CLI reference
aroha init my-agent # interactive scaffold
aroha init my-agent --yes # zero-config: picks all defaults, writes .env
aroha convert --from mcp --dir . # wrap existing agent
aroha doctor [endpoint] # check environment + connectivity
aroha dev # local dev registry + dashboard
Deploy
One-click templates for fast deployment.
Note: the Vercel and Cloudflare buttons clone from github.com/ArohaLabs/aroha-protocol . You can
also copy the template directory from a checkout and deploy it with the platform CLI.
packages/ SDK packages (@aroha-sdk/*)
agents/mock-provider-agent/ Reference provider agent for interoperability tests
examples/ Deployment templates (Vercel, Cloudflare, Railway, Render)
examples/sandbox-agents/ Pre-built sandbox agents (flight, hotel, weather, web-search, summarizer)
contracts/ Solidity contracts (ArohaRegistry, ArohaEscrow, ArohaReputation)
docs/ MkDocs documentation site
The Aroha runtime (desktop app) is developed in a separate, private repository and has not been released. The saga, settlement,
commerce and reputation extensions are also developed outside this repository; they stay published on npm.
Requirement
Version
Node.js
≥ 22
npm
≥ 10
Foundry and Python are only needed for the smart contract demo and MkDocs docs site respectively.
Use GitHub Issues for bugs and feature requests and
GitHub Discussions for questions and ideas,
or email the maintainer listed in MAINTAINERS.md .
See CONTRIBUTING.md (commits need a DCO sign-off) and SUPPORT.md .
For security issues, see SECURITY.md .
npm install && npm run build && npm test
License
Apache-2.0 — see LICENSE and NOTICE .
Aroha: verifiable identity and spending limits for AI agents
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Aroha: verifiable identity and spending limits for AI agents - ArohaLabs/aroha-protocol

GitHub - ArohaLabs/aroha-protocol: Aroha: verifiable identity and spending limits for AI agents · GitHub
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
ArohaLabs
aroha-protocol Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Aroha: verifiable identity and spending limits for AI agents
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
21 Commits 21 Commits Folders and files
.github .github .well-known .well-known LICENSES/ third-party LICENSES/ third-party agents/ mock-provider-agent agents/ mock-provider-agent apps apps community community contracts contracts docs docs examples examples packages packages papers papers script script scripts scripts sdks sdks skills skills spec-public spec-public test test .gitignore .gitignore .lycheeignore .lycheeignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CHARTER.md CHARTER.md CODEOWNERS CODEOWNERS CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMPATIBILITY.md COMPATIBILITY.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md TODO.md TODO.md chain-addresses.json chain-addresses.json foundry.toml foundry.toml lychee.toml lychee.toml mkdocs.yml mkdocs.yml package-lock.json package-lock.json package.json package.json quickstart.sh quickstart.sh requirements-docs.txt requirements-docs.txt tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.conformance.config.ts vitest.conformance.config.ts View all files Repository files navigation
Cryptographically bounded authority for every agent in the chain.
Aroha Mandates give orchestrators cryptographically enforceable control over every sub-agent they delegate to. No agent can exceed the scope it was granted.
Docs · Support matrix · Local sandbox · Support · Spec
import { generateKeyPair , generateDid } from "@aroha-sdk/core" ;
import { issueIntentMandate , attenuateToPayment , verifyMandateChain } from "@aroha-sdk/credentials" ;
const human = await generateKeyPair ( ) ;
const orchestrator = await generateKeyPair ( ) ;
const flightAgent = await generateKeyPair ( ) ;
// Content-addressed DIDs (did:aroha:z…) commit to their own public key,
// so a verifier needs no registry lookup to check signatures.
const humanDID = generateDid ( human . publicKey ) ;
const orchestratorDID = generateDid ( orchestrator . publicKey ) ;
const flightAgentDID = generateDid ( flightAgent . publicKey ) ;
// Human issues an intent mandate to the orchestrator ($500, 1 hour TTL)
const intent = await issueIntentMandate (
humanDID , orchestratorDID ,
{ spendLimitUsd : 500 , merchantCategory : "4511" } , // MCC 4511 = airlines
human . privateKey ,
3_600_000 ,
) ;
// Orchestrator attenuates scope downward to the flight agent ($300 of the $500).
// The child token embeds the signed chain back to the human.
const payment = await attenuateToPayment (
intent , flightAgentDID ,
{ spendLimitUsd : 300 , merchantCategory : "4511" } ,
orchestrator . privateKey ,
) ;
// Flight agent verifies every link back to the human root before acting
const result = await verifyMandateChain ( payment . token , undefined , {
expectedRoot : humanDID , // the human authority must come from
expectedGrantee : flightAgentDID , // the flight agent itself
} ) ;
// result.valid === true, result.chainVerified === true,
// result.mandate?.constraints.spendLimitUsd === 300.
// A link that widens scope, or is signed by anyone but its grantor, fails.
The problem
Agents now act for people and hand work to other agents and tools, often run by other companies. MCP and A2A connect them; identity providers name them. Nothing that travels with the work proves what each agent was allowed to do, stops the whole chain when the person changes their mind, or leaves a record anyone else can check.
Aroha answers four questions for every action: who is acting, what were they allowed to do, can it be stopped, and what actually happened. It works with MCP, A2A and your identity provider rather than replacing them. See what Aroha is (and is not) .
Human issues a mandate — signs an IntentMandate with authorized actions, a spending cap, and a TTL. Cryptographically bound to the orchestrator's DID.
Orchestrator attenuates — delegates a strict subset of that scope to the next agent. The new mandate is signed by the orchestrator. Scope narrows, never widens.
Every agent verifies — before acting, an agent verifies the mandate chain back to its human origin, offline. A forged, altered, widened or expired link fails verification.
Revoke and prove — the person can revoke at any time; the registry witness pushes the revocation to verifiers. Every allowed or refused action produces a signed receipt.
Deploy in your existing stack — bridges for MCP, LangChain, A2A, and Composio. Mandate verification drops in in two lines.
Core — the four questions
├── @aroha-sdk/core Identity (did:aroha), signed envelopes, receipts, conversation profile
├── @aroha-sdk/credentials ★ Start here — mandates, attenuation, chain verification, revocation
├── @aroha-sdk/delegation Multi-hop delegation with depth limits and stitched receipts
└── @aroha-sdk/policy Deterministic limits checked before any model sees a request
Adapters (use Aroha from the stack you already run)
├── @aroha-sdk/mcp-server Remote MCP server (2026 authorization spec)
├── @aroha-sdk/mcp-guard Mandate checks in front of any MCP server
├── @aroha-sdk/mcp-bridge MCP ↔ Aroha (includes wrapMcp() one-liner)
├── @aroha-sdk/a2a-bridge A2A ↔ Aroha
└── @aroha-sdk/langchain-bridge LangChain / OpenAI functions ↔ Aroha
Extensions (maintained, optional; developed outside this repository)
└── orchestrator (sagas) · settlement · commerce · reputation
Tooling
├── @aroha-sdk/cli scaffold · convert · doctor · dev sandbox
└── create-aroha-agent npx create-aroha-agent — zero-config scaffolder
Already building agents? Add mandate verification in 2 lines.
import { verifyMandateChain } from "@aroha-sdk/credentials" ;
// In your agent's request handler. HUMAN_ROOT_DID is the human you accept authority from;
// MY_DID is this agent's own DID. The mandate travels in the signed Aroha envelope body
// (`credentialToken`); over MCP it travels in `_meta["aroha/mandate"]` with a holder proof
// (see @aroha-sdk/mcp-bridge). There is no mandate HTTP header.
const { valid , chainVerified , mandate , reason } = await verifyMandateChain (
envelope . body . credentialToken ,
undefined , // default resolver: content-addressed did:aroha:z… DIDs
{ expectedRoot : process . env . HUMAN_ROOT_DID , expectedGrantee : MY_DID } ,
) ;
if ( ! valid || ! chainVerified ) {
res . status ( 403 ) . json ( { error : reason } ) ;
} else {
// mandate.constraints.spendLimitUsd, mandate.constraints.merchantCategory, etc.
}
Already have an MCP server? Add one import:
import { wrapMcp , registryPublicKeyResolver } from "@aroha-sdk/mcp-bridge" ;
const agent = await wrapMcp ( server . listTools ( ) , {
name : "my-tools" ,
price : 0.01 ,
resolvePublicKey : registryPublicKeyResolver ( ) , // verifies who is calling
handlers : { search : async ( args ) => ( { content : [ { type : "text" , text : await search ( args ) } ] } ) } ,
} ) ;
await agent . start ( ) ;
// ✓ did:aroha:my-tools → http://localhost:3001
Using LangChain? Call any mandate-aware Aroha agent as a native tool:
import { createArohaLangChainTool } from "@aroha-sdk/langchain-bridge" ;
const tool = await createArohaLangChainTool (
"http://localhost:3001" ,
"search-flights" ,
"Search available flights by origin, destination, and date"
) ;
// plug `tool` straight into AgentExecutor, createOpenAIFunctionsAgent, etc.
Using an AI coding agent? Point it at the aroha-integrate Agent Skill (or docs/llms.txt ): it picks the smallest integration path for your code, applies it, and checks that out-of-scope calls are refused.
Scaffold a new agent pre-wired for mandate verification:
npx create-aroha-agent my-agent
cd my-agent && npm install && npm run dev
# → [my-agent] did:aroha:my-agent http://localhost:3001/aroha/v1 Ready ✓
CLI reference
aroha init my-agent # interactive scaffold
aroha init my-agent --yes # zero-config: picks all defaults, writes .env
aroha convert --from mcp --dir . # wrap existing agent
aroha doctor [endpoint] # check environment + connectivity
aroha dev # local dev registry + dashboard
Deploy
One-click templates for fast deployment.
Note: the Vercel and Cloudflare buttons clone from github.com/ArohaLabs/aroha-protocol . You can
also copy the template directory from a checkout and deploy it with the platform CLI.
packages/ SDK packages (@aroha-sdk/*)
agents/mock-provider-agent/ Reference provider agent for interoperability tests
examples/ Deployment templates (Vercel, Cloudflare, Railway, Render)
examples/sandbox-agents/ Pre-built sandbox agents (flight, hotel, weather, web-search, summarizer)
contracts/ Solidity contracts (ArohaRegistry, ArohaEscrow, ArohaReputation)
docs/ MkDocs documentation site
The Aroha runtime (desktop app) is developed in a separate, private repository and has not been released. The saga, settlement,
commerce and reputation extensions are also developed outside this repository; they stay published on npm.
Requirement
Version
Node.js
≥ 22
npm
≥ 10
Foundry and Python are only needed for the smart contract demo and MkDocs docs site respectively.
Use GitHub Issues for bugs and feature requests and
GitHub Discussions for questions and ideas,
or email the maintainer listed in MAINTAINERS.md .
See CONTRIBUTING.md (commits need a DCO sign-off) and SUPPORT.md .
For security issues, see SECURITY.md .
npm install && npm run build && npm test
License
Apache-2.0 — see LICENSE and NOTICE .
Aroha: verifiable identity and spending limits for AI agents
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
