---
source: "https://blog.emailengine.app/how-ai-changed-the-build-vs-buy-equation-in-saas/"
hn_url: "https://news.ycombinator.com/item?id=50246230"
title: "AI Changed the Build vs. Buy Equation in SaaS"
article_title: "How AI Changed the Build vs. Buy Equation in SaaS"
image: "https://blog.emailengine.app/content/images/size/w1200/2026/10/AI-to--o--laud_-ko--ik-u--hes-su--steemis.png"
author: "andris9"
captured_at: "2026-10-26T16:33:15Z"
capture_tool: "hn-digest"
hn_id: 50246230
score: 1
comments: 0
posted_at: "2026-10-26T16:12:05Z"
tags:
  - hacker-news
---

# AI Changed the Build vs. Buy Equation in SaaS

- HN: [50246230](https://news.ycombinator.com/item?id=50246230)
- Source: [blog.emailengine.app](https://blog.emailengine.app/how-ai-changed-the-build-vs-buy-equation-in-saas/)
- Score: 1
- Comments: 0
- Posted: 2026-10-26T16:12:05Z

## Translation

Title: AI Changed the Build vs. Buy Equation in SaaS
Article title: How AI Changed the Build vs. Buy Equation in SaaS
Description: Many old common-sense rules about software development no longer apply, or at least the reasoning behind them has changed quite a bit.
For example, the conventional advice when starting a new SaaS was not to implement a password reset function until you actually needed one. New users were unlikely

Article text:
How AI Changed the Build vs. Buy Equation in SaaS
Many old common-sense rules about software development no longer apply, or at least the reasoning behind them has changed quite a bit.
For example, the conventional advice when starting a new SaaS was not to implement a password reset function until you actually needed one. New users were unlikely to forget their passwords right away, and if someone did, you could handle the occasional support request manually. Implementing a proper password reset flow took development time, and that time was better spent on functionality that actually mattered.
This made perfect sense back then, but not so much anymore. Today, a password reset function is just another line item in an AI coding plan. It still needs testing, but it no longer takes away significant development time from the core product.
The same goes for the old advice to avoid building non-core features and use external SaaS services instead. Why spend precious development time on something you can buy for a small monthly fee? Think of testimonial widgets, lead forms, pricing tables, documentation chatbots, and so on.
But there is also a downside to using all these external services. Your data ends up scattered across different systems, and integrating everything with your existing backend often requires development work anyway. If you already need to build an integration, why not implement the entire feature in-house? With AI coding, it might just be another line item in the plan.
I've been doing exactly this for my own company . I've replaced several external SaaS services by letting AI add functionality to my existing backoffice system, which was originally developed long before I started using AI for coding. The landing page, testimonials, and even the AI chat on the documentation pages are now either part of that system or generated from it.
The cost was a few days of tweaking AI coding plans and testing the results. The savings are around 300 eur per month in SaaS subscriptions, plus better integration and, in some cases, better quality. For example, my own documentation chatbot now uses better and faster AI models for roughly the same cost as the external service it replaced.
I'm not saying everyone should start building everything themselves. But the economics of "build vs. buy" have changed significantly, and a lot of advice that used to be perfectly reasonable deserves to be reconsidered.
And before anyone points out that developing chatbots in-house can introduce security risks, I spent quite a bit of time working with Mythos to identify potential vulnerabilities and fix them. All incoming chat messages are screened for potential jailbreak attempts (users trying funny things get blocked), usage is capped by a daily budget (once the limit is reached, the chatbot simply stops responding), and so on. Building it yourself doesn't mean ignoring security. It just means you're responsible for it.
Kuidas ma kasutan AI-d töös, kus ma enam ise koodi ei kirjuta
Email is the hardest easy problem, and I built a business in it
AI Makes Documenting Small Open Source Projects Viable

## Original Extract

Many old common-sense rules about software development no longer apply, or at least the reasoning behind them has changed quite a bit.
For example, the conventional advice when starting a new SaaS was not to implement a password reset function until you actually needed one. New users were unlikely

How AI Changed the Build vs. Buy Equation in SaaS
Many old common-sense rules about software development no longer apply, or at least the reasoning behind them has changed quite a bit.
For example, the conventional advice when starting a new SaaS was not to implement a password reset function until you actually needed one. New users were unlikely to forget their passwords right away, and if someone did, you could handle the occasional support request manually. Implementing a proper password reset flow took development time, and that time was better spent on functionality that actually mattered.
This made perfect sense back then, but not so much anymore. Today, a password reset function is just another line item in an AI coding plan. It still needs testing, but it no longer takes away significant development time from the core product.
The same goes for the old advice to avoid building non-core features and use external SaaS services instead. Why spend precious development time on something you can buy for a small monthly fee? Think of testimonial widgets, lead forms, pricing tables, documentation chatbots, and so on.
But there is also a downside to using all these external services. Your data ends up scattered across different systems, and integrating everything with your existing backend often requires development work anyway. If you already need to build an integration, why not implement the entire feature in-house? With AI coding, it might just be another line item in the plan.
I've been doing exactly this for my own company . I've replaced several external SaaS services by letting AI add functionality to my existing backoffice system, which was originally developed long before I started using AI for coding. The landing page, testimonials, and even the AI chat on the documentation pages are now either part of that system or generated from it.
The cost was a few days of tweaking AI coding plans and testing the results. The savings are around 300 eur per month in SaaS subscriptions, plus better integration and, in some cases, better quality. For example, my own documentation chatbot now uses better and faster AI models for roughly the same cost as the external service it replaced.
I'm not saying everyone should start building everything themselves. But the economics of "build vs. buy" have changed significantly, and a lot of advice that used to be perfectly reasonable deserves to be reconsidered.
And before anyone points out that developing chatbots in-house can introduce security risks, I spent quite a bit of time working with Mythos to identify potential vulnerabilities and fix them. All incoming chat messages are screened for potential jailbreak attempts (users trying funny things get blocked), usage is capped by a daily budget (once the limit is reached, the chatbot simply stops responding), and so on. Building it yourself doesn't mean ignoring security. It just means you're responsible for it.
Kuidas ma kasutan AI-d töös, kus ma enam ise koodi ei kirjuta
Email is the hardest easy problem, and I built a business in it
AI Makes Documenting Small Open Source Projects Viable
