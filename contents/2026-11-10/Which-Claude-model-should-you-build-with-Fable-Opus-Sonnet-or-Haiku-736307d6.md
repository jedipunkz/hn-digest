---
source: "https://claude.dev/blog/which-claude-model-should-you-build-with/"
hn_url: "https://news.ycombinator.com/item?id=50435447"
title: "Which Claude model should you build with? Fable, Opus, Sonnet, or Haiku"
article_title: "Which Claude model should you build with? Fable, Opus, Sonnet, or Haiku / claude.dev Blog"
image: "https://claude.dev/blog/which-claude-model-should-you-build-with/og.png"
author: "882542F3884314B"
captured_at: "2026-11-10T03:13:43Z"
capture_tool: "hn-digest"
hn_id: 50435447
score: 2
comments: 0
posted_at: "2026-11-10T02:45:00Z"
tags:
  - hacker-news
---

# Which Claude model should you build with? Fable, Opus, Sonnet, or Haiku

- HN: [50435447](https://news.ycombinator.com/item?id=50435447)
- Source: [claude.dev](https://claude.dev/blog/which-claude-model-should-you-build-with/)
- Score: 2
- Comments: 0
- Posted: 2026-11-10T02:45:00Z

## Translation

Title: Which Claude model should you build with? Fable, Opus, Sonnet, or Haiku
Article title: Which Claude model should you build with? Fable, Opus, Sonnet, or Haiku / claude.dev Blog
Description: Which Claude model to pick, what they cost, and how to use them together.

Article text:
Which Claude model should you build with? Fable, Opus, Sonnet, or Haiku / claude.dev Blog [ H ] HOME [ M ] MODS [ D ] DOCS [ T ] TERMINAL Try Claude Code [ H ] HOME [ M ] MODS [ D ] DOCS [ T ] TERMINAL Playbooks Which Claude model should you build with? Fable, Opus, Sonnet, or Haiku
Which Claude model to pick, what they cost, and how to use them together.
The current Claude lineup has four generally available models: Claude Haiku 5.5, Claude Sonnet 5.5, Claude Opus 5.5 and Claude Fable 5.2. Claude Mythos 5.2 is our most capable model for cybersecurity and life sciences research, available through our verification programs.
If you build on the Claude Platform or work in Claude Code, you may be wondering which model to use for your common tasks. This article will help answer that question and how to think about the trade-offs.
The quick answer is that you should start most work on Opus 5.5 at its default effort level, medium .
There's a reason it's the default model in Claude Code. It costs half as much per input and output token as Fable 5.2, and it's the model we recommend starting with for most coding and knowledge work.
Move down to Sonnet 5.5 for well-scoped everyday work, like a bug fix or a feature built against a spec. Just because you have a clear spec doesn't mean you should use Sonnet, though. A migration or a performance optimization can have clear specs and still be difficult enough to warrant Opus 5.5, or Fable 5.2 for a major one.
Move further down to Haiku 5.5 for the simplest high-volume work, such as classification, routing, and extraction. Haiku 5.5 is also good to use with subagents that work on a larger model's side tasks.
Move up to Claude Fable 5.2 for the hardest technical projects in coding, finance and the sciences. Also move up when Opus 5.5 at high still falls short, or for long unattended runs.
Before you move models, try adjusting your effort level. It's a finer dial for both cost and capability. Past high , one more effort level can cost as much as a bigger model, so compare both on your own work. If you're past high , test the next model up at a lower effort instead. Conversely, one level down can get you the result you need at a lower rate.
Higher effort helps when the model needs to think longer, call more tools, or check its own work more. Just keep in mind that it also adds latency. Additionally, moving to a higher effort level won't give a model knowledge or skill it doesn't have. For that, you still need to move up a model.
Prices are Claude API list prices. For benchmark results, see the Fable 5.2 page .
When you choose between models, you're making a trade-off. When you pick a larger model, for example, you are choosing capability over latency. These are the main trade-offs to think through.
Capability and cost per finished task
Generally, bigger models cost more per token and are more capable, but the cost vs capability trade-off isn't universally true at the task level.
A higher price per token doesn't always mean a higher cost per task, or a better result. As we showed in Spending your effort , on Anthropic's runs of Terminal-Bench 3.0, the previous Fable, 5.1, at max scored 58.0%, while Opus 5.5 at high scored 58.9% with half the tokens.
On the hardest work, the price gap can close. On FrontierCode 2.0, Fable 5.2 at high scored about level with Opus 5.5 at max (45.3% vs. 45.5%) for roughly half the cost per task. So before you raise Opus 5.5 past high on a hard task, test Fable 5.2. For everyday work, Opus 5.5 at medium is still the place to start. For benchmark results, see the launch post .
However, the Claude Platform's cost and intelligence measurements also show the limits of a small model. These generally do well on high-volume work with checkable outputs and are less suited to long agentic loops.
When you compare two models, price the hardest tasks in your set as well as the typical ones, because the hardest tasks carry most of the spend. On one 20-problem search benchmark, two problems carried 43% of the spend.
A faster setup finishes sooner and can miss more. Think through how long a person can wait before you pick your model and effort level. For replies a person waits on, start Sonnet 5.5 at low or medium .
Often next generation models are faster than their predecessors. Fable 5.2 generates output up to 30% faster than Fable 5.1. Sonnet 5.5 runs 30% faster than Sonnet 5.
Effort level and latency are roughly correlated. Determine if the accuracy you gain is worth the extra latency. For example, on SWE-bench Pro, Opus 5.5 at medium scored about 2.5 points lower for about 70% of the cost as the high effort setting. At low it scored about 8 points lower for about a third of the cost. At xhigh it scored about 1.4 points higher for 2.5 times the cost of high .
Often providing a means of self-verification is more cost effective and adds less latency than changing effort levels. When a model isn't good enough for a task, take these steps in order:
Check that the model can verify its own work, with a test, a build or a script that calls the endpoint.
Move to the next model up, starting at a lower effort level than you used on the smaller one.
If you can check outcomes automatically, run everything at a low level and re-run only the failures at a high one. On the same SWE-bench Pro subset, Opus 5.5 at low with failures re-run at high passed about 97% of tasks for about $0.17 each, against 95.3% for $0.29 with every task at high .
If you need a faster reply from the same model, you can pay for speed instead of giving up capability. On the Claude API, Opus 5.5 and Fable 5.2 both offer fast mode , with higher output speed at premium pricing.
Every model in the current generation has a 1M-token context window. Large context can help the model complete the task but it costs you on every turn, because each turn resends it. That being said, on long runs most of your input bills as cache reads .
Opus 5.5 charges $0.20 per million tokens for a cache read, and Sonnet 5.5 charges $0.10. Fable 5.2 charges the same $0.20, so on a long, cache-heavy run the gap between Fable and Opus comes down to uncached input and output, where Fable costs twice as much per token.
In What a task costs on Opus 5.5 , a turn's cache read costs about $0.03 at 150K tokens of context. At 20K it costs about $0.004.
One takeaway from this is that a lean main context matters more the bigger the main model is. Keep the expensive model's context for the decisions, and let cheaper subagents do the verbose reading.
Haiku 5.5 is the only model where the token price varies by prompt length. A prompt of up to 100K tokens bills at the lower tier, and a prompt over 100K bills at five times that rate. The long-prompt tier still costs a quarter of Sonnet 5.5's per-token price. So Haiku 5.5 pays off most on short prompts, where it costs a twentieth of Sonnet 5.5's input price.
When every call needs a long prompt, compare Haiku 5.5 with Sonnet 5.5 on cost per finished task, not list price.
The prompt cache belongs to the model that wrote it. Your first turn on a new model pays the cache-write price on the whole conversation. To make that turn smaller, run /compact first in Claude Code, or start a fresh session with a short written plan. On the API, open the new model's conversation with a summary of the work so far instead of the full history.
Each thinking block records the model that produced it, and each model reads only some other models' blocks. On the Claude API and Claude Models on Google Cloud (Google managed), Opus 5.5 reads Sonnet 5.5's and Haiku 5.5's thinking blocks. Sonnet 5.5 reads Haiku 5.5's. Fable 5.2 reads Opus 5.5's, Sonnet 5.5's and Haiku 5.5's. On the Claude API, Fable 5.1 reads Opus 5.5's.
A move down a tier loses the earlier reasoning, because a model never reads the blocks of a model above it. Other moves can lose it too. When the new model can't read the earlier model's thinking blocks, the API drops them and the request still succeeds. The new model then works without that reasoning.
Preserved thinking lists every case. A move from Fable 5.2 to any other generally available model drops its reasoning.
Put subagents on smaller models
A subagent runs in its own context window and hands back a summary. You pay for its tokens at the rates of the model it runs on. In Claude Code, set model: haiku or model: sonnet in a subagent's definition. The CLAUDE_CODE_SUBAGENT_MODEL environment variable sets the default for subagents that don't specify a model.
Save this as .claude/agents/log-reader.md :
---
name: log-reader
description: Reads test output and logs, and reports each failing case. Use after a test run.
tools: Read, Grep, Glob, Bash
model: haiku
---
Read the logs you're given. For each failure, report the file, the test name and the first error line. Don't edit any files.
Keep small models on work where a mistake is cheap to spot, like finding files, running tests and reading logs. Keep code edits and judgment calls on the main model. When a small model misreads a search result, the main model opens the wrong file and spends tokens on it.
In current versions of Claude Code, the built-in Explore subagent runs on your session's model instead of always running on Haiku. On the Claude API, that model is capped at Opus. To keep exploration cheap, define your own subagent named Explore with model: haiku . You can pin the model the haiku alias uses with ANTHROPIC_DEFAULT_HAIKU_MODEL .
On the Anthropic API, the haiku alias now resolves to Haiku 5.5, which needs Claude Code v2.1.293 or later. On Amazon Bedrock, Google Cloud, Microsoft Foundry and Claude Platform on AWS, it still resolves to Haiku 4.5. Claude Code's model configuration docs have the details.
Claude Code also has the opusplan alias, which plans with Opus in plan mode and then carries out the plan with Sonnet. That puts your code edits on Sonnet. The advice above does the reverse. It keeps code edits on the main model and gives smaller models only work where a mistake is cheap to spot. Measure opusplan on your own tasks before you make it your default.
Escalate hard decisions to an advisor
The advisor tool (beta) pairs a faster, cheaper executor model with a stronger advisor model. The executor runs the task and can ask the advisor for guidance partway through. Most of the tokens bill at the executor's rate. The advisor helps only when it adds capability the executor lacks and the executor asks for advice, which it does less often at low effort.
Where it might pay off: agentic work where most turns are routine. The platform's cost guide is cautious: of the four pairings in its advisor table, none clearly beat its executor alone, and an advisor can help only when the executor consults it. Measure your executor's consult rate first. Check which advisors your executor accepts. A Sonnet 5.5 executor accepts Fable 5.2, Opus 5.5, Opus 5 and Sonnet 5.5, for example, and a Fable 5.2 executor accepts only Fable 5.2 or Mythos 5.2.
Where it adds little: one of those pairings was an Opus 5.5 executor at high with a Fable 5.1 advisor on an internal coding benchmark. It scored 90.1% at $2.92 per attempt. That's 1.7 points over Opus 5.5 alone at high for about 2.1 times the money. A gain that size is at the edge of run-to-run noise, and about what raising effort would have bought.
Delegate bulk work to cheaper workers
Subagents, above, handle side tasks while one main model keeps the thread. This pattern is for work bigger than one context window. A top-tier model leads, splits the task, hands the pieces to many cheaper workers in parallel, and merges what comes back. On an internal defect-search benchmark over about 22 million tokens of source code, an Opus 5.5 lead over Sonnet 5.5 workers, up to 25 at a time, cost about a third less than Opus 5.5 alone at its default effort. It scored about the same.
On the Claude Platform, multiagent orchestration in Claude Managed Agents gives you a coordinator agent and a roster

[truncated]

## Original Extract

Which Claude model to pick, what they cost, and how to use them together.

Which Claude model should you build with? Fable, Opus, Sonnet, or Haiku / claude.dev Blog [ H ] HOME [ M ] MODS [ D ] DOCS [ T ] TERMINAL Try Claude Code [ H ] HOME [ M ] MODS [ D ] DOCS [ T ] TERMINAL Playbooks Which Claude model should you build with? Fable, Opus, Sonnet, or Haiku
Which Claude model to pick, what they cost, and how to use them together.
The current Claude lineup has four generally available models: Claude Haiku 5.5, Claude Sonnet 5.5, Claude Opus 5.5 and Claude Fable 5.2. Claude Mythos 5.2 is our most capable model for cybersecurity and life sciences research, available through our verification programs.
If you build on the Claude Platform or work in Claude Code, you may be wondering which model to use for your common tasks. This article will help answer that question and how to think about the trade-offs.
The quick answer is that you should start most work on Opus 5.5 at its default effort level, medium .
There's a reason it's the default model in Claude Code. It costs half as much per input and output token as Fable 5.2, and it's the model we recommend starting with for most coding and knowledge work.
Move down to Sonnet 5.5 for well-scoped everyday work, like a bug fix or a feature built against a spec. Just because you have a clear spec doesn't mean you should use Sonnet, though. A migration or a performance optimization can have clear specs and still be difficult enough to warrant Opus 5.5, or Fable 5.2 for a major one.
Move further down to Haiku 5.5 for the simplest high-volume work, such as classification, routing, and extraction. Haiku 5.5 is also good to use with subagents that work on a larger model's side tasks.
Move up to Claude Fable 5.2 for the hardest technical projects in coding, finance and the sciences. Also move up when Opus 5.5 at high still falls short, or for long unattended runs.
Before you move models, try adjusting your effort level. It's a finer dial for both cost and capability. Past high , one more effort level can cost as much as a bigger model, so compare both on your own work. If you're past high , test the next model up at a lower effort instead. Conversely, one level down can get you the result you need at a lower rate.
Higher effort helps when the model needs to think longer, call more tools, or check its own work more. Just keep in mind that it also adds latency. Additionally, moving to a higher effort level won't give a model knowledge or skill it doesn't have. For that, you still need to move up a model.
Prices are Claude API list prices. For benchmark results, see the Fable 5.2 page .
When you choose between models, you're making a trade-off. When you pick a larger model, for example, you are choosing capability over latency. These are the main trade-offs to think through.
Capability and cost per finished task
Generally, bigger models cost more per token and are more capable, but the cost vs capability trade-off isn't universally true at the task level.
A higher price per token doesn't always mean a higher cost per task, or a better result. As we showed in Spending your effort , on Anthropic's runs of Terminal-Bench 3.0, the previous Fable, 5.1, at max scored 58.0%, while Opus 5.5 at high scored 58.9% with half the tokens.
On the hardest work, the price gap can close. On FrontierCode 2.0, Fable 5.2 at high scored about level with Opus 5.5 at max (45.3% vs. 45.5%) for roughly half the cost per task. So before you raise Opus 5.5 past high on a hard task, test Fable 5.2. For everyday work, Opus 5.5 at medium is still the place to start. For benchmark results, see the launch post .
However, the Claude Platform's cost and intelligence measurements also show the limits of a small model. These generally do well on high-volume work with checkable outputs and are less suited to long agentic loops.
When you compare two models, price the hardest tasks in your set as well as the typical ones, because the hardest tasks carry most of the spend. On one 20-problem search benchmark, two problems carried 43% of the spend.
A faster setup finishes sooner and can miss more. Think through how long a person can wait before you pick your model and effort level. For replies a person waits on, start Sonnet 5.5 at low or medium .
Often next generation models are faster than their predecessors. Fable 5.2 generates output up to 30% faster than Fable 5.1. Sonnet 5.5 runs 30% faster than Sonnet 5.
Effort level and latency are roughly correlated. Determine if the accuracy you gain is worth the extra latency. For example, on SWE-bench Pro, Opus 5.5 at medium scored about 2.5 points lower for about 70% of the cost as the high effort setting. At low it scored about 8 points lower for about a third of the cost. At xhigh it scored about 1.4 points higher for 2.5 times the cost of high .
Often providing a means of self-verification is more cost effective and adds less latency than changing effort levels. When a model isn't good enough for a task, take these steps in order:
Check that the model can verify its own work, with a test, a build or a script that calls the endpoint.
Move to the next model up, starting at a lower effort level than you used on the smaller one.
If you can check outcomes automatically, run everything at a low level and re-run only the failures at a high one. On the same SWE-bench Pro subset, Opus 5.5 at low with failures re-run at high passed about 97% of tasks for about $0.17 each, against 95.3% for $0.29 with every task at high .
If you need a faster reply from the same model, you can pay for speed instead of giving up capability. On the Claude API, Opus 5.5 and Fable 5.2 both offer fast mode , with higher output speed at premium pricing.
Every model in the current generation has a 1M-token context window. Large context can help the model complete the task but it costs you on every turn, because each turn resends it. That being said, on long runs most of your input bills as cache reads .
Opus 5.5 charges $0.20 per million tokens for a cache read, and Sonnet 5.5 charges $0.10. Fable 5.2 charges the same $0.20, so on a long, cache-heavy run the gap between Fable and Opus comes down to uncached input and output, where Fable costs twice as much per token.
In What a task costs on Opus 5.5 , a turn's cache read costs about $0.03 at 150K tokens of context. At 20K it costs about $0.004.
One takeaway from this is that a lean main context matters more the bigger the main model is. Keep the expensive model's context for the decisions, and let cheaper subagents do the verbose reading.
Haiku 5.5 is the only model where the token price varies by prompt length. A prompt of up to 100K tokens bills at the lower tier, and a prompt over 100K bills at five times that rate. The long-prompt tier still costs a quarter of Sonnet 5.5's per-token price. So Haiku 5.5 pays off most on short prompts, where it costs a twentieth of Sonnet 5.5's input price.
When every call needs a long prompt, compare Haiku 5.5 with Sonnet 5.5 on cost per finished task, not list price.
The prompt cache belongs to the model that wrote it. Your first turn on a new model pays the cache-write price on the whole conversation. To make that turn smaller, run /compact first in Claude Code, or start a fresh session with a short written plan. On the API, open the new model's conversation with a summary of the work so far instead of the full history.
Each thinking block records the model that produced it, and each model reads only some other models' blocks. On the Claude API and Claude Models on Google Cloud (Google managed), Opus 5.5 reads Sonnet 5.5's and Haiku 5.5's thinking blocks. Sonnet 5.5 reads Haiku 5.5's. Fable 5.2 reads Opus 5.5's, Sonnet 5.5's and Haiku 5.5's. On the Claude API, Fable 5.1 reads Opus 5.5's.
A move down a tier loses the earlier reasoning, because a model never reads the blocks of a model above it. Other moves can lose it too. When the new model can't read the earlier model's thinking blocks, the API drops them and the request still succeeds. The new model then works without that reasoning.
Preserved thinking lists every case. A move from Fable 5.2 to any other generally available model drops its reasoning.
Put subagents on smaller models
A subagent runs in its own context window and hands back a summary. You pay for its tokens at the rates of the model it runs on. In Claude Code, set model: haiku or model: sonnet in a subagent's definition. The CLAUDE_CODE_SUBAGENT_MODEL environment variable sets the default for subagents that don't specify a model.
Save this as .claude/agents/log-reader.md :
---
name: log-reader
description: Reads test output and logs, and reports each failing case. Use after a test run.
tools: Read, Grep, Glob, Bash
model: haiku
---
Read the logs you're given. For each failure, report the file, the test name and the first error line. Don't edit any files.
Keep small models on work where a mistake is cheap to spot, like finding files, running tests and reading logs. Keep code edits and judgment calls on the main model. When a small model misreads a search result, the main model opens the wrong file and spends tokens on it.
In current versions of Claude Code, the built-in Explore subagent runs on your session's model instead of always running on Haiku. On the Claude API, that model is capped at Opus. To keep exploration cheap, define your own subagent named Explore with model: haiku . You can pin the model the haiku alias uses with ANTHROPIC_DEFAULT_HAIKU_MODEL .
On the Anthropic API, the haiku alias now resolves to Haiku 5.5, which needs Claude Code v2.1.293 or later. On Amazon Bedrock, Google Cloud, Microsoft Foundry and Claude Platform on AWS, it still resolves to Haiku 4.5. Claude Code's model configuration docs have the details.
Claude Code also has the opusplan alias, which plans with Opus in plan mode and then carries out the plan with Sonnet. That puts your code edits on Sonnet. The advice above does the reverse. It keeps code edits on the main model and gives smaller models only work where a mistake is cheap to spot. Measure opusplan on your own tasks before you make it your default.
Escalate hard decisions to an advisor
The advisor tool (beta) pairs a faster, cheaper executor model with a stronger advisor model. The executor runs the task and can ask the advisor for guidance partway through. Most of the tokens bill at the executor's rate. The advisor helps only when it adds capability the executor lacks and the executor asks for advice, which it does less often at low effort.
Where it might pay off: agentic work where most turns are routine. The platform's cost guide is cautious: of the four pairings in its advisor table, none clearly beat its executor alone, and an advisor can help only when the executor consults it. Measure your executor's consult rate first. Check which advisors your executor accepts. A Sonnet 5.5 executor accepts Fable 5.2, Opus 5.5, Opus 5 and Sonnet 5.5, for example, and a Fable 5.2 executor accepts only Fable 5.2 or Mythos 5.2.
Where it adds little: one of those pairings was an Opus 5.5 executor at high with a Fable 5.1 advisor on an internal coding benchmark. It scored 90.1% at $2.92 per attempt. That's 1.7 points over Opus 5.5 alone at high for about 2.1 times the money. A gain that size is at the edge of run-to-run noise, and about what raising effort would have bought.
Delegate bulk work to cheaper workers
Subagents, above, handle side tasks while one main model keeps the thread. This pattern is for work bigger than one context window. A top-tier model leads, splits the task, hands the pieces to many cheaper workers in parallel, and merges what comes back. On an internal defect-search benchmark over about 22 million tokens of source code, an Opus 5.5 lead over Sonnet 5.5 workers, up to 25 at a time, cost about a third less than Opus 5.5 alone at its default effort. It scored about the same.
On the Claude Platform, multiagent orchestration in Claude Managed Agents gives you a coordinator agent and a roster

[truncated]
