---
source: "https://banes-lab.com/"
hn_url: "https://news.ycombinator.com/item?id=49861827"
title: "Practical lessons on building software with LLMs and automated governance"
article_title: "Bane's Lab — Structured collaboration with LLMs"
image: "https://banes-lab.com/assets/social/home-w600-h315.og.generated.gif"
author: "jay-baleine"
captured_at: "2026-09-27T00:17:55Z"
capture_tool: "hn-digest"
hn_id: 49861827
score: 1
comments: 0
posted_at: "2026-09-27T00:00:06Z"
tags:
  - hacker-news
---

# Practical lessons on building software with LLMs and automated governance

- HN: [49861827](https://news.ycombinator.com/item?id=49861827)
- Source: [banes-lab.com](https://banes-lab.com/)
- Score: 1
- Comments: 0
- Posted: 2026-09-27T00:00:06Z

## Translation

Title: Practical lessons on building software with LLMs and automated governance
Article title: Bane's Lab — Structured collaboration with LLMs
Description: Bane

Article text:
Structured collaboration with LLMs
Bane's Lab documents the method I use to build software with LLMs, together with the instruction format, the software architecture and the ontology that the method relies on. I have used it since 2024, on every project I build, this site included. The method does not make a model's output correct. What it does is make incorrect output visible and refuse it, through checks that run on every change rather than through instructions the model is asked to remember.
The site is organized around the methodology, which is one page of six tabs read in order, from Start to Ship. The grammar, architecture and ontology pages supply material that particular sections of the methodology rely on, and the visualization shows where each of their sections enters that order. The anatomy page works in the other direction: it shows the source of this site, and its sections link back to the chapters that describe it.
 Start 01 - The loop 02 - Who does what 03 - The stance 04 - Adversarial by default 05 - Resolving a message 06 - Three encodings 07 - Where a rule lives 08 - Rules with names 09 - A seat is a contract 10 - The behaviour document 11 - The drop-in Introduction 01 - What PAG is 02 - Why it works 03 - PAG and the method Guide 01 - Writing a first document 02 - Document structure 03 - Semantic operations 04 - Node design 05 - Writing constraints 06 - Well-formedness  Plan 12 - Worth before work 13 - The plan is a graph 14 - Execute the template 15 - Ask where it appears 16 - When rules collide Patterns 01 - Instruction patterns 02 - From intent to structure 03 - Genesis stages 04 - Algorithm examples 05 - Integrating algorithms  Build 17 - Detect, log, fix 18 - The gate holds the line 19 - The check comes first 20 - A check matches a shape 21 - Tools live in the tree 22 - One home 23 - The filesystem is the architecture 24 - Fail at the boundary 25 - Placement is a grammar  Verify 26 - It looked right 27 - Verify the verifier 28 - A report, not a checkbox 29 - Unknown is not pass 30 - One correct answer 31 - Derived state 32 - Counting copies 33 - Documentation is code 34 - Moves and renames 35 - Coverage is derived Validation 01 - Validation gates 02 - Limits  Model 01 - A system is a graph 02 - Definitions own what, code owns how 03 - The layer spine 04 - The direction axis  Principles 01 - Principles are typed 02 - Every record has a kind 03 - The canon is grouped twice 04 - Computation and resource 05 - Execution joins the halves 06 - The structural domain 07 - A tension has a mechanism 08 - Separate, trade, or mitigate  Decay 01 - An anti-pattern is a decay path 02 - Seven controls, seven classes 03 - Never and always 04 - Debt and leverage  Coverage 01 - From intent to predicate 02 - What can drift, seen through how it drifts 03 - A cell that resists an invariant 04 - The honest gaps  Glossary 01 - The principle architecture 02 - Architectural rules and p
[truncated]
The method itself comes in six parts that follow the order of the work, from starting a project to shipping it. Each section describes one practice, how it fails when it is missing, and a test you can run against your own work to check it.
Engineering leads · Software architects · LLM-assisted developers
Pattern Abstract Grammar is a structured format for writing instructions to a model. A document declares what type of instruction it is, draws its verbs from a closed vocabulary and ends every step on a gate with checkable evidence. The page covers the grammar, its validation rules and a set of templates.
Software architects · LLM-assisted developers · Tool authors
Architecture Principles, decay and coverage
The page covers software architecture for systems in which a model writes much of the code. It models a system as a graph, holds each principle as a typed record, traces each anti-pattern back to the control whose absence caused it, and derives coverage from a grid rather than from a count.
Engineering leads · Software architects · LLM-assisted developers
Ontology Principles, lexicon and algorithms
The ontology is the data behind the architecture page, and the site's own checks read it too. It holds every principle with its relations and repairs, every defined term, every algorithm contract, and how each tension between two principles is resolved.
Software architects · LLM-assisted developers · Tool authors
Anatomy Source trees, walks and definitions
The anatomy page shows the client source of this site, parsed on every build. Each file is shown with its syntax walk, its definitions and the calls between them, together with the diagnoses the parser ran over the whole tree.
Software architects · LLM-assisted developers · Tool authors
FAQ Origin, practice and limits
Here I answer the questions I am asked most often about the methodology, including where it came from and where it stops being useful.
Each page is also published as Markdown and as JSON, and you can query them via the API routes.
/json/api · /api.md · /llms.txt · /llms-full.txt
Model training is permitted; credit Jay Baleine and cite this site.

## Original Extract

Bane

Structured collaboration with LLMs
Bane's Lab documents the method I use to build software with LLMs, together with the instruction format, the software architecture and the ontology that the method relies on. I have used it since 2024, on every project I build, this site included. The method does not make a model's output correct. What it does is make incorrect output visible and refuse it, through checks that run on every change rather than through instructions the model is asked to remember.
The site is organized around the methodology, which is one page of six tabs read in order, from Start to Ship. The grammar, architecture and ontology pages supply material that particular sections of the methodology rely on, and the visualization shows where each of their sections enters that order. The anatomy page works in the other direction: it shows the source of this site, and its sections link back to the chapters that describe it.
 Start 01 - The loop 02 - Who does what 03 - The stance 04 - Adversarial by default 05 - Resolving a message 06 - Three encodings 07 - Where a rule lives 08 - Rules with names 09 - A seat is a contract 10 - The behaviour document 11 - The drop-in Introduction 01 - What PAG is 02 - Why it works 03 - PAG and the method Guide 01 - Writing a first document 02 - Document structure 03 - Semantic operations 04 - Node design 05 - Writing constraints 06 - Well-formedness  Plan 12 - Worth before work 13 - The plan is a graph 14 - Execute the template 15 - Ask where it appears 16 - When rules collide Patterns 01 - Instruction patterns 02 - From intent to structure 03 - Genesis stages 04 - Algorithm examples 05 - Integrating algorithms  Build 17 - Detect, log, fix 18 - The gate holds the line 19 - The check comes first 20 - A check matches a shape 21 - Tools live in the tree 22 - One home 23 - The filesystem is the architecture 24 - Fail at the boundary 25 - Placement is a grammar  Verify 26 - It looked right 27 - Verify the verifier 28 - A report, not a checkbox 29 - Unknown is not pass 30 - One correct answer 31 - Derived state 32 - Counting copies 33 - Documentation is code 34 - Moves and renames 35 - Coverage is derived Validation 01 - Validation gates 02 - Limits  Model 01 - A system is a graph 02 - Definitions own what, code owns how 03 - The layer spine 04 - The direction axis  Principles 01 - Principles are typed 02 - Every record has a kind 03 - The canon is grouped twice 04 - Computation and resource 05 - Execution joins the halves 06 - The structural domain 07 - A tension has a mechanism 08 - Separate, trade, or mitigate  Decay 01 - An anti-pattern is a decay path 02 - Seven controls, seven classes 03 - Never and always 04 - Debt and leverage  Coverage 01 - From intent to predicate 02 - What can drift, seen through how it drifts 03 - A cell that resists an invariant 04 - The honest gaps  Glossary 01 - The principle architecture 02 - Architectural rules and p
[truncated]
The method itself comes in six parts that follow the order of the work, from starting a project to shipping it. Each section describes one practice, how it fails when it is missing, and a test you can run against your own work to check it.
Engineering leads · Software architects · LLM-assisted developers
Pattern Abstract Grammar is a structured format for writing instructions to a model. A document declares what type of instruction it is, draws its verbs from a closed vocabulary and ends every step on a gate with checkable evidence. The page covers the grammar, its validation rules and a set of templates.
Software architects · LLM-assisted developers · Tool authors
Architecture Principles, decay and coverage
The page covers software architecture for systems in which a model writes much of the code. It models a system as a graph, holds each principle as a typed record, traces each anti-pattern back to the control whose absence caused it, and derives coverage from a grid rather than from a count.
Engineering leads · Software architects · LLM-assisted developers
Ontology Principles, lexicon and algorithms
The ontology is the data behind the architecture page, and the site's own checks read it too. It holds every principle with its relations and repairs, every defined term, every algorithm contract, and how each tension between two principles is resolved.
Software architects · LLM-assisted developers · Tool authors
Anatomy Source trees, walks and definitions
The anatomy page shows the client source of this site, parsed on every build. Each file is shown with its syntax walk, its definitions and the calls between them, together with the diagnoses the parser ran over the whole tree.
Software architects · LLM-assisted developers · Tool authors
FAQ Origin, practice and limits
Here I answer the questions I am asked most often about the methodology, including where it came from and where it stops being useful.
Each page is also published as Markdown and as JSON, and you can query them via the API routes.
/json/api · /api.md · /llms.txt · /llms-full.txt
Model training is permitted; credit Jay Baleine and cite this site.
