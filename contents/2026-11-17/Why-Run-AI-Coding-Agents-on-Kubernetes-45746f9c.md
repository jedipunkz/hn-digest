---
source: "https://www.okteto.com/blog/why-run-ai-coding-agents-on-kubernetes/"
hn_url: "https://news.ycombinator.com/item?id=50533978"
title: "Why Run AI Coding Agents on Kubernetes?"
article_title: "AI Coding Agents on Kubernetes: Security, Scale, and Control | Okteto"
image: ""
author: "pchico83"
captured_at: "2026-11-17T16:43:30Z"
capture_tool: "hn-digest"
hn_id: 50533978
score: 1
comments: 0
posted_at: "2026-11-17T16:37:26Z"
tags:
  - hacker-news
---

# Why Run AI Coding Agents on Kubernetes?

- HN: [50533978](https://news.ycombinator.com/item?id=50533978)
- Source: [www.okteto.com](https://www.okteto.com/blog/why-run-ai-coding-agents-on-kubernetes/)
- Score: 1
- Comments: 0
- Posted: 2026-11-17T16:37:26Z

## Translation

Title: Why Run AI Coding Agents on Kubernetes?
Article title: AI Coding Agents on Kubernetes: Security, Scale, and Control | Okteto

Article text:
Skip to main content Search Docs Tutorials Blog Community Get a demo Start for Free Why Run AI Coding Agents on Kubernetes?
AI is an amplifier. Coding agents multiply speed where your engineering platform has structure, and they multiply risk where it doesn't. That is why the most important question about agents is no longer which one your developers use. It is where those agents run.
In many companies, the answer today is "on every developer's laptop". Someone installs Claude Code, Codex, or Cursor, points it at a repository, and gives it a terminal. It works well enough for one developer and one agent. It stops working when the platform team is asked to roll agents out to hundreds of developers, each running several agents at once, in a company that has to answer to auditors.
The premise: agents run on your infrastructure ​
For organizations that require agent execution to stay within infrastructure they control, the question is how to provide that execution environment consistently across teams. Running the agent on your infrastructure and choosing where model inference happens are separate architectural decisions: an internal AI gateway, by itself, doesn't guarantee that prompts and tool outputs stay within your infrastructure.
Developer laptops can be secured, but they shouldn't have to serve as the execution platform for every autonomous task. For teams already operating Kubernetes, the cluster provides a familiar foundation for centrally managed agent workloads, with explicit identities, resource limits, network policies, and lifecycle controls. These controls are not automatic. The advantage is having a platform where they can be defined, enforced, and reviewed consistently, instead of building a second platform for agents.
The rest of this post compares the two operating models on four points: security, reproducibility, scale, and observability.
Agents can run securely on a laptop: agent harnesses offer file system and network sandboxing . But it depends on each developer to set it up, keep it on, and keep it current. Without it, the agent runs as the developer, with their cloud credentials, kubeconfig, SSH keys, and VPN access: Simon Willison's "lethal trifecta" of private data, untrusted content, and a way to send data out. The s1ngularity attack leaked 2,349 credentials from developer machines, and at PocketOS an agent found an API token in an unrelated file and deleted the production database.
On Kubernetes, the platform team defines which credentials and destinations each agent can reach. Each task runs in its own pod, with a short-lived token bound to the pod (or none), workload identity instead of static cloud keys, and RBAC on the Kubernetes API. The Pod Security Standards Restricted profile keeps pods non-root and unprivileged, a RuntimeClass can add a sandboxed runtime such as gVisor or Kata Containers , and NetworkPolicies limit egress to the destinations the agent needs. Agent Sandbox , a Kubernetes SIG Apps subproject, packages this pattern as a Sandbox API, with warm pools for fast startup.
Each developer configures their own agent: their own toolchain versions, their own MCP servers, their own permission settings, their own CLAUDE.md or AGENTS.md . Over time, these configurations diverge. Security researchers call this posture drift : an agent that starts with three approved tools ends up with eleven after developers add servers one at a time, and nobody reviews the result. So the same task, given to the same agent, behaves differently depending on whose laptop it runs on. When an agent's change fails in CI, you can't tell whether the change is wrong or the agent's environment was.
On Kubernetes, the agent's runtime is a versioned definition in Git: a container image pinned by digest, so it can't change under a moving tag, and a manifest with the permissions each kind of task needs. When the platform team updates the image or the policy, the change is reviewed like any other code, and every agent that uses that definition gets it.
Agents are most useful on long tasks and in parallel. A laptop is bad at both. An agent that runs for hours or days takes the machine with it: it runs tests, compiles, and competes with the developer for CPU and memory, and the task dies when the laptop sleeps, hangs, or loses its connection. The laptop's resources also set a hard ceiling on how many agents can run at once. As a principal engineer at Faire put it, there are ways to parallelize agents on your local machine, "but it's much more complicated" .
On Kubernetes, agent tasks can run as Jobs, scheduled across the cluster, with the cluster autoscaler or Karpenter adding nodes when demand grows. A developer can start ten agents and close the laptop. activeDeadlineSeconds stops runaway tasks, ttlSecondsAfterFinished deletes finished Jobs and their pods, and ResourceQuotas cap how much compute each team's agents can use.
When agents run on laptops, the platform team has no inventory of which agents exist, what they can reach, what they run, or what they cost. In most organizations, there are already agents running things that nobody notices. You can't govern what you can't see, and you can't enforce governance through prompts and process documents on hundreds of machines you don't control.
Kubernetes provides a foundation for collecting workload logs, resource metrics, and API audit events. To understand an agent's actions, correlate that infrastructure telemetry with agent-level records: task identifiers, tool calls, execution results, and model usage. Together, these records help platform teams investigate failures, review access, and attribute costs.
Local and Kubernetes execution side by side ​
Agents belong where your workloads already live ​
Kubernetes doesn't make agents secure by default. It gives platform teams a foundation to provide governed execution for agents, under the same identities, policies, and observability as the rest of their workloads.

## Original Extract

Skip to main content Search Docs Tutorials Blog Community Get a demo Start for Free Why Run AI Coding Agents on Kubernetes?
AI is an amplifier. Coding agents multiply speed where your engineering platform has structure, and they multiply risk where it doesn't. That is why the most important question about agents is no longer which one your developers use. It is where those agents run.
In many companies, the answer today is "on every developer's laptop". Someone installs Claude Code, Codex, or Cursor, points it at a repository, and gives it a terminal. It works well enough for one developer and one agent. It stops working when the platform team is asked to roll agents out to hundreds of developers, each running several agents at once, in a company that has to answer to auditors.
The premise: agents run on your infrastructure ​
For organizations that require agent execution to stay within infrastructure they control, the question is how to provide that execution environment consistently across teams. Running the agent on your infrastructure and choosing where model inference happens are separate architectural decisions: an internal AI gateway, by itself, doesn't guarantee that prompts and tool outputs stay within your infrastructure.
Developer laptops can be secured, but they shouldn't have to serve as the execution platform for every autonomous task. For teams already operating Kubernetes, the cluster provides a familiar foundation for centrally managed agent workloads, with explicit identities, resource limits, network policies, and lifecycle controls. These controls are not automatic. The advantage is having a platform where they can be defined, enforced, and reviewed consistently, instead of building a second platform for agents.
The rest of this post compares the two operating models on four points: security, reproducibility, scale, and observability.
Agents can run securely on a laptop: agent harnesses offer file system and network sandboxing . But it depends on each developer to set it up, keep it on, and keep it current. Without it, the agent runs as the developer, with their cloud credentials, kubeconfig, SSH keys, and VPN access: Simon Willison's "lethal trifecta" of private data, untrusted content, and a way to send data out. The s1ngularity attack leaked 2,349 credentials from developer machines, and at PocketOS an agent found an API token in an unrelated file and deleted the production database.
On Kubernetes, the platform team defines which credentials and destinations each agent can reach. Each task runs in its own pod, with a short-lived token bound to the pod (or none), workload identity instead of static cloud keys, and RBAC on the Kubernetes API. The Pod Security Standards Restricted profile keeps pods non-root and unprivileged, a RuntimeClass can add a sandboxed runtime such as gVisor or Kata Containers , and NetworkPolicies limit egress to the destinations the agent needs. Agent Sandbox , a Kubernetes SIG Apps subproject, packages this pattern as a Sandbox API, with warm pools for fast startup.
Each developer configures their own agent: their own toolchain versions, their own MCP servers, their own permission settings, their own CLAUDE.md or AGENTS.md . Over time, these configurations diverge. Security researchers call this posture drift : an agent that starts with three approved tools ends up with eleven after developers add servers one at a time, and nobody reviews the result. So the same task, given to the same agent, behaves differently depending on whose laptop it runs on. When an agent's change fails in CI, you can't tell whether the change is wrong or the agent's environment was.
On Kubernetes, the agent's runtime is a versioned definition in Git: a container image pinned by digest, so it can't change under a moving tag, and a manifest with the permissions each kind of task needs. When the platform team updates the image or the policy, the change is reviewed like any other code, and every agent that uses that definition gets it.
Agents are most useful on long tasks and in parallel. A laptop is bad at both. An agent that runs for hours or days takes the machine with it: it runs tests, compiles, and competes with the developer for CPU and memory, and the task dies when the laptop sleeps, hangs, or loses its connection. The laptop's resources also set a hard ceiling on how many agents can run at once. As a principal engineer at Faire put it, there are ways to parallelize agents on your local machine, "but it's much more complicated" .
On Kubernetes, agent tasks can run as Jobs, scheduled across the cluster, with the cluster autoscaler or Karpenter adding nodes when demand grows. A developer can start ten agents and close the laptop. activeDeadlineSeconds stops runaway tasks, ttlSecondsAfterFinished deletes finished Jobs and their pods, and ResourceQuotas cap how much compute each team's agents can use.
When agents run on laptops, the platform team has no inventory of which agents exist, what they can reach, what they run, or what they cost. In most organizations, there are already agents running things that nobody notices. You can't govern what you can't see, and you can't enforce governance through prompts and process documents on hundreds of machines you don't control.
Kubernetes provides a foundation for collecting workload logs, resource metrics, and API audit events. To understand an agent's actions, correlate that infrastructure telemetry with agent-level records: task identifiers, tool calls, execution results, and model usage. Together, these records help platform teams investigate failures, review access, and attribute costs.
Local and Kubernetes execution side by side ​
Agents belong where your workloads already live ​
Kubernetes doesn't make agents secure by default. It gives platform teams a foundation to provide governed execution for agents, under the same identities, policies, and observability as the rest of their workloads.
