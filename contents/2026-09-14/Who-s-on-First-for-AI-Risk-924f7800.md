---
source: "https://sdarchitect.blog/2026/09/13/ai-risk-a-users-guide-part-ii-whos-on-first-for-ai-risk/"
hn_url: "https://news.ycombinator.com/item?id=49694992"
title: "'Who's on First' for AI Risk?"
article_title: "AI Risk – a User’s Guide (Part II): ‘Who’s on First’ for AI Risk? – Home: sdarchitect.blog"
image: "https://i0.wp.com/sdarchitect.blog/wp-content/uploads/2026/09/ai-risk-ii-2.png?fit=1200%2C675&ssl=1"
author: "SanjeevSharma"
captured_at: "2026-09-14T11:58:53Z"
capture_tool: "hn-digest"
hn_id: 49694992
score: 2
comments: 0
posted_at: "2026-09-14T11:16:45Z"
tags:
  - hacker-news
---

# 'Who's on First' for AI Risk?

- HN: [49694992](https://news.ycombinator.com/item?id=49694992)
- Source: [sdarchitect.blog](https://sdarchitect.blog/2026/09/13/ai-risk-a-users-guide-part-ii-whos-on-first-for-ai-risk/)
- Score: 2
- Comments: 0
- Posted: 2026-09-14T11:16:45Z

## Translation

Title: 'Who's on First' for AI Risk?
Article title: AI Risk – a User’s Guide (Part II): ‘Who’s on First’ for AI Risk? – Home: sdarchitect.blog
Description: Buying the AI Doesn't Transfer the Risk There is a comforting story a lot of enterprises tell themselves when they adopt a third-party AI tool: "we bought this from a reputable vendor who has indemnified us in their contract, so the vendor owns the risk." I understand the appeal of that story. It is
[truncated]

Article text:
AI Risk – a User’s Guide (Part II): ‘Who’s on First’ for AI Risk? – Home: sdarchitect.blog
Skip to content
Home: sdarchitect.blog
Sanjeev Sharma: My thoughts on Cloud, DevOps, Data Strategy, and Life…
AI Risk – a User’s Guide (Part II): ‘Who’s on First’ for AI Risk?
Buying the AI Doesn’t Transfer the Risk
There is a comforting story a lot of enterprises tell themselves when they adopt a third-party AI tool: “we bought this from a reputable vendor who has indemnified us in their contract, so the vendor owns the risk.” I understand the appeal of that story. It is the same logic that’s worked for decades with traditional software procurement where you buy the license, the vendor is liable for defects, you move on. I want to walk through why that story doesn’t hold for AI systems, and why the legal reality is far messier and far more your problem than most procurement teams currently assume.
Liability is a chain, not a handoff
Under the EU AI Act, which just went into force in August 2026, and increasingly under other emerging frameworks, liability for an AI system flows through a chain of roles, not a single handoff point. The provider: the vendor who builds the model or system, carries specific obligations under Articles 9 through 17 and Article 43. These include, risk management, technical documentation, conformity assessment, and accuracy and security testing at design time. That is their job, and it is a real and substantial one.
But the deployer (that is you), the enterprise putting that system into production for your actual customers or employees, carries a completely separate, and equally substantial, set of obligations under Articles 26 and 27. These include, human oversight, using the system within its intended purpose, monitoring and logging, incident reporting, and impact assessments.
Notice what is missing from that second list: nothing about it depends on how good the vendor’s original risk management was. You own human oversight and monitoring regardless of how carefully the provider built the underlying model or harness. Your compliance obligation exists independently of theirs. This is the part that catches procurement and legal teams off guard as they treat the vendor relationship as risk transfer, when the regulation actually treats it as risk splitting.
The trigger that quietly changes your role
Here is the part of this I think deserves the most attention, because it is the one most likely to blindside an organization that thinks it has this figured out: certain actions can silently convert you from a deployer into a full provider under Article 25 of the act. Something as simple as rebranding a vendor’s AI system under your own product name. Fine-tuning it on your own data – which more and more organizations are doing to make the AI better suited for their use cases. Or repurposing it for a use case meaningfully different from what it was designed and certified for, as per the vendor language.
Any of those three moves, and I would bet a meaningful percentage of enterprises reading this have done at least one if not all of them somewhere in their AI stack without realizing the regulatory implications, shifts you into the provider category. That means you also inherit the full weight of Articles 9 through 17: the design-time risk management, the technical documentation, the conformity assessment obligations, etc. that you assumed were the vendor’s problem, not yours.
I have watched engineering teams fine-tune a vendor’s foundation model on proprietary customer data as a completely routine technical decision, made without any awareness that it might have just converted their legal posture. That is the gap I want every technical leader reading this to close: your engineering choices about how you use a third-party AI system are, in a very real sense, legal and risk decisions now. Treat them that way.
Mobley v. Workday: the case that made this concrete
I regularly bring up Mobley v. Workday as an example lawsuit to follow, because it is not a hypothetical scenario. It is a live case, which while still pending a decision, shows liability now spanning deployer, vendor, and infrastructure provider simultaneously. That is the pattern to watch: courts are increasingly unwilling to let any single party in the AI supply chain fully externalize responsibility onto another. The deployer who used the tool, the vendor who built it, and the infrastructure provider who hosted it can all end up named, and all end up exposed.
This should reshape how you think about vendor contracts. A vendor’s terms of service disclaiming liability for model outputs does not erase your regulatory exposure as a deployer , and it increasingly does not fully insulate the vendor either. Everyone in the chain has skin in the game now, which means everyone in the chain needs their own governance evidence, independently defensible, regardless of what the contract says about indemnification.
What this means for procurement, practically
If you take one operational change away from this post, make it this: AI vendor risk cannot live solely inside legal’s contract review process anymore. It needs a technical governance review layered on top, one that specifically asks:
What is this vendor’s provider-side documentation, and have we actually reviewed it, not just accepted it exists?
Are we fine-tuning, rebranding, or repurposing this system in any way that could trigger Article 25?
Do we have our own deployer-side obligations in place: human oversight, monitoring, logging, incident reporting, actually implemented, or just assumed to be the vendor’s problem?
If this system fails, who inside our organization owns the incident response, and have they rehearsed it?
None of these questions are hard to ask. What is hard is building the organizational muscle to ask them before deployment rather than during a post-incident scramble. That is where I am headed next in this series. What you as an organization need is the operational cycle that moves you on from figuring out “who is liable” from a legal perspective, to a continuous governance discipline you run to properly quantify and qualify your risk, and most importantly, the operational discipline to respond when the risk exceeds your enterprise’s risk appetite or tolerance. More on these topics in coming posts.
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

Buying the AI Doesn't Transfer the Risk There is a comforting story a lot of enterprises tell themselves when they adopt a third-party AI tool: "we bought this from a reputable vendor who has indemnified us in their contract, so the vendor owns the risk." I understand the appeal of that story. It is
[truncated]

AI Risk – a User’s Guide (Part II): ‘Who’s on First’ for AI Risk? – Home: sdarchitect.blog
Skip to content
Home: sdarchitect.blog
Sanjeev Sharma: My thoughts on Cloud, DevOps, Data Strategy, and Life…
AI Risk – a User’s Guide (Part II): ‘Who’s on First’ for AI Risk?
Buying the AI Doesn’t Transfer the Risk
There is a comforting story a lot of enterprises tell themselves when they adopt a third-party AI tool: “we bought this from a reputable vendor who has indemnified us in their contract, so the vendor owns the risk.” I understand the appeal of that story. It is the same logic that’s worked for decades with traditional software procurement where you buy the license, the vendor is liable for defects, you move on. I want to walk through why that story doesn’t hold for AI systems, and why the legal reality is far messier and far more your problem than most procurement teams currently assume.
Liability is a chain, not a handoff
Under the EU AI Act, which just went into force in August 2026, and increasingly under other emerging frameworks, liability for an AI system flows through a chain of roles, not a single handoff point. The provider: the vendor who builds the model or system, carries specific obligations under Articles 9 through 17 and Article 43. These include, risk management, technical documentation, conformity assessment, and accuracy and security testing at design time. That is their job, and it is a real and substantial one.
But the deployer (that is you), the enterprise putting that system into production for your actual customers or employees, carries a completely separate, and equally substantial, set of obligations under Articles 26 and 27. These include, human oversight, using the system within its intended purpose, monitoring and logging, incident reporting, and impact assessments.
Notice what is missing from that second list: nothing about it depends on how good the vendor’s original risk management was. You own human oversight and monitoring regardless of how carefully the provider built the underlying model or harness. Your compliance obligation exists independently of theirs. This is the part that catches procurement and legal teams off guard as they treat the vendor relationship as risk transfer, when the regulation actually treats it as risk splitting.
The trigger that quietly changes your role
Here is the part of this I think deserves the most attention, because it is the one most likely to blindside an organization that thinks it has this figured out: certain actions can silently convert you from a deployer into a full provider under Article 25 of the act. Something as simple as rebranding a vendor’s AI system under your own product name. Fine-tuning it on your own data – which more and more organizations are doing to make the AI better suited for their use cases. Or repurposing it for a use case meaningfully different from what it was designed and certified for, as per the vendor language.
Any of those three moves, and I would bet a meaningful percentage of enterprises reading this have done at least one if not all of them somewhere in their AI stack without realizing the regulatory implications, shifts you into the provider category. That means you also inherit the full weight of Articles 9 through 17: the design-time risk management, the technical documentation, the conformity assessment obligations, etc. that you assumed were the vendor’s problem, not yours.
I have watched engineering teams fine-tune a vendor’s foundation model on proprietary customer data as a completely routine technical decision, made without any awareness that it might have just converted their legal posture. That is the gap I want every technical leader reading this to close: your engineering choices about how you use a third-party AI system are, in a very real sense, legal and risk decisions now. Treat them that way.
Mobley v. Workday: the case that made this concrete
I regularly bring up Mobley v. Workday as an example lawsuit to follow, because it is not a hypothetical scenario. It is a live case, which while still pending a decision, shows liability now spanning deployer, vendor, and infrastructure provider simultaneously. That is the pattern to watch: courts are increasingly unwilling to let any single party in the AI supply chain fully externalize responsibility onto another. The deployer who used the tool, the vendor who built it, and the infrastructure provider who hosted it can all end up named, and all end up exposed.
This should reshape how you think about vendor contracts. A vendor’s terms of service disclaiming liability for model outputs does not erase your regulatory exposure as a deployer , and it increasingly does not fully insulate the vendor either. Everyone in the chain has skin in the game now, which means everyone in the chain needs their own governance evidence, independently defensible, regardless of what the contract says about indemnification.
What this means for procurement, practically
If you take one operational change away from this post, make it this: AI vendor risk cannot live solely inside legal’s contract review process anymore. It needs a technical governance review layered on top, one that specifically asks:
What is this vendor’s provider-side documentation, and have we actually reviewed it, not just accepted it exists?
Are we fine-tuning, rebranding, or repurposing this system in any way that could trigger Article 25?
Do we have our own deployer-side obligations in place: human oversight, monitoring, logging, incident reporting, actually implemented, or just assumed to be the vendor’s problem?
If this system fails, who inside our organization owns the incident response, and have they rehearsed it?
None of these questions are hard to ask. What is hard is building the organizational muscle to ask them before deployment rather than during a post-incident scramble. That is where I am headed next in this series. What you as an organization need is the operational cycle that moves you on from figuring out “who is liable” from a legal perspective, to a continuous governance discipline you run to properly quantify and qualify your risk, and most importantly, the operational discipline to respond when the risk exceeds your enterprise’s risk appetite or tolerance. More on these topics in coming posts.
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
