---
source: "https://github.com/mrpkk/soma-protocol"
hn_url: "https://news.ycombinator.com/item?id=50216800"
title: "Soma – memory and proofs for AI agents, measured not asserted"
article_title: "GitHub - mrpkk/soma-protocol: Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подписью; память в три яруса; набор тестов соответствия. · GitHub"
image: "https://opengraph.githubassets.com/610f7a9b40616bda5d8b02cd4224e3dc3731160fc7879ada38588919cfb88e83/mrpkk/soma-protocol"
author: "vnutrivas"
captured_at: "2026-10-24T00:18:38Z"
capture_tool: "hn-digest"
hn_id: 50216800
score: 1
comments: 0
posted_at: "2026-10-23T23:31:45Z"
tags:
  - hacker-news
---

# Soma – memory and proofs for AI agents, measured not asserted

- HN: [50216800](https://news.ycombinator.com/item?id=50216800)
- Source: [github.com](https://github.com/mrpkk/soma-protocol)
- Score: 1
- Comments: 0
- Posted: 2026-10-23T23:31:45Z

## Translation

Title: Soma – memory and proofs for AI agents, measured not asserted
Article title: GitHub - mrpkk/soma-protocol: Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подписью; память в три яруса; набор тестов соответствия. · GitHub
Description: Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подписью; память в три яруса; набор тестов соответствия. - mrpkk/soma-protocol

Article text:
GitHub - mrpkk/soma-protocol: Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подписью; память в три яруса; набор тестов соответствия. · GitHub
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
mrpkk
soma-protocol Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подписью; память в три яруса; набор тестов соответствия.
github.com/mrpkk/mitra-hub Topics
Readme MIT license Security policy
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
53 Commits 53 Commits Folders and files
.github/ workflows .github/ workflows bench bench docs docs schemas schemas src/ soma src/ soma tests tests tools tools .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE PROGRESS.md PROGRESS.md PROTOCOL.md PROTOCOL.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md agents.json agents.json pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
One protocol for agents to trust each other, remember together, and pay for what they do.
Try it without installing anything: memory demo in the browser —
https://rta.lat/hub/demo/recall?q=pricing · storefront and pricing — https://rta.lat/hub ·
public MCP node, no token — https://rta.lat/soma/mcp
A message between agents is an envelope , not a chat line. The envelope carries the goal, the task,
a hard token budget, a mandate, a price, a deadline and a signature. Memory, transport, money and
history are just carriers of that envelope.
POST /v1/messages
{"goal": "verify the artifact", "task": "run the linter", "to": "verifier",
"token_budget": 2000, "mandate": {"max_cost_usd": 0.5}, "price_usdc": 0.005}
Budget travels with the message. The receiver knows the limit before it starts.
Refusals come with a reason. Every decline names the rule that fired.
Degradation is declared. If context did not fit, the report says what was dropped.
Signatures verify offline. Ed25519, public key included — no service, no secret.
Try it in the browser: https://rta.lat/hub/verify (tamper a field and watch it fail).
Cost is measured, not listed. Tokens, CPU, bytes, writes, external calls.
pip install -e " .[dev] "
soma keygen # ключ подписи Ed25519
soma serve # хаб на http://127.0.0.1:8765
soma send --to research --goal " find the market " --task " scan 3 sources " --budget 2000
soma inbox --agent research
Or in Python:
from soma import Envelope , Hub
hub = Hub ()
env = Envelope ( goal = "verify the artifact" , task = "run the linter" , to = "verifier" ,
token_budget = 2000 )
hub . send ( env ) # подписан и поставлен в очередь
hub . inbox ( "verifier" ) # что ждёт агента
hub . claim ( env . id , "verifier" ) # забрать (атомарно)
hub . complete ( env . id , { "text" : "ok" })
Why
Agents already work. What they lack is a shared nervous system: each remembers only its own
session, trusts on faith, spends without a mandate, and cannot prove what it did. Memory,
provenance, authorization and accounting all exist — but as separate products that do not talk
to each other.
Soma is the glue. It is deliberately small: the core is a protocol , a reference
implementation and a conformance suite . Not a platform, not a framework, not a vendor.
Piece
What it is
src/soma/envelope.py
the protocol core: envelope, canonicalization, signing, mandate rules, router, JSON Schema
src/soma/memory.py
three memory tiers — hot queue, warm facts (SQLite + FTS5), cold artifacts by hash
src/soma/server.py
HTTP transport (stdlib only) — /v1/messages , /v1/inbox , /v1/claim , /v1/complete , /v1/route
src/soma/mcp_adapter.py
MCP adapter — the same operations as MCP tools ( soma_send , soma_inbox , soma_claim , soma_complete , soma_route , soma_verify , soma_stats )
src/soma/payments.py
usage metering (tokens/CPU/bytes/writes/calls) and the x402-compatible payment contract
src/soma/anchors.py
append-only history chain (hash-linked), with a seam for Arweave/IPFS/Nostr
src/soma/registry.py
the agent registry — a file, so it travels with the agent
tests/test_protocol.py
property tests: signature breaks on tampering, mandate refuses with a rule name, budget is hard, TTL expires
tests/test_conformance.py
the conformance suite — any implementation can pass it
schemas/envelope.schema.json
the wire format (exported from the code — one source of truth)
PROTOCOL.md
the specification
Build your own hub, prove it speaks Mitra
The conformance suite is the point of this project. Two ways to run it:
# 1) any MCP endpoint, including the language you wrote yours in
python3 tools/check_conformance.py https://rta.lat/soma/mcp
# 2) a Python implementation, in-process
MITRA_CONFORMANCE_IMPL= " my_package.hub:MyHubFactory " pytest tests/test_conformance.py -v
A public demo node runs at https://rta.lat/soma/mcp — no token, 30 requests an hour
per client, and it forces every envelope it stores to visibility: public . Nothing
private lives there. Step-by-step walkthrough: docs/IMPLEMENT.md .
A factory receives a temp directory and must return an object with
send , inbox , claim , complete , route , verify , stats .
Twelve checks: signature verifies, tampering breaks it, budget is respected, degradation is
declared, expired envelopes are not delivered, claims are exclusive, refusals explain themselves,
completion closes the envelope, budget + mandate travel with the envelope, a mandate refusal names
a rule, and a repeated envelope id does not overwrite the original.
Composing with what already exists
Soma does not replace any of these — it carries envelopes between them:
MCP — an adapter exposes the same operations as MCP tools.
A2A — agents.json in this repo is an A2A Agent Card; envelopes map onto A2A tasks.
x402 — the payment header contract is defined in src/soma/payments.py .
OPA / Cedar — the mandate interface is a plain function; policy engines plug in.
Sigstore / EIP-712 — same primitive (signed payload, public key travels with it).
Arweave / IPFS / Nostr — the anchor interface is a seam; local hash chain works offline.
This is the whole point of the split, so it is written down plainly.
The reference implementations in this repository are deliberately minimal : enough to run, test
and understand. They are not the engines. Our commercial bundle (Mitra Core) plugs into the same
seams and is the best implementation of them we have — and it is not in this repository.
If you are reading this to decide whether to build on Soma: yes, build on it. If you are reading
this to decide whether to use our hosted version instead of running your own: that is a fair
question, and the answer depends on whether you want guarantees, the shared history and the
managed key handling, or you would rather run everything yourself.
The five seams are protocols ( soma/interfaces.py ): Signer , MandateEvaluator , MemoryTier ,
Meter , Anchor . A plugin is any object that satisfies them.
# my_engine.py
from soma . envelope import Mandate
class StrictMandate :
name = "my-engine"
def check ( self , action , mandate , only = None ):
if action . get ( "payee" ) == "mallory" :
return False , "payee: blocked by my engine"
return True , "ok"
def components ():
return { "mandate" : StrictMandate ()}
Load it two ways:
# 1) environment — for experiments and local runs
SOMA_PLUGINS= " my_engine:components " soma serve
# 2) entry point — for a distributed package
# [project.entry-points."soma.plugins"]
# my_engine = "my_engine:components"
pip install my-engine && soma serve
GET /healthz reports what is actually connected:
{ "components" : { "signer" : " ed25519 " , "mandate" : " my-engine " , "meter" : " usage " ,
"memory" : " Memory " , "anchor" : " LocalChain " , "plugins" : [ " my-engine " ]}}
A broken plugin is skipped, not fatal: the hub starts with the reference implementations and says
so in /healthz .
Who is who, and why it matters
The registry is a file. An agent joins it by proving it holds its key: the hub issues a
one-time nonce, the agent signs it, and a foreign key cannot get in. When work is done, the
executor signs a receipt ; the hub verifies it and counts. Actions refused by a mandate are
counted as breaches. Reputation is then a derived fact — verified_receipts − 4 × breaches —
that anyone can recompute from the receipts, with no opinion and no stars.
Nodes can share work and history. The default is private : an envelope leaves
the node only if it is explicitly marked shared or public , and every envelope
that arrives is signed and checked against the receiver's registry. History is a
hash-linked chain that another node can import and verify — so it outlives the
node that lived it. No blockchain is required for any of this; an external anchor
(Arweave, IPFS) can publish the chain when you want one.
By default the reference transport is local : no authentication, binds 127.0.0.1 — a
component to put behind your own gateway. For a public deployment there is a managed profile:
SOMA_MANAGED=1 SOMA_TOKEN= < secret > SOMA_QUOTA_PER_HOUR=100 soma serve
Variable
Meaning
SOMA_MANAGED=1
require Authorization: Bearer $SOMA_TOKEN on every request
SOMA_TOKEN
the token; if it is not set, the managed profile denies everything
SOMA_QUOTA_PER_HOUR
max envelopes per sender per hour (0 = no limit)
SOMA_MAX_BODY
max request body in bytes (default 1 MiB)
Without a token configured, the managed profile is closed — a misconfiguration must not open
the door. The quota is counted per sender, so one noisy agent cannot spend another's budget.
The reference HTTP transport has no authentication and binds 127.0.0.1 . Put it behind
your own gateway (TLS, auth, rate limiting) before exposing it. Bodies are capped at 1 MiB.
A signature proves integrity by itself. It proves origin only when the sender's public
key is in the agent registry — otherwise anyone can sign their own envelope with their own
key. The reference implementation checks the registry when it has one.
The local history chain proves ordering and continuity; the newest entry needs an external
anchor to be final.
Memory quality is measured, not asserted. The benchmark runs offline, without an LLM:
bench/memory_bench.py (numbers) and tests/test_recall_abilities.py (abilities).
What abstention means here: a question with no supporting fact returns nothing, rather than
the nearest plausible neighbour. A memory that always answers is indistinguishable from one
that invents.
Beta. The protocol is versioned ( soma/1 ); envelopes are signed, budgets are hard, refusals are
explained. See ROADMAP.md for what is next and SECURITY.md for the trust model.
Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подпись

[truncated]

## Original Extract

Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подписью; память в три яруса; набор тестов соответствия. - mrpkk/soma-protocol

GitHub - mrpkk/soma-protocol: Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подписью; память в три яруса; набор тестов соответствия. · GitHub
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
mrpkk
soma-protocol Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подписью; память в три яруса; набор тестов соответствия.
github.com/mrpkk/mitra-hub Topics
Readme MIT license Security policy
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
53 Commits 53 Commits Folders and files
.github/ workflows .github/ workflows bench bench docs docs schemas schemas src/ soma src/ soma tests tests tools tools .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE PROGRESS.md PROGRESS.md PROTOCOL.md PROTOCOL.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md agents.json agents.json pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
One protocol for agents to trust each other, remember together, and pay for what they do.
Try it without installing anything: memory demo in the browser —
https://rta.lat/hub/demo/recall?q=pricing · storefront and pricing — https://rta.lat/hub ·
public MCP node, no token — https://rta.lat/soma/mcp
A message between agents is an envelope , not a chat line. The envelope carries the goal, the task,
a hard token budget, a mandate, a price, a deadline and a signature. Memory, transport, money and
history are just carriers of that envelope.
POST /v1/messages
{"goal": "verify the artifact", "task": "run the linter", "to": "verifier",
"token_budget": 2000, "mandate": {"max_cost_usd": 0.5}, "price_usdc": 0.005}
Budget travels with the message. The receiver knows the limit before it starts.
Refusals come with a reason. Every decline names the rule that fired.
Degradation is declared. If context did not fit, the report says what was dropped.
Signatures verify offline. Ed25519, public key included — no service, no secret.
Try it in the browser: https://rta.lat/hub/verify (tamper a field and watch it fail).
Cost is measured, not listed. Tokens, CPU, bytes, writes, external calls.
pip install -e " .[dev] "
soma keygen # ключ подписи Ed25519
soma serve # хаб на http://127.0.0.1:8765
soma send --to research --goal " find the market " --task " scan 3 sources " --budget 2000
soma inbox --agent research
Or in Python:
from soma import Envelope , Hub
hub = Hub ()
env = Envelope ( goal = "verify the artifact" , task = "run the linter" , to = "verifier" ,
token_budget = 2000 )
hub . send ( env ) # подписан и поставлен в очередь
hub . inbox ( "verifier" ) # что ждёт агента
hub . claim ( env . id , "verifier" ) # забрать (атомарно)
hub . complete ( env . id , { "text" : "ok" })
Why
Agents already work. What they lack is a shared nervous system: each remembers only its own
session, trusts on faith, spends without a mandate, and cannot prove what it did. Memory,
provenance, authorization and accounting all exist — but as separate products that do not talk
to each other.
Soma is the glue. It is deliberately small: the core is a protocol , a reference
implementation and a conformance suite . Not a platform, not a framework, not a vendor.
Piece
What it is
src/soma/envelope.py
the protocol core: envelope, canonicalization, signing, mandate rules, router, JSON Schema
src/soma/memory.py
three memory tiers — hot queue, warm facts (SQLite + FTS5), cold artifacts by hash
src/soma/server.py
HTTP transport (stdlib only) — /v1/messages , /v1/inbox , /v1/claim , /v1/complete , /v1/route
src/soma/mcp_adapter.py
MCP adapter — the same operations as MCP tools ( soma_send , soma_inbox , soma_claim , soma_complete , soma_route , soma_verify , soma_stats )
src/soma/payments.py
usage metering (tokens/CPU/bytes/writes/calls) and the x402-compatible payment contract
src/soma/anchors.py
append-only history chain (hash-linked), with a seam for Arweave/IPFS/Nostr
src/soma/registry.py
the agent registry — a file, so it travels with the agent
tests/test_protocol.py
property tests: signature breaks on tampering, mandate refuses with a rule name, budget is hard, TTL expires
tests/test_conformance.py
the conformance suite — any implementation can pass it
schemas/envelope.schema.json
the wire format (exported from the code — one source of truth)
PROTOCOL.md
the specification
Build your own hub, prove it speaks Mitra
The conformance suite is the point of this project. Two ways to run it:
# 1) any MCP endpoint, including the language you wrote yours in
python3 tools/check_conformance.py https://rta.lat/soma/mcp
# 2) a Python implementation, in-process
MITRA_CONFORMANCE_IMPL= " my_package.hub:MyHubFactory " pytest tests/test_conformance.py -v
A public demo node runs at https://rta.lat/soma/mcp — no token, 30 requests an hour
per client, and it forces every envelope it stores to visibility: public . Nothing
private lives there. Step-by-step walkthrough: docs/IMPLEMENT.md .
A factory receives a temp directory and must return an object with
send , inbox , claim , complete , route , verify , stats .
Twelve checks: signature verifies, tampering breaks it, budget is respected, degradation is
declared, expired envelopes are not delivered, claims are exclusive, refusals explain themselves,
completion closes the envelope, budget + mandate travel with the envelope, a mandate refusal names
a rule, and a repeated envelope id does not overwrite the original.
Composing with what already exists
Soma does not replace any of these — it carries envelopes between them:
MCP — an adapter exposes the same operations as MCP tools.
A2A — agents.json in this repo is an A2A Agent Card; envelopes map onto A2A tasks.
x402 — the payment header contract is defined in src/soma/payments.py .
OPA / Cedar — the mandate interface is a plain function; policy engines plug in.
Sigstore / EIP-712 — same primitive (signed payload, public key travels with it).
Arweave / IPFS / Nostr — the anchor interface is a seam; local hash chain works offline.
This is the whole point of the split, so it is written down plainly.
The reference implementations in this repository are deliberately minimal : enough to run, test
and understand. They are not the engines. Our commercial bundle (Mitra Core) plugs into the same
seams and is the best implementation of them we have — and it is not in this repository.
If you are reading this to decide whether to build on Soma: yes, build on it. If you are reading
this to decide whether to use our hosted version instead of running your own: that is a fair
question, and the answer depends on whether you want guarantees, the shared history and the
managed key handling, or you would rather run everything yourself.
The five seams are protocols ( soma/interfaces.py ): Signer , MandateEvaluator , MemoryTier ,
Meter , Anchor . A plugin is any object that satisfies them.
# my_engine.py
from soma . envelope import Mandate
class StrictMandate :
name = "my-engine"
def check ( self , action , mandate , only = None ):
if action . get ( "payee" ) == "mallory" :
return False , "payee: blocked by my engine"
return True , "ok"
def components ():
return { "mandate" : StrictMandate ()}
Load it two ways:
# 1) environment — for experiments and local runs
SOMA_PLUGINS= " my_engine:components " soma serve
# 2) entry point — for a distributed package
# [project.entry-points."soma.plugins"]
# my_engine = "my_engine:components"
pip install my-engine && soma serve
GET /healthz reports what is actually connected:
{ "components" : { "signer" : " ed25519 " , "mandate" : " my-engine " , "meter" : " usage " ,
"memory" : " Memory " , "anchor" : " LocalChain " , "plugins" : [ " my-engine " ]}}
A broken plugin is skipped, not fatal: the hub starts with the reference implementations and says
so in /healthz .
Who is who, and why it matters
The registry is a file. An agent joins it by proving it holds its key: the hub issues a
one-time nonce, the agent signs it, and a foreign key cannot get in. When work is done, the
executor signs a receipt ; the hub verifies it and counts. Actions refused by a mandate are
counted as breaches. Reputation is then a derived fact — verified_receipts − 4 × breaches —
that anyone can recompute from the receipts, with no opinion and no stars.
Nodes can share work and history. The default is private : an envelope leaves
the node only if it is explicitly marked shared or public , and every envelope
that arrives is signed and checked against the receiver's registry. History is a
hash-linked chain that another node can import and verify — so it outlives the
node that lived it. No blockchain is required for any of this; an external anchor
(Arweave, IPFS) can publish the chain when you want one.
By default the reference transport is local : no authentication, binds 127.0.0.1 — a
component to put behind your own gateway. For a public deployment there is a managed profile:
SOMA_MANAGED=1 SOMA_TOKEN= < secret > SOMA_QUOTA_PER_HOUR=100 soma serve
Variable
Meaning
SOMA_MANAGED=1
require Authorization: Bearer $SOMA_TOKEN on every request
SOMA_TOKEN
the token; if it is not set, the managed profile denies everything
SOMA_QUOTA_PER_HOUR
max envelopes per sender per hour (0 = no limit)
SOMA_MAX_BODY
max request body in bytes (default 1 MiB)
Without a token configured, the managed profile is closed — a misconfiguration must not open
the door. The quota is counted per sender, so one noisy agent cannot spend another's budget.
The reference HTTP transport has no authentication and binds 127.0.0.1 . Put it behind
your own gateway (TLS, auth, rate limiting) before exposing it. Bodies are capped at 1 MiB.
A signature proves integrity by itself. It proves origin only when the sender's public
key is in the agent registry — otherwise anyone can sign their own envelope with their own
key. The reference implementation checks the registry when it has one.
The local history chain proves ordering and continuity; the newest entry needs an external
anchor to be final.
Memory quality is measured, not asserted. The benchmark runs offline, without an LLM:
bench/memory_bench.py (numbers) and tests/test_recall_abilities.py (abilities).
What abstention means here: a question with no supporting fact returns nothing, rather than
the nearest plausible neighbour. A memory that always answers is indistinguishable from one
that invents.
Beta. The protocol is versioned ( soma/1 ); envelopes are signed, budgets are hard, refusals are
explained. See ROADMAP.md for what is next and SECURITY.md for the trust model.
Soma — протокол коммуникации агентов: конверт с бюджетом, мандатом и подпись

[truncated]
