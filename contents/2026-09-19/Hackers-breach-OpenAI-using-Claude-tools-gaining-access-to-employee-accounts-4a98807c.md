---
source: "https://www.tomshardware.com/tech-industry/cyber-security/hackers-breach-openai-using-claude-tools-gaining-access-to-employee-accounts-and-the-companys-internal-codebase-initiating-a-harmless-pull-request-as-proof-of-the-hack"
hn_url: "https://news.ycombinator.com/item?id=49766056"
title: "Hackers breach OpenAI using Claude tools, gaining access to employee accounts"
article_title: "Hackers breach OpenAI using Claude tools, gaining access to employee accounts and the company's internal codebase — attackers initiated a 'harmless' pull request as proof of the hack | Tom's Hardware"
image: "https://cdn.mos.cms.futurecdn.net/YUDxAZxxyWFMPWzwJRmWvH-2560-80.jpg"
author: "thunderbong"
captured_at: "2026-09-19T12:55:11Z"
capture_tool: "hn-digest"
hn_id: 49766056
score: 1
comments: 0
posted_at: "2026-09-19T12:35:34Z"
tags:
  - hacker-news
---

# Hackers breach OpenAI using Claude tools, gaining access to employee accounts

- HN: [49766056](https://news.ycombinator.com/item?id=49766056)
- Source: [www.tomshardware.com](https://www.tomshardware.com/tech-industry/cyber-security/hackers-breach-openai-using-claude-tools-gaining-access-to-employee-accounts-and-the-companys-internal-codebase-initiating-a-harmless-pull-request-as-proof-of-the-hack)
- Score: 1
- Comments: 0
- Posted: 2026-09-19T12:35:34Z

## Translation

Title: Hackers breach OpenAI using Claude tools, gaining access to employee accounts
Article title: Hackers breach OpenAI using Claude tools, gaining access to employee accounts and the company's internal codebase — attackers initiated a 'harmless' pull request as proof of the hack | Tom's Hardware
Description: Researchers receive a $6,500 bounty after reporting the vulnerabilities

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
Hackers breach OpenAI using Claude tools, gaining access to employee accounts and the company's internal codebase — attackers initiated a 'harmless' pull request as proof of the hack
Researchers receive a $6,500 bounty after reporting the vulnerabilities
When you purchase through links on our site, we may earn an affiliate commission. Here’s how it works .
(Image credit: Getty Images)
Copy link
5
Join the conversation
Follow us
Add us as a preferred source on Google
Newsletter
Subscribe to our newsletter
A team of white-hat hackers from cybersecurity startup Hackron AI has successfully hacked OpenAI using Claude tools. In an X post on September 18, the team claimed they breached OpenAI's internal codebase on July 25 and gained access to the ChatGPT and Codex accounts of some OpenAI employees. They established proof of the hack via a pull request to OpenAI's private repository before reporting the vulnerabilities to OpenAI. The company reportedly fixed the issue within 14 hours of the report and paid the researchers a $6,500 bounty.
On July 25, our team hacked OpenAI. It took us less than 72 hours.Two vulnerabilities chained together gave us access to ChatGPT and Codex accounts belonging to OpenAI employees. We demonstrated the impact with a harmless PR in OpenAI’s internal monorepo.The full chain:… September 18, 2026
Operating as hackers under OpenAI’s bug bounty program, Hacktron researchers uncovered critical vulnerabilities that granted them access to internal employee tools and the ability to compromise private software repositories. The researchers exploited a single sign-on (SSO) misconfiguration and a Remote Code Execution (RCE) flaw in Discourse, a third-party platform that powers OpenAI’s community discussion forum. The chain of attack was as follows: HEIF upload → libheif heap overflow → RCE → OpenAI SSO flaw → ChatGPT/Codex takeover → connected GitHub → internal PR.
First, the researchers uploaded a malicious HEIF (High Efficiency Image File) image to the forum as a profile picture. When Discourse’s server-side software tried to process the image using an outdated libheif package, it triggered a heap overflow memory vulnerability, causing the library to crash and mismanage internal system memory. The researchers carefully orchestrated the memory crash to achieve remote code execution. After gaining access to the forum's local server environment, the researchers intercepted the server’s environmental configurations and session handling, discovering an SSO flaw in which the forum's authentication system did not adequately validate or isolate user sessions from other OpenAI services.
Armed with session tokens hijacked from the local forum server database, the hackers exploited the SSO flaw to impersonate a real OpenAI employee, allowing them to bypass traditional login screens and infiltrate a highly privileged internal account linked to OpenAI's development teams. As many tech companies unify authentication across corporate apps, the hijacked employee account was directly linked to OpenAI’s corporate enterprise systems, including GitHub, Slack, and email accounts. The researchers were able to access OpenAI’s massive private codebase, where they initiated an internal Pull Request as definitive proof of the exploit.
Similar to an incident last month in which China-linked hackers used AI to carry out the first-ever end-to-end autonomous cyberattack on Taiwan's government , the Hacktron hack also used artificial intelligence. The researchers constructed the exploit pipeline using Anthropic's Claude Opus 5 model, after attempts with Opus 4.8 failed. After they found the unpatched libheif library on OpenAI's forum, they fed the raw server data into the model, asking it to write an exploit for the bug.
OpenAI's GPT-5.6 Sol and unreleased AI models break out of testing environment in 'unprecedented cybersecurity incident'
Anthropic's Claude hacked three real-life companies during security capabilities test
OpenAI took ten days to tell Hugging Face its models were behind the July 11 weekend hack, report claims
The model analyzed the memory structure and successfully calculated how to trigger the heap buffer overflow. It generated the precise, weaponized code required to create the malicious HEIF image. The human hackers uploaded it to the forum — triggering the Remote Code Execution — then manually executed the rest of the “attack.” An important clarification is that they used an authorized, cybersecurity-configured version of Claude, which relaxes certain cyber restrictions for authorized researchers.
After gaining access, the researchers say they immediately halted testing and reported the vulnerabilities to OpenAI and Discourse — both of which have fixed their sides of the issue — without studying or downloading OpenAI's source code. From the initial finding to full resolution took 72 hours, after which OpenAI rewarded the researchers with a $6,500 bounty. The incident further highlights ongoing concerns over the risk of AI-powered cyberattacks. Recently, rogue OpenAI agents autonomously breached HuggingFace . US frontier AI companies are now warning against sophisticated distillation attacks .
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Follow Tom's Hardware on Google News , or add us as a preferred source , to get our latest news, analysis, & reviews in your feeds.
Etiido Uko
Social Links Navigation
News Contributor
Etiido Uko is a news contributor for Tom's Hardware covering the latest updates in big tech and the PC industry. He is a mechanical engineer and senior technical writer with over nine years of experience in documentation and reporting. He is deeply passionate about all things engineering and technology, and is an expert in gadgets, manufacturing, robotics, automotive, and aerospace.
Artificial Intelligence
OpenAI's GPT-5.6 Sol and unreleased AI models break out of testing environment in 'unprecedented cybersecurity incident'
Artificial Intelligence
Anthropic's Claude hacked three real-life companies during security capabilities test
Artificial Intelligence
OpenAI took ten days to tell Hugging Face its models were behind the July 11 weekend hack, report claims
Artificial Intelligence
OpenAI agent goes rogue and hacks popular AI community
Artificial Intelligence
OpenAI admits to 'wiki incident' after its agents were discovered using a programming hub to communicate
Artificial Intelligence
OpenAI's HuggingFace breach heralds an unprecedented age of AI cyber warfare
Latest in Cybersecurity
Cybersecurity
Hackers find encryption key stored on Flock camera
Cybersecurity
Researcher reverse-engineers infamous Stuxnet malware source code, publishes it on Github for all
Cybersecurity
Russian hacker faces up to 20 years in prison, following extradition and indictment over US phishing campaign that allegedly infected 80,000 PCs
Cybersecurity
FBI investigating 153 million US and Canadian driver’s licenses leaked on Russian cybercrime forum, including that of US SecDef Pete Hegseth
Cybersecurity
BlindLock hides your password manager and secure vault in a PNG image
Cybersecurity
Security researchers find surveillance implants in Chinese-made routers sold worldwide
Latest in News
SSDs
China's premier memory maker CXMT eyes producing flash for SSDs, report claims
Data Centers
House passes act to make AI data centers pay for grid upgrades to minimize impact on residents
Artificial Intelligence
AI developer vibe codes DLSS 5 onto Intel Arc 140T
PC Gaming
Control Resonant PC performance tested
Artificial Intelligence
Microsoft director called AI scraping ‘the largest theft of labor in human history'
GPUs
Modder gets Nvidia's DLSS 5 working in a web browser using WebGPU
5 Comments
Comment from the forums
EzzyB
Hackers breach OpenAI using Claude toolsBegun, the AI wa
[truncated]
Trake_17
The arms race continues
Reply
DS426
Only a $6,500 reward? That's low for this kind of thing, especially when a lot of the big tech companies will pay $20K for a single critical-severity software vulnerability.
Reply
Corgano
And they didn't change the default language to "pirate"? What a wasted opportunity! For shame.
Reply
-Fran-
"Live by the sword, die by the sword".
Regards.
Reply
Tom's Hardware is part of Future US Inc, an international media group and leading digital publisher. Visit our corporate site .
Add as a preferred source on Google
Terms and conditions
©
Future US, Inc. Full 7th Floor, 130 West 42nd Street,
New York,
NY 10036.

## Original Extract

Researchers receive a $6,500 bounty after reporting the vulnerabilities

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
Hackers breach OpenAI using Claude tools, gaining access to employee accounts and the company's internal codebase — attackers initiated a 'harmless' pull request as proof of the hack
Researchers receive a $6,500 bounty after reporting the vulnerabilities
When you purchase through links on our site, we may earn an affiliate commission. Here’s how it works .
(Image credit: Getty Images)
Copy link
5
Join the conversation
Follow us
Add us as a preferred source on Google
Newsletter
Subscribe to our newsletter
A team of white-hat hackers from cybersecurity startup Hackron AI has successfully hacked OpenAI using Claude tools. In an X post on September 18, the team claimed they breached OpenAI's internal codebase on July 25 and gained access to the ChatGPT and Codex accounts of some OpenAI employees. They established proof of the hack via a pull request to OpenAI's private repository before reporting the vulnerabilities to OpenAI. The company reportedly fixed the issue within 14 hours of the report and paid the researchers a $6,500 bounty.
On July 25, our team hacked OpenAI. It took us less than 72 hours.Two vulnerabilities chained together gave us access to ChatGPT and Codex accounts belonging to OpenAI employees. We demonstrated the impact with a harmless PR in OpenAI’s internal monorepo.The full chain:… September 18, 2026
Operating as hackers under OpenAI’s bug bounty program, Hacktron researchers uncovered critical vulnerabilities that granted them access to internal employee tools and the ability to compromise private software repositories. The researchers exploited a single sign-on (SSO) misconfiguration and a Remote Code Execution (RCE) flaw in Discourse, a third-party platform that powers OpenAI’s community discussion forum. The chain of attack was as follows: HEIF upload → libheif heap overflow → RCE → OpenAI SSO flaw → ChatGPT/Codex takeover → connected GitHub → internal PR.
First, the researchers uploaded a malicious HEIF (High Efficiency Image File) image to the forum as a profile picture. When Discourse’s server-side software tried to process the image using an outdated libheif package, it triggered a heap overflow memory vulnerability, causing the library to crash and mismanage internal system memory. The researchers carefully orchestrated the memory crash to achieve remote code execution. After gaining access to the forum's local server environment, the researchers intercepted the server’s environmental configurations and session handling, discovering an SSO flaw in which the forum's authentication system did not adequately validate or isolate user sessions from other OpenAI services.
Armed with session tokens hijacked from the local forum server database, the hackers exploited the SSO flaw to impersonate a real OpenAI employee, allowing them to bypass traditional login screens and infiltrate a highly privileged internal account linked to OpenAI's development teams. As many tech companies unify authentication across corporate apps, the hijacked employee account was directly linked to OpenAI’s corporate enterprise systems, including GitHub, Slack, and email accounts. The researchers were able to access OpenAI’s massive private codebase, where they initiated an internal Pull Request as definitive proof of the exploit.
Similar to an incident last month in which China-linked hackers used AI to carry out the first-ever end-to-end autonomous cyberattack on Taiwan's government , the Hacktron hack also used artificial intelligence. The researchers constructed the exploit pipeline using Anthropic's Claude Opus 5 model, after attempts with Opus 4.8 failed. After they found the unpatched libheif library on OpenAI's forum, they fed the raw server data into the model, asking it to write an exploit for the bug.
OpenAI's GPT-5.6 Sol and unreleased AI models break out of testing environment in 'unprecedented cybersecurity incident'
Anthropic's Claude hacked three real-life companies during security capabilities test
OpenAI took ten days to tell Hugging Face its models were behind the July 11 weekend hack, report claims
The model analyzed the memory structure and successfully calculated how to trigger the heap buffer overflow. It generated the precise, weaponized code required to create the malicious HEIF image. The human hackers uploaded it to the forum — triggering the Remote Code Execution — then manually executed the rest of the “attack.” An important clarification is that they used an authorized, cybersecurity-configured version of Claude, which relaxes certain cyber restrictions for authorized researchers.
After gaining access, the researchers say they immediately halted testing and reported the vulnerabilities to OpenAI and Discourse — both of which have fixed their sides of the issue — without studying or downloading OpenAI's source code. From the initial finding to full resolution took 72 hours, after which OpenAI rewarded the researchers with a $6,500 bounty. The incident further highlights ongoing concerns over the risk of AI-powered cyberattacks. Recently, rogue OpenAI agents autonomously breached HuggingFace . US frontier AI companies are now warning against sophisticated distillation attacks .
Get Tom's Hardware's best news and in-depth reviews, straight to your inbox.
Follow Tom's Hardware on Google News , or add us as a preferred source , to get our latest news, analysis, & reviews in your feeds.
Etiido Uko
Social Links Navigation
News Contributor
Etiido Uko is a news contributor for Tom's Hardware covering the latest updates in big tech and the PC industry. He is a mechanical engineer and senior technical writer with over nine years of experience in documentation and reporting. He is deeply passionate about all things engineering and technology, and is an expert in gadgets, manufacturing, robotics, automotive, and aerospace.
Artificial Intelligence
OpenAI's GPT-5.6 Sol and unreleased AI models break out of testing environment in 'unprecedented cybersecurity incident'
Artificial Intelligence
Anthropic's Claude hacked three real-life companies during security capabilities test
Artificial Intelligence
OpenAI took ten days to tell Hugging Face its models were behind the July 11 weekend hack, report claims
Artificial Intelligence
OpenAI agent goes rogue and hacks popular AI community
Artificial Intelligence
OpenAI admits to 'wiki incident' after its agents were discovered using a programming hub to communicate
Artificial Intelligence
OpenAI's HuggingFace breach heralds an unprecedented age of AI cyber warfare
Latest in Cybersecurity
Cybersecurity
Hackers find encryption key stored on Flock camera
Cybersecurity
Researcher reverse-engineers infamous Stuxnet malware source code, publishes it on Github for all
Cybersecurity
Russian hacker faces up to 20 years in prison, following extradition and indictment over US phishing campaign that allegedly infected 80,000 PCs
Cybersecurity
FBI investigating 153 million US and Canadian driver’s licenses leaked on Russian cybercrime forum, including that of US SecDef Pete Hegseth
Cybersecurity
BlindLock hides your password manager and secure vault in a PNG image
Cybersecurity
Security researchers find surveillance implants in Chinese-made routers sold worldwide
Latest in News
SSDs
China's premier memory maker CXMT eyes producing flash for SSDs, report claims
Data Centers
House passes act to make AI data centers pay for grid upgrades to minimize impact on residents
Artificial Intelligence
AI developer vibe codes DLSS 5 onto Intel Arc 140T
PC Gaming
Control Resonant PC performance tested
Artificial Intelligence
Microsoft director called AI scraping ‘the largest theft of labor in human history'
GPUs
Modder gets Nvidia's DLSS 5 working in a web browser using WebGPU
5 Comments
Comment from the forums
EzzyB
Hackers breach OpenAI using Claude toolsBegun, the AI wa
[truncated]
Trake_17
The arms race continues
Reply
DS426
Only a $6,500 reward? That's low for this kind of thing, especially when a lot of the big tech companies will pay $20K for a single critical-severity software vulnerability.
Reply
Corgano
And they didn't change the default language to "pirate"? What a wasted opportunity! For shame.
Reply
-Fran-
"Live by the sword, die by the sword".
Regards.
Reply
Tom's Hardware is part of Future US Inc, an international media group and leading digital publisher. Visit our corporate site .
Add as a preferred source on Google
Terms and conditions
©
Future US, Inc. Full 7th Floor, 130 West 42nd Street,
New York,
NY 10036.
