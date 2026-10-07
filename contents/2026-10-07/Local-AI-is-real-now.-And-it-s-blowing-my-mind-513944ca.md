---
source: "https://hive.technology/lab-notes/local-ai-is-real/"
hn_url: "https://news.ycombinator.com/item?id=49988005"
title: "Local AI is real now. And it's blowing my mind"
article_title: "Local AI is real now · Hive.Tech™"
image: "https://hive.technology/lab-notes/local-ai-is-real/og.png"
author: "tylermhall"
captured_at: "2026-10-07T04:38:23Z"
capture_tool: "hn-digest"
hn_id: 49988005
score: 1
comments: 1
posted_at: "2026-10-07T04:02:20Z"
tags:
  - hacker-news
---

# Local AI is real now. And it's blowing my mind

- HN: [49988005](https://news.ycombinator.com/item?id=49988005)
- Source: [hive.technology](https://hive.technology/lab-notes/local-ai-is-real/)
- Score: 1
- Comments: 1
- Posted: 2026-10-07T04:02:20Z

## Translation

Title: Local AI is real now. And it's blowing my mind
Article title: Local AI is real now · Hive.Tech™
Description: There is a whole class of work I want an agent chewing on all night that is not worth frontier prices. A year of trying local sucked. Then the hardware caught up: four open models on one machine, on solar, taking real work. The box, the numbers, and what it cannot do yet.

Article text:
[ HIVE.TECH ] ™
NEXUS
DEV
WORK
COMPANY
// DEV > LAB NOTES · 261002
There is a whole class of work I want an agent chewing on all night that is not worth frontier prices. A year of trying local sucked. Then the hardware caught up: four open models on one machine, on solar, taking real work. The box, the numbers, and what it cannot do yet.
FOR AGENTS: this article as plain text with a reading guide, local-ai-is-real.agents.md .
DISCUSS THIS WITH AN AGENT COPY
Read https://hive.technology/lab-notes/local-ai-is-real/local-ai-is-real.agents.md,
the article with a brief for you. Then talk it through with me: which of the
four jobs in it I actually have, what my machine can hold, which one model I
should pull first, and what I would still send to a paid API. Push back where
the article is wrong for my setup.
I want an agent chewing on server logs all night. Run regressions over years of data. Look for small optimizations. Finish a pass and start another one.
I'm still using frontier models like crazy: roughly 10 billion tokens in the last 90 days, most of it cached input. But there's a whole class of work I haven't bothered giving them. Continuous inference for jobs that need modest intelligence probably isn't worth it. I'd like to let those jobs run without counting tokens.
My bench already runs on solar. Could I use some of that power for inference? And could I do private customer work the same way, with planning, development, and the finished result all staying on the machine?
I'd been trying local models for about a year. Mostly, it sucked. Get something running, give it actual work, go back to the frontier tools.
Last week the lab got a Mac Studio: an Apple M5 Ultra with 256 GB of unified memory.
The Mac Studio on the bench: Apple M5 Ultra, 256 GB
The rest is free with your email. Lab Notes arrives Tuesday mornings, and one click takes you off.
We collect your email address to send what you asked for. It is not sold or shared, and is kept until you ask us to delete it. What we collect and why .
With 256 GB, the temptation is to load the biggest model that fits. I tried that. DeepSeek V4 Flash beside all three smaller models put the machine at 220 GB in use, with no swap. It ran. There was hardly any room left to work.
The weights need memory, but so do the KV caches that grow as the models carry context. That lineup left about 20 GB for them.
I ended up with four slots: big, small text, small vision, and decision. A profile picks the models that fit together, with room to use them.
The everyday Work profile puts Qwen3.8 Flash Next in the big slot at 99 GB. Gemma 4 26B, at 4-bit, gets most of the small work. Qwen 3.6 35B-A3B, with 3 billion parameters active per token, gets the pictures. Clef-Flash 9B takes the decision slot.
A profile budgets each model at the size the registry expects (99, 21, 21, and 18.5 GB here), plus 15 GB for macOS. That leaves 82 GB for context. The measured resident sizes in the table below run smaller; the budget keeps the margin.
For the Smart profile, DeepSeek takes the big slot, budgeted at its 156 GB on disk (128 GB resident once loaded). Qwen 3.6 and Clef-Flash come out. Gemma stays because it's the small model I lean on. That leaves 64 GB.
Keeping useful models loaded matters. Qwen takes 8 seconds to load from disk, Gemma 15 seconds, DeepSeek 32 seconds. I don't want to keep paying that wait between jobs.
I wanted to know how they ran together, too. On 2 October I benchmarked the box while the small models took traffic on the same GPU throughout.
Gemma, Qwen 3.6, and DeepSeek used oMLX 0.7.0. Flash Next used MTPLX 2.12.0. Requests were streamed at temperature zero, with unique prompts so the prompt cache couldn't answer. The main measurements below are medians of three runs per cell. Generation counts thinking tokens as well as answer tokens, since you wait for both; some Qwen 3.6 runs spent their whole output budget thinking.
Flash Next read a 17,000-token prompt in about 8 seconds. Gemma took about 4 seconds. Qwen took under 3 seconds. The small models generated 91 to 115 tokens per second across the tested contexts. I could stop treating every request as a demonstration and get on with it.
Flash Next's generation medians were 51 to 58 tokens per second, with individual runs ranging from 44 to 70 tokens per second. It uses speculative decoding: a small draft model guesses ahead and the big one checks. About 55 to 65 percent of draft tokens were accepted. Speed follows that acceptance.
DeepSeek generated 27 to 31 tokens per second. Its starred prefill numbers came from single uncached requests in the first pass, including 13,139 tokens in 21.8 seconds. I never ran the clean prefill pass against it.
I measured speed. I didn't run the same prompts against a cloud API, measure whether tasks were done right, or measure the share of my work that moved local.
The work itself still needs boundaries. I tried a local model as plan-manager. It drifted.
So Claude Code still reads the repository, works through the ambiguity with me, and writes the plan and task contracts. That's the arrangement in Plans for agents . Local opencode or pi sessions take one contract each. Do the work. Write the result to the task file. Stop.
The frontier reads the results back. It doesn't read all the trying: intermediate diffs, test output, and context.
When privacy is the concern, I take the plan-manager seat. Opencode or pi runs against the bench, and planning, development, and producing the result happen in that session. The model is ours and no inference leaves the box.
That's now usable. Local inference is a growing part of my stack.
The fourth slot gets me closer to those jobs I wanted running all night. Clef-Flash is a 9-billion-parameter decision model from Cloudflare. Give it a state and typed questions with allowed answers. It returns probabilities, with no free text and no sampling step.
Here's a request and what came back from the box:
curl -s -X POST http://studio.local:8004/v1/systemone -H 'Content-Type: application/json' -d '{
"model": "clef-flash",
"state": "Checkout is returning errors and orders are blocked.",
"questions": {
"outage": {"type": "noul", "instructions": "Is a service down?"},
"urgency": {"type": "score", "criteria": ["Can wait", "This week", "Today"]},
"technical": {"type": "noul", "instructions": "Is this a technical problem?"}
}}'
outage 0.78
urgency Can wait 0.10 This week 0.06 Today 0.83
technical 0.92
I can ask about an event without parsing a paragraph afterward. Getting it onto Apple Silicon through PyTorch took a week and one retired attempt with the 27B model.
Those timings cover ten distinct states, about 310 tokens and three questions each. Cold after idle, Flash's first decision took about 2 seconds.
The 27B gave the outage example 0.92 where Flash gave 0.78. I haven't measured which is better calibrated. The 9B kept the seat for the memory. I'll get into Clef properly in a separate piece next week.
By then I had several runtimes taking work and needed to see what the box was doing. My agents built this page in two days.
Each runtime's own class supplies a snapshot every 5 seconds. Every field carries evidence and age. Stale means null with a reason. Profile cards show what would unload, what would load, and the memory afterward. "Short by 18 GB" tells me something before I press the button.
A collector records chip counters every 15 seconds: temperature, power, fans, memory pressure, swap, disk writes, and runtime state. It keeps 30 days at about 10 MB per day. There's an event log, plus per-runtime logs with keys redacted. One runtime had logged an API key in plain text.
It's a page for my box, not a product.
Next I want a router in front of the four slots. Small when it can do the job, large when it must, paid API when the local models are out of their depth. I may release it one day.
You can start smaller. One small model takes 15 to 20 GB resident, so a 32 to 48 GB machine can run that slot.
For now, this one runs on the solar that already powers my bench. It's taking real work. I can finally start finding out which of those always-running jobs are worth keeping.
#ai/local #ai/models #ai/tools
The thought of the week and what shipped. One email, Tuesday morning.
We collect your email address to send what you asked for. It is not sold or shared, and is kept until you ask us to delete it. What we collect and why .
JOIN THE DISCORD
DEV > LAB NOTES · 261002
Hive.Tech™, Hive.Nexus™, and Hive.Dev™ are trademarks of Hive Technology.

## Original Extract

There is a whole class of work I want an agent chewing on all night that is not worth frontier prices. A year of trying local sucked. Then the hardware caught up: four open models on one machine, on solar, taking real work. The box, the numbers, and what it cannot do yet.

[ HIVE.TECH ] ™
NEXUS
DEV
WORK
COMPANY
// DEV > LAB NOTES · 261002
There is a whole class of work I want an agent chewing on all night that is not worth frontier prices. A year of trying local sucked. Then the hardware caught up: four open models on one machine, on solar, taking real work. The box, the numbers, and what it cannot do yet.
FOR AGENTS: this article as plain text with a reading guide, local-ai-is-real.agents.md .
DISCUSS THIS WITH AN AGENT COPY
Read https://hive.technology/lab-notes/local-ai-is-real/local-ai-is-real.agents.md,
the article with a brief for you. Then talk it through with me: which of the
four jobs in it I actually have, what my machine can hold, which one model I
should pull first, and what I would still send to a paid API. Push back where
the article is wrong for my setup.
I want an agent chewing on server logs all night. Run regressions over years of data. Look for small optimizations. Finish a pass and start another one.
I'm still using frontier models like crazy: roughly 10 billion tokens in the last 90 days, most of it cached input. But there's a whole class of work I haven't bothered giving them. Continuous inference for jobs that need modest intelligence probably isn't worth it. I'd like to let those jobs run without counting tokens.
My bench already runs on solar. Could I use some of that power for inference? And could I do private customer work the same way, with planning, development, and the finished result all staying on the machine?
I'd been trying local models for about a year. Mostly, it sucked. Get something running, give it actual work, go back to the frontier tools.
Last week the lab got a Mac Studio: an Apple M5 Ultra with 256 GB of unified memory.
The Mac Studio on the bench: Apple M5 Ultra, 256 GB
The rest is free with your email. Lab Notes arrives Tuesday mornings, and one click takes you off.
We collect your email address to send what you asked for. It is not sold or shared, and is kept until you ask us to delete it. What we collect and why .
With 256 GB, the temptation is to load the biggest model that fits. I tried that. DeepSeek V4 Flash beside all three smaller models put the machine at 220 GB in use, with no swap. It ran. There was hardly any room left to work.
The weights need memory, but so do the KV caches that grow as the models carry context. That lineup left about 20 GB for them.
I ended up with four slots: big, small text, small vision, and decision. A profile picks the models that fit together, with room to use them.
The everyday Work profile puts Qwen3.8 Flash Next in the big slot at 99 GB. Gemma 4 26B, at 4-bit, gets most of the small work. Qwen 3.6 35B-A3B, with 3 billion parameters active per token, gets the pictures. Clef-Flash 9B takes the decision slot.
A profile budgets each model at the size the registry expects (99, 21, 21, and 18.5 GB here), plus 15 GB for macOS. That leaves 82 GB for context. The measured resident sizes in the table below run smaller; the budget keeps the margin.
For the Smart profile, DeepSeek takes the big slot, budgeted at its 156 GB on disk (128 GB resident once loaded). Qwen 3.6 and Clef-Flash come out. Gemma stays because it's the small model I lean on. That leaves 64 GB.
Keeping useful models loaded matters. Qwen takes 8 seconds to load from disk, Gemma 15 seconds, DeepSeek 32 seconds. I don't want to keep paying that wait between jobs.
I wanted to know how they ran together, too. On 2 October I benchmarked the box while the small models took traffic on the same GPU throughout.
Gemma, Qwen 3.6, and DeepSeek used oMLX 0.7.0. Flash Next used MTPLX 2.12.0. Requests were streamed at temperature zero, with unique prompts so the prompt cache couldn't answer. The main measurements below are medians of three runs per cell. Generation counts thinking tokens as well as answer tokens, since you wait for both; some Qwen 3.6 runs spent their whole output budget thinking.
Flash Next read a 17,000-token prompt in about 8 seconds. Gemma took about 4 seconds. Qwen took under 3 seconds. The small models generated 91 to 115 tokens per second across the tested contexts. I could stop treating every request as a demonstration and get on with it.
Flash Next's generation medians were 51 to 58 tokens per second, with individual runs ranging from 44 to 70 tokens per second. It uses speculative decoding: a small draft model guesses ahead and the big one checks. About 55 to 65 percent of draft tokens were accepted. Speed follows that acceptance.
DeepSeek generated 27 to 31 tokens per second. Its starred prefill numbers came from single uncached requests in the first pass, including 13,139 tokens in 21.8 seconds. I never ran the clean prefill pass against it.
I measured speed. I didn't run the same prompts against a cloud API, measure whether tasks were done right, or measure the share of my work that moved local.
The work itself still needs boundaries. I tried a local model as plan-manager. It drifted.
So Claude Code still reads the repository, works through the ambiguity with me, and writes the plan and task contracts. That's the arrangement in Plans for agents . Local opencode or pi sessions take one contract each. Do the work. Write the result to the task file. Stop.
The frontier reads the results back. It doesn't read all the trying: intermediate diffs, test output, and context.
When privacy is the concern, I take the plan-manager seat. Opencode or pi runs against the bench, and planning, development, and producing the result happen in that session. The model is ours and no inference leaves the box.
That's now usable. Local inference is a growing part of my stack.
The fourth slot gets me closer to those jobs I wanted running all night. Clef-Flash is a 9-billion-parameter decision model from Cloudflare. Give it a state and typed questions with allowed answers. It returns probabilities, with no free text and no sampling step.
Here's a request and what came back from the box:
curl -s -X POST http://studio.local:8004/v1/systemone -H 'Content-Type: application/json' -d '{
"model": "clef-flash",
"state": "Checkout is returning errors and orders are blocked.",
"questions": {
"outage": {"type": "noul", "instructions": "Is a service down?"},
"urgency": {"type": "score", "criteria": ["Can wait", "This week", "Today"]},
"technical": {"type": "noul", "instructions": "Is this a technical problem?"}
}}'
outage 0.78
urgency Can wait 0.10 This week 0.06 Today 0.83
technical 0.92
I can ask about an event without parsing a paragraph afterward. Getting it onto Apple Silicon through PyTorch took a week and one retired attempt with the 27B model.
Those timings cover ten distinct states, about 310 tokens and three questions each. Cold after idle, Flash's first decision took about 2 seconds.
The 27B gave the outage example 0.92 where Flash gave 0.78. I haven't measured which is better calibrated. The 9B kept the seat for the memory. I'll get into Clef properly in a separate piece next week.
By then I had several runtimes taking work and needed to see what the box was doing. My agents built this page in two days.
Each runtime's own class supplies a snapshot every 5 seconds. Every field carries evidence and age. Stale means null with a reason. Profile cards show what would unload, what would load, and the memory afterward. "Short by 18 GB" tells me something before I press the button.
A collector records chip counters every 15 seconds: temperature, power, fans, memory pressure, swap, disk writes, and runtime state. It keeps 30 days at about 10 MB per day. There's an event log, plus per-runtime logs with keys redacted. One runtime had logged an API key in plain text.
It's a page for my box, not a product.
Next I want a router in front of the four slots. Small when it can do the job, large when it must, paid API when the local models are out of their depth. I may release it one day.
You can start smaller. One small model takes 15 to 20 GB resident, so a 32 to 48 GB machine can run that slot.
For now, this one runs on the solar that already powers my bench. It's taking real work. I can finally start finding out which of those always-running jobs are worth keeping.
#ai/local #ai/models #ai/tools
The thought of the week and what shipped. One email, Tuesday morning.
We collect your email address to send what you asked for. It is not sold or shared, and is kept until you ask us to delete it. What we collect and why .
JOIN THE DISCORD
DEV > LAB NOTES · 261002
Hive.Tech™, Hive.Nexus™, and Hive.Dev™ are trademarks of Hive Technology.
