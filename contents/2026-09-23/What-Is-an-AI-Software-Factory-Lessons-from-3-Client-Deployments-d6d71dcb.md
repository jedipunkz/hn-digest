---
source: "https://camplight.net/ai/ai-software-factory/"
hn_url: "https://news.ycombinator.com/item?id=49813704"
title: "What Is an AI Software Factory? Lessons from 3 Client Deployments"
article_title: "Best AI Software Factory Platforms: 7 Features to Compare"
image: "https://camplight.net/wp-content/uploads/2026/09/best-ai-software-factory-platform-nest.webp"
author: "altras"
captured_at: "2026-09-23T10:53:52Z"
capture_tool: "hn-digest"
hn_id: 49813704
score: 4
comments: 0
posted_at: "2026-09-23T09:57:06Z"
tags:
  - hacker-news
---

# What Is an AI Software Factory? Lessons from 3 Client Deployments

- HN: [49813704](https://news.ycombinator.com/item?id=49813704)
- Source: [camplight.net](https://camplight.net/ai/ai-software-factory/)
- Score: 4
- Comments: 0
- Posted: 2026-09-23T09:57:06Z

## Translation

Title: What Is an AI Software Factory? Lessons from 3 Client Deployments
Article title: Best AI Software Factory Platforms: 7 Features to Compare
Description: What makes the best AI software factories work? Compare seven essential capabilities, informed by Camplight's experience across three client deployments.

Article text:
Don’t get fooled by dark factory vendors — live session, Sept 29.
Save my seat
×
Company
Close Company
Open Company
Company
Camplight is your unfair advantage in digital product innovation.
We’re a Tech Partner and Venture Agency for B2B established (often non-tech) businesses that want to build and accelerate digital products with startup speed and enterprise standards.
Discover our mission to evolve human collaboration through digital craftsmanship, shared ownership, and long-term impact.
Explore our transparent, participatory, and radically human way of working.
Join a team of self-managed visionaries building digital products that matter. Together.
Services
Close Services
Open Services
Services
Digital Innovation needs more than execution. It needs a true partner.
From accelerating digital product delivery to launching new ventures, our services are built to reduce risk and unlock momentum. Fast.
When your B2B product stalls, we identify what’s blocking growth and fix it. Validation-first approach ensures you scale, not patch.
Most new products fail because teams build before they validate. We test market demand first, then build what sells.
Looking for a tech co-founder with real skin in the game? Let’s build it together, from idea to traction.
Industries
Close Industries
Open Industries
Industries
Innovation looks different in every industry and that’s exactly why we don’t stay in one lane.
We bring patterns, insights, and strategies from across sectors to help you build smarter in your own.
Designing financial tools with trust, precision, and user-centricity.
Reimagining how people learn, grow, and teach through tech.
Human-first applications of AI. Smart, not hype.
Discover
Close Discover
Open Discover
Discover
Digital innovation is complex. We make it more transparent.
Through open playbooks, experiments, and lessons from the field, we help established B2B companies move faster with less guesswork
Stories, lessons, and honest takes from the edge of digital product, tech innovation, and venture building.
Tools, templates, and guides from our playbooks. Made to help you move your digital innovation smarter, faster and with less risk.
Where we connect, learn, and share across ecosystems and expert circles.
Services
Accelerate Existing Product​
What Is an AI Software Factory? Lessons From 3 Client Deployments
I’m writing this in bed after a long day of trying to explain something that seems obvious, but every time I articulate it, it feels vague. I think the current blog article the best one on explaining what a software factory is! But you tell me…
An AI software factory is a managed system (like cloud servers are managed hardware) for turning defined requirements into validated software changes using AI agents, connected tools, automated checks, and human oversight.
If you are comparing the best AI software factory platforms, do not start with the number of agents or the interface. Start with whether the system connects intent, context, execution, verification, approval, release, and feedback.
Jeez, that was a lot… Let’s hold our breath for a little bit more.
It includes the agent workflow: how work starts, what context is available, what actions are permitted, and how results are accepted.
At Camplight, we have spent the last nine months building software factories and have deployed three for clients .
Those implementations are covered by NDAs. We cannot publish the client systems yet.
Instead, we have generalized our implementation experience into Nest : a public model that makes the operating questions and building blocks visible without revealing confidential deployments. It’s based on our open-source OrgOps infra .
Nest visualizes Camplight’s model for a leading human-led AI software factory, based on patterns from three client deployments.
This is not a prediction about something we might eventually build. It is a way to explain work we are already doing.
The screenshots below present a unified, generalized model rather than any of the three confidential client environments. They do not imply that every deployment contains every interface shown. I want to add a disclaimer that the names, budgets, timings, and performance figures are illustrative, not published client results.
When comparing leading AI software factory platforms, treat these screens as an evaluation framework rather than a universal product checklist. I guess I will have to write another blog article in couple of months because the space is moving so quick…
The useful question is not whether your company needs an interface identical to Nest.
It is whether you can answer the operational questions behind it.
This practical guide builds on our analysis of Uber’s AI software factory and our examination of dark-factory readiness . Here, we focus on what the system contains, how work moves through it, and what leaders need to govern.
What makes an AI software factory different from a coding agent?
A coding agent can perform substantial development work. For example, GitHub’s Copilot cloud agent can investigate a repository, plan changes, modify code, and run tests in a development environment. A software factory therefore cannot be distinguished simply by saying, “Our AI does more than autocomplete.”
The distinction is the system around the execution .
In the model we use, a software factory connects a business request to the context, tools, people, validation, and release process needed to deliver it.
Defined intent -> relevant context -> execution -> verification -> required approval -> release -> feedback
Different parts of that flow can use different mechanisms. Some require an agent. Others are better handled by ordinary code, an existing pipeline, or a person. (hint: usually the bottle neck is around verification because evals and guardrails are a moving target)
An AI software factory also does not replace continuous integration and delivery. CI/CD already provides mechanisms for building, testing, and deploying changes; the factory needs to connect agent-generated work to those mechanisms.
A coding agent performs work. A software factory defines how that work becomes an accepted, accountable software change that anyone can trigger.
The important bit here is “anyone” but we’ll come back to this later.
The term itself is used in different ways. Cortex describes an organizational software-delivery system , while StrongDM describes a deliberately non-interactive approach without human code review. Our scope is a human-led AI software factory , with explicit decisions about autonomy, verification, and intervention. This is because nobody has reached full dark state
What the best AI software factories have in common: seven building blocks
We organized Nest around seven areas that leaders can use to compare AI software factory platforms: operational visibility, agent management, configuration, projects, team assembly, reusable capabilities, and human collaboration.
These are evaluation criteria for a leading AI software factory, not a requirement to build seven new applications.
1. An operational dashboard: what is happening, and what needs attention?
The best AI software factory platforms turn spend, tasks, issues, projects, and attention points into one operational view.
The dashboard is where a person should be able to understand the state of the factory without reconstructing it from conversations. I loved my millennial days with mIRC but chatting is super tiring.
What is running? What has finished? What is waiting for review? Which issues need intervention? What has execution cost? asl pls?
Nest brings those questions together through Spend, Tasks, Issues, Active Projects, and Community .
The important design choice is to connect visibility to action! A blocked task should lead to its context. A spending anomaly should lead to the relevant workflow. A review request should lead to the artifact and the criteria for accepting it.
For an executive, I would also distinguish activity metrics from delivery metrics .
“Agents completed 100 tasks” describes activity. It does not establish that the company received 100 useful outcomes.
My preferred evaluation would combine accepted changes, elapsed delivery time, review effort, rework, and execution cost. For comparable work, one useful measure is:
Cost per accepted change = total execution and review cost for a defined batch of work / accepted changes in that batch
That calculation should include failed attempts and retries, not only the successful final run.
The dashboard also includes small celebrations and shared work from the community. That is deliberate. I want this environment to show what people are accomplishing together, not just a growing queue of machine activity.
The management question: Can someone see where their attention would be most useful right now?
2. Agent management: organize capabilities around the work
Leading AI software factories make each agent’s role, activity, performance, and status visible.
An agent directory should do more than list names and avatars. It should make responsibilities understandable.
What does each agent do? Where does it operate? Who owns its configuration? Is it active, paused, or waiting for help?
This is where the question of horizontal versus vertical agents becomes useful, provided we define it.
A horizontal capability might support several teams, such as a reusable research workflow. A domain-specific capability might operate within one product, repository, or business process.
These are design choices, not competing religions.
A shared capability can reduce duplicated work. A capability close to one domain can have a clearer context and tighter boundaries. The right choice depends on the work and the organization.
Team Topologies provides a useful reference for thinking about team boundaries, shared platforms, and interaction modes. Applying those ideas to agents is an architectural analogy, not a claim that its human team types map directly onto AI roles.
I would not begin by recreating the company’s org chart as a collection of bots.
I would begin with an outcome, identify the capabilities it requires, and decide where those capabilities should live.
The management question: Does each agent have a clear responsibility and a human owner?
3. Agent configuration: instructions are only part of the system
A mature AI software factory connects instructions to tools, skills, models, workspaces, and access boundaries.
The Agent Details screen makes an important distinction visible: an agent is more than its prompt.
In Nest, its configuration includes a description of its responsibility, model selection, tool connections, skills, workspace, and additional instructions.
Consider an implementation agent. “Write good code” is not a sufficient operating instruction.
It needs a defined change, relevant repository context, acceptance criteria, access to the appropriate tools, and an expected output. That output might be a proposed code change accompanied by test results, rather than permission to release directly.
There is also a boundary that deserves particular attention:
Instructions describe what an agent should do. Authorization controls determine what it can do.
“Do not touch production” should not be the only thing preventing production access. Permissions need enforcement outside the model, with access scoped to the task and high-impact actions subject to appropriate approval. OWASP’s agent-security guidance explicitly addresses least privilege, tool authorization, and human approval controls.
That is why configuration is an operating concern, not just a prompt-writing exercise.
The management question: What can this agent access and change, and who is accountable for that decision?
4. Projects: keep the context attached to the outcome
Top platforms keep requests, evidence, human decisions, artifacts, budgets, and agent activity in one traceable context.
The Projects screen shows how the pieces come t

[truncated]

## Original Extract

What makes the best AI software factories work? Compare seven essential capabilities, informed by Camplight's experience across three client deployments.

Don’t get fooled by dark factory vendors — live session, Sept 29.
Save my seat
×
Company
Close Company
Open Company
Company
Camplight is your unfair advantage in digital product innovation.
We’re a Tech Partner and Venture Agency for B2B established (often non-tech) businesses that want to build and accelerate digital products with startup speed and enterprise standards.
Discover our mission to evolve human collaboration through digital craftsmanship, shared ownership, and long-term impact.
Explore our transparent, participatory, and radically human way of working.
Join a team of self-managed visionaries building digital products that matter. Together.
Services
Close Services
Open Services
Services
Digital Innovation needs more than execution. It needs a true partner.
From accelerating digital product delivery to launching new ventures, our services are built to reduce risk and unlock momentum. Fast.
When your B2B product stalls, we identify what’s blocking growth and fix it. Validation-first approach ensures you scale, not patch.
Most new products fail because teams build before they validate. We test market demand first, then build what sells.
Looking for a tech co-founder with real skin in the game? Let’s build it together, from idea to traction.
Industries
Close Industries
Open Industries
Industries
Innovation looks different in every industry and that’s exactly why we don’t stay in one lane.
We bring patterns, insights, and strategies from across sectors to help you build smarter in your own.
Designing financial tools with trust, precision, and user-centricity.
Reimagining how people learn, grow, and teach through tech.
Human-first applications of AI. Smart, not hype.
Discover
Close Discover
Open Discover
Discover
Digital innovation is complex. We make it more transparent.
Through open playbooks, experiments, and lessons from the field, we help established B2B companies move faster with less guesswork
Stories, lessons, and honest takes from the edge of digital product, tech innovation, and venture building.
Tools, templates, and guides from our playbooks. Made to help you move your digital innovation smarter, faster and with less risk.
Where we connect, learn, and share across ecosystems and expert circles.
Services
Accelerate Existing Product​
What Is an AI Software Factory? Lessons From 3 Client Deployments
I’m writing this in bed after a long day of trying to explain something that seems obvious, but every time I articulate it, it feels vague. I think the current blog article the best one on explaining what a software factory is! But you tell me…
An AI software factory is a managed system (like cloud servers are managed hardware) for turning defined requirements into validated software changes using AI agents, connected tools, automated checks, and human oversight.
If you are comparing the best AI software factory platforms, do not start with the number of agents or the interface. Start with whether the system connects intent, context, execution, verification, approval, release, and feedback.
Jeez, that was a lot… Let’s hold our breath for a little bit more.
It includes the agent workflow: how work starts, what context is available, what actions are permitted, and how results are accepted.
At Camplight, we have spent the last nine months building software factories and have deployed three for clients .
Those implementations are covered by NDAs. We cannot publish the client systems yet.
Instead, we have generalized our implementation experience into Nest : a public model that makes the operating questions and building blocks visible without revealing confidential deployments. It’s based on our open-source OrgOps infra .
Nest visualizes Camplight’s model for a leading human-led AI software factory, based on patterns from three client deployments.
This is not a prediction about something we might eventually build. It is a way to explain work we are already doing.
The screenshots below present a unified, generalized model rather than any of the three confidential client environments. They do not imply that every deployment contains every interface shown. I want to add a disclaimer that the names, budgets, timings, and performance figures are illustrative, not published client results.
When comparing leading AI software factory platforms, treat these screens as an evaluation framework rather than a universal product checklist. I guess I will have to write another blog article in couple of months because the space is moving so quick…
The useful question is not whether your company needs an interface identical to Nest.
It is whether you can answer the operational questions behind it.
This practical guide builds on our analysis of Uber’s AI software factory and our examination of dark-factory readiness . Here, we focus on what the system contains, how work moves through it, and what leaders need to govern.
What makes an AI software factory different from a coding agent?
A coding agent can perform substantial development work. For example, GitHub’s Copilot cloud agent can investigate a repository, plan changes, modify code, and run tests in a development environment. A software factory therefore cannot be distinguished simply by saying, “Our AI does more than autocomplete.”
The distinction is the system around the execution .
In the model we use, a software factory connects a business request to the context, tools, people, validation, and release process needed to deliver it.
Defined intent -> relevant context -> execution -> verification -> required approval -> release -> feedback
Different parts of that flow can use different mechanisms. Some require an agent. Others are better handled by ordinary code, an existing pipeline, or a person. (hint: usually the bottle neck is around verification because evals and guardrails are a moving target)
An AI software factory also does not replace continuous integration and delivery. CI/CD already provides mechanisms for building, testing, and deploying changes; the factory needs to connect agent-generated work to those mechanisms.
A coding agent performs work. A software factory defines how that work becomes an accepted, accountable software change that anyone can trigger.
The important bit here is “anyone” but we’ll come back to this later.
The term itself is used in different ways. Cortex describes an organizational software-delivery system , while StrongDM describes a deliberately non-interactive approach without human code review. Our scope is a human-led AI software factory , with explicit decisions about autonomy, verification, and intervention. This is because nobody has reached full dark state
What the best AI software factories have in common: seven building blocks
We organized Nest around seven areas that leaders can use to compare AI software factory platforms: operational visibility, agent management, configuration, projects, team assembly, reusable capabilities, and human collaboration.
These are evaluation criteria for a leading AI software factory, not a requirement to build seven new applications.
1. An operational dashboard: what is happening, and what needs attention?
The best AI software factory platforms turn spend, tasks, issues, projects, and attention points into one operational view.
The dashboard is where a person should be able to understand the state of the factory without reconstructing it from conversations. I loved my millennial days with mIRC but chatting is super tiring.
What is running? What has finished? What is waiting for review? Which issues need intervention? What has execution cost? asl pls?
Nest brings those questions together through Spend, Tasks, Issues, Active Projects, and Community .
The important design choice is to connect visibility to action! A blocked task should lead to its context. A spending anomaly should lead to the relevant workflow. A review request should lead to the artifact and the criteria for accepting it.
For an executive, I would also distinguish activity metrics from delivery metrics .
“Agents completed 100 tasks” describes activity. It does not establish that the company received 100 useful outcomes.
My preferred evaluation would combine accepted changes, elapsed delivery time, review effort, rework, and execution cost. For comparable work, one useful measure is:
Cost per accepted change = total execution and review cost for a defined batch of work / accepted changes in that batch
That calculation should include failed attempts and retries, not only the successful final run.
The dashboard also includes small celebrations and shared work from the community. That is deliberate. I want this environment to show what people are accomplishing together, not just a growing queue of machine activity.
The management question: Can someone see where their attention would be most useful right now?
2. Agent management: organize capabilities around the work
Leading AI software factories make each agent’s role, activity, performance, and status visible.
An agent directory should do more than list names and avatars. It should make responsibilities understandable.
What does each agent do? Where does it operate? Who owns its configuration? Is it active, paused, or waiting for help?
This is where the question of horizontal versus vertical agents becomes useful, provided we define it.
A horizontal capability might support several teams, such as a reusable research workflow. A domain-specific capability might operate within one product, repository, or business process.
These are design choices, not competing religions.
A shared capability can reduce duplicated work. A capability close to one domain can have a clearer context and tighter boundaries. The right choice depends on the work and the organization.
Team Topologies provides a useful reference for thinking about team boundaries, shared platforms, and interaction modes. Applying those ideas to agents is an architectural analogy, not a claim that its human team types map directly onto AI roles.
I would not begin by recreating the company’s org chart as a collection of bots.
I would begin with an outcome, identify the capabilities it requires, and decide where those capabilities should live.
The management question: Does each agent have a clear responsibility and a human owner?
3. Agent configuration: instructions are only part of the system
A mature AI software factory connects instructions to tools, skills, models, workspaces, and access boundaries.
The Agent Details screen makes an important distinction visible: an agent is more than its prompt.
In Nest, its configuration includes a description of its responsibility, model selection, tool connections, skills, workspace, and additional instructions.
Consider an implementation agent. “Write good code” is not a sufficient operating instruction.
It needs a defined change, relevant repository context, acceptance criteria, access to the appropriate tools, and an expected output. That output might be a proposed code change accompanied by test results, rather than permission to release directly.
There is also a boundary that deserves particular attention:
Instructions describe what an agent should do. Authorization controls determine what it can do.
“Do not touch production” should not be the only thing preventing production access. Permissions need enforcement outside the model, with access scoped to the task and high-impact actions subject to appropriate approval. OWASP’s agent-security guidance explicitly addresses least privilege, tool authorization, and human approval controls.
That is why configuration is an operating concern, not just a prompt-writing exercise.
The management question: What can this agent access and change, and who is accountable for that decision?
4. Projects: keep the context attached to the outcome
Top platforms keep requests, evidence, human decisions, artifacts, budgets, and agent activity in one traceable context.
The Projects screen shows how the pieces come t

[truncated]
