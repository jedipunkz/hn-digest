---
source: "https://github.com/rickbauer327-cell/solo-stack-starter"
hn_url: "https://news.ycombinator.com/item?id=49723991"
title: "I gave Claude Code $100 and 30 days to turn a profit"
article_title: "GitHub - rickbauer327-cell/solo-stack-starter: Five free Claude Code skills for the business side of freelancing: business context, meeting notes to actions, invoices, weekly review, LinkedIn posts · GitHub"
image: "https://opengraph.githubassets.com/d115a754bd2d71e4e46b337280fe8de96d0c3b94c99df66f99fd25a8089b2756/rickbauer327-cell/solo-stack-starter"
author: "solo_stacker"
captured_at: "2026-09-16T09:42:27Z"
capture_tool: "hn-digest"
hn_id: 49723991
score: 1
comments: 0
posted_at: "2026-09-16T09:31:15Z"
tags:
  - hacker-news
---

# I gave Claude Code $100 and 30 days to turn a profit

- HN: [49723991](https://news.ycombinator.com/item?id=49723991)
- Source: [github.com](https://github.com/rickbauer327-cell/solo-stack-starter)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T09:31:15Z

## Translation

Title: I gave Claude Code $100 and 30 days to turn a profit
Article title: GitHub - rickbauer327-cell/solo-stack-starter: Five free Claude Code skills for the business side of freelancing: business context, meeting notes to actions, invoices, weekly review, LinkedIn posts · GitHub
Description: Five free Claude Code skills for the business side of freelancing: business context, meeting notes to actions, invoices, weekly review, LinkedIn posts - rickbauer327-cell/solo-stack-starter

Article text:
GitHub - rickbauer327-cell/solo-stack-starter: Five free Claude Code skills for the business side of freelancing: business context, meeting notes to actions, invoices, weekly review, LinkedIn posts · GitHub
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
rickbauer327-cell
/
solo-stack-starter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.claude-plugin .claude-plugin examples examples skills skills templates templates .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Five free Claude Code skills for the business side of working for yourself.
Claude Code is great at code. This makes it useful for the other two-thirds of a freelancer's week: meeting notes, invoices, weekly reviews, and posts — all reading one shared context file so you explain your business once.
# macOS / Linux
git clone https://github.com/rickbauer327-cell/solo-stack-starter
cp -r solo-stack-starter/skills/ * ~ /.claude/skills/
# Windows
git clone https: // github.com / rickbauer327 - cell / solo - stack - starter
Copy-Item - Recurse solo - stack - starter\skills\ * " $ env: USERPROFILE \.claude\skills\ "
Or as a plugin, inside Claude Code:
/plugin marketplace add rickbauer327-cell/solo-stack-starter
/plugin install solo-stack-starter@solo-stack-starter
Then in the folder where you keep your business files:
/setup-business
The five skills
Skill
What it does
/setup-business
Interviews you (in small batches) and writes .solo/business.md — services, rates, terms, tone, clients. Every other skill reads it first.
/meeting-notes-to-actions
Notes or transcript → decisions, owned actions, open questions, and a follow-up email under 150 words. Never upgrades "we could" into "we will".
/invoice-generator
Time log or line items → numbered HTML invoice (prints to PDF) + markdown copy. Scans invoices/ so numbers never repeat.
/weekly-review
Reads your time log, this week's meetings, overdue invoices → what happened, what's stuck, one "must" for next week.
/linkedin-post
Idea or link → three variants in your voice. Hooks under 12 words, no engagement bait.
The pattern
Every skill here follows the same shape, and you can use it for your own:
---
name : skill-name
description : Use when the user <says or needs X>. Produces <Y>.
---
## Step 1 — Load context read .solo/business.md
## Step 2 — Inputs max 3 questions at a time
## Step 3 — Write an output path + a template
## Rules 3–6 opinions about what good looks like
## Do not hard constraints (never invent prices, dates, results)
Three things that matter more than they look: the description is the trigger, so write it with the words you actually say; write files, not chat, so skills can build on each other; and constraints ("under 150 words") beat instructions ("be concise").
Solo Stack has 26 skills covering the whole client lifecycle — lead research, cold outreach, discovery-call prep, proposals, pricing, SOW, kickoff, weekly client updates, scope-creep guard, late-payment sequence, monthly finance summary, case studies, testimonials, a one-file portfolio site, inbox triage, SOPs, decision memos — plus a playbook on chaining and customizing them.
→ Get Solo Stack on Gumroad — $19 during launch.
Five free Claude Code skills for the business side of freelancing: business context, meeting notes to actions, invoices, weekly review, LinkedIn posts
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Five free Claude Code skills for the business side of freelancing: business context, meeting notes to actions, invoices, weekly review, LinkedIn posts - rickbauer327-cell/solo-stack-starter

GitHub - rickbauer327-cell/solo-stack-starter: Five free Claude Code skills for the business side of freelancing: business context, meeting notes to actions, invoices, weekly review, LinkedIn posts · GitHub
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
rickbauer327-cell
/
solo-stack-starter
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
1 Commit 1 Commit Folders and files
.claude-plugin .claude-plugin examples examples skills skills templates templates .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Five free Claude Code skills for the business side of working for yourself.
Claude Code is great at code. This makes it useful for the other two-thirds of a freelancer's week: meeting notes, invoices, weekly reviews, and posts — all reading one shared context file so you explain your business once.
# macOS / Linux
git clone https://github.com/rickbauer327-cell/solo-stack-starter
cp -r solo-stack-starter/skills/ * ~ /.claude/skills/
# Windows
git clone https: // github.com / rickbauer327 - cell / solo - stack - starter
Copy-Item - Recurse solo - stack - starter\skills\ * " $ env: USERPROFILE \.claude\skills\ "
Or as a plugin, inside Claude Code:
/plugin marketplace add rickbauer327-cell/solo-stack-starter
/plugin install solo-stack-starter@solo-stack-starter
Then in the folder where you keep your business files:
/setup-business
The five skills
Skill
What it does
/setup-business
Interviews you (in small batches) and writes .solo/business.md — services, rates, terms, tone, clients. Every other skill reads it first.
/meeting-notes-to-actions
Notes or transcript → decisions, owned actions, open questions, and a follow-up email under 150 words. Never upgrades "we could" into "we will".
/invoice-generator
Time log or line items → numbered HTML invoice (prints to PDF) + markdown copy. Scans invoices/ so numbers never repeat.
/weekly-review
Reads your time log, this week's meetings, overdue invoices → what happened, what's stuck, one "must" for next week.
/linkedin-post
Idea or link → three variants in your voice. Hooks under 12 words, no engagement bait.
The pattern
Every skill here follows the same shape, and you can use it for your own:
---
name : skill-name
description : Use when the user <says or needs X>. Produces <Y>.
---
## Step 1 — Load context read .solo/business.md
## Step 2 — Inputs max 3 questions at a time
## Step 3 — Write an output path + a template
## Rules 3–6 opinions about what good looks like
## Do not hard constraints (never invent prices, dates, results)
Three things that matter more than they look: the description is the trigger, so write it with the words you actually say; write files, not chat, so skills can build on each other; and constraints ("under 150 words") beat instructions ("be concise").
Solo Stack has 26 skills covering the whole client lifecycle — lead research, cold outreach, discovery-call prep, proposals, pricing, SOW, kickoff, weekly client updates, scope-creep guard, late-payment sequence, monthly finance summary, case studies, testimonials, a one-file portfolio site, inbox triage, SOPs, decision memos — plus a playbook on chaining and customizing them.
→ Get Solo Stack on Gumroad — $19 during launch.
Five free Claude Code skills for the business side of freelancing: business context, meeting notes to actions, invoices, weekly review, LinkedIn posts
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
