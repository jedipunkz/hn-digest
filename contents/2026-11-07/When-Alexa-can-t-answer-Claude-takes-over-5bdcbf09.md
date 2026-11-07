---
source: "https://github.com/rg321/hey-claude"
hn_url: "https://news.ycombinator.com/item?id=50400738"
title: "When Alexa can't answer, Claude takes over"
article_title: "GitHub - rg321/hey-claude: When Alexa can't answer, Claude takes over. AI home assistant powered by Claude. · GitHub"
image: "https://opengraph.githubassets.com/6cccc4bd17e75d75d0ae26c8671da59b8fe83b82667617c05367125c06b36f03/rg321/hey-claude"
author: "mparnisari"
captured_at: "2026-11-07T01:39:42Z"
capture_tool: "hn-digest"
hn_id: 50400738
score: 2
comments: 0
posted_at: "2026-11-07T01:29:23Z"
tags:
  - hacker-news
---

# When Alexa can't answer, Claude takes over

- HN: [50400738](https://news.ycombinator.com/item?id=50400738)
- Source: [github.com](https://github.com/rg321/hey-claude)
- Score: 2
- Comments: 0
- Posted: 2026-11-07T01:29:23Z

## Translation

Title: When Alexa can't answer, Claude takes over
Article title: GitHub - rg321/hey-claude: When Alexa can't answer, Claude takes over. AI home assistant powered by Claude. · GitHub
Description: When Alexa can't answer, Claude takes over. AI home assistant powered by Claude. - rg321/hey-claude

Article text:
GitHub - rg321/hey-claude: When Alexa can't answer, Claude takes over. AI home assistant powered by Claude. · GitHub
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
rg321
hey-claude Public Notifications You must be signed in to change notification settings
Star 34 ( 34 ) You must be signed in to star a repository
When Alexa can't answer, Claude takes over. AI home assistant powered by Claude.
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
59 Commits 59 Commits Folders and files
.github/ workflows .github/ workflows core core devices devices .gitignore .gitignore BACKLOG.md BACKLOG.md CLAUDE.md.example CLAUDE.md.example README.md README.md architecture.png architecture.png config.example.json config.example.json package-lock.json package-lock.json package.json package.json View all files Repository files navigation
🪔 "navratri shuru ho rahe hain, mujhe iska matlab batao" — Alexa had no clue. Claude explained the festival.
📹 "mummy ke tv pe cctv chalaa do, pahle usko on kar lena" — Woke up TV, started camera stream, cast it. All in one go.
📺 "tv pe netflix chalaao aur koee badhiya movie play karo" — Powered on TV, launched Netflix, and confirmed.
😂 "mainu ek punjabi chutkula suna de yaar" — Told a joke. In Punjabi.
When Alexa can't answer, Claude takes over.
A lightweight AI home assistant that catches failed Alexa commands, routes them to Claude, and executes them — controlling your TV, speakers, cameras, and more.
One Claude session per command — Fast, cheap (Haiku), no long-running process
CLAUDE.md as memory — Claude reads it every session, learns across sessions by writing back to it
Plug-and-play devices — Drop a folder in devices/ , write a device.md , and Claude auto-discovers it
Failure blacklist in poller — Only routes genuinely failed commands to Claude; successful Alexa responses are skipped
60s dedup window — Prevents feedback loops when Claude's actions trigger new Alexa events
Dynamic response delay — Waits for Alexa to finish speaking (~20 chars/sec) before Claude responds
Processing beeps — Audio feedback while Claude is working so you know it's alive
Alexa can't handle it ("Sorry, I don't know that")
A poller watches Alexa's voice history every 2s and catches the failure
A watcher (100ms polling) detects the new command and spawns a Claude Haiku session
Claude understands the intent, executes the action (TV on/off, launch apps, play music, etc.)
The response is spoken back through Alexa in a male voice (Matthew via SSML)
While Claude is processing, you hear a beep sound so you know it's working.
Poller ( core/poller.js ) — Connects to Alexa via alexa-remote2 , pulls voice history, filters failed commands
Watcher ( core/watch.sh ) — Detects new failed commands, handles deduplication, spawns Claude sessions
Claude sessions — One Haiku session per command, reads CLAUDE.md for context, executes via bash
LG TV control ( devices/lg-webos-tv/lg_tv.js ) — Direct WebSocket (SSAP) control: power, volume, apps, screenshots
Alexa control ( core/alexa/ ) — Speak, volume, text commands, announcements
CCTV streaming ( devices/cpplus-dvr/stream.sh ) — Live camera feed on TV: DVR (RTSP) → ffmpeg (HLS) → TV browser
Jio STB ( devices/jio-stb/ ) — DLNA casting, YouTube/Netflix via DIAL
├── core/
│ ├── poller.js # Polls Alexa voice history, filters failures
│ ├── watch.sh # Watches for commands, spawns Claude, auto-loads device docs
│ └── alexa/ # Alexa voice interface
│ ├── auth.js # Authentication (generates cookie_data.json)
│ └── control.js # speak, announce, textcommand, volume
│
├── devices/ # Plug-and-play device integrations
│ ├── lg-webos-tv/
│ │ ├── lg_tv.js # WebSocket (SSAP) control
│ │ └── device.md # Instructions for Claude (auto-included)
│ ├── cpplus-dvr/
│ │ ├── stream.sh # RTSP → HLS → TV browser
│ │ ├── dvr.sh # Snapshot, record, motion detection
│ │ ├── player.html # Fullscreen video player
│ │ └── device.md # Instructions for Claude (auto-included)
│ ├── jio-stb/
│ │ ├── cast.sh # DLNA casting, volume, CCTV
│ │ ├── dial.sh # YouTube, Netflix via DIAL
│ │ └── device.md # Instructions for Claude (auto-included)
│
├── config.json # Your device IPs, MACs, network (gitignored)
├── config.example.json # Template
├── CLAUDE.md # Your preferences + learnings (gitignored)
├── CLAUDE.md.example # Template
├── STATE.md # Device state (gitignored)
└── README.md
Adding new devices
Create a folder under devices/ (e.g. devices/philips-hue/ )
Add your control script (e.g. index.js or control.sh )
Write a device.md describing the commands Claude can use
Add device config to config.json
Done — watch.sh auto-includes your device.md in every Claude session
Claude CLI installed and authenticated
Devices on the same local network
git clone https://github.com/rg321/hey-claude.git
cd hey-claude
cd core/alexa && npm install && cd ../..
2. Authenticate with Alexa
cd core/alexa
node auth.js
# Opens a browser — log in with your Amazon account
# Saves cookie_data.json (gitignored)
3. Configure devices
cp config.example.json config.json
Edit config.json with your device IPs, MAC addresses, and network details. This file is gitignored.
cp CLAUDE.md.example CLAUDE.md
Add your personal preferences (language, device nicknames, aliases). Device instructions are auto-loaded from devices/*/device.md — no need to add them here.
# Start the poller (watches Alexa voice history)
nohup node core/poller.js > poller.log 2>&1 &
# Start the watcher (spawns Claude sessions for failed commands)
nohup bash core/watch.sh > /dev/null 2>&1 &
Now speak to Alexa. If she can't handle it, Claude will.
Device
Protocol
What it can do
LG WebOS TV
WebSocket (SSAP)
Power on/off, volume, launch apps (YouTube, Netflix, Spotify), screenshots, screen on/off
Alexa Echo Dot
alexa-remote2
Speak, play music, text commands, announcements
CP Plus DVR
RTSP + HTTP
Live camera feed on TV, snapshots, record clips, motion detection
Jio Set Top Box
DLNA + DIAL
Cast media, YouTube, Netflix, volume, CCTV streaming
Smart Lights
—
Planned
CCTV on TV — Say "show camera on TV" and it spins up a live feed: DVR → ffmpeg (HLS) → fullscreen TV browser. Auto-stops when you switch away.
Contributions are welcome! Feel free to open issues or submit PRs — whether it's adding support for new devices, improving the failure detection, or making the pipeline faster.
When Alexa can't answer, Claude takes over. AI home assistant powered by Claude.
4 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

When Alexa can't answer, Claude takes over. AI home assistant powered by Claude. - rg321/hey-claude

GitHub - rg321/hey-claude: When Alexa can't answer, Claude takes over. AI home assistant powered by Claude. · GitHub
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
rg321
hey-claude Public Notifications You must be signed in to change notification settings
Star 34 ( 34 ) You must be signed in to star a repository
When Alexa can't answer, Claude takes over. AI home assistant powered by Claude.
4 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
59 Commits 59 Commits Folders and files
.github/ workflows .github/ workflows core core devices devices .gitignore .gitignore BACKLOG.md BACKLOG.md CLAUDE.md.example CLAUDE.md.example README.md README.md architecture.png architecture.png config.example.json config.example.json package-lock.json package-lock.json package.json package.json View all files Repository files navigation
🪔 "navratri shuru ho rahe hain, mujhe iska matlab batao" — Alexa had no clue. Claude explained the festival.
📹 "mummy ke tv pe cctv chalaa do, pahle usko on kar lena" — Woke up TV, started camera stream, cast it. All in one go.
📺 "tv pe netflix chalaao aur koee badhiya movie play karo" — Powered on TV, launched Netflix, and confirmed.
😂 "mainu ek punjabi chutkula suna de yaar" — Told a joke. In Punjabi.
When Alexa can't answer, Claude takes over.
A lightweight AI home assistant that catches failed Alexa commands, routes them to Claude, and executes them — controlling your TV, speakers, cameras, and more.
One Claude session per command — Fast, cheap (Haiku), no long-running process
CLAUDE.md as memory — Claude reads it every session, learns across sessions by writing back to it
Plug-and-play devices — Drop a folder in devices/ , write a device.md , and Claude auto-discovers it
Failure blacklist in poller — Only routes genuinely failed commands to Claude; successful Alexa responses are skipped
60s dedup window — Prevents feedback loops when Claude's actions trigger new Alexa events
Dynamic response delay — Waits for Alexa to finish speaking (~20 chars/sec) before Claude responds
Processing beeps — Audio feedback while Claude is working so you know it's alive
Alexa can't handle it ("Sorry, I don't know that")
A poller watches Alexa's voice history every 2s and catches the failure
A watcher (100ms polling) detects the new command and spawns a Claude Haiku session
Claude understands the intent, executes the action (TV on/off, launch apps, play music, etc.)
The response is spoken back through Alexa in a male voice (Matthew via SSML)
While Claude is processing, you hear a beep sound so you know it's working.
Poller ( core/poller.js ) — Connects to Alexa via alexa-remote2 , pulls voice history, filters failed commands
Watcher ( core/watch.sh ) — Detects new failed commands, handles deduplication, spawns Claude sessions
Claude sessions — One Haiku session per command, reads CLAUDE.md for context, executes via bash
LG TV control ( devices/lg-webos-tv/lg_tv.js ) — Direct WebSocket (SSAP) control: power, volume, apps, screenshots
Alexa control ( core/alexa/ ) — Speak, volume, text commands, announcements
CCTV streaming ( devices/cpplus-dvr/stream.sh ) — Live camera feed on TV: DVR (RTSP) → ffmpeg (HLS) → TV browser
Jio STB ( devices/jio-stb/ ) — DLNA casting, YouTube/Netflix via DIAL
├── core/
│ ├── poller.js # Polls Alexa voice history, filters failures
│ ├── watch.sh # Watches for commands, spawns Claude, auto-loads device docs
│ └── alexa/ # Alexa voice interface
│ ├── auth.js # Authentication (generates cookie_data.json)
│ └── control.js # speak, announce, textcommand, volume
│
├── devices/ # Plug-and-play device integrations
│ ├── lg-webos-tv/
│ │ ├── lg_tv.js # WebSocket (SSAP) control
│ │ └── device.md # Instructions for Claude (auto-included)
│ ├── cpplus-dvr/
│ │ ├── stream.sh # RTSP → HLS → TV browser
│ │ ├── dvr.sh # Snapshot, record, motion detection
│ │ ├── player.html # Fullscreen video player
│ │ └── device.md # Instructions for Claude (auto-included)
│ ├── jio-stb/
│ │ ├── cast.sh # DLNA casting, volume, CCTV
│ │ ├── dial.sh # YouTube, Netflix via DIAL
│ │ └── device.md # Instructions for Claude (auto-included)
│
├── config.json # Your device IPs, MACs, network (gitignored)
├── config.example.json # Template
├── CLAUDE.md # Your preferences + learnings (gitignored)
├── CLAUDE.md.example # Template
├── STATE.md # Device state (gitignored)
└── README.md
Adding new devices
Create a folder under devices/ (e.g. devices/philips-hue/ )
Add your control script (e.g. index.js or control.sh )
Write a device.md describing the commands Claude can use
Add device config to config.json
Done — watch.sh auto-includes your device.md in every Claude session
Claude CLI installed and authenticated
Devices on the same local network
git clone https://github.com/rg321/hey-claude.git
cd hey-claude
cd core/alexa && npm install && cd ../..
2. Authenticate with Alexa
cd core/alexa
node auth.js
# Opens a browser — log in with your Amazon account
# Saves cookie_data.json (gitignored)
3. Configure devices
cp config.example.json config.json
Edit config.json with your device IPs, MAC addresses, and network details. This file is gitignored.
cp CLAUDE.md.example CLAUDE.md
Add your personal preferences (language, device nicknames, aliases). Device instructions are auto-loaded from devices/*/device.md — no need to add them here.
# Start the poller (watches Alexa voice history)
nohup node core/poller.js > poller.log 2>&1 &
# Start the watcher (spawns Claude sessions for failed commands)
nohup bash core/watch.sh > /dev/null 2>&1 &
Now speak to Alexa. If she can't handle it, Claude will.
Device
Protocol
What it can do
LG WebOS TV
WebSocket (SSAP)
Power on/off, volume, launch apps (YouTube, Netflix, Spotify), screenshots, screen on/off
Alexa Echo Dot
alexa-remote2
Speak, play music, text commands, announcements
CP Plus DVR
RTSP + HTTP
Live camera feed on TV, snapshots, record clips, motion detection
Jio Set Top Box
DLNA + DIAL
Cast media, YouTube, Netflix, volume, CCTV streaming
Smart Lights
—
Planned
CCTV on TV — Say "show camera on TV" and it spins up a live feed: DVR → ffmpeg (HLS) → fullscreen TV browser. Auto-stops when you switch away.
Contributions are welcome! Feel free to open issues or submit PRs — whether it's adding support for new devices, improving the failure detection, or making the pipeline faster.
When Alexa can't answer, Claude takes over. AI home assistant powered by Claude.
4 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
