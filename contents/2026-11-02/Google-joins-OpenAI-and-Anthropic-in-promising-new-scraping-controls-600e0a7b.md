---
source: "https://www.engadget.com/2303225/google-openai-anthropic-ai-crawling-controls/"
hn_url: "https://news.ycombinator.com/item?id=50331090"
title: "Google joins OpenAI and Anthropic in promising new scraping controls"
article_title: "Google Joins OpenAI And Anthropic In Promising New Scraping Controls"
image: "https://www.engadget.com/img/gallery/google-joins-openai-and-anthropic-in-promising-new-scraping-controls/l-intro-1793442580.jpg"
author: "geox"
captured_at: "2026-11-02T00:19:19Z"
capture_tool: "hn-digest"
hn_id: 50331090
score: 2
comments: 0
posted_at: "2026-11-01T23:39:35Z"
tags:
  - hacker-news
---

# Google joins OpenAI and Anthropic in promising new scraping controls

- HN: [50331090](https://news.ycombinator.com/item?id=50331090)
- Source: [www.engadget.com](https://www.engadget.com/2303225/google-openai-anthropic-ai-crawling-controls/)
- Score: 2
- Comments: 0
- Posted: 2026-11-01T23:39:35Z

## Translation

Title: Google joins OpenAI and Anthropic in promising new scraping controls
Article title: Google Joins OpenAI And Anthropic In Promising New Scraping Controls
Description: Google, OpenAI, Anthropic and Microsoft pledge to support a new common signal for robots.txt that blocks scraping for AI training.

Article text:
Google Joins OpenAI And Anthropic In Promising New Scraping Controls
Latest
Google joins OpenAI and Anthropic in promising new scraping controls
The new robots.txt directive will prevent the companies from scraping content for generative AI training.
tadamichi/Shutterstock
Add Engadget on Google:
Preferred Source
Google Discover
Google , Anthropic, OpenAI and Microsoft have committed to honoring a new robots.txt directive, which will allow publishers to specify their preferences when it comes to the scraping of their content for generative AI training. The companies have decided to adopt a proposal under consideration at IETF, the open standards organization for the internet, written by Anthropic technical lead for standards Kevin Kelley.
As OpenAI and Google explained in their announcements, website owners currently have to identify individual crawlers and add separate rules for each provider. Kelley's proposal will change that by getting multiple AI companies to recognize and honor a single common signal. Publishers have to set up a new user agent in the robots.txt file that's specifically for AI training, which will look like this:
User-Agent: __usage_ai-training
Disallow: /
Allow: /press-kit/
That will block crawlers specifically meant to gather data for AI training. The "Allow" section gives website operators the option to allow scraping for AI training for specific parts of the page, such as if there's a press kit on it. But this new user-agent section doesn't prevent crawling altogether or prevent content from showing up on Google search. For instance, Google will specifically interpret it as a stop sign specifically for the "Google-Extended" user agent, which crawls websites for materials used to train Gemini models.
Another thing to note is that this directive cannot undo any past scraping the companies have already done. If they'd crawled a page previously, they already have its content. This will just prevent them from scraping websites or specific pages in the future. Anthropic has already started supporting the new directive, while Microsoft said it will start supporting it on February 1, 2027. Google said it will support it in early 2027, as well, while OpenAI pledged to implement the new signal in the coming months.

## Original Extract

Google, OpenAI, Anthropic and Microsoft pledge to support a new common signal for robots.txt that blocks scraping for AI training.

Google Joins OpenAI And Anthropic In Promising New Scraping Controls
Latest
Google joins OpenAI and Anthropic in promising new scraping controls
The new robots.txt directive will prevent the companies from scraping content for generative AI training.
tadamichi/Shutterstock
Add Engadget on Google:
Preferred Source
Google Discover
Google , Anthropic, OpenAI and Microsoft have committed to honoring a new robots.txt directive, which will allow publishers to specify their preferences when it comes to the scraping of their content for generative AI training. The companies have decided to adopt a proposal under consideration at IETF, the open standards organization for the internet, written by Anthropic technical lead for standards Kevin Kelley.
As OpenAI and Google explained in their announcements, website owners currently have to identify individual crawlers and add separate rules for each provider. Kelley's proposal will change that by getting multiple AI companies to recognize and honor a single common signal. Publishers have to set up a new user agent in the robots.txt file that's specifically for AI training, which will look like this:
User-Agent: __usage_ai-training
Disallow: /
Allow: /press-kit/
That will block crawlers specifically meant to gather data for AI training. The "Allow" section gives website operators the option to allow scraping for AI training for specific parts of the page, such as if there's a press kit on it. But this new user-agent section doesn't prevent crawling altogether or prevent content from showing up on Google search. For instance, Google will specifically interpret it as a stop sign specifically for the "Google-Extended" user agent, which crawls websites for materials used to train Gemini models.
Another thing to note is that this directive cannot undo any past scraping the companies have already done. If they'd crawled a page previously, they already have its content. This will just prevent them from scraping websites or specific pages in the future. Anthropic has already started supporting the new directive, while Microsoft said it will start supporting it on February 1, 2027. Google said it will support it in early 2027, as well, while OpenAI pledged to implement the new signal in the coming months.
