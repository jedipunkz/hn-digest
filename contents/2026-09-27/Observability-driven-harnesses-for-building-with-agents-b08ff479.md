---
source: "https://www.datadoghq.com/blog/ai/harness-first-agents/"
hn_url: "https://news.ycombinator.com/item?id=49868243"
title: "Observability-driven harnesses for building with agents"
article_title: "Closing the verification loop: Observability-driven harnesses for building with agents | Datadog"
image: "https://web-assets.dd-static.net/42588/1776351980-harness-first-agents-harness-first-agents-hero.png"
author: "sdeframond"
captured_at: "2026-09-27T17:33:15Z"
capture_tool: "hn-digest"
hn_id: 49868243
score: 2
comments: 0
posted_at: "2026-09-27T16:33:49Z"
tags:
  - hacker-news
---

# Observability-driven harnesses for building with agents

- HN: [49868243](https://news.ycombinator.com/item?id=49868243)
- Source: [www.datadoghq.com](https://www.datadoghq.com/blog/ai/harness-first-agents/)
- Score: 2
- Comments: 0
- Posted: 2026-09-27T16:33:49Z

## Translation

Title: Observability-driven harnesses for building with agents
Article title: Closing the verification loop: Observability-driven harnesses for building with agents | Datadog
Description: Learn how Datadog verifies AI-generated systems at scale using deterministic testing, formal methods, and observability-driven feedback loops.

Article text:
Closing the verification loop: Observability-driven harnesses for building with agents | Datadog
Datadog named a Leader in the Gartner® Magic Quadrant™ for Observability Platforms
Leader in the Gartner® Magic Quadrant™
30;
},
handleResize() {
if (window.innerWidth >= 1024) {
this.mobileOpen = false;
this.dropdownOpen = 'none';
}
},
checkAnnouncementBanner() {
const announcementBanner = document.querySelector('.announcement-banner') || document.querySelector('.announcement-banner--large');
if (announcementBanner) {
this.hasAnnouncementBanner = true;
} else {
this.hasAnnouncementBanner = false;
}
}
}" x-init="checkAnnouncementBanner()" x-on:scroll.window="handleScroll" x-on:resize.window="handleResize"> Product {
this.openCategory = category;
const productMenu = document.querySelector('.product-menu');
window.DD_RUM?.onReady(function() {
if (productMenu.classList.contains('show')) {
window.DD_RUM.addAction(`Product Category ${category} Hover`)
}
})
}, 160);
},
clearCategory() {
clearTimeout(this.timeoutID);
}
}" x-init="
const menu = document.querySelector('.product-menu');
var observer = new MutationObserver(function(mutations) {
mutations.forEach(function(mutation) {
if (mutation.attributeName === 'class' && !mutation.target.classList.contains('show')) {
openCategory = 'observability';
}
});
});
observer.observe(menu, { attributes: true });
"> The integrated platform for monitoring & security
dashboard
Platform Capabilities
End-to-end, simplified visibility into your stack’s health & performance
Application Performance Monitoring
Bring Your Own Cloud Log Management
Outpace AI-powered attacks with unified security and observability
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
Optimize front-end performance and enhance user experiences
Build, test, secure and ship quality code faster
Internal Developer Portal (IDP)
Integrated, streamlined workflows for faster time-to-resolution
Internal Developer Portal (IDP)
Monitor and improve model performance. Pinpoint root causes and detect anomalies
Built-in features & integrations that power the Datadog platform
Amazon Web Services Monitoring
Product
host-map
Infrastructure
Application Performance Monitoring
Bring Your Own Cloud Log Management
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
Internal Developer Portal (IDP)
Internal Developer Portal (IDP)
dashboard
Platform Capabilities
Amazon Web Services Monitoring
AI Closing the verification loop: Observability-driven harnesses for building with agents
AI agents can now produce software faster than any team can verify it. The bottleneck has moved from writing code to trusting what was written.
We have seen this pattern before. Early programmers resisted compilers because they could write better assembly by hand. Often they were right. Compilers earned trust because the languages they translate have precise semantics: The programmer defines what the program does; the compiler has freedom over how it is implemented. Automation has consistently won only when paired with verification.
With AI agents, building trust is more challenging than in the case of compilers. AI agents ingest unrestricted natural language, sometimes from untrusted sources, and translate it into running code. We must find new ways to verify the outputs of these new program synthesis engines.
At Datadog, we see this as our opportunity: preventing “vibe-coding” from spiraling into “yolo-deploys.” Our approach is harness-first engineering : instead of reading every line of agent-generated code, invest in automated checks that can tell us with high confidence, in seconds, whether the code is correct. The agent generates code, the harness verifies it, production telemetry validates it, and if something is wrong, the feedback updates the harness and the agent tries again. The specific methods to develop harnesses vary in rigor—deterministic simulation testing, formal specifications, shadow evaluation, observability-driven feedback loops—but the principle remains the same: make the verification fast and automatic, and let the harness do the work that human review cannot scale to do.
We have been building toward this vision for the past year. BitsEvolve , our LLM-guided evolutionary optimizer, uses production-driven feedback loops to keep evolved code honest. It shipped 10x speedups on key ingestion functions, 1.53x on a DeBERTa encoder for sensitive data scanning , and 1.57x on Toto , our timeseries forecasting model—all verified against live traffic. We learned that if the harness is tight enough, the LLM can explore freely and the results hold. A good harness makes iteration cheap. A weak harness cannot be compensated for by better models or more human review.
Then in late 2025, we observed a sharp jump in model capabilities. Until then, BitsEvolve operated at file and function level. We began asking what would happen if we pushed the harness-first approach to full systems. This post is a first in a series in which we describe how we evolved harness-first engineering to operate at system scale.
In this first post, we walk through two projects: redis-rust , where we learned the methodology through trial and error, and Helix , a Kafka-compatible streaming engine where we refined it. In both cases, the harness proved strong enough to replace code review as the primary source of correctness: redis-rust reached production-like staging with comparable latency and an 87% memory reduction after agent-guided iteration, while Helix sustained millions of deterministic simulation runs and achieved about 93% of peak disk throughput, without sacrificing Kafka-semantic guarantees. The pattern was consistent: once invariants were explicit and continuously checked, the agent could safely move faster than humans could review.
redis-rust: Learning the methodology
redis-rust was our first attempt at pushing a coding agent to build a full system. We ran it with a single agent (Claude Code with Opus 4.5) and learned primarily by discovering issues as the codebase evolved.
Within a few hours of back-and-forth on architectural ideas—actor-per-shard design, conflict-free replicated data types (CRDTs)—the agent produced a working Redis-compatible server. It compiled and passed tests, but many details were subtly wrong in ways we initially lacked the infrastructure to detect. Error messages drifted from Redis compatibility in ways that seemed reasonable but were incorrect. The agent also had a tendency to over-engineer abstractions that we later simplified.
So we started building verification steps one layer at a time. Each layer was motivated by something that slipped through the layer below it.
We began with a shadow-state oracle: a simple HashMap running alongside the real executor that compared responses after every operation. That caught basic semantic bugs but could not exercise timing-dependent paths, so we added deterministic simulation testing ( DST ) with fault injection.
DST required invariants to check against, which led us to write TLA+ specifications for the replication and gossip protocols. For the CRDT merge properties, we needed mathematical guarantees, so we added Kani, a Rust verification tool, for bounded proofs. For system-level correctness, we ran Maelstrom, a distributed systems testing framework, with the Knossos linearizability checker at 1, 3, and 5 nodes. We also ran the official Redis Tcl compatibility suite for the implemented commands. Design decisions, trade-offs, and verification methods are documented in the technical report .
The next step was empirical verification using real traffic. We used Ephemera, our internal caching system that operates a cluster of Redis shards behind a data/control plane API. For an apples-to-apples comparison, we set up a redis-rust shadow cluster that received the same workload as its Redis 8.4 counterpart.
Using pup as a Datadog interface, the agent verified redis-rust was functional in a staging environment with nominal latency differences:
redis-rust’s latency profile was initially comparable to Redis. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> redis-rust’s latency profile was initially comparable to Redis. Close dialog
However, it used 8x more memory than Redis 8.4 as it was hardcoded to pre-allocate 512 x 8 KB buffers (4 MB) at startup optimized for an exhaustive micro-benchmark, among other things. Within minutes, the agent suggested and implemented three optimizations for an 87% reduction in memory footprint:
The agent’s memory footprint optimization with metrics as feedback. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> The agent’s memory footprint optimization with metrics as feedback. Close dialog
We’ll continue tuning performance using metrics for memory, network, and latency along with CPU profiles.
Helix: Improving in the harness
Helix , a Kafka-like streaming service on object storage, is another full system we built using coding agents based on what we learned from redis-rust. This time, we ran with multiple coding agents, primarily Claude Code and Codex. The workflow was constraint-first: design artifacts are the contracts, semantics are stated explicitly (bringing our experience operating Kafka for over a decade), and every artifact is coupled to a feedback mechanism through a verification pyramid that can falsify mistakes.
Each layer trades off speed against rigor:
Shared invariants flow from TLA+ specifications into Stateright, DST, Kani, and staging telemetry, with DST as the primary verification layer. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Shared invariants flow from TLA+ specifications into Stateright, DST, Kani, and staging telemetry, with DST as the primary verification layer. Close dialog
Contracts before code
We described core invariants up front—replicated log plus object storage, partition model, failure boundaries, Kafka compatibility—then had the agent design each subsystem independently: Raft (verified with TLA+), write-ahead log (WAL), tiering, DST, service layer, Kafka wire protocol. The agent is not allowed to invent system meaning. What is durable vs. acknowledged? What is committed vs. visible? What happens on crash at each boundary?
Antithesis and AWS call this “semi-formal methods”: specifications and invariants explicit enough to be checked, and cheap enough to run continuously. The mental overhead is roughly 2–3x the effort of writing the code itself. It pays back immediately—explicit invariants turn every agent iteration into an objective pass/fail decision instead of a judgment call.
DST is the workhorse. Popularized by FoundationDB and TigerBeetle , DST abstracts physical time, makes execution deterministic, and injects faults synthetically. Each run takes about 5 seconds and exercises actual production code through randomized scenarios with fault injection.
TLA+ specs provide the map. They define the state variables, actions, and invariants that would otherwise take hours to extract from implementation code. We generate them from architecture decision records (ADRs), catching ambiguities early. Model checking and Kani escalate when stronger guarantees are needed. Telemetry grounds everything empirically. The lightest mechanism that can falsify a hypothesis is used first.
Architecture decision records (ADRs) generate TLA+ specifications, which define invariants reused across Stateright, DST, Kani, and staging telemetry to keep verification layers aligned. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Architecture decision records (ADRs) generate TLA+ specifications, which define

