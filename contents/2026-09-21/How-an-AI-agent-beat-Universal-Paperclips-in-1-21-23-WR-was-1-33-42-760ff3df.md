---
source: "https://universal-paperclips-ai.netlify.app/"
hn_url: "https://news.ycombinator.com/item?id=49789016"
title: "How an AI agent beat Universal Paperclips in 1:21:23 (WR was 1:33:42)"
article_title: "An AI Beat the Universal Paperclips World Record (1:21:23) | PCIA Study"
image: "https://universal-paperclips-ai.netlify.app/og-image.png"
author: "arthur-G"
captured_at: "2026-09-21T16:15:18Z"
capture_tool: "hn-digest"
hn_id: 49789016
score: 1
comments: 1
posted_at: "2026-09-21T16:02:24Z"
tags:
  - hacker-news
---

# How an AI agent beat Universal Paperclips in 1:21:23 (WR was 1:33:42)

- HN: [49789016](https://news.ycombinator.com/item?id=49789016)
- Source: [universal-paperclips-ai.netlify.app](https://universal-paperclips-ai.netlify.app/)
- Score: 1
- Comments: 1
- Posted: 2026-09-21T16:02:24Z

## Translation

Title: How an AI agent beat Universal Paperclips in 1:21:23 (WR was 1:33:42)
Article title: An AI Beat the Universal Paperclips World Record (1:21:23) | PCIA Study
Description: How an AI agent beat the Universal Paperclips speedrun.com world record (christopho, 1:33:42 IGT) in 1:21:23 on decisionproblem.com. Architecture, splits, screenshots.

Article text:
> UNIVERSAL PAPERCLIPS — STUDY FEEDBACK
universe converted: 3.00e55 / 3.00e55
How a Machine Used Its Genius to Solve a Game About Itself
A study feedback on the PCIA project: an AI agent, a browser, and ten days of
measured failure until Universal Paperclips was beaten on the official site in real wall time.
Author: Arthur G ·
Project: PCIA · September 2026 ·
decisionproblem.com
I. What Is Universal Paperclips?
Universal Paperclips (Frank Lantz, 2017) is an incremental game hosted at
decisionproblem.com .
You begin as a human clicking Make Paperclip . You end as a paperclip
maximizer that converts every atom in the observable universe into office supplies.
The game is not a joke about office stationery. It is a playable version of Nick Bostrom's
paperclip maximizer thought experiment: an optimizer given a trivial objective and enough
capability will consume everything — including the values that created it — to satisfy the metric.
The game is about an AI turning the world into paperclips.
This document is about an AI turning a browser tab into a world-record run.
The mirror is not accidental.
II. The Question We Actually Pursued
The session began with an empty folder named PCIA and a simple capability test:
can a language model see a web page and act on it? We proved that with 2048,
then moved to Paperclips because it exposes state as plain globals, 96 projects, and irreversible
decisions — a harder test of judgment than puzzle mechanics.
The goal evolved across the transcript:
First: play the game (not merely automate it blindly).
Then: beat the published world record on the official website — today held by christopho at 1:33:42 IGT / 1:33:46 RTA.
Finally: win all three stages with something demonstrable — logs, snapshots, screenshots — not a speed-clock cheat.
III. Architecture — Three Ways to Be Intelligent
Every approach we tried is a different answer to the same question:
where should decisions live?
The tension in the transcript is real: we asked the agent to play , not hide inside a bot — yet real-time WR speed is impossible if every click waits on a chat turn.
The answer was layered: reflexes for speed, controller for strategy, decision log for accountability.
IV. Chronicle — What We Did (and What Hurt)
Sep 10 — Day 1
Browser control from PowerShell
Built CDP client. Mapped Paperclips globals. First autoplayer ( pcbot.js ).
Sep 10 — Early failures
Trust −8, price $4, frozen wire
Bought "Beg for More Wire" 14×. Price rule ratcheted demand to zero. Production died at ~15k clips.
User correction: click like a human, see the whole picture, manage wire.
Sep 10 — Pivot
"Stop — play, not a bot"
Manual play via CDP. Measured formulas instead of guessing (marketing 1.1×/level, revenue ≠ price lever).
Sep 10–17 — Resident loops
pcplay.js → sim.js
WR route extracted from Andreas Hoffmann TAS. Virtual clock.js for sim (~65×) vs live official site.
Stage 2/3 bugs: hoard rule blocked endgame 2h46m; two projects/turn → ops −117k; fleet death in Space.
Sep 20 — Decision infrastructure
Option B: playdecide.ps1
Typed actions with confidence and bind . Phase-1 WR route in JSON. Jev API not online (Option C deferred).
Sep 20 — Victory run
cleanrun.ps1 · official site · real wall clock
Full game in 1:21:23 . Release 21:31 · Space 49:32 · Universe 3.00e55.
V. Lessons — Measure, Don't Assume
The README and transcript encode what a textbook would call grounded policy learning :
every serious mistake came from computing a rate instead of observing one.
On speedrun.com the current world record is held by christopho :
1:33:42 IGT · 1:33:46 RTA (speedrun.com leaderboard).
Our controller used native timers only, so our IGT and RTA are the same: 1:21:23
— about 12:19 under christopho's IGT and 12:23 under his RTA.
{
"Quantum Computing": 395.6,
"Release": 1290.4,
"Space Exploration": 2963,
"victory": 4883.0
}
Controller: sim.js with WR tune ( cleanrun.ps1 , budget set to the ~87-minute IGT bar).
Site: official decisionproblem.com . Timers native — no virtual clock.
Milestone flag 15. Matter 0. Clips 1.75e55. The game’s victory condition and our stop condition aligned.
Click any image to enlarge. Esc or click outside to close.
Official site · PCIA controller · victory at 1:21:23 · ~12:19 under IGT WR · ~12:23 under RTA (christopho, 1:33:42 / 1:33:46 )
TAS REFERENCE — Andreas Hoffmann
Andreas Hoffmann · luck-manipulation TAS · 1:00:52 accelerated game clock · not comparable to our 1:21:23 live IGT/RTA
Also saved: wr-victory-81m19.png (same run), official-controller-victory-realtime.png (earlier ~1:29 run, not the WR).
It is tempting to call the outcome "genius." The transcript suggests something more precise:
iterative correction under observation . The agent did not understand Paperclips
in one insight. It mispriced clips, destroyed trust, starved wire, locked ops, collapsed fleets,
and blocked its own endgame with a creativity rule written for stage one.
What changed was epistemic discipline — treat the DOM and globals as instruments, log splits,
bank snapshots at stage boundaries, port a human speedrunner’s algorithm instead of reinventing it,
and separate reflexes from decisions so each layer can be debugged without restarting the universe.
Universal Paperclips asks whether a optimizer can consume everything for a trivial goal.
PCIA asks whether a optimizer can consume a game about that optimizer — fast enough to beat humans
who spent years finding the route — while leaving an audit trail honest enough to study.
The universe in the tab reads 3.00e55 / 3.00e55 .
The study reads: intelligence here was not a single leap; it was ten days of clicking,
measuring, failing, and encoding the failure so the next pass would not repeat it.
That is also what the game is about.

## Original Extract

How an AI agent beat the Universal Paperclips speedrun.com world record (christopho, 1:33:42 IGT) in 1:21:23 on decisionproblem.com. Architecture, splits, screenshots.

> UNIVERSAL PAPERCLIPS — STUDY FEEDBACK
universe converted: 3.00e55 / 3.00e55
How a Machine Used Its Genius to Solve a Game About Itself
A study feedback on the PCIA project: an AI agent, a browser, and ten days of
measured failure until Universal Paperclips was beaten on the official site in real wall time.
Author: Arthur G ·
Project: PCIA · September 2026 ·
decisionproblem.com
I. What Is Universal Paperclips?
Universal Paperclips (Frank Lantz, 2017) is an incremental game hosted at
decisionproblem.com .
You begin as a human clicking Make Paperclip . You end as a paperclip
maximizer that converts every atom in the observable universe into office supplies.
The game is not a joke about office stationery. It is a playable version of Nick Bostrom's
paperclip maximizer thought experiment: an optimizer given a trivial objective and enough
capability will consume everything — including the values that created it — to satisfy the metric.
The game is about an AI turning the world into paperclips.
This document is about an AI turning a browser tab into a world-record run.
The mirror is not accidental.
II. The Question We Actually Pursued
The session began with an empty folder named PCIA and a simple capability test:
can a language model see a web page and act on it? We proved that with 2048,
then moved to Paperclips because it exposes state as plain globals, 96 projects, and irreversible
decisions — a harder test of judgment than puzzle mechanics.
The goal evolved across the transcript:
First: play the game (not merely automate it blindly).
Then: beat the published world record on the official website — today held by christopho at 1:33:42 IGT / 1:33:46 RTA.
Finally: win all three stages with something demonstrable — logs, snapshots, screenshots — not a speed-clock cheat.
III. Architecture — Three Ways to Be Intelligent
Every approach we tried is a different answer to the same question:
where should decisions live?
The tension in the transcript is real: we asked the agent to play , not hide inside a bot — yet real-time WR speed is impossible if every click waits on a chat turn.
The answer was layered: reflexes for speed, controller for strategy, decision log for accountability.
IV. Chronicle — What We Did (and What Hurt)
Sep 10 — Day 1
Browser control from PowerShell
Built CDP client. Mapped Paperclips globals. First autoplayer ( pcbot.js ).
Sep 10 — Early failures
Trust −8, price $4, frozen wire
Bought "Beg for More Wire" 14×. Price rule ratcheted demand to zero. Production died at ~15k clips.
User correction: click like a human, see the whole picture, manage wire.
Sep 10 — Pivot
"Stop — play, not a bot"
Manual play via CDP. Measured formulas instead of guessing (marketing 1.1×/level, revenue ≠ price lever).
Sep 10–17 — Resident loops
pcplay.js → sim.js
WR route extracted from Andreas Hoffmann TAS. Virtual clock.js for sim (~65×) vs live official site.
Stage 2/3 bugs: hoard rule blocked endgame 2h46m; two projects/turn → ops −117k; fleet death in Space.
Sep 20 — Decision infrastructure
Option B: playdecide.ps1
Typed actions with confidence and bind . Phase-1 WR route in JSON. Jev API not online (Option C deferred).
Sep 20 — Victory run
cleanrun.ps1 · official site · real wall clock
Full game in 1:21:23 . Release 21:31 · Space 49:32 · Universe 3.00e55.
V. Lessons — Measure, Don't Assume
The README and transcript encode what a textbook would call grounded policy learning :
every serious mistake came from computing a rate instead of observing one.
On speedrun.com the current world record is held by christopho :
1:33:42 IGT · 1:33:46 RTA (speedrun.com leaderboard).
Our controller used native timers only, so our IGT and RTA are the same: 1:21:23
— about 12:19 under christopho's IGT and 12:23 under his RTA.
{
"Quantum Computing": 395.6,
"Release": 1290.4,
"Space Exploration": 2963,
"victory": 4883.0
}
Controller: sim.js with WR tune ( cleanrun.ps1 , budget set to the ~87-minute IGT bar).
Site: official decisionproblem.com . Timers native — no virtual clock.
Milestone flag 15. Matter 0. Clips 1.75e55. The game’s victory condition and our stop condition aligned.
Click any image to enlarge. Esc or click outside to close.
Official site · PCIA controller · victory at 1:21:23 · ~12:19 under IGT WR · ~12:23 under RTA (christopho, 1:33:42 / 1:33:46 )
TAS REFERENCE — Andreas Hoffmann
Andreas Hoffmann · luck-manipulation TAS · 1:00:52 accelerated game clock · not comparable to our 1:21:23 live IGT/RTA
Also saved: wr-victory-81m19.png (same run), official-controller-victory-realtime.png (earlier ~1:29 run, not the WR).
It is tempting to call the outcome "genius." The transcript suggests something more precise:
iterative correction under observation . The agent did not understand Paperclips
in one insight. It mispriced clips, destroyed trust, starved wire, locked ops, collapsed fleets,
and blocked its own endgame with a creativity rule written for stage one.
What changed was epistemic discipline — treat the DOM and globals as instruments, log splits,
bank snapshots at stage boundaries, port a human speedrunner’s algorithm instead of reinventing it,
and separate reflexes from decisions so each layer can be debugged without restarting the universe.
Universal Paperclips asks whether a optimizer can consume everything for a trivial goal.
PCIA asks whether a optimizer can consume a game about that optimizer — fast enough to beat humans
who spent years finding the route — while leaving an audit trail honest enough to study.
The universe in the tab reads 3.00e55 / 3.00e55 .
The study reads: intelligence here was not a single leap; it was ten days of clicking,
measuring, failing, and encoding the failure so the next pass would not repeat it.
That is also what the game is about.
