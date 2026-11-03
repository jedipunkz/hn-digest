---
source: "https://www.daweisun.me/blog/stl-planning-rl/"
hn_url: "https://news.ycombinator.com/item?id=50345968"
title: "From Words to Waypoints: Teaching a Small LLM to Plan for Robots"
article_title: "From Words to Waypoints: Teaching a Small LLM to Plan for Robots — Dawei Sun"
image: ""
author: "sundw"
captured_at: "2026-11-03T00:36:59Z"
capture_tool: "hn-digest"
hn_id: 50345968
score: 1
comments: 0
posted_at: "2026-11-02T23:51:20Z"
tags:
  - hacker-news
---

# From Words to Waypoints: Teaching a Small LLM to Plan for Robots

- HN: [50345968](https://news.ycombinator.com/item?id=50345968)
- Source: [www.daweisun.me](https://www.daweisun.me/blog/stl-planning-rl/)
- Score: 1
- Comments: 0
- Posted: 2026-11-02T23:51:20Z

## Translation

Title: From Words to Waypoints: Teaching a Small LLM to Plan for Robots
Article title: From Words to Waypoints: Teaching a Small LLM to Plan for Robots — Dawei Sun

Article text:
Dawei Sun Blog
About
From Words to Waypoints: Teaching a Small LLM to Plan for Robots
Large language models are increasingly asked to act as the high-level planner of a robot: given a task in words or in a formal language, decide where to go, in which order, and when. For tasks written in temporal logic, the classical approach is to encode the whole problem as a mixed-integer linear program (MILP) and hand it to a solver. That works, but it can take a very long time, and the cost grows quickly with the length of the task and the number of conditions. A MILP solver is a powerful but blind search: every operator, at every time step, becomes a set of binary variables and constraints, and the solver has no notion of what they mean. A language model can understand the semantics of the temporal operators, for instance that “stay out of D until K is reached” means visiting K before entering D, and use common sense to come up with a plan directly. This post does not compare the two approaches rigorously, and we have no evidence that an LLM planner is better than a MILP solver; we are showing that the LLM route is possible.
We wanted to know whether a small open model can learn to do this well. Concretely, the goal was to take Qwen3.5-9B and train it with reinforcement learning until it solves robot planning tasks written in temporal logic about as reliably as a much larger model, DeepSeek-V4.1-Flash-1017, which we use as the reference (“the teacher”).
This post tells the story of how we got there, in the order it happened: a first attempt that learned nothing, the fix that made learning possible, and the step that made it fast.
A robot moves in a 2D workspace with rectangular obstacles and named regions. Its task is a specification in Signal Temporal Logic (STL), a language for properties of trajectories over time. The setting follows Multi-agent Motion Planning from Signal Temporal Logic Specifications , restricted to one robot with a speed limit. Any feasible path is accepted; we do not ask for the optimal one. A specification combines two kinds of pieces. Atoms say where the robot is: in(B) holds when the robot is inside region B, and out(A) when it is outside region A, where A, B, K, D and Goal are names of regions on the map. Temporal operators say when an atom has to hold. They are written as single capital letters followed by a time window in seconds: F means “eventually” (at some time in the window), G means “always” (at every time in the window), and p U q means “p holds until q does”. Operators can be nested, as in the last example. Below are some examples of STL specifications:
F [7,21] in(B) : be inside region B at some time between 7 and 21 seconds;
G [0,30] out(A) : stay out of region A for the first 30 seconds;
out(D) U [0,136] in(K) : stay out of D until K is reached, within 136 seconds;
F [153,169] G [0,5] in(Goal) : at some time between 153 and 169, enter the goal and stay there for 5 seconds.
The model reads the map and a list of such clauses and must answer with a timed path, a list of waypoints [[t, x, y], ...] . The robot is a small square that must never touch an obstacle and cannot move faster than its speed limit.
Everything the model knows about the task comes from its prompt. The system prompt explains the robot and path model, defines the STL atoms and operators with a couple of examples, and fixes the answer format; the user message lists the workspace, every obstacle and region as coordinates, the robot’s size, speed limit and start, and the specification clauses. There are no solved examples in the prompt, and every model in this post, the teacher included, gets the same prompt.
You are a high-level motion planner for a single mobile robot in a 2D workspace. Your job is to produce a timed path for the robot that satisfies a Signal Temporal Logic (STL) specification. Any feasible path is accepted; it does not need to be optimal.
# Robot and path model
- The robot is an axis-aligned square with half-width r: at position (x, y) its body occupies [x-r, x+r] x [y-r, y+r]. It translates only (no rotation).
- A path is a list of waypoints [[t0, x0, y0], [t1, x1, y1], ...] with t0 = 0 and (x0, y0) = the start position. Times must be non-decreasing.
- Between consecutive waypoints the robot moves in a straight line at constant speed. The speed on each segment, distance / (t_{k+1} - t_k), must not exceed v_max. To wait, repeat the same position at a later time.
- After the last waypoint the robot stays there forever.
- The robot body must always stay inside the workspace and must never overlap any obstacle (touching the boundary is fine).
# STL specification
All regions are axis-aligned rectangles. Atomic propositions:
- in(R): the robot body is fully inside region R (equivalently, x in [xmin+r, xmax-r] and y in [ymin+r, ymax-r]).
- out(R): the robot body is fully outside region R (it does not overlap R).
Operators (a, b are absolute durations in seconds, measured from the time the formula is evaluated):
- p & q: both hold. p | q: at least one holds.
- F[a,b] p ("eventually"): p holds at some time in [t+a, t+b].
- G[a,b] p ("always"): p holds at every time in [t+a, t+b].
- p U[a,b] q ("until"): q holds at some time t' in [t+a, t+b], and p holds at every time in [t, t'].
Formulas can be nested, e.g. F[0,30] G[0,5] in(A) means "at some time t in [0,30], enter A and stay in A for 5 seconds (during [t, t+5])"; G[0,50] F[0,20] in(B) means "during [0,50], B is always visited again within 20 seconds".
The whole specification is evaluated at time t = 0, and every listed item (phi1, phi2, ...) must hold.
# Planning
You have no tools. Reason carefully about ge
[truncated]
An RL environment needs a stream of tasks that are solvable, checkable, and of the right difficulty. We generate them path first : the generator draws a map, picks a plan (an order of region visits, waits, dwell times, regions to keep away from), builds a path for that plan, and only then derives the STL clauses from the path’s actual timing, with randomized slack in the time windows. Every instance therefore has a known solution. The generator also runs a naive baseline (visit the regions in order, no waiting, no avoidance) and resamples most instances that the baseline already solves, so that the clauses actually interact.
The instances come in eight levels, meant as the rungs of a ladder:
Why a ladder? We train with GRPO, which samples several answers per task (eight, in our case) and reinforces the ones that scored above the group’s average. If all eight fail, or all eight succeed, every answer gets the same score and the task contributes nothing to learning. The policy has to succeed sometimes . A ladder of levels makes it likely that some rung sits in that zone for the current model, both across the instances of a level and across the eight samples of one instance.
For evaluation we hold out 550 instances (50 each from L1–L7 and 200 from L8) and sample one answer per instance. On this set the teacher scores 0.90 overall and 0.80 on L8. The RL setup stays the same throughout the story: in every step we take 32 training tasks, half of them from L8, sample eight answers for each, and give each answer at most 32k tokens. More details: GRPO with the clip-higher trick from DAPO (0.2/0.28), no KL penalty, token-mean loss, LoRA rank 64 on all linear layers, learning rate 1e-5, up to 40 assistant turns per episode, trained with verl.
The first thing we tried was the simplest: give the base model the task and let RL do the rest. The model reasons in text and writes down the path.
It learned nothing. At the start, the base model solved 28% of L1, 24% of L2, 6% of L3 and essentially nothing from L4 upwards (0.06 overall). After 28 RL steps the training reward was still between 0.00 and 0.04, and the evaluation had not moved (0.062, 0.044 and 0.045 at steps 0, 10 and 20). In the first five steps, every group of eight samples on L5, L7 and L8 failed completely, and so did 86% of the groups on L3. The ladder we had built did not exist for this model: almost every rung was out of reach, so there was nothing to reinforce.
Why does it fail? Let us look at an example. In this L2 task the robot must wait in region E for a while, then reach region A without passing through region H, and later come back to E. The model writes almost 500 lines of reasoning; here is some of the model’s response where it works out the path:
- Distance: √((17.8-17.3)² + (13.3-6.5)²) = √(0.25 + 46.24) = √46.49 ≈ 6.82.
- Stay in E from t=6.82 to t=13.82. Duration 7s.
...
- Dist d1 = √((17.8-12.0)² + (13.3-17.9)²) = √(33.64 + 21.16) = √54.8 ≈ 7.40?
- (Previous rough calc was 4.67, mistake in distance estimation).
...
- Segment 3: E (17.3, 6.5) → Point U (17.3, 17.9).
- X is constant 17.3. 17.3 > 15.3. Always out(H). Good.
- E → U: d = 11.4. Time 11.4. Arrive 13.82 + 11.4 = 25.22.
...
<answer>[[0, 17.8, 13.3], [6.82, 17.3, 6.5], [13.82, 17.3, 6.5], [25.22, 17.3, 17.9], ...]</answer>
The checker rejects the answer: one of the segments runs through an obstacle.
The response shows what goes wrong. The plan itself is right: wait in E, go around H, reach A in time. But the model then tries to turn the plan into numbers by itself, computing distances, arrival times and clearances by hand. It even catches one of its own arithmetic mistakes along the way, and still misses another one. LLMs make mistakes in this kind of calculation all the time, and one slip anywhere in a long chain of it is enough to make the whole path infeasible.
The example is typical. We sorted all 550 answers at step 0 by the first problem the checker found:
Only 1% of the answers fail on the logic of the specification; almost all failures are calculation errors in the geometry or the timing.
So the model knows what to do but cannot compute the numbers , and RL cannot teach that from a reward that is almost always zero. We needed to give the model a calculator.
Attempt 2: give the model tools
plan_path(start, t_start, goal_region, avoid=[...], hold_until=...) returns the shortest collision-free path segment from a point to a region at full speed, optionally keeping out of some regions and waiting at the goal. It is complete: if it says there is no path, there is none.
check_path(waypoints) runs the exact checker on a candidate path and explains what fails, for example 'in(A)' never holds within [0, 38] (it holds at times: [41.2, 47]) .
In the system prompt, the paragraph that told the model to work out the geometry by itself is replaced by a description of the tools and a suggested way to use them:
# Tools
You can call these tools (each call costs part of a limited budget):
- plan_path: computes the shortest collision-free path segment from a start point to a region or point, departing at a given time at maximum speed, optionally avoiding listed regions and optionally holding at the goal until a given time. Use it instead of computing geometry yourself.
- check_path: checks a full path and reports speed/collision/boundary violations and which specification items fail and why.
A good strategy: decide the order and timing of region visits that satisfies the specification, build the path segment by segment with plan_path (each segment starts where and when the previous one ended; insert waiting by choosing a later t_start or using hold_until), concatenate the waypoint lists, verify the full path with check_path, fix any reported problem, then give the final answer.
Both tools are fast. plan_path takes about 1–11 ms on average depending on the map (at most 37 ms on the largest mazes), and check_path at most 3 ms, so the exact geometry costs the model almost nothing, in contrast to solving the whole problem as a MILP. Why the gap is so large: the standard encoding of STL as a MILP (Raman et al., Model Predictive Control with Signal Temporal Logic Specifications , CDC 2014) introduces binary variables for every atom and every operator at every time step, and solving a MILP is NP-ha

[truncated]

## Original Extract

Dawei Sun Blog
About
From Words to Waypoints: Teaching a Small LLM to Plan for Robots
Large language models are increasingly asked to act as the high-level planner of a robot: given a task in words or in a formal language, decide where to go, in which order, and when. For tasks written in temporal logic, the classical approach is to encode the whole problem as a mixed-integer linear program (MILP) and hand it to a solver. That works, but it can take a very long time, and the cost grows quickly with the length of the task and the number of conditions. A MILP solver is a powerful but blind search: every operator, at every time step, becomes a set of binary variables and constraints, and the solver has no notion of what they mean. A language model can understand the semantics of the temporal operators, for instance that “stay out of D until K is reached” means visiting K before entering D, and use common sense to come up with a plan directly. This post does not compare the two approaches rigorously, and we have no evidence that an LLM planner is better than a MILP solver; we are showing that the LLM route is possible.
We wanted to know whether a small open model can learn to do this well. Concretely, the goal was to take Qwen3.5-9B and train it with reinforcement learning until it solves robot planning tasks written in temporal logic about as reliably as a much larger model, DeepSeek-V4.1-Flash-1017, which we use as the reference (“the teacher”).
This post tells the story of how we got there, in the order it happened: a first attempt that learned nothing, the fix that made learning possible, and the step that made it fast.
A robot moves in a 2D workspace with rectangular obstacles and named regions. Its task is a specification in Signal Temporal Logic (STL), a language for properties of trajectories over time. The setting follows Multi-agent Motion Planning from Signal Temporal Logic Specifications , restricted to one robot with a speed limit. Any feasible path is accepted; we do not ask for the optimal one. A specification combines two kinds of pieces. Atoms say where the robot is: in(B) holds when the robot is inside region B, and out(A) when it is outside region A, where A, B, K, D and Goal are names of regions on the map. Temporal operators say when an atom has to hold. They are written as single capital letters followed by a time window in seconds: F means “eventually” (at some time in the window), G means “always” (at every time in the window), and p U q means “p holds until q does”. Operators can be nested, as in the last example. Below are some examples of STL specifications:
F [7,21] in(B) : be inside region B at some time between 7 and 21 seconds;
G [0,30] out(A) : stay out of region A for the first 30 seconds;
out(D) U [0,136] in(K) : stay out of D until K is reached, within 136 seconds;
F [153,169] G [0,5] in(Goal) : at some time between 153 and 169, enter the goal and stay there for 5 seconds.
The model reads the map and a list of such clauses and must answer with a timed path, a list of waypoints [[t, x, y], ...] . The robot is a small square that must never touch an obstacle and cannot move faster than its speed limit.
Everything the model knows about the task comes from its prompt. The system prompt explains the robot and path model, defines the STL atoms and operators with a couple of examples, and fixes the answer format; the user message lists the workspace, every obstacle and region as coordinates, the robot’s size, speed limit and start, and the specification clauses. There are no solved examples in the prompt, and every model in this post, the teacher included, gets the same prompt.
You are a high-level motion planner for a single mobile robot in a 2D workspace. Your job is to produce a timed path for the robot that satisfies a Signal Temporal Logic (STL) specification. Any feasible path is accepted; it does not need to be optimal.
# Robot and path model
- The robot is an axis-aligned square with half-width r: at position (x, y) its body occupies [x-r, x+r] x [y-r, y+r]. It translates only (no rotation).
- A path is a list of waypoints [[t0, x0, y0], [t1, x1, y1], ...] with t0 = 0 and (x0, y0) = the start position. Times must be non-decreasing.
- Between consecutive waypoints the robot moves in a straight line at constant speed. The speed on each segment, distance / (t_{k+1} - t_k), must not exceed v_max. To wait, repeat the same position at a later time.
- After the last waypoint the robot stays there forever.
- The robot body must always stay inside the workspace and must never overlap any obstacle (touching the boundary is fine).
# STL specification
All regions are axis-aligned rectangles. Atomic propositions:
- in(R): the robot body is fully inside region R (equivalently, x in [xmin+r, xmax-r] and y in [ymin+r, ymax-r]).
- out(R): the robot body is fully outside region R (it does not overlap R).
Operators (a, b are absolute durations in seconds, measured from the time the formula is evaluated):
- p & q: both hold. p | q: at least one holds.
- F[a,b] p ("eventually"): p holds at some time in [t+a, t+b].
- G[a,b] p ("always"): p holds at every time in [t+a, t+b].
- p U[a,b] q ("until"): q holds at some time t' in [t+a, t+b], and p holds at every time in [t, t'].
Formulas can be nested, e.g. F[0,30] G[0,5] in(A) means "at some time t in [0,30], enter A and stay in A for 5 seconds (during [t, t+5])"; G[0,50] F[0,20] in(B) means "during [0,50], B is always visited again within 20 seconds".
The whole specification is evaluated at time t = 0, and every listed item (phi1, phi2, ...) must hold.
# Planning
You have no tools. Reason carefully about ge
[truncated]
An RL environment needs a stream of tasks that are solvable, checkable, and of the right difficulty. We generate them path first : the generator draws a map, picks a plan (an order of region visits, waits, dwell times, regions to keep away from), builds a path for that plan, and only then derives the STL clauses from the path’s actual timing, with randomized slack in the time windows. Every instance therefore has a known solution. The generator also runs a naive baseline (visit the regions in order, no waiting, no avoidance) and resamples most instances that the baseline already solves, so that the clauses actually interact.
The instances come in eight levels, meant as the rungs of a ladder:
Why a ladder? We train with GRPO, which samples several answers per task (eight, in our case) and reinforces the ones that scored above the group’s average. If all eight fail, or all eight succeed, every answer gets the same score and the task contributes nothing to learning. The policy has to succeed sometimes . A ladder of levels makes it likely that some rung sits in that zone for the current model, both across the instances of a level and across the eight samples of one instance.
For evaluation we hold out 550 instances (50 each from L1–L7 and 200 from L8) and sample one answer per instance. On this set the teacher scores 0.90 overall and 0.80 on L8. The RL setup stays the same throughout the story: in every step we take 32 training tasks, half of them from L8, sample eight answers for each, and give each answer at most 32k tokens. More details: GRPO with the clip-higher trick from DAPO (0.2/0.28), no KL penalty, token-mean loss, LoRA rank 64 on all linear layers, learning rate 1e-5, up to 40 assistant turns per episode, trained with verl.
The first thing we tried was the simplest: give the base model the task and let RL do the rest. The model reasons in text and writes down the path.
It learned nothing. At the start, the base model solved 28% of L1, 24% of L2, 6% of L3 and essentially nothing from L4 upwards (0.06 overall). After 28 RL steps the training reward was still between 0.00 and 0.04, and the evaluation had not moved (0.062, 0.044 and 0.045 at steps 0, 10 and 20). In the first five steps, every group of eight samples on L5, L7 and L8 failed completely, and so did 86% of the groups on L3. The ladder we had built did not exist for this model: almost every rung was out of reach, so there was nothing to reinforce.
Why does it fail? Let us look at an example. In this L2 task the robot must wait in region E for a while, then reach region A without passing through region H, and later come back to E. The model writes almost 500 lines of reasoning; here is some of the model’s response where it works out the path:
- Distance: √((17.8-17.3)² + (13.3-6.5)²) = √(0.25 + 46.24) = √46.49 ≈ 6.82.
- Stay in E from t=6.82 to t=13.82. Duration 7s.
...
- Dist d1 = √((17.8-12.0)² + (13.3-17.9)²) = √(33.64 + 21.16) = √54.8 ≈ 7.40?
- (Previous rough calc was 4.67, mistake in distance estimation).
...
- Segment 3: E (17.3, 6.5) → Point U (17.3, 17.9).
- X is constant 17.3. 17.3 > 15.3. Always out(H). Good.
- E → U: d = 11.4. Time 11.4. Arrive 13.82 + 11.4 = 25.22.
...
<answer>[[0, 17.8, 13.3], [6.82, 17.3, 6.5], [13.82, 17.3, 6.5], [25.22, 17.3, 17.9], ...]</answer>
The checker rejects the answer: one of the segments runs through an obstacle.
The response shows what goes wrong. The plan itself is right: wait in E, go around H, reach A in time. But the model then tries to turn the plan into numbers by itself, computing distances, arrival times and clearances by hand. It even catches one of its own arithmetic mistakes along the way, and still misses another one. LLMs make mistakes in this kind of calculation all the time, and one slip anywhere in a long chain of it is enough to make the whole path infeasible.
The example is typical. We sorted all 550 answers at step 0 by the first problem the checker found:
Only 1% of the answers fail on the logic of the specification; almost all failures are calculation errors in the geometry or the timing.
So the model knows what to do but cannot compute the numbers , and RL cannot teach that from a reward that is almost always zero. We needed to give the model a calculator.
Attempt 2: give the model tools
plan_path(start, t_start, goal_region, avoid=[...], hold_until=...) returns the shortest collision-free path segment from a point to a region at full speed, optionally keeping out of some regions and waiting at the goal. It is complete: if it says there is no path, there is none.
check_path(waypoints) runs the exact checker on a candidate path and explains what fails, for example 'in(A)' never holds within [0, 38] (it holds at times: [41.2, 47]) .
In the system prompt, the paragraph that told the model to work out the geometry by itself is replaced by a description of the tools and a suggested way to use them:
# Tools
You can call these tools (each call costs part of a limited budget):
- plan_path: computes the shortest collision-free path segment from a start point to a region or point, departing at a given time at maximum speed, optionally avoiding listed regions and optionally holding at the goal until a given time. Use it instead of computing geometry yourself.
- check_path: checks a full path and reports speed/collision/boundary violations and which specification items fail and why.
A good strategy: decide the order and timing of region visits that satisfies the specification, build the path segment by segment with plan_path (each segment starts where and when the previous one ended; insert waiting by choosing a later t_start or using hold_until), concatenate the waypoint lists, verify the full path with check_path, fix any reported problem, then give the final answer.
Both tools are fast. plan_path takes about 1–11 ms on average depending on the map (at most 37 ms on the largest mazes), and check_path at most 3 ms, so the exact geometry costs the model almost nothing, in contrast to solving the whole problem as a MILP. Why the gap is so large: the standard encoding of STL as a MILP (Raman et al., Model Predictive Control with Signal Temporal Logic Specifications , CDC 2014) introduces binary variables for every atom and every operator at every time step, and solving a MILP is NP-ha

[truncated]
