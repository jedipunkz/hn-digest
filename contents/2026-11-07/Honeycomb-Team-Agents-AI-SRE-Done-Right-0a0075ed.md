---
source: "https://www.honeycomb.io/blog/honeycomb-team-agents-ai-sre-done-right"
hn_url: "https://news.ycombinator.com/item?id=50405124"
title: "Honeycomb Team Agents: AI SRE Done Right"
article_title: "Honeycomb Team Agents: AI SRE Done Right"
image: "https://cdn.sanity.io/images/927dxq0h/production/143ed4c976a8e648aa2ff34592bdbf54d7a00800-5760x3240.png?rect=0,115,5760,3010&w=1200&h=627&fit=crop"
author: "mooreds"
captured_at: "2026-11-07T15:28:25Z"
capture_tool: "hn-digest"
hn_id: 50405124
score: 1
comments: 0
posted_at: "2026-11-07T14:59:40Z"
tags:
  - hacker-news
---

# Honeycomb Team Agents: AI SRE Done Right

- HN: [50405124](https://news.ycombinator.com/item?id=50405124)
- Source: [www.honeycomb.io](https://www.honeycomb.io/blog/honeycomb-team-agents-ai-sre-done-right)
- Score: 1
- Comments: 0
- Posted: 2026-11-07T14:59:40Z

## Translation

Title: Honeycomb Team Agents: AI SRE Done Right
Description: Introducing Honeycomb Team Agents: AI SRE coworkers that proactively check releases, instrumentation, cost, and agent quality, not just respond to alerts.

Article text:
Honeycomb Team Agents: AI SRE Done Right Going to AWS re:Invent 2026? Join us at booth 1871 for demos, giveaways, and more. Going to AWS re:Invent 2026? Join us at booth 1871 for demos, giveaways, and more. Meet with us Platform Understand Your Systems Investigate with Agents Observe AI Systems Manage Telemetry Operate at Scale Platform Overview Understand and debug complex software and AI systems with full production context. Explore See how your apps, services and experiences connect.
Distributed Tracing Follow requests across services and AI workflows.
Log Analytics Explore structured logs to find issues faster.
Time Series Metrics Track trends and investigate underlying causes.
Frontend Observability Connect user experiences to backend behavior.
Service Map Visualize service dependencies and system health.
BubbleUp Identify what is driving unusual behavior.
Let AI explore production data to resolve issues faster.
Honeycomb Intelligence Investigate and resolve issues with AI.
Canvas Collaborate in an AI-guided investigation workspace.
Team Agents AI SRE coworkers that join your team proactively
Honeycomb MCP Give AI tools access to production context.
Anomaly Detection Surface meaningful changes before users are affected.
Understand agent behavior, performance and cost in production.
AI Ecosystem Track agent behavior, performance and cost across your fleet.
Agent Observability Monitor and debug AI agents end to end.
Collect, control and route telemetry across your stack.
OpenTelemetry Instrument systems with open standards.
Telemetry Pipeline Collect, refine and route telemetry at scale.
Refinery as a Service Control data volume without losing valuable context.
Maintain reliability and efficiency at enterprise scale.
Private Cloud Run Honeycomb in a dedicated environment.
SLOs Set and track customer-focused reliability goals.
App Integrations Connect the tools across your stack.
Honeycomb Services Get expert support to accelerate results.
Partners Discover Honeycomb’s partner ecosystem.
By Use Case By Industry & Size Featured Technologies Explore solutions for specific engineering challenges.
Incident Response Investigate and resolve production issues faster.
LLM Observability Understand LLM quality, latency and cost.
Agent Observability Trace agent decisions, tools and handoffs.
Migrations Maintain visibility throughout your cloud migration.
Frontend Development Connect user experiences to system performance.
The industries and company sizes Honeycomb supports.
Financial Services Operate complex financial systems reliably.
Retail & Ecommerce Protect digital experiences and critical transactions.
Software & Technology Build, operate, and innovate software faster.
Enterprise Scale observability across complex teams and systems.
AI Startups Build and scale reliable AI-native products.
Connect Honeycomb with the technologies across your stack.
OpenTelemetry Send telemetry using open standards.
Amazon Web Services Understand application and infrastructure behavior across AWS.
Microsoft Azure Gain deeper context across your Azure environment.
Kubernetes Understand application and cluster behavior together.
Google Cloud Observe services and workloads across Google Cloud.
AI Agents Connect agent workflows with production telemetry.
All Integrations Explore every Honeycomb integration.
See how Honeycomb improves your business and engineering outcomes.
Predictable Costs Control observability costs as telemetry grows.
DevOps & Releases Ship faster with greater confidence.
Meet Customer SLAs Protect reliability and customer commitments.
See how teams use Honeycomb to achieve measurable results.
Understand how Honeycomb differs from other approaches.
Observability Engineering Download the second edition free today. Explore Why Honeycomb Resources Learn Hub
Blog Articles Read technical insights and product perspectives.
Getting Started Learn the fundamentals of using Honeycomb.
Docs Explore Honeycomb’s technical documentation.
Guides Find practical guidance for common challenges.
Success Stories See how customers achieve real results.
Webinars & Videos Watch expert sessions, product demos & walkthroughs.
Academy Develop skills through structured learning.
Events Discover upcoming Honeycomb events and sessions.
Office Hours Get practical guidance from Honeycomb experts.
Pollinators Slack Connect with observability practitioners on Slack.
Partners Discover Honeycomb’s partner ecosystem.
Support Get help using Honeycomb.
About Us Meet the company behind Honeycomb.
Careers Explore job opportunities at Honeycomb.
News Read company news and announcements.
Log in Sign up Get a demo Platform Solutions Why Honeycomb Resources Pricing Platform
Understand and debug complex software and AI systems with full production context. Understand Your Systems
Honeycomb Team Agents: AI SRE Done Right
Honeycomb Team Agents: AI SRE Done Right
Honeycomb Team Agents are AI SRE coworkers that join your team proactively, not just reactively. Personal agents answer questions about production with evidence; team agents run on goals you define, checking releases, instrumentation, cost, agent quality, and resolving incidents with autonomy you match to risk.
Canvas agents can now read your code, incidents, runbooks and tickets, to get to the right solution, first time. Also new: AI Ecosystem and LLM cost tracking in early access, Anomaly Detection GA, and onboarding from your coding agent.
The best alert is the one that never fires.
Picture an idyllic Tuesday afternoon: last night's release nudged checkout p99 up 41%, but an agent noticed. Nobody got paged, and your SLOs didn't start burning. Instead, the agent acted and already narrowed it down to one retry change. It filed a ticket for the Payments team and flagged the PR.
By the time anyone would have noticed, the fix is in review. That's what SRE was always supposed to feel like.
This Innovation Week, we're previewing Honeycomb Team Agents: AI SRE agents that join your team as coworkers and proactively check on releases and hygiene, not just responding to alerts. They arrive at AWS re:Invent on December 1.
In theory, reliable systems, fewer pages, and teams that own their services all sounded great. But in practice, SREs are buried in reactive work. The goal of every team owning their services' reliability requires that those teams learn SRE best practices and this can take forever. And AI has cranked the rate of change to eleven. Code ships faster than ever, and the agents writing it can't see production at all.
The first wave of “AI SRE” claimed to help fight the fires, but that was only ever a small portion of what an SRE does. Those tools are reactive: they wake up when an alert fires, make a suggestion, and stop there. Worse, they reason over fragmented logs and metrics, which is how you end up with questionable accuracy… which nobody trusts in production.
AI SRE deserves better. Here's what that can look like when it's done right.
Honeycomb agents come in two kinds. Every user gets a personal agent . Ask it anything about production, in Honeycomb or Slack, and it answers with evidence, remembering how you work. Honeycomb customer teams get team agents , which are essentially AI SRE coworkers that run in the background on goals you define, and report back with findings.
AI SRE agents that go beyond incidents
Real SRE work goes beyond incident response: checking releases, keeping instrumentation current, watching cost, keeping runbooks current, and making sure the AI features you shipped last month still behave as expected. Incidents are the part everyone sees, but the rest is what prevents them.
Setting a team agent goal takes a sentence, and from then on, it does the rest, whether that's after a deploy, when an alert fires, or on a schedule.
While you can create custom goals, here are six examples teams might hand over to a team agent:
Verify checkout releases (runs on deploys). Check every checkout service deploy for latency and error regressions in the 30 minutes after it ships.
Flag Lambda spend (runs on alerts). Protect margins by flagging Lambda spend that threatens your margin SLO.
Improve instrumentation (runs on a schedule). Review your Honeycomb datasets and recommend instrumentation improvements, filed in Linear.
Investigate alerts (runs on alerts). Pick up triggers and burn alerts, recommend remediation, open an incident channel, and page a responder.
Improve agent quality (runs on deploys). Compare Braintrust eval scores with live conversation traces after every prompt or model change.
Maintain the knowledge doc (runs on a schedule). Review findings and Slack threads, then draft updates to your Notion or Confluence docs.
Only one of these directly relates to incidents. The other five are the proactive work that most AI SRE tools aren't built for and most human SRE teams aren't staffed to keep up with.
Built on best-in-class tracing
AI suffers from the “garbage in, garbage out” problem, but at speed and scale. Most AI SRE tools sit on top of someone else's telemetry and reason over whatever fragments they can reach. Honeycomb's agents start from high-cardinality, full-context traces . They query them fast enough to keep asking until they find the answer, not just an answer. And you can augment that context beyond production telemetry with MCP connectors that bring in deploys, documentation, cost, and other data.
Every finding links to the query, trace, or deploy behind it, and every number says how it knows: measured, derived, estimated, or unset. Since the agents live where the telemetry does, they can tell you what your instrumentation is missing. A tool layered on borrowed data has no reason to.
SRE expertise has always been spread too thin. Team Agents start with Honeycomb's playbook on day one, a decade of observability and SRE practice built into their skills. Then they learn yours.
Team memory tracks what broke, how often, and whether the fix held, so your team only has to learn something once. Every entry shows where it came from, and anyone on the team can correct it. Your personal agent keeps your corrections private and can suggest what belongs in team memory. A new teammate inherits hard-won context on day one, and the agent investigating next week's alert already knows about last month's.
A lot of the work to avoid the next incident happens between incidents. This is when team agents work on the goals and schedules you set. Whether it's a hygiene audit or closely watching every new deploy on a revenue-impacting service, team agents surface findings: what's true, with evidence, and what the agent already did about it. Using the MCP connectors, those actions might be paging a responder, updating a knowledge base article, drafting a pull request, or submitting a ticket.
Proactive doesn't mean unsupervised. You match autonomy to risk, per goal and per tool:
Observe (human-led): It pulls together context during your investigation. A great place to start and reasonable for high risk, high complexity services with a large blast radius.
Recommend (guided analysis): It proposes root causes and next steps for review.
Act with review (human-approved): It prepares the remediation, and a person makes the call.
Operate (autonomous): It monitors, detects, and resolves issues within your boundaries. Earned with consistently accepted recommendations, and reasonable for lower risk services with strong SLO coverage.
Novel failures with a big blast radius stay human-led, while well-understood services with strong SLO coverage can earn more room. Every agent has its own identity and per-action permissions, so nothing borrows a login and every action is recorded.
Generally available on December 1
Personal agents will be included on every Honeycomb plan. Team agents will be available with Honeycomb Intelligence on Pro and Enterprise plans, priced by hours of agent work rather than tokens.
Team Agents will be generally available at re:Invent on Decembe

[truncated]

## Original Extract

Introducing Honeycomb Team Agents: AI SRE coworkers that proactively check releases, instrumentation, cost, and agent quality, not just respond to alerts.

Honeycomb Team Agents: AI SRE Done Right Going to AWS re:Invent 2026? Join us at booth 1871 for demos, giveaways, and more. Going to AWS re:Invent 2026? Join us at booth 1871 for demos, giveaways, and more. Meet with us Platform Understand Your Systems Investigate with Agents Observe AI Systems Manage Telemetry Operate at Scale Platform Overview Understand and debug complex software and AI systems with full production context. Explore See how your apps, services and experiences connect.
Distributed Tracing Follow requests across services and AI workflows.
Log Analytics Explore structured logs to find issues faster.
Time Series Metrics Track trends and investigate underlying causes.
Frontend Observability Connect user experiences to backend behavior.
Service Map Visualize service dependencies and system health.
BubbleUp Identify what is driving unusual behavior.
Let AI explore production data to resolve issues faster.
Honeycomb Intelligence Investigate and resolve issues with AI.
Canvas Collaborate in an AI-guided investigation workspace.
Team Agents AI SRE coworkers that join your team proactively
Honeycomb MCP Give AI tools access to production context.
Anomaly Detection Surface meaningful changes before users are affected.
Understand agent behavior, performance and cost in production.
AI Ecosystem Track agent behavior, performance and cost across your fleet.
Agent Observability Monitor and debug AI agents end to end.
Collect, control and route telemetry across your stack.
OpenTelemetry Instrument systems with open standards.
Telemetry Pipeline Collect, refine and route telemetry at scale.
Refinery as a Service Control data volume without losing valuable context.
Maintain reliability and efficiency at enterprise scale.
Private Cloud Run Honeycomb in a dedicated environment.
SLOs Set and track customer-focused reliability goals.
App Integrations Connect the tools across your stack.
Honeycomb Services Get expert support to accelerate results.
Partners Discover Honeycomb’s partner ecosystem.
By Use Case By Industry & Size Featured Technologies Explore solutions for specific engineering challenges.
Incident Response Investigate and resolve production issues faster.
LLM Observability Understand LLM quality, latency and cost.
Agent Observability Trace agent decisions, tools and handoffs.
Migrations Maintain visibility throughout your cloud migration.
Frontend Development Connect user experiences to system performance.
The industries and company sizes Honeycomb supports.
Financial Services Operate complex financial systems reliably.
Retail & Ecommerce Protect digital experiences and critical transactions.
Software & Technology Build, operate, and innovate software faster.
Enterprise Scale observability across complex teams and systems.
AI Startups Build and scale reliable AI-native products.
Connect Honeycomb with the technologies across your stack.
OpenTelemetry Send telemetry using open standards.
Amazon Web Services Understand application and infrastructure behavior across AWS.
Microsoft Azure Gain deeper context across your Azure environment.
Kubernetes Understand application and cluster behavior together.
Google Cloud Observe services and workloads across Google Cloud.
AI Agents Connect agent workflows with production telemetry.
All Integrations Explore every Honeycomb integration.
See how Honeycomb improves your business and engineering outcomes.
Predictable Costs Control observability costs as telemetry grows.
DevOps & Releases Ship faster with greater confidence.
Meet Customer SLAs Protect reliability and customer commitments.
See how teams use Honeycomb to achieve measurable results.
Understand how Honeycomb differs from other approaches.
Observability Engineering Download the second edition free today. Explore Why Honeycomb Resources Learn Hub
Blog Articles Read technical insights and product perspectives.
Getting Started Learn the fundamentals of using Honeycomb.
Docs Explore Honeycomb’s technical documentation.
Guides Find practical guidance for common challenges.
Success Stories See how customers achieve real results.
Webinars & Videos Watch expert sessions, product demos & walkthroughs.
Academy Develop skills through structured learning.
Events Discover upcoming Honeycomb events and sessions.
Office Hours Get practical guidance from Honeycomb experts.
Pollinators Slack Connect with observability practitioners on Slack.
Partners Discover Honeycomb’s partner ecosystem.
Support Get help using Honeycomb.
About Us Meet the company behind Honeycomb.
Careers Explore job opportunities at Honeycomb.
News Read company news and announcements.
Log in Sign up Get a demo Platform Solutions Why Honeycomb Resources Pricing Platform
Understand and debug complex software and AI systems with full production context. Understand Your Systems
Honeycomb Team Agents: AI SRE Done Right
Honeycomb Team Agents: AI SRE Done Right
Honeycomb Team Agents are AI SRE coworkers that join your team proactively, not just reactively. Personal agents answer questions about production with evidence; team agents run on goals you define, checking releases, instrumentation, cost, agent quality, and resolving incidents with autonomy you match to risk.
Canvas agents can now read your code, incidents, runbooks and tickets, to get to the right solution, first time. Also new: AI Ecosystem and LLM cost tracking in early access, Anomaly Detection GA, and onboarding from your coding agent.
The best alert is the one that never fires.
Picture an idyllic Tuesday afternoon: last night's release nudged checkout p99 up 41%, but an agent noticed. Nobody got paged, and your SLOs didn't start burning. Instead, the agent acted and already narrowed it down to one retry change. It filed a ticket for the Payments team and flagged the PR.
By the time anyone would have noticed, the fix is in review. That's what SRE was always supposed to feel like.
This Innovation Week, we're previewing Honeycomb Team Agents: AI SRE agents that join your team as coworkers and proactively check on releases and hygiene, not just responding to alerts. They arrive at AWS re:Invent on December 1.
In theory, reliable systems, fewer pages, and teams that own their services all sounded great. But in practice, SREs are buried in reactive work. The goal of every team owning their services' reliability requires that those teams learn SRE best practices and this can take forever. And AI has cranked the rate of change to eleven. Code ships faster than ever, and the agents writing it can't see production at all.
The first wave of “AI SRE” claimed to help fight the fires, but that was only ever a small portion of what an SRE does. Those tools are reactive: they wake up when an alert fires, make a suggestion, and stop there. Worse, they reason over fragmented logs and metrics, which is how you end up with questionable accuracy… which nobody trusts in production.
AI SRE deserves better. Here's what that can look like when it's done right.
Honeycomb agents come in two kinds. Every user gets a personal agent . Ask it anything about production, in Honeycomb or Slack, and it answers with evidence, remembering how you work. Honeycomb customer teams get team agents , which are essentially AI SRE coworkers that run in the background on goals you define, and report back with findings.
AI SRE agents that go beyond incidents
Real SRE work goes beyond incident response: checking releases, keeping instrumentation current, watching cost, keeping runbooks current, and making sure the AI features you shipped last month still behave as expected. Incidents are the part everyone sees, but the rest is what prevents them.
Setting a team agent goal takes a sentence, and from then on, it does the rest, whether that's after a deploy, when an alert fires, or on a schedule.
While you can create custom goals, here are six examples teams might hand over to a team agent:
Verify checkout releases (runs on deploys). Check every checkout service deploy for latency and error regressions in the 30 minutes after it ships.
Flag Lambda spend (runs on alerts). Protect margins by flagging Lambda spend that threatens your margin SLO.
Improve instrumentation (runs on a schedule). Review your Honeycomb datasets and recommend instrumentation improvements, filed in Linear.
Investigate alerts (runs on alerts). Pick up triggers and burn alerts, recommend remediation, open an incident channel, and page a responder.
Improve agent quality (runs on deploys). Compare Braintrust eval scores with live conversation traces after every prompt or model change.
Maintain the knowledge doc (runs on a schedule). Review findings and Slack threads, then draft updates to your Notion or Confluence docs.
Only one of these directly relates to incidents. The other five are the proactive work that most AI SRE tools aren't built for and most human SRE teams aren't staffed to keep up with.
Built on best-in-class tracing
AI suffers from the “garbage in, garbage out” problem, but at speed and scale. Most AI SRE tools sit on top of someone else's telemetry and reason over whatever fragments they can reach. Honeycomb's agents start from high-cardinality, full-context traces . They query them fast enough to keep asking until they find the answer, not just an answer. And you can augment that context beyond production telemetry with MCP connectors that bring in deploys, documentation, cost, and other data.
Every finding links to the query, trace, or deploy behind it, and every number says how it knows: measured, derived, estimated, or unset. Since the agents live where the telemetry does, they can tell you what your instrumentation is missing. A tool layered on borrowed data has no reason to.
SRE expertise has always been spread too thin. Team Agents start with Honeycomb's playbook on day one, a decade of observability and SRE practice built into their skills. Then they learn yours.
Team memory tracks what broke, how often, and whether the fix held, so your team only has to learn something once. Every entry shows where it came from, and anyone on the team can correct it. Your personal agent keeps your corrections private and can suggest what belongs in team memory. A new teammate inherits hard-won context on day one, and the agent investigating next week's alert already knows about last month's.
A lot of the work to avoid the next incident happens between incidents. This is when team agents work on the goals and schedules you set. Whether it's a hygiene audit or closely watching every new deploy on a revenue-impacting service, team agents surface findings: what's true, with evidence, and what the agent already did about it. Using the MCP connectors, those actions might be paging a responder, updating a knowledge base article, drafting a pull request, or submitting a ticket.
Proactive doesn't mean unsupervised. You match autonomy to risk, per goal and per tool:
Observe (human-led): It pulls together context during your investigation. A great place to start and reasonable for high risk, high complexity services with a large blast radius.
Recommend (guided analysis): It proposes root causes and next steps for review.
Act with review (human-approved): It prepares the remediation, and a person makes the call.
Operate (autonomous): It monitors, detects, and resolves issues within your boundaries. Earned with consistently accepted recommendations, and reasonable for lower risk services with strong SLO coverage.
Novel failures with a big blast radius stay human-led, while well-understood services with strong SLO coverage can earn more room. Every agent has its own identity and per-action permissions, so nothing borrows a login and every action is recorded.
Generally available on December 1
Personal agents will be included on every Honeycomb plan. Team agents will be available with Honeycomb Intelligence on Pro and Enterprise plans, priced by hours of agent work rather than tokens.
Team Agents will be generally available at re:Invent on Decembe

[truncated]
