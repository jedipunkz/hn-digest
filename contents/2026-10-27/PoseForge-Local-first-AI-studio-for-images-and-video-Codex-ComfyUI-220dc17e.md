---
source: "https://github.com/vishwakulkarni/PoseForge"
hn_url: "https://news.ycombinator.com/item?id=50259155"
title: "PoseForge – Local-first AI studio for images and video (Codex, ComfyUI)"
article_title: "GitHub - vishwakulkarni/PoseForge: Local-first AI pose transformation studio. Transform identity photos using Codex CLI, Google Antigravity, or ComfyUI. · GitHub"
image: "https://opengraph.githubassets.com/6c509f51422ce4a3142133d65ef08d83eb27146639aac6db98b45299821a28ec/vishwakulkarni/PoseForge"
author: "vishwakulkarni"
captured_at: "2026-10-27T14:02:24Z"
capture_tool: "hn-digest"
hn_id: 50259155
score: 1
comments: 0
posted_at: "2026-10-27T13:40:28Z"
tags:
  - hacker-news
---

# PoseForge – Local-first AI studio for images and video (Codex, ComfyUI)

- HN: [50259155](https://news.ycombinator.com/item?id=50259155)
- Source: [github.com](https://github.com/vishwakulkarni/PoseForge)
- Score: 1
- Comments: 0
- Posted: 2026-10-27T13:40:28Z

## Translation

Title: PoseForge – Local-first AI studio for images and video (Codex, ComfyUI)
Article title: GitHub - vishwakulkarni/PoseForge: Local-first AI pose transformation studio. Transform identity photos using Codex CLI, Google Antigravity, or ComfyUI. · GitHub
Description: Local-first AI pose transformation studio. Transform identity photos using Codex CLI, Google Antigravity, or ComfyUI. - vishwakulkarni/PoseForge

Article text:
GitHub - vishwakulkarni/PoseForge: Local-first AI pose transformation studio. Transform identity photos using Codex CLI, Google Antigravity, or ComfyUI. · GitHub
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
vishwakulkarni
PoseForge Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Local-first AI pose transformation studio. Transform identity photos using Codex CLI, Google Antigravity, or ComfyUI.
vishwakulkarni.github.io/PoseForge/ Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
106 Commits 106 Commits Folders and files
.github .github .serena .serena db db demo-output demo-output docs docs engines engines lib lib public/ images public/ images research research routes routes scripts scripts social social storage/ pose-library/ seed storage/ pose-library/ seed tests tests web web .env.example .env.example .gitignore .gitignore ADVANCED-STUDIO-PLAN.md ADVANCED-STUDIO-PLAN.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md CREDITS.md CREDITS.md IMPLEMENTATION-PLAN-PoseForge-Fluid-Studio.md IMPLEMENTATION-PLAN-PoseForge-Fluid-Studio.md LAUNCH_CHECKLIST.md LAUNCH_CHECKLIST.md LAUNCH_POSTS.md LAUNCH_POSTS.md LICENSE LICENSE PRD-PoseForge-Fluid-Studio.md PRD-PoseForge-Fluid-Studio.md PRIVACY.md PRIVACY.md README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md index.ts index.ts launch-readiness.json launch-readiness.json package-lock.json package-lock.json package.json package.json server.js server.js studio-changes.md studio-changes.md View all files Repository files navigation
Open the landing page and documentation · The full Studio runs locally because generation, storage, and its embedded database are intentionally not hosted on GitHub Pages.
Turn the AI subscription you already pay for into a private, local-first photo studio
Give PoseForge an identity photo and a pose reference. It returns the same
person—or family—in the new pose and composition, with controls for camera,
lighting, styling, fidelity, and output format.
PoseForge exists so Codex, Google Antigravity, and ComfyUI users can turn their
existing AI access into a repeatable visual workflow instead of rebuilding the
same multi-image prompt in chat for every photograph.
Watch the 15-second generation walkthrough ·
Download the lightweight 10-second MP4 ·
Quickstart ·
How it works ·
Examples ·
Landing page source ·
Documentation ·
Roadmap ·
Support ·
Discussions
Local-first, not automatically offline. Your workspace, embedded PGlite
database, character library, pose library, and generated files stay on your
machine. ComfyUI can keep inference fully local. Codex CLI, Google
Antigravity, and hosted API engines send only the references, prompt, and
settings selected for that generation to their provider. See
Privacy and data flow .
Prerequisite: Node.js 20.9+. PGlite is embedded, so Docker and a separate database server are not required.
git clone https://github.com/vishwakulkarni/PoseForge.git
cd PoseForge
npm install
npm run dev
Open http://localhost:3000 . npm install installs both
the API and web dependencies. The embedded PGlite database is created and
migrated automatically on first start. Settings can be configured in the app;
copy .env.example to .env only if you want to set options such as API keys
up front. To check a fresh install end to end (database and the 16 bundled
offline poses), run npm run setup .
No image engine is required to explore the interface. To generate images,
authenticate Codex or Google Antigravity, connect local ComfyUI, or configure
a hosted provider in Settings .
Choose the identity. Upload one to four people, or reuse saved character
references.
Choose the pose. Upload a pose, split a pose sheet, or select one of the
bundled references.
Direct and generate. Set composition, camera, lighting, styling, and
fidelity; then run through Codex, Antigravity, ComfyUI, or a hosted API.
PoseForge stores successful pose uploads in the reusable pose library and
keeps generation history, usage, latency, cost estimates, and outputs together
on the local machine.
Consistent characters, on any provider
PoseForge keeps the parts that make a result look like the same person on your
machine, then hands each provider the same carefully prepared inputs:
Characters live locally. Saved character references, per-character angle
profiles, your pose library, and every generation are stored in the embedded
PGlite database and local image storage, not in a provider account.
A local pipeline does the preparation. References are normalized, poses can
be split from multi-pose sheets, and the multi-image prompt is assembled with
your identity, pose, camera, lighting, and fidelity controls before any engine
is called, so engines receive consistent, well-formed input.
The same characters carry through images and videos. Reuse one saved
character across Studio stills and Advanced Studio image and video nodes, with
up to four people per generation, instead of re-describing them each time.
Switch providers without starting over. Characters, poses, history,
projects, and metrics all stay put when you change engine, so you can try
Codex, Antigravity, Dreamina, Higgsfield, ComfyUI, or a hosted API and compare
results from identical inputs.
Metrics are calculated locally. Usage, cost, latency, and reliability come
from your own stored runs; nothing is sent to a PoseForge cloud, because there
isn't one.
Output quality still depends on the engine and model you choose; PoseForge's job
is to give each one the best, most consistent starting point.
Choose where the image is made
Consumer subscriptions and API billing are often separate products. Signed-in
CLI behavior, quotas, and retention remain subject to the provider account and
terms. Database-backed API keys and ComfyUI workflow JSON are stored in plain
text in the local database; prefer environment variables where supported and
read SECURITY.md before use on a shared machine.
Studio supports seven interchangeable engines. Advanced Studio adds two
more providers, Higgsfield and Dreamina , for image and video nodes, and
uses fal.ai for video too:
Higgsfield and Dreamina usage is metered in credits and shown on the
Metrics page next to token and dollar spend. Selected inputs leave your
machine for both providers. Adding another engine is a small adapter change; see
the engine interface .
The people below are fictional, AI-created demo subjects. These visuals show
the intended character → pose → result contract without publishing private
user photos. They are product demonstrations, not provider quality benchmarks.
Individual editorial pose transfer
Identity stays recognizable while the reference contributes body position,
camera framing, and composition; the final scene adds new lighting and art
direction.
Indian family using an American family pose
The middle photograph supplies only the arrangement: parents seated on either
side while their two-year-old stands between them. The result keeps the Indian
family identities and applies that family pose in a new festive setting.
Reproducible 10-second demo reel
The README animation combines both transformations into a character → pose →
result story. AGY assembled the local media pipeline; sharp renders 300
frames and FFmpeg creates the MP4 and optimized GIF without network assets.
For generation benchmarks, record the exact engine, model, prompt, runtime,
local hardware, and provider date. Do not compare providers using unrecorded
launch assets or publish identifiable photos without permission.
Studio, Advanced Studio, and Metrics
Studio: character → pose → result
The guided workbench. Pick up to four identity slots, a pose reference (or split
a multi-pose sheet), then direct per-person styling, camera, lighting, and
fidelity. A live workflow canvas shows what feeds the generation, recipes save
your settings, and up to six variants can be queued at once. Normal mode keeps
the controls compact; Advanced mode adds per-person direction and identity and pose
fidelity controls.
Advanced Studio: free-form node workflows
A multi-project node canvas for workflows that do not start from a character
and a pose. Start from a template or a blank canvas, wire nodes together, and
run them from the canvas:
Node types include text/prompt, image and media inputs, image and video
generators, and groups. Generator options (aspect ratios, resolutions,
reference-image limits) come from what each engine can actually do rather than
a global list. Generator nodes can run on Higgsfield and Dreamina as well as
the Studio engines (see below). Projects autosave, edges can be cut, and the canvas has
keyboard shortcuts and a selection toolbar. Open it at /studio-advanced .
The bundled "Example: baby running" project: prompts feed image generators, and
their outputs chain into video generators. Fictional, AI-created subjects.
Metrics: know what each generation cost
The Metrics page summarizes every generation stored locally: lifetime spend,
tokens, success rate, median and P95 latency, queue wait, and active days. It
adds a cumulative usage chart (group by day, week, month, or run; tokens or
cost), a per-engine and per-model breakdown, recurring failure groups, and
Codex and Antigravity account-limit windows. Switch between Historical and
Session scope, and export JSON or CSV. Values are marked actual, estimated,
partial, or local, so you can tell provider-reported usage from estimates; see
the metrics reference .
A read-only static demo of the whole app, including sample Metrics data, is
published from main to GitHub Pages: https://vishwakulkarni.github.io/PoseForge/ .
Generation is disabled there; run PoseForge locally to create images.
Studio: guided character → pose → result workbench described above.
Advanced Studio: node-canvas image, video, and storyboard workflows.
Characters and poses: reusable local libraries with normalized image
storage, pose filtering, and best-effort tagging.
History: generation details, inputs, outputs, deletion, and rerun context.
ID Photos: local formatting for U.S. and Indian passport, visa, e-Visa,
and OCI profiles, with optional AI assistance kept separate.
Metrics: usage, spend, latency, reliability, and exports described above.
Documentation: repository Markdown rendered inside the application.
PoseForge runs as one Node.js server. Express owns /api , /storage , database
rules, generation queues, and engine adapters; Next.js owns the interface. The
browser talks to one lo

[truncated]

## Original Extract

Local-first AI pose transformation studio. Transform identity photos using Codex CLI, Google Antigravity, or ComfyUI. - vishwakulkarni/PoseForge

GitHub - vishwakulkarni/PoseForge: Local-first AI pose transformation studio. Transform identity photos using Codex CLI, Google Antigravity, or ComfyUI. · GitHub
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
vishwakulkarni
PoseForge Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Local-first AI pose transformation studio. Transform identity photos using Codex CLI, Google Antigravity, or ComfyUI.
vishwakulkarni.github.io/PoseForge/ Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
106 Commits 106 Commits Folders and files
.github .github .serena .serena db db demo-output demo-output docs docs engines engines lib lib public/ images public/ images research research routes routes scripts scripts social social storage/ pose-library/ seed storage/ pose-library/ seed tests tests web web .env.example .env.example .gitignore .gitignore ADVANCED-STUDIO-PLAN.md ADVANCED-STUDIO-PLAN.md ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md CREDITS.md CREDITS.md IMPLEMENTATION-PLAN-PoseForge-Fluid-Studio.md IMPLEMENTATION-PLAN-PoseForge-Fluid-Studio.md LAUNCH_CHECKLIST.md LAUNCH_CHECKLIST.md LAUNCH_POSTS.md LAUNCH_POSTS.md LICENSE LICENSE PRD-PoseForge-Fluid-Studio.md PRD-PoseForge-Fluid-Studio.md PRIVACY.md PRIVACY.md README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md index.ts index.ts launch-readiness.json launch-readiness.json package-lock.json package-lock.json package.json package.json server.js server.js studio-changes.md studio-changes.md View all files Repository files navigation
Open the landing page and documentation · The full Studio runs locally because generation, storage, and its embedded database are intentionally not hosted on GitHub Pages.
Turn the AI subscription you already pay for into a private, local-first photo studio
Give PoseForge an identity photo and a pose reference. It returns the same
person—or family—in the new pose and composition, with controls for camera,
lighting, styling, fidelity, and output format.
PoseForge exists so Codex, Google Antigravity, and ComfyUI users can turn their
existing AI access into a repeatable visual workflow instead of rebuilding the
same multi-image prompt in chat for every photograph.
Watch the 15-second generation walkthrough ·
Download the lightweight 10-second MP4 ·
Quickstart ·
How it works ·
Examples ·
Landing page source ·
Documentation ·
Roadmap ·
Support ·
Discussions
Local-first, not automatically offline. Your workspace, embedded PGlite
database, character library, pose library, and generated files stay on your
machine. ComfyUI can keep inference fully local. Codex CLI, Google
Antigravity, and hosted API engines send only the references, prompt, and
settings selected for that generation to their provider. See
Privacy and data flow .
Prerequisite: Node.js 20.9+. PGlite is embedded, so Docker and a separate database server are not required.
git clone https://github.com/vishwakulkarni/PoseForge.git
cd PoseForge
npm install
npm run dev
Open http://localhost:3000 . npm install installs both
the API and web dependencies. The embedded PGlite database is created and
migrated automatically on first start. Settings can be configured in the app;
copy .env.example to .env only if you want to set options such as API keys
up front. To check a fresh install end to end (database and the 16 bundled
offline poses), run npm run setup .
No image engine is required to explore the interface. To generate images,
authenticate Codex or Google Antigravity, connect local ComfyUI, or configure
a hosted provider in Settings .
Choose the identity. Upload one to four people, or reuse saved character
references.
Choose the pose. Upload a pose, split a pose sheet, or select one of the
bundled references.
Direct and generate. Set composition, camera, lighting, styling, and
fidelity; then run through Codex, Antigravity, ComfyUI, or a hosted API.
PoseForge stores successful pose uploads in the reusable pose library and
keeps generation history, usage, latency, cost estimates, and outputs together
on the local machine.
Consistent characters, on any provider
PoseForge keeps the parts that make a result look like the same person on your
machine, then hands each provider the same carefully prepared inputs:
Characters live locally. Saved character references, per-character angle
profiles, your pose library, and every generation are stored in the embedded
PGlite database and local image storage, not in a provider account.
A local pipeline does the preparation. References are normalized, poses can
be split from multi-pose sheets, and the multi-image prompt is assembled with
your identity, pose, camera, lighting, and fidelity controls before any engine
is called, so engines receive consistent, well-formed input.
The same characters carry through images and videos. Reuse one saved
character across Studio stills and Advanced Studio image and video nodes, with
up to four people per generation, instead of re-describing them each time.
Switch providers without starting over. Characters, poses, history,
projects, and metrics all stay put when you change engine, so you can try
Codex, Antigravity, Dreamina, Higgsfield, ComfyUI, or a hosted API and compare
results from identical inputs.
Metrics are calculated locally. Usage, cost, latency, and reliability come
from your own stored runs; nothing is sent to a PoseForge cloud, because there
isn't one.
Output quality still depends on the engine and model you choose; PoseForge's job
is to give each one the best, most consistent starting point.
Choose where the image is made
Consumer subscriptions and API billing are often separate products. Signed-in
CLI behavior, quotas, and retention remain subject to the provider account and
terms. Database-backed API keys and ComfyUI workflow JSON are stored in plain
text in the local database; prefer environment variables where supported and
read SECURITY.md before use on a shared machine.
Studio supports seven interchangeable engines. Advanced Studio adds two
more providers, Higgsfield and Dreamina , for image and video nodes, and
uses fal.ai for video too:
Higgsfield and Dreamina usage is metered in credits and shown on the
Metrics page next to token and dollar spend. Selected inputs leave your
machine for both providers. Adding another engine is a small adapter change; see
the engine interface .
The people below are fictional, AI-created demo subjects. These visuals show
the intended character → pose → result contract without publishing private
user photos. They are product demonstrations, not provider quality benchmarks.
Individual editorial pose transfer
Identity stays recognizable while the reference contributes body position,
camera framing, and composition; the final scene adds new lighting and art
direction.
Indian family using an American family pose
The middle photograph supplies only the arrangement: parents seated on either
side while their two-year-old stands between them. The result keeps the Indian
family identities and applies that family pose in a new festive setting.
Reproducible 10-second demo reel
The README animation combines both transformations into a character → pose →
result story. AGY assembled the local media pipeline; sharp renders 300
frames and FFmpeg creates the MP4 and optimized GIF without network assets.
For generation benchmarks, record the exact engine, model, prompt, runtime,
local hardware, and provider date. Do not compare providers using unrecorded
launch assets or publish identifiable photos without permission.
Studio, Advanced Studio, and Metrics
Studio: character → pose → result
The guided workbench. Pick up to four identity slots, a pose reference (or split
a multi-pose sheet), then direct per-person styling, camera, lighting, and
fidelity. A live workflow canvas shows what feeds the generation, recipes save
your settings, and up to six variants can be queued at once. Normal mode keeps
the controls compact; Advanced mode adds per-person direction and identity and pose
fidelity controls.
Advanced Studio: free-form node workflows
A multi-project node canvas for workflows that do not start from a character
and a pose. Start from a template or a blank canvas, wire nodes together, and
run them from the canvas:
Node types include text/prompt, image and media inputs, image and video
generators, and groups. Generator options (aspect ratios, resolutions,
reference-image limits) come from what each engine can actually do rather than
a global list. Generator nodes can run on Higgsfield and Dreamina as well as
the Studio engines (see below). Projects autosave, edges can be cut, and the canvas has
keyboard shortcuts and a selection toolbar. Open it at /studio-advanced .
The bundled "Example: baby running" project: prompts feed image generators, and
their outputs chain into video generators. Fictional, AI-created subjects.
Metrics: know what each generation cost
The Metrics page summarizes every generation stored locally: lifetime spend,
tokens, success rate, median and P95 latency, queue wait, and active days. It
adds a cumulative usage chart (group by day, week, month, or run; tokens or
cost), a per-engine and per-model breakdown, recurring failure groups, and
Codex and Antigravity account-limit windows. Switch between Historical and
Session scope, and export JSON or CSV. Values are marked actual, estimated,
partial, or local, so you can tell provider-reported usage from estimates; see
the metrics reference .
A read-only static demo of the whole app, including sample Metrics data, is
published from main to GitHub Pages: https://vishwakulkarni.github.io/PoseForge/ .
Generation is disabled there; run PoseForge locally to create images.
Studio: guided character → pose → result workbench described above.
Advanced Studio: node-canvas image, video, and storyboard workflows.
Characters and poses: reusable local libraries with normalized image
storage, pose filtering, and best-effort tagging.
History: generation details, inputs, outputs, deletion, and rerun context.
ID Photos: local formatting for U.S. and Indian passport, visa, e-Visa,
and OCI profiles, with optional AI assistance kept separate.
Metrics: usage, spend, latency, reliability, and exports described above.
Documentation: repository Markdown rendered inside the application.
PoseForge runs as one Node.js server. Express owns /api , /storage , database
rules, generation queues, and engine adapters; Next.js owns the interface. The
browser talks to one lo

[truncated]
