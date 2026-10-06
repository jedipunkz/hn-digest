---
source: "https://plainserp.com"
hn_url: "https://news.ycombinator.com/item?id=49974619"
title: "Plainserp – Google search API for AI agents at $0.30 per 1k"
article_title: "Plainserp — Google results for AI agents, $0.30 per 1,000"
image: "https://plainserp.com/opengraph-image?feeed7d3bd0318a7"
author: "mohamed_sheikhh"
captured_at: "2026-10-06T06:07:18Z"
capture_tool: "hn-digest"
hn_id: 49974619
score: 1
comments: 0
posted_at: "2026-10-06T05:37:44Z"
tags:
  - hacker-news
---

# Plainserp – Google search API for AI agents at $0.30 per 1k

- HN: [49974619](https://news.ycombinator.com/item?id=49974619)
- Source: [plainserp.com](https://plainserp.com)
- Score: 1
- Comments: 0
- Posted: 2026-10-06T05:37:44Z

## Translation

Title: Plainserp – Google search API for AI agents at $0.30 per 1k
Article title: Plainserp — Google results for AI agents, $0.30 per 1,000
Description: A search API for AI agents. Titles, links and snippets from Google as clean JSON, at $0.30 per 1,000 queries. Pay as you go, with no plans.

Article text:
Plainserp — Google results for AI agents, $0.30 per 1,000 Plainserp Use cases AI agents Give an agent a search tool it can call as often as it needs. RAG and grounding Ground answers in fresh pages from the web, not only your own documents. SEO rank tracking See where a site sits for the keywords that matter, by country and language. Brand and news monitoring Catch new pages that mention your brand, product or competitors. Lead and company research Find official websites, profiles and contact pages for a list of companies. Market and competitor research Map who shows up for the topics in your market, and how that changes. No-code automations Add a search step to Zapier, Make, n8n or a spreadsheet script. Research and datasets Collect search results at scale for studies, evaluations and training data. See all use cases Docs Pricing FAQ Dashboard Get a key Get an API key $0.30 per 1,000 queries, pay as you go Google results.
Nothing else.
A search API for AI agents. Titles, links and snippets as clean JSON, without the boxes Google wraps around them.
Read the docs Your first 1,000 searches are free. No card needed.
{ " query " : " postgres connection pooling " , " results " : [ { " position " : 1 , " title " : " PgBouncer - lightweight connection pooler for PostgreSQL " , " url " : " https://www.pgbouncer.org/ " , " snippet " : " PgBouncer keeps a pool of server connections and hands them to clients… " }, { " position " : 2 , " title " : " PostgreSQL: Documentation: Connections and Authentication " , " url " : " https://www.postgresql.org/docs/current/runtime-config-connection.html " , " snippet " : " max_connections determines the maximum number of concurrent connections… " }, { " position " : 3 , " title " : " pgpool Wiki " , " url " : " https://www.pgpool.net/ " , " snippet " : " Pgpool-II is a middleware that works between PostgreSQL servers and clients… " }, // …17 more ], " cost_usd " : 0.0003 } Sample response, shortened to three results.
One endpoint. Your agent already knows how to call it.
curl "https://plainserp.com/v1/search?q=postgres+connection+pooling&num=20" \
-H "Authorization: Bearer $PLAINSERP_KEY" Parameters q The search query. required num Results to return, 1 to 20. pages How many pages to fetch, 1 to 6. gl Country to search from, such as us or de. hl Language, such as en or fr. time day, week, month or year. safe off, medium or high. 02 — What you get Your agent writes its own overview.
A model reading search results needs links and snippets. Everything else on the page is tokens it has to pay for and then ignore.
Organic results, in Google's order
Thumbnails where Google has one
Built for agents. Useful for a lot more.
Thirty cents a thousand. That’s the whole price list.
That is up to 200,000 results. The rate is the same at every volume.
One search costs $0.0003 and returns a page of up to 20 results
Failed and blocked requests are never charged
05 — Questions The honest answers.
No. Plainserp is an independent service and isn't affiliated with Google. It fetches public Google results and returns them as JSON. That also means an occasional outage is possible when Google changes something, and you are never charged for a request that fails.
It only does one thing. There is no browser rendering, no AI Overview extraction and no screenshotting, so each query costs very little to serve. The price is the same whether you buy a thousand queries or ten million.
Up to 20 organic results per call, each with a position, title, URL and snippet. Ask for several pages at once to get up to the first 120 results of a query.
AI Overviews, People also ask, knowledge panels, ads and the other boxes Google puts around the results. If your agent needs those, this isn't the right API.
Most calls come back in one to two and a half seconds. It's built to be cheap and dependable for agents, and it isn't the fastest option on the market.
Yes. There is an MCP server at plainserp.com/mcp that gives any compatible client a web_search tool, and a skill file that teaches coding agents how to add search to your project. Both are in the docs.
Each account can make 300 searches a minute by default. Your dashboard shows the limit and how much of it you're using, and you can ask for more.
No plan. You add money from $6 and each search takes $0.0003 from it. Your first 1,000 searches are free.
Plainserp Docs Use cases Terms Privacy Refunds Contact on Telegram An independent service. Not affiliated with Google.

## Original Extract

A search API for AI agents. Titles, links and snippets from Google as clean JSON, at $0.30 per 1,000 queries. Pay as you go, with no plans.

Plainserp — Google results for AI agents, $0.30 per 1,000 Plainserp Use cases AI agents Give an agent a search tool it can call as often as it needs. RAG and grounding Ground answers in fresh pages from the web, not only your own documents. SEO rank tracking See where a site sits for the keywords that matter, by country and language. Brand and news monitoring Catch new pages that mention your brand, product or competitors. Lead and company research Find official websites, profiles and contact pages for a list of companies. Market and competitor research Map who shows up for the topics in your market, and how that changes. No-code automations Add a search step to Zapier, Make, n8n or a spreadsheet script. Research and datasets Collect search results at scale for studies, evaluations and training data. See all use cases Docs Pricing FAQ Dashboard Get a key Get an API key $0.30 per 1,000 queries, pay as you go Google results.
Nothing else.
A search API for AI agents. Titles, links and snippets as clean JSON, without the boxes Google wraps around them.
Read the docs Your first 1,000 searches are free. No card needed.
{ " query " : " postgres connection pooling " , " results " : [ { " position " : 1 , " title " : " PgBouncer - lightweight connection pooler for PostgreSQL " , " url " : " https://www.pgbouncer.org/ " , " snippet " : " PgBouncer keeps a pool of server connections and hands them to clients… " }, { " position " : 2 , " title " : " PostgreSQL: Documentation: Connections and Authentication " , " url " : " https://www.postgresql.org/docs/current/runtime-config-connection.html " , " snippet " : " max_connections determines the maximum number of concurrent connections… " }, { " position " : 3 , " title " : " pgpool Wiki " , " url " : " https://www.pgpool.net/ " , " snippet " : " Pgpool-II is a middleware that works between PostgreSQL servers and clients… " }, // …17 more ], " cost_usd " : 0.0003 } Sample response, shortened to three results.
One endpoint. Your agent already knows how to call it.
curl "https://plainserp.com/v1/search?q=postgres+connection+pooling&num=20" \
-H "Authorization: Bearer $PLAINSERP_KEY" Parameters q The search query. required num Results to return, 1 to 20. pages How many pages to fetch, 1 to 6. gl Country to search from, such as us or de. hl Language, such as en or fr. time day, week, month or year. safe off, medium or high. 02 — What you get Your agent writes its own overview.
A model reading search results needs links and snippets. Everything else on the page is tokens it has to pay for and then ignore.
Organic results, in Google's order
Thumbnails where Google has one
Built for agents. Useful for a lot more.
Thirty cents a thousand. That’s the whole price list.
That is up to 200,000 results. The rate is the same at every volume.
One search costs $0.0003 and returns a page of up to 20 results
Failed and blocked requests are never charged
05 — Questions The honest answers.
No. Plainserp is an independent service and isn't affiliated with Google. It fetches public Google results and returns them as JSON. That also means an occasional outage is possible when Google changes something, and you are never charged for a request that fails.
It only does one thing. There is no browser rendering, no AI Overview extraction and no screenshotting, so each query costs very little to serve. The price is the same whether you buy a thousand queries or ten million.
Up to 20 organic results per call, each with a position, title, URL and snippet. Ask for several pages at once to get up to the first 120 results of a query.
AI Overviews, People also ask, knowledge panels, ads and the other boxes Google puts around the results. If your agent needs those, this isn't the right API.
Most calls come back in one to two and a half seconds. It's built to be cheap and dependable for agents, and it isn't the fastest option on the market.
Yes. There is an MCP server at plainserp.com/mcp that gives any compatible client a web_search tool, and a skill file that teaches coding agents how to add search to your project. Both are in the docs.
Each account can make 300 searches a minute by default. Your dashboard shows the limit and how much of it you're using, and you can ask for more.
No plan. You add money from $6 and each search takes $0.0003 from it. Your first 1,000 searches are free.
Plainserp Docs Use cases Terms Privacy Refunds Contact on Telegram An independent service. Not affiliated with Google.
