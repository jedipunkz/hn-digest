---
source: "https://www.i-kh.net/p/my-spec-first-workflow-for-ai-coding"
hn_url: "https://news.ycombinator.com/item?id=50194903"
title: "My Spec-First Workflow for AI Coding Agents"
article_title: "My Spec-First Workflow for AI Coding Agents"
image: "https://substackcdn.com/image/fetch/$s_!jP0y!,f_auto,q_auto:best,fl_progressive:steep/https%3A%2F%2Fiourikhramtsov.substack.com%2Ftwitter%2Fsubscribe-card.jpg%3Fv%3D1694042200%26version%3D9"
author: "ajorgensen"
captured_at: "2026-10-22T13:31:30Z"
capture_tool: "hn-digest"
hn_id: 50194903
score: 1
comments: 0
posted_at: "2026-10-22T13:17:10Z"
tags:
  - hacker-news
---

# My Spec-First Workflow for AI Coding Agents

- HN: [50194903](https://news.ycombinator.com/item?id=50194903)
- Source: [www.i-kh.net](https://www.i-kh.net/p/my-spec-first-workflow-for-ai-coding)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T13:17:10Z

## Translation

Title: My Spec-First Workflow for AI Coding Agents
Description: If you’ve used AI agents enough, you may have noticed that for really simple tasks you can just prompt them, and for more complex tasks you may want to generate a plan first.

Article text:
My Spec-First Workflow for AI Coding Agents
Iouri Khramtsov
Subscribe Sign in My Spec-First Workflow for AI Coding Agents
Iouri Khramtsov Aug 28, 2026 1 Share If you’ve used AI agents enough, you may have noticed that for really simple tasks you can just prompt them, and for more complex tasks you may want to generate a plan first.
And if you want the agent to code a complex feature end-to-end, you may want to add another step before planning: a requirements and design document, also called a “spec” file.
For me, once I’ve dialed in the technique, these spec files have been responsible for the single biggest productivity bump I’ve ever had. Looking back at the features and projects I’ve delivered since I started using them, when I compare the old estimates with the actual delivery times, it’s safe to say that my productivity rose around 2x while improving the quality—as observed by the number of bugs found by our manual QA and the number of post-deployment fixes (more on that here ).
But how exactly should you write these spec files to get the most out of them? I’m sure many people have arrived at their own approaches, many of them better than mine, but here’s how I do it.
Note that I use Cursor (at work) and VS Code+Copilot (at home), but in theory this process is very harness-agnostic and should work with most agents.
The first thing I do is copy+paste the template below into the working folder and fill out the obvious parts: the title, the overview sentence, and a rough sketch of the initial requirements. (By the way, if the template looks suspiciously like a miniature version of my method for planning larger projects , that’s not a coincidence).
Note a few blocks of text that are intended strictly for the agents that will be updating this file:
Document rules: deals with a variety of annoying quirks of various models.
High-level design: contains a sentence telling the agent to keep the designs high-level.
Open Questions section: contains rules for updating the open questions.
The rest are the instructions and examples for you, the human driving the AI.
# Title - short description of the project
A statement in a sentence or two describing what you'd like to do.
Example: we'd like to add the ability to send email notifications in our system.
## Document rules
Do not delete this instruction paragraph. Follow it carefully when editing
this document. This file should contain final designs only; do not track
changes or discuss pros/cons of alternatives, or discuss previous
implementation plans unless explicitly prompted. Avoid leaving large
paragraph of text, they're hard to read; find another way to organize
text if necessary. Do not use headings past h3 (`###`), find another way
to organize text if needed. Keep changes brief and to the point - avoid
adding any text that doesn't need to be added. Avoid large blocks of code,
though feel free to add focused blocks of code if it really helps
communicate the design.
All target designs/future designs should use either future tense ("X will
use...") or describe the change to be made ("Update X to..."); present
tense ("X runs only on...") should be used only when describing the
currently implemented designs. Using present sense for future target
states may confuse the reader.
## Requirements
Outline the requirements. Feel free to break them up into subsections
if needed, or can just start in point form and expand later.
Example:
We'll want to send emails on:
- Account creation
- Account deletion
- Manager-initiated account activation
- Manager-initiated account deactivation
New accounts are created through self-serve signup only, and account-creation
emails are sent for those self-serve signups.
For account deletion, send the email for both self-serve deletions and
manager-initiated deletions.
... And so on.
## High-level design
To be filled out later.
This section should contain high-level designs only. Avoid detailed designs,
that's for the planning/implementation agent later.
## Open questions
- **Example open question**: Example open question details.
Do not remove this s
[truncated]
Now is a good time to pause and think through the requirements a bit more, filling them out as much as you can.
Here are some of the things I do next:
Talk to my product manager, the designer, or any other relevant people to refine the edge cases and clarify specific behaviors.
Get mock-ups from the designer, take the screenshots, give those screenshots to my coding agent, and ask the agent to (1) describe what it sees, and (2) add those details to the requirements section. (I usually ask it to describe the changes in terms of any existing components in the codebase as much as possible.)
Grab any third-party docs I need and either reference them in the spec file, or get the agent to summarize them.
Ask the coding agent to review the requirements for any inconsistencies or gaps.
For the last point, there’s a specific prompt that I like to use:
We’re working on `my-spec-file.md`. Do the following:
* Review the file and any relevant documents and code. Think through how you’d implement it. Are there any gaps in the requirements? Any conflicts? Other concerns? Any unexpected interactions with existing functionality? Emit a list of issues. We’re looking only at the issues with the requirements at the moment, not technical design.
* Review each issue from the previous step. For every one that is actually a non-issue, ignore it. For any issue that has a single obvious best answer or any that can be answered with “just do what the code does now”, update `my-spec-file.md` to reflect that. For all others, add them to the “Open questions” section.
In a system of any complexity, this will almost always find something to add to the document. And a smarter model (like say, Anthropic’s Sonnet/Opus, or OpenAI’s Terra/Sol) may also raise some open questions that I haven’t thought of.
Though sometimes they raise dumb questions too, or make wrong edits.
Step 2: Carefully review the updated requirements
If you usually skip reviewing AI’s output, I won’t judge you much (I’ve done that at times), but this is really not the time to do it. At the requirements stage, a wrong assumption by AI will take 2 min to detect and fix. If it’s missed, it could cost you 30 min of rework at best or become a production incident at worst.
Think about it like this: you’re ultimately responsible for these requirements. The AI can help you, but it just cannot know as much about the requirements as you or the humans around you.
I usually do a couple of cycles of prompt for gaps/open questions → review → apply fixes/answer questions → prompt for gaps again.
Once the requirements have stabilized, it’s time to move on to the design.
On simpler features you can just directly prompt for one, but more often than not I do something like this:
First, I put down a few technical decisions that I know upfront: e.g. use a specific library, or a specific third-party service, or that service X should own objects Y. This is also a good time to add sizing constraints: should scale to at least 2 million of {whatever object} across N users, should have p95 latency <500ms, and so on.
Second, I usually have a few open questions of my own, like how to design a certain part, or whether a certain thing is even possible.
Then, I systematically work through various parts using the /tech-design skill below:
---
name: tech-design
description: 'Come up with a technical design for a problem. Use when asked a challenging technical question or to design a solution for one specific problem. Researches best practices, generates options, evaluates tradeoffs, and recommends the best approach. Best for one focused problem — invoke separately for each independent decision.'
---
# Come up with a technical design for a problem
Considers and selects a technical design for one problem.
Best for one specific problem, not when there are multiple
independent decisions to be made. If there are multiple
independent decisions, it's best to use this skill for each decision.
## When to use
- Use this skill for answering a challenging technical question
or coming up with a technical design for one specific problem.
- Avoid using this to solve multiple problems at once. If there
are multiple independent problems/questions, it's best to invoke
this skill for each question/problem separately.
## Instructions
- Read the relevant documentation.
- If there are AGENTS.md, CLAUDE.md, cursor rules, or any other rules
files, review those.
- If there is code to review, review the relevant code.
- Research online to see what other people have done, what the best
practices are, any insights from blogs, research papers, or people
who've done similar things before.
- Think through the problem.
- Come up with the parameters to evaluate the solution on. Examples
could be feasibility, implementation cost, dollar cost, architectural
fit, simplicity/understandability, security, performance, and others.
Don't just take this list, come up with an appropriate list yourself.
- Come up with some design options.
- Consider pros/cons of each design option, tradeoffs, gotchas, and
evaluate it against each of the previously selected evaluation
parameters. Critique each solution.
- Try to improve each design option based on the critique.
- Do a final review/ranking/evaluation of the design options and select
the best one.
[truncated]
/tech-design a solution to this open question from my-spec.md:
/tech-design the best approach to implement (whatever thing)
You can find additional context in (list of files)
Note that this skill doesn’t mention anything about modifying the spec document; in fact, it explicitly asks not to make changes. This is because I prefer to review the decision and, if I’m not happy with the result, re-prompt with additional information or ask follow-up questions and refine the design.
When I’m happy with the result, I explicitly ask:
Update my-spec-file.md to reflect the recommended approach.
Aside from the usual human review, I also like to run this variant of the review prompt a couple times before settling on the design:
We’re about to start implementing my-spec-file.md. Before doing that, I’d like to do a review of the technical design. Do the following:
* Review the file and any relevant documents and code. Think through how you’d implement it, in detail. Do you have any concerns? Are there any gaps in the requirements or design? Any conflicts? Any remaining significant decisions to be made? Any unexpected interactions with existing functionality? Emit a list of issues.
* Review each issue from the previous step. For every one that is actually a non-issue, ignore it. For any issue that has a single obvious best answer or any that can be answered with “just do what the code does now”, update my-spec-file.md to reflect that. For all others, add them to the “Open questions” section.
This usually produces a number of edits and follow-up questions, which I review, correct, and answer. Sometimes I can answer the open questions immediately, and sometimes I invoke the /tech-design skill to answer them.
Sometimes these reviews reveal new missed requirements or unexpected interactions that force me to step back and rethink the approach.
I usually do a couple of these review passes before moving on to the next step.
Like what you’re reading? Subscribe for more
At this point, the design is largely complete and now we’re moving on to the implementation.
For simpler features, we can, again, just tell the agent to plan and implement. However, for larger features I find it useful to break up the work into implementation phases.
For features spanning multiple repos, I can ask the agent to figure out which PRs should be created in what order. Some models can be overzealous in splitting the PRs, so I might ask to bundle the changes in as few PRs as practically possible.
For a large implementation in a single repo, I ask to break up the changes into phases con

[truncated]

## Original Extract

If you’ve used AI agents enough, you may have noticed that for really simple tasks you can just prompt them, and for more complex tasks you may want to generate a plan first.

My Spec-First Workflow for AI Coding Agents
Iouri Khramtsov
Subscribe Sign in My Spec-First Workflow for AI Coding Agents
Iouri Khramtsov Aug 28, 2026 1 Share If you’ve used AI agents enough, you may have noticed that for really simple tasks you can just prompt them, and for more complex tasks you may want to generate a plan first.
And if you want the agent to code a complex feature end-to-end, you may want to add another step before planning: a requirements and design document, also called a “spec” file.
For me, once I’ve dialed in the technique, these spec files have been responsible for the single biggest productivity bump I’ve ever had. Looking back at the features and projects I’ve delivered since I started using them, when I compare the old estimates with the actual delivery times, it’s safe to say that my productivity rose around 2x while improving the quality—as observed by the number of bugs found by our manual QA and the number of post-deployment fixes (more on that here ).
But how exactly should you write these spec files to get the most out of them? I’m sure many people have arrived at their own approaches, many of them better than mine, but here’s how I do it.
Note that I use Cursor (at work) and VS Code+Copilot (at home), but in theory this process is very harness-agnostic and should work with most agents.
The first thing I do is copy+paste the template below into the working folder and fill out the obvious parts: the title, the overview sentence, and a rough sketch of the initial requirements. (By the way, if the template looks suspiciously like a miniature version of my method for planning larger projects , that’s not a coincidence).
Note a few blocks of text that are intended strictly for the agents that will be updating this file:
Document rules: deals with a variety of annoying quirks of various models.
High-level design: contains a sentence telling the agent to keep the designs high-level.
Open Questions section: contains rules for updating the open questions.
The rest are the instructions and examples for you, the human driving the AI.
# Title - short description of the project
A statement in a sentence or two describing what you'd like to do.
Example: we'd like to add the ability to send email notifications in our system.
## Document rules
Do not delete this instruction paragraph. Follow it carefully when editing
this document. This file should contain final designs only; do not track
changes or discuss pros/cons of alternatives, or discuss previous
implementation plans unless explicitly prompted. Avoid leaving large
paragraph of text, they're hard to read; find another way to organize
text if necessary. Do not use headings past h3 (`###`), find another way
to organize text if needed. Keep changes brief and to the point - avoid
adding any text that doesn't need to be added. Avoid large blocks of code,
though feel free to add focused blocks of code if it really helps
communicate the design.
All target designs/future designs should use either future tense ("X will
use...") or describe the change to be made ("Update X to..."); present
tense ("X runs only on...") should be used only when describing the
currently implemented designs. Using present sense for future target
states may confuse the reader.
## Requirements
Outline the requirements. Feel free to break them up into subsections
if needed, or can just start in point form and expand later.
Example:
We'll want to send emails on:
- Account creation
- Account deletion
- Manager-initiated account activation
- Manager-initiated account deactivation
New accounts are created through self-serve signup only, and account-creation
emails are sent for those self-serve signups.
For account deletion, send the email for both self-serve deletions and
manager-initiated deletions.
... And so on.
## High-level design
To be filled out later.
This section should contain high-level designs only. Avoid detailed designs,
that's for the planning/implementation agent later.
## Open questions
- **Example open question**: Example open question details.
Do not remove this s
[truncated]
Now is a good time to pause and think through the requirements a bit more, filling them out as much as you can.
Here are some of the things I do next:
Talk to my product manager, the designer, or any other relevant people to refine the edge cases and clarify specific behaviors.
Get mock-ups from the designer, take the screenshots, give those screenshots to my coding agent, and ask the agent to (1) describe what it sees, and (2) add those details to the requirements section. (I usually ask it to describe the changes in terms of any existing components in the codebase as much as possible.)
Grab any third-party docs I need and either reference them in the spec file, or get the agent to summarize them.
Ask the coding agent to review the requirements for any inconsistencies or gaps.
For the last point, there’s a specific prompt that I like to use:
We’re working on `my-spec-file.md`. Do the following:
* Review the file and any relevant documents and code. Think through how you’d implement it. Are there any gaps in the requirements? Any conflicts? Other concerns? Any unexpected interactions with existing functionality? Emit a list of issues. We’re looking only at the issues with the requirements at the moment, not technical design.
* Review each issue from the previous step. For every one that is actually a non-issue, ignore it. For any issue that has a single obvious best answer or any that can be answered with “just do what the code does now”, update `my-spec-file.md` to reflect that. For all others, add them to the “Open questions” section.
In a system of any complexity, this will almost always find something to add to the document. And a smarter model (like say, Anthropic’s Sonnet/Opus, or OpenAI’s Terra/Sol) may also raise some open questions that I haven’t thought of.
Though sometimes they raise dumb questions too, or make wrong edits.
Step 2: Carefully review the updated requirements
If you usually skip reviewing AI’s output, I won’t judge you much (I’ve done that at times), but this is really not the time to do it. At the requirements stage, a wrong assumption by AI will take 2 min to detect and fix. If it’s missed, it could cost you 30 min of rework at best or become a production incident at worst.
Think about it like this: you’re ultimately responsible for these requirements. The AI can help you, but it just cannot know as much about the requirements as you or the humans around you.
I usually do a couple of cycles of prompt for gaps/open questions → review → apply fixes/answer questions → prompt for gaps again.
Once the requirements have stabilized, it’s time to move on to the design.
On simpler features you can just directly prompt for one, but more often than not I do something like this:
First, I put down a few technical decisions that I know upfront: e.g. use a specific library, or a specific third-party service, or that service X should own objects Y. This is also a good time to add sizing constraints: should scale to at least 2 million of {whatever object} across N users, should have p95 latency <500ms, and so on.
Second, I usually have a few open questions of my own, like how to design a certain part, or whether a certain thing is even possible.
Then, I systematically work through various parts using the /tech-design skill below:
---
name: tech-design
description: 'Come up with a technical design for a problem. Use when asked a challenging technical question or to design a solution for one specific problem. Researches best practices, generates options, evaluates tradeoffs, and recommends the best approach. Best for one focused problem — invoke separately for each independent decision.'
---
# Come up with a technical design for a problem
Considers and selects a technical design for one problem.
Best for one specific problem, not when there are multiple
independent decisions to be made. If there are multiple
independent decisions, it's best to use this skill for each decision.
## When to use
- Use this skill for answering a challenging technical question
or coming up with a technical design for one specific problem.
- Avoid using this to solve multiple problems at once. If there
are multiple independent problems/questions, it's best to invoke
this skill for each question/problem separately.
## Instructions
- Read the relevant documentation.
- If there are AGENTS.md, CLAUDE.md, cursor rules, or any other rules
files, review those.
- If there is code to review, review the relevant code.
- Research online to see what other people have done, what the best
practices are, any insights from blogs, research papers, or people
who've done similar things before.
- Think through the problem.
- Come up with the parameters to evaluate the solution on. Examples
could be feasibility, implementation cost, dollar cost, architectural
fit, simplicity/understandability, security, performance, and others.
Don't just take this list, come up with an appropriate list yourself.
- Come up with some design options.
- Consider pros/cons of each design option, tradeoffs, gotchas, and
evaluate it against each of the previously selected evaluation
parameters. Critique each solution.
- Try to improve each design option based on the critique.
- Do a final review/ranking/evaluation of the design options and select
the best one.
[truncated]
/tech-design a solution to this open question from my-spec.md:
/tech-design the best approach to implement (whatever thing)
You can find additional context in (list of files)
Note that this skill doesn’t mention anything about modifying the spec document; in fact, it explicitly asks not to make changes. This is because I prefer to review the decision and, if I’m not happy with the result, re-prompt with additional information or ask follow-up questions and refine the design.
When I’m happy with the result, I explicitly ask:
Update my-spec-file.md to reflect the recommended approach.
Aside from the usual human review, I also like to run this variant of the review prompt a couple times before settling on the design:
We’re about to start implementing my-spec-file.md. Before doing that, I’d like to do a review of the technical design. Do the following:
* Review the file and any relevant documents and code. Think through how you’d implement it, in detail. Do you have any concerns? Are there any gaps in the requirements or design? Any conflicts? Any remaining significant decisions to be made? Any unexpected interactions with existing functionality? Emit a list of issues.
* Review each issue from the previous step. For every one that is actually a non-issue, ignore it. For any issue that has a single obvious best answer or any that can be answered with “just do what the code does now”, update my-spec-file.md to reflect that. For all others, add them to the “Open questions” section.
This usually produces a number of edits and follow-up questions, which I review, correct, and answer. Sometimes I can answer the open questions immediately, and sometimes I invoke the /tech-design skill to answer them.
Sometimes these reviews reveal new missed requirements or unexpected interactions that force me to step back and rethink the approach.
I usually do a couple of these review passes before moving on to the next step.
Like what you’re reading? Subscribe for more
At this point, the design is largely complete and now we’re moving on to the implementation.
For simpler features, we can, again, just tell the agent to plan and implement. However, for larger features I find it useful to break up the work into implementation phases.
For features spanning multiple repos, I can ask the agent to figure out which PRs should be created in what order. Some models can be overzealous in splitting the PRs, so I might ask to bundle the changes in as few PRs as practically possible.
For a large implementation in a single repo, I ask to break up the changes into phases con

[truncated]
