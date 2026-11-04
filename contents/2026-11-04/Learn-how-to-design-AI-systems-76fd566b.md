---
source: "https://github.com/daneb/ai-system-design-primer/"
hn_url: "https://news.ycombinator.com/item?id=50367602"
title: "Learn how to design AI systems"
article_title: "GitHub - daneb/ai-system-design-primer: Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards. · GitHub"
image: "https://opengraph.githubassets.com/e47ed9ab9fea74c74b4587839b998c2c91594b8b0f170e8bfb7f0de71a4594e4/daneb/ai-system-design-primer"
author: "danebalia"
captured_at: "2026-11-04T16:54:29Z"
capture_tool: "hn-digest"
hn_id: 50367602
score: 1
comments: 1
posted_at: "2026-11-04T16:08:10Z"
tags:
  - hacker-news
---

# Learn how to design AI systems

- HN: [50367602](https://news.ycombinator.com/item?id=50367602)
- Source: [github.com](https://github.com/daneb/ai-system-design-primer/)
- Score: 1
- Comments: 1
- Posted: 2026-11-04T16:08:10Z

## Translation

Title: Learn how to design AI systems
Article title: GitHub - daneb/ai-system-design-primer: Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards. · GitHub
Description: Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards. - daneb/ai-system-design-primer

Article text:
GitHub - daneb/ai-system-design-primer: Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
daneb
ai-system-design-primer Public Notifications You must be signed in to change notification settings
Star 6 ( 6 ) You must be signed in to star a repository
Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards.
Readme CC-BY-4.0, MIT licenses found Code of conduct
Security policy Activity Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
33 Commits 33 Commits Folders and files
.github .github adr adr docs docs images images labs labs research research scripts scripts templates templates .gitignore .gitignore CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSE-CODE LICENSE-CODE PLAN.md PLAN.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md STYLE.md STYLE.md lychee.toml lychee.toml View all files Repository files navigation
Learn how to design systems built on large language models: the trade-offs, the components, and the
patterns that have converged across the industry, along with the ones that are still moving.
If you are…
Read
New to building on models
Orientation and the foundations 0.1 – 0.3 , then the trade-offs in Part 2
Preparing for an interview or design review
The method , two solved designs , then the practice questions ; pace yourself with the study guide
Building a system now
The component you need and its deep dive, the numbers and the cloud mapping
Reviewing risk or security
2.4 Capability vs security , 3.2 , 3.12 , 3.17 and the governance crosswalk
Motivation
Building on models is now a core engineering skill, and the field moves fast. Most of what's written
about it is a product announcement, a framework tutorial or a single team's experience. This primer is
an organised collection of what has converged across providers and researchers, with every trade-off
named and every number dated and sourced.
It is modelled on Donne Martin's System Design Primer :
the same arc of method, trade-offs, components, solved designs and reference numbers, rebuilt for
systems whose core component is probabilistic, priced per token, and open to prompt injection.
Learn how to design AI systems. Each topic is a short page you can read in ten minutes, with a
diagram, when to use it, when not to, and its disadvantages. A deep dive per topic holds the provider
survey, numbers, risk controls and cloud mappings.
Prepare for the AI system design conversation, in an interview or a design review. Work through the
method , the solved designs
and the practice questions .
Every page ends in question-and-answer flashcards. Build the deck with either command (standard
library only):
python3 scripts/primer.py anki # writes build/ai-system-design-primer.apkg: double-click to import
python3 scripts/primer.py flashcards # writes build/flashcards.tsv: plain text for Anki's File > Import
The package has one sub-deck per part. Re-importing a newer build updates answers in place and keeps
your review history; a reworded question arrives as a new card. Cards are tagged by section and
maturity, so you can study one part, or suspend maturity::frontier cards when they go stale.
Most topics come with a small, runnable lab: a simulation or calculation, numpy or the standard
library only, with the prediction written down before the run and an honest note when it was wrong.
Pull requests are welcome to fix errors, improve sections and add new ones. Read STYLE.md
for how sections are researched and written (learner-first, several providers per topic, every claim
sourced) and CONTRIBUTING.md for the mechanics. Decisions about the primer itself
are recorded in adr/ . Work still to do is in ROADMAP.md .
Index of AI system design topics
Summaries of AI system design topics, including pros and cons. Everything is a trade-off.
Each section links to its deep dive. Maturity is marked S (stable core), S/F (stable concept,
volatile specifics) or F (frontier, expect rewrites).
0.1 Tokens, context and sampling S
0.2 Inside a model call: prefill, decode and reasoning S
0.3 Tool calling and structured output S
1 How to approach an AI system design question S
2.1 Cost vs quality vs latency S · deep dive
2.2 Determinism vs autonomy S/F · deep dive
2.3 Context stuffing vs retrieval S/F · deep dive
2.4 Capability vs security S/F · deep dive
2.5 Single-agent vs multi-agent S/F · deep dive
2.6 Prompting vs RAG vs fine-tuning S · deep dive
2.7 Throughput vs per-request latency in serving S · deep dive
Part 3: Components, in request order
3.1 AI gateway and router S · deep dive
3.2 Input guardrails S · deep dive
3.3 Context engineering S/F · deep dive
3.4 Retrieval pipeline S · deep dive
3.5 Agentic retrieval and the long-context boundary F · deep dive
3.6 Vector stores S/F · deep dive
3.7 Model layer: selection, structured outputs, streaming S/F · deep dive
3.8 Serving internals S/F · deep dive
3.9 Workflows and agents S · deep dive
3.10 Multi-agent topologies and protocols (MCP, A2A) F · deep dive
3.11 Agent memory F · deep dive
3.12 Output guardrails and validation S · deep dive
3.14 Asynchronism: batch, durable execution, back-pressure S · deep dive
3.16 Observability S/F · deep dive
3.18 Model adaptation S · deep dive
4.1 Design an LLM inference API S
4.2 Design an enterprise RAG assistant S/F · deep dive
4.3 Design an enterprise AI gateway S
4.4 Design a regulated document-review assistant S/F
4.5 Design a customer-support agent F
4.6 Design a coding agent platform F
4.7 Design an eval and release pipeline S
4.8 Additional design questions S/F
5.1 Numbers every AI engineer should know S/F
5.2 Cloud mapping: Azure, AWS, Google Cloud F
5.4 Governance crosswalk: OWASP, FINOS, EU AI Act F
5.5 Real-world architectures and engineering blogs F
Suggested topics to review based on your timeline.
How to approach an AI system design question
How to tackle an AI system design question, in an interview or a design review.
The classic four steps, with two changes: the quality bar comes first , because a model-backed
feature is right only some of the time; and hardening is its own step , because prompt injection,
data leakage and runaway agents are design constraints, not add-ons.
Use cases, constraints and the quality bar. Who, what, how many, how fast, what data, and how
you'll know it works.
Back-of-the-envelope. Tokens per request, requests per second, cost per day, latency.
High-level design on the complexity ladder. Single call → workflow → single agent →
multi-agent. Start at the lowest rung that meets the bar.
Core components. Retrieval, context, tools, model choice, output checks.
Harden. Which legs of the lethal trifecta are present, and which does the design break?
Scale and optimise. Bottlenecks in order; caching, routing, batching.
Full walkthrough with a worked example: How to approach an AI system design question .
AI system design questions with solutions
Common AI system design questions with sample discussions, numbers and diagrams.
More practice: additional design questions .
Numbers every AI engineer should know : tokens, speed, cost
ratios, memory, retrieval and eval figures, each dated and sourced.
Cloud mapping : each capability on Azure, AWS, Google Cloud,
OpenAI, Anthropic and open source.
Pattern crosswalk : the same patterns under each
provider's and author's names.
Governance crosswalk : OWASP LLM and Agentic Top 10,
FINOS, EU AI Act, mapped to primer sections.
Real-world architectures : engineering write-ups
and independent studies.
How this primer differs from the original, and from a blog post
Trade-offs first. Every pattern says which trade-off it moves and in which direction.
Quality bar first. Evals are part of the method, not an afterthought.
Maturity is labelled and dated. Each page carries last_verified and review_by ; the checker
fails when a review is overdue. Frontier pages are reviewed every three months.
Several providers on every topic. Anthropic, OpenAI, Google, AWS and Microsoft, with the names
each uses.
Vendor-neutral prose. Product names live in the deep dives and the cloud mapping.
Risk and controls on every topic, with the evidence an auditor would ask for.
Runnable labs with predictions recorded before the run.
Part
Sections
Published
Part 0: Orientation
4
4
Part 1: Method
1
1
Part 2: Trade-offs
7
7
Part 3: Components, in request order
18
18
Part 4: Solved designs
8
8
Part 5: Appendix
6
6
See ROADMAP.md for what's next and PLAN.md for the build plan and the
verification debt (claims still waiting for a first-party source).
Donne Martin's System Design Primer , whose
structure this follows.
The provider guides, papers and engineering write-ups cited in each section; see
real-world architectures .
Open an issue to discuss errors, missing topics or new solved designs. Report security problems, or
guidance that would make a system less safe, privately as described in SECURITY.md .
Participation is covered by the Code of Conduct .
Content (pages, diagrams, flashcards, lab READMEs) is licensed under
Creative Commons Attribution 4.0 International (CC BY 4.0), the same licence as the
original System Design Primer.
All diagrams and charts are original to this repository and covered by the content licence.
Code ( scripts/ , the lab scripts and the CI workflow) is licensed under the MIT License .
Contributions are accepted under the same terms. The reasoning is in ADR 0008 .
The author researched and wrote this primer with an AI assistant, and reviewed every section. The
method is in STYLE.md : primary sources from several providers for every claim, a dated
last_verified on every page, runnable labs instead of asserted numbers, and a public list of what
isn't yet verified. Mistakes are still possible; please report them with a source.
Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards.
Readme CC-BY-4.0, MIT licenses found Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards. - daneb/ai-system-design-primer

GitHub - daneb/ai-system-design-primer: Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
daneb
ai-system-design-primer Public Notifications You must be signed in to change notification settings
Star 6 ( 6 ) You must be signed in to star a repository
Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards.
Readme CC-BY-4.0, MIT licenses found Code of conduct
Security policy Activity Stars
0 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
33 Commits 33 Commits Folders and files
.github .github adr adr docs docs images images labs labs research research scripts scripts templates templates .gitignore .gitignore CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSE-CODE LICENSE-CODE PLAN.md PLAN.md README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md STYLE.md STYLE.md lychee.toml lychee.toml View all files Repository files navigation
Learn how to design systems built on large language models: the trade-offs, the components, and the
patterns that have converged across the industry, along with the ones that are still moving.
If you are…
Read
New to building on models
Orientation and the foundations 0.1 – 0.3 , then the trade-offs in Part 2
Preparing for an interview or design review
The method , two solved designs , then the practice questions ; pace yourself with the study guide
Building a system now
The component you need and its deep dive, the numbers and the cloud mapping
Reviewing risk or security
2.4 Capability vs security , 3.2 , 3.12 , 3.17 and the governance crosswalk
Motivation
Building on models is now a core engineering skill, and the field moves fast. Most of what's written
about it is a product announcement, a framework tutorial or a single team's experience. This primer is
an organised collection of what has converged across providers and researchers, with every trade-off
named and every number dated and sourced.
It is modelled on Donne Martin's System Design Primer :
the same arc of method, trade-offs, components, solved designs and reference numbers, rebuilt for
systems whose core component is probabilistic, priced per token, and open to prompt injection.
Learn how to design AI systems. Each topic is a short page you can read in ten minutes, with a
diagram, when to use it, when not to, and its disadvantages. A deep dive per topic holds the provider
survey, numbers, risk controls and cloud mappings.
Prepare for the AI system design conversation, in an interview or a design review. Work through the
method , the solved designs
and the practice questions .
Every page ends in question-and-answer flashcards. Build the deck with either command (standard
library only):
python3 scripts/primer.py anki # writes build/ai-system-design-primer.apkg: double-click to import
python3 scripts/primer.py flashcards # writes build/flashcards.tsv: plain text for Anki's File > Import
The package has one sub-deck per part. Re-importing a newer build updates answers in place and keeps
your review history; a reworded question arrives as a new card. Cards are tagged by section and
maturity, so you can study one part, or suspend maturity::frontier cards when they go stale.
Most topics come with a small, runnable lab: a simulation or calculation, numpy or the standard
library only, with the prediction written down before the run and an honest note when it was wrong.
Pull requests are welcome to fix errors, improve sections and add new ones. Read STYLE.md
for how sections are researched and written (learner-first, several providers per topic, every claim
sourced) and CONTRIBUTING.md for the mechanics. Decisions about the primer itself
are recorded in adr/ . Work still to do is in ROADMAP.md .
Index of AI system design topics
Summaries of AI system design topics, including pros and cons. Everything is a trade-off.
Each section links to its deep dive. Maturity is marked S (stable core), S/F (stable concept,
volatile specifics) or F (frontier, expect rewrites).
0.1 Tokens, context and sampling S
0.2 Inside a model call: prefill, decode and reasoning S
0.3 Tool calling and structured output S
1 How to approach an AI system design question S
2.1 Cost vs quality vs latency S · deep dive
2.2 Determinism vs autonomy S/F · deep dive
2.3 Context stuffing vs retrieval S/F · deep dive
2.4 Capability vs security S/F · deep dive
2.5 Single-agent vs multi-agent S/F · deep dive
2.6 Prompting vs RAG vs fine-tuning S · deep dive
2.7 Throughput vs per-request latency in serving S · deep dive
Part 3: Components, in request order
3.1 AI gateway and router S · deep dive
3.2 Input guardrails S · deep dive
3.3 Context engineering S/F · deep dive
3.4 Retrieval pipeline S · deep dive
3.5 Agentic retrieval and the long-context boundary F · deep dive
3.6 Vector stores S/F · deep dive
3.7 Model layer: selection, structured outputs, streaming S/F · deep dive
3.8 Serving internals S/F · deep dive
3.9 Workflows and agents S · deep dive
3.10 Multi-agent topologies and protocols (MCP, A2A) F · deep dive
3.11 Agent memory F · deep dive
3.12 Output guardrails and validation S · deep dive
3.14 Asynchronism: batch, durable execution, back-pressure S · deep dive
3.16 Observability S/F · deep dive
3.18 Model adaptation S · deep dive
4.1 Design an LLM inference API S
4.2 Design an enterprise RAG assistant S/F · deep dive
4.3 Design an enterprise AI gateway S
4.4 Design a regulated document-review assistant S/F
4.5 Design a customer-support agent F
4.6 Design a coding agent platform F
4.7 Design an eval and release pipeline S
4.8 Additional design questions S/F
5.1 Numbers every AI engineer should know S/F
5.2 Cloud mapping: Azure, AWS, Google Cloud F
5.4 Governance crosswalk: OWASP, FINOS, EU AI Act F
5.5 Real-world architectures and engineering blogs F
Suggested topics to review based on your timeline.
How to approach an AI system design question
How to tackle an AI system design question, in an interview or a design review.
The classic four steps, with two changes: the quality bar comes first , because a model-backed
feature is right only some of the time; and hardening is its own step , because prompt injection,
data leakage and runaway agents are design constraints, not add-ons.
Use cases, constraints and the quality bar. Who, what, how many, how fast, what data, and how
you'll know it works.
Back-of-the-envelope. Tokens per request, requests per second, cost per day, latency.
High-level design on the complexity ladder. Single call → workflow → single agent →
multi-agent. Start at the lowest rung that meets the bar.
Core components. Retrieval, context, tools, model choice, output checks.
Harden. Which legs of the lethal trifecta are present, and which does the design break?
Scale and optimise. Bottlenecks in order; caching, routing, batching.
Full walkthrough with a worked example: How to approach an AI system design question .
AI system design questions with solutions
Common AI system design questions with sample discussions, numbers and diagrams.
More practice: additional design questions .
Numbers every AI engineer should know : tokens, speed, cost
ratios, memory, retrieval and eval figures, each dated and sourced.
Cloud mapping : each capability on Azure, AWS, Google Cloud,
OpenAI, Anthropic and open source.
Pattern crosswalk : the same patterns under each
provider's and author's names.
Governance crosswalk : OWASP LLM and Agentic Top 10,
FINOS, EU AI Act, mapped to primer sections.
Real-world architectures : engineering write-ups
and independent studies.
How this primer differs from the original, and from a blog post
Trade-offs first. Every pattern says which trade-off it moves and in which direction.
Quality bar first. Evals are part of the method, not an afterthought.
Maturity is labelled and dated. Each page carries last_verified and review_by ; the checker
fails when a review is overdue. Frontier pages are reviewed every three months.
Several providers on every topic. Anthropic, OpenAI, Google, AWS and Microsoft, with the names
each uses.
Vendor-neutral prose. Product names live in the deep dives and the cloud mapping.
Risk and controls on every topic, with the evidence an auditor would ask for.
Runnable labs with predictions recorded before the run.
Part
Sections
Published
Part 0: Orientation
4
4
Part 1: Method
1
1
Part 2: Trade-offs
7
7
Part 3: Components, in request order
18
18
Part 4: Solved designs
8
8
Part 5: Appendix
6
6
See ROADMAP.md for what's next and PLAN.md for the build plan and the
verification debt (claims still waiting for a first-party source).
Donne Martin's System Design Primer , whose
structure this follows.
The provider guides, papers and engineering write-ups cited in each section; see
real-world architectures .
Open an issue to discuss errors, missing topics or new solved designs. Report security problems, or
guidance that would make a system less safe, privately as described in SECURITY.md .
Participation is covered by the Code of Conduct .
Content (pages, diagrams, flashcards, lab READMEs) is licensed under
Creative Commons Attribution 4.0 International (CC BY 4.0), the same licence as the
original System Design Primer.
All diagrams and charts are original to this repository and covered by the content licence.
Code ( scripts/ , the lab scripts and the CI workflow) is licensed under the MIT License .
Contributions are accepted under the same terms. The reasoning is in ADR 0008 .
The author researched and wrote this primer with an AI assistant, and reviewed every section. The
method is in STYLE.md : primary sources from several providers for every claim, a dated
last_verified on every page, runnable labs instead of asserted numbers, and a public list of what
isn't yet verified. Mistakes are still possible; please report them with a source.
Learn how to design systems built on LLMs. Prep for the system design interview. Includes Anki flashcards.
Readme CC-BY-4.0, MIT licenses found Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
