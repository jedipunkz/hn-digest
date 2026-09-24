---
source: "https://freshjots.com/"
hn_url: "https://news.ycombinator.com/item?id=49834382"
title: "Logs you, your AI and cron jobs write to that you can prove weren't altered"
article_title: "Fresh Jots — tamper-evident notes your scripts, agents, and you share"
image: "https://freshjots.com/og/home.png?v=2"
author: "arsphy"
captured_at: "2026-09-24T18:15:19Z"
capture_tool: "hn-digest"
hn_id: 49834382
score: 2
comments: 1
posted_at: "2026-09-24T17:53:45Z"
tags:
  - hacker-news
---

# Logs you, your AI and cron jobs write to that you can prove weren't altered

- HN: [49834382](https://news.ycombinator.com/item?id=49834382)
- Source: [freshjots.com](https://freshjots.com/)
- Score: 2
- Comments: 1
- Posted: 2026-09-24T17:53:45Z

## Translation

Title: Logs you, your AI and cron jobs write to that you can prove weren't altered
Article title: Fresh Jots — tamper-evident notes your scripts, agents, and you share
Description: Give your scripts and AI their own tamper-evident notes. Prove nothing changed, encrypt with your own key, write over MCP or one curl. Free to start.

Article text:
Skip to content
nav-menu#clickOutside">
Fresh Jots
nav-menu#toggle"
data-nav-menu-target="button"
aria-controls="mobile-nav"
aria-expanded="false"
aria-label="Open menu">
Blog
Developers
Sign in
Sign up
Blog
Developers
Sign in
Sign up
Tamper-evident · Private · AI-ready
Tamper-evident notes for you, your scripts and AI — and you can prove it.
Write manually. Write from the CLI, your project code, or an MCP client; set a dead man's switch: if the writes stop, you get an email.
Tamper-evident — every entry hashed & timestamped
Append-only option — add lines, never rewrite them
Encrypted option — you hold the only key, so not even we can read it
Export everything — leave anytime
No credit card · Free read-only MCP for your AI · No ads, no AI training on your content
Demo: running the command freshjots append deploys "v2 ok" in a terminal instantly adds that line to the top of the "deploys" note in Fresh Jots.
No credit card, no trial clock. Sign up and you're in a working notebook in under a minute.
2
Receive a token — or connect your AI
A 14-day API trial for the full Dev surface. Or add https://freshjots.com/mcp in Claude, Cursor, or any MCP client and it reads your notes over OAuth — free and read-only, with agent writes on paid plans.
Send a note to your Fresh Jots account in three lines:
curl -fsSL https://freshjots.com/install.sh | sh
export FRESHJOTS_TOKEN=mn_your_token
freshjots append cron-jobs-prod "backup ok"
Installs the CLI, sets your token, and appends a line — the note is created on the first write.
Set up a token so Claude, Codex, or any agent appends to a notebook you can read on your phone.
Rich and plain notes, the REST API, dead man's alerts, quick capture — everything that ships.
One note store — read it in the browser, on your phone, or straight from code.
Every script gets its own notebook — append-only notes plus an atomic append API. Your scripts write, you read: a tamper-resistant log you own. Transform any note into a webhook inbox, or into a dead man's switch that emails you when the writes stop.
One-click MCP connect for Claude, Cursor, and other tools. Free to read and search; upgrade to let your agents write — session notes, research dumps, transcripts you can read on your phone.
Write everything in a single place, reachable by all of your devices. REST API, webhooks, an ingest inbox, and official Bash, npm, Ruby, and Python clients — capture in one curl or one tap ( see integrations ).
The home page to the left, the open note on the right. That's the basic UI.
Plain notes from cron, CI, webhooks. Rich notes for everything else. Same notebook, same search.
An official client in your language
Homebrew brew install Goran-Arsov/freshjots/freshjots
RubyGems gem install freshjots
GitHub Actions Goran-Arsov/freshjots-append@v1
What's inside — and your data stays yours
Every note is hashed and anchored to the Bitcoin blockchain, so you can prove an AI session or log existed by a certain time and hasn't changed since — evidence anyone can check without trusting us. Sessions captured with the Fresh Jots hook get an immediate, independent timestamp on top: a trusted authority stamps each one the moment you save it, proving the exact time, not just the day. How it works .
Client-encrypt a note with your own key and store the ciphertext — we keep bytes we can't read, search, or decrypt. How it works .
Rich notes for writing, plain notes for streaming output. Search across both.
Every edit persists as you type. Restore any earlier snapshot — the paragraph you cut last week is still there.
.txt, .docx, .pdf per note. Full archive as a single .zip from Options — even after you cancel.
One tap on iOS or Android. No app store, no updates, no native bloat. Set it up .
Not on the free plan, not anywhere.
We don't read your notes, and our Terms forbid training models on them or selling them to anyone who would.
One account, all your devices. Not per device. Not per feature.
Up to 10 notes. No card, no trial clock.
≈ $2.00 / month, billed annually
Up to 1,000 notes. Your AI reads and writes over MCP.
Personal + REST API, 10K notes, dead man's alerts, webhooks, and more.
2-seat minimum ($358/yr floor)
Up to 25 seats, shared notes + audit log.
Yes. Every account can keep up to 10 notes free — no card, no trial clock. When you need more, subscribe; see the pricing page .
Do you read my notes, or train AI on them?
No. We don't browse your notes, and our Terms commit us not to train machine-learning models on them or sell them to anyone who would. More on privacy →
Always. Download any note as .txt, .docx, or .pdf, or your full archive as a single .zip from Options — even after you cancel. How exports work →
What do my scripts and AI need in order to write here?
Scripts use an API token — included with the Dev plan, or free for 14 days with the API trial. Your AI reads over MCP on every plan, free; agent writes come with the Personal plan ($24/yr).
Coming from another notes app?
One prompt: paste this into Codex CLI and your sessions auto-archive to Fresh Jots
Connect Fresh Jots to any AI assistant over MCP — just talk to your notes
`/save-to-freshjots <title> :: <body>` — one running engineering journal, one titled append
Start writing. Upgrade only if you hit the ceiling.
No ads · No AI training · Export anytime
No credit card. 10 notes free, forever.

## Original Extract

Give your scripts and AI their own tamper-evident notes. Prove nothing changed, encrypt with your own key, write over MCP or one curl. Free to start.

Skip to content
nav-menu#clickOutside">
Fresh Jots
nav-menu#toggle"
data-nav-menu-target="button"
aria-controls="mobile-nav"
aria-expanded="false"
aria-label="Open menu">
Blog
Developers
Sign in
Sign up
Blog
Developers
Sign in
Sign up
Tamper-evident · Private · AI-ready
Tamper-evident notes for you, your scripts and AI — and you can prove it.
Write manually. Write from the CLI, your project code, or an MCP client; set a dead man's switch: if the writes stop, you get an email.
Tamper-evident — every entry hashed & timestamped
Append-only option — add lines, never rewrite them
Encrypted option — you hold the only key, so not even we can read it
Export everything — leave anytime
No credit card · Free read-only MCP for your AI · No ads, no AI training on your content
Demo: running the command freshjots append deploys "v2 ok" in a terminal instantly adds that line to the top of the "deploys" note in Fresh Jots.
No credit card, no trial clock. Sign up and you're in a working notebook in under a minute.
2
Receive a token — or connect your AI
A 14-day API trial for the full Dev surface. Or add https://freshjots.com/mcp in Claude, Cursor, or any MCP client and it reads your notes over OAuth — free and read-only, with agent writes on paid plans.
Send a note to your Fresh Jots account in three lines:
curl -fsSL https://freshjots.com/install.sh | sh
export FRESHJOTS_TOKEN=mn_your_token
freshjots append cron-jobs-prod "backup ok"
Installs the CLI, sets your token, and appends a line — the note is created on the first write.
Set up a token so Claude, Codex, or any agent appends to a notebook you can read on your phone.
Rich and plain notes, the REST API, dead man's alerts, quick capture — everything that ships.
One note store — read it in the browser, on your phone, or straight from code.
Every script gets its own notebook — append-only notes plus an atomic append API. Your scripts write, you read: a tamper-resistant log you own. Transform any note into a webhook inbox, or into a dead man's switch that emails you when the writes stop.
One-click MCP connect for Claude, Cursor, and other tools. Free to read and search; upgrade to let your agents write — session notes, research dumps, transcripts you can read on your phone.
Write everything in a single place, reachable by all of your devices. REST API, webhooks, an ingest inbox, and official Bash, npm, Ruby, and Python clients — capture in one curl or one tap ( see integrations ).
The home page to the left, the open note on the right. That's the basic UI.
Plain notes from cron, CI, webhooks. Rich notes for everything else. Same notebook, same search.
An official client in your language
Homebrew brew install Goran-Arsov/freshjots/freshjots
RubyGems gem install freshjots
GitHub Actions Goran-Arsov/freshjots-append@v1
What's inside — and your data stays yours
Every note is hashed and anchored to the Bitcoin blockchain, so you can prove an AI session or log existed by a certain time and hasn't changed since — evidence anyone can check without trusting us. Sessions captured with the Fresh Jots hook get an immediate, independent timestamp on top: a trusted authority stamps each one the moment you save it, proving the exact time, not just the day. How it works .
Client-encrypt a note with your own key and store the ciphertext — we keep bytes we can't read, search, or decrypt. How it works .
Rich notes for writing, plain notes for streaming output. Search across both.
Every edit persists as you type. Restore any earlier snapshot — the paragraph you cut last week is still there.
.txt, .docx, .pdf per note. Full archive as a single .zip from Options — even after you cancel.
One tap on iOS or Android. No app store, no updates, no native bloat. Set it up .
Not on the free plan, not anywhere.
We don't read your notes, and our Terms forbid training models on them or selling them to anyone who would.
One account, all your devices. Not per device. Not per feature.
Up to 10 notes. No card, no trial clock.
≈ $2.00 / month, billed annually
Up to 1,000 notes. Your AI reads and writes over MCP.
Personal + REST API, 10K notes, dead man's alerts, webhooks, and more.
2-seat minimum ($358/yr floor)
Up to 25 seats, shared notes + audit log.
Yes. Every account can keep up to 10 notes free — no card, no trial clock. When you need more, subscribe; see the pricing page .
Do you read my notes, or train AI on them?
No. We don't browse your notes, and our Terms commit us not to train machine-learning models on them or sell them to anyone who would. More on privacy →
Always. Download any note as .txt, .docx, or .pdf, or your full archive as a single .zip from Options — even after you cancel. How exports work →
What do my scripts and AI need in order to write here?
Scripts use an API token — included with the Dev plan, or free for 14 days with the API trial. Your AI reads over MCP on every plan, free; agent writes come with the Personal plan ($24/yr).
Coming from another notes app?
One prompt: paste this into Codex CLI and your sessions auto-archive to Fresh Jots
Connect Fresh Jots to any AI assistant over MCP — just talk to your notes
`/save-to-freshjots <title> :: <body>` — one running engineering journal, one titled append
Start writing. Upgrade only if you hit the ceiling.
No ads · No AI training · Export anytime
No credit card. 10 notes free, forever.
