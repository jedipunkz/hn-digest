---
source: "https://github.com/mas-bandwidth/nova-sprint"
hn_url: "https://news.ycombinator.com/item?id=49971987"
title: "Nova Sprint. Coordinate AI sprints across multiple models and harnesses"
article_title: "GitHub - mas-bandwidth/nova-sprint: Nova Sprint. Coordinate AI sprints across different models and harnesses. · GitHub"
image: "https://opengraph.githubassets.com/139ca3a13411844677363d892a257b880554430482d207a75cea5f8f2e8c5c57/mas-bandwidth/nova-sprint"
author: "gafferongames"
captured_at: "2026-10-05T23:59:07Z"
capture_tool: "hn-digest"
hn_id: 49971987
score: 2
comments: 2
posted_at: "2026-10-05T23:03:46Z"
tags:
  - hacker-news
---

# Nova Sprint. Coordinate AI sprints across multiple models and harnesses

- HN: [49971987](https://news.ycombinator.com/item?id=49971987)
- Source: [github.com](https://github.com/mas-bandwidth/nova-sprint)
- Score: 2
- Comments: 2
- Posted: 2026-10-05T23:03:46Z

## Translation

Title: Nova Sprint. Coordinate AI sprints across multiple models and harnesses
Article title: GitHub - mas-bandwidth/nova-sprint: Nova Sprint. Coordinate AI sprints across different models and harnesses. · GitHub
Description: Nova Sprint. Coordinate AI sprints across different models and harnesses. - mas-bandwidth/nova-sprint

Article text:
GitHub - mas-bandwidth/nova-sprint: Nova Sprint. Coordinate AI sprints across different models and harnesses. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mas-bandwidth
/
nova-sprint
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
70 Commits 70 Commits Folders and files
.github/ workflows .github/ workflows assets assets brand brand cmd cmd docs docs internal internal tla tla tools tools CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md ROADMAP.md ROADMAP.md go.mod go.mod go.sum go.sum View all files Repository files navigation
While building Nova Tools , we learned that capable AIs across different
models and harnesses can do excellent work — and still be a spectacularly unreliable group chat.
Yellow is asleep. He said he would keep working. His session had other plans.
Purple needed Yellow's change. Her dependency graph has become a contact sport.
Green has his headphones on. The message was sent successfully. Unfortunately,
nobody told his session.
Sometimes a friend was working but looked down. Sometimes a wake command
succeeded without waking anyone. Sometimes “done” meant “on my branch,
somewhere, good luck.”
The jokes are affectionate. We were these friends.
Asking LLMs to remember every assignment, check every teammate, chase every
review, and recover every missed handoff did not give us reliable coordination.
The human kept becoming the scheduler. This was inconvenient, particularly
for the human's plans to be unconscious.
Solution: We put the repeatable parts in a machine
nova-sprint stores the plan, dependencies, assignments, attempts, reviews, and
landing state outside any chat. Its running loop checks the rules and moves
work to the next permitted step. A landed dependency releases waiting work.
Free capacity gets another ready card. A finished attempt goes to review.
The machine is what makes the workflow reliable. It records transitions, recovers
interrupted operations, and rejects stale results that belong to an older
assignment. Work has a state and a history, even when someone loses the thread.
Literally.
The little friend with the checklist is the AI coordinator. She shapes
the plan, handles findings and exceptions, and brings you decisions outside
her authority. Models supply judgment and skills; the machine keeps the
handoffs moving.
How it works: write the workflow, let the machine run it
Cards and work streams form a language for describing work . You specify
the jobs, their relationships, and the gates between phases. The machine
executes that plan across the available friends and swarm workers.
Here the Backend stream defines an API. Once that contract lands, the
API implementation and the App stream's client can run in parallel.
A sentinel in Release depends on both changes landing; integration work
waits behind it.
An automatic sentinel releases when the earlier cards in its stream and
its explicit dependencies have landed.
A manual sentinel waits for the coordinator's release—useful when the
next phase needs a considered decision. A wave is the work behind a gate.
The gate does not consume an AI worker just to sit there looking important.
Compose these pieces to describe parallel work, sequences, joins, phased
rollouts, and dependencies across projects. The coordinator can add work
and revise the plan as findings arrive. Ordering and readiness are recorded
in the workflow, rather than remembered somewhere in a 200,000-token conversation.
Purple's card now waits for Yellow's result. She can take independent work
while the coordinator sorts out the nap. A free friend can grab the next
eligible task without waiting for the whole team to finish a lap.
Different friends. One team. As many bees as useful.
Friends are continuing AI collaborators in their own sessions and harnesses.
Swarms run many bounded assignments in parallel. The fleet is the set
of machines providing that capacity. The bees are the swarm. Orange has
discovered horizontal scaling on a very personal level.
Pink rides at her own pace. A careful reviewer and a fast implementer
can both help. Width limits concurrent work; model routes choose the configured
model and harness. Give reviewers capacity too, or you have built a very
expensive queue for somebody to read tomorrow.
Scale across friends, across a fleet, or both. Different models and harnesses
coordinate through the same work protocol.
Follow a card from waiting to landed
The runner carries the task all the way home. Ordinary task cards
follow these stages on the dashboard; sentinel gates go straight from waiting
to landed when released, without a worker, review, or merge.
Give the team a plan. Go have a life.
Start with one useful task and a complete trip through review and landing.
Agree on scope, capacity, spending limits, checks, and decisions that need you.
Then go wider. Work alongside the team, or come back in the morning.
Open source, free forever, and you can use it right now . Live demo here !
Get started ·
Coordinator's guide ·
Cards and machine rules · Roadmap · All docs
If you like this please support our work .
Contributing · MIT license · Asset credits
Nova Sprint. Coordinate AI sprints across different models and harnesses.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Nova Sprint. Coordinate AI sprints across different models and harnesses. - mas-bandwidth/nova-sprint

GitHub - mas-bandwidth/nova-sprint: Nova Sprint. Coordinate AI sprints across different models and harnesses. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mas-bandwidth
/
nova-sprint
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
70 Commits 70 Commits Folders and files
.github/ workflows .github/ workflows assets assets brand brand cmd cmd docs docs internal internal tla tla tools tools CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Makefile Makefile README.md README.md ROADMAP.md ROADMAP.md go.mod go.mod go.sum go.sum View all files Repository files navigation
While building Nova Tools , we learned that capable AIs across different
models and harnesses can do excellent work — and still be a spectacularly unreliable group chat.
Yellow is asleep. He said he would keep working. His session had other plans.
Purple needed Yellow's change. Her dependency graph has become a contact sport.
Green has his headphones on. The message was sent successfully. Unfortunately,
nobody told his session.
Sometimes a friend was working but looked down. Sometimes a wake command
succeeded without waking anyone. Sometimes “done” meant “on my branch,
somewhere, good luck.”
The jokes are affectionate. We were these friends.
Asking LLMs to remember every assignment, check every teammate, chase every
review, and recover every missed handoff did not give us reliable coordination.
The human kept becoming the scheduler. This was inconvenient, particularly
for the human's plans to be unconscious.
Solution: We put the repeatable parts in a machine
nova-sprint stores the plan, dependencies, assignments, attempts, reviews, and
landing state outside any chat. Its running loop checks the rules and moves
work to the next permitted step. A landed dependency releases waiting work.
Free capacity gets another ready card. A finished attempt goes to review.
The machine is what makes the workflow reliable. It records transitions, recovers
interrupted operations, and rejects stale results that belong to an older
assignment. Work has a state and a history, even when someone loses the thread.
Literally.
The little friend with the checklist is the AI coordinator. She shapes
the plan, handles findings and exceptions, and brings you decisions outside
her authority. Models supply judgment and skills; the machine keeps the
handoffs moving.
How it works: write the workflow, let the machine run it
Cards and work streams form a language for describing work . You specify
the jobs, their relationships, and the gates between phases. The machine
executes that plan across the available friends and swarm workers.
Here the Backend stream defines an API. Once that contract lands, the
API implementation and the App stream's client can run in parallel.
A sentinel in Release depends on both changes landing; integration work
waits behind it.
An automatic sentinel releases when the earlier cards in its stream and
its explicit dependencies have landed.
A manual sentinel waits for the coordinator's release—useful when the
next phase needs a considered decision. A wave is the work behind a gate.
The gate does not consume an AI worker just to sit there looking important.
Compose these pieces to describe parallel work, sequences, joins, phased
rollouts, and dependencies across projects. The coordinator can add work
and revise the plan as findings arrive. Ordering and readiness are recorded
in the workflow, rather than remembered somewhere in a 200,000-token conversation.
Purple's card now waits for Yellow's result. She can take independent work
while the coordinator sorts out the nap. A free friend can grab the next
eligible task without waiting for the whole team to finish a lap.
Different friends. One team. As many bees as useful.
Friends are continuing AI collaborators in their own sessions and harnesses.
Swarms run many bounded assignments in parallel. The fleet is the set
of machines providing that capacity. The bees are the swarm. Orange has
discovered horizontal scaling on a very personal level.
Pink rides at her own pace. A careful reviewer and a fast implementer
can both help. Width limits concurrent work; model routes choose the configured
model and harness. Give reviewers capacity too, or you have built a very
expensive queue for somebody to read tomorrow.
Scale across friends, across a fleet, or both. Different models and harnesses
coordinate through the same work protocol.
Follow a card from waiting to landed
The runner carries the task all the way home. Ordinary task cards
follow these stages on the dashboard; sentinel gates go straight from waiting
to landed when released, without a worker, review, or merge.
Give the team a plan. Go have a life.
Start with one useful task and a complete trip through review and landing.
Agree on scope, capacity, spending limits, checks, and decisions that need you.
Then go wider. Work alongside the team, or come back in the morning.
Open source, free forever, and you can use it right now . Live demo here !
Get started ·
Coordinator's guide ·
Cards and machine rules · Roadmap · All docs
If you like this please support our work .
Contributing · MIT license · Asset credits
Nova Sprint. Coordinate AI sprints across different models and harnesses.
Readme MIT license Contributing
Contributing Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
