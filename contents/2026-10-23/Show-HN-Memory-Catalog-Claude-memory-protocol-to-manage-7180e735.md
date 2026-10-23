---
source: "https://github.com/donalddominko/Memory-Catalog"
hn_url: "https://news.ycombinator.com/item?id=50212810"
title: "Show HN: Memory-Catalog – Claude memory protocol to manage"
article_title: "GitHub - donalddominko/Memory-Catalog: A protocol for organising an AI assistant's long-term memory into an index, backlog, log and \"already considered\" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off. · GitHub"
image: "https://repository-images.githubusercontent.com/1393911103/646f4697-a4e6-4e09-9d77-168813c8293a"
author: "donalddominko"
captured_at: "2026-10-23T18:07:10Z"
capture_tool: "hn-digest"
hn_id: 50212810
score: 1
comments: 0
posted_at: "2026-10-23T17:50:10Z"
tags:
  - hacker-news
---

# Show HN: Memory-Catalog – Claude memory protocol to manage

- HN: [50212810](https://news.ycombinator.com/item?id=50212810)
- Source: [github.com](https://github.com/donalddominko/Memory-Catalog)
- Score: 1
- Comments: 0
- Posted: 2026-10-23T17:50:10Z

## Translation

Title: Show HN: Memory-Catalog – Claude memory protocol to manage
Article title: GitHub - donalddominko/Memory-Catalog: A protocol for organising an AI assistant's long-term memory into an index, backlog, log and "already considered" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off. · GitHub
Description: A protocol for organising an AI assistant's long-term memory into an index, backlog, log and "already considered" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off. - donalddominko/Memory-Catalog

Article text:
GitHub - donalddominko/Memory-Catalog: A protocol for organising an AI assistant's long-term memory into an index, backlog, log and "already considered" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off. · GitHub
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
donalddominko
Memory-Catalog Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
A protocol for organising an AI assistant's long-term memory into an index, backlog, log and "already considered" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off.
Readme AGPL-3.0 and 2 other licenses found Contributing
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.github/ workflows .github/ workflows examples examples templates templates tools tools .gitignore .gitignore .personal-allow .personal-allow CHANGELOG.md CHANGELOG.md CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md INSTALL.md INSTALL.md LICENSE LICENSE LICENSE-DOCS LICENSE-DOCS LICENSE-TEMPLATES LICENSE-TEMPLATES NOTICE NOTICE README.md README.md SPEC.md SPEC.md View all files Repository files navigation
A protocol for organising an AI assistant's long-term memory so it stays useful as it grows: an index of your work streams, a backlog of what's next, a dated log of what happened, a register of ideas already considered , and links from every entry back to the exact chat turn it came from.
If you find it useful, a ⭐ helps others discover it.
It was built on Claude's memory and Projects. Everything is plain Markdown, so the structure carries over to any assistant with file-based memory; Claude-specific parts are named as such in the spec.
Assistant memory tends to turn into a pile. Status is repeated in several places and drifts apart. You can't tell which decisions you made and which the assistant suggested. Ideas you rejected come back as new. Cleaning up risks deleting the one fact that wasn't recorded anywhere else.
The catalog is built on five principles:
One fact, one home. Status lives only in the backlog, history only in the log. Detail that a repository already holds stays there, with a pointer to it. ( SPEC §3 , §11 )
You adopt, the assistant proposes. Anything from the assistant's analysis is filed only after you confirm it. ( SPEC §5 , the adoption gate)
Provenance. Every entry carries a checkpoint such as ⟨a1b2c3d4 t13→t20⟩ : proposed at turn 13 of that chat, accepted at turn 20. Verbatim anchor phrases let the link be verified or repaired if the chat is edited. ( SPEC §6 )
Nothing is lost silently. Ideas you declined, and ideas raised but never discussed, go into an "already considered" register so they don't come back as if new. ( SPEC §5a )
Proof before removal. A cleanup removes nothing that can't be proven to exist elsewhere, and you get a byte-verified backup first. (Safety rule S2, SPEC §13 and §16 )
The five safety rules (S1–S5) that protect your data are in SPEC §13 . New to the vocabulary? See Terms, SPEC §0 .
Requirements, checking the installation, updating and removal are in INSTALL.md . The short version:
The catalog. Open a new chat outside any Project , attach SPEC.md and send:
Set up the catalog system from this spec. Show me what you'll create before writing anything.
The assistant proposes the catalog files and how your existing memory is sorted into them, asks a few questions, and writes only after you approve. Check it worked with status digest .
A Project board (optional). Open a chat inside the Project and send:
Until the nightly sync exists, a board shares status only between chats of that Project. To bring a Project's news into your main catalog, say "read the Project name board and update the catalog" in a main chat. ( SPEC §7 )
Say these in plain words; exact wording doesn't matter. Where: MAIN = a chat outside any Project · PROJECT = a chat inside a Project · ANY = either. This table is copied from SPEC §12 , which is the source of truth.
File
Holds
index.md
Every work stream, its status, where its detail lives
backlog.md
Open tasks by stream: the only place task status lives
log.md
Dated history, older months rolled up
principles.md
Rules of thumb you've adopted, with checkpoints
considered.md
Ideas declined (with your reason) or raised but not discussed
checkpoints.md
Chat keys, turns and anchor phrases for provenance
quotes.md
Planned. Exact key sentences, append-only
Alongside the catalog, each work stream has an area file ( /areas/<stream>.md ) holding its current state. Each Project gets _board.md , _feed.md , _inbox.md and _considered.md . Full reference: SPEC §2 .
SPEC.md the full specification (install this)
INSTALL.md how to install, check, update and remove it; how to run the tools
CHANGELOG.md spec versions
templates/catalog/ empty starter files for main memory
templates/project/ empty starter files for a Project
templates/area.md empty starter file for a work stream
examples/ a worked example with a fictional user
tools/ checks used by the safety rules (bash)
CONTRIBUTING.md how to suggest changes
CLA.md contributor licence agreement
LICENSE, LICENSE-DOCS, LICENSE-TEMPLATES the three licences (see Licence below)
NOTICE copyright notice and reservation of text-and-data-mining rights
The tools prove that a cleanup lost nothing, that a backup is complete, and that nothing personal is published. What each proves: SPEC §15 . How to run them: INSTALL.md .
Specification v0.15 . The index, backlog, log, principles, checkpoints and "already considered" register are in daily use. The quotes file, Project feeds and inboxes, and the nightly sync are specified but not yet running. See the changelog .
If Memory Catalog is useful to you, please leave a star ⭐ at the top of this page. It helps other people find the project and shows that it's worth continuing to develop.
Copyright (C) 2026 Donald Dominko. Different parts use different licences:
Commercial licences for the code, without the AGPL conditions, are available from AI Partner Solutions Ltd. To ask, use the contact form at www.aipartner.london .
Text and data mining, including AI training, is reserved: it is permitted only under the terms of the applicable licence or a separate licence. See NOTICE .
Versions up to and including spec v0.14 (commit e32b809 ) were released under the MIT licence and remain available under it.
A protocol for organising an AI assistant's long-term memory into an index, backlog, log and "already considered" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off.
Readme AGPL-3.0 and 2 other licenses found Contributing
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A protocol for organising an AI assistant's long-term memory into an index, backlog, log and "already considered" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off. - donalddominko/Memory-Catalog

GitHub - donalddominko/Memory-Catalog: A protocol for organising an AI assistant's long-term memory into an index, backlog, log and "already considered" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off. · GitHub
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
donalddominko
Memory-Catalog Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
A protocol for organising an AI assistant's long-term memory into an index, backlog, log and "already considered" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off.
Readme AGPL-3.0 and 2 other licenses found Contributing
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.github/ workflows .github/ workflows examples examples templates templates tools tools .gitignore .gitignore .personal-allow .personal-allow CHANGELOG.md CHANGELOG.md CLA.md CLA.md CONTRIBUTING.md CONTRIBUTING.md INSTALL.md INSTALL.md LICENSE LICENSE LICENSE-DOCS LICENSE-DOCS LICENSE-TEMPLATES LICENSE-TEMPLATES NOTICE NOTICE README.md README.md SPEC.md SPEC.md View all files Repository files navigation
A protocol for organising an AI assistant's long-term memory so it stays useful as it grows: an index of your work streams, a backlog of what's next, a dated log of what happened, a register of ideas already considered , and links from every entry back to the exact chat turn it came from.
If you find it useful, a ⭐ helps others discover it.
It was built on Claude's memory and Projects. Everything is plain Markdown, so the structure carries over to any assistant with file-based memory; Claude-specific parts are named as such in the spec.
Assistant memory tends to turn into a pile. Status is repeated in several places and drifts apart. You can't tell which decisions you made and which the assistant suggested. Ideas you rejected come back as new. Cleaning up risks deleting the one fact that wasn't recorded anywhere else.
The catalog is built on five principles:
One fact, one home. Status lives only in the backlog, history only in the log. Detail that a repository already holds stays there, with a pointer to it. ( SPEC §3 , §11 )
You adopt, the assistant proposes. Anything from the assistant's analysis is filed only after you confirm it. ( SPEC §5 , the adoption gate)
Provenance. Every entry carries a checkpoint such as ⟨a1b2c3d4 t13→t20⟩ : proposed at turn 13 of that chat, accepted at turn 20. Verbatim anchor phrases let the link be verified or repaired if the chat is edited. ( SPEC §6 )
Nothing is lost silently. Ideas you declined, and ideas raised but never discussed, go into an "already considered" register so they don't come back as if new. ( SPEC §5a )
Proof before removal. A cleanup removes nothing that can't be proven to exist elsewhere, and you get a byte-verified backup first. (Safety rule S2, SPEC §13 and §16 )
The five safety rules (S1–S5) that protect your data are in SPEC §13 . New to the vocabulary? See Terms, SPEC §0 .
Requirements, checking the installation, updating and removal are in INSTALL.md . The short version:
The catalog. Open a new chat outside any Project , attach SPEC.md and send:
Set up the catalog system from this spec. Show me what you'll create before writing anything.
The assistant proposes the catalog files and how your existing memory is sorted into them, asks a few questions, and writes only after you approve. Check it worked with status digest .
A Project board (optional). Open a chat inside the Project and send:
Until the nightly sync exists, a board shares status only between chats of that Project. To bring a Project's news into your main catalog, say "read the Project name board and update the catalog" in a main chat. ( SPEC §7 )
Say these in plain words; exact wording doesn't matter. Where: MAIN = a chat outside any Project · PROJECT = a chat inside a Project · ANY = either. This table is copied from SPEC §12 , which is the source of truth.
File
Holds
index.md
Every work stream, its status, where its detail lives
backlog.md
Open tasks by stream: the only place task status lives
log.md
Dated history, older months rolled up
principles.md
Rules of thumb you've adopted, with checkpoints
considered.md
Ideas declined (with your reason) or raised but not discussed
checkpoints.md
Chat keys, turns and anchor phrases for provenance
quotes.md
Planned. Exact key sentences, append-only
Alongside the catalog, each work stream has an area file ( /areas/<stream>.md ) holding its current state. Each Project gets _board.md , _feed.md , _inbox.md and _considered.md . Full reference: SPEC §2 .
SPEC.md the full specification (install this)
INSTALL.md how to install, check, update and remove it; how to run the tools
CHANGELOG.md spec versions
templates/catalog/ empty starter files for main memory
templates/project/ empty starter files for a Project
templates/area.md empty starter file for a work stream
examples/ a worked example with a fictional user
tools/ checks used by the safety rules (bash)
CONTRIBUTING.md how to suggest changes
CLA.md contributor licence agreement
LICENSE, LICENSE-DOCS, LICENSE-TEMPLATES the three licences (see Licence below)
NOTICE copyright notice and reservation of text-and-data-mining rights
The tools prove that a cleanup lost nothing, that a backup is complete, and that nothing personal is published. What each proves: SPEC §15 . How to run them: INSTALL.md .
Specification v0.15 . The index, backlog, log, principles, checkpoints and "already considered" register are in daily use. The quotes file, Project feeds and inboxes, and the nightly sync are specified but not yet running. See the changelog .
If Memory Catalog is useful to you, please leave a star ⭐ at the top of this page. It helps other people find the project and shows that it's worth continuing to develop.
Copyright (C) 2026 Donald Dominko. Different parts use different licences:
Commercial licences for the code, without the AGPL conditions, are available from AI Partner Solutions Ltd. To ask, use the contact form at www.aipartner.london .
Text and data mining, including AI training, is reserved: it is permitted only under the terms of the applicable licence or a separate licence. See NOTICE .
Versions up to and including spec v0.14 (commit e32b809 ) were released under the MIT licence and remain available under it.
A protocol for organising an AI assistant's long-term memory into an index, backlog, log and "already considered" register, with chat-level provenance, a byte-verified backup before any removal, and privacy-safe hand-off.
Readme AGPL-3.0 and 2 other licenses found Contributing
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
