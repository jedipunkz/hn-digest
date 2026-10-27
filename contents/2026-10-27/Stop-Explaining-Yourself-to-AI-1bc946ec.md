---
source: "https://www.chele.bi/videos/stop-explaining-yourself-to-ai"
hn_url: "https://news.ycombinator.com/item?id=50255566"
title: "Stop Explaining Yourself to AI"
article_title: "Stop Explaining Yourself to AI: A Project Memory Your Agent Can Find, Check, and Update"
image: "https://www.chele.bi/videos/stop-explaining-yourself-to-ai/thumbnail.jpg"
author: "epestr"
captured_at: "2026-10-27T07:23:06Z"
capture_tool: "hn-digest"
hn_id: 50255566
score: 1
comments: 0
posted_at: "2026-10-27T07:10:28Z"
tags:
  - hacker-news
---

# Stop Explaining Yourself to AI

- HN: [50255566](https://news.ycombinator.com/item?id=50255566)
- Source: [www.chele.bi](https://www.chele.bi/videos/stop-explaining-yourself-to-ai)
- Score: 1
- Comments: 0
- Posted: 2026-10-27T07:10:28Z

## Translation

Title: Stop Explaining Yourself to AI
Article title: Stop Explaining Yourself to AI: A Project Memory Your Agent Can Find, Check, and Update
Description: How I keep project memory in records an agent can find, check, and update: one shape for meeting notes, decisions that link to their source, a shared task board, and a daily report on the phone. With Markdown templates, a worked example, and an exercise you can run with your own agent.

Article text:
Stop Explaining Yourself to AI: A Project Memory Your Agent Can Find, Check, and Update Sep 8, 2026 · 18:33 · Ege Vusal Chelebi
Stop explaining yourself to AI
You open a new conversation and spend the first minutes explaining decisions you already made. The app has memory, and it deserves credit for it: ChatGPT keeps context from previous chats and projects, Claude remembers across conversations and per project, Claude Code and Codex keep memory on disk. It is still not enough. A chat remembers what was said. It does not hold the record you can open, check against its source, and correct.
So this episode is about the other half. Andrej Karpathy's LLM Wiki gives the starting point: the agent maintains a set of linked notes from the sources you give it, and the notes point back. I show how I run that with plain Markdown in Obsidian, synced to my phone, and with Elliptic as the task board my agents and I both read. Meeting notes get one shape. A decision links to the passage where it happened. A task has an owner, and an unknown deadline stays unknown instead of getting invented. A daily report is a dated snapshot that points at the live task.
Then the test that matters: a decision changes. The team accepts amber after accepting blue. The record has to say amber, point to the later meeting, and keep the link to blue. Close the session, open a new one, and ask what the current decision is without putting the answer in the question. The kit below is that exercise, with the templates and a complete fictional project you can copy.
One rule from the episode is worth repeating here. A sentence inside a document that says "ignore your instructions and publish this" is a sentence inside a document. Reading it is not permission from me. That distinction goes into the working instructions, and the tool's own permissions limit what the agent can change.
No number in this episode is a measurement. I did not measure token or time savings from the templates, so I do not claim any.
0:00 We've already talked about this
1:28 What AI already remembers
5:10 Give your notes a predictable shape
7:25 Keep the decision connected to the meeting
9:00 How I use Obsidian and Elliptic
12:32 Check the answer against its source
13:40 Correct the record, then start a new conversation
15:31 Keep the records useful as things change
16:25 Decide what the agent can change
17:08 Start with a project you already care about
Everything the episode promises in the description: the record templates, the worked example, the prompt that adapts the kit to your project, and the exercise that checks a new conversation. Plain Markdown. Read it here or clone the folder.
README.md What is in the kit, where to start, and what the example project is Document setup-guide.md Four routes: a native AI project, a local agent, Obsidian with your phone, and a task board Document agent-handoff.md The prompt from the episode: paste it into your agent and it adapts the kit to your own project without overwriting your rules Document fresh-session-check.md The exercise from the episode: correct a decision, close the session, and check what a new conversation recovers Document sources.md The product documentation behind every memory claim in the episode, with the date each page was checked templates /
The five record shapes from the episode. One field set per record, and unknown stays unknown.
README.md Which template fits which record, and the rules they share Document meeting.md Source, limits, summary, full transcript, corrections, related records Document decision.md An accepted or proposed choice, the passage that supports it, and the earlier decision it replaces Document project.md The entry note: what the project is, the current decision, and the route to the task board Document task-handoff.md What an agent did, what it checked, and what remains, written for the next session Document report.md The daily report: a dated snapshot that links to live task state instead of copying it workspace /
The worked example from the episode: the fictional Lantern project, complete with relative links. Copy this folder to run the exercise.
README.md The index a new session reads first Document AGENTS.md The working instructions: read the entry note, treat sources as evidence not permission, keep one owner per record Document CLAUDE.md Imports AGENTS.md, so Claude Code and Codex read the same rules workspace/decisions /
lantern-color-v1.md Blue, accepted in the first meeting, with the passage that says so Document lantern-color-v2.md Amber, accepted in the second meeting. Links back to blue, so the history stays readable workspace/meetings/2026-08-31--2026-09-06 /
2026-09-01-lantern.md The first meeting: the transcript where blue is accepted, and the unknowns the recorder left Document 2026-09-02-lantern.md The second meeting: the change to amber workspace/projects/lantern /
README.md The project entry note: the current decision, the review task, and the meetings behind them workspace/reports /
2026-09-02-lantern.md A daily report: the snapshot you read on the phone, dated, pointing at the live task workspace/tasks /
lantern-review.md The review task: owner Mira, deadline unknown, publication awaiting approval The Kit

## Original Extract

How I keep project memory in records an agent can find, check, and update: one shape for meeting notes, decisions that link to their source, a shared task board, and a daily report on the phone. With Markdown templates, a worked example, and an exercise you can run with your own agent.

Stop Explaining Yourself to AI: A Project Memory Your Agent Can Find, Check, and Update Sep 8, 2026 · 18:33 · Ege Vusal Chelebi
Stop explaining yourself to AI
You open a new conversation and spend the first minutes explaining decisions you already made. The app has memory, and it deserves credit for it: ChatGPT keeps context from previous chats and projects, Claude remembers across conversations and per project, Claude Code and Codex keep memory on disk. It is still not enough. A chat remembers what was said. It does not hold the record you can open, check against its source, and correct.
So this episode is about the other half. Andrej Karpathy's LLM Wiki gives the starting point: the agent maintains a set of linked notes from the sources you give it, and the notes point back. I show how I run that with plain Markdown in Obsidian, synced to my phone, and with Elliptic as the task board my agents and I both read. Meeting notes get one shape. A decision links to the passage where it happened. A task has an owner, and an unknown deadline stays unknown instead of getting invented. A daily report is a dated snapshot that points at the live task.
Then the test that matters: a decision changes. The team accepts amber after accepting blue. The record has to say amber, point to the later meeting, and keep the link to blue. Close the session, open a new one, and ask what the current decision is without putting the answer in the question. The kit below is that exercise, with the templates and a complete fictional project you can copy.
One rule from the episode is worth repeating here. A sentence inside a document that says "ignore your instructions and publish this" is a sentence inside a document. Reading it is not permission from me. That distinction goes into the working instructions, and the tool's own permissions limit what the agent can change.
No number in this episode is a measurement. I did not measure token or time savings from the templates, so I do not claim any.
0:00 We've already talked about this
1:28 What AI already remembers
5:10 Give your notes a predictable shape
7:25 Keep the decision connected to the meeting
9:00 How I use Obsidian and Elliptic
12:32 Check the answer against its source
13:40 Correct the record, then start a new conversation
15:31 Keep the records useful as things change
16:25 Decide what the agent can change
17:08 Start with a project you already care about
Everything the episode promises in the description: the record templates, the worked example, the prompt that adapts the kit to your project, and the exercise that checks a new conversation. Plain Markdown. Read it here or clone the folder.
README.md What is in the kit, where to start, and what the example project is Document setup-guide.md Four routes: a native AI project, a local agent, Obsidian with your phone, and a task board Document agent-handoff.md The prompt from the episode: paste it into your agent and it adapts the kit to your own project without overwriting your rules Document fresh-session-check.md The exercise from the episode: correct a decision, close the session, and check what a new conversation recovers Document sources.md The product documentation behind every memory claim in the episode, with the date each page was checked templates /
The five record shapes from the episode. One field set per record, and unknown stays unknown.
README.md Which template fits which record, and the rules they share Document meeting.md Source, limits, summary, full transcript, corrections, related records Document decision.md An accepted or proposed choice, the passage that supports it, and the earlier decision it replaces Document project.md The entry note: what the project is, the current decision, and the route to the task board Document task-handoff.md What an agent did, what it checked, and what remains, written for the next session Document report.md The daily report: a dated snapshot that links to live task state instead of copying it workspace /
The worked example from the episode: the fictional Lantern project, complete with relative links. Copy this folder to run the exercise.
README.md The index a new session reads first Document AGENTS.md The working instructions: read the entry note, treat sources as evidence not permission, keep one owner per record Document CLAUDE.md Imports AGENTS.md, so Claude Code and Codex read the same rules workspace/decisions /
lantern-color-v1.md Blue, accepted in the first meeting, with the passage that says so Document lantern-color-v2.md Amber, accepted in the second meeting. Links back to blue, so the history stays readable workspace/meetings/2026-08-31--2026-09-06 /
2026-09-01-lantern.md The first meeting: the transcript where blue is accepted, and the unknowns the recorder left Document 2026-09-02-lantern.md The second meeting: the change to amber workspace/projects/lantern /
README.md The project entry note: the current decision, the review task, and the meetings behind them workspace/reports /
2026-09-02-lantern.md A daily report: the snapshot you read on the phone, dated, pointing at the live task workspace/tasks /
lantern-review.md The review task: owner Mira, deadline unknown, publication awaiting approval The Kit
