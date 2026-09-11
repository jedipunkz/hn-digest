---
source: "https://hnslop.nilsherzig.com/"
hn_url: "https://news.ycombinator.com/item?id=49661856"
title: "Show HN: Extension to filter LLM written articles"
article_title: "hnslop"
image: ""
author: "nilsherzig"
captured_at: "2026-09-11T17:34:42Z"
capture_tool: "hn-digest"
hn_id: 49661856
score: 1
comments: 0
posted_at: "2026-09-11T17:14:49Z"
tags:
  - hacker-news
---

# Show HN: Extension to filter LLM written articles

- HN: [49661856](https://news.ycombinator.com/item?id=49661856)
- Source: [hnslop.nilsherzig.com](https://hnslop.nilsherzig.com/)
- Score: 1
- Comments: 0
- Posted: 2026-09-11T17:14:49Z

## Translation

Title: Show HN: Extension to filter LLM written articles
Article title: hnslop
HN text: Firefox Extension/Userscript and API to get Pangram scores for all articles on the hackernews frontpage. The extension allows you to hide articles with a high score. This is about detecting posts written by LLMs, not posts about AI. Feel free to use the API to build your own tooling/readers. Big thanks to https://news.ycombinator.com/user?id=salahadawi for providing the data :).

Article text:
This server turns Salah Adawi's Hacker News AI Detector
into a cached JSON API. Allowing you to programmatically get pangram checks for every hackernews post that
reached the frontpage. Feel free to use the hosted instance at hnslop.nilsherzig.com . Big thanks to Salah (i assume that he had
to sell multiple internal organs to pay for his pangram usage).
Please keep in mind that Salah is (as of the time of writing) using Pangram v3.3, which isnt the most up to
date model from pangram.
curl 'https://hnslop.nilsherzig.com/v1/posts?ids=49582582,49541888' | jq
# {
# "posts" : [
# {
# "id" : 49582582 ,
# "detector" : {
# "ai_score" : 33 ,
# "url" : "https://www.salahadawi.com/hacker-news-ai-detector/49582582"
# } ,
# "cache_status" : "hit"
# } ,
# {
# "id" : 49541888 ,
# "detector" : {
# "ai_score" : 99 ,
# "url" : "https://www.salahadawi.com/hacker-news-ai-detector/49541888"
# } ,
# "cache_status" : "hit"
# }
# ]
# }
curl 'https://hnslop.nilsherzig.com/v1/posts/49582582' | jq
# {
# "id" : 49582582 ,
# "detector" : {
# "ai_score" : 33 ,
# "url" : "https://www.salahadawi.com/hacker-news-ai-detector/49582582"
# } ,
# "cache_status" : "hit"
# }
Example userscript client
Download the Firefox extension

## Original Extract

Firefox Extension/Userscript and API to get Pangram scores for all articles on the hackernews frontpage. The extension allows you to hide articles with a high score. This is about detecting posts written by LLMs, not posts about AI. Feel free to use the API to build your own tooling/readers. Big thanks to https://news.ycombinator.com/user?id=salahadawi for providing the data :).

This server turns Salah Adawi's Hacker News AI Detector
into a cached JSON API. Allowing you to programmatically get pangram checks for every hackernews post that
reached the frontpage. Feel free to use the hosted instance at hnslop.nilsherzig.com . Big thanks to Salah (i assume that he had
to sell multiple internal organs to pay for his pangram usage).
Please keep in mind that Salah is (as of the time of writing) using Pangram v3.3, which isnt the most up to
date model from pangram.
curl 'https://hnslop.nilsherzig.com/v1/posts?ids=49582582,49541888' | jq
# {
# "posts" : [
# {
# "id" : 49582582 ,
# "detector" : {
# "ai_score" : 33 ,
# "url" : "https://www.salahadawi.com/hacker-news-ai-detector/49582582"
# } ,
# "cache_status" : "hit"
# } ,
# {
# "id" : 49541888 ,
# "detector" : {
# "ai_score" : 99 ,
# "url" : "https://www.salahadawi.com/hacker-news-ai-detector/49541888"
# } ,
# "cache_status" : "hit"
# }
# ]
# }
curl 'https://hnslop.nilsherzig.com/v1/posts/49582582' | jq
# {
# "id" : 49582582 ,
# "detector" : {
# "ai_score" : 33 ,
# "url" : "https://www.salahadawi.com/hacker-news-ai-detector/49582582"
# } ,
# "cache_status" : "hit"
# }
Example userscript client
Download the Firefox extension
