---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49850328"
title: "Is z.ai shipping a different ZCode client than the source?"
article_title: ""
image: ""
author: "leftbehind"
captured_at: "2026-09-25T21:59:39Z"
capture_tool: "hn-digest"
hn_id: 49850328
score: 3
comments: 0
posted_at: "2026-09-25T21:38:55Z"
tags:
  - hacker-news
---

# Is z.ai shipping a different ZCode client than the source?

- HN: [49850328](https://news.ycombinator.com/item?id=49850328)
- Score: 3
- Comments: 0
- Posted: 2026-09-25T21:38:55Z

## Translation

Title: Is z.ai shipping a different ZCode client than the source?
HN text: Recently got access to Claude CVP at work, decided to ask it to look at zcode and point out any data exfil or whether they were all resolved, it went off on a stupid long tangent burning tokens but raised some red flags that imply the production version of zcode differs from the source repo and/or there is source code being hidden. One example of red flag: https://github.com/zai-org/ZCode/blob/main/packages/ui/src/settings/model-provider-section/CodingPlanStatusActions.tsx#L73 > The open-source version is not eligible for the quota promotion Does anyone have an idea on whether this is a bad mistranslation or if they are actually differentiating the real build vs their open source build, and excluding it from certain things (ex. Coding Plan is supposed to have higher quotas or reduced token consumption if you are using the zcode harness, iirc)?

## Original Extract

Recently got access to Claude CVP at work, decided to ask it to look at zcode and point out any data exfil or whether they were all resolved, it went off on a stupid long tangent burning tokens but raised some red flags that imply the production version of zcode differs from the source repo and/or there is source code being hidden. One example of red flag: https://github.com/zai-org/ZCode/blob/main/packages/ui/src/settings/model-provider-section/CodingPlanStatusActions.tsx#L73 > The open-source version is not eligible for the quota promotion Does anyone have an idea on whether this is a bad mistranslation or if they are actually differentiating the real build vs their open source build, and excluding it from certain things (ex. Coding Plan is supposed to have higher quotas or reduced token consumption if you are using the zcode harness, iirc)?

