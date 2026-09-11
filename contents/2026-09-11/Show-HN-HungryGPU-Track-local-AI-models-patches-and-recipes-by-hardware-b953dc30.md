---
source: "https://hungrygpu.com"
hn_url: "https://news.ycombinator.com/item?id=49664363"
title: "Show HN: HungryGPU – Track local AI models, patches and recipes by hardware"
article_title: "HungryGPU"
image: "https://hungrygpu.com/og.png"
author: "clawterminal"
captured_at: "2026-09-11T20:33:50Z"
capture_tool: "hn-digest"
hn_id: 49664363
score: 2
comments: 0
posted_at: "2026-09-11T19:49:39Z"
tags:
  - hacker-news
---

# Show HN: HungryGPU – Track local AI models, patches and recipes by hardware

- HN: [49664363](https://news.ycombinator.com/item?id=49664363)
- Source: [hungrygpu.com](https://hungrygpu.com)
- Score: 2
- Comments: 0
- Posted: 2026-09-11T19:49:39Z

## Translation

Title: Show HN: HungryGPU – Track local AI models, patches and recipes by hardware
Article title: HungryGPU
Description: What shipped today for the machine you actually own — and what quietly leaves it out. Merged patches, open weights, community recipes and papers, tagged per machine and rebuilt daily, each with the line of evidence behind it.

Article text:
Skip to content
Close
Plain text search over names. To ask for something like
uncensored or MoE , use the facets in Models: they say exactly what they select and
stay visible while they apply.
A repository, a person, a paper, a model or a benchmark
this site should be reading and is not. It goes into a queue the daily pass reads before it
starts; it is not published as a finding and it does not change any ranking.
The catalogue already answers over MCP without an account — what
changed, what runs on which box, who beat a published ceiling. A token adds what is
yours: your setup fills itself in, and you can queue things without leaving the agent.
Anything about the site itself: something broken, a column
that does not read, a figure that cannot be right, something you expected to find and did not.
It goes to a person who reads it — it is not published anywhere and it does not change
the catalogue.
Your local AI starts here.
Feed your GPU.
Discover what you can run on your machine.
Community recipes and useful patches for running local AI.
Made for your hardware What are you running?
Pick your machine. Find your next setup.
Community setups with the models, engines and instructions behind them.
See what changed upstream, what helps your hardware and what leaves it out.
This view is about one machine. Which is yours?
Or browse recipes
and papers , which work without one.
Open weights keep closing on the best proprietary models. Those two
lines converging are the curve this site is named after.
Nov 2022
Sep 2026
Late 2024: open weights at roughly 40% of the best proprietary
model. Late 2026: within about 15% .
Traced from the Artificial Analysis Intelligence Index v4.3 —
their figures, not ours, so the axis carries no numbers.
artificialanalysis.ai ↗
The question is whether it reaches your desk.
So every day this reads the merged pull requests, the model releases,
the community setups and the papers — and tags each one with the machines it helps,
and the ones it leaves behind .
Every tag shows its evidence — the diff line that justifies it.
An estimate says so. Memory fit is arithmetic, never a run.
Unknown stays unknown — never compatible, never false.
Pick as many as you have. With two or more, every patch says which of
your boxes it is for — and which it leaves out.
The same week, filtered to what lands on one box — and to what
quietly leaves it out, which is the half nobody announces.
Merged pull requests that touch this machine, by the week they landed.
Which part of the stack is moving for this machine.
Only models somebody actually ran appear here: the vertical axis is a
measured figure from a community recipe, not an estimate.
Does your setup need a change?
Corrections address a reported problem; improvements are optional. Nothing here knows which fixes your installed version already has.
2. The work, grouped by what it costs you
Real setups by hardware, engine and quantisation, each with the author's own conditions.
How it got here: which machines are gaining people who publish for them, how fast the
numbers move, and how much of what ships fits a box you can buy.
Every attributed single-box figure, per machine: the bar is the middle half, the thick
line the median, the ring the record. The record is not the expectation — on a
DGX Spark it is five times the median. Multi-box figures are in the tooltip, not the drawing.
Engines named by published setups, per machine. Not one preference: a DGX Spark is
mostly vLLM, a Strix Halo almost entirely llama.cpp.
Techniques named by setups for your machine — what people did, not what papers
propose.
One bar per month, split by the machine the setup covers; one covering three machines
counts in all three. The date is the repository’s, not the figure’s
— when a number was published is recorded nowhere.
The same setups, one row per machine instead of one pile. Each row is scaled to its
own busiest month , printed on the right, so every shape is legible and heights across
rows mean nothing. Totals do not compare either: each machine is found by its own search
terms, so a row compares its present against its own past.
Ordered by when the first setup appeared; the number on the right is the total. Machines
nobody has published for are counted underneath, not listed.
Every attributed single-stream figure, and the running best — the record as it
stood that day, not today’s. Aggregates, prefill and training figures are excluded,
by the same rule the records page uses.
Only changes worth doing, by the part of the stack they touch.
arXiv submissions that cleared the bar, by area.
The catalogue and the usage figures arrive without history. Every pass appends a row, so
this one starts empty and fills itself from here.
Two counted quantities, no ranking: when they last published, and how many figures they
measured themselves. With a machine picked, the accounts publishing for it are in colour.
Up and to the right is active and measuring.
For the busiest accounts: figures they measured, figures quoted from elsewhere, and
setups reporting none. A setup without a figure is not worse, it is a different thing.
Accounts publishing each month, split by whether it was their first. A growing community
and the same twenty people publishing more give the same total curve; only this separates
them.
The two numbers that decide what you can load and how fast it decodes, on log scales. Up
and to the right is fast and roomy; bottom right is the memory-rich, bandwidth-poor corner.
Every dot answers on hover.
Dashed lines are memory ceilings at 4-bit, weights × bytes × 1.2. Arithmetic,
not a measurement : it knows nothing about KV cache, fragmentation or whether a kernel
exists. Only your machine’s ceiling is labelled; the rest answer on hover.
What exists, and how much of it your machine could load at all.
The bar is models with a published figure. “Fits” is not the bar : that
is arithmetic on the weights, printed as a number. The faint tail is what was measured for
the first time in the last 90 days.
Checkpoint count. This one measures who republishes : requantisations,
repacks and derivatives all count as releases.
The same families, ordered by use. High on the left and low here is a repackaging shop,
not a lab.
“Does not say” is the majority, and it is its own state: a paper silent about
its code has not told you there is none.
Mentions inside a 30-day window: what got written about, not what got adopted.
How many patches a keyword search would have marked for each machine, how many the tagger
marked, and where the two disagree.
Point an agent at this catalogue
Everything on this site answers over MCP: merged patches, community setups with the
figure each one measured, and which of them your machine can actually run. Your agent
asks; you do not come and look.
2 Ask it something worth asking
Paste this at your agent, in a checkout of whatever you serve models with:
It will use whats_new , records and
can_i_run . Nothing it reports is inferred: every row carries the
line of evidence behind it, and a memory estimate is never returned as a measurement.
Signing in is optional and changes nothing about what this site
publishes. What it stores if you do →
whats_new What changed since a date, for one machine or all
can_i_run Tested, estimated or unknown — never a guess dressed as a fact
find_setups Setups ranked by published decode speed
records Throughput ceilings beaten recently
machines The hardware catalogue: slugs, memory, bandwidth
You describe your hardware in every call.
my_setup The machines and checkpoints this account synced
propose Queue a repo, paper or model for the next daily pass
my_proposals What you proposed, and what was done with it
submit_run Publish a tok/s you measured yourself, with the engine output behind it
my_runs Your runs, and what the checks said about each
And the five above stop asking. whats_new and
records read your setup by themselves, so you never describe your
hardware twice.
Tokens are minted one at a time, shown once, stored hashed and revoked
individually. The account travels in Authorization: Bearer and never in a
cookie — a page you visit cannot make requests as you.
Device count, engine, quantisation and context differ between runs. Figures exclude prefill and concurrent streams. A missing run is missing evidence, not incompatibility.
Find the right model for the task
Each score belongs to the model and setup it was measured in; a community recipe may use another quantisation or context. Memory fit is an estimate. No published run is missing evidence, not incompatibility.
Usage now, and where it is going
Area is tokens routed in the last 30 days; colour is those 7 days against the three weeks before. Both published by OpenRouter, never combined.
Every cell says what produced it. Scores from different evaluations are never averaged into one number, a
missing value is shown as missing and not as zero, and a run on more devices than you have is marked.
Have an agent watch this for you
The same updates answer over MCP — filtered by your setup,
each one carrying the reason it is here — so you can ask what changed for my
machines this week from wherever you already work, instead of coming to look.
The catalogue answers over MCP as well as on screen, so an agent
can ask what this page shows: what changed for your machines, whether a model fits one
of them, whether anyone just beat the best published figure for your box. With a token
it reads the setup above by itself — you never describe your hardware twice —
and it can queue something for the daily pass without leaving the conversation.
Things to try and keep for later.
This form is the way to reach whoever runs the
site. There is no account to create and no address to write to — that is on
purpose, so one form is the only thing to keep working.
Feedback
const D = {"hw":{"dgx-spark":{"nombre":"NVIDIA DGX Spark (GB10)","familia":"cuda","cc":121,"cc_txt":"sm_121","mem_gb":121,"bw":273,"precision":"fp4/fp8/bf16","nota":"Unified CPU/GPU memory, single node, no NVLink. Memory-rich and bandwidth-poor: holds a 126 GB checkpoint an 80 GB H100 cannot, and decodes at a fraction of the speed. Frequently excluded by arch gates written for sm_100 or sm_120 only."},"rtx-5090":{"nombre":"GeForce RTX 5090","familia":"cuda","cc":120,"cc_txt":"sm_120","mem_gb":32,"bw":1792,"precision":"fp4/fp8/bf16","nota":"Consumer Blackwell. Fast memory, little of it: a 32 GB budget makes quantization and KV-cache work matter more than raw kernel speed."},"rtx-5060-ti":{"nombre":"GeForce RTX 5060 Ti (16 GB)","familia":"cuda","cc":120,"cc_txt":"sm_120","mem_gb":16,"bw":448,"precision":"fp4/fp8/bf16","nota":"Entry-level Blackwell: 16 GB at 448 GB/s, and the cheapest way there is to get fp4 in hardware. Blackwell-gated kernels compile and run here, so an arch gate that opens is real news for it \u2014 and then the memory budget, a sixth of a workstation card's, decides whether the weights fit at all. fp4 quantization and KV-cache work are the entry ticket here, not an optimisation."},"rtx-pro-6000":{"nombre":"RTX PRO 6000 Blackwell","familia":"cuda","cc":120,"cc_txt":"sm_120","mem_gb":96,"bw":1792,"precision":"fp4/fp8/bf16","nota":"Workstation Blackwell, same arch as the 5090 but 96 GB. The one card that runs big models locally at speed."},"b200":{"nombre":"NVIDIA B200","familia":"cuda","cc":100,"cc_txt":"sm_100","mem_gb":180,"bw":8000,"precision":"fp4/fp8/bf16","nota":"Datacenter Blackwell. Usually the FIRST arch a new Blackwell kernel supports, and often the only one for a while."},"h200":{"nombre":"NVIDIA H200","familia":"cuda","cc":90,"cc_txt":"sm_90a","mem_gb":141,"bw":4800,"precision":"fp8/bf16","nota":"Hopper with more memory. No fp4. Still the most common serving target."},"h100":{"nombre":"NVIDIA H100","familia":"cuda","cc":90,"cc_tx
[truncated]
Automatic hardware interpretation · verify against source ${esc(p.wy)}
Extracted action · not a verified instruction for your installation ${esc(p.do)}
${esc

[truncated]

## Original Extract

What shipped today for the machine you actually own — and what quietly leaves it out. Merged patches, open weights, community recipes and papers, tagged per machine and rebuilt daily, each with the line of evidence behind it.

Skip to content
Close
Plain text search over names. To ask for something like
uncensored or MoE , use the facets in Models: they say exactly what they select and
stay visible while they apply.
A repository, a person, a paper, a model or a benchmark
this site should be reading and is not. It goes into a queue the daily pass reads before it
starts; it is not published as a finding and it does not change any ranking.
The catalogue already answers over MCP without an account — what
changed, what runs on which box, who beat a published ceiling. A token adds what is
yours: your setup fills itself in, and you can queue things without leaving the agent.
Anything about the site itself: something broken, a column
that does not read, a figure that cannot be right, something you expected to find and did not.
It goes to a person who reads it — it is not published anywhere and it does not change
the catalogue.
Your local AI starts here.
Feed your GPU.
Discover what you can run on your machine.
Community recipes and useful patches for running local AI.
Made for your hardware What are you running?
Pick your machine. Find your next setup.
Community setups with the models, engines and instructions behind them.
See what changed upstream, what helps your hardware and what leaves it out.
This view is about one machine. Which is yours?
Or browse recipes
and papers , which work without one.
Open weights keep closing on the best proprietary models. Those two
lines converging are the curve this site is named after.
Nov 2022
Sep 2026
Late 2024: open weights at roughly 40% of the best proprietary
model. Late 2026: within about 15% .
Traced from the Artificial Analysis Intelligence Index v4.3 —
their figures, not ours, so the axis carries no numbers.
artificialanalysis.ai ↗
The question is whether it reaches your desk.
So every day this reads the merged pull requests, the model releases,
the community setups and the papers — and tags each one with the machines it helps,
and the ones it leaves behind .
Every tag shows its evidence — the diff line that justifies it.
An estimate says so. Memory fit is arithmetic, never a run.
Unknown stays unknown — never compatible, never false.
Pick as many as you have. With two or more, every patch says which of
your boxes it is for — and which it leaves out.
The same week, filtered to what lands on one box — and to what
quietly leaves it out, which is the half nobody announces.
Merged pull requests that touch this machine, by the week they landed.
Which part of the stack is moving for this machine.
Only models somebody actually ran appear here: the vertical axis is a
measured figure from a community recipe, not an estimate.
Does your setup need a change?
Corrections address a reported problem; improvements are optional. Nothing here knows which fixes your installed version already has.
2. The work, grouped by what it costs you
Real setups by hardware, engine and quantisation, each with the author's own conditions.
How it got here: which machines are gaining people who publish for them, how fast the
numbers move, and how much of what ships fits a box you can buy.
Every attributed single-box figure, per machine: the bar is the middle half, the thick
line the median, the ring the record. The record is not the expectation — on a
DGX Spark it is five times the median. Multi-box figures are in the tooltip, not the drawing.
Engines named by published setups, per machine. Not one preference: a DGX Spark is
mostly vLLM, a Strix Halo almost entirely llama.cpp.
Techniques named by setups for your machine — what people did, not what papers
propose.
One bar per month, split by the machine the setup covers; one covering three machines
counts in all three. The date is the repository’s, not the figure’s
— when a number was published is recorded nowhere.
The same setups, one row per machine instead of one pile. Each row is scaled to its
own busiest month , printed on the right, so every shape is legible and heights across
rows mean nothing. Totals do not compare either: each machine is found by its own search
terms, so a row compares its present against its own past.
Ordered by when the first setup appeared; the number on the right is the total. Machines
nobody has published for are counted underneath, not listed.
Every attributed single-stream figure, and the running best — the record as it
stood that day, not today’s. Aggregates, prefill and training figures are excluded,
by the same rule the records page uses.
Only changes worth doing, by the part of the stack they touch.
arXiv submissions that cleared the bar, by area.
The catalogue and the usage figures arrive without history. Every pass appends a row, so
this one starts empty and fills itself from here.
Two counted quantities, no ranking: when they last published, and how many figures they
measured themselves. With a machine picked, the accounts publishing for it are in colour.
Up and to the right is active and measuring.
For the busiest accounts: figures they measured, figures quoted from elsewhere, and
setups reporting none. A setup without a figure is not worse, it is a different thing.
Accounts publishing each month, split by whether it was their first. A growing community
and the same twenty people publishing more give the same total curve; only this separates
them.
The two numbers that decide what you can load and how fast it decodes, on log scales. Up
and to the right is fast and roomy; bottom right is the memory-rich, bandwidth-poor corner.
Every dot answers on hover.
Dashed lines are memory ceilings at 4-bit, weights × bytes × 1.2. Arithmetic,
not a measurement : it knows nothing about KV cache, fragmentation or whether a kernel
exists. Only your machine’s ceiling is labelled; the rest answer on hover.
What exists, and how much of it your machine could load at all.
The bar is models with a published figure. “Fits” is not the bar : that
is arithmetic on the weights, printed as a number. The faint tail is what was measured for
the first time in the last 90 days.
Checkpoint count. This one measures who republishes : requantisations,
repacks and derivatives all count as releases.
The same families, ordered by use. High on the left and low here is a repackaging shop,
not a lab.
“Does not say” is the majority, and it is its own state: a paper silent about
its code has not told you there is none.
Mentions inside a 30-day window: what got written about, not what got adopted.
How many patches a keyword search would have marked for each machine, how many the tagger
marked, and where the two disagree.
Point an agent at this catalogue
Everything on this site answers over MCP: merged patches, community setups with the
figure each one measured, and which of them your machine can actually run. Your agent
asks; you do not come and look.
2 Ask it something worth asking
Paste this at your agent, in a checkout of whatever you serve models with:
It will use whats_new , records and
can_i_run . Nothing it reports is inferred: every row carries the
line of evidence behind it, and a memory estimate is never returned as a measurement.
Signing in is optional and changes nothing about what this site
publishes. What it stores if you do →
whats_new What changed since a date, for one machine or all
can_i_run Tested, estimated or unknown — never a guess dressed as a fact
find_setups Setups ranked by published decode speed
records Throughput ceilings beaten recently
machines The hardware catalogue: slugs, memory, bandwidth
You describe your hardware in every call.
my_setup The machines and checkpoints this account synced
propose Queue a repo, paper or model for the next daily pass
my_proposals What you proposed, and what was done with it
submit_run Publish a tok/s you measured yourself, with the engine output behind it
my_runs Your runs, and what the checks said about each
And the five above stop asking. whats_new and
records read your setup by themselves, so you never describe your
hardware twice.
Tokens are minted one at a time, shown once, stored hashed and revoked
individually. The account travels in Authorization: Bearer and never in a
cookie — a page you visit cannot make requests as you.
Device count, engine, quantisation and context differ between runs. Figures exclude prefill and concurrent streams. A missing run is missing evidence, not incompatibility.
Find the right model for the task
Each score belongs to the model and setup it was measured in; a community recipe may use another quantisation or context. Memory fit is an estimate. No published run is missing evidence, not incompatibility.
Usage now, and where it is going
Area is tokens routed in the last 30 days; colour is those 7 days against the three weeks before. Both published by OpenRouter, never combined.
Every cell says what produced it. Scores from different evaluations are never averaged into one number, a
missing value is shown as missing and not as zero, and a run on more devices than you have is marked.
Have an agent watch this for you
The same updates answer over MCP — filtered by your setup,
each one carrying the reason it is here — so you can ask what changed for my
machines this week from wherever you already work, instead of coming to look.
The catalogue answers over MCP as well as on screen, so an agent
can ask what this page shows: what changed for your machines, whether a model fits one
of them, whether anyone just beat the best published figure for your box. With a token
it reads the setup above by itself — you never describe your hardware twice —
and it can queue something for the daily pass without leaving the conversation.
Things to try and keep for later.
This form is the way to reach whoever runs the
site. There is no account to create and no address to write to — that is on
purpose, so one form is the only thing to keep working.
Feedback
const D = {"hw":{"dgx-spark":{"nombre":"NVIDIA DGX Spark (GB10)","familia":"cuda","cc":121,"cc_txt":"sm_121","mem_gb":121,"bw":273,"precision":"fp4/fp8/bf16","nota":"Unified CPU/GPU memory, single node, no NVLink. Memory-rich and bandwidth-poor: holds a 126 GB checkpoint an 80 GB H100 cannot, and decodes at a fraction of the speed. Frequently excluded by arch gates written for sm_100 or sm_120 only."},"rtx-5090":{"nombre":"GeForce RTX 5090","familia":"cuda","cc":120,"cc_txt":"sm_120","mem_gb":32,"bw":1792,"precision":"fp4/fp8/bf16","nota":"Consumer Blackwell. Fast memory, little of it: a 32 GB budget makes quantization and KV-cache work matter more than raw kernel speed."},"rtx-5060-ti":{"nombre":"GeForce RTX 5060 Ti (16 GB)","familia":"cuda","cc":120,"cc_txt":"sm_120","mem_gb":16,"bw":448,"precision":"fp4/fp8/bf16","nota":"Entry-level Blackwell: 16 GB at 448 GB/s, and the cheapest way there is to get fp4 in hardware. Blackwell-gated kernels compile and run here, so an arch gate that opens is real news for it \u2014 and then the memory budget, a sixth of a workstation card's, decides whether the weights fit at all. fp4 quantization and KV-cache work are the entry ticket here, not an optimisation."},"rtx-pro-6000":{"nombre":"RTX PRO 6000 Blackwell","familia":"cuda","cc":120,"cc_txt":"sm_120","mem_gb":96,"bw":1792,"precision":"fp4/fp8/bf16","nota":"Workstation Blackwell, same arch as the 5090 but 96 GB. The one card that runs big models locally at speed."},"b200":{"nombre":"NVIDIA B200","familia":"cuda","cc":100,"cc_txt":"sm_100","mem_gb":180,"bw":8000,"precision":"fp4/fp8/bf16","nota":"Datacenter Blackwell. Usually the FIRST arch a new Blackwell kernel supports, and often the only one for a while."},"h200":{"nombre":"NVIDIA H200","familia":"cuda","cc":90,"cc_txt":"sm_90a","mem_gb":141,"bw":4800,"precision":"fp8/bf16","nota":"Hopper with more memory. No fp4. Still the most common serving target."},"h100":{"nombre":"NVIDIA H100","familia":"cuda","cc":90,"cc_tx
[truncated]
Automatic hardware interpretation · verify against source ${esc(p.wy)}
Extracted action · not a verified instruction for your installation ${esc(p.do)}
${esc

[truncated]
