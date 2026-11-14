---
source: "https://huggingface.co/blog/ibm-granite/granite-4-2"
hn_url: "https://news.ycombinator.com/item?id=50502172"
title: "Granite 4.2 LLMs: How They're Built"
article_title: "Granite 4.2 LLMs: How They're Built"
image: "https://cdn-uploads.huggingface.co/production/uploads/65cc2c288ebd392213e58899/XHWV2L_wZsFe3ekksCgvQ.webp"
author: "gmays"
captured_at: "2026-11-14T23:40:27Z"
capture_tool: "hn-digest"
hn_id: 50502172
score: 1
comments: 0
posted_at: "2026-11-14T23:27:36Z"
tags:
  - hacker-news
---

# Granite 4.2 LLMs: How They're Built

- HN: [50502172](https://news.ycombinator.com/item?id=50502172)
- Source: [huggingface.co](https://huggingface.co/blog/ibm-granite/granite-4-2)
- Score: 1
- Comments: 0
- Posted: 2026-11-14T23:27:36Z

## Translation

Title: Granite 4.2 LLMs: How They're Built
Description: A Blog post by IBM Granite on Hugging Face

Article text:
Granite 4.2 LLMs: How They're Built
Hugging Face Search models, datasets, users... Ctrl K Models
Back to Articles a]:hidden">
Granite 4.2 LLMs: How They're Built
SFT: Data Preparation & Quality Control Data Quality Control
Reinforcement Learning: A Multi-Stage, Multi-Environment Pipeline Training Methodology
Foundational RL: Build the Skills
Agentic RL: Learning to Act (8B / 30B)
Agentic AI Infrastructure for Scalable RL
Getting Started (Transformers) Installation
Basic Inference (Thinking Mode)
Tool Calling Basic Tool Calling
Multi-Turn Conversations History Thinking Truncation
Parsing Thinking vs. Final Answer
Using with Agentic Coding Harnesses OpenCode
A technical walkthrough of how we built the Granite 4.2 reasoning model family.
TL;DR: Granite 4.2 is our first family of dense, decoder-only reasoning LLMs, released in three sizes: 3B, 8B, and 30B . These models are post-trained from Granite-4.1 base models. Granite-4.1 base models were pre-trained from scratch on roughly 15T tokens with a five-phase strategy that extends the context window to 512K tokens, supervised fine-tuned on chain-of-thought, reasoning, and agentic-trajectory data, then post-trained with a multi-stage reinforcement learning pipeline . That pipeline includes agentic RL, where the 8B and 30B models learn to act with tools inside real sandboxed environments. Every model has a thinking / non-thinking switch, a low-effort thinking mode that spends a short reasoning budget on easy questions, and native tool calling. All Granite 4.2 models are released under the Apache 2.0 license.
Granite 4.2 is the reasoning-focused release of the Granite language-model family. Earlier Granite releases were strong instruction-following assistants; Granite 4.2 adds explicit reasoning. Every model can produce a chain of thought before its answer and can run in thinking or non-thinking mode depending on how much deliberation a task needs. A low-effort mode falls between the two, spending a short reasoning budget on easy questions.
The three sizes ( 3B, 8B, and 30B ) share the same architectural design and follow the same training pipeline (base model, SFT, then multi-stage RL), each at its own scale. All three are strong reasoners and instruction followers. The clearest capability split shows up in post-training. The 8B and 30B models additionally go through an agentic RL block that teaches them to operate as agents: calling tools, editing and running code, driving a terminal, and searching the web inside real environments. Every model supports native tool calling. Served through an OpenAI-compatible endpoint (for example, with vLLM), it emits tool calls in the OpenAI function-calling format and plugs into agentic harnesses without extra glue. Granite 4.2 is also supported in SGLang, see the SGLang cookbook for a ready-to-serve recipe.
The rest of this post walks through the build: architecture, pre-training, supervised fine-tuning, the multi-stage RL pipeline, and results.
Granite 4.2 models are built on a decoder-only dense transformer architecture with the following core components:
Attention: Grouped Query Attention (GQA) with 40 attention heads and 8 KV heads
Position Embedding: Rotary Position Embedding (RoPE) with θ = 10,000,000
Feed-Forward: MLP with SwiGLU activation
Normalization: RMSNorm (ε = 1e-5)
Embeddings: Separate input/output embeddings (not tied)
Granite 4.2 models are post-trained from Granite-4.1 base models. Granite-4.1 base models were trained from scratch on approximately 15 trillion tokens using a five-phase training strategy. Phases 1–2 focus on foundational pre-training, phases 3–4 perform mid-training with progressively higher-quality data annealing, and phase 5 introduces long-context training, extending the context window to 512K tokens . Each phase uses a distinct data mixture and learning-rate schedule, gradually shifting from broad web-scale data toward more curated, high-quality sources.
The pre-training recipe closely follows the previous generation; for a detailed treatment of the data blend, phase schedule, and long-context extension, see the Granite 4.1 blog .
SFT: Data Preparation & Quality Control
Supervised fine-tuning (SFT) turns the base model into a reliable instruction-following, reasoning, and tool-using assistant. The SFT data mixture combines agentic (31.6%) and non-agentic (68.4%) data, totaling approximately 7.2 million samples, or roughly 100B tokens, of which about 65B are trainable.
The agentic corpus covers a broad range of domains, including software engineering (SWE, 69%), tool calling (12.1%), terminal use (8.0%), math (3.5%), search (0.8%), and action (0.2%). These samples and trajectories are generated using a diverse set of agent scaffolds and harnesses, including OpenHands, OpenCode, Terminus-2, SWE-agent, OpenResearcher, MiniSWE, OpenSeeker, EnvScaler, Gemini CLI, Hermes, Codex, and Goose. The agentic data combines samples from both open-source datasets and our own synthetically generated RL environments, spanning a variety of agent–harness combinations.
The non-agentic corpus consists of several major categories: instruction following (18.8%), coding (18.8%), math (14.6%), multilingual (7.0%), science (5.4%), reasoning (3.0%), and safety (0.8%).
We apply multiple stages of quality control before a sample enters the final SFT mixture. First, data from different sources is normalized and reformatted into a consistent OpenAI Chat format, making the conversation structure and tool interactions uniform across datasets and scaffolds.
We then use GPT-OSS-120B and Gemma 4 as LLM-based judges to assess sample quality. Low-scoring samples are removed, as are samples containing hallucinated or fabricated information, invalid tool interactions, or tool calls to functions that are not defined in the corresponding tool list. Several targeted, dataset-specific heuristic rules are also applied where appropriate to further improve quality and remove known sources of noise.
Finally, we perform both local and global deduplication. Deduplication is based on SHA-256 hashes computed over the combination of the tools and messages fields, removing duplicate samples both within individual data sources and across the overall SFT mixture.
The complete corpus is first globally shuffled to reduce ordering effects and ensure that samples from different domains are well mixed during training. The shuffled corpus is then partitioned into equally sized .parquet shards, which are tokenized using the model's tokenizer and chat template and prepared for large-scale distributed training.
Before launching the final large-scale runs, we tune hyperparameters on representative configurations, sweeping learning-rate schedules, initial learning rates, and warm-up ratios to find settings that train stably across model sizes. The final training configuration is summarized below:
For the 30B model, we additionally perform a second phase of SFT focused specifically on agentic coding. In this phase, agentic, SWE, and coding data are upsampled to increase their effective contribution to the training distribution, while approximately 16% of the mixture is retained as replay data from the original SFT corpus.
The 30B model is then fine-tuned for roughly one additional epoch at a lower learning rate of 3.0e-6. This targeted second phase increases the model's exposure to agentic coding trajectories without discarding the capabilities acquired during the initial SFT phase.
Reinforcement Learning: A Multi-Stage, Multi-Environment Pipeline
After SFT, we apply a multi-stage, multi-environment reinforcement learning pipeline . Rather than a single RL pass, we run a chain of focused stages spanning many environments: math, code, science, instruction following, tool use, and structured output, then software engineering, terminal use, and web search. Each stage is an independent RL run that targets one capability and warm-starts from the previous stage's checkpoint.
Figure 1. The staged RL curriculum. Foundational RL (verifiable rewards + skill boosters) runs for all sizes; the agentic RL block (SWE → Terminal → Search) runs for 8B and 30B only. Every model finishes with RLHF. Each stage is a separate GRPO run that warm-starts from the previous checkpoint.
Every stage trains with asynchronous GRPO ( Group Relative Policy Optimization ), so the generator and trainer halves of the loop never block on each other. A pool of generation workers keeps sampling responses and dropping the finished trajectories into a shared buffer; once the buffer holds a full step's worth, the trainer pulls that batch, takes an optimizer step, and streams the updated parameters back to the generator workers without pausing them. A refresh can land partway through a rollout, leaving a single trajectory stitched together from two adjacent policy versions. We allow this instead of paying to prevent it: the workers reuse their existing KV cache rather than rebuilding it after each refresh, and the one guardrail is a limit that keeps them from drifting more than a single update behind the trainer, which bounds how off-policy any sample can get. Whatever mismatch survives that limit is handled in the objective by truncated importance sampling , which clamps the train-versus-generation log-probability ratio to a fixed ceiling so a handful of stale tokens cannot dominate an update.
Advantages are group-relative with a leave-one-out baseline : each response is judged against the mean reward of the other samples drawn for the same prompt, which removes the need for a separate value network. To make this concrete, take RLVR , the first and longest-running stage: each step pairs 256 prompts with 16 sampled responses apiece for a 4,096 -example batch, which the trainer consumes in a single optimizer step before the next rollout begins. Later stages keep this machinery unchanged and adjust only the per-stage shape, shown next.
RL training configuration
The pipeline keeps a common backbone of hyperparameters across every stage, which makes the curriculum easier to run and compare. A handful of knobs are fixed everywhere:
What changes from stage to stage is the shape of each run: how many prompts and generations per step, how long the context is, whether the agent loop runs, and how hard we pull back toward the reference policy. The table below gives the exact settings for the 30B chain, stage by stage:
Parameters shown for the 30B model. Global batch size = prompts/step × generations/prompt (e.g. 256 × 16 = 4096 for RLVR). The 3B and 8B models use the same recipe and hyperparameters with fewer stages (see How the Three Sizes Differ ); the stage list is what changes, not the knobs.
The KL schedule follows the reward type: explore freely where the reward is objective and verifiable (RLVR and SWE 2 run at KL 0), and stay close to the reference where the objective is preference, safety, or a narrow skill graft (RLHF and the code booster use KL 0.05). The rollout-turns column counts the environment interactions GRPO itself sees per rollout. In every case the model is trained on complete, real-environment trajectories.
Each stage is a separate RL run with a single objective and its own reward signal. When it finishes, its policy is exported to Hugging Face format and becomes the base model for the next stage, so the pipeline is a sequence of warm-starts:
SFT ─▶ RLVR ─▶ Skill boosters ─▶ SWE agent ─▶ Terminal ─▶ Search ─▶ RLHF
└──────── foundational RL ────────┘ └──────── agentic RL (8B / 30B) ────────┘
The 8B and 30B models follow the full ladder. The 3B model takes a shortened path: foundational RL and alignment, without the agentic block.
A stage is defined mostly by how it is rewarded . Across the pipeline there are three reward types, and a single stage can use more than one:
Verifiable rewards are objective and hard to game, so the pipeline front-loads them. Judge- and preference-based rewards handle open-ended qualities that no checker can express. Agentic

[truncated]

## Original Extract

A Blog post by IBM Granite on Hugging Face

Granite 4.2 LLMs: How They're Built
Hugging Face Search models, datasets, users... Ctrl K Models
Back to Articles a]:hidden">
Granite 4.2 LLMs: How They're Built
SFT: Data Preparation & Quality Control Data Quality Control
Reinforcement Learning: A Multi-Stage, Multi-Environment Pipeline Training Methodology
Foundational RL: Build the Skills
Agentic RL: Learning to Act (8B / 30B)
Agentic AI Infrastructure for Scalable RL
Getting Started (Transformers) Installation
Basic Inference (Thinking Mode)
Tool Calling Basic Tool Calling
Multi-Turn Conversations History Thinking Truncation
Parsing Thinking vs. Final Answer
Using with Agentic Coding Harnesses OpenCode
A technical walkthrough of how we built the Granite 4.2 reasoning model family.
TL;DR: Granite 4.2 is our first family of dense, decoder-only reasoning LLMs, released in three sizes: 3B, 8B, and 30B . These models are post-trained from Granite-4.1 base models. Granite-4.1 base models were pre-trained from scratch on roughly 15T tokens with a five-phase strategy that extends the context window to 512K tokens, supervised fine-tuned on chain-of-thought, reasoning, and agentic-trajectory data, then post-trained with a multi-stage reinforcement learning pipeline . That pipeline includes agentic RL, where the 8B and 30B models learn to act with tools inside real sandboxed environments. Every model has a thinking / non-thinking switch, a low-effort thinking mode that spends a short reasoning budget on easy questions, and native tool calling. All Granite 4.2 models are released under the Apache 2.0 license.
Granite 4.2 is the reasoning-focused release of the Granite language-model family. Earlier Granite releases were strong instruction-following assistants; Granite 4.2 adds explicit reasoning. Every model can produce a chain of thought before its answer and can run in thinking or non-thinking mode depending on how much deliberation a task needs. A low-effort mode falls between the two, spending a short reasoning budget on easy questions.
The three sizes ( 3B, 8B, and 30B ) share the same architectural design and follow the same training pipeline (base model, SFT, then multi-stage RL), each at its own scale. All three are strong reasoners and instruction followers. The clearest capability split shows up in post-training. The 8B and 30B models additionally go through an agentic RL block that teaches them to operate as agents: calling tools, editing and running code, driving a terminal, and searching the web inside real environments. Every model supports native tool calling. Served through an OpenAI-compatible endpoint (for example, with vLLM), it emits tool calls in the OpenAI function-calling format and plugs into agentic harnesses without extra glue. Granite 4.2 is also supported in SGLang, see the SGLang cookbook for a ready-to-serve recipe.
The rest of this post walks through the build: architecture, pre-training, supervised fine-tuning, the multi-stage RL pipeline, and results.
Granite 4.2 models are built on a decoder-only dense transformer architecture with the following core components:
Attention: Grouped Query Attention (GQA) with 40 attention heads and 8 KV heads
Position Embedding: Rotary Position Embedding (RoPE) with θ = 10,000,000
Feed-Forward: MLP with SwiGLU activation
Normalization: RMSNorm (ε = 1e-5)
Embeddings: Separate input/output embeddings (not tied)
Granite 4.2 models are post-trained from Granite-4.1 base models. Granite-4.1 base models were trained from scratch on approximately 15 trillion tokens using a five-phase training strategy. Phases 1–2 focus on foundational pre-training, phases 3–4 perform mid-training with progressively higher-quality data annealing, and phase 5 introduces long-context training, extending the context window to 512K tokens . Each phase uses a distinct data mixture and learning-rate schedule, gradually shifting from broad web-scale data toward more curated, high-quality sources.
The pre-training recipe closely follows the previous generation; for a detailed treatment of the data blend, phase schedule, and long-context extension, see the Granite 4.1 blog .
SFT: Data Preparation & Quality Control
Supervised fine-tuning (SFT) turns the base model into a reliable instruction-following, reasoning, and tool-using assistant. The SFT data mixture combines agentic (31.6%) and non-agentic (68.4%) data, totaling approximately 7.2 million samples, or roughly 100B tokens, of which about 65B are trainable.
The agentic corpus covers a broad range of domains, including software engineering (SWE, 69%), tool calling (12.1%), terminal use (8.0%), math (3.5%), search (0.8%), and action (0.2%). These samples and trajectories are generated using a diverse set of agent scaffolds and harnesses, including OpenHands, OpenCode, Terminus-2, SWE-agent, OpenResearcher, MiniSWE, OpenSeeker, EnvScaler, Gemini CLI, Hermes, Codex, and Goose. The agentic data combines samples from both open-source datasets and our own synthetically generated RL environments, spanning a variety of agent–harness combinations.
The non-agentic corpus consists of several major categories: instruction following (18.8%), coding (18.8%), math (14.6%), multilingual (7.0%), science (5.4%), reasoning (3.0%), and safety (0.8%).
We apply multiple stages of quality control before a sample enters the final SFT mixture. First, data from different sources is normalized and reformatted into a consistent OpenAI Chat format, making the conversation structure and tool interactions uniform across datasets and scaffolds.
We then use GPT-OSS-120B and Gemma 4 as LLM-based judges to assess sample quality. Low-scoring samples are removed, as are samples containing hallucinated or fabricated information, invalid tool interactions, or tool calls to functions that are not defined in the corresponding tool list. Several targeted, dataset-specific heuristic rules are also applied where appropriate to further improve quality and remove known sources of noise.
Finally, we perform both local and global deduplication. Deduplication is based on SHA-256 hashes computed over the combination of the tools and messages fields, removing duplicate samples both within individual data sources and across the overall SFT mixture.
The complete corpus is first globally shuffled to reduce ordering effects and ensure that samples from different domains are well mixed during training. The shuffled corpus is then partitioned into equally sized .parquet shards, which are tokenized using the model's tokenizer and chat template and prepared for large-scale distributed training.
Before launching the final large-scale runs, we tune hyperparameters on representative configurations, sweeping learning-rate schedules, initial learning rates, and warm-up ratios to find settings that train stably across model sizes. The final training configuration is summarized below:
For the 30B model, we additionally perform a second phase of SFT focused specifically on agentic coding. In this phase, agentic, SWE, and coding data are upsampled to increase their effective contribution to the training distribution, while approximately 16% of the mixture is retained as replay data from the original SFT corpus.
The 30B model is then fine-tuned for roughly one additional epoch at a lower learning rate of 3.0e-6. This targeted second phase increases the model's exposure to agentic coding trajectories without discarding the capabilities acquired during the initial SFT phase.
Reinforcement Learning: A Multi-Stage, Multi-Environment Pipeline
After SFT, we apply a multi-stage, multi-environment reinforcement learning pipeline . Rather than a single RL pass, we run a chain of focused stages spanning many environments: math, code, science, instruction following, tool use, and structured output, then software engineering, terminal use, and web search. Each stage is an independent RL run that targets one capability and warm-starts from the previous stage's checkpoint.
Figure 1. The staged RL curriculum. Foundational RL (verifiable rewards + skill boosters) runs for all sizes; the agentic RL block (SWE → Terminal → Search) runs for 8B and 30B only. Every model finishes with RLHF. Each stage is a separate GRPO run that warm-starts from the previous checkpoint.
Every stage trains with asynchronous GRPO ( Group Relative Policy Optimization ), so the generator and trainer halves of the loop never block on each other. A pool of generation workers keeps sampling responses and dropping the finished trajectories into a shared buffer; once the buffer holds a full step's worth, the trainer pulls that batch, takes an optimizer step, and streams the updated parameters back to the generator workers without pausing them. A refresh can land partway through a rollout, leaving a single trajectory stitched together from two adjacent policy versions. We allow this instead of paying to prevent it: the workers reuse their existing KV cache rather than rebuilding it after each refresh, and the one guardrail is a limit that keeps them from drifting more than a single update behind the trainer, which bounds how off-policy any sample can get. Whatever mismatch survives that limit is handled in the objective by truncated importance sampling , which clamps the train-versus-generation log-probability ratio to a fixed ceiling so a handful of stale tokens cannot dominate an update.
Advantages are group-relative with a leave-one-out baseline : each response is judged against the mean reward of the other samples drawn for the same prompt, which removes the need for a separate value network. To make this concrete, take RLVR , the first and longest-running stage: each step pairs 256 prompts with 16 sampled responses apiece for a 4,096 -example batch, which the trainer consumes in a single optimizer step before the next rollout begins. Later stages keep this machinery unchanged and adjust only the per-stage shape, shown next.
RL training configuration
The pipeline keeps a common backbone of hyperparameters across every stage, which makes the curriculum easier to run and compare. A handful of knobs are fixed everywhere:
What changes from stage to stage is the shape of each run: how many prompts and generations per step, how long the context is, whether the agent loop runs, and how hard we pull back toward the reference policy. The table below gives the exact settings for the 30B chain, stage by stage:
Parameters shown for the 30B model. Global batch size = prompts/step × generations/prompt (e.g. 256 × 16 = 4096 for RLVR). The 3B and 8B models use the same recipe and hyperparameters with fewer stages (see How the Three Sizes Differ ); the stage list is what changes, not the knobs.
The KL schedule follows the reward type: explore freely where the reward is objective and verifiable (RLVR and SWE 2 run at KL 0), and stay close to the reference where the objective is preference, safety, or a narrow skill graft (RLHF and the code booster use KL 0.05). The rollout-turns column counts the environment interactions GRPO itself sees per rollout. In every case the model is trained on complete, real-environment trajectories.
Each stage is a separate RL run with a single objective and its own reward signal. When it finishes, its policy is exported to Hugging Face format and becomes the base model for the next stage, so the pipeline is a sequence of warm-starts:
SFT ─▶ RLVR ─▶ Skill boosters ─▶ SWE agent ─▶ Terminal ─▶ Search ─▶ RLHF
└──────── foundational RL ────────┘ └──────── agentic RL (8B / 30B) ────────┘
The 8B and 30B models follow the full ladder. The 3B model takes a shortened path: foundational RL and alignment, without the agentic block.
A stage is defined mostly by how it is rewarded . Across the pipeline there are three reward types, and a single stage can use more than one:
Verifiable rewards are objective and hard to game, so the pipeline front-loads them. Judge- and preference-based rewards handle open-ended qualities that no checker can express. Agentic

[truncated]
