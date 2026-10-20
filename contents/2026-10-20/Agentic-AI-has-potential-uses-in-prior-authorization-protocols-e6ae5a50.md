---
source: "https://civicpivot.blogspot.com/2026/09/healthcares-ai-arms-race-can-agentic-ai.html"
hn_url: "https://news.ycombinator.com/item?id=50170755"
title: "Agentic AI has potential uses in prior-authorization protocols"
article_title: "Healthcare’s AI Arms Race: Can Agentic AI End It?"
image: "https://blogger.googleusercontent.com/img/a/AVvXsEhWbS--yhBxHKTxC-a3qMTViwoWp71MSPOih-mdGNebyaEwleD4bF1VgAH3u0YFy0N-ivMvx02ylPY64n9CdKOmYo935d49sAn1lBcOConPEBci7ZINQ53zSTE1mOA087zu7gE6MrQ_uS37bANP8YwZHj8lqEyVjqWVKQE69TETylHcQqeh0ZDgwTkIMRA=w1200-h630-p-k-no-nu"
author: "civicpivot"
captured_at: "2026-10-20T18:57:38Z"
capture_tool: "hn-digest"
hn_id: 50170755
score: 1
comments: 0
posted_at: "2026-10-20T18:40:17Z"
tags:
  - hacker-news
---

# Agentic AI has potential uses in prior-authorization protocols

