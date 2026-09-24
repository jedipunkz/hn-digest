---
source: "https://github.com/ahmtsahin/deutsch-dna"
hn_url: "https://news.ycombinator.com/item?id=49833890"
title: "Show HN: DeutschDNA – German Tutor Skill for Claude Code and Codex"
article_title: "GitHub - ahmtsahin/deutsch-dna: DeutschDNA: a German tutor skill for Claude Code and Codex that remembers your mistakes and the hints that helped. · GitHub"
image: "https://opengraph.githubassets.com/692e4bcf0c619a2bfe5a637e878bee530d8f49673716e71d2fd4939332a3ecb4/ahmtsahin/deutsch-dna"
author: "impala64"
captured_at: "2026-09-24T18:15:49Z"
capture_tool: "hn-digest"
hn_id: 49833890
score: 1
comments: 0
posted_at: "2026-09-24T17:22:37Z"
tags:
  - hacker-news
---

# Show HN: DeutschDNA – German Tutor Skill for Claude Code and Codex

- HN: [49833890](https://news.ycombinator.com/item?id=49833890)
- Source: [github.com](https://github.com/ahmtsahin/deutsch-dna)
- Score: 1
- Comments: 0
- Posted: 2026-09-24T17:22:37Z

## Translation

Title: Show HN: DeutschDNA – German Tutor Skill for Claude Code and Codex
Article title: GitHub - ahmtsahin/deutsch-dna: DeutschDNA: a German tutor skill for Claude Code and Codex that remembers your mistakes and the hints that helped. · GitHub
Description: DeutschDNA: a German tutor skill for Claude Code and Codex that remembers your mistakes and the hints that helped. - ahmtsahin/deutsch-dna

Article text:
GitHub - ahmtsahin/deutsch-dna: DeutschDNA: a German tutor skill for Claude Code and Codex that remembers your mistakes and the hints that helped. · GitHub
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
ahmtsahin
/
deutsch-dna
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
11 Commits 11 Commits Folders and files
.github/ workflows .github/ workflows agents agents demo demo docs docs evals evals references references scripts scripts tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
A German coaching skill for Claude Code and Codex.
You learn German. Your tutor remembers what helped.
Your tutor remembers your mistakes, the hints that helped, and your progress across conversations.
Install and start · Actual conversation · 103 days later
Yesterday: with a hint. Today: a new sentence, without help.
Scripted learner inputs, real engine records. View the result without animation .
Requires Python 3.10+ and Claude Code or Codex. No runtime packages, extra API key, server, or database to set up. Your existing model teaches; progress is stored locally in ~/.deutschdna .
git clone https://github.com/ahmtsahin/deutsch-dna.git " $HOME /.claude/skills/deutsch-dna "
Start a chat with /deutsch-dna .
git clone https://github.com/ahmtsahin/deutsch-dna.git " $HOME /.agents/skills/deutsch-dna "
Start a chat with $deutsch-dna .
The tutor gives you one small German task right away. You can ask for explanations in your own language. Your name and goals are optional.
The clone commands work in macOS/Linux terminals and PowerShell. For project-only installation, Python on Windows, or the one-time permission setup that lets the tutor save progress, see installation and permissions .
A new chat remembers the old hint. This is a timed transcript replay of actual Claude Code replies to scripted learner messages. Both chats happened on the same day, using an isolated learning history. Displayed excerpts preserve the tutor's words; waiting time is compressed.
Full conversation · Captured replies and saved evidence · Static view
Every session starts with your board
Your patterns, what is due first, and which mistakes came back. The tutor then quotes one of your own sentences and puts it in a new situation. This board belongs to the scripted four-month learner and comes straight from the engine.
Say this
What happens
“Correct my German.”
Minimal corrections, with related mistakes grouped by their root cause.
“Give me a hint.”
You repair the sentence; the tutor remembers the help you needed.
“Let's review.”
Due patterns return in fresh situations. Copying the old answer cannot advance mastery.
“I have a meeting tomorrow.”
A roleplay can use your goal and a pattern you are practising.
“Let's practise words.”
Vocabulary from your scenes returns in spaced reviews.
“How am I doing?”
Your recorded progress, with the sentences behind it.
“That wasn't a mistake.”
The tutor can undo the correction and restore its schedule.
During a roleplay, ordinary corrections wait until the debrief. Start with “Restoranda konuşalım” or “Roleplay Restaurant”; finish with “bitir” or “stop roleplay”. The reply includes up to three corrections and useful vocabulary.
Your first minute, roleplays, and more examples .
Your first mistake is still there when you need it. This scripted four-month story runs through the real engine. The callback, original sentence, and saved hint come from its records. Static view .
The first example and teaching milestones survive the rolling history limits. Progress scores describe tracked patterns, not overall German proficiency. Text practice works directly; microphone capture and pronunciation scoring are not included.
python scripts/demo.py --learning-loop
This uses a fresh demo directory. For the four-month story, run python scripts/demo.py ; for the restaurant scene, add --speak .
CLI guide · Command contract · Tests, recording, and GIF generation
Planned: Goethe B1/B2, telc B1/B2, and telc Deutsch Beruf exam modes.
Exploring: voice practice, pronunciation fingerprint, and a local dashboard.
DeutschDNA: a German tutor skill for Claude Code and Codex that remembers your mistakes and the hints that helped.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

DeutschDNA: a German tutor skill for Claude Code and Codex that remembers your mistakes and the hints that helped. - ahmtsahin/deutsch-dna

GitHub - ahmtsahin/deutsch-dna: DeutschDNA: a German tutor skill for Claude Code and Codex that remembers your mistakes and the hints that helped. · GitHub
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
ahmtsahin
/
deutsch-dna
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
11 Commits 11 Commits Folders and files
.github/ workflows .github/ workflows agents agents demo demo docs docs evals evals references references scripts scripts tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md SKILL.md SKILL.md View all files Repository files navigation
A German coaching skill for Claude Code and Codex.
You learn German. Your tutor remembers what helped.
Your tutor remembers your mistakes, the hints that helped, and your progress across conversations.
Install and start · Actual conversation · 103 days later
Yesterday: with a hint. Today: a new sentence, without help.
Scripted learner inputs, real engine records. View the result without animation .
Requires Python 3.10+ and Claude Code or Codex. No runtime packages, extra API key, server, or database to set up. Your existing model teaches; progress is stored locally in ~/.deutschdna .
git clone https://github.com/ahmtsahin/deutsch-dna.git " $HOME /.claude/skills/deutsch-dna "
Start a chat with /deutsch-dna .
git clone https://github.com/ahmtsahin/deutsch-dna.git " $HOME /.agents/skills/deutsch-dna "
Start a chat with $deutsch-dna .
The tutor gives you one small German task right away. You can ask for explanations in your own language. Your name and goals are optional.
The clone commands work in macOS/Linux terminals and PowerShell. For project-only installation, Python on Windows, or the one-time permission setup that lets the tutor save progress, see installation and permissions .
A new chat remembers the old hint. This is a timed transcript replay of actual Claude Code replies to scripted learner messages. Both chats happened on the same day, using an isolated learning history. Displayed excerpts preserve the tutor's words; waiting time is compressed.
Full conversation · Captured replies and saved evidence · Static view
Every session starts with your board
Your patterns, what is due first, and which mistakes came back. The tutor then quotes one of your own sentences and puts it in a new situation. This board belongs to the scripted four-month learner and comes straight from the engine.
Say this
What happens
“Correct my German.”
Minimal corrections, with related mistakes grouped by their root cause.
“Give me a hint.”
You repair the sentence; the tutor remembers the help you needed.
“Let's review.”
Due patterns return in fresh situations. Copying the old answer cannot advance mastery.
“I have a meeting tomorrow.”
A roleplay can use your goal and a pattern you are practising.
“Let's practise words.”
Vocabulary from your scenes returns in spaced reviews.
“How am I doing?”
Your recorded progress, with the sentences behind it.
“That wasn't a mistake.”
The tutor can undo the correction and restore its schedule.
During a roleplay, ordinary corrections wait until the debrief. Start with “Restoranda konuşalım” or “Roleplay Restaurant”; finish with “bitir” or “stop roleplay”. The reply includes up to three corrections and useful vocabulary.
Your first minute, roleplays, and more examples .
Your first mistake is still there when you need it. This scripted four-month story runs through the real engine. The callback, original sentence, and saved hint come from its records. Static view .
The first example and teaching milestones survive the rolling history limits. Progress scores describe tracked patterns, not overall German proficiency. Text practice works directly; microphone capture and pronunciation scoring are not included.
python scripts/demo.py --learning-loop
This uses a fresh demo directory. For the four-month story, run python scripts/demo.py ; for the restaurant scene, add --speak .
CLI guide · Command contract · Tests, recording, and GIF generation
Planned: Goethe B1/B2, telc B1/B2, and telc Deutsch Beruf exam modes.
Exploring: voice practice, pronunciation fingerprint, and a local dashboard.
DeutschDNA: a German tutor skill for Claude Code and Codex that remembers your mistakes and the hints that helped.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
