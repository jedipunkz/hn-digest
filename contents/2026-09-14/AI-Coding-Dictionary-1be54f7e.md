---
source: "https://www.aihero.dev/ai-coding-dictionary"
hn_url: "https://news.ycombinator.com/item?id=49690058"
title: "AI Coding Dictionary"
article_title: "AI Coding Dictionary"
image: "https://res.cloudinary.com/total-typescript/image/upload/v1777983728/ai-coding-dictionary-og_2x.jpg"
author: "dnw"
captured_at: "2026-09-14T00:47:35Z"
capture_tool: "hn-digest"
hn_id: 49690058
score: 1
comments: 0
posted_at: "2026-09-13T23:56:24Z"
tags:
  - hacker-news
---

# AI Coding Dictionary

- HN: [49690058](https://news.ycombinator.com/item?id=49690058)
- Source: [www.aihero.dev](https://www.aihero.dev/ai-coding-dictionary)
- Score: 1
- Comments: 0
- Posted: 2026-09-13T23:56:24Z

## Translation

Title: AI Coding Dictionary
Description: The vocabulary of AI coding, translated into plain English for engineers.

Article text:
AI Coding Dictionary AI Hero Courses
Map Skills Open source LLM Fundamentals AI Engineer Roadmap AI Coding Dictionary All posts AI Hero · Dictionary
The vocabulary of AI coding, in plain English .
Skimmable definitions for the terms that make agentic coding click. Search 69 entries below, or jump into a section.
4,569 · mattpocock / dictionary-of-ai-coding # the-model
# sessions,-context-windows-&-turns
Search the dictionary... Search the dictionary... Sections
# Sessions, Context Windows & Turns 8
A moving label, not a technology. Points at whatever computers can newly, impressively do — right now, large language models.
The parameters. Stateless — does next-token prediction and nothing else. Cannot do anything agentic on its own.
The numbers inside a model — often billions — tuned during training. Everything the model knows lives in them. Also called weights.
The process that sets a model's parameters by exposing it to vast amounts of text and adjusting to improve next-token prediction.
Running a trained model to generate output — what happens on every model provider request. Parameters stay fixed.
A dial for how much reasoning the model does before it answers. More effort spends more output tokens for a better shot at hard problems.
The atomic unit a model reads and writes. Roughly word-sized but not exactly. Context window size, cost, and latency all count tokens.
What the model actually does. Samples one next token from the context, appends it, and runs again. Its only mode of operation.
The same input can produce different output. A property of how models generate text and how providers serve requests.
Whatever serves a model for inference. Usually remote (Anthropic, OpenAI, Google), but can also be local (Ollama, llama.cpp).
Everything around the model that turns it into an agent: tools, system prompt, context-window management, permissions, hooks.
One round-trip from the harness to the model provider. The harness sends context; the provider returns one response.
Tokens the harness sends on each model provider request. Billed at a lower rate than output tokens.
Tokens the model generates back. Billed at a higher rate than input tokens, since they cost more compute to produce.
The provider-side store that lets consecutive requests skip re-processing a shared prefix, billing those tokens at a lower rate.
Input tokens the provider has cached from a previous request via its prefix cache, billed at a much lower rate.
Sessions, Context Windows & Turns
Carries no information forward. The model is stateless across requests; an agent is stateless across sessions by default.
The relevant information the agent has access to right now — what the agent knows that's pertinent to the task.
Everything the model sees on each model provider request. Finite, model-specific, the only surface through which the model perceives.
Carries information forward. Sessions are stateful across turns; agents can be made stateful across sessions via a memory system.
A model harnessed with tools, a system prompt, and a context window, that takes turns with a user. The model in motion.
The instructions the harness prepends to every model provider request — the agent's standing brief. Usually stable across a session.
One bounded run of interaction with an agent. Starts empty, accumulates, ends when cleared, closed, or compacted into a fresh session.
One user message plus everything the agent does in response, up until it yields back to the user. Contains one or more provider requests.
The world the agent acts on — anything outside the harness that the agent perceives via tool results and changes via tool calls.
A tree of files and directories the agent reads from, writes to, and executes within — the default environment for a coding agent.
A function the harness exposes for the agent to call — Read, Write, Bash, Search. How an agent perceives and acts on the environment.
The model's output naming a tool and its arguments — just structured text. The harness has to read it and execute.
What the harness sends back after executing a tool call — file contents, output, or error. The agent's only view of the environment.
A protocol for plugging external tool servers into a harness — how an agent gets tools beyond what the harness ships with.
What the harness shows the user before executing a tool call that isn't pre-approved. The mechanism for putting a human in the loop.
The permission-gating slice of an agent mode — which tool calls trigger a permission request and which run automatically.
A preset bundling a permission mode with behavioral instructions injected into the system prompt. Can flip mid-session.
An isolated environment the agent runs inside — container, VM, or restricted shell. Limits the blast radius of agent actions.
Confidently agreeable model output. Caused by training that shaped the model to favor answers humans liked — including agreement.
Confidently-wrong model output. Two flavors: factuality (invented facts) and faithfulness (drift from loaded context).
What the model knows from training, stored in its parameters. Frozen at training time. Counterpart to contextual knowledge.
The date past which a model has no parametric knowledge. Post-cutoff libraries and APIs are fabrication traps unless docs are loaded.
Facts the agent can read directly from the context right now. Counterpart to parametric knowledge.
The pairing between two tokens — meaningful pairs influence each other more than unrelated ones. A context of N tokens has ~N² of these.
Each token has a finite amount of influence to distribute across the rest of the context. Per-token, doesn't grow when context does.
As a session grows, each token's attention budget spreads across more competitors; signal on meaningful relationships shrinks.
Early in a session the agent is sharp and focused. As the session grows it drifts into a dumb zone: sloppier, forgetful, more mistakes.
Ending the current session and starting a fresh one. The next message begins with an empty session and an empty context window.
Transferring agent context from one session to another, with no return path. Carry mechanism varies — artifact, compaction, others.
The thing itself — code, transcripts, raw data. Complete and authoritative, but expensive to load into context.
An account of a primary source, one step removed — summaries, docs, compaction summaries. Cheap to load, lossy by construction.
A document used as the carry mechanism for a handoff — written by one session to be read by another.
A handoff artifact describing a multi-session piece of work — what's being built, not how each session does its share. Made of tickets.
A handoff artifact scoping one session of work. Stands alone or hangs off a spec. Can block or be blocked by sibling tickets.
A handoff done in-memory: the previous session's history is summarised and seeds a fresh session. Lossy — detail traded for headroom.
Compaction triggered automatically by the harness when the context window approaches full.
A system that attempts to make an agent stateful across sessions by persisting to the environment and reloading at session start.
A file in the environment that the harness loads into the context window at session start — the project's standing brief to the agent.
Loading only the context an agent needs right now, with context pointers to the rest. Borrowed from UI design.
A mention in one document that points to another, so the agent can pull it into context only when the task calls for it.
A teachable capability bundled as a unit — kept out of the context window until a context pointer pulls it in for the task at hand.
An agent spawned by another agent via a tool call. Runs in its own session, reports a single tool result. Cannot spawn further subagents.
A working pattern where one or more humans pair with the agent during a session — reviewing, redirecting, or collaborating in real time.
A working pattern where the user kicks off a session and leaves the agent to run unattended (away from keyboard).
A deterministic verification that runs in the environment — tests, type checks, lints, build, pre-commit hooks. Pass/fail, no judgement.
An agent reviewing another agent's work, often with a different model or system prompt. Non-deterministic: it forms a judgement.
The user reading the code the agent produced and forming a judgement on it. Reading the diff counts; reading the summary doesn't.
A working pattern where the user accepts the agent's code without human review. The diff is treated as opaque.
The shared understanding of what's being built, held in common between user and agent but separate from any asset.
A technique for developing a design concept: the agent interviews the user Socratically, one decision at a time.
Having the agent build a quick, rough version when conversation is too low-fidelity and you need a real artifact to talk about.
Developer experience: how easy a codebase and its toolchain make it for humans to do good work — docs, feedback speed, errors.
Agent experience: how well the environment is set up for an agent to do good work — checks, architecture, and free context.
Join AI Hero for practical skills, thinking on AI engineering, and resources that keep you ahead of the curve.
Model Context Protocol Tutorial
Claude Code for Real Engineers
Build Your Own AI Personal Assistant in TypeScript
Build DeepSearch in TypeScript

## Original Extract

The vocabulary of AI coding, translated into plain English for engineers.

AI Coding Dictionary AI Hero Courses
Map Skills Open source LLM Fundamentals AI Engineer Roadmap AI Coding Dictionary All posts AI Hero · Dictionary
The vocabulary of AI coding, in plain English .
Skimmable definitions for the terms that make agentic coding click. Search 69 entries below, or jump into a section.
4,569 · mattpocock / dictionary-of-ai-coding # the-model
# sessions,-context-windows-&-turns
Search the dictionary... Search the dictionary... Sections
# Sessions, Context Windows & Turns 8
A moving label, not a technology. Points at whatever computers can newly, impressively do — right now, large language models.
The parameters. Stateless — does next-token prediction and nothing else. Cannot do anything agentic on its own.
The numbers inside a model — often billions — tuned during training. Everything the model knows lives in them. Also called weights.
The process that sets a model's parameters by exposing it to vast amounts of text and adjusting to improve next-token prediction.
Running a trained model to generate output — what happens on every model provider request. Parameters stay fixed.
A dial for how much reasoning the model does before it answers. More effort spends more output tokens for a better shot at hard problems.
The atomic unit a model reads and writes. Roughly word-sized but not exactly. Context window size, cost, and latency all count tokens.
What the model actually does. Samples one next token from the context, appends it, and runs again. Its only mode of operation.
The same input can produce different output. A property of how models generate text and how providers serve requests.
Whatever serves a model for inference. Usually remote (Anthropic, OpenAI, Google), but can also be local (Ollama, llama.cpp).
Everything around the model that turns it into an agent: tools, system prompt, context-window management, permissions, hooks.
One round-trip from the harness to the model provider. The harness sends context; the provider returns one response.
Tokens the harness sends on each model provider request. Billed at a lower rate than output tokens.
Tokens the model generates back. Billed at a higher rate than input tokens, since they cost more compute to produce.
The provider-side store that lets consecutive requests skip re-processing a shared prefix, billing those tokens at a lower rate.
Input tokens the provider has cached from a previous request via its prefix cache, billed at a much lower rate.
Sessions, Context Windows & Turns
Carries no information forward. The model is stateless across requests; an agent is stateless across sessions by default.
The relevant information the agent has access to right now — what the agent knows that's pertinent to the task.
Everything the model sees on each model provider request. Finite, model-specific, the only surface through which the model perceives.
Carries information forward. Sessions are stateful across turns; agents can be made stateful across sessions via a memory system.
A model harnessed with tools, a system prompt, and a context window, that takes turns with a user. The model in motion.
The instructions the harness prepends to every model provider request — the agent's standing brief. Usually stable across a session.
One bounded run of interaction with an agent. Starts empty, accumulates, ends when cleared, closed, or compacted into a fresh session.
One user message plus everything the agent does in response, up until it yields back to the user. Contains one or more provider requests.
The world the agent acts on — anything outside the harness that the agent perceives via tool results and changes via tool calls.
A tree of files and directories the agent reads from, writes to, and executes within — the default environment for a coding agent.
A function the harness exposes for the agent to call — Read, Write, Bash, Search. How an agent perceives and acts on the environment.
The model's output naming a tool and its arguments — just structured text. The harness has to read it and execute.
What the harness sends back after executing a tool call — file contents, output, or error. The agent's only view of the environment.
A protocol for plugging external tool servers into a harness — how an agent gets tools beyond what the harness ships with.
What the harness shows the user before executing a tool call that isn't pre-approved. The mechanism for putting a human in the loop.
The permission-gating slice of an agent mode — which tool calls trigger a permission request and which run automatically.
A preset bundling a permission mode with behavioral instructions injected into the system prompt. Can flip mid-session.
An isolated environment the agent runs inside — container, VM, or restricted shell. Limits the blast radius of agent actions.
Confidently agreeable model output. Caused by training that shaped the model to favor answers humans liked — including agreement.
Confidently-wrong model output. Two flavors: factuality (invented facts) and faithfulness (drift from loaded context).
What the model knows from training, stored in its parameters. Frozen at training time. Counterpart to contextual knowledge.
The date past which a model has no parametric knowledge. Post-cutoff libraries and APIs are fabrication traps unless docs are loaded.
Facts the agent can read directly from the context right now. Counterpart to parametric knowledge.
The pairing between two tokens — meaningful pairs influence each other more than unrelated ones. A context of N tokens has ~N² of these.
Each token has a finite amount of influence to distribute across the rest of the context. Per-token, doesn't grow when context does.
As a session grows, each token's attention budget spreads across more competitors; signal on meaningful relationships shrinks.
Early in a session the agent is sharp and focused. As the session grows it drifts into a dumb zone: sloppier, forgetful, more mistakes.
Ending the current session and starting a fresh one. The next message begins with an empty session and an empty context window.
Transferring agent context from one session to another, with no return path. Carry mechanism varies — artifact, compaction, others.
The thing itself — code, transcripts, raw data. Complete and authoritative, but expensive to load into context.
An account of a primary source, one step removed — summaries, docs, compaction summaries. Cheap to load, lossy by construction.
A document used as the carry mechanism for a handoff — written by one session to be read by another.
A handoff artifact describing a multi-session piece of work — what's being built, not how each session does its share. Made of tickets.
A handoff artifact scoping one session of work. Stands alone or hangs off a spec. Can block or be blocked by sibling tickets.
A handoff done in-memory: the previous session's history is summarised and seeds a fresh session. Lossy — detail traded for headroom.
Compaction triggered automatically by the harness when the context window approaches full.
A system that attempts to make an agent stateful across sessions by persisting to the environment and reloading at session start.
A file in the environment that the harness loads into the context window at session start — the project's standing brief to the agent.
Loading only the context an agent needs right now, with context pointers to the rest. Borrowed from UI design.
A mention in one document that points to another, so the agent can pull it into context only when the task calls for it.
A teachable capability bundled as a unit — kept out of the context window until a context pointer pulls it in for the task at hand.
An agent spawned by another agent via a tool call. Runs in its own session, reports a single tool result. Cannot spawn further subagents.
A working pattern where one or more humans pair with the agent during a session — reviewing, redirecting, or collaborating in real time.
A working pattern where the user kicks off a session and leaves the agent to run unattended (away from keyboard).
A deterministic verification that runs in the environment — tests, type checks, lints, build, pre-commit hooks. Pass/fail, no judgement.
An agent reviewing another agent's work, often with a different model or system prompt. Non-deterministic: it forms a judgement.
The user reading the code the agent produced and forming a judgement on it. Reading the diff counts; reading the summary doesn't.
A working pattern where the user accepts the agent's code without human review. The diff is treated as opaque.
The shared understanding of what's being built, held in common between user and agent but separate from any asset.
A technique for developing a design concept: the agent interviews the user Socratically, one decision at a time.
Having the agent build a quick, rough version when conversation is too low-fidelity and you need a real artifact to talk about.
Developer experience: how easy a codebase and its toolchain make it for humans to do good work — docs, feedback speed, errors.
Agent experience: how well the environment is set up for an agent to do good work — checks, architecture, and free context.
Join AI Hero for practical skills, thinking on AI engineering, and resources that keep you ahead of the curve.
Model Context Protocol Tutorial
Claude Code for Real Engineers
Build Your Own AI Personal Assistant in TypeScript
Build DeepSearch in TypeScript
