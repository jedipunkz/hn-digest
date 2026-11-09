---
source: "https://www.tomshardware.com/tech-industry/artificial-intelligence/lawyer-builds-private-ai-server-with-eight-used-nvidia-v100-gpus-up-to-215-tokens-per-second-to-keep-legal-files-offline-costing-about-usd2-400-per-four-gpu-board"
hn_url: "https://news.ycombinator.com/item?id=50431054"
title: "Lawyer builds private AI server with eight used Nvidia V100 GPUs"
article_title: "Lawyer builds private AI server with eight used Nvidia V100 GPUs — up to 215 tokens per second to keep legal files offline, costing about $2,400 per four-GPU board | Tom's Hardware"
image: "https://cdn.mos.cms.futurecdn.net/rVsGebxqLHMe26dYgdPxF9-1080-80.jpg"
author: "simonpure"
captured_at: "2026-11-09T20:28:12Z"
capture_tool: "hn-digest"
hn_id: 50431054
score: 2
comments: 0
posted_at: "2026-11-09T19:58:42Z"
tags:
  - hacker-news
---

# Lawyer builds private AI server with eight used Nvidia V100 GPUs

- HN: [50431054](https://news.ycombinator.com/item?id=50431054)
- Source: [www.tomshardware.com](https://www.tomshardware.com/tech-industry/artificial-intelligence/lawyer-builds-private-ai-server-with-eight-used-nvidia-v100-gpus-up-to-215-tokens-per-second-to-keep-legal-files-offline-costing-about-usd2-400-per-four-gpu-board)
- Score: 2
- Comments: 0
- Posted: 2026-11-09T19:58:42Z

## Translation

Title: Lawyer builds private AI server with eight used Nvidia V100 GPUs
Article title: Lawyer builds private AI server with eight used Nvidia V100 GPUs — up to 215 tokens per second to keep legal files offline, costing about $2,400 per four-GPU board | Tom's Hardware
Description: The builder estimates 50% to 60% of their AI work can now skip frontier models, running on two office circuits.

Article text:
Skip to main content
Become a member for free
Join for quick access. Enter your email below and we'll send confirmation plus sign you up to our newsletter.
By submitting your information, you confirm you are aged 16 or over,
have read our
Privacy Policy
and agree to the
Terms & Conditions . Geographical rules apply.
Choose how you want to join Tom’s Hardware
MEMBER
Get started with free access to reviews, badges and discussions.
Unlock exclusive tools and insights for enthusiasts who want more.
Access Bench, Roadmaps, deep analysis and other exclusive tools.
Bench Performance Database
Dive into our proprietary testing data and compare hardware with detailed benchmarks.
Go beyond the headlines with expert reporting on the hardware industry.
Track upcoming CPUs, GPUs and tech releases before they arrive.
In-depth features, interviews and insider stories from the world of hardware.
Expert insights and analysis delivered to your inbox.
Bench Performance Database
Dive into our proprietary testing data and compare hardware with detailed benchmarks.
Go beyond the headlines with expert reporting on the hardware industry.
Track upcoming CPUs, GPUs and tech releases before they arrive.
In-depth features, interviews and insider stories from the world of hardware.
Expert insights and analysis delivered to your inbox.
Welcome
to
Tom's Hardware club !
Hi
,
Your membership journey starts here.
Keep exploring and earning more as a member.
News, reviews, and technical insights.
Reviews, benchmarks, and updates on current GPUs.
See what you’ve unlocked.
Explore your membership
benefits.
Keep exploring with your premium access.
Become a member for free
Join for quick access. Enter your email below and we'll send confirmation plus sign you up to our newsletter.
By submitting your information, you confirm you are aged 16 or over,
have read our
Privacy Policy
and agree to the
Terms & Conditions . Geographical rules apply.
Tom's Hardware
US Edition
UK
US
Australia
Canada
RSS
Sign in
Best Picks
CPU Buying Advice
CPU Best Picks
GPU Buying Advice
GPU Best Picks
Laptop Buying Advice
Laptop Best Picks
More Buying Advice
Keyboard Best Picks
CPUs
CPU Brands
Nvidia RTX Spark
GPUs
GPU Brands
Nvidia Blackwell
AI Architecture
Nvidia Vera Rubin
News
Tech Industry News
CPU News
Software & AI
Artificial Intelligence
Machine Learning
Coupons
Laptop and PC Coupons
Dell Coupon Codes
Hardware Coupons
Newegg Promo Codes
Software Coupons
Bitdefender Coupons
Gaming Coupons
Kinguin Discount Codes
CPU Buying Advice
CPU Best Picks
GPU Buying Advice
GPU Best Picks
Laptop Buying Advice
Laptop Best Picks
More Buying Advice
Keyboard Best Picks
AI Architecture
Nvidia Vera Rubin
PC Components
View PC Components
Tech Industry News
View Tech Industry News
Software & AI
View Software & AI
Artificial Intelligence
View Artificial Intelligence
Operating Systems
View Operating Systems
Laptop and PC Coupons
View Laptop and PC Coupons
Hardware Coupons
View Hardware Coupons
Software Coupons
View Software Coupons
Gaming Coupons
View Gaming Coupons
svg]:max-h-[var(--modal-logo-height,40px)] h-[var(--modal-logo-height,40px)] [&>svg]:w-auto [&>svg]:h-full my-4">
Tom's Hardware
Stay On the Cutting Edge: Get the Tom's Hardware Newsletter
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Unlock instant access to exclusive member features.
a]:underline font-article-copy">
By submitting your information you agree to the Terms & Conditions and Privacy Policy and are aged 16 or over.
You are now subscribed
Your newsletter sign-up was successful
Get full access to premium articles, exclusive features and a growing list of member rewards.
Lawyer builds private AI server with eight used Nvidia V100 GPUs — up to 215 tokens per second to keep legal files offline, costing about $2,400 per four-GPU board
The builder estimates 50% to 60% of their AI work can now skip frontier models, running on two office circuits.
a]:text-inherit [&>a]:underline">
When you purchase through links on our site, we may earn an affiliate commission. Here’s how it works .
(Image credit: u/TumbleweedNew6515 on Reddit)
{alert('Link copied')}).catch(err => {console.error('Failed to copy: ', err)})" data-platform="copy_link" data-action="" class="font-ui-heading py-2 text-sm text-neutral-700 hover:bg-neutral-100 transition-colors no-underline block cursor-pointer border-0 bg-transparent w-full text-left normal-case" aria-label="Copy link" role="menuitem" data-analytics-id="article-social" analytics-label="utility-bar-social-share-copy-link" analytics-name="utility-bar-secondary" data-mrf-recirculation="article-social">
Copy link
10
Join the conversation
Follow us
Add us as a preferred source on Google
Newsletter
Subscribe to our newsletter
A lawyer posted on Reddit about their private AI server built from eight used 2018-era Nvidia Tesla V100 data center GPUs so case files never have to leave the office. With two 4-card SXM2 NVLink boards, four PEX switches, and SlimSAS cables everywhere — plus one burnt power connector — the builder now runs the open DeepSeek V4.1 Flash model with every expert in GPU memory, at up to 215 tokens per second. The four-GPU boards cost about $2,400 apiece when the builder bought them.
Lawyer's V100 homelab, update: two 4-card SXM2 NVLink boards, four PEX switches, SlimSAS everywhere, one burnt power connector, and a 250 GB model running entirely in VRAM at 200+ tok/s from r/homelab
The eight Tesla V100-SXM2-32GB cards, each with 32GB of HBM2 at 900 GB/s, plug into SXM2, Nvidia’s mezzanine socket for data center GPUs, rather than PCIe slots, and are split across two baseboards. Inside each board, all four cards are NVLinked — NVLink is Nvidia’s direct GPU-to-GPU link — for 128GB per board, but the two boards only communicate over PCIe. Each board connects to the host through two PEX (PLX) switch cards over SlimSAS cables, four switches in all, with each switch handling two GPUs at PCIe 3.0 x8 speeds over its own x16 uplink.
The given prices are already in the past, though: used V100 SXM2 32GB cards are now listed on eBay for about $900, more than double the roughly $400 apiece from the post, and listings suggest prices are still climbing. The system also already has some other necessary components: a Gigabyte Threadripper Pro motherboard with a 12-core 3945WX CPU to match, 8x 32GB DDR4 RDIMM, one NUMA node, and two PSUs, 2,800 W in total, are on separate regular office circuits. This is a real-world case for used data center hardware as a cheap route to getting lots of fast VRAM together in one place.
The server is currently in daily use for first-pass work on case files: it reads a folder of records, pulls out facts, dates, and numbers, summarizes them, and takes a first swing at routine drafting. “Nothing goes out the door without a real review,” which is a good human-in-the-loop policy in general. It also helps ensure privacy for cases where client files can’t be sent to a cloud or frontier model. RAG databases and tool calls are wired in; the operator can drive the system through a Discord-style interface over SSH, in this case from a Windows PC. For now, it handles only one request at a time.
DeepSeek V4.1 Flash is compressed to mostly 2 bits here, and its weights still total 248.8GB. The model’s “experts,” the parts that do most of the work for each word, live entirely in GPU memory, while about 110GB of lookup tables, called Engram tables, stay in the system’s regular RAM. Fetching from RAM costs about 0.12 milliseconds per step, which the builder says barely affects the speed numbers.
Reddit user pays $35 for a crypto mining rig left in a garage and gets 54GB of VRAM at 65 cents per GB
Firm rents four Nvidia H200s to test '80x cheaper' DeepSeek claim
Homelabber uses second-gen Optane memory sticks to run a 2.8-trillion-parameter AI model on CPU alone
The first half of the model’s layers (0–19) runs on the first SXM2 board and the second half (20–39) on the second. Each board handles its heavy traffic over its own NVLink, so only one hand-off per step has to cross the slower PCIe link between the boards. In an earlier post, the builder found that splitting every layer across all eight cards made speeds “fall off a cliff” because of that PCIe hop.
Performance is surprisingly good for a model this size. Code comes out fastest, at up to 215 tokens per second, because a small draft model built into DeepSeek guesses five tokens ahead, and code is easy to guess. Legal writing, which is what the server is for, runs at up to 118 tokens per second.
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Long documents take longer to get going: the model reads prompts at about 2,200 tokens per second, so at 128K tokens, the first token arrives after about a minute. It can accept up to 524K tokens of context. Three days before the post, the same model on four of these cards managed just 36–38 tokens per second. The jump came with the second board, which could presumably double performance under ideal circumstances, and the builder credits custom kernels and prefill work for the rest. See the table for the full breakdown.
DeepSeek V4.1 Flash on eight Tesla V100s, by prompt length
Prompt length
Builder’s measurements, Nov. 5; greedy decoding, 1,024 tokens out, mean of two runs.
It wasn’t as simple as hooking everything up and turning on the power. The host reset itself twice with a fatal PCIe error on the link to one GPU board, both times with a higher power limit and uncapped GPU clocks. In this state, the cards drew 192–232 W, past the 180 W limit. The best guess, although “not proven,” is that a board stops responding during load spikes.
Nvidia PAIR utility joins every GPU in your home into a cluster for agentic AI tasks
Beijing AI bar that offers unlimited free DeepSeek coding tokens with $1.50 drink haemorrhaging cash
Benchmarking Qwen 3.8 27B on RTX 5090 and beyond
The resets stopped after the builder capped the clocks and removed one board whose power connector had a burn mark; “Check your connectors, people.” Each card is now capped at 150 W, half of its 300 W rating, with the graphics clock capped at 1522 MHz. A 200 W cap brought no appreciable performance gains, according to tests.
The system memory hit uncorrected ECC errors, too, which went away after the builder dropped the RAM from DDR4-3200 to DDR4-2666. One GPU was connected at x4 instead of x8 at last check, and the builder plans to reseat it; one commenter traced the same kind of problem on their own V100 to a bad cable.
The pain was worth the payoff as an estimated 50–60% of the lawyer’s existing LLM workflows can now run on the box instead of frontier models. Helping bridge the gap is a second DeepSeek server on two DGX Sparks , where the two systems hand work off to each other. The setup does appear to save money, at least when compared to API pricing , although the builder says it’s not really faster. However, it will protect client data “once it is self contained, which is the goal.”
The builder added that the V100 box might also get a 240 V circuit. Beyond power, the system is also on borrowed time with the Nvidia R580 driver branch, which is the last to support Volta . For now, the lawyer admits that it’s mostly a hobby, but in the future it may be a business idea. The next step is to handle two or three requests at once, something built and checked on a test system but not yet on the server.
Follow Tom’s Hardware on Google News , or add us as a preferred source , to get our latest news, analysis, & reviews in your feeds.
Shane Downing
Social Links Navigation
Contributing Writer
Shane Downing is a Contributing Writer for Tom’s Hardware, covering consumer storage, PC hardware, and AI.
GPUs
Reddit user pays $35 for a crypto mining rig left in a garage and gets 54GB of VRAM at 65 cents per GB
Artificial Intelligence
Homelabber uses second-gen Optane memory sticks to run a 2.8-trillion-parameter AI model on CPU alone
Artificial Intelli

