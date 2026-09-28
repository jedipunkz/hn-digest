---
source: "https://github.com/noroles/noroles"
hn_url: "https://news.ycombinator.com/item?id=49881509"
title: "NoRoles – run a company of people and AI agents on permissions"
article_title: "GitHub - noroles/NoRoles: Run a company of people and AI agents on permissions instead of roles. · GitHub"
image: "https://opengraph.githubassets.com/1dfb2d7d9de97ebe3747769ba5a5e7806e9568df34b37e63cd2e36889b012721/noroles/NoRoles"
author: "ilyacherepanov"
captured_at: "2026-09-28T18:31:28Z"
capture_tool: "hn-digest"
hn_id: 49881509
score: 2
comments: 0
posted_at: "2026-09-28T17:34:47Z"
tags:
  - hacker-news
---

# NoRoles – run a company of people and AI agents on permissions

- HN: [49881509](https://news.ycombinator.com/item?id=49881509)
- Source: [github.com](https://github.com/noroles/noroles)
- Score: 2
- Comments: 0
- Posted: 2026-09-28T17:34:47Z

## Translation

Title: NoRoles – run a company of people and AI agents on permissions
Article title: GitHub - noroles/NoRoles: Run a company of people and AI agents on permissions instead of roles. · GitHub
Description: Run a company of people and AI agents on permissions instead of roles. - noroles/NoRoles

Article text:
GitHub - noroles/NoRoles: Run a company of people and AI agents on permissions instead of roles. · GitHub
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
noroles
/
NoRoles
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
.github/ workflows .github/ workflows npm npm .gitignore .gitignore LICENSE LICENSE MANIFESTO.md MANIFESTO.md README.md README.md SPEC.md SPEC.md View all files Repository files navigation
Run a company of people and AI agents on permissions instead of roles. https://noroles.com
Agents now write, build, sell and pay. Companies still run on titles, managers and approvals, so the work waits. NoRoles replaces the org chart with one open list: who can say yes to the few actions that can't be undone (money leaving, contracts, public words, personal data). Everything else, anyone just does. An agent acts with exactly its person's permissions, never more.
npx noroles init my-company
What is here
MANIFESTO.md
eighteen principles: how we think and work
SPEC.md
the model: permissions, mandates, rules, eight laws, and what the orchestrator must enforce
npm/
the noroles command line: start a company, ask for a yes, sign it, run agents with only the keys their mandate allows
How work runs
noroles open first-website # the holders say yes to a mandate once
noroles ask --as me-agent --mandate first-website --permission money.spend \
--summary " buy example.com for 1 year " --amount 12 --currency USD --to Namecheap
noroles yes < id > # a person, in a terminal, sees the exact action and signs
noroles do < id > # runs the approved command once, with only its own key
noroles run first-website --as me-agent -- claude
noroles mcp-config first-website --as me-agent > .mcp.json # every tool call goes through NoRoles
See npm/README.md for what the code enforces today and what it does not yet.
Issues and pull requests are welcome. Every change to how approvals work needs a test in npm/test/laws.test.js that names the law or attack it covers.
Run a company of people and AI agents on permissions instead of roles.
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Run a company of people and AI agents on permissions instead of roles. - noroles/NoRoles

GitHub - noroles/NoRoles: Run a company of people and AI agents on permissions instead of roles. · GitHub
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
noroles
/
NoRoles
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
.github/ workflows .github/ workflows npm npm .gitignore .gitignore LICENSE LICENSE MANIFESTO.md MANIFESTO.md README.md README.md SPEC.md SPEC.md View all files Repository files navigation
Run a company of people and AI agents on permissions instead of roles. https://noroles.com
Agents now write, build, sell and pay. Companies still run on titles, managers and approvals, so the work waits. NoRoles replaces the org chart with one open list: who can say yes to the few actions that can't be undone (money leaving, contracts, public words, personal data). Everything else, anyone just does. An agent acts with exactly its person's permissions, never more.
npx noroles init my-company
What is here
MANIFESTO.md
eighteen principles: how we think and work
SPEC.md
the model: permissions, mandates, rules, eight laws, and what the orchestrator must enforce
npm/
the noroles command line: start a company, ask for a yes, sign it, run agents with only the keys their mandate allows
How work runs
noroles open first-website # the holders say yes to a mandate once
noroles ask --as me-agent --mandate first-website --permission money.spend \
--summary " buy example.com for 1 year " --amount 12 --currency USD --to Namecheap
noroles yes < id > # a person, in a terminal, sees the exact action and signs
noroles do < id > # runs the approved command once, with only its own key
noroles run first-website --as me-agent -- claude
noroles mcp-config first-website --as me-agent > .mcp.json # every tool call goes through NoRoles
See npm/README.md for what the code enforces today and what it does not yet.
Issues and pull requests are welcome. Every change to how approvals work needs a test in npm/test/laws.test.js that names the law or attack it covers.
Run a company of people and AI agents on permissions instead of roles.
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
