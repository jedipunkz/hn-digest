---
source: "https://github.com/landoncrabtree/etymon"
hn_url: "https://news.ycombinator.com/item?id=49926105"
title: "Etymon – Get rid of all your Claude.md, AGENTS.md, cursor/rules, mcp.json, etc."
article_title: "GitHub - landoncrabtree/etymon: Replace all of your harness-specific files: CLAUDE.md, AGENTS.md, opencode.json, codex.json, .cursor/rules/, .mcp.json, etc. with Etymon. · GitHub"
image: "https://opengraph.githubassets.com/53f1650208d6cd05402f58f2b181d4476b771db881725b4e9de85492ebd89bd9/landoncrabtree/etymon"
author: "landoncrabtree2"
captured_at: "2026-10-01T20:10:11Z"
capture_tool: "hn-digest"
hn_id: 49926105
score: 2
comments: 2
posted_at: "2026-10-01T19:35:49Z"
tags:
  - hacker-news
---

# Etymon – Get rid of all your Claude.md, AGENTS.md, cursor/rules, mcp.json, etc.

- HN: [49926105](https://news.ycombinator.com/item?id=49926105)
- Source: [github.com](https://github.com/landoncrabtree/etymon)
- Score: 2
- Comments: 2
- Posted: 2026-10-01T19:35:49Z

## Translation

Title: Etymon – Get rid of all your Claude.md, AGENTS.md, cursor/rules, mcp.json, etc.
Article title: GitHub - landoncrabtree/etymon: Replace all of your harness-specific files: CLAUDE.md, AGENTS.md, opencode.json, codex.json, .cursor/rules/, .mcp.json, etc. with Etymon. · GitHub
Description: Replace all of your harness-specific files: CLAUDE.md, AGENTS.md, opencode.json, codex.json, .cursor/rules/, .mcp.json, etc. with Etymon. - landoncrabtree/etymon

Article text:
GitHub - landoncrabtree/etymon: Replace all of your harness-specific files: CLAUDE.md, AGENTS.md, opencode.json, codex.json, .cursor/rules/, .mcp.json, etc. with Etymon. · GitHub
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
landoncrabtree
/
etymon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
8 Commits 8 Commits Folders and files
.agents .agents .github/ workflows .github/ workflows .impeccable .impeccable bin bin docs docs scripts scripts src src testcases testcases tests tests website website .gitignore .gitignore .prettierrc.json .prettierrc.json DESIGN.md DESIGN.md IMPLEMENTATION.md IMPLEMENTATION.md LICENSE LICENSE PRODUCT.md PRODUCT.md PROJECT_PLAN.md PROJECT_PLAN.md README.md README.md TECHNICAL_DETAILS.md TECHNICAL_DETAILS.md package-lock.json package-lock.json package.json package.json test_harness.sh test_harness.sh tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Replace CLAUDE.md , AGENTS.md , opencode.json , .codex/config.toml , .cursor/rules/ , .mcp.json , and other tool-specific setup with one shared manifest and lockfile: etymon.toml and etymon.lock .
Make your repositories portable and AI harness-agnostic with Etymon.
npm install -g etymon
etymon
# Or run without installing globally
npx etymon
Node.js 22+ and Git for repository sources.
etymon init
etymon skills add vercel-labs/skills --skill find-skills
etymon mcp add io.github.upstash/context7 --remote 0
# Create your own skills, agents, and rules
etymon skills create
etymon agents create
etymon rules create
etymon sync --harness claude,codex,opencode
Contributors
etymon init
etymon convert
etymon sync --harness codex,opencode
Contributors
etymon skills find react
etymon mcp find context7
etymon mcp info io.github.upstash/context7
etymon skills add ./skills/checks
etymon agents add augmnt/agents/api-designer.md
etymon rules add ./guidance.md --dest-dir src/api
etymon mcp add ./mcp.json
etymon mcp create
etymon convert --harness claude # Import one tool
etymon list
etymon update # Choose newer external versions
etymon sync # Apply the locked versions
etymon sync --offline # Use cached resources
etymon sync --dry-run # Review changes
etymon sync --allow-lossy # Allow conversion losses, with warnings
etymon doctor
etymon skills remove checks
etymon agents remove reviewer
etymon rules remove guidance
etymon mcp remove docs
etymon skills create --global
etymon sync --global --harness claude,codex
etymon uninstall
Full CLI reference , including headless creation, source resolution, and global options.
Interface
Sources
Skills
Local folders, skills.sh , Git
MCP
Local JSON, MCP Registry
Agents
Local files, Git, file URLs
Rules & Instructions
Local files, Git, file URLs
Claude Code, Codex, Copilot CLI/cloud, VS Code, Gemini, Kiro, Pi, Oh My Pi, OpenCode, Cursor, Antigravity, Roo, Cline, Kilo, Continue, Windsurf, Amp, and Zed.
Supported interfaces by tool · Rule compatibility
Humans: docs/CONTRIBUTING.md .
Agents: run etymon sync , then read the generated project instructions and development skills.
Replace all of your harness-specific files: CLAUDE.md, AGENTS.md, opencode.json, codex.json, .cursor/rules/, .mcp.json, etc. with Etymon.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Replace all of your harness-specific files: CLAUDE.md, AGENTS.md, opencode.json, codex.json, .cursor/rules/, .mcp.json, etc. with Etymon. - landoncrabtree/etymon

GitHub - landoncrabtree/etymon: Replace all of your harness-specific files: CLAUDE.md, AGENTS.md, opencode.json, codex.json, .cursor/rules/, .mcp.json, etc. with Etymon. · GitHub
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
landoncrabtree
/
etymon
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
8 Commits 8 Commits Folders and files
.agents .agents .github/ workflows .github/ workflows .impeccable .impeccable bin bin docs docs scripts scripts src src testcases testcases tests tests website website .gitignore .gitignore .prettierrc.json .prettierrc.json DESIGN.md DESIGN.md IMPLEMENTATION.md IMPLEMENTATION.md LICENSE LICENSE PRODUCT.md PRODUCT.md PROJECT_PLAN.md PROJECT_PLAN.md README.md README.md TECHNICAL_DETAILS.md TECHNICAL_DETAILS.md package-lock.json package-lock.json package.json package.json test_harness.sh test_harness.sh tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Replace CLAUDE.md , AGENTS.md , opencode.json , .codex/config.toml , .cursor/rules/ , .mcp.json , and other tool-specific setup with one shared manifest and lockfile: etymon.toml and etymon.lock .
Make your repositories portable and AI harness-agnostic with Etymon.
npm install -g etymon
etymon
# Or run without installing globally
npx etymon
Node.js 22+ and Git for repository sources.
etymon init
etymon skills add vercel-labs/skills --skill find-skills
etymon mcp add io.github.upstash/context7 --remote 0
# Create your own skills, agents, and rules
etymon skills create
etymon agents create
etymon rules create
etymon sync --harness claude,codex,opencode
Contributors
etymon init
etymon convert
etymon sync --harness codex,opencode
Contributors
etymon skills find react
etymon mcp find context7
etymon mcp info io.github.upstash/context7
etymon skills add ./skills/checks
etymon agents add augmnt/agents/api-designer.md
etymon rules add ./guidance.md --dest-dir src/api
etymon mcp add ./mcp.json
etymon mcp create
etymon convert --harness claude # Import one tool
etymon list
etymon update # Choose newer external versions
etymon sync # Apply the locked versions
etymon sync --offline # Use cached resources
etymon sync --dry-run # Review changes
etymon sync --allow-lossy # Allow conversion losses, with warnings
etymon doctor
etymon skills remove checks
etymon agents remove reviewer
etymon rules remove guidance
etymon mcp remove docs
etymon skills create --global
etymon sync --global --harness claude,codex
etymon uninstall
Full CLI reference , including headless creation, source resolution, and global options.
Interface
Sources
Skills
Local folders, skills.sh , Git
MCP
Local JSON, MCP Registry
Agents
Local files, Git, file URLs
Rules & Instructions
Local files, Git, file URLs
Claude Code, Codex, Copilot CLI/cloud, VS Code, Gemini, Kiro, Pi, Oh My Pi, OpenCode, Cursor, Antigravity, Roo, Cline, Kilo, Continue, Windsurf, Amp, and Zed.
Supported interfaces by tool · Rule compatibility
Humans: docs/CONTRIBUTING.md .
Agents: run etymon sync , then read the generated project instructions and development skills.
Replace all of your harness-specific files: CLAUDE.md, AGENTS.md, opencode.json, codex.json, .cursor/rules/, .mcp.json, etc. with Etymon.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
