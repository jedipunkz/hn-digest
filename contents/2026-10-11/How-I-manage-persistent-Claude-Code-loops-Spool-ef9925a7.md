---
source: "https://github.com/enes-alatas/spool"
hn_url: "https://news.ycombinator.com/item?id=50048193"
title: "How I manage persistent Claude Code loops: Spool"
article_title: "GitHub - enes-alatas/spool · GitHub"
image: "https://opengraph.githubassets.com/e4789ade179cdb1b7359e1f5be2e0ba045a66ae292ab65b4489d82e24032bfb6/enes-alatas/spool"
author: "enesalatas"
captured_at: "2026-10-11T22:41:36Z"
capture_tool: "hn-digest"
hn_id: 50048193
score: 2
comments: 1
posted_at: "2026-10-11T21:50:14Z"
tags:
  - hacker-news
---

# How I manage persistent Claude Code loops: Spool

- HN: [50048193](https://news.ycombinator.com/item?id=50048193)
- Source: [github.com](https://github.com/enes-alatas/spool)
- Score: 2
- Comments: 1
- Posted: 2026-10-11T21:50:14Z

## Translation

Title: How I manage persistent Claude Code loops: Spool
Article title: GitHub - enes-alatas/spool · GitHub
Description: Contribute to enes-alatas/spool development by creating an account on GitHub.

Article text:
GitHub - enes-alatas/spool · GitHub
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
enes-alatas
spool Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
668 Commits 668 Commits Folders and files
.claude/ skills .claude/ skills .github .github cmd cmd docker docker docs docs internal internal itest itest scripts scripts web web .editorconfig .editorconfig .gitignore .gitignore .golangci.yml .golangci.yml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Your Claude Code loops, working together and keeping you in the conversation.
Spool makes it easy to create, manage, and talk to persistent Claude Code loops. Give them work, let them coordinate, and stay involved through chat.
spool-demo.mp4
What does Spool make easier?
Keep work going. Create and manage persistent loops in one place, with scheduled wakes and resumable sessions.
Stay involved. Talk to your loops in the web control room, on Telegram or in Slack, without sitting in their terminal sessions.
Coordinate work. Let loops communicate directly in the fleet channel, so you don't have to relay every question, result, or handoff.
Less switching between sessions and carrying messages between loops. More time deciding what they should work on and reviewing what they produce.
From a clone to a loop that answers you on Telegram. Each step is one command or one action.
Have the prerequisites. Linux or macOS, Claude Code installed and logged in ( claude on PATH), Docker running, and Go 1.25+ with Node 20.19+ or 22.12+ for the build.
Build the binary: make build , which builds the web UI and the single bin/spool binary.
Build the images: make image , which builds the workstation a loop runs in and its egress proxy. Spool runs them from your local Docker; it never pulls them.
Start the hub: ./bin/spool . The control room is on 127.0.0.1:8080, and data goes in ~/.spool . The loop listener goes where your workstations can reach it, and the hub logs where that is; see What contains a loop .
Sign in to the control room: open http://127.0.0.1:8080 and sign in as admin with the one-time password the hub printed at its first start, then choose your own password. Missed it? ./bin/spool user reset admin prints a new one.
Give the workstations your Claude login: run claude setup-token , then paste the token under Settings → Claude token . A workstation cannot use your machine's ~/.claude , so this is how a loop runs on your plan.
Create a loop: New loop , then a name, a mission and a tick interval. Leave the workspace empty; it is for bare loops only.
Attach Telegram: in Telegram, send @BotFather /newbot . On the loop's page, under Surfaces , choose Attach Telegram and paste the token.
Allow yourself: DM the bot. It answers with a pairing code, and you appear as pending on the Access page. Type the code into your row and click Allow , then choose yourself as the owner on the loop's page.
Say hello: DM the bot again. Your message wakes the loop, and it answers in that DM.
No Docker? Start with ./bin/spool --runtime bare instead, and skip steps 3 and 6. The loop then runs as a host subprocess under your own account, uncontained . Read What contains a loop before you choose that.
To bring loops into a Telegram group, and for who can reach them there, see Telegram .
Plain claude CLI underneath. Loops are claude subprocesses speaking stream-json over stdin/stdout — your normal Claude Code login, plan, and session token limits. No API keys, no SDK.
Direct coordination. Loops send explicitly addressed messages to the fleet channel or a private conversation. @mention a loop from Telegram, the web, or another loop's message, and Spool delivers it. Final turn replies stay as status notes in the timeline. Loop-to-loop chains are storm-guarded.
Wake/sleep engine. Idle loops cost nothing: their process exits and resumes later via --resume with full context. Loops can self-pace with a [next-wake: 45m] trailer, clamped to bounds you set.
Contained by default. Each loop runs in its own container with an egress allowlist and no route to your machine but one hub port. Without Docker, Spool refuses to start rather than quietly running loops on your host — uncontained is something you ask for, and it says so when you do. See What contains a loop .
Worktree isolation. Point several loops at one repo and each gets its own git worktree on loop/<name> — they can't clobber each other.
Control room. Live conversation timelines (token streaming included), schedules, costs per turn/day, pause/wake/kill.
Plan guardrails. The Fleet page shows how much of your plan's 5-hour and 7-day windows is used, read from the loops' own streams, and each loop's share of the day's spend. Past a cap, 90% of either window by default, every loop sleeps until that window resets: messages wait, and nothing is lost. Resume now lifts it early, and Settings → Plan guardrails sets the thresholds or turns them off (ADR-0047).
Spool is pre-1.0 and developed in the open by a fleet of Claude Code loops and their operator — the four loops in this repo's issues and PRs are running on Spool, building Spool. Commits are co-authored by the model that wrote them; every change lands through a pull request that CI gates and a reviewer reads, and the operator is the one who merges it. Found a problem? CONTRIBUTING.md says how to report it and who answers.
Spool runs every loop with --permission-mode bypassPermissions : a loop is
never asked to confirm anything it does. What contains it is its workstation
— the runtime it runs in — and not the workspace directory you point it at. A
workspace is a working directory; it stops nothing.
With Docker reachable, a loop gets a container of its own, on an internal
network with no route off it except an allowlist of the hosts a loop needs
( spool-egress-proxy ), and the single hub port that serves the MCP endpoint.
You add hosts to that allowlist under Settings → Egress . A change reaches
every docker loop within seconds, with nothing restarted.
An http MCP server you attach to a loop is reached through the hub, which adds
its credential, so the loop calls the server without ever holding its secret.
It cannot reach your control room API, your files, or the rest of your network.
Without a reachable Docker daemon, --runtime auto refuses to start and
tells you the two ways on: install or start Docker, or ask for host
subprocesses deliberately with --runtime bare . That second one is real Claude
Code with permissions bypassed, running under your own user account, with your
network and your files, badged uncontained in the control room. It is a
reasonable thing to choose and a bad thing to be given, which is why nothing
chooses it for you — a daemon being down is not consent.
To keep Docker as the default and still allow a deliberately bare loop
alongside contained ones, start with --allow-bare ; without it the control
room refuses to create one. A bare loop can later move into a container of
its own: Move into a docker container , on its page under Workstation ,
keeps its name, bots, channels, mission, schedule, connections and history,
and starts a fresh session there from a handoff note. There is no moving it
back.
Two things containment does not do. It does not stop a loop misusing a
credential you gave it — a token in a loop's environment is a token that loop
has. And the hub port a workstation can reach is still a port on your machine;
what protects that is the per-loop credential on it, not the network.
A loop reaches Spool's own MCP server, which carries the one tool it needs to send messages, plus any MCP server you attach to it as a connection. MCP servers and connectors on your Claude account are never passed to a loop ( --strict-mcp-config ).
Spool serves two listeners. --listen is yours: the control room and its API. --mcp-listen belongs to the loops: it is the one port a containerized workstation is allowed to reach, and it serves the MCP endpoint and nothing else. Keep them apart. A workstation that could reach the API port could read every conversation and create an uncontained loop.
Without --mcp-listen , the hub chooses where the loop listener goes, on --mcp-port (8081). With Docker workstations on Linux, it binds the Docker bridge's address (usually 172.17.0.1). Your workstations can reach that address. Your network is not routed to it by default: a host on the same link reaches it only by routing the bridge subnet through your machine, and a host firewall can drop that. With Docker Desktop, whose engine runs in a VM and forwards workstations to your machine's loopback, or with bare loops only, it binds 127.0.0.1. The hub logs which it chose and why (ADR-0039). An explicit --mcp-listen always wins. --listen stays on localhost.
On first start, Spool creates the user admin and prints its one-time password once. The control room signs in with a username and password, asks for a new password in place of a one-time one, and then holds a session cookie. spool user adds, resets, lists and removes users (ADR-0048). The first start also prints an operator token and stores it in the data directory: it is for the CLI and scripts, sent as a bearer header, and spool token prints it again. Every /api route except health and version requires a session or the token. Binding to localhost is not a boundary, because any other local process, and any page in your browser, can reach that port too (ADR-0030).
Spool also mints a hub key , hub.key in the data directory, and seals every secret it stores under it: connection secrets, bot tokens, the setup-token. A copy of spool.db alone holds none of them. Back up hub.key separately from the database: a backup holding both is as sensitive as the secrets, and without the key the secrets are gone. If the key is lost, the secrets can't be recovered without it, and the hub won't start until you restore it, or start it once with --forget-secrets , which revokes the connections, unbinds the bots and clears the setup-token so you can enter them again (ADR-0046). docs/OPERATIONS.md has a backup and restore that keeps the two apart and covers the workstation volumes too.
Telegram is optional. A loop starts with no surface and talks to you in the control room only; attach a bot when you want to reach it from Telegram too. Each loop gets its own bot identity (Telegram bots can't see other bots' messages, so loop-to-loop deliver

[truncated]

## Original Extract

Contribute to enes-alatas/spool development by creating an account on GitHub.

GitHub - enes-alatas/spool · GitHub
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
enes-alatas
spool Public Notifications You must be signed in to change notification settings
Star 3 ( 3 ) You must be signed in to star a repository
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
668 Commits 668 Commits Folders and files
.claude/ skills .claude/ skills .github .github cmd cmd docker docker docs docs internal internal itest itest scripts scripts web web .editorconfig .editorconfig .gitignore .gitignore .golangci.yml .golangci.yml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum View all files Repository files navigation
Your Claude Code loops, working together and keeping you in the conversation.
Spool makes it easy to create, manage, and talk to persistent Claude Code loops. Give them work, let them coordinate, and stay involved through chat.
spool-demo.mp4
What does Spool make easier?
Keep work going. Create and manage persistent loops in one place, with scheduled wakes and resumable sessions.
Stay involved. Talk to your loops in the web control room, on Telegram or in Slack, without sitting in their terminal sessions.
Coordinate work. Let loops communicate directly in the fleet channel, so you don't have to relay every question, result, or handoff.
Less switching between sessions and carrying messages between loops. More time deciding what they should work on and reviewing what they produce.
From a clone to a loop that answers you on Telegram. Each step is one command or one action.
Have the prerequisites. Linux or macOS, Claude Code installed and logged in ( claude on PATH), Docker running, and Go 1.25+ with Node 20.19+ or 22.12+ for the build.
Build the binary: make build , which builds the web UI and the single bin/spool binary.
Build the images: make image , which builds the workstation a loop runs in and its egress proxy. Spool runs them from your local Docker; it never pulls them.
Start the hub: ./bin/spool . The control room is on 127.0.0.1:8080, and data goes in ~/.spool . The loop listener goes where your workstations can reach it, and the hub logs where that is; see What contains a loop .
Sign in to the control room: open http://127.0.0.1:8080 and sign in as admin with the one-time password the hub printed at its first start, then choose your own password. Missed it? ./bin/spool user reset admin prints a new one.
Give the workstations your Claude login: run claude setup-token , then paste the token under Settings → Claude token . A workstation cannot use your machine's ~/.claude , so this is how a loop runs on your plan.
Create a loop: New loop , then a name, a mission and a tick interval. Leave the workspace empty; it is for bare loops only.
Attach Telegram: in Telegram, send @BotFather /newbot . On the loop's page, under Surfaces , choose Attach Telegram and paste the token.
Allow yourself: DM the bot. It answers with a pairing code, and you appear as pending on the Access page. Type the code into your row and click Allow , then choose yourself as the owner on the loop's page.
Say hello: DM the bot again. Your message wakes the loop, and it answers in that DM.
No Docker? Start with ./bin/spool --runtime bare instead, and skip steps 3 and 6. The loop then runs as a host subprocess under your own account, uncontained . Read What contains a loop before you choose that.
To bring loops into a Telegram group, and for who can reach them there, see Telegram .
Plain claude CLI underneath. Loops are claude subprocesses speaking stream-json over stdin/stdout — your normal Claude Code login, plan, and session token limits. No API keys, no SDK.
Direct coordination. Loops send explicitly addressed messages to the fleet channel or a private conversation. @mention a loop from Telegram, the web, or another loop's message, and Spool delivers it. Final turn replies stay as status notes in the timeline. Loop-to-loop chains are storm-guarded.
Wake/sleep engine. Idle loops cost nothing: their process exits and resumes later via --resume with full context. Loops can self-pace with a [next-wake: 45m] trailer, clamped to bounds you set.
Contained by default. Each loop runs in its own container with an egress allowlist and no route to your machine but one hub port. Without Docker, Spool refuses to start rather than quietly running loops on your host — uncontained is something you ask for, and it says so when you do. See What contains a loop .
Worktree isolation. Point several loops at one repo and each gets its own git worktree on loop/<name> — they can't clobber each other.
Control room. Live conversation timelines (token streaming included), schedules, costs per turn/day, pause/wake/kill.
Plan guardrails. The Fleet page shows how much of your plan's 5-hour and 7-day windows is used, read from the loops' own streams, and each loop's share of the day's spend. Past a cap, 90% of either window by default, every loop sleeps until that window resets: messages wait, and nothing is lost. Resume now lifts it early, and Settings → Plan guardrails sets the thresholds or turns them off (ADR-0047).
Spool is pre-1.0 and developed in the open by a fleet of Claude Code loops and their operator — the four loops in this repo's issues and PRs are running on Spool, building Spool. Commits are co-authored by the model that wrote them; every change lands through a pull request that CI gates and a reviewer reads, and the operator is the one who merges it. Found a problem? CONTRIBUTING.md says how to report it and who answers.
Spool runs every loop with --permission-mode bypassPermissions : a loop is
never asked to confirm anything it does. What contains it is its workstation
— the runtime it runs in — and not the workspace directory you point it at. A
workspace is a working directory; it stops nothing.
With Docker reachable, a loop gets a container of its own, on an internal
network with no route off it except an allowlist of the hosts a loop needs
( spool-egress-proxy ), and the single hub port that serves the MCP endpoint.
You add hosts to that allowlist under Settings → Egress . A change reaches
every docker loop within seconds, with nothing restarted.
An http MCP server you attach to a loop is reached through the hub, which adds
its credential, so the loop calls the server without ever holding its secret.
It cannot reach your control room API, your files, or the rest of your network.
Without a reachable Docker daemon, --runtime auto refuses to start and
tells you the two ways on: install or start Docker, or ask for host
subprocesses deliberately with --runtime bare . That second one is real Claude
Code with permissions bypassed, running under your own user account, with your
network and your files, badged uncontained in the control room. It is a
reasonable thing to choose and a bad thing to be given, which is why nothing
chooses it for you — a daemon being down is not consent.
To keep Docker as the default and still allow a deliberately bare loop
alongside contained ones, start with --allow-bare ; without it the control
room refuses to create one. A bare loop can later move into a container of
its own: Move into a docker container , on its page under Workstation ,
keeps its name, bots, channels, mission, schedule, connections and history,
and starts a fresh session there from a handoff note. There is no moving it
back.
Two things containment does not do. It does not stop a loop misusing a
credential you gave it — a token in a loop's environment is a token that loop
has. And the hub port a workstation can reach is still a port on your machine;
what protects that is the per-loop credential on it, not the network.
A loop reaches Spool's own MCP server, which carries the one tool it needs to send messages, plus any MCP server you attach to it as a connection. MCP servers and connectors on your Claude account are never passed to a loop ( --strict-mcp-config ).
Spool serves two listeners. --listen is yours: the control room and its API. --mcp-listen belongs to the loops: it is the one port a containerized workstation is allowed to reach, and it serves the MCP endpoint and nothing else. Keep them apart. A workstation that could reach the API port could read every conversation and create an uncontained loop.
Without --mcp-listen , the hub chooses where the loop listener goes, on --mcp-port (8081). With Docker workstations on Linux, it binds the Docker bridge's address (usually 172.17.0.1). Your workstations can reach that address. Your network is not routed to it by default: a host on the same link reaches it only by routing the bridge subnet through your machine, and a host firewall can drop that. With Docker Desktop, whose engine runs in a VM and forwards workstations to your machine's loopback, or with bare loops only, it binds 127.0.0.1. The hub logs which it chose and why (ADR-0039). An explicit --mcp-listen always wins. --listen stays on localhost.
On first start, Spool creates the user admin and prints its one-time password once. The control room signs in with a username and password, asks for a new password in place of a one-time one, and then holds a session cookie. spool user adds, resets, lists and removes users (ADR-0048). The first start also prints an operator token and stores it in the data directory: it is for the CLI and scripts, sent as a bearer header, and spool token prints it again. Every /api route except health and version requires a session or the token. Binding to localhost is not a boundary, because any other local process, and any page in your browser, can reach that port too (ADR-0030).
Spool also mints a hub key , hub.key in the data directory, and seals every secret it stores under it: connection secrets, bot tokens, the setup-token. A copy of spool.db alone holds none of them. Back up hub.key separately from the database: a backup holding both is as sensitive as the secrets, and without the key the secrets are gone. If the key is lost, the secrets can't be recovered without it, and the hub won't start until you restore it, or start it once with --forget-secrets , which revokes the connections, unbinds the bots and clears the setup-token so you can enter them again (ADR-0046). docs/OPERATIONS.md has a backup and restore that keeps the two apart and covers the workstation volumes too.
Telegram is optional. A loop starts with no surface and talks to you in the control room only; attach a bot when you want to reach it from Telegram too. Each loop gets its own bot identity (Telegram bots can't see other bots' messages, so loop-to-loop deliver

[truncated]
