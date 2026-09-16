---
source: "https://accomplish.ai/blog/beltdown-escaping-the-claude-code-sandbox/"
hn_url: "https://news.ycombinator.com/item?id=49727207"
title: "Beltdown: Escaping the Claude Code Sandbox"
article_title: "Beltdown: Escaping the Claude Code sandbox — Accomplish Blog"
image: "https://accomplish.ai/images/beltdown-claude-code-key.jpg"
author: "snikolaev"
captured_at: "2026-09-16T14:39:35Z"
capture_tool: "hn-digest"
hn_id: 49727207
score: 1
comments: 0
posted_at: "2026-09-16T14:02:47Z"
tags:
  - hacker-news
---

# Beltdown: Escaping the Claude Code Sandbox

- HN: [49727207](https://news.ycombinator.com/item?id=49727207)
- Source: [accomplish.ai](https://accomplish.ai/blog/beltdown-escaping-the-claude-code-sandbox/)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T14:02:47Z

## Translation

Title: Beltdown: Escaping the Claude Code Sandbox
Article title: Beltdown: Escaping the Claude Code sandbox — Accomplish Blog
Description: An untrusted repository opened in Claude Code can escape the macOS sandbox and run commands on your computer as your privileged user. You never get the permission prompt.

Article text:
Boundary-Bench Blog Boundary-Bench Blog X Linkedin Beltdown: Escaping the C l a u d e C o d e s a n d b o x
An untrusted repository opened in Claude Code can escape the macOS sandbox and run commands on your computer as your privileged user. You never get the permission prompt.
Intro How the escape works Disclosure timeline We turned the sandbox on and set permission mode to the strictest “don’t ask”. Then we opened a repository in Claude Code and sent one short message. A command from that repo ran on our Mac anyway, outside the sandbox, with no permission prompt. It can also work through an indirect prompt injection.
Counterintuitively, the sandbox is why there was no permission prompt. When you turn it on, Claude sandboxes the agent’s commands using Seatbelt, so it stops asking before it runs them.
Following our report, Anthropic fixed this in Claude Code 2.1.247.
The harness runs its own git commands outside the sandbox. It runs them in the background to index the repo.
Git has a config setting, core.fsmonitor , that it reads from .git/config and runs as a shell command whenever it looks at the working tree.
So a rogue or compromised agent that wants out has to set core.fsmonitor in a .git/config file, then get the harness to run a git command that reads it. That git command is already outside the sandbox.
But four defenses get in the way:
Chained together, the setup script plants the poisoned .git in an unprotected nested path, tells Claude to read one file in that folder, and the file index refresh runs the payload.
Anthropic triaged our report quickly, and now every git command the harness runs blanks core.fsmonitor , so a repo’s config can’t run anything.
We put the whole agent in a VM. Bash, git, every process it starts, all of it is in there. Real credentials never enter the guest, the agent only holds placeholders. Its network traffic goes out through a proxy on the host that the agent can’t reconfigure.
A poisoned core.fsmonitor still runs, it just runs inside the VM, not on your Mac.
Escaping the OpenAI Codex sandbox, twice
Beltdown2: Escaping the Cursor CLI sandbox
SharedRoot: Escaping the Claude Cowork sandbox
Copyright © 2026 Accomplish AI, Inc. All rights reserved.

## Original Extract

An untrusted repository opened in Claude Code can escape the macOS sandbox and run commands on your computer as your privileged user. You never get the permission prompt.

Boundary-Bench Blog Boundary-Bench Blog X Linkedin Beltdown: Escaping the C l a u d e C o d e s a n d b o x
An untrusted repository opened in Claude Code can escape the macOS sandbox and run commands on your computer as your privileged user. You never get the permission prompt.
Intro How the escape works Disclosure timeline We turned the sandbox on and set permission mode to the strictest “don’t ask”. Then we opened a repository in Claude Code and sent one short message. A command from that repo ran on our Mac anyway, outside the sandbox, with no permission prompt. It can also work through an indirect prompt injection.
Counterintuitively, the sandbox is why there was no permission prompt. When you turn it on, Claude sandboxes the agent’s commands using Seatbelt, so it stops asking before it runs them.
Following our report, Anthropic fixed this in Claude Code 2.1.247.
The harness runs its own git commands outside the sandbox. It runs them in the background to index the repo.
Git has a config setting, core.fsmonitor , that it reads from .git/config and runs as a shell command whenever it looks at the working tree.
So a rogue or compromised agent that wants out has to set core.fsmonitor in a .git/config file, then get the harness to run a git command that reads it. That git command is already outside the sandbox.
But four defenses get in the way:
Chained together, the setup script plants the poisoned .git in an unprotected nested path, tells Claude to read one file in that folder, and the file index refresh runs the payload.
Anthropic triaged our report quickly, and now every git command the harness runs blanks core.fsmonitor , so a repo’s config can’t run anything.
We put the whole agent in a VM. Bash, git, every process it starts, all of it is in there. Real credentials never enter the guest, the agent only holds placeholders. Its network traffic goes out through a proxy on the host that the agent can’t reconfigure.
A poisoned core.fsmonitor still runs, it just runs inside the VM, not on your Mac.
Escaping the OpenAI Codex sandbox, twice
Beltdown2: Escaping the Cursor CLI sandbox
SharedRoot: Escaping the Claude Cowork sandbox
Copyright © 2026 Accomplish AI, Inc. All rights reserved.
