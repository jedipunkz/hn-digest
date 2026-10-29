---
source: "https://santurcesoftware.com/mcp-tools-should-disclose-llm-and-api-calls"
hn_url: "https://news.ycombinator.com/item?id=50290905"
title: "MCP tools should disclose LLM and external API dependencies"
article_title: "Your Test Harness Should Know When an MCP Tool Calls an LLM | Santurce Software | Santurce Software"
image: "https://santurcesoftware.com/static/img/social-card.png"
author: "boredumb"
captured_at: "2026-10-29T16:40:10Z"
capture_tool: "hn-digest"
hn_id: 50290905
score: 1
comments: 0
posted_at: "2026-10-29T15:44:34Z"
tags:
  - hacker-news
---

# MCP tools should disclose LLM and external API dependencies

- HN: [50290905](https://news.ycombinator.com/item?id=50290905)
- Source: [santurcesoftware.com](https://santurcesoftware.com/mcp-tools-should-disclose-llm-and-api-calls)
- Score: 1
- Comments: 0
- Posted: 2026-10-29T15:44:34Z

## Translation

Title: MCP tools should disclose LLM and external API dependencies
Article title: Your Test Harness Should Know When an MCP Tool Calls an LLM | Santurce Software | Santurce Software
Description: MCP tools are dependencies of your tests. Disclosing LLM inference and external API calls helps you choose assertions, fixtures and verification checks.

Article text:
Skip to content
Menu
Services Industries Selected work About Insights Discuss your project ↗
Home
Your Test Harness Should Know When an MCP Tool Calls an LLM
Engineering notes · October 2026
Your test harness should know when an MCP tool calls an LLM
If I am writing tests around your MCP tool, I want to know whether it reads a record, calls another service or asks another model. that all changes the test I should write.
We have been building harnesses for automations and getting the agent to do the work is one part of it. Writing code that checks whether the work actually happened is another.
Some of that is straightforward. Parse the response. Check the ID. Read the record back and compare the fields. But once your verification code starts calling MCP tools, you need to understand what those tools are doing too.
Is the tool reading application state? Calling a third-party API? Sending the whole thing to another LLM and asking whether it looks right? Those need different tests, different fixtures and different expectations when something fails.
I would like MCP to give us a standard way to declare that. It is fairly important information when the tool is a dependency of your tests.
Take this hypothetical example:
result = mcp.call("verify_project_setup", {
"project_id": project_id,
"requirements": requirements,
})
assert result["passed"] is True
The Python is easy enough. The interesting part is whatever is behind verify_project_setup .
That could be a database read followed by a few comparisons. It could fetch the project from an external service. It could ask a model to review a description of the project. Or it could fetch the actual records and have a model evaluate those.
Every one of those can return {"passed": true} . We now have four rather different tests behind the same line of code.
For the database comparison, I want known records and clear comparison rules. For the API call, I need to account for credentials, remote state and service failures. For the model review, I need the grading criteria, model configuration and some understanding of how much the results vary. If it does both, I need to account for both.
Valid JSON tells me the response parsed. It tells me very little about what made the test pass.
This affects the tests you write
If there is an LLM generating prose under the hood, comparing the entire response with a fixed string may not be particularly useful. I might check its structure, check specific facts, run it against examples with known expected outcomes or measure the results over repeated runs. Exact comparisons can still make sense for constrained outputs. I need to know what I am dealing with before choosing.
If it calls an external API, I probably want a local fixture for testing how my code handles the response, and a separate integration test for the actual service. Passing against my fixture does not tell me the live integration works. A passing live call does not exercise every failure case my fixture can cover.
And if the tool does both, mocking the business API might still leave a model running inside the test. Mocking the entire MCP response with {"passed": true} can remove the verification logic I meant to test in the first place. These details affect where I put the mocks, what I assert, how long I wait and which tests belong in which CI job.
There are perfectly good reasons to use an LLM to review an automation. Some requirements are awkward to express as a handful of field comparisons. I may want it to assess whether a response addresses the request or whether a document includes enough explanation.
But now that model is part of my test system. What did we give it to inspect? What are the grading rules? Is it looking at the resulting records, or just reading the agent saying it finished? If I change the model, does the pass rate change?
An agent saying “I created the record”, followed by another model saying “yes, that sounds complete”, has not established that the record exists. Reading it back and checking the returned ID gives me an actual state check. A model can help review that evidence as well. I just need to know which part it is responsible for.
You dont want to discover that your clever state check is actually a model reviewing generated text after spending an afternoon chasing an intermittent test failure.
What about the existing MCP annotations?
MCP already gives us useful hints:
readOnlyHint describes whether a tool modifies its environment.
idempotentHint describes whether repeating identical arguments produces additional effects on its
environment.
openWorldHint describes interaction with an open world of external entities.
destructiveHint describes potentially destructive updates.
Those matter. They help with setup, cleanup and whether repeating an operation creates more side effects. But idempotent does not mean identical output, and read-only does not mean there is no model involved. openWorldHint can cover a business API, a hosted model or both.
What I am asking for is a little more detail alongside those hints: does this tool invoke an LLM, and which external APIs can it call?
I initially thought of this as deterministic versus LLM-backed, but that is too loose. A local model does not need an external API. A live API can return changing data without involving a model at all. No LLM does not automatically mean reproducible. These should be separate declarations.
What we added to our own server
We started with our own website. Its MCP server has two tools: get_company_profile reads data already in the Go process, and submit_project_inquiry validates an authorised introduction and sends it to Mailchimp when intake is enabled. Neither calls an LLM.
The inquiry tool now declares this in the accompanying source:
{
"name": "submit_project_inquiry",
"_meta": {
"com.santurcesoftware/execution": {
"schemaVersion": 1,
"llm": "never",
"externalApi": "conditional",
"externalApiProviders": ["mailchimp"]
}
}
}
The profile tool declares "llm": "never" , "externalApi": "never" and an empty provider list. We also state these dependencies in the descriptions.
This is our private extension, not a standard MCP annotation. It supplements the standard hints. An ordinary client is not expected to understand the key. The MCP maintainers' discussion of annotations recommends namespaced metadata as a place to try ideas before proposing a shared convention.
Here, “external API” means an application-level request to a separately operated downstream service. It excludes the incoming MCP connection and is not a complete inventory of databases, telemetry or network traffic. The declaration covers our implementation and known dependencies; it does not claim access to Mailchimp's internal implementation.
What could happen, and what did happen?
Knowing that a tool can call Mailchimp is useful. Knowing whether a particular test reached that call is useful too.
The request might stop at validation. Intake might be disabled. A local rate limit might reject it. An error response alone does not tell me whether we attempted the external call.
So our tool results also include:
{
"_meta": {
"com.santurcesoftware/execution": {
"schemaVersion": 1,
"llmInvoked": false,
"externalApiAttempted": true
}
}
}
We set the API observation immediately before the HTTP client attempts the request. It remains true for transport errors and provider failures. It does not prove remote receipt or completion. Invalid arguments returned as JSON-RPC errors dont carry these tool-result observations.
Our tests use a local mock provider so we can compare the reported path with the request the mock received. Disabled intake and local rate limits should show no attempt. A provider returning HTTP 200 with the wrong acknowledgement should still fail. That is the sort of detail I want the harness checking.
The declaration tells me how to set up the test. The observation helps explain which branch ran. The mock and the assertions check the behaviour. I still need all of those pieces.
A harness could check these fields when it discovers the tools. If a test suite assumes there is no model inference, a tool declaring otherwise should make that assumption fail clearly. A missing or unfamiliar declaration should mean we dont know, rather than quietly being treated as no LLM.
For a client that understands our private fields, that could look like this:
# Illustrative harness policy, not a standard MCP client API.
execution = tool.get("_meta", {}).get(
"com.santurcesoftware/execution", {}
)
if execution.get("schemaVersion") != 1:
raise TestConfigurationError("Unknown execution declaration")
if execution.get("llm") != "never":
raise TestConfigurationError("Suite excludes LLM inference")
if execution.get("externalApi") != "never":
raise TestConfigurationError("Suite excludes downstream APIs")
# Run behavioural assertions in a controlled environment.
Another suite could allow Mailchimp, point our server at a mock endpoint and inspect the requests. A live integration suite could use the actual provider. A model evaluation suite could allow inference and save the model configuration and grading rules with the results.
The point is that the harness can make that choice deliberately.
A few details we need to get right
A useful shared vocabulary could start with never , conditional , always and unknown , with precise definitions for preflight failures and nested dependencies. Our current server only emits the first two.
MCP's client-mediated sampling makes some inference visible to the client. Direct server-side provider calls still need disclosure. Cache hits introduce another distinction: no model invoked during this call does not mean the returned content was never model-generated.
These fields will not explain an entire verification algorithm. I may still need the source, comparison rules, fixtures, model settings and traces. But being able to discover the major dependencies is a useful starting point.
A server can also report this incorrectly. In a test deployment I control, I can check the declaration against instrumented clients and network restrictions. For a remote service, it remains a claim from that service.
If I import a library into my test runner, I expect to understand the dependencies it brings with it. Calling that code through MCP does not make those dependencies any less relevant.
We already have useful annotations for effects. I would like us to add a shared way to describe LLM calls, external APIs and which paths were attempted. It would make writing the harness and understanding what its tests actually prove a good deal easier.
Related: Building Agentic Lab Workflows with Janus, Hermes and Ollama , Observing AI Workflows with Rust and Langfuse , and our MCP agent guide .
Tool Annotations as Risk
Vocabulary
Custom software. Sound engineering.
Built in Puerto Rico. Working remotely.

## Original Extract

MCP tools are dependencies of your tests. Disclosing LLM inference and external API calls helps you choose assertions, fixtures and verification checks.

Skip to content
Menu
Services Industries Selected work About Insights Discuss your project ↗
Home
Your Test Harness Should Know When an MCP Tool Calls an LLM
Engineering notes · October 2026
Your test harness should know when an MCP tool calls an LLM
If I am writing tests around your MCP tool, I want to know whether it reads a record, calls another service or asks another model. that all changes the test I should write.
We have been building harnesses for automations and getting the agent to do the work is one part of it. Writing code that checks whether the work actually happened is another.
Some of that is straightforward. Parse the response. Check the ID. Read the record back and compare the fields. But once your verification code starts calling MCP tools, you need to understand what those tools are doing too.
Is the tool reading application state? Calling a third-party API? Sending the whole thing to another LLM and asking whether it looks right? Those need different tests, different fixtures and different expectations when something fails.
I would like MCP to give us a standard way to declare that. It is fairly important information when the tool is a dependency of your tests.
Take this hypothetical example:
result = mcp.call("verify_project_setup", {
"project_id": project_id,
"requirements": requirements,
})
assert result["passed"] is True
The Python is easy enough. The interesting part is whatever is behind verify_project_setup .
That could be a database read followed by a few comparisons. It could fetch the project from an external service. It could ask a model to review a description of the project. Or it could fetch the actual records and have a model evaluate those.
Every one of those can return {"passed": true} . We now have four rather different tests behind the same line of code.
For the database comparison, I want known records and clear comparison rules. For the API call, I need to account for credentials, remote state and service failures. For the model review, I need the grading criteria, model configuration and some understanding of how much the results vary. If it does both, I need to account for both.
Valid JSON tells me the response parsed. It tells me very little about what made the test pass.
This affects the tests you write
If there is an LLM generating prose under the hood, comparing the entire response with a fixed string may not be particularly useful. I might check its structure, check specific facts, run it against examples with known expected outcomes or measure the results over repeated runs. Exact comparisons can still make sense for constrained outputs. I need to know what I am dealing with before choosing.
If it calls an external API, I probably want a local fixture for testing how my code handles the response, and a separate integration test for the actual service. Passing against my fixture does not tell me the live integration works. A passing live call does not exercise every failure case my fixture can cover.
And if the tool does both, mocking the business API might still leave a model running inside the test. Mocking the entire MCP response with {"passed": true} can remove the verification logic I meant to test in the first place. These details affect where I put the mocks, what I assert, how long I wait and which tests belong in which CI job.
There are perfectly good reasons to use an LLM to review an automation. Some requirements are awkward to express as a handful of field comparisons. I may want it to assess whether a response addresses the request or whether a document includes enough explanation.
But now that model is part of my test system. What did we give it to inspect? What are the grading rules? Is it looking at the resulting records, or just reading the agent saying it finished? If I change the model, does the pass rate change?
An agent saying “I created the record”, followed by another model saying “yes, that sounds complete”, has not established that the record exists. Reading it back and checking the returned ID gives me an actual state check. A model can help review that evidence as well. I just need to know which part it is responsible for.
You dont want to discover that your clever state check is actually a model reviewing generated text after spending an afternoon chasing an intermittent test failure.
What about the existing MCP annotations?
MCP already gives us useful hints:
readOnlyHint describes whether a tool modifies its environment.
idempotentHint describes whether repeating identical arguments produces additional effects on its
environment.
openWorldHint describes interaction with an open world of external entities.
destructiveHint describes potentially destructive updates.
Those matter. They help with setup, cleanup and whether repeating an operation creates more side effects. But idempotent does not mean identical output, and read-only does not mean there is no model involved. openWorldHint can cover a business API, a hosted model or both.
What I am asking for is a little more detail alongside those hints: does this tool invoke an LLM, and which external APIs can it call?
I initially thought of this as deterministic versus LLM-backed, but that is too loose. A local model does not need an external API. A live API can return changing data without involving a model at all. No LLM does not automatically mean reproducible. These should be separate declarations.
What we added to our own server
We started with our own website. Its MCP server has two tools: get_company_profile reads data already in the Go process, and submit_project_inquiry validates an authorised introduction and sends it to Mailchimp when intake is enabled. Neither calls an LLM.
The inquiry tool now declares this in the accompanying source:
{
"name": "submit_project_inquiry",
"_meta": {
"com.santurcesoftware/execution": {
"schemaVersion": 1,
"llm": "never",
"externalApi": "conditional",
"externalApiProviders": ["mailchimp"]
}
}
}
The profile tool declares "llm": "never" , "externalApi": "never" and an empty provider list. We also state these dependencies in the descriptions.
This is our private extension, not a standard MCP annotation. It supplements the standard hints. An ordinary client is not expected to understand the key. The MCP maintainers' discussion of annotations recommends namespaced metadata as a place to try ideas before proposing a shared convention.
Here, “external API” means an application-level request to a separately operated downstream service. It excludes the incoming MCP connection and is not a complete inventory of databases, telemetry or network traffic. The declaration covers our implementation and known dependencies; it does not claim access to Mailchimp's internal implementation.
What could happen, and what did happen?
Knowing that a tool can call Mailchimp is useful. Knowing whether a particular test reached that call is useful too.
The request might stop at validation. Intake might be disabled. A local rate limit might reject it. An error response alone does not tell me whether we attempted the external call.
So our tool results also include:
{
"_meta": {
"com.santurcesoftware/execution": {
"schemaVersion": 1,
"llmInvoked": false,
"externalApiAttempted": true
}
}
}
We set the API observation immediately before the HTTP client attempts the request. It remains true for transport errors and provider failures. It does not prove remote receipt or completion. Invalid arguments returned as JSON-RPC errors dont carry these tool-result observations.
Our tests use a local mock provider so we can compare the reported path with the request the mock received. Disabled intake and local rate limits should show no attempt. A provider returning HTTP 200 with the wrong acknowledgement should still fail. That is the sort of detail I want the harness checking.
The declaration tells me how to set up the test. The observation helps explain which branch ran. The mock and the assertions check the behaviour. I still need all of those pieces.
A harness could check these fields when it discovers the tools. If a test suite assumes there is no model inference, a tool declaring otherwise should make that assumption fail clearly. A missing or unfamiliar declaration should mean we dont know, rather than quietly being treated as no LLM.
For a client that understands our private fields, that could look like this:
# Illustrative harness policy, not a standard MCP client API.
execution = tool.get("_meta", {}).get(
"com.santurcesoftware/execution", {}
)
if execution.get("schemaVersion") != 1:
raise TestConfigurationError("Unknown execution declaration")
if execution.get("llm") != "never":
raise TestConfigurationError("Suite excludes LLM inference")
if execution.get("externalApi") != "never":
raise TestConfigurationError("Suite excludes downstream APIs")
# Run behavioural assertions in a controlled environment.
Another suite could allow Mailchimp, point our server at a mock endpoint and inspect the requests. A live integration suite could use the actual provider. A model evaluation suite could allow inference and save the model configuration and grading rules with the results.
The point is that the harness can make that choice deliberately.
A few details we need to get right
A useful shared vocabulary could start with never , conditional , always and unknown , with precise definitions for preflight failures and nested dependencies. Our current server only emits the first two.
MCP's client-mediated sampling makes some inference visible to the client. Direct server-side provider calls still need disclosure. Cache hits introduce another distinction: no model invoked during this call does not mean the returned content was never model-generated.
These fields will not explain an entire verification algorithm. I may still need the source, comparison rules, fixtures, model settings and traces. But being able to discover the major dependencies is a useful starting point.
A server can also report this incorrectly. In a test deployment I control, I can check the declaration against instrumented clients and network restrictions. For a remote service, it remains a claim from that service.
If I import a library into my test runner, I expect to understand the dependencies it brings with it. Calling that code through MCP does not make those dependencies any less relevant.
We already have useful annotations for effects. I would like us to add a shared way to describe LLM calls, external APIs and which paths were attempted. It would make writing the harness and understanding what its tests actually prove a good deal easier.
Related: Building Agentic Lab Workflows with Janus, Hermes and Ollama , Observing AI Workflows with Rust and Langfuse , and our MCP agent guide .
Tool Annotations as Risk
Vocabulary
Custom software. Sound engineering.
Built in Puerto Rico. Working remotely.
