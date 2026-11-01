---
source: "https://castov.com"
hn_url: "https://news.ycombinator.com/item?id=50329149"
title: "Show HN: Castov – A local-first desktop automation engine using local AI"
article_title: "Castov | Developer Tools & MCP Server (100% Client-Side, Offline & CLI)"
image: "https://castov.com/opengraph_image.jpg"
author: "castov"
captured_at: "2026-11-01T20:15:33Z"
capture_tool: "hn-digest"
hn_id: 50329149
score: 2
comments: 0
posted_at: "2026-11-01T19:59:55Z"
tags:
  - hacker-news
---

# Show HN: Castov – A local-first desktop automation engine using local AI

- HN: [50329149](https://news.ycombinator.com/item?id=50329149)
- Source: [castov.com](https://castov.com)
- Score: 2
- Comments: 0
- Posted: 2026-11-01T19:59:55Z

## Translation

Title: Show HN: Castov – A local-first desktop automation engine using local AI
Article title: Castov | Developer Tools & MCP Server (100% Client-Side, Offline & CLI)
Description: Fast, 100% client-side developer platform: Web + CLI + MCP Server for Cursor and Claude. Zero-data logging, JSON to TypeScript, Cron multi-timezone, and cryptographic pipelines.

Article text:
Castov | Developer Tools & MCP Server (100% Client-Side, Offline & CLI) Skip to main content Search tools & conversions... Ctrl+K Tools →
Castov DEV Tools MCP Server AI Converters Guides Radar & Hubs Search... ⌘ K 100% Client-Side • Zero-Data Logging • Works 100% Offline The Complete Developer Toolkit
Web + CLI + AI (MCP)
All your mission-critical utilities executed locally in your browser, terminal, and AI agents. Zero server tracking, sub-millisecond transforms, and seamless multi-stage pipelines.
Smart Omnibox: Ctrl + K Terminal CLI: npx castov Cursor & Claude: castov mcp Explore 61 Tools MCP Server (13 Tools) CLI Guide Instant Timestamp & Date Inspector castov-playground ~ zsh npm i -g castov Try: Secret 32B Security Cron Parser DevOps CIDR Calculator Network Fix Port 3000 Instant Fix Crypto UUID Utility Castov DevShell v0.2.0 (Zero-Latency Browser Sandbox) Simulates the full castov-cli suite locally via WebCrypto and instant transforms.
Click any quick-run chip above or type help to see commands.
Built for Speed, Privacy & AI Why Developers Choose Castov
Five architectural decisions that make Castov the preferred toolkit for modern software teams.
100% of cryptography, JWT inspection, JSON formatting, and timestamp parsing executes in local browser memory. Not a single keystroke is sent to external servers.
Pass outputs seamlessly across utilities in one click: JSON directly to TypeScript interfaces, Cron schedules into world timezones, and diffs to schema generators without clipboard juggling.
The exact same 13 utilities work on the Web, in your local terminal via castov , and inside Cursor or Claude via our Grade-A certified Model Context Protocol server.
Share inputs, diffs, and settings with teammates instantly. Payloads are compressed into the URL hash fragment ( # ) so servers never see what you shared.
Zero-friction leverage unlocked on registration
Connect your GitHub account to sync pinned utilities, presets, and dev preferences across all your machines. Collaborate with teammates using RSA-OAEP encrypted Team Vaults to share .env files without cleartext exposure in chat channels.
GLOBAL DISTRIBUTION & MCP ECOSYSTEM Deploy Castov Everywhere You Build
Whether inside your local terminal, Docker containers, AI editors, or automated CI/CD pipelines — Castov provides the instant execution layer for developers and MCP AI agents.
Docker Hub Homebrew Tap Open VSX & VS Code GitHub Action Official Container Verified Distribution Docker Hub
Run Castov CLI and MCP runtime in an isolated, non-root Alpine container with zero host dependencies.
terminal execution Copy docker run --rm -it castovdev/castov:latest --help Lightweight Alpine runtime (~50MB footprint) Strict non-root user execution (USER node) Direct MCP transport streaming via stdin/stdout Advanced Capabilities
We didn't just build formatters. We built the architecture required by modern engineering teams.
A visual node-editor for building data pipelines. Connects peers in real-time using WebRTC and CRDTs (Yjs) with direct peer-to-peer data channels coordinated through public signaling servers (which may receive connection metadata). Workflow graphs, node configurations, and payload data are never stored on Castov servers.
Steganography Environment Vaults
End-to-end encrypted environment variables. We hide your secure AES keys inside standard image pixels (Steganography) for invisible sharing, entirely in your browser.
API Mocker (Service Worker) Intercepting GET /api/v1/users {
"status": 200,
"data": [
{ "id": "usr_91x", "role": "admin" }
]
} Test your frontend without a backend. Our Local API Mocker runs entirely in your browser's Service Worker to simulate payloads globally.
From time manipulation to cryptography, everything is beautifully structured and instantly accessible.
Convert, validate, compare and work with Unix timestamps.
Calculate differences, durations, and add or subtract time from a date.
Convert, compare and browse time across the world's time zones.
ISO 8601, RFC 3339, cron and other formats developers work with daily.
Unix Timestamp Batch Converter
Docker Architecture Visualizer
App & Screenshot Mockup Studio
Color Code Converter & Palette Studio
Git Disaster Recovery & Time-Traveler
SVG Keyframe Animator & Studio
Experimental tools: WebRTC P2P mesh, steganography, AST parsing, and in-browser WebGL computing.
Generate local fetch interceptor scripts to mock API endpoints.
P2P Ghost Tunnel Lab Share localhost via direct P2P WebRTC connection.
Steganography Env Vault Lab Hide encrypted .env files inside PNG image pixels.
Holographic JSON Explorer Lab Explore massive JSON architectures in a 3D WebGL universe.
Genetic Regex Auto-Healer Lab Evolve perfect Regex using local genetic algorithms.
Chrono-Debug API Replay Lab Replay past API states using a visual timeline scrubber.
AST Conflict Telepathy Lab Semantically merge Git conflicts using AST logic parsing.
NexusGrid Lab Decentralized WebGPU Supercomputer running local AI workloads.
Sonic-Gap Bridge Lab Air-gapped ultrasonic acoustic data transfer protocol.
Quantum Simulator Lab Browser-native quantum circuit execution via matrix tensors.
Everything you need to know about Castov.
Yes. All tools are free, with no account and no usage limits. The site may show ads in the future to cover running costs. Is my data sent to a server? Time, date and converter tools do not upload your input to Castov servers. Some preferences, such as theme, custom time zones and selected models, are stored locally on your device. Some features use the network by design: Env Vault uploads an encrypted copy, P2P Ghost Tunnel uses public STUN discovery with manual token exchange, and Developer Canvas connects to public WebRTC signaling servers (which may receive connection metadata; workflow graphs and payloads are never stored on Castov servers). What is a Unix timestamp? A Unix timestamp is the number of seconds since 1 January 1970 at 00:00:00 UTC. It identifies one moment in time regardless of time zone, which is why programs and APIs use it. Should I use seconds or milliseconds? Use whichever your system produces. A 10-digit value is in seconds and a 13-digit value is in milliseconds. The converters can detect the unit for you. Which time zones are supported? All IANA time zones supported by your browser, such as America/New_York, Europe/Paris, Africa/Algiers and Asia/Tokyo, with daylight saving time rules applied for the date you choose. Ready to supercharge your workflow?
Free browser-based developer tools. Some collaboration, cloud, and external API features may require an account or additional configuration.
Systems Operational 00:00:00 UTC • EPOCH -------- Client-First • No Payload Logging $ npm i -g castov Star on GitHub DEVELOPER UPDATES & RELEASES Precision utilities designed for daily engineering workflows
Never miss new developer tools, CLI feature updates, or cheat sheet releases. No spam, no behavioral tracking.
Subscribe Castov The high-velocity developer toolbox. Client-side cryptography, zero payload logging, and instant execution without network latency. Standard infrastructure logs and aggregate analytics may apply.
Node.js 18+ No Payload Logging Core Utilities
VS Code Extensions Suite 13 Extensions
Interactive Cheat Sheets Popular

## Original Extract

Fast, 100% client-side developer platform: Web + CLI + MCP Server for Cursor and Claude. Zero-data logging, JSON to TypeScript, Cron multi-timezone, and cryptographic pipelines.

Castov | Developer Tools & MCP Server (100% Client-Side, Offline & CLI) Skip to main content Search tools & conversions... Ctrl+K Tools →
Castov DEV Tools MCP Server AI Converters Guides Radar & Hubs Search... ⌘ K 100% Client-Side • Zero-Data Logging • Works 100% Offline The Complete Developer Toolkit
Web + CLI + AI (MCP)
All your mission-critical utilities executed locally in your browser, terminal, and AI agents. Zero server tracking, sub-millisecond transforms, and seamless multi-stage pipelines.
Smart Omnibox: Ctrl + K Terminal CLI: npx castov Cursor & Claude: castov mcp Explore 61 Tools MCP Server (13 Tools) CLI Guide Instant Timestamp & Date Inspector castov-playground ~ zsh npm i -g castov Try: Secret 32B Security Cron Parser DevOps CIDR Calculator Network Fix Port 3000 Instant Fix Crypto UUID Utility Castov DevShell v0.2.0 (Zero-Latency Browser Sandbox) Simulates the full castov-cli suite locally via WebCrypto and instant transforms.
Click any quick-run chip above or type help to see commands.
Built for Speed, Privacy & AI Why Developers Choose Castov
Five architectural decisions that make Castov the preferred toolkit for modern software teams.
100% of cryptography, JWT inspection, JSON formatting, and timestamp parsing executes in local browser memory. Not a single keystroke is sent to external servers.
Pass outputs seamlessly across utilities in one click: JSON directly to TypeScript interfaces, Cron schedules into world timezones, and diffs to schema generators without clipboard juggling.
The exact same 13 utilities work on the Web, in your local terminal via castov , and inside Cursor or Claude via our Grade-A certified Model Context Protocol server.
Share inputs, diffs, and settings with teammates instantly. Payloads are compressed into the URL hash fragment ( # ) so servers never see what you shared.
Zero-friction leverage unlocked on registration
Connect your GitHub account to sync pinned utilities, presets, and dev preferences across all your machines. Collaborate with teammates using RSA-OAEP encrypted Team Vaults to share .env files without cleartext exposure in chat channels.
GLOBAL DISTRIBUTION & MCP ECOSYSTEM Deploy Castov Everywhere You Build
Whether inside your local terminal, Docker containers, AI editors, or automated CI/CD pipelines — Castov provides the instant execution layer for developers and MCP AI agents.
Docker Hub Homebrew Tap Open VSX & VS Code GitHub Action Official Container Verified Distribution Docker Hub
Run Castov CLI and MCP runtime in an isolated, non-root Alpine container with zero host dependencies.
terminal execution Copy docker run --rm -it castovdev/castov:latest --help Lightweight Alpine runtime (~50MB footprint) Strict non-root user execution (USER node) Direct MCP transport streaming via stdin/stdout Advanced Capabilities
We didn't just build formatters. We built the architecture required by modern engineering teams.
A visual node-editor for building data pipelines. Connects peers in real-time using WebRTC and CRDTs (Yjs) with direct peer-to-peer data channels coordinated through public signaling servers (which may receive connection metadata). Workflow graphs, node configurations, and payload data are never stored on Castov servers.
Steganography Environment Vaults
End-to-end encrypted environment variables. We hide your secure AES keys inside standard image pixels (Steganography) for invisible sharing, entirely in your browser.
API Mocker (Service Worker) Intercepting GET /api/v1/users {
"status": 200,
"data": [
{ "id": "usr_91x", "role": "admin" }
]
} Test your frontend without a backend. Our Local API Mocker runs entirely in your browser's Service Worker to simulate payloads globally.
From time manipulation to cryptography, everything is beautifully structured and instantly accessible.
Convert, validate, compare and work with Unix timestamps.
Calculate differences, durations, and add or subtract time from a date.
Convert, compare and browse time across the world's time zones.
ISO 8601, RFC 3339, cron and other formats developers work with daily.
Unix Timestamp Batch Converter
Docker Architecture Visualizer
App & Screenshot Mockup Studio
Color Code Converter & Palette Studio
Git Disaster Recovery & Time-Traveler
SVG Keyframe Animator & Studio
Experimental tools: WebRTC P2P mesh, steganography, AST parsing, and in-browser WebGL computing.
Generate local fetch interceptor scripts to mock API endpoints.
P2P Ghost Tunnel Lab Share localhost via direct P2P WebRTC connection.
Steganography Env Vault Lab Hide encrypted .env files inside PNG image pixels.
Holographic JSON Explorer Lab Explore massive JSON architectures in a 3D WebGL universe.
Genetic Regex Auto-Healer Lab Evolve perfect Regex using local genetic algorithms.
Chrono-Debug API Replay Lab Replay past API states using a visual timeline scrubber.
AST Conflict Telepathy Lab Semantically merge Git conflicts using AST logic parsing.
NexusGrid Lab Decentralized WebGPU Supercomputer running local AI workloads.
Sonic-Gap Bridge Lab Air-gapped ultrasonic acoustic data transfer protocol.
Quantum Simulator Lab Browser-native quantum circuit execution via matrix tensors.
Everything you need to know about Castov.
Yes. All tools are free, with no account and no usage limits. The site may show ads in the future to cover running costs. Is my data sent to a server? Time, date and converter tools do not upload your input to Castov servers. Some preferences, such as theme, custom time zones and selected models, are stored locally on your device. Some features use the network by design: Env Vault uploads an encrypted copy, P2P Ghost Tunnel uses public STUN discovery with manual token exchange, and Developer Canvas connects to public WebRTC signaling servers (which may receive connection metadata; workflow graphs and payloads are never stored on Castov servers). What is a Unix timestamp? A Unix timestamp is the number of seconds since 1 January 1970 at 00:00:00 UTC. It identifies one moment in time regardless of time zone, which is why programs and APIs use it. Should I use seconds or milliseconds? Use whichever your system produces. A 10-digit value is in seconds and a 13-digit value is in milliseconds. The converters can detect the unit for you. Which time zones are supported? All IANA time zones supported by your browser, such as America/New_York, Europe/Paris, Africa/Algiers and Asia/Tokyo, with daylight saving time rules applied for the date you choose. Ready to supercharge your workflow?
Free browser-based developer tools. Some collaboration, cloud, and external API features may require an account or additional configuration.
Systems Operational 00:00:00 UTC • EPOCH -------- Client-First • No Payload Logging $ npm i -g castov Star on GitHub DEVELOPER UPDATES & RELEASES Precision utilities designed for daily engineering workflows
Never miss new developer tools, CLI feature updates, or cheat sheet releases. No spam, no behavioral tracking.
Subscribe Castov The high-velocity developer toolbox. Client-side cryptography, zero payload logging, and instant execution without network latency. Standard infrastructure logs and aggregate analytics may apply.
Node.js 18+ No Payload Logging Core Utilities
VS Code Extensions Suite 13 Extensions
Interactive Cheat Sheets Popular
