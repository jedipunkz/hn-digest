---
source: "https://github.com/delexw/miod"
hn_url: "https://news.ycombinator.com/item?id=49999983"
title: "A Claude Code mod plays MIDI music when it works"
article_title: "GitHub - delexw/miod: Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens. · GitHub"
image: "https://opengraph.githubassets.com/c39da5db9e8ec4147017eff99242f779d0cc05b7e432a7bf80107b89c8f716d3/delexw/miod"
author: "delexw"
captured_at: "2026-10-07T23:28:00Z"
capture_tool: "hn-digest"
hn_id: 49999983
score: 3
comments: 1
posted_at: "2026-10-07T23:06:24Z"
tags:
  - hacker-news
---

# A Claude Code mod plays MIDI music when it works

- HN: [49999983](https://news.ycombinator.com/item?id=49999983)
- Source: [github.com](https://github.com/delexw/miod)
- Score: 3
- Comments: 1
- Posted: 2026-10-07T23:06:24Z

## Translation

Title: A Claude Code mod plays MIDI music when it works
Article title: GitHub - delexw/miod: Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens. · GitHub
Description: Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens. - delexw/miod

Article text:
GitHub - delexw/miod: Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens. · GitHub
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
delexw
miod Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
.claude-plugin .claude-plugin docs docs hooks hooks tests tests types types .gitignore .gitignore LICENSE LICENSE README.md README.md tsconfig.json tsconfig.json View all files Repository files navigation
Only tested on macOS. Not sure yet how it behaves on Linux or Windows.
I'm an engineer, but these days I solve the boring problems entirely by vibe coding. Somewhere along the way I felt I had lost the creativity and curiosity I used to have when writing code myself.
So I started thinking about how to make vibe coding fun again. miod is the first try: while the agent does the work, you get to hear it think, read, edit, fail and fix, with a little piece of music no one has heard before.
Feel free to add more moods, flavours and styles. Each mood is one row in hooks/moods.ts , each flavour one row in hooks/flavours.ts , each sound style one entry in hooks/styles.ts , and each wave look one file in hooks/waves/ .
Feeling it's a bit too noisy? 🙉 No hard feelings. Just turn it off with /miod off .
A Claude Code mod (MIDI + mod) that writes music on the spot while Claude works. No music files, no library, no AI model.
Every task picks its own flavour (lo-fi, chiptune, ambient, jazzy and 8 more), key, scale (16 of them), chord progression, tempo, sound and theme.
The mood follows what the agent is doing. Every mood has 10 variants, and a new one is picked as the task goes: whenever the activity changes, and every 4 phrases. The task's flavour bends it too, so editing in a lo-fi task sounds nothing like editing in a chiptune one.
The faster it spends tokens, the busier the music.
Above the prompt, one line shows the mood, flavour, key, tempo and energy, with a wave under it that scrolls with the music, coloured by mood:
Tasks flow into each other: the next task fades in, in a related key. If no new task comes, a closing chord fades out at the end of the phrase.
Needs Claude Code 2.1.290 or newer, on macOS.
claude plugin marketplace add delexw/miod
claude plugin install miod@miod
Then start a new session.
/miod is it on or off?
/miod off stop the music
/miod on play again from the next task
/miod wave list the wave looks: bars, line, mirror, dots, pulse
/miod wave dots switch the wave to the dots look
How the mood is picked
Each row below is the mood's first variant. All 10 are in MOODS in hooks/moods.ts .
When several happen in one 8-beat phrase, the one higher in STRONGEST_FIRST ( hooks/activity.ts ) wins. A filling context window lifts the melody and, past 80%, adds a low drone.
The scales follow how musicians describe each mode ( musical-u , Jiang et al. 2024 ).
Mood variant: add a row to that mood's list in MOODS in hooks/moods.ts .
New mood: add the name to Activity and STRONGEST_FIRST in hooks/activity.ts , map tools or commands to it there, and give it a list of at least 10 variants in MOODS .
Flavour: add a row to FLAVOURS in hooks/flavours.ts . Each field nudges every mood: brighter or darker scale, busier or calmer, drums up or down, swing, bass line.
Scale or chords: add to SCALES or PROGRESSIONS in hooks/theory.ts .
Wave look: add a file to hooks/waves/ that exports a WaveLook (how many rows, and a draw function), then list it in WAVES in hooks/waves/index.ts .
Style: add a wave function to STYLES in hooks/styles.ts . It must return values between -1 and 1.
claude plugin validate .
claude plugin test .
Run a checkout without installing it with claude --plugin-dir <path to this repo> .
Only tested on macOS, where Claude Code plays the audio through afplay .
The sound is simple chiptune-style tones.
miod only hears the session it's loaded in.
Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens. - delexw/miod

GitHub - delexw/miod: Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens. · GitHub
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
delexw
miod Public Notifications You must be signed in to change notification settings
Star 1 ( 1 ) You must be signed in to star a repository
Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens.
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
14 Commits 14 Commits Folders and files
.claude-plugin .claude-plugin docs docs hooks hooks tests tests types types .gitignore .gitignore LICENSE LICENSE README.md README.md tsconfig.json tsconfig.json View all files Repository files navigation
Only tested on macOS. Not sure yet how it behaves on Linux or Windows.
I'm an engineer, but these days I solve the boring problems entirely by vibe coding. Somewhere along the way I felt I had lost the creativity and curiosity I used to have when writing code myself.
So I started thinking about how to make vibe coding fun again. miod is the first try: while the agent does the work, you get to hear it think, read, edit, fail and fix, with a little piece of music no one has heard before.
Feel free to add more moods, flavours and styles. Each mood is one row in hooks/moods.ts , each flavour one row in hooks/flavours.ts , each sound style one entry in hooks/styles.ts , and each wave look one file in hooks/waves/ .
Feeling it's a bit too noisy? 🙉 No hard feelings. Just turn it off with /miod off .
A Claude Code mod (MIDI + mod) that writes music on the spot while Claude works. No music files, no library, no AI model.
Every task picks its own flavour (lo-fi, chiptune, ambient, jazzy and 8 more), key, scale (16 of them), chord progression, tempo, sound and theme.
The mood follows what the agent is doing. Every mood has 10 variants, and a new one is picked as the task goes: whenever the activity changes, and every 4 phrases. The task's flavour bends it too, so editing in a lo-fi task sounds nothing like editing in a chiptune one.
The faster it spends tokens, the busier the music.
Above the prompt, one line shows the mood, flavour, key, tempo and energy, with a wave under it that scrolls with the music, coloured by mood:
Tasks flow into each other: the next task fades in, in a related key. If no new task comes, a closing chord fades out at the end of the phrase.
Needs Claude Code 2.1.290 or newer, on macOS.
claude plugin marketplace add delexw/miod
claude plugin install miod@miod
Then start a new session.
/miod is it on or off?
/miod off stop the music
/miod on play again from the next task
/miod wave list the wave looks: bars, line, mirror, dots, pulse
/miod wave dots switch the wave to the dots look
How the mood is picked
Each row below is the mood's first variant. All 10 are in MOODS in hooks/moods.ts .
When several happen in one 8-beat phrase, the one higher in STRONGEST_FIRST ( hooks/activity.ts ) wins. A filling context window lifts the melody and, past 80%, adds a low drone.
The scales follow how musicians describe each mode ( musical-u , Jiang et al. 2024 ).
Mood variant: add a row to that mood's list in MOODS in hooks/moods.ts .
New mood: add the name to Activity and STRONGEST_FIRST in hooks/activity.ts , map tools or commands to it there, and give it a list of at least 10 variants in MOODS .
Flavour: add a row to FLAVOURS in hooks/flavours.ts . Each field nudges every mood: brighter or darker scale, busier or calmer, drums up or down, swing, bass line.
Scale or chords: add to SCALES or PROGRESSIONS in hooks/theory.ts .
Wave look: add a file to hooks/waves/ that exports a WaveLook (how many rows, and a draw function), then list it in WAVES in hooks/waves/index.ts .
Style: add a wave function to STYLES in hooks/styles.ts . It must return values between -1 and 1.
claude plugin validate .
claude plugin test .
Run a checkout without installing it with claude --plugin-dir <path to this repo> .
Only tested on macOS, where Claude Code plays the audio through afplay .
The sound is simple chiptune-style tones.
miod only hears the session it's loaded in.
Generative music for Claude Code. Every task gets its own tune, and the mood follows what the agent is doing and how fast it spends tokens.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
