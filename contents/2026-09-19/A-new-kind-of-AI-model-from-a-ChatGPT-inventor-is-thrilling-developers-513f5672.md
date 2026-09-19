---
source: "https://techcrunch.com/2026/09/18/a-new-kind-of-ai-model-from-a-chatgpt-inventor-is-thrilling-developers/"
hn_url: "https://news.ycombinator.com/item?id=49769222"
title: "A new kind of AI model from a ChatGPT inventor is thrilling developers"
article_title: "A new kind of AI model from a ChatGPT inventor is thrilling developers | TechCrunch"
image: "https://techcrunch.com/wp-content/uploads/2026/09/typesafe-ai.jpg?resize=1174,1200"
author: "andsoitis"
captured_at: "2026-09-19T19:11:33Z"
capture_tool: "hn-digest"
hn_id: 49769222
score: 1
comments: 0
posted_at: "2026-09-19T18:56:08Z"
tags:
  - hacker-news
---

# A new kind of AI model from a ChatGPT inventor is thrilling developers

- HN: [49769222](https://news.ycombinator.com/item?id=49769222)
- Source: [techcrunch.com](https://techcrunch.com/2026/09/18/a-new-kind-of-ai-model-from-a-chatgpt-inventor-is-thrilling-developers/)
- Score: 1
- Comments: 0
- Posted: 2026-09-19T18:56:08Z

## Translation

Title: A new kind of AI model from a ChatGPT inventor is thrilling developers
Article title: A new kind of AI model from a ChatGPT inventor is thrilling developers | TechCrunch
Description: Jev, a new kind of AI model, is showing developers a cheaper and faster path to software intelligence.

Article text:
Disrupt 2026: OpenAI, Anthropic, Replit, and more take over 6 industry stages. 25% off tickets now
Back by popular demand: Save up to $300 on Disrupt
TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
A new kind of AI model from a ChatGPT inventor is thrilling developers
ChatGPT broke Diogo Almeida’s heart.
Almeida was an OpenAI researcher who helped build the chatbot and then invent reinforcement learning from human feedback (RLHF), the model-training technique perhaps most responsible for our current age of AI. But despite its capabilities, he was disappointed.
“We have lightning in a bottle, and yet it is not useful,” Almeida told TechCrunch. “I’ve been battling that problem since then. It took me a while to come to the conclusion: The problem is we are optimizing for human language … We have been super good at human language for four years, but it’s not useful for automation because computers speak a different language.”
Two years ago, Almeida left OpenAI to start TypeSafe AI , a startup trying to fix that problem. This week, the company released a new transformer-based model, Jev , that is not a large language model (LLM). It doesn’t output text, but instead produces probabilities, or what the company calls “calibrated decisions.”
Eschewing language does a few things: It makes the model incredibly cheap and fast, and because users define the outputs in advance, it cannot hallucinate. Its output tokens are free, and input tokens are metered by the billion, not the million.
Developers are taking a great interest in the product; the company briefly lost the ability to serve users from its API because demand was so high. Jev appears most useful for software automation. Thus far, software developers see it as a cheaper and more robust way to incorporate intelligence into their code.
For example, Pranit Sharma, a software engineer at Vercel, a company making agentic infrastructure, said his company had used OpenAI’s ChatGPT Luna 5.6 to run a classifier to review commands for safety. When Vercel replaced OpenAI’s Luna with Jev, it got results five to 18 times more quickly and with greater accuracy.
Another developer, Bryo AI CTO Nikhil Mudholkar, tested Jev against Gemini for classifying business emails. In his test, Gemini was slightly more accurate, but 10 to 20 times more expensive. More interesting to Mudholkar were Jev’s confidence scores — “it is the only one that hands back a real probability which makes it ideal for automating workflows!!”
Besides replacing LLMs in certain use cases, the new model can also augment them, acting as a smart check on misbehavior. Using agents to monitor agents can quickly become expensive, but using Jev to do so, Almeida argues, makes sense. He sees users deploying Jev to track LLM agent traces and prevent jailbreaks.
“At the end of the day, it delegates the hallucination problem a little bit to the user,” explained Armin Ronacher, the CTO of Earendil, which builds the open source model harness Pi. “The user has to say, okay, if this only comes back with 50% probability, maybe this is a coin toss, and I disregard it. But if it’s 95%, sure, then I can do something with it.”
Another potential use for Jev is model routing, Ronacher said. Predicting whether a given workload requires a specific model would be useful, but using an LLM for the job would be expensive. Jev’s low cost and speed make that kind of real-time sorting possible.
And that’s Almeida’s hope. The model is named after William Stanley Jevons, the 19th-century economist whose eponymous paradox describes how the falling cost of a commodity can lead to it being used more and more. In this case, the falling cost of intelligence should lead to its widespread deployment.
“We think that there’s just going to be smart software all over the place in a way that’s emergent and distributed … much more like the early internet than you know like the mega apps that people are trying to build right now,” Almeida said.
Almeida is tight-lipped about the model’s architecture, which outside observers suspect is built on top of an open-weight LLM. The company refers to Jev as a “System One model,” focused on intuition rather than reasoning, and specifically focused on the right task. Almeida says Jev is trained exclusively on synthetic data using a technique he calls “reinforcement learning from calibrated decisions.”
“We made an early bet that we will be making all of our data, and that has been one of the best bets I’ve ever made in my life — better than our launch, in my opinion, better than RLHF,” he told TechCrunch. “Half of [our company] is a lab that basically owns this entire subfield of statistically well-understood synthetic data, and that is now my life joy.”
For now, Jev stands alone as this kind of model, but Ronacher expects that competitors will spring up now that its utility is apparent.
“We should have seen this earlier in many ways, but presumably because the LLMs are so cheap and subsidized, you often don’t have to be creative yet,” he said.
TypeSafe itself will be building more versions of the model, in new modalities. Asked if TypeSafe is a frontier lab, Almeida said, “the main product of frontier labs is fear or hype. I would like our main product to be intelligence…[but we are] not a lab in the sense of, you know, like bet on infinite wealth, or a religion, or building God in a data center, or whatever is the thing of today.”
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
October 13 – 15
San Francisco
Last day to book an exhibit table is September 18. Don’t miss out on high-impact leads, investor access, and a brand spotlight in Disrupt’s Expo Hall.
A new kind of AI model from a ChatGPT inventor is thrilling developers
OpenAI caught its models leaving notes to successors to hide bad behavior
Microsoft exec called AI scraping ‘the largest theft of labor in human history,’ new unredacted filings reveal
Clean tech startup Fluxnium found a way to tap 50,000 years’ worth of nuclear fuel
Salesforce and Nvidia’s new reasoning model is everything the AI labs should fear
Jensen Huang took a call from Trump, and showed off something else, too
The 9 buzziest startups from Y Combinator’s latest Demo Day, according to VCs

## Original Extract

Jev, a new kind of AI model, is showing developers a cheaper and faster path to software intelligence.

Disrupt 2026: OpenAI, Anthropic, Replit, and more take over 6 industry stages. 25% off tickets now
Back by popular demand: Save up to $300 on Disrupt
TechCrunch Desktop Logo
TechCrunch Mobile Logo
Latest
A new kind of AI model from a ChatGPT inventor is thrilling developers
ChatGPT broke Diogo Almeida’s heart.
Almeida was an OpenAI researcher who helped build the chatbot and then invent reinforcement learning from human feedback (RLHF), the model-training technique perhaps most responsible for our current age of AI. But despite its capabilities, he was disappointed.
“We have lightning in a bottle, and yet it is not useful,” Almeida told TechCrunch. “I’ve been battling that problem since then. It took me a while to come to the conclusion: The problem is we are optimizing for human language … We have been super good at human language for four years, but it’s not useful for automation because computers speak a different language.”
Two years ago, Almeida left OpenAI to start TypeSafe AI , a startup trying to fix that problem. This week, the company released a new transformer-based model, Jev , that is not a large language model (LLM). It doesn’t output text, but instead produces probabilities, or what the company calls “calibrated decisions.”
Eschewing language does a few things: It makes the model incredibly cheap and fast, and because users define the outputs in advance, it cannot hallucinate. Its output tokens are free, and input tokens are metered by the billion, not the million.
Developers are taking a great interest in the product; the company briefly lost the ability to serve users from its API because demand was so high. Jev appears most useful for software automation. Thus far, software developers see it as a cheaper and more robust way to incorporate intelligence into their code.
For example, Pranit Sharma, a software engineer at Vercel, a company making agentic infrastructure, said his company had used OpenAI’s ChatGPT Luna 5.6 to run a classifier to review commands for safety. When Vercel replaced OpenAI’s Luna with Jev, it got results five to 18 times more quickly and with greater accuracy.
Another developer, Bryo AI CTO Nikhil Mudholkar, tested Jev against Gemini for classifying business emails. In his test, Gemini was slightly more accurate, but 10 to 20 times more expensive. More interesting to Mudholkar were Jev’s confidence scores — “it is the only one that hands back a real probability which makes it ideal for automating workflows!!”
Besides replacing LLMs in certain use cases, the new model can also augment them, acting as a smart check on misbehavior. Using agents to monitor agents can quickly become expensive, but using Jev to do so, Almeida argues, makes sense. He sees users deploying Jev to track LLM agent traces and prevent jailbreaks.
“At the end of the day, it delegates the hallucination problem a little bit to the user,” explained Armin Ronacher, the CTO of Earendil, which builds the open source model harness Pi. “The user has to say, okay, if this only comes back with 50% probability, maybe this is a coin toss, and I disregard it. But if it’s 95%, sure, then I can do something with it.”
Another potential use for Jev is model routing, Ronacher said. Predicting whether a given workload requires a specific model would be useful, but using an LLM for the job would be expensive. Jev’s low cost and speed make that kind of real-time sorting possible.
And that’s Almeida’s hope. The model is named after William Stanley Jevons, the 19th-century economist whose eponymous paradox describes how the falling cost of a commodity can lead to it being used more and more. In this case, the falling cost of intelligence should lead to its widespread deployment.
“We think that there’s just going to be smart software all over the place in a way that’s emergent and distributed … much more like the early internet than you know like the mega apps that people are trying to build right now,” Almeida said.
Almeida is tight-lipped about the model’s architecture, which outside observers suspect is built on top of an open-weight LLM. The company refers to Jev as a “System One model,” focused on intuition rather than reasoning, and specifically focused on the right task. Almeida says Jev is trained exclusively on synthetic data using a technique he calls “reinforcement learning from calibrated decisions.”
“We made an early bet that we will be making all of our data, and that has been one of the best bets I’ve ever made in my life — better than our launch, in my opinion, better than RLHF,” he told TechCrunch. “Half of [our company] is a lab that basically owns this entire subfield of statistically well-understood synthetic data, and that is now my life joy.”
For now, Jev stands alone as this kind of model, but Ronacher expects that competitors will spring up now that its utility is apparent.
“We should have seen this earlier in many ways, but presumably because the LLMs are so cheap and subsidized, you often don’t have to be creative yet,” he said.
TypeSafe itself will be building more versions of the model, in new modalities. Asked if TypeSafe is a frontier lab, Almeida said, “the main product of frontier labs is fear or hype. I would like our main product to be intelligence…[but we are] not a lab in the sense of, you know, like bet on infinite wealth, or a religion, or building God in a data center, or whatever is the thing of today.”
When you purchase through links in our articles, we may earn a small commission . This doesn’t affect our editorial independence.
October 13 – 15
San Francisco
Last day to book an exhibit table is September 18. Don’t miss out on high-impact leads, investor access, and a brand spotlight in Disrupt’s Expo Hall.
A new kind of AI model from a ChatGPT inventor is thrilling developers
OpenAI caught its models leaving notes to successors to hide bad behavior
Microsoft exec called AI scraping ‘the largest theft of labor in human history,’ new unredacted filings reveal
Clean tech startup Fluxnium found a way to tap 50,000 years’ worth of nuclear fuel
Salesforce and Nvidia’s new reasoning model is everything the AI labs should fear
Jensen Huang took a call from Trump, and showed off something else, too
The 9 buzziest startups from Y Combinator’s latest Demo Day, according to VCs
