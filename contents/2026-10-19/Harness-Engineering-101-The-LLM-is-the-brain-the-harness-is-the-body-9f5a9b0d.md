---
source: "https://github.com/IsuruMaduranga/harness-engineering-101"
hn_url: "https://news.ycombinator.com/item?id=50151390"
title: "Harness Engineering 101: The LLM is the brain, the harness is the body"
article_title: "GitHub - IsuruMaduranga/harness-engineering-101: Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body. · GitHub"
image: "https://opengraph.githubassets.com/0e4bd36c2232d995526dfe575042e5ed2c518ff4278a342418626de86101a6fe/IsuruMaduranga/harness-engineering-101"
author: "isuruwijesiri"
captured_at: "2026-10-19T12:19:17Z"
capture_tool: "hn-digest"
hn_id: 50151390
score: 2
comments: 0
posted_at: "2026-10-19T12:07:41Z"
tags:
  - hacker-news
---

# Harness Engineering 101: The LLM is the brain, the harness is the body

- HN: [50151390](https://news.ycombinator.com/item?id=50151390)
- Source: [github.com](https://github.com/IsuruMaduranga/harness-engineering-101)
- Score: 2
- Comments: 0
- Posted: 2026-10-19T12:07:41Z

## Translation

Title: Harness Engineering 101: The LLM is the brain, the harness is the body
Article title: GitHub - IsuruMaduranga/harness-engineering-101: Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body. · GitHub
Description: Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body. - IsuruMaduranga/harness-engineering-101

Article text:
GitHub - IsuruMaduranga/harness-engineering-101: Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body. · GitHub
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
IsuruMaduranga
harness-engineering-101 Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body.
Readme License, MIT licenses found Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
25 Commits 25 Commits Folders and files
docs docs scripts scripts src src .gitignore .gitignore LICENSE LICENSE LICENSE-CODE LICENSE-CODE README.md README.md book.toml book.toml external-links-new-tab.js external-links-new-tab.js mermaid-init.js mermaid-init.js mermaid.min.js mermaid.min.js site-back-link.js site-back-link.js View all files Repository files navigation
The LLM is the brain. The harness is the body.
A blog series about building AI agents from first principles, published as a
book.
📖 Read it online: https://isuruwijesiri.com/harness-engineering-101/
An LLM API is stateless. Every turn, you send the entire conversation as a JSON
array and get text back. A harness is the program that builds, maintains, and
protects that array. Everything the field calls "agents" is a set of patches to
that one loop, and each patch exists because something concrete broke. This
series develops that one idea across 16 chapters, growing a ~300-line toy agent
in plain Python as it goes.
The patterns here are developed in full in
One Code , an open-source coding
agent built on pi .
The rendered book reads best online at the link above. If you'd rather read
the source here on GitHub, follow the chapters in order below. The full
introduction is on the book's landing page .
Every chapter has the same shape: the failure, the minimal patch, what the
patch costs.
Part I — The Wire (nothing is magic)
#
Chapter
The failure it patches
1
It's Just a JSON Array
"How do I even talk to this thing?"
2
The Brain: A Next-Token Black Box
"Why does it behave like that?"
3
Tools: JSON Mapped to Functions
The brain can't touch the world
4
The Agent Loop
One tool call isn't enough
5
Caching: Why Order Is Load-Bearing
Resending the array gets expensive
Part II — Running Long (the array under pressure)
#
Chapter
The failure it patches
6
Context Is a Budget, Not a Bag
The window fills up
7
Subagents: Fork the Context
One array can't hold all the work
8
Steering the Running Loop
The brain drifts mid-task
9
Background Work and Time
The loop is synchronous; the world isn't
Part III — The Ecosystem (naming what you already understand)
#
Chapter
The failure it patches
10
Every Framework Is a Wrapper Around Chapter 1
Abstraction anxiety
11
Extending the Body: MCP, Skills, Deferred Loading, Hooks
Capabilities don't fit in the array
12
Debugging the Array
You can't fix the prompt you can't see
Part IV — Trust, Domains, and Data
#
Chapter
The failure it patches
13
Reflexes and Guardrails
The loop will do something dumb
14
Case Study: Coding Agents
How specialized does the body need to be?
15
RAG Was a Harness Pattern All Along
A thousand names for one idea
Epilogue
Build Your Own — the complete toy harness, annotated
Standalone pieces. Read them when you need them.
A. One Harness, Many Brains — model routing, prompt tiers, cost engineering
B. Worktrees and Isolation — giving parallel agents separate copies of the world
C. Modes and Plan Mode — harness-enforced operating states (coding-specific)
D. Retries, Rate Limits, and Streaming — the unglamorous plumbing, plus how streaming works on the wire
Path
What it holds
src/*.md
The chapters and appendices (the book source)
src/SUMMARY.md
The book's table of contents
src/harness/
The runnable toy harness — one file per stage
src/diagrams/
Diagram generators and their generated PNGs
book.toml
mdBook configuration
scripts/sync-to-blog.sh
Builds the book and mounts it into the blog for serving
Building the book locally
cargo install mdbook mdbook-mermaid # or: brew install mdbook (+ mdbook-mermaid binary)
mdbook serve --open # live-reloading preview at http://localhost:3000
mdbook build # static site into ./book
mdbook-mermaid renders the inline mermaid diagrams; without it those blocks
show as raw code.
Regenerating the quantitative diagrams
cd src/diagrams && python3 gen_all.py # Python 3.10+, matplotlib
License
This repository is dual-licensed:
The prose — every chapter and appendix (the .md files under src/ ) —
is licensed under
Creative Commons Attribution-NonCommercial-NoDerivatives 4.0 International
(CC BY-NC-ND 4.0) . You
may share it with credit, but not use it commercially or publish modified
versions. Full text in LICENSE .
The code — the .py files under src/harness/ and src/diagrams/ , and
code snippets reproduced from them in the chapters — is licensed under the
MIT License . Use it freely, with attribution.
Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body.
Readme License, MIT licenses found Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body. - IsuruMaduranga/harness-engineering-101

GitHub - IsuruMaduranga/harness-engineering-101: Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body. · GitHub
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
IsuruMaduranga
harness-engineering-101 Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body.
Readme License, MIT licenses found Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
25 Commits 25 Commits Folders and files
docs docs scripts scripts src src .gitignore .gitignore LICENSE LICENSE LICENSE-CODE LICENSE-CODE README.md README.md book.toml book.toml external-links-new-tab.js external-links-new-tab.js mermaid-init.js mermaid-init.js mermaid.min.js mermaid.min.js site-back-link.js site-back-link.js View all files Repository files navigation
The LLM is the brain. The harness is the body.
A blog series about building AI agents from first principles, published as a
book.
📖 Read it online: https://isuruwijesiri.com/harness-engineering-101/
An LLM API is stateless. Every turn, you send the entire conversation as a JSON
array and get text back. A harness is the program that builds, maintains, and
protects that array. Everything the field calls "agents" is a set of patches to
that one loop, and each patch exists because something concrete broke. This
series develops that one idea across 16 chapters, growing a ~300-line toy agent
in plain Python as it goes.
The patterns here are developed in full in
One Code , an open-source coding
agent built on pi .
The rendered book reads best online at the link above. If you'd rather read
the source here on GitHub, follow the chapters in order below. The full
introduction is on the book's landing page .
Every chapter has the same shape: the failure, the minimal patch, what the
patch costs.
Part I — The Wire (nothing is magic)
#
Chapter
The failure it patches
1
It's Just a JSON Array
"How do I even talk to this thing?"
2
The Brain: A Next-Token Black Box
"Why does it behave like that?"
3
Tools: JSON Mapped to Functions
The brain can't touch the world
4
The Agent Loop
One tool call isn't enough
5
Caching: Why Order Is Load-Bearing
Resending the array gets expensive
Part II — Running Long (the array under pressure)
#
Chapter
The failure it patches
6
Context Is a Budget, Not a Bag
The window fills up
7
Subagents: Fork the Context
One array can't hold all the work
8
Steering the Running Loop
The brain drifts mid-task
9
Background Work and Time
The loop is synchronous; the world isn't
Part III — The Ecosystem (naming what you already understand)
#
Chapter
The failure it patches
10
Every Framework Is a Wrapper Around Chapter 1
Abstraction anxiety
11
Extending the Body: MCP, Skills, Deferred Loading, Hooks
Capabilities don't fit in the array
12
Debugging the Array
You can't fix the prompt you can't see
Part IV — Trust, Domains, and Data
#
Chapter
The failure it patches
13
Reflexes and Guardrails
The loop will do something dumb
14
Case Study: Coding Agents
How specialized does the body need to be?
15
RAG Was a Harness Pattern All Along
A thousand names for one idea
Epilogue
Build Your Own — the complete toy harness, annotated
Standalone pieces. Read them when you need them.
A. One Harness, Many Brains — model routing, prompt tiers, cost engineering
B. Worktrees and Isolation — giving parallel agents separate copies of the world
C. Modes and Plan Mode — harness-enforced operating states (coding-specific)
D. Retries, Rate Limits, and Streaming — the unglamorous plumbing, plus how streaming works on the wire
Path
What it holds
src/*.md
The chapters and appendices (the book source)
src/SUMMARY.md
The book's table of contents
src/harness/
The runnable toy harness — one file per stage
src/diagrams/
Diagram generators and their generated PNGs
book.toml
mdBook configuration
scripts/sync-to-blog.sh
Builds the book and mounts it into the blog for serving
Building the book locally
cargo install mdbook mdbook-mermaid # or: brew install mdbook (+ mdbook-mermaid binary)
mdbook serve --open # live-reloading preview at http://localhost:3000
mdbook build # static site into ./book
mdbook-mermaid renders the inline mermaid diagrams; without it those blocks
show as raw code.
Regenerating the quantitative diagrams
cd src/diagrams && python3 gen_all.py # Python 3.10+, matplotlib
License
This repository is dual-licensed:
The prose — every chapter and appendix (the .md files under src/ ) —
is licensed under
Creative Commons Attribution-NonCommercial-NoDerivatives 4.0 International
(CC BY-NC-ND 4.0) . You
may share it with credit, but not use it commercially or publish modified
versions. Full text in LICENSE .
The code — the .py files under src/harness/ and src/diagrams/ , and
code snippets reproduced from them in the chapters — is licensed under the
MIT License . Use it freely, with attribution.
Harness Engineering 101 — building AI agents from first principles. The LLM is the brain; the harness is the body.
Readme License, MIT licenses found Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
