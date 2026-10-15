---
source: "https://www.pagerduty.com/resources/ai/learn/sre-agent-benchmark-report/"
hn_url: "https://news.ycombinator.com/item?id=50110584"
title: "A Better Benchmark Harness for SRE Agents"
article_title: "SRE Agent Benchmark Report - 2026 | PagerDuty"
image: "https://cdn.pagerduty.com/wp-content/uploads/2026/10/image.png"
author: "Epixors"
captured_at: "2026-10-15T21:58:17Z"
capture_tool: "hn-digest"
hn_id: 50110584
score: 1
comments: 0
posted_at: "2026-10-15T21:29:55Z"
tags:
  - hacker-news
---

# A Better Benchmark Harness for SRE Agents

- HN: [50110584](https://news.ycombinator.com/item?id=50110584)
- Source: [www.pagerduty.com](https://www.pagerduty.com/resources/ai/learn/sre-agent-benchmark-report/)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T21:29:55Z

## Translation

Title: A Better Benchmark Harness for SRE Agents
Article title: SRE Agent Benchmark Report - 2026 | PagerDuty
Description: Summary This report shows how to create a benchmark harness that injects realistic faults into Kubernetes microservices and has SRE agents work through

Article text:
SRE Agent Benchmark Report - 2026 | PagerDuty
A new era of reliability is arriving in October.
See it first
✕
The platform for mission-critical work in the modern enterprise.
End-to-end orchestration for rapid issue resolution.
Revolutionize operations at the speed of AI.
Accelerate critical work across the enterprise.
Redefine operations with AI Agents that ignite growth.
A single source of truth for system status.
Generative AI for critical operations work.
Bridge support and engineering teams.
Reduce alert noise and accelerate triage.
Learn more about new Generally Available and Early Access features
The platform for mission-critical work in the modern enterprise.
Customize your PagerDuty experience.
Get more value from PagerDuty.
View our trust and compliance practices.
Secure, reliable, extensible, and scalable.
Work how you want with PagerDuty.
Learn more about new Generally Available and Early Access features
Learn how strategic technology investments powered FOX's digital and operational transformation.
We're empowering teams with the time and efficiency to build the future.
See how we are building resilience and accelerating change.
Learn about our efforts around global employee engagement, sustainability, and more.
Get on-demand insights on AI-first operations.
A Better Benchmark Harness for SRE Agents
This report shows how to create a benchmark harness that injects realistic faults into Kubernetes microservices and has SRE agents work through them. Each agent is scored on five things: locating the fault, diagnosing the root cause, identifying a mitigation, grounding its conclusions in evidence, and explaining its reasoning clearly. A panel of judges built on different models does the scoring, which keeps any single model’s bias out of the results. Initial comparative results demonstrate that PagerDuty’s SRE Agent consistently led the benchmark across 17 scenarios on issue localization, diagnosis, evidence grounding, mitigation recommendations, and explanation quality. The results also point to two factors that matter most: the strength of the underlying model and how much context the agent can reach. Both give teams a clear direction for improving how AI resolves incidents.
Why Build a Benchmark Harness for SRE Agents
A reliable benchmark measures whether changes make an SRE agent better at resolving incidents. Building a basic SRE agent is easy. You take a powerful LLM, give it access to code and observability via MCP and put it all together in a simple harness. Job done. But to get from basic to great requires asking additional questions. How well does it work now? Is it solving the problem effectively? Is it doing it efficiently? How do you know the complexity you have added actually made things better? As you work to improve your agent, you need a reliable way to quantify the changes.
The obvious first step is to replay the agent on past incidents and compare its output with what your post-incident review says happened. We do this, and it’s useful, but it has hard limits as an evaluation method:
The reference data is only as good as the information you have from post-incident reviews, observability, collaboration tools, and other systems that you use during incident resolution. For example, often reviews are written for people, not for scoring an agent, so they’re incomplete or contested, and may not have all the required information.
Replaying past signals, events and actions is hard, and hindsight often leaks through. The agent has to see code, config, and telemetry as they were when the incident started, and nothing after. Any gap in that tooling lets the fix, postmortem, or later telemetry leak into the answer, including the title of the incident.
It only works for agents you control. We can wire time-travel tooling into our own agents. A third-party agent queries live systems through its own integrations, so it can’t be replayed fairly.
Past incidents can’t be changed. They can’t be re-run to measure variance, modified to test a specific capability, or evaluated with a data source removed (or added!).
You get the incidents you had, not necessarily the ones you need. Complex incidents are rare, many trace back to third-party outages whose cause isn’t visible in our own telemetry or code, and coverage follows whatever happened to break.
Injecting faults into a live system addresses each of these. The ground truth is exact because we caused it. Every agent investigates the present through its normal integrations. Every scenario can be repeated, varied, and ablated. And we choose which faults, technologies, and languages to cover instead of waiting for them to fail. Replay still has a role for signals a synthetic system can’t reproduce, such as incident-call transcripts (see Impact of transcript data), but for benchmarking we needed something new.
Another benefit of having a good benchmarking suite is for our customers: there are now many general-purpose and SRE-specific AI agents, and they want to know how PagerDuty’s SRE Agent compares. Answering that honestly means testing every agent the way a customer would run it, on the same faults and against the same ground truth. A well-designed benchmarking suite will allow us to do this as fairly as possible.
Our first question was whether we could adopt or adapt a benchmark that already exists. There are a number of public benchmark harnesses available , but they all presented a few drawbacks:
They normally rely on calling agents directly, with structured outputs shaped to the benchmark. When interacting with third-party agents, it is not always possible to define the interaction process in a way that is optimised for that benchmark. For example, PagerDuty’s SRE Agent uses a chat interface, it can be guided to provide a given output shape but not forced to. Should it be punished because it doesn’t output JSON?
Frequently, they run on publicly available microservice demo apps (such as OpenTelemetry’s Astronomy Shop or the SAIL Group’s DeathStarBench ) which are captured in LLM training. This means that the models underlying SRE agents already know these systems, and in some cases they are even able to recall failure injection points from training data (we saw this happen multiple times when we tested our agents on ITBench).
Few benchmarks included code as a source of data for agents to use, potentially because of the issue specified in the previous point or because the code contains the answer explicitly written in a feature flag. One of the key questions we wanted to ask is “what is the value of providing code”. To answer this question we needed a clean codebase that did not contain the answer.
Most public harnesses are built around a single fault-injection mechanism, e.g. through changing a feature flag to trigger the event or using Chaos Mesh to break the Kubernetes (k8s) substrate. Ours uses four: Helm value overrides, Chaos Mesh, real PR merges, and AWS control-plane changes, with the ability to add more as we desire such as third party failures or security incidents. This allows us to test a wider range of issues and probe agents against more complex scenarios.
For these reasons, we created our own benchmark harness. It is a purpose-built financial application that runs on Kubernetes microservices, emits telemetry, has alerting rules on Grafana and sends those alerts to PagerDuty (or the agent’s own SaaS platform). With this standard infrastructure in place, we can integrate the stack with a variety of agents, and grant them access to metrics, logs, alerting and code to use in their RCA.
This accounted for all the basic information that is ingested by all SRE agents but does not include some of the highest value information that we have for PagerDuty’s SRE Agent:
Live transcripts of incident calls are accessible via the Scribe Agent; we have found that the interaction between human and agent responders provides significant value in complex incidents.
Historical incidents and memories saved by the agent allow it to learn how a specific system works, and increase performance over a generic agent.
Neither of these are easy to include in a benchmarking suite. LLMs can generate a fake incident call, but you don’t want them leaking the answer. Past incidents need to be similar enough to be useful but not the same. Memories should be clues, not solutions. Likewise, the general-purpose agents were run in default configurations, without SRE-specific instruction files, skills, or subagents that a customer might add.
We are publishing this report to outline how we are testing our agents and highlight the strengths and limitations of our work. Future reports will cover different aspects of that testing in more detail as we continue to work on it. This report is focused on the underlying principles we have adopted when building the benchmark, and the initial results we have achieved using it.
This benchmark has been undertaken knowing that there is a conflict of interest, we are both testing and building agents based on the same vision of what an SRE agent should be. To reduce this bias, we have implemented the following protections:
A similar, neutral, prompting approach is used across all agents, with variations only to account for differences in how the agents take inputs, provide outputs and ensure they are providing the information required for the scoring.
Scoring is done using a 3-model judge ensemble from 3 different providers to prevent same provider bias (e.g. an agent using one model family might be preferentially scored by another model from the same family).
Scoring against pre-defined sub-criteria written at scenario creation.
A contamination scanner that runs on every sweep, checking for leakage from e.g., source code.
We are publishing our methodology, allowing people to critique it and raise concerns with our approach.
At the moment, we are keeping this codebase private (since we do not want it leaking into the LLMs training data), but we intend to make it available in the future to allow others to validate our results.
How Does the Benchmark Methodology Work
Our benchmark harness deploys a realistic, purpose-built financial app that runs on a number of interconnected Kubernetes services. It is a mono-repo consisting of 11 micro-services written in Go, Python, TypeScript and Elixir using Kafka, Redis, PostgreSQL and Secrets Manager, all running on EKS on AWS. They are instrumented to emit metrics and logs to Grafana Cloud . On the Grafana stack a number of alerting rules were created on telemetry such as HTTP errors, request latency, Kubernetes infrastructure health, etc. Additional services are stubbed out and will be added in future iterations using additional coding languages (e.g. Java) and technologies (e.g. MySQL) to provide a variety of scenarios and issues that can be used to test the agents. This organic growth has allowed us to test whilst we have continued to grow the benchmark.
Each service has a runbook that outlines how it works and how to debug “common faults”. These have been written by an LLM to replicate typical runbooks from effective teams and have been made available to all agents.
An external load-generation service posts requests to the financial app’s endpoints and exercises the whole system. This load-generation system does not emit MELT (metrics, events, logs and traces) data to the shared Grafana Cloud and is hidden from the agents.
We designed a number of self-contained incident scenarios, where operationally realistic faults are injected, leading alert rules to fire. These alerts are the trigger for agents to start their analyses . Faults can be injected using four different approaches:
Helm value overrides – allowing for changes in the EKS configuration, e.g., port number, pod count or load-generation workloads
Chaos Mesh scenario injection. For this reason, we have had to prevent the SRE agents from having access to the EKS logs, which would leak the presence of Chaos

[truncated]

## Original Extract

Summary This report shows how to create a benchmark harness that injects realistic faults into Kubernetes microservices and has SRE agents work through

SRE Agent Benchmark Report - 2026 | PagerDuty
A new era of reliability is arriving in October.
See it first
✕
The platform for mission-critical work in the modern enterprise.
End-to-end orchestration for rapid issue resolution.
Revolutionize operations at the speed of AI.
Accelerate critical work across the enterprise.
Redefine operations with AI Agents that ignite growth.
A single source of truth for system status.
Generative AI for critical operations work.
Bridge support and engineering teams.
Reduce alert noise and accelerate triage.
Learn more about new Generally Available and Early Access features
The platform for mission-critical work in the modern enterprise.
Customize your PagerDuty experience.
Get more value from PagerDuty.
View our trust and compliance practices.
Secure, reliable, extensible, and scalable.
Work how you want with PagerDuty.
Learn more about new Generally Available and Early Access features
Learn how strategic technology investments powered FOX's digital and operational transformation.
We're empowering teams with the time and efficiency to build the future.
See how we are building resilience and accelerating change.
Learn about our efforts around global employee engagement, sustainability, and more.
Get on-demand insights on AI-first operations.
A Better Benchmark Harness for SRE Agents
This report shows how to create a benchmark harness that injects realistic faults into Kubernetes microservices and has SRE agents work through them. Each agent is scored on five things: locating the fault, diagnosing the root cause, identifying a mitigation, grounding its conclusions in evidence, and explaining its reasoning clearly. A panel of judges built on different models does the scoring, which keeps any single model’s bias out of the results. Initial comparative results demonstrate that PagerDuty’s SRE Agent consistently led the benchmark across 17 scenarios on issue localization, diagnosis, evidence grounding, mitigation recommendations, and explanation quality. The results also point to two factors that matter most: the strength of the underlying model and how much context the agent can reach. Both give teams a clear direction for improving how AI resolves incidents.
Why Build a Benchmark Harness for SRE Agents
A reliable benchmark measures whether changes make an SRE agent better at resolving incidents. Building a basic SRE agent is easy. You take a powerful LLM, give it access to code and observability via MCP and put it all together in a simple harness. Job done. But to get from basic to great requires asking additional questions. How well does it work now? Is it solving the problem effectively? Is it doing it efficiently? How do you know the complexity you have added actually made things better? As you work to improve your agent, you need a reliable way to quantify the changes.
The obvious first step is to replay the agent on past incidents and compare its output with what your post-incident review says happened. We do this, and it’s useful, but it has hard limits as an evaluation method:
The reference data is only as good as the information you have from post-incident reviews, observability, collaboration tools, and other systems that you use during incident resolution. For example, often reviews are written for people, not for scoring an agent, so they’re incomplete or contested, and may not have all the required information.
Replaying past signals, events and actions is hard, and hindsight often leaks through. The agent has to see code, config, and telemetry as they were when the incident started, and nothing after. Any gap in that tooling lets the fix, postmortem, or later telemetry leak into the answer, including the title of the incident.
It only works for agents you control. We can wire time-travel tooling into our own agents. A third-party agent queries live systems through its own integrations, so it can’t be replayed fairly.
Past incidents can’t be changed. They can’t be re-run to measure variance, modified to test a specific capability, or evaluated with a data source removed (or added!).
You get the incidents you had, not necessarily the ones you need. Complex incidents are rare, many trace back to third-party outages whose cause isn’t visible in our own telemetry or code, and coverage follows whatever happened to break.
Injecting faults into a live system addresses each of these. The ground truth is exact because we caused it. Every agent investigates the present through its normal integrations. Every scenario can be repeated, varied, and ablated. And we choose which faults, technologies, and languages to cover instead of waiting for them to fail. Replay still has a role for signals a synthetic system can’t reproduce, such as incident-call transcripts (see Impact of transcript data), but for benchmarking we needed something new.
Another benefit of having a good benchmarking suite is for our customers: there are now many general-purpose and SRE-specific AI agents, and they want to know how PagerDuty’s SRE Agent compares. Answering that honestly means testing every agent the way a customer would run it, on the same faults and against the same ground truth. A well-designed benchmarking suite will allow us to do this as fairly as possible.
Our first question was whether we could adopt or adapt a benchmark that already exists. There are a number of public benchmark harnesses available , but they all presented a few drawbacks:
They normally rely on calling agents directly, with structured outputs shaped to the benchmark. When interacting with third-party agents, it is not always possible to define the interaction process in a way that is optimised for that benchmark. For example, PagerDuty’s SRE Agent uses a chat interface, it can be guided to provide a given output shape but not forced to. Should it be punished because it doesn’t output JSON?
Frequently, they run on publicly available microservice demo apps (such as OpenTelemetry’s Astronomy Shop or the SAIL Group’s DeathStarBench ) which are captured in LLM training. This means that the models underlying SRE agents already know these systems, and in some cases they are even able to recall failure injection points from training data (we saw this happen multiple times when we tested our agents on ITBench).
Few benchmarks included code as a source of data for agents to use, potentially because of the issue specified in the previous point or because the code contains the answer explicitly written in a feature flag. One of the key questions we wanted to ask is “what is the value of providing code”. To answer this question we needed a clean codebase that did not contain the answer.
Most public harnesses are built around a single fault-injection mechanism, e.g. through changing a feature flag to trigger the event or using Chaos Mesh to break the Kubernetes (k8s) substrate. Ours uses four: Helm value overrides, Chaos Mesh, real PR merges, and AWS control-plane changes, with the ability to add more as we desire such as third party failures or security incidents. This allows us to test a wider range of issues and probe agents against more complex scenarios.
For these reasons, we created our own benchmark harness. It is a purpose-built financial application that runs on Kubernetes microservices, emits telemetry, has alerting rules on Grafana and sends those alerts to PagerDuty (or the agent’s own SaaS platform). With this standard infrastructure in place, we can integrate the stack with a variety of agents, and grant them access to metrics, logs, alerting and code to use in their RCA.
This accounted for all the basic information that is ingested by all SRE agents but does not include some of the highest value information that we have for PagerDuty’s SRE Agent:
Live transcripts of incident calls are accessible via the Scribe Agent; we have found that the interaction between human and agent responders provides significant value in complex incidents.
Historical incidents and memories saved by the agent allow it to learn how a specific system works, and increase performance over a generic agent.
Neither of these are easy to include in a benchmarking suite. LLMs can generate a fake incident call, but you don’t want them leaking the answer. Past incidents need to be similar enough to be useful but not the same. Memories should be clues, not solutions. Likewise, the general-purpose agents were run in default configurations, without SRE-specific instruction files, skills, or subagents that a customer might add.
We are publishing this report to outline how we are testing our agents and highlight the strengths and limitations of our work. Future reports will cover different aspects of that testing in more detail as we continue to work on it. This report is focused on the underlying principles we have adopted when building the benchmark, and the initial results we have achieved using it.
This benchmark has been undertaken knowing that there is a conflict of interest, we are both testing and building agents based on the same vision of what an SRE agent should be. To reduce this bias, we have implemented the following protections:
A similar, neutral, prompting approach is used across all agents, with variations only to account for differences in how the agents take inputs, provide outputs and ensure they are providing the information required for the scoring.
Scoring is done using a 3-model judge ensemble from 3 different providers to prevent same provider bias (e.g. an agent using one model family might be preferentially scored by another model from the same family).
Scoring against pre-defined sub-criteria written at scenario creation.
A contamination scanner that runs on every sweep, checking for leakage from e.g., source code.
We are publishing our methodology, allowing people to critique it and raise concerns with our approach.
At the moment, we are keeping this codebase private (since we do not want it leaking into the LLMs training data), but we intend to make it available in the future to allow others to validate our results.
How Does the Benchmark Methodology Work
Our benchmark harness deploys a realistic, purpose-built financial app that runs on a number of interconnected Kubernetes services. It is a mono-repo consisting of 11 micro-services written in Go, Python, TypeScript and Elixir using Kafka, Redis, PostgreSQL and Secrets Manager, all running on EKS on AWS. They are instrumented to emit metrics and logs to Grafana Cloud . On the Grafana stack a number of alerting rules were created on telemetry such as HTTP errors, request latency, Kubernetes infrastructure health, etc. Additional services are stubbed out and will be added in future iterations using additional coding languages (e.g. Java) and technologies (e.g. MySQL) to provide a variety of scenarios and issues that can be used to test the agents. This organic growth has allowed us to test whilst we have continued to grow the benchmark.
Each service has a runbook that outlines how it works and how to debug “common faults”. These have been written by an LLM to replicate typical runbooks from effective teams and have been made available to all agents.
An external load-generation service posts requests to the financial app’s endpoints and exercises the whole system. This load-generation system does not emit MELT (metrics, events, logs and traces) data to the shared Grafana Cloud and is hidden from the agents.
We designed a number of self-contained incident scenarios, where operationally realistic faults are injected, leading alert rules to fire. These alerts are the trigger for agents to start their analyses . Faults can be injected using four different approaches:
Helm value overrides – allowing for changes in the EKS configuration, e.g., port number, pod count or load-generation workloads
Chaos Mesh scenario injection. For this reason, we have had to prevent the SRE agents from having access to the EKS logs, which would leak the presence of Chaos

[truncated]
