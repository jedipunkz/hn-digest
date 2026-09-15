---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49708448"
title: "How do you even investigate Segfaults with OpenAI Agents?"
article_title: ""
image: ""
author: "gnatolf"
captured_at: "2026-09-15T07:16:04Z"
capture_tool: "hn-digest"
hn_id: 49708448
score: 1
comments: 2
posted_at: "2026-09-15T06:18:23Z"
tags:
  - hacker-news
---

# How do you even investigate Segfaults with OpenAI Agents?

- HN: [49708448](https://news.ycombinator.com/item?id=49708448)
- Score: 1
- Comments: 2
- Posted: 2026-09-15T06:18:23Z

## Translation

Title: How do you even investigate Segfaults with OpenAI Agents?
HN text: I have a "simple" segfault because of an out of bounds read/write in one of my libraries. I cannot get any Astra/Sol/Luna agent to finish the probe work on this without it running into the 'security work' blocker. This is an open code base, locally checked out. Any tips or hints? There is not a sliver of nefarious intent behind me asking "why does a high index XY cause an out-of-bound read". And I can't see why this would trigger any safeguards.

## Original Extract

I have a "simple" segfault because of an out of bounds read/write in one of my libraries. I cannot get any Astra/Sol/Luna agent to finish the probe work on this without it running into the 'security work' blocker. This is an open code base, locally checked out. Any tips or hints? There is not a sliver of nefarious intent behind me asking "why does a high index XY cause an out-of-bound read". And I can't see why this would trigger any safeguards.

