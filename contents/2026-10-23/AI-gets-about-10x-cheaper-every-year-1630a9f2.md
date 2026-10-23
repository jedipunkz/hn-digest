---
source: "https://stevekrouse.com/10x"
hn_url: "https://news.ycombinator.com/item?id=50215307"
title: "AI gets about 10x cheaper every year"
article_title: "AI gets about 10x cheaper every year"
image: ""
author: "stevekrouse"
captured_at: "2026-10-23T21:52:39Z"
capture_tool: "hn-digest"
hn_id: 50215307
score: 1
comments: 0
posted_at: "2026-10-23T21:04:14Z"
tags:
  - hacker-news
---

# AI gets about 10x cheaper every year

- HN: [50215307](https://news.ycombinator.com/item?id=50215307)
- Source: [stevekrouse.com](https://stevekrouse.com/10x)
- Score: 1
- Comments: 0
- Posted: 2026-10-23T21:04:14Z

## Translation

Title: AI gets about 10x cheaper every year

Article text:
Skip to content Steve Krouse Essays Notes Podcasts Quotes Books AI gets about 10x cheaper every year
The cost to use a given level of AI falls about 10x every 12 months.
Now in October 2026, 20 months later, I wanted to see if that was still true, or if the rate had sped up or slowed down. 1
Living on an exponential, it's really hard to get an intuition for it. It helps to zoom out and make charts.
I charted the price decreases using four different benchmarks:
I confirmed Sam's claim: AI continues to get about 10x cheaper every year.
Humanity's Last Exam (HLE) tests AI models with difficult expert-written questions across a wide range of subjects.
In this chart, each colored line represents a different level of intelligence on HLE. For each level of intelligence, you can walk along the graph to see subsequent models that would give you the same level of intelligence, or higher, but cheaper.
For example, start with o3-mini, released in January 2025. It scored 12% on HLE and launched at $1.93 per million tokens. 2
Just over a year later, DeepSeek V4 Flash offered that level of intelligence for about 18 cents per million tokens. That's 11x cheaper.
But when you play it out further, the story isn't as clear. By September 2026, MiMo V2.6 Flash still cost about 18 cents per million tokens. That's only 11x cheaper in about 1.7 years, rather than the roughly 50x the rule would predict. I hypothesize that this flattening happens when things get too cheap to meter, and people don't care about making arbitrary, old levels of intelligence even cheaper. 3
Averaging the five levels of intelligence gives 10x cheaper every 0.8 years .
Elo measures which model's answers people prefer in head-to-head comparisons. It lets us look back further, to March 2023.
Here the fitted rate is 10x cheaper every 1.3 years , which again might be slower because the oldest levels of intelligence aren't continuing to get cheaper. 3
GPQA Diamond is a test of graduate-level science questions. The question here is the same: how much does it cost to get a model that scores at least 30%, 40%, 50%, and so on?
Terminal-Bench measures how well models complete tasks in a terminal.
I admire Moore's Law and think we should have another law people can point to about how fast AI is getting too cheap to meter.
I'd propose “Lee's Law.” The earliest prediction I found was Kai-Fu Lee predicting inference costs would get 10x cheaper a year in May 2024 . 4 Although there's a strong argument for calling it “Altman's Law,” because he's helped popularize the idea, as well as the phrase “intelligence too cheap to meter” .
For context, Moore's Law says the number of transistors on a chip doubles every two years. That pace of improvement has utterly changed our industry and world.
Lee's Law is 7x faster. Which is insane. Lee's Law makes Moore's Law look like a line.
Thanks to Charlie , Tom , and Jackson for feedback on this post.
I've long been curious about if these trends are continuing. My friend Swyx made this chart , which I was constantly referencing. But I couldn't keep relying on him to keep it up to date, and I wanted to explore the data and slice and dice it my own way. But who has the time? With GPT-6 Astra Ultrafast, I was able to do it. That's a big step up from even a month ago, when every edit to a chart took 5 minutes, which meant my feedback loop was so slow that a project would take all day or I'd give up. Now an edit is 10 seconds, so I can iterate to a non-slop version of what I want in minutes, and accomplish this whole project in mere hours, mostly between meetings. Chart data , dated prices and sources , calculations .
Prices use a mix of 3:1 input to output tokens, without caching.
If someone wanted to just hit o3-mini's 12% on HLE, I bet they could make a model that is even cheaper than MiMo2.6F (35% HLE), which is already almost too cheap to meter. And anyways, everyone I know would spend the extra 6 cents for GLM 5.3 Flash (40% HLE). My guess is that these older levels could keep getting 10x cheaper if the market cared.
Guido Appenzeller later charted the trend and called it LLMflation. People were already calling it Altman's Law in February 2025 . Earlier charts include ARK's January 2024 report and Charles Frye's May 2024 chart . Swyx's January 2025 chart used chat Elo. Epoch's September 2026 study also accounts for how many tokens models use, and finds a similar rate.

## Original Extract

Skip to content Steve Krouse Essays Notes Podcasts Quotes Books AI gets about 10x cheaper every year
The cost to use a given level of AI falls about 10x every 12 months.
Now in October 2026, 20 months later, I wanted to see if that was still true, or if the rate had sped up or slowed down. 1
Living on an exponential, it's really hard to get an intuition for it. It helps to zoom out and make charts.
I charted the price decreases using four different benchmarks:
I confirmed Sam's claim: AI continues to get about 10x cheaper every year.
Humanity's Last Exam (HLE) tests AI models with difficult expert-written questions across a wide range of subjects.
In this chart, each colored line represents a different level of intelligence on HLE. For each level of intelligence, you can walk along the graph to see subsequent models that would give you the same level of intelligence, or higher, but cheaper.
For example, start with o3-mini, released in January 2025. It scored 12% on HLE and launched at $1.93 per million tokens. 2
Just over a year later, DeepSeek V4 Flash offered that level of intelligence for about 18 cents per million tokens. That's 11x cheaper.
But when you play it out further, the story isn't as clear. By September 2026, MiMo V2.6 Flash still cost about 18 cents per million tokens. That's only 11x cheaper in about 1.7 years, rather than the roughly 50x the rule would predict. I hypothesize that this flattening happens when things get too cheap to meter, and people don't care about making arbitrary, old levels of intelligence even cheaper. 3
Averaging the five levels of intelligence gives 10x cheaper every 0.8 years .
Elo measures which model's answers people prefer in head-to-head comparisons. It lets us look back further, to March 2023.
Here the fitted rate is 10x cheaper every 1.3 years , which again might be slower because the oldest levels of intelligence aren't continuing to get cheaper. 3
GPQA Diamond is a test of graduate-level science questions. The question here is the same: how much does it cost to get a model that scores at least 30%, 40%, 50%, and so on?
Terminal-Bench measures how well models complete tasks in a terminal.
I admire Moore's Law and think we should have another law people can point to about how fast AI is getting too cheap to meter.
I'd propose “Lee's Law.” The earliest prediction I found was Kai-Fu Lee predicting inference costs would get 10x cheaper a year in May 2024 . 4 Although there's a strong argument for calling it “Altman's Law,” because he's helped popularize the idea, as well as the phrase “intelligence too cheap to meter” .
For context, Moore's Law says the number of transistors on a chip doubles every two years. That pace of improvement has utterly changed our industry and world.
Lee's Law is 7x faster. Which is insane. Lee's Law makes Moore's Law look like a line.
Thanks to Charlie , Tom , and Jackson for feedback on this post.
I've long been curious about if these trends are continuing. My friend Swyx made this chart , which I was constantly referencing. But I couldn't keep relying on him to keep it up to date, and I wanted to explore the data and slice and dice it my own way. But who has the time? With GPT-6 Astra Ultrafast, I was able to do it. That's a big step up from even a month ago, when every edit to a chart took 5 minutes, which meant my feedback loop was so slow that a project would take all day or I'd give up. Now an edit is 10 seconds, so I can iterate to a non-slop version of what I want in minutes, and accomplish this whole project in mere hours, mostly between meetings. Chart data , dated prices and sources , calculations .
Prices use a mix of 3:1 input to output tokens, without caching.
If someone wanted to just hit o3-mini's 12% on HLE, I bet they could make a model that is even cheaper than MiMo2.6F (35% HLE), which is already almost too cheap to meter. And anyways, everyone I know would spend the extra 6 cents for GLM 5.3 Flash (40% HLE). My guess is that these older levels could keep getting 10x cheaper if the market cared.
Guido Appenzeller later charted the trend and called it LLMflation. People were already calling it Altman's Law in February 2025 . Earlier charts include ARK's January 2024 report and Charles Frye's May 2024 chart . Swyx's January 2025 chart used chat Elo. Epoch's September 2026 study also accounts for how many tokens models use, and finds a similar rate.
