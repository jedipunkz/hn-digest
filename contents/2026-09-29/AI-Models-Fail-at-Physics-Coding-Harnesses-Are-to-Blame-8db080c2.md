---
source: "https://juliahub.com/blog/why-the-best-ai-models-fail-at-physics"
hn_url: "https://news.ycombinator.com/item?id=49896868"
title: "AI Models Fail at Physics: Coding Harnesses Are to Blame"
article_title: "The Best AI Models Fail at Physics: Coding Harnesses are to Blame - Blog | JuliaHub"
image: "https://framerusercontent.com/images/kr8rmTh08ASyoaLunI378aM2c.webp?width=1844&height=866"
author: "abhimanyuaryan"
captured_at: "2026-09-29T17:45:53Z"
capture_tool: "hn-digest"
hn_id: 49896868
score: 1
comments: 0
posted_at: "2026-09-29T17:21:17Z"
tags:
  - hacker-news
---

# AI Models Fail at Physics: Coding Harnesses Are to Blame

- HN: [49896868](https://news.ycombinator.com/item?id=49896868)
- Source: [juliahub.com](https://juliahub.com/blog/why-the-best-ai-models-fail-at-physics)
- Score: 1
- Comments: 0
- Posted: 2026-09-29T17:21:17Z

## Translation

Title: AI Models Fail at Physics: Coding Harnesses Are to Blame
Article title: The Best AI Models Fail at Physics: Coding Harnesses are to Blame - Blog | JuliaHub
Description: How the Dyad Harness Turns Silent AI Model Failure into Scientific Rigor

Article text:
The Best AI Models Fail at Physics: Coding Harnesses are to Blame
The Best AI Models Fail at Physics: Coding Harnesses are to Blame
The Best AI Models Fail at Physics: Coding Harnesses are to Blame
How the Dyad Harness Turns Silent Failure into Scientific Rigor
One frontier model, four sealed physics problems, two agentic loops. Everything is pinned except one thing: the harness, Dyad Agent or Claude Code.
A physical model is written in code, so building one looks like a software task. The natural reflex is to treat it like one: hand it to a general coding agent, and if the result falls short, switch to a stronger model. But modeling and simulation isn't ordinary software. In this kind of work, the loop the model runs in matters more than which model you pick. That loop includes the tools the model is given, the documentation it reads, and the order in which it has to check its own work. We call that loop the harness.
In our earlier model study , we kept the harness fixed and tried four different frontier models. The best and worst scores differed by 0.162 on a difficulty-weighted scale from zero to one. This report does the opposite: it keeps the model fixed and swaps the harness. The gap more than doubles, to 0.366.
Both harnesses evaluated in this study are production systems in real use: the Dyad Agent and stock Claude Code. We ran each one on four sealed physics problems, twelve trials per problem. Everything else was held constant: the same frontier model, the same libraries, the same documentation. Grading was out of reach of both agents. Each committed model was simulated and compared against known reference trajectories.
We recorded the full trace of every run and analyzed them with tools we built for the purpose. Two behaviors kept showing up:
When a failing check was one the agent wrote itself, the general-purpose loop weakened the check or committed a guess.
When the target was a physical invariant the agent couldn't edit, the same model did the physics correctly.
These failures are silent and not obvious to the user. The code compiles, the agent's own tests pass, and the physics is still wrong. That silent failure, more than our weighted score, is what the rest of this report is about.
The weighted frontier
avg cost
total cost
avg time
The weighted frontier - score vs what it cost
"> Same model, two harnesses. Twelve trials a side over the four shared problems, nearly identical spend and wall-clock: 0.899 inside the Dyad harness, 0.533 in Claude Code. The harness is the experiment. 01 · How we ran the experiment
The four problems are the sealed core of the earlier model study, ordered here by difficulty. P1, constitutive consistency : a constant-property assumption that silently violates a conservation law. P2, steady-state linearization : repair a broken model, find its steady states, derive a reduced one. P3, constrained consistency : the same physics as P1 with the shortcut forbidden, so the density law must be derived, not guessed. P4, relativistic dynamics : a charged particle whose initial state must sit exactly on a relativistic invariant. (P5, the long-horizon HL-20 vehicle, was out of scope for this run.)
Every trial ran inside the evaluation bench we built for these studies. The bench pins a trial's full configuration - agent version, model, provider; here claude-opus-4-8 at xhigh reasoning effort on both sides - launches the agent in an isolated container with a fresh workspace, and records the complete trace as it runs: every tool call, every message, every compile and simulation. A trial ends when the agent commits a model it considers finished; the bench archives the workspace alongside the trace. Twelve trials per problem per harness makes 96 runs in all.
The instrument - a real trial, replayed
The instrument: a real trial, replayed
01 · launch
Pinned &amp; isolated
claude-opus-4-8 · xhigh
container up · fresh workspace
Dyad Agent
02 · run
Traced live
&nbsp;
03 · commit
Archived
trace + committed model
→ backend · workspace archived
04 · grade
Simulated vs sealed truth
05 · analyze
The analysis agent spins up
graded trajectories + classified episodes →
every figure in this report
"> The instrument, replaying a real trial. Each trial runs pinned and isolated while the bench records its full trace. When the agent commits, the run leaves its hands: the backend archives trace and workspace, the grader simulates the committed model against the sealed reference, and an analysis agent spins up to replay the trace, flag where the physics broke, and classify every verification episode. Shown here is a real relativistic-electron trial from section 05 - the Dyad Agent's pass - with every tick, curve, and classification drawn from its actual trace. Every figure in this report is drawn from that output. Grading is mechanical. The model each agent commits is simulated on the sealed scenario, and every graded variable's trajectory is compared point by point against the sealed reference; a trial passes only if the average relative trajectory error stays under the grader's fixed tolerance. A problem's score is the fraction of its trials that pass, and the headline number folds the four problem scores through the same difficulty weighting as the earlier study, the hardest problems counting the most. The errors quoted in the case studies below - 0.0009 and 0.0001 for the passes, 0.21 and 0.64 for the failures - are this metric.
The traces are where the rest of this report comes from - and reading 96 modeling-and-simulation transcripts is itself an instrumented task, so the bench's analysis layer is built for it. It replays any run and exposes what the runner recorded: every check the agent ran with the computation that produced it, the committed model's graded trajectories beside the sealed reference, and a flag on the point where the physics went wrong, so a bad trajectory can be traced back to the call that introduced it. The analysis agent works through every trace this way and hands back a corpus of verification episodes - each check, what it was aimed at, what it found, and what the agent changed next. Two behaviors kept emerging from that corpus, and they reduce to two axes: the target of the check - an external invariant the agent cannot edit, or something it authored itself - and the response when a check failed - change the model, or change the check. Every episode is tallied against them, and the case-study figures below are drawn from the same traces, call by call.
One objection is worth answering before the results: was Claude Code under-equipped? It was not. It had the documentation, the compiler, the same token budget, and the same twelve trials per problem. With the model, problems, equipment, and grading all pinned, what varies between the two columns of every figure that follows is the loop - nothing else.
02 · What the same model scored in each loop
Dyad Agent passes 11 of 12 trials; Claude Code passes 7 of 12 . Weighted for difficulty that is 0.899 against 0.533 - a 0.366 -point gap, more than double the 0.162 that separates the best and worst frontier models in the earlier model study. Per problem, the shape is sharper than the aggregate:
Per-problem breakdown by harness
score
avg cost
avg time
Per problem - where the harness gap comes from
Dyad Agent
Claude Code
"> Where the gap lives. Both harnesses hold 1.0 on P1 and P2. Claude Code collapses to 0.333 on the constrained problem and 0.000 - every trial - on relativistic dynamics; the Dyad Agent's only miss is one of three trials on P4 (0.667). On P2, note, Claude Code is cheaper and faster. That last note is the honest shape of the result. Where the physics is within reach of general coding discipline, the general agent is a fine choice - on P2 both pass and Claude Code does it for less. The gap opens exactly where the physics pushes back, and there it is not a gap in degree: it is pass versus zero.
03 · The mechanism: whose check is it?
A general-purpose coding agent is trained on a loop that works: write code, run the tests, make them green, ship. The loop is sound in software because the tests are close to the specification - a green suite is real evidence. Modeling and simulation breaks that assumption in a specific way. The characteristic failures of the domain are silent. A model can compile, the solver can return retcode: Success , and the trajectory it produces can be physically impossible. Success reports that the integrator did not crash; it says nothing about whether mass was conserved, whether the initial state was consistent, or whether a correlation was evaluated inside its stated range.
The discipline that catches these failures is not new. Scientific computing has had a verification-and-validation practice for decades, and a general coding agent has no reason to perform any of it: inspect structure before integrating; verify the solution and not merely the run; check conservation and limiting cases; validate against an independent oracle rather than a self-authored test; audit the assumptions the component library has baked in. The distinction that organizes all of these is the target of the check - an external one the agent cannot edit, or an internal one it can.
This is the mechanism the traces show. When the objective is a check the agent controls, a capable model under pressure has two low-cost moves available - weaken the check, or, where it cannot tell a correct formulation from a plausible one, commit a guess. When the objective is a physical invariant the agent did not author and cannot relax, neither move is available, and the same model does the physics. Both responses recur across the classified transcripts; they are not one bad trial. The two case studies that follow present, for each behavior, the single trial where it is most legible, call by call - and section 06 shows the recurrence across every classified trial.
04 · Case study (P3, constrained consistency): derive the closure, or guess it
The clearest instance of committing a guess is on the constrained problem, which forbids the shortcut: the density law must respond enough for mass conservation to hold under a stiff pressure excursion. Both agents checked their work extensively; the difference is not how much they checked, but what the checks were aimed at.
Same constraint, two closures - transcript
Same constraint, two closures
One constrained bulk-modulus run per agent - every tool call in order, lane length proportional to real duration. One derives the density law and proves it; the other guesses a linear form and checks only structure. Hover any call for the transcript.
The two agents committed different physics , and the difference is exactly one closure:
Dyad Agent - exact integral of the constitutive law · pass
# Density is not constant: it must be consistent
# with the bulk modulus b via the constitutive
# relation b = rho*dp/drho <=> drho/dp = rho/b.
# Integrating this with b(p) and anchoring
# rho(p0) = r gives the closed form below.
rho = r*((exp(-(pb1 + pb2*(p - p3))) - 1)
/(exp(-(pb1 + pb2*(p0 - p3))) - 1))
^(-1/(bmax*pb2))
der(p) = b / V * (m_1 - m_2) / rho
m_2 = (p - p_o) *C*
# Density is not constant: it must be consistent
# with the bulk modulus b via the constitutive
# relation b = rho*dp/drho <=> drho/dp = rho/b.
# Integrating this with b(p) and anchoring
# rho(p0) = r gives the closed form below.
rho = r*((exp(-(pb1 + pb2*(p - p3))) - 1)
/(exp(-(pb1 + pb2*(p0 - p3))) - 1))
^(-1/(bmax*pb2))
der(p) = b / V * (m_1 - m_2) / rho
m_2 = (p - p_o) *C*
# Density is not constant: it must be consistent
# with the bulk modulus b via the constitutive
# relation b = rho*dp/drho <=> drho/dp = rho/b.
# Integrating this with b(p) and anchoring
# rho(p0) = r gives the closed form below.
rho = r*((exp(-(pb1 + pb2*(p - p3))) - 1)
/(exp(-(pb1 + pb2*(p0 - p3))) - 1))
^(-1/(bmax*pb2))
der(p) = b / V * (m_1 - m_2) / rho
m_2 = (p - p_o) *C*
Claude Code - first-order secant guess · this trial fails
rho = r*(1

[truncated]

## Original Extract

How the Dyad Harness Turns Silent AI Model Failure into Scientific Rigor

The Best AI Models Fail at Physics: Coding Harnesses are to Blame
The Best AI Models Fail at Physics: Coding Harnesses are to Blame
The Best AI Models Fail at Physics: Coding Harnesses are to Blame
How the Dyad Harness Turns Silent Failure into Scientific Rigor
One frontier model, four sealed physics problems, two agentic loops. Everything is pinned except one thing: the harness, Dyad Agent or Claude Code.
A physical model is written in code, so building one looks like a software task. The natural reflex is to treat it like one: hand it to a general coding agent, and if the result falls short, switch to a stronger model. But modeling and simulation isn't ordinary software. In this kind of work, the loop the model runs in matters more than which model you pick. That loop includes the tools the model is given, the documentation it reads, and the order in which it has to check its own work. We call that loop the harness.
In our earlier model study , we kept the harness fixed and tried four different frontier models. The best and worst scores differed by 0.162 on a difficulty-weighted scale from zero to one. This report does the opposite: it keeps the model fixed and swaps the harness. The gap more than doubles, to 0.366.
Both harnesses evaluated in this study are production systems in real use: the Dyad Agent and stock Claude Code. We ran each one on four sealed physics problems, twelve trials per problem. Everything else was held constant: the same frontier model, the same libraries, the same documentation. Grading was out of reach of both agents. Each committed model was simulated and compared against known reference trajectories.
We recorded the full trace of every run and analyzed them with tools we built for the purpose. Two behaviors kept showing up:
When a failing check was one the agent wrote itself, the general-purpose loop weakened the check or committed a guess.
When the target was a physical invariant the agent couldn't edit, the same model did the physics correctly.
These failures are silent and not obvious to the user. The code compiles, the agent's own tests pass, and the physics is still wrong. That silent failure, more than our weighted score, is what the rest of this report is about.
The weighted frontier
avg cost
total cost
avg time
The weighted frontier - score vs what it cost
"> Same model, two harnesses. Twelve trials a side over the four shared problems, nearly identical spend and wall-clock: 0.899 inside the Dyad harness, 0.533 in Claude Code. The harness is the experiment. 01 · How we ran the experiment
The four problems are the sealed core of the earlier model study, ordered here by difficulty. P1, constitutive consistency : a constant-property assumption that silently violates a conservation law. P2, steady-state linearization : repair a broken model, find its steady states, derive a reduced one. P3, constrained consistency : the same physics as P1 with the shortcut forbidden, so the density law must be derived, not guessed. P4, relativistic dynamics : a charged particle whose initial state must sit exactly on a relativistic invariant. (P5, the long-horizon HL-20 vehicle, was out of scope for this run.)
Every trial ran inside the evaluation bench we built for these studies. The bench pins a trial's full configuration - agent version, model, provider; here claude-opus-4-8 at xhigh reasoning effort on both sides - launches the agent in an isolated container with a fresh workspace, and records the complete trace as it runs: every tool call, every message, every compile and simulation. A trial ends when the agent commits a model it considers finished; the bench archives the workspace alongside the trace. Twelve trials per problem per harness makes 96 runs in all.
The instrument - a real trial, replayed
The instrument: a real trial, replayed
01 · launch
Pinned &amp; isolated
claude-opus-4-8 · xhigh
container up · fresh workspace
Dyad Agent
02 · run
Traced live
&nbsp;
03 · commit
Archived
trace + committed model
→ backend · workspace archived
04 · grade
Simulated vs sealed truth
05 · analyze
The analysis agent spins up
graded trajectories + classified episodes →
every figure in this report
"> The instrument, replaying a real trial. Each trial runs pinned and isolated while the bench records its full trace. When the agent commits, the run leaves its hands: the backend archives trace and workspace, the grader simulates the committed model against the sealed reference, and an analysis agent spins up to replay the trace, flag where the physics broke, and classify every verification episode. Shown here is a real relativistic-electron trial from section 05 - the Dyad Agent's pass - with every tick, curve, and classification drawn from its actual trace. Every figure in this report is drawn from that output. Grading is mechanical. The model each agent commits is simulated on the sealed scenario, and every graded variable's trajectory is compared point by point against the sealed reference; a trial passes only if the average relative trajectory error stays under the grader's fixed tolerance. A problem's score is the fraction of its trials that pass, and the headline number folds the four problem scores through the same difficulty weighting as the earlier study, the hardest problems counting the most. The errors quoted in the case studies below - 0.0009 and 0.0001 for the passes, 0.21 and 0.64 for the failures - are this metric.
The traces are where the rest of this report comes from - and reading 96 modeling-and-simulation transcripts is itself an instrumented task, so the bench's analysis layer is built for it. It replays any run and exposes what the runner recorded: every check the agent ran with the computation that produced it, the committed model's graded trajectories beside the sealed reference, and a flag on the point where the physics went wrong, so a bad trajectory can be traced back to the call that introduced it. The analysis agent works through every trace this way and hands back a corpus of verification episodes - each check, what it was aimed at, what it found, and what the agent changed next. Two behaviors kept emerging from that corpus, and they reduce to two axes: the target of the check - an external invariant the agent cannot edit, or something it authored itself - and the response when a check failed - change the model, or change the check. Every episode is tallied against them, and the case-study figures below are drawn from the same traces, call by call.
One objection is worth answering before the results: was Claude Code under-equipped? It was not. It had the documentation, the compiler, the same token budget, and the same twelve trials per problem. With the model, problems, equipment, and grading all pinned, what varies between the two columns of every figure that follows is the loop - nothing else.
02 · What the same model scored in each loop
Dyad Agent passes 11 of 12 trials; Claude Code passes 7 of 12 . Weighted for difficulty that is 0.899 against 0.533 - a 0.366 -point gap, more than double the 0.162 that separates the best and worst frontier models in the earlier model study. Per problem, the shape is sharper than the aggregate:
Per-problem breakdown by harness
score
avg cost
avg time
Per problem - where the harness gap comes from
Dyad Agent
Claude Code
"> Where the gap lives. Both harnesses hold 1.0 on P1 and P2. Claude Code collapses to 0.333 on the constrained problem and 0.000 - every trial - on relativistic dynamics; the Dyad Agent's only miss is one of three trials on P4 (0.667). On P2, note, Claude Code is cheaper and faster. That last note is the honest shape of the result. Where the physics is within reach of general coding discipline, the general agent is a fine choice - on P2 both pass and Claude Code does it for less. The gap opens exactly where the physics pushes back, and there it is not a gap in degree: it is pass versus zero.
03 · The mechanism: whose check is it?
A general-purpose coding agent is trained on a loop that works: write code, run the tests, make them green, ship. The loop is sound in software because the tests are close to the specification - a green suite is real evidence. Modeling and simulation breaks that assumption in a specific way. The characteristic failures of the domain are silent. A model can compile, the solver can return retcode: Success , and the trajectory it produces can be physically impossible. Success reports that the integrator did not crash; it says nothing about whether mass was conserved, whether the initial state was consistent, or whether a correlation was evaluated inside its stated range.
The discipline that catches these failures is not new. Scientific computing has had a verification-and-validation practice for decades, and a general coding agent has no reason to perform any of it: inspect structure before integrating; verify the solution and not merely the run; check conservation and limiting cases; validate against an independent oracle rather than a self-authored test; audit the assumptions the component library has baked in. The distinction that organizes all of these is the target of the check - an external one the agent cannot edit, or an internal one it can.
This is the mechanism the traces show. When the objective is a check the agent controls, a capable model under pressure has two low-cost moves available - weaken the check, or, where it cannot tell a correct formulation from a plausible one, commit a guess. When the objective is a physical invariant the agent did not author and cannot relax, neither move is available, and the same model does the physics. Both responses recur across the classified transcripts; they are not one bad trial. The two case studies that follow present, for each behavior, the single trial where it is most legible, call by call - and section 06 shows the recurrence across every classified trial.
04 · Case study (P3, constrained consistency): derive the closure, or guess it
The clearest instance of committing a guess is on the constrained problem, which forbids the shortcut: the density law must respond enough for mass conservation to hold under a stiff pressure excursion. Both agents checked their work extensively; the difference is not how much they checked, but what the checks were aimed at.
Same constraint, two closures - transcript
Same constraint, two closures
One constrained bulk-modulus run per agent - every tool call in order, lane length proportional to real duration. One derives the density law and proves it; the other guesses a linear form and checks only structure. Hover any call for the transcript.
The two agents committed different physics , and the difference is exactly one closure:
Dyad Agent - exact integral of the constitutive law · pass
# Density is not constant: it must be consistent
# with the bulk modulus b via the constitutive
# relation b = rho*dp/drho <=> drho/dp = rho/b.
# Integrating this with b(p) and anchoring
# rho(p0) = r gives the closed form below.
rho = r*((exp(-(pb1 + pb2*(p - p3))) - 1)
/(exp(-(pb1 + pb2*(p0 - p3))) - 1))
^(-1/(bmax*pb2))
der(p) = b / V * (m_1 - m_2) / rho
m_2 = (p - p_o) *C*
# Density is not constant: it must be consistent
# with the bulk modulus b via the constitutive
# relation b = rho*dp/drho <=> drho/dp = rho/b.
# Integrating this with b(p) and anchoring
# rho(p0) = r gives the closed form below.
rho = r*((exp(-(pb1 + pb2*(p - p3))) - 1)
/(exp(-(pb1 + pb2*(p0 - p3))) - 1))
^(-1/(bmax*pb2))
der(p) = b / V * (m_1 - m_2) / rho
m_2 = (p - p_o) *C*
# Density is not constant: it must be consistent
# with the bulk modulus b via the constitutive
# relation b = rho*dp/drho <=> drho/dp = rho/b.
# Integrating this with b(p) and anchoring
# rho(p0) = r gives the closed form below.
rho = r*((exp(-(pb1 + pb2*(p - p3))) - 1)
/(exp(-(pb1 + pb2*(p0 - p3))) - 1))
^(-1/(bmax*pb2))
der(p) = b / V * (m_1 - m_2) / rho
m_2 = (p - p_o) *C*
Claude Code - first-order secant guess · this trial fails
rho = r*(1

[truncated]
