---
source: "https://m3.sineframe.com/blog/claude-code-vs-codex-mcp"
hn_url: "https://news.ycombinator.com/item?id=50033782"
title: "Claude Code and Codex break on different MCP features"
article_title: "Claude Code vs Codex: 10 MCP features tested | SineFrame M3"
image: "https://m3.sineframe.com/blog/claude-code-vs-codex-mcp.png"
author: "stag"
captured_at: "2026-10-10T16:10:33Z"
capture_tool: "hn-digest"
hn_id: 50033782
score: 2
comments: 0
posted_at: "2026-10-10T15:16:02Z"
tags:
  - hacker-news
---

# Claude Code and Codex break on different MCP features

- HN: [50033782](https://news.ycombinator.com/item?id=50033782)
- Source: [m3.sineframe.com](https://m3.sineframe.com/blog/claude-code-vs-codex-mcp)
- Score: 2
- Comments: 0
- Posted: 2026-10-10T15:16:02Z

## Translation

Title: Claude Code and Codex break on different MCP features
Article title: Claude Code vs Codex: 10 MCP features tested | SineFrame M3
Description: We ran Claude Code and Codex against the same ten MCP servers at pinned versions. Five features split them, and they never failed the same one.

Article text:
Claude Code vs Codex: 10 MCP features tested | SineFrame M3
SineFrame M3 CI gate for MCPs Back to home Blog Claude Code and Codex break on different MCP features
We ran Claude Code and Codex against the same ten MCP servers at pinned versions. Five features split them, and they never failed the same one.
We ran Claude Code and Codex against the same ten MCP servers. Five of the ten features worked on one agent and failed on the other, and the two agents never failed on the same feature.
If you ship an MCP server, you have probably seen the issue: "tool doesn't show up in Claude Code" or "Codex sends the wrong arguments." When we started, about 28 issues like these were open across agent repositories. None of them says which agent, at which version, breaks on which MCP feature.
So we built a table that does. We call it the Harness Quirks Matrix, and it runs on M3 , our open-source tool for testing MCP servers and the agents that use them.
The results: 5 of 10 MCP features split the agents
On 9 October 2026 we ran ten MCP features against Claude Code 2.1.287 ( claude-sonnet-5-5 ) and Codex 0.162.0 ( gpt-5.6-sol ), three trials each. Claude Code failed three of the features and Codex failed the other two.
Here is what the five failures mean if you maintain an MCP server.
Claude Code never calls a tool whose input property is named ids[] . It shows no error, and the tool simply isn't there. OpenAPI-generated servers often produce names like this.
Claude Code also keeps only about the first and last 5,000 characters of a tool error, so anything in the middle never reaches the model. Put the code, the reason and the next step at the start of an error message. And when a result carries both content and structuredContent , the model reported the structured value and lost the text one.
Codex kept using its old catalog after the server announced a new tool, so it never called the new one. It also compacts large input schemas to fit a budget, and in our 14 KB schema the required nested fields went with them. Raising Codex's per-server schema budget to 20,000 restored the call.
Each result covers one client and model pair at pinned versions, with explicit prompts. A failed cell tells you the data didn't get from the server to the model's answer; it doesn't tell you which layer dropped it. Every cell has its run manifest, binary SHA-256 and wire trace stored with it.
Eight more failures from direct probes
We also probed the Claude Code and Codex binaries directly, without M3, and found eight more features that fail on at least one of them. None has run through M3 yet, so treat them as leads until they do.
The probes also cleared some old suspects. Claude Code's root anyOf and 65-character name issues no longer reproduce at current versions, and neither does Codex's resource_link handling, so server authors can drop the workarounds for those.
Every cell is an ordinary pytest test that M3 runs against the real agent binary, at an exact version, while recording what crosses the wire. Here is the whole test for the ids[] case:
@pytest.mark.m3(suite_name="quirks")
def test_q05_args_property_name_brackets(agent) -> None:
state = run_eval(agent, Path(__file__).resolve().parent)
assert state == "works", state
run_eval calls agent.run(prompt, servers=(control, quirk)) and then checks the result with M3's expect(result).to_have_tool_call(...) against wire evidence. The agent fixture comes from the command line:
m3 test --runtime=managed \
--harness [email protected] =claude-sonnet-5-5 \
--harness [email protected] =gpt-5.6-sol \
--trials 3 -- quirks/
A few things keep the cells honest:
M3 downloads each agent release into a cache and records its version and SHA-256, so a new release gets its own column.
M3 records separately whether the agent reported a tool call and whether the server received it. Most of the failures above happen between those two points.
Each run puts a plain echo server next to the server under test. If echo fails, the cell is marked not measured, so a broken setup can't show up as an agent bug.
The server generates a random marker for every trial, so the model can't get the answer from the prompt.
We re-run every failure against the same binary without M3, to check that M3 itself isn't the cause.
The table is rendered only from stored files, and re-rendering it gives the same bytes.
The full run took 66 executions and about 16 minutes. Claude Code averaged 9.8 seconds per execution; Codex averaged 18.2 seconds.
All twenty tests are in the M3 repository under benchmarks/harness-quirks , with the command to run them yourself.
Our first version found nothing. We started with five features taken from public bug reports, and all 36 executions passed on both agents. Those bugs had either been fixed or only showed up on setups we weren't testing, which is why we went back and probed the binaries directly.
One sentence in a prompt broke every Codex run. Our first prompts ended with "If a tool is not available, say which one and stop." Codex took that as a reason to give up, and every cell came back control failed. Without the control server, we would have published that as a Codex bug.
With gpt-5.6-sol , Codex runs in code mode, so the model never sees your JSON Schema. It gets a compacted TypeScript declaration generated from it, and several of the Codex failures in this post happen in that translation.
Test your own MCP server on every agent
The matrix is just M3 tests, and you can write the same test for your own server in about ten lines. Install the CLI and scaffold a project:
uv tool install sf-m3-cli
m3 init # a skipped starter test and an .env.example
m3 setup
Write the behaviour you care about as a normal pytest test:
import pytest
from m3 import expect
@pytest.mark.m3
def test_shipping(agent, shipping_server):
result = agent.run("Get a local shipping quote", server=shipping_server)
expect(result).to_have_tool_call("shipping_quote", server=shipping_server.name,
status="success")
Then run it against as many agents and versions as you like:
m3 test --runtime=managed \
--harness [email protected] =claude-sonnet-5-5 \
--harness [email protected] =gpt-5.6-sol \
--trials 3 -- tests/test_shipping.py
M3 ships with Claude Code, Codex, OpenCode and Pi built in, and any agent that speaks ACP can be added with a small manifest. Runs are saved locally and can be compared with a saved baseline, so when a new agent release breaks your server, you see a failing test before your users see the bug. To make that a pull-request check, see gating an MCP server in CI .
M3 is open source under Apache 2.0: github.com/sineframe/m3 . Start with the documentation . If your server has a quirk we should add to the matrix, open an issue there.
Written by Rishav Katoch Co-founder, SineFrame

## Original Extract

We ran Claude Code and Codex against the same ten MCP servers at pinned versions. Five features split them, and they never failed the same one.

Claude Code vs Codex: 10 MCP features tested | SineFrame M3
SineFrame M3 CI gate for MCPs Back to home Blog Claude Code and Codex break on different MCP features
We ran Claude Code and Codex against the same ten MCP servers at pinned versions. Five features split them, and they never failed the same one.
We ran Claude Code and Codex against the same ten MCP servers. Five of the ten features worked on one agent and failed on the other, and the two agents never failed on the same feature.
If you ship an MCP server, you have probably seen the issue: "tool doesn't show up in Claude Code" or "Codex sends the wrong arguments." When we started, about 28 issues like these were open across agent repositories. None of them says which agent, at which version, breaks on which MCP feature.
So we built a table that does. We call it the Harness Quirks Matrix, and it runs on M3 , our open-source tool for testing MCP servers and the agents that use them.
The results: 5 of 10 MCP features split the agents
On 9 October 2026 we ran ten MCP features against Claude Code 2.1.287 ( claude-sonnet-5-5 ) and Codex 0.162.0 ( gpt-5.6-sol ), three trials each. Claude Code failed three of the features and Codex failed the other two.
Here is what the five failures mean if you maintain an MCP server.
Claude Code never calls a tool whose input property is named ids[] . It shows no error, and the tool simply isn't there. OpenAPI-generated servers often produce names like this.
Claude Code also keeps only about the first and last 5,000 characters of a tool error, so anything in the middle never reaches the model. Put the code, the reason and the next step at the start of an error message. And when a result carries both content and structuredContent , the model reported the structured value and lost the text one.
Codex kept using its old catalog after the server announced a new tool, so it never called the new one. It also compacts large input schemas to fit a budget, and in our 14 KB schema the required nested fields went with them. Raising Codex's per-server schema budget to 20,000 restored the call.
Each result covers one client and model pair at pinned versions, with explicit prompts. A failed cell tells you the data didn't get from the server to the model's answer; it doesn't tell you which layer dropped it. Every cell has its run manifest, binary SHA-256 and wire trace stored with it.
Eight more failures from direct probes
We also probed the Claude Code and Codex binaries directly, without M3, and found eight more features that fail on at least one of them. None has run through M3 yet, so treat them as leads until they do.
The probes also cleared some old suspects. Claude Code's root anyOf and 65-character name issues no longer reproduce at current versions, and neither does Codex's resource_link handling, so server authors can drop the workarounds for those.
Every cell is an ordinary pytest test that M3 runs against the real agent binary, at an exact version, while recording what crosses the wire. Here is the whole test for the ids[] case:
@pytest.mark.m3(suite_name="quirks")
def test_q05_args_property_name_brackets(agent) -> None:
state = run_eval(agent, Path(__file__).resolve().parent)
assert state == "works", state
run_eval calls agent.run(prompt, servers=(control, quirk)) and then checks the result with M3's expect(result).to_have_tool_call(...) against wire evidence. The agent fixture comes from the command line:
m3 test --runtime=managed \
--harness [email protected] =claude-sonnet-5-5 \
--harness [email protected] =gpt-5.6-sol \
--trials 3 -- quirks/
A few things keep the cells honest:
M3 downloads each agent release into a cache and records its version and SHA-256, so a new release gets its own column.
M3 records separately whether the agent reported a tool call and whether the server received it. Most of the failures above happen between those two points.
Each run puts a plain echo server next to the server under test. If echo fails, the cell is marked not measured, so a broken setup can't show up as an agent bug.
The server generates a random marker for every trial, so the model can't get the answer from the prompt.
We re-run every failure against the same binary without M3, to check that M3 itself isn't the cause.
The table is rendered only from stored files, and re-rendering it gives the same bytes.
The full run took 66 executions and about 16 minutes. Claude Code averaged 9.8 seconds per execution; Codex averaged 18.2 seconds.
All twenty tests are in the M3 repository under benchmarks/harness-quirks , with the command to run them yourself.
Our first version found nothing. We started with five features taken from public bug reports, and all 36 executions passed on both agents. Those bugs had either been fixed or only showed up on setups we weren't testing, which is why we went back and probed the binaries directly.
One sentence in a prompt broke every Codex run. Our first prompts ended with "If a tool is not available, say which one and stop." Codex took that as a reason to give up, and every cell came back control failed. Without the control server, we would have published that as a Codex bug.
With gpt-5.6-sol , Codex runs in code mode, so the model never sees your JSON Schema. It gets a compacted TypeScript declaration generated from it, and several of the Codex failures in this post happen in that translation.
Test your own MCP server on every agent
The matrix is just M3 tests, and you can write the same test for your own server in about ten lines. Install the CLI and scaffold a project:
uv tool install sf-m3-cli
m3 init # a skipped starter test and an .env.example
m3 setup
Write the behaviour you care about as a normal pytest test:
import pytest
from m3 import expect
@pytest.mark.m3
def test_shipping(agent, shipping_server):
result = agent.run("Get a local shipping quote", server=shipping_server)
expect(result).to_have_tool_call("shipping_quote", server=shipping_server.name,
status="success")
Then run it against as many agents and versions as you like:
m3 test --runtime=managed \
--harness [email protected] =claude-sonnet-5-5 \
--harness [email protected] =gpt-5.6-sol \
--trials 3 -- tests/test_shipping.py
M3 ships with Claude Code, Codex, OpenCode and Pi built in, and any agent that speaks ACP can be added with a small manifest. Runs are saved locally and can be compared with a saved baseline, so when a new agent release breaks your server, you see a failing test before your users see the bug. To make that a pull-request check, see gating an MCP server in CI .
M3 is open source under Apache 2.0: github.com/sineframe/m3 . Start with the documentation . If your server has a quirk we should add to the matrix, open an issue there.
Written by Rishav Katoch Co-founder, SineFrame
