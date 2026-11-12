---
source: ""
hn_url: "https://news.ycombinator.com/item?id=50467233"
title: "Claude Code's API billing depends on how often the cache expires"
article_title: ""
image: ""
author: "MaxBozhenko"
captured_at: "2026-11-12T15:01:47Z"
capture_tool: "hn-digest"
hn_id: 50467233
score: 2
comments: 0
posted_at: "2026-11-12T14:02:38Z"
tags:
  - hacker-news
---

# Claude Code's API billing depends on how often the cache expires

- HN: [50467233](https://news.ycombinator.com/item?id=50467233)
- Score: 2
- Comments: 0
- Posted: 2026-11-12T14:02:38Z

## Translation

Title: Claude Code's API billing depends on how often the cache expires
HN text: Hello, HN! I’m Max, CTO at Hystax. We use Claude a lot in our development work, and we’re also trying to make coding agents cheaper to run. We built the open-source Kompressor for compressing identical tool outputs that agents constantly resend. In our Codex tests, we noticed a real cost reduction of 22.7%. But testing it with Claude Code, we found something completely different. In Claude Code, the cache plays a more crucial role than compression. Claude Code marks its messages with caching points. As a result, almost every request consists of a long cached prefix and a new, truncated final part. Based on practical experience with Opus 5.5, reading from the cache appears to cost $0.20 per million tokens, accounting for about 92–95% of the tokens charged for prompts. Meanwhile, a five-minute cache entry costs $5, which is 25 times more expensive than reading from the cache. When an agent needs to run a build or a longer training cycle, they face the problem of the cache expiring. The next request then has to write the prefix to the cache again. We saw this 33 times in a cycle of 89-task runs without compression. As a result, 927,000 tokens were rewritten, which is approximately one-third of the volume billed in full. Until now, we’ve deliberately left the Claude Code text uncompressed because overwriting any byte within a cached prefix would very likely cause it to become invalid. But here’s what we came up with instead: – Compress only the uncached part. On the first transmission, only the new part is compressed. Everything already cached stays byte-for-byte unchanged. Fragments shorter than 4,000 characters are sent in full, since compressing each fragment would be much more expensive than sending the entire dialogue as is. – Keep the cache active while the tools are running; don’t let it expire.
While the agent waits for a tool response (model building/training), Kompressor resends the last request every 4.5 minutes with max_tokens: 1, so Anthropic re-reads the cached dialogue, and the five-minute timer restarts. Pings stop on the agent’s next request, after ~50 min, on the first failed attempt, or if a ping finds nothing in the cache. – Switch to the one-hour cache only when necessary.
After the first pause longer than 5 min, Kompressor requests a one-hour cache for that conversation. We ran Claude Code on the Opus 5.5 model through the Terminal-Bench 2.1 benchmark (89 tasks) and calculated the cost based on Anthropic's actual billing, including pings. The cost per run decreased from $35.35 to $28.30 (-20%). The percentage of tasks solved by the agent without compression was 89.9%, and 90.6% when using Kompressor. The number of tokens lost due to cache expiration decreased from 927,000 to 13,000. But there’s a caveat: the agent never solves the same task twice in exactly the same way, and the entire difference came down to just 4 out of the 89 tasks. We expect these savings to persist from run to run: about 11% from keeping the cache up to date, and a few more percentage points from compression. This effect is most noticeable on tasks where the agent waits a long time – more than five minutes – for a command to execute: model training, compilation, and long simulations. In those cases, costs decreased by 34%.
To try Kompressor, see the GitHub repo below. Setup is in the README. The default configuration already includes compression of new text only and a switch to an hourly cache on long pauses; “keep-warm” is off (the only request that the proxy sends on its own). To enable, add to your config: [proxy.providers.anthropic]
cache_keep_warm = true
and then run: kompressor wrap claude
The benchmark above was run with all these settings enabled. Here are cost calculations: https://optscale.ai/blog/claude-code-prompt-caching
If you run long tasks in Claude Code, how often does the cache expire while a tool is still running? Have you seen this with other providers too?

## Original Extract

Hello, HN! I’m Max, CTO at Hystax. We use Claude a lot in our development work, and we’re also trying to make coding agents cheaper to run. We built the open-source Kompressor for compressing identical tool outputs that agents constantly resend. In our Codex tests, we noticed a real cost reduction of 22.7%. But testing it with Claude Code, we found something completely different. In Claude Code, the cache plays a more crucial role than compression. Claude Code marks its messages with caching points. As a result, almost every request consists of a long cached prefix and a new, truncated final part. Based on practical experience with Opus 5.5, reading from the cache appears to cost $0.20 per million tokens, accounting for about 92–95% of the tokens charged for prompts. Meanwhile, a five-minute cache entry costs $5, which is 25 times more expensive than reading from the cache. When an agent needs to run a build or a longer training cycle, they face the problem of the cache expiring. The next request then has to write the prefix to the cache again. We saw this 33 times in a cycle of 89-task runs without compression. As a result, 927,000 tokens were rewritten, which is approximately one-third of the volume billed in full. Until now, we’ve deliberately left the Claude Code text uncompressed because overwriting any byte within a cached prefix would very likely cause it to become invalid. But here’s what we came up with instead: – Compress only the uncached part. On the first transmission, only the new part is compressed. Everything already cached stays byte-for-byte unchanged. Fragments shorter than 4,000 characters are sent in full, since compressing each fragment would be much more expensive than sending the entire dialogue as is. – Keep the cache active while the tools are running; don’t let it expire.
While the agent waits for a tool response (model building/training), Kompressor resends the last request every 4.5 minutes with max_tokens: 1, so Anthropic re-reads the cached dialogue, and the five-minute timer restarts. Pings stop on the agent’s next request, after ~50 min, on the first failed attempt, or if a ping finds nothing in the cache. – Switch to the one-hour cache only when necessary.
After the first pause longer than 5 min, Kompressor requests a one-hour cache for that conversation. We ran Claude Code on the Opus 5.5 model through the Terminal-Bench 2.1 benchmark (89 tasks) and calculated the cost based on Anthropic's actual billing, including pings. The cost per run decreased from $35.35 to $28.30 (-20%). The percentage of tasks solved by the agent without compression was 89.9%, and 90.6% when using Kompressor. The number of tokens lost due to cache expiration decreased from 927,000 to 13,000. But there’s a caveat: the agent never solves the same task twice in exactly the same way, and the entire difference came down to just 4 out of the 89 tasks. We expect these savings to persist from run to run: about 11% from keeping the cache up to date, and a few more percentage points from compression. This effect is most noticeable on tasks where the agent waits a long time – more than five minutes – for a command to execute: model training, compilation, and long simulations. In those cases, costs decreased by 34%.
To try Kompressor, see the GitHub repo below. Setup is in the README. The default configuration already includes compression of new text only and a switch to an hourly cache on long pauses; “keep-warm” is off (the only request that the proxy sends on its own). To enable, add to your config: [proxy.providers.anthropic]
cache_keep_warm = true
and then run: kompressor wrap claude
The benchmark above was run with all these settings enabled. Here are cost calculations: https://optscale.ai/blog/claude-code-prompt-caching
If you run long tasks in Claude Code, how often does the cache expire while a tool is still running? Have you seen this with other providers too?

