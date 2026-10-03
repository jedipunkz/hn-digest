---
source: "https://github.com/taylormatematica-beep/repoguard"
hn_url: "https://news.ycombinator.com/item?id=49946607"
title: "Show HN: RepoGuard – Architecture linter for AI-generated code (Cursor, Claude)"
article_title: "GitHub - taylormatematica-beep/repoguard: The Architecture Guardian for AI-assisted code. Generate strict .cursorrules and audit PRs. · GitHub"
image: "https://opengraph.githubassets.com/52b565dcb43254d88a39d949f172cb7e1cda9ea8cb940c53bd49c079c4459af1/taylormatematica-beep/repoguard"
author: "taylormatematic"
captured_at: "2026-10-03T19:02:01Z"
capture_tool: "hn-digest"
hn_id: 49946607
score: 2
comments: 0
posted_at: "2026-10-03T18:37:10Z"
tags:
  - hacker-news
---

# Show HN: RepoGuard – Architecture linter for AI-generated code (Cursor, Claude)

- HN: [49946607](https://news.ycombinator.com/item?id=49946607)
- Source: [github.com](https://github.com/taylormatematica-beep/repoguard)
- Score: 2
- Comments: 0
- Posted: 2026-10-03T18:37:10Z

## Translation

Title: Show HN: RepoGuard – Architecture linter for AI-generated code (Cursor, Claude)
Article title: GitHub - taylormatematica-beep/repoguard: The Architecture Guardian for AI-assisted code. Generate strict .cursorrules and audit PRs. · GitHub
Description: The Architecture Guardian for AI-assisted code. Generate strict .cursorrules and audit PRs. - taylormatematica-beep/repoguard

Article text:
GitHub - taylormatematica-beep/repoguard: The Architecture Guardian for AI-assisted code. Generate strict .cursorrules and audit PRs. · GitHub
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
taylormatematica-beep
/
repoguard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
39 Commits 39 Commits Folders and files
bin bin br br curriculohero curriculohero golpehero golpehero test test README.md README.md action.yml action.yml index.html index.html package.json package.json View all files Repository files navigation
The Architecture Guardian for AI-Assisted Codebases.
Stop AI from turning your repository into architectural spaghetti across TypeScript, Python, and Golang.
AI coding assistants ( Cursor, GitHub Copilot, Claude Code, Windsurf ) write 300 lines of code in 5 seconds. However, without strict repository guardrails, they frequently:
Bypass Architectural Layers: Run raw database queries (Prisma, Drizzle, SQLAlchemy, GORM) directly inside UI components or HTTP handlers.
Reinvent Existing Helpers: Write duplicate date/string utilities instead of importing from /utils or shared packages.
Escape Type Safety & Error Handling: Scatter : any in TypeScript or discard errors with _ = err in Go to pass quick compilation.
Leak Sensitive Secrets: Hardcode mock API keys or prefix private secrets with NEXT_PUBLIC_ , bundling them into client-side JS.
RepoGuard acts as an automated architecture supervisor: it generates strict, customized .cursorrules , CLAUDE.md , and .windsurfrules context files, verifies pre-commit diffs, and runs inline audits on every Pull Request.
Run directly in any repository (zero installation required):
npx repoguard-rules init
Or install globally:
npm install -g repoguard-rules
repoguard init
What happens in 2 seconds:
🔍 Auto-detects your tech stack (Next.js, NestJS, Express, FastAPI, Django, Gin, Fiber, Prisma, GORM, etc.).
📝 Generates tailored .cursorrules (for Cursor AI).
🤖 Generates a comprehensive CLAUDE.md (for Claude Code).
🌊 Generates .windsurfrules (for Windsurf IDE).
🛡️ Generates .github/copilot-instructions.md (for GitHub Copilot).
⚙️ Configures pre-commit guard hooks & CI workflow .
Command
Description
npx repoguard-rules init
Scans codebase and generates tailored AI context files.
npx repoguard-rules audit
Evaluates entire codebase and returns an Architectural Health Score (A+ to F) .
npx repoguard-rules audit --format=sarif
Generates standard OASIS SARIF v2.1.0 for GitHub Code Scanning integration.
npx repoguard-rules audit --format=json
Outputs machine-readable JSON for custom CI/CD pipelines.
npx repoguard-rules diff
Audits uncommitted git diffs against architectural rules in real-time.
npx repoguard-rules hook install
Configures local .git/hooks/pre-commit to prevent rule breaches.
npx repoguard-rules rules
Displays all 12 built-in architectural rules and descriptions.
Ignoring Files & Folders ( .repoguardignore )
Add a .repoguardignore file to your root directory to skip specific files or directories:
# .repoguardignore
legacy/
migrations/
test/fixtures/
🛡️ Built-in Architectural Rules
Rule ID
Category
Severity
Guardrail Enforced
RULE-01
Architecture
Error
Prohibits raw ORM/DB queries in UI components and Controllers (TS/JS).
RULE-PY-01
Architecture
Warning / Critical
Enforces FastAPI layer separation; forbids direct DB queries and raw commits ( db.commit() ) inside route handlers.
RULE-GO-01
Architecture
Warning / Critical
Enforces Clean Architecture in Go; prohibits raw database/GORM operations inside Gin, Fiber, or Echo HTTP handlers.
RULE-GO-02
Error Handling
Warning
Flags unchecked errors silenced via blank identifier ( _ = err ) in Go.
RULE-02
Security
Critical
Flags hardcoded secrets, private keys, and API tokens.
RULE-09
Security
Critical
Flags private secrets exposed via public prefixes ( NEXT_PUBLIC_*SECRET* , VITE_*SECRET* ).
RULE-03
Type Safety
Warning
Forbids lazy : any and as any escape hatches in TypeScript.
RULE-04
Code Quality
Info
Enforces structured logging instead of raw console.log .
RULE-05
Next.js / SSR
Error
Prevents hydration mismatch from browser globals ( window / localStorage ).
RULE-06
Security
Critical
Detects SQL injection hazards in raw query string interpolations.
RULE-07
API Design
Warning
Enforces schema validation (Zod/Pydantic) on incoming request payloads.
RULE-08
DRY Principle
Info
Prevents AI assistants from duplicating existing common utility helpers.
🤖 GitHub Action & Security Integration
Add continuous architectural enforcement to your CI/CD pipeline:
# .github/workflows/repoguard.yml
name : RepoGuard Architecture Audit
on : [pull_request]
jobs :
audit :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : actions/setup-node@v4
with :
node-version : 20
- run : npx repoguard-rules audit --strict
GitHub Code Scanning (SARIF):
- run : npx repoguard-rules audit --format=sarif > repoguard.sarif
- uses : github/codeql-action/upload-sarif@v3
with :
sarif_file : repoguard.sarif
💎 Plans & Enterprise Upgrades
RepoGuard is 100% free and open-source for public repositories and local development. For automated CI/CD PR enforcement, private teams, and custom architectural rule engines:
👉 Subscribe to Developer Pro ($12/mo) • Upgrade Team ($39/mo) • 🇧🇷 Pagar no PIX (R$ 67 à vista)
Special thanks to the open source engineers contributing to RepoGuard:
@taylormatematica-beep (Lead Maintainer & Author)
@NihalPN — Authored RULE-PY-01 & FastAPI architectural guardrails (PR #3)
🌐 Documentation & Live Hub: https://taylormatematica-beep.github.io/repoguard/
📦 NPM Registry: https://www.npmjs.com/package/repoguard-rules
🐱 Product Hunt: https://www.producthunt.com/products/repoguard
If RepoGuard helps keep your AI coding clean, consider giving this repository a ⭐ Star !
The Architecture Guardian for AI-assisted code. Generate strict .cursorrules and audit PRs.
taylormatematica-beep.github.io/repoguard/ Topics
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

The Architecture Guardian for AI-assisted code. Generate strict .cursorrules and audit PRs. - taylormatematica-beep/repoguard

GitHub - taylormatematica-beep/repoguard: The Architecture Guardian for AI-assisted code. Generate strict .cursorrules and audit PRs. · GitHub
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
taylormatematica-beep
/
repoguard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
39 Commits 39 Commits Folders and files
bin bin br br curriculohero curriculohero golpehero golpehero test test README.md README.md action.yml action.yml index.html index.html package.json package.json View all files Repository files navigation
The Architecture Guardian for AI-Assisted Codebases.
Stop AI from turning your repository into architectural spaghetti across TypeScript, Python, and Golang.
AI coding assistants ( Cursor, GitHub Copilot, Claude Code, Windsurf ) write 300 lines of code in 5 seconds. However, without strict repository guardrails, they frequently:
Bypass Architectural Layers: Run raw database queries (Prisma, Drizzle, SQLAlchemy, GORM) directly inside UI components or HTTP handlers.
Reinvent Existing Helpers: Write duplicate date/string utilities instead of importing from /utils or shared packages.
Escape Type Safety & Error Handling: Scatter : any in TypeScript or discard errors with _ = err in Go to pass quick compilation.
Leak Sensitive Secrets: Hardcode mock API keys or prefix private secrets with NEXT_PUBLIC_ , bundling them into client-side JS.
RepoGuard acts as an automated architecture supervisor: it generates strict, customized .cursorrules , CLAUDE.md , and .windsurfrules context files, verifies pre-commit diffs, and runs inline audits on every Pull Request.
Run directly in any repository (zero installation required):
npx repoguard-rules init
Or install globally:
npm install -g repoguard-rules
repoguard init
What happens in 2 seconds:
🔍 Auto-detects your tech stack (Next.js, NestJS, Express, FastAPI, Django, Gin, Fiber, Prisma, GORM, etc.).
📝 Generates tailored .cursorrules (for Cursor AI).
🤖 Generates a comprehensive CLAUDE.md (for Claude Code).
🌊 Generates .windsurfrules (for Windsurf IDE).
🛡️ Generates .github/copilot-instructions.md (for GitHub Copilot).
⚙️ Configures pre-commit guard hooks & CI workflow .
Command
Description
npx repoguard-rules init
Scans codebase and generates tailored AI context files.
npx repoguard-rules audit
Evaluates entire codebase and returns an Architectural Health Score (A+ to F) .
npx repoguard-rules audit --format=sarif
Generates standard OASIS SARIF v2.1.0 for GitHub Code Scanning integration.
npx repoguard-rules audit --format=json
Outputs machine-readable JSON for custom CI/CD pipelines.
npx repoguard-rules diff
Audits uncommitted git diffs against architectural rules in real-time.
npx repoguard-rules hook install
Configures local .git/hooks/pre-commit to prevent rule breaches.
npx repoguard-rules rules
Displays all 12 built-in architectural rules and descriptions.
Ignoring Files & Folders ( .repoguardignore )
Add a .repoguardignore file to your root directory to skip specific files or directories:
# .repoguardignore
legacy/
migrations/
test/fixtures/
🛡️ Built-in Architectural Rules
Rule ID
Category
Severity
Guardrail Enforced
RULE-01
Architecture
Error
Prohibits raw ORM/DB queries in UI components and Controllers (TS/JS).
RULE-PY-01
Architecture
Warning / Critical
Enforces FastAPI layer separation; forbids direct DB queries and raw commits ( db.commit() ) inside route handlers.
RULE-GO-01
Architecture
Warning / Critical
Enforces Clean Architecture in Go; prohibits raw database/GORM operations inside Gin, Fiber, or Echo HTTP handlers.
RULE-GO-02
Error Handling
Warning
Flags unchecked errors silenced via blank identifier ( _ = err ) in Go.
RULE-02
Security
Critical
Flags hardcoded secrets, private keys, and API tokens.
RULE-09
Security
Critical
Flags private secrets exposed via public prefixes ( NEXT_PUBLIC_*SECRET* , VITE_*SECRET* ).
RULE-03
Type Safety
Warning
Forbids lazy : any and as any escape hatches in TypeScript.
RULE-04
Code Quality
Info
Enforces structured logging instead of raw console.log .
RULE-05
Next.js / SSR
Error
Prevents hydration mismatch from browser globals ( window / localStorage ).
RULE-06
Security
Critical
Detects SQL injection hazards in raw query string interpolations.
RULE-07
API Design
Warning
Enforces schema validation (Zod/Pydantic) on incoming request payloads.
RULE-08
DRY Principle
Info
Prevents AI assistants from duplicating existing common utility helpers.
🤖 GitHub Action & Security Integration
Add continuous architectural enforcement to your CI/CD pipeline:
# .github/workflows/repoguard.yml
name : RepoGuard Architecture Audit
on : [pull_request]
jobs :
audit :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : actions/setup-node@v4
with :
node-version : 20
- run : npx repoguard-rules audit --strict
GitHub Code Scanning (SARIF):
- run : npx repoguard-rules audit --format=sarif > repoguard.sarif
- uses : github/codeql-action/upload-sarif@v3
with :
sarif_file : repoguard.sarif
💎 Plans & Enterprise Upgrades
RepoGuard is 100% free and open-source for public repositories and local development. For automated CI/CD PR enforcement, private teams, and custom architectural rule engines:
👉 Subscribe to Developer Pro ($12/mo) • Upgrade Team ($39/mo) • 🇧🇷 Pagar no PIX (R$ 67 à vista)
Special thanks to the open source engineers contributing to RepoGuard:
@taylormatematica-beep (Lead Maintainer & Author)
@NihalPN — Authored RULE-PY-01 & FastAPI architectural guardrails (PR #3)
🌐 Documentation & Live Hub: https://taylormatematica-beep.github.io/repoguard/
📦 NPM Registry: https://www.npmjs.com/package/repoguard-rules
🐱 Product Hunt: https://www.producthunt.com/products/repoguard
If RepoGuard helps keep your AI coding clean, consider giving this repository a ⭐ Star !
The Architecture Guardian for AI-assisted code. Generate strict .cursorrules and audit PRs.
taylormatematica-beep.github.io/repoguard/ Topics
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
