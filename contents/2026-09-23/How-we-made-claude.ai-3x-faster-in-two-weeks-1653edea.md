---
source: "https://claude.dev/blog/how-we-made-claude-ai-faster/"
hn_url: "https://news.ycombinator.com/item?id=49821196"
title: "How we made claude.ai 3x faster in two weeks"
article_title: "How we made claude.ai 3x faster in two weeks / claude.dev"
image: "https://claude.dev/blog/how-we-made-claude-ai-faster/og.png"
author: "matthieu_bl"
captured_at: "2026-09-23T19:29:18Z"
capture_tool: "hn-digest"
hn_id: 49821196
score: 3
comments: 0
posted_at: "2026-09-23T19:23:40Z"
tags:
  - hacker-news
---

# How we made claude.ai 3x faster in two weeks

- HN: [49821196](https://news.ycombinator.com/item?id=49821196)
- Source: [claude.dev](https://claude.dev/blog/how-we-made-claude-ai-faster/)
- Score: 3
- Comments: 0
- Posted: 2026-09-23T19:23:40Z

## Translation

Title: How we made claude.ai 3x faster in two weeks
Article title: How we made claude.ai 3x faster in two weeks / claude.dev
Description: Inside our performance sprint: the benchmarks Claude built, the loop each Slack thread ran, and the guardrails that let us ship 3,000 changes safely.

Article text:
How we made claude.ai 3x faster in two weeks / claude.dev [ H ] HOME [ D ] DOCS [ T ] TERMINAL Try Claude Code [ H ] HOME [ D ] DOCS [ T ] TERMINAL Engineering How we made claude.ai 3x faster in two weeks
Once Claude can measure something, it can make it faster. So we kept finding more things to measure.
This August, we made the core user experience of claude.ai and the Claude desktop app about 3x faster in a two-week sprint. Users had been telling us it was slow, and they were right. We ran everything from a single Slack channel, with Claude in every thread.
We focused on four journeys that make up 95% of user activity. At the 75th percentile, time to a typeable page on a fresh load of claude.ai went from 3.1 seconds to 0.55, starting a new Claude Code session went from 0.8 seconds to 0.3, and loading a Claude Cowork cloud session went from 2.6 seconds to 0.73. In aggregate, we estimate that saves tens of thousands of user-hours of waiting every day.
We used Claude Tag (beta), running an internal research model roughly comparable to Opus 5.5. Claude found bottlenecks, built benchmarks, shipped improvements, and watched every deploy. We steered by setting goals, making tradeoffs, and approving every change. With that approach, we merged more than three thousand changes without a single customer-facing incident or rollback. This post covers what we shipped, how we measured it, and the loop we built with Claude to do it safely.
Before the sprint, we created a Slack channel with the following standing instructions :
@Claude Your job is to facilitate all things related to the performance of the claude.ai website and desktop app. Your responsibilities include monitoring deploys for performance regressions, assessing the accuracy and comprehensiveness of existing telemetry, maintaining well-curated observability dashboards, proactively implementing solutions for observed issues and low-hanging fruit, proposing performance project opportunities, and communicating with your human teammates. […]
The ultimate goal for this channel is for you to become as autonomous as possible, but today we know that isn’t yet possible.
We asked Claude to analyze usage data through the Datadog MCP server. It identified the four highest-impact user journeys: launching the app, starting a conversation, loading an existing conversation, and sending a message. Between web and desktop, and across our products, those journeys came to thirteen distinct measurements. To establish baselines, we added instrumentation until they were directly comparable: each started with a user interaction, ended once the result was rendered, and disambiguated client and server work.
We kicked off the sprint with a list of about twenty hand-picked projects, each targeting a specific journey. Claude estimated the impact of each project in milliseconds, and we aggregated those estimates to set our targets for the sprint. Some of the projects were fairly large, but we thought we could probably achieve most of them within two weeks.
We hit twelve of the thirteen targets by day three.
The planned projects landed early. For faster launches, we baked a static composer into the HTML so users can type during React initialization, and precompiled a V8 code cache so the desktop shell’s main process doesn’t recompile from scratch. For faster navigations, we kept the composer mounted between conversations, prefetched sessions when the user hovered over them, and cut sidebar re-renders by 90%.
We had also left room for Claude to identify opportunities and propose new workstreams. Those workstreams quickly ramped into full projects of their own, which far exceeded our initial targets. So we set new targets, then looked for more things to measure:
@Claude we’ve ended up funding nearly every project in the original projects list and more. let’s do a refresh […] what have we not explored, what can we hill climb on, where is the most opportunity at this point? […] i am open to WACKY ideas
From the start, we knew we wanted to iterate faster than our deploy cadence. Claude could work asynchronously for many hours, even overnight, and we wanted to let it validate its prototypes without waiting for field reads. To achieve that, we looked for other ways to measure performance in the lab.
S Sam 12:34 PM What can we do instead of wall-clock timing? Can we measure JS instruction counts for instance?
Claude App 12:42 PM Yes. For pure-JS hot paths, literal instruction counts: run the benchmark under Valgrind with node --predictable and compare to a checked-in baseline — one run, no statistics needed.
For browser paths there’s no instruction counting under Chromium, but there’s a ladder of other deterministic counts: React commits per interaction, function call counts from V8’s precise coverage, layout and style-recalc counts, DOM mutations. Which do you want first?
S Sam 12:49 PM Let’s explore valgrind + Ir + --predictable in one thread and each of the browser/react benches in new threads. Ping me in all of them. You know what we want. Let’s go
Eleven minutes later, five threads were running, each focused on a different measurement: instruction counts, V8 call counts, React commits, style recalculations, and DOM mutations.
We treated every new benchmark with some skepticism. Each one had two jobs: first, a metric Claude could move in the lab; second, a guardrail in CI with a number that could only ratchet down. If a benchmark was flaky, or if it didn’t actually correlate with user latency, we threw it out rather than let Claude climb the wrong hill.
@Claude please prove that hill climbing against each of these can result in measurable wall clock perf wins. we’ll unship the benches for any candidates that cannot prove that
Wall-clock time is what users feel, but it’s noisy, and milliseconds are too flaky to use as a CI gate. Instruction counts were appealing because they were deterministic, but we still needed Claude to prove they tracked wall-clock time.
So we asked Claude to drive the count down on two hot paths: the routine that assembles a conversation’s message tree, and a scanner for status lines in Claude Code output. Claude profiled both with Valgrind and found that a quarter of the first path’s instructions were megamorphic dictionary lookups, resolving the same message ID three separate times.
An hour later it had cut instructions on both paths by 48% and 31%, and wall-clock time had dropped 78% and 44%. We checked in two new ratchets. From then on, any PR that raised the instruction counts of those paths failed CI, and a daily job lowered each ceiling whenever the count went down.
Message-tree assembly
each message ID resolved
once instead of three times
−48%
−78%
4.6x faster
Status-line scanner
cheap first-character
check before the regex
−31%
−44%
1.8x faster
Two hot paths, instruction count against wall-clock time. Counts under Valgrind with node --predictable ; timings on the same benchmark under plain node with the JIT warm.
Does the count track the clock?
CPU instructions vs. wall-clock time
Message-tree assembly
4.6x faster
each message ID resolved
once instead of three times
CPU instructions
−48%
Wall-clock
−78%
Status-line scanner
1.8x faster
cheap first-character check
before the regex
CPU instructions
−31%
Wall-clock
−44%
Two hot paths, instruction count against wall-clock time. Counts under Valgrind with node --predictable ; timings on the same benchmark under plain node with the JIT warm.
That led us to the central lesson of the sprint. With Claude, measuring something makes it tractable.
Measurement used to be step zero: you’d add a metric, wait for data to roll in, and only then start to understand the problem. With Claude, it’s step one of the climb. As soon as Claude had a number to beat, it could start optimizing. This meant the highest-leverage thing we could do was find more things to measure.
All of this ran in the same Slack channel, with multiple engineers and Claude jamming in every thread. From there, the sprint settled into a loop :
Someone would open a thread about a slow stretch of a journey, often with a screenshot or recording.
Claude would trace the flow, then find or build a benchmark that demonstrated the problem.
Once it had a promising result in the lab, Claude would come back with a PR — often several, sized for risk and review, with anything user-visible behind a flag.
After it shipped, Claude watched the deploy and read the field data.
If performance improved, Claude locked in the win by ratcheting the benchmark down; if not, it turned the flag off and iterated.
Then it went looking for the next slow spot in the same journey.
It got faster
ratchet the
benchmark down
It didn’t
turn the flag off,
iterate
next slow spot
One thread in the loop
Someone opens a thread; Claude takes it from there
Claude
Open a thread
a slow stretch, a recording
Build a benchmark
trace it, reproduce in the lab
Pull request
several, sized for review
Deploy
behind a flag
Read the field
real users, by build and platform
It got faster
ratchet the
benchmark down
It didn’t
turn the flag off,
iterate
next slow spot
An example: someone shared a screen recording that showed sidebar rows popping in after the page loaded. Chat and Cowork rows resolved at different times, making the page feel janky. None of our existing monitors detected it. The closest we had was Cumulative Layout Shift , but each shift only scored about 0.008 — well within the good threshold of 0.1.
Issac had the idea to reference the underlying Layout Instability API directly. Claude created a telemetry event that mapped the sources of each layout-shift entry to a named region (e.g. sidebar, transcript) and phase (e.g. before first paint, after typeable). It added an integration test that opened the page with a populated sidebar, held the sidebar’s data until after first paint, and failed on any shift in any named region. Claude used that as a benchmark to prove a fix: the test went red 20 of 20 runs on main, and green 20 of 20 on the PR.
After the event deployed, Claude read the field data and found that 31% of web page loads moved something after the page was usable, without any user interaction. From there, Claude worked through the causes by name: a header row that arrived late, a caret that slid sideways once the user’s name loaded, a list that moved when the scrollbar popped in. Claude fixed the top offenders as a batch, and when they were gone, it found the next batch.
That was one thread. During the sprint, we ran more than a hundred and fifty at a time.
Once the loop worked on one thread, running it on more was just a matter of opening them. Instead of closing a thread once its original request had been fulfilled, Claude would keep going . An individual thread would put up fifty, sometimes a hundred, optimization PRs. Increasingly, it was Claude, not one of us, opening new threads to chase opportunities it had found on its own, as part of a separate investigation or nightly job. Shelley, one of the engineers in the channel, observed, “[This model] is a numbers demon.”
Every measurement found something to improve. Claude ran a React hook census and found 6,900 hooks and 900 store subscriptions in the composer’s typing path, re-rendering on every keystroke. Claude counted style recalculations and found a single :root:has() selector adding 24 milliseconds to every DOM change. Claude traced code paths after first paint and found a leftover location.reload() causing half a million hidden reloads a day that none of our load metrics could see. Claude read profiler samples from idle tabs and found identical cache snapshots being cloned into IndexedDB twice a minute, all on the main thread.
We rarely knew where a thread would lead. In a sweep for CPU hitches, Claude noticed that highlighting a finished code block could freeze the page for about a second. It dug in the lab and found the culprit: em dashes. If a reply’s markdown contained any non-Latin-1 character, like an

[truncated]

## Original Extract

Inside our performance sprint: the benchmarks Claude built, the loop each Slack thread ran, and the guardrails that let us ship 3,000 changes safely.

How we made claude.ai 3x faster in two weeks / claude.dev [ H ] HOME [ D ] DOCS [ T ] TERMINAL Try Claude Code [ H ] HOME [ D ] DOCS [ T ] TERMINAL Engineering How we made claude.ai 3x faster in two weeks
Once Claude can measure something, it can make it faster. So we kept finding more things to measure.
This August, we made the core user experience of claude.ai and the Claude desktop app about 3x faster in a two-week sprint. Users had been telling us it was slow, and they were right. We ran everything from a single Slack channel, with Claude in every thread.
We focused on four journeys that make up 95% of user activity. At the 75th percentile, time to a typeable page on a fresh load of claude.ai went from 3.1 seconds to 0.55, starting a new Claude Code session went from 0.8 seconds to 0.3, and loading a Claude Cowork cloud session went from 2.6 seconds to 0.73. In aggregate, we estimate that saves tens of thousands of user-hours of waiting every day.
We used Claude Tag (beta), running an internal research model roughly comparable to Opus 5.5. Claude found bottlenecks, built benchmarks, shipped improvements, and watched every deploy. We steered by setting goals, making tradeoffs, and approving every change. With that approach, we merged more than three thousand changes without a single customer-facing incident or rollback. This post covers what we shipped, how we measured it, and the loop we built with Claude to do it safely.
Before the sprint, we created a Slack channel with the following standing instructions :
@Claude Your job is to facilitate all things related to the performance of the claude.ai website and desktop app. Your responsibilities include monitoring deploys for performance regressions, assessing the accuracy and comprehensiveness of existing telemetry, maintaining well-curated observability dashboards, proactively implementing solutions for observed issues and low-hanging fruit, proposing performance project opportunities, and communicating with your human teammates. […]
The ultimate goal for this channel is for you to become as autonomous as possible, but today we know that isn’t yet possible.
We asked Claude to analyze usage data through the Datadog MCP server. It identified the four highest-impact user journeys: launching the app, starting a conversation, loading an existing conversation, and sending a message. Between web and desktop, and across our products, those journeys came to thirteen distinct measurements. To establish baselines, we added instrumentation until they were directly comparable: each started with a user interaction, ended once the result was rendered, and disambiguated client and server work.
We kicked off the sprint with a list of about twenty hand-picked projects, each targeting a specific journey. Claude estimated the impact of each project in milliseconds, and we aggregated those estimates to set our targets for the sprint. Some of the projects were fairly large, but we thought we could probably achieve most of them within two weeks.
We hit twelve of the thirteen targets by day three.
The planned projects landed early. For faster launches, we baked a static composer into the HTML so users can type during React initialization, and precompiled a V8 code cache so the desktop shell’s main process doesn’t recompile from scratch. For faster navigations, we kept the composer mounted between conversations, prefetched sessions when the user hovered over them, and cut sidebar re-renders by 90%.
We had also left room for Claude to identify opportunities and propose new workstreams. Those workstreams quickly ramped into full projects of their own, which far exceeded our initial targets. So we set new targets, then looked for more things to measure:
@Claude we’ve ended up funding nearly every project in the original projects list and more. let’s do a refresh […] what have we not explored, what can we hill climb on, where is the most opportunity at this point? […] i am open to WACKY ideas
From the start, we knew we wanted to iterate faster than our deploy cadence. Claude could work asynchronously for many hours, even overnight, and we wanted to let it validate its prototypes without waiting for field reads. To achieve that, we looked for other ways to measure performance in the lab.
S Sam 12:34 PM What can we do instead of wall-clock timing? Can we measure JS instruction counts for instance?
Claude App 12:42 PM Yes. For pure-JS hot paths, literal instruction counts: run the benchmark under Valgrind with node --predictable and compare to a checked-in baseline — one run, no statistics needed.
For browser paths there’s no instruction counting under Chromium, but there’s a ladder of other deterministic counts: React commits per interaction, function call counts from V8’s precise coverage, layout and style-recalc counts, DOM mutations. Which do you want first?
S Sam 12:49 PM Let’s explore valgrind + Ir + --predictable in one thread and each of the browser/react benches in new threads. Ping me in all of them. You know what we want. Let’s go
Eleven minutes later, five threads were running, each focused on a different measurement: instruction counts, V8 call counts, React commits, style recalculations, and DOM mutations.
We treated every new benchmark with some skepticism. Each one had two jobs: first, a metric Claude could move in the lab; second, a guardrail in CI with a number that could only ratchet down. If a benchmark was flaky, or if it didn’t actually correlate with user latency, we threw it out rather than let Claude climb the wrong hill.
@Claude please prove that hill climbing against each of these can result in measurable wall clock perf wins. we’ll unship the benches for any candidates that cannot prove that
Wall-clock time is what users feel, but it’s noisy, and milliseconds are too flaky to use as a CI gate. Instruction counts were appealing because they were deterministic, but we still needed Claude to prove they tracked wall-clock time.
So we asked Claude to drive the count down on two hot paths: the routine that assembles a conversation’s message tree, and a scanner for status lines in Claude Code output. Claude profiled both with Valgrind and found that a quarter of the first path’s instructions were megamorphic dictionary lookups, resolving the same message ID three separate times.
An hour later it had cut instructions on both paths by 48% and 31%, and wall-clock time had dropped 78% and 44%. We checked in two new ratchets. From then on, any PR that raised the instruction counts of those paths failed CI, and a daily job lowered each ceiling whenever the count went down.
Message-tree assembly
each message ID resolved
once instead of three times
−48%
−78%
4.6x faster
Status-line scanner
cheap first-character
check before the regex
−31%
−44%
1.8x faster
Two hot paths, instruction count against wall-clock time. Counts under Valgrind with node --predictable ; timings on the same benchmark under plain node with the JIT warm.
Does the count track the clock?
CPU instructions vs. wall-clock time
Message-tree assembly
4.6x faster
each message ID resolved
once instead of three times
CPU instructions
−48%
Wall-clock
−78%
Status-line scanner
1.8x faster
cheap first-character check
before the regex
CPU instructions
−31%
Wall-clock
−44%
Two hot paths, instruction count against wall-clock time. Counts under Valgrind with node --predictable ; timings on the same benchmark under plain node with the JIT warm.
That led us to the central lesson of the sprint. With Claude, measuring something makes it tractable.
Measurement used to be step zero: you’d add a metric, wait for data to roll in, and only then start to understand the problem. With Claude, it’s step one of the climb. As soon as Claude had a number to beat, it could start optimizing. This meant the highest-leverage thing we could do was find more things to measure.
All of this ran in the same Slack channel, with multiple engineers and Claude jamming in every thread. From there, the sprint settled into a loop :
Someone would open a thread about a slow stretch of a journey, often with a screenshot or recording.
Claude would trace the flow, then find or build a benchmark that demonstrated the problem.
Once it had a promising result in the lab, Claude would come back with a PR — often several, sized for risk and review, with anything user-visible behind a flag.
After it shipped, Claude watched the deploy and read the field data.
If performance improved, Claude locked in the win by ratcheting the benchmark down; if not, it turned the flag off and iterated.
Then it went looking for the next slow spot in the same journey.
It got faster
ratchet the
benchmark down
It didn’t
turn the flag off,
iterate
next slow spot
One thread in the loop
Someone opens a thread; Claude takes it from there
Claude
Open a thread
a slow stretch, a recording
Build a benchmark
trace it, reproduce in the lab
Pull request
several, sized for review
Deploy
behind a flag
Read the field
real users, by build and platform
It got faster
ratchet the
benchmark down
It didn’t
turn the flag off,
iterate
next slow spot
An example: someone shared a screen recording that showed sidebar rows popping in after the page loaded. Chat and Cowork rows resolved at different times, making the page feel janky. None of our existing monitors detected it. The closest we had was Cumulative Layout Shift , but each shift only scored about 0.008 — well within the good threshold of 0.1.
Issac had the idea to reference the underlying Layout Instability API directly. Claude created a telemetry event that mapped the sources of each layout-shift entry to a named region (e.g. sidebar, transcript) and phase (e.g. before first paint, after typeable). It added an integration test that opened the page with a populated sidebar, held the sidebar’s data until after first paint, and failed on any shift in any named region. Claude used that as a benchmark to prove a fix: the test went red 20 of 20 runs on main, and green 20 of 20 on the PR.
After the event deployed, Claude read the field data and found that 31% of web page loads moved something after the page was usable, without any user interaction. From there, Claude worked through the causes by name: a header row that arrived late, a caret that slid sideways once the user’s name loaded, a list that moved when the scrollbar popped in. Claude fixed the top offenders as a batch, and when they were gone, it found the next batch.
That was one thread. During the sprint, we ran more than a hundred and fifty at a time.
Once the loop worked on one thread, running it on more was just a matter of opening them. Instead of closing a thread once its original request had been fulfilled, Claude would keep going . An individual thread would put up fifty, sometimes a hundred, optimization PRs. Increasingly, it was Claude, not one of us, opening new threads to chase opportunities it had found on its own, as part of a separate investigation or nightly job. Shelley, one of the engineers in the channel, observed, “[This model] is a numbers demon.”
Every measurement found something to improve. Claude ran a React hook census and found 6,900 hooks and 900 store subscriptions in the composer’s typing path, re-rendering on every keystroke. Claude counted style recalculations and found a single :root:has() selector adding 24 milliseconds to every DOM change. Claude traced code paths after first paint and found a leftover location.reload() causing half a million hidden reloads a day that none of our load metrics could see. Claude read profiler samples from idle tabs and found identical cache snapshots being cloned into IndexedDB twice a minute, all on the main thread.
We rarely knew where a thread would lead. In a sweep for CPU hitches, Claude noticed that highlighting a finished code block could freeze the page for about a second. It dug in the lab and found the culprit: em dashes. If a reply’s markdown contained any non-Latin-1 character, like an

[truncated]
