---
source: "https://bump.sh/blog/arazzo-is-solving-the-right-problem-for-ai-tools-the-spec-just-isnt-there-yet/"
hn_url: "https://news.ycombinator.com/item?id=49711694"
title: "Arazzo is solving the right problem for AI tools. The spec just isn't there yet"
article_title: "Arazzo is solving the right problem for AI tools. The spec just isn't there yet. · Bump.sh"
image: "https://bump.sh/website/images/blog/arazzo-is-solving-the-right-problem-the-spec-isnt-there-yet.png"
author: "quaelor"
captured_at: "2026-09-15T13:11:32Z"
capture_tool: "hn-digest"
hn_id: 49711694
score: 1
comments: 0
posted_at: "2026-09-15T12:44:57Z"
tags:
  - hacker-news
---

# Arazzo is solving the right problem for AI tools. The spec just isn't there yet

- HN: [49711694](https://news.ycombinator.com/item?id=49711694)
- Source: [bump.sh](https://bump.sh/blog/arazzo-is-solving-the-right-problem-for-ai-tools-the-spec-just-isnt-there-yet/)
- Score: 1
- Comments: 0
- Posted: 2026-09-15T12:44:57Z

## Translation

Title: Arazzo is solving the right problem for AI tools. The spec just isn't there yet
Article title: Arazzo is solving the right problem for AI tools. The spec just isn't there yet. · Bump.sh
Description: From API contract to documentation portal: streamline your workflow and deliver the best API experience. Built for engineers and tech writers.

Article text:
Products
Want a personalized product tour?
API doc platform
Publish API doc portals from OpenAPI and AsyncAPI documents.
Automated changelog, versioning and governance.
Documentation that scales with your API ecosystem.
Managed MCP platform
Turn your API ecosystem into deterministic, production-ready MCP servers.
Define how agents consume your APIs, with built-in authentication and observability.
Resources
Want a personalized product tour?
Help center
How to use Bump.sh
Product updates
New features, improvements, and bug fixes
The ultimate OpenAPI guide
Master OpenAPI from the start
Arazzo complete guide
Learn Arazzo API workflow capabilities
Products
API doc platform
Publish API doc portals from OpenAPI and AsyncAPI documents.
Automated changelog, versioning and governance.
Documentation that scales with your API ecosystem.
Managed MCP platform
Turn your API ecosystem into deterministic, production-ready MCP servers.
Define how agents consume your APIs, with built-in authentication and observability.
Resources
Help center
How to use Bump.sh
Product updates
New features, improvements, and bug fixes
The ultimate OpenAPI guide
Master OpenAPI from the start
Arazzo complete guide
Learn Arazzo API workflow capabilities
Arazzo is solving the right problem for AI tools. The spec just isn't there yet.
">
MCP and the emergence of AI-orchestrated API workflows
">
My relationship with Arazzo
">
We tried to describe our API workflows using Arazzo. Here’s what fell short.
">
Conditions
">
The missing piece: take a break and ask a human
">
Arazzo’s dependency on OpenAPI: its biggest strength and weakness
">
Arazzo and Flower, at a glance
">
Let’s hope Arazzo becomes what AI needs it to be
">
MCP and the emergence of AI-orchestrated API workflows
">
My relationship with Arazzo
">
We tried to describe our API workflows using Arazzo. Here’s what fell short.
">
Conditions
">
The missing piece: take a break and ask a human
">
Arazzo’s dependency on OpenAPI: its biggest strength and weakness
">
Arazzo and Flower, at a glance
">
Let’s hope Arazzo becomes what AI needs it to be
One of the biggest struggles with AI tools is making them as deterministic and reliable as possible. We saw a surge in API calls made directly by AIs that inevitably faced the same issue. A huge part of the guess work is about what actions need to be chained in order to accomplish a real business use case, showing the need to document workflows on top of standard API (per endpoint) documentation.
That’s exactly what Arazzo is made for: pushed by the OpenAPI Initiative (the folks behind OpenAPI), this specification is made to document workflows, aka action chainings. While it struggled to get traction upon release in 2024, there’s now a clear use case around API reliability for AI tools.
In this article, I’ll talk about Arazzo’s current strengths and limits, and how our internal data model filled the gap while waiting for future versions of the specification.
MCP and the emergence of AI-orchestrated API workflows
2025 was a huge shakeup in the API world: it was now easier than ever to run API workflows by asking LLMs in natural language, or even through autonomous agents, by different means: MCP servers with one tool per endpoint, skills, CLIs, direct cURL commands, etc., while leaving the AI tool with the burden of understanding how to use the API to achieve a specific user-requested action.
Limits were quickly reached: guessing API chains led to hallucinations, unreliable successes, slowness, and heavy token costs. The need for upper-level orchestration became clear: we couldn’t leave uncertainty to AI tools. Arazzo solved some of these problems: Sebastien wrote a great article on that specific topic .
I’ve been following Arazzo development since 2023, when our focus was still on human-only documentation. We thought about supporting Arazzo to document workflows as a layer on top of our clients’ existing API documentation. We didn’t get much traction, as it’s quite natural for a developer to roam around an API doc for a while to get a good view of the API’s capabilities and chainings. For complex API ecosystems, a diagram or a global getting-started guide was often more than enough to onboard their users.
We also thought about supporting it in the API Explorer, to run real workflows instead of single API calls. But as the API Explorer is mostly used during the early discovery process (either pre-sale or in early onboarding), fiddling with API calls and knowing that these options exist in the API was enough.
It was no game changer for us, hence not something we wanted to invest heavily in.
The specification found itself truly useful when we started thinking about providing our own agentic runner, the MCP Platform.
We tried to describe our API workflows using Arazzo. Here’s what fell short.
For the MCP Platform, we needed to support a standard allowing our users to describe their workflows. That’s why we started digging into Arazzo, to understand its philosophy, see if it could fit our clients’ needs, and what would be needed to support it.
We also discussed what data model to pick to store users’ workflows. We have a dedicated internal data model for API documentation: OpenAPI and AsyncAPI documents are converted into our own data model, ensuring a cohesive experience no matter what specification was used. We ended up doing the same for workflows, as we didn’t want to depend on the direction a workflow specification like Arazzo would take and wanted to have the latitude to support another specification if needed. For that reason, we designed our own internal data model for workflows, called Flower.
MCP Platform users can either upload a Flower or an Arazzo document. Arazzo documents are then automatically converted to Flower documents to be stored. The conversion is handled using arazzo-to-flower , our open-source gem that maps Arazzo constructs to their Flower equivalents.
Quickly, while working on Arazzo support and writing our first workflows (nothing is better than eating your own dog food), we found limits in the current state of the specification, and started thinking about opening our Flower data model to our users, to be used as an alternative way of describing workflows. Check out its repository, it’s open-source! .
Almost all our users describe their workflows using Flower instead of Arazzo. Let’s talk about the main reasons.
From the start, we had difficulties understanding the success/failure mechanisms. Arazzo splits the logic into two layers that both happen to be called “criteria,” which doesn’t help readability. First, a step-level successCriteria : a list of assertions that decides whether the step itself succeeded or failed, routing you to either the onSuccess or the onFailure array. Then, inside each action of those arrays, there’s a second, separate criteria field, deciding whether that specific action should run. Same object (a list of assertions), same name, two different jobs: one gates the step, the other gates the action. Add the type field on top ( end / goto for success actions, end / retry / goto for failure actions, each behaving slightly differently) and you get a lot of spec to hold in your head for what is, in practice, “run this step, then decide what’s next.”
In Flower, we kept it simple: no two-tier criteria system, just one condition per step, determining the next action:
next : run the next defined step in the workflow,
goto : run a specific flow/step,
retry : re-run the current step,
Arazzo uses one syntax (dot notation, e.g. $steps.findProduct.outputs.pets ) to grab a value, and a different one (JSONPath) as soon as you need to filter or select from a list. Two syntaxes to learn for one spec, and the filtering power only lives in that second one. There’s no way to reuse it to shape an actual output: you can check that a condition holds, but not extract “only the matching items” as an output the rest of the workflow can use.
That’s a real limitation once workflows return anything more complex than a flat object: think an inventory check that needs “only the products still in stock.” In Flower, we kept dot notation for the simple cases, and added JMESPath on top for everything else: filtering, slicing, aggregating. So $response.body.users[?age > \ 30`].name` filters and reshapes the response in one expression that can be used in other steps or returned as an output.
The missing piece: take a break and ask a human
Sometimes, a human interaction is needed during a workflow execution, to reliably execute the next steps and provide the right output. Imagine you have a workflow that’s used to get a specific product and returns its details.
User request through an LLM: “How many microphones does the Yamaha Pacifica 212 have?”.
First step executed: find-product with query “Yamaha Pacifica 212”,
Step output: an array of 2 items: {yamaha-212-vintage,yamaha-212-pro} 😱
In that case, using workflow conditions, you can either:
Run a new step to return details of the first product,
Run a new step to return details of the two products,
Simply return product ids, and let the LLM ask the user which guitar is the one he wants information about,
Ask the user, as part of your workflow, to pick one of the two products. That’s elicitation.
I described a simple use case here, but it’s a game-changer in a lot of use cases. The alternatives all have trade-offs: risk of returning the wrong information, increasing data cost, and forcing the LLM to chain workflows, hence increasing risk of failures.
Elicitations are supported in AI-first protocols like MCP, but not in Arazzo yet. It should come soon enough, as their last release now mentions Actor-in-loop as part of their roadmap .
Arazzo’s dependency on OpenAPI: its biggest strength and weakness
Arazzo is built following the same philosophy that brought us OpenAPI, and for teams with complex API ecosystems already documenting their APIs using OpenAPI, that’s its biggest strength: reference an operationId and you get its parameters and schemas for free.
But that dependency is also its biggest weakness. The two documents version separately, so they drift: rename a parameter in your OpenAPI file, and the Arazzo workflow breaks silently. You still end up duplicating inputs, examples, and descriptions across both files. And if you don’t have an OpenAPI doc, you’re writing both API and workflow documents from scratch. For teams getting started and wanting to quickly iterate, that’s an obstacle, not a shortcut.
Flower doesn’t have that problem: no separate file to keep in sync, so no drift and no duplication. It’s also less verbose, with the conditions I described earlier being a good example.
Arazzo and Flower, at a glance
Let’s hope Arazzo becomes what AI needs it to be
Flower originally was designed only as our internal data model. Clients now use it daily because it fills many gaps left by Arazzo, like the unwanted complexity on simple use cases, and the deeper data transformation capabilities. That being said, we are rooting for a new version of Arazzo that would make Flower irrelevant: we need a strong workflow standard.
Let’s do a rematch upon Arazzo next release!
We think you might like these articles too.
A few years ago most API designers, developers, and technical writers would have had very little reason to bump into JSONPath, but its starting to get more and more relevant as more tools and standards start relying on it. So what is JSONPath, what is it used for, and how can you get up to speed with using it?
After releasing the Arazzo Cheat Sheet in January, we wanted to go further and provide a step-by-step Arazzo guide. Shout-out to Phil Sturgeon for putting it together, after his great work on the OpenAPI complete guide!
OpenAPI won't make your APIs AI-ready. But Arazzo can.
When we started the Bump.sh adventure, having API documentation was mostly something reserved for A-players like Stripe, Twilio, or SquareUp. For everyone else, it was at best a h

[truncated]

## Original Extract

From API contract to documentation portal: streamline your workflow and deliver the best API experience. Built for engineers and tech writers.

Products
Want a personalized product tour?
API doc platform
Publish API doc portals from OpenAPI and AsyncAPI documents.
Automated changelog, versioning and governance.
Documentation that scales with your API ecosystem.
Managed MCP platform
Turn your API ecosystem into deterministic, production-ready MCP servers.
Define how agents consume your APIs, with built-in authentication and observability.
Resources
Want a personalized product tour?
Help center
How to use Bump.sh
Product updates
New features, improvements, and bug fixes
The ultimate OpenAPI guide
Master OpenAPI from the start
Arazzo complete guide
Learn Arazzo API workflow capabilities
Products
API doc platform
Publish API doc portals from OpenAPI and AsyncAPI documents.
Automated changelog, versioning and governance.
Documentation that scales with your API ecosystem.
Managed MCP platform
Turn your API ecosystem into deterministic, production-ready MCP servers.
Define how agents consume your APIs, with built-in authentication and observability.
Resources
Help center
How to use Bump.sh
Product updates
New features, improvements, and bug fixes
The ultimate OpenAPI guide
Master OpenAPI from the start
Arazzo complete guide
Learn Arazzo API workflow capabilities
Arazzo is solving the right problem for AI tools. The spec just isn't there yet.
">
MCP and the emergence of AI-orchestrated API workflows
">
My relationship with Arazzo
">
We tried to describe our API workflows using Arazzo. Here’s what fell short.
">
Conditions
">
The missing piece: take a break and ask a human
">
Arazzo’s dependency on OpenAPI: its biggest strength and weakness
">
Arazzo and Flower, at a glance
">
Let’s hope Arazzo becomes what AI needs it to be
">
MCP and the emergence of AI-orchestrated API workflows
">
My relationship with Arazzo
">
We tried to describe our API workflows using Arazzo. Here’s what fell short.
">
Conditions
">
The missing piece: take a break and ask a human
">
Arazzo’s dependency on OpenAPI: its biggest strength and weakness
">
Arazzo and Flower, at a glance
">
Let’s hope Arazzo becomes what AI needs it to be
One of the biggest struggles with AI tools is making them as deterministic and reliable as possible. We saw a surge in API calls made directly by AIs that inevitably faced the same issue. A huge part of the guess work is about what actions need to be chained in order to accomplish a real business use case, showing the need to document workflows on top of standard API (per endpoint) documentation.
That’s exactly what Arazzo is made for: pushed by the OpenAPI Initiative (the folks behind OpenAPI), this specification is made to document workflows, aka action chainings. While it struggled to get traction upon release in 2024, there’s now a clear use case around API reliability for AI tools.
In this article, I’ll talk about Arazzo’s current strengths and limits, and how our internal data model filled the gap while waiting for future versions of the specification.
MCP and the emergence of AI-orchestrated API workflows
2025 was a huge shakeup in the API world: it was now easier than ever to run API workflows by asking LLMs in natural language, or even through autonomous agents, by different means: MCP servers with one tool per endpoint, skills, CLIs, direct cURL commands, etc., while leaving the AI tool with the burden of understanding how to use the API to achieve a specific user-requested action.
Limits were quickly reached: guessing API chains led to hallucinations, unreliable successes, slowness, and heavy token costs. The need for upper-level orchestration became clear: we couldn’t leave uncertainty to AI tools. Arazzo solved some of these problems: Sebastien wrote a great article on that specific topic .
I’ve been following Arazzo development since 2023, when our focus was still on human-only documentation. We thought about supporting Arazzo to document workflows as a layer on top of our clients’ existing API documentation. We didn’t get much traction, as it’s quite natural for a developer to roam around an API doc for a while to get a good view of the API’s capabilities and chainings. For complex API ecosystems, a diagram or a global getting-started guide was often more than enough to onboard their users.
We also thought about supporting it in the API Explorer, to run real workflows instead of single API calls. But as the API Explorer is mostly used during the early discovery process (either pre-sale or in early onboarding), fiddling with API calls and knowing that these options exist in the API was enough.
It was no game changer for us, hence not something we wanted to invest heavily in.
The specification found itself truly useful when we started thinking about providing our own agentic runner, the MCP Platform.
We tried to describe our API workflows using Arazzo. Here’s what fell short.
For the MCP Platform, we needed to support a standard allowing our users to describe their workflows. That’s why we started digging into Arazzo, to understand its philosophy, see if it could fit our clients’ needs, and what would be needed to support it.
We also discussed what data model to pick to store users’ workflows. We have a dedicated internal data model for API documentation: OpenAPI and AsyncAPI documents are converted into our own data model, ensuring a cohesive experience no matter what specification was used. We ended up doing the same for workflows, as we didn’t want to depend on the direction a workflow specification like Arazzo would take and wanted to have the latitude to support another specification if needed. For that reason, we designed our own internal data model for workflows, called Flower.
MCP Platform users can either upload a Flower or an Arazzo document. Arazzo documents are then automatically converted to Flower documents to be stored. The conversion is handled using arazzo-to-flower , our open-source gem that maps Arazzo constructs to their Flower equivalents.
Quickly, while working on Arazzo support and writing our first workflows (nothing is better than eating your own dog food), we found limits in the current state of the specification, and started thinking about opening our Flower data model to our users, to be used as an alternative way of describing workflows. Check out its repository, it’s open-source! .
Almost all our users describe their workflows using Flower instead of Arazzo. Let’s talk about the main reasons.
From the start, we had difficulties understanding the success/failure mechanisms. Arazzo splits the logic into two layers that both happen to be called “criteria,” which doesn’t help readability. First, a step-level successCriteria : a list of assertions that decides whether the step itself succeeded or failed, routing you to either the onSuccess or the onFailure array. Then, inside each action of those arrays, there’s a second, separate criteria field, deciding whether that specific action should run. Same object (a list of assertions), same name, two different jobs: one gates the step, the other gates the action. Add the type field on top ( end / goto for success actions, end / retry / goto for failure actions, each behaving slightly differently) and you get a lot of spec to hold in your head for what is, in practice, “run this step, then decide what’s next.”
In Flower, we kept it simple: no two-tier criteria system, just one condition per step, determining the next action:
next : run the next defined step in the workflow,
goto : run a specific flow/step,
retry : re-run the current step,
Arazzo uses one syntax (dot notation, e.g. $steps.findProduct.outputs.pets ) to grab a value, and a different one (JSONPath) as soon as you need to filter or select from a list. Two syntaxes to learn for one spec, and the filtering power only lives in that second one. There’s no way to reuse it to shape an actual output: you can check that a condition holds, but not extract “only the matching items” as an output the rest of the workflow can use.
That’s a real limitation once workflows return anything more complex than a flat object: think an inventory check that needs “only the products still in stock.” In Flower, we kept dot notation for the simple cases, and added JMESPath on top for everything else: filtering, slicing, aggregating. So $response.body.users[?age > \ 30`].name` filters and reshapes the response in one expression that can be used in other steps or returned as an output.
The missing piece: take a break and ask a human
Sometimes, a human interaction is needed during a workflow execution, to reliably execute the next steps and provide the right output. Imagine you have a workflow that’s used to get a specific product and returns its details.
User request through an LLM: “How many microphones does the Yamaha Pacifica 212 have?”.
First step executed: find-product with query “Yamaha Pacifica 212”,
Step output: an array of 2 items: {yamaha-212-vintage,yamaha-212-pro} 😱
In that case, using workflow conditions, you can either:
Run a new step to return details of the first product,
Run a new step to return details of the two products,
Simply return product ids, and let the LLM ask the user which guitar is the one he wants information about,
Ask the user, as part of your workflow, to pick one of the two products. That’s elicitation.
I described a simple use case here, but it’s a game-changer in a lot of use cases. The alternatives all have trade-offs: risk of returning the wrong information, increasing data cost, and forcing the LLM to chain workflows, hence increasing risk of failures.
Elicitations are supported in AI-first protocols like MCP, but not in Arazzo yet. It should come soon enough, as their last release now mentions Actor-in-loop as part of their roadmap .
Arazzo’s dependency on OpenAPI: its biggest strength and weakness
Arazzo is built following the same philosophy that brought us OpenAPI, and for teams with complex API ecosystems already documenting their APIs using OpenAPI, that’s its biggest strength: reference an operationId and you get its parameters and schemas for free.
But that dependency is also its biggest weakness. The two documents version separately, so they drift: rename a parameter in your OpenAPI file, and the Arazzo workflow breaks silently. You still end up duplicating inputs, examples, and descriptions across both files. And if you don’t have an OpenAPI doc, you’re writing both API and workflow documents from scratch. For teams getting started and wanting to quickly iterate, that’s an obstacle, not a shortcut.
Flower doesn’t have that problem: no separate file to keep in sync, so no drift and no duplication. It’s also less verbose, with the conditions I described earlier being a good example.
Arazzo and Flower, at a glance
Let’s hope Arazzo becomes what AI needs it to be
Flower originally was designed only as our internal data model. Clients now use it daily because it fills many gaps left by Arazzo, like the unwanted complexity on simple use cases, and the deeper data transformation capabilities. That being said, we are rooting for a new version of Arazzo that would make Flower irrelevant: we need a strong workflow standard.
Let’s do a rematch upon Arazzo next release!
We think you might like these articles too.
A few years ago most API designers, developers, and technical writers would have had very little reason to bump into JSONPath, but its starting to get more and more relevant as more tools and standards start relying on it. So what is JSONPath, what is it used for, and how can you get up to speed with using it?
After releasing the Arazzo Cheat Sheet in January, we wanted to go further and provide a step-by-step Arazzo guide. Shout-out to Phil Sturgeon for putting it together, after his great work on the OpenAPI complete guide!
OpenAPI won't make your APIs AI-ready. But Arazzo can.
When we started the Bump.sh adventure, having API documentation was mostly something reserved for A-players like Stripe, Twilio, or SquareUp. For everyone else, it was at best a h

[truncated]
