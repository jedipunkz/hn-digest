---
source: "https://github.com/beneadie/deep_dog_2"
hn_url: "https://news.ycombinator.com/item?id=49686186"
title: "DeepDog 2: Outperforms Gemini and OpenAI deep research at a fraction of the cost"
article_title: "GitHub - beneadie/deep_dog_2: pacakage version of deep dog 2 research agent · GitHub"
image: "https://opengraph.githubassets.com/954b1963c61056a9740aeb473418f77283cf5f85360fd84801379ddbd082220d/beneadie/deep_dog_2"
author: "beneadie01"
captured_at: "2026-09-13T17:56:17Z"
capture_tool: "hn-digest"
hn_id: 49686186
score: 2
comments: 1
posted_at: "2026-09-13T17:14:23Z"
tags:
  - hacker-news
---

# DeepDog 2: Outperforms Gemini and OpenAI deep research at a fraction of the cost

- HN: [49686186](https://news.ycombinator.com/item?id=49686186)
- Source: [github.com](https://github.com/beneadie/deep_dog_2)
- Score: 2
- Comments: 1
- Posted: 2026-09-13T17:14:23Z

## Translation

Title: DeepDog 2: Outperforms Gemini and OpenAI deep research at a fraction of the cost
Article title: GitHub - beneadie/deep_dog_2: pacakage version of deep dog 2 research agent · GitHub
Description: pacakage version of deep dog 2 research agent. Contribute to beneadie/deep_dog_2 development by creating an account on GitHub.

Article text:
GitHub - beneadie/deep_dog_2: pacakage version of deep dog 2 research agent · GitHub
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
beneadie
/
deep_dog_2
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
assets assets deep_research deep_research scripts scripts tests tests .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md compress_demo.ps1 compress_demo.ps1 deepdog2_demo_web_hd.mp4 deepdog2_demo_web_hd.mp4 pyproject.toml pyproject.toml quickstart.py quickstart.py requirements.txt requirements.txt run_research.py run_research.py View all files Repository files navigation
Turn a question into a cited Markdown report. A supervisor plans the research, delegates work to platform agents, reviews their findings, and writes the final report.
Deep Dog 2 installs directly from GitHub as a Python package in your existing project. Call run_research() with just a question to use the defaults, or pass a per-run configuration to choose models, search, agent types and research budgets. See the quickstart for pip and uv installation, or clone the source to edit the engine and run the repository examples.
At the time of publication, Deep Dog 2 ranked 5th overall and 1st among open-source research agents on the DeepResearch Bench . The published run used a relatively economical profile: a 15-minute research window, a 20-iteration supervisor cap, at most 3 Exa searches per sub-agent, DeepSeek V4 Pro as supervisor, and DeepSeek V4 Flash for sub-agents.
Role-specific models — use different models for the supervisor, initial draft, and platform sub-agents.
Reflection and delegation — the supervisor plans research, delegates distinct questions, evaluates findings, and decides whether more work is needed.
Platform specialization — Web is enabled by default; Reddit, Substack, General, PubMed, Arxiv, and SEC agents can be enabled through configuration.
Evidence-first reports — findings are collected into a source registry, then final inline citations and the ## Sources section are validated before the report is returned.
Operational control — time limits, iteration caps, search budgets, fallback chains, and output modes can be tuned for local experiments or more economical deployments.
Supervisor + sub-agent architecture — configurable platform agents in config.py ; ResearchWeb is enabled by default, with Reddit, Substack, General, PubMed, Arxiv and SEC agents available through configuration
Multiple search providers — Tavily and/or Exa ( WEB_SEARCH_ENGINE at deep_research/config.py:658 )
Provider-agnostic models — DeepSeek, MiMo, Meta Muse, Gemini, OpenAI, GLM, and OpenRouter-hosted models via a single get_model() factory ( deep_research/config.py:883 )
Model fallback chains — per-role fallback lists ( SUBAGENT_MODEL_FALLBACK_CHAIN , SUPERVISOR_MODEL_FALLBACK_CHAIN , DRAFT_REPORT_MODEL_FALLBACK_CHAIN )
Cited reports — Markdown report, source metadata and optional research trace; applications can save the returned results
LangGraph execution — recursion limit, timeouts, and observability logging
Deep Dog 2 retains the draft-first, iterative refinement idea from Deep Dog 1, while making reflection, delegation, and source handling explicit:
User question
|
v
clarify_with_user → write_research_brief → write_draft_report
|
v
supervisor research loop
┌───────────────┼────────────────┐
│ │ │
reflect delegate conclude research
(think_tool) (parallel agents)
|
v
Web / Reddit / Substack / ...
|
v
findings + sources
|
└── repeat until complete
|
v
final report → citation validation → output
Scope the question. The input is converted into a structured research brief.
Create a scaffold. An initial draft establishes a useful report structure before live research begins. It is not treated as evidence.
Reflect and delegate. The supervisor uses internal reflection to identify gaps and delegates focused, non-overlapping research tasks to platform agents.
Research in parallel. Agents search, read, save, and compress findings using the tools available for their platform. Their results are returned with source metadata and citations.
Evaluate. The supervisor reviews the findings and may request another round.
Finalize the report. The final writer combines the brief, draft, and findings. Citation checks validate the relationship between inline citations and the final sources list.
Supervisor reflection is used for planning and control; it is not copied into the final research report. Optional subtopic evaluation and parallel subtopic reports can run after the main report when enabled in deep_research/config.py .
The supported prompt family is OPEN . Older configurations using LEGACY must switch to OPEN .
At the time of publication, Deep Dog 2 ranked 5th overall and 1st among open-source research agents on the DeepResearch Bench benchmark. These results were achieved with a relatively economical configuration: a maximum of 3 Exa searches per sub-agent, 15 minutes of research time, 20 total iterations, and DeepSeek V4 Pro and DeepSeek V4 Flash as the supervisor and sub-agent base models.
These results are a historical reproducibility profile, not a promise about current defaults. The current package defaults to DeepSeek V4 Flash for all model roles and uses a different supervisor iteration default. Benchmark rankings and scores may change as the leaderboard changes.
The following is a pricing-based estimate for one research task. It is not a controlled cost benchmark: the systems use different architectures, search providers and token budgets.
Deep Dog 2’s range is an estimate based on the DeepSeek V4 pricing (off-peak) and Exa pricing , using off-peak DeepSeek rates. Actual cost varies with prompt length, model output, cache hits, the number of delegated agents and how quickly the supervisor concludes. Exa currently advertises $20 in sign-up credits and $10 in credits each month ; Tavily provides 1,000 free credits per month , equivalent to $8 at its $0.008 per-credit pay-as-you-go rate.
Costs are easy to change by changing the configuration: shorten the research window, lower the supervisor or sub-agent iteration limits, reduce searches per sub-agent, or use fewer agents. Nemotron 3.5 Lightning has also worked successfully as a sub-agent model in Deep Dog 2 and costs approximately half as much as DeepSeek V4 Flash in the tested setup. Google’s Gemini figures are its own published estimates; see the Gemini Deep Research pricing section for the assumptions behind them.
For a detailed explanation of the reflection and delegation methods used, see the engineering article Deep Dog 2: How Reflection and Structured Delegation Improve Supervisor–Subagent Research Systems .
Requires Python 3.11+ , Git , a DeepSeek API key and an Exa API key. In your own project directory, use your existing Python environment or create one below.
Choose standard Python tooling:
# Standard Python tooling
python -m venv .venv
Or, using uv :
uv venv --python 3.11 --seed
--seed includes pip so either installer below works in this environment.
# macOS / Linux
source .venv/bin/activate
# Windows PowerShell
.venv\Scripts\Activate.ps1
Install the package directly from GitHub. Choose pip or uv:
# pip
python -m pip install " git+https://github.com/beneadie/deep_dog_2.git "
Or, with uv:
uv pip install " git+https://github.com/beneadie/deep_dog_2.git "
The installer fetches the repository, builds the package and installs its dependencies. You do not need a local source checkout or a PyPI release. This works because the repository includes Python packaging metadata in pyproject.toml ; a Git repository needs a Python package build configuration to support this kind of install. The installed distribution is named deep-dog-2 ; the Python import is deep_research .
Add your keys to a .env file in your project directory:
DEEPSEEK_API_KEY = your-deepseek-key
EXA_API_KEY = your-exa-key
Keep .env out of version control by adding it to your project's .gitignore . If those keys are already set as environment variables, no .env file is needed; existing environment variables take precedence.
Create simple_research.py in your project with the following code. The two settings near the top are optional examples: change the model or maximum research time, and leave the rest of the configuration at its defaults:
import asyncio
from dotenv import load_dotenv
load_dotenv ()
from deep_research . integration import RunConfig , run_research
MODEL = "deepseek-v4-flash"
MAX_MINUTES = 10
config = RunConfig (
supervisor_model_fallback_chain = [ MODEL ],
subagent_model_fallback_chain = [ MODEL ],
draft_report_model_fallback_chain = [ MODEL ],
research_time_max_minutes = MAX_MINUTES ,
)
result = asyncio . run (
run_research (
"How is geothermal energy developing in Europe?" ,
config = config ,
)
)
print ( result . status )
print ( result . final_report )
Run it with your project environment activated:
python simple_research.py
The example prints live progress, the run status and the Markdown report. It does not save a report file by default. Only DeepSeek and Exa credentials are needed here: the example uses DeepSeek V4 Flash for all roles and Exa for Web research. The maximum research window is set to 10 minutes above; other settings remain at their library defaults. Final writing can finish after the research window.
Updating the installed package
An installation uses a snapshot of the repository; it does not update automatically when new commits are published. To fetch and reinstall the current default-branch version, use the matching installer in your project environment:
# pip
python -m pip install --upgrade --force-reinstall " git+https://github.com/beneadie/deep_dog_2.git "
Or, with uv:
uv pip install --upgrade --reinstall-package deep-dog-2 " git+https://github.com/beneadie/deep_dog_2.git "
Reinstallation also picks up code changes that keep the same package version number. These commands can update dependencies too. For a reproducible deployment, append @<full-commit-hash> to the Git URL to install a specific revision. See pip's Git installation documentation and uv's package installation documentation .
To modify the engine or use the repository's ready-made scripts, clone the source:
git clone https://github.com/beneadie/deep_dog_2.git
cd deep_dog_2
Create and activate an environment as above, then choose an editable install:
# pip
python -m pip install -e .
Or, with uv:
uv pip install -e .
The . selects the current checkout; -e makes local source edits available without reinstalling. To update a checkout, pull the upstre

[truncated]

## Original Extract

pacakage version of deep dog 2 research agent. Contribute to beneadie/deep_dog_2 development by creating an account on GitHub.

GitHub - beneadie/deep_dog_2: pacakage version of deep dog 2 research agent · GitHub
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
beneadie
/
deep_dog_2
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
assets assets deep_research deep_research scripts scripts tests tests .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md compress_demo.ps1 compress_demo.ps1 deepdog2_demo_web_hd.mp4 deepdog2_demo_web_hd.mp4 pyproject.toml pyproject.toml quickstart.py quickstart.py requirements.txt requirements.txt run_research.py run_research.py View all files Repository files navigation
Turn a question into a cited Markdown report. A supervisor plans the research, delegates work to platform agents, reviews their findings, and writes the final report.
Deep Dog 2 installs directly from GitHub as a Python package in your existing project. Call run_research() with just a question to use the defaults, or pass a per-run configuration to choose models, search, agent types and research budgets. See the quickstart for pip and uv installation, or clone the source to edit the engine and run the repository examples.
At the time of publication, Deep Dog 2 ranked 5th overall and 1st among open-source research agents on the DeepResearch Bench . The published run used a relatively economical profile: a 15-minute research window, a 20-iteration supervisor cap, at most 3 Exa searches per sub-agent, DeepSeek V4 Pro as supervisor, and DeepSeek V4 Flash for sub-agents.
Role-specific models — use different models for the supervisor, initial draft, and platform sub-agents.
Reflection and delegation — the supervisor plans research, delegates distinct questions, evaluates findings, and decides whether more work is needed.
Platform specialization — Web is enabled by default; Reddit, Substack, General, PubMed, Arxiv, and SEC agents can be enabled through configuration.
Evidence-first reports — findings are collected into a source registry, then final inline citations and the ## Sources section are validated before the report is returned.
Operational control — time limits, iteration caps, search budgets, fallback chains, and output modes can be tuned for local experiments or more economical deployments.
Supervisor + sub-agent architecture — configurable platform agents in config.py ; ResearchWeb is enabled by default, with Reddit, Substack, General, PubMed, Arxiv and SEC agents available through configuration
Multiple search providers — Tavily and/or Exa ( WEB_SEARCH_ENGINE at deep_research/config.py:658 )
Provider-agnostic models — DeepSeek, MiMo, Meta Muse, Gemini, OpenAI, GLM, and OpenRouter-hosted models via a single get_model() factory ( deep_research/config.py:883 )
Model fallback chains — per-role fallback lists ( SUBAGENT_MODEL_FALLBACK_CHAIN , SUPERVISOR_MODEL_FALLBACK_CHAIN , DRAFT_REPORT_MODEL_FALLBACK_CHAIN )
Cited reports — Markdown report, source metadata and optional research trace; applications can save the returned results
LangGraph execution — recursion limit, timeouts, and observability logging
Deep Dog 2 retains the draft-first, iterative refinement idea from Deep Dog 1, while making reflection, delegation, and source handling explicit:
User question
|
v
clarify_with_user → write_research_brief → write_draft_report
|
v
supervisor research loop
┌───────────────┼────────────────┐
│ │ │
reflect delegate conclude research
(think_tool) (parallel agents)
|
v
Web / Reddit / Substack / ...
|
v
findings + sources
|
└── repeat until complete
|
v
final report → citation validation → output
Scope the question. The input is converted into a structured research brief.
Create a scaffold. An initial draft establishes a useful report structure before live research begins. It is not treated as evidence.
Reflect and delegate. The supervisor uses internal reflection to identify gaps and delegates focused, non-overlapping research tasks to platform agents.
Research in parallel. Agents search, read, save, and compress findings using the tools available for their platform. Their results are returned with source metadata and citations.
Evaluate. The supervisor reviews the findings and may request another round.
Finalize the report. The final writer combines the brief, draft, and findings. Citation checks validate the relationship between inline citations and the final sources list.
Supervisor reflection is used for planning and control; it is not copied into the final research report. Optional subtopic evaluation and parallel subtopic reports can run after the main report when enabled in deep_research/config.py .
The supported prompt family is OPEN . Older configurations using LEGACY must switch to OPEN .
At the time of publication, Deep Dog 2 ranked 5th overall and 1st among open-source research agents on the DeepResearch Bench benchmark. These results were achieved with a relatively economical configuration: a maximum of 3 Exa searches per sub-agent, 15 minutes of research time, 20 total iterations, and DeepSeek V4 Pro and DeepSeek V4 Flash as the supervisor and sub-agent base models.
These results are a historical reproducibility profile, not a promise about current defaults. The current package defaults to DeepSeek V4 Flash for all model roles and uses a different supervisor iteration default. Benchmark rankings and scores may change as the leaderboard changes.
The following is a pricing-based estimate for one research task. It is not a controlled cost benchmark: the systems use different architectures, search providers and token budgets.
Deep Dog 2’s range is an estimate based on the DeepSeek V4 pricing (off-peak) and Exa pricing , using off-peak DeepSeek rates. Actual cost varies with prompt length, model output, cache hits, the number of delegated agents and how quickly the supervisor concludes. Exa currently advertises $20 in sign-up credits and $10 in credits each month ; Tavily provides 1,000 free credits per month , equivalent to $8 at its $0.008 per-credit pay-as-you-go rate.
Costs are easy to change by changing the configuration: shorten the research window, lower the supervisor or sub-agent iteration limits, reduce searches per sub-agent, or use fewer agents. Nemotron 3.5 Lightning has also worked successfully as a sub-agent model in Deep Dog 2 and costs approximately half as much as DeepSeek V4 Flash in the tested setup. Google’s Gemini figures are its own published estimates; see the Gemini Deep Research pricing section for the assumptions behind them.
For a detailed explanation of the reflection and delegation methods used, see the engineering article Deep Dog 2: How Reflection and Structured Delegation Improve Supervisor–Subagent Research Systems .
Requires Python 3.11+ , Git , a DeepSeek API key and an Exa API key. In your own project directory, use your existing Python environment or create one below.
Choose standard Python tooling:
# Standard Python tooling
python -m venv .venv
Or, using uv :
uv venv --python 3.11 --seed
--seed includes pip so either installer below works in this environment.
# macOS / Linux
source .venv/bin/activate
# Windows PowerShell
.venv\Scripts\Activate.ps1
Install the package directly from GitHub. Choose pip or uv:
# pip
python -m pip install " git+https://github.com/beneadie/deep_dog_2.git "
Or, with uv:
uv pip install " git+https://github.com/beneadie/deep_dog_2.git "
The installer fetches the repository, builds the package and installs its dependencies. You do not need a local source checkout or a PyPI release. This works because the repository includes Python packaging metadata in pyproject.toml ; a Git repository needs a Python package build configuration to support this kind of install. The installed distribution is named deep-dog-2 ; the Python import is deep_research .
Add your keys to a .env file in your project directory:
DEEPSEEK_API_KEY = your-deepseek-key
EXA_API_KEY = your-exa-key
Keep .env out of version control by adding it to your project's .gitignore . If those keys are already set as environment variables, no .env file is needed; existing environment variables take precedence.
Create simple_research.py in your project with the following code. The two settings near the top are optional examples: change the model or maximum research time, and leave the rest of the configuration at its defaults:
import asyncio
from dotenv import load_dotenv
load_dotenv ()
from deep_research . integration import RunConfig , run_research
MODEL = "deepseek-v4-flash"
MAX_MINUTES = 10
config = RunConfig (
supervisor_model_fallback_chain = [ MODEL ],
subagent_model_fallback_chain = [ MODEL ],
draft_report_model_fallback_chain = [ MODEL ],
research_time_max_minutes = MAX_MINUTES ,
)
result = asyncio . run (
run_research (
"How is geothermal energy developing in Europe?" ,
config = config ,
)
)
print ( result . status )
print ( result . final_report )
Run it with your project environment activated:
python simple_research.py
The example prints live progress, the run status and the Markdown report. It does not save a report file by default. Only DeepSeek and Exa credentials are needed here: the example uses DeepSeek V4 Flash for all roles and Exa for Web research. The maximum research window is set to 10 minutes above; other settings remain at their library defaults. Final writing can finish after the research window.
Updating the installed package
An installation uses a snapshot of the repository; it does not update automatically when new commits are published. To fetch and reinstall the current default-branch version, use the matching installer in your project environment:
# pip
python -m pip install --upgrade --force-reinstall " git+https://github.com/beneadie/deep_dog_2.git "
Or, with uv:
uv pip install --upgrade --reinstall-package deep-dog-2 " git+https://github.com/beneadie/deep_dog_2.git "
Reinstallation also picks up code changes that keep the same package version number. These commands can update dependencies too. For a reproducible deployment, append @<full-commit-hash> to the Git URL to install a specific revision. See pip's Git installation documentation and uv's package installation documentation .
To modify the engine or use the repository's ready-made scripts, clone the source:
git clone https://github.com/beneadie/deep_dog_2.git
cd deep_dog_2
Create and activate an environment as above, then choose an editable install:
# pip
python -m pip install -e .
Or, with uv:
uv pip install -e .
The . selects the current checkout; -e makes local source edits available without reinstalling. To update a checkout, pull the upstre

[truncated]
