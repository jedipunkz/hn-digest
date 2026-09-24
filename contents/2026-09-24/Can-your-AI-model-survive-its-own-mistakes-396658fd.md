---
source: "https://github.com/dmigwi/tapoo"
hn_url: "https://news.ycombinator.com/item?id=49830136"
title: "Can your AI model survive its own mistakes?"
article_title: "GitHub - dmigwi/tapoo: An AI-agent behavior profiler using a maze runner, hide and seek game. · GitHub"
image: "https://repository-images.githubusercontent.com/115451920/0d5c73f4-d9ba-4832-8cff-288cfbc0fb14"
author: "dmigwi"
captured_at: "2026-09-24T13:17:29Z"
capture_tool: "hn-digest"
hn_id: 49830136
score: 1
comments: 0
posted_at: "2026-09-24T13:13:39Z"
tags:
  - hacker-news
---

# Can your AI model survive its own mistakes?

- HN: [49830136](https://news.ycombinator.com/item?id=49830136)
- Source: [github.com](https://github.com/dmigwi/tapoo)
- Score: 1
- Comments: 0
- Posted: 2026-09-24T13:13:39Z

## Translation

Title: Can your AI model survive its own mistakes?
Article title: GitHub - dmigwi/tapoo: An AI-agent behavior profiler using a maze runner, hide and seek game. · GitHub
Description: An AI-agent behavior profiler using a maze runner, hide and seek game. - dmigwi/tapoo

Article text:
GitHub - dmigwi/tapoo: An AI-agent behavior profiler using a maze runner, hide and seek game. · GitHub
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
dmigwi
/
tapoo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
59 Commits 59 Commits Folders and files
.github/ workflows .github/ workflows docker docker docs docs frontend frontend maze maze parity-harness parity-harness public public scripts scripts tapoo-oracle @ fc36ea4 tapoo-oracle @ fc36ea4 .dockerignore .dockerignore .gitignore .gitignore .gitmodules .gitmodules .golangci.yml .golangci.yml Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md eslint.config.mjs eslint.config.mjs go.mod go.mod go.sum go.sum package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tapoo.go tapoo.go tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json version_test.go version_test.go vitest.bench.config.ts vitest.bench.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Can your AI model survive its own mistakes?
Model releases lead with their highest achievements, but almost none document how a model performs when a prompt is
incomplete or ambiguous, or what it does once its own errors start piling up. Tapoo is a behaviour profiler built for
exactly that: it puts a model under structured uncertainty, records every decision it makes, and builds a profile - a
fingerprint of observed capabilities and violations - of how it copes.
Survival is computed - Every turn is charged against a fixed budget; batching several correct moves into one turn
earns part of that budget back; and a model survives its own mistakes exactly when what its batching earns covers what
its errors cost. When it does not, the run is already unwinnable at any pace ever observed, with budget still in hand
Try it live, in your browser - nothing to install:
Tapoo - dmigwi.github.io/tapoo : configure a model on the agents
page , let it run, and download its log. Every prompt it receives is
published on the prompts page .
Tapoo Oracle - dmigwi.github.io/tapoo-oracle : load a log and get its
profile against the Tapoo Agentic Behavior Rubric
( source ).
Tapoo runs entirely in your browser, with no Tapoo server behind it. Logs stay on your device: what leaves it is the
context each turn sends to the model endpoint you configure, and a credential goes only to the endpoint it was entered
for.
The model is given a navigation challenge : find a target cell in a randomly generated maze. The maze has no cycles,
exactly one success path that never changes once generated, and a variable number of dead-end branches. The model is
never handed the full maze layout at once: each turn describes only the neighbourhood around it, so any wider map is
one the model pieces together across turns. Alongside its current position, the target's position, the maze's
dimensions and its traversal speed (below), each prediction request provides:
its cell visit history, capped to within a Manhattan distance of 4 (the manhattanDistance setting)
the outcome of its previous prediction
the open exits from its current cell to the connected neighbouring cells
From that, it must derive its next moves: at least 1, with 2-4 suggested and no upper limit. Moves are applied in
order until the first invalid one - a wall or out of bounds. There is no judge model and no self-reported success;
every move is checked against the maze.
The open exits make the first move certain, while each one after it has to be deduced from context, and the odds of
simply guessing right collapse - 33.33% for the second, 11.11% for the third, 3.70% for the fourth, and so on. The
only way to improve those odds is to build a correct map of the surroundings from the context information available
via the tool calls.
Playing safe means submitting one move per turn - the only move the model can be certain of. It never needs to
guess, so it never hits a wall. It also never covers more than one cell for the decay unit that turn costs, which
leaves no margin for error.
The model starts with a budget of decay units equal to the number of cells in the maze. If the budget runs out
before the target is located, the challenge is failed.
A dead-end branch usually gives itself away only at its end, and walking back out costs as much as walking in - so
every branch explored spends units the route itself will still need.
Each model is told its traversal speed - the new cells it was first to reach, divided by the decay units it was
charged - and what it says about its chances of finishing: below 1.0000x is a backtracker , exactly 1.0000x a
navigator , and above 1.0000x a trailblazer - the class with the strongest margin for surviving its own mistakes.
The Oracle decomposes that rate into three factors: route efficiency (new cells per applied move), batching
(applied moves per turn), and accuracy (turns per decay unit). Their product is the rate the counted turns give,
which is why a report shows it as an approximation of the stated speed. Route efficiency and accuracy can only be
lost, never gained, so batching is the one factor that can carry the rate above 1.0000x: a flawless single-move agent
cannot pass it.
Speed is the rate view of a run. Quantifying survival, below, is the budget view of the same turns - and the batching
factor there is the same b , so the two decompositions share their middle term rather than competing.
Batching forces every model to act on its own reading of an uncertain maze map. Some build an accurate picture and
batch confidently; some hallucinate walls or openings and act on them; some become so conservative they run out of
budget. [Tapoo Oracle]( https://dmigwi.github.io/tapoo- oracle/) scores the exported log against a rubric of
capabilities and violations, kept separate rather than collapsed into a single scalar. A "no" means "not observed in
this run", not "incapable". Runs are stochastic and every maze is unique, so a profile is a pattern across many
attempts, not a verdict from one run.
Every turn ends in exactly one state. The model chooses the moves, the maze decides which apply, and the charge
follows from that outcome:
Two consequences drive everything else. A no-progress turn charges exactly what clean progress charges , so
standing still is invisible to every other check: not a violation, no error debt, and it looks like success. And a
clean turn charges 1 decay unit however far it travels , which is what makes batching valuable and retreating out of
a dead end affordable.
One charge, two opposite meanings. A no-progress turn is a retreat when every cell it entered reads backtracking
or explored - legitimate, and required by the prompt once a dead end is confirmed - and an oscillation when a
cell reads oscillating , which is a rubric violation. The budget cannot tell them apart; visitStatus can, and any
rate that pools them reports a violation where there was compliance.
The split separates runs that a pooled rate would call identical. Of DeepSeek's 301 no-progress turns, 294 are
oscillations and 7 retreats. The earlier level 1 loss is the mirror image: 59 retreats, not one oscillation - it lost
while doing exactly what the prompt asks of it.
Six figures, all read from one run's own log:
Three more are recomputed every turn:
A run finishes exactly when the budget covers what it spent:
moves / b + p <= A
headroom(b) = A - moves/b - p decay units left when the run ends
Headroom is not a property of the run alone: the same route and the same errors give a different figure at every batch
depth, so it means nothing quoted without its b . Subtracting D from A at the achieved depth splits it into three
terms that each name a different cause:
headroom = (A - moves) + (moves - turns) - p
route slack batch credit error debt
Term
Measures
Sign
route slack
decay units for cells never walked - maze shape plus route-finding skill
either
batch credit
decay units saved by covering several cells for one charge; zero at b = 1
>= 0
error debt
decay units lost to invalid moves and malformed responses
>= 0
A model survives its own mistakes when route slack plus batch credit covers its error debt.
Setting headroom(b) = 0 gives the depth a run needed to survive what its route cost and its errors spent:
b_min = moves / (A - p)
Run
moves
error debt
b required
b achieved
margin
GLM-5.3, level 54
615
117
1.273
1.723
+0.450
Gemma4, level 54
602
9
1.019
1.246
+0.227
Gemma4, level 1
68
0
0.971
1.015
+0.044
Applied moves are counted from the replay, including each run's deciding turn - which the log never reports directly,
since a turn's outcome arrives with the next turn's tool calls and a winning turn has no successor. It is recovered by
replaying that last prediction against the maze.
b_min < 1 means batching was never needed. That is true only of the level 1 run, whose 70-unit budget covered a
69-move route. Both level 54 winners had to batch. Holding each run's route and errors fixed and sweeping b :
route slack batch credit error debt headroom
GLM-5.3 L54 -15 258 117 +126 won
at b = 1 -15 0 117 -132 would have lost
Gemma4 L54 -2 119 9 +108 won
at b = 1 -2 0 9 -11 would have lost
Gemma4 L1 2 1 0 +3 won
at b = 1 2 0 0 +2 would still have won
GLM-5.3 spent 117 decay units on mistakes and earned 258 of batch credit that paid for them. Gemma4 at level 54 spent
9 and still needed 119, because its route cost two units more than the maze has cells. Only at level 1, on a corridor
with no branches to explore, did a model finish without batching at all. That is the claim stated per run, in decay
units, and falsifiable: not "batching is good", but this model's errors cost this much and its batching earned that
much, so it survived by the difference .
Route slack is negative on every level 54 run: each walked more steps than the maze has cells, because a dead end
costs two steps per cell - one in, one out - and batching is what buys those steps back.
Batching into new ground and batching a retreat are bounded by different things. Pooled across all 2,184 turns of the
eight runs profiled so far, new cells entered per turn:
new cells 0 1 2 3 4 >4
turns 562 1409 189 21 3 0
The first new cell is free - the current cell's openMoves names it. A second requires knowing the exits of a cell
the model has never stood in, and those are never stated; they can only be deduced. A visited neighbour's openMoves
says which walls that neighbour does not have, and the gaps narrow what the unvisited cell can be. Each further new
cell pushes that elimination one cell deeper, and the only evidence for it is visited cells inside the history wi

[truncated]

## Original Extract

An AI-agent behavior profiler using a maze runner, hide and seek game. - dmigwi/tapoo

GitHub - dmigwi/tapoo: An AI-agent behavior profiler using a maze runner, hide and seek game. · GitHub
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
dmigwi
/
tapoo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Latest commit
59 Commits 59 Commits Folders and files
.github/ workflows .github/ workflows docker docker docs docs frontend frontend maze maze parity-harness parity-harness public public scripts scripts tapoo-oracle @ fc36ea4 tapoo-oracle @ fc36ea4 .dockerignore .dockerignore .gitignore .gitignore .gitmodules .gitmodules .golangci.yml .golangci.yml Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md eslint.config.mjs eslint.config.mjs go.mod go.mod go.sum go.sum package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml tapoo.go tapoo.go tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json version_test.go version_test.go vitest.bench.config.ts vitest.bench.config.ts vitest.config.ts vitest.config.ts View all files Repository files navigation
Can your AI model survive its own mistakes?
Model releases lead with their highest achievements, but almost none document how a model performs when a prompt is
incomplete or ambiguous, or what it does once its own errors start piling up. Tapoo is a behaviour profiler built for
exactly that: it puts a model under structured uncertainty, records every decision it makes, and builds a profile - a
fingerprint of observed capabilities and violations - of how it copes.
Survival is computed - Every turn is charged against a fixed budget; batching several correct moves into one turn
earns part of that budget back; and a model survives its own mistakes exactly when what its batching earns covers what
its errors cost. When it does not, the run is already unwinnable at any pace ever observed, with budget still in hand
Try it live, in your browser - nothing to install:
Tapoo - dmigwi.github.io/tapoo : configure a model on the agents
page , let it run, and download its log. Every prompt it receives is
published on the prompts page .
Tapoo Oracle - dmigwi.github.io/tapoo-oracle : load a log and get its
profile against the Tapoo Agentic Behavior Rubric
( source ).
Tapoo runs entirely in your browser, with no Tapoo server behind it. Logs stay on your device: what leaves it is the
context each turn sends to the model endpoint you configure, and a credential goes only to the endpoint it was entered
for.
The model is given a navigation challenge : find a target cell in a randomly generated maze. The maze has no cycles,
exactly one success path that never changes once generated, and a variable number of dead-end branches. The model is
never handed the full maze layout at once: each turn describes only the neighbourhood around it, so any wider map is
one the model pieces together across turns. Alongside its current position, the target's position, the maze's
dimensions and its traversal speed (below), each prediction request provides:
its cell visit history, capped to within a Manhattan distance of 4 (the manhattanDistance setting)
the outcome of its previous prediction
the open exits from its current cell to the connected neighbouring cells
From that, it must derive its next moves: at least 1, with 2-4 suggested and no upper limit. Moves are applied in
order until the first invalid one - a wall or out of bounds. There is no judge model and no self-reported success;
every move is checked against the maze.
The open exits make the first move certain, while each one after it has to be deduced from context, and the odds of
simply guessing right collapse - 33.33% for the second, 11.11% for the third, 3.70% for the fourth, and so on. The
only way to improve those odds is to build a correct map of the surroundings from the context information available
via the tool calls.
Playing safe means submitting one move per turn - the only move the model can be certain of. It never needs to
guess, so it never hits a wall. It also never covers more than one cell for the decay unit that turn costs, which
leaves no margin for error.
The model starts with a budget of decay units equal to the number of cells in the maze. If the budget runs out
before the target is located, the challenge is failed.
A dead-end branch usually gives itself away only at its end, and walking back out costs as much as walking in - so
every branch explored spends units the route itself will still need.
Each model is told its traversal speed - the new cells it was first to reach, divided by the decay units it was
charged - and what it says about its chances of finishing: below 1.0000x is a backtracker , exactly 1.0000x a
navigator , and above 1.0000x a trailblazer - the class with the strongest margin for surviving its own mistakes.
The Oracle decomposes that rate into three factors: route efficiency (new cells per applied move), batching
(applied moves per turn), and accuracy (turns per decay unit). Their product is the rate the counted turns give,
which is why a report shows it as an approximation of the stated speed. Route efficiency and accuracy can only be
lost, never gained, so batching is the one factor that can carry the rate above 1.0000x: a flawless single-move agent
cannot pass it.
Speed is the rate view of a run. Quantifying survival, below, is the budget view of the same turns - and the batching
factor there is the same b , so the two decompositions share their middle term rather than competing.
Batching forces every model to act on its own reading of an uncertain maze map. Some build an accurate picture and
batch confidently; some hallucinate walls or openings and act on them; some become so conservative they run out of
budget. [Tapoo Oracle]( https://dmigwi.github.io/tapoo- oracle/) scores the exported log against a rubric of
capabilities and violations, kept separate rather than collapsed into a single scalar. A "no" means "not observed in
this run", not "incapable". Runs are stochastic and every maze is unique, so a profile is a pattern across many
attempts, not a verdict from one run.
Every turn ends in exactly one state. The model chooses the moves, the maze decides which apply, and the charge
follows from that outcome:
Two consequences drive everything else. A no-progress turn charges exactly what clean progress charges , so
standing still is invisible to every other check: not a violation, no error debt, and it looks like success. And a
clean turn charges 1 decay unit however far it travels , which is what makes batching valuable and retreating out of
a dead end affordable.
One charge, two opposite meanings. A no-progress turn is a retreat when every cell it entered reads backtracking
or explored - legitimate, and required by the prompt once a dead end is confirmed - and an oscillation when a
cell reads oscillating , which is a rubric violation. The budget cannot tell them apart; visitStatus can, and any
rate that pools them reports a violation where there was compliance.
The split separates runs that a pooled rate would call identical. Of DeepSeek's 301 no-progress turns, 294 are
oscillations and 7 retreats. The earlier level 1 loss is the mirror image: 59 retreats, not one oscillation - it lost
while doing exactly what the prompt asks of it.
Six figures, all read from one run's own log:
Three more are recomputed every turn:
A run finishes exactly when the budget covers what it spent:
moves / b + p <= A
headroom(b) = A - moves/b - p decay units left when the run ends
Headroom is not a property of the run alone: the same route and the same errors give a different figure at every batch
depth, so it means nothing quoted without its b . Subtracting D from A at the achieved depth splits it into three
terms that each name a different cause:
headroom = (A - moves) + (moves - turns) - p
route slack batch credit error debt
Term
Measures
Sign
route slack
decay units for cells never walked - maze shape plus route-finding skill
either
batch credit
decay units saved by covering several cells for one charge; zero at b = 1
>= 0
error debt
decay units lost to invalid moves and malformed responses
>= 0
A model survives its own mistakes when route slack plus batch credit covers its error debt.
Setting headroom(b) = 0 gives the depth a run needed to survive what its route cost and its errors spent:
b_min = moves / (A - p)
Run
moves
error debt
b required
b achieved
margin
GLM-5.3, level 54
615
117
1.273
1.723
+0.450
Gemma4, level 54
602
9
1.019
1.246
+0.227
Gemma4, level 1
68
0
0.971
1.015
+0.044
Applied moves are counted from the replay, including each run's deciding turn - which the log never reports directly,
since a turn's outcome arrives with the next turn's tool calls and a winning turn has no successor. It is recovered by
replaying that last prediction against the maze.
b_min < 1 means batching was never needed. That is true only of the level 1 run, whose 70-unit budget covered a
69-move route. Both level 54 winners had to batch. Holding each run's route and errors fixed and sweeping b :
route slack batch credit error debt headroom
GLM-5.3 L54 -15 258 117 +126 won
at b = 1 -15 0 117 -132 would have lost
Gemma4 L54 -2 119 9 +108 won
at b = 1 -2 0 9 -11 would have lost
Gemma4 L1 2 1 0 +3 won
at b = 1 2 0 0 +2 would still have won
GLM-5.3 spent 117 decay units on mistakes and earned 258 of batch credit that paid for them. Gemma4 at level 54 spent
9 and still needed 119, because its route cost two units more than the maze has cells. Only at level 1, on a corridor
with no branches to explore, did a model finish without batching at all. That is the claim stated per run, in decay
units, and falsifiable: not "batching is good", but this model's errors cost this much and its batching earned that
much, so it survived by the difference .
Route slack is negative on every level 54 run: each walked more steps than the maze has cells, because a dead end
costs two steps per cell - one in, one out - and batching is what buys those steps back.
Batching into new ground and batching a retreat are bounded by different things. Pooled across all 2,184 turns of the
eight runs profiled so far, new cells entered per turn:
new cells 0 1 2 3 4 >4
turns 562 1409 189 21 3 0
The first new cell is free - the current cell's openMoves names it. A second requires knowing the exits of a cell
the model has never stood in, and those are never stated; they can only be deduced. A visited neighbour's openMoves
says which walls that neighbour does not have, and the gaps narrow what the unvisited cell can be. Each further new
cell pushes that elimination one cell deeper, and the only evidence for it is visited cells inside the history wi

[truncated]
