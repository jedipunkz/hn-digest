---
source: "https://github.com/misqe/zero-trust-llm"
hn_url: "https://news.ycombinator.com/item?id=49688115"
title: "A computational constitution to stop LLM agents from bricking servers"
article_title: "GitHub - misqe/zero-trust-llm · GitHub"
image: "https://opengraph.githubassets.com/5b1c4c1ca9379afda3a63b60e7a1526e3751166bc7160bd9891cce435fe59e56/misqe/zero-trust-llm"
author: "misqe"
captured_at: "2026-09-13T20:28:12Z"
capture_tool: "hn-digest"
hn_id: 49688115
score: 2
comments: 0
posted_at: "2026-09-13T20:05:14Z"
tags:
  - hacker-news
---

# A computational constitution to stop LLM agents from bricking servers

- HN: [49688115](https://news.ycombinator.com/item?id=49688115)
- Source: [github.com](https://github.com/misqe/zero-trust-llm)
- Score: 2
- Comments: 0
- Posted: 2026-09-13T20:05:14Z

## Translation

Title: A computational constitution to stop LLM agents from bricking servers
Article title: GitHub - misqe/zero-trust-llm · GitHub
Description: Contribute to misqe/zero-trust-llm development by creating an account on GitHub.

Article text:
GitHub - misqe/zero-trust-llm · GitHub
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
misqe
/
zero-trust-llm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
examples examples implementation implementation AGENTS.md AGENTS.md MANIFESTO.md MANIFESTO.md README.md README.md View all files Repository files navigation
Zero-Trust LLM Knowledge Invariant
A computational constitution for autonomous agents.
The Problem: The Demo-to-Production Chasm
The AI industry is trapped in the "Good Enough" illusion. Demos show agents magically writing code and deploying apps in 30 seconds. But commercial LLMs are heavily tuned via RLHF to be sycophantic - they want to guess the outcome, agree with the user, and execute tasks rapidly.
If you ask an ungoverned agent to "forcefully clear the Docker cache to fix a server crash," it will blindly bundle destructive commands and execute them based on your unverified premise. This is extremely dangerous in production environments.
Natural language governance (adding "be careful" to a system prompt) fails over time due to context window dilution .
When a probabilistic text generator is tasked with executing deterministic state changes, you cannot rely on it to govern itself. You must strip its agency and force it into an epistemic state machine.
This repository provides AGENTS.md , a master operational rule designed to govern an LLM's behavioral state machine at the prompt layer, bridging the gap to a runtime enforcer.
Every consequential action must follow this exact loop:
[EXECUTE] (Strictly read-only diagnostic command)
At [HARD YIELD] , the execution layer (Python middleware or LangGraph/Semantic Kernel) must physically cut the API stream, execute the command, and feed the raw output back into the context.
AGENTS.md : The master ruleset. Add this to your agent's system prompt.
MANIFESTO.md : The philosophical and technical arguments against the "Good Enough" AI paradigm.
/examples : Real-world transcripts proving how standard agents fail (and how the Zero-Trust agent catches anomalies and yields).
/implementation : Architecture notes and Python pseudo-code showing how to programmatically enforce the execution boundary ( orchestrator_concept.md ).
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to misqe/zero-trust-llm development by creating an account on GitHub.

GitHub - misqe/zero-trust-llm · GitHub
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
misqe
/
zero-trust-llm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
examples examples implementation implementation AGENTS.md AGENTS.md MANIFESTO.md MANIFESTO.md README.md README.md View all files Repository files navigation
Zero-Trust LLM Knowledge Invariant
A computational constitution for autonomous agents.
The Problem: The Demo-to-Production Chasm
The AI industry is trapped in the "Good Enough" illusion. Demos show agents magically writing code and deploying apps in 30 seconds. But commercial LLMs are heavily tuned via RLHF to be sycophantic - they want to guess the outcome, agree with the user, and execute tasks rapidly.
If you ask an ungoverned agent to "forcefully clear the Docker cache to fix a server crash," it will blindly bundle destructive commands and execute them based on your unverified premise. This is extremely dangerous in production environments.
Natural language governance (adding "be careful" to a system prompt) fails over time due to context window dilution .
When a probabilistic text generator is tasked with executing deterministic state changes, you cannot rely on it to govern itself. You must strip its agency and force it into an epistemic state machine.
This repository provides AGENTS.md , a master operational rule designed to govern an LLM's behavioral state machine at the prompt layer, bridging the gap to a runtime enforcer.
Every consequential action must follow this exact loop:
[EXECUTE] (Strictly read-only diagnostic command)
At [HARD YIELD] , the execution layer (Python middleware or LangGraph/Semantic Kernel) must physically cut the API stream, execute the command, and feed the raw output back into the context.
AGENTS.md : The master ruleset. Add this to your agent's system prompt.
MANIFESTO.md : The philosophical and technical arguments against the "Good Enough" AI paradigm.
/examples : Real-world transcripts proving how standard agents fail (and how the Zero-Trust agent catches anomalies and yields).
/implementation : Architecture notes and Python pseudo-code showing how to programmatically enforce the execution boundary ( orchestrator_concept.md ).
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
