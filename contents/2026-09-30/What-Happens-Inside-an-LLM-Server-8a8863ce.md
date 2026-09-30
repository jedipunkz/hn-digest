---
source: "https://muhammadraza.me/2026/what-happens-inside-an-llm-server/"
hn_url: "https://news.ycombinator.com/item?id=49910406"
title: "What Happens Inside an LLM Server"
article_title: "What Happens Inside an LLM Server When 10 People Send a Prompt"
image: "https://muhammadraza.me/og/base_hu_1612ebdcd5a03ed8.png"
author: "mr_o47"
captured_at: "2026-09-30T15:41:39Z"
capture_tool: "hn-digest"
hn_id: 49910406
score: 2
comments: 0
posted_at: "2026-09-30T15:39:09Z"
tags:
  - hacker-news
---

# What Happens Inside an LLM Server

- HN: [49910406](https://news.ycombinator.com/item?id=49910406)
- Source: [muhammadraza.me](https://muhammadraza.me/2026/what-happens-inside-an-llm-server/)
- Score: 2
- Comments: 0
- Posted: 2026-09-30T15:39:09Z

## Translation

Title: What Happens Inside an LLM Server
Article title: What Happens Inside an LLM Server When 10 People Send a Prompt
Description: What an LLM server does when ten people send a prompt at once: prefill, decode, the KV cache, continuous batching, PagedAttention, and more, with animations.

Article text:
Skip to content Muhammad Raza Writing
What Happens Inside an LLM Server When 10 People Send a Prompt
What an LLM server does when ten people send a prompt at once: prefill, decode, the KV cache, continuous batching, PagedAttention, and more, with animations.
I have been running models on my own machine for a while. Then I gave a few friends the URL, and a question came up that I could not answer: what does the server do when all of them press send at once?
My guess was simple. One person gets about 60 tokens per second. Ten people split that, so each one gets 6. That guess is wrong. These are my notes on why.
Short answer: A GPU running an LLM spends most of its time reading the model from memory, not doing math. Once the weights are read for a step, making a token for ten people costs little more than making one. The real limit is the memory each conversation needs, called the KV cache. Most of the tricks inside an LLM server exist to fit more conversations into that memory.
The animations run on a simple model: Llama 3.1 8B at 16-bit, on a GPU with about 1 TB/s of memory bandwidth and 165 TFLOPS, close to an RTX 4090. Real servers get less. The shape holds.
A prompt is two different jobs
The server does two things with a message. First it reads the whole prompt in one pass. That is prefill . Then it writes the reply one token at a time, one pass per token. That is decode .
I assumed prefill was the heavy part. Press send on the short question and watch the meters.
Every pass reads all 16 GB of weights. For a short question the math is tiny, so the GPU spends the pass waiting on memory. The pasted doc flips this: 2,000 tokens is enough math to keep the GPU busy. Per token, decode is the expensive job: the Sarathi paper measured it at about 200 times the cost of a prefill token.
It is the same memory limit I ran into in GGUF vs MLX . I had not connected it to serving.
Every friend brings their own memory
So that it never redoes old work, the model keeps a key and a value for every token it has read, in every layer. That is the KV cache . For Llama 3.1 8B it is 128 KiB per token, so an 8K-token chat takes about 1 GB.
Everyone shares the weights. Nobody shares a KV cache. Move the sliders.
On a 24 GB card with 16-bit weights, four friends at 8K fit and ten do not. This was my first real surprise. What limits the number of people is not speed. It is this memory.
Each decode pass reads all the weights anyway, so the server can make a token for ten people in the same pass. The weights come in once and get used ten times. That is batching.
The simple way is to start a batch and wait until every reply in it is done. Short replies finish early, and their slots sit empty. Continuous batching puts a new request into a free slot on the next step. The Orca paper called this iteration-level scheduling.
Same GPU, same replies: 84 steps with static batching, 58 with continuous batching.
Reserved memory is wasted memory
Older servers reserved KV memory for the longest chat allowed. Most chats are short, so most of that memory sat empty. The vLLM team measured 60 to 80% of KV memory lost this way in the systems they tested.
PagedAttention takes its fix from operating systems. It hands out memory in pages of 16 tokens, only when a chat needs them, anywhere there is room.
With pages, all ten friends fit where four fit before. Everyone also starts with the same system prompt, so the server stores those pages once and shares them.
The friend who pastes a whole document
Then one friend pastes an 8K-token document. The prefill for it takes about 800 ms in this model. If the server runs it as one pass, nine other replies stop mid-sentence.
Chunked prefill cuts the document into pieces and runs each piece in the same pass as the normal decode steps.
Each step gets slower, 26 ms instead of 18 ms, but nobody freezes. The friend with the document waits about the same time. vLLM turns this on by default.
This one surprised me the most. A small draft model guesses the next few tokens. The big model checks all the guesses in one pass, keeps the ones that are right, and fixes the first wrong one. This is speculative decoding , and the paper shows that the output is the same as the big model’s output on its own.
It works because a memory-bound pass costs the same for one token or five. On a busy server that stops being true. With enough users the GPU is doing real math, the checking costs real time, and in this model the guessing makes things slower.
Back to my first guess. This is the same model, with each friend at a 2K-token chat:
Not 6 tokens per second each. 54. Each friend loses a little speed because every pass also reads their KV cache, and the server does almost nine times the work. These numbers come from the model, not a benchmark.
For llama.cpp’s llama-server :
llama-server -m Meta-Llama-3.1-8B-Instruct-Q4_K_M.gguf -c 32768 -np 4 --metrics -np sets the number of slots, one per friend at a time. The catch: when you set it, each slot gets -c divided by the number of slots. Here that is 8,192 tokens per friend, not 32,768. Continuous batching is on by default.
vllm serve meta-llama/Llama-3.1-8B-Instruct --max-model-len 8192 --max-num-seqs 16 --max-model-len caps the chat length, and with it the KV cache per friend. --max-num-seqs caps how many requests run at once (default 128). vLLM takes 92% of GPU memory and gives what is left after the weights to the KV cache. Prefix caching, the shared-pages trick, is on by default.
In vLLM I would watch vllm:time_to_first_token_seconds , vllm:inter_token_latency_seconds , vllm:num_requests_waiting , and vllm:kv_cache_usage_perc . If KV cache usage sits near 100% and the waiting count climbs, the fix is more memory or shorter chats, not a faster GPU.
Decode is memory-bound, so one pass can serve many people for close to the price of one.
The KV cache decides how many people fit, not speed.
Pages, chunks, and guessing ahead are all ways to spend memory and passes better.
Orca: A Distributed Serving System for Transformer-Based Generative Models (OSDI 2022)
Efficient Memory Management for Large Language Model Serving with PagedAttention (SOSP 2023)
vLLM launch post , the source of the 60 to 80% figure
SARATHI: Efficient LLM Inference by Piggybacking Decodes with Chunked Prefills
Sarathi-Serve: Taming Throughput-Latency Tradeoff in LLM Inference (OSDI 2024)
Fast Inference from Transformers via Speculative Decoding (ICML 2023)
Mastering LLM Techniques: Inference Optimization (NVIDIA)
Bringing coding agents into your team?
I help engineering teams put coding agents to work safely: choosing the right tools, setting up guardrails like permissions, hooks, and shared memory, and wiring agents into code review and CI/CD.
“Muhammad is an extraordinarily diligent and driven problem solver who proactively discovers opportunities to deliver value.” Client review
Deep dives on DevOps, AWS, and AI agents. Unsubscribe anytime.
Independent DevOps consultant and former AWS Professional Services engineer. I help teams run reliable AWS infrastructure, spend less on it, and ship faster.
Found this useful? Share it on
X or
Hacker News .
ai Jun 26
GGUF vs MLX on Mac (Apple Silicon): Which Format to Pick
ai Apr 26
Building CodeWiki: Compiling Codebases Into Living Wikis With LLMs
ai Sep 26
How to Convert GGUF to MLX (and When You Should Not)
From “it just works” to “I know why it works.” © 2026

## Original Extract

What an LLM server does when ten people send a prompt at once: prefill, decode, the KV cache, continuous batching, PagedAttention, and more, with animations.

Skip to content Muhammad Raza Writing
What Happens Inside an LLM Server When 10 People Send a Prompt
What an LLM server does when ten people send a prompt at once: prefill, decode, the KV cache, continuous batching, PagedAttention, and more, with animations.
I have been running models on my own machine for a while. Then I gave a few friends the URL, and a question came up that I could not answer: what does the server do when all of them press send at once?
My guess was simple. One person gets about 60 tokens per second. Ten people split that, so each one gets 6. That guess is wrong. These are my notes on why.
Short answer: A GPU running an LLM spends most of its time reading the model from memory, not doing math. Once the weights are read for a step, making a token for ten people costs little more than making one. The real limit is the memory each conversation needs, called the KV cache. Most of the tricks inside an LLM server exist to fit more conversations into that memory.
The animations run on a simple model: Llama 3.1 8B at 16-bit, on a GPU with about 1 TB/s of memory bandwidth and 165 TFLOPS, close to an RTX 4090. Real servers get less. The shape holds.
A prompt is two different jobs
The server does two things with a message. First it reads the whole prompt in one pass. That is prefill . Then it writes the reply one token at a time, one pass per token. That is decode .
I assumed prefill was the heavy part. Press send on the short question and watch the meters.
Every pass reads all 16 GB of weights. For a short question the math is tiny, so the GPU spends the pass waiting on memory. The pasted doc flips this: 2,000 tokens is enough math to keep the GPU busy. Per token, decode is the expensive job: the Sarathi paper measured it at about 200 times the cost of a prefill token.
It is the same memory limit I ran into in GGUF vs MLX . I had not connected it to serving.
Every friend brings their own memory
So that it never redoes old work, the model keeps a key and a value for every token it has read, in every layer. That is the KV cache . For Llama 3.1 8B it is 128 KiB per token, so an 8K-token chat takes about 1 GB.
Everyone shares the weights. Nobody shares a KV cache. Move the sliders.
On a 24 GB card with 16-bit weights, four friends at 8K fit and ten do not. This was my first real surprise. What limits the number of people is not speed. It is this memory.
Each decode pass reads all the weights anyway, so the server can make a token for ten people in the same pass. The weights come in once and get used ten times. That is batching.
The simple way is to start a batch and wait until every reply in it is done. Short replies finish early, and their slots sit empty. Continuous batching puts a new request into a free slot on the next step. The Orca paper called this iteration-level scheduling.
Same GPU, same replies: 84 steps with static batching, 58 with continuous batching.
Reserved memory is wasted memory
Older servers reserved KV memory for the longest chat allowed. Most chats are short, so most of that memory sat empty. The vLLM team measured 60 to 80% of KV memory lost this way in the systems they tested.
PagedAttention takes its fix from operating systems. It hands out memory in pages of 16 tokens, only when a chat needs them, anywhere there is room.
With pages, all ten friends fit where four fit before. Everyone also starts with the same system prompt, so the server stores those pages once and shares them.
The friend who pastes a whole document
Then one friend pastes an 8K-token document. The prefill for it takes about 800 ms in this model. If the server runs it as one pass, nine other replies stop mid-sentence.
Chunked prefill cuts the document into pieces and runs each piece in the same pass as the normal decode steps.
Each step gets slower, 26 ms instead of 18 ms, but nobody freezes. The friend with the document waits about the same time. vLLM turns this on by default.
This one surprised me the most. A small draft model guesses the next few tokens. The big model checks all the guesses in one pass, keeps the ones that are right, and fixes the first wrong one. This is speculative decoding , and the paper shows that the output is the same as the big model’s output on its own.
It works because a memory-bound pass costs the same for one token or five. On a busy server that stops being true. With enough users the GPU is doing real math, the checking costs real time, and in this model the guessing makes things slower.
Back to my first guess. This is the same model, with each friend at a 2K-token chat:
Not 6 tokens per second each. 54. Each friend loses a little speed because every pass also reads their KV cache, and the server does almost nine times the work. These numbers come from the model, not a benchmark.
For llama.cpp’s llama-server :
llama-server -m Meta-Llama-3.1-8B-Instruct-Q4_K_M.gguf -c 32768 -np 4 --metrics -np sets the number of slots, one per friend at a time. The catch: when you set it, each slot gets -c divided by the number of slots. Here that is 8,192 tokens per friend, not 32,768. Continuous batching is on by default.
vllm serve meta-llama/Llama-3.1-8B-Instruct --max-model-len 8192 --max-num-seqs 16 --max-model-len caps the chat length, and with it the KV cache per friend. --max-num-seqs caps how many requests run at once (default 128). vLLM takes 92% of GPU memory and gives what is left after the weights to the KV cache. Prefix caching, the shared-pages trick, is on by default.
In vLLM I would watch vllm:time_to_first_token_seconds , vllm:inter_token_latency_seconds , vllm:num_requests_waiting , and vllm:kv_cache_usage_perc . If KV cache usage sits near 100% and the waiting count climbs, the fix is more memory or shorter chats, not a faster GPU.
Decode is memory-bound, so one pass can serve many people for close to the price of one.
The KV cache decides how many people fit, not speed.
Pages, chunks, and guessing ahead are all ways to spend memory and passes better.
Orca: A Distributed Serving System for Transformer-Based Generative Models (OSDI 2022)
Efficient Memory Management for Large Language Model Serving with PagedAttention (SOSP 2023)
vLLM launch post , the source of the 60 to 80% figure
SARATHI: Efficient LLM Inference by Piggybacking Decodes with Chunked Prefills
Sarathi-Serve: Taming Throughput-Latency Tradeoff in LLM Inference (OSDI 2024)
Fast Inference from Transformers via Speculative Decoding (ICML 2023)
Mastering LLM Techniques: Inference Optimization (NVIDIA)
Bringing coding agents into your team?
I help engineering teams put coding agents to work safely: choosing the right tools, setting up guardrails like permissions, hooks, and shared memory, and wiring agents into code review and CI/CD.
“Muhammad is an extraordinarily diligent and driven problem solver who proactively discovers opportunities to deliver value.” Client review
Deep dives on DevOps, AWS, and AI agents. Unsubscribe anytime.
Independent DevOps consultant and former AWS Professional Services engineer. I help teams run reliable AWS infrastructure, spend less on it, and ship faster.
Found this useful? Share it on
X or
Hacker News .
ai Jun 26
GGUF vs MLX on Mac (Apple Silicon): Which Format to Pick
ai Apr 26
Building CodeWiki: Compiling Codebases Into Living Wikis With LLMs
ai Sep 26
How to Convert GGUF to MLX (and When You Should Not)
From “it just works” to “I know why it works.” © 2026
