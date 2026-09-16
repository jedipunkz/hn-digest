---
source: "https://buttondown.com/hillelwayne/archive/the-llms-yearn-for-the-spines/"
hn_url: "https://news.ycombinator.com/item?id=49733100"
title: "The LLMs Yearn for the Spines"
article_title: "The LLMs yearn for the spines • Buttondown"
image: "https://assets.buttondown.email/images/ba001783-041b-4c65-998c-4d8ff9bbc3e9.png?w=960&fit=max"
author: "srijan4"
captured_at: "2026-09-16T21:55:39Z"
capture_tool: "hn-digest"
hn_id: 49733100
score: 5
comments: 0
posted_at: "2026-09-16T21:13:10Z"
tags:
  - hacker-news
---

# The LLMs Yearn for the Spines

- HN: [49733100](https://news.ycombinator.com/item?id=49733100)
- Source: [buttondown.com](https://buttondown.com/hillelwayne/archive/the-llms-yearn-for-the-spines/)
- Score: 5
- Comments: 0
- Posted: 2026-09-16T21:13:10Z

## Translation

Title: The LLMs Yearn for the Spines
Article title: The LLMs yearn for the spines • Buttondown
Description: You can't get them to stop talking about spines!

Article text:
The LLMs yearn for the spines • Buttondown
Computer Things
Archives
Search...
Log in
Subscribe
September 16, 2026
The LLMs yearn for the spines
You can't get them to stop talking about spines!
Earlier this year I worked on a couple of AI-generated TLA+ specifications, and one thing I've noticed was that they all used the word "spine" somewhere. Then I saw some non-TLA+ public projects use "spine" too and got curious if it was a new LLMism (LLisM?).
If it is, we should see it appear in a discontinuous jump in the number of code projects that use "spine". The proper way to thoroughly explore this would be via the GitHub archive data dump but that'd take forever and I have a full time job now. Instead, I'm going to use GitHub Search . We can't filter code samples by date but can filter pull requests, which seems like an acceptable enough proxy.
Here's a graph of the number of public PRs each year with "spine" in the title:
I stopped 2026 at 2026-09-01, meaning the first nine months of this year have seen 20 times more "spine" PRs as all of 2025. Now it could be the case that that GitHub grew, like, a lot in the last year. Maybe the 20x increase is due there being 20x as many pull requests. We can account for this by comparing it against all public PRs : 1
If "spine" wasn't an LLisM, we would expect only a 1.5xish increase from 2025, not a 20x increase. One other possible confounder could be that LLMs don't like the word spine but PRs Georg submitted 15,000 changes to a chiropractor repo. I will deal with this issue by ignoring it.
Can we figure out which models specifically are obsessed with spines? Let's try breaking the change down by month:
Looking at LLM timeline , the only widely-used model released in May was Opus 4.8, which came out too late in the month to explain the spike. GPT 5.5 came out in April, which could be the cause. But I'm not convinced by this. For one, the spine count is still 1.5xing each month before April, and I definitely remember seeing spines as early as February. Also, I spot checked and many of the PRs that use "spine" are coauthored by Claude or Cursor. So I think it's not specific to one LLM, but could be convinced otherwise.
Avdi isn't crazy, LLMs are obsessed with gates:
That's one out of every 300 PRs, BTW. If we include PRs where "gate" appears anywhere , that number jumps to 1 in 26 .
The reason I'm only looking at titles is because if we include PR bodies, we get an artifact when searching for "ladder". In 2022, over 30,000 PRs used the word "ladder". Looking at a sample , almost all of them are due to automated upgrades of boto3, which has "ABR package ladders" in the upgrade description. I felt really good about finding this until, out of curiosity, I asked a cleanroom LLM to "figure out the spike". It got the same answer much faster 😢
Anyway, looking at just PR titles removes that artifact for "ladder":
I was going to make an "LLM love ladder" joke, but turns out I was misremembering and the meme is "I love lamp". I checked, they don't love "lamp" any more than the average dev.
They do, however, love "lane":
Maybe this is explained by "CPU lanes"? I dunno.
Since formal verification is a hot topic in AI, of course we see a lot more PRs talking 'bout proofs:
This isn't as big of an increase, possibly because we already had "proof of work". What is a big increase is truth:
Finally, while I was writing this, a coworker mentioned she saw the word "seam" way more often, so I checked that too:
I think with some more time and sophistication, it should be possible to figure out which words are particular to specific LLMs. I'm not that sophisticated, though!
I put the tool I vibecoded to do this analysis up as a gist . Go tell me how cool I am for making reproducible datasets and/or how I screwed everything up by making an obvious mistake.
Systems Distributed Talk Livewatch
My Systems Distributed talk Logic for Programmers is going live Sept 23 ! I'll be doing a livechat and answering questions from viewers as it plays, and giving background notes on the Making Of. Should be starting 11 AM central / noon Eastern.
Gergely Orosz has insider info showing that GitHub has seen 71.7M more LLM-authored PRs this year than last year. I believe this is both public and private repos, meaning that there is probably now more agentic activity on GitHub than public activity. ↩
If you're reading this on the web, you can subscribe here. Updates are once a week. My main website is here .
Logic for Programmers is now available in print !

## Original Extract

You can't get them to stop talking about spines!

The LLMs yearn for the spines • Buttondown
Computer Things
Archives
Search...
Log in
Subscribe
September 16, 2026
The LLMs yearn for the spines
You can't get them to stop talking about spines!
Earlier this year I worked on a couple of AI-generated TLA+ specifications, and one thing I've noticed was that they all used the word "spine" somewhere. Then I saw some non-TLA+ public projects use "spine" too and got curious if it was a new LLMism (LLisM?).
If it is, we should see it appear in a discontinuous jump in the number of code projects that use "spine". The proper way to thoroughly explore this would be via the GitHub archive data dump but that'd take forever and I have a full time job now. Instead, I'm going to use GitHub Search . We can't filter code samples by date but can filter pull requests, which seems like an acceptable enough proxy.
Here's a graph of the number of public PRs each year with "spine" in the title:
I stopped 2026 at 2026-09-01, meaning the first nine months of this year have seen 20 times more "spine" PRs as all of 2025. Now it could be the case that that GitHub grew, like, a lot in the last year. Maybe the 20x increase is due there being 20x as many pull requests. We can account for this by comparing it against all public PRs : 1
If "spine" wasn't an LLisM, we would expect only a 1.5xish increase from 2025, not a 20x increase. One other possible confounder could be that LLMs don't like the word spine but PRs Georg submitted 15,000 changes to a chiropractor repo. I will deal with this issue by ignoring it.
Can we figure out which models specifically are obsessed with spines? Let's try breaking the change down by month:
Looking at LLM timeline , the only widely-used model released in May was Opus 4.8, which came out too late in the month to explain the spike. GPT 5.5 came out in April, which could be the cause. But I'm not convinced by this. For one, the spine count is still 1.5xing each month before April, and I definitely remember seeing spines as early as February. Also, I spot checked and many of the PRs that use "spine" are coauthored by Claude or Cursor. So I think it's not specific to one LLM, but could be convinced otherwise.
Avdi isn't crazy, LLMs are obsessed with gates:
That's one out of every 300 PRs, BTW. If we include PRs where "gate" appears anywhere , that number jumps to 1 in 26 .
The reason I'm only looking at titles is because if we include PR bodies, we get an artifact when searching for "ladder". In 2022, over 30,000 PRs used the word "ladder". Looking at a sample , almost all of them are due to automated upgrades of boto3, which has "ABR package ladders" in the upgrade description. I felt really good about finding this until, out of curiosity, I asked a cleanroom LLM to "figure out the spike". It got the same answer much faster 😢
Anyway, looking at just PR titles removes that artifact for "ladder":
I was going to make an "LLM love ladder" joke, but turns out I was misremembering and the meme is "I love lamp". I checked, they don't love "lamp" any more than the average dev.
They do, however, love "lane":
Maybe this is explained by "CPU lanes"? I dunno.
Since formal verification is a hot topic in AI, of course we see a lot more PRs talking 'bout proofs:
This isn't as big of an increase, possibly because we already had "proof of work". What is a big increase is truth:
Finally, while I was writing this, a coworker mentioned she saw the word "seam" way more often, so I checked that too:
I think with some more time and sophistication, it should be possible to figure out which words are particular to specific LLMs. I'm not that sophisticated, though!
I put the tool I vibecoded to do this analysis up as a gist . Go tell me how cool I am for making reproducible datasets and/or how I screwed everything up by making an obvious mistake.
Systems Distributed Talk Livewatch
My Systems Distributed talk Logic for Programmers is going live Sept 23 ! I'll be doing a livechat and answering questions from viewers as it plays, and giving background notes on the Making Of. Should be starting 11 AM central / noon Eastern.
Gergely Orosz has insider info showing that GitHub has seen 71.7M more LLM-authored PRs this year than last year. I believe this is both public and private repos, meaning that there is probably now more agentic activity on GitHub than public activity. ↩
If you're reading this on the web, you can subscribe here. Updates are once a week. My main website is here .
Logic for Programmers is now available in print !
