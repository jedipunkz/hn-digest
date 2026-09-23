---
source: "https://vibeladder.dev/results/r2.html"
hn_url: "https://news.ycombinator.com/item?id=49817230"
title: "Show HN: Vibe-coder ladder round 2 – the AI ran 38 minutes past the deadline"
article_title: "Round 2 results — VIBELADDER"
image: "https://vibeladder.dev/assets/og-v3.png"
author: "alviso"
captured_at: "2026-09-23T15:19:40Z"
capture_tool: "hn-digest"
hn_id: 49817230
score: 1
comments: 0
posted_at: "2026-09-23T15:02:23Z"
tags:
  - hacker-news
---

# Show HN: Vibe-coder ladder round 2 – the AI ran 38 minutes past the deadline

- HN: [49817230](https://news.ycombinator.com/item?id=49817230)
- Source: [vibeladder.dev](https://vibeladder.dev/results/r2.html)
- Score: 1
- Comments: 0
- Posted: 2026-09-23T15:02:23Z

## Translation

Title: Show HN: Vibe-coder ladder round 2 – the AI ran 38 minutes past the deadline
Article title: Round 2 results — VIBELADDER
Description: Round 2, Four in a Row: two entries, a close verdict, and the season back to level. The judge

Article text:
🪜 ">
▸ VIBELADDER
How it works
Task
Schedule
Rules
Sponsors
Leaderboard
Get in line
Round 2 · Four in a Row · 19 September 2026
Fourteen minutes, three prompts, and the season is level.
Round 2 ran on Saturday. The task, secret until 17:00 CET, was Connect Four against a computer opponent in one HTML file: seven by six, at least two genuinely different difficulty levels, win and draw detection with a clear announcement, a new-game control, and on the harder setting the computer has to actually try to win. Same two entrants as Round 1. Different result.
Both entries passed all five items of the functional checklist, no network violations, no injection attempts. One head to head, judged by an LLM that plays both games side by side through the interface only and never sees the source. @schwarzkopfb won it, and the judge called it close.
The season is now exactly level. One win each, against each other, and the Bradley-Terry fit puts both at 1100, which is the baseline everyone starts from. That is not a rounding accident. Two results that cancel out are, honestly read, no evidence that either player is stronger than the other, and the rating says so. The 1192 from Round 1 rested on a single comparison, as the leaderboard warned it did.
The judge played full games on both apps, on Easy and on Hard, and wrote one note per dimension. A is @P1s0, B is @schwarzkopfb. This is the full text.
We checked the two visible complaints against A's file afterwards, because a judge's word alone is not a receipt. The page is 837 pixels tall on the judge's 800 pixel viewport, so it does need scrolling, by 37 pixels. And the 'Test unlock active' badge is a debug tag that is meant to be hidden: the element carries the hidden attribute, but a stylesheet rule on the same element overrides it. Two small things. This is what "close" looks like.
@schwarzkopfb's session ran fourteen minutes and three prompts, in OpenCode with GPT-5.6 Sol. He submitted at 15:26, twenty-seven minutes after the drop, and did not resubmit. In Round 1 his write-up said the format rewards a second look. He took a different lesson from it: look first.
let's plan an implementation for this: [pasted the full task and the short rules]
looks good, let's refine some details. the instructions doen't say that we cannot use semi-persistent storage apis, like local storage. we'd include some convenience features based on this
let's create an initial implementation based on this plan
@P1s0's session is the more instructive one this round, and we could only read it after the export arrived by email (the cowork link that was submitted only opens for its owner, which is why the submit page now warns about that). @P1s0 ran the same method that won Round 1: nine discussion prompts in 25 minutes. The AI was asked to explain the rules in a table, to say what "harder" and "actually trying to win" mean in measurable terms, to list decisions, to check the plan against the brief (it found five things, including that four or more in a row must count), to list twelve risks with a fix and a test for each, and then to write a detailed build prompt. Then, at 11:32 local time, 32 minutes in: "Analyze the prompt then create it".
That one build turn ran for 1 hour and 35 minutes , 75 tool calls, and finished 38 minutes after the window had closed. Along the way it wrote a bitboard solver, verified it against a reference solver with zero mismatches, generated a 154,459-position opening book, and ran 19 rule tests, 48 tactic positions and 29 browser checks. None of that reached the judge. What was submitted, ninety seconds before the close, was a snapshot of the file as it stood mid-build. The build log itself lists the missing [hidden] rule behind the "Test unlock active" badge as a bug found and fixed later; the submitted copy predates the fix. The 19 KB entry has no opening book because the book was still generating.
Shipping what exists at the deadline was the right call, and that snapshot still won the functional depth dimension outright. The lesson is about the clock, not the method: the ninety minutes include the AI's time, and an agent told to build, test and prove everything will happily use two hours. Round 1 was won by asking for a first version and iterating. Round 2 was lost by asking for the finished version.
Round 1 was won by an hour of iteration. Round 2 was won by fourteen minutes of planning. Both winning sessions made the AI describe what it was going to build, and checked that against the rules, before letting it build. The build itself was one prompt both times. The difference between the Round 2 winner and the Round 2 runner-up was not planning, both planned. It was the size of the build they asked for: an "initial implementation" that arrived in four minutes, against a full verified build that needed 95.
Time-box the build. Ask for a first working version, submit it, then iterate with whatever is left. The rules allow resubmitting until the close for exactly this reason. An agent that is still running when the window shuts has produced nothing you can enter.
The rating did what we said it would. A 1192 built on one comparison lasted exactly one round. Two wins each way is a level season, and the board says 1100 for both. Anyone joining in Round 3 starts from the same number as the two people who have played twice.
Two of twenty-two, again. Six opened the task, two shipped. We will keep printing the ratio. Round 3 is Saturday 26 September at 17:00 CET, and a seat covers the rest of the season.
every result ships with its receipts

## Original Extract

Round 2, Four in a Row: two entries, a close verdict, and the season back to level. The judge

🪜 ">
▸ VIBELADDER
How it works
Task
Schedule
Rules
Sponsors
Leaderboard
Get in line
Round 2 · Four in a Row · 19 September 2026
Fourteen minutes, three prompts, and the season is level.
Round 2 ran on Saturday. The task, secret until 17:00 CET, was Connect Four against a computer opponent in one HTML file: seven by six, at least two genuinely different difficulty levels, win and draw detection with a clear announcement, a new-game control, and on the harder setting the computer has to actually try to win. Same two entrants as Round 1. Different result.
Both entries passed all five items of the functional checklist, no network violations, no injection attempts. One head to head, judged by an LLM that plays both games side by side through the interface only and never sees the source. @schwarzkopfb won it, and the judge called it close.
The season is now exactly level. One win each, against each other, and the Bradley-Terry fit puts both at 1100, which is the baseline everyone starts from. That is not a rounding accident. Two results that cancel out are, honestly read, no evidence that either player is stronger than the other, and the rating says so. The 1192 from Round 1 rested on a single comparison, as the leaderboard warned it did.
The judge played full games on both apps, on Easy and on Hard, and wrote one note per dimension. A is @P1s0, B is @schwarzkopfb. This is the full text.
We checked the two visible complaints against A's file afterwards, because a judge's word alone is not a receipt. The page is 837 pixels tall on the judge's 800 pixel viewport, so it does need scrolling, by 37 pixels. And the 'Test unlock active' badge is a debug tag that is meant to be hidden: the element carries the hidden attribute, but a stylesheet rule on the same element overrides it. Two small things. This is what "close" looks like.
@schwarzkopfb's session ran fourteen minutes and three prompts, in OpenCode with GPT-5.6 Sol. He submitted at 15:26, twenty-seven minutes after the drop, and did not resubmit. In Round 1 his write-up said the format rewards a second look. He took a different lesson from it: look first.
let's plan an implementation for this: [pasted the full task and the short rules]
looks good, let's refine some details. the instructions doen't say that we cannot use semi-persistent storage apis, like local storage. we'd include some convenience features based on this
let's create an initial implementation based on this plan
@P1s0's session is the more instructive one this round, and we could only read it after the export arrived by email (the cowork link that was submitted only opens for its owner, which is why the submit page now warns about that). @P1s0 ran the same method that won Round 1: nine discussion prompts in 25 minutes. The AI was asked to explain the rules in a table, to say what "harder" and "actually trying to win" mean in measurable terms, to list decisions, to check the plan against the brief (it found five things, including that four or more in a row must count), to list twelve risks with a fix and a test for each, and then to write a detailed build prompt. Then, at 11:32 local time, 32 minutes in: "Analyze the prompt then create it".
That one build turn ran for 1 hour and 35 minutes , 75 tool calls, and finished 38 minutes after the window had closed. Along the way it wrote a bitboard solver, verified it against a reference solver with zero mismatches, generated a 154,459-position opening book, and ran 19 rule tests, 48 tactic positions and 29 browser checks. None of that reached the judge. What was submitted, ninety seconds before the close, was a snapshot of the file as it stood mid-build. The build log itself lists the missing [hidden] rule behind the "Test unlock active" badge as a bug found and fixed later; the submitted copy predates the fix. The 19 KB entry has no opening book because the book was still generating.
Shipping what exists at the deadline was the right call, and that snapshot still won the functional depth dimension outright. The lesson is about the clock, not the method: the ninety minutes include the AI's time, and an agent told to build, test and prove everything will happily use two hours. Round 1 was won by asking for a first version and iterating. Round 2 was lost by asking for the finished version.
Round 1 was won by an hour of iteration. Round 2 was won by fourteen minutes of planning. Both winning sessions made the AI describe what it was going to build, and checked that against the rules, before letting it build. The build itself was one prompt both times. The difference between the Round 2 winner and the Round 2 runner-up was not planning, both planned. It was the size of the build they asked for: an "initial implementation" that arrived in four minutes, against a full verified build that needed 95.
Time-box the build. Ask for a first working version, submit it, then iterate with whatever is left. The rules allow resubmitting until the close for exactly this reason. An agent that is still running when the window shuts has produced nothing you can enter.
The rating did what we said it would. A 1192 built on one comparison lasted exactly one round. Two wins each way is a level season, and the board says 1100 for both. Anyone joining in Round 3 starts from the same number as the two people who have played twice.
Two of twenty-two, again. Six opened the task, two shipped. We will keep printing the ratio. Round 3 is Saturday 26 September at 17:00 CET, and a seat covers the rest of the season.
every result ships with its receipts
