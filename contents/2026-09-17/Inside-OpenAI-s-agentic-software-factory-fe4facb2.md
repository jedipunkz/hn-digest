---
source: "https://newsletter.pragmaticengineer.com/p/openai-software-factory"
hn_url: "https://news.ycombinator.com/item?id=49745280"
title: "Inside OpenAI’s agentic software factory"
article_title: "Inside OpenAI’s agentic software factory - by Gergely Orosz"
image: "https://substackcdn.com/image/fetch/$s_!Eiw9!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F1e73721e-e4d8-473f-a5a1-b6211c14f6ca_2048x1762.png"
author: "gfortaine"
captured_at: "2026-09-17T20:02:37Z"
capture_tool: "hn-digest"
hn_id: 49745280
score: 1
comments: 0
posted_at: "2026-09-17T19:17:25Z"
tags:
  - hacker-news
---

# Inside OpenAI’s agentic software factory

- HN: [49745280](https://news.ycombinator.com/item?id=49745280)
- Source: [newsletter.pragmaticengineer.com](https://newsletter.pragmaticengineer.com/p/openai-software-factory)
- Score: 1
- Comments: 0
- Posted: 2026-09-17T19:17:25Z

## Translation

Title: Inside OpenAI’s agentic software factory
Article title: Inside OpenAI’s agentic software factory - by Gergely Orosz
Description: A deepdive into how Codex has “taken over” OpenAI, how the frontier lab builds its agentic software factory, and the engineering challenges of one billion users. Details from inside OpenAI

Article text:
Inside OpenAI’s agentic software factory - by Gergely Orosz
The Pragmatic Engineer
Subscribe Sign in Deepdives Inside OpenAI’s agentic software factory
A deepdive into how Codex has “taken over” OpenAI, how the frontier lab builds its agentic software factory, and the engineering challenges of one billion users. Details from inside OpenAI
Gergely Orosz Sep 15, 2026 ∙ Paid 260 7 17 Share It’s rare to work with an unlimited token budget, but at OpenAI, that’s what all engineers, researchers, finance colleagues, and marketing folks do. Recently, I visited one of the world’s leading frontier labs to find out how OpenAI operates today – and for a glimpse at where software engineering might be headed as a profession.
Plenty has changed since I visited OpenAI’s headquarters last year. Within a year, Codex has gone from a “nice-to-have” tool to being the backbone of pretty much everything at the company.
To learn more, I talked with seven engineering leaders and engineers there: Venkat Venkataramani (VP of Engineering, Applied Infra), Sulman Choudhry (Head of Engineering, ChatGPT), Andrew Ambrosino (Lead, Desktop), Joe Gershenson (Lead, Core Agent team), Akshay Nathan (Engineering Lead, Productivity), Ahmed Ibrahim (Engineer, Codex) and Steve Coffey (Engineer, Responses API). Thanks to all for taking part!
Codex takes over at OpenAI. In a matter of months, nearly all OpenAI’s non-engineers moved over to Codex and ChatGPT Work without a mandate from above for it.
Death of the IDE & pull requests. IDE usage has been down since January when Codex usage started to surge. PRs and code reviews need to be rethought.
OpenAI’s agentic software factory. OpenAI has built a “software factory” with several automated, agentic feedback loops: for example, Perf Factory monitors production and kicks off Codex agents to automatically fix performance issues.
How engineering tooling & practices are changing. Hand-built internal tools are slowly being replaced by Codex, which is increasingly preferred for debugging over specialized tools. Harness efficiency is critical in software factories.
Engineering for a billion users: how OpenAI scales up its infra. They buy first and take it in-house later. Also, geographic infra distribution, capacity planning tactics and challenges.
Making OpenAI’s API more reliable and performant . CPUs are becoming a bottleneck, doing slower deployments on purpose, and solving load challenges.
How the software engineering job is changing. Engineering specializations are disappearing, judgment and agency are more important, and it only takes one or two engineers for previously “impossible” rewrites and migrations to succeed.
Before we start, a scheduling update: I’m in New York for the week, attending the LDX3 conference and visiting a few startups and tech companies in the city, so there will be no edition of The Pulse on Thursday. Normal service resumes next week!
The bottom of this article could be cut off in some email clients. Read the full article uninterrupted, online.
The takeaway from my visit to the company’s headquarters which really sticks out is that Codex – and lately Codex and ChatGPT Work – have taken over everything there, starting in around January. Desktop lead, Andrew Ambrosino, told me:
“The big theme of the past months has been that everything is now a coding agent. Whether the visible code is your output or not, agents write your artifacts.
Think of it like this: your entire life is via software. You have these powerful tools (agents) in your computer, and the ability to loop and reason and write code is the ability to do everything.”
The token usage chart below shows this sudden adoption surge:
Codex usage since August 2025 at OpenAI by department. Source: OpenAI In a four-month period, non-engineering orgs like finance, recruitment, and legal went from ~0% usage of Codex to 90% usage. Now, almost all OpenAI employees use Codex and ChatGPT Work weekly. So, what happened?
OpenAI released the Codex app for Mac in February and for Windows in March, and ChatGPT Work (powered by the Codex harness) in July. Following that, non-engineers there moved all their workflows over to Codex and then to Work. Caveat: OpenAI’s internal version of Codex is a lot more advanced than its external counterpart because it’s plugged into pretty much every OpenAI system – similar to how Ramp’s Inspect AI agent has been wired up.
The fascinating part of this is that OpenAI got close to 40% adoption across non-engineering teams at a time when the Codex app was hostile to non-engineering users (hard to use). Between February and April, the Codex app still showed the code on-screen, but even so, non-technical colleagues outside of engineering still used it because it could do complex work like researching and creating a presentation, document, spreadsheet, or tasks that produce rich output. Today, those folks are very heavy users of it.
Being able to work for longer on more complex things drove adoption. OpenAI added the /goal setting to Codex, where you can set up a goal for the agent and it keeps working until it is complete. Between April and May, usage surged from 60% to 90%. Andrew believes improvement in the harness’s handling of long-running tasks was one cause of this:
“The number one thing that is changing is that people are starting to use threads for much longer, and this longer usage has been a breakthrough. It’s surprising to see the sheer length of time that people spend on a thread – even days! They often set a goal and then have the model crank.
Codex being good at longer-running tasks seems to cause people to do fewer things in parallel. This is because a long-running agent often spins off other agents to do other things, reducing the surface area that you, as a human, have to manage.”
“Awareness overhang” is another cause of the rapid adoption, the Codex team believes. As Akshay Nathan, Engineering Lead, Productivity team, told me:
“For a long time, we had a ‘capability overhang’: the models were capable but the products didn’t fully bring that out. Now, we’re seeing an awareness gap. Some people have figured out they can use Codex to monitor Slack, update Airtable, or create onboarding materials. But many others still use it for one task and then discover more uses from teammates via word-of-mouth.
But there’s still so much more, under the surface, that you can do with Codex.”
Role-specific and team-specific plugins are created and distributed. Another thing that sped up adoption is that each group started to distribute useful role-specific workflows as plugins. Andrew explained why it’s important to not just offer a generic coding agent:
“If you build a product that can do anything, teams need a way to make it their own. You can’t just give everyone an empty box. Skills and plugins let teams adapt the agent to their work. Sometimes, we also need a new app capability, like a browser that the agent can use alongside those skills. But the same building blocks already cover a lot of different roles.”
Subject matter experts are embedded in ChatGPT Work engineering teams. The models have become “smarter” than developers in some domains, so devs cannot channel “taste” into the harness in those areas. So, people who are domain experts are onboarded onto engineering teams. This is one outcome of ChatGPT Work being used by so many non-engineering domains: experts embedded with engineering advise developers on things like what a good slide deck, spreadsheet, or business report looks like.
Of course, domain experts being in engineering teams is a decades-old best practice for building quality products. It seems like this gets rediscovered in different contexts every few years!
OpenAI is fully dependent on Codex and Work. This is so much the case that in the event of even a minor outage, internal messages from colleagues alert the Codex and Work teams at the same time as – or before – automated alerts.
Basically, work happens through Codex and Work, and pretty much nothing else. From the outside, this dependence on a single shared harness is particularly eye-catching; two years ago, there were no AI agents, only advanced AI autocomplete!
2. Death of the IDE & pull requests
Late last year, the Codex team was torn about whether to release the Codex desktop app. Andrew recalls the hesitation:
“In December 2025, we weren’t entirely sure if we would release the Codex app. We had the Codex CLI as a terminal, and there are large, feature-rich IDEs out there. So, would there be space for a dev tool that is between a terminal and an IDE? In my head, there was this future where it would not work out, and be the kind of ‘misfit’ like the iPad was.
A lot of people buy iPads and then never use them: they either use their smaller, more portable smartphone (which could be the equivalent of the CLI in this metaphor) , or their feature-rich laptop (the equivalent of the IDE).
Also, don’t forget that in November, Antigravity came out as a VS Code fork. This added to the feeling that perhaps we should have also forked VS Code for the Codex app. But still, we dismissed the temptation and went with our gut feeling that as AI agents get better, IDEs will matter less.”
Indeed, since January, IDE usage has gone down and OpenAI’s bet looks like a good one. However, the Codex app is becoming a little more akin to an IDE: for example, the ability to edit files inside the app was shipped in June.
CI/CD systems are seeing massive load increases. One sign of productivity gains from Codex is the amount of additional code flowing through OpenAI’s dev infra systems. More code being created and pushed leads to new scaling challenges which the team is currently heads-down on solving. Venkat Venkataramani, VP of Engineering, Applied Infra, said:
“The number of pull requests (PRs) per engineer is growing like a hockey stick (at a very high, accelerating rate) . Every part of the build-test-deploy pipeline is seeing dramatically more load.
We’re talking about roughly a 10x increase in load on some systems. At most companies, that kind of growth might happen over two or three years. At OpenAI, we see it in about six months.
That level of acceleration exposes bottlenecks everywhere: version control has to handle far more code being written and pushed, CI/CD systems have to scale with it, and production release processes have to absorb a much higher rate of change.
Every month, we wake up to a new set of infrastructure scaling challenges to solve. Just when we think we’ve created enough capacity for the next phase of growth, the model unlocks another wave of capabilities, which creates a new set of bottlenecks somewhere else in the system.”
In this context, PRs and code reviews are being rethought. They have “core primitives” in software engineering, but this level of development acceleration is an opportunity to reimagine them. Again, from Venkat:
“The question we ought to ask ourselves in the middle of all this development acceleration is how do we reimagine many things we took for granted. For example, how do we reimagine the CI (continuous integration) and CD (continuous deployment) process? What does observability mean in this world, and how should people interact with pull requests?
If you ask me, the way we do code review today makes less and less sense, and the same is true for pull requests.
We’re now seeing agentic code reviews that look at code changes through a series of different lenses. In the past, it would have been impractical for a cloud infrastructure engineer and a security engineer to review every single code change. With agents, that becomes possible.
We can rethink how code is deployed with agents, too. We are building an agent that “handholds” a change all the way to production — whether it’s a code change or a change behind a feature flag. It observes the relevant monitoring graphs, but can also build its own dashboard to monitor important signals. More of our code changes are going to

[truncated]

## Original Extract

A deepdive into how Codex has “taken over” OpenAI, how the frontier lab builds its agentic software factory, and the engineering challenges of one billion users. Details from inside OpenAI

Inside OpenAI’s agentic software factory - by Gergely Orosz
The Pragmatic Engineer
Subscribe Sign in Deepdives Inside OpenAI’s agentic software factory
A deepdive into how Codex has “taken over” OpenAI, how the frontier lab builds its agentic software factory, and the engineering challenges of one billion users. Details from inside OpenAI
Gergely Orosz Sep 15, 2026 ∙ Paid 260 7 17 Share It’s rare to work with an unlimited token budget, but at OpenAI, that’s what all engineers, researchers, finance colleagues, and marketing folks do. Recently, I visited one of the world’s leading frontier labs to find out how OpenAI operates today – and for a glimpse at where software engineering might be headed as a profession.
Plenty has changed since I visited OpenAI’s headquarters last year. Within a year, Codex has gone from a “nice-to-have” tool to being the backbone of pretty much everything at the company.
To learn more, I talked with seven engineering leaders and engineers there: Venkat Venkataramani (VP of Engineering, Applied Infra), Sulman Choudhry (Head of Engineering, ChatGPT), Andrew Ambrosino (Lead, Desktop), Joe Gershenson (Lead, Core Agent team), Akshay Nathan (Engineering Lead, Productivity), Ahmed Ibrahim (Engineer, Codex) and Steve Coffey (Engineer, Responses API). Thanks to all for taking part!
Codex takes over at OpenAI. In a matter of months, nearly all OpenAI’s non-engineers moved over to Codex and ChatGPT Work without a mandate from above for it.
Death of the IDE & pull requests. IDE usage has been down since January when Codex usage started to surge. PRs and code reviews need to be rethought.
OpenAI’s agentic software factory. OpenAI has built a “software factory” with several automated, agentic feedback loops: for example, Perf Factory monitors production and kicks off Codex agents to automatically fix performance issues.
How engineering tooling & practices are changing. Hand-built internal tools are slowly being replaced by Codex, which is increasingly preferred for debugging over specialized tools. Harness efficiency is critical in software factories.
Engineering for a billion users: how OpenAI scales up its infra. They buy first and take it in-house later. Also, geographic infra distribution, capacity planning tactics and challenges.
Making OpenAI’s API more reliable and performant . CPUs are becoming a bottleneck, doing slower deployments on purpose, and solving load challenges.
How the software engineering job is changing. Engineering specializations are disappearing, judgment and agency are more important, and it only takes one or two engineers for previously “impossible” rewrites and migrations to succeed.
Before we start, a scheduling update: I’m in New York for the week, attending the LDX3 conference and visiting a few startups and tech companies in the city, so there will be no edition of The Pulse on Thursday. Normal service resumes next week!
The bottom of this article could be cut off in some email clients. Read the full article uninterrupted, online.
The takeaway from my visit to the company’s headquarters which really sticks out is that Codex – and lately Codex and ChatGPT Work – have taken over everything there, starting in around January. Desktop lead, Andrew Ambrosino, told me:
“The big theme of the past months has been that everything is now a coding agent. Whether the visible code is your output or not, agents write your artifacts.
Think of it like this: your entire life is via software. You have these powerful tools (agents) in your computer, and the ability to loop and reason and write code is the ability to do everything.”
The token usage chart below shows this sudden adoption surge:
Codex usage since August 2025 at OpenAI by department. Source: OpenAI In a four-month period, non-engineering orgs like finance, recruitment, and legal went from ~0% usage of Codex to 90% usage. Now, almost all OpenAI employees use Codex and ChatGPT Work weekly. So, what happened?
OpenAI released the Codex app for Mac in February and for Windows in March, and ChatGPT Work (powered by the Codex harness) in July. Following that, non-engineers there moved all their workflows over to Codex and then to Work. Caveat: OpenAI’s internal version of Codex is a lot more advanced than its external counterpart because it’s plugged into pretty much every OpenAI system – similar to how Ramp’s Inspect AI agent has been wired up.
The fascinating part of this is that OpenAI got close to 40% adoption across non-engineering teams at a time when the Codex app was hostile to non-engineering users (hard to use). Between February and April, the Codex app still showed the code on-screen, but even so, non-technical colleagues outside of engineering still used it because it could do complex work like researching and creating a presentation, document, spreadsheet, or tasks that produce rich output. Today, those folks are very heavy users of it.
Being able to work for longer on more complex things drove adoption. OpenAI added the /goal setting to Codex, where you can set up a goal for the agent and it keeps working until it is complete. Between April and May, usage surged from 60% to 90%. Andrew believes improvement in the harness’s handling of long-running tasks was one cause of this:
“The number one thing that is changing is that people are starting to use threads for much longer, and this longer usage has been a breakthrough. It’s surprising to see the sheer length of time that people spend on a thread – even days! They often set a goal and then have the model crank.
Codex being good at longer-running tasks seems to cause people to do fewer things in parallel. This is because a long-running agent often spins off other agents to do other things, reducing the surface area that you, as a human, have to manage.”
“Awareness overhang” is another cause of the rapid adoption, the Codex team believes. As Akshay Nathan, Engineering Lead, Productivity team, told me:
“For a long time, we had a ‘capability overhang’: the models were capable but the products didn’t fully bring that out. Now, we’re seeing an awareness gap. Some people have figured out they can use Codex to monitor Slack, update Airtable, or create onboarding materials. But many others still use it for one task and then discover more uses from teammates via word-of-mouth.
But there’s still so much more, under the surface, that you can do with Codex.”
Role-specific and team-specific plugins are created and distributed. Another thing that sped up adoption is that each group started to distribute useful role-specific workflows as plugins. Andrew explained why it’s important to not just offer a generic coding agent:
“If you build a product that can do anything, teams need a way to make it their own. You can’t just give everyone an empty box. Skills and plugins let teams adapt the agent to their work. Sometimes, we also need a new app capability, like a browser that the agent can use alongside those skills. But the same building blocks already cover a lot of different roles.”
Subject matter experts are embedded in ChatGPT Work engineering teams. The models have become “smarter” than developers in some domains, so devs cannot channel “taste” into the harness in those areas. So, people who are domain experts are onboarded onto engineering teams. This is one outcome of ChatGPT Work being used by so many non-engineering domains: experts embedded with engineering advise developers on things like what a good slide deck, spreadsheet, or business report looks like.
Of course, domain experts being in engineering teams is a decades-old best practice for building quality products. It seems like this gets rediscovered in different contexts every few years!
OpenAI is fully dependent on Codex and Work. This is so much the case that in the event of even a minor outage, internal messages from colleagues alert the Codex and Work teams at the same time as – or before – automated alerts.
Basically, work happens through Codex and Work, and pretty much nothing else. From the outside, this dependence on a single shared harness is particularly eye-catching; two years ago, there were no AI agents, only advanced AI autocomplete!
2. Death of the IDE & pull requests
Late last year, the Codex team was torn about whether to release the Codex desktop app. Andrew recalls the hesitation:
“In December 2025, we weren’t entirely sure if we would release the Codex app. We had the Codex CLI as a terminal, and there are large, feature-rich IDEs out there. So, would there be space for a dev tool that is between a terminal and an IDE? In my head, there was this future where it would not work out, and be the kind of ‘misfit’ like the iPad was.
A lot of people buy iPads and then never use them: they either use their smaller, more portable smartphone (which could be the equivalent of the CLI in this metaphor) , or their feature-rich laptop (the equivalent of the IDE).
Also, don’t forget that in November, Antigravity came out as a VS Code fork. This added to the feeling that perhaps we should have also forked VS Code for the Codex app. But still, we dismissed the temptation and went with our gut feeling that as AI agents get better, IDEs will matter less.”
Indeed, since January, IDE usage has gone down and OpenAI’s bet looks like a good one. However, the Codex app is becoming a little more akin to an IDE: for example, the ability to edit files inside the app was shipped in June.
CI/CD systems are seeing massive load increases. One sign of productivity gains from Codex is the amount of additional code flowing through OpenAI’s dev infra systems. More code being created and pushed leads to new scaling challenges which the team is currently heads-down on solving. Venkat Venkataramani, VP of Engineering, Applied Infra, said:
“The number of pull requests (PRs) per engineer is growing like a hockey stick (at a very high, accelerating rate) . Every part of the build-test-deploy pipeline is seeing dramatically more load.
We’re talking about roughly a 10x increase in load on some systems. At most companies, that kind of growth might happen over two or three years. At OpenAI, we see it in about six months.
That level of acceleration exposes bottlenecks everywhere: version control has to handle far more code being written and pushed, CI/CD systems have to scale with it, and production release processes have to absorb a much higher rate of change.
Every month, we wake up to a new set of infrastructure scaling challenges to solve. Just when we think we’ve created enough capacity for the next phase of growth, the model unlocks another wave of capabilities, which creates a new set of bottlenecks somewhere else in the system.”
In this context, PRs and code reviews are being rethought. They have “core primitives” in software engineering, but this level of development acceleration is an opportunity to reimagine them. Again, from Venkat:
“The question we ought to ask ourselves in the middle of all this development acceleration is how do we reimagine many things we took for granted. For example, how do we reimagine the CI (continuous integration) and CD (continuous deployment) process? What does observability mean in this world, and how should people interact with pull requests?
If you ask me, the way we do code review today makes less and less sense, and the same is true for pull requests.
We’re now seeing agentic code reviews that look at code changes through a series of different lenses. In the past, it would have been impractical for a cloud infrastructure engineer and a security engineer to review every single code change. With agents, that becomes possible.
We can rethink how code is deployed with agents, too. We are building an agent that “handholds” a change all the way to production — whether it’s a code change or a change behind a feature flag. It observes the relevant monitoring graphs, but can also build its own dashboard to monitor important signals. More of our code changes are going to

[truncated]
