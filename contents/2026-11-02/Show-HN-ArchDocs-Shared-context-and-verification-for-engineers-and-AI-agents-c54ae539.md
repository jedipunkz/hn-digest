---
source: "https://archdocs.io"
hn_url: "https://news.ycombinator.com/item?id=50340570"
title: "Show HN: ArchDocs – Shared context and verification for engineers and AI agents"
article_title: "ArchDocs - The system of record for software engineering"
image: "https://archdocs.io/og/home.png"
author: "AylonArmstrong"
captured_at: "2026-11-02T17:53:08Z"
capture_tool: "hn-digest"
hn_id: 50340570
score: 2
comments: 0
posted_at: "2026-11-02T17:44:32Z"
tags:
  - hacker-news
---

# Show HN: ArchDocs – Shared context and verification for engineers and AI agents

- HN: [50340570](https://news.ycombinator.com/item?id=50340570)
- Source: [archdocs.io](https://archdocs.io)
- Score: 2
- Comments: 0
- Posted: 2026-11-02T17:44:32Z

## Translation

Title: Show HN: ArchDocs – Shared context and verification for engineers and AI agents
Article title: ArchDocs - The system of record for software engineering
Description: Design it. Build it. Verify it. Connect design docs, architecture, decisions and recorded work for your team and your coding agents, then review what the record supports.
HN text: Hi HN, My name is Aylon and I'm building ArchDocs. I started building ArchDocs in late 2024 as an architecture diagramming tool with reusable references and objects (still exists today and is usable). As I've been using AI to code more and more, one of the things I've run into is that my agent doesn't have enough context around past decisions or findings that could affect an implementation, causing some nasty side-effects like redundant investigations, incorrect assumptions, repeating previously caught mistakes and ultimately wasting precious tokens. Let's say for example, you need to do a migration that affects 10 backend services. Your agent starts working on the first one and during the implementation and verification process, it discovers an issue that will undeniably affect the other 9 services. The next agent can start on the second service without knowing what was learned on the first. There's another problem I run up against time and time again. I need a dependable way to verify and log what an agent has done, the outcome of that and any deviations or issues that arise from it, so I can tell whether the change actually did what I intended. Enter Design Docs inside ArchDocs. The workflow is simple. You or your agent can craft a spec and ask questions for people to answer. Writing the implementation plan comes next, followed by doing the actual work. While you or your agent works, you track important milestones, deviations, decisions and more in the work log. Need to pause or have someone else do the work? No problem, craft a handoff or let ArchDocs AI do it for you. Finally, you or your agent logs a summarized version of what was done and how it was verified in an evidence of work submission. In that migration example, your agent can mark the issue in the work log as a "finding" and that will then be surfaced when future agents execute on the other 9 services. On the verification side, you may ask, "How does this help if the agent or engineer is just writing down their verification?" This is a good question and that's only one part of the verification. We have this concept of Agent attested (paired with human verified) and human attested. The last pillar and the most important is ArchDocs' own independent, automated verifications. Right now, we just have GitHub tracking for check results but we will soon be expanding this into so much more. Imagine one of the automated verifications is that a specific log with a specific attribute shows up in Datadog. Maybe even a specific event showing up in PostHog. There's a lot of ways we're going to take this and we're excited to share them once done. You can explore an example design doc without signing up. It includes the spec, implementation plan, work log and verification evidence: https://archdocs.io/templates/design-doc-example

Article text:
ArchDocs - The system of record for software engineering Arch Docs Product
The system of record for software engineering.
Design it. Build it. Verify it.
Connect design docs, architecture, decisions, and recorded work. Give your team and coding agents the same plan, then review what the record supports and what still needs confirmation.
Free includes architecture docs, MCP access and one design doc within Free plan limits. Unlimited design docs start on Team.
A team is cutting the tender response window for API-connected carriers from 15 minutes to 5. EDI carriers keep 15 minutes. Follow the change from impact and decisions to review and the next session.
Tender load Reads and writes Consumes Load status ETA updates Planned and recorded S Service Dispatch Service Planned; no recorded change S Service Carrier Gateway Recorded outside the plan S Service Tracking Service A API Public Shipper API D Datastore Bookings DB Q Queue/Topic booking-events Tender load Reads and writes Consumes Load status ETA updates Planned and recorded S Service Dispatch Service Planned; no recorded change S Service Carrier Gateway Recorded outside the plan S Service Tracking Service A API Public Shipper API D Datastore Bookings DB Q Queue/Topic booking-events DD-0021 · Shorter tender window for API carriers
One planned change has no record yet.
Carrier Gateway was in the plan. Confirm whether channel handling was implemented.
Plan versus recorded, approved plan
Map the system your team actually runs
Draw services, queues, databases and APIs once, then reuse them on every diagram. Archdocs shows what depends on each one, from the docs your team keeps.
01 . 1 Endpoint cards, data models, sequence diagrams and cloud blocks on one canvas.
01 . 2 One object per service, so a rename or new owner updates every diagram that uses it.
01 . 3 Every diagram is YAML you can review in Git, with a pull request check that flags docs that may be stale.
Plan · Architecture docs, shared objects and the MCP server are on every plan.
Object Bookings DB Kind Datastore Engine PostgreSQL 16 Change Upgrade to PostgreSQL 17 Dispatch Service doc
Bookings DB is PostgreSQL 16 in all three docs.
Changing Bookings DB affects 9 objects across 7 docs. Example
Decide before you build, and keep the reasons
A design doc holds the spec, the open questions and the decision, next to the services it changes. Decision records keep the options you rejected and why.
02 . 1 Open questions stay open until a person answers them.
02 . 2 Decision records show the considered options and what superseded them.
02 . 3 Named reviewers approve the spec, and later edits are recorded as amendments.
Plan · One design doc on Free; unlimited design docs on Team and above. Review and approvals on Business and Enterprise.
Product screenshots from a sample workspace.
Give people and coding agents the same plan
The implementation plan lists the changes and tasks. Engineers work from it in Archdocs, and coding agents read the same spec, plan and open questions over MCP.
03 . 1 Agents get the open questions with the plan, and Archdocs tells them to ask rather than decide.
03 . 2 Evidence an agent submits is labeled as the agent's report until a person reviews it.
03 . 3 Agents ask for a task's context before they start, and leave a handoff when they stop.
03 . 4 Send plan tasks to Linear when your team tracks work there. Linear is on Business and Enterprise, in early access.
Plan · The MCP server is on every plan. One design doc on Free; unlimited design docs on Team and above.
1 . Asks for context Task: shorten the tender window for API carriers. Current: DD-0021, approved by a person; ADR-0003. Open question: Q1, assigned to Maya Chen. 2 . Takes a task Channel tag in Carrier Gateway, from the implementation plan. 3 . Reports Marks the task done. Recorded as Agent-reported; no recorded change for Carrier Gateway yet. 4 . Leaves a handoff E-4 by Claude Code, run by Priya Raman: Done, Blocked on Q1, Next: confirm the Carrier Gateway change. fig. 03.1 · An agent session with context, work and a handoff Recreated from Archdocs views with sample data. Product screenshot from a sample workspace. The agent session is recreated from Archdocs views with sample data.
Review what landed against the plan
Compare the planned changes with the changes your team and agents recorded, next to pull requests and checks from GitHub.
04 . 1 Gaps show both ways: planned work with no recorded change, and changes nobody planned.
04 . 2 Pull request and check results come from GitHub; Archdocs doesn't read your code.
04 . 3 Acceptance criteria say what has to be true, and each piece of evidence shows where it came from.
Plan · Plan versus actual on every plan. GitHub evidence on Business and Enterprise. Acceptance criteria on every plan, up to 10 per design doc on Free.
This compares the plan with recorded activity. Missing records do not establish whether implementation occurred.
Dispatch Service : Per-channel tender window
Tracking Service : Shipper countdown uses the new window
Product screenshots from a sample workspace.
Pick up where the last session stopped
Each design doc keeps a work log of what was done, what's blocked and what comes next, with every entry marked as reported by a person or an agent. When someone picks the task up, they ask for its context and get the decisions, open questions and history linked to that work.
05 . 1 Handoffs list what's done, what's blocked and what's next, with a resume prompt for the next session.
05 . 2 Commits cited in an entry are checked against GitHub when Archdocs can read the repository.
05 . 3 Task context labels every record with who approved it and how it was found.
05 . 4 Agents ask for it over MCP; people use the Context page. Records you can't open are left out.
Plan · Work logs come with design docs on every plan. Task context is on Team, Business and Enterprise.
DD-0021 · Shorter tender window for API carriers Work log Work log
E- 3 Person-reported Priya Raman
Per-channel tender window merged in #482
E- 4 Agent-reported Claude Code run by Priya Raman
Left a handoff for the next session.
Left by Claude Code run by Priya Raman
Channel tag in Carrier Gateway marked done by the agent. No change recorded for it yet; confirm next.
2 of 3 checked in the Implementation plan
Q1 · What happens when a carrier accepts after the window closes?
Copy resume prompt fig. 05.2 · Context page: records linked to the task Context
Scope: DD-0021 and lanefold/dispatch
DD-0021 · Shorter tender window for API carriers
ADR-0003 · Prefer carrier REST APIs over EDI where available
Q1 · What happens when a carrier accepts after the window closes?
DD-0021 · Shorter tender window for API carriers
DD-0021 plans to change Carrier Gateway, not yet recorded as done
ADR-0004 · Queue late acceptances for re-offer
ADR-0002 · Tender loads over EDI 204 before carrier APIs
Records you cannot open are not shown.
Recreated from Archdocs views with sample data.
I built ArchDocs so a service is defined once and used in every document that needs it. Design docs, decisions and recorded changes now build on those same objects.
Yes. Free covers up to three people, ten architecture docs, shared objects, the MCP server and one design doc. Team adds unlimited design docs and task context. Business adds reviews, approvals and GitHub evidence.
Markdown works well for many teams. Archdocs becomes useful when decisions, open questions, architecture references, and implementation records need to stay connected across people and changes. You can import a markdown plan as is.
Shared objects reduce duplicate edits. People and agents can contribute records, and supported integrations add signals about work and documentation that may need review. Your team still decides what is correct.
Yes. Code stays in GitHub and work tracking stays in Linear. Archdocs links pull requests, checks and issues to the design doc they belong to, so the engineering record points at those tools instead of replacing them. GitHub is on Team and above; Linear is on Business and Enterprise, in early access.
With one proposal. Write it as a design doc, mention the services it changes, and record the work as it lands. Free includes one design doc and ten architecture docs.
You could keep a folder of docs and a script or two. The work that lasts is maintaining permissions, history, authority labels, attribution, connectors and evidence over time, for people and agents at once. That is what Archdocs maintains for you.
Diagrams export as SVG, PNG or PDF, and every diagram is YAML you can copy. A design doc exports as markdown, whole or one variation at a time, including work log entries. The audit log exports as CSV or JSON. Approvals, acceptance criteria results and decision records are not included in those exports today.
Yes, if your client supports remote MCP servers with OAuth. Connect to https://mcp.archdocs.io/mcp and sign in through your browser. The agent can read and update architecture docs, write design docs and decisions, record changes, add work log entries and ask for a task's context (task context is on Team and above).
From the design doc's work log. The latest handoff says what's done, what's blocked and what's next, and task context gathers the approved decisions, open questions and history linked to the work. Agents get the same over MCP. Work logs come with design docs on every plan. Task context is on Team, Business and Enterprise.
Architecture documents are stored in PostgreSQL in the United States. The API and database run on Northflank, the web apps run on Vercel, and uploaded files are stored in Cloudflare R2.
Start with one engineering change.
Explore the architecture editor for free, or see how Design Docs connect a proposal, its decisions, and the record of the work.
Arch Docs The system of record for software engineering. Design it. Build it. Verify it.
Evidence and acceptance criteria
© 2026 ArchDocs. All rights reserved.

## Original Extract

Design it. Build it. Verify it. Connect design docs, architecture, decisions and recorded work for your team and your coding agents, then review what the record supports.

Hi HN, My name is Aylon and I'm building ArchDocs. I started building ArchDocs in late 2024 as an architecture diagramming tool with reusable references and objects (still exists today and is usable). As I've been using AI to code more and more, one of the things I've run into is that my agent doesn't have enough context around past decisions or findings that could affect an implementation, causing some nasty side-effects like redundant investigations, incorrect assumptions, repeating previously caught mistakes and ultimately wasting precious tokens. Let's say for example, you need to do a migration that affects 10 backend services. Your agent starts working on the first one and during the implementation and verification process, it discovers an issue that will undeniably affect the other 9 services. The next agent can start on the second service without knowing what was learned on the first. There's another problem I run up against time and time again. I need a dependable way to verify and log what an agent has done, the outcome of that and any deviations or issues that arise from it, so I can tell whether the change actually did what I intended. Enter Design Docs inside ArchDocs. The workflow is simple. You or your agent can craft a spec and ask questions for people to answer. Writing the implementation plan comes next, followed by doing the actual work. While you or your agent works, you track important milestones, deviations, decisions and more in the work log. Need to pause or have someone else do the work? No problem, craft a handoff or let ArchDocs AI do it for you. Finally, you or your agent logs a summarized version of what was done and how it was verified in an evidence of work submission. In that migration example, your agent can mark the issue in the work log as a "finding" and that will then be surfaced when future agents execute on the other 9 services. On the verification side, you may ask, "How does this help if the agent or engineer is just writing down their verification?" This is a good question and that's only one part of the verification. We have this concept of Agent attested (paired with human verified) and human attested. The last pillar and the most important is ArchDocs' own independent, automated verifications. Right now, we just have GitHub tracking for check results but we will soon be expanding this into so much more. Imagine one of the automated verifications is that a specific log with a specific attribute shows up in Datadog. Maybe even a specific event showing up in PostHog. There's a lot of ways we're going to take this and we're excited to share them once done. You can explore an example design doc without signing up. It includes the spec, implementation plan, work log and verification evidence: https://archdocs.io/templates/design-doc-example

ArchDocs - The system of record for software engineering Arch Docs Product
The system of record for software engineering.
Design it. Build it. Verify it.
Connect design docs, architecture, decisions, and recorded work. Give your team and coding agents the same plan, then review what the record supports and what still needs confirmation.
Free includes architecture docs, MCP access and one design doc within Free plan limits. Unlimited design docs start on Team.
A team is cutting the tender response window for API-connected carriers from 15 minutes to 5. EDI carriers keep 15 minutes. Follow the change from impact and decisions to review and the next session.
Tender load Reads and writes Consumes Load status ETA updates Planned and recorded S Service Dispatch Service Planned; no recorded change S Service Carrier Gateway Recorded outside the plan S Service Tracking Service A API Public Shipper API D Datastore Bookings DB Q Queue/Topic booking-events Tender load Reads and writes Consumes Load status ETA updates Planned and recorded S Service Dispatch Service Planned; no recorded change S Service Carrier Gateway Recorded outside the plan S Service Tracking Service A API Public Shipper API D Datastore Bookings DB Q Queue/Topic booking-events DD-0021 · Shorter tender window for API carriers
One planned change has no record yet.
Carrier Gateway was in the plan. Confirm whether channel handling was implemented.
Plan versus recorded, approved plan
Map the system your team actually runs
Draw services, queues, databases and APIs once, then reuse them on every diagram. Archdocs shows what depends on each one, from the docs your team keeps.
01 . 1 Endpoint cards, data models, sequence diagrams and cloud blocks on one canvas.
01 . 2 One object per service, so a rename or new owner updates every diagram that uses it.
01 . 3 Every diagram is YAML you can review in Git, with a pull request check that flags docs that may be stale.
Plan · Architecture docs, shared objects and the MCP server are on every plan.
Object Bookings DB Kind Datastore Engine PostgreSQL 16 Change Upgrade to PostgreSQL 17 Dispatch Service doc
Bookings DB is PostgreSQL 16 in all three docs.
Changing Bookings DB affects 9 objects across 7 docs. Example
Decide before you build, and keep the reasons
A design doc holds the spec, the open questions and the decision, next to the services it changes. Decision records keep the options you rejected and why.
02 . 1 Open questions stay open until a person answers them.
02 . 2 Decision records show the considered options and what superseded them.
02 . 3 Named reviewers approve the spec, and later edits are recorded as amendments.
Plan · One design doc on Free; unlimited design docs on Team and above. Review and approvals on Business and Enterprise.
Product screenshots from a sample workspace.
Give people and coding agents the same plan
The implementation plan lists the changes and tasks. Engineers work from it in Archdocs, and coding agents read the same spec, plan and open questions over MCP.
03 . 1 Agents get the open questions with the plan, and Archdocs tells them to ask rather than decide.
03 . 2 Evidence an agent submits is labeled as the agent's report until a person reviews it.
03 . 3 Agents ask for a task's context before they start, and leave a handoff when they stop.
03 . 4 Send plan tasks to Linear when your team tracks work there. Linear is on Business and Enterprise, in early access.
Plan · The MCP server is on every plan. One design doc on Free; unlimited design docs on Team and above.
1 . Asks for context Task: shorten the tender window for API carriers. Current: DD-0021, approved by a person; ADR-0003. Open question: Q1, assigned to Maya Chen. 2 . Takes a task Channel tag in Carrier Gateway, from the implementation plan. 3 . Reports Marks the task done. Recorded as Agent-reported; no recorded change for Carrier Gateway yet. 4 . Leaves a handoff E-4 by Claude Code, run by Priya Raman: Done, Blocked on Q1, Next: confirm the Carrier Gateway change. fig. 03.1 · An agent session with context, work and a handoff Recreated from Archdocs views with sample data. Product screenshot from a sample workspace. The agent session is recreated from Archdocs views with sample data.
Review what landed against the plan
Compare the planned changes with the changes your team and agents recorded, next to pull requests and checks from GitHub.
04 . 1 Gaps show both ways: planned work with no recorded change, and changes nobody planned.
04 . 2 Pull request and check results come from GitHub; Archdocs doesn't read your code.
04 . 3 Acceptance criteria say what has to be true, and each piece of evidence shows where it came from.
Plan · Plan versus actual on every plan. GitHub evidence on Business and Enterprise. Acceptance criteria on every plan, up to 10 per design doc on Free.
This compares the plan with recorded activity. Missing records do not establish whether implementation occurred.
Dispatch Service : Per-channel tender window
Tracking Service : Shipper countdown uses the new window
Product screenshots from a sample workspace.
Pick up where the last session stopped
Each design doc keeps a work log of what was done, what's blocked and what comes next, with every entry marked as reported by a person or an agent. When someone picks the task up, they ask for its context and get the decisions, open questions and history linked to that work.
05 . 1 Handoffs list what's done, what's blocked and what's next, with a resume prompt for the next session.
05 . 2 Commits cited in an entry are checked against GitHub when Archdocs can read the repository.
05 . 3 Task context labels every record with who approved it and how it was found.
05 . 4 Agents ask for it over MCP; people use the Context page. Records you can't open are left out.
Plan · Work logs come with design docs on every plan. Task context is on Team, Business and Enterprise.
DD-0021 · Shorter tender window for API carriers Work log Work log
E- 3 Person-reported Priya Raman
Per-channel tender window merged in #482
E- 4 Agent-reported Claude Code run by Priya Raman
Left a handoff for the next session.
Left by Claude Code run by Priya Raman
Channel tag in Carrier Gateway marked done by the agent. No change recorded for it yet; confirm next.
2 of 3 checked in the Implementation plan
Q1 · What happens when a carrier accepts after the window closes?
Copy resume prompt fig. 05.2 · Context page: records linked to the task Context
Scope: DD-0021 and lanefold/dispatch
DD-0021 · Shorter tender window for API carriers
ADR-0003 · Prefer carrier REST APIs over EDI where available
Q1 · What happens when a carrier accepts after the window closes?
DD-0021 · Shorter tender window for API carriers
DD-0021 plans to change Carrier Gateway, not yet recorded as done
ADR-0004 · Queue late acceptances for re-offer
ADR-0002 · Tender loads over EDI 204 before carrier APIs
Records you cannot open are not shown.
Recreated from Archdocs views with sample data.
I built ArchDocs so a service is defined once and used in every document that needs it. Design docs, decisions and recorded changes now build on those same objects.
Yes. Free covers up to three people, ten architecture docs, shared objects, the MCP server and one design doc. Team adds unlimited design docs and task context. Business adds reviews, approvals and GitHub evidence.
Markdown works well for many teams. Archdocs becomes useful when decisions, open questions, architecture references, and implementation records need to stay connected across people and changes. You can import a markdown plan as is.
Shared objects reduce duplicate edits. People and agents can contribute records, and supported integrations add signals about work and documentation that may need review. Your team still decides what is correct.
Yes. Code stays in GitHub and work tracking stays in Linear. Archdocs links pull requests, checks and issues to the design doc they belong to, so the engineering record points at those tools instead of replacing them. GitHub is on Team and above; Linear is on Business and Enterprise, in early access.
With one proposal. Write it as a design doc, mention the services it changes, and record the work as it lands. Free includes one design doc and ten architecture docs.
You could keep a folder of docs and a script or two. The work that lasts is maintaining permissions, history, authority labels, attribution, connectors and evidence over time, for people and agents at once. That is what Archdocs maintains for you.
Diagrams export as SVG, PNG or PDF, and every diagram is YAML you can copy. A design doc exports as markdown, whole or one variation at a time, including work log entries. The audit log exports as CSV or JSON. Approvals, acceptance criteria results and decision records are not included in those exports today.
Yes, if your client supports remote MCP servers with OAuth. Connect to https://mcp.archdocs.io/mcp and sign in through your browser. The agent can read and update architecture docs, write design docs and decisions, record changes, add work log entries and ask for a task's context (task context is on Team and above).
From the design doc's work log. The latest handoff says what's done, what's blocked and what's next, and task context gathers the approved decisions, open questions and history linked to the work. Agents get the same over MCP. Work logs come with design docs on every plan. Task context is on Team, Business and Enterprise.
Architecture documents are stored in PostgreSQL in the United States. The API and database run on Northflank, the web apps run on Vercel, and uploaded files are stored in Cloudflare R2.
Start with one engineering change.
Explore the architecture editor for free, or see how Design Docs connect a proposal, its decisions, and the record of the work.
Arch Docs The system of record for software engineering. Design it. Build it. Verify it.
Evidence and acceptance criteria
© 2026 ArchDocs. All rights reserved.
