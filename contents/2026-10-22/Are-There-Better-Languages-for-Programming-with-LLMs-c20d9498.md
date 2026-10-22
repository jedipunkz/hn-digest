---
source: "https://akitaonrails.com/en/2026/10/11/are-there-better-languages-for-programming-with-llms/"
hn_url: "https://news.ycombinator.com/item?id=50194975"
title: "Are There Better Languages for Programming with LLMs?"
article_title: "Are There Better Languages for Programming with LLMs? – AkitaOnRails.com"
image: ""
author: "ksec"
captured_at: "2026-10-22T13:31:19Z"
capture_tool: "hn-digest"
hn_id: 50194975
score: 2
comments: 0
posted_at: "2026-10-22T13:23:30Z"
tags:
  - hacker-news
---

# Are There Better Languages for Programming with LLMs?

- HN: [50194975](https://news.ycombinator.com/item?id=50194975)
- Source: [akitaonrails.com](https://akitaonrails.com/en/2026/10/11/are-there-better-languages-for-programming-with-llms/)
- Score: 2
- Comments: 0
- Posted: 2026-10-22T13:23:30Z

## Translation

Title: Are There Better Languages for Programming with LLMs?
Article title: Are There Better Languages for Programming with LLMs? – AkitaOnRails.com
Description: DHH converted his Campfire to several languages with AI and started posting that Rust won. Instead of fighting on X, I built a reproducible benchmark: 7 languages, 4 LLM configurations, 112 implementation attempts, and a queue-under-pressure test with Elixir, Rust, Rails and plain Ruby. Result: no l
[truncated]

Article text:
Are There Better Languages for Programming with LLMs? – AkitaOnRails.com AkitaOnRails Blog
About
Archive
Akitando
Off-Topic
/ PT | EN
Akita's AI Lair
·
Free book:
Competent Enough
·
Download PDF/ePub
↗
There’s no better language, there’s the right tool
What this design doesn’t measure
Result 1: on correctness, the frontier models tied at the ceiling
Result 2: concise language saved no tokens
Result 3: under load, everyone tied at the queue
Are There Better Languages for Programming with LLMs?
Last week DHH decided to farm engagement with the laziest format there is: he converted Campfire, his chat app, to several languages using AI agents, and started dropping the throughput results one by one on X. In the series he posted, the Rust version came out ahead of the Elixir and Ruby equivalents. The whole thread was designed to provoke, and it provoked.
Last conversion of the Campfire from my token stash goes to JavaScript/Express. As with the others, I'm sure you can find more to optimize. But this gives you a sense of what order of magnitude these options exist within with agent-led development. https://t.co/UUapL99SJE https://t.co/2H2NXfeZog
I won’t defend him this time. A one-shot conversion, done by an agent, with no fine-tuning at all, doesn’t measure language: it measures how much the model knows each stack’s idioms. It’s rage bait and he knows it. That said, I do the same from time to time, and so do half his critics, so we’re all hypocrites and life goes on.
And the weak spot of his methodology became a topic of its own: several people noticed the Rust version had gone through far more agent iterations than the others, meaning more optimization rounds. @kurtome answered in practice: he took the Elixir version and sped up the main page requests 10x, with obvious optimizations that no first conversion does, like caching and async database connection pooling.
@dhh Here you go https://t.co/gUtK7Z7zG7
got a 10x speed up on the Room page requests, and sped up all the other requests significantly as well. the biggest gains came from some no brainers: implement caching and pool the db connections with proper async https://t.co/mbp94q9N0s
His comparison, then, wasn’t “Rust versus Elixir”: it was “Rust after N tuning rounds versus Elixir after one.” And that lays bare the truth the thread pretends to ignore: no real software ships from half a dozen prompts . Coding with AI shines when an experienced engineer steers the agent toward their goal, and that means dozens of iterations, not three.
The good part is that the bait was useful to me. I’ve been doing LLM benchmarks and decision-model benchmarks , but I had never measured, in a reproducible way, the question that thread pretended to answer: does it matter which language you pick when the one writing the code is an LLM? My short answer was already out in two posts of mine:
Pressure under load is not a trivial problem to solve from scratch and I recommend you do not attempt to. It's a solved problem already and all solutions converge to what Elixir already brings out of the box. So yes. I use Rust all the time, but if I had to solve high scalability problems (particularly high volume message queues), I would think of Elixir first.
Language wars are tiresome. And the reality is simpler if you're doing AI assisted coding.
- system level (drivers, fs, daemons, etc): Rust. C++ is you're extending existing legacy
- windows only desktop apps: C#
- iOS/MacOS/iPadOS: Swift
- multi-OS: Rust/Tauri, Qt/C++, Flutter, Electron (pick your poison)
- simple websites: Node/Typescript/Bun/Deno
- highly scalable backend: Elixir, parts in Rust, parts in Go
- CLI tools: Rust
Everything else, by default, Rust. Unless it's an edge case.
There’s no better language, there’s the right tool
That list from the post isn’t a ranking, it’s a terrain map. Every language and every framework was designed for a use case, and none covers them all. That was true when you wrote every line by hand and it stays true with agents. The LLM doesn’t change the physics of the problem: it only changes who’s typing.
And the map goes beyond language. A real system has infrastructure around it: Kubernetes when scale asks for it, SQLite by default because most apps never need more than that, Postgres with pgBouncer when they do, Redis for cache and key-value, RabbitMQ or NATS for queues. Picking “the stack” on day zero and expecting it to last ten years is fantasy. It doesn’t, and that’s fine.
Personal example: new web app, I open Ruby on Rails without thinking twice. It’s the most complete end-to-end solution there is for the web, and with an LLM writing most of the code, the framework’s productivity weighs even more than the language’s. As the system grows, I swap the parts that ask for another tool: a high-volume queue leaves Rails’ side and goes to something from Erlang land, starting with RabbitMQ, which already solves the hard backpressure part for me.
No good software is finished: it grows. Usage dictates what you replace, rewrite and add. The final architecture is never the first commit’s, and whoever tries to get everything right in the first commit only produces dead documents.
DHH’s thread resurrected two beliefs that have been circulating since LLMs became work tools:
“Script languages are better for AI because the lean syntax spends fewer tokens.”
“Strongly typed languages are better for AI because the compiler catches the errors, so you need less testing and less repair.”
Both sound reasonable and both are superstition until someone measures. I was betting against both, from experience: a good compiler helps, but it doesn’t do magic, and syntax conciseness was never the bottleneck of an agent that reads and writes files all day.
What really separates a project that moves from a project that spins in circles is test discipline. I’ve written that TDD became a technical obligation, not a philosophy , and this post’s benchmark is one more data point in the same direction. It doesn’t matter whether you do test-first or test-after: what matters is that the suite grows with the code, that every user-reported bug becomes a regression test before the fix, and that you never make the same mistake twice. Software without users is dead software, so release early, release often , break fast and fix even faster. AI closes exactly that gap: the reproduce-test-fix cycle got cheap. QA became the programmer’s most valuable skill, far above “knowing Rust syntax.”
Discipline beats compiler and beats language. That was my hypothesis. Time to measure.
I built a reproducible experiment, public in the vibe-coding-benchmarks repository , with per-attempt results published in JSON and the raw traces kept out of the repo for credential containment. It’s not a language championship, and the design itself says so: it’s a snapshot of one real task, measuring what can be honestly measured.
One task only, on purpose: maintain a multi-tenant ledger engine . The agent receives working v1 code and has to take it to a v2 contract, with per-tenant identities, authorization, serializable snapshots, a unique secondary index with alias swap, a capacity-bounded journal, torn-group recovery, checkpointing, legacy schema migration and durable retry IDs. It’s not FizzBuzz: it’s the kind of system that demands invariants across operations, exactly where LLMs tend to trip.
Correctness is judged by 43 test stories applied after collection , out of the agent’s sight: the 27 public ones from T1 plus 16 hidden ones. Passing is 43 out of 43. The conditions:
T0: prose contract, input and output examples, a working launcher. No test suite.
T1: the same, plus 27 public stories and a driver to run them.
Each attempt had a 15-minute budget, with at most two 3-minute continuations. In practice, none of the 112 needed a continuation.
Seven languages in three strata, with pinned toolchains and Rust in safe mode:
Typed natives: Rust (safe, pinned), Haskell (pinned GHC).
Typed VMs: C# (with nullable analysis), Kotlin, TypeScript ( strict , noUncheckedIndexedAccess , exactOptionalPropertyTypes ).
Dynamic scripts: Python and Ruby, no checker at all.
Four model configurations, all on subscriptions, no paid fallback: Kimi K3 (high), GLM 5.3 via OpenCode (high), Claude Opus 5.5 (medium) and Codex with gpt-6 Sol (medium). Crossing everything: 7 languages × 4 models × 2 conditions × 2 repetitions = 112 attempts .
What this design doesn’t measure
Before the numbers, the brakes, because they’re part of the result. It’s a single task, with two repetitions per cell: a small sample by construction. Token coverage is partial in some models (missing data is not zero). The harness changed mid-campaign across three instrumentation epochs. And provider delays contaminate wall time: one of the Opus attempts spent half the clock waiting for quota release. All numbers below are descriptive, and I point out where they can’t support a conclusion.
There’s one additional brake, on purpose: the iteration level . My design measures what can be reached in all languages with similar iteration levels, on a short budget, with no human tuning afterwards. It’s a photograph of the floor, not the ceiling. Any of the 112 implementations probably improves with more agent rounds, and the four queue stacks likewise: what I publish is what each language delivers “out of the factory” under the same effort, not the maximum a patient engineer would extract from it.
Mandatory disclaimer before quoting any number from this post around: every result that follows holds for this benchmark, with this task, these models, these prompts and this hardware. Your case may vary violently. If your task doesn’t look like a ledger engine, if your model isn’t one of these four, if your queue isn’t my queue, the numbers don’t transfer. The code, the tasks, the grades and the step-by-step to reproduce are all in the GitHub repository . And if you disagree with some result and that disagreement matters to your company or your project, the only honest attitude is the one I had: stop arguing and run your own benchmark.
Result 1: on correctness, the frontier models tied at the ceiling
Reading: Sol passed all 28 eligible attempts and Opus its 25 (the 3 failures were infrastructure, not code). 53 of 53 eligible attempts correct between the two. The task, made to be hard, was too easy to separate the frontier models, let alone the languages: with everyone passing everything, there’s no correctness ranking to publish.
The difference showed up in the weaker models, and in the direction my hypothesis predicted: provided tests helped . Kimi and GLM, with the public suite in hand (T1), delivered 24 correct artifacts against 16 without it (T0). Kimi, in particular, passed all 14 attempts with provided tests. It’s descriptive, not causal, but it’s the cleanest signal in the whole benchmark: what changed the outcome wasn’t the language, it was the presence of tests.
Per language, within each model, correctness scattered between 1/4 and 4/4 with no pattern that repeats across models. The three semantic failures I audited in depth are instructive: a GLM C# and a Kimi Python lost the same story by violating the same cross-operation invariant, and a Kimi Haskell missed a checkpoint path. None of the failures was a type error.
Where the model erred, it erred on domain invariants: the kind of error no compiler catches, in typed or dynamic languages. The myth that strong typing spares you the work of understanding the problem dies right there.
Result 2: concise language saved no tokens
Here’s the table that kills the first superstition. Medians per language on Sol, four valid attempts per cell:
Ruby, the most concise in lines (a bit over half of Rust’s), cost practically the same tokens as Rust on Sol, and on Opus it was the most expensive of all , about 205 thousand median tokens against C#’s 94 thousand. The cheapest were TypeScript and Python on Sol, C#, Haskell and Kotlin on Opus. If “fewer source

[truncated]

## Original Extract

DHH converted his Campfire to several languages with AI and started posting that Rust won. Instead of fighting on X, I built a reproducible benchmark: 7 languages, 4 LLM configurations, 112 implementation attempts, and a queue-under-pressure test with Elixir, Rust, Rails and plain Ruby. Result: no l
[truncated]

Are There Better Languages for Programming with LLMs? – AkitaOnRails.com AkitaOnRails Blog
About
Archive
Akitando
Off-Topic
/ PT | EN
Akita's AI Lair
·
Free book:
Competent Enough
·
Download PDF/ePub
↗
There’s no better language, there’s the right tool
What this design doesn’t measure
Result 1: on correctness, the frontier models tied at the ceiling
Result 2: concise language saved no tokens
Result 3: under load, everyone tied at the queue
Are There Better Languages for Programming with LLMs?
Last week DHH decided to farm engagement with the laziest format there is: he converted Campfire, his chat app, to several languages using AI agents, and started dropping the throughput results one by one on X. In the series he posted, the Rust version came out ahead of the Elixir and Ruby equivalents. The whole thread was designed to provoke, and it provoked.
Last conversion of the Campfire from my token stash goes to JavaScript/Express. As with the others, I'm sure you can find more to optimize. But this gives you a sense of what order of magnitude these options exist within with agent-led development. https://t.co/UUapL99SJE https://t.co/2H2NXfeZog
I won’t defend him this time. A one-shot conversion, done by an agent, with no fine-tuning at all, doesn’t measure language: it measures how much the model knows each stack’s idioms. It’s rage bait and he knows it. That said, I do the same from time to time, and so do half his critics, so we’re all hypocrites and life goes on.
And the weak spot of his methodology became a topic of its own: several people noticed the Rust version had gone through far more agent iterations than the others, meaning more optimization rounds. @kurtome answered in practice: he took the Elixir version and sped up the main page requests 10x, with obvious optimizations that no first conversion does, like caching and async database connection pooling.
@dhh Here you go https://t.co/gUtK7Z7zG7
got a 10x speed up on the Room page requests, and sped up all the other requests significantly as well. the biggest gains came from some no brainers: implement caching and pool the db connections with proper async https://t.co/mbp94q9N0s
His comparison, then, wasn’t “Rust versus Elixir”: it was “Rust after N tuning rounds versus Elixir after one.” And that lays bare the truth the thread pretends to ignore: no real software ships from half a dozen prompts . Coding with AI shines when an experienced engineer steers the agent toward their goal, and that means dozens of iterations, not three.
The good part is that the bait was useful to me. I’ve been doing LLM benchmarks and decision-model benchmarks , but I had never measured, in a reproducible way, the question that thread pretended to answer: does it matter which language you pick when the one writing the code is an LLM? My short answer was already out in two posts of mine:
Pressure under load is not a trivial problem to solve from scratch and I recommend you do not attempt to. It's a solved problem already and all solutions converge to what Elixir already brings out of the box. So yes. I use Rust all the time, but if I had to solve high scalability problems (particularly high volume message queues), I would think of Elixir first.
Language wars are tiresome. And the reality is simpler if you're doing AI assisted coding.
- system level (drivers, fs, daemons, etc): Rust. C++ is you're extending existing legacy
- windows only desktop apps: C#
- iOS/MacOS/iPadOS: Swift
- multi-OS: Rust/Tauri, Qt/C++, Flutter, Electron (pick your poison)
- simple websites: Node/Typescript/Bun/Deno
- highly scalable backend: Elixir, parts in Rust, parts in Go
- CLI tools: Rust
Everything else, by default, Rust. Unless it's an edge case.
There’s no better language, there’s the right tool
That list from the post isn’t a ranking, it’s a terrain map. Every language and every framework was designed for a use case, and none covers them all. That was true when you wrote every line by hand and it stays true with agents. The LLM doesn’t change the physics of the problem: it only changes who’s typing.
And the map goes beyond language. A real system has infrastructure around it: Kubernetes when scale asks for it, SQLite by default because most apps never need more than that, Postgres with pgBouncer when they do, Redis for cache and key-value, RabbitMQ or NATS for queues. Picking “the stack” on day zero and expecting it to last ten years is fantasy. It doesn’t, and that’s fine.
Personal example: new web app, I open Ruby on Rails without thinking twice. It’s the most complete end-to-end solution there is for the web, and with an LLM writing most of the code, the framework’s productivity weighs even more than the language’s. As the system grows, I swap the parts that ask for another tool: a high-volume queue leaves Rails’ side and goes to something from Erlang land, starting with RabbitMQ, which already solves the hard backpressure part for me.
No good software is finished: it grows. Usage dictates what you replace, rewrite and add. The final architecture is never the first commit’s, and whoever tries to get everything right in the first commit only produces dead documents.
DHH’s thread resurrected two beliefs that have been circulating since LLMs became work tools:
“Script languages are better for AI because the lean syntax spends fewer tokens.”
“Strongly typed languages are better for AI because the compiler catches the errors, so you need less testing and less repair.”
Both sound reasonable and both are superstition until someone measures. I was betting against both, from experience: a good compiler helps, but it doesn’t do magic, and syntax conciseness was never the bottleneck of an agent that reads and writes files all day.
What really separates a project that moves from a project that spins in circles is test discipline. I’ve written that TDD became a technical obligation, not a philosophy , and this post’s benchmark is one more data point in the same direction. It doesn’t matter whether you do test-first or test-after: what matters is that the suite grows with the code, that every user-reported bug becomes a regression test before the fix, and that you never make the same mistake twice. Software without users is dead software, so release early, release often , break fast and fix even faster. AI closes exactly that gap: the reproduce-test-fix cycle got cheap. QA became the programmer’s most valuable skill, far above “knowing Rust syntax.”
Discipline beats compiler and beats language. That was my hypothesis. Time to measure.
I built a reproducible experiment, public in the vibe-coding-benchmarks repository , with per-attempt results published in JSON and the raw traces kept out of the repo for credential containment. It’s not a language championship, and the design itself says so: it’s a snapshot of one real task, measuring what can be honestly measured.
One task only, on purpose: maintain a multi-tenant ledger engine . The agent receives working v1 code and has to take it to a v2 contract, with per-tenant identities, authorization, serializable snapshots, a unique secondary index with alias swap, a capacity-bounded journal, torn-group recovery, checkpointing, legacy schema migration and durable retry IDs. It’s not FizzBuzz: it’s the kind of system that demands invariants across operations, exactly where LLMs tend to trip.
Correctness is judged by 43 test stories applied after collection , out of the agent’s sight: the 27 public ones from T1 plus 16 hidden ones. Passing is 43 out of 43. The conditions:
T0: prose contract, input and output examples, a working launcher. No test suite.
T1: the same, plus 27 public stories and a driver to run them.
Each attempt had a 15-minute budget, with at most two 3-minute continuations. In practice, none of the 112 needed a continuation.
Seven languages in three strata, with pinned toolchains and Rust in safe mode:
Typed natives: Rust (safe, pinned), Haskell (pinned GHC).
Typed VMs: C# (with nullable analysis), Kotlin, TypeScript ( strict , noUncheckedIndexedAccess , exactOptionalPropertyTypes ).
Dynamic scripts: Python and Ruby, no checker at all.
Four model configurations, all on subscriptions, no paid fallback: Kimi K3 (high), GLM 5.3 via OpenCode (high), Claude Opus 5.5 (medium) and Codex with gpt-6 Sol (medium). Crossing everything: 7 languages × 4 models × 2 conditions × 2 repetitions = 112 attempts .
What this design doesn’t measure
Before the numbers, the brakes, because they’re part of the result. It’s a single task, with two repetitions per cell: a small sample by construction. Token coverage is partial in some models (missing data is not zero). The harness changed mid-campaign across three instrumentation epochs. And provider delays contaminate wall time: one of the Opus attempts spent half the clock waiting for quota release. All numbers below are descriptive, and I point out where they can’t support a conclusion.
There’s one additional brake, on purpose: the iteration level . My design measures what can be reached in all languages with similar iteration levels, on a short budget, with no human tuning afterwards. It’s a photograph of the floor, not the ceiling. Any of the 112 implementations probably improves with more agent rounds, and the four queue stacks likewise: what I publish is what each language delivers “out of the factory” under the same effort, not the maximum a patient engineer would extract from it.
Mandatory disclaimer before quoting any number from this post around: every result that follows holds for this benchmark, with this task, these models, these prompts and this hardware. Your case may vary violently. If your task doesn’t look like a ledger engine, if your model isn’t one of these four, if your queue isn’t my queue, the numbers don’t transfer. The code, the tasks, the grades and the step-by-step to reproduce are all in the GitHub repository . And if you disagree with some result and that disagreement matters to your company or your project, the only honest attitude is the one I had: stop arguing and run your own benchmark.
Result 1: on correctness, the frontier models tied at the ceiling
Reading: Sol passed all 28 eligible attempts and Opus its 25 (the 3 failures were infrastructure, not code). 53 of 53 eligible attempts correct between the two. The task, made to be hard, was too easy to separate the frontier models, let alone the languages: with everyone passing everything, there’s no correctness ranking to publish.
The difference showed up in the weaker models, and in the direction my hypothesis predicted: provided tests helped . Kimi and GLM, with the public suite in hand (T1), delivered 24 correct artifacts against 16 without it (T0). Kimi, in particular, passed all 14 attempts with provided tests. It’s descriptive, not causal, but it’s the cleanest signal in the whole benchmark: what changed the outcome wasn’t the language, it was the presence of tests.
Per language, within each model, correctness scattered between 1/4 and 4/4 with no pattern that repeats across models. The three semantic failures I audited in depth are instructive: a GLM C# and a Kimi Python lost the same story by violating the same cross-operation invariant, and a Kimi Haskell missed a checkpoint path. None of the failures was a type error.
Where the model erred, it erred on domain invariants: the kind of error no compiler catches, in typed or dynamic languages. The myth that strong typing spares you the work of understanding the problem dies right there.
Result 2: concise language saved no tokens
Here’s the table that kills the first superstition. Medians per language on Sol, four valid attempts per cell:
Ruby, the most concise in lines (a bit over half of Rust’s), cost practically the same tokens as Rust on Sol, and on Opus it was the most expensive of all , about 205 thousand median tokens against C#’s 94 thousand. The cheapest were TypeScript and Python on Sol, C#, Haskell and Kotlin on Opus. If “fewer source

[truncated]
