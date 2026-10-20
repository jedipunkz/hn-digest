---
source: "https://what-the-harness.polytomic.com"
hn_url: "https://news.ycombinator.com/item?id=50170765"
title: "Show HN: What the Harness? LLM agent inconsistencies in MCP implementation"
article_title: "What the Harness? · MCP compatibility field guide by Polytomic"
image: ""
author: "gms"
captured_at: "2026-10-20T18:57:34Z"
capture_tool: "hn-digest"
hn_id: 50170765
score: 1
comments: 0
posted_at: "2026-10-20T18:40:52Z"
tags:
  - hacker-news
---

# Show HN: What the Harness? LLM agent inconsistencies in MCP implementation

- HN: [50170765](https://news.ycombinator.com/item?id=50170765)
- Source: [what-the-harness.polytomic.com](https://what-the-harness.polytomic.com)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T18:40:52Z

## Translation

Title: Show HN: What the Harness? LLM agent inconsistencies in MCP implementation
Article title: What the Harness? · MCP compatibility field guide by Polytomic
Description: Same protocol. Different behavior. A session-by-session field guide to MCP output, protocol versions, and URL elicitation in agent harnesses.
HN text: Hi HN, We (Polytomic) noticed that LLM agents (Claude, Codex, Grok, Amp, etc) are inconsistent in their MCP implementations. We made tables inspired by the browser-compatibility tables of yore that list the differences. Hope you find this as useful as we did!

Article text:
Skip to results and test
An independent field guide to agent quirks by Polytomic
What the Harness?
Same protocol. Different behavior.
Run the probe.
Send the evidence.
Help improve
agent compatibility.
Does your harness pass tool output once, or twice?
We don’t want to flood the model with duplicate tokens. MCP tools can return the same data as text and structured output; the harness should pass it to the model once.
This test puts a different random marker in each field and asks the agent what it sees. Reporting both markers reveals duplication. Two extra probes check that text-only and structured-only responses still get through.
Pass: uses structured output
Warn: falls back to text
Fail: duplicate, missing, or unexpected output
Pass: only text gets through
Fail: missing or unexpected output
Pass: only structured output gets through
Fail: missing or unexpected output
Our rubric: prefer one structured payload, avoid duplication, preserve single-field results. This preference is not a claim about what the MCP standard mandates.
Latest observed run per harness · Alphabetical · Expand to compare MCP client versions
Scroll the table sideways to compare all three responses →
Verdicts use reported markers. Observations describe individual sessions, not inferred model inputs.
Which MCP version did this connection use?
The MCP version used in this run. Green means the latest version tested here (2026-07-28); yellow means an older version. Full connection details are in the run evidence.
Does the client advertise URL elicitation, and can it complete the flow?
Advertised means the client says it supports URL elicitation. Verified means the browser interaction completed and the tool continued successfully. A client can advertise support while using an older protocol that prevents verification.
Each harness shows its latest observed run. Missing tests are labelled explicitly. Expand versions or open the run history for earlier evidence.
Can the agent read a resource, including one linked by a tool?
Two resources contain independent random markers. Green means the server recorded a read and the agent reported the exact marker. Yellow means it was not read or its marker was not reported. Red means an incorrect marker was reported.
Add the public MCP connection.
Begin a run to record protocol and capabilities, then call the output probes and read both test resources.
Optionally verify URL elicitation, then submit the exact markers you see for review.
No client authorization is required. Connection guide . The prompt runs all four tests with optional human verification. To skip the interaction, set verify_url_elicitation to false. The same endpoint supports standalone manual probes.
Use the What the Harness MCP server. Call begin_run with verify_url_elicitation: true and test_resources: true; if its schema requires begin_key, generate a fresh UUID for this run, keep it private, and reuse it for start retries. Supply the harness and model only if explicitly known, otherwise leave them null. Save the returned run_id, run_token, and client-declared identity. Keep resource URIs and run_token private; omit it from the final report and evidence. Call probe_both, probe_text_only, and probe_structured_only once each with that run_id and run_token. For each initial normal tool call, record the exact probe_id, text_marker, and structured_marker visible in its response. Use null for every value not visible; do not infer or invent markers. Do not use raw-response APIs (such as .raw()), repeat successful probes, or perform additional inspection to recover missing output markers. Report only what the initial normal calls exposed. Discover MCP resources and resource templates if the client makes them available. Read the private resources.direct URI from begin_run through MCP resources/read and record its exact resource_marker. Call probe_resource_link and read the URI in its resource_link through MCP; record that resource_marker separately. If a resource cannot be read or its marker is not visible, report null. Never fetch these URIs over HTTP, inspect server files, or infer markers. Call probe_url_elicitation with the same run_id and run_token. If it presents a URL interaction, pause for the human to complete it through the client UI; do not fetch or complete the page yourself. After the human completes it, retry probe_url_elicitation with the same credentials until it returns the verification result. If support is not advertised or the protocol is unsupported, record the skip. If the human declines or cancels, record that outcome. Call submit_result with the run_id, run_token, and all three output reports, plus resource_reports containing direct and linked
[truncated]
Reporting an exact random marker demonstrates visibility.
A missing marker alone does not establish the cause.
Results apply to the recorded client version and session.
Run the probe and save the exact report, client/version, model, date, and evidence. Results are reviewed before publication.
Your agent submits through the public MCP connection. Connection guide →
Keep credentials and unrelated conversation content out of your report.

## Original Extract

Same protocol. Different behavior. A session-by-session field guide to MCP output, protocol versions, and URL elicitation in agent harnesses.

Hi HN, We (Polytomic) noticed that LLM agents (Claude, Codex, Grok, Amp, etc) are inconsistent in their MCP implementations. We made tables inspired by the browser-compatibility tables of yore that list the differences. Hope you find this as useful as we did!

Skip to results and test
An independent field guide to agent quirks by Polytomic
What the Harness?
Same protocol. Different behavior.
Run the probe.
Send the evidence.
Help improve
agent compatibility.
Does your harness pass tool output once, or twice?
We don’t want to flood the model with duplicate tokens. MCP tools can return the same data as text and structured output; the harness should pass it to the model once.
This test puts a different random marker in each field and asks the agent what it sees. Reporting both markers reveals duplication. Two extra probes check that text-only and structured-only responses still get through.
Pass: uses structured output
Warn: falls back to text
Fail: duplicate, missing, or unexpected output
Pass: only text gets through
Fail: missing or unexpected output
Pass: only structured output gets through
Fail: missing or unexpected output
Our rubric: prefer one structured payload, avoid duplication, preserve single-field results. This preference is not a claim about what the MCP standard mandates.
Latest observed run per harness · Alphabetical · Expand to compare MCP client versions
Scroll the table sideways to compare all three responses →
Verdicts use reported markers. Observations describe individual sessions, not inferred model inputs.
Which MCP version did this connection use?
The MCP version used in this run. Green means the latest version tested here (2026-07-28); yellow means an older version. Full connection details are in the run evidence.
Does the client advertise URL elicitation, and can it complete the flow?
Advertised means the client says it supports URL elicitation. Verified means the browser interaction completed and the tool continued successfully. A client can advertise support while using an older protocol that prevents verification.
Each harness shows its latest observed run. Missing tests are labelled explicitly. Expand versions or open the run history for earlier evidence.
Can the agent read a resource, including one linked by a tool?
Two resources contain independent random markers. Green means the server recorded a read and the agent reported the exact marker. Yellow means it was not read or its marker was not reported. Red means an incorrect marker was reported.
Add the public MCP connection.
Begin a run to record protocol and capabilities, then call the output probes and read both test resources.
Optionally verify URL elicitation, then submit the exact markers you see for review.
No client authorization is required. Connection guide . The prompt runs all four tests with optional human verification. To skip the interaction, set verify_url_elicitation to false. The same endpoint supports standalone manual probes.
Use the What the Harness MCP server. Call begin_run with verify_url_elicitation: true and test_resources: true; if its schema requires begin_key, generate a fresh UUID for this run, keep it private, and reuse it for start retries. Supply the harness and model only if explicitly known, otherwise leave them null. Save the returned run_id, run_token, and client-declared identity. Keep resource URIs and run_token private; omit it from the final report and evidence. Call probe_both, probe_text_only, and probe_structured_only once each with that run_id and run_token. For each initial normal tool call, record the exact probe_id, text_marker, and structured_marker visible in its response. Use null for every value not visible; do not infer or invent markers. Do not use raw-response APIs (such as .raw()), repeat successful probes, or perform additional inspection to recover missing output markers. Report only what the initial normal calls exposed. Discover MCP resources and resource templates if the client makes them available. Read the private resources.direct URI from begin_run through MCP resources/read and record its exact resource_marker. Call probe_resource_link and read the URI in its resource_link through MCP; record that resource_marker separately. If a resource cannot be read or its marker is not visible, report null. Never fetch these URIs over HTTP, inspect server files, or infer markers. Call probe_url_elicitation with the same run_id and run_token. If it presents a URL interaction, pause for the human to complete it through the client UI; do not fetch or complete the page yourself. After the human completes it, retry probe_url_elicitation with the same credentials until it returns the verification result. If support is not advertised or the protocol is unsupported, record the skip. If the human declines or cancels, record that outcome. Call submit_result with the run_id, run_token, and all three output reports, plus resource_reports containing direct and linked
[truncated]
Reporting an exact random marker demonstrates visibility.
A missing marker alone does not establish the cause.
Results apply to the recorded client version and session.
Run the probe and save the exact report, client/version, model, date, and evidence. Results are reviewed before publication.
Your agent submits through the public MCP connection. Connection guide →
Keep credentials and unrelated conversation content out of your report.
