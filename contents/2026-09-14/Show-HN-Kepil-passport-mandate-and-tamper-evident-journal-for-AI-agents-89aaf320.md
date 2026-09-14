---
source: "https://github.com/oleg-vdv/kepil"
hn_url: "https://news.ycombinator.com/item?id=49692366"
title: "Show HN: Kepil – passport, mandate and tamper-evident journal for AI agents"
article_title: "GitHub - oleg-vdv/kepil: Accountability layer for AI agents: passport, mandate, action gateway, tamper-evident journal — and undo · GitHub"
image: "https://opengraph.githubassets.com/0e1704ed7ae0e39472297465b739d8c80fc27bf327623985113c3bcc9032bde9/oleg-vdv/kepil"
author: "ipgleg"
captured_at: "2026-09-14T05:51:47Z"
capture_tool: "hn-digest"
hn_id: 49692366
score: 2
comments: 0
posted_at: "2026-09-14T05:28:36Z"
tags:
  - hacker-news
---

# Show HN: Kepil – passport, mandate and tamper-evident journal for AI agents

- HN: [49692366](https://news.ycombinator.com/item?id=49692366)
- Source: [github.com](https://github.com/oleg-vdv/kepil)
- Score: 2
- Comments: 0
- Posted: 2026-09-14T05:28:36Z

## Translation

Title: Show HN: Kepil – passport, mandate and tamper-evident journal for AI agents
Article title: GitHub - oleg-vdv/kepil: Accountability layer for AI agents: passport, mandate, action gateway, tamper-evident journal — and undo · GitHub
Description: Accountability layer for AI agents: passport, mandate, action gateway, tamper-evident journal — and undo - oleg-vdv/kepil

Article text:
GitHub - oleg-vdv/kepil: Accountability layer for AI agents: passport, mandate, action gateway, tamper-evident journal — and undo · GitHub
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
oleg-vdv
/
kepil
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github/ workflows .github/ workflows docs docs schemas schemas src/ kepil src/ kepil tests tests .gitignore .gitignore LICENSE LICENSE NOTICE.md NOTICE.md README.md README.md README.ru.md README.ru.md pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
Accountability layer for AI agents. Give every agent a passport, put every
action through one gate, and keep a log that cannot be rewritten afterwards.
53% of organisations have had an AI agent exceed its intended permissions.
48% of agents in production run with no monitoring at all. Only 22% treat an
agent as an entity with its own identity.
— Cloud Security Alliance and State of AI Agent Security, 2026
Kepil is what the other 78% are missing: identity, mandate, enforcement,
evidence — and the part nobody else does, undo .
pip install kepil
python -m kepil.admin # http://localhost:7317
Русская версия: README.ru.md
Passport. Every agent version gets an immutable card: who built it, who runs
it, what it does, what it will never do, its risk class, its autonomy class,
its limits, and when its risks are due for review. A new version is a new card;
the old one is kept forever.
Mandate. A machine-readable power of attorney for one job: allowed actions,
allowed systems, spending limits, a validity window, and which action types must
be confirmed by a human. Anything not explicitly allowed is refused.
Gate. The single point through which an agent touches the outside world.
Every action is checked against the mandate before a model is even called.
Fail-closed: any error inside the check means refusal, never a pass.
Journal. Append-only JSONL where every record carries the hash of the one
before it. Editing or deleting a record is detectable — by anyone, using an
independent implementation:
npx proofbyte-agent-trace verify data/journal.jsonl
Undo. The journal is a graph of actions, and every profession declares its
compensating action. Kepil walks that graph backwards and stops honestly at the
first step that cannot be undone. Agent platforms record what happened; this one
puts it back.
Confirmations on your phone. Irreversible actions arrive in Telegram with
two buttons — approve or return — so being accountable does not mean sitting at
a laptop.
Kepil ships an MCP server, so an editor, an assistant or another agent can work
through it — and every action still passes the same gate into the same journal.
{
"mcpServers" : {
"kepil" : { "command" : " python " , "args" : [ " -m " , " kepil.mcp " ] }
}
}
Seven tools: list professions, create an order, run a step, read order status,
see what is waiting for a human, verify the journal, read an agent passport.
One tool is deliberately missing: confirmation. If a model could approve an
irreversible action, the human would drop out of the chain and the whole design
would be pointless. The confirmation card goes to a person — in the panel or in
Telegram — and no MCP client can press it. A test enforces this.
Guard your existing automations
Kepil has a small JSON API, so an n8n workflow, a Make scenario or your own
script can ask permission before acting:
curl -X POST http://localhost:7317/api/check -H " Authorization: Bearer $KEPIL_API_TOKEN " -H " Content-Type: application/json " -d ' {"order_id":"ord-0042","action":"send:message","system":"whatsapp.local"} '
{ "decision" : " await_human " , "allowed" : false , "needs_human" : true ,
"reason" : " необратимое действие: требуется подтверждение человека " }
The answer is recorded in the journal, so later you can show on what grounds the
automation did — or did not do — something. For n8n there is a ready node:
n8n-nodes-kepil .
The API stays off until you set a token (panel → Settings, or
KEPIL_API_TOKEN ). A panel bound to localhost is protected by the binding; a
programmatic interface is not, so it is disabled by default.
An agent here is never fully autonomous
AgentPassport refuses to be constructed with the autonomy class where a human
can no longer cancel a decision. That is a deliberate architectural limit rather
than a missing feature — see
ADR-0002 . The gate enforces the
same rule regardless of what a profession definition claims.
Professions: behaviour as data, not code
An agent's job is a JSON description: ordered steps, boundaries, limits,
irreversible action patterns, rollback rules. Adding a new kind of work means
adding a file — or filling in a form in the panel. The dangerous parts stay in
code and under test.
Five ship with the project: inbound leads, process automation, bookkeeping
documents, AI-adoption audit, public-procurement packages.
An order's journal is a sequence of actions and every profession declares the
compensating action for each, so the panel can walk it backwards: pick a window,
and the pass runs from the last action towards earlier ones, stopping at the
first one that cannot be undone. What will happen is shown before the button
is pressed, naming the step where the pass will stop — an undo promise that
quietly fails is worse than no undo at all. The result is recorded as an
operator's decision, which is why neither the MCP server nor the JSON API can
roll anything back: an agent undoing its own actions would be signing in
somebody else's name.
python -m kepil.admin opens an operator console: orders, professions, agent
passports, a meter (actions, tokens, cost, human time replaced), the compliance
generator, the journal with chain verification and anchoring, and settings.
State is plain JSON files under KEPIL_DATA (default ./data ). No database:
you can open them, read them, and attach them to a dispute.
Documentation requirements differ by country and change faster than code, so the
texts live outside the engine. The neutral pack shipped here follows
international practice (ISO/IEC 42001, record-keeping in the spirit of the EU AI
Act). Jurisdiction packs — for example Kazakhstan's AI Law No. 230-VIII with
order No. 95/НҚ — are dropped into $KEPIL_DATA/packs as files.
Zero dependencies. The core runs on the Python 3.11+ standard library, and
CI fails the build if a third-party import appears. That keeps Kepil
installable inside an air-gapped perimeter, and keeps the supply-chain attack
surface of a tool that sees every action at zero.
Values never enter the journal — only types, counts and hashes.
The verifier is a separate implementation in another language. Proof that
only its own author can check is not proof.
Project
Role
agent-trace
Independent journal verification and evidence packs (MIT)
AI-Gateway
PII and secret masking between your apps and external models
AutoGov
Discovery of shadow automations and the credentials they can reach
Status
Alpha, 115 tests. Interfaces may still change. Nothing here is a legal opinion:
before relying on generated documents, have them reviewed by a lawyer in your
jurisdiction.
AGPL-3.0-or-later. Running a network service built on Kepil obliges you to
release your own source under the same terms — or to take a commercial licence.
See NOTICE.md .
Accountability layer for AI agents: passport, mandate, action gateway, tamper-evident journal — and undo
Readme AGPL-3.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Accountability layer for AI agents: passport, mandate, action gateway, tamper-evident journal — and undo - oleg-vdv/kepil

GitHub - oleg-vdv/kepil: Accountability layer for AI agents: passport, mandate, action gateway, tamper-evident journal — and undo · GitHub
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
oleg-vdv
/
kepil
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github/ workflows .github/ workflows docs docs schemas schemas src/ kepil src/ kepil tests tests .gitignore .gitignore LICENSE LICENSE NOTICE.md NOTICE.md README.md README.md README.ru.md README.ru.md pyproject.toml pyproject.toml server.json server.json View all files Repository files navigation
Accountability layer for AI agents. Give every agent a passport, put every
action through one gate, and keep a log that cannot be rewritten afterwards.
53% of organisations have had an AI agent exceed its intended permissions.
48% of agents in production run with no monitoring at all. Only 22% treat an
agent as an entity with its own identity.
— Cloud Security Alliance and State of AI Agent Security, 2026
Kepil is what the other 78% are missing: identity, mandate, enforcement,
evidence — and the part nobody else does, undo .
pip install kepil
python -m kepil.admin # http://localhost:7317
Русская версия: README.ru.md
Passport. Every agent version gets an immutable card: who built it, who runs
it, what it does, what it will never do, its risk class, its autonomy class,
its limits, and when its risks are due for review. A new version is a new card;
the old one is kept forever.
Mandate. A machine-readable power of attorney for one job: allowed actions,
allowed systems, spending limits, a validity window, and which action types must
be confirmed by a human. Anything not explicitly allowed is refused.
Gate. The single point through which an agent touches the outside world.
Every action is checked against the mandate before a model is even called.
Fail-closed: any error inside the check means refusal, never a pass.
Journal. Append-only JSONL where every record carries the hash of the one
before it. Editing or deleting a record is detectable — by anyone, using an
independent implementation:
npx proofbyte-agent-trace verify data/journal.jsonl
Undo. The journal is a graph of actions, and every profession declares its
compensating action. Kepil walks that graph backwards and stops honestly at the
first step that cannot be undone. Agent platforms record what happened; this one
puts it back.
Confirmations on your phone. Irreversible actions arrive in Telegram with
two buttons — approve or return — so being accountable does not mean sitting at
a laptop.
Kepil ships an MCP server, so an editor, an assistant or another agent can work
through it — and every action still passes the same gate into the same journal.
{
"mcpServers" : {
"kepil" : { "command" : " python " , "args" : [ " -m " , " kepil.mcp " ] }
}
}
Seven tools: list professions, create an order, run a step, read order status,
see what is waiting for a human, verify the journal, read an agent passport.
One tool is deliberately missing: confirmation. If a model could approve an
irreversible action, the human would drop out of the chain and the whole design
would be pointless. The confirmation card goes to a person — in the panel or in
Telegram — and no MCP client can press it. A test enforces this.
Guard your existing automations
Kepil has a small JSON API, so an n8n workflow, a Make scenario or your own
script can ask permission before acting:
curl -X POST http://localhost:7317/api/check -H " Authorization: Bearer $KEPIL_API_TOKEN " -H " Content-Type: application/json " -d ' {"order_id":"ord-0042","action":"send:message","system":"whatsapp.local"} '
{ "decision" : " await_human " , "allowed" : false , "needs_human" : true ,
"reason" : " необратимое действие: требуется подтверждение человека " }
The answer is recorded in the journal, so later you can show on what grounds the
automation did — or did not do — something. For n8n there is a ready node:
n8n-nodes-kepil .
The API stays off until you set a token (panel → Settings, or
KEPIL_API_TOKEN ). A panel bound to localhost is protected by the binding; a
programmatic interface is not, so it is disabled by default.
An agent here is never fully autonomous
AgentPassport refuses to be constructed with the autonomy class where a human
can no longer cancel a decision. That is a deliberate architectural limit rather
than a missing feature — see
ADR-0002 . The gate enforces the
same rule regardless of what a profession definition claims.
Professions: behaviour as data, not code
An agent's job is a JSON description: ordered steps, boundaries, limits,
irreversible action patterns, rollback rules. Adding a new kind of work means
adding a file — or filling in a form in the panel. The dangerous parts stay in
code and under test.
Five ship with the project: inbound leads, process automation, bookkeeping
documents, AI-adoption audit, public-procurement packages.
An order's journal is a sequence of actions and every profession declares the
compensating action for each, so the panel can walk it backwards: pick a window,
and the pass runs from the last action towards earlier ones, stopping at the
first one that cannot be undone. What will happen is shown before the button
is pressed, naming the step where the pass will stop — an undo promise that
quietly fails is worse than no undo at all. The result is recorded as an
operator's decision, which is why neither the MCP server nor the JSON API can
roll anything back: an agent undoing its own actions would be signing in
somebody else's name.
python -m kepil.admin opens an operator console: orders, professions, agent
passports, a meter (actions, tokens, cost, human time replaced), the compliance
generator, the journal with chain verification and anchoring, and settings.
State is plain JSON files under KEPIL_DATA (default ./data ). No database:
you can open them, read them, and attach them to a dispute.
Documentation requirements differ by country and change faster than code, so the
texts live outside the engine. The neutral pack shipped here follows
international practice (ISO/IEC 42001, record-keeping in the spirit of the EU AI
Act). Jurisdiction packs — for example Kazakhstan's AI Law No. 230-VIII with
order No. 95/НҚ — are dropped into $KEPIL_DATA/packs as files.
Zero dependencies. The core runs on the Python 3.11+ standard library, and
CI fails the build if a third-party import appears. That keeps Kepil
installable inside an air-gapped perimeter, and keeps the supply-chain attack
surface of a tool that sees every action at zero.
Values never enter the journal — only types, counts and hashes.
The verifier is a separate implementation in another language. Proof that
only its own author can check is not proof.
Project
Role
agent-trace
Independent journal verification and evidence packs (MIT)
AI-Gateway
PII and secret masking between your apps and external models
AutoGov
Discovery of shadow automations and the credentials they can reach
Status
Alpha, 115 tests. Interfaces may still change. Nothing here is a legal opinion:
before relying on generated documents, have them reviewed by a lawyer in your
jurisdiction.
AGPL-3.0-or-later. Running a network service built on Kepil obliges you to
release your own source under the same terms — or to take a commercial licence.
See NOTICE.md .
Accountability layer for AI agents: passport, mandate, action gateway, tamper-evident journal — and undo
Readme AGPL-3.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
