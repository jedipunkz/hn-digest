---
source: "https://github.com/plurnk/plurnk"
hn_url: "https://news.ycombinator.com/item?id=49720466"
title: "Show HN: Plurnk (Yet *Another* AI Harness)"
article_title: "GitHub - plurnk/plurnk: Plurnk Client - Pattern Lookup Universal Resource NetworK · GitHub"
image: "https://repository-images.githubusercontent.com/1241322899/91a85d43-3fd2-45ca-aa40-f1d0d3611c1d"
author: "wikitopian"
captured_at: "2026-09-15T23:59:42Z"
capture_tool: "hn-digest"
hn_id: 49720466
score: 1
comments: 0
posted_at: "2026-09-15T23:57:14Z"
tags:
  - hacker-news
---

# Show HN: Plurnk (Yet *Another* AI Harness)

- HN: [49720466](https://news.ycombinator.com/item?id=49720466)
- Source: [github.com](https://github.com/plurnk/plurnk)
- Score: 1
- Comments: 0
- Posted: 2026-09-15T23:57:14Z

## Translation

Title: Show HN: Plurnk (Yet *Another* AI Harness)
Article title: GitHub - plurnk/plurnk: Plurnk Client - Pattern Lookup Universal Resource NetworK · GitHub
Description: Plurnk Client - Pattern Lookup Universal Resource NetworK - plurnk/plurnk
HN text: I am desperate for feedback on my homebrewed harness, "Plurnk." I've been dogfooding it on my RTX5070Ti with only 16GB VRAM as my daily driver for weeks now and it's genuinely better than everything else available (I'm biased). 1. Curation, not Compaction The model's fully responsible (bitter lesson) for deterministically curating its own context. 2. ANTLR Grammar The dozen verbs for the harness are optimized for model pretraining, looking like markdown while being fully integrated with an EBNF grammar that plugs into an AST for superior workflow and recovery. 3. Universal Resources Everything's addressable by the model through a pseudo-URI interface, including the log entries. 4. Omnipatterns Everything in the repo goes through treesitters which build a complete graph that the model can search with glob, regex, jsonpath, xpath, graph, and sqlite fulltext. 5. Standards MCP2 tools, AG-UI client interface, A2A "agent to agent", plugs into everything through OAI spec + models.dev 6. Interfaces CLI, TUI, and a Neovim plugin, easy to extend, very pluggable 7. Lean Agent Service/Client architecture with lean TUI allows dozens of agents to run on weak hardware. 200MB footprint. 8. Lean Context Achieves all of the "batteries includes" features of a fat harness while the sysprompt is only slightly larger than pi agent's. It started as a proof of concept for my bespoke opinions on what "bitter lesson" actually means, and it's gone so well that I believe others trying to build agentic workflows, especially with local and humble constraints, could benefit from the proven concepts.

Article text:
GitHub - plurnk/plurnk: Plurnk Client - Pattern Lookup Universal Resource NetworK · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
plurnk
/
plurnk
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
496 Commits 496 Commits Folders and files
.githooks .githooks bin bin completions completions conformance conformance man man scripts scripts src src test test .env.defaults .env.defaults .gitignore .gitignore .nvmrc .nvmrc AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md SPEC.md SPEC.md TUI.md TUI.md package-lock.json package-lock.json package.json package.json pid pid tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json View all files Repository files navigation
The Bitter Lesson applies to the harness, too.
The model should decide what to remember, where to look, when to delegate, and
how to proceed. Plurnk is an agentic operating system and programming language
that puts those decisions in the model's hands.
Files, tools, and the agent's own context become an addressable environment.
Composable operations let it search that environment, make precise changes,
curate its memory, and build its own delegation topology. The runtime provides
reliable machinery; the model supplies the strategy.
Use local or cloud models to build software, investigate a codebase, or automate
work across tools. Interact from your terminal or editor, or compose Plurnk
with ordinary shell pipelines.
This repository provides the CLI and interactive terminal client for
plurnk-service , the shared daemon.
Curation, not compaction. The agent retrieves the passages it needs and
removes stale items or individual lines from its active context. Its working
set changes; source material and original execution evidence survive.
Precision without ceremony. Line ranges, character regions, and
hash-anchored edits make surgical changes possible. Stale anchors reject
conflicting edits before they overwrite the wrong text.
The model chooses the topology. Fork with existing context, start a
worker with a fresh log, or delegate pure inference without an agent loop.
Parent and child models can use different endpoints: a cloud model can
orchestrate local workers through the same primitives.
Execution with evidence. ANTLR parses model output into executable
operations. The runtime records their results and returns structured errors
the model can act on. Full packet digests make the work inspectable. A
compatible local server can run under a grammar you wrote, carried verbatim.
Workspaces and worker conversations live in the daemon, independently of the
client session. Choose your model, context limits, tools, and capability
policies without replacing the environment. No Plurnk account is required.
The pattern engine connects discovery and context management. Path globs
combine with full-text search, regex, JSONPath, XPath, and symbol-graph queries.
The model can search for phrases, inspect structured data, or follow symbol
relationships without writing a script for each question.
For example, these model-side operations find TypeScript files matching
retry or timeout , then trim older READ receipts to their first 16 lines:
### FIND_ (src/**/*.{ts,tsx})
~retry OR timeout
### KILL_ (log:///1/[1-7]/*/READ) <17,-1>
The second operation targets READ results from turns 1–7 of loop 1. It curates
the log, not the source files. One expression can manage many entries: the
agent has bulk operations over its own context, not just over your code.
Requires Node.js 26+, npm, Git, and a local or cloud model endpoint.
npm install -g @plurnk/plurnk @plurnk/plurnk-service
In one terminal, configure a model and start the daemon. This example uses
DeepSeek; see model configuration
for other providers and local servers.
export DEEPSEEK_API_KEY= " your-api-key "
export PLURNK_MODEL=deepseek/deepseek-v4-flash
plurnk-service start
In another terminal, open a project:
cd /path/to/your/project
plurnk --workspace= " myProject "
Give it a task in ordinary language. The model uses the operation language;
you do not need to learn it to use Plurnk. Run the same command later to return
to that workspace's conversation.
The client connects to 127.0.0.1:1066 by default and never starts the daemon.
Provider credentials belong in the daemon's environment. Proposals are accepted
automatically by default ( PLURNK_CLIENT_YOLO=1 ); set it to 0 or start a prompt
with ? to review them. Neither overrides capability restrictions.
plurnk --workspace= " myProject "
A scrollback-native TUI with multiline prompts, streaming reasoning when the
provider supplies it, Markdown and Mermaid rendering, and slash commands for
managing models, workers, and tools. Run /help to explore; /model and
/child select the conversation and delegated models.
plurnk " Explain how this project's request handling works "
plurnk " Summarize this repository " > overview.md
git diff | plurnk " Review this patch for correctness "
plurnk --json " Explain the test layout " | jq -r .response
One-shot commands put the answer on stdout and progress on stderr. --json
returns one structured document containing the answer, operation trace,
diagnostics, and usage.
Use plurnk --help for CLI options. plurnk models lists available model
routes.
For an editor-native interface, use plurnk.nvim
against the same daemon.
Plurnk uses AG-UI for clients, MCP for external tools, Agent Skills
for reusable instructions, and A2A for remote agents. Tools, skills, and
agents can be discovered and enabled per worker without restarting the daemon.
Their documentation is retrieved on demand rather than loading every tool's
schema into every prompt.
Configuration uses cascading environment variables and .env files, including
the XDG user configuration at ~/.config/plurnk/.env . Model discovery uses
models.dev . Inspect the complete, documented configuration
catalog with:
plurnk-service config defaults
Documentation
Installation and configuration
Architecture and extension points
Client reference and terminal design
Questions, bug reports, and feedback are welcome in
GitHub issues . See the shared
contributing guide
for development guidance.
Plurnk Client - Pattern Lookup Universal Resource NetworK
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Plurnk Client - Pattern Lookup Universal Resource NetworK - plurnk/plurnk

I am desperate for feedback on my homebrewed harness, "Plurnk." I've been dogfooding it on my RTX5070Ti with only 16GB VRAM as my daily driver for weeks now and it's genuinely better than everything else available (I'm biased). 1. Curation, not Compaction The model's fully responsible (bitter lesson) for deterministically curating its own context. 2. ANTLR Grammar The dozen verbs for the harness are optimized for model pretraining, looking like markdown while being fully integrated with an EBNF grammar that plugs into an AST for superior workflow and recovery. 3. Universal Resources Everything's addressable by the model through a pseudo-URI interface, including the log entries. 4. Omnipatterns Everything in the repo goes through treesitters which build a complete graph that the model can search with glob, regex, jsonpath, xpath, graph, and sqlite fulltext. 5. Standards MCP2 tools, AG-UI client interface, A2A "agent to agent", plugs into everything through OAI spec + models.dev 6. Interfaces CLI, TUI, and a Neovim plugin, easy to extend, very pluggable 7. Lean Agent Service/Client architecture with lean TUI allows dozens of agents to run on weak hardware. 200MB footprint. 8. Lean Context Achieves all of the "batteries includes" features of a fat harness while the sysprompt is only slightly larger than pi agent's. It started as a proof of concept for my bespoke opinions on what "bitter lesson" actually means, and it's gone so well that I believe others trying to build agentic workflows, especially with local and humble constraints, could benefit from the proven concepts.

GitHub - plurnk/plurnk: Plurnk Client - Pattern Lookup Universal Resource NetworK · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
plurnk
/
plurnk
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
496 Commits 496 Commits Folders and files
.githooks .githooks bin bin completions completions conformance conformance man man scripts scripts src src test test .env.defaults .env.defaults .gitignore .gitignore .nvmrc .nvmrc AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md SPEC.md SPEC.md TUI.md TUI.md package-lock.json package-lock.json package.json package.json pid pid tsconfig.build.json tsconfig.build.json tsconfig.json tsconfig.json View all files Repository files navigation
The Bitter Lesson applies to the harness, too.
The model should decide what to remember, where to look, when to delegate, and
how to proceed. Plurnk is an agentic operating system and programming language
that puts those decisions in the model's hands.
Files, tools, and the agent's own context become an addressable environment.
Composable operations let it search that environment, make precise changes,
curate its memory, and build its own delegation topology. The runtime provides
reliable machinery; the model supplies the strategy.
Use local or cloud models to build software, investigate a codebase, or automate
work across tools. Interact from your terminal or editor, or compose Plurnk
with ordinary shell pipelines.
This repository provides the CLI and interactive terminal client for
plurnk-service , the shared daemon.
Curation, not compaction. The agent retrieves the passages it needs and
removes stale items or individual lines from its active context. Its working
set changes; source material and original execution evidence survive.
Precision without ceremony. Line ranges, character regions, and
hash-anchored edits make surgical changes possible. Stale anchors reject
conflicting edits before they overwrite the wrong text.
The model chooses the topology. Fork with existing context, start a
worker with a fresh log, or delegate pure inference without an agent loop.
Parent and child models can use different endpoints: a cloud model can
orchestrate local workers through the same primitives.
Execution with evidence. ANTLR parses model output into executable
operations. The runtime records their results and returns structured errors
the model can act on. Full packet digests make the work inspectable. A
compatible local server can run under a grammar you wrote, carried verbatim.
Workspaces and worker conversations live in the daemon, independently of the
client session. Choose your model, context limits, tools, and capability
policies without replacing the environment. No Plurnk account is required.
The pattern engine connects discovery and context management. Path globs
combine with full-text search, regex, JSONPath, XPath, and symbol-graph queries.
The model can search for phrases, inspect structured data, or follow symbol
relationships without writing a script for each question.
For example, these model-side operations find TypeScript files matching
retry or timeout , then trim older READ receipts to their first 16 lines:
### FIND_ (src/**/*.{ts,tsx})
~retry OR timeout
### KILL_ (log:///1/[1-7]/*/READ) <17,-1>
The second operation targets READ results from turns 1–7 of loop 1. It curates
the log, not the source files. One expression can manage many entries: the
agent has bulk operations over its own context, not just over your code.
Requires Node.js 26+, npm, Git, and a local or cloud model endpoint.
npm install -g @plurnk/plurnk @plurnk/plurnk-service
In one terminal, configure a model and start the daemon. This example uses
DeepSeek; see model configuration
for other providers and local servers.
export DEEPSEEK_API_KEY= " your-api-key "
export PLURNK_MODEL=deepseek/deepseek-v4-flash
plurnk-service start
In another terminal, open a project:
cd /path/to/your/project
plurnk --workspace= " myProject "
Give it a task in ordinary language. The model uses the operation language;
you do not need to learn it to use Plurnk. Run the same command later to return
to that workspace's conversation.
The client connects to 127.0.0.1:1066 by default and never starts the daemon.
Provider credentials belong in the daemon's environment. Proposals are accepted
automatically by default ( PLURNK_CLIENT_YOLO=1 ); set it to 0 or start a prompt
with ? to review them. Neither overrides capability restrictions.
plurnk --workspace= " myProject "
A scrollback-native TUI with multiline prompts, streaming reasoning when the
provider supplies it, Markdown and Mermaid rendering, and slash commands for
managing models, workers, and tools. Run /help to explore; /model and
/child select the conversation and delegated models.
plurnk " Explain how this project's request handling works "
plurnk " Summarize this repository " > overview.md
git diff | plurnk " Review this patch for correctness "
plurnk --json " Explain the test layout " | jq -r .response
One-shot commands put the answer on stdout and progress on stderr. --json
returns one structured document containing the answer, operation trace,
diagnostics, and usage.
Use plurnk --help for CLI options. plurnk models lists available model
routes.
For an editor-native interface, use plurnk.nvim
against the same daemon.
Plurnk uses AG-UI for clients, MCP for external tools, Agent Skills
for reusable instructions, and A2A for remote agents. Tools, skills, and
agents can be discovered and enabled per worker without restarting the daemon.
Their documentation is retrieved on demand rather than loading every tool's
schema into every prompt.
Configuration uses cascading environment variables and .env files, including
the XDG user configuration at ~/.config/plurnk/.env . Model discovery uses
models.dev . Inspect the complete, documented configuration
catalog with:
plurnk-service config defaults
Documentation
Installation and configuration
Architecture and extension points
Client reference and terminal design
Questions, bug reports, and feedback are welcome in
GitHub issues . See the shared
contributing guide
for development guidance.
Plurnk Client - Pattern Lookup Universal Resource NetworK
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
