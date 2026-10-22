---
source: "https://kuron.pages.dev/"
hn_url: "https://news.ycombinator.com/item?id=50191903"
title: "Show HN: Kuron – An LLM acts as your app so you can try ideas before building"
article_title: "kuron — an LLM acts as the app"
image: ""
author: "kofdai"
captured_at: "2026-10-22T07:11:13Z"
capture_tool: "hn-digest"
hn_id: 50191903
score: 1
comments: 0
posted_at: "2026-10-22T06:28:12Z"
tags:
  - hacker-news
---

# Show HN: Kuron – An LLM acts as your app so you can try ideas before building

- HN: [50191903](https://news.ycombinator.com/item?id=50191903)
- Source: [kuron.pages.dev](https://kuron.pages.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T06:28:12Z

## Translation

Title: Show HN: Kuron – An LLM acts as your app so you can try ideas before building
Article title: kuron — an LLM acts as the app
Description: Instead of generating code, an LLM with tools acts as the whole system, so you can try an idea by hand within minutes and find out whether it is worth building.

Article text:
kuron 空論 · “armchair theory”
Instead of generating code, an LLM with tools acts as the app.
Write an idea in one line — even one that would take ~50 hours to build. Within about a minute it is running on minimal screens and you can use it by hand, so you can find out whether it is worth building. Nothing is generated per idea: the LLM sets the system up, then is the system on every action.
one line ─▶ setup agent (LLM + tools) ─▶ brief · data · screens · timers · checks · reality estimate
│
button / timer / camera / mic ─▶ event agent (LLM + tools) ─▶ reads what it needs, changes the world,
│ speaks, plays sound, draws
"this should…" ─▶ request agent ─▶ changes its own brief, screens, data, checks
│
world.json ◀────┘ ─▶ minimal face: website, phone, desktop, IoT panel, watch…
The code is only a set of generic tools exposed over MCP — store and read data, define a screen, notify, look through the camera, listen to the mic, play and shape sound, speak, draw, set timers — plus a minimal renderer. Every tool checks the shape of its input and answers in words; the LLM corrects itself. There is no business logic, no word list, no per-idea exception in the code.
When a prototype behaves wrongly, you don’t patch it: you tell it how it should behave (“urgent patients go first”), and the request agent rewrites its own brief, screens and checks.
One English idea file for a small clinic: reception tablet, doctor’s PC, stock PC, patient’s phone and waiting-room display. Claude set it up from the idea in 65 s with 16 tool calls (about $0.41) , then acted as the system on every button. Each reply below is what it actually did; only the waiting time between steps is shortened.
Recorded on an M1 Max through Claude Code: Claude Sonnet 5.5 for setup and requests, Claude Haiku 5.5 for button presses. The screens are drawn by the same renderer the tool uses. The raw recording is replay.json .
A prototype where the LLM does everything is dangerously convincing. So during setup the LLM also estimates, for every action, how a shippable version would do it — ordinary code , a model , an outside service , or not yet possible — with an error rate and a delay. Flip Ideal → Existing tech and the system behaves that way: late, sometimes wrong, sound leaking past the earphones.
model urgency from symptoms, conditions, age
model rewrite the waiting-room notice
model summary, cautions, notes for the doctor
model drug, dose, days from the draft
service allergy and interaction check
code reduce stock, reorder below minimum
After you have tried it, kuron turns what the system actually did into a spec — data, screens, golden examples taken from real tool calls, this table — and an implementation plan for a coding agent.
UI and app generators (Lovable, v0, Figma Make…) build the screens and the code behind them, and prototypes usually fake the judgement with fixed data. kuron builds the behaviour : the face stays minimal, while the judgement, the data and the hand-offs between roles really happen.
Ideas whose value is a judgement — triage, recommendations, re-planning, estimates, drafted messages
Workflows across roles and devices — run a whole clinic day and see where people get stuck
Ideas that depend on technology or integrations that don’t exist yet — try the experience first
Showing an idea and changing it while you talk
Deciding whether to build, and what will be expensive
Pixel-level UI design — the face is deliberately minimal
Production — about 10 s and a few cents per action, and results vary slightly between runs
Anything needing an instant reaction (games, fine real-time audio)
Open kuron on claude.ai . It starts with the clinic above already set up; press any button, ask for a change, flip to existing tech, or write your own idea. It runs on your own Claude account (you are asked once), so you need to be signed in to claude.ai. The browser version has no microphone and runs timers when you press them.
Needs Node.js 20+ and Claude Code (logged in). No npm dependencies.
git clone https://github.com/Ag3497120/kuron
cd kuron
node bin/kuron.js dev examples # http://127.0.0.1:4870/
Pick an example — a task list, an IoT fridge with a phone app, a sales deal desk, earphones that let only important sounds through — or press + New and write your own idea in one line. Each example sets itself up the first time you open it (about a minute).
--preset local runs everything on Ollama with tool calling instead of Claude. The tools are a plain MCP server, so any MCP client can drive the same system: node bin/kuron.js hands --url http://127.0.0.1:4870 .

## Original Extract

Instead of generating code, an LLM with tools acts as the whole system, so you can try an idea by hand within minutes and find out whether it is worth building.

kuron 空論 · “armchair theory”
Instead of generating code, an LLM with tools acts as the app.
Write an idea in one line — even one that would take ~50 hours to build. Within about a minute it is running on minimal screens and you can use it by hand, so you can find out whether it is worth building. Nothing is generated per idea: the LLM sets the system up, then is the system on every action.
one line ─▶ setup agent (LLM + tools) ─▶ brief · data · screens · timers · checks · reality estimate
│
button / timer / camera / mic ─▶ event agent (LLM + tools) ─▶ reads what it needs, changes the world,
│ speaks, plays sound, draws
"this should…" ─▶ request agent ─▶ changes its own brief, screens, data, checks
│
world.json ◀────┘ ─▶ minimal face: website, phone, desktop, IoT panel, watch…
The code is only a set of generic tools exposed over MCP — store and read data, define a screen, notify, look through the camera, listen to the mic, play and shape sound, speak, draw, set timers — plus a minimal renderer. Every tool checks the shape of its input and answers in words; the LLM corrects itself. There is no business logic, no word list, no per-idea exception in the code.
When a prototype behaves wrongly, you don’t patch it: you tell it how it should behave (“urgent patients go first”), and the request agent rewrites its own brief, screens and checks.
One English idea file for a small clinic: reception tablet, doctor’s PC, stock PC, patient’s phone and waiting-room display. Claude set it up from the idea in 65 s with 16 tool calls (about $0.41) , then acted as the system on every button. Each reply below is what it actually did; only the waiting time between steps is shortened.
Recorded on an M1 Max through Claude Code: Claude Sonnet 5.5 for setup and requests, Claude Haiku 5.5 for button presses. The screens are drawn by the same renderer the tool uses. The raw recording is replay.json .
A prototype where the LLM does everything is dangerously convincing. So during setup the LLM also estimates, for every action, how a shippable version would do it — ordinary code , a model , an outside service , or not yet possible — with an error rate and a delay. Flip Ideal → Existing tech and the system behaves that way: late, sometimes wrong, sound leaking past the earphones.
model urgency from symptoms, conditions, age
model rewrite the waiting-room notice
model summary, cautions, notes for the doctor
model drug, dose, days from the draft
service allergy and interaction check
code reduce stock, reorder below minimum
After you have tried it, kuron turns what the system actually did into a spec — data, screens, golden examples taken from real tool calls, this table — and an implementation plan for a coding agent.
UI and app generators (Lovable, v0, Figma Make…) build the screens and the code behind them, and prototypes usually fake the judgement with fixed data. kuron builds the behaviour : the face stays minimal, while the judgement, the data and the hand-offs between roles really happen.
Ideas whose value is a judgement — triage, recommendations, re-planning, estimates, drafted messages
Workflows across roles and devices — run a whole clinic day and see where people get stuck
Ideas that depend on technology or integrations that don’t exist yet — try the experience first
Showing an idea and changing it while you talk
Deciding whether to build, and what will be expensive
Pixel-level UI design — the face is deliberately minimal
Production — about 10 s and a few cents per action, and results vary slightly between runs
Anything needing an instant reaction (games, fine real-time audio)
Open kuron on claude.ai . It starts with the clinic above already set up; press any button, ask for a change, flip to existing tech, or write your own idea. It runs on your own Claude account (you are asked once), so you need to be signed in to claude.ai. The browser version has no microphone and runs timers when you press them.
Needs Node.js 20+ and Claude Code (logged in). No npm dependencies.
git clone https://github.com/Ag3497120/kuron
cd kuron
node bin/kuron.js dev examples # http://127.0.0.1:4870/
Pick an example — a task list, an IoT fridge with a phone app, a sales deal desk, earphones that let only important sounds through — or press + New and write your own idea in one line. Each example sets itself up the first time you open it (about a minute).
--preset local runs everything on Ollama with tool calling instead of Claude. The tools are a plain MCP server, so any MCP client can drive the same system: node bin/kuron.js hands --url http://127.0.0.1:4870 .
