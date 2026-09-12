---
source: "https://canonry.ai/blog/classifying-ai-crawlers-user-fetches-server-logs"
hn_url: "https://news.ycombinator.com/item?id=49668078"
title: "Classifying AI Website Fetches on Behalf of Users"
article_title: "Classifying AI Crawlers and User Fetches from Server Logs | Canonry"
image: "https://canonry.ai/og-canonry.png"
author: "arberx"
captured_at: "2026-09-12T03:24:35Z"
capture_tool: "hn-digest"
hn_id: 49668078
score: 1
comments: 0
posted_at: "2026-09-12T02:25:29Z"
tags:
  - hacker-news
---

# Classifying AI Website Fetches on Behalf of Users

- HN: [49668078](https://news.ycombinator.com/item?id=49668078)
- Source: [canonry.ai](https://canonry.ai/blog/classifying-ai-crawlers-user-fetches-server-logs)
- Score: 1
- Comments: 0
- Posted: 2026-09-12T02:25:29Z

## Translation

Title: Classifying AI Website Fetches on Behalf of Users
Article title: Classifying AI Crawlers and User Fetches from Server Logs | Canonry
Description: ChatGPT and Claude can fetch a page through provider infrastructure without sending the user to it. Canonry separates these user fetches from crawlers, citations, and referrals.

Article text:
Canonry For Businesses For Marketing Teams For Agencies Products Case Studies Blog About Talk to us Canonry
Based in New York. We work across markets.
Classifying AI Crawlers and User Fetches from Server Logs
Arber Xhindoli · August 13, 2026 · 3 min read
Human ideas · AI-written · Human-reviewed
Arber Xhindoli developed the ideas and examples, directed the argument, and checked the final article. AI wrote the prose.
When the fetch happens on the provider's network
One AI traffic number hides different events
Server logs are where AI crawlers and user fetches show up. I couldn't find a good self-hosted way to separate the two, so I built it into Canonry.
When the fetch happens on the provider's network
A question in ChatGPT or Claude goes to the provider, not your site. If the app needs a live page, it may send a separate request to your server.
user action -> provider backend -> your site
When that request carries the expected user-agent and matches its published IP range, the path behaves like a server-side proxy. Your origin sees a provider-controlled address instead of the person's browser. The provider gets the page back and may use it in the answer. The person may never visit your site.
Traditional browser analytics start after a page loads and executes JavaScript. In this server-side fetch path, the provider retrieves the HTML without running the page's JavaScript, so the GA4 browser tag and marketing pixels do not fire. The request appears in edge or origin logs while the analytics dashboard records no visit.
OpenAI says certain user actions may visit a page with ChatGPT-User . Anthropic says Claude may use Claude-User to retrieve content for a user's query. These are different from automatic crawlers such as GPTBot , OAI-SearchBot , and ClaudeBot .
OpenAI publishes the IP ranges used by ChatGPT-User , and Anthropic publishes its bot ranges . The providers do not promise that every user-triggered fetch follows this server-side path. A user-agent by itself is only a claim, so unmatched requests have to remain unverified.
One AI traffic number hides different events
Canonry reads server traffic from Cloudflare, Cloud Run, Vercel, and WordPress. It stores hourly rollups by identity, normalized path, response status, and verification state. User fetches stay separate from crawler traffic. Citations come from saved answer evidence. Referrals come from browser visits attributed to an AI source.
During one seven-day window, 947 requests used AI user-fetch identities on canonry.ai . Another 1,667 used four crawler identities.
Of the 927 requests labeled ChatGPT-User , Canonry's stored May 17 IP snapshot matched 64. At 9:27 p.m. ET on August 13, I rechecked the same logs against OpenAI's current file and got 76 matches. The other 851 did not match that snapshot.
Of those unmatched requests, 789 arrived in five hourly bursts. Nearly all returned a redirect or a 404 while probing paths such as /.env and /.ssh/id_rsa . That pattern resembles a scanner using an AI user-agent, though the IP mismatch alone does not identify the caller.
These are requests, not users or citations. The main job is keeping crawls, user fetches, citations, and referrals from becoming one misleading number.
The traffic classifier and source integrations are in the Canonry repository .
No. Every question reaches the provider's backend, but only some questions cause an external page fetch. Those requests can use ChatGPT-User or Claude-User.
When the provider fetches a page server-side, the site's server sees its egress IP and user-agent instead of the person's browser. A match against the published range supports that identity.
No. A fetch proves that a request reached the site. Citation tracking requires inspecting the answer and its source links.
Inspect the technical workflow, run it on your own site, or add live visibility reporting to an agency portal.
How AI traffic appears in server logs
© 2026 Canonry (formerly AI NYC)

## Original Extract

ChatGPT and Claude can fetch a page through provider infrastructure without sending the user to it. Canonry separates these user fetches from crawlers, citations, and referrals.

Canonry For Businesses For Marketing Teams For Agencies Products Case Studies Blog About Talk to us Canonry
Based in New York. We work across markets.
Classifying AI Crawlers and User Fetches from Server Logs
Arber Xhindoli · August 13, 2026 · 3 min read
Human ideas · AI-written · Human-reviewed
Arber Xhindoli developed the ideas and examples, directed the argument, and checked the final article. AI wrote the prose.
When the fetch happens on the provider's network
One AI traffic number hides different events
Server logs are where AI crawlers and user fetches show up. I couldn't find a good self-hosted way to separate the two, so I built it into Canonry.
When the fetch happens on the provider's network
A question in ChatGPT or Claude goes to the provider, not your site. If the app needs a live page, it may send a separate request to your server.
user action -> provider backend -> your site
When that request carries the expected user-agent and matches its published IP range, the path behaves like a server-side proxy. Your origin sees a provider-controlled address instead of the person's browser. The provider gets the page back and may use it in the answer. The person may never visit your site.
Traditional browser analytics start after a page loads and executes JavaScript. In this server-side fetch path, the provider retrieves the HTML without running the page's JavaScript, so the GA4 browser tag and marketing pixels do not fire. The request appears in edge or origin logs while the analytics dashboard records no visit.
OpenAI says certain user actions may visit a page with ChatGPT-User . Anthropic says Claude may use Claude-User to retrieve content for a user's query. These are different from automatic crawlers such as GPTBot , OAI-SearchBot , and ClaudeBot .
OpenAI publishes the IP ranges used by ChatGPT-User , and Anthropic publishes its bot ranges . The providers do not promise that every user-triggered fetch follows this server-side path. A user-agent by itself is only a claim, so unmatched requests have to remain unverified.
One AI traffic number hides different events
Canonry reads server traffic from Cloudflare, Cloud Run, Vercel, and WordPress. It stores hourly rollups by identity, normalized path, response status, and verification state. User fetches stay separate from crawler traffic. Citations come from saved answer evidence. Referrals come from browser visits attributed to an AI source.
During one seven-day window, 947 requests used AI user-fetch identities on canonry.ai . Another 1,667 used four crawler identities.
Of the 927 requests labeled ChatGPT-User , Canonry's stored May 17 IP snapshot matched 64. At 9:27 p.m. ET on August 13, I rechecked the same logs against OpenAI's current file and got 76 matches. The other 851 did not match that snapshot.
Of those unmatched requests, 789 arrived in five hourly bursts. Nearly all returned a redirect or a 404 while probing paths such as /.env and /.ssh/id_rsa . That pattern resembles a scanner using an AI user-agent, though the IP mismatch alone does not identify the caller.
These are requests, not users or citations. The main job is keeping crawls, user fetches, citations, and referrals from becoming one misleading number.
The traffic classifier and source integrations are in the Canonry repository .
No. Every question reaches the provider's backend, but only some questions cause an external page fetch. Those requests can use ChatGPT-User or Claude-User.
When the provider fetches a page server-side, the site's server sees its egress IP and user-agent instead of the person's browser. A match against the published range supports that identity.
No. A fetch proves that a request reached the site. Citation tracking requires inspecting the answer and its source links.
Inspect the technical workflow, run it on your own site, or add live visibility reporting to an agency portal.
How AI traffic appears in server logs
© 2026 Canonry (formerly AI NYC)
