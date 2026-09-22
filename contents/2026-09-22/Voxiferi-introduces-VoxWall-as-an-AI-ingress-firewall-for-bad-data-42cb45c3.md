---
source: "https://opensourcewatch.beehiiv.com/p/new-voxiferi-introduces-voxwall-as-an-ai-ingress-firewall-for-bad-data"
hn_url: "https://news.ycombinator.com/item?id=49808404"
title: "Voxiferi introduces VoxWall as an 'AI ingress firewall' for bad data"
article_title: "Voxiferi introduces VoxWall as an ‘AI ingress firewall’ for bad data"
image: "https://media.beehiiv.com/cdn-cgi/image/fit=scale-down,quality=80,width=1200,onerror=redirect/uploads/asset/file/5d5f04d9-8670-4df7-8f6c-de91bdef05f2/AI_Firewall.jpeg?t=1790103372"
author: "CrankyBear"
captured_at: "2026-09-22T21:49:09Z"
capture_tool: "hn-digest"
hn_id: 49808404
score: 1
comments: 0
posted_at: "2026-09-22T21:30:35Z"
tags:
  - hacker-news
---

# Voxiferi introduces VoxWall as an 'AI ingress firewall' for bad data

- HN: [49808404](https://news.ycombinator.com/item?id=49808404)
- Source: [opensourcewatch.beehiiv.com](https://opensourcewatch.beehiiv.com/p/new-voxiferi-introduces-voxwall-as-an-ai-ingress-firewall-for-bad-data)
- Score: 1
- Comments: 0
- Posted: 2026-09-22T21:30:35Z

## Translation

Title: Voxiferi introduces VoxWall as an 'AI ingress firewall' for bad data
Article title: Voxiferi introduces VoxWall as an ‘AI ingress firewall’ for bad data
Description: Open-source veteran Dick Morrell’s Voxiferi Labs has open-sourced VoxWall, a policy-driven gateway intended to stop bad data before it reaches AI training, Retrieval-Augmented Generation, inference, or workloads.

Article text:
Voxiferi introduces VoxWall as an ‘AI ingress firewall’ for bad data Open Source Watch Authors Login Subscribe Home
Voxiferi introduces VoxWall as an ‘AI ingress firewall’ for bad data
Voxiferi introduces VoxWall as an ‘AI ingress firewall’ for bad data
Open-source veteran Dick Morrell’s Voxiferi Labs has open-sourced VoxWall, a policy-driven gateway intended to stop bad data before it reaches AI training, Retrieval-Augmented Generation, inference, or workloads.
The next-generation open-source AI firewall.
The idea behind Voxiferi Labs' new open-source VoxWall , a firewall for bad AI training data, is simple. Block poor-quality data from ever reaching an AI's Large Language Model (LLM) in the first place.
VoxWall parses incoming records, validates them against defined schemas and policies, assesses their usefulness and expected processing cost, then either forwards them with an audit trail or quarantines them .
Although related to what Voxferi calls Morrell’s Law of AI data , which states AI accuracy and operating cost are primarily determined by the structure and quality of data at intake, VoxWall isn't an implementation of that law.
Instead, according to the company, "VoxWall is purely a perimeter firewall. It refuses malicious model files, config injection, credential leakage, and unsafe tool calls at the boundary, before they reach a model."
In short, rather than primarily asking a model to recognize prompt injection, poisoned documents, malformed records, or unreliable retrieved content after the fact, VoxWall’s thesis is that unsuitable inputs should be identified or rejected at ingestion.
The far more advanced Morrell's Law implementations measure structure coefficients, score provenance, and validate data at intake. To run Morrell’s Law in practice, you'll need VoxGate or VoxSuite.
Still, “AI needs an ingress firewall,” said Voxiferi founder Richard “Dick” Morrell , and VoxWall is a good, open-source start. The AI wall arrives as enterprises confront the security complications of retrieval-augmented generation (RAG) . RAG systems routinely ingest documents, emails, websites, PDFs, trouble tickets, and knowledge-base entries, then convert that data into LLM prompts. It's a fine idea, but RAG can also vastly increase the chances of getting dangerous answers.
Worse, RAG makes the AI data ingestion chain an attractive target for hackers. An attacker who can place malicious text in a company's source materials can infect an AI via indirect prompt injection . Poorly governed document ingestion can also expose confidential material to the wrong user, preserve stale or fraudulent records, or burn up inference and embedding resources on low-value data.
Existing RAG security guidance calls for layered protections: document hashing and integrity checks at ingestion, source provenance, authorization before retrieval, isolation of retrieved text from system instructions, scanning for injection attempts, logging, output validation, and tightly constrained access to downstream tools.
VoxWall provides a firewall that blocks bad records from ever entering a RAG pipeline. For example, the system could potentially reject malformed JSON, records that fail an organization’s declared schema, items from an unauthorized source, or inputs whose quality might poison the LLM or simply don't justify their downstream compute cost.
The approach is closer to a combination of an API gateway, data-quality system, provenance service, and security control point than to a conventional LLM guardrail. The key difference is architectural: rather than trying only to detect dangerous or nonsensical material after it has been embedded, VoxWall seeks to prevent questionable material from entering the AI supply chain.
That distinction could matter most for organizations operating in regulated, sensitive, or disconnected environments. Voxiferi has emphasized sovereign, local, and closed-network deployments in its broader positioning. These are all areas where sending raw data to a cloud moderation service, or relying on a hosted model provider to inspect inputs, may be unacceptable.
VoxWall’s central insight makes perfect sense: AI systems cannot reliably repair bad source data simply by applying more parameters or compute. NIST’s generative-AI risk guidance likewise treats the integrity and confidentiality of AI code and data as core security concerns, rather than model-level concerns.
But the practical assessment will depend on implementation details that aren't available yet. So to judge this, admittedly an alpha program, you'll need to join the 110-thousand people who've already downloaded it and check it out for yourself.
Morrell tells me, "VoxWall is the world's very first AI appliance builder to allow you to take a redundant PC or mini server and turn it into an offline listening solution to place in front of an LLM to stop threats dead. It's based on RHEL / Rocky Linux 9 (because I'm ex-Red Hat) with a four-part installer that stops known AI attacks dead."
The downloadable VoxWall installer is a shell script wrapped in a zip file. The program is open-sourced under the Mozilla license.
If VoxWall delivers on its promises, its strongest role may be as a pre-RAG and pre-inference intake control for companies building private or sovereign AI systems, which could make the data boundary, not the model, the more important place to enforce AI security policy. Given LLM guardrails' lousy security record , we desperately need a new and better approach to AI security; VoxWall may be it.
Noteworthy Linux, AI, and open-source stories:
US District Court Decision in AI’s Favor Worries Open-Source Developers
Why AI companies are really pumping the brakes on their models
OpenAI and Anthropic could be two of the biggest AI IPOs Wall Street has seen in years.
But investors don’t have to wait for those names to hit the public markets to get exposure to the AI boom.
MarketBeat’s 7 AI Stocks to Buy Now report reveals 7 publicly traded companies already positioned to benefit as the next wave of AI investment moves beyond the private model providers.
These are the stocks investors can buy today, before the IPO crowd rushes in.
What's what with open-source news.

## Original Extract

Open-source veteran Dick Morrell’s Voxiferi Labs has open-sourced VoxWall, a policy-driven gateway intended to stop bad data before it reaches AI training, Retrieval-Augmented Generation, inference, or workloads.

Voxiferi introduces VoxWall as an ‘AI ingress firewall’ for bad data Open Source Watch Authors Login Subscribe Home
Voxiferi introduces VoxWall as an ‘AI ingress firewall’ for bad data
Voxiferi introduces VoxWall as an ‘AI ingress firewall’ for bad data
Open-source veteran Dick Morrell’s Voxiferi Labs has open-sourced VoxWall, a policy-driven gateway intended to stop bad data before it reaches AI training, Retrieval-Augmented Generation, inference, or workloads.
The next-generation open-source AI firewall.
The idea behind Voxiferi Labs' new open-source VoxWall , a firewall for bad AI training data, is simple. Block poor-quality data from ever reaching an AI's Large Language Model (LLM) in the first place.
VoxWall parses incoming records, validates them against defined schemas and policies, assesses their usefulness and expected processing cost, then either forwards them with an audit trail or quarantines them .
Although related to what Voxferi calls Morrell’s Law of AI data , which states AI accuracy and operating cost are primarily determined by the structure and quality of data at intake, VoxWall isn't an implementation of that law.
Instead, according to the company, "VoxWall is purely a perimeter firewall. It refuses malicious model files, config injection, credential leakage, and unsafe tool calls at the boundary, before they reach a model."
In short, rather than primarily asking a model to recognize prompt injection, poisoned documents, malformed records, or unreliable retrieved content after the fact, VoxWall’s thesis is that unsuitable inputs should be identified or rejected at ingestion.
The far more advanced Morrell's Law implementations measure structure coefficients, score provenance, and validate data at intake. To run Morrell’s Law in practice, you'll need VoxGate or VoxSuite.
Still, “AI needs an ingress firewall,” said Voxiferi founder Richard “Dick” Morrell , and VoxWall is a good, open-source start. The AI wall arrives as enterprises confront the security complications of retrieval-augmented generation (RAG) . RAG systems routinely ingest documents, emails, websites, PDFs, trouble tickets, and knowledge-base entries, then convert that data into LLM prompts. It's a fine idea, but RAG can also vastly increase the chances of getting dangerous answers.
Worse, RAG makes the AI data ingestion chain an attractive target for hackers. An attacker who can place malicious text in a company's source materials can infect an AI via indirect prompt injection . Poorly governed document ingestion can also expose confidential material to the wrong user, preserve stale or fraudulent records, or burn up inference and embedding resources on low-value data.
Existing RAG security guidance calls for layered protections: document hashing and integrity checks at ingestion, source provenance, authorization before retrieval, isolation of retrieved text from system instructions, scanning for injection attempts, logging, output validation, and tightly constrained access to downstream tools.
VoxWall provides a firewall that blocks bad records from ever entering a RAG pipeline. For example, the system could potentially reject malformed JSON, records that fail an organization’s declared schema, items from an unauthorized source, or inputs whose quality might poison the LLM or simply don't justify their downstream compute cost.
The approach is closer to a combination of an API gateway, data-quality system, provenance service, and security control point than to a conventional LLM guardrail. The key difference is architectural: rather than trying only to detect dangerous or nonsensical material after it has been embedded, VoxWall seeks to prevent questionable material from entering the AI supply chain.
That distinction could matter most for organizations operating in regulated, sensitive, or disconnected environments. Voxiferi has emphasized sovereign, local, and closed-network deployments in its broader positioning. These are all areas where sending raw data to a cloud moderation service, or relying on a hosted model provider to inspect inputs, may be unacceptable.
VoxWall’s central insight makes perfect sense: AI systems cannot reliably repair bad source data simply by applying more parameters or compute. NIST’s generative-AI risk guidance likewise treats the integrity and confidentiality of AI code and data as core security concerns, rather than model-level concerns.
But the practical assessment will depend on implementation details that aren't available yet. So to judge this, admittedly an alpha program, you'll need to join the 110-thousand people who've already downloaded it and check it out for yourself.
Morrell tells me, "VoxWall is the world's very first AI appliance builder to allow you to take a redundant PC or mini server and turn it into an offline listening solution to place in front of an LLM to stop threats dead. It's based on RHEL / Rocky Linux 9 (because I'm ex-Red Hat) with a four-part installer that stops known AI attacks dead."
The downloadable VoxWall installer is a shell script wrapped in a zip file. The program is open-sourced under the Mozilla license.
If VoxWall delivers on its promises, its strongest role may be as a pre-RAG and pre-inference intake control for companies building private or sovereign AI systems, which could make the data boundary, not the model, the more important place to enforce AI security policy. Given LLM guardrails' lousy security record , we desperately need a new and better approach to AI security; VoxWall may be it.
Noteworthy Linux, AI, and open-source stories:
US District Court Decision in AI’s Favor Worries Open-Source Developers
Why AI companies are really pumping the brakes on their models
OpenAI and Anthropic could be two of the biggest AI IPOs Wall Street has seen in years.
But investors don’t have to wait for those names to hit the public markets to get exposure to the AI boom.
MarketBeat’s 7 AI Stocks to Buy Now report reveals 7 publicly traded companies already positioned to benefit as the next wave of AI investment moves beyond the private model providers.
These are the stocks investors can buy today, before the IPO crowd rushes in.
What's what with open-source news.
