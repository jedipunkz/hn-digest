---
source: "https://thoughtbot.com/blog/what-causes-ai-fatigue-for-engineers-and-how-to-ease-the-strain"
hn_url: "https://news.ycombinator.com/item?id=50111887"
title: "What causes AI fatigue for engineers and how to ease the strain"
article_title: "What causes AI fatigue for engineers and how to ease the strain"
image: "https://images.prismic.io/thoughtbot-website/aEccKLh8WN-LV5_o_development-article-background.png?auto=format%2Ccompress&blend-mode=normal&blend-x=0&blend-y=0&blend=https%3A%2F%2Fimages.prismic.io%2Fthoughtbot-website%2FZ1tQFpbqstJ98cGT_development.png&mark-fit=max&mark-h=400&mark-w=800&mark-x=356&mark-y=100&mark64=aHR0cHM6Ly9hc3NldHMuaW1naXgubmV0L350ZXh0Lz90eHQtY2xyPWY1ZjVmNSZ0eHQtZm9udD1JQk1QbGV4U2Fuc0pQLVNlbWlCb2xkJnR4dC1sZWFkPTAmdHh0LXBhZD0wJnR4dC1zaXplPTY0JnR4dC10cmFjaz0wJnR4dDY0PVYyaGhkQ0JqWVhWelpYTWdRVWtnWm1GMGFXZDFaU0JtYjNJZ1pXNW5hVzVsWlhKeklHRnVaQ0JvYjNjZ2RHOGdaV0Z6WlNCMGFHVWdjM1J5WVdsdSZ3PTgwMA%3D%3D&txt-align=center%2Cmiddle&txt-color=f5f5f5&txt-fit=max&txt-font=IBMPlexSansJP-SemiBold&txt-size=24&txt-x=391&txt-y=526&txt=Richard+Newman%2C+et+al"
author: "devonnull"
captured_at: "2026-10-16T00:12:35Z"
capture_tool: "hn-digest"
hn_id: 50111887
score: 4
comments: 0
posted_at: "2026-10-15T23:43:23Z"
tags:
  - hacker-news
---

# What causes AI fatigue for engineers and how to ease the strain

