---
source: "https://github.com/ankush1982/ai-secure-by-design"
hn_url: "https://news.ycombinator.com/item?id=50367537"
title: "Deterministic security and privacy checks for code from AI assistants"
article_title: "GitHub - ankush1982/ai-secure-by-design: Platform-neutral security and privacy checks for AI-assisted development: a deterministic CLI, MCP server, git hooks and CI action, plus editable scaffolds mapped to NIST SSDF, ISO 27001:2022, DPDP Act 2023 and the AI Security Wiki. · GitHub"
image: "https://repository-images.githubusercontent.com/1265859782/20d09b67-2e05-4738-904a-5312579e331a"
author: "ankush82"
captured_at: "2026-11-04T16:54:35Z"
capture_tool: "hn-digest"
hn_id: 50367537
score: 1
comments: 0
posted_at: "2026-11-04T16:03:48Z"
tags:
  - hacker-news
---

# Deterministic security and privacy checks for code from AI assistants

- HN: [50367537](https://news.ycombinator.com/item?id=50367537)
- Source: [github.com](https://github.com/ankush1982/ai-secure-by-design)
- Score: 1
- Comments: 0
- Posted: 2026-11-04T16:03:48Z

## Translation

Title: Deterministic security and privacy checks for code from AI assistants
Article title: GitHub - ankush1982/ai-secure-by-design: Platform-neutral security and privacy checks for AI-assisted development: a deterministic CLI, MCP server, git hooks and CI action, plus editable scaffolds mapped to NIST SSDF, ISO 27001:2022, DPDP Act 2023 and the AI Security Wiki. · GitHub
Description: Platform-neutral security and privacy checks for AI-assisted development: a deterministic CLI, MCP server, git hooks and CI action, plus editable scaffolds mapped to NIST SSDF, ISO 27001:2022, DPDP Act 2023 and the AI Security Wiki. - ankush1982/ai-secure-by-design

Article text:
GitHub - ankush1982/ai-secure-by-design: Platform-neutral security and privacy checks for AI-assisted development: a deterministic CLI, MCP server, git hooks and CI action, plus editable scaffolds mapped to NIST SSDF, ISO 27001:2022, DPDP Act 2023 and the AI Security Wiki. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
ankush1982
ai-secure-by-design Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Platform-neutral security and privacy checks for AI-assisted development: a deterministic CLI, MCP server, git hooks and CI action, plus editable scaffolds mapped to NIST SSDF, ISO 27001:2022, DPDP Act 2023 and the AI Security Wiki.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
1 fork Report repository Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
32 Commits 32 Commits Folders and files
.claude-plugin .claude-plugin .claude .claude .github .github agents agents commands commands dist dist hooks hooks instructions instructions mcp-server mcp-server regimes regimes rules rules schemas schemas scripts scripts skills skills templates templates .gitattributes .gitattributes .gitignore .gitignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml CONTRIBUTING.md CONTRIBUTING.md LEGAL-DISCLAIMER.md LEGAL-DISCLAIMER.md LICENSE LICENSE NOTICE NOTICE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md action.yml action.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
The same security and privacy checks, whichever AI wrote the code.
A deterministic engine, plus editable scaffolds for consent, audit logs, data rights and AI oversight.
Quick start ·
What it covers ·
What it detects ·
How it works ·
Platforms ·
What this is not ·
Roadmap ·
Contributing
AI coding tools write code quickly and forget rules quietly. Secure & Private by Design moves the rules out of the prompt and into an engine that gives the same answer every time, then runs that engine in your editor, your git hooks and your CI.
Tell it what your product does and which regimes you target. It profiles your codebase, asks you to classify your data, and deploys dev-loop controls (hooks, skills, slash commands, subagents, an MCP server) and product-code scaffolds tailored to your stack.
Secrets and personal data get blocked. Hardcoded keys, Aadhaar numbers, PAN, mobile numbers and emails are stopped in Claude Code, at commit and in CI, and reports never repeat the value.
Checks are code, not prompts. Claude, GPT, Gemini and a human running the CLI all see the same findings, each with a stable rule ID, a severity and a location.
Your manifest is the source of truth. One YAML file records your data categories, regimes, retention and consent. Only the regimes you list are enforced, and the tool never classifies a field for you.
You own the scaffolds. Consent, data rights, audit logging, RBAC, PHI redaction and AI decision logs land in your repo as ordinary code you can edit or delete.
Everything runs locally. No telemetry and no hosted service.
sbd check on the sample app in mcp-server/test/fixtures/clinic-app (trimmed):
Manifest: privacy-security-manifest.yaml (schemaVersion 0.2)
Regimes: nist-ssdf-1.1, iso-27001-2022, dpdp-2023
1. [CRITICAL] SBD-SCAF-001 Missing src/privacy/consent.<ext> (Consent scaffold).
→ Copy templates/typescript-express/consent.ts to src/privacy/consent.ts, or run `sbd deploy`.
10. [HIGH] SBD-FIELD-001 prisma/schema.prisma:5 — Field dateOfBirth looks personal or sensitive but no dataCategory lists it.
12. [HIGH] SBD-CODE-002 src/routes/patients.ts:19 — Log call references a field the manifest classifies as personal or sensitive.
console.log('loaded patient', patient.email);
→ Route the value through the redacting logger in src/security/audit_log.
13. [HIGH] SBD-CODE-003 src/services/triage.ts:6 — Model inference call with no decision-log call within 15 lines.
16. [MEDIUM] SBD-CODE-001 src/routes/patients.ts:12 — Handler mutates data without an auditLog() call within 40 lines.
17 finding(s): 2 critical, 11 high, 3 medium, 1 low, 0 info.
This report lists control gaps. It is not a compliance determination.
What it covers
Regime
Scope
NIST SSDF 1.1
Secure SDLC baseline — applies to every project
ISO 27001:2022
Annex A controls relevant to the dev/SDLC slice (A.5.15, A.5.34, A.8.5, A.8.10–.12, A.8.15, A.8.24–.32)
DPDP Act 2023 (India)
Notice, consent, data principal rights, retention, breach notification, children's data gate, Significant Data Fiduciary flagging
AI Security Wiki (optional)
The community AI Security Wiki control library (46 AI controls) plus AI-specific checks on model output, model loading, MCP servers and agent permissions. A practitioner reference, not a law. See AI Security Wiki integration
Stack
Status
TypeScript + Express
v0.1 ✅ (templates in templates/typescript-express/ )
TypeScript + Next.js
Planned (v0.2+). No templates yet.
Python + FastAPI
v0.2 ✅ (templates in templates/python-fastapi/ ; needs Python 3.10+ and fastapi ). sbd check also reads Python code and models.
Go, Rust, Java, .NET
Later
What it detects
sbd check looks for three kinds of sensitive data. Rules are patterns over your code and data files, so a clean report does not prove the data is absent.
Any secret or literal PII blocks. Every rule above is HIGH or CRITICAL , so each one fails at the default --fail-on high :
Placeholder values such as @example.com addresses are skipped, and the secret rules skip lockfiles, snapshots and fixtures/ folders. To accept reviewed sample data, add sbd: ignore SBD-PII-001 to the line and record why; it shows up in the diff for review, and secret rules cannot be ignored this way. Setting --fail-on critical or none , or running git commit --no-verify , turns local blocking off, so keep the GitHub Action as a required check. Without Node.js, the fallback hook scripts cover secrets only.
The checks that must give the same answer every time are code, not prompts. One engine, thin adapters per platform:
schemas/ + regimes/ + rules/ data, schema-validated
│
sbd engine (mcp-server/src/core) check · scan · validate · plan · explain
output: text, JSON, SARIF 2.1.0; exit codes 0 / 1 / 2
│
┌──────────┼──────────────┬───────────────┬────────────────────────┐
MCP server git hooks CI action instruction packs
stdio+HTTP (pre-commit) (action.yml) CLAUDE.md, AGENTS.md, GEMINI.md,
Copilot, Cursor (one source)
An AI assistant explains findings and proposes fixes. It does not decide what the findings are, so Claude, GPT, Gemini or a human running the CLI all see the same list, and CI enforces it whichever of them wrote the code.
git clone https://github.com/ankush1982/ai-secure-by-design
cd /path/to/your/repo
node /path/to/ai-secure-by-design/dist/sbd.mjs check
Or add it to a Node project as a dev dependency ( npm i -D github:ankush1982/ai-secure-by-design#v0.2.0 ) and run npx sbd check . It is not published to the npm registry yet.
To set up a project without an AI assistant, choose your regimes and run:
sbd init --non-interactive --regime nist-ssdf-1.1,iso-27001-2022 # writes a draft manifest
sbd validate
sbd deploy --dry-run # shows what would change
sbd deploy # copies the scaffolds
sbd plan # every scaffold now shows as present
The draft never decides for you. It leaves dataCategories empty and lists the personal-looking fields as TODO(user): comments, because classification is yours. Without --regime it leaves regimes empty, and validate , plan and deploy stop with a schema error until you add one. Scaffolds that depend on your data, such as consent and DSR handling, are required only once you declare personal data in the manifest.
Exit codes: 0 no findings at or above --fail-on (default high ), 1 findings, 2 usage or runtime error.
With --staged or --diff , sbd check gates the change: it reports findings on added lines, and manifest findings only if the change touches the manifest. Project-wide gaps such as missing scaffolds appear in a full sbd check , which the CI template runs on pushes to the main branch.
Each finding has a stable rule ID ( SBD-SECRET-001 , SBD-PII-001 , SBD-CODE-002 , SBD-SCAF-004 , …), a severity, a location and the regime controls it maps to. Rules live in rules/*.yaml ; add a rule there and every adapter picks it up.
Git hooks (any editor, any AI)
Add to your repo's .pre-commit-config.yaml ( pre-commit ):
repos :
- repo : https://github.com/ankush1982/ai-secure-by-design
rev : v0.2.0 # pin it
hooks :
- id : sbd-secrets
- id : sbd-check
- id : sbd-provenance
Then pre-commit install --hook-type pre-commit --hook-type commit-msg .
Copy templates/.github/workflows/secure-by-design.yml . It uses this repo's action.yml , uploads SARIF to code scanning and fails the build on findings at or above fail-on . Pin the action to a tag or commit SHA.
MCP server (Claude, ChatGPT, Gemini, Copilot, Cursor, …)
node /path/to/ai-secure-by-design/dist/sbd-mcp.mjs --root /path/to/your/repo # stdio (default)
node /path/to/ai-secure-by-design/dist/sbd-mcp.mjs --root /path/to/your/repo --http # http://127.0.0.1:8765/mcp
Tools (all read-only): sbd_validate_manifest , sbd_check , sbd_scan_text , sbd_review_diff , sbd_list_controls , sbd_explain_control , sbd_plan_scaffold , sbd_plan_init , sbd_plan_deploy , sbd_evidence_plan . sbd_plan_init and sbd_plan_deploy return what sbd init and sbd deploy would do. Only the CLI writes files.
Prompts : sbd-init , sbd-check , sbd-deploy (from instructions/prompts/ ).
Resources : sbd://regimes/{id} , sbd://schemas/{manifest,regime,rules} , sbd://manifest .
Every path argument is confined to --root .
HTTP mode binds to 127.0.0.1 with DNS-rebinding protection. It will not bind to a non-loopback address unless SBD_HTTP_TOKEN is set, and then every request needs Authorization: Bearer <token> . Set SBD_ALLOWED_HOSTS to the public hostname when hosting it. The server reads your source tree, so only expose it deliberately.
/plugin marketplace add ankush1982/ai-secure-by-design
/plugin install secure-by-design@ai-secure-by-design
The engine ships prebuilt in dist/ : dist/sbd.mjs (CLI), dist/sbd-mcp.mjs (MCP server) and their shared code in dist/chunks/ , so it needs only Node.js 20+ with no npm install . If Node is missing, the PreToolUse secret hook falls back to a grep -E + jq scan. The bundles are rebuilt with npm run build , and npm test fails if they no longer match the source in mcp-server/src/ .
Slash commands: /sbd-init , /sbd-deploy , /sbd-check (now reports the engine's results instead of doing the checks in prose), /sbd-add , /sbd-evidence .
T

[truncated]

## Original Extract

Platform-neutral security and privacy checks for AI-assisted development: a deterministic CLI, MCP server, git hooks and CI action, plus editable scaffolds mapped to NIST SSDF, ISO 27001:2022, DPDP Act 2023 and the AI Security Wiki. - ankush1982/ai-secure-by-design

GitHub - ankush1982/ai-secure-by-design: Platform-neutral security and privacy checks for AI-assisted development: a deterministic CLI, MCP server, git hooks and CI action, plus editable scaffolds mapped to NIST SSDF, ISO 27001:2022, DPDP Act 2023 and the AI Security Wiki. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
ankush1982
ai-secure-by-design Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Platform-neutral security and privacy checks for AI-assisted development: a deterministic CLI, MCP server, git hooks and CI action, plus editable scaffolds mapped to NIST SSDF, ISO 27001:2022, DPDP Act 2023 and the AI Security Wiki.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
1 fork Report repository Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Latest commit
32 Commits 32 Commits Folders and files
.claude-plugin .claude-plugin .claude .claude .github .github agents agents commands commands dist dist hooks hooks instructions instructions mcp-server mcp-server regimes regimes rules rules schemas schemas scripts scripts skills skills templates templates .gitattributes .gitattributes .gitignore .gitignore .pre-commit-hooks.yaml .pre-commit-hooks.yaml CONTRIBUTING.md CONTRIBUTING.md LEGAL-DISCLAIMER.md LEGAL-DISCLAIMER.md LICENSE LICENSE NOTICE NOTICE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md action.yml action.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
The same security and privacy checks, whichever AI wrote the code.
A deterministic engine, plus editable scaffolds for consent, audit logs, data rights and AI oversight.
Quick start ·
What it covers ·
What it detects ·
How it works ·
Platforms ·
What this is not ·
Roadmap ·
Contributing
AI coding tools write code quickly and forget rules quietly. Secure & Private by Design moves the rules out of the prompt and into an engine that gives the same answer every time, then runs that engine in your editor, your git hooks and your CI.
Tell it what your product does and which regimes you target. It profiles your codebase, asks you to classify your data, and deploys dev-loop controls (hooks, skills, slash commands, subagents, an MCP server) and product-code scaffolds tailored to your stack.
Secrets and personal data get blocked. Hardcoded keys, Aadhaar numbers, PAN, mobile numbers and emails are stopped in Claude Code, at commit and in CI, and reports never repeat the value.
Checks are code, not prompts. Claude, GPT, Gemini and a human running the CLI all see the same findings, each with a stable rule ID, a severity and a location.
Your manifest is the source of truth. One YAML file records your data categories, regimes, retention and consent. Only the regimes you list are enforced, and the tool never classifies a field for you.
You own the scaffolds. Consent, data rights, audit logging, RBAC, PHI redaction and AI decision logs land in your repo as ordinary code you can edit or delete.
Everything runs locally. No telemetry and no hosted service.
sbd check on the sample app in mcp-server/test/fixtures/clinic-app (trimmed):
Manifest: privacy-security-manifest.yaml (schemaVersion 0.2)
Regimes: nist-ssdf-1.1, iso-27001-2022, dpdp-2023
1. [CRITICAL] SBD-SCAF-001 Missing src/privacy/consent.<ext> (Consent scaffold).
→ Copy templates/typescript-express/consent.ts to src/privacy/consent.ts, or run `sbd deploy`.
10. [HIGH] SBD-FIELD-001 prisma/schema.prisma:5 — Field dateOfBirth looks personal or sensitive but no dataCategory lists it.
12. [HIGH] SBD-CODE-002 src/routes/patients.ts:19 — Log call references a field the manifest classifies as personal or sensitive.
console.log('loaded patient', patient.email);
→ Route the value through the redacting logger in src/security/audit_log.
13. [HIGH] SBD-CODE-003 src/services/triage.ts:6 — Model inference call with no decision-log call within 15 lines.
16. [MEDIUM] SBD-CODE-001 src/routes/patients.ts:12 — Handler mutates data without an auditLog() call within 40 lines.
17 finding(s): 2 critical, 11 high, 3 medium, 1 low, 0 info.
This report lists control gaps. It is not a compliance determination.
What it covers
Regime
Scope
NIST SSDF 1.1
Secure SDLC baseline — applies to every project
ISO 27001:2022
Annex A controls relevant to the dev/SDLC slice (A.5.15, A.5.34, A.8.5, A.8.10–.12, A.8.15, A.8.24–.32)
DPDP Act 2023 (India)
Notice, consent, data principal rights, retention, breach notification, children's data gate, Significant Data Fiduciary flagging
AI Security Wiki (optional)
The community AI Security Wiki control library (46 AI controls) plus AI-specific checks on model output, model loading, MCP servers and agent permissions. A practitioner reference, not a law. See AI Security Wiki integration
Stack
Status
TypeScript + Express
v0.1 ✅ (templates in templates/typescript-express/ )
TypeScript + Next.js
Planned (v0.2+). No templates yet.
Python + FastAPI
v0.2 ✅ (templates in templates/python-fastapi/ ; needs Python 3.10+ and fastapi ). sbd check also reads Python code and models.
Go, Rust, Java, .NET
Later
What it detects
sbd check looks for three kinds of sensitive data. Rules are patterns over your code and data files, so a clean report does not prove the data is absent.
Any secret or literal PII blocks. Every rule above is HIGH or CRITICAL , so each one fails at the default --fail-on high :
Placeholder values such as @example.com addresses are skipped, and the secret rules skip lockfiles, snapshots and fixtures/ folders. To accept reviewed sample data, add sbd: ignore SBD-PII-001 to the line and record why; it shows up in the diff for review, and secret rules cannot be ignored this way. Setting --fail-on critical or none , or running git commit --no-verify , turns local blocking off, so keep the GitHub Action as a required check. Without Node.js, the fallback hook scripts cover secrets only.
The checks that must give the same answer every time are code, not prompts. One engine, thin adapters per platform:
schemas/ + regimes/ + rules/ data, schema-validated
│
sbd engine (mcp-server/src/core) check · scan · validate · plan · explain
output: text, JSON, SARIF 2.1.0; exit codes 0 / 1 / 2
│
┌──────────┼──────────────┬───────────────┬────────────────────────┐
MCP server git hooks CI action instruction packs
stdio+HTTP (pre-commit) (action.yml) CLAUDE.md, AGENTS.md, GEMINI.md,
Copilot, Cursor (one source)
An AI assistant explains findings and proposes fixes. It does not decide what the findings are, so Claude, GPT, Gemini or a human running the CLI all see the same list, and CI enforces it whichever of them wrote the code.
git clone https://github.com/ankush1982/ai-secure-by-design
cd /path/to/your/repo
node /path/to/ai-secure-by-design/dist/sbd.mjs check
Or add it to a Node project as a dev dependency ( npm i -D github:ankush1982/ai-secure-by-design#v0.2.0 ) and run npx sbd check . It is not published to the npm registry yet.
To set up a project without an AI assistant, choose your regimes and run:
sbd init --non-interactive --regime nist-ssdf-1.1,iso-27001-2022 # writes a draft manifest
sbd validate
sbd deploy --dry-run # shows what would change
sbd deploy # copies the scaffolds
sbd plan # every scaffold now shows as present
The draft never decides for you. It leaves dataCategories empty and lists the personal-looking fields as TODO(user): comments, because classification is yours. Without --regime it leaves regimes empty, and validate , plan and deploy stop with a schema error until you add one. Scaffolds that depend on your data, such as consent and DSR handling, are required only once you declare personal data in the manifest.
Exit codes: 0 no findings at or above --fail-on (default high ), 1 findings, 2 usage or runtime error.
With --staged or --diff , sbd check gates the change: it reports findings on added lines, and manifest findings only if the change touches the manifest. Project-wide gaps such as missing scaffolds appear in a full sbd check , which the CI template runs on pushes to the main branch.
Each finding has a stable rule ID ( SBD-SECRET-001 , SBD-PII-001 , SBD-CODE-002 , SBD-SCAF-004 , …), a severity, a location and the regime controls it maps to. Rules live in rules/*.yaml ; add a rule there and every adapter picks it up.
Git hooks (any editor, any AI)
Add to your repo's .pre-commit-config.yaml ( pre-commit ):
repos :
- repo : https://github.com/ankush1982/ai-secure-by-design
rev : v0.2.0 # pin it
hooks :
- id : sbd-secrets
- id : sbd-check
- id : sbd-provenance
Then pre-commit install --hook-type pre-commit --hook-type commit-msg .
Copy templates/.github/workflows/secure-by-design.yml . It uses this repo's action.yml , uploads SARIF to code scanning and fails the build on findings at or above fail-on . Pin the action to a tag or commit SHA.
MCP server (Claude, ChatGPT, Gemini, Copilot, Cursor, …)
node /path/to/ai-secure-by-design/dist/sbd-mcp.mjs --root /path/to/your/repo # stdio (default)
node /path/to/ai-secure-by-design/dist/sbd-mcp.mjs --root /path/to/your/repo --http # http://127.0.0.1:8765/mcp
Tools (all read-only): sbd_validate_manifest , sbd_check , sbd_scan_text , sbd_review_diff , sbd_list_controls , sbd_explain_control , sbd_plan_scaffold , sbd_plan_init , sbd_plan_deploy , sbd_evidence_plan . sbd_plan_init and sbd_plan_deploy return what sbd init and sbd deploy would do. Only the CLI writes files.
Prompts : sbd-init , sbd-check , sbd-deploy (from instructions/prompts/ ).
Resources : sbd://regimes/{id} , sbd://schemas/{manifest,regime,rules} , sbd://manifest .
Every path argument is confined to --root .
HTTP mode binds to 127.0.0.1 with DNS-rebinding protection. It will not bind to a non-loopback address unless SBD_HTTP_TOKEN is set, and then every request needs Authorization: Bearer <token> . Set SBD_ALLOWED_HOSTS to the public hostname when hosting it. The server reads your source tree, so only expose it deliberately.
/plugin marketplace add ankush1982/ai-secure-by-design
/plugin install secure-by-design@ai-secure-by-design
The engine ships prebuilt in dist/ : dist/sbd.mjs (CLI), dist/sbd-mcp.mjs (MCP server) and their shared code in dist/chunks/ , so it needs only Node.js 20+ with no npm install . If Node is missing, the PreToolUse secret hook falls back to a grep -E + jq scan. The bundles are rebuilt with npm run build , and npm test fails if they no longer match the source in mcp-server/src/ .
Slash commands: /sbd-init , /sbd-deploy , /sbd-check (now reports the engine's results instead of doing the checks in prose), /sbd-add , /sbd-evidence .
T

[truncated]
