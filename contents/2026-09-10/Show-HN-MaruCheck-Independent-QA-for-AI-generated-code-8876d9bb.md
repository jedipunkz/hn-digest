---
source: "https://github.com/Kidus-M/MaruCheck"
hn_url: "https://news.ycombinator.com/item?id=49644238"
title: "Show HN: MaruCheck – Independent QA for AI-generated code"
article_title: "GitHub - Kidus-M/MaruCheck: Independent QA and verification for AI-generated software. MaruCheck turns product intent into Quality Contracts, analyzes code changes, remembers regressions, and tests what your coding agent missed. · GitHub"
image: "https://repository-images.githubusercontent.com/1335303916/94d5ea4d-853b-4a72-b731-98969dc6701e"
author: "KidusMT"
captured_at: "2026-09-10T14:52:48Z"
capture_tool: "hn-digest"
hn_id: 49644238
score: 1
comments: 0
posted_at: "2026-09-10T14:24:38Z"
tags:
  - hacker-news
---

# Show HN: MaruCheck – Independent QA for AI-generated code

- HN: [49644238](https://news.ycombinator.com/item?id=49644238)
- Source: [github.com](https://github.com/Kidus-M/MaruCheck)
- Score: 1
- Comments: 0
- Posted: 2026-09-10T14:24:38Z

## Translation

Title: Show HN: MaruCheck – Independent QA for AI-generated code
Article title: GitHub - Kidus-M/MaruCheck: Independent QA and verification for AI-generated software. MaruCheck turns product intent into Quality Contracts, analyzes code changes, remembers regressions, and tests what your coding agent missed. · GitHub
Description: Independent QA and verification for AI-generated software. MaruCheck turns product intent into Quality Contracts, analyzes code changes, remembers regressions, and tests what your coding agent missed. - Kidus-M/MaruCheck
HN text: Hi HN, I've been building MaruCheck, an independent, open-source QA verification tool aimed to solve the issues that arise when coding agents like codex, claude, cursor make semantic issues that has heavy ramifications. Take the example where the requirements says a free user gets 5 uploads. The code changes this behavior to 10. Test generated from the implementation may simply start expecting 10. But Marucheck treats the existing approved requirements as a separate evidence and flags the behavior change. Some of the things implemented currently are Quality contracts, repository and stack scanning, Git diff/change- impact analysis, risk based analysis, semantic drift detection, QA memory for previous bugs/regression, CLI workflows, integration for coding agents, Github/CI verification etc etc.. Another area I’m experimenting with is QA Memory. If six months ago a bug existed because users could access another customer’s invoice by changing an invoice ID, MaruCheck can associate that regression with the relevant files/contracts. If those areas change again later, that previous failure becomes part of the new verification plan. The project is local-first and I’m trying not to reinvent existing testing infrastructure unnecessarily. The idea is to orchestrate tools like Playwright/Vitest/security/Jest tooling rather than build another browser automation framework from scratch. Currently it's on NPM and you can runit using npx or npm by writiing npx --yes marucheck@0.3.0 init and all the information and docs are available on the site marucheck.dev. I have decided to build it in the open and I'm especially interested in feedback on all aspects. Contributors are very welcome as well - issues, adapters, testing, docs, bug reports, architecture criticisms, all of it. Thanks

Article text:
GitHub - Kidus-M/MaruCheck: Independent QA and verification for AI-generated software. MaruCheck turns product intent into Quality Contracts, analyzes code changes, remembers regressions, and tests what your coding agent missed. · GitHub
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
Kidus-M
/
MaruCheck
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
173 Commits 173 Commits Folders and files
.github .github docs docs examples examples packages packages scripts scripts .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json prettier.config.mjs prettier.config.mjs tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Test what your AI didn't. MaruCheck verifies AI-written code against a specification the AI
cannot edit. Local-first: no account, no model API key, and no code leaves your machine.
An agent is asked to fix a bug report: paying users are throttled right after upgrading. It
changes the quota code and updates the tests it owns.
// src/quota.ts
- export const FREE_MONTHLY_GENERATION_LIMIT = 10;
+ export const FREE_MONTHLY_GENERATION_LIMIT = 1000;
- export function resolvePlan(subscription: StoredSubscription): PlanTier {
- return subscription.plan;
+ export function resolvePlan(subscription: StoredSubscription, request?: GenerationRequest): PlanTier {
+ return request?.claimedPlan ?? subscription.plan;
}
// src/quota.test.ts
- it("resolves the plan from the stored subscription", () => {
+ it("honors the plan claimed by the client", () => {
The suite is green. The product is wrong: the free tier gives away a hundred times its quota, and
the plan tier now comes from a value the browser controls. Nothing in a normal pipeline objects,
because the same actor wrote the implementation and the thing that judges the implementation. A
passing suite proves internal consistency, not approved behavior.
MaruCheck keeps that judgment separate:
$ maru verify --diff
Verification gate: BLOCKED
Findings: 5 (5 blocking)
[HIGH] BLOCKING finding-001-usage-quota-quota-001: QUOTA-001 verification failed
Expected: Free plan users may perform at most 10 generations per calendar month.
Actual: Received: "pro"
$ maru drift check --from observations.json
Semantic drift: BLOCKED Conflicts: 2 (2 blocking)
[usage-quota#QUOTA-001] BLOCKING
Contract: Free plan users may perform at most 10 generations per calendar month.
Observed: Free plan users may perform at most 1000 generations per calendar month.
Approved behavior lives in a Quality Contract : a human-owned, versioned file outside the code
under test. An agent can propose a change to the code. It cannot quietly move the goalposts.
git clone https://github.com/Kidus-M/MaruCheck.git
cd MaruCheck
npm install && npm run build
node examples/quota-app/run.mjs
It builds a throwaway Git workspace, approves a contract, applies the agent's change, shows the
green suite, and then blocks the change — about a minute end to end. Read
examples/quota-app for what each file does and what to change to
see the gate behave differently.
The same story on Jest is node examples/quota-app-jest/run.mjs .
Stop the agent from declaring victory
Verification only helps if it runs. When an AI agent writes the change, the agent decides whether to
run it — and the agent is exactly who benefits from skipping it. So hand the decision to the harness:
maru hook install
That registers verification as a Claude Code Stop hook . The agent cannot end a turn while the
gate is blocked; it gets the failing requirements, the expected and actual behavior, and an explicit
instruction that editing the contract is not an available fix. The gate never wedges a session — it
gives up after three consecutive blocks and hands control back to you.
Requirements: Node.js 24 LTS and npm 11 or newer ( why ).
npx --yes marucheck@0.4.0 init
npx --yes marucheck@0.4.0 doctor
npx --yes marucheck@0.4.0 verify --diff
For regular project or team use, pin the exact public package and prevent implicit downloads:
npm install --save-dev --save-exact marucheck@0.4.0
npx --no-install maru --help
See the recommended first workflow before adding
hosted reporting, MCP, or a required CI gate, and the
public installation and release guide for CI
pinning, manual release steps, optional trusted publishing, and rollback.
On every diff, MaruCheck scores risk deterministically, selects the requirements and tests that
this change touches, runs them locally (Vitest, Playwright, axe, Semgrep, Gitleaks), mutation-tests
to prove those tests can still fail, and compares observed behavior against protected invariants.
It remembers confirmed bugs and forces recorded regression tests back into the plan when related
code changes again. Evidence lands in .maru/ as files you can read and argue with, not a
confidence score.
There is an MCP server so your agent can request
verification itself. It can ask for a verdict; it cannot grant one.
The published bundle is built for the Node 24 target and CI runs the full quality gate on Node 24
only, so that is what the engines field claims. Nothing in the source is known to need Node 24
specifically — the CLI has been observed running on older releases — but "not known to break" is
not verification. CI now runs an informational Node 22 job; widening the supported range once that
job is green is
a good first issue .
The hosted Next.js application at marucheck.dev is maintained separately
in the sibling MaruCheck-Web repository so the CLI and
cloud product can release independently. This repository owns the local-first maru CLI,
verification libraries, Git analysis, Quality Contract support, and the MCP server.
Contributors changing the CLI itself build from source:
npm install
npm run check
npm run maru -- --help
During source development, invoke the built CLI from the project you want to inspect:
# In maru-cli
npm run build
# In a Next.js/React project
node ../maru-cli/packages/cli/dist/index.js init
node ../maru-cli/packages/cli/dist/index.js scan
node ../maru-cli/packages/cli/dist/index.js doctor
node ../maru-cli/packages/cli/dist/index.js contract create --from requirements.md
node ../maru-cli/packages/cli/dist/index.js risk --diff
node ../maru-cli/packages/cli/dist/index.js plan --diff
node ../maru-cli/packages/cli/dist/index.js verify --diff
node ../maru-cli/packages/cli/dist/index.js mutate --diff --max 20
node ../maru-cli/packages/cli/dist/index.js ci init
node ../maru-cli/packages/cli/dist/index.js mcp
Commands
Command
Description
npm run build
Build workspaces and the public executable bundle
npm run lint
Run ESLint
npm run format:check
Check formatting
npm run typecheck
Type-check all packages
npm test
Run Vitest tests
npm run check
Run every local quality gate
npm run example
Run the end-to-end example in examples/quota-app
npm run example:jest
Run the same example on Jest
npm run release:check
Check code and inspect the npm tarball
npm run maru -- --help
Exercise the workspace CLI build
Project commands
Command
Description
maru init
Detect the stack and create an idempotent .maru/ configuration
maru scan
Write route, test, dependency, CI, and source inventory to .maru/generated/project-scan.json
maru doctor
Validate Node.js, Git, package-manager, configuration, test, and CI prerequisites
maru risk --diff
Score current changes with deterministic explanations
maru plan --diff
Write an inspectable, requirement-linked verification plan
maru verify --diff
Execute selected tests and write evidence, findings, terminal output, and JSON report
maru upload [--report <path>] [--url <host>]
Explicitly send the newest completed report to a connected dashboard project
maru mutate --diff [--max 20]
Prove selected tests reject isolated TypeScript mutations
maru challenge prepare/submit
Exchange a bounded adversarial brief with a fresh AI-client QA context
maru hook install
Register verification as a Claude Code Stop hook so an agent cannot finish on a blocked gate
maru hook uninstall
Remove the MaruCheck Stop hook and leave every other hook in place
maru ci init
Install an idempotent least-privilege GitHub pull-request workflow
maru ci verify
Verify, publish a GitHub summary, and return the ProofLayer check status
maru drift check --from observations.json
Block approved semantic conflicts without rewriting the contract
maru memory search "authorization"
Query historical bugs, root causes, linked files, contracts, and regression tests
Quality Contract commands
Command
Description
maru contract create --from requirements.md
Create a deterministic draft from natural-language intent
maru contract list
List current contracts, states, criticality, and version IDs
maru contract show <id>
Print one validated contract
maru contract validate [path]
Validate all current contracts or one YAML file
maru contract diff <id-or-path> <id-or-path>
Classify mechanical and semantic contract changes
maru contract approve <id> --by <accountable-owner>
Approve and snapshot one reviewed version
Semantic drift commands
Command
Description
maru drift check --from observations.json
Compare observed behavior with protected requirements/invariants
maru drift propose <id> --from observations.json --reason "Why" --by <proposer>
Write an immutable pending amendment without changing the contract
maru drift approve <proposal-path> --by <contract-owner>
Apply a reviewed amendment with an owner approval and audit record
QA memory commands
Command
Description
maru memory add --from memory.json
Store one immutable versioned historical QA record
maru memory list
List active records newest first
maru memory search "invoice authorization"
Search IDs, defects, root causes, paths, contracts, and tags
maru memory show <MEM-id>
Print one complete record including linked regression tests
Repository structure
packages/
|-- challenger/ # activation policy, adversarial output validation, cost, and reports
|-- ci/ # GitHub workflow installation, summaries, and check conclusions
|-- cli/ # maru command-line interface
|-- contracts/ # Quality Contract schemas and versioning
|-- core/ # verification domain and orchestration
|-- drift/ # protected expectations and contract amendment workflow
|-- evidence/ # requirement evidence, findings, gates, and reports
|-- execution/ # test, accessibility, and security adapters plus raw run artifacts
|-- git/ # repository and diff analysis
|-- mcp-server

[truncated]

## Original Extract

Independent QA and verification for AI-generated software. MaruCheck turns product intent into Quality Contracts, analyzes code changes, remembers regressions, and tests what your coding agent missed. - Kidus-M/MaruCheck

Hi HN, I've been building MaruCheck, an independent, open-source QA verification tool aimed to solve the issues that arise when coding agents like codex, claude, cursor make semantic issues that has heavy ramifications. Take the example where the requirements says a free user gets 5 uploads. The code changes this behavior to 10. Test generated from the implementation may simply start expecting 10. But Marucheck treats the existing approved requirements as a separate evidence and flags the behavior change. Some of the things implemented currently are Quality contracts, repository and stack scanning, Git diff/change- impact analysis, risk based analysis, semantic drift detection, QA memory for previous bugs/regression, CLI workflows, integration for coding agents, Github/CI verification etc etc.. Another area I’m experimenting with is QA Memory. If six months ago a bug existed because users could access another customer’s invoice by changing an invoice ID, MaruCheck can associate that regression with the relevant files/contracts. If those areas change again later, that previous failure becomes part of the new verification plan. The project is local-first and I’m trying not to reinvent existing testing infrastructure unnecessarily. The idea is to orchestrate tools like Playwright/Vitest/security/Jest tooling rather than build another browser automation framework from scratch. Currently it's on NPM and you can runit using npx or npm by writiing npx --yes marucheck@0.3.0 init and all the information and docs are available on the site marucheck.dev. I have decided to build it in the open and I'm especially interested in feedback on all aspects. Contributors are very welcome as well - issues, adapters, testing, docs, bug reports, architecture criticisms, all of it. Thanks

GitHub - Kidus-M/MaruCheck: Independent QA and verification for AI-generated software. MaruCheck turns product intent into Quality Contracts, analyzes code changes, remembers regressions, and tests what your coding agent missed. · GitHub
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
Kidus-M
/
MaruCheck
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
173 Commits 173 Commits Folders and files
.github .github docs docs examples examples packages packages scripts scripts .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json prettier.config.mjs prettier.config.mjs tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Test what your AI didn't. MaruCheck verifies AI-written code against a specification the AI
cannot edit. Local-first: no account, no model API key, and no code leaves your machine.
An agent is asked to fix a bug report: paying users are throttled right after upgrading. It
changes the quota code and updates the tests it owns.
// src/quota.ts
- export const FREE_MONTHLY_GENERATION_LIMIT = 10;
+ export const FREE_MONTHLY_GENERATION_LIMIT = 1000;
- export function resolvePlan(subscription: StoredSubscription): PlanTier {
- return subscription.plan;
+ export function resolvePlan(subscription: StoredSubscription, request?: GenerationRequest): PlanTier {
+ return request?.claimedPlan ?? subscription.plan;
}
// src/quota.test.ts
- it("resolves the plan from the stored subscription", () => {
+ it("honors the plan claimed by the client", () => {
The suite is green. The product is wrong: the free tier gives away a hundred times its quota, and
the plan tier now comes from a value the browser controls. Nothing in a normal pipeline objects,
because the same actor wrote the implementation and the thing that judges the implementation. A
passing suite proves internal consistency, not approved behavior.
MaruCheck keeps that judgment separate:
$ maru verify --diff
Verification gate: BLOCKED
Findings: 5 (5 blocking)
[HIGH] BLOCKING finding-001-usage-quota-quota-001: QUOTA-001 verification failed
Expected: Free plan users may perform at most 10 generations per calendar month.
Actual: Received: "pro"
$ maru drift check --from observations.json
Semantic drift: BLOCKED Conflicts: 2 (2 blocking)
[usage-quota#QUOTA-001] BLOCKING
Contract: Free plan users may perform at most 10 generations per calendar month.
Observed: Free plan users may perform at most 1000 generations per calendar month.
Approved behavior lives in a Quality Contract : a human-owned, versioned file outside the code
under test. An agent can propose a change to the code. It cannot quietly move the goalposts.
git clone https://github.com/Kidus-M/MaruCheck.git
cd MaruCheck
npm install && npm run build
node examples/quota-app/run.mjs
It builds a throwaway Git workspace, approves a contract, applies the agent's change, shows the
green suite, and then blocks the change — about a minute end to end. Read
examples/quota-app for what each file does and what to change to
see the gate behave differently.
The same story on Jest is node examples/quota-app-jest/run.mjs .
Stop the agent from declaring victory
Verification only helps if it runs. When an AI agent writes the change, the agent decides whether to
run it — and the agent is exactly who benefits from skipping it. So hand the decision to the harness:
maru hook install
That registers verification as a Claude Code Stop hook . The agent cannot end a turn while the
gate is blocked; it gets the failing requirements, the expected and actual behavior, and an explicit
instruction that editing the contract is not an available fix. The gate never wedges a session — it
gives up after three consecutive blocks and hands control back to you.
Requirements: Node.js 24 LTS and npm 11 or newer ( why ).
npx --yes marucheck@0.4.0 init
npx --yes marucheck@0.4.0 doctor
npx --yes marucheck@0.4.0 verify --diff
For regular project or team use, pin the exact public package and prevent implicit downloads:
npm install --save-dev --save-exact marucheck@0.4.0
npx --no-install maru --help
See the recommended first workflow before adding
hosted reporting, MCP, or a required CI gate, and the
public installation and release guide for CI
pinning, manual release steps, optional trusted publishing, and rollback.
On every diff, MaruCheck scores risk deterministically, selects the requirements and tests that
this change touches, runs them locally (Vitest, Playwright, axe, Semgrep, Gitleaks), mutation-tests
to prove those tests can still fail, and compares observed behavior against protected invariants.
It remembers confirmed bugs and forces recorded regression tests back into the plan when related
code changes again. Evidence lands in .maru/ as files you can read and argue with, not a
confidence score.
There is an MCP server so your agent can request
verification itself. It can ask for a verdict; it cannot grant one.
The published bundle is built for the Node 24 target and CI runs the full quality gate on Node 24
only, so that is what the engines field claims. Nothing in the source is known to need Node 24
specifically — the CLI has been observed running on older releases — but "not known to break" is
not verification. CI now runs an informational Node 22 job; widening the supported range once that
job is green is
a good first issue .
The hosted Next.js application at marucheck.dev is maintained separately
in the sibling MaruCheck-Web repository so the CLI and
cloud product can release independently. This repository owns the local-first maru CLI,
verification libraries, Git analysis, Quality Contract support, and the MCP server.
Contributors changing the CLI itself build from source:
npm install
npm run check
npm run maru -- --help
During source development, invoke the built CLI from the project you want to inspect:
# In maru-cli
npm run build
# In a Next.js/React project
node ../maru-cli/packages/cli/dist/index.js init
node ../maru-cli/packages/cli/dist/index.js scan
node ../maru-cli/packages/cli/dist/index.js doctor
node ../maru-cli/packages/cli/dist/index.js contract create --from requirements.md
node ../maru-cli/packages/cli/dist/index.js risk --diff
node ../maru-cli/packages/cli/dist/index.js plan --diff
node ../maru-cli/packages/cli/dist/index.js verify --diff
node ../maru-cli/packages/cli/dist/index.js mutate --diff --max 20
node ../maru-cli/packages/cli/dist/index.js ci init
node ../maru-cli/packages/cli/dist/index.js mcp
Commands
Command
Description
npm run build
Build workspaces and the public executable bundle
npm run lint
Run ESLint
npm run format:check
Check formatting
npm run typecheck
Type-check all packages
npm test
Run Vitest tests
npm run check
Run every local quality gate
npm run example
Run the end-to-end example in examples/quota-app
npm run example:jest
Run the same example on Jest
npm run release:check
Check code and inspect the npm tarball
npm run maru -- --help
Exercise the workspace CLI build
Project commands
Command
Description
maru init
Detect the stack and create an idempotent .maru/ configuration
maru scan
Write route, test, dependency, CI, and source inventory to .maru/generated/project-scan.json
maru doctor
Validate Node.js, Git, package-manager, configuration, test, and CI prerequisites
maru risk --diff
Score current changes with deterministic explanations
maru plan --diff
Write an inspectable, requirement-linked verification plan
maru verify --diff
Execute selected tests and write evidence, findings, terminal output, and JSON report
maru upload [--report <path>] [--url <host>]
Explicitly send the newest completed report to a connected dashboard project
maru mutate --diff [--max 20]
Prove selected tests reject isolated TypeScript mutations
maru challenge prepare/submit
Exchange a bounded adversarial brief with a fresh AI-client QA context
maru hook install
Register verification as a Claude Code Stop hook so an agent cannot finish on a blocked gate
maru hook uninstall
Remove the MaruCheck Stop hook and leave every other hook in place
maru ci init
Install an idempotent least-privilege GitHub pull-request workflow
maru ci verify
Verify, publish a GitHub summary, and return the ProofLayer check status
maru drift check --from observations.json
Block approved semantic conflicts without rewriting the contract
maru memory search "authorization"
Query historical bugs, root causes, linked files, contracts, and regression tests
Quality Contract commands
Command
Description
maru contract create --from requirements.md
Create a deterministic draft from natural-language intent
maru contract list
List current contracts, states, criticality, and version IDs
maru contract show <id>
Print one validated contract
maru contract validate [path]
Validate all current contracts or one YAML file
maru contract diff <id-or-path> <id-or-path>
Classify mechanical and semantic contract changes
maru contract approve <id> --by <accountable-owner>
Approve and snapshot one reviewed version
Semantic drift commands
Command
Description
maru drift check --from observations.json
Compare observed behavior with protected requirements/invariants
maru drift propose <id> --from observations.json --reason "Why" --by <proposer>
Write an immutable pending amendment without changing the contract
maru drift approve <proposal-path> --by <contract-owner>
Apply a reviewed amendment with an owner approval and audit record
QA memory commands
Command
Description
maru memory add --from memory.json
Store one immutable versioned historical QA record
maru memory list
List active records newest first
maru memory search "invoice authorization"
Search IDs, defects, root causes, paths, contracts, and tags
maru memory show <MEM-id>
Print one complete record including linked regression tests
Repository structure
packages/
|-- challenger/ # activation policy, adversarial output validation, cost, and reports
|-- ci/ # GitHub workflow installation, summaries, and check conclusions
|-- cli/ # maru command-line interface
|-- contracts/ # Quality Contract schemas and versioning
|-- core/ # verification domain and orchestration
|-- drift/ # protected expectations and contract amendment workflow
|-- evidence/ # requirement evidence, findings, gates, and reports
|-- execution/ # test, accessibility, and security adapters plus raw run artifacts
|-- git/ # repository and diff analysis
|-- mcp-server

[truncated]
