---
source: "https://github.com/agency-project/agency"
hn_url: "https://news.ycombinator.com/item?id=50188214"
title: "Agency: Open-source platform for debugging, profiling, and analyzing AI agents"
article_title: "GitHub - agency-project/agency: Open-source platform for debugging, profiling, and analyzing AI agent execution. · GitHub"
image: "https://opengraph.githubassets.com/e94ce6f5b5bc6cef1302437404a898d55d566b663ed83b9cf24027341c267225/agency-project/agency"
author: "matt_d"
captured_at: "2026-10-21T22:08:40Z"
capture_tool: "hn-digest"
hn_id: 50188214
score: 2
comments: 0
posted_at: "2026-10-21T21:56:31Z"
tags:
  - hacker-news
---

# Agency: Open-source platform for debugging, profiling, and analyzing AI agents

- HN: [50188214](https://news.ycombinator.com/item?id=50188214)
- Source: [github.com](https://github.com/agency-project/agency)
- Score: 2
- Comments: 0
- Posted: 2026-10-21T21:56:31Z

## Translation

Title: Agency: Open-source platform for debugging, profiling, and analyzing AI agents
Article title: GitHub - agency-project/agency: Open-source platform for debugging, profiling, and analyzing AI agent execution. · GitHub
Description: Open-source platform for debugging, profiling, and analyzing AI agent execution. - agency-project/agency

Article text:
GitHub - agency-project/agency: Open-source platform for debugging, profiling, and analyzing AI agent execution. · GitHub
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
agency-project
agency Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
Open-source platform for debugging, profiling, and analyzing AI agent execution.
Contributing Activity Custom properties Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
407 Commits 407 Commits Folders and files
.github/ workflows .github/ workflows agency agency docs docs examples examples scripts scripts tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CONTRIBUTING.md CONTRIBUTING.md README.md README.md live_profiler_demo.py live_profiler_demo.py pyproject.toml pyproject.toml swebench_live.py swebench_live.py swebench_three_agents.py swebench_three_agents.py uv.lock uv.lock View all files Repository files navigation
Live debugging and systems profiling for AI agents.
Getting Started · Documentation · Examples · Contributing
Agency is an open-source platform developed at UC Berkeley Sky Computing Lab for running, debugging, profiling, and controlling AI agents .
As agents tackle longer tasks, invoke tools, and work in parallel, understanding what they're actually doing gets harder. Model traces only tell part of the story. What is an agent waiting for? Where is time being spent? Why is one run slower than another?
Agency connects LLM requests, tool calls, and system activity in one place, giving you the visibility to investigate execution, find bottlenecks, and understand what changed.
We are model and harness agnostic! Run Claude Code, Codex, OpenCode, or Agency's native harness with supported backends including OpenAI, Anthropic, Bedrock, and local models via vLLM. Get started
Live Execution — See what's happening
Think htop + gdb for AI agents .
Watch agents execute in real time. Inspect model requests, tool calls, outputs, failures, and waits across agents. Pause, resume, or stop supported executions directly from the dashboard.
When execution finishes, the same view becomes a navigable history of everything that happened.
Critical Path — Find what's slowing you down
See where wall-clock time actually goes.
Trace model requests, tool calls, processes, and dependencies across agents. Identify the work determining completion time, distinguish parallel execution from waiting, and zoom into bottlenecks.
Use the Critical Path Profiler after execution to investigate what held the run back.
Dig into the measurements behind execution.
Explore LLM latency, token usage, tool durations, errors, retries, CPU, memory, I/O, and optional GPU metrics. Follow linked records from high-level activity down to individual events and processes.
No invented numbers. Missing measurements stay missing.
Compare Runs — Understand what changed
Compare completed runs side by side.
Examine execution trajectories, model and tool activity, latency, resource usage, failures, and outcomes. Find where execution paths diverged and which measurements changed.
Agency compares recorded evidence, without guessing at causes.
Agency provides a sandboxed execution environment, Python API, supported agent harnesses, and integrated telemetry. Your application keeps control of its workflow and orchestration.
Compatibility depends on model capabilities, configuration, and external CLI versions. External harness CLIs must be installed separately.
Experimental: Integration with existing Claude Code and Codex CLI sessions on Linux is experimental. Profiling, interception, and execution controls may be limited. Python-orchestrated workflows currently provide the most complete experience.
See configuration and requirements .
API key for your configured model provider
git clone https://github.com/agency-project/agency.git
cd agency
uv sync --locked --python 3.12
export OPENAI_API_KEY= " your-api-key "
export OPENAI_MODEL= " gpt-6-luna "
uv run python examples/quickstart.py
This runs a native agent in an isolated sandbox. The configured model must support reasoning_effort="none" . Model API usage may incur charges.
First startup requires network access for the sandbox image and dependencies. No GPU, ZFS, or CRIU required.
AGENCY_PROFILE=1 AGENCY_PROFILE_SCOPE=workload \
uv run python examples/quickstart.py --webui
Open localhost:7860 for live execution and localhost:7860/profiler for profiling.
The dashboard binds to 0.0.0.0 by default; don't expose it to untrusted networks without appropriate access controls.
For Python integration, see examples/quickstart.py . You can configure the harness and model independently, define tasks with typed inputs and outputs, and collect execution records for later analysis.
Getting Started — Installation and first run
Configuration — Harnesses, models, and environments
Profiling — Recording and interpreting execution
Examples — Agent workflows and integrations
API Reference — Python interfaces
Architecture — Runtime and system design
Agency is open source. Contributions, bug reports, and feedback are welcome.
See CONTRIBUTING.md to get started.
Open-source platform for debugging, profiling, and analyzing AI agent execution.
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Open-source platform for debugging, profiling, and analyzing AI agent execution. - agency-project/agency

GitHub - agency-project/agency: Open-source platform for debugging, profiling, and analyzing AI agent execution. · GitHub
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
agency-project
agency Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
Open-source platform for debugging, profiling, and analyzing AI agent execution.
Contributing Activity Custom properties Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
407 Commits 407 Commits Folders and files
.github/ workflows .github/ workflows agency agency docs docs examples examples scripts scripts tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CONTRIBUTING.md CONTRIBUTING.md README.md README.md live_profiler_demo.py live_profiler_demo.py pyproject.toml pyproject.toml swebench_live.py swebench_live.py swebench_three_agents.py swebench_three_agents.py uv.lock uv.lock View all files Repository files navigation
Live debugging and systems profiling for AI agents.
Getting Started · Documentation · Examples · Contributing
Agency is an open-source platform developed at UC Berkeley Sky Computing Lab for running, debugging, profiling, and controlling AI agents .
As agents tackle longer tasks, invoke tools, and work in parallel, understanding what they're actually doing gets harder. Model traces only tell part of the story. What is an agent waiting for? Where is time being spent? Why is one run slower than another?
Agency connects LLM requests, tool calls, and system activity in one place, giving you the visibility to investigate execution, find bottlenecks, and understand what changed.
We are model and harness agnostic! Run Claude Code, Codex, OpenCode, or Agency's native harness with supported backends including OpenAI, Anthropic, Bedrock, and local models via vLLM. Get started
Live Execution — See what's happening
Think htop + gdb for AI agents .
Watch agents execute in real time. Inspect model requests, tool calls, outputs, failures, and waits across agents. Pause, resume, or stop supported executions directly from the dashboard.
When execution finishes, the same view becomes a navigable history of everything that happened.
Critical Path — Find what's slowing you down
See where wall-clock time actually goes.
Trace model requests, tool calls, processes, and dependencies across agents. Identify the work determining completion time, distinguish parallel execution from waiting, and zoom into bottlenecks.
Use the Critical Path Profiler after execution to investigate what held the run back.
Dig into the measurements behind execution.
Explore LLM latency, token usage, tool durations, errors, retries, CPU, memory, I/O, and optional GPU metrics. Follow linked records from high-level activity down to individual events and processes.
No invented numbers. Missing measurements stay missing.
Compare Runs — Understand what changed
Compare completed runs side by side.
Examine execution trajectories, model and tool activity, latency, resource usage, failures, and outcomes. Find where execution paths diverged and which measurements changed.
Agency compares recorded evidence, without guessing at causes.
Agency provides a sandboxed execution environment, Python API, supported agent harnesses, and integrated telemetry. Your application keeps control of its workflow and orchestration.
Compatibility depends on model capabilities, configuration, and external CLI versions. External harness CLIs must be installed separately.
Experimental: Integration with existing Claude Code and Codex CLI sessions on Linux is experimental. Profiling, interception, and execution controls may be limited. Python-orchestrated workflows currently provide the most complete experience.
See configuration and requirements .
API key for your configured model provider
git clone https://github.com/agency-project/agency.git
cd agency
uv sync --locked --python 3.12
export OPENAI_API_KEY= " your-api-key "
export OPENAI_MODEL= " gpt-6-luna "
uv run python examples/quickstart.py
This runs a native agent in an isolated sandbox. The configured model must support reasoning_effort="none" . Model API usage may incur charges.
First startup requires network access for the sandbox image and dependencies. No GPU, ZFS, or CRIU required.
AGENCY_PROFILE=1 AGENCY_PROFILE_SCOPE=workload \
uv run python examples/quickstart.py --webui
Open localhost:7860 for live execution and localhost:7860/profiler for profiling.
The dashboard binds to 0.0.0.0 by default; don't expose it to untrusted networks without appropriate access controls.
For Python integration, see examples/quickstart.py . You can configure the harness and model independently, define tasks with typed inputs and outputs, and collect execution records for later analysis.
Getting Started — Installation and first run
Configuration — Harnesses, models, and environments
Profiling — Recording and interpreting execution
Examples — Agent workflows and integrations
API Reference — Python interfaces
Architecture — Runtime and system design
Agency is open source. Contributions, bug reports, and feedback are welcome.
See CONTRIBUTING.md to get started.
Open-source platform for debugging, profiling, and analyzing AI agent execution.
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
