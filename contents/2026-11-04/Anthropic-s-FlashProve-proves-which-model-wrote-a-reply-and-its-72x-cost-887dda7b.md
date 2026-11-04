---
source: "https://theainewsreport.com/2026-11-04-anthropic-flashprove-provable-inference-explained.html"
hn_url: "https://news.ycombinator.com/item?id=50371853"
title: "Anthropic's FlashProve proves which model wrote a reply, and its 72x cost"
article_title: "Anthropic FlashProve explained: how a zero-knowledge proof…"
image: "https://theainewsreport.com/og-2026-11-04-anthropic-flashprove-provable-inference-explained.png"
author: "titusblair"
captured_at: "2026-11-04T22:27:47Z"
capture_tool: "hn-digest"
hn_id: 50371853
score: 1
comments: 1
posted_at: "2026-11-04T21:31:36Z"
tags:
  - hacker-news
---

# Anthropic's FlashProve proves which model wrote a reply, and its 72x cost

- HN: [50371853](https://news.ycombinator.com/item?id=50371853)
- Source: [theainewsreport.com](https://theainewsreport.com/2026-11-04-anthropic-flashprove-provable-inference-explained.html)
- Score: 1
- Comments: 1
- Posted: 2026-11-04T21:31:36Z

## Translation

Title: Anthropic's FlashProve proves which model wrote a reply, and its 72x cost
Article title: Anthropic FlashProve explained: how a zero-knowledge proof…
Description: Anthropic released an open-source prototype that lets anyone check that a reply came from a specific, committed model without seeing its weights. We walk the three steps in plain words, turn the benchmark into per-request cost and

Article text:
The Morning Wire AI NEWS REPORT
Anthropic FlashProve explained: how a zero-knowledge proof shows which model wrote a reply, what it costs at about 72x normal serving, and what it does not prove
Anthropic released an open-source prototype that lets anyone check that a reply came from a specific, committed model without seeing its weights. We walk the three steps in plain words, turn the benchmark into per-request cost and size, and list what the proof leaves out.
Provable inference means a user with no access to the weights can check that a reply came from one specific model the provider committed to.
FlashProve uses a zero-knowledge proof: the proof shows the math was done right without revealing the weights.
On one H100 it proves about 110 tokens a second, around 72x the GPU time of normal vLLM serving. Each proof is about 19 MB, and a CPU checks it in about 17 seconds.
It covers greedy decoding on an exact integer version of the weights, up to 1,024 tokens of context. Sampled outputs and side channels are out of scope.
It is a prototype under Apache 2.0, with code, a spec and Lean proofs of the top-level security claims.
When you pay for a model through an API, you trust that the provider ran the model it named. You cannot check. A provider could quietly swap in a cheaper model, or a modified one.
Provable inference turns that trust into a check. Anthropic's post lists what customers or regulators could confirm: that a reply came from the committed model on the given input, that the served model is the one they paid for, and that a model was not silently swapped.
Step 1, commit (once per model): the provider publishes a short fingerprint of the weights, called a commitment. The weights stay secret. Anthropic measured about 39 minutes for this step, and one commitment covers about 16 million requests.
Step 2, prove (per batch of requests): after answering, the provider builds a proof that the committed model turns this prompt into these output tokens. The proof is a zkSNARK, a short zero-knowledge argument. It is built from an interactive proof turned into a non-interactive one with hash functions, which Anthropic says also gives post-quantum security.
Step 3, verify (by anyone): with the commitment, the prompt, the reply and the proof, a checker confirms it on an ordinary CPU, no GPU and no weights needed.
Zero-knowledge means the proof reveals nothing about the weights or the middle steps. Soundness means a forged proof is rejected. Succinct means the proof is small and fast to check compared with re-running the model.
Why a mixture-of-experts model makes it feasible
Proving every multiplication of a 400 billion parameter model would be far too slow. FlashProve's trick is that a mixture-of-experts model only uses a few experts per token, so the prover's work scales with the active experts, not the whole model.
Anthropic tested it on Llama 4 Maverick: 400B total parameters, about 17B active per token, 4-bit weights. The prover streams weight segments from host memory, so it runs on a single GPU.
Anthropic's numbers, then the same numbers per request:
Anthropic's measurements (Llama 4 Maverick, one H100)
batch of 24 requests proved in ~223 s
prover throughput ~110 tokens/s
overhead vs vLLM serving ~72x GPU time (66-77x; ~94x at vLLM's top load)
proof size ~18.8 MB
verifier ~17 s on 24 CPU threads, ~21 s on one
commitment download ~6.5 MB
Per 1,024-token request (our arithmetic)
H100 time to prove 223 s / 24 = ~9.3 s
data to ship ~19 MB proof So proving every request would multiply GPU cost by roughly two orders of magnitude, which Anthropic itself says. Proofs are made in batches after the fact, so the user does not wait for them. A buyer could ask for proofs on a random sample of requests rather than all of them.
Anthropic lists the limits plainly, and they matter:
Greedy decoding only, on an exact integer version of the weights. Real serving often samples and uses floating point, so outputs can drift from the proved version at some position.
Context is limited to 1,024 tokens in this prototype.
Adversarially designed prompts are not covered.
The architecture is not hidden, and the design assumes a mixture-of-experts model.
Side channels, like how long a proof takes, are out of scope.
Only the Lean protocol is formally verified. The fast C++ code is not proven to match the paper or the Lean version.
And one scoping point from the post: proving that hardware was busy serving does not rule out that it was also used for training on the side.
Anthropic says earlier open-source systems are orders of magnitude slower. It also cites Attestable at 53 to 77 tokens a second on a dense 31B model, but Attestable has not released open-source code, so that comparison cannot be checked from the outside.
If you buy models through a reseller or gateway, ask how they prove which model ran.
Read the protocol spec and limits before relying on any proof claim.
Researchers: the prover, verifier and Lean proofs are open under Apache 2.0.
When, or whether, any provider will offer proofs in production. The post describes a prototype.
How the cost changes for longer contexts than 1,024 tokens.
Whether the fast C++ code matches the formally verified Lean protocol. Anthropic says it is not guaranteed.
The Attestable comparison rests on numbers Anthropic cites; Attestable's code is not public.
Anthropic, "Progress on provable inference for large language models" (November 2, 2026) — the original report
Hacker News discussion of the post (item 50368626)

## Original Extract

Anthropic released an open-source prototype that lets anyone check that a reply came from a specific, committed model without seeing its weights. We walk the three steps in plain words, turn the benchmark into per-request cost and

The Morning Wire AI NEWS REPORT
Anthropic FlashProve explained: how a zero-knowledge proof shows which model wrote a reply, what it costs at about 72x normal serving, and what it does not prove
Anthropic released an open-source prototype that lets anyone check that a reply came from a specific, committed model without seeing its weights. We walk the three steps in plain words, turn the benchmark into per-request cost and size, and list what the proof leaves out.
Provable inference means a user with no access to the weights can check that a reply came from one specific model the provider committed to.
FlashProve uses a zero-knowledge proof: the proof shows the math was done right without revealing the weights.
On one H100 it proves about 110 tokens a second, around 72x the GPU time of normal vLLM serving. Each proof is about 19 MB, and a CPU checks it in about 17 seconds.
It covers greedy decoding on an exact integer version of the weights, up to 1,024 tokens of context. Sampled outputs and side channels are out of scope.
It is a prototype under Apache 2.0, with code, a spec and Lean proofs of the top-level security claims.
When you pay for a model through an API, you trust that the provider ran the model it named. You cannot check. A provider could quietly swap in a cheaper model, or a modified one.
Provable inference turns that trust into a check. Anthropic's post lists what customers or regulators could confirm: that a reply came from the committed model on the given input, that the served model is the one they paid for, and that a model was not silently swapped.
Step 1, commit (once per model): the provider publishes a short fingerprint of the weights, called a commitment. The weights stay secret. Anthropic measured about 39 minutes for this step, and one commitment covers about 16 million requests.
Step 2, prove (per batch of requests): after answering, the provider builds a proof that the committed model turns this prompt into these output tokens. The proof is a zkSNARK, a short zero-knowledge argument. It is built from an interactive proof turned into a non-interactive one with hash functions, which Anthropic says also gives post-quantum security.
Step 3, verify (by anyone): with the commitment, the prompt, the reply and the proof, a checker confirms it on an ordinary CPU, no GPU and no weights needed.
Zero-knowledge means the proof reveals nothing about the weights or the middle steps. Soundness means a forged proof is rejected. Succinct means the proof is small and fast to check compared with re-running the model.
Why a mixture-of-experts model makes it feasible
Proving every multiplication of a 400 billion parameter model would be far too slow. FlashProve's trick is that a mixture-of-experts model only uses a few experts per token, so the prover's work scales with the active experts, not the whole model.
Anthropic tested it on Llama 4 Maverick: 400B total parameters, about 17B active per token, 4-bit weights. The prover streams weight segments from host memory, so it runs on a single GPU.
Anthropic's numbers, then the same numbers per request:
Anthropic's measurements (Llama 4 Maverick, one H100)
batch of 24 requests proved in ~223 s
prover throughput ~110 tokens/s
overhead vs vLLM serving ~72x GPU time (66-77x; ~94x at vLLM's top load)
proof size ~18.8 MB
verifier ~17 s on 24 CPU threads, ~21 s on one
commitment download ~6.5 MB
Per 1,024-token request (our arithmetic)
H100 time to prove 223 s / 24 = ~9.3 s
data to ship ~19 MB proof So proving every request would multiply GPU cost by roughly two orders of magnitude, which Anthropic itself says. Proofs are made in batches after the fact, so the user does not wait for them. A buyer could ask for proofs on a random sample of requests rather than all of them.
Anthropic lists the limits plainly, and they matter:
Greedy decoding only, on an exact integer version of the weights. Real serving often samples and uses floating point, so outputs can drift from the proved version at some position.
Context is limited to 1,024 tokens in this prototype.
Adversarially designed prompts are not covered.
The architecture is not hidden, and the design assumes a mixture-of-experts model.
Side channels, like how long a proof takes, are out of scope.
Only the Lean protocol is formally verified. The fast C++ code is not proven to match the paper or the Lean version.
And one scoping point from the post: proving that hardware was busy serving does not rule out that it was also used for training on the side.
Anthropic says earlier open-source systems are orders of magnitude slower. It also cites Attestable at 53 to 77 tokens a second on a dense 31B model, but Attestable has not released open-source code, so that comparison cannot be checked from the outside.
If you buy models through a reseller or gateway, ask how they prove which model ran.
Read the protocol spec and limits before relying on any proof claim.
Researchers: the prover, verifier and Lean proofs are open under Apache 2.0.
When, or whether, any provider will offer proofs in production. The post describes a prototype.
How the cost changes for longer contexts than 1,024 tokens.
Whether the fast C++ code matches the formally verified Lean protocol. Anthropic says it is not guaranteed.
The Attestable comparison rests on numbers Anthropic cites; Attestable's code is not public.
Anthropic, "Progress on provable inference for large language models" (November 2, 2026) — the original report
Hacker News discussion of the post (item 50368626)
