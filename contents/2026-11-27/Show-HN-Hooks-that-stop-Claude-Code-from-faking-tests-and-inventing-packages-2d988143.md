---
source: "https://github.com/John99101/claude-code-project-template"
hn_url: "https://news.ycombinator.com/item?id=50669015"
title: "Show HN: Hooks that stop Claude Code from faking tests and inventing packages"
article_title: "GitHub - John99101/claude-code-project-template: Make Claude Code know your project in 5 minutes. One command detects your stack and writes CLAUDE.md, rules, hooks, skills, agents and MCP config, plus a safety guard, a done gate that runs your tests, and /pr, /start-issue and /handoff commands. Plug\n[truncated]"
image: "https://opengraph.githubassets.com/1e013c509e13827b0b62ceca8da8fb38c2daa75ad28a1238d2b61055a115668e/John99101/claude-code-project-template"
author: "john6601"
captured_at: "2026-11-27T17:24:16Z"
capture_tool: "hn-digest"
hn_id: 50669015
score: 1
comments: 0
posted_at: "2026-11-27T17:01:08Z"
tags:
  - hacker-news
---

# Show HN: Hooks that stop Claude Code from faking tests and inventing packages

- HN: [50669015](https://news.ycombinator.com/item?id=50669015)
- Source: [github.com](https://github.com/John99101/claude-code-project-template)
- Score: 1
- Comments: 0
- Posted: 2026-11-27T17:01:08Z

## Translation

Title: Show HN: Hooks that stop Claude Code from faking tests and inventing packages
Article title: GitHub - John99101/claude-code-project-template: Make Claude Code know your project in 5 minutes. One command detects your stack and writes CLAUDE.md, rules, hooks, skills, agents and MCP config, plus a safety guard, a done gate that runs your tests, and /pr, /start-issue and /handoff commands. Plug
[truncated]
Description: Make Claude Code know your project in 5 minutes. One command detects your stack and writes CLAUDE.md, rules, hooks, skills, agents and MCP config, plus a safety guard, a done gate that runs your tests, and /pr, /start-issue and /handoff commands. Plugin or installer; Linux, macOS, Windows. - John991
[truncated]

Article text:
GitHub - John99101/claude-code-project-template: Make Claude Code know your project in 5 minutes. One command detects your stack and writes CLAUDE.md, rules, hooks, skills, agents and MCP config, plus a safety guard, a done gate that runs your tests, and /pr, /start-issue and /handoff commands. Plugin or installer; Linux, macOS, Windows. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
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
John99101
claude-code-project-template Public template Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Make Claude Code know your project in 5 minutes. One command detects your stack and writes CLAUDE.md, rules, hooks, skills, agents and MCP config, plus a safety guard, a done gate that runs your tests, and /pr, /start-issue and /handoff commands. Plugin or installer; Linux, macOS, Windows.
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
49 Commits 49 Commits Folders and files
.claude-plugin .claude-plugin .github .github agents agents docs docs hooks hooks rule-packs rule-packs rules rules scripts scripts skills skills templates templates tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md INSTALL.md INSTALL.md LICENSE LICENSE README.md README.md VERSION VERSION install.ps1 install.ps1 install.py install.py install.sh install.sh View all files Repository files navigation
Make Claude Code actually know your project, in five minutes.
One command, and Claude reads your repo, asks a few multiple-choice questions, and writes the whole setup for you:
CLAUDE.md , rules, permissions, hooks, skills, agents and MCP servers, all matched to your real stack.
Illustrative session. Every message from the kit in it is the hooks' real output.
/plugin install cct --marketplace John99101/claude-code-project-template
/cct:setup
Quick start · What you get · A day with the kit · FAQ · Full guide
Out of the box, Claude Code doesn't know how your tests run, which folders are generated, what your team's commit style is, or that it must never touch .env . So it guesses, and you correct it over and over.
Claude Code can learn all of this through CLAUDE.md , rules, settings, hooks, skills, agents and MCP. Setting them up by hand takes hours per project, is easy to get wrong, and goes stale. This kit does it for you, and keeps it current.
Plugin (easiest). In Claude Code:
/plugin install cct --marketplace John99101/claude-code-project-template
Then run /cct:setup in any project. Updates arrive through /plugin . On older Claude Code versions, add the marketplace first with /plugin marketplace add John99101/claude-code-project-template , then /plugin install cct@cct-kit .
Ask Claude. In Claude Code, inside your project:
Install github.com/John99101/claude-code-project-template into this project, following its INSTALL.md
Claude previews the install, applies it, then walks you through setup. The config lives in the repo, so your whole team shares it.
git clone --depth 1 https://github.com/John99101/claude-code-project-template /tmp/cct
python3 /tmp/cct/install.py --project ~ /code/my-app # one project, shared with your team via git
python3 /tmp/cct/install.py --global # every project on this device
On Windows, in PowerShell: .\install.ps1 --project C:\code\my-app (or py -3 install.py … ). Add --dry-run to preview. Existing files are never overwritten, and settings.json is merged with a backup. Then run /setup in Claude Code (or /setup --defaults to skip the questions).
New repository. Click Use this template , run python3 install.py --adopt once, restart Claude Code and run /setup .
--update pulls kit updates and keeps your changes; --uninstall removes the kit but keeps files you changed.
Requirements: Claude Code and Python 3.8+ (the hooks use only the standard library). Tested on every push on Linux, macOS and Windows .
Illustrative run on a TypeScript API:
[cct] The Claude Code kit is installed here, but setup isn't finished. Run /setup to finish it.
You /setup
Claude Found: TypeScript + Fastify, pnpm, Vitest, ESLint + Prettier, Postgres, Conventional Commits.
? How thorough should setup be? › Quick (recommended)
? How independently should I work? › Ask before big or risky changes (recommended)
? What must pass before I say I'm done? › pnpm lint, pnpm typecheck, pnpm test
? Run those automatically when I finish? › Yes, the fast checks, only related tests
? Connect any tools via MCP? › GitHub, Postgres
Plan: CLAUDE.md (stack, 7 commands, architecture) · 5 rules + typescript.md
settings: allow pnpm test/lint/typecheck · .mcp.json: GitHub, Postgres
Apply? › Apply
Claude PASS pnpm lint (2.1 s) · PASS pnpm typecheck (4.0 s) · PASS pnpm vitest related (1.3 s)
Done. Set GITHUB_PERSONAL_ACCESS_TOKEN and DATABASE_URI, then restart Claude Code.
Quick setup asks about a dozen questions; Thorough asks about 30. Monorepos get a short brief per package. See the full walkthrough .
A project brief Claude reads every session. CLAUDE.md with your stack, real commands, architecture, conventions and a "before you say it's done" checklist. Topic rules (code style, testing, API design, security, git) and language packs for Python, TypeScript, Go, Rust and Java/Kotlin load only for matching files.
Commands for the whole workflow
Claude also picks up write-tests , debug-issue , release-notes and remember-preference by itself when they fit, and has four specialist agents: planner, reviewer, security auditor and test writer.
Hooks that enforce, not suggest
Guard: blocks rm -rf ~ , curl … | sh , force-pushes to main and writes to .env or keys, even when they're hidden in sh -c , eval or $(…) . It asks before git reset --hard , sudo , DROP TABLE and publishing.
Format and lint on every edit with the tools your project already has (prettier, biome, eslint, ruff, black, gofmt, golangci-lint, rustfmt, clippy, shellcheck). It never downloads anything.
Done gate: your checks run before Claude finishes a turn that changed files; a failure sends Claude back to fix it.
Dependency check: before an install, every package is looked up in its registry. Names that don't exist are blocked (AI models invent package names, and attackers register them); packages under 30 days old or barely downloaded need your OK. New imports that no manifest declares are pointed out too.
Test guard: edits that add skip markers, delete test cases, remove assertions or only change an expected value need your OK, so Claude can't make a test pass by weakening it.
Trifecta check: when MCP servers and permissions give Claude private data, untrusted content and a way to send data out (Simon Willison's "lethal trifecta"), you're warned, and outbound commands such as curl -d need your OK.
Small-batch nudge: past 300 uncommitted lines (lockfiles and generated files excluded), you're told to commit the finished part.
Claim check: if Claude's final reply says "tests pass" or "verified" but no check ran in that turn, it's sent back to run it or correct the claim.
Duplicate hints: a new function that already exists elsewhere ( formatDate / format_date ) is pointed out, so Claude reuses it.
Status line: Opus · my-app (main) · $0.42 · guard standard · setup done .
Session notices. At session start you see what needs attention before Claude touches anything:
[cct] This machine doesn't quite match the project:
- pnpm-lock.yaml changed since the last install. Run `pnpm install`.
- .env is missing STRIPE_KEY (listed in .env.example).
- Node 20 is required (.nvmrc), but `node` is 18.19.1.
[cct] Teammates (Alex) changed the Claude config since your last session: .claude/rules/testing.md.
[cct] There's a handoff note from 2026-11-27 18:05 in .claude-handoff.md. Run /handoff resume to pick up where you left off.
It also flags commands in CLAUDE.md that no longer exist, linters your project configures but you haven't installed, instruction files for other AI tools that disagree with AGENTS.md , and, for teams with Windows, commands, hooks and scripts that won't work there.
/start-issue 482 → branch feat/482-rate-limit, plan, your OK
… Claude works … → every edit formatted and linted; the done gate runs related tests
/review-changes → reviewer and security auditor read the diff
/commit → /pr → PR opened from your template, test plan filled in
/handoff → end of day; tomorrow: /handoff resume
Built on research
Several features answer specific findings about AI-assisted coding:
Anything that makes Claude do extra work is capped, and everything that's just information costs nothing:
Lint feedback stops for a file after 3 reports in a row ( lint_max_repeats ), so Claude never loops on a linter it can't satisfy.
The done gate gives up after 3 tries per turn, and runs only the tests related to the change.
Notices are shown to you, not sent to Claude, so they use no tokens.
The status line shows the session's cost as you go.
Does it work with my stack? Setup detects Node (npm, pnpm, yarn, bun), Python (uv, poetry, pip), Go, Rust, Java/Kotlin and more, plus monorepos (pnpm, Turborepo, Nx, Cargo and Go workspaces). Anything it can't detect, it asks about.
Will it overwrite my existing CLAUDE.md or settings? No. Existing files are never overwritten; setup merges into them and asks before dropping anything you wrote. settings.json is merged with a backup.
Can my team share it? Yes. A project install lives in the repo, so everyone gets the same rules, guard and commands. Personal tweaks go in CLAUDE.local.md and settings.local.json , which stay out of git.
Does it work alongside other plugins? Yes. As a plugin, its commands are namespaced ( /cct:… ). See using it with other plugins .
Does it update itself? No, on purpose: the hooks run code on your machine, so new code only arrives when you choose. Plugin installs update through /plugin . Script installs use /template-update and can opt in to a weekly "new version" notice.
Can I turn things off? Every hook has a switch in .claude/cct.json . See the configuration reference .
Check
When
Test suite : every hook, the installer, updates, uninstall, the plugin and file formats
Every push, on Linux, macOS and Windows
Setup eval runs /setup for real on four sample projects; the question eval checks the questions it would ask
Weekly on GitHub, or any time locally ( how )
MCP catalog check : every server URL and package still exists
Weekly; opens an issue if something breaks
Repository layout
The repository uses Claude Code's standard plugin layout. The plugin uses these folders directly; the installer copies them into a project's .claude/ .
.
├── .claude-plugin/ # plugin manifest, marketplace entry, icon
├── skills/ # setup, start-issue, check, review-changes, commit, pr, handoff, template-update,
│ # re

[truncated]

## Original Extract

Make Claude Code know your project in 5 minutes. One command detects your stack and writes CLAUDE.md, rules, hooks, skills, agents and MCP config, plus a safety guard, a done gate that runs your tests, and /pr, /start-issue and /handoff commands. Plugin or installer; Linux, macOS, Windows. - John991
[truncated]

GitHub - John99101/claude-code-project-template: Make Claude Code know your project in 5 minutes. One command detects your stack and writes CLAUDE.md, rules, hooks, skills, agents and MCP config, plus a safety guard, a done gate that runs your tests, and /pr, /start-issue and /handoff commands. Plugin or installer; Linux, macOS, Windows. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects, and automated work
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
John99101
claude-code-project-template Public template Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Make Claude Code know your project in 5 minutes. One command detects your stack and writes CLAUDE.md, rules, hooks, skills, agents and MCP config, plus a safety guard, a done gate that runs your tests, and /pr, /start-issue and /handoff commands. Plugin or installer; Linux, macOS, Windows.
Readme MIT license Contributing
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
49 Commits 49 Commits Folders and files
.claude-plugin .claude-plugin .github .github agents agents docs docs hooks hooks rule-packs rule-packs rules rules scripts scripts skills skills templates templates tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md INSTALL.md INSTALL.md LICENSE LICENSE README.md README.md VERSION VERSION install.ps1 install.ps1 install.py install.py install.sh install.sh View all files Repository files navigation
Make Claude Code actually know your project, in five minutes.
One command, and Claude reads your repo, asks a few multiple-choice questions, and writes the whole setup for you:
CLAUDE.md , rules, permissions, hooks, skills, agents and MCP servers, all matched to your real stack.
Illustrative session. Every message from the kit in it is the hooks' real output.
/plugin install cct --marketplace John99101/claude-code-project-template
/cct:setup
Quick start · What you get · A day with the kit · FAQ · Full guide
Out of the box, Claude Code doesn't know how your tests run, which folders are generated, what your team's commit style is, or that it must never touch .env . So it guesses, and you correct it over and over.
Claude Code can learn all of this through CLAUDE.md , rules, settings, hooks, skills, agents and MCP. Setting them up by hand takes hours per project, is easy to get wrong, and goes stale. This kit does it for you, and keeps it current.
Plugin (easiest). In Claude Code:
/plugin install cct --marketplace John99101/claude-code-project-template
Then run /cct:setup in any project. Updates arrive through /plugin . On older Claude Code versions, add the marketplace first with /plugin marketplace add John99101/claude-code-project-template , then /plugin install cct@cct-kit .
Ask Claude. In Claude Code, inside your project:
Install github.com/John99101/claude-code-project-template into this project, following its INSTALL.md
Claude previews the install, applies it, then walks you through setup. The config lives in the repo, so your whole team shares it.
git clone --depth 1 https://github.com/John99101/claude-code-project-template /tmp/cct
python3 /tmp/cct/install.py --project ~ /code/my-app # one project, shared with your team via git
python3 /tmp/cct/install.py --global # every project on this device
On Windows, in PowerShell: .\install.ps1 --project C:\code\my-app (or py -3 install.py … ). Add --dry-run to preview. Existing files are never overwritten, and settings.json is merged with a backup. Then run /setup in Claude Code (or /setup --defaults to skip the questions).
New repository. Click Use this template , run python3 install.py --adopt once, restart Claude Code and run /setup .
--update pulls kit updates and keeps your changes; --uninstall removes the kit but keeps files you changed.
Requirements: Claude Code and Python 3.8+ (the hooks use only the standard library). Tested on every push on Linux, macOS and Windows .
Illustrative run on a TypeScript API:
[cct] The Claude Code kit is installed here, but setup isn't finished. Run /setup to finish it.
You /setup
Claude Found: TypeScript + Fastify, pnpm, Vitest, ESLint + Prettier, Postgres, Conventional Commits.
? How thorough should setup be? › Quick (recommended)
? How independently should I work? › Ask before big or risky changes (recommended)
? What must pass before I say I'm done? › pnpm lint, pnpm typecheck, pnpm test
? Run those automatically when I finish? › Yes, the fast checks, only related tests
? Connect any tools via MCP? › GitHub, Postgres
Plan: CLAUDE.md (stack, 7 commands, architecture) · 5 rules + typescript.md
settings: allow pnpm test/lint/typecheck · .mcp.json: GitHub, Postgres
Apply? › Apply
Claude PASS pnpm lint (2.1 s) · PASS pnpm typecheck (4.0 s) · PASS pnpm vitest related (1.3 s)
Done. Set GITHUB_PERSONAL_ACCESS_TOKEN and DATABASE_URI, then restart Claude Code.
Quick setup asks about a dozen questions; Thorough asks about 30. Monorepos get a short brief per package. See the full walkthrough .
A project brief Claude reads every session. CLAUDE.md with your stack, real commands, architecture, conventions and a "before you say it's done" checklist. Topic rules (code style, testing, API design, security, git) and language packs for Python, TypeScript, Go, Rust and Java/Kotlin load only for matching files.
Commands for the whole workflow
Claude also picks up write-tests , debug-issue , release-notes and remember-preference by itself when they fit, and has four specialist agents: planner, reviewer, security auditor and test writer.
Hooks that enforce, not suggest
Guard: blocks rm -rf ~ , curl … | sh , force-pushes to main and writes to .env or keys, even when they're hidden in sh -c , eval or $(…) . It asks before git reset --hard , sudo , DROP TABLE and publishing.
Format and lint on every edit with the tools your project already has (prettier, biome, eslint, ruff, black, gofmt, golangci-lint, rustfmt, clippy, shellcheck). It never downloads anything.
Done gate: your checks run before Claude finishes a turn that changed files; a failure sends Claude back to fix it.
Dependency check: before an install, every package is looked up in its registry. Names that don't exist are blocked (AI models invent package names, and attackers register them); packages under 30 days old or barely downloaded need your OK. New imports that no manifest declares are pointed out too.
Test guard: edits that add skip markers, delete test cases, remove assertions or only change an expected value need your OK, so Claude can't make a test pass by weakening it.
Trifecta check: when MCP servers and permissions give Claude private data, untrusted content and a way to send data out (Simon Willison's "lethal trifecta"), you're warned, and outbound commands such as curl -d need your OK.
Small-batch nudge: past 300 uncommitted lines (lockfiles and generated files excluded), you're told to commit the finished part.
Claim check: if Claude's final reply says "tests pass" or "verified" but no check ran in that turn, it's sent back to run it or correct the claim.
Duplicate hints: a new function that already exists elsewhere ( formatDate / format_date ) is pointed out, so Claude reuses it.
Status line: Opus · my-app (main) · $0.42 · guard standard · setup done .
Session notices. At session start you see what needs attention before Claude touches anything:
[cct] This machine doesn't quite match the project:
- pnpm-lock.yaml changed since the last install. Run `pnpm install`.
- .env is missing STRIPE_KEY (listed in .env.example).
- Node 20 is required (.nvmrc), but `node` is 18.19.1.
[cct] Teammates (Alex) changed the Claude config since your last session: .claude/rules/testing.md.
[cct] There's a handoff note from 2026-11-27 18:05 in .claude-handoff.md. Run /handoff resume to pick up where you left off.
It also flags commands in CLAUDE.md that no longer exist, linters your project configures but you haven't installed, instruction files for other AI tools that disagree with AGENTS.md , and, for teams with Windows, commands, hooks and scripts that won't work there.
/start-issue 482 → branch feat/482-rate-limit, plan, your OK
… Claude works … → every edit formatted and linted; the done gate runs related tests
/review-changes → reviewer and security auditor read the diff
/commit → /pr → PR opened from your template, test plan filled in
/handoff → end of day; tomorrow: /handoff resume
Built on research
Several features answer specific findings about AI-assisted coding:
Anything that makes Claude do extra work is capped, and everything that's just information costs nothing:
Lint feedback stops for a file after 3 reports in a row ( lint_max_repeats ), so Claude never loops on a linter it can't satisfy.
The done gate gives up after 3 tries per turn, and runs only the tests related to the change.
Notices are shown to you, not sent to Claude, so they use no tokens.
The status line shows the session's cost as you go.
Does it work with my stack? Setup detects Node (npm, pnpm, yarn, bun), Python (uv, poetry, pip), Go, Rust, Java/Kotlin and more, plus monorepos (pnpm, Turborepo, Nx, Cargo and Go workspaces). Anything it can't detect, it asks about.
Will it overwrite my existing CLAUDE.md or settings? No. Existing files are never overwritten; setup merges into them and asks before dropping anything you wrote. settings.json is merged with a backup.
Can my team share it? Yes. A project install lives in the repo, so everyone gets the same rules, guard and commands. Personal tweaks go in CLAUDE.local.md and settings.local.json , which stay out of git.
Does it work alongside other plugins? Yes. As a plugin, its commands are namespaced ( /cct:… ). See using it with other plugins .
Does it update itself? No, on purpose: the hooks run code on your machine, so new code only arrives when you choose. Plugin installs update through /plugin . Script installs use /template-update and can opt in to a weekly "new version" notice.
Can I turn things off? Every hook has a switch in .claude/cct.json . See the configuration reference .
Check
When
Test suite : every hook, the installer, updates, uninstall, the plugin and file formats
Every push, on Linux, macOS and Windows
Setup eval runs /setup for real on four sample projects; the question eval checks the questions it would ask
Weekly on GitHub, or any time locally ( how )
MCP catalog check : every server URL and package still exists
Weekly; opens an issue if something breaks
Repository layout
The repository uses Claude Code's standard plugin layout. The plugin uses these folders directly; the installer copies them into a project's .claude/ .
.
├── .claude-plugin/ # plugin manifest, marketplace entry, icon
├── skills/ # setup, start-issue, check, review-changes, commit, pr, handoff, template-update,
│ # re

[truncated]
