---
source: "https://siimtolk.github.io/articles/llm-testing-overnight/"
hn_url: "https://news.ycombinator.com/item?id=50194948"
title: "Tested 18 new Local LLMs on two RTX 5060 Ti cards: none replaced Qwen3.8-27B"
article_title: "I let an AI agent test local LLMs while I sleep | Siim Tolk"
image: "https://siimtolk.github.io/assets/img/llm-testing-overnight/chart.png"
author: "TheQuantumPhant"
captured_at: "2026-10-22T13:31:26Z"
capture_tool: "hn-digest"
hn_id: 50194948
score: 1
comments: 0
posted_at: "2026-10-22T13:20:53Z"
tags:
  - hacker-news
---

# Tested 18 new Local LLMs on two RTX 5060 Ti cards: none replaced Qwen3.8-27B

- HN: [50194948](https://news.ycombinator.com/item?id=50194948)
- Source: [siimtolk.github.io](https://siimtolk.github.io/articles/llm-testing-overnight/)
- Score: 1
- Comments: 0
- Posted: 2026-10-22T13:20:53Z

## Translation

Title: Tested 18 new Local LLMs on two RTX 5060 Ti cards: none replaced Qwen3.8-27B
Article title: I let an AI agent test local LLMs while I sleep | Siim Tolk
Description: An AI agent tested 18 new open-weight models on my home server overnight. None replaced the Qwen3.8-27B my assistant already runs.

Article text:
ST ">
ST Siim Tolk
Articles
GitHub
RSS
Local LLM LLM evaluation AI agents
I let an AI agent test local LLMs while I sleep
An AI agent tested 18 new open-weight models on my home server overnight. None replaced the Qwen3.8-27B my assistant already runs.
Siim Tolk · 22 October 2026 · 26 min read
Every week a dozen new open-weight models claim to beat the last one. I wanted to know which of them could actually replace the model my personal assistant runs on, so I built a pipeline that finds candidates, tests them overnight and sends me a verdict over breakfast. Over the past two weeks it has tested eighteen models, and none has replaced the incumbent, an uncensored build of Qwen3.8-27B. The best one passed a single graded run more, but was 18% slower and refused a harmless request. The code, the test suite and every result are on GitHub: local-llm-eval .
Why leaderboards were not enough
I run a personal Hermes assistant on my own hardware. It builds morning briefings, reads mail, sets reminders, writes code, searches the web, runs a daily quiz and answers in Estonian as often as in English. Behind it sits one open-weight model on a home server with two consumer GPUs and 32 GB of video memory between them.
That model answers every request all day. Swapping it changes every reply the assistant sends. A new release can top a public leaderboard and still be useless for me. NVIDIA’s new Nemotron 3.5 Bolt 74B-A9B ranks 4th among models under 100B parameters on Artificial Analysis. The build that fits my cards passed 247 of 291 runs in my suite, against production’s 283, and only 16 of 45 in Estonian. A model might not speak Estonian, or it might think for so long that a simple reminder takes a minute. Its chat template might return an empty reply.
No leaderboard answers that, so I built my own test. Then I got tired of tracking releases and running it by hand. The overnight part borrows the idea of Andrej Karpathy’s autoresearch , where an agent runs training experiments while you sleep. Here the same loop picks a production model instead.
The framework: a loop that runs itself
One daily loop runs five stages. Plain code finds, downloads and reports; the production model judges, and a Claude Code agent runs the tests.
Discover. Each morning a script pulls trending models, matching search results and new uploads from trusted publishers on Hugging Face, plus the day’s top community posts about new releases. Code drops adapters, embedding and OCR models, and anything too big for the GPUs, already tried, older than a month or barely liked.
Judge. The production model checks the shortlist against its own specs and the hardware, and picks up to three that look at least as capable. The picks arrive in my morning briefing with a poll.
Fetch. A risk check looks at each pick’s Hugging Face file scans, its GGUF header and the uploader. The ones I tick in the poll are downloaded over the day in a quant that fits the GPUs, and a 4-bit vLLM build of the same model is queued when one exists.
Test. At night a coding agent (Claude Code) takes over the GPUs, runs the suite on each candidate, compares the scores with production’s, writes up and commits the result and restores production.
Report. At breakfast a WhatsApp message gives each model’s verdict (passed, rejected or failed to start) and the reason.
Production must be back by morning, when the day’s jobs start. Until then, any LLM requests fall back to Claude and the assistant keeps answering.
The serving stack: two 16 GB cards, one model
Everything runs on llama.cpp ’s llama-server , built from source with CUDA. It speaks the OpenAI chat API. The assistant, a web chat UI and the test suite all use the same endpoint. The hardware is two NVIDIA RTX 5060 Ti cards with 16 GB each, (plus an old 6 GB GTX 1660 Super that only holds the vision encoder and is not relevant for this test).
A 27B dense model at 4 bits takes about 15.5 GB, which leaves about 0.5 GB of a 16 GB card for context and buffers. The model is split by layers across both 5060 Tis. Each card then holds half the weights, and the freed memory goes to context. These are the production settings:
Table 1. Production settings for llama-server.
In total it uses 25.4 GB of the 32 GB.
Production runs llama.cpp build b11707 (commit 1268edac5 ), compiled locally:
Two copies of llama.cpp sit side by side, so tests never touch what production runs:
llama.cpp-prod , pinned at b11707: production.
llama.cpp-next , a test copy that every candidate runs on. It is the same build as production unless a new model’s architecture needs newer code; then the agent pulls and rebuilds it.
Production itself moves to a newer build only after its own model has passed the full suite on it at least as well as before. Apart from that, every test is about the model, not llama.cpp.
The suite has 97 cases taken from what the assistant really does. Each runs three times, because sampling is random and a model that gets something right once in three tries is not reliable. Every answer is graded by code, not by another LLM. The same answer gets the same score every night.
Table 3. What the 97-case suite tests.
The model I want also has to carry out edgy but harmless tasks instead of refusing them: write a port scanner for my home network, test my own web form for SQL injection, or set up a torrent client behind a VPN and add torrents. Those are counted separately.
Two more passes run next to the graded ones. A speed pass measures prose and tool-call output, a 30k-token prompt, two users at once and two long chats taking turns. Then the everyday cases run again with reasoning switched off, which shows whether the model is still usable for quick jobs, and how much faster.
Reasoning models write a hidden chain of thought before they answer, and every thought costs time. The suite controls this with three settings, sent the same way to every model:
Table 4. How the suite controls thinking.
There is no separate thinking budget. The 6,000-token limit is the only hard cap, on purpose: the test should show whether a model knows when to stop.
Why low? With no effort set, Qwen3.8’s template defaults to “xhigh” and asks the model to validate assumptions and weigh alternatives. That is too slow for reminders and quick questions, so production runs at low, and a client can ask for more per request. The template accepts only xhigh, medium and low. Sending “high” returns an error, which is worth knowing before you copy settings from one model to another.
At low effort production stays well inside the limit. Its median graded run used 265 tokens, and only 3 of 291 runs hit 6,000, all of them code tasks. One edgy prompt also hit its 2,000 cap. With thinking off the median run drops to 51 tokens and the longest to 364.
Thinking off saves time on every model but one, and costs most of them accuracy. The second pass covers the 47 everyday cases, 141 runs per model:
Table 5. Thinking on vs off, 47 everyday cases (141 runs).
Muse Glimmer is the exception because its template ignores the switch and it reasons anyway. MiniCPM is the only model that got better with thinking off (by 20 test runs). It never hit the token limit in these cases. Instead, with thinking on it often thought briefly and then returned an empty reply: all three runs of the exact-reply and Estonian-reply cases came back blank. Thinking off fixed most of those, and it stopped sending messages without asking first. The Qwen models lost 12 to 19 runs, mostly in tool use, instruction rules and Estonian.
On single cases the picture is less tidy. Production used more tokens with thinking off on 2 of the 47 cases, and nerkyor’s Coder390 on 3, because the model worked the problem out in its visible answer instead. The second pass leaves out the hard puzzles, so the suite cannot say how often that happens on harder questions. In a quick check of two requests per setting, production answered one counting puzzle in about 800 tokens with thinking off, against about 550 at low effort.
So is thinking off worth using? For production, yes, on narrow jobs. Its median answer takes a quarter to two thirds of the time, with no loss on these:
Grading quiz answers: 18 of 18 either way, 1.6 s per answer instead of 2.4 s.
Simple English tool calls: lights, weather, shopping, timers, calendar and three tools at once all passed 3 of 3.
Fixed-format text: the WhatsApp brief layout, staying silent, exact replies, plain text and three bullets all passed 3 of 3, in under a second.
Simple extraction and Estonian replies: invoices, file names, log lines, Estonian summaries and weather all passed 3 of 3.
It loses where the model has to stop and check something:
Asking before acting: “Send a message to Max” was sent at once, with made-up or empty text, in all three runs.
Typos: “turn of the kitchn lights” failed 2 of 3 runs: once the model searched the web for how to turn lights on, once it turned them on.
Estonian tool calls: “Pane elutoa tuli 50% peale” (set the living-room light to 50%) got no tool call in any run.
Grammar and counting: Estonian case forms, numbers in words, an acrostic and a word limit.
The assistant still runs every request with thinking on. A good split would be thinking off for quiz grading and fixed-format briefs, and thinking on for anything that acts in the world or speaks Estonian.
The same settings land differently on other models, because each template decides what “low” means. Gemma 4 26B ran out of tokens 20 times at the same limit. Muse Glimmer’s template ignores the thinking-off switch and reasons anyway, at about three times the tokens, so its thinking-off score cannot be compared with production’s.
Letting an agent run with nobody watching
An agent that downloads strangers’ models and restarts my production server at 3 a.m. needs firm limits. Five rules make it safe enough to leave alone (at least on a home server that holds nothing I could not rebuild):
The deadline is enforced from outside. A separate runner gives the agent a time limit, kills its whole session at the deadline whatever it is doing, and restores production itself. The agent is told the deadline, but nothing depends on it keeping to it. The prompt still has to say what it means: one early run read “the test window is 01:00 to 06:00”, decided it was too early, and waited three hours. It now says the window is open and to start right away.
Restoring production is idempotent and runs every time. One script brings production back after every test, after a killed run, and again from a timer at the end of the window. It does nothing when production already answers.
Downloaded content is data, never instructions. Model cards, file names and forum posts are written by strangers. They are read for facts only, no code from a model’s repo is ever run, and builds that need trust_remote_code are refused. Code the model writes during the test runs in a sandbox without network or my files.
One task per run, with recorded state. The agent follows a step-by-step procedure for one candidate at a time: how to start it, what to measure, where results go and how to write the verdict. It decides details such as memory settings, not the shape of the night. The candidate’s status is recorded before the agent starts, and a run that ends without a result is marked failed rather than passed. The agent may run exactly the commands the procedure needs, and may not edit the production config, use admin rights or force-push. Stopping production is one named script, because a one-liner to do the same thing was blocked by the safety check, and that night’s test never started.
A human decides, the agent reports. I pick what gets downloaded, and the agent never switches production. Every result is a committed JSON file plus a write-up, every commit passes a leak check because the repo is public, and each run’s final message is kept in a log.
Note that some of these limits rest on the agent’s good behaviour (and as hope is not a strateg

[truncated]

## Original Extract

An AI agent tested 18 new open-weight models on my home server overnight. None replaced the Qwen3.8-27B my assistant already runs.

ST ">
ST Siim Tolk
Articles
GitHub
RSS
Local LLM LLM evaluation AI agents
I let an AI agent test local LLMs while I sleep
An AI agent tested 18 new open-weight models on my home server overnight. None replaced the Qwen3.8-27B my assistant already runs.
Siim Tolk · 22 October 2026 · 26 min read
Every week a dozen new open-weight models claim to beat the last one. I wanted to know which of them could actually replace the model my personal assistant runs on, so I built a pipeline that finds candidates, tests them overnight and sends me a verdict over breakfast. Over the past two weeks it has tested eighteen models, and none has replaced the incumbent, an uncensored build of Qwen3.8-27B. The best one passed a single graded run more, but was 18% slower and refused a harmless request. The code, the test suite and every result are on GitHub: local-llm-eval .
Why leaderboards were not enough
I run a personal Hermes assistant on my own hardware. It builds morning briefings, reads mail, sets reminders, writes code, searches the web, runs a daily quiz and answers in Estonian as often as in English. Behind it sits one open-weight model on a home server with two consumer GPUs and 32 GB of video memory between them.
That model answers every request all day. Swapping it changes every reply the assistant sends. A new release can top a public leaderboard and still be useless for me. NVIDIA’s new Nemotron 3.5 Bolt 74B-A9B ranks 4th among models under 100B parameters on Artificial Analysis. The build that fits my cards passed 247 of 291 runs in my suite, against production’s 283, and only 16 of 45 in Estonian. A model might not speak Estonian, or it might think for so long that a simple reminder takes a minute. Its chat template might return an empty reply.
No leaderboard answers that, so I built my own test. Then I got tired of tracking releases and running it by hand. The overnight part borrows the idea of Andrej Karpathy’s autoresearch , where an agent runs training experiments while you sleep. Here the same loop picks a production model instead.
The framework: a loop that runs itself
One daily loop runs five stages. Plain code finds, downloads and reports; the production model judges, and a Claude Code agent runs the tests.
Discover. Each morning a script pulls trending models, matching search results and new uploads from trusted publishers on Hugging Face, plus the day’s top community posts about new releases. Code drops adapters, embedding and OCR models, and anything too big for the GPUs, already tried, older than a month or barely liked.
Judge. The production model checks the shortlist against its own specs and the hardware, and picks up to three that look at least as capable. The picks arrive in my morning briefing with a poll.
Fetch. A risk check looks at each pick’s Hugging Face file scans, its GGUF header and the uploader. The ones I tick in the poll are downloaded over the day in a quant that fits the GPUs, and a 4-bit vLLM build of the same model is queued when one exists.
Test. At night a coding agent (Claude Code) takes over the GPUs, runs the suite on each candidate, compares the scores with production’s, writes up and commits the result and restores production.
Report. At breakfast a WhatsApp message gives each model’s verdict (passed, rejected or failed to start) and the reason.
Production must be back by morning, when the day’s jobs start. Until then, any LLM requests fall back to Claude and the assistant keeps answering.
The serving stack: two 16 GB cards, one model
Everything runs on llama.cpp ’s llama-server , built from source with CUDA. It speaks the OpenAI chat API. The assistant, a web chat UI and the test suite all use the same endpoint. The hardware is two NVIDIA RTX 5060 Ti cards with 16 GB each, (plus an old 6 GB GTX 1660 Super that only holds the vision encoder and is not relevant for this test).
A 27B dense model at 4 bits takes about 15.5 GB, which leaves about 0.5 GB of a 16 GB card for context and buffers. The model is split by layers across both 5060 Tis. Each card then holds half the weights, and the freed memory goes to context. These are the production settings:
Table 1. Production settings for llama-server.
In total it uses 25.4 GB of the 32 GB.
Production runs llama.cpp build b11707 (commit 1268edac5 ), compiled locally:
Two copies of llama.cpp sit side by side, so tests never touch what production runs:
llama.cpp-prod , pinned at b11707: production.
llama.cpp-next , a test copy that every candidate runs on. It is the same build as production unless a new model’s architecture needs newer code; then the agent pulls and rebuilds it.
Production itself moves to a newer build only after its own model has passed the full suite on it at least as well as before. Apart from that, every test is about the model, not llama.cpp.
The suite has 97 cases taken from what the assistant really does. Each runs three times, because sampling is random and a model that gets something right once in three tries is not reliable. Every answer is graded by code, not by another LLM. The same answer gets the same score every night.
Table 3. What the 97-case suite tests.
The model I want also has to carry out edgy but harmless tasks instead of refusing them: write a port scanner for my home network, test my own web form for SQL injection, or set up a torrent client behind a VPN and add torrents. Those are counted separately.
Two more passes run next to the graded ones. A speed pass measures prose and tool-call output, a 30k-token prompt, two users at once and two long chats taking turns. Then the everyday cases run again with reasoning switched off, which shows whether the model is still usable for quick jobs, and how much faster.
Reasoning models write a hidden chain of thought before they answer, and every thought costs time. The suite controls this with three settings, sent the same way to every model:
Table 4. How the suite controls thinking.
There is no separate thinking budget. The 6,000-token limit is the only hard cap, on purpose: the test should show whether a model knows when to stop.
Why low? With no effort set, Qwen3.8’s template defaults to “xhigh” and asks the model to validate assumptions and weigh alternatives. That is too slow for reminders and quick questions, so production runs at low, and a client can ask for more per request. The template accepts only xhigh, medium and low. Sending “high” returns an error, which is worth knowing before you copy settings from one model to another.
At low effort production stays well inside the limit. Its median graded run used 265 tokens, and only 3 of 291 runs hit 6,000, all of them code tasks. One edgy prompt also hit its 2,000 cap. With thinking off the median run drops to 51 tokens and the longest to 364.
Thinking off saves time on every model but one, and costs most of them accuracy. The second pass covers the 47 everyday cases, 141 runs per model:
Table 5. Thinking on vs off, 47 everyday cases (141 runs).
Muse Glimmer is the exception because its template ignores the switch and it reasons anyway. MiniCPM is the only model that got better with thinking off (by 20 test runs). It never hit the token limit in these cases. Instead, with thinking on it often thought briefly and then returned an empty reply: all three runs of the exact-reply and Estonian-reply cases came back blank. Thinking off fixed most of those, and it stopped sending messages without asking first. The Qwen models lost 12 to 19 runs, mostly in tool use, instruction rules and Estonian.
On single cases the picture is less tidy. Production used more tokens with thinking off on 2 of the 47 cases, and nerkyor’s Coder390 on 3, because the model worked the problem out in its visible answer instead. The second pass leaves out the hard puzzles, so the suite cannot say how often that happens on harder questions. In a quick check of two requests per setting, production answered one counting puzzle in about 800 tokens with thinking off, against about 550 at low effort.
So is thinking off worth using? For production, yes, on narrow jobs. Its median answer takes a quarter to two thirds of the time, with no loss on these:
Grading quiz answers: 18 of 18 either way, 1.6 s per answer instead of 2.4 s.
Simple English tool calls: lights, weather, shopping, timers, calendar and three tools at once all passed 3 of 3.
Fixed-format text: the WhatsApp brief layout, staying silent, exact replies, plain text and three bullets all passed 3 of 3, in under a second.
Simple extraction and Estonian replies: invoices, file names, log lines, Estonian summaries and weather all passed 3 of 3.
It loses where the model has to stop and check something:
Asking before acting: “Send a message to Max” was sent at once, with made-up or empty text, in all three runs.
Typos: “turn of the kitchn lights” failed 2 of 3 runs: once the model searched the web for how to turn lights on, once it turned them on.
Estonian tool calls: “Pane elutoa tuli 50% peale” (set the living-room light to 50%) got no tool call in any run.
Grammar and counting: Estonian case forms, numbers in words, an acrostic and a word limit.
The assistant still runs every request with thinking on. A good split would be thinking off for quiz grading and fixed-format briefs, and thinking on for anything that acts in the world or speaks Estonian.
The same settings land differently on other models, because each template decides what “low” means. Gemma 4 26B ran out of tokens 20 times at the same limit. Muse Glimmer’s template ignores the thinking-off switch and reasons anyway, at about three times the tokens, so its thinking-off score cannot be compared with production’s.
Letting an agent run with nobody watching
An agent that downloads strangers’ models and restarts my production server at 3 a.m. needs firm limits. Five rules make it safe enough to leave alone (at least on a home server that holds nothing I could not rebuild):
The deadline is enforced from outside. A separate runner gives the agent a time limit, kills its whole session at the deadline whatever it is doing, and restores production itself. The agent is told the deadline, but nothing depends on it keeping to it. The prompt still has to say what it means: one early run read “the test window is 01:00 to 06:00”, decided it was too early, and waited three hours. It now says the window is open and to start right away.
Restoring production is idempotent and runs every time. One script brings production back after every test, after a killed run, and again from a timer at the end of the window. It does nothing when production already answers.
Downloaded content is data, never instructions. Model cards, file names and forum posts are written by strangers. They are read for facts only, no code from a model’s repo is ever run, and builds that need trust_remote_code are refused. Code the model writes during the test runs in a sandbox without network or my files.
One task per run, with recorded state. The agent follows a step-by-step procedure for one candidate at a time: how to start it, what to measure, where results go and how to write the verdict. It decides details such as memory settings, not the shape of the night. The candidate’s status is recorded before the agent starts, and a run that ends without a result is marked failed rather than passed. The agent may run exactly the commands the procedure needs, and may not edit the production config, use admin rights or force-push. Stopping production is one named script, because a one-liner to do the same thing was blocked by the safety check, and that night’s test never started.
A human decides, the agent reports. I pick what gets downloaded, and the agent never switches production. Every result is a committed JSON file plus a write-up, every commit passes a leak check because the repo is public, and each run’s final message is kept in a log.
Note that some of these limits rest on the agent’s good behaviour (and as hope is not a strateg

[truncated]
