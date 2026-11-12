---
source: "https://www.harness.io/blog/why-ai-coding-agents-lose-context-after-pr"
hn_url: "https://news.ycombinator.com/item?id=50463445"
title: "Why AI Coding Agents Lose Context After the PR, and How to Keep It"
article_title: "Why AI Coding Agents Lose Context After the PR"
image: "https://cdn.prod.website-files.com/69e786e187c925d08ee02b92/6aeb5399161e9a06350b61e3_6aeb5243bad91f63189b381a_why-ai-coding-agents-lose-context-after-the-pr-and-how-to-keep-it-og.png"
author: "sonichigo"
captured_at: "2026-11-12T05:59:17Z"
capture_tool: "hn-digest"
hn_id: 50463445
score: 1
comments: 0
posted_at: "2026-11-12T05:35:21Z"
tags:
  - hacker-news
---

# Why AI Coding Agents Lose Context After the PR, and How to Keep It

- HN: [50463445](https://news.ycombinator.com/item?id=50463445)
- Source: [www.harness.io](https://www.harness.io/blog/why-ai-coding-agents-lose-context-after-pr)
- Score: 1
- Comments: 0
- Posted: 2026-11-12T05:35:21Z

## Translation

Title: Why AI Coding Agents Lose Context After the PR, and How to Keep It
Article title: Why AI Coding Agents Lose Context After the PR
Description: AI coding agents lose context once a PR enters the pipeline. See how code context and delivery context work together across the SDLC | Blog

Article text:
Why AI Coding Agents Lose Context After the PR
-->
-->
Cosmos Software Factory Agent Continuous Delivery & GitOps For the complete documentation index, see llms.txt . A Markdown version of this page is available by requesting this URL with the header "Accept: text/markdown". Platform Harness Autonomous SDLC Platform Deliver at machine speed.
Explore Overview Documentation Knowledge Graph Trust Center Popular Use Cases AI DLC Agent DLC Mythos Readiness AI Cost Management Software Delivery Agent Get code and agents from commit to production in minutes.
Discover bugs and resilience weaknesses before your users do.
Catch vulnerabilities before they ship and protect your runtime.
Own production and manage costs without surprises.
Take engineering work from ticket to code.
More products Incidents Developer Portal Resilience Testing AI Evals Runtime Configuration Why Harness Pricing Customers Resources Build
Platform Explore Overview Documentation Knowledge Graph Trust Center Popular Use Cases AI DLC Agent DLC Mythos Readiness AI Cost Management AGENTS Software Delivery Agent Get code and agents from commit to production in minutes.
Discover bugs and resilience weaknesses before your users do.
Catch vulnerabilities before they ship and protect your runtime.
Own production and manage costs without surprises.
Take engineering work from ticket to code.
About Us Careers Security Press & News Legal Partners Contact Us Sign in See a Demo Start for Free
Blog → Cosmos Software Factory Agent
Why AI Coding Agents Lose Context After the PR, and How to Keep It
| Harness Blog
AI coding agents lose context once a PR enters the pipeline. See how code context and delivery context work
AI coding agents need more than repository-level knowledge to deliver reliable software. Combining code context from Augment's Code Context Engine with delivery context from Harness's Software Delivery Knowledge Graph helps agents understand dependencies, production impact, and release risks across the SDLC.
A governed agentic workflow connects context, execution, and governance through automated reviews, policy gates, approvals, and continuous verification, ensuring AI generated changes are safe to ship.
While deeper two-way context sharing is still planned, Harness and Cosmos are building toward an Autonomous SDLC where delivery evidence continuously informs future engineering decisions.
The Context Cliff: What Happens After an AI Agent Opens a PR
An AI coding agent can read a repository, trace its dependencies, edit a dozen files, and open a pull request that passes its own tests. Then the PR enters the delivery pipeline, and much of what the agent learned stays behind.
The pipeline produces evidence the agent never sees: a failed build, a vulnerable dependency flagged by a scan, a latency regression after deployment, an incident traced to a change that looked safe. That evidence could improve the original change and prevent the next failure, but only if something carries it back into the engineering workflow.
Writing code is often no longer the slow part. Moving a change through build, test, security, deployment, and runtime is, and each of those stages is a question of risk. Context engineering across the SDLC is the work of connecting them.
This post covers what context an agent needs, how the Code Context Engine and the Software Delivery Knowledge Graph supply two halves of it, and what works today and what Harness has said is planned.
Context Engineering: The Right Information at the Right Time
Context engineering is the work of selecting, organizing, retrieving, and maintaining what an AI agent knows while it works on a task. It covers more than prompt wording. An agent may need architectural decisions, dependency relationships, repository conventions, test results, security findings, deployment history, and operating constraints.
Agent memory and context engineering do different jobs. Memory stores what might be useful across sessions. Context engineering decides what goes in front of the agent for the current task. A large memory store does not help when the retrieved material is outdated, contradictory, or irrelevant.
Augment Code measured the effect of retrieval quality in its February 2026 Context Engine MCP benchmark . Its most useful finding concerns model choice: a weaker model with great context (Sonnet with the MCP) can outperform a stronger model with poor context (Opus without it). MCP-enabled runs also needed fewer tool calls and conversation turns.
The test ran 300 Elasticsearch pull requests with three prompts each, 900 attempts in total, scored on correctness, completeness, best practices, code reuse, and unsolicited documentation. Adding the Code Context Engine improved results by 80% for Claude Code, 71% for Cursor, and 30% for Cursor with Composer-1 . The original post ran these against Opus 4.5, and Augment's product page now lists the same figures under Opus 4.6.
These are Augment's own results on one codebase. They show that retrieval quality moves agent output, and they do not predict the gain on your repository.
Two Halves of the Problem: Code Context and Delivery Context
A repository shows how code works. It does not show which services are business-critical, which environments receive a change, or what the last deployment did. Agents need both kinds of context.
Code context comes from the Code Context Engine, which Harness brought in with the Augment Code assets it acquired. It indexes a codebase so an agent can find relevant code and its relationships instead of grepping files. Harness's governed software factory post describes it as indexing in real time across 100M+ lines, including commit history and which code is still active. The factory post says generic text embeddings help little on code, because similar-looking code is often unrelated, so the Code Context Engine uses embedding models trained on code structure. It is exposed over MCP, so any compatible agent can use it.
Delivery context comes from the Software Delivery Knowledge Graph, a typed model of how software is built, tested, secured, deployed, and operated. Harness's governed software factory post reports 80+ entity types across CI, CD, security, cost, and runtime, and says an agent question costs about 12K tokens through the graph against 250K+ through raw APIs.
The underlying estimate is in Harness's earlier knowledge graph post . For one question, which pipelines have the highest failure rate across services and vulnerabilities, it gives about 250,000 to 350,000 input tokens and 5+ LLM calls through raw MCP API calls, against about 12,000 tokens and 2 to 3 calls through the graph. It labels the source internal benchmarking and does not describe the method, so read the figures as an order-of-magnitude illustration for one query, not a benchmark.
For how Harness checks the graph's answers, see the validation write-up , which reports 22 of 27 scenarios validated, about 81% coverage. The architecture write-up explains how schemas, synchronization, and declarative queries make those answers repeatable.
Code context explains the change. Delivery context explains its consequences.
What Works Today and What Is Planned
On October 8, 2026, Harness announced the acquisition of Cosmos, the Code Context Engine, and related technology. Cosmos is now the Harness Cosmos Software Factory Agent, which automates engineering work from idea to code.. It joins the Harness Autonomous SDLC Platform, where the Software Delivery, Security Testing, Runtime Protection, and Cost Management agents take that code to production and beyond.
What the factory post describes today (published October 13, updated October 19):
A scan flags a critical CVE and starts a Cosmos workflow
An Expert uses the Code Context Engine to find where the vulnerable dependency is imported, then queries the Knowledge Graph for the affected services running in production, their business criticality, and remaining SLO budget
Fixes open as PRs, and reviewers and the Verifier check them
Each merge triggers a pipeline, with OPA policy and approval gates set by the service's risk tier
Continuous Verification watches each deployment, and an audit trail records the path
Connecting the Code Context Engine to the Software Delivery Knowledge Graph is described in the announcement as planned work, so read that two-way link as direction and not as a shipped feature. Harness plans three flows:
Delivery context flows into coding, so testing expectations, security requirements, and deployment constraints inform a change before it is written
Codebase understanding flows into delivery, so the Harness agents can make better decisions about dependencies, risk, and the validation a change needs
Findings from review, testing, security checks, and deployment verification flow back to the engineering workflow with the evidence needed to correct and revalidate the change
The practical reading: an agent can use both kinds of context in one workflow today, and the tighter two-way connection between them is the part still being built.
The Agent Harness: Context, Execution, Governance
In this post, the agent harness is the architecture around a coding agent that makes autonomous work dependable. It has three layers:
Context gives agents the codebase, delivery, and runtime information they need to understand the task
Execution coordinates agents and tools to plan, implement, review, test, and deliver a change through repeatable workflows
Governance enforces access controls, security policies, approvals, and audit trails so autonomous work stays inside agreed limits
The Harness agents supply the governance layer. An agent's recommendation grants no permission to deploy: policy as code, RBAC, approval gates, continuous verification, and rollback decide what proceeds. In the governed software factory design , pipeline work runs on Delegates inside the customer's own network, and a Zero Trust Service checks every task before it runs. Shell scripts are parsed into an AST and checked against a command allowlist, container images against a registry allowlist, and each step against the resolved pipeline YAML. The service fails to close, so nothing executes if it is unreachable.
AI code review and governed pipelines answer different questions. Review tells you whether a diff looks correct. A governed pipeline decides whether the change may ship and checks how it behaved once it did.
A critical CVE in a shared library shows how the two kinds of context combine. The steps follow the flow Harness describes in its governed software factory post.
The diagram shows four stages, and steps 5 and 6 below are the return path through the Knowledge Graph. Solid arrows show a change moving to production and its results landing in the Software Delivery Knowledge Graph. The dashed arrow is the planned link that lets delivery context inform the next change.
Scope in code. The Security Testing Agent's scan names the dependency and the vulnerable version. An Expert uses the Code Context Engine to find where the dependency is imported, which functions call it, and which tests cover it.
Scope in production. The Software Delivery Knowledge Graph shows which affected services run in production, which are business-critical, and how much SLO budget each has left. The same library in an internal utility and in a customer-facing authentication service gets different priority.
Fix and review. Agents propose the upgrade, update code and tests, and open a pull request in each affected repository. Reviewers and the Verifier check correctness, security, and maintainability.
Gate. The Software Delivery Agent's pipeline applies policy and approval gates by the service's risk tier. A change that fails a mandatory check does not proceed.
Verify in production. Continuous Verification compares each deployment with a baseline captured when the deployment starts. If behavior deviates, the pipeline rolls the deployment back with the anomalou

[truncated]

## Original Extract

AI coding agents lose context once a PR enters the pipeline. See how code context and delivery context work together across the SDLC | Blog

Why AI Coding Agents Lose Context After the PR
-->
-->
Cosmos Software Factory Agent Continuous Delivery & GitOps For the complete documentation index, see llms.txt . A Markdown version of this page is available by requesting this URL with the header "Accept: text/markdown". Platform Harness Autonomous SDLC Platform Deliver at machine speed.
Explore Overview Documentation Knowledge Graph Trust Center Popular Use Cases AI DLC Agent DLC Mythos Readiness AI Cost Management Software Delivery Agent Get code and agents from commit to production in minutes.
Discover bugs and resilience weaknesses before your users do.
Catch vulnerabilities before they ship and protect your runtime.
Own production and manage costs without surprises.
Take engineering work from ticket to code.
More products Incidents Developer Portal Resilience Testing AI Evals Runtime Configuration Why Harness Pricing Customers Resources Build
Platform Explore Overview Documentation Knowledge Graph Trust Center Popular Use Cases AI DLC Agent DLC Mythos Readiness AI Cost Management AGENTS Software Delivery Agent Get code and agents from commit to production in minutes.
Discover bugs and resilience weaknesses before your users do.
Catch vulnerabilities before they ship and protect your runtime.
Own production and manage costs without surprises.
Take engineering work from ticket to code.
About Us Careers Security Press & News Legal Partners Contact Us Sign in See a Demo Start for Free
Blog → Cosmos Software Factory Agent
Why AI Coding Agents Lose Context After the PR, and How to Keep It
| Harness Blog
AI coding agents lose context once a PR enters the pipeline. See how code context and delivery context work
AI coding agents need more than repository-level knowledge to deliver reliable software. Combining code context from Augment's Code Context Engine with delivery context from Harness's Software Delivery Knowledge Graph helps agents understand dependencies, production impact, and release risks across the SDLC.
A governed agentic workflow connects context, execution, and governance through automated reviews, policy gates, approvals, and continuous verification, ensuring AI generated changes are safe to ship.
While deeper two-way context sharing is still planned, Harness and Cosmos are building toward an Autonomous SDLC where delivery evidence continuously informs future engineering decisions.
The Context Cliff: What Happens After an AI Agent Opens a PR
An AI coding agent can read a repository, trace its dependencies, edit a dozen files, and open a pull request that passes its own tests. Then the PR enters the delivery pipeline, and much of what the agent learned stays behind.
The pipeline produces evidence the agent never sees: a failed build, a vulnerable dependency flagged by a scan, a latency regression after deployment, an incident traced to a change that looked safe. That evidence could improve the original change and prevent the next failure, but only if something carries it back into the engineering workflow.
Writing code is often no longer the slow part. Moving a change through build, test, security, deployment, and runtime is, and each of those stages is a question of risk. Context engineering across the SDLC is the work of connecting them.
This post covers what context an agent needs, how the Code Context Engine and the Software Delivery Knowledge Graph supply two halves of it, and what works today and what Harness has said is planned.
Context Engineering: The Right Information at the Right Time
Context engineering is the work of selecting, organizing, retrieving, and maintaining what an AI agent knows while it works on a task. It covers more than prompt wording. An agent may need architectural decisions, dependency relationships, repository conventions, test results, security findings, deployment history, and operating constraints.
Agent memory and context engineering do different jobs. Memory stores what might be useful across sessions. Context engineering decides what goes in front of the agent for the current task. A large memory store does not help when the retrieved material is outdated, contradictory, or irrelevant.
Augment Code measured the effect of retrieval quality in its February 2026 Context Engine MCP benchmark . Its most useful finding concerns model choice: a weaker model with great context (Sonnet with the MCP) can outperform a stronger model with poor context (Opus without it). MCP-enabled runs also needed fewer tool calls and conversation turns.
The test ran 300 Elasticsearch pull requests with three prompts each, 900 attempts in total, scored on correctness, completeness, best practices, code reuse, and unsolicited documentation. Adding the Code Context Engine improved results by 80% for Claude Code, 71% for Cursor, and 30% for Cursor with Composer-1 . The original post ran these against Opus 4.5, and Augment's product page now lists the same figures under Opus 4.6.
These are Augment's own results on one codebase. They show that retrieval quality moves agent output, and they do not predict the gain on your repository.
Two Halves of the Problem: Code Context and Delivery Context
A repository shows how code works. It does not show which services are business-critical, which environments receive a change, or what the last deployment did. Agents need both kinds of context.
Code context comes from the Code Context Engine, which Harness brought in with the Augment Code assets it acquired. It indexes a codebase so an agent can find relevant code and its relationships instead of grepping files. Harness's governed software factory post describes it as indexing in real time across 100M+ lines, including commit history and which code is still active. The factory post says generic text embeddings help little on code, because similar-looking code is often unrelated, so the Code Context Engine uses embedding models trained on code structure. It is exposed over MCP, so any compatible agent can use it.
Delivery context comes from the Software Delivery Knowledge Graph, a typed model of how software is built, tested, secured, deployed, and operated. Harness's governed software factory post reports 80+ entity types across CI, CD, security, cost, and runtime, and says an agent question costs about 12K tokens through the graph against 250K+ through raw APIs.
The underlying estimate is in Harness's earlier knowledge graph post . For one question, which pipelines have the highest failure rate across services and vulnerabilities, it gives about 250,000 to 350,000 input tokens and 5+ LLM calls through raw MCP API calls, against about 12,000 tokens and 2 to 3 calls through the graph. It labels the source internal benchmarking and does not describe the method, so read the figures as an order-of-magnitude illustration for one query, not a benchmark.
For how Harness checks the graph's answers, see the validation write-up , which reports 22 of 27 scenarios validated, about 81% coverage. The architecture write-up explains how schemas, synchronization, and declarative queries make those answers repeatable.
Code context explains the change. Delivery context explains its consequences.
What Works Today and What Is Planned
On October 8, 2026, Harness announced the acquisition of Cosmos, the Code Context Engine, and related technology. Cosmos is now the Harness Cosmos Software Factory Agent, which automates engineering work from idea to code.. It joins the Harness Autonomous SDLC Platform, where the Software Delivery, Security Testing, Runtime Protection, and Cost Management agents take that code to production and beyond.
What the factory post describes today (published October 13, updated October 19):
A scan flags a critical CVE and starts a Cosmos workflow
An Expert uses the Code Context Engine to find where the vulnerable dependency is imported, then queries the Knowledge Graph for the affected services running in production, their business criticality, and remaining SLO budget
Fixes open as PRs, and reviewers and the Verifier check them
Each merge triggers a pipeline, with OPA policy and approval gates set by the service's risk tier
Continuous Verification watches each deployment, and an audit trail records the path
Connecting the Code Context Engine to the Software Delivery Knowledge Graph is described in the announcement as planned work, so read that two-way link as direction and not as a shipped feature. Harness plans three flows:
Delivery context flows into coding, so testing expectations, security requirements, and deployment constraints inform a change before it is written
Codebase understanding flows into delivery, so the Harness agents can make better decisions about dependencies, risk, and the validation a change needs
Findings from review, testing, security checks, and deployment verification flow back to the engineering workflow with the evidence needed to correct and revalidate the change
The practical reading: an agent can use both kinds of context in one workflow today, and the tighter two-way connection between them is the part still being built.
The Agent Harness: Context, Execution, Governance
In this post, the agent harness is the architecture around a coding agent that makes autonomous work dependable. It has three layers:
Context gives agents the codebase, delivery, and runtime information they need to understand the task
Execution coordinates agents and tools to plan, implement, review, test, and deliver a change through repeatable workflows
Governance enforces access controls, security policies, approvals, and audit trails so autonomous work stays inside agreed limits
The Harness agents supply the governance layer. An agent's recommendation grants no permission to deploy: policy as code, RBAC, approval gates, continuous verification, and rollback decide what proceeds. In the governed software factory design , pipeline work runs on Delegates inside the customer's own network, and a Zero Trust Service checks every task before it runs. Shell scripts are parsed into an AST and checked against a command allowlist, container images against a registry allowlist, and each step against the resolved pipeline YAML. The service fails to close, so nothing executes if it is unreachable.
AI code review and governed pipelines answer different questions. Review tells you whether a diff looks correct. A governed pipeline decides whether the change may ship and checks how it behaved once it did.
A critical CVE in a shared library shows how the two kinds of context combine. The steps follow the flow Harness describes in its governed software factory post.
The diagram shows four stages, and steps 5 and 6 below are the return path through the Knowledge Graph. Solid arrows show a change moving to production and its results landing in the Software Delivery Knowledge Graph. The dashed arrow is the planned link that lets delivery context inform the next change.
Scope in code. The Security Testing Agent's scan names the dependency and the vulnerable version. An Expert uses the Code Context Engine to find where the dependency is imported, which functions call it, and which tests cover it.
Scope in production. The Software Delivery Knowledge Graph shows which affected services run in production, which are business-critical, and how much SLO budget each has left. The same library in an internal utility and in a customer-facing authentication service gets different priority.
Fix and review. Agents propose the upgrade, update code and tests, and open a pull request in each affected repository. Reviewers and the Verifier check correctness, security, and maintainability.
Gate. The Software Delivery Agent's pipeline applies policy and approval gates by the service's risk tier. A change that fails a mandatory check does not proceed.
Verify in production. Continuous Verification compares each deployment with a baseline captured when the deployment starts. If behavior deviates, the pipeline rolls the deployment back with the anomalou

[truncated]
