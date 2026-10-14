---
source: "https://www.jcchouinard.com/google-search-console-ai-mode-click-tracking/"
hn_url: "https://news.ycombinator.com/item?id=50091642"
title: "Clicks, Impressions and Positions from AI Mode in Google Search Console"
article_title: "Clicks, Impressions and Positions from AI Mode in Google Search Console - JC Chouinard"
image: "https://www.jcchouinard.com/wp-content/uploads/2026/01/AI-Mode-in-Google-Search-Console.png"
author: "jcchouinard"
captured_at: "2026-10-14T18:20:44Z"
capture_tool: "hn-digest"
hn_id: 50091642
score: 1
comments: 0
posted_at: "2026-10-14T17:37:10Z"
tags:
  - hacker-news
---

# Clicks, Impressions and Positions from AI Mode in Google Search Console

- HN: [50091642](https://news.ycombinator.com/item?id=50091642)
- Source: [www.jcchouinard.com](https://www.jcchouinard.com/google-search-console-ai-mode-click-tracking/)
- Score: 1
- Comments: 0
- Posted: 2026-10-14T17:37:10Z

## Translation

Title: Clicks, Impressions and Positions from AI Mode in Google Search Console
Article title: Clicks, Impressions and Positions from AI Mode in Google Search Console - JC Chouinard
Description: Google has confirmed that Google AI Mode traffic data is in Google Search Console, but that there are no AI mode specific report. All the data get bundled Learn Google Search Console by JC Chouinard

Article text:
Clicks, Impressions and Positions from AI Mode in Google Search Console - JC Chouinard
-->
Subscribe now to Datacamp and save 63% - Subscribe now -->
JC Chouinard
JC Chouinard
menu
About
Clicks, Impressions and Positions from AI Mode in Google Search Console
Google has confirmed that Google AI Mode traffic data is in Google Search Console , but that there are no AI mode specific report. All the data get bundled in the performance report . In this article, I am interested in how the clicks are tracked, what gets counted as a position, what gets impressions, or not. Positions from clicks in AI Mode are showing up in Google Search Console in a more complex matter than for positions from clicks in AI Overviews .
Tracking AI Mode Clicks and Impressions in Google Search Console Data
For this experiment, I used Google Bulk Data Export , a Google data source that contains upsampled data from the Google Search Console performance report (yet still containing anonymized queries). Here is what I have tried out.
Ran ‘xyzk’ queries, clicked links
Searched for rarely visited pages
Recorded screen to correlate searches
Ran queries in multiple browsers
Tested logged-in versus logged-out states
Asked colleagues to run similar queries
Targeted personal blog and dead websites
Generated high volume to reveal queries
Analyzed “Google Links” position impact
Analyzed results on different websites
Summary of the AI Mode Click Tracking Case Study
Clicks in AI Mode show up in Google Search Console, but the queries are anonymized.
Clicks on Inline Links: tracked
Clicks on sidebar citations: tracked
Clicks on sidebar citations with view more: tracked
Clicks on subsequent queries in a conversation: Tracked as a new page.
Sometimes two links can be counted with same position
Running 100s of impressions in AI Mode isn’t enough to trigger non-anonymized queries.
How Do Click Show Up in Google Search Console?
Clicks in AI Mode do appear in Google Search Console, but the specific queries are almost exclusively anonymized . Increasing volume (100+ impressions) does not force query data to appear. The caveat is that Google links are also considered in the Position calculation.
Google Links are Tracked in AI Mode
When I ran this query, I have realized that my clicks were considered as position 8 and 9 instead of 1 and two.
When I tried to figure out how this could have happened, I noticed that the only way that would work is if all links are considered as position.
Aggregated Traffic in Google Search Console
If you click on twice on the same URL for different queries, since they are anonymized, they will be combined together.
If you have your first Click on /abc url in position 1 and your second click on /abc url in position 5, you will get one row in Google Search Console showing the date, the path, the query = null, and an. average position of 3 ((1+5)/2).
See if I can Generate Enough Query Volume for the Queries not to be Anonymized
I ran the query “la patrie hotels” 500 times in separate browsers and behind VPN, separate tabs and click on links, to see if we have enough volume to make sure that the queries are not anonymized. It looks like it wasn’t enough. Led to 101 clicks and none of them had query data.
1 rank = 1 position in GSC. All impressions and clicks are shown.
Show All Citations Sidebar (e.g. Links)
When you click on link icon, or show all, you get other links. Each link is tracked as a position.
Barry Schwartz made a post about AI Mode queries surfacing in GSC . This is very compelling and doesn’t necessarily make the current experiment wrong. One example that was out of scope of the current experiment what a user asking a prompt in AI Overviews and then expanding with a subsequent prompt that triggers AI Mode. It is possible that this may be calculated.
Another possibility is that Google simply changed. When I tested, AI mode was on google.ai domain, and now it is deployed to google.com. Other tracking changes probably have occured, someone should re-experiment with that. Util then, look at the regular expression to find AI Mode types of query patterns .
Senior SEO Strategist at Tripadvisor, ex- Seek.com.au. Specialized in technical SEO and AI Search. Writer in Python, Information Retrieval, SEO and machine learning. Guest author at SearchEngineJournal, SearchEngineLand and OnCrawl.
How to use Google Search Console (GSC) – Complete Guide
How to Validate Your Site in Google Search Console (GSC Example)
Upload Sitemaps to Google Search Console with the API (Python)
Sitemap Could not be Read (Sitemap Couldn’t fetch examples) in Google Search Console

## Original Extract

Google has confirmed that Google AI Mode traffic data is in Google Search Console, but that there are no AI mode specific report. All the data get bundled Learn Google Search Console by JC Chouinard

Clicks, Impressions and Positions from AI Mode in Google Search Console - JC Chouinard
-->
Subscribe now to Datacamp and save 63% - Subscribe now -->
JC Chouinard
JC Chouinard
menu
About
Clicks, Impressions and Positions from AI Mode in Google Search Console
Google has confirmed that Google AI Mode traffic data is in Google Search Console , but that there are no AI mode specific report. All the data get bundled in the performance report . In this article, I am interested in how the clicks are tracked, what gets counted as a position, what gets impressions, or not. Positions from clicks in AI Mode are showing up in Google Search Console in a more complex matter than for positions from clicks in AI Overviews .
Tracking AI Mode Clicks and Impressions in Google Search Console Data
For this experiment, I used Google Bulk Data Export , a Google data source that contains upsampled data from the Google Search Console performance report (yet still containing anonymized queries). Here is what I have tried out.
Ran ‘xyzk’ queries, clicked links
Searched for rarely visited pages
Recorded screen to correlate searches
Ran queries in multiple browsers
Tested logged-in versus logged-out states
Asked colleagues to run similar queries
Targeted personal blog and dead websites
Generated high volume to reveal queries
Analyzed “Google Links” position impact
Analyzed results on different websites
Summary of the AI Mode Click Tracking Case Study
Clicks in AI Mode show up in Google Search Console, but the queries are anonymized.
Clicks on Inline Links: tracked
Clicks on sidebar citations: tracked
Clicks on sidebar citations with view more: tracked
Clicks on subsequent queries in a conversation: Tracked as a new page.
Sometimes two links can be counted with same position
Running 100s of impressions in AI Mode isn’t enough to trigger non-anonymized queries.
How Do Click Show Up in Google Search Console?
Clicks in AI Mode do appear in Google Search Console, but the specific queries are almost exclusively anonymized . Increasing volume (100+ impressions) does not force query data to appear. The caveat is that Google links are also considered in the Position calculation.
Google Links are Tracked in AI Mode
When I ran this query, I have realized that my clicks were considered as position 8 and 9 instead of 1 and two.
When I tried to figure out how this could have happened, I noticed that the only way that would work is if all links are considered as position.
Aggregated Traffic in Google Search Console
If you click on twice on the same URL for different queries, since they are anonymized, they will be combined together.
If you have your first Click on /abc url in position 1 and your second click on /abc url in position 5, you will get one row in Google Search Console showing the date, the path, the query = null, and an. average position of 3 ((1+5)/2).
See if I can Generate Enough Query Volume for the Queries not to be Anonymized
I ran the query “la patrie hotels” 500 times in separate browsers and behind VPN, separate tabs and click on links, to see if we have enough volume to make sure that the queries are not anonymized. It looks like it wasn’t enough. Led to 101 clicks and none of them had query data.
1 rank = 1 position in GSC. All impressions and clicks are shown.
Show All Citations Sidebar (e.g. Links)
When you click on link icon, or show all, you get other links. Each link is tracked as a position.
Barry Schwartz made a post about AI Mode queries surfacing in GSC . This is very compelling and doesn’t necessarily make the current experiment wrong. One example that was out of scope of the current experiment what a user asking a prompt in AI Overviews and then expanding with a subsequent prompt that triggers AI Mode. It is possible that this may be calculated.
Another possibility is that Google simply changed. When I tested, AI mode was on google.ai domain, and now it is deployed to google.com. Other tracking changes probably have occured, someone should re-experiment with that. Util then, look at the regular expression to find AI Mode types of query patterns .
Senior SEO Strategist at Tripadvisor, ex- Seek.com.au. Specialized in technical SEO and AI Search. Writer in Python, Information Retrieval, SEO and machine learning. Guest author at SearchEngineJournal, SearchEngineLand and OnCrawl.
How to use Google Search Console (GSC) – Complete Guide
How to Validate Your Site in Google Search Console (GSC Example)
Upload Sitemaps to Google Search Console with the API (Python)
Sitemap Could not be Read (Sitemap Couldn’t fetch examples) in Google Search Console
