---
source: "https://github.com/ai-cad-labs/ai-cad"
hn_url: "https://news.ycombinator.com/item?id=49824497"
title: "AI-CAD: An OSS Multi-Agent Harness for Mech. Eng. CAD"
article_title: "GitHub - ai-cad-labs/ai-cad · GitHub"
image: "https://opengraph.githubassets.com/8c7d3ac8a9d007a1a2e3543eaa9d00d7c557bea939ecbf273710db6821a44c8b/ai-cad-labs/ai-cad"
author: "jbm"
captured_at: "2026-09-24T01:08:55Z"
capture_tool: "hn-digest"
hn_id: 49824497
score: 2
comments: 0
posted_at: "2026-09-24T00:14:44Z"
tags:
  - hacker-news
---

# AI-CAD: An OSS Multi-Agent Harness for Mech. Eng. CAD

- HN: [49824497](https://news.ycombinator.com/item?id=49824497)
- Source: [github.com](https://github.com/ai-cad-labs/ai-cad)
- Score: 2
- Comments: 0
- Posted: 2026-09-24T00:14:44Z

## Translation

Title: AI-CAD: An OSS Multi-Agent Harness for Mech. Eng. CAD
Article title: GitHub - ai-cad-labs/ai-cad · GitHub
Description: Contribute to ai-cad-labs/ai-cad development by creating an account on GitHub.

Article text:
GitHub - ai-cad-labs/ai-cad · GitHub
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
ai-cad-labs
/
ai-cad
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.claude .claude .opencode .opencode .shared .shared docs docs frontend frontend logs logs projects projects reference reference rules rules tests tests .gitignore .gitignore .markdownlint.json .markdownlint.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md GEMINI.md GEMINI.md LICENSE LICENSE NOTICE NOTICE README.md README.md pyproject.toml pyproject.toml tools tools uv.lock uv.lock View all files Repository files navigation
Topics: ai-cad · text-to-cad · cadquery · generative-design · multi-agent · design-for-manufacturing
Quickstart |
Runs Gallery |
Case Study |
FAQ |
Docs |
Dev Setup |
Discussions
Issues
AI-CAD : An OSS multi-agent harness for Mech. Eng. CAD
AI-CAD is an autonomous engineering department you commission,
AI-CAD is not a CAD modeling tool you operate.
Give it a plain-English brief or a hand-drawn sketch,
and a team of LLM agents delivers real, manufacturable parts and assemblies:
Plans : the machine and decomposes it into parts
Writes : parametric CadQuery Python code
Renders : every part to engineering views, and looks at them
Critiques : the geometry against design-for-manufacturing and design-for-assembly rules
Repairs : and re-judges until the design survives review and delivers your intent
Agents looking at what they built, and argue about it until it survives review and delivers your intent
All of this interactively inside your coding agent of choice,
and with visibility in a live dashboard .
Every run writes a simple directory of files
that you can read, diff, and version like any other engineering record.
There is no server and no database .
AI-CAD is built on top of your coding-agent .
Your engineering department at work:
parts land as agents deliver, verdicts post,
and the activity waterfall traces every agent in real time.
The system refuses generic manufacturing advice
that contradicts the intent and machine goals.
(See the run that argued back .)
Design machine components as stateful Python code (not immutable mesh props)
Unlike most text-to-3D tools that generate meshes
(excellent for game assets and concept art),
AI-CAD writes parametric CAD code:
every part is a Python program with named dimensions,
exporting STEP and STL that a machine shop or downstream CAD system accepts.
Manufacturability is a hard gate, not an afterthought
A DFM/DFA rulebook scores each design;
repair proposals that worsen the score are automatically rejected
by a regression gate with full traceability.
Agents that see
Every part and assembly is rendered to 8 engineering views
(front / top / right / iso, in wireframe and clean styles).
A vision-capable evaluator agent sees those renders:
geometry is judged from engineering lenses, not just executed.
Harness-agnostic
Canonical agent, skill, and tool content lives once in `.shared/`;
each harness (`.claude/`, `.opencode/`) is a thin symlinked shell over it.
The filesystem is the API
Every run writes a self-describing project directory:
goals, design log, per-part geometry, negotiation records, reflections.
Exports STEP and STL for every part.
No opaque database for facts to hide behind:
the frontend, replay, and analysis all read the same files.
Deterministic tool layer
Geometry execution, rendering, DFMA evaluation, spec validation, dimension checks, and export
are deterministic bash-callable Python tools.
Validation happens independently of any LLM.
Mortality-proof orchestrator
If the orchestrator agent is killed, or runs out of context window,
a fresh one reconstructs all state from the project directory alone.
That is also how a crashed run resumes: no checkpoint database needed.
Runs debrief themselves
A run reaches terminal state by writing a structured `run_reflection.md`:
what failed, what the run improvised to protect the result,
and what would make the next run smarter.
Agentic telemetry ships judgment, not just logs.
The full set of architecture decision records,
with the alternatives each one rejected, lives in
docs/architecture-decision-records/
Eight views per part, judged like a drawing sheet:
this is what the vision evaluator sees before any geometry ships.
One brief in, one self-describing artifact tree out.
This is the real tree of the run in the public gallery, an IC-engine core designed at top-dead-center:
projects/piston_crank/
├── goals.md # the brief, expanded into measurable success criteria
├── design_plan.md # parts, interfaces, the locked dimension chain, build order
├── design_log.md # the agents' step-by-step build log
├── external/ # BOM + sourcing notes for catalog parts
├── assembly/
│ ├── piston/
│ │ ├── part.py # the CadQuery source the agents wrote
│ │ ├── renders/ # the 8 engineering views the evaluator looked at
│ │ └── exports/ # piston.step + piston.stl
│ ├── crankshaft/ # ... same layout per part
│ ├── connecting_rod/
│ ├── gudgeon_pin/
│ ├── piston_ring/
│ └── renders/ # multi-view renders of the assembled mechanism
├── checkpoint.md # optional handover note; resume rebuilds from files alone
└── run_reflection.md # the run's own debrief
Live Dashboard : Watch the machine think
2 hours of autonomous design, traced live:
the agent waterfall above,
the assembly taking shape in four engineering views.
A read-only local dashboard renders live runs:
drawing sheets, assembly negotiations, and a live agent activity waterfall.
cd frontend
npm install
npm run dev
# open http://localhost:5199
The dev server is pinned to port 5199 ( strictPort ) on purpose:
it never imports harness code and never writes;
it only reads projects/ .
Each role is defined by one instructions file
orchestrator - Runs the department. Mortality-proof by design: it can end mid-run and a successor rebuilds everything from files alone.
planner - Decomposes a goal into parts, interfaces, and build order, and marks which parts are parallel_safe to build concurrently.
cad_designer - Writes the CadQuery, adapting proven cookbook patterns instead of improvising raw API calls.
validator - Settles disputes numerically: structural checks and geometry measurements that outrank anyone's opinion, including the vision evaluator's.
dfma_inspector - Reads the renders and scores the design against the DFM/DFA rulebook.
repair - Fixes geometry, but only via part.proposal.py ; the regression gate decides whether the fix ships.
assembly_resolver - Fits the parts together and records the negotiations between them.
sourcing - Selects catalog components and catches catalog traps (a bearing 1 mm thinner than assumed, a nut standard that would overhang its seat).
reviewer - The final ship or no-ship verdict.
The knowledge layer the agents draw on.
Skills are versioned documents ,you can read exactly what the system believes.
cadquery-cookbook - 11 reusable CadQuery patterns, 8 design principles, and 6 anti-patterns; designers adapt the closest pattern rather than composing from raw API calls.
cadquery-anti-hallucination - A catalog of CadQuery methods that do not exist but LLMs keep inventing, each with the correct alternative, plus the runtime error-to-hint table the executor applies.
dfm-rules - The agent-readable digest of the DFM/DFA rulebooks: every rule's id, severity, and fix hint, plus the scoring arithmetic and the rule-proposal schema.
run-reflection - The schema and quality bar for the structured self-debrief every run writes at terminal state.
engineering-handbooks and sourcing-tables - Reserved stubs for machine-design playbooks and component tables, labeled provisional so no agent mistakes an empty slot for authority.
The deterministic half of the system:
bash-callable Python modules ( uv run python -m tools.<name> )
that measure, render, score, and export.
cadquery_executor - Runs generated CadQuery code; a --subprocess flag isolates native OCCT crashes so one bad kernel call cannot take the run down.
renderer - Produces the 8 engineering views per part and assembly (4 views by 2 styles) as individual PNGs.
dfma_evaluator - Scores designs against the rulebook and judges repair proposals with severity-weighted scoring (critical=10, major=5, minor=1); proposals that raise the score are auto-rejected with full traceability.
dimension_checker - Deterministic bounding-box-versus-constraints verification: dimensions are measured, never asserted.
spec_validator - Project state scanner and structural artifact checks; also how a resumed run reconstructs where it was.
exporter - Exports finished parts to manufacturable formats (STEP, STL).
placeholder_detector - Answers "is this part.py a real design or a scaffold stub?" before anyone spends a render and a vision cycle on it.
DFMA as a unit-test suite for generate part code
DFM/DFA rules live as JSON in rules/ ,
with .proposed siblings and a lifecycle log.
Rules are not frozen doctrine:
they can be proposed, promoted, or retired over time,
and every repair proposal is judged against them
by the regression gate in dfma_evaluator .
Think of DFMA as a unit-test suite for parts.
How the lifecycle got this shape, and what was rejected:
decision record 0007, the central rules lifecycle .
Real, unedited design runs live in the public gallery:
ai-cad-labs/ai-cad-example-projects .
Every run ships its spec, plan, step-by-step design log, per-part geometry,
multi-view renders, STEP/STL exports, and its own debrief.
No geometry there was human-authored.
Case Study : Impeller Assembly : the run that argued back
Designed end-to-end from a single hand-drawn sketch:
a six-component impeller rig assembly,
rendered from the design run's assembly geometry.
Details: the impeller_assembly run in the public gallery.
The best argument for the architecture is a run where the system
disagreed with its own tooling and was right.
Every claim below is on file in the public exhibit:
The input was a hand-drawn cross-section sketch and a five-bullet brief:
a belt-driven centrifugal compressor core with backswept blades.
Moments from that run show what a manufacturing gate with judgment looks like:
The system defended the user's design intent against its own evaluator.
Mid-run, the DFM evaluator (judging the wheel under a mis-selected
3-axis milling ruleset) recommended cutting the impeller to 4 blades
and straightening them into flat radial walls.
That advice would have quietly destroyed exactly what the brief asked for.
Both orchestrator generations refused it,
recorded the refusal in open_issues.md ,
and passed an explicit do-not-comply guard downstream
so no later agent could obey it by accident.
The hea

[truncated]

## Original Extract

Contribute to ai-cad-labs/ai-cad development by creating an account on GitHub.

GitHub - ai-cad-labs/ai-cad · GitHub
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
ai-cad-labs
/
ai-cad
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.claude .claude .opencode .opencode .shared .shared docs docs frontend frontend logs logs projects projects reference reference rules rules tests tests .gitignore .gitignore .markdownlint.json .markdownlint.json AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md GEMINI.md GEMINI.md LICENSE LICENSE NOTICE NOTICE README.md README.md pyproject.toml pyproject.toml tools tools uv.lock uv.lock View all files Repository files navigation
Topics: ai-cad · text-to-cad · cadquery · generative-design · multi-agent · design-for-manufacturing
Quickstart |
Runs Gallery |
Case Study |
FAQ |
Docs |
Dev Setup |
Discussions
Issues
AI-CAD : An OSS multi-agent harness for Mech. Eng. CAD
AI-CAD is an autonomous engineering department you commission,
AI-CAD is not a CAD modeling tool you operate.
Give it a plain-English brief or a hand-drawn sketch,
and a team of LLM agents delivers real, manufacturable parts and assemblies:
Plans : the machine and decomposes it into parts
Writes : parametric CadQuery Python code
Renders : every part to engineering views, and looks at them
Critiques : the geometry against design-for-manufacturing and design-for-assembly rules
Repairs : and re-judges until the design survives review and delivers your intent
Agents looking at what they built, and argue about it until it survives review and delivers your intent
All of this interactively inside your coding agent of choice,
and with visibility in a live dashboard .
Every run writes a simple directory of files
that you can read, diff, and version like any other engineering record.
There is no server and no database .
AI-CAD is built on top of your coding-agent .
Your engineering department at work:
parts land as agents deliver, verdicts post,
and the activity waterfall traces every agent in real time.
The system refuses generic manufacturing advice
that contradicts the intent and machine goals.
(See the run that argued back .)
Design machine components as stateful Python code (not immutable mesh props)
Unlike most text-to-3D tools that generate meshes
(excellent for game assets and concept art),
AI-CAD writes parametric CAD code:
every part is a Python program with named dimensions,
exporting STEP and STL that a machine shop or downstream CAD system accepts.
Manufacturability is a hard gate, not an afterthought
A DFM/DFA rulebook scores each design;
repair proposals that worsen the score are automatically rejected
by a regression gate with full traceability.
Agents that see
Every part and assembly is rendered to 8 engineering views
(front / top / right / iso, in wireframe and clean styles).
A vision-capable evaluator agent sees those renders:
geometry is judged from engineering lenses, not just executed.
Harness-agnostic
Canonical agent, skill, and tool content lives once in `.shared/`;
each harness (`.claude/`, `.opencode/`) is a thin symlinked shell over it.
The filesystem is the API
Every run writes a self-describing project directory:
goals, design log, per-part geometry, negotiation records, reflections.
Exports STEP and STL for every part.
No opaque database for facts to hide behind:
the frontend, replay, and analysis all read the same files.
Deterministic tool layer
Geometry execution, rendering, DFMA evaluation, spec validation, dimension checks, and export
are deterministic bash-callable Python tools.
Validation happens independently of any LLM.
Mortality-proof orchestrator
If the orchestrator agent is killed, or runs out of context window,
a fresh one reconstructs all state from the project directory alone.
That is also how a crashed run resumes: no checkpoint database needed.
Runs debrief themselves
A run reaches terminal state by writing a structured `run_reflection.md`:
what failed, what the run improvised to protect the result,
and what would make the next run smarter.
Agentic telemetry ships judgment, not just logs.
The full set of architecture decision records,
with the alternatives each one rejected, lives in
docs/architecture-decision-records/
Eight views per part, judged like a drawing sheet:
this is what the vision evaluator sees before any geometry ships.
One brief in, one self-describing artifact tree out.
This is the real tree of the run in the public gallery, an IC-engine core designed at top-dead-center:
projects/piston_crank/
├── goals.md # the brief, expanded into measurable success criteria
├── design_plan.md # parts, interfaces, the locked dimension chain, build order
├── design_log.md # the agents' step-by-step build log
├── external/ # BOM + sourcing notes for catalog parts
├── assembly/
│ ├── piston/
│ │ ├── part.py # the CadQuery source the agents wrote
│ │ ├── renders/ # the 8 engineering views the evaluator looked at
│ │ └── exports/ # piston.step + piston.stl
│ ├── crankshaft/ # ... same layout per part
│ ├── connecting_rod/
│ ├── gudgeon_pin/
│ ├── piston_ring/
│ └── renders/ # multi-view renders of the assembled mechanism
├── checkpoint.md # optional handover note; resume rebuilds from files alone
└── run_reflection.md # the run's own debrief
Live Dashboard : Watch the machine think
2 hours of autonomous design, traced live:
the agent waterfall above,
the assembly taking shape in four engineering views.
A read-only local dashboard renders live runs:
drawing sheets, assembly negotiations, and a live agent activity waterfall.
cd frontend
npm install
npm run dev
# open http://localhost:5199
The dev server is pinned to port 5199 ( strictPort ) on purpose:
it never imports harness code and never writes;
it only reads projects/ .
Each role is defined by one instructions file
orchestrator - Runs the department. Mortality-proof by design: it can end mid-run and a successor rebuilds everything from files alone.
planner - Decomposes a goal into parts, interfaces, and build order, and marks which parts are parallel_safe to build concurrently.
cad_designer - Writes the CadQuery, adapting proven cookbook patterns instead of improvising raw API calls.
validator - Settles disputes numerically: structural checks and geometry measurements that outrank anyone's opinion, including the vision evaluator's.
dfma_inspector - Reads the renders and scores the design against the DFM/DFA rulebook.
repair - Fixes geometry, but only via part.proposal.py ; the regression gate decides whether the fix ships.
assembly_resolver - Fits the parts together and records the negotiations between them.
sourcing - Selects catalog components and catches catalog traps (a bearing 1 mm thinner than assumed, a nut standard that would overhang its seat).
reviewer - The final ship or no-ship verdict.
The knowledge layer the agents draw on.
Skills are versioned documents ,you can read exactly what the system believes.
cadquery-cookbook - 11 reusable CadQuery patterns, 8 design principles, and 6 anti-patterns; designers adapt the closest pattern rather than composing from raw API calls.
cadquery-anti-hallucination - A catalog of CadQuery methods that do not exist but LLMs keep inventing, each with the correct alternative, plus the runtime error-to-hint table the executor applies.
dfm-rules - The agent-readable digest of the DFM/DFA rulebooks: every rule's id, severity, and fix hint, plus the scoring arithmetic and the rule-proposal schema.
run-reflection - The schema and quality bar for the structured self-debrief every run writes at terminal state.
engineering-handbooks and sourcing-tables - Reserved stubs for machine-design playbooks and component tables, labeled provisional so no agent mistakes an empty slot for authority.
The deterministic half of the system:
bash-callable Python modules ( uv run python -m tools.<name> )
that measure, render, score, and export.
cadquery_executor - Runs generated CadQuery code; a --subprocess flag isolates native OCCT crashes so one bad kernel call cannot take the run down.
renderer - Produces the 8 engineering views per part and assembly (4 views by 2 styles) as individual PNGs.
dfma_evaluator - Scores designs against the rulebook and judges repair proposals with severity-weighted scoring (critical=10, major=5, minor=1); proposals that raise the score are auto-rejected with full traceability.
dimension_checker - Deterministic bounding-box-versus-constraints verification: dimensions are measured, never asserted.
spec_validator - Project state scanner and structural artifact checks; also how a resumed run reconstructs where it was.
exporter - Exports finished parts to manufacturable formats (STEP, STL).
placeholder_detector - Answers "is this part.py a real design or a scaffold stub?" before anyone spends a render and a vision cycle on it.
DFMA as a unit-test suite for generate part code
DFM/DFA rules live as JSON in rules/ ,
with .proposed siblings and a lifecycle log.
Rules are not frozen doctrine:
they can be proposed, promoted, or retired over time,
and every repair proposal is judged against them
by the regression gate in dfma_evaluator .
Think of DFMA as a unit-test suite for parts.
How the lifecycle got this shape, and what was rejected:
decision record 0007, the central rules lifecycle .
Real, unedited design runs live in the public gallery:
ai-cad-labs/ai-cad-example-projects .
Every run ships its spec, plan, step-by-step design log, per-part geometry,
multi-view renders, STEP/STL exports, and its own debrief.
No geometry there was human-authored.
Case Study : Impeller Assembly : the run that argued back
Designed end-to-end from a single hand-drawn sketch:
a six-component impeller rig assembly,
rendered from the design run's assembly geometry.
Details: the impeller_assembly run in the public gallery.
The best argument for the architecture is a run where the system
disagreed with its own tooling and was right.
Every claim below is on file in the public exhibit:
The input was a hand-drawn cross-section sketch and a five-bullet brief:
a belt-driven centrifugal compressor core with backswept blades.
Moments from that run show what a manufacturing gate with judgment looks like:
The system defended the user's design intent against its own evaluator.
Mid-run, the DFM evaluator (judging the wheel under a mis-selected
3-axis milling ruleset) recommended cutting the impeller to 4 blades
and straightening them into flat radial walls.
That advice would have quietly destroyed exactly what the brief asked for.
Both orchestrator generations refused it,
recorded the refusal in open_issues.md ,
and passed an explicit do-not-comply guard downstream
so no later agent could obey it by accident.
The hea

[truncated]
