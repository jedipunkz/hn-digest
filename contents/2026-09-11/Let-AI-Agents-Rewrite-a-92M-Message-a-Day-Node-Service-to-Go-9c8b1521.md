---
source: "https://www.checklyhq.com/blog/agentic-rewrite-nodejs-to-go/"
hn_url: "https://news.ycombinator.com/item?id=49662139"
title: "Let AI Agents Rewrite a 92M-Message-a-Day Node Service to Go"
article_title: "Rewriting a Node.js Service in Go With AI Agents"
image: "https://images.prismic.io/checklyhq/DkQueTTRxgYbz44l_OG.png?auto=format,compress"
author: "tnolet"
captured_at: "2026-09-11T17:34:32Z"
capture_tool: "hn-digest"
hn_id: 49662139
score: 1
comments: 0
posted_at: "2026-09-11T17:31:36Z"
tags:
  - hacker-news
---

# Let AI Agents Rewrite a 92M-Message-a-Day Node Service to Go

- HN: [49662139](https://news.ycombinator.com/item?id=49662139)
- Source: [www.checklyhq.com](https://www.checklyhq.com/blog/agentic-rewrite-nodejs-to-go/)
- Score: 1
- Comments: 0
- Posted: 2026-09-11T17:31:36Z

## Translation

Title: Let AI Agents Rewrite a 92M-Message-a-Day Node Service to Go
Article title: Rewriting a Node.js Service in Go With AI Agents
Description: How Checkly rewrote a Node.js service handling 92M messages a day into Go using AI agents, and the test harness that made it safe to ship.

Article text:
Rewriting a Node.js Service in Go With AI Agents Checkly - Home Product Detect
Uptime Monitoring Measure the availability of your digital footprint Heartbeat Monitoring Catch cron jobs and backups that fail silently Synthetic Monitoring Simulate real user interactions across your stack Testing Catch issues before production with an AI-powered test reporter Communicate
Status Pages Communicate app availability to your customers Alerts Contextual alerting to notify the team right away Resolve
AI Root Cause Analysis Automated root cause analysis powered by AI agents Traces Powerful OTel tracing for deeper insights Getting Started
Developers Checks in TypeScript, in your repo, deployed with your app SRE & Platform The synthetic layer for your observability stack QA Engineers Run your Playwright suite as production monitors Engineering Managers Give every team ownership of its own monitors Use Cases
Critical user flows Watch login, checkout, and signup from real browsers API & backend monitoring Validate endpoints, chains, and auth on a schedule Tests to production monitors Promote tagged specs from CI to scheduled checks Observability consolidation Replace the synthetic module bolted onto your APM Reliability for AI-generated code Verify what your agents ship before your users do Industries
E-commerce Protect checkout and revenue paths around the clock Financial services Prove uptime and latency against strict SLAs SaaS & B2B software Keep every tenant's core workflows verified Moving from another tool?
Documentation Learn key concepts and features API Docs Build on the Checkly REST API CLI Docs Bring monitoring to your terminal Quickstart Set up your first check fast Guides In-depth Playwright & OTel guides MCP Server Connect Checkly to your AI tools Reference
Agent Skills Terraform Pulumi All Integrations Changelog Resources Featured
Blog Read about the latest news at Checkly Learn Tips and best practices for learning Playwright and more Webinars Register or view webinars on-demand Events Meet the team in person Community
Public Roadmap See Checkly's active customer outcomes and product direction Community Slack Connect with the Checkly Community Customers Pricing Login Start for free Open Navigation Product Solutions Developers Resources Customers Pricing Start for free Login Blog / AI / Development / Playwright We Let AI Agents Rewrite a 92M-Message-a-Day Service in Go. Zero Incidents.
Our Results Daemon processes about 92 million messages a day. We recently rewrote it from Node.js to Go, and we let Claude Code write it.
We wanted to know whether we could trust an agentic rewrite for a critical, high-throughput production service rather than a prototype. It shipped with zero incidents, a 70% reduction in running pods, and a lighter database load. Go's stronger type system also proved a better fit for agents than JavaScript, adding protection against regressions and letting us ship faster and with more confidence.
What made it work was the test harness we built before the agent wrote a line. Here's how we designed it, and the principles you can reuse on your own legacy services.
Checkly is a monitoring platform that runs synthetic checks, automated scripts that emulate real users, and uptime checks that confirm a system component is operational. A runner component executes all of these and produces a result that has to be processed, stored, and alerted on.
Since we introduced uptime checks, and with the company's overall growth, the volume of checks run on our platform has doubled over the last year. Some components started degrading under that load. The most notable one was Results Daemon, a Node.js component written in vanilla JavaScript.
Results Daemon is a background worker. It consumes results from our runner, writes them to databases, determines the check outcome, issues alerts, and schedules retries as needed. It also publishes WebSocket updates to our CLI and UI. In total, this component processes approximately 92,000,000 messages every day , around 40,000,000 of them check results and the rest WebSocket publishes.
At that scale, it was becoming a bottleneck. It paged our on-call engineers more often, and limited type safety made every change harder to land safely. So we decided to rewrite Results Daemon in Go using agentic engineering.
We built the harness before we started the rewrite. If an agent is going to write the code, something other than a human reviewer has to define what correct means.
We built it on these design principles:
The harness tests the component as a black box. There is zero coupling between the code or language of the system under test and the harness itself.
Every test case provides an input and expects a deterministic output, with all outputs recorded in "golden files." These are generated against the legacy system and later used by the rewrite to assert byte-to-byte parity.
Non-deterministic fields, such as UUIDs or timestamps generated during the test itself, are written as <uuid> or <timestamp> and are still type-checked, to minimize the risk of differences in behavior slipping through.
Surrounding components (databases, queues, caches, other services) are categorized as boundaries. These are managed strictly by the harness, and the system under test is only pointed at them using environment variables.
Boundaries use real instances of the service in testing. If data is written to PostgreSQL, the harness uses a real PostgreSQL container rather than an emulated one.
For simpler boundaries such as SQS queues, we built our own emulator instead of using ElasticMQ or LocalStack. We found it more performant and simpler, both in our tests and in our assertions.
A set of "oracle" classes fetches the test output and asserts whether the test failed. For example, PostgresOracle.expectResultToMatchSnapshot(testId) fetches the relevant output and asserts its byte-level accuracy against the established golden file.
Three technical choices carried the harness:
Playwright , a testing framework built for reliability, with strong tooling for network interception and parallel test execution. Our own synthetic monitoring offering is built on Playwright, so we already knew it well, and its black-box model fit what we were doing here.
Docker Compose , the simplest way to start and tear down containers for our boundaries, both locally and in CI.
Toxiproxy , a TCP proxy that emulates network conditions and let us test how the system behaves when surrounding infrastructure fails.
With the architecture settled, the next step was building actual test cases. Results Daemon's outputs depend on two things:
The check result, the data object representing the outcome of a check execution. It holds the outcome state of the check run (Failed | Degraded | Success) and other metadata.
The check configuration at the time the result was received: retry rules for rescheduling, alert rules for notifications, and so on.
From there it followed quickly that behavior coverage depends directly on the diversity of the inputs. A harness that covers every combination of check result and configuration covers every possible code path, with zero coupling to the implementation. That gave us our first principle:
The quality of the harness depends on the quality of the inputs you can provide. The more diverse and realistic the inputs, the greater the coverage of code paths and behaviors. In our case, the number of outcomes is represented by (accountConfigs × groupConfigs × checkConfigs × resultOutcomes) .
We generated those inputs from our internal data lake. We extracted all account, group, and check configurations along with every result outcome from the last 24 hours, which is the longest interval we schedule checks at. We loaded all of it into a ClickHouse instance, and each dataset was "collapsed" into a unique set of configurations and outcomes, each tagged with its number of occurrences. We then used that data to seed realistic scenarios and pin them to business rules. For example, a re-dispatch takes its runtime from the account when the job pins none .
After generating the test cases, we used code coverage reports to estimate how effective the harness was, targeting between 90% and 100%. We also fed those reports back to the agent to review, identify gaps, and propose test cases for anything missed. We expected some gaps to remain, but with all relevant files sufficiently covered and every scoped feature included, we considered the harness ready for a test run.
Code coverage is a great metric to track as you start out with your initial set of test cases, since high code coverage means the core behaviors are covered. However, it does measure direct system actual behavior and is therefore likely to miss edge cases.
We also built a separate suite of tests that caught failure modes when infrastructure failures occur, e.g., PostgreSQL going down, using Toxiproxy. These were focused on how behavior changes and what data is lost when a piece of infrastructure goes down.
A prerequisite for this approach was an earlier migration to a monorepo running on Tilt , which lets us spin up the whole platform end to end in a local development environment.
We dispatched an instance of Claude Code, using Fable , and gave it one instruction: "build a Go service that consumes data from input queues, processes the messages, and writes the outputs to downstream applications such as databases, caches, and other queues," with the legacy implementation available as a reference. The main acceptance criterion was that the test harness passed against the new implementation.
The agent ran overnight and produced a deployable service of about 13,000 lines of application code, architecturally mirroring the legacy implementation. It also kept token usage within the daily limits of a $200 subscription.
We ran an earlier attempt with the same instructions using Opus . That implementation did not meet our bar and was discarded.
After reviewing the implementation, we deployed the application to every environment except production, using an internal account to push results so we could get feedback quickly. The review also surfaced anti-patterns we wanted gone and improvements we wanted in:
Removing configuration generated at runtime. The legacy application derived its configuration from partial values provided via environment variables. That has been a pain point in the past when debugging and making configuration changes, especially under pressure during incidents.
Improving end-to-end observability. The legacy service had high-level observability, but given the increase in throughput, there was clear room to do better.
We implemented both with human supervision, which gave us two more principles:
Human intervention is sometimes required, especially to find and fix anti-patterns an agent inherited from the legacy application or introduced itself. For us this meant removing all configuration derivations and making static environment variables the only way the application is configured, plus improving observability. That improved the system's operability and simplified the code in both the rewrite and the harness.
Take the rewrite as an opportunity to improve things. In our case that meant expanding end-to-end observability. At this message volume, recording low-level metrics such as database connection utilization, CPU, memory, and timings of individual code execution steps was a necessary addition for long-term operational excellence.
With those in place, the next step was production.
First deployment: what we got wrong
Before deploying, we had to decide how to migrate customers. The simplest approach was to spin up separate infrastructure (queues) and patch the consumers to push data to the new queue for accounts with a specific feature flag enabled.
Once the infrastructure was ready, we deployed the Go application across all environments and migrated internal accounts. We ran into issues almost immediately, mainly re

[truncated]

## Original Extract

How Checkly rewrote a Node.js service handling 92M messages a day into Go using AI agents, and the test harness that made it safe to ship.

Rewriting a Node.js Service in Go With AI Agents Checkly - Home Product Detect
Uptime Monitoring Measure the availability of your digital footprint Heartbeat Monitoring Catch cron jobs and backups that fail silently Synthetic Monitoring Simulate real user interactions across your stack Testing Catch issues before production with an AI-powered test reporter Communicate
Status Pages Communicate app availability to your customers Alerts Contextual alerting to notify the team right away Resolve
AI Root Cause Analysis Automated root cause analysis powered by AI agents Traces Powerful OTel tracing for deeper insights Getting Started
Developers Checks in TypeScript, in your repo, deployed with your app SRE & Platform The synthetic layer for your observability stack QA Engineers Run your Playwright suite as production monitors Engineering Managers Give every team ownership of its own monitors Use Cases
Critical user flows Watch login, checkout, and signup from real browsers API & backend monitoring Validate endpoints, chains, and auth on a schedule Tests to production monitors Promote tagged specs from CI to scheduled checks Observability consolidation Replace the synthetic module bolted onto your APM Reliability for AI-generated code Verify what your agents ship before your users do Industries
E-commerce Protect checkout and revenue paths around the clock Financial services Prove uptime and latency against strict SLAs SaaS & B2B software Keep every tenant's core workflows verified Moving from another tool?
Documentation Learn key concepts and features API Docs Build on the Checkly REST API CLI Docs Bring monitoring to your terminal Quickstart Set up your first check fast Guides In-depth Playwright & OTel guides MCP Server Connect Checkly to your AI tools Reference
Agent Skills Terraform Pulumi All Integrations Changelog Resources Featured
Blog Read about the latest news at Checkly Learn Tips and best practices for learning Playwright and more Webinars Register or view webinars on-demand Events Meet the team in person Community
Public Roadmap See Checkly's active customer outcomes and product direction Community Slack Connect with the Checkly Community Customers Pricing Login Start for free Open Navigation Product Solutions Developers Resources Customers Pricing Start for free Login Blog / AI / Development / Playwright We Let AI Agents Rewrite a 92M-Message-a-Day Service in Go. Zero Incidents.
Our Results Daemon processes about 92 million messages a day. We recently rewrote it from Node.js to Go, and we let Claude Code write it.
We wanted to know whether we could trust an agentic rewrite for a critical, high-throughput production service rather than a prototype. It shipped with zero incidents, a 70% reduction in running pods, and a lighter database load. Go's stronger type system also proved a better fit for agents than JavaScript, adding protection against regressions and letting us ship faster and with more confidence.
What made it work was the test harness we built before the agent wrote a line. Here's how we designed it, and the principles you can reuse on your own legacy services.
Checkly is a monitoring platform that runs synthetic checks, automated scripts that emulate real users, and uptime checks that confirm a system component is operational. A runner component executes all of these and produces a result that has to be processed, stored, and alerted on.
Since we introduced uptime checks, and with the company's overall growth, the volume of checks run on our platform has doubled over the last year. Some components started degrading under that load. The most notable one was Results Daemon, a Node.js component written in vanilla JavaScript.
Results Daemon is a background worker. It consumes results from our runner, writes them to databases, determines the check outcome, issues alerts, and schedules retries as needed. It also publishes WebSocket updates to our CLI and UI. In total, this component processes approximately 92,000,000 messages every day , around 40,000,000 of them check results and the rest WebSocket publishes.
At that scale, it was becoming a bottleneck. It paged our on-call engineers more often, and limited type safety made every change harder to land safely. So we decided to rewrite Results Daemon in Go using agentic engineering.
We built the harness before we started the rewrite. If an agent is going to write the code, something other than a human reviewer has to define what correct means.
We built it on these design principles:
The harness tests the component as a black box. There is zero coupling between the code or language of the system under test and the harness itself.
Every test case provides an input and expects a deterministic output, with all outputs recorded in "golden files." These are generated against the legacy system and later used by the rewrite to assert byte-to-byte parity.
Non-deterministic fields, such as UUIDs or timestamps generated during the test itself, are written as <uuid> or <timestamp> and are still type-checked, to minimize the risk of differences in behavior slipping through.
Surrounding components (databases, queues, caches, other services) are categorized as boundaries. These are managed strictly by the harness, and the system under test is only pointed at them using environment variables.
Boundaries use real instances of the service in testing. If data is written to PostgreSQL, the harness uses a real PostgreSQL container rather than an emulated one.
For simpler boundaries such as SQS queues, we built our own emulator instead of using ElasticMQ or LocalStack. We found it more performant and simpler, both in our tests and in our assertions.
A set of "oracle" classes fetches the test output and asserts whether the test failed. For example, PostgresOracle.expectResultToMatchSnapshot(testId) fetches the relevant output and asserts its byte-level accuracy against the established golden file.
Three technical choices carried the harness:
Playwright , a testing framework built for reliability, with strong tooling for network interception and parallel test execution. Our own synthetic monitoring offering is built on Playwright, so we already knew it well, and its black-box model fit what we were doing here.
Docker Compose , the simplest way to start and tear down containers for our boundaries, both locally and in CI.
Toxiproxy , a TCP proxy that emulates network conditions and let us test how the system behaves when surrounding infrastructure fails.
With the architecture settled, the next step was building actual test cases. Results Daemon's outputs depend on two things:
The check result, the data object representing the outcome of a check execution. It holds the outcome state of the check run (Failed | Degraded | Success) and other metadata.
The check configuration at the time the result was received: retry rules for rescheduling, alert rules for notifications, and so on.
From there it followed quickly that behavior coverage depends directly on the diversity of the inputs. A harness that covers every combination of check result and configuration covers every possible code path, with zero coupling to the implementation. That gave us our first principle:
The quality of the harness depends on the quality of the inputs you can provide. The more diverse and realistic the inputs, the greater the coverage of code paths and behaviors. In our case, the number of outcomes is represented by (accountConfigs × groupConfigs × checkConfigs × resultOutcomes) .
We generated those inputs from our internal data lake. We extracted all account, group, and check configurations along with every result outcome from the last 24 hours, which is the longest interval we schedule checks at. We loaded all of it into a ClickHouse instance, and each dataset was "collapsed" into a unique set of configurations and outcomes, each tagged with its number of occurrences. We then used that data to seed realistic scenarios and pin them to business rules. For example, a re-dispatch takes its runtime from the account when the job pins none .
After generating the test cases, we used code coverage reports to estimate how effective the harness was, targeting between 90% and 100%. We also fed those reports back to the agent to review, identify gaps, and propose test cases for anything missed. We expected some gaps to remain, but with all relevant files sufficiently covered and every scoped feature included, we considered the harness ready for a test run.
Code coverage is a great metric to track as you start out with your initial set of test cases, since high code coverage means the core behaviors are covered. However, it does measure direct system actual behavior and is therefore likely to miss edge cases.
We also built a separate suite of tests that caught failure modes when infrastructure failures occur, e.g., PostgreSQL going down, using Toxiproxy. These were focused on how behavior changes and what data is lost when a piece of infrastructure goes down.
A prerequisite for this approach was an earlier migration to a monorepo running on Tilt , which lets us spin up the whole platform end to end in a local development environment.
We dispatched an instance of Claude Code, using Fable , and gave it one instruction: "build a Go service that consumes data from input queues, processes the messages, and writes the outputs to downstream applications such as databases, caches, and other queues," with the legacy implementation available as a reference. The main acceptance criterion was that the test harness passed against the new implementation.
The agent ran overnight and produced a deployable service of about 13,000 lines of application code, architecturally mirroring the legacy implementation. It also kept token usage within the daily limits of a $200 subscription.
We ran an earlier attempt with the same instructions using Opus . That implementation did not meet our bar and was discarded.
After reviewing the implementation, we deployed the application to every environment except production, using an internal account to push results so we could get feedback quickly. The review also surfaced anti-patterns we wanted gone and improvements we wanted in:
Removing configuration generated at runtime. The legacy application derived its configuration from partial values provided via environment variables. That has been a pain point in the past when debugging and making configuration changes, especially under pressure during incidents.
Improving end-to-end observability. The legacy service had high-level observability, but given the increase in throughput, there was clear room to do better.
We implemented both with human supervision, which gave us two more principles:
Human intervention is sometimes required, especially to find and fix anti-patterns an agent inherited from the legacy application or introduced itself. For us this meant removing all configuration derivations and making static environment variables the only way the application is configured, plus improving observability. That improved the system's operability and simplified the code in both the rewrite and the harness.
Take the rewrite as an opportunity to improve things. In our case that meant expanding end-to-end observability. At this message volume, recording low-level metrics such as database connection utilization, CPU, memory, and timings of individual code execution steps was a necessary addition for long-term operational excellence.
With those in place, the next step was production.
First deployment: what we got wrong
Before deploying, we had to decide how to migrate customers. The simplest approach was to spin up separate infrastructure (queues) and patch the consumers to push data to the new queue for accounts with a specific feature flag enabled.
Once the infrastructure was ready, we deployed the Go application across all environments and migrated internal accounts. We ran into issues almost immediately, mainly re

[truncated]
