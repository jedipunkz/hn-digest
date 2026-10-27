---
source: "https://www.heal.dev/benchmark"
hn_url: "https://news.ycombinator.com/item?id=50258722"
title: "Static reasoning can't prove runtime behavior: benchmark of AI PR reviewers"
article_title: "Benchmark: Static reasoning can’t prove runtime behavior: a benchmark of what AI PR reviewers miss · heal"
image: "https://heal.dev/benchmark/opengraph-image-ccf4qc?c0bb0d55fe6e1b57"
author: "batikha"
captured_at: "2026-10-27T14:02:46Z"
capture_tool: "hn-digest"
hn_id: 50258722
score: 5
comments: 1
posted_at: "2026-10-27T13:14:03Z"
tags:
  - hacker-news
---

# Static reasoning can't prove runtime behavior: benchmark of AI PR reviewers

- HN: [50258722](https://news.ycombinator.com/item?id=50258722)
- Source: [www.heal.dev](https://www.heal.dev/benchmark)
- Score: 5
- Comments: 1
- Posted: 2026-10-27T13:14:03Z

## Translation

Title: Static reasoning can't prove runtime behavior: benchmark of AI PR reviewers
Article title: Benchmark: Static reasoning can’t prove runtime behavior: a benchmark of what AI PR reviewers miss · heal
Description: A benchmark of Heal's full product run against CodeRabbit, Greptile and cubic on 236 pull requests of five products: what each side finds, what the engineer gets, and where Heal goes next.

Article text:
Benchmark: Static reasoning can’t prove runtime behavior: a benchmark of what AI PR reviewers miss · heal New Our benchmark results are out! See how heal finds issues code review doesn't find read it vs code review
try now T Heal · benchmark · October 2026
Benchmark: Static reasoning can’t prove runtime behavior: a benchmark of what AI PR reviewers miss
Keep code reviewers. Don't stack them. Add the runtime.
We built Heal, an agent that finds and traverses product entrypoints like users do. This gives a different class of bugs than code reviewers.
Heal takes in a codebase and learns how to run it in a way that is production-like. It mocks, stubs, or emulates dependencies, and creates an environment in which to run and reproduce. We call that environment the Heal application sandbox .
The majority of bugs could be found by reading code, but the nuance is that Heal can afford to generate a lot of noisy candidates, and having the Heal application sandbox allows for filtering. Code reviewers optimize for generating less noise. We optimize for generating more candidates, then filter them thanks to the Heal application sandbox.
Heal is complementary to code review because it finds only runtime issues reached by user behavior (behavior, performance, security). 70 % of the defects are not found by code-review agents. Code review finds a lot more defects in total, mostly stylistic or in files Heal doesn't cover because they don't affect the runtime.
Sometimes behavior issues are open product questions and not bugs. Our next frontier is improving how this fits in the development loop.
As a separate insight, we find that stacking code reviewers is not really useful as they massively overlap.
of the defects Heal found are not found by code review.
Runtime validation vs code review
Results: what code review bots don't see
Results: what code review bots find and Heal doesn't
The truth about stacking: is a second code review bot the answer?
Heal builds a Heal application sandbox around any application and drives it as its users and their machines do, through every entry point: the UI, the API, the CLI, the SDK, webhooks and scheduled jobs, with the data under control, to find defects and prove each one with a failing test.
Heal agent acts as several users and machines drives observes screens, responses, exit codes, logs, stored data and what left for third parties Heal application sandbox · built from your repositories Entry points UI API CLI SDK webhooks jobs Your product under test · runs real frontends, services, workers, jobs, functions Storage runs real, seeded databases, caches, queues, files, search indexes Harness what Heal adds browser, API client, shell stubs: fault switch, journal counters, write gate Third parties stood in for emulator · vendor test mode mock · captured sink · off each defect is re-run by the system Reproduced defect a failing test, the steps, the cause in the code Figure 1. The Heal agent, the Heal application sandbox it builds around the application, and what it produces. How Heal finds defects
Scan the codebase, and map the features. A feature is an entrypoint into the product, for a human or a machine (UI, CLI, SDK, etc).
Scope the Heal application sandbox based on this map. For example, it creates seeds, or it mocks external APIs, or it creates a controlled environment for the UI. This also gives the agent a deep understanding of how the system works.
Find many defect candidates: Run static analysis on the codebase to identify a large number of defect candidates, powered by all the context that the agent has. This is similar to what code review does.
Run other tools that need a runtime, such as security scanners and UI checkers. Interestingly, Heal could be extended here by running any custom tool against the Heal application sandbox process.
Inject manual QA priors to find common issues, from a rich knowledge base of past defects and testing experience.
Prove candidates against the Heal application sandbox, writing a test to allow for reproduction. Each defect reported to the user comes with a test that doesn't work, and can be actually reached from using the application.
2 . Runtime validation vs code review
Code review and Heal both aim to catch defects before they get deployed to users. Ultimately, both look at automating a part of the agentic verification loop , and making it as fast, cheap, and seamless as possible.
Code review focuses on "what a developer would change" , and addresses things Heal doesn't cover, such as code style, code consistency, and adherence to internal guidelines. Heal focuses on "what users will encounter" , and catches behavior, security, and performance issues that happen at runtime.
The first kind is argument-based: a developer or an agent reads the comment, decides whether it is right, and fixes or dismisses it. The second kind is proof-based: the defect is a behaviour a user would meet, and the test is the evidence. This benchmark measures what each kind finds on the same products.
The goal of the benchmark is to answer three questions:
What is the benefit of having a runtime, a Heal application sandbox?
Is it worth running both code review and Heal? What is the difference?
Is it worth having several code reviewers?
The deeper question is: what would let PRs be merged much faster, with high confidence that there are no bugs?
We chose five products: two private codebases that we own, anonymised here, and three open-source ones, Formbricks , Sim and Infisical .
We replayed each product's real pull-request history, PR by PR, into a fresh private repository with the three review bots installed on their default settings, and ran Heal before and after the sequence of PRs tested.
To tell whether a bot and Heal found the same thing, a judge compares them: a language model that reads a bot's claim and a Heal defect on the same lines and says whether they name the same problem. It was calibrated on a manually curated sample.
This is a relatively small dataset, but it gives a useful signal. We stopped after spending several thousand dollars on code review. We would be interested in running a fuller benchmark: if you are building a code reviewer and would like to collaborate, hit us up on X .
Greptile with TREX Code review, with code snippet runtime validation
every PR replayed; CodeRabbit, Greptile and cubic on each PR of four products, and on 60, 50 and 45 of Product A's 60
We did not assess whether the code review flags were relevant, whereas Heal defects are proven and reproducible with a test. This may create a bias in favor of code reviewers.
Product A's sequence of PRs tested is its whole history, because the project is earlier stage. This is both favorable to code reviewers (everything got scanned) and to Heal (earlier projects likely contain more issues).
What was tried first. Before the replays, the three open-source products were measured on one-week windows of their live pull requests, 25 to 30 each, with the bots as the maintainers run them and a full run at each end of the week. That method did not work well. Correlating a full run's defects with a week of live history is too complex: the state the run saw and the lines the week's PRs changed drift apart, each repository had different bots on different settings, and a defect could seldom be tied to one PR with confidence. Replaying the PR history on a private copy removes all three problems, so those windows were dropped from this page and their data kept as a side check.
We couldn't validate if the code review comments were acted upon by the authors, as the replays were isolated from the real history. The reason we did not use the real history is code reviewers were not installed on the repos at the time, which would have made the data too scarce to action.
4 . Results: what code review bots don't see
Across the five products, 70 % of the user-facing defects that the full run proved on the code of the sequence of PRs tested (141 of 202) had never been found by any review bot.
More candidates, filtered by the Heal application sandbox
A bot hallucinates by nature and needs to manage its noise level. Heal's candidates are allowed to be wrong, then go to the running Heal application sandbox, where it's tested and either confirmed and reproduced or discarded.
Ultimately this leads to extra bugs found without increasing the noise level. Additionally, because Heal filters using the Heal application sandbox, it raises only bugs that impact the user. We'll see in the next section how this is complementary to code review.
Candidates found A 1,170 · B 2,036 Validated on the Heal application sandbox settled A 363 · B 746 in progress A 237 · B 306 Dismissed A 83 · B 312 · nobody reads them Reproduced A 280 · B 434 Defects filed A 151 · B 172, one test each Figure 3. Heal's filtering funnel on Products A and B. Examples: Four defects no code review bot could find
5 . Results: what code review bots find and Heal doesn't
So what do code reviewers find that Heal doesn't? The majority of findings are issues with files that don't impact the runtime (CI, analytics, tests) that Heal doesn't read. Those produce a large number of findings. The rest are mostly race conditions.
We sorted the bots' 1,003 distinct flags into four buckets:
In files Heal does not look at, by design: 654 (309 on Product A, 65 on B, 100 on Formbricks, 94 on Sim, 86 on Infisical). Heal starts from what a user can do and works inward, so it skips code no user ever reaches (CI configuration, database migration history, dependency versions, analytics, upgrade scripts, the product's own tests). That is where code review is the right tool.
On code Heal tests, but not reported by Heal: 202 (168 on Product A, 10 on B, 7 on Formbricks, 9 on Sim, 8 on Infisical). Caveat: we did not check whether these are real bugs. Most are race conditions, which a single pass through a feature does not always trigger; section 8 explains how we plan to close that gap. Some look plausible: A date that shifts by a day in time zones ahead of UTC
Edits typed during a save that get lost
On Sim, a file's access setting read before its database row is locked
Out of scope: 100 (80 on Product A, 18 on B, 1 on Formbricks, 1 on Sim, 0 on Infisical). These sit on code that a later PR rewrote before Heal ran, so Heal never saw it, and it is out of the set.
Examples: What Heal does not find, mostly outside the user path
6 . The truth about stacking: is a second code review bot the answer?
Several teams we talk to run multiple code review bots. On our sample, this gives very marginal improvements.
each bar = the 120 defects a Heal run proved
each bar = the 53 defects a Heal run proved
each bar = the 31 defects a Heal run proved
each bar = the 21 defects a Heal run proved
each bar = the 22 defects a Heal run proved
7 . Heal and code review find complementary issues
Keep the bot. Add the run. They do not find the same things, and the overlap is small.
It is on the diff a minute after the push, before anyone else reads it.
It reads files the run never maps: migrations, CI, scripts, documentation, the product's own tests. On Product A that was 309 distinct claims; on Formbricks, 100.
Nearly every claim, 98 to 99 % of them, comes with a proposed fix.
The defects that reach users and that no code review finds.
Every finding is proven and reproducible with a test script across the Heal application sandbox. This is a total of 1,085 regression tests across the five products.
The map of the product and the memory of every pass, which make the next run find more.
We found it interesting that a Heal defect sometimes questions the intent. Teams using Heal told us that a share of the defects sent them back to the decision behind the code: heal surfaced a bug, but the expected behavior it proposed was wrong, and the honest conclusion was that this part of the product was never really designed. This is an active area of development for us.
Heal exists so that code wr

[truncated]

## Original Extract

A benchmark of Heal's full product run against CodeRabbit, Greptile and cubic on 236 pull requests of five products: what each side finds, what the engineer gets, and where Heal goes next.

Benchmark: Static reasoning can’t prove runtime behavior: a benchmark of what AI PR reviewers miss · heal New Our benchmark results are out! See how heal finds issues code review doesn't find read it vs code review
try now T Heal · benchmark · October 2026
Benchmark: Static reasoning can’t prove runtime behavior: a benchmark of what AI PR reviewers miss
Keep code reviewers. Don't stack them. Add the runtime.
We built Heal, an agent that finds and traverses product entrypoints like users do. This gives a different class of bugs than code reviewers.
Heal takes in a codebase and learns how to run it in a way that is production-like. It mocks, stubs, or emulates dependencies, and creates an environment in which to run and reproduce. We call that environment the Heal application sandbox .
The majority of bugs could be found by reading code, but the nuance is that Heal can afford to generate a lot of noisy candidates, and having the Heal application sandbox allows for filtering. Code reviewers optimize for generating less noise. We optimize for generating more candidates, then filter them thanks to the Heal application sandbox.
Heal is complementary to code review because it finds only runtime issues reached by user behavior (behavior, performance, security). 70 % of the defects are not found by code-review agents. Code review finds a lot more defects in total, mostly stylistic or in files Heal doesn't cover because they don't affect the runtime.
Sometimes behavior issues are open product questions and not bugs. Our next frontier is improving how this fits in the development loop.
As a separate insight, we find that stacking code reviewers is not really useful as they massively overlap.
of the defects Heal found are not found by code review.
Runtime validation vs code review
Results: what code review bots don't see
Results: what code review bots find and Heal doesn't
The truth about stacking: is a second code review bot the answer?
Heal builds a Heal application sandbox around any application and drives it as its users and their machines do, through every entry point: the UI, the API, the CLI, the SDK, webhooks and scheduled jobs, with the data under control, to find defects and prove each one with a failing test.
Heal agent acts as several users and machines drives observes screens, responses, exit codes, logs, stored data and what left for third parties Heal application sandbox · built from your repositories Entry points UI API CLI SDK webhooks jobs Your product under test · runs real frontends, services, workers, jobs, functions Storage runs real, seeded databases, caches, queues, files, search indexes Harness what Heal adds browser, API client, shell stubs: fault switch, journal counters, write gate Third parties stood in for emulator · vendor test mode mock · captured sink · off each defect is re-run by the system Reproduced defect a failing test, the steps, the cause in the code Figure 1. The Heal agent, the Heal application sandbox it builds around the application, and what it produces. How Heal finds defects
Scan the codebase, and map the features. A feature is an entrypoint into the product, for a human or a machine (UI, CLI, SDK, etc).
Scope the Heal application sandbox based on this map. For example, it creates seeds, or it mocks external APIs, or it creates a controlled environment for the UI. This also gives the agent a deep understanding of how the system works.
Find many defect candidates: Run static analysis on the codebase to identify a large number of defect candidates, powered by all the context that the agent has. This is similar to what code review does.
Run other tools that need a runtime, such as security scanners and UI checkers. Interestingly, Heal could be extended here by running any custom tool against the Heal application sandbox process.
Inject manual QA priors to find common issues, from a rich knowledge base of past defects and testing experience.
Prove candidates against the Heal application sandbox, writing a test to allow for reproduction. Each defect reported to the user comes with a test that doesn't work, and can be actually reached from using the application.
2 . Runtime validation vs code review
Code review and Heal both aim to catch defects before they get deployed to users. Ultimately, both look at automating a part of the agentic verification loop , and making it as fast, cheap, and seamless as possible.
Code review focuses on "what a developer would change" , and addresses things Heal doesn't cover, such as code style, code consistency, and adherence to internal guidelines. Heal focuses on "what users will encounter" , and catches behavior, security, and performance issues that happen at runtime.
The first kind is argument-based: a developer or an agent reads the comment, decides whether it is right, and fixes or dismisses it. The second kind is proof-based: the defect is a behaviour a user would meet, and the test is the evidence. This benchmark measures what each kind finds on the same products.
The goal of the benchmark is to answer three questions:
What is the benefit of having a runtime, a Heal application sandbox?
Is it worth running both code review and Heal? What is the difference?
Is it worth having several code reviewers?
The deeper question is: what would let PRs be merged much faster, with high confidence that there are no bugs?
We chose five products: two private codebases that we own, anonymised here, and three open-source ones, Formbricks , Sim and Infisical .
We replayed each product's real pull-request history, PR by PR, into a fresh private repository with the three review bots installed on their default settings, and ran Heal before and after the sequence of PRs tested.
To tell whether a bot and Heal found the same thing, a judge compares them: a language model that reads a bot's claim and a Heal defect on the same lines and says whether they name the same problem. It was calibrated on a manually curated sample.
This is a relatively small dataset, but it gives a useful signal. We stopped after spending several thousand dollars on code review. We would be interested in running a fuller benchmark: if you are building a code reviewer and would like to collaborate, hit us up on X .
Greptile with TREX Code review, with code snippet runtime validation
every PR replayed; CodeRabbit, Greptile and cubic on each PR of four products, and on 60, 50 and 45 of Product A's 60
We did not assess whether the code review flags were relevant, whereas Heal defects are proven and reproducible with a test. This may create a bias in favor of code reviewers.
Product A's sequence of PRs tested is its whole history, because the project is earlier stage. This is both favorable to code reviewers (everything got scanned) and to Heal (earlier projects likely contain more issues).
What was tried first. Before the replays, the three open-source products were measured on one-week windows of their live pull requests, 25 to 30 each, with the bots as the maintainers run them and a full run at each end of the week. That method did not work well. Correlating a full run's defects with a week of live history is too complex: the state the run saw and the lines the week's PRs changed drift apart, each repository had different bots on different settings, and a defect could seldom be tied to one PR with confidence. Replaying the PR history on a private copy removes all three problems, so those windows were dropped from this page and their data kept as a side check.
We couldn't validate if the code review comments were acted upon by the authors, as the replays were isolated from the real history. The reason we did not use the real history is code reviewers were not installed on the repos at the time, which would have made the data too scarce to action.
4 . Results: what code review bots don't see
Across the five products, 70 % of the user-facing defects that the full run proved on the code of the sequence of PRs tested (141 of 202) had never been found by any review bot.
More candidates, filtered by the Heal application sandbox
A bot hallucinates by nature and needs to manage its noise level. Heal's candidates are allowed to be wrong, then go to the running Heal application sandbox, where it's tested and either confirmed and reproduced or discarded.
Ultimately this leads to extra bugs found without increasing the noise level. Additionally, because Heal filters using the Heal application sandbox, it raises only bugs that impact the user. We'll see in the next section how this is complementary to code review.
Candidates found A 1,170 · B 2,036 Validated on the Heal application sandbox settled A 363 · B 746 in progress A 237 · B 306 Dismissed A 83 · B 312 · nobody reads them Reproduced A 280 · B 434 Defects filed A 151 · B 172, one test each Figure 3. Heal's filtering funnel on Products A and B. Examples: Four defects no code review bot could find
5 . Results: what code review bots find and Heal doesn't
So what do code reviewers find that Heal doesn't? The majority of findings are issues with files that don't impact the runtime (CI, analytics, tests) that Heal doesn't read. Those produce a large number of findings. The rest are mostly race conditions.
We sorted the bots' 1,003 distinct flags into four buckets:
In files Heal does not look at, by design: 654 (309 on Product A, 65 on B, 100 on Formbricks, 94 on Sim, 86 on Infisical). Heal starts from what a user can do and works inward, so it skips code no user ever reaches (CI configuration, database migration history, dependency versions, analytics, upgrade scripts, the product's own tests). That is where code review is the right tool.
On code Heal tests, but not reported by Heal: 202 (168 on Product A, 10 on B, 7 on Formbricks, 9 on Sim, 8 on Infisical). Caveat: we did not check whether these are real bugs. Most are race conditions, which a single pass through a feature does not always trigger; section 8 explains how we plan to close that gap. Some look plausible: A date that shifts by a day in time zones ahead of UTC
Edits typed during a save that get lost
On Sim, a file's access setting read before its database row is locked
Out of scope: 100 (80 on Product A, 18 on B, 1 on Formbricks, 1 on Sim, 0 on Infisical). These sit on code that a later PR rewrote before Heal ran, so Heal never saw it, and it is out of the set.
Examples: What Heal does not find, mostly outside the user path
6 . The truth about stacking: is a second code review bot the answer?
Several teams we talk to run multiple code review bots. On our sample, this gives very marginal improvements.
each bar = the 120 defects a Heal run proved
each bar = the 53 defects a Heal run proved
each bar = the 31 defects a Heal run proved
each bar = the 21 defects a Heal run proved
each bar = the 22 defects a Heal run proved
7 . Heal and code review find complementary issues
Keep the bot. Add the run. They do not find the same things, and the overlap is small.
It is on the diff a minute after the push, before anyone else reads it.
It reads files the run never maps: migrations, CI, scripts, documentation, the product's own tests. On Product A that was 309 distinct claims; on Formbricks, 100.
Nearly every claim, 98 to 99 % of them, comes with a proposed fix.
The defects that reach users and that no code review finds.
Every finding is proven and reproducible with a test script across the Heal application sandbox. This is a total of 1,085 regression tests across the five products.
The map of the product and the memory of every pass, which make the next run find more.
We found it interesting that a Heal defect sometimes questions the intent. Teams using Heal told us that a share of the defects sent them back to the decision behind the code: heal surfaced a bug, but the expected behavior it proposed was wrong, and the honest conclusion was that this part of the product was never really designed. This is an active area of development for us.
Heal exists so that code wr

[truncated]
