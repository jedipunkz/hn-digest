---
source: "https://arxiv.org/abs/2606.24083"
hn_url: "https://news.ycombinator.com/item?id=50036322"
title: "CAVEWOMAN: How LLMs Behave Under Linguistic Input and Output Compression"
article_title: "[2606.24083] CAVEWOMAN: How Large Language Models Behave Under Linguistic Input and Output Compression"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "farlight"
captured_at: "2026-10-10T20:29:49Z"
capture_tool: "hn-digest"
hn_id: 50036322
score: 2
comments: 0
posted_at: "2026-10-10T19:31:57Z"
tags:
  - hacker-news
---

# CAVEWOMAN: How LLMs Behave Under Linguistic Input and Output Compression

- HN: [50036322](https://news.ycombinator.com/item?id=50036322)
- Source: [arxiv.org](https://arxiv.org/abs/2606.24083)
- Score: 2
- Comments: 0
- Posted: 2026-10-10T19:31:57Z

## Translation

Title: CAVEWOMAN: How LLMs Behave Under Linguistic Input and Output Compression
Article title: [2606.24083] CAVEWOMAN: How Large Language Models Behave Under Linguistic Input and Output Compression
Description: Abstract page for arXiv paper 2606.24083: CAVEWOMAN: How Large Language Models Behave Under Linguistic Input and Output Compression

Article text:
Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 23 Jun 2026]
Title: CAVEWOMAN: How Large Language Models Behave Under Linguistic Input and Output Compression
Abstract: "Talk short. Drop grammar. Save token." This caveman style is widely promoted as a way to cut inference cost, but whether it actually saves anything depends on which channel (the user's prompt or the model's response) is being compressed. We present Cavewoman, a two-channel evaluation protocol that scores every generation on task accuracy, realized per-item cost, and reference-text agreement against the model's unconstrained reference. We evaluate eight models on five datasets at five reduction levels, with both channels measured on the same items. Output compression cuts realized cost on most API models (1.4-2.4x per model, up to 3x in the best case) and on all four open-weight models under public-tier pricing. Input compression has the opposite effect, a strict lose-lose: it raises net cost rather than lowering it (~1.15x on the five-benchmark mean, up to 1.8x on the worst dataset and 2.7x under stronger compression), because models compensate with longer responses even as accuracy collapses. Under the same setting, surface text diverges from the unconstrained reference: on the non-reasoning models, roughly half of all generations are correct yet their surface text no longer entails the model's own unconstrained baseline generation. The divergence survives length-controlled re-scoring, multiple-comparisons correction, and replication under complementary semantic measures. Code and data are available at this https URL .
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

Abstract page for arXiv paper 2606.24083: CAVEWOMAN: How Large Language Models Behave Under Linguistic Input and Output Compression

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 23 Jun 2026]
Title: CAVEWOMAN: How Large Language Models Behave Under Linguistic Input and Output Compression
Abstract: "Talk short. Drop grammar. Save token." This caveman style is widely promoted as a way to cut inference cost, but whether it actually saves anything depends on which channel (the user's prompt or the model's response) is being compressed. We present Cavewoman, a two-channel evaluation protocol that scores every generation on task accuracy, realized per-item cost, and reference-text agreement against the model's unconstrained reference. We evaluate eight models on five datasets at five reduction levels, with both channels measured on the same items. Output compression cuts realized cost on most API models (1.4-2.4x per model, up to 3x in the best case) and on all four open-weight models under public-tier pricing. Input compression has the opposite effect, a strict lose-lose: it raises net cost rather than lowering it (~1.15x on the five-benchmark mean, up to 1.8x on the worst dataset and 2.7x under stronger compression), because models compensate with longer responses even as accuracy collapses. Under the same setting, surface text diverges from the unconstrained reference: on the non-reasoning models, roughly half of all generations are correct yet their surface text no longer entails the model's own unconstrained baseline generation. The divergence survives length-controlled re-scoring, multiple-comparisons correction, and replication under complementary semantic measures. Code and data are available at this https URL .
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
