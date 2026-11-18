---
source: "https://www.crawlreadiness.com/ai-access-index/2026-11"
hn_url: "https://news.ycombinator.com/item?id=50553175"
title: "AI Access Index: how 684 of the top thousand hostnames treat AI crawlers"
article_title: "AI Access Index, Nov 2026: how 684 of the top 1,000 hostnames treat AI crawlers"
image: "https://www.crawlreadiness.com/api/og/ai-access-index/2026-11"
author: "grubbomatic"
captured_at: "2026-11-18T23:17:53Z"
capture_tool: "hn-digest"
hn_id: 50553175
score: 1
comments: 0
posted_at: "2026-11-18T22:25:52Z"
tags:
  - hacker-news
---

# AI Access Index: how 684 of the top thousand hostnames treat AI crawlers

- HN: [50553175](https://news.ycombinator.com/item?id=50553175)
- Source: [www.crawlreadiness.com](https://www.crawlreadiness.com/ai-access-index/2026-11)
- Score: 1
- Comments: 0
- Posted: 2026-11-18T22:25:52Z

## Translation

Title: AI Access Index: how 684 of the top thousand hostnames treat AI crawlers
Article title: AI Access Index, Nov 2026: how 684 of the top 1,000 hostnames treat AI crawlers
Description: November 2026: 684 of the Tranco top 1,000 answered. 62 of 547 tested sites (11.3%) refused an AI crawler their robots.txt allows; 79 of 124 (63.7%) served one it blocks. Published with the data.

Article text:
AI Access Index, Nov 2026: how 684 of the top 1,000 hostnames treat AI crawlers CR Crawl Readiness Menu Tools ▾ llms.txt Generator Generate an llms.txt for your site Schema Wizard Generate JSON-LD structured data Schema Validator Validate structured data on a URL robots.txt Generator AI-crawler-aware robots.txt robots.txt Validator Audit robots.txt line by line Parity Check Compare human vs AI crawler view MCP Server Run these tools from Claude, Cursor, Cline Chrome Extension Check any site from your toolbar Guides Pricing Blog Access Index Dashboard → AI Access Index · monthly How the top 1,000 websites treat AI crawlers
On the 1st of every month, the same 1,000 sites, checked the same way: what each site's robots.txt tells AI crawlers, and what its server actually does when they arrive. It is not a ranking. Open or closed, a policy is a choice; the index measures which choices sites make, and whether their servers keep to them.
November 2026 edition Measured November 1, 2026 684 of 1000 sites answered Data under CC BY-NC 4.0
What robots.txt tells AI crawlers
Each site's policy is read the way its own report reads it: from what its robots.txt says to the 40 crawlers that fetch pages for AI search and answers, and the 9 that gather training data.
Open to AI 79.7% ( 545 ) AI search: yes · AI training: yes
AI search, no AI training 8.8% ( 60 ) AI search: yes · AI training: no
Selective 3.8% ( 26 ) Some AI companies in, others kept out
Closed to AI 7.7% ( 53 ) AI search: no · AI training: no
Shares of the 684 sites that answered. Open includes 119 sites with no robots.txt at all, which allows everything. Closed includes 13 hosts that close themselves to every crawler, search engines included: mostly API and CDN hostnames rather than a decision about AI.
Does the server do what robots.txt says?
robots.txt is a request, and the server answers first. After the check's own visit, each homepage was asked for again four times, each time with one AI crawler's user-agent string, and the answer was set beside what the site's robots.txt says about that crawler.
Blocked in robots.txt: share of the 684 sites that answered. The server columns: share of the sites whose homepage opened for the check's own request (about 525 ). Served while blocked: share of those that block that crawler. A refusal is a 401, 403, 429 or 503.
This is the first edition. From the December 2026 edition on, this section names every site that changed its AI policy, started or stopped blocking one of the four crawlers, or whose server started or stopped refusing a crawler its robots.txt allows.
Every site in the index, as it answered on November 1, 2026 . A site that is not in the sample can be checked the same way, free.
Share of the 684 sites that answered whose robots.txt disallows / for the crawler, by name or through a group it falls under. Every crawler the check knows is in summary.json ; what each one is, in the crawler reference .
The files and signals written for AI
Shares of the 684 sites that answered.
The same sites every month. The top 1,000 hostnames of Tranco list 8PXYV , the research-standard ranking, downloaded on October 11, 2026 and frozen, so a change in the index is a change in the sites, not in the list. Some are infrastructure hostnames with no website behind them; they count as not answering.
The same day and the same rules. On the 1st of each month, from a GitHub Actions cloud runner (this edition ran 13:27 to 13:37 UTC) : each site once, six at a time, nothing retried. It is the check this site runs for anyone , scoring rules v3 : robots.txt, llms.txt, the homepage and its markdown copy, agents.json and the TDM file as CrawlReadinessBot , then the homepage as GPTBot, OAI-SearchBot, ClaudeBot and PerplexityBot, keeping only the status. A site's own report says what the index says about it.
The four policies. Read from robots.txt's verdicts for the 40 search and answer crawlers and the 9 training crawlers. Open to AI : at most a fifth of the search crawlers blocked and training left mostly open. AI search, no AI training : search open, training blocked (or every training crawler the file names blocked). Closed to AI : every major AI company kept out. Selective : anything else, such as one company left open while the rest are blocked.
Refused and tested. A refusal is a 401, 403, 429 or 503. A crawler's refusal counts only when the homepage opened for the check's own request: when the server refused us too, it could be a wall against every bot. A 2xx to a crawler robots.txt blocks counts either way, so the share of sites that serve a blocked crawler includes some whose homepage did not open for us; the second figure on that tile, and the per-crawler table, leave them out. A site counts as tested when at least one crawler's answer can be read against its robots.txt. Each share names what it is a share of.
Limits. One pass, homepages only. The test sends a crawler's user-agent string, which is what firewall rules match; it does not come from the vendors' IP addresses, so a site that verifies crawlers by address could refuse us and admit the real crawler. Firewalls answer by network and by hour, and a 429 can be a rate limit, so the stricter figure leaves 429-only refusals out. The October 2026 study measured the same sites from a different network with a browser as its baseline, so its numbers are not part of this series.
Quote the numbers anywhere, the press included, with a credit. The files are licensed CC BY-NC 4.0 : free to share and adapt for non-commercial use, with credit. Commercial use of the files, such as in a product or a dataset you sell, needs our permission: ask at support@crawlreadiness.com . Cite as “Crawl Readiness, AI Access Index , November 2026 , www.crawlreadiness.com /ai-access-index/2026-11 .”
results.csv : one row per site in the sample, with its policy, the server's answer to each of the four crawlers and every signal.
changes.csv : every change since the edition before, site by site.
summary.json : every number on this page, with what each is a share of.
If you run a site in the sample and want its row out of the published files, write to support@crawlreadiness.com from an address at that domain.
The same check, free, with the fixes: run it on your site . Sites people have checked with the free tool are listed, unranked, in Checked sites .
Contact: support@crawlreadiness.com · Privacy · Terms

## Original Extract

November 2026: 684 of the Tranco top 1,000 answered. 62 of 547 tested sites (11.3%) refused an AI crawler their robots.txt allows; 79 of 124 (63.7%) served one it blocks. Published with the data.

AI Access Index, Nov 2026: how 684 of the top 1,000 hostnames treat AI crawlers CR Crawl Readiness Menu Tools ▾ llms.txt Generator Generate an llms.txt for your site Schema Wizard Generate JSON-LD structured data Schema Validator Validate structured data on a URL robots.txt Generator AI-crawler-aware robots.txt robots.txt Validator Audit robots.txt line by line Parity Check Compare human vs AI crawler view MCP Server Run these tools from Claude, Cursor, Cline Chrome Extension Check any site from your toolbar Guides Pricing Blog Access Index Dashboard → AI Access Index · monthly How the top 1,000 websites treat AI crawlers
On the 1st of every month, the same 1,000 sites, checked the same way: what each site's robots.txt tells AI crawlers, and what its server actually does when they arrive. It is not a ranking. Open or closed, a policy is a choice; the index measures which choices sites make, and whether their servers keep to them.
November 2026 edition Measured November 1, 2026 684 of 1000 sites answered Data under CC BY-NC 4.0
What robots.txt tells AI crawlers
Each site's policy is read the way its own report reads it: from what its robots.txt says to the 40 crawlers that fetch pages for AI search and answers, and the 9 that gather training data.
Open to AI 79.7% ( 545 ) AI search: yes · AI training: yes
AI search, no AI training 8.8% ( 60 ) AI search: yes · AI training: no
Selective 3.8% ( 26 ) Some AI companies in, others kept out
Closed to AI 7.7% ( 53 ) AI search: no · AI training: no
Shares of the 684 sites that answered. Open includes 119 sites with no robots.txt at all, which allows everything. Closed includes 13 hosts that close themselves to every crawler, search engines included: mostly API and CDN hostnames rather than a decision about AI.
Does the server do what robots.txt says?
robots.txt is a request, and the server answers first. After the check's own visit, each homepage was asked for again four times, each time with one AI crawler's user-agent string, and the answer was set beside what the site's robots.txt says about that crawler.
Blocked in robots.txt: share of the 684 sites that answered. The server columns: share of the sites whose homepage opened for the check's own request (about 525 ). Served while blocked: share of those that block that crawler. A refusal is a 401, 403, 429 or 503.
This is the first edition. From the December 2026 edition on, this section names every site that changed its AI policy, started or stopped blocking one of the four crawlers, or whose server started or stopped refusing a crawler its robots.txt allows.
Every site in the index, as it answered on November 1, 2026 . A site that is not in the sample can be checked the same way, free.
Share of the 684 sites that answered whose robots.txt disallows / for the crawler, by name or through a group it falls under. Every crawler the check knows is in summary.json ; what each one is, in the crawler reference .
The files and signals written for AI
Shares of the 684 sites that answered.
The same sites every month. The top 1,000 hostnames of Tranco list 8PXYV , the research-standard ranking, downloaded on October 11, 2026 and frozen, so a change in the index is a change in the sites, not in the list. Some are infrastructure hostnames with no website behind them; they count as not answering.
The same day and the same rules. On the 1st of each month, from a GitHub Actions cloud runner (this edition ran 13:27 to 13:37 UTC) : each site once, six at a time, nothing retried. It is the check this site runs for anyone , scoring rules v3 : robots.txt, llms.txt, the homepage and its markdown copy, agents.json and the TDM file as CrawlReadinessBot , then the homepage as GPTBot, OAI-SearchBot, ClaudeBot and PerplexityBot, keeping only the status. A site's own report says what the index says about it.
The four policies. Read from robots.txt's verdicts for the 40 search and answer crawlers and the 9 training crawlers. Open to AI : at most a fifth of the search crawlers blocked and training left mostly open. AI search, no AI training : search open, training blocked (or every training crawler the file names blocked). Closed to AI : every major AI company kept out. Selective : anything else, such as one company left open while the rest are blocked.
Refused and tested. A refusal is a 401, 403, 429 or 503. A crawler's refusal counts only when the homepage opened for the check's own request: when the server refused us too, it could be a wall against every bot. A 2xx to a crawler robots.txt blocks counts either way, so the share of sites that serve a blocked crawler includes some whose homepage did not open for us; the second figure on that tile, and the per-crawler table, leave them out. A site counts as tested when at least one crawler's answer can be read against its robots.txt. Each share names what it is a share of.
Limits. One pass, homepages only. The test sends a crawler's user-agent string, which is what firewall rules match; it does not come from the vendors' IP addresses, so a site that verifies crawlers by address could refuse us and admit the real crawler. Firewalls answer by network and by hour, and a 429 can be a rate limit, so the stricter figure leaves 429-only refusals out. The October 2026 study measured the same sites from a different network with a browser as its baseline, so its numbers are not part of this series.
Quote the numbers anywhere, the press included, with a credit. The files are licensed CC BY-NC 4.0 : free to share and adapt for non-commercial use, with credit. Commercial use of the files, such as in a product or a dataset you sell, needs our permission: ask at support@crawlreadiness.com . Cite as “Crawl Readiness, AI Access Index , November 2026 , www.crawlreadiness.com /ai-access-index/2026-11 .”
results.csv : one row per site in the sample, with its policy, the server's answer to each of the four crawlers and every signal.
changes.csv : every change since the edition before, site by site.
summary.json : every number on this page, with what each is a share of.
If you run a site in the sample and want its row out of the published files, write to support@crawlreadiness.com from an address at that domain.
The same check, free, with the fixes: run it on your site . Sites people have checked with the free tool are listed, unranked, in Checked sites .
Contact: support@crawlreadiness.com · Privacy · Terms
