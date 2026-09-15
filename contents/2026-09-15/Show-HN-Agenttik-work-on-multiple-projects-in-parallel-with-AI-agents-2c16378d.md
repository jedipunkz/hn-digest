---
source: "https://github.com/pausan/agenttik"
hn_url: "https://news.ycombinator.com/item?id=49720222"
title: "Show HN: Agenttik – work on multiple projects in parallel with AI agents"
article_title: "GitHub - pausan/agenttik: Manage multiple projects and agents with ease · GitHub"
image: "https://opengraph.githubassets.com/11938a0b666bbe45571b41021be556f857fa296f4e70b1ca5be73193d271af46/pausan/agenttik"
author: "psanchez"
captured_at: "2026-09-15T23:59:55Z"
capture_tool: "hn-digest"
hn_id: 49720222
score: 3
comments: 0
posted_at: "2026-09-15T23:24:53Z"
tags:
  - hacker-news
---

# Show HN: Agenttik – work on multiple projects in parallel with AI agents

- HN: [49720222](https://news.ycombinator.com/item?id=49720222)
- Source: [github.com](https://github.com/pausan/agenttik)
- Score: 3
- Comments: 0
- Posted: 2026-09-15T23:24:53Z

## Translation

Title: Show HN: Agenttik – work on multiple projects in parallel with AI agents
Article title: GitHub - pausan/agenttik: Manage multiple projects and agents with ease · GitHub
Description: Manage multiple projects and agents with ease. Contribute to pausan/agenttik development by creating an account on GitHub.
HN text: Hi! I've been reading HN daily for almost 16 years. Really glad this community
exist. This is my first Show HN, so I hope you find this tool useful/interesting. During my summer holidays I started playing with latest AI models on personal
projects. At one point I started working in parallel with the same project and
tried different approaches. The one that worked best for me was to clone the
repo and distribute work among different AI workers, work in parallel but with
different branches & folders. This allowed me to review and merge work
afterwards. The thing is, spawning several agents in parallel was challenging. I could work
on multiple features and bugfix at once, but it was hard to keep track what I
was working on. I was jumping from one thing to another, and back; from one
window to another. Switching tasks like this was mentally exhausting. Holidays were over, and after returning to my day to day work, I continued
working using a similar approach not on one, but also on multiple projects, but
I still could not find how to work in a more organized way without the mental
toll. This space is growing really fast, and I haven't been able to find the right
tool for myself, so I decided to build my own. I wanted a tool that would help me manage multiple projects at once, work with
multiple agents, even from different providers, where I could chose between my
personal and work subscriptions depending on the project, where I could easily
review changes without leaving the tool if I wanted, or go all in and delegate
fully to agents; one where I had quick shortcuts to move around, where I could
schedule work or enqueue work for later so agents could work on one task after
another. agenttik is my way of tackling this problem to work on multiple projects in
parallel, while allowing myself to be on the driver seat. I can give as much
autonomy as I want to the agents while keeping things organized in my head. I've been hesitating for hours on whether to submit or not, but hey "if you are not ashamed of your project, you are shipping too late" Honestly, this project has exactly one week of development, started from
scratch, and even though it is still early on, it has been extremely useful for
me. I've built in a week things that would have taken me months, not only built
the tool itself but other projects as well (most of the merit is still on the
models used, but would have had no way of distributing all the work, though). Feel free to try. Here's how I would get started: - Open a project you already have - Add some general prompt (in the project area) or add to your
AGENTS.md/CLAUDE.md (to instruct it to do regular commits and/or work with branches, and also instructions on how to build effective tests so that it can self-correct). - Chose your subscriptin of choice: Claude Code / Open AI subscription / ... - Pick big models (Opus High to Max / Fable Medium to Max / Astra Medium to Max)
for bigger autonomy and try not to do hand-holding Finally, start sending/enqueuing tasks on all the features you want your product
to have, bugfixes, ... use a new prompt per task (Ctrl+T / Cmd+T), I recommend
using big models to avoid investing your own time reviewing, and then once LLM
is done, review and archive that task or continue iterating. Keep adding tasks
as you think on them. Add more projects as you start feeling comfortable. It has been really addictive for me to see the speed of development when working
like this. I got totally hooked. Speed of iteration is insane. Happy to discuss any feedback you have even if it is about other tools or ways
of working with AI these days! I hope you like it!!

Article text:
GitHub - pausan/agenttik: Manage multiple projects and agents with ease · GitHub
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
pausan
/
agenttik
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
275 Commits 275 Commits Folders and files
.githooks .githooks .github/ workflows .github/ workflows app app docs/ screenshots docs/ screenshots e2e e2e scripts scripts specs specs web web .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
A local workspace for running AI coding sessions across multiple projects.
Use Claude Code, Codex and GitHub Copilot through their official CLIs, with your existing login,
in one desktop window or browser tab. Desktop builds support Linux, macOS ARM64,
and Windows x64.
Project workspaces: organize sessions and file tabs; archive and restore conversations.
Orchestrator: enable a separate pinned project in Settings to inspect work and manage tasks across projects, with editable instructions and a reset to the built-in prompt.
Live sessions: stream replies, resume conversations, and queue prompts per project.
Reusable work: pin prompts and schedule recurring tasks.
Model controls: choose models, effort, permissions, and favorite combinations.
Code tools: browse and edit files, preview Markdown/HTML, inspect Git diffs, stage changes, and write commits.
Usage tracking: view tokens, context usage, subscription allowance, and costs where reported.
Local state: SQLite history, remembered tabs, customizable shortcuts, and color themes. Use profiles for separate workspaces or private mode for a temporary one.
Claude Code is working; Codex, GitHub Copilot and OpenCode Go are implemented, with live-turn validation still pending. OpenCode Go supports direct use with a subscription key and an optional CLI; Settings defaults to the CLI when installed.
Captured from a fresh private instance and browser context with fictional projects
and demo conversation data. No live agent was run. Click an image to view it full size.
Open a changed file beside your project tasks. Compare the diff, stage individual
files, and prepare a commit from the workspace panel.
Give a recurring task its own prompt, model, and schedule. Review its next run and
history, or pause it from the same view.
Use Go 1.25+, Node.js 22.12+ with npm, and either a logged-in claude , codex or copilot CLI on PATH , or an OpenCode Go key configured in Settings → Subscriptions. The opencode CLI is optional.
make run-web
Open localhost:7717 , add a project folder, and start a session.
Or add one from the terminal with agenttik --init in the folder you want —
agenttik --init path/to/repo names another — which works whether or not the
app is already open, and shows up in an open window straight away.
Remote servers and private workspaces
Connect to another server with agenttik --remote host:7717 (or an HTTPS
URL), or choose Connect to remote server in the command palette. The
client checks /api/version first, then shows the server’s login if needed.
See remote connections .
Start a temporary, separate instance with agenttik --private . Its app data is
removed on exit; project files stay on disk. Manage local profiles in
Settings → Profiles . With multiple profiles, the picker before Shortcuts
switches between their isolated projects and tasks.
--addr and --data-dir change where it listens and where it keeps its
database, --web skips the desktop window, and --help and --version say
what this build is.
For the desktop app, install the native build dependencies ( make deps on
Debian/Ubuntu), then run make run . make build-windows-amd64 cross-compiles
the Windows x64 binary; build the macOS ARM64 target on macOS with
make build-macos-arm64 . macOS builds also create an ad-hoc-signed
agenttik.app and a ZIP ( bin/ for make build , dist/ for the ARM64 target).
Extract the release ZIP and move the app to Applications. No Apple developer
account is needed to build it. Downloaded builds are not notarized; after a
blocked launch, use System Settings → Privacy & Security → Open Anyway.
On macOS, Close to tray works even without Accessibility permission. The global
show/hide shortcut needs that permission; grant it in System Settings → Privacy
& Security → Accessibility, then restart agenttik. Shortcut failures appear in
agenttik's desktop settings while the tray menu remains usable.
Local access by default: agenttik binds to loopback. Anyone who can reach an
unauthenticated server has full control of it. Before exposing it through
Settings → Server , configure authentication and use a trusted network. See
server access .
Built with Go, SQLite, Vue 3, Nuxt UI, and Wails. See specs
for architecture and implementation details.
Manage multiple projects and agents with ease
Readme MIT license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Manage multiple projects and agents with ease. Contribute to pausan/agenttik development by creating an account on GitHub.

Hi! I've been reading HN daily for almost 16 years. Really glad this community
exist. This is my first Show HN, so I hope you find this tool useful/interesting. During my summer holidays I started playing with latest AI models on personal
projects. At one point I started working in parallel with the same project and
tried different approaches. The one that worked best for me was to clone the
repo and distribute work among different AI workers, work in parallel but with
different branches & folders. This allowed me to review and merge work
afterwards. The thing is, spawning several agents in parallel was challenging. I could work
on multiple features and bugfix at once, but it was hard to keep track what I
was working on. I was jumping from one thing to another, and back; from one
window to another. Switching tasks like this was mentally exhausting. Holidays were over, and after returning to my day to day work, I continued
working using a similar approach not on one, but also on multiple projects, but
I still could not find how to work in a more organized way without the mental
toll. This space is growing really fast, and I haven't been able to find the right
tool for myself, so I decided to build my own. I wanted a tool that would help me manage multiple projects at once, work with
multiple agents, even from different providers, where I could chose between my
personal and work subscriptions depending on the project, where I could easily
review changes without leaving the tool if I wanted, or go all in and delegate
fully to agents; one where I had quick shortcuts to move around, where I could
schedule work or enqueue work for later so agents could work on one task after
another. agenttik is my way of tackling this problem to work on multiple projects in
parallel, while allowing myself to be on the driver seat. I can give as much
autonomy as I want to the agents while keeping things organized in my head. I've been hesitating for hours on whether to submit or not, but hey "if you are not ashamed of your project, you are shipping too late" Honestly, this project has exactly one week of development, started from
scratch, and even though it is still early on, it has been extremely useful for
me. I've built in a week things that would have taken me months, not only built
the tool itself but other projects as well (most of the merit is still on the
models used, but would have had no way of distributing all the work, though). Feel free to try. Here's how I would get started: - Open a project you already have - Add some general prompt (in the project area) or add to your
AGENTS.md/CLAUDE.md (to instruct it to do regular commits and/or work with branches, and also instructions on how to build effective tests so that it can self-correct). - Chose your subscriptin of choice: Claude Code / Open AI subscription / ... - Pick big models (Opus High to Max / Fable Medium to Max / Astra Medium to Max)
for bigger autonomy and try not to do hand-holding Finally, start sending/enqueuing tasks on all the features you want your product
to have, bugfixes, ... use a new prompt per task (Ctrl+T / Cmd+T), I recommend
using big models to avoid investing your own time reviewing, and then once LLM
is done, review and archive that task or continue iterating. Keep adding tasks
as you think on them. Add more projects as you start feeling comfortable. It has been really addictive for me to see the speed of development when working
like this. I got totally hooked. Speed of iteration is insane. Happy to discuss any feedback you have even if it is about other tools or ways
of working with AI these days! I hope you like it!!

GitHub - pausan/agenttik: Manage multiple projects and agents with ease · GitHub
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
pausan
/
agenttik
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
275 Commits 275 Commits Folders and files
.githooks .githooks .github/ workflows .github/ workflows app app docs/ screenshots docs/ screenshots e2e e2e scripts scripts specs specs web web .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
A local workspace for running AI coding sessions across multiple projects.
Use Claude Code, Codex and GitHub Copilot through their official CLIs, with your existing login,
in one desktop window or browser tab. Desktop builds support Linux, macOS ARM64,
and Windows x64.
Project workspaces: organize sessions and file tabs; archive and restore conversations.
Orchestrator: enable a separate pinned project in Settings to inspect work and manage tasks across projects, with editable instructions and a reset to the built-in prompt.
Live sessions: stream replies, resume conversations, and queue prompts per project.
Reusable work: pin prompts and schedule recurring tasks.
Model controls: choose models, effort, permissions, and favorite combinations.
Code tools: browse and edit files, preview Markdown/HTML, inspect Git diffs, stage changes, and write commits.
Usage tracking: view tokens, context usage, subscription allowance, and costs where reported.
Local state: SQLite history, remembered tabs, customizable shortcuts, and color themes. Use profiles for separate workspaces or private mode for a temporary one.
Claude Code is working; Codex, GitHub Copilot and OpenCode Go are implemented, with live-turn validation still pending. OpenCode Go supports direct use with a subscription key and an optional CLI; Settings defaults to the CLI when installed.
Captured from a fresh private instance and browser context with fictional projects
and demo conversation data. No live agent was run. Click an image to view it full size.
Open a changed file beside your project tasks. Compare the diff, stage individual
files, and prepare a commit from the workspace panel.
Give a recurring task its own prompt, model, and schedule. Review its next run and
history, or pause it from the same view.
Use Go 1.25+, Node.js 22.12+ with npm, and either a logged-in claude , codex or copilot CLI on PATH , or an OpenCode Go key configured in Settings → Subscriptions. The opencode CLI is optional.
make run-web
Open localhost:7717 , add a project folder, and start a session.
Or add one from the terminal with agenttik --init in the folder you want —
agenttik --init path/to/repo names another — which works whether or not the
app is already open, and shows up in an open window straight away.
Remote servers and private workspaces
Connect to another server with agenttik --remote host:7717 (or an HTTPS
URL), or choose Connect to remote server in the command palette. The
client checks /api/version first, then shows the server’s login if needed.
See remote connections .
Start a temporary, separate instance with agenttik --private . Its app data is
removed on exit; project files stay on disk. Manage local profiles in
Settings → Profiles . With multiple profiles, the picker before Shortcuts
switches between their isolated projects and tasks.
--addr and --data-dir change where it listens and where it keeps its
database, --web skips the desktop window, and --help and --version say
what this build is.
For the desktop app, install the native build dependencies ( make deps on
Debian/Ubuntu), then run make run . make build-windows-amd64 cross-compiles
the Windows x64 binary; build the macOS ARM64 target on macOS with
make build-macos-arm64 . macOS builds also create an ad-hoc-signed
agenttik.app and a ZIP ( bin/ for make build , dist/ for the ARM64 target).
Extract the release ZIP and move the app to Applications. No Apple developer
account is needed to build it. Downloaded builds are not notarized; after a
blocked launch, use System Settings → Privacy & Security → Open Anyway.
On macOS, Close to tray works even without Accessibility permission. The global
show/hide shortcut needs that permission; grant it in System Settings → Privacy
& Security → Accessibility, then restart agenttik. Shortcut failures appear in
agenttik's desktop settings while the tray menu remains usable.
Local access by default: agenttik binds to loopback. Anyone who can reach an
unauthenticated server has full control of it. Before exposing it through
Settings → Server , configure authentication and use a trusted network. See
server access .
Built with Go, SQLite, Vue 3, Nuxt UI, and Wails. See specs
for architecture and implementation details.
Manage multiple projects and agents with ease
Readme MIT license Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
