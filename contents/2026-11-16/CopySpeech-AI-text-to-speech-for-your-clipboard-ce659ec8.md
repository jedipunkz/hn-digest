---
source: "https://github.com/ilyaizen/CopySpeech"
hn_url: "https://news.ycombinator.com/item?id=50524875"
title: "CopySpeech – AI text-to-speech for your clipboard"
article_title: "GitHub - ilyaizen/CopySpeech: Clipboard AI text-to-speech 🔊 · GitHub"
image: "https://opengraph.githubassets.com/432024f261a7af72fdb354a548bb698cc56d090a21559fc43c5d73043f335033/ilyaizen/CopySpeech"
author: "ilyaizen"
captured_at: "2026-11-16T23:41:52Z"
capture_tool: "hn-digest"
hn_id: 50524875
score: 4
comments: 0
posted_at: "2026-11-16T23:08:26Z"
tags:
  - hacker-news
---

# CopySpeech – AI text-to-speech for your clipboard

- HN: [50524875](https://news.ycombinator.com/item?id=50524875)
- Source: [github.com](https://github.com/ilyaizen/CopySpeech)
- Score: 4
- Comments: 0
- Posted: 2026-11-16T23:08:26Z

## Translation

Title: CopySpeech – AI text-to-speech for your clipboard
Article title: GitHub - ilyaizen/CopySpeech: Clipboard AI text-to-speech 🔊 · GitHub
Description: Clipboard AI text-to-speech 🔊. Contribute to ilyaizen/CopySpeech development by creating an account on GitHub.

Article text:
GitHub - ilyaizen/CopySpeech: Clipboard AI text-to-speech 🔊 · GitHub
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
ilyaizen
CopySpeech Public Notifications You must be signed in to change notification settings
Star 14 ( 14 ) You must be signed in to star a repository
copyspeech.ilyaizen.link Topics
Readme MIT license Contributing
2 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
431 Commits 431 Commits Folders and files
.github/ workflows .github/ workflows .pi/ extensions/ copyspeech .pi/ extensions/ copyspeech .vscode .vscode browser-extension browser-extension browser-native-host browser-native-host copyspeech-dsh copyspeech-dsh copyspeech-hermes copyspeech-hermes docs docs scripts scripts src-tauri src-tauri src src static static tools/ oxlint/ anti-slop tools/ oxlint/ anti-slop .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md BROWSER_EXTENSION.md BROWSER_EXTENSION.md CHANGELOG.md CHANGELOG.md DESIGN.md DESIGN.md Dockerfile Dockerfile GLOSSARY.md GLOSSARY.md IDEA.md IDEA.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md bun.lock bun.lock components.json components.json crowdin.yml crowdin.yml file-structure.txt file-structure.txt nginx.conf nginx.conf package.json package.json tsconfig.json tsconfig.json vite.config.js vite.config.js vitest-plugin-mock-state.ts vitest-plugin-mock-state.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
AI text-to-speech for your clipboard. Copy twice. Hear it.
CopySpeech reads your clipboard aloud. Copy the same text twice within 1.5 s, and an AI voice speaks it, with a HUD and word-level captions.
Download · Changelog · Contributing · Issues
Double-copy trigger. No hotkey needed. The tray and the Play page work too.
16 engine backends.
Local: Kokoro (default), Piper, Kitten, Qwen3-TTS, Pocket, and any CLI TTS tool.
Cloud: OpenAI, ElevenLabs, Cartesia, Fish Audio, xAI, Mistral, Google, Microsoft Azure, Edge (keyless), and any HTTP endpoint.
One-click local installs. Kokoro installs from onboarding. Installs stream progress and support optional CUDA.
Live captions. Words highlight in sync with the audio on the Play page, the HUD, and the browser companion. Captions need an engine that reports timings.
Streaming playback. Local engines start speaking after the first audio window.
Speed 0.25×–4×, pitch, and volume. Voice profiles store each combination per engine.
Duck other apps (Windows). Other apps get quieter while CopySpeech speaks.
Text cleanup. Strips markdown, expands URLs, and normalizes Unicode. It handles non-Latin scripts without damage.
Browser companion. Reads your selection through the app. Chrome, Edge, and Firefox are supported. The page highlights the passage as it plays.
Control server. A local HTTP API for scripts and tools. See Control server .
Interface. 9 languages, right-to-left layouts, tray, dark and light themes, and an auto-updater.
Experimental features are alpha. They can change or break between releases.
The Experimental Features (Alpha) switch lives in Settings → Advanced . It gates these features:
When the switch is off, these features pause. Their settings and saved readings stay on disk.
DSH partial auto-read (0.1.1). Reads only the first or last paragraphs of each automatic reply.
Not yet verified live, per the plugin notes:
Hermes follow-scroll playback.
The explicit-Save settings page in the Hermes plugin.
Get a build from Releases . Replace <version> with the release number. The latest is 0.3.0.
Onboarding asks you to pick an engine and a voice on first launch. Kokoro is the default.
Linux note. Wayland blocks global hotkeys. Use double-copy, the tray, or a compositor keybind that calls the control server .
Open chrome://extensions , turn on Developer mode , and choose Load unpacked on the extracted folder.
Grant per-site access in the extension settings.
On Windows, installed builds register the Chrome host at launch. Edge, Firefox, Linux, and macOS need manual setup.
Hermes Desktop plugin ( copyspeech-hermes )
Reads each new completed Hermes reply aloud through the running app. It can highlight spoken words and follow playback.
Create a junction so Hermes loads the repo file in place:
New-Item - ItemType Junction - Path " $ env: LOCALAPPDATA \hermes\desktop-plugins\copyspeech-hermes " - Target " <repo>\copyspeech-hermes "
Open Settings → Plugins → CopySpeech and turn on Read new replies aloud .
Turn off Hermes's built-in Read replies aloud . If both run, each reply plays twice.
Changes apply when you press Save . Edits to plugin.js reload in the app.
DSH is the DeepSeek Harness. This plugin adds a sidebar panel, a speaker action on each reply, a selection pill, and a playback bar with pause, ±5 s seek, and stop.
Ask the DSH plugin manager to run install_bundle with the absolute path to copyspeech-dsh/ .
Open the CopySpeech panel. The status line reads CopySpeech is running.
Send a message and click the speaker in the reply's action row.
The plugin needs the control server on 127.0.0.1:43117 . Full details are in copyspeech-dsh/README.md .
The Pi extension and the Claude Code hook speak final assistant replies. See docs/agent-voice.md .
A local HTTP API on 127.0.0.1:43117 :
Details are in docs/profile-engine-settings.md .
bun install
bun run tauri dev
Command
Purpose
bun run dev
Frontend only, on port 5173
bun run check
Svelte and TypeScript checks, then lint
bun run test
Frontend tests (Vitest)
bun run format
Format the repo
bun run tauri build
Production build
cd src-tauri && cargo test
Rust tests
bun run bump
Patch version bump ( bump:minor and bump:major also exist)
Layer
Stack
Backend
Rust, Tauri 2
Frontend
Svelte 5, SvelteKit, TypeScript, Tailwind CSS 4
Audio
rodio, Web Audio ( OfflineAudioContext )
Tooling
Bun, Vite+ ( vp ), shadcn-svelte
i18n
svelte-i18n, Crowdin
Read Understanding the Codebase before you change code. docs/engines.md explains how local engines connect.
The next release targets low time-to-first-audio and a small VRAM and RAM footprint. The goal is local engines that run well on low-VRAM GPUs and CPU-only machines.
Read the Contributing Guide for setup, style, and pull request steps.
copyspeech.ilyaizen.link Topics
Readme MIT license Contributing
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Clipboard AI text-to-speech 🔊. Contribute to ilyaizen/CopySpeech development by creating an account on GitHub.

GitHub - ilyaizen/CopySpeech: Clipboard AI text-to-speech 🔊 · GitHub
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
ilyaizen
CopySpeech Public Notifications You must be signed in to change notification settings
Star 14 ( 14 ) You must be signed in to star a repository
copyspeech.ilyaizen.link Topics
Readme MIT license Contributing
2 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
431 Commits 431 Commits Folders and files
.github/ workflows .github/ workflows .pi/ extensions/ copyspeech .pi/ extensions/ copyspeech .vscode .vscode browser-extension browser-extension browser-native-host browser-native-host copyspeech-dsh copyspeech-dsh copyspeech-hermes copyspeech-hermes docs docs scripts scripts src-tauri src-tauri src src static static tools/ oxlint/ anti-slop tools/ oxlint/ anti-slop .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md BROWSER_EXTENSION.md BROWSER_EXTENSION.md CHANGELOG.md CHANGELOG.md DESIGN.md DESIGN.md Dockerfile Dockerfile GLOSSARY.md GLOSSARY.md IDEA.md IDEA.md LICENSE LICENSE PRODUCT.md PRODUCT.md README.md README.md bun.lock bun.lock components.json components.json crowdin.yml crowdin.yml file-structure.txt file-structure.txt nginx.conf nginx.conf package.json package.json tsconfig.json tsconfig.json vite.config.js vite.config.js vitest-plugin-mock-state.ts vitest-plugin-mock-state.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
AI text-to-speech for your clipboard. Copy twice. Hear it.
CopySpeech reads your clipboard aloud. Copy the same text twice within 1.5 s, and an AI voice speaks it, with a HUD and word-level captions.
Download · Changelog · Contributing · Issues
Double-copy trigger. No hotkey needed. The tray and the Play page work too.
16 engine backends.
Local: Kokoro (default), Piper, Kitten, Qwen3-TTS, Pocket, and any CLI TTS tool.
Cloud: OpenAI, ElevenLabs, Cartesia, Fish Audio, xAI, Mistral, Google, Microsoft Azure, Edge (keyless), and any HTTP endpoint.
One-click local installs. Kokoro installs from onboarding. Installs stream progress and support optional CUDA.
Live captions. Words highlight in sync with the audio on the Play page, the HUD, and the browser companion. Captions need an engine that reports timings.
Streaming playback. Local engines start speaking after the first audio window.
Speed 0.25×–4×, pitch, and volume. Voice profiles store each combination per engine.
Duck other apps (Windows). Other apps get quieter while CopySpeech speaks.
Text cleanup. Strips markdown, expands URLs, and normalizes Unicode. It handles non-Latin scripts without damage.
Browser companion. Reads your selection through the app. Chrome, Edge, and Firefox are supported. The page highlights the passage as it plays.
Control server. A local HTTP API for scripts and tools. See Control server .
Interface. 9 languages, right-to-left layouts, tray, dark and light themes, and an auto-updater.
Experimental features are alpha. They can change or break between releases.
The Experimental Features (Alpha) switch lives in Settings → Advanced . It gates these features:
When the switch is off, these features pause. Their settings and saved readings stay on disk.
DSH partial auto-read (0.1.1). Reads only the first or last paragraphs of each automatic reply.
Not yet verified live, per the plugin notes:
Hermes follow-scroll playback.
The explicit-Save settings page in the Hermes plugin.
Get a build from Releases . Replace <version> with the release number. The latest is 0.3.0.
Onboarding asks you to pick an engine and a voice on first launch. Kokoro is the default.
Linux note. Wayland blocks global hotkeys. Use double-copy, the tray, or a compositor keybind that calls the control server .
Open chrome://extensions , turn on Developer mode , and choose Load unpacked on the extracted folder.
Grant per-site access in the extension settings.
On Windows, installed builds register the Chrome host at launch. Edge, Firefox, Linux, and macOS need manual setup.
Hermes Desktop plugin ( copyspeech-hermes )
Reads each new completed Hermes reply aloud through the running app. It can highlight spoken words and follow playback.
Create a junction so Hermes loads the repo file in place:
New-Item - ItemType Junction - Path " $ env: LOCALAPPDATA \hermes\desktop-plugins\copyspeech-hermes " - Target " <repo>\copyspeech-hermes "
Open Settings → Plugins → CopySpeech and turn on Read new replies aloud .
Turn off Hermes's built-in Read replies aloud . If both run, each reply plays twice.
Changes apply when you press Save . Edits to plugin.js reload in the app.
DSH is the DeepSeek Harness. This plugin adds a sidebar panel, a speaker action on each reply, a selection pill, and a playback bar with pause, ±5 s seek, and stop.
Ask the DSH plugin manager to run install_bundle with the absolute path to copyspeech-dsh/ .
Open the CopySpeech panel. The status line reads CopySpeech is running.
Send a message and click the speaker in the reply's action row.
The plugin needs the control server on 127.0.0.1:43117 . Full details are in copyspeech-dsh/README.md .
The Pi extension and the Claude Code hook speak final assistant replies. See docs/agent-voice.md .
A local HTTP API on 127.0.0.1:43117 :
Details are in docs/profile-engine-settings.md .
bun install
bun run tauri dev
Command
Purpose
bun run dev
Frontend only, on port 5173
bun run check
Svelte and TypeScript checks, then lint
bun run test
Frontend tests (Vitest)
bun run format
Format the repo
bun run tauri build
Production build
cd src-tauri && cargo test
Rust tests
bun run bump
Patch version bump ( bump:minor and bump:major also exist)
Layer
Stack
Backend
Rust, Tauri 2
Frontend
Svelte 5, SvelteKit, TypeScript, Tailwind CSS 4
Audio
rodio, Web Audio ( OfflineAudioContext )
Tooling
Bun, Vite+ ( vp ), shadcn-svelte
i18n
svelte-i18n, Crowdin
Read Understanding the Codebase before you change code. docs/engines.md explains how local engines connect.
The next release targets low time-to-first-audio and a small VRAM and RAM footprint. The goal is local engines that run well on low-VRAM GPUs and CPU-only machines.
Read the Contributing Guide for setup, style, and pull request steps.
copyspeech.ilyaizen.link Topics
Readme MIT license Contributing
2 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
