---
source: "https://github.com/drawcms/drawcms"
hn_url: "https://news.ycombinator.com/item?id=49834430"
title: "Show HN: DrawCMS – An open-source animated diagramming tool for AI Agents"
article_title: "GitHub - drawcms/drawcms: The animated diagram editor AI agents can draw on. Animated technical diagrams (architecture, sequence, BPMN, cloud icons) built by you and your agent together on the same live canvas, via WebMCP. Open-source, local-first. · GitHub"
image: "https://repository-images.githubusercontent.com/1357305997/1ebf7b19-22ac-457f-8e77-970919b995ce"
author: "dimas3399"
captured_at: "2026-09-24T18:15:17Z"
capture_tool: "hn-digest"
hn_id: 49834430
score: 1
comments: 0
posted_at: "2026-09-24T17:57:11Z"
tags:
  - hacker-news
---

# Show HN: DrawCMS – An open-source animated diagramming tool for AI Agents

- HN: [49834430](https://news.ycombinator.com/item?id=49834430)
- Source: [github.com](https://github.com/drawcms/drawcms)
- Score: 1
- Comments: 0
- Posted: 2026-09-24T17:57:11Z

## Translation

Title: Show HN: DrawCMS – An open-source animated diagramming tool for AI Agents
Article title: GitHub - drawcms/drawcms: The animated diagram editor AI agents can draw on. Animated technical diagrams (architecture, sequence, BPMN, cloud icons) built by you and your agent together on the same live canvas, via WebMCP. Open-source, local-first. · GitHub
Description: The animated diagram editor AI agents can draw on. Animated technical diagrams (architecture, sequence, BPMN, cloud icons) built by you and your agent together on the same live canvas, via WebMCP. Open-source, local-first. - drawcms/drawcms
HN text: Hello, I'm building DrawCMS. It's an open source animated diagram editor you can self-host. You can ask a compatible AI agent to read your repository and create an editable diagram, then walk through the flow step by step on the canvas. AI makes it easy to write code quickly. Keeping track of how that code works can be harder. You finish a feature, move on, and a week later you're trying to remember where a request goes, which service writes to the database, or why a background job runs. It supports architecture, sequence, network and etc diagrams, infrastructure icons also . You can also draw and edit everything manually. Repo: https://github.com/drawcms/drawcms Quickstart (Claude Code, Codex, OpenClaw, Hermes): npx skills add drawcms/drawcms-skill -g Thanks for reading, and I hope you try it out! I would love to hear any feedback

Article text:
GitHub - drawcms/drawcms: The animated diagram editor AI agents can draw on. Animated technical diagrams (architecture, sequence, BPMN, cloud icons) built by you and your agent together on the same live canvas, via WebMCP. Open-source, local-first. · GitHub
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
drawcms
/
drawcms
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
58 Commits 58 Commits Folders and files
.github/ workflows .github/ workflows adr adr blog blog docs docs public public site-blog site-blog site site src src .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CLA.md CLA.md COMMERCIAL-LICENSE.md COMMERCIAL-LICENSE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md eslint.config.mjs eslint.config.mjs netlify.toml netlify.toml next.config.ts next.config.ts open-next.config.ts open-next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs render.yaml render.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts wrangler.jsonc wrangler.jsonc View all files Repository files navigation
Diagrams that explain how a system moves, not just how it is wired.
Describe a flow to your AI agent, or draw it yourself. DrawCMS turns it into an animated technical diagram you can narrate step by step, export as a GIF, and keep as a portable file. Local-first, no account required.
Documentation ·
Get started ·
Agent tools ·
Blog ·
License
The agent drives the editor. No API keys, no screenshots, no simulated clicks — it calls the editor's own tools.
The result. A sequence diagram with connector motion and a narrated step order
Install, then describe your diagram
Works with Claude Code, Codex CLI, Cursor, and OpenCode.
npx skills add drawcms/drawcms-skill -g
Send this to your agent:
Use DrawCMS to diagram a checkout request: the browser hits the API gateway,
the gateway publishes to a queue, a worker writes to Postgres, and the
response flows back. Animate the request path and narrate it in three steps.
Then keep going — “add a Redis cache”, “highlight the failure path”, or “make the queue pulse slower”.
No repository required. Start from a description, or point the agent at a repository for a source-backed architecture diagram.
Build the structure
Explain the movement
Narrate the story
75 palette elements across 17 diagram types — architecture, sequence, UML, ER, BPMN, data flow, network, and more.
Attach motion presets to nodes and connectors independently, so the diagram shows how data actually travels.
Turn a static picture into an ordered walkthrough: each step highlights its own nodes and connectors.
Core concepts ↗
Motion presets ↗
Presentation steps ↗
Every diagram is one portable .drawcms file. Export it as PNG or an animated GIF, or hand the file to someone else and they get the motion and the narration with it.
The editor engine and the web application are a single Next.js project. Clone it, install it, run it — that is the entire setup.
git clone https://github.com/drawcms/drawcms.git drawcms
cd drawcms
npm install
npm run dev
Open http://localhost:3002 . The app is the editor and redirects there from every route.
17 diagram types, their element sets, and the templates that ship with them
Picking a diagram type in the elements panel narrows the palette to the elements that notation actually uses, so a sequence diagram offers lifelines and activations rather than BPMN gateways.
Eighteen templates open with motion presets and presentation steps already applied, so a new document is animated before you touch it:
Architecture request flow · Deployment pipeline · Incident timeline · Secure sign-in sequence · Flowchart · Mind map · Org chart · Entity relationship diagram · Data flow diagram · Timeline · Class diagram · State diagram · Deployment diagram · Component diagram · Use case diagram · Network diagram · Activity diagram · User flow diagram
Alongside the notation sets, the palette carries AWS, GCP, Azure, and infrastructure icons (Docker, Kubernetes, Redis, PostgreSQL, and more), plus container elements — groups, regions, security and trust boundaries, processing stages, sequence frames, swimlanes, and BPMN pools — that accept dragged-in children.
Motion is the point
The story is authored, not inferred
A static diagram shows topology. Animation shows behaviour: which way a request travels, what waits, what retries.
You decide the order and the wording of each step. Nothing is guessed from the graph.
Local-first and portable
Agents are first-class
Autosave to browser storage, explicit save when you want it, and one self-contained .drawcms file. No account, no server.
The WebMCP toolset is part of the editor, so an agent co-edits the canvas you are looking at.
The engineering behind it
One app, no package dance — the engine and the host live in the same Next.js project. There is no build-then-link step between them.
A documented boundary — both layers talk through the editor's public API ( src/editor/index.ts ) and a persistence boundary, so storage backends swap without touching canvas code.
A visual grammar, not a shape list — elements carry semantic purpose, suitable and unsuitable uses, compatible diagram types, and motion guidance. Agents query it before they draw, and drawcms_validate_diagram checks the result against it.
Undoable agent edits — drawcms_edit_diagram applies a batch as a single history entry, reversible with one Mod + Z . Agent work and manual work share one undo stack.
Motion separate from structure — retiming a diagram never rewrites its nodes, and rewriting narration never touches motion.
Accessibility as a constraint — keyboard-complete chrome, named controls, and prefers-reduced-motion respected by every preset.
Deliberately webpack — dev and build use webpack rather than Turbopack, on purpose.
DrawCMS is not a general-purpose drawing tool and not a Mermaid renderer. It exists to turn a system you understand into something someone else can follow.
From description to exported artifact
Step
What happens
Describe
You prompt an agent, or place elements yourself.
Recommend
drawcms_recommend_visuals maps entity roles and relationship meanings onto elements, connector types, routing, and motion presets.
Build
drawcms_replace_diagram lays the whole diagram out — lifeline columns for sequences, ranked layers for flowcharts and architecture, obstacle-aware connector routes.
Refine
drawcms_edit_diagram and drawcms_set_motion adjust structure and timing incrementally, each as one undoable batch.
Narrate
drawcms_set_story writes the scenes and steps, or you right-click a selection and add it as a step.
Validate
drawcms_validate_diagram reports unregistered elements, shape-purpose mismatches, unsuitable motion, and choreography problems.
Ship
Export PNG or animated GIF, or save the .drawcms document.
Build diagrams with an AI agent (WebMCP)
The editor registers a WebMCP toolset with the browser's native navigator.modelContext API. In an agent-capable browser — ChatGPT's agentic browsing, or any WebMCP-compatible host — the agent creates animated diagrams directly on the live canvas. No API keys, no servers, no simulated pointer input. Browsers without WebMCP render the ordinary editor.
Because the agent edits the same canvas you do, working together is co-editing rather than file handoff. Agent authoring details live in docs/webmcp.md ; the CLI skill is documented in docs/agent-skill.md .
Mod is Cmd on Apple platforms and Ctrl elsewhere. Menus format their hints from the same table that binds the keys, so a label and its binding cannot drift.
src/
app/ Next.js application shell — layout, editor route, theme
toggle, local persistence wiring (the "host" layer)
editor/ The editor engine — canvas, element catalog, document
format, motion, presentation, persistence boundary, WebMCP
tools, templates, and the test suite
public/ Static assets (cloud provider icons, GIF worker)
docs/ Documentation content, rendered by site/
blog/ Blog posts, rendered by site-blog/
src/app wires the engine into a full product: autosave with a save-status pill, theming, and the “Made with DrawCMS” badge.
React 19 · React Flow v12 · GSAP · Next.js 16 · Tailwind CSS v4 · Vitest
npm run dev # editor on :3002 with hot reload
npm run ci # lint → format:check → typecheck → test → build → site builds
npm run docs:dev # documentation site on :4321
npm run blog:dev # blog on :4322
CI runs exactly npm ci && npm run ci on Node 24. The suite is 512 tests.
Guides and release notes live at drawcms.com/docs , the blog at drawcms.com/blog .
The docs and blog are part of this repository. Pages live in docs/ , posts in blog/ . Two Astro sites render them: site/ (docs, served at drawcms.com/docs ) and site-blog/ (blog, served at drawcms.com/blog ), both deployed as static assets on Cloudflare Workers next to the app worker and routed by path.
A merged pull request touching the sites or their content publishes. Blog posts support draft: true for merge-before-publish workflows. Deploys run through GitHub Actions — set the CLOUDFLARE_API_TOKEN and CLOUDFLARE_ACCOUNT_ID repository secrets — or manually with npm run docs:deploy and npm run blog:deploy . The main app deploys separately and never rebuilds for docs-only changes.
Quick start · Core concepts · Document format
Agent skill · WebMCP · Plugin and host API · API versioning
Self-hosting · Cloud · Upgrading
Accessibility · Browser support · Performance · Design system
Importer limitations — .drawio and .excalidraw imports produce a non-blocking report of what could not be carried across
Real-time multiplayer editing, hosted sharing, and automatic Mermaid parsing are intentionally outside the current scope of this repository. Hosted collaboration lives in DrawCMS Cloud.
Contributions are welcome after accepting the Contributor License Agreement , which enables the project's dual AGPL and commercial licensing. Read CONTRIBUTING.md before opening a pull request.
The official repository is drawcms/drawcms — please open issues and pull requests here. ( dimasna/drawcms-app is a hackathon-purpose repo only.)
Please report security issues privately using the process in SECURITY.md .
DrawCMS is open-source software licensed under GNU AGPL v3.0 only . If you modify it and make that version available to users over a network, the AGPL generally requires you to offer those users the corresponding source code under the AGPL.
The AGPL does not prohibit a compliant competing hosted service. The DrawCMS name and branding are covered separately by TRADEMARKS.md . Organ

[truncated]

## Original Extract

The animated diagram editor AI agents can draw on. Animated technical diagrams (architecture, sequence, BPMN, cloud icons) built by you and your agent together on the same live canvas, via WebMCP. Open-source, local-first. - drawcms/drawcms

Hello, I'm building DrawCMS. It's an open source animated diagram editor you can self-host. You can ask a compatible AI agent to read your repository and create an editable diagram, then walk through the flow step by step on the canvas. AI makes it easy to write code quickly. Keeping track of how that code works can be harder. You finish a feature, move on, and a week later you're trying to remember where a request goes, which service writes to the database, or why a background job runs. It supports architecture, sequence, network and etc diagrams, infrastructure icons also . You can also draw and edit everything manually. Repo: https://github.com/drawcms/drawcms Quickstart (Claude Code, Codex, OpenClaw, Hermes): npx skills add drawcms/drawcms-skill -g Thanks for reading, and I hope you try it out! I would love to hear any feedback

GitHub - drawcms/drawcms: The animated diagram editor AI agents can draw on. Animated technical diagrams (architecture, sequence, BPMN, cloud icons) built by you and your agent together on the same live canvas, via WebMCP. Open-source, local-first. · GitHub
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
drawcms
/
drawcms
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
58 Commits 58 Commits Folders and files
.github/ workflows .github/ workflows adr adr blog blog docs docs public public site-blog site-blog site site src src .gitignore .gitignore .prettierignore .prettierignore .prettierrc.json .prettierrc.json CLA.md CLA.md COMMERCIAL-LICENSE.md COMMERCIAL-LICENSE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md TRADEMARKS.md TRADEMARKS.md eslint.config.mjs eslint.config.mjs netlify.toml netlify.toml next.config.ts next.config.ts open-next.config.ts open-next.config.ts package-lock.json package-lock.json package.json package.json postcss.config.mjs postcss.config.mjs render.yaml render.yaml tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts wrangler.jsonc wrangler.jsonc View all files Repository files navigation
Diagrams that explain how a system moves, not just how it is wired.
Describe a flow to your AI agent, or draw it yourself. DrawCMS turns it into an animated technical diagram you can narrate step by step, export as a GIF, and keep as a portable file. Local-first, no account required.
Documentation ·
Get started ·
Agent tools ·
Blog ·
License
The agent drives the editor. No API keys, no screenshots, no simulated clicks — it calls the editor's own tools.
The result. A sequence diagram with connector motion and a narrated step order
Install, then describe your diagram
Works with Claude Code, Codex CLI, Cursor, and OpenCode.
npx skills add drawcms/drawcms-skill -g
Send this to your agent:
Use DrawCMS to diagram a checkout request: the browser hits the API gateway,
the gateway publishes to a queue, a worker writes to Postgres, and the
response flows back. Animate the request path and narrate it in three steps.
Then keep going — “add a Redis cache”, “highlight the failure path”, or “make the queue pulse slower”.
No repository required. Start from a description, or point the agent at a repository for a source-backed architecture diagram.
Build the structure
Explain the movement
Narrate the story
75 palette elements across 17 diagram types — architecture, sequence, UML, ER, BPMN, data flow, network, and more.
Attach motion presets to nodes and connectors independently, so the diagram shows how data actually travels.
Turn a static picture into an ordered walkthrough: each step highlights its own nodes and connectors.
Core concepts ↗
Motion presets ↗
Presentation steps ↗
Every diagram is one portable .drawcms file. Export it as PNG or an animated GIF, or hand the file to someone else and they get the motion and the narration with it.
The editor engine and the web application are a single Next.js project. Clone it, install it, run it — that is the entire setup.
git clone https://github.com/drawcms/drawcms.git drawcms
cd drawcms
npm install
npm run dev
Open http://localhost:3002 . The app is the editor and redirects there from every route.
17 diagram types, their element sets, and the templates that ship with them
Picking a diagram type in the elements panel narrows the palette to the elements that notation actually uses, so a sequence diagram offers lifelines and activations rather than BPMN gateways.
Eighteen templates open with motion presets and presentation steps already applied, so a new document is animated before you touch it:
Architecture request flow · Deployment pipeline · Incident timeline · Secure sign-in sequence · Flowchart · Mind map · Org chart · Entity relationship diagram · Data flow diagram · Timeline · Class diagram · State diagram · Deployment diagram · Component diagram · Use case diagram · Network diagram · Activity diagram · User flow diagram
Alongside the notation sets, the palette carries AWS, GCP, Azure, and infrastructure icons (Docker, Kubernetes, Redis, PostgreSQL, and more), plus container elements — groups, regions, security and trust boundaries, processing stages, sequence frames, swimlanes, and BPMN pools — that accept dragged-in children.
Motion is the point
The story is authored, not inferred
A static diagram shows topology. Animation shows behaviour: which way a request travels, what waits, what retries.
You decide the order and the wording of each step. Nothing is guessed from the graph.
Local-first and portable
Agents are first-class
Autosave to browser storage, explicit save when you want it, and one self-contained .drawcms file. No account, no server.
The WebMCP toolset is part of the editor, so an agent co-edits the canvas you are looking at.
The engineering behind it
One app, no package dance — the engine and the host live in the same Next.js project. There is no build-then-link step between them.
A documented boundary — both layers talk through the editor's public API ( src/editor/index.ts ) and a persistence boundary, so storage backends swap without touching canvas code.
A visual grammar, not a shape list — elements carry semantic purpose, suitable and unsuitable uses, compatible diagram types, and motion guidance. Agents query it before they draw, and drawcms_validate_diagram checks the result against it.
Undoable agent edits — drawcms_edit_diagram applies a batch as a single history entry, reversible with one Mod + Z . Agent work and manual work share one undo stack.
Motion separate from structure — retiming a diagram never rewrites its nodes, and rewriting narration never touches motion.
Accessibility as a constraint — keyboard-complete chrome, named controls, and prefers-reduced-motion respected by every preset.
Deliberately webpack — dev and build use webpack rather than Turbopack, on purpose.
DrawCMS is not a general-purpose drawing tool and not a Mermaid renderer. It exists to turn a system you understand into something someone else can follow.
From description to exported artifact
Step
What happens
Describe
You prompt an agent, or place elements yourself.
Recommend
drawcms_recommend_visuals maps entity roles and relationship meanings onto elements, connector types, routing, and motion presets.
Build
drawcms_replace_diagram lays the whole diagram out — lifeline columns for sequences, ranked layers for flowcharts and architecture, obstacle-aware connector routes.
Refine
drawcms_edit_diagram and drawcms_set_motion adjust structure and timing incrementally, each as one undoable batch.
Narrate
drawcms_set_story writes the scenes and steps, or you right-click a selection and add it as a step.
Validate
drawcms_validate_diagram reports unregistered elements, shape-purpose mismatches, unsuitable motion, and choreography problems.
Ship
Export PNG or animated GIF, or save the .drawcms document.
Build diagrams with an AI agent (WebMCP)
The editor registers a WebMCP toolset with the browser's native navigator.modelContext API. In an agent-capable browser — ChatGPT's agentic browsing, or any WebMCP-compatible host — the agent creates animated diagrams directly on the live canvas. No API keys, no servers, no simulated pointer input. Browsers without WebMCP render the ordinary editor.
Because the agent edits the same canvas you do, working together is co-editing rather than file handoff. Agent authoring details live in docs/webmcp.md ; the CLI skill is documented in docs/agent-skill.md .
Mod is Cmd on Apple platforms and Ctrl elsewhere. Menus format their hints from the same table that binds the keys, so a label and its binding cannot drift.
src/
app/ Next.js application shell — layout, editor route, theme
toggle, local persistence wiring (the "host" layer)
editor/ The editor engine — canvas, element catalog, document
format, motion, presentation, persistence boundary, WebMCP
tools, templates, and the test suite
public/ Static assets (cloud provider icons, GIF worker)
docs/ Documentation content, rendered by site/
blog/ Blog posts, rendered by site-blog/
src/app wires the engine into a full product: autosave with a save-status pill, theming, and the “Made with DrawCMS” badge.
React 19 · React Flow v12 · GSAP · Next.js 16 · Tailwind CSS v4 · Vitest
npm run dev # editor on :3002 with hot reload
npm run ci # lint → format:check → typecheck → test → build → site builds
npm run docs:dev # documentation site on :4321
npm run blog:dev # blog on :4322
CI runs exactly npm ci && npm run ci on Node 24. The suite is 512 tests.
Guides and release notes live at drawcms.com/docs , the blog at drawcms.com/blog .
The docs and blog are part of this repository. Pages live in docs/ , posts in blog/ . Two Astro sites render them: site/ (docs, served at drawcms.com/docs ) and site-blog/ (blog, served at drawcms.com/blog ), both deployed as static assets on Cloudflare Workers next to the app worker and routed by path.
A merged pull request touching the sites or their content publishes. Blog posts support draft: true for merge-before-publish workflows. Deploys run through GitHub Actions — set the CLOUDFLARE_API_TOKEN and CLOUDFLARE_ACCOUNT_ID repository secrets — or manually with npm run docs:deploy and npm run blog:deploy . The main app deploys separately and never rebuilds for docs-only changes.
Quick start · Core concepts · Document format
Agent skill · WebMCP · Plugin and host API · API versioning
Self-hosting · Cloud · Upgrading
Accessibility · Browser support · Performance · Design system
Importer limitations — .drawio and .excalidraw imports produce a non-blocking report of what could not be carried across
Real-time multiplayer editing, hosted sharing, and automatic Mermaid parsing are intentionally outside the current scope of this repository. Hosted collaboration lives in DrawCMS Cloud.
Contributions are welcome after accepting the Contributor License Agreement , which enables the project's dual AGPL and commercial licensing. Read CONTRIBUTING.md before opening a pull request.
The official repository is drawcms/drawcms — please open issues and pull requests here. ( dimasna/drawcms-app is a hackathon-purpose repo only.)
Please report security issues privately using the process in SECURITY.md .
DrawCMS is open-source software licensed under GNU AGPL v3.0 only . If you modify it and make that version available to users over a network, the AGPL generally requires you to offer those users the corresponding source code under the AGPL.
The AGPL does not prohibit a compliant competing hosted service. The DrawCMS name and branding are covered separately by TRADEMARKS.md . Organ

[truncated]
