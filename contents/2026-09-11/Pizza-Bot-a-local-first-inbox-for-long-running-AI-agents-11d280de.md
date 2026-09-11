---
source: "https://github.com/pizza-bot-app/pizza-bot"
hn_url: "https://news.ycombinator.com/item?id=49657189"
title: "Pizza Bot: a local-first inbox for long-running AI agents"
article_title: "GitHub - pizza-bot-app/pizza-bot: A local-first inbox for long-running AI agents, built with DeepAgents and LangGraph. · GitHub"
image: "https://opengraph.githubassets.com/b2b44429e1db22e80e66bd70161655e4e8c649d15106b8fb164eaf796a3a3f14/pizza-bot-app/pizza-bot"
author: "joshcsimmons"
captured_at: "2026-09-11T13:18:30Z"
capture_tool: "hn-digest"
hn_id: 49657189
score: 1
comments: 0
posted_at: "2026-09-11T12:20:16Z"
tags:
  - hacker-news
---

# Pizza Bot: a local-first inbox for long-running AI agents

- HN: [49657189](https://news.ycombinator.com/item?id=49657189)
- Source: [github.com](https://github.com/pizza-bot-app/pizza-bot)
- Score: 1
- Comments: 0
- Posted: 2026-09-11T12:20:16Z

## Translation

Title: Pizza Bot: a local-first inbox for long-running AI agents
Article title: GitHub - pizza-bot-app/pizza-bot: A local-first inbox for long-running AI agents, built with DeepAgents and LangGraph. · GitHub
Description: A local-first inbox for long-running AI agents, built with DeepAgents and LangGraph. - pizza-bot-app/pizza-bot

Article text:
GitHub - pizza-bot-app/pizza-bot: A local-first inbox for long-running AI agents, built with DeepAgents and LangGraph. · GitHub
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
pizza-bot-app
/
pizza-bot
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
75 Commits 75 Commits Folders and files
.github .github apps apps assets assets docs docs examples/ plugins examples/ plugins packages packages plugins plugins scripts scripts skills skills tests/ langgraph-compat tests/ langgraph-compat .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json turbo.json turbo.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Pizza Bot is an inbox for long-running AI work. Start or schedule a task, return
to your day, and let completed work collect in Unread while runs waiting for
your decision collect in Action . Agents keep working when you navigate away
or disconnect; the api-server process must remain running.
Pizza Bot uses a stateful DeepAgents/LangGraph runtime with the same React
experience in Electron and the browser. The desktop app, web app, and terminal
CLI all communicate with the api-server over HTTP/SSE.
Pizza Bot was developed at Amazon and is released under the Apache 2.0 license.
Work asynchronously. Switch conversations without stopping their runs.
Return to the right queue. Completed work lands in Unread; durable approval
requests land in Action.
Organize conversations. Group threads into folders without hiding matching
work from the global Unread and Action queues.
Resume real work. Checkpointed runs survive client disconnects, and cron or
webhook triggers can start work without an open conversation.
Delegate to specialists. Skills become tool-scoped workers whose progress
appears in the Activity panel.
Bring your model provider. Amazon Bedrock, Anthropic, Google Gemini,
OpenAI, OpenRouter, and Ollama are supported.
Keep control of consequential actions. Human-in-the-loop approvals,
long-term memory, file attachments, and desktop notifications are built into
the workflow.
Grant local access explicitly. Add individual read-only or writable folders
under Settings > Files ; Pizza Bot receives no default home-directory access.
Installers for macOS (Intel and Apple silicon), Windows, and Linux (x64 and
arm64) are attached to every release , with a SHA256SUMS to
check a download against. The macOS builds are signed and notarized; the Linux
packages are not signed, so verify them against the checksums.
Node.js 24 or newer is required.
npm install
npm run build
npm run dev
npm run dev starts the Vite frontend and Electron desktop shell. The shell
forks and supervises its own api-server, matching the packaged application's
process model. Configure a model under Settings > Providers before starting
a live run.
See Running from source for isolated data roots, browser and
CLI development, desktop packages, and remote backends.
Experience
Best for
Start here
Electron desktop
Local inbox with an embedded backend
npm run dev
Browser
Web development or static deployment
Browser development
Terminal CLI
Scripts, terminals, and remote backends
CLI
Standalone backend
Remote Electron, browsers, containers, or Linux services
Backend guide
A running api-server needs access to at least one model provider; HTTP clients
do not. Configure Amazon Bedrock, Anthropic, Google Gemini, OpenAI, OpenRouter,
or Ollama in Settings > Providers . Bedrock accepts an AWS profile, AWS
access keys, or a Bedrock API key, with an optional region override; otherwise
AWS_REGION or us-west-2 is used. Bedrock combines its native catalog with
the regional Mantle catalog and routes models through Converse, OpenAI
Responses or Chat Completions, or Anthropic Messages according to their
advertised API family.
OpenAI and Anthropic also accept custom base URLs for compatible endpoints;
OpenAI can explicitly select Responses or Chat Completions, and Anthropic
supports x-api-key or bearer authentication. Select a model with
PIZZA_MODEL=<provider>:<id> . The desktop protects entered secrets with
Electron safeStorage ; server configuration persists only environment-variable
references.
Add MCP servers from the UI or <PIZZA_DATA_ROOT>/.mcp.json . Add Agent Skills
under <PIZZA_DATA_ROOT>/skills , or install plugins that package MCP servers and
skills together. Skills become available after their declared tools are enabled
and connected. A custom skill can replace a Built-in or Plugin skill with the
same id without modifying the original; removing the customization reveals the
Built-in or Plugin version again. The Built-in Pizza Bot Guide can explain
features, suggest workflows, help with setup, and point to project documentation.
See Extending Pizza Bot for configuration, environment
references, skill authoring, approval gates, and plugin installation.
apps/ api-server (Hono) | cli | desktop-shell (Electron) | web (React)
packages/ core | runtime-langgraph | inference-providers | plugin-api | plugin-sdk | storage | logging
plugins/ bundled Plugin packages and their packaging workspace
skills/ optional Built-in Agent Skills
tests/ LangGraph compatibility and protocol conformance
The production graph engine is isolated to packages/runtime-langgraph ;
frontends consume protocol projections rather than importing runtime or model
bindings.
Running - desktop, browser, CLI, and package commands.
Extending - MCP servers, skills, and plugins.
Architecture - system boundaries, event model,
persistence, transports, and design decisions.
Standalone backend - authentication, remote
Electron, static browser deployment, Docker, Compose, and Kubernetes.
Contributing - development setup, CI checks, worktrees,
releases, and layering rules.
Security - network defaults, credentials, local data, and
plugin trust.
Logging - diagnostics, retention, viewing, and redaction.
Roadmap - exploratory directions and the principles used to
evaluate them.
Code of Conduct - community participation
expectations.
Local-first by default. The api-server binds to 127.0.0.1 ; non-loopback
binding requires authentication and an explicit origin allowlist.
Application state stays local. Threads, checkpoints, memories,
attachments, and logs live under <PIZZA_DATA_ROOT>
( ~/.pizza-bot-oss by default). Model and tool requests go to the providers
and endpoints you configure.
Local folders require an explicit grant. Each folder added under
Settings > Files is read-only unless you allow writes. Remote grants name
paths on the backend host.
MCP servers and plugins are trusted. Their commands and materializers can
execute with your user account's permissions. Install only sources you trust.
See SECURITY.md for the complete security model and vulnerability
reporting process.
Pizza Bot was designed, built, and brought into the open by its Executive Chefs
and Sous Chefs:
Pizza Bot was also shaped by more than 2,000 users across Amazon who tested
earlier versions and shared feedback from real-world use. Their bug reports,
ideas, and candid input helped make Pizza Bot ready for a broader community.
Thank you to everyone who contributed.
Apache-2.0 . See NOTICE for attribution notices.
A local-first inbox for long-running AI agents, built with DeepAgents and LangGraph.
github.com/pizza-bot-app/pizza-bot/releases Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A local-first inbox for long-running AI agents, built with DeepAgents and LangGraph. - pizza-bot-app/pizza-bot

GitHub - pizza-bot-app/pizza-bot: A local-first inbox for long-running AI agents, built with DeepAgents and LangGraph. · GitHub
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
pizza-bot-app
/
pizza-bot
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
75 Commits 75 Commits Folders and files
.github .github apps apps assets assets docs docs examples/ plugins examples/ plugins packages packages plugins plugins scripts scripts skills skills tests/ langgraph-compat tests/ langgraph-compat .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json turbo.json turbo.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Pizza Bot is an inbox for long-running AI work. Start or schedule a task, return
to your day, and let completed work collect in Unread while runs waiting for
your decision collect in Action . Agents keep working when you navigate away
or disconnect; the api-server process must remain running.
Pizza Bot uses a stateful DeepAgents/LangGraph runtime with the same React
experience in Electron and the browser. The desktop app, web app, and terminal
CLI all communicate with the api-server over HTTP/SSE.
Pizza Bot was developed at Amazon and is released under the Apache 2.0 license.
Work asynchronously. Switch conversations without stopping their runs.
Return to the right queue. Completed work lands in Unread; durable approval
requests land in Action.
Organize conversations. Group threads into folders without hiding matching
work from the global Unread and Action queues.
Resume real work. Checkpointed runs survive client disconnects, and cron or
webhook triggers can start work without an open conversation.
Delegate to specialists. Skills become tool-scoped workers whose progress
appears in the Activity panel.
Bring your model provider. Amazon Bedrock, Anthropic, Google Gemini,
OpenAI, OpenRouter, and Ollama are supported.
Keep control of consequential actions. Human-in-the-loop approvals,
long-term memory, file attachments, and desktop notifications are built into
the workflow.
Grant local access explicitly. Add individual read-only or writable folders
under Settings > Files ; Pizza Bot receives no default home-directory access.
Installers for macOS (Intel and Apple silicon), Windows, and Linux (x64 and
arm64) are attached to every release , with a SHA256SUMS to
check a download against. The macOS builds are signed and notarized; the Linux
packages are not signed, so verify them against the checksums.
Node.js 24 or newer is required.
npm install
npm run build
npm run dev
npm run dev starts the Vite frontend and Electron desktop shell. The shell
forks and supervises its own api-server, matching the packaged application's
process model. Configure a model under Settings > Providers before starting
a live run.
See Running from source for isolated data roots, browser and
CLI development, desktop packages, and remote backends.
Experience
Best for
Start here
Electron desktop
Local inbox with an embedded backend
npm run dev
Browser
Web development or static deployment
Browser development
Terminal CLI
Scripts, terminals, and remote backends
CLI
Standalone backend
Remote Electron, browsers, containers, or Linux services
Backend guide
A running api-server needs access to at least one model provider; HTTP clients
do not. Configure Amazon Bedrock, Anthropic, Google Gemini, OpenAI, OpenRouter,
or Ollama in Settings > Providers . Bedrock accepts an AWS profile, AWS
access keys, or a Bedrock API key, with an optional region override; otherwise
AWS_REGION or us-west-2 is used. Bedrock combines its native catalog with
the regional Mantle catalog and routes models through Converse, OpenAI
Responses or Chat Completions, or Anthropic Messages according to their
advertised API family.
OpenAI and Anthropic also accept custom base URLs for compatible endpoints;
OpenAI can explicitly select Responses or Chat Completions, and Anthropic
supports x-api-key or bearer authentication. Select a model with
PIZZA_MODEL=<provider>:<id> . The desktop protects entered secrets with
Electron safeStorage ; server configuration persists only environment-variable
references.
Add MCP servers from the UI or <PIZZA_DATA_ROOT>/.mcp.json . Add Agent Skills
under <PIZZA_DATA_ROOT>/skills , or install plugins that package MCP servers and
skills together. Skills become available after their declared tools are enabled
and connected. A custom skill can replace a Built-in or Plugin skill with the
same id without modifying the original; removing the customization reveals the
Built-in or Plugin version again. The Built-in Pizza Bot Guide can explain
features, suggest workflows, help with setup, and point to project documentation.
See Extending Pizza Bot for configuration, environment
references, skill authoring, approval gates, and plugin installation.
apps/ api-server (Hono) | cli | desktop-shell (Electron) | web (React)
packages/ core | runtime-langgraph | inference-providers | plugin-api | plugin-sdk | storage | logging
plugins/ bundled Plugin packages and their packaging workspace
skills/ optional Built-in Agent Skills
tests/ LangGraph compatibility and protocol conformance
The production graph engine is isolated to packages/runtime-langgraph ;
frontends consume protocol projections rather than importing runtime or model
bindings.
Running - desktop, browser, CLI, and package commands.
Extending - MCP servers, skills, and plugins.
Architecture - system boundaries, event model,
persistence, transports, and design decisions.
Standalone backend - authentication, remote
Electron, static browser deployment, Docker, Compose, and Kubernetes.
Contributing - development setup, CI checks, worktrees,
releases, and layering rules.
Security - network defaults, credentials, local data, and
plugin trust.
Logging - diagnostics, retention, viewing, and redaction.
Roadmap - exploratory directions and the principles used to
evaluate them.
Code of Conduct - community participation
expectations.
Local-first by default. The api-server binds to 127.0.0.1 ; non-loopback
binding requires authentication and an explicit origin allowlist.
Application state stays local. Threads, checkpoints, memories,
attachments, and logs live under <PIZZA_DATA_ROOT>
( ~/.pizza-bot-oss by default). Model and tool requests go to the providers
and endpoints you configure.
Local folders require an explicit grant. Each folder added under
Settings > Files is read-only unless you allow writes. Remote grants name
paths on the backend host.
MCP servers and plugins are trusted. Their commands and materializers can
execute with your user account's permissions. Install only sources you trust.
See SECURITY.md for the complete security model and vulnerability
reporting process.
Pizza Bot was designed, built, and brought into the open by its Executive Chefs
and Sous Chefs:
Pizza Bot was also shaped by more than 2,000 users across Amazon who tested
earlier versions and shared feedback from real-world use. Their bug reports,
ideas, and candid input helped make Pizza Bot ready for a broader community.
Thank you to everyone who contributed.
Apache-2.0 . See NOTICE for attribution notices.
A local-first inbox for long-running AI agents, built with DeepAgents and LangGraph.
github.com/pizza-bot-app/pizza-bot/releases Topics
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
