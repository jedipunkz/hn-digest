---
source: "https://blog.slopcop.com/comparing-jev-gemini-and-openai-decision-models/"
hn_url: "https://news.ycombinator.com/item?id=50474571"
title: "Comparing Jev, Gemini, and OpenAI decision models"
article_title: "Comparing Jev, Gemini, and OpenAI decision models"
image: "https://storage.ghost.io/c/f5/e4/f5e49182-2f16-4855-8727-fde7524e784f/content/images/2026/11/night-answers.png"
author: "mfiguiere"
captured_at: "2026-11-12T22:04:26Z"
capture_tool: "hn-digest"
hn_id: 50474571
score: 1
comments: 0
posted_at: "2026-11-12T21:29:57Z"
tags:
  - hacker-news
---

# Comparing Jev, Gemini, and OpenAI decision models

- HN: [50474571](https://news.ycombinator.com/item?id=50474571)
- Source: [blog.slopcop.com](https://blog.slopcop.com/comparing-jev-gemini-and-openai-decision-models/)
- Score: 1
- Comments: 0
- Posted: 2026-11-12T21:29:57Z

## Translation

Title: Comparing Jev, Gemini, and OpenAI decision models
Description: One of the ways that Mjolnir makes your agents smarter is by delivering relevant github events to them so they stay looped in to what's happening with the rest of the team.
Some of these can be done deterministically: activity on your own issues or PRs (second entry from

Article text:
Skip to content
SlopCop
Home Blog
Sign in
Subscribe
Case notes.
Comparing Jev, Gemini, and OpenAI decision models
Jonathan Ellis Nov 12, 2026 3 min read
One of the ways that Mjolnir makes your agents smarter is by delivering relevant github events to them so they stay looped in to what's happening with the rest of the team.
Some of these can be done deterministically: activity on your own issues or PRs (second entry from the top here) is always considered relevant. But what if a teammate opens an issue or a PR that is related to what you're working on? We want to surface those, too.
For events that we can't classify programatically, Mjolnir uses a classifier to determine relevance. (Provided free out of the box via a SlopCop-hosted proxy, with a local-sized "almost as good as GLiDE" model coming soon for those who want to keep their data completely private.) We started with Jev, because it was there first, but when Google released Gemini Decisions I decided to perform a bakeoff to see if we could do better.
We ended up choosing GLiDE from Fastino Labs for our revision. As a bonus, not only does it give the best results for our use case, it's also one of the cheapest options:
It's also open-ish weights . (Commercial use requires a separate agreement if your annual revenue is over $10M.)
The only real ding against GLiDE is that it was the second-slowest service tested; this is where Gemini and OpenAI most excelled. I guess having some of the top inference-speed experts in the world on tap, with massive hardware footprints, is worth something!
I asked Fable to go through our eval notebooks for generalizable lessons:
Never use the incumbent's verdicts as truth, or labels built from them. In our first label set, a strong adjudicator saw a pair only when the first labeller disagreed with Jev; most pairs (where both said no) were never checked. Against those labels Jev's recall looked like 0.91; against an independent panel it is 0.57. Any truth that is filtered through the incumbent's answer flatters the incumbent.
Give the judges more context than the contenders, from more than one vendor. Our judges read the whole session transcript; the contenders see only the bounded production evidence. Two judges from different labs, plus a third model that breaks ties after seeing both rationales unattributed, keeps one model family's quirks out of the truth.
Sample by stratum, weight back to the population, and stratify on something that isn't competing. Relevant pairs are 11% of traffic, so a uniform sample wastes most labels on easy negatives. We stratified on the incumbent's verdict and on an independent reranker score (Voyage, not a contender), then weighted each pair by its stratum's size. Stratifying on a contender's own score would bias the comparison toward it.
When building your dataset, record why a pair is relevant, not only whether. The judges' "own / dependency / adjacent" field showed that every model loses most of its recall on dependency items. That changed the follow-up work from choosing a model to improving evidence.
If you don't have the time or the data to perform a proper eval, Jev is a good place to start: inexpensive, smart, and reasonably fast. Just know that you can probably do better depending on your domain and your priorities, just like with coding models .
David Baker Effendi Nov 6, 2026
Code intelligence has escaped the code
You're paying too much for code review
David Baker Effendi Oct 6, 2026
Catching Resources Held Across Async Suspension
Each new post by email, when it is published.

## Original Extract

One of the ways that Mjolnir makes your agents smarter is by delivering relevant github events to them so they stay looped in to what's happening with the rest of the team.
Some of these can be done deterministically: activity on your own issues or PRs (second entry from

Skip to content
SlopCop
Home Blog
Sign in
Subscribe
Case notes.
Comparing Jev, Gemini, and OpenAI decision models
Jonathan Ellis Nov 12, 2026 3 min read
One of the ways that Mjolnir makes your agents smarter is by delivering relevant github events to them so they stay looped in to what's happening with the rest of the team.
Some of these can be done deterministically: activity on your own issues or PRs (second entry from the top here) is always considered relevant. But what if a teammate opens an issue or a PR that is related to what you're working on? We want to surface those, too.
For events that we can't classify programatically, Mjolnir uses a classifier to determine relevance. (Provided free out of the box via a SlopCop-hosted proxy, with a local-sized "almost as good as GLiDE" model coming soon for those who want to keep their data completely private.) We started with Jev, because it was there first, but when Google released Gemini Decisions I decided to perform a bakeoff to see if we could do better.
We ended up choosing GLiDE from Fastino Labs for our revision. As a bonus, not only does it give the best results for our use case, it's also one of the cheapest options:
It's also open-ish weights . (Commercial use requires a separate agreement if your annual revenue is over $10M.)
The only real ding against GLiDE is that it was the second-slowest service tested; this is where Gemini and OpenAI most excelled. I guess having some of the top inference-speed experts in the world on tap, with massive hardware footprints, is worth something!
I asked Fable to go through our eval notebooks for generalizable lessons:
Never use the incumbent's verdicts as truth, or labels built from them. In our first label set, a strong adjudicator saw a pair only when the first labeller disagreed with Jev; most pairs (where both said no) were never checked. Against those labels Jev's recall looked like 0.91; against an independent panel it is 0.57. Any truth that is filtered through the incumbent's answer flatters the incumbent.
Give the judges more context than the contenders, from more than one vendor. Our judges read the whole session transcript; the contenders see only the bounded production evidence. Two judges from different labs, plus a third model that breaks ties after seeing both rationales unattributed, keeps one model family's quirks out of the truth.
Sample by stratum, weight back to the population, and stratify on something that isn't competing. Relevant pairs are 11% of traffic, so a uniform sample wastes most labels on easy negatives. We stratified on the incumbent's verdict and on an independent reranker score (Voyage, not a contender), then weighted each pair by its stratum's size. Stratifying on a contender's own score would bias the comparison toward it.
When building your dataset, record why a pair is relevant, not only whether. The judges' "own / dependency / adjacent" field showed that every model loses most of its recall on dependency items. That changed the follow-up work from choosing a model to improving evidence.
If you don't have the time or the data to perform a proper eval, Jev is a good place to start: inexpensive, smart, and reasonably fast. Just know that you can probably do better depending on your domain and your priorities, just like with coding models .
David Baker Effendi Nov 6, 2026
Code intelligence has escaped the code
You're paying too much for code review
David Baker Effendi Oct 6, 2026
Catching Resources Held Across Async Suspension
Each new post by email, when it is published.
