---
source: "https://airlock.lunareng.com/docs"
hn_url: "https://news.ycombinator.com/item?id=50419615"
title: "Show HN: Airlock – Policy-as-code for fleets of AI coding agents"
article_title: "Airlock"
image: ""
author: "lovetocode"
captured_at: "2026-11-09T01:28:04Z"
capture_tool: "hn-digest"
hn_id: 50419615
score: 1
comments: 0
posted_at: "2026-11-09T00:59:27Z"
tags:
  - hacker-news
---

# Show HN: Airlock – Policy-as-code for fleets of AI coding agents

- HN: [50419615](https://news.ycombinator.com/item?id=50419615)
- Source: [airlock.lunareng.com](https://airlock.lunareng.com/docs)
- Score: 1
- Comments: 0
- Posted: 2026-11-09T00:59:27Z

## Translation

Title: Show HN: Airlock – Policy-as-code for fleets of AI coding agents
Article title: Airlock
Description: Your agents work around the clock. Your team decides how far they go. And in the morning, you can see what got done, and what it cost.
HN text: My name is Matt and I am building Airlock. Airlock is going to be an open standard on managing autonomous fleets of engineering agents. I am seeking constructive feedback on the direction of this project and also asking what are your major pain points with managing AI use at larger businesses. Right now I am running my own company using this framework with a total of 23 workflows across four large scale projects. Airlock is still really early, only 3-4 weeks old. Airlock's configuration standard is aimed to be open source and I will also open source the backplane once I gain a bit more confidence in what I have built here and it is a bit more ready for public scrutiny. Today the schema itself and the reusable GitHub actions are publicly available. I will selfishly figure out how to monetize this at some point but my desired customer is going to be larger enterprises. Airlock kind of does a lot of things but if I had to sum it up in several bullet points it would be: 1) Allows for the configuration of agent behaviors and life cycles under version control. Each agent run proves who it is with GitHub’s OIDC token, and Airlock trades that for a GitHub token that lasts one hour and can only touch the repos and permissions listed in that agent’s role file. 2) Allows for configuring levels of autonomy starting with a repository's base autonomy level with stricter overrides down to individual files within a code base. (e.g. you trust agents for most of the repo but perhaps you want extra gating around the 'billing' or 'auth' modules). 3) Supports a standard autonomy scale. At level 4 agents are leading a project's direction while at level 1 they are dependent on humans to lead the direction of a project. 4) Provides an out-of-the-box OTEL collector that stores a subset of agent actions you would care about in the event of an audit. I’m exploring storing agent prompts and user prompts so that only the customer can decrypt them, and haven’t decided how. 5) Comes with a shadow mode that allows you to run Airlock and observe how it would behave in your org without turning it on. To try it you just need to install the GitHub App, create a policy repository and run on GitHub hosted runners. I personally use self-hosted runners on EC2 spot instances. Visit https://airlock.lunareng.com/docs/quickstart to get started. I have been really passionate lately about exploring how businesses of the [near] future will effectively manage autonomous development using agents at scale. I want to help answer the question of how can the business better manage AI spend, prioritization of token usage, auditability of agent actions and compliance with various governance, risk and compliance frameworks. I am coming from the perspective of being an engineer for 15+ years but I have somehow managed to also develop a good understanding of how larger businesses operate. My vision of the future looks like a world where engineers don't use agent harnesses directly from their computer [for coding]. Instead engineering teams are managing a larger backlog of work that is automatically picked up by these agents. Engineering will instead be less about writing code and more about effectively deploying technology to solve business problems. Thanks for taking a look and I look forward to meeting you in the comments. References:
https://github.com/lunareng/airlock-actions
https://airlock.lunareng.com/schema/role.schema.json
https://airlock.lunareng.com/schema/autonomy.schema.json
https://airlock.lunareng.com/schema/config.schema.json
https://www.anthropic.com/institute/measuring-pace-of-ai-dev...
https://felixrieseberg.com/how-to-be-more-ambitious-a-tutori...

Article text:
Airlock

## Original Extract

Your agents work around the clock. Your team decides how far they go. And in the morning, you can see what got done, and what it cost.

My name is Matt and I am building Airlock. Airlock is going to be an open standard on managing autonomous fleets of engineering agents. I am seeking constructive feedback on the direction of this project and also asking what are your major pain points with managing AI use at larger businesses. Right now I am running my own company using this framework with a total of 23 workflows across four large scale projects. Airlock is still really early, only 3-4 weeks old. Airlock's configuration standard is aimed to be open source and I will also open source the backplane once I gain a bit more confidence in what I have built here and it is a bit more ready for public scrutiny. Today the schema itself and the reusable GitHub actions are publicly available. I will selfishly figure out how to monetize this at some point but my desired customer is going to be larger enterprises. Airlock kind of does a lot of things but if I had to sum it up in several bullet points it would be: 1) Allows for the configuration of agent behaviors and life cycles under version control. Each agent run proves who it is with GitHub’s OIDC token, and Airlock trades that for a GitHub token that lasts one hour and can only touch the repos and permissions listed in that agent’s role file. 2) Allows for configuring levels of autonomy starting with a repository's base autonomy level with stricter overrides down to individual files within a code base. (e.g. you trust agents for most of the repo but perhaps you want extra gating around the 'billing' or 'auth' modules). 3) Supports a standard autonomy scale. At level 4 agents are leading a project's direction while at level 1 they are dependent on humans to lead the direction of a project. 4) Provides an out-of-the-box OTEL collector that stores a subset of agent actions you would care about in the event of an audit. I’m exploring storing agent prompts and user prompts so that only the customer can decrypt them, and haven’t decided how. 5) Comes with a shadow mode that allows you to run Airlock and observe how it would behave in your org without turning it on. To try it you just need to install the GitHub App, create a policy repository and run on GitHub hosted runners. I personally use self-hosted runners on EC2 spot instances. Visit https://airlock.lunareng.com/docs/quickstart to get started. I have been really passionate lately about exploring how businesses of the [near] future will effectively manage autonomous development using agents at scale. I want to help answer the question of how can the business better manage AI spend, prioritization of token usage, auditability of agent actions and compliance with various governance, risk and compliance frameworks. I am coming from the perspective of being an engineer for 15+ years but I have somehow managed to also develop a good understanding of how larger businesses operate. My vision of the future looks like a world where engineers don't use agent harnesses directly from their computer [for coding]. Instead engineering teams are managing a larger backlog of work that is automatically picked up by these agents. Engineering will instead be less about writing code and more about effectively deploying technology to solve business problems. Thanks for taking a look and I look forward to meeting you in the comments. References:
https://github.com/lunareng/airlock-actions
https://airlock.lunareng.com/schema/role.schema.json
https://airlock.lunareng.com/schema/autonomy.schema.json
https://airlock.lunareng.com/schema/config.schema.json
https://www.anthropic.com/institute/measuring-pace-of-ai-dev...
https://felixrieseberg.com/how-to-be-more-ambitious-a-tutori...

Airlock
