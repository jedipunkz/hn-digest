---
source: "https://platform.claude.com/docs/en/managed-agents/multiagent-orchestration"
hn_url: "https://news.ycombinator.com/item?id=50031369"
title: "Claude agents can now write their own multiagent workflows"
article_title: "Multiagent orchestration - Claude Platform Docs"
image: "https://platform.claude.com/web-api/og/docs/en/managed-agents/multiagent-orchestration?design-rev=3"
author: "joshcsimmons"
captured_at: "2026-10-10T10:58:09Z"
capture_tool: "hn-digest"
hn_id: 50031369
score: 1
comments: 0
posted_at: "2026-10-10T10:13:21Z"
tags:
  - hacker-news
---

# Claude agents can now write their own multiagent workflows

- HN: [50031369](https://news.ycombinator.com/item?id=50031369)
- Source: [platform.claude.com](https://platform.claude.com/docs/en/managed-agents/multiagent-orchestration)
- Score: 1
- Comments: 0
- Posted: 2026-10-10T10:13:21Z

## Translation

Title: Claude agents can now write their own multiagent workflows
Article title: Multiagent orchestration - Claude Platform Docs
Description: Coordinate multiple agents within a single session.

Article text:
Multiagent orchestration - Claude Platform Docs Claude Platform Docs Messages
Coordinate multiple agents within a single session.
Multiagent orchestration lets one agent coordinate with others to complete complex work. Agents can act in parallel with their own isolated context, which helps improve output quality and can also improve time to completion.
Not sure a multiagent setup fits your problem? See when to use multiagent systems (and when not to) .
The agent that a session runs can hand work to other agents in two ways. With subagents , it delegates tasks itself and reads what each subagent reports. With dynamic workflows , it writes a workflow: a program that runs many agents in the background and combines their results. It can also consult an advisor model for guidance while it does the work itself.
You decide which of these the agent can use, and the agent determines when to use them. To guide that choice, tell the agent in its system prompt when to use a workflow run. See Tell the agent when to use a run . You can also limit the agent to agents that you list.
With subagents, the agent itself determines what happens next. A subagent's thread stays available until you archive it, so the agent can send it follow-up messages. With dynamic workflows, Claude writes a program to orchestrate agents without Claude's direct involvement. Context and results are passed programmatically from one agent to another, freeing up the main session thread to communicate with the user and check in on one or several running workflows to report on progress. The agent can't send follow-up messages to a run's threads, and the server archives each one by the end of its run.
You set these up in the multiagent block of the agent's definition, which has a type . With the multiagent_20261001 type, an agent can use all three together, and you can turn each one on or off. By default, subagents and workflows are both enabled. subagents and workflows each have inline_agents enabled, the setting for agents that the agent or a workflow defines itself :
{
"multiagent" : { "type" : "multiagent_20261001" }
} 
To set which agents the agent can call, see Predefined and inline agents . To turn a setting off, see Turn on dynamic workflows .
Predefined and inline agents 
Every subagent, and every agent in a workflow run, is one of two kinds:
Predefined agent: An agent that you have already created , and that you list in the multiagent block. It uses its own configuration: model, system prompt, tools, MCP servers, and skills.
Inline agent: An agent that is not saved. The agent that the session runs, or a workflow, defines it when it hands out the work. It uses that agent's model, tools, MCP servers, and skills.
Both kinds work under subagents and under workflows :
subagents.predefined_agents and workflows.predefined_agents are two separate lists. An agent in one list is not added to the other. Both lists are empty by default. An entry of either list takes one of these forms:
{"type": "agent", "id": agent.id} references a previously created agent by ID. If no version is specified, the reference is pinned to the agent's latest version when the agent that lists it is created, or when an update sends the list.
{"type": "agent", "id": agent.id, "version": agent.version} pins a specific agent version.
agent.id alone, as a string, is short for {"type": "agent", "id": agent.id} .
{"type": "self"} lists the agent itself, so that copies of it can do the work. If the session was created with agent configuration overrides , those overrides also apply to these copies. Entries referenced by ID are unaffected.
The rules for these entries, and for the agents that they name, apply to both lists. See List the subagents .
Inline agents are on by default under both settings. To allow only the agents that you list, set inline_agents to {"type": "disabled"} under subagents , under workflows , or under both. A setting with inline agents off needs at least one agent in its predefined_agents list. With an empty list, the request fails with a 400 error. On an update, the server checks the settings as they stand after the update.
The following agent allows only the agents that it lists. It can delegate to one agent and to copies of itself, and a workflow can use version 2 of another agent:
{
"multiagent" : {
"type" : "multiagent_20261001" ,
"subagents" : {
"type" : "enabled" ,
"inline_agents" : { "type" : "disabled" },
"predefined_agents" : [ "agent_01J8XkN5uT3vHpLqRfWdY2" , { "type" : "self" }]
},
"workflows" : {
"type" : "enabled" ,
"inline_agents" : { "type" : "disabled" },
"predefined_agents" : [
{ "type" : "agent" , "id" : "agent_01Lm4cV8yQ2tNs7XbKdR5h" , "version" : 2 }
]
}
}
} 
Delegate to subagents 
Multiagent coordination is best suited for complex tasks that either require work across a variety of surfaces, or where multiple well-scoped tasks contribute to an overall goal.
Parallelization: Fan out independent subtasks simultaneously (searching multiple sources, analyzing separate files) and have the agent synthesize the results.
Specialization: Route to agents with domain-focused system prompts and tools, such as a security agent or a documentation agent, rather than loading a single agent with every capability.
Escalation: Consult a more capable agent or model for a subset of complex subtasks. To consult a model, give the session an advisor .
All agents share the same sandbox, filesystem, and vault credentials , but each agent runs in its own session thread , a context-isolated event stream with its own conversation history. The agent that the session runs reports activity in the primary thread , which is the session-level event stream . Additional threads are spawned at runtime when it delegates work. A workflow run also creates threads.
A subagent's thread is persistent. The agent can send a follow-up to a subagent it called earlier, and that subagent retains everything from its previous turns.
Which configuration a subagent uses depends on whether it is a predefined or an inline agent . Session-level agent configuration overrides apply to the agent that the session runs and to its self copies. Each agent keeps its own conversation history.
When defining your agent , set subagents.predefined_agents in the multiagent block to list the agents that it can delegate to:
# Create the subagents, then read their IDs from the lockfile.
ant apply reviewer.md test-writer.md
REVIEWER_AGENT_ID = $( jq -er '.resources["./reviewer.md"].id' claude-lock.json )
TEST_WRITER_AGENT_ID = $( jq -er '.resources["./test-writer.md"].id' claude-lock.json )
# Write the agent's definition, listing each subagent by ID.
cat > engineering-lead.md << EOF
---
name: Engineering Lead
model: claude-opus-5-5
tools:
- type: agent_toolset_20260401
multiagent:
type: multiagent_20261001
subagents:
type: enabled
predefined_agents:
- type: agent
id: $REVIEWER_AGENT_ID
- type: agent
id: $TEST_WRITER_AGENT_ID
---
You coordinate engineering work. Delegate code review to the reviewer agent
and test writing to the test agent.
EOF
# Create the agent.
ant apply engineering-lead.md reviewer.md test-writer.md reviewer.md   ---
name : reviewer
model : claude-haiku-5-5
---
You are a code reviewer. test-writer.md   ---
name : test-writer
model : claude-haiku-5-5
---
You write unit tests.
The agent can also delegate to inline agents unless you turn them off. For that setting, and for the forms that an entry of subagents.predefined_agents takes, see Predefined and inline agents .
"advisor": {"type": "enabled", "model": "<model id>"} gives the session's primary thread an advisor it can consult mid-turn. The advisor is a setting in the multiagent block, not an entry in this list. See Give the session an advisor .
Dynamic workflows are also enabled by default with this type, so this agent can plan large work that runs many agents in the background. To turn them off, or to list agents that a workflow can use, see Turn on dynamic workflows .
The following rules apply to the agents you list in subagents.predefined_agents , and also to the agents you list in workflows.predefined_agents :
Pinning: The agent's configuration, including its subagents.predefined_agents list, is snapshotted when the agent is created or updated. Referenced agents stay pinned to the versions resolved then and don't pick up later updates to their definitions. An update that doesn't send the list keeps the versions that are already pinned. To delegate to a newer version of a referenced agent, update the agent so its subagents.predefined_agents list references that version.
One level: The agent can delegate to only one level of agents. Referencing another agent that has multiagent set fails the create or update request with a 400 validation error.
Up to 20 agents: subagents.predefined_agents can list up to 20 unique agents. The limit is per list: workflows.predefined_agents can also list up to 20. The agent can call multiple copies of each agent, within the session's thread limit .
Inference geography: The agent and every agent that you list, in subagents.predefined_agents or in workflows.predefined_agents , must pin the same inference geography ( model.inference_geo in the agent definition ), or none of them can pin one. A mismatch in either list is rejected with a 400 validation error. That check runs both when the agent is saved and when a session-create override changes any of the pins.
Create a session referencing the agent. The agent delegates to the agents you list in subagents.predefined_agents as needed. It can also delegate to inline agents unless you turn them off.
session = client.beta.sessions.create(
agent = lead_agent.id,
environment_id = environment.id,
)
Connect agents to MCP servers 
MCP servers are agent-scoped: each agent definition declares its own servers and tools. An inline agent has no agent definition, so it uses the MCP servers and tools of the agent that the session runs. Vault credentials are session-scoped: vault_ids passed at session creation apply to every thread. Two implications for your integration:
To authenticate MCP servers, include a vault credential for every MCP server used across all agents.
To limit an agent's access, declare only the servers it needs in its agent definition. You can't limit an inline agent this way. To allow only the agents you list, disable inline_agents in both subagents and workflows , and list at least one agent in each. See Predefined and inline agents .
Agent configuration overrides at session creation can replace the MCP servers of the agent that the session runs and those of its self copies.
With a limited environment , session creation fails with a 400 error when the agent, or an agent you list in subagents.predefined_agents or workflows.predefined_agents , declares an MCP server whose host is not in allowed_hosts . Setting allow_mcp_servers: true in the environment's networking turns this check off.
Create the researcher, which declares the GitHub MCP server, and the agent that delegates to the researcher:
# Create the researcher, then read its ID from the lockfile.
ant apply researcher.md
research_agent_id = $( jq -er '.resources["./researcher.md"].id' claude-lock.json )
# Write the agent's definition, listing the researcher by ID.
cat > lead.md << EOF
---
name: lead
model: claude-opus-5-5
tools:
- type: agent_toolset_20260401
multiagent:
type: multiagent_20261001
subagents:
type: enabled
predefined_agents:
- type: agent
id: $research_agent_id
---
EOF
# Create the agent.
ant apply lead.md researcher.md researcher.md   ---
name : researcher
model : claude-haiku-5-5
mcp_servers :
- type : url
name : github
url : https://api.githubcopilot.com/mcp/
tools :
- type : mcp_toolset
mcp_server_name : github
---
Then create the session with the vault that holds the GitHub credential:
session = client.beta.sessions.create(
agent = lead_agent.id,
environment_id = environment.id,
vault_ids = [vault.id],
)
print (session.id)
In this example, onl

[truncated]

## Original Extract

Coordinate multiple agents within a single session.

Multiagent orchestration - Claude Platform Docs Claude Platform Docs Messages
Coordinate multiple agents within a single session.
Multiagent orchestration lets one agent coordinate with others to complete complex work. Agents can act in parallel with their own isolated context, which helps improve output quality and can also improve time to completion.
Not sure a multiagent setup fits your problem? See when to use multiagent systems (and when not to) .
The agent that a session runs can hand work to other agents in two ways. With subagents , it delegates tasks itself and reads what each subagent reports. With dynamic workflows , it writes a workflow: a program that runs many agents in the background and combines their results. It can also consult an advisor model for guidance while it does the work itself.
You decide which of these the agent can use, and the agent determines when to use them. To guide that choice, tell the agent in its system prompt when to use a workflow run. See Tell the agent when to use a run . You can also limit the agent to agents that you list.
With subagents, the agent itself determines what happens next. A subagent's thread stays available until you archive it, so the agent can send it follow-up messages. With dynamic workflows, Claude writes a program to orchestrate agents without Claude's direct involvement. Context and results are passed programmatically from one agent to another, freeing up the main session thread to communicate with the user and check in on one or several running workflows to report on progress. The agent can't send follow-up messages to a run's threads, and the server archives each one by the end of its run.
You set these up in the multiagent block of the agent's definition, which has a type . With the multiagent_20261001 type, an agent can use all three together, and you can turn each one on or off. By default, subagents and workflows are both enabled. subagents and workflows each have inline_agents enabled, the setting for agents that the agent or a workflow defines itself :
{
"multiagent" : { "type" : "multiagent_20261001" }
} 
To set which agents the agent can call, see Predefined and inline agents . To turn a setting off, see Turn on dynamic workflows .
Predefined and inline agents 
Every subagent, and every agent in a workflow run, is one of two kinds:
Predefined agent: An agent that you have already created , and that you list in the multiagent block. It uses its own configuration: model, system prompt, tools, MCP servers, and skills.
Inline agent: An agent that is not saved. The agent that the session runs, or a workflow, defines it when it hands out the work. It uses that agent's model, tools, MCP servers, and skills.
Both kinds work under subagents and under workflows :
subagents.predefined_agents and workflows.predefined_agents are two separate lists. An agent in one list is not added to the other. Both lists are empty by default. An entry of either list takes one of these forms:
{"type": "agent", "id": agent.id} references a previously created agent by ID. If no version is specified, the reference is pinned to the agent's latest version when the agent that lists it is created, or when an update sends the list.
{"type": "agent", "id": agent.id, "version": agent.version} pins a specific agent version.
agent.id alone, as a string, is short for {"type": "agent", "id": agent.id} .
{"type": "self"} lists the agent itself, so that copies of it can do the work. If the session was created with agent configuration overrides , those overrides also apply to these copies. Entries referenced by ID are unaffected.
The rules for these entries, and for the agents that they name, apply to both lists. See List the subagents .
Inline agents are on by default under both settings. To allow only the agents that you list, set inline_agents to {"type": "disabled"} under subagents , under workflows , or under both. A setting with inline agents off needs at least one agent in its predefined_agents list. With an empty list, the request fails with a 400 error. On an update, the server checks the settings as they stand after the update.
The following agent allows only the agents that it lists. It can delegate to one agent and to copies of itself, and a workflow can use version 2 of another agent:
{
"multiagent" : {
"type" : "multiagent_20261001" ,
"subagents" : {
"type" : "enabled" ,
"inline_agents" : { "type" : "disabled" },
"predefined_agents" : [ "agent_01J8XkN5uT3vHpLqRfWdY2" , { "type" : "self" }]
},
"workflows" : {
"type" : "enabled" ,
"inline_agents" : { "type" : "disabled" },
"predefined_agents" : [
{ "type" : "agent" , "id" : "agent_01Lm4cV8yQ2tNs7XbKdR5h" , "version" : 2 }
]
}
}
} 
Delegate to subagents 
Multiagent coordination is best suited for complex tasks that either require work across a variety of surfaces, or where multiple well-scoped tasks contribute to an overall goal.
Parallelization: Fan out independent subtasks simultaneously (searching multiple sources, analyzing separate files) and have the agent synthesize the results.
Specialization: Route to agents with domain-focused system prompts and tools, such as a security agent or a documentation agent, rather than loading a single agent with every capability.
Escalation: Consult a more capable agent or model for a subset of complex subtasks. To consult a model, give the session an advisor .
All agents share the same sandbox, filesystem, and vault credentials , but each agent runs in its own session thread , a context-isolated event stream with its own conversation history. The agent that the session runs reports activity in the primary thread , which is the session-level event stream . Additional threads are spawned at runtime when it delegates work. A workflow run also creates threads.
A subagent's thread is persistent. The agent can send a follow-up to a subagent it called earlier, and that subagent retains everything from its previous turns.
Which configuration a subagent uses depends on whether it is a predefined or an inline agent . Session-level agent configuration overrides apply to the agent that the session runs and to its self copies. Each agent keeps its own conversation history.
When defining your agent , set subagents.predefined_agents in the multiagent block to list the agents that it can delegate to:
# Create the subagents, then read their IDs from the lockfile.
ant apply reviewer.md test-writer.md
REVIEWER_AGENT_ID = $( jq -er '.resources["./reviewer.md"].id' claude-lock.json )
TEST_WRITER_AGENT_ID = $( jq -er '.resources["./test-writer.md"].id' claude-lock.json )
# Write the agent's definition, listing each subagent by ID.
cat > engineering-lead.md << EOF
---
name: Engineering Lead
model: claude-opus-5-5
tools:
- type: agent_toolset_20260401
multiagent:
type: multiagent_20261001
subagents:
type: enabled
predefined_agents:
- type: agent
id: $REVIEWER_AGENT_ID
- type: agent
id: $TEST_WRITER_AGENT_ID
---
You coordinate engineering work. Delegate code review to the reviewer agent
and test writing to the test agent.
EOF
# Create the agent.
ant apply engineering-lead.md reviewer.md test-writer.md reviewer.md   ---
name : reviewer
model : claude-haiku-5-5
---
You are a code reviewer. test-writer.md   ---
name : test-writer
model : claude-haiku-5-5
---
You write unit tests.
The agent can also delegate to inline agents unless you turn them off. For that setting, and for the forms that an entry of subagents.predefined_agents takes, see Predefined and inline agents .
"advisor": {"type": "enabled", "model": "<model id>"} gives the session's primary thread an advisor it can consult mid-turn. The advisor is a setting in the multiagent block, not an entry in this list. See Give the session an advisor .
Dynamic workflows are also enabled by default with this type, so this agent can plan large work that runs many agents in the background. To turn them off, or to list agents that a workflow can use, see Turn on dynamic workflows .
The following rules apply to the agents you list in subagents.predefined_agents , and also to the agents you list in workflows.predefined_agents :
Pinning: The agent's configuration, including its subagents.predefined_agents list, is snapshotted when the agent is created or updated. Referenced agents stay pinned to the versions resolved then and don't pick up later updates to their definitions. An update that doesn't send the list keeps the versions that are already pinned. To delegate to a newer version of a referenced agent, update the agent so its subagents.predefined_agents list references that version.
One level: The agent can delegate to only one level of agents. Referencing another agent that has multiagent set fails the create or update request with a 400 validation error.
Up to 20 agents: subagents.predefined_agents can list up to 20 unique agents. The limit is per list: workflows.predefined_agents can also list up to 20. The agent can call multiple copies of each agent, within the session's thread limit .
Inference geography: The agent and every agent that you list, in subagents.predefined_agents or in workflows.predefined_agents , must pin the same inference geography ( model.inference_geo in the agent definition ), or none of them can pin one. A mismatch in either list is rejected with a 400 validation error. That check runs both when the agent is saved and when a session-create override changes any of the pins.
Create a session referencing the agent. The agent delegates to the agents you list in subagents.predefined_agents as needed. It can also delegate to inline agents unless you turn them off.
session = client.beta.sessions.create(
agent = lead_agent.id,
environment_id = environment.id,
)
Connect agents to MCP servers 
MCP servers are agent-scoped: each agent definition declares its own servers and tools. An inline agent has no agent definition, so it uses the MCP servers and tools of the agent that the session runs. Vault credentials are session-scoped: vault_ids passed at session creation apply to every thread. Two implications for your integration:
To authenticate MCP servers, include a vault credential for every MCP server used across all agents.
To limit an agent's access, declare only the servers it needs in its agent definition. You can't limit an inline agent this way. To allow only the agents you list, disable inline_agents in both subagents and workflows , and list at least one agent in each. See Predefined and inline agents .
Agent configuration overrides at session creation can replace the MCP servers of the agent that the session runs and those of its self copies.
With a limited environment , session creation fails with a 400 error when the agent, or an agent you list in subagents.predefined_agents or workflows.predefined_agents , declares an MCP server whose host is not in allowed_hosts . Setting allow_mcp_servers: true in the environment's networking turns this check off.
Create the researcher, which declares the GitHub MCP server, and the agent that delegates to the researcher:
# Create the researcher, then read its ID from the lockfile.
ant apply researcher.md
research_agent_id = $( jq -er '.resources["./researcher.md"].id' claude-lock.json )
# Write the agent's definition, listing the researcher by ID.
cat > lead.md << EOF
---
name: lead
model: claude-opus-5-5
tools:
- type: agent_toolset_20260401
multiagent:
type: multiagent_20261001
subagents:
type: enabled
predefined_agents:
- type: agent
id: $research_agent_id
---
EOF
# Create the agent.
ant apply lead.md researcher.md researcher.md   ---
name : researcher
model : claude-haiku-5-5
mcp_servers :
- type : url
name : github
url : https://api.githubcopilot.com/mcp/
tools :
- type : mcp_toolset
mcp_server_name : github
---
Then create the session with the vault that holds the GitHub credential:
session = client.beta.sessions.create(
agent = lead_agent.id,
environment_id = environment.id,
vault_ids = [vault.id],
)
print (session.id)
In this example, onl

[truncated]
