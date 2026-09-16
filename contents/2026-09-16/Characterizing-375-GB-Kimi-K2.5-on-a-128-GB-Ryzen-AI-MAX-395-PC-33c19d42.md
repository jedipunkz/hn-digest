---
source: "https://zenodo.org/records/22755791"
hn_url: "https://news.ycombinator.com/item?id=49733551"
title: "Characterizing ~375 GB Kimi K2.5 on a 128 GB Ryzen AI MAX+ 395 PC"
article_title: "QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference | Zenodo"
image: ""
author: "NHApplied"
captured_at: "2026-09-16T21:55:32Z"
capture_tool: "hn-digest"
hn_id: 49733551
score: 1
comments: 0
posted_at: "2026-09-16T21:52:46Z"
tags:
  - hacker-news
---

# Characterizing ~375 GB Kimi K2.5 on a 128 GB Ryzen AI MAX+ 395 PC

- HN: [49733551](https://news.ycombinator.com/item?id=49733551)
- Source: [zenodo.org](https://zenodo.org/records/22755791)
- Score: 1
- Comments: 0
- Posted: 2026-09-16T21:52:46Z

## Translation

Title: Characterizing ~375 GB Kimi K2.5 on a 128 GB Ryzen AI MAX+ 395 PC
Article title: QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference | Zenodo
Description: QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference is a complementary systems-characterization paper to the QES v2.7 white paper. The study examines where generated-token time is spent when executing Kimi K2.5 on a single 128 GB AMD Ryzen AI MA
[truncated]

Article text:
QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference | Zenodo
Skip to main
You are using an outdated browser. Please upgrade your browser to improve your experience.
Zenodo is currently experiencing slowness and intermittent outages due to heavy automated traffic from bots and AI crawlers. We are aware of the problem, and our team is focused on stabilizing the service. Thank you for your patience.
QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference
QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference is a complementary systems-characterization paper to the QES v2.7 white paper.
The study examines where generated-token time is spent when executing Kimi K2.5 on a single 128 GB AMD Ryzen AI MAX+ 395 system using storage-backed bounded expert residency. It measures generated-only expert traffic, routing locality, cache effectiveness, sustained decode behaviour and package energy.
The measured 128-token workload contains 61,440 generated expert requests. A bounded locality cache records 4,734 generated-only cache hits (7.705%), reducing expert-store traffic from a 7.126 GiB/token no-cache logical requirement to 6.577 GiB/token and avoiding 70.285 GiB of expert traffic over the run. Mean expert-delivery wait falls by approximately 5%, while repeated cache-enabled runs remain close to 0.438 tokens/s compared with the frozen 0.432091702 tokens/s baseline. A package-power repeat measures 113.47 J/generated token versus 117.50 J/generated token for the no-cache reference.
The results support storage-backed bounded residency as a feasible engineering route for very large sparse Mixture-of-Experts models where full model residency is unavailable and latency tolerance exists. The work does not claim interactive serving parity, universal model compatibility or production readiness.
This publication intentionally reports measured behaviour and architecture-level findings without disclosing the private QES implementation, source code, expert-store construction, scheduling logic, residency-control mechanisms or model-graph integration.
Parent publication: Nigel Hutchinson / NH Applied, QES: Bounded-Memory Execution of Large Mixture-of-Experts Models on a Single 128 GB System — From Qwen3-235B to Kimi K2.5 , v2.7. DOI: 10.5281/zenodo.22730031
QES_Characterization_Complementary_Paper_v1_0.pdf
More info on how stats are collected....
10.5281/zenodo.22755791
Markdown
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.22755791.svg)](https://doi.org/10.5281/zenodo.22755791)
reStructuredText
.. image:: https://zenodo.org/badge/DOI/10.5281/zenodo.22755791.svg
:target: https://doi.org/10.5281/zenodo.22755791
HTML
<a href="https://doi.org/10.5281/zenodo.22755791"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.22755791.svg" alt="DOI"></a>
Image URL
https://zenodo.org/badge/DOI/10.5281/zenodo.22755791.svg
Target URL
https://doi.org/10.5281/zenodo.22755791
Resource type
Report
Publisher
Zenodo
Languages
English
Rights
Powered by
CERN Data Centre & InvenioRDM
This site uses cookies. Find out more on how we use cookies

## Original Extract

QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference is a complementary systems-characterization paper to the QES v2.7 white paper. The study examines where generated-token time is spent when executing Kimi K2.5 on a single 128 GB AMD Ryzen AI MA
[truncated]

QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference | Zenodo
Skip to main
You are using an outdated browser. Please upgrade your browser to improve your experience.
Zenodo is currently experiencing slowness and intermittent outages due to heavy automated traffic from bots and AI crawlers. We are aware of the problem, and our team is focused on stabilizing the service. Thank you for your patience.
QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference
QES Characterization: Data-Path, Cache Locality and Energy in Single-Node Trillion-Parameter MoE Inference is a complementary systems-characterization paper to the QES v2.7 white paper.
The study examines where generated-token time is spent when executing Kimi K2.5 on a single 128 GB AMD Ryzen AI MAX+ 395 system using storage-backed bounded expert residency. It measures generated-only expert traffic, routing locality, cache effectiveness, sustained decode behaviour and package energy.
The measured 128-token workload contains 61,440 generated expert requests. A bounded locality cache records 4,734 generated-only cache hits (7.705%), reducing expert-store traffic from a 7.126 GiB/token no-cache logical requirement to 6.577 GiB/token and avoiding 70.285 GiB of expert traffic over the run. Mean expert-delivery wait falls by approximately 5%, while repeated cache-enabled runs remain close to 0.438 tokens/s compared with the frozen 0.432091702 tokens/s baseline. A package-power repeat measures 113.47 J/generated token versus 117.50 J/generated token for the no-cache reference.
The results support storage-backed bounded residency as a feasible engineering route for very large sparse Mixture-of-Experts models where full model residency is unavailable and latency tolerance exists. The work does not claim interactive serving parity, universal model compatibility or production readiness.
This publication intentionally reports measured behaviour and architecture-level findings without disclosing the private QES implementation, source code, expert-store construction, scheduling logic, residency-control mechanisms or model-graph integration.
Parent publication: Nigel Hutchinson / NH Applied, QES: Bounded-Memory Execution of Large Mixture-of-Experts Models on a Single 128 GB System — From Qwen3-235B to Kimi K2.5 , v2.7. DOI: 10.5281/zenodo.22730031
QES_Characterization_Complementary_Paper_v1_0.pdf
More info on how stats are collected....
10.5281/zenodo.22755791
Markdown
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.22755791.svg)](https://doi.org/10.5281/zenodo.22755791)
reStructuredText
.. image:: https://zenodo.org/badge/DOI/10.5281/zenodo.22755791.svg
:target: https://doi.org/10.5281/zenodo.22755791
HTML
<a href="https://doi.org/10.5281/zenodo.22755791"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.22755791.svg" alt="DOI"></a>
Image URL
https://zenodo.org/badge/DOI/10.5281/zenodo.22755791.svg
Target URL
https://doi.org/10.5281/zenodo.22755791
Resource type
Report
Publisher
Zenodo
Languages
English
Rights
Powered by
CERN Data Centre & InvenioRDM
This site uses cookies. Find out more on how we use cookies
