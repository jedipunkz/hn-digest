---
source: "https://github.com/FrigadeHQ/hob"
hn_url: "https://news.ycombinator.com/item?id=50246238"
title: "Show HN: Hob - Open-source, on device alternative to Claude Cowork and ChatGPT"
article_title: "GitHub - FrigadeHQ/hob: An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline. · GitHub"
image: "https://repository-images.githubusercontent.com/1419901875/bef3a5ac-8f6b-4f29-8f2a-30f31c1771d7"
author: "pancomplex"
captured_at: "2026-10-26T16:33:13Z"
capture_tool: "hn-digest"
hn_id: 50246238
score: 3
comments: 0
posted_at: "2026-10-26T16:12:29Z"
tags:
  - hacker-news
---

# Show HN: Hob - Open-source, on device alternative to Claude Cowork and ChatGPT

- HN: [50246238](https://news.ycombinator.com/item?id=50246238)
- Source: [github.com](https://github.com/FrigadeHQ/hob)
- Score: 3
- Comments: 0
- Posted: 2026-10-26T16:12:29Z

## Translation

Title: Show HN: Hob - Open-source, on device alternative to Claude Cowork and ChatGPT
Article title: GitHub - FrigadeHQ/hob: An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline. · GitHub
Description: An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline. - FrigadeHQ/hob
HN text: Hey HN, I'm Christian and this is a new OSS project I've been noodling on for the last couple of weeks. Hob is a free, MIT-licensed desktop app for Mac, Windows and Linux that works on your files the way the Claude and ChatGPT desktop apps do, except the model runs on your own computer. It uses open-weight models (currently Gemma 4 and gpt-oss). All your chats and files never leave your system. The only things that go online are web search (can be turned off in settings) and any MCP connectors you add yourself. Main features: - Automatically finds the best open-weight model for your hardware on first install (Mac/Linux/Windows) - Writes sandboxed Python code for ad hoc tools (data processing, document conversion, graphs etc.) - Built-in web search (powered by SearXNG) - MCP support - Image recognition - CRUD to your local filesystem - Works fully offline with no internet Before building this I tried every local AI app I could find, but it seemed each one had different problems. The main thing I noticed is that none of them could do what the ChatGPT and Claude apps do with your files, e.g. read and change files in a folder, look at images, run code in a sandbox, draw a chart. The most advanced thing any of them did was paste a plain text file into a small context window, which basically failed as soon as you had a file larger than a few KB. I couldn't find anything that could easily parse images and bigger PDFs. The OSS solutions that were more advanced all seemed to default to hosted models (either OSS or closed), which defeats the entire point of the privacy that on-device OSS offers. They also seemed to push towards an account/open-core style business. I build it using Electron bridged with llama.cpp in its own process. On initial load, the system will try and figure out if you have a compatible graphics card and if not default to CPU-only (I ran into a lot of issues of graphics card beeing seemingly compatible but then crashing on first query, there's a whole saga to this lol). Code execution happens in a sandbox that can only write to the chat's folder and has no network unless asked (sandbox-exec on macOS, Landlock on Linux and a low-integrity process on Windows). Images go to a small Qwen3-VL model. Dictation is a local Whisper model. The hardest part was definitely to get the smaller models to use tools reliably. Ran into a ton of issues of 4B models happily saying "here's your chart" without ever calling the chart tool or get stuck in an infinite call loop. I spent a fair amount of the work building guards vs tweaking prompts (that didn't really work for small models). I understand the pushback that local models are weaker than frontier cloud models. This is true, but for most tasks for most normal people, I'd argue you rarely need the newest frontier model. From my own evals, systems with 16+ GB of RAM do really well, and if you have a GPU or more RAM, you get very close to what you'll get out of the box with a frontier hosted model. Excited to hear what you think and see how it runs on other systems!

Article text:
GitHub - FrigadeHQ/hob: An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline. · GitHub
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
FrigadeHQ
hob Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
189 Commits 189 Commits Folders and files
.github/ workflows .github/ workflows app app catalog catalog design design docs docs packages/ picker packages/ picker site site .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
An easier way to use AI on your own computer.
A free, open-source alternative to Claude and ChatGPT.
Visit hob.chat → ·
Download for Mac · Windows · Linux
Apple silicon Macs, Windows and Linux, with 8 GB of memory or more.
Fill out PDFs, work through documents, and much more. Free, private, and works offline.
Hob picks the best frontier open-weight model that fits your computer and sets everything up for you.
Help · Contributing · MIT license
Made by Frigade .
An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline. - FrigadeHQ/hob

Hey HN, I'm Christian and this is a new OSS project I've been noodling on for the last couple of weeks. Hob is a free, MIT-licensed desktop app for Mac, Windows and Linux that works on your files the way the Claude and ChatGPT desktop apps do, except the model runs on your own computer. It uses open-weight models (currently Gemma 4 and gpt-oss). All your chats and files never leave your system. The only things that go online are web search (can be turned off in settings) and any MCP connectors you add yourself. Main features: - Automatically finds the best open-weight model for your hardware on first install (Mac/Linux/Windows) - Writes sandboxed Python code for ad hoc tools (data processing, document conversion, graphs etc.) - Built-in web search (powered by SearXNG) - MCP support - Image recognition - CRUD to your local filesystem - Works fully offline with no internet Before building this I tried every local AI app I could find, but it seemed each one had different problems. The main thing I noticed is that none of them could do what the ChatGPT and Claude apps do with your files, e.g. read and change files in a folder, look at images, run code in a sandbox, draw a chart. The most advanced thing any of them did was paste a plain text file into a small context window, which basically failed as soon as you had a file larger than a few KB. I couldn't find anything that could easily parse images and bigger PDFs. The OSS solutions that were more advanced all seemed to default to hosted models (either OSS or closed), which defeats the entire point of the privacy that on-device OSS offers. They also seemed to push towards an account/open-core style business. I build it using Electron bridged with llama.cpp in its own process. On initial load, the system will try and figure out if you have a compatible graphics card and if not default to CPU-only (I ran into a lot of issues of graphics card beeing seemingly compatible but then crashing on first query, there's a whole saga to this lol). Code execution happens in a sandbox that can only write to the chat's folder and has no network unless asked (sandbox-exec on macOS, Landlock on Linux and a low-integrity process on Windows). Images go to a small Qwen3-VL model. Dictation is a local Whisper model. The hardest part was definitely to get the smaller models to use tools reliably. Ran into a ton of issues of 4B models happily saying "here's your chart" without ever calling the chart tool or get stuck in an infinite call loop. I spent a fair amount of the work building guards vs tweaking prompts (that didn't really work for small models). I understand the pushback that local models are weaker than frontier cloud models. This is true, but for most tasks for most normal people, I'd argue you rarely need the newest frontier model. From my own evals, systems with 16+ GB of RAM do really well, and if you have a GPU or more RAM, you get very close to what you'll get out of the box with a frontier hosted model. Excited to hear what you think and see how it runs on other systems!

GitHub - FrigadeHQ/hob: An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline. · GitHub
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
FrigadeHQ
hob Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
189 Commits 189 Commits Folders and files
.github/ workflows .github/ workflows app app catalog catalog design design docs docs packages/ picker packages/ picker site site .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
An easier way to use AI on your own computer.
A free, open-source alternative to Claude and ChatGPT.
Visit hob.chat → ·
Download for Mac · Windows · Linux
Apple silicon Macs, Windows and Linux, with 8 GB of memory or more.
Fill out PDFs, work through documents, and much more. Free, private, and works offline.
Hob picks the best frontier open-weight model that fits your computer and sets everything up for you.
Help · Contributing · MIT license
Made by Frigade .
An easier way to use AI on your own computer. A free, open-source alternative to Claude and ChatGPT that runs open-weight models like Gemma and gpt-oss locally on Mac, Windows and Linux. Private, and works offline.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
