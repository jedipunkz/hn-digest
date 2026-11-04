---
source: "https://github.com/akustikrausch/amigent"
hn_url: "https://news.ycombinator.com/item?id=50372207"
title: "Show HN: Amigent – AI agent and remote desktop for Commodore Amiga over 50 KB/s"
article_title: "GitHub - akustikrausch/amigent: Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux. · GitHub"
image: "https://opengraph.githubassets.com/33b7daec45f0de4988c629ab98f58501c53acc7edafc47bd5f49ae7796062ee3/akustikrausch/amigent"
author: "akustikrausch"
captured_at: "2026-11-04T22:27:33Z"
capture_tool: "hn-digest"
hn_id: 50372207
score: 2
comments: 0
posted_at: "2026-11-04T22:07:35Z"
tags:
  - hacker-news
---

# Show HN: Amigent – AI agent and remote desktop for Commodore Amiga over 50 KB/s

- HN: [50372207](https://news.ycombinator.com/item?id=50372207)
- Source: [github.com](https://github.com/akustikrausch/amigent)
- Score: 2
- Comments: 0
- Posted: 2026-11-04T22:07:35Z

## Translation

Title: Show HN: Amigent – AI agent and remote desktop for Commodore Amiga over 50 KB/s
Article title: GitHub - akustikrausch/amigent: Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux. · GitHub
Description: Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux. - akustikrausch/amigent

Article text:
GitHub - akustikrausch/amigent: Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
akustikrausch
amigent Public Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
139 Commits 139 Commits Folders and files
.github/ workflows .github/ workflows amiga amiga docs docs scripts scripts src/ amigent src/ amigent tests tests .gitignore .gitignore .mcp.json.example .mcp.json.example CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
_____ .__ __
/ _ \ _____ |__| ____ ____ ____ _/ |_
/ /_\ \ / \ | | / ___\ _/ __ \ / \ \ __\
/ | \| Y Y \| | / /_/ >\ ___/ | | \ | |
\____|__ /|__|_| /|__| \___ / \___ >|___| / |__|
\/ \/ /_____/ \/ \/
complete remote control for a real Amiga, by AI or manually
tuned for slow Amigas and slow networks
Amigent connects your PC or Mac to your real Amiga over the home network. Then you can:
let an AI agent do the work , for example Claude Code or Codex: install
software, update programs, change settings, copy files, click through
programs. Just tell it what you want.
use the Amiga yourself from your computer , as a remote desktop like
TeamViewer or VNC: a window with the live screen, your own mouse and
keyboard, file transfer and a shared clipboard.
and both of these quickly, even on a slow machine. Amigent is tuned for
slow Amigas and slow networks: a 7 MHz 68000 sends its whole screen in
under a second, and over a plipbox a click shows its result in a fraction of a second.
How it stays fast .
Both at the same time: you see every click the agent makes and can take over at
any moment.
1. Copy amigentd to the Amiga ( amigentd.000 on a
68000), or unpack the archive from the
releases .
2. Start it with amigentd SETUP in a Shell, or double-click its icon.
3. Paste this into your AI agent:
Install Amigent for me: run "pipx install amigent" (if pipx or Python 3.10
or newer is missing, install them first: Windows winget, macOS Homebrew,
Linux the package manager). Then run amigent-setup; it finds my Amiga on
the network by itself. Register the MCP server "amigent" with the command
amigent-server, and run amigent-viewer --install so the remote desktop
shows up as a program. Tell me if you need to be restarted.
The agent sets up the Amiga over the network; from then on amigentd starts
with it.
Windows: Amigent appears in the Start menu. For an encrypted connection
openssl is needed, which comes with Git for Windows.
macOS: Amigent appears in ~/Applications . On its first start, allow it
access to the local network, or it cannot reach the Amiga. Cmd is the right
Amiga key.
Linux: there is no program entry; start the remote desktop with
amigent-viewer .
Without an AI agent: see the manual installation .
An Amiga with AmigaOS 2.04 or newer, any CPU from the 68000 on, a network connection
(a plipbox on the parallel port is enough) and a TCP/IP stack such as Roadshow.
A PC or Mac with Windows, macOS or Linux and Python 3.10 or newer.
"Check which programs on the Amiga are outdated."
"Install the game from my Downloads folder and add it to the Games menu."
"Copy the demo to the Amiga and start it."
"Open the remote desktop."
Made for slow machines and slow wires
Most of the work in Amigent went into making a slow Amiga on a slow network
feel quick, for the AI agent and for you at the mouse:
The screen travels as the Amiga keeps it. Chipset screens go out as
bitplanes, graphics-card screens as small palette tiles, and only the tiles
that changed, often as the difference to what you already have.
Your input comes first. Mouse moves, clicks and keys are carried out
even while a picture is still on its way, and the area around the pointer is
sent before the rest.
The Amiga keeps its time. While nothing happens, Amigent asks rarely
and costs the Amiga only a few percent of its CPU.
Files go in one go. Whole folders travel in a single stream, and files
that are already there stay where they are.
Measured with a 640×256 Workbench screen:
On a 1920×1080 graphics-card screen over a plipbox (about 50 KB/s):
The orange bars are what moving the whole screen as raw 24-bit pixels costs on
that wire, every time; the blue ones are Amigent, measured on a real A1200.
All numbers, how they were taken and how to repeat them are in
docs/PERFORMANCE.md .
New address after a restart? Amigent finds the Amiga again by itself. It
takes only the one that proves to be yours.
Several Amigas? amigent-fleet shows them side by side with a preview.
After a big OS update , amigent-repair puts back what the update removed.
Safe at home. A key keeps other computers on the network out, and the
connection can be encrypted. Don't open the Amiga to the internet.
More in the guide and the FAQ .
MIT, see LICENSE . By Andreas Wendorf (Akustikrausch).
Changes are in the changelog .
Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux. - akustikrausch/amigent

GitHub - akustikrausch/amigent: Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
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
akustikrausch
amigent Public Notifications You must be signed in to change notification settings
Star 9 ( 9 ) You must be signed in to star a repository
Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
139 Commits 139 Commits Folders and files
.github/ workflows .github/ workflows amiga amiga docs docs scripts scripts src/ amigent src/ amigent tests tests .gitignore .gitignore .mcp.json.example .mcp.json.example CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
_____ .__ __
/ _ \ _____ |__| ____ ____ ____ _/ |_
/ /_\ \ / \ | | / ___\ _/ __ \ / \ \ __\
/ | \| Y Y \| | / /_/ >\ ___/ | | \ | |
\____|__ /|__|_| /|__| \___ / \___ >|___| / |__|
\/ \/ /_____/ \/ \/
complete remote control for a real Amiga, by AI or manually
tuned for slow Amigas and slow networks
Amigent connects your PC or Mac to your real Amiga over the home network. Then you can:
let an AI agent do the work , for example Claude Code or Codex: install
software, update programs, change settings, copy files, click through
programs. Just tell it what you want.
use the Amiga yourself from your computer , as a remote desktop like
TeamViewer or VNC: a window with the live screen, your own mouse and
keyboard, file transfer and a shared clipboard.
and both of these quickly, even on a slow machine. Amigent is tuned for
slow Amigas and slow networks: a 7 MHz 68000 sends its whole screen in
under a second, and over a plipbox a click shows its result in a fraction of a second.
How it stays fast .
Both at the same time: you see every click the agent makes and can take over at
any moment.
1. Copy amigentd to the Amiga ( amigentd.000 on a
68000), or unpack the archive from the
releases .
2. Start it with amigentd SETUP in a Shell, or double-click its icon.
3. Paste this into your AI agent:
Install Amigent for me: run "pipx install amigent" (if pipx or Python 3.10
or newer is missing, install them first: Windows winget, macOS Homebrew,
Linux the package manager). Then run amigent-setup; it finds my Amiga on
the network by itself. Register the MCP server "amigent" with the command
amigent-server, and run amigent-viewer --install so the remote desktop
shows up as a program. Tell me if you need to be restarted.
The agent sets up the Amiga over the network; from then on amigentd starts
with it.
Windows: Amigent appears in the Start menu. For an encrypted connection
openssl is needed, which comes with Git for Windows.
macOS: Amigent appears in ~/Applications . On its first start, allow it
access to the local network, or it cannot reach the Amiga. Cmd is the right
Amiga key.
Linux: there is no program entry; start the remote desktop with
amigent-viewer .
Without an AI agent: see the manual installation .
An Amiga with AmigaOS 2.04 or newer, any CPU from the 68000 on, a network connection
(a plipbox on the parallel port is enough) and a TCP/IP stack such as Roadshow.
A PC or Mac with Windows, macOS or Linux and Python 3.10 or newer.
"Check which programs on the Amiga are outdated."
"Install the game from my Downloads folder and add it to the Games menu."
"Copy the demo to the Amiga and start it."
"Open the remote desktop."
Made for slow machines and slow wires
Most of the work in Amigent went into making a slow Amiga on a slow network
feel quick, for the AI agent and for you at the mouse:
The screen travels as the Amiga keeps it. Chipset screens go out as
bitplanes, graphics-card screens as small palette tiles, and only the tiles
that changed, often as the difference to what you already have.
Your input comes first. Mouse moves, clicks and keys are carried out
even while a picture is still on its way, and the area around the pointer is
sent before the rest.
The Amiga keeps its time. While nothing happens, Amigent asks rarely
and costs the Amiga only a few percent of its CPU.
Files go in one go. Whole folders travel in a single stream, and files
that are already there stay where they are.
Measured with a 640×256 Workbench screen:
On a 1920×1080 graphics-card screen over a plipbox (about 50 KB/s):
The orange bars are what moving the whole screen as raw 24-bit pixels costs on
that wire, every time; the blue ones are Amigent, measured on a real A1200.
All numbers, how they were taken and how to repeat them are in
docs/PERFORMANCE.md .
New address after a restart? Amigent finds the Amiga again by itself. It
takes only the one that proves to be yours.
Several Amigas? amigent-fleet shows them side by side with a preview.
After a big OS update , amigent-repair puts back what the update removed.
Safe at home. A key keeps other computers on the network out, and the
connection can be encrypted. Don't open the Amiga to the internet.
More in the guide and the FAQ .
MIT, see LICENSE . By Andreas Wendorf (Akustikrausch).
Changes are in the changelog .
Complete remote control for a real Amiga, by AI or manually: an MCP server for AI agents and a remote desktop like TeamViewer, tuned for slow Amigas and slow networks. AmigaOS 2.04 and up, any CPU from the 68000, Windows, macOS and Linux.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
