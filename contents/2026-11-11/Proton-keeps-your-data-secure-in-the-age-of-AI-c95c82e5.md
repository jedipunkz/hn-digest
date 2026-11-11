---
source: "https://proton.me/blog/how-proton-keeps-your-data-secure"
hn_url: "https://news.ycombinator.com/item?id=50451004"
title: "Proton keeps your data secure in the age of AI"
article_title: "How Proton keeps your data secure in the age of AI | Proton"
image: "https://pmecdn.protonweb.com/image-transformation/?s=c&image=images/w_1024,h_512,c_scale/f_auto,q_auto/v1794309174/wp-pme/image-5_313212519cc/image-5_313212519cc.png?_i=AA"
author: "fold3"
captured_at: "2026-11-11T08:55:11Z"
capture_tool: "hn-digest"
hn_id: 50451004
score: 1
comments: 1
posted_at: "2026-11-11T08:34:00Z"
tags:
  - hacker-news
---

# Proton keeps your data secure in the age of AI

- HN: [50451004](https://news.ycombinator.com/item?id=50451004)
- Source: [proton.me](https://proton.me/blog/how-proton-keeps-your-data-secure)
- Score: 1
- Comments: 1
- Posted: 2026-11-11T08:34:00Z

## Translation

Title: Proton keeps your data secure in the age of AI
Article title: How Proton keeps your data secure in the age of AI | Proton
Description: We have added powerful new AI review layers on top of existing human ones, and much of our security posture hasn

Article text:
How Proton keeps your data secure in the age of AI | Proton Products Products
Discover Proton Discover Proton
Proton VPN (new window) Cloud storage Proton Drive Password manager Proton Pass AI assistant Lumo AI Video conferencing Proton Meet Private scheduling Proton Calendar Collaborative documents Proton Docs Collaborative spreadsheets Proton Sheets Two-factor authenticator Proton Authenticator Bitcoin wallet Proton Wallet Encrypted notes Standard Notes (new window) Hide-my-email aliases SimpleLogin (new window) Discover Proton Personal plans Proton Choose an internet where your privacy comes first Create a free account All-in-one business suite Proton Workspace Business email Proton Mail Business VPN Proton VPN Business AI assistant Lumo AI Business video conferencing Proton Meet Business calendar Proton Calendar Business password manager Proton Pass Business cloud storage Proton Drive Business documents Proton Docs Business spreadsheets Proton Sheets Industries Tech Healthcare Nonprofits News and media Business resources Trust center HIPAA GDPR Data Breach Observatory ISO 27001 Case studies Business blog Switch from Microsoft 365 Google Workspace Discover Proton for Business Business plans Proton Workspace Your whole team, one encrypted workspace Explore Workspace Proton Mail Proton VPN Proton Drive Proton Pass Lumo AI Proton Meet Proton Calendar Proton Authenticator Proton Wallet Scan to get the mobile app
Web app (new window) Desktop client add-on Bridge Scan to get the mobile app
CLI Web app Web app (new window) Scan to get the mobile app
Command line interface (new window) Extensions Chrome Safari Firefox Edge Brave Web app Web app (new window) Scan to get the mobile app
Web app (new window) Scan to get the mobile app
Web app (new window) Scan to get the mobile app
Web app (new window) Scan to get the mobile app
Web app (new window) Company About us Impact Team Careers We’re hiring Proton Foundation Proton blog Privacy Open source Zero-access encryption Big Tech alternatives Connect Partners and affiliates Social media Media contact Discover Proton Personal plans Proton We believe a better world starts with privacy and digital freedom Discover Proton Get started with Proton Beginner video guides (new window) Learn about encryption Business resources News Proton blog Product updates Reports and insights Investigative reports Proton VPN observatory (new window) Data breach observatory Tech fines tracker Europe’s tech sovereignty watch European alternatives Age verification explained Proton’s guide to AI Videos about privacy Mass surveillance explained (new window) How employers spy on you (new window) The data you leak in 24 hours (new window) How to disappear online (new window) See all videos (new window) Tools What is my IP address? (new window) Password generator Username generator PDF watermark Image watermark Flock camera map Support Help center Contact us EN Create a free account Sign in Open menu Blog homepage Proton updates
Subscribe by RSS Search for blog posts Create account How Proton keeps your data secure in the age of AI
AI has already transformed how many companies operate, whether it’s writing code, answering support tickets, or handling data. It has certainly changed how we work at Proton. So it is fair to ask what the arrival of AI means for the security of the products you depend on.
The short answer: We have added powerful new AI review layers on top of existing human ones, and much of our security posture hasn’t changed at all. It’s worth discussing all this in more detail, though, so let’s dig in.
The foundations remain the same
Proton uses end-to-end and zero-access encryption . We do not hold the keys to your emails, files, or passwords, so we cannot read them, and no AI tool can, either. Our documented security principles have not changed since June 2024, before AI was even a major factor in our work here, and they all still apply. They include basing decisions on risk, centralizing governance, and granting exceptions only in an explicit, controlled way. We have not loosened them for AI. We have applied them to it.
Proton is ISO 27001 certified . In July 2025, we also completed our first SOC 2 Type II audit , carried out by Schellman, an independent auditing firm. The auditors examined our controls for access management, incident response, system monitoring and risk assessment. These same controls apply no matter how we build software, with or without AI. Proton services undergo regular security audits both internally and externally, and we have another SOC 2 audit scheduled for 2027. And as always, our code remains open source and available for independent review and verification. Anybody, human or AI, can scan that code to verify its security.
The controls for AI-written code
At Proton, the same unified controls apply to all code, whether it was written by AI, a human, or a combination of both. AI-generated code doesn’t get to take any shortcuts or bypass any checks. These unified controls include:
A person is required to approve every code change. As has been the policy for years, every change must be reviewed, and one or more people are expected to ultimately approve those changes. AI reviewers have been added to the process, but they cannot deploy code to production on their own. The specific people required to sign off depends on which codebases have been modified: For example, core cryptography code changes always need to be reviewed by our expert cryptography team.
Security scanners are still required. Automated security scans must run on code changes.
People are always accountable. Engineers can hand tasks to AI, but they stay responsible for the result and must follow our policies and use their professional judgment. Code generated with the help of AI must meet or exceed human standards.
Designs are reviewed before major features are put in to code. Engineers present designs to their peers, who check the security, dependencies, testing, roll-out and roll-back plans.
The controls around AI and our infrastructure
AI agents get only the access they need. We apply the same least-privilege rules to AI agents as to employees. Anything an employee wasn’t allowed to do before AI is still not allowed.
Cyberdefense is strengthened. Our internal Security Information and Event Management (SIEM) system records actions within our security perimeter, whether they come from a person or an AI agent. Our human security analysts act on every suspicious alert.
Our infrastructure is protected. Only authorized devices can reach our internal systems. Authorization requires physical tokens that AI systems cannot physically manipulate on their own. We monitor all servers and employee laptops for indicators of compromise.
Should security companies use AI?
Because of AI, attacks that once might have required a nation-state can now be executed by almost anyone. A defender relying on manual processes and hand-crafted tools will be hopelessly outgunned. An attacker only needs to find one flaw, while a defender needs to close all of them. AI is the only tool that makes that defense possible. Cybersecurity is an arms race, and the teams working on security at Proton have been some of the earliest adopters of AI, because the safety of your data depends on winning that race.
AI reviews code changes, alongside humans. AI reviewers complement our human reviewers and flag issues earlier in development, catching defects pre-merge that manual reviews would have only caught much later, if at all. They add to human review; today, our policy is that every change is subject to human approval.
AI helps discover vulnerabilities first. We run continuous AI-powered security reviews against our own code so we find weaknesses before anyone else does. Alongside our own models, we use non-public cyber models from OpenAI and Anthropic for threat analysis and code scanning. These models receive source code and threat intelligence only, never customer content.
AI strengthens security monitoring. Our security operations center watches our perimeter and endpoints for compromise, with Proton Sentinel protecting customer accounts. AI tooling built by Proton lets a team of analysts patrol an attack surface that has grown far beyond what unaided humans could monitor, surfacing the signals a human team then investigates and acts on.
Proton’s security controls have been strengthened by the introduction of AI, as we now have advanced capabilities that were not available before. Things can be done faster and more thoroughly with AI, and given the threats we face, they must be, or we won’t be able to keep up.
No company can remove every risk, and we won’t claim to. What we can promise is a disciplined approach: We manage risks explicitly, keep improving our security, and keep a person accountable for every change that reaches production software.
We are scientists, engineers, and specialists from around the world drawn together by a shared vision of protecting freedom and privacy online. Proton was born out of a desire to build an internet that puts people before profits, and we're working to create a world where everyone is in control of their digital lives.
English System status (new window)
Request a feature (new window)
© 2026 Proton AG. All rights reserved.

## Original Extract

We have added powerful new AI review layers on top of existing human ones, and much of our security posture hasn

How Proton keeps your data secure in the age of AI | Proton Products Products
Discover Proton Discover Proton
Proton VPN (new window) Cloud storage Proton Drive Password manager Proton Pass AI assistant Lumo AI Video conferencing Proton Meet Private scheduling Proton Calendar Collaborative documents Proton Docs Collaborative spreadsheets Proton Sheets Two-factor authenticator Proton Authenticator Bitcoin wallet Proton Wallet Encrypted notes Standard Notes (new window) Hide-my-email aliases SimpleLogin (new window) Discover Proton Personal plans Proton Choose an internet where your privacy comes first Create a free account All-in-one business suite Proton Workspace Business email Proton Mail Business VPN Proton VPN Business AI assistant Lumo AI Business video conferencing Proton Meet Business calendar Proton Calendar Business password manager Proton Pass Business cloud storage Proton Drive Business documents Proton Docs Business spreadsheets Proton Sheets Industries Tech Healthcare Nonprofits News and media Business resources Trust center HIPAA GDPR Data Breach Observatory ISO 27001 Case studies Business blog Switch from Microsoft 365 Google Workspace Discover Proton for Business Business plans Proton Workspace Your whole team, one encrypted workspace Explore Workspace Proton Mail Proton VPN Proton Drive Proton Pass Lumo AI Proton Meet Proton Calendar Proton Authenticator Proton Wallet Scan to get the mobile app
Web app (new window) Desktop client add-on Bridge Scan to get the mobile app
CLI Web app Web app (new window) Scan to get the mobile app
Command line interface (new window) Extensions Chrome Safari Firefox Edge Brave Web app Web app (new window) Scan to get the mobile app
Web app (new window) Scan to get the mobile app
Web app (new window) Scan to get the mobile app
Web app (new window) Scan to get the mobile app
Web app (new window) Company About us Impact Team Careers We’re hiring Proton Foundation Proton blog Privacy Open source Zero-access encryption Big Tech alternatives Connect Partners and affiliates Social media Media contact Discover Proton Personal plans Proton We believe a better world starts with privacy and digital freedom Discover Proton Get started with Proton Beginner video guides (new window) Learn about encryption Business resources News Proton blog Product updates Reports and insights Investigative reports Proton VPN observatory (new window) Data breach observatory Tech fines tracker Europe’s tech sovereignty watch European alternatives Age verification explained Proton’s guide to AI Videos about privacy Mass surveillance explained (new window) How employers spy on you (new window) The data you leak in 24 hours (new window) How to disappear online (new window) See all videos (new window) Tools What is my IP address? (new window) Password generator Username generator PDF watermark Image watermark Flock camera map Support Help center Contact us EN Create a free account Sign in Open menu Blog homepage Proton updates
Subscribe by RSS Search for blog posts Create account How Proton keeps your data secure in the age of AI
AI has already transformed how many companies operate, whether it’s writing code, answering support tickets, or handling data. It has certainly changed how we work at Proton. So it is fair to ask what the arrival of AI means for the security of the products you depend on.
The short answer: We have added powerful new AI review layers on top of existing human ones, and much of our security posture hasn’t changed at all. It’s worth discussing all this in more detail, though, so let’s dig in.
The foundations remain the same
Proton uses end-to-end and zero-access encryption . We do not hold the keys to your emails, files, or passwords, so we cannot read them, and no AI tool can, either. Our documented security principles have not changed since June 2024, before AI was even a major factor in our work here, and they all still apply. They include basing decisions on risk, centralizing governance, and granting exceptions only in an explicit, controlled way. We have not loosened them for AI. We have applied them to it.
Proton is ISO 27001 certified . In July 2025, we also completed our first SOC 2 Type II audit , carried out by Schellman, an independent auditing firm. The auditors examined our controls for access management, incident response, system monitoring and risk assessment. These same controls apply no matter how we build software, with or without AI. Proton services undergo regular security audits both internally and externally, and we have another SOC 2 audit scheduled for 2027. And as always, our code remains open source and available for independent review and verification. Anybody, human or AI, can scan that code to verify its security.
The controls for AI-written code
At Proton, the same unified controls apply to all code, whether it was written by AI, a human, or a combination of both. AI-generated code doesn’t get to take any shortcuts or bypass any checks. These unified controls include:
A person is required to approve every code change. As has been the policy for years, every change must be reviewed, and one or more people are expected to ultimately approve those changes. AI reviewers have been added to the process, but they cannot deploy code to production on their own. The specific people required to sign off depends on which codebases have been modified: For example, core cryptography code changes always need to be reviewed by our expert cryptography team.
Security scanners are still required. Automated security scans must run on code changes.
People are always accountable. Engineers can hand tasks to AI, but they stay responsible for the result and must follow our policies and use their professional judgment. Code generated with the help of AI must meet or exceed human standards.
Designs are reviewed before major features are put in to code. Engineers present designs to their peers, who check the security, dependencies, testing, roll-out and roll-back plans.
The controls around AI and our infrastructure
AI agents get only the access they need. We apply the same least-privilege rules to AI agents as to employees. Anything an employee wasn’t allowed to do before AI is still not allowed.
Cyberdefense is strengthened. Our internal Security Information and Event Management (SIEM) system records actions within our security perimeter, whether they come from a person or an AI agent. Our human security analysts act on every suspicious alert.
Our infrastructure is protected. Only authorized devices can reach our internal systems. Authorization requires physical tokens that AI systems cannot physically manipulate on their own. We monitor all servers and employee laptops for indicators of compromise.
Should security companies use AI?
Because of AI, attacks that once might have required a nation-state can now be executed by almost anyone. A defender relying on manual processes and hand-crafted tools will be hopelessly outgunned. An attacker only needs to find one flaw, while a defender needs to close all of them. AI is the only tool that makes that defense possible. Cybersecurity is an arms race, and the teams working on security at Proton have been some of the earliest adopters of AI, because the safety of your data depends on winning that race.
AI reviews code changes, alongside humans. AI reviewers complement our human reviewers and flag issues earlier in development, catching defects pre-merge that manual reviews would have only caught much later, if at all. They add to human review; today, our policy is that every change is subject to human approval.
AI helps discover vulnerabilities first. We run continuous AI-powered security reviews against our own code so we find weaknesses before anyone else does. Alongside our own models, we use non-public cyber models from OpenAI and Anthropic for threat analysis and code scanning. These models receive source code and threat intelligence only, never customer content.
AI strengthens security monitoring. Our security operations center watches our perimeter and endpoints for compromise, with Proton Sentinel protecting customer accounts. AI tooling built by Proton lets a team of analysts patrol an attack surface that has grown far beyond what unaided humans could monitor, surfacing the signals a human team then investigates and acts on.
Proton’s security controls have been strengthened by the introduction of AI, as we now have advanced capabilities that were not available before. Things can be done faster and more thoroughly with AI, and given the threats we face, they must be, or we won’t be able to keep up.
No company can remove every risk, and we won’t claim to. What we can promise is a disciplined approach: We manage risks explicitly, keep improving our security, and keep a person accountable for every change that reaches production software.
We are scientists, engineers, and specialists from around the world drawn together by a shared vision of protecting freedom and privacy online. Proton was born out of a desire to build an internet that puts people before profits, and we're working to create a world where everyone is in control of their digital lives.
English System status (new window)
Request a feature (new window)
© 2026 Proton AG. All rights reserved.
