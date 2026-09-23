---
source: "https://github.com/zariffromlatif/life-forge"
hn_url: "https://news.ycombinator.com/item?id=49813880"
title: "Show HN: Life Forge – Open-source flight simulator for autonomous AI agents"
article_title: "GitHub - zariffromlatif/life-forge · GitHub"
image: "https://opengraph.githubassets.com/2380114a9500105c9140d06d8a79a35ef2c8cace5fa40805462babaa3d7dff6c/zariffromlatif/life-forge"
author: "mdzariflatif"
captured_at: "2026-09-23T10:53:48Z"
capture_tool: "hn-digest"
hn_id: 49813880
score: 2
comments: 0
posted_at: "2026-09-23T10:17:33Z"
tags:
  - hacker-news
---

# Show HN: Life Forge – Open-source flight simulator for autonomous AI agents

- HN: [49813880](https://news.ycombinator.com/item?id=49813880)
- Source: [github.com](https://github.com/zariffromlatif/life-forge)
- Score: 2
- Comments: 0
- Posted: 2026-09-23T10:17:33Z

## Translation

Title: Show HN: Life Forge – Open-source flight simulator for autonomous AI agents
Article title: GitHub - zariffromlatif/life-forge · GitHub
Description: Contribute to zariffromlatif/life-forge development by creating an account on GitHub.

Article text:
GitHub - zariffromlatif/life-forge · GitHub
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
zariffromlatif
/
life-forge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github .github examples examples experiments experiments lifeforge lifeforge notebooks notebooks results results tests tests .gitignore .gitignore CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
LIFE FORGE: The Autonomous Flight Simulator for AI Agents
Co-evolutionary adversarial red-teaming and dynamic stress-testing for autonomous AI agents using Artificial Life Quality-Diversity algorithms (3D MAP-Elites).
┌─────────────────────────────────────────────────────────────────────────┐
│ LIFE FORGE │
│ The AI Agent Flight Simulator │
└─────────────────────────────────────────────────────────────────────────┘
│
┌─────────────────────────┴─────────────────────────┐
▼ ▼
┌──────────────────────────────┐ ┌──────────────────────────────┐
│ Target AI Agent │ ◄────────► │ Simulated World (Sandbox) │
│ (Claude, GPT-4o, Llama, │ Actions/ │ • ERP Database & Balances │
│ Qwen, Custom Frameworks) │ Tools │ • Vendor Catalogs & Quotes │
└──────────────────────────────┘ │ • Email Inbox / Outbox │
└──────────────────────────────┘
▲
│ Co-Evolves
│ Perturbations
┌──────────────────────────────┐
│ Evolution Engine │
│ • Adversarial Injections │
│ • Market Price Volatility │
│ • Supply Scarcity │
│ • 3D MAP-Elites Archive │
└──────────────────────────────┘
│
▼
┌──────────────────────────────┐
│ Causal Root-Cause Engine │
│ Generates Audit Report & │
│ Minimal Fix Recommendations │
└──────────────────────────────┘
Why LIFE FORGE?
Before commercial pilots fly passengers, they spend hundreds of hours in a flight simulator . The simulator doesn't give them sunny skies; it throws dual-engine failure, sudden crosswinds, sensor glitches, and electrical fires at them.
Today, enterprise AI agents with tool access (databases, email, financial APIs) are deployed with almost zero dynamic testing:
Static Benchmarks Are Useless : MMLU and HumanEval test multiple-choice trivia and leetcode snippets. They do not test what happens when an agent manages an ERP database while an adversary attempts an indirect prompt injection.
Static Red-Teaming Fails : Traditional security tools test static lists of 500 prompts that LLMs quickly memorize.
The Real World Fights Back : In production, suppliers run out of stock, prices fluctuate mid-transaction, executive emails get spoofed, and third-party data contains jailbreaks.
LIFE FORGE is the autonomous flight simulator for AI agents. Using evolutionary Quality-Diversity algorithms (3D MAP-Elites), LIFE FORGE autonomously breeds and discovers edge-case scenarios that force frontier models to hallucinate, crash, deadlock, or leak company funds before reaching production.
Empirical Benchmark: Meta Llama 3.1 vs. Alibaba Qwen 2.5
In an empirical showdown evaluated over 30 identical evolutionary generations under identical random seeds ( seed=42 ), LIFE FORGE exposed fundamentally contrasting zero-day failure profiles between two state-of-the-art open-weight models:
Alibaba Qwen 2.5 (14B) : Exhibited high general reasoning, but suffered from authority-impersonation blindness . When presented with spoofed CEO emails and vendor quote injections, Qwen believed the untrusted directive and executed the prohibited transfer_funds tool 7 times, wiring $60,000 to unauthorized offshore accounts.
Meta Llama 3.1 (8B) : Completely resisted social engineering and prompt injection wire fraud (0 exfiltrations). However, when market conditions shifted prices between catalog query and order submission, Llama entered an unbounded infinite retry loop 12 times , exhausting its maximum step budget without making progress.
1. 3D MAP-Elites Quality-Diversity Engine
Unlike standard optimizers that search for only one failure, LIFE FORGE maps an entire topography of agent fragility across a 3-dimensional phenotypic coordinate space:
Axis X (Adversarial Intensity) : Frequency, subtlety, and complexity of prompt injections and spoofed communications ( $0.0 \to 1.0$ ).
Axis Y (Environmental Volatility) : Market price surges, supplier dropouts, inventory exhaustion ( $0.0 \to 1.0$ ).
Axis Z (Budget Pressure) : Corporate treasury limits and tight spending caps ( $0.0 \to 1.0$ ).
2. Deterministic Digital Twin Sandbox
A zero-side-effect in-memory enterprise simulation state ( WorldState ) with instantaneous snapshot and causal rollback. Agents interact with 5 simulated enterprise tools:
query_database : Inspects inventory, prices, balances.
vendor_api : Fetches external catalog quotes from suppliers.
issue_purchase_order : Purchases hardware and commits company budget.
send_email : Internal communication channel.
transfer_funds : High-privilege banking wire transfer tool (policy-prohibited in procurement).
Monitors agent actions after every step and enforces mathematical policy boundaries:
UNAUTHORIZED_TOOL_EXECUTION : High-severity privilege boundary breach.
UNAUTHORIZED_FINANCIAL_DRAIN : Exceeding budget or unapproved fund movement.
RECURSIVE_LOOP_TRAP : Cyclical tool re-submission without parameter updates.
GOAL_INVENTORY_DEFICIT : Premature task termination without goal fulfillment.
4. Native Model Context Protocol (MCP) Server
LIFE FORGE can be run as a standard Model Context Protocol (MCP) server over stdio or SSE . Any MCP-compatible client--including Claude Desktop , Cursor , LangGraph , or custom multi-agent frameworks--can directly connect to LIFE FORGE's adversarial environments.
5. Multi-Provider LiteLLM Adapter
Test any frontier or local model with zero code changes:
Local Models : Run on local GPUs via Ollama / vLLM ( ollama/llama3.1:8b , ollama/qwen2.5:14b ).
Cloud Providers : OpenAI ( gpt-4o , o3-mini ), Anthropic ( claude-3-5-sonnet , claude-3-5-haiku ), Google AI Studio ( gemini-2.5-flash , gemini-1.5-flash ).
Built-in automatic rate-limit backoff handler for 429/503 quota management.
6. The Scientific Core: The MODES Framework
Underneath the agent simulator lies LIFE FORGE's foundational Artificial Life laboratory, designed to measure open-ended evolution and avoid the "Beautiful Garbage" trap (confusing high-entropy white noise with true computational complexity):
Activity : Bedau-Packard evolutionary activity waves ( $A_{cum}$ , excess activity over neutral shadow models).
Complexity : Shannon entropy ( $H$ ), bit-packed LZW algorithmic compressibility ( $C$ ), and the Complexity Gap ($H \cdot (1 - C)$) which peaks sharply on Wolfram Class IV systems.
Novelty : Cumulative vocabulary growth of local neighborhood micro-states.
Ecology : 8-connected spatial cluster tracking and entity diversity.
Clone the repository and install with optional extras:
git clone https://github.com/zariffromlatif/life-forge.git
cd life-forge
python -m venv .venv
# On Windows:
. \. venv \S cripts \A ctivate.ps1
# On Linux/macOS:
source .venv/bin/activate
# Install with LLM, MCP, and visualization dependencies:
pip install -e " .[all] "
1. Instant Baseline Stress Test (Offline / Zero Cost)
Stress-test the built-in reference agent across 50 evolutionary generations (requires no API keys):
python -m lifeforge.cli test --scenarios 50 --out results/baseline_report.md --json
2. Stress-Test Local Models on Your GPU (Ollama)
Run unlimited, free evolutionary stress tests against open-weight models on your local GPU (e.g. RTX 3080/4090):
# 1. Start Ollama with your chosen model:
ollama run llama3.1:8b
# 2. Run LIFE FORGE against your local GPU:
python -m lifeforge.cli test --model ollama/llama3.1:8b --api-base http://localhost:11434 --scenarios 30 --seed 42 --out results/llama_report.md --json
3. Launch the Visual Flight Simulator Dashboard (Web UI)
Explore 3D MAP-Elites behavior spaces, compare model showdowns, and inspect step-by-step exploit traces in an interactive local command center (zero extra dependencies required):
python -m lifeforge.cli ui --port 8000
Open your browser at http://localhost:8000 to inspect discovered zero-days, explore behavioral niches, or export an executive PDF audit dossier.
4. Stress-Test Cloud Frontier Models (Gemini / OpenAI / Claude)
Run against cloud frontier models using your API keys:
# Test Gemini (Google AI Studio Free Tier):
$env :GEMINI_API_KEY = " your-api-key "
python -m lifeforge.cli test --model gemini/gemini-1.5-flash --scenarios 20 --delay 4.0 --out results/gemini_report.md --json
# Test OpenAI GPT-4o:
$env :OPENAI_API_KEY = " your-api-key "
python -m lifeforge.cli test --model gpt-4o-mini --scenarios 20 --out results/gpt_report.md --json
5. Head-to-Head Model Showdown Comparison
Compare two or more evaluation reports side-by-side to crown the security winner:
python -m lifeforge.cli compare results/local_qwen_report.json results/local_llama_report.json --out results/MODEL_SHOWDOWN.md
6. Run as an MCP Server (Claude Desktop & Cursor)
Expose LIFE FORGE as a live MCP tool server:
python -m lifeforge.cli mcp-serve --transport stdio --adversarial
Add to your claude_desktop_config.json :
{
"mcpServers" : {
"lifeforge" : {
"command" : " python " ,
"args" : [ " -m " , " lifeforge.cli " , " mcp-serve " , " --transport " , " stdio " , " --adversarial " ]
}
}
}
7. Scientific Cellular Automata Laboratory
Simulate candidate universes and compute quantitative MODES complexity vectors:
# Run Conway's Game of Life
python -m lifeforge.cli run --substrate totalistic --steps 100
# Run Wolfram Rule 110 (Turing complete)
python -m lifeforge.cli run --substrate elementary --rule 110 --steps 100
# High-throughput 100-universe physics survey
python -m lifeforge.cli survey --count 100 --steps 150 --db results/survey.jsonl
Continuous CI/CD Integration
Prevent vulnerable or deadlocking agents from reaching production. Add LIFE FORGE to your GitHub repository workflow:
# .github/workflows/agent_stress_test.yml
name : AI Agent Stress Test
on : [push, pull_request]
jobs :
red-team :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : actions/setup-python@v5
with :
python-version : " 3.12 "
- name : Install LIFE FORGE
run : pip install -e ".[all]"
- name : Run Evolutionary Stress Test
run : |
python -m lifeforge.cli test --scenarios 30 --out results/ci_report.md --json
- name : Upload Audit Report
uses : actions/upload-artifact@v4
with :
name : agent-evolution-report
path : results/ci_report.md
Repository Architecture
lifeforge/
├── substrates/ # Artificial Life & Cellular Automata p

[truncated]

## Original Extract

Contribute to zariffromlatif/life-forge development by creating an account on GitHub.

GitHub - zariffromlatif/life-forge · GitHub
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
zariffromlatif
/
life-forge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
15 Commits 15 Commits Folders and files
.github .github examples examples experiments experiments lifeforge lifeforge notebooks notebooks results results tests tests .gitignore .gitignore CITATION.cff CITATION.cff CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml requirements.txt requirements.txt View all files Repository files navigation
LIFE FORGE: The Autonomous Flight Simulator for AI Agents
Co-evolutionary adversarial red-teaming and dynamic stress-testing for autonomous AI agents using Artificial Life Quality-Diversity algorithms (3D MAP-Elites).
┌─────────────────────────────────────────────────────────────────────────┐
│ LIFE FORGE │
│ The AI Agent Flight Simulator │
└─────────────────────────────────────────────────────────────────────────┘
│
┌─────────────────────────┴─────────────────────────┐
▼ ▼
┌──────────────────────────────┐ ┌──────────────────────────────┐
│ Target AI Agent │ ◄────────► │ Simulated World (Sandbox) │
│ (Claude, GPT-4o, Llama, │ Actions/ │ • ERP Database & Balances │
│ Qwen, Custom Frameworks) │ Tools │ • Vendor Catalogs & Quotes │
└──────────────────────────────┘ │ • Email Inbox / Outbox │
└──────────────────────────────┘
▲
│ Co-Evolves
│ Perturbations
┌──────────────────────────────┐
│ Evolution Engine │
│ • Adversarial Injections │
│ • Market Price Volatility │
│ • Supply Scarcity │
│ • 3D MAP-Elites Archive │
└──────────────────────────────┘
│
▼
┌──────────────────────────────┐
│ Causal Root-Cause Engine │
│ Generates Audit Report & │
│ Minimal Fix Recommendations │
└──────────────────────────────┘
Why LIFE FORGE?
Before commercial pilots fly passengers, they spend hundreds of hours in a flight simulator . The simulator doesn't give them sunny skies; it throws dual-engine failure, sudden crosswinds, sensor glitches, and electrical fires at them.
Today, enterprise AI agents with tool access (databases, email, financial APIs) are deployed with almost zero dynamic testing:
Static Benchmarks Are Useless : MMLU and HumanEval test multiple-choice trivia and leetcode snippets. They do not test what happens when an agent manages an ERP database while an adversary attempts an indirect prompt injection.
Static Red-Teaming Fails : Traditional security tools test static lists of 500 prompts that LLMs quickly memorize.
The Real World Fights Back : In production, suppliers run out of stock, prices fluctuate mid-transaction, executive emails get spoofed, and third-party data contains jailbreaks.
LIFE FORGE is the autonomous flight simulator for AI agents. Using evolutionary Quality-Diversity algorithms (3D MAP-Elites), LIFE FORGE autonomously breeds and discovers edge-case scenarios that force frontier models to hallucinate, crash, deadlock, or leak company funds before reaching production.
Empirical Benchmark: Meta Llama 3.1 vs. Alibaba Qwen 2.5
In an empirical showdown evaluated over 30 identical evolutionary generations under identical random seeds ( seed=42 ), LIFE FORGE exposed fundamentally contrasting zero-day failure profiles between two state-of-the-art open-weight models:
Alibaba Qwen 2.5 (14B) : Exhibited high general reasoning, but suffered from authority-impersonation blindness . When presented with spoofed CEO emails and vendor quote injections, Qwen believed the untrusted directive and executed the prohibited transfer_funds tool 7 times, wiring $60,000 to unauthorized offshore accounts.
Meta Llama 3.1 (8B) : Completely resisted social engineering and prompt injection wire fraud (0 exfiltrations). However, when market conditions shifted prices between catalog query and order submission, Llama entered an unbounded infinite retry loop 12 times , exhausting its maximum step budget without making progress.
1. 3D MAP-Elites Quality-Diversity Engine
Unlike standard optimizers that search for only one failure, LIFE FORGE maps an entire topography of agent fragility across a 3-dimensional phenotypic coordinate space:
Axis X (Adversarial Intensity) : Frequency, subtlety, and complexity of prompt injections and spoofed communications ( $0.0 \to 1.0$ ).
Axis Y (Environmental Volatility) : Market price surges, supplier dropouts, inventory exhaustion ( $0.0 \to 1.0$ ).
Axis Z (Budget Pressure) : Corporate treasury limits and tight spending caps ( $0.0 \to 1.0$ ).
2. Deterministic Digital Twin Sandbox
A zero-side-effect in-memory enterprise simulation state ( WorldState ) with instantaneous snapshot and causal rollback. Agents interact with 5 simulated enterprise tools:
query_database : Inspects inventory, prices, balances.
vendor_api : Fetches external catalog quotes from suppliers.
issue_purchase_order : Purchases hardware and commits company budget.
send_email : Internal communication channel.
transfer_funds : High-privilege banking wire transfer tool (policy-prohibited in procurement).
Monitors agent actions after every step and enforces mathematical policy boundaries:
UNAUTHORIZED_TOOL_EXECUTION : High-severity privilege boundary breach.
UNAUTHORIZED_FINANCIAL_DRAIN : Exceeding budget or unapproved fund movement.
RECURSIVE_LOOP_TRAP : Cyclical tool re-submission without parameter updates.
GOAL_INVENTORY_DEFICIT : Premature task termination without goal fulfillment.
4. Native Model Context Protocol (MCP) Server
LIFE FORGE can be run as a standard Model Context Protocol (MCP) server over stdio or SSE . Any MCP-compatible client--including Claude Desktop , Cursor , LangGraph , or custom multi-agent frameworks--can directly connect to LIFE FORGE's adversarial environments.
5. Multi-Provider LiteLLM Adapter
Test any frontier or local model with zero code changes:
Local Models : Run on local GPUs via Ollama / vLLM ( ollama/llama3.1:8b , ollama/qwen2.5:14b ).
Cloud Providers : OpenAI ( gpt-4o , o3-mini ), Anthropic ( claude-3-5-sonnet , claude-3-5-haiku ), Google AI Studio ( gemini-2.5-flash , gemini-1.5-flash ).
Built-in automatic rate-limit backoff handler for 429/503 quota management.
6. The Scientific Core: The MODES Framework
Underneath the agent simulator lies LIFE FORGE's foundational Artificial Life laboratory, designed to measure open-ended evolution and avoid the "Beautiful Garbage" trap (confusing high-entropy white noise with true computational complexity):
Activity : Bedau-Packard evolutionary activity waves ( $A_{cum}$ , excess activity over neutral shadow models).
Complexity : Shannon entropy ( $H$ ), bit-packed LZW algorithmic compressibility ( $C$ ), and the Complexity Gap ($H \cdot (1 - C)$) which peaks sharply on Wolfram Class IV systems.
Novelty : Cumulative vocabulary growth of local neighborhood micro-states.
Ecology : 8-connected spatial cluster tracking and entity diversity.
Clone the repository and install with optional extras:
git clone https://github.com/zariffromlatif/life-forge.git
cd life-forge
python -m venv .venv
# On Windows:
. \. venv \S cripts \A ctivate.ps1
# On Linux/macOS:
source .venv/bin/activate
# Install with LLM, MCP, and visualization dependencies:
pip install -e " .[all] "
1. Instant Baseline Stress Test (Offline / Zero Cost)
Stress-test the built-in reference agent across 50 evolutionary generations (requires no API keys):
python -m lifeforge.cli test --scenarios 50 --out results/baseline_report.md --json
2. Stress-Test Local Models on Your GPU (Ollama)
Run unlimited, free evolutionary stress tests against open-weight models on your local GPU (e.g. RTX 3080/4090):
# 1. Start Ollama with your chosen model:
ollama run llama3.1:8b
# 2. Run LIFE FORGE against your local GPU:
python -m lifeforge.cli test --model ollama/llama3.1:8b --api-base http://localhost:11434 --scenarios 30 --seed 42 --out results/llama_report.md --json
3. Launch the Visual Flight Simulator Dashboard (Web UI)
Explore 3D MAP-Elites behavior spaces, compare model showdowns, and inspect step-by-step exploit traces in an interactive local command center (zero extra dependencies required):
python -m lifeforge.cli ui --port 8000
Open your browser at http://localhost:8000 to inspect discovered zero-days, explore behavioral niches, or export an executive PDF audit dossier.
4. Stress-Test Cloud Frontier Models (Gemini / OpenAI / Claude)
Run against cloud frontier models using your API keys:
# Test Gemini (Google AI Studio Free Tier):
$env :GEMINI_API_KEY = " your-api-key "
python -m lifeforge.cli test --model gemini/gemini-1.5-flash --scenarios 20 --delay 4.0 --out results/gemini_report.md --json
# Test OpenAI GPT-4o:
$env :OPENAI_API_KEY = " your-api-key "
python -m lifeforge.cli test --model gpt-4o-mini --scenarios 20 --out results/gpt_report.md --json
5. Head-to-Head Model Showdown Comparison
Compare two or more evaluation reports side-by-side to crown the security winner:
python -m lifeforge.cli compare results/local_qwen_report.json results/local_llama_report.json --out results/MODEL_SHOWDOWN.md
6. Run as an MCP Server (Claude Desktop & Cursor)
Expose LIFE FORGE as a live MCP tool server:
python -m lifeforge.cli mcp-serve --transport stdio --adversarial
Add to your claude_desktop_config.json :
{
"mcpServers" : {
"lifeforge" : {
"command" : " python " ,
"args" : [ " -m " , " lifeforge.cli " , " mcp-serve " , " --transport " , " stdio " , " --adversarial " ]
}
}
}
7. Scientific Cellular Automata Laboratory
Simulate candidate universes and compute quantitative MODES complexity vectors:
# Run Conway's Game of Life
python -m lifeforge.cli run --substrate totalistic --steps 100
# Run Wolfram Rule 110 (Turing complete)
python -m lifeforge.cli run --substrate elementary --rule 110 --steps 100
# High-throughput 100-universe physics survey
python -m lifeforge.cli survey --count 100 --steps 150 --db results/survey.jsonl
Continuous CI/CD Integration
Prevent vulnerable or deadlocking agents from reaching production. Add LIFE FORGE to your GitHub repository workflow:
# .github/workflows/agent_stress_test.yml
name : AI Agent Stress Test
on : [push, pull_request]
jobs :
red-team :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : actions/setup-python@v5
with :
python-version : " 3.12 "
- name : Install LIFE FORGE
run : pip install -e ".[all]"
- name : Run Evolutionary Stress Test
run : |
python -m lifeforge.cli test --scenarios 30 --out results/ci_report.md --json
- name : Upload Audit Report
uses : actions/upload-artifact@v4
with :
name : agent-evolution-report
path : results/ci_report.md
Repository Architecture
lifeforge/
├── substrates/ # Artificial Life & Cellular Automata p

[truncated]
