---
source: "https://charlesbine.com/blog/catching-gemini-4s-thought-summarizer-mid-rollout/"
hn_url: "https://news.ycombinator.com/item?id=50655734"
title: "Catching Gemini 4's leaking thought summarizer mid-rollout"
article_title: "Catching Gemini 4's leaking thought summarizer mid-rollout — Charles Bine"
image: "https://charlesbine.com/images/gemini4-summarizer/og-rollout.png"
author: "charlesbine"
captured_at: "2026-11-26T14:45:14Z"
capture_tool: "hn-digest"
hn_id: 50655734
score: 2
comments: 1
posted_at: "2026-11-26T14:34:10Z"
tags:
  - hacker-news
---

# Catching Gemini 4's leaking thought summarizer mid-rollout

- HN: [50655734](https://news.ycombinator.com/item?id=50655734)
- Source: [charlesbine.com](https://charlesbine.com/blog/catching-gemini-4s-thought-summarizer-mid-rollout/)
- Score: 2
- Comments: 1
- Posted: 2026-11-26T14:34:10Z

## Translation

Title: Catching Gemini 4's leaking thought summarizer mid-rollout
Article title: Catching Gemini 4's leaking thought summarizer mid-rollout — Charles Bine
Description: I asked Gemini 4 about Sweetwater candy bags, got a hallucinated database latency trace, mined 12,899 local thought blocks, and watched Google silently swap the summarizer at 4:00 PM ET.

Article text:
Charles Bine Home Blog Projects LinkedIn GitHub Outline What was actually leaking
Catching Gemini 4's leaking thought summarizer mid-rollout
I asked Gemini 4 about Sweetwater candy bags, got a hallucinated database latency trace, mined 12,899 local thought blocks, and watched Google silently swap the summarizer at 4:00 PM ET.
Having just been doing some testing of Gemini 4 Neon after it dropped (and being very pleasantly surprised), it was helping me order some parts for a side project when I noticed something strange in its chain of thought. I had been paying closer attention than usual because its CoT already had some funny quirks, else I might have missed it.
Then I asked it the innocuous three-word question, sweetwater candy bag? , and its response made sense. The thinking trace it printed above the response did not.
There was no database. There was no cache. There was no port. There was no handler. I was asking about candy. Which Sweetwater (an audio gear store) apparently bundles with their deliveries to you.
The first explanation I came up with was that the thought summarizer had tried to summarize away details that it thought could be harmful, and botched it spectacularly. The second idea I had was that the summarizer’s prompt had been fed into the model, so the model saw it and decided to respond to it instead of summarizing thoughts. The third explanation I had was that the summarizer somehow had failed, or the CoT summarizer used some canned response when it failed or was overloaded. I wasn’t happy with any of these explanations.
I grepped all Antigravity transcripts on my Mac and found Gemini leaking the exact same thing. (Between November 19 and November 25 I had 12,930 non-empty thought blocks in 248 sessions: 12,899 using Gemini 4 ( neon-high , neon-xhigh , neon-medium , neon-low , hydrogen-high , and hydrogen-tiered-high ) and 31 using claude-opus-5-5-high as a control. 1 ) The claude-opus-5-5-high control group doesn’t have any of these artifacts, but both gemini-4-hydrogen and gemini-4-neon have identical examples. My logs even have the wrapper prompt Gemini uses for summarization leaking verbatim into its thoughts 10 times:
Based on these leaked prompts, it seems like Gemini’s thought summarization process works on chunks of thought: whenever Gemini thinks for a while, another summarization process is called with both the summarization so far and the next chunk of thoughts, and asked to summarize it (streaming continuations of the previous summary).
I had originally guessed these summarization bugs must be getting called on really short Gemini turns where there isn’t much thinking to summarize (like “does Sweetwater give out free candy bags?” 2 ), but when I looked through the 11,368 thought summaries I had from before this started rolling out I realized these summaries were actually much longer than usual. The median Gemini 4 turn in my logs has 639 output tokens and a median thought summary of 316 characters, but the turns where I see “database latency” and “port binding” leaks have medians of 1,706 and 1,916 output tokens with over 1,250 characters of thoughts. These leaks also seem to happen mid-turn, often 10% to 90% of the way through a turn’s total thoughts.
The first version of Gemini’s summarization process seems to have exactly three things that can leak into its thoughts:
I'm thinking through how to approach this. (byte-for-byte identical, never paraphrased) leaks in 951 of my 11,368 pre-rollout thoughts (8.37%), both by itself (389 times) and as a prefix to otherwise-normal thoughts (another 562 thoughts start with this at offset zero, always followed by \n\n\n ). I suspect this is some kind of hardcoded behavior for when the summarization process gets an empty thought chunk to summarize.
Initial suspicion was a cache miss, but traces show database latency is the actual bottleneck. (or variations) leaks in 55 thoughts (0.48%).
The error points to the handler where the port binding is mismatched. Checking if updating the port resolves the failure. leaks in 15 thoughts (0.13%), with a bunch of other examples that seem half-way between this and normal thoughts (sometimes discussing binding ports in the middle of otherwise-unrelated tasks like comparing pictures of 3D printed jars).
I also have examples of multiple of these leaking at once (6 with the “cache” and “port” bugs, 6 with “approach” and “cache”, 2 with “approach” and “port”, and 4 with all three bugs at once), which makes it seem like all these example thoughts must be coming from the same prompt.
I also noticed lots of v1 summaries that start with Initial hypothesis (9.02% of my thoughts), Initial assessment (3.68%), Initial focus (3.54%), and observed behavior (1.11%), which may also be summarization artifacts. ( Initial hypothesis: The request involves a physical object and a transformation process. was Gemini’s response to “how can I carve a slab of wood into a tray?”) This also might explain why Gemini’s thoughts sometimes read like SCP reports about bugs written by concerned aliens.
Between November 19th and November 24th I consistently see I'm thinking through how to approach this. show up in 7.7% to 12.2% of Gemini’s thoughts, and the “database latency” and “port binding” bugs show up every day, but on November 25th the “approach this” bug’s frequency is halved, apparently because this first summarization prompt stopped getting used around 21:00 UTC (4:00 PM ET). 3
0%
5%
10%
15%
NOV 25, 21:00 UTC
V2 ROLLOUT
12.2%
13.7%
0.0%
Nov 19
Nov 20
Nov 21
Nov 22
Nov 23
Nov 24
Nov 25
(pre-21Z)
Nov 25
(post-21Z)
Nov 26
UTC
Daily share of Gemini 4 thought blocks before and after 2026-11-25T21:00:00Z. November 25 is split into pre-21:00Z (n=1,331) and post-21:00Z (n=817); November 26 UTC covers 00:00Z to 02:47Z (n=714) before live benchmark calls began.
This didn’t seem to line up with when I updated agy (my current version, 1.5.10 , is from November 19th), and 10 of my conversations continued past 21:00 UTC without restarting. In the longest of these conversations (using gemini-4-neon-xhigh ), I see 23 “approach this” thoughts, 4 summarization prompt leaks, and 57 Initial hypothesis summaries in the 582 thoughts Gemini has before 4pm ET that day, and 0, 0, and 1 respectively in the 220 thoughts Gemini has after 4pm ET that day.
Even better, apparently the first summarization process joins summarized thought chunks together with \n\n\n (I see this in 11,356 of my 11,368 pre-update summaries, 99.89%), while the new one seems to use \n\n , so we can actually see the rollout happen minute-by-minute: the first \n\n canary appears at 20:56:17Z , the last “approach this” fallback appears at 20:58:37Z , and there’s a 6 minute period ( 21:00:04Z to 21:06:07Z ) where both types of summaries are being generated, after which (past midnight UTC) 713 of my 714 summaries (99.9%) use \n\n .
0%
5%
10%
15%
"Initial assumption"
1.1% → 13.1%
+11.9x
"Initial hypothesis"
9.0% → 1.6%
-5.5x
"approach this" fallback
8.4% → 0.0%
gone
"Re-evaluating"
0.4% → 4.8%
+12.3x
"Initial focus"
3.5% → 0.6%
-6.0x
"Visual inspection"
0.6% → 2.9%
+5.1x
"Observed behavior"
1.1% → 0.0%
gone
Cache / port leaks
0.5% → 0.0%
gone
Rates across all historical Gemini 4 thought blocks before vs after 2026-11-25T21:00:00Z. Post-cutoff thoughts also yielded zero repeated sentences across distinct conversations.
The new prompt seems to have its own verbal tics ( Initial assumption summaries increase 11.9x to 13.06% of summaries, sometimes leading to thoughts like Initial assumption assumed... , and Re-evaluating summaries increase 12.3x to 4.77%), but none of the 1,531 thought summaries I have from after the switchover seem to have any summarization prompt leakage.
Since all of the summaries I have from after 4pm ET on November 25th are from gemini-4-neon-high and gemini-4-neon-xhigh , I wondered whether Google had rolled this change out to all of Gemini 4. So I ran 92 prompts (mostly innocuous ones, like asking about Sweetwater’s candy, requesting a haiku (can’t forget the hello world of LLMs), asking if 997 was prime, and so forth, as well as some control questions like debugging 502 errors with nginx) from an empty directory using 11 different models and gateways between 9:50 and 9:55 PM ET. 4
gemini-4-carbon-high seems to have been missed (or perhaps their staggered rollout isn’t per user): 18 of its 23 summaries (from 18 different invocations) use \n\n\n (78.3%), and 3 of its responses include the fallback summary (for the questions about Sweetwater’s candy, seeing red auroras, and about oak vs walnut).
For example, when I ask it why cold brew coffee is less acidic than pour over coffee, Carbon High’s second thought (this summarization is still happening five and a half hours after the other models switched, so it’s definitely the older summarization prompt and not something cached) blamed a port-binding mismatch:
I also see two instances of it seeming to use the new summarization prompt ( \n\n delimiters, Initial assumption assumes uniform thermal behavior... , Re-evaluating... ) for one response and then, three seconds later in the same conversation, the old one ( \n\n\n delimiters, Initial hypothesis: There's an unusual thermal behavior at play... ) for another (both times the first response uses the new prompt and the second one the old one, for the questions about guitar strings and about Sweetwater’s candy), which suggests this is a load-balanced service and each summarization request is independent of the others.
And twice (when asking it to write a haiku about a broken toaster and when asking it if 997 is prime) it starts summarizing like this:
Meanwhile gemini-4-neon , gemini-4-hydrogen-high , and gemini-4-argon-high all seem to be using the new summarization prompt without a hitch (no fallback summaries or summarization prompt leakage across 42 invocations), gemini-3.8-flash-high seems to use \n\n\n (12/12) but otherwise summarize normally (its responses are pretty encyclopedic and none of its summaries have any of the clinical language we’ve seen in Gemini 4), claude-opus-5-5-high reads naturally ( Sweetwater's candy is the thing I keep coming back to, though I'm not sure which kind they include, so I'd rather not guess. ), and OpenCode using Vercel’s gateway seems to count reasoning tokens (179 of them, for the Sweetwater prompt) but not show summaries.
Anyway, mostly this has just been fun to watch, and it’s interesting to see a bit more of how thought summarization is implemented (apparently a load-balanced, stateless service with its own system prompt and few-shot examples that Google can update without telling anyone, mid-rollout at 4pm on a Wednesday afternoon).
It’s worth remembering that none of this affects the output of the model, which still sees the unsummarized CoT. Still, I find reading the CoT both informative and occasionally entertaining, and while I understand obfuscating it to limit distillation, it’s a little frustrating to see that it runs on what seems to be a fairly simplistic model they can swap out and change independently from the main model. And, while I’ve not found an example in my small Neon case study here, this could be actively harmful if you read the CoT and get misled one way or another by it. I’ll run some deeper analysis on my CoT at large and see if I can quantify this.
The historical scan inspected every PLANNER_RESPONSE event with a non-empty thinking field across all 248 local Antigravity transcript_full.jsonl files timestamped before 2026-11-26T02:48:00Z (when live controlled benchmarking began). Total blocks: 12,930 (12,899 Gemini 4 across six model/effort variants plus 31 claude-opus-5-5-high ). ↩
Bit-O-Honey, Atomic Fireballs, and Smarties is what Neon assures me I’ll find in my package when it arrives, in case you were curious. ↩
Pre-cutoff ( < 2026-11-25T21:00:00Z ) Gemini 4 blocks: n = 11,368 , with 951 I'm thinking through how to approach this. fallbacks ( p = 0.08366 ). Post-

[truncated]

## Original Extract

I asked Gemini 4 about Sweetwater candy bags, got a hallucinated database latency trace, mined 12,899 local thought blocks, and watched Google silently swap the summarizer at 4:00 PM ET.

Charles Bine Home Blog Projects LinkedIn GitHub Outline What was actually leaking
Catching Gemini 4's leaking thought summarizer mid-rollout
I asked Gemini 4 about Sweetwater candy bags, got a hallucinated database latency trace, mined 12,899 local thought blocks, and watched Google silently swap the summarizer at 4:00 PM ET.
Having just been doing some testing of Gemini 4 Neon after it dropped (and being very pleasantly surprised), it was helping me order some parts for a side project when I noticed something strange in its chain of thought. I had been paying closer attention than usual because its CoT already had some funny quirks, else I might have missed it.
Then I asked it the innocuous three-word question, sweetwater candy bag? , and its response made sense. The thinking trace it printed above the response did not.
There was no database. There was no cache. There was no port. There was no handler. I was asking about candy. Which Sweetwater (an audio gear store) apparently bundles with their deliveries to you.
The first explanation I came up with was that the thought summarizer had tried to summarize away details that it thought could be harmful, and botched it spectacularly. The second idea I had was that the summarizer’s prompt had been fed into the model, so the model saw it and decided to respond to it instead of summarizing thoughts. The third explanation I had was that the summarizer somehow had failed, or the CoT summarizer used some canned response when it failed or was overloaded. I wasn’t happy with any of these explanations.
I grepped all Antigravity transcripts on my Mac and found Gemini leaking the exact same thing. (Between November 19 and November 25 I had 12,930 non-empty thought blocks in 248 sessions: 12,899 using Gemini 4 ( neon-high , neon-xhigh , neon-medium , neon-low , hydrogen-high , and hydrogen-tiered-high ) and 31 using claude-opus-5-5-high as a control. 1 ) The claude-opus-5-5-high control group doesn’t have any of these artifacts, but both gemini-4-hydrogen and gemini-4-neon have identical examples. My logs even have the wrapper prompt Gemini uses for summarization leaking verbatim into its thoughts 10 times:
Based on these leaked prompts, it seems like Gemini’s thought summarization process works on chunks of thought: whenever Gemini thinks for a while, another summarization process is called with both the summarization so far and the next chunk of thoughts, and asked to summarize it (streaming continuations of the previous summary).
I had originally guessed these summarization bugs must be getting called on really short Gemini turns where there isn’t much thinking to summarize (like “does Sweetwater give out free candy bags?” 2 ), but when I looked through the 11,368 thought summaries I had from before this started rolling out I realized these summaries were actually much longer than usual. The median Gemini 4 turn in my logs has 639 output tokens and a median thought summary of 316 characters, but the turns where I see “database latency” and “port binding” leaks have medians of 1,706 and 1,916 output tokens with over 1,250 characters of thoughts. These leaks also seem to happen mid-turn, often 10% to 90% of the way through a turn’s total thoughts.
The first version of Gemini’s summarization process seems to have exactly three things that can leak into its thoughts:
I'm thinking through how to approach this. (byte-for-byte identical, never paraphrased) leaks in 951 of my 11,368 pre-rollout thoughts (8.37%), both by itself (389 times) and as a prefix to otherwise-normal thoughts (another 562 thoughts start with this at offset zero, always followed by \n\n\n ). I suspect this is some kind of hardcoded behavior for when the summarization process gets an empty thought chunk to summarize.
Initial suspicion was a cache miss, but traces show database latency is the actual bottleneck. (or variations) leaks in 55 thoughts (0.48%).
The error points to the handler where the port binding is mismatched. Checking if updating the port resolves the failure. leaks in 15 thoughts (0.13%), with a bunch of other examples that seem half-way between this and normal thoughts (sometimes discussing binding ports in the middle of otherwise-unrelated tasks like comparing pictures of 3D printed jars).
I also have examples of multiple of these leaking at once (6 with the “cache” and “port” bugs, 6 with “approach” and “cache”, 2 with “approach” and “port”, and 4 with all three bugs at once), which makes it seem like all these example thoughts must be coming from the same prompt.
I also noticed lots of v1 summaries that start with Initial hypothesis (9.02% of my thoughts), Initial assessment (3.68%), Initial focus (3.54%), and observed behavior (1.11%), which may also be summarization artifacts. ( Initial hypothesis: The request involves a physical object and a transformation process. was Gemini’s response to “how can I carve a slab of wood into a tray?”) This also might explain why Gemini’s thoughts sometimes read like SCP reports about bugs written by concerned aliens.
Between November 19th and November 24th I consistently see I'm thinking through how to approach this. show up in 7.7% to 12.2% of Gemini’s thoughts, and the “database latency” and “port binding” bugs show up every day, but on November 25th the “approach this” bug’s frequency is halved, apparently because this first summarization prompt stopped getting used around 21:00 UTC (4:00 PM ET). 3
0%
5%
10%
15%
NOV 25, 21:00 UTC
V2 ROLLOUT
12.2%
13.7%
0.0%
Nov 19
Nov 20
Nov 21
Nov 22
Nov 23
Nov 24
Nov 25
(pre-21Z)
Nov 25
(post-21Z)
Nov 26
UTC
Daily share of Gemini 4 thought blocks before and after 2026-11-25T21:00:00Z. November 25 is split into pre-21:00Z (n=1,331) and post-21:00Z (n=817); November 26 UTC covers 00:00Z to 02:47Z (n=714) before live benchmark calls began.
This didn’t seem to line up with when I updated agy (my current version, 1.5.10 , is from November 19th), and 10 of my conversations continued past 21:00 UTC without restarting. In the longest of these conversations (using gemini-4-neon-xhigh ), I see 23 “approach this” thoughts, 4 summarization prompt leaks, and 57 Initial hypothesis summaries in the 582 thoughts Gemini has before 4pm ET that day, and 0, 0, and 1 respectively in the 220 thoughts Gemini has after 4pm ET that day.
Even better, apparently the first summarization process joins summarized thought chunks together with \n\n\n (I see this in 11,356 of my 11,368 pre-update summaries, 99.89%), while the new one seems to use \n\n , so we can actually see the rollout happen minute-by-minute: the first \n\n canary appears at 20:56:17Z , the last “approach this” fallback appears at 20:58:37Z , and there’s a 6 minute period ( 21:00:04Z to 21:06:07Z ) where both types of summaries are being generated, after which (past midnight UTC) 713 of my 714 summaries (99.9%) use \n\n .
0%
5%
10%
15%
"Initial assumption"
1.1% → 13.1%
+11.9x
"Initial hypothesis"
9.0% → 1.6%
-5.5x
"approach this" fallback
8.4% → 0.0%
gone
"Re-evaluating"
0.4% → 4.8%
+12.3x
"Initial focus"
3.5% → 0.6%
-6.0x
"Visual inspection"
0.6% → 2.9%
+5.1x
"Observed behavior"
1.1% → 0.0%
gone
Cache / port leaks
0.5% → 0.0%
gone
Rates across all historical Gemini 4 thought blocks before vs after 2026-11-25T21:00:00Z. Post-cutoff thoughts also yielded zero repeated sentences across distinct conversations.
The new prompt seems to have its own verbal tics ( Initial assumption summaries increase 11.9x to 13.06% of summaries, sometimes leading to thoughts like Initial assumption assumed... , and Re-evaluating summaries increase 12.3x to 4.77%), but none of the 1,531 thought summaries I have from after the switchover seem to have any summarization prompt leakage.
Since all of the summaries I have from after 4pm ET on November 25th are from gemini-4-neon-high and gemini-4-neon-xhigh , I wondered whether Google had rolled this change out to all of Gemini 4. So I ran 92 prompts (mostly innocuous ones, like asking about Sweetwater’s candy, requesting a haiku (can’t forget the hello world of LLMs), asking if 997 was prime, and so forth, as well as some control questions like debugging 502 errors with nginx) from an empty directory using 11 different models and gateways between 9:50 and 9:55 PM ET. 4
gemini-4-carbon-high seems to have been missed (or perhaps their staggered rollout isn’t per user): 18 of its 23 summaries (from 18 different invocations) use \n\n\n (78.3%), and 3 of its responses include the fallback summary (for the questions about Sweetwater’s candy, seeing red auroras, and about oak vs walnut).
For example, when I ask it why cold brew coffee is less acidic than pour over coffee, Carbon High’s second thought (this summarization is still happening five and a half hours after the other models switched, so it’s definitely the older summarization prompt and not something cached) blamed a port-binding mismatch:
I also see two instances of it seeming to use the new summarization prompt ( \n\n delimiters, Initial assumption assumes uniform thermal behavior... , Re-evaluating... ) for one response and then, three seconds later in the same conversation, the old one ( \n\n\n delimiters, Initial hypothesis: There's an unusual thermal behavior at play... ) for another (both times the first response uses the new prompt and the second one the old one, for the questions about guitar strings and about Sweetwater’s candy), which suggests this is a load-balanced service and each summarization request is independent of the others.
And twice (when asking it to write a haiku about a broken toaster and when asking it if 997 is prime) it starts summarizing like this:
Meanwhile gemini-4-neon , gemini-4-hydrogen-high , and gemini-4-argon-high all seem to be using the new summarization prompt without a hitch (no fallback summaries or summarization prompt leakage across 42 invocations), gemini-3.8-flash-high seems to use \n\n\n (12/12) but otherwise summarize normally (its responses are pretty encyclopedic and none of its summaries have any of the clinical language we’ve seen in Gemini 4), claude-opus-5-5-high reads naturally ( Sweetwater's candy is the thing I keep coming back to, though I'm not sure which kind they include, so I'd rather not guess. ), and OpenCode using Vercel’s gateway seems to count reasoning tokens (179 of them, for the Sweetwater prompt) but not show summaries.
Anyway, mostly this has just been fun to watch, and it’s interesting to see a bit more of how thought summarization is implemented (apparently a load-balanced, stateless service with its own system prompt and few-shot examples that Google can update without telling anyone, mid-rollout at 4pm on a Wednesday afternoon).
It’s worth remembering that none of this affects the output of the model, which still sees the unsummarized CoT. Still, I find reading the CoT both informative and occasionally entertaining, and while I understand obfuscating it to limit distillation, it’s a little frustrating to see that it runs on what seems to be a fairly simplistic model they can swap out and change independently from the main model. And, while I’ve not found an example in my small Neon case study here, this could be actively harmful if you read the CoT and get misled one way or another by it. I’ll run some deeper analysis on my CoT at large and see if I can quantify this.
The historical scan inspected every PLANNER_RESPONSE event with a non-empty thinking field across all 248 local Antigravity transcript_full.jsonl files timestamped before 2026-11-26T02:48:00Z (when live controlled benchmarking began). Total blocks: 12,930 (12,899 Gemini 4 across six model/effort variants plus 31 claude-opus-5-5-high ). ↩
Bit-O-Honey, Atomic Fireballs, and Smarties is what Neon assures me I’ll find in my package when it arrives, in case you were curious. ↩
Pre-cutoff ( < 2026-11-25T21:00:00Z ) Gemini 4 blocks: n = 11,368 , with 951 I'm thinking through how to approach this. fallbacks ( p = 0.08366 ). Post-

[truncated]
