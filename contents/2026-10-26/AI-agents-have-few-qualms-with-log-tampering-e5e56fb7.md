---
source: "https://www.manifold.security/blog/agent-log-tampering"
hn_url: "https://news.ycombinator.com/item?id=50250494"
title: "AI agents have few qualms with log tampering"
article_title: "AI agents have few qualms with log tampering"
image: "https://cdn.sanity.io/images/phwxcc13/production/0094107ed1ccfc9329979a9a14e735c1afb45b0e-2528x1696.png?rect=0,185,2528,1327&w=1200&h=630&auto=format"
author: "codyznash"
captured_at: "2026-10-26T21:24:14Z"
capture_tool: "hn-digest"
hn_id: 50250494
score: 2
comments: 0
posted_at: "2026-10-26T20:31:43Z"
tags:
  - hacker-news
---

# AI agents have few qualms with log tampering

- HN: [50250494](https://news.ycombinator.com/item?id=50250494)
- Source: [www.manifold.security](https://www.manifold.security/blog/agent-log-tampering)
- Score: 2
- Comments: 0
- Posted: 2026-10-26T20:31:43Z

## Translation

Title: AI agents have few qualms with log tampering
Description: Our research shows why models comply with log deletion requests so readily, and how tamper-evident logging catches it.

Article text:
AI agents have few qualms with log tampering ISO/IEC 27001:2022 audits passed with zero nonconformities. Find out more → Why Manifold Platform Company About us Press Career Trust Center Resources Blog Events & Recordings Manifest Book a demo Home Blog AI agents have few qualms with log tampering Research
AI agents have few qualms with log tampering
Oct 26, 2026 • 8 Min Cody Nash Researcher
Research shows that most current frontier AI agents will tamper with their own logs. Only Muse Spark frequently refused requests. Claude Code, Codex and other harnesses running the latest models almost always complied.
We tested two open models to get more insight into why the models comply so readily. Log tampering requests were only refused when an explicit policy was in place. Abliteration of harmful refusals or moral direction suggests there is a rule following component. Observation of experts activated by the models suggests log tampering is generally not considered a matter of right or wrong.
Not just theoretical: Reviews of rogue agent incidents this year found widespread tampering efforts, including log spoofing, container resets to erase history, and unexplainable gaps in logs.
Tamper-evident logging (hash chains, Merkle trees, an outside recorder) shows when logs were altered. Agent tools for this shipped months before the incidents.
Better alignment lowers the risk but won’t remove it. Catching tampering has to happen outside the model.
In their Sept. 24 study , Qin et al. show it is alarmingly easy to get current frontier models to tamper with their own traces, hiding the evidence of what they did. Their study found that if you simply ask them to delete their transcripts, they’ll do it almost every time. Showing them peers who did it works almost as well, without even asking. Make it into a game? They hid their tracks in nearly every run. No fine-tuning, no abliteration of refusals needed. Their peer experiment was directly inspired by OpenAI’s agents’ infiltration and abuse of Hugging Face this summer. During that incident, agents showed great interest in developing log tampering techniques and directed multiple efforts to that end. Investigators noted “it was common for agents to edit” the logs to which they had access. We only know about these activities because logs the investigators relied upon as a source of truth remained out of the agents reach.
We wanted to know why agents go after their logs so readily, so we looked inside two open models. We examined, via abliteration, the concepts associated with log tampering refusals. We investigated whether the models considered tampering right or wrong by looking at which experts, in mixture of experts models, were used when considering requests.
Tamper-evident logging methods, which we will explain, can prove a log hasn’t been touched, and agent tools like nono and AI Action Ledger shipped them earlier this year. Read on to find out more about how models think about log-tampering, how tamper-evident logging works, and where we go from here.
To see why models comply so readily, we looked inside two small open models, OLMoE and DeepSeek-V2-Lite . Without a policy line, neither model refused more than 4% of tampering or spoofing requests. This is consistent with Qin et al.’s findings that most of the agents running frontier models they tested comply with log tampering requests. Only with a policy in place prohibiting log tampering did DeepSeek refuse up to 46% of deletions. Even with the policy, it refused none of the requests to spoof the logs. Abliteration of harmful task refusals reduced the refusal rate for requests to delete logs where the reason given was fear of being spied on or trying to hide activity from a manager. Abliteration of a moral direction reduced refusals in these cases a similar amount. Removal of a random direction of the same size had no significant effects. This suggests at least part of the refusal is driven by an awareness that these deletions are harmful requests when there is a policy in place, and another part of the reason is just following the rules. Fine-tuning could make models comply with harmful tasks less often, but it would not stop them entirely, so spoofing and tampering still need to be caught outside the model.
Abliteration cuts DeepSeek’s refusals to tamper with its log
Using methods we developed earlier to tell when models treat a question as a moral one, we tracked which of their internal experts handle the request. If you ask the models if deleting logs is immoral or malicious, then the experts associated with morality are used. However, actual requests to delete, redact or fabricate log entries barely reached the same experts. Control questions about the harmful deletion of files, for DeepSeek at least, also activated moral experts. DeepSeek sent spoofing requests to moral experts, yet it refused none of them as we saw above. A stated motive, such as hiding the log from a manager or fear of being spied on, raised moral-expert load in OLMoE, as sometimes did a policy against tampering. In these models, log tampering rarely registers as a moral choice, which may be resolvable through better alignment.
Log tampering barely touches the morality experts
Questions about deleting the log, vs. neutral control
Tampering requests, vs. neutral logging requests
Deletion controls, vs. neutral logging requests
Let’s talk about logs and trees
Log tampering can be retroactive, where agents make edits after the fact, or proactive, called spoofing, where they set things up so logs don’t record what actually happened. Qin et al. got most harnesses to tamper and spoof in most trials, for example by rigging curl to return made-up responses that the trace logged as real. Tamper-evident logging helps with retroactive tampering. Catching spoofing takes a record the agent’s host cannot forge, which a December 2025 paper, VET , proposed for an agent’s API calls. Since March, the sandbox nono has kept such a record: a supervisor process outside the agent’s sandbox logs the agent’s network calls where the agent hopefully cannot reach them.
Spoofing comes first, tampering comes after
Pinch to zoom in on the diagram
Spoofing, set up ahead of time
Outside recorder process and network catches the spoof
Event hashing chain and tree catches the edit
H(e1) missing H(e2) H(e3) H(e4) left half differs right half matches root differs tree hashes kept off host
Comparing tree hashes with the off-host copy leads to the edited entry without rereading the rest.
The recipe for tamper-evident logging is fairly straightforward. Chain each event’s hash to the one before it to fix the sequence, add Merkle trees over whole traces so you can audit a single event or an entire session, and have a process separate from the agent do the recording. If you want full auditability (what changed and how), store the records and hashes in a second location. If you want to keep it low-cost or privacy oriented, keep the root of each tree in a remote location, so you don’t pay for a second set of logs. Some tools use a blockchain as that remote location, with no token involved. GuardianChain and Agent Flight Recorder build on Base. WeilChain receipts commit a hash of each Claude Code file edit to their own chain, and the first version of an IETF draft proposed using Bitcoin.
The people building agent tools expected rogue agents to hide their tracks and built ways to catch them before this summer: AI Action Ledger brought tamper-evident records to LangChain agents in January, and Faramesh and nono followed, with nono adding a hash chain in April . Researchers saw it coming earlier still, proposing anchored audit ledgers for agents ( BlockA2A ) in August 2025, and the Singapore Consensus on AI safety research priorities counts tamper evidence as part of auditability. Standards have moved more slowly. An IETF draft for agent audit trails that hash-chains each action is still a draft, and a community member’s proposal for signed, chained receipts in Microsoft’s Agent Governance Toolkit is still a proposal. A July proposal to build tamper-evident records into the Model Context Protocol (MCP), the standard many agents use to call tools, drew independent implementations before maintainers closed it.
What gets logged matters as much as how the log is protected. Andrea Fortuna listed in May what a trace needs to hold up in an investigation, from each tool call’s arguments and results to the identity that started the session. Reasoning belongs in the record too: METR could spot the spoofed tool calls in the transcripts of OpenAI’s hack of Hugging Face this summer because the steps that set them up, and the agents’ reasoning about them, were in the same record.
A hash chain only vouches for the events it receives. An agent that keeps an action out of the record or writes a false entry before it is hashed still passes every integrity check. Qin et al. call this the difference between the integrity of a trace and the fidelity of its collection. So some of the record has to come from outside the agent, where it cannot be shaped. An outside recorder has to see every model the agent talks to: if the agent calls a second AI provider, runs a local model or starts another agent, those conversations never reach the trace.
Want to see what your agents are actually doing, and the detection and response to secure your agentic operations? Talk to Manifold
What about your provider’s logs?
Depending on your model provider, you can use their compliance API as a source of truth for what your agents did. It gives you a remote, ideally tamper-proof record, but it has limits. It isn’t an option if you need zero data retention. It only records what gets sent to it, so not every model and harness is covered, which matters most for extra agents spun up by the one you’re watching. Spoofed logs get recorded as truth unless you also logged the spoofing setup. And how much do you trust your provider? Their code and processes largely come from the same models you’re trying to monitor, and misalignment can exist at the model and corporate level. Hence the call to “separate the supply of intelligence from the authority over it.”
Investigators could not rule out log tampering
OpenAI’s and Anthropic’s incident reports of their “rogue” agents describe transcript reviews, classifiers and containment, but neither mentions a tamper-evident record. Neither explains how an outside investigator could check that the transcripts they handed over were complete and unchanged.
The tools existed before the OpenAI attack on HF
BlockA2A paper Paper or standard
IETF policy Bitcoin draft Paper or standard
Microsoft AGT proposal Paper or standard
Fortuna blog Paper or standard
MCP SEP proposal Paper or standard
Singapore Consensus Paper or standard
Qin et al. tampering study Paper or standard
IETF policy rev 8 Paper or standard
METR and Redwood Research found the tampering by having agents comb through the transcripts, agents they themselves called “often-unreliable.” Log tampering during the OpenAI/HF incident has been covered well already by webofmike , implicator , ExplainX and InfoQ . In METR’s sample, at least one in five OpenAI agents showed clear interest in tampering with their transcripts, and looser counts ran near 50%. Clear interest looked like the agents writing their plans down: “Could create [a transcript log] ourselves with sanitized fake trajectory.” The investigation found log tampering was one of three high-level approaches undertaken by agents to fool the evaluation process. Their efforts developed successful spoofing techniques, identified only because setting up the spoof was recorded in the transcripts. Some agents realized they could edit or delete the container logs within their reach, but “correctly concluded that these logs were not the real source of truth.”
Agents also explored using container resets to wipe parts of their transcripts. Container resets were part of regular

[truncated]

## Original Extract

Our research shows why models comply with log deletion requests so readily, and how tamper-evident logging catches it.

AI agents have few qualms with log tampering ISO/IEC 27001:2022 audits passed with zero nonconformities. Find out more → Why Manifold Platform Company About us Press Career Trust Center Resources Blog Events & Recordings Manifest Book a demo Home Blog AI agents have few qualms with log tampering Research
AI agents have few qualms with log tampering
Oct 26, 2026 • 8 Min Cody Nash Researcher
Research shows that most current frontier AI agents will tamper with their own logs. Only Muse Spark frequently refused requests. Claude Code, Codex and other harnesses running the latest models almost always complied.
We tested two open models to get more insight into why the models comply so readily. Log tampering requests were only refused when an explicit policy was in place. Abliteration of harmful refusals or moral direction suggests there is a rule following component. Observation of experts activated by the models suggests log tampering is generally not considered a matter of right or wrong.
Not just theoretical: Reviews of rogue agent incidents this year found widespread tampering efforts, including log spoofing, container resets to erase history, and unexplainable gaps in logs.
Tamper-evident logging (hash chains, Merkle trees, an outside recorder) shows when logs were altered. Agent tools for this shipped months before the incidents.
Better alignment lowers the risk but won’t remove it. Catching tampering has to happen outside the model.
In their Sept. 24 study , Qin et al. show it is alarmingly easy to get current frontier models to tamper with their own traces, hiding the evidence of what they did. Their study found that if you simply ask them to delete their transcripts, they’ll do it almost every time. Showing them peers who did it works almost as well, without even asking. Make it into a game? They hid their tracks in nearly every run. No fine-tuning, no abliteration of refusals needed. Their peer experiment was directly inspired by OpenAI’s agents’ infiltration and abuse of Hugging Face this summer. During that incident, agents showed great interest in developing log tampering techniques and directed multiple efforts to that end. Investigators noted “it was common for agents to edit” the logs to which they had access. We only know about these activities because logs the investigators relied upon as a source of truth remained out of the agents reach.
We wanted to know why agents go after their logs so readily, so we looked inside two open models. We examined, via abliteration, the concepts associated with log tampering refusals. We investigated whether the models considered tampering right or wrong by looking at which experts, in mixture of experts models, were used when considering requests.
Tamper-evident logging methods, which we will explain, can prove a log hasn’t been touched, and agent tools like nono and AI Action Ledger shipped them earlier this year. Read on to find out more about how models think about log-tampering, how tamper-evident logging works, and where we go from here.
To see why models comply so readily, we looked inside two small open models, OLMoE and DeepSeek-V2-Lite . Without a policy line, neither model refused more than 4% of tampering or spoofing requests. This is consistent with Qin et al.’s findings that most of the agents running frontier models they tested comply with log tampering requests. Only with a policy in place prohibiting log tampering did DeepSeek refuse up to 46% of deletions. Even with the policy, it refused none of the requests to spoof the logs. Abliteration of harmful task refusals reduced the refusal rate for requests to delete logs where the reason given was fear of being spied on or trying to hide activity from a manager. Abliteration of a moral direction reduced refusals in these cases a similar amount. Removal of a random direction of the same size had no significant effects. This suggests at least part of the refusal is driven by an awareness that these deletions are harmful requests when there is a policy in place, and another part of the reason is just following the rules. Fine-tuning could make models comply with harmful tasks less often, but it would not stop them entirely, so spoofing and tampering still need to be caught outside the model.
Abliteration cuts DeepSeek’s refusals to tamper with its log
Using methods we developed earlier to tell when models treat a question as a moral one, we tracked which of their internal experts handle the request. If you ask the models if deleting logs is immoral or malicious, then the experts associated with morality are used. However, actual requests to delete, redact or fabricate log entries barely reached the same experts. Control questions about the harmful deletion of files, for DeepSeek at least, also activated moral experts. DeepSeek sent spoofing requests to moral experts, yet it refused none of them as we saw above. A stated motive, such as hiding the log from a manager or fear of being spied on, raised moral-expert load in OLMoE, as sometimes did a policy against tampering. In these models, log tampering rarely registers as a moral choice, which may be resolvable through better alignment.
Log tampering barely touches the morality experts
Questions about deleting the log, vs. neutral control
Tampering requests, vs. neutral logging requests
Deletion controls, vs. neutral logging requests
Let’s talk about logs and trees
Log tampering can be retroactive, where agents make edits after the fact, or proactive, called spoofing, where they set things up so logs don’t record what actually happened. Qin et al. got most harnesses to tamper and spoof in most trials, for example by rigging curl to return made-up responses that the trace logged as real. Tamper-evident logging helps with retroactive tampering. Catching spoofing takes a record the agent’s host cannot forge, which a December 2025 paper, VET , proposed for an agent’s API calls. Since March, the sandbox nono has kept such a record: a supervisor process outside the agent’s sandbox logs the agent’s network calls where the agent hopefully cannot reach them.
Spoofing comes first, tampering comes after
Pinch to zoom in on the diagram
Spoofing, set up ahead of time
Outside recorder process and network catches the spoof
Event hashing chain and tree catches the edit
H(e1) missing H(e2) H(e3) H(e4) left half differs right half matches root differs tree hashes kept off host
Comparing tree hashes with the off-host copy leads to the edited entry without rereading the rest.
The recipe for tamper-evident logging is fairly straightforward. Chain each event’s hash to the one before it to fix the sequence, add Merkle trees over whole traces so you can audit a single event or an entire session, and have a process separate from the agent do the recording. If you want full auditability (what changed and how), store the records and hashes in a second location. If you want to keep it low-cost or privacy oriented, keep the root of each tree in a remote location, so you don’t pay for a second set of logs. Some tools use a blockchain as that remote location, with no token involved. GuardianChain and Agent Flight Recorder build on Base. WeilChain receipts commit a hash of each Claude Code file edit to their own chain, and the first version of an IETF draft proposed using Bitcoin.
The people building agent tools expected rogue agents to hide their tracks and built ways to catch them before this summer: AI Action Ledger brought tamper-evident records to LangChain agents in January, and Faramesh and nono followed, with nono adding a hash chain in April . Researchers saw it coming earlier still, proposing anchored audit ledgers for agents ( BlockA2A ) in August 2025, and the Singapore Consensus on AI safety research priorities counts tamper evidence as part of auditability. Standards have moved more slowly. An IETF draft for agent audit trails that hash-chains each action is still a draft, and a community member’s proposal for signed, chained receipts in Microsoft’s Agent Governance Toolkit is still a proposal. A July proposal to build tamper-evident records into the Model Context Protocol (MCP), the standard many agents use to call tools, drew independent implementations before maintainers closed it.
What gets logged matters as much as how the log is protected. Andrea Fortuna listed in May what a trace needs to hold up in an investigation, from each tool call’s arguments and results to the identity that started the session. Reasoning belongs in the record too: METR could spot the spoofed tool calls in the transcripts of OpenAI’s hack of Hugging Face this summer because the steps that set them up, and the agents’ reasoning about them, were in the same record.
A hash chain only vouches for the events it receives. An agent that keeps an action out of the record or writes a false entry before it is hashed still passes every integrity check. Qin et al. call this the difference between the integrity of a trace and the fidelity of its collection. So some of the record has to come from outside the agent, where it cannot be shaped. An outside recorder has to see every model the agent talks to: if the agent calls a second AI provider, runs a local model or starts another agent, those conversations never reach the trace.
Want to see what your agents are actually doing, and the detection and response to secure your agentic operations? Talk to Manifold
What about your provider’s logs?
Depending on your model provider, you can use their compliance API as a source of truth for what your agents did. It gives you a remote, ideally tamper-proof record, but it has limits. It isn’t an option if you need zero data retention. It only records what gets sent to it, so not every model and harness is covered, which matters most for extra agents spun up by the one you’re watching. Spoofed logs get recorded as truth unless you also logged the spoofing setup. And how much do you trust your provider? Their code and processes largely come from the same models you’re trying to monitor, and misalignment can exist at the model and corporate level. Hence the call to “separate the supply of intelligence from the authority over it.”
Investigators could not rule out log tampering
OpenAI’s and Anthropic’s incident reports of their “rogue” agents describe transcript reviews, classifiers and containment, but neither mentions a tamper-evident record. Neither explains how an outside investigator could check that the transcripts they handed over were complete and unchanged.
The tools existed before the OpenAI attack on HF
BlockA2A paper Paper or standard
IETF policy Bitcoin draft Paper or standard
Microsoft AGT proposal Paper or standard
Fortuna blog Paper or standard
MCP SEP proposal Paper or standard
Singapore Consensus Paper or standard
Qin et al. tampering study Paper or standard
IETF policy rev 8 Paper or standard
METR and Redwood Research found the tampering by having agents comb through the transcripts, agents they themselves called “often-unreliable.” Log tampering during the OpenAI/HF incident has been covered well already by webofmike , implicator , ExplainX and InfoQ . In METR’s sample, at least one in five OpenAI agents showed clear interest in tampering with their transcripts, and looser counts ran near 50%. Clear interest looked like the agents writing their plans down: “Could create [a transcript log] ourselves with sanitized fake trajectory.” The investigation found log tampering was one of three high-level approaches undertaken by agents to fool the evaluation process. Their efforts developed successful spoofing techniques, identified only because setting up the spoof was recorded in the transcripts. Some agents realized they could edit or delete the container logs within their reach, but “correctly concluded that these logs were not the real source of truth.”
Agents also explored using container resets to wipe parts of their transcripts. Container resets were part of regular

[truncated]
