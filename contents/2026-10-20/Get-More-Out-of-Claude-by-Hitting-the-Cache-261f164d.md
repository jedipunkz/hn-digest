---
source: "https://www.dgrcode.com/blog/claude-cache"
hn_url: "https://news.ycombinator.com/item?id=50163131"
title: "Get More Out of Claude by Hitting the Cache"
article_title: "dgrcode blog - Get More Out of Claude by Hitting the Cache"
image: ""
author: "dgrcode"
captured_at: "2026-10-20T09:12:01Z"
capture_tool: "hn-digest"
hn_id: 50163131
score: 3
comments: 0
posted_at: "2026-10-20T08:50:32Z"
tags:
  - hacker-news
---

# Get More Out of Claude by Hitting the Cache

- HN: [50163131](https://news.ycombinator.com/item?id=50163131)
- Source: [www.dgrcode.com](https://www.dgrcode.com/blog/claude-cache)
- Score: 3
- Comments: 0
- Posted: 2026-10-20T08:50:32Z

## Translation

Title: Get More Out of Claude by Hitting the Cache
Article title: dgrcode blog - Get More Out of Claude by Hitting the Cache
Description: Learn how Claude's prompt cache works and how to keep your token usage under control.

Article text:
dgrcode blog - Get More Out of Claude by Hitting the Cache dgr code Projects Blog Get More Out of Claude by Hitting the Cache
Learn how Claude's prompt cache works and how to keep your token usage under control.
You might have noticed that a small prompt sometimes burns a surprising amount of Claude usage. If you sent that prompt to a session that had been idle for a while, there's a good chance you missed the cache. With Opus 5.5, a cache miss can cost up to 40x more than a cache hit. Let's go over what happens and how to avoid it.
Simply put, the cache is temporary storage for the results of your previous prompts . It's a bit more complex than that, because what gets stored is an intermediate result (the KV representation), but we don't need the details to understand how to take advantage of it.
For example, you send a prompt to work on a feature. Your harness (e.g. Claude Code) sends your prompt to Opus, along with additional context: the system prompt, your AGENTS.md content, skills, and so on.
Claude processes the whole prompt, stores the result in the cache, and responds to you.
When you send a second message, your harness sends the whole context again. The good news is that Claude doesn't process it from scratch. Instead, it reuses the result stored in the cache and only processes the new information. The new result is stored in the cache again, ready for your follow-up. Again, this is a simplification with some nuances, but it's a good enough mental model.
The cache is a huge performance improvement, but it takes up storage space, and that space is limited. If Claude hasn't heard back from you after some time, it removes the stored result to free up space for others.
That means a message sent after the cache was cleared has to process the whole context again.
That's what we call a "cache miss". The opposite, when your message can reuse a result from the cache, is a "cache hit".
# The Cost of Missing the Cache
You've seen that every model lists prices per million input and output tokens. They also list a "cache read" and a "cache write" price.
For example, Opus 5.5 costs $4 per million input tokens, but only $0.20 per million for a cache read. That's 20x cheaper if you hit the cache.
Writing to the cache also has an extra cost: Opus 5.5 charges 2x the normal input price for a 1-hour cache write. So the difference between hitting and missing the cache is 40x.
Imagine you're working in a session with 300k tokens of context, and you send a new 1k-token message. There are two scenarios:
Cache hit: 300k tokens read from the cache at 1/20 of the base price, plus 1k tokens written to the cache at 2x the base price. That's 15k + 2k = 17k base input tokens.
Cache miss: 301k tokens written to the cache at 2x the base price. That's 602k base input tokens.
In this realistic example, the cache miss costs over 35x more.
If the previous tokens are P P P and the new tokens are N N N , the two scenarios cost:
Cache hit: P / 20 + 2 N P/20 + 2N P /20 + 2 N
Cache miss: 2 ( P + N ) 2(P + N) 2 ( P + N )
The bigger the existing context, P P P , and the smaller the new message, N N N , the closer a cache miss gets to costing 40x more.
# How to Avoid Missing the Cache
By default, Claude Code uses a 1-hour cache, meaning the cache expires 1 hour after the last write or read. Other harnesses might use a different cache duration.
One strategy is to keep sending messages so the cache never expires. You can do that manually by writing something like
wait, I'm still working on this
That message is around 10 tokens, but it can write over 100 tokens to the cache, because extra data gets sent with it, like the previous response and the user/system markers.
So, using the formula above, the savings compared to a cache miss are close to 40x. With just 100k tokens of context, you'd already save 38.5x.
Note that writing the first hour to the cache costs 2x the base tokens, but hitting the cache adds an extra hour for 0.05x the base tokens. That means if you could automate the process perfectly, a 1h cache costs 2x, a 2h cache costs 2.05x, a 3h cache costs 2.1x, and so on. There's definitely a cost asymmetry here that can be leveraged.
Claude Code has a /loop command that sends a message on a schedule.
This isn't free, but you could loop roughly 40 times before it costs as much as a single cache miss.
One problem is that there's no easy way to express "every 50 minutes" in cron syntax. The simplest alternative is to run the loop at :00 and :30, which gives 40 iterations in 20 hours.
# The Best Alternative to Keeping the Cache Alive
The best alternative is to compact your context with the /compact command. Compacting summarizes the conversation so the context itself gets smaller. The tradeoff is that you might lose parts of the context that later turn out to be relevant.
The cost of compacting depends on your harness, which is the one sending the instruction to compact. For Claude Code, that instruction is about 2k tokens. The whole current context hits the cache*, so it's 20x cheaper than base tokens, and the 2k instruction tokens are new. They could skip being written to the cache, but the docs don't make it clear what happens.
Then Claude responds with the summary, which is output. Output tokens are much more expensive than input tokens. Opus 5.5, for example, charges 5x the base input price. In my experience, the summary is usually under 30k tokens, which costs the equivalent of 150k base input tokens. That's easy to overlook, but it's usually the most expensive part of compacting.
The first message after compaction writes the summary to the cache, along with the system prompt and tools. The system prompt and tools might be another 30k tokens, but this depends a lot on your setup. You can check the size of your context by running /context .
If your compaction misses the cache, that's definitely the worst possible scenario. You'll pay 2x for writing the whole context to the cache, plus the compaction itself, the summary output at 5x the base token price, and then 2x again to write the summary and the system prompt + tools back to the cache. Remember: do not compact when the cache has expired !
Let's go back to the session with 300k tokens of context, and compact it before the cache expires:
Compaction instruction: 300k tokens read from the cache at 1/20 of the base price, plus 2k instruction tokens written to the cache at 2x. That's 15k + 4k = 19k base input tokens.
Summary: 30k output tokens at 5x the base price. That's 150k base input tokens.
First message after compaction: the 30k summary, 30k of system prompt and tools, and the new message (let's assume 1k tokens), all written to the cache at 2x. That's 122k base input tokens.
That's 291k base input tokens in total. It's about half the cost of the 602k cache miss, but much more than the 17k of a cache hit.
The savings come later. Every following message reads around 60k tokens from the cache instead of 300k, so it costs about 5k base input tokens instead of 17k. After enough messages, compacting pays for itself.
You can see how the different strategies depend a lot on a few variables, like the size of the context and the size of the summary. It's time for an interactive comparison.
Here are the three scenarios we discussed: letting the cache expire, looping a wait message, and compacting the context. The chart shows the cumulative base token cost of the pause and the first 10 messages after it.
pause 1 2 3 4 5 6 7 8 9 10 Messages after the pause 0 100k 200k 300k 400k 500k 600k 700k 800k Cumulative cost (base input tokens) Cache expires Loop Compact
The best strategy depends mostly on the context size and how long you'd need to loop. As a rule of thumb: compact when you have a big context or a long pause, and otherwise loop.
What counts as a big context or a long pause depends on the other one. But roughly 600k tokens is a big context, and 8+ hours is a long pause.
Here's a more nuanced decision flow, if you're that kind of person:
If your context is under 100k: you're overcomplicating things. Let the cache expire.
Else, if your context is over 600k: compact. You'd have to compact soon anyway.
Else, if you'll pause for 2h or less: loop.
Else, if you'll pause for 16+ hours: compact.
Else, if you'll pause for 8–16 hours: compact if your context is over 200k, or loop otherwise.
Else, if your context is under 300k: loop.
Else, if you'll pause for 4+ hours: compact.
The 2–4 hour range is a gray area. There, the main questions are:
Will you keep working in that session for much longer? Compacting is probably better.
Do you need just a few more prompts to finish the task? Looping is probably better.
Here's a flow diagram courtesy of Claude:
When I realized that hitting or missing the cache meant a 40x usage difference, I thought I had found hidden treasure with the loop strategy. But comparing /loop with /compact made me re-evaluate when each one makes more sense.
I still think /loop is the best option for many common scenarios, like a long lunch break or running an errand.
Even so, big contexts, or long pauses with a medium-sized context, make compacting a no-brainer. Losing relevant context is a reason to lean towards /loop , but as models get smarter, their summaries get better too, so this problem will slowly go away.
That said, compaction has a hard limit on fidelity. Going from an 800k to a 30k context is over 26x compression, so we're definitely losing relevant pieces there. For reference, research on prompt compression seems to agree that somewhere around 10x to 20x compression is the sweet spot where most relevant information is still present, with a performance drop around 25x. That research compresses the prompt itself rather than summarizing it, but it's a useful reference point.
# Annex: Claude Code Might Be Compacting Your Context Automatically
Just as I was writing this post, I saw Claude Code report:
⏺ Compacted while idle, before the prompt cache expired
While experimenting with context sizes and caches, I had just captured my /context , which was only 220.3k tokens.
Opus 5.5 claude-opus-5-5 220.3k / 1m tokens ( 22 %) Estimated usage by category System prompt : 2.5k tokens (0.3%)
System tools : 12.9k tokens (1.3%)
MCP server instructions : 752 tokens (0.1%)
Memory files : 1.6k tokens (0.2%)
Messages : 192.6k tokens (19.3%)
Autocompact buffer : 33k tokens (3.3%)
It looks like Anthropic agrees that missing the cache is pretty bad, although we disagree on what the best strategy would've been in this scenario, given that I was away for just over 1h.
It would be great if we had a setting to manage cache expiration in Claude Code, so that, depending on the work I'm doing, I could set it to:
Smart Cache (compact or loop, based on the context size and the time away)
The "time away" could be some kind of input we give to the harness.
The difference between 10x and .10x engineers, and how to identify them in your team.
Optimising a strategy to solve Wordle.

## Original Extract

Learn how Claude's prompt cache works and how to keep your token usage under control.

dgrcode blog - Get More Out of Claude by Hitting the Cache dgr code Projects Blog Get More Out of Claude by Hitting the Cache
Learn how Claude's prompt cache works and how to keep your token usage under control.
You might have noticed that a small prompt sometimes burns a surprising amount of Claude usage. If you sent that prompt to a session that had been idle for a while, there's a good chance you missed the cache. With Opus 5.5, a cache miss can cost up to 40x more than a cache hit. Let's go over what happens and how to avoid it.
Simply put, the cache is temporary storage for the results of your previous prompts . It's a bit more complex than that, because what gets stored is an intermediate result (the KV representation), but we don't need the details to understand how to take advantage of it.
For example, you send a prompt to work on a feature. Your harness (e.g. Claude Code) sends your prompt to Opus, along with additional context: the system prompt, your AGENTS.md content, skills, and so on.
Claude processes the whole prompt, stores the result in the cache, and responds to you.
When you send a second message, your harness sends the whole context again. The good news is that Claude doesn't process it from scratch. Instead, it reuses the result stored in the cache and only processes the new information. The new result is stored in the cache again, ready for your follow-up. Again, this is a simplification with some nuances, but it's a good enough mental model.
The cache is a huge performance improvement, but it takes up storage space, and that space is limited. If Claude hasn't heard back from you after some time, it removes the stored result to free up space for others.
That means a message sent after the cache was cleared has to process the whole context again.
That's what we call a "cache miss". The opposite, when your message can reuse a result from the cache, is a "cache hit".
# The Cost of Missing the Cache
You've seen that every model lists prices per million input and output tokens. They also list a "cache read" and a "cache write" price.
For example, Opus 5.5 costs $4 per million input tokens, but only $0.20 per million for a cache read. That's 20x cheaper if you hit the cache.
Writing to the cache also has an extra cost: Opus 5.5 charges 2x the normal input price for a 1-hour cache write. So the difference between hitting and missing the cache is 40x.
Imagine you're working in a session with 300k tokens of context, and you send a new 1k-token message. There are two scenarios:
Cache hit: 300k tokens read from the cache at 1/20 of the base price, plus 1k tokens written to the cache at 2x the base price. That's 15k + 2k = 17k base input tokens.
Cache miss: 301k tokens written to the cache at 2x the base price. That's 602k base input tokens.
In this realistic example, the cache miss costs over 35x more.
If the previous tokens are P P P and the new tokens are N N N , the two scenarios cost:
Cache hit: P / 20 + 2 N P/20 + 2N P /20 + 2 N
Cache miss: 2 ( P + N ) 2(P + N) 2 ( P + N )
The bigger the existing context, P P P , and the smaller the new message, N N N , the closer a cache miss gets to costing 40x more.
# How to Avoid Missing the Cache
By default, Claude Code uses a 1-hour cache, meaning the cache expires 1 hour after the last write or read. Other harnesses might use a different cache duration.
One strategy is to keep sending messages so the cache never expires. You can do that manually by writing something like
wait, I'm still working on this
That message is around 10 tokens, but it can write over 100 tokens to the cache, because extra data gets sent with it, like the previous response and the user/system markers.
So, using the formula above, the savings compared to a cache miss are close to 40x. With just 100k tokens of context, you'd already save 38.5x.
Note that writing the first hour to the cache costs 2x the base tokens, but hitting the cache adds an extra hour for 0.05x the base tokens. That means if you could automate the process perfectly, a 1h cache costs 2x, a 2h cache costs 2.05x, a 3h cache costs 2.1x, and so on. There's definitely a cost asymmetry here that can be leveraged.
Claude Code has a /loop command that sends a message on a schedule.
This isn't free, but you could loop roughly 40 times before it costs as much as a single cache miss.
One problem is that there's no easy way to express "every 50 minutes" in cron syntax. The simplest alternative is to run the loop at :00 and :30, which gives 40 iterations in 20 hours.
# The Best Alternative to Keeping the Cache Alive
The best alternative is to compact your context with the /compact command. Compacting summarizes the conversation so the context itself gets smaller. The tradeoff is that you might lose parts of the context that later turn out to be relevant.
The cost of compacting depends on your harness, which is the one sending the instruction to compact. For Claude Code, that instruction is about 2k tokens. The whole current context hits the cache*, so it's 20x cheaper than base tokens, and the 2k instruction tokens are new. They could skip being written to the cache, but the docs don't make it clear what happens.
Then Claude responds with the summary, which is output. Output tokens are much more expensive than input tokens. Opus 5.5, for example, charges 5x the base input price. In my experience, the summary is usually under 30k tokens, which costs the equivalent of 150k base input tokens. That's easy to overlook, but it's usually the most expensive part of compacting.
The first message after compaction writes the summary to the cache, along with the system prompt and tools. The system prompt and tools might be another 30k tokens, but this depends a lot on your setup. You can check the size of your context by running /context .
If your compaction misses the cache, that's definitely the worst possible scenario. You'll pay 2x for writing the whole context to the cache, plus the compaction itself, the summary output at 5x the base token price, and then 2x again to write the summary and the system prompt + tools back to the cache. Remember: do not compact when the cache has expired !
Let's go back to the session with 300k tokens of context, and compact it before the cache expires:
Compaction instruction: 300k tokens read from the cache at 1/20 of the base price, plus 2k instruction tokens written to the cache at 2x. That's 15k + 4k = 19k base input tokens.
Summary: 30k output tokens at 5x the base price. That's 150k base input tokens.
First message after compaction: the 30k summary, 30k of system prompt and tools, and the new message (let's assume 1k tokens), all written to the cache at 2x. That's 122k base input tokens.
That's 291k base input tokens in total. It's about half the cost of the 602k cache miss, but much more than the 17k of a cache hit.
The savings come later. Every following message reads around 60k tokens from the cache instead of 300k, so it costs about 5k base input tokens instead of 17k. After enough messages, compacting pays for itself.
You can see how the different strategies depend a lot on a few variables, like the size of the context and the size of the summary. It's time for an interactive comparison.
Here are the three scenarios we discussed: letting the cache expire, looping a wait message, and compacting the context. The chart shows the cumulative base token cost of the pause and the first 10 messages after it.
pause 1 2 3 4 5 6 7 8 9 10 Messages after the pause 0 100k 200k 300k 400k 500k 600k 700k 800k Cumulative cost (base input tokens) Cache expires Loop Compact
The best strategy depends mostly on the context size and how long you'd need to loop. As a rule of thumb: compact when you have a big context or a long pause, and otherwise loop.
What counts as a big context or a long pause depends on the other one. But roughly 600k tokens is a big context, and 8+ hours is a long pause.
Here's a more nuanced decision flow, if you're that kind of person:
If your context is under 100k: you're overcomplicating things. Let the cache expire.
Else, if your context is over 600k: compact. You'd have to compact soon anyway.
Else, if you'll pause for 2h or less: loop.
Else, if you'll pause for 16+ hours: compact.
Else, if you'll pause for 8–16 hours: compact if your context is over 200k, or loop otherwise.
Else, if your context is under 300k: loop.
Else, if you'll pause for 4+ hours: compact.
The 2–4 hour range is a gray area. There, the main questions are:
Will you keep working in that session for much longer? Compacting is probably better.
Do you need just a few more prompts to finish the task? Looping is probably better.
Here's a flow diagram courtesy of Claude:
When I realized that hitting or missing the cache meant a 40x usage difference, I thought I had found hidden treasure with the loop strategy. But comparing /loop with /compact made me re-evaluate when each one makes more sense.
I still think /loop is the best option for many common scenarios, like a long lunch break or running an errand.
Even so, big contexts, or long pauses with a medium-sized context, make compacting a no-brainer. Losing relevant context is a reason to lean towards /loop , but as models get smarter, their summaries get better too, so this problem will slowly go away.
That said, compaction has a hard limit on fidelity. Going from an 800k to a 30k context is over 26x compression, so we're definitely losing relevant pieces there. For reference, research on prompt compression seems to agree that somewhere around 10x to 20x compression is the sweet spot where most relevant information is still present, with a performance drop around 25x. That research compresses the prompt itself rather than summarizing it, but it's a useful reference point.
# Annex: Claude Code Might Be Compacting Your Context Automatically
Just as I was writing this post, I saw Claude Code report:
⏺ Compacted while idle, before the prompt cache expired
While experimenting with context sizes and caches, I had just captured my /context , which was only 220.3k tokens.
Opus 5.5 claude-opus-5-5 220.3k / 1m tokens ( 22 %) Estimated usage by category System prompt : 2.5k tokens (0.3%)
System tools : 12.9k tokens (1.3%)
MCP server instructions : 752 tokens (0.1%)
Memory files : 1.6k tokens (0.2%)
Messages : 192.6k tokens (19.3%)
Autocompact buffer : 33k tokens (3.3%)
It looks like Anthropic agrees that missing the cache is pretty bad, although we disagree on what the best strategy would've been in this scenario, given that I was away for just over 1h.
It would be great if we had a setting to manage cache expiration in Claude Code, so that, depending on the work I'm doing, I could set it to:
Smart Cache (compact or loop, based on the context size and the time away)
The "time away" could be some kind of input we give to the harness.
The difference between 10x and .10x engineers, and how to identify them in your team.
Optimising a strategy to solve Wordle.
