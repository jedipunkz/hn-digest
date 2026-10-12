---
source: "https://huntback.io/blog/ddos-amplification-lab-ai-copilot"
hn_url: "https://news.ycombinator.com/item?id=50053689"
title: "A Russian operator's DDoS-amplification lab, with an AI co-pilot"
article_title: "A Russian operator's DDoS-amplification lab, with an AI co-pilot | huntback blog"
image: "https://huntback.io/blog/img/og/ddos-amplification-lab-ai-copilot.png"
author: "noktec"
captured_at: "2026-10-12T10:19:54Z"
capture_tool: "hn-digest"
hn_id: 50053689
score: 1
comments: 0
posted_at: "2026-10-12T09:29:11Z"
tags:
  - hacker-news
---

# A Russian operator's DDoS-amplification lab, with an AI co-pilot

- HN: [50053689](https://news.ycombinator.com/item?id=50053689)
- Source: [huntback.io](https://huntback.io/blog/ddos-amplification-lab-ai-copilot)
- Score: 1
- Comments: 0
- Posted: 2026-10-12T09:29:11Z

## Translation

Title: A Russian operator's DDoS-amplification lab, with an AI co-pilot
Article title: A Russian operator's DDoS-amplification lab, with an AI co-pilot | huntback blog
Description: An exposed server mixed a cover application with a multi-protocol DDoS-amplification scanner, a self-healing worker daemon and the now-familiar signs of an AI-assisted operator, Russian-language throughout.

Article text:
huntback Deceive Hunt Code Surface Fraud CVE Intel Pricing Blog Login Start free Blog · Threat intel · 2026-10-09 · 8 min A Russian operator's DDoS-amplification lab, with an AI co-pilot
On 84.32.34.6 , a host at UAB Cherry Servers, an exposed directory showed a pattern we are seeing more of: a legitimate-looking application used as cover, an offensive tool beside it, a persistence worker to keep things running, and an AI assistant woven through the operator's workflow. Russian-language throughout.
A reflection-DDoS reconnaissance tool. ampscanner.py measures the bandwidth amplification factor of UDP reflectors across multiple protocols, the homework before a reflected flood.
Self-healing persistence. watchdog.sh relaunches an ssh_daemon.py worker inside a detached screen session whenever it dies.
An AI co-pilot, run with the safety rail off. A claude.md brief and a persona.md show the operator running the work with a model in the loop, and the agent config carries a dangerously-skip-permissions autonomy flag, the exact kind of leftover that outlives a model-backend swap.
Secrets in the open: a GitHub token, a private key and several key-value credentials sat in the exposed directory.
Figure 1. Observed window and the operator's build-to-operate sequence (from shell history).
The directory scored malicious on harvested secrets and an offensive scanner sitting next to the application code. We recovered and classified the contents automatically; sensors place the host on 2026-09-30. Alongside a Rust game backend (plausibly cover, or simply the box's day job) sat the interesting material.
How reflection amplification works, and what this tool prepares
Open UDP reflectors (DNS/NTP/SSDP...) BAF measured Victim (spoofed src) solid = recon today · dashed = the reflected flood it prepares
Figure 2. The scanner measures reflectors today; the dashed path is the reflected flood it sets up.
A multi-protocol amplification scanner
The centrepiece is ampscanner.py , a multi-protocol amplification scanner that measures the bandwidth amplification factor (BAF) of reflectors across DNS and other UDP services by sending a small query and measuring the response. A Russian comment states it is "read-only, no spoofing, just measuring BAF from a real IP", which is exactly the reconnaissance step before a reflection/amplification DDoS: find the reflectors that turn a small spoofed packet into a large one aimed at a victim. Measuring BAF is the homework; the attack is the exam.
Persistence is handled by watchdog.sh , which checks whether an ssh_daemon.py worker is running and, if not, relaunches it inside a detached screen session, logging as it goes. Pairing a long-running worker with a watchdog keeps an operation alive across crashes and restarts without a full C2 framework.
As on other operator boxes we have mapped recently, the tell-tale files are there: a claude.md project brief and a persona.md defining an assistant's role. The operator runs their offensive work with an AI model in the loop; it is becoming a default, not an exception. Tellingly, the agent config carries a dangerously-skip-permissions flag, the operator has switched off the rail that would otherwise prompt before destructive actions. That is exactly the kind of artifact that survives a model-backend swap: it describes how the operator runs the agent, not which provider they happen to rent this week, which is why our detection keys on it rather than on any one provider's API domains. We also found harvested secrets in the open directory, a GitHub token, a private key and several key-value credentials, exactly the material an exposed working directory should never contain.
Comments and operator notes are in Russian, placing a Russian-speaking operator behind the box. Language is a lead, not a named-group attribution. The combination, cover application, offensive scanner, self-healing worker and an AI co-pilot, is the recurring shape of the mid-tier operator in 2026.
Do not be a reflector: close or rate-limit internet-exposed UDP services (open DNS resolvers, NTP monlist, memcached, SSDP); these scans exist to find yours.
Rotate anything you have ever committed or exposed: tokens and keys in a working directory are stolen the moment the directory is indexed.
Deception catches the recon: an amplification scanner touching a decoy is unambiguous, hostile, early signal.
Browse live finds or start free and hunt your own.
New attacker-infrastructure writeups and live indicators, straight to your inbox. No spam, unsubscribe anytime.
Classify any IP, browse live attacker infrastructure, or deploy your own sensors. Free, no card.
huntback Deception-led security that hunts back.

## Original Extract

An exposed server mixed a cover application with a multi-protocol DDoS-amplification scanner, a self-healing worker daemon and the now-familiar signs of an AI-assisted operator, Russian-language throughout.

huntback Deceive Hunt Code Surface Fraud CVE Intel Pricing Blog Login Start free Blog · Threat intel · 2026-10-09 · 8 min A Russian operator's DDoS-amplification lab, with an AI co-pilot
On 84.32.34.6 , a host at UAB Cherry Servers, an exposed directory showed a pattern we are seeing more of: a legitimate-looking application used as cover, an offensive tool beside it, a persistence worker to keep things running, and an AI assistant woven through the operator's workflow. Russian-language throughout.
A reflection-DDoS reconnaissance tool. ampscanner.py measures the bandwidth amplification factor of UDP reflectors across multiple protocols, the homework before a reflected flood.
Self-healing persistence. watchdog.sh relaunches an ssh_daemon.py worker inside a detached screen session whenever it dies.
An AI co-pilot, run with the safety rail off. A claude.md brief and a persona.md show the operator running the work with a model in the loop, and the agent config carries a dangerously-skip-permissions autonomy flag, the exact kind of leftover that outlives a model-backend swap.
Secrets in the open: a GitHub token, a private key and several key-value credentials sat in the exposed directory.
Figure 1. Observed window and the operator's build-to-operate sequence (from shell history).
The directory scored malicious on harvested secrets and an offensive scanner sitting next to the application code. We recovered and classified the contents automatically; sensors place the host on 2026-09-30. Alongside a Rust game backend (plausibly cover, or simply the box's day job) sat the interesting material.
How reflection amplification works, and what this tool prepares
Open UDP reflectors (DNS/NTP/SSDP...) BAF measured Victim (spoofed src) solid = recon today · dashed = the reflected flood it prepares
Figure 2. The scanner measures reflectors today; the dashed path is the reflected flood it sets up.
A multi-protocol amplification scanner
The centrepiece is ampscanner.py , a multi-protocol amplification scanner that measures the bandwidth amplification factor (BAF) of reflectors across DNS and other UDP services by sending a small query and measuring the response. A Russian comment states it is "read-only, no spoofing, just measuring BAF from a real IP", which is exactly the reconnaissance step before a reflection/amplification DDoS: find the reflectors that turn a small spoofed packet into a large one aimed at a victim. Measuring BAF is the homework; the attack is the exam.
Persistence is handled by watchdog.sh , which checks whether an ssh_daemon.py worker is running and, if not, relaunches it inside a detached screen session, logging as it goes. Pairing a long-running worker with a watchdog keeps an operation alive across crashes and restarts without a full C2 framework.
As on other operator boxes we have mapped recently, the tell-tale files are there: a claude.md project brief and a persona.md defining an assistant's role. The operator runs their offensive work with an AI model in the loop; it is becoming a default, not an exception. Tellingly, the agent config carries a dangerously-skip-permissions flag, the operator has switched off the rail that would otherwise prompt before destructive actions. That is exactly the kind of artifact that survives a model-backend swap: it describes how the operator runs the agent, not which provider they happen to rent this week, which is why our detection keys on it rather than on any one provider's API domains. We also found harvested secrets in the open directory, a GitHub token, a private key and several key-value credentials, exactly the material an exposed working directory should never contain.
Comments and operator notes are in Russian, placing a Russian-speaking operator behind the box. Language is a lead, not a named-group attribution. The combination, cover application, offensive scanner, self-healing worker and an AI co-pilot, is the recurring shape of the mid-tier operator in 2026.
Do not be a reflector: close or rate-limit internet-exposed UDP services (open DNS resolvers, NTP monlist, memcached, SSDP); these scans exist to find yours.
Rotate anything you have ever committed or exposed: tokens and keys in a working directory are stolen the moment the directory is indexed.
Deception catches the recon: an amplification scanner touching a decoy is unambiguous, hostile, early signal.
Browse live finds or start free and hunt your own.
New attacker-infrastructure writeups and live indicators, straight to your inbox. No spam, unsubscribe anytime.
Classify any IP, browse live attacker infrastructure, or deploy your own sensors. Free, no card.
huntback Deception-led security that hunts back.
