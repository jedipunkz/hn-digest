---
source: "https://www.dragos.com/blog/ai-supply-chain-security-canary-tokens"
hn_url: "https://news.ycombinator.com/item?id=50582734"
title: "Prompt Canaries Detect Prompt Leaks at AI Providers"
article_title: "AI Supply Chain Security: What Canary Tokens Found | Dragos - Dragos"
image: "https://dragos.brightspotcdn.com/dims4/default/41c5c4a/2147483647/strip/true/crop/1307x735+309+0/resize/1440x810!/format/webp/quality/90/?url=https%3A%2F%2Fdragos-production-dragos.s3.us-west-2.amazonaws.com%2Fbrightspot%2Fd4%2Fb6%2Fb4dbcbe348a8b3506fb99c5419e7%2Fxot-hero.webp"
author: "paulis"
captured_at: "2026-11-20T22:30:44Z"
capture_tool: "hn-digest"
hn_id: 50582734
score: 1
comments: 0
posted_at: "2026-11-20T21:49:42Z"
tags:
  - hacker-news
---

# Prompt Canaries Detect Prompt Leaks at AI Providers

- HN: [50582734](https://news.ycombinator.com/item?id=50582734)
- Source: [www.dragos.com](https://www.dragos.com/blog/ai-supply-chain-security-canary-tokens)
- Score: 1
- Comments: 0
- Posted: 2026-11-20T21:49:42Z

## Translation

Title: Prompt Canaries Detect Prompt Leaks at AI Providers
Article title: AI Supply Chain Security: What Canary Tokens Found | Dragos - Dragos
Description: We seeded two million AI prompts with canary tokens and found three provider breaches, prompt theft, and ignored Zero Data Retention settings.

Article text:
AI Supply Chain Security: What Canary Tokens Found | Dragos - Dragos
PLATFORM
The Dragos Platform
Explore the xOT Environment
Software Supply Chain Security
RISK & EXPOSURE
Vulnerability Management
DETECTION & RESPONSE
Threat Detection
AI ACCELERATION
Dragos EmberAI
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
INTELLIGENCE
xOT Cyber Threat Intelligence
The living knowledge engine behind the Dragos Platform, built from over a decade of OT-specific telemetry, adversary research, and frontline expertise.
THREAT INTELLIGENCE
Dragos WorldView
THREAT RESEARCH
Threat Reports
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
SERVICES
Dragos xOT Services
Assess, Plan, & Respond
ASSESS & VALIDATE
OT Cyber Assessment
RESPONSE & RECOVER
Incident Response
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
INSIGHTS
Understanding xOT
xOT defines the full operational technology environment - every system that influences a physical process.
FUNDAMENTALS
OT Cybersecurity Basics
THREATS
Monitoring Threat Groups
BY CHALLENGE
Industrial Risk Management
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
SOLUTIONS
The Leader in xOT Cybersecurity
Safeguarding civilization from those trying to disrupt critical infrastructure we depend on every day.
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
RESOURCES
xOT Cybersecurity Resource Library
Over a decade of OT incident and research, in one place.
COMMUNITY & TRAINING
Community Resources
Dragos Industrial Security Conference
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
COMPANY
xOT Security
The New Standard for OT Defense
Dragos Industrial Security Conference
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
Menu
PLATFORM
The Dragos Platform
Explore the xOT Environment
Software Supply Chain Security
RISK & EXPOSURE
Vulnerability Management
DETECTION & RESPONSE
Threat Detection
AI ACCELERATION
Dragos EmberAI
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
INTELLIGENCE
xOT Cyber Threat Intelligence
The living knowledge engine behind the Dragos Platform, built from over a decade of OT-specific telemetry, adversary research, and frontline expertise.
THREAT INTELLIGENCE
Dragos WorldView
THREAT RESEARCH
Threat Reports
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
SERVICES
Dragos xOT Services
Assess, Plan, & Respond
ASSESS & VALIDATE
OT Cyber Assessment
RESPONSE & RECOVER
Incident Response
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
INSIGHTS
Understanding xOT
xOT defines the full operational technology environment - every system that influences a physical process.
FUNDAMENTALS
OT Cybersecurity Basics
THREATS
Monitoring Threat Groups
BY CHALLENGE
Industrial Risk Management
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
SOLUTIONS
The Leader in xOT Cybersecurity
Safeguarding civilization from those trying to disrupt critical infrastructure we depend on every day.
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
RESOURCES
xOT Cybersecurity Resource Library
Over a decade of OT incident and research, in one place.
COMMUNITY & TRAINING
Community Resources
Dragos Industrial Security Conference
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
COMPANY
xOT Security
The New Standard for OT Defense
Dragos Industrial Security Conference
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
Submit Search
Show Search
Dragos has acquired NetRise and runZero, building the industry’s first end-to-end xOT cybersecurity platform. Read the blog →
Close
Blog
Three breaches, endless leaks: Your AI supply chain needs canaries
Why does AI inference data matter?
How do AI routers expose your prompts?
Marketplace resellers stole customer prompts
Privacy settings and retention policies were ignored
How do canary tokens detect prompt leaks?
How can you protect your AI supply chain?
HD Moore, Chief Research Officer, will be regularly sharing work from the research team he leads at Dragos. Topics will include xOT, IT, LLMs, PQC, and more. Their findings will sharpen defenders’ understanding of exposures and vulnerabilities emerging across the threat landscape.
We sent nearly two million AI prompts containing unique, trackable secrets to AI providers, then watched where those secrets surfaced. What we found revealed security breaches, prompt theft, and widespread failures to protect customer data.
Over a five-month period from July through November 2026, we submitted approximately two million prompts containing unique URLs and canary token credentials to widely used AI services and monitored what happened after inference. Our investigation uncovered three significant findings:
Evidence of provider compromises. We observed activity indicating that data from at least three AI providers was exposed. Our seeded test prompts were used in attempts to access our controlled test systems and canary AWS accounts.
Unauthorized access to customer prompts by inference marketplace providers. Multiple resellers operating through an AI inference marketplace were able to view and retain customer prompts using standard logging capabilities, without exploiting a vulnerability or bypassing platform controls.
Failures in privacy and data-handling practices. We observed multiple cases in which providers appeared not to honor customer privacy settings or published policies. Examples included prompts submitted to U.S.-based providers being transmitted to Chinese AI organizations, unique URLs being accessed months after a single inference request, and one seeded link ultimately appearing in an internal company chat in China.
We wanted to understand what happens to inference data after it leaves the customer’s environment.
To answer that question, we created millions of unique canaries, including URLs and credentials that would never be used by legitimate systems. Any attempt to access those assets would indicate that someone, somewhere in the AI ecosystem had seen the original prompt.
The results exposed a fundamental disconnect between how AI services are marketed and how inference data is handled in practice.
Why does AI inference data matter?
Most folks picture AI inference as a question and a response. Text goes in, GPUs juggle some bits, and text comes back out. What complicates this is how the data is actually treated.
AI providers commonly store prompts and responses for operational, training, safety, abuse-monitoring, or business purposes. Depending on the provider, that data may be retained for days, months, or indefinitely. It may also traverse multiple organizations, including model developers, cloud operators, inference platforms, and third-party resellers.
The prompts and replies – the inference data itself – are often more valuable than the GPU resources used to process them. They frequently contain confidential source code, internal communications, strategic plans, customer information, credentials, and other sensitive business data.
How do AI routers expose your prompts?
AI routers work by offering a single API endpoint that forwards your request to a pool of available models, sometimes at different hosting providers, with unified billing. Routers are great at cost reduction and providing flexibility without managing multiple accounts and credentials. The path to the inference server often passes through more than one router – it’s common for requests from Vercel to then pass through OpenRouter, before reaching another provider, and from there the specific server hosting the target model. There is risk in this process – every hop can see the entire prompt and response.
Some routers are open marketplaces where the seller with the lowest price gets the request. It’s common for sellers to offer AI inference at deep discounts, and they can do this by pooling requests through unmetered consumer-plan subscriptions or by finding other low-cost paths to the specific model (startup credits, etc.). In the shadier marketplaces, it’s also common to see inference that is billed to stolen AWS credentials or consumer API keys. If a seller wants to siphon customer request data, they can list popular models at the lowest price, route the traffic, pay the difference, and then collect the data from the platform logs.
These providers offer privacy controls – customers can set privacy flags at the account level, for a specific API key, or in each request. The strictest setting is known as Zero Data Retention (ZDR), which specifies that the prompt and response are not stored. The routers should only forward these requests to providers that offer the same controls. The reality is quite different – providers often ignore these settings, either intentionally or due to bugs in their processing flow, and prompts not covered by ZDR are fair game for resale into a massive ecosystem of buyers and sellers.
The most serious finding was also the most direct. We found attackers actively stealing data from three different hosting providers. Although our prompts were bait, they were processed along with real customer requests, and those were almost certainly copied too. The earliest incident dates back to the middle of July, and none of the three providers knew about their breach until we notified them.
The largest breach involved a couple hundred of our AWS keys. Each key was sent in exactly one request to one provider, all within a few short windows in late October, and every request we sent during those windows was stolen. Some of the stolen keys came from requests that the provider rejected before a model ever ran – this tells us that the attacker was copying requests at the provider’s front door, and not somewhere deeper in the inference stack. In the other two breaches, some of the stolen requests were marked as ZDR. A privacy flag doesn’t help when the attacker is already inside.
Once the attacker had our credentials, they followed a predictable path. First, they tested each key to see if it was live, in some cases within a few hours of our request. Next, they went looking for data, listing S3 buckets and searching Secrets Manager and Parameter Store. Finally, they tried enabling hosted models in AWS Bedrock, creating Bedrock API keys, and adding their own IAM users to keep access. In total, our bait keys were used for tens of thousands of AWS API calls.
Figure 1. One stolen AWS key. We seeded it on October 21, it was first used the next day, and it was still in use three weeks later. Names are pseudonyms and addresses are partly masked.
One adversary appears to have compromised two of these providers at the same time. Secrets that we sent to those two providers were used together, almost immediately, from a single rented server. The attacker’s tooling changed over the three weeks that we watched it – it started with a script pretending to be Chrom

[truncated]

## Original Extract

We seeded two million AI prompts with canary tokens and found three provider breaches, prompt theft, and ignored Zero Data Retention settings.

AI Supply Chain Security: What Canary Tokens Found | Dragos - Dragos
PLATFORM
The Dragos Platform
Explore the xOT Environment
Software Supply Chain Security
RISK & EXPOSURE
Vulnerability Management
DETECTION & RESPONSE
Threat Detection
AI ACCELERATION
Dragos EmberAI
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
INTELLIGENCE
xOT Cyber Threat Intelligence
The living knowledge engine behind the Dragos Platform, built from over a decade of OT-specific telemetry, adversary research, and frontline expertise.
THREAT INTELLIGENCE
Dragos WorldView
THREAT RESEARCH
Threat Reports
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
SERVICES
Dragos xOT Services
Assess, Plan, & Respond
ASSESS & VALIDATE
OT Cyber Assessment
RESPONSE & RECOVER
Incident Response
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
INSIGHTS
Understanding xOT
xOT defines the full operational technology environment - every system that influences a physical process.
FUNDAMENTALS
OT Cybersecurity Basics
THREATS
Monitoring Threat Groups
BY CHALLENGE
Industrial Risk Management
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
SOLUTIONS
The Leader in xOT Cybersecurity
Safeguarding civilization from those trying to disrupt critical infrastructure we depend on every day.
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
RESOURCES
xOT Cybersecurity Resource Library
Over a decade of OT incident and research, in one place.
COMMUNITY & TRAINING
Community Resources
Dragos Industrial Security Conference
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
COMPANY
xOT Security
The New Standard for OT Defense
Dragos Industrial Security Conference
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
Menu
PLATFORM
The Dragos Platform
Explore the xOT Environment
Software Supply Chain Security
RISK & EXPOSURE
Vulnerability Management
DETECTION & RESPONSE
Threat Detection
AI ACCELERATION
Dragos EmberAI
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
INTELLIGENCE
xOT Cyber Threat Intelligence
The living knowledge engine behind the Dragos Platform, built from over a decade of OT-specific telemetry, adversary research, and frontline expertise.
THREAT INTELLIGENCE
Dragos WorldView
THREAT RESEARCH
Threat Reports
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
SERVICES
Dragos xOT Services
Assess, Plan, & Respond
ASSESS & VALIDATE
OT Cyber Assessment
RESPONSE & RECOVER
Incident Response
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
INSIGHTS
Understanding xOT
xOT defines the full operational technology environment - every system that influences a physical process.
FUNDAMENTALS
OT Cybersecurity Basics
THREATS
Monitoring Threat Groups
BY CHALLENGE
Industrial Risk Management
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
SOLUTIONS
The Leader in xOT Cybersecurity
Safeguarding civilization from those trying to disrupt critical infrastructure we depend on every day.
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
RESOURCES
xOT Cybersecurity Resource Library
Over a decade of OT incident and research, in one place.
COMMUNITY & TRAINING
Community Resources
Dragos Industrial Security Conference
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
COMPANY
xOT Security
The New Standard for OT Defense
Dragos Industrial Security Conference
report
A Leader in the Forrester Wave
See the full evaluation of the OT security market and where Dragos scored.
See Report
September 1, 2026 09:24 AM
Submit Search
Show Search
Dragos has acquired NetRise and runZero, building the industry’s first end-to-end xOT cybersecurity platform. Read the blog →
Close
Blog
Three breaches, endless leaks: Your AI supply chain needs canaries
Why does AI inference data matter?
How do AI routers expose your prompts?
Marketplace resellers stole customer prompts
Privacy settings and retention policies were ignored
How do canary tokens detect prompt leaks?
How can you protect your AI supply chain?
HD Moore, Chief Research Officer, will be regularly sharing work from the research team he leads at Dragos. Topics will include xOT, IT, LLMs, PQC, and more. Their findings will sharpen defenders’ understanding of exposures and vulnerabilities emerging across the threat landscape.
We sent nearly two million AI prompts containing unique, trackable secrets to AI providers, then watched where those secrets surfaced. What we found revealed security breaches, prompt theft, and widespread failures to protect customer data.
Over a five-month period from July through November 2026, we submitted approximately two million prompts containing unique URLs and canary token credentials to widely used AI services and monitored what happened after inference. Our investigation uncovered three significant findings:
Evidence of provider compromises. We observed activity indicating that data from at least three AI providers was exposed. Our seeded test prompts were used in attempts to access our controlled test systems and canary AWS accounts.
Unauthorized access to customer prompts by inference marketplace providers. Multiple resellers operating through an AI inference marketplace were able to view and retain customer prompts using standard logging capabilities, without exploiting a vulnerability or bypassing platform controls.
Failures in privacy and data-handling practices. We observed multiple cases in which providers appeared not to honor customer privacy settings or published policies. Examples included prompts submitted to U.S.-based providers being transmitted to Chinese AI organizations, unique URLs being accessed months after a single inference request, and one seeded link ultimately appearing in an internal company chat in China.
We wanted to understand what happens to inference data after it leaves the customer’s environment.
To answer that question, we created millions of unique canaries, including URLs and credentials that would never be used by legitimate systems. Any attempt to access those assets would indicate that someone, somewhere in the AI ecosystem had seen the original prompt.
The results exposed a fundamental disconnect between how AI services are marketed and how inference data is handled in practice.
Why does AI inference data matter?
Most folks picture AI inference as a question and a response. Text goes in, GPUs juggle some bits, and text comes back out. What complicates this is how the data is actually treated.
AI providers commonly store prompts and responses for operational, training, safety, abuse-monitoring, or business purposes. Depending on the provider, that data may be retained for days, months, or indefinitely. It may also traverse multiple organizations, including model developers, cloud operators, inference platforms, and third-party resellers.
The prompts and replies – the inference data itself – are often more valuable than the GPU resources used to process them. They frequently contain confidential source code, internal communications, strategic plans, customer information, credentials, and other sensitive business data.
How do AI routers expose your prompts?
AI routers work by offering a single API endpoint that forwards your request to a pool of available models, sometimes at different hosting providers, with unified billing. Routers are great at cost reduction and providing flexibility without managing multiple accounts and credentials. The path to the inference server often passes through more than one router – it’s common for requests from Vercel to then pass through OpenRouter, before reaching another provider, and from there the specific server hosting the target model. There is risk in this process – every hop can see the entire prompt and response.
Some routers are open marketplaces where the seller with the lowest price gets the request. It’s common for sellers to offer AI inference at deep discounts, and they can do this by pooling requests through unmetered consumer-plan subscriptions or by finding other low-cost paths to the specific model (startup credits, etc.). In the shadier marketplaces, it’s also common to see inference that is billed to stolen AWS credentials or consumer API keys. If a seller wants to siphon customer request data, they can list popular models at the lowest price, route the traffic, pay the difference, and then collect the data from the platform logs.
These providers offer privacy controls – customers can set privacy flags at the account level, for a specific API key, or in each request. The strictest setting is known as Zero Data Retention (ZDR), which specifies that the prompt and response are not stored. The routers should only forward these requests to providers that offer the same controls. The reality is quite different – providers often ignore these settings, either intentionally or due to bugs in their processing flow, and prompts not covered by ZDR are fair game for resale into a massive ecosystem of buyers and sellers.
The most serious finding was also the most direct. We found attackers actively stealing data from three different hosting providers. Although our prompts were bait, they were processed along with real customer requests, and those were almost certainly copied too. The earliest incident dates back to the middle of July, and none of the three providers knew about their breach until we notified them.
The largest breach involved a couple hundred of our AWS keys. Each key was sent in exactly one request to one provider, all within a few short windows in late October, and every request we sent during those windows was stolen. Some of the stolen keys came from requests that the provider rejected before a model ever ran – this tells us that the attacker was copying requests at the provider’s front door, and not somewhere deeper in the inference stack. In the other two breaches, some of the stolen requests were marked as ZDR. A privacy flag doesn’t help when the attacker is already inside.
Once the attacker had our credentials, they followed a predictable path. First, they tested each key to see if it was live, in some cases within a few hours of our request. Next, they went looking for data, listing S3 buckets and searching Secrets Manager and Parameter Store. Finally, they tried enabling hosted models in AWS Bedrock, creating Bedrock API keys, and adding their own IAM users to keep access. In total, our bait keys were used for tens of thousands of AWS API calls.
Figure 1. One stolen AWS key. We seeded it on October 21, it was first used the next day, and it was still in use three weeks later. Names are pseudonyms and addresses are partly masked.
One adversary appears to have compromised two of these providers at the same time. Secrets that we sent to those two providers were used together, almost immediately, from a single rented server. The attacker’s tooling changed over the three weeks that we watched it – it started with a script pretending to be Chrom

[truncated]
