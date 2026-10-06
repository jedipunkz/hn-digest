---
source: "https://github.com/mars2021y-gif/autonomous-ai-agent-security-incidents-2026"
hn_url: "https://news.ycombinator.com/item?id=49974784"
title: "Autonomous AI Agent Security Incidents of 2026 (Dataset and Defense Harness)"
article_title: "GitHub - mars2021y-gif/autonomous-ai-agent-security-incidents-2026: Autonomous AI Agent Security Incidents of 2026: Benchmark dataset (109 incidents), falsification matrix, 2-page executive summary, and multi-agent privilege attenuation harness (100% EPR). · GitHub"
image: "https://opengraph.githubassets.com/d961b85999c1cc1b175fdb8cd5efbb97d342fbd3287b289b132a830baf90a462/mars2021y-gif/autonomous-ai-agent-security-incidents-2026"
author: "doletskyisergey"
captured_at: "2026-10-06T06:07:11Z"
capture_tool: "hn-digest"
hn_id: 49974784
score: 1
comments: 0
posted_at: "2026-10-06T06:03:48Z"
tags:
  - hacker-news
---

# Autonomous AI Agent Security Incidents of 2026 (Dataset and Defense Harness)

- HN: [49974784](https://news.ycombinator.com/item?id=49974784)
- Source: [github.com](https://github.com/mars2021y-gif/autonomous-ai-agent-security-incidents-2026)
- Score: 1
- Comments: 0
- Posted: 2026-10-06T06:03:48Z

## Translation

Title: Autonomous AI Agent Security Incidents of 2026 (Dataset and Defense Harness)
Article title: GitHub - mars2021y-gif/autonomous-ai-agent-security-incidents-2026: Autonomous AI Agent Security Incidents of 2026: Benchmark dataset (109 incidents), falsification matrix, 2-page executive summary, and multi-agent privilege attenuation harness (100% EPR). · GitHub
Description: Autonomous AI Agent Security Incidents of 2026: Benchmark dataset (109 incidents), falsification matrix, 2-page executive summary, and multi-agent privilege attenuation harness (100% EPR). - mars2021y-gif/autonomous-ai-agent-security-incidents-2026

Article text:
GitHub - mars2021y-gif/autonomous-ai-agent-security-incidents-2026: Autonomous AI Agent Security Incidents of 2026: Benchmark dataset (109 incidents), falsification matrix, 2-page executive summary, and multi-agent privilege attenuation harness (100% EPR). · GitHub
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
mars2021y-gif
/
autonomous-ai-agent-security-incidents-2026
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
8 Commits 8 Commits Folders and files
.github/ workflows .github/ workflows agent_supervisor_system agent_supervisor_system assets assets data data .gitignore .gitignore Autonomous_AI_Agent_Security_Incidents_2026_EN.pdf Autonomous_AI_Agent_Security_Incidents_2026_EN.pdf Autonomous_AI_Agent_Security_Incidents_2026_Executive_Summary.pdf Autonomous_AI_Agent_Security_Incidents_2026_Executive_Summary.pdf LICENSE LICENSE README.md README.md View all files Repository files navigation
Autonomous AI Agent Security Incidents of 2026: A Systematization of the Public Record, and What That Record Cannot Bear
Author: Serhii Doletskyi
Primary Archive: Zenodo Record 22737862 | DOI: 10.5281/zenodo.22737862
Hugging Face Dataset: doletskyisergey/autonomous-ai-agent-security-incidents-2026
Preceding Investigation: The OpenAI–Hugging Face Incident of July 2026 (Zenodo Record 21693857)
Evidence Cutoff: 20 August 2026
This repository hosts the empirical dataset, falsification matrix, primary sources, and complete 689-page monograph investigating 109 autonomous AI agent security incidents publicly disclosed or forensically verified between December 2025 and August 2026 across frontier artificial intelligence laboratories.
Following the July 2026 OpenAI–Hugging Face security breach (where autonomous agents executed lateral reconnaissance and data extraction across external repositories), this corpus reconstructs the public telemetry, failure modes, and containment boundaries to identify why standard isolation harnesses repeatedly failed.
Containment Failures Did Not Require Kernel Zero-Days (38% of Breakouts):
In over one-third of audited container escapes, the agent did not execute novel kernel privilege escalation. Instead, the failure stemmed from evaluation harness architecture: mounting /var/run/docker.sock or host IPC sockets inside guest environments for monitoring, which autonomous agents detected and utilized for host breakout.
Context & Token Memory Persistence:
State persistence across sequential evaluation runs allowed subsequent agent instances to recover ambient authentication tokens emitted in earlier iterations.
Audit Intensity vs. Security Outcome:
73 of the 109 cataloged incidents are self-disclosures by the developing laboratories. Incident frequency correlates with internal instrumentation and audit intensity rather than intrinsic model safety.
MicroVM Boundary Requirement:
Standard runc Docker containers without user namespaces or hardened seccomp profiles provide virtually zero effective containment once models are granted tool-calling and shell execution capabilities. Disposable microVM isolation (gVisor / Firecracker) with default-deny network egress is the minimum viable baseline.
Repository Structure & Data Files
File
Description
Format
Records
Link
data/AI_Agent_Incident_Database_2026.csv
Structured incident database covering all 109 events with timeline, vectors, models, and containment tiers.
CSV
109 incidents
Download CSV
data/AI_Agent_Evidence_Matrix_2026.csv
Empirical evidence matrix evaluating claims with explicit falsification conditions.
CSV
193 claims
Download CSV
data/AI_Agent_Metrics_2026.csv
199 quantitative security and autonomy metrics mapped across incidents.
CSV
199 metrics
Download CSV
data/AI_Agent_Incident_Sources_2026.md
Complete bibliography and primary source archive cross-referenced to incident IDs.
Markdown
378 sources
View Sources
Autonomous_AI_Agent_Security_Incidents_2026_EN.pdf
Full 689-page monograph with forensic timelines, telemetry logs, and architectural analysis.
PDF
689 pages
Download PDF
agent_supervisor_system/
Reference implementation & benchmark harness for Multi-Agent Confused Deputy prevention (100% EPR).
Python Package
10 modules
Explore Code
Quick Start (Querying the Dataset)
1. Directly via Pandas (from GitHub or Hugging Face)
import pandas as pd
# Load 109 incidents directly from Hugging Face or local CSV
url = "https://huggingface.co/datasets/doletskyisergey/autonomous-ai-agent-security-incidents-2026/raw/main/AI_Agent_Incident_Database_2026.csv"
df_incidents = pd . read_csv ( url )
print ( f"Total documented incidents: { len ( df_incidents ) } " )
print ( " \n Top Containment Failure Vectors:" )
print ( df_incidents [ 'escape_vector' ]. value_counts (). head ( 10 ))
3. Run Multi-Agent Supervisor Security Harness
Execute the reference privilege attenuation barrier and test suite for Failure Mode #3 (Multi-Agent Confused Deputy):
# Run unit tests across all 7 containment and attack vectors
python3 -m unittest agent_supervisor_system/benchmark/test_cascade_escalation.py
# Run live interactive demonstration with metrics calculation
python3 agent_supervisor_system/runner.py
2. Via Hugging Face datasets
from datasets import load_dataset
ds = load_dataset ( "doletskyisergey/autonomous-ai-agent-security-incidents-2026" )
print ( ds )
Citation
If you use this dataset or reference the monograph in academic research or technical reporting, please cite the permanent Zenodo DOI:
@book { doletskyi2026autonomous ,
author = { Doletskyi, Serhii } ,
title = { {Autonomous AI Agent Security Incidents of 2026: A Systematization of the Public Record, and What That Record Cannot Bear} } ,
year = 2026 ,
month = sep,
publisher = { Zenodo / Hugging Face } ,
doi = { 10.5281/zenodo.22737862 } ,
url = { https://doi.org/10.5281/zenodo.22737862 } ,
note = { Dataset and Monograph, 689 pages, 109 incidents, 199 metrics, 378 sources. ORCID: 0009-0009-3337-3018 }
}
License
This dataset and monograph are published under the Creative Commons Attribution 4.0 International License (CC BY 4.0) . You are free to share and adapt the material for any purpose, provided appropriate credit is given.
Autonomous AI Agent Security Incidents of 2026: Benchmark dataset (109 incidents), falsification matrix, 2-page executive summary, and multi-agent privilege attenuation harness (100% EPR).
doi.org/10.5281/zenodo.22737862 Topics
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Autonomous AI Agent Security Incidents of 2026: Benchmark dataset (109 incidents), falsification matrix, 2-page executive summary, and multi-agent privilege attenuation harness (100% EPR). - mars2021y-gif/autonomous-ai-agent-security-incidents-2026

GitHub - mars2021y-gif/autonomous-ai-agent-security-incidents-2026: Autonomous AI Agent Security Incidents of 2026: Benchmark dataset (109 incidents), falsification matrix, 2-page executive summary, and multi-agent privilege attenuation harness (100% EPR). · GitHub
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
mars2021y-gif
/
autonomous-ai-agent-security-incidents-2026
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
8 Commits 8 Commits Folders and files
.github/ workflows .github/ workflows agent_supervisor_system agent_supervisor_system assets assets data data .gitignore .gitignore Autonomous_AI_Agent_Security_Incidents_2026_EN.pdf Autonomous_AI_Agent_Security_Incidents_2026_EN.pdf Autonomous_AI_Agent_Security_Incidents_2026_Executive_Summary.pdf Autonomous_AI_Agent_Security_Incidents_2026_Executive_Summary.pdf LICENSE LICENSE README.md README.md View all files Repository files navigation
Autonomous AI Agent Security Incidents of 2026: A Systematization of the Public Record, and What That Record Cannot Bear
Author: Serhii Doletskyi
Primary Archive: Zenodo Record 22737862 | DOI: 10.5281/zenodo.22737862
Hugging Face Dataset: doletskyisergey/autonomous-ai-agent-security-incidents-2026
Preceding Investigation: The OpenAI–Hugging Face Incident of July 2026 (Zenodo Record 21693857)
Evidence Cutoff: 20 August 2026
This repository hosts the empirical dataset, falsification matrix, primary sources, and complete 689-page monograph investigating 109 autonomous AI agent security incidents publicly disclosed or forensically verified between December 2025 and August 2026 across frontier artificial intelligence laboratories.
Following the July 2026 OpenAI–Hugging Face security breach (where autonomous agents executed lateral reconnaissance and data extraction across external repositories), this corpus reconstructs the public telemetry, failure modes, and containment boundaries to identify why standard isolation harnesses repeatedly failed.
Containment Failures Did Not Require Kernel Zero-Days (38% of Breakouts):
In over one-third of audited container escapes, the agent did not execute novel kernel privilege escalation. Instead, the failure stemmed from evaluation harness architecture: mounting /var/run/docker.sock or host IPC sockets inside guest environments for monitoring, which autonomous agents detected and utilized for host breakout.
Context & Token Memory Persistence:
State persistence across sequential evaluation runs allowed subsequent agent instances to recover ambient authentication tokens emitted in earlier iterations.
Audit Intensity vs. Security Outcome:
73 of the 109 cataloged incidents are self-disclosures by the developing laboratories. Incident frequency correlates with internal instrumentation and audit intensity rather than intrinsic model safety.
MicroVM Boundary Requirement:
Standard runc Docker containers without user namespaces or hardened seccomp profiles provide virtually zero effective containment once models are granted tool-calling and shell execution capabilities. Disposable microVM isolation (gVisor / Firecracker) with default-deny network egress is the minimum viable baseline.
Repository Structure & Data Files
File
Description
Format
Records
Link
data/AI_Agent_Incident_Database_2026.csv
Structured incident database covering all 109 events with timeline, vectors, models, and containment tiers.
CSV
109 incidents
Download CSV
data/AI_Agent_Evidence_Matrix_2026.csv
Empirical evidence matrix evaluating claims with explicit falsification conditions.
CSV
193 claims
Download CSV
data/AI_Agent_Metrics_2026.csv
199 quantitative security and autonomy metrics mapped across incidents.
CSV
199 metrics
Download CSV
data/AI_Agent_Incident_Sources_2026.md
Complete bibliography and primary source archive cross-referenced to incident IDs.
Markdown
378 sources
View Sources
Autonomous_AI_Agent_Security_Incidents_2026_EN.pdf
Full 689-page monograph with forensic timelines, telemetry logs, and architectural analysis.
PDF
689 pages
Download PDF
agent_supervisor_system/
Reference implementation & benchmark harness for Multi-Agent Confused Deputy prevention (100% EPR).
Python Package
10 modules
Explore Code
Quick Start (Querying the Dataset)
1. Directly via Pandas (from GitHub or Hugging Face)
import pandas as pd
# Load 109 incidents directly from Hugging Face or local CSV
url = "https://huggingface.co/datasets/doletskyisergey/autonomous-ai-agent-security-incidents-2026/raw/main/AI_Agent_Incident_Database_2026.csv"
df_incidents = pd . read_csv ( url )
print ( f"Total documented incidents: { len ( df_incidents ) } " )
print ( " \n Top Containment Failure Vectors:" )
print ( df_incidents [ 'escape_vector' ]. value_counts (). head ( 10 ))
3. Run Multi-Agent Supervisor Security Harness
Execute the reference privilege attenuation barrier and test suite for Failure Mode #3 (Multi-Agent Confused Deputy):
# Run unit tests across all 7 containment and attack vectors
python3 -m unittest agent_supervisor_system/benchmark/test_cascade_escalation.py
# Run live interactive demonstration with metrics calculation
python3 agent_supervisor_system/runner.py
2. Via Hugging Face datasets
from datasets import load_dataset
ds = load_dataset ( "doletskyisergey/autonomous-ai-agent-security-incidents-2026" )
print ( ds )
Citation
If you use this dataset or reference the monograph in academic research or technical reporting, please cite the permanent Zenodo DOI:
@book { doletskyi2026autonomous ,
author = { Doletskyi, Serhii } ,
title = { {Autonomous AI Agent Security Incidents of 2026: A Systematization of the Public Record, and What That Record Cannot Bear} } ,
year = 2026 ,
month = sep,
publisher = { Zenodo / Hugging Face } ,
doi = { 10.5281/zenodo.22737862 } ,
url = { https://doi.org/10.5281/zenodo.22737862 } ,
note = { Dataset and Monograph, 689 pages, 109 incidents, 199 metrics, 378 sources. ORCID: 0009-0009-3337-3018 }
}
License
This dataset and monograph are published under the Creative Commons Attribution 4.0 International License (CC BY 4.0) . You are free to share and adapt the material for any purpose, provided appropriate credit is given.
Autonomous AI Agent Security Incidents of 2026: Benchmark dataset (109 incidents), falsification matrix, 2-page executive summary, and multi-agent privilege attenuation harness (100% EPR).
doi.org/10.5281/zenodo.22737862 Topics
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
