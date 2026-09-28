---
source: "https://arxiv.org/abs/2608.10257"
hn_url: "https://news.ycombinator.com/item?id=49881549"
title: "GenRec: An LLM-Backed Recommendation Ranker at NetflixConference"
article_title: "[2608.10257] GenRec: An LLM-Backed Recommendation Ranker at Netflix"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "throwthrowrow"
captured_at: "2026-09-28T18:31:21Z"
capture_tool: "hn-digest"
hn_id: 49881549
score: 3
comments: 0
posted_at: "2026-09-28T17:37:47Z"
tags:
  - hacker-news
---

# GenRec: An LLM-Backed Recommendation Ranker at NetflixConference

- HN: [49881549](https://news.ycombinator.com/item?id=49881549)
- Source: [arxiv.org](https://arxiv.org/abs/2608.10257)
- Score: 3
- Comments: 0
- Posted: 2026-09-28T17:37:47Z

## Translation

Title: GenRec: An LLM-Backed Recommendation Ranker at NetflixConference
Article title: [2608.10257] GenRec: An LLM-Backed Recommendation Ranker at Netflix
Description: Abstract page for arXiv paper 2608.10257: GenRec: An LLM-Backed Recommendation Ranker at Netflix

Article text:
Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Information Retrieval
[Submitted on 10 Aug 2026 ( v1 ), last revised 21 Aug 2026 (this version, v2)]
Title: GenRec: An LLM-Backed Recommendation Ranker at Netflix
Abstract: Large language models (LLMs) are reshaping recommender systems by enabling richer modeling of users, content, and context directly in natural language. At Netflix, we are exploring this direction through GenRec, an LLM-backed recommendation ranker built on top of an in-house foundational LLM. GenRec follows a two-phase framework: Phase 1 adapts an open-source LLM to Netflix data, developing deep understanding of the catalog and member behavior while balancing capabilities such as content understanding and instruction following. Phase 2 post-trains this foundation model with recommendation-ranking specific data, labels, and reward signals, aiming to align the ranker with business requirements and long-term member satisfaction.
This paper focuses on Phase 2 and the transition from a traditional discriminative ranker with thousands of engineered features to an LLM-backed ranker driven by verbalized user histories and context. We describe our design for input verbalization and context engineering, post-training data construction, reward integration, model architecture, and a cost-constrained serving design based on a prefill-only inference approach. We report results from a large-scale A/B test comparing GenRec against the current production ranker model, where we show that a GenRec model trained with substantially fewer Phase-2 labeled training examples and input signals can achieve statistically significant gains in offline and online metrics. We discuss how LLM-backed recommenders could shift the recommendation paradigm: from feature engineering to context engineering, and from bespoke architectures to shared foundation backbones. We also outline practical lessons for serving such systems under real-world resource constraints.
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

Abstract page for arXiv paper 2608.10257: GenRec: An LLM-Backed Recommendation Ranker at Netflix

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Information Retrieval
[Submitted on 10 Aug 2026 ( v1 ), last revised 21 Aug 2026 (this version, v2)]
Title: GenRec: An LLM-Backed Recommendation Ranker at Netflix
Abstract: Large language models (LLMs) are reshaping recommender systems by enabling richer modeling of users, content, and context directly in natural language. At Netflix, we are exploring this direction through GenRec, an LLM-backed recommendation ranker built on top of an in-house foundational LLM. GenRec follows a two-phase framework: Phase 1 adapts an open-source LLM to Netflix data, developing deep understanding of the catalog and member behavior while balancing capabilities such as content understanding and instruction following. Phase 2 post-trains this foundation model with recommendation-ranking specific data, labels, and reward signals, aiming to align the ranker with business requirements and long-term member satisfaction.
This paper focuses on Phase 2 and the transition from a traditional discriminative ranker with thousands of engineered features to an LLM-backed ranker driven by verbalized user histories and context. We describe our design for input verbalization and context engineering, post-training data construction, reward integration, model architecture, and a cost-constrained serving design based on a prefill-only inference approach. We report results from a large-scale A/B test comparing GenRec against the current production ranker model, where we show that a GenRec model trained with substantially fewer Phase-2 labeled training examples and input signals can achieve statistically significant gains in offline and online metrics. We discuss how LLM-backed recommenders could shift the recommendation paradigm: from feature engineering to context engineering, and from bespoke architectures to shared foundation backbones. We also outline practical lessons for serving such systems under real-world resource constraints.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