[truncated]

## Original Extract

Learn how Datadog verifies AI-generated systems at scale using deterministic testing, formal methods, and observability-driven feedback loops.

Closing the verification loop: Observability-driven harnesses for building with agents | Datadog
Datadog named a Leader in the Gartner® Magic Quadrant™ for Observability Platforms
Leader in the Gartner® Magic Quadrant™
30;
},
handleResize() {
if (window.innerWidth >= 1024) {
this.mobileOpen = false;
this.dropdownOpen = 'none';
}
},
checkAnnouncementBanner() {
const announcementBanner = document.querySelector('.announcement-banner') || document.querySelector('.announcement-banner--large');
if (announcementBanner) {
this.hasAnnouncementBanner = true;
} else {
this.hasAnnouncementBanner = false;
}
}
}" x-init="checkAnnouncementBanner()" x-on:scroll.window="handleScroll" x-on:resize.window="handleResize"> Product {
this.openCategory = category;
const productMenu = document.querySelector('.product-menu');
window.DD_RUM?.onReady(function() {
if (productMenu.classList.contains('show')) {
window.DD_RUM.addAction(`Product Category ${category} Hover`)
}
})
}, 160);
},
clearCategory() {
clearTimeout(this.timeoutID);
}
}" x-init="
const menu = document.querySelector('.product-menu');
var observer = new MutationObserver(function(mutations) {
mutations.forEach(function(mutation) {
if (mutation.attributeName === 'class' && !mutation.target.classList.contains('show')) {
openCategory = 'observability';
}
});
});
observer.observe(menu, { attributes: true });
"> The integrated platform for monitoring & security
dashboard
Platform Capabilities
End-to-end, simplified visibility into your stack’s health & performance
Application Performance Monitoring
Bring Your Own Cloud Log Management
Outpace AI-powered attacks with unified security and observability
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
Optimize front-end performance and enhance user experiences
Build, test, secure and ship quality code faster
Internal Developer Portal (IDP)
Integrated, streamlined workflows for faster time-to-resolution
Internal Developer Portal (IDP)
Monitor and improve model performance. Pinpoint root causes and detect anomalies
Built-in features & integrations that power the Datadog platform
Amazon Web Services Monitoring
Product
host-map
Infrastructure
Application Performance Monitoring
Bring Your Own Cloud Log Management
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
Internal Developer Portal (IDP)
Internal Developer Portal (IDP)
dashboard
Platform Capabilities
Amazon Web Services Monitoring
AI Closing the verification loop: Observability-driven harnesses for building with agents
AI agents can now produce software faster than any team can verify it. The bottleneck has moved from writing code to trusting what was written.
We have seen this pattern before. Early programmers resisted compilers because they could write better assembly by hand. Often they were right. Compilers earned trust because the languages they translate have precise semantics: The programmer defines what the program does; the compiler has freedom over how it is implemented. Automation has consistently won only when paired with verification.
With AI agents, building trust is more challenging than in the case of compilers. AI agents ingest unrestricted natural language, sometimes from untrusted sources, and translate it into running code. We must find new ways to verify the outputs of these new program synthesis engines.
At Datadog, we see this as our opportunity: preventing “vibe-coding” from spiraling into “yolo-deploys.” Our approach is harness-first engineering : instead of reading every line of agent-generated code, invest in automated checks that can tell us with high confidence, in seconds, whether the code is correct. The agent generates code, the harness verifies it, production telemetry validates it, and if something is wrong, the feedback updates the harness and the agent tries again. The specific methods to develop harnesses vary in rigor—deterministic simulation testing, formal specifications, shadow evaluation, observability-driven feedback loops—but the principle remains the same: make the verification fast and automatic, and let the harness do the work that human review cannot scale to do.
We have been building toward this vision for the past year. BitsEvolve , our LLM-guided evolutionary optimizer, uses production-driven feedback loops to keep evolved code honest. It shipped 10x speedups on key ingestion functions, 1.53x on a DeBERTa encoder for sensitive data scanning , and 1.57x on Toto , our timeseries forecasting model—all verified against live traffic. We learned that if the harness is tight enough, the LLM can explore freely and the results hold. A good harness makes iteration cheap. A weak harness cannot be compensated for by better models or more human review.
Then in late 2025, we observed a sharp jump in model capabilities. Until then, BitsEvolve operated at file and function level. We began asking what would happen if we pushed the harness-first approach to full systems. This post is a first in a series in which we describe how we evolved harness-first engineering to operate at system scale.
In this first post, we walk through two projects: redis-rust , where we learned the methodology through trial and error, and Helix , a Kafka-compatible streaming engine where we refined it. In both cases, the harness proved strong enough to replace code review as the primary source of correctness: redis-rust reached production-like staging with comparable latency and an 87% memory reduction after agent-guided iteration, while Helix sustained millions of deterministic simulation runs and achieved about 93% of peak disk throughput, without sacrificing Kafka-semantic guarantees. The pattern was consistent: once invariants were explicit and continuously checked, the agent could safely move faster than humans could review.
redis-rust: Learning the methodology
redis-rust was our first attempt at pushing a coding agent to build a full system. We ran it with a single agent (Claude Code with Opus 4.5) and learned primarily by discovering issues as the codebase evolved.
Within a few hours of back-and-forth on architectural ideas—actor-per-shard design, conflict-free replicated data types (CRDTs)—the agent produced a working Redis-compatible server. It compiled and passed tests, but many details were subtly wrong in ways we initially lacked the infrastructure to detect. Error messages drifted from Redis compatibility in ways that seemed reasonable but were incorrect. The agent also had a tendency to over-engineer abstractions that we later simplified.
So we started building verification steps one layer at a time. Each layer was motivated by something that slipped through the layer below it.
We began with a shadow-state oracle: a simple HashMap running alongside the real executor that compared responses after every operation. That caught basic semantic bugs but could not exercise timing-dependent paths, so we added deterministic simulation testing ( DST ) with fault injection.
DST required invariants to check against, which led us to write TLA+ specifications for the replication and gossip protocols. For the CRDT merge properties, we needed mathematical guarantees, so we added Kani, a Rust verification tool, for bounded proofs. For system-level correctness, we ran Maelstrom, a distributed systems testing framework, with the Knossos linearizability checker at 1, 3, and 5 nodes. We also ran the official Redis Tcl compatibility suite for the implemented commands. Design decisions, trade-offs, and verification methods are documented in the technical report .
The next step was empirical verification using real traffic. We used Ephemera, our internal caching system that operates a cluster of Redis shards behind a data/control plane API. For an apples-to-apples comparison, we set up a redis-rust shadow cluster that received the same workload as its Redis 8.4 counterpart.
Using pup as a Datadog interface, the agent verified redis-rust was functional in a staging environment with nominal latency differences:
redis-rust’s latency profile was initially comparable to Redis. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> redis-rust’s latency profile was initially comparable to Redis. Close dialog
However, it used 8x more memory than Redis 8.4 as it was hardcoded to pre-allocate 512 x 8 KB buffers (4 MB) at startup optimized for an exhaustive micro-benchmark, among other things. Within minutes, the agent suggested and implemented three optimizations for an 87% reduction in memory footprint:
The agent’s memory footprint optimization with metrics as feedback. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> The agent’s memory footprint optimization with metrics as feedback. Close dialog
We’ll continue tuning performance using metrics for memory, network, and latency along with CPU profiles.
Helix: Improving in the harness
Helix , a Kafka-like streaming service on object storage, is another full system we built using coding agents based on what we learned from redis-rust. This time, we ran with multiple coding agents, primarily Claude Code and Codex. The workflow was constraint-first: design artifacts are the contracts, semantics are stated explicitly (bringing our experience operating Kafka for over a decade), and every artifact is coupled to a feedback mechanism through a verification pyramid that can falsify mistakes.
Each layer trades off speed against rigor:
Shared invariants flow from TLA+ specifications into Stateright, DST, Kani, and staging telemetry, with DST as the primary verification layer. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Shared invariants flow from TLA+ specifications into Stateright, DST, Kani, and staging telemetry, with DST as the primary verification layer. Close dialog
Contracts before code
We described core invariants up front—replicated log plus object storage, partition model, failure boundaries, Kafka compatibility—then had the agent design each subsystem independently: Raft (verified with TLA+), write-ahead log (WAL), tiering, DST, service layer, Kafka wire protocol. The agent is not allowed to invent system meaning. What is durable vs. acknowledged? What is committed vs. visible? What happens on crash at each boundary?
Antithesis and AWS call this “semi-formal methods”: specifications and invariants explicit enough to be checked, and cheap enough to run continuously. The mental overhead is roughly 2–3x the effort of writing the code itself. It pays back immediately—explicit invariants turn every agent iteration into an objective pass/fail decision instead of a judgment call.
DST is the workhorse. Popularized by FoundationDB and TigerBeetle , DST abstracts physical time, makes execution deterministic, and injects faults synthetically. Each run takes about 5 seconds and exercises actual production code through randomized scenarios with fault injection.
TLA+ specs provide the map. They define the state variables, actions, and invariants that would otherwise take hours to extract from implementation code. We generate them from architecture decision records (ADRs), catching ambiguities early. Model checking and Kani escalate when stronger guarantees are needed. Telemetry grounds everything empirically. The lightest mechanism that can falsify a hypothesis is used first.
Architecture decision records (ADRs) generate TLA+ specifications, which define invariants reused across Stateright, DST, Kani, and staging telemetry to keep verification layers aligned. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Architecture decision records (ADRs) generate TLA+ specifications, which define

[truncated]
