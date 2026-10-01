---
source: "https://github.com/sriti-ai/sriti-core"
hn_url: "https://news.ycombinator.com/item?id=49921633"
title: "Sriti Core – Local-first LLM router that cascades Ollama to cloud to frontier"
article_title: "GitHub - sriti-ai/sriti-core: Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models. · GitHub"
image: "https://opengraph.githubassets.com/fc401173239d55b6ad6e23f6d9abad79dd0bdc4262124f383d7d6fa4eac54529/sriti-ai/sriti-core"
author: "sudevs"
captured_at: "2026-10-01T14:15:57Z"
capture_tool: "hn-digest"
hn_id: 49921633
score: 2
comments: 0
posted_at: "2026-10-01T13:46:06Z"
tags:
  - hacker-news
---

# Sriti Core – Local-first LLM router that cascades Ollama to cloud to frontier

- HN: [49921633](https://news.ycombinator.com/item?id=49921633)
- Source: [github.com](https://github.com/sriti-ai/sriti-core)
- Score: 2
- Comments: 0
- Posted: 2026-10-01T13:46:06Z

## Translation

Title: Sriti Core – Local-first LLM router that cascades Ollama to cloud to frontier
Article title: GitHub - sriti-ai/sriti-core: Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models. · GitHub
Description: Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models. - sriti-ai/sriti-core

Article text:
GitHub - sriti-ai/sriti-core: Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models. · GitHub
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
sriti-ai
/
sriti-core
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
assets assets sriti sriti tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md __init__.py __init__.py pyproject.toml pyproject.toml pytest.ini pytest.ini requirements.txt requirements.txt View all files Repository files navigation
Intelligent, Reliability-Aware Model Routing & Cascade Engine for LLMs
Key Features •
Architecture •
Quickstart •
How It Works •
Configuration •
API Reference •
Testing •
Troubleshooting •
License
Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models.
Instead of routing every single request to expensive frontier models (like Claude 3.5 Sonnet or GPT-4o), Sriti dynamically categorizes the request, checks an encrypted semantic cache, and cascades through a hierarchy of models:
Tier 3 (Local / Fast): Free, local models running on edge devices or on-prem hardware via Ollama (e.g., Qwen 2.5, Llama 3.2, MiniCPM).
Tier 2 (Balanced Cloud): High-throughput, cost-effective inference providers (e.g., Groq, Fireworks, DeepInfra).
Tier 1 (Frontier): Capable reasoning models invoked only when task complexity warrants it, or when lower tiers fail verification or latency SLOs.
Sriti continuously learns model reliability, evaluates output quality gates, and guarantees strict latency SLOs while dramatically slashing cloud API bills.
🎯 In-Process Task Classifier: Zero-shot task classification using fast ONNX runtime ( fastembed BAAI/bge-small-en-v1.5) mapping prompts across categories (e.g., Code, Extraction, Summarization, QA, Math, Reasoning).
⚡ Semantically Aware Caching: Vector similarity caching backed by Redis/Valkey with task-dependent cosine thresholds. Uncacheable tasks (e.g., non-deterministic reasoning, math) bypass the cache automatically.
🔒 Encrypted Cache at Rest: Optional AES-GCM (256-bit) authenticated encryption for cached prompt and completion payloads.
🗜️ In-Process Prompt Compression: LLMLingua-2 integration to compress lengthy contexts before sending them over the wire, with an adaptive quality gate to ensure semantic integrity.
📊 Dynamic Cascade with Quality Gates: Evaluates lower-tier completions using structural validation (e.g. JSON validation) or semantic quality scoring. If a tier's response fails, it automatically escalates to the next tier.
📈 Online Reliability Profile Learner: Real-time Exponential Moving Average (EMA) tracking of per-model latency, success rate, and error states without needing external analytical databases.
🧠 Case-Based Memory (KNN Routing): Remembers past routing decisions and dynamically adjusts candidate selection bias based on historical outcome rewards.
🛡️ Lightweight & Self-Contained: Uses a native NumPy cosine similarity vector store requiring no proprietary Redis modules (no RediSearch or RedisJSON needed). Works seamlessly on vanilla Redis or Valkey.
🔌 Unified LLM Execution: Powered by LiteLLM to seamlessly talk to 100+ LLM providers and local inference runtimes.
User / Application Request
│
▼
┌──────────────────────────────────┐
│ 1. Fast ONNX Classifier │ ──> Task Category & Embeddings
└──────────────────────────────────┘
│
▼
┌──────────────────────────────────┐
│ 2. Semantic Cache (Valkey) │ ──[Hit]──> Return Cached Response
└──────────────────────────────────┘
│ [Miss]
▼
┌──────────────────────────────────┐
│ 3. Model Cascading Engine │
│ │
│ ┌────────────────────────────┐ │
│ │ Tier 3: Local Edge Models │ │ ──[Success]──> Accept & Learn
│ └────────────────────────────┘ │
│ │ [Fail/Timeout] │
│ ┌────────────────────────────┐ │
│ │ Tier 2: Cost-Effective LLM │ │ ──[Success]──> Accept & Learn
│ └────────────────────────────┘ │
│ │ [Fail/Timeout] │
│ ┌────────────────────────────┐ │
│ │ Tier 1: Frontier Models │ │ ──[Guaranteed Execution]
│ └────────────────────────────┘ │
└──────────────────────────────────┘
│
▼
┌──────────────────────────────────┐
│ 4. Update Reliability & Cache │
└──────────────────────────────────┘
Quickstart
pip install sriti-core
From source:
git clone https://github.com/sriti-ai/sriti-core.git
cd sriti-core
pip install -e " .[dev] "
2. Configure Environment
Create a .env file with your credentials and configuration:
SRITI_REDIS_URL = redis://localhost:6379/0
# Optional: Provider API keys (only configure what you use)
ANTHROPIC_API_KEY = sk-ant-...
OPENAI_API_KEY = sk-...
GROQ_API_KEY = gsk_...
# Optional: AES-256 GCM key for cache encryption (base64-encoded 32 bytes)
# CACHE_ENCRYPTION_KEY=...
3. Start the Sriti Server
uvicorn sriti.app:app --host 0.0.0.0 --port 8100
4. Send a Request via HTTP
curl -X POST http://localhost:8100/complete \
-H " Content-Type: application/json " \
-d ' {
"prompt": "Summarise the following paragraph in one sentence: ...",
"minimum_tier": "tier_3",
"cacheable": true,
"requires_json": false
} '
5. Python Client SDK
You can also use the bundled async Python client:
import asyncio
from sriti . client import SritiModelClient
async def main ():
client = SritiModelClient ( base_url = "http://localhost:8100" )
response = await client . complete (
prompt = "Explain how semantic caching reduces LLM inference costs." ,
minimum_tier = "tier_3" ,
cacheable = True
)
print ( f"Response: { response . text } " )
print ( f"Tier Used: { response . tier_used } " )
print ( f"Latency: { response . latency_ms } ms" )
print ( f"Cost: $ { response . cost_usd } " )
asyncio . run ( main ())
How It Works
Prompts are embedded using fastembed and compared against predefined semantic task anchors (or an optional lightweight classification head). The classifier runs completely locally in under 10ms.
Tasks have configurable similarity thresholds in policy.yaml . For example, conversational questions allow a similarity threshold of 0.75 , while structured data extraction requires 0.92 .
Vectors and encrypted payloads are stored in Valkey/Redis.
3. Reliability-Driven Cascading
For each tier, active candidates in models.yaml are evaluated based on their Pareto score: a composite metric factoring in latency, cost, and historical success probability.
If a lower tier experiences an API error, context window exhaustion, or fails a verification check (e.g. invalid JSON syntax), Sriti escalates seamlessly to the next tier.
Sriti's configuration is managed through two clean YAML files in sriti/config/ :
models.yaml : Defines model providers, context windows, cost per token, and tier mappings ( default_tier: 1 | 2 | 3 ).
policy.yaml : Defines cache TTLs, per-task similarity thresholds, latency SLOs, and quality gate cutoffs.
The test suite lives in the source repository and is not included in the published wheel. Clone the repo first, then run:
git clone https://github.com/sriti-ai/sriti-core.git
cd sriti-core
pip install -e " .[dev] "
python -m pytest
All 39 unit tests run without requiring live LLM API keys or a running Redis instance (mocked in-memory).
The uvicorn binary is not on your PATH. Run it via Python instead:
python -m uvicorn sriti.app:app --host 0.0.0.0 --port 8100
ModuleNotFoundError: No module named 'fastapi' (or other missing modules)
You installed sriti-core but the server dependencies are not in your active environment. Install them:
pip install sriti-core[dev]
# or from source:
pip install -e " .[dev] "
Could not find a version that satisfies the requirement sriti-core (from versions: none)
Your Python version is below 3.11. Verify:
python --version
sriti-core requires Python 3.11+. If your default Python is older (e.g. Anaconda 3.8), create a compatible environment:
conda create -n sriti-env python=3.11 -y
conda activate sriti-env
pip install sriti-core
pip install -e ".[dev]" fails with setup.py not found
Your pip version is too old to support pyproject.toml -based editable installs. Upgrade pip first:
pip install --upgrade pip
pip install -e " .[dev] "
pytest collects 0 items or fails with ModuleNotFoundError
Your shell is picking up the system/Anaconda pytest instead of the one in your active environment. Use:
python -m pytest
macOS: wrong Python is used even after conda activate
If which python still points to Anaconda's base Python, your shell may not have conda initialized. Run:
source /opt/anaconda3/etc/profile.d/conda.sh
conda activate sriti-env
which python # should now show the sriti-env path
License
This project is licensed under the Apache License 2.0.
Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models.
Contributing Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models. - sriti-ai/sriti-core

GitHub - sriti-ai/sriti-core: Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models. · GitHub
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
sriti-ai
/
sriti-core
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
6 Commits 6 Commits Folders and files
assets assets sriti sriti tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MANIFEST.in MANIFEST.in README.md README.md __init__.py __init__.py pyproject.toml pyproject.toml pytest.ini pytest.ini requirements.txt requirements.txt View all files Repository files navigation
Intelligent, Reliability-Aware Model Routing & Cascade Engine for LLMs
Key Features •
Architecture •
Quickstart •
How It Works •
Configuration •
API Reference •
Testing •
Troubleshooting •
License
Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models.
Instead of routing every single request to expensive frontier models (like Claude 3.5 Sonnet or GPT-4o), Sriti dynamically categorizes the request, checks an encrypted semantic cache, and cascades through a hierarchy of models:
Tier 3 (Local / Fast): Free, local models running on edge devices or on-prem hardware via Ollama (e.g., Qwen 2.5, Llama 3.2, MiniCPM).
Tier 2 (Balanced Cloud): High-throughput, cost-effective inference providers (e.g., Groq, Fireworks, DeepInfra).
Tier 1 (Frontier): Capable reasoning models invoked only when task complexity warrants it, or when lower tiers fail verification or latency SLOs.
Sriti continuously learns model reliability, evaluates output quality gates, and guarantees strict latency SLOs while dramatically slashing cloud API bills.
🎯 In-Process Task Classifier: Zero-shot task classification using fast ONNX runtime ( fastembed BAAI/bge-small-en-v1.5) mapping prompts across categories (e.g., Code, Extraction, Summarization, QA, Math, Reasoning).
⚡ Semantically Aware Caching: Vector similarity caching backed by Redis/Valkey with task-dependent cosine thresholds. Uncacheable tasks (e.g., non-deterministic reasoning, math) bypass the cache automatically.
🔒 Encrypted Cache at Rest: Optional AES-GCM (256-bit) authenticated encryption for cached prompt and completion payloads.
🗜️ In-Process Prompt Compression: LLMLingua-2 integration to compress lengthy contexts before sending them over the wire, with an adaptive quality gate to ensure semantic integrity.
📊 Dynamic Cascade with Quality Gates: Evaluates lower-tier completions using structural validation (e.g. JSON validation) or semantic quality scoring. If a tier's response fails, it automatically escalates to the next tier.
📈 Online Reliability Profile Learner: Real-time Exponential Moving Average (EMA) tracking of per-model latency, success rate, and error states without needing external analytical databases.
🧠 Case-Based Memory (KNN Routing): Remembers past routing decisions and dynamically adjusts candidate selection bias based on historical outcome rewards.
🛡️ Lightweight & Self-Contained: Uses a native NumPy cosine similarity vector store requiring no proprietary Redis modules (no RediSearch or RedisJSON needed). Works seamlessly on vanilla Redis or Valkey.
🔌 Unified LLM Execution: Powered by LiteLLM to seamlessly talk to 100+ LLM providers and local inference runtimes.
User / Application Request
│
▼
┌──────────────────────────────────┐
│ 1. Fast ONNX Classifier │ ──> Task Category & Embeddings
└──────────────────────────────────┘
│
▼
┌──────────────────────────────────┐
│ 2. Semantic Cache (Valkey) │ ──[Hit]──> Return Cached Response
└──────────────────────────────────┘
│ [Miss]
▼
┌──────────────────────────────────┐
│ 3. Model Cascading Engine │
│ │
│ ┌────────────────────────────┐ │
│ │ Tier 3: Local Edge Models │ │ ──[Success]──> Accept & Learn
│ └────────────────────────────┘ │
│ │ [Fail/Timeout] │
│ ┌────────────────────────────┐ │
│ │ Tier 2: Cost-Effective LLM │ │ ──[Success]──> Accept & Learn
│ └────────────────────────────┘ │
│ │ [Fail/Timeout] │
│ ┌────────────────────────────┐ │
│ │ Tier 1: Frontier Models │ │ ──[Guaranteed Execution]
│ └────────────────────────────┘ │
└──────────────────────────────────┘
│
▼
┌──────────────────────────────────┐
│ 4. Update Reliability & Cache │
└──────────────────────────────────┘
Quickstart
pip install sriti-core
From source:
git clone https://github.com/sriti-ai/sriti-core.git
cd sriti-core
pip install -e " .[dev] "
2. Configure Environment
Create a .env file with your credentials and configuration:
SRITI_REDIS_URL = redis://localhost:6379/0
# Optional: Provider API keys (only configure what you use)
ANTHROPIC_API_KEY = sk-ant-...
OPENAI_API_KEY = sk-...
GROQ_API_KEY = gsk_...
# Optional: AES-256 GCM key for cache encryption (base64-encoded 32 bytes)
# CACHE_ENCRYPTION_KEY=...
3. Start the Sriti Server
uvicorn sriti.app:app --host 0.0.0.0 --port 8100
4. Send a Request via HTTP
curl -X POST http://localhost:8100/complete \
-H " Content-Type: application/json " \
-d ' {
"prompt": "Summarise the following paragraph in one sentence: ...",
"minimum_tier": "tier_3",
"cacheable": true,
"requires_json": false
} '
5. Python Client SDK
You can also use the bundled async Python client:
import asyncio
from sriti . client import SritiModelClient
async def main ():
client = SritiModelClient ( base_url = "http://localhost:8100" )
response = await client . complete (
prompt = "Explain how semantic caching reduces LLM inference costs." ,
minimum_tier = "tier_3" ,
cacheable = True
)
print ( f"Response: { response . text } " )
print ( f"Tier Used: { response . tier_used } " )
print ( f"Latency: { response . latency_ms } ms" )
print ( f"Cost: $ { response . cost_usd } " )
asyncio . run ( main ())
How It Works
Prompts are embedded using fastembed and compared against predefined semantic task anchors (or an optional lightweight classification head). The classifier runs completely locally in under 10ms.
Tasks have configurable similarity thresholds in policy.yaml . For example, conversational questions allow a similarity threshold of 0.75 , while structured data extraction requires 0.92 .
Vectors and encrypted payloads are stored in Valkey/Redis.
3. Reliability-Driven Cascading
For each tier, active candidates in models.yaml are evaluated based on their Pareto score: a composite metric factoring in latency, cost, and historical success probability.
If a lower tier experiences an API error, context window exhaustion, or fails a verification check (e.g. invalid JSON syntax), Sriti escalates seamlessly to the next tier.
Sriti's configuration is managed through two clean YAML files in sriti/config/ :
models.yaml : Defines model providers, context windows, cost per token, and tier mappings ( default_tier: 1 | 2 | 3 ).
policy.yaml : Defines cache TTLs, per-task similarity thresholds, latency SLOs, and quality gate cutoffs.
The test suite lives in the source repository and is not included in the published wheel. Clone the repo first, then run:
git clone https://github.com/sriti-ai/sriti-core.git
cd sriti-core
pip install -e " .[dev] "
python -m pytest
All 39 unit tests run without requiring live LLM API keys or a running Redis instance (mocked in-memory).
The uvicorn binary is not on your PATH. Run it via Python instead:
python -m uvicorn sriti.app:app --host 0.0.0.0 --port 8100
ModuleNotFoundError: No module named 'fastapi' (or other missing modules)
You installed sriti-core but the server dependencies are not in your active environment. Install them:
pip install sriti-core[dev]
# or from source:
pip install -e " .[dev] "
Could not find a version that satisfies the requirement sriti-core (from versions: none)
Your Python version is below 3.11. Verify:
python --version
sriti-core requires Python 3.11+. If your default Python is older (e.g. Anaconda 3.8), create a compatible environment:
conda create -n sriti-env python=3.11 -y
conda activate sriti-env
pip install sriti-core
pip install -e ".[dev]" fails with setup.py not found
Your pip version is too old to support pyproject.toml -based editable installs. Upgrade pip first:
pip install --upgrade pip
pip install -e " .[dev] "
pytest collects 0 items or fails with ModuleNotFoundError
Your shell is picking up the system/Anaconda pytest instead of the one in your active environment. Use:
python -m pytest
macOS: wrong Python is used even after conda activate
If which python still points to Anaconda's base Python, your shell may not have conda initialized. Run:
source /opt/anaconda3/etc/profile.d/conda.sh
conda activate sriti-env
which python # should now show the sriti-env path
License
This project is licensed under the Apache License 2.0.
Sriti Core is an open-source, ultra-efficient intelligent routing proxy and cascading engine for Large Language Models.
Contributing Activity Custom properties Stars
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
