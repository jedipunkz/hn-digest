---
source: "https://mirevox.dev/"
hn_url: "https://news.ycombinator.com/item?id=50594796"
title: "Modular workflow governance and evals for AI-DLC"
article_title: "Mirevox — Agentic control plane modules"
image: ""
author: "ssadruddin"
captured_at: "2026-11-22T00:12:40Z"
capture_tool: "hn-digest"
hn_id: 50594796
score: 1
comments: 0
posted_at: "2026-11-21T23:48:33Z"
tags:
  - hacker-news
---

# Modular workflow governance and evals for AI-DLC

- HN: [50594796](https://news.ycombinator.com/item?id=50594796)
- Source: [mirevox.dev](https://mirevox.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-11-21T23:48:33Z

## Translation

Title: Modular workflow governance and evals for AI-DLC
Article title: Mirevox — Agentic control plane modules
Description: Open-source, self-hostable modules for routing, English governance, AI-DLC stage gates, shadow versioning, shared context, and AI egress evidence—without competing on foundation models.

Article text:
Mirevox — Agentic control plane modules Mirevox Stack AI-DLC Router Evals Stage Gate Versions Context Lineage AI-DORA GitHub (opens in a new tab) Stack AI-DLC Router Evals Stage Gate Versions Context Lineage AI-DORA request route eval egress project_id stage history ai-dora Mirevox
A self-hosted control plane for AI-assisted delivery.
It sits on top of the models and harnesses you already use. It doesn't try to replace them.
Status: all seven modules and the control-plane hub ship runnable v0.1 slices with Docker Compose and unit tests.
The tools already exist. The problem is the gaps between them.
If you've wired up AI tooling for a team, you already know the pieces. LiteLLM or Portkey handle routing. Langfuse, LangSmith or Braintrust handle traces and evals. LangGraph handles orchestration, and OpenMetadata or Atlan handle lineage. Each one is good at its own job. None of them agree on what a project is, which request belongs to which stage, or where the data went after it left the router.
So you end up writing glue. You have keys in three places, evals in four dashboards, and a spreadsheet saying which model is allowed to see which data. Mirevox is meant to be that glue, written once, with contracts that travel well enough that you can swap out any of the tools underneath. Apache-2.0. Every module repo and this site ship with Docker Compose and runnable v0.1 slices.
Models are at the bottom: Foundry, Bedrock, Vertex, or whatever you run locally. Harnesses sit above them—the workspace around the model, like VS Code with an agent or a chat UI. Mirevox is the layer above the harness. It's the part a non-technical person needs to get work through, and the part an operator needs to see what happened. It isn't another terminal harness.
Not Mirevox. Hyperscalers and local runtimes own this layer.
Workspace around the model—tools, retries, memory. We sit on top.
route → eval / egress evidence, joined by request / trace id.
project → stage → cross-tool history, plus a shared multi-repo index.
Delivery & durability observability
code turnover, PR review tax, AI attribution share, and complexity-adjusted throughput.
Optional showcase connectivity—modules stay standalone, not a monolith.
Built for the whole AI development lifecycle, not just the model call
Most AI tooling stops at the request. Real delivery takes longer than that. Someone has an idea, it turns into requirements, then a design decision, then a spec, then code, then tests. The modules are designed to follow that whole path: high-level requirements, then functional and non-functional requirements, then ADRs, then specs. At each step you get a clear answer to "can I move on yet?" SpecKit or OpenSpec can plug in, but you don't need them.
01 Pre-requirements Intake, constraints, and whether the work should start at all.
02 HLR High-level requirements in language reviewers can argue with.
03 FR / NFR Functional and non-functional detail—enough to build against.
04 ADR Architecture decisions recorded before code pretends they were obvious.
05 Spec / plan SpecKit, OpenSpec, or your own template—Stage Gate does not lock the method.
06 Build Harness work with a clear gate: what evidence advances the stage.
07 Verify Tests, reviews, and sign-off tied back to the same project id.
Stages are a metaphor borrowed from ordinary workflows (book a slot, meet the gate, move on)—applied to AI delivery. Advance criteria stay readable; SpecKit and OpenSpec are options, not the only path.
Nearly everything here falls into one of two paths. Both can lean on the Context Engine, so each repo gets indexed once and shared, instead of being re-ingested every time an agent starts up.
A request comes in from VS Code or the home chatbot, both pointing at the same endpoint. Smart Router decides where it goes—local model or provider—based on the task and what you've said is allowed. On the way through, Evals & English Governance checks the conversation against rules you wrote in plain English, and Data Lineage records what data left and where it went.
Someone starts a project and gets a project ID. That ID ties together Confluence or SharePoint pages, the GitHub or Bitbucket repo, and the test results. Stage Gate shows which stage they're in and what's still needed. Shadow Versioning keeps the history across those tools so nobody has to learn what a rebase is just to see what changed. Context Engine is the shared index underneath, while AI-DORA measures post-merge code durability, review tax, and throughput across repos.
Foundation models are a capital game, and the hyperscalers are winning it. Competing there would waste everyone's time. The same goes for cloud Terraform portals, which were the 2020 problem.
What's left is the unglamorous part: routing between models, checking what they do, and keeping track of projects and data. The platforms will probably absorb it eventually. Until then it's useful, so it's here, open source, for anyone to run or build a business on.
There is a control plane (opens in a new tab) repo, but it's mostly meant to bring the modules up together and connect them. Each module is its own repo with its own contract. You can run all seven or just the router. Nothing is hosted for you; it runs on your own machine.
There are seven. Each one has its own page saying what it's meant to do and how far along it is. Check the repo before you depend on anything.
Smart Router runnable Model access gets scattered across provider keys, endpoints, deployment rules, and separate UIs. Even a VS Code assistant and a home chatbot can end up repeating the same integration work.
Evals & English Governance runnable Each model platform has its own eval screens, scoring terms, and policy controls. Teams end up comparing results by hand, while governance rules live in documents that do not map cleanly to provider settings.
Stage Gate runnable AI-assisted projects can jump from an appointment straight into generated code. The missing middle—prerequisites, requirements, decisions, and clear advance criteria—is usually spread across documents and understood by only a few people.
Shadow Versioning runnable One project can have its brief in Confluence, approvals in SharePoint, code in Git, and evidence in a testing system. Humans see one piece of work, but the tools expose unrelated histories and technical version language.
Context Engine runnable An agent working across 15 repositories can spend more tokens rediscovering the codebase than answering the question. Every tool builds its own partial picture, then throws much of that work away.
Data Lineage runnable Once a prompt crosses a gateway, it can be hard to explain which provider, region, model, or internal service handled it. Existing lineage tools cover data estates well, but AI request paths often need a smaller, trace-linked view.
AI-DORA runnable AI coding agents inflate commit counts and PR volume, hiding code churn, review bottlenecks, and rework. Traditional DORA dashboards track deployment frequency without telling you if half the AI-written lines were reverted two weeks later.
Open-source modules for the agentic control plane—use them, adapt them, commercialize them. Licensed under Apache-2.0. Not a foundation model company.

## Original Extract

Open-source, self-hostable modules for routing, English governance, AI-DLC stage gates, shadow versioning, shared context, and AI egress evidence—without competing on foundation models.

Mirevox — Agentic control plane modules Mirevox Stack AI-DLC Router Evals Stage Gate Versions Context Lineage AI-DORA GitHub (opens in a new tab) Stack AI-DLC Router Evals Stage Gate Versions Context Lineage AI-DORA request route eval egress project_id stage history ai-dora Mirevox
A self-hosted control plane for AI-assisted delivery.
It sits on top of the models and harnesses you already use. It doesn't try to replace them.
Status: all seven modules and the control-plane hub ship runnable v0.1 slices with Docker Compose and unit tests.
The tools already exist. The problem is the gaps between them.
If you've wired up AI tooling for a team, you already know the pieces. LiteLLM or Portkey handle routing. Langfuse, LangSmith or Braintrust handle traces and evals. LangGraph handles orchestration, and OpenMetadata or Atlan handle lineage. Each one is good at its own job. None of them agree on what a project is, which request belongs to which stage, or where the data went after it left the router.
So you end up writing glue. You have keys in three places, evals in four dashboards, and a spreadsheet saying which model is allowed to see which data. Mirevox is meant to be that glue, written once, with contracts that travel well enough that you can swap out any of the tools underneath. Apache-2.0. Every module repo and this site ship with Docker Compose and runnable v0.1 slices.
Models are at the bottom: Foundry, Bedrock, Vertex, or whatever you run locally. Harnesses sit above them—the workspace around the model, like VS Code with an agent or a chat UI. Mirevox is the layer above the harness. It's the part a non-technical person needs to get work through, and the part an operator needs to see what happened. It isn't another terminal harness.
Not Mirevox. Hyperscalers and local runtimes own this layer.
Workspace around the model—tools, retries, memory. We sit on top.
route → eval / egress evidence, joined by request / trace id.
project → stage → cross-tool history, plus a shared multi-repo index.
Delivery & durability observability
code turnover, PR review tax, AI attribution share, and complexity-adjusted throughput.
Optional showcase connectivity—modules stay standalone, not a monolith.
Built for the whole AI development lifecycle, not just the model call
Most AI tooling stops at the request. Real delivery takes longer than that. Someone has an idea, it turns into requirements, then a design decision, then a spec, then code, then tests. The modules are designed to follow that whole path: high-level requirements, then functional and non-functional requirements, then ADRs, then specs. At each step you get a clear answer to "can I move on yet?" SpecKit or OpenSpec can plug in, but you don't need them.
01 Pre-requirements Intake, constraints, and whether the work should start at all.
02 HLR High-level requirements in language reviewers can argue with.
03 FR / NFR Functional and non-functional detail—enough to build against.
04 ADR Architecture decisions recorded before code pretends they were obvious.
05 Spec / plan SpecKit, OpenSpec, or your own template—Stage Gate does not lock the method.
06 Build Harness work with a clear gate: what evidence advances the stage.
07 Verify Tests, reviews, and sign-off tied back to the same project id.
Stages are a metaphor borrowed from ordinary workflows (book a slot, meet the gate, move on)—applied to AI delivery. Advance criteria stay readable; SpecKit and OpenSpec are options, not the only path.
Nearly everything here falls into one of two paths. Both can lean on the Context Engine, so each repo gets indexed once and shared, instead of being re-ingested every time an agent starts up.
A request comes in from VS Code or the home chatbot, both pointing at the same endpoint. Smart Router decides where it goes—local model or provider—based on the task and what you've said is allowed. On the way through, Evals & English Governance checks the conversation against rules you wrote in plain English, and Data Lineage records what data left and where it went.
Someone starts a project and gets a project ID. That ID ties together Confluence or SharePoint pages, the GitHub or Bitbucket repo, and the test results. Stage Gate shows which stage they're in and what's still needed. Shadow Versioning keeps the history across those tools so nobody has to learn what a rebase is just to see what changed. Context Engine is the shared index underneath, while AI-DORA measures post-merge code durability, review tax, and throughput across repos.
Foundation models are a capital game, and the hyperscalers are winning it. Competing there would waste everyone's time. The same goes for cloud Terraform portals, which were the 2020 problem.
What's left is the unglamorous part: routing between models, checking what they do, and keeping track of projects and data. The platforms will probably absorb it eventually. Until then it's useful, so it's here, open source, for anyone to run or build a business on.
There is a control plane (opens in a new tab) repo, but it's mostly meant to bring the modules up together and connect them. Each module is its own repo with its own contract. You can run all seven or just the router. Nothing is hosted for you; it runs on your own machine.
There are seven. Each one has its own page saying what it's meant to do and how far along it is. Check the repo before you depend on anything.
Smart Router runnable Model access gets scattered across provider keys, endpoints, deployment rules, and separate UIs. Even a VS Code assistant and a home chatbot can end up repeating the same integration work.
Evals & English Governance runnable Each model platform has its own eval screens, scoring terms, and policy controls. Teams end up comparing results by hand, while governance rules live in documents that do not map cleanly to provider settings.
Stage Gate runnable AI-assisted projects can jump from an appointment straight into generated code. The missing middle—prerequisites, requirements, decisions, and clear advance criteria—is usually spread across documents and understood by only a few people.
Shadow Versioning runnable One project can have its brief in Confluence, approvals in SharePoint, code in Git, and evidence in a testing system. Humans see one piece of work, but the tools expose unrelated histories and technical version language.
Context Engine runnable An agent working across 15 repositories can spend more tokens rediscovering the codebase than answering the question. Every tool builds its own partial picture, then throws much of that work away.
Data Lineage runnable Once a prompt crosses a gateway, it can be hard to explain which provider, region, model, or internal service handled it. Existing lineage tools cover data estates well, but AI request paths often need a smaller, trace-linked view.
AI-DORA runnable AI coding agents inflate commit counts and PR volume, hiding code churn, review bottlenecks, and rework. Traditional DORA dashboards track deployment frequency without telling you if half the AI-written lines were reverted two weeks later.
Open-source modules for the agentic control plane—use them, adapt them, commercialize them. Licensed under Apache-2.0. Not a foundation model company.
