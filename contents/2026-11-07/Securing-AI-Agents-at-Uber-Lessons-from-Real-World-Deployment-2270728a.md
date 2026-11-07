---
source: "https://www.uber.com/us/en/blog/securing-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=50405330"
title: "Securing AI Agents [at Uber]: Lessons from Real-World Deployment"
article_title: "Securing AI Agents: Lessons From Real-World Deployment"
image: "https://tb-static.uber.com/prod/udam-assets/78f00b91-097e-4094-961c-04e03225132c.png"
author: "bluepeter"
captured_at: "2026-11-07T15:28:20Z"
capture_tool: "hn-digest"
hn_id: 50405330
score: 1
comments: 0
posted_at: "2026-11-07T15:23:58Z"
tags:
  - hacker-news
---

# Securing AI Agents [at Uber]: Lessons from Real-World Deployment

- HN: [50405330](https://news.ycombinator.com/item?id=50405330)
- Source: [www.uber.com](https://www.uber.com/us/en/blog/securing-ai-agents/)
- Score: 1
- Comments: 0
- Posted: 2026-11-07T15:23:58Z

## Translation

Title: Securing AI Agents [at Uber]: Lessons from Real-World Deployment
Article title: Securing AI Agents: Lessons From Real-World Deployment
Description: Uber

Article text:
Skip to main content
×
Browser not supported
This probably isn't the experience you were expecting. Internet Explorer isn't supported on Uber.com. Try switching to a different browser to view our site.
Explore Products Advertising Learn more about advertising on Uber. Reach consumers as they go anywhere and get anything.
Resources for driving and delivering with Uber
Experiences and information for people on the move
Ordering meals for delivery is just the beginning with Uber Eats
Putting stores within reach of a world of customers
Transforming the way companies move and feed their people
Moving care forward together with medical providers
Enhancing campus transportation
Expanding the reach of public transportation
The technology behind Uber Engineering
Doing the right thing for cities and communities globally
Uber news and updates in your country
Product, how-to, and policy content—and more
Securing AI Agents: Lessons From Real-World Deployment
Director of Engineering & Science
Senior Staff Software Engineer
Share this article Facebook Linkedin X social Link Introduction
Agentic AI introduces a fundamentally different security model from traditional software. Unlike conventional services that execute predefined logic, agents dynamically reason over content from users, webpages, documents, tools, and other agents—and can translate that reasoning into real-world actions. This shifts part of the attack surface from traditional identity and API controls into the semantic layer, creating risks such as prompt injection, tool and memory poisoning, goal hijacking, behavioral drift, and hallucination-driven actions. Traditional authorization can determine who is allowed to act, but securing agentic systems requires answering a broader question: is the proposed action consistent with what the user requested, within the agent’s defined persona and intents, appropriate for the user’s role, and valid given the full session context—including the actions and decisions that preceded it?
To address this challenge, we developed a defense-in-depth security framework spanning the full agent lifecycle: provisioning-time validation, runtime enforcement, and continuous assurance. The framework combines an agent registry, personas and explicit intents, constraints, tool and MCP validation, context-aware guardrails, and continuous red teaming with end-to-end observability and replay. At runtime, we introduced prompt provenance and evidence chaining to correlate privileged actions back to what the user actually requested, providing an additional defense against indirect prompt injection, goal hijacking, and unintended agent behavior.
Securing agents requires more than adding guardrails; it requires designing security into the agent architecture itself. By sharing our framework, implementation approach, tradeoffs, and lessons learned, we aim to provide engineering teams with practical patterns for building these protections into agentic systems from the outset, rather than retrofitting them after agents reach production.
Beyond Authorization: Governing Agent Behavior
We built this framework because traditional security controls weren’t designed to govern how AI agents reason and act at runtime. Identity and authorization systems can determine whether a user or service has permission to access a resource, but they can’t determine whether an agent’s proposed action is appropriate in context. As agents gained greater autonomy and access to production tools and sensitive data, we needed controls that could govern behavior, not just access, and continuously adapt as agents, capabilities, and attack techniques evolved.
Before developing the framework, our primary alternatives were traditional identity, authorization, and policy-enforcement controls supplemented by AI-specific point solutions. We evaluated several commercial AI security and guardrail platforms, including technologies designed to detect direct prompt injection, jailbreaks, unsafe content, and sensitive data exposure. We also considered traditional policy engines and authorization technologies for controlling access to tools and data. These technologies remain important parts of a defense-in-depth architecture, but none fully addressed the broader threat model we were trying to solve: an agent manipulated into an unintended action even when the user, agent, tool, and every API call are legitimately authorized.
Figure 1 : Extending access controls with behavioral enforcement to ensure AI agents act as intended.
Designing these controls for production introduced a second set of challenges. Security effectiveness had to be balanced against accuracy, latency, cost, developer experience, and agent autonomy. Stronger enforcement can reduce unsafe actions but also increase false blocks, so we moved away from uniform thresholds toward continuous, per-agent calibration based on each agent’s behavior and risk profile. Runtime evaluation also introduced latency and inference costs, leading us to establish strict latency budgets and favor small, specialized models over large general-purpose LLMs where possible.
Developer adoption was equally important. A technically strong control provides limited value if onboarding complexity or false positives encourage teams to bypass it, so we invested in automated onboarding and guardrail assignment to make the secure path the easiest path. We also found that generic controls could provide broad baseline coverage, but some of the strongest protections—particularly persona, intent, and semantic constraint enforcement—required deeper integration and understanding of each agent.
These tradeoffs shaped the framework as much as the underlying security model. Our central lesson was that security, accuracy, latency, cost, agent autonomy, and developer experience can’t be optimized independently. Effective agent guardrails require treating them as a distributed systems and engineering problem as much as a security problem, with controls designed to operate continuously and reliably at production scale.
Uber Secure AI is a defense-in-depth architecture that secures agents throughout their lifecycle across three integrated layers: provisioning-time validation, runtime enforcement, and continuous assurance. It’s built on a simple premise: traditional identity and authorization remain essential, but they’re no longer sufficient when software can reason over untrusted content, dynamically select tools, and take actions on a user's behalf. Agentic systems require an additional semantic and behavioral control layer that can determine not only whether an agent can perform an action, but whether it should perform that action given its defined purpose, the user's request and permissions, and the context of the session.
Figure 2 : A three-layer architecture securing AI agents from provisioning and runtime enforcement through continuous assurance.
The architecture begins at provisioning, where the Uber AI Registry provides the authoritative record for an agent's identity, ownership, persona, permitted intents, constraints, dependencies, and risk tier. These properties collectively establish a behavioral contract describing what the agent is designed and permitted to do.
Rather than relying entirely on developers to declare that contract manually, Argus, our automated agent security analysis platform, analyzes the agent's source code, prompts, configurations, dependencies, exposed endpoints, and tool-invocation paths. This analysis identifies security weaknesses while also deriving structured intents from the agent's implementation, allowing declared behavior to be compared with observed capabilities.
The same principle extends to the agent's supply chain. Third-party MCP servers and tools are evaluated using techniques such as static and dynamic analysis, software composition analysis, reputation signals, and semantic inspection of tool descriptions and metadata to identify vulnerabilities, indirect prompt injection, tool poisoning, and other malicious behavior. LLMs (Large Language Models) also undergo security evaluation before deployment, while monorepo scanning identifies unregistered agents, unauthorized MCP endpoints, and attempts to operate outside required security controls.
Together, these capabilities establish what the agent is, what it is intended to do, which tools and resources it may use, and the security controls that must follow it into production.
At runtime, the architecture builds on the familiar separation between policy administration, policy decision, and policy enforcement, but extends the policies from resource authorization into semantic behavior.
The AI Registry acts as the PAP (policy administration point), where agent policies, personas, intents, constraints, and required guardrails are defined and managed. AI Guard acts as the PDP (policy decision point), evaluating interactions and proposed actions against those policies. Agent Guard runtime hooks, along with enforcement logic integrated through LangFx or the AI Guard SDK, act as PEPs (policy enforcement points), invoking AI Guard at critical checkpoints and applying its verdict before an interaction or action is allowed to proceed.
Conceptually, the model extends traditional authorization, as shown in Figure 3.
Figure 3 : Extending access control to evaluate not just who and what, but why and whether an action is appropriate.
An agent may therefore possess valid credentials to invoke a tool—and the user may legitimately have access to the underlying resource—while the specific action can still be denied because it falls outside the agent's defined intent or is inconsistent with the overall context.
Because guardrail decisions occur directly in the agent’s execution path, performance is a core architectural requirement. Runtime controls must operate within a strict latency budget, making it impractical to rely on general-purpose SaaS LLMs for every security decision, where network round trips and remote inference can introduce significant and unpredictable latency.
Instead, AI Guard uses small, specialized models hosted locally, including models fine-tuned for specific guardrail functions, to keep inference close to the execution path and minimize response time. These models execute in parallel and, where appropriate, leverage in-memory vector search for low-latency semantic matching. By combining local inference with models optimized for narrowly defined security tasks, the architecture reduces latency and cost while allowing individual guardrails to be continuously calibrated to the behavior and risk profile of each agent.
Secure AI isn’t designed as a standalone security gateway that product teams must independently adopt. It’s embedded at common infrastructure points across Uber's existing AI development and engineering stack so that agents can inherit protections as they’re created, deployed, and operated.
Integrations with Uber's code and no-code agent development platforms discover and register agents early in the development lifecycle, establishing identity, ownership, persona, intents, risk, and required controls. LangFx and the AI Guard SDK provide standard integration paths for applications to invoke AI Guard, while Agent Guard provides runtime interception and enforcement at agent and tool-execution boundaries.
AI-specific supply-chain defenses integrate with existing development and DevOps workflows to identify vulnerable or malicious dependencies, MCP servers, tools, and other risks before deployment. At runtime, AI Guard becomes a shared decision service for security, privacy, and safety policies across application, model, and tool checkpoints.
This approach makes agent security an inherited platform capability rather than an implementation responsibility repeated by every product team. Security follows the agent from registration and development through model interaction, tool execution, and production operation.
The third layer of t

[truncated]

## Original Extract

Uber

Skip to main content
×
Browser not supported
This probably isn't the experience you were expecting. Internet Explorer isn't supported on Uber.com. Try switching to a different browser to view our site.
Explore Products Advertising Learn more about advertising on Uber. Reach consumers as they go anywhere and get anything.
Resources for driving and delivering with Uber
Experiences and information for people on the move
Ordering meals for delivery is just the beginning with Uber Eats
Putting stores within reach of a world of customers
Transforming the way companies move and feed their people
Moving care forward together with medical providers
Enhancing campus transportation
Expanding the reach of public transportation
The technology behind Uber Engineering
Doing the right thing for cities and communities globally
Uber news and updates in your country
Product, how-to, and policy content—and more
Securing AI Agents: Lessons From Real-World Deployment
Director of Engineering & Science
Senior Staff Software Engineer
Share this article Facebook Linkedin X social Link Introduction
Agentic AI introduces a fundamentally different security model from traditional software. Unlike conventional services that execute predefined logic, agents dynamically reason over content from users, webpages, documents, tools, and other agents—and can translate that reasoning into real-world actions. This shifts part of the attack surface from traditional identity and API controls into the semantic layer, creating risks such as prompt injection, tool and memory poisoning, goal hijacking, behavioral drift, and hallucination-driven actions. Traditional authorization can determine who is allowed to act, but securing agentic systems requires answering a broader question: is the proposed action consistent with what the user requested, within the agent’s defined persona and intents, appropriate for the user’s role, and valid given the full session context—including the actions and decisions that preceded it?
To address this challenge, we developed a defense-in-depth security framework spanning the full agent lifecycle: provisioning-time validation, runtime enforcement, and continuous assurance. The framework combines an agent registry, personas and explicit intents, constraints, tool and MCP validation, context-aware guardrails, and continuous red teaming with end-to-end observability and replay. At runtime, we introduced prompt provenance and evidence chaining to correlate privileged actions back to what the user actually requested, providing an additional defense against indirect prompt injection, goal hijacking, and unintended agent behavior.
Securing agents requires more than adding guardrails; it requires designing security into the agent architecture itself. By sharing our framework, implementation approach, tradeoffs, and lessons learned, we aim to provide engineering teams with practical patterns for building these protections into agentic systems from the outset, rather than retrofitting them after agents reach production.
Beyond Authorization: Governing Agent Behavior
We built this framework because traditional security controls weren’t designed to govern how AI agents reason and act at runtime. Identity and authorization systems can determine whether a user or service has permission to access a resource, but they can’t determine whether an agent’s proposed action is appropriate in context. As agents gained greater autonomy and access to production tools and sensitive data, we needed controls that could govern behavior, not just access, and continuously adapt as agents, capabilities, and attack techniques evolved.
Before developing the framework, our primary alternatives were traditional identity, authorization, and policy-enforcement controls supplemented by AI-specific point solutions. We evaluated several commercial AI security and guardrail platforms, including technologies designed to detect direct prompt injection, jailbreaks, unsafe content, and sensitive data exposure. We also considered traditional policy engines and authorization technologies for controlling access to tools and data. These technologies remain important parts of a defense-in-depth architecture, but none fully addressed the broader threat model we were trying to solve: an agent manipulated into an unintended action even when the user, agent, tool, and every API call are legitimately authorized.
Figure 1 : Extending access controls with behavioral enforcement to ensure AI agents act as intended.
Designing these controls for production introduced a second set of challenges. Security effectiveness had to be balanced against accuracy, latency, cost, developer experience, and agent autonomy. Stronger enforcement can reduce unsafe actions but also increase false blocks, so we moved away from uniform thresholds toward continuous, per-agent calibration based on each agent’s behavior and risk profile. Runtime evaluation also introduced latency and inference costs, leading us to establish strict latency budgets and favor small, specialized models over large general-purpose LLMs where possible.
Developer adoption was equally important. A technically strong control provides limited value if onboarding complexity or false positives encourage teams to bypass it, so we invested in automated onboarding and guardrail assignment to make the secure path the easiest path. We also found that generic controls could provide broad baseline coverage, but some of the strongest protections—particularly persona, intent, and semantic constraint enforcement—required deeper integration and understanding of each agent.
These tradeoffs shaped the framework as much as the underlying security model. Our central lesson was that security, accuracy, latency, cost, agent autonomy, and developer experience can’t be optimized independently. Effective agent guardrails require treating them as a distributed systems and engineering problem as much as a security problem, with controls designed to operate continuously and reliably at production scale.
Uber Secure AI is a defense-in-depth architecture that secures agents throughout their lifecycle across three integrated layers: provisioning-time validation, runtime enforcement, and continuous assurance. It’s built on a simple premise: traditional identity and authorization remain essential, but they’re no longer sufficient when software can reason over untrusted content, dynamically select tools, and take actions on a user's behalf. Agentic systems require an additional semantic and behavioral control layer that can determine not only whether an agent can perform an action, but whether it should perform that action given its defined purpose, the user's request and permissions, and the context of the session.
Figure 2 : A three-layer architecture securing AI agents from provisioning and runtime enforcement through continuous assurance.
The architecture begins at provisioning, where the Uber AI Registry provides the authoritative record for an agent's identity, ownership, persona, permitted intents, constraints, dependencies, and risk tier. These properties collectively establish a behavioral contract describing what the agent is designed and permitted to do.
Rather than relying entirely on developers to declare that contract manually, Argus, our automated agent security analysis platform, analyzes the agent's source code, prompts, configurations, dependencies, exposed endpoints, and tool-invocation paths. This analysis identifies security weaknesses while also deriving structured intents from the agent's implementation, allowing declared behavior to be compared with observed capabilities.
The same principle extends to the agent's supply chain. Third-party MCP servers and tools are evaluated using techniques such as static and dynamic analysis, software composition analysis, reputation signals, and semantic inspection of tool descriptions and metadata to identify vulnerabilities, indirect prompt injection, tool poisoning, and other malicious behavior. LLMs (Large Language Models) also undergo security evaluation before deployment, while monorepo scanning identifies unregistered agents, unauthorized MCP endpoints, and attempts to operate outside required security controls.
Together, these capabilities establish what the agent is, what it is intended to do, which tools and resources it may use, and the security controls that must follow it into production.
At runtime, the architecture builds on the familiar separation between policy administration, policy decision, and policy enforcement, but extends the policies from resource authorization into semantic behavior.
The AI Registry acts as the PAP (policy administration point), where agent policies, personas, intents, constraints, and required guardrails are defined and managed. AI Guard acts as the PDP (policy decision point), evaluating interactions and proposed actions against those policies. Agent Guard runtime hooks, along with enforcement logic integrated through LangFx or the AI Guard SDK, act as PEPs (policy enforcement points), invoking AI Guard at critical checkpoints and applying its verdict before an interaction or action is allowed to proceed.
Conceptually, the model extends traditional authorization, as shown in Figure 3.
Figure 3 : Extending access control to evaluate not just who and what, but why and whether an action is appropriate.
An agent may therefore possess valid credentials to invoke a tool—and the user may legitimately have access to the underlying resource—while the specific action can still be denied because it falls outside the agent's defined intent or is inconsistent with the overall context.
Because guardrail decisions occur directly in the agent’s execution path, performance is a core architectural requirement. Runtime controls must operate within a strict latency budget, making it impractical to rely on general-purpose SaaS LLMs for every security decision, where network round trips and remote inference can introduce significant and unpredictable latency.
Instead, AI Guard uses small, specialized models hosted locally, including models fine-tuned for specific guardrail functions, to keep inference close to the execution path and minimize response time. These models execute in parallel and, where appropriate, leverage in-memory vector search for low-latency semantic matching. By combining local inference with models optimized for narrowly defined security tasks, the architecture reduces latency and cost while allowing individual guardrails to be continuously calibrated to the behavior and risk profile of each agent.
Secure AI isn’t designed as a standalone security gateway that product teams must independently adopt. It’s embedded at common infrastructure points across Uber's existing AI development and engineering stack so that agents can inherit protections as they’re created, deployed, and operated.
Integrations with Uber's code and no-code agent development platforms discover and register agents early in the development lifecycle, establishing identity, ownership, persona, intents, risk, and required controls. LangFx and the AI Guard SDK provide standard integration paths for applications to invoke AI Guard, while Agent Guard provides runtime interception and enforcement at agent and tool-execution boundaries.
AI-specific supply-chain defenses integrate with existing development and DevOps workflows to identify vulnerable or malicious dependencies, MCP servers, tools, and other risks before deployment. At runtime, AI Guard becomes a shared decision service for security, privacy, and safety policies across application, model, and tool checkpoints.
This approach makes agent security an inherited platform capability rather than an implementation responsibility repeated by every product team. Security follows the agent from registration and development through model interaction, tool execution, and production operation.
The third layer of t

[truncated]
