---
source: "https://github.com/cookiy-ai/sell-sessions-skill"
hn_url: "https://news.ycombinator.com/item?id=49720176"
title: "Show HN: Open-source tool to sell your Claude Code and Codex sessions"
article_title: "GitHub - cookiy-ai/sell-sessions-skill · GitHub"
image: "https://opengraph.githubassets.com/c8cc3bbe0ca7f0b6bd7db723e2dd3cb0a7d1dacb21892d9235b161e8c2b952c9/cookiy-ai/sell-sessions-skill"
author: "YouchaoDong"
captured_at: "2026-09-15T23:59:57Z"
capture_tool: "hn-digest"
hn_id: 49720176
score: 1
comments: 1
posted_at: "2026-09-15T23:18:50Z"
tags:
  - hacker-news
---

# Show HN: Open-source tool to sell your Claude Code and Codex sessions

- HN: [49720176](https://news.ycombinator.com/item?id=49720176)
- Source: [github.com](https://github.com/cookiy-ai/sell-sessions-skill)
- Score: 1
- Comments: 1
- Posted: 2026-09-15T23:18:50Z

## Translation

Title: Show HN: Open-source tool to sell your Claude Code and Codex sessions
Article title: GitHub - cookiy-ai/sell-sessions-skill · GitHub
Description: Contribute to cookiy-ai/sell-sessions-skill development by creating an account on GitHub.

Article text:
GitHub - cookiy-ai/sell-sessions-skill · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
cookiy-ai
/
sell-sessions-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
assets assets .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
Your agent history is worth money.
Sell the Claude Code and Codex sessions you already ran. Open-source skill + local uploader.
Runs locally · You choose every session · Secrets and PII scrubbed on your machine · Paid on every sale
Website ·
Install ·
How it works ·
Privacy ·
FAQ
One command. Your sessions. Your price.
npx @cookiyai/sell-agent-sessions
That opens a local web app on your machine that lists your Claude Code and Codex sessions, lets you pick the ones you want to license, scrubs secrets and PII before anything leaves your laptop, and shows you an estimated value per session. Installing or opening it uploads nothing .
Creators are talking about it 🎬
Launched publicly on September 10, 2026 . First creator video crossed 80K views in 4 days .
rishi-82k-views.mp4
@rishiexplainsai · Instagram · 82.4K views · 320 likes · 246 comments
rishi-5k-views.mp4
@rishiexplainsai · Instagram · 5K+ views · posted this week
rishi-new-sep12.mp4
@rishiexplainsai · new episode, going live this week
What builders say after trying it
creator-02.mp4
"Nothing goes without your approval. If you find something confidential, you just redact it."
creator-05.mp4
"Buyers aren't after random prompts. They want sessions from engineers who know what they're doing."
creator-03.mp4
"I'm the guy who reads every Terms page. Nothing is sent until you approve it yourself."
creator-08.mp4
"I run infrastructure for a living. No daemons, no phone-home. Only what you approve gets sent."
creator-06.mp4
"No new hours, no second job grind. The work is already finished."
creator-09.mp4
"I'm the paranoid friend who tapes his webcam, and even I signed off."
creator-04.mp4
"The valuable piece isn't the code. It's the corrections."
creator-10.mp4
"900 sessions, 40 projects. I had no idea it ran this deep."
Batch 1 of the creator program: 20 videos, September 2026. 200 more in production.
Why your sessions are worth money
Every session where you planned a task, corrected the model, ran tests and shipped is a worked example of how real engineering gets done. AI labs already train on this kind of data. Cookiy Earn lets you decide which sessions get licensed, and pays you for each sale.
Estimates are shown before upload and are not offers. Actual prices are set when a licensing sale completes and depend on quality, difficulty, completeness, model, token volume and buyer demand. A session can earn again each time it is licensed to a new buyer.
Choose your sessions. The local uploader scans ~/.claude , ~/.codex and their platform equivalents and lists everything by project. You tick what to offer.
Scrub and estimate, locally. Detected secrets and PII are redacted on your device. Each session gets an estimated value before upload.
Match and earn. Cookiy licenses your sessions to buyers and pays you on every completed sale, including repeat sales of the same session.
Local first. Session files never leave your machine until you select them and confirm.
Nothing on install. Installing or opening the uploader does not upload anything.
Redaction before upload. Secrets and personally identifiable information are detected and stripped on your device. You can review before confirming.
No background daemons, no phone-home. The uploader is a plain local web app you can stop any time.
Open source, MIT. Read the skill . Inspect the uploader package .
Honest limits. Automated scrubbing reduces risk a lot, but no scrubber catches everything. Review what you upload.
Option 1: run the uploader directly
npx @cookiyai/sell-agent-sessions
Requires Node.js and npm. Opens http://localhost:4318 in your browser.
Option 2: install the skill into your agent
npx skills add cookiy-ai/sell-sessions-skill --global
Then tell Claude Code, Codex, Cursor or any skill-aware agent:
I want to sell my agent sessions to earn income.
The agent explains the product, installs the uploader on your host machine and launches it for you.
Option 3: paste this into your agent
Follow https://github.com/cookiy-ai/sell-sessions-skill to install the skill and start the local session uploader for me.
Why would anyone pay for my sessions?
Real agent sessions show how models handle practical, multi-step work: planning, tool use, coding, debugging, corrections and final outcomes. Labs use them for training, evaluation and product improvement. Difficult, coherent, completed sessions are worth more than short or abandoned ones.
How much is a session worth?
The uploader shows an estimate per session before upload. Actual prices depend on quality, task difficulty, model, token volume, completeness, licensing terms and buyer demand. Estimates are not offers.
Does installing the uploader expose my data?
No. Installing or opening the uploader uploads nothing. Your sessions stay on your device until you select them and confirm.
What exactly gets uploaded?
Only the sessions you select and confirm, after local redaction. Selecting alone does not start an upload.
When do I get paid?
When a buyer completes a licensing purchase, not when you upload. A session may earn again if licensed to another buyer. Track submissions and earnings in your Cookiy Earn account.
Do I lose ownership?
You license selected sessions. Nothing is transferred automatically. Exclusivity, deletion, buyer rights, payout timing and taxes are covered by the current Cookiy Earn terms .
Why do I need to install something?
Your history lives on your computer. The uploader needs local access to list sessions, let you choose, scrub private data, estimate value and transfer only what you selected.
Star history
Built by the team behind the Cookiy User Research Skill (1.5K+ stars) and Cookiy.ai , the agentic user-research platform.
Readme MIT license Activity Custom properties Stars
12 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to cookiy-ai/sell-sessions-skill development by creating an account on GitHub.

GitHub - cookiy-ai/sell-sessions-skill · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
cookiy-ai
/
sell-sessions-skill
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
assets assets .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
Your agent history is worth money.
Sell the Claude Code and Codex sessions you already ran. Open-source skill + local uploader.
Runs locally · You choose every session · Secrets and PII scrubbed on your machine · Paid on every sale
Website ·
Install ·
How it works ·
Privacy ·
FAQ
One command. Your sessions. Your price.
npx @cookiyai/sell-agent-sessions
That opens a local web app on your machine that lists your Claude Code and Codex sessions, lets you pick the ones you want to license, scrubs secrets and PII before anything leaves your laptop, and shows you an estimated value per session. Installing or opening it uploads nothing .
Creators are talking about it 🎬
Launched publicly on September 10, 2026 . First creator video crossed 80K views in 4 days .
rishi-82k-views.mp4
@rishiexplainsai · Instagram · 82.4K views · 320 likes · 246 comments
rishi-5k-views.mp4
@rishiexplainsai · Instagram · 5K+ views · posted this week
rishi-new-sep12.mp4
@rishiexplainsai · new episode, going live this week
What builders say after trying it
creator-02.mp4
"Nothing goes without your approval. If you find something confidential, you just redact it."
creator-05.mp4
"Buyers aren't after random prompts. They want sessions from engineers who know what they're doing."
creator-03.mp4
"I'm the guy who reads every Terms page. Nothing is sent until you approve it yourself."
creator-08.mp4
"I run infrastructure for a living. No daemons, no phone-home. Only what you approve gets sent."
creator-06.mp4
"No new hours, no second job grind. The work is already finished."
creator-09.mp4
"I'm the paranoid friend who tapes his webcam, and even I signed off."
creator-04.mp4
"The valuable piece isn't the code. It's the corrections."
creator-10.mp4
"900 sessions, 40 projects. I had no idea it ran this deep."
Batch 1 of the creator program: 20 videos, September 2026. 200 more in production.
Why your sessions are worth money
Every session where you planned a task, corrected the model, ran tests and shipped is a worked example of how real engineering gets done. AI labs already train on this kind of data. Cookiy Earn lets you decide which sessions get licensed, and pays you for each sale.
Estimates are shown before upload and are not offers. Actual prices are set when a licensing sale completes and depend on quality, difficulty, completeness, model, token volume and buyer demand. A session can earn again each time it is licensed to a new buyer.
Choose your sessions. The local uploader scans ~/.claude , ~/.codex and their platform equivalents and lists everything by project. You tick what to offer.
Scrub and estimate, locally. Detected secrets and PII are redacted on your device. Each session gets an estimated value before upload.
Match and earn. Cookiy licenses your sessions to buyers and pays you on every completed sale, including repeat sales of the same session.
Local first. Session files never leave your machine until you select them and confirm.
Nothing on install. Installing or opening the uploader does not upload anything.
Redaction before upload. Secrets and personally identifiable information are detected and stripped on your device. You can review before confirming.
No background daemons, no phone-home. The uploader is a plain local web app you can stop any time.
Open source, MIT. Read the skill . Inspect the uploader package .
Honest limits. Automated scrubbing reduces risk a lot, but no scrubber catches everything. Review what you upload.
Option 1: run the uploader directly
npx @cookiyai/sell-agent-sessions
Requires Node.js and npm. Opens http://localhost:4318 in your browser.
Option 2: install the skill into your agent
npx skills add cookiy-ai/sell-sessions-skill --global
Then tell Claude Code, Codex, Cursor or any skill-aware agent:
I want to sell my agent sessions to earn income.
The agent explains the product, installs the uploader on your host machine and launches it for you.
Option 3: paste this into your agent
Follow https://github.com/cookiy-ai/sell-sessions-skill to install the skill and start the local session uploader for me.
Why would anyone pay for my sessions?
Real agent sessions show how models handle practical, multi-step work: planning, tool use, coding, debugging, corrections and final outcomes. Labs use them for training, evaluation and product improvement. Difficult, coherent, completed sessions are worth more than short or abandoned ones.
How much is a session worth?
The uploader shows an estimate per session before upload. Actual prices depend on quality, task difficulty, model, token volume, completeness, licensing terms and buyer demand. Estimates are not offers.
Does installing the uploader expose my data?
No. Installing or opening the uploader uploads nothing. Your sessions stay on your device until you select them and confirm.
What exactly gets uploaded?
Only the sessions you select and confirm, after local redaction. Selecting alone does not start an upload.
When do I get paid?
When a buyer completes a licensing purchase, not when you upload. A session may earn again if licensed to another buyer. Track submissions and earnings in your Cookiy Earn account.
Do I lose ownership?
You license selected sessions. Nothing is transferred automatically. Exclusivity, deletion, buyer rights, payout timing and taxes are covered by the current Cookiy Earn terms .
Why do I need to install something?
Your history lives on your computer. The uploader needs local access to list sessions, let you choose, scrub private data, estimate value and transfer only what you selected.
Star history
Built by the team behind the Cookiy User Research Skill (1.5K+ stars) and Cookiy.ai , the agentic user-research platform.
Readme MIT license Activity Custom properties Stars
12 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
