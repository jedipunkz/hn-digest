---
source: "https://explyt.ai/docs/explyt-test/whats-new-explyt"
hn_url: "https://news.ycombinator.com/item?id=49841139"
title: "Explyt 5.20 – background task panel for AI agent runs in JetBrains IDEs"
article_title: "Explyt 5.20: stop waiting in the chat — a background task panel and notifications outside the IDE | Explyt Tools"
image: "https://explyt.ai/docs/img/whats-new/background_tasks_panel.png"
author: "Explyt-ai"
captured_at: "2026-09-25T07:11:28Z"
capture_tool: "hn-digest"
hn_id: 49841139
score: 1
comments: 0
posted_at: "2026-09-25T07:07:26Z"
tags:
  - hacker-news
---

# Explyt 5.20 – background task panel for AI agent runs in JetBrains IDEs

- HN: [49841139](https://news.ycombinator.com/item?id=49841139)
- Source: [explyt.ai](https://explyt.ai/docs/explyt-test/whats-new-explyt)
- Score: 1
- Comments: 0
- Posted: 2026-09-25T07:07:26Z

## Translation

Title: Explyt 5.20 – background task panel for AI agent runs in JetBrains IDEs
Article title: Explyt 5.20: stop waiting in the chat — a background task panel and notifications outside the IDE | Explyt Tools
Description: Explyt 5.20 collects terminals, run configurations, and subagents in one background task panel, adds system notifications, frees Memory Bank space in two stages, and shortens interface delays.

Article text:
Skip to main content Explyt Docs Getting started Reference What's new Explyt Spring Try Explyt GitHub On this page Explyt 5.20: stop waiting in the chat — a background task panel and notifications outside the IDE
Part of the agent's work outlives a single response: a terminal command, a run configuration with tests, a subagent with its own investigation. Each of these processes gets its own card in the chat and keeps running while the agent answers your next messages. To find out what was still running, you had to scroll through the chat history and check the cards one by one, and the result of a long task only became visible once you came back to the IDE.
Explyt 5.20 collects active processes in one panel next to the chat input: the background task indicator opens a list of terminals, run configurations, and subagents with their status and elapsed time, and from the panel you can jump to a card or stop a process. When the IDE is inactive, Explyt can send system notifications: a response is ready, the agent asked a question, is waiting for confirmation, or has prepared a review report.
The rest of the release removes smaller but frequent interruptions: Memory Bank frees space for new entries in two stages, the inline code generation window receives focus right away, long subagent results are no longer clipped, and password-storage reads have left the IDE interface thread.
Memory Bank keeps saving knowledge when full
Inline code generation and subagent results
After updating from a previous version
Background tasks in one panel ​
A background task indicator now appears next to the message queue in the chat input. It shows work that continues independently of the current response:
Click the indicator to open the panel: it shows each task's status and elapsed time. From the panel you can jump to the original card in the chat or stop a process you no longer need.
The panel is most useful when the agent is running tests, building the project, and delegating a separate investigation to a subagent at the same time. Previously you had to scroll through the chat and find each card. Active work is now collected next to the input, so before closing the IDE or switching branches it is easier to confirm that nothing important is still running.
Stopping a background task interrupts only the selected process and does not undo the actions it has already completed. Check its output before running it again: the command may already have changed files or produced a partial result. The panel, jumping to a card, and stopping processes are described on the Background tasks page.
Explyt can mirror important events through macOS, Windows, and Linux system notifications. A notification appears when the IDE is inactive and the agent:
requested confirmation for an action;
This fits a long build, an analysis of a large project, or a review during which you switch to a browser, a terminal, or another IDE. You no longer need to return every few minutes just to check the chat.
The feature is off by default. Enable it under Settings | Explyt | Notifications . Your operating system must also allow notifications from the IDE. If notifications do not appear, check both settings and the Do Not Disturb mode.
These notifications are unrelated to the Memory notifications setting in the Memory Bank section: that one shows the results of automatic memory extraction inside the IDE.
Memory Bank keeps saving knowledge when full ​
Memory Bank preserves knowledge across chats, but the number of entries is limited to 200. Previously, when the bank was full, automatic cleanup removed only entries that had not been used for more than 30 days. An active project might have no such entries, so new knowledge was rejected even though automatic cleanup was enabled.
Space is now freed in two stages. First, Explyt still removes entries that have not been used for more than 30 days. If that is not enough, Explyt may remove Project , Reference , and Feedback entries that have not been used for more than 7 days. User entries and the entry that was just saved are protected from automatic removal.
This helps long-lived projects where the agent regularly records decisions and the constraints it discovers. The bank no longer stops accepting new knowledge simply because every existing entry was used relatively recently.
The second stage has a downside: a useful entry can also be evicted after 7 days without use. Do not save temporary task state, file contents, or facts that are easy to recover from the repository: they take up limited space and cause more valuable entries to be removed sooner.
Inline code generation and subagent results ​
The request window in inline code generation ( Explyt: Generate Code ) now receives focus as soon as it opens. You can start typing the request without an extra mouse click. While focus stays in the request window, Undo reverts the text of the request itself; the latest change to the source file stays in place. The difference is most noticeable in a keyboard-only workflow: open generation, type an instruction, fix it, and submit without risking an accidental code rollback.
The rendering of long subagent results is fixed as well. Panels that have not received a height yet are now measured using the width of their parent panel. As a result, Markdown with tables and long lines is no longer clipped until the next window resize.
The interface no longer waits for PasswordSafe ​
Provider availability checks and model selection no longer read PasswordSafe synchronously on the interface thread. Identical concurrent provider checks are combined into one. This reduces the chance of brief freezes when the IDE starts, a chat opens, or you switch models, especially when the system key store takes time to respond.
Chat history is analyzed once ​
The internal analyzer now extracts chat histories once per run and reuses the result across modes and open projects. Previously, the same set of files could be reread in full several times. In projects with a large chat history, this reduces repeated background work and leaves background threads available for user tasks.
Batch-processing dependencies updated ​
The batch-processing service no longer includes vulnerable dependencies, its system and Java packages have been updated, and the container build no longer leaves obsolete JAR files behind. The change requires no plugin configuration, but it reduces the surface of known vulnerabilities in the server environment.
After updating from a previous version ​
Existing chats and settings keep working. For the new features, check the following:
To receive notifications outside the IDE, enable them under Settings | Explyt | Notifications and allow notifications from the IDE in your operating system settings.
Before stopping a background command, open its card and check whether it has already changed any files.
If Memory Bank often becomes full, remove temporary knowledge and keep only facts that cannot be recovered quickly from the project. Project , Reference , and Feedback entries unused for more than 7 days can now be removed automatically.
Open Settings | Plugins | Installed , find the Explyt plugin, and click Update if an update is available. Restart the IDE after installation if prompted.
If the Explyt plugin is not installed yet, download it from the Explyt download page and follow the installation instructions.
Memory Bank keeps saving knowledge when full
Inline code generation and subagent results
Performance and security The interface no longer waits for PasswordSafe
Batch-processing dependencies updated
After updating from a previous version

## Original Extract

Explyt 5.20 collects terminals, run configurations, and subagents in one background task panel, adds system notifications, frees Memory Bank space in two stages, and shortens interface delays.

Skip to main content Explyt Docs Getting started Reference What's new Explyt Spring Try Explyt GitHub On this page Explyt 5.20: stop waiting in the chat — a background task panel and notifications outside the IDE
Part of the agent's work outlives a single response: a terminal command, a run configuration with tests, a subagent with its own investigation. Each of these processes gets its own card in the chat and keeps running while the agent answers your next messages. To find out what was still running, you had to scroll through the chat history and check the cards one by one, and the result of a long task only became visible once you came back to the IDE.
Explyt 5.20 collects active processes in one panel next to the chat input: the background task indicator opens a list of terminals, run configurations, and subagents with their status and elapsed time, and from the panel you can jump to a card or stop a process. When the IDE is inactive, Explyt can send system notifications: a response is ready, the agent asked a question, is waiting for confirmation, or has prepared a review report.
The rest of the release removes smaller but frequent interruptions: Memory Bank frees space for new entries in two stages, the inline code generation window receives focus right away, long subagent results are no longer clipped, and password-storage reads have left the IDE interface thread.
Memory Bank keeps saving knowledge when full
Inline code generation and subagent results
After updating from a previous version
Background tasks in one panel ​
A background task indicator now appears next to the message queue in the chat input. It shows work that continues independently of the current response:
Click the indicator to open the panel: it shows each task's status and elapsed time. From the panel you can jump to the original card in the chat or stop a process you no longer need.
The panel is most useful when the agent is running tests, building the project, and delegating a separate investigation to a subagent at the same time. Previously you had to scroll through the chat and find each card. Active work is now collected next to the input, so before closing the IDE or switching branches it is easier to confirm that nothing important is still running.
Stopping a background task interrupts only the selected process and does not undo the actions it has already completed. Check its output before running it again: the command may already have changed files or produced a partial result. The panel, jumping to a card, and stopping processes are described on the Background tasks page.
Explyt can mirror important events through macOS, Windows, and Linux system notifications. A notification appears when the IDE is inactive and the agent:
requested confirmation for an action;
This fits a long build, an analysis of a large project, or a review during which you switch to a browser, a terminal, or another IDE. You no longer need to return every few minutes just to check the chat.
The feature is off by default. Enable it under Settings | Explyt | Notifications . Your operating system must also allow notifications from the IDE. If notifications do not appear, check both settings and the Do Not Disturb mode.
These notifications are unrelated to the Memory notifications setting in the Memory Bank section: that one shows the results of automatic memory extraction inside the IDE.
Memory Bank keeps saving knowledge when full ​
Memory Bank preserves knowledge across chats, but the number of entries is limited to 200. Previously, when the bank was full, automatic cleanup removed only entries that had not been used for more than 30 days. An active project might have no such entries, so new knowledge was rejected even though automatic cleanup was enabled.
Space is now freed in two stages. First, Explyt still removes entries that have not been used for more than 30 days. If that is not enough, Explyt may remove Project , Reference , and Feedback entries that have not been used for more than 7 days. User entries and the entry that was just saved are protected from automatic removal.
This helps long-lived projects where the agent regularly records decisions and the constraints it discovers. The bank no longer stops accepting new knowledge simply because every existing entry was used relatively recently.
The second stage has a downside: a useful entry can also be evicted after 7 days without use. Do not save temporary task state, file contents, or facts that are easy to recover from the repository: they take up limited space and cause more valuable entries to be removed sooner.
Inline code generation and subagent results ​
The request window in inline code generation ( Explyt: Generate Code ) now receives focus as soon as it opens. You can start typing the request without an extra mouse click. While focus stays in the request window, Undo reverts the text of the request itself; the latest change to the source file stays in place. The difference is most noticeable in a keyboard-only workflow: open generation, type an instruction, fix it, and submit without risking an accidental code rollback.
The rendering of long subagent results is fixed as well. Panels that have not received a height yet are now measured using the width of their parent panel. As a result, Markdown with tables and long lines is no longer clipped until the next window resize.
The interface no longer waits for PasswordSafe ​
Provider availability checks and model selection no longer read PasswordSafe synchronously on the interface thread. Identical concurrent provider checks are combined into one. This reduces the chance of brief freezes when the IDE starts, a chat opens, or you switch models, especially when the system key store takes time to respond.
Chat history is analyzed once ​
The internal analyzer now extracts chat histories once per run and reuses the result across modes and open projects. Previously, the same set of files could be reread in full several times. In projects with a large chat history, this reduces repeated background work and leaves background threads available for user tasks.
Batch-processing dependencies updated ​
The batch-processing service no longer includes vulnerable dependencies, its system and Java packages have been updated, and the container build no longer leaves obsolete JAR files behind. The change requires no plugin configuration, but it reduces the surface of known vulnerabilities in the server environment.
After updating from a previous version ​
Existing chats and settings keep working. For the new features, check the following:
To receive notifications outside the IDE, enable them under Settings | Explyt | Notifications and allow notifications from the IDE in your operating system settings.
Before stopping a background command, open its card and check whether it has already changed any files.
If Memory Bank often becomes full, remove temporary knowledge and keep only facts that cannot be recovered quickly from the project. Project , Reference , and Feedback entries unused for more than 7 days can now be removed automatically.
Open Settings | Plugins | Installed , find the Explyt plugin, and click Update if an update is available. Restart the IDE after installation if prompted.
If the Explyt plugin is not installed yet, download it from the Explyt download page and follow the installation instructions.
Memory Bank keeps saving knowledge when full
Inline code generation and subagent results
Performance and security The interface no longer waits for PasswordSafe
Batch-processing dependencies updated
After updating from a previous version
