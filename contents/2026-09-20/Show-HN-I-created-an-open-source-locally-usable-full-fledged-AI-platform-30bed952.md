---
source: "https://github.com/theguysudo/ENZO"
hn_url: "https://news.ycombinator.com/item?id=49771118"
title: "Show HN: I created an open source locally usable full fledged AI platform"
article_title: "GitHub - theguysudo/ENZO: Self-hosted AI workspace with agents, skills, and tools (Gmail, Calendar) that runs entirely on your own provider API keys (BYOK). Bring your own keys — Groq, OpenRouter, NVIDIA, Hugging Face, Google AI. · GitHub"
image: "https://opengraph.githubassets.com/fc642ce27217d8b1a825789ee845eb838e5e0d8d04083900aa185862312e9fb9/theguysudo/ENZO"
author: "theguysudo"
captured_at: "2026-09-20T00:43:15Z"
capture_tool: "hn-digest"
hn_id: 49771118
score: 7
comments: 5
posted_at: "2026-09-19T23:47:35Z"
tags:
  - hacker-news
---

# Show HN: I created an open source locally usable full fledged AI platform

- HN: [49771118](https://news.ycombinator.com/item?id=49771118)
- Source: [github.com](https://github.com/theguysudo/ENZO)
- Score: 7
- Comments: 5
- Posted: 2026-09-19T23:47:35Z

## Translation

Title: Show HN: I created an open source locally usable full fledged AI platform
Article title: GitHub - theguysudo/ENZO: Self-hosted AI workspace with agents, skills, and tools (Gmail, Calendar) that runs entirely on your own provider API keys (BYOK). Bring your own keys — Groq, OpenRouter, NVIDIA, Hugging Face, Google AI. · GitHub
Description: Self-hosted AI workspace with agents, skills, and tools (Gmail, Calendar) that runs entirely on your own provider API keys (BYOK). Bring your own keys — Groq, OpenRouter, NVIDIA, Hugging Face, Google AI. - theguysudo/ENZO
HN text: hi to all the readers this post is for my recent opensource project called ENZO now answering what is enzo so enzo is an opensource platform where i clubbed all the free available api for anyone use under one hood with more than 2000 models available to use for chatting coding researching and much more now answering the most common question of why you should put your time looking the project so it has few distinct feature meaning it has a dedicated agents tab where you can describe your need and create a special agent just for one specific task with master ability in that domain second it has the ability to connect your gmail drive and calendar and then you can ask it to perform some specific tasks like reading you the most important mail of the day or finding recruiter mails and creating personalized reply based on your data which it stores locally on your device third the coding mode offers a dedicated preview window where you can see your code running and have a look of it feels and edit it in realtime as well as all the modes are packed with dedicated skills which delivers promising results fourth the ui features some additional things such as music tab where you can listen to any music want and it has a custom personalized feature which runs in background and an llm understands your taste and recommends similar kind of music you like fifth the most important why your trust it with your api key then to explain i would say enzo a dedicated vault which manages all your api and to secure it the vault as aes 256 bit encryption which prevents any person or any middle man to look at your api key and since the whole program runs locally on your device you have complete freedom to oversee all the backend work happening and it also features password lock which if you enable saves a backup key and then locks your whole platform work behind a pass screen though it is not foolproof as any third party or malware containing extension can still fetch login tokens from your browser so its security also depends upon how you access it concluding all of it. i urge to anyone who reads this to have a look at the platform even if you hate it just curse it in the comment its fine or if you would like to drop any feedback i would highly encourage that and since its my first work open source platform i know it has a lot of errors and bugs so i apologize upfront for it and if you consider my work worthy please drop a star on the repo that'll make my day

Article text:
GitHub - theguysudo/ENZO: Self-hosted AI workspace with agents, skills, and tools (Gmail, Calendar) that runs entirely on your own provider API keys (BYOK). Bring your own keys — Groq, OpenRouter, NVIDIA, Hugging Face, Google AI. · GitHub
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
theguysudo
/
ENZO
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
41 Commits 41 Commits Folders and files
.github/ workflows .github/ workflows docs docs icon icon notebooks notebooks scripts scripts skills-bundled skills-bundled src src synthetic-nature synthetic-nature tests tests traffic traffic .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md TRAFFIC.md TRAFFIC.md docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh index.ts index.ts package-lock.json package-lock.json package.json package.json security-implementation-log.md security-implementation-log.md skills-lock.json skills-lock.json tsconfig.json tsconfig.json View all files Repository files navigation
Quickstart ·
Live demo ·
What's inside ·
Security ·
What's new ·
Usage guide ·
Changelog
Chat with 300+ models. Build agents that write their own operating manuals. Research, generate code, run it all — on your keys, on your infrastructure. When you send a message, the request goes from your browser through ENZO to the provider you picked, and you pay that provider their normal price. Nothing sits in between taking a cut. There is no ENZO account, no usage meter, no subscription.
Inside ENZO
Count
What it gives you
Models in one catalog
300+
across 9 providers, health-checked live
Injectable agent skills
74 bundled
domain playbooks the agent loop pulls in per run
Self-drafting agents
2-pass builder
plain-English task in, operating manual out
CI pipeline stages
7
including a black-box security pentest on every push
Pentest assertions
44
auth bypass, IDOR, hostile payloads, stream integrity
Unit + security tests
298
agent, vault, crypto and model suites
TypeScript (strict)
~44,000 lines
one language, strict mode throughout
Releases
5
v1.0.0 → v1.4.0, everything in the changelog
Quickstart
git clone https://github.com/theguysudo/ENZO.git
cd enzo
docker compose up -d
# → http://localhost:5001
That's the whole install. No accounts, no mandatory env, no database server. Open the app, press Login , and pick any provider:
Keys are saved encrypted in your browser (passphrase-protected vault, with a recovery file you can download). You can wipe them anytime from the Vault.
On a fresh self-hosted instance the first live-validated key you paste claims the instance — it's written to the container .env and sealed into the enzo-memory volume, so every server-side feature (agents, skills, memory) unlocks immediately and survives restarts. No master key to configure, no setup wizard — paste a working key and go. (Pre-seed a provider key in compose env instead if you'd rather not have the claim window at all; the threat model states this trade plainly.)
Try it hosted first: https://enzo-hub.duckdns.org — the same app, running on our infrastructure. This repo is exactly that code, minus Google sign-in (self-hosted login is just your provider keys) with a trimmed default theme set for a small download.
Run on Google Colab — zero install
Don't want the local hassle — or want ENZO reachable from any device? The Colab notebook does the whole setup for you: it clones this repo, installs the dependencies, builds the UI and boots the server, then hands you a URL.
Click the button — the notebook opens in Colab (a free Google account is enough).
Run all cells ( Runtime → Run all , or Ctrl/⌘ + F9 ) — install + build takes ~4–5 min the first time; every cell is idempotent, so a re-run reuses what's already there instead of starting over.
Take your URL — the last cell prints two links: a Colab link that works in the browser you're already in, and a Cloudflare tunnel link that works from your phone or any other device (free, no account). The tunnel URL changes each session; the Colab link is bound to your session.
Stays up for the whole session. The notebook arms two disconnect-prevention mechanisms before handing you the URL: a keep-alive that resets Colab's ~90-minute idle timer every 60 seconds while the tab is open, and a watchdog thread that pings the server every 5 minutes and restarts it automatically if it ever dies. Free Colab caps a session at ~12 hours, so a 6-hour run fits comfortably; when a session does end, one click on Run all brings everything back.
The notebook runs on Google's hardware, so Colab's terms apply while you're there — but your model keys stay yours : paste any provider key after boot, exactly like self-hosting, and nothing is ever stored on our side. When the Colab session ends, everything on the VM is gone.
Surface
What it does
Terminal
Streaming chat with 300+ models — normal, thinking, research and coding modes — with a live ECG-style health trace in the toolbar that flatlines red the moment the catalog is unreachable
Model marketplace
One unified catalog across all 9 providers — cards now carry the platform's own cover art and a research panel with real download counts, licences and benchmarks
Music player
Search any song and play it in the marketplace — keyless, no YouTube API key — with a 5-band equalizer that actually re-shapes the audio
Agent builder
Describe a task in plain English; ENZO drafts the agent's full operating manual, and the agent keeps training itself on your activity from then on
Research mode
A deep-research loop that writes its own queries, reads what it finds, and decides when it's done — under hard budgets so it can't burn your key
Code-gen
Writes a coding project, boots it, previews it live, and tells you when it's broken
Vault
Every key sealed in the browser, attached per-request, wipeable in one click
How the agent builder works
Pass 1 — analysis. A two-pass drafter reads the domain of your task ("an agent that researches MUN country positions") and derives what the manual needs to cover: tacit knowledge, decision heuristics, edge cases.
The race. When a draft needs a model, ~10 free candidates from your own providers fire simultaneously — the first to answer wins, stragglers are aborted, and a brain-health scoreboard reorders future races by which models actually deliver. Dead or rate-limited free-tier models can no longer collapse a draft.
Honest provenance. Every agent records which model actually drafted it — and says so plainly when nothing was reachable.
It doesn't stop. A per-agent neural layer folds in domain-matched platform activity on a 90-second cadence, distills lessons into memory, and injects a live NEURAL FOCUS block into every run. Watch it in the agent's Neural tab.
Most "AI workspaces" hold your keys, meter your usage, or need a subscription to exist. ENZO is built the other way around:
(Competitor column is about the category, not specific products — details vary.)
The full threat model is written down — checkable, with the code that makes each claim true — in docs/SECURITY.md . The short version:
Keys are sealed in your browser with AES-256-GCM under a non-extractable WebCrypto key. It can be used to decrypt your keys while never being copied — no JavaScript can export its bytes, ours or an attacker's. Optional passphrase mode re-seals everything under PBKDF2-SHA256 (600,000 iterations) and deletes the device key entirely.
One module touches key storage, and CI enforces it. A pipeline stage greps the frontend for any raw localStorage key read and fails the build on a hit — a missed key-access site is a red build, never a production bug.
Every push runs a 44-assertion black-box pentest against a booted server — auth bypass, hostile payloads, IDOR, stream integrity — plus a keyless-boot proof: the server must start with zero provider keys. That's the BYOK guarantee, tested, not promised.
The limits are stated up front. Self-hosted mode stores the first key you claim in the container .env (sealed in the memory volume) so scheduled agents can run while your browser is closed — that trade is documented, not hidden. docs/SECURITY.md covers what's protected, what isn't, and why.
In-chat file converter — attach a PDF, spreadsheet, CSV, JSON, TXT or Markdown file in the terminal chat and it's parsed to real text/rows in your browser (files never leave the device; only the reasoning step uses your own key, like normal chat). The agent extracts, merges, or cross-converts — and CSV / Excel download buttons appear right on the reply . Built for research papers: "extract every table and merge into one CSV" now works end to end, and scanned PDFs report their missing text layer instead of failing.
Run it on Google Colab — one click on the Open In Colab button (see Run on Google Colab ): the notebook clones, installs, builds and boots ENZO on Google's hardware, hands you a URL for your browser plus a tunnel URL for your phone, and arms a keep-alive + watchdog so it stays up for 6+ hours.
Self-healing Docker pulls — the container now verifies its dependencies on every start and installs anything missing before the server boots ( ENZO_AUTO_INSTALL=0 to skip). A pulled image can't boot broken.
What's new in v1.3.0 (previous)
Music player — search any song and play it straight from the marketplace, keyless (no YouTube API key, no quota): a collapsed corner pill expands into a full player card — vinyl disc hero, queue walking, shuffle/loop/like, keyboard controls. For You turns your own listening history (kept device-local) into song seeds through your own provider key.
Real equalizer — a 5-band Web Audio EQ (bass / low-mid / mid / presence / air, ±12 dB, preamp, five presets) that genuinely re-shapes the frequency response when you opt in. Enhance starts off — normal playback is untouched. Tracks the enhancer can't stream fall back to the YouTube engine automatically; playback never breaks.
Marketplace, redesigned — every model card now carries the platform's own cover art, brand colour on hover, live health dot + latency, and a research panel with real facts: HuggingFace download counts, licences, knowledge cutoffs, Artificial Analysis scores and a Wikipedia-backed family summary — pulled keyless from public endpoints, never guessed.
NYC Subway theme — the workspace's new flagship backdrop: an AI-animated subway ride through a tunnel, with a handheld-camera tremble, monochrome film grade and animated recording grain added in code (the video ships clean).
A quieter interface — the whole workspace went monochrome + a single coral accent: the terminal toggle switch reb

[truncated]

## Original Extract

Self-hosted AI workspace with agents, skills, and tools (Gmail, Calendar) that runs entirely on your own provider API keys (BYOK). Bring your own keys — Groq, OpenRouter, NVIDIA, Hugging Face, Google AI. - theguysudo/ENZO

hi to all the readers this post is for my recent opensource project called ENZO now answering what is enzo so enzo is an opensource platform where i clubbed all the free available api for anyone use under one hood with more than 2000 models available to use for chatting coding researching and much more now answering the most common question of why you should put your time looking the project so it has few distinct feature meaning it has a dedicated agents tab where you can describe your need and create a special agent just for one specific task with master ability in that domain second it has the ability to connect your gmail drive and calendar and then you can ask it to perform some specific tasks like reading you the most important mail of the day or finding recruiter mails and creating personalized reply based on your data which it stores locally on your device third the coding mode offers a dedicated preview window where you can see your code running and have a look of it feels and edit it in realtime as well as all the modes are packed with dedicated skills which delivers promising results fourth the ui features some additional things such as music tab where you can listen to any music want and it has a custom personalized feature which runs in background and an llm understands your taste and recommends similar kind of music you like fifth the most important why your trust it with your api key then to explain i would say enzo a dedicated vault which manages all your api and to secure it the vault as aes 256 bit encryption which prevents any person or any middle man to look at your api key and since the whole program runs locally on your device you have complete freedom to oversee all the backend work happening and it also features password lock which if you enable saves a backup key and then locks your whole platform work behind a pass screen though it is not foolproof as any third party or malware containing extension can still fetch login tokens from your browser so its security also depends upon how you access it concluding all of it. i urge to anyone who reads this to have a look at the platform even if you hate it just curse it in the comment its fine or if you would like to drop any feedback i would highly encourage that and since its my first work open source platform i know it has a lot of errors and bugs so i apologize upfront for it and if you consider my work worthy please drop a star on the repo that'll make my day

GitHub - theguysudo/ENZO: Self-hosted AI workspace with agents, skills, and tools (Gmail, Calendar) that runs entirely on your own provider API keys (BYOK). Bring your own keys — Groq, OpenRouter, NVIDIA, Hugging Face, Google AI. · GitHub
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
theguysudo
/
ENZO
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
41 Commits 41 Commits Folders and files
.github/ workflows .github/ workflows docs docs icon icon notebooks notebooks scripts scripts skills-bundled skills-bundled src src synthetic-nature synthetic-nature tests tests traffic traffic .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore Dockerfile Dockerfile LICENSE LICENSE README.md README.md TRAFFIC.md TRAFFIC.md docker-compose.yml docker-compose.yml docker-entrypoint.sh docker-entrypoint.sh index.ts index.ts package-lock.json package-lock.json package.json package.json security-implementation-log.md security-implementation-log.md skills-lock.json skills-lock.json tsconfig.json tsconfig.json View all files Repository files navigation
Quickstart ·
Live demo ·
What's inside ·
Security ·
What's new ·
Usage guide ·
Changelog
Chat with 300+ models. Build agents that write their own operating manuals. Research, generate code, run it all — on your keys, on your infrastructure. When you send a message, the request goes from your browser through ENZO to the provider you picked, and you pay that provider their normal price. Nothing sits in between taking a cut. There is no ENZO account, no usage meter, no subscription.
Inside ENZO
Count
What it gives you
Models in one catalog
300+
across 9 providers, health-checked live
Injectable agent skills
74 bundled
domain playbooks the agent loop pulls in per run
Self-drafting agents
2-pass builder
plain-English task in, operating manual out
CI pipeline stages
7
including a black-box security pentest on every push
Pentest assertions
44
auth bypass, IDOR, hostile payloads, stream integrity
Unit + security tests
298
agent, vault, crypto and model suites
TypeScript (strict)
~44,000 lines
one language, strict mode throughout
Releases
5
v1.0.0 → v1.4.0, everything in the changelog
Quickstart
git clone https://github.com/theguysudo/ENZO.git
cd enzo
docker compose up -d
# → http://localhost:5001
That's the whole install. No accounts, no mandatory env, no database server. Open the app, press Login , and pick any provider:
Keys are saved encrypted in your browser (passphrase-protected vault, with a recovery file you can download). You can wipe them anytime from the Vault.
On a fresh self-hosted instance the first live-validated key you paste claims the instance — it's written to the container .env and sealed into the enzo-memory volume, so every server-side feature (agents, skills, memory) unlocks immediately and survives restarts. No master key to configure, no setup wizard — paste a working key and go. (Pre-seed a provider key in compose env instead if you'd rather not have the claim window at all; the threat model states this trade plainly.)
Try it hosted first: https://enzo-hub.duckdns.org — the same app, running on our infrastructure. This repo is exactly that code, minus Google sign-in (self-hosted login is just your provider keys) with a trimmed default theme set for a small download.
Run on Google Colab — zero install
Don't want the local hassle — or want ENZO reachable from any device? The Colab notebook does the whole setup for you: it clones this repo, installs the dependencies, builds the UI and boots the server, then hands you a URL.
Click the button — the notebook opens in Colab (a free Google account is enough).
Run all cells ( Runtime → Run all , or Ctrl/⌘ + F9 ) — install + build takes ~4–5 min the first time; every cell is idempotent, so a re-run reuses what's already there instead of starting over.
Take your URL — the last cell prints two links: a Colab link that works in the browser you're already in, and a Cloudflare tunnel link that works from your phone or any other device (free, no account). The tunnel URL changes each session; the Colab link is bound to your session.
Stays up for the whole session. The notebook arms two disconnect-prevention mechanisms before handing you the URL: a keep-alive that resets Colab's ~90-minute idle timer every 60 seconds while the tab is open, and a watchdog thread that pings the server every 5 minutes and restarts it automatically if it ever dies. Free Colab caps a session at ~12 hours, so a 6-hour run fits comfortably; when a session does end, one click on Run all brings everything back.
The notebook runs on Google's hardware, so Colab's terms apply while you're there — but your model keys stay yours : paste any provider key after boot, exactly like self-hosting, and nothing is ever stored on our side. When the Colab session ends, everything on the VM is gone.
Surface
What it does
Terminal
Streaming chat with 300+ models — normal, thinking, research and coding modes — with a live ECG-style health trace in the toolbar that flatlines red the moment the catalog is unreachable
Model marketplace
One unified catalog across all 9 providers — cards now carry the platform's own cover art and a research panel with real download counts, licences and benchmarks
Music player
Search any song and play it in the marketplace — keyless, no YouTube API key — with a 5-band equalizer that actually re-shapes the audio
Agent builder
Describe a task in plain English; ENZO drafts the agent's full operating manual, and the agent keeps training itself on your activity from then on
Research mode
A deep-research loop that writes its own queries, reads what it finds, and decides when it's done — under hard budgets so it can't burn your key
Code-gen
Writes a coding project, boots it, previews it live, and tells you when it's broken
Vault
Every key sealed in the browser, attached per-request, wipeable in one click
How the agent builder works
Pass 1 — analysis. A two-pass drafter reads the domain of your task ("an agent that researches MUN country positions") and derives what the manual needs to cover: tacit knowledge, decision heuristics, edge cases.
The race. When a draft needs a model, ~10 free candidates from your own providers fire simultaneously — the first to answer wins, stragglers are aborted, and a brain-health scoreboard reorders future races by which models actually deliver. Dead or rate-limited free-tier models can no longer collapse a draft.
Honest provenance. Every agent records which model actually drafted it — and says so plainly when nothing was reachable.
It doesn't stop. A per-agent neural layer folds in domain-matched platform activity on a 90-second cadence, distills lessons into memory, and injects a live NEURAL FOCUS block into every run. Watch it in the agent's Neural tab.
Most "AI workspaces" hold your keys, meter your usage, or need a subscription to exist. ENZO is built the other way around:
(Competitor column is about the category, not specific products — details vary.)
The full threat model is written down — checkable, with the code that makes each claim true — in docs/SECURITY.md . The short version:
Keys are sealed in your browser with AES-256-GCM under a non-extractable WebCrypto key. It can be used to decrypt your keys while never being copied — no JavaScript can export its bytes, ours or an attacker's. Optional passphrase mode re-seals everything under PBKDF2-SHA256 (600,000 iterations) and deletes the device key entirely.
One module touches key storage, and CI enforces it. A pipeline stage greps the frontend for any raw localStorage key read and fails the build on a hit — a missed key-access site is a red build, never a production bug.
Every push runs a 44-assertion black-box pentest against a booted server — auth bypass, hostile payloads, IDOR, stream integrity — plus a keyless-boot proof: the server must start with zero provider keys. That's the BYOK guarantee, tested, not promised.
The limits are stated up front. Self-hosted mode stores the first key you claim in the container .env (sealed in the memory volume) so scheduled agents can run while your browser is closed — that trade is documented, not hidden. docs/SECURITY.md covers what's protected, what isn't, and why.
In-chat file converter — attach a PDF, spreadsheet, CSV, JSON, TXT or Markdown file in the terminal chat and it's parsed to real text/rows in your browser (files never leave the device; only the reasoning step uses your own key, like normal chat). The agent extracts, merges, or cross-converts — and CSV / Excel download buttons appear right on the reply . Built for research papers: "extract every table and merge into one CSV" now works end to end, and scanned PDFs report their missing text layer instead of failing.
Run it on Google Colab — one click on the Open In Colab button (see Run on Google Colab ): the notebook clones, installs, builds and boots ENZO on Google's hardware, hands you a URL for your browser plus a tunnel URL for your phone, and arms a keep-alive + watchdog so it stays up for 6+ hours.
Self-healing Docker pulls — the container now verifies its dependencies on every start and installs anything missing before the server boots ( ENZO_AUTO_INSTALL=0 to skip). A pulled image can't boot broken.
What's new in v1.3.0 (previous)
Music player — search any song and play it straight from the marketplace, keyless (no YouTube API key, no quota): a collapsed corner pill expands into a full player card — vinyl disc hero, queue walking, shuffle/loop/like, keyboard controls. For You turns your own listening history (kept device-local) into song seeds through your own provider key.
Real equalizer — a 5-band Web Audio EQ (bass / low-mid / mid / presence / air, ±12 dB, preamp, five presets) that genuinely re-shapes the frequency response when you opt in. Enhance starts off — normal playback is untouched. Tracks the enhancer can't stream fall back to the YouTube engine automatically; playback never breaks.
Marketplace, redesigned — every model card now carries the platform's own cover art, brand colour on hover, live health dot + latency, and a research panel with real facts: HuggingFace download counts, licences, knowledge cutoffs, Artificial Analysis scores and a Wikipedia-backed family summary — pulled keyless from public endpoints, never guessed.
NYC Subway theme — the workspace's new flagship backdrop: an AI-animated subway ride through a tunnel, with a handheld-camera tremble, monochrome film grade and animated recording grain added in code (the video ships clean).
A quieter interface — the whole workspace went monochrome + a single coral accent: the terminal toggle switch reb

[truncated]