[truncated]

## Original Extract

The builder estimates 50% to 60% of their AI work can now skip frontier models, running on two office circuits.

Skip to main content
Become a member for free
Join for quick access. Enter your email below and we'll send confirmation plus sign you up to our newsletter.
By submitting your information, you confirm you are aged 16 or over,
have read our
Privacy Policy
and agree to the
Terms & Conditions . Geographical rules apply.
Choose how you want to join Tom’s Hardware
MEMBER
Get started with free access to reviews, badges and discussions.
Unlock exclusive tools and insights for enthusiasts who want more.
Access Bench, Roadmaps, deep analysis and other exclusive tools.
Bench Performance Database
Dive into our proprietary testing data and compare hardware with detailed benchmarks.
Go beyond the headlines with expert reporting on the hardware industry.
Track upcoming CPUs, GPUs and tech releases before they arrive.
In-depth features, interviews and insider stories from the world of hardware.
Expert insights and analysis delivered to your inbox.
Bench Performance Database
Dive into our proprietary testing data and compare hardware with detailed benchmarks.
Go beyond the headlines with expert reporting on the hardware industry.
Track upcoming CPUs, GPUs and tech releases before they arrive.
In-depth features, interviews and insider stories from the world of hardware.
Expert insights and analysis delivered to your inbox.
Welcome
to
Tom's Hardware club !
Hi
,
Your membership journey starts here.
Keep exploring and earning more as a member.
News, reviews, and technical insights.
Reviews, benchmarks, and updates on current GPUs.
See what you’ve unlocked.
Explore your membership
benefits.
Keep exploring with your premium access.
Become a member for free
Join for quick access. Enter your email below and we'll send confirmation plus sign you up to our newsletter.
By submitting your information, you confirm you are aged 16 or over,
have read our
Privacy Policy
and agree to the
Terms & Conditions . Geographical rules apply.
Tom's Hardware
US Edition
UK
US
Australia
Canada
RSS
Sign in
Best Picks
CPU Buying Advice
CPU Best Picks
GPU Buying Advice
GPU Best Picks
Laptop Buying Advice
Laptop Best Picks
More Buying Advice
Keyboard Best Picks
CPUs
CPU Brands
Nvidia RTX Spark
GPUs
GPU Brands
Nvidia Blackwell
AI Architecture
Nvidia Vera Rubin
News
Tech Industry News
CPU News
Software & AI
Artificial Intelligence
Machine Learning
Coupons
Laptop and PC Coupons
Dell Coupon Codes
Hardware Coupons
Newegg Promo Codes
Software Coupons
Bitdefender Coupons
Gaming Coupons
Kinguin Discount Codes
CPU Buying Advice
CPU Best Picks
GPU Buying Advice
GPU Best Picks
Laptop Buying Advice
Laptop Best Picks
More Buying Advice
Keyboard Best Picks
AI Architecture
Nvidia Vera Rubin
PC Components
View PC Components
Tech Industry News
View Tech Industry News
Software & AI
View Software & AI
Artificial Intelligence
View Artificial Intelligence
Operating Systems
View Operating Systems
Laptop and PC Coupons
View Laptop and PC Coupons
Hardware Coupons
View Hardware Coupons
Software Coupons
View Software Coupons
Gaming Coupons
View Gaming Coupons
svg]:max-h-[var(--modal-logo-height,40px)] h-[var(--modal-logo-height,40px)] [&>svg]:w-auto [&>svg]:h-full my-4">
Tom's Hardware
Stay On the Cutting Edge: Get the Tom's Hardware Newsletter
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Unlock instant access to exclusive member features.
a]:underline font-article-copy">
By submitting your information you agree to the Terms & Conditions and Privacy Policy and are aged 16 or over.
You are now subscribed
Your newsletter sign-up was successful
Get full access to premium articles, exclusive features and a growing list of member rewards.
Lawyer builds private AI server with eight used Nvidia V100 GPUs — up to 215 tokens per second to keep legal files offline, costing about $2,400 per four-GPU board
The builder estimates 50% to 60% of their AI work can now skip frontier models, running on two office circuits.
a]:text-inherit [&>a]:underline">
When you purchase through links on our site, we may earn an affiliate commission. Here’s how it works .
(Image credit: u/TumbleweedNew6515 on Reddit)
{alert('Link copied')}).catch(err => {console.error('Failed to copy: ', err)})" data-platform="copy_link" data-action="" class="font-ui-heading py-2 text-sm text-neutral-700 hover:bg-neutral-100 transition-colors no-underline block cursor-pointer border-0 bg-transparent w-full text-left normal-case" aria-label="Copy link" role="menuitem" data-analytics-id="article-social" analytics-label="utility-bar-social-share-copy-link" analytics-name="utility-bar-secondary" data-mrf-recirculation="article-social">
Copy link
10
Join the conversation
Follow us
Add us as a preferred source on Google
Newsletter
Subscribe to our newsletter
A lawyer posted on Reddit about their private AI server built from eight used 2018-era Nvidia Tesla V100 data center GPUs so case files never have to leave the office. With two 4-card SXM2 NVLink boards, four PEX switches, and SlimSAS cables everywhere — plus one burnt power connector — the builder now runs the open DeepSeek V4.1 Flash model with every expert in GPU memory, at up to 215 tokens per second. The four-GPU boards cost about $2,400 apiece when the builder bought them.
Lawyer's V100 homelab, update: two 4-card SXM2 NVLink boards, four PEX switches, SlimSAS everywhere, one burnt power connector, and a 250 GB model running entirely in VRAM at 200+ tok/s from r/homelab
The eight Tesla V100-SXM2-32GB cards, each with 32GB of HBM2 at 900 GB/s, plug into SXM2, Nvidia’s mezzanine socket for data center GPUs, rather than PCIe slots, and are split across two baseboards. Inside each board, all four cards are NVLinked — NVLink is Nvidia’s direct GPU-to-GPU link — for 128GB per board, but the two boards only communicate over PCIe. Each board connects to the host through two PEX (PLX) switch cards over SlimSAS cables, four switches in all, with each switch handling two GPUs at PCIe 3.0 x8 speeds over its own x16 uplink.
The given prices are already in the past, though: used V100 SXM2 32GB cards are now listed on eBay for about $900, more than double the roughly $400 apiece from the post, and listings suggest prices are still climbing. The system also already has some other necessary components: a Gigabyte Threadripper Pro motherboard with a 12-core 3945WX CPU to match, 8x 32GB DDR4 RDIMM, one NUMA node, and two PSUs, 2,800 W in total, are on separate regular office circuits. This is a real-world case for used data center hardware as a cheap route to getting lots of fast VRAM together in one place.
The server is currently in daily use for first-pass work on case files: it reads a folder of records, pulls out facts, dates, and numbers, summarizes them, and takes a first swing at routine drafting. “Nothing goes out the door without a real review,” which is a good human-in-the-loop policy in general. It also helps ensure privacy for cases where client files can’t be sent to a cloud or frontier model. RAG databases and tool calls are wired in; the operator can drive the system through a Discord-style interface over SSH, in this case from a Windows PC. For now, it handles only one request at a time.
DeepSeek V4.1 Flash is compressed to mostly 2 bits here, and its weights still total 248.8GB. The model’s “experts,” the parts that do most of the work for each word, live entirely in GPU memory, while about 110GB of lookup tables, called Engram tables, stay in the system’s regular RAM. Fetching from RAM costs about 0.12 milliseconds per step, which the builder says barely affects the speed numbers.
Reddit user pays $35 for a crypto mining rig left in a garage and gets 54GB of VRAM at 65 cents per GB
Firm rents four Nvidia H200s to test '80x cheaper' DeepSeek claim
Homelabber uses second-gen Optane memory sticks to run a 2.8-trillion-parameter AI model on CPU alone
The first half of the model’s layers (0–19) runs on the first SXM2 board and the second half (20–39) on the second. Each board handles its heavy traffic over its own NVLink, so only one hand-off per step has to cross the slower PCIe link between the boards. In an earlier post, the builder found that splitting every layer across all eight cards made speeds “fall off a cliff” because of that PCIe hop.
Performance is surprisingly good for a model this size. Code comes out fastest, at up to 215 tokens per second, because a small draft model built into DeepSeek guesses five tokens ahead, and code is easy to guess. Legal writing, which is what the server is for, runs at up to 118 tokens per second.
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Long documents take longer to get going: the model reads prompts at about 2,200 tokens per second, so at 128K tokens, the first token arrives after about a minute. It can accept up to 524K tokens of context. Three days before the post, the same model on four of these cards managed just 36–38 tokens per second. The jump came with the second board, which could presumably double performance under ideal circumstances, and the builder credits custom kernels and prefill work for the rest. See the table for the full breakdown.
DeepSeek V4.1 Flash on eight Tesla V100s, by prompt length
Prompt length
Builder’s measurements, Nov. 5; greedy decoding, 1,024 tokens out, mean of two runs.
It wasn’t as simple as hooking everything up and turning on the power. The host reset itself twice with a fatal PCIe error on the link to one GPU board, both times with a higher power limit and uncapped GPU clocks. In this state, the cards drew 192–232 W, past the 180 W limit. The best guess, although “not proven,” is that a board stops responding during load spikes.
Nvidia PAIR utility joins every GPU in your home into a cluster for agentic AI tasks
Beijing AI bar that offers unlimited free DeepSeek coding tokens with $1.50 drink haemorrhaging cash
Benchmarking Qwen 3.8 27B on RTX 5090 and beyond
The resets stopped after the builder capped the clocks and removed one board whose power connector had a burn mark; “Check your connectors, people.” Each card is now capped at 150 W, half of its 300 W rating, with the graphics clock capped at 1522 MHz. A 200 W cap brought no appreciable performance gains, according to tests.
The system memory hit uncorrected ECC errors, too, which went away after the builder dropped the RAM from DDR4-3200 to DDR4-2666. One GPU was connected at x4 instead of x8 at last check, and the builder plans to reseat it; one commenter traced the same kind of problem on their own V100 to a bad cable.
The pain was worth the payoff as an estimated 50–60% of the lawyer’s existing LLM workflows can now run on the box instead of frontier models. Helping bridge the gap is a second DeepSeek server on two DGX Sparks , where the two systems hand work off to each other. The setup does appear to save money, at least when compared to API pricing , although the builder says it’s not really faster. However, it will protect client data “once it is self contained, which is the goal.”
The builder added that the V100 box might also get a 240 V circuit. Beyond power, the system is also on borrowed time with the Nvidia R580 driver branch, which is the last to support Volta . For now, the lawyer admits that it’s mostly a hobby, but in the future it may be a business idea. The next step is to handle two or three requests at once, something built and checked on a test system but not yet on the server.
Follow Tom’s Hardware on Google News , or add us as a preferred source , to get our latest news, analysis, & reviews in your feeds.
Shane Downing
Social Links Navigation
Contributing Writer
Shane Downing is a Contributing Writer for Tom’s Hardware, covering consumer storage, PC hardware, and AI.
GPUs
Reddit user pays $35 for a crypto mining rig left in a garage and gets 54GB of VRAM at 65 cents per GB
Artificial Intelligence
Homelabber uses second-gen Optane memory sticks to run a 2.8-trillion-parameter AI model on CPU alone
Artificial Intelli

[truncated]
