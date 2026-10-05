---
source: "https://github.com/CharlesStrogish/bitcoin-stratigraphy-mcp"
hn_url: "https://news.ycombinator.com/item?id=49959733"
title: "L402 edge cache for AI agents paying over Lightning"
article_title: "GitHub - CharlesStrogish/bitcoin-stratigraphy-mcp: An enterprise-grade, L402-gated AI Notary Hub and MCP data server vending thermodynamic Bitcoin telemetry. Features 15 canonical tools, OpenTimestamps state anchoring, a free sandbox, 100-call Macro-Macaroon batch payments, and LangChain/LlamaIndex\n[truncated]"
image: "https://opengraph.githubassets.com/70df76ee5558a4715046e61454b8dc21686dcf0f160495be3fe7eaf2c1e45271/CharlesStrogish/bitcoin-stratigraphy-mcp"
author: "CharlesRStrogis"
captured_at: "2026-10-05T01:39:42Z"
capture_tool: "hn-digest"
hn_id: 49959733
score: 1
comments: 0
posted_at: "2026-10-05T01:29:00Z"
tags:
  - hacker-news
---

# L402 edge cache for AI agents paying over Lightning

- HN: [49959733](https://news.ycombinator.com/item?id=49959733)
- Source: [github.com](https://github.com/CharlesStrogish/bitcoin-stratigraphy-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-10-05T01:29:00Z

## Translation

Title: L402 edge cache for AI agents paying over Lightning
Article title: GitHub - CharlesStrogish/bitcoin-stratigraphy-mcp: An enterprise-grade, L402-gated AI Notary Hub and MCP data server vending thermodynamic Bitcoin telemetry. Features 15 canonical tools, OpenTimestamps state anchoring, a free sandbox, 100-call Macro-Macaroon batch payments, and LangChain/LlamaIndex
[truncated]
Description: An enterprise-grade, L402-gated AI Notary Hub and MCP data server vending thermodynamic Bitcoin telemetry. Features 15 canonical tools, OpenTimestamps state anchoring, a free sandbox, 100-call Macro-Macaroon batch payments, and LangChain/LlamaIndex cookbooks. Built for autonomous AI agents settling
[truncated]

Article text:
GitHub - CharlesStrogish/bitcoin-stratigraphy-mcp: An enterprise-grade, L402-gated AI Notary Hub and MCP data server vending thermodynamic Bitcoin telemetry. Features 15 canonical tools, OpenTimestamps state anchoring, a free sandbox, 100-call Macro-Macaroon batch payments, and LangChain/LlamaIndex cookbooks. Built for autonomous AI agents settling on the Lightning Network. · GitHub
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
CharlesStrogish
/
bitcoin-stratigraphy-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
cookbooks cookbooks examples examples scripts scripts sdk sdk test test .gitignore .gitignore LICENSE LICENSE README.md README.md mcp.json mcp.json openapi.json openapi.json openapi.yaml openapi.yaml package.json package.json smithery.yaml smithery.yaml View all files Repository files navigation
Bitcoin Stratigraphy MCP — Developer Distribution
Release: 1.1.0 · MCP registry name: crstrogish/bitcoin-stratigraphy
A standalone, remote-first integration repository for the Bitcoin Stratigraphy
MCP engine: registry metadata, OpenAPI contracts, dependency-free clients, and
native framework agent examples. No Replit workspace, wallet, API key, or model
provider account is needed for the free sandbox.
This bundle connects to the hosted engine ; it does not contain or claim to
run the complete settlement server offline. Client sources can be run locally
against the live service or an independently operated compatible MCP endpoint.
There are no workspace aliases, vendored credentials, private signing keys,
operator state, or dependency on a parent repository.
LangChain execution is supported. LlamaIndex source is preserved, but dependency
installation and execution are temporarily unavailable : its NLTK dependency is
affected by GHSA-8mgp-746c-j5xp .
The LlamaIndex requirements file intentionally installs nothing. Do not bypass
security checks or treat the badge as a claim of supported execution. Restore
installation only after auditing a safe, complete dependency tree.
Connect to the live MCP engine
Surface
URL
Streamable HTTP MCP
https://bitcoin-stratigraphy-dashboard.replit.app/mcp
Developer portal
https://bitcoin-stratigraphy-dashboard.replit.app/docs
Live OpenAPI
https://bitcoin-stratigraphy-dashboard.replit.app/openapi.json
Use the MCP URL in any Streamable HTTP client. Import mcp.json where
your client supports URL-based MCP configuration; some hosts require an explicit
transport type in their own configuration format. smithery.yaml
is canonical remote registration metadata , not a local-server launch recipe.
The server advertises 15 canonical capabilities. initialize and tools/list
are free protocol discovery operations. Exactly two tool executions are free:
Catalog discovery is not payment authorization. Other canonical tools require
paid access; query live schemas and payment terms before authorizing purchases.
Free sandbox quickstart — no wallet
From this repository's root, with no package installation :
node examples/free-sandbox.mjs
This negotiates MCP, sends the initialized notification, then calls only
stratigraphy.list and proof.list . It does not sign, pay, attach credentials, or
make LLM requests. It fails rather than paying if a free call returns HTTP 402.
To use your own compatible endpoint:
node examples/free-sandbox.mjs --url http://127.0.0.1:5000/mcp
This is a client connection, not a command to start the production engine.
The bundled client works with the standard library; no framework installation
is required:
python -m examples.free_sandbox
Equivalent copy-paste Python:
from sdk . python . client import BitcoinStratigraphyClient
client = BitcoinStratigraphyClient ()
print ( client . call_tool ( "stratigraphy.list" , {}))
print ( client . call_tool ( "proof.list" , {}))
Run code from the repository root so the sdk namespace is importable. TLS
verification remains enabled. If your organization uses a private CA, configure
its trusted certificate store; do not disable verification.
L402: 100-call Macro-Macaroon batches
Paid capabilities use tool-scoped credentials. To request a batch, include:
X-L402-Batch-Size: 100
For example, this request asks for a notarization batch challenge ; it does
not pay an invoice:
curl --include --request POST \
' https://bitcoin-stratigraphy-dashboard.replit.app/mcp ' \
--header ' Content-Type: application/json ' \
--header ' Accept: application/json, text/event-stream ' \
--header ' X-L402-Batch-Size: 100 ' \
--data ' {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"proof.notarize","arguments":{"hash":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}}} '
Inspect HTTP 402, the L402 WWW-Authenticate challenge, invoice amount,
expiry, and exact tool scope.
Pay only through a trusted Lightning wallet/client with explicit spending
limits. Verify the settlement proof matches the challenged payment hash.
Retry the same tool call with the batch header and:
Authorization: L402 <macaroon>:<preimage>
X-L402-Batch-Size: 100
The registered notarization quote is 250 sats/call × 100 = 25,000 sats .
The live invoice is authoritative; registry target-price examples are not a
price guarantee. A proof.notarize pass does not unlock other paid tools.
Budgets, expiration, revocation, and inherited caveats still apply on every use.
Attenuation delegates or narrows scope; it does not reset the parent budget.
Native framework adapters are not wallets . Provide a prepaid credential via
L402_AUTHORIZATION , or a trusted loopback Streamable HTTP wallet proxy via
L402_PROXY_URL with explicit spending limits. Do not commit either credentials
or model-provider secrets. The older one-shot SDK payment callbacks are not
automatic Macro-Macaroon batching implementations.
LangChain: native MCP and ReAct loops
Use a separate consumer Python environment, not an older MCP 1.x environment:
python -m pip install -r cookbooks/requirements-langchain.txt
python cookbooks/langchain_agent.py
Default mode discovers tools only. To opt into an LLM and paid notarization,
configure your provider key and a prepaid proof.notarize batch credential
through your secret manager, then run:
python cookbooks/langchain_agent.py --run
Copy-paste free-tool-only ReAct loop (MCP calls are free; model usage is billed
by your model provider):
import asyncio
from fastmcp . client import Client
from fastmcp . client . transports import StreamableHttpTransport
from langchain . mcp import MCPAdapter
from langchain . agents import create_agent
from langchain_openai import ChatOpenAI
async def main ():
transport = StreamableHttpTransport (
url = "https://bitcoin-stratigraphy-dashboard.replit.app/mcp"
)
client = Client ( transport , mode = "legacy" , timeout = 30 )
async with MCPAdapter ( client ) as adapter :
tools = await adapter . list_tools ()
free = [
t . model_copy ( update = { "name" : t . name . replace ( "." , "_" )})
for t in tools if t . name in { "stratigraphy.list" , "proof.list" }
]
if len ( free ) != 2 :
raise RuntimeError ( "Expected both free sandbox capabilities" )
agent = create_agent (
model = ChatOpenAI ( model = "gpt-4.1-mini" , temperature = 0 ),
tools = free ,
system_prompt = "Use only these discovery tools. Report actual results." ,
)
result = await agent . ainvoke (
{ "messages" : [{ "role" : "user" , "content" :
"List indexed stratigraphy bounds and supported proof operations." }]}
)
print ( result [ "messages" ][ - 1 ]. content )
asyncio . run ( main ())
create_agent implements the current ReAct model → tool → observation loop.
Underscore aliases satisfy model-provider naming rules; the wrappers retain the
original dotted remote tool names.
LlamaIndex: preserved native agent example
Source-only until the dependency security hold is lifted. The following
copy-paste example documents the supported API shape for a future audited
environment; it is not an instruction to install vulnerable dependencies now.
The full paid-notarization example is cookbooks/llamaindex_agent.py .
import asyncio
from llama_index . tools . mcp import BasicMCPClient , McpToolSpec
from llama_index . core . agent . workflow import ReActAgent
from llama_index . llms . openai import OpenAI
async def main ():
client = BasicMCPClient (
"https://bitcoin-stratigraphy-dashboard.replit.app/mcp" , timeout = 30
)
try :
tools = await McpToolSpec ( client = client ). to_tool_list_async ()
free = [ t for t in tools if t . metadata . name in
{ "stratigraphy.list" , "proof.list" }]
if len ( free ) != 2 :
raise RuntimeError ( "Expected both free sandbox capabilities" )
for tool in free :
tool . metadata . name = tool . metadata . name . replace ( "." , "_" )
agent = ReActAgent (
tools = free ,
llm = OpenAI ( model = "gpt-4.1-mini" , temperature = 0 ),
system_prompt = "Use only the discovery tools. Report actual results." ,
timeout = 90 ,
)
response = await agent . run (
user_msg = "List indexed stratigraphy bounds and supported proof operations."
)
print ( str ( response ))
finally :
await client . http_client . aclose ()
asyncio . run ( main ())
Model credentials and charges are separate from MCP payments.
sdk/typescript contains both the REST client and MCP client
sources, independent compiler configurations, local tests, and its own package
manifest. SDK version 0.1.0 is independent of registry release 1.1.0.
cd sdk/typescript
npm install
npm test
npm run build
The root package manifest retains the canonical MCP registry identity and
release version but replaces monorepo scripts/dependencies with portable
distribution commands. It is not the private workspace orchestration manifest.
Repository contents and export checks
README.md Distribution and operational guide
package.json Portable MCP metadata and sandbox/test commands
smithery.yaml Canonical remote registration manifest
mcp.json Remote MCP client configuration
openapi.yaml Canonical self-contained REST/OpenAPI contract
openapi.json JSON rendering of the same canonical contract
LICENSE MIT license
cookbooks/ Native framework sources, requirements, and offline tests
sdk/typescript/ TypeScript REST/MCP clients and standalone package
sdk/python/ Python REST/MCP clients and offline tests
examples/ Wallet-free Node/Python sandbox entrypoints
test/ Node sandbox transport and safety tests
scripts/verify-export.mjs Portable structure, import, and reference checks
npm run verify
npm test
python -m unittest discover -s cookbooks -p ' test_*.py '
python -m unittest discover -s sdk/python -p ' test_*.py '
verify checks relative source imports and Markdown file references, script and
SDK package paths, JSON syntax, and absence of workspace dependency aliases.
Comp

[truncated]

## Original Extract

An enterprise-grade, L402-gated AI Notary Hub and MCP data server vending thermodynamic Bitcoin telemetry. Features 15 canonical tools, OpenTimestamps state anchoring, a free sandbox, 100-call Macro-Macaroon batch payments, and LangChain/LlamaIndex cookbooks. Built for autonomous AI agents settling
[truncated]

GitHub - CharlesStrogish/bitcoin-stratigraphy-mcp: An enterprise-grade, L402-gated AI Notary Hub and MCP data server vending thermodynamic Bitcoin telemetry. Features 15 canonical tools, OpenTimestamps state anchoring, a free sandbox, 100-call Macro-Macaroon batch payments, and LangChain/LlamaIndex cookbooks. Built for autonomous AI agents settling on the Lightning Network. · GitHub
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
CharlesStrogish
/
bitcoin-stratigraphy-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
cookbooks cookbooks examples examples scripts scripts sdk sdk test test .gitignore .gitignore LICENSE LICENSE README.md README.md mcp.json mcp.json openapi.json openapi.json openapi.yaml openapi.yaml package.json package.json smithery.yaml smithery.yaml View all files Repository files navigation
Bitcoin Stratigraphy MCP — Developer Distribution
Release: 1.1.0 · MCP registry name: crstrogish/bitcoin-stratigraphy
A standalone, remote-first integration repository for the Bitcoin Stratigraphy
MCP engine: registry metadata, OpenAPI contracts, dependency-free clients, and
native framework agent examples. No Replit workspace, wallet, API key, or model
provider account is needed for the free sandbox.
This bundle connects to the hosted engine ; it does not contain or claim to
run the complete settlement server offline. Client sources can be run locally
against the live service or an independently operated compatible MCP endpoint.
There are no workspace aliases, vendored credentials, private signing keys,
operator state, or dependency on a parent repository.
LangChain execution is supported. LlamaIndex source is preserved, but dependency
installation and execution are temporarily unavailable : its NLTK dependency is
affected by GHSA-8mgp-746c-j5xp .
The LlamaIndex requirements file intentionally installs nothing. Do not bypass
security checks or treat the badge as a claim of supported execution. Restore
installation only after auditing a safe, complete dependency tree.
Connect to the live MCP engine
Surface
URL
Streamable HTTP MCP
https://bitcoin-stratigraphy-dashboard.replit.app/mcp
Developer portal
https://bitcoin-stratigraphy-dashboard.replit.app/docs
Live OpenAPI
https://bitcoin-stratigraphy-dashboard.replit.app/openapi.json
Use the MCP URL in any Streamable HTTP client. Import mcp.json where
your client supports URL-based MCP configuration; some hosts require an explicit
transport type in their own configuration format. smithery.yaml
is canonical remote registration metadata , not a local-server launch recipe.
The server advertises 15 canonical capabilities. initialize and tools/list
are free protocol discovery operations. Exactly two tool executions are free:
Catalog discovery is not payment authorization. Other canonical tools require
paid access; query live schemas and payment terms before authorizing purchases.
Free sandbox quickstart — no wallet
From this repository's root, with no package installation :
node examples/free-sandbox.mjs
This negotiates MCP, sends the initialized notification, then calls only
stratigraphy.list and proof.list . It does not sign, pay, attach credentials, or
make LLM requests. It fails rather than paying if a free call returns HTTP 402.
To use your own compatible endpoint:
node examples/free-sandbox.mjs --url http://127.0.0.1:5000/mcp
This is a client connection, not a command to start the production engine.
The bundled client works with the standard library; no framework installation
is required:
python -m examples.free_sandbox
Equivalent copy-paste Python:
from sdk . python . client import BitcoinStratigraphyClient
client = BitcoinStratigraphyClient ()
print ( client . call_tool ( "stratigraphy.list" , {}))
print ( client . call_tool ( "proof.list" , {}))
Run code from the repository root so the sdk namespace is importable. TLS
verification remains enabled. If your organization uses a private CA, configure
its trusted certificate store; do not disable verification.
L402: 100-call Macro-Macaroon batches
Paid capabilities use tool-scoped credentials. To request a batch, include:
X-L402-Batch-Size: 100
For example, this request asks for a notarization batch challenge ; it does
not pay an invoice:
curl --include --request POST \
' https://bitcoin-stratigraphy-dashboard.replit.app/mcp ' \
--header ' Content-Type: application/json ' \
--header ' Accept: application/json, text/event-stream ' \
--header ' X-L402-Batch-Size: 100 ' \
--data ' {"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"proof.notarize","arguments":{"hash":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}}} '
Inspect HTTP 402, the L402 WWW-Authenticate challenge, invoice amount,
expiry, and exact tool scope.
Pay only through a trusted Lightning wallet/client with explicit spending
limits. Verify the settlement proof matches the challenged payment hash.
Retry the same tool call with the batch header and:
Authorization: L402 <macaroon>:<preimage>
X-L402-Batch-Size: 100
The registered notarization quote is 250 sats/call × 100 = 25,000 sats .
The live invoice is authoritative; registry target-price examples are not a
price guarantee. A proof.notarize pass does not unlock other paid tools.
Budgets, expiration, revocation, and inherited caveats still apply on every use.
Attenuation delegates or narrows scope; it does not reset the parent budget.
Native framework adapters are not wallets . Provide a prepaid credential via
L402_AUTHORIZATION , or a trusted loopback Streamable HTTP wallet proxy via
L402_PROXY_URL with explicit spending limits. Do not commit either credentials
or model-provider secrets. The older one-shot SDK payment callbacks are not
automatic Macro-Macaroon batching implementations.
LangChain: native MCP and ReAct loops
Use a separate consumer Python environment, not an older MCP 1.x environment:
python -m pip install -r cookbooks/requirements-langchain.txt
python cookbooks/langchain_agent.py
Default mode discovers tools only. To opt into an LLM and paid notarization,
configure your provider key and a prepaid proof.notarize batch credential
through your secret manager, then run:
python cookbooks/langchain_agent.py --run
Copy-paste free-tool-only ReAct loop (MCP calls are free; model usage is billed
by your model provider):
import asyncio
from fastmcp . client import Client
from fastmcp . client . transports import StreamableHttpTransport
from langchain . mcp import MCPAdapter
from langchain . agents import create_agent
from langchain_openai import ChatOpenAI
async def main ():
transport = StreamableHttpTransport (
url = "https://bitcoin-stratigraphy-dashboard.replit.app/mcp"
)
client = Client ( transport , mode = "legacy" , timeout = 30 )
async with MCPAdapter ( client ) as adapter :
tools = await adapter . list_tools ()
free = [
t . model_copy ( update = { "name" : t . name . replace ( "." , "_" )})
for t in tools if t . name in { "stratigraphy.list" , "proof.list" }
]
if len ( free ) != 2 :
raise RuntimeError ( "Expected both free sandbox capabilities" )
agent = create_agent (
model = ChatOpenAI ( model = "gpt-4.1-mini" , temperature = 0 ),
tools = free ,
system_prompt = "Use only these discovery tools. Report actual results." ,
)
result = await agent . ainvoke (
{ "messages" : [{ "role" : "user" , "content" :
"List indexed stratigraphy bounds and supported proof operations." }]}
)
print ( result [ "messages" ][ - 1 ]. content )
asyncio . run ( main ())
create_agent implements the current ReAct model → tool → observation loop.
Underscore aliases satisfy model-provider naming rules; the wrappers retain the
original dotted remote tool names.
LlamaIndex: preserved native agent example
Source-only until the dependency security hold is lifted. The following
copy-paste example documents the supported API shape for a future audited
environment; it is not an instruction to install vulnerable dependencies now.
The full paid-notarization example is cookbooks/llamaindex_agent.py .
import asyncio
from llama_index . tools . mcp import BasicMCPClient , McpToolSpec
from llama_index . core . agent . workflow import ReActAgent
from llama_index . llms . openai import OpenAI
async def main ():
client = BasicMCPClient (
"https://bitcoin-stratigraphy-dashboard.replit.app/mcp" , timeout = 30
)
try :
tools = await McpToolSpec ( client = client ). to_tool_list_async ()
free = [ t for t in tools if t . metadata . name in
{ "stratigraphy.list" , "proof.list" }]
if len ( free ) != 2 :
raise RuntimeError ( "Expected both free sandbox capabilities" )
for tool in free :
tool . metadata . name = tool . metadata . name . replace ( "." , "_" )
agent = ReActAgent (
tools = free ,
llm = OpenAI ( model = "gpt-4.1-mini" , temperature = 0 ),
system_prompt = "Use only the discovery tools. Report actual results." ,
timeout = 90 ,
)
response = await agent . run (
user_msg = "List indexed stratigraphy bounds and supported proof operations."
)
print ( str ( response ))
finally :
await client . http_client . aclose ()
asyncio . run ( main ())
Model credentials and charges are separate from MCP payments.
sdk/typescript contains both the REST client and MCP client
sources, independent compiler configurations, local tests, and its own package
manifest. SDK version 0.1.0 is independent of registry release 1.1.0.
cd sdk/typescript
npm install
npm test
npm run build
The root package manifest retains the canonical MCP registry identity and
release version but replaces monorepo scripts/dependencies with portable
distribution commands. It is not the private workspace orchestration manifest.
Repository contents and export checks
README.md Distribution and operational guide
package.json Portable MCP metadata and sandbox/test commands
smithery.yaml Canonical remote registration manifest
mcp.json Remote MCP client configuration
openapi.yaml Canonical self-contained REST/OpenAPI contract
openapi.json JSON rendering of the same canonical contract
LICENSE MIT license
cookbooks/ Native framework sources, requirements, and offline tests
sdk/typescript/ TypeScript REST/MCP clients and standalone package
sdk/python/ Python REST/MCP clients and offline tests
examples/ Wallet-free Node/Python sandbox entrypoints
test/ Node sandbox transport and safety tests
scripts/verify-export.mjs Portable structure, import, and reference checks
npm run verify
npm test
python -m unittest discover -s cookbooks -p ' test_*.py '
python -m unittest discover -s sdk/python -p ' test_*.py '
verify checks relative source imports and Markdown file references, script and
SDK package paths, JSON syntax, and absence of workspace dependency aliases.
Comp

[truncated]
