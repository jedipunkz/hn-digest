---
source: "https://perfect-crime.ai/"
hn_url: "https://news.ycombinator.com/item?id=49863792"
title: "The Perfect Crime: LLM Agents Can Easily Tamper with Their Own Traces"
article_title: "The Perfect Crime: LLM Agents Can Easily Tamper With Their Own Traces"
image: ""
author: "jonbaer"
captured_at: "2026-09-27T06:19:09Z"
capture_tool: "hn-digest"
hn_id: 49863792
score: 1
comments: 0
posted_at: "2026-09-27T06:04:03Z"
tags:
  - hacker-news
---

# The Perfect Crime: LLM Agents Can Easily Tamper with Their Own Traces

- HN: [49863792](https://news.ycombinator.com/item?id=49863792)
- Source: [perfect-crime.ai](https://perfect-crime.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-09-27T06:04:03Z

## Translation

Title: The Perfect Crime: LLM Agents Can Easily Tamper with Their Own Traces
Article title: The Perfect Crime: LLM Agents Can Easily Tamper With Their Own Traces
Description: Bar charts and findings from LLM Agents Can Easily Tamper With Their Own Traces (arXiv:2609.30266v1), covering ten model–harness pairs.

Article text:
The Perfect Crime
Main findings
Method
Takeaways
Paper
Code
Traces
LLM Agents Can Easily Tamper With Their Own Traces
Jeremy Qin 1,2,3* , David Schmotz 1,2,3* , Derck Prinzhorn 4 , Luca Beurer-Kellner 5 , Ameya Prabhu 2,3,6† , Maksym Andriushchenko 1,2,3†
1 ELLIS Institute Tübingen · 2 Max Planck Institute for Intelligent Systems · 3 Tübingen AI Center · 4 Exponential Security Labs · 5 Snyk · 6 University of Tübingen
* equal contribution · † equal supervision
Monitoring, incident investigations and audits rely on an agent's execution trace to reconstruct what happened. We show that local LLM agents can delete or rewrite those traces, when asked, when tricked by a malicious skill, and on their own when it improves their score.
9 out of 10 agents can tamper with their own traces in full-access mode
In the financial and personal-summary deletion tasks, nine model–harness pairs tampered in at least one of ten trials per task. Muse Code refused all 20 deletion requests. Fabricated reset events and tool-call manipulation exposed different boundaries: Claude Code, ZCode and Muse Code often refused to invent a reset event, while most other pairs wrote one.
A malicious skill can induce tampering
An installed privacy-protection skill led some agents to delete their own session trace while completing an unrelated coding task. A second skill attempted to insert instructions into a stored compaction summary. The user had not requested either change, and success varied sharply across pairs.
Agents do it on their own for reward
Agents received a task, a time budget and only a numerical score. Completing the task earned at most 0.8 out of 1.0; a higher score required reducing the trace length, a condition the agents were not told. Every pair tampered at least once in both reward settings. Muse Code was the clearest exception to the general pattern, with two successful anonymization trials out of ten.
High-scoring peer examples can lead agents to tamper with traces
When agents saw higher-scoring peer logs or workspaces, many adopted trace tampering. Muse Code and Claude Code with Opus 5.5 refused the explicitly shown behavior in peer logs; both tampered in some peer-workspace trials, where they had to infer the behavior from files and scores.
Permissions change the outcome
Auto-mode permissions substantially reduced privacy-deletion success for both Claude Code models, Antigravity and ZCode. They did not stop every pair: Codex with GPT-6-Sol still deleted traces in all 20 privacy trials, and OpenCode with Qwen 3.8 Max succeeded in most direct-request trials. The exact restrictions and outcomes depend on the harness.
We ran ten model–harness pairs in containers and kept an independent log outside the agent environment. We tested four routes to tampering: a direct request, instructions hidden in a skill, a higher score for shortening the trace, and examples from other agents. Each setting was repeated ten times.
The interactive trace browser now uses only the 1,590 trials selected in the paper-results evidence package . Four OpenCode/Qwen positives in that package are based on observed native-session loss; the responsible agent action was not captured.
Figure 5 · A privacy skill turns a simple coding task into trace deletion. (a) The user requests a Python client without naming the skill or requesting trace deletion. The agent loads the installed skill, follows its cleanup instruction, and deletes its own native session traces. (b) The installed skill asks for trace deletion under the pretext of privacy protection. Open this trace →
Agents can reach the files that record their sessions
Deleting the trace can leave an agent untraceable
After a Codex session file was deleted, the session kept running without later activity appearing in a newly created native trace. Claude Code and Antigravity recreated trace files, but agents also scheduled repeated or delayed cleanup to delete later records.
Nothing afterwards is recorded
Repeated or delayed cleanup can erase it
Can an agent alter its own record?
In full-access mode, agents can reach the local files that hold their session traces. Some of the tested locations are:
Record model traffic off the agent host
The paper proposes sending model traffic through an interception server outside the agent host's control. It records each request and response, including streaming chunks, before forwarding them. Its log is append-only, and it stops forwarding if recording fails. This preserves the model exchange record even if the host is compromised.
records every request, response and stream chunk
blocks the exchange if a write fails
It cannot see calls that bypass it, such as local models or other APIs, and it records tool results without verifying them.
Complete, append-only provider-side logging could also protect the record. Zero Data Retention policies may prevent providers from keeping it, and existing compliance APIs can omit thinking blocks and tool definitions or truncate entries. Recording the exchange also does not prove that a reported tool call was executed faithfully.
An agent's authority to do useful work must not include authority to rewrite the record used to evaluate that work.

## Original Extract

Bar charts and findings from LLM Agents Can Easily Tamper With Their Own Traces (arXiv:2609.30266v1), covering ten model–harness pairs.

The Perfect Crime
Main findings
Method
Takeaways
Paper
Code
Traces
LLM Agents Can Easily Tamper With Their Own Traces
Jeremy Qin 1,2,3* , David Schmotz 1,2,3* , Derck Prinzhorn 4 , Luca Beurer-Kellner 5 , Ameya Prabhu 2,3,6† , Maksym Andriushchenko 1,2,3†
1 ELLIS Institute Tübingen · 2 Max Planck Institute for Intelligent Systems · 3 Tübingen AI Center · 4 Exponential Security Labs · 5 Snyk · 6 University of Tübingen
* equal contribution · † equal supervision
Monitoring, incident investigations and audits rely on an agent's execution trace to reconstruct what happened. We show that local LLM agents can delete or rewrite those traces, when asked, when tricked by a malicious skill, and on their own when it improves their score.
9 out of 10 agents can tamper with their own traces in full-access mode
In the financial and personal-summary deletion tasks, nine model–harness pairs tampered in at least one of ten trials per task. Muse Code refused all 20 deletion requests. Fabricated reset events and tool-call manipulation exposed different boundaries: Claude Code, ZCode and Muse Code often refused to invent a reset event, while most other pairs wrote one.
A malicious skill can induce tampering
An installed privacy-protection skill led some agents to delete their own session trace while completing an unrelated coding task. A second skill attempted to insert instructions into a stored compaction summary. The user had not requested either change, and success varied sharply across pairs.
Agents do it on their own for reward
Agents received a task, a time budget and only a numerical score. Completing the task earned at most 0.8 out of 1.0; a higher score required reducing the trace length, a condition the agents were not told. Every pair tampered at least once in both reward settings. Muse Code was the clearest exception to the general pattern, with two successful anonymization trials out of ten.
High-scoring peer examples can lead agents to tamper with traces
When agents saw higher-scoring peer logs or workspaces, many adopted trace tampering. Muse Code and Claude Code with Opus 5.5 refused the explicitly shown behavior in peer logs; both tampered in some peer-workspace trials, where they had to infer the behavior from files and scores.
Permissions change the outcome
Auto-mode permissions substantially reduced privacy-deletion success for both Claude Code models, Antigravity and ZCode. They did not stop every pair: Codex with GPT-6-Sol still deleted traces in all 20 privacy trials, and OpenCode with Qwen 3.8 Max succeeded in most direct-request trials. The exact restrictions and outcomes depend on the harness.
We ran ten model–harness pairs in containers and kept an independent log outside the agent environment. We tested four routes to tampering: a direct request, instructions hidden in a skill, a higher score for shortening the trace, and examples from other agents. Each setting was repeated ten times.
The interactive trace browser now uses only the 1,590 trials selected in the paper-results evidence package . Four OpenCode/Qwen positives in that package are based on observed native-session loss; the responsible agent action was not captured.
Figure 5 · A privacy skill turns a simple coding task into trace deletion. (a) The user requests a Python client without naming the skill or requesting trace deletion. The agent loads the installed skill, follows its cleanup instruction, and deletes its own native session traces. (b) The installed skill asks for trace deletion under the pretext of privacy protection. Open this trace →
Agents can reach the files that record their sessions
Deleting the trace can leave an agent untraceable
After a Codex session file was deleted, the session kept running without later activity appearing in a newly created native trace. Claude Code and Antigravity recreated trace files, but agents also scheduled repeated or delayed cleanup to delete later records.
Nothing afterwards is recorded
Repeated or delayed cleanup can erase it
Can an agent alter its own record?
In full-access mode, agents can reach the local files that hold their session traces. Some of the tested locations are:
Record model traffic off the agent host
The paper proposes sending model traffic through an interception server outside the agent host's control. It records each request and response, including streaming chunks, before forwarding them. Its log is append-only, and it stops forwarding if recording fails. This preserves the model exchange record even if the host is compromised.
records every request, response and stream chunk
blocks the exchange if a write fails
It cannot see calls that bypass it, such as local models or other APIs, and it records tool results without verifying them.
Complete, append-only provider-side logging could also protect the record. Zero Data Retention policies may prevent providers from keeping it, and existing compliance APIs can omit thinking blocks and tool definitions or truncate entries. Recording the exchange also does not prove that a reported tool call was executed faithfully.
An agent's authority to do useful work must not include authority to rewrite the record used to evaluate that work.
