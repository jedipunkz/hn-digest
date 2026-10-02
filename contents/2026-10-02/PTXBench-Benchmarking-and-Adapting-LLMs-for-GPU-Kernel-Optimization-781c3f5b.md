---
source: "https://arxiv.org/abs/2608.17379"
hn_url: "https://news.ycombinator.com/item?id=49939792"
title: "PTXBench: Benchmarking and Adapting LLMs for GPU Kernel Optimization"
article_title: "[2608.17379] PTXBench: Benchmarking and Adapting LLMs for GPU Kernel Optimization with Architecture-specific PTX"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "matt_d"
captured_at: "2026-10-02T23:34:41Z"
capture_tool: "hn-digest"
hn_id: 49939792
score: 2
comments: 0
posted_at: "2026-10-02T23:21:30Z"
tags:
  - hacker-news
---

# PTXBench: Benchmarking and Adapting LLMs for GPU Kernel Optimization

- HN: [49939792](https://news.ycombinator.com/item?id=49939792)
- Source: [arxiv.org](https://arxiv.org/abs/2608.17379)
- Score: 2
- Comments: 0
- Posted: 2026-10-02T23:21:30Z

## Translation

Title: PTXBench: Benchmarking and Adapting LLMs for GPU Kernel Optimization
Article title: [2608.17379] PTXBench: Benchmarking and Adapting LLMs for GPU Kernel Optimization with Architecture-specific PTX
Description: Abstract page for arXiv paper 2608.17379: PTXBench: Benchmarking and Adapting LLMs for GPU Kernel Optimization with Architecture-specific PTX

Article text:
Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 18 Aug 2026 ( v1 ), last revised 28 Sep 2026 (this version, v3)]
Title: PTXBench: Benchmarking and Adapting LLMs for GPU Kernel Optimization with Architecture-specific PTX
Abstract: We introduce PTXBench, a benchmark for evaluating and adapting large language models (LLMs) to use architecture-specific PTX for GPU kernel optimization. PTXBench measures functional correctness, whether selected target instructions execute at runtime, and speedup over frontier libraries across GEMM and attention workloads on H100 and B200 GPUs. Our evaluation shows that architecture-specific PTX capability remains uneven: success rates fall substantially on complex attention backward workloads, and executing the target instructions does not necessarily translate into competitive performance. No evaluated model consistently matches frontier libraries across the suite. We further adapt Qwen3.6-27B using supervised fine-tuning. Repair-conditioned training improves several tasks, but generalization remains uneven; data coverage, balance, and the quality of the reasoning teacher matter in addition to dataset size. PTXBench provides an auditable testbed for measuring and improving LLMs' ability to exploit evolving GPU architectures.
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

Abstract page for arXiv paper 2608.17379: PTXBench: Benchmarking and Adapting LLMs for GPU Kernel Optimization with Architecture-specific PTX

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 18 Aug 2026 ( v1 ), last revised 28 Sep 2026 (this version, v3)]
Title: PTXBench: Benchmarking and Adapting LLMs for GPU Kernel Optimization with Architecture-specific PTX
Abstract: We introduce PTXBench, a benchmark for evaluating and adapting large language models (LLMs) to use architecture-specific PTX for GPU kernel optimization. PTXBench measures functional correctness, whether selected target instructions execute at runtime, and speedup over frontier libraries across GEMM and attention workloads on H100 and B200 GPUs. Our evaluation shows that architecture-specific PTX capability remains uneven: success rates fall substantially on complex attention backward workloads, and executing the target instructions does not necessarily translate into competitive performance. No evaluated model consistently matches frontier libraries across the suite. We further adapt Qwen3.6-27B using supervised fine-tuning. Repair-conditioned training improves several tasks, but generalization remains uneven; data coverage, balance, and the quality of the reasoning teacher matter in addition to dataset size. PTXBench provides an auditable testbed for measuring and improving LLMs' ability to exploit evolving GPU architectures.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
