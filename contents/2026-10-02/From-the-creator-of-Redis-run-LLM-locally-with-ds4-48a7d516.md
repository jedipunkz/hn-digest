---
source: "https://dwarfstar.sh/"
hn_url: "https://news.ycombinator.com/item?id=49936575"
title: "From the creator of Redis; run LLM locally with ds4"
article_title: "DwarfStar 4 (ds4): Local DeepSeek V4.1, Qwen and GLM"
image: "https://dwarfstar.sh/og-preview-v2.png"
author: "fibo"
captured_at: "2026-10-02T19:01:17Z"
capture_tool: "hn-digest"
hn_id: 49936575
score: 4
comments: 0
posted_at: "2026-10-02T18:01:16Z"
tags:
  - hacker-news
---

# From the creator of Redis; run LLM locally with ds4

- HN: [49936575](https://news.ycombinator.com/item?id=49936575)
- Source: [dwarfstar.sh](https://dwarfstar.sh/)
- Score: 4
- Comments: 0
- Posted: 2026-10-02T18:01:16Z

## Translation

Title: From the creator of Redis; run LLM locally with ds4
Article title: DwarfStar 4 (ds4): Local DeepSeek V4.1, Qwen and GLM
Description: Docs, benchmarks and setup notes for ds4, antirez

Article text:
Skip to content dwarfstar .sh Docs Hardware Benchmarks Blog About ⌘K 19k Get started Docs Hardware Benchmarks Blog About GitHub ↗ Get started → SEARCH · DOCS, BLOG, PAGES
0 HITS · NOTHING IN THE DRAWING SET
DS4 · LOCAL FRONTIER INFERENCE
Run frontier open weights locally with ds4.
DwarfStar 4 is a narrow C inference engine for high-memory Mac,
CUDA and ROCm machines. It supports DeepSeek V4 and V4.1 Flash,
GLM 5.x and Qwen3.8 Flash Next, with text and vision models, local
APIs, a CLI and a native agent in one stack.
SUPPORTED: DEEPSEEK V4 / V4.1 + GLM 5.x + QWEN3.8 · MIT LICENSE
· C / METAL / CUDA / ROCM
· QWEN ON 64GB
DeepSeek V4 Flash is a large mixture-of-experts model. The usual
path is remote serving; ds4 starts from the opposite constraint.
Asymmetric quantization targets the routed experts while preserving
critical paths. The model becomes practical on high-memory machines.
The local engine exposes a CLI, HTTP APIs and a native agent, all
sharing the same model state and cache.
Local frontier inference, narrow on purpose
Not a generic GGUF runner. ds4 follows a small, opportunistic set of model families and validates each supported layout end to end.
Compress the routed experts, keep critical shared paths precise.
That is how the supported routed-MoE builds fit their target machines.
Save long prefixes to SSD and resume by prompt hash. Restarts do
not have to mean full re-prefill.
Use ./ds4 for chat, ./ds4-server for local
APIs and ./ds4-agent for persistent coding sessions.
How the ds4 stack fits together
Project GGUFs, a self-contained engine and agent-facing interfaces, checked against official model outputs.
MODEL DeepSeek V4 / V4.1 GLM 5.x · Qwen3.8 supported GGUF layouts only asymmetric 2-bit + imatrix load ENGINE ds4 engine written in C metal · cuda · rocm parallel · batch · speculate KV cache RAM ⇄ SSD · survives restarts serve ./ds4 interactive CLI ./ds4-server OpenAI + Anthropic API ./ds4-agent native coding agent personal → distributed memory classes RUNTIME MAP · SIMPLIFIED. SEE
ARCHITECTURE NOTES FOR THE FULL DRAWING.
Download the project GGUF, build for your backend, then start the CLI or server. Generic GGUF files are not the target.
STEP 2 · BUILD FOR YOUR BACKEND
Apple Silicon Mac 64 GB+ depending on model · Metal
NVIDIA DGX Spark or a generic CUDA Linux box
AMD Strix Halo Framework Desktop & similar · ROCm
Mac Studio 512 GB V4.1 Q4 or PRO-class headroom
ds4 hardware fit: local, streamed and distributed
Pick your platform and memory: get a conservative starting path and understand which execution modes apply.
V4 Flash Q2 is the baseline. At 128 GB, GLM 5.3 Q2 and Qwen Q4 also fit; V4.1 Q2 streams from SSD.
./download_model.sh ds4f-q2 && make REF · M5 MAX 128GB · 32K CTX: 34.4 T/S GEN · 557 T/S PREFILL
Estimates from the ds4 benchmark table .
Full guide in Hardware and
Installation .
ds4 benchmarks: prefill and generation
Reference rows from upstream. Read prefill and generation separately, especially for long-context agent workloads.
Use ds4 from Codex, Claude Code and OpenCode
ds4-server speaks OpenAI and Anthropic-style APIs, so local coding agents can connect to your own machine with a base URL.
Start with the quickstart, check the hardware matrix, then connect
your editor, agent or API client to the local server.
Community site for DwarfStar (ds4), a project by Salvatore Sanfilippo
(antirez), with docs, benchmarks and notes around the local inference
engine.

## Original Extract

Docs, benchmarks and setup notes for ds4, antirez

Skip to content dwarfstar .sh Docs Hardware Benchmarks Blog About ⌘K 19k Get started Docs Hardware Benchmarks Blog About GitHub ↗ Get started → SEARCH · DOCS, BLOG, PAGES
0 HITS · NOTHING IN THE DRAWING SET
DS4 · LOCAL FRONTIER INFERENCE
Run frontier open weights locally with ds4.
DwarfStar 4 is a narrow C inference engine for high-memory Mac,
CUDA and ROCm machines. It supports DeepSeek V4 and V4.1 Flash,
GLM 5.x and Qwen3.8 Flash Next, with text and vision models, local
APIs, a CLI and a native agent in one stack.
SUPPORTED: DEEPSEEK V4 / V4.1 + GLM 5.x + QWEN3.8 · MIT LICENSE
· C / METAL / CUDA / ROCM
· QWEN ON 64GB
DeepSeek V4 Flash is a large mixture-of-experts model. The usual
path is remote serving; ds4 starts from the opposite constraint.
Asymmetric quantization targets the routed experts while preserving
critical paths. The model becomes practical on high-memory machines.
The local engine exposes a CLI, HTTP APIs and a native agent, all
sharing the same model state and cache.
Local frontier inference, narrow on purpose
Not a generic GGUF runner. ds4 follows a small, opportunistic set of model families and validates each supported layout end to end.
Compress the routed experts, keep critical shared paths precise.
That is how the supported routed-MoE builds fit their target machines.
Save long prefixes to SSD and resume by prompt hash. Restarts do
not have to mean full re-prefill.
Use ./ds4 for chat, ./ds4-server for local
APIs and ./ds4-agent for persistent coding sessions.
How the ds4 stack fits together
Project GGUFs, a self-contained engine and agent-facing interfaces, checked against official model outputs.
MODEL DeepSeek V4 / V4.1 GLM 5.x · Qwen3.8 supported GGUF layouts only asymmetric 2-bit + imatrix load ENGINE ds4 engine written in C metal · cuda · rocm parallel · batch · speculate KV cache RAM ⇄ SSD · survives restarts serve ./ds4 interactive CLI ./ds4-server OpenAI + Anthropic API ./ds4-agent native coding agent personal → distributed memory classes RUNTIME MAP · SIMPLIFIED. SEE
ARCHITECTURE NOTES FOR THE FULL DRAWING.
Download the project GGUF, build for your backend, then start the CLI or server. Generic GGUF files are not the target.
STEP 2 · BUILD FOR YOUR BACKEND
Apple Silicon Mac 64 GB+ depending on model · Metal
NVIDIA DGX Spark or a generic CUDA Linux box
AMD Strix Halo Framework Desktop & similar · ROCm
Mac Studio 512 GB V4.1 Q4 or PRO-class headroom
ds4 hardware fit: local, streamed and distributed
Pick your platform and memory: get a conservative starting path and understand which execution modes apply.
V4 Flash Q2 is the baseline. At 128 GB, GLM 5.3 Q2 and Qwen Q4 also fit; V4.1 Q2 streams from SSD.
./download_model.sh ds4f-q2 && make REF · M5 MAX 128GB · 32K CTX: 34.4 T/S GEN · 557 T/S PREFILL
Estimates from the ds4 benchmark table .
Full guide in Hardware and
Installation .
ds4 benchmarks: prefill and generation
Reference rows from upstream. Read prefill and generation separately, especially for long-context agent workloads.
Use ds4 from Codex, Claude Code and OpenCode
ds4-server speaks OpenAI and Anthropic-style APIs, so local coding agents can connect to your own machine with a base URL.
Start with the quickstart, check the hardware matrix, then connect
your editor, agent or API client to the local server.
Community site for DwarfStar (ds4), a project by Salvatore Sanfilippo
(antirez), with docs, benchmarks and notes around the local inference
engine.
