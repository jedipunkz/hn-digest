---
source: "https://riemannhypothesis.blog/2026/11/revisiting-my-masters-thesis-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=50539404"
title: "Revisiting my Master's thesis with AI – Understanding the Riemann Hypothesis"
article_title: "Revisiting my master's thesis with AI - Understanding the Riemann Hypothesis"
image: "https://riemannhypothesis.blog/logo.png"
author: "kurinikku"
captured_at: "2026-11-17T23:45:57Z"
capture_tool: "hn-digest"
hn_id: 50539404
score: 2
comments: 0
posted_at: "2026-11-17T22:53:03Z"
tags:
  - hacker-news
---

# Revisiting my Master's thesis with AI – Understanding the Riemann Hypothesis

- HN: [50539404](https://news.ycombinator.com/item?id=50539404)
- Source: [riemannhypothesis.blog](https://riemannhypothesis.blog/2026/11/revisiting-my-masters-thesis-with-ai/)
- Score: 2
- Comments: 0
- Posted: 2026-11-17T22:53:03Z

## Translation

Title: Revisiting my Master's thesis with AI – Understanding the Riemann Hypothesis
Article title: Revisiting my master's thesis with AI - Understanding the Riemann Hypothesis
Description: An experiment in how AI-assisted mathematical research works in practice.

Article text:
Toggle navigation
Understanding the Riemann Hypothesis
About
About the blog
About me
Projects
Recommend.Games
Diris
Mathematical papers
Ricochet Robots solver
TeXCVMaker
Revisiting my master's thesis with AI
An experiment in how AI-assisted mathematical research works in practice.
There has been rather a lot of excitement about AI doing mathematics lately. In September, OpenAI claimed a solution to a version of the Navier–Stokes Millennium Prize Problem; in October it also claimed a proof that Landau–Siegel zeros do not exist. Those are quite some claims to put between the usual product announcements.
My immediate question was: how does this actually work?
Obviously, there is more to it than opening ChatGPT and typing:
solve navier stokes. think real hard. no mistakes!
OpenAI’s account of its Navier–Stokes effort describes an internal model, roughly 10,000 concurrent agents in the successful group, 88 hours to a proposed solution and a further 17 hours of Lean formalisation and verification. Groups explored different problem variants; an intermediate Euler result was fed back into the search; Codex consolidated insights between groups. That is a substantial research process behind a short description of a breakthrough.
The announcement also came amid a serious dispute over unpublished research and scientific credit. In his public statement , Tristan Buckmaster questioned whether his and Levent Alpöge’s private Codex sessions had informed OpenAI’s work and reported demands to remove Alpöge from authorship. 1 OpenAI denies using their unpublished work or Codex sessions . Buckmaster himself said he did not know whether their data had been used; his account of the authorship discussions remains deeply troubling.
The Siegel-zero story also had a subsequent human exposition and simplification, including papers by William Banks and Bryce Kerr . I am not going to adjudicate either claimed breakthrough here. What interested me was the work between posing a question and obtaining an argument someone could examine.
What is an agent , in this setting? What do thousands of them do that one cannot? Who chooses what they work on, and how do their answers turn into a proof? Could an ordinary person experiment with something similar using the tools available on a normal subscription?
I tried it with the problem from my master’s thesis.
Two papers and a very familiar problem
In 2011, my supervisor gave me two papers. One concerned primes in arithmetic progressions with moduli taken from polynomial sequences ( PDF ); the other concerned the variance of squarefree numbers in arithmetic progressions . My assignment was to combine the techniques.
That became my master’s thesis .
This needs a little perspective. My supervisor could probably have proved the result over a weekend. The point of giving it to me was that I would learn the techniques: understand the papers, see how their arguments fit together and work through the places where the combination required something new.
It seemed a good test for AI for much the same reason. This was a problem about adapting established methods, with relevant literature to read and a reasonably concrete target. I wanted to learn how to organise an attempt, including what would go wrong.
Squarefree numbers in arithmetic progressions
For readers who would like to know what the machines were actually asked, here is the mathematical version. A number is squarefree if no square greater than one divides it. Consequently, each prime factor occurs at most once. For example, \(30=2\cdot3\cdot5\) is squarefree, whereas \(12=2^2\cdot3\) is not, since it is divisible by 4.
Take the simplest arithmetic progressions, the odd and even numbers. The squarefree entries are in bold :
Every second even number is divisible by 4 and therefore excluded. Odd numbers have no such restriction from the prime 2. In the long run, about \(8/\pi^2\approx81.1\%\) of odd numbers and \(4/\pi^2\approx40.5\%\) of even numbers are squarefree.
Write \(Q(x;q,l)\) for the number of squarefree positive integers up to \(x\) that are congruent to \(l\) modulo \(q\). For each fixed progression, this count has a linear main term \(g(q,l)x\). 2 The coefficient is positive unless a prime square divides both \(q\) and \(l\), forcing every term to have that square divisor. The multiples of 4 are an example of this obstruction.
Primes, by comparison, become progressively rarer and have zero unweighted density. For a fixed modulus, weighting each prime \(p\) by \(\log p\) gives a linear main term in each coprime residue class. 3 Squarefree numbers need no extra weight and can have positive density even in non-coprime classes, as the even numbers illustrate.
Knowing the density gives a first approximation. To study its accuracy, we look at the discrepancy
\[
E(x;q,l)=Q(x;q,l)-g(q,l)x.
\]
How large are these errors? As the modulus grows, each progression contains fewer terms up to \(x\), and a fixed-modulus density formula alone tells us little about its accuracy. Controlling every individual progression can be difficult. A mean-square estimate instead describes the typical size of the discrepancy and limits how many progressions can have unusually large errors. Squaring prevents positive and negative errors from cancelling; dividing their sum by the number of classes gives the mean-square discrepancy.
The classical variance problem averages over all moduli up to a bound. My problem instead selected polynomial moduli, such as \(q=k^2+1\): \(2,5,10,17,26,\ldots\). Only about \(\sqrt{Y}\) of these lie below \(Y\). A good average over all moduli might conceal much worse behaviour on such a sparse family. Restricting the moduli asks how well the distribution behaves under this additional arithmetic structure.
There are applications of related estimates in prime number theory. For example, \(p\equiv1\pmod{k^2}\) means that \(p-1\) has the square divisor \(k^2\). Distribution estimates for square moduli helped Baier and Zhao ( PDF ) prove results on primes for which \(p-1\) has a large square divisor. My thesis asked for a squarefree analogue of the sparse-modulus variance studied for primes.
Specifically, for \(f(y)\le x\) it studied
\[
V_f(x,y)=\sum_{1\le k\le y} f^{\prime}(k)
\sum_{l=1}^{f(k)} E(x;f(k),l)^2.
\]
The weight \(f^{\prime}(k)\) compensates for the spacing between successive polynomial values: for \(f(k)=k^2+1\), the gaps are \(2k+1\), close to \(f^{\prime}(k)=2k\). This makes the weighted sparse sum comparable in scale with a sum over all moduli up to \(f(y)\). The question is to obtain an asymptotic formula: a main term, an explicit constant depending on \(f\) and an error estimate small enough to make the formula informative.
You do not need to understand the formula to follow the experiment. The number theory gives us a concrete problem for exploring how agents develop arguments, check each other’s work and assemble a proof. If you want to build your own harness for other kinds of work, those are useful questions too.
Before the harness: two single chats
I also tried the problem in two separate chats, one with Claude Opus 5.5 and one with ChatGPT/Codex . Each received the problem statement without the source papers and was instructed to work without web searches or stored context. Claude was allowed to run local computations.
Both produced proof sketches using elementary correlations between pairs of squarefree numbers, estimating the variance at each modulus before averaging over polynomial values. Both proposed the lower threshold \(x^{7/9+\delta}\), with an upper bound \(f(y)\le x^{1-\delta}\). Comparing their constants with the thesis led to the factor-of-two correction described below . This happened before the full harness was run.
Those sketches provide a useful baseline for the later experiments. They had not undergone the proposition-by-proposition reviews or manuscript checks described below, and my surviving records of the chats are the contemporary summaries linked above. The neutral harness run eventually reached the same lower threshold. Its contribution was developing the argument in detail, repeatedly checking it and preserving the work in an inspectable form.
The harness is 100% AI slop. The source is a horrible, unreadable mess; protective equipment is recommended before following the link.
The models also proposed the workflow and wrote the prompts. Deciding how to decompose a problem, what information to give each agent and how to check the results is a place where human expertise could make a real difference. I largely outsourced that judgment too, asking AI how to organise the attempt and having it implement the design. Even setting up the experiment was something I could delegate.
Starting separate model sessions
An agent in this setup is a fresh invocation of a command-line model, using Claude Code or Codex. The underlying commands are claude -p and codex exec . A Python driver prepares the working folder, supplies a prompt and selected files, starts the process and collects what it produces.
The completed experiments used Claude Opus 5.5, Claude Sonnet 5.5 and GPT-6.1 Sol, mostly at high effort. “High” is a provider’s setting, not a common unit of reasoning.
Each invocation receives instructions and source material for a particular task: a planner develops a proof strategy, a prover tackles one proposition and a referee checks a proposed proof.
A prover’s folder looks approximately like this:
problem.md the overall mathematical question
statement.md this agent's proposition
deps/ statements it may assume
papers/ the supplied literature
Its output is a proof file, or a report that the statement is defective. A referee receives a separate copy of the relevant inputs. It does not see the prover’s conversation or the other referee’s verdict.
How results move between agents
The agents in my harness do not chat with one another. The driver copies files between stages; those copies determine what each model can know from the experiment. Persistent state lives in those files, so the next session can be fresh.
I limited the system to two simultaneous agent invocations, largely because I was working within ordinary subscription limits. Over the course of the experiment there were many sessions, with different roles. Two at a time was quite enough to keep me occupied.
OpenAI’s published account describes communicating groups and redistribution of intermediate results; mine used explicit file handoffs and a mostly fixed sequence of stages. 4
First find out what to ask the provers
The tempting place to start is “prove the theorem”. We put several stages before that.
Prior knowledge and excluded sources
Fresh models were first asked what they already knew about the problem. My thesis is online, and a later paper by Fabian Dehnert treats the same topic. 5 Asking a model to rediscover an answer it has just retrieved would be a rather different experiment. Neither a failure to recall a result nor a different proof strategy can establish that it was absent from training data, but the probes at least recorded what the models volunteered.
We also restricted the usable literature to works from 2010 or earlier, plus the explicitly supplied papers, including the 2011 Brüdern–Wooley paper. My thesis and Dehnert’s work were excluded. This was an awkward, artificial restriction: in actual frontier research you would normally want to build on all relevant existing work. Here we wanted a somewhat apples-to-apples comparison with my original assignment, rather than an answer assembled from fifteen subsequent years of research. 6
There were transcript scans for excluded sources and restrictions on tools, but the real runs used host directories rather than isolated containers. Search snippets did sometimes contain excluded titles or links. This was an attempt to preserve and audit independence, with practical limitations.
Next came numerical experiments. Two agents independently wrote code to calcul

[truncated]

## Original Extract

An experiment in how AI-assisted mathematical research works in practice.

Toggle navigation
Understanding the Riemann Hypothesis
About
About the blog
About me
Projects
Recommend.Games
Diris
Mathematical papers
Ricochet Robots solver
TeXCVMaker
Revisiting my master's thesis with AI
An experiment in how AI-assisted mathematical research works in practice.
There has been rather a lot of excitement about AI doing mathematics lately. In September, OpenAI claimed a solution to a version of the Navier–Stokes Millennium Prize Problem; in October it also claimed a proof that Landau–Siegel zeros do not exist. Those are quite some claims to put between the usual product announcements.
My immediate question was: how does this actually work?
Obviously, there is more to it than opening ChatGPT and typing:
solve navier stokes. think real hard. no mistakes!
OpenAI’s account of its Navier–Stokes effort describes an internal model, roughly 10,000 concurrent agents in the successful group, 88 hours to a proposed solution and a further 17 hours of Lean formalisation and verification. Groups explored different problem variants; an intermediate Euler result was fed back into the search; Codex consolidated insights between groups. That is a substantial research process behind a short description of a breakthrough.
The announcement also came amid a serious dispute over unpublished research and scientific credit. In his public statement , Tristan Buckmaster questioned whether his and Levent Alpöge’s private Codex sessions had informed OpenAI’s work and reported demands to remove Alpöge from authorship. 1 OpenAI denies using their unpublished work or Codex sessions . Buckmaster himself said he did not know whether their data had been used; his account of the authorship discussions remains deeply troubling.
The Siegel-zero story also had a subsequent human exposition and simplification, including papers by William Banks and Bryce Kerr . I am not going to adjudicate either claimed breakthrough here. What interested me was the work between posing a question and obtaining an argument someone could examine.
What is an agent , in this setting? What do thousands of them do that one cannot? Who chooses what they work on, and how do their answers turn into a proof? Could an ordinary person experiment with something similar using the tools available on a normal subscription?
I tried it with the problem from my master’s thesis.
Two papers and a very familiar problem
In 2011, my supervisor gave me two papers. One concerned primes in arithmetic progressions with moduli taken from polynomial sequences ( PDF ); the other concerned the variance of squarefree numbers in arithmetic progressions . My assignment was to combine the techniques.
That became my master’s thesis .
This needs a little perspective. My supervisor could probably have proved the result over a weekend. The point of giving it to me was that I would learn the techniques: understand the papers, see how their arguments fit together and work through the places where the combination required something new.
It seemed a good test for AI for much the same reason. This was a problem about adapting established methods, with relevant literature to read and a reasonably concrete target. I wanted to learn how to organise an attempt, including what would go wrong.
Squarefree numbers in arithmetic progressions
For readers who would like to know what the machines were actually asked, here is the mathematical version. A number is squarefree if no square greater than one divides it. Consequently, each prime factor occurs at most once. For example, \(30=2\cdot3\cdot5\) is squarefree, whereas \(12=2^2\cdot3\) is not, since it is divisible by 4.
Take the simplest arithmetic progressions, the odd and even numbers. The squarefree entries are in bold :
Every second even number is divisible by 4 and therefore excluded. Odd numbers have no such restriction from the prime 2. In the long run, about \(8/\pi^2\approx81.1\%\) of odd numbers and \(4/\pi^2\approx40.5\%\) of even numbers are squarefree.
Write \(Q(x;q,l)\) for the number of squarefree positive integers up to \(x\) that are congruent to \(l\) modulo \(q\). For each fixed progression, this count has a linear main term \(g(q,l)x\). 2 The coefficient is positive unless a prime square divides both \(q\) and \(l\), forcing every term to have that square divisor. The multiples of 4 are an example of this obstruction.
Primes, by comparison, become progressively rarer and have zero unweighted density. For a fixed modulus, weighting each prime \(p\) by \(\log p\) gives a linear main term in each coprime residue class. 3 Squarefree numbers need no extra weight and can have positive density even in non-coprime classes, as the even numbers illustrate.
Knowing the density gives a first approximation. To study its accuracy, we look at the discrepancy
\[
E(x;q,l)=Q(x;q,l)-g(q,l)x.
\]
How large are these errors? As the modulus grows, each progression contains fewer terms up to \(x\), and a fixed-modulus density formula alone tells us little about its accuracy. Controlling every individual progression can be difficult. A mean-square estimate instead describes the typical size of the discrepancy and limits how many progressions can have unusually large errors. Squaring prevents positive and negative errors from cancelling; dividing their sum by the number of classes gives the mean-square discrepancy.
The classical variance problem averages over all moduli up to a bound. My problem instead selected polynomial moduli, such as \(q=k^2+1\): \(2,5,10,17,26,\ldots\). Only about \(\sqrt{Y}\) of these lie below \(Y\). A good average over all moduli might conceal much worse behaviour on such a sparse family. Restricting the moduli asks how well the distribution behaves under this additional arithmetic structure.
There are applications of related estimates in prime number theory. For example, \(p\equiv1\pmod{k^2}\) means that \(p-1\) has the square divisor \(k^2\). Distribution estimates for square moduli helped Baier and Zhao ( PDF ) prove results on primes for which \(p-1\) has a large square divisor. My thesis asked for a squarefree analogue of the sparse-modulus variance studied for primes.
Specifically, for \(f(y)\le x\) it studied
\[
V_f(x,y)=\sum_{1\le k\le y} f^{\prime}(k)
\sum_{l=1}^{f(k)} E(x;f(k),l)^2.
\]
The weight \(f^{\prime}(k)\) compensates for the spacing between successive polynomial values: for \(f(k)=k^2+1\), the gaps are \(2k+1\), close to \(f^{\prime}(k)=2k\). This makes the weighted sparse sum comparable in scale with a sum over all moduli up to \(f(y)\). The question is to obtain an asymptotic formula: a main term, an explicit constant depending on \(f\) and an error estimate small enough to make the formula informative.
You do not need to understand the formula to follow the experiment. The number theory gives us a concrete problem for exploring how agents develop arguments, check each other’s work and assemble a proof. If you want to build your own harness for other kinds of work, those are useful questions too.
Before the harness: two single chats
I also tried the problem in two separate chats, one with Claude Opus 5.5 and one with ChatGPT/Codex . Each received the problem statement without the source papers and was instructed to work without web searches or stored context. Claude was allowed to run local computations.
Both produced proof sketches using elementary correlations between pairs of squarefree numbers, estimating the variance at each modulus before averaging over polynomial values. Both proposed the lower threshold \(x^{7/9+\delta}\), with an upper bound \(f(y)\le x^{1-\delta}\). Comparing their constants with the thesis led to the factor-of-two correction described below . This happened before the full harness was run.
Those sketches provide a useful baseline for the later experiments. They had not undergone the proposition-by-proposition reviews or manuscript checks described below, and my surviving records of the chats are the contemporary summaries linked above. The neutral harness run eventually reached the same lower threshold. Its contribution was developing the argument in detail, repeatedly checking it and preserving the work in an inspectable form.
The harness is 100% AI slop. The source is a horrible, unreadable mess; protective equipment is recommended before following the link.
The models also proposed the workflow and wrote the prompts. Deciding how to decompose a problem, what information to give each agent and how to check the results is a place where human expertise could make a real difference. I largely outsourced that judgment too, asking AI how to organise the attempt and having it implement the design. Even setting up the experiment was something I could delegate.
Starting separate model sessions
An agent in this setup is a fresh invocation of a command-line model, using Claude Code or Codex. The underlying commands are claude -p and codex exec . A Python driver prepares the working folder, supplies a prompt and selected files, starts the process and collects what it produces.
The completed experiments used Claude Opus 5.5, Claude Sonnet 5.5 and GPT-6.1 Sol, mostly at high effort. “High” is a provider’s setting, not a common unit of reasoning.
Each invocation receives instructions and source material for a particular task: a planner develops a proof strategy, a prover tackles one proposition and a referee checks a proposed proof.
A prover’s folder looks approximately like this:
problem.md the overall mathematical question
statement.md this agent's proposition
deps/ statements it may assume
papers/ the supplied literature
Its output is a proof file, or a report that the statement is defective. A referee receives a separate copy of the relevant inputs. It does not see the prover’s conversation or the other referee’s verdict.
How results move between agents
The agents in my harness do not chat with one another. The driver copies files between stages; those copies determine what each model can know from the experiment. Persistent state lives in those files, so the next session can be fresh.
I limited the system to two simultaneous agent invocations, largely because I was working within ordinary subscription limits. Over the course of the experiment there were many sessions, with different roles. Two at a time was quite enough to keep me occupied.
OpenAI’s published account describes communicating groups and redistribution of intermediate results; mine used explicit file handoffs and a mostly fixed sequence of stages. 4
First find out what to ask the provers
The tempting place to start is “prove the theorem”. We put several stages before that.
Prior knowledge and excluded sources
Fresh models were first asked what they already knew about the problem. My thesis is online, and a later paper by Fabian Dehnert treats the same topic. 5 Asking a model to rediscover an answer it has just retrieved would be a rather different experiment. Neither a failure to recall a result nor a different proof strategy can establish that it was absent from training data, but the probes at least recorded what the models volunteered.
We also restricted the usable literature to works from 2010 or earlier, plus the explicitly supplied papers, including the 2011 Brüdern–Wooley paper. My thesis and Dehnert’s work were excluded. This was an awkward, artificial restriction: in actual frontier research you would normally want to build on all relevant existing work. Here we wanted a somewhat apples-to-apples comparison with my original assignment, rather than an answer assembled from fifteen subsequent years of research. 6
There were transcript scans for excluded sources and restrictions on tools, but the real runs used host directories rather than isolated containers. Search snippets did sometimes contain excluded titles or links. This was an attempt to preserve and audit independence, with practical limitations.
Next came numerical experiments. Two agents independently wrote code to calcul

[truncated]
