---
source: "https://yanahchat.com"
hn_url: "https://news.ycombinator.com/item?id=50360513"
title: "Local AI app that runs an 80B model on an 8GB Mac"
article_title: "Yanah.ai — Free Private AI On Your MacBook"
image: ""
author: "giovanniwants"
captured_at: "2026-11-04T01:56:02Z"
capture_tool: "hn-digest"
hn_id: 50360513
score: 1
comments: 1
posted_at: "2026-11-04T01:30:30Z"
tags:
  - hacker-news
---

# Local AI app that runs an 80B model on an 8GB Mac

- HN: [50360513](https://news.ycombinator.com/item?id=50360513)
- Source: [yanahchat.com](https://yanahchat.com)
- Score: 1
- Comments: 1
- Posted: 2026-11-04T01:30:30Z

## Translation

Title: Local AI app that runs an 80B model on an 8GB Mac
Article title: Yanah.ai — Free Private AI On Your MacBook

Article text:
YANAH.ai
Download Yanah Chat and you'll have a private, working chatbot, even when the internet is off.
Needs an M1 Mac or newer (late 2020 onward)
File size ~40GB standard install (up to ~150GB with upcoming 177B)
Donate
Contact
Architecture
Performance
FAQ (Everyone)
FAQ (Devs)
Thinking Forward
Private Local AI
AI companies having
all your private chats?
Ya nah.
Private: your data and your AI stay on your Mac, under your control. No cloud, no subscriptions, no one harvesting your conversations. Huge AI models that shouldn't fit on a laptop, running at chat speed. Private, no-fuss installation. No internet needed to chat. Environmentally sound AI, democratized for the masses.
HD Space needed is ~40GB standard install (up to 150GB with upcoming 177B).
Donate
Yanah Chat is free. If it's useful to you, a donation helps keep it that way.
How can an AI run on my laptop? Doesn't AI need giant data centers? The AI has to hold all the information of the internet, right? And it needs giant servers to process all that information for chat, right?
It sounds impossible, but an AI doesn't actually store the internet. This is the strangest thing for people to understand. It doesn't hold all knowledge and then look it up when you ask a question.
Think of it less like an encyclopedia and more like someone who has read an enormous amount and absorbed how language and ideas fit together. When you ask a question, it doesn't look anything up. It writes the answer one word at a time, each time predicting the most sensible next word given everything so far. What it learned is stored as billions of numbers called "parameters," and those fit in a file on your drive.
The giant data centers are mostly for training AIs and serving millions of people at once. Running one for a single person is a far smaller job, and Yanah uses new techniques to make it fit on a regular Mac. A common assumption is that you need a billion-dollar data center because a good AI would need trillions of terabytes of hard drive space. You don't. A genuinely great AI can be fit into less than 50 gigabytes. And it doesn't even need the internet to run.
Yanah running locally on a MacBook Pro. No cloud. No subscription.
*Models tested were a 64gb MacBook Pro M1 and an M3 MacBook Pro M3. Those are confirmed. Most other stats and claims were made by AI Calculating the math. I have made every effort to ensure the information is rock solid. See full benchmarks →
How Yanah runs an 80-billion-parameter AI locally on an 8GB Mac at 5–8 tokens/sec—while retaining 90–92% of full FP16 intelligence.
Why everyone assumes this is impossible
If you tell a systems engineer—or ask another AI—whether an 80-billion-parameter model ( 28.99 GB on disk) can run locally on an 8 GB MacBook with Wi-Fi and Bluetooth turned off at 5 to 8 tokens per second , their first instinct is to call it impossible.
In standard engines like stock mlx-lm or llama.cpp , it is impossible:
The RAM Wall: Loading or memory-mapping a 29 GB weight file into 8 GB of unified memory triggers an immediate Out-of-Memory (OOM) crash or grinds macOS swap to a standstill ( < 0.1 tokens/sec ).
The 2-Bit Quality Cliff: Crushing an entire model with standard uniform 2-bit quantization destroys its reasoning—dropping hidden-state fidelity to ~20–22% of the original FP16 model and causing repetitive loops.
Yanah ( moe_prefetch ) solves both problems by exploiting the physical structure of Mixture-of-Experts (MoE) models, applying surgical three-tier quantization, and treating Apple Silicon's NVMe SSD as a predictive, zero-copy extension of Unified Memory.
1. The core secret: 80B on disk, only 3B active per word
Yanah's 80B tier runs Qwen3-Next-80B-A3B , a hybrid Mixture-of-Experts architecture with 48 layers :
36 Gated DeltaNet layers: A fast, constant-memory recurrent state-space architecture whose memory footprint never grows with conversation length.
12 Full-Attention layers: Placed every 4th layer for exact long-range recall.
1 Shared Expert + 512 Routed Specialist Experts per layer: 24,576 routed experts in total across the model.
On disk, those 24,576 routed experts make up ~96% of the model (~77B parameters) . The "always-on" dense backbone—token embeddings, Attention, Gated DeltaNet, Routers, Shared Experts, and lm_head —is only ~3B parameters (~4% of the model, or 1.50 GB in RAM) .
For any single word piece (token), the router in each layer wakes up only 10 out of 512 experts (1.95%) . The other 502 experts per layer (98.05%) do nothing . You never need all 80 billion parameters in RAM at once—you only need the 1.50 GB backbone permanently in RAM, plus a way to stream the 10 chosen experts per layer off the SSD right before the GPU needs them.
2. Surgical hybrid quantization: keeping 90–92% of FP16 intelligence
People assume we use a standard 2-bit model to fit on constrained hardware. We don't.
Uniform 2-bit quantization breaks Qwen3-Next in two places: (1) its 36 Gated DeltaNet layers update a rolling recurrent memory state on every token, so 2-bit projections inject ~37% noise into short-term memory at every step; and (2) a 2-bit router only has 4 weight levels ( 0, 1, 2, 3 ), which scrambles 512-way expert rankings by more than 55%.
Instead, Yanah uses a three-tier hybrid precision layout ( quantize_2bit.py ):
Why "2-bit g32 " is actually 3.0 bits per weight: Because Yanah uses a tight group size of 32 with BF16 affine scales and biases, one-third of every 1.125 MiB expert record ( 384 KiB ) is high-precision BF16 scale and bias metadata . And because 10 routed experts are weighted and summed on top of the 4-bit Shared Expert at every layer, uncorrelated quantization noise across the 10 experts partially cancels out.
Across all 48 layers of Qwen3-80B, the measured relative RMS weight error of Yanah's 2-bit g32 experts sits between 36.04% (Layer 0) and 36.91% (Layer 41), with a median of 36.68% .
3. Zero-copy SSD expert streaming ( moe_prefetch )
Yanah's Source-Available / Fair-Source streaming engine ( moe_prefetch v00.1.00 , adapting mechanisms from SlotStream for 8 GB Macs) bypasses macOS swap completely:
Lazy load & expert stripping ( mlx_backend.py ): At startup, Yanah opens the checkpoint with load_model(lazy=True) , strips all 48 switch_mlp expert stacks before evaluation, and materializes only the 1.50 GB dense backbone into RAM.
Contiguous page-aligned records ( repack.py ): In standard checkpoints, one expert's 9 tensors ( gate_proj , up_proj , down_proj × weight , scales , biases ) sit scattered across multi-gigabyte shards. Yanah repacks all 9 pieces back-to-back into a single 16 KiB page-aligned record inside experts.bin so each expert is fetched in one preadv system call .
Zero-copy Unified Memory DMA ( MLXExpertMemory ): Because Apple Silicon CPU and GPU share physical RAM pages, background I/O threads call preadv directly into writable uint8 MLX array buffers via Python's buffer protocol. MLX's mx.gather_qmm kernel indexes those pool slots directly— an expert is never copied in memory after the SSD read .
4. The five speed engines behind 5–8 tok/s on 8 GB (and 15–19 tok/s on 24 GB+)
If the GPU waited idly for the SSD at every layer, decoding would crawl. Yanah overlaps SSD I/O with GPU compute through five mechanisms:
1. Two-layer GPU router lookahead ( predictor.py & route_and_forecast ): Right after Layer L finishes its Attention or Gated DeltaNet step, its hidden state already predicts Layer L+1 's routing with ~73% agreement ("attention tap"). In the exact same GPU sync that reads Layer L 's router, Yanah runs Layer L+1 and L+2 's 8-bit routers on the GPU (zero extra GPU syncs, zero CPU matmuls) and starts background SSD reads for the top candidates ( prob >= 0.062 ) two layers ahead of time.
2. CLOCK cache pool + quarantine speculative partition ( pool.py ): A global CLOCK cache keeps hot experts resident in RAM. To prevent wrong lookahead guesses from evicting useful cached experts, all speculative prefetches read into an isolated 64-slot quarantine partition ( spec_slot_cap ). A prefetched expert joins the main cache ( adopt() ) only when a router actually demands it; unused forecasts are discarded with zero cache evictions .
3. Priority I/O lane budgeting ( io_lanes.py ): Demand misses run as high-priority batches at Queue Depth 32–64 . Background speculative reads share an 8-lane budget ( IOLaneBudget ) that automatically pauses the instant a demand batch starts, and promotes any in-flight forecast ticket if a router demands it mid-read.
4. Resident-first GPU overlap & single-sync lazy eval ( moe_ready / moe_rest ): During decode, if 7 of a layer's 10 required experts are already in RAM and 3 are still reading from SSD, the GPU immediately starts computing the 7 resident experts + the 4-bit Shared Expert while waiting on the 3 SSD misses , then merges the remaining columns bit-identically. Lazy evaluation ( mx.async_eval + memory fence() ) cuts CPU↔GPU syncs in half from 96 to 48 per token.
5. Scan-resistant prefill sweep ( sweep.py ): Prompts of 16+ tokens touch almost all 512 experts per layer. Instead of churning the cache pool, prompts use a dedicated double-buffered sweep ( 2 × 32 slots at Queue Depth 12) that copies resident experts without touching reference bits, merges adjacent disk reads, and admits only a fair-share quota of the most-used experts on the final chunk.
5. The hard 8 GB RAM & bandwidth math
Here is the exact process memory budget on an 8 GB Mac (leaving ~3 GB for macOS with zero swap):
Why the token speed math works:
Per-token expert work: 48 layers × 10 experts = 480 expert uses per token . At 2-bit ( 1.125 MiB /expert), even a 100% cold-miss token is only 540 MiB —not 29 GB .
On an 8 GB Mac (~2,300 expert slots in RAM): Between resident CLOCK hits and 2-layer lookahead adoption, ~55–65% of expert requests are already in RAM or pre-streamed. That leaves only ~190–240 MiB of net SSD wait per token (~70–90 ms at 2.5–3.0 GB/s NVMe) plus ~50 ms of overlapped GPU compute = 5 to 8 tokens/sec .
On a 24 GB+ Mac (~16+ GB expert pool): Over 60% of the entire 80B model stays resident in RAM (85–90%+ hit rate), dropping SSD wait to ~10–20 ms/token and running at 15 to 19 tokens/sec .
6. Built-in server, instant state rollback & memory
3-level conversation & persona state reuse ( mlx_server.py ): Takes zero-copy snapshots of the immutable 74 MiB Gated DeltaNet recurrent arrays at prompt boundaries and system-prompt anchors ( cache_anchor ), so multi-turn chats and persona switches resume instantly without re-reading the prompt.
Adaptive MTP speculation with zero-compute rollback ( speculative.py ): Records per-position Gated DeltaNet states during verification so rejected draft tokens roll back in 0 ms , while an adaptive wall-clock governor automatically switches drafting off on 8 GB pools when SSD misses outweigh draft gains.
Markdown + SQLite long-term memory: Yanah's cross-conversation memory lives in human-readable Markdown files indexed by local SQLite for zero-lag retrieval, separate from the model's KV cache.
*Models tested were a 64gb MacBook Pro M1 and an M3 MacBook Pro M3. Those are confirmed. Most other stats and claims were made by AI Calculating the math. I have made every effort to ensure the information is rock solid.
A 177B model runs faster on a MacBook Pro than a dense 27B model does.
That's the whole story in one number. On a 64GB Apple Silicon MacBook Pro, a 177B MoE model generates about 20 tokens/sec, faster than a dense 27B model on the same machine, because only a small slice of the big model works on each word.
* Performance figures are estimates based on internal testing and published benchmarks. Results may vary by Mac model, memory, storage, model choice and conversation length.
Why a bigger MoE can beat a smaller dense model
Speed tracks the parameters computed per token, not the total size.
A dense 27B computes all 27B parameters on every token, which saturates the memory bus.
An MoE model wakes only its a

[truncated]

## Original Extract

YANAH.ai
Download Yanah Chat and you'll have a private, working chatbot, even when the internet is off.
Needs an M1 Mac or newer (late 2020 onward)
File size ~40GB standard install (up to ~150GB with upcoming 177B)
Donate
Contact
Architecture
Performance
FAQ (Everyone)
FAQ (Devs)
Thinking Forward
Private Local AI
AI companies having
all your private chats?
Ya nah.
Private: your data and your AI stay on your Mac, under your control. No cloud, no subscriptions, no one harvesting your conversations. Huge AI models that shouldn't fit on a laptop, running at chat speed. Private, no-fuss installation. No internet needed to chat. Environmentally sound AI, democratized for the masses.
HD Space needed is ~40GB standard install (up to 150GB with upcoming 177B).
Donate
Yanah Chat is free. If it's useful to you, a donation helps keep it that way.
How can an AI run on my laptop? Doesn't AI need giant data centers? The AI has to hold all the information of the internet, right? And it needs giant servers to process all that information for chat, right?
It sounds impossible, but an AI doesn't actually store the internet. This is the strangest thing for people to understand. It doesn't hold all knowledge and then look it up when you ask a question.
Think of it less like an encyclopedia and more like someone who has read an enormous amount and absorbed how language and ideas fit together. When you ask a question, it doesn't look anything up. It writes the answer one word at a time, each time predicting the most sensible next word given everything so far. What it learned is stored as billions of numbers called "parameters," and those fit in a file on your drive.
The giant data centers are mostly for training AIs and serving millions of people at once. Running one for a single person is a far smaller job, and Yanah uses new techniques to make it fit on a regular Mac. A common assumption is that you need a billion-dollar data center because a good AI would need trillions of terabytes of hard drive space. You don't. A genuinely great AI can be fit into less than 50 gigabytes. And it doesn't even need the internet to run.
Yanah running locally on a MacBook Pro. No cloud. No subscription.
*Models tested were a 64gb MacBook Pro M1 and an M3 MacBook Pro M3. Those are confirmed. Most other stats and claims were made by AI Calculating the math. I have made every effort to ensure the information is rock solid. See full benchmarks →
How Yanah runs an 80-billion-parameter AI locally on an 8GB Mac at 5–8 tokens/sec—while retaining 90–92% of full FP16 intelligence.
Why everyone assumes this is impossible
If you tell a systems engineer—or ask another AI—whether an 80-billion-parameter model ( 28.99 GB on disk) can run locally on an 8 GB MacBook with Wi-Fi and Bluetooth turned off at 5 to 8 tokens per second , their first instinct is to call it impossible.
In standard engines like stock mlx-lm or llama.cpp , it is impossible:
The RAM Wall: Loading or memory-mapping a 29 GB weight file into 8 GB of unified memory triggers an immediate Out-of-Memory (OOM) crash or grinds macOS swap to a standstill ( < 0.1 tokens/sec ).
The 2-Bit Quality Cliff: Crushing an entire model with standard uniform 2-bit quantization destroys its reasoning—dropping hidden-state fidelity to ~20–22% of the original FP16 model and causing repetitive loops.
Yanah ( moe_prefetch ) solves both problems by exploiting the physical structure of Mixture-of-Experts (MoE) models, applying surgical three-tier quantization, and treating Apple Silicon's NVMe SSD as a predictive, zero-copy extension of Unified Memory.
1. The core secret: 80B on disk, only 3B active per word
Yanah's 80B tier runs Qwen3-Next-80B-A3B , a hybrid Mixture-of-Experts architecture with 48 layers :
36 Gated DeltaNet layers: A fast, constant-memory recurrent state-space architecture whose memory footprint never grows with conversation length.
12 Full-Attention layers: Placed every 4th layer for exact long-range recall.
1 Shared Expert + 512 Routed Specialist Experts per layer: 24,576 routed experts in total across the model.
On disk, those 24,576 routed experts make up ~96% of the model (~77B parameters) . The "always-on" dense backbone—token embeddings, Attention, Gated DeltaNet, Routers, Shared Experts, and lm_head —is only ~3B parameters (~4% of the model, or 1.50 GB in RAM) .
For any single word piece (token), the router in each layer wakes up only 10 out of 512 experts (1.95%) . The other 502 experts per layer (98.05%) do nothing . You never need all 80 billion parameters in RAM at once—you only need the 1.50 GB backbone permanently in RAM, plus a way to stream the 10 chosen experts per layer off the SSD right before the GPU needs them.
2. Surgical hybrid quantization: keeping 90–92% of FP16 intelligence
People assume we use a standard 2-bit model to fit on constrained hardware. We don't.
Uniform 2-bit quantization breaks Qwen3-Next in two places: (1) its 36 Gated DeltaNet layers update a rolling recurrent memory state on every token, so 2-bit projections inject ~37% noise into short-term memory at every step; and (2) a 2-bit router only has 4 weight levels ( 0, 1, 2, 3 ), which scrambles 512-way expert rankings by more than 55%.
Instead, Yanah uses a three-tier hybrid precision layout ( quantize_2bit.py ):
Why "2-bit g32 " is actually 3.0 bits per weight: Because Yanah uses a tight group size of 32 with BF16 affine scales and biases, one-third of every 1.125 MiB expert record ( 384 KiB ) is high-precision BF16 scale and bias metadata . And because 10 routed experts are weighted and summed on top of the 4-bit Shared Expert at every layer, uncorrelated quantization noise across the 10 experts partially cancels out.
Across all 48 layers of Qwen3-80B, the measured relative RMS weight error of Yanah's 2-bit g32 experts sits between 36.04% (Layer 0) and 36.91% (Layer 41), with a median of 36.68% .
3. Zero-copy SSD expert streaming ( moe_prefetch )
Yanah's Source-Available / Fair-Source streaming engine ( moe_prefetch v00.1.00 , adapting mechanisms from SlotStream for 8 GB Macs) bypasses macOS swap completely:
Lazy load & expert stripping ( mlx_backend.py ): At startup, Yanah opens the checkpoint with load_model(lazy=True) , strips all 48 switch_mlp expert stacks before evaluation, and materializes only the 1.50 GB dense backbone into RAM.
Contiguous page-aligned records ( repack.py ): In standard checkpoints, one expert's 9 tensors ( gate_proj , up_proj , down_proj × weight , scales , biases ) sit scattered across multi-gigabyte shards. Yanah repacks all 9 pieces back-to-back into a single 16 KiB page-aligned record inside experts.bin so each expert is fetched in one preadv system call .
Zero-copy Unified Memory DMA ( MLXExpertMemory ): Because Apple Silicon CPU and GPU share physical RAM pages, background I/O threads call preadv directly into writable uint8 MLX array buffers via Python's buffer protocol. MLX's mx.gather_qmm kernel indexes those pool slots directly— an expert is never copied in memory after the SSD read .
4. The five speed engines behind 5–8 tok/s on 8 GB (and 15–19 tok/s on 24 GB+)
If the GPU waited idly for the SSD at every layer, decoding would crawl. Yanah overlaps SSD I/O with GPU compute through five mechanisms:
1. Two-layer GPU router lookahead ( predictor.py & route_and_forecast ): Right after Layer L finishes its Attention or Gated DeltaNet step, its hidden state already predicts Layer L+1 's routing with ~73% agreement ("attention tap"). In the exact same GPU sync that reads Layer L 's router, Yanah runs Layer L+1 and L+2 's 8-bit routers on the GPU (zero extra GPU syncs, zero CPU matmuls) and starts background SSD reads for the top candidates ( prob >= 0.062 ) two layers ahead of time.
2. CLOCK cache pool + quarantine speculative partition ( pool.py ): A global CLOCK cache keeps hot experts resident in RAM. To prevent wrong lookahead guesses from evicting useful cached experts, all speculative prefetches read into an isolated 64-slot quarantine partition ( spec_slot_cap ). A prefetched expert joins the main cache ( adopt() ) only when a router actually demands it; unused forecasts are discarded with zero cache evictions .
3. Priority I/O lane budgeting ( io_lanes.py ): Demand misses run as high-priority batches at Queue Depth 32–64 . Background speculative reads share an 8-lane budget ( IOLaneBudget ) that automatically pauses the instant a demand batch starts, and promotes any in-flight forecast ticket if a router demands it mid-read.
4. Resident-first GPU overlap & single-sync lazy eval ( moe_ready / moe_rest ): During decode, if 7 of a layer's 10 required experts are already in RAM and 3 are still reading from SSD, the GPU immediately starts computing the 7 resident experts + the 4-bit Shared Expert while waiting on the 3 SSD misses , then merges the remaining columns bit-identically. Lazy evaluation ( mx.async_eval + memory fence() ) cuts CPU↔GPU syncs in half from 96 to 48 per token.
5. Scan-resistant prefill sweep ( sweep.py ): Prompts of 16+ tokens touch almost all 512 experts per layer. Instead of churning the cache pool, prompts use a dedicated double-buffered sweep ( 2 × 32 slots at Queue Depth 12) that copies resident experts without touching reference bits, merges adjacent disk reads, and admits only a fair-share quota of the most-used experts on the final chunk.
5. The hard 8 GB RAM & bandwidth math
Here is the exact process memory budget on an 8 GB Mac (leaving ~3 GB for macOS with zero swap):
Why the token speed math works:
Per-token expert work: 48 layers × 10 experts = 480 expert uses per token . At 2-bit ( 1.125 MiB /expert), even a 100% cold-miss token is only 540 MiB —not 29 GB .
On an 8 GB Mac (~2,300 expert slots in RAM): Between resident CLOCK hits and 2-layer lookahead adoption, ~55–65% of expert requests are already in RAM or pre-streamed. That leaves only ~190–240 MiB of net SSD wait per token (~70–90 ms at 2.5–3.0 GB/s NVMe) plus ~50 ms of overlapped GPU compute = 5 to 8 tokens/sec .
On a 24 GB+ Mac (~16+ GB expert pool): Over 60% of the entire 80B model stays resident in RAM (85–90%+ hit rate), dropping SSD wait to ~10–20 ms/token and running at 15 to 19 tokens/sec .
6. Built-in server, instant state rollback & memory
3-level conversation & persona state reuse ( mlx_server.py ): Takes zero-copy snapshots of the immutable 74 MiB Gated DeltaNet recurrent arrays at prompt boundaries and system-prompt anchors ( cache_anchor ), so multi-turn chats and persona switches resume instantly without re-reading the prompt.
Adaptive MTP speculation with zero-compute rollback ( speculative.py ): Records per-position Gated DeltaNet states during verification so rejected draft tokens roll back in 0 ms , while an adaptive wall-clock governor automatically switches drafting off on 8 GB pools when SSD misses outweigh draft gains.
Markdown + SQLite long-term memory: Yanah's cross-conversation memory lives in human-readable Markdown files indexed by local SQLite for zero-lag retrieval, separate from the model's KV cache.
*Models tested were a 64gb MacBook Pro M1 and an M3 MacBook Pro M3. Those are confirmed. Most other stats and claims were made by AI Calculating the math. I have made every effort to ensure the information is rock solid.
A 177B model runs faster on a MacBook Pro than a dense 27B model does.
That's the whole story in one number. On a 64GB Apple Silicon MacBook Pro, a 177B MoE model generates about 20 tokens/sec, faster than a dense 27B model on the same machine, because only a small slice of the big model works on each word.
* Performance figures are estimates based on internal testing and published benchmarks. Results may vary by Mac model, memory, storage, model choice and conversation length.
Why a bigger MoE can beat a smaller dense model
Speed tracks the parameters computed per token, not the total size.
A dense 27B computes all 27B parameters on every token, which saturates the memory bus.
An MoE model wakes only its a

[truncated]
