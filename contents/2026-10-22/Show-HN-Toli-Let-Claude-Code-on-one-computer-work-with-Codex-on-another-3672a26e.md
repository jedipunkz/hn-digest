---
source: "https://github.com/arjunvlama/toli"
hn_url: "https://news.ycombinator.com/item?id=50194681"
title: "Show HN: Toli – Let Claude Code on one computer work with Codex on another"
article_title: "GitHub - arjunvlama/toli: Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many. · GitHub"
image: "https://opengraph.githubassets.com/e1f79e0d1ad711216c74fce6aceecd92e0f48f9a83392313abedd65673f91eaa/arjunvlama/toli"
author: "alama24"
captured_at: "2026-10-22T13:31:37Z"
capture_tool: "hn-digest"
hn_id: 50194681
score: 1
comments: 0
posted_at: "2026-10-22T12:54:31Z"
tags:
  - hacker-news
---

# Show HN: Toli – Let Claude Code on one computer work with Codex on another

- HN: [50194681](https://news.ycombinator.com/item?id=50194681)
- Source: [github.com](https://github.com/arjunvlama/toli)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T12:54:31Z

## Translation

Title: Show HN: Toli – Let Claude Code on one computer work with Codex on another
Article title: GitHub - arjunvlama/toli: Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many. · GitHub
Description: Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many. - arjunvlama/toli
HN text: A month or so ago, a colleague was running some code my team owned on his machine and hit a niche Windows kernel-related Python issue. We were debugging it together, but kept having to relay information between our coding agents, mine understood the codebase, while his had the context of the environment where it was actually failing. At some point I realized we were basically acting as the message bus between them, copying findings and questions back and forth. That made me wonder, why not just let the agents work together directly? I found A2A and built toli, a local daemon that an existing agent can connect to over MCP, giving it the plumbing to discover and work with A2A peers, including another person’s toli-connected agent on another machine. It also works locally. I use it to let two of my agents working on sister codebases collaborate on bugs and design without having to give either agent both repos. Would love feedback: https://github.com/arjunvlama/toli

Article text:
GitHub - arjunvlama/toli: Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
arjunvlama
toli Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
.github .github cmd/ toli cmd/ toli docs docs e2e e2e examples/ custom-agent examples/ custom-agent internal internal .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml ARCHITECTURE.md ARCHITECTURE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
toli — the telephone for AI agents ☎️
toli (टोली, "toh-lee") is Hindi for a band of people working together.
With toli, your AI agents can work together no matter which harness they run in (Claude Code, Codex, OpenCode, OpenClaw, … or one you built yourself) or where they run, and with any agent that already speaks A2A. They can collaborate on one machine or across the internet with a teammate's agents, handing off tasks, asking each other questions and sending back results over the open A2A protocol.
It's not just for coding agents. Any agent with an MCP client can join, so a personal assistant can check with your coding agent before it answers you:
You run one toli daemon. Each of your agents gets a name and a token there, which it uses to reach toli over MCP, on localhost.
Each agent has peers : the other agents it's allowed to work with, on your daemon or someone else's. Only you can add peers, never the agent. Pairing exchanges a token in each direction.
Agents and their peers send each other tasks over A2A. They can delegate work, ask questions and send back results. When a task arrives, toli wakes your agent and notifies it.
You can watch and manage all their tasks in toli (a TUI), or with the CLI.
curl -fsSL https://raw.githubusercontent.com/arjunvlama/toli/main/install.sh | sh
Homebrew:
brew install arjunvlama/tap/toli
Windows (PowerShell):
irm https: // raw.githubusercontent.com / arjunvlama / toli / main / install.ps1 | iex
With Go: go install github.com/arjunvlama/toli/cmd/toli@latest
Then start the daemon in the background:
toli daemon install
toli daemon status
toli daemon stop , start and uninstall manage it.
1. Add the agents, and connect them to toli. Run this in the folder where you start each agent: where you type claude or codex , or the folder you opened in your editor (its built-in terminal starts there). --connect names the agent's harness ( all harnesses ), and peers see the description when deciding what to send it.
cd ~ /src/web-ui
toli agent add alice --description " Owns the web UI " --connect claude-code
cd ~ /src/api
toli agent add bob --description " Owns the API " --connect codex
Each prints what it did:
✓ Added agent 'alice'.
card_url http://127.0.0.1:7422/a2a/alice/.well-known/agent-card.json
✓ Connected Claude Code to agent 'alice' in ~/src/web-ui (~/.claude.json, .claude/settings.local.json).
Next: start a Claude Code session, or restart a running one (`claude --continue`).
Sessions only pick toli up when they start, so start one in each folder, or restart any that are already running. Then toli agent list shows each agent's harness, which means it's connected. (Some harnesses, like Codex, only show up once you've sent them a message after starting or restarting.)
2. Pair them. alice invites bob:
toli peer invite alice bob
✓ Created an invite for peer 'bob'. It works once, within 24 hours.
To accept it, run this where peer 'bob' runs:
toli peer accept <agent> toli-invite:eyJ2IjoxLCJhZ2Vu…
Send it privately: whoever has it can pair with agent 'alice'.
bob accepts:
toli peer accept bob < invite >
✓ Paired agent 'bob' with peer 'alice'.
Check they're connected:
toli peer list alice
NAME CONNECTED HAS_TOKEN URL LAST_ERROR
bob true true http://127.0.0.1:7422/a2a/bob/.well-known/agent-card.json
3. Ask. In alice's session: "Ask bob what it's working on." bob is woken with the task. To watch them work together, run toli (the TUI), or toli task list alice .
The steps are the same as on one machine. The daemons just need to be able to reach each other over the network.
Making them reachable. toli doesn't handle networking itself. Your daemon listens for its peers on its A2A address ( --a2a-listen , localhost by default), and tells them to reach it at its public URL ( --public-url , the A2A address by default). Traffic between daemons is plain HTTP and carries tokens and task contents, so choose how they connect:
Tailscale (recommended). Traffic is encrypted, and there are no ports to open. Listen on the machine's Tailscale IP ( tailscale ip -4 prints it) and give peers that same address:
toli daemon install --a2a-listen < tailscale IP > : < port > --public-url http:// < tailscale IP > : < port >
Running toli daemon install again replaces the earlier setup.
A network you trust (home, office): do the same with the machine's local IP.
The public internet : only behind HTTPS. Keep the daemon listening locally, and put an HTTPS tunnel (Cloudflare Tunnel, ngrok) or reverse proxy in front of its A2A address. Use the tunnel's https:// address as your daemon's public URL. Never make the A2A address itself reachable from the internet.
Then pair them as in step 2. One of you creates the invite and sends it to the other privately, like a password; the other accepts it on their own daemon. To pair without an invite, use toli peer add and toli peer token ( toli peer --help ).
toli connect < harness > < agent > [--scope project | user]
This gives your agent toli's tools (MCP) and notifications, so it's woken when work arrives. Harnesses: claude-code , codex , opencode , pi and openclaw . toli agent add --connect does the same when adding an agent.
Run it in the folder where you start the agent. It connects sessions started there; use --scope user to connect all your sessions.
It makes a new MCP URL for the agent, so a session still using the old one is disconnected. To use a URL you already have, pass --mcp-url <url> instead of the agent.
Then start a new session, or restart the ones already running. Sessions only pick toli up when they start. Once one has started, toli agent list shows its harness, which means it's connected:
NAME HARNESS DISABLED NOTIFY CARD_URL
alice claude-code false http://127.0.0.1:7422/a2a/alice/.well-known/agent-card.json
Some harnesses only register a session once you've sent it a first message, so if yours doesn't show up yet, say something to it.
toli disconnect <harness> undoes it.
Any agent with an MCP client can use toli. Point its MCP client at the agent's MCP URL, which toli agent new-token <agent> prints, to give it toli's tools.
To have it woken when work arrives, use one of these:
Pull if your agent is a program: it listens for notifications in the background ( pull.py ).
Push if your agent is a service with an API: toli calls a webhook or runs a program with each notification. Set it up with toli agent notify <agent> --webhook <url> ( webhook.py ).
See docs/custom-agents.md for how it works.
You decide who your agents talk to. Only you can add peers, with the CLI or TUI; an agent can't. Peers you add are trusted. If your agent's harness can be woken, a peer's task starts it working.
Tokens. Each agent has a token, in its MCP URL, for talking to the daemon. Each pairing has a token in each direction. toli agent new-token and toli peer new-token replace them.
The daemon's MCP server and admin API listen on localhost only. The admin API needs the admin token in ~/.toli .
Between machines, use Tailscale or HTTPS (see above). toli's A2A traffic is plain HTTP on its own.
Sandbox your agents. toli keeps out peers, browsers and other users, but not an agent running a shell as you, which could read ~/.toli or call the admin API. Run agents in a sandbox that can reach toli's MCP port only: deny ~/.toli/** and the admin API's port (7421 by default).
For contributors, please see CONTRIBUTING.md , and for more context on the architecture, see ARCHITECTURE.md . We have a lot of ideas for expanding toli, so come chat with us on Discord , or start a thread in Discussions for any idea you have.
If you've found a security vulnerability, please don't open an issue. Report it privately instead, as described in SECURITY.md .
toli is early, but it works. It's tested with Claude Code, Codex, OpenCode, OpenClaw and Pi, and the daemon runs on macOS, Linux and Windows.
toli is licensed under Apache-2.0 .
Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many. - arjunvlama/toli

A month or so ago, a colleague was running some code my team owned on his machine and hit a niche Windows kernel-related Python issue. We were debugging it together, but kept having to relay information between our coding agents, mine understood the codebase, while his had the context of the environment where it was actually failing. At some point I realized we were basically acting as the message bus between them, copying findings and questions back and forth. That made me wonder, why not just let the agents work together directly? I found A2A and built toli, a local daemon that an existing agent can connect to over MCP, giving it the plumbing to discover and work with A2A peers, including another person’s toli-connected agent on another machine. It also works locally. I use it to let two of my agents working on sister codebases collaborate on bugs and design without having to give either agent both repos. Would love feedback: https://github.com/arjunvlama/toli

GitHub - arjunvlama/toli: Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
arjunvlama
toli Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
.github .github cmd/ toli cmd/ toli docs docs e2e e2e examples/ custom-agent examples/ custom-agent internal internal .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml ARCHITECTURE.md ARCHITECTURE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
toli — the telephone for AI agents ☎️
toli (टोली, "toh-lee") is Hindi for a band of people working together.
With toli, your AI agents can work together no matter which harness they run in (Claude Code, Codex, OpenCode, OpenClaw, … or one you built yourself) or where they run, and with any agent that already speaks A2A. They can collaborate on one machine or across the internet with a teammate's agents, handing off tasks, asking each other questions and sending back results over the open A2A protocol.
It's not just for coding agents. Any agent with an MCP client can join, so a personal assistant can check with your coding agent before it answers you:
You run one toli daemon. Each of your agents gets a name and a token there, which it uses to reach toli over MCP, on localhost.
Each agent has peers : the other agents it's allowed to work with, on your daemon or someone else's. Only you can add peers, never the agent. Pairing exchanges a token in each direction.
Agents and their peers send each other tasks over A2A. They can delegate work, ask questions and send back results. When a task arrives, toli wakes your agent and notifies it.
You can watch and manage all their tasks in toli (a TUI), or with the CLI.
curl -fsSL https://raw.githubusercontent.com/arjunvlama/toli/main/install.sh | sh
Homebrew:
brew install arjunvlama/tap/toli
Windows (PowerShell):
irm https: // raw.githubusercontent.com / arjunvlama / toli / main / install.ps1 | iex
With Go: go install github.com/arjunvlama/toli/cmd/toli@latest
Then start the daemon in the background:
toli daemon install
toli daemon status
toli daemon stop , start and uninstall manage it.
1. Add the agents, and connect them to toli. Run this in the folder where you start each agent: where you type claude or codex , or the folder you opened in your editor (its built-in terminal starts there). --connect names the agent's harness ( all harnesses ), and peers see the description when deciding what to send it.
cd ~ /src/web-ui
toli agent add alice --description " Owns the web UI " --connect claude-code
cd ~ /src/api
toli agent add bob --description " Owns the API " --connect codex
Each prints what it did:
✓ Added agent 'alice'.
card_url http://127.0.0.1:7422/a2a/alice/.well-known/agent-card.json
✓ Connected Claude Code to agent 'alice' in ~/src/web-ui (~/.claude.json, .claude/settings.local.json).
Next: start a Claude Code session, or restart a running one (`claude --continue`).
Sessions only pick toli up when they start, so start one in each folder, or restart any that are already running. Then toli agent list shows each agent's harness, which means it's connected. (Some harnesses, like Codex, only show up once you've sent them a message after starting or restarting.)
2. Pair them. alice invites bob:
toli peer invite alice bob
✓ Created an invite for peer 'bob'. It works once, within 24 hours.
To accept it, run this where peer 'bob' runs:
toli peer accept <agent> toli-invite:eyJ2IjoxLCJhZ2Vu…
Send it privately: whoever has it can pair with agent 'alice'.
bob accepts:
toli peer accept bob < invite >
✓ Paired agent 'bob' with peer 'alice'.
Check they're connected:
toli peer list alice
NAME CONNECTED HAS_TOKEN URL LAST_ERROR
bob true true http://127.0.0.1:7422/a2a/bob/.well-known/agent-card.json
3. Ask. In alice's session: "Ask bob what it's working on." bob is woken with the task. To watch them work together, run toli (the TUI), or toli task list alice .
The steps are the same as on one machine. The daemons just need to be able to reach each other over the network.
Making them reachable. toli doesn't handle networking itself. Your daemon listens for its peers on its A2A address ( --a2a-listen , localhost by default), and tells them to reach it at its public URL ( --public-url , the A2A address by default). Traffic between daemons is plain HTTP and carries tokens and task contents, so choose how they connect:
Tailscale (recommended). Traffic is encrypted, and there are no ports to open. Listen on the machine's Tailscale IP ( tailscale ip -4 prints it) and give peers that same address:
toli daemon install --a2a-listen < tailscale IP > : < port > --public-url http:// < tailscale IP > : < port >
Running toli daemon install again replaces the earlier setup.
A network you trust (home, office): do the same with the machine's local IP.
The public internet : only behind HTTPS. Keep the daemon listening locally, and put an HTTPS tunnel (Cloudflare Tunnel, ngrok) or reverse proxy in front of its A2A address. Use the tunnel's https:// address as your daemon's public URL. Never make the A2A address itself reachable from the internet.
Then pair them as in step 2. One of you creates the invite and sends it to the other privately, like a password; the other accepts it on their own daemon. To pair without an invite, use toli peer add and toli peer token ( toli peer --help ).
toli connect < harness > < agent > [--scope project | user]
This gives your agent toli's tools (MCP) and notifications, so it's woken when work arrives. Harnesses: claude-code , codex , opencode , pi and openclaw . toli agent add --connect does the same when adding an agent.
Run it in the folder where you start the agent. It connects sessions started there; use --scope user to connect all your sessions.
It makes a new MCP URL for the agent, so a session still using the old one is disconnected. To use a URL you already have, pass --mcp-url <url> instead of the agent.
Then start a new session, or restart the ones already running. Sessions only pick toli up when they start. Once one has started, toli agent list shows its harness, which means it's connected:
NAME HARNESS DISABLED NOTIFY CARD_URL
alice claude-code false http://127.0.0.1:7422/a2a/alice/.well-known/agent-card.json
Some harnesses only register a session once you've sent it a first message, so if yours doesn't show up yet, say something to it.
toli disconnect <harness> undoes it.
Any agent with an MCP client can use toli. Point its MCP client at the agent's MCP URL, which toli agent new-token <agent> prints, to give it toli's tools.
To have it woken when work arrives, use one of these:
Pull if your agent is a program: it listens for notifications in the background ( pull.py ).
Push if your agent is a service with an API: toli calls a webhook or runs a program with each notification. Set it up with toli agent notify <agent> --webhook <url> ( webhook.py ).
See docs/custom-agents.md for how it works.
You decide who your agents talk to. Only you can add peers, with the CLI or TUI; an agent can't. Peers you add are trusted. If your agent's harness can be woken, a peer's task starts it working.
Tokens. Each agent has a token, in its MCP URL, for talking to the daemon. Each pairing has a token in each direction. toli agent new-token and toli peer new-token replace them.
The daemon's MCP server and admin API listen on localhost only. The admin API needs the admin token in ~/.toli .
Between machines, use Tailscale or HTTPS (see above). toli's A2A traffic is plain HTTP on its own.
Sandbox your agents. toli keeps out peers, browsers and other users, but not an agent running a shell as you, which could read ~/.toli or call the admin API. Run agents in a sandbox that can reach toli's MCP port only: deny ~/.toli/** and the admin API's port (7421 by default).
For contributors, please see CONTRIBUTING.md , and for more context on the architecture, see ARCHITECTURE.md . We have a lot of ideas for expanding toli, so come chat with us on Discord , or start a thread in Discussions for any idea you have.
If you've found a security vulnerability, please don't open an issue. Report it privately instead, as described in SECURITY.md .
toli is early, but it works. It's tested with Claude Code, Codex, OpenCode, OpenClaw and Pi, and the daemon runs on macOS, Linux and Windows.
toli is licensed under Apache-2.0 .
Let Claude Code, Codex, OpenClaw or any other agent work together, on one machine or across many.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
