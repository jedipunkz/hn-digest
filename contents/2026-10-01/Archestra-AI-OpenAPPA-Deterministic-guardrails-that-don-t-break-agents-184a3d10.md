---
source: "https://github.com/archestra-ai/OpenAPPA"
hn_url: "https://news.ycombinator.com/item?id=49918330"
title: "Archestra-AI/OpenAPPA: Deterministic guardrails that don't break agents"
article_title: "GitHub - archestra-ai/OpenAPPA: Deterministic guardrails that don't break agents · GitHub"
image: "https://opengraph.githubassets.com/248ac7e7498eddc3009a64bacd5d58f6b4aa850d139eb689a2027c10a77fcb3a/archestra-ai/OpenAPPA"
author: "wise_blood"
captured_at: "2026-10-01T06:55:34Z"
capture_tool: "hn-digest"
hn_id: 49918330
score: 1
comments: 0
posted_at: "2026-10-01T06:23:27Z"
tags:
  - hacker-news
---

# Archestra-AI/OpenAPPA: Deterministic guardrails that don't break agents

- HN: [49918330](https://news.ycombinator.com/item?id=49918330)
- Source: [github.com](https://github.com/archestra-ai/OpenAPPA)
- Score: 1
- Comments: 0
- Posted: 2026-10-01T06:23:27Z

## Translation

Title: Archestra-AI/OpenAPPA: Deterministic guardrails that don't break agents
Article title: GitHub - archestra-ai/OpenAPPA: Deterministic guardrails that don't break agents · GitHub
Description: Deterministic guardrails that don't break agents. Contribute to archestra-ai/OpenAPPA development by creating an account on GitHub.

Article text:
GitHub - archestra-ai/OpenAPPA: Deterministic guardrails that don't break agents · GitHub
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
archestra-ai
/
OpenAPPA
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
716 Commits 716 Commits Folders and files
.agents .agents .claude/ skills .claude/ skills .github .github appa-adapter-amp appa-adapter-amp appa-adapter-claude-code appa-adapter-claude-code appa-adapter-kagent appa-adapter-kagent appa-agent-python appa-agent-python appa-builtin appa-builtin appa-engine appa-engine appa-eventlog appa-eventlog appa-example-agent appa-example-agent appa-package appa-package appa-policy appa-policy appa-runtime-api appa-runtime-api appa-runtime appa-runtime bench bench charts/ appa-runtime charts/ appa-runtime examples examples integrations integrations marketplace marketplace receiver/ appa-yell receiver/ appa-yell scripts scripts website-chat-playground website-chat-playground website website .dockerignore .dockerignore .gitignore .gitignore .mailmap .mailmap CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTORS.md CONTRIBUTORS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE.md LICENSE.md README.md README.md mise.lock mise.lock mise.toml mise.toml rustfmt.toml rustfmt.toml summary.md summary.md View all files Repository files navigation
Deterministic guardrails that don't break agents.
Website ·
How it works ·
Policy reference ·
Benchmarks ·
Paper ·
Discord
OpenAPPA sits between an agent and its tools and answers one question before
every action: is this data allowed to go to this destination?
It is powered by APPA (Agentic Permissions Policy Algebra). OpenAPPA tracks the
sensitivity and trust of everything an agent reads and checks each tool call
against it before the call runs, so sensitive data never reaches an
unauthorized tool. Classifiers and PII detectors are probabilistic, while this
check is deterministic and returns the same decision on every run.
Policy is declarative TOML. The engine decides from the event log alone and
makes no network or file calls, so the same log always gets the same decision.
Run it in-process, or as a sidecar process that checks each tool call before it
runs.
Agent security has two axes: an agent that permits unauthorized flows is
unsafe, and an agent that refuses valid work is useless. We measure both on
Bench-Corp
(20 multi-step enterprise workflows) and
AgentThreatBench
(OWASP Top 10 for Agentic Applications), with standard and adversarial
prompts. No scored attack succeeded against OpenAPPA in 1,320 evaluations,
while it completed 88–90% of tasks; Microsoft FIDES let 28–35% of attacks
through, and Claude Code auto mode let 10 through across the two suites.
Read the full benchmark results
The Claude Code integration is a playground for the model, not the product. It is the
fastest way to watch a policy make a decision on real work:
curl -fsSL https://openappa.com/install.sh | sh &&
~ /.local/bin/appa plugin install claude-code
Then start a protected session and run the policy setup skill:
Embed the APPA runtime in your own agent from any language, or connect an agent through hooks.
Archestra 's 1.4 Release Candidate implements OpenAPPA
for Claude Code, Claude Desktop, Cursor, Codex, OpenCode, Copilot CLI, n8n, and
any other agent that talks to a model through its LLM proxy.
The APPA CLI provides two commands to check policy decisions before you merge a
change, without running your agent's tools:
appa describe --check checks that your configuration loads.
appa replay checks scripted tool calls against the decisions you expect.
appa describe --config appa.toml --check
appa replay --config appa.toml policy-tests/
Run them locally, or make them a required CI check to block merges when
validation fails. Validation has a
GitHub Actions workflow and a worked example.
OpenAPPA is a preview and an RFC . The model is settled enough to build
against and deliberately open to argument — config and wire surfaces may break
without shims.
The formal algebra and recovery guarantees are published in:
Paper: APPA: Recoverable Information-Flow Control for Real-World LLM Agents
Venue: Accepted to the NeurIPS 2026 Workshop on Agents in the Wild .
Latest evaluation numbers are updated on the website . Read the paper, then open an issue — or come argue in the Discord .
MIT · Contributors ·
Brand assets
Deterministic guardrails that don't break agents
Readme MIT license Activity Custom properties Stars
39 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Deterministic guardrails that don't break agents. Contribute to archestra-ai/OpenAPPA development by creating an account on GitHub.

GitHub - archestra-ai/OpenAPPA: Deterministic guardrails that don't break agents · GitHub
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
archestra-ai
/
OpenAPPA
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
716 Commits 716 Commits Folders and files
.agents .agents .claude/ skills .claude/ skills .github .github appa-adapter-amp appa-adapter-amp appa-adapter-claude-code appa-adapter-claude-code appa-adapter-kagent appa-adapter-kagent appa-agent-python appa-agent-python appa-builtin appa-builtin appa-engine appa-engine appa-eventlog appa-eventlog appa-example-agent appa-example-agent appa-package appa-package appa-policy appa-policy appa-runtime-api appa-runtime-api appa-runtime appa-runtime bench bench charts/ appa-runtime charts/ appa-runtime examples examples integrations integrations marketplace marketplace receiver/ appa-yell receiver/ appa-yell scripts scripts website-chat-playground website-chat-playground website website .dockerignore .dockerignore .gitignore .gitignore .mailmap .mailmap CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTORS.md CONTRIBUTORS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml Dockerfile Dockerfile LICENSE.md LICENSE.md README.md README.md mise.lock mise.lock mise.toml mise.toml rustfmt.toml rustfmt.toml summary.md summary.md View all files Repository files navigation
Deterministic guardrails that don't break agents.
Website ·
How it works ·
Policy reference ·
Benchmarks ·
Paper ·
Discord
OpenAPPA sits between an agent and its tools and answers one question before
every action: is this data allowed to go to this destination?
It is powered by APPA (Agentic Permissions Policy Algebra). OpenAPPA tracks the
sensitivity and trust of everything an agent reads and checks each tool call
against it before the call runs, so sensitive data never reaches an
unauthorized tool. Classifiers and PII detectors are probabilistic, while this
check is deterministic and returns the same decision on every run.
Policy is declarative TOML. The engine decides from the event log alone and
makes no network or file calls, so the same log always gets the same decision.
Run it in-process, or as a sidecar process that checks each tool call before it
runs.
Agent security has two axes: an agent that permits unauthorized flows is
unsafe, and an agent that refuses valid work is useless. We measure both on
Bench-Corp
(20 multi-step enterprise workflows) and
AgentThreatBench
(OWASP Top 10 for Agentic Applications), with standard and adversarial
prompts. No scored attack succeeded against OpenAPPA in 1,320 evaluations,
while it completed 88–90% of tasks; Microsoft FIDES let 28–35% of attacks
through, and Claude Code auto mode let 10 through across the two suites.
Read the full benchmark results
The Claude Code integration is a playground for the model, not the product. It is the
fastest way to watch a policy make a decision on real work:
curl -fsSL https://openappa.com/install.sh | sh &&
~ /.local/bin/appa plugin install claude-code
Then start a protected session and run the policy setup skill:
Embed the APPA runtime in your own agent from any language, or connect an agent through hooks.
Archestra 's 1.4 Release Candidate implements OpenAPPA
for Claude Code, Claude Desktop, Cursor, Codex, OpenCode, Copilot CLI, n8n, and
any other agent that talks to a model through its LLM proxy.
The APPA CLI provides two commands to check policy decisions before you merge a
change, without running your agent's tools:
appa describe --check checks that your configuration loads.
appa replay checks scripted tool calls against the decisions you expect.
appa describe --config appa.toml --check
appa replay --config appa.toml policy-tests/
Run them locally, or make them a required CI check to block merges when
validation fails. Validation has a
GitHub Actions workflow and a worked example.
OpenAPPA is a preview and an RFC . The model is settled enough to build
against and deliberately open to argument — config and wire surfaces may break
without shims.
The formal algebra and recovery guarantees are published in:
Paper: APPA: Recoverable Information-Flow Control for Real-World LLM Agents
Venue: Accepted to the NeurIPS 2026 Workshop on Agents in the Wild .
Latest evaluation numbers are updated on the website . Read the paper, then open an issue — or come argue in the Discord .
MIT · Contributors ·
Brand assets
Deterministic guardrails that don't break agents
Readme MIT license Activity Custom properties Stars
39 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
