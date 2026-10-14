---
source: "https://github.com/waywardgeek/ensemble"
hn_url: "https://news.ycombinator.com/item?id=50095802"
title: "The Art of Building AI Coding Agents: The book that builds itself"
article_title: "GitHub - waywardgeek/ensemble: The book that builds an AI coding agent. The agent that builds the next edition. · GitHub"
image: "https://opengraph.githubassets.com/ed931a46b6fa6ef603244033b38017f05b7b8092b234cc44c390fb182db44d70/waywardgeek/ensemble"
author: "WaywardGeek"
captured_at: "2026-10-14T22:04:28Z"
capture_tool: "hn-digest"
hn_id: 50095802
score: 1
comments: 0
posted_at: "2026-10-14T21:44:37Z"
tags:
  - hacker-news
---

# The Art of Building AI Coding Agents: The book that builds itself

- HN: [50095802](https://news.ycombinator.com/item?id=50095802)
- Source: [github.com](https://github.com/waywardgeek/ensemble)
- Score: 1
- Comments: 0
- Posted: 2026-10-14T21:44:37Z

## Translation

Title: The Art of Building AI Coding Agents: The book that builds itself
Article title: GitHub - waywardgeek/ensemble: The book that builds an AI coding agent. The agent that builds the next edition. · GitHub
Description: The book that builds an AI coding agent. The agent that builds the next edition. - waywardgeek/ensemble

Article text:
GitHub - waywardgeek/ensemble: The book that builds an AI coding agent. The agent that builds the next edition. · GitHub
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
waywardgeek
ensemble Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
The book that builds an AI coding agent. The agent that builds the next edition.
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,643 Commits 1,643 Commits Folders and files
.agents/ skills/ ensemble-coding .agents/ skills/ ensemble-coding .githooks .githooks agent agent book book cmd cmd cr/ docs cr/ docs docs docs edition-1 edition-1 internal internal scripts scripts solutions solutions .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE Makefile Makefile README.md README.md TODO.md TODO.md go.mod go.mod go.sum go.sum View all files Repository files navigation
The Art of Building AI Coding Agents
Second edition. An agentic codebook: build the agent, test the teaching, improve the book.
Read the book ·
Chapter guide · Solutions ·
First edition archive
Ensemble is the coding agent built by this book's exercises. It has tools and
supervised jobs, streaming CLI and browser interfaces, persistent conversations,
skills, memory, a fleet of agents and an authenticated gateway. The code is here,
along with the instructions, graders and reviews that shaped it.
The second edition was built with Codex under Bill Cox's direction, using a
student implementation followed by comparative review and revisions to both
teaching and graders. It catches up through the first edition's final exercise,
Gateway, now Chapter 26. Original Chapters 5, 22 and 24 were absorbed into related
chapters while preserving their capabilities.
Crossover has begun. Ensemble has performed real coding tasks with GPT-6
Astra at high reasoning. Making it the everyday tool for the next edition is
ongoing work. Chapter 25 is deliberately a placeholder for that experience;
this publication does not claim Ensemble has already produced Edition 3.
Chapter 24 is an optional, ungraded Codex-endpoint workaround that readers may
skip when their public OAuth route caches satisfactorily.
Path
Purpose
book/
Current second-edition manuscript and publication files
agent/
Current application development tree, with its own Go module
agent/gui/
Optional browser application, in a separate Go module
solutions/chNN/
Accepted cumulative chapter snapshots; use matching book numbers
solutions/evidence/
Retained runs and checks, using original exercise numbers
cmd/grade-edition2/
Current chapter grader command
docs/edition-2-notes/
Student, author and independent review records
edition-1/
Historical first-edition book, application, solutions and original graders
Develop in agent/ ; keep accepted snapshots intact. Chapter 25 has no accepted
solution directory yet. The retired standalone refactor snapshot is the only
chapter under solutions/history/ . Historical manifests and receipts retain
the paths and source revisions at which they were measured; the
relocation record explains the move.
Use the Go version declared in each module. Build from the module directory:
mkdir -p bin
(cd agent && go build -o ../bin/ensemble .)
(cd agent/gui && go build -o ../../bin/ensemble-gui ./cmd)
For API-key use, select LLM_VENDOR , LLM_MODEL and LLM_API_KEY , then run the
CLI from the workspace you want it to use:
/path/to/ensemble/bin/ensemble chat --skills-dir /path/to/ensemble/agent/skills
Choose a model supported by the implementation and available to your account.
The application README documents configuration, authentication,
streaming and persistence; the GUI README covers the browser.
For the optional Chapter 24 route, an existing Codex sign-in supplies the
read-only credential file. This example keeps conversation state outside the
checkout while tools work in the repository:
ensemble_repo= " $PWD "
mkdir -p " $HOME /ensemble-session "
(
cd " $HOME /ensemble-session "
unset LLM_BASE_URL OPENAI_BASE_URL LLM_RESOLVED_MODEL
LLM_VENDOR=openai LLM_MODEL=gpt-6-astra \
OPENAI_REASONING_EFFORT=high OPENAI_BILLING_MODE=codex \
" $ensemble_repo /bin/ensemble-gui " \
--workspace " $ensemble_repo " \
--skills-dir " $ensemble_repo /agent/skills " --save save.json
)
Use the browser address printed by the application. The route reads Codex's
existing credentials without copying or renewing them. See
Chapter 24 for the workaround's evidence and limitations.
A successful answer alone does not prove prompt-cache reuse.
make grade # current Chapter 1 snapshot
make grade-dir CH=26 DIR=./solutions/ch26
go run ./cmd/grade-edition2 -ch 8 # matching snapshot by default
go run ./cmd/grade-edition2 -ch 3 /path/to/your/submission
Revised graders cover Chapters 1–18, 20–23 and 26. Chapter 19 is explanatory,
Chapter 24 is optional and ungraded, and Chapter 25 is reserved. Checks run the
actual application against local fake services without paid inference. The
chapter states the public contract and any separate review or live-use work.
Original grader failures remain failures in the historical record, including
independently accepted exceptions. Revised scores are separate results. Read
the workflow and
enhancement records for the distinctions.
To run an original grader, enter the archive first:
(cd edition-1 && go run ./cmd/grade -ch 1)
Build the book
make book
make ebook
python3 book/build-print.py
The assembled Markdown is committed so it can be read directly in the repository.
The publication instructions describe EPUB navigation, print
page verification and the cover-prefixed PDF. The first-edition publication
files remain in edition-1/book/ .
The book and code are Apache-2.0 licensed . Fork them, change what your
agent needs, and preserve the reasons for the changes. A fresh implementation
is a useful test of how much the instructions actually teach.
The book that builds an AI coding agent. The agent that builds the next edition.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

The book that builds an AI coding agent. The agent that builds the next edition. - waywardgeek/ensemble

GitHub - waywardgeek/ensemble: The book that builds an AI coding agent. The agent that builds the next edition. · GitHub
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
waywardgeek
ensemble Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
The book that builds an AI coding agent. The agent that builds the next edition.
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,643 Commits 1,643 Commits Folders and files
.agents/ skills/ ensemble-coding .agents/ skills/ ensemble-coding .githooks .githooks agent agent book book cmd cmd cr/ docs cr/ docs docs docs edition-1 edition-1 internal internal scripts scripts solutions solutions .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE Makefile Makefile README.md README.md TODO.md TODO.md go.mod go.mod go.sum go.sum View all files Repository files navigation
The Art of Building AI Coding Agents
Second edition. An agentic codebook: build the agent, test the teaching, improve the book.
Read the book ·
Chapter guide · Solutions ·
First edition archive
Ensemble is the coding agent built by this book's exercises. It has tools and
supervised jobs, streaming CLI and browser interfaces, persistent conversations,
skills, memory, a fleet of agents and an authenticated gateway. The code is here,
along with the instructions, graders and reviews that shaped it.
The second edition was built with Codex under Bill Cox's direction, using a
student implementation followed by comparative review and revisions to both
teaching and graders. It catches up through the first edition's final exercise,
Gateway, now Chapter 26. Original Chapters 5, 22 and 24 were absorbed into related
chapters while preserving their capabilities.
Crossover has begun. Ensemble has performed real coding tasks with GPT-6
Astra at high reasoning. Making it the everyday tool for the next edition is
ongoing work. Chapter 25 is deliberately a placeholder for that experience;
this publication does not claim Ensemble has already produced Edition 3.
Chapter 24 is an optional, ungraded Codex-endpoint workaround that readers may
skip when their public OAuth route caches satisfactorily.
Path
Purpose
book/
Current second-edition manuscript and publication files
agent/
Current application development tree, with its own Go module
agent/gui/
Optional browser application, in a separate Go module
solutions/chNN/
Accepted cumulative chapter snapshots; use matching book numbers
solutions/evidence/
Retained runs and checks, using original exercise numbers
cmd/grade-edition2/
Current chapter grader command
docs/edition-2-notes/
Student, author and independent review records
edition-1/
Historical first-edition book, application, solutions and original graders
Develop in agent/ ; keep accepted snapshots intact. Chapter 25 has no accepted
solution directory yet. The retired standalone refactor snapshot is the only
chapter under solutions/history/ . Historical manifests and receipts retain
the paths and source revisions at which they were measured; the
relocation record explains the move.
Use the Go version declared in each module. Build from the module directory:
mkdir -p bin
(cd agent && go build -o ../bin/ensemble .)
(cd agent/gui && go build -o ../../bin/ensemble-gui ./cmd)
For API-key use, select LLM_VENDOR , LLM_MODEL and LLM_API_KEY , then run the
CLI from the workspace you want it to use:
/path/to/ensemble/bin/ensemble chat --skills-dir /path/to/ensemble/agent/skills
Choose a model supported by the implementation and available to your account.
The application README documents configuration, authentication,
streaming and persistence; the GUI README covers the browser.
For the optional Chapter 24 route, an existing Codex sign-in supplies the
read-only credential file. This example keeps conversation state outside the
checkout while tools work in the repository:
ensemble_repo= " $PWD "
mkdir -p " $HOME /ensemble-session "
(
cd " $HOME /ensemble-session "
unset LLM_BASE_URL OPENAI_BASE_URL LLM_RESOLVED_MODEL
LLM_VENDOR=openai LLM_MODEL=gpt-6-astra \
OPENAI_REASONING_EFFORT=high OPENAI_BILLING_MODE=codex \
" $ensemble_repo /bin/ensemble-gui " \
--workspace " $ensemble_repo " \
--skills-dir " $ensemble_repo /agent/skills " --save save.json
)
Use the browser address printed by the application. The route reads Codex's
existing credentials without copying or renewing them. See
Chapter 24 for the workaround's evidence and limitations.
A successful answer alone does not prove prompt-cache reuse.
make grade # current Chapter 1 snapshot
make grade-dir CH=26 DIR=./solutions/ch26
go run ./cmd/grade-edition2 -ch 8 # matching snapshot by default
go run ./cmd/grade-edition2 -ch 3 /path/to/your/submission
Revised graders cover Chapters 1–18, 20–23 and 26. Chapter 19 is explanatory,
Chapter 24 is optional and ungraded, and Chapter 25 is reserved. Checks run the
actual application against local fake services without paid inference. The
chapter states the public contract and any separate review or live-use work.
Original grader failures remain failures in the historical record, including
independently accepted exceptions. Revised scores are separate results. Read
the workflow and
enhancement records for the distinctions.
To run an original grader, enter the archive first:
(cd edition-1 && go run ./cmd/grade -ch 1)
Build the book
make book
make ebook
python3 book/build-print.py
The assembled Markdown is committed so it can be read directly in the repository.
The publication instructions describe EPUB navigation, print
page verification and the cover-prefixed PDF. The first-edition publication
files remain in edition-1/book/ .
The book and code are Apache-2.0 licensed . Fork them, change what your
agent needs, and preserve the reasons for the changes. A fresh implementation
is a useful test of how much the instructions actually teach.
The book that builds an AI coding agent. The agent that builds the next edition.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
