---
source: "https://www.theregister.com/databases/2026/10/29/partner-content-built-to-bend-postgress-inventor-and-a-steward-on-what-ai-really-needs/5305957"
hn_url: "https://news.ycombinator.com/item?id=50291348"
title: "Built to bend: Postgres's inventor and a steward on what AI needs"
article_title: "PARTNER CONTENT: Built to Bend: Postgres’s Inventor and a Steward on What AI Really Needs"
image: "https://image.theregister.com/5305959.jpg?imageId=5305959&x=0&y=0&cropw=100&croph=100&panox=0&panoy=0&panow=100&panoh=100&width=1200&height=683"
author: "DougN7"
captured_at: "2026-10-29T16:39:43Z"
capture_tool: "hn-digest"
hn_id: 50291348
score: 1
comments: 0
posted_at: "2026-10-29T16:10:06Z"
tags:
  - hacker-news
---

# Built to bend: Postgres's inventor and a steward on what AI needs

- HN: [50291348](https://news.ycombinator.com/item?id=50291348)
- Source: [www.theregister.com](https://www.theregister.com/databases/2026/10/29/partner-content-built-to-bend-postgress-inventor-and-a-steward-on-what-ai-really-needs/5305957)
- Score: 1
- Comments: 0
- Posted: 2026-10-29T16:10:06Z

## Translation

Title: Built to bend: Postgres's inventor and a steward on what AI needs
Article title: PARTNER CONTENT: Built to Bend: Postgres’s Inventor and a Steward on What AI Really Needs
Description: PARTNER CONTENT: Michael Stonebraker and Bruce Momjian on why enterprise AI needs a database that no one controls, and what PostgreSQL 19 says about the road ahead

Article text:
Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Built to bend: Postgres’s inventor and a steward on what AI really needs
PARTNER CONTENT: Michael Stonebraker and Bruce Momjian on why enterprise AI needs a database that no one controls, and what PostgreSQL 19 says about the road ahead
Michael Stonebraker introduced Postgres ® at Berkeley in 1986. Forty years later, the database he invented is a leading choice for a generation of AI projects, and he is among the sharpest skeptics of how those projects are being built. Bruce Momjian is cofounder and core team member of the PostgreSQL Global Development Group, and he works at EnterpriseDB (EDB). In upcoming episodes of AI & Data Horizons by EnterpriseDB (EDB), the Turing Award winner and MIT professor and the Postgres community advocate both make the case that the future belongs to systems that bend without breaking, and to communities that no single company controls.
Stonebraker’s first exhibit is text-to-SQL, the promise that a business user can ask a question in plain English and get an answer from the warehouse. Public benchmarks such as Spider and BIRD report accuracy above 80 percent. His team tried the same thing on MIT’s own data warehouse—1,400 tables of student, faculty, building, and salary information—using real queries from its SQL log.
“We got zero-percent accuracy. Not 80, not 70, not 40—zero,” he said. Adding agentic AI, retrieval, and prompt engineering lifted that level into the teens. Telling the model every table and join term reached the 30s. “Still not ready for prime time.”
The reasons are the stuff of every enterprise estate. There is no training data, because the warehouse is not on the public internet. The data is idiosyncratic: among other things, MIT buildings have numbers, not names. The schema has rotted into cryptic table and column names and overlapping views. And real queries are harder than benchmark ones. His team has published its own benchmark, Beaver, and says nobody has broken out of accuracy levels in the 30-percent range.
Use the model for what it’s good at
The deeper problem, Stonebraker argues, is architectural. Large language models are good with text, but “what they are no good at is doing joins of structured data,” which is most of what an enterprise holds. Working with the city of Munich’s transport department, Stonebraker’s team found that answering a single citizen complaint about crosswalk timing meant combining regulations, CAD drawings, trolley schedules, and school-zone maps. Orchestrating that with a language model at the center was, he said, “ridiculously expensive.”
The team’s answer, a system called Rubicon, uses models only on text, then casts the results into tables and hands the rest to a tabular engine. The orchestrator “looks like a query optimizer.” Stonebraker says it is cheaper by a factor of 10 and considerably more accurate than agent frameworks built around a model.
He expects agentic AI to slide from the top of the hype cycle into the valley of despair before a more balanced view emerges. Enterprises hate non-determinism: Approve a loan and reverse it on review, and “that’s called a lawsuit.” Organizations also need protection from rogue agents. The guardrails inside the database, he notes, largely exist already: decades of access control, complete logs, and triggers that can stop a transaction in its tracks. Data-heavy agents, he suggests, belong inside the data layer as user-defined functions, if the tooling to debug and version them improves.
Stonebraker is equally blunt about the frontier. The models have already used the real training data on the internet, and he knows of no enterprise that will let a model train on its proprietary data. AGI, he said, is “basically marketing hype” from foundation model vendors justifying their valuations.
Why Postgres ? “Postgres is the intergalactic answer to open source,” Stonebraker said. Its unique feature is that “it’s not controlled by any one company.” EDB, Microsoft, Google, and others each contribute committers, and nobody steers. He dates its rise to Oracle buying MySQL, when “everybody said: vendor control.” For enterprises, that independence is the foundation of sovereignty : retaining control over how and where their data systems operate.
Momjian describes the same structure from the inside. Postgres has no office, no assets, and no legal entity, yet it produces “software better than billion-dollar companies.” Its forecasting tool is the community itself, which spots needs early from China, Russia, and the United States. “We only need one person in the world to be supersmart, and that person to tell us what they think.”
Both experts describe a design philosophy that favors flexibility over the last few percent of speed. Stonebraker built Postgres to be extended, and that is why it absorbed JSON, full-text search, GIS, and now AI without a rewrite. Momjian adds a second call from the 1990s: Lean on the operating system rather than absorb it, which made Postgres a natural fit for the cloud.
Momjian resists optimizations that buy 5 or 10 percent at the cost of rigidity, since a mistuned system can lose 40 percent. “If you take a Formula One car to the supermarket, there’s no place for the groceries.” The project plans on a horizon few vendors can match. “We’re looking at what would be the impact 5 years from now, 10 years from now.”
The new release puts that discipline on show. Several previewed features, including SQL property graph queries and partition merging and splitting, were reverted, because reliability comes first. AI-assisted review is finding problems quickly; the August 2026 patch release for PostgreSQL 18 addressed 28 CVEs.
What remains is practical: REPACK CONCURRENTLY for busy tables, smarter parallel autovacuum, and sequence synchronization for logical replication. Plan Advice, from PostgreSQL major contributor and committer Robert Haas, is the community’s answer to Oracle-style optimizer hints it rejected for years. “Somebody found a way of doing it that didn’t have the downsides,” Momjian said.
The two diverge on how fast AI will write SQL. Momjian expects agents to create databases and generate much of the code. Stonebraker says that will not happen inside the enterprise until data is documented, silos are opened, and the politics change. They agree on the destination. AI needs a relational system to give it permanence and memory, and, in Momjian’s words, that is “a game-changer.”
Forty years on, the inventor and the steward are betting on the same database.
BattMan aims to turbocharge US batteries without robbin' critical minerals
DoE-backed consortium hoping to scale, says orgs can apply to become, literally, part of BattMan
Oracle wants to bill you for every AI task, not just every seat
Big Red says AI units will add to subscriptions, as it juggles $90BN-plus capex, mounting debt and a $225BN revenue goal for 2030
HPE puts storage to work on AI’s expensive memory problem
PARTNER CONTENT: HPE Alletra MP X10000 cache offload delivers up to 20× faster time to first token in testing validated by Kamiwaza. Apparently, remembering things beats calculating them again.
Built to bend: Postgres’s inventor and a steward on what AI really needs
PARTNER CONTENT: Michael Stonebraker and Bruce Momjian on why enterprise AI needs a database that no one controls, and what PostgreSQL 19 says about the road ahead
PHP2Shell is a warning shot for vulnerability disclosure
Keeping secrets is difficult in an open world, but we've done it before and must do it again
Belgium's tax agency confirms cyberattack, pulls services offline
Officials confirm this week’s tax deadlines extended to November as a result
Anthropic now lets Claude users jump out of plane without parachute
Indian certification authority eMudhra’s private keys compromised – one of them was the number ‘4’
OpenClaw creator says he spent $10 million in tokens last month
PHP blames supply chain security biz Chainguard for putting huge chunks of the web at risk
PHP2Shell exploit bypasses protections designed to stop remote code execution
There's a third PHP2Shell bug, but it won't get an emergency fix
Oracle wants to bill you for every AI task, not just every seat
Big Red says AI units will add to subscriptions, as it juggles $90BN-plus capex, mounting debt and a $225BN revenue goal for 2030
GitHub powering up four million extra cores to stop its service going down
Re-architects site to cope with AI slop-pocalypse
Microsoft faces potential vote: Shareholders curious about its plans to tackle emissions
Greenhouse gas output has risen dramatically since the company promised to become carbon-negative by 2030
UK ministers had plenty of time for AI datacenter bosses, but none for environmental groups
Investigation finds businesses dominated govt meetings while communities were left out of decisions about AI Growth Zones
Humanity can't afford not to spend $800B on AI, Anthropic insists
It's more important than chips, 'leccy and the net, after all
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Systemd adds support for Rust
What's even better than Rust in the Linux kernel? Rust in systemd!
FreeBSD founder ports Omarchy to FreeBSD in ten days, ten releases
Daemonarchy is unaffiliated with the Linux original, built with AI help, and could complicate FreeBSD's debate over AI contributions
SparkyOS: Lightweight general-purpose Arch, on a pendrive
Interesting new project from the minds behind SparkyLinux
Red Hat helps teach about tech writing
But learns tough lessons about AI in the process
NetBSD may have found a new role: The AI-free FOSS OS
Increasing awareness and adoption, followed by increasing donations
OpenAI dots inspire open source imitators amid technical difficulties
Persistent agents appeal to AI companies, but demand isn't obvious
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.

## Original Extract

PARTNER CONTENT: Michael Stonebraker and Bruce Momjian on why enterprise AI needs a database that no one controls, and what PostgreSQL 19 says about the road ahead

Jump to main content
Search
TOPICS
Special Features
All Special Features
Cloud Infrastructure Month 2026
Modernizing Financial Services with FIS and AWS
Make it real with Capgemini and AWS
Nutanix: Scale Kubernetes. Not Chaos.
Built to bend: Postgres’s inventor and a steward on what AI really needs
PARTNER CONTENT: Michael Stonebraker and Bruce Momjian on why enterprise AI needs a database that no one controls, and what PostgreSQL 19 says about the road ahead
Michael Stonebraker introduced Postgres ® at Berkeley in 1986. Forty years later, the database he invented is a leading choice for a generation of AI projects, and he is among the sharpest skeptics of how those projects are being built. Bruce Momjian is cofounder and core team member of the PostgreSQL Global Development Group, and he works at EnterpriseDB (EDB). In upcoming episodes of AI & Data Horizons by EnterpriseDB (EDB), the Turing Award winner and MIT professor and the Postgres community advocate both make the case that the future belongs to systems that bend without breaking, and to communities that no single company controls.
Stonebraker’s first exhibit is text-to-SQL, the promise that a business user can ask a question in plain English and get an answer from the warehouse. Public benchmarks such as Spider and BIRD report accuracy above 80 percent. His team tried the same thing on MIT’s own data warehouse—1,400 tables of student, faculty, building, and salary information—using real queries from its SQL log.
“We got zero-percent accuracy. Not 80, not 70, not 40—zero,” he said. Adding agentic AI, retrieval, and prompt engineering lifted that level into the teens. Telling the model every table and join term reached the 30s. “Still not ready for prime time.”
The reasons are the stuff of every enterprise estate. There is no training data, because the warehouse is not on the public internet. The data is idiosyncratic: among other things, MIT buildings have numbers, not names. The schema has rotted into cryptic table and column names and overlapping views. And real queries are harder than benchmark ones. His team has published its own benchmark, Beaver, and says nobody has broken out of accuracy levels in the 30-percent range.
Use the model for what it’s good at
The deeper problem, Stonebraker argues, is architectural. Large language models are good with text, but “what they are no good at is doing joins of structured data,” which is most of what an enterprise holds. Working with the city of Munich’s transport department, Stonebraker’s team found that answering a single citizen complaint about crosswalk timing meant combining regulations, CAD drawings, trolley schedules, and school-zone maps. Orchestrating that with a language model at the center was, he said, “ridiculously expensive.”
The team’s answer, a system called Rubicon, uses models only on text, then casts the results into tables and hands the rest to a tabular engine. The orchestrator “looks like a query optimizer.” Stonebraker says it is cheaper by a factor of 10 and considerably more accurate than agent frameworks built around a model.
He expects agentic AI to slide from the top of the hype cycle into the valley of despair before a more balanced view emerges. Enterprises hate non-determinism: Approve a loan and reverse it on review, and “that’s called a lawsuit.” Organizations also need protection from rogue agents. The guardrails inside the database, he notes, largely exist already: decades of access control, complete logs, and triggers that can stop a transaction in its tracks. Data-heavy agents, he suggests, belong inside the data layer as user-defined functions, if the tooling to debug and version them improves.
Stonebraker is equally blunt about the frontier. The models have already used the real training data on the internet, and he knows of no enterprise that will let a model train on its proprietary data. AGI, he said, is “basically marketing hype” from foundation model vendors justifying their valuations.
Why Postgres ? “Postgres is the intergalactic answer to open source,” Stonebraker said. Its unique feature is that “it’s not controlled by any one company.” EDB, Microsoft, Google, and others each contribute committers, and nobody steers. He dates its rise to Oracle buying MySQL, when “everybody said: vendor control.” For enterprises, that independence is the foundation of sovereignty : retaining control over how and where their data systems operate.
Momjian describes the same structure from the inside. Postgres has no office, no assets, and no legal entity, yet it produces “software better than billion-dollar companies.” Its forecasting tool is the community itself, which spots needs early from China, Russia, and the United States. “We only need one person in the world to be supersmart, and that person to tell us what they think.”
Both experts describe a design philosophy that favors flexibility over the last few percent of speed. Stonebraker built Postgres to be extended, and that is why it absorbed JSON, full-text search, GIS, and now AI without a rewrite. Momjian adds a second call from the 1990s: Lean on the operating system rather than absorb it, which made Postgres a natural fit for the cloud.
Momjian resists optimizations that buy 5 or 10 percent at the cost of rigidity, since a mistuned system can lose 40 percent. “If you take a Formula One car to the supermarket, there’s no place for the groceries.” The project plans on a horizon few vendors can match. “We’re looking at what would be the impact 5 years from now, 10 years from now.”
The new release puts that discipline on show. Several previewed features, including SQL property graph queries and partition merging and splitting, were reverted, because reliability comes first. AI-assisted review is finding problems quickly; the August 2026 patch release for PostgreSQL 18 addressed 28 CVEs.
What remains is practical: REPACK CONCURRENTLY for busy tables, smarter parallel autovacuum, and sequence synchronization for logical replication. Plan Advice, from PostgreSQL major contributor and committer Robert Haas, is the community’s answer to Oracle-style optimizer hints it rejected for years. “Somebody found a way of doing it that didn’t have the downsides,” Momjian said.
The two diverge on how fast AI will write SQL. Momjian expects agents to create databases and generate much of the code. Stonebraker says that will not happen inside the enterprise until data is documented, silos are opened, and the politics change. They agree on the destination. AI needs a relational system to give it permanence and memory, and, in Momjian’s words, that is “a game-changer.”
Forty years on, the inventor and the steward are betting on the same database.
BattMan aims to turbocharge US batteries without robbin' critical minerals
DoE-backed consortium hoping to scale, says orgs can apply to become, literally, part of BattMan
Oracle wants to bill you for every AI task, not just every seat
Big Red says AI units will add to subscriptions, as it juggles $90BN-plus capex, mounting debt and a $225BN revenue goal for 2030
HPE puts storage to work on AI’s expensive memory problem
PARTNER CONTENT: HPE Alletra MP X10000 cache offload delivers up to 20× faster time to first token in testing validated by Kamiwaza. Apparently, remembering things beats calculating them again.
Built to bend: Postgres’s inventor and a steward on what AI really needs
PARTNER CONTENT: Michael Stonebraker and Bruce Momjian on why enterprise AI needs a database that no one controls, and what PostgreSQL 19 says about the road ahead
PHP2Shell is a warning shot for vulnerability disclosure
Keeping secrets is difficult in an open world, but we've done it before and must do it again
Belgium's tax agency confirms cyberattack, pulls services offline
Officials confirm this week’s tax deadlines extended to November as a result
Anthropic now lets Claude users jump out of plane without parachute
Indian certification authority eMudhra’s private keys compromised – one of them was the number ‘4’
OpenClaw creator says he spent $10 million in tokens last month
PHP blames supply chain security biz Chainguard for putting huge chunks of the web at risk
PHP2Shell exploit bypasses protections designed to stop remote code execution
There's a third PHP2Shell bug, but it won't get an emergency fix
Oracle wants to bill you for every AI task, not just every seat
Big Red says AI units will add to subscriptions, as it juggles $90BN-plus capex, mounting debt and a $225BN revenue goal for 2030
GitHub powering up four million extra cores to stop its service going down
Re-architects site to cope with AI slop-pocalypse
Microsoft faces potential vote: Shareholders curious about its plans to tackle emissions
Greenhouse gas output has risen dramatically since the company promised to become carbon-negative by 2030
UK ministers had plenty of time for AI datacenter bosses, but none for environmental groups
Investigation finds businesses dominated govt meetings while communities were left out of decisions about AI Growth Zones
Humanity can't afford not to spend $800B on AI, Anthropic insists
It's more important than chips, 'leccy and the net, after all
PLUS: US takes down Iranian propaganda sites; Marketing company asks 'Why Do We Have Your Information?' And more!
PLUS: China upgrades smartphone surveillance tools; Ring eases anti-snooping stance; and more
Voting village reports have been so successful, says Jeff Moss, that the whole of DEF CON will now be included
Went at equivalent of $3.5B+ valuation for entire firm, though portion sold not specified
On the plus side, infosec's a good bet for a long, stable career
Systemd adds support for Rust
What's even better than Rust in the Linux kernel? Rust in systemd!
FreeBSD founder ports Omarchy to FreeBSD in ten days, ten releases
Daemonarchy is unaffiliated with the Linux original, built with AI help, and could complicate FreeBSD's debate over AI contributions
SparkyOS: Lightweight general-purpose Arch, on a pendrive
Interesting new project from the minds behind SparkyLinux
Red Hat helps teach about tech writing
But learns tough lessons about AI in the process
NetBSD may have found a new role: The AI-free FOSS OS
Increasing awareness and adoption, followed by increasing donations
OpenAI dots inspire open source imitators amid technical difficulties
Persistent agents appeal to AI companies, but demand isn't obvious
Contact us
Advertise with us
Who we are
Newsletter
The Next Platform
DevClass
Blocks and Files
Situation Publishing
Cookies Policy
Privacy Policy
Ts & Cs
Do not share my personal information
Your Consent Options
Copyright. All rights reserved © 1998-2026.
