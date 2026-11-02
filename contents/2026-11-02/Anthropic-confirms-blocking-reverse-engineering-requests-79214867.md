---
source: "https://www.tomshardware.com/tech-industry/artificial-intelligence/anthropic-confirms-new-claude-code-limits-to-prevent-game-cloning-users-report-opus-5-5-is-also-actively-blocking-reverse-engineering-tasks"
hn_url: "https://news.ycombinator.com/item?id=50339956"
title: "Anthropic confirms blocking reverse engineering requests"
article_title: "Anthropic confirms new Claude Code limits to prevent game cloning — users report Opus 5.5 is also actively blocking reverse engineering tasks | Tom's Hardware"
image: "https://cdn.mos.cms.futurecdn.net/qziRh5FuRAUAAnNS3QuCiU-2070-80.jpg"
author: "zof3"
captured_at: "2026-11-02T17:53:30Z"
capture_tool: "hn-digest"
hn_id: 50339956
score: 4
comments: 0
posted_at: "2026-11-02T17:07:28Z"
tags:
  - hacker-news
---

# Anthropic confirms blocking reverse engineering requests

- HN: [50339956](https://news.ycombinator.com/item?id=50339956)
- Source: [www.tomshardware.com](https://www.tomshardware.com/tech-industry/artificial-intelligence/anthropic-confirms-new-claude-code-limits-to-prevent-game-cloning-users-report-opus-5-5-is-also-actively-blocking-reverse-engineering-tasks)
- Score: 4
- Comments: 0
- Posted: 2026-11-02T17:07:28Z

## Translation

Title: Anthropic confirms blocking reverse engineering requests
Article title: Anthropic confirms new Claude Code limits to prevent game cloning — users report Opus 5.5 is also actively blocking reverse engineering tasks | Tom's Hardware
Description: Meanwhile, we now have direct reports of reverse engineering efforts being blocked, while online complaints grow.

Article text:
Skip to main content
Join Tom’s Hardware today
Upgrade to Tom’s Hardware Premium
Explore
GO PREMIUM
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
Get club Access Quick
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
Tom's Hardware
Stay On the Cutting Edge: Get the Tom's Hardware Newsletter
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Unlock instant access to exclusive member features.
By submitting your information you agree to the Terms & Conditions and Privacy Policy and are aged 16 or over.
You are now subscribed
Your newsletter sign-up was successful
Get full access to premium articles, exclusive features and a growing list of member rewards.
Anthropic confirms new Claude Code limits to prevent game cloning — users report Opus 5.5 is also actively blocking reverse engineering tasks
Meanwhile, we now have direct reports of reverse engineering efforts being blocked, while online complaints grow.
When you purchase through links on our site, we may earn an affiliate commission. Here’s how it works .
(Image credit: Getty Images)
Copy link
3
Join the conversation
Follow us
Add us as a preferred source on Google
Newsletter
Subscribe to our newsletter
Anthropic's recent move to have Claude Code deny video game cloning and possibly reverse engineering efforts has come under heavy fire from the software development community. Hot on the heels of Friday's report on the matter, an Anthropic employee responded to the tweet with the purported rule update with a direct statement confirming that the firm added guardrails to Claude Code to prevent game recreation.
In the full tweet, which you can read below, Addy Osmani says that "[Anthropic] added two sentences to Claude Code's prompt on Opus 5.5 to steer it away from recreating someone else's specific game because using Claude to infringe IP is against [its] usage policy." He goes on to state that work on original games is meant to continue working, and that Anthropic is iterating on the rules to try and "strike the right balance."
We added two sentences to Claude Code's prompt on Opus 5.5 to steer it away from recreating someone else's specific game because using Claude to infringe IP is against our usage policy.
Original games, genre games, and games built from your own designs and assets are meant to keep working.
We'll continue to iterate to ensure we're striking the right balance. In the meantime, if Claude turns down something that's yours, send it my way. October 30, 2026
There's no telling what those sentences might be, though given Osmani's choice of words, it's possible that said sentences are indeed what X user Chasm posted a screenshot of: "Do build original games, games in a genre, and games from the user’s own designs and assets. Don’t build a recreation of a specific existing game, even if it is renamed or called original."
In our previous report, we had already linked to recent issues raised at Claude Code's own GitHub repository, comprised of users complaining that the AI assistant was refusing to work on their games or on reverse-engineering tasks. As far as we can ascertain, it appears the game-related blocking is limited to the combination of using Claude Code and the Opus 5.5 model. Should that prove true, then users should be able to sidestep the limitation by calling upon Opus using another harness like Ollama or forwarding specific requests through a wrapper of some sort.
Regarding the blocking of reverse engineering, we have received direct confirmations. The first comes from our own alumnus Jarred Walton , who was trying to run checks on the software he's developing. He notes that such work had been performed with no issue until yesterday, October 31, and that switching to Opus 4.8 let him continue his work.
Anthropic's Claude AI blocks video game cloning and reverse-engineering
Anthropic says Claude thwarted bioweapon research from state-sponsored actors
Game decompilation hub freezes leaderboard after AI turns progress into a hollow numbers game amid explosion of 'AI slop'
Walton theorizes that Anthropic may have added a new rule to Code specifically for reverse-engineering. Considering that timing, plus the fact that other reports of refusals all popped up around the last week, may indicate that the company could be fiddling with cybersecurity-related rules in the same manner as they are for video game-related content.
We also received screenshots from a staffer's direct acquaintance, who asked Opus 5.5 (via Code) to try and reverse engineer a product activation, a request that likewise triggered safeguards. Meanwhile, notices that Claude Code appears to block reverse-engineering continue to mount, now with a Redditor allegedly having their account suspended , and then reinstated, over using Opus 5.5 and Reverse Engineer Anything . The company has an ongoing Cyber Verification Program , in which only approved entities get access to the full set of frontier models' cybersecurity capabilities.
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Bruno Ferreira
Social Links Navigation
Contributor
Bruno Ferreira is a contributing writer for Tom's Hardware. He has decades of experience with PC hardware and assorted sundries, alongside a career as a developer. He's obsessed with detail and has a tendency to ramble on the topics he loves. When not doing that, he's usually playing games, or at live music shows and festivals.
Artificial Intelligence
Anthropic says Claude thwarted bioweapon research from state-sponsored actors
Retro Gaming
Game decompilation hub freezes leaderboard after AI turns progress into a hollow numbers game amid explosion of 'AI slop'
Artificial Intelligence
Anthropic claims popular Chinese AI model has Mythos-class hacking abilities
Artificial Intelligence
Claude nukes a developer's 700 GB home directory while testing deletion safeguards; automatic model safety downgrade may have contributed to the screw-up
Artificial Intelligence
Chinese military researchers and tech giants caught using Claude
Artificial Intelligence
Russian freelancers use Claude to program autonomous combat drone swarm
Latest in Artificial Intelligence
Artificial Intelligence
Elon Musk mocks OpenAI CEO Sam Altman's claim that PCs having RAM is terrible
Artificial Intelligence
Brazilian lawyers busted for attempting to manipulate court's AI with prompt injection hidden in filing
Artificial Intelligence
AI agent swarms unleashed in Minecraft to ‘build a thriving civilization’
Artificial Intelligence
Firm that lets you rent out your idle Mac for AI inference workloads warns of plummeting token costs
Artificial Intelligence
Anthropic's Claude AI blocks video game cloning and reverse-engineering
Artificial Intelligence
AI agent used to perform controversial brain surgery on a mouse, Anthropic Sonnet controlled motorized surgery rig
Latest in News
Raspberry Pi
16GB Raspberry Pi 5 hiked to $385, more than triple its launch price — makers are already shifting to Pico/ESP32 for small projects, mini PCs for more ambitious tasks
Data Centers
'America hates data centers' — New data highlights public sentiment
Gaming Monitors
Gamer's ASUS TUF monitor blocks bullet fired through wall during Call of Duty session — says 280Hz monitor is 'the best $280 I ever spent'
Cybersecurity
Hackers claim to have stolen 23TB of data from Russian statistics body, including Ukraine war casualty figures
Artificial Intelligen
[truncated]
ThatMouse
OpenCode doesn't have this issue.
Reply
hotaru251
and this works for a finite time frame as chinese models will allow it and overtake current claude soon enoguh anyway.
Reply
Tom's Hardware is part of Future US Inc, an international media group and leading digital publisher. Visit our corporate site .
Add as a preferred source on Google
Terms and conditions
©
Future US, Inc. Full 7th Floor, 130 West 42nd Street,
New York,
NY 10036.

## Original Extract

Meanwhile, we now have direct reports of reverse engineering efforts being blocked, while online complaints grow.

Skip to main content
Join Tom’s Hardware today
Upgrade to Tom’s Hardware Premium
Explore
GO PREMIUM
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
Get club Access Quick
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
Tom's Hardware
Stay On the Cutting Edge: Get the Tom's Hardware Newsletter
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Unlock instant access to exclusive member features.
By submitting your information you agree to the Terms & Conditions and Privacy Policy and are aged 16 or over.
You are now subscribed
Your newsletter sign-up was successful
Get full access to premium articles, exclusive features and a growing list of member rewards.
Anthropic confirms new Claude Code limits to prevent game cloning — users report Opus 5.5 is also actively blocking reverse engineering tasks
Meanwhile, we now have direct reports of reverse engineering efforts being blocked, while online complaints grow.
When you purchase through links on our site, we may earn an affiliate commission. Here’s how it works .
(Image credit: Getty Images)
Copy link
3
Join the conversation
Follow us
Add us as a preferred source on Google
Newsletter
Subscribe to our newsletter
Anthropic's recent move to have Claude Code deny video game cloning and possibly reverse engineering efforts has come under heavy fire from the software development community. Hot on the heels of Friday's report on the matter, an Anthropic employee responded to the tweet with the purported rule update with a direct statement confirming that the firm added guardrails to Claude Code to prevent game recreation.
In the full tweet, which you can read below, Addy Osmani says that "[Anthropic] added two sentences to Claude Code's prompt on Opus 5.5 to steer it away from recreating someone else's specific game because using Claude to infringe IP is against [its] usage policy." He goes on to state that work on original games is meant to continue working, and that Anthropic is iterating on the rules to try and "strike the right balance."
We added two sentences to Claude Code's prompt on Opus 5.5 to steer it away from recreating someone else's specific game because using Claude to infringe IP is against our usage policy.
Original games, genre games, and games built from your own designs and assets are meant to keep working.
We'll continue to iterate to ensure we're striking the right balance. In the meantime, if Claude turns down something that's yours, send it my way. October 30, 2026
There's no telling what those sentences might be, though given Osmani's choice of words, it's possible that said sentences are indeed what X user Chasm posted a screenshot of: "Do build original games, games in a genre, and games from the user’s own designs and assets. Don’t build a recreation of a specific existing game, even if it is renamed or called original."
In our previous report, we had already linked to recent issues raised at Claude Code's own GitHub repository, comprised of users complaining that the AI assistant was refusing to work on their games or on reverse-engineering tasks. As far as we can ascertain, it appears the game-related blocking is limited to the combination of using Claude Code and the Opus 5.5 model. Should that prove true, then users should be able to sidestep the limitation by calling upon Opus using another harness like Ollama or forwarding specific requests through a wrapper of some sort.
Regarding the blocking of reverse engineering, we have received direct confirmations. The first comes from our own alumnus Jarred Walton , who was trying to run checks on the software he's developing. He notes that such work had been performed with no issue until yesterday, October 31, and that switching to Opus 4.8 let him continue his work.
Anthropic's Claude AI blocks video game cloning and reverse-engineering
Anthropic says Claude thwarted bioweapon research from state-sponsored actors
Game decompilation hub freezes leaderboard after AI turns progress into a hollow numbers game amid explosion of 'AI slop'
Walton theorizes that Anthropic may have added a new rule to Code specifically for reverse-engineering. Considering that timing, plus the fact that other reports of refusals all popped up around the last week, may indicate that the company could be fiddling with cybersecurity-related rules in the same manner as they are for video game-related content.
We also received screenshots from a staffer's direct acquaintance, who asked Opus 5.5 (via Code) to try and reverse engineer a product activation, a request that likewise triggered safeguards. Meanwhile, notices that Claude Code appears to block reverse-engineering continue to mount, now with a Redditor allegedly having their account suspended , and then reinstated, over using Opus 5.5 and Reverse Engineer Anything . The company has an ongoing Cyber Verification Program , in which only approved entities get access to the full set of frontier models' cybersecurity capabilities.
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Bruno Ferreira
Social Links Navigation
Contributor
Bruno Ferreira is a contributing writer for Tom's Hardware. He has decades of experience with PC hardware and assorted sundries, alongside a career as a developer. He's obsessed with detail and has a tendency to ramble on the topics he loves. When not doing that, he's usually playing games, or at live music shows and festivals.
Artificial Intelligence
Anthropic says Claude thwarted bioweapon research from state-sponsored actors
Retro Gaming
Game decompilation hub freezes leaderboard after AI turns progress into a hollow numbers game amid explosion of 'AI slop'
Artificial Intelligence
Anthropic claims popular Chinese AI model has Mythos-class hacking abilities
Artificial Intelligence
Claude nukes a developer's 700 GB home directory while testing deletion safeguards; automatic model safety downgrade may have contributed to the screw-up
Artificial Intelligence
Chinese military researchers and tech giants caught using Claude
Artificial Intelligence
Russian freelancers use Claude to program autonomous combat drone swarm
Latest in Artificial Intelligence
Artificial Intelligence
Elon Musk mocks OpenAI CEO Sam Altman's claim that PCs having RAM is terrible
Artificial Intelligence
Brazilian lawyers busted for attempting to manipulate court's AI with prompt injection hidden in filing
Artificial Intelligence
AI agent swarms unleashed in Minecraft to ‘build a thriving civilization’
Artificial Intelligence
Firm that lets you rent out your idle Mac for AI inference workloads warns of plummeting token costs
Artificial Intelligence
Anthropic's Claude AI blocks video game cloning and reverse-engineering
Artificial Intelligence
AI agent used to perform controversial brain surgery on a mouse, Anthropic Sonnet controlled motorized surgery rig
Latest in News
Raspberry Pi
16GB Raspberry Pi 5 hiked to $385, more than triple its launch price — makers are already shifting to Pico/ESP32 for small projects, mini PCs for more ambitious tasks
Data Centers
'America hates data centers' — New data highlights public sentiment
Gaming Monitors
Gamer's ASUS TUF monitor blocks bullet fired through wall during Call of Duty session — says 280Hz monitor is 'the best $280 I ever spent'
Cybersecurity
Hackers claim to have stolen 23TB of data from Russian statistics body, including Ukraine war casualty figures
Artificial Intelligen
[truncated]
ThatMouse
OpenCode doesn't have this issue.
Reply
hotaru251
and this works for a finite time frame as chinese models will allow it and overtake current claude soon enoguh anyway.
Reply
Tom's Hardware is part of Future US Inc, an international media group and leading digital publisher. Visit our corporate site .
Add as a preferred source on Google
Terms and conditions
©
Future US, Inc. Full 7th Floor, 130 West 42nd Street,
New York,
NY 10036.
