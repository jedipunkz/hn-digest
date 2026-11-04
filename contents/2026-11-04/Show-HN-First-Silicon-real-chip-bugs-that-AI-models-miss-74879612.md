---
source: "https://umwelt.vision/first-silicon/"
hn_url: "https://news.ycombinator.com/item?id=50368283"
title: "Show HN: First Silicon, real chip bugs that AI models miss"
article_title: "First Silicon: a benchmark of real hardware bugs for AI models"
image: "https://umwelt.vision/img/og-first-silicon.jpg"
author: "Arno_Bighub"
captured_at: "2026-11-04T16:54:07Z"
capture_tool: "hn-digest"
hn_id: 50368283
score: 1
comments: 0
posted_at: "2026-11-04T16:50:41Z"
tags:
  - hacker-news
---

# Show HN: First Silicon, real chip bugs that AI models miss

- HN: [50368283](https://news.ycombinator.com/item?id=50368283)
- Source: [umwelt.vision](https://umwelt.vision/first-silicon/)
- Score: 1
- Comments: 0
- Posted: 2026-11-04T16:50:41Z

## Translation

Title: Show HN: First Silicon, real chip bugs that AI models miss
Article title: First Silicon: a benchmark of real hardware bugs for AI models
Description: First Silicon: can frontier models find real hardware bugs? 86 fixes made by engineers in open-source chips, 84 of them real bugs, judged by an exact formal checker.

Article text:
First Silicon: a benchmark of real hardware bugs for AI models
Umwelt
About
Research
Careers
Get in touch
Research · October 2026
First Silicon
Can frontier models find real hardware bugs without raising false alarms? We collected 86 fixes that engineers made in open-source chip designs, 84 of them real bugs, hid which version each model sees, and let a formal checker decide.
The best configuration, Claude Opus 5.5 at maximum effort, catches 16 of the 86 cases cleanly, and two of its alarms in three are false. Most bugs escape every model: 67 of the 84 genuine fixes are caught cleanly by none of the four models at default settings.
With 86 cases, the benchmark separates large models from small ones (Opus 5.5 against Haiku 5.5: 10 bugs against 1, p = 0.012 in an exact McNemar test), but not the strongest from each other (p ≥ 0.23).
Real bugs. 70 come from OpenTitan and 16 from 10 other projects. Each fix provably changes behaviour at the outputs, as certified by a formal miter; a trivial template (“outputs are zero after reset”) does not catch it; and none was used to train our own models. Two later turned out to be intentional behaviour changes, which we keep and flag.
A blind task. The model reads a documentation excerpt, the port and signal lists, and one version of the module, buggy or fixed, without being told which. It writes 3 to 8 SystemVerilog assertions describing the intended behaviour. File names are neutral, and the two versions of a bug never share a batch.
An exact judge. Yosys and SymbiYosys run every assertion on both versions, with bounded model checking and then a PDR proof. A failure on the buggy version is a catch only if the same assertion is proven on the fixed one; any failure on the fixed version is a false alarm. Under STRICT, a single false alarm cancels the catch.
The name comes from the first chips back from the fab, where a missed bug costs a new mask set and months of delay.
We let Opus 5.5 call the judge’s bounded model checker up to five times, on the shown version only. The method was designed on 20 bugs and confirmed once on 61 others.
The checker did not close the gap. When an assertion received a counterexample, the model dropped or changed it 79 % of the time, including 8 assertions that had correctly exposed a bug: it found the bug, the tool confirmed it, and the model decided its own assertion was wrong.
The triage keeps a counterexample only when two independent reviewers quote, word for word, the sentence of the documentation it contradicts. It is more precise than random removal and stable across passes, but it did not meet our pre-registered criteria. The open problem is judging a counterexample: is the code wrong, or the assertion?
Squares, left to right: Opus 5.5 max, Opus 5.5, Fable 5.1, Sonnet 5.5, Haiku 5.5, with the checker, with checker and triage.
The designs are public, and some answers refer to remembered upstream code, which can only inflate scores; only 3 bugs were fixed after the models’ training cutoff, and none is caught. Six prompts had an enum table that did not match the source; without them, STRICT is 14, 12, 8, 7 and 3 out of 80. Each prompt was run once, the judge works at module level with a restricted assertion subset, and the benchmark has not yet been externally validated.
We publish the outcome of every bug, five bugs in full, and a salted SHA-256 commitment for each bug and prompt; the salt is revealed on audit. The other 81 bugs stay private, so the benchmark stays usable for future models. We can evaluate yours.
Want your model evaluated on First Silicon, or your own design checked? Write to us below, or at contact@umwelt.vision .
Want your model on First Silicon, or your design checked? Write to us.

## Original Extract

First Silicon: can frontier models find real hardware bugs? 86 fixes made by engineers in open-source chips, 84 of them real bugs, judged by an exact formal checker.

First Silicon: a benchmark of real hardware bugs for AI models
Umwelt
About
Research
Careers
Get in touch
Research · October 2026
First Silicon
Can frontier models find real hardware bugs without raising false alarms? We collected 86 fixes that engineers made in open-source chip designs, 84 of them real bugs, hid which version each model sees, and let a formal checker decide.
The best configuration, Claude Opus 5.5 at maximum effort, catches 16 of the 86 cases cleanly, and two of its alarms in three are false. Most bugs escape every model: 67 of the 84 genuine fixes are caught cleanly by none of the four models at default settings.
With 86 cases, the benchmark separates large models from small ones (Opus 5.5 against Haiku 5.5: 10 bugs against 1, p = 0.012 in an exact McNemar test), but not the strongest from each other (p ≥ 0.23).
Real bugs. 70 come from OpenTitan and 16 from 10 other projects. Each fix provably changes behaviour at the outputs, as certified by a formal miter; a trivial template (“outputs are zero after reset”) does not catch it; and none was used to train our own models. Two later turned out to be intentional behaviour changes, which we keep and flag.
A blind task. The model reads a documentation excerpt, the port and signal lists, and one version of the module, buggy or fixed, without being told which. It writes 3 to 8 SystemVerilog assertions describing the intended behaviour. File names are neutral, and the two versions of a bug never share a batch.
An exact judge. Yosys and SymbiYosys run every assertion on both versions, with bounded model checking and then a PDR proof. A failure on the buggy version is a catch only if the same assertion is proven on the fixed one; any failure on the fixed version is a false alarm. Under STRICT, a single false alarm cancels the catch.
The name comes from the first chips back from the fab, where a missed bug costs a new mask set and months of delay.
We let Opus 5.5 call the judge’s bounded model checker up to five times, on the shown version only. The method was designed on 20 bugs and confirmed once on 61 others.
The checker did not close the gap. When an assertion received a counterexample, the model dropped or changed it 79 % of the time, including 8 assertions that had correctly exposed a bug: it found the bug, the tool confirmed it, and the model decided its own assertion was wrong.
The triage keeps a counterexample only when two independent reviewers quote, word for word, the sentence of the documentation it contradicts. It is more precise than random removal and stable across passes, but it did not meet our pre-registered criteria. The open problem is judging a counterexample: is the code wrong, or the assertion?
Squares, left to right: Opus 5.5 max, Opus 5.5, Fable 5.1, Sonnet 5.5, Haiku 5.5, with the checker, with checker and triage.
The designs are public, and some answers refer to remembered upstream code, which can only inflate scores; only 3 bugs were fixed after the models’ training cutoff, and none is caught. Six prompts had an enum table that did not match the source; without them, STRICT is 14, 12, 8, 7 and 3 out of 80. Each prompt was run once, the judge works at module level with a restricted assertion subset, and the benchmark has not yet been externally validated.
We publish the outcome of every bug, five bugs in full, and a salted SHA-256 commitment for each bug and prompt; the salt is revealed on audit. The other 81 bugs stay private, so the benchmark stays usable for future models. We can evaluate yours.
Want your model evaluated on First Silicon, or your own design checked? Write to us below, or at contact@umwelt.vision .
Want your model on First Silicon, or your design checked? Write to us.
