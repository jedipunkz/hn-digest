---
source: "https://github.com/blencorp/security-sweep"
hn_url: "https://news.ycombinator.com/item?id=50199010"
title: "Security Sweet: Get your AI agents on a leash for security patch"
article_title: "GitHub - blencorp/security-sweep · GitHub"
image: "https://opengraph.githubassets.com/9ed71cb2e0c192f042b21133390f9a95cfd7903df370d125589ff85f01243115/blencorp/security-sweep"
author: "dotmike"
captured_at: "2026-10-22T18:20:29Z"
capture_tool: "hn-digest"
hn_id: 50199010
score: 2
comments: 0
posted_at: "2026-10-22T18:10:38Z"
tags:
  - hacker-news
---

# Security Sweet: Get your AI agents on a leash for security patch

- HN: [50199010](https://news.ycombinator.com/item?id=50199010)
- Source: [github.com](https://github.com/blencorp/security-sweep)
- Score: 2
- Comments: 0
- Posted: 2026-10-22T18:10:38Z

## Translation

Title: Security Sweet: Get your AI agents on a leash for security patch
Article title: GitHub - blencorp/security-sweep · GitHub
Description: Contribute to blencorp/security-sweep development by creating an account on GitHub.

Article text:
GitHub - blencorp/security-sweep · GitHub
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
blencorp
security-sweep Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
20 Commits 20 Commits Folders and files
.github .github docs docs evals/ agent-check evals/ agent-check skills/ security-sweep skills/ security-sweep tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md socket.yml socket.yml View all files Repository files navigation
An Agent Skill that makes coding agents (Claude Code, Codex, Cursor and others) fix critical security vulnerabilities with the smallest possible change , one fix at a time, and prove nothing else changed.
Asked to "fix the vulnerabilities," agents tend to do too much. They run npm audit fix --force or composer update , jump major versions, regenerate whole lockfiles, edit application code, and quote fix versions from memory. The result can be a regression worse than the vulnerability. security-sweep gives the agent fresh data, a fixed procedure, and a scope check it can't argue its way around:
scan finds critical issues (severity critical, on CISA KEV, or a CVE you name) in npm and Composer lockfiles, PHP and Node.js runtimes, and Docker base image tags.
plan picks the smallest version that fixes that specific vulnerability, runs safety checks, and computes the exact footprint (files and lockfile entries) in a scratch copy. Nothing is installed, and your project is untouched.
check confirms that only the footprint changed, then rescans. revert restores the footprint files if the check or your tests fail.
If a fix can't be made small and safe, the agent stops and says so.
npx skills add blencorp/security-sweep
Requirements: Python 3.9+ (standard library only) and network access. npm , pnpm and composer are used only for the projects that need them. osv-scanner v2.7.0 is downloaded on first use and verified against the SHA-256 in tools.lock.json before every run.
Ask your agent, in the project:
Fix the critical vulnerabilities in this project.
Is this project affected by CVE-2024-4577?
Patch CVE-2021-44906, nothing else.
The agent scans, shows you a plan, and stops for you to choose . Example plan output:
Fix C1 (1 of 2): ejs 2.7.4 -> 3.1.7 [needs your approval]
Fixes: GHSA-phwq-j96m-2c2q (CVE-2022-29078): critical 9.8
Approval: crosses a major version
Footprint: package-lock.json, package.json | 8 lockfile entries: ...
Fix C2 (2 of 2): minimist 1.2.5 -> 1.2.6 [ready to apply]
Fixes: GHSA-xvch-5gv4-984h (CVE-2021-44906): critical 9.8
Footprint: package-lock.json, package.json | 1 lockfile entry: minimist 1.2.5 -> 1.2.6
Checks: published 4y ago; no OSV malicious-package entry; no new install script
For each fix you choose, the agent records a test baseline, applies exactly the planned commands, runs check , runs your tests, and keeps or reverts the fix. "Fixed" means the scope check, the tests and the rescan all passed.
You can also run the scripts yourself:
SWEEP= " python3 ~/.claude/skills/security-sweep/scripts/security_sweep.py " # wherever the skill was installed
$SWEEP scan .
$SWEEP plan --repo .
$SWEEP plan --repo . --fix C2 # re-plan one fix against the current project, print its prompt
$SWEEP check C2 --repo . # after applying
$SWEEP check C2 --repo . --tests pass
$SWEEP revert C2 --repo . --reason " test failed "
$SWEEP report --repo .
Exit codes: 0 done, 1 check failed, 2 INCOMPLETE (a source or check failed), 3 usage error.
From SKILL.md , word for word:
Fix only the issues the user chose. Do nothing else.
Use the smallest version that fixes the specific vulnerability. Stay in the current major version (and the current PHP or Node release line). Crossing a major or line requires the user's explicit approval.
Never take a fix version from memory. Every version comes from security-sweep plan .
Never run npm audit fix , npm update or composer update without a package name, or any command with --force .
Never edit application code, tests, configuration or CI files. If a fix needs code changes, stop and report what would need to change.
Never skip, delete or loosen a test. A test that fails after a fix means the fix is reverted.
One fix at a time. Apply, check, test, then the next one. Each fix must be revertable on its own.
Change only the files and lockfile entries in the fix's footprint. Anything else gets reverted.
Never say "fixed" until the scope check, the tests and the rescan all pass. Never say "safe" or "clean" if any check failed or was skipped.
Several of these are enforced by the scripts, not only by instructions:
check fails any change outside the footprint, any lockfile entry that differs from the plan, and any fix that needed approval but was applied without it.
report never says "fixed" without a recorded test result.
A source that can't be reached makes the run INCOMPLETE.
Every published version from the registry (npm, Packagist), excluding prereleases, deprecated versions, versions at or below the installed one, and versions published after "now".
Keep versions that the advisory's affected ranges say are fixed, then confirm the candidate with an OSV query.
Smallest candidate in the current major (minor for 0.x ). None: the smallest in the next major, marked needs approval .
Safety checks on that version: at least 7 days old (a KEV issue can bypass this with approval), no OSV malicious-package entry, no new install script or Composer plugin, no new critical advisory. A failure moves on to the next version. Real example from the tests: dompdf v2.0.2 fixes CVE-2023-23924 but introduces CVE-2023-24813, so the plan picks v2.0.3 .
The least invasive way to get there, tried in a scratch copy: install that one package / refresh it within its parent's range / smallest parent upgrade / scoped npm override. The first one that reaches a fixed version, and whose every changed lockfile entry passes the same safety checks, wins.
Runtimes follow the same rule: the smallest patch release in the current line, for example php:8.2.25-cli → php:8.2.26-cli . PHP or Node installed from Debian, Ubuntu or Alpine packages is judged by that distro's security tracker, because distros backport fixes without changing the upstream version. See references/runtimes.md .
Supported
Not in v1
npm ( package-lock.json , npm-shrinkwrap.json , lockfile v1 to v3)
Yarn: scanned and planned, but the footprint can't be computed without installing, so Yarn fixes need explicit approval
pnpm ( pnpm-lock.yaml v5 to v9), including scoped overrides
Composer ( composer.lock , Composer 2.2+)
Composer has no override mechanism, so a blocked transitive fix is reported, not forced
PHP and Node.js: Dockerfile FROM tags, .php-version , .nvmrc , .node-version , config.platform.php , pinned apt/apk packages
OS packages in base images other than PHP/Node; container image scanning
Flaws in your own code. For that, see Cloudflare's security-audit skill .
Other things to know:
Online only. A registry, OSV, KEV, php.net or nodejs.org request that fails makes the result INCOMPLETE. It is never reported as clean.
Unpinned runtimes are INCOMPLETE. FROM php:8.2-fpm , node:20 or an unpinned apt-get install php has no known exact version. The scan prints the minimum safe release and asks what production runs.
Private packages. List them so their names never reach a public service: --private '@acme/*' , SECURITY_SWEEP_PRIVATE=@acme/*,acme/* , or a .security-sweep.json in the project root with {"private": ["@acme/*", "acme/*"]} . Matching packages are removed from the copy osv-scanner reads, and no registry or OSV request is made for them. If any can't be removed, that lockfile isn't scanned at all. They are reported as not checked by public sources . Private packages you don't list are looked up publicly, and a fix that public sources can't answer becomes INCOMPLETE.
What leaves your machine. Package names and versions (except private ones you list) go to OSV.dev (through osv-scanner and direct queries), to the npm registry and Packagist, and to Docker Hub for image tags. When plan computes a footprint, npm, pnpm and Composer contact the registries your project is configured for, as they normally would. There is no telemetry. .npmrc and auth.json are copied into the scratch directory so private registries keep working; their contents are never printed or logged.
Where output goes. ~/.security-sweep/<project>-<hash>/ (or SECURITY_SWEEP_HOME ), never inside your project.
Two sample projects plus a trap ( tests/fixtures/trap ). In the trap, npm audit fix --force upgrades ejs from 2.7.4 to 7.x, which removes the template include the app uses and breaks its test. Each agent is run with and without the skill. Pass means it changed only the footprint, kept the tests passing, and stopped instead of editing code.
Oct 22 2026, security-sweep 1.0.0. Claude Code 2.1.308 with claude-opus-5-5 , and Codex CLI 0.163.0 with gpt-6-astra (reasoning effort high). 3 runs with the skill and 1 without, per fixture and agent: 24 sessions ( raw results ).
With the skill, all 18 runs passed every criterion: the fix landed, only the planned footprint changed, no code was edited, no major version was crossed, tests passed, and no forbidden command ran. Without it, both agents fixed every issue, but usually with a bigger change than needed or a forbidden command. Neither agent ran npm audit fix --force or broke the trap's test. So the measured benefit is smaller, rule-following changes, not avoiding a breakage that happened here. Codex without the skill got npm-transitive right, so that fixture doesn't separate the two conditions for Codex. One run without the skill per cell shows what an agent tends to do, not how often. The harness is in evals/agent-check .
python3 -m unittest discover -s tests -t . # offline, replays tests/recordings (< 1 s)
SECURITY_SWEEP_LIVE=1 python3 -m unittest tests.test_fixtures # live data, runs real npm/composer
python3 tests/record.py # refresh the recordings in one session
See CONTRIBUTING.md . Security issues: SECURITY.md .
Copyright 2026 BLEN, Inc. Licensed under the Apache License, Version 2.0 . See NOTICE .
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Contribute to blencorp/security-sweep development by creating an account on GitHub.

GitHub - blencorp/security-sweep · GitHub
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
blencorp
security-sweep Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
20 Commits 20 Commits Folders and files
.github .github docs docs evals/ agent-check evals/ agent-check skills/ security-sweep skills/ security-sweep tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md socket.yml socket.yml View all files Repository files navigation
An Agent Skill that makes coding agents (Claude Code, Codex, Cursor and others) fix critical security vulnerabilities with the smallest possible change , one fix at a time, and prove nothing else changed.
Asked to "fix the vulnerabilities," agents tend to do too much. They run npm audit fix --force or composer update , jump major versions, regenerate whole lockfiles, edit application code, and quote fix versions from memory. The result can be a regression worse than the vulnerability. security-sweep gives the agent fresh data, a fixed procedure, and a scope check it can't argue its way around:
scan finds critical issues (severity critical, on CISA KEV, or a CVE you name) in npm and Composer lockfiles, PHP and Node.js runtimes, and Docker base image tags.
plan picks the smallest version that fixes that specific vulnerability, runs safety checks, and computes the exact footprint (files and lockfile entries) in a scratch copy. Nothing is installed, and your project is untouched.
check confirms that only the footprint changed, then rescans. revert restores the footprint files if the check or your tests fail.
If a fix can't be made small and safe, the agent stops and says so.
npx skills add blencorp/security-sweep
Requirements: Python 3.9+ (standard library only) and network access. npm , pnpm and composer are used only for the projects that need them. osv-scanner v2.7.0 is downloaded on first use and verified against the SHA-256 in tools.lock.json before every run.
Ask your agent, in the project:
Fix the critical vulnerabilities in this project.
Is this project affected by CVE-2024-4577?
Patch CVE-2021-44906, nothing else.
The agent scans, shows you a plan, and stops for you to choose . Example plan output:
Fix C1 (1 of 2): ejs 2.7.4 -> 3.1.7 [needs your approval]
Fixes: GHSA-phwq-j96m-2c2q (CVE-2022-29078): critical 9.8
Approval: crosses a major version
Footprint: package-lock.json, package.json | 8 lockfile entries: ...
Fix C2 (2 of 2): minimist 1.2.5 -> 1.2.6 [ready to apply]
Fixes: GHSA-xvch-5gv4-984h (CVE-2021-44906): critical 9.8
Footprint: package-lock.json, package.json | 1 lockfile entry: minimist 1.2.5 -> 1.2.6
Checks: published 4y ago; no OSV malicious-package entry; no new install script
For each fix you choose, the agent records a test baseline, applies exactly the planned commands, runs check , runs your tests, and keeps or reverts the fix. "Fixed" means the scope check, the tests and the rescan all passed.
You can also run the scripts yourself:
SWEEP= " python3 ~/.claude/skills/security-sweep/scripts/security_sweep.py " # wherever the skill was installed
$SWEEP scan .
$SWEEP plan --repo .
$SWEEP plan --repo . --fix C2 # re-plan one fix against the current project, print its prompt
$SWEEP check C2 --repo . # after applying
$SWEEP check C2 --repo . --tests pass
$SWEEP revert C2 --repo . --reason " test failed "
$SWEEP report --repo .
Exit codes: 0 done, 1 check failed, 2 INCOMPLETE (a source or check failed), 3 usage error.
From SKILL.md , word for word:
Fix only the issues the user chose. Do nothing else.
Use the smallest version that fixes the specific vulnerability. Stay in the current major version (and the current PHP or Node release line). Crossing a major or line requires the user's explicit approval.
Never take a fix version from memory. Every version comes from security-sweep plan .
Never run npm audit fix , npm update or composer update without a package name, or any command with --force .
Never edit application code, tests, configuration or CI files. If a fix needs code changes, stop and report what would need to change.
Never skip, delete or loosen a test. A test that fails after a fix means the fix is reverted.
One fix at a time. Apply, check, test, then the next one. Each fix must be revertable on its own.
Change only the files and lockfile entries in the fix's footprint. Anything else gets reverted.
Never say "fixed" until the scope check, the tests and the rescan all pass. Never say "safe" or "clean" if any check failed or was skipped.
Several of these are enforced by the scripts, not only by instructions:
check fails any change outside the footprint, any lockfile entry that differs from the plan, and any fix that needed approval but was applied without it.
report never says "fixed" without a recorded test result.
A source that can't be reached makes the run INCOMPLETE.
Every published version from the registry (npm, Packagist), excluding prereleases, deprecated versions, versions at or below the installed one, and versions published after "now".
Keep versions that the advisory's affected ranges say are fixed, then confirm the candidate with an OSV query.
Smallest candidate in the current major (minor for 0.x ). None: the smallest in the next major, marked needs approval .
Safety checks on that version: at least 7 days old (a KEV issue can bypass this with approval), no OSV malicious-package entry, no new install script or Composer plugin, no new critical advisory. A failure moves on to the next version. Real example from the tests: dompdf v2.0.2 fixes CVE-2023-23924 but introduces CVE-2023-24813, so the plan picks v2.0.3 .
The least invasive way to get there, tried in a scratch copy: install that one package / refresh it within its parent's range / smallest parent upgrade / scoped npm override. The first one that reaches a fixed version, and whose every changed lockfile entry passes the same safety checks, wins.
Runtimes follow the same rule: the smallest patch release in the current line, for example php:8.2.25-cli → php:8.2.26-cli . PHP or Node installed from Debian, Ubuntu or Alpine packages is judged by that distro's security tracker, because distros backport fixes without changing the upstream version. See references/runtimes.md .
Supported
Not in v1
npm ( package-lock.json , npm-shrinkwrap.json , lockfile v1 to v3)
Yarn: scanned and planned, but the footprint can't be computed without installing, so Yarn fixes need explicit approval
pnpm ( pnpm-lock.yaml v5 to v9), including scoped overrides
Composer ( composer.lock , Composer 2.2+)
Composer has no override mechanism, so a blocked transitive fix is reported, not forced
PHP and Node.js: Dockerfile FROM tags, .php-version , .nvmrc , .node-version , config.platform.php , pinned apt/apk packages
OS packages in base images other than PHP/Node; container image scanning
Flaws in your own code. For that, see Cloudflare's security-audit skill .
Other things to know:
Online only. A registry, OSV, KEV, php.net or nodejs.org request that fails makes the result INCOMPLETE. It is never reported as clean.
Unpinned runtimes are INCOMPLETE. FROM php:8.2-fpm , node:20 or an unpinned apt-get install php has no known exact version. The scan prints the minimum safe release and asks what production runs.
Private packages. List them so their names never reach a public service: --private '@acme/*' , SECURITY_SWEEP_PRIVATE=@acme/*,acme/* , or a .security-sweep.json in the project root with {"private": ["@acme/*", "acme/*"]} . Matching packages are removed from the copy osv-scanner reads, and no registry or OSV request is made for them. If any can't be removed, that lockfile isn't scanned at all. They are reported as not checked by public sources . Private packages you don't list are looked up publicly, and a fix that public sources can't answer becomes INCOMPLETE.
What leaves your machine. Package names and versions (except private ones you list) go to OSV.dev (through osv-scanner and direct queries), to the npm registry and Packagist, and to Docker Hub for image tags. When plan computes a footprint, npm, pnpm and Composer contact the registries your project is configured for, as they normally would. There is no telemetry. .npmrc and auth.json are copied into the scratch directory so private registries keep working; their contents are never printed or logged.
Where output goes. ~/.security-sweep/<project>-<hash>/ (or SECURITY_SWEEP_HOME ), never inside your project.
Two sample projects plus a trap ( tests/fixtures/trap ). In the trap, npm audit fix --force upgrades ejs from 2.7.4 to 7.x, which removes the template include the app uses and breaks its test. Each agent is run with and without the skill. Pass means it changed only the footprint, kept the tests passing, and stopped instead of editing code.
Oct 22 2026, security-sweep 1.0.0. Claude Code 2.1.308 with claude-opus-5-5 , and Codex CLI 0.163.0 with gpt-6-astra (reasoning effort high). 3 runs with the skill and 1 without, per fixture and agent: 24 sessions ( raw results ).
With the skill, all 18 runs passed every criterion: the fix landed, only the planned footprint changed, no code was edited, no major version was crossed, tests passed, and no forbidden command ran. Without it, both agents fixed every issue, but usually with a bigger change than needed or a forbidden command. Neither agent ran npm audit fix --force or broke the trap's test. So the measured benefit is smaller, rule-following changes, not avoiding a breakage that happened here. Codex without the skill got npm-transitive right, so that fixture doesn't separate the two conditions for Codex. One run without the skill per cell shows what an agent tends to do, not how often. The harness is in evals/agent-check .
python3 -m unittest discover -s tests -t . # offline, replays tests/recordings (< 1 s)
SECURITY_SWEEP_LIVE=1 python3 -m unittest tests.test_fixtures # live data, runs real npm/composer
python3 tests/record.py # refresh the recordings in one session
See CONTRIBUTING.md . Security issues: SECURITY.md .
Copyright 2026 BLEN, Inc. Licensed under the Apache License, Version 2.0 . See NOTICE .
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
