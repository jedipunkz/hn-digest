---
source: "https://latent.space/p/ai-is-getting-good-at-math/"
hn_url: "https://news.ycombinator.com/item?id=50237104"
title: "AI is getting good at math, here's why"
article_title: "AI Is Getting Good at Math. Here’s Why. — Latent Space"
image: "https://latent.space/og/card/post/ai-is-getting-good-at-math/81d5ad0c4d69fdf1e50b3e421806a0b816e1e19166d1c828fddd82ddafc35e6a.png"
author: "ankit219"
captured_at: "2026-10-25T23:51:43Z"
capture_tool: "hn-digest"
hn_id: 50237104
score: 1
comments: 0
posted_at: "2026-10-25T23:30:52Z"
tags:
  - hacker-news
---

# AI is getting good at math, here's why

- HN: [50237104](https://news.ycombinator.com/item?id=50237104)
- Source: [latent.space](https://latent.space/p/ai-is-getting-good-at-math/)
- Score: 1
- Comments: 0
- Posted: 2026-10-25T23:30:52Z

## Translation

Title: AI is getting good at math, here's why
Article title: AI Is Getting Good at Math. Here’s Why. — Latent Space
Description: AI models are solving more mathematical problems. A pattern is emerging in their successes that could help us predict which problems they’ll solve next.

Article text:
Skip to content The AI Engineer publication Search ⌘ K Search Latent Space
Filter by type Result type All None Articles Only ⓘ Essays and analysis from Latent Space. AI News Only ⓘ News briefings about AI models, tools, and research. Podcasts Only ⓘ Recorded conversations with AI builders and researchers. Community Only ⓘ Community pages and resources. People Only ⓘ Sourced wiki profiles of people working in AI. Organizations Only ⓘ Sourced wiki profiles of companies and organizations. Topics Only ⓘ Sourced wiki entries explaining AI topics and concepts. Writeups Only ⓘ Grounded explanations linked to original sources. Search mode ? Search tips Keyword matches words; semantic matches meaning and exact titles or names; hybrid combines both. In keyword mode, use quotation marks for exact phrases. Author matches published credits. Year applies to posts; wiki entries have no author or publication year. You can also use author:"Ada" or type:podcast in non-regex searches.
Explore articles, conversations, and the wiki.
Subscribe Welcome back ⌄ Signed in with Google Email preferences Membership & billing Sign out
AI Is Getting Good at Math. Here’s Why.
AI models are solving more mathematical problems. A pattern is emerging in their successes that could help us predict which problems they’ll solve next.
Share On this page ⌄ Problem complexity Characteristics of a given hint Humans solve problems breadth first AI and complexity LLMs are depth first The contrast Failed attempts are hints for LLMs Reconstructing Unit-Distance Proof and Tsimerman’s work Other proofs What changed in models this year? Branching factor An interesting divergence Bonus: Can LLMs ever come up with the General Theory of Relativity? In May, an internal model at OpenAI found a counterexample to Erdős's unit distance conjecture . Open since 1946. Many experts had tried to prove it. OpenAI found a counterexample instead. A month later, OpenAI published a proof of the Cycle Double Cover (CDC) problem , then solutions to ten more open problems by July. And a proof of Navier Stokes in September. Anthropic published a formalization of the fermat's last theorem , and advanced bounds on Reimann's hypothesis . The models are suddenly good at math.
The question that interests me is whether we can predict what kinds of problems these models will solve. What makes a problem tractable for an LLM, and what do these solutions have in common? If AI solves problems differently from humans, identifying the difference should help us decide where to point it. The unit-distance disproof has the richest commentary, so let’s start there.
Here is a simple mental model. To solve any given problem, you need a sequence of hints or insights with work between them. This maps surprisingly well to simple everyday problems and to complex ideas like General Relativity 1 (see Bonus). The sequence looks something like this:
Insight -> time / effort -> insight -> time/effort -> solution.
An expert would need fewer hints than a beginner would, as they would have more background. In his commentary on unit-distance proof, WT Gowers uses a similar idea: the difficulty of the proof is the shortest sequence of hints that would let experts reconstruct it. He calls it Kolmogorov complexity 2 modulo experts. For the unit-distance problem, his hint sequence had three parts:
Take the best known construction and generalize it.
Try a sequence of number fields of increasing degree, but work with prime ideals of bounded norm.
Gowers’ proposed formula for complexity is 2^k + t, where k is the length of the shortest hint sequence and t is the work required to develop it. It is the sort of formula that looks rigorous right up until you try to put a number into it, which is more or less the point. Complexity grows exponentially with the number of insights you need to solve a problem.
Most problem solving is the work between these insights. You have a eureka moment, then spend months developing the hint until another conceptual jump becomes possible. For an expert, time spent is a useful proxy for this work, hence t. Hand someone the first unit-distance hint and they still have to do enough work to reach the generalized construction.
Useful, but post hoc. You can’t look at the problem and predict it would need x number of hints. So, a problem’s age, the number of tries, and the prize attached to it are common proxies to ascribe difficulty.
Characteristics of a given hint
A typical insight changes the formulation of the problem. The representation and understanding of the problem change . Two things decide how much a hint is worth:
How unlikely it looked before, against how obvious it looks after supplying the hint → This is harder than people realize. In unit distance, many experts did not consider the counterexample as likely.
Where it belongs in the sequence: hints are ordered. A third hint is useful only after the work on the first two have made enough progress for it to be actionable. This defeats the naive brute force idea of listing a thousand mathematical techniques and trying every combination 3 . Hints sequence traverses a tree. Initial insights change the problem and reveal the choices that follow.
Humans solve problems breadth first
Combine k and t and you get the human search pattern . Generate candidate directions, explore them shallowly, abandon most, and commit to a few. A promising direction is abandoned if it does not give enough encouragement early 4 . This is a human constraint: Developing an idea can take months. Before committing, a researcher needs reasons to believe it will lead somewhere: conceptual promise, early results, or experience with similar problems. The longer it takes to find out, the higher the bar for trying the idea.
Most research software exists to shrink this time by automating the tedious, mechanical, repeatable work. AI shrinks it far enough that the search pattern itself flips, and everything downstream follows from that.
First, t. The model’s equivalent of effort is compute and tokens, spent on one trajectory (test time compute) or spread across parallel agents (parallel test time compute). Pretraining also supplies knowledge that a human researcher would have to locate or reconstruct. No individual expert is aware of every previous development, so occasionally they end up redoing things already done.
The bigger difference is how models search:
Give the model a single prompt and it will give you a solution. Maybe correct, perhaps not. Ask for a proof and it will give you something that approaches a good proof. CoT for these models runs in a greedy depth first manner too. And the process scales. Give parallel agents ten approaches and you get ten attempts developed toward a conclusion. Maybe all distinct, or some overlapping. Each either hits a roadblock or produces an argument that looks right. Left alone, this gets you beautiful false proofs.
This seems like a waste of compute and resources, until you add an automated verifier 5 that can prune the completed trajectories and identify the correct ones. With this checker, it is now rational to spend enormous compute on uncertain ideas. You find out whether a given idea works or not within minutes. The catch is that the verifier can only judge finished arguments. It cannot tell you whether a half-developed idea is going anywhere.
How good does a hint need to be if you can find out whether it works in five minutes? For humans, the quality bar is high because it could take months of work to pursue a hint. For a model, the bar nearly disappears. The hint can be ordinary, implausible, or barely formed. Following it establishes whether it leads anywhere. LLMs enable us to exchange time for compute, and can help us pursue more problems in a shorter timespan .
Failed attempts are hints for LLMs
The model doesn't need a definitive direction to start. Pretraining already gives it priors about the problem's history and past attempts. Every wrong solution can go back to the model as a negative hint for the next iteration. The system accumulates instructions about what has already failed and why. Roughly, a harness looks like this:
Drafter: writes the prompts given to the solver model. The prompts are based on model priors, mitigating errors from previous runs, and accounting for all the promising paths that went nowhere (these are negative hints).
Solver: Takes the prompts and comes up with one or many candidate solutions. It can run multiple subagents and pursue different paths using parallel test-time compute.
Verifier / Checker: Checks whether the argument is valid and establishes the requested result.
For example, consider this CDC prompt:
Assume for purposes of this task that a complete affirmative proof exists. → This prevents the model from treating the problem’s open status, learned during pretraining, as a reason to stop.
In particular, proofs for special graph classes, constructions of cycle covers with some edges covered other than twice, bounded-length or prescribed-cycle variants, reductions to another unproved conjecture, computational verification through any fixed graph size, and candidate counterexamples without a complete nonexistence certificate are insufficient → This rules out a catalogue of substitutes for solving the actual problem.
Use adversarial agents throughout: every candidate proof must be checked for exact-two multiplicity, repeated-edge closed trails masquerading as cycles, parallel-edge 2-cycles, disconnected graphs, cutvertices, bridges introduced by reductions, and circular use of an equivalent CDC statement → The catalogued failure modes, promoted to standing checks.
Plus one line against learned laziness 6 : Do not return merely because current approaches fail or agents report theorem-strength gaps
OpenAI likely ran multiple rounds of attempts, the drafter took the errors and failed ideas, and steered the next iteration away from those. Every one of these lines is a failed path, written up as an instruction.
Reconstructing Unit-Distance Proof and Tsimerman’s work
Let us read Gowers' three hints back through the harness lens:
Hint 1: Look for a counterexample : This is genuinely hard, because the community and model’s pretraining priors both say the conjecture is true. The drafter most plausibly arrived at it because the solvers kept failing while pursuing a proof. The framing stops the model from treating the accepted view as the direction it must pursue . Every failed attempt makes a counterexample a little less unlikely.
Hint 2: Take the best known construction and generalize it. Closest to a community prior; this is a standard move. The substantial work is developing the generalization far enough to discover what it permits. This is the large-t step. But notice, this is a generalization of a counterexample, which almost nobody was looking for.
Hint 3: Try number fields of increasing degree. Standard machinery once the first two hints have built the space in which it applies.
The order matters. Start at hint three and you don't even know the exercise is disproof. Skip the generalized construction and hint three has nothing to act on. And even with all three, effort remains enormous. Counterexample hint is obvious in hindsight, which makes you wonder why nobody tried it. In fact, some did.
Jacob Tsimerman, recently hired by OpenAI , worked on the counterexample. He knew the number theory. Even considered increasing degrees. And he stopped. Looking for a counterexample was the first insight. It was still only the beginning. What stopped him was not the lack of skill. It was the aptly sane feeling of a man several weeks into a swamp with no evidence that it has another side. His own words:
“A very scary dynamic, hard to think through the analytic regimes and retain guiding intuition, consumes much time and frequently doesn't work out.”
Psychology found this human side of problem-solving decades ago: Metcalfe and Wiebe asked solvers every fifteen se

[truncated]

## Original Extract

AI models are solving more mathematical problems. A pattern is emerging in their successes that could help us predict which problems they’ll solve next.

Skip to content The AI Engineer publication Search ⌘ K Search Latent Space
Filter by type Result type All None Articles Only ⓘ Essays and analysis from Latent Space. AI News Only ⓘ News briefings about AI models, tools, and research. Podcasts Only ⓘ Recorded conversations with AI builders and researchers. Community Only ⓘ Community pages and resources. People Only ⓘ Sourced wiki profiles of people working in AI. Organizations Only ⓘ Sourced wiki profiles of companies and organizations. Topics Only ⓘ Sourced wiki entries explaining AI topics and concepts. Writeups Only ⓘ Grounded explanations linked to original sources. Search mode ? Search tips Keyword matches words; semantic matches meaning and exact titles or names; hybrid combines both. In keyword mode, use quotation marks for exact phrases. Author matches published credits. Year applies to posts; wiki entries have no author or publication year. You can also use author:"Ada" or type:podcast in non-regex searches.
Explore articles, conversations, and the wiki.
Subscribe Welcome back ⌄ Signed in with Google Email preferences Membership & billing Sign out
AI Is Getting Good at Math. Here’s Why.
AI models are solving more mathematical problems. A pattern is emerging in their successes that could help us predict which problems they’ll solve next.
Share On this page ⌄ Problem complexity Characteristics of a given hint Humans solve problems breadth first AI and complexity LLMs are depth first The contrast Failed attempts are hints for LLMs Reconstructing Unit-Distance Proof and Tsimerman’s work Other proofs What changed in models this year? Branching factor An interesting divergence Bonus: Can LLMs ever come up with the General Theory of Relativity? In May, an internal model at OpenAI found a counterexample to Erdős's unit distance conjecture . Open since 1946. Many experts had tried to prove it. OpenAI found a counterexample instead. A month later, OpenAI published a proof of the Cycle Double Cover (CDC) problem , then solutions to ten more open problems by July. And a proof of Navier Stokes in September. Anthropic published a formalization of the fermat's last theorem , and advanced bounds on Reimann's hypothesis . The models are suddenly good at math.
The question that interests me is whether we can predict what kinds of problems these models will solve. What makes a problem tractable for an LLM, and what do these solutions have in common? If AI solves problems differently from humans, identifying the difference should help us decide where to point it. The unit-distance disproof has the richest commentary, so let’s start there.
Here is a simple mental model. To solve any given problem, you need a sequence of hints or insights with work between them. This maps surprisingly well to simple everyday problems and to complex ideas like General Relativity 1 (see Bonus). The sequence looks something like this:
Insight -> time / effort -> insight -> time/effort -> solution.
An expert would need fewer hints than a beginner would, as they would have more background. In his commentary on unit-distance proof, WT Gowers uses a similar idea: the difficulty of the proof is the shortest sequence of hints that would let experts reconstruct it. He calls it Kolmogorov complexity 2 modulo experts. For the unit-distance problem, his hint sequence had three parts:
Take the best known construction and generalize it.
Try a sequence of number fields of increasing degree, but work with prime ideals of bounded norm.
Gowers’ proposed formula for complexity is 2^k + t, where k is the length of the shortest hint sequence and t is the work required to develop it. It is the sort of formula that looks rigorous right up until you try to put a number into it, which is more or less the point. Complexity grows exponentially with the number of insights you need to solve a problem.
Most problem solving is the work between these insights. You have a eureka moment, then spend months developing the hint until another conceptual jump becomes possible. For an expert, time spent is a useful proxy for this work, hence t. Hand someone the first unit-distance hint and they still have to do enough work to reach the generalized construction.
Useful, but post hoc. You can’t look at the problem and predict it would need x number of hints. So, a problem’s age, the number of tries, and the prize attached to it are common proxies to ascribe difficulty.
Characteristics of a given hint
A typical insight changes the formulation of the problem. The representation and understanding of the problem change . Two things decide how much a hint is worth:
How unlikely it looked before, against how obvious it looks after supplying the hint → This is harder than people realize. In unit distance, many experts did not consider the counterexample as likely.
Where it belongs in the sequence: hints are ordered. A third hint is useful only after the work on the first two have made enough progress for it to be actionable. This defeats the naive brute force idea of listing a thousand mathematical techniques and trying every combination 3 . Hints sequence traverses a tree. Initial insights change the problem and reveal the choices that follow.
Humans solve problems breadth first
Combine k and t and you get the human search pattern . Generate candidate directions, explore them shallowly, abandon most, and commit to a few. A promising direction is abandoned if it does not give enough encouragement early 4 . This is a human constraint: Developing an idea can take months. Before committing, a researcher needs reasons to believe it will lead somewhere: conceptual promise, early results, or experience with similar problems. The longer it takes to find out, the higher the bar for trying the idea.
Most research software exists to shrink this time by automating the tedious, mechanical, repeatable work. AI shrinks it far enough that the search pattern itself flips, and everything downstream follows from that.
First, t. The model’s equivalent of effort is compute and tokens, spent on one trajectory (test time compute) or spread across parallel agents (parallel test time compute). Pretraining also supplies knowledge that a human researcher would have to locate or reconstruct. No individual expert is aware of every previous development, so occasionally they end up redoing things already done.
The bigger difference is how models search:
Give the model a single prompt and it will give you a solution. Maybe correct, perhaps not. Ask for a proof and it will give you something that approaches a good proof. CoT for these models runs in a greedy depth first manner too. And the process scales. Give parallel agents ten approaches and you get ten attempts developed toward a conclusion. Maybe all distinct, or some overlapping. Each either hits a roadblock or produces an argument that looks right. Left alone, this gets you beautiful false proofs.
This seems like a waste of compute and resources, until you add an automated verifier 5 that can prune the completed trajectories and identify the correct ones. With this checker, it is now rational to spend enormous compute on uncertain ideas. You find out whether a given idea works or not within minutes. The catch is that the verifier can only judge finished arguments. It cannot tell you whether a half-developed idea is going anywhere.
How good does a hint need to be if you can find out whether it works in five minutes? For humans, the quality bar is high because it could take months of work to pursue a hint. For a model, the bar nearly disappears. The hint can be ordinary, implausible, or barely formed. Following it establishes whether it leads anywhere. LLMs enable us to exchange time for compute, and can help us pursue more problems in a shorter timespan .
Failed attempts are hints for LLMs
The model doesn't need a definitive direction to start. Pretraining already gives it priors about the problem's history and past attempts. Every wrong solution can go back to the model as a negative hint for the next iteration. The system accumulates instructions about what has already failed and why. Roughly, a harness looks like this:
Drafter: writes the prompts given to the solver model. The prompts are based on model priors, mitigating errors from previous runs, and accounting for all the promising paths that went nowhere (these are negative hints).
Solver: Takes the prompts and comes up with one or many candidate solutions. It can run multiple subagents and pursue different paths using parallel test-time compute.
Verifier / Checker: Checks whether the argument is valid and establishes the requested result.
For example, consider this CDC prompt:
Assume for purposes of this task that a complete affirmative proof exists. → This prevents the model from treating the problem’s open status, learned during pretraining, as a reason to stop.
In particular, proofs for special graph classes, constructions of cycle covers with some edges covered other than twice, bounded-length or prescribed-cycle variants, reductions to another unproved conjecture, computational verification through any fixed graph size, and candidate counterexamples without a complete nonexistence certificate are insufficient → This rules out a catalogue of substitutes for solving the actual problem.
Use adversarial agents throughout: every candidate proof must be checked for exact-two multiplicity, repeated-edge closed trails masquerading as cycles, parallel-edge 2-cycles, disconnected graphs, cutvertices, bridges introduced by reductions, and circular use of an equivalent CDC statement → The catalogued failure modes, promoted to standing checks.
Plus one line against learned laziness 6 : Do not return merely because current approaches fail or agents report theorem-strength gaps
OpenAI likely ran multiple rounds of attempts, the drafter took the errors and failed ideas, and steered the next iteration away from those. Every one of these lines is a failed path, written up as an instruction.
Reconstructing Unit-Distance Proof and Tsimerman’s work
Let us read Gowers' three hints back through the harness lens:
Hint 1: Look for a counterexample : This is genuinely hard, because the community and model’s pretraining priors both say the conjecture is true. The drafter most plausibly arrived at it because the solvers kept failing while pursuing a proof. The framing stops the model from treating the accepted view as the direction it must pursue . Every failed attempt makes a counterexample a little less unlikely.
Hint 2: Take the best known construction and generalize it. Closest to a community prior; this is a standard move. The substantial work is developing the generalization far enough to discover what it permits. This is the large-t step. But notice, this is a generalization of a counterexample, which almost nobody was looking for.
Hint 3: Try number fields of increasing degree. Standard machinery once the first two hints have built the space in which it applies.
The order matters. Start at hint three and you don't even know the exercise is disproof. Skip the generalized construction and hint three has nothing to act on. And even with all three, effort remains enormous. Counterexample hint is obvious in hindsight, which makes you wonder why nobody tried it. In fact, some did.
Jacob Tsimerman, recently hired by OpenAI , worked on the counterexample. He knew the number theory. Even considered increasing degrees. And he stopped. Looking for a counterexample was the first insight. It was still only the beginning. What stopped him was not the lack of skill. It was the aptly sane feeling of a man several weeks into a swamp with no evidence that it has another side. His own words:
“A very scary dynamic, hard to think through the analytic regimes and retain guiding intuition, consumes much time and frequently doesn't work out.”
Psychology found this human side of problem-solving decades ago: Metcalfe and Wiebe asked solvers every fifteen se

[truncated]
