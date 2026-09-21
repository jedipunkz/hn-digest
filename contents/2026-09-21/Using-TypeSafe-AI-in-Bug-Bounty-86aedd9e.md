---
source: "https://www.lampysecurity.com/post/using-typesafe-ai-in-bug-bounty"
hn_url: "https://news.ycombinator.com/item?id=49788625"
title: "Using TypeSafe AI in Bug Bounty"
article_title: "Using TypeSafe AI in Bug Bounty"
image: "https://static.wixstatic.com/media/5b6a44_a0f13fc37977474fac28e30f3ae2e0f1~mv2.jpg/v1/fill/w_1000,h_558,al_c,q_85,usm_0.66_1.00_0.01/5b6a44_a0f13fc37977474fac28e30f3ae2e0f1~mv2.jpg"
author: "lampysecurity"
captured_at: "2026-09-21T16:15:27Z"
capture_tool: "hn-digest"
hn_id: 49788625
score: 1
comments: 0
posted_at: "2026-09-21T15:34:32Z"
tags:
  - hacker-news
---

# Using TypeSafe AI in Bug Bounty

- HN: [49788625](https://news.ycombinator.com/item?id=49788625)
- Source: [www.lampysecurity.com](https://www.lampysecurity.com/post/using-typesafe-ai-in-bug-bounty)
- Score: 1
- Comments: 0
- Posted: 2026-09-21T15:34:32Z

## Translation

Title: Using TypeSafe AI in Bug Bounty
Description: Today I got access to TypeSafe. If you aren

Article text:
-->
Using TypeSafe AI in Bug Bounty
top of page Home
Search Using TypeSafe AI in Bug Bounty
Today I got access to TypeSafe. If you aren't familiar with the product, it's an ultra-fast and incredibly cheap AI model. When I say cheap, I'm talking $0.042 per million input tokens, with output tokens completely free .
They already had $5 credited to my account, and even with fairly heavy usage and zero optimization, I've only spent $1.29 on the first day. But just wait until you hear what I was able to accomplish with that.
So, what's different about TypeSafe?
It doesn't work like a traditional chatbot such as ChatGPT or Claude. You don't have conversations with it or ask it questions in the same way. What it does excel at is categorization and decision-making at extremely high speeds .
It can accept three types of questions:
Choice – Returns a predefined multiple-choice answer along with a confidence/probability score.
Score – Returns a score based on a defined rubric along with a confidence/probability score.
Noul – Returns a simple true or false.
This structure allows the AI to respond in near real time, enabling fast decision-making without the latency you normally experience with most AI models.
Once I got access, I immediately wanted to incorporate it into my bug bounty workflow. I wanted something that could handle both large scale and high speed .
When you're doing bug bounty reconnaissance and enumeration, you can quickly end up with a massive number of websites and subdomains to process. Traditionally, you have a few options, each with its own downsides:
Manual review – Can take days to weeks to complete
Eyeballer – Takes minutes, but results can be mixed because of the limited model
Frontier models – Powerful, but expensive. They can consume millions of tokens just to review webpages
Local LLMs – Flexible, but slow, and they don't always have the context windows needed to process full webpages
With my first test using TypeSafe, I was able to process 17 subdomains per second .
That's approximately 1,040 domains per minute .
To accomplish this, I first came up with a set of questions broad enough to apply to almost every website.
Does the visible page text below contain enough information to judge what this page actually is?
This helps eliminate false negatives. Some websites time out or don't load correctly because many of them are non-production environments.
What kind of page is this, based on its title, headers, and visible text?
Clearly, in a bug bounty situation, you want to pull out things like S3 buckets, directory listings, and internal tools. "Broken" is a bit too broad of a category, and I would probably break it up further in future versions.
Does this look like an internal admin, staff, or Ops tool that wasn't meant to be public?
Pretty self-explanatory. We're continuing to look for accidental exposure.
Does this look like a login or management tool for a network security application, VPN gateway, or remote access product?
While many of these products could be classified as internal tools, we want to specifically call them out because they are often intentionally exposed to the internet.
Does this look like the login page of an IT management, RMM, or monitoring tool?
This includes products like Grafana, SolarWinds, Splunk, etc.
Does this response directly expose sensitive data itself rather than just linking to it?
Because TypeSafe is looking at the HTML itself rather than screenshots, we can quickly identify things like API keys sitting directly on a landing page. This would be rare, but it's also something that's too easy to overlook.
How interesting is this host for security testing?
This is the first time we're using the Score question type. It gives us a useful number to work with when determining whether a host is worth investigating further.
The responses were saved to a CSV file, making them easy to review and even easier for AI to filter through. Once I had the results, I had Claude take the top 20 highest-scoring results across different categories and put them into text files.
I use the Chrome extension Open Multiple URLs to quickly review the results manually. After a quick review, I pass the interesting results to Claude for further analysis to look for sensitive or exposed content.
This was just a quick post to show what Jev is able to do. I think Jev is going to be a valuable tool for anyone in the offensive security space, and I can't wait to see some of the other projects people build with it. We're already starting to see them pop up.
https://github.com/browser-use/jev-ultrafast
This was only my first day using Jev, but working with this speed, cheap cost and flexibility is what sells me.
Instead of sending thousands of hosts directly to an expensive frontier model or spending hours manually reviewing them, Jev can act as a high-speed filtering layer. This is a must have for your enumeration layer at scale.
For offensive security, I think that's where this gets really interesting. The goal isn't to replace manual testing or more powerful AI models. It's to make sure you're spending your time and compute on the things that actually deserve it.
Go ahead and check it out. Let me know if you make any cool offensive security projects with it.
https://typesafe.ai/
Cyber Security - Technology - Life

## Original Extract

Today I got access to TypeSafe. If you aren

-->
Using TypeSafe AI in Bug Bounty
top of page Home
Search Using TypeSafe AI in Bug Bounty
Today I got access to TypeSafe. If you aren't familiar with the product, it's an ultra-fast and incredibly cheap AI model. When I say cheap, I'm talking $0.042 per million input tokens, with output tokens completely free .
They already had $5 credited to my account, and even with fairly heavy usage and zero optimization, I've only spent $1.29 on the first day. But just wait until you hear what I was able to accomplish with that.
So, what's different about TypeSafe?
It doesn't work like a traditional chatbot such as ChatGPT or Claude. You don't have conversations with it or ask it questions in the same way. What it does excel at is categorization and decision-making at extremely high speeds .
It can accept three types of questions:
Choice – Returns a predefined multiple-choice answer along with a confidence/probability score.
Score – Returns a score based on a defined rubric along with a confidence/probability score.
Noul – Returns a simple true or false.
This structure allows the AI to respond in near real time, enabling fast decision-making without the latency you normally experience with most AI models.
Once I got access, I immediately wanted to incorporate it into my bug bounty workflow. I wanted something that could handle both large scale and high speed .
When you're doing bug bounty reconnaissance and enumeration, you can quickly end up with a massive number of websites and subdomains to process. Traditionally, you have a few options, each with its own downsides:
Manual review – Can take days to weeks to complete
Eyeballer – Takes minutes, but results can be mixed because of the limited model
Frontier models – Powerful, but expensive. They can consume millions of tokens just to review webpages
Local LLMs – Flexible, but slow, and they don't always have the context windows needed to process full webpages
With my first test using TypeSafe, I was able to process 17 subdomains per second .
That's approximately 1,040 domains per minute .
To accomplish this, I first came up with a set of questions broad enough to apply to almost every website.
Does the visible page text below contain enough information to judge what this page actually is?
This helps eliminate false negatives. Some websites time out or don't load correctly because many of them are non-production environments.
What kind of page is this, based on its title, headers, and visible text?
Clearly, in a bug bounty situation, you want to pull out things like S3 buckets, directory listings, and internal tools. "Broken" is a bit too broad of a category, and I would probably break it up further in future versions.
Does this look like an internal admin, staff, or Ops tool that wasn't meant to be public?
Pretty self-explanatory. We're continuing to look for accidental exposure.
Does this look like a login or management tool for a network security application, VPN gateway, or remote access product?
While many of these products could be classified as internal tools, we want to specifically call them out because they are often intentionally exposed to the internet.
Does this look like the login page of an IT management, RMM, or monitoring tool?
This includes products like Grafana, SolarWinds, Splunk, etc.
Does this response directly expose sensitive data itself rather than just linking to it?
Because TypeSafe is looking at the HTML itself rather than screenshots, we can quickly identify things like API keys sitting directly on a landing page. This would be rare, but it's also something that's too easy to overlook.
How interesting is this host for security testing?
This is the first time we're using the Score question type. It gives us a useful number to work with when determining whether a host is worth investigating further.
The responses were saved to a CSV file, making them easy to review and even easier for AI to filter through. Once I had the results, I had Claude take the top 20 highest-scoring results across different categories and put them into text files.
I use the Chrome extension Open Multiple URLs to quickly review the results manually. After a quick review, I pass the interesting results to Claude for further analysis to look for sensitive or exposed content.
This was just a quick post to show what Jev is able to do. I think Jev is going to be a valuable tool for anyone in the offensive security space, and I can't wait to see some of the other projects people build with it. We're already starting to see them pop up.
https://github.com/browser-use/jev-ultrafast
This was only my first day using Jev, but working with this speed, cheap cost and flexibility is what sells me.
Instead of sending thousands of hosts directly to an expensive frontier model or spending hours manually reviewing them, Jev can act as a high-speed filtering layer. This is a must have for your enumeration layer at scale.
For offensive security, I think that's where this gets really interesting. The goal isn't to replace manual testing or more powerful AI models. It's to make sure you're spending your time and compute on the things that actually deserve it.
Go ahead and check it out. Let me know if you make any cool offensive security projects with it.
https://typesafe.ai/
Cyber Security - Technology - Life
