---
source: "https://github.com/Guksu/suta"
hn_url: "https://news.ycombinator.com/item?id=50040224"
title: "Show HN: Suta – A skill to reduce AI slop in generated UIs"
article_title: "GitHub - Guksu/suta: AI 슬롭을 없애는 UI Skill · GitHub"
image: "https://opengraph.githubassets.com/7834acba03ac15b47f90de58851e7e2ff82bccaef3f3c0447bba01e448f4795c/Guksu/suta"
author: "yung3152"
captured_at: "2026-10-11T06:50:42Z"
capture_tool: "hn-digest"
hn_id: 50040224
score: 1
comments: 0
posted_at: "2026-10-11T06:10:15Z"
tags:
  - hacker-news
---

# Show HN: Suta – A skill to reduce AI slop in generated UIs

- HN: [50040224](https://news.ycombinator.com/item?id=50040224)
- Source: [github.com](https://github.com/Guksu/suta)
- Score: 1
- Comments: 0
- Posted: 2026-10-11T06:10:15Z

## Translation

Title: Show HN: Suta – A skill to reduce AI slop in generated UIs
Article title: GitHub - Guksu/suta: AI 슬롭을 없애는 UI Skill · GitHub
Description: AI 슬롭을 없애는 UI Skill. Contribute to Guksu/suta development by creating an account on GitHub.

Article text:
GitHub - Guksu/suta: AI 슬롭을 없애는 UI Skill · GitHub
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
Guksu
suta Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
134 Commits 134 Commits Folders and files
.agents .agents .claude-plugin .claude-plugin .claude .claude .codex-plugin .codex-plugin .github/ workflows .github/ workflows demo demo docs docs evals evals scripts scripts skills/ suta skills/ suta .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.en.md README.en.md README.ko.md README.ko.md README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Layout, color and motion guidelines, with code for 56 UI patterns, for Claude Code and Codex.
I built suta because AI-generated UI kept repeating the same cards, excessive colors and unnecessary decoration. I wanted to reduce that design AI slop.
한국어 · Live demo · Install · Full guide
Same request, model and settings in Claude Code. Top: without suta (one run). Bottom: run 3 of 4 with suta 0.8. Results vary; generation took longer with suta. Method, all results and limitations · Photo credits
Claude Code — run in the chat:
/plugin marketplace add Guksu/suta
/plugin install suta@suta
Codex — run in a terminal, then approve suta's post-edit check in /hooks :
codex plugin marketplace add Guksu/suta
codex plugin add suta@suta
Other agents / skill only — installs the guidelines and patterns without the plugin's post-edit check:
npx skills add Guksu/suta --skill suta
Install options and migration · Updating · skills.sh listing
After installing, ask your agent:
Use suta to build a login screen with email, password, a sign-in button and a password reset link. Keep it simple and make it usable with a keyboard.
The agent reads the guidelines, selects relevant patterns and adapts their code to your project. Plugin installs also report layout and motion issues after edits. The skill and pattern documentation are in Korean; one recorded check used an English request and produced English UI copy and a reply in English.
Try the patterns in the live demo . Browser checks are still needed; static checks cannot catch every UI problem.
I use suta in my own work and will keep fixing issues and updating the guidelines and patterns. If something feels wrong, open an issue with your request and what happened.
Releases · Screen comparisons · Benchmark · Contributing
The code is MIT (© 2026 Guksu). Licenses for the photos in the comparison image are in Image credits .
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

AI 슬롭을 없애는 UI Skill. Contribute to Guksu/suta development by creating an account on GitHub.

GitHub - Guksu/suta: AI 슬롭을 없애는 UI Skill · GitHub
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
Guksu
suta Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
134 Commits 134 Commits Folders and files
.agents .agents .claude-plugin .claude-plugin .claude .claude .codex-plugin .codex-plugin .github/ workflows .github/ workflows demo demo docs docs evals evals scripts scripts skills/ suta skills/ suta .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.en.md README.en.md README.ko.md README.ko.md README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Layout, color and motion guidelines, with code for 56 UI patterns, for Claude Code and Codex.
I built suta because AI-generated UI kept repeating the same cards, excessive colors and unnecessary decoration. I wanted to reduce that design AI slop.
한국어 · Live demo · Install · Full guide
Same request, model and settings in Claude Code. Top: without suta (one run). Bottom: run 3 of 4 with suta 0.8. Results vary; generation took longer with suta. Method, all results and limitations · Photo credits
Claude Code — run in the chat:
/plugin marketplace add Guksu/suta
/plugin install suta@suta
Codex — run in a terminal, then approve suta's post-edit check in /hooks :
codex plugin marketplace add Guksu/suta
codex plugin add suta@suta
Other agents / skill only — installs the guidelines and patterns without the plugin's post-edit check:
npx skills add Guksu/suta --skill suta
Install options and migration · Updating · skills.sh listing
After installing, ask your agent:
Use suta to build a login screen with email, password, a sign-in button and a password reset link. Keep it simple and make it usable with a keyboard.
The agent reads the guidelines, selects relevant patterns and adapts their code to your project. Plugin installs also report layout and motion issues after edits. The skill and pattern documentation are in Korean; one recorded check used an English request and produced English UI copy and a reply in English.
Try the patterns in the live demo . Browser checks are still needed; static checks cannot catch every UI problem.
I use suta in my own work and will keep fixing issues and updating the guidelines and patterns. If something feels wrong, open an issue with your request and what happened.
Releases · Screen comparisons · Benchmark · Contributing
The code is MIT (© 2026 Guksu). Licenses for the photos in the comparison image are in Image credits .
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
