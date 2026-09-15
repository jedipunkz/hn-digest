---
source: "https://ai-visibility.lastminutedealshq.com/who-blocks-ai-search"
hn_url: "https://news.ycombinator.com/item?id=49711489"
title: "16 of 32 well-known sites block at least one AI search crawler"
article_title: "Which well-known sites block AI search crawlers"
image: "https://ai-visibility.lastminutedealshq.com/og-image.png"
author: "reesecalder"
captured_at: "2026-09-15T13:11:35Z"
capture_tool: "hn-digest"
hn_id: 49711489
score: 1
comments: 0
posted_at: "2026-09-15T12:25:09Z"
tags:
  - hacker-news
---

# 16 of 32 well-known sites block at least one AI search crawler

- HN: [49711489](https://news.ycombinator.com/item?id=49711489)
- Source: [ai-visibility.lastminutedealshq.com](https://ai-visibility.lastminutedealshq.com/who-blocks-ai-search)
- Score: 1
- Comments: 0
- Posted: 2026-09-15T12:25:09Z

## Translation

Title: 16 of 32 well-known sites block at least one AI search crawler
Article title: Which well-known sites block AI search crawlers
Description: A check of 32 recognizable sites: 16 block at least one AI search crawler and 10 are invisible to both ChatGPT and Perplexity. Read from their public robots.txt, September 2026.

Article text:
AI visibility checker llms.txt generator and validator robots.txt checker Crawler census Who blocks AI search Full block list Guides AI crawlers
Which well-known sites block AI search crawlers
A check of 32 recognizable sites: 16 block at least one AI search crawler and 10 are invisible to both ChatGPT and Perplexity. Read from their public robots.txt, September 2026.
I ran the AI-visibility checker on 32 well-known sites to see how many let the AI search crawlers read them. 16 of the 32 block at least one. 10 block both the crawler ChatGPT uses for citations (OAI-SearchBot) and the one Perplexity uses (PerplexityBot), so those sites cannot be cited by either engine right now. Another 6 block only Perplexity.
News publishers make up most of the blockers, and for them it is often a deliberate choice tied to content licensing. This page does not judge that. It reads each site's public robots.txt and reports what the robots.txt allows. For a business that wants to show up in AI answers, a blocked search crawler means the same thing either way: no citations from that engine.
Each site links to its full result. Re-checked automatically, most recently on 7 September 2026. A question mark means the site's robots.txt could not be fetched on the last run, so the honest answer is that I do not know.
Tracking started 2026-09-07. Nothing on this list has changed its robots.txt since then. Any change will be listed here with the date it was first seen.
Correction. npr.org was first published here as allowing both crawlers. That was wrong. The checker could not fetch npr.org's robots.txt and treated the failure as permission. npr.org blocks OAI-SearchBot and PerplexityBot. Corrected 7 September 2026.
Correction. figma.com was first published here as blocking both crawlers. That was wrong. Figma's robots.txt allows OAI-SearchBot and PerplexityBot on its homepage and a curated list of pages using rules anchored with a dollar sign, such as "Allow: /$", and blocks the rest of the site. My matcher did not handle that anchor, so the allow rules were missed and the site read as fully blocked. Corrected 7 September 2026. The table reports homepage access, so figma now shows as reachable, with most of the rest of the site still disallowed.
A green check means the site allows that engine's search crawler. A red mark means its robots.txt blocks it. ChatGPT Search uses OAI-SearchBot; Perplexity uses PerplexityBot.
This page covers 32 names most people recognise. The full list has every site in the census that blocks an AI search crawler, several hundred of them, each with its own result.
Want to know where your own site stands? Check it here , see the full 5,000-site crawler census , or run the same checks from the command line with aicheck .

## Original Extract

A check of 32 recognizable sites: 16 block at least one AI search crawler and 10 are invisible to both ChatGPT and Perplexity. Read from their public robots.txt, September 2026.

AI visibility checker llms.txt generator and validator robots.txt checker Crawler census Who blocks AI search Full block list Guides AI crawlers
Which well-known sites block AI search crawlers
A check of 32 recognizable sites: 16 block at least one AI search crawler and 10 are invisible to both ChatGPT and Perplexity. Read from their public robots.txt, September 2026.
I ran the AI-visibility checker on 32 well-known sites to see how many let the AI search crawlers read them. 16 of the 32 block at least one. 10 block both the crawler ChatGPT uses for citations (OAI-SearchBot) and the one Perplexity uses (PerplexityBot), so those sites cannot be cited by either engine right now. Another 6 block only Perplexity.
News publishers make up most of the blockers, and for them it is often a deliberate choice tied to content licensing. This page does not judge that. It reads each site's public robots.txt and reports what the robots.txt allows. For a business that wants to show up in AI answers, a blocked search crawler means the same thing either way: no citations from that engine.
Each site links to its full result. Re-checked automatically, most recently on 7 September 2026. A question mark means the site's robots.txt could not be fetched on the last run, so the honest answer is that I do not know.
Tracking started 2026-09-07. Nothing on this list has changed its robots.txt since then. Any change will be listed here with the date it was first seen.
Correction. npr.org was first published here as allowing both crawlers. That was wrong. The checker could not fetch npr.org's robots.txt and treated the failure as permission. npr.org blocks OAI-SearchBot and PerplexityBot. Corrected 7 September 2026.
Correction. figma.com was first published here as blocking both crawlers. That was wrong. Figma's robots.txt allows OAI-SearchBot and PerplexityBot on its homepage and a curated list of pages using rules anchored with a dollar sign, such as "Allow: /$", and blocks the rest of the site. My matcher did not handle that anchor, so the allow rules were missed and the site read as fully blocked. Corrected 7 September 2026. The table reports homepage access, so figma now shows as reachable, with most of the rest of the site still disallowed.
A green check means the site allows that engine's search crawler. A red mark means its robots.txt blocks it. ChatGPT Search uses OAI-SearchBot; Perplexity uses PerplexityBot.
This page covers 32 names most people recognise. The full list has every site in the census that blocks an AI search crawler, several hundred of them, each with its own result.
Want to know where your own site stands? Check it here , see the full 5,000-site crawler census , or run the same checks from the command line with aicheck .
