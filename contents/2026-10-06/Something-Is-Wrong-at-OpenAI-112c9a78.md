---
source: "https://www.vincentschmalbach.com/something-wrong-at-openai/"
hn_url: "https://news.ycombinator.com/item?id=49974751"
title: "Something Is Wrong at OpenAI"
article_title: "Something Is Wrong at OpenAI - Vincent Schmalbach"
image: "https://www.vincentschmalbach.com/wp-content/uploads/2021/05/Vincent-Schmalbach-draft01.jpg"
author: "vincent_s"
captured_at: "2026-10-06T06:07:15Z"
capture_tool: "hn-digest"
hn_id: 49974751
score: 1
comments: 0
posted_at: "2026-10-06T05:57:48Z"
tags:
  - hacker-news
---

# Something Is Wrong at OpenAI

- HN: [49974751](https://news.ycombinator.com/item?id=49974751)
- Source: [www.vincentschmalbach.com](https://www.vincentschmalbach.com/something-wrong-at-openai/)
- Score: 1
- Comments: 0
- Posted: 2026-10-06T05:57:48Z

## Translation

Title: Something Is Wrong at OpenAI
Article title: Something Is Wrong at OpenAI - Vincent Schmalbach
Description: Something is wrong at OpenAI. Each model it has released since GPT-5.5 is smarter on paper and worse to work with, and GPT-6.1 Sol is the worst so far.

Article text:
Something Is Wrong at OpenAI - Vincent Schmalbach
Skip to content
Software Developer
Home
October 5, 2026
by Vincent Schmalbach
Something is wrong at OpenAI. Each model it has released since GPT-5.5 is smarter on paper and worse to work with, and GPT-6.1 Sol is the worst so far. OpenAI says 6.1 Sol is "built for complex refactors, deep codebase investigations, and long-running agents across apps." In my sessions it works for a day or more on things I never asked for. On the one task where I want it to keep going all day, it stops after 15 or 20 minutes. GPT-5.5 is the last OpenAI model that still works for me, and it is being removed from the Codex subscription next week, on October 14.
I'm writing this while working on vroni.com , my AI coding tool for turning tasks into pull requests.
I have written about this since August. Switching from GPT-5.5 to GPT-5.6 made me less productive , GPT-5.6 Sol used 2.25x the tokens of GPT-5.5 , and GPT-6 Astra never knew when to stop either. Going back to GPT-5.6 Sol does not help, because it has the same problems. Back then my complaint was mostly time and quota. With GPT-6 Sol and 6.1 Sol the work itself is going into the wrong direction. In about 50% of my chat they work rather well, in the other 50% they go down rabbit holes and work on things I never asked for.
One example is a building of an assistant that should clean up my inbox by archiving mail automatically. Sol worked about 100 hours on it, once for almost 36 hours straight, and used about 4.15 billion tokens with its subagents. The result is 150 commits, about 53,500 lines of Python and 834 tests. First time it even looked at my mailbox was four and a half days in, and its first run found nothing to archive. After a week it had not archived a single email. Then I asked it to build a rule what can be auto-archived, gave it some explanations and pasted six emails that should obviously go as examples, like old login links and outdated new-device notices. It added those six senders' domains to an allowlist, wrote a test that checks those exact subject lines, archived the six and reported "All six emails are archived." That's a complete joke, a total waste of my time (and tokens). And this little email assistant I wanted it to build is now a gigantic code base and absolutely unusable and does nothing. This is stuff that GPT-5.5 is perfectly capable doing, and if you give the same task to Opus 5.5 it nails it and produces a production ready tool. As I said above, stuff like this happens in about 50% of cases, rather randomly. So I do have some other sessions where it just does what I want and it's quite good at it.
In one of my SaaS products I gave it a free hand to build the remaining features, that were already planned in some markdown files. After more than 18 hours, none of the 15 feature projects has started. It wrote a 768-line JSON policy file to accept one warning in a dev-only Tailwind dependency. It also added a security rule for some outgoing connections. And then Sol wrote its own SOCKS5 proxy, 733 lines across 11 files, in about four hours. I did not ask for a SOCKS5 proxy, this has nothing to do with the actual tasks at hand.
The one session where I want it to run all day I told it 33 times to keep going, with messages like "ok go on never stop" and "I told you to go on and on and on and never stop." 6.1 Sol ended each turn after 11 to 32 minutes with a status report.
It is no better as a reviewer. In one of my projects Claude Opus fixed some cookie related stuff, and 6.1 Sol reviewed the outcome. My rule was to keep going until Sol had nothing left to complain about. Big mistake. After 19 rounds of fixing what it complained about it still had something. The result was an insanely overengineered solution to problems Sol more or less made up.
I am not the only one. Another Codex user reports on GitHub that tasks that took about 20 minutes with GPT-5.6 Sol take around 50 with GPT-6 Sol. Another describes Codex expanding scope, creating its own work and not stopping when told.
Right now the only way I get useful work out of Sol is a second Claude Opus session that watches what Sol does and corrects it. I keep running Sol because OpenAI gave me credits and both my accounts still have usage resets left. Even burning free tokens is hard with this model. GPT-5.5 is the only OpenAI model I still trust, so I will use it until it retires on October 14 . After that I will move to Claude completely.
Give Vroni a GitHub issue, bug report, spec, or rough idea. It reads the repo, plans the change, writes code, runs checks, and works toward a review-ready pull request.
Usually a new article and a few links I found interesting.
No spam. Unsubscribe with one click.
More writing around this topic.
Anyone Can Build an App Now, but Who Is Going to Pay for It?
With Claude Code or Cursor, a developer can build a booking app for hair salons over a weekend. It can have online…
Codex Users Are Fleeing to Claude Code
Codex users are leaving for Claude Code. OpenAI cut the Pro 200 plan from 20x to 10x the Plus allowance at DevDay…
How to Make Videos with Claude Code
Claude Code can't generate video directly. It isn't a video model like Veo or Sora. It writes text and code. You can…
Your email address will not be published. Required fields are marked *
I write about software, AI, Laravel, SEO, and what it takes to build online products. In 2025, 73k people read the blog. Some posts were picked up by publications and developer communities.
A few older posts got the most attention. I keep them here, separate from the newest writing below.
Google Now Defaults to Not Indexing Your Content
The post that led to a Guardian column and a long developer-community discussion.
AI Exponentializes Your Tech Debt
Why AI makes weak engineering foundations more expensive, not less.
Only Experts Can Write Good Prompts
Good AI work still depends on technical judgment and taste.
Specific Laravel advice for people working in real codebases.
The newest posts from the blog. This updates automatically when I publish.
Something is wrong at OpenAI. Each model it has released since GPT-5.5 is smarter on paper and worse...
Anyone Can Build an App Now, but Who Is Going to Pay for It?
With Claude Code or Cursor, a developer can build a booking app for hair salons over a weekend....
Codex Users Are Fleeing to Claude Code
Codex users are leaving for Claude Code. OpenAI cut the Pro 200 plan from 20x to 10x the...
Places that have mentioned, quoted, or discussed my writing and commentary.
Actual feedback from clients I've worked with.
"HIRE VINCENT. Seriously. I've worked with many many developers over the years (remember rentacoder?), and Vincent is that rare breed of web developer who just gets it. You have a chat, you express what you need, and he comes back with what you asked for. His english is fluent, he understands what you need, and he's not looking to overcomplicate things. It's a joy and pleasure to work with him, not to have to micromanage and just overall a great experience. Honestly I wish I could keep him on we just don't have the budget anymore - that is the only reason the contract ended."
"What stood out about working with Vincent was how methodical and efficient he was. He started by really digging into our requirements, asking clarifying questions until he had a solid understanding. He then built exactly what was needed and kept us informed throughout. Communication was clear and proactive at all times. Everything stayed on track, both in terms of time and budget. The end result was excellent and has been running smoothly since launch. It's that combination of practical focus and attention to detail that makes the difference. I would definitely recommend Vincent for quality development."
"Having Vincent on retainer has been excellent. He's tackled some major technical challenges for us, notably significantly improving our website's performance. He also expertly managed a complex upgrade of our core platform (Laravel) to the latest version, ensuring we stay current and secure. We're now rolling out the multilingual capabilities he built, which is exciting. Vincent is reliable, technically skilled, and a valuable partner for keeping our site running optimally and moving forward. Highly recommend his expertise."
"I've been working with Vincent for several months, and he's one of the most independent developers I've collaborated with. He doesn't need micromanagement - he's fully capable of making decisions and driving the project forward, which has freed me up to focus on other parts of the business.
Vincent works very well asynchronously: communication stays focused, there is little unnecessary back-and-forth, and he still raises the important questions when needed. The result is a developer who delivers without overhead, and that's been genuinely valuable."
"We work with Vincent regularly on Laravel/Vue projects, both new builds and existing codebases that have grown over the years. In both cases, he delivers clean code and reliable implementation. We especially value how quickly he gets into mature projects and finds solutions that fit the existing state of the codebase. Absolutely recommended."
"Vincent was easy to work with and produced high-quality code. Would recommend."
"Vincent is one of those rare developers who can jump into a complex codebase, become useful immediately, and solve real problems without hand-holding.
He worked on performance optimization, customer-specific features, and a major custom reporting system. He also turned customer discussions into clear technical requirements and documented his work so the rest of the team could build on it.
Vincent does not just close tickets, he thinks through the product, the edge cases, and the customer impact. If you need a senior Laravel developer who can move a SaaS product forward, I strongly recommend him."
Have something I should look at?
Send what you have. I can turn it into a plan and then build it.
My Book: Rapid SaaS with Laravel
Launch Your SaaS in Days, Not Months. A practical, no-BS guide to building production-ready SaaS apps with Laravel 12: speed, billing, teams, AI integration, and the parts you need before customers can pay you.
Vincent Schmalbach is an entrepreneur, software developer, and SEO expert. He works on SaaS products, business apps, automation, and technical project rescue, with 15+ years of software and online business experience.
Posts
Something Is Wrong at OpenAI
Anyone Can Build an App Now, but Who Is Going to Pay for It?
Codex Users Are Fleeing to Claude Code
How to Make Videos with Claude Code
OpenAI Is Paying Me to Move My Business to Anthropic
OpenAI Codex Is Retiring the Last Model that Behaves Like an Assistant
ChatGPT Plus and Pro Will Get Ads in 2027
The Economic Incentives Behind the Push to Pause AI
The Cost of AI Is Now My Main Blocker
Why I Am Bullish on ChatGPT Ads
We Urgently Need a New Major Search Engine
Setting Up Pi With DeepSeek V4.1 Flash on OpenRouter
Node.js and TypeScript Development
Customer Portals, Dashboards and Internal Tools
Software Development Company Alternative
Software Audit and Code Review
Custom Laravel App Development
Laravel Development Consulting
Technical Team Leadership & Mentoring
Laravel Performance Optimization
Laravel Application Maintenance & Support
|
Resources | Privacy Policy |
Imprint
I mostly send new articles and a few links I found interesting. Sometimes it's something else.
No spam. Unsubscribe with one click.

## Original Extract

Something is wrong at OpenAI. Each model it has released since GPT-5.5 is smarter on paper and worse to work with, and GPT-6.1 Sol is the worst so far.

Something Is Wrong at OpenAI - Vincent Schmalbach
Skip to content
Software Developer
Home
October 5, 2026
by Vincent Schmalbach
Something is wrong at OpenAI. Each model it has released since GPT-5.5 is smarter on paper and worse to work with, and GPT-6.1 Sol is the worst so far. OpenAI says 6.1 Sol is "built for complex refactors, deep codebase investigations, and long-running agents across apps." In my sessions it works for a day or more on things I never asked for. On the one task where I want it to keep going all day, it stops after 15 or 20 minutes. GPT-5.5 is the last OpenAI model that still works for me, and it is being removed from the Codex subscription next week, on October 14.
I'm writing this while working on vroni.com , my AI coding tool for turning tasks into pull requests.
I have written about this since August. Switching from GPT-5.5 to GPT-5.6 made me less productive , GPT-5.6 Sol used 2.25x the tokens of GPT-5.5 , and GPT-6 Astra never knew when to stop either. Going back to GPT-5.6 Sol does not help, because it has the same problems. Back then my complaint was mostly time and quota. With GPT-6 Sol and 6.1 Sol the work itself is going into the wrong direction. In about 50% of my chat they work rather well, in the other 50% they go down rabbit holes and work on things I never asked for.
One example is a building of an assistant that should clean up my inbox by archiving mail automatically. Sol worked about 100 hours on it, once for almost 36 hours straight, and used about 4.15 billion tokens with its subagents. The result is 150 commits, about 53,500 lines of Python and 834 tests. First time it even looked at my mailbox was four and a half days in, and its first run found nothing to archive. After a week it had not archived a single email. Then I asked it to build a rule what can be auto-archived, gave it some explanations and pasted six emails that should obviously go as examples, like old login links and outdated new-device notices. It added those six senders' domains to an allowlist, wrote a test that checks those exact subject lines, archived the six and reported "All six emails are archived." That's a complete joke, a total waste of my time (and tokens). And this little email assistant I wanted it to build is now a gigantic code base and absolutely unusable and does nothing. This is stuff that GPT-5.5 is perfectly capable doing, and if you give the same task to Opus 5.5 it nails it and produces a production ready tool. As I said above, stuff like this happens in about 50% of cases, rather randomly. So I do have some other sessions where it just does what I want and it's quite good at it.
In one of my SaaS products I gave it a free hand to build the remaining features, that were already planned in some markdown files. After more than 18 hours, none of the 15 feature projects has started. It wrote a 768-line JSON policy file to accept one warning in a dev-only Tailwind dependency. It also added a security rule for some outgoing connections. And then Sol wrote its own SOCKS5 proxy, 733 lines across 11 files, in about four hours. I did not ask for a SOCKS5 proxy, this has nothing to do with the actual tasks at hand.
The one session where I want it to run all day I told it 33 times to keep going, with messages like "ok go on never stop" and "I told you to go on and on and on and never stop." 6.1 Sol ended each turn after 11 to 32 minutes with a status report.
It is no better as a reviewer. In one of my projects Claude Opus fixed some cookie related stuff, and 6.1 Sol reviewed the outcome. My rule was to keep going until Sol had nothing left to complain about. Big mistake. After 19 rounds of fixing what it complained about it still had something. The result was an insanely overengineered solution to problems Sol more or less made up.
I am not the only one. Another Codex user reports on GitHub that tasks that took about 20 minutes with GPT-5.6 Sol take around 50 with GPT-6 Sol. Another describes Codex expanding scope, creating its own work and not stopping when told.
Right now the only way I get useful work out of Sol is a second Claude Opus session that watches what Sol does and corrects it. I keep running Sol because OpenAI gave me credits and both my accounts still have usage resets left. Even burning free tokens is hard with this model. GPT-5.5 is the only OpenAI model I still trust, so I will use it until it retires on October 14 . After that I will move to Claude completely.
Give Vroni a GitHub issue, bug report, spec, or rough idea. It reads the repo, plans the change, writes code, runs checks, and works toward a review-ready pull request.
Usually a new article and a few links I found interesting.
No spam. Unsubscribe with one click.
More writing around this topic.
Anyone Can Build an App Now, but Who Is Going to Pay for It?
With Claude Code or Cursor, a developer can build a booking app for hair salons over a weekend. It can have online…
Codex Users Are Fleeing to Claude Code
Codex users are leaving for Claude Code. OpenAI cut the Pro 200 plan from 20x to 10x the Plus allowance at DevDay…
How to Make Videos with Claude Code
Claude Code can't generate video directly. It isn't a video model like Veo or Sora. It writes text and code. You can…
Your email address will not be published. Required fields are marked *
I write about software, AI, Laravel, SEO, and what it takes to build online products. In 2025, 73k people read the blog. Some posts were picked up by publications and developer communities.
A few older posts got the most attention. I keep them here, separate from the newest writing below.
Google Now Defaults to Not Indexing Your Content
The post that led to a Guardian column and a long developer-community discussion.
AI Exponentializes Your Tech Debt
Why AI makes weak engineering foundations more expensive, not less.
Only Experts Can Write Good Prompts
Good AI work still depends on technical judgment and taste.
Specific Laravel advice for people working in real codebases.
The newest posts from the blog. This updates automatically when I publish.
Something is wrong at OpenAI. Each model it has released since GPT-5.5 is smarter on paper and worse...
Anyone Can Build an App Now, but Who Is Going to Pay for It?
With Claude Code or Cursor, a developer can build a booking app for hair salons over a weekend....
Codex Users Are Fleeing to Claude Code
Codex users are leaving for Claude Code. OpenAI cut the Pro 200 plan from 20x to 10x the...
Places that have mentioned, quoted, or discussed my writing and commentary.
Actual feedback from clients I've worked with.
"HIRE VINCENT. Seriously. I've worked with many many developers over the years (remember rentacoder?), and Vincent is that rare breed of web developer who just gets it. You have a chat, you express what you need, and he comes back with what you asked for. His english is fluent, he understands what you need, and he's not looking to overcomplicate things. It's a joy and pleasure to work with him, not to have to micromanage and just overall a great experience. Honestly I wish I could keep him on we just don't have the budget anymore - that is the only reason the contract ended."
"What stood out about working with Vincent was how methodical and efficient he was. He started by really digging into our requirements, asking clarifying questions until he had a solid understanding. He then built exactly what was needed and kept us informed throughout. Communication was clear and proactive at all times. Everything stayed on track, both in terms of time and budget. The end result was excellent and has been running smoothly since launch. It's that combination of practical focus and attention to detail that makes the difference. I would definitely recommend Vincent for quality development."
"Having Vincent on retainer has been excellent. He's tackled some major technical challenges for us, notably significantly improving our website's performance. He also expertly managed a complex upgrade of our core platform (Laravel) to the latest version, ensuring we stay current and secure. We're now rolling out the multilingual capabilities he built, which is exciting. Vincent is reliable, technically skilled, and a valuable partner for keeping our site running optimally and moving forward. Highly recommend his expertise."
"I've been working with Vincent for several months, and he's one of the most independent developers I've collaborated with. He doesn't need micromanagement - he's fully capable of making decisions and driving the project forward, which has freed me up to focus on other parts of the business.
Vincent works very well asynchronously: communication stays focused, there is little unnecessary back-and-forth, and he still raises the important questions when needed. The result is a developer who delivers without overhead, and that's been genuinely valuable."
"We work with Vincent regularly on Laravel/Vue projects, both new builds and existing codebases that have grown over the years. In both cases, he delivers clean code and reliable implementation. We especially value how quickly he gets into mature projects and finds solutions that fit the existing state of the codebase. Absolutely recommended."
"Vincent was easy to work with and produced high-quality code. Would recommend."
"Vincent is one of those rare developers who can jump into a complex codebase, become useful immediately, and solve real problems without hand-holding.
He worked on performance optimization, customer-specific features, and a major custom reporting system. He also turned customer discussions into clear technical requirements and documented his work so the rest of the team could build on it.
Vincent does not just close tickets, he thinks through the product, the edge cases, and the customer impact. If you need a senior Laravel developer who can move a SaaS product forward, I strongly recommend him."
Have something I should look at?
Send what you have. I can turn it into a plan and then build it.
My Book: Rapid SaaS with Laravel
Launch Your SaaS in Days, Not Months. A practical, no-BS guide to building production-ready SaaS apps with Laravel 12: speed, billing, teams, AI integration, and the parts you need before customers can pay you.
Vincent Schmalbach is an entrepreneur, software developer, and SEO expert. He works on SaaS products, business apps, automation, and technical project rescue, with 15+ years of software and online business experience.
Posts
Something Is Wrong at OpenAI
Anyone Can Build an App Now, but Who Is Going to Pay for It?
Codex Users Are Fleeing to Claude Code
How to Make Videos with Claude Code
OpenAI Is Paying Me to Move My Business to Anthropic
OpenAI Codex Is Retiring the Last Model that Behaves Like an Assistant
ChatGPT Plus and Pro Will Get Ads in 2027
The Economic Incentives Behind the Push to Pause AI
The Cost of AI Is Now My Main Blocker
Why I Am Bullish on ChatGPT Ads
We Urgently Need a New Major Search Engine
Setting Up Pi With DeepSeek V4.1 Flash on OpenRouter
Node.js and TypeScript Development
Customer Portals, Dashboards and Internal Tools
Software Development Company Alternative
Software Audit and Code Review
Custom Laravel App Development
Laravel Development Consulting
Technical Team Leadership & Mentoring
Laravel Performance Optimization
Laravel Application Maintenance & Support
|
Resources | Privacy Policy |
Imprint
I mostly send new articles and a few links I found interesting. Sometimes it's something else.
No spam. Unsubscribe with one click.
