---
source: "https://nfltr.xyz/blog/distributed-claude-code"
hn_url: "https://news.ycombinator.com/item?id=49888279"
title: "Show HN: How distributed Claude Code is built on a tunnelling service"
article_title: "How distributed Claude Code is built with nfltr | NFLTR Blog"
image: "https://nfltr.xyz/og.png"
author: "abbiya"
captured_at: "2026-09-29T04:34:29Z"
capture_tool: "hn-digest"
hn_id: 49888279
score: 1
comments: 0
posted_at: "2026-09-29T04:27:01Z"
tags:
  - hacker-news
---

# Show HN: How distributed Claude Code is built on a tunnelling service

- HN: [49888279](https://news.ycombinator.com/item?id=49888279)
- Source: [nfltr.xyz](https://nfltr.xyz/blog/distributed-claude-code)
- Score: 1
- Comments: 0
- Posted: 2026-09-29T04:27:01Z

## Translation

Title: Show HN: How distributed Claude Code is built on a tunnelling service
Article title: How distributed Claude Code is built with nfltr | NFLTR Blog
Description: How nfltr turns Claude Code
HN text: built a distributed claude code that can access/bring all computing resources to one view

Article text:
How distributed Claude Code is built with nfltr | NFLTR Blog
NFLTR › Blog
How distributed Claude Code is built with nfltr
Claude Code already knows how to split up work. It starts subagents, lets them run in the background, and picks up their results when they finish. That works well as long as everything the work needs is on the machine where Claude Code runs.
The dataset lives on a VM and should not leave it.
The failing service is only reachable from inside one network.
Three independent fixes would go faster on three machines at once.
The machine doing the work gets restarted halfway through.
nfltr keeps Claude Code's model and removes the one-machine limit. Here is how it is built, what went wrong along the way, and how we test it.
The idea: Claude Code's hub, made distributed
In nfltr, one Claude session is the hub. It gets a small set of tools that mirror Claude Code's own subagent model:
spawn_agent starts an agent from a brief and returns its id at once.
send_message reaches an agent. A running agent gets it in its live session; a finished one continues in the same session and workspace.
stop_agent cancels an agent. Its partial work comes back as evidence.
list_agents shows each agent's status, machine, progress and last activity.
wait_for_agents blocks until results arrive and returns them.
list_nodes shows the joined machines, their descriptions, and how many more agents each may start.
start_monitor runs a command on a machine without a model. Each line it prints becomes an event for the hub.
The difference from subagents is where the agents run: on any machine you have joined, and they keep running when the hub goes away.
nfltr orch "<goal>" starts Claude Code as the hub in your terminal. Or add the hub to your own session with claude mcp add nfltr -- nfltr mcp --toolset hub ; with nfltr orch hub install-hook , a finished agent wakes an idle session by itself.
Machines join once, with nfltr node join --max-agents 2 . A node stays idle until a spawn needs it, then launches a Claude Code agent for that task; the agent exits when the task ends. One join per machine, no configuration per task.
Your terminal or Claude Code session
the hub: decides what runs where
outbound
nfltr.xyz relay
carries messages, runs no model
outbound
laptop
the checkout, git push
agents 1/2
QA VMs ×3
nightly suites and logs
agents 1/2 each
app VM
the service, its logs
agents 0/2
database VM
private data stays here
agents 1/2
Every arrow is a connection its source opens toward the relay; no machine accepts inbound connections. Filled dots are running agents, hollow dots free slots under the machine's --max-agents . The machines are the ones from the demos.
hub
relay
node
agent
spawn_agent → id
launch an agent
start Claude Code
register (outbound)
the brief
progress
result
stored, then delivered once
agent exits, slot free
The hub calls spawn_agent with a brief and, say, a machine, and gets an agent id at once.
No agent is free there, so the hub asks that machine's node, through the relay, to launch one.
The node starts a Claude Code agent for this task.
The agent connects out to the relay and registers.
The hub sends the brief; the agent decides how to do the work.
Progress comes back; list_agents shows it.
The result, with any branch and commit pushed, comes back.
The hub stores it and marks it delivered as wait_for_agents returns it: exactly once, even across a hub restart.
The agent exits and the slot is free again.
nfltr never picks a machine, a model or a plan for you. A machine or labels named in a spawn are hard filters; among the matches, the runtime only finds a free slot. If several machines could run a monitor, start_monitor asks for one to be named rather than choosing. Model, effort and budgets pass through when set; unset, Claude Code on that machine decides.
The same goes for what we tell the model. In one campaign the hub spawned agents in only 2 of 7 passing runs. Instead of an instruction to delegate, we added facts to the tool results: free capacity, and that agents run in parallel, each with its own workspace and tokens. The choice stays with the hub.
Capabilities are opt-in per machine
What an agent may do on a machine is that machine's decision:
Without --allow-all-tools , a node's agents can edit files, but commands that need approval (shell, git push , reading a path outside the checkout) are refused, and the turn fails with tools_not_allowed , naming the flag.
Monitors run only with --allow-monitors : they use that machine's credentials.
A node clones only repositories it lists with --allow-repo .
The API key comes from the environment or the saved config, never from the command line.
Every result arrives exactly once
Agents keep running when the hub disconnects. The hub stores each completion and marks it delivered before wait_for_agents returns it, so a hub that reattaches with the same id gets everything it missed, none of it twice. Monitor events work the same way. In your own Claude Code session the hub id comes from the project directory, so reopening the project reattaches.
hub
relay
node
start_monitor
command runs, no model
turn ends; idle
a line: event #1, kept
pull (cursor)
event #1
stored; the next pull acknowledges it
wake: a new turn reads it as data
The hub calls start_monitor on a node joined with --allow-monitors .
The node runs the command; no model runs while it watches.
The session ends its turn and sits idle.
The command prints a line; the node numbers it and keeps it until acknowledged.
The hub's long-poll pull through the relay carries its cursor.
The node answers with the new event.
The hub stores it before its next pull acknowledges it, so nothing is lost or repeated.
The session wakes, and wait_for_agents returns the line as untrusted data.
A hub sees and controls only its own agents and monitors. A second session in the same project cannot share its hub; it fails and names the holder.
Agents get a clean Claude config
In real-Claude runs, agents picked up the host user's hooks and memory, and one used Claude Code's cross-session messaging to ask another local session for data. Now agents load no user settings, hooks, memory, plugins or MCP servers from the host (the workspace's project context still applies), and cross-session messaging is off. The node verifies both before any task; each is an opt-in per node.
What the relay does, and doesn't
Every node, agent and hub opens a long-lived connection out to the relay, over TLS, authenticated with your account key. No machine opens an inbound port, joins a VPN or needs a firewall change, so it works behind NAT. The relay routes each message to the agent it is addressed to, stamping the sender from the authenticated connection so nobody can speak as another agent. It runs no model, runs none of your code, and makes no decisions: it never picks a machine, a worker or a retry.
How a result survives a relay restart
Exactly-once delivery is a chain, not one component. The hub keeps its agents' state in its own durable store; each agent keeps its finished result until the hub acknowledges it; every task's events carry a gapless sequence number; and a resumed stream replays what was missed and then says so explicitly.
A result across a relay restart
hub
relay
agent
task
accepted, progress 1…n
relay restarts
keeps working
redial redial
resume after n
missed events, then "replay complete"
result
stored, ack result released
The hub dispatches a task to an agent through the relay.
The agent accepts it and streams progress, each event numbered.
The relay restarts (a deploy, a crash).
The agent keeps working; its machine has everything it needs.
Hub and agent redial at once; a spawn issued meanwhile waits for them instead of failing.
The hub resumes the task from the last event it has.
The agent replays the events the hub missed, then marks the replay complete.
The hub stores it and acknowledges; only then does the agent drop its copy and exit.
One owner per task. When processes share a task store (relay replicas, a restarted hub), each task records an owner and an epoch, and each process holds a renewed lease. A task whose owner's lease lapsed can be taken over with the next epoch, and the store rejects any write from a stale owner, so a restart or a second replica never runs a task twice or loses it. A hub restarted after kill -9 takes the dead process's tasks over at once.
No polling for capacity. The relay publishes a small feed of changes (a slot freed, a worker or node connected); a spawn waiting for a machine wakes on the event. With nothing happening, the feed costs nothing.
Holding for a reconnecting agent. A dispatch or resume for an agent that is reconnecting is held briefly rather than failed.
Artifacts. Where artifacts pass through the relay's store, they are addressed by SHA-256, checked chunk by chunk and as a whole, and resumed from the last acknowledged offset; a mismatch fails closed.
Dormant when idle. Periodic work first checks a cheap revision number and reads nothing if it has not changed. Idle after the final hub-on-nodes soak, the relay used 0.23 % CPU (2026-09-28); the production relay idled at about 100 MB and 1 % CPU (2026-09-27).
Deploys. nfltr.xyz runs on one VM. A deploy drains, restarts and waits for health before taking traffic; clients ride out the gap for up to 2 minutes, and nodes redial within a second, jittered so a fleet does not reconnect in lockstep. The site is served directly rather than through a CDN proxy, which cut long polls after 100 seconds and answered every request with an error during deploys.
Connections are encrypted to the relay with TLS. On top of that, the frames between the hub and its agents and nodes (briefs, progress, results, artifacts, monitor lines) are encrypted end to end with X25519 and AES-256-GCM by default, so the relay forwards ciphertext. That is not the whole picture, and we would rather you hear it from us:
The dashboard summary is readable by default. The hub also sends the relay a summary of each task for the dashboard: the brief, progress, the result text (up to 12 KiB), steers, usage, branch and commit, and artifact names and hashes. The relay stores it. With --dashboard-digest status the summary carries only state, timing, the machine, usage and counts, and the dashboard says the content stayed on your machine; off sends none. Without one of them, treat the relay as able to read your prompts and results.
The end-to-end keys are authenticated only with a pairing key. Without one they protect against a relay that only watches; a relay could in principle sit in the middle of the key exchange, though no code in it does. With --e2ee-key-file , a random secret you copy to your machines yourself, each end proves it holds it before anything is sent, and a relay that swaps keys is refused. It proves the peer is one of your machines, not which one.
No downgrade. An end with encryption on refuses unencrypted messages, naming the flag.
Metadata is always visible: agent and machine ids, labels and machine descriptions, timing and message sizes.
The dashboard can act on your work. The relay carries your own dashboard actions on tasks you started (answer, approve, reject, abort, steer, pause) to the hub, unless the hub runs --dashboard-commands=false , which refuses and logs them. It cannot start new work on your machines.
What does not go through the relay: your repositories and data, which stay on the machines that work with them, apart from what agents put in their answers.
The relay is exercised on every rung of the testing ladder : SIGTERM and SIGKILL at random points, deploy windows of 30–90 s, the history audit, the linearizability checks of the ownership leases, and the production canary through the real relay. Relay-side defects those runs found and fixed:
The relay kept a context for every finished request on long-lived connections: heap growth of 4.6 MB in 33 minutes, none after the fix.
An idle hub re-sent its whole dashboard su

[truncated]

## Original Extract

How nfltr turns Claude Code

built a distributed claude code that can access/bring all computing resources to one view

How distributed Claude Code is built with nfltr | NFLTR Blog
NFLTR › Blog
How distributed Claude Code is built with nfltr
Claude Code already knows how to split up work. It starts subagents, lets them run in the background, and picks up their results when they finish. That works well as long as everything the work needs is on the machine where Claude Code runs.
The dataset lives on a VM and should not leave it.
The failing service is only reachable from inside one network.
Three independent fixes would go faster on three machines at once.
The machine doing the work gets restarted halfway through.
nfltr keeps Claude Code's model and removes the one-machine limit. Here is how it is built, what went wrong along the way, and how we test it.
The idea: Claude Code's hub, made distributed
In nfltr, one Claude session is the hub. It gets a small set of tools that mirror Claude Code's own subagent model:
spawn_agent starts an agent from a brief and returns its id at once.
send_message reaches an agent. A running agent gets it in its live session; a finished one continues in the same session and workspace.
stop_agent cancels an agent. Its partial work comes back as evidence.
list_agents shows each agent's status, machine, progress and last activity.
wait_for_agents blocks until results arrive and returns them.
list_nodes shows the joined machines, their descriptions, and how many more agents each may start.
start_monitor runs a command on a machine without a model. Each line it prints becomes an event for the hub.
The difference from subagents is where the agents run: on any machine you have joined, and they keep running when the hub goes away.
nfltr orch "<goal>" starts Claude Code as the hub in your terminal. Or add the hub to your own session with claude mcp add nfltr -- nfltr mcp --toolset hub ; with nfltr orch hub install-hook , a finished agent wakes an idle session by itself.
Machines join once, with nfltr node join --max-agents 2 . A node stays idle until a spawn needs it, then launches a Claude Code agent for that task; the agent exits when the task ends. One join per machine, no configuration per task.
Your terminal or Claude Code session
the hub: decides what runs where
outbound
nfltr.xyz relay
carries messages, runs no model
outbound
laptop
the checkout, git push
agents 1/2
QA VMs ×3
nightly suites and logs
agents 1/2 each
app VM
the service, its logs
agents 0/2
database VM
private data stays here
agents 1/2
Every arrow is a connection its source opens toward the relay; no machine accepts inbound connections. Filled dots are running agents, hollow dots free slots under the machine's --max-agents . The machines are the ones from the demos.
hub
relay
node
agent
spawn_agent → id
launch an agent
start Claude Code
register (outbound)
the brief
progress
result
stored, then delivered once
agent exits, slot free
The hub calls spawn_agent with a brief and, say, a machine, and gets an agent id at once.
No agent is free there, so the hub asks that machine's node, through the relay, to launch one.
The node starts a Claude Code agent for this task.
The agent connects out to the relay and registers.
The hub sends the brief; the agent decides how to do the work.
Progress comes back; list_agents shows it.
The result, with any branch and commit pushed, comes back.
The hub stores it and marks it delivered as wait_for_agents returns it: exactly once, even across a hub restart.
The agent exits and the slot is free again.
nfltr never picks a machine, a model or a plan for you. A machine or labels named in a spawn are hard filters; among the matches, the runtime only finds a free slot. If several machines could run a monitor, start_monitor asks for one to be named rather than choosing. Model, effort and budgets pass through when set; unset, Claude Code on that machine decides.
The same goes for what we tell the model. In one campaign the hub spawned agents in only 2 of 7 passing runs. Instead of an instruction to delegate, we added facts to the tool results: free capacity, and that agents run in parallel, each with its own workspace and tokens. The choice stays with the hub.
Capabilities are opt-in per machine
What an agent may do on a machine is that machine's decision:
Without --allow-all-tools , a node's agents can edit files, but commands that need approval (shell, git push , reading a path outside the checkout) are refused, and the turn fails with tools_not_allowed , naming the flag.
Monitors run only with --allow-monitors : they use that machine's credentials.
A node clones only repositories it lists with --allow-repo .
The API key comes from the environment or the saved config, never from the command line.
Every result arrives exactly once
Agents keep running when the hub disconnects. The hub stores each completion and marks it delivered before wait_for_agents returns it, so a hub that reattaches with the same id gets everything it missed, none of it twice. Monitor events work the same way. In your own Claude Code session the hub id comes from the project directory, so reopening the project reattaches.
hub
relay
node
start_monitor
command runs, no model
turn ends; idle
a line: event #1, kept
pull (cursor)
event #1
stored; the next pull acknowledges it
wake: a new turn reads it as data
The hub calls start_monitor on a node joined with --allow-monitors .
The node runs the command; no model runs while it watches.
The session ends its turn and sits idle.
The command prints a line; the node numbers it and keeps it until acknowledged.
The hub's long-poll pull through the relay carries its cursor.
The node answers with the new event.
The hub stores it before its next pull acknowledges it, so nothing is lost or repeated.
The session wakes, and wait_for_agents returns the line as untrusted data.
A hub sees and controls only its own agents and monitors. A second session in the same project cannot share its hub; it fails and names the holder.
Agents get a clean Claude config
In real-Claude runs, agents picked up the host user's hooks and memory, and one used Claude Code's cross-session messaging to ask another local session for data. Now agents load no user settings, hooks, memory, plugins or MCP servers from the host (the workspace's project context still applies), and cross-session messaging is off. The node verifies both before any task; each is an opt-in per node.
What the relay does, and doesn't
Every node, agent and hub opens a long-lived connection out to the relay, over TLS, authenticated with your account key. No machine opens an inbound port, joins a VPN or needs a firewall change, so it works behind NAT. The relay routes each message to the agent it is addressed to, stamping the sender from the authenticated connection so nobody can speak as another agent. It runs no model, runs none of your code, and makes no decisions: it never picks a machine, a worker or a retry.
How a result survives a relay restart
Exactly-once delivery is a chain, not one component. The hub keeps its agents' state in its own durable store; each agent keeps its finished result until the hub acknowledges it; every task's events carry a gapless sequence number; and a resumed stream replays what was missed and then says so explicitly.
A result across a relay restart
hub
relay
agent
task
accepted, progress 1…n
relay restarts
keeps working
redial redial
resume after n
missed events, then "replay complete"
result
stored, ack result released
The hub dispatches a task to an agent through the relay.
The agent accepts it and streams progress, each event numbered.
The relay restarts (a deploy, a crash).
The agent keeps working; its machine has everything it needs.
Hub and agent redial at once; a spawn issued meanwhile waits for them instead of failing.
The hub resumes the task from the last event it has.
The agent replays the events the hub missed, then marks the replay complete.
The hub stores it and acknowledges; only then does the agent drop its copy and exit.
One owner per task. When processes share a task store (relay replicas, a restarted hub), each task records an owner and an epoch, and each process holds a renewed lease. A task whose owner's lease lapsed can be taken over with the next epoch, and the store rejects any write from a stale owner, so a restart or a second replica never runs a task twice or loses it. A hub restarted after kill -9 takes the dead process's tasks over at once.
No polling for capacity. The relay publishes a small feed of changes (a slot freed, a worker or node connected); a spawn waiting for a machine wakes on the event. With nothing happening, the feed costs nothing.
Holding for a reconnecting agent. A dispatch or resume for an agent that is reconnecting is held briefly rather than failed.
Artifacts. Where artifacts pass through the relay's store, they are addressed by SHA-256, checked chunk by chunk and as a whole, and resumed from the last acknowledged offset; a mismatch fails closed.
Dormant when idle. Periodic work first checks a cheap revision number and reads nothing if it has not changed. Idle after the final hub-on-nodes soak, the relay used 0.23 % CPU (2026-09-28); the production relay idled at about 100 MB and 1 % CPU (2026-09-27).
Deploys. nfltr.xyz runs on one VM. A deploy drains, restarts and waits for health before taking traffic; clients ride out the gap for up to 2 minutes, and nodes redial within a second, jittered so a fleet does not reconnect in lockstep. The site is served directly rather than through a CDN proxy, which cut long polls after 100 seconds and answered every request with an error during deploys.
Connections are encrypted to the relay with TLS. On top of that, the frames between the hub and its agents and nodes (briefs, progress, results, artifacts, monitor lines) are encrypted end to end with X25519 and AES-256-GCM by default, so the relay forwards ciphertext. That is not the whole picture, and we would rather you hear it from us:
The dashboard summary is readable by default. The hub also sends the relay a summary of each task for the dashboard: the brief, progress, the result text (up to 12 KiB), steers, usage, branch and commit, and artifact names and hashes. The relay stores it. With --dashboard-digest status the summary carries only state, timing, the machine, usage and counts, and the dashboard says the content stayed on your machine; off sends none. Without one of them, treat the relay as able to read your prompts and results.
The end-to-end keys are authenticated only with a pairing key. Without one they protect against a relay that only watches; a relay could in principle sit in the middle of the key exchange, though no code in it does. With --e2ee-key-file , a random secret you copy to your machines yourself, each end proves it holds it before anything is sent, and a relay that swaps keys is refused. It proves the peer is one of your machines, not which one.
No downgrade. An end with encryption on refuses unencrypted messages, naming the flag.
Metadata is always visible: agent and machine ids, labels and machine descriptions, timing and message sizes.
The dashboard can act on your work. The relay carries your own dashboard actions on tasks you started (answer, approve, reject, abort, steer, pause) to the hub, unless the hub runs --dashboard-commands=false , which refuses and logs them. It cannot start new work on your machines.
What does not go through the relay: your repositories and data, which stay on the machines that work with them, apart from what agents put in their answers.
The relay is exercised on every rung of the testing ladder : SIGTERM and SIGKILL at random points, deploy windows of 30–90 s, the history audit, the linearizability checks of the ownership leases, and the production canary through the real relay. Relay-side defects those runs found and fixed:
The relay kept a context for every finished request on long-lived connections: heap growth of 4.6 MB in 33 minutes, none after the fix.
An idle hub re-sent its whole dashboard su

[truncated]
