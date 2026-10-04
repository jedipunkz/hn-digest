---
source: "https://quickchat.ai/post/claude-code-stockfish-chess-opening-mistakes"
hn_url: "https://news.ycombinator.com/item?id=49958259"
title: "Claude Code Found My Pirc and King's Indian Mistakes"
article_title: "Claude Code Found My Pirc and King's Indian Mistakes | Quickchat AI - AI Agents"
image: "https://quickchat.ai/blog-assets/posts/claude-code-stockfish-chess-opening-mistakes_bg.png"
author: "piotrgrudzien"
captured_at: "2026-10-04T22:48:21Z"
capture_tool: "hn-digest"
hn_id: 49958259
score: 1
comments: 0
posted_at: "2026-10-04T21:50:03Z"
tags:
  - hacker-news
---

# Claude Code Found My Pirc and King's Indian Mistakes

- HN: [49958259](https://news.ycombinator.com/item?id=49958259)
- Source: [quickchat.ai](https://quickchat.ai/post/claude-code-stockfish-chess-opening-mistakes)
- Score: 1
- Comments: 0
- Posted: 2026-10-04T21:50:03Z

## Translation

Title: Claude Code Found My Pirc and King's Indian Mistakes
Article title: Claude Code Found My Pirc and King's Indian Mistakes | Quickchat AI - AI Agents
Description: I ran Claude Code and Stockfish 19 over 2,000 of my chess.com blitz games to find the nine Pirc and King

Article text:
Claude Code Found My Pirc and King's Indian Mistakes | Quickchat AI - AI Agents Copy logo to clipboard (png)
Copy logo to clipboard (svg)
Download brand assets Use Cases
Customer Support Resolve tickets and cut response times
Sales & Lead Generation Qualify leads and recommend products
Ecommerce Recover carts and guide shoppers
IT & Internal Helpdesk Answer employee and internal questions
Enterprise Security, scale, and control for large teams
Platform The AI agent platform Releases What's new in Quickchat AI Contact us Company
Why Us Pricing Resources Log in Start for free Start for free Solutions Use Cases
Customer Support Sales & Lead Generation Ecommerce IT & Internal Helpdesk Enterprise Channels & Integrations
Website chat WordPress WhatsApp Messenger Instagram Discord Telegram Shopify Intercom HubSpot Zendesk More integrations Platform Releases Contact us Why Us Pricing Resources Company
Claude Code Found My Pirc and King's Indian Mistakes
; the original PNG/JPG stays as the
fallback and remains the og:image for social scrapers. --> Intro
Setting long-running tasks for AI Agents to work on computers is the way to work in 2026. I have 15-20 Claude Code / Codex agents running for me every day. I wanted to see if they can do useful work for me in the realm of chess training.
At my chess level, you have to have some kind of an opening repertoire. Even if you spend ~0h per week actually training chess. You do find yourself often falling into trouble in similar kinds of positions. I wanted Claude to go ahead and do two things for me:
Find 5-10 positions I often misplay
Explain to me how to do better in those positions in a way digestible for a 2000-2200 ELO player (not grandmaster- or engine-speak)
I found both Claude’s results and process quite impressive so I’m sharing them here. Key elements that caught my attention:
the scoring system Claude devised for detecting and ranking teachable positions
how it calibrated findings to my playing strength
the actual 9 positions it ended up highlighting
I used Claude Code and Stockfish 19 on a two-core Linux VM to go through 2,000 of my chess.com blitz games, find the Pirc and King’s Indian positions I keep going wrong in, and turn the result into an eleven-page PDF brochure written for a 2000 to 2200 player. The session ran for four hours and thirteen minutes on its own while I was at home with the laptop closed, and I checked on it from my phone.
The nine positions and what I learned from them are in What I learned about my chess . The brochure itself is a PDF you can download .
Why the Pirc and the King’s Indian keep getting me into trouble
As Black I have played the Pirc against 1.e4 and the King’s Indian against 1.d4 for a very long time. Both give good counterattacking chances and both are on thin ice as was once explained to me by GM Mateusz Bartel . White plays simple, natural moves, I make one small inaccuracy, and by move 15 or 20 I am close to lost. It happens far less than it used to, but it still happens, and I had a feeling it happened in the same handful of positions.
Isn’t a chess engine or chess.com enough?
At my level, learning openings from engine recommendations makes no sense. You will learn what the best move is but won’t understand why and how to proceed. It’s almost as bad as writing blog posts using AI.
chess.com’s Game Review is fun but it’s aimed at beginner players and its suggestions are very basic. I wonder if the team at chess.com is working on making Game Review better calibrated to individual player’s strength?
By far the best way to train chess is with a coach or with chess books. But that takes a lot of time. So let’s use AI instead.
Position 6 of the brochure , rendered from its data: the Austrian Attack after 6.Be3 Ng4 7.Bg1. I played 7…e5 in all three games that reached it. The engine wants 7…c5.
What I wanted from 2,000 games
I wanted five to ten positions: the specific moments where my habits and the engine disagree, with the move I should play and the reason in words a 2200 player uses. No ocean of engine lines, and no advice to develop my pieces and castle early. A recurring-position audit is a run that groups every game by the exact position where my results start to slip, so the positions are ranked by how often I reach them and how much they cost, not by how badly one game went.
Doing that by hand means loading hundreds of games into an analysis board, noting the move where the evaluation drops, and keeping a tally. It is weeks of evenings, which is why nobody does it.
Why run this as an unattended Claude Code session on a VM?
These days I only ever work with ten to twenty AI sessions running at the same time, so doing this project the same way was the natural choice. Each session gets its own virtual machine, which changes three things. It can run for hours without my laptop being open. It can install whatever it needs, in this case a PDF typesetter and two fonts. And it can be reached from a phone , because the terminal is a web page.
The machine for this run had two CPU cores, 4 GB of memory, Stockfish 19 and the python-chess library already installed, and a Claude Code session that starts with permission prompts switched off. The VM does not idle or sleep, which matters when the engine is going to run for three and a half hours. It also does not notify you when the work is done; if you want a ping, you ask for one, and I did.
The first message asks for a plan, not for work. It states who I am, what I want, what the machine has, how long it may run, and what the deliverable is. It also contains the one rule that let me publish this post: my username and my opponents’ names stay out of everything it prints.
I'm a 2200+ blitz player on chess.com. As Black I play the Pirc and the King's Indian. They give good counterattacking chances but they're risky: White plays simple natural moves, I make one small inaccuracy, and by move 15-20 I'm close to lost. It happens less than it used to, but it still happens, and I want to know exactly where.
I want you to go through all my games and find the 5-10 positions where this keeps happening, then teach me those positions.
What you have here: Stockfish 19, python-chess in .venv, 2 CPU cores and 4 GB of RAM. My chess.com username is in /home/dev/chess/.chesscom_username. Read it from there and never print it. I'm going to publish screenshots of this session, so keep my username, my opponents' names and anything about the hosting of this machine out of logs, filenames, printed output and the report. Call my opponents "White".
I think I have about 1,900 blitz games. Keep the ones where I'm Black and the opening is the Pirc, the King's Indian or the Modern. Classify by the actual move sequence, not only by the ECO code chess.com attaches.
What I want at the end is a PDF brochure in results/. Page one is a one-pager: the positions at a glance. Then one page per position with a diagram from my side of the board, the line that gets there, what I usually play there and how often (from my own games), what I should play instead, and why, in the words a 2000-2200 FIDE player would use. No beginner advice. No pages of engine lines. If a move is only good for reasons Stockfish finds at depth 30, I can't learn it; prefer plan-level explanations and say when the engine's reason is not a human reason. Make it look good: this is something I want to print and keep.
You can run for several hours and I won't be watching, so budget the engine time for two cores, save intermediate results so nothing is lost if something crashes, and keep a short progress file in reports/ that I can read from my phone.
I did a planning session earlier and saved its notes i
[truncated]
“Say when the engine’s reason is not a human reason” is what separates a brochure for me from a brochure for a computer. The planning note the prompt mentions, pipeline/DESIGN.md , is a file from an earlier session: about 5,000 words of measurements and arithmetic, with the engine speed on this machine, the node budgets, the scoring ideas and the privacy rules. It opens like this:
Design note: finding my recurring Pirc / King’s Indian trouble spots. Written in a planning session on 2026-10-03. It records what was measured on this machine and the approach that came out of that planning. It is a starting point, not a spec. Check the numbers, change what does not hold up, and present your own plan before running anything.
Its decision table, which the session kept almost unchanged apart from the budgets:
The full note is published next to the brochure. Handing the next session a file is cheaper than typing the same context twice, and it is also what let the session’s plan come back in five minutes.
The plan came back five and a half minutes later , in seven numbered sections and about 1,200 words. It covered the download and the classification, an engine budget in nodes rather than seconds, the selection rules, how it would keep the advice at my level, how it would check its own work, what it had changed from my notes, and a timeline. I was supposed to read it on the laptop before leaving, and I did.
The end of the plan, captured in the desktop browser at 19:10: the timeline table, the deliverables and the assumptions it wanted confirmed.
The plan’s estimates and what actually happened:
I sent two small adjustments and said go. The reply confirmed both and started installing the typesetter in the background while it wrote the first scripts.
How Claude Code downloaded and classified 2,000 chess.com games
The chess.com public API serves a player’s games as monthly archives, 29 of them in my case, with the full PGN of every game. The download took 41 seconds : 2,000 games, 1,904 of them 3+2 blitz, 955 of those with me as Black.
Download and classification output, rendered from the session transcript.
Classification was by move sequence, not by the opening name chess.com attaches. A game counts as mine if Black plays both …d6 and …g6 in the first ten moves and does not push …d5 or an early …c5; it is a King’s Indian if White has both c4 and d4 by move 10, and a Pirc if White has e4 without c4. That gave 791 games: 532 Pirc and 259 King’s Indian . The classifier was cross-checked against chess.com’s labels afterwards and kept 99 percent of the games chess.com calls Pirc and 92 percent of the ones it calls King’s Indian. My Modern Defense bucket stayed empty, because I play …Nf6 early in every game.
Two facts came out of this stage before any engine ran. I score 54 percent in the Pirc and 49 percent in the King’s Indian. And a bigger gap: after move 12 my expected score is 40 percent or less in 43 percent of my King’s Indian games, against 28 percent of my Pirc games . The thin ice was mostly on the 1.d4 side. I would have guessed the opposite.
How Stockfish screened 791 games on two cores
The engine work was three passes with fixed node budgets, so that a throttled core would stretch the clock without changing a result. Before launching, the session ran a smoke test, measured about 170,000 nodes per second per core, less than my notes had assumed, and cut the budgets to fit: 100,000 nodes per position for screening and 600,000 for confirmation .
Concept diagram of the pipeline. Every box is a script the session wrote during the run; the cache in the middle is what makes it resumable.
The unit of work was a position, not a game. Opening positions repeat heavily within one player’s repertoire, so every position was stored once under its FEN in a small SQLite cache, keyed by the budget it was analysed at. Everything downstream (the per-game tables, the error list, the clusters, the brochure) is computed from that cache plus the raw games. Any stage can be re-run at any time, and a crash costs one position.
The screening pass covered moves 4 to 20 of every game in two phases: 9,155 unique positions out of 14,203 raw ones in the first, 11,563 in the second, with games already decided by move 12 skipped in the second. 2

[truncated]

## Original Extract

I ran Claude Code and Stockfish 19 over 2,000 of my chess.com blitz games to find the nine Pirc and King

Claude Code Found My Pirc and King's Indian Mistakes | Quickchat AI - AI Agents Copy logo to clipboard (png)
Copy logo to clipboard (svg)
Download brand assets Use Cases
Customer Support Resolve tickets and cut response times
Sales & Lead Generation Qualify leads and recommend products
Ecommerce Recover carts and guide shoppers
IT & Internal Helpdesk Answer employee and internal questions
Enterprise Security, scale, and control for large teams
Platform The AI agent platform Releases What's new in Quickchat AI Contact us Company
Why Us Pricing Resources Log in Start for free Start for free Solutions Use Cases
Customer Support Sales & Lead Generation Ecommerce IT & Internal Helpdesk Enterprise Channels & Integrations
Website chat WordPress WhatsApp Messenger Instagram Discord Telegram Shopify Intercom HubSpot Zendesk More integrations Platform Releases Contact us Why Us Pricing Resources Company
Claude Code Found My Pirc and King's Indian Mistakes
; the original PNG/JPG stays as the
fallback and remains the og:image for social scrapers. --> Intro
Setting long-running tasks for AI Agents to work on computers is the way to work in 2026. I have 15-20 Claude Code / Codex agents running for me every day. I wanted to see if they can do useful work for me in the realm of chess training.
At my chess level, you have to have some kind of an opening repertoire. Even if you spend ~0h per week actually training chess. You do find yourself often falling into trouble in similar kinds of positions. I wanted Claude to go ahead and do two things for me:
Find 5-10 positions I often misplay
Explain to me how to do better in those positions in a way digestible for a 2000-2200 ELO player (not grandmaster- or engine-speak)
I found both Claude’s results and process quite impressive so I’m sharing them here. Key elements that caught my attention:
the scoring system Claude devised for detecting and ranking teachable positions
how it calibrated findings to my playing strength
the actual 9 positions it ended up highlighting
I used Claude Code and Stockfish 19 on a two-core Linux VM to go through 2,000 of my chess.com blitz games, find the Pirc and King’s Indian positions I keep going wrong in, and turn the result into an eleven-page PDF brochure written for a 2000 to 2200 player. The session ran for four hours and thirteen minutes on its own while I was at home with the laptop closed, and I checked on it from my phone.
The nine positions and what I learned from them are in What I learned about my chess . The brochure itself is a PDF you can download .
Why the Pirc and the King’s Indian keep getting me into trouble
As Black I have played the Pirc against 1.e4 and the King’s Indian against 1.d4 for a very long time. Both give good counterattacking chances and both are on thin ice as was once explained to me by GM Mateusz Bartel . White plays simple, natural moves, I make one small inaccuracy, and by move 15 or 20 I am close to lost. It happens far less than it used to, but it still happens, and I had a feeling it happened in the same handful of positions.
Isn’t a chess engine or chess.com enough?
At my level, learning openings from engine recommendations makes no sense. You will learn what the best move is but won’t understand why and how to proceed. It’s almost as bad as writing blog posts using AI.
chess.com’s Game Review is fun but it’s aimed at beginner players and its suggestions are very basic. I wonder if the team at chess.com is working on making Game Review better calibrated to individual player’s strength?
By far the best way to train chess is with a coach or with chess books. But that takes a lot of time. So let’s use AI instead.
Position 6 of the brochure , rendered from its data: the Austrian Attack after 6.Be3 Ng4 7.Bg1. I played 7…e5 in all three games that reached it. The engine wants 7…c5.
What I wanted from 2,000 games
I wanted five to ten positions: the specific moments where my habits and the engine disagree, with the move I should play and the reason in words a 2200 player uses. No ocean of engine lines, and no advice to develop my pieces and castle early. A recurring-position audit is a run that groups every game by the exact position where my results start to slip, so the positions are ranked by how often I reach them and how much they cost, not by how badly one game went.
Doing that by hand means loading hundreds of games into an analysis board, noting the move where the evaluation drops, and keeping a tally. It is weeks of evenings, which is why nobody does it.
Why run this as an unattended Claude Code session on a VM?
These days I only ever work with ten to twenty AI sessions running at the same time, so doing this project the same way was the natural choice. Each session gets its own virtual machine, which changes three things. It can run for hours without my laptop being open. It can install whatever it needs, in this case a PDF typesetter and two fonts. And it can be reached from a phone , because the terminal is a web page.
The machine for this run had two CPU cores, 4 GB of memory, Stockfish 19 and the python-chess library already installed, and a Claude Code session that starts with permission prompts switched off. The VM does not idle or sleep, which matters when the engine is going to run for three and a half hours. It also does not notify you when the work is done; if you want a ping, you ask for one, and I did.
The first message asks for a plan, not for work. It states who I am, what I want, what the machine has, how long it may run, and what the deliverable is. It also contains the one rule that let me publish this post: my username and my opponents’ names stay out of everything it prints.
I'm a 2200+ blitz player on chess.com. As Black I play the Pirc and the King's Indian. They give good counterattacking chances but they're risky: White plays simple natural moves, I make one small inaccuracy, and by move 15-20 I'm close to lost. It happens less than it used to, but it still happens, and I want to know exactly where.
I want you to go through all my games and find the 5-10 positions where this keeps happening, then teach me those positions.
What you have here: Stockfish 19, python-chess in .venv, 2 CPU cores and 4 GB of RAM. My chess.com username is in /home/dev/chess/.chesscom_username. Read it from there and never print it. I'm going to publish screenshots of this session, so keep my username, my opponents' names and anything about the hosting of this machine out of logs, filenames, printed output and the report. Call my opponents "White".
I think I have about 1,900 blitz games. Keep the ones where I'm Black and the opening is the Pirc, the King's Indian or the Modern. Classify by the actual move sequence, not only by the ECO code chess.com attaches.
What I want at the end is a PDF brochure in results/. Page one is a one-pager: the positions at a glance. Then one page per position with a diagram from my side of the board, the line that gets there, what I usually play there and how often (from my own games), what I should play instead, and why, in the words a 2000-2200 FIDE player would use. No beginner advice. No pages of engine lines. If a move is only good for reasons Stockfish finds at depth 30, I can't learn it; prefer plan-level explanations and say when the engine's reason is not a human reason. Make it look good: this is something I want to print and keep.
You can run for several hours and I won't be watching, so budget the engine time for two cores, save intermediate results so nothing is lost if something crashes, and keep a short progress file in reports/ that I can read from my phone.
I did a planning session earlier and saved its notes i
[truncated]
“Say when the engine’s reason is not a human reason” is what separates a brochure for me from a brochure for a computer. The planning note the prompt mentions, pipeline/DESIGN.md , is a file from an earlier session: about 5,000 words of measurements and arithmetic, with the engine speed on this machine, the node budgets, the scoring ideas and the privacy rules. It opens like this:
Design note: finding my recurring Pirc / King’s Indian trouble spots. Written in a planning session on 2026-10-03. It records what was measured on this machine and the approach that came out of that planning. It is a starting point, not a spec. Check the numbers, change what does not hold up, and present your own plan before running anything.
Its decision table, which the session kept almost unchanged apart from the budgets:
The full note is published next to the brochure. Handing the next session a file is cheaper than typing the same context twice, and it is also what let the session’s plan come back in five minutes.
The plan came back five and a half minutes later , in seven numbered sections and about 1,200 words. It covered the download and the classification, an engine budget in nodes rather than seconds, the selection rules, how it would keep the advice at my level, how it would check its own work, what it had changed from my notes, and a timeline. I was supposed to read it on the laptop before leaving, and I did.
The end of the plan, captured in the desktop browser at 19:10: the timeline table, the deliverables and the assumptions it wanted confirmed.
The plan’s estimates and what actually happened:
I sent two small adjustments and said go. The reply confirmed both and started installing the typesetter in the background while it wrote the first scripts.
How Claude Code downloaded and classified 2,000 chess.com games
The chess.com public API serves a player’s games as monthly archives, 29 of them in my case, with the full PGN of every game. The download took 41 seconds : 2,000 games, 1,904 of them 3+2 blitz, 955 of those with me as Black.
Download and classification output, rendered from the session transcript.
Classification was by move sequence, not by the opening name chess.com attaches. A game counts as mine if Black plays both …d6 and …g6 in the first ten moves and does not push …d5 or an early …c5; it is a King’s Indian if White has both c4 and d4 by move 10, and a Pirc if White has e4 without c4. That gave 791 games: 532 Pirc and 259 King’s Indian . The classifier was cross-checked against chess.com’s labels afterwards and kept 99 percent of the games chess.com calls Pirc and 92 percent of the ones it calls King’s Indian. My Modern Defense bucket stayed empty, because I play …Nf6 early in every game.
Two facts came out of this stage before any engine ran. I score 54 percent in the Pirc and 49 percent in the King’s Indian. And a bigger gap: after move 12 my expected score is 40 percent or less in 43 percent of my King’s Indian games, against 28 percent of my Pirc games . The thin ice was mostly on the 1.d4 side. I would have guessed the opposite.
How Stockfish screened 791 games on two cores
The engine work was three passes with fixed node budgets, so that a throttled core would stretch the clock without changing a result. Before launching, the session ran a smoke test, measured about 170,000 nodes per second per core, less than my notes had assumed, and cut the budgets to fit: 100,000 nodes per position for screening and 600,000 for confirmation .
Concept diagram of the pipeline. Every box is a script the session wrote during the run; the cache in the middle is what makes it resumable.
The unit of work was a position, not a game. Opening positions repeat heavily within one player’s repertoire, so every position was stored once under its FEN in a small SQLite cache, keyed by the budget it was analysed at. Everything downstream (the per-game tables, the error list, the clusters, the brochure) is computed from that cache plus the raw games. Any stage can be re-run at any time, and a crash costs one position.
The screening pass covered moves 4 to 20 of every game in two phases: 9,155 unique positions out of 14,203 raw ones in the first, 11,563 in the second, with games already decided by move 12 skipped in the second. 2

[truncated]
