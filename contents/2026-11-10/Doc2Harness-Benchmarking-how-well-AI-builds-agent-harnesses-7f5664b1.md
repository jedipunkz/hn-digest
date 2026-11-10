---
source: "https://arxiv.org/abs/2610.20344"
hn_url: "https://news.ycombinator.com/item?id=50445107"
title: "Doc2Harness: Benchmarking how well AI builds agent harnesses"
article_title: "[2610.20344] Doc2Harness: Benchmarking LLMs on Long-Horizon Harness Engineering of Coding Agents"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "waximabbax"
captured_at: "2026-11-10T20:16:26Z"
capture_tool: "hn-digest"
hn_id: 50445107
score: 1
comments: 0
posted_at: "2026-11-10T19:31:36Z"
tags:
  - hacker-news
---

# Doc2Harness: Benchmarking how well AI builds agent harnesses

- HN: [50445107](https://news.ycombinator.com/item?id=50445107)
- Source: [arxiv.org](https://arxiv.org/abs/2610.20344)
- Score: 1
- Comments: 0
- Posted: 2026-11-10T19:31:36Z

## Translation

Title: Doc2Harness: Benchmarking how well AI builds agent harnesses
Article title: [2610.20344] Doc2Harness: Benchmarking LLMs on Long-Horizon Harness Engineering of Coding Agents
Description: Abstract page for arXiv paper 2610.20344: Doc2Harness: Benchmarking LLMs on Long-Horizon Harness Engineering of Coding Agents

Article text:
Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
Computer Science > Software Engineering
[Submitted on 14 Oct 2026]
Title: Doc2Harness: Benchmarking LLMs on Long-Horizon Harness Engineering of Coding Agents
Abstract: Agents driven by large language models are increasingly deployed on long-horizon tasks, where they use tools and modify real workspaces. How well they perform depends as much on the harness around the model as on the model itself. Existing benchmarks measure downstream task success under a fixed harness, leaving harness engineering itself unexplored. We propose Doc2Harness, a two-stage benchmark that evaluates long-horizon harness engineering directly. In Stage 1, a build model operating inside a fixed builder harness receives an interface-only contract skeleton and a development document that consolidates behavior across twelve frontier harnesses, and then constructs a complete candidate harness in an isolated container. In Stage 2, the frozen candidate harness runs in containers separate from the construction environment, one per test case, and is driven end to end by a scripted mock model. We score each test case using four observable channels, namely the requests sent to the model, the files on disk, the process exit code, and the command output; positive controls rule out vacuous passes. Across nineteen frontier models, each evaluated with three independent builds on 328 test cases, GPT-6 Astra achieves the highest pass rate at 53.05%, while the weakest model reaches 27.44%, a gap of roughly twenty-six percentage points. Even the strongest model fails nearly half of the test cases, and failures cluster around how cross-cutting functionality is split across modules. To our knowledge, Doc2Harness provides the first deterministic benchmark for measuring long-horizon harness engineering.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

## Original Extract

Abstract page for arXiv paper 2610.20344: Doc2Harness: Benchmarking LLMs on Long-Horizon Harness Engineering of Coding Agents

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
Computer Science > Software Engineering
[Submitted on 14 Oct 2026]
Title: Doc2Harness: Benchmarking LLMs on Long-Horizon Harness Engineering of Coding Agents
Abstract: Agents driven by large language models are increasingly deployed on long-horizon tasks, where they use tools and modify real workspaces. How well they perform depends as much on the harness around the model as on the model itself. Existing benchmarks measure downstream task success under a fixed harness, leaving harness engineering itself unexplored. We propose Doc2Harness, a two-stage benchmark that evaluates long-horizon harness engineering directly. In Stage 1, a build model operating inside a fixed builder harness receives an interface-only contract skeleton and a development document that consolidates behavior across twelve frontier harnesses, and then constructs a complete candidate harness in an isolated container. In Stage 2, the frozen candidate harness runs in containers separate from the construction environment, one per test case, and is driven end to end by a scripted mock model. We score each test case using four observable channels, namely the requests sent to the model, the files on disk, the process exit code, and the command output; positive controls rule out vacuous passes. Across nineteen frontier models, each evaluated with three independent builds on 328 test cases, GPT-6 Astra achieves the highest pass rate at 53.05%, while the weakest model reaches 27.44%, a gap of roughly twenty-six percentage points. Even the strongest model fails nearly half of the test cases, and failures cluster around how cross-cutting functionality is split across modules. To our knowledge, Doc2Harness provides the first deterministic benchmark for measuring long-horizon harness engineering.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
