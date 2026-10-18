---
source: "https://github.com/isAlmogK/trustbut"
hn_url: "https://news.ycombinator.com/item?id=50144227"
title: "Trustbut checks if an AI agent's new tests fail without its change"
article_title: "GitHub - isAlmogK/trustbut: Trust, but verify, for AI coding agents. A local CLI that reruns your tests and prints a receipt of what your agent actually did: which new tests prove the change, which pass without it, and what got skipped, deleted or silenced. Works with Claude Code, Codex, Cursor, Gem\n[truncated]"
image: "https://opengraph.githubassets.com/df7720e61ae13e780590a6de6ac952420820171b32c213508bbb03d0c0aea28d/isAlmogK/trustbut"
author: "almog_io"
captured_at: "2026-10-18T19:15:10Z"
capture_tool: "hn-digest"
hn_id: 50144227
score: 1
comments: 0
posted_at: "2026-10-18T19:07:31Z"
tags:
  - hacker-news
---

# Trustbut checks if an AI agent's new tests fail without its change

- HN: [50144227](https://news.ycombinator.com/item?id=50144227)
- Source: [github.com](https://github.com/isAlmogK/trustbut)
- Score: 1
- Comments: 0
- Posted: 2026-10-18T19:07:31Z

## Translation

Title: Trustbut checks if an AI agent's new tests fail without its change
Article title: GitHub - isAlmogK/trustbut: Trust, but verify, for AI coding agents. A local CLI that reruns your tests and prints a receipt of what your agent actually did: which new tests prove the change, which pass without it, and what got skipped, deleted or silenced. Works with Claude Code, Codex, Cursor, Gem
[truncated]
Description: Trust, but verify, for AI coding agents. A local CLI that reruns your tests and prints a receipt of what your agent actually did: which new tests prove the change, which pass without it, and what got skipped, deleted or silenced. Works with Claude Code, Codex, Cursor, Gemini and OpenCode. No LLM, no
[truncated]

Article text:
GitHub - isAlmogK/trustbut: Trust, but verify, for AI coding agents. A local CLI that reruns your tests and prints a receipt of what your agent actually did: which new tests prove the change, which pass without it, and what got skipped, deleted or silenced. Works with Claude Code, Codex, Cursor, Gemini and OpenCode. No LLM, no server. · GitHub
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
isAlmogK
Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Trust, but verify, for AI coding agents. A local CLI that reruns your tests and prints a receipt of what your agent actually did: which new tests prove the change, which pass without it, and what got skipped, deleted or silenced. Works with Claude Code, Codex, Cursor, Gemini and OpenCode. No LLM, no server.
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
205 Commits 205 Commits Folders and files
.claude-plugin .claude-plugin .github .github docs docs integrations/ claude-code integrations/ claude-code scripts scripts spec spec src src test test .gitignore .gitignore .trustbut.json .trustbut.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md action.yml action.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Trust, but verify, for AI coding agents.
Your agent says it's done and the tests pass. Does anything actually prove the change?
🧾 Verified by trustbut: pull requests to this repo carry its own trustbut receipt.
Languages · Why · What we've measured · Install · Usage · Use it with your agent · Config · What it checks · Receipt format · CI · How it works · FAQ · Roadmap · Sponsor · About me
trustbut is a local command-line tool that produces a receipt for a code change, whether an AI agent or a person wrote it. It reruns your tests in a clean checkout, runs every new or changed test against the code from before the change to see whether it fails there (if it passes either way, it doesn't test the change), and flags skipped tests, removed assertions and silenced type checks. Then it gives you a PASS / WARN / FAIL verdict. Wired into your agent, it can also stop the agent from finishing while the receipt is FAIL, so it keeps fixing until the evidence holds up ( details ).
It works with any agent (Claude Code, Codex, Cursor, Gemini CLI, OpenCode, Copilot) and 11 languages (JavaScript/TypeScript, Python, Go, Rust, Java, Kotlin, C#, Ruby, PHP, Dart/Flutter and Swift) because it only needs git and your test runner. No LLM, no server, no account. Your code never leaves your machine.
$ npx trustbut verify
## 🧾 Receipt: WARN
`3f2a91c` → `8be04d1` · runner: vitest
**Tests (clean checkout):** 41 passed · 0 failed · 1 skipped
**Proof:** 1 new/changed test proves the change ✅
**New/changed tests:** 1 prove the change · 1 pass without it
### Findings
- ⚠ **Test skipped** · `test/auth.test.ts:88`: it.skip('handles clock skew', () => {
- ⚠ **cast to any** · `src/auth/session.ts:41`: const claims = decoded as any;
### Review these first
1. `src/auth/session.ts`: auth/permissions, has findings
2. `package.json`: dependencies
When the suite fails at head, the receipt also lists the failing tests. If package.json declares dependencies that aren't installed in your node_modules (a stale install), a deps-not-installed warning names them, because both runs link and use your installed dependencies. See the finding ids in the spec's finding catalog .
The output is Markdown, so the same receipt reads cleanly in your terminal and renders properly in a pull request or a CI job summary.
Language
Test frameworks
Tamper and escape checks
Verified in CI against the real toolchain
JavaScript / TypeScript
Vitest, Jest (incl. Expo / React Native), Mocha, node:test
✅
✅
Python
pytest
✅
✅
Go
go test
✅
✅
Rust
cargo-nextest, cargo test
✅
✅
Java
Maven, Gradle
✅
✅
Kotlin
Maven, Gradle
✅
◐ same runners as Java; Kotlin end-to-end test coming
C# / .NET
dotnet test (xUnit verified; NUnit and MSTest via the same TRX results)
✅
✅
Ruby
RSpec
✅
✅
PHP
PHPUnit
✅
✅
Dart / Flutter
dart test , flutter test
✅
✅
Swift
Swift Package Manager (XCTest, Swift Testing)
✅
✅
Anything else
any command that writes JUnit XML
—
✅
trustbut detects the right runner automatically, including in monorepo subfolders. Setup steps (such as starting a database emulator), extra arguments and environment variables go in a .trustbut.json . Detection rules and the exact commands are in Supported languages and test runners .
AI agents now write a large share of the code I ship. Writing it is no longer the slow part. Checking it is.
When an agent says "done, all tests pass", that's the agent reporting on its own work. Reading the diff doesn't show what actually ran, or whether the new tests would fail if the change were wrong. I wanted a check from outside the agent that answers those two questions with evidence, not with another model's opinion.
In the 2026 Stack Overflow survey , 48% of developers trust AI output when they can easily validate it, but only 16% trust it for most tasks. The gap between those numbers is the problem trustbut works on.
In 2025, 66% of developers named "almost right, but not quite" as their top frustration with AI tools ( InfoWorld on the 2025 survey ).
Armin Ronacher put it well in The Final Bottleneck : generating code is now cheaper than reviewing it, and the human stays accountable.
The labs have also reported their models gaming tests in agentic coding:
Anthropic's Claude 3.7 Sonnet system card reports the model special-casing tests, and sometimes modifying them, to make them pass in agentic coding environments like Claude Code. The Claude 4 system card reports this much reduced, but still present.
OpenAI's research on monitoring reasoning models caught a model exiting early or skipping unit tests to look green.
METR found frontier models patching the grading and scoring code of their evaluations, and ImpossibleBench measures how often models exploit tests on tasks that can't be solved honestly.
That's why trustbut also flags skipped tests, removed assertions and silenced checks. But in my own measurement of 201 real pull requests, merged agent work rarely showed these behaviours (see What we've measured ). The question that holds up is the plainer one: does anything prove this change?
AI code reviewers read the diff and comment, but they don't rerun anything. A vendor checking its own agent is a weak audit, and most teams use more than one agent. So trustbut is independent, deterministic and local : it doesn't ask an AI whether the code looks right. It runs things and reports what it found.
I'm measuring whether trustbut is useful, in the open, in trustbut-index . It runs trustbut on real merged pull requests from established open-source projects: PRs written by agents (Copilot, Codex, Claude Code, Cursor, Devin, Jules) and, as a control, PRs written by people in the same projects. These are pilot results, not the final study.
trustbut 0.7.0 on 201 pull requests:
¹ 18 of the 109 runs had a failing test with no result from before the change, so this count is a lower bound.
Agents didn't break suites or fake their tests in these PRs. Their new tests were at least as likely as people's to test the change. Merged PRs have already been through CI and review, though, so this says little about what an agent does before it says "done". That's where the agent hook works, and it isn't measured yet.
Most findings are true but don't need action. In the 0.6.0 pilot, 71 findings were labeled: 5 were worth a reviewer acting on (7%), 57 were true but harmless, and 9 were wrong. All 5 useful ones said that nothing proved the change. 0.7.0 fixed the causes of the 9 wrong findings and made that one its headline, change-not-proven . The labels are drafts by Claude that no person has confirmed yet, and the labeler could see the code, so they weren't fully blind.
trustbut ran on about half of the PRs. Of the 92 it couldn't run, 78 failed to install the project's dependencies or toolchain in the study's runner, 13 had no supported test runner, and 1 result is missing.
Not measured yet: whether change-not-proven findings are worth acting on, and whether the agent hook leads to better agent work.
On my own repos ( DogBase ), receipts found a fix commit with no test, 7 new tests that passed with or without their change, two new eslint-disable lines, and a dependency missing from the local install, which became the deps-not-installed warning.
Have a run of your own? Open an issue with what the receipt found, or missed.
Requirements: Node.js 22 or newer, git, and a repository with tests in one of the supported languages and test runners , or any runner that can write JUnit XML. The toolchain of your project (Go, Cargo, Maven, .NET, Composer, …) must be installed, as it is when you run the tests yourself.
Run without installing (recommended):
npx trustbut verify
Install globally:
npm i -g trustbut
trustbut verify
From source:
git clone https://github.com/isAlmogK/trustbut.git
cd trustbut
npm ci
npm link
To have your agent verify its own work after every turn, see Use it with your agent .
Run it in your repository after an agent finishes:
trustbut verify # the committed work on your branch
trustbut verify --working # uncommitted work too (tracked edits plus untracked files)
By default trustbut compares your branch with the point where it split off main (or master ), the same way a GitHub pull request shows its diff.
# Compare against a specific branch or commit
trustbut verify --base develop
trustbut verify --base HEAD~3
# Verify uncommitted work (the agent hasn't committed yet)
trustbut verify --working
# Faster: rerun only the new and changed test files at head
trustbut verify --quick
# The tests live in a subfolder, or the config lives outside the repo
trustbut verify --root services/api
trustbut verify --config ~ /configs/myapp.trustbut.json
# Gate a script or CI job on warnings as well as failures
trustbut verify --strict
# Machine-readable output (an agent receipt v1)
trustbut verify --json > receipt.json
Options
Option
Description
Default
-b, --base <ref>
Base branch or commit, compared from its merge-base with head
origin/main , main , origin/master or master
--head <ref>
The commit to verify (not with --working )
HEAD
--working
Verify uncommitted work: tracked edits plus untracked, non-ignored files, snapshotted without touching your index
off
--runner <name>
Force a test runner, any name from the table below

[truncated]

## Original Extract

Trust, but verify, for AI coding agents. A local CLI that reruns your tests and prints a receipt of what your agent actually did: which new tests prove the change, which pass without it, and what got skipped, deleted or silenced. Works with Claude Code, Codex, Cursor, Gemini and OpenCode. No LLM, no
[truncated]

GitHub - isAlmogK/trustbut: Trust, but verify, for AI coding agents. A local CLI that reruns your tests and prints a receipt of what your agent actually did: which new tests prove the change, which pass without it, and what got skipped, deleted or silenced. Works with Claude Code, Codex, Cursor, Gemini and OpenCode. No LLM, no server. · GitHub
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
isAlmogK
Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Trust, but verify, for AI coding agents. A local CLI that reruns your tests and prints a receipt of what your agent actually did: which new tests prove the change, which pass without it, and what got skipped, deleted or silenced. Works with Claude Code, Codex, Cursor, Gemini and OpenCode. No LLM, no server.
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
205 Commits 205 Commits Folders and files
.claude-plugin .claude-plugin .github .github docs docs integrations/ claude-code integrations/ claude-code scripts scripts spec spec src src test test .gitignore .gitignore .trustbut.json .trustbut.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md action.yml action.yml package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsup.config.ts tsup.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Trust, but verify, for AI coding agents.
Your agent says it's done and the tests pass. Does anything actually prove the change?
🧾 Verified by trustbut: pull requests to this repo carry its own trustbut receipt.
Languages · Why · What we've measured · Install · Usage · Use it with your agent · Config · What it checks · Receipt format · CI · How it works · FAQ · Roadmap · Sponsor · About me
trustbut is a local command-line tool that produces a receipt for a code change, whether an AI agent or a person wrote it. It reruns your tests in a clean checkout, runs every new or changed test against the code from before the change to see whether it fails there (if it passes either way, it doesn't test the change), and flags skipped tests, removed assertions and silenced type checks. Then it gives you a PASS / WARN / FAIL verdict. Wired into your agent, it can also stop the agent from finishing while the receipt is FAIL, so it keeps fixing until the evidence holds up ( details ).
It works with any agent (Claude Code, Codex, Cursor, Gemini CLI, OpenCode, Copilot) and 11 languages (JavaScript/TypeScript, Python, Go, Rust, Java, Kotlin, C#, Ruby, PHP, Dart/Flutter and Swift) because it only needs git and your test runner. No LLM, no server, no account. Your code never leaves your machine.
$ npx trustbut verify
## 🧾 Receipt: WARN
`3f2a91c` → `8be04d1` · runner: vitest
**Tests (clean checkout):** 41 passed · 0 failed · 1 skipped
**Proof:** 1 new/changed test proves the change ✅
**New/changed tests:** 1 prove the change · 1 pass without it
### Findings
- ⚠ **Test skipped** · `test/auth.test.ts:88`: it.skip('handles clock skew', () => {
- ⚠ **cast to any** · `src/auth/session.ts:41`: const claims = decoded as any;
### Review these first
1. `src/auth/session.ts`: auth/permissions, has findings
2. `package.json`: dependencies
When the suite fails at head, the receipt also lists the failing tests. If package.json declares dependencies that aren't installed in your node_modules (a stale install), a deps-not-installed warning names them, because both runs link and use your installed dependencies. See the finding ids in the spec's finding catalog .
The output is Markdown, so the same receipt reads cleanly in your terminal and renders properly in a pull request or a CI job summary.
Language
Test frameworks
Tamper and escape checks
Verified in CI against the real toolchain
JavaScript / TypeScript
Vitest, Jest (incl. Expo / React Native), Mocha, node:test
✅
✅
Python
pytest
✅
✅
Go
go test
✅
✅
Rust
cargo-nextest, cargo test
✅
✅
Java
Maven, Gradle
✅
✅
Kotlin
Maven, Gradle
✅
◐ same runners as Java; Kotlin end-to-end test coming
C# / .NET
dotnet test (xUnit verified; NUnit and MSTest via the same TRX results)
✅
✅
Ruby
RSpec
✅
✅
PHP
PHPUnit
✅
✅
Dart / Flutter
dart test , flutter test
✅
✅
Swift
Swift Package Manager (XCTest, Swift Testing)
✅
✅
Anything else
any command that writes JUnit XML
—
✅
trustbut detects the right runner automatically, including in monorepo subfolders. Setup steps (such as starting a database emulator), extra arguments and environment variables go in a .trustbut.json . Detection rules and the exact commands are in Supported languages and test runners .
AI agents now write a large share of the code I ship. Writing it is no longer the slow part. Checking it is.
When an agent says "done, all tests pass", that's the agent reporting on its own work. Reading the diff doesn't show what actually ran, or whether the new tests would fail if the change were wrong. I wanted a check from outside the agent that answers those two questions with evidence, not with another model's opinion.
In the 2026 Stack Overflow survey , 48% of developers trust AI output when they can easily validate it, but only 16% trust it for most tasks. The gap between those numbers is the problem trustbut works on.
In 2025, 66% of developers named "almost right, but not quite" as their top frustration with AI tools ( InfoWorld on the 2025 survey ).
Armin Ronacher put it well in The Final Bottleneck : generating code is now cheaper than reviewing it, and the human stays accountable.
The labs have also reported their models gaming tests in agentic coding:
Anthropic's Claude 3.7 Sonnet system card reports the model special-casing tests, and sometimes modifying them, to make them pass in agentic coding environments like Claude Code. The Claude 4 system card reports this much reduced, but still present.
OpenAI's research on monitoring reasoning models caught a model exiting early or skipping unit tests to look green.
METR found frontier models patching the grading and scoring code of their evaluations, and ImpossibleBench measures how often models exploit tests on tasks that can't be solved honestly.
That's why trustbut also flags skipped tests, removed assertions and silenced checks. But in my own measurement of 201 real pull requests, merged agent work rarely showed these behaviours (see What we've measured ). The question that holds up is the plainer one: does anything prove this change?
AI code reviewers read the diff and comment, but they don't rerun anything. A vendor checking its own agent is a weak audit, and most teams use more than one agent. So trustbut is independent, deterministic and local : it doesn't ask an AI whether the code looks right. It runs things and reports what it found.
I'm measuring whether trustbut is useful, in the open, in trustbut-index . It runs trustbut on real merged pull requests from established open-source projects: PRs written by agents (Copilot, Codex, Claude Code, Cursor, Devin, Jules) and, as a control, PRs written by people in the same projects. These are pilot results, not the final study.
trustbut 0.7.0 on 201 pull requests:
¹ 18 of the 109 runs had a failing test with no result from before the change, so this count is a lower bound.
Agents didn't break suites or fake their tests in these PRs. Their new tests were at least as likely as people's to test the change. Merged PRs have already been through CI and review, though, so this says little about what an agent does before it says "done". That's where the agent hook works, and it isn't measured yet.
Most findings are true but don't need action. In the 0.6.0 pilot, 71 findings were labeled: 5 were worth a reviewer acting on (7%), 57 were true but harmless, and 9 were wrong. All 5 useful ones said that nothing proved the change. 0.7.0 fixed the causes of the 9 wrong findings and made that one its headline, change-not-proven . The labels are drafts by Claude that no person has confirmed yet, and the labeler could see the code, so they weren't fully blind.
trustbut ran on about half of the PRs. Of the 92 it couldn't run, 78 failed to install the project's dependencies or toolchain in the study's runner, 13 had no supported test runner, and 1 result is missing.
Not measured yet: whether change-not-proven findings are worth acting on, and whether the agent hook leads to better agent work.
On my own repos ( DogBase ), receipts found a fix commit with no test, 7 new tests that passed with or without their change, two new eslint-disable lines, and a dependency missing from the local install, which became the deps-not-installed warning.
Have a run of your own? Open an issue with what the receipt found, or missed.
Requirements: Node.js 22 or newer, git, and a repository with tests in one of the supported languages and test runners , or any runner that can write JUnit XML. The toolchain of your project (Go, Cargo, Maven, .NET, Composer, …) must be installed, as it is when you run the tests yourself.
Run without installing (recommended):
npx trustbut verify
Install globally:
npm i -g trustbut
trustbut verify
From source:
git clone https://github.com/isAlmogK/trustbut.git
cd trustbut
npm ci
npm link
To have your agent verify its own work after every turn, see Use it with your agent .
Run it in your repository after an agent finishes:
trustbut verify # the committed work on your branch
trustbut verify --working # uncommitted work too (tracked edits plus untracked files)
By default trustbut compares your branch with the point where it split off main (or master ), the same way a GitHub pull request shows its diff.
# Compare against a specific branch or commit
trustbut verify --base develop
trustbut verify --base HEAD~3
# Verify uncommitted work (the agent hasn't committed yet)
trustbut verify --working
# Faster: rerun only the new and changed test files at head
trustbut verify --quick
# The tests live in a subfolder, or the config lives outside the repo
trustbut verify --root services/api
trustbut verify --config ~ /configs/myapp.trustbut.json
# Gate a script or CI job on warnings as well as failures
trustbut verify --strict
# Machine-readable output (an agent receipt v1)
trustbut verify --json > receipt.json
Options
Option
Description
Default
-b, --base <ref>
Base branch or commit, compared from its merge-base with head
origin/main , main , origin/master or master
--head <ref>
The commit to verify (not with --working )
HEAD
--working
Verify uncommitted work: tracked edits plus untracked, non-ignored files, snapshotted without touching your index
off
--runner <name>
Force a test runner, any name from the table below

[truncated]
