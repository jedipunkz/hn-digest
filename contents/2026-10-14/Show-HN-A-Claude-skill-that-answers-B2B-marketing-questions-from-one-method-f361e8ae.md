---
source: "https://github.com/zachbrown29/lintel-marketing-playbooks"
hn_url: "https://news.ycombinator.com/item?id=50095555"
title: "Show HN: A Claude skill that answers B2B marketing questions from one method"
article_title: "GitHub - zachbrown29/lintel-marketing-playbooks: A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B. · GitHub"
image: "https://opengraph.githubassets.com/297a5b209a1f013b1db353239ecf18c21efc8a5d06fdcb18201e23072fce2f68/zachbrown29/lintel-marketing-playbooks"
author: "zachbrown29"
captured_at: "2026-10-14T22:04:34Z"
capture_tool: "hn-digest"
hn_id: 50095555
score: 1
comments: 0
posted_at: "2026-10-14T21:25:42Z"
tags:
  - hacker-news
---

# Show HN: A Claude skill that answers B2B marketing questions from one method

- HN: [50095555](https://news.ycombinator.com/item?id=50095555)
- Source: [github.com](https://github.com/zachbrown29/lintel-marketing-playbooks)
- Score: 1
- Comments: 0
- Posted: 2026-10-14T21:25:42Z

## Translation

Title: Show HN: A Claude skill that answers B2B marketing questions from one method
Article title: GitHub - zachbrown29/lintel-marketing-playbooks: A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B. · GitHub
Description: A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B. - zachbrown29/lintel-marketing-playbooks

Article text:
GitHub - zachbrown29/lintel-marketing-playbooks: A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B. · GitHub
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
zachbrown29
lintel-marketing-playbooks Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
references references CREDITS.md CREDITS.md LICENSE LICENSE NOTICE.md NOTICE.md README.md README.md SKILL.md SKILL.md View all files Repository files navigation
A Claude skill that answers B2B software marketing questions with a consistent method instead of generic advice: diagnose first, name the one framework that owns the question, give the build steps and the numbers, and say plainly when the answer is not here.
Built for seed to Series B B2B software companies, and deliberately kept there. At Series C and beyond the method and the build sequences still hold, but the benchmarks were written for smaller companies and the skill will say so instead of handing you a number from the wrong stage.
Questions like these, where most answers you can find are either a listicle or a consultant's opener:
We are at $800k ARR with no marketer. What should we spend, and who should we hire first?
Our demo requests fell off a cliff. What broke?
Should we hire an SDR? Our ACV is $18k and we get six inbound demos a month.
Sales says the leads are garbage, marketing says sales never follows up. How do we settle it?
Leadership will not let us publish prices. What goes on the pricing page?
It answers with a stage, a build sequence in order, a decision rule with its threshold, and a benchmark, and it names the framework behind each one so you can argue with it.
npx skills add zachbrown29/lintel-marketing-playbooks
Or copy this folder into your skills directory by hand:
Claude Code: ~/.claude/skills/b2b-marketing-playbooks/
Claude Desktop: the skills folder for your install
Then ask a marketing question. The router loads first and pulls in only the topic file it needs.
SKILL.md is the router: the method, the 13 topic files and what each one owns, the maturity ladder, and the rules that keep answers consistent. references/diagnosis.md always gets read first, because a complaint and its cause usually live in different files. Everything else loads on demand, so a question costs one topic file, not 118,000 words.
The four tiers, used the same way in every file: Reactive (happens when someone remembers), Building (written down, someone owns it), Operating (runs on a cadence, measured), Compounding (improves itself and survives a person leaving). Work on the next tier, not the last one.
It does not write the copy, run the campaign, or replace knowing your own customers. It does not cover consumer marketing, ecommerce, or enterprise sales motions above the mid-market. It holds no data about any company, so the quality of an answer depends on what you tell it about yours.
Known gaps it will admit to rather than bluff about: review-site programs, website design and information architecture, app store optimisation, layoff communication, executive compensation, and later-stage benchmarks.
Every benchmark is either credited to whoever published it or marked as our own judgement from practice. No figure here is credited to an analyst firm we could not check: those citations were stripped rather than dressed up, which is why most numbers read as judgement. One benchmark keeps its original credit (SDR meeting volume, to The Bridge Group) and says on its own line that we have not verified it against the report. Models created by other people (MEDDIC, SPIN, NPS, Jobs to Be Done, the Bowtie, and others) are credited in CREDITS.md , with a short list of credits still to verify.
It has been tested. On 25 real questions, scored side by side against the 3.5M-word library it was condensed from, it held up: better on 16 questions, worse on 5, level on 4. The judging was not blind and the margin was close to judge noise, so read that as "the short version does not lose much", not as a win.
MIT, see LICENSE . Trademarks belong to their owners and nothing here is endorsed by them, see NOTICE.md . This is general guidance, not professional advice.
A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B. - zachbrown29/lintel-marketing-playbooks

GitHub - zachbrown29/lintel-marketing-playbooks: A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B. · GitHub
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
zachbrown29
lintel-marketing-playbooks Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
4 Commits 4 Commits Folders and files
references references CREDITS.md CREDITS.md LICENSE LICENSE NOTICE.md NOTICE.md README.md README.md SKILL.md SKILL.md View all files Repository files navigation
A Claude skill that answers B2B software marketing questions with a consistent method instead of generic advice: diagnose first, name the one framework that owns the question, give the build steps and the numbers, and say plainly when the answer is not here.
Built for seed to Series B B2B software companies, and deliberately kept there. At Series C and beyond the method and the build sequences still hold, but the benchmarks were written for smaller companies and the skill will say so instead of handing you a number from the wrong stage.
Questions like these, where most answers you can find are either a listicle or a consultant's opener:
We are at $800k ARR with no marketer. What should we spend, and who should we hire first?
Our demo requests fell off a cliff. What broke?
Should we hire an SDR? Our ACV is $18k and we get six inbound demos a month.
Sales says the leads are garbage, marketing says sales never follows up. How do we settle it?
Leadership will not let us publish prices. What goes on the pricing page?
It answers with a stage, a build sequence in order, a decision rule with its threshold, and a benchmark, and it names the framework behind each one so you can argue with it.
npx skills add zachbrown29/lintel-marketing-playbooks
Or copy this folder into your skills directory by hand:
Claude Code: ~/.claude/skills/b2b-marketing-playbooks/
Claude Desktop: the skills folder for your install
Then ask a marketing question. The router loads first and pulls in only the topic file it needs.
SKILL.md is the router: the method, the 13 topic files and what each one owns, the maturity ladder, and the rules that keep answers consistent. references/diagnosis.md always gets read first, because a complaint and its cause usually live in different files. Everything else loads on demand, so a question costs one topic file, not 118,000 words.
The four tiers, used the same way in every file: Reactive (happens when someone remembers), Building (written down, someone owns it), Operating (runs on a cadence, measured), Compounding (improves itself and survives a person leaving). Work on the next tier, not the last one.
It does not write the copy, run the campaign, or replace knowing your own customers. It does not cover consumer marketing, ecommerce, or enterprise sales motions above the mid-market. It holds no data about any company, so the quality of an answer depends on what you tell it about yours.
Known gaps it will admit to rather than bluff about: review-site programs, website design and information architecture, app store optimisation, layoff communication, executive compensation, and later-stage benchmarks.
Every benchmark is either credited to whoever published it or marked as our own judgement from practice. No figure here is credited to an analyst firm we could not check: those citations were stripped rather than dressed up, which is why most numbers read as judgement. One benchmark keeps its original credit (SDR meeting volume, to The Bridge Group) and says on its own line that we have not verified it against the report. Models created by other people (MEDDIC, SPIN, NPS, Jobs to Be Done, the Bowtie, and others) are credited in CREDITS.md , with a short list of credits still to verify.
It has been tested. On 25 real questions, scored side by side against the 3.5M-word library it was condensed from, it held up: better on 16 questions, worse on 5, level on 4. The judging was not blind and the margin was close to judge noise, so read that as "the short version does not lose much", not as a win.
MIT, see LICENSE . Trademarks belong to their owners and nothing here is endorsed by them, see NOTICE.md . This is general guidance, not professional advice.
A Claude skill for B2B software marketing questions: diagnose first, name the framework that owns it, give the numbers. Seed to Series B.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
