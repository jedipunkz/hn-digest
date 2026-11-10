---
source: "https://github.com/DiegoZaluski/co-sh"
hn_url: "https://news.ycombinator.com/item?id=50439141"
title: "Show HN: Cosh – An open-source AI agent orchestrator"
article_title: "GitHub - DiegoZaluski/co-sh: Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use. · GitHub"
image: "https://repository-images.githubusercontent.com/1256090475/d05da64e-dbd1-461e-9a48-3085e0e1b6ea"
author: "Zaluski"
captured_at: "2026-11-10T12:11:23Z"
capture_tool: "hn-digest"
hn_id: 50439141
score: 1
comments: 0
posted_at: "2026-11-10T12:05:13Z"
tags:
  - hacker-news
---

# Show HN: Cosh – An open-source AI agent orchestrator

- HN: [50439141](https://news.ycombinator.com/item?id=50439141)
- Source: [github.com](https://github.com/DiegoZaluski/co-sh)
- Score: 1
- Comments: 0
- Posted: 2026-11-10T12:05:13Z

## Translation

Title: Show HN: Cosh – An open-source AI agent orchestrator
Article title: GitHub - DiegoZaluski/co-sh: Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use. · GitHub
Description: Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use. - DiegoZaluski/co-sh
HN text: Cosh is an open-source AI agent built in Rust that brings together multiple agents like Codex and Goose in a unified environment for orchestration and collaboration.

Article text:
GitHub - DiegoZaluski/co-sh: Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use. · GitHub
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
DiegoZaluski
co-sh Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use.
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,273 Commits 1,273 Commits Folders and files
.github/ workflows .github/ workflows assets assets crates crates docs docs src src .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE NOTICE NOTICE README.md README.md download.ps1 download.ps1 download.sh download.sh View all files Repository files navigation
English | Português | Español | Русский | 中文 | 日本語
Cosh is a general-purpose coding agent for the terminal, with software engineering tools and computer tools that let it interact with graphical applications.
Multi-agent mode ( /agents ) — run external agents (Gemini, Claude, Codex, Cline, …) alongside the main agent in split panes, all reachable from the same prompt. Each agent keeps its own session, model and optional git worktree, and the main agent coordinates them — dispatching tasks and collecting reports as one team.
Computer tools via accessibility tree — the agent reads and operates graphical interfaces through a structured, compact representation, falling back to screenshots only when needed. Fewer images processed, fewer tokens consumed.
TUI built from scratch in Rust (Ratatui), inspired by OpenCode's UX. Providers, models and operating modes are configured right in the interface — almost zero manual setup.
Independent crates — the development tools live in crates decoupled from the application, reusable as libraries outside Cosh.
Tested on Linux and Windows. macOS hasn't been tested yet — if anything fails, open an issue with as much detail as possible.
Linux & macOS (also works on Windows via Git Bash / WSL):
curl -fsSL https://raw.githubusercontent.com/DiegoZaluski/co-sh/main/download.sh | bash
Windows (PowerShell):
irm https: // raw.githubusercontent.com / DiegoZaluski / co - sh / main / download.ps1 | iex
Note: the curl | bash command does not work in PowerShell or CMD — Windows has no native bash . Use the PowerShell script above instead (Git Bash and WSL also work).
Optional variables: COSH_BIN_DIR (install directory), COSH_VERSION (specific version) and COSH_VARIANT ( slim , the default, or slim-embed , with local RAG via fastembed). On PowerShell, set them as $env:COSH_VERSION = "v0.1.0" before running the command.
After installing, run cosh in your terminal. If the install directory isn't on your PATH , the script prints the command to add it. If anything goes wrong during installation, open an issue — I'll be glad to help!
Crate
Description
cosh-tools
Development tools and computer tools
cosh-tui
TUI widget library (markdown/diff rendering, layout) — the components the app's terminal interface is built on
cosh-sdk
SDK for building agents on top of cosh
cosh-recall
Memory and context recall
cosh-onnx
ONNX decision engine (behind the onnx feature)
Roadmap
Laya/ONNX decision engine — optional onnx feature backed by the cosh-onnx crate; see the crate guide .
Graph-based harness orchestration — a new session mode for coordinating multiple harnesses.
Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use. - DiegoZaluski/co-sh

Cosh is an open-source AI agent built in Rust that brings together multiple agents like Codex and Goose in a unified environment for orchestration and collaboration.

GitHub - DiegoZaluski/co-sh: Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use. · GitHub
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
DiegoZaluski
co-sh Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use.
Readme Apache-2.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
1,273 Commits 1,273 Commits Folders and files
.github/ workflows .github/ workflows assets assets crates crates docs docs src src .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE NOTICE NOTICE README.md README.md download.ps1 download.ps1 download.sh download.sh View all files Repository files navigation
English | Português | Español | Русский | 中文 | 日本語
Cosh is a general-purpose coding agent for the terminal, with software engineering tools and computer tools that let it interact with graphical applications.
Multi-agent mode ( /agents ) — run external agents (Gemini, Claude, Codex, Cline, …) alongside the main agent in split panes, all reachable from the same prompt. Each agent keeps its own session, model and optional git worktree, and the main agent coordinates them — dispatching tasks and collecting reports as one team.
Computer tools via accessibility tree — the agent reads and operates graphical interfaces through a structured, compact representation, falling back to screenshots only when needed. Fewer images processed, fewer tokens consumed.
TUI built from scratch in Rust (Ratatui), inspired by OpenCode's UX. Providers, models and operating modes are configured right in the interface — almost zero manual setup.
Independent crates — the development tools live in crates decoupled from the application, reusable as libraries outside Cosh.
Tested on Linux and Windows. macOS hasn't been tested yet — if anything fails, open an issue with as much detail as possible.
Linux & macOS (also works on Windows via Git Bash / WSL):
curl -fsSL https://raw.githubusercontent.com/DiegoZaluski/co-sh/main/download.sh | bash
Windows (PowerShell):
irm https: // raw.githubusercontent.com / DiegoZaluski / co - sh / main / download.ps1 | iex
Note: the curl | bash command does not work in PowerShell or CMD — Windows has no native bash . Use the PowerShell script above instead (Git Bash and WSL also work).
Optional variables: COSH_BIN_DIR (install directory), COSH_VERSION (specific version) and COSH_VARIANT ( slim , the default, or slim-embed , with local RAG via fastembed). On PowerShell, set them as $env:COSH_VERSION = "v0.1.0" before running the command.
After installing, run cosh in your terminal. If the install directory isn't on your PATH , the script prints the command to add it. If anything goes wrong during installation, open an issue — I'll be glad to help!
Crate
Description
cosh-tools
Development tools and computer tools
cosh-tui
TUI widget library (markdown/diff rendering, layout) — the components the app's terminal interface is built on
cosh-sdk
SDK for building agents on top of cosh
cosh-recall
Memory and context recall
cosh-onnx
ONNX decision engine (behind the onnx feature)
Roadmap
Laya/ONNX decision engine — optional onnx feature backed by the cosh-onnx crate; see the crate guide .
Graph-based harness orchestration — a new session mode for coordinating multiple harnesses.
Open-source coding agent for the terminal, written in Rust. Fast TUI, LSP integration, and computer use.
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
