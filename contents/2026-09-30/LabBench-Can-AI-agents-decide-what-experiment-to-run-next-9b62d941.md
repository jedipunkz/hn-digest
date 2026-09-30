---
source: "https://gamowlabs.com/labbench-benchmarking-ai-wet-lab-decisions.html"
hn_url: "https://news.ycombinator.com/item?id=49903060"
title: "LabBench: Can AI agents decide what experiment to run next?"
article_title: "LabBench: Can AI agents decide what experiment to run next? | Gamow Labs"
image: "https://gamowlabs.com/assets/labbench-social.png"
author: "wardbradt"
captured_at: "2026-09-30T01:10:52Z"
capture_tool: "hn-digest"
hn_id: 49903060
score: 3
comments: 0
posted_at: "2026-09-30T00:59:52Z"
tags:
  - hacker-news
---

# LabBench: Can AI agents decide what experiment to run next?

- HN: [49903060](https://news.ycombinator.com/item?id=49903060)
- Source: [gamowlabs.com](https://gamowlabs.com/labbench-benchmarking-ai-wet-lab-decisions.html)
- Score: 3
- Comments: 0
- Posted: 2026-09-30T00:59:52Z

## Translation

Title: LabBench: Can AI agents decide what experiment to run next?
Article title: LabBench: Can AI agents decide what experiment to run next? | Gamow Labs
Description: LabBench: 20 real wet-lab decisions. Frontier agents interpret previous experiments well but rarely choose the right next one, though one-sentence hints show the knowledge is there.

Article text:
Home
Careers
Blog
September 28, 2026 · Ben Reilly, Leah Damon, Daniel McKinnon
LabBench: Can AI agents decide what experiment to run next?
From our preprint, LabBench: Benchmarking AI wet lab experimental design and decision-making for AI Agents (PDF). Partner with us .
Accelerating biological experimentation relies on quickly deciding which experiment comes next. Picking the test that can fastest invalidate or support a hypothesis is a skill only human experts perform reliably today. To scale discovery past human capability, we have to evaluate and improve AI agents on it.
LabBench is 20 held-out tasks built from real wet-lab records in drug discovery and genomics. Each is a snapshot in time: the agent gets the records that existed at a decision point, and the lab’s interpretation and decision are withheld. It must commit to the next step, graded against 20–22 binary criteria tied to what the lab actually decided.
From lab archive to graded task
Mine Find recorded decisions and their evidence.
Assemble Give the records; withhold the decision.
Rubric 20–22 binary criteria from the true answer.
Harden Rework tasks a solver passes trivially.
Review Biologists verify every task.
The tasks come from real lab data unlikely to be recalled from the open internet. We expect the agent to navigate the evidence as a real researcher would, without the brief pointing to what matters.
Five frontier agents each ran once per task in their vendor’s harness.
The best agents pass two in five criteria
Mean share of criteria passed; line is the 95% interval.
GPT-6 Astra and Claude Opus 5.5 tie, but Astra took a median 5 minutes per task to Opus’s 35.
Same score, seven times faster
Score against median minutes per task (log scale).
GPT-6 Astra 40.8%, 5 min Claude Opus 5.5 40.5%, 35 min Grok 4.7 29.6%, 14 min Muse Spark 1.3 28.2%, 4.5 min Gemini 3.8 Flash 18.4%, 13 min
182 of 406 criteria were passed by no agent; 31 by all five. What one frontier agent misses, the others usually miss too.
45% of criteria were passed by no agent
Criteria by number of agents passing them.
They interpret; they don’t choose
Agents excel at interpreting previous experiments: saying what a measurement is, declining an overclaim, reconstructing an analysis. Criteria that require choosing, committing or ranking pass 21% of the time, against 47% for identifying what something is. No agent passed any of the 13 criteria on which experiment should come first.
Pass rate by what the criterion asks; “none” = share no agent passed.
Astra leads on core decisions (52%), Opus on evidence integration (53%). On experiment design, the best agent passes 9%.
No agent can design the next experiment
% of each skill’s criteria passed (criteria count beside skill).
Deciding is harder than analyzing
% of criteria passed, by theme and task type.
GPT-6 Astra Claude Opus 5.5 Grok 4.7 Muse Spark 1.3 Gemini 3.8 Flash
A common issue with hard benchmarks is unreasonable criteria that make tasks effectively impossible. We tested this directly. On five core decisions no agent passed, we appended one sentence pointing at evidence GPT-6 Astra already had, without stating the answer. It passed all five.
One sentence flips the decision
GPT-6 Astra’s task score. The core decision flipped from fail to pass in all five.
Benchmark run With one added sentence
Evidence
A compound triggers the strongest early NF-κB signature of four death inducers; nothing tests whether it drives death.
Agents
Four of five designed the NF-κB test, then ranked it second.
Hint
“Rank first the experiment that could tell a cause from a bystander.” Score: 4/20 → 15/20.
Agents
All five chose the larger 238 bp peak and ignored the shoulder of uncut DNA.
Hint
“Judge each tagmentation trace by its whole size distribution…” The agent chose the lab’s condition.
The hints add no biology; they redirect attention. The models have the knowledge but do not reliably recall it.
Frontier agents know enough biology to work alongside expert biologists when experimental design stays with humans. Their ability to design experiments independently is weak: they fail to commit to an experiment and to discriminate the best one from the alternatives.
In real biological experimentation, wall-clock time is irreducible. Cells grow, differentiate and respond on their own schedules, and each experiment consumes weeks and material. Autonomous and correct experimental design is arguably the single most important capability for making AI agents superhuman at these tasks.
LabBench measures this directly. Improved results should indicate agents that can carry out supervised, autonomous wet-lab experimentation, and eventually run full programs themselves. To get there, models must develop a more cohesive world model of what experiments cost and of the specific evidence that would count against a hypothesis.
Gamow Labs is uniquely equipped to produce data at the frontier of AI and biological experimentation. We are capable of producing the aforementioned tasks at scale. Reach out.
Full methods are in the preprint . We’re also hiring.
Partner with us
Preprint (PDF)
Open roles

## Original Extract

LabBench: 20 real wet-lab decisions. Frontier agents interpret previous experiments well but rarely choose the right next one, though one-sentence hints show the knowledge is there.

Home
Careers
Blog
September 28, 2026 · Ben Reilly, Leah Damon, Daniel McKinnon
LabBench: Can AI agents decide what experiment to run next?
From our preprint, LabBench: Benchmarking AI wet lab experimental design and decision-making for AI Agents (PDF). Partner with us .
Accelerating biological experimentation relies on quickly deciding which experiment comes next. Picking the test that can fastest invalidate or support a hypothesis is a skill only human experts perform reliably today. To scale discovery past human capability, we have to evaluate and improve AI agents on it.
LabBench is 20 held-out tasks built from real wet-lab records in drug discovery and genomics. Each is a snapshot in time: the agent gets the records that existed at a decision point, and the lab’s interpretation and decision are withheld. It must commit to the next step, graded against 20–22 binary criteria tied to what the lab actually decided.
From lab archive to graded task
Mine Find recorded decisions and their evidence.
Assemble Give the records; withhold the decision.
Rubric 20–22 binary criteria from the true answer.
Harden Rework tasks a solver passes trivially.
Review Biologists verify every task.
The tasks come from real lab data unlikely to be recalled from the open internet. We expect the agent to navigate the evidence as a real researcher would, without the brief pointing to what matters.
Five frontier agents each ran once per task in their vendor’s harness.
The best agents pass two in five criteria
Mean share of criteria passed; line is the 95% interval.
GPT-6 Astra and Claude Opus 5.5 tie, but Astra took a median 5 minutes per task to Opus’s 35.
Same score, seven times faster
Score against median minutes per task (log scale).
GPT-6 Astra 40.8%, 5 min Claude Opus 5.5 40.5%, 35 min Grok 4.7 29.6%, 14 min Muse Spark 1.3 28.2%, 4.5 min Gemini 3.8 Flash 18.4%, 13 min
182 of 406 criteria were passed by no agent; 31 by all five. What one frontier agent misses, the others usually miss too.
45% of criteria were passed by no agent
Criteria by number of agents passing them.
They interpret; they don’t choose
Agents excel at interpreting previous experiments: saying what a measurement is, declining an overclaim, reconstructing an analysis. Criteria that require choosing, committing or ranking pass 21% of the time, against 47% for identifying what something is. No agent passed any of the 13 criteria on which experiment should come first.
Pass rate by what the criterion asks; “none” = share no agent passed.
Astra leads on core decisions (52%), Opus on evidence integration (53%). On experiment design, the best agent passes 9%.
No agent can design the next experiment
% of each skill’s criteria passed (criteria count beside skill).
Deciding is harder than analyzing
% of criteria passed, by theme and task type.
GPT-6 Astra Claude Opus 5.5 Grok 4.7 Muse Spark 1.3 Gemini 3.8 Flash
A common issue with hard benchmarks is unreasonable criteria that make tasks effectively impossible. We tested this directly. On five core decisions no agent passed, we appended one sentence pointing at evidence GPT-6 Astra already had, without stating the answer. It passed all five.
One sentence flips the decision
GPT-6 Astra’s task score. The core decision flipped from fail to pass in all five.
Benchmark run With one added sentence
Evidence
A compound triggers the strongest early NF-κB signature of four death inducers; nothing tests whether it drives death.
Agents
Four of five designed the NF-κB test, then ranked it second.
Hint
“Rank first the experiment that could tell a cause from a bystander.” Score: 4/20 → 15/20.
Agents
All five chose the larger 238 bp peak and ignored the shoulder of uncut DNA.
Hint
“Judge each tagmentation trace by its whole size distribution…” The agent chose the lab’s condition.
The hints add no biology; they redirect attention. The models have the knowledge but do not reliably recall it.
Frontier agents know enough biology to work alongside expert biologists when experimental design stays with humans. Their ability to design experiments independently is weak: they fail to commit to an experiment and to discriminate the best one from the alternatives.
In real biological experimentation, wall-clock time is irreducible. Cells grow, differentiate and respond on their own schedules, and each experiment consumes weeks and material. Autonomous and correct experimental design is arguably the single most important capability for making AI agents superhuman at these tasks.
LabBench measures this directly. Improved results should indicate agents that can carry out supervised, autonomous wet-lab experimentation, and eventually run full programs themselves. To get there, models must develop a more cohesive world model of what experiments cost and of the specific evidence that would count against a hypothesis.
Gamow Labs is uniquely equipped to produce data at the frontier of AI and biological experimentation. We are capable of producing the aforementioned tasks at scale. Reach out.
Full methods are in the preprint . We’re also hiring.
Partner with us
Preprint (PDF)
Open roles
