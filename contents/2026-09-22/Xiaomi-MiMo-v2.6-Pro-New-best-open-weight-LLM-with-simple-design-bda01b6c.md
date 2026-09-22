---
source: "https://sebastianraschka.com/blog/2026/mimo-v2-6-pro-architecture-training-notes.html"
hn_url: "https://news.ycombinator.com/item?id=49801771"
title: "Xiaomi MiMo v2.6 Pro: New best open-weight LLM with simple design"
article_title: "MiMo-V2.6 Pro Training Notes | Sebastian Raschka, PhD"
image: "https://sebastianraschka.com/images/blog/2026/mimo-v2-6-pro-architecture-training-notes/mimo-v2-6-pro-comparison.webp"
author: "ModelForge"
captured_at: "2026-09-22T14:23:21Z"
capture_tool: "hn-digest"
hn_id: 49801771
score: 1
comments: 0
posted_at: "2026-09-22T14:21:34Z"
tags:
  - hacker-news
---

# Xiaomi MiMo v2.6 Pro: New best open-weight LLM with simple design

- HN: [49801771](https://news.ycombinator.com/item?id=49801771)
- Source: [sebastianraschka.com](https://sebastianraschka.com/blog/2026/mimo-v2-6-pro-architecture-training-notes.html)
- Score: 1
- Comments: 0
- Posted: 2026-09-22T14:21:34Z

## Translation

Title: Xiaomi MiMo v2.6 Pro: New best open-weight LLM with simple design
Article title: MiMo-V2.6 Pro Training Notes | Sebastian Raschka, PhD
Description: Notes on MiMo-V2.6 Pro’s GQA and sliding-window attention, agent training tasks, reward signals, and large RL batches.

Article text:
Skip to main content
Sebastian
Raschka
-->
🌙
Search this website
Submit search
Home
Blog
Blog -->
Books
AI Newsletter -->
Courses
LLM Gallery
LLMs From Scratch
Reasoning Models
Talks
More
Blog Archive
Research
About
Contact -->
More -->
MiMo-V2.6 Pro Architecture and Training Notes
Sep 22, 2026
by Sebastian Raschka
Xiaomi’s new MiMo-V2.6 Pro is “simply” the best (for now). Despite its simple architecture design it’s currently No.1 in the open-weight benchmarks (weighted average).
With “simple,” I mean a classic Grouped Query Attention (GQA) with Sliding Window Attention (SWA) at a tiny 128-token window size.
So, that underlines one of the points I’ve been trying to make in recent months: most of the progress still comes from the data and post-training recipe improvements. Fancy attention variants are just mostly efficiency tweaks.
What are some of the training data improvements and recipe improvements? The MiMo team shared a pretty detailed technical report . Lots to carefully digest there, but in short, there are a few things that stood out:
An increase in agent tasks; also training across different harnesses (the average DeepSWE pass@1 accuracy on held-out harnesses improved from approximately 50% -> 66%).
Better reward signals: they replaced a simple correctness verifier with an agentic grader that looks at the execution traces as well.
Large RL batches (1,568 prompts × 16 rollouts = 25,088 trajectories) and 2.7–3.7 billion training tokens per update (unclear, though, what the predecessor used).
Source: website version of my Substack note .

## Original Extract

Notes on MiMo-V2.6 Pro’s GQA and sliding-window attention, agent training tasks, reward signals, and large RL batches.

Skip to main content
Sebastian
Raschka
-->
🌙
Search this website
Submit search
Home
Blog
Blog -->
Books
AI Newsletter -->
Courses
LLM Gallery
LLMs From Scratch
Reasoning Models
Talks
More
Blog Archive
Research
About
Contact -->
More -->
MiMo-V2.6 Pro Architecture and Training Notes
Sep 22, 2026
by Sebastian Raschka
Xiaomi’s new MiMo-V2.6 Pro is “simply” the best (for now). Despite its simple architecture design it’s currently No.1 in the open-weight benchmarks (weighted average).
With “simple,” I mean a classic Grouped Query Attention (GQA) with Sliding Window Attention (SWA) at a tiny 128-token window size.
So, that underlines one of the points I’ve been trying to make in recent months: most of the progress still comes from the data and post-training recipe improvements. Fancy attention variants are just mostly efficiency tweaks.
What are some of the training data improvements and recipe improvements? The MiMo team shared a pretty detailed technical report . Lots to carefully digest there, but in short, there are a few things that stood out:
An increase in agent tasks; also training across different harnesses (the average DeepSWE pass@1 accuracy on held-out harnesses improved from approximately 50% -> 66%).
Better reward signals: they replaced a simple correctness verifier with an agentic grader that looks at the execution traces as well.
Large RL batches (1,568 prompts × 16 rollouts = 25,088 trajectories) and 2.7–3.7 billion training tokens per update (unclear, though, what the predecessor used).
Source: website version of my Substack note .
