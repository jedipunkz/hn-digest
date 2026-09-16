---
source: "https://airbus.github.io/scikit-decide/"
hn_url: "https://news.ycombinator.com/item?id=49723963"
title: "Scikit-Decide AI Framework for RL, Auto Planning and Scheduling"
article_title: "Scikit-decide"
image: ""
author: "Bluestein"
captured_at: "2026-09-16T09:42:28Z"
capture_tool: "hn-digest"
hn_id: 49723963
score: 1
comments: 0
posted_at: "2026-09-16T09:26:41Z"
tags:
  - hacker-news
---

# Scikit-Decide AI Framework for RL, Auto Planning and Scheduling

- HN: [49723963](https://news.ycombinator.com/item?id=49723963)
- Source: [airbus.github.io](https://airbus.github.io/scikit-decide/)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T09:26:41Z

## Translation

Title: Scikit-Decide AI Framework for RL, Auto Planning and Scheduling
Article title: Scikit-decide
Description: This is scikit-decide documentation

Article text:
Scikit-decide
Home
Install
Guide
Notebooks
Code generators
Reference
Contribute
GitHub
(opens new window)
Home
Install
Guide
Notebooks
Code generators
Reference
Contribute
GitHub
(opens new window) Home
AI framework for Reinforcement Learning, Automated Planning and Scheduling
Describe your decision-making problem once and auto-match compatible solvers.
Enjoy a growing list of domains & solvers catalog, supported by the community.
Scikit-decide is open source and is able to wrap existing state-of-the-art domains/solvers.
Scikit-decide is an AI framework for Reinforcement Learning, Automated Planning and Scheduling.
This framework was initiated at Airbus (opens new window) AI Research and notably received contributions through the ANITI (opens new window) and TUPLES (opens new window) projects, and also from ANU (opens new window) .
Please refer to the Guide and Reference sections at the top to learn how to use scikit-decide.
Problem solving: describe your decision-making problem once and auto-match compatible solvers.
For instance planning/scheduling problems can be solved by RL solvers using GNNs.
Growing catalog: enjoy a growing list of domains & solvers catalog, supported by the community.
Open & Extensible: scikit-decide is open source and is able to wrap existing state-of-the-art domains/solvers.
Domains available: Gym(nasium) (opens new window) environments for reinforcement learning (RL)
PDDL (opens new window) (Planning Domain Definition Language) via unified-planning (opens new window) and plado (opens new window) libraries
encoding in gym(nasium) spaces compatible with RL
graph representations for RL (inspired by Lifted Learning Graph (opens new window) ) 🆕
RDDL (opens new window) (Relational Dynamic Influence Diagram Language) using pyrddl-gym (opens new window) library.
Flight planning, based on openap (opens new window) or in-house Poll-Schumann for performance model
Scheduling, based on rcpsp problem from discrete-optimization (opens new window) library
Toy domains like: maze, mastermind, rock-paper-scissors
Solvers available: RL solvers from ray.rllib and stable-baselines3
existing algos with action masking
adaptation of RL algos for graph observation, based on GNNs from pytorch-geometric (opens new window) 🆕
(with sb3 and ray.rllib (old api stack) in last release, only with sb3 on master since migration of ray.rllib to new api stack,
work in progress to do it with ray.rllib + new api stack)
autoregressive models with action masking component by component for parametric actions 🆕
(only with sb3 for now, work in progress for ray.rllib + new api stack)
Planning solvers from unified-planning (opens new window) library
RDDL solvers jax and gurobi-based based on pyRDDLGym-jax and pyRDDLGym-gurobi from pyrddl-gym project (opens new window)
Search solvers coded in scikit-decide library:
A*, AO*, Improved-LAO*
Value Iteration (VI), Policy Iteration (PI)
Labeled RTDP, Learning Real-Time A*
LDFS (Label-correcting Depth-First Search), Iterative Deepening A*
SSiPP (Short-Sighted Planning), FRET (Find, Revise, Eliminate Traps)
iDual (LP-based SSP solver), Goal Probability and Cost Iteration (GPCI)
Best First Width Search, Iterated Width (IW), Rollout IW (RIW)
Monte Carlo Tree Search (MCTS), POMCP
DESPOT, SARSOP, Witness (POMDP solvers)
RTDP-Bel (belief-space RTDP), HSVI / GoalHSVI
SSPReplan, SSPDetHindsight, SSPPlanMerger (determinization approaches)
Multi-Agent RTDP, Multi-Agent Heuristic meta-solver (MAHD)
(Probabilistic) PDDL (PPDDL) solvers:
FF planner
FFReplan / PPDDLReplan (replanning with pluggable inner solvers)
FFDetHindsight / PPDDLDetHindsight (determinization in hindsight)
RFF / PPDDLPlanMerger (plan aggregation into a policy)
PDDL heuristics (with their probabilistic extensions):
Delete-Relaxation heuristics
PDDL+ parser and simulators with Probabilistic PDDL extensions
Lifted applicable action filtering using Clingo
Z3-based event synchronization in python using z3-solver (opens new window)
Evolution strategy: Cartesian Genetic Programming (CGP)
Scheduling solvers from discrete-optimization (opens new window) ,
itself wrapping ortools (opens new window) , gurobi (opens new window) ,
toulbar (opens new window) , minizinc (opens new window) ,
deap (opens new window) (genetic algorithm), didppy (opens new window) (dynamic programming),
and coding local search (hill climber, simulated annealing), Large Neighborhood Search (LNS), and
genetic programming based hyper-heuristic (GPHH)
Tuning solvers hyperparameters hyperparameters definition

## Original Extract

This is scikit-decide documentation

Scikit-decide
Home
Install
Guide
Notebooks
Code generators
Reference
Contribute
GitHub
(opens new window)
Home
Install
Guide
Notebooks
Code generators
Reference
Contribute
GitHub
(opens new window) Home
AI framework for Reinforcement Learning, Automated Planning and Scheduling
Describe your decision-making problem once and auto-match compatible solvers.
Enjoy a growing list of domains & solvers catalog, supported by the community.
Scikit-decide is open source and is able to wrap existing state-of-the-art domains/solvers.
Scikit-decide is an AI framework for Reinforcement Learning, Automated Planning and Scheduling.
This framework was initiated at Airbus (opens new window) AI Research and notably received contributions through the ANITI (opens new window) and TUPLES (opens new window) projects, and also from ANU (opens new window) .
Please refer to the Guide and Reference sections at the top to learn how to use scikit-decide.
Problem solving: describe your decision-making problem once and auto-match compatible solvers.
For instance planning/scheduling problems can be solved by RL solvers using GNNs.
Growing catalog: enjoy a growing list of domains & solvers catalog, supported by the community.
Open & Extensible: scikit-decide is open source and is able to wrap existing state-of-the-art domains/solvers.
Domains available: Gym(nasium) (opens new window) environments for reinforcement learning (RL)
PDDL (opens new window) (Planning Domain Definition Language) via unified-planning (opens new window) and plado (opens new window) libraries
encoding in gym(nasium) spaces compatible with RL
graph representations for RL (inspired by Lifted Learning Graph (opens new window) ) 🆕
RDDL (opens new window) (Relational Dynamic Influence Diagram Language) using pyrddl-gym (opens new window) library.
Flight planning, based on openap (opens new window) or in-house Poll-Schumann for performance model
Scheduling, based on rcpsp problem from discrete-optimization (opens new window) library
Toy domains like: maze, mastermind, rock-paper-scissors
Solvers available: RL solvers from ray.rllib and stable-baselines3
existing algos with action masking
adaptation of RL algos for graph observation, based on GNNs from pytorch-geometric (opens new window) 🆕
(with sb3 and ray.rllib (old api stack) in last release, only with sb3 on master since migration of ray.rllib to new api stack,
work in progress to do it with ray.rllib + new api stack)
autoregressive models with action masking component by component for parametric actions 🆕
(only with sb3 for now, work in progress for ray.rllib + new api stack)
Planning solvers from unified-planning (opens new window) library
RDDL solvers jax and gurobi-based based on pyRDDLGym-jax and pyRDDLGym-gurobi from pyrddl-gym project (opens new window)
Search solvers coded in scikit-decide library:
A*, AO*, Improved-LAO*
Value Iteration (VI), Policy Iteration (PI)
Labeled RTDP, Learning Real-Time A*
LDFS (Label-correcting Depth-First Search), Iterative Deepening A*
SSiPP (Short-Sighted Planning), FRET (Find, Revise, Eliminate Traps)
iDual (LP-based SSP solver), Goal Probability and Cost Iteration (GPCI)
Best First Width Search, Iterated Width (IW), Rollout IW (RIW)
Monte Carlo Tree Search (MCTS), POMCP
DESPOT, SARSOP, Witness (POMDP solvers)
RTDP-Bel (belief-space RTDP), HSVI / GoalHSVI
SSPReplan, SSPDetHindsight, SSPPlanMerger (determinization approaches)
Multi-Agent RTDP, Multi-Agent Heuristic meta-solver (MAHD)
(Probabilistic) PDDL (PPDDL) solvers:
FF planner
FFReplan / PPDDLReplan (replanning with pluggable inner solvers)
FFDetHindsight / PPDDLDetHindsight (determinization in hindsight)
RFF / PPDDLPlanMerger (plan aggregation into a policy)
PDDL heuristics (with their probabilistic extensions):
Delete-Relaxation heuristics
PDDL+ parser and simulators with Probabilistic PDDL extensions
Lifted applicable action filtering using Clingo
Z3-based event synchronization in python using z3-solver (opens new window)
Evolution strategy: Cartesian Genetic Programming (CGP)
Scheduling solvers from discrete-optimization (opens new window) ,
itself wrapping ortools (opens new window) , gurobi (opens new window) ,
toulbar (opens new window) , minizinc (opens new window) ,
deap (opens new window) (genetic algorithm), didppy (opens new window) (dynamic programming),
and coding local search (hill climber, simulated annealing), Large Neighborhood Search (LNS), and
genetic programming based hyper-heuristic (GPHH)
Tuning solvers hyperparameters hyperparameters definition
