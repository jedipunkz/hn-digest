---
source: "https://jonathannen.com/avoid-claude-ship/"
hn_url: "https://news.ycombinator.com/item?id=50633354"
title: "Reasons to avoid Claude Ship (for now)"
article_title: "Reasons to avoid Claude Ship (for now)"
image: "https://jonathannen.com/images/2026-11-24-avoid-claude-ship-thumb.jpg"
author: "speckx"
captured_at: "2026-11-24T19:58:44Z"
capture_tool: "hn-digest"
hn_id: 50633354
score: 2
comments: 0
posted_at: "2026-11-24T19:36:19Z"
tags:
  - hacker-news
---

# Reasons to avoid Claude Ship (for now)

- HN: [50633354](https://news.ycombinator.com/item?id=50633354)
- Source: [jonathannen.com](https://jonathannen.com/avoid-claude-ship/)
- Score: 2
- Comments: 0
- Posted: 2026-11-24T19:36:19Z

## Translation

Title: Reasons to avoid Claude Ship (for now)
Description: Claude Ship is pretty good and it matches how I already work. I'm still avoiding it for now.

Article text:
<
Reasons to avoid Claude Ship (for now)
One of my general principles is to not treat your harness like a pet: if something better comes along, grab it, refactor your process, and move on. I think Claude Ship , which Anthropic announced this week, bucks the trend. Whilst I've been using a similar approach for some time, I'm holding fast in this case. This post has my four reasons.
Ship is Anthropic's Code Review with two new agents. PR Steward keeps a PR green (failing CI, merge conflicts, review comments), and Claude Approvals reads a policy file in the repo and approves what it allows. In Claude Code it's the /ship command.
My setup has the same shape. One agent rates each new PR, and a cold read checks the diff against its description. The simple ones get "Merge When Green". The rest go to one or two separate agents for review, and Claude or Codex works through the comments. That's 170+ PRs a day, and close to 60% of them merge with nobody reading them.
(1) The SDLC is changing rapidly . The challenge isn't supercharging your current SDLC, it's building a new one specific to the code and the task you have. That's work you should keep for yourself (for now), and the after-the-PR part of it is the bit you should really have your hands all over.
(2) Building your own shouldn't be hard . I could rebuild my Ship equivalent in a day (caveat: the bits of Ship that I need , anyway). Standing it up in your own process is what takes time [1] .
(3) I run more checks . The cold read is one. An agent gets the diff with no description and says what the change does. A second agent grades that against the PR description (how concordant are they), and the grade feeds automerge.
(4) A separate model is still useful today . My coding agent varies. Today I do the coding in Claude, so a Claude review is Claude reading Claude [2] . Copilot is really good at knowing the norms of the codebase. Codex is good at finding security and performance problems, particularly the "P2" ones that might be tolerable but can easily slip the net.
All that said, I want to make sure I'm across what is emerging. And it's clear Ship (or a Ship-shaped capability) will only continue to grow. I've started with using it on ancillary repos (static websites are one example). This blog is one, and I would turn it on here.
Also read: My AI Coding Workflow (Revisited) , The Cold Read , Build Your Own Lightsaber , and Start With One .
Discuss and share via the meta page
. Filed under AI , Code , and 100PR .
Anthropic recommends running the approval policy in observe-only mode first, starting with one repository. Mine started with one patch bump . ↩
If the same model writes the code and checks it, whatever it misread in the task it misreads again in the check. See Views of the Same Thing . ↩

## Original Extract

Claude Ship is pretty good and it matches how I already work. I'm still avoiding it for now.

<
Reasons to avoid Claude Ship (for now)
One of my general principles is to not treat your harness like a pet: if something better comes along, grab it, refactor your process, and move on. I think Claude Ship , which Anthropic announced this week, bucks the trend. Whilst I've been using a similar approach for some time, I'm holding fast in this case. This post has my four reasons.
Ship is Anthropic's Code Review with two new agents. PR Steward keeps a PR green (failing CI, merge conflicts, review comments), and Claude Approvals reads a policy file in the repo and approves what it allows. In Claude Code it's the /ship command.
My setup has the same shape. One agent rates each new PR, and a cold read checks the diff against its description. The simple ones get "Merge When Green". The rest go to one or two separate agents for review, and Claude or Codex works through the comments. That's 170+ PRs a day, and close to 60% of them merge with nobody reading them.
(1) The SDLC is changing rapidly . The challenge isn't supercharging your current SDLC, it's building a new one specific to the code and the task you have. That's work you should keep for yourself (for now), and the after-the-PR part of it is the bit you should really have your hands all over.
(2) Building your own shouldn't be hard . I could rebuild my Ship equivalent in a day (caveat: the bits of Ship that I need , anyway). Standing it up in your own process is what takes time [1] .
(3) I run more checks . The cold read is one. An agent gets the diff with no description and says what the change does. A second agent grades that against the PR description (how concordant are they), and the grade feeds automerge.
(4) A separate model is still useful today . My coding agent varies. Today I do the coding in Claude, so a Claude review is Claude reading Claude [2] . Copilot is really good at knowing the norms of the codebase. Codex is good at finding security and performance problems, particularly the "P2" ones that might be tolerable but can easily slip the net.
All that said, I want to make sure I'm across what is emerging. And it's clear Ship (or a Ship-shaped capability) will only continue to grow. I've started with using it on ancillary repos (static websites are one example). This blog is one, and I would turn it on here.
Also read: My AI Coding Workflow (Revisited) , The Cold Read , Build Your Own Lightsaber , and Start With One .
Discuss and share via the meta page
. Filed under AI , Code , and 100PR .
Anthropic recommends running the approval policy in observe-only mode first, starting with one repository. Mine started with one patch bump . ↩
If the same model writes the code and checks it, whatever it misread in the task it misreads again in the check. See Views of the Same Thing . ↩