- HN: [50111887](https://news.ycombinator.com/item?id=50111887)
- Source: [thoughtbot.com](https://thoughtbot.com/blog/what-causes-ai-fatigue-for-engineers-and-how-to-ease-the-strain)
- Score: 4
- Comments: 0
- Posted: 2026-10-15T23:43:23Z

## Translation

Title: What causes AI fatigue for engineers and how to ease the strain
Description: AI fatigue is setting in for engineering teams. Our research shows most feel pressure to assess every AI trend, struggle to evaluate tools, and wrestle with evaluating AI-generated code. Here’s some practical advice for navigating the pain points.

Article text:
Skip to main content
thoughtbot
thoughtbot
Close Menu
We're live
Live on Twitch!
thoughtbot is livestreaming
Work alongside the thoughtbot team as we collaborate with each other and our clients, live. Ask us anything, we're live right now!
Work alongside the thoughtbot team as we collaborate with each other and our clients, live. Ask us anything, we're live right now!
Menu
Let’s get started!
Back
View all Services
Development
Ruby on Rails
Back
View all Resources
Development
Tech Leadership Resources
Giant Robots Smashing Into Other Giant Robots Podcast
What causes AI fatigue for engineers and how to ease the strain
Copy as Markdown
AI fatigue is hitting some development teams hard. In our original research study , engineering, product and technology leaders from US healthcare companies told us they’re feeling pressured to evaluate a firehose of new AI trends and tools while struggling to navigate AI-generated code alongside legacy systems.
This isn’t entirely new. The industry has gone through waves like this before, from the rise of the internet and the advent of social media to mobile development and every new round of web technology in between. But the rapid progression of LLMs speeds up this cycle, which can leave teams to feel like they’re losing control.
Here’s what our study uncovered about the causes of AI fatigue and some practical ideas for successfully navigating these pain points.
73% feel pressure to evaluate every AI trend
Nearly three-quarters of survey respondents told us they feel significant or moderate pressure to evaluate every new AI trend despite limited bandwidth.
So, what’s a developer or product manager to do when the tenth email this month asks about another new AI tool? Start by looking behind the question. Ask executives what they’re hoping to achieve by adopting an AI trend. Are they looking for speed? Security? Once you know the objective, you can steer the conversation to what your team’s already doing to reach it.
It’s also 100% fine to admit you don’t know a given tool yet. In fact, change your goal from understanding every new tool to keeping a general eye on trends, so you can choose a promising tool to try out later.
59% struggle to evaluate AI technologies consistently
More than half of technology leaders told us they struggle significantly or moderately to evaluate new AI technologies consistently across governance, security, and operational requirements.
This seems like a messy problem to untangle, but the good news is you already know how to do much of this work. Start with the framework you already use to evaluate technology service providers. If you’re giving your whole development team access to Claude or ChatGPT, it means you’re putting data on someone else’s server. Set up and evaluate a data management agreement just as you would with any other Cloud provider.
You’ll also want to consider how strong a given tool’s adoption might be in the marketplace and how likely it is that the tool will be around later. Finally, evaluate whether the company uses trusted business practices and aligns with your industry’s priorities.
When you test out a new AI tool , make it as low risk as possible by setting up a developer sandbox and running short pilot studies that take weeks, not months. Then try basing your experiment around this simple framework: apply the tool to a concrete problem, build a working baseline, and improve iteratively. Stay away from mission-critical systems and use mock or anonymized data to keep personally identifiable information (PII) and protected health information (PHI) out.
Pilots show how a tool behaves, what it costs, and where the trade-offs live. This information all informs later governance decisions. Once you decide to move into production, you’ll want to follow your organization’s usual controls for new technology tools and collaborate with key partners in product and compliance.
62% say engineering fatigue is growing
Nearly two-thirds of survey respondents said engineering fatigue is growing significantly or moderately as teams manage AI-generated code alongside legacy systems.
Reviewing AI-generated code happens along a continuum of approach and time investment:
Pairing with an LLM: It’s a powerful partner but checking every step can be slow.
Reviewing afterward: Going over every piece of AI code can bury developers in work.
Letting AI review AI: This works for catching errors but only up to a point
At the far end of this effort continuum is the largely hands-off dark factory approach. You don’t watch what happens inside the AI code factory. Instead, you define tests and controls that measure success and limit risk.
Take a financial app that recommends trades. Tests might require expected returns between 5% and 20% and block the AI from executing trades. The AI can write code overnight, but it’s never allowed to change the tests. It stops only when it can’t meet the criteria or reaches a judgment call.
Our team sometimes takes what we call a dim factory approach in our agentic coding. Most of the lights are off (we’re not hands-on with the AI coding process), but we can still walk through the factory and see what’s generally happening without reviewing every detail. It reduces the burden on developers while still maintaining additional oversight.
The right approach to AI-generated code depends on your culture, your risk tolerance, and whether you’re prioritizing feature speed or maintainability. Legacy systems also come with legacy cultures, and those need navigating, too.
What all these stats hint at but don’t outright state is that there’s a gigantic emotional component to rapid AI change, too. Many developers take pride in crafting code, so naturally, there can be significant grief in turning over this hard-won skill to AI.
One way to re-frame the situation: Focus on the new craft you’re developing as you direct an LLM. As with any emerging technology, you’ll make mistakes, but you’ll also work through them and learn. Focus on personal skill growth vs. letting imposter syndrome creep in.
Another big source of mixed emotions about AI: Decision fatigue. Constantly evaluating output from a dozen agents leaves no room to breathe, especially as AI pushes productivity expectations higher. This volume can leave people feeling burned out. In a recent workplace survey from Lenny’s Letter, 55.7% of tech workers reported significant burnout, up from 44.7% in last year’s survey. There’s no easy answer, but as you’re trying to navigate all the change, focus on the decisions that matter most, set a manageable pace, and accept that everyone has a cognitive load limit.
There’s no one-size-fits-all fix
How your engineering team approaches AI depends on business context and whether you’re defining the team’s overall relationship with AI or simply matching a tool to a problem. Plus, all this context and technology keep changing at a rapid pace.
Want help navigating it? Our AI and machine learning expertise can guide you through governance, focused pilots, and strategies for managing AI-generated code. Let’s talk.
If you enjoyed this post, you might also like:
How to Use ChatGPT to Find Custom Software Consultants
From idea to impact: The role of rapid prototyping in AgeTech
Diversity, Equity, Inclusion and Building Great Teams - One Recruiter’s Rambling Thoughts on Inclusive Hiring
Sign up to receive a weekly recap from thoughtbot
Check out our feeds
About thoughtbot
We've been helping engineering teams deliver exceptional products for over 20 years. Our designers, developers, and product managers work closely with teams to solve your toughest software challenges through collaborative design and development. Learn more about us .
The design of a robot and thoughtbot are registered trademarks of thoughtbot, inc.

## Original Extract

AI fatigue is setting in for engineering teams. Our research shows most feel pressure to assess every AI trend, struggle to evaluate tools, and wrestle with evaluating AI-generated code. Here’s some practical advice for navigating the pain points.

Skip to main content
thoughtbot
thoughtbot
Close Menu
We're live
Live on Twitch!
thoughtbot is livestreaming
Work alongside the thoughtbot team as we collaborate with each other and our clients, live. Ask us anything, we're live right now!
Work alongside the thoughtbot team as we collaborate with each other and our clients, live. Ask us anything, we're live right now!
Menu
Let’s get started!
Back
View all Services
Development
Ruby on Rails
Back
View all Resources
Development
Tech Leadership Resources
Giant Robots Smashing Into Other Giant Robots Podcast
What causes AI fatigue for engineers and how to ease the strain
Copy as Markdown
AI fatigue is hitting some development teams hard. In our original research study , engineering, product and technology leaders from US healthcare companies told us they’re feeling pressured to evaluate a firehose of new AI trends and tools while struggling to navigate AI-generated code alongside legacy systems.
This isn’t entirely new. The industry has gone through waves like this before, from the rise of the internet and the advent of social media to mobile development and every new round of web technology in between. But the rapid progression of LLMs speeds up this cycle, which can leave teams to feel like they’re losing control.
Here’s what our study uncovered about the causes of AI fatigue and some practical ideas for successfully navigating these pain points.
73% feel pressure to evaluate every AI trend
Nearly three-quarters of survey respondents told us they feel significant or moderate pressure to evaluate every new AI trend despite limited bandwidth.
So, what’s a developer or product manager to do when the tenth email this month asks about another new AI tool? Start by looking behind the question. Ask executives what they’re hoping to achieve by adopting an AI trend. Are they looking for speed? Security? Once you know the objective, you can steer the conversation to what your team’s already doing to reach it.
It’s also 100% fine to admit you don’t know a given tool yet. In fact, change your goal from understanding every new tool to keeping a general eye on trends, so you can choose a promising tool to try out later.
59% struggle to evaluate AI technologies consistently
More than half of technology leaders told us they struggle significantly or moderately to evaluate new AI technologies consistently across governance, security, and operational requirements.
This seems like a messy problem to untangle, but the good news is you already know how to do much of this work. Start with the framework you already use to evaluate technology service providers. If you’re giving your whole development team access to Claude or ChatGPT, it means you’re putting data on someone else’s server. Set up and evaluate a data management agreement just as you would with any other Cloud provider.
You’ll also want to consider how strong a given tool’s adoption might be in the marketplace and how likely it is that the tool will be around later. Finally, evaluate whether the company uses trusted business practices and aligns with your industry’s priorities.
When you test out a new AI tool , make it as low risk as possible by setting up a developer sandbox and running short pilot studies that take weeks, not months. Then try basing your experiment around this simple framework: apply the tool to a concrete problem, build a working baseline, and improve iteratively. Stay away from mission-critical systems and use mock or anonymized data to keep personally identifiable information (PII) and protected health information (PHI) out.
Pilots show how a tool behaves, what it costs, and where the trade-offs live. This information all informs later governance decisions. Once you decide to move into production, you’ll want to follow your organization’s usual controls for new technology tools and collaborate with key partners in product and compliance.
62% say engineering fatigue is growing
Nearly two-thirds of survey respondents said engineering fatigue is growing significantly or moderately as teams manage AI-generated code alongside legacy systems.
Reviewing AI-generated code happens along a continuum of approach and time investment:
Pairing with an LLM: It’s a powerful partner but checking every step can be slow.
Reviewing afterward: Going over every piece of AI code can bury developers in work.
Letting AI review AI: This works for catching errors but only up to a point
At the far end of this effort continuum is the largely hands-off dark factory approach. You don’t watch what happens inside the AI code factory. Instead, you define tests and controls that measure success and limit risk.
Take a financial app that recommends trades. Tests might require expected returns between 5% and 20% and block the AI from executing trades. The AI can write code overnight, but it’s never allowed to change the tests. It stops only when it can’t meet the criteria or reaches a judgment call.
Our team sometimes takes what we call a dim factory approach in our agentic coding. Most of the lights are off (we’re not hands-on with the AI coding process), but we can still walk through the factory and see what’s generally happening without reviewing every detail. It reduces the burden on developers while still maintaining additional oversight.
The right approach to AI-generated code depends on your culture, your risk tolerance, and whether you’re prioritizing feature speed or maintainability. Legacy systems also come with legacy cultures, and those need navigating, too.
What all these stats hint at but don’t outright state is that there’s a gigantic emotional component to rapid AI change, too. Many developers take pride in crafting code, so naturally, there can be significant grief in turning over this hard-won skill to AI.
One way to re-frame the situation: Focus on the new craft you’re developing as you direct an LLM. As with any emerging technology, you’ll make mistakes, but you’ll also work through them and learn. Focus on personal skill growth vs. letting imposter syndrome creep in.
Another big source of mixed emotions about AI: Decision fatigue. Constantly evaluating output from a dozen agents leaves no room to breathe, especially as AI pushes productivity expectations higher. This volume can leave people feeling burned out. In a recent workplace survey from Lenny’s Letter, 55.7% of tech workers reported significant burnout, up from 44.7% in last year’s survey. There’s no easy answer, but as you’re trying to navigate all the change, focus on the decisions that matter most, set a manageable pace, and accept that everyone has a cognitive load limit.
There’s no one-size-fits-all fix
How your engineering team approaches AI depends on business context and whether you’re defining the team’s overall relationship with AI or simply matching a tool to a problem. Plus, all this context and technology keep changing at a rapid pace.
Want help navigating it? Our AI and machine learning expertise can guide you through governance, focused pilots, and strategies for managing AI-generated code. Let’s talk.
If you enjoyed this post, you might also like:
How to Use ChatGPT to Find Custom Software Consultants
From idea to impact: The role of rapid prototyping in AgeTech
Diversity, Equity, Inclusion and Building Great Teams - One Recruiter’s Rambling Thoughts on Inclusive Hiring
Sign up to receive a weekly recap from thoughtbot
Check out our feeds
About thoughtbot
We've been helping engineering teams deliver exceptional products for over 20 years. Our designers, developers, and product managers work closely with teams to solve your toughest software challenges through collaborative design and development. Learn more about us .
The design of a robot and thoughtbot are registered trademarks of thoughtbot, inc.
