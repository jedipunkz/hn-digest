---
source: "https://github.com/s4ps4n/prompt-hospital"
hn_url: "https://news.ycombinator.com/item?id=49909739"
title: "I built an isometric office where my AI models work (and Slack off)"
article_title: "GitHub - s4ps4n/prompt-hospital: Prompt Hospital — an AI-orchestrator visualization in the style of Theme Hospital. / Prompt Hospital — визуализация ИИ-оркестратора в стиле Theme Hospital · GitHub"
image: "https://opengraph.githubassets.com/791b31af6601b1d9ea550db874da37105ff8d8ecedb325754e8aeae8ddb8f3e8/s4ps4n/prompt-hospital"
author: "sapsanius"
captured_at: "2026-09-30T15:42:01Z"
capture_tool: "hn-digest"
hn_id: 49909739
score: 2
comments: 0
posted_at: "2026-09-30T14:47:26Z"
tags:
  - hacker-news
---

# I built an isometric office where my AI models work (and Slack off)

- HN: [49909739](https://news.ycombinator.com/item?id=49909739)
- Source: [github.com](https://github.com/s4ps4n/prompt-hospital)
- Score: 2
- Comments: 0
- Posted: 2026-09-30T14:47:26Z

## Translation

Title: I built an isometric office where my AI models work (and Slack off)
Article title: GitHub - s4ps4n/prompt-hospital: Prompt Hospital — an AI-orchestrator visualization in the style of Theme Hospital. / Prompt Hospital — визуализация ИИ-оркестратора в стиле Theme Hospital · GitHub
Description: Prompt Hospital — an AI-orchestrator visualization in the style of Theme Hospital. / Prompt Hospital — визуализация ИИ-оркестратора в стиле Theme Hospital - s4ps4n/prompt-hospital

Article text:
GitHub - s4ps4n/prompt-hospital: Prompt Hospital — an AI-orchestrator visualization in the style of Theme Hospital. / Prompt Hospital — визуализация ИИ-оркестратора в стиле Theme Hospital · GitHub
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
s4ps4n
/
prompt-hospital
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
49 Commits 49 Commits Folders and files
docs docs orchestrator orchestrator public public scripts scripts src src .gitignore .gitignore .oxlintrc.json .oxlintrc.json INSTALL.md INSTALL.md INSTALL.ru.md INSTALL.ru.md LICENSE LICENSE README.md README.md README.ru.md README.ru.md RELEASE_NOTES.md RELEASE_NOTES.md ROADMAP.md ROADMAP.md ROADMAP.ru.md ROADMAP.ru.md index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Your AI models finally have an office.
An isometric, Theme Hospital–style control panel for an AI team: each model gets a room,
tasks arrive as envelopes, and when you drag one onto a model — the model actually runs.
▸ Live demo ·
Quick start ·
Add your model ·
Roadmap ·
Русский
An isometric scene where each model of the team sits in its own room. Tasks are envelopes in Hermes's tray (the coordinator). Drag an envelope onto a model → the task lands in the orchestrator journal → the dispatcher actually runs the model with that task.
Frontend — Vite + React + TypeScript + SVG, offline, no external graphics libraries.
Journal — single source of truth ( journal.json ): models, tasks, statuses, history.
API — server.py , serves the journal and accepts operations.
Dispatcher — dispatch.py , picks up assigned tasks and runs executors.
prompt-hospital/ # frontend (this repository)
src/journal/ # journal: types, ops (pure operations), selectors, store, remote
src/scene/ # isometric scene, rooms
src/characters/ # chibi characters, statuses
src/animations/ # idle animations (walk/coffee/sleep), gait
src/app/ src/ui/ # layout, model card, tray, drag-and-drop
src/theme/ # palette, CSS animations
orchestrator/ # control-panel backend (in this same repository)
journal.example.json # starter journal — copy to journal.json
journal.py # operation CLI (same ops as the frontend OPS)
server.py # HTTP API: GET /journal, POST /op
dispatch.py # dispatcher: task → actually run the model
Architecture
┌────────────────────────────────────────────┐
│ Office (React, example.com) │
│ drag-and-drop, card, catalog, tray │
└───────────────┬────────────────────────────┘
GET /journal│ POST /op (operations)
▼
┌────────────────────────────────────────────┐
│ nginx (Basic Auth) │
│ /journal, /op → orchestrator-api │
└───────────────┬────────────────────────────┘
▼
┌────────────────────────────────────────────┐
│ server.py (API :8090, docker) │
│ GET /journal, POST /op → journal.py │
└───────────────┬────────────────────────────┘
▼
┌────────────────────────────────────────────┐
│ journal.json (source of truth) │
└───────────────┬────────────────────────────┘
│ polls on a loop (--loop)
▼
┌────────────────────────────────────────────┐
│ dispatch.py — actually runs models │
│ codex exec / claude -p │
└────────────────────────────────────────────┘
Installation
git clone https://github.com/s4ps4n/prompt-hospital.git
cd prompt-hospital
npm install
npm run dev # dev server at http://localhost:5173
Journal integration is enabled with an env var at build time:
VITE_JOURNAL_URL=/journal npm run build # office reads/writes the journal at /journal
npm run preview # http://localhost:4173
Without VITE_JOURNAL_URL the office runs standalone (journal in localStorage, seeded tasks).
cd orchestrator
cp journal.example.json journal.json
python3 journal.py list # view the journal
python3 server.py 8090 # run the API (locally)
python3 dispatch.py --once # one dispatcher pass
Wiring it together
Create the journal — cp orchestrator/journal.example.json orchestrator/journal.json .
Run the API — python3 orchestrator/server.py 8090 (serves /journal , accepts /op ).
Build the frontend with VITE_JOURNAL_URL=/journal and serve the static files behind
nginx, proxying /journal and /op to the API (see "Deploy").
ORCHESTRATOR_API=https://example.com \
ORCHESTRATOR_AUTH= ' user:pass ' \
python3 orchestrator/dispatch.py --loop 30
Journal operations
journal.py (and POST /op {name, args} ) supports the same operations as the frontend:
Roles: coordinator, executor, architect, fullstack, sysadmin, designer, UX/UI, acceptor, reviewer .
Bookkeeping rule: whenever you start a model on a task, immediately addTask + assign in the journal; when done, complete . Otherwise the office shows a lie.
A model has two sides: how it looks in the office and what actually runs when it gets a task.
1. Appearance — add an entry to CATALOG in src/journal/catalog.ts :
{ model : 'phi' , name : 'Phi' , provider : 'Microsoft' , hair : '#6fb7e8' , color : '#3f87b8' , skin : '#f2c9a0' , style : 'bob' , acc : 'glasses' } ,
style : cap | bob | quiff | bun | spiky ; acc (optional): wings | glasses | headset .
Models not in the catalog still get a stable look — by prefix ( qwen-27b → Qwen), by provider,
or by a hash of the name — so this step is about a face of its own, not about working at all.
2. Execution — in orchestrator/dispatch.py , inside tick() , add a branch:
model = ( w . get ( 'model' ) or '' ). lower ()
if 'codex' in model :
ok = run_codex ( task [ 'id' ], task [ 'title' ])
elif 'my-model' in model :
ok = run_my_model ( task [ 'title' ])
Then write run_my_model — the call to your executor (CLI/API). The dispatcher records
complete / block by itself.
A PR titled Add model: <Name> with a screenshot of the room is the easiest first contribution.
Docker + nginx (Basic Auth) + Traefik with HTTPS — step by step in INSTALL.md, Level 2 .
A. Journal — single source of truth, operation CLI.
B. Monitor — office reads the journal (read-only).
C. Write — drag-and-drop writes to the journal via POST /op .
D. Execute — the dispatcher actually runs the model and closes the task.
By default the office runs at level C (writes to the journal); D is enabled by running dispatch.py .
How the author runs it; any orchestrator that speaks HTTP works the same way.
npm test # vitest — journal, operations, interaction, remote mode
npm run lint # oxlint
npm run build # tsc + vite build
Roadmap
✅ Office and rooms — models, tray, journal, idle animations
✅ Own orchestrator — journal, API, dispatcher that actually runs models
🔜 Stats — who worked how much, per model
🔜 The Warden — a hand that slaps idle models back to Hermes for a task
Inspired by Theme Hospital (Bullfrog, 1997); no assets, names or code from the original are used.
MIT © s4ps4n.
Made in Ulan-Ude by @s4ps4n at ITTEK — we build web products, online stores and AI tooling.
Prompt Hospital — an AI-orchestrator visualization in the style of Theme Hospital. / Prompt Hospital — визуализация ИИ-оркестратора в стиле Theme Hospital
prompthospital.site/ Resources
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Prompt Hospital — an AI-orchestrator visualization in the style of Theme Hospital. / Prompt Hospital — визуализация ИИ-оркестратора в стиле Theme Hospital - s4ps4n/prompt-hospital

GitHub - s4ps4n/prompt-hospital: Prompt Hospital — an AI-orchestrator visualization in the style of Theme Hospital. / Prompt Hospital — визуализация ИИ-оркестратора в стиле Theme Hospital · GitHub
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
s4ps4n
/
prompt-hospital
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
49 Commits 49 Commits Folders and files
docs docs orchestrator orchestrator public public scripts scripts src src .gitignore .gitignore .oxlintrc.json .oxlintrc.json INSTALL.md INSTALL.md INSTALL.ru.md INSTALL.ru.md LICENSE LICENSE README.md README.md README.ru.md README.ru.md RELEASE_NOTES.md RELEASE_NOTES.md ROADMAP.md ROADMAP.md ROADMAP.ru.md ROADMAP.ru.md index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.app.json tsconfig.app.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Your AI models finally have an office.
An isometric, Theme Hospital–style control panel for an AI team: each model gets a room,
tasks arrive as envelopes, and when you drag one onto a model — the model actually runs.
▸ Live demo ·
Quick start ·
Add your model ·
Roadmap ·
Русский
An isometric scene where each model of the team sits in its own room. Tasks are envelopes in Hermes's tray (the coordinator). Drag an envelope onto a model → the task lands in the orchestrator journal → the dispatcher actually runs the model with that task.
Frontend — Vite + React + TypeScript + SVG, offline, no external graphics libraries.
Journal — single source of truth ( journal.json ): models, tasks, statuses, history.
API — server.py , serves the journal and accepts operations.
Dispatcher — dispatch.py , picks up assigned tasks and runs executors.
prompt-hospital/ # frontend (this repository)
src/journal/ # journal: types, ops (pure operations), selectors, store, remote
src/scene/ # isometric scene, rooms
src/characters/ # chibi characters, statuses
src/animations/ # idle animations (walk/coffee/sleep), gait
src/app/ src/ui/ # layout, model card, tray, drag-and-drop
src/theme/ # palette, CSS animations
orchestrator/ # control-panel backend (in this same repository)
journal.example.json # starter journal — copy to journal.json
journal.py # operation CLI (same ops as the frontend OPS)
server.py # HTTP API: GET /journal, POST /op
dispatch.py # dispatcher: task → actually run the model
Architecture
┌────────────────────────────────────────────┐
│ Office (React, example.com) │
│ drag-and-drop, card, catalog, tray │
└───────────────┬────────────────────────────┘
GET /journal│ POST /op (operations)
▼
┌────────────────────────────────────────────┐
│ nginx (Basic Auth) │
│ /journal, /op → orchestrator-api │
└───────────────┬────────────────────────────┘
▼
┌────────────────────────────────────────────┐
│ server.py (API :8090, docker) │
│ GET /journal, POST /op → journal.py │
└───────────────┬────────────────────────────┘
▼
┌────────────────────────────────────────────┐
│ journal.json (source of truth) │
└───────────────┬────────────────────────────┘
│ polls on a loop (--loop)
▼
┌────────────────────────────────────────────┐
│ dispatch.py — actually runs models │
│ codex exec / claude -p │
└────────────────────────────────────────────┘
Installation
git clone https://github.com/s4ps4n/prompt-hospital.git
cd prompt-hospital
npm install
npm run dev # dev server at http://localhost:5173
Journal integration is enabled with an env var at build time:
VITE_JOURNAL_URL=/journal npm run build # office reads/writes the journal at /journal
npm run preview # http://localhost:4173
Without VITE_JOURNAL_URL the office runs standalone (journal in localStorage, seeded tasks).
cd orchestrator
cp journal.example.json journal.json
python3 journal.py list # view the journal
python3 server.py 8090 # run the API (locally)
python3 dispatch.py --once # one dispatcher pass
Wiring it together
Create the journal — cp orchestrator/journal.example.json orchestrator/journal.json .
Run the API — python3 orchestrator/server.py 8090 (serves /journal , accepts /op ).
Build the frontend with VITE_JOURNAL_URL=/journal and serve the static files behind
nginx, proxying /journal and /op to the API (see "Deploy").
ORCHESTRATOR_API=https://example.com \
ORCHESTRATOR_AUTH= ' user:pass ' \
python3 orchestrator/dispatch.py --loop 30
Journal operations
journal.py (and POST /op {name, args} ) supports the same operations as the frontend:
Roles: coordinator, executor, architect, fullstack, sysadmin, designer, UX/UI, acceptor, reviewer .
Bookkeeping rule: whenever you start a model on a task, immediately addTask + assign in the journal; when done, complete . Otherwise the office shows a lie.
A model has two sides: how it looks in the office and what actually runs when it gets a task.
1. Appearance — add an entry to CATALOG in src/journal/catalog.ts :
{ model : 'phi' , name : 'Phi' , provider : 'Microsoft' , hair : '#6fb7e8' , color : '#3f87b8' , skin : '#f2c9a0' , style : 'bob' , acc : 'glasses' } ,
style : cap | bob | quiff | bun | spiky ; acc (optional): wings | glasses | headset .
Models not in the catalog still get a stable look — by prefix ( qwen-27b → Qwen), by provider,
or by a hash of the name — so this step is about a face of its own, not about working at all.
2. Execution — in orchestrator/dispatch.py , inside tick() , add a branch:
model = ( w . get ( 'model' ) or '' ). lower ()
if 'codex' in model :
ok = run_codex ( task [ 'id' ], task [ 'title' ])
elif 'my-model' in model :
ok = run_my_model ( task [ 'title' ])
Then write run_my_model — the call to your executor (CLI/API). The dispatcher records
complete / block by itself.
A PR titled Add model: <Name> with a screenshot of the room is the easiest first contribution.
Docker + nginx (Basic Auth) + Traefik with HTTPS — step by step in INSTALL.md, Level 2 .
A. Journal — single source of truth, operation CLI.
B. Monitor — office reads the journal (read-only).
C. Write — drag-and-drop writes to the journal via POST /op .
D. Execute — the dispatcher actually runs the model and closes the task.
By default the office runs at level C (writes to the journal); D is enabled by running dispatch.py .
How the author runs it; any orchestrator that speaks HTTP works the same way.
npm test # vitest — journal, operations, interaction, remote mode
npm run lint # oxlint
npm run build # tsc + vite build
Roadmap
✅ Office and rooms — models, tray, journal, idle animations
✅ Own orchestrator — journal, API, dispatcher that actually runs models
🔜 Stats — who worked how much, per model
🔜 The Warden — a hand that slaps idle models back to Hermes for a task
Inspired by Theme Hospital (Bullfrog, 1997); no assets, names or code from the original are used.
MIT © s4ps4n.
Made in Ulan-Ude by @s4ps4n at ITTEK — we build web products, online stores and AI tooling.
Prompt Hospital — an AI-orchestrator visualization in the style of Theme Hospital. / Prompt Hospital — визуализация ИИ-оркестратора в стиле Theme Hospital
prompthospital.site/ Resources
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
