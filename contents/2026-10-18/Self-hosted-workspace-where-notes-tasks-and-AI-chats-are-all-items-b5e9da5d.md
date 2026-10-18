---
source: "https://github.com/sianachi/nixsites"
hn_url: "https://news.ycombinator.com/item?id=50143999"
title: "Self-hosted workspace where notes, tasks and AI chats are all items"
article_title: "GitHub - sianachi/nixsites · GitHub"
image: "https://opengraph.githubassets.com/67dbcc8d662dc3a694cedf3107905b947348f9962a8a3ab7b516288a77ed98c6/sianachi/nixsites"
author: "osinachi88"
captured_at: "2026-10-18T19:15:14Z"
capture_tool: "hn-digest"
hn_id: 50143999
score: 2
comments: 0
posted_at: "2026-10-18T18:39:42Z"
tags:
  - hacker-news
---

# Self-hosted workspace where notes, tasks and AI chats are all items

- HN: [50143999](https://news.ycombinator.com/item?id=50143999)
- Source: [github.com](https://github.com/sianachi/nixsites)
- Score: 2
- Comments: 0
- Posted: 2026-10-18T18:39:42Z

## Translation

Title: Self-hosted workspace where notes, tasks and AI chats are all items
Article title: GitHub - sianachi/nixsites · GitHub
Description: Contribute to sianachi/nixsites development by creating an account on GitHub.

Article text:
GitHub - sianachi/nixsites · GitHub
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
sianachi
nixsites Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,339 Commits 1,339 Commits Folders and files
.github/ workflows .github/ workflows Nix.Frontend Nix.Frontend apps apps assets assets backend backend deploy deploy docs docs fixtures fixtures packages packages scripts scripts .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore .prettierrc .prettierrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE Nix.slnx Nix.slnx README.md README.md eslint.config.js eslint.config.js mise.toml mise.toml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
A self-hosted workspace for one person or a small group: notes, tasks, habits, journal, files,
calendar and numbers in one item tree, on a laptop and on a phone. Everything stays exportable.
One tree, many views. A body (note, canvas or spreadsheet) is separate from how its children are
shown: list, board, calendar, timeline, gallery, spreadsheet, chart, habit tracker, finance ledger.
Structure you shape yourself. Property schemas, formulas, rollups, templates and automations,
plus a companion that builds structure on request, with every plan previewed before it runs.
Real-time editing. CRDT-synced notes with presence, version history and password-locked subtrees.
Quiet reminders and sync. Due tasks and habits arrive as push notifications; Google and Outlook
calendars sync both ways with a visible log.
Safe by construction. Permissions are enforced by Postgres row-level security, and the web app,
CLI and MCP server share one authorization path.
No lock-in. Lossless .nix archives, plus PDF, DOCX and Markdown export; Markdown, DOCX, PDF
and Obsidian import.
Scriptable. nixctl is a JSON CLI, and nixctl mcp exposes the workspace to AI tools.
Publish a subtree as a static site. This documentation is one: https://nix.nixsites.com .
The full list is in docs/features.md .
One tree of items, shown the way each job needs. All data here is made up.
On a phone. Nix installs as an app, with the same views, a share target and offline copies of notes you have opened.
git clone https://github.com/sianachi/Nix.git && cd Nix
mise install # Node, pnpm, .NET and Go at the pinned versions
pnpm install --frozen-lockfile
pnpm dev # infrastructure, database, identity, then every service
Open http://localhost:5173 and sign in as dev@nix.localhost with NixDev-Password1! .
These are local development credentials only.
pnpm dev is safe to re-run. Use bash scripts/dev.sh --help to start only some services, or
--skip-stack when the Docker infrastructure is already up. If a port is taken, see .env.example .
Working in JetBrains Rider? node --experimental-strip-types scripts/rider-run-configs.mts writes
Run/Debug configurations for each service.
Production is Docker Compose with CI-built multi-architecture images at ghcr.io/sianachi/nix .
Follow the deployment guide and the
backup and restore guide . Back up before every upgrade.
ASP.NET Core 10 and Postgres 16 (row-level security, pgvector) · React 19, TypeScript, Tailwind v4 ·
Node 22 collaboration service (CRDT over WebSocket) · Go workers on RabbitMQ · OIDC login (Zitadel).
apps/web React app, installable as a PWA
apps/collab Collaboration service
apps/go-workers Import, export, plugin, calendar and notification workers
apps/cli nixctl and the MCP server
apps/sites Static site builder and server
backend/ Core API (ASP.NET) and the migrator
packages/ Shared libraries
deploy/ Compose manifests, images, seed and backup scripts
Contributing
Read AGENTS.md first; it links the focused guides for backend, frontend, workers and
validation. Day-to-day commands are in docs/development.md . Before opening a PR, run ./scripts/changed-path-checks.sh --working-tree and then
./scripts/validate-changed.sh --working-tree . Generated OpenAPI and client code is never edited by hand.
Bug reports and ideas are welcome as issues.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to sianachi/nixsites development by creating an account on GitHub.

GitHub - sianachi/nixsites · GitHub
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
sianachi
nixsites Public Notifications You must be signed in to change notification settings
Star 5 ( 5 ) You must be signed in to star a repository
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,339 Commits 1,339 Commits Folders and files
.github/ workflows .github/ workflows Nix.Frontend Nix.Frontend apps apps assets assets backend backend deploy deploy docs docs fixtures fixtures packages packages scripts scripts .dockerignore .dockerignore .editorconfig .editorconfig .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore .prettierrc .prettierrc AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE Nix.slnx Nix.slnx README.md README.md eslint.config.js eslint.config.js mise.toml mise.toml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
A self-hosted workspace for one person or a small group: notes, tasks, habits, journal, files,
calendar and numbers in one item tree, on a laptop and on a phone. Everything stays exportable.
One tree, many views. A body (note, canvas or spreadsheet) is separate from how its children are
shown: list, board, calendar, timeline, gallery, spreadsheet, chart, habit tracker, finance ledger.
Structure you shape yourself. Property schemas, formulas, rollups, templates and automations,
plus a companion that builds structure on request, with every plan previewed before it runs.
Real-time editing. CRDT-synced notes with presence, version history and password-locked subtrees.
Quiet reminders and sync. Due tasks and habits arrive as push notifications; Google and Outlook
calendars sync both ways with a visible log.
Safe by construction. Permissions are enforced by Postgres row-level security, and the web app,
CLI and MCP server share one authorization path.
No lock-in. Lossless .nix archives, plus PDF, DOCX and Markdown export; Markdown, DOCX, PDF
and Obsidian import.
Scriptable. nixctl is a JSON CLI, and nixctl mcp exposes the workspace to AI tools.
Publish a subtree as a static site. This documentation is one: https://nix.nixsites.com .
The full list is in docs/features.md .
One tree of items, shown the way each job needs. All data here is made up.
On a phone. Nix installs as an app, with the same views, a share target and offline copies of notes you have opened.
git clone https://github.com/sianachi/Nix.git && cd Nix
mise install # Node, pnpm, .NET and Go at the pinned versions
pnpm install --frozen-lockfile
pnpm dev # infrastructure, database, identity, then every service
Open http://localhost:5173 and sign in as dev@nix.localhost with NixDev-Password1! .
These are local development credentials only.
pnpm dev is safe to re-run. Use bash scripts/dev.sh --help to start only some services, or
--skip-stack when the Docker infrastructure is already up. If a port is taken, see .env.example .
Working in JetBrains Rider? node --experimental-strip-types scripts/rider-run-configs.mts writes
Run/Debug configurations for each service.
Production is Docker Compose with CI-built multi-architecture images at ghcr.io/sianachi/nix .
Follow the deployment guide and the
backup and restore guide . Back up before every upgrade.
ASP.NET Core 10 and Postgres 16 (row-level security, pgvector) · React 19, TypeScript, Tailwind v4 ·
Node 22 collaboration service (CRDT over WebSocket) · Go workers on RabbitMQ · OIDC login (Zitadel).
apps/web React app, installable as a PWA
apps/collab Collaboration service
apps/go-workers Import, export, plugin, calendar and notification workers
apps/cli nixctl and the MCP server
apps/sites Static site builder and server
backend/ Core API (ASP.NET) and the migrator
packages/ Shared libraries
deploy/ Compose manifests, images, seed and backup scripts
Contributing
Read AGENTS.md first; it links the focused guides for backend, frontend, workers and
validation. Day-to-day commands are in docs/development.md . Before opening a PR, run ./scripts/changed-path-checks.sh --working-tree and then
./scripts/validate-changed.sh --working-tree . Generated OpenAPI and client code is never edited by hand.
Bug reports and ideas are welcome as issues.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
