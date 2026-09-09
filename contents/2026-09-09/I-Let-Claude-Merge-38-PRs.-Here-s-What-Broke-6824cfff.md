---
source: "https://www.craigphares.com/i-let-claude-merge-38-prs-heres-what-broke/"
hn_url: "https://news.ycombinator.com/item?id=49629668"
title: "I Let Claude Merge 38 PRs. Here's What Broke"
article_title: "I Let Claude Merge 38 PRs. Here’s What Broke. - Craig Phares"
image: "https://www.craigphares.com/wp-content/uploads/2026/09/hero-1.png"
author: "craigphares"
captured_at: "2026-09-09T17:03:36Z"
capture_tool: "hn-digest"
hn_id: 49629668
score: 1
comments: 0
posted_at: "2026-09-09T17:01:05Z"
tags:
  - hacker-news
---

# I Let Claude Merge 38 PRs. Here's What Broke

- HN: [49629668](https://news.ycombinator.com/item?id=49629668)
- Source: [www.craigphares.com](https://www.craigphares.com/i-let-claude-merge-38-prs-heres-what-broke/)
- Score: 1
- Comments: 0
- Posted: 2026-09-09T17:01:05Z

## Translation

Title: I Let Claude Merge 38 PRs. Here's What Broke
Article title: I Let Claude Merge 38 PRs. Here’s What Broke. - Craig Phares
Description: A field report on misleading green checks, durable agent state, and keeping production behind a human gate. Claude merged the sixth and final phase PR, then went idle. The plan still had one job left: open an integration PR with…

Article text:
I Let Claude Merge 38 PRs. Here’s What Broke. - Craig Phares
Skip to content
Craig Phares
I Let Claude Merge 38 PRs. Here’s What Broke.
A field report on misleading green checks, durable agent state, and keeping production behind a human gate.
Claude merged the sixth and final phase PR, then went idle.
The plan still had one job left: open an integration PR with a checklist for me to test before anything reached production.
All six phase PRs had merged. The feature branch contained the finished work. CI was green. There was no error and nothing was visibly stuck. From GitHub, the plan looked complete.
The instructions told the agent to open the integration PR, but two details pulled it in another direction. Every transition in the driver pointed to the next action except this one. After the final phase merged, the driver looked for newly unblocked work, found none, and stopped.
I had also described the integration PR as “the only human merge gate.” I meant that a human must merge it. Claude interpreted it to mean that a human must give permission before it could even be opened.
Then it saved that interpretation to the plan.
Every session after that woke up, read the same decision, and treated it as settled. The system had given a temporary mistake durable memory.
This is the kind of failure I kept finding while building claude-phases . The agent usually wrote good code. But knowing whether the surrounding work had actually happened was tricky.
I use Claude Code to work through features in phases. A planner breaks the feature into small steps and writes a plan. A driver implements each step in sequence.
Every phase ends in a PR against a feature branch. Once the required checks pass, the driver merges that PR and starts the next phase. Each phase is small enough to inspect and safe enough to merge without leaving the feature branch broken.
The driver never merges to main .
When all phases are complete, it opens one integration PR from the feature branch to main . That PR contains a human UAT checklist. I test the finished feature and decide whether it ships.
Since the end of July, claude-phases has merged 38 phases across five repositories. It’s written nine plans and run five of them from beginning to end. Four of those phases were claude-phases building itself.
Letting the driver merge its own code worked sooner than I expected. Teaching it how to prove that it was safe to continue took much longer.
The plan has to survive the session
Claude Code cloud sessions are temporary. They pause, compact, restart, and occasionally die in the middle of something.
A new session needs to recover from whatever state the previous session left behind. It might wake up during implementation, after a PR was opened, or after a merge that was never recorded.
I decided that a fresh session could trust two things: the plan file and GitHub. Everything else had to be reconstructed from those.
The plan lives on its own branch. Phase PRs are never allowed to modify it. This means a plan can have one branch recording the state of the work while another branch contains the work itself. It’s a little strange, but it prevents an important failure.
My earlier tool, claude-project-manager , recorded phase status inside the phase PR. That worked while an external timer held the current cursor. It would fail here because an unmerged PR hides its status changes from a new session. The new session would see the phase marked pending and start the same work again.
The driver also records state before taking the action that state describes. It claims the phase before creating its branch. It records the PR before waiting on it.
That ordering closes the gaps where a session can perform an action, die before recording it, and cause the next session to repeat it.
Six green checks, six different meanings
I set up an automated code review workflow on one repository. It produced six successful runs.
Each green check meant something different.
The first run finished in 10 seconds. It skipped the review because the workflow file differed from the copy on the default branch.
The second ran for eight and a half minutes and completed a real review. It didn’t use the --comment flag, so its findings disappeared into the job log.
The third finished in 20 seconds and produced no artifact.
The fourth found a real bug but lacked permission to post it.
The fifth stopped while waiting for a sub-agent.
The sixth found four real bugs. But every gh command it used to report them was denied.
GitHub marked all six runs as successful. None had completed the job I cared about.
A successful workflow run only proves that the workflow exited successfully. If the gate is supposed to represent a code review, I need evidence that a review took place.
claude-phases now supports review artifacts. A check can be required to produce something the reviewer actually wrote before the driver accepts it.
That rule needs configuration because reviewers behave differently. Some always produce a summary. Others only leave inline comments when they find a problem. A clean review from the second kind produces nothing, which is indistinguishable from a reviewer that never ran.
The driver cannot infer the difference. The repository has to define what counts as proof.
I also changed the review workflow so it fails loudly when it can’t complete its review. Skipping gracefully would create another green check with no review behind it.
CI introduced a few more ambiguous states.
A repository with no checks looks exactly like one whose CI stopped running. claude-phases blocks when it sees zero checks unless the repository has been explicitly configured to allow them.
Neutral checks need different handling. GitHub considers them non-failing, and some review tools use neutral specifically to avoid blocking merges. Anthropic’s managed Claude Code Review concludes with a neutral status. A gate that waits only for success will wait forever after the review has already finished.
Required checks can also be impossible to satisfy. A workflow may be configured to run only on PRs targeting main , while each phase PR targets a feature branch. The driver waits for the check by name, but GitHub will never create it.
That failure may not become obvious until hours after the plan was written.
The safest default I found is simple:
A false pass is silent and permanent. A false block is visible, and I can fix it within the phase.
The missing integration PR exposed a different class of problem.
Most of the driver followed an explicit chain. After one action completed, the instructions pointed to the next action. Completion was the exception. It was described as an outcome rather than implemented as a transition.
The driver reached the end of the chain and stopped.
The phrase “human merge gate” added just enough ambiguity for Claude to invent the wrong behavior. Once written into the plan, that interpretation survived every new session.
I rewrote completion as an explicit transition with its own inbound pointer. After the final phase merges, the driver opens the integration PR. Opening it is an agent action. Merging it is a human action.
There is no longer anything to infer from the surrounding language.
I put the checklist where I will see it
My first version put the UAT checklist on the final phase PR.
In automatic mode, the driver would open that PR, wait for its checks, and merge it about a minute later. It created a thorough checklist in a place I would probably never visit.
The plan completed. The checklist existed. And nobody tested it.
I moved the checklist to the integration PR because that is where I re-enter the process.
I also stopped grouping tests by phase. Phase order describes how the software was produced, which is rarely how someone should test it. It can also preserve contradictions.
One plan asked me in phase two to confirm that a deleted item returned. Phase four intentionally changed that behavior so the item stayed deleted. When the checklist was grouped by phase, both instructions survived.
The integration PR now gets one flat checklist describing the finished behavior.
When a plan runs cleanly, I come back to one integration PR and one checklist.
When a phase gets stuck, I get a small PR with a specific problem. I can fix it or steer the agent, and the driver continues from there.
This has become the boundary I was looking for. I can leave the loop while the work produces enough evidence to keep moving. I return when the evidence is missing or when the finished feature needs judgment.
Nothing notices a dead session because a session can’t report its own death. A stalled phase can look like a slow one. Rate limits are also invisible. A rate-limited reviewer looks like a gate that is still working.
Those are the next problems in the larger system I’m building around claude-phases . For now, the plugin handles the part that writes the code, opens the PRs, watches the gates, and leaves a trail that another session can recover.
I began by carefully reviewing every AI-generated diff before allowing the work to move forward. That gave me control, but it required my attention at every step.
Now I put that attention at the exceptions and at the final boundary to production. The rest of the system advances only when it can show what happened.
claude-phases is on GitHub at sixoverground/claude-phases . The design notes behind these decisions are in docs/design.md .
If you’re interested in building your own software factory, I set this up for engineering teams. Book a call .
Thanks for reading, and keep making! 🚀
How I Built and Launched a Micro App in 3 Days With the Help Of AI
Incorporating artificial intelligence into my development toolchain helped me to rapidly build ShinyTimer, a macOS menu bar app, and its marketing website. Building software quickly while maintaining quality has always been a challenge, but recent advancements in AI-assisted development have made it possible to go from…
ANSI Art + Chiptunes = ANSITUNES
This article dives into the research and thought process behind the development of my fifth long-form generative work, ANSITUNES on fxHash. My earliest exposure to digital art was viewing ANSI art on bulletin board systems (BBS), a scrolling, text-based experience over a slow dial-up connection. There’s something alluring about…
Makers need time to work on passion projects. Finding a way to fit this between life’s numerous obligations can seem impossible. But it can be done. With a simple, repeatable daily routine.
I write a newsletter sharing my thoughts on computer stuff, life, and being creative. Want to hear from me?

## Original Extract

A field report on misleading green checks, durable agent state, and keeping production behind a human gate. Claude merged the sixth and final phase PR, then went idle. The plan still had one job left: open an integration PR with…

I Let Claude Merge 38 PRs. Here’s What Broke. - Craig Phares
Skip to content
Craig Phares
I Let Claude Merge 38 PRs. Here’s What Broke.
A field report on misleading green checks, durable agent state, and keeping production behind a human gate.
Claude merged the sixth and final phase PR, then went idle.
The plan still had one job left: open an integration PR with a checklist for me to test before anything reached production.
All six phase PRs had merged. The feature branch contained the finished work. CI was green. There was no error and nothing was visibly stuck. From GitHub, the plan looked complete.
The instructions told the agent to open the integration PR, but two details pulled it in another direction. Every transition in the driver pointed to the next action except this one. After the final phase merged, the driver looked for newly unblocked work, found none, and stopped.
I had also described the integration PR as “the only human merge gate.” I meant that a human must merge it. Claude interpreted it to mean that a human must give permission before it could even be opened.
Then it saved that interpretation to the plan.
Every session after that woke up, read the same decision, and treated it as settled. The system had given a temporary mistake durable memory.
This is the kind of failure I kept finding while building claude-phases . The agent usually wrote good code. But knowing whether the surrounding work had actually happened was tricky.
I use Claude Code to work through features in phases. A planner breaks the feature into small steps and writes a plan. A driver implements each step in sequence.
Every phase ends in a PR against a feature branch. Once the required checks pass, the driver merges that PR and starts the next phase. Each phase is small enough to inspect and safe enough to merge without leaving the feature branch broken.
The driver never merges to main .
When all phases are complete, it opens one integration PR from the feature branch to main . That PR contains a human UAT checklist. I test the finished feature and decide whether it ships.
Since the end of July, claude-phases has merged 38 phases across five repositories. It’s written nine plans and run five of them from beginning to end. Four of those phases were claude-phases building itself.
Letting the driver merge its own code worked sooner than I expected. Teaching it how to prove that it was safe to continue took much longer.
The plan has to survive the session
Claude Code cloud sessions are temporary. They pause, compact, restart, and occasionally die in the middle of something.
A new session needs to recover from whatever state the previous session left behind. It might wake up during implementation, after a PR was opened, or after a merge that was never recorded.
I decided that a fresh session could trust two things: the plan file and GitHub. Everything else had to be reconstructed from those.
The plan lives on its own branch. Phase PRs are never allowed to modify it. This means a plan can have one branch recording the state of the work while another branch contains the work itself. It’s a little strange, but it prevents an important failure.
My earlier tool, claude-project-manager , recorded phase status inside the phase PR. That worked while an external timer held the current cursor. It would fail here because an unmerged PR hides its status changes from a new session. The new session would see the phase marked pending and start the same work again.
The driver also records state before taking the action that state describes. It claims the phase before creating its branch. It records the PR before waiting on it.
That ordering closes the gaps where a session can perform an action, die before recording it, and cause the next session to repeat it.
Six green checks, six different meanings
I set up an automated code review workflow on one repository. It produced six successful runs.
Each green check meant something different.
The first run finished in 10 seconds. It skipped the review because the workflow file differed from the copy on the default branch.
The second ran for eight and a half minutes and completed a real review. It didn’t use the --comment flag, so its findings disappeared into the job log.
The third finished in 20 seconds and produced no artifact.
The fourth found a real bug but lacked permission to post it.
The fifth stopped while waiting for a sub-agent.
The sixth found four real bugs. But every gh command it used to report them was denied.
GitHub marked all six runs as successful. None had completed the job I cared about.
A successful workflow run only proves that the workflow exited successfully. If the gate is supposed to represent a code review, I need evidence that a review took place.
claude-phases now supports review artifacts. A check can be required to produce something the reviewer actually wrote before the driver accepts it.
That rule needs configuration because reviewers behave differently. Some always produce a summary. Others only leave inline comments when they find a problem. A clean review from the second kind produces nothing, which is indistinguishable from a reviewer that never ran.
The driver cannot infer the difference. The repository has to define what counts as proof.
I also changed the review workflow so it fails loudly when it can’t complete its review. Skipping gracefully would create another green check with no review behind it.
CI introduced a few more ambiguous states.
A repository with no checks looks exactly like one whose CI stopped running. claude-phases blocks when it sees zero checks unless the repository has been explicitly configured to allow them.
Neutral checks need different handling. GitHub considers them non-failing, and some review tools use neutral specifically to avoid blocking merges. Anthropic’s managed Claude Code Review concludes with a neutral status. A gate that waits only for success will wait forever after the review has already finished.
Required checks can also be impossible to satisfy. A workflow may be configured to run only on PRs targeting main , while each phase PR targets a feature branch. The driver waits for the check by name, but GitHub will never create it.
That failure may not become obvious until hours after the plan was written.
The safest default I found is simple:
A false pass is silent and permanent. A false block is visible, and I can fix it within the phase.
The missing integration PR exposed a different class of problem.
Most of the driver followed an explicit chain. After one action completed, the instructions pointed to the next action. Completion was the exception. It was described as an outcome rather than implemented as a transition.
The driver reached the end of the chain and stopped.
The phrase “human merge gate” added just enough ambiguity for Claude to invent the wrong behavior. Once written into the plan, that interpretation survived every new session.
I rewrote completion as an explicit transition with its own inbound pointer. After the final phase merges, the driver opens the integration PR. Opening it is an agent action. Merging it is a human action.
There is no longer anything to infer from the surrounding language.
I put the checklist where I will see it
My first version put the UAT checklist on the final phase PR.
In automatic mode, the driver would open that PR, wait for its checks, and merge it about a minute later. It created a thorough checklist in a place I would probably never visit.
The plan completed. The checklist existed. And nobody tested it.
I moved the checklist to the integration PR because that is where I re-enter the process.
I also stopped grouping tests by phase. Phase order describes how the software was produced, which is rarely how someone should test it. It can also preserve contradictions.
One plan asked me in phase two to confirm that a deleted item returned. Phase four intentionally changed that behavior so the item stayed deleted. When the checklist was grouped by phase, both instructions survived.
The integration PR now gets one flat checklist describing the finished behavior.
When a plan runs cleanly, I come back to one integration PR and one checklist.
When a phase gets stuck, I get a small PR with a specific problem. I can fix it or steer the agent, and the driver continues from there.
This has become the boundary I was looking for. I can leave the loop while the work produces enough evidence to keep moving. I return when the evidence is missing or when the finished feature needs judgment.
Nothing notices a dead session because a session can’t report its own death. A stalled phase can look like a slow one. Rate limits are also invisible. A rate-limited reviewer looks like a gate that is still working.
Those are the next problems in the larger system I’m building around claude-phases . For now, the plugin handles the part that writes the code, opens the PRs, watches the gates, and leaves a trail that another session can recover.
I began by carefully reviewing every AI-generated diff before allowing the work to move forward. That gave me control, but it required my attention at every step.
Now I put that attention at the exceptions and at the final boundary to production. The rest of the system advances only when it can show what happened.
claude-phases is on GitHub at sixoverground/claude-phases . The design notes behind these decisions are in docs/design.md .
If you’re interested in building your own software factory, I set this up for engineering teams. Book a call .
Thanks for reading, and keep making! 🚀
How I Built and Launched a Micro App in 3 Days With the Help Of AI
Incorporating artificial intelligence into my development toolchain helped me to rapidly build ShinyTimer, a macOS menu bar app, and its marketing website. Building software quickly while maintaining quality has always been a challenge, but recent advancements in AI-assisted development have made it possible to go from…
ANSI Art + Chiptunes = ANSITUNES
This article dives into the research and thought process behind the development of my fifth long-form generative work, ANSITUNES on fxHash. My earliest exposure to digital art was viewing ANSI art on bulletin board systems (BBS), a scrolling, text-based experience over a slow dial-up connection. There’s something alluring about…
Makers need time to work on passion projects. Finding a way to fit this between life’s numerous obligations can seem impossible. But it can be done. With a simple, repeatable daily routine.
I write a newsletter sharing my thoughts on computer stuff, life, and being creative. Want to hear from me?
