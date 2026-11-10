---
source: "https://github.com/Eigin-AI/Eigin"
hn_url: "https://news.ycombinator.com/item?id=50439089"
title: "Show HN: Eigin. Local-first and no-shell personal AI for iOS"
article_title: "GitHub - Eigin-AI/Eigin: Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in. · GitHub"
image: "https://opengraph.githubassets.com/aa8f01dd537cd3b979b834c0fdef27cfeb0915d629e96031a4e8f70e88ad8772/Eigin-AI/Eigin"
author: "danielrothmann"
captured_at: "2026-11-10T12:11:27Z"
capture_tool: "hn-digest"
hn_id: 50439089
score: 1
comments: 0
posted_at: "2026-11-10T12:00:34Z"
tags:
  - hacker-news
---

# Show HN: Eigin. Local-first and no-shell personal AI for iOS

- HN: [50439089](https://news.ycombinator.com/item?id=50439089)
- Source: [github.com](https://github.com/Eigin-AI/Eigin)
- Score: 1
- Comments: 0
- Posted: 2026-11-10T12:00:34Z

## Translation

Title: Show HN: Eigin. Local-first and no-shell personal AI for iOS
Article title: GitHub - Eigin-AI/Eigin: Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in. · GitHub
Description: Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in. - Eigin-AI/Eigin
HN text: Hey HN, we're Daniel and Therkel and we're working on Eigin, a personal agent built to be private, safe and under your control. What makes Eigin different from the mainstream personal agents is it's local-first architecture (chats, memory, agent runtime on your device), closed-set tool model (fixed set of tools, permissions in execution path) and it being open source under AGPL. Personal agents are exciting as an emerging category. We believe that technology this personal should belong to the people using it. More info at eigin.ai It's in open beta, and we'd be very happy to hear your thoughts.

Article text:
GitHub - Eigin-AI/Eigin: Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in. · GitHub
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
Eigin-AI
Eigin Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in.
Readme AGPL-3.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
700 Commits 700 Commits Folders and files
.buildkite .buildkite .gitlab .gitlab .pi .pi .zed .zed LICENSES LICENSES config config design design docs docs infra/ scaleway infra/ scaleway src src tools tools .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md cliff.toml cliff.toml justfile justfile opencode.json opencode.json package-lock.json package-lock.json package.json package.json View all files Repository files navigation
The open application layer for personal AI
Eigin is a personal AI agent built to be private, safe and under your control.
What would you tell an AI, if you knew it was truly yours?
Our places for reflection used to be private. Now we think out loud with AI, and every word lands on a corporate server.
Technology this personal should belong to the people who use it. That's why we're building Eigin.
Available on iOS. See what's new in the changelog .
Personal context. An on-device knowledge graph of your people, goals and routines, with an editable agent character.
Model choice. Use our European cloud AI, connect an OpenAI-compatible provider or run a local model. Your agent and memory stay with you.
Connected tools. Native calendar, reminders and health, plus a library of useful integrations and MCP.
Background tasks. Goals tracked as intents, with scheduled sessions within the permissions you set.
Local voice. On-device speech recognition and text-to-speech, with a hands-free call mode.
Eigin is local-first. Its agent runtime, memory and tools run on your device. Chats, memories and settings are stored locally. Inference runs locally, through Eigin's private cloud AI or your chosen AI provider.
Eigin is a closed-set agent. Its runtime has a fixed set of declared tools, with no shell or virtual computer. It asks before changing your data or acting on your behalf.
Eigin is yours. Edit its character and memory, connect your own tools and choose your models. The code is open source, so you can change the application itself and run your own version.
Read about the architecture and security and compliance .
For the Rust SDK, browser app, and headless CLI, see the SDK development guide .
Install Rust and just . For all-target generation, also install Node.js 22.17+ and npm 10+. For iOS, you also need macOS, Xcode 26 with the iOS 26 SDK, XcodeGen and SwiftFormat .
just hooks # install git hooks (one-time)
just config init # create your .env
just generate # run code generation
just run ios simulator # build and launch in the simulator
See the build guide for setup details.
Open a bug report or feature request .
Join the conversation on Matrix .
Report vulnerabilities through SECURITY.md .
How is Eigin funded? The app is free for personal use, funded by optional cloud AI and commercial plans. Local models and your own provider are supported.
Can I self-host? The app runs on your device. The optional relay ships as a Docker image with Terraform configuration for deployment.
What if the terms change? Code released under AGPL stays available under that license. You can fork it and run your own version.
Eigin is free and open source under AGPL v3 . Third-party attributions are in LICENSES/ .
Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in.
Readme AGPL-3.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository Contributors
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in. - Eigin-AI/Eigin

Hey HN, we're Daniel and Therkel and we're working on Eigin, a personal agent built to be private, safe and under your control. What makes Eigin different from the mainstream personal agents is it's local-first architecture (chats, memory, agent runtime on your device), closed-set tool model (fixed set of tools, permissions in execution path) and it being open source under AGPL. Personal agents are exciting as an emerging category. We believe that technology this personal should belong to the people using it. More info at eigin.ai It's in open beta, and we'd be very happy to hear your thoughts.

GitHub - Eigin-AI/Eigin: Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in. · GitHub
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
Eigin-AI
Eigin Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in.
Readme AGPL-3.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
700 Commits 700 Commits Folders and files
.buildkite .buildkite .gitlab .gitlab .pi .pi .zed .zed LICENSES LICENSES config config design design docs docs infra/ scaleway infra/ scaleway src src tools tools .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .worktreeinclude .worktreeinclude AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md cliff.toml cliff.toml justfile justfile opencode.json opencode.json package-lock.json package-lock.json package.json package.json View all files Repository files navigation
The open application layer for personal AI
Eigin is a personal AI agent built to be private, safe and under your control.
What would you tell an AI, if you knew it was truly yours?
Our places for reflection used to be private. Now we think out loud with AI, and every word lands on a corporate server.
Technology this personal should belong to the people who use it. That's why we're building Eigin.
Available on iOS. See what's new in the changelog .
Personal context. An on-device knowledge graph of your people, goals and routines, with an editable agent character.
Model choice. Use our European cloud AI, connect an OpenAI-compatible provider or run a local model. Your agent and memory stay with you.
Connected tools. Native calendar, reminders and health, plus a library of useful integrations and MCP.
Background tasks. Goals tracked as intents, with scheduled sessions within the permissions you set.
Local voice. On-device speech recognition and text-to-speech, with a hands-free call mode.
Eigin is local-first. Its agent runtime, memory and tools run on your device. Chats, memories and settings are stored locally. Inference runs locally, through Eigin's private cloud AI or your chosen AI provider.
Eigin is a closed-set agent. Its runtime has a fixed set of declared tools, with no shell or virtual computer. It asks before changing your data or acting on your behalf.
Eigin is yours. Edit its character and memory, connect your own tools and choose your models. The code is open source, so you can change the application itself and run your own version.
Read about the architecture and security and compliance .
For the Rust SDK, browser app, and headless CLI, see the SDK development guide .
Install Rust and just . For all-target generation, also install Node.js 22.17+ and npm 10+. For iOS, you also need macOS, Xcode 26 with the iOS 26 SDK, XcodeGen and SwiftFormat .
just hooks # install git hooks (one-time)
just config init # create your .env
just generate # run code generation
just run ios simulator # build and launch in the simulator
See the build guide for setup details.
Open a bug report or feature request .
Join the conversation on Matrix .
Report vulnerabilities through SECURITY.md .
How is Eigin funded? The app is free for personal use, funded by optional cloud AI and commercial plans. Local models and your own provider are supported.
Can I self-host? The app runs on your device. The optional relay ships as a Docker image with Terraform configuration for deployment.
What if the terms change? Code released under AGPL stays available under that license. You can fork it and run your own version.
Eigin is free and open source under AGPL v3 . Third-party attributions are in LICENSES/ .
Open application layer for personal AI. This is a read-only mirror. Find us at code.eigin.ai to join in.
Readme AGPL-3.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository Contributors
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
