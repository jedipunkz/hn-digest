---
source: "https://github.com/eltzanis/SystemOneAgents"
hn_url: "https://news.ycombinator.com/item?id=50539789"
title: "System 1 Agents – a method for local, reliable and fast agentic AI"
article_title: "GitHub - eltzanis/SystemOneAgents: A method to develop System 1 Agents · GitHub"
image: "https://opengraph.githubassets.com/80ac9e13344b7b74fff56a230074c3c6e2c233bb4da47fa930aeb313be5e446c/eltzanis/SystemOneAgents"
author: "mklont"
captured_at: "2026-11-17T23:45:53Z"
capture_tool: "hn-digest"
hn_id: 50539789
score: 1
comments: 0
posted_at: "2026-11-17T23:37:38Z"
tags:
  - hacker-news
---

# System 1 Agents – a method for local, reliable and fast agentic AI

- HN: [50539789](https://news.ycombinator.com/item?id=50539789)
- Source: [github.com](https://github.com/eltzanis/SystemOneAgents)
- Score: 1
- Comments: 0
- Posted: 2026-11-17T23:37:38Z

## Translation

Title: System 1 Agents – a method for local, reliable and fast agentic AI
Article title: GitHub - eltzanis/SystemOneAgents: A method to develop System 1 Agents · GitHub
Description: A method to develop System 1 Agents. Contribute to eltzanis/SystemOneAgents development by creating an account on GitHub.

Article text:
GitHub - eltzanis/SystemOneAgents: A method to develop System 1 Agents · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
eltzanis
SystemOneAgents Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
A method to develop System 1 Agents
Readme Apache-2.0 license Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
.claude/ skills/ reproduce-results .claude/ skills/ reproduce-results code code data data docs docs figures figures results results .gitattributes .gitattributes .gitignore .gitignore CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md SystemOneAgents_logo.jpg SystemOneAgents_logo.jpg View all files Repository files navigation
Local, reliable and fast agentic AI without text generation
Code, data and results of “System 1 agents: a method for local, reliable and fast agentic AI”
A System 1 agent calls tools without generating text. A deterministic compiler turns a tool's JSON schema and the user's request into typed multiple-choice questions: which tool, which value for each argument, which item comes next in a list. A decision model answers each question with a probability for every option, in one forward pass. Code owns the control flow and assembles the call, so every tool name, argument name and value comes from the catalogue, the schema or the request: the model judges; code assembles .
The agent declines a request that no tool can do, and asks for a required value the request does not give instead of inventing one. Every question and answer can be stored, so a run can be audited and replayed decision by decision.
flowchart TD
U[User request] --> A[System 1 agent]
C[MCP tool catalogue<br/>tool descriptions + JSON schemas] --> A
A --> T{Choose tool or decline}
T -->|Tool selected| F{Capability check}
T -->|None of these| X[Decline request]
F -->|Not supported| X
F -->|Supported| P[Deterministic compiler]
U --> P
C --> P
P --> Q[Finite-choice questions<br/>for each argument]
Q <--> D[Local decision model<br/>option probabilities]
Q --> V{Required values available?}
V -->|No| M[Prompt user using<br/>schema-derived template]
M --> P
V -->|Yes| E[Assemble schema-grounded<br/>MCP tool call]
E --> R[Execute tool / record result]
R --> S{Another step needed?}
S -->|Yes| A
S -->|No| Z[End]
D --> T
D --> F
Loading
How the agent, the compiler and the decision-model back-ends work, and how they were developed: docs/ARCHITECTURE.md .
Six System 1 agents were compared with nine local LLM agents, each with reasoning on and off, on 157 MCP tools from 18 medical-imaging and biomedical servers ( evaluation set ):
Accuracy: the best System 1 agent made 93.9% of requests fully correct, the best LLM agent 91.3%. Over the 108 System 1 × LLM pairs, System 1 was significantly ahead in 73 and the LLM in 11.
No invented content: no invented tool, argument or value was detected in 10,888 System 1 calls; 5.5–19.7% of the LLM agents' calls contained at least one (lower bound).
Missing values: the System 1 agents handled 86.2–93.7% of underspecified requests correctly (right tool, stated values right, the missing value left for the user), the LLM agents 0.0–5.7%.
Cost: the OpenJev System 1 agent took 2.6 s per request and 0.228 Wh per fully correct request; Qwen3.5-27B with reasoning off took 4.5 s and 0.413 Wh.
Every table of the paper is in results/analysis/tables.md .
Folder
Contents
code/
the System 1 agent (agent, compiler, decision-model back-ends), the LLM-agent baseline, the scorer, and scripts to run, replay and analyse; see code/README.md
data/
the evaluation requests with their labels, and the two MCP tool catalogues; see data/README.md
results/
every agent's records, the stored decision-model answers, the LLM completions, the analysis outputs and the fabrication review
figures/
the manuscript figures (PDF, SVG, PNG)
docs/ARCHITECTURE.md
how a System 1 agent works, and the method used to develop it
.claude/skills/reproduce-results/
a skill that guides a coding agent (for example Claude Code) through reproducing the results
Requirements
To
You need
reproduce the paper's results
Python 3.11 or later (3.12 was used) and Node.js (24 was used; the biomedical scorer runs the ClinicalTrials.gov server's own normalisers). A CPU is enough.
run System 1 agents
in addition, llama.cpp's llama-server and an NVIDIA GPU with room for the decision model (from 7.2 GiB for decider-4b; see the models )
The commands below are for a Linux shell; on Windows, use WSL. The study ran on Ubuntu 24.04.
Reproduce the paper's results (CPU, a few minutes)
This rebuilds every System 1 call from the stored model answers, rescores every agent, and regenerates the tables, the fabrication analysis, Table 1 and the figures.
git clone https://github.com/eltzanis/SystemOneAgents && cd SystemOneAgents
(cd data && sha256sum -c --quiet SHA256SUMS.txt) && (cd results && sha256sum -c --quiet SHA256SUMS.txt) # files intact
cd code
conda create -n s1agents python=3.12 && conda activate s1agents
pip install -e " .[test,figures] " # exact versions used: pip install -r requirements-lock.txt
python -m pytest # no GPU, no model
# replay: rebuild every System 1 call from the stored answers, compare with the released records
python scripts/replay_s1.py --catalogue ../data/catalogues/imaging_tools.json \
--data ../data/requests/imaging_test.jsonl ../data/requests/external.jsonl \
--answers ../results/system1/answers --records ../results/system1/records
python scripts/replay_s1.py --catalogue ../data/catalogues/biomedical_tools.json \
--data ../data/requests/biomedical_test.jsonl \
--answers ../results/system1/answers --records ../results/system1/records
# analysis: every table of the manuscript
python scripts/analyse.py --sensitivity --out runs/analysis
# fabrication analysis (Fig. 4a, Extended Data Table 7)
python scripts/fabrication.py --check-sample --out runs/fabrication
# Table 1: System 1 against LLM agents, category by category
python scripts/scorecard.py --out runs/scorecard
# figures
python scripts/figures.py --out runs/figures
What each step should give:
The figures were drawn with Liberation Sans and Liberation Mono 2.1.5 and URW Nimbus Sans, which the study's Ubuntu 24.04 desktop had. Without them, matplotlib falls back to another font: the plotted values stay the same, but the lettering and so the PNG files differ. On a minimal Ubuntu 22.04 (for example WSL), sudo apt install fonts-liberation2 fonts-urw-base35 gave byte-identical PNGs.
To have a coding agent run and check all of this, open the repository in Claude Code and ask it to reproduce the results: it uses the reproduce-results skill.
Build a System 1 agent for your own tools
cd code
pip install -e .
2. Get llama.cpp and a decision model
The agent reads its decision model from llama.cpp's llama-server . The study used release b11929; code/README.md gives the download commands and checksums. Its paths are set at the top of code/configs/models.yaml ( llama_server , cuda_libs ): change them to use your own build.
The six decision models of the study, all used without fine-tuning (measured on the study's GPU, over the evaluation set):
decider-4b is the quickest to try; OpenJev and pplx-decider-27b were the most accurate. Check each model's licence on its model card: OpenJev is CC-BY-NC-4.0 (non-commercial), and the Perplexity models may require accepting their terms on Hugging Face.
Model files go to code/models/<folder>/ , with the folders of code/configs/models.yaml . For decider-4b, from code/ :
hf download Mapika/decider-4b-GGUF decider-4b-v3-Q8_0.gguf tokenizer.json tokenizer_config.json decider_config.json \
--revision d3a78581751d518bbfa9e871a6d84059554d53c5 --local-dir models/decider-4b
sha256sum models/decider-4b/decider-4b-v3-Q8_0.gguf # d29bda3381301c81d1ecaf8adf2f9df94f724f0179be4ddb11c898f38713ffd0
The sources, revisions and checksums of the other models are in code/README.md .
s1a serve decider-4b # llama-server on 127.0.0.1, port from configs/models.yaml; leave it running
4. Try it on the study's tools
s1a plan " Segment the liver and the spleen in /data/ct/p01.nii.gz " \
--catalogue ../data/catalogues/imaging_tools.json --model decider-4b
plan selects a tool and fills its arguments without calling anything. It prints every decision (question, top options, probabilities) and the resulting call, and appends the decisions to runs/traces/ .
5. Connect your own MCP servers
List your servers in a YAML file as in code/configs/servers.example.yaml (each one started over stdio with command , args , and optionally cwd and env ). Then save their tools as a catalogue and use it:
s1a catalog --servers my_servers.yaml --out my_tools.json # start each server once and save its tools
s1a plan " ... " --catalogue my_tools.json --model decider-4b # dry run: decisions and call, nothing executed
s1a run " ... " --catalogue my_tools.json --servers my_servers.yaml --model decider-4b
run plans the call, asks you for any required value the request does not give, asks for confirmation ( --yes skips it) and calls the tool. --max-steps N lets the agent continue with further calls after the first. --args-model uses a second decision model for the argument questions.
code/configs/agent.yaml holds the agent's settings (pass another file with --settings ):
from s1agent import build_agent
agent = build_agent ( "decider-4b" , "../data/catalogues/imaging_tools.json" ) # the model's llama-server must be running
plan = agent . plan ( "Segment the liver in /data/ct/p01.nii.gz" )
if plan . filled is None :
print ( "declined: no tool fits, or the selection confidence is below select_min_conf" )
else :
print ( plan . selection . tool . name , plan . selection . confidence )
print ( plan . filled . args , plan . filled . missing ) # the call's arguments, and required values the request does not give
Add another decision model
A model that llama-server can serve with one of the three supported read-outs is one more entry in code/configs/models.yaml : decider-llamacpp (log-probabilities of the option labels), pplx-llamacpp (hidden state at a read-out marker) or systemone-llamacpp (llama-server's own decision endpoint). Any other model needs a back-end: a subclass of Backend in code/src/s1agent/backends/ whose _choice(state, question, options, kind) returns a probability for each option. lexical.py is the shortest example.
Rerunning the System 1 and LLM agents needs llama.cpp b11929 and the model files. code/README.md gives their sources, revisions and checksums, and the run commands ( scripts/run_s1.py , scripts/run_llm.py ). The runs used one NVIDIA RTX PRO 5000 Blackwell (48 GB) on Ubuntu 24.04.
A decision mo

[truncated]

## Original Extract

A method to develop System 1 Agents. Contribute to eltzanis/SystemOneAgents development by creating an account on GitHub.

GitHub - eltzanis/SystemOneAgents: A method to develop System 1 Agents · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
eltzanis
SystemOneAgents Public Notifications You must be signed in to change notification settings
Star 7 ( 7 ) You must be signed in to star a repository
A method to develop System 1 Agents
Readme Apache-2.0 license Cite this repository Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
.claude/ skills/ reproduce-results .claude/ skills/ reproduce-results code code data data docs docs figures figures results results .gitattributes .gitattributes .gitignore .gitignore CITATION.cff CITATION.cff LICENSE LICENSE README.md README.md SystemOneAgents_logo.jpg SystemOneAgents_logo.jpg View all files Repository files navigation
Local, reliable and fast agentic AI without text generation
Code, data and results of “System 1 agents: a method for local, reliable and fast agentic AI”
A System 1 agent calls tools without generating text. A deterministic compiler turns a tool's JSON schema and the user's request into typed multiple-choice questions: which tool, which value for each argument, which item comes next in a list. A decision model answers each question with a probability for every option, in one forward pass. Code owns the control flow and assembles the call, so every tool name, argument name and value comes from the catalogue, the schema or the request: the model judges; code assembles .
The agent declines a request that no tool can do, and asks for a required value the request does not give instead of inventing one. Every question and answer can be stored, so a run can be audited and replayed decision by decision.
flowchart TD
U[User request] --> A[System 1 agent]
C[MCP tool catalogue<br/>tool descriptions + JSON schemas] --> A
A --> T{Choose tool or decline}
T -->|Tool selected| F{Capability check}
T -->|None of these| X[Decline request]
F -->|Not supported| X
F -->|Supported| P[Deterministic compiler]
U --> P
C --> P
P --> Q[Finite-choice questions<br/>for each argument]
Q <--> D[Local decision model<br/>option probabilities]
Q --> V{Required values available?}
V -->|No| M[Prompt user using<br/>schema-derived template]
M --> P
V -->|Yes| E[Assemble schema-grounded<br/>MCP tool call]
E --> R[Execute tool / record result]
R --> S{Another step needed?}
S -->|Yes| A
S -->|No| Z[End]
D --> T
D --> F
Loading
How the agent, the compiler and the decision-model back-ends work, and how they were developed: docs/ARCHITECTURE.md .
Six System 1 agents were compared with nine local LLM agents, each with reasoning on and off, on 157 MCP tools from 18 medical-imaging and biomedical servers ( evaluation set ):
Accuracy: the best System 1 agent made 93.9% of requests fully correct, the best LLM agent 91.3%. Over the 108 System 1 × LLM pairs, System 1 was significantly ahead in 73 and the LLM in 11.
No invented content: no invented tool, argument or value was detected in 10,888 System 1 calls; 5.5–19.7% of the LLM agents' calls contained at least one (lower bound).
Missing values: the System 1 agents handled 86.2–93.7% of underspecified requests correctly (right tool, stated values right, the missing value left for the user), the LLM agents 0.0–5.7%.
Cost: the OpenJev System 1 agent took 2.6 s per request and 0.228 Wh per fully correct request; Qwen3.5-27B with reasoning off took 4.5 s and 0.413 Wh.
Every table of the paper is in results/analysis/tables.md .
Folder
Contents
code/
the System 1 agent (agent, compiler, decision-model back-ends), the LLM-agent baseline, the scorer, and scripts to run, replay and analyse; see code/README.md
data/
the evaluation requests with their labels, and the two MCP tool catalogues; see data/README.md
results/
every agent's records, the stored decision-model answers, the LLM completions, the analysis outputs and the fabrication review
figures/
the manuscript figures (PDF, SVG, PNG)
docs/ARCHITECTURE.md
how a System 1 agent works, and the method used to develop it
.claude/skills/reproduce-results/
a skill that guides a coding agent (for example Claude Code) through reproducing the results
Requirements
To
You need
reproduce the paper's results
Python 3.11 or later (3.12 was used) and Node.js (24 was used; the biomedical scorer runs the ClinicalTrials.gov server's own normalisers). A CPU is enough.
run System 1 agents
in addition, llama.cpp's llama-server and an NVIDIA GPU with room for the decision model (from 7.2 GiB for decider-4b; see the models )
The commands below are for a Linux shell; on Windows, use WSL. The study ran on Ubuntu 24.04.
Reproduce the paper's results (CPU, a few minutes)
This rebuilds every System 1 call from the stored model answers, rescores every agent, and regenerates the tables, the fabrication analysis, Table 1 and the figures.
git clone https://github.com/eltzanis/SystemOneAgents && cd SystemOneAgents
(cd data && sha256sum -c --quiet SHA256SUMS.txt) && (cd results && sha256sum -c --quiet SHA256SUMS.txt) # files intact
cd code
conda create -n s1agents python=3.12 && conda activate s1agents
pip install -e " .[test,figures] " # exact versions used: pip install -r requirements-lock.txt
python -m pytest # no GPU, no model
# replay: rebuild every System 1 call from the stored answers, compare with the released records
python scripts/replay_s1.py --catalogue ../data/catalogues/imaging_tools.json \
--data ../data/requests/imaging_test.jsonl ../data/requests/external.jsonl \
--answers ../results/system1/answers --records ../results/system1/records
python scripts/replay_s1.py --catalogue ../data/catalogues/biomedical_tools.json \
--data ../data/requests/biomedical_test.jsonl \
--answers ../results/system1/answers --records ../results/system1/records
# analysis: every table of the manuscript
python scripts/analyse.py --sensitivity --out runs/analysis
# fabrication analysis (Fig. 4a, Extended Data Table 7)
python scripts/fabrication.py --check-sample --out runs/fabrication
# Table 1: System 1 against LLM agents, category by category
python scripts/scorecard.py --out runs/scorecard
# figures
python scripts/figures.py --out runs/figures
What each step should give:
The figures were drawn with Liberation Sans and Liberation Mono 2.1.5 and URW Nimbus Sans, which the study's Ubuntu 24.04 desktop had. Without them, matplotlib falls back to another font: the plotted values stay the same, but the lettering and so the PNG files differ. On a minimal Ubuntu 22.04 (for example WSL), sudo apt install fonts-liberation2 fonts-urw-base35 gave byte-identical PNGs.
To have a coding agent run and check all of this, open the repository in Claude Code and ask it to reproduce the results: it uses the reproduce-results skill.
Build a System 1 agent for your own tools
cd code
pip install -e .
2. Get llama.cpp and a decision model
The agent reads its decision model from llama.cpp's llama-server . The study used release b11929; code/README.md gives the download commands and checksums. Its paths are set at the top of code/configs/models.yaml ( llama_server , cuda_libs ): change them to use your own build.
The six decision models of the study, all used without fine-tuning (measured on the study's GPU, over the evaluation set):
decider-4b is the quickest to try; OpenJev and pplx-decider-27b were the most accurate. Check each model's licence on its model card: OpenJev is CC-BY-NC-4.0 (non-commercial), and the Perplexity models may require accepting their terms on Hugging Face.
Model files go to code/models/<folder>/ , with the folders of code/configs/models.yaml . For decider-4b, from code/ :
hf download Mapika/decider-4b-GGUF decider-4b-v3-Q8_0.gguf tokenizer.json tokenizer_config.json decider_config.json \
--revision d3a78581751d518bbfa9e871a6d84059554d53c5 --local-dir models/decider-4b
sha256sum models/decider-4b/decider-4b-v3-Q8_0.gguf # d29bda3381301c81d1ecaf8adf2f9df94f724f0179be4ddb11c898f38713ffd0
The sources, revisions and checksums of the other models are in code/README.md .
s1a serve decider-4b # llama-server on 127.0.0.1, port from configs/models.yaml; leave it running
4. Try it on the study's tools
s1a plan " Segment the liver and the spleen in /data/ct/p01.nii.gz " \
--catalogue ../data/catalogues/imaging_tools.json --model decider-4b
plan selects a tool and fills its arguments without calling anything. It prints every decision (question, top options, probabilities) and the resulting call, and appends the decisions to runs/traces/ .
5. Connect your own MCP servers
List your servers in a YAML file as in code/configs/servers.example.yaml (each one started over stdio with command , args , and optionally cwd and env ). Then save their tools as a catalogue and use it:
s1a catalog --servers my_servers.yaml --out my_tools.json # start each server once and save its tools
s1a plan " ... " --catalogue my_tools.json --model decider-4b # dry run: decisions and call, nothing executed
s1a run " ... " --catalogue my_tools.json --servers my_servers.yaml --model decider-4b
run plans the call, asks you for any required value the request does not give, asks for confirmation ( --yes skips it) and calls the tool. --max-steps N lets the agent continue with further calls after the first. --args-model uses a second decision model for the argument questions.
code/configs/agent.yaml holds the agent's settings (pass another file with --settings ):
from s1agent import build_agent
agent = build_agent ( "decider-4b" , "../data/catalogues/imaging_tools.json" ) # the model's llama-server must be running
plan = agent . plan ( "Segment the liver in /data/ct/p01.nii.gz" )
if plan . filled is None :
print ( "declined: no tool fits, or the selection confidence is below select_min_conf" )
else :
print ( plan . selection . tool . name , plan . selection . confidence )
print ( plan . filled . args , plan . filled . missing ) # the call's arguments, and required values the request does not give
Add another decision model
A model that llama-server can serve with one of the three supported read-outs is one more entry in code/configs/models.yaml : decider-llamacpp (log-probabilities of the option labels), pplx-llamacpp (hidden state at a read-out marker) or systemone-llamacpp (llama-server's own decision endpoint). Any other model needs a back-end: a subclass of Backend in code/src/s1agent/backends/ whose _choice(state, question, options, kind) returns a probability for each option. lexical.py is the shortest example.
Rerunning the System 1 and LLM agents needs llama.cpp b11929 and the model files. code/README.md gives their sources, revisions and checksums, and the run commands ( scripts/run_s1.py , scripts/run_llm.py ). The runs used one NVIDIA RTX PRO 5000 Blackwell (48 GB) on Ubuntu 24.04.
A decision mo

[truncated]
