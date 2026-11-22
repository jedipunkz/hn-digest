---
source: "https://github.com/MiniMax-AI/OpenAgentCore"
hn_url: "https://news.ycombinator.com/item?id=50604001"
title: "OpenAgentCore: The OpenAI Agents API, open source and self-hosted"
article_title: "GitHub - MiniMax-AI/OpenAgentCore: The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK. · GitHub"
image: "https://repository-images.githubusercontent.com/1379023782/d9f6d7eb-fdbd-43b8-9693-b59e5696ed27"
author: "gfortaine"
captured_at: "2026-11-22T21:39:50Z"
capture_tool: "hn-digest"
hn_id: 50604001
score: 1
comments: 0
posted_at: "2026-11-22T21:18:43Z"
tags:
  - hacker-news
---

# OpenAgentCore: The OpenAI Agents API, open source and self-hosted

- HN: [50604001](https://news.ycombinator.com/item?id=50604001)
- Source: [github.com](https://github.com/MiniMax-AI/OpenAgentCore)
- Score: 1
- Comments: 0
- Posted: 2026-11-22T21:18:43Z

## Translation

Title: OpenAgentCore: The OpenAI Agents API, open source and self-hosted
Article title: GitHub - MiniMax-AI/OpenAgentCore: The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK. · GitHub
Description: The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK. - MiniMax-AI/OpenAgentCore

Article text:
GitHub - MiniMax-AI/OpenAgentCore: The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK. · GitHub
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
MiniMax-AI
OpenAgentCore Public Notifications You must be signed in to change notification settings
Star 346 ( 346 ) You must be signed in to star a repository
The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
44 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,074 Commits 1,074 Commits Folders and files
.github .github apps apps contracts/ agents-api contracts/ agents-api deploy deploy docs docs example example internal internal packages packages scripts scripts services services website website .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md README.zh-CN.md README.zh-CN.md docs.json docs.json go.mod go.mod go.sum go.sum go.work go.work package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK.
Website · Documentation · Quickstart · Contributing
Same API as OpenAI. Set OPENAI_BASE_URL and the official OpenAI SDK , or plain HTTP, works unchanged.
Your agent, your model. Each Session runs Codex, Claude Code or MiniMax Code natively, as it runs on your machine, with the model provider you configure. See Harness capabilities .
Compute that sleeps, workspaces that stay. On microsandbox, which keeps every workspace on independent storage, idle sandboxes release compute and resume on a compatible node with their files intact. E2B pauses idle sandboxes in place. Your own Linux, macOS or Windows machine also runs Sessions. See Sandbox Provider protocol .
Replaceable parts. Sandbox Providers, workspace filesystems, Harnesses and model providers plug in through defined protocols .
Install Core, Web and PostgreSQL on Linux or macOS with Docker :
curl -fsSL https://openagentcore.dev/install.sh | bash
On Windows, in PowerShell:
irm https: // github.com / MiniMax - AI / OpenAgentCore / releases / latest / download / install.ps1 | iex
In Web, sign in with the Core key the installer prints, configure the public address, set a default model, issue a Project API key and add sandbox capacity (a node, E2B or OpenSandbox). Install Core and Web covers each step.
Then call it with the OpenAI SDK, with OPENAI_BASE_URL set to your API base URL ( https://core.example/v1 ) and OPENAI_API_KEY to the Project API key:
from openai import OpenAI
client = OpenAI () # reads OPENAI_BASE_URL and OPENAI_API_KEY
session = client . beta . agents . sessions . create (
environment = { "type" : "openai_hosted" },
input = "Create /workspace/hello.txt with a short greeting, then describe it." ,
extra_body = { "agent" : { "x_agents_core" : { "harness" : "codex" }}},
)
print ( session . id )
Run your first Session waits for the result and reads it back.
Applications and operators use these Core APIs:
Core keeps durable execution state. The Runtime runs the chosen Harness inside the Environment. Each connection is a defined protocol, so any part can be replaced on its own. See Architecture .
I want to
Start with
Install and operate an installation
Install Core and Web , then Operations
Build an application on the API
Run your first Session , then the Agents API guide
See a complete application
Examples
Run agents on my own machine
Run self-hosted Sessions
Check Harness capabilities and limits
Harness capabilities
Understand the design
Architecture
Add a sandbox, Harness or other component
Develop OpenAgentCore
Contributing and license
Read CONTRIBUTING.md before changing code. Report bugs and ask questions in GitHub Issues . Released under the MIT License .
The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
44 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK. - MiniMax-AI/OpenAgentCore

GitHub - MiniMax-AI/OpenAgentCore: The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK. · GitHub
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
MiniMax-AI
OpenAgentCore Public Notifications You must be signed in to change notification settings
Star 346 ( 346 ) You must be signed in to star a repository
The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
44 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,074 Commits 1,074 Commits Folders and files
.github .github apps apps contracts/ agents-api contracts/ agents-api deploy deploy docs docs example example internal internal packages packages scripts scripts services services website website .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md README.zh-CN.md README.zh-CN.md docs.json docs.json go.mod go.mod go.sum go.sum go.work go.work package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK.
Website · Documentation · Quickstart · Contributing
Same API as OpenAI. Set OPENAI_BASE_URL and the official OpenAI SDK , or plain HTTP, works unchanged.
Your agent, your model. Each Session runs Codex, Claude Code or MiniMax Code natively, as it runs on your machine, with the model provider you configure. See Harness capabilities .
Compute that sleeps, workspaces that stay. On microsandbox, which keeps every workspace on independent storage, idle sandboxes release compute and resume on a compatible node with their files intact. E2B pauses idle sandboxes in place. Your own Linux, macOS or Windows machine also runs Sessions. See Sandbox Provider protocol .
Replaceable parts. Sandbox Providers, workspace filesystems, Harnesses and model providers plug in through defined protocols .
Install Core, Web and PostgreSQL on Linux or macOS with Docker :
curl -fsSL https://openagentcore.dev/install.sh | bash
On Windows, in PowerShell:
irm https: // github.com / MiniMax - AI / OpenAgentCore / releases / latest / download / install.ps1 | iex
In Web, sign in with the Core key the installer prints, configure the public address, set a default model, issue a Project API key and add sandbox capacity (a node, E2B or OpenSandbox). Install Core and Web covers each step.
Then call it with the OpenAI SDK, with OPENAI_BASE_URL set to your API base URL ( https://core.example/v1 ) and OPENAI_API_KEY to the Project API key:
from openai import OpenAI
client = OpenAI () # reads OPENAI_BASE_URL and OPENAI_API_KEY
session = client . beta . agents . sessions . create (
environment = { "type" : "openai_hosted" },
input = "Create /workspace/hello.txt with a short greeting, then describe it." ,
extra_body = { "agent" : { "x_agents_core" : { "harness" : "codex" }}},
)
print ( session . id )
Run your first Session waits for the result and reads it back.
Applications and operators use these Core APIs:
Core keeps durable execution state. The Runtime runs the chosen Harness inside the Environment. Each connection is a defined protocol, so any part can be replaced on its own. See Architecture .
I want to
Start with
Install and operate an installation
Install Core and Web , then Operations
Build an application on the API
Run your first Session , then the Agents API guide
See a complete application
Examples
Run agents on my own machine
Run self-hosted Sessions
Check Harness capabilities and limits
Harness capabilities
Understand the design
Architecture
Add a sandbox, Harness or other component
Develop OpenAgentCore
Contributing and license
Read CONTRIBUTING.md before changing code. Report bugs and ask questions in GitHub Issues . Released under the MIT License .
The OpenAI Agents API, open source and self-hosted. Run Codex, Claude Code or MiniMax Code on your own machines with your own models, through the official OpenAI SDK.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
44 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
