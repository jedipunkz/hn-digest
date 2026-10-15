---
source: "https://gimletlabs.ai/blog/inside-agentic-ai-workloads"
hn_url: "https://news.ycombinator.com/item?id=50107596"
title: "Inside Agentic AI Workloads"
article_title: "Inside Agentic AI Workloads | Gimlet Blog"
image: "https://gimletlabs.ai/blog/images/inside-agentic-ai-workloads/preview_dark_mode.png"
author: "matt_d"
captured_at: "2026-10-15T18:02:18Z"
capture_tool: "hn-digest"
hn_id: 50107596
score: 1
comments: 0
posted_at: "2026-10-15T18:01:20Z"
tags:
  - hacker-news
---

# Inside Agentic AI Workloads

- HN: [50107596](https://news.ycombinator.com/item?id=50107596)
- Source: [gimletlabs.ai](https://gimletlabs.ai/blog/inside-agentic-ai-workloads)
- Score: 1
- Comments: 0
- Posted: 2026-10-15T18:01:20Z

## Translation

Title: Inside Agentic AI Workloads
Article title: Inside Agentic AI Workloads | Gimlet Blog
Description: Agentic workloads repeatedly re-enter model inference over growing state. We characterize how trajectories, KV-cache retention, and tool use shape serving demand.

Article text:
Inside Agentic AI Workloads | Gimlet Blog Bringing Cerebras to Gimlet Cloud Explore
About Join us Blog Contact Home About Join us Blog Contact Inside Agentic AI Workloads
Name Ankita Nayak LinkedIn profile of Ankita Nayak Linkedin
Name Nishil Talati LinkedIn profile of Nishil Talati Linkedin
Agent workloads are more than just long prompts
KV-cache retention shapes serving performance
Tool use evolves during agent execution
How a trajectory unfolds determines total serving demand
Takeaway: Efficient agent serving requires optimizing the whole trajectory
Agent workloads are more than just long prompts
KV-cache retention shapes serving performance
Tool use evolves during agent execution
How a trajectory unfolds determines total serving demand
Takeaway: Efficient agent serving requires optimizing the whole trajectory
tl;dr: In collaboration with researchers at UIUC and Intel, we evaluated ReAct-style agent workloads across complete trajectories. Serving demand depends on how each trajectory unfolds, whether cached state survives, and how tool use evolves. Serving systems must therefore manage scheduling, cached state, and resources across the complete trajectory rather than optimize individual model calls in isolation.
Traditional LLM serving is built around individual requests. A prompt comes in, the model generates a response, and the request completes. Each request is treated independently. The serving stack is designed to batch as many of these independent requests as possible, keep the GPUs busy, and optimize for metrics such as tokens per second and time to first token.
Agentic AI breaks that assumption. An agent’s execution is inherently iterative. It repeatedly invokes the model, makes tool calls, incorporates observations, and maintains an increasingly large context throughout the task. As AI agents become an important part of software development, data analysis, and research, these workflows account for a growing share of inference demand. Serving them efficiently requires understanding the complete execution trajectory rather than treating each model invocation independently.
In our new paper, Agentic AI Workload Characterization , presented at the 2026 IEEE International Symposium on Workload Characterization , Gimlet Labs collaborated with UIUC’s Yichao Yuan and Professor Nishil Talati, and Intel’s Souvik Kundu. The study examines how agents behave, accumulate state, and consume inference resources, and what these patterns imply for serving infrastructure.
Fig. 1. A ReAct-style agent loop. Each turn generates reasoning, messages, and optional tool calls. Tool results become part of the next turn’s context, causing the agent’s state to grow throughout execution.
We traced and analyzed ReAct-style agents running through Claude Code across five workload domains. These included data analytics and data-agent tasks from ADE-Bench and DABStep, software engineering from SWE-bench Pro, terminal-based tasks from Terminal-Bench 2.0, and general-assistant tasks from GAIA. For DABStep and SWE-bench Pro, we randomly sampled 100 tasks from each benchmark. We compared Gemma4-31B and Qwen3.6-27B with thinking enabled and disabled, which we refer to as their thinking and instant configurations. The models were served on two NVIDIA H100 GPUs, and we linked each complete agent trajectory to its request-level vLLM traces. We looked at three connected layers of the execution stack: agent execution behavior and evolving context, runtime serving performance, and tool invocations external to language models. Four findings stood out.
Agent workloads are more than just long prompts
They repeatedly re-enter model inference while maintaining a growing state.
The traces reveal two ways a task becomes demanding. Some tasks accumulate large contexts as individual turns add substantial reasoning or tool output. Others require many turns, repeatedly re-entering model inference.
For example, Gemma Instant on ADE averages 109 turns per task, with a maximum of 786, yet its mean context stays at 60K tokens. Qwen Thinking on SWE-bench Pro shows a different balance, with just 41 turns on average but 79K tokens of context. Relative to the other benchmarks, SWE-bench Pro is heavy on context for every configuration, with the highest mean across all four (69K to 80K tokens). It is also heavy on the number of turns for most configurations. A workload can be demanding because it invokes the model frequently, maintains a large context, or both.
Fig. 2. Two dimensions of agent workload across benchmarks: (a) turns per task and (b) accumulated context per task. Workloads can be turn-heavy, context-heavy, or both.
Both these dimensions translate into serving cost. More turns mean more model calls to schedule and run, while larger contexts mean more KV-cache memory held across the task. Each turn also builds on earlier model responses and tool results, so the serving system is handling a stateful trajectory rather than a stream of independent requests.
KV-cache retention shapes serving performance
With retained KV state, cache reuse makes agents decode-dominated, while cache thrashing can sharply reduce throughput.
The growing contexts described above do not have to be processed from scratch at every turn. Each invocation carries forward most of the existing trajectory and appends only new reasoning, messages, tool calls, and observations. If the corresponding KV state remains available, the serving system can reuse the existing prefix and prefill only the appended tokens. Performance therefore depends not only on context size, but also on whether cached state is retained across turns.
When memory is sufficient to avoid cache eviction, empirical cache-hit rates measured by the serving system range from 84.6% to 99.5%, closely matching theoretical ratios based on trajectory prefix reuse (Fig. 3). In calculating the theoretical ratios, the initial request and newly appended tokens are treated as cache misses because they require first-time processing. Across configurations, mean raw input reaches 559.8 tokens per output token, while mean newly appended input never exceeds 7.3 tokens per output token.
Fig. 3. Most input tokens are served from the prefix cache. Cache-miss input, including newly appended content, requires prefill processing.
This reuse changes where execution time is spent. End-to-end runtime is split between LLM and tool execution (Fig. 4a), with the balance varying across benchmarks. Within LLM execution, decode accounts for 94.4% to 97% of execution time on average despite the large contexts (Fig. 4b). Prefill accounts for only 3% to 5.6% when the cached state remains available. Agent inference is therefore better characterized as repeated decode over growing cached state than as repeated processing of independent long prompts.
Fig. 4. (a) End-to-end execution time split between LLM and tool execution. (b) LLM execution time split between prefill and decode, measured at controlled concurrency to avoid cache thrashing.
That efficiency can disappear under memory pressure. A serving instance typically runs agents for many users concurrently, and their growing KV caches compete for the same limited GPU memory. In the Terminal-Bench concurrency experiment, the prefix-cache hit rate fell from 98.85% at 15 concurrent agents to 16.78% at 20 (Fig. 5a). At 25 concurrent agents, per-request generation throughput fell from 30.08 to 10.30 tokens per second (Fig. 5b), a reduction of about 66%.
Fig. 5. Effects of increasing concurrency: (a) prefix-cache hit rate, (b) per-request generation throughput, and (c) prefill share of request time. From 15 to 25 concurrent agents, throughput falls by about 66%.
When KV state no longer fits in GPU memory, the serving system must either recompute an evicted prefix or retrieve it from a secondary memory tier such as DRAM or SSD. Our experiment uses recomputation. The resulting prefill work competes with requests already generating tokens, slowing decode across the inference batch. Offloading avoids recomputation, but retrieval can still be costly because of transfer latency and limited bandwidth between memory tiers. In our measurements, prefill remains below 10% of request time (Fig. 5c), yet generation throughput falls substantially.
Together, these results explain why KV-cache retention is crucial. When the cached state survives across turns, agents reuse most of their input and spend nearly all of their LLM time in decode. When it is evicted, that advantage disappears, and the recomputation slows every request sharing the GPU, not only the agent that lost its state. For agent serving, the KV cache is long-lived execution state, and whether it stays resident directly affects throughput across repeated model invocations and tool calls.
Tool use evolves during agent execution
Changes in tool intent reveal how an agent’s work progresses.
KV-cache retention shapes the model-execution portion of an agent’s runtime, but model execution is only one part of the complete task. Between model invocations, agents call tools that vary widely in both latency and the work they perform. File operations, shell commands, web requests, and tools that launch subagents have different execution and output characteristics. A subagent can perform an entire agent execution within one parent-level call. Two workloads with similar tool-call counts can therefore have very different execution profiles depending on the tools they use.
Tool calls also reveal how an agent’s work evolves during a task. An agent fixing a complex software bug, for example, often starts much like a human developer by exploring directories and reading files before writing patches and running tests. Categorizing tool calls as read/explore or execute/write turns this intuitive pattern into a measurable runtime signal. When averaged across all five benchmarks, all four model configurations show read/explore activity declining and execute/write activity increasing as tasks progress.
Fig. 6. Tool intent over the course of agent execution, averaged across all benchmarks. All configurations shift from exploration (blue) toward action (orange).
Tool use varies across tasks and over the course of execution, so tool-call counts alone do not capture the work underway. Tracking tool intent provides an observable signal of shifts in an agent’s activity that could inform scheduling and resource allocation as the task progresses. For example, when deciding whether to keep an agent’s cached context resident, a scheduler could use its phase to anticipate which tools it may call next and how long those calls may take. Agent-serving systems can benefit from tracking tool behavior alongside model execution.
How a trajectory unfolds determines total serving demand
Serving demand accumulates across turns and must be managed at the trajectory level.
Tool activity provides one view into how an agent’s work progresses, but total serving demand depends on the complete run. Each model invocation consumes the context accumulated so far, while each model response or tool result can extend that context. Reasoning and failures illustrate how these effects accumulate across a trajectory.
Reasoning adds work to every model call. Before acting, a thinking model generates tokens that an instant model skips. Averaged across the five benchmarks, thinking accounts for 53.2% of Gemma Thinking’s output and 35.1% of Qwen Thinking’s output. Generating these tokens requires additional decode work, and retaining them increases the KV-cache state carried into subsequent turns.
That additional work, however, does not necessarily produce a larger context or more steps. As Fig. 2 shows, on SWE-bench Pro, Qwen Thinking averaged 41 turns, compared with 62 for Qwen Instant, while their mean contexts remained similar at 79K and 80K tokens. Although this relationship may vary by model and workload, the result suggests that more deliberation at each step can reduce the work r

[truncated]

## Original Extract

Agentic workloads repeatedly re-enter model inference over growing state. We characterize how trajectories, KV-cache retention, and tool use shape serving demand.

Inside Agentic AI Workloads | Gimlet Blog Bringing Cerebras to Gimlet Cloud Explore
About Join us Blog Contact Home About Join us Blog Contact Inside Agentic AI Workloads
Name Ankita Nayak LinkedIn profile of Ankita Nayak Linkedin
Name Nishil Talati LinkedIn profile of Nishil Talati Linkedin
Agent workloads are more than just long prompts
KV-cache retention shapes serving performance
Tool use evolves during agent execution
How a trajectory unfolds determines total serving demand
Takeaway: Efficient agent serving requires optimizing the whole trajectory
Agent workloads are more than just long prompts
KV-cache retention shapes serving performance
Tool use evolves during agent execution
How a trajectory unfolds determines total serving demand
Takeaway: Efficient agent serving requires optimizing the whole trajectory
tl;dr: In collaboration with researchers at UIUC and Intel, we evaluated ReAct-style agent workloads across complete trajectories. Serving demand depends on how each trajectory unfolds, whether cached state survives, and how tool use evolves. Serving systems must therefore manage scheduling, cached state, and resources across the complete trajectory rather than optimize individual model calls in isolation.
Traditional LLM serving is built around individual requests. A prompt comes in, the model generates a response, and the request completes. Each request is treated independently. The serving stack is designed to batch as many of these independent requests as possible, keep the GPUs busy, and optimize for metrics such as tokens per second and time to first token.
Agentic AI breaks that assumption. An agent’s execution is inherently iterative. It repeatedly invokes the model, makes tool calls, incorporates observations, and maintains an increasingly large context throughout the task. As AI agents become an important part of software development, data analysis, and research, these workflows account for a growing share of inference demand. Serving them efficiently requires understanding the complete execution trajectory rather than treating each model invocation independently.
In our new paper, Agentic AI Workload Characterization , presented at the 2026 IEEE International Symposium on Workload Characterization , Gimlet Labs collaborated with UIUC’s Yichao Yuan and Professor Nishil Talati, and Intel’s Souvik Kundu. The study examines how agents behave, accumulate state, and consume inference resources, and what these patterns imply for serving infrastructure.
Fig. 1. A ReAct-style agent loop. Each turn generates reasoning, messages, and optional tool calls. Tool results become part of the next turn’s context, causing the agent’s state to grow throughout execution.
We traced and analyzed ReAct-style agents running through Claude Code across five workload domains. These included data analytics and data-agent tasks from ADE-Bench and DABStep, software engineering from SWE-bench Pro, terminal-based tasks from Terminal-Bench 2.0, and general-assistant tasks from GAIA. For DABStep and SWE-bench Pro, we randomly sampled 100 tasks from each benchmark. We compared Gemma4-31B and Qwen3.6-27B with thinking enabled and disabled, which we refer to as their thinking and instant configurations. The models were served on two NVIDIA H100 GPUs, and we linked each complete agent trajectory to its request-level vLLM traces. We looked at three connected layers of the execution stack: agent execution behavior and evolving context, runtime serving performance, and tool invocations external to language models. Four findings stood out.
Agent workloads are more than just long prompts
They repeatedly re-enter model inference while maintaining a growing state.
The traces reveal two ways a task becomes demanding. Some tasks accumulate large contexts as individual turns add substantial reasoning or tool output. Others require many turns, repeatedly re-entering model inference.
For example, Gemma Instant on ADE averages 109 turns per task, with a maximum of 786, yet its mean context stays at 60K tokens. Qwen Thinking on SWE-bench Pro shows a different balance, with just 41 turns on average but 79K tokens of context. Relative to the other benchmarks, SWE-bench Pro is heavy on context for every configuration, with the highest mean across all four (69K to 80K tokens). It is also heavy on the number of turns for most configurations. A workload can be demanding because it invokes the model frequently, maintains a large context, or both.
Fig. 2. Two dimensions of agent workload across benchmarks: (a) turns per task and (b) accumulated context per task. Workloads can be turn-heavy, context-heavy, or both.
Both these dimensions translate into serving cost. More turns mean more model calls to schedule and run, while larger contexts mean more KV-cache memory held across the task. Each turn also builds on earlier model responses and tool results, so the serving system is handling a stateful trajectory rather than a stream of independent requests.
KV-cache retention shapes serving performance
With retained KV state, cache reuse makes agents decode-dominated, while cache thrashing can sharply reduce throughput.
The growing contexts described above do not have to be processed from scratch at every turn. Each invocation carries forward most of the existing trajectory and appends only new reasoning, messages, tool calls, and observations. If the corresponding KV state remains available, the serving system can reuse the existing prefix and prefill only the appended tokens. Performance therefore depends not only on context size, but also on whether cached state is retained across turns.
When memory is sufficient to avoid cache eviction, empirical cache-hit rates measured by the serving system range from 84.6% to 99.5%, closely matching theoretical ratios based on trajectory prefix reuse (Fig. 3). In calculating the theoretical ratios, the initial request and newly appended tokens are treated as cache misses because they require first-time processing. Across configurations, mean raw input reaches 559.8 tokens per output token, while mean newly appended input never exceeds 7.3 tokens per output token.
Fig. 3. Most input tokens are served from the prefix cache. Cache-miss input, including newly appended content, requires prefill processing.
This reuse changes where execution time is spent. End-to-end runtime is split between LLM and tool execution (Fig. 4a), with the balance varying across benchmarks. Within LLM execution, decode accounts for 94.4% to 97% of execution time on average despite the large contexts (Fig. 4b). Prefill accounts for only 3% to 5.6% when the cached state remains available. Agent inference is therefore better characterized as repeated decode over growing cached state than as repeated processing of independent long prompts.
Fig. 4. (a) End-to-end execution time split between LLM and tool execution. (b) LLM execution time split between prefill and decode, measured at controlled concurrency to avoid cache thrashing.
That efficiency can disappear under memory pressure. A serving instance typically runs agents for many users concurrently, and their growing KV caches compete for the same limited GPU memory. In the Terminal-Bench concurrency experiment, the prefix-cache hit rate fell from 98.85% at 15 concurrent agents to 16.78% at 20 (Fig. 5a). At 25 concurrent agents, per-request generation throughput fell from 30.08 to 10.30 tokens per second (Fig. 5b), a reduction of about 66%.
Fig. 5. Effects of increasing concurrency: (a) prefix-cache hit rate, (b) per-request generation throughput, and (c) prefill share of request time. From 15 to 25 concurrent agents, throughput falls by about 66%.
When KV state no longer fits in GPU memory, the serving system must either recompute an evicted prefix or retrieve it from a secondary memory tier such as DRAM or SSD. Our experiment uses recomputation. The resulting prefill work competes with requests already generating tokens, slowing decode across the inference batch. Offloading avoids recomputation, but retrieval can still be costly because of transfer latency and limited bandwidth between memory tiers. In our measurements, prefill remains below 10% of request time (Fig. 5c), yet generation throughput falls substantially.
Together, these results explain why KV-cache retention is crucial. When the cached state survives across turns, agents reuse most of their input and spend nearly all of their LLM time in decode. When it is evicted, that advantage disappears, and the recomputation slows every request sharing the GPU, not only the agent that lost its state. For agent serving, the KV cache is long-lived execution state, and whether it stays resident directly affects throughput across repeated model invocations and tool calls.
Tool use evolves during agent execution
Changes in tool intent reveal how an agent’s work progresses.
KV-cache retention shapes the model-execution portion of an agent’s runtime, but model execution is only one part of the complete task. Between model invocations, agents call tools that vary widely in both latency and the work they perform. File operations, shell commands, web requests, and tools that launch subagents have different execution and output characteristics. A subagent can perform an entire agent execution within one parent-level call. Two workloads with similar tool-call counts can therefore have very different execution profiles depending on the tools they use.
Tool calls also reveal how an agent’s work evolves during a task. An agent fixing a complex software bug, for example, often starts much like a human developer by exploring directories and reading files before writing patches and running tests. Categorizing tool calls as read/explore or execute/write turns this intuitive pattern into a measurable runtime signal. When averaged across all five benchmarks, all four model configurations show read/explore activity declining and execute/write activity increasing as tasks progress.
Fig. 6. Tool intent over the course of agent execution, averaged across all benchmarks. All configurations shift from exploration (blue) toward action (orange).
Tool use varies across tasks and over the course of execution, so tool-call counts alone do not capture the work underway. Tracking tool intent provides an observable signal of shifts in an agent’s activity that could inform scheduling and resource allocation as the task progresses. For example, when deciding whether to keep an agent’s cached context resident, a scheduler could use its phase to anticipate which tools it may call next and how long those calls may take. Agent-serving systems can benefit from tracking tool behavior alongside model execution.
How a trajectory unfolds determines total serving demand
Serving demand accumulates across turns and must be managed at the trajectory level.
Tool activity provides one view into how an agent’s work progresses, but total serving demand depends on the complete run. Each model invocation consumes the context accumulated so far, while each model response or tool result can extend that context. Reasoning and failures illustrate how these effects accumulate across a trajectory.
Reasoning adds work to every model call. Before acting, a thinking model generates tokens that an instant model skips. Averaged across the five benchmarks, thinking accounts for 53.2% of Gemma Thinking’s output and 35.1% of Qwen Thinking’s output. Generating these tokens requires additional decode work, and retaining them increases the KV-cache state carried into subsequent turns.
That additional work, however, does not necessarily produce a larger context or more steps. As Fig. 2 shows, on SWE-bench Pro, Qwen Thinking averaged 41 turns, compared with 62 for Qwen Instant, while their mean contexts remained similar at 79K and 80K tokens. Although this relationship may vary by model and workload, the result suggests that more deliberation at each step can reduce the work r

[truncated]
