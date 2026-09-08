---
source: "https://metalbear.com/blog/ai-native-sdlc-infrastructure/"
hn_url: "https://news.ycombinator.com/item?id=49617552"
title: "The AI-Native SDLC Starts with Your Infrastructure"
article_title: "The AI-Native SDLC Starts With Your Infrastructure | mirrord by MetalBear"
image: "https://metalbear.com/blog/ai-native-sdlc-infrastructure/thumbnail.png"
author: "ioanarebeca"
captured_at: "2026-09-08T22:24:42Z"
capture_tool: "hn-digest"
hn_id: 49617552
score: 1
comments: 0
posted_at: "2026-09-08T21:45:55Z"
tags:
  - hacker-news
---

# The AI-Native SDLC Starts with Your Infrastructure

- HN: [49617552](https://news.ycombinator.com/item?id=49617552)
- Source: [metalbear.com](https://metalbear.com/blog/ai-native-sdlc-infrastructure/)
- Score: 1
- Comments: 0
- Posted: 2026-09-08T21:45:55Z

## Translation

Title: The AI-Native SDLC Starts with Your Infrastructure
Article title: The AI-Native SDLC Starts With Your Infrastructure | mirrord by MetalBear
Description: Anthropic

Article text:
mirrord Overview
Autonomous Agents
For CI
Preview Environments Pricing
Docs
Customers
Blog
Contact
5k stars Sign in
Try Free →
mirrord Overview Kubernetes Dev Platform for AI-Powered Teams mirrord for AI Agents Turn Agents into Autonomous Developers Real Environment CI Run CI Tests Against a Real k8s Environment Preview Environments Live PR Previews Without Deploying Pricing
Docs
Customers
Blog
Contact
GitHub Sign in Try mirrord for free
Book a Demo
Back to blog The AI-Native SDLC Starts With Your Infrastructure
Stage 4: where the agent checks its own work
Why verification is the hard part
The agent never sees the running system
Running several agents against one cluster
Anthropic published a playbook for restructuring the software lifecycle around coding agents. Its premise is that the traditional SDLC was designed when writing code was the slow part, agents made that part fast, and the constraint moved to the stages around it.
The framework has six stages, and each one commits an artifact the next stage can read. Planning produces an intent.md . Design turns that into a spec.md . Build produces a plan.md before any code is edited. Deploy puts the review policy in a REVIEW.md .
It is more specific than most process documents, but it also leaves out an important detail that can decide whether the rest of the process actually works, which is what the agent’s code runs against when it checks itself.
Stage 4: where the agent checks its own work #
Stage 4 is the feedback loop, where the agent checks its own work before an engineer sees it. The playbook asks you to give it something to check against, whether tests, a build, or a screenshot diff. It tells you to stop the agent from turning a red test green by editing the test, using a hook that blocks edits to test files during a fix. For UI work it suggests wiring in a browser or screenshot tool over MCP.
Then it goes further than most organizations have, and asks you to treat the coding agent’s own configuration as software: evals running in CI that re-test CLAUDE.md , the skills, and the hooks whenever any of them change, with every production incident turned into a permanent eval.
What Stage 4 asks you to have in place before any of that works is a test suite and a build that run locally with one command each.
That prerequisite is where the playbook stops and your infrastructure starts. It tells you the agent needs tests it can run. It does not say what those tests should run against, and for a service that talks to a dozen others (plus databases, queues, third-party APIs, etc.), which describes most real-world software, that is most of the question.
Why verification is the hard part #
If the tests run against fake copies of those dozen services on the agent’s machine, then tests passing tells you the code works against the fakes. Whether it works against the ones in the cluster is a different question.
A developer running those same tests knows roughly how far to trust them. They know the fake billing service was written a year ago and that the real one changed its auth header two months ago. They know the fake search endpoint always returns the same three results, while the real one paginates. And they know that nothing in the fake set has ever rate-limited them or timed out. The agent knows none of that. It sees the tests pass and reports the work finished.
This is not a flaw in the playbook so much as the edge of what it can cover. Every other stage works on files in a repository, and Anthropic can be specific about those because Claude Code is theirs. The services, the databases, the queues, the message brokers, and everything else the code talks to are yours. No model vendor can tell you what those look like, so the playbook tells you to have a check and stops there.
The agent never sees the running system #
Look at what the artifact chain holds: intent.md , spec.md , plan.md , CLAUDE.md , the skills, REVIEW.md . Each one records something a person decided and wrote down.
None of it lets the agent look at the system as it runs right now in production (or staging). Not what the upstream service returns when you call it, not what is sitting on the queue, not what the staging database’s schema actually is today, which may be several migrations behind the branch the agent is working on. CLAUDE.md tells the agent what the organization decided, not what is actually running.
mirrord lets the agent’s code run against the real services in your staging cluster instead of fakes on its machine.
The code still runs locally, or on a CI runner or sandbox. What changes is everything around it: the process reads the same environment variables and secrets as the service it is standing in for in the cluster, its outbound calls go out through the cluster’s network, and traffic inside the cluster can be routed to it.
We built this for developers, who mostly use it to shorten their feedback loop by cutting out the deploy-and-wait cycle. Agents get more out of it, because a developer can still judge how stale a set of fakes has become but an agent cannot. Running the check against the services in the cluster means nobody has to make that judgment.
It helps before the check, too. The same connection lets the agent see what the API actually returns and what the messages on the queue actually contain, rather than working from documentation.
Running several agents against one cluster #
The playbook suggests running several Claude Code sessions at the same time, each in its own git worktree, with subagents inside a session.
Worktrees keep the code separate. They do nothing about the cluster those sessions check against. Point five agents at the same staging service and they interfere with each other and with the engineers already using it, which is usually where an organization decides that agents and shared staging do not mix and goes back to giving each agent its own (slow, expensive, shallow) copy.
The mirrord operator solves this. Traffic is filtered by header so each agent’s session only receives its own requests, queues are split so each session gets a private slice of a shared topic, and databases are branched so a session that writes does not disturb anyone else. One staging cluster serves as many agent sessions as you need to run against it.
If you are adopting the playbook, the question worth answering first is the one it doesn’t ask: what will the agent’s code be running against when it checks itself?
mirrord is open source. The operator, which is what handles the concurrent sessions above, is part of our commercial product. If you want to see what Stage 4 looks like running against your cluster, start at metalbear.com/mirrord/docs .
mirrord is a Kubernetes development platform that lets developers and AI coding agents test code in a production-like environment before deploying it. Your service runs wherever you're working, locally, in CI, or in an agent's sandbox, while mirrord proxies its traffic, environment variables, and files to and from a shared staging cluster, so it behaves as if it were deployed without actually being deployed.
Engineering teams at companies like monday.com, National Australia Bank, and SurveyMonkey use mirrord to iterate and ship faster, while spending less on dev environment infrastructure.
On this page Stage 4: where the agent checks its own work
Why verification is the hard part
The agent never sees the running system
Running several agents against one cluster
CTO and Co-founder of MetalBear. Eyal leads engineering at MetalBear, with a career spent building and managing engineering teams across security, fintech, and developer tools. He’s focused on making cloud native development as fast and intuitive as local development.
You may also like How To Reduce Token Costs of AI Coding Agents Aug 12, 2026 · 10 min read Four Layers of Validation in Kubernetes with Claude Code May 21, 2026 · 11 min read How To Prevent Token Burn Using mirrord With E2E Tests Apr 1, 2026 · 6 min read Try mirrord Want to dig deeper?
With mirrord, cloud developers can run local code in the context of their Kubernetes cluster — streamlining coding, debugging, testing, and troubleshooting.
Get hands-on experience with a ready-to-use development environment.
Developer infrastructure for AI‑powered teams.

## Original Extract

Anthropic

mirrord Overview
Autonomous Agents
For CI
Preview Environments Pricing
Docs
Customers
Blog
Contact
5k stars Sign in
Try Free →
mirrord Overview Kubernetes Dev Platform for AI-Powered Teams mirrord for AI Agents Turn Agents into Autonomous Developers Real Environment CI Run CI Tests Against a Real k8s Environment Preview Environments Live PR Previews Without Deploying Pricing
Docs
Customers
Blog
Contact
GitHub Sign in Try mirrord for free
Book a Demo
Back to blog The AI-Native SDLC Starts With Your Infrastructure
Stage 4: where the agent checks its own work
Why verification is the hard part
The agent never sees the running system
Running several agents against one cluster
Anthropic published a playbook for restructuring the software lifecycle around coding agents. Its premise is that the traditional SDLC was designed when writing code was the slow part, agents made that part fast, and the constraint moved to the stages around it.
The framework has six stages, and each one commits an artifact the next stage can read. Planning produces an intent.md . Design turns that into a spec.md . Build produces a plan.md before any code is edited. Deploy puts the review policy in a REVIEW.md .
It is more specific than most process documents, but it also leaves out an important detail that can decide whether the rest of the process actually works, which is what the agent’s code runs against when it checks itself.
Stage 4: where the agent checks its own work #
Stage 4 is the feedback loop, where the agent checks its own work before an engineer sees it. The playbook asks you to give it something to check against, whether tests, a build, or a screenshot diff. It tells you to stop the agent from turning a red test green by editing the test, using a hook that blocks edits to test files during a fix. For UI work it suggests wiring in a browser or screenshot tool over MCP.
Then it goes further than most organizations have, and asks you to treat the coding agent’s own configuration as software: evals running in CI that re-test CLAUDE.md , the skills, and the hooks whenever any of them change, with every production incident turned into a permanent eval.
What Stage 4 asks you to have in place before any of that works is a test suite and a build that run locally with one command each.
That prerequisite is where the playbook stops and your infrastructure starts. It tells you the agent needs tests it can run. It does not say what those tests should run against, and for a service that talks to a dozen others (plus databases, queues, third-party APIs, etc.), which describes most real-world software, that is most of the question.
Why verification is the hard part #
If the tests run against fake copies of those dozen services on the agent’s machine, then tests passing tells you the code works against the fakes. Whether it works against the ones in the cluster is a different question.
A developer running those same tests knows roughly how far to trust them. They know the fake billing service was written a year ago and that the real one changed its auth header two months ago. They know the fake search endpoint always returns the same three results, while the real one paginates. And they know that nothing in the fake set has ever rate-limited them or timed out. The agent knows none of that. It sees the tests pass and reports the work finished.
This is not a flaw in the playbook so much as the edge of what it can cover. Every other stage works on files in a repository, and Anthropic can be specific about those because Claude Code is theirs. The services, the databases, the queues, the message brokers, and everything else the code talks to are yours. No model vendor can tell you what those look like, so the playbook tells you to have a check and stops there.
The agent never sees the running system #
Look at what the artifact chain holds: intent.md , spec.md , plan.md , CLAUDE.md , the skills, REVIEW.md . Each one records something a person decided and wrote down.
None of it lets the agent look at the system as it runs right now in production (or staging). Not what the upstream service returns when you call it, not what is sitting on the queue, not what the staging database’s schema actually is today, which may be several migrations behind the branch the agent is working on. CLAUDE.md tells the agent what the organization decided, not what is actually running.
mirrord lets the agent’s code run against the real services in your staging cluster instead of fakes on its machine.
The code still runs locally, or on a CI runner or sandbox. What changes is everything around it: the process reads the same environment variables and secrets as the service it is standing in for in the cluster, its outbound calls go out through the cluster’s network, and traffic inside the cluster can be routed to it.
We built this for developers, who mostly use it to shorten their feedback loop by cutting out the deploy-and-wait cycle. Agents get more out of it, because a developer can still judge how stale a set of fakes has become but an agent cannot. Running the check against the services in the cluster means nobody has to make that judgment.
It helps before the check, too. The same connection lets the agent see what the API actually returns and what the messages on the queue actually contain, rather than working from documentation.
Running several agents against one cluster #
The playbook suggests running several Claude Code sessions at the same time, each in its own git worktree, with subagents inside a session.
Worktrees keep the code separate. They do nothing about the cluster those sessions check against. Point five agents at the same staging service and they interfere with each other and with the engineers already using it, which is usually where an organization decides that agents and shared staging do not mix and goes back to giving each agent its own (slow, expensive, shallow) copy.
The mirrord operator solves this. Traffic is filtered by header so each agent’s session only receives its own requests, queues are split so each session gets a private slice of a shared topic, and databases are branched so a session that writes does not disturb anyone else. One staging cluster serves as many agent sessions as you need to run against it.
If you are adopting the playbook, the question worth answering first is the one it doesn’t ask: what will the agent’s code be running against when it checks itself?
mirrord is open source. The operator, which is what handles the concurrent sessions above, is part of our commercial product. If you want to see what Stage 4 looks like running against your cluster, start at metalbear.com/mirrord/docs .
mirrord is a Kubernetes development platform that lets developers and AI coding agents test code in a production-like environment before deploying it. Your service runs wherever you're working, locally, in CI, or in an agent's sandbox, while mirrord proxies its traffic, environment variables, and files to and from a shared staging cluster, so it behaves as if it were deployed without actually being deployed.
Engineering teams at companies like monday.com, National Australia Bank, and SurveyMonkey use mirrord to iterate and ship faster, while spending less on dev environment infrastructure.
On this page Stage 4: where the agent checks its own work
Why verification is the hard part
The agent never sees the running system
Running several agents against one cluster
CTO and Co-founder of MetalBear. Eyal leads engineering at MetalBear, with a career spent building and managing engineering teams across security, fintech, and developer tools. He’s focused on making cloud native development as fast and intuitive as local development.
You may also like How To Reduce Token Costs of AI Coding Agents Aug 12, 2026 · 10 min read Four Layers of Validation in Kubernetes with Claude Code May 21, 2026 · 11 min read How To Prevent Token Burn Using mirrord With E2E Tests Apr 1, 2026 · 6 min read Try mirrord Want to dig deeper?
With mirrord, cloud developers can run local code in the context of their Kubernetes cluster — streamlining coding, debugging, testing, and troubleshooting.
Get hands-on experience with a ready-to-use development environment.
Developer infrastructure for AI‑powered teams.
