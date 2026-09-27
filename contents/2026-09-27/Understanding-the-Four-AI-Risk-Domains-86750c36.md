---
source: "https://sdarchitect.blog/2026/09/27/ai-risk-a-users-guide-part-iv-understanding-the-four-ai-risk-domains/"
hn_url: "https://news.ycombinator.com/item?id=49866122"
title: "Understanding the Four AI Risk Domains"
article_title: "AI Risk – A user’s guide (Part IV): Understanding the Four AI Risk Domains – Home: sdarchitect.blog"
image: "https://i0.wp.com/sdarchitect.blog/wp-content/uploads/2026/09/generated-image-8.png?fit=1200%2C675&ssl=1"
author: "SanjeevSharma"
captured_at: "2026-09-27T12:48:58Z"
capture_tool: "hn-digest"
hn_id: 49866122
score: 1
comments: 0
posted_at: "2026-09-27T12:30:59Z"
tags:
  - hacker-news
---

# Understanding the Four AI Risk Domains

- HN: [49866122](https://news.ycombinator.com/item?id=49866122)
- Source: [sdarchitect.blog](https://sdarchitect.blog/2026/09/27/ai-risk-a-users-guide-part-iv-understanding-the-four-ai-risk-domains/)
- Score: 1
- Comments: 0
- Posted: 2026-09-27T12:30:59Z

## Translation

Title: Understanding the Four AI Risk Domains
Article title: AI Risk – A user’s guide (Part IV): Understanding the Four AI Risk Domains – Home: sdarchitect.blog
Description: Whenever I ask a room of CIOs and CISOs to list the AI risks they are actively managing, I get answers clustered heavily in one area, almost always technical: hallucination, model drift, the occasional mention of robustness. What I rarely hear, unprompted, is a comprehensive answer that spans al the
[truncated]

Article text:
AI Risk – A user’s guide (Part IV): Understanding the Four AI Risk Domains – Home: sdarchitect.blog
Skip to content
Home: sdarchitect.blog
Sanjeev Sharma: My thoughts on Cloud, DevOps, Data Strategy, and Life…
AI Risk – A user’s guide (Part IV): Understanding the Four AI Risk Domains
Whenever I ask a room of CIOs and CISOs to list the AI risks they are actively managing, I get answers clustered heavily in one area, almost always technical: hallucination, model drift, the occasional mention of robustness. What I rarely hear, unprompted, is a comprehensive answer that spans al thel four domains that actually matter. So let me lay them out explicitly, the way I did at the AI Risk Summit last month, because I think the gaps between what leaders are watching and what they should be watching are exactly where the next round of incidents will come from.
Technical: the domain everyone already ‘half-knows’
Technical risk covers hallucination, model drift, robustness failures, and unexplainability. This is the domain that gets the most airtime, partly because it’s the most intuitive to a technical audience and partly because it maps cleanly onto things engineers already know how to measure: accuracy metrics, confidence scores, test suites.
But I would push back gently on treating this as the primary risk domain just because it is the most familiar. Unexplainability, specifically, deserves more weight than it usually gets. A model that is accurate 95% of the time but can’t explain why it made any individual decision is a serious liability in any regulated or high-stakes context, because “trust me, it’s usually right” is not a defensible position in front of a regulator, a customer, or certainly a courtroom.
Societal and ethical: the domain that erodes trust slowly
Societal and ethical risk covers harmful bias, privacy violations, manipulation, and autonomy erosion. I want to focus on that last item: autonomy erosion because it is the one I think gets the least attention and deserves the most.
Autonomy erosion is what happens when AI systems, cumulatively, start making decisions that quietly narrow human choice without anyone explicitly deciding that should happen. A recommendation engine that increasingly determines what content, products, or opportunities a person even sees is exercising a form of control that is easy to miss because no single decision looks dramatic. This is a slow-burn risk, and slow-burn risks are exactly the kind that governance programs built around discrete incident response tend to miss, because there is no single moment that triggers an alert.
Operational and organizational: the domain hiding in your org chart
Operational and organizational risk covers accountability gaps, third-party dependencies, and over-reliance. I have written about third-party dependency already in my last post in the series , the liability chain that runs from provider to deployer, but accountability gaps and over-reliance deserve their own attention here.
Accountability gaps happen when an AI system’s decisions don’t map cleanly onto any single owner in the org chart. Who is accountable when a model embedded across three different teams’ workflows produces a bad outcome? If your honest answer is “it’s unclear,” you have an operational risk sitting right there, independent of anything the model itself is doing wrong technically.
Over-reliance is the subtler cousin of this problem: teams that stop exercising independent judgment because the AI system has been reliable enough, often enough, that skepticism starts to feel like friction. I have watched this happen with automation before AI existed too. It is a human factors problem as old as automation itself, just accelerated by how competent generative AI systems feel in day-to-day use.
Adversarial and security: the domain growing eight times faster than anything else
Adversarial and security risk covers prompt injection, data poisoning, exfiltration, and AI-enabled attacks. And here’s the number that should reorder every security team’s priority list: adversarial exploitation of AI grew eightfold from 2022 to 2025, making it the fastest-growing risk category of the four, by a wide margin.
I think this growth rate reflects something structural about how attackers operate: they go where the new attack surface is, and AI systems, especially agentic ones with tool access and the ability to take real-world actions, represent an enormous, largely unhardened new attack surface. Prompt injection specifically deserves board-level attention because it is not a theoretical vulnerability; it’s a practical technique that’s been demonstrated repeatedly against production systems, and it gets more dangerous, not less, as we give AI agents more autonomy and more tool access in the name of productivity.
This is a topic close to home for me given my work on autonomous DevOps and SRE agents. An agent with the ability to modify infrastructure, restart services, or roll back deployments is exactly the kind of high-value target that adversarial exploitation growth should make every platform engineering leader nervous about. The more capability you give an agent, the more valuable it becomes as an attack target, and the more catastrophic a successful prompt injection or data poisoning attack becomes.
Why you need all four, not your favorite one or two
The organizations I see struggling most with AI governance aren’t the ones ignoring risk entirely, they are the ones who have built real competence in one or two of these domains and quietly assumed that competence generalizes. A team with excellent model eval practices covering technical risk can still be completely exposed on the adversarial and security domain, because the skills, tools, and mindset required to red-team a model for prompt injection are genuinely different from the skills required to measure its accuracy or fairness.
My advice: use these four domains as a literal audit checklist for your own AI governance program, domain by domain, system by system. Don’t let strength in one area create false confidence about the others. The eightfold growth in adversarial exploitation alone should be reason enough to make sure the domain you are least comfortable with isn’t the one you’re least prepared for.
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
Join 160 other subscribers
Sign me up
Have a WordPress.com account? Log in now.
Home: sdarchitect.blog Copy shortlink View post in Reader Manage subscriptions Sign up Log in Report this content Collapse this bar

## Original Extract

Whenever I ask a room of CIOs and CISOs to list the AI risks they are actively managing, I get answers clustered heavily in one area, almost always technical: hallucination, model drift, the occasional mention of robustness. What I rarely hear, unprompted, is a comprehensive answer that spans al the
[truncated]

AI Risk – A user’s guide (Part IV): Understanding the Four AI Risk Domains – Home: sdarchitect.blog
Skip to content
Home: sdarchitect.blog
Sanjeev Sharma: My thoughts on Cloud, DevOps, Data Strategy, and Life…
AI Risk – A user’s guide (Part IV): Understanding the Four AI Risk Domains
Whenever I ask a room of CIOs and CISOs to list the AI risks they are actively managing, I get answers clustered heavily in one area, almost always technical: hallucination, model drift, the occasional mention of robustness. What I rarely hear, unprompted, is a comprehensive answer that spans al thel four domains that actually matter. So let me lay them out explicitly, the way I did at the AI Risk Summit last month, because I think the gaps between what leaders are watching and what they should be watching are exactly where the next round of incidents will come from.
Technical: the domain everyone already ‘half-knows’
Technical risk covers hallucination, model drift, robustness failures, and unexplainability. This is the domain that gets the most airtime, partly because it’s the most intuitive to a technical audience and partly because it maps cleanly onto things engineers already know how to measure: accuracy metrics, confidence scores, test suites.
But I would push back gently on treating this as the primary risk domain just because it is the most familiar. Unexplainability, specifically, deserves more weight than it usually gets. A model that is accurate 95% of the time but can’t explain why it made any individual decision is a serious liability in any regulated or high-stakes context, because “trust me, it’s usually right” is not a defensible position in front of a regulator, a customer, or certainly a courtroom.
Societal and ethical: the domain that erodes trust slowly
Societal and ethical risk covers harmful bias, privacy violations, manipulation, and autonomy erosion. I want to focus on that last item: autonomy erosion because it is the one I think gets the least attention and deserves the most.
Autonomy erosion is what happens when AI systems, cumulatively, start making decisions that quietly narrow human choice without anyone explicitly deciding that should happen. A recommendation engine that increasingly determines what content, products, or opportunities a person even sees is exercising a form of control that is easy to miss because no single decision looks dramatic. This is a slow-burn risk, and slow-burn risks are exactly the kind that governance programs built around discrete incident response tend to miss, because there is no single moment that triggers an alert.
Operational and organizational: the domain hiding in your org chart
Operational and organizational risk covers accountability gaps, third-party dependencies, and over-reliance. I have written about third-party dependency already in my last post in the series , the liability chain that runs from provider to deployer, but accountability gaps and over-reliance deserve their own attention here.
Accountability gaps happen when an AI system’s decisions don’t map cleanly onto any single owner in the org chart. Who is accountable when a model embedded across three different teams’ workflows produces a bad outcome? If your honest answer is “it’s unclear,” you have an operational risk sitting right there, independent of anything the model itself is doing wrong technically.
Over-reliance is the subtler cousin of this problem: teams that stop exercising independent judgment because the AI system has been reliable enough, often enough, that skepticism starts to feel like friction. I have watched this happen with automation before AI existed too. It is a human factors problem as old as automation itself, just accelerated by how competent generative AI systems feel in day-to-day use.
Adversarial and security: the domain growing eight times faster than anything else
Adversarial and security risk covers prompt injection, data poisoning, exfiltration, and AI-enabled attacks. And here’s the number that should reorder every security team’s priority list: adversarial exploitation of AI grew eightfold from 2022 to 2025, making it the fastest-growing risk category of the four, by a wide margin.
I think this growth rate reflects something structural about how attackers operate: they go where the new attack surface is, and AI systems, especially agentic ones with tool access and the ability to take real-world actions, represent an enormous, largely unhardened new attack surface. Prompt injection specifically deserves board-level attention because it is not a theoretical vulnerability; it’s a practical technique that’s been demonstrated repeatedly against production systems, and it gets more dangerous, not less, as we give AI agents more autonomy and more tool access in the name of productivity.
This is a topic close to home for me given my work on autonomous DevOps and SRE agents. An agent with the ability to modify infrastructure, restart services, or roll back deployments is exactly the kind of high-value target that adversarial exploitation growth should make every platform engineering leader nervous about. The more capability you give an agent, the more valuable it becomes as an attack target, and the more catastrophic a successful prompt injection or data poisoning attack becomes.
Why you need all four, not your favorite one or two
The organizations I see struggling most with AI governance aren’t the ones ignoring risk entirely, they are the ones who have built real competence in one or two of these domains and quietly assumed that competence generalizes. A team with excellent model eval practices covering technical risk can still be completely exposed on the adversarial and security domain, because the skills, tools, and mindset required to red-team a model for prompt injection are genuinely different from the skills required to measure its accuracy or fairness.
My advice: use these four domains as a literal audit checklist for your own AI governance program, domain by domain, system by system. Don’t let strength in one area create false confidence about the others. The eightfold growth in adversarial exploitation alone should be reason enough to make sure the domain you are least comfortable with isn’t the one you’re least prepared for.
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
Join 160 other subscribers
Sign me up
Have a WordPress.com account? Log in now.
Home: sdarchitect.blog Copy shortlink View post in Reader Manage subscriptions Sign up Log in Report this content Collapse this bar
