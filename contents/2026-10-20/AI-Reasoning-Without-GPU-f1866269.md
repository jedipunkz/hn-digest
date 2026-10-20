---
source: "https://fleetingswallow.com/introducing-minotaur-a-gpuless-ai-architecture/"
hn_url: "https://news.ycombinator.com/item?id=50166901"
title: "AI Reasoning Without GPU"
article_title: "Introducing Minotaur: A GPU'less AI Architecture"
image: "http://fleetingswallow.com/wp-content/uploads/2026/10/AI-Emporium_-Happier-Robots-Brighter-Tomorrow-1024x576.png"
author: "edwardchalk"
captured_at: "2026-10-20T14:44:01Z"
capture_tool: "hn-digest"
hn_id: 50166901
score: 1
comments: 0
posted_at: "2026-10-20T14:22:41Z"
tags:
  - hacker-news
---

# AI Reasoning Without GPU

- HN: [50166901](https://news.ycombinator.com/item?id=50166901)
- Source: [fleetingswallow.com](https://fleetingswallow.com/introducing-minotaur-a-gpuless-ai-architecture/)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T14:22:41Z

## Translation

Title: AI Reasoning Without GPU
Article title: Introducing Minotaur: A GPU'less AI Architecture
Description: Introducing Minotaur: A GPU’less AI Architecture This post introduces Minotaur, an AI architecture we have built on ELC that works out its answers by measuring rather than predicting. It needs no GPU and no training, and every answer it gives can show you exactly what it was measured on. Introductio
[truncated]

Article text:
Introducing Minotaur: A GPU'less AI Architecture
Skip to content
Fleeting Swallow
Introducing Minotaur: A GPU’less AI Architecture
Introducing Minotaur: A GPU’less AI Architecture
This post introduces Minotaur, an AI architecture we have built on ELC that works out its answers by measuring rather than predicting. It needs no GPU and no training, and every answer it gives can show you exactly what it was measured on.
Ask a large language model whether a company’s records follow a regulation and you’ll get a fluent, confident answer. What you won’t get is a measurement. The model has produced what a good answer usually looks like. If it’s wrong, there’s no step you can point to and say: there , that’s where it went astray.
For a lot of everyday uses that’s fine. For audit, compliance and regulation it isn’t. An auditor, a regulator or a court needs an answer that can be replayed, that says what it rests on, and that never claims more than the evidence allows.
Minotaur is our attempt at that kind of AI. It takes its name from MINOS , the minimal ordinal representation of systems that we set out recently on the MINOS page . The TAUR stands for Tension-Apportioning Unit Relations , which is a fairly exact description of what goes on inside it. Every element divides what reaches it between what it lets through and what it holds back, with a drive and a constraint pulling against each other. All of Minotaur’s reasoning is done on the relations between elements like that.
In practice it works like this. You give it a rulebook and the records of how the rules were applied, and you ask it, in ordinary words, whether the records follow the rules. Minotaur measures the answer. If a practice departs from a rule, it tells you by how much and at what point. If something in the chain has gone wrong, it tells you which part. And every answer stays tied to the data it came from, so if that data changes, the answer is withdrawn.
A large language model is really doing two quite different jobs at once:
Language: reading what you’ve written, and writing something sensible back.
Reasoning: working out what follows from what.
Both are done the same way, by predicting the next token, and both run through the same enormous set of weights.
The language part has been good for a while. By 2024 the leading models were scoring above 90% on the standard tests of language and general knowledge, and those tests had stopped telling one model from another. Most of the visible progress since then has been in reasoning. The 2025 Stanford AI Index reports jumps of 49 points on a test of graduate-level science questions and 67 points on a test of fixing real software, each within a year of the test being introduced. It names models that reason before they answer as a new driver. When OpenAI released o1, its first model of that kind, people rating its answers strongly preferred it to its predecessor for coding, maths and data analysis, but not for some tasks that were purely about language. What got better was the reasoning.
So if most of the gain has been in the reasoning, do we need a language model to do the reasoning at all?
Minotaur is built on the answer no . Language is handled at the two ends, where a question comes in and an answer goes out, and those ends are small and replaceable. The reasoning in the middle is done by measurement on ELC, and no language model is involved in it.
The computing power that language models consume isn’t an engineering accident. It follows from how they store what they know.
An LLM holds its knowledge implicitly, spread across hundreds of billions or even trillions of weights learned from text. Nothing in the network marks where a particular fact or rule lives. So every word it produces has to pass through all of those weights, or a large part of them, at roughly two arithmetic operations per weight. A few hundred words from a frontier model comes to something like a thousand million million operations, and learning the weights in the first place takes thousands of GPUs running for months. Hence the data centres.
Minotaur keeps its knowledge out in the open. Part of it is the structure itself: the three roles of Energy, Logic and Control, and the MINOS mathematics behind them. That structure is given, not learned. The rest is the record: the rule as a person accepted it, and the data as it was captured.
So there’s nothing to train and no network to run. Checking 4,000 cases against a rule takes a few tens of sums per case, about a hundred thousand operations in all, against the thousand million million a language model spends on one answer. That’s a difference of roughly ten billion times. An ordinary laptop processor does billions of operations a second, and in our tests a question that sent Minotaur back to the record came back in under a fifth of a second.
So “GPU’less” doesn’t mean we’ve squeezed something large onto a small machine. Nothing in Minotaur needs a GPU in the first place.
All of this rests on one idea, set out at length on the MINOS page . It is that any system, whatever it’s about, can be described with the same three roles:
something drives it ( Energy );
something constrains it ( Control );
and the system has its own way of holding the two together ( Logic ).
The page makes the case for why three roles should be enough , and follows the pattern from the atom up to the stars .
Why does that matter for AI? Think about what a compliance check normally involves. A rule is a paragraph of text and a record is a spreadsheet. They’re different kinds of thing, so something has to translate between them: a person, a piece of hand-written code, or a language model. Mistakes get in during that translation, and the reasons for a conclusion tend to get lost there too.
In ELC, a rule and a record are the same kind of thing. Each is a relation between a drive and a constraint, so the two can be measured against each other directly, with nothing translating in between.
The same idea makes every quantity comparable. MINOS only ever reads a system in shares : of what reaches a part of the system, how much gets through and how much is held back. A share is always a share of that part’s own unit. So days before a foreclosure, nights in hospital and parts per million of a gas all end up as the same kind of number. That’s why one machine can check a mortgage rule, a Medicare rule and an emissions rule without being rebuilt for each. The MINOS page shows how much of nature already works in shares .
Because every part of the chain is described in the same structure, a failure always happens somewhere in particular, and it can be traced to that place. The MINOS page calls this localising failure instead of blaming the whole system .
MINOS also separates what can be read from a system from the frame it happens to be read in. Some things change when the frame changes and some don’t. That difference is what lets Minotaur tell a badly calibrated instrument from a broken rule. It’s the same idea as the reconciliation puzzle on the MINOS page.
None of this would work if every field kept its own vocabulary. It works because there’s one representation for all of them, a lingua franca in which a rule, a record and a verdict can each be written down and measured against the others.
Once a rule and a record are both in ELC form, Minotaur works on them in that form. It doesn’t turn them back into words or symbols to reason about them. We call the things it does with them ELC-native operations , and there are three parts to them: getting things into ELC, reading them, and the operations themselves.
Nothing in the world arrives in ELC form. Rules arrive as prose, and records arrive as columns of days, pounds and parts per million. The ELC adapter translates the external world into the ELC world, and translates the answers back out again. The ELC adapter is patent pending.
On the way in, it works out which column of a record is which quantity, and what its unit and its noise are. It turns each accepted rule, such as “appeal: days ≥ 90”, into a relation over that quantity, without anyone writing code for it. And it expresses every value as a share of its own unit, so that the record speaks the same language as everything else inside.
On the way out it reverses all of that, so the person reading the answer sees days and parts per million, not shares.
Once inside, the record and the rule are set side by side. The record is held as a picture of the practice as it actually is, built from its own cases. The rule is held as a picture of the practice as it should be.
Then both are put to the same simple question, case by case: of what arrived, how much got through, and how much was held back? The record answers with what the servicer actually did. The rule answers with what the servicer should have done.
So a paragraph of law and a spreadsheet, which started out as completely different things, have become answers to the same question in the same terms, and two answers to the same question can be laid side by side and measured against each other.
Everything Minotaur concludes comes from a handful of operations, and each one is a question that MINOS lets us ask about the record and the rule:
Compare asks whether the record and the rule are one relation. Wherever cases arrive, does the practice let through and hold back what the rule says it should? If the two agree to within the noise the record declares, the practice follows the rule. If not, the comparison says how far it departs.
Locate asks where they part company. Departures are rarely spread evenly. Locate follows the drive, the quantity reaching the system, and finds the stretch where the practice stops letting through what the rule lets through.
Rescale asks whether the fault is in the frame or in the relation. If the record agrees with the rule once its frame is changed, then the rule and the practice were both fine, and it was the measurement that was off.
Carry asks what the rule would do with a given case. In MINOS a relation carries a state on from one part of a system to the next, so carrying a case through the rule tells us what it would let through or hold. Minotaur only stands behind the answer where the record has real cases nearby.
Rank asks which of several records departs most. Records declare different amounts of noise, so each departure is weighed against its own record’s noise before they’re compared.
Nothing in these five operations predicts anything or rewrites any symbols. Each one asks a question about relations the rule and the record share, and answers it with a few sums per case. That’s why the whole thing runs comfortably without a GPU.
Some people hear all this and take it for a return to the expert systems and logic engines of the 1970s and 80s. In those systems knowledge was written down as symbols and rules, and reasoning meant applying the rules to rewrite the symbols. Minotaur doesn’t work that way at any point.
A symbolic system works on symbols. A rule like “if days ≥ 90 then appeal” is held as a string of tokens, and reasoning means matching and rewriting those tokens according to rules of inference. Whatever meaning there is lives in what the symbols are taken to stand for. So the system knows only what its labels say, and its conclusions inherit every mistake made in putting the world into words.
A language model also works on tokens, statistically rather than logically. It produces the most likely next one, and its answers rest on the statistics of the text it learned from.
Minotaur doesn’t work on tokens at all. Words appear only at its edges, as the rulebook comes in and the answer goes out. Once the ELC adapter has done its job, there are no symbols left to manipulate and nothing to predict. There are only relations carrying shares of real quantities, and measurements between them.
A conclusion isn’t derived from premises; it’s measured on the data , to a stated precision, and it rests on the exact bytes

[truncated]

## Original Extract

Introducing Minotaur: A GPU’less AI Architecture This post introduces Minotaur, an AI architecture we have built on ELC that works out its answers by measuring rather than predicting. It needs no GPU and no training, and every answer it gives can show you exactly what it was measured on. Introductio
[truncated]

Introducing Minotaur: A GPU'less AI Architecture
Skip to content
Fleeting Swallow
Introducing Minotaur: A GPU’less AI Architecture
Introducing Minotaur: A GPU’less AI Architecture
This post introduces Minotaur, an AI architecture we have built on ELC that works out its answers by measuring rather than predicting. It needs no GPU and no training, and every answer it gives can show you exactly what it was measured on.
Ask a large language model whether a company’s records follow a regulation and you’ll get a fluent, confident answer. What you won’t get is a measurement. The model has produced what a good answer usually looks like. If it’s wrong, there’s no step you can point to and say: there , that’s where it went astray.
For a lot of everyday uses that’s fine. For audit, compliance and regulation it isn’t. An auditor, a regulator or a court needs an answer that can be replayed, that says what it rests on, and that never claims more than the evidence allows.
Minotaur is our attempt at that kind of AI. It takes its name from MINOS , the minimal ordinal representation of systems that we set out recently on the MINOS page . The TAUR stands for Tension-Apportioning Unit Relations , which is a fairly exact description of what goes on inside it. Every element divides what reaches it between what it lets through and what it holds back, with a drive and a constraint pulling against each other. All of Minotaur’s reasoning is done on the relations between elements like that.
In practice it works like this. You give it a rulebook and the records of how the rules were applied, and you ask it, in ordinary words, whether the records follow the rules. Minotaur measures the answer. If a practice departs from a rule, it tells you by how much and at what point. If something in the chain has gone wrong, it tells you which part. And every answer stays tied to the data it came from, so if that data changes, the answer is withdrawn.
A large language model is really doing two quite different jobs at once:
Language: reading what you’ve written, and writing something sensible back.
Reasoning: working out what follows from what.
Both are done the same way, by predicting the next token, and both run through the same enormous set of weights.
The language part has been good for a while. By 2024 the leading models were scoring above 90% on the standard tests of language and general knowledge, and those tests had stopped telling one model from another. Most of the visible progress since then has been in reasoning. The 2025 Stanford AI Index reports jumps of 49 points on a test of graduate-level science questions and 67 points on a test of fixing real software, each within a year of the test being introduced. It names models that reason before they answer as a new driver. When OpenAI released o1, its first model of that kind, people rating its answers strongly preferred it to its predecessor for coding, maths and data analysis, but not for some tasks that were purely about language. What got better was the reasoning.
So if most of the gain has been in the reasoning, do we need a language model to do the reasoning at all?
Minotaur is built on the answer no . Language is handled at the two ends, where a question comes in and an answer goes out, and those ends are small and replaceable. The reasoning in the middle is done by measurement on ELC, and no language model is involved in it.
The computing power that language models consume isn’t an engineering accident. It follows from how they store what they know.
An LLM holds its knowledge implicitly, spread across hundreds of billions or even trillions of weights learned from text. Nothing in the network marks where a particular fact or rule lives. So every word it produces has to pass through all of those weights, or a large part of them, at roughly two arithmetic operations per weight. A few hundred words from a frontier model comes to something like a thousand million million operations, and learning the weights in the first place takes thousands of GPUs running for months. Hence the data centres.
Minotaur keeps its knowledge out in the open. Part of it is the structure itself: the three roles of Energy, Logic and Control, and the MINOS mathematics behind them. That structure is given, not learned. The rest is the record: the rule as a person accepted it, and the data as it was captured.
So there’s nothing to train and no network to run. Checking 4,000 cases against a rule takes a few tens of sums per case, about a hundred thousand operations in all, against the thousand million million a language model spends on one answer. That’s a difference of roughly ten billion times. An ordinary laptop processor does billions of operations a second, and in our tests a question that sent Minotaur back to the record came back in under a fifth of a second.
So “GPU’less” doesn’t mean we’ve squeezed something large onto a small machine. Nothing in Minotaur needs a GPU in the first place.
All of this rests on one idea, set out at length on the MINOS page . It is that any system, whatever it’s about, can be described with the same three roles:
something drives it ( Energy );
something constrains it ( Control );
and the system has its own way of holding the two together ( Logic ).
The page makes the case for why three roles should be enough , and follows the pattern from the atom up to the stars .
Why does that matter for AI? Think about what a compliance check normally involves. A rule is a paragraph of text and a record is a spreadsheet. They’re different kinds of thing, so something has to translate between them: a person, a piece of hand-written code, or a language model. Mistakes get in during that translation, and the reasons for a conclusion tend to get lost there too.
In ELC, a rule and a record are the same kind of thing. Each is a relation between a drive and a constraint, so the two can be measured against each other directly, with nothing translating in between.
The same idea makes every quantity comparable. MINOS only ever reads a system in shares : of what reaches a part of the system, how much gets through and how much is held back. A share is always a share of that part’s own unit. So days before a foreclosure, nights in hospital and parts per million of a gas all end up as the same kind of number. That’s why one machine can check a mortgage rule, a Medicare rule and an emissions rule without being rebuilt for each. The MINOS page shows how much of nature already works in shares .
Because every part of the chain is described in the same structure, a failure always happens somewhere in particular, and it can be traced to that place. The MINOS page calls this localising failure instead of blaming the whole system .
MINOS also separates what can be read from a system from the frame it happens to be read in. Some things change when the frame changes and some don’t. That difference is what lets Minotaur tell a badly calibrated instrument from a broken rule. It’s the same idea as the reconciliation puzzle on the MINOS page.
None of this would work if every field kept its own vocabulary. It works because there’s one representation for all of them, a lingua franca in which a rule, a record and a verdict can each be written down and measured against the others.
Once a rule and a record are both in ELC form, Minotaur works on them in that form. It doesn’t turn them back into words or symbols to reason about them. We call the things it does with them ELC-native operations , and there are three parts to them: getting things into ELC, reading them, and the operations themselves.
Nothing in the world arrives in ELC form. Rules arrive as prose, and records arrive as columns of days, pounds and parts per million. The ELC adapter translates the external world into the ELC world, and translates the answers back out again. The ELC adapter is patent pending.
On the way in, it works out which column of a record is which quantity, and what its unit and its noise are. It turns each accepted rule, such as “appeal: days ≥ 90”, into a relation over that quantity, without anyone writing code for it. And it expresses every value as a share of its own unit, so that the record speaks the same language as everything else inside.
On the way out it reverses all of that, so the person reading the answer sees days and parts per million, not shares.
Once inside, the record and the rule are set side by side. The record is held as a picture of the practice as it actually is, built from its own cases. The rule is held as a picture of the practice as it should be.
Then both are put to the same simple question, case by case: of what arrived, how much got through, and how much was held back? The record answers with what the servicer actually did. The rule answers with what the servicer should have done.
So a paragraph of law and a spreadsheet, which started out as completely different things, have become answers to the same question in the same terms, and two answers to the same question can be laid side by side and measured against each other.
Everything Minotaur concludes comes from a handful of operations, and each one is a question that MINOS lets us ask about the record and the rule:
Compare asks whether the record and the rule are one relation. Wherever cases arrive, does the practice let through and hold back what the rule says it should? If the two agree to within the noise the record declares, the practice follows the rule. If not, the comparison says how far it departs.
Locate asks where they part company. Departures are rarely spread evenly. Locate follows the drive, the quantity reaching the system, and finds the stretch where the practice stops letting through what the rule lets through.
Rescale asks whether the fault is in the frame or in the relation. If the record agrees with the rule once its frame is changed, then the rule and the practice were both fine, and it was the measurement that was off.
Carry asks what the rule would do with a given case. In MINOS a relation carries a state on from one part of a system to the next, so carrying a case through the rule tells us what it would let through or hold. Minotaur only stands behind the answer where the record has real cases nearby.
Rank asks which of several records departs most. Records declare different amounts of noise, so each departure is weighed against its own record’s noise before they’re compared.
Nothing in these five operations predicts anything or rewrites any symbols. Each one asks a question about relations the rule and the record share, and answers it with a few sums per case. That’s why the whole thing runs comfortably without a GPU.
Some people hear all this and take it for a return to the expert systems and logic engines of the 1970s and 80s. In those systems knowledge was written down as symbols and rules, and reasoning meant applying the rules to rewrite the symbols. Minotaur doesn’t work that way at any point.
A symbolic system works on symbols. A rule like “if days ≥ 90 then appeal” is held as a string of tokens, and reasoning means matching and rewriting those tokens according to rules of inference. Whatever meaning there is lives in what the symbols are taken to stand for. So the system knows only what its labels say, and its conclusions inherit every mistake made in putting the world into words.
A language model also works on tokens, statistically rather than logically. It produces the most likely next one, and its answers rest on the statistics of the text it learned from.
Minotaur doesn’t work on tokens at all. Words appear only at its edges, as the rulebook comes in and the answer goes out. Once the ELC adapter has done its job, there are no symbols left to manipulate and nothing to predict. There are only relations carrying shares of real quantities, and measurements between them.
A conclusion isn’t derived from premises; it’s measured on the data , to a stated precision, and it rests on the exact bytes

[truncated]
