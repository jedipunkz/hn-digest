---
source: "https://keln.ai/blog/ai-inference-provider-benchmark-2026"
hn_url: "https://news.ycombinator.com/item?id=50431519"
title: "We Benchmarked Fireworks, Together, Baseten, Modal, DeepInfra, Z.ai and Keln"
article_title: "Modal vs Fireworks vs Together AI vs Keln: Benchmark | Keln"
image: "https://keln.ai/blog/ai-inference-provider-benchmark-2026-preview.png"
author: "kris-holy"
captured_at: "2026-11-09T20:27:52Z"
capture_tool: "hn-digest"
hn_id: 50431519
score: 1
comments: 0
posted_at: "2026-11-09T20:26:56Z"
tags:
  - hacker-news
---

# We Benchmarked Fireworks, Together, Baseten, Modal, DeepInfra, Z.ai and Keln

- HN: [50431519](https://news.ycombinator.com/item?id=50431519)
- Source: [keln.ai](https://keln.ai/blog/ai-inference-provider-benchmark-2026)
- Score: 1
- Comments: 0
- Posted: 2026-11-09T20:26:56Z

## Translation

Title: We Benchmarked Fireworks, Together, Baseten, Modal, DeepInfra, Z.ai and Keln
Article title: Modal vs Fireworks vs Together AI vs Keln: Benchmark | Keln
Description: We benchmarked Keln, Fireworks, Together, Modal, Baseten, DeepInfra and Z.ai on real agent sessions. Compare speed, reliability, prompt caching and cost.

Article text:
Skip to content Keln Keln home Models & Pricing Docs Blog For providers ↗ Log in Sign up Models & Pricing Docs Blog For providers ↗ Contact Sign up Log in Blog Benchmarks Benchmark DeepSeek V4.1 Flash 1017 and GLM-5.3 Flash November 4, 2026 8 min read AI Inference Benchmark: Keln vs Fireworks, Together AI, Baseten, Modal, DeepInfra & Z.ai
We tested DeepSeek V4.1 Flash 1017 and GLM-5.3 Flash across six inference providers and Keln, using real agent sessions with tool calling and maximum reasoning enabled. Each provider ran the same conversations, in the same order, with identical pauses between turns. You can find the full methodology at the end.
Copy page .
Sources: Keln cube from the gateway benchmark post; Together, Fireworks, Baseten, Modal and Z.ai
from the models.dev provider set; DeepInfra from the lobehub icon set. Brand colours are set in
blog.css under .bench--direct. -->
Benchmark results at a glance
Most sides answered everything we sent on both models. Baseten refused about half the DeepSeek turns with rate-limit errors: these prompts are tens of thousands of tokens each, and just 2 sessions at once put us over our account's limit. On GLM-5.3 Flash the breaks kept the load lower, so Baseten refused fewer turns there.
Modal cut 1 stream on DeepSeek, while on all the other providers answers ended normally. An app calling Modal directly gets the cut stream and has to resend the turn itself, because there is no second provider to send it to.
Running inference across multiple providers can hurt prompt caching when conversations move between them. On DeepSeek, Keln switched providers on just 1 of 1,440 turns, keeping its cache hit rate at 94.3%, close to the best direct provider at 95%. On GLM-5.3 Flash, Keln reached 90.5%, the highest of the 6 sides.
Fireworks had the most cold turns on DeepSeek: 165, compared with 11 on Keln. Each cold turn meant reading most of the prompt again instead of using cached tokens.
Based on each side's usage . A cold turn is any turn after the first with under 10% cached.
We also tested how well providers retain cached prompts when a conversation goes idle. On GLM-5.3 Flash, we added breaks of 1–20 minutes to 15% of turns, using the same timing across all providers.
0 20 40 60 80 63.0% Keln 10 turns 38.4% DeepInfra 10 turns 16.5% Z.ai 10 turns 0.0% Baseten 10 turns 0.0% Modal 10 turns 0.0% Fireworks 10 turns
After a break of 1 to 5 min 5 to 10 min 10 to 15 min 15 to 20 min Turns with no break Whole run
Keln 77.1% 6 turns 91.7% 8 turns 78.2% 11 turns 63.0% 10 turns 94.0% 90.5%
DeepInfra 97.5% 6 turns 80.2% 8 turns 59.3% 11 turns 38.4% 10 turns 94.0% 88.7%
Z.ai 83.3% 6 turns 81.0% 8 turns 33.6% 11 turns 16.5% 10 turns 93.8% 85.9%
Baseten 69.3% 6 turns 42.6% 8 turns 35.5% 11 turns 0.0% 10 turns 91.6% 80.4%
Modal 49.1% 6 turns 19.7% 8 turns 0.0% 11 turns 0.0% 10 turns 91.9% 78.8%
Fireworks 97.8% 6 turns 81.0% 8 turns 22.3% 11 turns 0.0% 10 turns 96.3% 87.1%
After a 15–20 minute break, Keln still had 63.0% of the prompt cached, compared with 38.4% on DeepInfra and 16.5% on Z.ai. Baseten, Modal and Fireworks had 0%.
Keln had the highest median decode speed on GLM-5.3 Flash at 432 tokens/s, 75% faster than the next fastest provider, Baseten, at 247 tokens/s.
On DeepSeek, Keln and Together were nearly tied at 432 and 437 tokens/s, both well ahead of the other providers.
0 100 200 300 400 500 432 Keln 240 turns 247 Baseten 217 turns 204 Fireworks 240 turns 176 DeepInfra 240 turns 129 Modal 240 turns 46 Z.ai 240 turns
Decode speed, median DeepSeek V4.1 Flash 1017 GLM-5.3 Flash
Keln 432 tok/s first token 0.35 / 0.89 s 432 tok/s first token 1.01 / 2.97 s
Fireworks 234 tok/s first token 1.21 / 2.45 s 204 tok/s first token 0.79 / 1.66 s
Together 437 tok/s first token 0.22 / 0.65 s –
Baseten 208 tok/s first token 0.38 / 1.07 s 247 tok/s first token 0.38 / 1.59 s
Modal 228 tok/s first token 0.72 / 1.45 s 129 tok/s first token 0.99 / 2.48 s
DeepInfra 98 tok/s first token 1.05 / 3.19 s 176 tok/s first token 1.16 / 5.20 s
Z.ai – 46 tok/s first token 4.98 / 27.11 s
First token is p50 / p95 and counts the first reasoning token.
We set reasoning effort to maximum for every request. All providers returned broadly similar amounts of reasoning tokens, including Keln.
In our previous benchmark , two gateways returned less than half as many reasoning tokens for the same requests. We didn't see that issue here.
Modal doesn't report reasoning tokens for GLM-5.3 Flash, so those results aren't included.
Fireworks was the most expensive on DeepSeek, largely due to its 165 cold turns. It processed 2.5 times as many uncached prompt tokens as Keln.
On GLM-5.3 Flash, Keln was among the more expensive options. Its pricing covers more than model inference: Keln manages capacity across providers, monitors performance and output quality, and handles failover and retries. You get all of that through one API, without managing provider accounts, rate limits or caching yourself.
Baseten's cost per 1,000 turns counts only the turns it answered.
Going directly to an inference provider can be cheaper, and performance can be excellent when capacity is available. But you're limited to that provider's infrastructure and your account's rate limits.
We saw those limits in this benchmark. Baseten rejected requests due to rate limits, Modal cut off a response mid-stream, and Fireworks had to process previously cached prompts again.
Keln takes a different approach. It runs inference across multiple providers, managing capacity, performance, and reliability behind one API. In this benchmark, Keln completed every turn, matched or exceeded the fastest direct providers on decode speed, and retained more cached tokens after longer breaks.
Keln wasn't the cheapest option, but the price includes the infrastructure management, failover, and retries you'd otherwise have to handle yourself.
We post benchmarks with the method included. The previous one compared OpenRouter, Vercel, Ramp and Keln on DeepSeek V4.1 Flash for 24 hours: read it here .
We used 8 sessions from the public armand0e Claude Code corpus on Hugging Face, replayed with native tool calls. Each recorded assistant turn was sent as a new request, with the original tool results appended, so conversations grew as they did in the original sessions.
Each run used up to 30 rounds per session, with 2 concurrent sessions per side and fresh salts and conversation IDs for each pass. The median turn contained 40,665 tokens on DeepSeek and 38,578 on GLM-5.3 Flash.
All requests used reasoning_effort: max in the format supported by each API, temperature 1, streaming enabled, and no max_tokens limit.
Pauses between turns followed the original sessions, capped at 60 seconds, with additional breaks to test caching. On DeepSeek, 5% of turns had breaks of 30 seconds to 2 minutes. On GLM-5.3 Flash, 15% had breaks of 1–20 minutes. We used the same random seed across all sides, so each received the same pauses on the same turns.
The DeepSeek benchmark ran on November 3, 2026, from 02:30 to 08:53 UTC, with 8,640 turns total (1,440 per side). The GLM-5.3 Flash run made 3 passes in 14 hours on November 4; a pass replays the 8 sessions once. This post reports the third pass, from 03:45 to 08:32 UTC, with 1,440 turns total (roughly 240 per side).
We tested Keln through its public API. For direct providers, we used Keln's own accounts and each provider's model identifiers, reasoning settings, and per-conversation cache hints. These included the documented user field or, for Baseten, the x-session-affinity header. Fireworks served DeepSeek through its standard endpoint.
Cache hit rate is the percentage of prompt tokens reported as cached in each provider's usage data. A provider switch is counted when consecutive turns in the same conversation run on different providers. Decode speed is calculated as output tokens divided by generation time for each turn.
We calculated costs using each provider's published model prices as of November 4, 2026, applied to its reported usage. Fireworks offers two DeepSeek pricing tiers; we used the standard tier corresponding to the endpoint tested. Modal shows its rates in its console, so we used the prices shown on our account.
For Keln, we used the November 4 price list: DeepSeek at $0.30 per million input tokens, $0.01 per million cached input tokens, and $1.20 per million output tokens; GLM-5.3 Flash at $0.15, $0.05, and $0.50 respectively. The DeepSeek run was billed at the previous cached-input rate of $0.0108 per million tokens, resulting in $2.09 per 1,000 turns.
Results reflect this workload, our account limits, and the conditions during the test. Performance may vary with other workloads or concurrency levels.
The raw benchmark rows and response bodies are not publicly available.
Claude Setup Run open models in the desktop app. No code. Codex Setup Run open models in Codex, terminal or desktop DeepSeek V4.1 Flash 1017 GLM-5.3 Flash Benchmark Prompt caching Pricing Try it yourself Keln Quickstart Send your first request in under 2 minutes Claude Setup Run open models in the desktop app. No code. Codex Setup Run open models in Codex, terminal or desktop Run your agent sessions on Keln
Cheat detection, high cache hit rate and normalized output.
Built by a team from NVIDIA and Google.
© 2026 Keln, Inc. 2261 Market Street STE 32997, San Francisco, CA 94114

## Original Extract

We benchmarked Keln, Fireworks, Together, Modal, Baseten, DeepInfra and Z.ai on real agent sessions. Compare speed, reliability, prompt caching and cost.

Skip to content Keln Keln home Models & Pricing Docs Blog For providers ↗ Log in Sign up Models & Pricing Docs Blog For providers ↗ Contact Sign up Log in Blog Benchmarks Benchmark DeepSeek V4.1 Flash 1017 and GLM-5.3 Flash November 4, 2026 8 min read AI Inference Benchmark: Keln vs Fireworks, Together AI, Baseten, Modal, DeepInfra & Z.ai
We tested DeepSeek V4.1 Flash 1017 and GLM-5.3 Flash across six inference providers and Keln, using real agent sessions with tool calling and maximum reasoning enabled. Each provider ran the same conversations, in the same order, with identical pauses between turns. You can find the full methodology at the end.
Copy page .
Sources: Keln cube from the gateway benchmark post; Together, Fireworks, Baseten, Modal and Z.ai
from the models.dev provider set; DeepInfra from the lobehub icon set. Brand colours are set in
blog.css under .bench--direct. -->
Benchmark results at a glance
Most sides answered everything we sent on both models. Baseten refused about half the DeepSeek turns with rate-limit errors: these prompts are tens of thousands of tokens each, and just 2 sessions at once put us over our account's limit. On GLM-5.3 Flash the breaks kept the load lower, so Baseten refused fewer turns there.
Modal cut 1 stream on DeepSeek, while on all the other providers answers ended normally. An app calling Modal directly gets the cut stream and has to resend the turn itself, because there is no second provider to send it to.
Running inference across multiple providers can hurt prompt caching when conversations move between them. On DeepSeek, Keln switched providers on just 1 of 1,440 turns, keeping its cache hit rate at 94.3%, close to the best direct provider at 95%. On GLM-5.3 Flash, Keln reached 90.5%, the highest of the 6 sides.
Fireworks had the most cold turns on DeepSeek: 165, compared with 11 on Keln. Each cold turn meant reading most of the prompt again instead of using cached tokens.
Based on each side's usage . A cold turn is any turn after the first with under 10% cached.
We also tested how well providers retain cached prompts when a conversation goes idle. On GLM-5.3 Flash, we added breaks of 1–20 minutes to 15% of turns, using the same timing across all providers.
0 20 40 60 80 63.0% Keln 10 turns 38.4% DeepInfra 10 turns 16.5% Z.ai 10 turns 0.0% Baseten 10 turns 0.0% Modal 10 turns 0.0% Fireworks 10 turns
After a break of 1 to 5 min 5 to 10 min 10 to 15 min 15 to 20 min Turns with no break Whole run
Keln 77.1% 6 turns 91.7% 8 turns 78.2% 11 turns 63.0% 10 turns 94.0% 90.5%
DeepInfra 97.5% 6 turns 80.2% 8 turns 59.3% 11 turns 38.4% 10 turns 94.0% 88.7%
Z.ai 83.3% 6 turns 81.0% 8 turns 33.6% 11 turns 16.5% 10 turns 93.8% 85.9%
Baseten 69.3% 6 turns 42.6% 8 turns 35.5% 11 turns 0.0% 10 turns 91.6% 80.4%
Modal 49.1% 6 turns 19.7% 8 turns 0.0% 11 turns 0.0% 10 turns 91.9% 78.8%
Fireworks 97.8% 6 turns 81.0% 8 turns 22.3% 11 turns 0.0% 10 turns 96.3% 87.1%
After a 15–20 minute break, Keln still had 63.0% of the prompt cached, compared with 38.4% on DeepInfra and 16.5% on Z.ai. Baseten, Modal and Fireworks had 0%.
Keln had the highest median decode speed on GLM-5.3 Flash at 432 tokens/s, 75% faster than the next fastest provider, Baseten, at 247 tokens/s.
On DeepSeek, Keln and Together were nearly tied at 432 and 437 tokens/s, both well ahead of the other providers.
0 100 200 300 400 500 432 Keln 240 turns 247 Baseten 217 turns 204 Fireworks 240 turns 176 DeepInfra 240 turns 129 Modal 240 turns 46 Z.ai 240 turns
Decode speed, median DeepSeek V4.1 Flash 1017 GLM-5.3 Flash
Keln 432 tok/s first token 0.35 / 0.89 s 432 tok/s first token 1.01 / 2.97 s
Fireworks 234 tok/s first token 1.21 / 2.45 s 204 tok/s first token 0.79 / 1.66 s
Together 437 tok/s first token 0.22 / 0.65 s –
Baseten 208 tok/s first token 0.38 / 1.07 s 247 tok/s first token 0.38 / 1.59 s
Modal 228 tok/s first token 0.72 / 1.45 s 129 tok/s first token 0.99 / 2.48 s
DeepInfra 98 tok/s first token 1.05 / 3.19 s 176 tok/s first token 1.16 / 5.20 s
Z.ai – 46 tok/s first token 4.98 / 27.11 s
First token is p50 / p95 and counts the first reasoning token.
We set reasoning effort to maximum for every request. All providers returned broadly similar amounts of reasoning tokens, including Keln.
In our previous benchmark , two gateways returned less than half as many reasoning tokens for the same requests. We didn't see that issue here.
Modal doesn't report reasoning tokens for GLM-5.3 Flash, so those results aren't included.
Fireworks was the most expensive on DeepSeek, largely due to its 165 cold turns. It processed 2.5 times as many uncached prompt tokens as Keln.
On GLM-5.3 Flash, Keln was among the more expensive options. Its pricing covers more than model inference: Keln manages capacity across providers, monitors performance and output quality, and handles failover and retries. You get all of that through one API, without managing provider accounts, rate limits or caching yourself.
Baseten's cost per 1,000 turns counts only the turns it answered.
Going directly to an inference provider can be cheaper, and performance can be excellent when capacity is available. But you're limited to that provider's infrastructure and your account's rate limits.
We saw those limits in this benchmark. Baseten rejected requests due to rate limits, Modal cut off a response mid-stream, and Fireworks had to process previously cached prompts again.
Keln takes a different approach. It runs inference across multiple providers, managing capacity, performance, and reliability behind one API. In this benchmark, Keln completed every turn, matched or exceeded the fastest direct providers on decode speed, and retained more cached tokens after longer breaks.
Keln wasn't the cheapest option, but the price includes the infrastructure management, failover, and retries you'd otherwise have to handle yourself.
We post benchmarks with the method included. The previous one compared OpenRouter, Vercel, Ramp and Keln on DeepSeek V4.1 Flash for 24 hours: read it here .
We used 8 sessions from the public armand0e Claude Code corpus on Hugging Face, replayed with native tool calls. Each recorded assistant turn was sent as a new request, with the original tool results appended, so conversations grew as they did in the original sessions.
Each run used up to 30 rounds per session, with 2 concurrent sessions per side and fresh salts and conversation IDs for each pass. The median turn contained 40,665 tokens on DeepSeek and 38,578 on GLM-5.3 Flash.
All requests used reasoning_effort: max in the format supported by each API, temperature 1, streaming enabled, and no max_tokens limit.
Pauses between turns followed the original sessions, capped at 60 seconds, with additional breaks to test caching. On DeepSeek, 5% of turns had breaks of 30 seconds to 2 minutes. On GLM-5.3 Flash, 15% had breaks of 1–20 minutes. We used the same random seed across all sides, so each received the same pauses on the same turns.
The DeepSeek benchmark ran on November 3, 2026, from 02:30 to 08:53 UTC, with 8,640 turns total (1,440 per side). The GLM-5.3 Flash run made 3 passes in 14 hours on November 4; a pass replays the 8 sessions once. This post reports the third pass, from 03:45 to 08:32 UTC, with 1,440 turns total (roughly 240 per side).
We tested Keln through its public API. For direct providers, we used Keln's own accounts and each provider's model identifiers, reasoning settings, and per-conversation cache hints. These included the documented user field or, for Baseten, the x-session-affinity header. Fireworks served DeepSeek through its standard endpoint.
Cache hit rate is the percentage of prompt tokens reported as cached in each provider's usage data. A provider switch is counted when consecutive turns in the same conversation run on different providers. Decode speed is calculated as output tokens divided by generation time for each turn.
We calculated costs using each provider's published model prices as of November 4, 2026, applied to its reported usage. Fireworks offers two DeepSeek pricing tiers; we used the standard tier corresponding to the endpoint tested. Modal shows its rates in its console, so we used the prices shown on our account.
For Keln, we used the November 4 price list: DeepSeek at $0.30 per million input tokens, $0.01 per million cached input tokens, and $1.20 per million output tokens; GLM-5.3 Flash at $0.15, $0.05, and $0.50 respectively. The DeepSeek run was billed at the previous cached-input rate of $0.0108 per million tokens, resulting in $2.09 per 1,000 turns.
Results reflect this workload, our account limits, and the conditions during the test. Performance may vary with other workloads or concurrency levels.
The raw benchmark rows and response bodies are not publicly available.
Claude Setup Run open models in the desktop app. No code. Codex Setup Run open models in Codex, terminal or desktop DeepSeek V4.1 Flash 1017 GLM-5.3 Flash Benchmark Prompt caching Pricing Try it yourself Keln Quickstart Send your first request in under 2 minutes Claude Setup Run open models in the desktop app. No code. Codex Setup Run open models in Codex, terminal or desktop Run your agent sessions on Keln
Cheat detection, high cache hit rate and normalized output.
Built by a team from NVIDIA and Google.
© 2026 Keln, Inc. 2261 Market Street STE 32997, San Francisco, CA 94114
