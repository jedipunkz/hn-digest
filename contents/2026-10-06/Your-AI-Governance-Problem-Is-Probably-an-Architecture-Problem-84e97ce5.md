---
source: "https://www.jamiecadvisory.com/perspectives/your-ai-governance-problem-is-probably-an-architecture-problem/"
hn_url: "https://news.ycombinator.com/item?id=49978556"
title: "Your AI Governance Problem Is Probably an Architecture Problem"
article_title: "Your AI Governance Problem Is Probably an Architecture Problem | Jamie Advisory"
image: "https://www.jamiecadvisory.com/assets/jamie-c-advisory-social.png"
author: "mooreds"
captured_at: "2026-10-06T13:57:41Z"
capture_tool: "hn-digest"
hn_id: 49978556
score: 1
comments: 0
posted_at: "2026-10-06T13:55:11Z"
tags:
  - hacker-news
---

# Your AI Governance Problem Is Probably an Architecture Problem

- HN: [49978556](https://news.ycombinator.com/item?id=49978556)
- Source: [www.jamiecadvisory.com](https://www.jamiecadvisory.com/perspectives/your-ai-governance-problem-is-probably-an-architecture-problem/)
- Score: 1
- Comments: 0
- Posted: 2026-10-06T13:55:11Z

## Translation

Title: Your AI Governance Problem Is Probably an Architecture Problem
Article title: Your AI Governance Problem Is Probably an Architecture Problem | Jamie Advisory
Description: Policy describes the organisation you would like to have. Architecture describes the organisation you actually built.

Article text:
Skip to main content
Jamie C
Advisory
Menu
Home
Perspectives
Profiles
Contact
Perspective 010
Your AI Governance Problem Is Probably an Architecture Problem
Policy describes the organisation you would like to have. Architecture describes the organisation you actually built.
AI governance has become an industry remarkably quickly. There are frameworks, policies, principles, committees, risk classifications, acceptable-use standards and increasingly elaborate approval processes. Organisations are defining which models employees can use, what information can be submitted to them and which use cases require additional scrutiny. Regulators are adding another layer, and vendors are helpfully selling products designed to manage all of it.
Much of this is necessary. But underneath it sits a fairly fundamental problem: you cannot govern what you cannot see, and many organisations still cannot describe the architecture through which their AI actually operates.
They know which products they have approved. They may know which models sit behind those products. They probably have policies describing what employees should and should not do. What they often cannot describe with the same confidence is what happens to information once somebody starts using them.
A policy can say that confidential information must not be submitted to an external model, but that only establishes the expected behaviour. The architecture determines whether the organisation can enforce that expectation, detect when it is breached and understand what happened afterwards.
It determines whether prompts leave the enterprise boundary, where they are processed, whether they are retained, which systems an AI service can reach, what information can be retrieved and where outputs are stored. It also determines whether any of those interactions can be reconstructed later if something goes wrong.
These are not primarily policy questions. They are architecture questions, and the distinction matters because organisations have spent decades discovering that writing something into a security policy does not magically cause the underlying technology to obey it. AI has apparently given us an opportunity to learn this lesson again, only faster and with considerably larger budgets.
Governance without architectural control therefore relies heavily on people consistently doing what the policy says. That is important, but it has never been a particularly robust control model on its own. If the only thing preventing sensitive information crossing a boundary is an employee remembering that paragraph twelve of the acceptable-use policy says they should not do it, the organisation does not really have a control. It has a request.
The boundary is becoming harder to see
Traditional enterprise systems generally gave us reasonably understandable boundaries. There was an application, some infrastructure, a collection of interfaces and identifiable stores of data. Information moved between systems through routes that architects and security teams could map, inspect and spend several months arguing about.
AI complicates that picture because the model is increasingly only one component in a much larger chain. A single business process might involve an employee application, an enterprise AI platform, a retrieval layer, several internal data sources, an external foundation model and tools capable of taking actions in other systems. Some of those components may be operated by the enterprise, some by technology providers and some by providers used by those providers.
Agents make the problem more interesting again. The system is no longer necessarily responding to a prompt and returning some text. It may decide which information to retrieve, which tools to invoke and which systems to interact with. As organisations give AI greater access to enterprise systems, the question of what the model is becomes progressively less useful than understanding what the overall system is capable of doing.
This changes the governance question. Asking which models the organisation is using is still useful, but it is nowhere near enough. The more important questions are what those models can see, what they can do, which boundaries they can cross and where the information involved actually goes.
Those questions require an architectural answer.
Sovereignty is about control, not geography
This is also why discussions about AI sovereignty can become unnecessarily simplistic. The argument is sometimes reduced to internally hosted models being safe and external models being dangerous, as though putting something inside your own data centre automatically confers wisdom upon everyone involved.
An internally hosted model with unrestricted access to sensitive information, weak identity controls and inadequate logging may represent a substantial risk. An externally hosted model operating behind strong contractual, technical and data controls may represent a much smaller one. Location matters, particularly where regulation and jurisdiction are involved, but it is only one part of the problem.
The more useful question is where control resides. Who controls the infrastructure and the data? Who determines identity and access? Where is information processed and retained? Which jurisdiction applies? Which other organisations participate in the service? Can the enterprise inspect what happened? Can it change provider without rebuilding everything surrounding the model?
Seen this way, sovereignty is not simply a hosting decision. It is the degree to which an organisation retains meaningful control over its technology, information and future choices. That needs to be designed into the architecture rather than appended to a governance checklist after procurement has already signed the contract.
There is another consequence of treating AI governance primarily as policy. If an organisation cannot reconstruct how an AI-enabled system reached a decision or performed an action, its governance framework eventually depends on trust rather than evidence.
That becomes increasingly important as AI moves from assisting people to participating directly in business processes. An organisation does not necessarily need to preserve every token involved in every interaction forever, but it should understand what evidence is required for the level of risk involved. For important processes, that may include which model was used, what information was available to it, which sources were retrieved, which tools it could access, what actions were attempted and which controls were applied.
Most of those records will never be examined. That is not the point. We do not build auditability because we expect every transaction to require an investigation; we build it because controls that cannot be demonstrated become very difficult to rely upon when something consequential happens.
Without that evidence, an AI governance committee risks becoming a group of senior people periodically confirming that everyone still agrees with the policy. Enterprises already have enough committees capable of doing that.
None of this makes policy, governance frameworks or responsible AI principles unnecessary. It makes them dependent on something more fundamental. Before creating another approval process, organisations should be able to map their AI estate with enough precision to understand which models and platforms are being used, which information they can reach, which trust boundaries they cross, which external organisations are involved, which actions they can perform and what controls exist at each boundary.
That map will probably be uncomfortable. It may reveal duplicate platforms, uncontrolled experimentation, overlapping vendor relationships and AI capabilities embedded inside products nobody originally thought of as AI systems. It may also reveal that some carefully written governance requirements cannot actually be enforced by the architecture currently underneath them.
Good governance should expose those things, because the objective is not to prevent organisations from using AI. Organisations that understand their boundaries should be able to move faster: they know where experimentation is safe, where additional controls are necessary and where the potential consequences justify greater scrutiny.
The alternative is governance by prohibition. When nobody can confidently explain what is happening, restricting everything starts to look rational. That protects the organisation from some risks, but it also protects it remarkably effectively from innovation.
AI governance therefore cannot belong solely to risk, security, legal or technology. Each has part of the answer, and each sees a different part of the risk. But underneath all of them sits the architecture that determines what the organisation can actually control.
Policy describes the organisation you would like to have. Architecture describes the organisation you actually built.
Interested in discussing technology leadership, AI governance or board advisory work?
Read the Executive Profile , explore the Board Profile , return to the Perspectives Library or get in touch .
Previous: Perspective 009 — The Handoffs Are Where Transformation Goes to Die
Clear judgement for consequential technology decisions.
Privacy: this site uses no cookies. Anonymous analytics help improve the site.

## Original Extract

Policy describes the organisation you would like to have. Architecture describes the organisation you actually built.

Skip to main content
Jamie C
Advisory
Menu
Home
Perspectives
Profiles
Contact
Perspective 010
Your AI Governance Problem Is Probably an Architecture Problem
Policy describes the organisation you would like to have. Architecture describes the organisation you actually built.
AI governance has become an industry remarkably quickly. There are frameworks, policies, principles, committees, risk classifications, acceptable-use standards and increasingly elaborate approval processes. Organisations are defining which models employees can use, what information can be submitted to them and which use cases require additional scrutiny. Regulators are adding another layer, and vendors are helpfully selling products designed to manage all of it.
Much of this is necessary. But underneath it sits a fairly fundamental problem: you cannot govern what you cannot see, and many organisations still cannot describe the architecture through which their AI actually operates.
They know which products they have approved. They may know which models sit behind those products. They probably have policies describing what employees should and should not do. What they often cannot describe with the same confidence is what happens to information once somebody starts using them.
A policy can say that confidential information must not be submitted to an external model, but that only establishes the expected behaviour. The architecture determines whether the organisation can enforce that expectation, detect when it is breached and understand what happened afterwards.
It determines whether prompts leave the enterprise boundary, where they are processed, whether they are retained, which systems an AI service can reach, what information can be retrieved and where outputs are stored. It also determines whether any of those interactions can be reconstructed later if something goes wrong.
These are not primarily policy questions. They are architecture questions, and the distinction matters because organisations have spent decades discovering that writing something into a security policy does not magically cause the underlying technology to obey it. AI has apparently given us an opportunity to learn this lesson again, only faster and with considerably larger budgets.
Governance without architectural control therefore relies heavily on people consistently doing what the policy says. That is important, but it has never been a particularly robust control model on its own. If the only thing preventing sensitive information crossing a boundary is an employee remembering that paragraph twelve of the acceptable-use policy says they should not do it, the organisation does not really have a control. It has a request.
The boundary is becoming harder to see
Traditional enterprise systems generally gave us reasonably understandable boundaries. There was an application, some infrastructure, a collection of interfaces and identifiable stores of data. Information moved between systems through routes that architects and security teams could map, inspect and spend several months arguing about.
AI complicates that picture because the model is increasingly only one component in a much larger chain. A single business process might involve an employee application, an enterprise AI platform, a retrieval layer, several internal data sources, an external foundation model and tools capable of taking actions in other systems. Some of those components may be operated by the enterprise, some by technology providers and some by providers used by those providers.
Agents make the problem more interesting again. The system is no longer necessarily responding to a prompt and returning some text. It may decide which information to retrieve, which tools to invoke and which systems to interact with. As organisations give AI greater access to enterprise systems, the question of what the model is becomes progressively less useful than understanding what the overall system is capable of doing.
This changes the governance question. Asking which models the organisation is using is still useful, but it is nowhere near enough. The more important questions are what those models can see, what they can do, which boundaries they can cross and where the information involved actually goes.
Those questions require an architectural answer.
Sovereignty is about control, not geography
This is also why discussions about AI sovereignty can become unnecessarily simplistic. The argument is sometimes reduced to internally hosted models being safe and external models being dangerous, as though putting something inside your own data centre automatically confers wisdom upon everyone involved.
An internally hosted model with unrestricted access to sensitive information, weak identity controls and inadequate logging may represent a substantial risk. An externally hosted model operating behind strong contractual, technical and data controls may represent a much smaller one. Location matters, particularly where regulation and jurisdiction are involved, but it is only one part of the problem.
The more useful question is where control resides. Who controls the infrastructure and the data? Who determines identity and access? Where is information processed and retained? Which jurisdiction applies? Which other organisations participate in the service? Can the enterprise inspect what happened? Can it change provider without rebuilding everything surrounding the model?
Seen this way, sovereignty is not simply a hosting decision. It is the degree to which an organisation retains meaningful control over its technology, information and future choices. That needs to be designed into the architecture rather than appended to a governance checklist after procurement has already signed the contract.
There is another consequence of treating AI governance primarily as policy. If an organisation cannot reconstruct how an AI-enabled system reached a decision or performed an action, its governance framework eventually depends on trust rather than evidence.
That becomes increasingly important as AI moves from assisting people to participating directly in business processes. An organisation does not necessarily need to preserve every token involved in every interaction forever, but it should understand what evidence is required for the level of risk involved. For important processes, that may include which model was used, what information was available to it, which sources were retrieved, which tools it could access, what actions were attempted and which controls were applied.
Most of those records will never be examined. That is not the point. We do not build auditability because we expect every transaction to require an investigation; we build it because controls that cannot be demonstrated become very difficult to rely upon when something consequential happens.
Without that evidence, an AI governance committee risks becoming a group of senior people periodically confirming that everyone still agrees with the policy. Enterprises already have enough committees capable of doing that.
None of this makes policy, governance frameworks or responsible AI principles unnecessary. It makes them dependent on something more fundamental. Before creating another approval process, organisations should be able to map their AI estate with enough precision to understand which models and platforms are being used, which information they can reach, which trust boundaries they cross, which external organisations are involved, which actions they can perform and what controls exist at each boundary.
That map will probably be uncomfortable. It may reveal duplicate platforms, uncontrolled experimentation, overlapping vendor relationships and AI capabilities embedded inside products nobody originally thought of as AI systems. It may also reveal that some carefully written governance requirements cannot actually be enforced by the architecture currently underneath them.
Good governance should expose those things, because the objective is not to prevent organisations from using AI. Organisations that understand their boundaries should be able to move faster: they know where experimentation is safe, where additional controls are necessary and where the potential consequences justify greater scrutiny.
The alternative is governance by prohibition. When nobody can confidently explain what is happening, restricting everything starts to look rational. That protects the organisation from some risks, but it also protects it remarkably effectively from innovation.
AI governance therefore cannot belong solely to risk, security, legal or technology. Each has part of the answer, and each sees a different part of the risk. But underneath all of them sits the architecture that determines what the organisation can actually control.
Policy describes the organisation you would like to have. Architecture describes the organisation you actually built.
Interested in discussing technology leadership, AI governance or board advisory work?
Read the Executive Profile , explore the Board Profile , return to the Perspectives Library or get in touch .
Previous: Perspective 009 — The Handoffs Are Where Transformation Goes to Die
Clear judgement for consequential technology decisions.
Privacy: this site uses no cookies. Anonymous analytics help improve the site.
