---
source: "https://peras.vertias.eu/"
hn_url: "https://news.ycombinator.com/item?id=50180215"
title: "Show HN: Peras – Pause Claude Code and Codex before they burn your weekly limit"
article_title: "Peras: stop Claude Code or Codex before it burns your limit"
image: "https://peras.vertias.eu/og.png"
author: "vertias3u"
captured_at: "2026-10-21T13:30:08Z"
capture_tool: "hn-digest"
hn_id: 50180215
score: 1
comments: 0
posted_at: "2026-10-21T13:01:26Z"
tags:
  - hacker-news
---

# Show HN: Peras – Pause Claude Code and Codex before they burn your weekly limit

- HN: [50180215](https://news.ycombinator.com/item?id=50180215)
- Source: [peras.vertias.eu](https://peras.vertias.eu/)
- Score: 1
- Comments: 0
- Posted: 2026-10-21T13:01:26Z

## Translation

Title: Show HN: Peras – Pause Claude Code and Codex before they burn your weekly limit
Article title: Peras: stop Claude Code or Codex before it burns your limit
Description: Peras watches your Claude Pro/Max or ChatGPT plan usage, tells you where the week is headed, and pauses the agent (or asks first) at the line you set. Free, MIT, local.
HN text: Hello HN! Today I am presenting you with Peras, the tool which governs your Claude Code and Codex usage, and enforces limits. I built it because I got tired of constantly having to spam /usage to check my usage, and decide if a task is worth risking the agent running out mid-task. Peras enforces a limit with 2 main tiers. At 75% weekly usage you will get a notification, and at 90% Peras will automatically pause the agent. These limits are the default and they are customizable to your own liking. You may ask "how does it know my usage in Claude Code or Codex?" Well, it's pretty simple. Peras gets the Claude Code usage % from the status line, and Codex's one comes from its local app-server. The "pause" Peras executes is a hook that refuses the agent's tool calls and ends its turn. A new feature I implemented is one that makes agents aware of their own usage. When you send your prompt, Peras attaches the usage %. It sends it at the start of the session, then again at 50, 75, 90 and 100%, or if a reset is coming in 30 minutes, so the agent can hurry up and actually capitalize on the remaining usage, if any. Just to clarify, this feature is informational and the pause is what enforces the limit. Before I switched to native SwiftUI, I used SwiftBar. The problem was that for the mascots patrolling their usage bars, I had to re-encode every menu image 8 times a second! There was also stuff SwiftBar wasn't capable of, like the menu being able to appear over fullscreen apps, pinning it as a floating panel, etc. What is the difference between the free and paid version? Well, it is MIT, first of all. The CLI, the hooks, the SwiftBar menu, and the app's source are all on GitHub, but building the app needs macOS 13+, Node 18+, "xcode-select --install" and then finally "npm run build:app". The €4.99 buys you a signed copy which is notarized by Apple, with updates included. The guard works with no menu at all. To install, just run: npx peras-cli setup
npx peras-cli codex setup
And now we get to the limits part. For now it's subscription plans only, and when running "claude -p", the pause uses the last interactive reading, so it can come late. Also I feel I have to mention that the guard rail is not a security boundary. It runs inside the agent. The menu is macOS only. The CLI and hooks pass tests on Linux in CI, but I don't use them there day-to-day. Last thing before I end this post, I would really appreciate feedback from you, whether install worked for Codex users, whether you have encountered any bugs, or whatever you feel like sharing.

Article text:
Ver tias
/ Peras
Star on GitHub
75% · line
p e r a s
Stop Claude Code or Codex before it burns the rest of your limit.
Peras watches your plan's usage, tells you where the week is headed, and pauses the agent (or asks first) at the line you set.
Copy
v0.4 · free, MIT, local · Node 18+, no dependencies · peras remove takes the hooks out
The pear sits next to your clock with each agent's 5-hour usage. Click it for the weekly bars, when each window resets, and a pause button for Claude Code and Codex separately.
Clawd and the Codex companion walk along their 5-hour bars, and stand still while their agent is paused.
Press ⌃⌥P to open it from anywhere, full-screen apps included.
Prefer it quiet? Show the pear only, and pick which notifications you get.
The CLI and its guard stay free. Build the app yourself from the source , or use the free SwiftBar menu: peras menubar install
Thu 15:00
Drawn by the Peras app itself, on the same demo week.
At your line, one of three things happens.
A desktop notification. The agent keeps going.
The next tool call waits for your OK, once per session. Even from your phone.
Tool calls are refused and the turn ends, until it lifts or you run peras resume .
Plan the week, not just the stop.
Let an agent use ten percentage points today, with a fresh allowance at local midnight.
Set an estimated weekly allowance for a folder. Parallel account activity can count too; tracking gaps can miss usage.
Weekly charts show observed usage, last-known remaining allowance, recording gaps and retained interventions.
Allowances count observed increases in your weekly percentage. Daily tracking begins with the first reading each day. Reminders require the menu bar and a fresh reading. When one agent pauses, the menu can point out fresh allowance on the other.
One command adds the same hooks to Codex desktop and the Codex CLI. Its own status, lines, pauses, allowances and history work there too: peras codex status , peras codex pause , and the rest.
Codex's lines, history and pauses are kept apart from Claude Code's.
Usage comes from your ChatGPT plan, read through Codex's own local app-server. Peras reads no login files.
Codex has no approval prompt, so an ask line holds its tools until you run peras codex resume .
Everything stays on your machine. The event log holds no prompts and no tool commands.
Ver tias
Spend the last of your week on purpose.

## Original Extract

Peras watches your Claude Pro/Max or ChatGPT plan usage, tells you where the week is headed, and pauses the agent (or asks first) at the line you set. Free, MIT, local.

Hello HN! Today I am presenting you with Peras, the tool which governs your Claude Code and Codex usage, and enforces limits. I built it because I got tired of constantly having to spam /usage to check my usage, and decide if a task is worth risking the agent running out mid-task. Peras enforces a limit with 2 main tiers. At 75% weekly usage you will get a notification, and at 90% Peras will automatically pause the agent. These limits are the default and they are customizable to your own liking. You may ask "how does it know my usage in Claude Code or Codex?" Well, it's pretty simple. Peras gets the Claude Code usage % from the status line, and Codex's one comes from its local app-server. The "pause" Peras executes is a hook that refuses the agent's tool calls and ends its turn. A new feature I implemented is one that makes agents aware of their own usage. When you send your prompt, Peras attaches the usage %. It sends it at the start of the session, then again at 50, 75, 90 and 100%, or if a reset is coming in 30 minutes, so the agent can hurry up and actually capitalize on the remaining usage, if any. Just to clarify, this feature is informational and the pause is what enforces the limit. Before I switched to native SwiftUI, I used SwiftBar. The problem was that for the mascots patrolling their usage bars, I had to re-encode every menu image 8 times a second! There was also stuff SwiftBar wasn't capable of, like the menu being able to appear over fullscreen apps, pinning it as a floating panel, etc. What is the difference between the free and paid version? Well, it is MIT, first of all. The CLI, the hooks, the SwiftBar menu, and the app's source are all on GitHub, but building the app needs macOS 13+, Node 18+, "xcode-select --install" and then finally "npm run build:app". The €4.99 buys you a signed copy which is notarized by Apple, with updates included. The guard works with no menu at all. To install, just run: npx peras-cli setup
npx peras-cli codex setup
And now we get to the limits part. For now it's subscription plans only, and when running "claude -p", the pause uses the last interactive reading, so it can come late. Also I feel I have to mention that the guard rail is not a security boundary. It runs inside the agent. The menu is macOS only. The CLI and hooks pass tests on Linux in CI, but I don't use them there day-to-day. Last thing before I end this post, I would really appreciate feedback from you, whether install worked for Codex users, whether you have encountered any bugs, or whatever you feel like sharing.

Ver tias
/ Peras
Star on GitHub
75% · line
p e r a s
Stop Claude Code or Codex before it burns the rest of your limit.
Peras watches your plan's usage, tells you where the week is headed, and pauses the agent (or asks first) at the line you set.
Copy
v0.4 · free, MIT, local · Node 18+, no dependencies · peras remove takes the hooks out
The pear sits next to your clock with each agent's 5-hour usage. Click it for the weekly bars, when each window resets, and a pause button for Claude Code and Codex separately.
Clawd and the Codex companion walk along their 5-hour bars, and stand still while their agent is paused.
Press ⌃⌥P to open it from anywhere, full-screen apps included.
Prefer it quiet? Show the pear only, and pick which notifications you get.
The CLI and its guard stay free. Build the app yourself from the source , or use the free SwiftBar menu: peras menubar install
Thu 15:00
Drawn by the Peras app itself, on the same demo week.
At your line, one of three things happens.
A desktop notification. The agent keeps going.
The next tool call waits for your OK, once per session. Even from your phone.
Tool calls are refused and the turn ends, until it lifts or you run peras resume .
Plan the week, not just the stop.
Let an agent use ten percentage points today, with a fresh allowance at local midnight.
Set an estimated weekly allowance for a folder. Parallel account activity can count too; tracking gaps can miss usage.
Weekly charts show observed usage, last-known remaining allowance, recording gaps and retained interventions.
Allowances count observed increases in your weekly percentage. Daily tracking begins with the first reading each day. Reminders require the menu bar and a fresh reading. When one agent pauses, the menu can point out fresh allowance on the other.
One command adds the same hooks to Codex desktop and the Codex CLI. Its own status, lines, pauses, allowances and history work there too: peras codex status , peras codex pause , and the rest.
Codex's lines, history and pauses are kept apart from Claude Code's.
Usage comes from your ChatGPT plan, read through Codex's own local app-server. Peras reads no login files.
Codex has no approval prompt, so an ask line holds its tools until you run peras codex resume .
Everything stays on your machine. The event log holds no prompts and no tool commands.
Ver tias
Spend the last of your week on purpose.
