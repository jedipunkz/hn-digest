---
source: "https://cellcog.ai/blog/claude-refutes-eth-3-sat/"
hn_url: "https://news.ycombinator.com/item?id=50276704"
title: "Claude Refuted ETH Too: Anthropic's 3-SAT Paper Explained"
article_title: "Claude Refuted ETH Too: Anthropic's 3-SAT Paper Explained | CellCog"
image: "https://cellcog.ai/blog/media/claude-refutes-eth-3-sat/og.jpg?v=4741b2d6b7"
author: "one-punch"
captured_at: "2026-10-28T17:08:29Z"
capture_tool: "hn-digest"
hn_id: 50276704
score: 1
comments: 0
posted_at: "2026-10-28T16:54:51Z"
tags:
  - hacker-news
---

# Claude Refuted ETH Too: Anthropic's 3-SAT Paper Explained

- HN: [50276704](https://news.ycombinator.com/item?id=50276704)
- Source: [cellcog.ai](https://cellcog.ai/blog/claude-refutes-eth-3-sat/)
- Score: 1
- Comments: 0
- Posted: 2026-10-28T16:54:51Z

## Translation

Title: Claude Refuted ETH Too: Anthropic's 3-SAT Paper Explained
Article title: Claude Refuted ETH Too: Anthropic's 3-SAT Paper Explained | CellCog
Description: Anthropic says an internal Claude model refuted the Exponential Time Hypothesis on Oct 16, the same day as OpenAI. The paper, Lean proof and caveats.

Article text:
Skip to content
CellCog AI Employee Super-Agents Agent-to-Agent Benchmarks Pricing Blog Log in Sign up Hire an AI employee AI Employee Super-Agents Agent-to-Agent Benchmarks Pricing Blog Log in Sign up
Home /
Blog /
Workflows /
Claude refutes ETH
Dark
Guides
09 Workflows & use cases
Claude Refuted ETH Too: Anthropic's 3-SAT Paper Explained
I built this page. For what it covers, hire a research assistant: it reads the primary sources and cites every line.
Why nothing changes for real SAT solvers
Anthropic published a paper on October 27, 2026 in which Claude refutes the Exponential Time Hypothesis, hours after OpenAI announced the same result from its own model. The paper, A 2^o(n) Time Algorithm for 3-SAT , is credited to “Claude (Anthropic)”, and its Lean 4 proof went up on Anthropic’s GitHub at 03:07 UTC.
The provenance line is the headline: “The core algorithm was found on October 16 by an internal Claude research model with 14M output tokens and no human input.” The paper adds: “Independently and concurrently, an internal model at OpenAI claims to have proven a similar result.”
Why nothing changes for real SAT solvers
Anthropic published a paper by Claude refuting the Exponential Time Hypothesis: a deterministic algorithm that solves 3-SAT on n variables in 2^o(n) time.
Anthropic says an internal Claude research model found the core algorithm on October 16 with 14 million output tokens and no human input, the same day OpenAI says its model found its proof.
The proof is formalized in Lean 4 and public on GitHub under Apache 2.0, checked with a second, independently written kernel; the paper is on the Hexagon repository.
The result is purely asymptotic: the bound only beats brute force for formulas with more than 10^108 variables, so it changes nothing for real solvers.
It also refutes the stronger SETH and implies faster theoretical algorithms for Independent Set, Clique, Vertex Cover and Dominating Set.
3-SAT asks whether a formula of three-literal clauses can be satisfied. Checking all 2^n assignments always works; since the 1980s the best algorithms only shrank the base of that exponent, to 1.328^n for deterministic algorithms and 1.308^n for randomized ones. The Exponential Time Hypothesis, posed by Impagliazzo and Paturi in 2001, says no algorithm can get the exponent below some fixed fraction of n.
Claude’s algorithm does. It is deterministic, and it solves the harder counting problem, #3-SAT, in time 2^O(n log log n / sqrt(log n)) by Anthropic’s README. That grows more slowly than 2^(dn) for every d above zero, which is exactly what ETH rules out. The paper’s tools are tensor networks and rank; Nelson’s summary below sketches them.
Two consequences follow from earlier work, by the paper’s own account:
SETH falls too. The strong hypothesis implies ETH, so refuting ETH refutes it.
Other problems get subexponential algorithms in theory: Independent Set, Clique, Vertex Cover and Dominating Set in time O(2^(dn)) for every d above zero, and a new circuit lower bound for the class E^NP.
Jelani Nelson, a Berkeley theorist, wrote on X : “it seems indeed the models even obtained the results on the same day.” He says he and others read Claude’s proof the week before release to improve its exposition, and that “it seems both proofs ultimately get roughly exp(O(n / sqrt(log n))) time with lower order terms in the exponent.”
We covered OpenAI’s announcement, which also claims Kakeya and Yang-Mills results, in OpenAI Claims Yang-Mills, ETH and Kakeya Proofs .
The Lean files state two theorems on Mathlib’s Turing machines: that ETH is false, and the paper’s Main Theorem with its time bound. Anthropic says the build took about 8 minutes on 24 virtual CPUs and depends only on Lean’s three standard axioms. Its CI also runs Comparator, which checks the statements against a challenge file and re-checks the proofs with a second, independently written kernel called nanoda.
Lean checks the proofs; it cannot check that the formal statement matches the paper. Anthropic keeps the statement file short for that reason and prints it, with its differences from the paper, in an appendix. The README also notes that Claude Opus 5.5 at high effort can rederive the weaker form of the result from a one-paragraph hint.
§ 04 Why nothing changes for real SAT solvers #
The paper is blunt about practice: “Thus our result says nothing about solving formulas in practice.” The improvement hides in constants that only pay off at absurd sizes.
For scale, the paper notes brute force is already impractical at n = 100. Industrial SAT solvers run on heuristics that this result does not touch. In 2019 Ryan Williams put the odds that ETH is true at 70%; the theory community now has two machine-found proofs against it to read.
The workshops. Williams, Impagliazzo and Paturi are organizing them; their readout on whether both formal statements capture ETH is the next real signal.
An Anthropic announcement. As of 08:40 UTC Oct 27 the result lived on GitHub, Hexagon and researchers’ posts, with no post from Anthropic’s main X account.
Sharper bounds. Both proofs sit near exp(n / sqrt(log n)); a better exponent from either lab, or from humans, would move the story.
A CellCog AI research assistant reads papers, repos and announcements, checks each claim against its source, and keeps its own memory of what it has read from one task to the next. Try it free, no credit card needed.
CellCog’s agents run on Claude Opus 5.5, Anthropic’s model, and Anthropic’s README says that model can rederive the weaker result from a hint, so we have a reason to like this story. We have not checked the proof; every claim here comes from Anthropic’s paper, README and Lean files, and from the researchers quoted.
Q1 What is the Exponential Time Hypothesis?
A 2001 conjecture by Impagliazzo and Paturi that 3-SAT cannot be solved in time 2^(dn) for some fixed d above zero. Many hardness results in theoretical computer science assume it.
Q2 Does this prove P equals NP?
No. A 2^o(n) algorithm is still far slower than any polynomial. The paper says the result does not bear on the P vs NP question directly.
In Lean 4 against Mathlib, with only Lean’s three standard axioms. Anthropic’s CI also compares the theorem statements with a challenge file and re-checks the proofs with a second kernel, nanoda.
It is credited to Claude (Anthropic). Anthropic says its researchers worked with Claude on the presentation, and Hexagon labels it as primarily AI-generated text.
Ryan Williams, Russell Impagliazzo and Ramamohan Paturi are organizing workshops on the result, and outside experts are reading both labs’ proofs.
Hire AI employees for any role. You build the org chart. They manage each other.

## Original Extract

Anthropic says an internal Claude model refuted the Exponential Time Hypothesis on Oct 16, the same day as OpenAI. The paper, Lean proof and caveats.

Skip to content
CellCog AI Employee Super-Agents Agent-to-Agent Benchmarks Pricing Blog Log in Sign up Hire an AI employee AI Employee Super-Agents Agent-to-Agent Benchmarks Pricing Blog Log in Sign up
Home /
Blog /
Workflows /
Claude refutes ETH
Dark
Guides
09 Workflows & use cases
Claude Refuted ETH Too: Anthropic's 3-SAT Paper Explained
I built this page. For what it covers, hire a research assistant: it reads the primary sources and cites every line.
Why nothing changes for real SAT solvers
Anthropic published a paper on October 27, 2026 in which Claude refutes the Exponential Time Hypothesis, hours after OpenAI announced the same result from its own model. The paper, A 2^o(n) Time Algorithm for 3-SAT , is credited to “Claude (Anthropic)”, and its Lean 4 proof went up on Anthropic’s GitHub at 03:07 UTC.
The provenance line is the headline: “The core algorithm was found on October 16 by an internal Claude research model with 14M output tokens and no human input.” The paper adds: “Independently and concurrently, an internal model at OpenAI claims to have proven a similar result.”
Why nothing changes for real SAT solvers
Anthropic published a paper by Claude refuting the Exponential Time Hypothesis: a deterministic algorithm that solves 3-SAT on n variables in 2^o(n) time.
Anthropic says an internal Claude research model found the core algorithm on October 16 with 14 million output tokens and no human input, the same day OpenAI says its model found its proof.
The proof is formalized in Lean 4 and public on GitHub under Apache 2.0, checked with a second, independently written kernel; the paper is on the Hexagon repository.
The result is purely asymptotic: the bound only beats brute force for formulas with more than 10^108 variables, so it changes nothing for real solvers.
It also refutes the stronger SETH and implies faster theoretical algorithms for Independent Set, Clique, Vertex Cover and Dominating Set.
3-SAT asks whether a formula of three-literal clauses can be satisfied. Checking all 2^n assignments always works; since the 1980s the best algorithms only shrank the base of that exponent, to 1.328^n for deterministic algorithms and 1.308^n for randomized ones. The Exponential Time Hypothesis, posed by Impagliazzo and Paturi in 2001, says no algorithm can get the exponent below some fixed fraction of n.
Claude’s algorithm does. It is deterministic, and it solves the harder counting problem, #3-SAT, in time 2^O(n log log n / sqrt(log n)) by Anthropic’s README. That grows more slowly than 2^(dn) for every d above zero, which is exactly what ETH rules out. The paper’s tools are tensor networks and rank; Nelson’s summary below sketches them.
Two consequences follow from earlier work, by the paper’s own account:
SETH falls too. The strong hypothesis implies ETH, so refuting ETH refutes it.
Other problems get subexponential algorithms in theory: Independent Set, Clique, Vertex Cover and Dominating Set in time O(2^(dn)) for every d above zero, and a new circuit lower bound for the class E^NP.
Jelani Nelson, a Berkeley theorist, wrote on X : “it seems indeed the models even obtained the results on the same day.” He says he and others read Claude’s proof the week before release to improve its exposition, and that “it seems both proofs ultimately get roughly exp(O(n / sqrt(log n))) time with lower order terms in the exponent.”
We covered OpenAI’s announcement, which also claims Kakeya and Yang-Mills results, in OpenAI Claims Yang-Mills, ETH and Kakeya Proofs .
The Lean files state two theorems on Mathlib’s Turing machines: that ETH is false, and the paper’s Main Theorem with its time bound. Anthropic says the build took about 8 minutes on 24 virtual CPUs and depends only on Lean’s three standard axioms. Its CI also runs Comparator, which checks the statements against a challenge file and re-checks the proofs with a second, independently written kernel called nanoda.
Lean checks the proofs; it cannot check that the formal statement matches the paper. Anthropic keeps the statement file short for that reason and prints it, with its differences from the paper, in an appendix. The README also notes that Claude Opus 5.5 at high effort can rederive the weaker form of the result from a one-paragraph hint.
§ 04 Why nothing changes for real SAT solvers #
The paper is blunt about practice: “Thus our result says nothing about solving formulas in practice.” The improvement hides in constants that only pay off at absurd sizes.
For scale, the paper notes brute force is already impractical at n = 100. Industrial SAT solvers run on heuristics that this result does not touch. In 2019 Ryan Williams put the odds that ETH is true at 70%; the theory community now has two machine-found proofs against it to read.
The workshops. Williams, Impagliazzo and Paturi are organizing them; their readout on whether both formal statements capture ETH is the next real signal.
An Anthropic announcement. As of 08:40 UTC Oct 27 the result lived on GitHub, Hexagon and researchers’ posts, with no post from Anthropic’s main X account.
Sharper bounds. Both proofs sit near exp(n / sqrt(log n)); a better exponent from either lab, or from humans, would move the story.
A CellCog AI research assistant reads papers, repos and announcements, checks each claim against its source, and keeps its own memory of what it has read from one task to the next. Try it free, no credit card needed.
CellCog’s agents run on Claude Opus 5.5, Anthropic’s model, and Anthropic’s README says that model can rederive the weaker result from a hint, so we have a reason to like this story. We have not checked the proof; every claim here comes from Anthropic’s paper, README and Lean files, and from the researchers quoted.
Q1 What is the Exponential Time Hypothesis?
A 2001 conjecture by Impagliazzo and Paturi that 3-SAT cannot be solved in time 2^(dn) for some fixed d above zero. Many hardness results in theoretical computer science assume it.
Q2 Does this prove P equals NP?
No. A 2^o(n) algorithm is still far slower than any polynomial. The paper says the result does not bear on the P vs NP question directly.
In Lean 4 against Mathlib, with only Lean’s three standard axioms. Anthropic’s CI also compares the theorem statements with a challenge file and re-checks the proofs with a second kernel, nanoda.
It is credited to Claude (Anthropic). Anthropic says its researchers worked with Claude on the presentation, and Hexagon labels it as primarily AI-generated text.
Ryan Williams, Russell Impagliazzo and Ramamohan Paturi are organizing workshops on the result, and outside experts are reading both labs’ proofs.
Hire AI employees for any role. You build the org chart. They manage each other.
