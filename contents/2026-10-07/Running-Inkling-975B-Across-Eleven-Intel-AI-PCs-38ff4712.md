---
source: "https://arxiv.org/abs/2610.07219"
hn_url: "https://news.ycombinator.com/item?id=49996430"
title: "Running Inkling (975B) Across Eleven Intel AI PCs"
article_title: "[2610.07219] Cascadia: Resident 975B MoE Inference on Eleven AI PCs"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "jackwsmth"
captured_at: "2026-10-07T18:31:48Z"
capture_tool: "hn-digest"
hn_id: 49996430
score: 2
comments: 0
posted_at: "2026-10-07T18:01:11Z"
tags:
  - hacker-news
---

# Running Inkling (975B) Across Eleven Intel AI PCs

- HN: [49996430](https://news.ycombinator.com/item?id=49996430)
- Source: [arxiv.org](https://arxiv.org/abs/2610.07219)
- Score: 2
- Comments: 0
- Posted: 2026-10-07T18:01:11Z

## Translation

Title: Running Inkling (975B) Across Eleven Intel AI PCs
Article title: [2610.07219] Cascadia: Resident 975B MoE Inference on Eleven AI PCs
Description: Abstract page for arXiv paper 2610.07219: Cascadia: Resident 975B MoE Inference on Eleven AI PCs

Article text:
Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 5 Oct 2026]
Title: Cascadia: Resident 975B MoE Inference on Eleven AI PCs
Abstract: Mixture-of-experts models make nearly trillion-parameter capacity accessible with sparse per-token computation, provided that the serving system can distribute the weights and coordinate their execution. We present Cascadia's resident execution of Inkling, a 975B-total/41B-active-parameter model, on eleven Intel Core Ultra X7 358H AI PCs, each with 64 GB of memory, Arc B390 integrated graphics and gigabit Ethernet. We contribute a custom resident MoE engine that preserves Inkling's routing rules, constructs compressed graphs for OpenVINO's fused iGPU primitives, and coordinates FP16 expert computation with FP32 output restoration. The engine fits six consecutive decoder layers per machine and represents dense feed-forward blocks as all-active expert slices, reducing measured dense-layer call time from approximately 8.1 to 4.5 ms. A streaming pipeline coordinates concurrent generation, while captured-state draft evaluation measures agreement with the deployed numerical path. Paired measurements at fifteen concurrency levels from 1 to 176 streams reach 60.29 aggregate decode tokens/s at 88 streams, with 46.87 tokens/s over the complete serving phases. At fifteen streams, median first-token latency is 6.05 s. Raising the context budget from the 1,024-position default, real prompts of 1k to 64k tokens recover the embedded code in all 19 measured answers, with first-token time growing as $aN+bN^2$ and decode latency growing approximately linearly, both bounded by a single-threaded CPU attention loop rather than by memory, which holds 512k positions per stream. Evaluation on captured fleet states separates the effects of vocabulary selection and weight quantization on draft agreement. Together, these contributions establish an execution and evaluation approach for large sparse models on distributed client systems with shared CPU-GPU memory.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

## Original Extract

Abstract page for arXiv paper 2610.07219: Cascadia: Resident 975B MoE Inference on Eleven AI PCs

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 5 Oct 2026]
Title: Cascadia: Resident 975B MoE Inference on Eleven AI PCs
Abstract: Mixture-of-experts models make nearly trillion-parameter capacity accessible with sparse per-token computation, provided that the serving system can distribute the weights and coordinate their execution. We present Cascadia's resident execution of Inkling, a 975B-total/41B-active-parameter model, on eleven Intel Core Ultra X7 358H AI PCs, each with 64 GB of memory, Arc B390 integrated graphics and gigabit Ethernet. We contribute a custom resident MoE engine that preserves Inkling's routing rules, constructs compressed graphs for OpenVINO's fused iGPU primitives, and coordinates FP16 expert computation with FP32 output restoration. The engine fits six consecutive decoder layers per machine and represents dense feed-forward blocks as all-active expert slices, reducing measured dense-layer call time from approximately 8.1 to 4.5 ms. A streaming pipeline coordinates concurrent generation, while captured-state draft evaluation measures agreement with the deployed numerical path. Paired measurements at fifteen concurrency levels from 1 to 176 streams reach 60.29 aggregate decode tokens/s at 88 streams, with 46.87 tokens/s over the complete serving phases. At fifteen streams, median first-token latency is 6.05 s. Raising the context budget from the 1,024-position default, real prompts of 1k to 64k tokens recover the embedded code in all 19 measured answers, with first-token time growing as $aN+bN^2$ and decode latency growing approximately linearly, both bounded by a single-threaded CPU attention loop rather than by memory, which holds 512k positions per stream. Evaluation on captured fleet states separates the effects of vocabulary selection and weight quantization on draft agreement. Together, these contributions establish an execution and evaluation approach for large sparse models on distributed client systems with shared CPU-GPU memory.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
