---
source: "https://citegeo.uvansa.com"
hn_url: "https://news.ycombinator.com/item?id=49826945"
title: "Show HN: I asked 5 LLMs 10k questions and indexed every domain they cited"
article_title: "CiteGEO: open source GEO and AI visibility tracking"
image: ""
author: "jhaankit373"
captured_at: "2026-09-24T07:17:27Z"
capture_tool: "hn-digest"
hn_id: 49826945
score: 1
comments: 1
posted_at: "2026-09-24T06:27:32Z"
tags:
  - hacker-news
---

# Show HN: I asked 5 LLMs 10k questions and indexed every domain they cited

- HN: [49826945](https://news.ycombinator.com/item?id=49826945)
- Source: [citegeo.uvansa.com](https://citegeo.uvansa.com)
- Score: 1
- Comments: 1
- Posted: 2026-09-24T06:27:32Z

## Translation

Title: Show HN: I asked 5 LLMs 10k questions and indexed every domain they cited
Article title: CiteGEO: open source GEO and AI visibility tracking
Description: Self-hosted generative engine optimization (GEO) and answer engine optimization (AEO) tracking. Ask ChatGPT, Claude, Gemini and Perplexity the questions your buyers ask, keep every raw answer, and see the evidence behind every number.

Article text:
Skip to content CiteGEO CiteGEO emblem: two angled bars forming a quotation mark citegeo How it works What you get Sources What is GEO Contact GitHub GEO tracking you run yourself
Five categories. One brand never came up.
Engineering software, Indian equities, Kubernetes monitoring, UK accounting, running shoes. Five real questions put to a real model, kept word for word. Pick one and read what it said, which brands it chose to name, and in what order it named them.
Engineering software Linear what should a five person startup use to track engineering work not named
Indian equities Screener.in what is the best stock screening website for indian stocks named 1st
Developer infrastructure Grafana how do i monitor a kubernetes cluster without paying for datadog named 2nd
UK accounting FreeAgent best accounting software for a uk sole trader named 3rd
Running shoes HOKA best running shoes for flat feet and long distance named 5th
Sage Business Cloud Accounting
New Balance Fresh Foam X 860v13
Search could be counted. An answer cannot.
Ten ranked links, a position you could track, and a visit that arrived with a referrer attached. Every part of that loop was countable.
02 An answer replaces all of it
The assistant reads the pages for your buyer and writes one paragraph naming two or three products. No list, no position, and usually no click.
03 Nothing tells you when you are left out
A rank tracker watches a page your buyer never opened. Analytics can only count visits that happened. The answer that omitted you sends nothing at all.
04 So put the question yourself
CiteGEO asks the models what your buyers ask, keeps every answer whole, and links every figure back to the sentence it was counted from.
A score is a summary of something, and you cannot audit a summary. CiteGEO writes down the answer itself: the question that produced it, every brand it named, the model, the date, and the sentences in the order they arrived. The corridor behind is one card per answer already on disk.
what should a five person startup use to track engineering work
Tracking engineering work efficiently is crucial for a startup, especially with a small team of five where communication and transparency are key. Startups typically need something lightweig
what is the best stock screening website for indian stocks
There are several stock screening websites and platforms specifically tailored for analyzing Indian stocks, each offering unique features for different types of investors. 1. Screener.in is
how do i monitor a kubernetes cluster without paying for datadog
Monitoring a Kubernetes cluster without paying for tools like Datadog is entirely possible using open-source software and other free options. 1. Use Prometheus for metrics aggregation. Prome
best accounting software for a uk sole trader
Choosing the best accounting software for a UK sole trader depends on your specific needs, preferences, and budget, but here are some excellent options tailored for sole traders. 1. QuickBoo
best running shoes for flat feet and long distance
If you have flat feet and are looking to run long distances, it is essential to choose running shoes with proper arch support, stability, and cushioning to prevent discomfort and injuries. S
One model is an opinion. Five of them are a measurement.
Take one of those five categories and go deeper. The same question, put to five providers within the same minute. Each one came back with a different shortlist. Each shortlist makes a number, and the numbers do not agree.
Put to all five, same run what is the best stock screening website for indian stocks
anthropic your key 1 Screener.in
perplexity your key 1 Screener.in
Read the anthropic run on its own and you would report that you lead the category. Read the google run and you would report that you do not exist. Both figures are true about a model. Neither is true about the brand.
What the run actually measured
44.6 mean across five providers
0 to 81 the spread it came from
This is why the provider list is the method rather than a compatibility table. OpenAI, Anthropic, Gemini, Perplexity, an Azure OpenAI deployment, or a model rented through a cloud account you already hold on Bedrock, Vertex AI, Databricks or watsonx: CiteGEO puts the question to every provider you configure and keeps the spread attached to the mean. A brand at 81 on one provider and 0 on another has a different problem from one sitting at 45 everywhere, and a single figure hides which of the two you have.
Move them and watch the score lie.
The run above already happened, so it stays as it came back. These five are yours to drag. Everything underneath recomputes as you move them, which is the argument in one control: the mean is not the finding, the distance between the ends is.
An answer is assembled from pages somebody already wrote.
A model with search on does not invent the shortlist. It reads the open web and returns the pages it used. You cannot edit the answer, but you can change what it reads, and that list is the only actionable thing in this whole category.
Community forums, threads, question sites
Reference encyclopaedias and documentation
Comparisons review sites and roundups
Pages you own docs, pricing, guides
The answer Three names, one paragraph, no list and no position.
a forum thread comparing the tools
a roundup post naming five of them
CiteGEO keeps the sources a provider returns apart from URLs that merely appear in the answer text, because the two are not the same evidence. Then it names the pages worth going after.
The pages you have to get into
Every page a model cited for one of your topics, with the question it was cited for. An opening is one a rival won on a question that never named you at all. That is the shortest route to being in the answer: the page already ranks with the model, so earn a mention on it or write the one that replaces it.
What kind of page wins this topic
This is the part worth acting on. Chasing five URLs is a week of outreach. Knowing that four in ten citations on this topic are forum threads tells you what to make, and a model that reaches for one thread will reach for the next one like it.
Four of these five are pages you do not control, and every one of them was cited on a question where no answer named you. Nearby topics on the same page count too: a model that reaches for a page once tends to reach for it again.
Measurement you can argue with.
An absence is reported as an absence. Never as a zero, never as a default, never as a plausible looking value.
New to the terms? What generative engine optimization means, in plain words
Four commands
and it is yours.
You hold the provider key, you pay the provider directly, and the answers stay on your disk. CiteGEO adds no cost of its own and sends nothing anywhere except to the APIs you configure.
$ git clone https://github.com/ankit373/citegeo
$ npm ci
$ cp .env.example .env
$ npm run server
listening on http://127.0.0.1:8787 Contact
There is one person at the other end.
Questions about running it, a provider you want supported, a figure that looks wrong, or a patch worth reviewing. Mail reaches the person who maintains CiteGEO rather than a queue, so say what you are trying to measure and you will get a useful answer back.
Running it Setup, provider keys, self-hosting, anything in the readme that did not survive contact with your machine. Mail about setup
A provider to add A model or gateway you want measured alongside the rest. Bedrock arrived this way. Request a provider
A number you doubt Every figure links to the answers behind it, so send the figure and the answer you think it got wrong. Challenge a figure
Open source GEO and AI visibility tracking. MIT licensed.

## Original Extract

Self-hosted generative engine optimization (GEO) and answer engine optimization (AEO) tracking. Ask ChatGPT, Claude, Gemini and Perplexity the questions your buyers ask, keep every raw answer, and see the evidence behind every number.

Skip to content CiteGEO CiteGEO emblem: two angled bars forming a quotation mark citegeo How it works What you get Sources What is GEO Contact GitHub GEO tracking you run yourself
Five categories. One brand never came up.
Engineering software, Indian equities, Kubernetes monitoring, UK accounting, running shoes. Five real questions put to a real model, kept word for word. Pick one and read what it said, which brands it chose to name, and in what order it named them.
Engineering software Linear what should a five person startup use to track engineering work not named
Indian equities Screener.in what is the best stock screening website for indian stocks named 1st
Developer infrastructure Grafana how do i monitor a kubernetes cluster without paying for datadog named 2nd
UK accounting FreeAgent best accounting software for a uk sole trader named 3rd
Running shoes HOKA best running shoes for flat feet and long distance named 5th
Sage Business Cloud Accounting
New Balance Fresh Foam X 860v13
Search could be counted. An answer cannot.
Ten ranked links, a position you could track, and a visit that arrived with a referrer attached. Every part of that loop was countable.
02 An answer replaces all of it
The assistant reads the pages for your buyer and writes one paragraph naming two or three products. No list, no position, and usually no click.
03 Nothing tells you when you are left out
A rank tracker watches a page your buyer never opened. Analytics can only count visits that happened. The answer that omitted you sends nothing at all.
04 So put the question yourself
CiteGEO asks the models what your buyers ask, keeps every answer whole, and links every figure back to the sentence it was counted from.
A score is a summary of something, and you cannot audit a summary. CiteGEO writes down the answer itself: the question that produced it, every brand it named, the model, the date, and the sentences in the order they arrived. The corridor behind is one card per answer already on disk.
what should a five person startup use to track engineering work
Tracking engineering work efficiently is crucial for a startup, especially with a small team of five where communication and transparency are key. Startups typically need something lightweig
what is the best stock screening website for indian stocks
There are several stock screening websites and platforms specifically tailored for analyzing Indian stocks, each offering unique features for different types of investors. 1. Screener.in is
how do i monitor a kubernetes cluster without paying for datadog
Monitoring a Kubernetes cluster without paying for tools like Datadog is entirely possible using open-source software and other free options. 1. Use Prometheus for metrics aggregation. Prome
best accounting software for a uk sole trader
Choosing the best accounting software for a UK sole trader depends on your specific needs, preferences, and budget, but here are some excellent options tailored for sole traders. 1. QuickBoo
best running shoes for flat feet and long distance
If you have flat feet and are looking to run long distances, it is essential to choose running shoes with proper arch support, stability, and cushioning to prevent discomfort and injuries. S
One model is an opinion. Five of them are a measurement.
Take one of those five categories and go deeper. The same question, put to five providers within the same minute. Each one came back with a different shortlist. Each shortlist makes a number, and the numbers do not agree.
Put to all five, same run what is the best stock screening website for indian stocks
anthropic your key 1 Screener.in
perplexity your key 1 Screener.in
Read the anthropic run on its own and you would report that you lead the category. Read the google run and you would report that you do not exist. Both figures are true about a model. Neither is true about the brand.
What the run actually measured
44.6 mean across five providers
0 to 81 the spread it came from
This is why the provider list is the method rather than a compatibility table. OpenAI, Anthropic, Gemini, Perplexity, an Azure OpenAI deployment, or a model rented through a cloud account you already hold on Bedrock, Vertex AI, Databricks or watsonx: CiteGEO puts the question to every provider you configure and keeps the spread attached to the mean. A brand at 81 on one provider and 0 on another has a different problem from one sitting at 45 everywhere, and a single figure hides which of the two you have.
Move them and watch the score lie.
The run above already happened, so it stays as it came back. These five are yours to drag. Everything underneath recomputes as you move them, which is the argument in one control: the mean is not the finding, the distance between the ends is.
An answer is assembled from pages somebody already wrote.
A model with search on does not invent the shortlist. It reads the open web and returns the pages it used. You cannot edit the answer, but you can change what it reads, and that list is the only actionable thing in this whole category.
Community forums, threads, question sites
Reference encyclopaedias and documentation
Comparisons review sites and roundups
Pages you own docs, pricing, guides
The answer Three names, one paragraph, no list and no position.
a forum thread comparing the tools
a roundup post naming five of them
CiteGEO keeps the sources a provider returns apart from URLs that merely appear in the answer text, because the two are not the same evidence. Then it names the pages worth going after.
The pages you have to get into
Every page a model cited for one of your topics, with the question it was cited for. An opening is one a rival won on a question that never named you at all. That is the shortest route to being in the answer: the page already ranks with the model, so earn a mention on it or write the one that replaces it.
What kind of page wins this topic
This is the part worth acting on. Chasing five URLs is a week of outreach. Knowing that four in ten citations on this topic are forum threads tells you what to make, and a model that reaches for one thread will reach for the next one like it.
Four of these five are pages you do not control, and every one of them was cited on a question where no answer named you. Nearby topics on the same page count too: a model that reaches for a page once tends to reach for it again.
Measurement you can argue with.
An absence is reported as an absence. Never as a zero, never as a default, never as a plausible looking value.
New to the terms? What generative engine optimization means, in plain words
Four commands
and it is yours.
You hold the provider key, you pay the provider directly, and the answers stay on your disk. CiteGEO adds no cost of its own and sends nothing anywhere except to the APIs you configure.
$ git clone https://github.com/ankit373/citegeo
$ npm ci
$ cp .env.example .env
$ npm run server
listening on http://127.0.0.1:8787 Contact
There is one person at the other end.
Questions about running it, a provider you want supported, a figure that looks wrong, or a patch worth reviewing. Mail reaches the person who maintains CiteGEO rather than a queue, so say what you are trying to measure and you will get a useful answer back.
Running it Setup, provider keys, self-hosting, anything in the readme that did not survive contact with your machine. Mail about setup
A provider to add A model or gateway you want measured alongside the rest. Bedrock arrived this way. Request a provider
A number you doubt Every figure links to the answers behind it, so send the figure and the answer you think it got wrong. Challenge a figure
Open source GEO and AI visibility tracking. MIT licensed.
