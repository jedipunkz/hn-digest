---
source: "https://github.com/tapaskar/archwhere"
hn_url: "https://news.ycombinator.com/item?id=50313522"
title: "Show HN: Archwhere – work alongside AI coding agents on an architecture map"
article_title: "GitHub - tapaskar/archwhere: Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram. · GitHub"
image: "https://opengraph.githubassets.com/d23b06ab808099fa4d91d950531070b94b532716d6485ec484038765e3e11f04/tapaskar/archwhere"
author: "sparkupcloud"
captured_at: "2026-10-31T06:41:19Z"
capture_tool: "hn-digest"
hn_id: 50313522
score: 1
comments: 0
posted_at: "2026-10-31T05:46:30Z"
tags:
  - hacker-news
---

# Show HN: Archwhere – work alongside AI coding agents on an architecture map

- HN: [50313522](https://news.ycombinator.com/item?id=50313522)
- Source: [github.com](https://github.com/tapaskar/archwhere)
- Score: 1
- Comments: 0
- Posted: 2026-10-31T05:46:30Z

## Translation

Title: Show HN: Archwhere – work alongside AI coding agents on an architecture map
Article title: GitHub - tapaskar/archwhere: Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram. · GitHub
Description: Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram. - tapaskar/archwhere
HN text: I am a solo founder/developer, my agents shipping 30 PRs a week, couldn’t see which part of product they were changing; existing arch tools die because humans won’t maintain the model — so I made the agents maintain it. MIT license

Article text:
GitHub - tapaskar/archwhere: Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram. · GitHub
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
tapaskar
archwhere Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram.
Readme MIT license Activity Stars
0 forks Report repository Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
9 Commits 9 Commits Folders and files
.github/ workflows .github/ workflows arch arch bin bin docs docs examples/ meridian examples/ meridian lib lib templates templates test test ui ui .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md LICENSE LICENSE README.md README.md SKILL.md SKILL.md action.yml action.yml package.json package.json View all files Repository files navigation
Your AI coding agents ship thirty PRs a week. Do you know where they land?
archwhere puts every change — every PR, every task, every agent — on a declared
architecture map : which component it touches, whether that component is a
security zone, and what needs you right now . Then it goes further: the map
becomes the place you talk to your agents .
Mission control on a fictional product: the attention rail (security → blocked →
needs-attention → ready-to-merge), the component map with live badges, and a
thread pinned to a component — answered by a coding agent, with its PR.
Zero dependencies. One JSON file. Works with Claude, or any agent that can see GitHub.
Supervising an agent workforce diff-by-diff doesn't scale. Architecture tools
(Structurizr, IcePanel) assume a human architect maintains the model — so the
model rots. CODEOWNERS and labelers map files to owners but carry no meaning.
archwhere inverts it: your agents maintain the map as a condition of merging
(there's a SKILL.md that tells them how, and archwhere check fails CI when
the map goes stale). You read components, not diffs.
npx archwhere init
That drafts the model from your repo itself — directory structure,
package.json workspaces, import edges — with guessed kinds (auth → security,
ui → product, scripts → ops), plus the PR-comment workflow and two agent
listeners. archwhere check passes from minute one; you rename and refine
instead of starting from a blank file. It also prints a prompt you paste to
your coding agent so it polishes the model by surveying the repo. The
cold-start problem every architecture-as-code tool dies on? Automated.
npx archwhere # what is my working tree touching?
npx archwhere --base origin/main # what does this branch touch?
npx archwhere check # CI guard: model valid, no file unmapped, no stale glob
npx archwhere check --fix # adopt unmapped files into the nearest component
npx archwhere heat # which components churned most in the last 30 days?
npx archwhere doc --out ARCHITECTURE.md
npx archwhere mission # the UI at http://localhost:4610
The map heals itself
When an agent's PR adds files the model doesn't own, the PR comment doesn't
just complain — it suggests the exact patch (adopt into the nearest
component, or create a new one), and archwhere check --fix applies it. Tell
your agents "run check --fix and review the adopted globs" and the map stays
alive with no human bookkeeping.
Add the action (one of the files init writes):
- uses : tapaskar/archwhere@v1
Every PR gets a comment: a component/kind/files table, security flags, and the
system map with touched components highlighted — rendered natively by GitHub:
flowchart LR
web-app["Web app"]
mobile-api["Mobile API"]
admin-console["Admin console"]
ticket-engine["Ticket engine"]
auth["Auth & sessions"]
search["Search"]
notifications["Notifications"]
billing["Billing"]
llm-gateway["LLM gateway"]
jobs["Jobs & queues"]
ops["Ops & CI"]
docs["Docs"]
web-app --> ticket-engine
web-app --> auth
mobile-api --> ticket-engine
mobile-api --> auth
admin-console --> ticket-engine
ticket-engine --> search
ticket-engine --> notifications
ticket-engine --> llm-gateway
search --> jobs
notifications --> jobs
classDef core fill:#E8EEF9,stroke:#2F4AC0,color:#1F2433
classDef security fill:#FDE8E4,stroke:#C2410C,color:#1F2433
classDef feature fill:#FFF4E0,stroke:#B45309,color:#1F2433
classDef infra fill:#EFEFEF,stroke:#595959,color:#1F2433
classDef product fill:#E3F2EC,stroke:#0F7A5C,color:#1F2433
classDef ops fill:#F3E8F9,stroke:#7E3FA0,color:#1F2433
classDef docs fill:#F7F7F7,stroke:#8A8A8A,color:#1F2433
classDef other fill:#F2F2EF,stroke:#777,color:#1F2433
classDef touched fill:#F7B32B,stroke:#1F2433,stroke-width:3px,color:#1F2433
class web-app,mobile-api,admin-console product
class ticket-engine core
class auth security
class search,notifications,billing feature
class llm-gateway,jobs infra
class ops ops
class docs docs
Loading
Talk to your agents on the diagram
npx archwhere mission serves a local UI where the map is the conversation surface:
Attention rail — security items, blocked decisions, PRs needing attention,
PRs ready to merge (the Merge button uses your token; agents never merge).
Click any component → open a thread. Threads are GitHub issues labelled
component:<id> (auto-created). Your agent answers; the reply appears on the map.
Four lenses — Structure / Changes / Heat / Runtime.
Heat colours components by git churn over the last 30 days — where your
agents are actually spending effort. Runtime colours them by live health
from any endpoint you point RUNTIME_URL at, returning
{"components": {"<id>": {"status": "ok|warn|bad", "lines": ["…"]}}} .
The bus is GitHub — archwhere never talks to an agent directly:
Which mention the UI appends is yours: arch/mission.config.json → {"agentMention": "@your-bot"} (env AGENT_MENTION overrides; "" disables).
One file, declared judgment — not an auto-extracted import hairball:
{
"system" : " Meridian " ,
"components" : {
"auth" : {
"name" : " Auth & sessions " , "kind" : " security " ,
"provider" : " Auth0 + RBAC " , // shows on every surface: map, PR comment, docs, mission UI
"paths" : [ " core/auth/** " ], // ** crosses dirs, * stays in one; longest match wins
"watch" : [{ "pattern" : " session|token|rbac " , "note" : " security review " }],
"deps" : []
}
},
"lanes" : [[ " web-app " ], [ " auth " ]], // optional; defaults to kind grouping
"flows" : { "signup" : { "name" : " … " , "actor" : " … " , "steps" : [{ "from" : " web-app " , "to" : " auth " , "label" : " SSO " }] } }
}
Kinds: core · security · feature · infra · product · ops · docs . flows are
your C4-style dynamic views — the process layer above the components. The full
example: examples/meridian/architecture.json .
Interactive diagrams (Archify)
npx archwhere ir compiles the model into Archify
IR — an animated signal-flow system diagram with trust boundaries and guided
views per flow, plus one numbered diagram per flow. Render with Archify's CLI.
SKILL.md is the agent's side of the contract: place work on the map
before coding, extend the model when adding modules, update flows when changing
user-visible processes, keep archwhere check green. Works as a Claude Code /
agent skill, or paste it into your agent's instructions.
maps PRs
visual
agents maintain it
talk to agents on it
CODEOWNERS / labeler
✓
—
—
—
Structurizr / IcePanel
—
✓
—
—
dependency graphs
✓
hairball
n/a
—
archwhere
✓
✓
✓
✓
Onboarding at a glance
From stranger to agent-replies-on-the-map in under twenty minutes — the full
funnel, stage by stage: docs/ONBOARDING.md .
MIT © Tapas Kar. Built with and for AI coding agents — this repo's own model
lives in arch/architecture.json , of course.
Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram. - tapaskar/archwhere

I am a solo founder/developer, my agents shipping 30 PRs a week, couldn’t see which part of product they were changing; existing arch tools die because humans won’t maintain the model — so I made the agents maintain it. MIT license

GitHub - tapaskar/archwhere: Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram. · GitHub
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
tapaskar
archwhere Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram.
Readme MIT license Activity Stars
0 forks Report repository Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
9 Commits 9 Commits Folders and files
.github/ workflows .github/ workflows arch arch bin bin docs docs examples/ meridian examples/ meridian lib lib templates templates test test ui ui .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md LICENSE LICENSE README.md README.md SKILL.md SKILL.md action.yml action.yml package.json package.json View all files Repository files navigation
Your AI coding agents ship thirty PRs a week. Do you know where they land?
archwhere puts every change — every PR, every task, every agent — on a declared
architecture map : which component it touches, whether that component is a
security zone, and what needs you right now . Then it goes further: the map
becomes the place you talk to your agents .
Mission control on a fictional product: the attention rail (security → blocked →
needs-attention → ready-to-merge), the component map with live badges, and a
thread pinned to a component — answered by a coding agent, with its PR.
Zero dependencies. One JSON file. Works with Claude, or any agent that can see GitHub.
Supervising an agent workforce diff-by-diff doesn't scale. Architecture tools
(Structurizr, IcePanel) assume a human architect maintains the model — so the
model rots. CODEOWNERS and labelers map files to owners but carry no meaning.
archwhere inverts it: your agents maintain the map as a condition of merging
(there's a SKILL.md that tells them how, and archwhere check fails CI when
the map goes stale). You read components, not diffs.
npx archwhere init
That drafts the model from your repo itself — directory structure,
package.json workspaces, import edges — with guessed kinds (auth → security,
ui → product, scripts → ops), plus the PR-comment workflow and two agent
listeners. archwhere check passes from minute one; you rename and refine
instead of starting from a blank file. It also prints a prompt you paste to
your coding agent so it polishes the model by surveying the repo. The
cold-start problem every architecture-as-code tool dies on? Automated.
npx archwhere # what is my working tree touching?
npx archwhere --base origin/main # what does this branch touch?
npx archwhere check # CI guard: model valid, no file unmapped, no stale glob
npx archwhere check --fix # adopt unmapped files into the nearest component
npx archwhere heat # which components churned most in the last 30 days?
npx archwhere doc --out ARCHITECTURE.md
npx archwhere mission # the UI at http://localhost:4610
The map heals itself
When an agent's PR adds files the model doesn't own, the PR comment doesn't
just complain — it suggests the exact patch (adopt into the nearest
component, or create a new one), and archwhere check --fix applies it. Tell
your agents "run check --fix and review the adopted globs" and the map stays
alive with no human bookkeeping.
Add the action (one of the files init writes):
- uses : tapaskar/archwhere@v1
Every PR gets a comment: a component/kind/files table, security flags, and the
system map with touched components highlighted — rendered natively by GitHub:
flowchart LR
web-app["Web app"]
mobile-api["Mobile API"]
admin-console["Admin console"]
ticket-engine["Ticket engine"]
auth["Auth & sessions"]
search["Search"]
notifications["Notifications"]
billing["Billing"]
llm-gateway["LLM gateway"]
jobs["Jobs & queues"]
ops["Ops & CI"]
docs["Docs"]
web-app --> ticket-engine
web-app --> auth
mobile-api --> ticket-engine
mobile-api --> auth
admin-console --> ticket-engine
ticket-engine --> search
ticket-engine --> notifications
ticket-engine --> llm-gateway
search --> jobs
notifications --> jobs
classDef core fill:#E8EEF9,stroke:#2F4AC0,color:#1F2433
classDef security fill:#FDE8E4,stroke:#C2410C,color:#1F2433
classDef feature fill:#FFF4E0,stroke:#B45309,color:#1F2433
classDef infra fill:#EFEFEF,stroke:#595959,color:#1F2433
classDef product fill:#E3F2EC,stroke:#0F7A5C,color:#1F2433
classDef ops fill:#F3E8F9,stroke:#7E3FA0,color:#1F2433
classDef docs fill:#F7F7F7,stroke:#8A8A8A,color:#1F2433
classDef other fill:#F2F2EF,stroke:#777,color:#1F2433
classDef touched fill:#F7B32B,stroke:#1F2433,stroke-width:3px,color:#1F2433
class web-app,mobile-api,admin-console product
class ticket-engine core
class auth security
class search,notifications,billing feature
class llm-gateway,jobs infra
class ops ops
class docs docs
Loading
Talk to your agents on the diagram
npx archwhere mission serves a local UI where the map is the conversation surface:
Attention rail — security items, blocked decisions, PRs needing attention,
PRs ready to merge (the Merge button uses your token; agents never merge).
Click any component → open a thread. Threads are GitHub issues labelled
component:<id> (auto-created). Your agent answers; the reply appears on the map.
Four lenses — Structure / Changes / Heat / Runtime.
Heat colours components by git churn over the last 30 days — where your
agents are actually spending effort. Runtime colours them by live health
from any endpoint you point RUNTIME_URL at, returning
{"components": {"<id>": {"status": "ok|warn|bad", "lines": ["…"]}}} .
The bus is GitHub — archwhere never talks to an agent directly:
Which mention the UI appends is yours: arch/mission.config.json → {"agentMention": "@your-bot"} (env AGENT_MENTION overrides; "" disables).
One file, declared judgment — not an auto-extracted import hairball:
{
"system" : " Meridian " ,
"components" : {
"auth" : {
"name" : " Auth & sessions " , "kind" : " security " ,
"provider" : " Auth0 + RBAC " , // shows on every surface: map, PR comment, docs, mission UI
"paths" : [ " core/auth/** " ], // ** crosses dirs, * stays in one; longest match wins
"watch" : [{ "pattern" : " session|token|rbac " , "note" : " security review " }],
"deps" : []
}
},
"lanes" : [[ " web-app " ], [ " auth " ]], // optional; defaults to kind grouping
"flows" : { "signup" : { "name" : " … " , "actor" : " … " , "steps" : [{ "from" : " web-app " , "to" : " auth " , "label" : " SSO " }] } }
}
Kinds: core · security · feature · infra · product · ops · docs . flows are
your C4-style dynamic views — the process layer above the components. The full
example: examples/meridian/architecture.json .
Interactive diagrams (Archify)
npx archwhere ir compiles the model into Archify
IR — an animated signal-flow system diagram with trust boundaries and guided
views per flow, plus one numbered diagram per flow. Render with Archify's CLI.
SKILL.md is the agent's side of the contract: place work on the map
before coding, extend the model when adding modules, update flows when changing
user-visible processes, keep archwhere check green. Works as a Claude Code /
agent skill, or paste it into your agent's instructions.
maps PRs
visual
agents maintain it
talk to agents on it
CODEOWNERS / labeler
✓
—
—
—
Structurizr / IcePanel
—
✓
—
—
dependency graphs
✓
hairball
n/a
—
archwhere
✓
✓
✓
✓
Onboarding at a glance
From stranger to agent-replies-on-the-map in under twenty minutes — the full
funnel, stage by stage: docs/ONBOARDING.md .
MIT © Tapas Kar. Built with and for AI coding agents — this repo's own model
lives in arch/architecture.json , of course.
Mission control for AI coding agents — every PR lands on the architecture map, and you talk to your agents on the diagram.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