- HN: [50170755](https://news.ycombinator.com/item?id=50170755)
- Source: [civicpivot.blogspot.com](https://civicpivot.blogspot.com/2026/09/healthcares-ai-arms-race-can-agentic-ai.html)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T18:40:17Z

## Translation

Title: Agentic AI has potential uses in prior-authorization protocols
Article title: Healthcare’s AI Arms Race: Can Agentic AI End It?
Description: Why agentic AI’s most important healthcare test may be fixing the relationship between payers and providers In 2026, a physician...

Article text:
Skip to main content
Reshaping modern policy, data science, and medical infrastructure through a classical philosophical lens.
Healthcare’s AI Arms Race: Can Agentic AI End It?
Why agentic AI’s most important healthcare test may be fixing the relationship between payers and providers
In 2026, a physician can use AI to write a prior-authorization request. A health plan can use AI to review it. If the request is denied, another AI can draft the appeal.
When I first thought through that sequence, it seemed almost absurd. We've built intelligent tools for every stage of the disagreement, but not for preventing the disagreement itself! Are we using AI in healthcare to solve the wrong problem?
This is the strange place that healthcare has reached. Providers are getting better at assembling claims, contesting denials, and recovering revenue, and payers are getting better at applying medical policies, identifying inconsistencies, and controlling utilization. Yet neither side necessarily knows what the other needs until something goes wrong.
We've automated the argument without resolving the disagreement.
This problem is not new, but the conditions surrounding it are.
New CMS requirements are shortening prior-authorization decision timelines and requiring affected payers to provide more specific reasons for denials. By January 2027, these payers must implement new application programming interfaces, or APIs, intended to improve the exchange of information among patients, providers, and health plans. Major insurers have also pledged to reduce authorization requirements, standardize electronic submissions, and move more approvals toward real-time decisions. (
CMS
,
HHS
)
At the same time, AI is moving beyond generating text. A new class of “agentic” systems can interpret changing conditions, consult multiple sources, and coordinate a sequence of actions toward a goal.
For the first time, regulatory pressure, interoperable data, and agentic AI are arriving together.
That convergence creates a choice. We can use AI to help providers submit more requests while helping payers scrutinize them more efficiently. The administrative arms race will accelerate, but the underlying relationship will remain unchanged.
The Problem Is Not Just Paperwork
Payer-provider friction is often described as an administrative burden. That is true, but incomplete. The deeper problem is misalignment.
Clinical evidence and payer requirements live in different systems. Policy criteria may become visible to the provider only after a request is submitted. Each organization optimizes its own part of the transaction, even when doing so creates more work across the system.
A denial, then, is not merely a failed claim. It is often a delayed signal that two organizations were working from different information, rules, or assumptions.
Traditional Revenue Cycle Management is built to respond to that signal: correct the claim, assemble the appeal, schedule the peer-to-peer review, and pursue payment. It treats friction as something to manage.
A more intelligent revenue cycle would treat friction as information and move that information upstream.
Why Faster Is Not the Same as Smarter
At first, faster automation seems like the obvious answer. If an authorization takes fourteen days, build a system that processes it in one.
But speed alone does not fix the underlying problem. If a provider submits an incomplete request in seconds and a payer denies it in seconds, congratulations, we have achieved real-time dysfunction!
Traditional automation asks:
How can I complete this task faster?
Agentic orchestration asks:
Why did this task become necessary, what information is missing, and which action should happen next?
That is a much larger shift. The goal is no longer to automate one isolated step, checking eligibility, extracting a diagnosis, or generating an appeal. It is to coordinate the entire process so that the right information reaches the right system before a preventable failure occurs.
What the Architecture Actually Looks Like
The breakthrough is not a swarm of chatbots sending messages to one another. It is an orchestration layer built on top of healthcare’s emerging interoperability infrastructure.
Imagine that a physician orders a knee MRI in an electronic health record. The order can be represented using
FHIR
, a standard for exchanging healthcare data through APIs. Three implementation guides developed through the HL7 Da Vinci Project can then support different parts of the authorization process:
Coverage Requirements Discovery (CRD)
asks the payer whether authorization is required and which coverage rules apply.
Documentation Templates and Rules (DTR)
identifies the required evidence and helps retrieve it from the patient’s record.
Prior Authorization Support (PAS)
packages the request, sends it to the payer, and returns the response.
These standards create the rails. Agentic AI coordinates what travels across them.
An orchestration agent could maintain the state of the request and determine which service to call next. A retrieval agent could search both structured data and unstructured clinical notes for the patient’s physical therapy history. A policy engine could compare that evidence with the payer’s current criteria. A submission agent could then assemble the required FHIR resources and send a completed authorization request.
The distinction here is crucial: the language model should help interpret messy clinical information, but it should not become the source of truth.
If an agent claims that a patient completed six weeks of physical therapy, it should link that claim to the exact encounters or notes supporting it. If it applies a coverage rule, it should record the policy version, benefit plan, jurisdiction, and date. Eligibility and policy requirements should come from authoritative systems, while deterministic rules should handle decisions that do not require interpretation.
The agent does not replace the API, policy engine, or clinical reviewer. It coordinates them.
This architecture is becoming possible at a particularly important moment. Under CMS-0057-F, affected payers must implement Prior Authorization APIs that can identify documentation requirements, accept electronic requests, and return one of three structured responses: approval, a specific reason for denial, or a request for more information. The API requirements begin in 2027, creating a real technical foundation on which agentic workflows could operate, not just a hypothetical one.
The Difficult Part Begins After the Demo
On a diagram, this process looks almost suspiciously clean. Real healthcare data is not.
A physical-therapy session might appear as a coded encounter, a scanned document, one sentence buried in a progress note, or nowhere at all. The same procedure may be represented by different codes in different systems. A payer might update its policy while a request is still being reviewed. A patient’s insurance could change midway through treatment.
A safe system therefore has to solve several hard problems at once.
Identity resolution:
It must correctly match the patient, provider, health plan, and policy. A brilliantly reasoned answer about the wrong patient is still a dangerous answer.
Semantic interoperability:
Moving information is not the same as understanding it. The system must reconcile different codes, local terminology, abbreviations, and free text.
Policy versioning:
Every recommendation must be tied to the exact rule used at that moment. Otherwise, neither side can reconstruct why the agent acted as it did.
Grounded retrieval:
Every extracted clinical claim should point back to its source. “The model probably saw it somewhere in the chart” is not an acceptable audit trail.
Workflow state:
Prior authorization is not one prompt followed by one answer. A request may pause for documentation, resume days later, retry after a technical failure, or escalate to human review. The orchestrator must remember what happened without losing information or accidentally submitting the same request twice.
Security and permissions:
Each agent should receive only the information and authority required for its task. An agent allowed to retrieve documentation should not automatically be allowed to approve care or modify the medical record.
Human escalation:
The system must recognize genuine ambiguity. When the evidence is missing, contradictory, or dependent on clinical judgment, automation should stop and explain why a person needs to intervene.
This is why a compelling demonstration is not the same as a deployable healthcare system. The hardest work is not getting an AI model to generate an answer. It is building provenance, permissions, policy logic, failure recovery, and accountability around that answer.
The strongest architecture would therefore combine several kinds of intelligence rather than asking a language model to do everything. FHIR would provide the exchange standard. Policy engines would apply repeatable rules. Identity and access systems would define trust boundaries. Language models would interpret unstructured information. Agentic AI would coordinate the workflow across them.
That may sound less dramatic than “autonomous AI transforms healthcare.” But it's actually more radical. Healthcare has spent decades building systems that exchange transactions. Agentic orchestration could allow those systems to coordinate around a shared goal.
What would that difference look like for an actual patient?
What One Day Instead of Fourteen Could Mean
Consider a patient who needs a knee MRI.
Today, a provider may submit the request without knowing that the payer requires documentation of prior physical therapy. Days later, the payer asks for additional records. Staff search the chart and resubmit the request. A second review finds the documentation insufficient, and the case moves to a peer-to-peer discussion. The physician spends 35 minutes defending a clinically reasonable decision. Nearly two weeks after the original order, the MRI is approved.
Everyone may have followed the process. The process still failed the patient.
In an orchestrated model, the policy agent identifies the physical-therapy requirement when the MRI is ordered. The clinical agent locates the relevant history, and the documentation agent flags any remaining gap. A complete, policy-aligned request is submitted once and approved the same day.
Not every case should receive automatic approval. Medicine contains genuine ambiguity, and difficult decisions still require human judgment. The value lies in separating those cases from the ones delayed only because one system could not see what another system required.
Speed is the visible benefit. Shared understanding is the deeper one.
The Business Case Is a Friction Case
Administrative misalignment is expensive because it creates work on both sides of the transaction.
The enterprise-wide value of operational optimization is clear. On the provider side, mitigating denials and collection costs transforms administrative friction into bottom-line savings. For large payers, reducing processing overhead, preventing unnecessary clinical spend, and strengthening Medicare Advantage operations drives significant, long-term financial health. Actual results will depend on scale, baseline denial rates, contracts, data quality, adoption, and the services included. What matters is the mechanism behind them:
Earlier alignment leads to more complete submissions. More complete submissions lead to fewer denials and appeals. Fewer administrative failures mean lower costs and faster care.
The strongest business case is therefore not that payers win or that providers win. It is that both are currently paying for the same friction in different places.
The Hard Part Is Not the Model
It would be tempting to treat this as a technology problem with a technology solution. It is not.
A shared intelligence layer raises difficult questions. Who determines

[truncated]

## Original Extract

Why agentic AI’s most important healthcare test may be fixing the relationship between payers and providers In 2026, a physician...

Skip to main content
Reshaping modern policy, data science, and medical infrastructure through a classical philosophical lens.
Healthcare’s AI Arms Race: Can Agentic AI End It?
Why agentic AI’s most important healthcare test may be fixing the relationship between payers and providers
In 2026, a physician can use AI to write a prior-authorization request. A health plan can use AI to review it. If the request is denied, another AI can draft the appeal.
When I first thought through that sequence, it seemed almost absurd. We've built intelligent tools for every stage of the disagreement, but not for preventing the disagreement itself! Are we using AI in healthcare to solve the wrong problem?
This is the strange place that healthcare has reached. Providers are getting better at assembling claims, contesting denials, and recovering revenue, and payers are getting better at applying medical policies, identifying inconsistencies, and controlling utilization. Yet neither side necessarily knows what the other needs until something goes wrong.
We've automated the argument without resolving the disagreement.
This problem is not new, but the conditions surrounding it are.
New CMS requirements are shortening prior-authorization decision timelines and requiring affected payers to provide more specific reasons for denials. By January 2027, these payers must implement new application programming interfaces, or APIs, intended to improve the exchange of information among patients, providers, and health plans. Major insurers have also pledged to reduce authorization requirements, standardize electronic submissions, and move more approvals toward real-time decisions. (
CMS
,
HHS
)
At the same time, AI is moving beyond generating text. A new class of “agentic” systems can interpret changing conditions, consult multiple sources, and coordinate a sequence of actions toward a goal.
For the first time, regulatory pressure, interoperable data, and agentic AI are arriving together.
That convergence creates a choice. We can use AI to help providers submit more requests while helping payers scrutinize them more efficiently. The administrative arms race will accelerate, but the underlying relationship will remain unchanged.
The Problem Is Not Just Paperwork
Payer-provider friction is often described as an administrative burden. That is true, but incomplete. The deeper problem is misalignment.
Clinical evidence and payer requirements live in different systems. Policy criteria may become visible to the provider only after a request is submitted. Each organization optimizes its own part of the transaction, even when doing so creates more work across the system.
A denial, then, is not merely a failed claim. It is often a delayed signal that two organizations were working from different information, rules, or assumptions.
Traditional Revenue Cycle Management is built to respond to that signal: correct the claim, assemble the appeal, schedule the peer-to-peer review, and pursue payment. It treats friction as something to manage.
A more intelligent revenue cycle would treat friction as information and move that information upstream.
Why Faster Is Not the Same as Smarter
At first, faster automation seems like the obvious answer. If an authorization takes fourteen days, build a system that processes it in one.
But speed alone does not fix the underlying problem. If a provider submits an incomplete request in seconds and a payer denies it in seconds, congratulations, we have achieved real-time dysfunction!
Traditional automation asks:
How can I complete this task faster?
Agentic orchestration asks:
Why did this task become necessary, what information is missing, and which action should happen next?
That is a much larger shift. The goal is no longer to automate one isolated step, checking eligibility, extracting a diagnosis, or generating an appeal. It is to coordinate the entire process so that the right information reaches the right system before a preventable failure occurs.
What the Architecture Actually Looks Like
The breakthrough is not a swarm of chatbots sending messages to one another. It is an orchestration layer built on top of healthcare’s emerging interoperability infrastructure.
Imagine that a physician orders a knee MRI in an electronic health record. The order can be represented using
FHIR
, a standard for exchanging healthcare data through APIs. Three implementation guides developed through the HL7 Da Vinci Project can then support different parts of the authorization process:
Coverage Requirements Discovery (CRD)
asks the payer whether authorization is required and which coverage rules apply.
Documentation Templates and Rules (DTR)
identifies the required evidence and helps retrieve it from the patient’s record.
Prior Authorization Support (PAS)
packages the request, sends it to the payer, and returns the response.
These standards create the rails. Agentic AI coordinates what travels across them.
An orchestration agent could maintain the state of the request and determine which service to call next. A retrieval agent could search both structured data and unstructured clinical notes for the patient’s physical therapy history. A policy engine could compare that evidence with the payer’s current criteria. A submission agent could then assemble the required FHIR resources and send a completed authorization request.
The distinction here is crucial: the language model should help interpret messy clinical information, but it should not become the source of truth.
If an agent claims that a patient completed six weeks of physical therapy, it should link that claim to the exact encounters or notes supporting it. If it applies a coverage rule, it should record the policy version, benefit plan, jurisdiction, and date. Eligibility and policy requirements should come from authoritative systems, while deterministic rules should handle decisions that do not require interpretation.
The agent does not replace the API, policy engine, or clinical reviewer. It coordinates them.
This architecture is becoming possible at a particularly important moment. Under CMS-0057-F, affected payers must implement Prior Authorization APIs that can identify documentation requirements, accept electronic requests, and return one of three structured responses: approval, a specific reason for denial, or a request for more information. The API requirements begin in 2027, creating a real technical foundation on which agentic workflows could operate, not just a hypothetical one.
The Difficult Part Begins After the Demo
On a diagram, this process looks almost suspiciously clean. Real healthcare data is not.
A physical-therapy session might appear as a coded encounter, a scanned document, one sentence buried in a progress note, or nowhere at all. The same procedure may be represented by different codes in different systems. A payer might update its policy while a request is still being reviewed. A patient’s insurance could change midway through treatment.
A safe system therefore has to solve several hard problems at once.
Identity resolution:
It must correctly match the patient, provider, health plan, and policy. A brilliantly reasoned answer about the wrong patient is still a dangerous answer.
Semantic interoperability:
Moving information is not the same as understanding it. The system must reconcile different codes, local terminology, abbreviations, and free text.
Policy versioning:
Every recommendation must be tied to the exact rule used at that moment. Otherwise, neither side can reconstruct why the agent acted as it did.
Grounded retrieval:
Every extracted clinical claim should point back to its source. “The model probably saw it somewhere in the chart” is not an acceptable audit trail.
Workflow state:
Prior authorization is not one prompt followed by one answer. A request may pause for documentation, resume days later, retry after a technical failure, or escalate to human review. The orchestrator must remember what happened without losing information or accidentally submitting the same request twice.
Security and permissions:
Each agent should receive only the information and authority required for its task. An agent allowed to retrieve documentation should not automatically be allowed to approve care or modify the medical record.
Human escalation:
The system must recognize genuine ambiguity. When the evidence is missing, contradictory, or dependent on clinical judgment, automation should stop and explain why a person needs to intervene.
This is why a compelling demonstration is not the same as a deployable healthcare system. The hardest work is not getting an AI model to generate an answer. It is building provenance, permissions, policy logic, failure recovery, and accountability around that answer.
The strongest architecture would therefore combine several kinds of intelligence rather than asking a language model to do everything. FHIR would provide the exchange standard. Policy engines would apply repeatable rules. Identity and access systems would define trust boundaries. Language models would interpret unstructured information. Agentic AI would coordinate the workflow across them.
That may sound less dramatic than “autonomous AI transforms healthcare.” But it's actually more radical. Healthcare has spent decades building systems that exchange transactions. Agentic orchestration could allow those systems to coordinate around a shared goal.
What would that difference look like for an actual patient?
What One Day Instead of Fourteen Could Mean
Consider a patient who needs a knee MRI.
Today, a provider may submit the request without knowing that the payer requires documentation of prior physical therapy. Days later, the payer asks for additional records. Staff search the chart and resubmit the request. A second review finds the documentation insufficient, and the case moves to a peer-to-peer discussion. The physician spends 35 minutes defending a clinically reasonable decision. Nearly two weeks after the original order, the MRI is approved.
Everyone may have followed the process. The process still failed the patient.
In an orchestrated model, the policy agent identifies the physical-therapy requirement when the MRI is ordered. The clinical agent locates the relevant history, and the documentation agent flags any remaining gap. A complete, policy-aligned request is submitted once and approved the same day.
Not every case should receive automatic approval. Medicine contains genuine ambiguity, and difficult decisions still require human judgment. The value lies in separating those cases from the ones delayed only because one system could not see what another system required.
Speed is the visible benefit. Shared understanding is the deeper one.
The Business Case Is a Friction Case
Administrative misalignment is expensive because it creates work on both sides of the transaction.
The enterprise-wide value of operational optimization is clear. On the provider side, mitigating denials and collection costs transforms administrative friction into bottom-line savings. For large payers, reducing processing overhead, preventing unnecessary clinical spend, and strengthening Medicare Advantage operations drives significant, long-term financial health. Actual results will depend on scale, baseline denial rates, contracts, data quality, adoption, and the services included. What matters is the mechanism behind them:
Earlier alignment leads to more complete submissions. More complete submissions lead to fewer denials and appeals. Fewer administrative failures mean lower costs and faster care.
The strongest business case is therefore not that payers win or that providers win. It is that both are currently paying for the same friction in different places.
The Hard Part Is Not the Model
It would be tempting to treat this as a technology problem with a technology solution. It is not.
A shared intelligence layer raises difficult questions. Who determines

[truncated]
