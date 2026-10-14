---
source: "https://evlat.kalaomer.com/"
hn_url: "https://news.ycombinator.com/item?id=50087078"
title: "Show HN: Evlat, see which of your AI agents waiting on you"
article_title: "Evlat: which agent is waiting on you?"
image: "https://evlat.kalaomer.com/og.png"
author: "tdgunes"
captured_at: "2026-10-14T13:21:20Z"
capture_tool: "hn-digest"
hn_id: 50087078
score: 2
comments: 0
posted_at: "2026-10-14T13:00:23Z"
tags:
  - hacker-news
---

# Show HN: Evlat, see which of your AI agents waiting on you

- HN: [50087078](https://news.ycombinator.com/item?id=50087078)
- Source: [evlat.kalaomer.com](https://evlat.kalaomer.com/)
- Score: 2
- Comments: 0
- Posted: 2026-10-14T13:00:23Z

## Translation

Title: Show HN: Evlat, see which of your AI agents waiting on you
Article title: Evlat: which agent is waiting on you?
Description: Evlat docks a thin black strip to the edge of your Mac

Article text:
Skip to content
evlat
How it works
Works with
FAQ
Docs
Changelog
EN
Download
Which of your agents is waiting on you?
A thin strip on the edge of your Mac’s screen. One click takes you straight back to whichever agent needs you.
Download for Mac
Source on GitHub
macOS 14+ · Apple silicon · free · no permissions
Projects
roadmap.md
screenshot.png
Works with
Claude Code waiting for approval
Anything you run evlat watch · 62%
Not just agents. Anything that takes a while.
Start the long one (an image build, a training run, the full e2e suite, a deploy) with evlat watch in front, then go do something else. It keeps a ring on the bar while it runs: green when it exits 0, red with the exit code when it doesn’t. Your terminal doesn’t notice a thing.
sped up; in real life these take minutes to hours
Runs it exactly as it is: colours, input, Ctrl-C and the exit code all pass through. If Evlat isn’t running, the command simply runs.
evlat signal from your own scripts
For work that knows how far along it is: --progress 0.4 fills the ring, --waiting , --done or --failed set its state. Rows expire on their own.
The same command works on a server you reach over ssh: its jobs arrive tagged with the machine’s name. How servers connect
A training run on the GPU box, Claude Code on the devbox, Codex on the build server. Their rows sit on your Mac’s bar next to your local ones, each tagged with the machine it runs on. Evlat reaches them over the ssh you already use.
gpu-01
devbox
build-01
Tagged by machine
Every row carries the name of the machine it runs on. An agent waiting on a server pulses amber on your Mac, the same as one on your laptop. For Claude Code, Codex and OpenCode you answer it right on the bar; anything else, in that server’s terminal.
When a server can’t be heard, its rows dim and say no connection instead of showing a stale state. Evlat keeps retrying, and they light up again once it’s back.
Add a server by the name you give ssh , such as devbox or me@server . Evlat carries a socket on the server to your Mac: no port opened, no account, no relay in between. It can set up the hooks and the evlat command there, or hand you one block to run yourself.
Settings → Servers · passwordless ssh · Claude Code, Codex, Antigravity and OpenCode
Evlat hooks into your agents, and every session reports to it as it goes. Each one is a ring on the strip at your screen’s edge: turning while the agent works, pulsing amber when it needs you, green when it’s done. Everything else waits for your cursor.
The strip grows into the screen: every session’s name, what it’s doing and for how long. Waiting ones sit on top, your Claude and Codex limits underneath.
A card slides out with what it’s stuck on: the exact command it wants to run or Claude’s question, and for Claude Code and OpenCode the buttons to answer right there. Open in Terminal finds the app that owns the session (a terminal, VS Code, Orca) and brings it to the front. No permissions asked.
Or press ⇧ ⌘ Space . A balloon opens right out of the bar and takes your keyboard without pulling Evlat to the front. It runs on your own Claude Code install, so it knows what your terminal knows.
Its eyes lock onto the file as it comes closer, and the file lands in the chat.
Yes. There is no paid tier and no account. The source is on GitHub under the Functional Source License ( FSL-1.1-ALv2 ).
Yes. The license allows any use, including inside a company, except offering Evlat or something substantially like it as a competing commercial product. Each version also becomes Apache 2.0 two years after its release.
No. Evlat has no account and no telemetry. Sessions, commands and usage stay on your Mac, and the hooks talk to a listener on 127.0.0.1 . The only thing it fetches from the internet is its own update, from GitHub, which it looks for every hour. The chat balloon runs your own Claude Code, so what you ask there goes wherever your Claude Code already sends it. More in the privacy notice .
Claude Code, Codex, Antigravity (its app, IDE and agy CLI) and OpenCode, through the hooks and plugin Evlat installs for you. What each one shows is on What works where . Anything else (a build, a test run, a deploy, your own script) shows up with evlat watch or evlat signal .
Evlat adds its own entries to Claude Code’s ~/.claude/settings.json and Codex’s ~/.codex/hooks.json and Antigravity’s ~/.gemini/config/hooks.json , next to any hooks you already have, and before the first write it keeps a .evlat.bak copy beside each one. For OpenCode it writes one file of its own, ~/.config/opencode/plugins/evlat.js . No other file is touched. To uninstall, remove the hooks in Evlat’s settings, quit it and move it to the Trash.
macOS 14 Sonoma or later, on Apple silicon. There is no Intel build and no Windows or Linux app. The servers it watches can be any machine you reach over ssh.
Every release is signed and notarised by Apple. Evlat checks GitHub for a new version every hour and updates through Sparkle, which verifies each update against Evlat’s signing key before installing it. By default a new version downloads in the background and installs when Evlat quits; you can turn that off in Settings. Release notes
In the GitHub issues . The source, the license and every release are in the same repository.
Open Evlat.dmg and drag Evlat into Applications.
It is signed and notarised, so it opens like any other app. It keeps itself up to date: Check for Updates… is in its menu.
Right-click the mascot → install hooks for Claude Code, Codex, Antigravity and OpenCode. New sessions start reporting straight away; the ones already open show up too.
Download Evlat
Read the source
3 MB · macOS 14+ · Apple silicon · updates itself
Evlat is Turkish for “child”. It keeps an eye on yours.

## Original Extract

Evlat docks a thin black strip to the edge of your Mac

Skip to content
evlat
How it works
Works with
FAQ
Docs
Changelog
EN
Download
Which of your agents is waiting on you?
A thin strip on the edge of your Mac’s screen. One click takes you straight back to whichever agent needs you.
Download for Mac
Source on GitHub
macOS 14+ · Apple silicon · free · no permissions
Projects
roadmap.md
screenshot.png
Works with
Claude Code waiting for approval
Anything you run evlat watch · 62%
Not just agents. Anything that takes a while.
Start the long one (an image build, a training run, the full e2e suite, a deploy) with evlat watch in front, then go do something else. It keeps a ring on the bar while it runs: green when it exits 0, red with the exit code when it doesn’t. Your terminal doesn’t notice a thing.
sped up; in real life these take minutes to hours
Runs it exactly as it is: colours, input, Ctrl-C and the exit code all pass through. If Evlat isn’t running, the command simply runs.
evlat signal from your own scripts
For work that knows how far along it is: --progress 0.4 fills the ring, --waiting , --done or --failed set its state. Rows expire on their own.
The same command works on a server you reach over ssh: its jobs arrive tagged with the machine’s name. How servers connect
A training run on the GPU box, Claude Code on the devbox, Codex on the build server. Their rows sit on your Mac’s bar next to your local ones, each tagged with the machine it runs on. Evlat reaches them over the ssh you already use.
gpu-01
devbox
build-01
Tagged by machine
Every row carries the name of the machine it runs on. An agent waiting on a server pulses amber on your Mac, the same as one on your laptop. For Claude Code, Codex and OpenCode you answer it right on the bar; anything else, in that server’s terminal.
When a server can’t be heard, its rows dim and say no connection instead of showing a stale state. Evlat keeps retrying, and they light up again once it’s back.
Add a server by the name you give ssh , such as devbox or me@server . Evlat carries a socket on the server to your Mac: no port opened, no account, no relay in between. It can set up the hooks and the evlat command there, or hand you one block to run yourself.
Settings → Servers · passwordless ssh · Claude Code, Codex, Antigravity and OpenCode
Evlat hooks into your agents, and every session reports to it as it goes. Each one is a ring on the strip at your screen’s edge: turning while the agent works, pulsing amber when it needs you, green when it’s done. Everything else waits for your cursor.
The strip grows into the screen: every session’s name, what it’s doing and for how long. Waiting ones sit on top, your Claude and Codex limits underneath.
A card slides out with what it’s stuck on: the exact command it wants to run or Claude’s question, and for Claude Code and OpenCode the buttons to answer right there. Open in Terminal finds the app that owns the session (a terminal, VS Code, Orca) and brings it to the front. No permissions asked.
Or press ⇧ ⌘ Space . A balloon opens right out of the bar and takes your keyboard without pulling Evlat to the front. It runs on your own Claude Code install, so it knows what your terminal knows.
Its eyes lock onto the file as it comes closer, and the file lands in the chat.
Yes. There is no paid tier and no account. The source is on GitHub under the Functional Source License ( FSL-1.1-ALv2 ).
Yes. The license allows any use, including inside a company, except offering Evlat or something substantially like it as a competing commercial product. Each version also becomes Apache 2.0 two years after its release.
No. Evlat has no account and no telemetry. Sessions, commands and usage stay on your Mac, and the hooks talk to a listener on 127.0.0.1 . The only thing it fetches from the internet is its own update, from GitHub, which it looks for every hour. The chat balloon runs your own Claude Code, so what you ask there goes wherever your Claude Code already sends it. More in the privacy notice .
Claude Code, Codex, Antigravity (its app, IDE and agy CLI) and OpenCode, through the hooks and plugin Evlat installs for you. What each one shows is on What works where . Anything else (a build, a test run, a deploy, your own script) shows up with evlat watch or evlat signal .
Evlat adds its own entries to Claude Code’s ~/.claude/settings.json and Codex’s ~/.codex/hooks.json and Antigravity’s ~/.gemini/config/hooks.json , next to any hooks you already have, and before the first write it keeps a .evlat.bak copy beside each one. For OpenCode it writes one file of its own, ~/.config/opencode/plugins/evlat.js . No other file is touched. To uninstall, remove the hooks in Evlat’s settings, quit it and move it to the Trash.
macOS 14 Sonoma or later, on Apple silicon. There is no Intel build and no Windows or Linux app. The servers it watches can be any machine you reach over ssh.
Every release is signed and notarised by Apple. Evlat checks GitHub for a new version every hour and updates through Sparkle, which verifies each update against Evlat’s signing key before installing it. By default a new version downloads in the background and installs when Evlat quits; you can turn that off in Settings. Release notes
In the GitHub issues . The source, the license and every release are in the same repository.
Open Evlat.dmg and drag Evlat into Applications.
It is signed and notarised, so it opens like any other app. It keeps itself up to date: Check for Updates… is in its menu.
Right-click the mascot → install hooks for Claude Code, Codex, Antigravity and OpenCode. New sessions start reporting straight away; the ones already open show up too.
Download Evlat
Read the source
3 MB · macOS 14+ · Apple silicon · updates itself
Evlat is Turkish for “child”. It keeps an eye on yours.
