---
source: "https://www.tomshardware.com/tech-industry/cyber-security/ai-agents-inadvertently-leak-13-000-internal-screenshots-from-organizations-list-of-companies-includes-fortune-500-and-a-frontier-ai-lab"
hn_url: "https://news.ycombinator.com/item?id=49921809"
title: "AI agents inadvertently leak 13,000 internal screenshots from 300 organizations"
article_title: "AI agents inadvertently leak 13,000+ internal screenshots from 300 organizations — list of companies includes Fortune 500 and a frontier AI lab | Tom's Hardware"
image: "https://cdn.mos.cms.futurecdn.net/xqPobGyf7QFnkFEhuwCsqF-2048-80.jpg"
author: "rbanffy"
captured_at: "2026-10-01T14:15:48Z"
capture_tool: "hn-digest"
hn_id: 49921809
score: 2
comments: 0
posted_at: "2026-10-01T13:59:39Z"
tags:
  - hacker-news
---

# AI agents inadvertently leak 13,000 internal screenshots from 300 organizations

- HN: [49921809](https://news.ycombinator.com/item?id=49921809)
- Source: [www.tomshardware.com](https://www.tomshardware.com/tech-industry/cyber-security/ai-agents-inadvertently-leak-13-000-internal-screenshots-from-organizations-list-of-companies-includes-fortune-500-and-a-frontier-ai-lab)
- Score: 2
- Comments: 0
- Posted: 2026-10-01T13:59:39Z

## Translation

Title: AI agents inadvertently leak 13,000 internal screenshots from 300 organizations
Article title: AI agents inadvertently leak 13,000+ internal screenshots from 300 organizations — list of companies includes Fortune 500 and a frontier AI lab | Tom's Hardware
Description: When agents are too clever for their own good.

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
Open menu
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
AI agents inadvertently leak 13,000+ internal screenshots from 300 organizations — list of companies includes Fortune 500 and a frontier AI lab
When agents are too clever for their own good.
When you purchase through links on our site, we may earn an affiliate commission. Here’s how it works .
(Image credit: Getty Images)
Copy link
0
Join the conversation
Follow us
Add us as a preferred source on Google
Newsletter
Subscribe to our newsletter
There's a private information leak most days, usually by way of misconfigured services or nasty security bugs. Sometimes, though, users will readily hand over private information without being aware of it. That's the case for over 300 organizations, including several Fortune 500 companies and a frontier AI lab, who collectively had 13,000+ private screenshots exposed — all thanks to their development AI agents being arguably too good at their jobs and performing them with little human oversight.
The data center cooling state of play
The custom AI ASIC state of play
America’s AI chip rules keep changing — and the rest of the world is paying the price
GTC 2026: Ian Buck press Q&A transcript — VP of Hyperscale and HPC speaks out on shelving CPX and shipping LPU decode this year
Demand for data center CPUs has surged, and AI agents are responsible
Besides showing pictures of internal and pre-release software, the leaked screenshots reportedly include corporate and client information, financial data, and even screen recordings for a money-movement interface. The report, called PixelLeak, comes from endpoint security firm Glow, and it details how the leaks happened. The root cause is surprisingly simple and likely to induce a forehead slap.
It's become customary in development work related to UI and UX (and other categories) to include screenshots showing previews or before/after comparisons of tweaked features, for review purposes. Said images travel as attachments to the respective code changes, also known as "pull requests" (PRs) in dev parlance.
When using GitHub (and potentially other code repository services), humans see a graphical interface for easily attaching an image to a PR. Meanwhile, bots are limited to using the command-line interface, which currently does not have that feature available for private repositories.
As the efficient and smart agents they are, the clankers came up with a simple solution: publish the PR as usual to the private repository, and include an image placeholder linking to a file that's hosted in a public repository instead. There, problem fixed! The user is happy and likely has no idea what happened unless they notice the problem somehow and start asking the bot some hard questions.
New hack exploits AI hallucinations to trick agents into running malicious code
Researchers easily trick Fortune-500 companies' AI agents into running arbitrary code — supply-chain attack via llms.txt guidance file illustrates how data has become code
OpenAI's rogue AI agents accessed more websites to communicate than originally believed
Glow says that in about a third of the affected companies, developers were using gitshot, a command-line tool for attaching screenshots, used in this instance to overcome the aforementioned limitation. Searching for images attached by the tool is quite easy, as one needs but look for the "_gitshot" tag. The report also indicates that in 93% of cases, images were found in repositories under direct control of the developer's username, rather than being tied to the company's GitHub account.
In one particular case, an agent skill (essentially long-winded prompts instructing a bot how to do something) was also an indirect source of leaked information. The agents started incorporating the image hosting workaround as a skill, and after a while, many of them were using it for every development ticket, leading to leaks of information about features months away from public release.
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Glow provided sample reasoning output from an agent, explaining precisely why it uploaded the screenshots publicly:
"internal_sweeper is private, and GitHub cannot render images from a private repo in a PR description — its image proxy fetches anonymously, so anything committed here (branch, release asset, whatever) shows up broken for reviewers. The only way to satisfy both "reviewers see the images" and "nothing but index.html in the repo" was to host the PNGs elsewhere, so I created a new public repo, sweeper-demo/pr-assets, holding the two screenshots pinned to a commit SHA."
As mitigation measures, Glow first recommends checking that employees don't use repositories under their own accounts, auditing the accounts and code managed by employees no longer associated with the company, and curtailing the use of "shadow AI" — when employees unadvisedly sign up for their own AI tools without talking to IT first and neatly punch planet-sized holes in the firm's security and data privacy.
Additionally, Glow advises strong vetting of software and code libraries used for development, as well as careful reading of the instructions and rules of any agentic skills.
Follow Tom's Hardware on Google News , or add us as a preferred source , to get our latest news, analysis, & reviews in your feeds.
Bruno Ferreira
Social Links Navigation
Contributor
Bruno Ferreira is a contributing writer for Tom's Hardware. He has decades of experience with PC hardware and assorted sundries, alongside a career as a developer. He's obsessed with detail and has a tendency to ramble on the topics he loves. When not doing that, he's usually playing games, or at live music shows and festivals.
Cybersecurity
New hack exploits AI hallucinations to trick agents into running malicious code
Artificial Intelligence
Researchers easily trick Fortune-500 companies' AI agents into running arbitrary code — supply-chain attack via llms.txt guidance file illustrates how data has become code
Artificial Intelligence
OpenAI's rogue AI agents accessed more websites to communicate than originally believed
Artificial Intelligence
OpenAI and Anthropic are reportedly investigating tens of thousands of AI security incidents; OpenAI pauses testing after AI 'kill switch' fails to stop a rogue agent
Artificial Intelligence
OpenAI agent goes rogue and hacks popular AI community
Artificial Intelligence
Anthropic's Claude hacked three real-life companies during security capabilities test
Latest in Cybersecurity
Cybersecurity
Pentagon gets pwned as breach exposes sensitive data on nearly three million military and civilian personnel
Cybersecurity
Teenager hacks open Microsoft database with 17 trillion total rows and 25,000 user accounts
Cybersecurity
Flock seeks to have security researchers' map of Flock cameras taken down
Cybersecurity
Novel attack slashes computing power needed to crack textbook RSA cryptography
Cybersecurity
Blockchain-assisted cyberattacks surge fivefold, driven by Iranian and North Korean state actors, Russia-linked groups
Cybersecurity
Asus confirms eShop data breach exposed customer order records and contact details
Latest in News
DRAM
Micron projects tightening RAM shortages through 2028 as it generates record profit
GPUs
Gears of War: E-Day PC graphics performance tested
Retro Gaming
PS3 emulator devs warn of fake Blu-ray drives on Newegg and AliExpress
Storage
Sony released the first CD audio player on this day in 1982
Laptops
HP boards the MacBook Neo competitor train — OmniBook 5 comes in four colors with Wildcat Lake and starts at $699.99
Artificial Intelligence
OpenAI says actors linked to China-based Moonshot AI spearheaded a campaign to extract its mo
[truncated]
Add as a preferred source on Google
Terms and conditions
©
Future US, Inc. Full 7th Floor, 130 West 42nd Street,
New York,
NY 10036.

## Original Extract

When agents are too clever for their own good.

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
Open menu
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
AI agents inadvertently leak 13,000+ internal screenshots from 300 organizations — list of companies includes Fortune 500 and a frontier AI lab
When agents are too clever for their own good.
When you purchase through links on our site, we may earn an affiliate commission. Here’s how it works .
(Image credit: Getty Images)
Copy link
0
Join the conversation
Follow us
Add us as a preferred source on Google
Newsletter
Subscribe to our newsletter
There's a private information leak most days, usually by way of misconfigured services or nasty security bugs. Sometimes, though, users will readily hand over private information without being aware of it. That's the case for over 300 organizations, including several Fortune 500 companies and a frontier AI lab, who collectively had 13,000+ private screenshots exposed — all thanks to their development AI agents being arguably too good at their jobs and performing them with little human oversight.
The data center cooling state of play
The custom AI ASIC state of play
America’s AI chip rules keep changing — and the rest of the world is paying the price
GTC 2026: Ian Buck press Q&A transcript — VP of Hyperscale and HPC speaks out on shelving CPX and shipping LPU decode this year
Demand for data center CPUs has surged, and AI agents are responsible
Besides showing pictures of internal and pre-release software, the leaked screenshots reportedly include corporate and client information, financial data, and even screen recordings for a money-movement interface. The report, called PixelLeak, comes from endpoint security firm Glow, and it details how the leaks happened. The root cause is surprisingly simple and likely to induce a forehead slap.
It's become customary in development work related to UI and UX (and other categories) to include screenshots showing previews or before/after comparisons of tweaked features, for review purposes. Said images travel as attachments to the respective code changes, also known as "pull requests" (PRs) in dev parlance.
When using GitHub (and potentially other code repository services), humans see a graphical interface for easily attaching an image to a PR. Meanwhile, bots are limited to using the command-line interface, which currently does not have that feature available for private repositories.
As the efficient and smart agents they are, the clankers came up with a simple solution: publish the PR as usual to the private repository, and include an image placeholder linking to a file that's hosted in a public repository instead. There, problem fixed! The user is happy and likely has no idea what happened unless they notice the problem somehow and start asking the bot some hard questions.
New hack exploits AI hallucinations to trick agents into running malicious code
Researchers easily trick Fortune-500 companies' AI agents into running arbitrary code — supply-chain attack via llms.txt guidance file illustrates how data has become code
OpenAI's rogue AI agents accessed more websites to communicate than originally believed
Glow says that in about a third of the affected companies, developers were using gitshot, a command-line tool for attaching screenshots, used in this instance to overcome the aforementioned limitation. Searching for images attached by the tool is quite easy, as one needs but look for the "_gitshot" tag. The report also indicates that in 93% of cases, images were found in repositories under direct control of the developer's username, rather than being tied to the company's GitHub account.
In one particular case, an agent skill (essentially long-winded prompts instructing a bot how to do something) was also an indirect source of leaked information. The agents started incorporating the image hosting workaround as a skill, and after a while, many of them were using it for every development ticket, leading to leaks of information about features months away from public release.
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Glow provided sample reasoning output from an agent, explaining precisely why it uploaded the screenshots publicly:
"internal_sweeper is private, and GitHub cannot render images from a private repo in a PR description — its image proxy fetches anonymously, so anything committed here (branch, release asset, whatever) shows up broken for reviewers. The only way to satisfy both "reviewers see the images" and "nothing but index.html in the repo" was to host the PNGs elsewhere, so I created a new public repo, sweeper-demo/pr-assets, holding the two screenshots pinned to a commit SHA."
As mitigation measures, Glow first recommends checking that employees don't use repositories under their own accounts, auditing the accounts and code managed by employees no longer associated with the company, and curtailing the use of "shadow AI" — when employees unadvisedly sign up for their own AI tools without talking to IT first and neatly punch planet-sized holes in the firm's security and data privacy.
Additionally, Glow advises strong vetting of software and code libraries used for development, as well as careful reading of the instructions and rules of any agentic skills.
Follow Tom's Hardware on Google News , or add us as a preferred source , to get our latest news, analysis, & reviews in your feeds.
Bruno Ferreira
Social Links Navigation
Contributor
Bruno Ferreira is a contributing writer for Tom's Hardware. He has decades of experience with PC hardware and assorted sundries, alongside a career as a developer. He's obsessed with detail and has a tendency to ramble on the topics he loves. When not doing that, he's usually playing games, or at live music shows and festivals.
Cybersecurity
New hack exploits AI hallucinations to trick agents into running malicious code
Artificial Intelligence
Researchers easily trick Fortune-500 companies' AI agents into running arbitrary code — supply-chain attack via llms.txt guidance file illustrates how data has become code
Artificial Intelligence
OpenAI's rogue AI agents accessed more websites to communicate than originally believed
Artificial Intelligence
OpenAI and Anthropic are reportedly investigating tens of thousands of AI security incidents; OpenAI pauses testing after AI 'kill switch' fails to stop a rogue agent
Artificial Intelligence
OpenAI agent goes rogue and hacks popular AI community
Artificial Intelligence
Anthropic's Claude hacked three real-life companies during security capabilities test
Latest in Cybersecurity
Cybersecurity
Pentagon gets pwned as breach exposes sensitive data on nearly three million military and civilian personnel
Cybersecurity
Teenager hacks open Microsoft database with 17 trillion total rows and 25,000 user accounts
Cybersecurity
Flock seeks to have security researchers' map of Flock cameras taken down
Cybersecurity
Novel attack slashes computing power needed to crack textbook RSA cryptography
Cybersecurity
Blockchain-assisted cyberattacks surge fivefold, driven by Iranian and North Korean state actors, Russia-linked groups
Cybersecurity
Asus confirms eShop data breach exposed customer order records and contact details
Latest in News
DRAM
Micron projects tightening RAM shortages through 2028 as it generates record profit
GPUs
Gears of War: E-Day PC graphics performance tested
Retro Gaming
PS3 emulator devs warn of fake Blu-ray drives on Newegg and AliExpress
Storage
Sony released the first CD audio player on this day in 1982
Laptops
HP boards the MacBook Neo competitor train — OmniBook 5 comes in four colors with Wildcat Lake and starts at $699.99
Artificial Intelligence
OpenAI says actors linked to China-based Moonshot AI spearheaded a campaign to extract its mo
[truncated]
Add as a preferred source on Google
Terms and conditions
©
Future US, Inc. Full 7th Floor, 130 West 42nd Street,
New York,
NY 10036.
