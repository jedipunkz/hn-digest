---
source: "https://github.com/mnestiq/mnestiq"
hn_url: "https://news.ycombinator.com/item?id=50180226"
title: "Show HN: Mnestiq – tamper-evident logs of what an AI agent saw and did"
article_title: "GitHub - mnestiq/mnestiq: Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections. · GitHub"
image: "https://opengraph.githubassets.com/ea4d971d869b8199320814a95b75c89f3cec87bce5f035bfbba1101145b5d7e3/mnestiq/mnestiq"
author: "lynceus"
captured_at: "2026-10-21T13:30:07Z"
capture_tool: "hn-digest"
hn_id: 50180226
score: 1
comments: 0
posted_at: "2026-10-21T13:01:54Z"
tags:
  - hacker-news
---

# Show HN: Mnestiq – tamper-evident logs of what an AI agent saw and did

- HN: [50180226](https://news.ycombinator.com/item?id=50180226)
- Source: [github.com](https://github.com/mnestiq/mnestiq)
- Score: 1
- Comments: 0
- Posted: 2026-10-21T13:01:54Z

## Translation

Title: Show HN: Mnestiq – tamper-evident logs of what an AI agent saw and did
Article title: GitHub - mnestiq/mnestiq: Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections. · GitHub
Description: Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections. - mnestiq/mnestiq

Article text:
GitHub - mnestiq/mnestiq: Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections. · GitHub
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
mnestiq
mnestiq Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections.
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
33 Commits 33 Commits Folders and files
.github .github docs docs examples examples python python spec spec .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
A flight recorder for AI agents. Tamper-evident, signed evidence of what an agent
saw, decided and did, built for incident response rather than dashboards.
When an agent emails the customer list to an attacker, observability traces tell you
that it happened. Evidence tells you which run and step did it, that the
instruction came from a web page rather than the user, that an auto-approve policy
waved it through, and lets you prove none of that was edited after the fact.
firewall: 10.0.4.17:53120 -> 198.51.100.9:25 (50 bytes) [no agent identity]
mnestiq: agent=support-bot run=c7831ff7 step=6 sandbox=container:7f3a9c operator=acme-helpdesk
untrusted input earlier in the run: step 4: web via fetch_page
authorized by policy:email-auto-approve (policy): recipient not checked against customer domain
Attribution: every event carries run_id , step_id , parent_run_id ,
agent_id , operator_id and sandbox_id (auto-detected for containers and k8s pods).
Egress identity: outbound HTTP gets X-Mnestiq-Run-Id / X-Mnestiq-Egress-Id
headers, and every outbound request and raw TCP connection (SMTP, databases, sockets)
is recorded with its local and remote IP and port: the join key to flow and firewall logs.
Provenance: every piece of model context is tagged at capture time:
system , user , model , tool_output , retrieved_doc , web , agent_msg ,
memory . That's what makes prompt-injection analysis possible afterwards.
Tamper evidence: records are hash-chained (RFC 8785 canonical JSON, SHA-256)
and periodically sealed by Ed25519-signed Merkle checkpoints (RFC 9162 trees).
Edits, deletions, insertions and reordering fail verification.
Redaction that keeps the proof: strip prompts, tool arguments and tool output
before handing evidence to an auditor; the chain still verifies.
Drop-in SDK capture: wraps the Anthropic and OpenAI clients (sync and async,
Chat Completions and Responses APIs). Recording fails open: it never breaks the agent.
No code changes needed: mnestiq proxy records an agent's model calls from outside it,
whatever language or framework it is written in. Point its base URL at the proxy.
pip install " mnestiq[anthropic] " # or mnestiq[openai]
mnestiq keygen --out keys
import anthropic
from mnestiq import FileSink , Recorder
from mnestiq . keys import load_private_key
from mnestiq . egress import instrument_egress
from mnestiq . integrations . anthropic import instrument_anthropic
recorder = Recorder ( FileSink ( "evidence.jsonl" ), agent_id = "support-bot" ,
operator_id = "acme-helpdesk" , signing_key = load_private_key ( "keys/signing.key" ))
@ recorder . tool ( source = "web" ) # output of this tool is untrusted web content
def fetch_page ( url : str ) -> str : ...
client = instrument_anthropic ( anthropic . Anthropic (), recorder )
instrument_egress () # stamp + record outbound connections during runs
with recorder . run ( effective_permissions = { "email" : "send" }):
client . messages . create ( model = "claude-sonnet-5-5" , max_tokens = 1024 , messages = [...])
recorder . close () # writes the final signed checkpoint
mnestiq verify evidence.jsonl --trusted-key keys/signing.pub
mnestiq inspect evidence.jsonl
mnestiq redact evidence.jsonl evidence.redacted.jsonl
A key file the agent can read is fine for trying it out. In production, keep the key where
the agent can't read it: signer=SignerClient.from_env() with mnestiq signer serve running
under another account, or signer=AzureKeyVaultSigner(...) . See docs/signing.md .
Or record from outside the agent (any language)
mnestiq proxy evidence.jsonl --agent support-bot --key keys/signing.key
export OPENAI_BASE_URL=http://127.0.0.1:8787/openai/v1 # and/or
export ANTHROPIC_BASE_URL=http://127.0.0.1:8787/anthropic
The agent runs unchanged. Every model call it makes is written to evidence.jsonl , signed.
What the proxy can and cannot see, and how to deploy it so an agent can't go around it:
docs/proxy.md .
Or try a demo first (no API key)
A scripted support agent reads a poisoned web page and emails customer data out. You then
investigate the evidence it left behind.
git clone https://github.com/mnestiq/mnestiq && cd mnestiq
python -m venv .venv && . .venv/bin/activate # Windows: .venv\Scripts\activate
pip install -e python httpx
python examples/injection_incident.py --out evidence
mnestiq dashboard evidence --trusted-key evidence/incident.pub
Dashboard
mnestiq dashboard evidence/ --trusted-key keys/signing.pub
A local investigation view at the link it prints ( http://127.0.0.1:8765/?key=... , whose key keeps
other accounts on the machine out): runs and sub-runs, a step-by-step
timeline with every piece of model context colored by provenance, findings (below), the
verification status of every record (sealed, unsealed, or tampered), and a connection
finder: paste an ip:port from a firewall or flow log and jump to the agent step that
opened it. It live-updates while agents write.
MNQ-001 is the signature of a hijack: an injected instruction has to put the attacker's
address, account or URL into a tool call. MNQ-004 covers hijacks with no destination, such as
"change the password to ...", "delete file 13" or "cancel order 4100". MNQ-005 covers leaks
through the answer itself, such as an injected image link that carries the customer's email
address. MNQ-006 covers data handed to someone who simply asked for it. Findings describe what was observed, not a
verdict: an agent that visits a link a colleague posted triggers MNQ-001 too. The same
rules are available in Python as mnestiq.findings .
It stays on your machine: it binds to loopback only, rejects non-local Host headers,
renders evidence strictly as text, and its Content-Security-Policy forbids the page from
loading or contacting anything but the local server.
To see it with live data (no API keys needed):
python examples/live_agent.py --out evidence # keeps a scripted agent running
mnestiq dashboard evidence --trusted-key evidence/live.pub
See examples/injection_incident.py for a full
incident: an injected web page leads to exfiltration over a raw TCP connection; a
firewall log entry is traced back to the exact agent step; a tampered approval is
caught; a redacted copy is shared.
The recorder runs inside your agent, so it is built to never be the thing that breaks it.
Path
What
spec/SPEC.md
The evidence format, hashing and verification rules.
spec/schema/
JSON Schema (2020-12) for records.
python/
Recorder, verifier and CLI.
examples/
Runnable demos.
Security model in one paragraph
The format makes evidence tamper-evident after it is written . It does not make
a compromised host tell the truth at write time. On its own, a file can be cut back
to any earlier checkpoint without detection; keep the latest checkpoint somewhere
else with HeadFile and verify with --head to close that. Keep the signing key
out of the agent's reach, in a signing service or Azure Key Vault
( docs/signing.md ), so it can't be copied and used to re-sign history.
A signature only means something if you pin the recorder's public key with
--trusted-key ; otherwise someone who rewrites the whole file can re-sign it. Keys change
only by a signed hand-over. See SPEC.md, sections 2 and 6 , and the
threat model : every attack in it is a test, and you can try your own
on the break-it kit .
v0.2 draft. Done: spec, verifier, Python recorder, Anthropic + OpenAI capture,
egress identity propagation (httpx, requests, raw TCP), local dashboard, signing outside
the agent (signing service, Azure Key Vault) with signed key hand-overs, heartbeats,
RFC 3161 timestamps from public authorities, checked by the verifier, and a model API proxy
that records agents in any language from outside them.
External anchoring of checkpoint heads; S3 Object Lock sink.
Automatic run propagation into executors (today: run.bind ); aiohttp; subprocess egress.
OpenAI Agents SDK and LangGraph capture; streaming output capture.
TypeScript SDK, an MCP proxy, and an OpenTelemetry (GenAI) trace receiver.
Evidence test vectors for other implementations; Merkle inclusion proofs for selective disclosure.
cd python
python -m venv .venv && source .venv/bin/activate # Windows: .venv\Scripts\activate
pip install -e " .[dev,anthropic,openai,azure,timestamps] " httpx requests ruff mypy
pytest
ruff check src tests
mypy src
The test suite never touches the network beyond this machine: tests/_netguard.py
blocks any connection or DNS lookup that isn't loopback, so a test that tried would
fail rather than quietly reach the internet.
Please report vulnerabilities privately; see SECURITY.md .
Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections.
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections. - mnestiq/mnestiq

GitHub - mnestiq/mnestiq: Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections. · GitHub
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
mnestiq
mnestiq Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections.
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
33 Commits 33 Commits Folders and files
.github .github docs docs examples examples python python spec spec .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
A flight recorder for AI agents. Tamper-evident, signed evidence of what an agent
saw, decided and did, built for incident response rather than dashboards.
When an agent emails the customer list to an attacker, observability traces tell you
that it happened. Evidence tells you which run and step did it, that the
instruction came from a web page rather than the user, that an auto-approve policy
waved it through, and lets you prove none of that was edited after the fact.
firewall: 10.0.4.17:53120 -> 198.51.100.9:25 (50 bytes) [no agent identity]
mnestiq: agent=support-bot run=c7831ff7 step=6 sandbox=container:7f3a9c operator=acme-helpdesk
untrusted input earlier in the run: step 4: web via fetch_page
authorized by policy:email-auto-approve (policy): recipient not checked against customer domain
Attribution: every event carries run_id , step_id , parent_run_id ,
agent_id , operator_id and sandbox_id (auto-detected for containers and k8s pods).
Egress identity: outbound HTTP gets X-Mnestiq-Run-Id / X-Mnestiq-Egress-Id
headers, and every outbound request and raw TCP connection (SMTP, databases, sockets)
is recorded with its local and remote IP and port: the join key to flow and firewall logs.
Provenance: every piece of model context is tagged at capture time:
system , user , model , tool_output , retrieved_doc , web , agent_msg ,
memory . That's what makes prompt-injection analysis possible afterwards.
Tamper evidence: records are hash-chained (RFC 8785 canonical JSON, SHA-256)
and periodically sealed by Ed25519-signed Merkle checkpoints (RFC 9162 trees).
Edits, deletions, insertions and reordering fail verification.
Redaction that keeps the proof: strip prompts, tool arguments and tool output
before handing evidence to an auditor; the chain still verifies.
Drop-in SDK capture: wraps the Anthropic and OpenAI clients (sync and async,
Chat Completions and Responses APIs). Recording fails open: it never breaks the agent.
No code changes needed: mnestiq proxy records an agent's model calls from outside it,
whatever language or framework it is written in. Point its base URL at the proxy.
pip install " mnestiq[anthropic] " # or mnestiq[openai]
mnestiq keygen --out keys
import anthropic
from mnestiq import FileSink , Recorder
from mnestiq . keys import load_private_key
from mnestiq . egress import instrument_egress
from mnestiq . integrations . anthropic import instrument_anthropic
recorder = Recorder ( FileSink ( "evidence.jsonl" ), agent_id = "support-bot" ,
operator_id = "acme-helpdesk" , signing_key = load_private_key ( "keys/signing.key" ))
@ recorder . tool ( source = "web" ) # output of this tool is untrusted web content
def fetch_page ( url : str ) -> str : ...
client = instrument_anthropic ( anthropic . Anthropic (), recorder )
instrument_egress () # stamp + record outbound connections during runs
with recorder . run ( effective_permissions = { "email" : "send" }):
client . messages . create ( model = "claude-sonnet-5-5" , max_tokens = 1024 , messages = [...])
recorder . close () # writes the final signed checkpoint
mnestiq verify evidence.jsonl --trusted-key keys/signing.pub
mnestiq inspect evidence.jsonl
mnestiq redact evidence.jsonl evidence.redacted.jsonl
A key file the agent can read is fine for trying it out. In production, keep the key where
the agent can't read it: signer=SignerClient.from_env() with mnestiq signer serve running
under another account, or signer=AzureKeyVaultSigner(...) . See docs/signing.md .
Or record from outside the agent (any language)
mnestiq proxy evidence.jsonl --agent support-bot --key keys/signing.key
export OPENAI_BASE_URL=http://127.0.0.1:8787/openai/v1 # and/or
export ANTHROPIC_BASE_URL=http://127.0.0.1:8787/anthropic
The agent runs unchanged. Every model call it makes is written to evidence.jsonl , signed.
What the proxy can and cannot see, and how to deploy it so an agent can't go around it:
docs/proxy.md .
Or try a demo first (no API key)
A scripted support agent reads a poisoned web page and emails customer data out. You then
investigate the evidence it left behind.
git clone https://github.com/mnestiq/mnestiq && cd mnestiq
python -m venv .venv && . .venv/bin/activate # Windows: .venv\Scripts\activate
pip install -e python httpx
python examples/injection_incident.py --out evidence
mnestiq dashboard evidence --trusted-key evidence/incident.pub
Dashboard
mnestiq dashboard evidence/ --trusted-key keys/signing.pub
A local investigation view at the link it prints ( http://127.0.0.1:8765/?key=... , whose key keeps
other accounts on the machine out): runs and sub-runs, a step-by-step
timeline with every piece of model context colored by provenance, findings (below), the
verification status of every record (sealed, unsealed, or tampered), and a connection
finder: paste an ip:port from a firewall or flow log and jump to the agent step that
opened it. It live-updates while agents write.
MNQ-001 is the signature of a hijack: an injected instruction has to put the attacker's
address, account or URL into a tool call. MNQ-004 covers hijacks with no destination, such as
"change the password to ...", "delete file 13" or "cancel order 4100". MNQ-005 covers leaks
through the answer itself, such as an injected image link that carries the customer's email
address. MNQ-006 covers data handed to someone who simply asked for it. Findings describe what was observed, not a
verdict: an agent that visits a link a colleague posted triggers MNQ-001 too. The same
rules are available in Python as mnestiq.findings .
It stays on your machine: it binds to loopback only, rejects non-local Host headers,
renders evidence strictly as text, and its Content-Security-Policy forbids the page from
loading or contacting anything but the local server.
To see it with live data (no API keys needed):
python examples/live_agent.py --out evidence # keeps a scripted agent running
mnestiq dashboard evidence --trusted-key evidence/live.pub
See examples/injection_incident.py for a full
incident: an injected web page leads to exfiltration over a raw TCP connection; a
firewall log entry is traced back to the exact agent step; a tampered approval is
caught; a redacted copy is shared.
The recorder runs inside your agent, so it is built to never be the thing that breaks it.
Path
What
spec/SPEC.md
The evidence format, hashing and verification rules.
spec/schema/
JSON Schema (2020-12) for records.
python/
Recorder, verifier and CLI.
examples/
Runnable demos.
Security model in one paragraph
The format makes evidence tamper-evident after it is written . It does not make
a compromised host tell the truth at write time. On its own, a file can be cut back
to any earlier checkpoint without detection; keep the latest checkpoint somewhere
else with HeadFile and verify with --head to close that. Keep the signing key
out of the agent's reach, in a signing service or Azure Key Vault
( docs/signing.md ), so it can't be copied and used to re-sign history.
A signature only means something if you pin the recorder's public key with
--trusted-key ; otherwise someone who rewrites the whole file can re-sign it. Keys change
only by a signed hand-over. See SPEC.md, sections 2 and 6 , and the
threat model : every attack in it is a test, and you can try your own
on the break-it kit .
v0.2 draft. Done: spec, verifier, Python recorder, Anthropic + OpenAI capture,
egress identity propagation (httpx, requests, raw TCP), local dashboard, signing outside
the agent (signing service, Azure Key Vault) with signed key hand-overs, heartbeats,
RFC 3161 timestamps from public authorities, checked by the verifier, and a model API proxy
that records agents in any language from outside them.
External anchoring of checkpoint heads; S3 Object Lock sink.
Automatic run propagation into executors (today: run.bind ); aiohttp; subprocess egress.
OpenAI Agents SDK and LangGraph capture; streaming output capture.
TypeScript SDK, an MCP proxy, and an OpenTelemetry (GenAI) trace receiver.
Evidence test vectors for other implementations; Merkle inclusion proofs for selective disclosure.
cd python
python -m venv .venv && source .venv/bin/activate # Windows: .venv\Scripts\activate
pip install -e " .[dev,anthropic,openai,azure,timestamps] " httpx requests ruff mypy
pytest
ruff check src tests
mypy src
The test suite never touches the network beyond this machine: tests/_netguard.py
blocks any connection or DNS lookup that isn't loopback, so a test that tried would
fail rather than quietly reach the internet.
Please report vulnerabilities privately; see SECURITY.md .
Signed, verifiable records of what your AI agents did: model calls, tool calls and network connections.
Readme Apache-2.0 license Security policy
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
