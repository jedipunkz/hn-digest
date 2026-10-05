---
source: "https://www.atlassian.com/blog/ai-at-work/ai-sdlc-transformation-playbook"
hn_url: "https://news.ycombinator.com/item?id=49967903"
title: "The AI SDLC transformation playbook"
article_title: "The AI SDLC transformation playbook - Inside Atlassian"
image: "https://atlassianblog.wpengine.com/wp-content/uploads/2026/09/blog_email-1120x545px-a-1.png"
author: "nnutter"
captured_at: "2026-10-05T17:59:23Z"
capture_tool: "hn-digest"
hn_id: 49967903
score: 1
comments: 0
posted_at: "2026-10-05T17:38:05Z"
tags:
  - hacker-news
---

# The AI SDLC transformation playbook

- HN: [49967903](https://news.ycombinator.com/item?id=49967903)
- Source: [www.atlassian.com](https://www.atlassian.com/blog/ai-at-work/ai-sdlc-transformation-playbook)
- Score: 1
- Comments: 0
- Posted: 2026-10-05T17:38:05Z

## Translation

Title: The AI SDLC transformation playbook
Article title: The AI SDLC transformation playbook - Inside Atlassian
Description: Teams across the industry are sprinting to reinvent their software development processes with agentic AI. Yet, navigating the fast-changing landscape of tools and opinions can be daunting for even the most AI-forward organizations. This guide shares Atlassian’s approach and best practices, informed
[truncated]

Article text:
The AI SDLC transformation playbook - Inside Atlassian
Skip to content
AI that knows your business
Connect people, knowledge, and work to move faster together
The AI SDLC transformation playbook
Research and insights on how teams can deliver results with AI.
Your time is valuable. We promise to send only what’s actually worth reading.
Share on Facebook (Opens in new window) Facebook
Share on X (Opens in new window) X
Share on LinkedIn (Opens in new window) LinkedIn
Share on Mail (Opens in new window) Mail
Develop AI-empowered teams that deliver results
Proven methods and insights to help teams collaborate
Put the future of work into practice
Better ways to build and ship software
A look inside Atlassian’s approach to building apps and teams
Get more value from AI with connected context
AI-powered apps – driven by your team’s knowledge
Deliver service at high velocity
Explore posts for all Atlassian apps
Sign up now
Dismiss
Subscribe to our newsletter
Research and insights on how teams can deliver results with AI.
Your time is valuable. We promise to send only what’s actually worth reading.
Develop AI-empowered teams that deliver results
Proven methods and insights to help teams collaborate
Put the future of work into practice
Better ways to build and ship software
A look inside Atlassian’s approach to building apps and teams
Get more value from AI with connected context
AI-powered apps – driven by your team’s knowledge
Deliver service at high velocity
Explore posts for all Atlassian apps
Research and insights on how teams can deliver results with AI.
Your time is valuable. We promise to send only what’s actually worth reading.
The AI SDLC transformation playbook
Teams across the industry are sprinting to reinvent their software development processes with agentic AI. Yet, navigating the fast-changing landscape of tools and opinions can be daunting for even the most AI-forward organizations.
This guide shares Atlassian’s approach and best practices, informed by our own internal AI SDLC transformation and collaborative learnings from top engineering organizations. As our approach continues to evolve, the foundational elements we describe here are the ones we’ve found consistently mattered.
Read the full report below, or download your copy here .
Making the shift to AI-native SDLC
Traditional vs. AI-native SDLC workflows
Platform investments
Context graph
Making the shift to AI-native SDLC
The AI-native software development lifecycle (SDLC) goes far beyond having developers adopt coding agents. As less time is spent writing code by hand, coding itself is evolving from a synchronous, single-threaded task into an asynchronous, multi-threaded one. This exposes the processes and steps before and after code (planning, design, review, and maintenance) as the largest bottlenecks and opportunities. Furthermore, team collaboration becomes an even greater bottleneck as the rate of work accelerates.
To successfully transform the entire SDLC, organizations must simultaneously invest in key platform foundations—context and measurement—while also shifting from traditional SDLC processes to AI-native ways of working. Together, these investments and workflow shifts become the modern software factory: faster, more autonomous, more event-driven, and reshaped around critical human judgment.
Platform investments and new SDLC workflows are mutually dependent: new SDLC workflows only work effectively when supported by powerful enterprise context and feedback loops. Inversely, reliable context and measurement are only possible when metadata from new SDLC workflows is consistently captured and accessible. As part of that loop, developer–agent interactions themselves become a new form of context, making the whole workflow smarter over time.
Systems of record tie these pillars together, wherever development happens. They support humans in applying the critical judgment that moves software forward: planning, evolving designs, aligning on goals, acting on customer feedback, and responsibly shipping code. In turn, agents rely on those same systems for the structured, connected data that lets people and AI work effectively together.
Check out The Agentic Pivot: Engineering leaders share the reality of AI in the SDLC – key findings from a 2026 survey of 1,000+ software professionals on the state of AI-native engineering within their organizations.
Traditional vs. AI-native SDLC workflows
The software development lifecycle (SDLC) moves software from idea to production through planning, design, engineering, testing, deployment, and maintenance. Agile and DevOps reshaped how teams work together in the traditional SDLC by tightening feedback loops, breaking down silos, and letting teams own the code they ship and run. Engineering organizations have come a long way, but haven’t eliminated the delays that come with every hand-off.
An AI-native SDLC is the next evolution. It builds on the same vision and collapses remaining hand-offs into continuous, AI-supported loops where product, design, and engineering co-create in real time, and where agents help humans do the work at a new pace and scale.
This playbook describes the key SDLC workflow shifts summarized in the table below. Today, most organizations sit somewhere between the traditional and AI-native models. To realize the full benefits of an AI-native SDLC at enterprise scale, it’s essential that work activity and knowledge are captured in systems of record in order to provide AI with continuous context and consistent measurement abilities.
AI agents may have the intelligence to write, refactor, and review code, but they face the same challenge development teams always have: making sense of the context surrounding each task. Humans continuously gather that context, whether its understanding of code, product strategy, architectural decisions, standards, and constraints, accumulated over years of working in the same system. Agents need the same context supplied to them for every task in order to produce high-quality output.
To provide effective context to agents, organizations need a context graph containing interconnected data spanning code repositories, work items, documentation, technical standards, meeting recordings, and chat. This data must then be made accessible to both humans and agents through a permissions-aware layer, so that AI can act as a knowledgeable teammate rather than guessing at the shape of the system.
To tackle this problem, we’ve developed a context engine called the Teamwork Graph. It combines data from 50+ connected sources with Jira, Confluence, and Loom: a living map of work, code, people, decisions, and dependencies that helps agents understand not just the task, but the system around it. Governance and permissions are enforced at the data layer, so what a person can see is exactly what an agent can act on. The Teamwork Graph is reachable from wherever teams work: Atlassian apps, terminals, IDEs, or any agent surface via MCP or CLI. As agents and humans work in these surfaces, their activity flows back into the graph, so the system of record maintains itself autonomously and context compounds every time the graph is used.
In internal benchmarks, agents with access to the Teamwork Graph delivered 44% more accurate results while using 48% fewer tokens compared to agents operating without it. Because every agent run generates more structured context, each subsequent run becomes higher quality and lower cost.
Amidst such rapid change to how teams work, it’s difficult to know what’s working and what’s not without data. With the right feedback loops and metrics, organizations can confidently track the impact of AI tools and workflows alongside overall cost, productivity, and ROI while enabling teams to continuously improve their own processes.
To produce actionable insights, work and agent activity must be observable and measurable. By integrating data from systems of record, third-party APIs, and agent sessions, organizations can build real-time visibility into AI effectiveness, impact, and ROI, and pair it with qualitative signals from the developers doing the work.
At Atlassian, we use DX to measure and report on how AI is impacting our SDLC and overall productivity. DX combines quantitative data from tools like Jira, Cursor, and Claude Code with qualitative developer feedback. Analyzing and benchmarking this dataset helps us set targets, guide investment, and ensure organizational excellence.
“When an agent starts in the wrong place, everything slows down and costs more. Developers end up explaining where to look, why the code works the way it does, and what else it touches. Teamwork Graph puts all of that in front of the agent from the start, so it can focus on getting the work done.”
– Mark Walz, Chief Technology Officer , SpotOn
In the traditional SDLC, planning relies on manual effort and committee-driven consensus. Product managers and analysts spend weeks drafting requirements and aligning stakeholders. Because those defining the work often have incomplete knowledge of the underlying codebase, specifications frequently miss architectural constraints and technical debt, resulting in friction and costly rework during implementation.
An AI-native approach transforms this phase by using agents to synthesize customer needs and organizational context into requirements and specs, drafting from a real understanding of both the business and the codebase. Instead of human teams writing specifications from scratch, they focus on applying judgment: refining, deciding, and ensuring plans are grounded in reality before engineering work begins.
Atlassian operationalizes this vision through the Teamwork Graph, which brings together people, goals, code, and knowledge across Atlassian and connected third-party apps, making it available to any agent. Building on this rich context, teams evaluate ideas against customer feedback in Jira Product Discovery, then use Jira to co-create specs, PRDs, and technical blueprints. Once finalized, these plans can be instantly converted into actionable, tracked work items.
In the traditional SDLC, design is an isolated phase between requirements gathering and implementation. Designers spend significant time crafting static mockups and manually written briefs, then hand them off to engineering. Because designers lack direct access to the underlying codebase, even minor visual tweaks or layout adjustments require developer intervention, creating recurring bottlenecks and dragging out release timelines.
An AI-native SDLC reimagines design as a continuous thread integrated with planning and development, rather than a distinct step. AI-assisted tooling empowers designers to build live, interactive prototypes, generate UI code, and commit bug fixes independently, bridging the gap between interface concept and production without relying on developers for every layout change.
Atlassian brings design artifacts into the Teamwork Graph, connecting design intent to code and work in flight. Designers, PMs, and developers can generate rapid prototypes and design updates directly from shared context, feeding them into code repositories and treating the resulting code as a design spec. To keep collaboration fluid, teams use Loom to capture feedback, where AI ingests the transcripts and summaries as prompts, making revisions a part of the handshake between crafts.
In the traditional SDLC, work is scoped and logged in backlogs before any code is written. Code quality is constrained by individual developer skill and domain expertise, and velocity is bounded by human throughput, the physical limit of how quickly engineers can write, test, and refactor code.
An AI-native SDLC turns development into a fluid, automated process. Tasks are spawned from operational signals and captured automatically as work happens. In this model, code quality is no longer determined by raw developer skill, it depends on the quality of context supplie

[truncated]

## Original Extract

Teams across the industry are sprinting to reinvent their software development processes with agentic AI. Yet, navigating the fast-changing landscape of tools and opinions can be daunting for even the most AI-forward organizations. This guide shares Atlassian’s approach and best practices, informed
[truncated]

The AI SDLC transformation playbook - Inside Atlassian
Skip to content
AI that knows your business
Connect people, knowledge, and work to move faster together
The AI SDLC transformation playbook
Research and insights on how teams can deliver results with AI.
Your time is valuable. We promise to send only what’s actually worth reading.
Share on Facebook (Opens in new window) Facebook
Share on X (Opens in new window) X
Share on LinkedIn (Opens in new window) LinkedIn
Share on Mail (Opens in new window) Mail
Develop AI-empowered teams that deliver results
Proven methods and insights to help teams collaborate
Put the future of work into practice
Better ways to build and ship software
A look inside Atlassian’s approach to building apps and teams
Get more value from AI with connected context
AI-powered apps – driven by your team’s knowledge
Deliver service at high velocity
Explore posts for all Atlassian apps
Sign up now
Dismiss
Subscribe to our newsletter
Research and insights on how teams can deliver results with AI.
Your time is valuable. We promise to send only what’s actually worth reading.
Develop AI-empowered teams that deliver results
Proven methods and insights to help teams collaborate
Put the future of work into practice
Better ways to build and ship software
A look inside Atlassian’s approach to building apps and teams
Get more value from AI with connected context
AI-powered apps – driven by your team’s knowledge
Deliver service at high velocity
Explore posts for all Atlassian apps
Research and insights on how teams can deliver results with AI.
Your time is valuable. We promise to send only what’s actually worth reading.
The AI SDLC transformation playbook
Teams across the industry are sprinting to reinvent their software development processes with agentic AI. Yet, navigating the fast-changing landscape of tools and opinions can be daunting for even the most AI-forward organizations.
This guide shares Atlassian’s approach and best practices, informed by our own internal AI SDLC transformation and collaborative learnings from top engineering organizations. As our approach continues to evolve, the foundational elements we describe here are the ones we’ve found consistently mattered.
Read the full report below, or download your copy here .
Making the shift to AI-native SDLC
Traditional vs. AI-native SDLC workflows
Platform investments
Context graph
Making the shift to AI-native SDLC
The AI-native software development lifecycle (SDLC) goes far beyond having developers adopt coding agents. As less time is spent writing code by hand, coding itself is evolving from a synchronous, single-threaded task into an asynchronous, multi-threaded one. This exposes the processes and steps before and after code (planning, design, review, and maintenance) as the largest bottlenecks and opportunities. Furthermore, team collaboration becomes an even greater bottleneck as the rate of work accelerates.
To successfully transform the entire SDLC, organizations must simultaneously invest in key platform foundations—context and measurement—while also shifting from traditional SDLC processes to AI-native ways of working. Together, these investments and workflow shifts become the modern software factory: faster, more autonomous, more event-driven, and reshaped around critical human judgment.
Platform investments and new SDLC workflows are mutually dependent: new SDLC workflows only work effectively when supported by powerful enterprise context and feedback loops. Inversely, reliable context and measurement are only possible when metadata from new SDLC workflows is consistently captured and accessible. As part of that loop, developer–agent interactions themselves become a new form of context, making the whole workflow smarter over time.
Systems of record tie these pillars together, wherever development happens. They support humans in applying the critical judgment that moves software forward: planning, evolving designs, aligning on goals, acting on customer feedback, and responsibly shipping code. In turn, agents rely on those same systems for the structured, connected data that lets people and AI work effectively together.
Check out The Agentic Pivot: Engineering leaders share the reality of AI in the SDLC – key findings from a 2026 survey of 1,000+ software professionals on the state of AI-native engineering within their organizations.
Traditional vs. AI-native SDLC workflows
The software development lifecycle (SDLC) moves software from idea to production through planning, design, engineering, testing, deployment, and maintenance. Agile and DevOps reshaped how teams work together in the traditional SDLC by tightening feedback loops, breaking down silos, and letting teams own the code they ship and run. Engineering organizations have come a long way, but haven’t eliminated the delays that come with every hand-off.
An AI-native SDLC is the next evolution. It builds on the same vision and collapses remaining hand-offs into continuous, AI-supported loops where product, design, and engineering co-create in real time, and where agents help humans do the work at a new pace and scale.
This playbook describes the key SDLC workflow shifts summarized in the table below. Today, most organizations sit somewhere between the traditional and AI-native models. To realize the full benefits of an AI-native SDLC at enterprise scale, it’s essential that work activity and knowledge are captured in systems of record in order to provide AI with continuous context and consistent measurement abilities.
AI agents may have the intelligence to write, refactor, and review code, but they face the same challenge development teams always have: making sense of the context surrounding each task. Humans continuously gather that context, whether its understanding of code, product strategy, architectural decisions, standards, and constraints, accumulated over years of working in the same system. Agents need the same context supplied to them for every task in order to produce high-quality output.
To provide effective context to agents, organizations need a context graph containing interconnected data spanning code repositories, work items, documentation, technical standards, meeting recordings, and chat. This data must then be made accessible to both humans and agents through a permissions-aware layer, so that AI can act as a knowledgeable teammate rather than guessing at the shape of the system.
To tackle this problem, we’ve developed a context engine called the Teamwork Graph. It combines data from 50+ connected sources with Jira, Confluence, and Loom: a living map of work, code, people, decisions, and dependencies that helps agents understand not just the task, but the system around it. Governance and permissions are enforced at the data layer, so what a person can see is exactly what an agent can act on. The Teamwork Graph is reachable from wherever teams work: Atlassian apps, terminals, IDEs, or any agent surface via MCP or CLI. As agents and humans work in these surfaces, their activity flows back into the graph, so the system of record maintains itself autonomously and context compounds every time the graph is used.
In internal benchmarks, agents with access to the Teamwork Graph delivered 44% more accurate results while using 48% fewer tokens compared to agents operating without it. Because every agent run generates more structured context, each subsequent run becomes higher quality and lower cost.
Amidst such rapid change to how teams work, it’s difficult to know what’s working and what’s not without data. With the right feedback loops and metrics, organizations can confidently track the impact of AI tools and workflows alongside overall cost, productivity, and ROI while enabling teams to continuously improve their own processes.
To produce actionable insights, work and agent activity must be observable and measurable. By integrating data from systems of record, third-party APIs, and agent sessions, organizations can build real-time visibility into AI effectiveness, impact, and ROI, and pair it with qualitative signals from the developers doing the work.
At Atlassian, we use DX to measure and report on how AI is impacting our SDLC and overall productivity. DX combines quantitative data from tools like Jira, Cursor, and Claude Code with qualitative developer feedback. Analyzing and benchmarking this dataset helps us set targets, guide investment, and ensure organizational excellence.
“When an agent starts in the wrong place, everything slows down and costs more. Developers end up explaining where to look, why the code works the way it does, and what else it touches. Teamwork Graph puts all of that in front of the agent from the start, so it can focus on getting the work done.”
– Mark Walz, Chief Technology Officer , SpotOn
In the traditional SDLC, planning relies on manual effort and committee-driven consensus. Product managers and analysts spend weeks drafting requirements and aligning stakeholders. Because those defining the work often have incomplete knowledge of the underlying codebase, specifications frequently miss architectural constraints and technical debt, resulting in friction and costly rework during implementation.
An AI-native approach transforms this phase by using agents to synthesize customer needs and organizational context into requirements and specs, drafting from a real understanding of both the business and the codebase. Instead of human teams writing specifications from scratch, they focus on applying judgment: refining, deciding, and ensuring plans are grounded in reality before engineering work begins.
Atlassian operationalizes this vision through the Teamwork Graph, which brings together people, goals, code, and knowledge across Atlassian and connected third-party apps, making it available to any agent. Building on this rich context, teams evaluate ideas against customer feedback in Jira Product Discovery, then use Jira to co-create specs, PRDs, and technical blueprints. Once finalized, these plans can be instantly converted into actionable, tracked work items.
In the traditional SDLC, design is an isolated phase between requirements gathering and implementation. Designers spend significant time crafting static mockups and manually written briefs, then hand them off to engineering. Because designers lack direct access to the underlying codebase, even minor visual tweaks or layout adjustments require developer intervention, creating recurring bottlenecks and dragging out release timelines.
An AI-native SDLC reimagines design as a continuous thread integrated with planning and development, rather than a distinct step. AI-assisted tooling empowers designers to build live, interactive prototypes, generate UI code, and commit bug fixes independently, bridging the gap between interface concept and production without relying on developers for every layout change.
Atlassian brings design artifacts into the Teamwork Graph, connecting design intent to code and work in flight. Designers, PMs, and developers can generate rapid prototypes and design updates directly from shared context, feeding them into code repositories and treating the resulting code as a design spec. To keep collaboration fluid, teams use Loom to capture feedback, where AI ingests the transcripts and summaries as prompts, making revisions a part of the handshake between crafts.
In the traditional SDLC, work is scoped and logged in backlogs before any code is written. Code quality is constrained by individual developer skill and domain expertise, and velocity is bounded by human throughput, the physical limit of how quickly engineers can write, test, and refactor code.
An AI-native SDLC turns development into a fluid, automated process. Tasks are spawned from operational signals and captured automatically as work happens. In this model, code quality is no longer determined by raw developer skill, it depends on the quality of context supplie

[truncated]
