---
source: "https://github.com/master5d/viberuler"
hn_url: "https://news.ycombinator.com/item?id=50009232"
title: "Show HN: Viberuler – measure where your AI coding context goes"
article_title: "GitHub - master5d/viberuler: The benchmark for vibe coders. npx viberuler — tokens per dollar, LoC shipped, meme ranks, global leaderboard. · GitHub"
image: "https://opengraph.githubassets.com/d63883912a8585ebeb9c14ab11b95f2fb564c8cb2e5c61f8c9e33ca57eef53ed/master5d/viberuler"
author: "daocat"
captured_at: "2026-10-08T18:30:43Z"
capture_tool: "hn-digest"
hn_id: 50009232
score: 3
comments: 0
posted_at: "2026-10-08T17:48:50Z"
tags:
  - hacker-news
---

# Show HN: Viberuler – measure where your AI coding context goes

- HN: [50009232](https://news.ycombinator.com/item?id=50009232)
- Source: [github.com](https://github.com/master5d/viberuler)
- Score: 3
- Comments: 0
- Posted: 2026-10-08T17:48:50Z

## Translation

Title: Show HN: Viberuler – measure where your AI coding context goes
Article title: GitHub - master5d/viberuler: The benchmark for vibe coders. npx viberuler — tokens per dollar, LoC shipped, meme ranks, global leaderboard. · GitHub
Description: The benchmark for vibe coders. npx viberuler — tokens per dollar, LoC shipped, meme ranks, global leaderboard. - master5d/viberuler

Article text:
GitHub - master5d/viberuler: The benchmark for vibe coders. npx viberuler — tokens per dollar, LoC shipped, meme ranks, global leaderboard. · GitHub
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
master5d
viberuler Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
The benchmark for vibe coders. npx viberuler — tokens per dollar, LoC shipped, meme ranks, global leaderboard.
Readme MIT license Contributing
3 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
217 Commits 217 Commits Folders and files
.github/ workflows .github/ workflows assets assets dashboard dashboard design design docs docs integrations/ statusline integrations/ statusline logs logs packages packages scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md LAUNCH.md LAUNCH.md LICENSE LICENSE METHODOLOGY.md METHODOLOGY.md PRIVACY.md PRIVACY.md README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
The benchmark for vibe coders.
How hard do you actually vibe? There's only one way to find out:
viberuler scans your machine — locally, in seconds — and computes your VIBE SCORE :
Then it prints a scorecard you'll screenshot before you can stop yourself.
tokens per dollar is the headline stat. Anyone can burn tokens. Burning them efficiently is the game.
Tokens are casino chips. They're denominated the way they are so spending them
doesn't feel like spending money — and every new way of working (vibe it, ramble it,
let the agent figure it out) lowers the effort and raises the burn. viberuler is the
cashier's window: it converts the chips back into dollars — API-equivalent, total and
per platform — before you sit back down at the table.
LoC counts what you wrote, not what you have. It used to be the size of your repo trees — which credited you with vendored code you never touched and with every line a compiler emitted. Now it's the lines you added in your own commits, generated output excluded. That cut the author's own headline number by 17%, and we shipped the smaller true one. Why, in detail →
The excluded lines are still shown. A third of the author's diff turned out to be machine output — regenerated types, bundles, lockfiles. It isn't scored and never leaves your machine; it's a baseline to shrink. A number you can't see is a number you can't reduce.
It's a mirror, not a performance review. Every number here is scoped to one machine and one person: yours. Score a junior on tokens burned and you'll get a junior who burns tokens — Goodhart shows up inside a week, and you paid for the tokens. Screenshot it, flex it, compare it with people who opted in; the moment it lands in someone's perf review it stops measuring anything at all.
Prompt Peasant → Vibe Apprentice → Token Burner → Context Goblin → Ship Machine → GIGACHAD SHIPPER → Singularity Adjacent
(No data? You get NPC (no vibes detected) . We're sorry. We're not sorry.)
💰 Token Billionaire — ≥1B tokens
🪦 Free Tier Martyr — ≥1M tokens under $1
🗄️ Cache Whisperer — >90% cache reads
🌐 Polyglot — 5+ languages
🐘 Monorepo Menace — 100K+ lines authored in one repo
🔥 Streak Freak — 100-day streak
🌙 3AM Committer — 10+ night commits
💥 YOLO Force Pusher — 20+ history rewrites
🎰 High Roller — $1,000+ API-equivalent burned
🎲 Table Hopper — $1+ of burn on 3+ platforms
The leaderboard
npx viberuler --submit
GitHub device-flow login → your score goes live at viberuler.dev/u/<you> as a Certificate of Vibe Measurement (LoC · tok/$ · streak · agents · rank · title), built for flexing. Global rank. Efficiency percentile. Prefilled share links — X · LinkedIn · Facebook · Bluesky.
Share to Stories — the certificate page also renders a vertical 9:16 story card (Spotify-Wrapped-style stat reveal) and a Share to Stories button. On mobile it hands the card straight to your phone's native share sheet — Instagram, WhatsApp, Facebook, Messenger; on desktop it downloads the card to post. (Stories are app-only, so this is the only way in — same mechanism Wrapped uses.)
The default run makes zero network calls . Zero.
--submit sends aggregates only — fourteen fields: aggregate stats, achievement ids, your coding-agent names, commit streak, and ship outcomes (features shipped / PRs merged). No paths, no repo names, no prompts, no code. Ever.
Before anything is sent, the CLI prints the exact JSON payload and asks.
--share also sends nothing: it prints a URL carrying your display numbers. Opening or posting that link is your choice, and the card it renders is marked unverified.
Time metrics come from timestamps already inside your transcripts. Nothing watches your screen, and no window titles, paths, or filenames are read, stored, or sent.
Don't trust us — read the ~140 lines: packages/cli/src/payload.ts and packages/cli/src/submit.ts . Details: PRIVACY.md .
Full formula, price table, normalization and honest disclaimers in METHODOLOGY.md . Short version:
VIBE = 1000·log₁₀(1 + LoC/1000) # shipping volume
+ 500·log₁₀(1 + tokens/1M) # AI leverage
+ 800·efficiency_percentile # tokens/$ vs the world
+ 300·log₁₀(1 + projects·10) # breadth
+ min(streak, 365) + 50·achievements
Logarithms everywhere — whales get compressed, newcomers have room to climb.
npx viberuler wrapped --month 2026-06
Your Vibe Wrapped for the month — commits, busiest day, streak, top language, and Claude Code tokens/cost for that window. 100% local; screenshot and flex.
npx viberuler # scan + scorecard (100% local)
npx viberuler audit # audit your rig — see below
npx viberuler --submit # push to the global leaderboard
npx viberuler payload # show exactly what --submit WOULD send
npx viberuler --json # machine-readable
npx viberuler --scan-dir ~/code --since 2026-01-01
npx viberuler --scan-dir ~/work --scan-dir ~/oss # repeatable — scans ALL repos under each root
npx viberuler --github <handle> # add your stars (the only other network call)
npx viberuler --share # print a shareable card URL (nothing is sent)
npx viberuler audit --idle-gap 5 # pause (minutes) that stops counting as attention
npx viberuler audit --compare 2026-06-01..2026-07-01 # that window vs the equally long one after it
npx viberuler audit --market # reprice your mix at live market rates (one catalog GET, opt-in)
A bare run scans every git repo under your home dir . If your code lives elsewhere (or in several places), point --scan-dir at each root — it's repeatable, and every metric (LoC, commits, features, PRs) is summed across all repos found, so your certificate reflects your whole body of work, not one project.
npx viberuler audit
Your tokens per dollar score says how efficiently you burn tokens. audit says how efficiently your rig is set up. 100% local, reads your Claude Code transcripts:
Token economy — cache-hit rate, and what prompt caching actually saved you in API-equivalent dollars.
Context amplification — how many times the average token you admit gets re-read on later turns. Measured on the main thread alone : pooling in short-lived subagent contexts halves the number and understates what a token really costs in the thread you live in. (On a real rig: 1088× .)
Subagents — how hard they compress: work pulled in inside a subagent vs the summary handed back. They aren't free (they cost real overhead), but at 1000× amplification, keeping tokens out of the main thread is the whole game.
🧊 Cold context — what a session costs before you type a word : system prompt, tool names, agent and skill descriptions, memory. Every subagent spawn re-pays it. (On a real rig: 50.2K tokens at session start, 33.1K re-paid on each of 3,234 spawns.)
👻 Ghost tokens — the tokens an output-rewriting plugin promises to save you, measured instead of promised. Oversized results (>4KB) were 54% of everything admitted; the famous "dedupe repeat reads" trick was worth 2% .
Top tools — which tools are actually filling your context, ranked by calls and by tokens.
☠️ Dead weight — MCP servers and plugins that load on every session, spawn a process, inject their tool schemas… and get called zero times .
That last one is the point. On the author's rig it found two MCP servers burning 1.5 GB of RAM across 76 processes for 0 calls in 10,700+ sessions . Removing them dropped the median cold context from 49.9K to 41.1K tokens — a 17.5% discount on every session since, and proof that deferred tool schemas do not make an unused server free.
Measuring beats guessing, which is an awkward conclusion for a tool that scores you on vibes.
Context waste — sizes and levers, never a savings claim
A crowded genre of tools promises to cut your tokens "by up to 90%". They're real
projects; none of them publishes a methodology behind the number. audit takes the
other side: it names where context went, on your own transcripts.
CONTEXT WASTE
5.7M tok · 2,555 calls oversized single results → slice / grep before read
2.2M tok · 4,068 calls subagent-returned tokens → tighter subagent contracts
1.9M tok · 3,394 calls whole-file reads never edited → outline-first / symbol reads
Three rules hold this honest, and they are enforced in the output, not just in docs:
No "you would save N%". The counterfactual is unknowable — a read that changed
nothing may be the read that told you not to change something.
No total. The classes overlap by construction (an oversized read can also be
exploratory), so summing them would double-count and manufacture a headline.
audit --compare A..B puts two windows side by side — install an optimizer,
compare before and after — and labels the delta an observation, not causation ,
because workload differs between windows. A window with no sessions says "not enough
data" instead of printing a flattering zero.
Your mix at market rates ( audit --market )
The chips-to-dollars view, extended to the whole exchange: your measured token mix,
repriced at the list rates of a curated set of current models — open-weight flagships
(Kimi K3, DeepSeek, GLM, Qwen, Llama, MiniMax) and the closed tiers. Rates come from
one anonymous GET to the OpenRouter public catalog (this flag is opt-in network,
like --github ), cached locally for a day, with an offline fallback to a bundled
snapshot — the output labels which source answered. Arithmetic, not advice: tokenizers,
quality, and cache mechanics differ, so the table says what your volume costs per
counter, never what you would save.
The same table answers maximum AI performance per dollar : where the catalog carries
an Artificial Analysis intelligence index, each counter shows what one index point
costs on your mix ( $/pt ), the ranking is by that, and the cheapest wears 🏆 —
points-per-dollar was the first cut, but it reads 0.00 o

[truncated]

## Original Extract

The benchmark for vibe coders. npx viberuler — tokens per dollar, LoC shipped, meme ranks, global leaderboard. - master5d/viberuler

GitHub - master5d/viberuler: The benchmark for vibe coders. npx viberuler — tokens per dollar, LoC shipped, meme ranks, global leaderboard. · GitHub
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
master5d
viberuler Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
The benchmark for vibe coders. npx viberuler — tokens per dollar, LoC shipped, meme ranks, global leaderboard.
Readme MIT license Contributing
3 forks Report repository master Branches Tags Go to file Code Open more actions menu Latest commit
217 Commits 217 Commits Folders and files
.github/ workflows .github/ workflows assets assets dashboard dashboard design design docs docs integrations/ statusline integrations/ statusline logs logs packages packages scripts scripts .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md DESIGN.md DESIGN.md LAUNCH.md LAUNCH.md LICENSE LICENSE METHODOLOGY.md METHODOLOGY.md PRIVACY.md PRIVACY.md README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
The benchmark for vibe coders.
How hard do you actually vibe? There's only one way to find out:
viberuler scans your machine — locally, in seconds — and computes your VIBE SCORE :
Then it prints a scorecard you'll screenshot before you can stop yourself.
tokens per dollar is the headline stat. Anyone can burn tokens. Burning them efficiently is the game.
Tokens are casino chips. They're denominated the way they are so spending them
doesn't feel like spending money — and every new way of working (vibe it, ramble it,
let the agent figure it out) lowers the effort and raises the burn. viberuler is the
cashier's window: it converts the chips back into dollars — API-equivalent, total and
per platform — before you sit back down at the table.
LoC counts what you wrote, not what you have. It used to be the size of your repo trees — which credited you with vendored code you never touched and with every line a compiler emitted. Now it's the lines you added in your own commits, generated output excluded. That cut the author's own headline number by 17%, and we shipped the smaller true one. Why, in detail →
The excluded lines are still shown. A third of the author's diff turned out to be machine output — regenerated types, bundles, lockfiles. It isn't scored and never leaves your machine; it's a baseline to shrink. A number you can't see is a number you can't reduce.
It's a mirror, not a performance review. Every number here is scoped to one machine and one person: yours. Score a junior on tokens burned and you'll get a junior who burns tokens — Goodhart shows up inside a week, and you paid for the tokens. Screenshot it, flex it, compare it with people who opted in; the moment it lands in someone's perf review it stops measuring anything at all.
Prompt Peasant → Vibe Apprentice → Token Burner → Context Goblin → Ship Machine → GIGACHAD SHIPPER → Singularity Adjacent
(No data? You get NPC (no vibes detected) . We're sorry. We're not sorry.)
💰 Token Billionaire — ≥1B tokens
🪦 Free Tier Martyr — ≥1M tokens under $1
🗄️ Cache Whisperer — >90% cache reads
🌐 Polyglot — 5+ languages
🐘 Monorepo Menace — 100K+ lines authored in one repo
🔥 Streak Freak — 100-day streak
🌙 3AM Committer — 10+ night commits
💥 YOLO Force Pusher — 20+ history rewrites
🎰 High Roller — $1,000+ API-equivalent burned
🎲 Table Hopper — $1+ of burn on 3+ platforms
The leaderboard
npx viberuler --submit
GitHub device-flow login → your score goes live at viberuler.dev/u/<you> as a Certificate of Vibe Measurement (LoC · tok/$ · streak · agents · rank · title), built for flexing. Global rank. Efficiency percentile. Prefilled share links — X · LinkedIn · Facebook · Bluesky.
Share to Stories — the certificate page also renders a vertical 9:16 story card (Spotify-Wrapped-style stat reveal) and a Share to Stories button. On mobile it hands the card straight to your phone's native share sheet — Instagram, WhatsApp, Facebook, Messenger; on desktop it downloads the card to post. (Stories are app-only, so this is the only way in — same mechanism Wrapped uses.)
The default run makes zero network calls . Zero.
--submit sends aggregates only — fourteen fields: aggregate stats, achievement ids, your coding-agent names, commit streak, and ship outcomes (features shipped / PRs merged). No paths, no repo names, no prompts, no code. Ever.
Before anything is sent, the CLI prints the exact JSON payload and asks.
--share also sends nothing: it prints a URL carrying your display numbers. Opening or posting that link is your choice, and the card it renders is marked unverified.
Time metrics come from timestamps already inside your transcripts. Nothing watches your screen, and no window titles, paths, or filenames are read, stored, or sent.
Don't trust us — read the ~140 lines: packages/cli/src/payload.ts and packages/cli/src/submit.ts . Details: PRIVACY.md .
Full formula, price table, normalization and honest disclaimers in METHODOLOGY.md . Short version:
VIBE = 1000·log₁₀(1 + LoC/1000) # shipping volume
+ 500·log₁₀(1 + tokens/1M) # AI leverage
+ 800·efficiency_percentile # tokens/$ vs the world
+ 300·log₁₀(1 + projects·10) # breadth
+ min(streak, 365) + 50·achievements
Logarithms everywhere — whales get compressed, newcomers have room to climb.
npx viberuler wrapped --month 2026-06
Your Vibe Wrapped for the month — commits, busiest day, streak, top language, and Claude Code tokens/cost for that window. 100% local; screenshot and flex.
npx viberuler # scan + scorecard (100% local)
npx viberuler audit # audit your rig — see below
npx viberuler --submit # push to the global leaderboard
npx viberuler payload # show exactly what --submit WOULD send
npx viberuler --json # machine-readable
npx viberuler --scan-dir ~/code --since 2026-01-01
npx viberuler --scan-dir ~/work --scan-dir ~/oss # repeatable — scans ALL repos under each root
npx viberuler --github <handle> # add your stars (the only other network call)
npx viberuler --share # print a shareable card URL (nothing is sent)
npx viberuler audit --idle-gap 5 # pause (minutes) that stops counting as attention
npx viberuler audit --compare 2026-06-01..2026-07-01 # that window vs the equally long one after it
npx viberuler audit --market # reprice your mix at live market rates (one catalog GET, opt-in)
A bare run scans every git repo under your home dir . If your code lives elsewhere (or in several places), point --scan-dir at each root — it's repeatable, and every metric (LoC, commits, features, PRs) is summed across all repos found, so your certificate reflects your whole body of work, not one project.
npx viberuler audit
Your tokens per dollar score says how efficiently you burn tokens. audit says how efficiently your rig is set up. 100% local, reads your Claude Code transcripts:
Token economy — cache-hit rate, and what prompt caching actually saved you in API-equivalent dollars.
Context amplification — how many times the average token you admit gets re-read on later turns. Measured on the main thread alone : pooling in short-lived subagent contexts halves the number and understates what a token really costs in the thread you live in. (On a real rig: 1088× .)
Subagents — how hard they compress: work pulled in inside a subagent vs the summary handed back. They aren't free (they cost real overhead), but at 1000× amplification, keeping tokens out of the main thread is the whole game.
🧊 Cold context — what a session costs before you type a word : system prompt, tool names, agent and skill descriptions, memory. Every subagent spawn re-pays it. (On a real rig: 50.2K tokens at session start, 33.1K re-paid on each of 3,234 spawns.)
👻 Ghost tokens — the tokens an output-rewriting plugin promises to save you, measured instead of promised. Oversized results (>4KB) were 54% of everything admitted; the famous "dedupe repeat reads" trick was worth 2% .
Top tools — which tools are actually filling your context, ranked by calls and by tokens.
☠️ Dead weight — MCP servers and plugins that load on every session, spawn a process, inject their tool schemas… and get called zero times .
That last one is the point. On the author's rig it found two MCP servers burning 1.5 GB of RAM across 76 processes for 0 calls in 10,700+ sessions . Removing them dropped the median cold context from 49.9K to 41.1K tokens — a 17.5% discount on every session since, and proof that deferred tool schemas do not make an unused server free.
Measuring beats guessing, which is an awkward conclusion for a tool that scores you on vibes.
Context waste — sizes and levers, never a savings claim
A crowded genre of tools promises to cut your tokens "by up to 90%". They're real
projects; none of them publishes a methodology behind the number. audit takes the
other side: it names where context went, on your own transcripts.
CONTEXT WASTE
5.7M tok · 2,555 calls oversized single results → slice / grep before read
2.2M tok · 4,068 calls subagent-returned tokens → tighter subagent contracts
1.9M tok · 3,394 calls whole-file reads never edited → outline-first / symbol reads
Three rules hold this honest, and they are enforced in the output, not just in docs:
No "you would save N%". The counterfactual is unknowable — a read that changed
nothing may be the read that told you not to change something.
No total. The classes overlap by construction (an oversized read can also be
exploratory), so summing them would double-count and manufacture a headline.
audit --compare A..B puts two windows side by side — install an optimizer,
compare before and after — and labels the delta an observation, not causation ,
because workload differs between windows. A window with no sessions says "not enough
data" instead of printing a flattering zero.
Your mix at market rates ( audit --market )
The chips-to-dollars view, extended to the whole exchange: your measured token mix,
repriced at the list rates of a curated set of current models — open-weight flagships
(Kimi K3, DeepSeek, GLM, Qwen, Llama, MiniMax) and the closed tiers. Rates come from
one anonymous GET to the OpenRouter public catalog (this flag is opt-in network,
like --github ), cached locally for a day, with an offline fallback to a bundled
snapshot — the output labels which source answered. Arithmetic, not advice: tokenizers,
quality, and cache mechanics differ, so the table says what your volume costs per
counter, never what you would save.
The same table answers maximum AI performance per dollar : where the catalog carries
an Artificial Analysis intelligence index, each counter shows what one index point
costs on your mix ( $/pt ), the ranking is by that, and the cheapest wears 🏆 —
points-per-dollar was the first cut, but it reads 0.00 o

[truncated]
