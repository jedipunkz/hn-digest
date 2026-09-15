---
source: "https://github.com/andreylukin/bough"
hn_url: "https://news.ycombinator.com/item?id=49711939"
title: "Show HN: Bough, the agent I built to replace Claude Code at work"
article_title: "GitHub - andreylukin/bough: A terminal coding agent where the model writes one program instead of calling tools one at a time. · GitHub"
image: "https://opengraph.githubassets.com/93573d239c53d4310bc1ab9b03e44f1010c158173fede0a91295e55018e663a7/andreylukin/bough"
author: "alukin"
captured_at: "2026-09-15T13:11:19Z"
capture_tool: "hn-digest"
hn_id: 49711939
score: 2
comments: 0
posted_at: "2026-09-15T13:04:27Z"
tags:
  - hacker-news
---

# Show HN: Bough, the agent I built to replace Claude Code at work

- HN: [49711939](https://news.ycombinator.com/item?id=49711939)
- Source: [github.com](https://github.com/andreylukin/bough)
- Score: 2
- Comments: 0
- Posted: 2026-09-15T13:04:27Z

## Translation

Title: Show HN: Bough, the agent I built to replace Claude Code at work
Article title: GitHub - andreylukin/bough: A terminal coding agent where the model writes one program instead of calling tools one at a time. · GitHub
Description: A terminal coding agent where the model writes one program instead of calling tools one at a time. - andreylukin/bough

Article text:
GitHub - andreylukin/bough: A terminal coding agent where the model writes one program instead of calling tools one at a time. · GitHub
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
andreylukin
/
bough
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2,106 Commits 2,106 Commits Folders and files
.design-sync .design-sync .githooks .githooks .github .github Formula Formula assets assets bench bench go go scripts/ demo scripts/ demo .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md EXAMPLES.md EXAMPLES.md LICENSE LICENSE README.md README.md SCREENSHOTS.md SCREENSHOTS.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md install.sh install.sh View all files Repository files navigation
A terminal coding agent where the model writes one program instead of calling tools one at a time.
Code mode. The model acts by writing JavaScript in a ```js block, which bough runs; what it prints goes back to the model. tools.view , tools.patch , tools.bash , tools.spawn and every MCP tool are ordinary functions inside it. A patch and the test run that checks it can be one step, and the model branches on results in code rather than in another round trip.
Here is one step from a recording of this demo, verbatim as openai/gpt-6-astra wrote it: patch, format, test and review the diff, in one program.
console . log ( tools . patch ( "wordfreq.go" ,
"\tsort.Slice(all, func(i, j int) bool { return all[i].N > all[j].N })\n\treturn all[:n]" ,
"\tsort.Slice(all, func(i, j int) bool {\n\t\tif all[i].N == all[j].N {\n\t\t\treturn all[i].Word < all[j].Word\n\t\t}\n\t\treturn all[i].N > all[j].N\n\t})\n\tif n > len(all) {\n\t\tn = len(all)\n\t}\n\treturn all[:n]" ) ) ;
console . log ( tools . bash ( "gofmt -w wordfreq.go; go test ./...; git diff --check; git diff -- wordfreq.go" ) ) ;
From the author's own sessions (Sept 2–15, 2026, mostly gpt-6-astra ): 735 turns ran 5,521 programs that made 10,542 tool calls, 1.9 per program. 52% of programs made more than one call and 30% branched or looped on results ( if , for , try ). Counts are call sites in the program text, so loops make them a floor. An agent that makes one tool call per model round trip would have needed about 10.5k round trips for that work; one that makes parallel calls, fewer.
Everything is a plugin. The provider, the loop, the tools, the UI, MCP, hooks and skills are rows in a YAML file. Swap one, disable one, or save the file mid-session and the running process reconciles.
One binary, your keys, no telemetry. Anthropic, OpenAI, OpenRouter or Cerebras. Sessions are append-only JSONL under ~/.bough/history , so resume, search and switching models mid-conversation just work.
curl -fsSL https://raw.githubusercontent.com/andreylukin/bough/main/install.sh | sh
printf ' say hello\n ' | bough --headless --set llm.plugin=llm-echo # no key needed
Then add a key and start it in a repo:
echo ' OPENROUTER_API_KEY=sk-or-... ' >> ~ /.bough/env # or ANTHROPIC_ / OPENAI_ / CEREBRAS_API_KEY
bough
Started inside a git checkout, bough can edit files there; anywhere else it only reads. Prefer a browser? Run bough serve from your repo and open the link: the first visit asks for a key and a folder, then starts a session.
macOS and Linux, x86-64 and arm64. Also brew tap andreylukin/bough https://github.com/andreylukin/bough && brew install bough , or cd go && go build ./cmd/bough (Go 1.27+).
bough
the terminal UI: / palette, @file , !shell , esc esc to rewind, -c to resume
bough serve
a web control room for every session. The first visit walks you through adding a key and picking a repo; after that it shows live transcripts, questions waiting on you, and pasted images
bough --headless
stdin in, events out ( --json ), for scripts and benchmarks
Skills load by name. Here a session names the parallel skill, searches the web with parallel-cli from inside its program, and answers with sources:
Plugin
Every part of bough (the provider, loop, tools, TUI, web UI) is a plugin mounted as a row in bough.yml . Rows find each other only through service keys, and saving the file mid-session remounts just what changed. PLUGINS.md
Hook
A plain .js file in ~/.bough/hooks/<event>/ or ./.bough/hooks/<event>/ that runs when that event fires: session-start , user-prompt-submit , pre-code-exec , post-result or stop . It can rewrite the prompt, block a program before it runs, or amend its result. Files are re-read on every fire, so edits apply without a restart. Hooks
Skill
A SKILL.md in ~/.claude/skills/<name>/ or ./.claude/skills/<name>/ , the same layout Claude Code uses. Mention the skill's name in a prompt and its instructions are added to that turn. Skills
Project
A definition in ~/.bough/projects/<slug>/ : its repos and branches ( project.yml ) plus a setup.sh that builds its container image. bough --project <slug> starts a session that can change that project's code.
Orb
A project session's workspace: git worktrees of the project's repos plus a Linux container built from the project image. Every shell command runs inside the container, and each turn is checkpointed. macOS (Apple container ) for now. orbs.md
LLM wiki
Markdown pages an agent compiles from your session logs: decisions, root causes, gotchas, each citing the log entry it came from. Never injected into prompts. Below
A session with no --project is local : it runs on your machine and can write only inside the git checkout it started in.
Having the model act in code is not a new idea: CodeAct showed executable code actions beat JSON tool calls, smolagents builds agents that think in Python, and Cloudflare's Code Mode has the model write TypeScript against MCP servers. bough takes that idea to a full interactive coding agent, where patching, the shell, subagents, background jobs and every MCP tool are functions in the same runtime.
bough is closest to Claude Code, opencode and pi, and borrows their conventions on purpose: AGENTS.md / CLAUDE.md , skills, hooks, MCP, subagents. The differences are architectural. Other agents expose a list of tools and loop once per call; bough exposes one program runner, so the model batches work and branches on results in code. And where those tools are applications you configure, bough is a small kernel where each part, including the UI, is a replaceable row.
It is a one-person project in daily use. Known gaps:
No Windows build. macOS and Linux only.
Project sessions (containers) need macOS; the Linux container runtimes are stubs.
Local sessions have no per-turn file checkpoints, so rewinding the conversation does not undo edits. Use git.
The usage numbers above come from one person's sessions, not a controlled benchmark.
The shell is not sandboxed by default. Here is what bough does and does not limit:
File tools stay in your repo. Started inside a git checkout, tools.write and tools.patch work only under that checkout and refuse any other path, symlinks included. Started anywhere else, the session is read-only and those tools don't exist.
The shell runs as you. tools.bash has your files, credentials and network, like a script you ran yourself. The prompt tells the model not to change files outside the checkout with the shell, but nothing enforces that. Keep work committed and read what it proposes.
For an enforced boundary, use a project session. bough --project <slug> runs the shell in a Linux container (Apple container , so macOS for now) against git worktrees, and snapshots every turn so /undo restores files. See orbs .
No telemetry. bough talks to your LLM provider, the MCP servers you configure, the public models.dev price list, and GitHub when you run bough update .
bough is a small kernel of services, events and a row loader, and everything else is a plugin (the same shape as DeepSeek Harness ). The provider, the loop, the tools, subagents, history, MCP, hooks, skills, the TUI and the web UI are rows in bough.yml that find each other only through service keys. ./bough.yml (else ~/.bough/bough.yml ) overrides the shipped rows by id:
- id : llm
plugin : llm-openrouter
config :
model : openai/gpt-6-astra
Save it mid-session and only the changed rows and their dependents remount; the conversation survives because context is rebuilt from the session log. bough rows shows the live tree. ~/.bough/init.js adds tools, commands and whole providers in a few lines ( INIT.md ); a new row is a Go plugin ( PLUGINS.md ).
Every session is an append-only log. bough wiki install has an agent compile those logs, every few minutes, into ~/.bough/wiki : markdown pages of decisions, root causes and gotchas, where every claim cites the exact log entry it came from. Nothing is injected into prompts; an agent reads the wiki when asked, and bough wiki check verifies every citation. bough serve shows each page as claims beside their evidence.
EXAMPLES.md walks through all of it with working snippets: plugin rows, what the model's programs look like, subagents, init.js , hooks, skills and rules, MCP, the wiki, scripting and replay. The full reference is go/README.md .
./.githooks/install
cd go && go build ./cmd/bough
go test -race -parallel 4 ./...
(cd internal/serve/web && bun install && bun run build) # web UI; dist/ is committed
(cd tests/web && npm ci && npx playwright install chromium && npm test)
The README recordings are scripts: scripts/demo . Conventions and traps are in AGENTS.md .
A terminal coding agent where the model writes one program instead of calling tools one at a time.
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A terminal coding agent where the model writes one program instead of calling tools one at a time. - andreylukin/bough

GitHub - andreylukin/bough: A terminal coding agent where the model writes one program instead of calling tools one at a time. · GitHub
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
andreylukin
/
bough
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
2,106 Commits 2,106 Commits Folders and files
.design-sync .design-sync .githooks .githooks .github .github Formula Formula assets assets bench bench go go scripts/ demo scripts/ demo .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md EXAMPLES.md EXAMPLES.md LICENSE LICENSE README.md README.md SCREENSHOTS.md SCREENSHOTS.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md install.sh install.sh View all files Repository files navigation
A terminal coding agent where the model writes one program instead of calling tools one at a time.
Code mode. The model acts by writing JavaScript in a ```js block, which bough runs; what it prints goes back to the model. tools.view , tools.patch , tools.bash , tools.spawn and every MCP tool are ordinary functions inside it. A patch and the test run that checks it can be one step, and the model branches on results in code rather than in another round trip.
Here is one step from a recording of this demo, verbatim as openai/gpt-6-astra wrote it: patch, format, test and review the diff, in one program.
console . log ( tools . patch ( "wordfreq.go" ,
"\tsort.Slice(all, func(i, j int) bool { return all[i].N > all[j].N })\n\treturn all[:n]" ,
"\tsort.Slice(all, func(i, j int) bool {\n\t\tif all[i].N == all[j].N {\n\t\t\treturn all[i].Word < all[j].Word\n\t\t}\n\t\treturn all[i].N > all[j].N\n\t})\n\tif n > len(all) {\n\t\tn = len(all)\n\t}\n\treturn all[:n]" ) ) ;
console . log ( tools . bash ( "gofmt -w wordfreq.go; go test ./...; git diff --check; git diff -- wordfreq.go" ) ) ;
From the author's own sessions (Sept 2–15, 2026, mostly gpt-6-astra ): 735 turns ran 5,521 programs that made 10,542 tool calls, 1.9 per program. 52% of programs made more than one call and 30% branched or looped on results ( if , for , try ). Counts are call sites in the program text, so loops make them a floor. An agent that makes one tool call per model round trip would have needed about 10.5k round trips for that work; one that makes parallel calls, fewer.
Everything is a plugin. The provider, the loop, the tools, the UI, MCP, hooks and skills are rows in a YAML file. Swap one, disable one, or save the file mid-session and the running process reconciles.
One binary, your keys, no telemetry. Anthropic, OpenAI, OpenRouter or Cerebras. Sessions are append-only JSONL under ~/.bough/history , so resume, search and switching models mid-conversation just work.
curl -fsSL https://raw.githubusercontent.com/andreylukin/bough/main/install.sh | sh
printf ' say hello\n ' | bough --headless --set llm.plugin=llm-echo # no key needed
Then add a key and start it in a repo:
echo ' OPENROUTER_API_KEY=sk-or-... ' >> ~ /.bough/env # or ANTHROPIC_ / OPENAI_ / CEREBRAS_API_KEY
bough
Started inside a git checkout, bough can edit files there; anywhere else it only reads. Prefer a browser? Run bough serve from your repo and open the link: the first visit asks for a key and a folder, then starts a session.
macOS and Linux, x86-64 and arm64. Also brew tap andreylukin/bough https://github.com/andreylukin/bough && brew install bough , or cd go && go build ./cmd/bough (Go 1.27+).
bough
the terminal UI: / palette, @file , !shell , esc esc to rewind, -c to resume
bough serve
a web control room for every session. The first visit walks you through adding a key and picking a repo; after that it shows live transcripts, questions waiting on you, and pasted images
bough --headless
stdin in, events out ( --json ), for scripts and benchmarks
Skills load by name. Here a session names the parallel skill, searches the web with parallel-cli from inside its program, and answers with sources:
Plugin
Every part of bough (the provider, loop, tools, TUI, web UI) is a plugin mounted as a row in bough.yml . Rows find each other only through service keys, and saving the file mid-session remounts just what changed. PLUGINS.md
Hook
A plain .js file in ~/.bough/hooks/<event>/ or ./.bough/hooks/<event>/ that runs when that event fires: session-start , user-prompt-submit , pre-code-exec , post-result or stop . It can rewrite the prompt, block a program before it runs, or amend its result. Files are re-read on every fire, so edits apply without a restart. Hooks
Skill
A SKILL.md in ~/.claude/skills/<name>/ or ./.claude/skills/<name>/ , the same layout Claude Code uses. Mention the skill's name in a prompt and its instructions are added to that turn. Skills
Project
A definition in ~/.bough/projects/<slug>/ : its repos and branches ( project.yml ) plus a setup.sh that builds its container image. bough --project <slug> starts a session that can change that project's code.
Orb
A project session's workspace: git worktrees of the project's repos plus a Linux container built from the project image. Every shell command runs inside the container, and each turn is checkpointed. macOS (Apple container ) for now. orbs.md
LLM wiki
Markdown pages an agent compiles from your session logs: decisions, root causes, gotchas, each citing the log entry it came from. Never injected into prompts. Below
A session with no --project is local : it runs on your machine and can write only inside the git checkout it started in.
Having the model act in code is not a new idea: CodeAct showed executable code actions beat JSON tool calls, smolagents builds agents that think in Python, and Cloudflare's Code Mode has the model write TypeScript against MCP servers. bough takes that idea to a full interactive coding agent, where patching, the shell, subagents, background jobs and every MCP tool are functions in the same runtime.
bough is closest to Claude Code, opencode and pi, and borrows their conventions on purpose: AGENTS.md / CLAUDE.md , skills, hooks, MCP, subagents. The differences are architectural. Other agents expose a list of tools and loop once per call; bough exposes one program runner, so the model batches work and branches on results in code. And where those tools are applications you configure, bough is a small kernel where each part, including the UI, is a replaceable row.
It is a one-person project in daily use. Known gaps:
No Windows build. macOS and Linux only.
Project sessions (containers) need macOS; the Linux container runtimes are stubs.
Local sessions have no per-turn file checkpoints, so rewinding the conversation does not undo edits. Use git.
The usage numbers above come from one person's sessions, not a controlled benchmark.
The shell is not sandboxed by default. Here is what bough does and does not limit:
File tools stay in your repo. Started inside a git checkout, tools.write and tools.patch work only under that checkout and refuse any other path, symlinks included. Started anywhere else, the session is read-only and those tools don't exist.
The shell runs as you. tools.bash has your files, credentials and network, like a script you ran yourself. The prompt tells the model not to change files outside the checkout with the shell, but nothing enforces that. Keep work committed and read what it proposes.
For an enforced boundary, use a project session. bough --project <slug> runs the shell in a Linux container (Apple container , so macOS for now) against git worktrees, and snapshots every turn so /undo restores files. See orbs .
No telemetry. bough talks to your LLM provider, the MCP servers you configure, the public models.dev price list, and GitHub when you run bough update .
bough is a small kernel of services, events and a row loader, and everything else is a plugin (the same shape as DeepSeek Harness ). The provider, the loop, the tools, subagents, history, MCP, hooks, skills, the TUI and the web UI are rows in bough.yml that find each other only through service keys. ./bough.yml (else ~/.bough/bough.yml ) overrides the shipped rows by id:
- id : llm
plugin : llm-openrouter
config :
model : openai/gpt-6-astra
Save it mid-session and only the changed rows and their dependents remount; the conversation survives because context is rebuilt from the session log. bough rows shows the live tree. ~/.bough/init.js adds tools, commands and whole providers in a few lines ( INIT.md ); a new row is a Go plugin ( PLUGINS.md ).
Every session is an append-only log. bough wiki install has an agent compile those logs, every few minutes, into ~/.bough/wiki : markdown pages of decisions, root causes and gotchas, where every claim cites the exact log entry it came from. Nothing is injected into prompts; an agent reads the wiki when asked, and bough wiki check verifies every citation. bough serve shows each page as claims beside their evidence.
EXAMPLES.md walks through all of it with working snippets: plugin rows, what the model's programs look like, subagents, init.js , hooks, skills and rules, MCP, the wiki, scripting and replay. The full reference is go/README.md .
./.githooks/install
cd go && go build ./cmd/bough
go test -race -parallel 4 ./...
(cd internal/serve/web && bun install && bun run build) # web UI; dist/ is committed
(cd tests/web && npm ci && npx playwright install chromium && npm test)
The README recordings are scripts: scripts/demo . Conventions and traps are in AGENTS.md .
A terminal coding agent where the model writes one program instead of calling tools one at a time.
Readme Apache-2.0 license Code of conduct
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
