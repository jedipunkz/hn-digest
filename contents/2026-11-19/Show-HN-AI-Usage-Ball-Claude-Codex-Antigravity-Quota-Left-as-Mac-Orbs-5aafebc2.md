---
source: "https://github.com/aiusageball/ai-usage-ball"
hn_url: "https://news.ycombinator.com/item?id=50561944"
title: "Show HN: AI Usage Ball – Claude, Codex, Antigravity Quota Left as Mac Orbs"
article_title: "GitHub - aiusageball/ai-usage-ball: Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com · GitHub"
image: "https://opengraph.githubassets.com/77ae9f5242f64c349cadea45bb9361a02aed32cd6ad466f4e8438744c8307276/aiusageball/ai-usage-ball"
author: "xbyxy2000"
captured_at: "2026-11-19T16:21:23Z"
capture_tool: "hn-digest"
hn_id: 50561944
score: 1
comments: 0
posted_at: "2026-11-19T16:09:47Z"
tags:
  - hacker-news
---

# Show HN: AI Usage Ball – Claude, Codex, Antigravity Quota Left as Mac Orbs

- HN: [50561944](https://news.ycombinator.com/item?id=50561944)
- Source: [github.com](https://github.com/aiusageball/ai-usage-ball)
- Score: 1
- Comments: 0
- Posted: 2026-11-19T16:09:47Z

## Translation

Title: Show HN: AI Usage Ball – Claude, Codex, Antigravity Quota Left as Mac Orbs
Article title: GitHub - aiusageball/ai-usage-ball: Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com · GitHub
Description: Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com - aiusageball/ai-usage-ball

Article text:
GitHub - aiusageball/ai-usage-ball: Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com · GitHub
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
aiusageball
ai-usage-ball Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com
Readme License Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
70 Commits 70 Commits Folders and files
AiPulseWatch AiPulseWatch cloudflare cloudflare dashboard dashboard docs docs landing landing scripts scripts server server watch watch .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md liquid-frame.png liquid-frame.png liquid-loop.mp4 liquid-loop.mp4 watch-mockup.html watch-mockup.html View all files Repository files navigation
A source-available macOS desktop app that shows how much of your AI coding-tool
quota — Claude, Codex / ChatGPT, and Antigravity — you have left, as animated
liquid gauges. Session limits, weekly limits, and reset countdowns, at a
glance.
Demo uses sample data and a time-lapse; the window-to-desktop fly-out is a demo effect.
Try the signed, notarized build (Apple Silicon, macOS 12+): Download for Apple Silicon
or see the website for details. Or with Homebrew:
brew tap aiusageball/aiusageball
brew trust aiusageball/aiusageball
brew install --cask ai-usage-ball
Every download includes a 30-day trial with no card required, and the first 100 users who try
the app and share genuine product feedback can apply for a free lifetime
personal license.
This repository is the full source, published so anyone can audit exactly what
the app reads from your machine and how. It is source-available for
auditability and noncommercial self-builds — see License for what
you can and can't do with it.
Path
What it is
dashboard/
The desktop app — Tauri v2 (Rust) + React frontend. Main window, desktop widgets, the liquid-orb visualization.
server/
Python (FastAPI) backend the app talks to on 127.0.0.1:8000 . Reads local usage data (Claude via browser session cookie / CLI credentials, Codex via ~/.codex/auth.json , Antigravity via local process detection) and streams it to the frontend over SSE.
AiPulseWatch/
watchOS companion (SwiftUI), work in progress.
landing/
Source for the aiusageball.com marketing site (static HTML, no build step).
Why developers use it
Keep AI coding-tool limits visible before a session unexpectedly runs out.
See reset countdowns and weekly/session windows without opening each provider.
Audit the local readers that collect usage state from your own Mac.
Run the desktop UI as a Tauri app backed by a local FastAPI service.
Everything is read locally , from sessions/credentials already on your
machine — nothing is proxied through a server we run. The backend's polling
logic lives in server/server.py ; that file is the ground truth for exactly
what's read and from where.
cd server
python3 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
python server.py # serves http://127.0.0.1:8000
Desktop app:
cd dashboard
npm install
npm run tauri:dev
Building a distributable .app / .dmg : the release build bundles the
backend into the app as a self-contained sidecar (so end users don't need
Python). Build it first, then run the Tauri build:
# 1. Freeze the backend into a single binary (PyInstaller)
cd server && source venv/bin/activate && pip install pyinstaller
pyinstaller --onefile --name aipulse-server \
--collect-all uvicorn --collect-all fastapi --collect-all zeroconf \
--collect-all browser_cookie3 --collect-all anyio \
--add-data " ../dashboard/public/liquid-loop.mp4:. " server.py
# 2. Place it where Tauri expects the sidecar (name includes the target triple)
cp dist/aipulse-server \
../dashboard/src-tauri/binaries/aipulse-server-aarch64-apple-darwin
# 3. Build the app
cd ../dashboard && npm run tauri:build
Signing/notarizing requires your own Apple Developer identity (see
dashboard/src-tauri/tauri.conf.json → bundle.macOS.signingIdentity ;
entitlements.plist carries disable-library-validation , which the
PyInstaller sidecar needs to load its embedded Python under a hardened
runtime). Unsigned local builds work fine for development.
Licensed under the PolyForm Noncommercial License 1.0.0 : you're
free to read, audit, modify, and build the source for personal / noncommercial
use. Commercial use (selling it, bundling it, offering it as a paid service,
etc.) isn't permitted under this license — get in touch if you want that.
The official signed, notarized build at aiusageball.com
supports continued development.
Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com - aiusageball/ai-usage-ball

GitHub - aiusageball/ai-usage-ball: Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com · GitHub
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
aiusageball
ai-usage-ball Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com
Readme License Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
70 Commits 70 Commits Folders and files
AiPulseWatch AiPulseWatch cloudflare cloudflare dashboard dashboard docs docs landing landing scripts scripts server server watch watch .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md liquid-frame.png liquid-frame.png liquid-loop.mp4 liquid-loop.mp4 watch-mockup.html watch-mockup.html View all files Repository files navigation
A source-available macOS desktop app that shows how much of your AI coding-tool
quota — Claude, Codex / ChatGPT, and Antigravity — you have left, as animated
liquid gauges. Session limits, weekly limits, and reset countdowns, at a
glance.
Demo uses sample data and a time-lapse; the window-to-desktop fly-out is a demo effect.
Try the signed, notarized build (Apple Silicon, macOS 12+): Download for Apple Silicon
or see the website for details. Or with Homebrew:
brew tap aiusageball/aiusageball
brew trust aiusageball/aiusageball
brew install --cask ai-usage-ball
Every download includes a 30-day trial with no card required, and the first 100 users who try
the app and share genuine product feedback can apply for a free lifetime
personal license.
This repository is the full source, published so anyone can audit exactly what
the app reads from your machine and how. It is source-available for
auditability and noncommercial self-builds — see License for what
you can and can't do with it.
Path
What it is
dashboard/
The desktop app — Tauri v2 (Rust) + React frontend. Main window, desktop widgets, the liquid-orb visualization.
server/
Python (FastAPI) backend the app talks to on 127.0.0.1:8000 . Reads local usage data (Claude via browser session cookie / CLI credentials, Codex via ~/.codex/auth.json , Antigravity via local process detection) and streams it to the frontend over SSE.
AiPulseWatch/
watchOS companion (SwiftUI), work in progress.
landing/
Source for the aiusageball.com marketing site (static HTML, no build step).
Why developers use it
Keep AI coding-tool limits visible before a session unexpectedly runs out.
See reset countdowns and weekly/session windows without opening each provider.
Audit the local readers that collect usage state from your own Mac.
Run the desktop UI as a Tauri app backed by a local FastAPI service.
Everything is read locally , from sessions/credentials already on your
machine — nothing is proxied through a server we run. The backend's polling
logic lives in server/server.py ; that file is the ground truth for exactly
what's read and from where.
cd server
python3 -m venv venv && source venv/bin/activate
pip install -r requirements.txt
python server.py # serves http://127.0.0.1:8000
Desktop app:
cd dashboard
npm install
npm run tauri:dev
Building a distributable .app / .dmg : the release build bundles the
backend into the app as a self-contained sidecar (so end users don't need
Python). Build it first, then run the Tauri build:
# 1. Freeze the backend into a single binary (PyInstaller)
cd server && source venv/bin/activate && pip install pyinstaller
pyinstaller --onefile --name aipulse-server \
--collect-all uvicorn --collect-all fastapi --collect-all zeroconf \
--collect-all browser_cookie3 --collect-all anyio \
--add-data " ../dashboard/public/liquid-loop.mp4:. " server.py
# 2. Place it where Tauri expects the sidecar (name includes the target triple)
cp dist/aipulse-server \
../dashboard/src-tauri/binaries/aipulse-server-aarch64-apple-darwin
# 3. Build the app
cd ../dashboard && npm run tauri:build
Signing/notarizing requires your own Apple Developer identity (see
dashboard/src-tauri/tauri.conf.json → bundle.macOS.signingIdentity ;
entitlements.plist carries disable-library-validation , which the
PyInstaller sidecar needs to load its embedded Python under a hardened
runtime). Unsigned local builds work fine for development.
Licensed under the PolyForm Noncommercial License 1.0.0 : you're
free to read, audit, modify, and build the source for personal / noncommercial
use. Commercial use (selling it, bundling it, offering it as a paid service,
etc.) isn't permitted under this license — get in touch if you want that.
The official signed, notarized build at aiusageball.com
supports continued development.
Live liquid-orb usage gauges for Claude, Codex & Antigravity on macOS — source for aiusageball.com
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
