---
source: "https://jyoung.dev/blog/claude-code-parallel-sessions-memory/"
hn_url: "https://news.ycombinator.com/item?id=50166803"
title: "Claude Code Memory Usage: How Many Sessions Fit on One Machine"
article_title: "Claude Code Memory Usage: How Many Sessions Fit on One Machine — John Young"
image: "https://jyoung.dev/og/claude-code-parallel-sessions-memory.png"
author: "spreadstreet"
captured_at: "2026-10-20T14:44:09Z"
capture_tool: "hn-digest"
hn_id: 50166803
score: 1
comments: 0
posted_at: "2026-10-20T14:14:54Z"
tags:
  - hacker-news
---

# Claude Code Memory Usage: How Many Sessions Fit on One Machine

- HN: [50166803](https://news.ycombinator.com/item?id=50166803)
- Source: [jyoung.dev](https://jyoung.dev/blog/claude-code-parallel-sessions-memory/)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T14:14:54Z

## Translation

Title: Claude Code Memory Usage: How Many Sessions Fit on One Machine
Article title: Claude Code Memory Usage: How Many Sessions Fit on One Machine — John Young
Description: Claude Code memory usage, measured: three sessions ran 32 processes and MCP servers outweighed Claude in two. A RAM budget for running sessions in parallel.

Article text:
Skip to content John Young essays
Claude Code Memory Usage: How Many Sessions Fit on One Machine
October 19, 2026 ·
32 min read ·
28 sources
How many Claude Code sessions fit on one machine depends less on the claude binary than on everything each session drags along: on my workstation, three sessions ran 32 processes, only three were Claude, and MCP servers outweighed Claude in two of them.
Measure each session as its whole process tree in PSS (or the cgroup’s memory.peak ), not summed RSS or VSZ, then cut the multiplier before buying RAM by trimming servers, dropping npx , and sharing stateless servers over HTTP.
My trees stayed flat across a 20-minute log, which says nothing about hour four, so fence each session in its own cgroup with MemoryHigh and MemoryMax and size with (ceiling minus headroom) divided by measured peak tree PSS, where the 25% headroom is my own judgment.
§ 1 Count Processes, Not Sessions
§ 2 Measure the Session Tree in PSS
§ 3 Cut the Multiplier Before You Buy RAM
§ 4 Budget for the Session’s Peak, Not Its Start
§ 5 Fence Every Session in Its Own Cgroup
§ 6 Size CI Runners with the Same Math
§ 7 The Max-Sessions Sizing Table
This is about what happens to a computer’s working memory when you run several AI coding assistants side by side, and how to plan for it.
When you open an AI coding assistant, you think you are starting one program. You are actually starting a small crowd. It is like hiring one contractor who shows up with a van full of specialists, each with their own toolbox. Count only the contractor and you will badly underestimate how much room you need on the job site.
The author counted the programs running behind three assistant windows on his own workstation. There were 32 of them, and only three were the assistant itself. In two of the three windows, the add-on helpers used more memory than the assistant did. So how many windows one computer can hold depends less on the assistant and more on everything it brings along. To plan well, you need to measure the whole crowd fairly, check it at its worst moment, and put a fence around each window so one runaway can’t take down the rest.
Note that “memory” here means the computer’s working memory (RAM), the space that fills up and makes programs get shut down. It does not mean the assistant’s feature for remembering notes between chats.
Count the whole crowd, not just the windows
Each assistant window starts its own copy of every add-on helper you have set up. Three windows with four helpers each means twelve helpers, not four. The helpers are not shared by default, even when they do the same job. The author’s helpers for browsing and for deployment tools each showed up as several separate copies.
Two things make the crowd bigger. A teammate assistant that one window spawns counts as a full extra window with its own helpers. And some helpers keep running after you close the window that started them, so count again after closing things, not just while they run.
Measure fairly: don’t count shared stuff twice
Programs often share the same underlying pieces. The usual way of adding up memory charges each program for the whole shared piece, like splitting a restaurant bill by giving every diner the full price of the shared appetizer. On the author’s machine, that made the total look about 62% bigger than it really was. A fairer method splits each shared piece among the programs that use it.
There is an even worse number. One widely repeated horror story said an assistant “consumed 129GB.” That figure was the space a program had reserved on paper, not memory it was actually using. The same report showed it was really using about 415 MB. Reserving a seat is not the same as sitting in it.
Trim the extras before buying more memory
Every helper you remove from your setup is removed from every window, so cutting helpers is the cheapest fix. Some ways to do it:
Keep only the helpers a given project needs.
Skip a launcher that quietly starts a second program just to start the first. On the author’s machine, each helper started that way cost two programs of about 100 MB each, and roughly half of that was the launcher doing no useful work. That half is the author’s measurement; the fix of installing the helper directly is his inference, so he checks the program list afterward.
Share helpers that don’t keep per-window state, by running one copy that all windows talk to. One person cut from 172 programs to about 33 this way. The catch: if any one shared helper fails to start, the whole shared setup can go down with it. The author shares only helpers that hold no private state, and keeps the web browser helper separate per window, because each window drives its own browser.
One trap worth knowing: a feature that makes helpers cheaper for the assistant’s thinking space does not make them cheaper for the computer’s memory. The helpers still run. It saves one kind of room, not the other.
Plan for the worst moment, not the start
A window’s size at startup is a floor, not a budget. The author sampled his three windows every five minutes, but the run was cut short by a restart, so he only has about 20 minutes of data. In that time the windows stayed roughly flat; the biggest one rose by 63 MB and then gave most of it back. That is all 20 minutes can say. It says nothing about what happens at hour four.
For longer stretches, the evidence comes from other people’s bug reports, and it is worse: sessions that grew from under a gigabyte to many gigabytes and got killed. The maker has been fixing these one release at a time, so which version you run matters. The maker also shows a warning when a window gets too big. The advice is to restart the window and pick the conversation back up when you see it, and to restart between big tasks anyway.
Put a fence around each window
Without a fence, all windows draw from one shared pool, and one runaway window can freeze the whole machine. On Linux, you can give each window its own memory limit with two settings: a soft one that slows it down when it gets greedy, and a hard one that is the last resort. Then a leak costs you one window, and that window can be resumed.
Two honest limits. First, a fence protects the computer, not the window; the greedy window still gets shut down. Second, Macs don’t have this fencing tool, so on a Mac the options are fewer helpers and fewer windows.
The same math works for automated runs on a build server. Run them in a stripped-down mode so they only start the helpers you listed, put a hard memory cap on each, and record how much each run actually used. The author’s claim that the stripped-down mode saves memory is an inference from his measurements, not something the maker states.
Maximum windows is roughly: (total memory, minus a cushion) divided by the worst size one window reaches.
The cushion, about a quarter of the memory, is the author’s own rule of thumb, not a sourced figure. On his machine, the math gives about 119 windows if every window stays as small as it did in his 20-minute sample, but only about 12 if every window hit its hard limit at once. The gap between those numbers is the whole point. The true answer sits in between, and the fences are what make running well past 12 a reasonable bet. For most teams, though, memory is not the real limit. The number of changes a human can review limits useful parallel work long before the computer does.
Before you plan, list every program each window starts, not just the window.
Measure with the fair method, and look at the worst moment, not the average.
Trim helpers first. It is cheaper than more memory.
Fence each window if your system allows it, and restart long-running windows rather than hoping.
Re-measure after you change your helpers or update the assistant, because both change the answer.
Remember that this author’s own direct measurements cover about 20 minutes. Treat his numbers as a method to copy, not figures to rely on.
The technical terms, in plain words
Claude Code = an AI assistant for writing software that runs on your own computer
session = one open working window with the assistant
process = one running program; a single session is really many of them
MCP server = an add-on helper that gives the assistant extra abilities, such as controlling a web browser
stdio server = a helper that each window starts as its own private copy on your machine
RAM = the computer’s working memory, which fills up and causes shutdowns
RSS = the usual way of counting a program’s memory, which counts shared pieces in full for every program using them
PSS = the fairer count that splits each shared piece among the programs sharing it
VSZ = space a program has reserved on paper, not memory it is actually using
cgroup = a Linux fence that limits how much memory a group of programs can use
OOM killer = the system step that shuts a program down when memory runs out
npx = a launcher that starts a helper by first starting a second program
proxy = one shared middleman that many windows connect to, so helpers run once
heap = the part of a program’s memory it uses while working
teammate = an extra assistant a window can spawn, which counts as a full extra session
CI = an automated build server that runs checks without a person watching
WSL2 = a way of running Linux inside Windows, in its own small virtual computer
Keep reading: the full version, with the research and sources ↓
I counted the processes behind three Claude Code sessions on my workstation: 32 of them, and only three were Claude. In two of the three sessions, the MCP servers I had configured used more memory than Claude did. Memory here means RAM, the resident kind that fills a host and gets sessions OOM-killed, not the CLAUDE.md feature that shares the name. How many sessions one machine can hold depends less on the claude binary than on everything each session drags along with it. That means measuring the whole tree correctly, at its worst moment, inside a fence that keeps one bad session from taking down the rest.
A Claude Code session is not one process. It is a tree: the claude binary, every stdio MCP server it launched, and anything those servers launched in turn. With stdio servers, N sessions times M servers is the number that fills RAM. List the tree for each session before you budget a single megabyte.
Here is that list at 10:42 on October 19, 2026. The host is a WSL2 VM on a 128 GB Windows box running Claude Code 2.1.303, with three interactive sessions open side by side.
Session A runs two MCP servers: proto-mcp, a Go binary at about 75 MB, and playwright-mcp. Sessions B and C add cloud-run-mcp and gitbook-mcp, which pushes their servers past the agent by roughly two to one. Across the host that is three copies of playwright-mcp, three of proto-mcp, two of cloud-run-mcp, and two of gitbook-mcp. Each is a separate process with its own heap, doing the same job as its twins. ( ps --ppid <claude-pid> -o pid,rss,args gives you the same view of a session’s direct children on your own machine.)
This is the documented design, not a quirk of my config. Claude Code’s MCP docs state that “Stdio servers run as local processes on your machine” ( Claude Code Docs: MCP 1 ). The protocol defines stdio as the transport where “The client launches the MCP server as a subprocess” ( MCP Spec: Transports 2 ). Every session is its own client, so every session launches its own copy. Claude Code does deduplicate, but only inside one session. When a server is defined in more than one place, it “connects to it once, using the definition from the highest-precedence source” ( Claude Code Docs: MCP 1 ).
Other people’s hosts show the same shape at larger N. One user filed the expected count as “4 sessions with 6 MCP servers should spawn 6 MCP processes total, not 42” ( JotaNikkan-pro: #28860 3 ). A report from October puts the desktop app at roughly 10 GB of duplicate server processes across ten sessions ( lrgibs01: #99831 4 ). At the far end:
With N sessions and M co

[truncated]

## Original Extract

Claude Code memory usage, measured: three sessions ran 32 processes and MCP servers outweighed Claude in two. A RAM budget for running sessions in parallel.

Skip to content John Young essays
Claude Code Memory Usage: How Many Sessions Fit on One Machine
October 19, 2026 ·
32 min read ·
28 sources
How many Claude Code sessions fit on one machine depends less on the claude binary than on everything each session drags along: on my workstation, three sessions ran 32 processes, only three were Claude, and MCP servers outweighed Claude in two of them.
Measure each session as its whole process tree in PSS (or the cgroup’s memory.peak ), not summed RSS or VSZ, then cut the multiplier before buying RAM by trimming servers, dropping npx , and sharing stateless servers over HTTP.
My trees stayed flat across a 20-minute log, which says nothing about hour four, so fence each session in its own cgroup with MemoryHigh and MemoryMax and size with (ceiling minus headroom) divided by measured peak tree PSS, where the 25% headroom is my own judgment.
§ 1 Count Processes, Not Sessions
§ 2 Measure the Session Tree in PSS
§ 3 Cut the Multiplier Before You Buy RAM
§ 4 Budget for the Session’s Peak, Not Its Start
§ 5 Fence Every Session in Its Own Cgroup
§ 6 Size CI Runners with the Same Math
§ 7 The Max-Sessions Sizing Table
This is about what happens to a computer’s working memory when you run several AI coding assistants side by side, and how to plan for it.
When you open an AI coding assistant, you think you are starting one program. You are actually starting a small crowd. It is like hiring one contractor who shows up with a van full of specialists, each with their own toolbox. Count only the contractor and you will badly underestimate how much room you need on the job site.
The author counted the programs running behind three assistant windows on his own workstation. There were 32 of them, and only three were the assistant itself. In two of the three windows, the add-on helpers used more memory than the assistant did. So how many windows one computer can hold depends less on the assistant and more on everything it brings along. To plan well, you need to measure the whole crowd fairly, check it at its worst moment, and put a fence around each window so one runaway can’t take down the rest.
Note that “memory” here means the computer’s working memory (RAM), the space that fills up and makes programs get shut down. It does not mean the assistant’s feature for remembering notes between chats.
Count the whole crowd, not just the windows
Each assistant window starts its own copy of every add-on helper you have set up. Three windows with four helpers each means twelve helpers, not four. The helpers are not shared by default, even when they do the same job. The author’s helpers for browsing and for deployment tools each showed up as several separate copies.
Two things make the crowd bigger. A teammate assistant that one window spawns counts as a full extra window with its own helpers. And some helpers keep running after you close the window that started them, so count again after closing things, not just while they run.
Measure fairly: don’t count shared stuff twice
Programs often share the same underlying pieces. The usual way of adding up memory charges each program for the whole shared piece, like splitting a restaurant bill by giving every diner the full price of the shared appetizer. On the author’s machine, that made the total look about 62% bigger than it really was. A fairer method splits each shared piece among the programs that use it.
There is an even worse number. One widely repeated horror story said an assistant “consumed 129GB.” That figure was the space a program had reserved on paper, not memory it was actually using. The same report showed it was really using about 415 MB. Reserving a seat is not the same as sitting in it.
Trim the extras before buying more memory
Every helper you remove from your setup is removed from every window, so cutting helpers is the cheapest fix. Some ways to do it:
Keep only the helpers a given project needs.
Skip a launcher that quietly starts a second program just to start the first. On the author’s machine, each helper started that way cost two programs of about 100 MB each, and roughly half of that was the launcher doing no useful work. That half is the author’s measurement; the fix of installing the helper directly is his inference, so he checks the program list afterward.
Share helpers that don’t keep per-window state, by running one copy that all windows talk to. One person cut from 172 programs to about 33 this way. The catch: if any one shared helper fails to start, the whole shared setup can go down with it. The author shares only helpers that hold no private state, and keeps the web browser helper separate per window, because each window drives its own browser.
One trap worth knowing: a feature that makes helpers cheaper for the assistant’s thinking space does not make them cheaper for the computer’s memory. The helpers still run. It saves one kind of room, not the other.
Plan for the worst moment, not the start
A window’s size at startup is a floor, not a budget. The author sampled his three windows every five minutes, but the run was cut short by a restart, so he only has about 20 minutes of data. In that time the windows stayed roughly flat; the biggest one rose by 63 MB and then gave most of it back. That is all 20 minutes can say. It says nothing about what happens at hour four.
For longer stretches, the evidence comes from other people’s bug reports, and it is worse: sessions that grew from under a gigabyte to many gigabytes and got killed. The maker has been fixing these one release at a time, so which version you run matters. The maker also shows a warning when a window gets too big. The advice is to restart the window and pick the conversation back up when you see it, and to restart between big tasks anyway.
Put a fence around each window
Without a fence, all windows draw from one shared pool, and one runaway window can freeze the whole machine. On Linux, you can give each window its own memory limit with two settings: a soft one that slows it down when it gets greedy, and a hard one that is the last resort. Then a leak costs you one window, and that window can be resumed.
Two honest limits. First, a fence protects the computer, not the window; the greedy window still gets shut down. Second, Macs don’t have this fencing tool, so on a Mac the options are fewer helpers and fewer windows.
The same math works for automated runs on a build server. Run them in a stripped-down mode so they only start the helpers you listed, put a hard memory cap on each, and record how much each run actually used. The author’s claim that the stripped-down mode saves memory is an inference from his measurements, not something the maker states.
Maximum windows is roughly: (total memory, minus a cushion) divided by the worst size one window reaches.
The cushion, about a quarter of the memory, is the author’s own rule of thumb, not a sourced figure. On his machine, the math gives about 119 windows if every window stays as small as it did in his 20-minute sample, but only about 12 if every window hit its hard limit at once. The gap between those numbers is the whole point. The true answer sits in between, and the fences are what make running well past 12 a reasonable bet. For most teams, though, memory is not the real limit. The number of changes a human can review limits useful parallel work long before the computer does.
Before you plan, list every program each window starts, not just the window.
Measure with the fair method, and look at the worst moment, not the average.
Trim helpers first. It is cheaper than more memory.
Fence each window if your system allows it, and restart long-running windows rather than hoping.
Re-measure after you change your helpers or update the assistant, because both change the answer.
Remember that this author’s own direct measurements cover about 20 minutes. Treat his numbers as a method to copy, not figures to rely on.
The technical terms, in plain words
Claude Code = an AI assistant for writing software that runs on your own computer
session = one open working window with the assistant
process = one running program; a single session is really many of them
MCP server = an add-on helper that gives the assistant extra abilities, such as controlling a web browser
stdio server = a helper that each window starts as its own private copy on your machine
RAM = the computer’s working memory, which fills up and causes shutdowns
RSS = the usual way of counting a program’s memory, which counts shared pieces in full for every program using them
PSS = the fairer count that splits each shared piece among the programs sharing it
VSZ = space a program has reserved on paper, not memory it is actually using
cgroup = a Linux fence that limits how much memory a group of programs can use
OOM killer = the system step that shuts a program down when memory runs out
npx = a launcher that starts a helper by first starting a second program
proxy = one shared middleman that many windows connect to, so helpers run once
heap = the part of a program’s memory it uses while working
teammate = an extra assistant a window can spawn, which counts as a full extra session
CI = an automated build server that runs checks without a person watching
WSL2 = a way of running Linux inside Windows, in its own small virtual computer
Keep reading: the full version, with the research and sources ↓
I counted the processes behind three Claude Code sessions on my workstation: 32 of them, and only three were Claude. In two of the three sessions, the MCP servers I had configured used more memory than Claude did. Memory here means RAM, the resident kind that fills a host and gets sessions OOM-killed, not the CLAUDE.md feature that shares the name. How many sessions one machine can hold depends less on the claude binary than on everything each session drags along with it. That means measuring the whole tree correctly, at its worst moment, inside a fence that keeps one bad session from taking down the rest.
A Claude Code session is not one process. It is a tree: the claude binary, every stdio MCP server it launched, and anything those servers launched in turn. With stdio servers, N sessions times M servers is the number that fills RAM. List the tree for each session before you budget a single megabyte.
Here is that list at 10:42 on October 19, 2026. The host is a WSL2 VM on a 128 GB Windows box running Claude Code 2.1.303, with three interactive sessions open side by side.
Session A runs two MCP servers: proto-mcp, a Go binary at about 75 MB, and playwright-mcp. Sessions B and C add cloud-run-mcp and gitbook-mcp, which pushes their servers past the agent by roughly two to one. Across the host that is three copies of playwright-mcp, three of proto-mcp, two of cloud-run-mcp, and two of gitbook-mcp. Each is a separate process with its own heap, doing the same job as its twins. ( ps --ppid <claude-pid> -o pid,rss,args gives you the same view of a session’s direct children on your own machine.)
This is the documented design, not a quirk of my config. Claude Code’s MCP docs state that “Stdio servers run as local processes on your machine” ( Claude Code Docs: MCP 1 ). The protocol defines stdio as the transport where “The client launches the MCP server as a subprocess” ( MCP Spec: Transports 2 ). Every session is its own client, so every session launches its own copy. Claude Code does deduplicate, but only inside one session. When a server is defined in more than one place, it “connects to it once, using the definition from the highest-precedence source” ( Claude Code Docs: MCP 1 ).
Other people’s hosts show the same shape at larger N. One user filed the expected count as “4 sessions with 6 MCP servers should spawn 6 MCP processes total, not 42” ( JotaNikkan-pro: #28860 3 ). A report from October puts the desktop app at roughly 10 GB of duplicate server processes across ten sessions ( lrgibs01: #99831 4 ). At the far end:
With N sessions and M co

[truncated]
