---
source: "https://github.com/wyzowl/sfx-kit"
hn_url: "https://news.ycombinator.com/item?id=50166753"
title: "Show HN: SFX Kit – An AI-first sound effect kit with CC0 licence"
article_title: "GitHub - wyzowl/sfx-kit: 320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0. · GitHub"
image: "https://opengraph.githubassets.com/1108644c806959601c0eb1bf32028415747bb0171d99c74780a069f702eb21af/wyzowl/sfx-kit"
author: "mattbyrom"
captured_at: "2026-10-20T14:44:14Z"
capture_tool: "hn-digest"
hn_id: 50166753
score: 2
comments: 0
posted_at: "2026-10-20T14:11:13Z"
tags:
  - hacker-news
---

# Show HN: SFX Kit – An AI-first sound effect kit with CC0 licence

- HN: [50166753](https://news.ycombinator.com/item?id=50166753)
- Source: [github.com](https://github.com/wyzowl/sfx-kit)
- Score: 2
- Comments: 0
- Posted: 2026-10-20T14:11:13Z

## Translation

Title: Show HN: SFX Kit – An AI-first sound effect kit with CC0 licence
Article title: GitHub - wyzowl/sfx-kit: 320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0. · GitHub
Description: 320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0. - wyzowl/sfx-kit

Article text:
GitHub - wyzowl/sfx-kit: 320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0. · GitHub
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
wyzowl
sfx-kit Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0.
Readme CC0-1.0 license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
sounds sounds GUIDE.md GUIDE.md LICENSE LICENSE README.md README.md analysis.json analysis.json catalogue.md catalogue.md sounds.json sounds.json sounds.schema.json sounds.schema.json View all files Repository files navigation
Also on Hugging Face: huggingface.co/datasets/wyzowl/sfx-kit .
320 original sound effects for videos, animations and interfaces, made to be as easy for AI tools to use as for
people. Every sound comes with the data needed to choose it, place it in time and mix it: what it is, when to use it
and when not to, its length, its sync point and its level.
Free for any use, commercial or not: CC0 , no credit required.
Category
Sounds
What it's for
Whooshes
9
Movement, sliding, flying and transitions
Swishes and zips
9
Quick, sharp movement: flicks, cuts, and things zipping in or out
Swipes
6
Fingers swiping across screens: app demos, cards and carousels
Reverse whooshes
6
Air drawn in to a sudden stop: things collapsing into a point, and the build into a cut
Risers and fallers
12
Building anticipation before a reveal, and things descending or leaving
Bounce, spin and zoom
5
Bouncing, springing, spinning and zooming
Impacts
12
Objects landing, colliding and hitting
Cinematic hits
8
Big, bass-heavy hits for title slams, trailer beats and dramatic moments
Booms and bass
15
Deep booms and drops, sub drops, 808s and rumbles, for big reveals, title cards and scene changes
Pops
9
Elements appearing, scaling and popping into view
Appear and vanish
6
Things vanishing in a puff, dissolving, materialising, flipping over and bursting
Shape changes
8
Shapes and icons growing, shrinking, morphing, stretching, inflating and filling
Interface
36
Clicks, taps, toggles and switches, snaps, clips, claps and hovers
Keyboard
15
Typing, keypresses and text entry
Writing and drawing
16
Pens, pencils, markers and chalk: writing, drawing and ticking, erasing and highlighting
Phone
5
Phones vibrating, ringing and dialling
Mechanical
12
Locks, latches, switches, shutters and other mechanisms
Foley
12
Paper, card, fabric, plastic, glass and wood being handled
Everyday props
10
Pages, books, scissors, tape, boxes, zips, bells and water
Notifications
28
Success and failure, errors, alerts, messages and attention
Tones
8
Single notes: a ding for highlights, and steps that rise as text or list items appear
Shimmers
9
Highlights, sparkles and positive reveals
Logo stings
14
Short signatures for logo reveals, end cards and title moments
Celebration
5
Applause, party poppers, confetti and fireworks
Time
6
Clocks ticking, alarms, timers, stopwatches and countdowns
Money
7
Coins, a till, paying by card and counting
[truncated]
For AI tools and agents. Read GUIDE.md , which maps on-screen moments to sounds and explains placing
and fitting them, and catalogue.md , with one line per sound. For code, load
sounds.json . Then:
Find a sound by what happens on screen: its use_for , avoid_for and tags say where it fits, and its
fits_seconds what length of moment.
Place it so its sync point meets the moment: start it at that moment minus its sync.at . That one rule works
for every sound.
Levels: the sounds are matched in loudness as they are, and each says how far it comes down under speech
( under_voiceover_db ).
For people. Browse sounds/ by category, or search catalogue.md for a word.
sounds.json : how to read and use the data, in brief, its categories, and a record for every sound: name,
description, uses and things to avoid, tags, length, sync point and the moment it meets, timing, the length of
moment it fits, level, how far it can be adapted, tuning, and what it layers well with.
sounds.schema.json describes it.
catalogue.md : every sound in one line each, to read when choosing.
analysis.json : each sound's detailed measurements (loudness, spectrum, stereo image).
Format
WAV, 24-bit, 48 kHz; mono, or stereo that's centred and never panned
Loudness
Matched across the kit on one scale (the loudest 150 ms, K-weighted, in LUFS), with each category's intensities apart from each other
Peaks
True peak at or under −3 dBTP, no clipping
Ends
Trimmed close, starting and ending on silence, without clicks
Low end
Cut to what each sound needs, so nothing rumbles under a voice except the bass sounds, which are meant to
How they're made
Every sound is synthesised by code: struck materials, air, friction, bubbles and tuned instruments, modelled on how
real things make sound, and clean electronic tones for the sounds that are digital by nature. No recordings, samples
or other sound libraries went into any of them. The sounds, and the code that makes them, were created with Claude,
Anthropic's AI model, directed by Wyzowl, and every one was listened to and approved before release.
CC0 1.0 Universal : no rights reserved. Use the sounds and the data for anything, commercial or not, change
them, and share them, with no need to ask or credit. If you'd like to give credit, though, it's welcome: "SFX Kit by
Wyzowl".
320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0.
Readme CC0-1.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0. - wyzowl/sfx-kit

GitHub - wyzowl/sfx-kit: 320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0. · GitHub
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
wyzowl
sfx-kit Public Notifications You must be signed in to change notification settings
Star 2 ( 2 ) You must be signed in to star a repository
320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0.
Readme CC0-1.0 license Activity Custom properties Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
2 Commits 2 Commits Folders and files
sounds sounds GUIDE.md GUIDE.md LICENSE LICENSE README.md README.md analysis.json analysis.json catalogue.md catalogue.md sounds.json sounds.json sounds.schema.json sounds.schema.json View all files Repository files navigation
Also on Hugging Face: huggingface.co/datasets/wyzowl/sfx-kit .
320 original sound effects for videos, animations and interfaces, made to be as easy for AI tools to use as for
people. Every sound comes with the data needed to choose it, place it in time and mix it: what it is, when to use it
and when not to, its length, its sync point and its level.
Free for any use, commercial or not: CC0 , no credit required.
Category
Sounds
What it's for
Whooshes
9
Movement, sliding, flying and transitions
Swishes and zips
9
Quick, sharp movement: flicks, cuts, and things zipping in or out
Swipes
6
Fingers swiping across screens: app demos, cards and carousels
Reverse whooshes
6
Air drawn in to a sudden stop: things collapsing into a point, and the build into a cut
Risers and fallers
12
Building anticipation before a reveal, and things descending or leaving
Bounce, spin and zoom
5
Bouncing, springing, spinning and zooming
Impacts
12
Objects landing, colliding and hitting
Cinematic hits
8
Big, bass-heavy hits for title slams, trailer beats and dramatic moments
Booms and bass
15
Deep booms and drops, sub drops, 808s and rumbles, for big reveals, title cards and scene changes
Pops
9
Elements appearing, scaling and popping into view
Appear and vanish
6
Things vanishing in a puff, dissolving, materialising, flipping over and bursting
Shape changes
8
Shapes and icons growing, shrinking, morphing, stretching, inflating and filling
Interface
36
Clicks, taps, toggles and switches, snaps, clips, claps and hovers
Keyboard
15
Typing, keypresses and text entry
Writing and drawing
16
Pens, pencils, markers and chalk: writing, drawing and ticking, erasing and highlighting
Phone
5
Phones vibrating, ringing and dialling
Mechanical
12
Locks, latches, switches, shutters and other mechanisms
Foley
12
Paper, card, fabric, plastic, glass and wood being handled
Everyday props
10
Pages, books, scissors, tape, boxes, zips, bells and water
Notifications
28
Success and failure, errors, alerts, messages and attention
Tones
8
Single notes: a ding for highlights, and steps that rise as text or list items appear
Shimmers
9
Highlights, sparkles and positive reveals
Logo stings
14
Short signatures for logo reveals, end cards and title moments
Celebration
5
Applause, party poppers, confetti and fireworks
Time
6
Clocks ticking, alarms, timers, stopwatches and countdowns
Money
7
Coins, a till, paying by card and counting
[truncated]
For AI tools and agents. Read GUIDE.md , which maps on-screen moments to sounds and explains placing
and fitting them, and catalogue.md , with one line per sound. For code, load
sounds.json . Then:
Find a sound by what happens on screen: its use_for , avoid_for and tags say where it fits, and its
fits_seconds what length of moment.
Place it so its sync point meets the moment: start it at that moment minus its sync.at . That one rule works
for every sound.
Levels: the sounds are matched in loudness as they are, and each says how far it comes down under speech
( under_voiceover_db ).
For people. Browse sounds/ by category, or search catalogue.md for a word.
sounds.json : how to read and use the data, in brief, its categories, and a record for every sound: name,
description, uses and things to avoid, tags, length, sync point and the moment it meets, timing, the length of
moment it fits, level, how far it can be adapted, tuning, and what it layers well with.
sounds.schema.json describes it.
catalogue.md : every sound in one line each, to read when choosing.
analysis.json : each sound's detailed measurements (loudness, spectrum, stereo image).
Format
WAV, 24-bit, 48 kHz; mono, or stereo that's centred and never panned
Loudness
Matched across the kit on one scale (the loudest 150 ms, K-weighted, in LUFS), with each category's intensities apart from each other
Peaks
True peak at or under −3 dBTP, no clipping
Ends
Trimmed close, starting and ending on silence, without clicks
Low end
Cut to what each sound needs, so nothing rumbles under a voice except the bass sounds, which are meant to
How they're made
Every sound is synthesised by code: struck materials, air, friction, bubbles and tuned instruments, modelled on how
real things make sound, and clean electronic tones for the sounds that are digital by nature. No recordings, samples
or other sound libraries went into any of them. The sounds, and the code that makes them, were created with Claude,
Anthropic's AI model, directed by Wyzowl, and every one was listened to and approved before release.
CC0 1.0 Universal : no rights reserved. Use the sounds and the data for anything, commercial or not, change
them, and share them, with no need to ask or credit. If you'd like to give credit, though, it's welcome: "SFX Kit by
Wyzowl".
320 original sound effects for videos, animations and interfaces, with the data AI tools need to choose, place and mix each one. CC0.
Readme CC0-1.0 license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
