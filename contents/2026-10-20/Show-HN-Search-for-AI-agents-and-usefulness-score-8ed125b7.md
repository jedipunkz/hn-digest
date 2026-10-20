---
source: "https://xtract-production-2efc.up.railway.app/"
hn_url: "https://news.ycombinator.com/item?id=50160830"
title: "Show HN: Search for AI-agents and usefulness score"
article_title: "Xtract search API"
image: ""
author: "Ayobamiu"
captured_at: "2026-10-20T03:41:38Z"
capture_tool: "hn-digest"
hn_id: 50160830
score: 1
comments: 0
posted_at: "2026-10-20T02:57:52Z"
tags:
  - hacker-news
---

# Show HN: Search for AI-agents and usefulness score

- HN: [50160830](https://news.ycombinator.com/item?id=50160830)
- Source: [xtract-production-2efc.up.railway.app](https://xtract-production-2efc.up.railway.app/)
- Score: 1
- Comments: 0
- Posted: 2026-10-20T02:57:52Z

## Translation

Title: Show HN: Search for AI-agents and usefulness score
Article title: Xtract search API
HN text: We presented this project at innovationintelligencehackathon.com and today we are putting a url out for you to test and give us feedback. after talking to teams building agents that research companies to create outreach, we found few things that are very important to them. one of them is recency. When we prompted agents with outdated alert about companies, 4 out of 7 times, they came up with outdated insights, even when they had recent news in context. A bigger model didnt fix it. So we built search for ai-agents with usefulness score on top. So, when you search each page comes with score+price. this solve 3 problems; 1. token, SOTA agents read all pages to make sense of the useful ones 2. cost, when agents pay for pages they read, its is cost-effective to read only useful ones 3. quality, our score consider the task/query at hand and score useful pages, increasing quality of answers. we want to find out fast if we are really up to something here. we'd love to hear from you.

Article text:
Send a query and get each page back with a usefulness score and a price before your agent fetches it, plus a flag when a later page replaces it.
GET /v1/search · 50/day · prices simulated
Query Humane AI Pin , run on 2026-10-14, no usage history. Rank is the search engine's order. Three of the ten pages are flagged: two cover the November 2023 launch and one is a 2024 review. A later result says the Pin was shut down on 28 February 2025.
Get a key
Company email. Shown once. We store only a hash of it.
curl -s "ORIGIN/v1/search?q=Humane+AI+Pin" \
-H "Authorization: Bearer YOUR_KEY"
The raw JSON
{
"query": "Humane AI Pin",
"task": null,
"results": [
{
"url": "https://www.designboom.com/technology/humane-ai-pin-debut-wearable-screenless-smartphone-flash-apps-hands-11-09-2023/",
"title": "humane AI pin shuts down and discontinues all its call, text and smart conversation features",
"snippet": "Humane AI pin no longer works two years after its debut UPDATE (February 19th, 2025): Humane AI Pin announces it is shutting down and discontinuing all of its features, including calling, messaging, AI queries and responses, and cloud access. Less than two years after its debut, the hands-free device that has hoped to become a smartphone substitute no longer works. \u2026",
"keenable_rank": 1,
"usefulness": 0.791,
"price_cents": 3.2686,
"is_new": true,
"superseded_by": null
},
{
"url": "https://www.theverge.com/24126502/humane-ai-pin-review",
"title": "Humane AI Pin review: the post-smartphone future isn\u2019t here yet",
"snippet": "For $699 and $24 a month, this wearable computer promises to free you from your smartphone. There\u2019s only one problem: it just doesn\u2019t work. The idea behind [the Humane AI Pin](https://humane.com/) is a simple one: it\u2019s a phone without a screen. \u2026",
"keenable_rank": 2,
"usefulness": 0.3579,
"price_cents": 1.7527,
"is_new": true,
"superseded_by": {
"url": "https://www.designboom.com/technology/humane-ai-pin-debut-wearable-screenless-smartphone-flash-apps-hands-11-09-2023/",
"title": "humane AI pin shuts down and discontinues all its call, text and smart conversation features",
"published": "2025-02-19",
"event": "Humane discontinued AI Pin services effective February 28, 2025."
}
},
{
"url": "https://mashable.com/article/humane-ai-pin-discontinued-killed",
"title": "Humane has killed its Ai Pin less than a year after its release",
"snippet": "[Humane's Ai Pin](https://mashable.com/article/ai-fail-meta-rabbit-humane) is about to be
[truncated]
Sort by usefulness per cent: usefulness divided by price_cents .
A superseded page is the old story. superseded_by.event is the later fact.
After the agent answers, POST /v1/feedback with the URLs it fetched and the URLs it cited.
q is the search query, up to 300 characters. It is sent to the search engine as typed.
task (optional, up to 500 characters) is what your agent needs answered, for example task=Is+the+Humane+AI+Pin+still+sold%3F . Scores are sharper with it: each page is checked against the task, not just the query.
snippet is a short excerpt, 400 characters by default. Ask for up to 2,000 with snippet_chars . The full page is what the price is for.
usefulness runs from 0 to 1. Below 0.45 the row is marked skip, which is where reading stops.
price_cents is the price to read the page, in cents.
superseded_by is null, or the later page and its event.
scoring_cost_usd is the list-price cost of scoring this search, including when the answer is cached.
curl -s -X POST "ORIGIN/v1/feedback" \
-H "Authorization: Bearer YOUR_KEY" -H "Content-Type: application/json" \
-d '{"query": "Humane AI Pin",
"fetched": ["https://mashable.com/article/humane-ai-pin-discontinued-killed",
"https://wikipedia.org/wiki/Humane_AI_Pin"],
"cited": ["https://mashable.com/article/humane-ai-pin-discontinued-killed"]}'
A follow-up search after this report moved Mashable, fetched and cited, from 0.73 to 0.90. Wikipedia, fetched and left out of the citations, moved from 0.71 to 0.55.
Search results come from Keenable . Scores use OpenAI models.

## Original Extract

We presented this project at innovationintelligencehackathon.com and today we are putting a url out for you to test and give us feedback. after talking to teams building agents that research companies to create outreach, we found few things that are very important to them. one of them is recency. When we prompted agents with outdated alert about companies, 4 out of 7 times, they came up with outdated insights, even when they had recent news in context. A bigger model didnt fix it. So we built search for ai-agents with usefulness score on top. So, when you search each page comes with score+price. this solve 3 problems; 1. token, SOTA agents read all pages to make sense of the useful ones 2. cost, when agents pay for pages they read, its is cost-effective to read only useful ones 3. quality, our score consider the task/query at hand and score useful pages, increasing quality of answers. we want to find out fast if we are really up to something here. we'd love to hear from you.

Send a query and get each page back with a usefulness score and a price before your agent fetches it, plus a flag when a later page replaces it.
GET /v1/search · 50/day · prices simulated
Query Humane AI Pin , run on 2026-10-14, no usage history. Rank is the search engine's order. Three of the ten pages are flagged: two cover the November 2023 launch and one is a 2024 review. A later result says the Pin was shut down on 28 February 2025.
Get a key
Company email. Shown once. We store only a hash of it.
curl -s "ORIGIN/v1/search?q=Humane+AI+Pin" \
-H "Authorization: Bearer YOUR_KEY"
The raw JSON
{
"query": "Humane AI Pin",
"task": null,
"results": [
{
"url": "https://www.designboom.com/technology/humane-ai-pin-debut-wearable-screenless-smartphone-flash-apps-hands-11-09-2023/",
"title": "humane AI pin shuts down and discontinues all its call, text and smart conversation features",
"snippet": "Humane AI pin no longer works two years after its debut UPDATE (February 19th, 2025): Humane AI Pin announces it is shutting down and discontinuing all of its features, including calling, messaging, AI queries and responses, and cloud access. Less than two years after its debut, the hands-free device that has hoped to become a smartphone substitute no longer works. \u2026",
"keenable_rank": 1,
"usefulness": 0.791,
"price_cents": 3.2686,
"is_new": true,
"superseded_by": null
},
{
"url": "https://www.theverge.com/24126502/humane-ai-pin-review",
"title": "Humane AI Pin review: the post-smartphone future isn\u2019t here yet",
"snippet": "For $699 and $24 a month, this wearable computer promises to free you from your smartphone. There\u2019s only one problem: it just doesn\u2019t work. The idea behind [the Humane AI Pin](https://humane.com/) is a simple one: it\u2019s a phone without a screen. \u2026",
"keenable_rank": 2,
"usefulness": 0.3579,
"price_cents": 1.7527,
"is_new": true,
"superseded_by": {
"url": "https://www.designboom.com/technology/humane-ai-pin-debut-wearable-screenless-smartphone-flash-apps-hands-11-09-2023/",
"title": "humane AI pin shuts down and discontinues all its call, text and smart conversation features",
"published": "2025-02-19",
"event": "Humane discontinued AI Pin services effective February 28, 2025."
}
},
{
"url": "https://mashable.com/article/humane-ai-pin-discontinued-killed",
"title": "Humane has killed its Ai Pin less than a year after its release",
"snippet": "[Humane's Ai Pin](https://mashable.com/article/ai-fail-meta-rabbit-humane) is about to be
[truncated]
Sort by usefulness per cent: usefulness divided by price_cents .
A superseded page is the old story. superseded_by.event is the later fact.
After the agent answers, POST /v1/feedback with the URLs it fetched and the URLs it cited.
q is the search query, up to 300 characters. It is sent to the search engine as typed.
task (optional, up to 500 characters) is what your agent needs answered, for example task=Is+the+Humane+AI+Pin+still+sold%3F . Scores are sharper with it: each page is checked against the task, not just the query.
snippet is a short excerpt, 400 characters by default. Ask for up to 2,000 with snippet_chars . The full page is what the price is for.
usefulness runs from 0 to 1. Below 0.45 the row is marked skip, which is where reading stops.
price_cents is the price to read the page, in cents.
superseded_by is null, or the later page and its event.
scoring_cost_usd is the list-price cost of scoring this search, including when the answer is cached.
curl -s -X POST "ORIGIN/v1/feedback" \
-H "Authorization: Bearer YOUR_KEY" -H "Content-Type: application/json" \
-d '{"query": "Humane AI Pin",
"fetched": ["https://mashable.com/article/humane-ai-pin-discontinued-killed",
"https://wikipedia.org/wiki/Humane_AI_Pin"],
"cited": ["https://mashable.com/article/humane-ai-pin-discontinued-killed"]}'
A follow-up search after this report moved Mashable, fetched and cited, from 0.73 to 0.90. Wikipedia, fetched and left out of the citations, moved from 0.71 to 0.55.
Search results come from Keenable . Scores use OpenAI models.
