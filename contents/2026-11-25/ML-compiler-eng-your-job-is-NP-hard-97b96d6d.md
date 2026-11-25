---
source: "https://ianbarber.blog/2026/11/24/ml-compiler-eng-your-job-is-np-hard/"
hn_url: "https://news.ycombinator.com/item?id=50648253"
title: "ML compiler eng, your job is NP-hard"
article_title: "ML compiler eng, your job is NP-hard – Ian’s Blog"
image: "https://s0.wp.com/_si/?t=eyJpbWciOiJodHRwczpcL1wvaTAud3AuY29tXC9pYW5iYXJiZXIuYmxvZ1wvd3AtY29udGVudFwvdXBsb2Fkc1wvMjAyNFwvMTFcL2Nyb3BwZWQtbG9nby1yZW1vdmViZy1wcmV2aWV3LnBuZz9maXQ9NDI1JTJDNDI1JnNzbD0xIiwidHh0IjoiSWFuXHUyMDE5cyBCbG9nIiwidGVtcGxhdGUiOiJlZGdlIiwiZm9udCI6IiIsImJsb2dfaWQiOjIzODYzMzU2NX0.fq3U-V5MT1pTFs8OrXpT6sb83b9JfIPBwAcqO8AHRQQMQ"
author: "gmays"
captured_at: "2026-11-25T22:58:37Z"
capture_tool: "hn-digest"
hn_id: 50648253
score: 1
comments: 0
posted_at: "2026-11-25T22:06:04Z"
tags:
  - hacker-news
---

# ML compiler eng, your job is NP-hard

