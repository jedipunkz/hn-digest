---
source: "https://sdarchitect.blog/2026/09/06/ai-risk-a-users-guide-part-i-the-approval-nobody-signed/"
hn_url: "https://news.ycombinator.com/item?id=49694997"
title: "AI Risk: The Approval Nobody Signed Off"
article_title: "AI Risk – a User’s Guide (Part I): The Approval Nobody Signed off – Home: sdarchitect.blog"
image: "https://i0.wp.com/sdarchitect.blog/wp-content/uploads/2026/09/unapproved-autonomous-ai-decisions-bypassing-enterprise-governance.png?fit=1200%2C675&ssl=1"
author: "SanjeevSharma"
captured_at: "2026-09-14T11:58:51Z"
capture_tool: "hn-digest"
hn_id: 49694997
score: 5
comments: 0
posted_at: "2026-09-14T11:17:27Z"
tags:
  - hacker-news
---

# AI Risk: The Approval Nobody Signed Off

- HN: [49694997](https://news.ycombinator.com/item?id=49694997)
- Source: [sdarchitect.blog](https://sdarchitect.blog/2026/09/06/ai-risk-a-users-guide-part-i-the-approval-nobody-signed/)
- Score: 5
- Comments: 0
- Posted: 2026-09-14T11:17:27Z

## Translation

Title: AI Risk: The Approval Nobody Signed Off
Article title: AI Risk – a User’s Guide (Part I): The Approval Nobody Signed off – Home: sdarchitect.blog
Description: Here's a question I've been asking CIOs and CISOs at every conference session I speak at these days: “what if the riskiest AI decision your company made this year is one nobody approved?” Sit with that for a second. Not a decision your board debated. Not a model or Agentic deployment your risk commi
[truncated]

Article text:
AI Risk – a User’s Guide (Part I): The Approval Nobody Signed off – Home: sdarchitect.blog
Skip to content
Home: sdarchitect.blog
Sanjeev Sharma: My thoughts on Cloud, DevOps, Data Strategy, and Life…
AI Risk – a User’s Guide (Part I): The Approval Nobody Signed off
Here’s a question I’ve been asking CIOs and CISOs at every conference session I speak at these days: “what if the riskiest AI decision your company made this year is one nobody approved?”
Sit with that for a second. Not a decision your board debated. Not a model or Agentic deployment your risk committee signed off on. A decision, a real one, with real consequences, that an AI Agent leveraging a LLM made on its own, embedded three layers deep in some workflow nobody in the executive suite has looked at since it went live.
This is the uncomfortable truth I opened with at the AI Risk Summit in Half Moon Bay last month, and it’s the truth I want to unpack here, because I don’t think most leadership teams have actually internalized it yet.
Judgment calls happen every day, without a human-in-the-loop
When we built traditional software systems, every consequential branch in the logic was, in some sense, a decision someone made ahead of time. A product owner validated and approved the business rule. A compliance officer signed off on the workflow that decided who got a loan, what claim got flagged for fraud, and which action got escalated to a human for review. And an engineer wrote and tested the logic.
AI systems, especially the generative and agentic kind we’re all racing to adopt, don’t work that way. The “decision” does not synthesize from a pre-approved and validated workflow. It emerges, probabilistically, from a model doing inference at run time. Two customers with nearly identical profiles can get different answers from the same model on the same day, and there is no line of code you can point to and say “that is the logic behind this decision.”
That is not a bug. That’s the nature of LLM based AI. It is inherently non-deterministic. What it means is the old governance models, where risks got reviewed and approved before something shipped, now breaks down the moment the “something” is an agent using a AI model to make calls in production, continuously, at a scale and velocity no committee could ever keep pace with.
The liability doesn’t wait for your governance to catch up
I want to be precise about the stakes here, because “AI risk” can sound abstract until it isn’t. When a model hallucinates a fact in a customer-facing chatbot, misclassifies a loan applicant, or takes an autonomous action in an infrastructure pipeline that it should not have, the liability that follows does not pause and ask whether your governance program was mature enough to have caught it. Regulators, courts, and customers evaluate you on the outcome, not on your intention to eventually build the right guardrails or on the expectation of a right outcome.
This is exactly why the data I have been tracking, and that I shared at the summit, is so alarming. A very significant number of companies have already suffered reputational damage from AI errors. That is not a technology failure as that is how LLM based AI behaves. It is but a governance failure, and there is a real difference between those two framings that I think gets lost in a lot of the discourse.
When we say something is a “technology failure,” we are implicitly saying: the model was not good enough yet, and better models will fix it. When we say it is a “governance failure,” we are saying something much less comfortable: the organization did not build the right and rigorous oversight structure needed to catch and correct the mistake before it caused harm. And that gap will continue to exist regardless of how good the underlying model gets.
I lean firmly toward the second framing, and here’s why: no frontier model, no matter how advanced, will ever hit a zero error rate. Hallucination, drift, and edge-case failures are inherent properties of probabilistic/non-deterministic systems operating on a model’s ‘model’ of the world, no matter how comprehensive a ‘world model’ you build. If your risk posture depends on the model never being wrong, you have unfortunately already lost.
If your risk posture depends on the model never being wrong, you have unfortunately already lost.
Why AI outran governance so fast
I’ve spent a career in software delivery watching organizations struggle to keep governance in step with the pace of technical change: DevOps, cloud, containers, now agents. Each wave got faster in its pace of adoption. But nothing has moved as fast as generative AI, and there is a structural reason for that.
Traditional software governance assumed a deploy cadence measured in weeks or months, with discrete release gates. Generative AI systems are living, and even self-evolving systems. The same model can behave differently depending on prompt, context window, fine-tuning updates from a vendor you don’t control, drift that emerges silently over weeks of production traffic which the agent is now ‘learning’ from. There is no single “release” to put through gates with deterministic review and approval steps. There’s a continuous stream of judgment calls, and by the time a quarterly risk review catches up, thousands of those calls have already happened.
This is the core reason I keep telling audiences: GOVERN cannot be an atomic step in your process anymore. It has to be a continuous, cross-cutting function that runs in parallel with deployment, not a checkpoint that happens before it.
What “nobody approved it” actually means for you
Let me make this concrete with the kind of scenario I see in the field constantly. A customer support team adopts a generative AI tool to draft responses. Nobody in legal or risk formally reviewed the tool because it was rolled out as a “productivity aid,” not a decision making system. Six months later, that tool has generated thousands of customer-facing statements, some of them factually wrong, a few of them legally exposed, and no one owns the incident because no one ever officially “approved” the tool as an AI system requiring oversight.
This is not a hypothetical. This is the modal failure mode I am seeing across enterprises right now: risk accumulating in the shadows of well-intentioned productivity AI ‘bots’, invisible to governance until something breaks publicly. Shadow IT was nothing compared to the scale and impact of “Shadow AI”.
Shadow IT was nothing compared to the scale and impact of ‘Shadow AI’.
Where I think leaders should start
If you take one thing from this post, take this: the absence of a formal approval is not the absence of risk. It’s the presence of ungoverned risk, which is worse, because it is also unmeasured risk. You cannot manage what you have not even inventoried.
The fix is not complicated to state, even if it is hard to execute: every AI system touching your business, whether it was formally procured, quietly adopted by a team, or embedded inside a vendor’s product, needs to be on a list, with an owner, a risk tier, and a monitoring plan. That is the starting point I will build upon in the next few posts, where I will walk through what can turn this from an abstract mandate into something you can actually run on starting Monday morning.
The riskiest decision your company makes this year might already have happened. The question is whether you will know about its faults before your customers, your regulators, or the press does.
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
Founder and Principal - Data Capital Labs, Ex-Dell, Truist, IBM, Author: The DevOps Adoption Playbook, and DevOps For Dummies (IBM Edition). AI, Cloud, DevSecOps, Disruptive Tech.
View all posts by Sanjeev Sharma
This site uses Akismet to reduce spam. Learn how your comment data is processed.
Enter your email address to subscribe to this blog and receive notifications of new posts by email.
Subscribe
Subscribed
Home: sdarchitect.blog
Already have a WordPress.com account? Log in now.

## Original Extract

Here's a question I've been asking CIOs and CISOs at every conference session I speak at these days: “what if the riskiest AI decision your company made this year is one nobody approved?” Sit with that for a second. Not a decision your board debated. Not a model or Agentic deployment your risk commi
[truncated]

AI Risk – a User’s Guide (Part I): The Approval Nobody Signed off – Home: sdarchitect.blog
Skip to content
Home: sdarchitect.blog
Sanjeev Sharma: My thoughts on Cloud, DevOps, Data Strategy, and Life…
AI Risk – a User’s Guide (Part I): The Approval Nobody Signed off
Here’s a question I’ve been asking CIOs and CISOs at every conference session I speak at these days: “what if the riskiest AI decision your company made this year is one nobody approved?”
Sit with that for a second. Not a decision your board debated. Not a model or Agentic deployment your risk committee signed off on. A decision, a real one, with real consequences, that an AI Agent leveraging a LLM made on its own, embedded three layers deep in some workflow nobody in the executive suite has looked at since it went live.
This is the uncomfortable truth I opened with at the AI Risk Summit in Half Moon Bay last month, and it’s the truth I want to unpack here, because I don’t think most leadership teams have actually internalized it yet.
Judgment calls happen every day, without a human-in-the-loop
When we built traditional software systems, every consequential branch in the logic was, in some sense, a decision someone made ahead of time. A product owner validated and approved the business rule. A compliance officer signed off on the workflow that decided who got a loan, what claim got flagged for fraud, and which action got escalated to a human for review. And an engineer wrote and tested the logic.
AI systems, especially the generative and agentic kind we’re all racing to adopt, don’t work that way. The “decision” does not synthesize from a pre-approved and validated workflow. It emerges, probabilistically, from a model doing inference at run time. Two customers with nearly identical profiles can get different answers from the same model on the same day, and there is no line of code you can point to and say “that is the logic behind this decision.”
That is not a bug. That’s the nature of LLM based AI. It is inherently non-deterministic. What it means is the old governance models, where risks got reviewed and approved before something shipped, now breaks down the moment the “something” is an agent using a AI model to make calls in production, continuously, at a scale and velocity no committee could ever keep pace with.
The liability doesn’t wait for your governance to catch up
I want to be precise about the stakes here, because “AI risk” can sound abstract until it isn’t. When a model hallucinates a fact in a customer-facing chatbot, misclassifies a loan applicant, or takes an autonomous action in an infrastructure pipeline that it should not have, the liability that follows does not pause and ask whether your governance program was mature enough to have caught it. Regulators, courts, and customers evaluate you on the outcome, not on your intention to eventually build the right guardrails or on the expectation of a right outcome.
This is exactly why the data I have been tracking, and that I shared at the summit, is so alarming. A very significant number of companies have already suffered reputational damage from AI errors. That is not a technology failure as that is how LLM based AI behaves. It is but a governance failure, and there is a real difference between those two framings that I think gets lost in a lot of the discourse.
When we say something is a “technology failure,” we are implicitly saying: the model was not good enough yet, and better models will fix it. When we say it is a “governance failure,” we are saying something much less comfortable: the organization did not build the right and rigorous oversight structure needed to catch and correct the mistake before it caused harm. And that gap will continue to exist regardless of how good the underlying model gets.
I lean firmly toward the second framing, and here’s why: no frontier model, no matter how advanced, will ever hit a zero error rate. Hallucination, drift, and edge-case failures are inherent properties of probabilistic/non-deterministic systems operating on a model’s ‘model’ of the world, no matter how comprehensive a ‘world model’ you build. If your risk posture depends on the model never being wrong, you have unfortunately already lost.
If your risk posture depends on the model never being wrong, you have unfortunately already lost.
Why AI outran governance so fast
I’ve spent a career in software delivery watching organizations struggle to keep governance in step with the pace of technical change: DevOps, cloud, containers, now agents. Each wave got faster in its pace of adoption. But nothing has moved as fast as generative AI, and there is a structural reason for that.
Traditional software governance assumed a deploy cadence measured in weeks or months, with discrete release gates. Generative AI systems are living, and even self-evolving systems. The same model can behave differently depending on prompt, context window, fine-tuning updates from a vendor you don’t control, drift that emerges silently over weeks of production traffic which the agent is now ‘learning’ from. There is no single “release” to put through gates with deterministic review and approval steps. There’s a continuous stream of judgment calls, and by the time a quarterly risk review catches up, thousands of those calls have already happened.
This is the core reason I keep telling audiences: GOVERN cannot be an atomic step in your process anymore. It has to be a continuous, cross-cutting function that runs in parallel with deployment, not a checkpoint that happens before it.
What “nobody approved it” actually means for you
Let me make this concrete with the kind of scenario I see in the field constantly. A customer support team adopts a generative AI tool to draft responses. Nobody in legal or risk formally reviewed the tool because it was rolled out as a “productivity aid,” not a decision making system. Six months later, that tool has generated thousands of customer-facing statements, some of them factually wrong, a few of them legally exposed, and no one owns the incident because no one ever officially “approved” the tool as an AI system requiring oversight.
This is not a hypothetical. This is the modal failure mode I am seeing across enterprises right now: risk accumulating in the shadows of well-intentioned productivity AI ‘bots’, invisible to governance until something breaks publicly. Shadow IT was nothing compared to the scale and impact of “Shadow AI”.
Shadow IT was nothing compared to the scale and impact of ‘Shadow AI’.
Where I think leaders should start
If you take one thing from this post, take this: the absence of a formal approval is not the absence of risk. It’s the presence of ungoverned risk, which is worse, because it is also unmeasured risk. You cannot manage what you have not even inventoried.
The fix is not complicated to state, even if it is hard to execute: every AI system touching your business, whether it was formally procured, quietly adopted by a team, or embedded inside a vendor’s product, needs to be on a list, with an owner, a risk tier, and a monitoring plan. That is the starting point I will build upon in the next few posts, where I will walk through what can turn this from an abstract mandate into something you can actually run on starting Monday morning.
The riskiest decision your company makes this year might already have happened. The question is whether you will know about its faults before your customers, your regulators, or the press does.
Share on X (Opens in new window)
X
Share on Facebook (Opens in new window)
Facebook
Founder and Principal - Data Capital Labs, Ex-Dell, Truist, IBM, Author: The DevOps Adoption Playbook, and DevOps For Dummies (IBM Edition). AI, Cloud, DevSecOps, Disruptive Tech.
View all posts by Sanjeev Sharma
This site uses Akismet to reduce spam. Learn how your comment data is processed.
Enter your email address to subscribe to this blog and receive notifications of new posts by email.
Subscribe
Subscribed
Home: sdarchitect.blog
Already have a WordPress.com account? Log in now.
