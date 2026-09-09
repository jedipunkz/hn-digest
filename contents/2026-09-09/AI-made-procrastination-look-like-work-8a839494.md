---
source: "https://nedio.xyz/research/ai-procrastination-in-software-development"
hn_url: "https://news.ycombinator.com/item?id=49625196"
title: "AI made procrastination look like work"
article_title: "AI Made Procrastination Look Like Work | NEDIO"
image: "https://nedio.xyz/research/ai-procrastination-in-software-development/opengraph-image"
author: "thaziev"
captured_at: "2026-09-09T12:46:54Z"
capture_tool: "hn-digest"
hn_id: 49625196
score: 1
comments: 0
posted_at: "2026-09-09T12:10:28Z"
tags:
  - hacker-news
---

# AI made procrastination look like work

- HN: [49625196](https://news.ycombinator.com/item?id=49625196)
- Source: [nedio.xyz](https://nedio.xyz/research/ai-procrastination-in-software-development)
- Score: 1
- Comments: 0
- Posted: 2026-09-09T12:10:28Z

## Translation

Title: AI made procrastination look like work
Article title: AI Made Procrastination Look Like Work | NEDIO
Description: An agent can generate a defensible refactor in minutes while the bug you were avoiding still costs what it always did. On productive detours in AI-assisted development — and the one-line note that catches them.

Article text:
AI Made Procrastination Look Like Work | NEDIO NEDIO Sprint AI Review VS Code Pricing Compare Guides Research Start a Sprint Home
AI made procrastination look like work
An agent can produce a defensible refactor in minutes. The task you were avoiding still costs exactly what it did before, so the cheap work is now the one that feels like progress.
This is an argument, not a study: what changes when a useful detour becomes almost free, what the organizational data can and cannot tell us about it, and the smallest note that catches the trade while you are making it.
Consider a morning spent fixing an intermittent login failure.
You ask an agent to explain the authentication flow. It finds some awkward code and suggests a cleaner approach. You ask what that would look like. Then you ask it to implement the approach, add tests, and handle a few edge cases.
An hour later, you have a substantial diff. The new tests pass. You can explain why the code is better.
You still haven’t reproduced the login failure.
That was the unpleasant part. It meant digging through incomplete logs, questioning an assumption, perhaps discovering that code you wrote last month was wrong. The refactor gave you something useful and much more comfortable to do instead.
There is a kind of procrastination that survives every website blocker, because it happens inside your work.
Programmers already have names for nearby habits. Yak shaving sends you through a chain of preparatory tasks. Bikeshedding absorbs attention in details that are easier to discuss than the difficult question. Both predate AI, and both can leave useful work behind.
A cleaner authentication module has value. You can defend the refactor in review and mention the new tests in your update. Neither requires you to explain why you chose that work while the login failure was still waiting.
My hypothesis is about the relative cost of those choices. AI can make a useful detour much cheaper while leaving the discomfort of the original task intact.
An agent can generate a replacement module in minutes. Asking a colleague about an undocumented field still involves waiting for a reply. Investigating the failing case still risks discovering that your approach was wrong. The easier activity now offers an immediate, tangible result.
It supplies its own next steps
The replacement needs tests, the tests suggest edge cases, and the edge cases invite another abstraction. You can stay busy following that sequence without returning to the decision that sent you down it.
That is the part I find genuinely new. Yak shaving ran out of yak eventually. A prompt loop does not have a natural end, because every artifact it produces suggests the next one.
What the org data shows, and what it doesn’t
There is a related gap in organizational data. Faros AI’s 2025 analysis, published in July and covering data through June, looked at more than 10,000 developers across 1,255 teams. Developers on teams with high AI adoption merged 98% more pull requests, while PR review time increased 91%. At the company level, the researchers found no significant correlation between AI adoption and improvements in delivery metrics.
Review bottlenecks could explain that gap on their own; individual procrastination cannot be inferred from it. Two numbers from a vendor’s telemetry are not evidence for the argument in this essay. They are a reminder to ask what the increased activity actually accomplished.
What the research actually covers
An ICSE 2025 interview study by Saghi, Zimmermann, and Chattopadhyay asked fifteen developers about procrastination. Participants described task-related, personal, and external triggers, and both harmful and beneficial consequences of delay.
Those interviews offer reasons to examine the task and its context when someone puts work off. They do not establish whether AI changes that behavior, and fifteen interviews would not settle it if they tried.
So I don’t know whether AI makes developers procrastinate more overall. It can clearly help someone get past a blank page, understand unfamiliar code, or work out which question to ask. Those benefits can coexist with a new way of getting stuck.
An unfinished task proves very little
Debugging can take a day and leave you with three discarded hypotheses and no patch. A prototype can reveal that a feature is a bad idea. Even the authentication refactor may be necessary to understand the failure. The question is how the work advances the investigation.
Can you now reproduce the failure? Have you ruled out a cause? Did the refactor make a previously hidden state observable? If so, you can describe how it moved the investigation forward.
If all you can say is that the new code is cleaner, it is worth checking whether the original obstacle is still sitting exactly where you left it.
A merged PR tells you that a change was accepted. Whether it addressed the problem you sat down to solve requires looking back at that problem.
The smallest check I can suggest is to leave yourself a note before opening the agent.
Write down what you want to have changed by the end of this block of work. For an investigation, write down the uncertainty you want to reduce.
Before:
Goal — one reproducible case of the login failure, and
whether the session is missing before or after the redirect.
After:
What happened —
Still unresolved —
Changed direction because — Two fields before, three after. Anything longer becomes its own way to avoid the task. You might add logging, read the framework source, or have an agent build a test harness. The note gives you something to look back at when you’re forty minutes into rewriting middleware.
When you stop, add what actually happened and what remains unresolved. If you changed direction, record why. Discovering that the original goal was wrong is a perfectly good reason to abandon it.
You can do this in an issue, a text file, or on paper. There’s no need to turn it into a score. A completion percentage would introduce its own problems: smaller promises are easier to keep, difficult investigations look bad, and abandoned ideas can be valuable results.
The immediate value is having a record of what mattered before the work became absorbing. You get to compare it with your explanation afterwards.
When the same obstacle keeps showing up
If the same obstacle keeps appearing in those notes while the surrounding code keeps improving, you have something specific to examine. Perhaps you’re blocked. Perhaps the task needs to be split up. Perhaps you know the next step and would rather generate something else.
In that last case, another implementation is unlikely to help. Close the diff for a moment. Open the logs, send the question, or run the test you’ve been putting off. Let the agent help with that.
The authentication module can be prettier tomorrow. Right now, someone still can’t log in.
Saghi, Z., Zimmermann, T., & Chattopadhyay, S. (2025). “Code Today, Deadline Tomorrow: Procrastination Among Software Developers.” ICSE 2025 . Interviews with fifteen developers on why work gets delayed: task-related, personal, and external triggers, plus both harmful and beneficial consequences. Qualitative and small; it describes the behavior, it does not measure AI’s effect on it.
Faros AI (2025). “The AI productivity paradox.” Vendor telemetry across 10,000+ developers and 1,255 teams: 98% more merged PRs on high-adoption teams, 91% longer review times, and no significant company-level correlation with delivery metrics. Cited for the size of the activity-to-outcome gap, not for its cause.
It supplies its own next steps
An unfinished task proves little
Where this argument connects to the rest of the work on developer attention.
Context switching cost for developers
The throughput version of the same problem: why busy days produce nothing shippable.
What actually protects a block once you have decided what belongs inside it.
The other failure mode — losing the thread rather than trading it for a nicer one.
What the tool does with the notes you leave at the start and end of a session.
NEDIO is the timer I built around that note
A sprint starts with what you intend to change and ends with what actually happened. Free, no account needed to try it.
Focus experience for developers. One tab. One sprint. One commit.
Best Focus Apps for Developers
© 2026 NEDIO. All rights reserved.
Cookies & privacy. Essential cookies keep sign-in and the core features working. Optional ones help us measure usage and improve NEDIO. Privacy Policy · Cookie Policy · Do Not Sell/Share

## Original Extract

An agent can generate a defensible refactor in minutes while the bug you were avoiding still costs what it always did. On productive detours in AI-assisted development — and the one-line note that catches them.

AI Made Procrastination Look Like Work | NEDIO NEDIO Sprint AI Review VS Code Pricing Compare Guides Research Start a Sprint Home
AI made procrastination look like work
An agent can produce a defensible refactor in minutes. The task you were avoiding still costs exactly what it did before, so the cheap work is now the one that feels like progress.
This is an argument, not a study: what changes when a useful detour becomes almost free, what the organizational data can and cannot tell us about it, and the smallest note that catches the trade while you are making it.
Consider a morning spent fixing an intermittent login failure.
You ask an agent to explain the authentication flow. It finds some awkward code and suggests a cleaner approach. You ask what that would look like. Then you ask it to implement the approach, add tests, and handle a few edge cases.
An hour later, you have a substantial diff. The new tests pass. You can explain why the code is better.
You still haven’t reproduced the login failure.
That was the unpleasant part. It meant digging through incomplete logs, questioning an assumption, perhaps discovering that code you wrote last month was wrong. The refactor gave you something useful and much more comfortable to do instead.
There is a kind of procrastination that survives every website blocker, because it happens inside your work.
Programmers already have names for nearby habits. Yak shaving sends you through a chain of preparatory tasks. Bikeshedding absorbs attention in details that are easier to discuss than the difficult question. Both predate AI, and both can leave useful work behind.
A cleaner authentication module has value. You can defend the refactor in review and mention the new tests in your update. Neither requires you to explain why you chose that work while the login failure was still waiting.
My hypothesis is about the relative cost of those choices. AI can make a useful detour much cheaper while leaving the discomfort of the original task intact.
An agent can generate a replacement module in minutes. Asking a colleague about an undocumented field still involves waiting for a reply. Investigating the failing case still risks discovering that your approach was wrong. The easier activity now offers an immediate, tangible result.
It supplies its own next steps
The replacement needs tests, the tests suggest edge cases, and the edge cases invite another abstraction. You can stay busy following that sequence without returning to the decision that sent you down it.
That is the part I find genuinely new. Yak shaving ran out of yak eventually. A prompt loop does not have a natural end, because every artifact it produces suggests the next one.
What the org data shows, and what it doesn’t
There is a related gap in organizational data. Faros AI’s 2025 analysis, published in July and covering data through June, looked at more than 10,000 developers across 1,255 teams. Developers on teams with high AI adoption merged 98% more pull requests, while PR review time increased 91%. At the company level, the researchers found no significant correlation between AI adoption and improvements in delivery metrics.
Review bottlenecks could explain that gap on their own; individual procrastination cannot be inferred from it. Two numbers from a vendor’s telemetry are not evidence for the argument in this essay. They are a reminder to ask what the increased activity actually accomplished.
What the research actually covers
An ICSE 2025 interview study by Saghi, Zimmermann, and Chattopadhyay asked fifteen developers about procrastination. Participants described task-related, personal, and external triggers, and both harmful and beneficial consequences of delay.
Those interviews offer reasons to examine the task and its context when someone puts work off. They do not establish whether AI changes that behavior, and fifteen interviews would not settle it if they tried.
So I don’t know whether AI makes developers procrastinate more overall. It can clearly help someone get past a blank page, understand unfamiliar code, or work out which question to ask. Those benefits can coexist with a new way of getting stuck.
An unfinished task proves very little
Debugging can take a day and leave you with three discarded hypotheses and no patch. A prototype can reveal that a feature is a bad idea. Even the authentication refactor may be necessary to understand the failure. The question is how the work advances the investigation.
Can you now reproduce the failure? Have you ruled out a cause? Did the refactor make a previously hidden state observable? If so, you can describe how it moved the investigation forward.
If all you can say is that the new code is cleaner, it is worth checking whether the original obstacle is still sitting exactly where you left it.
A merged PR tells you that a change was accepted. Whether it addressed the problem you sat down to solve requires looking back at that problem.
The smallest check I can suggest is to leave yourself a note before opening the agent.
Write down what you want to have changed by the end of this block of work. For an investigation, write down the uncertainty you want to reduce.
Before:
Goal — one reproducible case of the login failure, and
whether the session is missing before or after the redirect.
After:
What happened —
Still unresolved —
Changed direction because — Two fields before, three after. Anything longer becomes its own way to avoid the task. You might add logging, read the framework source, or have an agent build a test harness. The note gives you something to look back at when you’re forty minutes into rewriting middleware.
When you stop, add what actually happened and what remains unresolved. If you changed direction, record why. Discovering that the original goal was wrong is a perfectly good reason to abandon it.
You can do this in an issue, a text file, or on paper. There’s no need to turn it into a score. A completion percentage would introduce its own problems: smaller promises are easier to keep, difficult investigations look bad, and abandoned ideas can be valuable results.
The immediate value is having a record of what mattered before the work became absorbing. You get to compare it with your explanation afterwards.
When the same obstacle keeps showing up
If the same obstacle keeps appearing in those notes while the surrounding code keeps improving, you have something specific to examine. Perhaps you’re blocked. Perhaps the task needs to be split up. Perhaps you know the next step and would rather generate something else.
In that last case, another implementation is unlikely to help. Close the diff for a moment. Open the logs, send the question, or run the test you’ve been putting off. Let the agent help with that.
The authentication module can be prettier tomorrow. Right now, someone still can’t log in.
Saghi, Z., Zimmermann, T., & Chattopadhyay, S. (2025). “Code Today, Deadline Tomorrow: Procrastination Among Software Developers.” ICSE 2025 . Interviews with fifteen developers on why work gets delayed: task-related, personal, and external triggers, plus both harmful and beneficial consequences. Qualitative and small; it describes the behavior, it does not measure AI’s effect on it.
Faros AI (2025). “The AI productivity paradox.” Vendor telemetry across 10,000+ developers and 1,255 teams: 98% more merged PRs on high-adoption teams, 91% longer review times, and no significant company-level correlation with delivery metrics. Cited for the size of the activity-to-outcome gap, not for its cause.
It supplies its own next steps
An unfinished task proves little
Where this argument connects to the rest of the work on developer attention.
Context switching cost for developers
The throughput version of the same problem: why busy days produce nothing shippable.
What actually protects a block once you have decided what belongs inside it.
The other failure mode — losing the thread rather than trading it for a nicer one.
What the tool does with the notes you leave at the start and end of a session.
NEDIO is the timer I built around that note
A sprint starts with what you intend to change and ends with what actually happened. Free, no account needed to try it.
Focus experience for developers. One tab. One sprint. One commit.
Best Focus Apps for Developers
© 2026 NEDIO. All rights reserved.
Cookies & privacy. Essential cookies keep sign-in and the core features working. Optional ones help us measure usage and improve NEDIO. Privacy Policy · Cookie Policy · Do Not Sell/Share
