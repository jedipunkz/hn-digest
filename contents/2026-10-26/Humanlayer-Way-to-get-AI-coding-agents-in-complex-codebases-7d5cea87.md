---
source: "https://github.com/humanlayer/humanlayer"
hn_url: "https://news.ycombinator.com/item?id=50246605"
title: "Humanlayer: Way to get AI coding agents in complex codebases"
article_title: "GitHub - humanlayer/humanlayer: The best way to get AI coding agents to solve hard problems in complex codebases. · GitHub"
image: "https://opengraph.githubassets.com/f0957b689018e8b6845ef37252c8fb3e1bf9b36c674565ceed5b6ad98fb38a9d/humanlayer/humanlayer"
author: "Garbage"
captured_at: "2026-10-26T16:33:08Z"
capture_tool: "hn-digest"
hn_id: 50246605
score: 1
comments: 0
posted_at: "2026-10-26T16:32:14Z"
tags:
  - hacker-news
---

# Humanlayer: Way to get AI coding agents in complex codebases

- HN: [50246605](https://news.ycombinator.com/item?id=50246605)
- Source: [github.com](https://github.com/humanlayer/humanlayer)
- Score: 1
- Comments: 0
- Posted: 2026-10-26T16:32:14Z

## Translation

Title: Humanlayer: Way to get AI coding agents in complex codebases
Article title: GitHub - humanlayer/humanlayer: The best way to get AI coding agents to solve hard problems in complex codebases. · GitHub
Description: The best way to get AI coding agents to solve hard problems in complex codebases. - humanlayer/humanlayer

Article text:
GitHub - humanlayer/humanlayer: The best way to get AI coding agents to solve hard problems in complex codebases. · GitHub
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
humanlayer
humanlayer Public Notifications You must be signed in to change notification settings
Star 12.6k ( 12.6k ) You must be signed in to star a repository
The best way to get AI coding agents to solve hard problems in complex codebases.
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
1.0k forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2,121 Commits 2,121 Commits Folders and files
.agents/ context-shards .agents/ context-shards .claude .claude .github .github .humanlayer .humanlayer .vscode .vscode apps apps ci-scripts ci-scripts contracts contracts docker/ durable-streams-server docker/ durable-streams-server golden golden packages packages scripts scripts src/ test/ learning src/ test/ learning tools/ oxlint tools/ oxlint vendor/ effect-orpc vendor/ effect-orpc .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc .oxfmtrc.jsonc .oxfmtrc.jsonc .oxlintrc.jsonc .oxlintrc.jsonc AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Caddyfile Caddyfile HUMANS.md HUMANS.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock bunfig.toml bunfig.toml docker-compose.cluster.yml docker-compose.cluster.yml docker-compose.test.yml docker-compose.test.yml docker-compose.yaml docker-compose.yaml doctor.config.ts doctor.config.ts package.json package.json tsconfig.json tsconfig.json turbo.json turbo.json vitest.config.ts vitest.config.ts View all files Repository files navigation
HumanLayer runs coding agents (Claude Code, Codex, OpenCode, Pi) on your machines and lets your team steer them from one place: tasks, sessions, plans, and reviews, synced between a desktop app, a web app, and a daemon on each host.
This repo holds all of it. It's a Turborepo monorepo that uses Bun .
apps/riptide-ui : the app, as a web app and inside the desktop shells
apps/riptide-electron , apps/riptide-native : the macOS desktop apps (Electron, Tauri)
apps/riptide-daemon : the daemon that runs agent sessions on a host
apps/cli : the humanlayer CLI, which also ships the daemon
apps/riptide-api : the oRPC + Hono API
apps/riptide-electric-proxy : the auth proxy in front of Electric , which syncs Postgres to clients
apps/riptide-cloud : the admin web app (organizations, invites, settings, install page)
apps/riptide-migrator , apps/updater-server , apps/stale-daemon-sweeper , apps/database-maintenance-function : AWS Lambda jobs
apps/riptide-claude-code-plugin , apps/riptide-opencode-plugin , apps/riptide-pi-extension : plugins that mirror Claude Code, OpenCode, and Pi sessions into HumanLayer
apps/riptide-humanlayer-claude-plugin : the built-in HumanLayer skills
apps/riptide-rpi-claude-plugin , apps/linear-cli : the Research-Plan-Implement skills and the Linear CLI
apps/devstack : runs a full local stack per worktree
packages/ : shared libraries ( db , schemas , ui , sync , ai , billing , ...)
contracts/ : the oRPC and Effect RPC contracts between the apps
You need Bun and Docker. Building apps/riptide-native also needs Rust.
bun install
cp .env.example .env
bun --bun run typecheck
bun run test
bun --bun run devstack help explains how to run the whole stack locally. CONTRIBUTING.md lists every check a pull request must pass.
A build from source sends nothing to us: Sentry, PostHog, and the WorkOS client ID come from env, and stay off when unset. .env.example lists them.
HumanLayer Cloud runs this code. Guides for running it yourself are in progress.
⚠️ note - If you use the Docker Compose file as a starter for self hosting, please note that it uses default/simple passwords that should be updated before you deploy.
Apache License 2.0 . A few vendored folders keep their MIT licenses; NOTICE lists them.
The best way to get AI coding agents to solve hard problems in complex codebases.
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
1.0k forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

The best way to get AI coding agents to solve hard problems in complex codebases. - humanlayer/humanlayer

GitHub - humanlayer/humanlayer: The best way to get AI coding agents to solve hard problems in complex codebases. · GitHub
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
humanlayer
humanlayer Public Notifications You must be signed in to change notification settings
Star 12.6k ( 12.6k ) You must be signed in to star a repository
The best way to get AI coding agents to solve hard problems in complex codebases.
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
1.0k forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2,121 Commits 2,121 Commits Folders and files
.agents/ context-shards .agents/ context-shards .claude .claude .github .github .humanlayer .humanlayer .vscode .vscode apps apps ci-scripts ci-scripts contracts contracts docker/ durable-streams-server docker/ durable-streams-server golden golden packages packages scripts scripts src/ test/ learning src/ test/ learning tools/ oxlint tools/ oxlint vendor/ effect-orpc vendor/ effect-orpc .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc .oxfmtrc.jsonc .oxfmtrc.jsonc .oxlintrc.jsonc .oxlintrc.jsonc AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Caddyfile Caddyfile HUMANS.md HUMANS.md LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md bun.lock bun.lock bunfig.toml bunfig.toml docker-compose.cluster.yml docker-compose.cluster.yml docker-compose.test.yml docker-compose.test.yml docker-compose.yaml docker-compose.yaml doctor.config.ts doctor.config.ts package.json package.json tsconfig.json tsconfig.json turbo.json turbo.json vitest.config.ts vitest.config.ts View all files Repository files navigation
HumanLayer runs coding agents (Claude Code, Codex, OpenCode, Pi) on your machines and lets your team steer them from one place: tasks, sessions, plans, and reviews, synced between a desktop app, a web app, and a daemon on each host.
This repo holds all of it. It's a Turborepo monorepo that uses Bun .
apps/riptide-ui : the app, as a web app and inside the desktop shells
apps/riptide-electron , apps/riptide-native : the macOS desktop apps (Electron, Tauri)
apps/riptide-daemon : the daemon that runs agent sessions on a host
apps/cli : the humanlayer CLI, which also ships the daemon
apps/riptide-api : the oRPC + Hono API
apps/riptide-electric-proxy : the auth proxy in front of Electric , which syncs Postgres to clients
apps/riptide-cloud : the admin web app (organizations, invites, settings, install page)
apps/riptide-migrator , apps/updater-server , apps/stale-daemon-sweeper , apps/database-maintenance-function : AWS Lambda jobs
apps/riptide-claude-code-plugin , apps/riptide-opencode-plugin , apps/riptide-pi-extension : plugins that mirror Claude Code, OpenCode, and Pi sessions into HumanLayer
apps/riptide-humanlayer-claude-plugin : the built-in HumanLayer skills
apps/riptide-rpi-claude-plugin , apps/linear-cli : the Research-Plan-Implement skills and the Linear CLI
apps/devstack : runs a full local stack per worktree
packages/ : shared libraries ( db , schemas , ui , sync , ai , billing , ...)
contracts/ : the oRPC and Effect RPC contracts between the apps
You need Bun and Docker. Building apps/riptide-native also needs Rust.
bun install
cp .env.example .env
bun --bun run typecheck
bun run test
bun --bun run devstack help explains how to run the whole stack locally. CONTRIBUTING.md lists every check a pull request must pass.
A build from source sends nothing to us: Sentry, PostHog, and the WorkOS client ID come from env, and stay off when unset. .env.example lists them.
HumanLayer Cloud runs this code. Guides for running it yourself are in progress.
⚠️ note - If you use the Docker Compose file as a starter for self hosting, please note that it uses default/simple passwords that should be updated before you deploy.
Apache License 2.0 . A few vendored folders keep their MIT licenses; NOTICE lists them.
The best way to get AI coding agents to solve hard problems in complex codebases.
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
1.0k forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
