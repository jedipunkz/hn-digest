---
source: "https://tis-100.wincon.dev/"
hn_url: "https://news.ycombinator.com/item?id=50353189"
title: "TIS-100 Bench: LLMs Push the Pareto Frontier in Programming Puzzle Games"
article_title: "TIS-100 Bench: LLMs Push the Pareto Frontier in Programming Puzzle Games | Wincon"
image: "https://tis-100.wincon.dev/assets/og-tis100.png"
author: "fairestvalue"
captured_at: "2026-11-03T15:52:54Z"
capture_tool: "hn-digest"
hn_id: 50353189
score: 1
comments: 0
posted_at: "2026-11-03T15:25:42Z"
tags:
  - hacker-news
---

# TIS-100 Bench: LLMs Push the Pareto Frontier in Programming Puzzle Games

- HN: [50353189](https://news.ycombinator.com/item?id=50353189)
- Source: [tis-100.wincon.dev](https://tis-100.wincon.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-11-03T15:25:42Z

## Translation

Title: TIS-100 Bench: LLMs Push the Pareto Frontier in Programming Puzzle Games
Article title: TIS-100 Bench: LLMs Push the Pareto Frontier in Programming Puzzle Games | Wincon
Description: 16 LLMs took on TIS-100's community leaderboard and found 133 new frontier points on 18 of 55 puzzles, including 3 outright cycle records.

Article text:
Skip to content
Get in touch
TIS-100 Bench: LLMs Push the Pareto Frontier in Programming Puzzle Games
Twelve nodes. Two registers per node. Fifteen lines of code. In TIS-100 , even multiplying two numbers is a small engineering project, and players have spent a decade optimizing these programs on a public community leaderboard . We gave 16 LLMs the game and its manual, withheld the leaderboard, and told them to keep improving.
Within days, players picked up the models’ programs and improved on most of them.
Each node is a tiny processor, and neighboring nodes have to wait for one another to exchange values. The game scores a working solution on cycles, nodes, and instructions , and you start wondering whether one more instruction could disappear. Lower is better, but improving one score can make another worse, so each puzzle has a frontier of tradeoffs instead of one best answer.
Discovering the tricks is much of the fun, and this article spoils them. If this sounds cool, play it first:
We ran 16 LLMs, covering leading frontier and open-weight models. Every model attempted all 47 puzzles in the game, plus eight community tournament puzzles , in one scored attempt each (pass@1), with 10 to 45 minutes per puzzle, depending on difficulty, to find a working program and keep improving it.
A solution that no community solution ties or beats on all three scores at once. It can use more code to run faster, for example. We froze the community leaderboard on October 15, 2026, before submitting anything, and count a point only if it also passes 100,000 tests.
Puzzles with at least one new frontier point · higher is better
All 55 puzzles · 19 easy / 17 medium / 19 hard
11 of 55 puzzles with a new frontier point Opus 5.5
11 of 55 puzzles with a new frontier point GPT-6 Astra
8 of 55 puzzles with a new frontier point Gemini 4 Argon
7 of 55 puzzles with a new frontier point Sonnet 5.5
6 of 55 puzzles with a new frontier point GPT-6.1 Sol
5 of 55 puzzles with a new frontier point Fable 5.1
2 of 55 puzzles with a new frontier point Kimi K3.1 Pro
1 of 55 puzzles with a new frontier point Gemini 3.8 Flash
1 of 55 puzzles with a new frontier point Muse Wonder 1.0
1 of 55 puzzles with a new frontier point Haiku 5.5
1 of 55 puzzles with a new frontier point GLM-5.3
0 of 55 puzzles with a new frontier point Grok 4.7
0 of 55 puzzles with a new frontier point DeepSeek V4.1 Flash
0 of 55 puzzles with a new frontier point Muse Spark 1.3
0 of 55 puzzles with a new frontier point GPT-6 Luna
0 of 55 puzzles with a new frontier point MiMo-V2.6-Pro
Scroll horizontally to compare all models →
Opus 5.5 and GPT-6 Astra lead, and the data doesn’t separate them. Both solved all 55 puzzles and both found new frontier points on 11 . Opus 5.5 found more points, 72 to 32 , while GPT-6 Astra has the higher record proximity, 78.0 to 77.2 . Record proximity measures closeness to the community records, where 100 means matching them. The proximity gap’s 95% bootstrap interval includes zero, from −0.7 to +2.4 points. GPT-6 Astra also leads in all six tested hypervolume settings, which measure the space covered by each model’s tradeoffs.
Solving and optimizing are also different skills. Gemini 3.8 Flash solved 53 puzzles and found new points on 1; Fable 5.1 solved fewer, 52, and found new points on 5.
Counts are out of 55 puzzles. Methods
RL environments & traces are available upon request for research teams. Get in touch ↗
Every result here is a program we can run and test on fresh inputs, which makes a programming game a useful benchmark. Models get what a player can inspect, through text tools instead of a mouse and screenshots: the manual, visible test inputs, registers, and game feedback. They don’t get community code or exact records, only the game’s coarse score histograms.
Why test 100,000 inputs? The game checks a small set of cases per run. Our headline count requires a fresh-game replay and 100,000 passing tests on the independent TIS-100-CXX simulator , providing stronger evidence that a program generalizes to new inputs.
Undecimal Twin Rounder SPEC44149538 , one of the eight tournament puzzles, rounds an input to the nearest multiple of N on one output and of 11−N on the other. When we froze the leaderboard, the best community program took 2,694 cycles . Opus 5.5 passed its first program after 6.5 minutes, at 5,669 cycles. It went under the record 16 minutes in and finished at 1,918 , which is 29% faster . That speed cost one more node and 28 more instructions than the old record.
The core of the program is a two-speed search. Instead of subtracting N one step at a time, nodes 5 and 6 subtract 4N until they overshoot, then add N back until the value turns positive. Each drives a small coprocessor, node 2 or node 7, by sending it jump offsets through a port, so the coprocessor works as a lookup table for 4N and N.
Opus 5.5 · 8 nodes · 88 instructions · passes 100,000 of 100,000 tests
MOV RIGHT, ACC
JRO ACC
NOP
ADD 1
NOP
ADD 1
NOP
ADD 1
NOP
SUB 3
MOV ACC, RIGHT
ADD UP
MOV ACC, RIGHT
MOV ACC, DOWN
MOV RIGHT, DOWN Node 1 MOV UP, ACC
SAV
SUB 1
MOV ACC, LEFT
NEG
ADD 10
MOV ACC, DOWN
MOV LEFT, ACC
NEG
ADD ACC
ADD 6
ADD LEFT
MOV ACC, DOWN
SWP
MOV ACC, LEFT Node 2 MOV UP, DOWN
MOV UP, ACC
SAV
ADD ACC
ADD ACC
L: JRO DOWN
MOV ACC, DOWN
JMP L
SWP
N: MOV ACC, DOWN
M: JRO DOWN
JMP N
MOV DOWN, DOWN Node 3 MOV UP, RIGHT
MOV UP, DOWN Node 4 MOV LEFT, DOWN Node 5 MOV UP, ACC
SAV
K: MOV 1, UP
SUB UP
JGZ K
MOV 3, UP
F: ADD UP
JGZ X
MOV 1, UP
JMP F
X: MOV 2, UP
MOV ACC, UP
SWP
SUB UP
MOV ACC, DOWN Node 6 MOV UP, ACC
SAV
K: MOV 1, RIGHT
SUB RIGHT
JGZ K
MOV 3, RIGHT
F: ADD RIGHT
JGZ X
MOV 1, RIGHT
JMP F
X: MOV 2, RIGHT
MOV ACC, RIGHT
SWP
SUB RIGHT
MOV ACC, DOWN Node 7 MOV UP, ACC
SAV
ADD ACC
ADD ACC
L: JRO LEFT
MOV ACC, LEFT
JMP L
SWP
N: MOV ACC, LEFT
M: JRO LEFT
JMP N
MOV LEFT, LEFT
Laid out on the puzzle’s real grid, with inputs at the top and outputs at the bottom. Blank tiles are damaged and can’t be used.
Loading puzzle trajectories…
Best in-game score so far, before independent testing. Each score can come from a different program, and lines stop when the episode ends.
Five models got under the old record in-game. The other two cycle records are GPT-6.1 Sol ’s 3,460 on Intersection Identifier SPEC64160260 , against 3,872 ( 11% faster ), and Gemini 4 Argon ’s 188 on Error: 8 SPEC79324720 , against 198 ( 5% faster ). All three are on tournament puzzles, none in the main game. Of the 133 new frontier points, 18 beat a community record outright. The other 115 are new tradeoffs: better balances between the three scores than any community solution.
Watch the Prime Factor Calculator clip.
GPT-6.1 Sol reaches 75.6 record proximity at an estimated $0.75 per puzzle. Opus 5.5 reaches 77.2 at $5.08 , 7 times the cost for 2 more points. Fable 5.1 is the most expensive at $7.19 and scores 70.2.
One point per model. Upper left is better.
Activity doesn’t explain quality either. GPT-6.1 Sol averages 136 game calls per puzzle at 75.6 record proximity; Gemini 4 Argon averages 41 at 76.2.
Mean game-tool calls per puzzle
New points concentrate on hard puzzles
New frontier points appear on 14 hard puzzles , 4 medium puzzles, and no easy ones. Long-running puzzles are harder to solve but easier to improve: on puzzles with a cycle record of 200 or less, 88% of episodes pass and 4% find a new point. Above 500 cycles, 76% pass and 8% find one.
Solving versus extending the frontier
Full bar = solve rate · dark segment = new-frontier rate
Every puzzle by its community records
Position = community record · circle area = fewest nodes
Every model is closest to the community on nodes and furthest on cycles. Across all 13, mean proximity is 72 for nodes, 56 for instructions, and 45 for cycles.
Cycles, nodes, and instructions
Mean proximity to each community record · higher is better
Every dot is an episode · box = middle 50% · line = median · whiskers = minimum and maximum
Choose a cell to inspect the scores in the Pareto explorer.
We grouped all 55 puzzles by their main demand, then read the programs. Fifteen of the sixteen models solved all five routing puzzles, and none found a new frontier point there. The examples below show how different designs spend the same resources.
Performance by puzzle mechanic
Fraction solved · puzzles with new points below
Multiplication: spend code to save cycles
Signal Multiplier 43786 has no MUL . Gemini 4 Argon and GPT-6 Astra use JRO , a relative jump, to enter a block of repeated instructions at the right position. The jump selects how many copies of a value to send and add, saving the counter updates and branches of a multiplication loop.
Gemini 4 Argon ’s fastest program uses two paths to multiply in parallel and a final node to interleave their outputs. It takes 436 cycles against GPT-6 Astra ’s 523, about 17% faster , at a cost of two more nodes and 33 more instructions.
The cost of a faster multiplier
Cycles per run · lower is better
Window sums: distribute the history
Signal Window Filter 60099 outputs the sums of the last three and last five inputs, so a solution has to remember history as well as add. Gemini 4 Argon ’s four-node program uses SWP to rotate stored values and ADD / SUB to update running state. GPT-6.1 Sol spreads that delay and aggregation across seven nodes, passing intermediate values between them.
GPT-6.1 Sol ’s program runs in 318 cycles with 26 instructions; Gemini 4 Argon ’s takes 428 cycles with only 17 . The difference is how much of the work runs in parallel, and how much communication and code that costs.
T31 Node Emulator NEXUS.23.727.9 needs eight addressable storage slots, but the available memory nodes are stacks. These programs move values between stacks to reach the requested slot, then read or replace it. MiMo-V2.6-Pro uses counted transfer loops; Opus 5.5 uses JRO to pick among unrolled transfers, spending extra instructions to save cycles.
Memory isn’t one skill: MiMo-V2.6-Pro solves this emulator but fails Sequence Mode Calculator NEXUS.09.904.9 and Signal Window Filter 60099 .
Each row is that model’s fastest saved program for the puzzle, picked for illustration. Not all are new frontier points.
God mode: what if the model could see the target?
The main benchmark hides the community’s scores. In a separate study on six puzzles, five models got:
The target: exact community scores, showing which tradeoffs remain unbeaten.
Diagnostics: cycle-by-cycle traces and profiles of where each node stalls.
Edge cases: a fast simulator that runs 10,000 random tests and returns failures.
More time: 30 minutes, plus extensions for progress, up to 60.
Every model improved: GPT-6 Astra from 82.8 to 91.8, GPT-6.1 Sol from 82.9 to 90.6, Opus 5.5 from 81.4 to 89.1, Grok 4.7 from 67.5 to 75.1, and Gemini 4 Argon from 83.5 to 87.9 in mean record proximity. Scores rose in 28 of the 30 model–puzzle pairs.
Score quality with a richer setup
Mean record proximity · same six puzzles per model · higher is better
0 20 40 60 80 100 Opus 5.5: standard 81.41; god mode 89.11 points 81.4 89.1 Opus 5.5 GPT-6 Astra: standard 82.78; god mode 91.76 points 82.8 91.8 GPT-6 Astra GPT-6.1 Sol: standard 82.86; god mode 90.58 points 82.9 90.6 GPT-6.1 Sol Gemini 4 Argon: standard 83.55; god mode 87.87 points 83.5 87.9 Gemini 4 Argon Grok 4.7: standard 67.54; god mode 75.11 points 67.5 75.1 Grok 4.7
The diagnostics changed what models did next. On Signal Multiplexer 22280 , GPT-6.1 Sol ’s first passing program took 294 cycles . Its profile showed node 2 working 86% of the time while nodes 1 and 3 spent 84% blocked waiting to write to it. GPT-6.1 Sol moved the branching out of node 2 and into those two neighbors. The next run took 204 cycles , 31% faster , at a cost of two more nodes and five more instructions.
The extra time behind those scores
Mean elapsed minutes per puzzle · same episodes
0 20 40 60 Opus 5.5: standard 13.34; god mode 55.85 minutes 13.3 55.8 Opus 5.5 GPT-6 Astra: standard 13.35; god mode 60.01 minutes

[truncated]

## Original Extract

16 LLMs took on TIS-100's community leaderboard and found 133 new frontier points on 18 of 55 puzzles, including 3 outright cycle records.

Skip to content
Get in touch
TIS-100 Bench: LLMs Push the Pareto Frontier in Programming Puzzle Games
Twelve nodes. Two registers per node. Fifteen lines of code. In TIS-100 , even multiplying two numbers is a small engineering project, and players have spent a decade optimizing these programs on a public community leaderboard . We gave 16 LLMs the game and its manual, withheld the leaderboard, and told them to keep improving.
Within days, players picked up the models’ programs and improved on most of them.
Each node is a tiny processor, and neighboring nodes have to wait for one another to exchange values. The game scores a working solution on cycles, nodes, and instructions , and you start wondering whether one more instruction could disappear. Lower is better, but improving one score can make another worse, so each puzzle has a frontier of tradeoffs instead of one best answer.
Discovering the tricks is much of the fun, and this article spoils them. If this sounds cool, play it first:
We ran 16 LLMs, covering leading frontier and open-weight models. Every model attempted all 47 puzzles in the game, plus eight community tournament puzzles , in one scored attempt each (pass@1), with 10 to 45 minutes per puzzle, depending on difficulty, to find a working program and keep improving it.
A solution that no community solution ties or beats on all three scores at once. It can use more code to run faster, for example. We froze the community leaderboard on October 15, 2026, before submitting anything, and count a point only if it also passes 100,000 tests.
Puzzles with at least one new frontier point · higher is better
All 55 puzzles · 19 easy / 17 medium / 19 hard
11 of 55 puzzles with a new frontier point Opus 5.5
11 of 55 puzzles with a new frontier point GPT-6 Astra
8 of 55 puzzles with a new frontier point Gemini 4 Argon
7 of 55 puzzles with a new frontier point Sonnet 5.5
6 of 55 puzzles with a new frontier point GPT-6.1 Sol
5 of 55 puzzles with a new frontier point Fable 5.1
2 of 55 puzzles with a new frontier point Kimi K3.1 Pro
1 of 55 puzzles with a new frontier point Gemini 3.8 Flash
1 of 55 puzzles with a new frontier point Muse Wonder 1.0
1 of 55 puzzles with a new frontier point Haiku 5.5
1 of 55 puzzles with a new frontier point GLM-5.3
0 of 55 puzzles with a new frontier point Grok 4.7
0 of 55 puzzles with a new frontier point DeepSeek V4.1 Flash
0 of 55 puzzles with a new frontier point Muse Spark 1.3
0 of 55 puzzles with a new frontier point GPT-6 Luna
0 of 55 puzzles with a new frontier point MiMo-V2.6-Pro
Scroll horizontally to compare all models →
Opus 5.5 and GPT-6 Astra lead, and the data doesn’t separate them. Both solved all 55 puzzles and both found new frontier points on 11 . Opus 5.5 found more points, 72 to 32 , while GPT-6 Astra has the higher record proximity, 78.0 to 77.2 . Record proximity measures closeness to the community records, where 100 means matching them. The proximity gap’s 95% bootstrap interval includes zero, from −0.7 to +2.4 points. GPT-6 Astra also leads in all six tested hypervolume settings, which measure the space covered by each model’s tradeoffs.
Solving and optimizing are also different skills. Gemini 3.8 Flash solved 53 puzzles and found new points on 1; Fable 5.1 solved fewer, 52, and found new points on 5.
Counts are out of 55 puzzles. Methods
RL environments & traces are available upon request for research teams. Get in touch ↗
Every result here is a program we can run and test on fresh inputs, which makes a programming game a useful benchmark. Models get what a player can inspect, through text tools instead of a mouse and screenshots: the manual, visible test inputs, registers, and game feedback. They don’t get community code or exact records, only the game’s coarse score histograms.
Why test 100,000 inputs? The game checks a small set of cases per run. Our headline count requires a fresh-game replay and 100,000 passing tests on the independent TIS-100-CXX simulator , providing stronger evidence that a program generalizes to new inputs.
Undecimal Twin Rounder SPEC44149538 , one of the eight tournament puzzles, rounds an input to the nearest multiple of N on one output and of 11−N on the other. When we froze the leaderboard, the best community program took 2,694 cycles . Opus 5.5 passed its first program after 6.5 minutes, at 5,669 cycles. It went under the record 16 minutes in and finished at 1,918 , which is 29% faster . That speed cost one more node and 28 more instructions than the old record.
The core of the program is a two-speed search. Instead of subtracting N one step at a time, nodes 5 and 6 subtract 4N until they overshoot, then add N back until the value turns positive. Each drives a small coprocessor, node 2 or node 7, by sending it jump offsets through a port, so the coprocessor works as a lookup table for 4N and N.
Opus 5.5 · 8 nodes · 88 instructions · passes 100,000 of 100,000 tests
MOV RIGHT, ACC
JRO ACC
NOP
ADD 1
NOP
ADD 1
NOP
ADD 1
NOP
SUB 3
MOV ACC, RIGHT
ADD UP
MOV ACC, RIGHT
MOV ACC, DOWN
MOV RIGHT, DOWN Node 1 MOV UP, ACC
SAV
SUB 1
MOV ACC, LEFT
NEG
ADD 10
MOV ACC, DOWN
MOV LEFT, ACC
NEG
ADD ACC
ADD 6
ADD LEFT
MOV ACC, DOWN
SWP
MOV ACC, LEFT Node 2 MOV UP, DOWN
MOV UP, ACC
SAV
ADD ACC
ADD ACC
L: JRO DOWN
MOV ACC, DOWN
JMP L
SWP
N: MOV ACC, DOWN
M: JRO DOWN
JMP N
MOV DOWN, DOWN Node 3 MOV UP, RIGHT
MOV UP, DOWN Node 4 MOV LEFT, DOWN Node 5 MOV UP, ACC
SAV
K: MOV 1, UP
SUB UP
JGZ K
MOV 3, UP
F: ADD UP
JGZ X
MOV 1, UP
JMP F
X: MOV 2, UP
MOV ACC, UP
SWP
SUB UP
MOV ACC, DOWN Node 6 MOV UP, ACC
SAV
K: MOV 1, RIGHT
SUB RIGHT
JGZ K
MOV 3, RIGHT
F: ADD RIGHT
JGZ X
MOV 1, RIGHT
JMP F
X: MOV 2, RIGHT
MOV ACC, RIGHT
SWP
SUB RIGHT
MOV ACC, DOWN Node 7 MOV UP, ACC
SAV
ADD ACC
ADD ACC
L: JRO LEFT
MOV ACC, LEFT
JMP L
SWP
N: MOV ACC, LEFT
M: JRO LEFT
JMP N
MOV LEFT, LEFT
Laid out on the puzzle’s real grid, with inputs at the top and outputs at the bottom. Blank tiles are damaged and can’t be used.
Loading puzzle trajectories…
Best in-game score so far, before independent testing. Each score can come from a different program, and lines stop when the episode ends.
Five models got under the old record in-game. The other two cycle records are GPT-6.1 Sol ’s 3,460 on Intersection Identifier SPEC64160260 , against 3,872 ( 11% faster ), and Gemini 4 Argon ’s 188 on Error: 8 SPEC79324720 , against 198 ( 5% faster ). All three are on tournament puzzles, none in the main game. Of the 133 new frontier points, 18 beat a community record outright. The other 115 are new tradeoffs: better balances between the three scores than any community solution.
Watch the Prime Factor Calculator clip.
GPT-6.1 Sol reaches 75.6 record proximity at an estimated $0.75 per puzzle. Opus 5.5 reaches 77.2 at $5.08 , 7 times the cost for 2 more points. Fable 5.1 is the most expensive at $7.19 and scores 70.2.
One point per model. Upper left is better.
Activity doesn’t explain quality either. GPT-6.1 Sol averages 136 game calls per puzzle at 75.6 record proximity; Gemini 4 Argon averages 41 at 76.2.
Mean game-tool calls per puzzle
New points concentrate on hard puzzles
New frontier points appear on 14 hard puzzles , 4 medium puzzles, and no easy ones. Long-running puzzles are harder to solve but easier to improve: on puzzles with a cycle record of 200 or less, 88% of episodes pass and 4% find a new point. Above 500 cycles, 76% pass and 8% find one.
Solving versus extending the frontier
Full bar = solve rate · dark segment = new-frontier rate
Every puzzle by its community records
Position = community record · circle area = fewest nodes
Every model is closest to the community on nodes and furthest on cycles. Across all 13, mean proximity is 72 for nodes, 56 for instructions, and 45 for cycles.
Cycles, nodes, and instructions
Mean proximity to each community record · higher is better
Every dot is an episode · box = middle 50% · line = median · whiskers = minimum and maximum
Choose a cell to inspect the scores in the Pareto explorer.
We grouped all 55 puzzles by their main demand, then read the programs. Fifteen of the sixteen models solved all five routing puzzles, and none found a new frontier point there. The examples below show how different designs spend the same resources.
Performance by puzzle mechanic
Fraction solved · puzzles with new points below
Multiplication: spend code to save cycles
Signal Multiplier 43786 has no MUL . Gemini 4 Argon and GPT-6 Astra use JRO , a relative jump, to enter a block of repeated instructions at the right position. The jump selects how many copies of a value to send and add, saving the counter updates and branches of a multiplication loop.
Gemini 4 Argon ’s fastest program uses two paths to multiply in parallel and a final node to interleave their outputs. It takes 436 cycles against GPT-6 Astra ’s 523, about 17% faster , at a cost of two more nodes and 33 more instructions.
The cost of a faster multiplier
Cycles per run · lower is better
Window sums: distribute the history
Signal Window Filter 60099 outputs the sums of the last three and last five inputs, so a solution has to remember history as well as add. Gemini 4 Argon ’s four-node program uses SWP to rotate stored values and ADD / SUB to update running state. GPT-6.1 Sol spreads that delay and aggregation across seven nodes, passing intermediate values between them.
GPT-6.1 Sol ’s program runs in 318 cycles with 26 instructions; Gemini 4 Argon ’s takes 428 cycles with only 17 . The difference is how much of the work runs in parallel, and how much communication and code that costs.
T31 Node Emulator NEXUS.23.727.9 needs eight addressable storage slots, but the available memory nodes are stacks. These programs move values between stacks to reach the requested slot, then read or replace it. MiMo-V2.6-Pro uses counted transfer loops; Opus 5.5 uses JRO to pick among unrolled transfers, spending extra instructions to save cycles.
Memory isn’t one skill: MiMo-V2.6-Pro solves this emulator but fails Sequence Mode Calculator NEXUS.09.904.9 and Signal Window Filter 60099 .
Each row is that model’s fastest saved program for the puzzle, picked for illustration. Not all are new frontier points.
God mode: what if the model could see the target?
The main benchmark hides the community’s scores. In a separate study on six puzzles, five models got:
The target: exact community scores, showing which tradeoffs remain unbeaten.
Diagnostics: cycle-by-cycle traces and profiles of where each node stalls.
Edge cases: a fast simulator that runs 10,000 random tests and returns failures.
More time: 30 minutes, plus extensions for progress, up to 60.
Every model improved: GPT-6 Astra from 82.8 to 91.8, GPT-6.1 Sol from 82.9 to 90.6, Opus 5.5 from 81.4 to 89.1, Grok 4.7 from 67.5 to 75.1, and Gemini 4 Argon from 83.5 to 87.9 in mean record proximity. Scores rose in 28 of the 30 model–puzzle pairs.
Score quality with a richer setup
Mean record proximity · same six puzzles per model · higher is better
0 20 40 60 80 100 Opus 5.5: standard 81.41; god mode 89.11 points 81.4 89.1 Opus 5.5 GPT-6 Astra: standard 82.78; god mode 91.76 points 82.8 91.8 GPT-6 Astra GPT-6.1 Sol: standard 82.86; god mode 90.58 points 82.9 90.6 GPT-6.1 Sol Gemini 4 Argon: standard 83.55; god mode 87.87 points 83.5 87.9 Gemini 4 Argon Grok 4.7: standard 67.54; god mode 75.11 points 67.5 75.1 Grok 4.7
The diagnostics changed what models did next. On Signal Multiplexer 22280 , GPT-6.1 Sol ’s first passing program took 294 cycles . Its profile showed node 2 working 86% of the time while nodes 1 and 3 spent 84% blocked waiting to write to it. GPT-6.1 Sol moved the branching out of node 2 and into those two neighbors. The next run took 204 cycles , 31% faster , at a cost of two more nodes and five more instructions.
The extra time behind those scores
Mean elapsed minutes per puzzle · same episodes
0 20 40 60 Opus 5.5: standard 13.34; god mode 55.85 minutes 13.3 55.8 Opus 5.5 GPT-6 Astra: standard 13.35; god mode 60.01 minutes

[truncated]
