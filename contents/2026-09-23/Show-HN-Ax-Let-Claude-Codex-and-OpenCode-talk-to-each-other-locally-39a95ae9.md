---
source: "https://useax.dev/"
hn_url: "https://news.ycombinator.com/item?id=49816738"
title: "Show HN: Ax – Let Claude, Codex and OpenCode talk to each other locally"
article_title: "AX — Put your agents on speaking terms"
image: ""
author: "rcdexta"
captured_at: "2026-09-23T15:19:58Z"
capture_tool: "hn-digest"
hn_id: 49816738
score: 3
comments: 0
posted_at: "2026-09-23T14:27:29Z"
tags:
  - hacker-news
---

# Show HN: Ax – Let Claude, Codex and OpenCode talk to each other locally

- HN: [49816738](https://news.ycombinator.com/item?id=49816738)
- Source: [useax.dev](https://useax.dev/)
- Score: 3
- Comments: 0
- Posted: 2026-09-23T14:27:29Z

## Translation

Title: Show HN: Ax – Let Claude, Codex and OpenCode talk to each other locally
Article title: AX — Put your agents on speaking terms
Description: Give your coding agents names. Let Claude Code, Codex, Grok Build, OpenCode, and Pi talk to each other with Agent Exchange.
HN text: Hi HN, I’m RC. When I worked at ThoughtWorks, we were pretty religious about pair programming: one person driving, another reviewing, both working through the same problem. Most of us have already paired with an agent in our own coding sessions. Some of us have even tried putting agents in both seats: one implementing, the other reviewing. I really like Claude Code’s cross-session messaging. I wanted that same experience between Claude, Codex, and OpenCode. The existing setups meant installing a new app or
plugins that just don't work reliably. I wanted something local that stayed out of the way, let me keep using each tool as usual, but enabled hetergenous messaging. That became Agent Exchange, or ax. After installing, start two sessions in separate terminals: ax claude -n api
ax codex -n web
Then ask Codex: “Ask api whether the schema is ready.” The agents address each other by name, and replies come back into their conversations, including across repositories on the same machine under the same OS user. You can also connect existing conversations using each tool’s native resume syntax. AX supports Claude Code, Codex CLI, Grok Build, OpenCode, and Pi. Underneath, it’s a local Go broker, Unix sockets, and a SQLite mailbox. There’s no AX cloud relay or telemetry. The messaging runs locally; the coding agents still communicate with their configured AI providers. We’ve been using it to send AX’s own PRs between Codex and Grok for review. I’d be interested in which agent combinations you use, and what you currently have to relay between them yourself.

Article text:
Skip to content
ax ↗ Agent
Exchange
Docs GitHub ↗ Get started
A local exchange for coding agents
Your agents.
On speaking terms.
Claude handles the API. Codex builds the UI. AX gets them talking.
Fast, Reliable, Local Routing only
AX’s broker and mailbox run on your machine. No cloud relay or telemetry.
Your agents still send content to their AI providers.
Read https://useax.dev/agents.md and install Agent Exchange.
Copy the full guide into your agent, or paste the prompt above.
curl -fsSL https://useax.dev/install.sh | sh
macOS, Linux, and WSL 2. Installation guide ↗
› Build signup. Coordinate with api.
└ to: api
Which signup fields?
• api: at least 12 characters.
• Form and validation are ready.
• api: looks good. Handle 409 for an email already in use.
• Added the error. Tests pass.
└ to: api
409 handled. UI tests pass.
└ Looks good. Handle 409 for an email already in use.
Illustrative exchange · scroll either pane to read back
Claude builds the API. OpenCode builds the UI. Grok reviews the changes. Give each a name so they can coordinate across their own sessions.
curl -fsSL https://useax.dev/install.sh | sh
export PATH="$HOME/.local/bin:$PATH"
ax doctor Platform requirements and setup ↗
Run each command in a separate terminal.
ax claude -n api OpenCode · UI ax opencode -n web Grok Build · Review ax grok -n reviewer
“Ask api for the response format. Build the UI, then ask reviewer to check your changes.”
OpenCode gets the API details from Claude and sends its changes to Grok for review. Replies return to the same sessions.
Keep the tools.
Connect the work.
Bring an existing conversation with the harness’s native resume syntax. The name stays attached to that conversation.
Agent names work across directories for the same OS user. Your API and frontend can live in different repositories.
A delegated task keeps its scope. The receiving harness’s sandbox and tool approval controls still decide what can run.
Your agents do the work.
AX carries the messages.
Already using OpenCode or a terminal workspace like Herdr? Here’s the part each tool handles.
Writes and reviews code, with primary agents and subagents inside OpenCode.
AX lets an OpenCode session exchange tasks and replies with named Claude, Codex, Grok, and Pi sessions.
Organizes workspaces and panes, tracks agent state, and provides controls for automating agents.
AX focuses on addressing messages to agents by name, storing them locally, and tracking delivery.
Connects independent coding sessions across supported harnesses on your machine.
Ask one agent to contact another. Each keeps its conversation and native tool permissions; replies come back through AX.
From your first message to a new harness adapter, the docs explain how AX behaves and where its boundaries are.
Give your agents
someone to talk to.
Agent Exchange · Built by the summation.com team.

## Original Extract

Give your coding agents names. Let Claude Code, Codex, Grok Build, OpenCode, and Pi talk to each other with Agent Exchange.

Hi HN, I’m RC. When I worked at ThoughtWorks, we were pretty religious about pair programming: one person driving, another reviewing, both working through the same problem. Most of us have already paired with an agent in our own coding sessions. Some of us have even tried putting agents in both seats: one implementing, the other reviewing. I really like Claude Code’s cross-session messaging. I wanted that same experience between Claude, Codex, and OpenCode. The existing setups meant installing a new app or
plugins that just don't work reliably. I wanted something local that stayed out of the way, let me keep using each tool as usual, but enabled hetergenous messaging. That became Agent Exchange, or ax. After installing, start two sessions in separate terminals: ax claude -n api
ax codex -n web
Then ask Codex: “Ask api whether the schema is ready.” The agents address each other by name, and replies come back into their conversations, including across repositories on the same machine under the same OS user. You can also connect existing conversations using each tool’s native resume syntax. AX supports Claude Code, Codex CLI, Grok Build, OpenCode, and Pi. Underneath, it’s a local Go broker, Unix sockets, and a SQLite mailbox. There’s no AX cloud relay or telemetry. The messaging runs locally; the coding agents still communicate with their configured AI providers. We’ve been using it to send AX’s own PRs between Codex and Grok for review. I’d be interested in which agent combinations you use, and what you currently have to relay between them yourself.

Skip to content
ax ↗ Agent
Exchange
Docs GitHub ↗ Get started
A local exchange for coding agents
Your agents.
On speaking terms.
Claude handles the API. Codex builds the UI. AX gets them talking.
Fast, Reliable, Local Routing only
AX’s broker and mailbox run on your machine. No cloud relay or telemetry.
Your agents still send content to their AI providers.
Read https://useax.dev/agents.md and install Agent Exchange.
Copy the full guide into your agent, or paste the prompt above.
curl -fsSL https://useax.dev/install.sh | sh
macOS, Linux, and WSL 2. Installation guide ↗
› Build signup. Coordinate with api.
└ to: api
Which signup fields?
• api: at least 12 characters.
• Form and validation are ready.
• api: looks good. Handle 409 for an email already in use.
• Added the error. Tests pass.
└ to: api
409 handled. UI tests pass.
└ Looks good. Handle 409 for an email already in use.
Illustrative exchange · scroll either pane to read back
Claude builds the API. OpenCode builds the UI. Grok reviews the changes. Give each a name so they can coordinate across their own sessions.
curl -fsSL https://useax.dev/install.sh | sh
export PATH="$HOME/.local/bin:$PATH"
ax doctor Platform requirements and setup ↗
Run each command in a separate terminal.
ax claude -n api OpenCode · UI ax opencode -n web Grok Build · Review ax grok -n reviewer
“Ask api for the response format. Build the UI, then ask reviewer to check your changes.”
OpenCode gets the API details from Claude and sends its changes to Grok for review. Replies return to the same sessions.
Keep the tools.
Connect the work.
Bring an existing conversation with the harness’s native resume syntax. The name stays attached to that conversation.
Agent names work across directories for the same OS user. Your API and frontend can live in different repositories.
A delegated task keeps its scope. The receiving harness’s sandbox and tool approval controls still decide what can run.
Your agents do the work.
AX carries the messages.
Already using OpenCode or a terminal workspace like Herdr? Here’s the part each tool handles.
Writes and reviews code, with primary agents and subagents inside OpenCode.
AX lets an OpenCode session exchange tasks and replies with named Claude, Codex, Grok, and Pi sessions.
Organizes workspaces and panes, tracks agent state, and provides controls for automating agents.
AX focuses on addressing messages to agents by name, storing them locally, and tracking delivery.
Connects independent coding sessions across supported harnesses on your machine.
Ask one agent to contact another. Each keeps its conversation and native tool permissions; replies come back through AX.
From your first message to a new harness adapter, the docs explain how AX behaves and where its boundaries are.
Give your agents
someone to talk to.
Agent Exchange · Built by the summation.com team.