- HN: [50648253](https://news.ycombinator.com/item?id=50648253)
- Source: [ianbarber.blog](https://ianbarber.blog/2026/11/24/ml-compiler-eng-your-job-is-np-hard/)
- Score: 1
- Comments: 0
- Posted: 2026-11-25T22:06:04Z

## Translation

Title: ML compiler eng, your job is NP-hard
Article title: ML compiler eng, your job is NP-hard – Ian’s Blog
Description: I mean, sorta. From the Tensor Seeks Layout paper from Amazon: Layout Selection is NP-hard even when each operator has only a constant number of inputs and each tensor admits only a small number of layouts, and hardness persists in restricted settings that resemble common matmul-heavy graphs. The pa
[truncated]

Article text:
ML compiler eng, your job is NP-hard
I mean, sorta. From the Tensor Seeks Layout paper from Amazon:
Layout Selection is NP-hard even when each operator has only a constant number of inputs and each tensor admits only a small number of layouts, and hardness persists in restricted settings that resemble common matmul-heavy graphs.
The paper is very focused on formal analysis of the layout problem, for Trainium workloads. They build a layout selection exact solver, and on ResNet-50 they find the proven optimum using their cost model takes 63.9M cycles. In contrast, the awkward heuristic rules take… 20.9M cycles. The authors attribute this to a wonky cost model, which we will come back to.
There has been a lot of work in recent months on what layouts actually are , but this paper is about where and when you do conversions between them.
To jump down a level of granularity, in the Triton compiler this decision is mainly handled by the RemoveLayoutConversions pass: during compilation, Triton throws in layout conversions every time one might be needed, and this pass prunes them. It has the very sophisticated rule that conversions either cost 0 or 32xbytes . “0” means the conversion is a register permute, where everything happens locally. The higher value is for shuffles and shared memory round trips.
This made me wonder how wrong the rule actually was, versus the more structured approach in Amazon’s paper. I ran some experiments and found, yeah, the heuristic is pretty wrong. Relative to an elementwise add, the heuristic overestimates by about 4x for fp32, 19x for fp16, and sometimes even more for fp8. If you update to use the proper cost model though, not a lot changes. No tutorial kernels shift, and in general the performance changes are a bit of a wash: in testing 8 were slower, 4 were faster, 9 didn’t change at all.
The main reason is rematerialization: instead of converting the layout, walk back up the chain and recompute or load in the format you actually need. Conversions are mostly about cases where you can only emit one way, and then consume a different way. In those cases, the layout picks are arbitrary 1 , but they are mostly fine: the majority of kernels didn’t require a choice, and where they did, the default pick seemed to be the right one the vast majority of times.
We talked earlier about OpenAI’s effort to remove heuristics from the compiler, but this seems like an area where a heuristic is doing a reasonable job. If you dig into the cases where there is a performance mismatch though, you will find that the problem is also a compiler, but in this case a different one.
The pass exists to cut down layout conversions, which is generally good! But the output is then fed into another compiler, ptxas . That can be sensitive: in one case we found a matmul that got 15–26x faster, purely because of the instruction ordering. ptxas decides register allocation, and if it decides there aren’t enough it will spill to memory, significantly impacting performance 2 .
The finding that layout search is computationally hard in Amazon’s paper is valid, and good to share with your executives 3 . But, in reality, you don’t have to check all that many options: there are generally just a few places to put a conversion in a specific kernel 4 . Picking an option requires relying a potentially inaccurate model of how expensive those ops are given the subsequent (closed-source!) software stack and the hardware itself. And that seems to be the really tricky part.
And the pass is commented“TODO: add a proper heuristic”, so it’s not like this was a surprise ↩︎
in this case it moved one local_load 6 ops earlier, which caused ptxas to drop from 255 registers to 40. I don’t know why! ↩︎
We have to take these wins in the brief time before someone puts out 150M lines of Lean to prove that NP isn’t that hard at all. ↩︎
To be fair, the Amazon paper is doing the same thing for entire graphs of 10k+ operators, but it seems like it is somewhat analogous ↩︎
If you enjoyed this, subscribe and the next post will land in your inbox. No spam, unsubscribe any time.
Prefer a feed reader? Follow along via RSS .
ML compiler eng, your job is NP-hard
This weight does not spark joy
Make this forward pass as fast as possible
Now is the winter of our disclosure

## Original Extract

I mean, sorta. From the Tensor Seeks Layout paper from Amazon: Layout Selection is NP-hard even when each operator has only a constant number of inputs and each tensor admits only a small number of layouts, and hardness persists in restricted settings that resemble common matmul-heavy graphs. The pa
[truncated]

ML compiler eng, your job is NP-hard
I mean, sorta. From the Tensor Seeks Layout paper from Amazon:
Layout Selection is NP-hard even when each operator has only a constant number of inputs and each tensor admits only a small number of layouts, and hardness persists in restricted settings that resemble common matmul-heavy graphs.
The paper is very focused on formal analysis of the layout problem, for Trainium workloads. They build a layout selection exact solver, and on ResNet-50 they find the proven optimum using their cost model takes 63.9M cycles. In contrast, the awkward heuristic rules take… 20.9M cycles. The authors attribute this to a wonky cost model, which we will come back to.
There has been a lot of work in recent months on what layouts actually are , but this paper is about where and when you do conversions between them.
To jump down a level of granularity, in the Triton compiler this decision is mainly handled by the RemoveLayoutConversions pass: during compilation, Triton throws in layout conversions every time one might be needed, and this pass prunes them. It has the very sophisticated rule that conversions either cost 0 or 32xbytes . “0” means the conversion is a register permute, where everything happens locally. The higher value is for shuffles and shared memory round trips.
This made me wonder how wrong the rule actually was, versus the more structured approach in Amazon’s paper. I ran some experiments and found, yeah, the heuristic is pretty wrong. Relative to an elementwise add, the heuristic overestimates by about 4x for fp32, 19x for fp16, and sometimes even more for fp8. If you update to use the proper cost model though, not a lot changes. No tutorial kernels shift, and in general the performance changes are a bit of a wash: in testing 8 were slower, 4 were faster, 9 didn’t change at all.
The main reason is rematerialization: instead of converting the layout, walk back up the chain and recompute or load in the format you actually need. Conversions are mostly about cases where you can only emit one way, and then consume a different way. In those cases, the layout picks are arbitrary 1 , but they are mostly fine: the majority of kernels didn’t require a choice, and where they did, the default pick seemed to be the right one the vast majority of times.
We talked earlier about OpenAI’s effort to remove heuristics from the compiler, but this seems like an area where a heuristic is doing a reasonable job. If you dig into the cases where there is a performance mismatch though, you will find that the problem is also a compiler, but in this case a different one.
The pass exists to cut down layout conversions, which is generally good! But the output is then fed into another compiler, ptxas . That can be sensitive: in one case we found a matmul that got 15–26x faster, purely because of the instruction ordering. ptxas decides register allocation, and if it decides there aren’t enough it will spill to memory, significantly impacting performance 2 .
The finding that layout search is computationally hard in Amazon’s paper is valid, and good to share with your executives 3 . But, in reality, you don’t have to check all that many options: there are generally just a few places to put a conversion in a specific kernel 4 . Picking an option requires relying a potentially inaccurate model of how expensive those ops are given the subsequent (closed-source!) software stack and the hardware itself. And that seems to be the really tricky part.
And the pass is commented“TODO: add a proper heuristic”, so it’s not like this was a surprise ↩︎
in this case it moved one local_load 6 ops earlier, which caused ptxas to drop from 255 registers to 40. I don’t know why! ↩︎
We have to take these wins in the brief time before someone puts out 150M lines of Lean to prove that NP isn’t that hard at all. ↩︎
To be fair, the Amazon paper is doing the same thing for entire graphs of 10k+ operators, but it seems like it is somewhat analogous ↩︎
If you enjoyed this, subscribe and the next post will land in your inbox. No spam, unsubscribe any time.
Prefer a feed reader? Follow along via RSS .
ML compiler eng, your job is NP-hard
This weight does not spark joy
Make this forward pass as fast as possible
Now is the winter of our disclosure
