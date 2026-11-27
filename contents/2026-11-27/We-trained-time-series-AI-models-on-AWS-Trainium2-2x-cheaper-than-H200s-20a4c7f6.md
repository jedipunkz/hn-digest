---
source: "https://twitter.com/martinmaritsch/status/2126195898795573437"
hn_url: "https://news.ycombinator.com/item?id=50663399"
title: "We trained time-series AI models on AWS Trainium2, 2x cheaper than H200s"
article_title: "Martin Maritsch on X: \"We trained OpenTSLM (Qwen3-4B backbone) on AWS Trainium2 and it was 2x cheaper than NVIDIA H200s.\nLoss was similar, and the one-node Trainium2 run took 1.5x as long as the two-node H200 run.\nAt AWS standard rates, that's $884 vs. $1816 for a full run.\nThe Trainium2 port of… /\n[truncated]"
image: "https://pbs.twimg.com/media/HYHD5vPbAAAZWuP?format=webp&name=large"
author: "masquare"
captured_at: "2026-11-27T06:59:51Z"
capture_tool: "hn-digest"
hn_id: 50663399
score: 1
comments: 0
posted_at: "2026-11-27T06:15:08Z"
tags:
  - hacker-news
---

# We trained time-series AI models on AWS Trainium2, 2x cheaper than H200s

- HN: [50663399](https://news.ycombinator.com/item?id=50663399)
- Source: [twitter.com](https://twitter.com/martinmaritsch/status/2126195898795573437)
- Score: 1
- Comments: 0
- Posted: 2026-11-27T06:15:08Z

## Translation

Title: We trained time-series AI models on AWS Trainium2, 2x cheaper than H200s
Article title: Martin Maritsch on X: "We trained OpenTSLM (Qwen3-4B backbone) on AWS Trainium2 and it was 2x cheaper than NVIDIA H200s.
Loss was similar, and the one-node Trainium2 run took 1.5x as long as the two-node H200 run.
At AWS standard rates, that's $884 vs. $1816 for a full run.
The Trainium2 port of… /
[truncated]
Description: We trained OpenTSLM (Qwen3-4B backbone) on AWS Trainium2 and it was 2x cheaper than NVIDIA H200s.
Loss was similar, and the one-node Trainium2 run took 1.5x as long as the two-node H200 run.
At AWS standard rates, that's $884 vs. $1816 for a full run.
The Trainium2 port of our architecture was s…

Article text:
Martin Maritsch on X: "We trained OpenTSLM (Qwen3-4B backbone) on AWS Trainium2 and it was 2x cheaper than NVIDIA H200s.
Loss was similar, and the one-node Trainium2 run took 1.5x as long as the two-node H200 run.
At AWS standard rates, that's $884 vs. $1816 for a full run.
The Trainium2 port of our architecture was straightforward:"
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
Martin Maritsch @martinmaritsch We trained OpenTSLM (Qwen3-4B backbone) on AWS Trainium2 and it was 2x cheaper than NVIDIA H200s.
Loss was similar, and the one-node Trainium2 run took 1.5x as long as the two-node H200 run.
At AWS standard rates, that's $884 vs. $1816 for a full run.
The Trainium2 port of our architecture was straightforward: Aionic Labs 6:06 AM · Nov 27, 2026 · 57 Views 1 2 2
Martin Maritsch @martinmaritsch 53m We kept the original model in PyTorch, precompiled fixed batch shapes and used 4-way tensor parallelism to fit the model on the chips.
More on OpenTSLM in our @awscloud blog: Reason over time-series data with OpenTSLM on Amazon SageMaker JumpStart | Amazon Web Services From aws.amazon.com 1 11

## Original Extract

We trained OpenTSLM (Qwen3-4B backbone) on AWS Trainium2 and it was 2x cheaper than NVIDIA H200s.
Loss was similar, and the one-node Trainium2 run took 1.5x as long as the two-node H200 run.
At AWS standard rates, that's $884 vs. $1816 for a full run.
The Trainium2 port of our architecture was s…

Martin Maritsch on X: "We trained OpenTSLM (Qwen3-4B backbone) on AWS Trainium2 and it was 2x cheaper than NVIDIA H200s.
Loss was similar, and the one-node Trainium2 run took 1.5x as long as the two-node H200 run.
At AWS standard rates, that's $884 vs. $1816 for a full run.
The Trainium2 port of our architecture was straightforward:"
See what’s happening and join the conversation
Continue with phone Continue with Apple Continue with Google or Log in with username or email Relevant people
Martin Maritsch @martinmaritsch We trained OpenTSLM (Qwen3-4B backbone) on AWS Trainium2 and it was 2x cheaper than NVIDIA H200s.
Loss was similar, and the one-node Trainium2 run took 1.5x as long as the two-node H200 run.
At AWS standard rates, that's $884 vs. $1816 for a full run.
The Trainium2 port of our architecture was straightforward: Aionic Labs 6:06 AM · Nov 27, 2026 · 57 Views 1 2 2
Martin Maritsch @martinmaritsch 53m We kept the original model in PyTorch, precompiled fixed batch shapes and used 4-way tensor parallelism to fit the model on the chips.
More on OpenTSLM in our @awscloud blog: Reason over time-series data with OpenTSLM on Amazon SageMaker JumpStart | Amazon Web Services From aws.amazon.com 1 11
