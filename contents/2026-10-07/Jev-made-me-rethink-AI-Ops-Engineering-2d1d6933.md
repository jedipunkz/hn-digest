---
source: "https://arminn.com/writing/jev-made-me-rethink-ai-ops-engineering/"
hn_url: "https://news.ycombinator.com/item?id=49991138"
title: "Jev made me rethink AI Ops Engineering"
article_title: "Jev made me rethink AI Ops Engineering — Armin Naimi"
image: ""
author: "arminn"
captured_at: "2026-10-07T11:33:19Z"
capture_tool: "hn-digest"
hn_id: 49991138
score: 1
comments: 1
posted_at: "2026-10-07T11:15:14Z"
tags:
  - hacker-news
---

# Jev made me rethink AI Ops Engineering

- HN: [49991138](https://news.ycombinator.com/item?id=49991138)
- Source: [arminn.com](https://arminn.com/writing/jev-made-me-rethink-ai-ops-engineering/)
- Score: 1
- Comments: 1
- Posted: 2026-10-07T11:15:14Z

## Translation

Title: Jev made me rethink AI Ops Engineering
Article title: Jev made me rethink AI Ops Engineering — Armin Naimi
Description: How AI Ops systems can combine code, decision models, LLMs, agents and human judgment.

Article text:
Jev made me rethink AI Ops Engineering — Armin Naimi Skip to content Armin Naimi AI Ops Engineer & Leader | Product Builder Jev made me rethink AI Ops Engineering
I’ve been thinking a lot about AI Ops Engineering lately.
Not “AI infrastructure” in the sense of hosting models.
I mean the engineering discipline around making AI actually operate inside a company.
Connecting models to business systems. Giving them the right context. Deciding what they are allowed to do. Routing work between humans, agents and software. Measuring what happens. Handling failures. And, increasingly, deciding which kind of intelligence should be used for which decision .
Jev made that last part much more concrete for me.
AI Ops is not just about agents
I think one mistake we’re making right now is treating the LLM as the center of every AI system.
We give Claude or GPT access to a bunch of tools, throw context at it and ask it to figure everything out.
But it also creates systems where the most expensive and flexible component is doing everything.
From an AI Ops perspective, that feels wrong.
The job isn’t to put an LLM everywhere.
The job is to build the right system.
Jev is interesting because it forces separation
When I first looked at Jev, what stood out was what it doesn’t do.
You give it state and ask it focused questions. It makes structured decisions that your software can use.
That sounds like a much smaller capability than an LLM.
And that is exactly why I find it interesting.
Because a huge amount of work inside an organisation is not really a generation problem.
Should we escalate this ticket?
Is this customer showing signs of churn?
Is this lead worth routing to sales?
Did something important change?
Does this require human approval?
We currently use LLMs for many of these because LLMs are good enough at them.
But “the model can do this” and “this is the right architecture” are two different questions.
An example from our own systems
We have a ticket classification workflow at work.
Claude gets a ticket, gathers context and works out what is going on.
It can determine whether it should be escalated.
It can write an escalation note.
Some of those jobs clearly need a generative model.
Classification is the obvious example.
then I don’t really need Claude to generate anything.
I need software to make a judgment.
That distinction sounds small, but I think it is fundamental to AI Ops Engineering.
The architecture starts to look different
I think we’ll increasingly build systems that look more like:
Data → decision model → deterministic workflow → LLM when needed → human when needed
Different parts of the system do different jobs.
Deterministic software handles things we can define clearly.
Models like Jev handle fast probabilistic judgments.
LLMs handle open-ended reasoning, generation and interaction.
Agents coordinate across tools and workflows.
Humans stay in the loop where judgment, responsibility or ambiguity makes that necessary.
That is AI Ops Engineering to me.
Not “how do we build the smartest agent?”
How do we design the operating system around intelligence?
The missing layer between code and LLMs
Traditional software gives us deterministic rules.
LLMs gave us something completely different.
here is a messy situation, work out what to do
There is a huge space between those two.
they opened several support tickets
somebody mentioned pricing concerns on a call
A human looks at that and develops a feeling:
This account might be at risk.
Writing deterministic rules for that gets ugly.
But calling a large LLM every time something changes also feels excessive.
This middle layer is where Jev made something click for me.
There are thousands of places inside a business where you want machine judgment , but not necessarily machine reasoning .
This matters because AI Ops happens at scale
One classification call doesn’t matter much.
Neither does one churn-risk check.
But an organisation is full of tiny decisions.
Once AI starts operating across all of these systems, the number of decisions becomes enormous.
And suddenly latency, cost, observability and reliability matter much more.
This is where AI Ops becomes an engineering problem rather than a prompting problem.
Which ones can use a smaller decision model?
Which ones should just be code?
When should we escalate to a human?
How do we measure whether the decisions are good?
What happens when confidence is low?
How do we trace why a workflow executed?
How do we improve the system over time?
The model is only one component.
Agents make this even more important
Agents are going to create more machine decisions, not fewer.
An agent working inside a company might have access to HubSpot, PostHog, support tickets, calls, internal documents and product data.
Before taking one meaningful action, it may need to make dozens of smaller decisions.
Should I retrieve more information?
Can I take this action automatically?
Which workflow should I invoke?
Today we often let the main LLM make all of those decisions.
I’m increasingly convinced that won’t be the architecture we end up with.
The main reasoning model should probably not be doing every tiny judgment inside the system.
AI Ops Engineering is partly about decomposing those decisions and putting them in the right place.
The interesting problem is orchestration
For me, this is becoming the real engineering challenge.
Not choosing Claude versus GPT versus Gemini.
Not finding one model that does everything.
But orchestrating different types of intelligence.
SQL and normal code for deterministic logic
Jev-like models for lightweight judgment
LLMs for reasoning and generation
agents for multi-step execution
humans for high-impact decisions
Then you need infrastructure connecting all of it.
That layer is what I mean when I talk about AI Ops Engineering .
Jev is one example of a bigger shift
I don’t know yet how important Jev itself will become.
I’m still early with it, and I want to test it against real workflows rather than draw conclusions from demos.
But the model itself is almost less interesting to me than the architectural idea behind it.
For the last few years, our answer to software needing intelligence has mostly been:
I think AI Ops Engineering requires a more mature answer:
What kind of intelligence does this part of the system actually need?
Sometimes that will be Claude.
Sometimes it will be an agent.
Sometimes it will be a model like Jev.
And quite often, it should still just be code.
The companies that get AI right probably won’t be the ones with the biggest prompts or the most agents.
They’ll be the ones that learn how to combine all of these pieces into reliable systems.

## Original Extract

How AI Ops systems can combine code, decision models, LLMs, agents and human judgment.

Jev made me rethink AI Ops Engineering — Armin Naimi Skip to content Armin Naimi AI Ops Engineer & Leader | Product Builder Jev made me rethink AI Ops Engineering
I’ve been thinking a lot about AI Ops Engineering lately.
Not “AI infrastructure” in the sense of hosting models.
I mean the engineering discipline around making AI actually operate inside a company.
Connecting models to business systems. Giving them the right context. Deciding what they are allowed to do. Routing work between humans, agents and software. Measuring what happens. Handling failures. And, increasingly, deciding which kind of intelligence should be used for which decision .
Jev made that last part much more concrete for me.
AI Ops is not just about agents
I think one mistake we’re making right now is treating the LLM as the center of every AI system.
We give Claude or GPT access to a bunch of tools, throw context at it and ask it to figure everything out.
But it also creates systems where the most expensive and flexible component is doing everything.
From an AI Ops perspective, that feels wrong.
The job isn’t to put an LLM everywhere.
The job is to build the right system.
Jev is interesting because it forces separation
When I first looked at Jev, what stood out was what it doesn’t do.
You give it state and ask it focused questions. It makes structured decisions that your software can use.
That sounds like a much smaller capability than an LLM.
And that is exactly why I find it interesting.
Because a huge amount of work inside an organisation is not really a generation problem.
Should we escalate this ticket?
Is this customer showing signs of churn?
Is this lead worth routing to sales?
Did something important change?
Does this require human approval?
We currently use LLMs for many of these because LLMs are good enough at them.
But “the model can do this” and “this is the right architecture” are two different questions.
An example from our own systems
We have a ticket classification workflow at work.
Claude gets a ticket, gathers context and works out what is going on.
It can determine whether it should be escalated.
It can write an escalation note.
Some of those jobs clearly need a generative model.
Classification is the obvious example.
then I don’t really need Claude to generate anything.
I need software to make a judgment.
That distinction sounds small, but I think it is fundamental to AI Ops Engineering.
The architecture starts to look different
I think we’ll increasingly build systems that look more like:
Data → decision model → deterministic workflow → LLM when needed → human when needed
Different parts of the system do different jobs.
Deterministic software handles things we can define clearly.
Models like Jev handle fast probabilistic judgments.
LLMs handle open-ended reasoning, generation and interaction.
Agents coordinate across tools and workflows.
Humans stay in the loop where judgment, responsibility or ambiguity makes that necessary.
That is AI Ops Engineering to me.
Not “how do we build the smartest agent?”
How do we design the operating system around intelligence?
The missing layer between code and LLMs
Traditional software gives us deterministic rules.
LLMs gave us something completely different.
here is a messy situation, work out what to do
There is a huge space between those two.
they opened several support tickets
somebody mentioned pricing concerns on a call
A human looks at that and develops a feeling:
This account might be at risk.
Writing deterministic rules for that gets ugly.
But calling a large LLM every time something changes also feels excessive.
This middle layer is where Jev made something click for me.
There are thousands of places inside a business where you want machine judgment , but not necessarily machine reasoning .
This matters because AI Ops happens at scale
One classification call doesn’t matter much.
Neither does one churn-risk check.
But an organisation is full of tiny decisions.
Once AI starts operating across all of these systems, the number of decisions becomes enormous.
And suddenly latency, cost, observability and reliability matter much more.
This is where AI Ops becomes an engineering problem rather than a prompting problem.
Which ones can use a smaller decision model?
Which ones should just be code?
When should we escalate to a human?
How do we measure whether the decisions are good?
What happens when confidence is low?
How do we trace why a workflow executed?
How do we improve the system over time?
The model is only one component.
Agents make this even more important
Agents are going to create more machine decisions, not fewer.
An agent working inside a company might have access to HubSpot, PostHog, support tickets, calls, internal documents and product data.
Before taking one meaningful action, it may need to make dozens of smaller decisions.
Should I retrieve more information?
Can I take this action automatically?
Which workflow should I invoke?
Today we often let the main LLM make all of those decisions.
I’m increasingly convinced that won’t be the architecture we end up with.
The main reasoning model should probably not be doing every tiny judgment inside the system.
AI Ops Engineering is partly about decomposing those decisions and putting them in the right place.
The interesting problem is orchestration
For me, this is becoming the real engineering challenge.
Not choosing Claude versus GPT versus Gemini.
Not finding one model that does everything.
But orchestrating different types of intelligence.
SQL and normal code for deterministic logic
Jev-like models for lightweight judgment
LLMs for reasoning and generation
agents for multi-step execution
humans for high-impact decisions
Then you need infrastructure connecting all of it.
That layer is what I mean when I talk about AI Ops Engineering .
Jev is one example of a bigger shift
I don’t know yet how important Jev itself will become.
I’m still early with it, and I want to test it against real workflows rather than draw conclusions from demos.
But the model itself is almost less interesting to me than the architectural idea behind it.
For the last few years, our answer to software needing intelligence has mostly been:
I think AI Ops Engineering requires a more mature answer:
What kind of intelligence does this part of the system actually need?
Sometimes that will be Claude.
Sometimes it will be an agent.
Sometimes it will be a model like Jev.
And quite often, it should still just be code.
The companies that get AI right probably won’t be the ones with the biggest prompts or the most agents.
They’ll be the ones that learn how to combine all of these pieces into reliable systems.
