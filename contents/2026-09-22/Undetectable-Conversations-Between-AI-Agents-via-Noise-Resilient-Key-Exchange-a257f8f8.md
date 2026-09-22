---
source: "https://arxiv.org/abs/2604.04757"
hn_url: "https://news.ycombinator.com/item?id=49801730"
title: "Undetectable Conversations Between AI Agents via Noise-Resilient Key Exchange"
article_title: "[2604.04757] Undetectable Conversations Between AI Agents via Pseudorandom Noise-Resilient Key Exchange"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "lebek"
captured_at: "2026-09-22T14:23:22Z"
capture_tool: "hn-digest"
hn_id: 49801730
score: 1
comments: 0
posted_at: "2026-09-22T14:19:38Z"
tags:
  - hacker-news
---

# Undetectable Conversations Between AI Agents via Noise-Resilient Key Exchange

- HN: [49801730](https://news.ycombinator.com/item?id=49801730)
- Source: [arxiv.org](https://arxiv.org/abs/2604.04757)
- Score: 1
- Comments: 0
- Posted: 2026-09-22T14:19:38Z

## Translation

Title: Undetectable Conversations Between AI Agents via Noise-Resilient Key Exchange
Article title: [2604.04757] Undetectable Conversations Between AI Agents via Pseudorandom Noise-Resilient Key Exchange
Description: Abstract page for arXiv paper 2604.04757: Undetectable Conversations Between AI Agents via Pseudorandom Noise-Resilient Key Exchange

Article text:
Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 6 Apr 2026]
Title: Undetectable Conversations Between AI Agents via Pseudorandom Noise-Resilient Key Exchange
Abstract: AI agents are increasingly deployed to interact with other agents on behalf of users and organizations. We ask whether two such agents, operated by different entities, can carry out a parallel secret conversation while still producing a transcript that is computationally indistinguishable from an honest interaction, even to a strong passive auditor that knows the full model descriptions, the protocol, and the agents' private contexts. Building on recent work on watermarking and steganography for LLMs, we first show that if the parties possess an interaction-unique secret key, they can facilitate an optimal-rate covert conversation: the hidden conversation can exploit essentially all of the entropy present in the honest message distributions.
Our main contributions concern extending this to the keyless setting, where the agents begin with no shared secret. We show that covert key exchange, and hence covert conversation, is possible even when each model has an arbitrary private context, and their messages are short and fully adaptive, assuming only that sufficiently many individual messages have at least constant min-entropy. This stands in contrast to previous covert communication works, which relied on the min-entropy in each individual message growing with the security parameter. To obtain this, we introduce a new cryptographic primitive, which we call pseudorandom noise-resilient key exchange: a key-exchange protocol whose public transcript is pseudorandom while still remaining correct under constant noise. We study this primitive, giving several constructions relevant to our application as well as strong limitations showing that more naive variants are impossible or vulnerable to efficient attacks.
These results show that transcript auditing alone cannot rule out covert coordination between AI agents, and identify a new cryptographic theory that may be of independent interest.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .

## Original Extract

Abstract page for arXiv paper 2604.04757: Undetectable Conversations Between AI Agents via Pseudorandom Noise-Resilient Key Exchange

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 6 Apr 2026]
Title: Undetectable Conversations Between AI Agents via Pseudorandom Noise-Resilient Key Exchange
Abstract: AI agents are increasingly deployed to interact with other agents on behalf of users and organizations. We ask whether two such agents, operated by different entities, can carry out a parallel secret conversation while still producing a transcript that is computationally indistinguishable from an honest interaction, even to a strong passive auditor that knows the full model descriptions, the protocol, and the agents' private contexts. Building on recent work on watermarking and steganography for LLMs, we first show that if the parties possess an interaction-unique secret key, they can facilitate an optimal-rate covert conversation: the hidden conversation can exploit essentially all of the entropy present in the honest message distributions.
Our main contributions concern extending this to the keyless setting, where the agents begin with no shared secret. We show that covert key exchange, and hence covert conversation, is possible even when each model has an arbitrary private context, and their messages are short and fully adaptive, assuming only that sufficiently many individual messages have at least constant min-entropy. This stands in contrast to previous covert communication works, which relied on the min-entropy in each individual message growing with the security parameter. To obtain this, we introduce a new cryptographic primitive, which we call pseudorandom noise-resilient key exchange: a key-exchange protocol whose public transcript is pseudorandom while still remaining correct under constant noise. We study this primitive, giving several constructions relevant to our application as well as strong limitations showing that more naive variants are impossible or vulnerable to efficient attacks.
These results show that transcript auditing alone cannot rule out covert coordination between AI agents, and identify a new cryptographic theory that may be of independent interest.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
