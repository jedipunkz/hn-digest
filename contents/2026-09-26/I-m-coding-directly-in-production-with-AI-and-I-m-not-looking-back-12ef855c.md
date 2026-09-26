---
source: "https://joydemo.com/blog/coding-directly-in-production-with-ai"
hn_url: "https://news.ycombinator.com/item?id=49856750"
title: "I'm coding directly in production with AI and I'm not looking back"
article_title: "I'm coding directly in production with AI and i'm not looking back | JoyDemo"
image: "https://joydemo.com/blog/coding-directly-in-production-with-ai/og.png?v=3"
author: "sh_tomer"
captured_at: "2026-09-26T14:53:57Z"
capture_tool: "hn-digest"
hn_id: 49856750
score: 3
comments: 7
posted_at: "2026-09-26T14:11:22Z"
tags:
  - hacker-news
---

# I'm coding directly in production with AI and I'm not looking back

- HN: [49856750](https://news.ycombinator.com/item?id=49856750)
- Source: [joydemo.com](https://joydemo.com/blog/coding-directly-in-production-with-ai)
- Score: 3
- Comments: 7
- Posted: 2026-09-26T14:11:22Z

## Translation

Title: I'm coding directly in production with AI and I'm not looking back
Article title: I'm coding directly in production with AI and i'm not looking back | JoyDemo
Description: How the JoyDemo team develops directly on production with AI, Codex CLI, Git worktrees, tests, and previews, cutting bugs by about 95%.

Article text:
JoyDemo
Home
Free Tools
Blog
Pricing
Start for free
← All posts
Inside JoyDemo
I'm coding directly in production with AI and i'm not looking back
How the JoyDemo team develops directly on production with AI, Codex CLI, Git worktrees, tests, and previews, cutting bugs by about 95%.
Moving development onto production has cut our bugs by about 95%. We used to find many of our bugs while moving changes from development to production. Now we build on the production host, using the runtime and setup the live product uses.
JoyDemo helps teams demonstrate software, train users, and show customers how to get more from a product. Reliability matters. Our team uses AI to make product changes directly on the production host, with tests and previews before changes reach the live app.
Fewer environment handoffs, fewer bugs
We used to build in one environment and deploy to another. A change could pass in development, then fail in production because the dependencies, settings, or runtime were different. That handoff was a steady source of bugs.
Now the AI works in the same runtime environment as the live product. We still keep proposed code separate from the live app until it's ready. Extensive automated tests run before every change is applied. If a change could have a bigger impact, we start a preview on the same host and review it before promoting it to the live app.
I don't mean that AI never gets things wrong. I attribute the improvement to removing environment differences and checking each change before it ships.
But how does this scale to a team?
We don't point every developer's AI at one shared folder. Each person gets a separate Git worktree, branch, and Codex CLI session on the production host. A worktree is simply that developer's own copy of the project. Their AI can edit and test it without colliding with another person's in-progress work or the files the live app is using.
Each worktree can run its own preview process. It uses the production host and the same runtime setup, with access to the services it needs. This gives each developer a realistic place to check a change while keeping the live app's code untouched.
When a change is ready, the developer runs the test suite and reviews the preview. We then promote that reviewed change to the live app. The important separation is between each person's workspace and the shared live code, not between a development machine and a production machine. We can work in parallel without changing the files currently serving customers.
If a bug gets through, the path to a fix is short. A developer can reproduce it on the production host, make a fix in their own worktree, run tests, and check the preview. Once reviewed, we promote that exact change to the live app. We don't need a long cycle to recreate the production setup somewhere else and transfer the fix across environments.
That speed matters for our product. When a demo flow doesn't work, a team may not be able to show a buyer how the software works or help a customer learn a feature. We can investigate the issue and release a fix within minutes.
Pieter Levels' post about coding on the production server without a separate development environment got me thinking about how much trouble the handoff between environments can create. Amjad Masad's post got me to weigh whether developing in production is too risky, what it makes possible, and which safeguards make the benefits outweigh the risk.
We took that inspiration and built a workflow for a team: separate worktrees and AI sessions, shared production context, extensive tests, and previews before higher-risk changes go live. AI does more of the implementation, but our team still decides what to build, reviews every change, and owns what reaches customers. That's why I'm not looking back.
Record interactive online product demos in seconds. Edit, publish, and embed them for free.
© 2026 JoyDemo. All rights reserved.

## Original Extract

How the JoyDemo team develops directly on production with AI, Codex CLI, Git worktrees, tests, and previews, cutting bugs by about 95%.

JoyDemo
Home
Free Tools
Blog
Pricing
Start for free
← All posts
Inside JoyDemo
I'm coding directly in production with AI and i'm not looking back
How the JoyDemo team develops directly on production with AI, Codex CLI, Git worktrees, tests, and previews, cutting bugs by about 95%.
Moving development onto production has cut our bugs by about 95%. We used to find many of our bugs while moving changes from development to production. Now we build on the production host, using the runtime and setup the live product uses.
JoyDemo helps teams demonstrate software, train users, and show customers how to get more from a product. Reliability matters. Our team uses AI to make product changes directly on the production host, with tests and previews before changes reach the live app.
Fewer environment handoffs, fewer bugs
We used to build in one environment and deploy to another. A change could pass in development, then fail in production because the dependencies, settings, or runtime were different. That handoff was a steady source of bugs.
Now the AI works in the same runtime environment as the live product. We still keep proposed code separate from the live app until it's ready. Extensive automated tests run before every change is applied. If a change could have a bigger impact, we start a preview on the same host and review it before promoting it to the live app.
I don't mean that AI never gets things wrong. I attribute the improvement to removing environment differences and checking each change before it ships.
But how does this scale to a team?
We don't point every developer's AI at one shared folder. Each person gets a separate Git worktree, branch, and Codex CLI session on the production host. A worktree is simply that developer's own copy of the project. Their AI can edit and test it without colliding with another person's in-progress work or the files the live app is using.
Each worktree can run its own preview process. It uses the production host and the same runtime setup, with access to the services it needs. This gives each developer a realistic place to check a change while keeping the live app's code untouched.
When a change is ready, the developer runs the test suite and reviews the preview. We then promote that reviewed change to the live app. The important separation is between each person's workspace and the shared live code, not between a development machine and a production machine. We can work in parallel without changing the files currently serving customers.
If a bug gets through, the path to a fix is short. A developer can reproduce it on the production host, make a fix in their own worktree, run tests, and check the preview. Once reviewed, we promote that exact change to the live app. We don't need a long cycle to recreate the production setup somewhere else and transfer the fix across environments.
That speed matters for our product. When a demo flow doesn't work, a team may not be able to show a buyer how the software works or help a customer learn a feature. We can investigate the issue and release a fix within minutes.
Pieter Levels' post about coding on the production server without a separate development environment got me thinking about how much trouble the handoff between environments can create. Amjad Masad's post got me to weigh whether developing in production is too risky, what it makes possible, and which safeguards make the benefits outweigh the risk.
We took that inspiration and built a workflow for a team: separate worktrees and AI sessions, shared production context, extensive tests, and previews before higher-risk changes go live. AI does more of the implementation, but our team still decides what to build, reviews every change, and owns what reaches customers. That's why I'm not looking back.
Record interactive online product demos in seconds. Edit, publish, and embed them for free.
© 2026 JoyDemo. All rights reserved.
